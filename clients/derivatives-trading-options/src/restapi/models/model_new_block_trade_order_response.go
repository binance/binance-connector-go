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

// checks if the NewBlockTradeOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NewBlockTradeOrderResponse{}

// NewBlockTradeOrderResponse struct for NewBlockTradeOrderResponse
type NewBlockTradeOrderResponse struct {
	BlockTradeSettlementKey *string                                  `json:"blockTradeSettlementKey,omitempty"`
	ExpireTime              *int64                                   `json:"expireTime,omitempty"`
	Liquidity               *string                                  `json:"liquidity,omitempty"`
	Status                  *string                                  `json:"status,omitempty"`
	Legs                    []ExtendBlockTradeOrderResponseLegsInner `json:"legs,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _NewBlockTradeOrderResponse NewBlockTradeOrderResponse

// NewNewBlockTradeOrderResponse instantiates a new NewBlockTradeOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewBlockTradeOrderResponse() *NewBlockTradeOrderResponse {
	this := NewBlockTradeOrderResponse{}
	return &this
}

// NewNewBlockTradeOrderResponseWithDefaults instantiates a new NewBlockTradeOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewBlockTradeOrderResponseWithDefaults() *NewBlockTradeOrderResponse {
	this := NewBlockTradeOrderResponse{}
	return &this
}

// GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field value if set, zero value otherwise.
func (o *NewBlockTradeOrderResponse) GetBlockTradeSettlementKey() string {
	if o == nil || common.IsNil(o.BlockTradeSettlementKey) {
		var ret string
		return ret
	}
	return *o.BlockTradeSettlementKey
}

// GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewBlockTradeOrderResponse) GetBlockTradeSettlementKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.BlockTradeSettlementKey) {
		return nil, false
	}
	return o.BlockTradeSettlementKey, true
}

// HasBlockTradeSettlementKey returns a boolean if a field has been set.
func (o *NewBlockTradeOrderResponse) HasBlockTradeSettlementKey() bool {
	if o != nil && !common.IsNil(o.BlockTradeSettlementKey) {
		return true
	}

	return false
}

// SetBlockTradeSettlementKey gets a reference to the given string and assigns it to the BlockTradeSettlementKey field.
func (o *NewBlockTradeOrderResponse) SetBlockTradeSettlementKey(v string) {
	o.BlockTradeSettlementKey = &v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *NewBlockTradeOrderResponse) GetExpireTime() int64 {
	if o == nil || common.IsNil(o.ExpireTime) {
		var ret int64
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewBlockTradeOrderResponse) GetExpireTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *NewBlockTradeOrderResponse) HasExpireTime() bool {
	if o != nil && !common.IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given int64 and assigns it to the ExpireTime field.
func (o *NewBlockTradeOrderResponse) SetExpireTime(v int64) {
	o.ExpireTime = &v
}

// GetLiquidity returns the Liquidity field value if set, zero value otherwise.
func (o *NewBlockTradeOrderResponse) GetLiquidity() string {
	if o == nil || common.IsNil(o.Liquidity) {
		var ret string
		return ret
	}
	return *o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewBlockTradeOrderResponse) GetLiquidityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Liquidity) {
		return nil, false
	}
	return o.Liquidity, true
}

// HasLiquidity returns a boolean if a field has been set.
func (o *NewBlockTradeOrderResponse) HasLiquidity() bool {
	if o != nil && !common.IsNil(o.Liquidity) {
		return true
	}

	return false
}

// SetLiquidity gets a reference to the given string and assigns it to the Liquidity field.
func (o *NewBlockTradeOrderResponse) SetLiquidity(v string) {
	o.Liquidity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NewBlockTradeOrderResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewBlockTradeOrderResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NewBlockTradeOrderResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NewBlockTradeOrderResponse) SetStatus(v string) {
	o.Status = &v
}

// GetLegs returns the Legs field value if set, zero value otherwise.
func (o *NewBlockTradeOrderResponse) GetLegs() []ExtendBlockTradeOrderResponseLegsInner {
	if o == nil || common.IsNil(o.Legs) {
		var ret []ExtendBlockTradeOrderResponseLegsInner
		return ret
	}
	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewBlockTradeOrderResponse) GetLegsOk() ([]ExtendBlockTradeOrderResponseLegsInner, bool) {
	if o == nil || common.IsNil(o.Legs) {
		return nil, false
	}
	return o.Legs, true
}

// HasLegs returns a boolean if a field has been set.
func (o *NewBlockTradeOrderResponse) HasLegs() bool {
	if o != nil && !common.IsNil(o.Legs) {
		return true
	}

	return false
}

// SetLegs gets a reference to the given []ExtendBlockTradeOrderResponseLegsInner and assigns it to the Legs field.
func (o *NewBlockTradeOrderResponse) SetLegs(v []ExtendBlockTradeOrderResponseLegsInner) {
	o.Legs = v
}

func (o NewBlockTradeOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewBlockTradeOrderResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Legs) {
		toSerialize["legs"] = o.Legs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NewBlockTradeOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varNewBlockTradeOrderResponse := _NewBlockTradeOrderResponse{}

	err = json.Unmarshal(data, &varNewBlockTradeOrderResponse)

	if err != nil {
		return err
	}

	*o = NewBlockTradeOrderResponse(varNewBlockTradeOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "blockTradeSettlementKey")
		delete(additionalProperties, "expireTime")
		delete(additionalProperties, "liquidity")
		delete(additionalProperties, "status")
		delete(additionalProperties, "legs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNewBlockTradeOrderResponse struct {
	value *NewBlockTradeOrderResponse
	isSet bool
}

func (v NullableNewBlockTradeOrderResponse) Get() *NewBlockTradeOrderResponse {
	return v.value
}

func (v *NullableNewBlockTradeOrderResponse) Set(val *NewBlockTradeOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNewBlockTradeOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNewBlockTradeOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewBlockTradeOrderResponse(val *NewBlockTradeOrderResponse) *NullableNewBlockTradeOrderResponse {
	return &NullableNewBlockTradeOrderResponse{value: val, isSet: true}
}

func (v NullableNewBlockTradeOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewBlockTradeOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
