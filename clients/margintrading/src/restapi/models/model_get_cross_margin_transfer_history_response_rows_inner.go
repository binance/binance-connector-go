/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetCrossMarginTransferHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCrossMarginTransferHistoryResponseRowsInner{}

// GetCrossMarginTransferHistoryResponseRowsInner struct for GetCrossMarginTransferHistoryResponseRowsInner
type GetCrossMarginTransferHistoryResponseRowsInner struct {
	Amount               *string `json:"amount,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Status               *string `json:"status,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	TxId                 *int64  `json:"txId,omitempty"`
	Type                 *string `json:"type,omitempty"`
	TransFrom            *string `json:"transFrom,omitempty"`
	TransTo              *string `json:"transTo,omitempty"`
	FromSymbol           *string `json:"fromSymbol,omitempty"`
	ToSymbol             *string `json:"toSymbol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCrossMarginTransferHistoryResponseRowsInner GetCrossMarginTransferHistoryResponseRowsInner

// NewGetCrossMarginTransferHistoryResponseRowsInner instantiates a new GetCrossMarginTransferHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCrossMarginTransferHistoryResponseRowsInner() *GetCrossMarginTransferHistoryResponseRowsInner {
	this := GetCrossMarginTransferHistoryResponseRowsInner{}
	return &this
}

// NewGetCrossMarginTransferHistoryResponseRowsInnerWithDefaults instantiates a new GetCrossMarginTransferHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCrossMarginTransferHistoryResponseRowsInnerWithDefaults() *GetCrossMarginTransferHistoryResponseRowsInner {
	this := GetCrossMarginTransferHistoryResponseRowsInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetTxId() int64 {
	if o == nil || common.IsNil(o.TxId) {
		var ret int64
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetTxIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) HasTxId() bool {
	if o != nil && !common.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int64 and assigns it to the TxId field.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) SetTxId(v int64) {
	o.TxId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) SetType(v string) {
	o.Type = &v
}

// GetTransFrom returns the TransFrom field value if set, zero value otherwise.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetTransFrom() string {
	if o == nil || common.IsNil(o.TransFrom) {
		var ret string
		return ret
	}
	return *o.TransFrom
}

// GetTransFromOk returns a tuple with the TransFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetTransFromOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransFrom) {
		return nil, false
	}
	return o.TransFrom, true
}

// HasTransFrom returns a boolean if a field has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) HasTransFrom() bool {
	if o != nil && !common.IsNil(o.TransFrom) {
		return true
	}

	return false
}

// SetTransFrom gets a reference to the given string and assigns it to the TransFrom field.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) SetTransFrom(v string) {
	o.TransFrom = &v
}

// GetTransTo returns the TransTo field value if set, zero value otherwise.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetTransTo() string {
	if o == nil || common.IsNil(o.TransTo) {
		var ret string
		return ret
	}
	return *o.TransTo
}

// GetTransToOk returns a tuple with the TransTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetTransToOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransTo) {
		return nil, false
	}
	return o.TransTo, true
}

// HasTransTo returns a boolean if a field has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) HasTransTo() bool {
	if o != nil && !common.IsNil(o.TransTo) {
		return true
	}

	return false
}

// SetTransTo gets a reference to the given string and assigns it to the TransTo field.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) SetTransTo(v string) {
	o.TransTo = &v
}

// GetFromSymbol returns the FromSymbol field value if set, zero value otherwise.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetFromSymbol() string {
	if o == nil || common.IsNil(o.FromSymbol) {
		var ret string
		return ret
	}
	return *o.FromSymbol
}

// GetFromSymbolOk returns a tuple with the FromSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetFromSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromSymbol) {
		return nil, false
	}
	return o.FromSymbol, true
}

// HasFromSymbol returns a boolean if a field has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) HasFromSymbol() bool {
	if o != nil && !common.IsNil(o.FromSymbol) {
		return true
	}

	return false
}

// SetFromSymbol gets a reference to the given string and assigns it to the FromSymbol field.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) SetFromSymbol(v string) {
	o.FromSymbol = &v
}

// GetToSymbol returns the ToSymbol field value if set, zero value otherwise.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetToSymbol() string {
	if o == nil || common.IsNil(o.ToSymbol) {
		var ret string
		return ret
	}
	return *o.ToSymbol
}

// GetToSymbolOk returns a tuple with the ToSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) GetToSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToSymbol) {
		return nil, false
	}
	return o.ToSymbol, true
}

// HasToSymbol returns a boolean if a field has been set.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) HasToSymbol() bool {
	if o != nil && !common.IsNil(o.ToSymbol) {
		return true
	}

	return false
}

// SetToSymbol gets a reference to the given string and assigns it to the ToSymbol field.
func (o *GetCrossMarginTransferHistoryResponseRowsInner) SetToSymbol(v string) {
	o.ToSymbol = &v
}

func (o GetCrossMarginTransferHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCrossMarginTransferHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
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
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.TransFrom) {
		toSerialize["transFrom"] = o.TransFrom
	}
	if !common.IsNil(o.TransTo) {
		toSerialize["transTo"] = o.TransTo
	}
	if !common.IsNil(o.FromSymbol) {
		toSerialize["fromSymbol"] = o.FromSymbol
	}
	if !common.IsNil(o.ToSymbol) {
		toSerialize["toSymbol"] = o.ToSymbol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCrossMarginTransferHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetCrossMarginTransferHistoryResponseRowsInner := _GetCrossMarginTransferHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetCrossMarginTransferHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetCrossMarginTransferHistoryResponseRowsInner(varGetCrossMarginTransferHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "txId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "transFrom")
		delete(additionalProperties, "transTo")
		delete(additionalProperties, "fromSymbol")
		delete(additionalProperties, "toSymbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCrossMarginTransferHistoryResponseRowsInner struct {
	value *GetCrossMarginTransferHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetCrossMarginTransferHistoryResponseRowsInner) Get() *GetCrossMarginTransferHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetCrossMarginTransferHistoryResponseRowsInner) Set(val *GetCrossMarginTransferHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCrossMarginTransferHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCrossMarginTransferHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCrossMarginTransferHistoryResponseRowsInner(val *GetCrossMarginTransferHistoryResponseRowsInner) *NullableGetCrossMarginTransferHistoryResponseRowsInner {
	return &NullableGetCrossMarginTransferHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetCrossMarginTransferHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCrossMarginTransferHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
