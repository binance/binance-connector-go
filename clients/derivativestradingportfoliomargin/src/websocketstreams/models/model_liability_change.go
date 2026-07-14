/*
Portfolio Margin WebSocket Market Streams

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the LiabilityChange type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LiabilityChange{}

// LiabilityChange struct for LiabilityChange
type LiabilityChange struct {
	// Event Time
	E *int64 `json:"E,omitempty"`
	// Asset
	Smalla *string `json:"a,omitempty"`
	// Type
	Smallt *string `json:"t,omitempty"`
	// Transaction ID
	T *int64 `json:"T,omitempty"`
	// Principal
	Smallp *string `json:"p,omitempty"`
	// Interest
	Smalli *string `json:"i,omitempty"`
	// Total Liability
	Smalll               *string `json:"l,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LiabilityChange LiabilityChange

// NewLiabilityChange instantiates a new LiabilityChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiabilityChange() *LiabilityChange {
	this := LiabilityChange{}
	return &this
}

// NewLiabilityChangeWithDefaults instantiates a new LiabilityChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiabilityChangeWithDefaults() *LiabilityChange {
	this := LiabilityChange{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *LiabilityChange) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiabilityChange) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *LiabilityChange) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *LiabilityChange) SetE(v int64) {
	o.E = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *LiabilityChange) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiabilityChange) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *LiabilityChange) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *LiabilityChange) SetSmalla(v string) {
	o.Smalla = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *LiabilityChange) GetSmallt() string {
	if o == nil || common.IsNil(o.Smallt) {
		var ret string
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiabilityChange) GetSmalltOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *LiabilityChange) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given string and assigns it to the T field.
func (o *LiabilityChange) SetSmallt(v string) {
	o.Smallt = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *LiabilityChange) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiabilityChange) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *LiabilityChange) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *LiabilityChange) SetT(v int64) {
	o.T = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *LiabilityChange) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiabilityChange) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *LiabilityChange) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *LiabilityChange) SetSmallp(v string) {
	o.Smallp = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *LiabilityChange) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiabilityChange) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *LiabilityChange) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *LiabilityChange) SetSmalli(v string) {
	o.Smalli = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *LiabilityChange) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiabilityChange) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *LiabilityChange) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *LiabilityChange) SetSmalll(v string) {
	o.Smalll = &v
}

func (o LiabilityChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiabilityChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LiabilityChange) UnmarshalJSON(data []byte) (err error) {
	varLiabilityChange := _LiabilityChange{}

	err = json.Unmarshal(data, &varLiabilityChange)

	if err != nil {
		return err
	}

	*o = LiabilityChange(varLiabilityChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "a")
		delete(additionalProperties, "t")
		delete(additionalProperties, "T")
		delete(additionalProperties, "p")
		delete(additionalProperties, "i")
		delete(additionalProperties, "l")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLiabilityChange struct {
	value *LiabilityChange
	isSet bool
}

func (v NullableLiabilityChange) Get() *LiabilityChange {
	return v.value
}

func (v *NullableLiabilityChange) Set(val *LiabilityChange) {
	v.value = val
	v.isSet = true
}

func (v NullableLiabilityChange) IsSet() bool {
	return v.isSet
}

func (v *NullableLiabilityChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiabilityChange(val *LiabilityChange) *NullableLiabilityChange {
	return &NullableLiabilityChange{value: val, isSet: true}
}

func (v NullableLiabilityChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiabilityChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
