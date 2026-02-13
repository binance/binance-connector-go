/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CmNotionalAndLeverageBracketsResponseInnerBracketsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CmNotionalAndLeverageBracketsResponseInnerBracketsInner{}

// CmNotionalAndLeverageBracketsResponseInnerBracketsInner struct for CmNotionalAndLeverageBracketsResponseInnerBracketsInner
type CmNotionalAndLeverageBracketsResponseInnerBracketsInner struct {
	Bracket              *int64   `json:"bracket,omitempty"`
	InitialLeverage      *int64   `json:"initialLeverage,omitempty"`
	QtyCap               *int64   `json:"qtyCap,omitempty"`
	QtyFloor             *int64   `json:"qtyFloor,omitempty"`
	MaintMarginRatio     *float32 `json:"maintMarginRatio,omitempty"`
	Cum                  *float32 `json:"cum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CmNotionalAndLeverageBracketsResponseInnerBracketsInner CmNotionalAndLeverageBracketsResponseInnerBracketsInner

// NewCmNotionalAndLeverageBracketsResponseInnerBracketsInner instantiates a new CmNotionalAndLeverageBracketsResponseInnerBracketsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmNotionalAndLeverageBracketsResponseInnerBracketsInner() *CmNotionalAndLeverageBracketsResponseInnerBracketsInner {
	this := CmNotionalAndLeverageBracketsResponseInnerBracketsInner{}
	return &this
}

// NewCmNotionalAndLeverageBracketsResponseInnerBracketsInnerWithDefaults instantiates a new CmNotionalAndLeverageBracketsResponseInnerBracketsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmNotionalAndLeverageBracketsResponseInnerBracketsInnerWithDefaults() *CmNotionalAndLeverageBracketsResponseInnerBracketsInner {
	this := CmNotionalAndLeverageBracketsResponseInnerBracketsInner{}
	return &this
}

// GetBracket returns the Bracket field value if set, zero value otherwise.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetBracket() int64 {
	if o == nil || common.IsNil(o.Bracket) {
		var ret int64
		return ret
	}
	return *o.Bracket
}

// GetBracketOk returns a tuple with the Bracket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetBracketOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Bracket) {
		return nil, false
	}
	return o.Bracket, true
}

// HasBracket returns a boolean if a field has been set.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) HasBracket() bool {
	if o != nil && !common.IsNil(o.Bracket) {
		return true
	}

	return false
}

// SetBracket gets a reference to the given int64 and assigns it to the Bracket field.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) SetBracket(v int64) {
	o.Bracket = &v
}

// GetInitialLeverage returns the InitialLeverage field value if set, zero value otherwise.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetInitialLeverage() int64 {
	if o == nil || common.IsNil(o.InitialLeverage) {
		var ret int64
		return ret
	}
	return *o.InitialLeverage
}

// GetInitialLeverageOk returns a tuple with the InitialLeverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetInitialLeverageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InitialLeverage) {
		return nil, false
	}
	return o.InitialLeverage, true
}

// HasInitialLeverage returns a boolean if a field has been set.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) HasInitialLeverage() bool {
	if o != nil && !common.IsNil(o.InitialLeverage) {
		return true
	}

	return false
}

// SetInitialLeverage gets a reference to the given int64 and assigns it to the InitialLeverage field.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) SetInitialLeverage(v int64) {
	o.InitialLeverage = &v
}

// GetQtyCap returns the QtyCap field value if set, zero value otherwise.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetQtyCap() int64 {
	if o == nil || common.IsNil(o.QtyCap) {
		var ret int64
		return ret
	}
	return *o.QtyCap
}

// GetQtyCapOk returns a tuple with the QtyCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetQtyCapOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QtyCap) {
		return nil, false
	}
	return o.QtyCap, true
}

// HasQtyCap returns a boolean if a field has been set.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) HasQtyCap() bool {
	if o != nil && !common.IsNil(o.QtyCap) {
		return true
	}

	return false
}

// SetQtyCap gets a reference to the given int64 and assigns it to the QtyCap field.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) SetQtyCap(v int64) {
	o.QtyCap = &v
}

// GetQtyFloor returns the QtyFloor field value if set, zero value otherwise.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetQtyFloor() int64 {
	if o == nil || common.IsNil(o.QtyFloor) {
		var ret int64
		return ret
	}
	return *o.QtyFloor
}

// GetQtyFloorOk returns a tuple with the QtyFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetQtyFloorOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QtyFloor) {
		return nil, false
	}
	return o.QtyFloor, true
}

// HasQtyFloor returns a boolean if a field has been set.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) HasQtyFloor() bool {
	if o != nil && !common.IsNil(o.QtyFloor) {
		return true
	}

	return false
}

// SetQtyFloor gets a reference to the given int64 and assigns it to the QtyFloor field.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) SetQtyFloor(v int64) {
	o.QtyFloor = &v
}

// GetMaintMarginRatio returns the MaintMarginRatio field value if set, zero value otherwise.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetMaintMarginRatio() float32 {
	if o == nil || common.IsNil(o.MaintMarginRatio) {
		var ret float32
		return ret
	}
	return *o.MaintMarginRatio
}

// GetMaintMarginRatioOk returns a tuple with the MaintMarginRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetMaintMarginRatioOk() (*float32, bool) {
	if o == nil || common.IsNil(o.MaintMarginRatio) {
		return nil, false
	}
	return o.MaintMarginRatio, true
}

// HasMaintMarginRatio returns a boolean if a field has been set.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) HasMaintMarginRatio() bool {
	if o != nil && !common.IsNil(o.MaintMarginRatio) {
		return true
	}

	return false
}

// SetMaintMarginRatio gets a reference to the given float32 and assigns it to the MaintMarginRatio field.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) SetMaintMarginRatio(v float32) {
	o.MaintMarginRatio = &v
}

// GetCum returns the Cum field value if set, zero value otherwise.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetCum() float32 {
	if o == nil || common.IsNil(o.Cum) {
		var ret float32
		return ret
	}
	return *o.Cum
}

// GetCumOk returns a tuple with the Cum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) GetCumOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Cum) {
		return nil, false
	}
	return o.Cum, true
}

// HasCum returns a boolean if a field has been set.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) HasCum() bool {
	if o != nil && !common.IsNil(o.Cum) {
		return true
	}

	return false
}

// SetCum gets a reference to the given float32 and assigns it to the Cum field.
func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) SetCum(v float32) {
	o.Cum = &v
}

func (o CmNotionalAndLeverageBracketsResponseInnerBracketsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmNotionalAndLeverageBracketsResponseInnerBracketsInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.QtyFloor) {
		toSerialize["qtyFloor"] = o.QtyFloor
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

func (o *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) UnmarshalJSON(data []byte) (err error) {
	varCmNotionalAndLeverageBracketsResponseInnerBracketsInner := _CmNotionalAndLeverageBracketsResponseInnerBracketsInner{}

	err = json.Unmarshal(data, &varCmNotionalAndLeverageBracketsResponseInnerBracketsInner)

	if err != nil {
		return err
	}

	*o = CmNotionalAndLeverageBracketsResponseInnerBracketsInner(varCmNotionalAndLeverageBracketsResponseInnerBracketsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bracket")
		delete(additionalProperties, "initialLeverage")
		delete(additionalProperties, "qtyCap")
		delete(additionalProperties, "qtyFloor")
		delete(additionalProperties, "maintMarginRatio")
		delete(additionalProperties, "cum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCmNotionalAndLeverageBracketsResponseInnerBracketsInner struct {
	value *CmNotionalAndLeverageBracketsResponseInnerBracketsInner
	isSet bool
}

func (v NullableCmNotionalAndLeverageBracketsResponseInnerBracketsInner) Get() *CmNotionalAndLeverageBracketsResponseInnerBracketsInner {
	return v.value
}

func (v *NullableCmNotionalAndLeverageBracketsResponseInnerBracketsInner) Set(val *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCmNotionalAndLeverageBracketsResponseInnerBracketsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCmNotionalAndLeverageBracketsResponseInnerBracketsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmNotionalAndLeverageBracketsResponseInnerBracketsInner(val *CmNotionalAndLeverageBracketsResponseInnerBracketsInner) *NullableCmNotionalAndLeverageBracketsResponseInnerBracketsInner {
	return &NullableCmNotionalAndLeverageBracketsResponseInnerBracketsInner{value: val, isSet: true}
}

func (v NullableCmNotionalAndLeverageBracketsResponseInnerBracketsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmNotionalAndLeverageBracketsResponseInnerBracketsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
