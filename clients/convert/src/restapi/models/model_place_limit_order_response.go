/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PlaceLimitOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlaceLimitOrderResponse{}

// PlaceLimitOrderResponse struct for PlaceLimitOrderResponse
type PlaceLimitOrderResponse struct {
	QuoteId              *string `json:"quoteId,omitempty"`
	Ratio                *string `json:"ratio,omitempty"`
	InverseRatio         *string `json:"inverseRatio,omitempty"`
	ValidTimestamp       *int64  `json:"validTimestamp,omitempty"`
	ToAmount             *string `json:"toAmount,omitempty"`
	FromAmount           *string `json:"fromAmount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlaceLimitOrderResponse PlaceLimitOrderResponse

// NewPlaceLimitOrderResponse instantiates a new PlaceLimitOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceLimitOrderResponse() *PlaceLimitOrderResponse {
	this := PlaceLimitOrderResponse{}
	return &this
}

// NewPlaceLimitOrderResponseWithDefaults instantiates a new PlaceLimitOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceLimitOrderResponseWithDefaults() *PlaceLimitOrderResponse {
	this := PlaceLimitOrderResponse{}
	return &this
}

// GetQuoteId returns the QuoteId field value if set, zero value otherwise.
func (o *PlaceLimitOrderResponse) GetQuoteId() string {
	if o == nil || common.IsNil(o.QuoteId) {
		var ret string
		return ret
	}
	return *o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceLimitOrderResponse) GetQuoteIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteId) {
		return nil, false
	}
	return o.QuoteId, true
}

// HasQuoteId returns a boolean if a field has been set.
func (o *PlaceLimitOrderResponse) HasQuoteId() bool {
	if o != nil && !common.IsNil(o.QuoteId) {
		return true
	}

	return false
}

// SetQuoteId gets a reference to the given string and assigns it to the QuoteId field.
func (o *PlaceLimitOrderResponse) SetQuoteId(v string) {
	o.QuoteId = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *PlaceLimitOrderResponse) GetRatio() string {
	if o == nil || common.IsNil(o.Ratio) {
		var ret string
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceLimitOrderResponse) GetRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *PlaceLimitOrderResponse) HasRatio() bool {
	if o != nil && !common.IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given string and assigns it to the Ratio field.
func (o *PlaceLimitOrderResponse) SetRatio(v string) {
	o.Ratio = &v
}

// GetInverseRatio returns the InverseRatio field value if set, zero value otherwise.
func (o *PlaceLimitOrderResponse) GetInverseRatio() string {
	if o == nil || common.IsNil(o.InverseRatio) {
		var ret string
		return ret
	}
	return *o.InverseRatio
}

// GetInverseRatioOk returns a tuple with the InverseRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceLimitOrderResponse) GetInverseRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.InverseRatio) {
		return nil, false
	}
	return o.InverseRatio, true
}

// HasInverseRatio returns a boolean if a field has been set.
func (o *PlaceLimitOrderResponse) HasInverseRatio() bool {
	if o != nil && !common.IsNil(o.InverseRatio) {
		return true
	}

	return false
}

// SetInverseRatio gets a reference to the given string and assigns it to the InverseRatio field.
func (o *PlaceLimitOrderResponse) SetInverseRatio(v string) {
	o.InverseRatio = &v
}

// GetValidTimestamp returns the ValidTimestamp field value if set, zero value otherwise.
func (o *PlaceLimitOrderResponse) GetValidTimestamp() int64 {
	if o == nil || common.IsNil(o.ValidTimestamp) {
		var ret int64
		return ret
	}
	return *o.ValidTimestamp
}

// GetValidTimestampOk returns a tuple with the ValidTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceLimitOrderResponse) GetValidTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ValidTimestamp) {
		return nil, false
	}
	return o.ValidTimestamp, true
}

// HasValidTimestamp returns a boolean if a field has been set.
func (o *PlaceLimitOrderResponse) HasValidTimestamp() bool {
	if o != nil && !common.IsNil(o.ValidTimestamp) {
		return true
	}

	return false
}

// SetValidTimestamp gets a reference to the given int64 and assigns it to the ValidTimestamp field.
func (o *PlaceLimitOrderResponse) SetValidTimestamp(v int64) {
	o.ValidTimestamp = &v
}

// GetToAmount returns the ToAmount field value if set, zero value otherwise.
func (o *PlaceLimitOrderResponse) GetToAmount() string {
	if o == nil || common.IsNil(o.ToAmount) {
		var ret string
		return ret
	}
	return *o.ToAmount
}

// GetToAmountOk returns a tuple with the ToAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceLimitOrderResponse) GetToAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAmount) {
		return nil, false
	}
	return o.ToAmount, true
}

// HasToAmount returns a boolean if a field has been set.
func (o *PlaceLimitOrderResponse) HasToAmount() bool {
	if o != nil && !common.IsNil(o.ToAmount) {
		return true
	}

	return false
}

// SetToAmount gets a reference to the given string and assigns it to the ToAmount field.
func (o *PlaceLimitOrderResponse) SetToAmount(v string) {
	o.ToAmount = &v
}

// GetFromAmount returns the FromAmount field value if set, zero value otherwise.
func (o *PlaceLimitOrderResponse) GetFromAmount() string {
	if o == nil || common.IsNil(o.FromAmount) {
		var ret string
		return ret
	}
	return *o.FromAmount
}

// GetFromAmountOk returns a tuple with the FromAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceLimitOrderResponse) GetFromAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAmount) {
		return nil, false
	}
	return o.FromAmount, true
}

// HasFromAmount returns a boolean if a field has been set.
func (o *PlaceLimitOrderResponse) HasFromAmount() bool {
	if o != nil && !common.IsNil(o.FromAmount) {
		return true
	}

	return false
}

// SetFromAmount gets a reference to the given string and assigns it to the FromAmount field.
func (o *PlaceLimitOrderResponse) SetFromAmount(v string) {
	o.FromAmount = &v
}

func (o PlaceLimitOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaceLimitOrderResponse) ToMap() (map[string]interface{}, error) {
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

func (o *PlaceLimitOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varPlaceLimitOrderResponse := _PlaceLimitOrderResponse{}

	err = json.Unmarshal(data, &varPlaceLimitOrderResponse)

	if err != nil {
		return err
	}

	*o = PlaceLimitOrderResponse(varPlaceLimitOrderResponse)

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

type NullablePlaceLimitOrderResponse struct {
	value *PlaceLimitOrderResponse
	isSet bool
}

func (v NullablePlaceLimitOrderResponse) Get() *PlaceLimitOrderResponse {
	return v.value
}

func (v *NullablePlaceLimitOrderResponse) Set(val *PlaceLimitOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceLimitOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceLimitOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceLimitOrderResponse(val *PlaceLimitOrderResponse) *NullablePlaceLimitOrderResponse {
	return &NullablePlaceLimitOrderResponse{value: val, isSet: true}
}

func (v NullablePlaceLimitOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceLimitOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
