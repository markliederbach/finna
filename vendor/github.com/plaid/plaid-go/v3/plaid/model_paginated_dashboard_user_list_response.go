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

// PaginatedDashboardUserListResponse Paginated list of dashboard users
type PaginatedDashboardUserListResponse struct {
	// List of dashboard users
	DashboardUsers []DashboardUser `json:"dashboard_users"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedDashboardUserListResponse PaginatedDashboardUserListResponse

// NewPaginatedDashboardUserListResponse instantiates a new PaginatedDashboardUserListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedDashboardUserListResponse(dashboardUsers []DashboardUser, nextCursor NullableString, requestId string) *PaginatedDashboardUserListResponse {
	this := PaginatedDashboardUserListResponse{}
	this.DashboardUsers = dashboardUsers
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewPaginatedDashboardUserListResponseWithDefaults instantiates a new PaginatedDashboardUserListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedDashboardUserListResponseWithDefaults() *PaginatedDashboardUserListResponse {
	this := PaginatedDashboardUserListResponse{}
	return &this
}

// GetDashboardUsers returns the DashboardUsers field value
func (o *PaginatedDashboardUserListResponse) GetDashboardUsers() []DashboardUser {
	if o == nil {
		var ret []DashboardUser
		return ret
	}

	return o.DashboardUsers
}

// GetDashboardUsersOk returns a tuple with the DashboardUsers field value
// and a boolean to check if the value has been set.
func (o *PaginatedDashboardUserListResponse) GetDashboardUsersOk() (*[]DashboardUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DashboardUsers, true
}

// SetDashboardUsers sets field value
func (o *PaginatedDashboardUserListResponse) SetDashboardUsers(v []DashboardUser) {
	o.DashboardUsers = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginatedDashboardUserListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedDashboardUserListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *PaginatedDashboardUserListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *PaginatedDashboardUserListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaginatedDashboardUserListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaginatedDashboardUserListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaginatedDashboardUserListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dashboard_users"] = o.DashboardUsers
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

func (o *PaginatedDashboardUserListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaginatedDashboardUserListResponse := _PaginatedDashboardUserListResponse{}

	if err = json.Unmarshal(bytes, &varPaginatedDashboardUserListResponse); err == nil {
		*o = PaginatedDashboardUserListResponse(varPaginatedDashboardUserListResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dashboard_users")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedDashboardUserListResponse struct {
	value *PaginatedDashboardUserListResponse
	isSet bool
}

func (v NullablePaginatedDashboardUserListResponse) Get() *PaginatedDashboardUserListResponse {
	return v.value
}

func (v *NullablePaginatedDashboardUserListResponse) Set(val *PaginatedDashboardUserListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedDashboardUserListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedDashboardUserListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedDashboardUserListResponse(val *PaginatedDashboardUserListResponse) *NullablePaginatedDashboardUserListResponse {
	return &NullablePaginatedDashboardUserListResponse{value: val, isSet: true}
}

func (v NullablePaginatedDashboardUserListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedDashboardUserListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


