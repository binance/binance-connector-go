/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AggregateTradeStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AggregateTradeStreamsResponse{}

// AggregateTradeStreamsResponse struct for AggregateTradeStreamsResponse
type AggregateTradeStreamsResponse struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smalla               *int64  `json:"a,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	Smallf               *int64  `json:"f,omitempty"`
	Smalll               *int64  `json:"l,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	Smallm               *bool   `json:"m,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AggregateTradeStreamsResponse AggregateTradeStreamsResponse

// NewAggregateTradeStreamsResponse instantiates a new AggregateTradeStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregateTradeStreamsResponse() *AggregateTradeStreamsResponse {
	this := AggregateTradeStreamsResponse{}
	return &this
}

// NewAggregateTradeStreamsResponseWithDefaults instantiates a new AggregateTradeStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregateTradeStreamsResponseWithDefaults() *AggregateTradeStreamsResponse {
	this := AggregateTradeStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AggregateTradeStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *AggregateTradeStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AggregateTradeStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AggregateTradeStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AggregateTradeStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AggregateTradeStreamsResponse) SetE(v int64) {
	o.E = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *AggregateTradeStreamsResponse) GetSmalla() int64 {
	if o == nil || common.IsNil(o.Smalla) {
		var ret int64
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamsResponse) GetSmallaOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *AggregateTradeStreamsResponse) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given int64 and assigns it to the A field.
func (o *AggregateTradeStreamsResponse) SetSmalla(v int64) {
	o.Smalla = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AggregateTradeStreamsResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamsResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AggregateTradeStreamsResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AggregateTradeStreamsResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AggregateTradeStreamsResponse) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamsResponse) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *AggregateTradeStreamsResponse) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AggregateTradeStreamsResponse) SetSmallp(v string) {
	o.Smallp = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AggregateTradeStreamsResponse) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamsResponse) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *AggregateTradeStreamsResponse) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AggregateTradeStreamsResponse) SetSmallq(v string) {
	o.Smallq = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *AggregateTradeStreamsResponse) GetSmallf() int64 {
	if o == nil || common.IsNil(o.Smallf) {
		var ret int64
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamsResponse) GetSmallfOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *AggregateTradeStreamsResponse) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given int64 and assigns it to the F field.
func (o *AggregateTradeStreamsResponse) SetSmallf(v int64) {
	o.Smallf = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AggregateTradeStreamsResponse) GetSmalll() int64 {
	if o == nil || common.IsNil(o.Smalll) {
		var ret int64
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamsResponse) GetSmalllOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *AggregateTradeStreamsResponse) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *AggregateTradeStreamsResponse) SetSmalll(v int64) {
	o.Smalll = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AggregateTradeStreamsResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamsResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AggregateTradeStreamsResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AggregateTradeStreamsResponse) SetT(v int64) {
	o.T = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *AggregateTradeStreamsResponse) GetSmallm() bool {
	if o == nil || common.IsNil(o.Smallm) {
		var ret bool
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamsResponse) GetSmallmOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *AggregateTradeStreamsResponse) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *AggregateTradeStreamsResponse) SetSmallm(v bool) {
	o.Smallm = &v
}

func (o AggregateTradeStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregateTradeStreamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.Smallf) {
		toSerialize["f"] = o.Smallf
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AggregateTradeStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varAggregateTradeStreamsResponse := _AggregateTradeStreamsResponse{}

	err = json.Unmarshal(data, &varAggregateTradeStreamsResponse)

	if err != nil {
		return err
	}

	*o = AggregateTradeStreamsResponse(varAggregateTradeStreamsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "a")
		delete(additionalProperties, "s")
		delete(additionalProperties, "p")
		delete(additionalProperties, "q")
		delete(additionalProperties, "f")
		delete(additionalProperties, "l")
		delete(additionalProperties, "T")
		delete(additionalProperties, "m")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAggregateTradeStreamsResponse struct {
	value *AggregateTradeStreamsResponse
	isSet bool
}

func (v NullableAggregateTradeStreamsResponse) Get() *AggregateTradeStreamsResponse {
	return v.value
}

func (v *NullableAggregateTradeStreamsResponse) Set(val *AggregateTradeStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregateTradeStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregateTradeStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregateTradeStreamsResponse(val *AggregateTradeStreamsResponse) *NullableAggregateTradeStreamsResponse {
	return &NullableAggregateTradeStreamsResponse{value: val, isSet: true}
}

func (v NullableAggregateTradeStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregateTradeStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
