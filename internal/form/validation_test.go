package form

import (
	"fmt"
	"strings"
	"testing"

	"github.com/danilo/scripts/github/dca/internal/dca"
)

func TestFormModel_ValidateAmount_Pass(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["amount"].Value = "100.50"

	err := form.validateAmount(form.Fields["amount"].Value)
	if err != nil {
		t.Errorf("Expected valid amount to pass, got error: %v", err)
	}
}

func TestFormModel_ValidateAmount_RejectZero(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["amount"].Value = "0"

	err := form.validateAmount(form.Fields["amount"].Value)
	if err == nil {
		t.Error("Expected zero amount to fail validation")
	}
	if err.Error() != "Amount must be positive" {
		t.Errorf("Expected exact error message 'Amount must be positive', got: %v", err)
	}
}

func TestFormModel_ValidateAmount_RejectNegative(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["amount"].Value = "-50"

	err := form.validateAmount(form.Fields["amount"].Value)
	if err == nil {
		t.Error("Expected negative amount to fail validation")
	}
}

func TestFormModel_ValidateAmount_RejectEmpty(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["amount"].Value = ""

	err := form.validateAmount(form.Fields["amount"].Value)
	if err == nil {
		t.Error("Expected empty amount to fail validation")
	}
}

func TestFormModel_ValidateAmount_RejectInvalid(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["amount"].Value = "abc"

	err := form.validateAmount(form.Fields["amount"].Value)
	if err == nil {
		t.Error("Expected invalid amount to fail validation")
	}
}

func TestFormModel_ValidateDate_Pass(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["date"].Value = "2024-01-15T10:30:00Z"

	err := form.validateDate(form.Fields["date"].Value)
	if err != nil {
		t.Errorf("Expected valid date to pass, got error: %v", err)
	}
}

func TestFormModel_ValidateDate_RejectInvalid(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["date"].Value = "01-15-2024"

	err := form.validateDate(form.Fields["date"].Value)
	if err == nil {
		t.Error("Expected invalid date format to fail validation")
	}
}

func TestFormModel_ValidateAsset_Pass(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["asset"].Value = "BTC"

	err := form.validateAsset(form.Fields["asset"].Value)
	if err != nil {
		t.Errorf("Expected valid asset to pass, got error: %v", err)
	}
}

func TestFormModel_ValidateAsset_RejectEmpty(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["asset"].Value = ""

	err := form.validateAsset(form.Fields["asset"].Value)
	if err == nil {
		t.Error("Expected empty asset to fail validation")
	}
}

func TestFormModel_ValidateAsset_RejectWhitespace(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["asset"].Value = "   "

	err := form.validateAsset(form.Fields["asset"].Value)
	if err == nil {
		t.Error("Expected whitespace-only asset to fail validation")
	}
}

func TestFormModel_ValidatePrice_Pass(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["price"].Value = "50000.00"

	err := form.validatePrice(form.Fields["price"].Value)
	if err != nil {
		t.Errorf("Expected valid price to pass, got error: %v", err)
	}
}

func TestFormModel_ValidatePrice_RejectZero(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["price"].Value = "0"

	err := form.validatePrice(form.Fields["price"].Value)
	if err == nil {
		t.Error("Expected zero price to fail validation")
	}
}

func TestFormModel_ValidatePrice_RejectNegative(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["price"].Value = "-100"

	err := form.validatePrice(form.Fields["price"].Value)
	if err == nil {
		t.Error("Expected negative price to fail validation")
	}
}

func TestCalculateSharesFromValues(t *testing.T) {
	tests := []struct {
		amount float64
		price  float64
		want   float64
	}{
		{amount: 100, price: 50, want: 2.0},
		{amount: 100, price: 33.33333333, want: 3.0}, // 8 decimal precision
		{amount: 100, price: 0, want: 0},             // division by zero
	}

	for _, tt := range tests {
		got := CalculateSharesFromValues(tt.amount, tt.price)
		if got != tt.want {
			t.Errorf("CalculateSharesFromValues(%f, %f) = %f, want %f", tt.amount, tt.price, got, tt.want)
		}
	}
}

func TestCalculateSharesFromValues_Precision(t *testing.T) {
	// Test 8 decimal precision
	amount := 100.0
	price := 33.33333333
	shares := CalculateSharesFromValues(amount, price)

	// Should round to 8 decimal places
	if shares <= 0 {
		t.Errorf("Expected positive shares, got: %f", shares)
	}
}

func TestRoundTo8Decimals(t *testing.T) {
	tests := []struct {
		val  float64
		want float64
	}{
		{val: 3.1415926535, want: 3.14159265},
		{val: 2.718281828, want: 2.71828183},
		{val: 1.0, want: 1.0},
		{val: 0.000000001, want: 0.0}, // smaller than 8 decimal precision
	}

	for _, tt := range tests {
		got := RoundTo8Decimals(tt.val)
		if got != tt.want {
			t.Errorf("RoundTo8Decimals(%f) = %f, want %f", tt.val, got, tt.want)
		}
	}
}

func TestFormModel_GetFieldFloat64(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["amount"].Value = "123.45"

	got := form.getFieldFloat64("amount")
	want := 123.45

	if got != want {
		t.Errorf("getFieldFloat64('amount') = %f, want %f", got, want)
	}
}

func TestFormModel_GetFieldFloat64_Empty(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["amount"].Value = ""

	got := form.getFieldFloat64("amount")
	if got != 0 {
		t.Errorf("getFieldFloat64('') should return 0, got %f", got)
	}
}

func TestFormModel_GetCurrentFieldKey(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")

	tests := []struct {
		field int
		want  string
	}{
		{field: 0, want: "amount"},
		{field: 1, want: "date"},
		{field: 2, want: "asset"},
		{field: 3, want: "price"},
		{field: 4, want: "shares"},
		{field: 5, want: "confirm"},
	}

	for _, tt := range tests {
		form.CurrentField = tt.field
		got := form.getCurrentFieldKey()
		if got != tt.want {
			t.Errorf("getCurrentFieldKey(%d) = %s, want %s", tt.field, got, tt.want)
		}
	}
}

func TestFormModel_TabForward(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.CurrentField = 0

	form.handleTabForward()
	if form.CurrentField != 1 {
		t.Errorf("After tab forward, field should be 1, got %d", form.CurrentField)
	}
}

func TestFormModel_TabBackward(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.CurrentField = 1

	form.handleTabBackward()
	if form.CurrentField != 0 {
		t.Errorf("After tab backward, field should be 0, got %d", form.CurrentField)
	}
}

func TestFormModel_HandleBackspace(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["amount"].Value = "123"
	form.CurrentField = 0

	form.handleBackspace()
	if form.Fields["amount"].Value != "12" {
		t.Errorf("After backspace, amount should be '12', got '%s'", form.Fields["amount"].Value)
	}
}

func TestFormModel_HandleInput(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["amount"].Value = ""
	form.CurrentField = 0

	form.handleInput("123")
	if form.Fields["amount"].Value != "123" {
		t.Errorf("After input, amount should be '123', got '%s'", form.Fields["amount"].Value)
	}
}

func TestFormModel_RenderForm(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")

	view := form.View()

	// Check that form renders successfully
	if view == "" {
		t.Error("Expected view to not be empty")
	}

	// Check for key elements
	if !strings.Contains(view, "Enter DCA Entry") {
		t.Error("Expected header 'Enter DCA Entry' in view")
	}
}

// TestFormModel_ValidateAmount_ExactErrorMessage tests exact error message format
func TestFormModel_ValidateAmount_ExactErrorMessage(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["amount"].Value = "0"

	err := form.validateAmount(form.Fields["amount"].Value)
	if err == nil {
		t.Error("Expected zero amount to fail validation")
	}
	if err.Error() != "Amount must be positive" {
		t.Errorf("Expected exact error message 'Amount must be positive', got: %v", err)
	}
}

// TestFormModel_ValidateAmount_NegativeExactErrorMessage tests exact error message for negative
func TestFormModel_ValidateAmount_NegativeExactErrorMessage(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["amount"].Value = "-100"

	err := form.validateAmount(form.Fields["amount"].Value)
	if err == nil {
		t.Error("Expected negative amount to fail validation")
	}
	if err.Error() != "Amount must be positive" {
		t.Errorf("Expected exact error message 'Amount must be positive', got: %v", err)
	}
}

// TestFormModel_ValidateDate_ExactErrorMessage tests exact error message format with YYYY-MM-DD
func TestFormModel_ValidateDate_ExactErrorMessage(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["date"].Value = "invalid-date"

	err := form.validateDate(form.Fields["date"].Value)
	if err == nil {
		t.Error("Expected invalid date to fail validation")
	}
	if err.Error() != "Use YYYY-MM-DD" {
		t.Errorf("Expected exact error message 'Use YYYY-MM-DD', got: %v", err)
	}
}

// TestFormModel_ValidateAsset_ExactErrorMessage tests exact error message format
func TestFormModel_ValidateAsset_ExactErrorMessage(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["asset"].Value = ""

	err := form.validateAsset(form.Fields["asset"].Value)
	if err == nil {
		t.Error("Expected empty asset to fail validation")
	}
	if err.Error() != "Asset ticker is required" {
		t.Errorf("Expected exact error message 'Asset ticker is required', got: %v", err)
	}
}

// TestFormModel_ValidatePrice_ExactErrorMessage tests exact error message format
func TestFormModel_ValidatePrice_ExactErrorMessage(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["price"].Value = "0"

	err := form.validatePrice(form.Fields["price"].Value)
	if err == nil {
		t.Error("Expected zero price to fail validation")
	}
	if err.Error() != "Price must be positive" {
		t.Errorf("Expected exact error message 'Price must be positive', got: %v", err)
	}
}

// TestFormModel_ValidatePrice_NegativeExactErrorMessage tests exact error message for negative
func TestFormModel_ValidatePrice_NegativeExactErrorMessage(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["price"].Value = "-50"

	err := form.validatePrice(form.Fields["price"].Value)
	if err == nil {
		t.Error("Expected negative price to fail validation")
	}
	if err.Error() != "Price must be positive" {
		t.Errorf("Expected exact error message 'Price must be positive', got: %v", err)
	}
}

// TestFormModel_InlineErrorDisplay tests error display with inline error
func TestFormModel_InlineErrorDisplay(t *testing.T) {
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}
	form := NewFormModel(entries, "test.json")
	form.Fields["amount"].Value = "0"
	form.Fields["amount"].Error = fmt.Errorf("Amount must be positive")

	view := form.View()

	// Check that error is displayed in the view
	if !strings.Contains(view, "❌") {
		t.Errorf("Expected ❌ error indicator in view, got: %s", view)
	}
	if !strings.Contains(view, "Amount must be positive") {
		t.Errorf("Expected error message in view, got: %s", view)
	}
}
