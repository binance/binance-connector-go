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

// checks if the GetLoanableAssetsDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLoanableAssetsDataResponse{}

// GetLoanableAssetsDataResponse struct for GetLoanableAssetsDataResponse
type GetLoanableAssetsDataResponse struct {
	Rows                 []GetLoanableAssetsDataResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLoanableAssetsDataResponse GetLoanableAssetsDataResponse

// NewGetLoanableAssetsDataResponse instantiates a new GetLoanableAssetsDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLoanableAssetsDataResponse() *GetLoanableAssetsDataResponse {
	this := GetLoanableAssetsDataResponse{}
	return &this
}

// NewGetLoanableAssetsDataResponseWithDefaults instantiates a new GetLoanableAssetsDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLoanableAssetsDataResponseWithDefaults() *GetLoanableAssetsDataResponse {
	this := GetLoanableAssetsDataResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetLoanableAssetsDataResponse) GetRows() []GetLoanableAssetsDataResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetLoanableAssetsDataResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanableAssetsDataResponse) GetRowsOk() ([]GetLoanableAssetsDataResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetLoanableAssetsDataResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetLoanableAssetsDataResponseRowsInner and assigns it to the Rows field.
func (o *GetLoanableAssetsDataResponse) SetRows(v []GetLoanableAssetsDataResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetLoanableAssetsDataResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanableAssetsDataResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetLoanableAssetsDataResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetLoanableAssetsDataResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetLoanableAssetsDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLoanableAssetsDataResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetLoanableAssetsDataResponse) UnmarshalJSON(data []byte) (err error) {
	varGetLoanableAssetsDataResponse := _GetLoanableAssetsDataResponse{}

	err = json.Unmarshal(data, &varGetLoanableAssetsDataResponse)

	if err != nil {
		return err
	}

	*o = GetLoanableAssetsDataResponse(varGetLoanableAssetsDataResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLoanableAssetsDataResponse struct {
	value *GetLoanableAssetsDataResponse
	isSet bool
}

func (v NullableGetLoanableAssetsDataResponse) Get() *GetLoanableAssetsDataResponse {
	return v.value
}

func (v *NullableGetLoanableAssetsDataResponse) Set(val *GetLoanableAssetsDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLoanableAssetsDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLoanableAssetsDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLoanableAssetsDataResponse(val *GetLoanableAssetsDataResponse) *NullableGetLoanableAssetsDataResponse {
	return &NullableGetLoanableAssetsDataResponse{value: val, isSet: true}
}

func (v NullableGetLoanableAssetsDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLoanableAssetsDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
