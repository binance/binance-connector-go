/*
Binance Algo REST API

OpenAPI Specification for the Binance Algo REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySubOrdersSpotAlgoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubOrdersSpotAlgoResponse{}

// QuerySubOrdersSpotAlgoResponse struct for QuerySubOrdersSpotAlgoResponse
type QuerySubOrdersSpotAlgoResponse struct {
	Total                *int64                                           `json:"total,omitempty"`
	ExecutedQty          *string                                          `json:"executedQty,omitempty"`
	ExecutedAmt          *string                                          `json:"executedAmt,omitempty"`
	SubOrders            []QuerySubOrdersFutureAlgoResponseSubOrdersInner `json:"subOrders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubOrdersSpotAlgoResponse QuerySubOrdersSpotAlgoResponse

// NewQuerySubOrdersSpotAlgoResponse instantiates a new QuerySubOrdersSpotAlgoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubOrdersSpotAlgoResponse() *QuerySubOrdersSpotAlgoResponse {
	this := QuerySubOrdersSpotAlgoResponse{}
	return &this
}

// NewQuerySubOrdersSpotAlgoResponseWithDefaults instantiates a new QuerySubOrdersSpotAlgoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubOrdersSpotAlgoResponseWithDefaults() *QuerySubOrdersSpotAlgoResponse {
	this := QuerySubOrdersSpotAlgoResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QuerySubOrdersSpotAlgoResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersSpotAlgoResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QuerySubOrdersSpotAlgoResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QuerySubOrdersSpotAlgoResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *QuerySubOrdersSpotAlgoResponse) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersSpotAlgoResponse) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *QuerySubOrdersSpotAlgoResponse) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *QuerySubOrdersSpotAlgoResponse) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetExecutedAmt returns the ExecutedAmt field value if set, zero value otherwise.
func (o *QuerySubOrdersSpotAlgoResponse) GetExecutedAmt() string {
	if o == nil || common.IsNil(o.ExecutedAmt) {
		var ret string
		return ret
	}
	return *o.ExecutedAmt
}

// GetExecutedAmtOk returns a tuple with the ExecutedAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersSpotAlgoResponse) GetExecutedAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedAmt) {
		return nil, false
	}
	return o.ExecutedAmt, true
}

// HasExecutedAmt returns a boolean if a field has been set.
func (o *QuerySubOrdersSpotAlgoResponse) HasExecutedAmt() bool {
	if o != nil && !common.IsNil(o.ExecutedAmt) {
		return true
	}

	return false
}

// SetExecutedAmt gets a reference to the given string and assigns it to the ExecutedAmt field.
func (o *QuerySubOrdersSpotAlgoResponse) SetExecutedAmt(v string) {
	o.ExecutedAmt = &v
}

// GetSubOrders returns the SubOrders field value if set, zero value otherwise.
func (o *QuerySubOrdersSpotAlgoResponse) GetSubOrders() []QuerySubOrdersFutureAlgoResponseSubOrdersInner {
	if o == nil || common.IsNil(o.SubOrders) {
		var ret []QuerySubOrdersFutureAlgoResponseSubOrdersInner
		return ret
	}
	return o.SubOrders
}

// GetSubOrdersOk returns a tuple with the SubOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersSpotAlgoResponse) GetSubOrdersOk() ([]QuerySubOrdersFutureAlgoResponseSubOrdersInner, bool) {
	if o == nil || common.IsNil(o.SubOrders) {
		return nil, false
	}
	return o.SubOrders, true
}

// HasSubOrders returns a boolean if a field has been set.
func (o *QuerySubOrdersSpotAlgoResponse) HasSubOrders() bool {
	if o != nil && !common.IsNil(o.SubOrders) {
		return true
	}

	return false
}

// SetSubOrders gets a reference to the given []QuerySubOrdersFutureAlgoResponseSubOrdersInner and assigns it to the SubOrders field.
func (o *QuerySubOrdersSpotAlgoResponse) SetSubOrders(v []QuerySubOrdersFutureAlgoResponseSubOrdersInner) {
	o.SubOrders = v
}

func (o QuerySubOrdersSpotAlgoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubOrdersSpotAlgoResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QuerySubOrdersSpotAlgoResponse) UnmarshalJSON(data []byte) (err error) {
	varQuerySubOrdersSpotAlgoResponse := _QuerySubOrdersSpotAlgoResponse{}

	err = json.Unmarshal(data, &varQuerySubOrdersSpotAlgoResponse)

	if err != nil {
		return err
	}

	*o = QuerySubOrdersSpotAlgoResponse(varQuerySubOrdersSpotAlgoResponse)

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

type NullableQuerySubOrdersSpotAlgoResponse struct {
	value *QuerySubOrdersSpotAlgoResponse
	isSet bool
}

func (v NullableQuerySubOrdersSpotAlgoResponse) Get() *QuerySubOrdersSpotAlgoResponse {
	return v.value
}

func (v *NullableQuerySubOrdersSpotAlgoResponse) Set(val *QuerySubOrdersSpotAlgoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubOrdersSpotAlgoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubOrdersSpotAlgoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubOrdersSpotAlgoResponse(val *QuerySubOrdersSpotAlgoResponse) *NullableQuerySubOrdersSpotAlgoResponse {
	return &NullableQuerySubOrdersSpotAlgoResponse{value: val, isSet: true}
}

func (v NullableQuerySubOrdersSpotAlgoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubOrdersSpotAlgoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
