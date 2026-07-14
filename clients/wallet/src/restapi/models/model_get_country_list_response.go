/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetCountryListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCountryListResponse{}

// GetCountryListResponse struct for GetCountryListResponse
type GetCountryListResponse struct {
	Countries []GetCountryListResponseCountriesInner `json:"countries,omitempty"`
	// Last data update timestamp (epoch milliseconds); 0 if empty.
	LastUpdated          *int64 `json:"lastUpdated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCountryListResponse GetCountryListResponse

// NewGetCountryListResponse instantiates a new GetCountryListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCountryListResponse() *GetCountryListResponse {
	this := GetCountryListResponse{}
	return &this
}

// NewGetCountryListResponseWithDefaults instantiates a new GetCountryListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCountryListResponseWithDefaults() *GetCountryListResponse {
	this := GetCountryListResponse{}
	return &this
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *GetCountryListResponse) GetCountries() []GetCountryListResponseCountriesInner {
	if o == nil || common.IsNil(o.Countries) {
		var ret []GetCountryListResponseCountriesInner
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCountryListResponse) GetCountriesOk() ([]GetCountryListResponseCountriesInner, bool) {
	if o == nil || common.IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *GetCountryListResponse) HasCountries() bool {
	if o != nil && !common.IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []GetCountryListResponseCountriesInner and assigns it to the Countries field.
func (o *GetCountryListResponse) SetCountries(v []GetCountryListResponseCountriesInner) {
	o.Countries = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *GetCountryListResponse) GetLastUpdated() int64 {
	if o == nil || common.IsNil(o.LastUpdated) {
		var ret int64
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCountryListResponse) GetLastUpdatedOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *GetCountryListResponse) HasLastUpdated() bool {
	if o != nil && !common.IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given int64 and assigns it to the LastUpdated field.
func (o *GetCountryListResponse) SetLastUpdated(v int64) {
	o.LastUpdated = &v
}

func (o GetCountryListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCountryListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Countries) {
		toSerialize["countries"] = o.Countries
	}
	if !common.IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCountryListResponse) UnmarshalJSON(data []byte) (err error) {
	varGetCountryListResponse := _GetCountryListResponse{}

	err = json.Unmarshal(data, &varGetCountryListResponse)

	if err != nil {
		return err
	}

	*o = GetCountryListResponse(varGetCountryListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "countries")
		delete(additionalProperties, "lastUpdated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCountryListResponse struct {
	value *GetCountryListResponse
	isSet bool
}

func (v NullableGetCountryListResponse) Get() *GetCountryListResponse {
	return v.value
}

func (v *NullableGetCountryListResponse) Set(val *GetCountryListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCountryListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCountryListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCountryListResponse(val *GetCountryListResponse) *NullableGetCountryListResponse {
	return &NullableGetCountryListResponse{value: val, isSet: true}
}

func (v NullableGetCountryListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCountryListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
