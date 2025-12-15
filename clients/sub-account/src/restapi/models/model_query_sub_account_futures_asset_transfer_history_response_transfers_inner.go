/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner{}

// QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner struct for QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner
type QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner struct {
	From                 *string `json:"from,omitempty"`
	To                   *string `json:"to,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	TranId               *int64  `json:"tranId,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner

// NewQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner instantiates a new QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner() *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner {
	this := QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner{}
	return &this
}

// NewQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInnerWithDefaults instantiates a new QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInnerWithDefaults() *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner {
	this := QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) GetFrom() string {
	if o == nil || common.IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) GetFromOk() (*string, bool) {
	if o == nil || common.IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) HasFrom() bool {
	if o != nil && !common.IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) GetTo() string {
	if o == nil || common.IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) GetToOk() (*string, bool) {
	if o == nil || common.IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) HasTo() bool {
	if o != nil && !common.IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) SetTo(v string) {
	o.To = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) SetAsset(v string) {
	o.Asset = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) SetQty(v string) {
	o.Qty = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) SetTranId(v int64) {
	o.TranId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) SetTime(v int64) {
	o.Time = &v
}

func (o QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !common.IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner := _QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner{}

	err = json.Unmarshal(data, &varQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner)

	if err != nil {
		return err
	}

	*o = QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner(varQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "tranId")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner struct {
	value *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner
	isSet bool
}

func (v NullableQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) Get() *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner {
	return v.value
}

func (v *NullableQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) Set(val *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner(val *QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) *NullableQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner {
	return &NullableQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner{value: val, isSet: true}
}

func (v NullableQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
