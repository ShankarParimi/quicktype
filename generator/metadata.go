// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    metadata, err := UnmarshalMetadata(bytes)
//    bytes, err = metadata.Marshal()

package metadata

import "encoding/json"

type Metadata map[string]Metadatum

func UnmarshalMetadata(data []byte) (Metadata, error) {
	var r Metadata
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Metadata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Metadatum struct {
	Alpha3            string              `json:"alpha_3"`
	ContinentCode     string              `json:"continent_code"`
	ContinentName     string              `json:"continent_name"`
	CountryName       string              `json:"country_name"`
	Currency          []string            `json:"currency"`
	DefaultLocale     string              `json:"default_locale"`
	DialCode          string              `json:"dial_code"`
	Flag              string              `json:"flag"`
	Locales           []Locale            `json:"locales"`
	NumericCode       string              `json:"numeric_code"`
	Sovereignty       string              `json:"sovereignty"`
	TimezoneOfCapital string              `json:"timezone_of_capital"`
	Timezones         map[string]Timezone `json:"timezones"`
}

func (r *Metadatum) GetAlpha3 () string  {
	return r.Alpha3 
}

func (r *Metadatum) GetContinentCode () string  {
	return r.ContinentCode 
}

func (r *Metadatum) GetContinentName () string  {
	return r.ContinentName 
}

func (r *Metadatum) GetCountryName () string  {
	return r.CountryName 
}

func (r *Metadatum) GetCurrency () []string  {
	return r.Currency 
}

func (r *Metadatum) GetDefaultLocale () string  {
	return r.DefaultLocale 
}

func (r *Metadatum) GetDialCode () string  {
	return r.DialCode 
}

func (r *Metadatum) GetFlag () string  {
	return r.Flag 
}

func (r *Metadatum) GetLocales () []Locale  {
	return r.Locales 
}

func (r *Metadatum) GetNumericCode () string  {
	return r.NumericCode 
}

func (r *Metadatum) GetSovereignty () string  {
	return r.Sovereignty 
}

func (r *Metadatum) GetTimezoneOfCapital () string  {
	return r.TimezoneOfCapital 
}

func (r *Metadatum) GetTimezones () map[string]Timezone  {
	return r.Timezones 
}

func NewMetadatum(alpha_3 string, continentCode string, continentName string, countryName string, currency [],string, defaultLocale string, dialCode string, flag string, locales []Locales, numericCode string, sovereignty string, timezoneOfCapital string, timezones map[string]Timezones) *Metadatum {
	return &Metadatum{
		Alpha3 : alpha_3,
		ContinentCode : continentCode,
		ContinentName : continentName,
		CountryName : countryName,
		Currency : currency,
		DefaultLocale : defaultLocale,
		DialCode : dialCode,
		Flag : flag,
		Locales : locales,
		NumericCode : numericCode,
		Sovereignty : sovereignty,
		TimezoneOfCapital : timezoneOfCapital,
		Timezones : timezones,
	}
}

type Locale struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (r *Locale) GetCode () string  {
	return r.Code 
}

func (r *Locale) GetName () string  {
	return r.Name 
}

func NewLocale(code string, name string) *Locale {
	return &Locale{
		Code : code,
		Name : name,
	}
}

type Timezone struct {
	UTCOffset string `json:"utc_offset"`
}

func (r *Timezone) GetUTCOffset () string  {
	return r.UTCOffset 
}

func NewTimezone(utcOffset string) *Timezone {
	return &Timezone{
		UTCOffset : utcOffset,
	}
}
