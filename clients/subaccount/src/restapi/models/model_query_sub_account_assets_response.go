/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySubAccountAssetsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountAssetsResponse{}

// QuerySubAccountAssetsResponse struct for QuerySubAccountAssetsResponse
type QuerySubAccountAssetsResponse struct {
	Balances             []QuerySubAccountAssetsResponseBalancesInner `json:"balances,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubAccountAssetsResponse QuerySubAccountAssetsResponse

// NewQuerySubAccountAssetsResponse instantiates a new QuerySubAccountAssetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountAssetsResponse() *QuerySubAccountAssetsResponse {
	this := QuerySubAccountAssetsResponse{}
	return &this
}

// NewQuerySubAccountAssetsResponseWithDefaults instantiates a new QuerySubAccountAssetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountAssetsResponseWithDefaults() *QuerySubAccountAssetsResponse {
	this := QuerySubAccountAssetsResponse{}
	return &this
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *QuerySubAccountAssetsResponse) GetBalances() []QuerySubAccountAssetsResponseBalancesInner {
	if o == nil || common.IsNil(o.Balances) {
		var ret []QuerySubAccountAssetsResponseBalancesInner
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountAssetsResponse) GetBalancesOk() ([]QuerySubAccountAssetsResponseBalancesInner, bool) {
	if o == nil || common.IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *QuerySubAccountAssetsResponse) HasBalances() bool {
	if o != nil && !common.IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []QuerySubAccountAssetsResponseBalancesInner and assigns it to the Balances field.
func (o *QuerySubAccountAssetsResponse) SetBalances(v []QuerySubAccountAssetsResponseBalancesInner) {
	o.Balances = v
}

func (o QuerySubAccountAssetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountAssetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubAccountAssetsResponse) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountAssetsResponse := _QuerySubAccountAssetsResponse{}

	err = json.Unmarshal(data, &varQuerySubAccountAssetsResponse)

	if err != nil {
		return err
	}

	*o = QuerySubAccountAssetsResponse(varQuerySubAccountAssetsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "balances")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountAssetsResponse struct {
	value *QuerySubAccountAssetsResponse
	isSet bool
}

func (v NullableQuerySubAccountAssetsResponse) Get() *QuerySubAccountAssetsResponse {
	return v.value
}

func (v *NullableQuerySubAccountAssetsResponse) Set(val *QuerySubAccountAssetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountAssetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountAssetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountAssetsResponse(val *QuerySubAccountAssetsResponse) *NullableQuerySubAccountAssetsResponse {
	return &NullableQuerySubAccountAssetsResponse{value: val, isSet: true}
}

func (v NullableQuerySubAccountAssetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountAssetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
