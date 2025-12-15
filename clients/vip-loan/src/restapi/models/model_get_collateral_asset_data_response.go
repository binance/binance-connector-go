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

// checks if the GetCollateralAssetDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCollateralAssetDataResponse{}

// GetCollateralAssetDataResponse struct for GetCollateralAssetDataResponse
type GetCollateralAssetDataResponse struct {
	Rows                 []GetCollateralAssetDataResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                    `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCollateralAssetDataResponse GetCollateralAssetDataResponse

// NewGetCollateralAssetDataResponse instantiates a new GetCollateralAssetDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCollateralAssetDataResponse() *GetCollateralAssetDataResponse {
	this := GetCollateralAssetDataResponse{}
	return &this
}

// NewGetCollateralAssetDataResponseWithDefaults instantiates a new GetCollateralAssetDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCollateralAssetDataResponseWithDefaults() *GetCollateralAssetDataResponse {
	this := GetCollateralAssetDataResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetCollateralAssetDataResponse) GetRows() []GetCollateralAssetDataResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetCollateralAssetDataResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralAssetDataResponse) GetRowsOk() ([]GetCollateralAssetDataResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetCollateralAssetDataResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetCollateralAssetDataResponseRowsInner and assigns it to the Rows field.
func (o *GetCollateralAssetDataResponse) SetRows(v []GetCollateralAssetDataResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetCollateralAssetDataResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralAssetDataResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetCollateralAssetDataResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetCollateralAssetDataResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetCollateralAssetDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCollateralAssetDataResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetCollateralAssetDataResponse) UnmarshalJSON(data []byte) (err error) {
	varGetCollateralAssetDataResponse := _GetCollateralAssetDataResponse{}

	err = json.Unmarshal(data, &varGetCollateralAssetDataResponse)

	if err != nil {
		return err
	}

	*o = GetCollateralAssetDataResponse(varGetCollateralAssetDataResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCollateralAssetDataResponse struct {
	value *GetCollateralAssetDataResponse
	isSet bool
}

func (v NullableGetCollateralAssetDataResponse) Get() *GetCollateralAssetDataResponse {
	return v.value
}

func (v *NullableGetCollateralAssetDataResponse) Set(val *GetCollateralAssetDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCollateralAssetDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCollateralAssetDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCollateralAssetDataResponse(val *GetCollateralAssetDataResponse) *NullableGetCollateralAssetDataResponse {
	return &NullableGetCollateralAssetDataResponse{value: val, isSet: true}
}

func (v NullableGetCollateralAssetDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCollateralAssetDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
