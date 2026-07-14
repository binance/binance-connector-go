/*
VIP Loan REST API

Access over-collateralized loan services, manage positions, and monitor collateral via the VIP Loan API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryVIPLoanFixedRateMarketResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryVIPLoanFixedRateMarketResponse{}

// QueryVIPLoanFixedRateMarketResponse struct for QueryVIPLoanFixedRateMarketResponse
type QueryVIPLoanFixedRateMarketResponse struct {
	// Total number of records
	Total *int64 `json:"total,omitempty"`
	// Current page data
	Rows                 []QueryVIPLoanFixedRateMarketResponseRowsInner `json:"rows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryVIPLoanFixedRateMarketResponse QueryVIPLoanFixedRateMarketResponse

// NewQueryVIPLoanFixedRateMarketResponse instantiates a new QueryVIPLoanFixedRateMarketResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryVIPLoanFixedRateMarketResponse() *QueryVIPLoanFixedRateMarketResponse {
	this := QueryVIPLoanFixedRateMarketResponse{}
	return &this
}

// NewQueryVIPLoanFixedRateMarketResponseWithDefaults instantiates a new QueryVIPLoanFixedRateMarketResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryVIPLoanFixedRateMarketResponseWithDefaults() *QueryVIPLoanFixedRateMarketResponse {
	this := QueryVIPLoanFixedRateMarketResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryVIPLoanFixedRateMarketResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryVIPLoanFixedRateMarketResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryVIPLoanFixedRateMarketResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryVIPLoanFixedRateMarketResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *QueryVIPLoanFixedRateMarketResponse) GetRows() []QueryVIPLoanFixedRateMarketResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []QueryVIPLoanFixedRateMarketResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryVIPLoanFixedRateMarketResponse) GetRowsOk() ([]QueryVIPLoanFixedRateMarketResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *QueryVIPLoanFixedRateMarketResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []QueryVIPLoanFixedRateMarketResponseRowsInner and assigns it to the Rows field.
func (o *QueryVIPLoanFixedRateMarketResponse) SetRows(v []QueryVIPLoanFixedRateMarketResponseRowsInner) {
	o.Rows = v
}

func (o QueryVIPLoanFixedRateMarketResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryVIPLoanFixedRateMarketResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryVIPLoanFixedRateMarketResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryVIPLoanFixedRateMarketResponse := _QueryVIPLoanFixedRateMarketResponse{}

	err = json.Unmarshal(data, &varQueryVIPLoanFixedRateMarketResponse)

	if err != nil {
		return err
	}

	*o = QueryVIPLoanFixedRateMarketResponse(varQueryVIPLoanFixedRateMarketResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "rows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryVIPLoanFixedRateMarketResponse struct {
	value *QueryVIPLoanFixedRateMarketResponse
	isSet bool
}

func (v NullableQueryVIPLoanFixedRateMarketResponse) Get() *QueryVIPLoanFixedRateMarketResponse {
	return v.value
}

func (v *NullableQueryVIPLoanFixedRateMarketResponse) Set(val *QueryVIPLoanFixedRateMarketResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryVIPLoanFixedRateMarketResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryVIPLoanFixedRateMarketResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryVIPLoanFixedRateMarketResponse(val *QueryVIPLoanFixedRateMarketResponse) *NullableQueryVIPLoanFixedRateMarketResponse {
	return &NullableQueryVIPLoanFixedRateMarketResponse{value: val, isSet: true}
}

func (v NullableQueryVIPLoanFixedRateMarketResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryVIPLoanFixedRateMarketResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
