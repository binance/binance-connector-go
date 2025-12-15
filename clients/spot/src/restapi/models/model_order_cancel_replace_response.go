/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderCancelReplaceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderCancelReplaceResponse{}

// OrderCancelReplaceResponse struct for OrderCancelReplaceResponse
type OrderCancelReplaceResponse struct {
	CancelResult         *string                                     `json:"cancelResult,omitempty"`
	NewOrderResult       *string                                     `json:"newOrderResult,omitempty"`
	CancelResponse       *OrderCancelReplaceResponseCancelResponse   `json:"cancelResponse,omitempty"`
	NewOrderResponse     *OrderCancelReplaceResponseNewOrderResponse `json:"newOrderResponse,omitempty"`
	Code                 *int64                                      `json:"code,omitempty"`
	Msg                  *string                                     `json:"msg,omitempty"`
	Data                 *OrderCancelReplaceResponseData             `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderCancelReplaceResponse OrderCancelReplaceResponse

// NewOrderCancelReplaceResponse instantiates a new OrderCancelReplaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCancelReplaceResponse() *OrderCancelReplaceResponse {
	this := OrderCancelReplaceResponse{}
	return &this
}

// NewOrderCancelReplaceResponseWithDefaults instantiates a new OrderCancelReplaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCancelReplaceResponseWithDefaults() *OrderCancelReplaceResponse {
	this := OrderCancelReplaceResponse{}
	return &this
}

// GetCancelResult returns the CancelResult field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponse) GetCancelResult() string {
	if o == nil || common.IsNil(o.CancelResult) {
		var ret string
		return ret
	}
	return *o.CancelResult
}

// GetCancelResultOk returns a tuple with the CancelResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponse) GetCancelResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.CancelResult) {
		return nil, false
	}
	return o.CancelResult, true
}

// HasCancelResult returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponse) HasCancelResult() bool {
	if o != nil && !common.IsNil(o.CancelResult) {
		return true
	}

	return false
}

// SetCancelResult gets a reference to the given string and assigns it to the CancelResult field.
func (o *OrderCancelReplaceResponse) SetCancelResult(v string) {
	o.CancelResult = &v
}

// GetNewOrderResult returns the NewOrderResult field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponse) GetNewOrderResult() string {
	if o == nil || common.IsNil(o.NewOrderResult) {
		var ret string
		return ret
	}
	return *o.NewOrderResult
}

// GetNewOrderResultOk returns a tuple with the NewOrderResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponse) GetNewOrderResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.NewOrderResult) {
		return nil, false
	}
	return o.NewOrderResult, true
}

// HasNewOrderResult returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponse) HasNewOrderResult() bool {
	if o != nil && !common.IsNil(o.NewOrderResult) {
		return true
	}

	return false
}

// SetNewOrderResult gets a reference to the given string and assigns it to the NewOrderResult field.
func (o *OrderCancelReplaceResponse) SetNewOrderResult(v string) {
	o.NewOrderResult = &v
}

// GetCancelResponse returns the CancelResponse field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponse) GetCancelResponse() OrderCancelReplaceResponseCancelResponse {
	if o == nil || common.IsNil(o.CancelResponse) {
		var ret OrderCancelReplaceResponseCancelResponse
		return ret
	}
	return *o.CancelResponse
}

// GetCancelResponseOk returns a tuple with the CancelResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponse) GetCancelResponseOk() (*OrderCancelReplaceResponseCancelResponse, bool) {
	if o == nil || common.IsNil(o.CancelResponse) {
		return nil, false
	}
	return o.CancelResponse, true
}

// HasCancelResponse returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponse) HasCancelResponse() bool {
	if o != nil && !common.IsNil(o.CancelResponse) {
		return true
	}

	return false
}

// SetCancelResponse gets a reference to the given OrderCancelReplaceResponseCancelResponse and assigns it to the CancelResponse field.
func (o *OrderCancelReplaceResponse) SetCancelResponse(v OrderCancelReplaceResponseCancelResponse) {
	o.CancelResponse = &v
}

// GetNewOrderResponse returns the NewOrderResponse field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponse) GetNewOrderResponse() OrderCancelReplaceResponseNewOrderResponse {
	if o == nil || common.IsNil(o.NewOrderResponse) {
		var ret OrderCancelReplaceResponseNewOrderResponse
		return ret
	}
	return *o.NewOrderResponse
}

// GetNewOrderResponseOk returns a tuple with the NewOrderResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponse) GetNewOrderResponseOk() (*OrderCancelReplaceResponseNewOrderResponse, bool) {
	if o == nil || common.IsNil(o.NewOrderResponse) {
		return nil, false
	}
	return o.NewOrderResponse, true
}

// HasNewOrderResponse returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponse) HasNewOrderResponse() bool {
	if o != nil && !common.IsNil(o.NewOrderResponse) {
		return true
	}

	return false
}

// SetNewOrderResponse gets a reference to the given OrderCancelReplaceResponseNewOrderResponse and assigns it to the NewOrderResponse field.
func (o *OrderCancelReplaceResponse) SetNewOrderResponse(v OrderCancelReplaceResponseNewOrderResponse) {
	o.NewOrderResponse = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *OrderCancelReplaceResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *OrderCancelReplaceResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponse) GetData() OrderCancelReplaceResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret OrderCancelReplaceResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponse) GetDataOk() (*OrderCancelReplaceResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderCancelReplaceResponseData and assigns it to the Data field.
func (o *OrderCancelReplaceResponse) SetData(v OrderCancelReplaceResponseData) {
	o.Data = &v
}

func (o OrderCancelReplaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCancelReplaceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CancelResult) {
		toSerialize["cancelResult"] = o.CancelResult
	}
	if !common.IsNil(o.NewOrderResult) {
		toSerialize["newOrderResult"] = o.NewOrderResult
	}
	if !common.IsNil(o.CancelResponse) {
		toSerialize["cancelResponse"] = o.CancelResponse
	}
	if !common.IsNil(o.NewOrderResponse) {
		toSerialize["newOrderResponse"] = o.NewOrderResponse
	}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderCancelReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderCancelReplaceResponse := _OrderCancelReplaceResponse{}

	err = json.Unmarshal(data, &varOrderCancelReplaceResponse)

	if err != nil {
		return err
	}

	*o = OrderCancelReplaceResponse(varOrderCancelReplaceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cancelResult")
		delete(additionalProperties, "newOrderResult")
		delete(additionalProperties, "cancelResponse")
		delete(additionalProperties, "newOrderResponse")
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderCancelReplaceResponse struct {
	value *OrderCancelReplaceResponse
	isSet bool
}

func (v NullableOrderCancelReplaceResponse) Get() *OrderCancelReplaceResponse {
	return v.value
}

func (v *NullableOrderCancelReplaceResponse) Set(val *OrderCancelReplaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplaceResponse(val *OrderCancelReplaceResponse) *NullableOrderCancelReplaceResponse {
	return &NullableOrderCancelReplaceResponse{value: val, isSet: true}
}

func (v NullableOrderCancelReplaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
