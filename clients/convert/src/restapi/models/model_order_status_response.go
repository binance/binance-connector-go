/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderStatusResponse{}

// OrderStatusResponse struct for OrderStatusResponse
type OrderStatusResponse struct {
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

type _OrderStatusResponse OrderStatusResponse

// NewOrderStatusResponse instantiates a new OrderStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderStatusResponse() *OrderStatusResponse {
	this := OrderStatusResponse{}
	return &this
}

// NewOrderStatusResponseWithDefaults instantiates a new OrderStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderStatusResponseWithDefaults() *OrderStatusResponse {
	this := OrderStatusResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *OrderStatusResponse) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetOrderStatus() string {
	if o == nil || common.IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasOrderStatus() bool {
	if o != nil && !common.IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *OrderStatusResponse) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetFromAsset returns the FromAsset field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetFromAsset() string {
	if o == nil || common.IsNil(o.FromAsset) {
		var ret string
		return ret
	}
	return *o.FromAsset
}

// GetFromAssetOk returns a tuple with the FromAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetFromAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAsset) {
		return nil, false
	}
	return o.FromAsset, true
}

// HasFromAsset returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasFromAsset() bool {
	if o != nil && !common.IsNil(o.FromAsset) {
		return true
	}

	return false
}

// SetFromAsset gets a reference to the given string and assigns it to the FromAsset field.
func (o *OrderStatusResponse) SetFromAsset(v string) {
	o.FromAsset = &v
}

// GetFromAmount returns the FromAmount field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetFromAmount() string {
	if o == nil || common.IsNil(o.FromAmount) {
		var ret string
		return ret
	}
	return *o.FromAmount
}

// GetFromAmountOk returns a tuple with the FromAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetFromAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAmount) {
		return nil, false
	}
	return o.FromAmount, true
}

// HasFromAmount returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasFromAmount() bool {
	if o != nil && !common.IsNil(o.FromAmount) {
		return true
	}

	return false
}

// SetFromAmount gets a reference to the given string and assigns it to the FromAmount field.
func (o *OrderStatusResponse) SetFromAmount(v string) {
	o.FromAmount = &v
}

// GetToAsset returns the ToAsset field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetToAsset() string {
	if o == nil || common.IsNil(o.ToAsset) {
		var ret string
		return ret
	}
	return *o.ToAsset
}

// GetToAssetOk returns a tuple with the ToAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetToAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAsset) {
		return nil, false
	}
	return o.ToAsset, true
}

// HasToAsset returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasToAsset() bool {
	if o != nil && !common.IsNil(o.ToAsset) {
		return true
	}

	return false
}

// SetToAsset gets a reference to the given string and assigns it to the ToAsset field.
func (o *OrderStatusResponse) SetToAsset(v string) {
	o.ToAsset = &v
}

// GetToAmount returns the ToAmount field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetToAmount() string {
	if o == nil || common.IsNil(o.ToAmount) {
		var ret string
		return ret
	}
	return *o.ToAmount
}

// GetToAmountOk returns a tuple with the ToAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetToAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAmount) {
		return nil, false
	}
	return o.ToAmount, true
}

// HasToAmount returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasToAmount() bool {
	if o != nil && !common.IsNil(o.ToAmount) {
		return true
	}

	return false
}

// SetToAmount gets a reference to the given string and assigns it to the ToAmount field.
func (o *OrderStatusResponse) SetToAmount(v string) {
	o.ToAmount = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetRatio() string {
	if o == nil || common.IsNil(o.Ratio) {
		var ret string
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasRatio() bool {
	if o != nil && !common.IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given string and assigns it to the Ratio field.
func (o *OrderStatusResponse) SetRatio(v string) {
	o.Ratio = &v
}

// GetInverseRatio returns the InverseRatio field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetInverseRatio() string {
	if o == nil || common.IsNil(o.InverseRatio) {
		var ret string
		return ret
	}
	return *o.InverseRatio
}

// GetInverseRatioOk returns a tuple with the InverseRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetInverseRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.InverseRatio) {
		return nil, false
	}
	return o.InverseRatio, true
}

// HasInverseRatio returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasInverseRatio() bool {
	if o != nil && !common.IsNil(o.InverseRatio) {
		return true
	}

	return false
}

// SetInverseRatio gets a reference to the given string and assigns it to the InverseRatio field.
func (o *OrderStatusResponse) SetInverseRatio(v string) {
	o.InverseRatio = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *OrderStatusResponse) SetCreateTime(v int64) {
	o.CreateTime = &v
}

func (o OrderStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *OrderStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderStatusResponse := _OrderStatusResponse{}

	err = json.Unmarshal(data, &varOrderStatusResponse)

	if err != nil {
		return err
	}

	*o = OrderStatusResponse(varOrderStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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

type NullableOrderStatusResponse struct {
	value *OrderStatusResponse
	isSet bool
}

func (v NullableOrderStatusResponse) Get() *OrderStatusResponse {
	return v.value
}

func (v *NullableOrderStatusResponse) Set(val *OrderStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatusResponse(val *OrderStatusResponse) *NullableOrderStatusResponse {
	return &NullableOrderStatusResponse{value: val, isSet: true}
}

func (v NullableOrderStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
