/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySubAccountAssetsAssetManagementResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountAssetsAssetManagementResponse{}

// QuerySubAccountAssetsAssetManagementResponse struct for QuerySubAccountAssetsAssetManagementResponse
type QuerySubAccountAssetsAssetManagementResponse struct {
	Balances             []QuerySubAccountAssetsAssetManagementResponseBalancesInner `json:"balances,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubAccountAssetsAssetManagementResponse QuerySubAccountAssetsAssetManagementResponse

// NewQuerySubAccountAssetsAssetManagementResponse instantiates a new QuerySubAccountAssetsAssetManagementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountAssetsAssetManagementResponse() *QuerySubAccountAssetsAssetManagementResponse {
	this := QuerySubAccountAssetsAssetManagementResponse{}
	return &this
}

// NewQuerySubAccountAssetsAssetManagementResponseWithDefaults instantiates a new QuerySubAccountAssetsAssetManagementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountAssetsAssetManagementResponseWithDefaults() *QuerySubAccountAssetsAssetManagementResponse {
	this := QuerySubAccountAssetsAssetManagementResponse{}
	return &this
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *QuerySubAccountAssetsAssetManagementResponse) GetBalances() []QuerySubAccountAssetsAssetManagementResponseBalancesInner {
	if o == nil || common.IsNil(o.Balances) {
		var ret []QuerySubAccountAssetsAssetManagementResponseBalancesInner
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountAssetsAssetManagementResponse) GetBalancesOk() ([]QuerySubAccountAssetsAssetManagementResponseBalancesInner, bool) {
	if o == nil || common.IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *QuerySubAccountAssetsAssetManagementResponse) HasBalances() bool {
	if o != nil && !common.IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []QuerySubAccountAssetsAssetManagementResponseBalancesInner and assigns it to the Balances field.
func (o *QuerySubAccountAssetsAssetManagementResponse) SetBalances(v []QuerySubAccountAssetsAssetManagementResponseBalancesInner) {
	o.Balances = v
}

func (o QuerySubAccountAssetsAssetManagementResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountAssetsAssetManagementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubAccountAssetsAssetManagementResponse) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountAssetsAssetManagementResponse := _QuerySubAccountAssetsAssetManagementResponse{}

	err = json.Unmarshal(data, &varQuerySubAccountAssetsAssetManagementResponse)

	if err != nil {
		return err
	}

	*o = QuerySubAccountAssetsAssetManagementResponse(varQuerySubAccountAssetsAssetManagementResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "balances")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountAssetsAssetManagementResponse struct {
	value *QuerySubAccountAssetsAssetManagementResponse
	isSet bool
}

func (v NullableQuerySubAccountAssetsAssetManagementResponse) Get() *QuerySubAccountAssetsAssetManagementResponse {
	return v.value
}

func (v *NullableQuerySubAccountAssetsAssetManagementResponse) Set(val *QuerySubAccountAssetsAssetManagementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountAssetsAssetManagementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountAssetsAssetManagementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountAssetsAssetManagementResponse(val *QuerySubAccountAssetsAssetManagementResponse) *NullableQuerySubAccountAssetsAssetManagementResponse {
	return &NullableQuerySubAccountAssetsAssetManagementResponse{value: val, isSet: true}
}

func (v NullableQuerySubAccountAssetsAssetManagementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountAssetsAssetManagementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
