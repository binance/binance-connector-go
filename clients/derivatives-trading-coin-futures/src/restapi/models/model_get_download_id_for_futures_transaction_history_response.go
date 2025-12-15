/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetDownloadIdForFuturesTransactionHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDownloadIdForFuturesTransactionHistoryResponse{}

// GetDownloadIdForFuturesTransactionHistoryResponse struct for GetDownloadIdForFuturesTransactionHistoryResponse
type GetDownloadIdForFuturesTransactionHistoryResponse struct {
	AvgCostTimestampOfLast30d *int64  `json:"avgCostTimestampOfLast30d,omitempty"`
	DownloadId                *string `json:"downloadId,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _GetDownloadIdForFuturesTransactionHistoryResponse GetDownloadIdForFuturesTransactionHistoryResponse

// NewGetDownloadIdForFuturesTransactionHistoryResponse instantiates a new GetDownloadIdForFuturesTransactionHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDownloadIdForFuturesTransactionHistoryResponse() *GetDownloadIdForFuturesTransactionHistoryResponse {
	this := GetDownloadIdForFuturesTransactionHistoryResponse{}
	return &this
}

// NewGetDownloadIdForFuturesTransactionHistoryResponseWithDefaults instantiates a new GetDownloadIdForFuturesTransactionHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDownloadIdForFuturesTransactionHistoryResponseWithDefaults() *GetDownloadIdForFuturesTransactionHistoryResponse {
	this := GetDownloadIdForFuturesTransactionHistoryResponse{}
	return &this
}

// GetAvgCostTimestampOfLast30d returns the AvgCostTimestampOfLast30d field value if set, zero value otherwise.
func (o *GetDownloadIdForFuturesTransactionHistoryResponse) GetAvgCostTimestampOfLast30d() int64 {
	if o == nil || common.IsNil(o.AvgCostTimestampOfLast30d) {
		var ret int64
		return ret
	}
	return *o.AvgCostTimestampOfLast30d
}

// GetAvgCostTimestampOfLast30dOk returns a tuple with the AvgCostTimestampOfLast30d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDownloadIdForFuturesTransactionHistoryResponse) GetAvgCostTimestampOfLast30dOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AvgCostTimestampOfLast30d) {
		return nil, false
	}
	return o.AvgCostTimestampOfLast30d, true
}

// HasAvgCostTimestampOfLast30d returns a boolean if a field has been set.
func (o *GetDownloadIdForFuturesTransactionHistoryResponse) HasAvgCostTimestampOfLast30d() bool {
	if o != nil && !common.IsNil(o.AvgCostTimestampOfLast30d) {
		return true
	}

	return false
}

// SetAvgCostTimestampOfLast30d gets a reference to the given int64 and assigns it to the AvgCostTimestampOfLast30d field.
func (o *GetDownloadIdForFuturesTransactionHistoryResponse) SetAvgCostTimestampOfLast30d(v int64) {
	o.AvgCostTimestampOfLast30d = &v
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise.
func (o *GetDownloadIdForFuturesTransactionHistoryResponse) GetDownloadId() string {
	if o == nil || common.IsNil(o.DownloadId) {
		var ret string
		return ret
	}
	return *o.DownloadId
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDownloadIdForFuturesTransactionHistoryResponse) GetDownloadIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.DownloadId) {
		return nil, false
	}
	return o.DownloadId, true
}

// HasDownloadId returns a boolean if a field has been set.
func (o *GetDownloadIdForFuturesTransactionHistoryResponse) HasDownloadId() bool {
	if o != nil && !common.IsNil(o.DownloadId) {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given string and assigns it to the DownloadId field.
func (o *GetDownloadIdForFuturesTransactionHistoryResponse) SetDownloadId(v string) {
	o.DownloadId = &v
}

func (o GetDownloadIdForFuturesTransactionHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDownloadIdForFuturesTransactionHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetDownloadIdForFuturesTransactionHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetDownloadIdForFuturesTransactionHistoryResponse := _GetDownloadIdForFuturesTransactionHistoryResponse{}

	err = json.Unmarshal(data, &varGetDownloadIdForFuturesTransactionHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetDownloadIdForFuturesTransactionHistoryResponse(varGetDownloadIdForFuturesTransactionHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "avgCostTimestampOfLast30d")
		delete(additionalProperties, "downloadId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDownloadIdForFuturesTransactionHistoryResponse struct {
	value *GetDownloadIdForFuturesTransactionHistoryResponse
	isSet bool
}

func (v NullableGetDownloadIdForFuturesTransactionHistoryResponse) Get() *GetDownloadIdForFuturesTransactionHistoryResponse {
	return v.value
}

func (v *NullableGetDownloadIdForFuturesTransactionHistoryResponse) Set(val *GetDownloadIdForFuturesTransactionHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDownloadIdForFuturesTransactionHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDownloadIdForFuturesTransactionHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDownloadIdForFuturesTransactionHistoryResponse(val *GetDownloadIdForFuturesTransactionHistoryResponse) *NullableGetDownloadIdForFuturesTransactionHistoryResponse {
	return &NullableGetDownloadIdForFuturesTransactionHistoryResponse{value: val, isSet: true}
}

func (v NullableGetDownloadIdForFuturesTransactionHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDownloadIdForFuturesTransactionHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
