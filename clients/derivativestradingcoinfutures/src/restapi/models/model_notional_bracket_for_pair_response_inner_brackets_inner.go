/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the NotionalBracketForPairResponseInnerBracketsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotionalBracketForPairResponseInnerBracketsInner{}

// NotionalBracketForPairResponseInnerBracketsInner struct for NotionalBracketForPairResponseInnerBracketsInner
type NotionalBracketForPairResponseInnerBracketsInner struct {
	Bracket              *int64   `json:"bracket,omitempty"`
	InitialLeverage      *int64   `json:"initialLeverage,omitempty"`
	QtyCap               *int64   `json:"qtyCap,omitempty"`
	QtylFloor            *int64   `json:"qtylFloor,omitempty"`
	MaintMarginRatio     *float32 `json:"maintMarginRatio,omitempty"`
	Cum                  *float32 `json:"cum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotionalBracketForPairResponseInnerBracketsInner NotionalBracketForPairResponseInnerBracketsInner

// NewNotionalBracketForPairResponseInnerBracketsInner instantiates a new NotionalBracketForPairResponseInnerBracketsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotionalBracketForPairResponseInnerBracketsInner() *NotionalBracketForPairResponseInnerBracketsInner {
	this := NotionalBracketForPairResponseInnerBracketsInner{}
	return &this
}

// NewNotionalBracketForPairResponseInnerBracketsInnerWithDefaults instantiates a new NotionalBracketForPairResponseInnerBracketsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotionalBracketForPairResponseInnerBracketsInnerWithDefaults() *NotionalBracketForPairResponseInnerBracketsInner {
	this := NotionalBracketForPairResponseInnerBracketsInner{}
	return &this
}

// GetBracket returns the Bracket field value if set, zero value otherwise.
func (o *NotionalBracketForPairResponseInnerBracketsInner) GetBracket() int64 {
	if o == nil || common.IsNil(o.Bracket) {
		var ret int64
		return ret
	}
	return *o.Bracket
}

// GetBracketOk returns a tuple with the Bracket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalBracketForPairResponseInnerBracketsInner) GetBracketOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Bracket) {
		return nil, false
	}
	return o.Bracket, true
}

// HasBracket returns a boolean if a field has been set.
func (o *NotionalBracketForPairResponseInnerBracketsInner) HasBracket() bool {
	if o != nil && !common.IsNil(o.Bracket) {
		return true
	}

	return false
}

// SetBracket gets a reference to the given int64 and assigns it to the Bracket field.
func (o *NotionalBracketForPairResponseInnerBracketsInner) SetBracket(v int64) {
	o.Bracket = &v
}

// GetInitialLeverage returns the InitialLeverage field value if set, zero value otherwise.
func (o *NotionalBracketForPairResponseInnerBracketsInner) GetInitialLeverage() int64 {
	if o == nil || common.IsNil(o.InitialLeverage) {
		var ret int64
		return ret
	}
	return *o.InitialLeverage
}

// GetInitialLeverageOk returns a tuple with the InitialLeverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalBracketForPairResponseInnerBracketsInner) GetInitialLeverageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InitialLeverage) {
		return nil, false
	}
	return o.InitialLeverage, true
}

// HasInitialLeverage returns a boolean if a field has been set.
func (o *NotionalBracketForPairResponseInnerBracketsInner) HasInitialLeverage() bool {
	if o != nil && !common.IsNil(o.InitialLeverage) {
		return true
	}

	return false
}

// SetInitialLeverage gets a reference to the given int64 and assigns it to the InitialLeverage field.
func (o *NotionalBracketForPairResponseInnerBracketsInner) SetInitialLeverage(v int64) {
	o.InitialLeverage = &v
}

// GetQtyCap returns the QtyCap field value if set, zero value otherwise.
func (o *NotionalBracketForPairResponseInnerBracketsInner) GetQtyCap() int64 {
	if o == nil || common.IsNil(o.QtyCap) {
		var ret int64
		return ret
	}
	return *o.QtyCap
}

// GetQtyCapOk returns a tuple with the QtyCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalBracketForPairResponseInnerBracketsInner) GetQtyCapOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QtyCap) {
		return nil, false
	}
	return o.QtyCap, true
}

// HasQtyCap returns a boolean if a field has been set.
func (o *NotionalBracketForPairResponseInnerBracketsInner) HasQtyCap() bool {
	if o != nil && !common.IsNil(o.QtyCap) {
		return true
	}

	return false
}

// SetQtyCap gets a reference to the given int64 and assigns it to the QtyCap field.
func (o *NotionalBracketForPairResponseInnerBracketsInner) SetQtyCap(v int64) {
	o.QtyCap = &v
}

// GetQtylFloor returns the QtylFloor field value if set, zero value otherwise.
func (o *NotionalBracketForPairResponseInnerBracketsInner) GetQtylFloor() int64 {
	if o == nil || common.IsNil(o.QtylFloor) {
		var ret int64
		return ret
	}
	return *o.QtylFloor
}

// GetQtylFloorOk returns a tuple with the QtylFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalBracketForPairResponseInnerBracketsInner) GetQtylFloorOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QtylFloor) {
		return nil, false
	}
	return o.QtylFloor, true
}

// HasQtylFloor returns a boolean if a field has been set.
func (o *NotionalBracketForPairResponseInnerBracketsInner) HasQtylFloor() bool {
	if o != nil && !common.IsNil(o.QtylFloor) {
		return true
	}

	return false
}

// SetQtylFloor gets a reference to the given int64 and assigns it to the QtylFloor field.
func (o *NotionalBracketForPairResponseInnerBracketsInner) SetQtylFloor(v int64) {
	o.QtylFloor = &v
}

// GetMaintMarginRatio returns the MaintMarginRatio field value if set, zero value otherwise.
func (o *NotionalBracketForPairResponseInnerBracketsInner) GetMaintMarginRatio() float32 {
	if o == nil || common.IsNil(o.MaintMarginRatio) {
		var ret float32
		return ret
	}
	return *o.MaintMarginRatio
}

// GetMaintMarginRatioOk returns a tuple with the MaintMarginRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalBracketForPairResponseInnerBracketsInner) GetMaintMarginRatioOk() (*float32, bool) {
	if o == nil || common.IsNil(o.MaintMarginRatio) {
		return nil, false
	}
	return o.MaintMarginRatio, true
}

// HasMaintMarginRatio returns a boolean if a field has been set.
func (o *NotionalBracketForPairResponseInnerBracketsInner) HasMaintMarginRatio() bool {
	if o != nil && !common.IsNil(o.MaintMarginRatio) {
		return true
	}

	return false
}

// SetMaintMarginRatio gets a reference to the given float32 and assigns it to the MaintMarginRatio field.
func (o *NotionalBracketForPairResponseInnerBracketsInner) SetMaintMarginRatio(v float32) {
	o.MaintMarginRatio = &v
}

// GetCum returns the Cum field value if set, zero value otherwise.
func (o *NotionalBracketForPairResponseInnerBracketsInner) GetCum() float32 {
	if o == nil || common.IsNil(o.Cum) {
		var ret float32
		return ret
	}
	return *o.Cum
}

// GetCumOk returns a tuple with the Cum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalBracketForPairResponseInnerBracketsInner) GetCumOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Cum) {
		return nil, false
	}
	return o.Cum, true
}

// HasCum returns a boolean if a field has been set.
func (o *NotionalBracketForPairResponseInnerBracketsInner) HasCum() bool {
	if o != nil && !common.IsNil(o.Cum) {
		return true
	}

	return false
}

// SetCum gets a reference to the given float32 and assigns it to the Cum field.
func (o *NotionalBracketForPairResponseInnerBracketsInner) SetCum(v float32) {
	o.Cum = &v
}

func (o NotionalBracketForPairResponseInnerBracketsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotionalBracketForPairResponseInnerBracketsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Bracket) {
		toSerialize["bracket"] = o.Bracket
	}
	if !common.IsNil(o.InitialLeverage) {
		toSerialize["initialLeverage"] = o.InitialLeverage
	}
	if !common.IsNil(o.QtyCap) {
		toSerialize["qtyCap"] = o.QtyCap
	}
	if !common.IsNil(o.QtylFloor) {
		toSerialize["qtylFloor"] = o.QtylFloor
	}
	if !common.IsNil(o.MaintMarginRatio) {
		toSerialize["maintMarginRatio"] = o.MaintMarginRatio
	}
	if !common.IsNil(o.Cum) {
		toSerialize["cum"] = o.Cum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotionalBracketForPairResponseInnerBracketsInner) UnmarshalJSON(data []byte) (err error) {
	varNotionalBracketForPairResponseInnerBracketsInner := _NotionalBracketForPairResponseInnerBracketsInner{}

	err = json.Unmarshal(data, &varNotionalBracketForPairResponseInnerBracketsInner)

	if err != nil {
		return err
	}

	*o = NotionalBracketForPairResponseInnerBracketsInner(varNotionalBracketForPairResponseInnerBracketsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bracket")
		delete(additionalProperties, "initialLeverage")
		delete(additionalProperties, "qtyCap")
		delete(additionalProperties, "qtylFloor")
		delete(additionalProperties, "maintMarginRatio")
		delete(additionalProperties, "cum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotionalBracketForPairResponseInnerBracketsInner struct {
	value *NotionalBracketForPairResponseInnerBracketsInner
	isSet bool
}

func (v NullableNotionalBracketForPairResponseInnerBracketsInner) Get() *NotionalBracketForPairResponseInnerBracketsInner {
	return v.value
}

func (v *NullableNotionalBracketForPairResponseInnerBracketsInner) Set(val *NotionalBracketForPairResponseInnerBracketsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableNotionalBracketForPairResponseInnerBracketsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableNotionalBracketForPairResponseInnerBracketsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotionalBracketForPairResponseInnerBracketsInner(val *NotionalBracketForPairResponseInnerBracketsInner) *NullableNotionalBracketForPairResponseInnerBracketsInner {
	return &NullableNotionalBracketForPairResponseInnerBracketsInner{value: val, isSet: true}
}

func (v NullableNotionalBracketForPairResponseInnerBracketsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotionalBracketForPairResponseInnerBracketsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
