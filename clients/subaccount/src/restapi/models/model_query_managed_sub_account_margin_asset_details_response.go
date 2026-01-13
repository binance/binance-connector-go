/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryManagedSubAccountMarginAssetDetailsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountMarginAssetDetailsResponse{}

// QueryManagedSubAccountMarginAssetDetailsResponse struct for QueryManagedSubAccountMarginAssetDetailsResponse
type QueryManagedSubAccountMarginAssetDetailsResponse struct {
	MarginLevel          *string                                                           `json:"marginLevel,omitempty"`
	TotalAssetOfBtc      *string                                                           `json:"totalAssetOfBtc,omitempty"`
	TotalLiabilityOfBtc  *string                                                           `json:"totalLiabilityOfBtc,omitempty"`
	TotalNetAssetOfBtc   *string                                                           `json:"totalNetAssetOfBtc,omitempty"`
	UserAssets           []QueryManagedSubAccountMarginAssetDetailsResponseUserAssetsInner `json:"userAssets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountMarginAssetDetailsResponse QueryManagedSubAccountMarginAssetDetailsResponse

// NewQueryManagedSubAccountMarginAssetDetailsResponse instantiates a new QueryManagedSubAccountMarginAssetDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountMarginAssetDetailsResponse() *QueryManagedSubAccountMarginAssetDetailsResponse {
	this := QueryManagedSubAccountMarginAssetDetailsResponse{}
	return &this
}

// NewQueryManagedSubAccountMarginAssetDetailsResponseWithDefaults instantiates a new QueryManagedSubAccountMarginAssetDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountMarginAssetDetailsResponseWithDefaults() *QueryManagedSubAccountMarginAssetDetailsResponse {
	this := QueryManagedSubAccountMarginAssetDetailsResponse{}
	return &this
}

// GetMarginLevel returns the MarginLevel field value if set, zero value otherwise.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) GetMarginLevel() string {
	if o == nil || common.IsNil(o.MarginLevel) {
		var ret string
		return ret
	}
	return *o.MarginLevel
}

// GetMarginLevelOk returns a tuple with the MarginLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) GetMarginLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginLevel) {
		return nil, false
	}
	return o.MarginLevel, true
}

// HasMarginLevel returns a boolean if a field has been set.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) HasMarginLevel() bool {
	if o != nil && !common.IsNil(o.MarginLevel) {
		return true
	}

	return false
}

// SetMarginLevel gets a reference to the given string and assigns it to the MarginLevel field.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) SetMarginLevel(v string) {
	o.MarginLevel = &v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value if set, zero value otherwise.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) GetTotalAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		return nil, false
	}
	return o.TotalAssetOfBtc, true
}

// HasTotalAssetOfBtc returns a boolean if a field has been set.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) HasTotalAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalAssetOfBtc gets a reference to the given string and assigns it to the TotalAssetOfBtc field.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = &v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value if set, zero value otherwise.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) GetTotalLiabilityOfBtc() string {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		return nil, false
	}
	return o.TotalLiabilityOfBtc, true
}

// HasTotalLiabilityOfBtc returns a boolean if a field has been set.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) HasTotalLiabilityOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalLiabilityOfBtc) {
		return true
	}

	return false
}

// SetTotalLiabilityOfBtc gets a reference to the given string and assigns it to the TotalLiabilityOfBtc field.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = &v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value if set, zero value otherwise.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) GetTotalNetAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		return nil, false
	}
	return o.TotalNetAssetOfBtc, true
}

// HasTotalNetAssetOfBtc returns a boolean if a field has been set.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) HasTotalNetAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalNetAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalNetAssetOfBtc gets a reference to the given string and assigns it to the TotalNetAssetOfBtc field.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = &v
}

// GetUserAssets returns the UserAssets field value if set, zero value otherwise.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) GetUserAssets() []QueryManagedSubAccountMarginAssetDetailsResponseUserAssetsInner {
	if o == nil || common.IsNil(o.UserAssets) {
		var ret []QueryManagedSubAccountMarginAssetDetailsResponseUserAssetsInner
		return ret
	}
	return o.UserAssets
}

// GetUserAssetsOk returns a tuple with the UserAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) GetUserAssetsOk() ([]QueryManagedSubAccountMarginAssetDetailsResponseUserAssetsInner, bool) {
	if o == nil || common.IsNil(o.UserAssets) {
		return nil, false
	}
	return o.UserAssets, true
}

// HasUserAssets returns a boolean if a field has been set.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) HasUserAssets() bool {
	if o != nil && !common.IsNil(o.UserAssets) {
		return true
	}

	return false
}

// SetUserAssets gets a reference to the given []QueryManagedSubAccountMarginAssetDetailsResponseUserAssetsInner and assigns it to the UserAssets field.
func (o *QueryManagedSubAccountMarginAssetDetailsResponse) SetUserAssets(v []QueryManagedSubAccountMarginAssetDetailsResponseUserAssetsInner) {
	o.UserAssets = v
}

func (o QueryManagedSubAccountMarginAssetDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountMarginAssetDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MarginLevel) {
		toSerialize["marginLevel"] = o.MarginLevel
	}
	if !common.IsNil(o.TotalAssetOfBtc) {
		toSerialize["totalAssetOfBtc"] = o.TotalAssetOfBtc
	}
	if !common.IsNil(o.TotalLiabilityOfBtc) {
		toSerialize["totalLiabilityOfBtc"] = o.TotalLiabilityOfBtc
	}
	if !common.IsNil(o.TotalNetAssetOfBtc) {
		toSerialize["totalNetAssetOfBtc"] = o.TotalNetAssetOfBtc
	}
	if !common.IsNil(o.UserAssets) {
		toSerialize["userAssets"] = o.UserAssets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryManagedSubAccountMarginAssetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountMarginAssetDetailsResponse := _QueryManagedSubAccountMarginAssetDetailsResponse{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountMarginAssetDetailsResponse)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountMarginAssetDetailsResponse(varQueryManagedSubAccountMarginAssetDetailsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "marginLevel")
		delete(additionalProperties, "totalAssetOfBtc")
		delete(additionalProperties, "totalLiabilityOfBtc")
		delete(additionalProperties, "totalNetAssetOfBtc")
		delete(additionalProperties, "userAssets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountMarginAssetDetailsResponse struct {
	value *QueryManagedSubAccountMarginAssetDetailsResponse
	isSet bool
}

func (v NullableQueryManagedSubAccountMarginAssetDetailsResponse) Get() *QueryManagedSubAccountMarginAssetDetailsResponse {
	return v.value
}

func (v *NullableQueryManagedSubAccountMarginAssetDetailsResponse) Set(val *QueryManagedSubAccountMarginAssetDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountMarginAssetDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountMarginAssetDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountMarginAssetDetailsResponse(val *QueryManagedSubAccountMarginAssetDetailsResponse) *NullableQueryManagedSubAccountMarginAssetDetailsResponse {
	return &NullableQueryManagedSubAccountMarginAssetDetailsResponse{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountMarginAssetDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountMarginAssetDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
