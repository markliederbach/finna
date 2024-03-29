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

// CreateEntityWatchlistScreeningReviewRequest Request input for creating a review for an entity screening
type CreateEntityWatchlistScreeningReviewRequest struct {
	// Hits to mark as a true positive after thorough manual review. These hits will never recur or be updated once dismissed. In most cases, confirmed hits indicate that the customer should be rejected.
	ConfirmedHits []string `json:"confirmed_hits"`
	// Hits to mark as a false positive after thorough manual review. These hits will never recur or be updated once dismissed.
	DismissedHits []string `json:"dismissed_hits"`
	// A comment submitted by a team member as part of reviewing a watchlist screening.
	Comment NullableString `json:"comment,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// ID of the associated entity screening.
	EntityWatchlistScreeningId string `json:"entity_watchlist_screening_id"`
}

// NewCreateEntityWatchlistScreeningReviewRequest instantiates a new CreateEntityWatchlistScreeningReviewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEntityWatchlistScreeningReviewRequest(confirmedHits []string, dismissedHits []string, entityWatchlistScreeningId string) *CreateEntityWatchlistScreeningReviewRequest {
	this := CreateEntityWatchlistScreeningReviewRequest{}
	this.ConfirmedHits = confirmedHits
	this.DismissedHits = dismissedHits
	this.EntityWatchlistScreeningId = entityWatchlistScreeningId
	return &this
}

// NewCreateEntityWatchlistScreeningReviewRequestWithDefaults instantiates a new CreateEntityWatchlistScreeningReviewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEntityWatchlistScreeningReviewRequestWithDefaults() *CreateEntityWatchlistScreeningReviewRequest {
	this := CreateEntityWatchlistScreeningReviewRequest{}
	return &this
}

// GetConfirmedHits returns the ConfirmedHits field value
func (o *CreateEntityWatchlistScreeningReviewRequest) GetConfirmedHits() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConfirmedHits
}

// GetConfirmedHitsOk returns a tuple with the ConfirmedHits field value
// and a boolean to check if the value has been set.
func (o *CreateEntityWatchlistScreeningReviewRequest) GetConfirmedHitsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfirmedHits, true
}

// SetConfirmedHits sets field value
func (o *CreateEntityWatchlistScreeningReviewRequest) SetConfirmedHits(v []string) {
	o.ConfirmedHits = v
}

// GetDismissedHits returns the DismissedHits field value
func (o *CreateEntityWatchlistScreeningReviewRequest) GetDismissedHits() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DismissedHits
}

// GetDismissedHitsOk returns a tuple with the DismissedHits field value
// and a boolean to check if the value has been set.
func (o *CreateEntityWatchlistScreeningReviewRequest) GetDismissedHitsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DismissedHits, true
}

// SetDismissedHits sets field value
func (o *CreateEntityWatchlistScreeningReviewRequest) SetDismissedHits(v []string) {
	o.DismissedHits = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEntityWatchlistScreeningReviewRequest) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEntityWatchlistScreeningReviewRequest) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *CreateEntityWatchlistScreeningReviewRequest) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *CreateEntityWatchlistScreeningReviewRequest) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *CreateEntityWatchlistScreeningReviewRequest) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *CreateEntityWatchlistScreeningReviewRequest) UnsetComment() {
	o.Comment.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreateEntityWatchlistScreeningReviewRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEntityWatchlistScreeningReviewRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreateEntityWatchlistScreeningReviewRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreateEntityWatchlistScreeningReviewRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreateEntityWatchlistScreeningReviewRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEntityWatchlistScreeningReviewRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreateEntityWatchlistScreeningReviewRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreateEntityWatchlistScreeningReviewRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetEntityWatchlistScreeningId returns the EntityWatchlistScreeningId field value
func (o *CreateEntityWatchlistScreeningReviewRequest) GetEntityWatchlistScreeningId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityWatchlistScreeningId
}

// GetEntityWatchlistScreeningIdOk returns a tuple with the EntityWatchlistScreeningId field value
// and a boolean to check if the value has been set.
func (o *CreateEntityWatchlistScreeningReviewRequest) GetEntityWatchlistScreeningIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistScreeningId, true
}

// SetEntityWatchlistScreeningId sets field value
func (o *CreateEntityWatchlistScreeningReviewRequest) SetEntityWatchlistScreeningId(v string) {
	o.EntityWatchlistScreeningId = v
}

func (o CreateEntityWatchlistScreeningReviewRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["confirmed_hits"] = o.ConfirmedHits
	}
	if true {
		toSerialize["dismissed_hits"] = o.DismissedHits
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["entity_watchlist_screening_id"] = o.EntityWatchlistScreeningId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateEntityWatchlistScreeningReviewRequest struct {
	value *CreateEntityWatchlistScreeningReviewRequest
	isSet bool
}

func (v NullableCreateEntityWatchlistScreeningReviewRequest) Get() *CreateEntityWatchlistScreeningReviewRequest {
	return v.value
}

func (v *NullableCreateEntityWatchlistScreeningReviewRequest) Set(val *CreateEntityWatchlistScreeningReviewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEntityWatchlistScreeningReviewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEntityWatchlistScreeningReviewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEntityWatchlistScreeningReviewRequest(val *CreateEntityWatchlistScreeningReviewRequest) *NullableCreateEntityWatchlistScreeningReviewRequest {
	return &NullableCreateEntityWatchlistScreeningReviewRequest{value: val, isSet: true}
}

func (v NullableCreateEntityWatchlistScreeningReviewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEntityWatchlistScreeningReviewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


