package assets

import (
	"github.com/danilo/scripts/github/dca/internal/dca"
)

// AssetHistoryModal represents a modal for displaying daily asset history
type AssetHistoryModal struct {
	AssetTicker   string
	Filename      string // Source filename for data loading (used by LoadMore)
	EntriesByDate []EntryByDate
	SelectedIndex int
	Loaded        bool
	Error         error
	Visible       bool

	// Pagination fields for infinite scroll
	Offset    int  // Number of days already loaded
	AllLoaded bool // True when all available data has been loaded
	Loading   bool // True when fetching more data
}

// EntryByDate represents entries aggregated by a single calendar date
type EntryByDate struct {
	Date             string
	TotalInvested    float64
	TotalShares      float64
	WeightedAvgPrice float64
	EntryCount       int
}

// NewAssetHistoryModal creates a new asset history modal
func NewAssetHistoryModal() *AssetHistoryModal {
	return &AssetHistoryModal{
		SelectedIndex: 0,
		Loaded:        false,
		Error:         nil,
		Visible:       false,
		Offset:        0,
		AllLoaded:     false,
		Loading:       false,
	}
}

// LoadData loads entries for the specified asset from the file and aggregates by date
func (m *AssetHistoryModal) LoadData(filename string, assetTicker string) error {
	// Load all entries from file
	data, err := dca.LoadEntries(filename)
	if err != nil {
		return err
	}

	// Get entries for the specific asset
	entries, exists := data.Entries[assetTicker]
	if !exists {
		// No entries for this asset - return empty data
		m.EntriesByDate = []EntryByDate{}
		m.Loaded = true
		m.Error = nil
		m.AllLoaded = true
		return nil
	}

	// Aggregate entries by date (all dates, we'll paginate on display)
	allEntries := AggregateByDate(entries)

	// Limit to first 10 days (batch size)
	const batchSize = 10
	if len(allEntries) > batchSize {
		m.EntriesByDate = allEntries[:batchSize]
	} else {
		m.EntriesByDate = allEntries
	}

	m.AssetTicker = assetTicker
	m.Filename = filename // Store filename for LoadMore
	m.Loaded = true
	m.Error = nil
	return nil
}

// LoadMore loads the next batch of 10 days for infinite scroll
func (m *AssetHistoryModal) LoadMore(filename string) error {
	if m.AllLoaded || m.Loading {
		return nil
	}

	m.Loading = true

	// Use stored filename with fallback to parameter for backward compatibility
	loadFilename := m.Filename
	if loadFilename == "" {
		loadFilename = filename
	}

	// Load all entries from file
	data, err := dca.LoadEntries(loadFilename)
	if err != nil {
		m.Loading = false
		m.Error = err
		return err
	}

	// Get entries for the specific asset
	entries, exists := data.Entries[m.AssetTicker]
	if !exists {
		m.AllLoaded = true
		m.Loading = false
		return nil
	}

	// Aggregate entries by date
	allEntries := AggregateByDate(entries)

	// Calculate how many days we already have loaded
	// Use len(EntriesByDate) as the current count
	currentCount := len(m.EntriesByDate)
	if currentCount >= len(allEntries) {
		m.AllLoaded = true
		m.Loading = false
		return nil
	}

	// Calculate how many more days we can load (batch of 10)
	const batchSize = 10
	remaining := len(allEntries) - currentCount
	toLoad := batchSize
	if remaining < batchSize {
		toLoad = remaining
	}

	// Append next batch
	if toLoad > 0 {
		m.EntriesByDate = append(m.EntriesByDate, allEntries[currentCount:currentCount+toLoad]...)
	}

	// Check if all data is now loaded
	if len(m.EntriesByDate) >= len(allEntries) {
		m.AllLoaded = true
	}

	m.Loading = false
	return nil
}

// AggregateByDate groups entries by calendar date and calculates daily metrics with
// cumulative weighted average. The cumulative weighted average at each day uses the
// PRD formula: sum(price_per_share × amount) / sum(amounts) for all entries up to that day.
func AggregateByDate(entries []dca.DCAEntry) []EntryByDate {
	// Group entries by date string (YYYY-MM-DD)
	dateGroups := make(map[string][]dca.DCAEntry)
	for _, entry := range entries {
		dateStr := entry.Date.Format("2006-01-02")
		dateGroups[dateStr] = append(dateGroups[dateStr], entry)
	}

	// Convert to slice and calculate daily metrics
	var result []EntryByDate
	for dateStr, dayEntries := range dateGroups {
		entry := calculateDayMetrics(dateStr, dayEntries)
		result = append(result, entry)
	}

	// Sort by date descending (newest first)
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i].Date < result[j].Date {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	// Calculate cumulative totals and cumulative weighted average using PRD formula
	// cumulativeWeightedAvg = sum(all price_per_share × amount up to this day) / sum(all amounts up to this day)
	var cumulativeTotalAmount float64
	var cumulativeSumPriceAmount float64
	for i := range result {
		cumulativeTotalAmount += result[i].TotalInvested
		// Calculate sum of (price_per_share × amount) for this day
		// weightedAvgPrice = sumPriceAmount / totalInvested, so sumPriceAmount = weightedAvgPrice × totalInvested
		cumulativeSumPriceAmount += result[i].WeightedAvgPrice * result[i].TotalInvested
		result[i].TotalInvested = RoundTo8Decimals(cumulativeTotalAmount)
		// Recalculate cumulative weighted average with PRD formula
		if cumulativeTotalAmount > 0 {
			result[i].WeightedAvgPrice = RoundTo8Decimals(cumulativeSumPriceAmount / cumulativeTotalAmount)
		}
	}

	return result
}

// calculateDayMetrics calculates the daily aggregation metrics for a group of entries.
// Returns per-day metrics including the weighted average price for that specific day.
// Note: Cumulative weighted average across all days is calculated in AggregateByDate.
func calculateDayMetrics(dateStr string, entries []dca.DCAEntry) EntryByDate {
	var totalInvested float64
	var sumPriceAmount float64
	var totalShares float64

	for _, entry := range entries {
		totalInvested += entry.Amount
		sumPriceAmount += entry.PricePerShare * entry.Amount
		totalShares += entry.Shares
	}

	// Weighted average price (PRD formula): sum(price_per_share × amount) / sum(amounts)
	var weightedAvgPrice float64
	if totalInvested > 0 {
		weightedAvgPrice = RoundTo8Decimals(sumPriceAmount / totalInvested)
	}

	return EntryByDate{
		Date:             dateStr,
		TotalInvested:    RoundTo8Decimals(totalInvested),
		TotalShares:      RoundTo8Decimals(totalShares),
		WeightedAvgPrice: weightedAvgPrice,
		EntryCount:       len(entries),
	}
}
