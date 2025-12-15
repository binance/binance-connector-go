/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ContractInfoStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ContractInfoStreamResponse{}

// ContractInfoStreamResponse struct for ContractInfoStreamResponse
type ContractInfoStreamResponse struct {
	Smalle               *string                              `json:"e,omitempty"`
	E                    *int64                               `json:"E,omitempty"`
	Smalls               *string                              `json:"s,omitempty"`
	Smallps              *string                              `json:"ps,omitempty"`
	Smallct              *string                              `json:"ct,omitempty"`
	Smalldt              *int64                               `json:"dt,omitempty"`
	Smallot              *int64                               `json:"ot,omitempty"`
	Smallcs              *string                              `json:"cs,omitempty"`
	Bks                  []ContractInfoStreamResponseBksInner `json:"bks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContractInfoStreamResponse ContractInfoStreamResponse

// NewContractInfoStreamResponse instantiates a new ContractInfoStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractInfoStreamResponse() *ContractInfoStreamResponse {
	this := ContractInfoStreamResponse{}
	return &this
}

// NewContractInfoStreamResponseWithDefaults instantiates a new ContractInfoStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractInfoStreamResponseWithDefaults() *ContractInfoStreamResponse {
	this := ContractInfoStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *ContractInfoStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *ContractInfoStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *ContractInfoStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *ContractInfoStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *ContractInfoStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *ContractInfoStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *ContractInfoStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *ContractInfoStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *ContractInfoStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *ContractInfoStreamResponse) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponse) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *ContractInfoStreamResponse) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *ContractInfoStreamResponse) SetSmallps(v string) {
	o.Smallps = &v
}

// GetCt returns the Ct field value if set, zero value otherwise.
func (o *ContractInfoStreamResponse) GetSmallct() string {
	if o == nil || common.IsNil(o.Smallct) {
		var ret string
		return ret
	}
	return *o.Smallct
}

// GetCtOk returns a tuple with the Ct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponse) GetSmallctOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallct) {
		return nil, false
	}
	return o.Smallct, true
}

// HasCt returns a boolean if a field has been set.
func (o *ContractInfoStreamResponse) HasSmallct() bool {
	if o != nil && !common.IsNil(o.Smallct) {
		return true
	}

	return false
}

// SetCt gets a reference to the given string and assigns it to the Ct field.
func (o *ContractInfoStreamResponse) SetSmallct(v string) {
	o.Smallct = &v
}

// GetDt returns the Dt field value if set, zero value otherwise.
func (o *ContractInfoStreamResponse) GetSmalldt() int64 {
	if o == nil || common.IsNil(o.Smalldt) {
		var ret int64
		return ret
	}
	return *o.Smalldt
}

// GetDtOk returns a tuple with the Dt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponse) GetSmalldtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalldt) {
		return nil, false
	}
	return o.Smalldt, true
}

// HasDt returns a boolean if a field has been set.
func (o *ContractInfoStreamResponse) HasSmalldt() bool {
	if o != nil && !common.IsNil(o.Smalldt) {
		return true
	}

	return false
}

// SetDt gets a reference to the given int64 and assigns it to the Dt field.
func (o *ContractInfoStreamResponse) SetSmalldt(v int64) {
	o.Smalldt = &v
}

// GetOt returns the Ot field value if set, zero value otherwise.
func (o *ContractInfoStreamResponse) GetSmallot() int64 {
	if o == nil || common.IsNil(o.Smallot) {
		var ret int64
		return ret
	}
	return *o.Smallot
}

// GetOtOk returns a tuple with the Ot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponse) GetSmallotOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallot) {
		return nil, false
	}
	return o.Smallot, true
}

// HasOt returns a boolean if a field has been set.
func (o *ContractInfoStreamResponse) HasSmallot() bool {
	if o != nil && !common.IsNil(o.Smallot) {
		return true
	}

	return false
}

// SetOt gets a reference to the given int64 and assigns it to the Ot field.
func (o *ContractInfoStreamResponse) SetSmallot(v int64) {
	o.Smallot = &v
}

// GetCs returns the Cs field value if set, zero value otherwise.
func (o *ContractInfoStreamResponse) GetSmallcs() string {
	if o == nil || common.IsNil(o.Smallcs) {
		var ret string
		return ret
	}
	return *o.Smallcs
}

// GetCsOk returns a tuple with the Cs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponse) GetSmallcsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallcs) {
		return nil, false
	}
	return o.Smallcs, true
}

// HasCs returns a boolean if a field has been set.
func (o *ContractInfoStreamResponse) HasSmallcs() bool {
	if o != nil && !common.IsNil(o.Smallcs) {
		return true
	}

	return false
}

// SetCs gets a reference to the given string and assigns it to the Cs field.
func (o *ContractInfoStreamResponse) SetSmallcs(v string) {
	o.Smallcs = &v
}

// GetBks returns the Bks field value if set, zero value otherwise.
func (o *ContractInfoStreamResponse) GetBks() []ContractInfoStreamResponseBksInner {
	if o == nil || common.IsNil(o.Bks) {
		var ret []ContractInfoStreamResponseBksInner
		return ret
	}
	return o.Bks
}

// GetBksOk returns a tuple with the Bks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponse) GetBksOk() ([]ContractInfoStreamResponseBksInner, bool) {
	if o == nil || common.IsNil(o.Bks) {
		return nil, false
	}
	return o.Bks, true
}

// HasBks returns a boolean if a field has been set.
func (o *ContractInfoStreamResponse) HasBks() bool {
	if o != nil && !common.IsNil(o.Bks) {
		return true
	}

	return false
}

// SetBks gets a reference to the given []ContractInfoStreamResponseBksInner and assigns it to the Bks field.
func (o *ContractInfoStreamResponse) SetBks(v []ContractInfoStreamResponseBksInner) {
	o.Bks = v
}

func (o ContractInfoStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractInfoStreamResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Smallct) {
		toSerialize["ct"] = o.Smallct
	}
	if !common.IsNil(o.Smalldt) {
		toSerialize["dt"] = o.Smalldt
	}
	if !common.IsNil(o.Smallot) {
		toSerialize["ot"] = o.Smallot
	}
	if !common.IsNil(o.Smallcs) {
		toSerialize["cs"] = o.Smallcs
	}
	if !common.IsNil(o.Bks) {
		toSerialize["bks"] = o.Bks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContractInfoStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varContractInfoStreamResponse := _ContractInfoStreamResponse{}

	err = json.Unmarshal(data, &varContractInfoStreamResponse)

	if err != nil {
		return err
	}

	*o = ContractInfoStreamResponse(varContractInfoStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "ct")
		delete(additionalProperties, "dt")
		delete(additionalProperties, "ot")
		delete(additionalProperties, "cs")
		delete(additionalProperties, "bks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContractInfoStreamResponse struct {
	value *ContractInfoStreamResponse
	isSet bool
}

func (v NullableContractInfoStreamResponse) Get() *ContractInfoStreamResponse {
	return v.value
}

func (v *NullableContractInfoStreamResponse) Set(val *ContractInfoStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContractInfoStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContractInfoStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractInfoStreamResponse(val *ContractInfoStreamResponse) *NullableContractInfoStreamResponse {
	return &NullableContractInfoStreamResponse{value: val, isSet: true}
}

func (v NullableContractInfoStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractInfoStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
