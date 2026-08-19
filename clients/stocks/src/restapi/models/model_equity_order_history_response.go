/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the EquityOrderHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EquityOrderHistoryResponse{}

// EquityOrderHistoryResponse struct for EquityOrderHistoryResponse
type EquityOrderHistoryResponse struct {
	// Total number of rows matching the filter.
	Total *int64 `json:"total,omitempty"`
	// Current page (echoes `current`).
	Page *int32 `json:"page,omitempty"`
	// Current page size (echoes `size`).
	Size *int32 `json:"size,omitempty"`
	// Order rows on this page. Empty array if nothing matches.
	Rows                 []EquityOrderHistoryResponseRowsInner `json:"rows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquityOrderHistoryResponse EquityOrderHistoryResponse

// NewEquityOrderHistoryResponse instantiates a new EquityOrderHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityOrderHistoryResponse() *EquityOrderHistoryResponse {
	this := EquityOrderHistoryResponse{}
	return &this
}

// NewEquityOrderHistoryResponseWithDefaults instantiates a new EquityOrderHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityOrderHistoryResponseWithDefaults() *EquityOrderHistoryResponse {
	this := EquityOrderHistoryResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *EquityOrderHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponse) GetPage() int32 {
	if o == nil || common.IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponse) GetPageOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponse) HasPage() bool {
	if o != nil && !common.IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *EquityOrderHistoryResponse) SetPage(v int32) {
	o.Page = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponse) GetSize() int32 {
	if o == nil || common.IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponse) GetSizeOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponse) HasSize() bool {
	if o != nil && !common.IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *EquityOrderHistoryResponse) SetSize(v int32) {
	o.Size = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponse) GetRows() []EquityOrderHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []EquityOrderHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponse) GetRowsOk() ([]EquityOrderHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []EquityOrderHistoryResponseRowsInner and assigns it to the Rows field.
func (o *EquityOrderHistoryResponse) SetRows(v []EquityOrderHistoryResponseRowsInner) {
	o.Rows = v
}

func (o EquityOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquityOrderHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !common.IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !common.IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquityOrderHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varEquityOrderHistoryResponse := _EquityOrderHistoryResponse{}

	err = json.Unmarshal(data, &varEquityOrderHistoryResponse)

	if err != nil {
		return err
	}

	*o = EquityOrderHistoryResponse(varEquityOrderHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "page")
		delete(additionalProperties, "size")
		delete(additionalProperties, "rows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquityOrderHistoryResponse struct {
	value *EquityOrderHistoryResponse
	isSet bool
}

func (v NullableEquityOrderHistoryResponse) Get() *EquityOrderHistoryResponse {
	return v.value
}

func (v *NullableEquityOrderHistoryResponse) Set(val *EquityOrderHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityOrderHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityOrderHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityOrderHistoryResponse(val *EquityOrderHistoryResponse) *NullableEquityOrderHistoryResponse {
	return &NullableEquityOrderHistoryResponse{value: val, isSet: true}
}

func (v NullableEquityOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityOrderHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
