/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OptionMarginAccountInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OptionMarginAccountInformationResponse{}

// OptionMarginAccountInformationResponse struct for OptionMarginAccountInformationResponse
type OptionMarginAccountInformationResponse struct {
	Asset                []OptionMarginAccountInformationResponseAssetInner `json:"asset,omitempty"`
	Greek                []OptionMarginAccountInformationResponseGreekInner `json:"greek,omitempty"`
	Time                 *int64                                             `json:"time,omitempty"`
	CanTrade             *bool                                              `json:"canTrade,omitempty"`
	CanDeposit           *bool                                              `json:"canDeposit,omitempty"`
	CanWithdraw          *bool                                              `json:"canWithdraw,omitempty"`
	ReduceOnly           *bool                                              `json:"reduceOnly,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OptionMarginAccountInformationResponse OptionMarginAccountInformationResponse

// NewOptionMarginAccountInformationResponse instantiates a new OptionMarginAccountInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionMarginAccountInformationResponse() *OptionMarginAccountInformationResponse {
	this := OptionMarginAccountInformationResponse{}
	return &this
}

// NewOptionMarginAccountInformationResponseWithDefaults instantiates a new OptionMarginAccountInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionMarginAccountInformationResponseWithDefaults() *OptionMarginAccountInformationResponse {
	this := OptionMarginAccountInformationResponse{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponse) GetAsset() []OptionMarginAccountInformationResponseAssetInner {
	if o == nil || common.IsNil(o.Asset) {
		var ret []OptionMarginAccountInformationResponseAssetInner
		return ret
	}
	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponse) GetAssetOk() ([]OptionMarginAccountInformationResponseAssetInner, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponse) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given []OptionMarginAccountInformationResponseAssetInner and assigns it to the Asset field.
func (o *OptionMarginAccountInformationResponse) SetAsset(v []OptionMarginAccountInformationResponseAssetInner) {
	o.Asset = v
}

// GetGreek returns the Greek field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponse) GetGreek() []OptionMarginAccountInformationResponseGreekInner {
	if o == nil || common.IsNil(o.Greek) {
		var ret []OptionMarginAccountInformationResponseGreekInner
		return ret
	}
	return o.Greek
}

// GetGreekOk returns a tuple with the Greek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponse) GetGreekOk() ([]OptionMarginAccountInformationResponseGreekInner, bool) {
	if o == nil || common.IsNil(o.Greek) {
		return nil, false
	}
	return o.Greek, true
}

// HasGreek returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponse) HasGreek() bool {
	if o != nil && !common.IsNil(o.Greek) {
		return true
	}

	return false
}

// SetGreek gets a reference to the given []OptionMarginAccountInformationResponseGreekInner and assigns it to the Greek field.
func (o *OptionMarginAccountInformationResponse) SetGreek(v []OptionMarginAccountInformationResponseGreekInner) {
	o.Greek = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponse) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponse) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponse) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *OptionMarginAccountInformationResponse) SetTime(v int64) {
	o.Time = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponse) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponse) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponse) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *OptionMarginAccountInformationResponse) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponse) GetCanDeposit() bool {
	if o == nil || common.IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponse) GetCanDepositOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponse) HasCanDeposit() bool {
	if o != nil && !common.IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *OptionMarginAccountInformationResponse) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponse) GetCanWithdraw() bool {
	if o == nil || common.IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponse) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponse) HasCanWithdraw() bool {
	if o != nil && !common.IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *OptionMarginAccountInformationResponse) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponse) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponse) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponse) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *OptionMarginAccountInformationResponse) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

func (o OptionMarginAccountInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionMarginAccountInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Greek) {
		toSerialize["greek"] = o.Greek
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
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
	if !common.IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OptionMarginAccountInformationResponse) UnmarshalJSON(data []byte) (err error) {
	varOptionMarginAccountInformationResponse := _OptionMarginAccountInformationResponse{}

	err = json.Unmarshal(data, &varOptionMarginAccountInformationResponse)

	if err != nil {
		return err
	}

	*o = OptionMarginAccountInformationResponse(varOptionMarginAccountInformationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "greek")
		delete(additionalProperties, "time")
		delete(additionalProperties, "canTrade")
		delete(additionalProperties, "canDeposit")
		delete(additionalProperties, "canWithdraw")
		delete(additionalProperties, "reduceOnly")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOptionMarginAccountInformationResponse struct {
	value *OptionMarginAccountInformationResponse
	isSet bool
}

func (v NullableOptionMarginAccountInformationResponse) Get() *OptionMarginAccountInformationResponse {
	return v.value
}

func (v *NullableOptionMarginAccountInformationResponse) Set(val *OptionMarginAccountInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionMarginAccountInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionMarginAccountInformationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionMarginAccountInformationResponse(val *OptionMarginAccountInformationResponse) *NullableOptionMarginAccountInformationResponse {
	return &NullableOptionMarginAccountInformationResponse{value: val, isSet: true}
}

func (v NullableOptionMarginAccountInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionMarginAccountInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
