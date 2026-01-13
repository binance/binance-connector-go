/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderCancelReplaceResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderCancelReplaceResponseData{}

// OrderCancelReplaceResponseData struct for OrderCancelReplaceResponseData
type OrderCancelReplaceResponseData struct {
	CancelResult         *string                                         `json:"cancelResult,omitempty"`
	NewOrderResult       *string                                         `json:"newOrderResult,omitempty"`
	CancelResponse       *OrderCancelReplaceResponseDataCancelResponse   `json:"cancelResponse,omitempty"`
	NewOrderResponse     *OrderCancelReplaceResponseDataNewOrderResponse `json:"newOrderResponse,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderCancelReplaceResponseData OrderCancelReplaceResponseData

// NewOrderCancelReplaceResponseData instantiates a new OrderCancelReplaceResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCancelReplaceResponseData() *OrderCancelReplaceResponseData {
	this := OrderCancelReplaceResponseData{}
	return &this
}

// NewOrderCancelReplaceResponseDataWithDefaults instantiates a new OrderCancelReplaceResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCancelReplaceResponseDataWithDefaults() *OrderCancelReplaceResponseData {
	this := OrderCancelReplaceResponseData{}
	return &this
}

// GetCancelResult returns the CancelResult field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseData) GetCancelResult() string {
	if o == nil || common.IsNil(o.CancelResult) {
		var ret string
		return ret
	}
	return *o.CancelResult
}

// GetCancelResultOk returns a tuple with the CancelResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseData) GetCancelResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.CancelResult) {
		return nil, false
	}
	return o.CancelResult, true
}

// HasCancelResult returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseData) HasCancelResult() bool {
	if o != nil && !common.IsNil(o.CancelResult) {
		return true
	}

	return false
}

// SetCancelResult gets a reference to the given string and assigns it to the CancelResult field.
func (o *OrderCancelReplaceResponseData) SetCancelResult(v string) {
	o.CancelResult = &v
}

// GetNewOrderResult returns the NewOrderResult field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseData) GetNewOrderResult() string {
	if o == nil || common.IsNil(o.NewOrderResult) {
		var ret string
		return ret
	}
	return *o.NewOrderResult
}

// GetNewOrderResultOk returns a tuple with the NewOrderResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseData) GetNewOrderResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.NewOrderResult) {
		return nil, false
	}
	return o.NewOrderResult, true
}

// HasNewOrderResult returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseData) HasNewOrderResult() bool {
	if o != nil && !common.IsNil(o.NewOrderResult) {
		return true
	}

	return false
}

// SetNewOrderResult gets a reference to the given string and assigns it to the NewOrderResult field.
func (o *OrderCancelReplaceResponseData) SetNewOrderResult(v string) {
	o.NewOrderResult = &v
}

// GetCancelResponse returns the CancelResponse field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseData) GetCancelResponse() OrderCancelReplaceResponseDataCancelResponse {
	if o == nil || common.IsNil(o.CancelResponse) {
		var ret OrderCancelReplaceResponseDataCancelResponse
		return ret
	}
	return *o.CancelResponse
}

// GetCancelResponseOk returns a tuple with the CancelResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseData) GetCancelResponseOk() (*OrderCancelReplaceResponseDataCancelResponse, bool) {
	if o == nil || common.IsNil(o.CancelResponse) {
		return nil, false
	}
	return o.CancelResponse, true
}

// HasCancelResponse returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseData) HasCancelResponse() bool {
	if o != nil && !common.IsNil(o.CancelResponse) {
		return true
	}

	return false
}

// SetCancelResponse gets a reference to the given OrderCancelReplaceResponseDataCancelResponse and assigns it to the CancelResponse field.
func (o *OrderCancelReplaceResponseData) SetCancelResponse(v OrderCancelReplaceResponseDataCancelResponse) {
	o.CancelResponse = &v
}

// GetNewOrderResponse returns the NewOrderResponse field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseData) GetNewOrderResponse() OrderCancelReplaceResponseDataNewOrderResponse {
	if o == nil || common.IsNil(o.NewOrderResponse) {
		var ret OrderCancelReplaceResponseDataNewOrderResponse
		return ret
	}
	return *o.NewOrderResponse
}

// GetNewOrderResponseOk returns a tuple with the NewOrderResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseData) GetNewOrderResponseOk() (*OrderCancelReplaceResponseDataNewOrderResponse, bool) {
	if o == nil || common.IsNil(o.NewOrderResponse) {
		return nil, false
	}
	return o.NewOrderResponse, true
}

// HasNewOrderResponse returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseData) HasNewOrderResponse() bool {
	if o != nil && !common.IsNil(o.NewOrderResponse) {
		return true
	}

	return false
}

// SetNewOrderResponse gets a reference to the given OrderCancelReplaceResponseDataNewOrderResponse and assigns it to the NewOrderResponse field.
func (o *OrderCancelReplaceResponseData) SetNewOrderResponse(v OrderCancelReplaceResponseDataNewOrderResponse) {
	o.NewOrderResponse = &v
}

func (o OrderCancelReplaceResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCancelReplaceResponseData) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderCancelReplaceResponseData) UnmarshalJSON(data []byte) (err error) {
	varOrderCancelReplaceResponseData := _OrderCancelReplaceResponseData{}

	err = json.Unmarshal(data, &varOrderCancelReplaceResponseData)

	if err != nil {
		return err
	}

	*o = OrderCancelReplaceResponseData(varOrderCancelReplaceResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cancelResult")
		delete(additionalProperties, "newOrderResult")
		delete(additionalProperties, "cancelResponse")
		delete(additionalProperties, "newOrderResponse")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderCancelReplaceResponseData struct {
	value *OrderCancelReplaceResponseData
	isSet bool
}

func (v NullableOrderCancelReplaceResponseData) Get() *OrderCancelReplaceResponseData {
	return v.value
}

func (v *NullableOrderCancelReplaceResponseData) Set(val *OrderCancelReplaceResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplaceResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplaceResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplaceResponseData(val *OrderCancelReplaceResponseData) *NullableOrderCancelReplaceResponseData {
	return &NullableOrderCancelReplaceResponseData{value: val, isSet: true}
}

func (v NullableOrderCancelReplaceResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplaceResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
