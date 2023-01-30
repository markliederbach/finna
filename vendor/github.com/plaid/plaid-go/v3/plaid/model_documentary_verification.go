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

// DocumentaryVerification data, images, analysis, and results from the `documentary_verification` step.
type DocumentaryVerification struct {
	// The outcome status for the associated Identity Verification attempt's `documentary_verification` step. This field will always have the same value as `steps.documentary_verification`.
	Status string `json:"status"`
	// An array of documents submitted to the `documentary_verification` step. Each entry represents one user submission, where each submission will contain both a front and back image, or just a front image, depending on the document type.  Note: Plaid will automatically let a user submit a new set of document images up to three times if we detect that a previous attempt might have failed due to user error. For example, if the first set of document images are blurry or obscured by glare, the user will be asked to capture their documents again, resulting in at least two separate entries within `documents`. If the overall `documentary_verification` is `failed`, the user has exhausted their retry attempts.
	Documents []DocumentaryVerificationDocument `json:"documents"`
}

// NewDocumentaryVerification instantiates a new DocumentaryVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentaryVerification(status string, documents []DocumentaryVerificationDocument) *DocumentaryVerification {
	this := DocumentaryVerification{}
	this.Status = status
	this.Documents = documents
	return &this
}

// NewDocumentaryVerificationWithDefaults instantiates a new DocumentaryVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentaryVerificationWithDefaults() *DocumentaryVerification {
	this := DocumentaryVerification{}
	return &this
}

// GetStatus returns the Status field value
func (o *DocumentaryVerification) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DocumentaryVerification) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DocumentaryVerification) SetStatus(v string) {
	o.Status = v
}

// GetDocuments returns the Documents field value
func (o *DocumentaryVerification) GetDocuments() []DocumentaryVerificationDocument {
	if o == nil {
		var ret []DocumentaryVerificationDocument
		return ret
	}

	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value
// and a boolean to check if the value has been set.
func (o *DocumentaryVerification) GetDocumentsOk() (*[]DocumentaryVerificationDocument, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Documents, true
}

// SetDocuments sets field value
func (o *DocumentaryVerification) SetDocuments(v []DocumentaryVerificationDocument) {
	o.Documents = v
}

func (o DocumentaryVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["documents"] = o.Documents
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentaryVerification struct {
	value *DocumentaryVerification
	isSet bool
}

func (v NullableDocumentaryVerification) Get() *DocumentaryVerification {
	return v.value
}

func (v *NullableDocumentaryVerification) Set(val *DocumentaryVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentaryVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentaryVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentaryVerification(val *DocumentaryVerification) *NullableDocumentaryVerification {
	return &NullableDocumentaryVerification{value: val, isSet: true}
}

func (v NullableDocumentaryVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentaryVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


