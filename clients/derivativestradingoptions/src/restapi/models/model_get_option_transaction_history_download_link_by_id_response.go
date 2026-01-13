/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetOptionTransactionHistoryDownloadLinkByIdResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOptionTransactionHistoryDownloadLinkByIdResponse{}

// GetOptionTransactionHistoryDownloadLinkByIdResponse struct for GetOptionTransactionHistoryDownloadLinkByIdResponse
type GetOptionTransactionHistoryDownloadLinkByIdResponse struct {
	DownloadId           *string `json:"downloadId,omitempty"`
	Status               *string `json:"status,omitempty"`
	Url                  *string `json:"url,omitempty"`
	Notified             *bool   `json:"notified,omitempty"`
	ExpirationTimestamp  *int64  `json:"expirationTimestamp,omitempty"`
	IsExpired            *string `json:"isExpired,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOptionTransactionHistoryDownloadLinkByIdResponse GetOptionTransactionHistoryDownloadLinkByIdResponse

// NewGetOptionTransactionHistoryDownloadLinkByIdResponse instantiates a new GetOptionTransactionHistoryDownloadLinkByIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOptionTransactionHistoryDownloadLinkByIdResponse() *GetOptionTransactionHistoryDownloadLinkByIdResponse {
	this := GetOptionTransactionHistoryDownloadLinkByIdResponse{}
	return &this
}

// NewGetOptionTransactionHistoryDownloadLinkByIdResponseWithDefaults instantiates a new GetOptionTransactionHistoryDownloadLinkByIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOptionTransactionHistoryDownloadLinkByIdResponseWithDefaults() *GetOptionTransactionHistoryDownloadLinkByIdResponse {
	this := GetOptionTransactionHistoryDownloadLinkByIdResponse{}
	return &this
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) GetDownloadId() string {
	if o == nil || common.IsNil(o.DownloadId) {
		var ret string
		return ret
	}
	return *o.DownloadId
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) GetDownloadIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.DownloadId) {
		return nil, false
	}
	return o.DownloadId, true
}

// HasDownloadId returns a boolean if a field has been set.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) HasDownloadId() bool {
	if o != nil && !common.IsNil(o.DownloadId) {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given string and assigns it to the DownloadId field.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) SetDownloadId(v string) {
	o.DownloadId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) SetUrl(v string) {
	o.Url = &v
}

// GetNotified returns the Notified field value if set, zero value otherwise.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) GetNotified() bool {
	if o == nil || common.IsNil(o.Notified) {
		var ret bool
		return ret
	}
	return *o.Notified
}

// GetNotifiedOk returns a tuple with the Notified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) GetNotifiedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Notified) {
		return nil, false
	}
	return o.Notified, true
}

// HasNotified returns a boolean if a field has been set.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) HasNotified() bool {
	if o != nil && !common.IsNil(o.Notified) {
		return true
	}

	return false
}

// SetNotified gets a reference to the given bool and assigns it to the Notified field.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) SetNotified(v bool) {
	o.Notified = &v
}

// GetExpirationTimestamp returns the ExpirationTimestamp field value if set, zero value otherwise.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) GetExpirationTimestamp() int64 {
	if o == nil || common.IsNil(o.ExpirationTimestamp) {
		var ret int64
		return ret
	}
	return *o.ExpirationTimestamp
}

// GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) GetExpirationTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpirationTimestamp) {
		return nil, false
	}
	return o.ExpirationTimestamp, true
}

// HasExpirationTimestamp returns a boolean if a field has been set.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) HasExpirationTimestamp() bool {
	if o != nil && !common.IsNil(o.ExpirationTimestamp) {
		return true
	}

	return false
}

// SetExpirationTimestamp gets a reference to the given int64 and assigns it to the ExpirationTimestamp field.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) SetExpirationTimestamp(v int64) {
	o.ExpirationTimestamp = &v
}

// GetIsExpired returns the IsExpired field value if set, zero value otherwise.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) GetIsExpired() string {
	if o == nil || common.IsNil(o.IsExpired) {
		var ret string
		return ret
	}
	return *o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) GetIsExpiredOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsExpired) {
		return nil, false
	}
	return o.IsExpired, true
}

// HasIsExpired returns a boolean if a field has been set.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) HasIsExpired() bool {
	if o != nil && !common.IsNil(o.IsExpired) {
		return true
	}

	return false
}

// SetIsExpired gets a reference to the given string and assigns it to the IsExpired field.
func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) SetIsExpired(v string) {
	o.IsExpired = &v
}

func (o GetOptionTransactionHistoryDownloadLinkByIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOptionTransactionHistoryDownloadLinkByIdResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetOptionTransactionHistoryDownloadLinkByIdResponse) UnmarshalJSON(data []byte) (err error) {
	varGetOptionTransactionHistoryDownloadLinkByIdResponse := _GetOptionTransactionHistoryDownloadLinkByIdResponse{}

	err = json.Unmarshal(data, &varGetOptionTransactionHistoryDownloadLinkByIdResponse)

	if err != nil {
		return err
	}

	*o = GetOptionTransactionHistoryDownloadLinkByIdResponse(varGetOptionTransactionHistoryDownloadLinkByIdResponse)

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

type NullableGetOptionTransactionHistoryDownloadLinkByIdResponse struct {
	value *GetOptionTransactionHistoryDownloadLinkByIdResponse
	isSet bool
}

func (v NullableGetOptionTransactionHistoryDownloadLinkByIdResponse) Get() *GetOptionTransactionHistoryDownloadLinkByIdResponse {
	return v.value
}

func (v *NullableGetOptionTransactionHistoryDownloadLinkByIdResponse) Set(val *GetOptionTransactionHistoryDownloadLinkByIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOptionTransactionHistoryDownloadLinkByIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOptionTransactionHistoryDownloadLinkByIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOptionTransactionHistoryDownloadLinkByIdResponse(val *GetOptionTransactionHistoryDownloadLinkByIdResponse) *NullableGetOptionTransactionHistoryDownloadLinkByIdResponse {
	return &NullableGetOptionTransactionHistoryDownloadLinkByIdResponse{value: val, isSet: true}
}

func (v NullableGetOptionTransactionHistoryDownloadLinkByIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOptionTransactionHistoryDownloadLinkByIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
