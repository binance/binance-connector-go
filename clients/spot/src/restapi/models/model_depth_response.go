/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DepthResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DepthResponse{}

// DepthResponse struct for DepthResponse
type DepthResponse struct {
	LastUpdateId         *int64     `json:"lastUpdateId,omitempty"`
	Bids                 [][]string `json:"bids,omitempty"`
	Asks                 [][]string `json:"asks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DepthResponse DepthResponse

// NewDepthResponse instantiates a new DepthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepthResponse() *DepthResponse {
	this := DepthResponse{}
	return &this
}

// NewDepthResponseWithDefaults instantiates a new DepthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepthResponseWithDefaults() *DepthResponse {
	this := DepthResponse{}
	return &this
}

// GetLastUpdateId returns the LastUpdateId field value if set, zero value otherwise.
func (o *DepthResponse) GetLastUpdateId() int64 {
	if o == nil || common.IsNil(o.LastUpdateId) {
		var ret int64
		return ret
	}
	return *o.LastUpdateId
}

// GetLastUpdateIdOk returns a tuple with the LastUpdateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepthResponse) GetLastUpdateIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastUpdateId) {
		return nil, false
	}
	return o.LastUpdateId, true
}

// HasLastUpdateId returns a boolean if a field has been set.
func (o *DepthResponse) HasLastUpdateId() bool {
	if o != nil && !common.IsNil(o.LastUpdateId) {
		return true
	}

	return false
}

// SetLastUpdateId gets a reference to the given int64 and assigns it to the LastUpdateId field.
func (o *DepthResponse) SetLastUpdateId(v int64) {
	o.LastUpdateId = &v
}

// GetBids returns the Bids field value if set, zero value otherwise.
func (o *DepthResponse) GetBids() [][]string {
	if o == nil || common.IsNil(o.Bids) {
		var ret [][]string
		return ret
	}
	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepthResponse) GetBidsOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Bids) {
		return nil, false
	}
	return o.Bids, true
}

// HasBids returns a boolean if a field has been set.
func (o *DepthResponse) HasBids() bool {
	if o != nil && !common.IsNil(o.Bids) {
		return true
	}

	return false
}

// SetBids gets a reference to the given [][]string and assigns it to the Bids field.
func (o *DepthResponse) SetBids(v [][]string) {
	o.Bids = v
}

// GetAsks returns the Asks field value if set, zero value otherwise.
func (o *DepthResponse) GetAsks() [][]string {
	if o == nil || common.IsNil(o.Asks) {
		var ret [][]string
		return ret
	}
	return o.Asks
}

// GetAsksOk returns a tuple with the Asks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepthResponse) GetAsksOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Asks) {
		return nil, false
	}
	return o.Asks, true
}

// HasAsks returns a boolean if a field has been set.
func (o *DepthResponse) HasAsks() bool {
	if o != nil && !common.IsNil(o.Asks) {
		return true
	}

	return false
}

// SetAsks gets a reference to the given [][]string and assigns it to the Asks field.
func (o *DepthResponse) SetAsks(v [][]string) {
	o.Asks = v
}

func (o DepthResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepthResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LastUpdateId) {
		toSerialize["lastUpdateId"] = o.LastUpdateId
	}
	if !common.IsNil(o.Bids) {
		toSerialize["bids"] = o.Bids
	}
	if !common.IsNil(o.Asks) {
		toSerialize["asks"] = o.Asks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DepthResponse) UnmarshalJSON(data []byte) (err error) {
	varDepthResponse := _DepthResponse{}

	err = json.Unmarshal(data, &varDepthResponse)

	if err != nil {
		return err
	}

	*o = DepthResponse(varDepthResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lastUpdateId")
		delete(additionalProperties, "bids")
		delete(additionalProperties, "asks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepthResponse struct {
	value *DepthResponse
	isSet bool
}

func (v NullableDepthResponse) Get() *DepthResponse {
	return v.value
}

func (v *NullableDepthResponse) Set(val *DepthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDepthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDepthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepthResponse(val *DepthResponse) *NullableDepthResponse {
	return &NullableDepthResponse{value: val, isSet: true}
}

func (v NullableDepthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
