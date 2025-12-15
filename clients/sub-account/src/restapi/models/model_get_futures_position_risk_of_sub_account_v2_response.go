/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFuturesPositionRiskOfSubAccountV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFuturesPositionRiskOfSubAccountV2Response{}

// GetFuturesPositionRiskOfSubAccountV2Response struct for GetFuturesPositionRiskOfSubAccountV2Response
type GetFuturesPositionRiskOfSubAccountV2Response struct {
	FuturePositionRiskVos   []GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner   `json:"futurePositionRiskVos,omitempty"`
	DeliveryPositionRiskVos []GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner `json:"deliveryPositionRiskVos,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _GetFuturesPositionRiskOfSubAccountV2Response GetFuturesPositionRiskOfSubAccountV2Response

// NewGetFuturesPositionRiskOfSubAccountV2Response instantiates a new GetFuturesPositionRiskOfSubAccountV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesPositionRiskOfSubAccountV2Response() *GetFuturesPositionRiskOfSubAccountV2Response {
	this := GetFuturesPositionRiskOfSubAccountV2Response{}
	return &this
}

// NewGetFuturesPositionRiskOfSubAccountV2ResponseWithDefaults instantiates a new GetFuturesPositionRiskOfSubAccountV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesPositionRiskOfSubAccountV2ResponseWithDefaults() *GetFuturesPositionRiskOfSubAccountV2Response {
	this := GetFuturesPositionRiskOfSubAccountV2Response{}
	return &this
}

// GetFuturePositionRiskVos returns the FuturePositionRiskVos field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2Response) GetFuturePositionRiskVos() []GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner {
	if o == nil || common.IsNil(o.FuturePositionRiskVos) {
		var ret []GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner
		return ret
	}
	return o.FuturePositionRiskVos
}

// GetFuturePositionRiskVosOk returns a tuple with the FuturePositionRiskVos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2Response) GetFuturePositionRiskVosOk() ([]GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner, bool) {
	if o == nil || common.IsNil(o.FuturePositionRiskVos) {
		return nil, false
	}
	return o.FuturePositionRiskVos, true
}

// HasFuturePositionRiskVos returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2Response) HasFuturePositionRiskVos() bool {
	if o != nil && !common.IsNil(o.FuturePositionRiskVos) {
		return true
	}

	return false
}

// SetFuturePositionRiskVos gets a reference to the given []GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner and assigns it to the FuturePositionRiskVos field.
func (o *GetFuturesPositionRiskOfSubAccountV2Response) SetFuturePositionRiskVos(v []GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) {
	o.FuturePositionRiskVos = v
}

// GetDeliveryPositionRiskVos returns the DeliveryPositionRiskVos field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2Response) GetDeliveryPositionRiskVos() []GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner {
	if o == nil || common.IsNil(o.DeliveryPositionRiskVos) {
		var ret []GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner
		return ret
	}
	return o.DeliveryPositionRiskVos
}

// GetDeliveryPositionRiskVosOk returns a tuple with the DeliveryPositionRiskVos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2Response) GetDeliveryPositionRiskVosOk() ([]GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner, bool) {
	if o == nil || common.IsNil(o.DeliveryPositionRiskVos) {
		return nil, false
	}
	return o.DeliveryPositionRiskVos, true
}

// HasDeliveryPositionRiskVos returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2Response) HasDeliveryPositionRiskVos() bool {
	if o != nil && !common.IsNil(o.DeliveryPositionRiskVos) {
		return true
	}

	return false
}

// SetDeliveryPositionRiskVos gets a reference to the given []GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner and assigns it to the DeliveryPositionRiskVos field.
func (o *GetFuturesPositionRiskOfSubAccountV2Response) SetDeliveryPositionRiskVos(v []GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) {
	o.DeliveryPositionRiskVos = v
}

func (o GetFuturesPositionRiskOfSubAccountV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesPositionRiskOfSubAccountV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FuturePositionRiskVos) {
		toSerialize["futurePositionRiskVos"] = o.FuturePositionRiskVos
	}
	if !common.IsNil(o.DeliveryPositionRiskVos) {
		toSerialize["deliveryPositionRiskVos"] = o.DeliveryPositionRiskVos
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFuturesPositionRiskOfSubAccountV2Response) UnmarshalJSON(data []byte) (err error) {
	varGetFuturesPositionRiskOfSubAccountV2Response := _GetFuturesPositionRiskOfSubAccountV2Response{}

	err = json.Unmarshal(data, &varGetFuturesPositionRiskOfSubAccountV2Response)

	if err != nil {
		return err
	}

	*o = GetFuturesPositionRiskOfSubAccountV2Response(varGetFuturesPositionRiskOfSubAccountV2Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "futurePositionRiskVos")
		delete(additionalProperties, "deliveryPositionRiskVos")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFuturesPositionRiskOfSubAccountV2Response struct {
	value *GetFuturesPositionRiskOfSubAccountV2Response
	isSet bool
}

func (v NullableGetFuturesPositionRiskOfSubAccountV2Response) Get() *GetFuturesPositionRiskOfSubAccountV2Response {
	return v.value
}

func (v *NullableGetFuturesPositionRiskOfSubAccountV2Response) Set(val *GetFuturesPositionRiskOfSubAccountV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesPositionRiskOfSubAccountV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesPositionRiskOfSubAccountV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesPositionRiskOfSubAccountV2Response(val *GetFuturesPositionRiskOfSubAccountV2Response) *NullableGetFuturesPositionRiskOfSubAccountV2Response {
	return &NullableGetFuturesPositionRiskOfSubAccountV2Response{value: val, isSet: true}
}

func (v NullableGetFuturesPositionRiskOfSubAccountV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesPositionRiskOfSubAccountV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
