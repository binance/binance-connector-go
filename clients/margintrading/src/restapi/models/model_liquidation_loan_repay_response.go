/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the LiquidationLoanRepayResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LiquidationLoanRepayResponse{}

// LiquidationLoanRepayResponse struct for LiquidationLoanRepayResponse
type LiquidationLoanRepayResponse struct {
	// Unique identifier for this repayment transaction
	RepayId *int64 `json:"repayId,omitempty"`
	// Asset used for repayment
	Asset *string `json:"asset,omitempty"`
	// Actual repayment amount
	Amount *string `json:"amount,omitempty"`
	// Repayment status: `SUCCESS` (completed) or `PENDING` (processing)
	Status *string `json:"status,omitempty"`
	// Unix timestamp (milliseconds) when the repayment was created
	CreateTime           *int64 `json:"createTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LiquidationLoanRepayResponse LiquidationLoanRepayResponse

// NewLiquidationLoanRepayResponse instantiates a new LiquidationLoanRepayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiquidationLoanRepayResponse() *LiquidationLoanRepayResponse {
	this := LiquidationLoanRepayResponse{}
	return &this
}

// NewLiquidationLoanRepayResponseWithDefaults instantiates a new LiquidationLoanRepayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiquidationLoanRepayResponseWithDefaults() *LiquidationLoanRepayResponse {
	this := LiquidationLoanRepayResponse{}
	return &this
}

// GetRepayId returns the RepayId field value if set, zero value otherwise.
func (o *LiquidationLoanRepayResponse) GetRepayId() int64 {
	if o == nil || common.IsNil(o.RepayId) {
		var ret int64
		return ret
	}
	return *o.RepayId
}

// GetRepayIdOk returns a tuple with the RepayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiquidationLoanRepayResponse) GetRepayIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RepayId) {
		return nil, false
	}
	return o.RepayId, true
}

// HasRepayId returns a boolean if a field has been set.
func (o *LiquidationLoanRepayResponse) HasRepayId() bool {
	if o != nil && !common.IsNil(o.RepayId) {
		return true
	}

	return false
}

// SetRepayId gets a reference to the given int64 and assigns it to the RepayId field.
func (o *LiquidationLoanRepayResponse) SetRepayId(v int64) {
	o.RepayId = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *LiquidationLoanRepayResponse) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiquidationLoanRepayResponse) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *LiquidationLoanRepayResponse) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *LiquidationLoanRepayResponse) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *LiquidationLoanRepayResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiquidationLoanRepayResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *LiquidationLoanRepayResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *LiquidationLoanRepayResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LiquidationLoanRepayResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiquidationLoanRepayResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LiquidationLoanRepayResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LiquidationLoanRepayResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *LiquidationLoanRepayResponse) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiquidationLoanRepayResponse) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *LiquidationLoanRepayResponse) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *LiquidationLoanRepayResponse) SetCreateTime(v int64) {
	o.CreateTime = &v
}

func (o LiquidationLoanRepayResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiquidationLoanRepayResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RepayId) {
		toSerialize["repayId"] = o.RepayId
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LiquidationLoanRepayResponse) UnmarshalJSON(data []byte) (err error) {
	varLiquidationLoanRepayResponse := _LiquidationLoanRepayResponse{}

	err = json.Unmarshal(data, &varLiquidationLoanRepayResponse)

	if err != nil {
		return err
	}

	*o = LiquidationLoanRepayResponse(varLiquidationLoanRepayResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "repayId")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "status")
		delete(additionalProperties, "createTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLiquidationLoanRepayResponse struct {
	value *LiquidationLoanRepayResponse
	isSet bool
}

func (v NullableLiquidationLoanRepayResponse) Get() *LiquidationLoanRepayResponse {
	return v.value
}

func (v *NullableLiquidationLoanRepayResponse) Set(val *LiquidationLoanRepayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLiquidationLoanRepayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLiquidationLoanRepayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiquidationLoanRepayResponse(val *LiquidationLoanRepayResponse) *NullableLiquidationLoanRepayResponse {
	return &NullableLiquidationLoanRepayResponse{value: val, isSet: true}
}

func (v NullableLiquidationLoanRepayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiquidationLoanRepayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
