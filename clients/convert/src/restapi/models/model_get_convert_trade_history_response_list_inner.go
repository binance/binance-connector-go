/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetConvertTradeHistoryResponseListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetConvertTradeHistoryResponseListInner{}

// GetConvertTradeHistoryResponseListInner struct for GetConvertTradeHistoryResponseListInner
type GetConvertTradeHistoryResponseListInner struct {
	QuoteId              *string `json:"quoteId,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	OrderStatus          *string `json:"orderStatus,omitempty"`
	FromAsset            *string `json:"fromAsset,omitempty"`
	FromAmount           *string `json:"fromAmount,omitempty"`
	ToAsset              *string `json:"toAsset,omitempty"`
	ToAmount             *string `json:"toAmount,omitempty"`
	Ratio                *string `json:"ratio,omitempty"`
	InverseRatio         *string `json:"inverseRatio,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetConvertTradeHistoryResponseListInner GetConvertTradeHistoryResponseListInner

// NewGetConvertTradeHistoryResponseListInner instantiates a new GetConvertTradeHistoryResponseListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConvertTradeHistoryResponseListInner() *GetConvertTradeHistoryResponseListInner {
	this := GetConvertTradeHistoryResponseListInner{}
	return &this
}

// NewGetConvertTradeHistoryResponseListInnerWithDefaults instantiates a new GetConvertTradeHistoryResponseListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConvertTradeHistoryResponseListInnerWithDefaults() *GetConvertTradeHistoryResponseListInner {
	this := GetConvertTradeHistoryResponseListInner{}
	return &this
}

// GetQuoteId returns the QuoteId field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponseListInner) GetQuoteId() string {
	if o == nil || common.IsNil(o.QuoteId) {
		var ret string
		return ret
	}
	return *o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponseListInner) GetQuoteIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteId) {
		return nil, false
	}
	return o.QuoteId, true
}

// HasQuoteId returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponseListInner) HasQuoteId() bool {
	if o != nil && !common.IsNil(o.QuoteId) {
		return true
	}

	return false
}

// SetQuoteId gets a reference to the given string and assigns it to the QuoteId field.
func (o *GetConvertTradeHistoryResponseListInner) SetQuoteId(v string) {
	o.QuoteId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponseListInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponseListInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponseListInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetConvertTradeHistoryResponseListInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponseListInner) GetOrderStatus() string {
	if o == nil || common.IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponseListInner) GetOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponseListInner) HasOrderStatus() bool {
	if o != nil && !common.IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *GetConvertTradeHistoryResponseListInner) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetFromAsset returns the FromAsset field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponseListInner) GetFromAsset() string {
	if o == nil || common.IsNil(o.FromAsset) {
		var ret string
		return ret
	}
	return *o.FromAsset
}

// GetFromAssetOk returns a tuple with the FromAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponseListInner) GetFromAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAsset) {
		return nil, false
	}
	return o.FromAsset, true
}

// HasFromAsset returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponseListInner) HasFromAsset() bool {
	if o != nil && !common.IsNil(o.FromAsset) {
		return true
	}

	return false
}

// SetFromAsset gets a reference to the given string and assigns it to the FromAsset field.
func (o *GetConvertTradeHistoryResponseListInner) SetFromAsset(v string) {
	o.FromAsset = &v
}

// GetFromAmount returns the FromAmount field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponseListInner) GetFromAmount() string {
	if o == nil || common.IsNil(o.FromAmount) {
		var ret string
		return ret
	}
	return *o.FromAmount
}

// GetFromAmountOk returns a tuple with the FromAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponseListInner) GetFromAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAmount) {
		return nil, false
	}
	return o.FromAmount, true
}

// HasFromAmount returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponseListInner) HasFromAmount() bool {
	if o != nil && !common.IsNil(o.FromAmount) {
		return true
	}

	return false
}

// SetFromAmount gets a reference to the given string and assigns it to the FromAmount field.
func (o *GetConvertTradeHistoryResponseListInner) SetFromAmount(v string) {
	o.FromAmount = &v
}

// GetToAsset returns the ToAsset field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponseListInner) GetToAsset() string {
	if o == nil || common.IsNil(o.ToAsset) {
		var ret string
		return ret
	}
	return *o.ToAsset
}

// GetToAssetOk returns a tuple with the ToAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponseListInner) GetToAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAsset) {
		return nil, false
	}
	return o.ToAsset, true
}

// HasToAsset returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponseListInner) HasToAsset() bool {
	if o != nil && !common.IsNil(o.ToAsset) {
		return true
	}

	return false
}

// SetToAsset gets a reference to the given string and assigns it to the ToAsset field.
func (o *GetConvertTradeHistoryResponseListInner) SetToAsset(v string) {
	o.ToAsset = &v
}

// GetToAmount returns the ToAmount field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponseListInner) GetToAmount() string {
	if o == nil || common.IsNil(o.ToAmount) {
		var ret string
		return ret
	}
	return *o.ToAmount
}

// GetToAmountOk returns a tuple with the ToAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponseListInner) GetToAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAmount) {
		return nil, false
	}
	return o.ToAmount, true
}

// HasToAmount returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponseListInner) HasToAmount() bool {
	if o != nil && !common.IsNil(o.ToAmount) {
		return true
	}

	return false
}

// SetToAmount gets a reference to the given string and assigns it to the ToAmount field.
func (o *GetConvertTradeHistoryResponseListInner) SetToAmount(v string) {
	o.ToAmount = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponseListInner) GetRatio() string {
	if o == nil || common.IsNil(o.Ratio) {
		var ret string
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponseListInner) GetRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponseListInner) HasRatio() bool {
	if o != nil && !common.IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given string and assigns it to the Ratio field.
func (o *GetConvertTradeHistoryResponseListInner) SetRatio(v string) {
	o.Ratio = &v
}

// GetInverseRatio returns the InverseRatio field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponseListInner) GetInverseRatio() string {
	if o == nil || common.IsNil(o.InverseRatio) {
		var ret string
		return ret
	}
	return *o.InverseRatio
}

// GetInverseRatioOk returns a tuple with the InverseRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponseListInner) GetInverseRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.InverseRatio) {
		return nil, false
	}
	return o.InverseRatio, true
}

// HasInverseRatio returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponseListInner) HasInverseRatio() bool {
	if o != nil && !common.IsNil(o.InverseRatio) {
		return true
	}

	return false
}

// SetInverseRatio gets a reference to the given string and assigns it to the InverseRatio field.
func (o *GetConvertTradeHistoryResponseListInner) SetInverseRatio(v string) {
	o.InverseRatio = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponseListInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponseListInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponseListInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *GetConvertTradeHistoryResponseListInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

func (o GetConvertTradeHistoryResponseListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConvertTradeHistoryResponseListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.QuoteId) {
		toSerialize["quoteId"] = o.QuoteId
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrderStatus) {
		toSerialize["orderStatus"] = o.OrderStatus
	}
	if !common.IsNil(o.FromAsset) {
		toSerialize["fromAsset"] = o.FromAsset
	}
	if !common.IsNil(o.FromAmount) {
		toSerialize["fromAmount"] = o.FromAmount
	}
	if !common.IsNil(o.ToAsset) {
		toSerialize["toAsset"] = o.ToAsset
	}
	if !common.IsNil(o.ToAmount) {
		toSerialize["toAmount"] = o.ToAmount
	}
	if !common.IsNil(o.Ratio) {
		toSerialize["ratio"] = o.Ratio
	}
	if !common.IsNil(o.InverseRatio) {
		toSerialize["inverseRatio"] = o.InverseRatio
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetConvertTradeHistoryResponseListInner) UnmarshalJSON(data []byte) (err error) {
	varGetConvertTradeHistoryResponseListInner := _GetConvertTradeHistoryResponseListInner{}

	err = json.Unmarshal(data, &varGetConvertTradeHistoryResponseListInner)

	if err != nil {
		return err
	}

	*o = GetConvertTradeHistoryResponseListInner(varGetConvertTradeHistoryResponseListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "quoteId")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderStatus")
		delete(additionalProperties, "fromAsset")
		delete(additionalProperties, "fromAmount")
		delete(additionalProperties, "toAsset")
		delete(additionalProperties, "toAmount")
		delete(additionalProperties, "ratio")
		delete(additionalProperties, "inverseRatio")
		delete(additionalProperties, "createTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetConvertTradeHistoryResponseListInner struct {
	value *GetConvertTradeHistoryResponseListInner
	isSet bool
}

func (v NullableGetConvertTradeHistoryResponseListInner) Get() *GetConvertTradeHistoryResponseListInner {
	return v.value
}

func (v *NullableGetConvertTradeHistoryResponseListInner) Set(val *GetConvertTradeHistoryResponseListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConvertTradeHistoryResponseListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConvertTradeHistoryResponseListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConvertTradeHistoryResponseListInner(val *GetConvertTradeHistoryResponseListInner) *NullableGetConvertTradeHistoryResponseListInner {
	return &NullableGetConvertTradeHistoryResponseListInner{value: val, isSet: true}
}

func (v NullableGetConvertTradeHistoryResponseListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConvertTradeHistoryResponseListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
