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

// checks if the SendQuoteRequestResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SendQuoteRequestResponse{}

// SendQuoteRequestResponse struct for SendQuoteRequestResponse
type SendQuoteRequestResponse struct {
	QuoteId              *string `json:"quoteId,omitempty"`
	Ratio                *string `json:"ratio,omitempty"`
	InverseRatio         *string `json:"inverseRatio,omitempty"`
	ValidTimestamp       *int64  `json:"validTimestamp,omitempty"`
	ToAmount             *string `json:"toAmount,omitempty"`
	FromAmount           *string `json:"fromAmount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SendQuoteRequestResponse SendQuoteRequestResponse

// NewSendQuoteRequestResponse instantiates a new SendQuoteRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendQuoteRequestResponse() *SendQuoteRequestResponse {
	this := SendQuoteRequestResponse{}
	return &this
}

// NewSendQuoteRequestResponseWithDefaults instantiates a new SendQuoteRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendQuoteRequestResponseWithDefaults() *SendQuoteRequestResponse {
	this := SendQuoteRequestResponse{}
	return &this
}

// GetQuoteId returns the QuoteId field value if set, zero value otherwise.
func (o *SendQuoteRequestResponse) GetQuoteId() string {
	if o == nil || common.IsNil(o.QuoteId) {
		var ret string
		return ret
	}
	return *o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendQuoteRequestResponse) GetQuoteIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteId) {
		return nil, false
	}
	return o.QuoteId, true
}

// HasQuoteId returns a boolean if a field has been set.
func (o *SendQuoteRequestResponse) HasQuoteId() bool {
	if o != nil && !common.IsNil(o.QuoteId) {
		return true
	}

	return false
}

// SetQuoteId gets a reference to the given string and assigns it to the QuoteId field.
func (o *SendQuoteRequestResponse) SetQuoteId(v string) {
	o.QuoteId = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *SendQuoteRequestResponse) GetRatio() string {
	if o == nil || common.IsNil(o.Ratio) {
		var ret string
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendQuoteRequestResponse) GetRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *SendQuoteRequestResponse) HasRatio() bool {
	if o != nil && !common.IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given string and assigns it to the Ratio field.
func (o *SendQuoteRequestResponse) SetRatio(v string) {
	o.Ratio = &v
}

// GetInverseRatio returns the InverseRatio field value if set, zero value otherwise.
func (o *SendQuoteRequestResponse) GetInverseRatio() string {
	if o == nil || common.IsNil(o.InverseRatio) {
		var ret string
		return ret
	}
	return *o.InverseRatio
}

// GetInverseRatioOk returns a tuple with the InverseRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendQuoteRequestResponse) GetInverseRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.InverseRatio) {
		return nil, false
	}
	return o.InverseRatio, true
}

// HasInverseRatio returns a boolean if a field has been set.
func (o *SendQuoteRequestResponse) HasInverseRatio() bool {
	if o != nil && !common.IsNil(o.InverseRatio) {
		return true
	}

	return false
}

// SetInverseRatio gets a reference to the given string and assigns it to the InverseRatio field.
func (o *SendQuoteRequestResponse) SetInverseRatio(v string) {
	o.InverseRatio = &v
}

// GetValidTimestamp returns the ValidTimestamp field value if set, zero value otherwise.
func (o *SendQuoteRequestResponse) GetValidTimestamp() int64 {
	if o == nil || common.IsNil(o.ValidTimestamp) {
		var ret int64
		return ret
	}
	return *o.ValidTimestamp
}

// GetValidTimestampOk returns a tuple with the ValidTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendQuoteRequestResponse) GetValidTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ValidTimestamp) {
		return nil, false
	}
	return o.ValidTimestamp, true
}

// HasValidTimestamp returns a boolean if a field has been set.
func (o *SendQuoteRequestResponse) HasValidTimestamp() bool {
	if o != nil && !common.IsNil(o.ValidTimestamp) {
		return true
	}

	return false
}

// SetValidTimestamp gets a reference to the given int64 and assigns it to the ValidTimestamp field.
func (o *SendQuoteRequestResponse) SetValidTimestamp(v int64) {
	o.ValidTimestamp = &v
}

// GetToAmount returns the ToAmount field value if set, zero value otherwise.
func (o *SendQuoteRequestResponse) GetToAmount() string {
	if o == nil || common.IsNil(o.ToAmount) {
		var ret string
		return ret
	}
	return *o.ToAmount
}

// GetToAmountOk returns a tuple with the ToAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendQuoteRequestResponse) GetToAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAmount) {
		return nil, false
	}
	return o.ToAmount, true
}

// HasToAmount returns a boolean if a field has been set.
func (o *SendQuoteRequestResponse) HasToAmount() bool {
	if o != nil && !common.IsNil(o.ToAmount) {
		return true
	}

	return false
}

// SetToAmount gets a reference to the given string and assigns it to the ToAmount field.
func (o *SendQuoteRequestResponse) SetToAmount(v string) {
	o.ToAmount = &v
}

// GetFromAmount returns the FromAmount field value if set, zero value otherwise.
func (o *SendQuoteRequestResponse) GetFromAmount() string {
	if o == nil || common.IsNil(o.FromAmount) {
		var ret string
		return ret
	}
	return *o.FromAmount
}

// GetFromAmountOk returns a tuple with the FromAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendQuoteRequestResponse) GetFromAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAmount) {
		return nil, false
	}
	return o.FromAmount, true
}

// HasFromAmount returns a boolean if a field has been set.
func (o *SendQuoteRequestResponse) HasFromAmount() bool {
	if o != nil && !common.IsNil(o.FromAmount) {
		return true
	}

	return false
}

// SetFromAmount gets a reference to the given string and assigns it to the FromAmount field.
func (o *SendQuoteRequestResponse) SetFromAmount(v string) {
	o.FromAmount = &v
}

func (o SendQuoteRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendQuoteRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.QuoteId) {
		toSerialize["quoteId"] = o.QuoteId
	}
	if !common.IsNil(o.Ratio) {
		toSerialize["ratio"] = o.Ratio
	}
	if !common.IsNil(o.InverseRatio) {
		toSerialize["inverseRatio"] = o.InverseRatio
	}
	if !common.IsNil(o.ValidTimestamp) {
		toSerialize["validTimestamp"] = o.ValidTimestamp
	}
	if !common.IsNil(o.ToAmount) {
		toSerialize["toAmount"] = o.ToAmount
	}
	if !common.IsNil(o.FromAmount) {
		toSerialize["fromAmount"] = o.FromAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SendQuoteRequestResponse) UnmarshalJSON(data []byte) (err error) {
	varSendQuoteRequestResponse := _SendQuoteRequestResponse{}

	err = json.Unmarshal(data, &varSendQuoteRequestResponse)

	if err != nil {
		return err
	}

	*o = SendQuoteRequestResponse(varSendQuoteRequestResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "quoteId")
		delete(additionalProperties, "ratio")
		delete(additionalProperties, "inverseRatio")
		delete(additionalProperties, "validTimestamp")
		delete(additionalProperties, "toAmount")
		delete(additionalProperties, "fromAmount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSendQuoteRequestResponse struct {
	value *SendQuoteRequestResponse
	isSet bool
}

func (v NullableSendQuoteRequestResponse) Get() *SendQuoteRequestResponse {
	return v.value
}

func (v *NullableSendQuoteRequestResponse) Set(val *SendQuoteRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSendQuoteRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSendQuoteRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendQuoteRequestResponse(val *SendQuoteRequestResponse) *NullableSendQuoteRequestResponse {
	return &NullableSendQuoteRequestResponse{value: val, isSet: true}
}

func (v NullableSendQuoteRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendQuoteRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
