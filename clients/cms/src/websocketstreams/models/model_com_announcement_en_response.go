/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 types of per-user business events in real time via the Binance SApi WebSocket channel, covering trading (market / limit), reward claim, fund transfer and market resolution, plus order book (bid/ask) quotes broadcast per topic. Business events are routed by exact `(topic, userId)` match; order book topics are broadcast to every subscriber. Delivery is at-most-once, ordering is not guaranteed, and offline messages are not cached — pull the latest state via the REST API after reconnecting and deduplicate on `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ComAnnouncementEnResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ComAnnouncementEnResponse{}

// ComAnnouncementEnResponse struct for ComAnnouncementEnResponse
type ComAnnouncementEnResponse struct {
	// Catalog ID
	CatalogId *int64 `json:"catalogId,omitempty"`
	// Catalog Name
	CatalogName *string `json:"catalogName,omitempty"`
	// Publish timestamp (milliseconds)
	PublishDate *int64 `json:"publishDate,omitempty"`
	// Announcement title
	Title *string `json:"title,omitempty"`
	// Announcement body
	Body *string `json:"body,omitempty"`
	// Disclaimer
	Disclaimer           *string `json:"disclaimer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComAnnouncementEnResponse ComAnnouncementEnResponse

// NewComAnnouncementEnResponse instantiates a new ComAnnouncementEnResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComAnnouncementEnResponse() *ComAnnouncementEnResponse {
	this := ComAnnouncementEnResponse{}
	return &this
}

// NewComAnnouncementEnResponseWithDefaults instantiates a new ComAnnouncementEnResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComAnnouncementEnResponseWithDefaults() *ComAnnouncementEnResponse {
	this := ComAnnouncementEnResponse{}
	return &this
}

// GetCatalogId returns the CatalogId field value if set, zero value otherwise.
func (o *ComAnnouncementEnResponse) GetCatalogId() int64 {
	if o == nil || common.IsNil(o.CatalogId) {
		var ret int64
		return ret
	}
	return *o.CatalogId
}

// GetCatalogIdOk returns a tuple with the CatalogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComAnnouncementEnResponse) GetCatalogIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CatalogId) {
		return nil, false
	}
	return o.CatalogId, true
}

// HasCatalogId returns a boolean if a field has been set.
func (o *ComAnnouncementEnResponse) HasCatalogId() bool {
	if o != nil && !common.IsNil(o.CatalogId) {
		return true
	}

	return false
}

// SetCatalogId gets a reference to the given int64 and assigns it to the CatalogId field.
func (o *ComAnnouncementEnResponse) SetCatalogId(v int64) {
	o.CatalogId = &v
}

// GetCatalogName returns the CatalogName field value if set, zero value otherwise.
func (o *ComAnnouncementEnResponse) GetCatalogName() string {
	if o == nil || common.IsNil(o.CatalogName) {
		var ret string
		return ret
	}
	return *o.CatalogName
}

// GetCatalogNameOk returns a tuple with the CatalogName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComAnnouncementEnResponse) GetCatalogNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CatalogName) {
		return nil, false
	}
	return o.CatalogName, true
}

// HasCatalogName returns a boolean if a field has been set.
func (o *ComAnnouncementEnResponse) HasCatalogName() bool {
	if o != nil && !common.IsNil(o.CatalogName) {
		return true
	}

	return false
}

// SetCatalogName gets a reference to the given string and assigns it to the CatalogName field.
func (o *ComAnnouncementEnResponse) SetCatalogName(v string) {
	o.CatalogName = &v
}

// GetPublishDate returns the PublishDate field value if set, zero value otherwise.
func (o *ComAnnouncementEnResponse) GetPublishDate() int64 {
	if o == nil || common.IsNil(o.PublishDate) {
		var ret int64
		return ret
	}
	return *o.PublishDate
}

// GetPublishDateOk returns a tuple with the PublishDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComAnnouncementEnResponse) GetPublishDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PublishDate) {
		return nil, false
	}
	return o.PublishDate, true
}

// HasPublishDate returns a boolean if a field has been set.
func (o *ComAnnouncementEnResponse) HasPublishDate() bool {
	if o != nil && !common.IsNil(o.PublishDate) {
		return true
	}

	return false
}

// SetPublishDate gets a reference to the given int64 and assigns it to the PublishDate field.
func (o *ComAnnouncementEnResponse) SetPublishDate(v int64) {
	o.PublishDate = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ComAnnouncementEnResponse) GetTitle() string {
	if o == nil || common.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComAnnouncementEnResponse) GetTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ComAnnouncementEnResponse) HasTitle() bool {
	if o != nil && !common.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ComAnnouncementEnResponse) SetTitle(v string) {
	o.Title = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ComAnnouncementEnResponse) GetBody() string {
	if o == nil || common.IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComAnnouncementEnResponse) GetBodyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ComAnnouncementEnResponse) HasBody() bool {
	if o != nil && !common.IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *ComAnnouncementEnResponse) SetBody(v string) {
	o.Body = &v
}

// GetDisclaimer returns the Disclaimer field value if set, zero value otherwise.
func (o *ComAnnouncementEnResponse) GetDisclaimer() string {
	if o == nil || common.IsNil(o.Disclaimer) {
		var ret string
		return ret
	}
	return *o.Disclaimer
}

// GetDisclaimerOk returns a tuple with the Disclaimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComAnnouncementEnResponse) GetDisclaimerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Disclaimer) {
		return nil, false
	}
	return o.Disclaimer, true
}

// HasDisclaimer returns a boolean if a field has been set.
func (o *ComAnnouncementEnResponse) HasDisclaimer() bool {
	if o != nil && !common.IsNil(o.Disclaimer) {
		return true
	}

	return false
}

// SetDisclaimer gets a reference to the given string and assigns it to the Disclaimer field.
func (o *ComAnnouncementEnResponse) SetDisclaimer(v string) {
	o.Disclaimer = &v
}

func (o ComAnnouncementEnResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComAnnouncementEnResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CatalogId) {
		toSerialize["catalogId"] = o.CatalogId
	}
	if !common.IsNil(o.CatalogName) {
		toSerialize["catalogName"] = o.CatalogName
	}
	if !common.IsNil(o.PublishDate) {
		toSerialize["publishDate"] = o.PublishDate
	}
	if !common.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !common.IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !common.IsNil(o.Disclaimer) {
		toSerialize["disclaimer"] = o.Disclaimer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ComAnnouncementEnResponse) UnmarshalJSON(data []byte) (err error) {
	varComAnnouncementEnResponse := _ComAnnouncementEnResponse{}

	err = json.Unmarshal(data, &varComAnnouncementEnResponse)

	if err != nil {
		return err
	}

	*o = ComAnnouncementEnResponse(varComAnnouncementEnResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "catalogId")
		delete(additionalProperties, "catalogName")
		delete(additionalProperties, "publishDate")
		delete(additionalProperties, "title")
		delete(additionalProperties, "body")
		delete(additionalProperties, "disclaimer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComAnnouncementEnResponse struct {
	value *ComAnnouncementEnResponse
	isSet bool
}

func (v NullableComAnnouncementEnResponse) Get() *ComAnnouncementEnResponse {
	return v.value
}

func (v *NullableComAnnouncementEnResponse) Set(val *ComAnnouncementEnResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableComAnnouncementEnResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableComAnnouncementEnResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComAnnouncementEnResponse(val *ComAnnouncementEnResponse) *NullableComAnnouncementEnResponse {
	return &NullableComAnnouncementEnResponse{value: val, isSet: true}
}

func (v NullableComAnnouncementEnResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComAnnouncementEnResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
