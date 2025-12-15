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

// checks if the QuerySubAccountFuturesAssetTransferHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountFuturesAssetTransferHistoryResponse{}

// QuerySubAccountFuturesAssetTransferHistoryResponse struct for QuerySubAccountFuturesAssetTransferHistoryResponse
type QuerySubAccountFuturesAssetTransferHistoryResponse struct {
	Success              *bool                                                              `json:"success,omitempty"`
	FuturesType          *int64                                                             `json:"futuresType,omitempty"`
	Transfers            []QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner `json:"transfers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubAccountFuturesAssetTransferHistoryResponse QuerySubAccountFuturesAssetTransferHistoryResponse

// NewQuerySubAccountFuturesAssetTransferHistoryResponse instantiates a new QuerySubAccountFuturesAssetTransferHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountFuturesAssetTransferHistoryResponse() *QuerySubAccountFuturesAssetTransferHistoryResponse {
	this := QuerySubAccountFuturesAssetTransferHistoryResponse{}
	return &this
}

// NewQuerySubAccountFuturesAssetTransferHistoryResponseWithDefaults instantiates a new QuerySubAccountFuturesAssetTransferHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountFuturesAssetTransferHistoryResponseWithDefaults() *QuerySubAccountFuturesAssetTransferHistoryResponse {
	this := QuerySubAccountFuturesAssetTransferHistoryResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetFuturesType returns the FuturesType field value if set, zero value otherwise.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) GetFuturesType() int64 {
	if o == nil || common.IsNil(o.FuturesType) {
		var ret int64
		return ret
	}
	return *o.FuturesType
}

// GetFuturesTypeOk returns a tuple with the FuturesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) GetFuturesTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FuturesType) {
		return nil, false
	}
	return o.FuturesType, true
}

// HasFuturesType returns a boolean if a field has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) HasFuturesType() bool {
	if o != nil && !common.IsNil(o.FuturesType) {
		return true
	}

	return false
}

// SetFuturesType gets a reference to the given int64 and assigns it to the FuturesType field.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) SetFuturesType(v int64) {
	o.FuturesType = &v
}

// GetTransfers returns the Transfers field value if set, zero value otherwise.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) GetTransfers() []QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner {
	if o == nil || common.IsNil(o.Transfers) {
		var ret []QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner
		return ret
	}
	return o.Transfers
}

// GetTransfersOk returns a tuple with the Transfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) GetTransfersOk() ([]QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner, bool) {
	if o == nil || common.IsNil(o.Transfers) {
		return nil, false
	}
	return o.Transfers, true
}

// HasTransfers returns a boolean if a field has been set.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) HasTransfers() bool {
	if o != nil && !common.IsNil(o.Transfers) {
		return true
	}

	return false
}

// SetTransfers gets a reference to the given []QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner and assigns it to the Transfers field.
func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) SetTransfers(v []QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner) {
	o.Transfers = v
}

func (o QuerySubAccountFuturesAssetTransferHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountFuturesAssetTransferHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.FuturesType) {
		toSerialize["futuresType"] = o.FuturesType
	}
	if !common.IsNil(o.Transfers) {
		toSerialize["transfers"] = o.Transfers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountFuturesAssetTransferHistoryResponse := _QuerySubAccountFuturesAssetTransferHistoryResponse{}

	err = json.Unmarshal(data, &varQuerySubAccountFuturesAssetTransferHistoryResponse)

	if err != nil {
		return err
	}

	*o = QuerySubAccountFuturesAssetTransferHistoryResponse(varQuerySubAccountFuturesAssetTransferHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "futuresType")
		delete(additionalProperties, "transfers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountFuturesAssetTransferHistoryResponse struct {
	value *QuerySubAccountFuturesAssetTransferHistoryResponse
	isSet bool
}

func (v NullableQuerySubAccountFuturesAssetTransferHistoryResponse) Get() *QuerySubAccountFuturesAssetTransferHistoryResponse {
	return v.value
}

func (v *NullableQuerySubAccountFuturesAssetTransferHistoryResponse) Set(val *QuerySubAccountFuturesAssetTransferHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountFuturesAssetTransferHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountFuturesAssetTransferHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountFuturesAssetTransferHistoryResponse(val *QuerySubAccountFuturesAssetTransferHistoryResponse) *NullableQuerySubAccountFuturesAssetTransferHistoryResponse {
	return &NullableQuerySubAccountFuturesAssetTransferHistoryResponse{value: val, isSet: true}
}

func (v NullableQuerySubAccountFuturesAssetTransferHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountFuturesAssetTransferHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
