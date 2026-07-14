/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPnLResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPnLResponse{}

// QueryPnLResponse struct for QueryPnLResponse
type QueryPnLResponse struct {
	ChainId              *string                              `json:"chainId,omitempty"`
	WalletAddress        *string                              `json:"walletAddress,omitempty"`
	Pnl                  map[string]interface{}               `json:"pnl,omitempty"`
	PnlList              []GetPortfolioResponsePositionsInner `json:"pnlList,omitempty"`
	TotalCount           *int32                               `json:"totalCount,omitempty"`
	TotalRealizedPnl     *string                              `json:"totalRealizedPnl,omitempty"`
	TotalUnrealizedPnl   *string                              `json:"totalUnrealizedPnl,omitempty"`
	TotalPnl             *string                              `json:"totalPnl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryPnLResponse QueryPnLResponse

// NewQueryPnLResponse instantiates a new QueryPnLResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPnLResponse() *QueryPnLResponse {
	this := QueryPnLResponse{}
	return &this
}

// NewQueryPnLResponseWithDefaults instantiates a new QueryPnLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPnLResponseWithDefaults() *QueryPnLResponse {
	this := QueryPnLResponse{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *QueryPnLResponse) GetChainId() string {
	if o == nil || common.IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPnLResponse) GetChainIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *QueryPnLResponse) HasChainId() bool {
	if o != nil && !common.IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *QueryPnLResponse) SetChainId(v string) {
	o.ChainId = &v
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise.
func (o *QueryPnLResponse) GetWalletAddress() string {
	if o == nil || common.IsNil(o.WalletAddress) {
		var ret string
		return ret
	}
	return *o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPnLResponse) GetWalletAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletAddress) {
		return nil, false
	}
	return o.WalletAddress, true
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *QueryPnLResponse) HasWalletAddress() bool {
	if o != nil && !common.IsNil(o.WalletAddress) {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given string and assigns it to the WalletAddress field.
func (o *QueryPnLResponse) SetWalletAddress(v string) {
	o.WalletAddress = &v
}

// GetPnl returns the Pnl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryPnLResponse) GetPnl() map[string]interface{} {
	if o == nil || common.IsNil(o.Pnl) {
		var ret map[string]interface{}
		return ret
	}
	return o.Pnl
}

// GetPnlOk returns a tuple with the Pnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryPnLResponse) GetPnlOk() (map[string]interface{}, bool) {
	if o == nil || common.IsNil(o.Pnl) {
		return map[string]interface{}{}, false
	}
	return o.Pnl, true
}

// HasPnl returns a boolean if a field has been set.
func (o *QueryPnLResponse) HasPnl() bool {
	if o != nil && !common.IsNil(o.Pnl) {
		return true
	}

	return false
}

// SetPnl gets a reference to the given map[string]interface{} and assigns it to the Pnl field.
func (o *QueryPnLResponse) SetPnl(v map[string]interface{}) {
	o.Pnl = v
}

// GetPnlList returns the PnlList field value if set, zero value otherwise.
func (o *QueryPnLResponse) GetPnlList() []GetPortfolioResponsePositionsInner {
	if o == nil || common.IsNil(o.PnlList) {
		var ret []GetPortfolioResponsePositionsInner
		return ret
	}
	return o.PnlList
}

// GetPnlListOk returns a tuple with the PnlList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPnLResponse) GetPnlListOk() ([]GetPortfolioResponsePositionsInner, bool) {
	if o == nil || common.IsNil(o.PnlList) {
		return nil, false
	}
	return o.PnlList, true
}

// HasPnlList returns a boolean if a field has been set.
func (o *QueryPnLResponse) HasPnlList() bool {
	if o != nil && !common.IsNil(o.PnlList) {
		return true
	}

	return false
}

// SetPnlList gets a reference to the given []GetPortfolioResponsePositionsInner and assigns it to the PnlList field.
func (o *QueryPnLResponse) SetPnlList(v []GetPortfolioResponsePositionsInner) {
	o.PnlList = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *QueryPnLResponse) GetTotalCount() int32 {
	if o == nil || common.IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPnLResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || common.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *QueryPnLResponse) HasTotalCount() bool {
	if o != nil && !common.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *QueryPnLResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetTotalRealizedPnl returns the TotalRealizedPnl field value if set, zero value otherwise.
func (o *QueryPnLResponse) GetTotalRealizedPnl() string {
	if o == nil || common.IsNil(o.TotalRealizedPnl) {
		var ret string
		return ret
	}
	return *o.TotalRealizedPnl
}

// GetTotalRealizedPnlOk returns a tuple with the TotalRealizedPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPnLResponse) GetTotalRealizedPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalRealizedPnl) {
		return nil, false
	}
	return o.TotalRealizedPnl, true
}

// HasTotalRealizedPnl returns a boolean if a field has been set.
func (o *QueryPnLResponse) HasTotalRealizedPnl() bool {
	if o != nil && !common.IsNil(o.TotalRealizedPnl) {
		return true
	}

	return false
}

// SetTotalRealizedPnl gets a reference to the given string and assigns it to the TotalRealizedPnl field.
func (o *QueryPnLResponse) SetTotalRealizedPnl(v string) {
	o.TotalRealizedPnl = &v
}

// GetTotalUnrealizedPnl returns the TotalUnrealizedPnl field value if set, zero value otherwise.
func (o *QueryPnLResponse) GetTotalUnrealizedPnl() string {
	if o == nil || common.IsNil(o.TotalUnrealizedPnl) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedPnl
}

// GetTotalUnrealizedPnlOk returns a tuple with the TotalUnrealizedPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPnLResponse) GetTotalUnrealizedPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedPnl) {
		return nil, false
	}
	return o.TotalUnrealizedPnl, true
}

// HasTotalUnrealizedPnl returns a boolean if a field has been set.
func (o *QueryPnLResponse) HasTotalUnrealizedPnl() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedPnl) {
		return true
	}

	return false
}

// SetTotalUnrealizedPnl gets a reference to the given string and assigns it to the TotalUnrealizedPnl field.
func (o *QueryPnLResponse) SetTotalUnrealizedPnl(v string) {
	o.TotalUnrealizedPnl = &v
}

// GetTotalPnl returns the TotalPnl field value if set, zero value otherwise.
func (o *QueryPnLResponse) GetTotalPnl() string {
	if o == nil || common.IsNil(o.TotalPnl) {
		var ret string
		return ret
	}
	return *o.TotalPnl
}

// GetTotalPnlOk returns a tuple with the TotalPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPnLResponse) GetTotalPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPnl) {
		return nil, false
	}
	return o.TotalPnl, true
}

// HasTotalPnl returns a boolean if a field has been set.
func (o *QueryPnLResponse) HasTotalPnl() bool {
	if o != nil && !common.IsNil(o.TotalPnl) {
		return true
	}

	return false
}

// SetTotalPnl gets a reference to the given string and assigns it to the TotalPnl field.
func (o *QueryPnLResponse) SetTotalPnl(v string) {
	o.TotalPnl = &v
}

func (o QueryPnLResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPnLResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	if !common.IsNil(o.WalletAddress) {
		toSerialize["walletAddress"] = o.WalletAddress
	}
	if o.Pnl != nil {
		toSerialize["pnl"] = o.Pnl
	}
	if !common.IsNil(o.PnlList) {
		toSerialize["pnlList"] = o.PnlList
	}
	if !common.IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !common.IsNil(o.TotalRealizedPnl) {
		toSerialize["totalRealizedPnl"] = o.TotalRealizedPnl
	}
	if !common.IsNil(o.TotalUnrealizedPnl) {
		toSerialize["totalUnrealizedPnl"] = o.TotalUnrealizedPnl
	}
	if !common.IsNil(o.TotalPnl) {
		toSerialize["totalPnl"] = o.TotalPnl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryPnLResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryPnLResponse := _QueryPnLResponse{}

	err = json.Unmarshal(data, &varQueryPnLResponse)

	if err != nil {
		return err
	}

	*o = QueryPnLResponse(varQueryPnLResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chainId")
		delete(additionalProperties, "walletAddress")
		delete(additionalProperties, "pnl")
		delete(additionalProperties, "pnlList")
		delete(additionalProperties, "totalCount")
		delete(additionalProperties, "totalRealizedPnl")
		delete(additionalProperties, "totalUnrealizedPnl")
		delete(additionalProperties, "totalPnl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPnLResponse struct {
	value *QueryPnLResponse
	isSet bool
}

func (v NullableQueryPnLResponse) Get() *QueryPnLResponse {
	return v.value
}

func (v *NullableQueryPnLResponse) Set(val *QueryPnLResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPnLResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPnLResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPnLResponse(val *QueryPnLResponse) *NullableQueryPnLResponse {
	return &NullableQueryPnLResponse{value: val, isSet: true}
}

func (v NullableQueryPnLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPnLResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
