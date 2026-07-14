/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetRegionListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRegionListResponse{}

// GetRegionListResponse struct for GetRegionListResponse
type GetRegionListResponse struct {
	// Echoed country code (lowercase).
	CountryCode *string                             `json:"countryCode,omitempty"`
	Regions     []GetRegionListResponseRegionsInner `json:"regions,omitempty"`
	// Last data update timestamp (epoch milliseconds); 0 if empty.
	LastUpdated          *int64 `json:"lastUpdated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRegionListResponse GetRegionListResponse

// NewGetRegionListResponse instantiates a new GetRegionListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRegionListResponse() *GetRegionListResponse {
	this := GetRegionListResponse{}
	return &this
}

// NewGetRegionListResponseWithDefaults instantiates a new GetRegionListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRegionListResponseWithDefaults() *GetRegionListResponse {
	this := GetRegionListResponse{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *GetRegionListResponse) GetCountryCode() string {
	if o == nil || common.IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionListResponse) GetCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *GetRegionListResponse) HasCountryCode() bool {
	if o != nil && !common.IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *GetRegionListResponse) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *GetRegionListResponse) GetRegions() []GetRegionListResponseRegionsInner {
	if o == nil || common.IsNil(o.Regions) {
		var ret []GetRegionListResponseRegionsInner
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionListResponse) GetRegionsOk() ([]GetRegionListResponseRegionsInner, bool) {
	if o == nil || common.IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *GetRegionListResponse) HasRegions() bool {
	if o != nil && !common.IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []GetRegionListResponseRegionsInner and assigns it to the Regions field.
func (o *GetRegionListResponse) SetRegions(v []GetRegionListResponseRegionsInner) {
	o.Regions = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *GetRegionListResponse) GetLastUpdated() int64 {
	if o == nil || common.IsNil(o.LastUpdated) {
		var ret int64
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionListResponse) GetLastUpdatedOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *GetRegionListResponse) HasLastUpdated() bool {
	if o != nil && !common.IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given int64 and assigns it to the LastUpdated field.
func (o *GetRegionListResponse) SetLastUpdated(v int64) {
	o.LastUpdated = &v
}

func (o GetRegionListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRegionListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !common.IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !common.IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRegionListResponse) UnmarshalJSON(data []byte) (err error) {
	varGetRegionListResponse := _GetRegionListResponse{}

	err = json.Unmarshal(data, &varGetRegionListResponse)

	if err != nil {
		return err
	}

	*o = GetRegionListResponse(varGetRegionListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "countryCode")
		delete(additionalProperties, "regions")
		delete(additionalProperties, "lastUpdated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRegionListResponse struct {
	value *GetRegionListResponse
	isSet bool
}

func (v NullableGetRegionListResponse) Get() *GetRegionListResponse {
	return v.value
}

func (v *NullableGetRegionListResponse) Set(val *GetRegionListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRegionListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRegionListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRegionListResponse(val *GetRegionListResponse) *NullableGetRegionListResponse {
	return &NullableGetRegionListResponse{value: val, isSet: true}
}

func (v NullableGetRegionListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRegionListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
