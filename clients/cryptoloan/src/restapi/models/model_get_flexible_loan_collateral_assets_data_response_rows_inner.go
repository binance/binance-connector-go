/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleLoanCollateralAssetsDataResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanCollateralAssetsDataResponseRowsInner{}

// GetFlexibleLoanCollateralAssetsDataResponseRowsInner struct for GetFlexibleLoanCollateralAssetsDataResponseRowsInner
type GetFlexibleLoanCollateralAssetsDataResponseRowsInner struct {
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	InitialLTV           *string `json:"initialLTV,omitempty"`
	MarginCallLTV        *string `json:"marginCallLTV,omitempty"`
	LiquidationLTV       *string `json:"liquidationLTV,omitempty"`
	MaxLimit             *string `json:"maxLimit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanCollateralAssetsDataResponseRowsInner GetFlexibleLoanCollateralAssetsDataResponseRowsInner

// NewGetFlexibleLoanCollateralAssetsDataResponseRowsInner instantiates a new GetFlexibleLoanCollateralAssetsDataResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanCollateralAssetsDataResponseRowsInner() *GetFlexibleLoanCollateralAssetsDataResponseRowsInner {
	this := GetFlexibleLoanCollateralAssetsDataResponseRowsInner{}
	return &this
}

// NewGetFlexibleLoanCollateralAssetsDataResponseRowsInnerWithDefaults instantiates a new GetFlexibleLoanCollateralAssetsDataResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanCollateralAssetsDataResponseRowsInnerWithDefaults() *GetFlexibleLoanCollateralAssetsDataResponseRowsInner {
	this := GetFlexibleLoanCollateralAssetsDataResponseRowsInner{}
	return &this
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetInitialLTV returns the InitialLTV field value if set, zero value otherwise.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) GetInitialLTV() string {
	if o == nil || common.IsNil(o.InitialLTV) {
		var ret string
		return ret
	}
	return *o.InitialLTV
}

// GetInitialLTVOk returns a tuple with the InitialLTV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) GetInitialLTVOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialLTV) {
		return nil, false
	}
	return o.InitialLTV, true
}

// HasInitialLTV returns a boolean if a field has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) HasInitialLTV() bool {
	if o != nil && !common.IsNil(o.InitialLTV) {
		return true
	}

	return false
}

// SetInitialLTV gets a reference to the given string and assigns it to the InitialLTV field.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) SetInitialLTV(v string) {
	o.InitialLTV = &v
}

// GetMarginCallLTV returns the MarginCallLTV field value if set, zero value otherwise.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) GetMarginCallLTV() string {
	if o == nil || common.IsNil(o.MarginCallLTV) {
		var ret string
		return ret
	}
	return *o.MarginCallLTV
}

// GetMarginCallLTVOk returns a tuple with the MarginCallLTV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) GetMarginCallLTVOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginCallLTV) {
		return nil, false
	}
	return o.MarginCallLTV, true
}

// HasMarginCallLTV returns a boolean if a field has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) HasMarginCallLTV() bool {
	if o != nil && !common.IsNil(o.MarginCallLTV) {
		return true
	}

	return false
}

// SetMarginCallLTV gets a reference to the given string and assigns it to the MarginCallLTV field.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) SetMarginCallLTV(v string) {
	o.MarginCallLTV = &v
}

// GetLiquidationLTV returns the LiquidationLTV field value if set, zero value otherwise.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) GetLiquidationLTV() string {
	if o == nil || common.IsNil(o.LiquidationLTV) {
		var ret string
		return ret
	}
	return *o.LiquidationLTV
}

// GetLiquidationLTVOk returns a tuple with the LiquidationLTV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) GetLiquidationLTVOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationLTV) {
		return nil, false
	}
	return o.LiquidationLTV, true
}

// HasLiquidationLTV returns a boolean if a field has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) HasLiquidationLTV() bool {
	if o != nil && !common.IsNil(o.LiquidationLTV) {
		return true
	}

	return false
}

// SetLiquidationLTV gets a reference to the given string and assigns it to the LiquidationLTV field.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) SetLiquidationLTV(v string) {
	o.LiquidationLTV = &v
}

// GetMaxLimit returns the MaxLimit field value if set, zero value otherwise.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) GetMaxLimit() string {
	if o == nil || common.IsNil(o.MaxLimit) {
		var ret string
		return ret
	}
	return *o.MaxLimit
}

// GetMaxLimitOk returns a tuple with the MaxLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) GetMaxLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxLimit) {
		return nil, false
	}
	return o.MaxLimit, true
}

// HasMaxLimit returns a boolean if a field has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) HasMaxLimit() bool {
	if o != nil && !common.IsNil(o.MaxLimit) {
		return true
	}

	return false
}

// SetMaxLimit gets a reference to the given string and assigns it to the MaxLimit field.
func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) SetMaxLimit(v string) {
	o.MaxLimit = &v
}

func (o GetFlexibleLoanCollateralAssetsDataResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanCollateralAssetsDataResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.InitialLTV) {
		toSerialize["initialLTV"] = o.InitialLTV
	}
	if !common.IsNil(o.MarginCallLTV) {
		toSerialize["marginCallLTV"] = o.MarginCallLTV
	}
	if !common.IsNil(o.LiquidationLTV) {
		toSerialize["liquidationLTV"] = o.LiquidationLTV
	}
	if !common.IsNil(o.MaxLimit) {
		toSerialize["maxLimit"] = o.MaxLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanCollateralAssetsDataResponseRowsInner := _GetFlexibleLoanCollateralAssetsDataResponseRowsInner{}

	err = json.Unmarshal(data, &varGetFlexibleLoanCollateralAssetsDataResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanCollateralAssetsDataResponseRowsInner(varGetFlexibleLoanCollateralAssetsDataResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "initialLTV")
		delete(additionalProperties, "marginCallLTV")
		delete(additionalProperties, "liquidationLTV")
		delete(additionalProperties, "maxLimit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanCollateralAssetsDataResponseRowsInner struct {
	value *GetFlexibleLoanCollateralAssetsDataResponseRowsInner
	isSet bool
}

func (v NullableGetFlexibleLoanCollateralAssetsDataResponseRowsInner) Get() *GetFlexibleLoanCollateralAssetsDataResponseRowsInner {
	return v.value
}

func (v *NullableGetFlexibleLoanCollateralAssetsDataResponseRowsInner) Set(val *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanCollateralAssetsDataResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanCollateralAssetsDataResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanCollateralAssetsDataResponseRowsInner(val *GetFlexibleLoanCollateralAssetsDataResponseRowsInner) *NullableGetFlexibleLoanCollateralAssetsDataResponseRowsInner {
	return &NullableGetFlexibleLoanCollateralAssetsDataResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanCollateralAssetsDataResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanCollateralAssetsDataResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
