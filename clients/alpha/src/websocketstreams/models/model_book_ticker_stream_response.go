/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BookTickerStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BookTickerStreamResponse{}

// BookTickerStreamResponse struct for BookTickerStreamResponse
type BookTickerStreamResponse struct {
	// eventType
	Smalle *string `json:"e,omitempty"`
	// eventTime
	E *int64 `json:"E,omitempty"`
	// transactionTime
	T *int64 `json:"T,omitempty"`
	// updateId
	Smallu *int64 `json:"u,omitempty"`
	// symbol
	Smalls *string `json:"s,omitempty"`
	// bid1Price
	Smallb *string `json:"b,omitempty"`
	// bid1Quantity
	B *string `json:"B,omitempty"`
	// ask1Price
	Smalla *string `json:"a,omitempty"`
	// ask1Quantity
	A                    *string `json:"A,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BookTickerStreamResponse BookTickerStreamResponse

// NewBookTickerStreamResponse instantiates a new BookTickerStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookTickerStreamResponse() *BookTickerStreamResponse {
	this := BookTickerStreamResponse{}
	return &this
}

// NewBookTickerStreamResponseWithDefaults instantiates a new BookTickerStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookTickerStreamResponseWithDefaults() *BookTickerStreamResponse {
	this := BookTickerStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *BookTickerStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *BookTickerStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *BookTickerStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *BookTickerStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *BookTickerStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *BookTickerStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *BookTickerStreamResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerStreamResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *BookTickerStreamResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *BookTickerStreamResponse) SetT(v int64) {
	o.T = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *BookTickerStreamResponse) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerStreamResponse) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *BookTickerStreamResponse) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *BookTickerStreamResponse) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *BookTickerStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *BookTickerStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *BookTickerStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *BookTickerStreamResponse) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerStreamResponse) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *BookTickerStreamResponse) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *BookTickerStreamResponse) SetSmallb(v string) {
	o.Smallb = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *BookTickerStreamResponse) GetB() string {
	if o == nil || common.IsNil(o.B) {
		var ret string
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerStreamResponse) GetBOk() (*string, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *BookTickerStreamResponse) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *BookTickerStreamResponse) SetB(v string) {
	o.B = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *BookTickerStreamResponse) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerStreamResponse) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *BookTickerStreamResponse) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *BookTickerStreamResponse) SetSmalla(v string) {
	o.Smalla = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *BookTickerStreamResponse) GetA() string {
	if o == nil || common.IsNil(o.A) {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookTickerStreamResponse) GetAOk() (*string, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *BookTickerStreamResponse) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *BookTickerStreamResponse) SetA(v string) {
	o.A = &v
}

func (o BookTickerStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookTickerStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
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

func (o *BookTickerStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varBookTickerStreamResponse := _BookTickerStreamResponse{}

	err = json.Unmarshal(data, &varBookTickerStreamResponse)

	if err != nil {
		return err
	}

	*o = BookTickerStreamResponse(varBookTickerStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
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

type NullableBookTickerStreamResponse struct {
	value *BookTickerStreamResponse
	isSet bool
}

func (v NullableBookTickerStreamResponse) Get() *BookTickerStreamResponse {
	return v.value
}

func (v *NullableBookTickerStreamResponse) Set(val *BookTickerStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBookTickerStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBookTickerStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookTickerStreamResponse(val *BookTickerStreamResponse) *NullableBookTickerStreamResponse {
	return &NullableBookTickerStreamResponse{value: val, isSet: true}
}

func (v NullableBookTickerStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookTickerStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
