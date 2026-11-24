/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetDownloadIdForCmFuturesTradeHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDownloadIdForCmFuturesTradeHistoryResponse{}

// GetDownloadIdForCmFuturesTradeHistoryResponse struct for GetDownloadIdForCmFuturesTradeHistoryResponse
type GetDownloadIdForCmFuturesTradeHistoryResponse struct {
	// Average time taken for data download in the past 30 days
	AvgCostTimestampOfLast30d *int64 `json:"avgCostTimestampOfLast30d,omitempty"`
	// Download ID.
	DownloadId           *string `json:"downloadId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDownloadIdForCmFuturesTradeHistoryResponse GetDownloadIdForCmFuturesTradeHistoryResponse

// NewGetDownloadIdForCmFuturesTradeHistoryResponse instantiates a new GetDownloadIdForCmFuturesTradeHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDownloadIdForCmFuturesTradeHistoryResponse() *GetDownloadIdForCmFuturesTradeHistoryResponse {
	this := GetDownloadIdForCmFuturesTradeHistoryResponse{}
	return &this
}

// NewGetDownloadIdForCmFuturesTradeHistoryResponseWithDefaults instantiates a new GetDownloadIdForCmFuturesTradeHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDownloadIdForCmFuturesTradeHistoryResponseWithDefaults() *GetDownloadIdForCmFuturesTradeHistoryResponse {
	this := GetDownloadIdForCmFuturesTradeHistoryResponse{}
	return &this
}

// GetAvgCostTimestampOfLast30d returns the AvgCostTimestampOfLast30d field value if set, zero value otherwise.
func (o *GetDownloadIdForCmFuturesTradeHistoryResponse) GetAvgCostTimestampOfLast30d() int64 {
	if o == nil || common.IsNil(o.AvgCostTimestampOfLast30d) {
		var ret int64
		return ret
	}
	return *o.AvgCostTimestampOfLast30d
}

// GetAvgCostTimestampOfLast30dOk returns a tuple with the AvgCostTimestampOfLast30d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDownloadIdForCmFuturesTradeHistoryResponse) GetAvgCostTimestampOfLast30dOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AvgCostTimestampOfLast30d) {
		return nil, false
	}
	return o.AvgCostTimestampOfLast30d, true
}

// HasAvgCostTimestampOfLast30d returns a boolean if a field has been set.
func (o *GetDownloadIdForCmFuturesTradeHistoryResponse) HasAvgCostTimestampOfLast30d() bool {
	if o != nil && !common.IsNil(o.AvgCostTimestampOfLast30d) {
		return true
	}

	return false
}

// SetAvgCostTimestampOfLast30d gets a reference to the given int64 and assigns it to the AvgCostTimestampOfLast30d field.
func (o *GetDownloadIdForCmFuturesTradeHistoryResponse) SetAvgCostTimestampOfLast30d(v int64) {
	o.AvgCostTimestampOfLast30d = &v
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise.
func (o *GetDownloadIdForCmFuturesTradeHistoryResponse) GetDownloadId() string {
	if o == nil || common.IsNil(o.DownloadId) {
		var ret string
		return ret
	}
	return *o.DownloadId
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDownloadIdForCmFuturesTradeHistoryResponse) GetDownloadIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.DownloadId) {
		return nil, false
	}
	return o.DownloadId, true
}

// HasDownloadId returns a boolean if a field has been set.
func (o *GetDownloadIdForCmFuturesTradeHistoryResponse) HasDownloadId() bool {
	if o != nil && !common.IsNil(o.DownloadId) {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given string and assigns it to the DownloadId field.
func (o *GetDownloadIdForCmFuturesTradeHistoryResponse) SetDownloadId(v string) {
	o.DownloadId = &v
}

func (o GetDownloadIdForCmFuturesTradeHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDownloadIdForCmFuturesTradeHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetDownloadIdForCmFuturesTradeHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetDownloadIdForCmFuturesTradeHistoryResponse := _GetDownloadIdForCmFuturesTradeHistoryResponse{}

	err = json.Unmarshal(data, &varGetDownloadIdForCmFuturesTradeHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetDownloadIdForCmFuturesTradeHistoryResponse(varGetDownloadIdForCmFuturesTradeHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "avgCostTimestampOfLast30d")
		delete(additionalProperties, "downloadId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDownloadIdForCmFuturesTradeHistoryResponse struct {
	value *GetDownloadIdForCmFuturesTradeHistoryResponse
	isSet bool
}

func (v NullableGetDownloadIdForCmFuturesTradeHistoryResponse) Get() *GetDownloadIdForCmFuturesTradeHistoryResponse {
	return v.value
}

func (v *NullableGetDownloadIdForCmFuturesTradeHistoryResponse) Set(val *GetDownloadIdForCmFuturesTradeHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDownloadIdForCmFuturesTradeHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDownloadIdForCmFuturesTradeHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDownloadIdForCmFuturesTradeHistoryResponse(val *GetDownloadIdForCmFuturesTradeHistoryResponse) *NullableGetDownloadIdForCmFuturesTradeHistoryResponse {
	return &NullableGetDownloadIdForCmFuturesTradeHistoryResponse{value: val, isSet: true}
}

func (v NullableGetDownloadIdForCmFuturesTradeHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDownloadIdForCmFuturesTradeHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
