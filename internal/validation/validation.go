package validation

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// IsValidAmount validates that the amount is a positive number
func IsValidAmount(value string) error {
	if value == "" {
		return fmt.Errorf("Amount must be positive")
	}
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("Amount must be positive")
	}
	if val <= 0 {
		return fmt.Errorf("Amount must be positive")
	}
	return nil
}

// IsValidPrice validates that the price is a positive number
func IsValidPrice(value string) error {
	if value == "" {
		return fmt.Errorf("Price must be positive")
	}
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("Price must be positive")
	}
	if val <= 0 {
		return fmt.Errorf("Price must be positive")
	}
	return nil
}

// IsValidAsset validates that the asset ticker is non-empty
func IsValidAsset(value string) error {
	if value == "" {
		return fmt.Errorf("Asset ticker is required")
	}
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("Asset ticker is required")
	}
	return nil
}

// IsValidDate validates that the date is in RFC3339 format
func IsValidDate(value string) error {
	if value == "" {
		return fmt.Errorf("Use YYYY-MM-DD")
	}
	_, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return fmt.Errorf("Use YYYY-MM-DD")
	}
	return nil
}

// RoundTo8Decimals rounds a float to 8 decimal places
func RoundTo8Decimals(val float64) float64 {
	return float64(int(val*1e8+.5)) / 1e8
}

// CalculateSharesFromValues calculates shares from float64 values
func CalculateSharesFromValues(amount, price float64) float64 {
	if price == 0 {
		return 0
	}
	shares := amount / price
	// Validate shares is a finite number
	if isNaNOrInf(shares) {
		return 0
	}
	return RoundTo8Decimals(shares)
}

// isNaNOrInf checks if a value is NaN or infinite
func isNaNOrInf(val float64) bool {
	return isNaN(val) || isInf(val)
}

// isNaN checks if a value is NaN (without using math.IsNaN for compatibility)
func isNaN(val float64) bool {
	return val != val
}

// isInf checks if a value is infinite (without using math.IsInf for compatibility)
func isInf(val float64) bool {
	return val == val && val != 0 && (val > 1e308 || val < -1e308)
}
