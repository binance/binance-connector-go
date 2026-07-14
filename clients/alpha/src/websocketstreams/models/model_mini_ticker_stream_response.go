/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MiniTickerStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MiniTickerStreamResponse{}

// MiniTickerStreamResponse struct for MiniTickerStreamResponse
type MiniTickerStreamResponse struct {
	// eventType
	Smalle *string `json:"e,omitempty"`
	// eventTime
	E *int64 `json:"E,omitempty"`
	// symbol
	Smalls *string `json:"s,omitempty"`
	// closePrice
	Smallc *string `json:"c,omitempty"`
	// openPrice
	Smallo *string `json:"o,omitempty"`
	// highPrice
	Smallh *string `json:"h,omitempty"`
	// lowPrice
	Smalll *string `json:"l,omitempty"`
	// volume
	Smallv *string `json:"v,omitempty"`
	// quoteVolume
	Smallq               *string `json:"q,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MiniTickerStreamResponse MiniTickerStreamResponse

// NewMiniTickerStreamResponse instantiates a new MiniTickerStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiniTickerStreamResponse() *MiniTickerStreamResponse {
	this := MiniTickerStreamResponse{}
	return &this
}

// NewMiniTickerStreamResponseWithDefaults instantiates a new MiniTickerStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiniTickerStreamResponseWithDefaults() *MiniTickerStreamResponse {
	this := MiniTickerStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MiniTickerStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *MiniTickerStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *MiniTickerStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MiniTickerStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *MiniTickerStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *MiniTickerStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *MiniTickerStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *MiniTickerStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *MiniTickerStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *MiniTickerStreamResponse) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerStreamResponse) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *MiniTickerStreamResponse) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *MiniTickerStreamResponse) SetSmallc(v string) {
	o.Smallc = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *MiniTickerStreamResponse) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerStreamResponse) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *MiniTickerStreamResponse) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *MiniTickerStreamResponse) SetSmallo(v string) {
	o.Smallo = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *MiniTickerStreamResponse) GetSmallh() string {
	if o == nil || common.IsNil(o.Smallh) {
		var ret string
		return ret
	}
	return *o.Smallh
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerStreamResponse) GetSmallhOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallh) {
		return nil, false
	}
	return o.Smallh, true
}

// HasH returns a boolean if a field has been set.
func (o *MiniTickerStreamResponse) HasSmallh() bool {
	if o != nil && !common.IsNil(o.Smallh) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *MiniTickerStreamResponse) SetSmallh(v string) {
	o.Smallh = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *MiniTickerStreamResponse) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerStreamResponse) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *MiniTickerStreamResponse) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *MiniTickerStreamResponse) SetSmalll(v string) {
	o.Smalll = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *MiniTickerStreamResponse) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerStreamResponse) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *MiniTickerStreamResponse) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *MiniTickerStreamResponse) SetSmallv(v string) {
	o.Smallv = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *MiniTickerStreamResponse) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniTickerStreamResponse) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *MiniTickerStreamResponse) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *MiniTickerStreamResponse) SetSmallq(v string) {
	o.Smallq = &v
}

func (o MiniTickerStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiniTickerStreamResponse) ToMap() (map[string]interface{}, error) {
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

func (o *MiniTickerStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varMiniTickerStreamResponse := _MiniTickerStreamResponse{}

	err = json.Unmarshal(data, &varMiniTickerStreamResponse)

	if err != nil {
		return err
	}

	*o = MiniTickerStreamResponse(varMiniTickerStreamResponse)

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

type NullableMiniTickerStreamResponse struct {
	value *MiniTickerStreamResponse
	isSet bool
}

func (v NullableMiniTickerStreamResponse) Get() *MiniTickerStreamResponse {
	return v.value
}

func (v *NullableMiniTickerStreamResponse) Set(val *MiniTickerStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMiniTickerStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMiniTickerStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiniTickerStreamResponse(val *MiniTickerStreamResponse) *NullableMiniTickerStreamResponse {
	return &NullableMiniTickerStreamResponse{value: val, isSet: true}
}

func (v NullableMiniTickerStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiniTickerStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
