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

// AddressData Data about the components comprising an address.
type AddressData struct {
	// The full city name
	City string `json:"city"`
	// The region or state. In API versions 2018-05-22 and earlier, this field is called `state`. Example: `\"NC\"`
	Region NullableString `json:"region"`
	// The full street address Example: `\"564 Main Street, APT 15\"`
	Street string `json:"street"`
	// The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.
	PostalCode NullableString `json:"postal_code"`
	// The ISO 3166-1 alpha-2 country code
	Country NullableString `json:"country"`
	AdditionalProperties map[string]interface{}
}

type _AddressData AddressData

// NewAddressData instantiates a new AddressData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressData(city string, region NullableString, street string, postalCode NullableString, country NullableString) *AddressData {
	this := AddressData{}
	this.City = city
	this.Region = region
	this.Street = street
	this.PostalCode = postalCode
	this.Country = country
	return &this
}

// NewAddressDataWithDefaults instantiates a new AddressData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressDataWithDefaults() *AddressData {
	this := AddressData{}
	return &this
}

// GetCity returns the City field value
func (o *AddressData) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *AddressData) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *AddressData) SetCity(v string) {
	o.City = v
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AddressData) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressData) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *AddressData) SetRegion(v string) {
	o.Region.Set(&v)
}

// GetStreet returns the Street field value
func (o *AddressData) GetStreet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Street
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
func (o *AddressData) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Street, true
}

// SetStreet sets field value
func (o *AddressData) SetStreet(v string) {
	o.Street = v
}

// GetPostalCode returns the PostalCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AddressData) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressData) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// SetPostalCode sets field value
func (o *AddressData) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AddressData) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressData) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *AddressData) SetCountry(v string) {
	o.Country.Set(&v)
}

func (o AddressData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["street"] = o.Street
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if true {
		toSerialize["country"] = o.Country.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AddressData) UnmarshalJSON(bytes []byte) (err error) {
	varAddressData := _AddressData{}

	if err = json.Unmarshal(bytes, &varAddressData); err == nil {
		*o = AddressData(varAddressData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		delete(additionalProperties, "street")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddressData struct {
	value *AddressData
	isSet bool
}

func (v NullableAddressData) Get() *AddressData {
	return v.value
}

func (v *NullableAddressData) Set(val *AddressData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressData(val *AddressData) *NullableAddressData {
	return &NullableAddressData{value: val, isSet: true}
}

func (v NullableAddressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


