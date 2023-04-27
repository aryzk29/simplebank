package util

// constant for currency
const (
	USD = "USD"
	EUR = "EUR"
	IDR = "IDR"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, IDR:
		return true
	}
	return false
}
