/*
Binance Fiat REST API

OpenAPI Specification for the Binance Fiat REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetOrderDetailResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOrderDetailResponseData{}

// GetOrderDetailResponseData struct for GetOrderDetailResponseData
type GetOrderDetailResponseData struct {
	OrderId              *string                `json:"orderId,omitempty"`
	OrderStatus          *string                `json:"orderStatus,omitempty"`
	Amount               *string                `json:"amount,omitempty"`
	Fee                  *string                `json:"fee,omitempty"`
	FiatCurrency         *string                `json:"fiatCurrency,omitempty"`
	ErrorCode            *string                `json:"errorCode,omitempty"`
	ErrorMessage         *string                `json:"errorMessage,omitempty"`
	Ext                  map[string]interface{} `json:"ext,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOrderDetailResponseData GetOrderDetailResponseData

// NewGetOrderDetailResponseData instantiates a new GetOrderDetailResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderDetailResponseData() *GetOrderDetailResponseData {
	this := GetOrderDetailResponseData{}
	return &this
}

// NewGetOrderDetailResponseDataWithDefaults instantiates a new GetOrderDetailResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderDetailResponseDataWithDefaults() *GetOrderDetailResponseData {
	this := GetOrderDetailResponseData{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetOrderDetailResponseData) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderDetailResponseData) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetOrderDetailResponseData) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *GetOrderDetailResponseData) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *GetOrderDetailResponseData) GetOrderStatus() string {
	if o == nil || common.IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderDetailResponseData) GetOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *GetOrderDetailResponseData) HasOrderStatus() bool {
	if o != nil && !common.IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *GetOrderDetailResponseData) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetOrderDetailResponseData) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderDetailResponseData) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetOrderDetailResponseData) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetOrderDetailResponseData) SetAmount(v string) {
	o.Amount = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetOrderDetailResponseData) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderDetailResponseData) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetOrderDetailResponseData) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetOrderDetailResponseData) SetFee(v string) {
	o.Fee = &v
}

// GetFiatCurrency returns the FiatCurrency field value if set, zero value otherwise.
func (o *GetOrderDetailResponseData) GetFiatCurrency() string {
	if o == nil || common.IsNil(o.FiatCurrency) {
		var ret string
		return ret
	}
	return *o.FiatCurrency
}

// GetFiatCurrencyOk returns a tuple with the FiatCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderDetailResponseData) GetFiatCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.FiatCurrency) {
		return nil, false
	}
	return o.FiatCurrency, true
}

// HasFiatCurrency returns a boolean if a field has been set.
func (o *GetOrderDetailResponseData) HasFiatCurrency() bool {
	if o != nil && !common.IsNil(o.FiatCurrency) {
		return true
	}

	return false
}

// SetFiatCurrency gets a reference to the given string and assigns it to the FiatCurrency field.
func (o *GetOrderDetailResponseData) SetFiatCurrency(v string) {
	o.FiatCurrency = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *GetOrderDetailResponseData) GetErrorCode() string {
	if o == nil || common.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderDetailResponseData) GetErrorCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *GetOrderDetailResponseData) HasErrorCode() bool {
	if o != nil && !common.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *GetOrderDetailResponseData) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *GetOrderDetailResponseData) GetErrorMessage() string {
	if o == nil || common.IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderDetailResponseData) GetErrorMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *GetOrderDetailResponseData) HasErrorMessage() bool {
	if o != nil && !common.IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *GetOrderDetailResponseData) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetExt returns the Ext field value if set, zero value otherwise.
func (o *GetOrderDetailResponseData) GetExt() map[string]interface{} {
	if o == nil || common.IsNil(o.Ext) {
		var ret map[string]interface{}
		return ret
	}
	return o.Ext
}

// GetExtOk returns a tuple with the Ext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderDetailResponseData) GetExtOk() (map[string]interface{}, bool) {
	if o == nil || common.IsNil(o.Ext) {
		return map[string]interface{}{}, false
	}
	return o.Ext, true
}

// HasExt returns a boolean if a field has been set.
func (o *GetOrderDetailResponseData) HasExt() bool {
	if o != nil && !common.IsNil(o.Ext) {
		return true
	}

	return false
}

// SetExt gets a reference to the given map[string]interface{} and assigns it to the Ext field.
func (o *GetOrderDetailResponseData) SetExt(v map[string]interface{}) {
	o.Ext = v
}

func (o GetOrderDetailResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderDetailResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrderStatus) {
		toSerialize["orderStatus"] = o.OrderStatus
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !common.IsNil(o.FiatCurrency) {
		toSerialize["fiatCurrency"] = o.FiatCurrency
	}
	if !common.IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !common.IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !common.IsNil(o.Ext) {
		toSerialize["ext"] = o.Ext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOrderDetailResponseData) UnmarshalJSON(data []byte) (err error) {
	varGetOrderDetailResponseData := _GetOrderDetailResponseData{}

	err = json.Unmarshal(data, &varGetOrderDetailResponseData)

	if err != nil {
		return err
	}

	*o = GetOrderDetailResponseData(varGetOrderDetailResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderStatus")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "fiatCurrency")
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "ext")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOrderDetailResponseData struct {
	value *GetOrderDetailResponseData
	isSet bool
}

func (v NullableGetOrderDetailResponseData) Get() *GetOrderDetailResponseData {
	return v.value
}

func (v *NullableGetOrderDetailResponseData) Set(val *GetOrderDetailResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderDetailResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderDetailResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderDetailResponseData(val *GetOrderDetailResponseData) *NullableGetOrderDetailResponseData {
	return &NullableGetOrderDetailResponseData{value: val, isSet: true}
}

func (v NullableGetOrderDetailResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderDetailResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
