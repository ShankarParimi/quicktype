// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    rzp, err := UnmarshalRzp(bytes)
//    bytes, err = rzp.Marshal()

package main

import "bytes"
import "errors"
import "encoding/json"

func UnmarshalRzp(data []byte) (Rzp, error) {
	var r Rzp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Rzp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Rzp struct {
	CountryCode CountryCode `json:"country_code"`
}

func (r *Rzp) GetCountryCode () CountryCode  {
	return r.CountryCode 
}

func NewRzp(countryCode [object Object]) *Rzp {
	return &Rzp{
		CountryCode : countryCode,
	}
}

type CountryCode struct {
	ContinentCode       string        `json:"continent_code"`
	ContinentName       string        `json:"continent_name"`
	SubDivisions1_Codes []string      `json:"sub_divisions_1_codes"`
	SubDivisions2_Codes []interface{} `json:"sub_divisions_2_codes"`
	Cities              []interface{} `json:"cities"`
	SubDivisions1       SubDivisions1 `json:"sub_divisions_1"`
}

func (r *CountryCode) GetContinentCode () string  {
	return r.ContinentCode 
}

func (r *CountryCode) GetContinentName () string  {
	return r.ContinentName 
}

func (r *CountryCode) GetSubDivisions1_Codes () []string  {
	return r.SubDivisions1_Codes 
}

func (r *CountryCode) GetSubDivisions2_Codes () []interface{}  {
	return r.SubDivisions2_Codes 
}

func (r *CountryCode) GetCities () []interface{}  {
	return r.Cities 
}

func (r *CountryCode) GetSubDivisions1 () SubDivisions1  {
	return r.SubDivisions1 
}

func NewCountryCode(continentCode string, continentName string, subDivisions_1Codes [],string, subDivisions_2Codes []SubDivisions2_Codes, cities []Cities, subDivisions_1 [object Object]) *CountryCode {
	return &CountryCode{
		ContinentCode : continentCode,
		ContinentName : continentName,
		SubDivisions1_Codes : subDivisions_1Codes,
		SubDivisions2_Codes : subDivisions_2Codes,
		Cities : cities,
		SubDivisions1 : subDivisions_1,
	}
}

type SubDivisions1 struct {
	SubDivision1_Code SubDivision1_Code `json:"sub_division_1_code"`
}

func (r *SubDivisions1) GetSubDivision1_Code () SubDivision1_Code  {
	return r.SubDivision1_Code 
}

func NewSubDivisions1(subDivision_1Code [object Object]) *SubDivisions1 {
	return &SubDivisions1{
		SubDivision1_Code : subDivision_1Code,
	}
}

type SubDivision1_Code struct {
	SubDivision1_Name string        `json:"sub_division_1_name"`
	SubDivisions2     SubDivisions2 `json:"sub_divisions_2"`
}

func (r *SubDivision1_Code) GetSubDivision1_Name () string  {
	return r.SubDivision1_Name 
}

func (r *SubDivision1_Code) GetSubDivisions2 () SubDivisions2  {
	return r.SubDivisions2 
}

func NewSubDivision1_Code(subDivision_1Name string, subDivisions_2 [object Object]) *SubDivision1_Code {
	return &SubDivision1_Code{
		SubDivision1_Name : subDivision_1Name,
		SubDivisions2 : subDivisions_2,
	}
}

type SubDivisions2 struct {
	SubDivision2_Code SubDivision2_Code `json:"sub_division_2_code"`
}

func (r *SubDivisions2) GetSubDivision2_Code () SubDivision2_Code  {
	return r.SubDivision2_Code 
}

func NewSubDivisions2(subDivision_2Code [object Object]) *SubDivisions2 {
	return &SubDivisions2{
		SubDivision2_Code : subDivision_2Code,
	}
}

type SubDivision2_Code struct {
	SubDivision2_Name string        `json:"sub_division_2_name"`
	Cities            []CityElement `json:"Cities"`
}

func (r *SubDivision2_Code) GetSubDivision2_Name () string  {
	return r.SubDivision2_Name 
}

func (r *SubDivision2_Code) GetCities () []CityElement  {
	return r.Cities 
}

func NewSubDivision2_Code(subDivision_2Name string, cities []Cities) *SubDivision2_Code {
	return &SubDivision2_Code{
		SubDivision2_Name : subDivision_2Name,
		Cities : cities,
	}
}

type CityClass struct {
	CityName string        `json:"city_name"`
	ZipCodes []interface{} `json:"zip_codes"`
}

func (r *CityClass) GetCityName () string  {
	return r.CityName 
}

func (r *CityClass) GetZipCodes () []interface{}  {
	return r.ZipCodes 
}

func NewCityClass(cityName string, zipCodes []ZipCodes) *CityClass {
	return &CityClass{
		CityName : cityName,
		ZipCodes : zipCodes,
	}
}

type CityElement struct {
	CityClass *CityClass
	String    *string
}

func (x *CityElement) UnmarshalJSON(data []byte) error {
	x.CityClass = nil
	var c CityClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.CityClass = &c
	}
	return nil
}

func (x *CityElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.CityClass != nil, x.CityClass, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
