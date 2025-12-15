/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderTradeUpdateOInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderTradeUpdateOInner{}

// OrderTradeUpdateOInner struct for OrderTradeUpdateOInner
type OrderTradeUpdateOInner struct {
	T                    *int64                          `json:"T,omitempty"`
	Smallt               *int64                          `json:"t,omitempty"`
	Smalls               *string                         `json:"s,omitempty"`
	Smallc               *string                         `json:"c,omitempty"`
	Oid                  *string                         `json:"oid,omitempty"`
	Smallp               *string                         `json:"p,omitempty"`
	Smallq               *string                         `json:"q,omitempty"`
	Stp                  *int64                          `json:"stp,omitempty"`
	Smallr               *bool                           `json:"r,omitempty"`
	Smallpo              *bool                           `json:"po,omitempty"`
	S                    *string                         `json:"S,omitempty"`
	Smalle               *string                         `json:"e,omitempty"`
	Smallec              *string                         `json:"ec,omitempty"`
	Smallf               *string                         `json:"f,omitempty"`
	Tif                  *string                         `json:"tif,omitempty"`
	Oty                  *string                         `json:"oty,omitempty"`
	Smallfi              []OrderTradeUpdateOInnerFiInner `json:"fi,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderTradeUpdateOInner OrderTradeUpdateOInner

// NewOrderTradeUpdateOInner instantiates a new OrderTradeUpdateOInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTradeUpdateOInner() *OrderTradeUpdateOInner {
	this := OrderTradeUpdateOInner{}
	return &this
}

// NewOrderTradeUpdateOInnerWithDefaults instantiates a new OrderTradeUpdateOInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTradeUpdateOInnerWithDefaults() *OrderTradeUpdateOInner {
	this := OrderTradeUpdateOInner{}
	return &this
}

// GetT returns the T field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *OrderTradeUpdateOInner) SetT(v int64) {
	o.T = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetSmallt() int64 {
	if o == nil || common.IsNil(o.Smallt) {
		var ret int64
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetSmalltOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *OrderTradeUpdateOInner) SetSmallt(v int64) {
	o.Smallt = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *OrderTradeUpdateOInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *OrderTradeUpdateOInner) SetSmallc(v string) {
	o.Smallc = &v
}

// GetOid returns the Oid field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetOid() string {
	if o == nil || common.IsNil(o.Oid) {
		var ret string
		return ret
	}
	return *o.Oid
}

// GetOidOk returns a tuple with the Oid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetOidOk() (*string, bool) {
	if o == nil || common.IsNil(o.Oid) {
		return nil, false
	}
	return o.Oid, true
}

// HasOid returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasOid() bool {
	if o != nil && !common.IsNil(o.Oid) {
		return true
	}

	return false
}

// SetOid gets a reference to the given string and assigns it to the Oid field.
func (o *OrderTradeUpdateOInner) SetOid(v string) {
	o.Oid = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *OrderTradeUpdateOInner) SetSmallp(v string) {
	o.Smallp = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *OrderTradeUpdateOInner) SetSmallq(v string) {
	o.Smallq = &v
}

// GetStp returns the Stp field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetStp() int64 {
	if o == nil || common.IsNil(o.Stp) {
		var ret int64
		return ret
	}
	return *o.Stp
}

// GetStpOk returns a tuple with the Stp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetStpOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Stp) {
		return nil, false
	}
	return o.Stp, true
}

// HasStp returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasStp() bool {
	if o != nil && !common.IsNil(o.Stp) {
		return true
	}

	return false
}

// SetStp gets a reference to the given int64 and assigns it to the Stp field.
func (o *OrderTradeUpdateOInner) SetStp(v int64) {
	o.Stp = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetSmallr() bool {
	if o == nil || common.IsNil(o.Smallr) {
		var ret bool
		return ret
	}
	return *o.Smallr
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetSmallrOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallr) {
		return nil, false
	}
	return o.Smallr, true
}

// HasR returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasSmallr() bool {
	if o != nil && !common.IsNil(o.Smallr) {
		return true
	}

	return false
}

// SetR gets a reference to the given bool and assigns it to the R field.
func (o *OrderTradeUpdateOInner) SetSmallr(v bool) {
	o.Smallr = &v
}

// GetPo returns the Po field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetSmallpo() bool {
	if o == nil || common.IsNil(o.Smallpo) {
		var ret bool
		return ret
	}
	return *o.Smallpo
}

// GetPoOk returns a tuple with the Po field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetSmallpoOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallpo) {
		return nil, false
	}
	return o.Smallpo, true
}

// HasPo returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasSmallpo() bool {
	if o != nil && !common.IsNil(o.Smallpo) {
		return true
	}

	return false
}

// SetPo gets a reference to the given bool and assigns it to the Po field.
func (o *OrderTradeUpdateOInner) SetSmallpo(v bool) {
	o.Smallpo = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *OrderTradeUpdateOInner) SetS(v string) {
	o.S = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *OrderTradeUpdateOInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetEc returns the Ec field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetSmallec() string {
	if o == nil || common.IsNil(o.Smallec) {
		var ret string
		return ret
	}
	return *o.Smallec
}

// GetEcOk returns a tuple with the Ec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetSmallecOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallec) {
		return nil, false
	}
	return o.Smallec, true
}

// HasEc returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasSmallec() bool {
	if o != nil && !common.IsNil(o.Smallec) {
		return true
	}

	return false
}

// SetEc gets a reference to the given string and assigns it to the Ec field.
func (o *OrderTradeUpdateOInner) SetSmallec(v string) {
	o.Smallec = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetSmallf() string {
	if o == nil || common.IsNil(o.Smallf) {
		var ret string
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetSmallfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *OrderTradeUpdateOInner) SetSmallf(v string) {
	o.Smallf = &v
}

// GetTif returns the Tif field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetTif() string {
	if o == nil || common.IsNil(o.Tif) {
		var ret string
		return ret
	}
	return *o.Tif
}

// GetTifOk returns a tuple with the Tif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetTifOk() (*string, bool) {
	if o == nil || common.IsNil(o.Tif) {
		return nil, false
	}
	return o.Tif, true
}

// HasTif returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasTif() bool {
	if o != nil && !common.IsNil(o.Tif) {
		return true
	}

	return false
}

// SetTif gets a reference to the given string and assigns it to the Tif field.
func (o *OrderTradeUpdateOInner) SetTif(v string) {
	o.Tif = &v
}

// GetOty returns the Oty field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetOty() string {
	if o == nil || common.IsNil(o.Oty) {
		var ret string
		return ret
	}
	return *o.Oty
}

// GetOtyOk returns a tuple with the Oty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetOtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Oty) {
		return nil, false
	}
	return o.Oty, true
}

// HasOty returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasOty() bool {
	if o != nil && !common.IsNil(o.Oty) {
		return true
	}

	return false
}

// SetOty gets a reference to the given string and assigns it to the Oty field.
func (o *OrderTradeUpdateOInner) SetOty(v string) {
	o.Oty = &v
}

// GetFi returns the Fi field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInner) GetSmallfi() []OrderTradeUpdateOInnerFiInner {
	if o == nil || common.IsNil(o.Smallfi) {
		var ret []OrderTradeUpdateOInnerFiInner
		return ret
	}
	return o.Smallfi
}

// GetFiOk returns a tuple with the Fi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInner) GetSmallfiOk() ([]OrderTradeUpdateOInnerFiInner, bool) {
	if o == nil || common.IsNil(o.Smallfi) {
		return nil, false
	}
	return o.Smallfi, true
}

// HasFi returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInner) HasSmallfi() bool {
	if o != nil && !common.IsNil(o.Smallfi) {
		return true
	}

	return false
}

// SetFi gets a reference to the given []OrderTradeUpdateOInnerFiInner and assigns it to the Fi field.
func (o *OrderTradeUpdateOInner) SetSmallfi(v []OrderTradeUpdateOInnerFiInner) {
	o.Smallfi = v
}

func (o OrderTradeUpdateOInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTradeUpdateOInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.Oid) {
		toSerialize["oid"] = o.Oid
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.Stp) {
		toSerialize["stp"] = o.Stp
	}
	if !common.IsNil(o.Smallr) {
		toSerialize["r"] = o.Smallr
	}
	if !common.IsNil(o.Smallpo) {
		toSerialize["po"] = o.Smallpo
	}
	if !common.IsNil(o.S) {
		toSerialize["S"] = o.S
	}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.Smallec) {
		toSerialize["ec"] = o.Smallec
	}
	if !common.IsNil(o.Smallf) {
		toSerialize["f"] = o.Smallf
	}
	if !common.IsNil(o.Tif) {
		toSerialize["tif"] = o.Tif
	}
	if !common.IsNil(o.Oty) {
		toSerialize["oty"] = o.Oty
	}
	if !common.IsNil(o.Smallfi) {
		toSerialize["fi"] = o.Smallfi
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderTradeUpdateOInner) UnmarshalJSON(data []byte) (err error) {
	varOrderTradeUpdateOInner := _OrderTradeUpdateOInner{}

	err = json.Unmarshal(data, &varOrderTradeUpdateOInner)

	if err != nil {
		return err
	}

	*o = OrderTradeUpdateOInner(varOrderTradeUpdateOInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "T")
		delete(additionalProperties, "t")
		delete(additionalProperties, "s")
		delete(additionalProperties, "c")
		delete(additionalProperties, "oid")
		delete(additionalProperties, "p")
		delete(additionalProperties, "q")
		delete(additionalProperties, "stp")
		delete(additionalProperties, "r")
		delete(additionalProperties, "po")
		delete(additionalProperties, "S")
		delete(additionalProperties, "e")
		delete(additionalProperties, "ec")
		delete(additionalProperties, "f")
		delete(additionalProperties, "tif")
		delete(additionalProperties, "oty")
		delete(additionalProperties, "fi")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderTradeUpdateOInner struct {
	value *OrderTradeUpdateOInner
	isSet bool
}

func (v NullableOrderTradeUpdateOInner) Get() *OrderTradeUpdateOInner {
	return v.value
}

func (v *NullableOrderTradeUpdateOInner) Set(val *OrderTradeUpdateOInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTradeUpdateOInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTradeUpdateOInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTradeUpdateOInner(val *OrderTradeUpdateOInner) *NullableOrderTradeUpdateOInner {
	return &NullableOrderTradeUpdateOInner{value: val, isSet: true}
}

func (v NullableOrderTradeUpdateOInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTradeUpdateOInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
