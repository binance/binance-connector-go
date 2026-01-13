/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FuturesAccountConfigurationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesAccountConfigurationResponse{}

// FuturesAccountConfigurationResponse struct for FuturesAccountConfigurationResponse
type FuturesAccountConfigurationResponse struct {
	FeeTier              *int64 `json:"feeTier,omitempty"`
	CanTrade             *bool  `json:"canTrade,omitempty"`
	CanDeposit           *bool  `json:"canDeposit,omitempty"`
	CanWithdraw          *bool  `json:"canWithdraw,omitempty"`
	DualSidePosition     *bool  `json:"dualSidePosition,omitempty"`
	UpdateTime           *int64 `json:"updateTime,omitempty"`
	MultiAssetsMargin    *bool  `json:"multiAssetsMargin,omitempty"`
	TradeGroupId         *int64 `json:"tradeGroupId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FuturesAccountConfigurationResponse FuturesAccountConfigurationResponse

// NewFuturesAccountConfigurationResponse instantiates a new FuturesAccountConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesAccountConfigurationResponse() *FuturesAccountConfigurationResponse {
	this := FuturesAccountConfigurationResponse{}
	return &this
}

// NewFuturesAccountConfigurationResponseWithDefaults instantiates a new FuturesAccountConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesAccountConfigurationResponseWithDefaults() *FuturesAccountConfigurationResponse {
	this := FuturesAccountConfigurationResponse{}
	return &this
}

// GetFeeTier returns the FeeTier field value if set, zero value otherwise.
func (o *FuturesAccountConfigurationResponse) GetFeeTier() int64 {
	if o == nil || common.IsNil(o.FeeTier) {
		var ret int64
		return ret
	}
	return *o.FeeTier
}

// GetFeeTierOk returns a tuple with the FeeTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountConfigurationResponse) GetFeeTierOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FeeTier) {
		return nil, false
	}
	return o.FeeTier, true
}

// HasFeeTier returns a boolean if a field has been set.
func (o *FuturesAccountConfigurationResponse) HasFeeTier() bool {
	if o != nil && !common.IsNil(o.FeeTier) {
		return true
	}

	return false
}

// SetFeeTier gets a reference to the given int64 and assigns it to the FeeTier field.
func (o *FuturesAccountConfigurationResponse) SetFeeTier(v int64) {
	o.FeeTier = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *FuturesAccountConfigurationResponse) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountConfigurationResponse) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *FuturesAccountConfigurationResponse) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *FuturesAccountConfigurationResponse) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *FuturesAccountConfigurationResponse) GetCanDeposit() bool {
	if o == nil || common.IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountConfigurationResponse) GetCanDepositOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *FuturesAccountConfigurationResponse) HasCanDeposit() bool {
	if o != nil && !common.IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *FuturesAccountConfigurationResponse) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *FuturesAccountConfigurationResponse) GetCanWithdraw() bool {
	if o == nil || common.IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountConfigurationResponse) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *FuturesAccountConfigurationResponse) HasCanWithdraw() bool {
	if o != nil && !common.IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *FuturesAccountConfigurationResponse) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetDualSidePosition returns the DualSidePosition field value if set, zero value otherwise.
func (o *FuturesAccountConfigurationResponse) GetDualSidePosition() bool {
	if o == nil || common.IsNil(o.DualSidePosition) {
		var ret bool
		return ret
	}
	return *o.DualSidePosition
}

// GetDualSidePositionOk returns a tuple with the DualSidePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountConfigurationResponse) GetDualSidePositionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DualSidePosition) {
		return nil, false
	}
	return o.DualSidePosition, true
}

// HasDualSidePosition returns a boolean if a field has been set.
func (o *FuturesAccountConfigurationResponse) HasDualSidePosition() bool {
	if o != nil && !common.IsNil(o.DualSidePosition) {
		return true
	}

	return false
}

// SetDualSidePosition gets a reference to the given bool and assigns it to the DualSidePosition field.
func (o *FuturesAccountConfigurationResponse) SetDualSidePosition(v bool) {
	o.DualSidePosition = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *FuturesAccountConfigurationResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountConfigurationResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *FuturesAccountConfigurationResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *FuturesAccountConfigurationResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetMultiAssetsMargin returns the MultiAssetsMargin field value if set, zero value otherwise.
func (o *FuturesAccountConfigurationResponse) GetMultiAssetsMargin() bool {
	if o == nil || common.IsNil(o.MultiAssetsMargin) {
		var ret bool
		return ret
	}
	return *o.MultiAssetsMargin
}

// GetMultiAssetsMarginOk returns a tuple with the MultiAssetsMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountConfigurationResponse) GetMultiAssetsMarginOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MultiAssetsMargin) {
		return nil, false
	}
	return o.MultiAssetsMargin, true
}

// HasMultiAssetsMargin returns a boolean if a field has been set.
func (o *FuturesAccountConfigurationResponse) HasMultiAssetsMargin() bool {
	if o != nil && !common.IsNil(o.MultiAssetsMargin) {
		return true
	}

	return false
}

// SetMultiAssetsMargin gets a reference to the given bool and assigns it to the MultiAssetsMargin field.
func (o *FuturesAccountConfigurationResponse) SetMultiAssetsMargin(v bool) {
	o.MultiAssetsMargin = &v
}

// GetTradeGroupId returns the TradeGroupId field value if set, zero value otherwise.
func (o *FuturesAccountConfigurationResponse) GetTradeGroupId() int64 {
	if o == nil || common.IsNil(o.TradeGroupId) {
		var ret int64
		return ret
	}
	return *o.TradeGroupId
}

// GetTradeGroupIdOk returns a tuple with the TradeGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountConfigurationResponse) GetTradeGroupIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeGroupId) {
		return nil, false
	}
	return o.TradeGroupId, true
}

// HasTradeGroupId returns a boolean if a field has been set.
func (o *FuturesAccountConfigurationResponse) HasTradeGroupId() bool {
	if o != nil && !common.IsNil(o.TradeGroupId) {
		return true
	}

	return false
}

// SetTradeGroupId gets a reference to the given int64 and assigns it to the TradeGroupId field.
func (o *FuturesAccountConfigurationResponse) SetTradeGroupId(v int64) {
	o.TradeGroupId = &v
}

func (o FuturesAccountConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesAccountConfigurationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FeeTier) {
		toSerialize["feeTier"] = o.FeeTier
	}
	if !common.IsNil(o.CanTrade) {
		toSerialize["canTrade"] = o.CanTrade
	}
	if !common.IsNil(o.CanDeposit) {
		toSerialize["canDeposit"] = o.CanDeposit
	}
	if !common.IsNil(o.CanWithdraw) {
		toSerialize["canWithdraw"] = o.CanWithdraw
	}
	if !common.IsNil(o.DualSidePosition) {
		toSerialize["dualSidePosition"] = o.DualSidePosition
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.MultiAssetsMargin) {
		toSerialize["multiAssetsMargin"] = o.MultiAssetsMargin
	}
	if !common.IsNil(o.TradeGroupId) {
		toSerialize["tradeGroupId"] = o.TradeGroupId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FuturesAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	varFuturesAccountConfigurationResponse := _FuturesAccountConfigurationResponse{}

	err = json.Unmarshal(data, &varFuturesAccountConfigurationResponse)

	if err != nil {
		return err
	}

	*o = FuturesAccountConfigurationResponse(varFuturesAccountConfigurationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feeTier")
		delete(additionalProperties, "canTrade")
		delete(additionalProperties, "canDeposit")
		delete(additionalProperties, "canWithdraw")
		delete(additionalProperties, "dualSidePosition")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "multiAssetsMargin")
		delete(additionalProperties, "tradeGroupId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFuturesAccountConfigurationResponse struct {
	value *FuturesAccountConfigurationResponse
	isSet bool
}

func (v NullableFuturesAccountConfigurationResponse) Get() *FuturesAccountConfigurationResponse {
	return v.value
}

func (v *NullableFuturesAccountConfigurationResponse) Set(val *FuturesAccountConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesAccountConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesAccountConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuturesAccountConfigurationResponse(val *FuturesAccountConfigurationResponse) *NullableFuturesAccountConfigurationResponse {
	return &NullableFuturesAccountConfigurationResponse{value: val, isSet: true}
}

func (v NullableFuturesAccountConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesAccountConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
