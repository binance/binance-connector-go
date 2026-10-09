/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetMarketDetailResponseMarketsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMarketDetailResponseMarketsInner{}

// GetMarketDetailResponseMarketsInner struct for GetMarketDetailResponseMarketsInner
type GetMarketDetailResponseMarketsInner struct {
	MarketId      *int64  `json:"marketId,omitempty"`
	ExternalId    *string `json:"externalId,omitempty"`
	Title         *string `json:"title,omitempty"`
	Question      *string `json:"question,omitempty"`
	Description   *string `json:"description,omitempty"`
	ConditionId   *string `json:"conditionId,omitempty"`
	Status        *string `json:"status,omitempty"`
	TradingStatus *string `json:"tradingStatus,omitempty"`
	TradeVolume   *string `json:"tradeVolume,omitempty"`
	Liquidity     *string `json:"liquidity,omitempty"`
	// Price precision of the market, in decimal places. The price tick is 10 to the power of minus `decimalPrecision`, for example `2` means a tick of 0.01. A limit price that is not on the tick is truncated down
	DecimalPrecision     *int32                                             `json:"decimalPrecision,omitempty"`
	Outcomes             []GetMarketDetailResponseMarketsInnerOutcomesInner `json:"outcomes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMarketDetailResponseMarketsInner GetMarketDetailResponseMarketsInner

// NewGetMarketDetailResponseMarketsInner instantiates a new GetMarketDetailResponseMarketsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketDetailResponseMarketsInner() *GetMarketDetailResponseMarketsInner {
	this := GetMarketDetailResponseMarketsInner{}
	return &this
}

// NewGetMarketDetailResponseMarketsInnerWithDefaults instantiates a new GetMarketDetailResponseMarketsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketDetailResponseMarketsInnerWithDefaults() *GetMarketDetailResponseMarketsInner {
	this := GetMarketDetailResponseMarketsInner{}
	return &this
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInner) GetMarketId() int64 {
	if o == nil || common.IsNil(o.MarketId) {
		var ret int64
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInner) GetMarketIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketId) {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInner) HasMarketId() bool {
	if o != nil && !common.IsNil(o.MarketId) {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given int64 and assigns it to the MarketId field.
func (o *GetMarketDetailResponseMarketsInner) SetMarketId(v int64) {
	o.MarketId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInner) GetExternalId() string {
	if o == nil || common.IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInner) GetExternalIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInner) HasExternalId() bool {
	if o != nil && !common.IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *GetMarketDetailResponseMarketsInner) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInner) GetTitle() string {
	if o == nil || common.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInner) GetTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInner) HasTitle() bool {
	if o != nil && !common.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetMarketDetailResponseMarketsInner) SetTitle(v string) {
	o.Title = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInner) GetQuestion() string {
	if o == nil || common.IsNil(o.Question) {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInner) GetQuestionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Question) {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInner) HasQuestion() bool {
	if o != nil && !common.IsNil(o.Question) {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *GetMarketDetailResponseMarketsInner) SetQuestion(v string) {
	o.Question = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInner) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInner) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetMarketDetailResponseMarketsInner) SetDescription(v string) {
	o.Description = &v
}

// GetConditionId returns the ConditionId field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInner) GetConditionId() string {
	if o == nil || common.IsNil(o.ConditionId) {
		var ret string
		return ret
	}
	return *o.ConditionId
}

// GetConditionIdOk returns a tuple with the ConditionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInner) GetConditionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ConditionId) {
		return nil, false
	}
	return o.ConditionId, true
}

// HasConditionId returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInner) HasConditionId() bool {
	if o != nil && !common.IsNil(o.ConditionId) {
		return true
	}

	return false
}

// SetConditionId gets a reference to the given string and assigns it to the ConditionId field.
func (o *GetMarketDetailResponseMarketsInner) SetConditionId(v string) {
	o.ConditionId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetMarketDetailResponseMarketsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTradingStatus returns the TradingStatus field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInner) GetTradingStatus() string {
	if o == nil || common.IsNil(o.TradingStatus) {
		var ret string
		return ret
	}
	return *o.TradingStatus
}

// GetTradingStatusOk returns a tuple with the TradingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInner) GetTradingStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.TradingStatus) {
		return nil, false
	}
	return o.TradingStatus, true
}

// HasTradingStatus returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInner) HasTradingStatus() bool {
	if o != nil && !common.IsNil(o.TradingStatus) {
		return true
	}

	return false
}

// SetTradingStatus gets a reference to the given string and assigns it to the TradingStatus field.
func (o *GetMarketDetailResponseMarketsInner) SetTradingStatus(v string) {
	o.TradingStatus = &v
}

// GetTradeVolume returns the TradeVolume field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInner) GetTradeVolume() string {
	if o == nil || common.IsNil(o.TradeVolume) {
		var ret string
		return ret
	}
	return *o.TradeVolume
}

// GetTradeVolumeOk returns a tuple with the TradeVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInner) GetTradeVolumeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TradeVolume) {
		return nil, false
	}
	return o.TradeVolume, true
}

// HasTradeVolume returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInner) HasTradeVolume() bool {
	if o != nil && !common.IsNil(o.TradeVolume) {
		return true
	}

	return false
}

// SetTradeVolume gets a reference to the given string and assigns it to the TradeVolume field.
func (o *GetMarketDetailResponseMarketsInner) SetTradeVolume(v string) {
	o.TradeVolume = &v
}

// GetLiquidity returns the Liquidity field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInner) GetLiquidity() string {
	if o == nil || common.IsNil(o.Liquidity) {
		var ret string
		return ret
	}
	return *o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInner) GetLiquidityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Liquidity) {
		return nil, false
	}
	return o.Liquidity, true
}

// HasLiquidity returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInner) HasLiquidity() bool {
	if o != nil && !common.IsNil(o.Liquidity) {
		return true
	}

	return false
}

// SetLiquidity gets a reference to the given string and assigns it to the Liquidity field.
func (o *GetMarketDetailResponseMarketsInner) SetLiquidity(v string) {
	o.Liquidity = &v
}

// GetDecimalPrecision returns the DecimalPrecision field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInner) GetDecimalPrecision() int32 {
	if o == nil || common.IsNil(o.DecimalPrecision) {
		var ret int32
		return ret
	}
	return *o.DecimalPrecision
}

// GetDecimalPrecisionOk returns a tuple with the DecimalPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInner) GetDecimalPrecisionOk() (*int32, bool) {
	if o == nil || common.IsNil(o.DecimalPrecision) {
		return nil, false
	}
	return o.DecimalPrecision, true
}

// HasDecimalPrecision returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInner) HasDecimalPrecision() bool {
	if o != nil && !common.IsNil(o.DecimalPrecision) {
		return true
	}

	return false
}

// SetDecimalPrecision gets a reference to the given int32 and assigns it to the DecimalPrecision field.
func (o *GetMarketDetailResponseMarketsInner) SetDecimalPrecision(v int32) {
	o.DecimalPrecision = &v
}

// GetOutcomes returns the Outcomes field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInner) GetOutcomes() []GetMarketDetailResponseMarketsInnerOutcomesInner {
	if o == nil || common.IsNil(o.Outcomes) {
		var ret []GetMarketDetailResponseMarketsInnerOutcomesInner
		return ret
	}
	return o.Outcomes
}

// GetOutcomesOk returns a tuple with the Outcomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInner) GetOutcomesOk() ([]GetMarketDetailResponseMarketsInnerOutcomesInner, bool) {
	if o == nil || common.IsNil(o.Outcomes) {
		return nil, false
	}
	return o.Outcomes, true
}

// HasOutcomes returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInner) HasOutcomes() bool {
	if o != nil && !common.IsNil(o.Outcomes) {
		return true
	}

	return false
}

// SetOutcomes gets a reference to the given []GetMarketDetailResponseMarketsInnerOutcomesInner and assigns it to the Outcomes field.
func (o *GetMarketDetailResponseMarketsInner) SetOutcomes(v []GetMarketDetailResponseMarketsInnerOutcomesInner) {
	o.Outcomes = v
}

func (o GetMarketDetailResponseMarketsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketDetailResponseMarketsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MarketId) {
		toSerialize["marketId"] = o.MarketId
	}
	if !common.IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !common.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !common.IsNil(o.Question) {
		toSerialize["question"] = o.Question
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.ConditionId) {
		toSerialize["conditionId"] = o.ConditionId
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TradingStatus) {
		toSerialize["tradingStatus"] = o.TradingStatus
	}
	if !common.IsNil(o.TradeVolume) {
		toSerialize["tradeVolume"] = o.TradeVolume
	}
	if !common.IsNil(o.Liquidity) {
		toSerialize["liquidity"] = o.Liquidity
	}
	if !common.IsNil(o.DecimalPrecision) {
		toSerialize["decimalPrecision"] = o.DecimalPrecision
	}
	if !common.IsNil(o.Outcomes) {
		toSerialize["outcomes"] = o.Outcomes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMarketDetailResponseMarketsInner) UnmarshalJSON(data []byte) (err error) {
	varGetMarketDetailResponseMarketsInner := _GetMarketDetailResponseMarketsInner{}

	err = json.Unmarshal(data, &varGetMarketDetailResponseMarketsInner)

	if err != nil {
		return err
	}

	*o = GetMarketDetailResponseMarketsInner(varGetMarketDetailResponseMarketsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "marketId")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "title")
		delete(additionalProperties, "question")
		delete(additionalProperties, "description")
		delete(additionalProperties, "conditionId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tradingStatus")
		delete(additionalProperties, "tradeVolume")
		delete(additionalProperties, "liquidity")
		delete(additionalProperties, "decimalPrecision")
		delete(additionalProperties, "outcomes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMarketDetailResponseMarketsInner struct {
	value *GetMarketDetailResponseMarketsInner
	isSet bool
}

func (v NullableGetMarketDetailResponseMarketsInner) Get() *GetMarketDetailResponseMarketsInner {
	return v.value
}

func (v *NullableGetMarketDetailResponseMarketsInner) Set(val *GetMarketDetailResponseMarketsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketDetailResponseMarketsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketDetailResponseMarketsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketDetailResponseMarketsInner(val *GetMarketDetailResponseMarketsInner) *NullableGetMarketDetailResponseMarketsInner {
	return &NullableGetMarketDetailResponseMarketsInner{value: val, isSet: true}
}

func (v NullableGetMarketDetailResponseMarketsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketDetailResponseMarketsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
