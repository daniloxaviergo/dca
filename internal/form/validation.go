package form

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// validateAmount validates that the amount is a positive number
func (m *FormModel) validateAmount(value string) error {
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

// validateDate validates that the date is in RFC3339 format
func (m *FormModel) validateDate(value string) error {
	if value == "" {
		return fmt.Errorf("Use YYYY-MM-DD")
	}
	_, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return fmt.Errorf("Use YYYY-MM-DD")
	}
	return nil
}

// validateAsset validates that the asset ticker is non-empty
func (m *FormModel) validateAsset(value string) error {
	if value == "" {
		return fmt.Errorf("Asset ticker is required")
	}
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("Asset ticker is required")
	}
	return nil
}

// validatePrice validates that the price is a positive number
func (m *FormModel) validatePrice(value string) error {
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
