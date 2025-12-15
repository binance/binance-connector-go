/*
Binance Algo REST API

OpenAPI Specification for the Binance Algo REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QuerySubOrdersFutureAlgoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubOrdersFutureAlgoResponse{}

// QuerySubOrdersFutureAlgoResponse struct for QuerySubOrdersFutureAlgoResponse
type QuerySubOrdersFutureAlgoResponse struct {
	Total                *int64                                           `json:"total,omitempty"`
	ExecutedQty          *string                                          `json:"executedQty,omitempty"`
	ExecutedAmt          *string                                          `json:"executedAmt,omitempty"`
	SubOrders            []QuerySubOrdersFutureAlgoResponseSubOrdersInner `json:"subOrders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubOrdersFutureAlgoResponse QuerySubOrdersFutureAlgoResponse

// NewQuerySubOrdersFutureAlgoResponse instantiates a new QuerySubOrdersFutureAlgoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubOrdersFutureAlgoResponse() *QuerySubOrdersFutureAlgoResponse {
	this := QuerySubOrdersFutureAlgoResponse{}
	return &this
}

// NewQuerySubOrdersFutureAlgoResponseWithDefaults instantiates a new QuerySubOrdersFutureAlgoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubOrdersFutureAlgoResponseWithDefaults() *QuerySubOrdersFutureAlgoResponse {
	this := QuerySubOrdersFutureAlgoResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QuerySubOrdersFutureAlgoResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponse) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponse) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponse) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *QuerySubOrdersFutureAlgoResponse) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetExecutedAmt returns the ExecutedAmt field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponse) GetExecutedAmt() string {
	if o == nil || common.IsNil(o.ExecutedAmt) {
		var ret string
		return ret
	}
	return *o.ExecutedAmt
}

// GetExecutedAmtOk returns a tuple with the ExecutedAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponse) GetExecutedAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedAmt) {
		return nil, false
	}
	return o.ExecutedAmt, true
}

// HasExecutedAmt returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponse) HasExecutedAmt() bool {
	if o != nil && !common.IsNil(o.ExecutedAmt) {
		return true
	}

	return false
}

// SetExecutedAmt gets a reference to the given string and assigns it to the ExecutedAmt field.
func (o *QuerySubOrdersFutureAlgoResponse) SetExecutedAmt(v string) {
	o.ExecutedAmt = &v
}

// GetSubOrders returns the SubOrders field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponse) GetSubOrders() []QuerySubOrdersFutureAlgoResponseSubOrdersInner {
	if o == nil || common.IsNil(o.SubOrders) {
		var ret []QuerySubOrdersFutureAlgoResponseSubOrdersInner
		return ret
	}
	return o.SubOrders
}

// GetSubOrdersOk returns a tuple with the SubOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponse) GetSubOrdersOk() ([]QuerySubOrdersFutureAlgoResponseSubOrdersInner, bool) {
	if o == nil || common.IsNil(o.SubOrders) {
		return nil, false
	}
	return o.SubOrders, true
}

// HasSubOrders returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponse) HasSubOrders() bool {
	if o != nil && !common.IsNil(o.SubOrders) {
		return true
	}

	return false
}

// SetSubOrders gets a reference to the given []QuerySubOrdersFutureAlgoResponseSubOrdersInner and assigns it to the SubOrders field.
func (o *QuerySubOrdersFutureAlgoResponse) SetSubOrders(v []QuerySubOrdersFutureAlgoResponseSubOrdersInner) {
	o.SubOrders = v
}

func (o QuerySubOrdersFutureAlgoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubOrdersFutureAlgoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.ExecutedAmt) {
		toSerialize["executedAmt"] = o.ExecutedAmt
	}
	if !common.IsNil(o.SubOrders) {
		toSerialize["subOrders"] = o.SubOrders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubOrdersFutureAlgoResponse) UnmarshalJSON(data []byte) (err error) {
	varQuerySubOrdersFutureAlgoResponse := _QuerySubOrdersFutureAlgoResponse{}

	err = json.Unmarshal(data, &varQuerySubOrdersFutureAlgoResponse)

	if err != nil {
		return err
	}

	*o = QuerySubOrdersFutureAlgoResponse(varQuerySubOrdersFutureAlgoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "executedAmt")
		delete(additionalProperties, "subOrders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubOrdersFutureAlgoResponse struct {
	value *QuerySubOrdersFutureAlgoResponse
	isSet bool
}

func (v NullableQuerySubOrdersFutureAlgoResponse) Get() *QuerySubOrdersFutureAlgoResponse {
	return v.value
}

func (v *NullableQuerySubOrdersFutureAlgoResponse) Set(val *QuerySubOrdersFutureAlgoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubOrdersFutureAlgoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubOrdersFutureAlgoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubOrdersFutureAlgoResponse(val *QuerySubOrdersFutureAlgoResponse) *NullableQuerySubOrdersFutureAlgoResponse {
	return &NullableQuerySubOrdersFutureAlgoResponse{value: val, isSet: true}
}

func (v NullableQuerySubOrdersFutureAlgoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubOrdersFutureAlgoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
