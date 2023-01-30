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

// EntityScreeningHitPhoneNumbers Phone number information associated with the entity screening hit
type EntityScreeningHitPhoneNumbers struct {
	Type PhoneType `json:"type"`
	// A phone number in E.164 format.
	PhoneNumber string `json:"phone_number"`
}

// NewEntityScreeningHitPhoneNumbers instantiates a new EntityScreeningHitPhoneNumbers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityScreeningHitPhoneNumbers(type_ PhoneType, phoneNumber string) *EntityScreeningHitPhoneNumbers {
	this := EntityScreeningHitPhoneNumbers{}
	this.Type = type_
	this.PhoneNumber = phoneNumber
	return &this
}

// NewEntityScreeningHitPhoneNumbersWithDefaults instantiates a new EntityScreeningHitPhoneNumbers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityScreeningHitPhoneNumbersWithDefaults() *EntityScreeningHitPhoneNumbers {
	this := EntityScreeningHitPhoneNumbers{}
	return &this
}

// GetType returns the Type field value
func (o *EntityScreeningHitPhoneNumbers) GetType() PhoneType {
	if o == nil {
		var ret PhoneType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitPhoneNumbers) GetTypeOk() (*PhoneType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EntityScreeningHitPhoneNumbers) SetType(v PhoneType) {
	o.Type = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *EntityScreeningHitPhoneNumbers) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitPhoneNumbers) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *EntityScreeningHitPhoneNumbers) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

func (o EntityScreeningHitPhoneNumbers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	return json.Marshal(toSerialize)
}

type NullableEntityScreeningHitPhoneNumbers struct {
	value *EntityScreeningHitPhoneNumbers
	isSet bool
}

func (v NullableEntityScreeningHitPhoneNumbers) Get() *EntityScreeningHitPhoneNumbers {
	return v.value
}

func (v *NullableEntityScreeningHitPhoneNumbers) Set(val *EntityScreeningHitPhoneNumbers) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityScreeningHitPhoneNumbers) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityScreeningHitPhoneNumbers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityScreeningHitPhoneNumbers(val *EntityScreeningHitPhoneNumbers) *NullableEntityScreeningHitPhoneNumbers {
	return &NullableEntityScreeningHitPhoneNumbers{value: val, isSet: true}
}

func (v NullableEntityScreeningHitPhoneNumbers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityScreeningHitPhoneNumbers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


