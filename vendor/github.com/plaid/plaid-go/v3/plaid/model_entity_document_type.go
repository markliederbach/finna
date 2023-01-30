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
	"fmt"
)

// EntityDocumentType The kind of official document represented by this object.  `bik` - Russian bank code  `business_number` - A number that uniquely identifies the business within a category of businesses  `imo` - Number assigned to the entity by the International Maritime Organization  `other` - Any document not covered by other categories  `swift` - Number identifying a bank and branch.  `tax_id` - Identification issued for the purpose of collecting taxes
type EntityDocumentType string

// List of EntityDocumentType
const (
	ENTITYDOCUMENTTYPE_BIK EntityDocumentType = "bik"
	ENTITYDOCUMENTTYPE_BUSINESS_NUMBER EntityDocumentType = "business_number"
	ENTITYDOCUMENTTYPE_IMO EntityDocumentType = "imo"
	ENTITYDOCUMENTTYPE_OTHER EntityDocumentType = "other"
	ENTITYDOCUMENTTYPE_SWIFT EntityDocumentType = "swift"
	ENTITYDOCUMENTTYPE_TAX_ID EntityDocumentType = "tax_id"
)

var allowedEntityDocumentTypeEnumValues = []EntityDocumentType{
	"bik",
	"business_number",
	"imo",
	"other",
	"swift",
	"tax_id",
}

func (v *EntityDocumentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntityDocumentType(value)
	for _, existing := range allowedEntityDocumentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntityDocumentType", value)
}

// NewEntityDocumentTypeFromValue returns a pointer to a valid EntityDocumentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntityDocumentTypeFromValue(v string) (*EntityDocumentType, error) {
	ev := EntityDocumentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntityDocumentType: valid values are %v", v, allowedEntityDocumentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntityDocumentType) IsValid() bool {
	for _, existing := range allowedEntityDocumentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityDocumentType value
func (v EntityDocumentType) Ptr() *EntityDocumentType {
	return &v
}

type NullableEntityDocumentType struct {
	value *EntityDocumentType
	isSet bool
}

func (v NullableEntityDocumentType) Get() *EntityDocumentType {
	return v.value
}

func (v *NullableEntityDocumentType) Set(val *EntityDocumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityDocumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityDocumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityDocumentType(val *EntityDocumentType) *NullableEntityDocumentType {
	return &NullableEntityDocumentType{value: val, isSet: true}
}

func (v NullableEntityDocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityDocumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

