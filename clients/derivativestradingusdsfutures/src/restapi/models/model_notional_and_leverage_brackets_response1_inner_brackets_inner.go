/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the NotionalAndLeverageBracketsResponse1InnerBracketsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotionalAndLeverageBracketsResponse1InnerBracketsInner{}

// NotionalAndLeverageBracketsResponse1InnerBracketsInner struct for NotionalAndLeverageBracketsResponse1InnerBracketsInner
type NotionalAndLeverageBracketsResponse1InnerBracketsInner struct {
	Bracket              *int64   `json:"bracket,omitempty"`
	InitialLeverage      *int64   `json:"initialLeverage,omitempty"`
	NotionalCap          *int64   `json:"notionalCap,omitempty"`
	NotionalFloor        *int64   `json:"notionalFloor,omitempty"`
	MaintMarginRatio     *float32 `json:"maintMarginRatio,omitempty"`
	Cum                  *float32 `json:"cum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotionalAndLeverageBracketsResponse1InnerBracketsInner NotionalAndLeverageBracketsResponse1InnerBracketsInner

// NewNotionalAndLeverageBracketsResponse1InnerBracketsInner instantiates a new NotionalAndLeverageBracketsResponse1InnerBracketsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotionalAndLeverageBracketsResponse1InnerBracketsInner() *NotionalAndLeverageBracketsResponse1InnerBracketsInner {
	this := NotionalAndLeverageBracketsResponse1InnerBracketsInner{}
	return &this
}

// NewNotionalAndLeverageBracketsResponse1InnerBracketsInnerWithDefaults instantiates a new NotionalAndLeverageBracketsResponse1InnerBracketsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotionalAndLeverageBracketsResponse1InnerBracketsInnerWithDefaults() *NotionalAndLeverageBracketsResponse1InnerBracketsInner {
	this := NotionalAndLeverageBracketsResponse1InnerBracketsInner{}
	return &this
}

// GetBracket returns the Bracket field value if set, zero value otherwise.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) GetBracket() int64 {
	if o == nil || common.IsNil(o.Bracket) {
		var ret int64
		return ret
	}
	return *o.Bracket
}

// GetBracketOk returns a tuple with the Bracket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) GetBracketOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Bracket) {
		return nil, false
	}
	return o.Bracket, true
}

// HasBracket returns a boolean if a field has been set.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) HasBracket() bool {
	if o != nil && !common.IsNil(o.Bracket) {
		return true
	}

	return false
}

// SetBracket gets a reference to the given int64 and assigns it to the Bracket field.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) SetBracket(v int64) {
	o.Bracket = &v
}

// GetInitialLeverage returns the InitialLeverage field value if set, zero value otherwise.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) GetInitialLeverage() int64 {
	if o == nil || common.IsNil(o.InitialLeverage) {
		var ret int64
		return ret
	}
	return *o.InitialLeverage
}

// GetInitialLeverageOk returns a tuple with the InitialLeverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) GetInitialLeverageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InitialLeverage) {
		return nil, false
	}
	return o.InitialLeverage, true
}

// HasInitialLeverage returns a boolean if a field has been set.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) HasInitialLeverage() bool {
	if o != nil && !common.IsNil(o.InitialLeverage) {
		return true
	}

	return false
}

// SetInitialLeverage gets a reference to the given int64 and assigns it to the InitialLeverage field.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) SetInitialLeverage(v int64) {
	o.InitialLeverage = &v
}

// GetNotionalCap returns the NotionalCap field value if set, zero value otherwise.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) GetNotionalCap() int64 {
	if o == nil || common.IsNil(o.NotionalCap) {
		var ret int64
		return ret
	}
	return *o.NotionalCap
}

// GetNotionalCapOk returns a tuple with the NotionalCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) GetNotionalCapOk() (*int64, bool) {
	if o == nil || common.IsNil(o.NotionalCap) {
		return nil, false
	}
	return o.NotionalCap, true
}

// HasNotionalCap returns a boolean if a field has been set.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) HasNotionalCap() bool {
	if o != nil && !common.IsNil(o.NotionalCap) {
		return true
	}

	return false
}

// SetNotionalCap gets a reference to the given int64 and assigns it to the NotionalCap field.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) SetNotionalCap(v int64) {
	o.NotionalCap = &v
}

// GetNotionalFloor returns the NotionalFloor field value if set, zero value otherwise.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) GetNotionalFloor() int64 {
	if o == nil || common.IsNil(o.NotionalFloor) {
		var ret int64
		return ret
	}
	return *o.NotionalFloor
}

// GetNotionalFloorOk returns a tuple with the NotionalFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) GetNotionalFloorOk() (*int64, bool) {
	if o == nil || common.IsNil(o.NotionalFloor) {
		return nil, false
	}
	return o.NotionalFloor, true
}

// HasNotionalFloor returns a boolean if a field has been set.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) HasNotionalFloor() bool {
	if o != nil && !common.IsNil(o.NotionalFloor) {
		return true
	}

	return false
}

// SetNotionalFloor gets a reference to the given int64 and assigns it to the NotionalFloor field.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) SetNotionalFloor(v int64) {
	o.NotionalFloor = &v
}

// GetMaintMarginRatio returns the MaintMarginRatio field value if set, zero value otherwise.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) GetMaintMarginRatio() float32 {
	if o == nil || common.IsNil(o.MaintMarginRatio) {
		var ret float32
		return ret
	}
	return *o.MaintMarginRatio
}

// GetMaintMarginRatioOk returns a tuple with the MaintMarginRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) GetMaintMarginRatioOk() (*float32, bool) {
	if o == nil || common.IsNil(o.MaintMarginRatio) {
		return nil, false
	}
	return o.MaintMarginRatio, true
}

// HasMaintMarginRatio returns a boolean if a field has been set.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) HasMaintMarginRatio() bool {
	if o != nil && !common.IsNil(o.MaintMarginRatio) {
		return true
	}

	return false
}

// SetMaintMarginRatio gets a reference to the given float32 and assigns it to the MaintMarginRatio field.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) SetMaintMarginRatio(v float32) {
	o.MaintMarginRatio = &v
}

// GetCum returns the Cum field value if set, zero value otherwise.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) GetCum() float32 {
	if o == nil || common.IsNil(o.Cum) {
		var ret float32
		return ret
	}
	return *o.Cum
}

// GetCumOk returns a tuple with the Cum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) GetCumOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Cum) {
		return nil, false
	}
	return o.Cum, true
}

// HasCum returns a boolean if a field has been set.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) HasCum() bool {
	if o != nil && !common.IsNil(o.Cum) {
		return true
	}

	return false
}

// SetCum gets a reference to the given float32 and assigns it to the Cum field.
func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) SetCum(v float32) {
	o.Cum = &v
}

func (o NotionalAndLeverageBracketsResponse1InnerBracketsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotionalAndLeverageBracketsResponse1InnerBracketsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Bracket) {
		toSerialize["bracket"] = o.Bracket
	}
	if !common.IsNil(o.InitialLeverage) {
		toSerialize["initialLeverage"] = o.InitialLeverage
	}
	if !common.IsNil(o.NotionalCap) {
		toSerialize["notionalCap"] = o.NotionalCap
	}
	if !common.IsNil(o.NotionalFloor) {
		toSerialize["notionalFloor"] = o.NotionalFloor
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

func (o *NotionalAndLeverageBracketsResponse1InnerBracketsInner) UnmarshalJSON(data []byte) (err error) {
	varNotionalAndLeverageBracketsResponse1InnerBracketsInner := _NotionalAndLeverageBracketsResponse1InnerBracketsInner{}

	err = json.Unmarshal(data, &varNotionalAndLeverageBracketsResponse1InnerBracketsInner)

	if err != nil {
		return err
	}

	*o = NotionalAndLeverageBracketsResponse1InnerBracketsInner(varNotionalAndLeverageBracketsResponse1InnerBracketsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bracket")
		delete(additionalProperties, "initialLeverage")
		delete(additionalProperties, "notionalCap")
		delete(additionalProperties, "notionalFloor")
		delete(additionalProperties, "maintMarginRatio")
		delete(additionalProperties, "cum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotionalAndLeverageBracketsResponse1InnerBracketsInner struct {
	value *NotionalAndLeverageBracketsResponse1InnerBracketsInner
	isSet bool
}

func (v NullableNotionalAndLeverageBracketsResponse1InnerBracketsInner) Get() *NotionalAndLeverageBracketsResponse1InnerBracketsInner {
	return v.value
}

func (v *NullableNotionalAndLeverageBracketsResponse1InnerBracketsInner) Set(val *NotionalAndLeverageBracketsResponse1InnerBracketsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableNotionalAndLeverageBracketsResponse1InnerBracketsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableNotionalAndLeverageBracketsResponse1InnerBracketsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotionalAndLeverageBracketsResponse1InnerBracketsInner(val *NotionalAndLeverageBracketsResponse1InnerBracketsInner) *NullableNotionalAndLeverageBracketsResponse1InnerBracketsInner {
	return &NullableNotionalAndLeverageBracketsResponse1InnerBracketsInner{value: val, isSet: true}
}

func (v NullableNotionalAndLeverageBracketsResponse1InnerBracketsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotionalAndLeverageBracketsResponse1InnerBracketsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
