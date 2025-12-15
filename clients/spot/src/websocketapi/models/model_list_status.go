/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ListStatus type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListStatus{}

// ListStatus struct for ListStatus
type ListStatus struct {
	E                    *int64             `json:"E,omitempty"`
	Smalls               *string            `json:"s,omitempty"`
	Smallg               *int64             `json:"g,omitempty"`
	Smallc               *string            `json:"c,omitempty"`
	Smalll               *string            `json:"l,omitempty"`
	L                    *string            `json:"L,omitempty"`
	Smallr               *string            `json:"r,omitempty"`
	C                    *string            `json:"C,omitempty"`
	T                    *int64             `json:"T,omitempty"`
	O                    []ListStatusOInner `json:"O,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListStatus ListStatus

// NewListStatus instantiates a new ListStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStatus() *ListStatus {
	this := ListStatus{}
	return &this
}

// NewListStatusWithDefaults instantiates a new ListStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStatusWithDefaults() *ListStatus {
	this := ListStatus{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *ListStatus) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatus) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *ListStatus) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *ListStatus) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *ListStatus) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatus) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *ListStatus) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *ListStatus) SetSmalls(v string) {
	o.Smalls = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *ListStatus) GetSmallg() int64 {
	if o == nil || common.IsNil(o.Smallg) {
		var ret int64
		return ret
	}
	return *o.Smallg
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatus) GetSmallgOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallg) {
		return nil, false
	}
	return o.Smallg, true
}

// HasG returns a boolean if a field has been set.
func (o *ListStatus) HasSmallg() bool {
	if o != nil && !common.IsNil(o.Smallg) {
		return true
	}

	return false
}

// SetG gets a reference to the given int64 and assigns it to the G field.
func (o *ListStatus) SetSmallg(v int64) {
	o.Smallg = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *ListStatus) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatus) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *ListStatus) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *ListStatus) SetSmallc(v string) {
	o.Smallc = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *ListStatus) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatus) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *ListStatus) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *ListStatus) SetSmalll(v string) {
	o.Smalll = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *ListStatus) GetL() string {
	if o == nil || common.IsNil(o.L) {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatus) GetLOk() (*string, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *ListStatus) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *ListStatus) SetL(v string) {
	o.L = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *ListStatus) GetSmallr() string {
	if o == nil || common.IsNil(o.Smallr) {
		var ret string
		return ret
	}
	return *o.Smallr
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatus) GetSmallrOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallr) {
		return nil, false
	}
	return o.Smallr, true
}

// HasR returns a boolean if a field has been set.
func (o *ListStatus) HasSmallr() bool {
	if o != nil && !common.IsNil(o.Smallr) {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *ListStatus) SetSmallr(v string) {
	o.Smallr = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *ListStatus) GetC() string {
	if o == nil || common.IsNil(o.C) {
		var ret string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatus) GetCOk() (*string, bool) {
	if o == nil || common.IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *ListStatus) HasC() bool {
	if o != nil && !common.IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *ListStatus) SetC(v string) {
	o.C = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *ListStatus) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatus) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *ListStatus) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *ListStatus) SetT(v int64) {
	o.T = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *ListStatus) GetO() []ListStatusOInner {
	if o == nil || common.IsNil(o.O) {
		var ret []ListStatusOInner
		return ret
	}
	return o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatus) GetOOk() ([]ListStatusOInner, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *ListStatus) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given []ListStatusOInner and assigns it to the O field.
func (o *ListStatus) SetO(v []ListStatusOInner) {
	o.O = v
}

func (o ListStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallg) {
		toSerialize["g"] = o.Smallg
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.L) {
		toSerialize["L"] = o.L
	}
	if !common.IsNil(o.Smallr) {
		toSerialize["r"] = o.Smallr
	}
	if !common.IsNil(o.C) {
		toSerialize["C"] = o.C
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.O) {
		toSerialize["O"] = o.O
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListStatus) UnmarshalJSON(data []byte) (err error) {
	varListStatus := _ListStatus{}

	err = json.Unmarshal(data, &varListStatus)

	if err != nil {
		return err
	}

	*o = ListStatus(varListStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "g")
		delete(additionalProperties, "c")
		delete(additionalProperties, "l")
		delete(additionalProperties, "L")
		delete(additionalProperties, "r")
		delete(additionalProperties, "C")
		delete(additionalProperties, "T")
		delete(additionalProperties, "O")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListStatus struct {
	value *ListStatus
	isSet bool
}

func (v NullableListStatus) Get() *ListStatus {
	return v.value
}

func (v *NullableListStatus) Set(val *ListStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableListStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableListStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStatus(val *ListStatus) *NullableListStatus {
	return &NullableListStatus{value: val, isSet: true}
}

func (v NullableListStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
