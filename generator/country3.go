// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    country3, err := UnmarshalCountry3(bytes)
//    bytes, err = country3.Marshal()

package country

import "encoding/json"

func UnmarshalCountry3(data []byte) (Country3, error) {
	var r Country3
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Country3) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Country3 struct {
	CountryInformation CountryInformation `json:"country_information"`
}

func (r *Country3) GetCountryInformation () CountryInformation  {
	return r.CountryInformation 
}

func NewCountry3(countryInformation [object Object]) *Country3 {
	return &Country3{
		CountryInformation : countryInformation,
	}
}

type CountryInformation struct {
	ContinentCode string        `json:"continent_code"`
	ContinentName string        `json:"continent_name"`
	CountryName   string        `json:"country_name"`
	Subdivisions1 Subdivisions1 `json:"subdivisions_1"`
}

func (r *CountryInformation) GetContinentCode () string  {
	return r.ContinentCode 
}

func (r *CountryInformation) GetContinentName () string  {
	return r.ContinentName 
}

func (r *CountryInformation) GetCountryName () string  {
	return r.CountryName 
}

func (r *CountryInformation) GetSubdivisions1 () Subdivisions1  {
	return r.Subdivisions1 
}

func NewCountryInformation(continentCode string, continentName string, countryName string, subdivisions_1 [object Object]) *CountryInformation {
	return &CountryInformation{
		ContinentCode : continentCode,
		ContinentName : continentName,
		CountryName : countryName,
		Subdivisions1 : subdivisions_1,
	}
}

type Subdivisions1 struct {
	Pincode           []int64 `json:"pincode"`
	SubDivision1_Name string  `json:"sub_division_1_name"`
}

func (r *Subdivisions1) GetPincode () []int64  {
	return r.Pincode 
}

func (r *Subdivisions1) GetSubDivision1_Name () string  {
	return r.SubDivision1_Name 
}

func NewSubdivisions1(pincode [],int64, subDivision_1Name string) *Subdivisions1 {
	return &Subdivisions1{
		Pincode : pincode,
		SubDivision1_Name : subDivision_1Name,
	}
}
