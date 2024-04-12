// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    countrySubdivisions, err := UnmarshalCountrySubdivisions(bytes)
//    bytes, err = countrySubdivisions.Marshal()

package country_subdivisions

import "encoding/json"

func UnmarshalCountrySubdivisions(data []byte) (CountrySubdivisions, error) {
	var r CountrySubdivisions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CountrySubdivisions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CountrySubdivisions struct {
	CountryName string           `json:"country_name"`
	States      map[string]State `json:"states"`
}

func (r *CountrySubdivisions) GetCountryName () string  {
	return r.CountryName 
}

func (r *CountrySubdivisions) GetStates () map[string]State  {
	return r.States 
}

func NewCountrySubdivisions(countryName string, states map[string]States) *CountrySubdivisions {
	return &CountrySubdivisions{
		CountryName : countryName,
		States : states,
	}
}

type State struct {
	Cities []City `json:"cities"`
	Name   string `json:"name"`
}

func (r *State) GetCities () []City  {
	return r.Cities 
}

func (r *State) GetName () string  {
	return r.Name 
}

func NewState(cities []Cities, name string) *State {
	return &State{
		Cities : cities,
		Name : name,
	}
}

type City struct {
	Name                   string   `json:"name"`
	RegionNameDistrictName string   `json:"region_name/district_name"`
	Timezone               string   `json:"timezone"`
	Zipcodes               []string `json:"zipcodes"`
}

func (r *City) GetName () string  {
	return r.Name 
}

func (r *City) GetRegionNameDistrictName () string  {
	return r.RegionNameDistrictName 
}

func (r *City) GetTimezone () string  {
	return r.Timezone 
}

func (r *City) GetZipcodes () []string  {
	return r.Zipcodes 
}

func NewCity(name string, regionName/districtName string, timezone string, zipcodes [],string) *City {
	return &City{
		Name : name,
		RegionNameDistrictName : regionName/districtName,
		Timezone : timezone,
		Zipcodes : zipcodes,
	}
}
