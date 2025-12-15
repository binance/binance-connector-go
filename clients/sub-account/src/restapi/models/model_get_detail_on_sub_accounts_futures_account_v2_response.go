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

// checks if the GetDetailOnSubAccountsFuturesAccountV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDetailOnSubAccountsFuturesAccountV2Response{}

// GetDetailOnSubAccountsFuturesAccountV2Response struct for GetDetailOnSubAccountsFuturesAccountV2Response
type GetDetailOnSubAccountsFuturesAccountV2Response struct {
	FutureAccountResp    *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp   `json:"futureAccountResp,omitempty"`
	DeliveryAccountResp  *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp `json:"deliveryAccountResp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDetailOnSubAccountsFuturesAccountV2Response GetDetailOnSubAccountsFuturesAccountV2Response

// NewGetDetailOnSubAccountsFuturesAccountV2Response instantiates a new GetDetailOnSubAccountsFuturesAccountV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDetailOnSubAccountsFuturesAccountV2Response() *GetDetailOnSubAccountsFuturesAccountV2Response {
	this := GetDetailOnSubAccountsFuturesAccountV2Response{}
	return &this
}

// NewGetDetailOnSubAccountsFuturesAccountV2ResponseWithDefaults instantiates a new GetDetailOnSubAccountsFuturesAccountV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDetailOnSubAccountsFuturesAccountV2ResponseWithDefaults() *GetDetailOnSubAccountsFuturesAccountV2Response {
	this := GetDetailOnSubAccountsFuturesAccountV2Response{}
	return &this
}

// GetFutureAccountResp returns the FutureAccountResp field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2Response) GetFutureAccountResp() GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp {
	if o == nil || common.IsNil(o.FutureAccountResp) {
		var ret GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp
		return ret
	}
	return *o.FutureAccountResp
}

// GetFutureAccountRespOk returns a tuple with the FutureAccountResp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2Response) GetFutureAccountRespOk() (*GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp, bool) {
	if o == nil || common.IsNil(o.FutureAccountResp) {
		return nil, false
	}
	return o.FutureAccountResp, true
}

// HasFutureAccountResp returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2Response) HasFutureAccountResp() bool {
	if o != nil && !common.IsNil(o.FutureAccountResp) {
		return true
	}

	return false
}

// SetFutureAccountResp gets a reference to the given GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp and assigns it to the FutureAccountResp field.
func (o *GetDetailOnSubAccountsFuturesAccountV2Response) SetFutureAccountResp(v GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) {
	o.FutureAccountResp = &v
}

// GetDeliveryAccountResp returns the DeliveryAccountResp field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2Response) GetDeliveryAccountResp() GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp {
	if o == nil || common.IsNil(o.DeliveryAccountResp) {
		var ret GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp
		return ret
	}
	return *o.DeliveryAccountResp
}

// GetDeliveryAccountRespOk returns a tuple with the DeliveryAccountResp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2Response) GetDeliveryAccountRespOk() (*GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp, bool) {
	if o == nil || common.IsNil(o.DeliveryAccountResp) {
		return nil, false
	}
	return o.DeliveryAccountResp, true
}

// HasDeliveryAccountResp returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2Response) HasDeliveryAccountResp() bool {
	if o != nil && !common.IsNil(o.DeliveryAccountResp) {
		return true
	}

	return false
}

// SetDeliveryAccountResp gets a reference to the given GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp and assigns it to the DeliveryAccountResp field.
func (o *GetDetailOnSubAccountsFuturesAccountV2Response) SetDeliveryAccountResp(v GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) {
	o.DeliveryAccountResp = &v
}

func (o GetDetailOnSubAccountsFuturesAccountV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDetailOnSubAccountsFuturesAccountV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FutureAccountResp) {
		toSerialize["futureAccountResp"] = o.FutureAccountResp
	}
	if !common.IsNil(o.DeliveryAccountResp) {
		toSerialize["deliveryAccountResp"] = o.DeliveryAccountResp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDetailOnSubAccountsFuturesAccountV2Response) UnmarshalJSON(data []byte) (err error) {
	varGetDetailOnSubAccountsFuturesAccountV2Response := _GetDetailOnSubAccountsFuturesAccountV2Response{}

	err = json.Unmarshal(data, &varGetDetailOnSubAccountsFuturesAccountV2Response)

	if err != nil {
		return err
	}

	*o = GetDetailOnSubAccountsFuturesAccountV2Response(varGetDetailOnSubAccountsFuturesAccountV2Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "futureAccountResp")
		delete(additionalProperties, "deliveryAccountResp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDetailOnSubAccountsFuturesAccountV2Response struct {
	value *GetDetailOnSubAccountsFuturesAccountV2Response
	isSet bool
}

func (v NullableGetDetailOnSubAccountsFuturesAccountV2Response) Get() *GetDetailOnSubAccountsFuturesAccountV2Response {
	return v.value
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountV2Response) Set(val *GetDetailOnSubAccountsFuturesAccountV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDetailOnSubAccountsFuturesAccountV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDetailOnSubAccountsFuturesAccountV2Response(val *GetDetailOnSubAccountsFuturesAccountV2Response) *NullableGetDetailOnSubAccountsFuturesAccountV2Response {
	return &NullableGetDetailOnSubAccountsFuturesAccountV2Response{value: val, isSet: true}
}

func (v NullableGetDetailOnSubAccountsFuturesAccountV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
