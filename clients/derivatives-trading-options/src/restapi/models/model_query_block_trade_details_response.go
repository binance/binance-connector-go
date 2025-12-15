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

// checks if the QueryBlockTradeDetailsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryBlockTradeDetailsResponse{}

// QueryBlockTradeDetailsResponse struct for QueryBlockTradeDetailsResponse
type QueryBlockTradeDetailsResponse struct {
	BlockTradeSettlementKey *string                                   `json:"blockTradeSettlementKey,omitempty"`
	ExpireTime              *int64                                    `json:"expireTime,omitempty"`
	Liquidity               *string                                   `json:"liquidity,omitempty"`
	Status                  *string                                   `json:"status,omitempty"`
	CreateTime              *int64                                    `json:"createTime,omitempty"`
	Legs                    []QueryBlockTradeDetailsResponseLegsInner `json:"legs,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _QueryBlockTradeDetailsResponse QueryBlockTradeDetailsResponse

// NewQueryBlockTradeDetailsResponse instantiates a new QueryBlockTradeDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryBlockTradeDetailsResponse() *QueryBlockTradeDetailsResponse {
	this := QueryBlockTradeDetailsResponse{}
	return &this
}

// NewQueryBlockTradeDetailsResponseWithDefaults instantiates a new QueryBlockTradeDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryBlockTradeDetailsResponseWithDefaults() *QueryBlockTradeDetailsResponse {
	this := QueryBlockTradeDetailsResponse{}
	return &this
}

// GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field value if set, zero value otherwise.
func (o *QueryBlockTradeDetailsResponse) GetBlockTradeSettlementKey() string {
	if o == nil || common.IsNil(o.BlockTradeSettlementKey) {
		var ret string
		return ret
	}
	return *o.BlockTradeSettlementKey
}

// GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBlockTradeDetailsResponse) GetBlockTradeSettlementKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.BlockTradeSettlementKey) {
		return nil, false
	}
	return o.BlockTradeSettlementKey, true
}

// HasBlockTradeSettlementKey returns a boolean if a field has been set.
func (o *QueryBlockTradeDetailsResponse) HasBlockTradeSettlementKey() bool {
	if o != nil && !common.IsNil(o.BlockTradeSettlementKey) {
		return true
	}

	return false
}

// SetBlockTradeSettlementKey gets a reference to the given string and assigns it to the BlockTradeSettlementKey field.
func (o *QueryBlockTradeDetailsResponse) SetBlockTradeSettlementKey(v string) {
	o.BlockTradeSettlementKey = &v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *QueryBlockTradeDetailsResponse) GetExpireTime() int64 {
	if o == nil || common.IsNil(o.ExpireTime) {
		var ret int64
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBlockTradeDetailsResponse) GetExpireTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *QueryBlockTradeDetailsResponse) HasExpireTime() bool {
	if o != nil && !common.IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given int64 and assigns it to the ExpireTime field.
func (o *QueryBlockTradeDetailsResponse) SetExpireTime(v int64) {
	o.ExpireTime = &v
}

// GetLiquidity returns the Liquidity field value if set, zero value otherwise.
func (o *QueryBlockTradeDetailsResponse) GetLiquidity() string {
	if o == nil || common.IsNil(o.Liquidity) {
		var ret string
		return ret
	}
	return *o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBlockTradeDetailsResponse) GetLiquidityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Liquidity) {
		return nil, false
	}
	return o.Liquidity, true
}

// HasLiquidity returns a boolean if a field has been set.
func (o *QueryBlockTradeDetailsResponse) HasLiquidity() bool {
	if o != nil && !common.IsNil(o.Liquidity) {
		return true
	}

	return false
}

// SetLiquidity gets a reference to the given string and assigns it to the Liquidity field.
func (o *QueryBlockTradeDetailsResponse) SetLiquidity(v string) {
	o.Liquidity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryBlockTradeDetailsResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBlockTradeDetailsResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryBlockTradeDetailsResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryBlockTradeDetailsResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *QueryBlockTradeDetailsResponse) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBlockTradeDetailsResponse) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *QueryBlockTradeDetailsResponse) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *QueryBlockTradeDetailsResponse) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetLegs returns the Legs field value if set, zero value otherwise.
func (o *QueryBlockTradeDetailsResponse) GetLegs() []QueryBlockTradeDetailsResponseLegsInner {
	if o == nil || common.IsNil(o.Legs) {
		var ret []QueryBlockTradeDetailsResponseLegsInner
		return ret
	}
	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBlockTradeDetailsResponse) GetLegsOk() ([]QueryBlockTradeDetailsResponseLegsInner, bool) {
	if o == nil || common.IsNil(o.Legs) {
		return nil, false
	}
	return o.Legs, true
}

// HasLegs returns a boolean if a field has been set.
func (o *QueryBlockTradeDetailsResponse) HasLegs() bool {
	if o != nil && !common.IsNil(o.Legs) {
		return true
	}

	return false
}

// SetLegs gets a reference to the given []QueryBlockTradeDetailsResponseLegsInner and assigns it to the Legs field.
func (o *QueryBlockTradeDetailsResponse) SetLegs(v []QueryBlockTradeDetailsResponseLegsInner) {
	o.Legs = v
}

func (o QueryBlockTradeDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryBlockTradeDetailsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryBlockTradeDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryBlockTradeDetailsResponse := _QueryBlockTradeDetailsResponse{}

	err = json.Unmarshal(data, &varQueryBlockTradeDetailsResponse)

	if err != nil {
		return err
	}

	*o = QueryBlockTradeDetailsResponse(varQueryBlockTradeDetailsResponse)

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

type NullableQueryBlockTradeDetailsResponse struct {
	value *QueryBlockTradeDetailsResponse
	isSet bool
}

func (v NullableQueryBlockTradeDetailsResponse) Get() *QueryBlockTradeDetailsResponse {
	return v.value
}

func (v *NullableQueryBlockTradeDetailsResponse) Set(val *QueryBlockTradeDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryBlockTradeDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryBlockTradeDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryBlockTradeDetailsResponse(val *QueryBlockTradeDetailsResponse) *NullableQueryBlockTradeDetailsResponse {
	return &NullableQueryBlockTradeDetailsResponse{value: val, isSet: true}
}

func (v NullableQueryBlockTradeDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryBlockTradeDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
