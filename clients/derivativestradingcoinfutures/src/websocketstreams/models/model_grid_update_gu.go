/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GridUpdateGu type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GridUpdateGu{}

// GridUpdateGu struct for GridUpdateGu
type GridUpdateGu struct {
	Smallsi              *int64  `json:"si,omitempty"`
	Smallst              *string `json:"st,omitempty"`
	Smallss              *string `json:"ss,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallr               *string `json:"r,omitempty"`
	Smallup              *string `json:"up,omitempty"`
	Smalluq              *string `json:"uq,omitempty"`
	Smalluf              *string `json:"uf,omitempty"`
	Smallmp              *string `json:"mp,omitempty"`
	Smallut              *int64  `json:"ut,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GridUpdateGu GridUpdateGu

// NewGridUpdateGu instantiates a new GridUpdateGu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridUpdateGu() *GridUpdateGu {
	this := GridUpdateGu{}
	return &this
}

// NewGridUpdateGuWithDefaults instantiates a new GridUpdateGu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridUpdateGuWithDefaults() *GridUpdateGu {
	this := GridUpdateGu{}
	return &this
}

// GetSi returns the Si field value if set, zero value otherwise.
func (o *GridUpdateGu) GetSmallsi() int64 {
	if o == nil || common.IsNil(o.Smallsi) {
		var ret int64
		return ret
	}
	return *o.Smallsi
}

// GetSiOk returns a tuple with the Si field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdateGu) GetSmallsiOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallsi) {
		return nil, false
	}
	return o.Smallsi, true
}

// HasSi returns a boolean if a field has been set.
func (o *GridUpdateGu) HasSmallsi() bool {
	if o != nil && !common.IsNil(o.Smallsi) {
		return true
	}

	return false
}

// SetSi gets a reference to the given int64 and assigns it to the Si field.
func (o *GridUpdateGu) SetSmallsi(v int64) {
	o.Smallsi = &v
}

// GetSt returns the St field value if set, zero value otherwise.
func (o *GridUpdateGu) GetSmallst() string {
	if o == nil || common.IsNil(o.Smallst) {
		var ret string
		return ret
	}
	return *o.Smallst
}

// GetStOk returns a tuple with the St field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdateGu) GetSmallstOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallst) {
		return nil, false
	}
	return o.Smallst, true
}

// HasSt returns a boolean if a field has been set.
func (o *GridUpdateGu) HasSmallst() bool {
	if o != nil && !common.IsNil(o.Smallst) {
		return true
	}

	return false
}

// SetSt gets a reference to the given string and assigns it to the St field.
func (o *GridUpdateGu) SetSmallst(v string) {
	o.Smallst = &v
}

// GetSs returns the Ss field value if set, zero value otherwise.
func (o *GridUpdateGu) GetSmallss() string {
	if o == nil || common.IsNil(o.Smallss) {
		var ret string
		return ret
	}
	return *o.Smallss
}

// GetSsOk returns a tuple with the Ss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdateGu) GetSmallssOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallss) {
		return nil, false
	}
	return o.Smallss, true
}

// HasSs returns a boolean if a field has been set.
func (o *GridUpdateGu) HasSmallss() bool {
	if o != nil && !common.IsNil(o.Smallss) {
		return true
	}

	return false
}

// SetSs gets a reference to the given string and assigns it to the Ss field.
func (o *GridUpdateGu) SetSmallss(v string) {
	o.Smallss = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *GridUpdateGu) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdateGu) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *GridUpdateGu) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *GridUpdateGu) SetSmalls(v string) {
	o.Smalls = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *GridUpdateGu) GetSmallr() string {
	if o == nil || common.IsNil(o.Smallr) {
		var ret string
		return ret
	}
	return *o.Smallr
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdateGu) GetSmallrOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallr) {
		return nil, false
	}
	return o.Smallr, true
}

// HasR returns a boolean if a field has been set.
func (o *GridUpdateGu) HasSmallr() bool {
	if o != nil && !common.IsNil(o.Smallr) {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *GridUpdateGu) SetSmallr(v string) {
	o.Smallr = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *GridUpdateGu) GetSmallup() string {
	if o == nil || common.IsNil(o.Smallup) {
		var ret string
		return ret
	}
	return *o.Smallup
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdateGu) GetSmallupOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallup) {
		return nil, false
	}
	return o.Smallup, true
}

// HasUp returns a boolean if a field has been set.
func (o *GridUpdateGu) HasSmallup() bool {
	if o != nil && !common.IsNil(o.Smallup) {
		return true
	}

	return false
}

// SetUp gets a reference to the given string and assigns it to the Up field.
func (o *GridUpdateGu) SetSmallup(v string) {
	o.Smallup = &v
}

// GetUq returns the Uq field value if set, zero value otherwise.
func (o *GridUpdateGu) GetSmalluq() string {
	if o == nil || common.IsNil(o.Smalluq) {
		var ret string
		return ret
	}
	return *o.Smalluq
}

// GetUqOk returns a tuple with the Uq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdateGu) GetSmalluqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalluq) {
		return nil, false
	}
	return o.Smalluq, true
}

// HasUq returns a boolean if a field has been set.
func (o *GridUpdateGu) HasSmalluq() bool {
	if o != nil && !common.IsNil(o.Smalluq) {
		return true
	}

	return false
}

// SetUq gets a reference to the given string and assigns it to the Uq field.
func (o *GridUpdateGu) SetSmalluq(v string) {
	o.Smalluq = &v
}

// GetUf returns the Uf field value if set, zero value otherwise.
func (o *GridUpdateGu) GetSmalluf() string {
	if o == nil || common.IsNil(o.Smalluf) {
		var ret string
		return ret
	}
	return *o.Smalluf
}

// GetUfOk returns a tuple with the Uf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdateGu) GetSmallufOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalluf) {
		return nil, false
	}
	return o.Smalluf, true
}

// HasUf returns a boolean if a field has been set.
func (o *GridUpdateGu) HasSmalluf() bool {
	if o != nil && !common.IsNil(o.Smalluf) {
		return true
	}

	return false
}

// SetUf gets a reference to the given string and assigns it to the Uf field.
func (o *GridUpdateGu) SetSmalluf(v string) {
	o.Smalluf = &v
}

// GetMp returns the Mp field value if set, zero value otherwise.
func (o *GridUpdateGu) GetSmallmp() string {
	if o == nil || common.IsNil(o.Smallmp) {
		var ret string
		return ret
	}
	return *o.Smallmp
}

// GetMpOk returns a tuple with the Mp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdateGu) GetSmallmpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallmp) {
		return nil, false
	}
	return o.Smallmp, true
}

// HasMp returns a boolean if a field has been set.
func (o *GridUpdateGu) HasSmallmp() bool {
	if o != nil && !common.IsNil(o.Smallmp) {
		return true
	}

	return false
}

// SetMp gets a reference to the given string and assigns it to the Mp field.
func (o *GridUpdateGu) SetSmallmp(v string) {
	o.Smallmp = &v
}

// GetUt returns the Ut field value if set, zero value otherwise.
func (o *GridUpdateGu) GetSmallut() int64 {
	if o == nil || common.IsNil(o.Smallut) {
		var ret int64
		return ret
	}
	return *o.Smallut
}

// GetUtOk returns a tuple with the Ut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdateGu) GetSmallutOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallut) {
		return nil, false
	}
	return o.Smallut, true
}

// HasUt returns a boolean if a field has been set.
func (o *GridUpdateGu) HasSmallut() bool {
	if o != nil && !common.IsNil(o.Smallut) {
		return true
	}

	return false
}

// SetUt gets a reference to the given int64 and assigns it to the Ut field.
func (o *GridUpdateGu) SetSmallut(v int64) {
	o.Smallut = &v
}

func (o GridUpdateGu) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GridUpdateGu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallsi) {
		toSerialize["si"] = o.Smallsi
	}
	if !common.IsNil(o.Smallst) {
		toSerialize["st"] = o.Smallst
	}
	if !common.IsNil(o.Smallss) {
		toSerialize["ss"] = o.Smallss
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallr) {
		toSerialize["r"] = o.Smallr
	}
	if !common.IsNil(o.Smallup) {
		toSerialize["up"] = o.Smallup
	}
	if !common.IsNil(o.Smalluq) {
		toSerialize["uq"] = o.Smalluq
	}
	if !common.IsNil(o.Smalluf) {
		toSerialize["uf"] = o.Smalluf
	}
	if !common.IsNil(o.Smallmp) {
		toSerialize["mp"] = o.Smallmp
	}
	if !common.IsNil(o.Smallut) {
		toSerialize["ut"] = o.Smallut
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GridUpdateGu) UnmarshalJSON(data []byte) (err error) {
	varGridUpdateGu := _GridUpdateGu{}

	err = json.Unmarshal(data, &varGridUpdateGu)

	if err != nil {
		return err
	}

	*o = GridUpdateGu(varGridUpdateGu)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "si")
		delete(additionalProperties, "st")
		delete(additionalProperties, "ss")
		delete(additionalProperties, "s")
		delete(additionalProperties, "r")
		delete(additionalProperties, "up")
		delete(additionalProperties, "uq")
		delete(additionalProperties, "uf")
		delete(additionalProperties, "mp")
		delete(additionalProperties, "ut")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGridUpdateGu struct {
	value *GridUpdateGu
	isSet bool
}

func (v NullableGridUpdateGu) Get() *GridUpdateGu {
	return v.value
}

func (v *NullableGridUpdateGu) Set(val *GridUpdateGu) {
	v.value = val
	v.isSet = true
}

func (v NullableGridUpdateGu) IsSet() bool {
	return v.isSet
}

func (v *NullableGridUpdateGu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridUpdateGu(val *GridUpdateGu) *NullableGridUpdateGu {
	return &NullableGridUpdateGu{value: val, isSet: true}
}

func (v NullableGridUpdateGu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridUpdateGu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
