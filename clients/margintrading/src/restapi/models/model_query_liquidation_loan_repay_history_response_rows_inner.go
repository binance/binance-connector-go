/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryLiquidationLoanRepayHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryLiquidationLoanRepayHistoryResponseRowsInner{}

// QueryLiquidationLoanRepayHistoryResponseRowsInner struct for QueryLiquidationLoanRepayHistoryResponseRowsInner
type QueryLiquidationLoanRepayHistoryResponseRowsInner struct {
	// Unique identifier for the repayment transaction
	RepayId *int64 `json:"repayId,omitempty"`
	// Asset used for repayment
	Asset *string `json:"asset,omitempty"`
	// The repayment amount
	Amount *string `json:"amount,omitempty"`
	// Repayment status: `SUCCESS` (completed) or `PENDING` (processing)
	Status *string `json:"status,omitempty"`
	// Unix timestamp (milliseconds) when the repayment was created
	CreateTime           *int64 `json:"createTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryLiquidationLoanRepayHistoryResponseRowsInner QueryLiquidationLoanRepayHistoryResponseRowsInner

// NewQueryLiquidationLoanRepayHistoryResponseRowsInner instantiates a new QueryLiquidationLoanRepayHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLiquidationLoanRepayHistoryResponseRowsInner() *QueryLiquidationLoanRepayHistoryResponseRowsInner {
	this := QueryLiquidationLoanRepayHistoryResponseRowsInner{}
	return &this
}

// NewQueryLiquidationLoanRepayHistoryResponseRowsInnerWithDefaults instantiates a new QueryLiquidationLoanRepayHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLiquidationLoanRepayHistoryResponseRowsInnerWithDefaults() *QueryLiquidationLoanRepayHistoryResponseRowsInner {
	this := QueryLiquidationLoanRepayHistoryResponseRowsInner{}
	return &this
}

// GetRepayId returns the RepayId field value if set, zero value otherwise.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) GetRepayId() int64 {
	if o == nil || common.IsNil(o.RepayId) {
		var ret int64
		return ret
	}
	return *o.RepayId
}

// GetRepayIdOk returns a tuple with the RepayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) GetRepayIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RepayId) {
		return nil, false
	}
	return o.RepayId, true
}

// HasRepayId returns a boolean if a field has been set.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) HasRepayId() bool {
	if o != nil && !common.IsNil(o.RepayId) {
		return true
	}

	return false
}

// SetRepayId gets a reference to the given int64 and assigns it to the RepayId field.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) SetRepayId(v int64) {
	o.RepayId = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

func (o QueryLiquidationLoanRepayHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLiquidationLoanRepayHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
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

func (o *QueryLiquidationLoanRepayHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryLiquidationLoanRepayHistoryResponseRowsInner := _QueryLiquidationLoanRepayHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varQueryLiquidationLoanRepayHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = QueryLiquidationLoanRepayHistoryResponseRowsInner(varQueryLiquidationLoanRepayHistoryResponseRowsInner)

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

type NullableQueryLiquidationLoanRepayHistoryResponseRowsInner struct {
	value *QueryLiquidationLoanRepayHistoryResponseRowsInner
	isSet bool
}

func (v NullableQueryLiquidationLoanRepayHistoryResponseRowsInner) Get() *QueryLiquidationLoanRepayHistoryResponseRowsInner {
	return v.value
}

func (v *NullableQueryLiquidationLoanRepayHistoryResponseRowsInner) Set(val *QueryLiquidationLoanRepayHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLiquidationLoanRepayHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLiquidationLoanRepayHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLiquidationLoanRepayHistoryResponseRowsInner(val *QueryLiquidationLoanRepayHistoryResponseRowsInner) *NullableQueryLiquidationLoanRepayHistoryResponseRowsInner {
	return &NullableQueryLiquidationLoanRepayHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableQueryLiquidationLoanRepayHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLiquidationLoanRepayHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
