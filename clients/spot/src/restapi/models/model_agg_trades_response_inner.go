/*
Spot REST API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AggTradesResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AggTradesResponseInner{}

// AggTradesResponseInner struct for AggTradesResponseInner
type AggTradesResponseInner struct {
	// Aggregate tradeId
	Smalla *int64 `json:"a,omitempty"`
	// Price
	Smallp *string `json:"p,omitempty"`
	// Quantity
	Smallq *string `json:"q,omitempty"`
	// First tradeId
	Smallf *int64 `json:"f,omitempty"`
	// Last tradeId
	Smalll *int64 `json:"l,omitempty"`
	// Timestamp
	T *int64 `json:"T,omitempty"`
	// Was the buyer the maker?
	Smallm *bool `json:"m,omitempty"`
	// Was the trade the best price match?
	M                    *bool `json:"M,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AggTradesResponseInner AggTradesResponseInner

// NewAggTradesResponseInner instantiates a new AggTradesResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggTradesResponseInner() *AggTradesResponseInner {
	this := AggTradesResponseInner{}
	return &this
}

// NewAggTradesResponseInnerWithDefaults instantiates a new AggTradesResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggTradesResponseInnerWithDefaults() *AggTradesResponseInner {
	this := AggTradesResponseInner{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *AggTradesResponseInner) GetSmalla() int64 {
	if o == nil || common.IsNil(o.Smalla) {
		var ret int64
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradesResponseInner) GetSmallaOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *AggTradesResponseInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given int64 and assigns it to the A field.
func (o *AggTradesResponseInner) SetSmalla(v int64) {
	o.Smalla = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AggTradesResponseInner) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradesResponseInner) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *AggTradesResponseInner) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AggTradesResponseInner) SetSmallp(v string) {
	o.Smallp = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AggTradesResponseInner) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradesResponseInner) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *AggTradesResponseInner) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AggTradesResponseInner) SetSmallq(v string) {
	o.Smallq = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *AggTradesResponseInner) GetSmallf() int64 {
	if o == nil || common.IsNil(o.Smallf) {
		var ret int64
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradesResponseInner) GetSmallfOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *AggTradesResponseInner) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given int64 and assigns it to the F field.
func (o *AggTradesResponseInner) SetSmallf(v int64) {
	o.Smallf = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AggTradesResponseInner) GetSmalll() int64 {
	if o == nil || common.IsNil(o.Smalll) {
		var ret int64
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradesResponseInner) GetSmalllOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *AggTradesResponseInner) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *AggTradesResponseInner) SetSmalll(v int64) {
	o.Smalll = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AggTradesResponseInner) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradesResponseInner) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AggTradesResponseInner) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AggTradesResponseInner) SetT(v int64) {
	o.T = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *AggTradesResponseInner) GetSmallm() bool {
	if o == nil || common.IsNil(o.Smallm) {
		var ret bool
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradesResponseInner) GetSmallmOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *AggTradesResponseInner) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *AggTradesResponseInner) SetSmallm(v bool) {
	o.Smallm = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *AggTradesResponseInner) GetM() bool {
	if o == nil || common.IsNil(o.M) {
		var ret bool
		return ret
	}
	return *o.M
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradesResponseInner) GetMOk() (*bool, bool) {
	if o == nil || common.IsNil(o.M) {
		return nil, false
	}
	return o.M, true
}

// HasM returns a boolean if a field has been set.
func (o *AggTradesResponseInner) HasM() bool {
	if o != nil && !common.IsNil(o.M) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *AggTradesResponseInner) SetM(v bool) {
	o.M = &v
}

func (o AggTradesResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggTradesResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
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
	if !common.IsNil(o.M) {
		toSerialize["M"] = o.M
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AggTradesResponseInner) UnmarshalJSON(data []byte) (err error) {
	varAggTradesResponseInner := _AggTradesResponseInner{}

	err = json.Unmarshal(data, &varAggTradesResponseInner)

	if err != nil {
		return err
	}

	*o = AggTradesResponseInner(varAggTradesResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "a")
		delete(additionalProperties, "p")
		delete(additionalProperties, "q")
		delete(additionalProperties, "f")
		delete(additionalProperties, "l")
		delete(additionalProperties, "T")
		delete(additionalProperties, "m")
		delete(additionalProperties, "M")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAggTradesResponseInner struct {
	value *AggTradesResponseInner
	isSet bool
}

func (v NullableAggTradesResponseInner) Get() *AggTradesResponseInner {
	return v.value
}

func (v *NullableAggTradesResponseInner) Set(val *AggTradesResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAggTradesResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAggTradesResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggTradesResponseInner(val *AggTradesResponseInner) *NullableAggTradesResponseInner {
	return &NullableAggTradesResponseInner{value: val, isSet: true}
}

func (v NullableAggTradesResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggTradesResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
