/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryLimitOpenOrdersResponseListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryLimitOpenOrdersResponseListInner{}

// QueryLimitOpenOrdersResponseListInner struct for QueryLimitOpenOrdersResponseListInner
type QueryLimitOpenOrdersResponseListInner struct {
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
	ExpiredTimestamp     *int64  `json:"expiredTimestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryLimitOpenOrdersResponseListInner QueryLimitOpenOrdersResponseListInner

// NewQueryLimitOpenOrdersResponseListInner instantiates a new QueryLimitOpenOrdersResponseListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLimitOpenOrdersResponseListInner() *QueryLimitOpenOrdersResponseListInner {
	this := QueryLimitOpenOrdersResponseListInner{}
	return &this
}

// NewQueryLimitOpenOrdersResponseListInnerWithDefaults instantiates a new QueryLimitOpenOrdersResponseListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLimitOpenOrdersResponseListInnerWithDefaults() *QueryLimitOpenOrdersResponseListInner {
	this := QueryLimitOpenOrdersResponseListInner{}
	return &this
}

// GetQuoteId returns the QuoteId field value if set, zero value otherwise.
func (o *QueryLimitOpenOrdersResponseListInner) GetQuoteId() string {
	if o == nil || common.IsNil(o.QuoteId) {
		var ret string
		return ret
	}
	return *o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLimitOpenOrdersResponseListInner) GetQuoteIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteId) {
		return nil, false
	}
	return o.QuoteId, true
}

// HasQuoteId returns a boolean if a field has been set.
func (o *QueryLimitOpenOrdersResponseListInner) HasQuoteId() bool {
	if o != nil && !common.IsNil(o.QuoteId) {
		return true
	}

	return false
}

// SetQuoteId gets a reference to the given string and assigns it to the QuoteId field.
func (o *QueryLimitOpenOrdersResponseListInner) SetQuoteId(v string) {
	o.QuoteId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QueryLimitOpenOrdersResponseListInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLimitOpenOrdersResponseListInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryLimitOpenOrdersResponseListInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *QueryLimitOpenOrdersResponseListInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *QueryLimitOpenOrdersResponseListInner) GetOrderStatus() string {
	if o == nil || common.IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLimitOpenOrdersResponseListInner) GetOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *QueryLimitOpenOrdersResponseListInner) HasOrderStatus() bool {
	if o != nil && !common.IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *QueryLimitOpenOrdersResponseListInner) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetFromAsset returns the FromAsset field value if set, zero value otherwise.
func (o *QueryLimitOpenOrdersResponseListInner) GetFromAsset() string {
	if o == nil || common.IsNil(o.FromAsset) {
		var ret string
		return ret
	}
	return *o.FromAsset
}

// GetFromAssetOk returns a tuple with the FromAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLimitOpenOrdersResponseListInner) GetFromAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAsset) {
		return nil, false
	}
	return o.FromAsset, true
}

// HasFromAsset returns a boolean if a field has been set.
func (o *QueryLimitOpenOrdersResponseListInner) HasFromAsset() bool {
	if o != nil && !common.IsNil(o.FromAsset) {
		return true
	}

	return false
}

// SetFromAsset gets a reference to the given string and assigns it to the FromAsset field.
func (o *QueryLimitOpenOrdersResponseListInner) SetFromAsset(v string) {
	o.FromAsset = &v
}

// GetFromAmount returns the FromAmount field value if set, zero value otherwise.
func (o *QueryLimitOpenOrdersResponseListInner) GetFromAmount() string {
	if o == nil || common.IsNil(o.FromAmount) {
		var ret string
		return ret
	}
	return *o.FromAmount
}

// GetFromAmountOk returns a tuple with the FromAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLimitOpenOrdersResponseListInner) GetFromAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAmount) {
		return nil, false
	}
	return o.FromAmount, true
}

// HasFromAmount returns a boolean if a field has been set.
func (o *QueryLimitOpenOrdersResponseListInner) HasFromAmount() bool {
	if o != nil && !common.IsNil(o.FromAmount) {
		return true
	}

	return false
}

// SetFromAmount gets a reference to the given string and assigns it to the FromAmount field.
func (o *QueryLimitOpenOrdersResponseListInner) SetFromAmount(v string) {
	o.FromAmount = &v
}

// GetToAsset returns the ToAsset field value if set, zero value otherwise.
func (o *QueryLimitOpenOrdersResponseListInner) GetToAsset() string {
	if o == nil || common.IsNil(o.ToAsset) {
		var ret string
		return ret
	}
	return *o.ToAsset
}

// GetToAssetOk returns a tuple with the ToAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLimitOpenOrdersResponseListInner) GetToAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAsset) {
		return nil, false
	}
	return o.ToAsset, true
}

// HasToAsset returns a boolean if a field has been set.
func (o *QueryLimitOpenOrdersResponseListInner) HasToAsset() bool {
	if o != nil && !common.IsNil(o.ToAsset) {
		return true
	}

	return false
}

// SetToAsset gets a reference to the given string and assigns it to the ToAsset field.
func (o *QueryLimitOpenOrdersResponseListInner) SetToAsset(v string) {
	o.ToAsset = &v
}

// GetToAmount returns the ToAmount field value if set, zero value otherwise.
func (o *QueryLimitOpenOrdersResponseListInner) GetToAmount() string {
	if o == nil || common.IsNil(o.ToAmount) {
		var ret string
		return ret
	}
	return *o.ToAmount
}

// GetToAmountOk returns a tuple with the ToAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLimitOpenOrdersResponseListInner) GetToAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAmount) {
		return nil, false
	}
	return o.ToAmount, true
}

// HasToAmount returns a boolean if a field has been set.
func (o *QueryLimitOpenOrdersResponseListInner) HasToAmount() bool {
	if o != nil && !common.IsNil(o.ToAmount) {
		return true
	}

	return false
}

// SetToAmount gets a reference to the given string and assigns it to the ToAmount field.
func (o *QueryLimitOpenOrdersResponseListInner) SetToAmount(v string) {
	o.ToAmount = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *QueryLimitOpenOrdersResponseListInner) GetRatio() string {
	if o == nil || common.IsNil(o.Ratio) {
		var ret string
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLimitOpenOrdersResponseListInner) GetRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *QueryLimitOpenOrdersResponseListInner) HasRatio() bool {
	if o != nil && !common.IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given string and assigns it to the Ratio field.
func (o *QueryLimitOpenOrdersResponseListInner) SetRatio(v string) {
	o.Ratio = &v
}

// GetInverseRatio returns the InverseRatio field value if set, zero value otherwise.
func (o *QueryLimitOpenOrdersResponseListInner) GetInverseRatio() string {
	if o == nil || common.IsNil(o.InverseRatio) {
		var ret string
		return ret
	}
	return *o.InverseRatio
}

// GetInverseRatioOk returns a tuple with the InverseRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLimitOpenOrdersResponseListInner) GetInverseRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.InverseRatio) {
		return nil, false
	}
	return o.InverseRatio, true
}

// HasInverseRatio returns a boolean if a field has been set.
func (o *QueryLimitOpenOrdersResponseListInner) HasInverseRatio() bool {
	if o != nil && !common.IsNil(o.InverseRatio) {
		return true
	}

	return false
}

// SetInverseRatio gets a reference to the given string and assigns it to the InverseRatio field.
func (o *QueryLimitOpenOrdersResponseListInner) SetInverseRatio(v string) {
	o.InverseRatio = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *QueryLimitOpenOrdersResponseListInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLimitOpenOrdersResponseListInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *QueryLimitOpenOrdersResponseListInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *QueryLimitOpenOrdersResponseListInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetExpiredTimestamp returns the ExpiredTimestamp field value if set, zero value otherwise.
func (o *QueryLimitOpenOrdersResponseListInner) GetExpiredTimestamp() int64 {
	if o == nil || common.IsNil(o.ExpiredTimestamp) {
		var ret int64
		return ret
	}
	return *o.ExpiredTimestamp
}

// GetExpiredTimestampOk returns a tuple with the ExpiredTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLimitOpenOrdersResponseListInner) GetExpiredTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpiredTimestamp) {
		return nil, false
	}
	return o.ExpiredTimestamp, true
}

// HasExpiredTimestamp returns a boolean if a field has been set.
func (o *QueryLimitOpenOrdersResponseListInner) HasExpiredTimestamp() bool {
	if o != nil && !common.IsNil(o.ExpiredTimestamp) {
		return true
	}

	return false
}

// SetExpiredTimestamp gets a reference to the given int64 and assigns it to the ExpiredTimestamp field.
func (o *QueryLimitOpenOrdersResponseListInner) SetExpiredTimestamp(v int64) {
	o.ExpiredTimestamp = &v
}

func (o QueryLimitOpenOrdersResponseListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLimitOpenOrdersResponseListInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.ExpiredTimestamp) {
		toSerialize["expiredTimestamp"] = o.ExpiredTimestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryLimitOpenOrdersResponseListInner) UnmarshalJSON(data []byte) (err error) {
	varQueryLimitOpenOrdersResponseListInner := _QueryLimitOpenOrdersResponseListInner{}

	err = json.Unmarshal(data, &varQueryLimitOpenOrdersResponseListInner)

	if err != nil {
		return err
	}

	*o = QueryLimitOpenOrdersResponseListInner(varQueryLimitOpenOrdersResponseListInner)

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
		delete(additionalProperties, "expiredTimestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryLimitOpenOrdersResponseListInner struct {
	value *QueryLimitOpenOrdersResponseListInner
	isSet bool
}

func (v NullableQueryLimitOpenOrdersResponseListInner) Get() *QueryLimitOpenOrdersResponseListInner {
	return v.value
}

func (v *NullableQueryLimitOpenOrdersResponseListInner) Set(val *QueryLimitOpenOrdersResponseListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLimitOpenOrdersResponseListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLimitOpenOrdersResponseListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLimitOpenOrdersResponseListInner(val *QueryLimitOpenOrdersResponseListInner) *NullableQueryLimitOpenOrdersResponseListInner {
	return &NullableQueryLimitOpenOrdersResponseListInner{value: val, isSet: true}
}

func (v NullableQueryLimitOpenOrdersResponseListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLimitOpenOrdersResponseListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
