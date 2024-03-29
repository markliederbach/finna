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

// ListEntityWatchlistScreeningRequest Request input for listing entity watchlist screenings
type ListEntityWatchlistScreeningRequest struct {
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// ID of the associated entity program.
	EntityWatchlistProgramId string `json:"entity_watchlist_program_id"`
	ClientUserId NullableString `json:"client_user_id,omitempty"`
	Status NullableWatchlistScreeningStatus `json:"status,omitempty"`
	Assignee NullableString `json:"assignee,omitempty"`
	// An identifier that determines which page of results you receive.
	Cursor NullableString `json:"cursor,omitempty"`
}

// NewListEntityWatchlistScreeningRequest instantiates a new ListEntityWatchlistScreeningRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEntityWatchlistScreeningRequest(entityWatchlistProgramId string) *ListEntityWatchlistScreeningRequest {
	this := ListEntityWatchlistScreeningRequest{}
	this.EntityWatchlistProgramId = entityWatchlistProgramId
	return &this
}

// NewListEntityWatchlistScreeningRequestWithDefaults instantiates a new ListEntityWatchlistScreeningRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEntityWatchlistScreeningRequestWithDefaults() *ListEntityWatchlistScreeningRequest {
	this := ListEntityWatchlistScreeningRequest{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ListEntityWatchlistScreeningRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntityWatchlistScreeningRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ListEntityWatchlistScreeningRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ListEntityWatchlistScreeningRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ListEntityWatchlistScreeningRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntityWatchlistScreeningRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ListEntityWatchlistScreeningRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ListEntityWatchlistScreeningRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetEntityWatchlistProgramId returns the EntityWatchlistProgramId field value
func (o *ListEntityWatchlistScreeningRequest) GetEntityWatchlistProgramId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityWatchlistProgramId
}

// GetEntityWatchlistProgramIdOk returns a tuple with the EntityWatchlistProgramId field value
// and a boolean to check if the value has been set.
func (o *ListEntityWatchlistScreeningRequest) GetEntityWatchlistProgramIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistProgramId, true
}

// SetEntityWatchlistProgramId sets field value
func (o *ListEntityWatchlistScreeningRequest) SetEntityWatchlistProgramId(v string) {
	o.EntityWatchlistProgramId = v
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListEntityWatchlistScreeningRequest) GetClientUserId() string {
	if o == nil || o.ClientUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientUserId.Get()
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListEntityWatchlistScreeningRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientUserId.Get(), o.ClientUserId.IsSet()
}

// HasClientUserId returns a boolean if a field has been set.
func (o *ListEntityWatchlistScreeningRequest) HasClientUserId() bool {
	if o != nil && o.ClientUserId.IsSet() {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given NullableString and assigns it to the ClientUserId field.
func (o *ListEntityWatchlistScreeningRequest) SetClientUserId(v string) {
	o.ClientUserId.Set(&v)
}
// SetClientUserIdNil sets the value for ClientUserId to be an explicit nil
func (o *ListEntityWatchlistScreeningRequest) SetClientUserIdNil() {
	o.ClientUserId.Set(nil)
}

// UnsetClientUserId ensures that no value is present for ClientUserId, not even an explicit nil
func (o *ListEntityWatchlistScreeningRequest) UnsetClientUserId() {
	o.ClientUserId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListEntityWatchlistScreeningRequest) GetStatus() WatchlistScreeningStatus {
	if o == nil || o.Status.Get() == nil {
		var ret WatchlistScreeningStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListEntityWatchlistScreeningRequest) GetStatusOk() (*WatchlistScreeningStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ListEntityWatchlistScreeningRequest) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableWatchlistScreeningStatus and assigns it to the Status field.
func (o *ListEntityWatchlistScreeningRequest) SetStatus(v WatchlistScreeningStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ListEntityWatchlistScreeningRequest) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ListEntityWatchlistScreeningRequest) UnsetStatus() {
	o.Status.Unset()
}

// GetAssignee returns the Assignee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListEntityWatchlistScreeningRequest) GetAssignee() string {
	if o == nil || o.Assignee.Get() == nil {
		var ret string
		return ret
	}
	return *o.Assignee.Get()
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListEntityWatchlistScreeningRequest) GetAssigneeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Assignee.Get(), o.Assignee.IsSet()
}

// HasAssignee returns a boolean if a field has been set.
func (o *ListEntityWatchlistScreeningRequest) HasAssignee() bool {
	if o != nil && o.Assignee.IsSet() {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given NullableString and assigns it to the Assignee field.
func (o *ListEntityWatchlistScreeningRequest) SetAssignee(v string) {
	o.Assignee.Set(&v)
}
// SetAssigneeNil sets the value for Assignee to be an explicit nil
func (o *ListEntityWatchlistScreeningRequest) SetAssigneeNil() {
	o.Assignee.Set(nil)
}

// UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
func (o *ListEntityWatchlistScreeningRequest) UnsetAssignee() {
	o.Assignee.Unset()
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListEntityWatchlistScreeningRequest) GetCursor() string {
	if o == nil || o.Cursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListEntityWatchlistScreeningRequest) GetCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *ListEntityWatchlistScreeningRequest) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableString and assigns it to the Cursor field.
func (o *ListEntityWatchlistScreeningRequest) SetCursor(v string) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *ListEntityWatchlistScreeningRequest) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *ListEntityWatchlistScreeningRequest) UnsetCursor() {
	o.Cursor.Unset()
}

func (o ListEntityWatchlistScreeningRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["entity_watchlist_program_id"] = o.EntityWatchlistProgramId
	}
	if o.ClientUserId.IsSet() {
		toSerialize["client_user_id"] = o.ClientUserId.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Assignee.IsSet() {
		toSerialize["assignee"] = o.Assignee.Get()
	}
	if o.Cursor.IsSet() {
		toSerialize["cursor"] = o.Cursor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableListEntityWatchlistScreeningRequest struct {
	value *ListEntityWatchlistScreeningRequest
	isSet bool
}

func (v NullableListEntityWatchlistScreeningRequest) Get() *ListEntityWatchlistScreeningRequest {
	return v.value
}

func (v *NullableListEntityWatchlistScreeningRequest) Set(val *ListEntityWatchlistScreeningRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListEntityWatchlistScreeningRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListEntityWatchlistScreeningRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEntityWatchlistScreeningRequest(val *ListEntityWatchlistScreeningRequest) *NullableListEntityWatchlistScreeningRequest {
	return &NullableListEntityWatchlistScreeningRequest{value: val, isSet: true}
}

func (v NullableListEntityWatchlistScreeningRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEntityWatchlistScreeningRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


