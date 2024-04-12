// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    types1, err := UnmarshalTypes1(bytes)
//    bytes, err = types1.Marshal()

package main

import "encoding/json"

func UnmarshalTypes1(data []byte) (Types1, error) {
	var r Types1
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Types1) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Types1 struct {
	CountryInformation map[string]CountryInformation `json:"country_information"`
}
func (r *Types1) GetTypes1() ([]byte, error) {
	return r
}
func (r *CountryInformation ) GetCountryInformation () map[string]CountryInformation  {
	return *r.CountryInformation 
}
func (r *CountryInformation ) SetCountryInformation (data map[string]CountryInformation ) {
	r.CountryInformation  = data
}

type CountryInformation struct {
	CountryName  string `json:"country_name"`
	CountryCode3 string `json:"country_code_3"`
	NumericCode  string `json:"numeric_code"`
	Sovereignty  string `json:"sovereignty"`
}
func (r *CountryInformation) GetCountryInformation() ([]byte, error) {
	return r
}
func (r *CountryName ) GetCountryName () string  {
	return *r.CountryName 
}
func (r *CountryName ) SetCountryName (data string ) {
	r.CountryName  = data
}
func (r *CountryCode3 ) GetCountryCode3 () string  {
	return *r.CountryCode3 
}
func (r *CountryCode3 ) SetCountryCode3 (data string ) {
	r.CountryCode3  = data
}
func (r *NumericCode ) GetNumericCode () string  {
	return *r.NumericCode 
}
func (r *NumericCode ) SetNumericCode (data string ) {
	r.NumericCode  = data
}
func (r *Sovereignty ) GetSovereignty () string  {
	return *r.Sovereignty 
}
func (r *Sovereignty ) SetSovereignty (data string ) {
	r.Sovereignty  = data
}

//// "start": "script/quicktype -l go -o generator/country1.go --generate-i18nify-stubs --package country /Users/debanshu.rout/repo/i18nify/data/country_v1.json",