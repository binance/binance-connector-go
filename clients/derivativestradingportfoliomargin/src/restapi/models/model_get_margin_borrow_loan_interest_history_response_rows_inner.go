/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetMarginBorrowLoanInterestHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMarginBorrowLoanInterestHistoryResponseRowsInner{}

// GetMarginBorrowLoanInterestHistoryResponseRowsInner struct for GetMarginBorrowLoanInterestHistoryResponseRowsInner
type GetMarginBorrowLoanInterestHistoryResponseRowsInner struct {
	TxId                 *int64  `json:"txId,omitempty"`
	InterestAccuredTime  *int64  `json:"interestAccuredTime,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	RawAsset             *string `json:"rawAsset,omitempty"`
	Principal            *string `json:"principal,omitempty"`
	Interest             *string `json:"interest,omitempty"`
	InterestRate         *string `json:"interestRate,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMarginBorrowLoanInterestHistoryResponseRowsInner GetMarginBorrowLoanInterestHistoryResponseRowsInner

// NewGetMarginBorrowLoanInterestHistoryResponseRowsInner instantiates a new GetMarginBorrowLoanInterestHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarginBorrowLoanInterestHistoryResponseRowsInner() *GetMarginBorrowLoanInterestHistoryResponseRowsInner {
	this := GetMarginBorrowLoanInterestHistoryResponseRowsInner{}
	return &this
}

// NewGetMarginBorrowLoanInterestHistoryResponseRowsInnerWithDefaults instantiates a new GetMarginBorrowLoanInterestHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarginBorrowLoanInterestHistoryResponseRowsInnerWithDefaults() *GetMarginBorrowLoanInterestHistoryResponseRowsInner {
	this := GetMarginBorrowLoanInterestHistoryResponseRowsInner{}
	return &this
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetTxId() int64 {
	if o == nil || common.IsNil(o.TxId) {
		var ret int64
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetTxIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) HasTxId() bool {
	if o != nil && !common.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int64 and assigns it to the TxId field.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) SetTxId(v int64) {
	o.TxId = &v
}

// GetInterestAccuredTime returns the InterestAccuredTime field value if set, zero value otherwise.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetInterestAccuredTime() int64 {
	if o == nil || common.IsNil(o.InterestAccuredTime) {
		var ret int64
		return ret
	}
	return *o.InterestAccuredTime
}

// GetInterestAccuredTimeOk returns a tuple with the InterestAccuredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetInterestAccuredTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InterestAccuredTime) {
		return nil, false
	}
	return o.InterestAccuredTime, true
}

// HasInterestAccuredTime returns a boolean if a field has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) HasInterestAccuredTime() bool {
	if o != nil && !common.IsNil(o.InterestAccuredTime) {
		return true
	}

	return false
}

// SetInterestAccuredTime gets a reference to the given int64 and assigns it to the InterestAccuredTime field.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) SetInterestAccuredTime(v int64) {
	o.InterestAccuredTime = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetRawAsset returns the RawAsset field value if set, zero value otherwise.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetRawAsset() string {
	if o == nil || common.IsNil(o.RawAsset) {
		var ret string
		return ret
	}
	return *o.RawAsset
}

// GetRawAssetOk returns a tuple with the RawAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetRawAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RawAsset) {
		return nil, false
	}
	return o.RawAsset, true
}

// HasRawAsset returns a boolean if a field has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) HasRawAsset() bool {
	if o != nil && !common.IsNil(o.RawAsset) {
		return true
	}

	return false
}

// SetRawAsset gets a reference to the given string and assigns it to the RawAsset field.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) SetRawAsset(v string) {
	o.RawAsset = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetPrincipal() string {
	if o == nil || common.IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetPrincipalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) HasPrincipal() bool {
	if o != nil && !common.IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) SetPrincipal(v string) {
	o.Principal = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetInterest() string {
	if o == nil || common.IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) HasInterest() bool {
	if o != nil && !common.IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) SetInterest(v string) {
	o.Interest = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetInterestRate() string {
	if o == nil || common.IsNil(o.InterestRate) {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) HasInterestRate() bool {
	if o != nil && !common.IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) SetInterestRate(v string) {
	o.InterestRate = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) SetType(v string) {
	o.Type = &v
}

func (o GetMarginBorrowLoanInterestHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarginBorrowLoanInterestHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TxId) {
		toSerialize["txId"] = o.TxId
	}
	if !common.IsNil(o.InterestAccuredTime) {
		toSerialize["interestAccuredTime"] = o.InterestAccuredTime
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.RawAsset) {
		toSerialize["rawAsset"] = o.RawAsset
	}
	if !common.IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !common.IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !common.IsNil(o.InterestRate) {
		toSerialize["interestRate"] = o.InterestRate
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMarginBorrowLoanInterestHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetMarginBorrowLoanInterestHistoryResponseRowsInner := _GetMarginBorrowLoanInterestHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetMarginBorrowLoanInterestHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetMarginBorrowLoanInterestHistoryResponseRowsInner(varGetMarginBorrowLoanInterestHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "txId")
		delete(additionalProperties, "interestAccuredTime")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "rawAsset")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "interest")
		delete(additionalProperties, "interestRate")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMarginBorrowLoanInterestHistoryResponseRowsInner struct {
	value *GetMarginBorrowLoanInterestHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetMarginBorrowLoanInterestHistoryResponseRowsInner) Get() *GetMarginBorrowLoanInterestHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetMarginBorrowLoanInterestHistoryResponseRowsInner) Set(val *GetMarginBorrowLoanInterestHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarginBorrowLoanInterestHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarginBorrowLoanInterestHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarginBorrowLoanInterestHistoryResponseRowsInner(val *GetMarginBorrowLoanInterestHistoryResponseRowsInner) *NullableGetMarginBorrowLoanInterestHistoryResponseRowsInner {
	return &NullableGetMarginBorrowLoanInterestHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetMarginBorrowLoanInterestHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarginBorrowLoanInterestHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
