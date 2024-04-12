// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    currency, err := UnmarshalCurrency(bytes)
//    bytes, err = currency.Marshal()

package currency

import "encoding/json"

func UnmarshalCurrency(data []byte) (Currency, error) {
	var r Currency
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Currency) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Currency struct {
	CurrencyInformation CurrencyInformation `json:"currency_information"`
}

func (r *Currency) GetCurrencyInformation () CurrencyInformation  {
	return r.CurrencyInformation 
}

func NewCurrency(currencyInformation [object Object]) *Currency {
	return &Currency{
		CurrencyInformation : currencyInformation,
	}
}

type CurrencyInformation struct {
}
