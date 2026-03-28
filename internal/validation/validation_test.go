package validation

import (
	"math"
	"testing"
)

func TestIsValidAmount_Pass(t *testing.T) {
	err := IsValidAmount("100.50")
	if err != nil {
		t.Errorf("Expected valid amount to pass, got error: %v", err)
	}
}

func TestIsValidAmount_RejectZero(t *testing.T) {
	err := IsValidAmount("0")
	if err == nil {
		t.Error("Expected zero amount to fail validation")
	}
}

func TestIsValidAmount_RejectNegative(t *testing.T) {
	err := IsValidAmount("-50")
	if err == nil {
		t.Error("Expected negative amount to fail validation")
	}
}

func TestIsValidAmount_RejectEmpty(t *testing.T) {
	err := IsValidAmount("")
	if err == nil {
		t.Error("Expected empty amount to fail validation")
	}
}

func TestIsValidAmount_RejectInvalid(t *testing.T) {
	err := IsValidAmount("abc")
	if err == nil {
		t.Error("Expected invalid amount to fail validation")
	}
}

func TestIsValidAmount_ExactErrorMessage(t *testing.T) {
	err := IsValidAmount("0")
	if err == nil {
		t.Error("Expected zero amount to fail validation")
	}
	if err.Error() != "Amount must be positive" {
		t.Errorf("Expected exact error message 'Amount must be positive', got: %v", err)
	}
}

func TestIsValidAmount_NegativeExactErrorMessage(t *testing.T) {
	err := IsValidAmount("-100")
	if err == nil {
		t.Error("Expected negative amount to fail validation")
	}
	if err.Error() != "Amount must be positive" {
		t.Errorf("Expected exact error message 'Amount must be positive', got: %v", err)
	}
}

func TestIsValidAmount_Whitespace(t *testing.T) {
	err := IsValidAmount("   ")
	if err == nil {
		t.Error("Expected whitespace-only amount to fail validation")
	}
}

func TestIsValidPrice_Pass(t *testing.T) {
	err := IsValidPrice("50000.00")
	if err != nil {
		t.Errorf("Expected valid price to pass, got error: %v", err)
	}
}

func TestIsValidPrice_RejectZero(t *testing.T) {
	err := IsValidPrice("0")
	if err == nil {
		t.Error("Expected zero price to fail validation")
	}
}

func TestIsValidPrice_RejectNegative(t *testing.T) {
	err := IsValidPrice("-100")
	if err == nil {
		t.Error("Expected negative price to fail validation")
	}
}

func TestIsValidPrice_RejectEmpty(t *testing.T) {
	err := IsValidPrice("")
	if err == nil {
		t.Error("Expected empty price to fail validation")
	}
}

func TestIsValidPrice_RejectInvalid(t *testing.T) {
	err := IsValidPrice("abc")
	if err == nil {
		t.Error("Expected invalid price to fail validation")
	}
}

func TestIsValidPrice_ExactErrorMessage(t *testing.T) {
	err := IsValidPrice("0")
	if err == nil {
		t.Error("Expected zero price to fail validation")
	}
	if err.Error() != "Price must be positive" {
		t.Errorf("Expected exact error message 'Price must be positive', got: %v", err)
	}
}

func TestIsValidPrice_NegativeExactErrorMessage(t *testing.T) {
	err := IsValidPrice("-50")
	if err == nil {
		t.Error("Expected negative price to fail validation")
	}
	if err.Error() != "Price must be positive" {
		t.Errorf("Expected exact error message 'Price must be positive', got: %v", err)
	}
}

func TestIsValidPrice_Whitespace(t *testing.T) {
	err := IsValidPrice("   ")
	if err == nil {
		t.Error("Expected whitespace-only price to fail validation")
	}
}

func TestIsValidAsset_Pass(t *testing.T) {
	err := IsValidAsset("BTC")
	if err != nil {
		t.Errorf("Expected valid asset to pass, got error: %v", err)
	}
}

func TestIsValidAsset_RejectEmpty(t *testing.T) {
	err := IsValidAsset("")
	if err == nil {
		t.Error("Expected empty asset to fail validation")
	}
}

func TestIsValidAsset_RejectWhitespace(t *testing.T) {
	err := IsValidAsset("   ")
	if err == nil {
		t.Error("Expected whitespace-only asset to fail validation")
	}
}

func TestIsValidAsset_ExactErrorMessage(t *testing.T) {
	err := IsValidAsset("")
	if err == nil {
		t.Error("Expected empty asset to fail validation")
	}
	if err.Error() != "Asset ticker is required" {
		t.Errorf("Expected exact error message 'Asset ticker is required', got: %v", err)
	}
}

func TestIsValidAsset_WithSpaces(t *testing.T) {
	err := IsValidAsset(" BTC ")
	if err != nil {
		t.Errorf("Expected asset with surrounding spaces to pass, got error: %v", err)
	}
}

func TestIsValidDate_Pass(t *testing.T) {
	err := IsValidDate("2024-01-15T10:30:00Z")
	if err != nil {
		t.Errorf("Expected valid date to pass, got error: %v", err)
	}
}

func TestIsValidDate_RejectInvalid(t *testing.T) {
	err := IsValidDate("01-15-2024")
	if err == nil {
		t.Error("Expected invalid date format to fail validation")
	}
}

func TestIsValidDate_RejectEmpty(t *testing.T) {
	err := IsValidDate("")
	if err == nil {
		t.Error("Expected empty date to fail validation")
	}
}

func TestIsValidDate_ExactErrorMessage(t *testing.T) {
	err := IsValidDate("invalid-date")
	if err == nil {
		t.Error("Expected invalid date to fail validation")
	}
	if err.Error() != "Use YYYY-MM-DD" {
		t.Errorf("Expected exact error message 'Use YYYY-MM-DD', got: %v", err)
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
		{val: 0.999999999, want: 1.0}, // rounding up
	}

	for _, tt := range tests {
		got := RoundTo8Decimals(tt.val)
		if got != tt.want {
			t.Errorf("RoundTo8Decimals(%f) = %f, want %f", tt.val, got, tt.want)
		}
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
		{amount: 150.50, price: 75.25, want: 2.0},    // precise calculation
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

func TestCalculateSharesFromValues_NaN(t *testing.T) {
	// Test that NaN is handled gracefully
	amount := 100.0
	price := 0.0
	shares := CalculateSharesFromValues(amount, price)

	if shares != 0 {
		t.Errorf("Expected 0 for NaN result, got: %f", shares)
	}
}

func TestCalculateSharesFromValues_Inf(t *testing.T) {
	// Test that infinity is handled gracefully
	amount := math.Inf(1) // positive infinity
	price := 1.0
	shares := CalculateSharesFromValues(amount, price)

	if shares != 0 {
		t.Errorf("Expected 0 for infinity result, got: %f", shares)
	}
}
