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

// DocumentRiskSignal Details about a certain reason as to why a document could potentially be fraudulent.
type DocumentRiskSignal struct {
	// The result from the risk signal check.
	Type NullableString `json:"type,omitempty"`
	// The field which the risk signal was computed for
	Field NullableString `json:"field,omitempty"`
	// A flag used to quickly identify if the signal indicates that this field is authentic or fraudulent
	HasFraudRisk NullableBool `json:"has_fraud_risk,omitempty"`
	InstitutionMetadata *DocumentRiskSignalInstitutionMetadata `json:"institution_metadata,omitempty"`
	// The expected value of the field, as seen on the document
	ExpectedValue NullableString `json:"expected_value,omitempty"`
	// The derived value obtained in the risk signal calculation process for this field
	ActualValue NullableString `json:"actual_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DocumentRiskSignal DocumentRiskSignal

// NewDocumentRiskSignal instantiates a new DocumentRiskSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentRiskSignal() *DocumentRiskSignal {
	this := DocumentRiskSignal{}
	return &this
}

// NewDocumentRiskSignalWithDefaults instantiates a new DocumentRiskSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentRiskSignalWithDefaults() *DocumentRiskSignal {
	this := DocumentRiskSignal{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentRiskSignal) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentRiskSignal) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *DocumentRiskSignal) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *DocumentRiskSignal) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *DocumentRiskSignal) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *DocumentRiskSignal) UnsetType() {
	o.Type.Unset()
}

// GetField returns the Field field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentRiskSignal) GetField() string {
	if o == nil || o.Field.Get() == nil {
		var ret string
		return ret
	}
	return *o.Field.Get()
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentRiskSignal) GetFieldOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Field.Get(), o.Field.IsSet()
}

// HasField returns a boolean if a field has been set.
func (o *DocumentRiskSignal) HasField() bool {
	if o != nil && o.Field.IsSet() {
		return true
	}

	return false
}

// SetField gets a reference to the given NullableString and assigns it to the Field field.
func (o *DocumentRiskSignal) SetField(v string) {
	o.Field.Set(&v)
}
// SetFieldNil sets the value for Field to be an explicit nil
func (o *DocumentRiskSignal) SetFieldNil() {
	o.Field.Set(nil)
}

// UnsetField ensures that no value is present for Field, not even an explicit nil
func (o *DocumentRiskSignal) UnsetField() {
	o.Field.Unset()
}

// GetHasFraudRisk returns the HasFraudRisk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentRiskSignal) GetHasFraudRisk() bool {
	if o == nil || o.HasFraudRisk.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasFraudRisk.Get()
}

// GetHasFraudRiskOk returns a tuple with the HasFraudRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentRiskSignal) GetHasFraudRiskOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasFraudRisk.Get(), o.HasFraudRisk.IsSet()
}

// HasHasFraudRisk returns a boolean if a field has been set.
func (o *DocumentRiskSignal) HasHasFraudRisk() bool {
	if o != nil && o.HasFraudRisk.IsSet() {
		return true
	}

	return false
}

// SetHasFraudRisk gets a reference to the given NullableBool and assigns it to the HasFraudRisk field.
func (o *DocumentRiskSignal) SetHasFraudRisk(v bool) {
	o.HasFraudRisk.Set(&v)
}
// SetHasFraudRiskNil sets the value for HasFraudRisk to be an explicit nil
func (o *DocumentRiskSignal) SetHasFraudRiskNil() {
	o.HasFraudRisk.Set(nil)
}

// UnsetHasFraudRisk ensures that no value is present for HasFraudRisk, not even an explicit nil
func (o *DocumentRiskSignal) UnsetHasFraudRisk() {
	o.HasFraudRisk.Unset()
}

// GetInstitutionMetadata returns the InstitutionMetadata field value if set, zero value otherwise.
func (o *DocumentRiskSignal) GetInstitutionMetadata() DocumentRiskSignalInstitutionMetadata {
	if o == nil || o.InstitutionMetadata == nil {
		var ret DocumentRiskSignalInstitutionMetadata
		return ret
	}
	return *o.InstitutionMetadata
}

// GetInstitutionMetadataOk returns a tuple with the InstitutionMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentRiskSignal) GetInstitutionMetadataOk() (*DocumentRiskSignalInstitutionMetadata, bool) {
	if o == nil || o.InstitutionMetadata == nil {
		return nil, false
	}
	return o.InstitutionMetadata, true
}

// HasInstitutionMetadata returns a boolean if a field has been set.
func (o *DocumentRiskSignal) HasInstitutionMetadata() bool {
	if o != nil && o.InstitutionMetadata != nil {
		return true
	}

	return false
}

// SetInstitutionMetadata gets a reference to the given DocumentRiskSignalInstitutionMetadata and assigns it to the InstitutionMetadata field.
func (o *DocumentRiskSignal) SetInstitutionMetadata(v DocumentRiskSignalInstitutionMetadata) {
	o.InstitutionMetadata = &v
}

// GetExpectedValue returns the ExpectedValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentRiskSignal) GetExpectedValue() string {
	if o == nil || o.ExpectedValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExpectedValue.Get()
}

// GetExpectedValueOk returns a tuple with the ExpectedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentRiskSignal) GetExpectedValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpectedValue.Get(), o.ExpectedValue.IsSet()
}

// HasExpectedValue returns a boolean if a field has been set.
func (o *DocumentRiskSignal) HasExpectedValue() bool {
	if o != nil && o.ExpectedValue.IsSet() {
		return true
	}

	return false
}

// SetExpectedValue gets a reference to the given NullableString and assigns it to the ExpectedValue field.
func (o *DocumentRiskSignal) SetExpectedValue(v string) {
	o.ExpectedValue.Set(&v)
}
// SetExpectedValueNil sets the value for ExpectedValue to be an explicit nil
func (o *DocumentRiskSignal) SetExpectedValueNil() {
	o.ExpectedValue.Set(nil)
}

// UnsetExpectedValue ensures that no value is present for ExpectedValue, not even an explicit nil
func (o *DocumentRiskSignal) UnsetExpectedValue() {
	o.ExpectedValue.Unset()
}

// GetActualValue returns the ActualValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentRiskSignal) GetActualValue() string {
	if o == nil || o.ActualValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.ActualValue.Get()
}

// GetActualValueOk returns a tuple with the ActualValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentRiskSignal) GetActualValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActualValue.Get(), o.ActualValue.IsSet()
}

// HasActualValue returns a boolean if a field has been set.
func (o *DocumentRiskSignal) HasActualValue() bool {
	if o != nil && o.ActualValue.IsSet() {
		return true
	}

	return false
}

// SetActualValue gets a reference to the given NullableString and assigns it to the ActualValue field.
func (o *DocumentRiskSignal) SetActualValue(v string) {
	o.ActualValue.Set(&v)
}
// SetActualValueNil sets the value for ActualValue to be an explicit nil
func (o *DocumentRiskSignal) SetActualValueNil() {
	o.ActualValue.Set(nil)
}

// UnsetActualValue ensures that no value is present for ActualValue, not even an explicit nil
func (o *DocumentRiskSignal) UnsetActualValue() {
	o.ActualValue.Unset()
}

func (o DocumentRiskSignal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Field.IsSet() {
		toSerialize["field"] = o.Field.Get()
	}
	if o.HasFraudRisk.IsSet() {
		toSerialize["has_fraud_risk"] = o.HasFraudRisk.Get()
	}
	if o.InstitutionMetadata != nil {
		toSerialize["institution_metadata"] = o.InstitutionMetadata
	}
	if o.ExpectedValue.IsSet() {
		toSerialize["expected_value"] = o.ExpectedValue.Get()
	}
	if o.ActualValue.IsSet() {
		toSerialize["actual_value"] = o.ActualValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DocumentRiskSignal) UnmarshalJSON(bytes []byte) (err error) {
	varDocumentRiskSignal := _DocumentRiskSignal{}

	if err = json.Unmarshal(bytes, &varDocumentRiskSignal); err == nil {
		*o = DocumentRiskSignal(varDocumentRiskSignal)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "field")
		delete(additionalProperties, "has_fraud_risk")
		delete(additionalProperties, "institution_metadata")
		delete(additionalProperties, "expected_value")
		delete(additionalProperties, "actual_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDocumentRiskSignal struct {
	value *DocumentRiskSignal
	isSet bool
}

func (v NullableDocumentRiskSignal) Get() *DocumentRiskSignal {
	return v.value
}

func (v *NullableDocumentRiskSignal) Set(val *DocumentRiskSignal) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentRiskSignal) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentRiskSignal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentRiskSignal(val *DocumentRiskSignal) *NullableDocumentRiskSignal {
	return &NullableDocumentRiskSignal{value: val, isSet: true}
}

func (v NullableDocumentRiskSignal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentRiskSignal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


