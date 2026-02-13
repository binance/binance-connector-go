/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryIsolatedMarginAccountInfoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryIsolatedMarginAccountInfoResponse{}

// QueryIsolatedMarginAccountInfoResponse struct for QueryIsolatedMarginAccountInfoResponse
type QueryIsolatedMarginAccountInfoResponse struct {
	Assets               []QueryIsolatedMarginAccountInfoResponseAssetsInner `json:"assets,omitempty"`
	TotalAssetOfBtc      *string                                             `json:"totalAssetOfBtc,omitempty"`
	TotalLiabilityOfBtc  *string                                             `json:"totalLiabilityOfBtc,omitempty"`
	TotalNetAssetOfBtc   *string                                             `json:"totalNetAssetOfBtc,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryIsolatedMarginAccountInfoResponse QueryIsolatedMarginAccountInfoResponse

// NewQueryIsolatedMarginAccountInfoResponse instantiates a new QueryIsolatedMarginAccountInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryIsolatedMarginAccountInfoResponse() *QueryIsolatedMarginAccountInfoResponse {
	this := QueryIsolatedMarginAccountInfoResponse{}
	return &this
}

// NewQueryIsolatedMarginAccountInfoResponseWithDefaults instantiates a new QueryIsolatedMarginAccountInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryIsolatedMarginAccountInfoResponseWithDefaults() *QueryIsolatedMarginAccountInfoResponse {
	this := QueryIsolatedMarginAccountInfoResponse{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponse) GetAssets() []QueryIsolatedMarginAccountInfoResponseAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []QueryIsolatedMarginAccountInfoResponseAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponse) GetAssetsOk() ([]QueryIsolatedMarginAccountInfoResponseAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponse) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []QueryIsolatedMarginAccountInfoResponseAssetsInner and assigns it to the Assets field.
func (o *QueryIsolatedMarginAccountInfoResponse) SetAssets(v []QueryIsolatedMarginAccountInfoResponseAssetsInner) {
	o.Assets = v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponse) GetTotalAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponse) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		return nil, false
	}
	return o.TotalAssetOfBtc, true
}

// HasTotalAssetOfBtc returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponse) HasTotalAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalAssetOfBtc gets a reference to the given string and assigns it to the TotalAssetOfBtc field.
func (o *QueryIsolatedMarginAccountInfoResponse) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = &v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponse) GetTotalLiabilityOfBtc() string {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponse) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		return nil, false
	}
	return o.TotalLiabilityOfBtc, true
}

// HasTotalLiabilityOfBtc returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponse) HasTotalLiabilityOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalLiabilityOfBtc) {
		return true
	}

	return false
}

// SetTotalLiabilityOfBtc gets a reference to the given string and assigns it to the TotalLiabilityOfBtc field.
func (o *QueryIsolatedMarginAccountInfoResponse) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = &v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponse) GetTotalNetAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponse) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		return nil, false
	}
	return o.TotalNetAssetOfBtc, true
}

// HasTotalNetAssetOfBtc returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponse) HasTotalNetAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalNetAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalNetAssetOfBtc gets a reference to the given string and assigns it to the TotalNetAssetOfBtc field.
func (o *QueryIsolatedMarginAccountInfoResponse) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = &v
}

func (o QueryIsolatedMarginAccountInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryIsolatedMarginAccountInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryIsolatedMarginAccountInfoResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryIsolatedMarginAccountInfoResponse := _QueryIsolatedMarginAccountInfoResponse{}

	err = json.Unmarshal(data, &varQueryIsolatedMarginAccountInfoResponse)

	if err != nil {
		return err
	}

	*o = QueryIsolatedMarginAccountInfoResponse(varQueryIsolatedMarginAccountInfoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assets")
		delete(additionalProperties, "totalAssetOfBtc")
		delete(additionalProperties, "totalLiabilityOfBtc")
		delete(additionalProperties, "totalNetAssetOfBtc")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryIsolatedMarginAccountInfoResponse struct {
	value *QueryIsolatedMarginAccountInfoResponse
	isSet bool
}

func (v NullableQueryIsolatedMarginAccountInfoResponse) Get() *QueryIsolatedMarginAccountInfoResponse {
	return v.value
}

func (v *NullableQueryIsolatedMarginAccountInfoResponse) Set(val *QueryIsolatedMarginAccountInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryIsolatedMarginAccountInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryIsolatedMarginAccountInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryIsolatedMarginAccountInfoResponse(val *QueryIsolatedMarginAccountInfoResponse) *NullableQueryIsolatedMarginAccountInfoResponse {
	return &NullableQueryIsolatedMarginAccountInfoResponse{value: val, isSet: true}
}

func (v NullableQueryIsolatedMarginAccountInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryIsolatedMarginAccountInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
