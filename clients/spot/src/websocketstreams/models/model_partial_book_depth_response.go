/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PartialBookDepthResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PartialBookDepthResponse{}

// PartialBookDepthResponse struct for PartialBookDepthResponse
type PartialBookDepthResponse struct {
	LastUpdateId         *int64     `json:"lastUpdateId,omitempty"`
	Bids                 [][]string `json:"bids,omitempty"`
	Asks                 [][]string `json:"asks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartialBookDepthResponse PartialBookDepthResponse

// NewPartialBookDepthResponse instantiates a new PartialBookDepthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialBookDepthResponse() *PartialBookDepthResponse {
	this := PartialBookDepthResponse{}
	return &this
}

// NewPartialBookDepthResponseWithDefaults instantiates a new PartialBookDepthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialBookDepthResponseWithDefaults() *PartialBookDepthResponse {
	this := PartialBookDepthResponse{}
	return &this
}

// GetLastUpdateId returns the LastUpdateId field value if set, zero value otherwise.
func (o *PartialBookDepthResponse) GetLastUpdateId() int64 {
	if o == nil || common.IsNil(o.LastUpdateId) {
		var ret int64
		return ret
	}
	return *o.LastUpdateId
}

// GetLastUpdateIdOk returns a tuple with the LastUpdateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialBookDepthResponse) GetLastUpdateIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastUpdateId) {
		return nil, false
	}
	return o.LastUpdateId, true
}

// HasLastUpdateId returns a boolean if a field has been set.
func (o *PartialBookDepthResponse) HasLastUpdateId() bool {
	if o != nil && !common.IsNil(o.LastUpdateId) {
		return true
	}

	return false
}

// SetLastUpdateId gets a reference to the given int64 and assigns it to the LastUpdateId field.
func (o *PartialBookDepthResponse) SetLastUpdateId(v int64) {
	o.LastUpdateId = &v
}

// GetBids returns the Bids field value if set, zero value otherwise.
func (o *PartialBookDepthResponse) GetBids() [][]string {
	if o == nil || common.IsNil(o.Bids) {
		var ret [][]string
		return ret
	}
	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialBookDepthResponse) GetBidsOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Bids) {
		return nil, false
	}
	return o.Bids, true
}

// HasBids returns a boolean if a field has been set.
func (o *PartialBookDepthResponse) HasBids() bool {
	if o != nil && !common.IsNil(o.Bids) {
		return true
	}

	return false
}

// SetBids gets a reference to the given [][]string and assigns it to the Bids field.
func (o *PartialBookDepthResponse) SetBids(v [][]string) {
	o.Bids = v
}

// GetAsks returns the Asks field value if set, zero value otherwise.
func (o *PartialBookDepthResponse) GetAsks() [][]string {
	if o == nil || common.IsNil(o.Asks) {
		var ret [][]string
		return ret
	}
	return o.Asks
}

// GetAsksOk returns a tuple with the Asks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialBookDepthResponse) GetAsksOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Asks) {
		return nil, false
	}
	return o.Asks, true
}

// HasAsks returns a boolean if a field has been set.
func (o *PartialBookDepthResponse) HasAsks() bool {
	if o != nil && !common.IsNil(o.Asks) {
		return true
	}

	return false
}

// SetAsks gets a reference to the given [][]string and assigns it to the Asks field.
func (o *PartialBookDepthResponse) SetAsks(v [][]string) {
	o.Asks = v
}

func (o PartialBookDepthResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartialBookDepthResponse) ToMap() (map[string]interface{}, error) {
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

func (o *PartialBookDepthResponse) UnmarshalJSON(data []byte) (err error) {
	varPartialBookDepthResponse := _PartialBookDepthResponse{}

	err = json.Unmarshal(data, &varPartialBookDepthResponse)

	if err != nil {
		return err
	}

	*o = PartialBookDepthResponse(varPartialBookDepthResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lastUpdateId")
		delete(additionalProperties, "bids")
		delete(additionalProperties, "asks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartialBookDepthResponse struct {
	value *PartialBookDepthResponse
	isSet bool
}

func (v NullablePartialBookDepthResponse) Get() *PartialBookDepthResponse {
	return v.value
}

func (v *NullablePartialBookDepthResponse) Set(val *PartialBookDepthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialBookDepthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialBookDepthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialBookDepthResponse(val *PartialBookDepthResponse) *NullablePartialBookDepthResponse {
	return &NullablePartialBookDepthResponse{value: val, isSet: true}
}

func (v NullablePartialBookDepthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialBookDepthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
