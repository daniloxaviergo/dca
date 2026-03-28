package form

import (
	"github.com/danilo/scripts/github/dca/internal/validation"
)

// validateAmount validates that the amount is a positive number
func (m *FormModel) validateAmount(value string) error {
	return validation.IsValidAmount(value)
}

// validateDate validates that the date is in RFC3339 format
func (m *FormModel) validateDate(value string) error {
	return validation.IsValidDate(value)
}

// validateAsset validates that the asset ticker is non-empty
func (m *FormModel) validateAsset(value string) error {
	return validation.IsValidAsset(value)
}

// validatePrice validates that the price is a positive number
func (m *FormModel) validatePrice(value string) error {
	return validation.IsValidPrice(value)
}
