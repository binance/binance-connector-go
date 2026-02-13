/*
Binance Fiat REST API

OpenAPI Specification for the Binance Fiat REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFiatDepositWithdrawHistoryResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFiatDepositWithdrawHistoryResponseDataInner{}

// GetFiatDepositWithdrawHistoryResponseDataInner struct for GetFiatDepositWithdrawHistoryResponseDataInner
type GetFiatDepositWithdrawHistoryResponseDataInner struct {
	OrderNo              *string `json:"orderNo,omitempty"`
	FiatCurrency         *string `json:"fiatCurrency,omitempty"`
	IndicatedAmount      *string `json:"indicatedAmount,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	TotalFee             *string `json:"totalFee,omitempty"`
	Method               *string `json:"method,omitempty"`
	Status               *string `json:"status,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFiatDepositWithdrawHistoryResponseDataInner GetFiatDepositWithdrawHistoryResponseDataInner

// NewGetFiatDepositWithdrawHistoryResponseDataInner instantiates a new GetFiatDepositWithdrawHistoryResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFiatDepositWithdrawHistoryResponseDataInner() *GetFiatDepositWithdrawHistoryResponseDataInner {
	this := GetFiatDepositWithdrawHistoryResponseDataInner{}
	return &this
}

// NewGetFiatDepositWithdrawHistoryResponseDataInnerWithDefaults instantiates a new GetFiatDepositWithdrawHistoryResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFiatDepositWithdrawHistoryResponseDataInnerWithDefaults() *GetFiatDepositWithdrawHistoryResponseDataInner {
	this := GetFiatDepositWithdrawHistoryResponseDataInner{}
	return &this
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetOrderNo() string {
	if o == nil || common.IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetOrderNoOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasOrderNo() bool {
	if o != nil && !common.IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetOrderNo(v string) {
	o.OrderNo = &v
}

// GetFiatCurrency returns the FiatCurrency field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetFiatCurrency() string {
	if o == nil || common.IsNil(o.FiatCurrency) {
		var ret string
		return ret
	}
	return *o.FiatCurrency
}

// GetFiatCurrencyOk returns a tuple with the FiatCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetFiatCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.FiatCurrency) {
		return nil, false
	}
	return o.FiatCurrency, true
}

// HasFiatCurrency returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasFiatCurrency() bool {
	if o != nil && !common.IsNil(o.FiatCurrency) {
		return true
	}

	return false
}

// SetFiatCurrency gets a reference to the given string and assigns it to the FiatCurrency field.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetFiatCurrency(v string) {
	o.FiatCurrency = &v
}

// GetIndicatedAmount returns the IndicatedAmount field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetIndicatedAmount() string {
	if o == nil || common.IsNil(o.IndicatedAmount) {
		var ret string
		return ret
	}
	return *o.IndicatedAmount
}

// GetIndicatedAmountOk returns a tuple with the IndicatedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetIndicatedAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.IndicatedAmount) {
		return nil, false
	}
	return o.IndicatedAmount, true
}

// HasIndicatedAmount returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasIndicatedAmount() bool {
	if o != nil && !common.IsNil(o.IndicatedAmount) {
		return true
	}

	return false
}

// SetIndicatedAmount gets a reference to the given string and assigns it to the IndicatedAmount field.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetIndicatedAmount(v string) {
	o.IndicatedAmount = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetAmount(v string) {
	o.Amount = &v
}

// GetTotalFee returns the TotalFee field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetTotalFee() string {
	if o == nil || common.IsNil(o.TotalFee) {
		var ret string
		return ret
	}
	return *o.TotalFee
}

// GetTotalFeeOk returns a tuple with the TotalFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetTotalFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalFee) {
		return nil, false
	}
	return o.TotalFee, true
}

// HasTotalFee returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasTotalFee() bool {
	if o != nil && !common.IsNil(o.TotalFee) {
		return true
	}

	return false
}

// SetTotalFee gets a reference to the given string and assigns it to the TotalFee field.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetTotalFee(v string) {
	o.TotalFee = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetMethod() string {
	if o == nil || common.IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasMethod() bool {
	if o != nil && !common.IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetMethod(v string) {
	o.Method = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetStatus(v string) {
	o.Status = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetFiatDepositWithdrawHistoryResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFiatDepositWithdrawHistoryResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderNo) {
		toSerialize["orderNo"] = o.OrderNo
	}
	if !common.IsNil(o.FiatCurrency) {
		toSerialize["fiatCurrency"] = o.FiatCurrency
	}
	if !common.IsNil(o.IndicatedAmount) {
		toSerialize["indicatedAmount"] = o.IndicatedAmount
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.TotalFee) {
		toSerialize["totalFee"] = o.TotalFee
	}
	if !common.IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
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

func (o *GetFiatDepositWithdrawHistoryResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varGetFiatDepositWithdrawHistoryResponseDataInner := _GetFiatDepositWithdrawHistoryResponseDataInner{}

	err = json.Unmarshal(data, &varGetFiatDepositWithdrawHistoryResponseDataInner)

	if err != nil {
		return err
	}

	*o = GetFiatDepositWithdrawHistoryResponseDataInner(varGetFiatDepositWithdrawHistoryResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderNo")
		delete(additionalProperties, "fiatCurrency")
		delete(additionalProperties, "indicatedAmount")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "totalFee")
		delete(additionalProperties, "method")
		delete(additionalProperties, "status")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFiatDepositWithdrawHistoryResponseDataInner struct {
	value *GetFiatDepositWithdrawHistoryResponseDataInner
	isSet bool
}

func (v NullableGetFiatDepositWithdrawHistoryResponseDataInner) Get() *GetFiatDepositWithdrawHistoryResponseDataInner {
	return v.value
}

func (v *NullableGetFiatDepositWithdrawHistoryResponseDataInner) Set(val *GetFiatDepositWithdrawHistoryResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFiatDepositWithdrawHistoryResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFiatDepositWithdrawHistoryResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFiatDepositWithdrawHistoryResponseDataInner(val *GetFiatDepositWithdrawHistoryResponseDataInner) *NullableGetFiatDepositWithdrawHistoryResponseDataInner {
	return &NullableGetFiatDepositWithdrawHistoryResponseDataInner{value: val, isSet: true}
}

func (v NullableGetFiatDepositWithdrawHistoryResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFiatDepositWithdrawHistoryResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
