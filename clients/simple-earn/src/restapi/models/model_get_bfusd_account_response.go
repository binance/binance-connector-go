/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetBfusdAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdAccountResponse{}

// GetBfusdAccountResponse struct for GetBfusdAccountResponse
type GetBfusdAccountResponse struct {
	BfusdAmount          *string `json:"bfusdAmount,omitempty"`
	UsdtProfit           *string `json:"usdtProfit,omitempty"`
	BfusdProfit          *string `json:"bfusdProfit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBfusdAccountResponse GetBfusdAccountResponse

// NewGetBfusdAccountResponse instantiates a new GetBfusdAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdAccountResponse() *GetBfusdAccountResponse {
	this := GetBfusdAccountResponse{}
	return &this
}

// NewGetBfusdAccountResponseWithDefaults instantiates a new GetBfusdAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdAccountResponseWithDefaults() *GetBfusdAccountResponse {
	this := GetBfusdAccountResponse{}
	return &this
}

// GetBfusdAmount returns the BfusdAmount field value if set, zero value otherwise.
func (o *GetBfusdAccountResponse) GetBfusdAmount() string {
	if o == nil || common.IsNil(o.BfusdAmount) {
		var ret string
		return ret
	}
	return *o.BfusdAmount
}

// GetBfusdAmountOk returns a tuple with the BfusdAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdAccountResponse) GetBfusdAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.BfusdAmount) {
		return nil, false
	}
	return o.BfusdAmount, true
}

// HasBfusdAmount returns a boolean if a field has been set.
func (o *GetBfusdAccountResponse) HasBfusdAmount() bool {
	if o != nil && !common.IsNil(o.BfusdAmount) {
		return true
	}

	return false
}

// SetBfusdAmount gets a reference to the given string and assigns it to the BfusdAmount field.
func (o *GetBfusdAccountResponse) SetBfusdAmount(v string) {
	o.BfusdAmount = &v
}

// GetUsdtProfit returns the UsdtProfit field value if set, zero value otherwise.
func (o *GetBfusdAccountResponse) GetUsdtProfit() string {
	if o == nil || common.IsNil(o.UsdtProfit) {
		var ret string
		return ret
	}
	return *o.UsdtProfit
}

// GetUsdtProfitOk returns a tuple with the UsdtProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdAccountResponse) GetUsdtProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UsdtProfit) {
		return nil, false
	}
	return o.UsdtProfit, true
}

// HasUsdtProfit returns a boolean if a field has been set.
func (o *GetBfusdAccountResponse) HasUsdtProfit() bool {
	if o != nil && !common.IsNil(o.UsdtProfit) {
		return true
	}

	return false
}

// SetUsdtProfit gets a reference to the given string and assigns it to the UsdtProfit field.
func (o *GetBfusdAccountResponse) SetUsdtProfit(v string) {
	o.UsdtProfit = &v
}

// GetBfusdProfit returns the BfusdProfit field value if set, zero value otherwise.
func (o *GetBfusdAccountResponse) GetBfusdProfit() string {
	if o == nil || common.IsNil(o.BfusdProfit) {
		var ret string
		return ret
	}
	return *o.BfusdProfit
}

// GetBfusdProfitOk returns a tuple with the BfusdProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdAccountResponse) GetBfusdProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.BfusdProfit) {
		return nil, false
	}
	return o.BfusdProfit, true
}

// HasBfusdProfit returns a boolean if a field has been set.
func (o *GetBfusdAccountResponse) HasBfusdProfit() bool {
	if o != nil && !common.IsNil(o.BfusdProfit) {
		return true
	}

	return false
}

// SetBfusdProfit gets a reference to the given string and assigns it to the BfusdProfit field.
func (o *GetBfusdAccountResponse) SetBfusdProfit(v string) {
	o.BfusdProfit = &v
}

func (o GetBfusdAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BfusdAmount) {
		toSerialize["bfusdAmount"] = o.BfusdAmount
	}
	if !common.IsNil(o.UsdtProfit) {
		toSerialize["usdtProfit"] = o.UsdtProfit
	}
	if !common.IsNil(o.BfusdProfit) {
		toSerialize["bfusdProfit"] = o.BfusdProfit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBfusdAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdAccountResponse := _GetBfusdAccountResponse{}

	err = json.Unmarshal(data, &varGetBfusdAccountResponse)

	if err != nil {
		return err
	}

	*o = GetBfusdAccountResponse(varGetBfusdAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bfusdAmount")
		delete(additionalProperties, "usdtProfit")
		delete(additionalProperties, "bfusdProfit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBfusdAccountResponse struct {
	value *GetBfusdAccountResponse
	isSet bool
}

func (v NullableGetBfusdAccountResponse) Get() *GetBfusdAccountResponse {
	return v.value
}

func (v *NullableGetBfusdAccountResponse) Set(val *GetBfusdAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdAccountResponse(val *GetBfusdAccountResponse) *NullableGetBfusdAccountResponse {
	return &NullableGetBfusdAccountResponse{value: val, isSet: true}
}

func (v NullableGetBfusdAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
