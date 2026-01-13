/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetRwusdAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRwusdAccountResponse{}

// GetRwusdAccountResponse struct for GetRwusdAccountResponse
type GetRwusdAccountResponse struct {
	RwusdAmount          *string `json:"rwusdAmount,omitempty"`
	TotalProfit          *string `json:"totalProfit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRwusdAccountResponse GetRwusdAccountResponse

// NewGetRwusdAccountResponse instantiates a new GetRwusdAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRwusdAccountResponse() *GetRwusdAccountResponse {
	this := GetRwusdAccountResponse{}
	return &this
}

// NewGetRwusdAccountResponseWithDefaults instantiates a new GetRwusdAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRwusdAccountResponseWithDefaults() *GetRwusdAccountResponse {
	this := GetRwusdAccountResponse{}
	return &this
}

// GetRwusdAmount returns the RwusdAmount field value if set, zero value otherwise.
func (o *GetRwusdAccountResponse) GetRwusdAmount() string {
	if o == nil || common.IsNil(o.RwusdAmount) {
		var ret string
		return ret
	}
	return *o.RwusdAmount
}

// GetRwusdAmountOk returns a tuple with the RwusdAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdAccountResponse) GetRwusdAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RwusdAmount) {
		return nil, false
	}
	return o.RwusdAmount, true
}

// HasRwusdAmount returns a boolean if a field has been set.
func (o *GetRwusdAccountResponse) HasRwusdAmount() bool {
	if o != nil && !common.IsNil(o.RwusdAmount) {
		return true
	}

	return false
}

// SetRwusdAmount gets a reference to the given string and assigns it to the RwusdAmount field.
func (o *GetRwusdAccountResponse) SetRwusdAmount(v string) {
	o.RwusdAmount = &v
}

// GetTotalProfit returns the TotalProfit field value if set, zero value otherwise.
func (o *GetRwusdAccountResponse) GetTotalProfit() string {
	if o == nil || common.IsNil(o.TotalProfit) {
		var ret string
		return ret
	}
	return *o.TotalProfit
}

// GetTotalProfitOk returns a tuple with the TotalProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdAccountResponse) GetTotalProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalProfit) {
		return nil, false
	}
	return o.TotalProfit, true
}

// HasTotalProfit returns a boolean if a field has been set.
func (o *GetRwusdAccountResponse) HasTotalProfit() bool {
	if o != nil && !common.IsNil(o.TotalProfit) {
		return true
	}

	return false
}

// SetTotalProfit gets a reference to the given string and assigns it to the TotalProfit field.
func (o *GetRwusdAccountResponse) SetTotalProfit(v string) {
	o.TotalProfit = &v
}

func (o GetRwusdAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRwusdAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RwusdAmount) {
		toSerialize["rwusdAmount"] = o.RwusdAmount
	}
	if !common.IsNil(o.TotalProfit) {
		toSerialize["totalProfit"] = o.TotalProfit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRwusdAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varGetRwusdAccountResponse := _GetRwusdAccountResponse{}

	err = json.Unmarshal(data, &varGetRwusdAccountResponse)

	if err != nil {
		return err
	}

	*o = GetRwusdAccountResponse(varGetRwusdAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rwusdAmount")
		delete(additionalProperties, "totalProfit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRwusdAccountResponse struct {
	value *GetRwusdAccountResponse
	isSet bool
}

func (v NullableGetRwusdAccountResponse) Get() *GetRwusdAccountResponse {
	return v.value
}

func (v *NullableGetRwusdAccountResponse) Set(val *GetRwusdAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRwusdAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRwusdAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRwusdAccountResponse(val *GetRwusdAccountResponse) *NullableGetRwusdAccountResponse {
	return &NullableGetRwusdAccountResponse{value: val, isSet: true}
}

func (v NullableGetRwusdAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRwusdAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
