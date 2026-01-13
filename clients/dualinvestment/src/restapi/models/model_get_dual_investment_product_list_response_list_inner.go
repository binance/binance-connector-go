/*
Binance Dual Investment REST API

OpenAPI Specification for the Binance Dual Investment REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetDualInvestmentProductListResponseListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDualInvestmentProductListResponseListInner{}

// GetDualInvestmentProductListResponseListInner struct for GetDualInvestmentProductListResponseListInner
type GetDualInvestmentProductListResponseListInner struct {
	Id                   *string  `json:"id,omitempty"`
	InvestCoin           *string  `json:"investCoin,omitempty"`
	ExercisedCoin        *string  `json:"exercisedCoin,omitempty"`
	StrikePrice          *string  `json:"strikePrice,omitempty"`
	Duration             *int64   `json:"duration,omitempty"`
	SettleDate           *int64   `json:"settleDate,omitempty"`
	PurchaseDecimal      *int64   `json:"purchaseDecimal,omitempty"`
	PurchaseEndTime      *int64   `json:"purchaseEndTime,omitempty"`
	CanPurchase          *bool    `json:"canPurchase,omitempty"`
	Apr                  *string  `json:"apr,omitempty"`
	OrderId              *int64   `json:"orderId,omitempty"`
	MinAmount            *string  `json:"minAmount,omitempty"`
	MaxAmount            *string  `json:"maxAmount,omitempty"`
	CreateTimestamp      *int64   `json:"createTimestamp,omitempty"`
	OptionType           *string  `json:"optionType,omitempty"`
	IsAutoCompoundEnable *bool    `json:"isAutoCompoundEnable,omitempty"`
	AutoCompoundPlanList []string `json:"autoCompoundPlanList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDualInvestmentProductListResponseListInner GetDualInvestmentProductListResponseListInner

// NewGetDualInvestmentProductListResponseListInner instantiates a new GetDualInvestmentProductListResponseListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDualInvestmentProductListResponseListInner() *GetDualInvestmentProductListResponseListInner {
	this := GetDualInvestmentProductListResponseListInner{}
	return &this
}

// NewGetDualInvestmentProductListResponseListInnerWithDefaults instantiates a new GetDualInvestmentProductListResponseListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDualInvestmentProductListResponseListInnerWithDefaults() *GetDualInvestmentProductListResponseListInner {
	this := GetDualInvestmentProductListResponseListInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetDualInvestmentProductListResponseListInner) SetId(v string) {
	o.Id = &v
}

// GetInvestCoin returns the InvestCoin field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetInvestCoin() string {
	if o == nil || common.IsNil(o.InvestCoin) {
		var ret string
		return ret
	}
	return *o.InvestCoin
}

// GetInvestCoinOk returns a tuple with the InvestCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetInvestCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.InvestCoin) {
		return nil, false
	}
	return o.InvestCoin, true
}

// HasInvestCoin returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasInvestCoin() bool {
	if o != nil && !common.IsNil(o.InvestCoin) {
		return true
	}

	return false
}

// SetInvestCoin gets a reference to the given string and assigns it to the InvestCoin field.
func (o *GetDualInvestmentProductListResponseListInner) SetInvestCoin(v string) {
	o.InvestCoin = &v
}

// GetExercisedCoin returns the ExercisedCoin field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetExercisedCoin() string {
	if o == nil || common.IsNil(o.ExercisedCoin) {
		var ret string
		return ret
	}
	return *o.ExercisedCoin
}

// GetExercisedCoinOk returns a tuple with the ExercisedCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetExercisedCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExercisedCoin) {
		return nil, false
	}
	return o.ExercisedCoin, true
}

// HasExercisedCoin returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasExercisedCoin() bool {
	if o != nil && !common.IsNil(o.ExercisedCoin) {
		return true
	}

	return false
}

// SetExercisedCoin gets a reference to the given string and assigns it to the ExercisedCoin field.
func (o *GetDualInvestmentProductListResponseListInner) SetExercisedCoin(v string) {
	o.ExercisedCoin = &v
}

// GetStrikePrice returns the StrikePrice field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetStrikePrice() string {
	if o == nil || common.IsNil(o.StrikePrice) {
		var ret string
		return ret
	}
	return *o.StrikePrice
}

// GetStrikePriceOk returns a tuple with the StrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetStrikePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrikePrice) {
		return nil, false
	}
	return o.StrikePrice, true
}

// HasStrikePrice returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasStrikePrice() bool {
	if o != nil && !common.IsNil(o.StrikePrice) {
		return true
	}

	return false
}

// SetStrikePrice gets a reference to the given string and assigns it to the StrikePrice field.
func (o *GetDualInvestmentProductListResponseListInner) SetStrikePrice(v string) {
	o.StrikePrice = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetDuration() int64 {
	if o == nil || common.IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetDurationOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasDuration() bool {
	if o != nil && !common.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *GetDualInvestmentProductListResponseListInner) SetDuration(v int64) {
	o.Duration = &v
}

// GetSettleDate returns the SettleDate field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetSettleDate() int64 {
	if o == nil || common.IsNil(o.SettleDate) {
		var ret int64
		return ret
	}
	return *o.SettleDate
}

// GetSettleDateOk returns a tuple with the SettleDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetSettleDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SettleDate) {
		return nil, false
	}
	return o.SettleDate, true
}

// HasSettleDate returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasSettleDate() bool {
	if o != nil && !common.IsNil(o.SettleDate) {
		return true
	}

	return false
}

// SetSettleDate gets a reference to the given int64 and assigns it to the SettleDate field.
func (o *GetDualInvestmentProductListResponseListInner) SetSettleDate(v int64) {
	o.SettleDate = &v
}

// GetPurchaseDecimal returns the PurchaseDecimal field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetPurchaseDecimal() int64 {
	if o == nil || common.IsNil(o.PurchaseDecimal) {
		var ret int64
		return ret
	}
	return *o.PurchaseDecimal
}

// GetPurchaseDecimalOk returns a tuple with the PurchaseDecimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetPurchaseDecimalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PurchaseDecimal) {
		return nil, false
	}
	return o.PurchaseDecimal, true
}

// HasPurchaseDecimal returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasPurchaseDecimal() bool {
	if o != nil && !common.IsNil(o.PurchaseDecimal) {
		return true
	}

	return false
}

// SetPurchaseDecimal gets a reference to the given int64 and assigns it to the PurchaseDecimal field.
func (o *GetDualInvestmentProductListResponseListInner) SetPurchaseDecimal(v int64) {
	o.PurchaseDecimal = &v
}

// GetPurchaseEndTime returns the PurchaseEndTime field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetPurchaseEndTime() int64 {
	if o == nil || common.IsNil(o.PurchaseEndTime) {
		var ret int64
		return ret
	}
	return *o.PurchaseEndTime
}

// GetPurchaseEndTimeOk returns a tuple with the PurchaseEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetPurchaseEndTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PurchaseEndTime) {
		return nil, false
	}
	return o.PurchaseEndTime, true
}

// HasPurchaseEndTime returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasPurchaseEndTime() bool {
	if o != nil && !common.IsNil(o.PurchaseEndTime) {
		return true
	}

	return false
}

// SetPurchaseEndTime gets a reference to the given int64 and assigns it to the PurchaseEndTime field.
func (o *GetDualInvestmentProductListResponseListInner) SetPurchaseEndTime(v int64) {
	o.PurchaseEndTime = &v
}

// GetCanPurchase returns the CanPurchase field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetCanPurchase() bool {
	if o == nil || common.IsNil(o.CanPurchase) {
		var ret bool
		return ret
	}
	return *o.CanPurchase
}

// GetCanPurchaseOk returns a tuple with the CanPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetCanPurchaseOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanPurchase) {
		return nil, false
	}
	return o.CanPurchase, true
}

// HasCanPurchase returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasCanPurchase() bool {
	if o != nil && !common.IsNil(o.CanPurchase) {
		return true
	}

	return false
}

// SetCanPurchase gets a reference to the given bool and assigns it to the CanPurchase field.
func (o *GetDualInvestmentProductListResponseListInner) SetCanPurchase(v bool) {
	o.CanPurchase = &v
}

// GetApr returns the Apr field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetApr() string {
	if o == nil || common.IsNil(o.Apr) {
		var ret string
		return ret
	}
	return *o.Apr
}

// GetAprOk returns a tuple with the Apr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetAprOk() (*string, bool) {
	if o == nil || common.IsNil(o.Apr) {
		return nil, false
	}
	return o.Apr, true
}

// HasApr returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasApr() bool {
	if o != nil && !common.IsNil(o.Apr) {
		return true
	}

	return false
}

// SetApr gets a reference to the given string and assigns it to the Apr field.
func (o *GetDualInvestmentProductListResponseListInner) SetApr(v string) {
	o.Apr = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetDualInvestmentProductListResponseListInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetMinAmount returns the MinAmount field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetMinAmount() string {
	if o == nil || common.IsNil(o.MinAmount) {
		var ret string
		return ret
	}
	return *o.MinAmount
}

// GetMinAmountOk returns a tuple with the MinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetMinAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinAmount) {
		return nil, false
	}
	return o.MinAmount, true
}

// HasMinAmount returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasMinAmount() bool {
	if o != nil && !common.IsNil(o.MinAmount) {
		return true
	}

	return false
}

// SetMinAmount gets a reference to the given string and assigns it to the MinAmount field.
func (o *GetDualInvestmentProductListResponseListInner) SetMinAmount(v string) {
	o.MinAmount = &v
}

// GetMaxAmount returns the MaxAmount field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetMaxAmount() string {
	if o == nil || common.IsNil(o.MaxAmount) {
		var ret string
		return ret
	}
	return *o.MaxAmount
}

// GetMaxAmountOk returns a tuple with the MaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetMaxAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxAmount) {
		return nil, false
	}
	return o.MaxAmount, true
}

// HasMaxAmount returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasMaxAmount() bool {
	if o != nil && !common.IsNil(o.MaxAmount) {
		return true
	}

	return false
}

// SetMaxAmount gets a reference to the given string and assigns it to the MaxAmount field.
func (o *GetDualInvestmentProductListResponseListInner) SetMaxAmount(v string) {
	o.MaxAmount = &v
}

// GetCreateTimestamp returns the CreateTimestamp field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetCreateTimestamp() int64 {
	if o == nil || common.IsNil(o.CreateTimestamp) {
		var ret int64
		return ret
	}
	return *o.CreateTimestamp
}

// GetCreateTimestampOk returns a tuple with the CreateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetCreateTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTimestamp) {
		return nil, false
	}
	return o.CreateTimestamp, true
}

// HasCreateTimestamp returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasCreateTimestamp() bool {
	if o != nil && !common.IsNil(o.CreateTimestamp) {
		return true
	}

	return false
}

// SetCreateTimestamp gets a reference to the given int64 and assigns it to the CreateTimestamp field.
func (o *GetDualInvestmentProductListResponseListInner) SetCreateTimestamp(v int64) {
	o.CreateTimestamp = &v
}

// GetOptionType returns the OptionType field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetOptionType() string {
	if o == nil || common.IsNil(o.OptionType) {
		var ret string
		return ret
	}
	return *o.OptionType
}

// GetOptionTypeOk returns a tuple with the OptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetOptionTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OptionType) {
		return nil, false
	}
	return o.OptionType, true
}

// HasOptionType returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasOptionType() bool {
	if o != nil && !common.IsNil(o.OptionType) {
		return true
	}

	return false
}

// SetOptionType gets a reference to the given string and assigns it to the OptionType field.
func (o *GetDualInvestmentProductListResponseListInner) SetOptionType(v string) {
	o.OptionType = &v
}

// GetIsAutoCompoundEnable returns the IsAutoCompoundEnable field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetIsAutoCompoundEnable() bool {
	if o == nil || common.IsNil(o.IsAutoCompoundEnable) {
		var ret bool
		return ret
	}
	return *o.IsAutoCompoundEnable
}

// GetIsAutoCompoundEnableOk returns a tuple with the IsAutoCompoundEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetIsAutoCompoundEnableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsAutoCompoundEnable) {
		return nil, false
	}
	return o.IsAutoCompoundEnable, true
}

// HasIsAutoCompoundEnable returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasIsAutoCompoundEnable() bool {
	if o != nil && !common.IsNil(o.IsAutoCompoundEnable) {
		return true
	}

	return false
}

// SetIsAutoCompoundEnable gets a reference to the given bool and assigns it to the IsAutoCompoundEnable field.
func (o *GetDualInvestmentProductListResponseListInner) SetIsAutoCompoundEnable(v bool) {
	o.IsAutoCompoundEnable = &v
}

// GetAutoCompoundPlanList returns the AutoCompoundPlanList field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponseListInner) GetAutoCompoundPlanList() []string {
	if o == nil || common.IsNil(o.AutoCompoundPlanList) {
		var ret []string
		return ret
	}
	return o.AutoCompoundPlanList
}

// GetAutoCompoundPlanListOk returns a tuple with the AutoCompoundPlanList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponseListInner) GetAutoCompoundPlanListOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AutoCompoundPlanList) {
		return nil, false
	}
	return o.AutoCompoundPlanList, true
}

// HasAutoCompoundPlanList returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponseListInner) HasAutoCompoundPlanList() bool {
	if o != nil && !common.IsNil(o.AutoCompoundPlanList) {
		return true
	}

	return false
}

// SetAutoCompoundPlanList gets a reference to the given []string and assigns it to the AutoCompoundPlanList field.
func (o *GetDualInvestmentProductListResponseListInner) SetAutoCompoundPlanList(v []string) {
	o.AutoCompoundPlanList = v
}

func (o GetDualInvestmentProductListResponseListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDualInvestmentProductListResponseListInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.StrikePrice) {
		toSerialize["strikePrice"] = o.StrikePrice
	}
	if !common.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !common.IsNil(o.SettleDate) {
		toSerialize["settleDate"] = o.SettleDate
	}
	if !common.IsNil(o.PurchaseDecimal) {
		toSerialize["purchaseDecimal"] = o.PurchaseDecimal
	}
	if !common.IsNil(o.PurchaseEndTime) {
		toSerialize["purchaseEndTime"] = o.PurchaseEndTime
	}
	if !common.IsNil(o.CanPurchase) {
		toSerialize["canPurchase"] = o.CanPurchase
	}
	if !common.IsNil(o.Apr) {
		toSerialize["apr"] = o.Apr
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.MinAmount) {
		toSerialize["minAmount"] = o.MinAmount
	}
	if !common.IsNil(o.MaxAmount) {
		toSerialize["maxAmount"] = o.MaxAmount
	}
	if !common.IsNil(o.CreateTimestamp) {
		toSerialize["createTimestamp"] = o.CreateTimestamp
	}
	if !common.IsNil(o.OptionType) {
		toSerialize["optionType"] = o.OptionType
	}
	if !common.IsNil(o.IsAutoCompoundEnable) {
		toSerialize["isAutoCompoundEnable"] = o.IsAutoCompoundEnable
	}
	if !common.IsNil(o.AutoCompoundPlanList) {
		toSerialize["autoCompoundPlanList"] = o.AutoCompoundPlanList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDualInvestmentProductListResponseListInner) UnmarshalJSON(data []byte) (err error) {
	varGetDualInvestmentProductListResponseListInner := _GetDualInvestmentProductListResponseListInner{}

	err = json.Unmarshal(data, &varGetDualInvestmentProductListResponseListInner)

	if err != nil {
		return err
	}

	*o = GetDualInvestmentProductListResponseListInner(varGetDualInvestmentProductListResponseListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "investCoin")
		delete(additionalProperties, "exercisedCoin")
		delete(additionalProperties, "strikePrice")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "settleDate")
		delete(additionalProperties, "purchaseDecimal")
		delete(additionalProperties, "purchaseEndTime")
		delete(additionalProperties, "canPurchase")
		delete(additionalProperties, "apr")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "minAmount")
		delete(additionalProperties, "maxAmount")
		delete(additionalProperties, "createTimestamp")
		delete(additionalProperties, "optionType")
		delete(additionalProperties, "isAutoCompoundEnable")
		delete(additionalProperties, "autoCompoundPlanList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDualInvestmentProductListResponseListInner struct {
	value *GetDualInvestmentProductListResponseListInner
	isSet bool
}

func (v NullableGetDualInvestmentProductListResponseListInner) Get() *GetDualInvestmentProductListResponseListInner {
	return v.value
}

func (v *NullableGetDualInvestmentProductListResponseListInner) Set(val *GetDualInvestmentProductListResponseListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDualInvestmentProductListResponseListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDualInvestmentProductListResponseListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDualInvestmentProductListResponseListInner(val *GetDualInvestmentProductListResponseListInner) *NullableGetDualInvestmentProductListResponseListInner {
	return &NullableGetDualInvestmentProductListResponseListInner{value: val, isSet: true}
}

func (v NullableGetDualInvestmentProductListResponseListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDualInvestmentProductListResponseListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
