/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the UmNotionalAndLeverageBracketsResponseInnerBracketsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UmNotionalAndLeverageBracketsResponseInnerBracketsInner{}

// UmNotionalAndLeverageBracketsResponseInnerBracketsInner struct for UmNotionalAndLeverageBracketsResponseInnerBracketsInner
type UmNotionalAndLeverageBracketsResponseInnerBracketsInner struct {
	Bracket              *int64   `json:"bracket,omitempty"`
	InitialLeverage      *int64   `json:"initialLeverage,omitempty"`
	NotionalCap          *int64   `json:"notionalCap,omitempty"`
	NotionalFloor        *int64   `json:"notionalFloor,omitempty"`
	MaintMarginRatio     *float32 `json:"maintMarginRatio,omitempty"`
	Cum                  *int64   `json:"cum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UmNotionalAndLeverageBracketsResponseInnerBracketsInner UmNotionalAndLeverageBracketsResponseInnerBracketsInner

// NewUmNotionalAndLeverageBracketsResponseInnerBracketsInner instantiates a new UmNotionalAndLeverageBracketsResponseInnerBracketsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmNotionalAndLeverageBracketsResponseInnerBracketsInner() *UmNotionalAndLeverageBracketsResponseInnerBracketsInner {
	this := UmNotionalAndLeverageBracketsResponseInnerBracketsInner{}
	return &this
}

// NewUmNotionalAndLeverageBracketsResponseInnerBracketsInnerWithDefaults instantiates a new UmNotionalAndLeverageBracketsResponseInnerBracketsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmNotionalAndLeverageBracketsResponseInnerBracketsInnerWithDefaults() *UmNotionalAndLeverageBracketsResponseInnerBracketsInner {
	this := UmNotionalAndLeverageBracketsResponseInnerBracketsInner{}
	return &this
}

// GetBracket returns the Bracket field value if set, zero value otherwise.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetBracket() int64 {
	if o == nil || common.IsNil(o.Bracket) {
		var ret int64
		return ret
	}
	return *o.Bracket
}

// GetBracketOk returns a tuple with the Bracket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetBracketOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Bracket) {
		return nil, false
	}
	return o.Bracket, true
}

// HasBracket returns a boolean if a field has been set.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) HasBracket() bool {
	if o != nil && !common.IsNil(o.Bracket) {
		return true
	}

	return false
}

// SetBracket gets a reference to the given int64 and assigns it to the Bracket field.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) SetBracket(v int64) {
	o.Bracket = &v
}

// GetInitialLeverage returns the InitialLeverage field value if set, zero value otherwise.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetInitialLeverage() int64 {
	if o == nil || common.IsNil(o.InitialLeverage) {
		var ret int64
		return ret
	}
	return *o.InitialLeverage
}

// GetInitialLeverageOk returns a tuple with the InitialLeverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetInitialLeverageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InitialLeverage) {
		return nil, false
	}
	return o.InitialLeverage, true
}

// HasInitialLeverage returns a boolean if a field has been set.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) HasInitialLeverage() bool {
	if o != nil && !common.IsNil(o.InitialLeverage) {
		return true
	}

	return false
}

// SetInitialLeverage gets a reference to the given int64 and assigns it to the InitialLeverage field.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) SetInitialLeverage(v int64) {
	o.InitialLeverage = &v
}

// GetNotionalCap returns the NotionalCap field value if set, zero value otherwise.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetNotionalCap() int64 {
	if o == nil || common.IsNil(o.NotionalCap) {
		var ret int64
		return ret
	}
	return *o.NotionalCap
}

// GetNotionalCapOk returns a tuple with the NotionalCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetNotionalCapOk() (*int64, bool) {
	if o == nil || common.IsNil(o.NotionalCap) {
		return nil, false
	}
	return o.NotionalCap, true
}

// HasNotionalCap returns a boolean if a field has been set.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) HasNotionalCap() bool {
	if o != nil && !common.IsNil(o.NotionalCap) {
		return true
	}

	return false
}

// SetNotionalCap gets a reference to the given int64 and assigns it to the NotionalCap field.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) SetNotionalCap(v int64) {
	o.NotionalCap = &v
}

// GetNotionalFloor returns the NotionalFloor field value if set, zero value otherwise.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetNotionalFloor() int64 {
	if o == nil || common.IsNil(o.NotionalFloor) {
		var ret int64
		return ret
	}
	return *o.NotionalFloor
}

// GetNotionalFloorOk returns a tuple with the NotionalFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetNotionalFloorOk() (*int64, bool) {
	if o == nil || common.IsNil(o.NotionalFloor) {
		return nil, false
	}
	return o.NotionalFloor, true
}

// HasNotionalFloor returns a boolean if a field has been set.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) HasNotionalFloor() bool {
	if o != nil && !common.IsNil(o.NotionalFloor) {
		return true
	}

	return false
}

// SetNotionalFloor gets a reference to the given int64 and assigns it to the NotionalFloor field.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) SetNotionalFloor(v int64) {
	o.NotionalFloor = &v
}

// GetMaintMarginRatio returns the MaintMarginRatio field value if set, zero value otherwise.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetMaintMarginRatio() float32 {
	if o == nil || common.IsNil(o.MaintMarginRatio) {
		var ret float32
		return ret
	}
	return *o.MaintMarginRatio
}

// GetMaintMarginRatioOk returns a tuple with the MaintMarginRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetMaintMarginRatioOk() (*float32, bool) {
	if o == nil || common.IsNil(o.MaintMarginRatio) {
		return nil, false
	}
	return o.MaintMarginRatio, true
}

// HasMaintMarginRatio returns a boolean if a field has been set.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) HasMaintMarginRatio() bool {
	if o != nil && !common.IsNil(o.MaintMarginRatio) {
		return true
	}

	return false
}

// SetMaintMarginRatio gets a reference to the given float32 and assigns it to the MaintMarginRatio field.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) SetMaintMarginRatio(v float32) {
	o.MaintMarginRatio = &v
}

// GetCum returns the Cum field value if set, zero value otherwise.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetCum() int64 {
	if o == nil || common.IsNil(o.Cum) {
		var ret int64
		return ret
	}
	return *o.Cum
}

// GetCumOk returns a tuple with the Cum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetCumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Cum) {
		return nil, false
	}
	return o.Cum, true
}

// HasCum returns a boolean if a field has been set.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) HasCum() bool {
	if o != nil && !common.IsNil(o.Cum) {
		return true
	}

	return false
}

// SetCum gets a reference to the given int64 and assigns it to the Cum field.
func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) SetCum(v int64) {
	o.Cum = &v
}

func (o UmNotionalAndLeverageBracketsResponseInnerBracketsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmNotionalAndLeverageBracketsResponseInnerBracketsInner) ToMap() (map[string]interface{}, error) {
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

func (o *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) UnmarshalJSON(data []byte) (err error) {
	varUmNotionalAndLeverageBracketsResponseInnerBracketsInner := _UmNotionalAndLeverageBracketsResponseInnerBracketsInner{}

	err = json.Unmarshal(data, &varUmNotionalAndLeverageBracketsResponseInnerBracketsInner)

	if err != nil {
		return err
	}

	*o = UmNotionalAndLeverageBracketsResponseInnerBracketsInner(varUmNotionalAndLeverageBracketsResponseInnerBracketsInner)

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

type NullableUmNotionalAndLeverageBracketsResponseInnerBracketsInner struct {
	value *UmNotionalAndLeverageBracketsResponseInnerBracketsInner
	isSet bool
}

func (v NullableUmNotionalAndLeverageBracketsResponseInnerBracketsInner) Get() *UmNotionalAndLeverageBracketsResponseInnerBracketsInner {
	return v.value
}

func (v *NullableUmNotionalAndLeverageBracketsResponseInnerBracketsInner) Set(val *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUmNotionalAndLeverageBracketsResponseInnerBracketsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUmNotionalAndLeverageBracketsResponseInnerBracketsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmNotionalAndLeverageBracketsResponseInnerBracketsInner(val *UmNotionalAndLeverageBracketsResponseInnerBracketsInner) *NullableUmNotionalAndLeverageBracketsResponseInnerBracketsInner {
	return &NullableUmNotionalAndLeverageBracketsResponseInnerBracketsInner{value: val, isSet: true}
}

func (v NullableUmNotionalAndLeverageBracketsResponseInnerBracketsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmNotionalAndLeverageBracketsResponseInnerBracketsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
