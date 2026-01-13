/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryMarginLoanRecordResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginLoanRecordResponseRowsInner{}

// QueryMarginLoanRecordResponseRowsInner struct for QueryMarginLoanRecordResponseRowsInner
type QueryMarginLoanRecordResponseRowsInner struct {
	TxId                 *int64  `json:"txId,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Principal            *string `json:"principal,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginLoanRecordResponseRowsInner QueryMarginLoanRecordResponseRowsInner

// NewQueryMarginLoanRecordResponseRowsInner instantiates a new QueryMarginLoanRecordResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginLoanRecordResponseRowsInner() *QueryMarginLoanRecordResponseRowsInner {
	this := QueryMarginLoanRecordResponseRowsInner{}
	return &this
}

// NewQueryMarginLoanRecordResponseRowsInnerWithDefaults instantiates a new QueryMarginLoanRecordResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginLoanRecordResponseRowsInnerWithDefaults() *QueryMarginLoanRecordResponseRowsInner {
	this := QueryMarginLoanRecordResponseRowsInner{}
	return &this
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *QueryMarginLoanRecordResponseRowsInner) GetTxId() int64 {
	if o == nil || common.IsNil(o.TxId) {
		var ret int64
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginLoanRecordResponseRowsInner) GetTxIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *QueryMarginLoanRecordResponseRowsInner) HasTxId() bool {
	if o != nil && !common.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int64 and assigns it to the TxId field.
func (o *QueryMarginLoanRecordResponseRowsInner) SetTxId(v int64) {
	o.TxId = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryMarginLoanRecordResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginLoanRecordResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryMarginLoanRecordResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryMarginLoanRecordResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *QueryMarginLoanRecordResponseRowsInner) GetPrincipal() string {
	if o == nil || common.IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginLoanRecordResponseRowsInner) GetPrincipalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *QueryMarginLoanRecordResponseRowsInner) HasPrincipal() bool {
	if o != nil && !common.IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *QueryMarginLoanRecordResponseRowsInner) SetPrincipal(v string) {
	o.Principal = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *QueryMarginLoanRecordResponseRowsInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginLoanRecordResponseRowsInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *QueryMarginLoanRecordResponseRowsInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *QueryMarginLoanRecordResponseRowsInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryMarginLoanRecordResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginLoanRecordResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryMarginLoanRecordResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryMarginLoanRecordResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o QueryMarginLoanRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginLoanRecordResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TxId) {
		toSerialize["txId"] = o.TxId
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMarginLoanRecordResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginLoanRecordResponseRowsInner := _QueryMarginLoanRecordResponseRowsInner{}

	err = json.Unmarshal(data, &varQueryMarginLoanRecordResponseRowsInner)

	if err != nil {
		return err
	}

	*o = QueryMarginLoanRecordResponseRowsInner(varQueryMarginLoanRecordResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "txId")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginLoanRecordResponseRowsInner struct {
	value *QueryMarginLoanRecordResponseRowsInner
	isSet bool
}

func (v NullableQueryMarginLoanRecordResponseRowsInner) Get() *QueryMarginLoanRecordResponseRowsInner {
	return v.value
}

func (v *NullableQueryMarginLoanRecordResponseRowsInner) Set(val *QueryMarginLoanRecordResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginLoanRecordResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginLoanRecordResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginLoanRecordResponseRowsInner(val *QueryMarginLoanRecordResponseRowsInner) *NullableQueryMarginLoanRecordResponseRowsInner {
	return &NullableQueryMarginLoanRecordResponseRowsInner{value: val, isSet: true}
}

func (v NullableQueryMarginLoanRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginLoanRecordResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
