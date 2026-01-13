/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetLimitPricePairsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLimitPricePairsResponse{}

// GetLimitPricePairsResponse struct for GetLimitPricePairsResponse
type GetLimitPricePairsResponse struct {
	CrossMarginSymbols   []string `json:"crossMarginSymbols,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLimitPricePairsResponse GetLimitPricePairsResponse

// NewGetLimitPricePairsResponse instantiates a new GetLimitPricePairsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLimitPricePairsResponse() *GetLimitPricePairsResponse {
	this := GetLimitPricePairsResponse{}
	return &this
}

// NewGetLimitPricePairsResponseWithDefaults instantiates a new GetLimitPricePairsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLimitPricePairsResponseWithDefaults() *GetLimitPricePairsResponse {
	this := GetLimitPricePairsResponse{}
	return &this
}

// GetCrossMarginSymbols returns the CrossMarginSymbols field value if set, zero value otherwise.
func (o *GetLimitPricePairsResponse) GetCrossMarginSymbols() []string {
	if o == nil || common.IsNil(o.CrossMarginSymbols) {
		var ret []string
		return ret
	}
	return o.CrossMarginSymbols
}

// GetCrossMarginSymbolsOk returns a tuple with the CrossMarginSymbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLimitPricePairsResponse) GetCrossMarginSymbolsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.CrossMarginSymbols) {
		return nil, false
	}
	return o.CrossMarginSymbols, true
}

// HasCrossMarginSymbols returns a boolean if a field has been set.
func (o *GetLimitPricePairsResponse) HasCrossMarginSymbols() bool {
	if o != nil && !common.IsNil(o.CrossMarginSymbols) {
		return true
	}

	return false
}

// SetCrossMarginSymbols gets a reference to the given []string and assigns it to the CrossMarginSymbols field.
func (o *GetLimitPricePairsResponse) SetCrossMarginSymbols(v []string) {
	o.CrossMarginSymbols = v
}

func (o GetLimitPricePairsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLimitPricePairsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CrossMarginSymbols) {
		toSerialize["crossMarginSymbols"] = o.CrossMarginSymbols
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLimitPricePairsResponse) UnmarshalJSON(data []byte) (err error) {
	varGetLimitPricePairsResponse := _GetLimitPricePairsResponse{}

	err = json.Unmarshal(data, &varGetLimitPricePairsResponse)

	if err != nil {
		return err
	}

	*o = GetLimitPricePairsResponse(varGetLimitPricePairsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "crossMarginSymbols")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLimitPricePairsResponse struct {
	value *GetLimitPricePairsResponse
	isSet bool
}

func (v NullableGetLimitPricePairsResponse) Get() *GetLimitPricePairsResponse {
	return v.value
}

func (v *NullableGetLimitPricePairsResponse) Set(val *GetLimitPricePairsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLimitPricePairsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLimitPricePairsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLimitPricePairsResponse(val *GetLimitPricePairsResponse) *NullableGetLimitPricePairsResponse {
	return &NullableGetLimitPricePairsResponse{value: val, isSet: true}
}

func (v NullableGetLimitPricePairsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLimitPricePairsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
