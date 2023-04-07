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

// DocumentaryVerificationDocument Images, extracted data, and analysis from a user's identity document
type DocumentaryVerificationDocument struct {
	Status DocumentStatus `json:"status"`
	// The `attempt` field begins with 1 and increments with each subsequent document upload.
	Attempt float32 `json:"attempt"`
	Images PhysicalDocumentImages `json:"images"`
	ExtractedData NullablePhysicalDocumentExtractedData `json:"extracted_data"`
	Analysis DocumentAnalysis `json:"analysis"`
}

// NewDocumentaryVerificationDocument instantiates a new DocumentaryVerificationDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentaryVerificationDocument(status DocumentStatus, attempt float32, images PhysicalDocumentImages, extractedData NullablePhysicalDocumentExtractedData, analysis DocumentAnalysis) *DocumentaryVerificationDocument {
	this := DocumentaryVerificationDocument{}
	this.Status = status
	this.Attempt = attempt
	this.Images = images
	this.ExtractedData = extractedData
	this.Analysis = analysis
	return &this
}

// NewDocumentaryVerificationDocumentWithDefaults instantiates a new DocumentaryVerificationDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentaryVerificationDocumentWithDefaults() *DocumentaryVerificationDocument {
	this := DocumentaryVerificationDocument{}
	return &this
}

// GetStatus returns the Status field value
func (o *DocumentaryVerificationDocument) GetStatus() DocumentStatus {
	if o == nil {
		var ret DocumentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DocumentaryVerificationDocument) GetStatusOk() (*DocumentStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DocumentaryVerificationDocument) SetStatus(v DocumentStatus) {
	o.Status = v
}

// GetAttempt returns the Attempt field value
func (o *DocumentaryVerificationDocument) GetAttempt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Attempt
}

// GetAttemptOk returns a tuple with the Attempt field value
// and a boolean to check if the value has been set.
func (o *DocumentaryVerificationDocument) GetAttemptOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attempt, true
}

// SetAttempt sets field value
func (o *DocumentaryVerificationDocument) SetAttempt(v float32) {
	o.Attempt = v
}

// GetImages returns the Images field value
func (o *DocumentaryVerificationDocument) GetImages() PhysicalDocumentImages {
	if o == nil {
		var ret PhysicalDocumentImages
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *DocumentaryVerificationDocument) GetImagesOk() (*PhysicalDocumentImages, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Images, true
}

// SetImages sets field value
func (o *DocumentaryVerificationDocument) SetImages(v PhysicalDocumentImages) {
	o.Images = v
}

// GetExtractedData returns the ExtractedData field value
// If the value is explicit nil, the zero value for PhysicalDocumentExtractedData will be returned
func (o *DocumentaryVerificationDocument) GetExtractedData() PhysicalDocumentExtractedData {
	if o == nil || o.ExtractedData.Get() == nil {
		var ret PhysicalDocumentExtractedData
		return ret
	}

	return *o.ExtractedData.Get()
}

// GetExtractedDataOk returns a tuple with the ExtractedData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentaryVerificationDocument) GetExtractedDataOk() (*PhysicalDocumentExtractedData, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtractedData.Get(), o.ExtractedData.IsSet()
}

// SetExtractedData sets field value
func (o *DocumentaryVerificationDocument) SetExtractedData(v PhysicalDocumentExtractedData) {
	o.ExtractedData.Set(&v)
}

// GetAnalysis returns the Analysis field value
func (o *DocumentaryVerificationDocument) GetAnalysis() DocumentAnalysis {
	if o == nil {
		var ret DocumentAnalysis
		return ret
	}

	return o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value
// and a boolean to check if the value has been set.
func (o *DocumentaryVerificationDocument) GetAnalysisOk() (*DocumentAnalysis, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Analysis, true
}

// SetAnalysis sets field value
func (o *DocumentaryVerificationDocument) SetAnalysis(v DocumentAnalysis) {
	o.Analysis = v
}

func (o DocumentaryVerificationDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["attempt"] = o.Attempt
	}
	if true {
		toSerialize["images"] = o.Images
	}
	if true {
		toSerialize["extracted_data"] = o.ExtractedData.Get()
	}
	if true {
		toSerialize["analysis"] = o.Analysis
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentaryVerificationDocument struct {
	value *DocumentaryVerificationDocument
	isSet bool
}

func (v NullableDocumentaryVerificationDocument) Get() *DocumentaryVerificationDocument {
	return v.value
}

func (v *NullableDocumentaryVerificationDocument) Set(val *DocumentaryVerificationDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentaryVerificationDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentaryVerificationDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentaryVerificationDocument(val *DocumentaryVerificationDocument) *NullableDocumentaryVerificationDocument {
	return &NullableDocumentaryVerificationDocument{value: val, isSet: true}
}

func (v NullableDocumentaryVerificationDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentaryVerificationDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

