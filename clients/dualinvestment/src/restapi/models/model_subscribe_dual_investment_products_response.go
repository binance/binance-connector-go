/*
Binance Dual Investment REST API

OpenAPI Specification for the Binance Dual Investment REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SubscribeDualInvestmentProductsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubscribeDualInvestmentProductsResponse{}

// SubscribeDualInvestmentProductsResponse struct for SubscribeDualInvestmentProductsResponse
type SubscribeDualInvestmentProductsResponse struct {
	PositionId           *int64  `json:"positionId,omitempty"`
	InvestCoin           *string `json:"investCoin,omitempty"`
	ExercisedCoin        *string `json:"exercisedCoin,omitempty"`
	SubscriptionAmount   *string `json:"subscriptionAmount,omitempty"`
	Duration             *int64  `json:"duration,omitempty"`
	AutoCompoundPlan     *string `json:"autoCompoundPlan,omitempty"`
	StrikePrice          *string `json:"strikePrice,omitempty"`
	SettleDate           *int64  `json:"settleDate,omitempty"`
	PurchaseStatus       *string `json:"purchaseStatus,omitempty"`
	Apr                  *string `json:"apr,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	PurchaseTime         *int64  `json:"purchaseTime,omitempty"`
	OptionType           *string `json:"optionType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscribeDualInvestmentProductsResponse SubscribeDualInvestmentProductsResponse

// NewSubscribeDualInvestmentProductsResponse instantiates a new SubscribeDualInvestmentProductsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribeDualInvestmentProductsResponse() *SubscribeDualInvestmentProductsResponse {
	this := SubscribeDualInvestmentProductsResponse{}
	return &this
}

// NewSubscribeDualInvestmentProductsResponseWithDefaults instantiates a new SubscribeDualInvestmentProductsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribeDualInvestmentProductsResponseWithDefaults() *SubscribeDualInvestmentProductsResponse {
	this := SubscribeDualInvestmentProductsResponse{}
	return &this
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetPositionId() int64 {
	if o == nil || common.IsNil(o.PositionId) {
		var ret int64
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetPositionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given int64 and assigns it to the PositionId field.
func (o *SubscribeDualInvestmentProductsResponse) SetPositionId(v int64) {
	o.PositionId = &v
}

// GetInvestCoin returns the InvestCoin field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetInvestCoin() string {
	if o == nil || common.IsNil(o.InvestCoin) {
		var ret string
		return ret
	}
	return *o.InvestCoin
}

// GetInvestCoinOk returns a tuple with the InvestCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetInvestCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.InvestCoin) {
		return nil, false
	}
	return o.InvestCoin, true
}

// HasInvestCoin returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasInvestCoin() bool {
	if o != nil && !common.IsNil(o.InvestCoin) {
		return true
	}

	return false
}

// SetInvestCoin gets a reference to the given string and assigns it to the InvestCoin field.
func (o *SubscribeDualInvestmentProductsResponse) SetInvestCoin(v string) {
	o.InvestCoin = &v
}

// GetExercisedCoin returns the ExercisedCoin field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetExercisedCoin() string {
	if o == nil || common.IsNil(o.ExercisedCoin) {
		var ret string
		return ret
	}
	return *o.ExercisedCoin
}

// GetExercisedCoinOk returns a tuple with the ExercisedCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetExercisedCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExercisedCoin) {
		return nil, false
	}
	return o.ExercisedCoin, true
}

// HasExercisedCoin returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasExercisedCoin() bool {
	if o != nil && !common.IsNil(o.ExercisedCoin) {
		return true
	}

	return false
}

// SetExercisedCoin gets a reference to the given string and assigns it to the ExercisedCoin field.
func (o *SubscribeDualInvestmentProductsResponse) SetExercisedCoin(v string) {
	o.ExercisedCoin = &v
}

// GetSubscriptionAmount returns the SubscriptionAmount field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetSubscriptionAmount() string {
	if o == nil || common.IsNil(o.SubscriptionAmount) {
		var ret string
		return ret
	}
	return *o.SubscriptionAmount
}

// GetSubscriptionAmountOk returns a tuple with the SubscriptionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetSubscriptionAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubscriptionAmount) {
		return nil, false
	}
	return o.SubscriptionAmount, true
}

// HasSubscriptionAmount returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasSubscriptionAmount() bool {
	if o != nil && !common.IsNil(o.SubscriptionAmount) {
		return true
	}

	return false
}

// SetSubscriptionAmount gets a reference to the given string and assigns it to the SubscriptionAmount field.
func (o *SubscribeDualInvestmentProductsResponse) SetSubscriptionAmount(v string) {
	o.SubscriptionAmount = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetDuration() int64 {
	if o == nil || common.IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetDurationOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasDuration() bool {
	if o != nil && !common.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *SubscribeDualInvestmentProductsResponse) SetDuration(v int64) {
	o.Duration = &v
}

// GetAutoCompoundPlan returns the AutoCompoundPlan field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetAutoCompoundPlan() string {
	if o == nil || common.IsNil(o.AutoCompoundPlan) {
		var ret string
		return ret
	}
	return *o.AutoCompoundPlan
}

// GetAutoCompoundPlanOk returns a tuple with the AutoCompoundPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetAutoCompoundPlanOk() (*string, bool) {
	if o == nil || common.IsNil(o.AutoCompoundPlan) {
		return nil, false
	}
	return o.AutoCompoundPlan, true
}

// HasAutoCompoundPlan returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasAutoCompoundPlan() bool {
	if o != nil && !common.IsNil(o.AutoCompoundPlan) {
		return true
	}

	return false
}

// SetAutoCompoundPlan gets a reference to the given string and assigns it to the AutoCompoundPlan field.
func (o *SubscribeDualInvestmentProductsResponse) SetAutoCompoundPlan(v string) {
	o.AutoCompoundPlan = &v
}

// GetStrikePrice returns the StrikePrice field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetStrikePrice() string {
	if o == nil || common.IsNil(o.StrikePrice) {
		var ret string
		return ret
	}
	return *o.StrikePrice
}

// GetStrikePriceOk returns a tuple with the StrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetStrikePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrikePrice) {
		return nil, false
	}
	return o.StrikePrice, true
}

// HasStrikePrice returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasStrikePrice() bool {
	if o != nil && !common.IsNil(o.StrikePrice) {
		return true
	}

	return false
}

// SetStrikePrice gets a reference to the given string and assigns it to the StrikePrice field.
func (o *SubscribeDualInvestmentProductsResponse) SetStrikePrice(v string) {
	o.StrikePrice = &v
}

// GetSettleDate returns the SettleDate field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetSettleDate() int64 {
	if o == nil || common.IsNil(o.SettleDate) {
		var ret int64
		return ret
	}
	return *o.SettleDate
}

// GetSettleDateOk returns a tuple with the SettleDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetSettleDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SettleDate) {
		return nil, false
	}
	return o.SettleDate, true
}

// HasSettleDate returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasSettleDate() bool {
	if o != nil && !common.IsNil(o.SettleDate) {
		return true
	}

	return false
}

// SetSettleDate gets a reference to the given int64 and assigns it to the SettleDate field.
func (o *SubscribeDualInvestmentProductsResponse) SetSettleDate(v int64) {
	o.SettleDate = &v
}

// GetPurchaseStatus returns the PurchaseStatus field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetPurchaseStatus() string {
	if o == nil || common.IsNil(o.PurchaseStatus) {
		var ret string
		return ret
	}
	return *o.PurchaseStatus
}

// GetPurchaseStatusOk returns a tuple with the PurchaseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetPurchaseStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.PurchaseStatus) {
		return nil, false
	}
	return o.PurchaseStatus, true
}

// HasPurchaseStatus returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasPurchaseStatus() bool {
	if o != nil && !common.IsNil(o.PurchaseStatus) {
		return true
	}

	return false
}

// SetPurchaseStatus gets a reference to the given string and assigns it to the PurchaseStatus field.
func (o *SubscribeDualInvestmentProductsResponse) SetPurchaseStatus(v string) {
	o.PurchaseStatus = &v
}

// GetApr returns the Apr field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetApr() string {
	if o == nil || common.IsNil(o.Apr) {
		var ret string
		return ret
	}
	return *o.Apr
}

// GetAprOk returns a tuple with the Apr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetAprOk() (*string, bool) {
	if o == nil || common.IsNil(o.Apr) {
		return nil, false
	}
	return o.Apr, true
}

// HasApr returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasApr() bool {
	if o != nil && !common.IsNil(o.Apr) {
		return true
	}

	return false
}

// SetApr gets a reference to the given string and assigns it to the Apr field.
func (o *SubscribeDualInvestmentProductsResponse) SetApr(v string) {
	o.Apr = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *SubscribeDualInvestmentProductsResponse) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetPurchaseTime returns the PurchaseTime field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetPurchaseTime() int64 {
	if o == nil || common.IsNil(o.PurchaseTime) {
		var ret int64
		return ret
	}
	return *o.PurchaseTime
}

// GetPurchaseTimeOk returns a tuple with the PurchaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetPurchaseTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PurchaseTime) {
		return nil, false
	}
	return o.PurchaseTime, true
}

// HasPurchaseTime returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasPurchaseTime() bool {
	if o != nil && !common.IsNil(o.PurchaseTime) {
		return true
	}

	return false
}

// SetPurchaseTime gets a reference to the given int64 and assigns it to the PurchaseTime field.
func (o *SubscribeDualInvestmentProductsResponse) SetPurchaseTime(v int64) {
	o.PurchaseTime = &v
}

// GetOptionType returns the OptionType field value if set, zero value otherwise.
func (o *SubscribeDualInvestmentProductsResponse) GetOptionType() string {
	if o == nil || common.IsNil(o.OptionType) {
		var ret string
		return ret
	}
	return *o.OptionType
}

// GetOptionTypeOk returns a tuple with the OptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeDualInvestmentProductsResponse) GetOptionTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OptionType) {
		return nil, false
	}
	return o.OptionType, true
}

// HasOptionType returns a boolean if a field has been set.
func (o *SubscribeDualInvestmentProductsResponse) HasOptionType() bool {
	if o != nil && !common.IsNil(o.OptionType) {
		return true
	}

	return false
}

// SetOptionType gets a reference to the given string and assigns it to the OptionType field.
func (o *SubscribeDualInvestmentProductsResponse) SetOptionType(v string) {
	o.OptionType = &v
}

func (o SubscribeDualInvestmentProductsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribeDualInvestmentProductsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PositionId) {
		toSerialize["positionId"] = o.PositionId
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
	if !common.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !common.IsNil(o.AutoCompoundPlan) {
		toSerialize["autoCompoundPlan"] = o.AutoCompoundPlan
	}
	if !common.IsNil(o.StrikePrice) {
		toSerialize["strikePrice"] = o.StrikePrice
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
	if !common.IsNil(o.PurchaseTime) {
		toSerialize["purchaseTime"] = o.PurchaseTime
	}
	if !common.IsNil(o.OptionType) {
		toSerialize["optionType"] = o.OptionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscribeDualInvestmentProductsResponse) UnmarshalJSON(data []byte) (err error) {
	varSubscribeDualInvestmentProductsResponse := _SubscribeDualInvestmentProductsResponse{}

	err = json.Unmarshal(data, &varSubscribeDualInvestmentProductsResponse)

	if err != nil {
		return err
	}

	*o = SubscribeDualInvestmentProductsResponse(varSubscribeDualInvestmentProductsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "positionId")
		delete(additionalProperties, "investCoin")
		delete(additionalProperties, "exercisedCoin")
		delete(additionalProperties, "subscriptionAmount")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "autoCompoundPlan")
		delete(additionalProperties, "strikePrice")
		delete(additionalProperties, "settleDate")
		delete(additionalProperties, "purchaseStatus")
		delete(additionalProperties, "apr")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "purchaseTime")
		delete(additionalProperties, "optionType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscribeDualInvestmentProductsResponse struct {
	value *SubscribeDualInvestmentProductsResponse
	isSet bool
}

func (v NullableSubscribeDualInvestmentProductsResponse) Get() *SubscribeDualInvestmentProductsResponse {
	return v.value
}

func (v *NullableSubscribeDualInvestmentProductsResponse) Set(val *SubscribeDualInvestmentProductsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeDualInvestmentProductsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeDualInvestmentProductsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeDualInvestmentProductsResponse(val *SubscribeDualInvestmentProductsResponse) *NullableSubscribeDualInvestmentProductsResponse {
	return &NullableSubscribeDualInvestmentProductsResponse{value: val, isSet: true}
}

func (v NullableSubscribeDualInvestmentProductsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeDualInvestmentProductsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
