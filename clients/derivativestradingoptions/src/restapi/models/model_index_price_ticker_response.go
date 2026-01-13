/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the IndexPriceTickerResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexPriceTickerResponse{}

// IndexPriceTickerResponse struct for IndexPriceTickerResponse
type IndexPriceTickerResponse struct {
	Time                 *int64  `json:"time,omitempty"`
	IndexPrice           *string `json:"indexPrice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IndexPriceTickerResponse IndexPriceTickerResponse

// NewIndexPriceTickerResponse instantiates a new IndexPriceTickerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexPriceTickerResponse() *IndexPriceTickerResponse {
	this := IndexPriceTickerResponse{}
	return &this
}

// NewIndexPriceTickerResponseWithDefaults instantiates a new IndexPriceTickerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexPriceTickerResponseWithDefaults() *IndexPriceTickerResponse {
	this := IndexPriceTickerResponse{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *IndexPriceTickerResponse) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceTickerResponse) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *IndexPriceTickerResponse) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *IndexPriceTickerResponse) SetTime(v int64) {
	o.Time = &v
}

// GetIndexPrice returns the IndexPrice field value if set, zero value otherwise.
func (o *IndexPriceTickerResponse) GetIndexPrice() string {
	if o == nil || common.IsNil(o.IndexPrice) {
		var ret string
		return ret
	}
	return *o.IndexPrice
}

// GetIndexPriceOk returns a tuple with the IndexPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceTickerResponse) GetIndexPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.IndexPrice) {
		return nil, false
	}
	return o.IndexPrice, true
}

// HasIndexPrice returns a boolean if a field has been set.
func (o *IndexPriceTickerResponse) HasIndexPrice() bool {
	if o != nil && !common.IsNil(o.IndexPrice) {
		return true
	}

	return false
}

// SetIndexPrice gets a reference to the given string and assigns it to the IndexPrice field.
func (o *IndexPriceTickerResponse) SetIndexPrice(v string) {
	o.IndexPrice = &v
}

func (o IndexPriceTickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexPriceTickerResponse) ToMap() (map[string]interface{}, error) {
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

func (o *IndexPriceTickerResponse) UnmarshalJSON(data []byte) (err error) {
	varIndexPriceTickerResponse := _IndexPriceTickerResponse{}

	err = json.Unmarshal(data, &varIndexPriceTickerResponse)

	if err != nil {
		return err
	}

	*o = IndexPriceTickerResponse(varIndexPriceTickerResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "indexPrice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndexPriceTickerResponse struct {
	value *IndexPriceTickerResponse
	isSet bool
}

func (v NullableIndexPriceTickerResponse) Get() *IndexPriceTickerResponse {
	return v.value
}

func (v *NullableIndexPriceTickerResponse) Set(val *IndexPriceTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceTickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexPriceTickerResponse(val *IndexPriceTickerResponse) *NullableIndexPriceTickerResponse {
	return &NullableIndexPriceTickerResponse{value: val, isSet: true}
}

func (v NullableIndexPriceTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
