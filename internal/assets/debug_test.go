package assets

import (
	"fmt"
	"testing"
)

func TestLoadMoreDebug2(t *testing.T) {
	m := NewAssetHistoryModal()
	m.AssetTicker = "BTC"

	// Create temp file with 25 entries
	testFile := createTestFileWith25Entries(t)
	defer removeTestFile(t, testFile)

	// Load initial data - should get 10 entries
	err := m.LoadData(testFile, "BTC")
	if err != nil {
		t.Fatalf("LoadData error: %v", err)
	}
	fmt.Printf("After LoadData: %d entries, AllLoaded=%v, Loading=%v\n",
		len(m.EntriesByDate), m.AllLoaded, m.Loading)

	// Reset state
	m.AllLoaded = false
	m.Loading = false

	// LoadMore - should get 10 more
	err = m.LoadMore(testFile)
	if err != nil {
		t.Fatalf("LoadMore error: %v", err)
	}
	fmt.Printf("After LoadMore: %d entries, AllLoaded=%v, Loading=%v\n",
		len(m.EntriesByDate), m.AllLoaded, m.Loading)
}
