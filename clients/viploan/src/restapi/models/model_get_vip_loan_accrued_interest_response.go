/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetVIPLoanAccruedInterestResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetVIPLoanAccruedInterestResponse{}

// GetVIPLoanAccruedInterestResponse struct for GetVIPLoanAccruedInterestResponse
type GetVIPLoanAccruedInterestResponse struct {
	Rows                 []GetVIPLoanAccruedInterestResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                       `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetVIPLoanAccruedInterestResponse GetVIPLoanAccruedInterestResponse

// NewGetVIPLoanAccruedInterestResponse instantiates a new GetVIPLoanAccruedInterestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVIPLoanAccruedInterestResponse() *GetVIPLoanAccruedInterestResponse {
	this := GetVIPLoanAccruedInterestResponse{}
	return &this
}

// NewGetVIPLoanAccruedInterestResponseWithDefaults instantiates a new GetVIPLoanAccruedInterestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVIPLoanAccruedInterestResponseWithDefaults() *GetVIPLoanAccruedInterestResponse {
	this := GetVIPLoanAccruedInterestResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetVIPLoanAccruedInterestResponse) GetRows() []GetVIPLoanAccruedInterestResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetVIPLoanAccruedInterestResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanAccruedInterestResponse) GetRowsOk() ([]GetVIPLoanAccruedInterestResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetVIPLoanAccruedInterestResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetVIPLoanAccruedInterestResponseRowsInner and assigns it to the Rows field.
func (o *GetVIPLoanAccruedInterestResponse) SetRows(v []GetVIPLoanAccruedInterestResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetVIPLoanAccruedInterestResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanAccruedInterestResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetVIPLoanAccruedInterestResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetVIPLoanAccruedInterestResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetVIPLoanAccruedInterestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVIPLoanAccruedInterestResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetVIPLoanAccruedInterestResponse) UnmarshalJSON(data []byte) (err error) {
	varGetVIPLoanAccruedInterestResponse := _GetVIPLoanAccruedInterestResponse{}

	err = json.Unmarshal(data, &varGetVIPLoanAccruedInterestResponse)

	if err != nil {
		return err
	}

	*o = GetVIPLoanAccruedInterestResponse(varGetVIPLoanAccruedInterestResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetVIPLoanAccruedInterestResponse struct {
	value *GetVIPLoanAccruedInterestResponse
	isSet bool
}

func (v NullableGetVIPLoanAccruedInterestResponse) Get() *GetVIPLoanAccruedInterestResponse {
	return v.value
}

func (v *NullableGetVIPLoanAccruedInterestResponse) Set(val *GetVIPLoanAccruedInterestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVIPLoanAccruedInterestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVIPLoanAccruedInterestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVIPLoanAccruedInterestResponse(val *GetVIPLoanAccruedInterestResponse) *NullableGetVIPLoanAccruedInterestResponse {
	return &NullableGetVIPLoanAccruedInterestResponse{value: val, isSet: true}
}

func (v NullableGetVIPLoanAccruedInterestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVIPLoanAccruedInterestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
