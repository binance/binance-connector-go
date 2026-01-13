/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetCollateralRecordResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCollateralRecordResponseRowsInner{}

// GetCollateralRecordResponseRowsInner struct for GetCollateralRecordResponseRowsInner
type GetCollateralRecordResponseRowsInner struct {
	Amount               *string `json:"amount,omitempty"`
	ProductId            *string `json:"productId,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty"`
	Type                 *string `json:"type,omitempty"`
	ProductName          *string `json:"productName,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCollateralRecordResponseRowsInner GetCollateralRecordResponseRowsInner

// NewGetCollateralRecordResponseRowsInner instantiates a new GetCollateralRecordResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCollateralRecordResponseRowsInner() *GetCollateralRecordResponseRowsInner {
	this := GetCollateralRecordResponseRowsInner{}
	return &this
}

// NewGetCollateralRecordResponseRowsInnerWithDefaults instantiates a new GetCollateralRecordResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCollateralRecordResponseRowsInnerWithDefaults() *GetCollateralRecordResponseRowsInner {
	this := GetCollateralRecordResponseRowsInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetCollateralRecordResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralRecordResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetCollateralRecordResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetCollateralRecordResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *GetCollateralRecordResponseRowsInner) GetProductId() string {
	if o == nil || common.IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralRecordResponseRowsInner) GetProductIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *GetCollateralRecordResponseRowsInner) HasProductId() bool {
	if o != nil && !common.IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *GetCollateralRecordResponseRowsInner) SetProductId(v string) {
	o.ProductId = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetCollateralRecordResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralRecordResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetCollateralRecordResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetCollateralRecordResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetCollateralRecordResponseRowsInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralRecordResponseRowsInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetCollateralRecordResponseRowsInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *GetCollateralRecordResponseRowsInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetCollateralRecordResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralRecordResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetCollateralRecordResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetCollateralRecordResponseRowsInner) SetType(v string) {
	o.Type = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *GetCollateralRecordResponseRowsInner) GetProductName() string {
	if o == nil || common.IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralRecordResponseRowsInner) GetProductNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *GetCollateralRecordResponseRowsInner) HasProductName() bool {
	if o != nil && !common.IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *GetCollateralRecordResponseRowsInner) SetProductName(v string) {
	o.ProductName = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetCollateralRecordResponseRowsInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralRecordResponseRowsInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetCollateralRecordResponseRowsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetCollateralRecordResponseRowsInner) SetOrderId(v int64) {
	o.OrderId = &v
}

func (o GetCollateralRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCollateralRecordResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCollateralRecordResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetCollateralRecordResponseRowsInner := _GetCollateralRecordResponseRowsInner{}

	err = json.Unmarshal(data, &varGetCollateralRecordResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetCollateralRecordResponseRowsInner(varGetCollateralRecordResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "productId")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "type")
		delete(additionalProperties, "productName")
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCollateralRecordResponseRowsInner struct {
	value *GetCollateralRecordResponseRowsInner
	isSet bool
}

func (v NullableGetCollateralRecordResponseRowsInner) Get() *GetCollateralRecordResponseRowsInner {
	return v.value
}

func (v *NullableGetCollateralRecordResponseRowsInner) Set(val *GetCollateralRecordResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCollateralRecordResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCollateralRecordResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCollateralRecordResponseRowsInner(val *GetCollateralRecordResponseRowsInner) *NullableGetCollateralRecordResponseRowsInner {
	return &NullableGetCollateralRecordResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetCollateralRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCollateralRecordResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
