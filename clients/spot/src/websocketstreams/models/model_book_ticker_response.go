/*
Spot WebSocket Market Streams

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BookTickerResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BookTickerResponse{}

// BookTickerResponse struct for BookTickerResponse
type BookTickerResponse struct {
	// Order book updateId
	Smallu *int64 `json:"u,omitempty"`
	// Symbol
	Smalls *string `json:"s,omitempty"`
	// Best bid price
	Smallb *string `json:"b,omitempty"`
	// Best bid qty
	B *string `json:"B,omitempty"`
	// Best ask price
	Smalla *string `json:"a,omitempty"`
	// Best ask qty
	A                    *string `json:"A,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BookTickerResponse BookTickerResponse

// NewBookTickerResponse instantiates a new BookTickerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookTickerResponse() *BookTickerResponse {
	this := BookTickerResponse{}
	return &this
}

// NewBookTickerResponseWithDefaults instantiates a new BookTickerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookTickerResponseWithDefaults() *BookTickerResponse {
	this := BookTickerResponse{}
	return &this
}

// GetU returns the U field value if set, zero value otherwise.
func (o *BookTickerResponse) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerResponse) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *BookTickerResponse) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *BookTickerResponse) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *BookTickerResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *BookTickerResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *BookTickerResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *BookTickerResponse) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerResponse) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *BookTickerResponse) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *BookTickerResponse) SetSmallb(v string) {
	o.Smallb = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *BookTickerResponse) GetB() string {
	if o == nil || common.IsNil(o.B) {
		var ret string
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerResponse) GetBOk() (*string, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *BookTickerResponse) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *BookTickerResponse) SetB(v string) {
	o.B = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *BookTickerResponse) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerResponse) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *BookTickerResponse) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *BookTickerResponse) SetSmalla(v string) {
	o.Smalla = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *BookTickerResponse) GetA() string {
	if o == nil || common.IsNil(o.A) {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerResponse) GetAOk() (*string, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *BookTickerResponse) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *BookTickerResponse) SetA(v string) {
	o.A = &v
}

func (o BookTickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookTickerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallb) {
		toSerialize["b"] = o.Smallb
	}
	if !common.IsNil(o.B) {
		toSerialize["B"] = o.B
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.A) {
		toSerialize["A"] = o.A
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BookTickerResponse) UnmarshalJSON(data []byte) (err error) {
	varBookTickerResponse := _BookTickerResponse{}

	err = json.Unmarshal(data, &varBookTickerResponse)

	if err != nil {
		return err
	}

	*o = BookTickerResponse(varBookTickerResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "u")
		delete(additionalProperties, "s")
		delete(additionalProperties, "b")
		delete(additionalProperties, "B")
		delete(additionalProperties, "a")
		delete(additionalProperties, "A")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBookTickerResponse struct {
	value *BookTickerResponse
	isSet bool
}

func (v NullableBookTickerResponse) Get() *BookTickerResponse {
	return v.value
}

func (v *NullableBookTickerResponse) Set(val *BookTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBookTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBookTickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookTickerResponse(val *BookTickerResponse) *NullableBookTickerResponse {
	return &NullableBookTickerResponse{value: val, isSet: true}
}

func (v NullableBookTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
