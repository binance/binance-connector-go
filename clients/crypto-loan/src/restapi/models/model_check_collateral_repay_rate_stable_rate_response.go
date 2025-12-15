/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CheckCollateralRepayRateStableRateResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckCollateralRepayRateStableRateResponse{}

// CheckCollateralRepayRateStableRateResponse struct for CheckCollateralRepayRateStableRateResponse
type CheckCollateralRepayRateStableRateResponse struct {
	LoanlCoin            *string `json:"loanlCoin,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	RepayAmount          *string `json:"repayAmount,omitempty"`
	Rate                 *string `json:"rate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CheckCollateralRepayRateStableRateResponse CheckCollateralRepayRateStableRateResponse

// NewCheckCollateralRepayRateStableRateResponse instantiates a new CheckCollateralRepayRateStableRateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckCollateralRepayRateStableRateResponse() *CheckCollateralRepayRateStableRateResponse {
	this := CheckCollateralRepayRateStableRateResponse{}
	return &this
}

// NewCheckCollateralRepayRateStableRateResponseWithDefaults instantiates a new CheckCollateralRepayRateStableRateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckCollateralRepayRateStableRateResponseWithDefaults() *CheckCollateralRepayRateStableRateResponse {
	this := CheckCollateralRepayRateStableRateResponse{}
	return &this
}

// GetLoanlCoin returns the LoanlCoin field value if set, zero value otherwise.
func (o *CheckCollateralRepayRateStableRateResponse) GetLoanlCoin() string {
	if o == nil || common.IsNil(o.LoanlCoin) {
		var ret string
		return ret
	}
	return *o.LoanlCoin
}

// GetLoanlCoinOk returns a tuple with the LoanlCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCollateralRepayRateStableRateResponse) GetLoanlCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanlCoin) {
		return nil, false
	}
	return o.LoanlCoin, true
}

// HasLoanlCoin returns a boolean if a field has been set.
func (o *CheckCollateralRepayRateStableRateResponse) HasLoanlCoin() bool {
	if o != nil && !common.IsNil(o.LoanlCoin) {
		return true
	}

	return false
}

// SetLoanlCoin gets a reference to the given string and assigns it to the LoanlCoin field.
func (o *CheckCollateralRepayRateStableRateResponse) SetLoanlCoin(v string) {
	o.LoanlCoin = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *CheckCollateralRepayRateStableRateResponse) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCollateralRepayRateStableRateResponse) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *CheckCollateralRepayRateStableRateResponse) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *CheckCollateralRepayRateStableRateResponse) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetRepayAmount returns the RepayAmount field value if set, zero value otherwise.
func (o *CheckCollateralRepayRateStableRateResponse) GetRepayAmount() string {
	if o == nil || common.IsNil(o.RepayAmount) {
		var ret string
		return ret
	}
	return *o.RepayAmount
}

// GetRepayAmountOk returns a tuple with the RepayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCollateralRepayRateStableRateResponse) GetRepayAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepayAmount) {
		return nil, false
	}
	return o.RepayAmount, true
}

// HasRepayAmount returns a boolean if a field has been set.
func (o *CheckCollateralRepayRateStableRateResponse) HasRepayAmount() bool {
	if o != nil && !common.IsNil(o.RepayAmount) {
		return true
	}

	return false
}

// SetRepayAmount gets a reference to the given string and assigns it to the RepayAmount field.
func (o *CheckCollateralRepayRateStableRateResponse) SetRepayAmount(v string) {
	o.RepayAmount = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *CheckCollateralRepayRateStableRateResponse) GetRate() string {
	if o == nil || common.IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCollateralRepayRateStableRateResponse) GetRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *CheckCollateralRepayRateStableRateResponse) HasRate() bool {
	if o != nil && !common.IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *CheckCollateralRepayRateStableRateResponse) SetRate(v string) {
	o.Rate = &v
}

func (o CheckCollateralRepayRateStableRateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckCollateralRepayRateStableRateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanlCoin) {
		toSerialize["loanlCoin"] = o.LoanlCoin
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.RepayAmount) {
		toSerialize["repayAmount"] = o.RepayAmount
	}
	if !common.IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckCollateralRepayRateStableRateResponse) UnmarshalJSON(data []byte) (err error) {
	varCheckCollateralRepayRateStableRateResponse := _CheckCollateralRepayRateStableRateResponse{}

	err = json.Unmarshal(data, &varCheckCollateralRepayRateStableRateResponse)

	if err != nil {
		return err
	}

	*o = CheckCollateralRepayRateStableRateResponse(varCheckCollateralRepayRateStableRateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanlCoin")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "repayAmount")
		delete(additionalProperties, "rate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCheckCollateralRepayRateStableRateResponse struct {
	value *CheckCollateralRepayRateStableRateResponse
	isSet bool
}

func (v NullableCheckCollateralRepayRateStableRateResponse) Get() *CheckCollateralRepayRateStableRateResponse {
	return v.value
}

func (v *NullableCheckCollateralRepayRateStableRateResponse) Set(val *CheckCollateralRepayRateStableRateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckCollateralRepayRateStableRateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckCollateralRepayRateStableRateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckCollateralRepayRateStableRateResponse(val *CheckCollateralRepayRateStableRateResponse) *NullableCheckCollateralRepayRateStableRateResponse {
	return &NullableCheckCollateralRepayRateStableRateResponse{value: val, isSet: true}
}

func (v NullableCheckCollateralRepayRateStableRateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckCollateralRepayRateStableRateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
