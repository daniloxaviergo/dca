package form

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/danilo/scripts/github/dca/internal/dca"
)

// FormStep represents the current step in the form
type FormStep int

const (
	StepAmount FormStep = iota
	StepDate
	StepAsset
	StepPrice
	StepShares
	StepConfirm
	StepDone
)

// FormField represents a single input field in the form
type FormField struct {
	Prompt   string
	Value    string
	Error    error
	ReadOnly bool
}

// FormModel manages the state of the form
type FormModel struct {
	Step         FormStep
	Fields       map[string]*FormField
	CurrentField int
	Entries      *dca.DCAData
	FilePath     string
	Submitted    bool
}

// Init starts the form with the first field focused
func (m *FormModel) Init() tea.Cmd {
	return nil
}

// Update handles user input and form state transitions
func (m *FormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			// Cancel without saving
			return m, tea.Quit
		case tea.KeyEnter:
			return m.handleEnter()
		case tea.KeyBackspace:
			return m.handleBackspace()
		case tea.KeyRunes:
			return m.handleInput(msg.String())
		case tea.KeyRight, tea.KeyTab:
			return m.handleTabForward()
		case tea.KeyLeft:
			return m.handleTabBackward()
		case tea.KeyUp, tea.KeyDown:
			return m, nil
		}
	case tea.QuitMsg:
		return m, tea.Quit
	}
	return m, nil
}

// handleEnter processes the Enter key press
func (m *FormModel) handleEnter() (tea.Model, tea.Cmd) {
	currentFieldKey := m.getCurrentFieldKey()
	field := m.Fields[currentFieldKey]

	switch m.Step {
	case StepAmount:
		if err := m.validateAmount(field.Value); err != nil {
			field.Error = err
			return m, nil
		}
		field.Error = nil
		m.Step = StepDate
		m.CurrentField = 1
	case StepDate:
		if err := m.validateDate(field.Value); err != nil {
			field.Error = err
			return m, nil
		}
		field.Error = nil
		m.Step = StepAsset
		m.CurrentField = 2
	case StepAsset:
		if err := m.validateAsset(field.Value); err != nil {
			field.Error = err
			return m, nil
		}
		field.Error = nil
		m.Step = StepPrice
		m.CurrentField = 3
	case StepPrice:
		if err := m.validatePrice(field.Value); err != nil {
			field.Error = err
			return m, nil
		}
		field.Error = nil
		// Calculate and update shares
		amount := m.getFieldFloat64("amount")
		price := m.getFieldFloat64("price")
		shares := CalculateSharesFromValues(amount, price)
		m.Fields["shares"].Value = fmt.Sprintf("%.8f", shares)
		m.Fields["shares"].ReadOnly = true
		m.Step = StepConfirm
		m.CurrentField = 5
	case StepConfirm:
		// Submit the form
		if err := m.saveEntry(); err != nil {
			return m, nil
		}
		m.Submitted = true
		m.Step = StepDone
		return m, func() tea.Msg {
			return formSubmittedMsg{}
		}
	}

	return m, nil
}

// handleBackspace processes the Backspace key press
func (m *FormModel) handleBackspace() (tea.Model, tea.Cmd) {
	currentFieldKey := m.getCurrentFieldKey()
	field := m.Fields[currentFieldKey]

	if !field.ReadOnly && len(field.Value) > 0 {
		// Remove last rune
		field.Value = strings.TrimSuffix(field.Value, string(field.Value[len(field.Value)-1]))
	}
	return m, nil
}

// handleInput processes character input
func (m *FormModel) handleInput(runes string) (tea.Model, tea.Cmd) {
	currentFieldKey := m.getCurrentFieldKey()
	field := m.Fields[currentFieldKey]

	if !field.ReadOnly {
		field.Value += runes
	}
	return m, nil
}

// handleTabForward moves focus to the next field
func (m *FormModel) handleTabForward() (tea.Model, tea.Cmd) {
	maxField := 5 // amount, date, asset, price, shares, confirm
	if m.CurrentField < maxField {
		m.CurrentField++
	}
	return m, nil
}

// handleTabBackward moves focus to the previous field
func (m *FormModel) handleTabBackward() (tea.Model, tea.Cmd) {
	if m.CurrentField > 0 {
		m.CurrentField--
	}
	return m, nil
}

// getCurrentFieldKey returns the key for the current field based on step
func (m *FormModel) getCurrentFieldKey() string {
	keys := []string{"amount", "date", "asset", "price", "shares", "confirm"}
	if m.CurrentField >= 0 && m.CurrentField < len(keys) {
		return keys[m.CurrentField]
	}
	return "amount"
}

// getFieldFloat64 safely parses a field value as float64
func (m *FormModel) getFieldFloat64(key string) float64 {
	val, _ := strconv.ParseFloat(m.Fields[key].Value, 64)
	return val
}

// saveEntry saves the current entry to the data file
func (m *FormModel) saveEntry() error {
	amount := m.getFieldFloat64("amount")
	dateStr := m.Fields["date"].Value
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return err
	}
	asset := m.Fields["asset"].Value
	price := m.getFieldFloat64("price")

	// Create entry with calculated shares
	shares := CalculateSharesFromValues(amount, price)

	entry := dca.DCAEntry{
		Amount:        amount,
		Date:          date,
		Asset:         asset,
		PricePerShare: price,
		Shares:        shares,
	}

	// Validate the entry
	if err := entry.Validate(); err != nil {
		return err
	}

	// Initialize entries map if nil
	if m.Entries.Entries == nil {
		m.Entries.Entries = make(map[string][]dca.DCAEntry)
	}

	// Add entry to the asset's list
	m.Entries.Entries[asset] = append(m.Entries.Entries[asset], entry)

	// Save to file
	if err := dca.SaveEntries(m.FilePath, m.Entries); err != nil {
		return err
	}

	return nil
}

// CalculateSharesFromValues calculates shares from float64 values
func CalculateSharesFromValues(amount, price float64) float64 {
	if price == 0 {
		return 0
	}
	shares := amount / price
	// Validate shares is a finite number
	if math.IsNaN(shares) || math.IsInf(shares, 0) {
		return 0
	}
	return RoundTo8Decimals(shares)
}

// RoundTo8Decimals rounds a float to 8 decimal places
func RoundTo8Decimals(val float64) float64 {
	return float64(int(val*1e8+.5)) / 1e8
}

// View renders the form UI
func (m *FormModel) View() string {
	var sb strings.Builder

	// Header
	sb.WriteString(m.renderHeader())

	// Form fields
	sb.WriteString("\n\n")
	sb.WriteString(m.renderForm())

	// Footer
	sb.WriteString("\n\n")
	sb.WriteString(m.renderFooter())

	return sb.String()
}

// renderHeader displays the form title
func (m *FormModel) renderHeader() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("159")).
		Background(lipgloss.Color("236")).
		Bold(true).
		Underline(true).
		Render("Enter DCA Entry")

	subtitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render("Fill in the form below to add a new DCA investment entry")

	return lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		subtitle,
	)
}

// renderForm displays all form fields
func (m *FormModel) renderForm() string {
	var fields []string

	fieldConfigs := []struct {
		key       string
		prompt    string
		step      FormStep
		readOnly  bool
		isConfirm bool
	}{
		{"amount", "Amount (USD)", StepAmount, false, false},
		{"date", "Date (RFC3339)", StepDate, false, false},
		{"asset", "Asset Ticker", StepAsset, false, false},
		{"price", "Price per Share", StepPrice, false, false},
		{"shares", "Calculated Shares", StepShares, true, false},
		{"confirm", "Confirm (y/n)", StepConfirm, false, true},
	}

	for i, config := range fieldConfigs {
		field := m.Fields[config.key]

		// Only show fields for current or previous steps
		if config.step > m.Step && !m.Submitted {
			continue
		}

		fieldStyle := lipgloss.NewStyle().PaddingLeft(2)

		// Style the field label
		labelStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Bold(true)

		prompt := config.prompt
		if config.isConfirm {
			prompt = "Confirm submission"
		}

		// Check if this is the active field
		if i == m.CurrentField && !m.Submitted {
			labelStyle = labelStyle.
				Foreground(lipgloss.Color("63")).
				Bold(true)
			fieldStyle = fieldStyle.Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(0, 1)
		}

		// Display value with validation error if present
		valueDisplay := field.Value
		if field.Error != nil && !m.Submitted {
			valueDisplay = fmt.Sprintf("❌ %s", field.Error.Error())
			fieldStyle = fieldStyle.Foreground(lipgloss.Color("196"))
		} else if config.readOnly && !m.Submitted {
			// Calculate and display shares dynamically
			if config.key == "shares" && m.Step >= StepPrice {
				amount := m.getFieldFloat64("amount")
				price := m.getFieldFloat64("price")
				shares := CalculateSharesFromValues(amount, price)
				valueDisplay = fmt.Sprintf("%.8f", shares)
			}
		}

		fieldLine := fmt.Sprintf("%s: %s", labelStyle.Render(prompt), valueDisplay)

		// Add helper text for confirm step
		if config.isConfirm {
			fieldLine = fmt.Sprintf("%s: [y]es / [n]o", labelStyle.Render("Confirm submission"))
			if m.Fields["confirm"].Value == "y" {
				fieldLine += lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Render(" ✓")
			} else if m.Fields["confirm"].Value == "n" {
				fieldLine += lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(" ✗")
			}
		}

		fields = append(fields, fieldLine)
	}

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1).
		Render(lipgloss.JoinVertical(lipgloss.Left, fields...))
}

// renderFooter displays navigation instructions
func (m *FormModel) renderFooter() string {
	if m.Submitted {
		success := lipgloss.NewStyle().
			Foreground(lipgloss.Color("82")).
			Bold(true).
			Render("✓ Entry saved successfully!")

		cancel := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Render("\nPress Ctrl+C to exit")

		return lipgloss.JoinVertical(lipgloss.Center, success, cancel)
	}

	// Navigation help text
	help := []string{}
	if m.Step >= StepAmount {
		help = append(help, "[Enter] Next / Submit")
	}
	if m.CurrentField > 0 && m.Step < StepConfirm {
		help = append(help, "[←] Previous")
	}
	if m.CurrentField < 5 && m.Step < StepConfirm {
		help = append(help, "[→] Next field")
	}
	help = append(help, "[Ctrl+C] Cancel")

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render(strings.Join(help, "  "))
}

// NewFormModel creates a new form model with default values
func NewFormModel(entries *dca.DCAData, filePath string) *FormModel {
	now := time.Now()
	defaultDate := now.Format(time.RFC3339)

	return &FormModel{
		Step:         StepAmount,
		CurrentField: 0,
		FilePath:     filePath,
		Entries:      entries,
		Fields: map[string]*FormField{
			"amount":  {Prompt: "Amount (USD)", Value: "", Error: nil},
			"date":    {Prompt: "Date", Value: defaultDate, Error: nil},
			"asset":   {Prompt: "Asset Ticker", Value: "", Error: nil},
			"price":   {Prompt: "Price per Share", Value: "", Error: nil},
			"shares":  {Prompt: "Calculated Shares", Value: "", Error: nil, ReadOnly: true},
			"confirm": {Prompt: "Confirm", Value: "", Error: nil},
		},
	}
}

// formSubmittedMsg is sent when the form is successfully submitted
type formSubmittedMsg struct{}
