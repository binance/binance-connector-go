/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TopTraderLongShortRatioPositionsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TopTraderLongShortRatioPositionsResponseInner{}

// TopTraderLongShortRatioPositionsResponseInner struct for TopTraderLongShortRatioPositionsResponseInner
type TopTraderLongShortRatioPositionsResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	LongShortRatio       *string `json:"longShortRatio,omitempty"`
	LongAccount          *string `json:"longAccount,omitempty"`
	ShortAccount         *string `json:"shortAccount,omitempty"`
	Timestamp            *string `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TopTraderLongShortRatioPositionsResponseInner TopTraderLongShortRatioPositionsResponseInner

// NewTopTraderLongShortRatioPositionsResponseInner instantiates a new TopTraderLongShortRatioPositionsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopTraderLongShortRatioPositionsResponseInner() *TopTraderLongShortRatioPositionsResponseInner {
	this := TopTraderLongShortRatioPositionsResponseInner{}
	return &this
}

// NewTopTraderLongShortRatioPositionsResponseInnerWithDefaults instantiates a new TopTraderLongShortRatioPositionsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopTraderLongShortRatioPositionsResponseInnerWithDefaults() *TopTraderLongShortRatioPositionsResponseInner {
	this := TopTraderLongShortRatioPositionsResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TopTraderLongShortRatioPositionsResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetLongShortRatio returns the LongShortRatio field value if set, zero value otherwise.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetLongShortRatio() string {
	if o == nil || common.IsNil(o.LongShortRatio) {
		var ret string
		return ret
	}
	return *o.LongShortRatio
}

// GetLongShortRatioOk returns a tuple with the LongShortRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetLongShortRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.LongShortRatio) {
		return nil, false
	}
	return o.LongShortRatio, true
}

// HasLongShortRatio returns a boolean if a field has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) HasLongShortRatio() bool {
	if o != nil && !common.IsNil(o.LongShortRatio) {
		return true
	}

	return false
}

// SetLongShortRatio gets a reference to the given string and assigns it to the LongShortRatio field.
func (o *TopTraderLongShortRatioPositionsResponseInner) SetLongShortRatio(v string) {
	o.LongShortRatio = &v
}

// GetLongAccount returns the LongAccount field value if set, zero value otherwise.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetLongAccount() string {
	if o == nil || common.IsNil(o.LongAccount) {
		var ret string
		return ret
	}
	return *o.LongAccount
}

// GetLongAccountOk returns a tuple with the LongAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetLongAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.LongAccount) {
		return nil, false
	}
	return o.LongAccount, true
}

// HasLongAccount returns a boolean if a field has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) HasLongAccount() bool {
	if o != nil && !common.IsNil(o.LongAccount) {
		return true
	}

	return false
}

// SetLongAccount gets a reference to the given string and assigns it to the LongAccount field.
func (o *TopTraderLongShortRatioPositionsResponseInner) SetLongAccount(v string) {
	o.LongAccount = &v
}

// GetShortAccount returns the ShortAccount field value if set, zero value otherwise.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetShortAccount() string {
	if o == nil || common.IsNil(o.ShortAccount) {
		var ret string
		return ret
	}
	return *o.ShortAccount
}

// GetShortAccountOk returns a tuple with the ShortAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetShortAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShortAccount) {
		return nil, false
	}
	return o.ShortAccount, true
}

// HasShortAccount returns a boolean if a field has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) HasShortAccount() bool {
	if o != nil && !common.IsNil(o.ShortAccount) {
		return true
	}

	return false
}

// SetShortAccount gets a reference to the given string and assigns it to the ShortAccount field.
func (o *TopTraderLongShortRatioPositionsResponseInner) SetShortAccount(v string) {
	o.ShortAccount = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetTimestamp() string {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetTimestampOk() (*string, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *TopTraderLongShortRatioPositionsResponseInner) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o TopTraderLongShortRatioPositionsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopTraderLongShortRatioPositionsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.LongShortRatio) {
		toSerialize["longShortRatio"] = o.LongShortRatio
	}
	if !common.IsNil(o.LongAccount) {
		toSerialize["longAccount"] = o.LongAccount
	}
	if !common.IsNil(o.ShortAccount) {
		toSerialize["shortAccount"] = o.ShortAccount
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TopTraderLongShortRatioPositionsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varTopTraderLongShortRatioPositionsResponseInner := _TopTraderLongShortRatioPositionsResponseInner{}

	err = json.Unmarshal(data, &varTopTraderLongShortRatioPositionsResponseInner)

	if err != nil {
		return err
	}

	*o = TopTraderLongShortRatioPositionsResponseInner(varTopTraderLongShortRatioPositionsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "longShortRatio")
		delete(additionalProperties, "longAccount")
		delete(additionalProperties, "shortAccount")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTopTraderLongShortRatioPositionsResponseInner struct {
	value *TopTraderLongShortRatioPositionsResponseInner
	isSet bool
}

func (v NullableTopTraderLongShortRatioPositionsResponseInner) Get() *TopTraderLongShortRatioPositionsResponseInner {
	return v.value
}

func (v *NullableTopTraderLongShortRatioPositionsResponseInner) Set(val *TopTraderLongShortRatioPositionsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTopTraderLongShortRatioPositionsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTopTraderLongShortRatioPositionsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopTraderLongShortRatioPositionsResponseInner(val *TopTraderLongShortRatioPositionsResponseInner) *NullableTopTraderLongShortRatioPositionsResponseInner {
	return &NullableTopTraderLongShortRatioPositionsResponseInner{value: val, isSet: true}
}

func (v NullableTopTraderLongShortRatioPositionsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopTraderLongShortRatioPositionsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
