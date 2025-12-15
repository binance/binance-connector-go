/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AcceptBlockTradeOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcceptBlockTradeOrderResponse{}

// AcceptBlockTradeOrderResponse struct for AcceptBlockTradeOrderResponse
type AcceptBlockTradeOrderResponse struct {
	BlockTradeSettlementKey *string                                  `json:"blockTradeSettlementKey,omitempty"`
	ExpireTime              *int64                                   `json:"expireTime,omitempty"`
	Liquidity               *string                                  `json:"liquidity,omitempty"`
	Status                  *string                                  `json:"status,omitempty"`
	CreateTime              *int64                                   `json:"createTime,omitempty"`
	Legs                    []AcceptBlockTradeOrderResponseLegsInner `json:"legs,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _AcceptBlockTradeOrderResponse AcceptBlockTradeOrderResponse

// NewAcceptBlockTradeOrderResponse instantiates a new AcceptBlockTradeOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptBlockTradeOrderResponse() *AcceptBlockTradeOrderResponse {
	this := AcceptBlockTradeOrderResponse{}
	return &this
}

// NewAcceptBlockTradeOrderResponseWithDefaults instantiates a new AcceptBlockTradeOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptBlockTradeOrderResponseWithDefaults() *AcceptBlockTradeOrderResponse {
	this := AcceptBlockTradeOrderResponse{}
	return &this
}

// GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field value if set, zero value otherwise.
func (o *AcceptBlockTradeOrderResponse) GetBlockTradeSettlementKey() string {
	if o == nil || common.IsNil(o.BlockTradeSettlementKey) {
		var ret string
		return ret
	}
	return *o.BlockTradeSettlementKey
}

// GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptBlockTradeOrderResponse) GetBlockTradeSettlementKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.BlockTradeSettlementKey) {
		return nil, false
	}
	return o.BlockTradeSettlementKey, true
}

// HasBlockTradeSettlementKey returns a boolean if a field has been set.
func (o *AcceptBlockTradeOrderResponse) HasBlockTradeSettlementKey() bool {
	if o != nil && !common.IsNil(o.BlockTradeSettlementKey) {
		return true
	}

	return false
}

// SetBlockTradeSettlementKey gets a reference to the given string and assigns it to the BlockTradeSettlementKey field.
func (o *AcceptBlockTradeOrderResponse) SetBlockTradeSettlementKey(v string) {
	o.BlockTradeSettlementKey = &v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *AcceptBlockTradeOrderResponse) GetExpireTime() int64 {
	if o == nil || common.IsNil(o.ExpireTime) {
		var ret int64
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptBlockTradeOrderResponse) GetExpireTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *AcceptBlockTradeOrderResponse) HasExpireTime() bool {
	if o != nil && !common.IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given int64 and assigns it to the ExpireTime field.
func (o *AcceptBlockTradeOrderResponse) SetExpireTime(v int64) {
	o.ExpireTime = &v
}

// GetLiquidity returns the Liquidity field value if set, zero value otherwise.
func (o *AcceptBlockTradeOrderResponse) GetLiquidity() string {
	if o == nil || common.IsNil(o.Liquidity) {
		var ret string
		return ret
	}
	return *o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptBlockTradeOrderResponse) GetLiquidityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Liquidity) {
		return nil, false
	}
	return o.Liquidity, true
}

// HasLiquidity returns a boolean if a field has been set.
func (o *AcceptBlockTradeOrderResponse) HasLiquidity() bool {
	if o != nil && !common.IsNil(o.Liquidity) {
		return true
	}

	return false
}

// SetLiquidity gets a reference to the given string and assigns it to the Liquidity field.
func (o *AcceptBlockTradeOrderResponse) SetLiquidity(v string) {
	o.Liquidity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AcceptBlockTradeOrderResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptBlockTradeOrderResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AcceptBlockTradeOrderResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AcceptBlockTradeOrderResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AcceptBlockTradeOrderResponse) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptBlockTradeOrderResponse) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AcceptBlockTradeOrderResponse) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *AcceptBlockTradeOrderResponse) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetLegs returns the Legs field value if set, zero value otherwise.
func (o *AcceptBlockTradeOrderResponse) GetLegs() []AcceptBlockTradeOrderResponseLegsInner {
	if o == nil || common.IsNil(o.Legs) {
		var ret []AcceptBlockTradeOrderResponseLegsInner
		return ret
	}
	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptBlockTradeOrderResponse) GetLegsOk() ([]AcceptBlockTradeOrderResponseLegsInner, bool) {
	if o == nil || common.IsNil(o.Legs) {
		return nil, false
	}
	return o.Legs, true
}

// HasLegs returns a boolean if a field has been set.
func (o *AcceptBlockTradeOrderResponse) HasLegs() bool {
	if o != nil && !common.IsNil(o.Legs) {
		return true
	}

	return false
}

// SetLegs gets a reference to the given []AcceptBlockTradeOrderResponseLegsInner and assigns it to the Legs field.
func (o *AcceptBlockTradeOrderResponse) SetLegs(v []AcceptBlockTradeOrderResponseLegsInner) {
	o.Legs = v
}

func (o AcceptBlockTradeOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptBlockTradeOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BlockTradeSettlementKey) {
		toSerialize["blockTradeSettlementKey"] = o.BlockTradeSettlementKey
	}
	if !common.IsNil(o.ExpireTime) {
		toSerialize["expireTime"] = o.ExpireTime
	}
	if !common.IsNil(o.Liquidity) {
		toSerialize["liquidity"] = o.Liquidity
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.Legs) {
		toSerialize["legs"] = o.Legs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcceptBlockTradeOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varAcceptBlockTradeOrderResponse := _AcceptBlockTradeOrderResponse{}

	err = json.Unmarshal(data, &varAcceptBlockTradeOrderResponse)

	if err != nil {
		return err
	}

	*o = AcceptBlockTradeOrderResponse(varAcceptBlockTradeOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "blockTradeSettlementKey")
		delete(additionalProperties, "expireTime")
		delete(additionalProperties, "liquidity")
		delete(additionalProperties, "status")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "legs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcceptBlockTradeOrderResponse struct {
	value *AcceptBlockTradeOrderResponse
	isSet bool
}

func (v NullableAcceptBlockTradeOrderResponse) Get() *AcceptBlockTradeOrderResponse {
	return v.value
}

func (v *NullableAcceptBlockTradeOrderResponse) Set(val *AcceptBlockTradeOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptBlockTradeOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptBlockTradeOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptBlockTradeOrderResponse(val *AcceptBlockTradeOrderResponse) *NullableAcceptBlockTradeOrderResponse {
	return &NullableAcceptBlockTradeOrderResponse{value: val, isSet: true}
}

func (v NullableAcceptBlockTradeOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptBlockTradeOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
