/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MiniTickerResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MiniTickerResponse{}

// MiniTickerResponse struct for MiniTickerResponse
type MiniTickerResponse struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	S                    *string `json:"s,omitempty"`
	C                    *string `json:"c,omitempty"`
	O                    *string `json:"o,omitempty"`
	H                    *string `json:"h,omitempty"`
	L                    *string `json:"l,omitempty"`
	V                    *string `json:"v,omitempty"`
	Q                    *string `json:"q,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MiniTickerResponse MiniTickerResponse

// NewMiniTickerResponse instantiates a new MiniTickerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiniTickerResponse() *MiniTickerResponse {
	this := MiniTickerResponse{}
	return &this
}

// NewMiniTickerResponseWithDefaults instantiates a new MiniTickerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiniTickerResponseWithDefaults() *MiniTickerResponse {
	this := MiniTickerResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *MiniTickerResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *MiniTickerResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *MiniTickerResponse) SetS(v string) {
	o.S = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetC() string {
	if o == nil || common.IsNil(o.C) {
		var ret string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetCOk() (*string, bool) {
	if o == nil || common.IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasC() bool {
	if o != nil && !common.IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *MiniTickerResponse) SetC(v string) {
	o.C = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetO() string {
	if o == nil || common.IsNil(o.O) {
		var ret string
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetOOk() (*string, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *MiniTickerResponse) SetO(v string) {
	o.O = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetH() string {
	if o == nil || common.IsNil(o.H) {
		var ret string
		return ret
	}
	return *o.H
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetHOk() (*string, bool) {
	if o == nil || common.IsNil(o.H) {
		return nil, false
	}
	return o.H, true
}

// HasH returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasH() bool {
	if o != nil && !common.IsNil(o.H) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *MiniTickerResponse) SetH(v string) {
	o.H = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetL() string {
	if o == nil || common.IsNil(o.L) {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetLOk() (*string, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *MiniTickerResponse) SetL(v string) {
	o.L = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *MiniTickerResponse) SetV(v string) {
	o.V = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *MiniTickerResponse) SetQ(v string) {
	o.Q = &v
}

func (o MiniTickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiniTickerResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.C) {
		toSerialize["c"] = o.C
	}
	if !common.IsNil(o.O) {
		toSerialize["o"] = o.O
	}
	if !common.IsNil(o.H) {
		toSerialize["h"] = o.H
	}
	if !common.IsNil(o.L) {
		toSerialize["l"] = o.L
	}
	if !common.IsNil(o.V) {
		toSerialize["v"] = o.V
	}
	if !common.IsNil(o.Q) {
		toSerialize["q"] = o.Q
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MiniTickerResponse) UnmarshalJSON(data []byte) (err error) {
	varMiniTickerResponse := _MiniTickerResponse{}

	err = json.Unmarshal(data, &varMiniTickerResponse)

	if err != nil {
		return err
	}

	*o = MiniTickerResponse(varMiniTickerResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "c")
		delete(additionalProperties, "o")
		delete(additionalProperties, "h")
		delete(additionalProperties, "l")
		delete(additionalProperties, "v")
		delete(additionalProperties, "q")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMiniTickerResponse struct {
	value *MiniTickerResponse
	isSet bool
}

func (v NullableMiniTickerResponse) Get() *MiniTickerResponse {
	return v.value
}

func (v *NullableMiniTickerResponse) Set(val *MiniTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMiniTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMiniTickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiniTickerResponse(val *MiniTickerResponse) *NullableMiniTickerResponse {
	return &NullableMiniTickerResponse{value: val, isSet: true}
}

func (v NullableMiniTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiniTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
