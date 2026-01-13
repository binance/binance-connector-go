/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AllMarketLiquidationOrderStreamsResponseO type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllMarketLiquidationOrderStreamsResponseO{}

// AllMarketLiquidationOrderStreamsResponseO struct for AllMarketLiquidationOrderStreamsResponseO
type AllMarketLiquidationOrderStreamsResponseO struct {
	Smalls               *string `json:"s,omitempty"`
	Smallps              *string `json:"ps,omitempty"`
	S                    *string `json:"S,omitempty"`
	Smallo               *string `json:"o,omitempty"`
	Smallf               *string `json:"f,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	Smallap              *string `json:"ap,omitempty"`
	X                    *string `json:"X,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	Smallz               *string `json:"z,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllMarketLiquidationOrderStreamsResponseO AllMarketLiquidationOrderStreamsResponseO

// NewAllMarketLiquidationOrderStreamsResponseO instantiates a new AllMarketLiquidationOrderStreamsResponseO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllMarketLiquidationOrderStreamsResponseO() *AllMarketLiquidationOrderStreamsResponseO {
	this := AllMarketLiquidationOrderStreamsResponseO{}
	return &this
}

// NewAllMarketLiquidationOrderStreamsResponseOWithDefaults instantiates a new AllMarketLiquidationOrderStreamsResponseO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllMarketLiquidationOrderStreamsResponseOWithDefaults() *AllMarketLiquidationOrderStreamsResponseO {
	this := AllMarketLiquidationOrderStreamsResponseO{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AllMarketLiquidationOrderStreamsResponseO) SetSmalls(v string) {
	o.Smalls = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *AllMarketLiquidationOrderStreamsResponseO) SetSmallps(v string) {
	o.Smallps = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AllMarketLiquidationOrderStreamsResponseO) SetS(v string) {
	o.S = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *AllMarketLiquidationOrderStreamsResponseO) SetSmallo(v string) {
	o.Smallo = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallf() string {
	if o == nil || common.IsNil(o.Smallf) {
		var ret string
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *AllMarketLiquidationOrderStreamsResponseO) SetSmallf(v string) {
	o.Smallf = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AllMarketLiquidationOrderStreamsResponseO) SetSmallq(v string) {
	o.Smallq = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AllMarketLiquidationOrderStreamsResponseO) SetSmallp(v string) {
	o.Smallp = &v
}

// GetAp returns the Ap field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallap() string {
	if o == nil || common.IsNil(o.Smallap) {
		var ret string
		return ret
	}
	return *o.Smallap
}

// GetApOk returns a tuple with the Ap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallapOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallap) {
		return nil, false
	}
	return o.Smallap, true
}

// HasAp returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) HasSmallap() bool {
	if o != nil && !common.IsNil(o.Smallap) {
		return true
	}

	return false
}

// SetAp gets a reference to the given string and assigns it to the Ap field.
func (o *AllMarketLiquidationOrderStreamsResponseO) SetSmallap(v string) {
	o.Smallap = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetX() string {
	if o == nil || common.IsNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetXOk() (*string, bool) {
	if o == nil || common.IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) HasX() bool {
	if o != nil && !common.IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *AllMarketLiquidationOrderStreamsResponseO) SetX(v string) {
	o.X = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *AllMarketLiquidationOrderStreamsResponseO) SetSmalll(v string) {
	o.Smalll = &v
}

// GetZ returns the Z field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallz() string {
	if o == nil || common.IsNil(o.Smallz) {
		var ret string
		return ret
	}
	return *o.Smallz
}

// GetZOk returns a tuple with the Z field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetSmallzOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallz) {
		return nil, false
	}
	return o.Smallz, true
}

// HasZ returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) HasSmallz() bool {
	if o != nil && !common.IsNil(o.Smallz) {
		return true
	}

	return false
}

// SetZ gets a reference to the given string and assigns it to the Z field.
func (o *AllMarketLiquidationOrderStreamsResponseO) SetSmallz(v string) {
	o.Smallz = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponseO) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AllMarketLiquidationOrderStreamsResponseO) SetT(v int64) {
	o.T = &v
}

func (o AllMarketLiquidationOrderStreamsResponseO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllMarketLiquidationOrderStreamsResponseO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallps) {
		toSerialize["ps"] = o.Smallps
	}
	if !common.IsNil(o.S) {
		toSerialize["S"] = o.S
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}
	if !common.IsNil(o.Smallf) {
		toSerialize["f"] = o.Smallf
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smallap) {
		toSerialize["ap"] = o.Smallap
	}
	if !common.IsNil(o.X) {
		toSerialize["X"] = o.X
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.Smallz) {
		toSerialize["z"] = o.Smallz
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AllMarketLiquidationOrderStreamsResponseO) UnmarshalJSON(data []byte) (err error) {
	varAllMarketLiquidationOrderStreamsResponseO := _AllMarketLiquidationOrderStreamsResponseO{}

	err = json.Unmarshal(data, &varAllMarketLiquidationOrderStreamsResponseO)

	if err != nil {
		return err
	}

	*o = AllMarketLiquidationOrderStreamsResponseO(varAllMarketLiquidationOrderStreamsResponseO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "S")
		delete(additionalProperties, "o")
		delete(additionalProperties, "f")
		delete(additionalProperties, "q")
		delete(additionalProperties, "p")
		delete(additionalProperties, "ap")
		delete(additionalProperties, "X")
		delete(additionalProperties, "l")
		delete(additionalProperties, "z")
		delete(additionalProperties, "T")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAllMarketLiquidationOrderStreamsResponseO struct {
	value *AllMarketLiquidationOrderStreamsResponseO
	isSet bool
}

func (v NullableAllMarketLiquidationOrderStreamsResponseO) Get() *AllMarketLiquidationOrderStreamsResponseO {
	return v.value
}

func (v *NullableAllMarketLiquidationOrderStreamsResponseO) Set(val *AllMarketLiquidationOrderStreamsResponseO) {
	v.value = val
	v.isSet = true
}

func (v NullableAllMarketLiquidationOrderStreamsResponseO) IsSet() bool {
	return v.isSet
}

func (v *NullableAllMarketLiquidationOrderStreamsResponseO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllMarketLiquidationOrderStreamsResponseO(val *AllMarketLiquidationOrderStreamsResponseO) *NullableAllMarketLiquidationOrderStreamsResponseO {
	return &NullableAllMarketLiquidationOrderStreamsResponseO{value: val, isSet: true}
}

func (v NullableAllMarketLiquidationOrderStreamsResponseO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllMarketLiquidationOrderStreamsResponseO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
