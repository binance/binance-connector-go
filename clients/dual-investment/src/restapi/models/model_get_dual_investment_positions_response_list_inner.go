/*
Binance Dual Investment REST API

OpenAPI Specification for the Binance Dual Investment REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetDualInvestmentPositionsResponseListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDualInvestmentPositionsResponseListInner{}

// GetDualInvestmentPositionsResponseListInner struct for GetDualInvestmentPositionsResponseListInner
type GetDualInvestmentPositionsResponseListInner struct {
	Id                   *string `json:"id,omitempty"`
	InvestCoin           *string `json:"investCoin,omitempty"`
	ExercisedCoin        *string `json:"exercisedCoin,omitempty"`
	SubscriptionAmount   *string `json:"subscriptionAmount,omitempty"`
	StrikePrice          *string `json:"strikePrice,omitempty"`
	Duration             *int64  `json:"duration,omitempty"`
	SettleDate           *int64  `json:"settleDate,omitempty"`
	PurchaseStatus       *string `json:"purchaseStatus,omitempty"`
	Apr                  *string `json:"apr,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	PurchaseEndTime      *int64  `json:"purchaseEndTime,omitempty"`
	OptionType           *string `json:"optionType,omitempty"`
	AutoCompoundPlan     *string `json:"autoCompoundPlan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDualInvestmentPositionsResponseListInner GetDualInvestmentPositionsResponseListInner

// NewGetDualInvestmentPositionsResponseListInner instantiates a new GetDualInvestmentPositionsResponseListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDualInvestmentPositionsResponseListInner() *GetDualInvestmentPositionsResponseListInner {
	this := GetDualInvestmentPositionsResponseListInner{}
	return &this
}

// NewGetDualInvestmentPositionsResponseListInnerWithDefaults instantiates a new GetDualInvestmentPositionsResponseListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDualInvestmentPositionsResponseListInnerWithDefaults() *GetDualInvestmentPositionsResponseListInner {
	this := GetDualInvestmentPositionsResponseListInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetDualInvestmentPositionsResponseListInner) SetId(v string) {
	o.Id = &v
}

// GetInvestCoin returns the InvestCoin field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetInvestCoin() string {
	if o == nil || common.IsNil(o.InvestCoin) {
		var ret string
		return ret
	}
	return *o.InvestCoin
}

// GetInvestCoinOk returns a tuple with the InvestCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetInvestCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.InvestCoin) {
		return nil, false
	}
	return o.InvestCoin, true
}

// HasInvestCoin returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasInvestCoin() bool {
	if o != nil && !common.IsNil(o.InvestCoin) {
		return true
	}

	return false
}

// SetInvestCoin gets a reference to the given string and assigns it to the InvestCoin field.
func (o *GetDualInvestmentPositionsResponseListInner) SetInvestCoin(v string) {
	o.InvestCoin = &v
}

// GetExercisedCoin returns the ExercisedCoin field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetExercisedCoin() string {
	if o == nil || common.IsNil(o.ExercisedCoin) {
		var ret string
		return ret
	}
	return *o.ExercisedCoin
}

// GetExercisedCoinOk returns a tuple with the ExercisedCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetExercisedCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExercisedCoin) {
		return nil, false
	}
	return o.ExercisedCoin, true
}

// HasExercisedCoin returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasExercisedCoin() bool {
	if o != nil && !common.IsNil(o.ExercisedCoin) {
		return true
	}

	return false
}

// SetExercisedCoin gets a reference to the given string and assigns it to the ExercisedCoin field.
func (o *GetDualInvestmentPositionsResponseListInner) SetExercisedCoin(v string) {
	o.ExercisedCoin = &v
}

// GetSubscriptionAmount returns the SubscriptionAmount field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetSubscriptionAmount() string {
	if o == nil || common.IsNil(o.SubscriptionAmount) {
		var ret string
		return ret
	}
	return *o.SubscriptionAmount
}

// GetSubscriptionAmountOk returns a tuple with the SubscriptionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetSubscriptionAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubscriptionAmount) {
		return nil, false
	}
	return o.SubscriptionAmount, true
}

// HasSubscriptionAmount returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasSubscriptionAmount() bool {
	if o != nil && !common.IsNil(o.SubscriptionAmount) {
		return true
	}

	return false
}

// SetSubscriptionAmount gets a reference to the given string and assigns it to the SubscriptionAmount field.
func (o *GetDualInvestmentPositionsResponseListInner) SetSubscriptionAmount(v string) {
	o.SubscriptionAmount = &v
}

// GetStrikePrice returns the StrikePrice field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetStrikePrice() string {
	if o == nil || common.IsNil(o.StrikePrice) {
		var ret string
		return ret
	}
	return *o.StrikePrice
}

// GetStrikePriceOk returns a tuple with the StrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetStrikePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrikePrice) {
		return nil, false
	}
	return o.StrikePrice, true
}

// HasStrikePrice returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasStrikePrice() bool {
	if o != nil && !common.IsNil(o.StrikePrice) {
		return true
	}

	return false
}

// SetStrikePrice gets a reference to the given string and assigns it to the StrikePrice field.
func (o *GetDualInvestmentPositionsResponseListInner) SetStrikePrice(v string) {
	o.StrikePrice = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetDuration() int64 {
	if o == nil || common.IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetDurationOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasDuration() bool {
	if o != nil && !common.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *GetDualInvestmentPositionsResponseListInner) SetDuration(v int64) {
	o.Duration = &v
}

// GetSettleDate returns the SettleDate field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetSettleDate() int64 {
	if o == nil || common.IsNil(o.SettleDate) {
		var ret int64
		return ret
	}
	return *o.SettleDate
}

// GetSettleDateOk returns a tuple with the SettleDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetSettleDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SettleDate) {
		return nil, false
	}
	return o.SettleDate, true
}

// HasSettleDate returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasSettleDate() bool {
	if o != nil && !common.IsNil(o.SettleDate) {
		return true
	}

	return false
}

// SetSettleDate gets a reference to the given int64 and assigns it to the SettleDate field.
func (o *GetDualInvestmentPositionsResponseListInner) SetSettleDate(v int64) {
	o.SettleDate = &v
}

// GetPurchaseStatus returns the PurchaseStatus field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetPurchaseStatus() string {
	if o == nil || common.IsNil(o.PurchaseStatus) {
		var ret string
		return ret
	}
	return *o.PurchaseStatus
}

// GetPurchaseStatusOk returns a tuple with the PurchaseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetPurchaseStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.PurchaseStatus) {
		return nil, false
	}
	return o.PurchaseStatus, true
}

// HasPurchaseStatus returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasPurchaseStatus() bool {
	if o != nil && !common.IsNil(o.PurchaseStatus) {
		return true
	}

	return false
}

// SetPurchaseStatus gets a reference to the given string and assigns it to the PurchaseStatus field.
func (o *GetDualInvestmentPositionsResponseListInner) SetPurchaseStatus(v string) {
	o.PurchaseStatus = &v
}

// GetApr returns the Apr field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetApr() string {
	if o == nil || common.IsNil(o.Apr) {
		var ret string
		return ret
	}
	return *o.Apr
}

// GetAprOk returns a tuple with the Apr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetAprOk() (*string, bool) {
	if o == nil || common.IsNil(o.Apr) {
		return nil, false
	}
	return o.Apr, true
}

// HasApr returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasApr() bool {
	if o != nil && !common.IsNil(o.Apr) {
		return true
	}

	return false
}

// SetApr gets a reference to the given string and assigns it to the Apr field.
func (o *GetDualInvestmentPositionsResponseListInner) SetApr(v string) {
	o.Apr = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetDualInvestmentPositionsResponseListInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetPurchaseEndTime returns the PurchaseEndTime field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetPurchaseEndTime() int64 {
	if o == nil || common.IsNil(o.PurchaseEndTime) {
		var ret int64
		return ret
	}
	return *o.PurchaseEndTime
}

// GetPurchaseEndTimeOk returns a tuple with the PurchaseEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetPurchaseEndTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PurchaseEndTime) {
		return nil, false
	}
	return o.PurchaseEndTime, true
}

// HasPurchaseEndTime returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasPurchaseEndTime() bool {
	if o != nil && !common.IsNil(o.PurchaseEndTime) {
		return true
	}

	return false
}

// SetPurchaseEndTime gets a reference to the given int64 and assigns it to the PurchaseEndTime field.
func (o *GetDualInvestmentPositionsResponseListInner) SetPurchaseEndTime(v int64) {
	o.PurchaseEndTime = &v
}

// GetOptionType returns the OptionType field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetOptionType() string {
	if o == nil || common.IsNil(o.OptionType) {
		var ret string
		return ret
	}
	return *o.OptionType
}

// GetOptionTypeOk returns a tuple with the OptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetOptionTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OptionType) {
		return nil, false
	}
	return o.OptionType, true
}

// HasOptionType returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasOptionType() bool {
	if o != nil && !common.IsNil(o.OptionType) {
		return true
	}

	return false
}

// SetOptionType gets a reference to the given string and assigns it to the OptionType field.
func (o *GetDualInvestmentPositionsResponseListInner) SetOptionType(v string) {
	o.OptionType = &v
}

// GetAutoCompoundPlan returns the AutoCompoundPlan field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponseListInner) GetAutoCompoundPlan() string {
	if o == nil || common.IsNil(o.AutoCompoundPlan) {
		var ret string
		return ret
	}
	return *o.AutoCompoundPlan
}

// GetAutoCompoundPlanOk returns a tuple with the AutoCompoundPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponseListInner) GetAutoCompoundPlanOk() (*string, bool) {
	if o == nil || common.IsNil(o.AutoCompoundPlan) {
		return nil, false
	}
	return o.AutoCompoundPlan, true
}

// HasAutoCompoundPlan returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponseListInner) HasAutoCompoundPlan() bool {
	if o != nil && !common.IsNil(o.AutoCompoundPlan) {
		return true
	}

	return false
}

// SetAutoCompoundPlan gets a reference to the given string and assigns it to the AutoCompoundPlan field.
func (o *GetDualInvestmentPositionsResponseListInner) SetAutoCompoundPlan(v string) {
	o.AutoCompoundPlan = &v
}

func (o GetDualInvestmentPositionsResponseListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDualInvestmentPositionsResponseListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.InvestCoin) {
		toSerialize["investCoin"] = o.InvestCoin
	}
	if !common.IsNil(o.ExercisedCoin) {
		toSerialize["exercisedCoin"] = o.ExercisedCoin
	}
	if !common.IsNil(o.SubscriptionAmount) {
		toSerialize["subscriptionAmount"] = o.SubscriptionAmount
	}
	if !common.IsNil(o.StrikePrice) {
		toSerialize["strikePrice"] = o.StrikePrice
	}
	if !common.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !common.IsNil(o.SettleDate) {
		toSerialize["settleDate"] = o.SettleDate
	}
	if !common.IsNil(o.PurchaseStatus) {
		toSerialize["purchaseStatus"] = o.PurchaseStatus
	}
	if !common.IsNil(o.Apr) {
		toSerialize["apr"] = o.Apr
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.PurchaseEndTime) {
		toSerialize["purchaseEndTime"] = o.PurchaseEndTime
	}
	if !common.IsNil(o.OptionType) {
		toSerialize["optionType"] = o.OptionType
	}
	if !common.IsNil(o.AutoCompoundPlan) {
		toSerialize["autoCompoundPlan"] = o.AutoCompoundPlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDualInvestmentPositionsResponseListInner) UnmarshalJSON(data []byte) (err error) {
	varGetDualInvestmentPositionsResponseListInner := _GetDualInvestmentPositionsResponseListInner{}

	err = json.Unmarshal(data, &varGetDualInvestmentPositionsResponseListInner)

	if err != nil {
		return err
	}

	*o = GetDualInvestmentPositionsResponseListInner(varGetDualInvestmentPositionsResponseListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "investCoin")
		delete(additionalProperties, "exercisedCoin")
		delete(additionalProperties, "subscriptionAmount")
		delete(additionalProperties, "strikePrice")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "settleDate")
		delete(additionalProperties, "purchaseStatus")
		delete(additionalProperties, "apr")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "purchaseEndTime")
		delete(additionalProperties, "optionType")
		delete(additionalProperties, "autoCompoundPlan")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDualInvestmentPositionsResponseListInner struct {
	value *GetDualInvestmentPositionsResponseListInner
	isSet bool
}

func (v NullableGetDualInvestmentPositionsResponseListInner) Get() *GetDualInvestmentPositionsResponseListInner {
	return v.value
}

func (v *NullableGetDualInvestmentPositionsResponseListInner) Set(val *GetDualInvestmentPositionsResponseListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDualInvestmentPositionsResponseListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDualInvestmentPositionsResponseListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDualInvestmentPositionsResponseListInner(val *GetDualInvestmentPositionsResponseListInner) *NullableGetDualInvestmentPositionsResponseListInner {
	return &NullableGetDualInvestmentPositionsResponseListInner{value: val, isSet: true}
}

func (v NullableGetDualInvestmentPositionsResponseListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDualInvestmentPositionsResponseListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
