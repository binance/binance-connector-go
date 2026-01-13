/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInformationResponse{}

// AccountInformationResponse struct for AccountInformationResponse
type AccountInformationResponse struct {
	Assets               []AccountInformationResponseAssetsInner    `json:"assets,omitempty"`
	Positions            []AccountInformationResponsePositionsInner `json:"positions,omitempty"`
	CanDeposit           *bool                                      `json:"canDeposit,omitempty"`
	CanTrade             *bool                                      `json:"canTrade,omitempty"`
	CanWithdraw          *bool                                      `json:"canWithdraw,omitempty"`
	FeeTier              *int64                                     `json:"feeTier,omitempty"`
	UpdateTime           *int64                                     `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountInformationResponse AccountInformationResponse

// NewAccountInformationResponse instantiates a new AccountInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInformationResponse() *AccountInformationResponse {
	this := AccountInformationResponse{}
	return &this
}

// NewAccountInformationResponseWithDefaults instantiates a new AccountInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInformationResponseWithDefaults() *AccountInformationResponse {
	this := AccountInformationResponse{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetAssets() []AccountInformationResponseAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []AccountInformationResponseAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetAssetsOk() ([]AccountInformationResponseAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []AccountInformationResponseAssetsInner and assigns it to the Assets field.
func (o *AccountInformationResponse) SetAssets(v []AccountInformationResponseAssetsInner) {
	o.Assets = v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetPositions() []AccountInformationResponsePositionsInner {
	if o == nil || common.IsNil(o.Positions) {
		var ret []AccountInformationResponsePositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetPositionsOk() ([]AccountInformationResponsePositionsInner, bool) {
	if o == nil || common.IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasPositions() bool {
	if o != nil && !common.IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []AccountInformationResponsePositionsInner and assigns it to the Positions field.
func (o *AccountInformationResponse) SetPositions(v []AccountInformationResponsePositionsInner) {
	o.Positions = v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetCanDeposit() bool {
	if o == nil || common.IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetCanDepositOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasCanDeposit() bool {
	if o != nil && !common.IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *AccountInformationResponse) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *AccountInformationResponse) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetCanWithdraw() bool {
	if o == nil || common.IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasCanWithdraw() bool {
	if o != nil && !common.IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *AccountInformationResponse) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetFeeTier returns the FeeTier field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetFeeTier() int64 {
	if o == nil || common.IsNil(o.FeeTier) {
		var ret int64
		return ret
	}
	return *o.FeeTier
}

// GetFeeTierOk returns a tuple with the FeeTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetFeeTierOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FeeTier) {
		return nil, false
	}
	return o.FeeTier, true
}

// HasFeeTier returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasFeeTier() bool {
	if o != nil && !common.IsNil(o.FeeTier) {
		return true
	}

	return false
}

// SetFeeTier gets a reference to the given int64 and assigns it to the FeeTier field.
func (o *AccountInformationResponse) SetFeeTier(v int64) {
	o.FeeTier = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountInformationResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o AccountInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !common.IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}
	if !common.IsNil(o.CanDeposit) {
		toSerialize["canDeposit"] = o.CanDeposit
	}
	if !common.IsNil(o.CanTrade) {
		toSerialize["canTrade"] = o.CanTrade
	}
	if !common.IsNil(o.CanWithdraw) {
		toSerialize["canWithdraw"] = o.CanWithdraw
	}
	if !common.IsNil(o.FeeTier) {
		toSerialize["feeTier"] = o.FeeTier
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInformationResponse) UnmarshalJSON(data []byte) (err error) {
	varAccountInformationResponse := _AccountInformationResponse{}

	err = json.Unmarshal(data, &varAccountInformationResponse)

	if err != nil {
		return err
	}

	*o = AccountInformationResponse(varAccountInformationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assets")
		delete(additionalProperties, "positions")
		delete(additionalProperties, "canDeposit")
		delete(additionalProperties, "canTrade")
		delete(additionalProperties, "canWithdraw")
		delete(additionalProperties, "feeTier")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInformationResponse struct {
	value *AccountInformationResponse
	isSet bool
}

func (v NullableAccountInformationResponse) Get() *AccountInformationResponse {
	return v.value
}

func (v *NullableAccountInformationResponse) Set(val *AccountInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInformationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInformationResponse(val *AccountInformationResponse) *NullableAccountInformationResponse {
	return &NullableAccountInformationResponse{value: val, isSet: true}
}

func (v NullableAccountInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
