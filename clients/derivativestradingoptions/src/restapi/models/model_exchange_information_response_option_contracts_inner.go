/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ExchangeInformationResponseOptionContractsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInformationResponseOptionContractsInner{}

// ExchangeInformationResponseOptionContractsInner struct for ExchangeInformationResponseOptionContractsInner
type ExchangeInformationResponseOptionContractsInner struct {
	BaseAsset            *string `json:"baseAsset,omitempty"`
	QuoteAsset           *string `json:"quoteAsset,omitempty"`
	Underlying           *string `json:"underlying,omitempty"`
	SettleAsset          *string `json:"settleAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeInformationResponseOptionContractsInner ExchangeInformationResponseOptionContractsInner

// NewExchangeInformationResponseOptionContractsInner instantiates a new ExchangeInformationResponseOptionContractsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInformationResponseOptionContractsInner() *ExchangeInformationResponseOptionContractsInner {
	this := ExchangeInformationResponseOptionContractsInner{}
	return &this
}

// NewExchangeInformationResponseOptionContractsInnerWithDefaults instantiates a new ExchangeInformationResponseOptionContractsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInformationResponseOptionContractsInnerWithDefaults() *ExchangeInformationResponseOptionContractsInner {
	this := ExchangeInformationResponseOptionContractsInner{}
	return &this
}

// GetBaseAsset returns the BaseAsset field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionContractsInner) GetBaseAsset() string {
	if o == nil || common.IsNil(o.BaseAsset) {
		var ret string
		return ret
	}
	return *o.BaseAsset
}

// GetBaseAssetOk returns a tuple with the BaseAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionContractsInner) GetBaseAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BaseAsset) {
		return nil, false
	}
	return o.BaseAsset, true
}

// HasBaseAsset returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionContractsInner) HasBaseAsset() bool {
	if o != nil && !common.IsNil(o.BaseAsset) {
		return true
	}

	return false
}

// SetBaseAsset gets a reference to the given string and assigns it to the BaseAsset field.
func (o *ExchangeInformationResponseOptionContractsInner) SetBaseAsset(v string) {
	o.BaseAsset = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionContractsInner) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionContractsInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionContractsInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *ExchangeInformationResponseOptionContractsInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionContractsInner) GetUnderlying() string {
	if o == nil || common.IsNil(o.Underlying) {
		var ret string
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionContractsInner) GetUnderlyingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionContractsInner) HasUnderlying() bool {
	if o != nil && !common.IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given string and assigns it to the Underlying field.
func (o *ExchangeInformationResponseOptionContractsInner) SetUnderlying(v string) {
	o.Underlying = &v
}

// GetSettleAsset returns the SettleAsset field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionContractsInner) GetSettleAsset() string {
	if o == nil || common.IsNil(o.SettleAsset) {
		var ret string
		return ret
	}
	return *o.SettleAsset
}

// GetSettleAssetOk returns a tuple with the SettleAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionContractsInner) GetSettleAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.SettleAsset) {
		return nil, false
	}
	return o.SettleAsset, true
}

// HasSettleAsset returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionContractsInner) HasSettleAsset() bool {
	if o != nil && !common.IsNil(o.SettleAsset) {
		return true
	}

	return false
}

// SetSettleAsset gets a reference to the given string and assigns it to the SettleAsset field.
func (o *ExchangeInformationResponseOptionContractsInner) SetSettleAsset(v string) {
	o.SettleAsset = &v
}

func (o ExchangeInformationResponseOptionContractsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInformationResponseOptionContractsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BaseAsset) {
		toSerialize["baseAsset"] = o.BaseAsset
	}
	if !common.IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}
	if !common.IsNil(o.Underlying) {
		toSerialize["underlying"] = o.Underlying
	}
	if !common.IsNil(o.SettleAsset) {
		toSerialize["settleAsset"] = o.SettleAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeInformationResponseOptionContractsInner) UnmarshalJSON(data []byte) (err error) {
	varExchangeInformationResponseOptionContractsInner := _ExchangeInformationResponseOptionContractsInner{}

	err = json.Unmarshal(data, &varExchangeInformationResponseOptionContractsInner)

	if err != nil {
		return err
	}

	*o = ExchangeInformationResponseOptionContractsInner(varExchangeInformationResponseOptionContractsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "baseAsset")
		delete(additionalProperties, "quoteAsset")
		delete(additionalProperties, "underlying")
		delete(additionalProperties, "settleAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInformationResponseOptionContractsInner struct {
	value *ExchangeInformationResponseOptionContractsInner
	isSet bool
}

func (v NullableExchangeInformationResponseOptionContractsInner) Get() *ExchangeInformationResponseOptionContractsInner {
	return v.value
}

func (v *NullableExchangeInformationResponseOptionContractsInner) Set(val *ExchangeInformationResponseOptionContractsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInformationResponseOptionContractsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInformationResponseOptionContractsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInformationResponseOptionContractsInner(val *ExchangeInformationResponseOptionContractsInner) *NullableExchangeInformationResponseOptionContractsInner {
	return &NullableExchangeInformationResponseOptionContractsInner{value: val, isSet: true}
}

func (v NullableExchangeInformationResponseOptionContractsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInformationResponseOptionContractsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
