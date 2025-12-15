/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AssetDividendRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AssetDividendRecordResponse{}

// AssetDividendRecordResponse struct for AssetDividendRecordResponse
type AssetDividendRecordResponse struct {
	Rows                 []AssetDividendRecordResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDividendRecordResponse AssetDividendRecordResponse

// NewAssetDividendRecordResponse instantiates a new AssetDividendRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDividendRecordResponse() *AssetDividendRecordResponse {
	this := AssetDividendRecordResponse{}
	return &this
}

// NewAssetDividendRecordResponseWithDefaults instantiates a new AssetDividendRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDividendRecordResponseWithDefaults() *AssetDividendRecordResponse {
	this := AssetDividendRecordResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *AssetDividendRecordResponse) GetRows() []AssetDividendRecordResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []AssetDividendRecordResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDividendRecordResponse) GetRowsOk() ([]AssetDividendRecordResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *AssetDividendRecordResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []AssetDividendRecordResponseRowsInner and assigns it to the Rows field.
func (o *AssetDividendRecordResponse) SetRows(v []AssetDividendRecordResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *AssetDividendRecordResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDividendRecordResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *AssetDividendRecordResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *AssetDividendRecordResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o AssetDividendRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetDividendRecordResponse) ToMap() (map[string]interface{}, error) {
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

func (o *AssetDividendRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varAssetDividendRecordResponse := _AssetDividendRecordResponse{}

	err = json.Unmarshal(data, &varAssetDividendRecordResponse)

	if err != nil {
		return err
	}

	*o = AssetDividendRecordResponse(varAssetDividendRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetDividendRecordResponse struct {
	value *AssetDividendRecordResponse
	isSet bool
}

func (v NullableAssetDividendRecordResponse) Get() *AssetDividendRecordResponse {
	return v.value
}

func (v *NullableAssetDividendRecordResponse) Set(val *AssetDividendRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDividendRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDividendRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDividendRecordResponse(val *AssetDividendRecordResponse) *NullableAssetDividendRecordResponse {
	return &NullableAssetDividendRecordResponse{value: val, isSet: true}
}

func (v NullableAssetDividendRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDividendRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
