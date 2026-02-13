/*
Binance Fiat REST API

OpenAPI Specification for the Binance Fiat REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFiatPaymentsHistoryResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFiatPaymentsHistoryResponseDataInner{}

// GetFiatPaymentsHistoryResponseDataInner struct for GetFiatPaymentsHistoryResponseDataInner
type GetFiatPaymentsHistoryResponseDataInner struct {
	OrderNo              *string `json:"orderNo,omitempty"`
	SourceAmount         *string `json:"sourceAmount,omitempty"`
	FiatCurrency         *string `json:"fiatCurrency,omitempty"`
	ObtainAmount         *string `json:"obtainAmount,omitempty"`
	CryptoCurrency       *string `json:"cryptoCurrency,omitempty"`
	TotalFee             *string `json:"totalFee,omitempty"`
	Price                *string `json:"price,omitempty"`
	Status               *string `json:"status,omitempty"`
	PaymentMethod        *string `json:"paymentMethod,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFiatPaymentsHistoryResponseDataInner GetFiatPaymentsHistoryResponseDataInner

// NewGetFiatPaymentsHistoryResponseDataInner instantiates a new GetFiatPaymentsHistoryResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFiatPaymentsHistoryResponseDataInner() *GetFiatPaymentsHistoryResponseDataInner {
	this := GetFiatPaymentsHistoryResponseDataInner{}
	return &this
}

// NewGetFiatPaymentsHistoryResponseDataInnerWithDefaults instantiates a new GetFiatPaymentsHistoryResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFiatPaymentsHistoryResponseDataInnerWithDefaults() *GetFiatPaymentsHistoryResponseDataInner {
	this := GetFiatPaymentsHistoryResponseDataInner{}
	return &this
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetOrderNo() string {
	if o == nil || common.IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetOrderNoOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) HasOrderNo() bool {
	if o != nil && !common.IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *GetFiatPaymentsHistoryResponseDataInner) SetOrderNo(v string) {
	o.OrderNo = &v
}

// GetSourceAmount returns the SourceAmount field value if set, zero value otherwise.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetSourceAmount() string {
	if o == nil || common.IsNil(o.SourceAmount) {
		var ret string
		return ret
	}
	return *o.SourceAmount
}

// GetSourceAmountOk returns a tuple with the SourceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetSourceAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.SourceAmount) {
		return nil, false
	}
	return o.SourceAmount, true
}

// HasSourceAmount returns a boolean if a field has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) HasSourceAmount() bool {
	if o != nil && !common.IsNil(o.SourceAmount) {
		return true
	}

	return false
}

// SetSourceAmount gets a reference to the given string and assigns it to the SourceAmount field.
func (o *GetFiatPaymentsHistoryResponseDataInner) SetSourceAmount(v string) {
	o.SourceAmount = &v
}

// GetFiatCurrency returns the FiatCurrency field value if set, zero value otherwise.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetFiatCurrency() string {
	if o == nil || common.IsNil(o.FiatCurrency) {
		var ret string
		return ret
	}
	return *o.FiatCurrency
}

// GetFiatCurrencyOk returns a tuple with the FiatCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetFiatCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.FiatCurrency) {
		return nil, false
	}
	return o.FiatCurrency, true
}

// HasFiatCurrency returns a boolean if a field has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) HasFiatCurrency() bool {
	if o != nil && !common.IsNil(o.FiatCurrency) {
		return true
	}

	return false
}

// SetFiatCurrency gets a reference to the given string and assigns it to the FiatCurrency field.
func (o *GetFiatPaymentsHistoryResponseDataInner) SetFiatCurrency(v string) {
	o.FiatCurrency = &v
}

// GetObtainAmount returns the ObtainAmount field value if set, zero value otherwise.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetObtainAmount() string {
	if o == nil || common.IsNil(o.ObtainAmount) {
		var ret string
		return ret
	}
	return *o.ObtainAmount
}

// GetObtainAmountOk returns a tuple with the ObtainAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetObtainAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ObtainAmount) {
		return nil, false
	}
	return o.ObtainAmount, true
}

// HasObtainAmount returns a boolean if a field has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) HasObtainAmount() bool {
	if o != nil && !common.IsNil(o.ObtainAmount) {
		return true
	}

	return false
}

// SetObtainAmount gets a reference to the given string and assigns it to the ObtainAmount field.
func (o *GetFiatPaymentsHistoryResponseDataInner) SetObtainAmount(v string) {
	o.ObtainAmount = &v
}

// GetCryptoCurrency returns the CryptoCurrency field value if set, zero value otherwise.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetCryptoCurrency() string {
	if o == nil || common.IsNil(o.CryptoCurrency) {
		var ret string
		return ret
	}
	return *o.CryptoCurrency
}

// GetCryptoCurrencyOk returns a tuple with the CryptoCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetCryptoCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CryptoCurrency) {
		return nil, false
	}
	return o.CryptoCurrency, true
}

// HasCryptoCurrency returns a boolean if a field has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) HasCryptoCurrency() bool {
	if o != nil && !common.IsNil(o.CryptoCurrency) {
		return true
	}

	return false
}

// SetCryptoCurrency gets a reference to the given string and assigns it to the CryptoCurrency field.
func (o *GetFiatPaymentsHistoryResponseDataInner) SetCryptoCurrency(v string) {
	o.CryptoCurrency = &v
}

// GetTotalFee returns the TotalFee field value if set, zero value otherwise.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetTotalFee() string {
	if o == nil || common.IsNil(o.TotalFee) {
		var ret string
		return ret
	}
	return *o.TotalFee
}

// GetTotalFeeOk returns a tuple with the TotalFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetTotalFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalFee) {
		return nil, false
	}
	return o.TotalFee, true
}

// HasTotalFee returns a boolean if a field has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) HasTotalFee() bool {
	if o != nil && !common.IsNil(o.TotalFee) {
		return true
	}

	return false
}

// SetTotalFee gets a reference to the given string and assigns it to the TotalFee field.
func (o *GetFiatPaymentsHistoryResponseDataInner) SetTotalFee(v string) {
	o.TotalFee = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *GetFiatPaymentsHistoryResponseDataInner) SetPrice(v string) {
	o.Price = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetFiatPaymentsHistoryResponseDataInner) SetStatus(v string) {
	o.Status = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetPaymentMethod() string {
	if o == nil || common.IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetPaymentMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) HasPaymentMethod() bool {
	if o != nil && !common.IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *GetFiatPaymentsHistoryResponseDataInner) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *GetFiatPaymentsHistoryResponseDataInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetFiatPaymentsHistoryResponseDataInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetFiatPaymentsHistoryResponseDataInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetFiatPaymentsHistoryResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFiatPaymentsHistoryResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderNo) {
		toSerialize["orderNo"] = o.OrderNo
	}
	if !common.IsNil(o.SourceAmount) {
		toSerialize["sourceAmount"] = o.SourceAmount
	}
	if !common.IsNil(o.FiatCurrency) {
		toSerialize["fiatCurrency"] = o.FiatCurrency
	}
	if !common.IsNil(o.ObtainAmount) {
		toSerialize["obtainAmount"] = o.ObtainAmount
	}
	if !common.IsNil(o.CryptoCurrency) {
		toSerialize["cryptoCurrency"] = o.CryptoCurrency
	}
	if !common.IsNil(o.TotalFee) {
		toSerialize["totalFee"] = o.TotalFee
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFiatPaymentsHistoryResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varGetFiatPaymentsHistoryResponseDataInner := _GetFiatPaymentsHistoryResponseDataInner{}

	err = json.Unmarshal(data, &varGetFiatPaymentsHistoryResponseDataInner)

	if err != nil {
		return err
	}

	*o = GetFiatPaymentsHistoryResponseDataInner(varGetFiatPaymentsHistoryResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderNo")
		delete(additionalProperties, "sourceAmount")
		delete(additionalProperties, "fiatCurrency")
		delete(additionalProperties, "obtainAmount")
		delete(additionalProperties, "cryptoCurrency")
		delete(additionalProperties, "totalFee")
		delete(additionalProperties, "price")
		delete(additionalProperties, "status")
		delete(additionalProperties, "paymentMethod")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFiatPaymentsHistoryResponseDataInner struct {
	value *GetFiatPaymentsHistoryResponseDataInner
	isSet bool
}

func (v NullableGetFiatPaymentsHistoryResponseDataInner) Get() *GetFiatPaymentsHistoryResponseDataInner {
	return v.value
}

func (v *NullableGetFiatPaymentsHistoryResponseDataInner) Set(val *GetFiatPaymentsHistoryResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFiatPaymentsHistoryResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFiatPaymentsHistoryResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFiatPaymentsHistoryResponseDataInner(val *GetFiatPaymentsHistoryResponseDataInner) *NullableGetFiatPaymentsHistoryResponseDataInner {
	return &NullableGetFiatPaymentsHistoryResponseDataInner{value: val, isSet: true}
}

func (v NullableGetFiatPaymentsHistoryResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFiatPaymentsHistoryResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
