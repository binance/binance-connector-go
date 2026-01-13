/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CompositeIndexSymbolInformationResponseInnerBaseAssetListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CompositeIndexSymbolInformationResponseInnerBaseAssetListInner{}

// CompositeIndexSymbolInformationResponseInnerBaseAssetListInner struct for CompositeIndexSymbolInformationResponseInnerBaseAssetListInner
type CompositeIndexSymbolInformationResponseInnerBaseAssetListInner struct {
	BaseAsset            *string `json:"baseAsset,omitempty"`
	QuoteAsset           *string `json:"quoteAsset,omitempty"`
	WeightInQuantity     *string `json:"weightInQuantity,omitempty"`
	WeightInPercentage   *string `json:"weightInPercentage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompositeIndexSymbolInformationResponseInnerBaseAssetListInner CompositeIndexSymbolInformationResponseInnerBaseAssetListInner

// NewCompositeIndexSymbolInformationResponseInnerBaseAssetListInner instantiates a new CompositeIndexSymbolInformationResponseInnerBaseAssetListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompositeIndexSymbolInformationResponseInnerBaseAssetListInner() *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner {
	this := CompositeIndexSymbolInformationResponseInnerBaseAssetListInner{}
	return &this
}

// NewCompositeIndexSymbolInformationResponseInnerBaseAssetListInnerWithDefaults instantiates a new CompositeIndexSymbolInformationResponseInnerBaseAssetListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompositeIndexSymbolInformationResponseInnerBaseAssetListInnerWithDefaults() *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner {
	this := CompositeIndexSymbolInformationResponseInnerBaseAssetListInner{}
	return &this
}

// GetBaseAsset returns the BaseAsset field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) GetBaseAsset() string {
	if o == nil || common.IsNil(o.BaseAsset) {
		var ret string
		return ret
	}
	return *o.BaseAsset
}

// GetBaseAssetOk returns a tuple with the BaseAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) GetBaseAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BaseAsset) {
		return nil, false
	}
	return o.BaseAsset, true
}

// HasBaseAsset returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) HasBaseAsset() bool {
	if o != nil && !common.IsNil(o.BaseAsset) {
		return true
	}

	return false
}

// SetBaseAsset gets a reference to the given string and assigns it to the BaseAsset field.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) SetBaseAsset(v string) {
	o.BaseAsset = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetWeightInQuantity returns the WeightInQuantity field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) GetWeightInQuantity() string {
	if o == nil || common.IsNil(o.WeightInQuantity) {
		var ret string
		return ret
	}
	return *o.WeightInQuantity
}

// GetWeightInQuantityOk returns a tuple with the WeightInQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) GetWeightInQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.WeightInQuantity) {
		return nil, false
	}
	return o.WeightInQuantity, true
}

// HasWeightInQuantity returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) HasWeightInQuantity() bool {
	if o != nil && !common.IsNil(o.WeightInQuantity) {
		return true
	}

	return false
}

// SetWeightInQuantity gets a reference to the given string and assigns it to the WeightInQuantity field.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) SetWeightInQuantity(v string) {
	o.WeightInQuantity = &v
}

// GetWeightInPercentage returns the WeightInPercentage field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) GetWeightInPercentage() string {
	if o == nil || common.IsNil(o.WeightInPercentage) {
		var ret string
		return ret
	}
	return *o.WeightInPercentage
}

// GetWeightInPercentageOk returns a tuple with the WeightInPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) GetWeightInPercentageOk() (*string, bool) {
	if o == nil || common.IsNil(o.WeightInPercentage) {
		return nil, false
	}
	return o.WeightInPercentage, true
}

// HasWeightInPercentage returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) HasWeightInPercentage() bool {
	if o != nil && !common.IsNil(o.WeightInPercentage) {
		return true
	}

	return false
}

// SetWeightInPercentage gets a reference to the given string and assigns it to the WeightInPercentage field.
func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) SetWeightInPercentage(v string) {
	o.WeightInPercentage = &v
}

func (o CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BaseAsset) {
		toSerialize["baseAsset"] = o.BaseAsset
	}
	if !common.IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}
	if !common.IsNil(o.WeightInQuantity) {
		toSerialize["weightInQuantity"] = o.WeightInQuantity
	}
	if !common.IsNil(o.WeightInPercentage) {
		toSerialize["weightInPercentage"] = o.WeightInPercentage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) UnmarshalJSON(data []byte) (err error) {
	varCompositeIndexSymbolInformationResponseInnerBaseAssetListInner := _CompositeIndexSymbolInformationResponseInnerBaseAssetListInner{}

	err = json.Unmarshal(data, &varCompositeIndexSymbolInformationResponseInnerBaseAssetListInner)

	if err != nil {
		return err
	}

	*o = CompositeIndexSymbolInformationResponseInnerBaseAssetListInner(varCompositeIndexSymbolInformationResponseInnerBaseAssetListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "baseAsset")
		delete(additionalProperties, "quoteAsset")
		delete(additionalProperties, "weightInQuantity")
		delete(additionalProperties, "weightInPercentage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompositeIndexSymbolInformationResponseInnerBaseAssetListInner struct {
	value *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner
	isSet bool
}

func (v NullableCompositeIndexSymbolInformationResponseInnerBaseAssetListInner) Get() *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner {
	return v.value
}

func (v *NullableCompositeIndexSymbolInformationResponseInnerBaseAssetListInner) Set(val *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCompositeIndexSymbolInformationResponseInnerBaseAssetListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCompositeIndexSymbolInformationResponseInnerBaseAssetListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompositeIndexSymbolInformationResponseInnerBaseAssetListInner(val *CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) *NullableCompositeIndexSymbolInformationResponseInnerBaseAssetListInner {
	return &NullableCompositeIndexSymbolInformationResponseInnerBaseAssetListInner{value: val, isSet: true}
}

func (v NullableCompositeIndexSymbolInformationResponseInnerBaseAssetListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompositeIndexSymbolInformationResponseInnerBaseAssetListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
