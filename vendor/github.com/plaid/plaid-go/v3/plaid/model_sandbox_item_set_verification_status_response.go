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

// SandboxItemSetVerificationStatusResponse SandboxItemSetVerificationStatusResponse defines the response schema for `/sandbox/item/set_verification_status`
type SandboxItemSetVerificationStatusResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SandboxItemSetVerificationStatusResponse SandboxItemSetVerificationStatusResponse

// NewSandboxItemSetVerificationStatusResponse instantiates a new SandboxItemSetVerificationStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxItemSetVerificationStatusResponse(requestId string) *SandboxItemSetVerificationStatusResponse {
	this := SandboxItemSetVerificationStatusResponse{}
	this.RequestId = requestId
	return &this
}

// NewSandboxItemSetVerificationStatusResponseWithDefaults instantiates a new SandboxItemSetVerificationStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxItemSetVerificationStatusResponseWithDefaults() *SandboxItemSetVerificationStatusResponse {
	this := SandboxItemSetVerificationStatusResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SandboxItemSetVerificationStatusResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxItemSetVerificationStatusResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxItemSetVerificationStatusResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxItemSetVerificationStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxItemSetVerificationStatusResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxItemSetVerificationStatusResponse := _SandboxItemSetVerificationStatusResponse{}

	if err = json.Unmarshal(bytes, &varSandboxItemSetVerificationStatusResponse); err == nil {
		*o = SandboxItemSetVerificationStatusResponse(varSandboxItemSetVerificationStatusResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxItemSetVerificationStatusResponse struct {
	value *SandboxItemSetVerificationStatusResponse
	isSet bool
}

func (v NullableSandboxItemSetVerificationStatusResponse) Get() *SandboxItemSetVerificationStatusResponse {
	return v.value
}

func (v *NullableSandboxItemSetVerificationStatusResponse) Set(val *SandboxItemSetVerificationStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxItemSetVerificationStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxItemSetVerificationStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxItemSetVerificationStatusResponse(val *SandboxItemSetVerificationStatusResponse) *NullableSandboxItemSetVerificationStatusResponse {
	return &NullableSandboxItemSetVerificationStatusResponse{value: val, isSet: true}
}

func (v NullableSandboxItemSetVerificationStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxItemSetVerificationStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


