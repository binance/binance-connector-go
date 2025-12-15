/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CheckVIPLoanCollateralAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckVIPLoanCollateralAccountResponse{}

// CheckVIPLoanCollateralAccountResponse struct for CheckVIPLoanCollateralAccountResponse
type CheckVIPLoanCollateralAccountResponse struct {
	Rows                 []CheckVIPLoanCollateralAccountResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                           `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CheckVIPLoanCollateralAccountResponse CheckVIPLoanCollateralAccountResponse

// NewCheckVIPLoanCollateralAccountResponse instantiates a new CheckVIPLoanCollateralAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckVIPLoanCollateralAccountResponse() *CheckVIPLoanCollateralAccountResponse {
	this := CheckVIPLoanCollateralAccountResponse{}
	return &this
}

// NewCheckVIPLoanCollateralAccountResponseWithDefaults instantiates a new CheckVIPLoanCollateralAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckVIPLoanCollateralAccountResponseWithDefaults() *CheckVIPLoanCollateralAccountResponse {
	this := CheckVIPLoanCollateralAccountResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *CheckVIPLoanCollateralAccountResponse) GetRows() []CheckVIPLoanCollateralAccountResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []CheckVIPLoanCollateralAccountResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVIPLoanCollateralAccountResponse) GetRowsOk() ([]CheckVIPLoanCollateralAccountResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *CheckVIPLoanCollateralAccountResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []CheckVIPLoanCollateralAccountResponseRowsInner and assigns it to the Rows field.
func (o *CheckVIPLoanCollateralAccountResponse) SetRows(v []CheckVIPLoanCollateralAccountResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CheckVIPLoanCollateralAccountResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVIPLoanCollateralAccountResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CheckVIPLoanCollateralAccountResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *CheckVIPLoanCollateralAccountResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o CheckVIPLoanCollateralAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckVIPLoanCollateralAccountResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CheckVIPLoanCollateralAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varCheckVIPLoanCollateralAccountResponse := _CheckVIPLoanCollateralAccountResponse{}

	err = json.Unmarshal(data, &varCheckVIPLoanCollateralAccountResponse)

	if err != nil {
		return err
	}

	*o = CheckVIPLoanCollateralAccountResponse(varCheckVIPLoanCollateralAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCheckVIPLoanCollateralAccountResponse struct {
	value *CheckVIPLoanCollateralAccountResponse
	isSet bool
}

func (v NullableCheckVIPLoanCollateralAccountResponse) Get() *CheckVIPLoanCollateralAccountResponse {
	return v.value
}

func (v *NullableCheckVIPLoanCollateralAccountResponse) Set(val *CheckVIPLoanCollateralAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckVIPLoanCollateralAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckVIPLoanCollateralAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckVIPLoanCollateralAccountResponse(val *CheckVIPLoanCollateralAccountResponse) *NullableCheckVIPLoanCollateralAccountResponse {
	return &NullableCheckVIPLoanCollateralAccountResponse{value: val, isSet: true}
}

func (v NullableCheckVIPLoanCollateralAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckVIPLoanCollateralAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
