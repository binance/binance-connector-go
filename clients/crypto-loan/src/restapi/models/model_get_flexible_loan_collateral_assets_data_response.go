/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleLoanCollateralAssetsDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanCollateralAssetsDataResponse{}

// GetFlexibleLoanCollateralAssetsDataResponse struct for GetFlexibleLoanCollateralAssetsDataResponse
type GetFlexibleLoanCollateralAssetsDataResponse struct {
	Rows                 []GetFlexibleLoanCollateralAssetsDataResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                                 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanCollateralAssetsDataResponse GetFlexibleLoanCollateralAssetsDataResponse

// NewGetFlexibleLoanCollateralAssetsDataResponse instantiates a new GetFlexibleLoanCollateralAssetsDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanCollateralAssetsDataResponse() *GetFlexibleLoanCollateralAssetsDataResponse {
	this := GetFlexibleLoanCollateralAssetsDataResponse{}
	return &this
}

// NewGetFlexibleLoanCollateralAssetsDataResponseWithDefaults instantiates a new GetFlexibleLoanCollateralAssetsDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanCollateralAssetsDataResponseWithDefaults() *GetFlexibleLoanCollateralAssetsDataResponse {
	this := GetFlexibleLoanCollateralAssetsDataResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetFlexibleLoanCollateralAssetsDataResponse) GetRows() []GetFlexibleLoanCollateralAssetsDataResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetFlexibleLoanCollateralAssetsDataResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponse) GetRowsOk() ([]GetFlexibleLoanCollateralAssetsDataResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetFlexibleLoanCollateralAssetsDataResponseRowsInner and assigns it to the Rows field.
func (o *GetFlexibleLoanCollateralAssetsDataResponse) SetRows(v []GetFlexibleLoanCollateralAssetsDataResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFlexibleLoanCollateralAssetsDataResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFlexibleLoanCollateralAssetsDataResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFlexibleLoanCollateralAssetsDataResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetFlexibleLoanCollateralAssetsDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanCollateralAssetsDataResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFlexibleLoanCollateralAssetsDataResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanCollateralAssetsDataResponse := _GetFlexibleLoanCollateralAssetsDataResponse{}

	err = json.Unmarshal(data, &varGetFlexibleLoanCollateralAssetsDataResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanCollateralAssetsDataResponse(varGetFlexibleLoanCollateralAssetsDataResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanCollateralAssetsDataResponse struct {
	value *GetFlexibleLoanCollateralAssetsDataResponse
	isSet bool
}

func (v NullableGetFlexibleLoanCollateralAssetsDataResponse) Get() *GetFlexibleLoanCollateralAssetsDataResponse {
	return v.value
}

func (v *NullableGetFlexibleLoanCollateralAssetsDataResponse) Set(val *GetFlexibleLoanCollateralAssetsDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanCollateralAssetsDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanCollateralAssetsDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanCollateralAssetsDataResponse(val *GetFlexibleLoanCollateralAssetsDataResponse) *NullableGetFlexibleLoanCollateralAssetsDataResponse {
	return &NullableGetFlexibleLoanCollateralAssetsDataResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanCollateralAssetsDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanCollateralAssetsDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
