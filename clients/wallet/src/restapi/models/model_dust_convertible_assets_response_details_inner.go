/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DustConvertibleAssetsResponseDetailsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DustConvertibleAssetsResponseDetailsInner{}

// DustConvertibleAssetsResponseDetailsInner struct for DustConvertibleAssetsResponseDetailsInner
type DustConvertibleAssetsResponseDetailsInner struct {
	Asset                    *string `json:"asset,omitempty"`
	AssetFullName            *string `json:"assetFullName,omitempty"`
	AmountFree               *string `json:"amountFree,omitempty"`
	Exchange                 *string `json:"exchange,omitempty"`
	ToQuotaAssetAmount       *string `json:"toQuotaAssetAmount,omitempty"`
	ToTargetAssetAmount      *string `json:"toTargetAssetAmount,omitempty"`
	ToTargetAssetOffExchange *string `json:"toTargetAssetOffExchange,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _DustConvertibleAssetsResponseDetailsInner DustConvertibleAssetsResponseDetailsInner

// NewDustConvertibleAssetsResponseDetailsInner instantiates a new DustConvertibleAssetsResponseDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDustConvertibleAssetsResponseDetailsInner() *DustConvertibleAssetsResponseDetailsInner {
	this := DustConvertibleAssetsResponseDetailsInner{}
	return &this
}

// NewDustConvertibleAssetsResponseDetailsInnerWithDefaults instantiates a new DustConvertibleAssetsResponseDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDustConvertibleAssetsResponseDetailsInnerWithDefaults() *DustConvertibleAssetsResponseDetailsInner {
	this := DustConvertibleAssetsResponseDetailsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *DustConvertibleAssetsResponseDetailsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *DustConvertibleAssetsResponseDetailsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAssetFullName returns the AssetFullName field value if set, zero value otherwise.
func (o *DustConvertibleAssetsResponseDetailsInner) GetAssetFullName() string {
	if o == nil || common.IsNil(o.AssetFullName) {
		var ret string
		return ret
	}
	return *o.AssetFullName
}

// GetAssetFullNameOk returns a tuple with the AssetFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) GetAssetFullNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssetFullName) {
		return nil, false
	}
	return o.AssetFullName, true
}

// HasAssetFullName returns a boolean if a field has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) HasAssetFullName() bool {
	if o != nil && !common.IsNil(o.AssetFullName) {
		return true
	}

	return false
}

// SetAssetFullName gets a reference to the given string and assigns it to the AssetFullName field.
func (o *DustConvertibleAssetsResponseDetailsInner) SetAssetFullName(v string) {
	o.AssetFullName = &v
}

// GetAmountFree returns the AmountFree field value if set, zero value otherwise.
func (o *DustConvertibleAssetsResponseDetailsInner) GetAmountFree() string {
	if o == nil || common.IsNil(o.AmountFree) {
		var ret string
		return ret
	}
	return *o.AmountFree
}

// GetAmountFreeOk returns a tuple with the AmountFree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) GetAmountFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmountFree) {
		return nil, false
	}
	return o.AmountFree, true
}

// HasAmountFree returns a boolean if a field has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) HasAmountFree() bool {
	if o != nil && !common.IsNil(o.AmountFree) {
		return true
	}

	return false
}

// SetAmountFree gets a reference to the given string and assigns it to the AmountFree field.
func (o *DustConvertibleAssetsResponseDetailsInner) SetAmountFree(v string) {
	o.AmountFree = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *DustConvertibleAssetsResponseDetailsInner) GetExchange() string {
	if o == nil || common.IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) GetExchangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) HasExchange() bool {
	if o != nil && !common.IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *DustConvertibleAssetsResponseDetailsInner) SetExchange(v string) {
	o.Exchange = &v
}

// GetToQuotaAssetAmount returns the ToQuotaAssetAmount field value if set, zero value otherwise.
func (o *DustConvertibleAssetsResponseDetailsInner) GetToQuotaAssetAmount() string {
	if o == nil || common.IsNil(o.ToQuotaAssetAmount) {
		var ret string
		return ret
	}
	return *o.ToQuotaAssetAmount
}

// GetToQuotaAssetAmountOk returns a tuple with the ToQuotaAssetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) GetToQuotaAssetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToQuotaAssetAmount) {
		return nil, false
	}
	return o.ToQuotaAssetAmount, true
}

// HasToQuotaAssetAmount returns a boolean if a field has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) HasToQuotaAssetAmount() bool {
	if o != nil && !common.IsNil(o.ToQuotaAssetAmount) {
		return true
	}

	return false
}

// SetToQuotaAssetAmount gets a reference to the given string and assigns it to the ToQuotaAssetAmount field.
func (o *DustConvertibleAssetsResponseDetailsInner) SetToQuotaAssetAmount(v string) {
	o.ToQuotaAssetAmount = &v
}

// GetToTargetAssetAmount returns the ToTargetAssetAmount field value if set, zero value otherwise.
func (o *DustConvertibleAssetsResponseDetailsInner) GetToTargetAssetAmount() string {
	if o == nil || common.IsNil(o.ToTargetAssetAmount) {
		var ret string
		return ret
	}
	return *o.ToTargetAssetAmount
}

// GetToTargetAssetAmountOk returns a tuple with the ToTargetAssetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) GetToTargetAssetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToTargetAssetAmount) {
		return nil, false
	}
	return o.ToTargetAssetAmount, true
}

// HasToTargetAssetAmount returns a boolean if a field has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) HasToTargetAssetAmount() bool {
	if o != nil && !common.IsNil(o.ToTargetAssetAmount) {
		return true
	}

	return false
}

// SetToTargetAssetAmount gets a reference to the given string and assigns it to the ToTargetAssetAmount field.
func (o *DustConvertibleAssetsResponseDetailsInner) SetToTargetAssetAmount(v string) {
	o.ToTargetAssetAmount = &v
}

// GetToTargetAssetOffExchange returns the ToTargetAssetOffExchange field value if set, zero value otherwise.
func (o *DustConvertibleAssetsResponseDetailsInner) GetToTargetAssetOffExchange() string {
	if o == nil || common.IsNil(o.ToTargetAssetOffExchange) {
		var ret string
		return ret
	}
	return *o.ToTargetAssetOffExchange
}

// GetToTargetAssetOffExchangeOk returns a tuple with the ToTargetAssetOffExchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) GetToTargetAssetOffExchangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToTargetAssetOffExchange) {
		return nil, false
	}
	return o.ToTargetAssetOffExchange, true
}

// HasToTargetAssetOffExchange returns a boolean if a field has been set.
func (o *DustConvertibleAssetsResponseDetailsInner) HasToTargetAssetOffExchange() bool {
	if o != nil && !common.IsNil(o.ToTargetAssetOffExchange) {
		return true
	}

	return false
}

// SetToTargetAssetOffExchange gets a reference to the given string and assigns it to the ToTargetAssetOffExchange field.
func (o *DustConvertibleAssetsResponseDetailsInner) SetToTargetAssetOffExchange(v string) {
	o.ToTargetAssetOffExchange = &v
}

func (o DustConvertibleAssetsResponseDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DustConvertibleAssetsResponseDetailsInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !common.IsNil(o.ToQuotaAssetAmount) {
		toSerialize["toQuotaAssetAmount"] = o.ToQuotaAssetAmount
	}
	if !common.IsNil(o.ToTargetAssetAmount) {
		toSerialize["toTargetAssetAmount"] = o.ToTargetAssetAmount
	}
	if !common.IsNil(o.ToTargetAssetOffExchange) {
		toSerialize["toTargetAssetOffExchange"] = o.ToTargetAssetOffExchange
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DustConvertibleAssetsResponseDetailsInner) UnmarshalJSON(data []byte) (err error) {
	varDustConvertibleAssetsResponseDetailsInner := _DustConvertibleAssetsResponseDetailsInner{}

	err = json.Unmarshal(data, &varDustConvertibleAssetsResponseDetailsInner)

	if err != nil {
		return err
	}

	*o = DustConvertibleAssetsResponseDetailsInner(varDustConvertibleAssetsResponseDetailsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "assetFullName")
		delete(additionalProperties, "amountFree")
		delete(additionalProperties, "exchange")
		delete(additionalProperties, "toQuotaAssetAmount")
		delete(additionalProperties, "toTargetAssetAmount")
		delete(additionalProperties, "toTargetAssetOffExchange")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDustConvertibleAssetsResponseDetailsInner struct {
	value *DustConvertibleAssetsResponseDetailsInner
	isSet bool
}

func (v NullableDustConvertibleAssetsResponseDetailsInner) Get() *DustConvertibleAssetsResponseDetailsInner {
	return v.value
}

func (v *NullableDustConvertibleAssetsResponseDetailsInner) Set(val *DustConvertibleAssetsResponseDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDustConvertibleAssetsResponseDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDustConvertibleAssetsResponseDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDustConvertibleAssetsResponseDetailsInner(val *DustConvertibleAssetsResponseDetailsInner) *NullableDustConvertibleAssetsResponseDetailsInner {
	return &NullableDustConvertibleAssetsResponseDetailsInner{value: val, isSet: true}
}

func (v NullableDustConvertibleAssetsResponseDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDustConvertibleAssetsResponseDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
