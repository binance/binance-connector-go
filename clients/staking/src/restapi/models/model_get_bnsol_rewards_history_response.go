/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetBnsolRewardsHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBnsolRewardsHistoryResponse{}

// GetBnsolRewardsHistoryResponse struct for GetBnsolRewardsHistoryResponse
type GetBnsolRewardsHistoryResponse struct {
	EstRewardsInSOL      *string                                   `json:"estRewardsInSOL,omitempty"`
	Rows                 []GetBnsolRewardsHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                    `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBnsolRewardsHistoryResponse GetBnsolRewardsHistoryResponse

// NewGetBnsolRewardsHistoryResponse instantiates a new GetBnsolRewardsHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBnsolRewardsHistoryResponse() *GetBnsolRewardsHistoryResponse {
	this := GetBnsolRewardsHistoryResponse{}
	return &this
}

// NewGetBnsolRewardsHistoryResponseWithDefaults instantiates a new GetBnsolRewardsHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBnsolRewardsHistoryResponseWithDefaults() *GetBnsolRewardsHistoryResponse {
	this := GetBnsolRewardsHistoryResponse{}
	return &this
}

// GetEstRewardsInSOL returns the EstRewardsInSOL field value if set, zero value otherwise.
func (o *GetBnsolRewardsHistoryResponse) GetEstRewardsInSOL() string {
	if o == nil || common.IsNil(o.EstRewardsInSOL) {
		var ret string
		return ret
	}
	return *o.EstRewardsInSOL
}

// GetEstRewardsInSOLOk returns a tuple with the EstRewardsInSOL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRewardsHistoryResponse) GetEstRewardsInSOLOk() (*string, bool) {
	if o == nil || common.IsNil(o.EstRewardsInSOL) {
		return nil, false
	}
	return o.EstRewardsInSOL, true
}

// HasEstRewardsInSOL returns a boolean if a field has been set.
func (o *GetBnsolRewardsHistoryResponse) HasEstRewardsInSOL() bool {
	if o != nil && !common.IsNil(o.EstRewardsInSOL) {
		return true
	}

	return false
}

// SetEstRewardsInSOL gets a reference to the given string and assigns it to the EstRewardsInSOL field.
func (o *GetBnsolRewardsHistoryResponse) SetEstRewardsInSOL(v string) {
	o.EstRewardsInSOL = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetBnsolRewardsHistoryResponse) GetRows() []GetBnsolRewardsHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetBnsolRewardsHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRewardsHistoryResponse) GetRowsOk() ([]GetBnsolRewardsHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetBnsolRewardsHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetBnsolRewardsHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetBnsolRewardsHistoryResponse) SetRows(v []GetBnsolRewardsHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetBnsolRewardsHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRewardsHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetBnsolRewardsHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetBnsolRewardsHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetBnsolRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBnsolRewardsHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EstRewardsInSOL) {
		toSerialize["estRewardsInSOL"] = o.EstRewardsInSOL
	}
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

func (o *GetBnsolRewardsHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetBnsolRewardsHistoryResponse := _GetBnsolRewardsHistoryResponse{}

	err = json.Unmarshal(data, &varGetBnsolRewardsHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetBnsolRewardsHistoryResponse(varGetBnsolRewardsHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "estRewardsInSOL")
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBnsolRewardsHistoryResponse struct {
	value *GetBnsolRewardsHistoryResponse
	isSet bool
}

func (v NullableGetBnsolRewardsHistoryResponse) Get() *GetBnsolRewardsHistoryResponse {
	return v.value
}

func (v *NullableGetBnsolRewardsHistoryResponse) Set(val *GetBnsolRewardsHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBnsolRewardsHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBnsolRewardsHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBnsolRewardsHistoryResponse(val *GetBnsolRewardsHistoryResponse) *NullableGetBnsolRewardsHistoryResponse {
	return &NullableGetBnsolRewardsHistoryResponse{value: val, isSet: true}
}

func (v NullableGetBnsolRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBnsolRewardsHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
