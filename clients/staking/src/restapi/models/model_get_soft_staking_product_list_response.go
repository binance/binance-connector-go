/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSoftStakingProductListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSoftStakingProductListResponse{}

// GetSoftStakingProductListResponse struct for GetSoftStakingProductListResponse
type GetSoftStakingProductListResponse struct {
	Status               *bool                                        `json:"status,omitempty"`
	TotalRewardsUsdt     *string                                      `json:"totalRewardsUsdt,omitempty"`
	Rows                 []GetSoftStakingProductListResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                       `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSoftStakingProductListResponse GetSoftStakingProductListResponse

// NewGetSoftStakingProductListResponse instantiates a new GetSoftStakingProductListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSoftStakingProductListResponse() *GetSoftStakingProductListResponse {
	this := GetSoftStakingProductListResponse{}
	return &this
}

// NewGetSoftStakingProductListResponseWithDefaults instantiates a new GetSoftStakingProductListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSoftStakingProductListResponseWithDefaults() *GetSoftStakingProductListResponse {
	this := GetSoftStakingProductListResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetSoftStakingProductListResponse) GetStatus() bool {
	if o == nil || common.IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingProductListResponse) GetStatusOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetSoftStakingProductListResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *GetSoftStakingProductListResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetTotalRewardsUsdt returns the TotalRewardsUsdt field value if set, zero value otherwise.
func (o *GetSoftStakingProductListResponse) GetTotalRewardsUsdt() string {
	if o == nil || common.IsNil(o.TotalRewardsUsdt) {
		var ret string
		return ret
	}
	return *o.TotalRewardsUsdt
}

// GetTotalRewardsUsdtOk returns a tuple with the TotalRewardsUsdt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingProductListResponse) GetTotalRewardsUsdtOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalRewardsUsdt) {
		return nil, false
	}
	return o.TotalRewardsUsdt, true
}

// HasTotalRewardsUsdt returns a boolean if a field has been set.
func (o *GetSoftStakingProductListResponse) HasTotalRewardsUsdt() bool {
	if o != nil && !common.IsNil(o.TotalRewardsUsdt) {
		return true
	}

	return false
}

// SetTotalRewardsUsdt gets a reference to the given string and assigns it to the TotalRewardsUsdt field.
func (o *GetSoftStakingProductListResponse) SetTotalRewardsUsdt(v string) {
	o.TotalRewardsUsdt = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetSoftStakingProductListResponse) GetRows() []GetSoftStakingProductListResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetSoftStakingProductListResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingProductListResponse) GetRowsOk() ([]GetSoftStakingProductListResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetSoftStakingProductListResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetSoftStakingProductListResponseRowsInner and assigns it to the Rows field.
func (o *GetSoftStakingProductListResponse) SetRows(v []GetSoftStakingProductListResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetSoftStakingProductListResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingProductListResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetSoftStakingProductListResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetSoftStakingProductListResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetSoftStakingProductListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSoftStakingProductListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TotalRewardsUsdt) {
		toSerialize["totalRewardsUsdt"] = o.TotalRewardsUsdt
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

func (o *GetSoftStakingProductListResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSoftStakingProductListResponse := _GetSoftStakingProductListResponse{}

	err = json.Unmarshal(data, &varGetSoftStakingProductListResponse)

	if err != nil {
		return err
	}

	*o = GetSoftStakingProductListResponse(varGetSoftStakingProductListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "totalRewardsUsdt")
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSoftStakingProductListResponse struct {
	value *GetSoftStakingProductListResponse
	isSet bool
}

func (v NullableGetSoftStakingProductListResponse) Get() *GetSoftStakingProductListResponse {
	return v.value
}

func (v *NullableGetSoftStakingProductListResponse) Set(val *GetSoftStakingProductListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSoftStakingProductListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSoftStakingProductListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSoftStakingProductListResponse(val *GetSoftStakingProductListResponse) *NullableGetSoftStakingProductListResponse {
	return &NullableGetSoftStakingProductListResponse{value: val, isSet: true}
}

func (v NullableGetSoftStakingProductListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSoftStakingProductListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
