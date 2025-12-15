/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetMovePositionHistoryForSubAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMovePositionHistoryForSubAccountResponse{}

// GetMovePositionHistoryForSubAccountResponse struct for GetMovePositionHistoryForSubAccountResponse
type GetMovePositionHistoryForSubAccountResponse struct {
	Total                         *int64                                                                          `json:"total,omitempty"`
	FutureMovePositionOrderVoList []GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner `json:"futureMovePositionOrderVoList,omitempty"`
	AdditionalProperties          map[string]interface{}
}

type _GetMovePositionHistoryForSubAccountResponse GetMovePositionHistoryForSubAccountResponse

// NewGetMovePositionHistoryForSubAccountResponse instantiates a new GetMovePositionHistoryForSubAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMovePositionHistoryForSubAccountResponse() *GetMovePositionHistoryForSubAccountResponse {
	this := GetMovePositionHistoryForSubAccountResponse{}
	return &this
}

// NewGetMovePositionHistoryForSubAccountResponseWithDefaults instantiates a new GetMovePositionHistoryForSubAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMovePositionHistoryForSubAccountResponseWithDefaults() *GetMovePositionHistoryForSubAccountResponse {
	this := GetMovePositionHistoryForSubAccountResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetMovePositionHistoryForSubAccountResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovePositionHistoryForSubAccountResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetMovePositionHistoryForSubAccountResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetMovePositionHistoryForSubAccountResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetFutureMovePositionOrderVoList returns the FutureMovePositionOrderVoList field value if set, zero value otherwise.
func (o *GetMovePositionHistoryForSubAccountResponse) GetFutureMovePositionOrderVoList() []GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner {
	if o == nil || common.IsNil(o.FutureMovePositionOrderVoList) {
		var ret []GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner
		return ret
	}
	return o.FutureMovePositionOrderVoList
}

// GetFutureMovePositionOrderVoListOk returns a tuple with the FutureMovePositionOrderVoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovePositionHistoryForSubAccountResponse) GetFutureMovePositionOrderVoListOk() ([]GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner, bool) {
	if o == nil || common.IsNil(o.FutureMovePositionOrderVoList) {
		return nil, false
	}
	return o.FutureMovePositionOrderVoList, true
}

// HasFutureMovePositionOrderVoList returns a boolean if a field has been set.
func (o *GetMovePositionHistoryForSubAccountResponse) HasFutureMovePositionOrderVoList() bool {
	if o != nil && !common.IsNil(o.FutureMovePositionOrderVoList) {
		return true
	}

	return false
}

// SetFutureMovePositionOrderVoList gets a reference to the given []GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner and assigns it to the FutureMovePositionOrderVoList field.
func (o *GetMovePositionHistoryForSubAccountResponse) SetFutureMovePositionOrderVoList(v []GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) {
	o.FutureMovePositionOrderVoList = v
}

func (o GetMovePositionHistoryForSubAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMovePositionHistoryForSubAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.FutureMovePositionOrderVoList) {
		toSerialize["futureMovePositionOrderVoList"] = o.FutureMovePositionOrderVoList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMovePositionHistoryForSubAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varGetMovePositionHistoryForSubAccountResponse := _GetMovePositionHistoryForSubAccountResponse{}

	err = json.Unmarshal(data, &varGetMovePositionHistoryForSubAccountResponse)

	if err != nil {
		return err
	}

	*o = GetMovePositionHistoryForSubAccountResponse(varGetMovePositionHistoryForSubAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "futureMovePositionOrderVoList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMovePositionHistoryForSubAccountResponse struct {
	value *GetMovePositionHistoryForSubAccountResponse
	isSet bool
}

func (v NullableGetMovePositionHistoryForSubAccountResponse) Get() *GetMovePositionHistoryForSubAccountResponse {
	return v.value
}

func (v *NullableGetMovePositionHistoryForSubAccountResponse) Set(val *GetMovePositionHistoryForSubAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMovePositionHistoryForSubAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMovePositionHistoryForSubAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMovePositionHistoryForSubAccountResponse(val *GetMovePositionHistoryForSubAccountResponse) *NullableGetMovePositionHistoryForSubAccountResponse {
	return &NullableGetMovePositionHistoryForSubAccountResponse{value: val, isSet: true}
}

func (v NullableGetMovePositionHistoryForSubAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMovePositionHistoryForSubAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
