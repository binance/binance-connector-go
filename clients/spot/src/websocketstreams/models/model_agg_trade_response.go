/*
Spot WebSocket Market Streams

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AggTradeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AggTradeResponse{}

// AggTradeResponse struct for AggTradeResponse
type AggTradeResponse struct {
	// Event type
	Smalle *string `json:"e,omitempty"`
	// Event time
	E *int64 `json:"E,omitempty"`
	// Symbol
	Smalls *string `json:"s,omitempty"`
	// Aggregate trade ID
	Smalla *int64 `json:"a,omitempty"`
	// Price
	Smallp *string `json:"p,omitempty"`
	// Quantity
	Smallq *string `json:"q,omitempty"`
	// First trade ID
	Smallf *int64 `json:"f,omitempty"`
	// Last trade ID
	Smalll *int64 `json:"l,omitempty"`
	// Trade time
	T *int64 `json:"T,omitempty"`
	// Is the buyer the market maker?
	Smallm *bool `json:"m,omitempty"`
	// Ignore
	M                    *bool `json:"M,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AggTradeResponse AggTradeResponse

// NewAggTradeResponse instantiates a new AggTradeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggTradeResponse() *AggTradeResponse {
	this := AggTradeResponse{}
	return &this
}

// NewAggTradeResponseWithDefaults instantiates a new AggTradeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggTradeResponseWithDefaults() *AggTradeResponse {
	this := AggTradeResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AggTradeResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradeResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *AggTradeResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AggTradeResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AggTradeResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradeResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AggTradeResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AggTradeResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AggTradeResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradeResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AggTradeResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AggTradeResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *AggTradeResponse) GetSmalla() int64 {
	if o == nil || common.IsNil(o.Smalla) {
		var ret int64
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradeResponse) GetSmallaOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *AggTradeResponse) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given int64 and assigns it to the A field.
func (o *AggTradeResponse) SetSmalla(v int64) {
	o.Smalla = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AggTradeResponse) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradeResponse) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *AggTradeResponse) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AggTradeResponse) SetSmallp(v string) {
	o.Smallp = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AggTradeResponse) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradeResponse) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *AggTradeResponse) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AggTradeResponse) SetSmallq(v string) {
	o.Smallq = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *AggTradeResponse) GetSmallf() int64 {
	if o == nil || common.IsNil(o.Smallf) {
		var ret int64
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradeResponse) GetSmallfOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *AggTradeResponse) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given int64 and assigns it to the F field.
func (o *AggTradeResponse) SetSmallf(v int64) {
	o.Smallf = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AggTradeResponse) GetSmalll() int64 {
	if o == nil || common.IsNil(o.Smalll) {
		var ret int64
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradeResponse) GetSmalllOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *AggTradeResponse) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *AggTradeResponse) SetSmalll(v int64) {
	o.Smalll = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AggTradeResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradeResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AggTradeResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AggTradeResponse) SetT(v int64) {
	o.T = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *AggTradeResponse) GetSmallm() bool {
	if o == nil || common.IsNil(o.Smallm) {
		var ret bool
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradeResponse) GetSmallmOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *AggTradeResponse) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *AggTradeResponse) SetSmallm(v bool) {
	o.Smallm = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *AggTradeResponse) GetM() bool {
	if o == nil || common.IsNil(o.M) {
		var ret bool
		return ret
	}
	return *o.M
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTradeResponse) GetMOk() (*bool, bool) {
	if o == nil || common.IsNil(o.M) {
		return nil, false
	}
	return o.M, true
}

// HasM returns a boolean if a field has been set.
func (o *AggTradeResponse) HasM() bool {
	if o != nil && !common.IsNil(o.M) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *AggTradeResponse) SetM(v bool) {
	o.M = &v
}

func (o AggTradeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggTradeResponse) ToMap() (map[string]interface{}, error) {
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

func (o *AggTradeResponse) UnmarshalJSON(data []byte) (err error) {
	varAggTradeResponse := _AggTradeResponse{}

	err = json.Unmarshal(data, &varAggTradeResponse)

	if err != nil {
		return err
	}

	*o = AggTradeResponse(varAggTradeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
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

type NullableAggTradeResponse struct {
	value *AggTradeResponse
	isSet bool
}

func (v NullableAggTradeResponse) Get() *AggTradeResponse {
	return v.value
}

func (v *NullableAggTradeResponse) Set(val *AggTradeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAggTradeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAggTradeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggTradeResponse(val *AggTradeResponse) *NullableAggTradeResponse {
	return &NullableAggTradeResponse{value: val, isSet: true}
}

func (v NullableAggTradeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggTradeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
