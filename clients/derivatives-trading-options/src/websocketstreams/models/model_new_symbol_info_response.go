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

// checks if the NewSymbolInfoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NewSymbolInfoResponse{}

// NewSymbolInfoResponse struct for NewSymbolInfoResponse
type NewSymbolInfoResponse struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smallu               *string `json:"u,omitempty"`
	Smallqa              *string `json:"qa,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Unit                 *int64  `json:"unit,omitempty"`
	Smallmq              *string `json:"mq,omitempty"`
	Smalld               *string `json:"d,omitempty"`
	Smallsp              *string `json:"sp,omitempty"`
	Smalled              *int64  `json:"ed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NewSymbolInfoResponse NewSymbolInfoResponse

// NewNewSymbolInfoResponse instantiates a new NewSymbolInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewSymbolInfoResponse() *NewSymbolInfoResponse {
	this := NewSymbolInfoResponse{}
	return &this
}

// NewNewSymbolInfoResponseWithDefaults instantiates a new NewSymbolInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewSymbolInfoResponseWithDefaults() *NewSymbolInfoResponse {
	this := NewSymbolInfoResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *NewSymbolInfoResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *NewSymbolInfoResponse) SetE(v int64) {
	o.E = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmallu() string {
	if o == nil || common.IsNil(o.Smallu) {
		var ret string
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmalluOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given string and assigns it to the U field.
func (o *NewSymbolInfoResponse) SetSmallu(v string) {
	o.Smallu = &v
}

// GetQa returns the Qa field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmallqa() string {
	if o == nil || common.IsNil(o.Smallqa) {
		var ret string
		return ret
	}
	return *o.Smallqa
}

// GetQaOk returns a tuple with the Qa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmallqaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallqa) {
		return nil, false
	}
	return o.Smallqa, true
}

// HasQa returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasSmallqa() bool {
	if o != nil && !common.IsNil(o.Smallqa) {
		return true
	}

	return false
}

// SetQa gets a reference to the given string and assigns it to the Qa field.
func (o *NewSymbolInfoResponse) SetSmallqa(v string) {
	o.Smallqa = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *NewSymbolInfoResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetUnit() int64 {
	if o == nil || common.IsNil(o.Unit) {
		var ret int64
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetUnitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasUnit() bool {
	if o != nil && !common.IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given int64 and assigns it to the Unit field.
func (o *NewSymbolInfoResponse) SetUnit(v int64) {
	o.Unit = &v
}

// GetMq returns the Mq field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmallmq() string {
	if o == nil || common.IsNil(o.Smallmq) {
		var ret string
		return ret
	}
	return *o.Smallmq
}

// GetMqOk returns a tuple with the Mq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmallmqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallmq) {
		return nil, false
	}
	return o.Smallmq, true
}

// HasMq returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasSmallmq() bool {
	if o != nil && !common.IsNil(o.Smallmq) {
		return true
	}

	return false
}

// SetMq gets a reference to the given string and assigns it to the Mq field.
func (o *NewSymbolInfoResponse) SetSmallmq(v string) {
	o.Smallmq = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmalld() string {
	if o == nil || common.IsNil(o.Smalld) {
		var ret string
		return ret
	}
	return *o.Smalld
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmalldOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalld) {
		return nil, false
	}
	return o.Smalld, true
}

// HasD returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasSmalld() bool {
	if o != nil && !common.IsNil(o.Smalld) {
		return true
	}

	return false
}

// SetD gets a reference to the given string and assigns it to the D field.
func (o *NewSymbolInfoResponse) SetSmalld(v string) {
	o.Smalld = &v
}

// GetSp returns the Sp field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmallsp() string {
	if o == nil || common.IsNil(o.Smallsp) {
		var ret string
		return ret
	}
	return *o.Smallsp
}

// GetSpOk returns a tuple with the Sp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmallspOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallsp) {
		return nil, false
	}
	return o.Smallsp, true
}

// HasSp returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasSmallsp() bool {
	if o != nil && !common.IsNil(o.Smallsp) {
		return true
	}

	return false
}

// SetSp gets a reference to the given string and assigns it to the Sp field.
func (o *NewSymbolInfoResponse) SetSmallsp(v string) {
	o.Smallsp = &v
}

// GetEd returns the Ed field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmalled() int64 {
	if o == nil || common.IsNil(o.Smalled) {
		var ret int64
		return ret
	}
	return *o.Smalled
}

// GetEdOk returns a tuple with the Ed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmalledOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalled) {
		return nil, false
	}
	return o.Smalled, true
}

// HasEd returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasSmalled() bool {
	if o != nil && !common.IsNil(o.Smalled) {
		return true
	}

	return false
}

// SetEd gets a reference to the given int64 and assigns it to the Ed field.
func (o *NewSymbolInfoResponse) SetSmalled(v int64) {
	o.Smalled = &v
}

func (o NewSymbolInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewSymbolInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.Smallqa) {
		toSerialize["qa"] = o.Smallqa
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !common.IsNil(o.Smallmq) {
		toSerialize["mq"] = o.Smallmq
	}
	if !common.IsNil(o.Smalld) {
		toSerialize["d"] = o.Smalld
	}
	if !common.IsNil(o.Smallsp) {
		toSerialize["sp"] = o.Smallsp
	}
	if !common.IsNil(o.Smalled) {
		toSerialize["ed"] = o.Smalled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NewSymbolInfoResponse) UnmarshalJSON(data []byte) (err error) {
	varNewSymbolInfoResponse := _NewSymbolInfoResponse{}

	err = json.Unmarshal(data, &varNewSymbolInfoResponse)

	if err != nil {
		return err
	}

	*o = NewSymbolInfoResponse(varNewSymbolInfoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "u")
		delete(additionalProperties, "qa")
		delete(additionalProperties, "s")
		delete(additionalProperties, "unit")
		delete(additionalProperties, "mq")
		delete(additionalProperties, "d")
		delete(additionalProperties, "sp")
		delete(additionalProperties, "ed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNewSymbolInfoResponse struct {
	value *NewSymbolInfoResponse
	isSet bool
}

func (v NullableNewSymbolInfoResponse) Get() *NewSymbolInfoResponse {
	return v.value
}

func (v *NullableNewSymbolInfoResponse) Set(val *NewSymbolInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNewSymbolInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNewSymbolInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewSymbolInfoResponse(val *NewSymbolInfoResponse) *NullableNewSymbolInfoResponse {
	return &NullableNewSymbolInfoResponse{value: val, isSet: true}
}

func (v NullableNewSymbolInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewSymbolInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
