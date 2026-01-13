/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryBorrowRepayRecordsInMarginAccountResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryBorrowRepayRecordsInMarginAccountResponseRowsInner{}

// QueryBorrowRepayRecordsInMarginAccountResponseRowsInner struct for QueryBorrowRepayRecordsInMarginAccountResponseRowsInner
type QueryBorrowRepayRecordsInMarginAccountResponseRowsInner struct {
	Type                 *string `json:"type,omitempty"`
	IsolatedSymbol       *string `json:"isolatedSymbol,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Interest             *string `json:"interest,omitempty"`
	Principal            *string `json:"principal,omitempty"`
	Status               *string `json:"status,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	TxId                 *int64  `json:"txId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryBorrowRepayRecordsInMarginAccountResponseRowsInner QueryBorrowRepayRecordsInMarginAccountResponseRowsInner

// NewQueryBorrowRepayRecordsInMarginAccountResponseRowsInner instantiates a new QueryBorrowRepayRecordsInMarginAccountResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryBorrowRepayRecordsInMarginAccountResponseRowsInner() *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner {
	this := QueryBorrowRepayRecordsInMarginAccountResponseRowsInner{}
	return &this
}

// NewQueryBorrowRepayRecordsInMarginAccountResponseRowsInnerWithDefaults instantiates a new QueryBorrowRepayRecordsInMarginAccountResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryBorrowRepayRecordsInMarginAccountResponseRowsInnerWithDefaults() *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner {
	this := QueryBorrowRepayRecordsInMarginAccountResponseRowsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) SetType(v string) {
	o.Type = &v
}

// GetIsolatedSymbol returns the IsolatedSymbol field value if set, zero value otherwise.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetIsolatedSymbol() string {
	if o == nil || common.IsNil(o.IsolatedSymbol) {
		var ret string
		return ret
	}
	return *o.IsolatedSymbol
}

// GetIsolatedSymbolOk returns a tuple with the IsolatedSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetIsolatedSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsolatedSymbol) {
		return nil, false
	}
	return o.IsolatedSymbol, true
}

// HasIsolatedSymbol returns a boolean if a field has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) HasIsolatedSymbol() bool {
	if o != nil && !common.IsNil(o.IsolatedSymbol) {
		return true
	}

	return false
}

// SetIsolatedSymbol gets a reference to the given string and assigns it to the IsolatedSymbol field.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) SetIsolatedSymbol(v string) {
	o.IsolatedSymbol = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetInterest() string {
	if o == nil || common.IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) HasInterest() bool {
	if o != nil && !common.IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) SetInterest(v string) {
	o.Interest = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetPrincipal() string {
	if o == nil || common.IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetPrincipalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) HasPrincipal() bool {
	if o != nil && !common.IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) SetPrincipal(v string) {
	o.Principal = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetTxId() int64 {
	if o == nil || common.IsNil(o.TxId) {
		var ret int64
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) GetTxIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) HasTxId() bool {
	if o != nil && !common.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int64 and assigns it to the TxId field.
func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) SetTxId(v int64) {
	o.TxId = &v
}

func (o QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.IsolatedSymbol) {
		toSerialize["isolatedSymbol"] = o.IsolatedSymbol
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !common.IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !common.IsNil(o.TxId) {
		toSerialize["txId"] = o.TxId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryBorrowRepayRecordsInMarginAccountResponseRowsInner := _QueryBorrowRepayRecordsInMarginAccountResponseRowsInner{}

	err = json.Unmarshal(data, &varQueryBorrowRepayRecordsInMarginAccountResponseRowsInner)

	if err != nil {
		return err
	}

	*o = QueryBorrowRepayRecordsInMarginAccountResponseRowsInner(varQueryBorrowRepayRecordsInMarginAccountResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "isolatedSymbol")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "interest")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "txId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryBorrowRepayRecordsInMarginAccountResponseRowsInner struct {
	value *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner
	isSet bool
}

func (v NullableQueryBorrowRepayRecordsInMarginAccountResponseRowsInner) Get() *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner {
	return v.value
}

func (v *NullableQueryBorrowRepayRecordsInMarginAccountResponseRowsInner) Set(val *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryBorrowRepayRecordsInMarginAccountResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryBorrowRepayRecordsInMarginAccountResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryBorrowRepayRecordsInMarginAccountResponseRowsInner(val *QueryBorrowRepayRecordsInMarginAccountResponseRowsInner) *NullableQueryBorrowRepayRecordsInMarginAccountResponseRowsInner {
	return &NullableQueryBorrowRepayRecordsInMarginAccountResponseRowsInner{value: val, isSet: true}
}

func (v NullableQueryBorrowRepayRecordsInMarginAccountResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryBorrowRepayRecordsInMarginAccountResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
