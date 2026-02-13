/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FlexibleLoanAdjustLtvResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FlexibleLoanAdjustLtvResponse{}

// FlexibleLoanAdjustLtvResponse struct for FlexibleLoanAdjustLtvResponse
type FlexibleLoanAdjustLtvResponse struct {
	LoanCoin             *string `json:"loanCoin,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	Direction            *string `json:"direction,omitempty"`
	AdjustmentAmount     *string `json:"adjustmentAmount,omitempty"`
	CurrentLTV           *string `json:"currentLTV,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FlexibleLoanAdjustLtvResponse FlexibleLoanAdjustLtvResponse

// NewFlexibleLoanAdjustLtvResponse instantiates a new FlexibleLoanAdjustLtvResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexibleLoanAdjustLtvResponse() *FlexibleLoanAdjustLtvResponse {
	this := FlexibleLoanAdjustLtvResponse{}
	return &this
}

// NewFlexibleLoanAdjustLtvResponseWithDefaults instantiates a new FlexibleLoanAdjustLtvResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexibleLoanAdjustLtvResponseWithDefaults() *FlexibleLoanAdjustLtvResponse {
	this := FlexibleLoanAdjustLtvResponse{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *FlexibleLoanAdjustLtvResponse) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanAdjustLtvResponse) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *FlexibleLoanAdjustLtvResponse) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *FlexibleLoanAdjustLtvResponse) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *FlexibleLoanAdjustLtvResponse) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanAdjustLtvResponse) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *FlexibleLoanAdjustLtvResponse) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *FlexibleLoanAdjustLtvResponse) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *FlexibleLoanAdjustLtvResponse) GetDirection() string {
	if o == nil || common.IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanAdjustLtvResponse) GetDirectionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *FlexibleLoanAdjustLtvResponse) HasDirection() bool {
	if o != nil && !common.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *FlexibleLoanAdjustLtvResponse) SetDirection(v string) {
	o.Direction = &v
}

// GetAdjustmentAmount returns the AdjustmentAmount field value if set, zero value otherwise.
func (o *FlexibleLoanAdjustLtvResponse) GetAdjustmentAmount() string {
	if o == nil || common.IsNil(o.AdjustmentAmount) {
		var ret string
		return ret
	}
	return *o.AdjustmentAmount
}

// GetAdjustmentAmountOk returns a tuple with the AdjustmentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanAdjustLtvResponse) GetAdjustmentAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdjustmentAmount) {
		return nil, false
	}
	return o.AdjustmentAmount, true
}

// HasAdjustmentAmount returns a boolean if a field has been set.
func (o *FlexibleLoanAdjustLtvResponse) HasAdjustmentAmount() bool {
	if o != nil && !common.IsNil(o.AdjustmentAmount) {
		return true
	}

	return false
}

// SetAdjustmentAmount gets a reference to the given string and assigns it to the AdjustmentAmount field.
func (o *FlexibleLoanAdjustLtvResponse) SetAdjustmentAmount(v string) {
	o.AdjustmentAmount = &v
}

// GetCurrentLTV returns the CurrentLTV field value if set, zero value otherwise.
func (o *FlexibleLoanAdjustLtvResponse) GetCurrentLTV() string {
	if o == nil || common.IsNil(o.CurrentLTV) {
		var ret string
		return ret
	}
	return *o.CurrentLTV
}

// GetCurrentLTVOk returns a tuple with the CurrentLTV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanAdjustLtvResponse) GetCurrentLTVOk() (*string, bool) {
	if o == nil || common.IsNil(o.CurrentLTV) {
		return nil, false
	}
	return o.CurrentLTV, true
}

// HasCurrentLTV returns a boolean if a field has been set.
func (o *FlexibleLoanAdjustLtvResponse) HasCurrentLTV() bool {
	if o != nil && !common.IsNil(o.CurrentLTV) {
		return true
	}

	return false
}

// SetCurrentLTV gets a reference to the given string and assigns it to the CurrentLTV field.
func (o *FlexibleLoanAdjustLtvResponse) SetCurrentLTV(v string) {
	o.CurrentLTV = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FlexibleLoanAdjustLtvResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanAdjustLtvResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FlexibleLoanAdjustLtvResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FlexibleLoanAdjustLtvResponse) SetStatus(v string) {
	o.Status = &v
}

func (o FlexibleLoanAdjustLtvResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlexibleLoanAdjustLtvResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !common.IsNil(o.AdjustmentAmount) {
		toSerialize["adjustmentAmount"] = o.AdjustmentAmount
	}
	if !common.IsNil(o.CurrentLTV) {
		toSerialize["currentLTV"] = o.CurrentLTV
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FlexibleLoanAdjustLtvResponse) UnmarshalJSON(data []byte) (err error) {
	varFlexibleLoanAdjustLtvResponse := _FlexibleLoanAdjustLtvResponse{}

	err = json.Unmarshal(data, &varFlexibleLoanAdjustLtvResponse)

	if err != nil {
		return err
	}

	*o = FlexibleLoanAdjustLtvResponse(varFlexibleLoanAdjustLtvResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "adjustmentAmount")
		delete(additionalProperties, "currentLTV")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFlexibleLoanAdjustLtvResponse struct {
	value *FlexibleLoanAdjustLtvResponse
	isSet bool
}

func (v NullableFlexibleLoanAdjustLtvResponse) Get() *FlexibleLoanAdjustLtvResponse {
	return v.value
}

func (v *NullableFlexibleLoanAdjustLtvResponse) Set(val *FlexibleLoanAdjustLtvResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexibleLoanAdjustLtvResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexibleLoanAdjustLtvResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexibleLoanAdjustLtvResponse(val *FlexibleLoanAdjustLtvResponse) *NullableFlexibleLoanAdjustLtvResponse {
	return &NullableFlexibleLoanAdjustLtvResponse{value: val, isSet: true}
}

func (v NullableFlexibleLoanAdjustLtvResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexibleLoanAdjustLtvResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
