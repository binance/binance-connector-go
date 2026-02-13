/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RpiOrderBookResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RpiOrderBookResponse{}

// RpiOrderBookResponse struct for RpiOrderBookResponse
type RpiOrderBookResponse struct {
	LastUpdateId         *int64                         `json:"lastUpdateId,omitempty"`
	E                    *int64                         `json:"E,omitempty"`
	T                    *int64                         `json:"T,omitempty"`
	Bids                 []RpiOrderBookResponseBidsItem `json:"bids,omitempty"`
	Asks                 []RpiOrderBookResponseAsksItem `json:"asks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RpiOrderBookResponse RpiOrderBookResponse

// NewRpiOrderBookResponse instantiates a new RpiOrderBookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpiOrderBookResponse() *RpiOrderBookResponse {
	this := RpiOrderBookResponse{}
	return &this
}

// NewRpiOrderBookResponseWithDefaults instantiates a new RpiOrderBookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpiOrderBookResponseWithDefaults() *RpiOrderBookResponse {
	this := RpiOrderBookResponse{}
	return &this
}

// GetLastUpdateId returns the LastUpdateId field value if set, zero value otherwise.
func (o *RpiOrderBookResponse) GetLastUpdateId() int64 {
	if o == nil || common.IsNil(o.LastUpdateId) {
		var ret int64
		return ret
	}
	return *o.LastUpdateId
}

// GetLastUpdateIdOk returns a tuple with the LastUpdateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiOrderBookResponse) GetLastUpdateIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastUpdateId) {
		return nil, false
	}
	return o.LastUpdateId, true
}

// HasLastUpdateId returns a boolean if a field has been set.
func (o *RpiOrderBookResponse) HasLastUpdateId() bool {
	if o != nil && !common.IsNil(o.LastUpdateId) {
		return true
	}

	return false
}

// SetLastUpdateId gets a reference to the given int64 and assigns it to the LastUpdateId field.
func (o *RpiOrderBookResponse) SetLastUpdateId(v int64) {
	o.LastUpdateId = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *RpiOrderBookResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiOrderBookResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *RpiOrderBookResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *RpiOrderBookResponse) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *RpiOrderBookResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiOrderBookResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *RpiOrderBookResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *RpiOrderBookResponse) SetT(v int64) {
	o.T = &v
}

// GetBids returns the Bids field value if set, zero value otherwise.
func (o *RpiOrderBookResponse) GetBids() []RpiOrderBookResponseBidsItem {
	if o == nil || common.IsNil(o.Bids) {
		var ret []RpiOrderBookResponseBidsItem
		return ret
	}
	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiOrderBookResponse) GetBidsOk() ([]RpiOrderBookResponseBidsItem, bool) {
	if o == nil || common.IsNil(o.Bids) {
		return nil, false
	}
	return o.Bids, true
}

// HasBids returns a boolean if a field has been set.
func (o *RpiOrderBookResponse) HasBids() bool {
	if o != nil && !common.IsNil(o.Bids) {
		return true
	}

	return false
}

// SetBids gets a reference to the given []RpiOrderBookResponseBidsItem and assigns it to the Bids field.
func (o *RpiOrderBookResponse) SetBids(v []RpiOrderBookResponseBidsItem) {
	o.Bids = v
}

// GetAsks returns the Asks field value if set, zero value otherwise.
func (o *RpiOrderBookResponse) GetAsks() []RpiOrderBookResponseAsksItem {
	if o == nil || common.IsNil(o.Asks) {
		var ret []RpiOrderBookResponseAsksItem
		return ret
	}
	return o.Asks
}

// GetAsksOk returns a tuple with the Asks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiOrderBookResponse) GetAsksOk() ([]RpiOrderBookResponseAsksItem, bool) {
	if o == nil || common.IsNil(o.Asks) {
		return nil, false
	}
	return o.Asks, true
}

// HasAsks returns a boolean if a field has been set.
func (o *RpiOrderBookResponse) HasAsks() bool {
	if o != nil && !common.IsNil(o.Asks) {
		return true
	}

	return false
}

// SetAsks gets a reference to the given []RpiOrderBookResponseAsksItem and assigns it to the Asks field.
func (o *RpiOrderBookResponse) SetAsks(v []RpiOrderBookResponseAsksItem) {
	o.Asks = v
}

func (o RpiOrderBookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpiOrderBookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LastUpdateId) {
		toSerialize["lastUpdateId"] = o.LastUpdateId
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
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

func (o *RpiOrderBookResponse) UnmarshalJSON(data []byte) (err error) {
	varRpiOrderBookResponse := _RpiOrderBookResponse{}

	err = json.Unmarshal(data, &varRpiOrderBookResponse)

	if err != nil {
		return err
	}

	*o = RpiOrderBookResponse(varRpiOrderBookResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lastUpdateId")
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "bids")
		delete(additionalProperties, "asks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRpiOrderBookResponse struct {
	value *RpiOrderBookResponse
	isSet bool
}

func (v NullableRpiOrderBookResponse) Get() *RpiOrderBookResponse {
	return v.value
}

func (v *NullableRpiOrderBookResponse) Set(val *RpiOrderBookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRpiOrderBookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRpiOrderBookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpiOrderBookResponse(val *RpiOrderBookResponse) *NullableRpiOrderBookResponse {
	return &NullableRpiOrderBookResponse{value: val, isSet: true}
}

func (v NullableRpiOrderBookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpiOrderBookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
