// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    country, err := UnmarshalCountry(bytes)
//    bytes, err = country.Marshal()

package country

import "encoding/json"

func UnmarshalCountry(data []byte) (Country, error) {
	var r Country
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Country) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Country struct {
	CountryInformation CountryInformation `json:"country_information"`
}

func (r *Country) GetCountryInformation () CountryInformation  {
	return r.CountryInformation 
}

func NewCountry(countryInformation [object Object]) *Country {
	return &Country{
		CountryInformation : countryInformation,
	}
}

type CountryInformation struct {
	My My `json:"MY"`
	In In `json:"IN"`
}

func (r *CountryInformation) GetMy () My  {
	return r.My 
}

func (r *CountryInformation) GetIn () In  {
	return r.In 
}

func NewCountryInformation(my [object Object], in [object Object]) *CountryInformation {
	return &CountryInformation{
		My : my,
		In : in,
	}
}

type In struct {
	CountryName   string          `json:"Country_name"`
	ContinentName string          `json:"continent_name"`
	ContinentCode string          `json:"continent_code"`
	Subdivisions1 INSubdivisions1 `json:"subdivisions_1"`
}

func (r *In) GetCountryName () string  {
	return r.CountryName 
}

func (r *In) GetContinentName () string  {
	return r.ContinentName 
}

func (r *In) GetContinentCode () string  {
	return r.ContinentCode 
}

func (r *In) GetSubdivisions1 () INSubdivisions1  {
	return r.Subdivisions1 
}

func NewIn(countryName string, continentName string, continentCode string, subdivisions_1 [object Object]) *In {
	return &In{
		CountryName : countryName,
		ContinentName : continentName,
		ContinentCode : continentCode,
		Subdivisions1 : subdivisions_1,
	}
}

type INSubdivisions1 struct {
	Ka Ka `json:"KA"`
	Tn Ka `json:"TN"`
	Mh Ka `json:"MH"`
}

func (r *INSubdivisions1) GetKa () Ka  {
	return r.Ka 
}

func (r *INSubdivisions1) GetTn () Ka  {
	return r.Tn 
}

func (r *INSubdivisions1) GetMh () Ka  {
	return r.Mh 
}

func NewINSubdivisions1(ka [object Object], tn [object Object], mh [object Object]) *INSubdivisions1 {
	return &INSubdivisions1{
		Ka : ka,
		Tn : tn,
		Mh : mh,
	}
}

type Ka struct {
	SubDivision1_Name string  `json:"sub_division_1_name"`
	Pincode           []int64 `json:"pincode"`
}

func (r *Ka) GetSubDivision1_Name () string  {
	return r.SubDivision1_Name 
}

func (r *Ka) GetPincode () []int64  {
	return r.Pincode 
}

func NewKa(subDivision_1Name string, pincode [],int64) *Ka {
	return &Ka{
		SubDivision1_Name : subDivision_1Name,
		Pincode : pincode,
	}
}

type My struct {
	CountryName   string          `json:"Country_name"`
	ContinentName string          `json:"continent_name"`
	ContinentCode string          `json:"continent_code"`
	Subdivisions1 MYSubdivisions1 `json:"subdivisions_1"`
}

func (r *My) GetCountryName () string  {
	return r.CountryName 
}

func (r *My) GetContinentName () string  {
	return r.ContinentName 
}

func (r *My) GetContinentCode () string  {
	return r.ContinentCode 
}

func (r *My) GetSubdivisions1 () MYSubdivisions1  {
	return r.Subdivisions1 
}

func NewMy(countryName string, continentName string, continentCode string, subdivisions_1 [object Object]) *My {
	return &My{
		CountryName : countryName,
		ContinentName : continentName,
		ContinentCode : continentCode,
		Subdivisions1 : subdivisions_1,
	}
}

type MYSubdivisions1 struct {
	Kedah  Ka `json:"Kedah"`
	Sdfsdf Ka `json:"sdfsdf"`
	Aaaaa  Ka `json:"aaaaa"`
}

func (r *MYSubdivisions1) GetKedah () Ka  {
	return r.Kedah 
}

func (r *MYSubdivisions1) GetSdfsdf () Ka  {
	return r.Sdfsdf 
}

func (r *MYSubdivisions1) GetAaaaa () Ka  {
	return r.Aaaaa 
}

func NewMYSubdivisions1(kedah [object Object], sdfsdf [object Object], aaaaa [object Object]) *MYSubdivisions1 {
	return &MYSubdivisions1{
		Kedah : kedah,
		Sdfsdf : sdfsdf,
		Aaaaa : aaaaa,
	}
}
