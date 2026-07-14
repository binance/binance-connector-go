/*
Portfolio Margin WebSocket Market Streams

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RiskLevelChange type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RiskLevelChange{}

// RiskLevelChange struct for RiskLevelChange
type RiskLevelChange struct {
	// Event Time
	E *int64 `json:"E,omitempty"`
	// uniMMR level
	Smallu *string `json:"u,omitempty"`
	// Risk level: MARGIN_CALL, REDUCE_ONLY, FORCE_LIQUIDATION
	Smalls *string `json:"s,omitempty"`
	// Account equity in USD value
	Smalleq *string `json:"eq,omitempty"`
	// Actual equity without collateral rate in USD value
	Smallae *string `json:"ae,omitempty"`
	// Total maintenance margin in USD value
	Smallm               *string `json:"m,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskLevelChange RiskLevelChange

// NewRiskLevelChange instantiates a new RiskLevelChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskLevelChange() *RiskLevelChange {
	this := RiskLevelChange{}
	return &this
}

// NewRiskLevelChangeWithDefaults instantiates a new RiskLevelChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskLevelChangeWithDefaults() *RiskLevelChange {
	this := RiskLevelChange{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *RiskLevelChange) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskLevelChange) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *RiskLevelChange) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *RiskLevelChange) SetE(v int64) {
	o.E = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *RiskLevelChange) GetSmallu() string {
	if o == nil || common.IsNil(o.Smallu) {
		var ret string
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskLevelChange) GetSmalluOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *RiskLevelChange) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given string and assigns it to the U field.
func (o *RiskLevelChange) SetSmallu(v string) {
	o.Smallu = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *RiskLevelChange) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskLevelChange) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *RiskLevelChange) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *RiskLevelChange) SetSmalls(v string) {
	o.Smalls = &v
}

// GetEq returns the Eq field value if set, zero value otherwise.
func (o *RiskLevelChange) GetSmalleq() string {
	if o == nil || common.IsNil(o.Smalleq) {
		var ret string
		return ret
	}
	return *o.Smalleq
}

// GetEqOk returns a tuple with the Eq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskLevelChange) GetSmalleqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalleq) {
		return nil, false
	}
	return o.Smalleq, true
}

// HasEq returns a boolean if a field has been set.
func (o *RiskLevelChange) HasSmalleq() bool {
	if o != nil && !common.IsNil(o.Smalleq) {
		return true
	}

	return false
}

// SetEq gets a reference to the given string and assigns it to the Eq field.
func (o *RiskLevelChange) SetSmalleq(v string) {
	o.Smalleq = &v
}

// GetAe returns the Ae field value if set, zero value otherwise.
func (o *RiskLevelChange) GetSmallae() string {
	if o == nil || common.IsNil(o.Smallae) {
		var ret string
		return ret
	}
	return *o.Smallae
}

// GetAeOk returns a tuple with the Ae field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskLevelChange) GetSmallaeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallae) {
		return nil, false
	}
	return o.Smallae, true
}

// HasAe returns a boolean if a field has been set.
func (o *RiskLevelChange) HasSmallae() bool {
	if o != nil && !common.IsNil(o.Smallae) {
		return true
	}

	return false
}

// SetAe gets a reference to the given string and assigns it to the Ae field.
func (o *RiskLevelChange) SetSmallae(v string) {
	o.Smallae = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *RiskLevelChange) GetSmallm() string {
	if o == nil || common.IsNil(o.Smallm) {
		var ret string
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskLevelChange) GetSmallmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *RiskLevelChange) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given string and assigns it to the M field.
func (o *RiskLevelChange) SetSmallm(v string) {
	o.Smallm = &v
}

func (o RiskLevelChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskLevelChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smalleq) {
		toSerialize["eq"] = o.Smalleq
	}
	if !common.IsNil(o.Smallae) {
		toSerialize["ae"] = o.Smallae
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RiskLevelChange) UnmarshalJSON(data []byte) (err error) {
	varRiskLevelChange := _RiskLevelChange{}

	err = json.Unmarshal(data, &varRiskLevelChange)

	if err != nil {
		return err
	}

	*o = RiskLevelChange(varRiskLevelChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "u")
		delete(additionalProperties, "s")
		delete(additionalProperties, "eq")
		delete(additionalProperties, "ae")
		delete(additionalProperties, "m")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskLevelChange struct {
	value *RiskLevelChange
	isSet bool
}

func (v NullableRiskLevelChange) Get() *RiskLevelChange {
	return v.value
}

func (v *NullableRiskLevelChange) Set(val *RiskLevelChange) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskLevelChange) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskLevelChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskLevelChange(val *RiskLevelChange) *NullableRiskLevelChange {
	return &NullableRiskLevelChange{value: val, isSet: true}
}

func (v NullableRiskLevelChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskLevelChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
