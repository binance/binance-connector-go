/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DustlogResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DustlogResponse{}

// DustlogResponse struct for DustlogResponse
type DustlogResponse struct {
	Total                *int64                                   `json:"total,omitempty"`
	UserAssetDribblets   []DustlogResponseUserAssetDribbletsInner `json:"userAssetDribblets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DustlogResponse DustlogResponse

// NewDustlogResponse instantiates a new DustlogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDustlogResponse() *DustlogResponse {
	this := DustlogResponse{}
	return &this
}

// NewDustlogResponseWithDefaults instantiates a new DustlogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDustlogResponseWithDefaults() *DustlogResponse {
	this := DustlogResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DustlogResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DustlogResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *DustlogResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetUserAssetDribblets returns the UserAssetDribblets field value if set, zero value otherwise.
func (o *DustlogResponse) GetUserAssetDribblets() []DustlogResponseUserAssetDribbletsInner {
	if o == nil || common.IsNil(o.UserAssetDribblets) {
		var ret []DustlogResponseUserAssetDribbletsInner
		return ret
	}
	return o.UserAssetDribblets
}

// GetUserAssetDribbletsOk returns a tuple with the UserAssetDribblets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponse) GetUserAssetDribbletsOk() ([]DustlogResponseUserAssetDribbletsInner, bool) {
	if o == nil || common.IsNil(o.UserAssetDribblets) {
		return nil, false
	}
	return o.UserAssetDribblets, true
}

// HasUserAssetDribblets returns a boolean if a field has been set.
func (o *DustlogResponse) HasUserAssetDribblets() bool {
	if o != nil && !common.IsNil(o.UserAssetDribblets) {
		return true
	}

	return false
}

// SetUserAssetDribblets gets a reference to the given []DustlogResponseUserAssetDribbletsInner and assigns it to the UserAssetDribblets field.
func (o *DustlogResponse) SetUserAssetDribblets(v []DustlogResponseUserAssetDribbletsInner) {
	o.UserAssetDribblets = v
}

func (o DustlogResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DustlogResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.UserAssetDribblets) {
		toSerialize["userAssetDribblets"] = o.UserAssetDribblets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DustlogResponse) UnmarshalJSON(data []byte) (err error) {
	varDustlogResponse := _DustlogResponse{}

	err = json.Unmarshal(data, &varDustlogResponse)

	if err != nil {
		return err
	}

	*o = DustlogResponse(varDustlogResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "userAssetDribblets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDustlogResponse struct {
	value *DustlogResponse
	isSet bool
}

func (v NullableDustlogResponse) Get() *DustlogResponse {
	return v.value
}

func (v *NullableDustlogResponse) Set(val *DustlogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDustlogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDustlogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDustlogResponse(val *DustlogResponse) *NullableDustlogResponse {
	return &NullableDustlogResponse{value: val, isSet: true}
}

func (v NullableDustlogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDustlogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
