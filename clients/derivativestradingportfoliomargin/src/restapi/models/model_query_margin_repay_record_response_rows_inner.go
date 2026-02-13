/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryMarginRepayRecordResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginRepayRecordResponseRowsInner{}

// QueryMarginRepayRecordResponseRowsInner struct for QueryMarginRepayRecordResponseRowsInner
type QueryMarginRepayRecordResponseRowsInner struct {
	Amount               *string `json:"amount,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Interest             *string `json:"interest,omitempty"`
	Principal            *string `json:"principal,omitempty"`
	Status               *string `json:"status,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	TxId                 *int64  `json:"txId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginRepayRecordResponseRowsInner QueryMarginRepayRecordResponseRowsInner

// NewQueryMarginRepayRecordResponseRowsInner instantiates a new QueryMarginRepayRecordResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginRepayRecordResponseRowsInner() *QueryMarginRepayRecordResponseRowsInner {
	this := QueryMarginRepayRecordResponseRowsInner{}
	return &this
}

// NewQueryMarginRepayRecordResponseRowsInnerWithDefaults instantiates a new QueryMarginRepayRecordResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginRepayRecordResponseRowsInnerWithDefaults() *QueryMarginRepayRecordResponseRowsInner {
	this := QueryMarginRepayRecordResponseRowsInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryMarginRepayRecordResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryMarginRepayRecordResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryMarginRepayRecordResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryMarginRepayRecordResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *QueryMarginRepayRecordResponseRowsInner) GetInterest() string {
	if o == nil || common.IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) GetInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) HasInterest() bool {
	if o != nil && !common.IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *QueryMarginRepayRecordResponseRowsInner) SetInterest(v string) {
	o.Interest = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *QueryMarginRepayRecordResponseRowsInner) GetPrincipal() string {
	if o == nil || common.IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) GetPrincipalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) HasPrincipal() bool {
	if o != nil && !common.IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *QueryMarginRepayRecordResponseRowsInner) SetPrincipal(v string) {
	o.Principal = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryMarginRepayRecordResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryMarginRepayRecordResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *QueryMarginRepayRecordResponseRowsInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *QueryMarginRepayRecordResponseRowsInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *QueryMarginRepayRecordResponseRowsInner) GetTxId() int64 {
	if o == nil || common.IsNil(o.TxId) {
		var ret int64
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) GetTxIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *QueryMarginRepayRecordResponseRowsInner) HasTxId() bool {
	if o != nil && !common.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int64 and assigns it to the TxId field.
func (o *QueryMarginRepayRecordResponseRowsInner) SetTxId(v int64) {
	o.TxId = &v
}

func (o QueryMarginRepayRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginRepayRecordResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *QueryMarginRepayRecordResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginRepayRecordResponseRowsInner := _QueryMarginRepayRecordResponseRowsInner{}

	err = json.Unmarshal(data, &varQueryMarginRepayRecordResponseRowsInner)

	if err != nil {
		return err
	}

	*o = QueryMarginRepayRecordResponseRowsInner(varQueryMarginRepayRecordResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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

type NullableQueryMarginRepayRecordResponseRowsInner struct {
	value *QueryMarginRepayRecordResponseRowsInner
	isSet bool
}

func (v NullableQueryMarginRepayRecordResponseRowsInner) Get() *QueryMarginRepayRecordResponseRowsInner {
	return v.value
}

func (v *NullableQueryMarginRepayRecordResponseRowsInner) Set(val *QueryMarginRepayRecordResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginRepayRecordResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginRepayRecordResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginRepayRecordResponseRowsInner(val *QueryMarginRepayRecordResponseRowsInner) *NullableQueryMarginRepayRecordResponseRowsInner {
	return &NullableQueryMarginRepayRecordResponseRowsInner{value: val, isSet: true}
}

func (v NullableQueryMarginRepayRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginRepayRecordResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
