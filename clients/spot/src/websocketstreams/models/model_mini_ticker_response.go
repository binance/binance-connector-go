/*
Spot WebSocket Market Streams

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MiniTickerResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MiniTickerResponse{}

// MiniTickerResponse struct for MiniTickerResponse
type MiniTickerResponse struct {
	// Event type
	Smalle *string `json:"e,omitempty"`
	// Event time
	E *int64 `json:"E,omitempty"`
	// Symbol
	Smalls *string `json:"s,omitempty"`
	// Close price
	Smallc *string `json:"c,omitempty"`
	// Open price
	Smallo *string `json:"o,omitempty"`
	// High price
	Smallh *string `json:"h,omitempty"`
	// Low price
	Smalll *string `json:"l,omitempty"`
	// Total traded base asset volume
	Smallv *string `json:"v,omitempty"`
	// Total traded quote asset volume
	Smallq               *string `json:"q,omitempty"`
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
func (o *MiniTickerResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *MiniTickerResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *MiniTickerResponse) SetSmallc(v string) {
	o.Smallc = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *MiniTickerResponse) SetSmallo(v string) {
	o.Smallo = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetSmallh() string {
	if o == nil || common.IsNil(o.Smallh) {
		var ret string
		return ret
	}
	return *o.Smallh
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetSmallhOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallh) {
		return nil, false
	}
	return o.Smallh, true
}

// HasH returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasSmallh() bool {
	if o != nil && !common.IsNil(o.Smallh) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *MiniTickerResponse) SetSmallh(v string) {
	o.Smallh = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *MiniTickerResponse) SetSmalll(v string) {
	o.Smalll = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *MiniTickerResponse) SetSmallv(v string) {
	o.Smallv = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *MiniTickerResponse) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerResponse) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *MiniTickerResponse) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *MiniTickerResponse) SetSmallq(v string) {
	o.Smallq = &v
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
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}
	if !common.IsNil(o.Smallh) {
		toSerialize["h"] = o.Smallh
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.Smallv) {
		toSerialize["v"] = o.Smallv
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
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
