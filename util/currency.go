package util

const (
	USD = "USD"
	EUR = "EUR"
	BRL = "BRL"
)

// IsSuppoertedCurrency return true if the currency is supported
func IsSuppoertedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, BRL:
		return true
	}

	return false
}
