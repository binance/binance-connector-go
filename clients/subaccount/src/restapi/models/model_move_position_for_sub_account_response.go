/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MovePositionForSubAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MovePositionForSubAccountResponse{}

// MovePositionForSubAccountResponse struct for MovePositionForSubAccountResponse
type MovePositionForSubAccountResponse struct {
	MovePositionOrders   []MovePositionForSubAccountResponseMovePositionOrdersInner `json:"movePositionOrders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MovePositionForSubAccountResponse MovePositionForSubAccountResponse

// NewMovePositionForSubAccountResponse instantiates a new MovePositionForSubAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovePositionForSubAccountResponse() *MovePositionForSubAccountResponse {
	this := MovePositionForSubAccountResponse{}
	return &this
}

// NewMovePositionForSubAccountResponseWithDefaults instantiates a new MovePositionForSubAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovePositionForSubAccountResponseWithDefaults() *MovePositionForSubAccountResponse {
	this := MovePositionForSubAccountResponse{}
	return &this
}

// GetMovePositionOrders returns the MovePositionOrders field value if set, zero value otherwise.
func (o *MovePositionForSubAccountResponse) GetMovePositionOrders() []MovePositionForSubAccountResponseMovePositionOrdersInner {
	if o == nil || common.IsNil(o.MovePositionOrders) {
		var ret []MovePositionForSubAccountResponseMovePositionOrdersInner
		return ret
	}
	return o.MovePositionOrders
}

// GetMovePositionOrdersOk returns a tuple with the MovePositionOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountResponse) GetMovePositionOrdersOk() ([]MovePositionForSubAccountResponseMovePositionOrdersInner, bool) {
	if o == nil || common.IsNil(o.MovePositionOrders) {
		return nil, false
	}
	return o.MovePositionOrders, true
}

// HasMovePositionOrders returns a boolean if a field has been set.
func (o *MovePositionForSubAccountResponse) HasMovePositionOrders() bool {
	if o != nil && !common.IsNil(o.MovePositionOrders) {
		return true
	}

	return false
}

// SetMovePositionOrders gets a reference to the given []MovePositionForSubAccountResponseMovePositionOrdersInner and assigns it to the MovePositionOrders field.
func (o *MovePositionForSubAccountResponse) SetMovePositionOrders(v []MovePositionForSubAccountResponseMovePositionOrdersInner) {
	o.MovePositionOrders = v
}

func (o MovePositionForSubAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MovePositionForSubAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MovePositionOrders) {
		toSerialize["movePositionOrders"] = o.MovePositionOrders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MovePositionForSubAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varMovePositionForSubAccountResponse := _MovePositionForSubAccountResponse{}

	err = json.Unmarshal(data, &varMovePositionForSubAccountResponse)

	if err != nil {
		return err
	}

	*o = MovePositionForSubAccountResponse(varMovePositionForSubAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "movePositionOrders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMovePositionForSubAccountResponse struct {
	value *MovePositionForSubAccountResponse
	isSet bool
}

func (v NullableMovePositionForSubAccountResponse) Get() *MovePositionForSubAccountResponse {
	return v.value
}

func (v *NullableMovePositionForSubAccountResponse) Set(val *MovePositionForSubAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMovePositionForSubAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMovePositionForSubAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovePositionForSubAccountResponse(val *MovePositionForSubAccountResponse) *NullableMovePositionForSubAccountResponse {
	return &NullableMovePositionForSubAccountResponse{value: val, isSet: true}
}

func (v NullableMovePositionForSubAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovePositionForSubAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
