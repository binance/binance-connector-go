/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderCancelReplaceResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderCancelReplaceResponseResult{}

// OrderCancelReplaceResponseResult struct for OrderCancelReplaceResponseResult
type OrderCancelReplaceResponseResult struct {
	CancelResult         *string                                           `json:"cancelResult,omitempty"`
	NewOrderResult       *string                                           `json:"newOrderResult,omitempty"`
	CancelResponse       *OrderCancelReplaceResponseResultCancelResponse   `json:"cancelResponse,omitempty"`
	NewOrderResponse     *OrderCancelReplaceResponseResultNewOrderResponse `json:"newOrderResponse,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderCancelReplaceResponseResult OrderCancelReplaceResponseResult

// NewOrderCancelReplaceResponseResult instantiates a new OrderCancelReplaceResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCancelReplaceResponseResult() *OrderCancelReplaceResponseResult {
	this := OrderCancelReplaceResponseResult{}
	return &this
}

// NewOrderCancelReplaceResponseResultWithDefaults instantiates a new OrderCancelReplaceResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCancelReplaceResponseResultWithDefaults() *OrderCancelReplaceResponseResult {
	this := OrderCancelReplaceResponseResult{}
	return &this
}

// GetCancelResult returns the CancelResult field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResult) GetCancelResult() string {
	if o == nil || common.IsNil(o.CancelResult) {
		var ret string
		return ret
	}
	return *o.CancelResult
}

// GetCancelResultOk returns a tuple with the CancelResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResult) GetCancelResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.CancelResult) {
		return nil, false
	}
	return o.CancelResult, true
}

// HasCancelResult returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResult) HasCancelResult() bool {
	if o != nil && !common.IsNil(o.CancelResult) {
		return true
	}

	return false
}

// SetCancelResult gets a reference to the given string and assigns it to the CancelResult field.
func (o *OrderCancelReplaceResponseResult) SetCancelResult(v string) {
	o.CancelResult = &v
}

// GetNewOrderResult returns the NewOrderResult field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResult) GetNewOrderResult() string {
	if o == nil || common.IsNil(o.NewOrderResult) {
		var ret string
		return ret
	}
	return *o.NewOrderResult
}

// GetNewOrderResultOk returns a tuple with the NewOrderResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResult) GetNewOrderResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.NewOrderResult) {
		return nil, false
	}
	return o.NewOrderResult, true
}

// HasNewOrderResult returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResult) HasNewOrderResult() bool {
	if o != nil && !common.IsNil(o.NewOrderResult) {
		return true
	}

	return false
}

// SetNewOrderResult gets a reference to the given string and assigns it to the NewOrderResult field.
func (o *OrderCancelReplaceResponseResult) SetNewOrderResult(v string) {
	o.NewOrderResult = &v
}

// GetCancelResponse returns the CancelResponse field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResult) GetCancelResponse() OrderCancelReplaceResponseResultCancelResponse {
	if o == nil || common.IsNil(o.CancelResponse) {
		var ret OrderCancelReplaceResponseResultCancelResponse
		return ret
	}
	return *o.CancelResponse
}

// GetCancelResponseOk returns a tuple with the CancelResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResult) GetCancelResponseOk() (*OrderCancelReplaceResponseResultCancelResponse, bool) {
	if o == nil || common.IsNil(o.CancelResponse) {
		return nil, false
	}
	return o.CancelResponse, true
}

// HasCancelResponse returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResult) HasCancelResponse() bool {
	if o != nil && !common.IsNil(o.CancelResponse) {
		return true
	}

	return false
}

// SetCancelResponse gets a reference to the given OrderCancelReplaceResponseResultCancelResponse and assigns it to the CancelResponse field.
func (o *OrderCancelReplaceResponseResult) SetCancelResponse(v OrderCancelReplaceResponseResultCancelResponse) {
	o.CancelResponse = &v
}

// GetNewOrderResponse returns the NewOrderResponse field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResult) GetNewOrderResponse() OrderCancelReplaceResponseResultNewOrderResponse {
	if o == nil || common.IsNil(o.NewOrderResponse) {
		var ret OrderCancelReplaceResponseResultNewOrderResponse
		return ret
	}
	return *o.NewOrderResponse
}

// GetNewOrderResponseOk returns a tuple with the NewOrderResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResult) GetNewOrderResponseOk() (*OrderCancelReplaceResponseResultNewOrderResponse, bool) {
	if o == nil || common.IsNil(o.NewOrderResponse) {
		return nil, false
	}
	return o.NewOrderResponse, true
}

// HasNewOrderResponse returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResult) HasNewOrderResponse() bool {
	if o != nil && !common.IsNil(o.NewOrderResponse) {
		return true
	}

	return false
}

// SetNewOrderResponse gets a reference to the given OrderCancelReplaceResponseResultNewOrderResponse and assigns it to the NewOrderResponse field.
func (o *OrderCancelReplaceResponseResult) SetNewOrderResponse(v OrderCancelReplaceResponseResultNewOrderResponse) {
	o.NewOrderResponse = &v
}

func (o OrderCancelReplaceResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCancelReplaceResponseResult) ToMap() (map[string]interface{}, error) {
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

func (o *OrderCancelReplaceResponseResult) UnmarshalJSON(data []byte) (err error) {
	varOrderCancelReplaceResponseResult := _OrderCancelReplaceResponseResult{}

	err = json.Unmarshal(data, &varOrderCancelReplaceResponseResult)

	if err != nil {
		return err
	}

	*o = OrderCancelReplaceResponseResult(varOrderCancelReplaceResponseResult)

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

type NullableOrderCancelReplaceResponseResult struct {
	value *OrderCancelReplaceResponseResult
	isSet bool
}

func (v NullableOrderCancelReplaceResponseResult) Get() *OrderCancelReplaceResponseResult {
	return v.value
}

func (v *NullableOrderCancelReplaceResponseResult) Set(val *OrderCancelReplaceResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplaceResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplaceResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplaceResponseResult(val *OrderCancelReplaceResponseResult) *NullableOrderCancelReplaceResponseResult {
	return &NullableOrderCancelReplaceResponseResult{value: val, isSet: true}
}

func (v NullableOrderCancelReplaceResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplaceResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
