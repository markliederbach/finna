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

// ReportToken ReportToken is a representation of a token that has a `report_type` field that can be one of `assets` or `income` and a `token` field which is the associated token with the `report_type`. The `token` can be an Asset Report token or Income Report token.
type ReportToken struct {
	// The report type. It can be `assets` or `income`.
	ReportType *string `json:"report_type,omitempty"`
	// The report token. It can be an asset report token or an income report token.
	Token *string `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReportToken ReportToken

// NewReportToken instantiates a new ReportToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportToken() *ReportToken {
	this := ReportToken{}
	return &this
}

// NewReportTokenWithDefaults instantiates a new ReportToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportTokenWithDefaults() *ReportToken {
	this := ReportToken{}
	return &this
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *ReportToken) GetReportType() string {
	if o == nil || o.ReportType == nil {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportToken) GetReportTypeOk() (*string, bool) {
	if o == nil || o.ReportType == nil {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *ReportToken) HasReportType() bool {
	if o != nil && o.ReportType != nil {
		return true
	}

	return false
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *ReportToken) SetReportType(v string) {
	o.ReportType = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ReportToken) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportToken) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ReportToken) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ReportToken) SetToken(v string) {
	o.Token = &v
}

func (o ReportToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReportType != nil {
		toSerialize["report_type"] = o.ReportType
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReportToken) UnmarshalJSON(bytes []byte) (err error) {
	varReportToken := _ReportToken{}

	if err = json.Unmarshal(bytes, &varReportToken); err == nil {
		*o = ReportToken(varReportToken)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "report_type")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReportToken struct {
	value *ReportToken
	isSet bool
}

func (v NullableReportToken) Get() *ReportToken {
	return v.value
}

func (v *NullableReportToken) Set(val *ReportToken) {
	v.value = val
	v.isSet = true
}

func (v NullableReportToken) IsSet() bool {
	return v.isSet
}

func (v *NullableReportToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportToken(val *ReportToken) *NullableReportToken {
	return &NullableReportToken{value: val, isSet: true}
}

func (v NullableReportToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


