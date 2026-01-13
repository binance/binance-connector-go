/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
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
	Smalls               *string `json:"s,omitempty"`
	Smallps              *string `json:"ps,omitempty"`
	Smallqa              *string `json:"qa,omitempty"`
	Smalld               *string `json:"d,omitempty"`
	Smallsp              *string `json:"sp,omitempty"`
	Smalldt              *int64  `json:"dt,omitempty"`
	Smallu               *int64  `json:"u,omitempty"`
	Smallot              *int64  `json:"ot,omitempty"`
	Smallcs              *string `json:"cs,omitempty"`
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

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *NewSymbolInfoResponse) SetSmallps(v string) {
	o.Smallps = &v
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

// GetDt returns the Dt field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmalldt() int64 {
	if o == nil || common.IsNil(o.Smalldt) {
		var ret int64
		return ret
	}
	return *o.Smalldt
}

// GetDtOk returns a tuple with the Dt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmalldtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalldt) {
		return nil, false
	}
	return o.Smalldt, true
}

// HasDt returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasSmalldt() bool {
	if o != nil && !common.IsNil(o.Smalldt) {
		return true
	}

	return false
}

// SetDt gets a reference to the given int64 and assigns it to the Dt field.
func (o *NewSymbolInfoResponse) SetSmalldt(v int64) {
	o.Smalldt = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmalluOk() (*int64, bool) {
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

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *NewSymbolInfoResponse) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetOt returns the Ot field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmallot() int64 {
	if o == nil || common.IsNil(o.Smallot) {
		var ret int64
		return ret
	}
	return *o.Smallot
}

// GetOtOk returns a tuple with the Ot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmallotOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallot) {
		return nil, false
	}
	return o.Smallot, true
}

// HasOt returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasSmallot() bool {
	if o != nil && !common.IsNil(o.Smallot) {
		return true
	}

	return false
}

// SetOt gets a reference to the given int64 and assigns it to the Ot field.
func (o *NewSymbolInfoResponse) SetSmallot(v int64) {
	o.Smallot = &v
}

// GetCs returns the Cs field value if set, zero value otherwise.
func (o *NewSymbolInfoResponse) GetSmallcs() string {
	if o == nil || common.IsNil(o.Smallcs) {
		var ret string
		return ret
	}
	return *o.Smallcs
}

// GetCsOk returns a tuple with the Cs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSymbolInfoResponse) GetSmallcsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallcs) {
		return nil, false
	}
	return o.Smallcs, true
}

// HasCs returns a boolean if a field has been set.
func (o *NewSymbolInfoResponse) HasSmallcs() bool {
	if o != nil && !common.IsNil(o.Smallcs) {
		return true
	}

	return false
}

// SetCs gets a reference to the given string and assigns it to the Cs field.
func (o *NewSymbolInfoResponse) SetSmallcs(v string) {
	o.Smallcs = &v
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
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallps) {
		toSerialize["ps"] = o.Smallps
	}
	if !common.IsNil(o.Smallqa) {
		toSerialize["qa"] = o.Smallqa
	}
	if !common.IsNil(o.Smalld) {
		toSerialize["d"] = o.Smalld
	}
	if !common.IsNil(o.Smallsp) {
		toSerialize["sp"] = o.Smallsp
	}
	if !common.IsNil(o.Smalldt) {
		toSerialize["dt"] = o.Smalldt
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.Smallot) {
		toSerialize["ot"] = o.Smallot
	}
	if !common.IsNil(o.Smallcs) {
		toSerialize["cs"] = o.Smallcs
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
		delete(additionalProperties, "s")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "qa")
		delete(additionalProperties, "d")
		delete(additionalProperties, "sp")
		delete(additionalProperties, "dt")
		delete(additionalProperties, "u")
		delete(additionalProperties, "ot")
		delete(additionalProperties, "cs")
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
