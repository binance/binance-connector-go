/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DepthResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DepthResponseResult{}

// DepthResponseResult struct for DepthResponseResult
type DepthResponseResult struct {
	LastUpdateId         *int64     `json:"lastUpdateId,omitempty"`
	Bids                 [][]string `json:"bids,omitempty"`
	Asks                 [][]string `json:"asks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DepthResponseResult DepthResponseResult

// NewDepthResponseResult instantiates a new DepthResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepthResponseResult() *DepthResponseResult {
	this := DepthResponseResult{}
	return &this
}

// NewDepthResponseResultWithDefaults instantiates a new DepthResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepthResponseResultWithDefaults() *DepthResponseResult {
	this := DepthResponseResult{}
	return &this
}

// GetLastUpdateId returns the LastUpdateId field value if set, zero value otherwise.
func (o *DepthResponseResult) GetLastUpdateId() int64 {
	if o == nil || common.IsNil(o.LastUpdateId) {
		var ret int64
		return ret
	}
	return *o.LastUpdateId
}

// GetLastUpdateIdOk returns a tuple with the LastUpdateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepthResponseResult) GetLastUpdateIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastUpdateId) {
		return nil, false
	}
	return o.LastUpdateId, true
}

// HasLastUpdateId returns a boolean if a field has been set.
func (o *DepthResponseResult) HasLastUpdateId() bool {
	if o != nil && !common.IsNil(o.LastUpdateId) {
		return true
	}

	return false
}

// SetLastUpdateId gets a reference to the given int64 and assigns it to the LastUpdateId field.
func (o *DepthResponseResult) SetLastUpdateId(v int64) {
	o.LastUpdateId = &v
}

// GetBids returns the Bids field value if set, zero value otherwise.
func (o *DepthResponseResult) GetBids() [][]string {
	if o == nil || common.IsNil(o.Bids) {
		var ret [][]string
		return ret
	}
	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepthResponseResult) GetBidsOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Bids) {
		return nil, false
	}
	return o.Bids, true
}

// HasBids returns a boolean if a field has been set.
func (o *DepthResponseResult) HasBids() bool {
	if o != nil && !common.IsNil(o.Bids) {
		return true
	}

	return false
}

// SetBids gets a reference to the given [][]string and assigns it to the Bids field.
func (o *DepthResponseResult) SetBids(v [][]string) {
	o.Bids = v
}

// GetAsks returns the Asks field value if set, zero value otherwise.
func (o *DepthResponseResult) GetAsks() [][]string {
	if o == nil || common.IsNil(o.Asks) {
		var ret [][]string
		return ret
	}
	return o.Asks
}

// GetAsksOk returns a tuple with the Asks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepthResponseResult) GetAsksOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Asks) {
		return nil, false
	}
	return o.Asks, true
}

// HasAsks returns a boolean if a field has been set.
func (o *DepthResponseResult) HasAsks() bool {
	if o != nil && !common.IsNil(o.Asks) {
		return true
	}

	return false
}

// SetAsks gets a reference to the given [][]string and assigns it to the Asks field.
func (o *DepthResponseResult) SetAsks(v [][]string) {
	o.Asks = v
}

func (o DepthResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepthResponseResult) ToMap() (map[string]interface{}, error) {
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

func (o *DepthResponseResult) UnmarshalJSON(data []byte) (err error) {
	varDepthResponseResult := _DepthResponseResult{}

	err = json.Unmarshal(data, &varDepthResponseResult)

	if err != nil {
		return err
	}

	*o = DepthResponseResult(varDepthResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lastUpdateId")
		delete(additionalProperties, "bids")
		delete(additionalProperties, "asks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepthResponseResult struct {
	value *DepthResponseResult
	isSet bool
}

func (v NullableDepthResponseResult) Get() *DepthResponseResult {
	return v.value
}

func (v *NullableDepthResponseResult) Set(val *DepthResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDepthResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDepthResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepthResponseResult(val *DepthResponseResult) *NullableDepthResponseResult {
	return &NullableDepthResponseResult{value: val, isSet: true}
}

func (v NullableDepthResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepthResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
