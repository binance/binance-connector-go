/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetVIPLoanOngoingOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetVIPLoanOngoingOrdersResponse{}

// GetVIPLoanOngoingOrdersResponse struct for GetVIPLoanOngoingOrdersResponse
type GetVIPLoanOngoingOrdersResponse struct {
	Rows                 []GetVIPLoanOngoingOrdersResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                     `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetVIPLoanOngoingOrdersResponse GetVIPLoanOngoingOrdersResponse

// NewGetVIPLoanOngoingOrdersResponse instantiates a new GetVIPLoanOngoingOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVIPLoanOngoingOrdersResponse() *GetVIPLoanOngoingOrdersResponse {
	this := GetVIPLoanOngoingOrdersResponse{}
	return &this
}

// NewGetVIPLoanOngoingOrdersResponseWithDefaults instantiates a new GetVIPLoanOngoingOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVIPLoanOngoingOrdersResponseWithDefaults() *GetVIPLoanOngoingOrdersResponse {
	this := GetVIPLoanOngoingOrdersResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponse) GetRows() []GetVIPLoanOngoingOrdersResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetVIPLoanOngoingOrdersResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponse) GetRowsOk() ([]GetVIPLoanOngoingOrdersResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetVIPLoanOngoingOrdersResponseRowsInner and assigns it to the Rows field.
func (o *GetVIPLoanOngoingOrdersResponse) SetRows(v []GetVIPLoanOngoingOrdersResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetVIPLoanOngoingOrdersResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetVIPLoanOngoingOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVIPLoanOngoingOrdersResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetVIPLoanOngoingOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varGetVIPLoanOngoingOrdersResponse := _GetVIPLoanOngoingOrdersResponse{}

	err = json.Unmarshal(data, &varGetVIPLoanOngoingOrdersResponse)

	if err != nil {
		return err
	}

	*o = GetVIPLoanOngoingOrdersResponse(varGetVIPLoanOngoingOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetVIPLoanOngoingOrdersResponse struct {
	value *GetVIPLoanOngoingOrdersResponse
	isSet bool
}

func (v NullableGetVIPLoanOngoingOrdersResponse) Get() *GetVIPLoanOngoingOrdersResponse {
	return v.value
}

func (v *NullableGetVIPLoanOngoingOrdersResponse) Set(val *GetVIPLoanOngoingOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVIPLoanOngoingOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVIPLoanOngoingOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVIPLoanOngoingOrdersResponse(val *GetVIPLoanOngoingOrdersResponse) *NullableGetVIPLoanOngoingOrdersResponse {
	return &NullableGetVIPLoanOngoingOrdersResponse{value: val, isSet: true}
}

func (v NullableGetVIPLoanOngoingOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVIPLoanOngoingOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
