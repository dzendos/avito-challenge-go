package types

import "time"

type Currency string

const (
	USD Currency = "USD"
	CNY Currency = "CNY"
	EUR Currency = "EUR"
	RUB Currency = "RUB"
)

type CurrencyRate struct {
	CharCode     string
	BaseCurrency string
	Rate         int64
	Date         time.Time
}

type EncodedCurrencies struct {
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type CurrenciesRate struct {
	EncodedCurrencies []EncodedCurrencies `xml:"Valute"`
}
