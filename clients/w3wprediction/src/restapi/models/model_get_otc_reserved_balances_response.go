/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOtcReservedBalancesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOtcReservedBalancesResponse{}

// GetOtcReservedBalancesResponse struct for GetOtcReservedBalancesResponse
type GetOtcReservedBalancesResponse struct {
	Balances             []GetOtcReservedBalancesResponseBalancesInner `json:"balances,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOtcReservedBalancesResponse GetOtcReservedBalancesResponse

// NewGetOtcReservedBalancesResponse instantiates a new GetOtcReservedBalancesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOtcReservedBalancesResponse() *GetOtcReservedBalancesResponse {
	this := GetOtcReservedBalancesResponse{}
	return &this
}

// NewGetOtcReservedBalancesResponseWithDefaults instantiates a new GetOtcReservedBalancesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOtcReservedBalancesResponseWithDefaults() *GetOtcReservedBalancesResponse {
	this := GetOtcReservedBalancesResponse{}
	return &this
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *GetOtcReservedBalancesResponse) GetBalances() []GetOtcReservedBalancesResponseBalancesInner {
	if o == nil || common.IsNil(o.Balances) {
		var ret []GetOtcReservedBalancesResponseBalancesInner
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcReservedBalancesResponse) GetBalancesOk() ([]GetOtcReservedBalancesResponseBalancesInner, bool) {
	if o == nil || common.IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *GetOtcReservedBalancesResponse) HasBalances() bool {
	if o != nil && !common.IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []GetOtcReservedBalancesResponseBalancesInner and assigns it to the Balances field.
func (o *GetOtcReservedBalancesResponse) SetBalances(v []GetOtcReservedBalancesResponseBalancesInner) {
	o.Balances = v
}

func (o GetOtcReservedBalancesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOtcReservedBalancesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOtcReservedBalancesResponse) UnmarshalJSON(data []byte) (err error) {
	varGetOtcReservedBalancesResponse := _GetOtcReservedBalancesResponse{}

	err = json.Unmarshal(data, &varGetOtcReservedBalancesResponse)

	if err != nil {
		return err
	}

	*o = GetOtcReservedBalancesResponse(varGetOtcReservedBalancesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "balances")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOtcReservedBalancesResponse struct {
	value *GetOtcReservedBalancesResponse
	isSet bool
}

func (v NullableGetOtcReservedBalancesResponse) Get() *GetOtcReservedBalancesResponse {
	return v.value
}

func (v *NullableGetOtcReservedBalancesResponse) Set(val *GetOtcReservedBalancesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOtcReservedBalancesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOtcReservedBalancesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOtcReservedBalancesResponse(val *GetOtcReservedBalancesResponse) *NullableGetOtcReservedBalancesResponse {
	return &NullableGetOtcReservedBalancesResponse{value: val, isSet: true}
}

func (v NullableGetOtcReservedBalancesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOtcReservedBalancesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
