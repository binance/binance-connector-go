/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetDownloadIdForUmFuturesOrderHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDownloadIdForUmFuturesOrderHistoryResponse{}

// GetDownloadIdForUmFuturesOrderHistoryResponse struct for GetDownloadIdForUmFuturesOrderHistoryResponse
type GetDownloadIdForUmFuturesOrderHistoryResponse struct {
	AvgCostTimestampOfLast30d *int64  `json:"avgCostTimestampOfLast30d,omitempty"`
	DownloadId                *string `json:"downloadId,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _GetDownloadIdForUmFuturesOrderHistoryResponse GetDownloadIdForUmFuturesOrderHistoryResponse

// NewGetDownloadIdForUmFuturesOrderHistoryResponse instantiates a new GetDownloadIdForUmFuturesOrderHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDownloadIdForUmFuturesOrderHistoryResponse() *GetDownloadIdForUmFuturesOrderHistoryResponse {
	this := GetDownloadIdForUmFuturesOrderHistoryResponse{}
	return &this
}

// NewGetDownloadIdForUmFuturesOrderHistoryResponseWithDefaults instantiates a new GetDownloadIdForUmFuturesOrderHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDownloadIdForUmFuturesOrderHistoryResponseWithDefaults() *GetDownloadIdForUmFuturesOrderHistoryResponse {
	this := GetDownloadIdForUmFuturesOrderHistoryResponse{}
	return &this
}

// GetAvgCostTimestampOfLast30d returns the AvgCostTimestampOfLast30d field value if set, zero value otherwise.
func (o *GetDownloadIdForUmFuturesOrderHistoryResponse) GetAvgCostTimestampOfLast30d() int64 {
	if o == nil || common.IsNil(o.AvgCostTimestampOfLast30d) {
		var ret int64
		return ret
	}
	return *o.AvgCostTimestampOfLast30d
}

// GetAvgCostTimestampOfLast30dOk returns a tuple with the AvgCostTimestampOfLast30d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDownloadIdForUmFuturesOrderHistoryResponse) GetAvgCostTimestampOfLast30dOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AvgCostTimestampOfLast30d) {
		return nil, false
	}
	return o.AvgCostTimestampOfLast30d, true
}

// HasAvgCostTimestampOfLast30d returns a boolean if a field has been set.
func (o *GetDownloadIdForUmFuturesOrderHistoryResponse) HasAvgCostTimestampOfLast30d() bool {
	if o != nil && !common.IsNil(o.AvgCostTimestampOfLast30d) {
		return true
	}

	return false
}

// SetAvgCostTimestampOfLast30d gets a reference to the given int64 and assigns it to the AvgCostTimestampOfLast30d field.
func (o *GetDownloadIdForUmFuturesOrderHistoryResponse) SetAvgCostTimestampOfLast30d(v int64) {
	o.AvgCostTimestampOfLast30d = &v
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise.
func (o *GetDownloadIdForUmFuturesOrderHistoryResponse) GetDownloadId() string {
	if o == nil || common.IsNil(o.DownloadId) {
		var ret string
		return ret
	}
	return *o.DownloadId
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDownloadIdForUmFuturesOrderHistoryResponse) GetDownloadIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.DownloadId) {
		return nil, false
	}
	return o.DownloadId, true
}

// HasDownloadId returns a boolean if a field has been set.
func (o *GetDownloadIdForUmFuturesOrderHistoryResponse) HasDownloadId() bool {
	if o != nil && !common.IsNil(o.DownloadId) {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given string and assigns it to the DownloadId field.
func (o *GetDownloadIdForUmFuturesOrderHistoryResponse) SetDownloadId(v string) {
	o.DownloadId = &v
}

func (o GetDownloadIdForUmFuturesOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDownloadIdForUmFuturesOrderHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AvgCostTimestampOfLast30d) {
		toSerialize["avgCostTimestampOfLast30d"] = o.AvgCostTimestampOfLast30d
	}
	if !common.IsNil(o.DownloadId) {
		toSerialize["downloadId"] = o.DownloadId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDownloadIdForUmFuturesOrderHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetDownloadIdForUmFuturesOrderHistoryResponse := _GetDownloadIdForUmFuturesOrderHistoryResponse{}

	err = json.Unmarshal(data, &varGetDownloadIdForUmFuturesOrderHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetDownloadIdForUmFuturesOrderHistoryResponse(varGetDownloadIdForUmFuturesOrderHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "avgCostTimestampOfLast30d")
		delete(additionalProperties, "downloadId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDownloadIdForUmFuturesOrderHistoryResponse struct {
	value *GetDownloadIdForUmFuturesOrderHistoryResponse
	isSet bool
}

func (v NullableGetDownloadIdForUmFuturesOrderHistoryResponse) Get() *GetDownloadIdForUmFuturesOrderHistoryResponse {
	return v.value
}

func (v *NullableGetDownloadIdForUmFuturesOrderHistoryResponse) Set(val *GetDownloadIdForUmFuturesOrderHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDownloadIdForUmFuturesOrderHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDownloadIdForUmFuturesOrderHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDownloadIdForUmFuturesOrderHistoryResponse(val *GetDownloadIdForUmFuturesOrderHistoryResponse) *NullableGetDownloadIdForUmFuturesOrderHistoryResponse {
	return &NullableGetDownloadIdForUmFuturesOrderHistoryResponse{value: val, isSet: true}
}

func (v NullableGetDownloadIdForUmFuturesOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDownloadIdForUmFuturesOrderHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
