/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BstocksCorporateActionRecordsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BstocksCorporateActionRecordsResponse{}

// BstocksCorporateActionRecordsResponse struct for BstocksCorporateActionRecordsResponse
type BstocksCorporateActionRecordsResponse struct {
	// Corporate action asset change records on this page.
	Items []CorporateActionRecordsResponseItemsInner `json:"items,omitempty"`
	// Pass this value as `lastTime` on the next request. Not returned on the last page.
	NextTime *int64 `json:"nextTime,omitempty"`
	// Pass this value as `lastId` on the next request. Not returned on the last page.
	NextId               *int64 `json:"nextId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BstocksCorporateActionRecordsResponse BstocksCorporateActionRecordsResponse

// NewBstocksCorporateActionRecordsResponse instantiates a new BstocksCorporateActionRecordsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBstocksCorporateActionRecordsResponse() *BstocksCorporateActionRecordsResponse {
	this := BstocksCorporateActionRecordsResponse{}
	return &this
}

// NewBstocksCorporateActionRecordsResponseWithDefaults instantiates a new BstocksCorporateActionRecordsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBstocksCorporateActionRecordsResponseWithDefaults() *BstocksCorporateActionRecordsResponse {
	this := BstocksCorporateActionRecordsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BstocksCorporateActionRecordsResponse) GetItems() []CorporateActionRecordsResponseItemsInner {
	if o == nil || common.IsNil(o.Items) {
		var ret []CorporateActionRecordsResponseItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BstocksCorporateActionRecordsResponse) GetItemsOk() ([]CorporateActionRecordsResponseItemsInner, bool) {
	if o == nil || common.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BstocksCorporateActionRecordsResponse) HasItems() bool {
	if o != nil && !common.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []CorporateActionRecordsResponseItemsInner and assigns it to the Items field.
func (o *BstocksCorporateActionRecordsResponse) SetItems(v []CorporateActionRecordsResponseItemsInner) {
	o.Items = v
}

// GetNextTime returns the NextTime field value if set, zero value otherwise.
func (o *BstocksCorporateActionRecordsResponse) GetNextTime() int64 {
	if o == nil || common.IsNil(o.NextTime) {
		var ret int64
		return ret
	}
	return *o.NextTime
}

// GetNextTimeOk returns a tuple with the NextTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BstocksCorporateActionRecordsResponse) GetNextTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.NextTime) {
		return nil, false
	}
	return o.NextTime, true
}

// HasNextTime returns a boolean if a field has been set.
func (o *BstocksCorporateActionRecordsResponse) HasNextTime() bool {
	if o != nil && !common.IsNil(o.NextTime) {
		return true
	}

	return false
}

// SetNextTime gets a reference to the given int64 and assigns it to the NextTime field.
func (o *BstocksCorporateActionRecordsResponse) SetNextTime(v int64) {
	o.NextTime = &v
}

// GetNextId returns the NextId field value if set, zero value otherwise.
func (o *BstocksCorporateActionRecordsResponse) GetNextId() int64 {
	if o == nil || common.IsNil(o.NextId) {
		var ret int64
		return ret
	}
	return *o.NextId
}

// GetNextIdOk returns a tuple with the NextId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BstocksCorporateActionRecordsResponse) GetNextIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.NextId) {
		return nil, false
	}
	return o.NextId, true
}

// HasNextId returns a boolean if a field has been set.
func (o *BstocksCorporateActionRecordsResponse) HasNextId() bool {
	if o != nil && !common.IsNil(o.NextId) {
		return true
	}

	return false
}

// SetNextId gets a reference to the given int64 and assigns it to the NextId field.
func (o *BstocksCorporateActionRecordsResponse) SetNextId(v int64) {
	o.NextId = &v
}

func (o BstocksCorporateActionRecordsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BstocksCorporateActionRecordsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !common.IsNil(o.NextTime) {
		toSerialize["nextTime"] = o.NextTime
	}
	if !common.IsNil(o.NextId) {
		toSerialize["nextId"] = o.NextId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BstocksCorporateActionRecordsResponse) UnmarshalJSON(data []byte) (err error) {
	varBstocksCorporateActionRecordsResponse := _BstocksCorporateActionRecordsResponse{}

	err = json.Unmarshal(data, &varBstocksCorporateActionRecordsResponse)

	if err != nil {
		return err
	}

	*o = BstocksCorporateActionRecordsResponse(varBstocksCorporateActionRecordsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "nextTime")
		delete(additionalProperties, "nextId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBstocksCorporateActionRecordsResponse struct {
	value *BstocksCorporateActionRecordsResponse
	isSet bool
}

func (v NullableBstocksCorporateActionRecordsResponse) Get() *BstocksCorporateActionRecordsResponse {
	return v.value
}

func (v *NullableBstocksCorporateActionRecordsResponse) Set(val *BstocksCorporateActionRecordsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBstocksCorporateActionRecordsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBstocksCorporateActionRecordsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBstocksCorporateActionRecordsResponse(val *BstocksCorporateActionRecordsResponse) *NullableBstocksCorporateActionRecordsResponse {
	return &NullableBstocksCorporateActionRecordsResponse{value: val, isSet: true}
}

func (v NullableBstocksCorporateActionRecordsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBstocksCorporateActionRecordsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
