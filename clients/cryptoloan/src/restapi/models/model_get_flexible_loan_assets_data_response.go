/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFlexibleLoanAssetsDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanAssetsDataResponse{}

// GetFlexibleLoanAssetsDataResponse struct for GetFlexibleLoanAssetsDataResponse
type GetFlexibleLoanAssetsDataResponse struct {
	Rows                 []GetFlexibleLoanAssetsDataResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                       `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanAssetsDataResponse GetFlexibleLoanAssetsDataResponse

// NewGetFlexibleLoanAssetsDataResponse instantiates a new GetFlexibleLoanAssetsDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanAssetsDataResponse() *GetFlexibleLoanAssetsDataResponse {
	this := GetFlexibleLoanAssetsDataResponse{}
	return &this
}

// NewGetFlexibleLoanAssetsDataResponseWithDefaults instantiates a new GetFlexibleLoanAssetsDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanAssetsDataResponseWithDefaults() *GetFlexibleLoanAssetsDataResponse {
	this := GetFlexibleLoanAssetsDataResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetFlexibleLoanAssetsDataResponse) GetRows() []GetFlexibleLoanAssetsDataResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetFlexibleLoanAssetsDataResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanAssetsDataResponse) GetRowsOk() ([]GetFlexibleLoanAssetsDataResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetFlexibleLoanAssetsDataResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetFlexibleLoanAssetsDataResponseRowsInner and assigns it to the Rows field.
func (o *GetFlexibleLoanAssetsDataResponse) SetRows(v []GetFlexibleLoanAssetsDataResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFlexibleLoanAssetsDataResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanAssetsDataResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFlexibleLoanAssetsDataResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFlexibleLoanAssetsDataResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetFlexibleLoanAssetsDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanAssetsDataResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFlexibleLoanAssetsDataResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanAssetsDataResponse := _GetFlexibleLoanAssetsDataResponse{}

	err = json.Unmarshal(data, &varGetFlexibleLoanAssetsDataResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanAssetsDataResponse(varGetFlexibleLoanAssetsDataResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanAssetsDataResponse struct {
	value *GetFlexibleLoanAssetsDataResponse
	isSet bool
}

func (v NullableGetFlexibleLoanAssetsDataResponse) Get() *GetFlexibleLoanAssetsDataResponse {
	return v.value
}

func (v *NullableGetFlexibleLoanAssetsDataResponse) Set(val *GetFlexibleLoanAssetsDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanAssetsDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanAssetsDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanAssetsDataResponse(val *GetFlexibleLoanAssetsDataResponse) *NullableGetFlexibleLoanAssetsDataResponse {
	return &NullableGetFlexibleLoanAssetsDataResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanAssetsDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanAssetsDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
