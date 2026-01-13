/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSummaryOfSubAccountsFuturesAccountV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSummaryOfSubAccountsFuturesAccountV2Response{}

// GetSummaryOfSubAccountsFuturesAccountV2Response struct for GetSummaryOfSubAccountsFuturesAccountV2Response
type GetSummaryOfSubAccountsFuturesAccountV2Response struct {
	FutureAccountSummaryResp   *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp   `json:"futureAccountSummaryResp,omitempty"`
	DeliveryAccountSummaryResp *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp `json:"deliveryAccountSummaryResp,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _GetSummaryOfSubAccountsFuturesAccountV2Response GetSummaryOfSubAccountsFuturesAccountV2Response

// NewGetSummaryOfSubAccountsFuturesAccountV2Response instantiates a new GetSummaryOfSubAccountsFuturesAccountV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSummaryOfSubAccountsFuturesAccountV2Response() *GetSummaryOfSubAccountsFuturesAccountV2Response {
	this := GetSummaryOfSubAccountsFuturesAccountV2Response{}
	return &this
}

// NewGetSummaryOfSubAccountsFuturesAccountV2ResponseWithDefaults instantiates a new GetSummaryOfSubAccountsFuturesAccountV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSummaryOfSubAccountsFuturesAccountV2ResponseWithDefaults() *GetSummaryOfSubAccountsFuturesAccountV2Response {
	this := GetSummaryOfSubAccountsFuturesAccountV2Response{}
	return &this
}

// GetFutureAccountSummaryResp returns the FutureAccountSummaryResp field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2Response) GetFutureAccountSummaryResp() GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp {
	if o == nil || common.IsNil(o.FutureAccountSummaryResp) {
		var ret GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp
		return ret
	}
	return *o.FutureAccountSummaryResp
}

// GetFutureAccountSummaryRespOk returns a tuple with the FutureAccountSummaryResp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2Response) GetFutureAccountSummaryRespOk() (*GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp, bool) {
	if o == nil || common.IsNil(o.FutureAccountSummaryResp) {
		return nil, false
	}
	return o.FutureAccountSummaryResp, true
}

// HasFutureAccountSummaryResp returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2Response) HasFutureAccountSummaryResp() bool {
	if o != nil && !common.IsNil(o.FutureAccountSummaryResp) {
		return true
	}

	return false
}

// SetFutureAccountSummaryResp gets a reference to the given GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp and assigns it to the FutureAccountSummaryResp field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2Response) SetFutureAccountSummaryResp(v GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) {
	o.FutureAccountSummaryResp = &v
}

// GetDeliveryAccountSummaryResp returns the DeliveryAccountSummaryResp field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2Response) GetDeliveryAccountSummaryResp() GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp {
	if o == nil || common.IsNil(o.DeliveryAccountSummaryResp) {
		var ret GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp
		return ret
	}
	return *o.DeliveryAccountSummaryResp
}

// GetDeliveryAccountSummaryRespOk returns a tuple with the DeliveryAccountSummaryResp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2Response) GetDeliveryAccountSummaryRespOk() (*GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp, bool) {
	if o == nil || common.IsNil(o.DeliveryAccountSummaryResp) {
		return nil, false
	}
	return o.DeliveryAccountSummaryResp, true
}

// HasDeliveryAccountSummaryResp returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2Response) HasDeliveryAccountSummaryResp() bool {
	if o != nil && !common.IsNil(o.DeliveryAccountSummaryResp) {
		return true
	}

	return false
}

// SetDeliveryAccountSummaryResp gets a reference to the given GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp and assigns it to the DeliveryAccountSummaryResp field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2Response) SetDeliveryAccountSummaryResp(v GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) {
	o.DeliveryAccountSummaryResp = &v
}

func (o GetSummaryOfSubAccountsFuturesAccountV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSummaryOfSubAccountsFuturesAccountV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FutureAccountSummaryResp) {
		toSerialize["futureAccountSummaryResp"] = o.FutureAccountSummaryResp
	}
	if !common.IsNil(o.DeliveryAccountSummaryResp) {
		toSerialize["deliveryAccountSummaryResp"] = o.DeliveryAccountSummaryResp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSummaryOfSubAccountsFuturesAccountV2Response) UnmarshalJSON(data []byte) (err error) {
	varGetSummaryOfSubAccountsFuturesAccountV2Response := _GetSummaryOfSubAccountsFuturesAccountV2Response{}

	err = json.Unmarshal(data, &varGetSummaryOfSubAccountsFuturesAccountV2Response)

	if err != nil {
		return err
	}

	*o = GetSummaryOfSubAccountsFuturesAccountV2Response(varGetSummaryOfSubAccountsFuturesAccountV2Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "futureAccountSummaryResp")
		delete(additionalProperties, "deliveryAccountSummaryResp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSummaryOfSubAccountsFuturesAccountV2Response struct {
	value *GetSummaryOfSubAccountsFuturesAccountV2Response
	isSet bool
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2Response) Get() *GetSummaryOfSubAccountsFuturesAccountV2Response {
	return v.value
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2Response) Set(val *GetSummaryOfSubAccountsFuturesAccountV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSummaryOfSubAccountsFuturesAccountV2Response(val *GetSummaryOfSubAccountsFuturesAccountV2Response) *NullableGetSummaryOfSubAccountsFuturesAccountV2Response {
	return &NullableGetSummaryOfSubAccountsFuturesAccountV2Response{value: val, isSet: true}
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
