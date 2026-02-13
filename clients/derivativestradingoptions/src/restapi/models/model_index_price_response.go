/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the IndexPriceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexPriceResponse{}

// IndexPriceResponse struct for IndexPriceResponse
type IndexPriceResponse struct {
	Time                 *int64  `json:"time,omitempty"`
	IndexPrice           *string `json:"indexPrice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IndexPriceResponse IndexPriceResponse

// NewIndexPriceResponse instantiates a new IndexPriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexPriceResponse() *IndexPriceResponse {
	this := IndexPriceResponse{}
	return &this
}

// NewIndexPriceResponseWithDefaults instantiates a new IndexPriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexPriceResponseWithDefaults() *IndexPriceResponse {
	this := IndexPriceResponse{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *IndexPriceResponse) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceResponse) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *IndexPriceResponse) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *IndexPriceResponse) SetTime(v int64) {
	o.Time = &v
}

// GetIndexPrice returns the IndexPrice field value if set, zero value otherwise.
func (o *IndexPriceResponse) GetIndexPrice() string {
	if o == nil || common.IsNil(o.IndexPrice) {
		var ret string
		return ret
	}
	return *o.IndexPrice
}

// GetIndexPriceOk returns a tuple with the IndexPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceResponse) GetIndexPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.IndexPrice) {
		return nil, false
	}
	return o.IndexPrice, true
}

// HasIndexPrice returns a boolean if a field has been set.
func (o *IndexPriceResponse) HasIndexPrice() bool {
	if o != nil && !common.IsNil(o.IndexPrice) {
		return true
	}

	return false
}

// SetIndexPrice gets a reference to the given string and assigns it to the IndexPrice field.
func (o *IndexPriceResponse) SetIndexPrice(v string) {
	o.IndexPrice = &v
}

func (o IndexPriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexPriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.IndexPrice) {
		toSerialize["indexPrice"] = o.IndexPrice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IndexPriceResponse) UnmarshalJSON(data []byte) (err error) {
	varIndexPriceResponse := _IndexPriceResponse{}

	err = json.Unmarshal(data, &varIndexPriceResponse)

	if err != nil {
		return err
	}

	*o = IndexPriceResponse(varIndexPriceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "indexPrice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndexPriceResponse struct {
	value *IndexPriceResponse
	isSet bool
}

func (v NullableIndexPriceResponse) Get() *IndexPriceResponse {
	return v.value
}

func (v *NullableIndexPriceResponse) Set(val *IndexPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexPriceResponse(val *IndexPriceResponse) *NullableIndexPriceResponse {
	return &NullableIndexPriceResponse{value: val, isSet: true}
}

func (v NullableIndexPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
