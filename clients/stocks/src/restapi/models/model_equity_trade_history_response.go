/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the EquityTradeHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EquityTradeHistoryResponse{}

// EquityTradeHistoryResponse struct for EquityTradeHistoryResponse
type EquityTradeHistoryResponse struct {
	// Total number of rows matching the filter.
	Total *int64 `json:"total,omitempty"`
	// Current page.
	Page *int32 `json:"page,omitempty"`
	// Current page size.
	Size *int32 `json:"size,omitempty"`
	// Trade rows on this page. Empty when nothing matches.
	Rows                 []EquityTradeHistoryResponseRowsInner `json:"rows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquityTradeHistoryResponse EquityTradeHistoryResponse

// NewEquityTradeHistoryResponse instantiates a new EquityTradeHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityTradeHistoryResponse() *EquityTradeHistoryResponse {
	this := EquityTradeHistoryResponse{}
	return &this
}

// NewEquityTradeHistoryResponseWithDefaults instantiates a new EquityTradeHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityTradeHistoryResponseWithDefaults() *EquityTradeHistoryResponse {
	this := EquityTradeHistoryResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *EquityTradeHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponse) GetPage() int32 {
	if o == nil || common.IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponse) GetPageOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponse) HasPage() bool {
	if o != nil && !common.IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *EquityTradeHistoryResponse) SetPage(v int32) {
	o.Page = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponse) GetSize() int32 {
	if o == nil || common.IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponse) GetSizeOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponse) HasSize() bool {
	if o != nil && !common.IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *EquityTradeHistoryResponse) SetSize(v int32) {
	o.Size = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponse) GetRows() []EquityTradeHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []EquityTradeHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponse) GetRowsOk() ([]EquityTradeHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []EquityTradeHistoryResponseRowsInner and assigns it to the Rows field.
func (o *EquityTradeHistoryResponse) SetRows(v []EquityTradeHistoryResponseRowsInner) {
	o.Rows = v
}

func (o EquityTradeHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquityTradeHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *EquityTradeHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varEquityTradeHistoryResponse := _EquityTradeHistoryResponse{}

	err = json.Unmarshal(data, &varEquityTradeHistoryResponse)

	if err != nil {
		return err
	}

	*o = EquityTradeHistoryResponse(varEquityTradeHistoryResponse)

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

type NullableEquityTradeHistoryResponse struct {
	value *EquityTradeHistoryResponse
	isSet bool
}

func (v NullableEquityTradeHistoryResponse) Get() *EquityTradeHistoryResponse {
	return v.value
}

func (v *NullableEquityTradeHistoryResponse) Set(val *EquityTradeHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityTradeHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityTradeHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityTradeHistoryResponse(val *EquityTradeHistoryResponse) *NullableEquityTradeHistoryResponse {
	return &NullableEquityTradeHistoryResponse{value: val, isSet: true}
}

func (v NullableEquityTradeHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityTradeHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
