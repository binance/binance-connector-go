/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner{}

// GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner struct for GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner
type GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner struct {
	Asset                *string `json:"asset,omitempty"`
	AssetFullName        *string `json:"assetFullName,omitempty"`
	AmountFree           *string `json:"amountFree,omitempty"`
	ToBTC                *string `json:"toBTC,omitempty"`
	ToBNB                *string `json:"toBNB,omitempty"`
	ToBNBOffExchange     *string `json:"toBNBOffExchange,omitempty"`
	Exchange             *string `json:"exchange,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner

// NewGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner instantiates a new GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner() *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner {
	this := GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner{}
	return &this
}

// NewGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInnerWithDefaults instantiates a new GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInnerWithDefaults() *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner {
	this := GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAssetFullName returns the AssetFullName field value if set, zero value otherwise.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetAssetFullName() string {
	if o == nil || common.IsNil(o.AssetFullName) {
		var ret string
		return ret
	}
	return *o.AssetFullName
}

// GetAssetFullNameOk returns a tuple with the AssetFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetAssetFullNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssetFullName) {
		return nil, false
	}
	return o.AssetFullName, true
}

// HasAssetFullName returns a boolean if a field has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) HasAssetFullName() bool {
	if o != nil && !common.IsNil(o.AssetFullName) {
		return true
	}

	return false
}

// SetAssetFullName gets a reference to the given string and assigns it to the AssetFullName field.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) SetAssetFullName(v string) {
	o.AssetFullName = &v
}

// GetAmountFree returns the AmountFree field value if set, zero value otherwise.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetAmountFree() string {
	if o == nil || common.IsNil(o.AmountFree) {
		var ret string
		return ret
	}
	return *o.AmountFree
}

// GetAmountFreeOk returns a tuple with the AmountFree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetAmountFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmountFree) {
		return nil, false
	}
	return o.AmountFree, true
}

// HasAmountFree returns a boolean if a field has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) HasAmountFree() bool {
	if o != nil && !common.IsNil(o.AmountFree) {
		return true
	}

	return false
}

// SetAmountFree gets a reference to the given string and assigns it to the AmountFree field.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) SetAmountFree(v string) {
	o.AmountFree = &v
}

// GetToBTC returns the ToBTC field value if set, zero value otherwise.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetToBTC() string {
	if o == nil || common.IsNil(o.ToBTC) {
		var ret string
		return ret
	}
	return *o.ToBTC
}

// GetToBTCOk returns a tuple with the ToBTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetToBTCOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToBTC) {
		return nil, false
	}
	return o.ToBTC, true
}

// HasToBTC returns a boolean if a field has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) HasToBTC() bool {
	if o != nil && !common.IsNil(o.ToBTC) {
		return true
	}

	return false
}

// SetToBTC gets a reference to the given string and assigns it to the ToBTC field.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) SetToBTC(v string) {
	o.ToBTC = &v
}

// GetToBNB returns the ToBNB field value if set, zero value otherwise.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetToBNB() string {
	if o == nil || common.IsNil(o.ToBNB) {
		var ret string
		return ret
	}
	return *o.ToBNB
}

// GetToBNBOk returns a tuple with the ToBNB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetToBNBOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToBNB) {
		return nil, false
	}
	return o.ToBNB, true
}

// HasToBNB returns a boolean if a field has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) HasToBNB() bool {
	if o != nil && !common.IsNil(o.ToBNB) {
		return true
	}

	return false
}

// SetToBNB gets a reference to the given string and assigns it to the ToBNB field.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) SetToBNB(v string) {
	o.ToBNB = &v
}

// GetToBNBOffExchange returns the ToBNBOffExchange field value if set, zero value otherwise.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetToBNBOffExchange() string {
	if o == nil || common.IsNil(o.ToBNBOffExchange) {
		var ret string
		return ret
	}
	return *o.ToBNBOffExchange
}

// GetToBNBOffExchangeOk returns a tuple with the ToBNBOffExchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetToBNBOffExchangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToBNBOffExchange) {
		return nil, false
	}
	return o.ToBNBOffExchange, true
}

// HasToBNBOffExchange returns a boolean if a field has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) HasToBNBOffExchange() bool {
	if o != nil && !common.IsNil(o.ToBNBOffExchange) {
		return true
	}

	return false
}

// SetToBNBOffExchange gets a reference to the given string and assigns it to the ToBNBOffExchange field.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) SetToBNBOffExchange(v string) {
	o.ToBNBOffExchange = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetExchange() string {
	if o == nil || common.IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) GetExchangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) HasExchange() bool {
	if o != nil && !common.IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) SetExchange(v string) {
	o.Exchange = &v
}

func (o GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.AssetFullName) {
		toSerialize["assetFullName"] = o.AssetFullName
	}
	if !common.IsNil(o.AmountFree) {
		toSerialize["amountFree"] = o.AmountFree
	}
	if !common.IsNil(o.ToBTC) {
		toSerialize["toBTC"] = o.ToBTC
	}
	if !common.IsNil(o.ToBNB) {
		toSerialize["toBNB"] = o.ToBNB
	}
	if !common.IsNil(o.ToBNBOffExchange) {
		toSerialize["toBNBOffExchange"] = o.ToBNBOffExchange
	}
	if !common.IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) UnmarshalJSON(data []byte) (err error) {
	varGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner := _GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner{}

	err = json.Unmarshal(data, &varGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner)

	if err != nil {
		return err
	}

	*o = GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner(varGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "assetFullName")
		delete(additionalProperties, "amountFree")
		delete(additionalProperties, "toBTC")
		delete(additionalProperties, "toBNB")
		delete(additionalProperties, "toBNBOffExchange")
		delete(additionalProperties, "exchange")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner struct {
	value *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner
	isSet bool
}

func (v NullableGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) Get() *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner {
	return v.value
}

func (v *NullableGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) Set(val *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner(val *GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) *NullableGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner {
	return &NullableGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner{value: val, isSet: true}
}

func (v NullableGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
