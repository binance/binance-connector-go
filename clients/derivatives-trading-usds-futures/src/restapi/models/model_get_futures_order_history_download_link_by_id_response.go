/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFuturesOrderHistoryDownloadLinkByIdResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFuturesOrderHistoryDownloadLinkByIdResponse{}

// GetFuturesOrderHistoryDownloadLinkByIdResponse struct for GetFuturesOrderHistoryDownloadLinkByIdResponse
type GetFuturesOrderHistoryDownloadLinkByIdResponse struct {
	DownloadId           *string `json:"downloadId,omitempty"`
	Status               *string `json:"status,omitempty"`
	Url                  *string `json:"url,omitempty"`
	Notified             *bool   `json:"notified,omitempty"`
	ExpirationTimestamp  *int64  `json:"expirationTimestamp,omitempty"`
	IsExpired            *string `json:"isExpired,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFuturesOrderHistoryDownloadLinkByIdResponse GetFuturesOrderHistoryDownloadLinkByIdResponse

// NewGetFuturesOrderHistoryDownloadLinkByIdResponse instantiates a new GetFuturesOrderHistoryDownloadLinkByIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesOrderHistoryDownloadLinkByIdResponse() *GetFuturesOrderHistoryDownloadLinkByIdResponse {
	this := GetFuturesOrderHistoryDownloadLinkByIdResponse{}
	return &this
}

// NewGetFuturesOrderHistoryDownloadLinkByIdResponseWithDefaults instantiates a new GetFuturesOrderHistoryDownloadLinkByIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesOrderHistoryDownloadLinkByIdResponseWithDefaults() *GetFuturesOrderHistoryDownloadLinkByIdResponse {
	this := GetFuturesOrderHistoryDownloadLinkByIdResponse{}
	return &this
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) GetDownloadId() string {
	if o == nil || common.IsNil(o.DownloadId) {
		var ret string
		return ret
	}
	return *o.DownloadId
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) GetDownloadIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.DownloadId) {
		return nil, false
	}
	return o.DownloadId, true
}

// HasDownloadId returns a boolean if a field has been set.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) HasDownloadId() bool {
	if o != nil && !common.IsNil(o.DownloadId) {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given string and assigns it to the DownloadId field.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) SetDownloadId(v string) {
	o.DownloadId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) SetUrl(v string) {
	o.Url = &v
}

// GetNotified returns the Notified field value if set, zero value otherwise.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) GetNotified() bool {
	if o == nil || common.IsNil(o.Notified) {
		var ret bool
		return ret
	}
	return *o.Notified
}

// GetNotifiedOk returns a tuple with the Notified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) GetNotifiedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Notified) {
		return nil, false
	}
	return o.Notified, true
}

// HasNotified returns a boolean if a field has been set.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) HasNotified() bool {
	if o != nil && !common.IsNil(o.Notified) {
		return true
	}

	return false
}

// SetNotified gets a reference to the given bool and assigns it to the Notified field.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) SetNotified(v bool) {
	o.Notified = &v
}

// GetExpirationTimestamp returns the ExpirationTimestamp field value if set, zero value otherwise.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) GetExpirationTimestamp() int64 {
	if o == nil || common.IsNil(o.ExpirationTimestamp) {
		var ret int64
		return ret
	}
	return *o.ExpirationTimestamp
}

// GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) GetExpirationTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpirationTimestamp) {
		return nil, false
	}
	return o.ExpirationTimestamp, true
}

// HasExpirationTimestamp returns a boolean if a field has been set.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) HasExpirationTimestamp() bool {
	if o != nil && !common.IsNil(o.ExpirationTimestamp) {
		return true
	}

	return false
}

// SetExpirationTimestamp gets a reference to the given int64 and assigns it to the ExpirationTimestamp field.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) SetExpirationTimestamp(v int64) {
	o.ExpirationTimestamp = &v
}

// GetIsExpired returns the IsExpired field value if set, zero value otherwise.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) GetIsExpired() string {
	if o == nil || common.IsNil(o.IsExpired) {
		var ret string
		return ret
	}
	return *o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) GetIsExpiredOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsExpired) {
		return nil, false
	}
	return o.IsExpired, true
}

// HasIsExpired returns a boolean if a field has been set.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) HasIsExpired() bool {
	if o != nil && !common.IsNil(o.IsExpired) {
		return true
	}

	return false
}

// SetIsExpired gets a reference to the given string and assigns it to the IsExpired field.
func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) SetIsExpired(v string) {
	o.IsExpired = &v
}

func (o GetFuturesOrderHistoryDownloadLinkByIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesOrderHistoryDownloadLinkByIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DownloadId) {
		toSerialize["downloadId"] = o.DownloadId
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !common.IsNil(o.Notified) {
		toSerialize["notified"] = o.Notified
	}
	if !common.IsNil(o.ExpirationTimestamp) {
		toSerialize["expirationTimestamp"] = o.ExpirationTimestamp
	}
	if !common.IsNil(o.IsExpired) {
		toSerialize["isExpired"] = o.IsExpired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFuturesOrderHistoryDownloadLinkByIdResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFuturesOrderHistoryDownloadLinkByIdResponse := _GetFuturesOrderHistoryDownloadLinkByIdResponse{}

	err = json.Unmarshal(data, &varGetFuturesOrderHistoryDownloadLinkByIdResponse)

	if err != nil {
		return err
	}

	*o = GetFuturesOrderHistoryDownloadLinkByIdResponse(varGetFuturesOrderHistoryDownloadLinkByIdResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "downloadId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "url")
		delete(additionalProperties, "notified")
		delete(additionalProperties, "expirationTimestamp")
		delete(additionalProperties, "isExpired")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFuturesOrderHistoryDownloadLinkByIdResponse struct {
	value *GetFuturesOrderHistoryDownloadLinkByIdResponse
	isSet bool
}

func (v NullableGetFuturesOrderHistoryDownloadLinkByIdResponse) Get() *GetFuturesOrderHistoryDownloadLinkByIdResponse {
	return v.value
}

func (v *NullableGetFuturesOrderHistoryDownloadLinkByIdResponse) Set(val *GetFuturesOrderHistoryDownloadLinkByIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesOrderHistoryDownloadLinkByIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesOrderHistoryDownloadLinkByIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesOrderHistoryDownloadLinkByIdResponse(val *GetFuturesOrderHistoryDownloadLinkByIdResponse) *NullableGetFuturesOrderHistoryDownloadLinkByIdResponse {
	return &NullableGetFuturesOrderHistoryDownloadLinkByIdResponse{value: val, isSet: true}
}

func (v NullableGetFuturesOrderHistoryDownloadLinkByIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesOrderHistoryDownloadLinkByIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
