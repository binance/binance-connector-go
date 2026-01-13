/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetUmFuturesTradeDownloadLinkByIdResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUmFuturesTradeDownloadLinkByIdResponse{}

// GetUmFuturesTradeDownloadLinkByIdResponse struct for GetUmFuturesTradeDownloadLinkByIdResponse
type GetUmFuturesTradeDownloadLinkByIdResponse struct {
	DownloadId           *string `json:"downloadId,omitempty"`
	Status               *string `json:"status,omitempty"`
	Url                  *string `json:"url,omitempty"`
	S3Link               *string `json:"s3Link,omitempty"`
	Notified             *bool   `json:"notified,omitempty"`
	ExpirationTimestamp  *int64  `json:"expirationTimestamp,omitempty"`
	IsExpired            *string `json:"isExpired,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUmFuturesTradeDownloadLinkByIdResponse GetUmFuturesTradeDownloadLinkByIdResponse

// NewGetUmFuturesTradeDownloadLinkByIdResponse instantiates a new GetUmFuturesTradeDownloadLinkByIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUmFuturesTradeDownloadLinkByIdResponse() *GetUmFuturesTradeDownloadLinkByIdResponse {
	this := GetUmFuturesTradeDownloadLinkByIdResponse{}
	return &this
}

// NewGetUmFuturesTradeDownloadLinkByIdResponseWithDefaults instantiates a new GetUmFuturesTradeDownloadLinkByIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUmFuturesTradeDownloadLinkByIdResponseWithDefaults() *GetUmFuturesTradeDownloadLinkByIdResponse {
	this := GetUmFuturesTradeDownloadLinkByIdResponse{}
	return &this
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetDownloadId() string {
	if o == nil || common.IsNil(o.DownloadId) {
		var ret string
		return ret
	}
	return *o.DownloadId
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetDownloadIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.DownloadId) {
		return nil, false
	}
	return o.DownloadId, true
}

// HasDownloadId returns a boolean if a field has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) HasDownloadId() bool {
	if o != nil && !common.IsNil(o.DownloadId) {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given string and assigns it to the DownloadId field.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) SetDownloadId(v string) {
	o.DownloadId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) SetUrl(v string) {
	o.Url = &v
}

// GetS3Link returns the S3Link field value if set, zero value otherwise.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetS3Link() string {
	if o == nil || common.IsNil(o.S3Link) {
		var ret string
		return ret
	}
	return *o.S3Link
}

// GetS3LinkOk returns a tuple with the S3Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetS3LinkOk() (*string, bool) {
	if o == nil || common.IsNil(o.S3Link) {
		return nil, false
	}
	return o.S3Link, true
}

// HasS3Link returns a boolean if a field has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) HasS3Link() bool {
	if o != nil && !common.IsNil(o.S3Link) {
		return true
	}

	return false
}

// SetS3Link gets a reference to the given string and assigns it to the S3Link field.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) SetS3Link(v string) {
	o.S3Link = &v
}

// GetNotified returns the Notified field value if set, zero value otherwise.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetNotified() bool {
	if o == nil || common.IsNil(o.Notified) {
		var ret bool
		return ret
	}
	return *o.Notified
}

// GetNotifiedOk returns a tuple with the Notified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetNotifiedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Notified) {
		return nil, false
	}
	return o.Notified, true
}

// HasNotified returns a boolean if a field has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) HasNotified() bool {
	if o != nil && !common.IsNil(o.Notified) {
		return true
	}

	return false
}

// SetNotified gets a reference to the given bool and assigns it to the Notified field.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) SetNotified(v bool) {
	o.Notified = &v
}

// GetExpirationTimestamp returns the ExpirationTimestamp field value if set, zero value otherwise.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetExpirationTimestamp() int64 {
	if o == nil || common.IsNil(o.ExpirationTimestamp) {
		var ret int64
		return ret
	}
	return *o.ExpirationTimestamp
}

// GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetExpirationTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpirationTimestamp) {
		return nil, false
	}
	return o.ExpirationTimestamp, true
}

// HasExpirationTimestamp returns a boolean if a field has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) HasExpirationTimestamp() bool {
	if o != nil && !common.IsNil(o.ExpirationTimestamp) {
		return true
	}

	return false
}

// SetExpirationTimestamp gets a reference to the given int64 and assigns it to the ExpirationTimestamp field.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) SetExpirationTimestamp(v int64) {
	o.ExpirationTimestamp = &v
}

// GetIsExpired returns the IsExpired field value if set, zero value otherwise.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetIsExpired() string {
	if o == nil || common.IsNil(o.IsExpired) {
		var ret string
		return ret
	}
	return *o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) GetIsExpiredOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsExpired) {
		return nil, false
	}
	return o.IsExpired, true
}

// HasIsExpired returns a boolean if a field has been set.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) HasIsExpired() bool {
	if o != nil && !common.IsNil(o.IsExpired) {
		return true
	}

	return false
}

// SetIsExpired gets a reference to the given string and assigns it to the IsExpired field.
func (o *GetUmFuturesTradeDownloadLinkByIdResponse) SetIsExpired(v string) {
	o.IsExpired = &v
}

func (o GetUmFuturesTradeDownloadLinkByIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUmFuturesTradeDownloadLinkByIdResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.S3Link) {
		toSerialize["s3Link"] = o.S3Link
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

func (o *GetUmFuturesTradeDownloadLinkByIdResponse) UnmarshalJSON(data []byte) (err error) {
	varGetUmFuturesTradeDownloadLinkByIdResponse := _GetUmFuturesTradeDownloadLinkByIdResponse{}

	err = json.Unmarshal(data, &varGetUmFuturesTradeDownloadLinkByIdResponse)

	if err != nil {
		return err
	}

	*o = GetUmFuturesTradeDownloadLinkByIdResponse(varGetUmFuturesTradeDownloadLinkByIdResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "downloadId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "url")
		delete(additionalProperties, "s3Link")
		delete(additionalProperties, "notified")
		delete(additionalProperties, "expirationTimestamp")
		delete(additionalProperties, "isExpired")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUmFuturesTradeDownloadLinkByIdResponse struct {
	value *GetUmFuturesTradeDownloadLinkByIdResponse
	isSet bool
}

func (v NullableGetUmFuturesTradeDownloadLinkByIdResponse) Get() *GetUmFuturesTradeDownloadLinkByIdResponse {
	return v.value
}

func (v *NullableGetUmFuturesTradeDownloadLinkByIdResponse) Set(val *GetUmFuturesTradeDownloadLinkByIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUmFuturesTradeDownloadLinkByIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUmFuturesTradeDownloadLinkByIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUmFuturesTradeDownloadLinkByIdResponse(val *GetUmFuturesTradeDownloadLinkByIdResponse) *NullableGetUmFuturesTradeDownloadLinkByIdResponse {
	return &NullableGetUmFuturesTradeDownloadLinkByIdResponse{value: val, isSet: true}
}

func (v NullableGetUmFuturesTradeDownloadLinkByIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUmFuturesTradeDownloadLinkByIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
