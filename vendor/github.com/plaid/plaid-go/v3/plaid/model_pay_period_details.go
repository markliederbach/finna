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

// PayPeriodDetails Details about the pay period.
type PayPeriodDetails struct {
	// The amount of the paycheck.
	CheckAmount NullableFloat64 `json:"check_amount,omitempty"`
	DistributionBreakdown *[]DistributionBreakdown `json:"distribution_breakdown,omitempty"`
	// The pay period end date, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format: \"yyyy-mm-dd\".
	EndDate NullableString `json:"end_date,omitempty"`
	// Total earnings before tax/deductions.
	GrossEarnings NullableFloat64 `json:"gross_earnings,omitempty"`
	// The date on which the paystub was issued, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (\"yyyy-mm-dd\").
	PayDate NullableString `json:"pay_date,omitempty"`
	// The frequency at which an individual is paid.
	PayFrequency NullableString `json:"pay_frequency,omitempty"`
	// The date on which the paystub was issued, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (\"yyyy-mm-dd\").
	PayDay NullableString `json:"pay_day,omitempty"`
	// The pay period start date, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format: \"yyyy-mm-dd\".
	StartDate NullableString `json:"start_date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PayPeriodDetails PayPeriodDetails

// NewPayPeriodDetails instantiates a new PayPeriodDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayPeriodDetails() *PayPeriodDetails {
	this := PayPeriodDetails{}
	return &this
}

// NewPayPeriodDetailsWithDefaults instantiates a new PayPeriodDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayPeriodDetailsWithDefaults() *PayPeriodDetails {
	this := PayPeriodDetails{}
	return &this
}

// GetCheckAmount returns the CheckAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayPeriodDetails) GetCheckAmount() float64 {
	if o == nil || o.CheckAmount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CheckAmount.Get()
}

// GetCheckAmountOk returns a tuple with the CheckAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayPeriodDetails) GetCheckAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CheckAmount.Get(), o.CheckAmount.IsSet()
}

// HasCheckAmount returns a boolean if a field has been set.
func (o *PayPeriodDetails) HasCheckAmount() bool {
	if o != nil && o.CheckAmount.IsSet() {
		return true
	}

	return false
}

// SetCheckAmount gets a reference to the given NullableFloat64 and assigns it to the CheckAmount field.
func (o *PayPeriodDetails) SetCheckAmount(v float64) {
	o.CheckAmount.Set(&v)
}
// SetCheckAmountNil sets the value for CheckAmount to be an explicit nil
func (o *PayPeriodDetails) SetCheckAmountNil() {
	o.CheckAmount.Set(nil)
}

// UnsetCheckAmount ensures that no value is present for CheckAmount, not even an explicit nil
func (o *PayPeriodDetails) UnsetCheckAmount() {
	o.CheckAmount.Unset()
}

// GetDistributionBreakdown returns the DistributionBreakdown field value if set, zero value otherwise.
func (o *PayPeriodDetails) GetDistributionBreakdown() []DistributionBreakdown {
	if o == nil || o.DistributionBreakdown == nil {
		var ret []DistributionBreakdown
		return ret
	}
	return *o.DistributionBreakdown
}

// GetDistributionBreakdownOk returns a tuple with the DistributionBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPeriodDetails) GetDistributionBreakdownOk() (*[]DistributionBreakdown, bool) {
	if o == nil || o.DistributionBreakdown == nil {
		return nil, false
	}
	return o.DistributionBreakdown, true
}

// HasDistributionBreakdown returns a boolean if a field has been set.
func (o *PayPeriodDetails) HasDistributionBreakdown() bool {
	if o != nil && o.DistributionBreakdown != nil {
		return true
	}

	return false
}

// SetDistributionBreakdown gets a reference to the given []DistributionBreakdown and assigns it to the DistributionBreakdown field.
func (o *PayPeriodDetails) SetDistributionBreakdown(v []DistributionBreakdown) {
	o.DistributionBreakdown = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayPeriodDetails) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayPeriodDetails) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *PayPeriodDetails) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *PayPeriodDetails) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *PayPeriodDetails) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *PayPeriodDetails) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetGrossEarnings returns the GrossEarnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayPeriodDetails) GetGrossEarnings() float64 {
	if o == nil || o.GrossEarnings.Get() == nil {
		var ret float64
		return ret
	}
	return *o.GrossEarnings.Get()
}

// GetGrossEarningsOk returns a tuple with the GrossEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayPeriodDetails) GetGrossEarningsOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GrossEarnings.Get(), o.GrossEarnings.IsSet()
}

// HasGrossEarnings returns a boolean if a field has been set.
func (o *PayPeriodDetails) HasGrossEarnings() bool {
	if o != nil && o.GrossEarnings.IsSet() {
		return true
	}

	return false
}

// SetGrossEarnings gets a reference to the given NullableFloat64 and assigns it to the GrossEarnings field.
func (o *PayPeriodDetails) SetGrossEarnings(v float64) {
	o.GrossEarnings.Set(&v)
}
// SetGrossEarningsNil sets the value for GrossEarnings to be an explicit nil
func (o *PayPeriodDetails) SetGrossEarningsNil() {
	o.GrossEarnings.Set(nil)
}

// UnsetGrossEarnings ensures that no value is present for GrossEarnings, not even an explicit nil
func (o *PayPeriodDetails) UnsetGrossEarnings() {
	o.GrossEarnings.Unset()
}

// GetPayDate returns the PayDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayPeriodDetails) GetPayDate() string {
	if o == nil || o.PayDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.PayDate.Get()
}

// GetPayDateOk returns a tuple with the PayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayPeriodDetails) GetPayDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayDate.Get(), o.PayDate.IsSet()
}

// HasPayDate returns a boolean if a field has been set.
func (o *PayPeriodDetails) HasPayDate() bool {
	if o != nil && o.PayDate.IsSet() {
		return true
	}

	return false
}

// SetPayDate gets a reference to the given NullableString and assigns it to the PayDate field.
func (o *PayPeriodDetails) SetPayDate(v string) {
	o.PayDate.Set(&v)
}
// SetPayDateNil sets the value for PayDate to be an explicit nil
func (o *PayPeriodDetails) SetPayDateNil() {
	o.PayDate.Set(nil)
}

// UnsetPayDate ensures that no value is present for PayDate, not even an explicit nil
func (o *PayPeriodDetails) UnsetPayDate() {
	o.PayDate.Unset()
}

// GetPayFrequency returns the PayFrequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayPeriodDetails) GetPayFrequency() string {
	if o == nil || o.PayFrequency.Get() == nil {
		var ret string
		return ret
	}
	return *o.PayFrequency.Get()
}

// GetPayFrequencyOk returns a tuple with the PayFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayPeriodDetails) GetPayFrequencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayFrequency.Get(), o.PayFrequency.IsSet()
}

// HasPayFrequency returns a boolean if a field has been set.
func (o *PayPeriodDetails) HasPayFrequency() bool {
	if o != nil && o.PayFrequency.IsSet() {
		return true
	}

	return false
}

// SetPayFrequency gets a reference to the given NullableString and assigns it to the PayFrequency field.
func (o *PayPeriodDetails) SetPayFrequency(v string) {
	o.PayFrequency.Set(&v)
}
// SetPayFrequencyNil sets the value for PayFrequency to be an explicit nil
func (o *PayPeriodDetails) SetPayFrequencyNil() {
	o.PayFrequency.Set(nil)
}

// UnsetPayFrequency ensures that no value is present for PayFrequency, not even an explicit nil
func (o *PayPeriodDetails) UnsetPayFrequency() {
	o.PayFrequency.Unset()
}

// GetPayDay returns the PayDay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayPeriodDetails) GetPayDay() string {
	if o == nil || o.PayDay.Get() == nil {
		var ret string
		return ret
	}
	return *o.PayDay.Get()
}

// GetPayDayOk returns a tuple with the PayDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayPeriodDetails) GetPayDayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayDay.Get(), o.PayDay.IsSet()
}

// HasPayDay returns a boolean if a field has been set.
func (o *PayPeriodDetails) HasPayDay() bool {
	if o != nil && o.PayDay.IsSet() {
		return true
	}

	return false
}

// SetPayDay gets a reference to the given NullableString and assigns it to the PayDay field.
func (o *PayPeriodDetails) SetPayDay(v string) {
	o.PayDay.Set(&v)
}
// SetPayDayNil sets the value for PayDay to be an explicit nil
func (o *PayPeriodDetails) SetPayDayNil() {
	o.PayDay.Set(nil)
}

// UnsetPayDay ensures that no value is present for PayDay, not even an explicit nil
func (o *PayPeriodDetails) UnsetPayDay() {
	o.PayDay.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayPeriodDetails) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayPeriodDetails) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *PayPeriodDetails) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableString and assigns it to the StartDate field.
func (o *PayPeriodDetails) SetStartDate(v string) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *PayPeriodDetails) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *PayPeriodDetails) UnsetStartDate() {
	o.StartDate.Unset()
}

func (o PayPeriodDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CheckAmount.IsSet() {
		toSerialize["check_amount"] = o.CheckAmount.Get()
	}
	if o.DistributionBreakdown != nil {
		toSerialize["distribution_breakdown"] = o.DistributionBreakdown
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.GrossEarnings.IsSet() {
		toSerialize["gross_earnings"] = o.GrossEarnings.Get()
	}
	if o.PayDate.IsSet() {
		toSerialize["pay_date"] = o.PayDate.Get()
	}
	if o.PayFrequency.IsSet() {
		toSerialize["pay_frequency"] = o.PayFrequency.Get()
	}
	if o.PayDay.IsSet() {
		toSerialize["pay_day"] = o.PayDay.Get()
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayPeriodDetails) UnmarshalJSON(bytes []byte) (err error) {
	varPayPeriodDetails := _PayPeriodDetails{}

	if err = json.Unmarshal(bytes, &varPayPeriodDetails); err == nil {
		*o = PayPeriodDetails(varPayPeriodDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "check_amount")
		delete(additionalProperties, "distribution_breakdown")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "gross_earnings")
		delete(additionalProperties, "pay_date")
		delete(additionalProperties, "pay_frequency")
		delete(additionalProperties, "pay_day")
		delete(additionalProperties, "start_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayPeriodDetails struct {
	value *PayPeriodDetails
	isSet bool
}

func (v NullablePayPeriodDetails) Get() *PayPeriodDetails {
	return v.value
}

func (v *NullablePayPeriodDetails) Set(val *PayPeriodDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePayPeriodDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePayPeriodDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayPeriodDetails(val *PayPeriodDetails) *NullablePayPeriodDetails {
	return &NullablePayPeriodDetails{value: val, isSet: true}
}

func (v NullablePayPeriodDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayPeriodDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


