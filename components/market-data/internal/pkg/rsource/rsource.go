package rsource

// RatesSource is the intrface that each of the sources must implement
type RatesSource interface {
	SetupQueue(string) error
	TeardownQueue() error
	GetExchangeRates() ([]byte, error)
	StoreExchangeRates([]byte) error
}
