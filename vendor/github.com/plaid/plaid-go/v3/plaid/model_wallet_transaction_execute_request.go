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

// WalletTransactionExecuteRequest WalletTransactionExecuteRequest defines the request schema for `/wallet/transaction/execute`
type WalletTransactionExecuteRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A random key provided by the client, per unique wallet transaction. Maximum of 128 characters.  The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. If a request to execute a wallet transaction fails due to a network connection error, then after a minimum delay of one minute, you can retry the request with the same idempotency key to guarantee that only a single wallet transaction is created. If the request was successfully processed, it will prevent any transaction that uses the same idempotency key, and was received within 24 hours of the first request, from being processed.
	IdempotencyKey string `json:"idempotency_key"`
	// The ID of the e-wallet to debit from
	WalletId string `json:"wallet_id"`
	Counterparty WalletTransactionCounterparty `json:"counterparty"`
	Amount WalletTransactionAmount `json:"amount"`
	// A reference for the transaction. This must be an alphanumeric string with at most 18 characters and must not contain any special characters or spaces.
	Reference string `json:"reference"`
}

// NewWalletTransactionExecuteRequest instantiates a new WalletTransactionExecuteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionExecuteRequest(idempotencyKey string, walletId string, counterparty WalletTransactionCounterparty, amount WalletTransactionAmount, reference string) *WalletTransactionExecuteRequest {
	this := WalletTransactionExecuteRequest{}
	this.IdempotencyKey = idempotencyKey
	this.WalletId = walletId
	this.Counterparty = counterparty
	this.Amount = amount
	this.Reference = reference
	return &this
}

// NewWalletTransactionExecuteRequestWithDefaults instantiates a new WalletTransactionExecuteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionExecuteRequestWithDefaults() *WalletTransactionExecuteRequest {
	this := WalletTransactionExecuteRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WalletTransactionExecuteRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionExecuteRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WalletTransactionExecuteRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WalletTransactionExecuteRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WalletTransactionExecuteRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionExecuteRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WalletTransactionExecuteRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WalletTransactionExecuteRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *WalletTransactionExecuteRequest) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionExecuteRequest) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *WalletTransactionExecuteRequest) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetWalletId returns the WalletId field value
func (o *WalletTransactionExecuteRequest) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionExecuteRequest) GetWalletIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *WalletTransactionExecuteRequest) SetWalletId(v string) {
	o.WalletId = v
}

// GetCounterparty returns the Counterparty field value
func (o *WalletTransactionExecuteRequest) GetCounterparty() WalletTransactionCounterparty {
	if o == nil {
		var ret WalletTransactionCounterparty
		return ret
	}

	return o.Counterparty
}

// GetCounterpartyOk returns a tuple with the Counterparty field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionExecuteRequest) GetCounterpartyOk() (*WalletTransactionCounterparty, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Counterparty, true
}

// SetCounterparty sets field value
func (o *WalletTransactionExecuteRequest) SetCounterparty(v WalletTransactionCounterparty) {
	o.Counterparty = v
}

// GetAmount returns the Amount field value
func (o *WalletTransactionExecuteRequest) GetAmount() WalletTransactionAmount {
	if o == nil {
		var ret WalletTransactionAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionExecuteRequest) GetAmountOk() (*WalletTransactionAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *WalletTransactionExecuteRequest) SetAmount(v WalletTransactionAmount) {
	o.Amount = v
}

// GetReference returns the Reference field value
func (o *WalletTransactionExecuteRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionExecuteRequest) GetReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *WalletTransactionExecuteRequest) SetReference(v string) {
	o.Reference = v
}

func (o WalletTransactionExecuteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	if true {
		toSerialize["wallet_id"] = o.WalletId
	}
	if true {
		toSerialize["counterparty"] = o.Counterparty
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["reference"] = o.Reference
	}
	return json.Marshal(toSerialize)
}

type NullableWalletTransactionExecuteRequest struct {
	value *WalletTransactionExecuteRequest
	isSet bool
}

func (v NullableWalletTransactionExecuteRequest) Get() *WalletTransactionExecuteRequest {
	return v.value
}

func (v *NullableWalletTransactionExecuteRequest) Set(val *WalletTransactionExecuteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionExecuteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionExecuteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionExecuteRequest(val *WalletTransactionExecuteRequest) *NullableWalletTransactionExecuteRequest {
	return &NullableWalletTransactionExecuteRequest{value: val, isSet: true}
}

func (v NullableWalletTransactionExecuteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionExecuteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


