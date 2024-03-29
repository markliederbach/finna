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

// PaginatedEntityWatchlistProgramListResponse Paginated list of entity watchlist screening programs
type PaginatedEntityWatchlistProgramListResponse struct {
	// List of entity watchlist screening programs
	EntityWatchlistPrograms []EntityWatchlistProgram `json:"entity_watchlist_programs"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedEntityWatchlistProgramListResponse PaginatedEntityWatchlistProgramListResponse

// NewPaginatedEntityWatchlistProgramListResponse instantiates a new PaginatedEntityWatchlistProgramListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEntityWatchlistProgramListResponse(entityWatchlistPrograms []EntityWatchlistProgram, nextCursor NullableString, requestId string) *PaginatedEntityWatchlistProgramListResponse {
	this := PaginatedEntityWatchlistProgramListResponse{}
	this.EntityWatchlistPrograms = entityWatchlistPrograms
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewPaginatedEntityWatchlistProgramListResponseWithDefaults instantiates a new PaginatedEntityWatchlistProgramListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEntityWatchlistProgramListResponseWithDefaults() *PaginatedEntityWatchlistProgramListResponse {
	this := PaginatedEntityWatchlistProgramListResponse{}
	return &this
}

// GetEntityWatchlistPrograms returns the EntityWatchlistPrograms field value
func (o *PaginatedEntityWatchlistProgramListResponse) GetEntityWatchlistPrograms() []EntityWatchlistProgram {
	if o == nil {
		var ret []EntityWatchlistProgram
		return ret
	}

	return o.EntityWatchlistPrograms
}

// GetEntityWatchlistProgramsOk returns a tuple with the EntityWatchlistPrograms field value
// and a boolean to check if the value has been set.
func (o *PaginatedEntityWatchlistProgramListResponse) GetEntityWatchlistProgramsOk() (*[]EntityWatchlistProgram, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistPrograms, true
}

// SetEntityWatchlistPrograms sets field value
func (o *PaginatedEntityWatchlistProgramListResponse) SetEntityWatchlistPrograms(v []EntityWatchlistProgram) {
	o.EntityWatchlistPrograms = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginatedEntityWatchlistProgramListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedEntityWatchlistProgramListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *PaginatedEntityWatchlistProgramListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *PaginatedEntityWatchlistProgramListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaginatedEntityWatchlistProgramListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaginatedEntityWatchlistProgramListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaginatedEntityWatchlistProgramListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_watchlist_programs"] = o.EntityWatchlistPrograms
	}
	if true {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaginatedEntityWatchlistProgramListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaginatedEntityWatchlistProgramListResponse := _PaginatedEntityWatchlistProgramListResponse{}

	if err = json.Unmarshal(bytes, &varPaginatedEntityWatchlistProgramListResponse); err == nil {
		*o = PaginatedEntityWatchlistProgramListResponse(varPaginatedEntityWatchlistProgramListResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entity_watchlist_programs")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedEntityWatchlistProgramListResponse struct {
	value *PaginatedEntityWatchlistProgramListResponse
	isSet bool
}

func (v NullablePaginatedEntityWatchlistProgramListResponse) Get() *PaginatedEntityWatchlistProgramListResponse {
	return v.value
}

func (v *NullablePaginatedEntityWatchlistProgramListResponse) Set(val *PaginatedEntityWatchlistProgramListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEntityWatchlistProgramListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEntityWatchlistProgramListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEntityWatchlistProgramListResponse(val *PaginatedEntityWatchlistProgramListResponse) *NullablePaginatedEntityWatchlistProgramListResponse {
	return &NullablePaginatedEntityWatchlistProgramListResponse{value: val, isSet: true}
}

func (v NullablePaginatedEntityWatchlistProgramListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEntityWatchlistProgramListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


