/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AllBookTickerStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllBookTickerStreamResponse{}

// AllBookTickerStreamResponse struct for AllBookTickerStreamResponse
type AllBookTickerStreamResponse struct {
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

type _AllBookTickerStreamResponse AllBookTickerStreamResponse

// NewAllBookTickerStreamResponse instantiates a new AllBookTickerStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllBookTickerStreamResponse() *AllBookTickerStreamResponse {
	this := AllBookTickerStreamResponse{}
	return &this
}

// NewAllBookTickerStreamResponseWithDefaults instantiates a new AllBookTickerStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllBookTickerStreamResponseWithDefaults() *AllBookTickerStreamResponse {
	this := AllBookTickerStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllBookTickerStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllBookTickerStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *AllBookTickerStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AllBookTickerStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllBookTickerStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllBookTickerStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AllBookTickerStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AllBookTickerStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AllBookTickerStreamResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllBookTickerStreamResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AllBookTickerStreamResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AllBookTickerStreamResponse) SetT(v int64) {
	o.T = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *AllBookTickerStreamResponse) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllBookTickerStreamResponse) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *AllBookTickerStreamResponse) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *AllBookTickerStreamResponse) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AllBookTickerStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllBookTickerStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AllBookTickerStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AllBookTickerStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *AllBookTickerStreamResponse) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllBookTickerStreamResponse) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *AllBookTickerStreamResponse) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *AllBookTickerStreamResponse) SetSmallb(v string) {
	o.Smallb = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *AllBookTickerStreamResponse) GetB() string {
	if o == nil || common.IsNil(o.B) {
		var ret string
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllBookTickerStreamResponse) GetBOk() (*string, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *AllBookTickerStreamResponse) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *AllBookTickerStreamResponse) SetB(v string) {
	o.B = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *AllBookTickerStreamResponse) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllBookTickerStreamResponse) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *AllBookTickerStreamResponse) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *AllBookTickerStreamResponse) SetSmalla(v string) {
	o.Smalla = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *AllBookTickerStreamResponse) GetA() string {
	if o == nil || common.IsNil(o.A) {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllBookTickerStreamResponse) GetAOk() (*string, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *AllBookTickerStreamResponse) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *AllBookTickerStreamResponse) SetA(v string) {
	o.A = &v
}

func (o AllBookTickerStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllBookTickerStreamResponse) ToMap() (map[string]interface{}, error) {
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

func (o *AllBookTickerStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varAllBookTickerStreamResponse := _AllBookTickerStreamResponse{}

	err = json.Unmarshal(data, &varAllBookTickerStreamResponse)

	if err != nil {
		return err
	}

	*o = AllBookTickerStreamResponse(varAllBookTickerStreamResponse)

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

type NullableAllBookTickerStreamResponse struct {
	value *AllBookTickerStreamResponse
	isSet bool
}

func (v NullableAllBookTickerStreamResponse) Get() *AllBookTickerStreamResponse {
	return v.value
}

func (v *NullableAllBookTickerStreamResponse) Set(val *AllBookTickerStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllBookTickerStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllBookTickerStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllBookTickerStreamResponse(val *AllBookTickerStreamResponse) *NullableAllBookTickerStreamResponse {
	return &NullableAllBookTickerStreamResponse{value: val, isSet: true}
}

func (v NullableAllBookTickerStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllBookTickerStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
