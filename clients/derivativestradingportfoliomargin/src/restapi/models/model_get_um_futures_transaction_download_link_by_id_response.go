/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetUmFuturesTransactionDownloadLinkByIdResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUmFuturesTransactionDownloadLinkByIdResponse{}

// GetUmFuturesTransactionDownloadLinkByIdResponse struct for GetUmFuturesTransactionDownloadLinkByIdResponse
type GetUmFuturesTransactionDownloadLinkByIdResponse struct {
	DownloadId           *string `json:"downloadId,omitempty"`
	Status               *string `json:"status,omitempty"`
	Url                  *string `json:"url,omitempty"`
	S3Link               *string `json:"s3Link,omitempty"`
	Notified             *bool   `json:"notified,omitempty"`
	ExpirationTimestamp  *int64  `json:"expirationTimestamp,omitempty"`
	IsExpired            *string `json:"isExpired,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUmFuturesTransactionDownloadLinkByIdResponse GetUmFuturesTransactionDownloadLinkByIdResponse

// NewGetUmFuturesTransactionDownloadLinkByIdResponse instantiates a new GetUmFuturesTransactionDownloadLinkByIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUmFuturesTransactionDownloadLinkByIdResponse() *GetUmFuturesTransactionDownloadLinkByIdResponse {
	this := GetUmFuturesTransactionDownloadLinkByIdResponse{}
	return &this
}

// NewGetUmFuturesTransactionDownloadLinkByIdResponseWithDefaults instantiates a new GetUmFuturesTransactionDownloadLinkByIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUmFuturesTransactionDownloadLinkByIdResponseWithDefaults() *GetUmFuturesTransactionDownloadLinkByIdResponse {
	this := GetUmFuturesTransactionDownloadLinkByIdResponse{}
	return &this
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetDownloadId() string {
	if o == nil || common.IsNil(o.DownloadId) {
		var ret string
		return ret
	}
	return *o.DownloadId
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetDownloadIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.DownloadId) {
		return nil, false
	}
	return o.DownloadId, true
}

// HasDownloadId returns a boolean if a field has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasDownloadId() bool {
	if o != nil && !common.IsNil(o.DownloadId) {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given string and assigns it to the DownloadId field.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetDownloadId(v string) {
	o.DownloadId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetUrl(v string) {
	o.Url = &v
}

// GetS3Link returns the S3Link field value if set, zero value otherwise.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetS3Link() string {
	if o == nil || common.IsNil(o.S3Link) {
		var ret string
		return ret
	}
	return *o.S3Link
}

// GetS3LinkOk returns a tuple with the S3Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetS3LinkOk() (*string, bool) {
	if o == nil || common.IsNil(o.S3Link) {
		return nil, false
	}
	return o.S3Link, true
}

// HasS3Link returns a boolean if a field has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasS3Link() bool {
	if o != nil && !common.IsNil(o.S3Link) {
		return true
	}

	return false
}

// SetS3Link gets a reference to the given string and assigns it to the S3Link field.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetS3Link(v string) {
	o.S3Link = &v
}

// GetNotified returns the Notified field value if set, zero value otherwise.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetNotified() bool {
	if o == nil || common.IsNil(o.Notified) {
		var ret bool
		return ret
	}
	return *o.Notified
}

// GetNotifiedOk returns a tuple with the Notified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetNotifiedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Notified) {
		return nil, false
	}
	return o.Notified, true
}

// HasNotified returns a boolean if a field has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasNotified() bool {
	if o != nil && !common.IsNil(o.Notified) {
		return true
	}

	return false
}

// SetNotified gets a reference to the given bool and assigns it to the Notified field.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetNotified(v bool) {
	o.Notified = &v
}

// GetExpirationTimestamp returns the ExpirationTimestamp field value if set, zero value otherwise.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetExpirationTimestamp() int64 {
	if o == nil || common.IsNil(o.ExpirationTimestamp) {
		var ret int64
		return ret
	}
	return *o.ExpirationTimestamp
}

// GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetExpirationTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpirationTimestamp) {
		return nil, false
	}
	return o.ExpirationTimestamp, true
}

// HasExpirationTimestamp returns a boolean if a field has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasExpirationTimestamp() bool {
	if o != nil && !common.IsNil(o.ExpirationTimestamp) {
		return true
	}

	return false
}

// SetExpirationTimestamp gets a reference to the given int64 and assigns it to the ExpirationTimestamp field.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetExpirationTimestamp(v int64) {
	o.ExpirationTimestamp = &v
}

// GetIsExpired returns the IsExpired field value if set, zero value otherwise.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetIsExpired() string {
	if o == nil || common.IsNil(o.IsExpired) {
		var ret string
		return ret
	}
	return *o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetIsExpiredOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsExpired) {
		return nil, false
	}
	return o.IsExpired, true
}

// HasIsExpired returns a boolean if a field has been set.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasIsExpired() bool {
	if o != nil && !common.IsNil(o.IsExpired) {
		return true
	}

	return false
}

// SetIsExpired gets a reference to the given string and assigns it to the IsExpired field.
func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetIsExpired(v string) {
	o.IsExpired = &v
}

func (o GetUmFuturesTransactionDownloadLinkByIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUmFuturesTransactionDownloadLinkByIdResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) UnmarshalJSON(data []byte) (err error) {
	varGetUmFuturesTransactionDownloadLinkByIdResponse := _GetUmFuturesTransactionDownloadLinkByIdResponse{}

	err = json.Unmarshal(data, &varGetUmFuturesTransactionDownloadLinkByIdResponse)

	if err != nil {
		return err
	}

	*o = GetUmFuturesTransactionDownloadLinkByIdResponse(varGetUmFuturesTransactionDownloadLinkByIdResponse)

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

type NullableGetUmFuturesTransactionDownloadLinkByIdResponse struct {
	value *GetUmFuturesTransactionDownloadLinkByIdResponse
	isSet bool
}

func (v NullableGetUmFuturesTransactionDownloadLinkByIdResponse) Get() *GetUmFuturesTransactionDownloadLinkByIdResponse {
	return v.value
}

func (v *NullableGetUmFuturesTransactionDownloadLinkByIdResponse) Set(val *GetUmFuturesTransactionDownloadLinkByIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUmFuturesTransactionDownloadLinkByIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUmFuturesTransactionDownloadLinkByIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUmFuturesTransactionDownloadLinkByIdResponse(val *GetUmFuturesTransactionDownloadLinkByIdResponse) *NullableGetUmFuturesTransactionDownloadLinkByIdResponse {
	return &NullableGetUmFuturesTransactionDownloadLinkByIdResponse{value: val, isSet: true}
}

func (v NullableGetUmFuturesTransactionDownloadLinkByIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUmFuturesTransactionDownloadLinkByIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
