/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFlexibleLoanOngoingOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanOngoingOrdersResponse{}

// GetFlexibleLoanOngoingOrdersResponse struct for GetFlexibleLoanOngoingOrdersResponse
type GetFlexibleLoanOngoingOrdersResponse struct {
	Rows                 []GetFlexibleLoanOngoingOrdersResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                          `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanOngoingOrdersResponse GetFlexibleLoanOngoingOrdersResponse

// NewGetFlexibleLoanOngoingOrdersResponse instantiates a new GetFlexibleLoanOngoingOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanOngoingOrdersResponse() *GetFlexibleLoanOngoingOrdersResponse {
	this := GetFlexibleLoanOngoingOrdersResponse{}
	return &this
}

// NewGetFlexibleLoanOngoingOrdersResponseWithDefaults instantiates a new GetFlexibleLoanOngoingOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanOngoingOrdersResponseWithDefaults() *GetFlexibleLoanOngoingOrdersResponse {
	this := GetFlexibleLoanOngoingOrdersResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetFlexibleLoanOngoingOrdersResponse) GetRows() []GetFlexibleLoanOngoingOrdersResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetFlexibleLoanOngoingOrdersResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanOngoingOrdersResponse) GetRowsOk() ([]GetFlexibleLoanOngoingOrdersResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetFlexibleLoanOngoingOrdersResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetFlexibleLoanOngoingOrdersResponseRowsInner and assigns it to the Rows field.
func (o *GetFlexibleLoanOngoingOrdersResponse) SetRows(v []GetFlexibleLoanOngoingOrdersResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFlexibleLoanOngoingOrdersResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanOngoingOrdersResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFlexibleLoanOngoingOrdersResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFlexibleLoanOngoingOrdersResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetFlexibleLoanOngoingOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanOngoingOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleLoanOngoingOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanOngoingOrdersResponse := _GetFlexibleLoanOngoingOrdersResponse{}

	err = json.Unmarshal(data, &varGetFlexibleLoanOngoingOrdersResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanOngoingOrdersResponse(varGetFlexibleLoanOngoingOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanOngoingOrdersResponse struct {
	value *GetFlexibleLoanOngoingOrdersResponse
	isSet bool
}

func (v NullableGetFlexibleLoanOngoingOrdersResponse) Get() *GetFlexibleLoanOngoingOrdersResponse {
	return v.value
}

func (v *NullableGetFlexibleLoanOngoingOrdersResponse) Set(val *GetFlexibleLoanOngoingOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanOngoingOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanOngoingOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanOngoingOrdersResponse(val *GetFlexibleLoanOngoingOrdersResponse) *NullableGetFlexibleLoanOngoingOrdersResponse {
	return &NullableGetFlexibleLoanOngoingOrdersResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanOngoingOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanOngoingOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
