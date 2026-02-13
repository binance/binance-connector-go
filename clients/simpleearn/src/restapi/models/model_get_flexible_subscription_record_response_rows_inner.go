/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFlexibleSubscriptionRecordResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleSubscriptionRecordResponseRowsInner{}

// GetFlexibleSubscriptionRecordResponseRowsInner struct for GetFlexibleSubscriptionRecordResponseRowsInner
type GetFlexibleSubscriptionRecordResponseRowsInner struct {
	Amount               *string `json:"amount,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	PurchaseId           *int64  `json:"purchaseId,omitempty"`
	ProductId            *string `json:"productId,omitempty"`
	Type                 *string `json:"type,omitempty"`
	SourceAccount        *string `json:"sourceAccount,omitempty"`
	AmtFromSpot          *string `json:"amtFromSpot,omitempty"`
	AmtFromFunding       *string `json:"amtFromFunding,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleSubscriptionRecordResponseRowsInner GetFlexibleSubscriptionRecordResponseRowsInner

// NewGetFlexibleSubscriptionRecordResponseRowsInner instantiates a new GetFlexibleSubscriptionRecordResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleSubscriptionRecordResponseRowsInner() *GetFlexibleSubscriptionRecordResponseRowsInner {
	this := GetFlexibleSubscriptionRecordResponseRowsInner{}
	return &this
}

// NewGetFlexibleSubscriptionRecordResponseRowsInnerWithDefaults instantiates a new GetFlexibleSubscriptionRecordResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleSubscriptionRecordResponseRowsInnerWithDefaults() *GetFlexibleSubscriptionRecordResponseRowsInner {
	this := GetFlexibleSubscriptionRecordResponseRowsInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetPurchaseId returns the PurchaseId field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetPurchaseId() int64 {
	if o == nil || common.IsNil(o.PurchaseId) {
		var ret int64
		return ret
	}
	return *o.PurchaseId
}

// GetPurchaseIdOk returns a tuple with the PurchaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetPurchaseIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PurchaseId) {
		return nil, false
	}
	return o.PurchaseId, true
}

// HasPurchaseId returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) HasPurchaseId() bool {
	if o != nil && !common.IsNil(o.PurchaseId) {
		return true
	}

	return false
}

// SetPurchaseId gets a reference to the given int64 and assigns it to the PurchaseId field.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) SetPurchaseId(v int64) {
	o.PurchaseId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetProductId() string {
	if o == nil || common.IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetProductIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) HasProductId() bool {
	if o != nil && !common.IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) SetProductId(v string) {
	o.ProductId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) SetType(v string) {
	o.Type = &v
}

// GetSourceAccount returns the SourceAccount field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetSourceAccount() string {
	if o == nil || common.IsNil(o.SourceAccount) {
		var ret string
		return ret
	}
	return *o.SourceAccount
}

// GetSourceAccountOk returns a tuple with the SourceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetSourceAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.SourceAccount) {
		return nil, false
	}
	return o.SourceAccount, true
}

// HasSourceAccount returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) HasSourceAccount() bool {
	if o != nil && !common.IsNil(o.SourceAccount) {
		return true
	}

	return false
}

// SetSourceAccount gets a reference to the given string and assigns it to the SourceAccount field.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) SetSourceAccount(v string) {
	o.SourceAccount = &v
}

// GetAmtFromSpot returns the AmtFromSpot field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetAmtFromSpot() string {
	if o == nil || common.IsNil(o.AmtFromSpot) {
		var ret string
		return ret
	}
	return *o.AmtFromSpot
}

// GetAmtFromSpotOk returns a tuple with the AmtFromSpot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetAmtFromSpotOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmtFromSpot) {
		return nil, false
	}
	return o.AmtFromSpot, true
}

// HasAmtFromSpot returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) HasAmtFromSpot() bool {
	if o != nil && !common.IsNil(o.AmtFromSpot) {
		return true
	}

	return false
}

// SetAmtFromSpot gets a reference to the given string and assigns it to the AmtFromSpot field.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) SetAmtFromSpot(v string) {
	o.AmtFromSpot = &v
}

// GetAmtFromFunding returns the AmtFromFunding field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetAmtFromFunding() string {
	if o == nil || common.IsNil(o.AmtFromFunding) {
		var ret string
		return ret
	}
	return *o.AmtFromFunding
}

// GetAmtFromFundingOk returns a tuple with the AmtFromFunding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetAmtFromFundingOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmtFromFunding) {
		return nil, false
	}
	return o.AmtFromFunding, true
}

// HasAmtFromFunding returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) HasAmtFromFunding() bool {
	if o != nil && !common.IsNil(o.AmtFromFunding) {
		return true
	}

	return false
}

// SetAmtFromFunding gets a reference to the given string and assigns it to the AmtFromFunding field.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) SetAmtFromFunding(v string) {
	o.AmtFromFunding = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetFlexibleSubscriptionRecordResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetFlexibleSubscriptionRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleSubscriptionRecordResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.PurchaseId) {
		toSerialize["purchaseId"] = o.PurchaseId
	}
	if !common.IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.SourceAccount) {
		toSerialize["sourceAccount"] = o.SourceAccount
	}
	if !common.IsNil(o.AmtFromSpot) {
		toSerialize["amtFromSpot"] = o.AmtFromSpot
	}
	if !common.IsNil(o.AmtFromFunding) {
		toSerialize["amtFromFunding"] = o.AmtFromFunding
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleSubscriptionRecordResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleSubscriptionRecordResponseRowsInner := _GetFlexibleSubscriptionRecordResponseRowsInner{}

	err = json.Unmarshal(data, &varGetFlexibleSubscriptionRecordResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetFlexibleSubscriptionRecordResponseRowsInner(varGetFlexibleSubscriptionRecordResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "time")
		delete(additionalProperties, "purchaseId")
		delete(additionalProperties, "productId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "sourceAccount")
		delete(additionalProperties, "amtFromSpot")
		delete(additionalProperties, "amtFromFunding")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleSubscriptionRecordResponseRowsInner struct {
	value *GetFlexibleSubscriptionRecordResponseRowsInner
	isSet bool
}

func (v NullableGetFlexibleSubscriptionRecordResponseRowsInner) Get() *GetFlexibleSubscriptionRecordResponseRowsInner {
	return v.value
}

func (v *NullableGetFlexibleSubscriptionRecordResponseRowsInner) Set(val *GetFlexibleSubscriptionRecordResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleSubscriptionRecordResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleSubscriptionRecordResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleSubscriptionRecordResponseRowsInner(val *GetFlexibleSubscriptionRecordResponseRowsInner) *NullableGetFlexibleSubscriptionRecordResponseRowsInner {
	return &NullableGetFlexibleSubscriptionRecordResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetFlexibleSubscriptionRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleSubscriptionRecordResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
