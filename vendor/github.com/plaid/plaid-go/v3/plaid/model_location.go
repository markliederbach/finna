/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.128.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// Location A representation of where a transaction took place
type Location struct {
	// The street address where the transaction occurred.
	Address NullableString `json:"address"`
	// The city where the transaction occurred.
	City NullableString `json:"city"`
	// The region or state where the transaction occurred. In API versions 2018-05-22 and earlier, this field is called `state`.
	Region NullableString `json:"region"`
	// The postal code where the transaction occurred. In API versions 2018-05-22 and earlier, this field is called `zip`.
	PostalCode NullableString `json:"postal_code"`
	// The ISO 3166-1 alpha-2 country code where the transaction occurred.
	Country NullableString `json:"country"`
	// The latitude where the transaction occurred.
	Lat NullableFloat64 `json:"lat"`
	// The longitude where the transaction occurred.
	Lon NullableFloat64 `json:"lon"`
	// The merchant defined store number where the transaction occurred.
	StoreNumber NullableString `json:"store_number"`
	AdditionalProperties map[string]interface{}
}

type _Location Location

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation(address NullableString, city NullableString, region NullableString, postalCode NullableString, country NullableString, lat NullableFloat64, lon NullableFloat64, storeNumber NullableString) *Location {
	this := Location{}
	this.Address = address
	this.City = city
	this.Region = region
	this.PostalCode = postalCode
	this.Country = country
	this.Lat = lat
	this.Lon = lon
	this.StoreNumber = storeNumber
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Location) GetAddress() string {
	if o == nil || o.Address.Get() == nil {
		var ret string
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *Location) SetAddress(v string) {
	o.Address.Set(&v)
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Location) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *Location) SetCity(v string) {
	o.City.Set(&v)
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Location) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *Location) SetRegion(v string) {
	o.Region.Set(&v)
}

// GetPostalCode returns the PostalCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Location) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// SetPostalCode sets field value
func (o *Location) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Location) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *Location) SetCountry(v string) {
	o.Country.Set(&v)
}

// GetLat returns the Lat field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Location) GetLat() float64 {
	if o == nil || o.Lat.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Lat.Get()
}

// GetLatOk returns a tuple with the Lat field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetLatOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Lat.Get(), o.Lat.IsSet()
}

// SetLat sets field value
func (o *Location) SetLat(v float64) {
	o.Lat.Set(&v)
}

// GetLon returns the Lon field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Location) GetLon() float64 {
	if o == nil || o.Lon.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Lon.Get()
}

// GetLonOk returns a tuple with the Lon field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetLonOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Lon.Get(), o.Lon.IsSet()
}

// SetLon sets field value
func (o *Location) SetLon(v float64) {
	o.Lon.Set(&v)
}

// GetStoreNumber returns the StoreNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Location) GetStoreNumber() string {
	if o == nil || o.StoreNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.StoreNumber.Get()
}

// GetStoreNumberOk returns a tuple with the StoreNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetStoreNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StoreNumber.Get(), o.StoreNumber.IsSet()
}

// SetStoreNumber sets field value
func (o *Location) SetStoreNumber(v string) {
	o.StoreNumber.Set(&v)
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address.Get()
	}
	if true {
		toSerialize["city"] = o.City.Get()
	}
	if true {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if true {
		toSerialize["country"] = o.Country.Get()
	}
	if true {
		toSerialize["lat"] = o.Lat.Get()
	}
	if true {
		toSerialize["lon"] = o.Lon.Get()
	}
	if true {
		toSerialize["store_number"] = o.StoreNumber.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Location) UnmarshalJSON(bytes []byte) (err error) {
	varLocation := _Location{}

	if err = json.Unmarshal(bytes, &varLocation); err == nil {
		*o = Location(varLocation)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		delete(additionalProperties, "lat")
		delete(additionalProperties, "lon")
		delete(additionalProperties, "store_number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


