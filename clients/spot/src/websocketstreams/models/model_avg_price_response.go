/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AvgPriceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AvgPriceResponse{}

// AvgPriceResponse struct for AvgPriceResponse
type AvgPriceResponse struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	S                    *string `json:"s,omitempty"`
	I                    *string `json:"i,omitempty"`
	W                    *string `json:"w,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AvgPriceResponse AvgPriceResponse

// NewAvgPriceResponse instantiates a new AvgPriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvgPriceResponse() *AvgPriceResponse {
	this := AvgPriceResponse{}
	return &this
}

// NewAvgPriceResponseWithDefaults instantiates a new AvgPriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvgPriceResponseWithDefaults() *AvgPriceResponse {
	this := AvgPriceResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AvgPriceResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AvgPriceResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AvgPriceResponse) SetS(v string) {
	o.S = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetI() string {
	if o == nil || common.IsNil(o.I) {
		var ret string
		return ret
	}
	return *o.I
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetIOk() (*string, bool) {
	if o == nil || common.IsNil(o.I) {
		return nil, false
	}
	return o.I, true
}

// HasI returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasI() bool {
	if o != nil && !common.IsNil(o.I) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *AvgPriceResponse) SetI(v string) {
	o.I = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetW() string {
	if o == nil || common.IsNil(o.W) {
		var ret string
		return ret
	}
	return *o.W
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetWOk() (*string, bool) {
	if o == nil || common.IsNil(o.W) {
		return nil, false
	}
	return o.W, true
}

// HasW returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasW() bool {
	if o != nil && !common.IsNil(o.W) {
		return true
	}

	return false
}

// SetW gets a reference to the given string and assigns it to the W field.
func (o *AvgPriceResponse) SetW(v string) {
	o.W = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AvgPriceResponse) SetT(v int64) {
	o.T = &v
}

func (o AvgPriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvgPriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.S) {
		toSerialize["s"] = o.S
	}
	if !common.IsNil(o.I) {
		toSerialize["i"] = o.I
	}
	if !common.IsNil(o.W) {
		toSerialize["w"] = o.W
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AvgPriceResponse) UnmarshalJSON(data []byte) (err error) {
	varAvgPriceResponse := _AvgPriceResponse{}

	err = json.Unmarshal(data, &varAvgPriceResponse)

	if err != nil {
		return err
	}

	*o = AvgPriceResponse(varAvgPriceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "i")
		delete(additionalProperties, "w")
		delete(additionalProperties, "T")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAvgPriceResponse struct {
	value *AvgPriceResponse
	isSet bool
}

func (v NullableAvgPriceResponse) Get() *AvgPriceResponse {
	return v.value
}

func (v *NullableAvgPriceResponse) Set(val *AvgPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAvgPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAvgPriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvgPriceResponse(val *AvgPriceResponse) *NullableAvgPriceResponse {
	return &NullableAvgPriceResponse{value: val, isSet: true}
}

func (v NullableAvgPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvgPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
