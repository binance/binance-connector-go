/*
Stocks Trading WebSocket Streams

WebSocket stream definitions for Binance Stocks Trading. Base URL: wss://nbstream.binance.com/equity
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuoteStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuoteStreamResponse{}

// QuoteStreamResponse struct for QuoteStreamResponse
type QuoteStreamResponse struct {
	// Event type, always `\"quote\"`.
	Smalle *string `json:"e,omitempty"`
	// Event time — epoch milliseconds when the server pushed the message.
	E *int64 `json:"E,omitempty"`
	// Symbol (UPPERCASE ticker), e.g. `\"AAPL\"`.
	Smalls *string `json:"s,omitempty"`
	// Best bid price.
	Smallbp *string `json:"bp,omitempty"`
	// Best ask price.
	Smallap *string `json:"ap,omitempty"`
	// Best bid size (shares).
	Smallbs *int32 `json:"bs,omitempty"`
	// Best ask size (shares).
	Smallas *int32 `json:"as,omitempty"`
	// Source quote timestamp (epoch milliseconds); may be null.
	T                    *int64 `json:"T,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuoteStreamResponse QuoteStreamResponse

// NewQuoteStreamResponse instantiates a new QuoteStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteStreamResponse() *QuoteStreamResponse {
	this := QuoteStreamResponse{}
	return &this
}

// NewQuoteStreamResponseWithDefaults instantiates a new QuoteStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteStreamResponseWithDefaults() *QuoteStreamResponse {
	this := QuoteStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *QuoteStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *QuoteStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *QuoteStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *QuoteStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *QuoteStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *QuoteStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *QuoteStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *QuoteStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *QuoteStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetBp returns the Bp field value if set, zero value otherwise.
func (o *QuoteStreamResponse) GetSmallbp() string {
	if o == nil || common.IsNil(o.Smallbp) {
		var ret string
		return ret
	}
	return *o.Smallbp
}

// GetBpOk returns a tuple with the Bp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteStreamResponse) GetSmallbpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallbp) {
		return nil, false
	}
	return o.Smallbp, true
}

// HasBp returns a boolean if a field has been set.
func (o *QuoteStreamResponse) HasSmallbp() bool {
	if o != nil && !common.IsNil(o.Smallbp) {
		return true
	}

	return false
}

// SetBp gets a reference to the given string and assigns it to the Bp field.
func (o *QuoteStreamResponse) SetSmallbp(v string) {
	o.Smallbp = &v
}

// GetAp returns the Ap field value if set, zero value otherwise.
func (o *QuoteStreamResponse) GetSmallap() string {
	if o == nil || common.IsNil(o.Smallap) {
		var ret string
		return ret
	}
	return *o.Smallap
}

// GetApOk returns a tuple with the Ap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteStreamResponse) GetSmallapOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallap) {
		return nil, false
	}
	return o.Smallap, true
}

// HasAp returns a boolean if a field has been set.
func (o *QuoteStreamResponse) HasSmallap() bool {
	if o != nil && !common.IsNil(o.Smallap) {
		return true
	}

	return false
}

// SetAp gets a reference to the given string and assigns it to the Ap field.
func (o *QuoteStreamResponse) SetSmallap(v string) {
	o.Smallap = &v
}

// GetBs returns the Bs field value if set, zero value otherwise.
func (o *QuoteStreamResponse) GetSmallbs() int32 {
	if o == nil || common.IsNil(o.Smallbs) {
		var ret int32
		return ret
	}
	return *o.Smallbs
}

// GetBsOk returns a tuple with the Bs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteStreamResponse) GetSmallbsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Smallbs) {
		return nil, false
	}
	return o.Smallbs, true
}

// HasBs returns a boolean if a field has been set.
func (o *QuoteStreamResponse) HasSmallbs() bool {
	if o != nil && !common.IsNil(o.Smallbs) {
		return true
	}

	return false
}

// SetBs gets a reference to the given int32 and assigns it to the Bs field.
func (o *QuoteStreamResponse) SetSmallbs(v int32) {
	o.Smallbs = &v
}

// GetAs returns the As field value if set, zero value otherwise.
func (o *QuoteStreamResponse) GetSmallas() int32 {
	if o == nil || common.IsNil(o.Smallas) {
		var ret int32
		return ret
	}
	return *o.Smallas
}

// GetAsOk returns a tuple with the As field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteStreamResponse) GetSmallasOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Smallas) {
		return nil, false
	}
	return o.Smallas, true
}

// HasAs returns a boolean if a field has been set.
func (o *QuoteStreamResponse) HasSmallas() bool {
	if o != nil && !common.IsNil(o.Smallas) {
		return true
	}

	return false
}

// SetAs gets a reference to the given int32 and assigns it to the As field.
func (o *QuoteStreamResponse) SetSmallas(v int32) {
	o.Smallas = &v
}

// GetT returns the T field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuoteStreamResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuoteStreamResponse) GetTOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *QuoteStreamResponse) HasT() bool {
	if o != nil && common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given NullableInt64 and assigns it to the T field.
func (o *QuoteStreamResponse) SetT(v int64) {
	o.T = &v
}

func (o QuoteStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallbp) {
		toSerialize["bp"] = o.Smallbp
	}
	if !common.IsNil(o.Smallap) {
		toSerialize["ap"] = o.Smallap
	}
	if !common.IsNil(o.Smallbs) {
		toSerialize["bs"] = o.Smallbs
	}
	if !common.IsNil(o.Smallas) {
		toSerialize["as"] = o.Smallas
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuoteStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varQuoteStreamResponse := _QuoteStreamResponse{}

	err = json.Unmarshal(data, &varQuoteStreamResponse)

	if err != nil {
		return err
	}

	*o = QuoteStreamResponse(varQuoteStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "bp")
		delete(additionalProperties, "ap")
		delete(additionalProperties, "bs")
		delete(additionalProperties, "as")
		delete(additionalProperties, "T")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuoteStreamResponse struct {
	value *QuoteStreamResponse
	isSet bool
}

func (v NullableQuoteStreamResponse) Get() *QuoteStreamResponse {
	return v.value
}

func (v *NullableQuoteStreamResponse) Set(val *QuoteStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteStreamResponse(val *QuoteStreamResponse) *NullableQuoteStreamResponse {
	return &NullableQuoteStreamResponse{value: val, isSet: true}
}

func (v NullableQuoteStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
