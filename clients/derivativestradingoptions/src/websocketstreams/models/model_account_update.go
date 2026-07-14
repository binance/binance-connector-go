/*
Options WebSocket Market Streams

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountUpdate{}

// AccountUpdate struct for AccountUpdate
type AccountUpdate struct {
	// Event Time
	E *int64 `json:"E,omitempty"`
	// Transaction Time
	T *int64 `json:"T,omitempty"`
	// Account equity in USDT
	Smalleq *string `json:"eq,omitempty"`
	// Account adjusted equity in USDT
	Aeq *string `json:"aeq,omitempty"`
	// Account wallet balance in USDT
	Smallb *string `json:"b,omitempty"`
	// Position value
	Smallm *string `json:"m,omitempty"`
	// Unrealized PnL
	Smallu *string `json:"u,omitempty"`
	// Initial margin in USDT
	Smalli *string `json:"i,omitempty"`
	// Maintenance margin in USDT
	M                    *string `json:"M,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUpdate AccountUpdate

// NewAccountUpdate instantiates a new AccountUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdate() *AccountUpdate {
	this := AccountUpdate{}
	return &this
}

// NewAccountUpdateWithDefaults instantiates a new AccountUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateWithDefaults() *AccountUpdate {
	this := AccountUpdate{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AccountUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AccountUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AccountUpdate) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AccountUpdate) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AccountUpdate) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AccountUpdate) SetT(v int64) {
	o.T = &v
}

// GetEq returns the Eq field value if set, zero value otherwise.
func (o *AccountUpdate) GetSmalleq() string {
	if o == nil || common.IsNil(o.Smalleq) {
		var ret string
		return ret
	}
	return *o.Smalleq
}

// GetEqOk returns a tuple with the Eq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetSmalleqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalleq) {
		return nil, false
	}
	return o.Smalleq, true
}

// HasEq returns a boolean if a field has been set.
func (o *AccountUpdate) HasSmalleq() bool {
	if o != nil && !common.IsNil(o.Smalleq) {
		return true
	}

	return false
}

// SetEq gets a reference to the given string and assigns it to the Eq field.
func (o *AccountUpdate) SetSmalleq(v string) {
	o.Smalleq = &v
}

// GetAeq returns the Aeq field value if set, zero value otherwise.
func (o *AccountUpdate) GetAeq() string {
	if o == nil || common.IsNil(o.Aeq) {
		var ret string
		return ret
	}
	return *o.Aeq
}

// GetAeqOk returns a tuple with the Aeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetAeqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Aeq) {
		return nil, false
	}
	return o.Aeq, true
}

// HasAeq returns a boolean if a field has been set.
func (o *AccountUpdate) HasAeq() bool {
	if o != nil && !common.IsNil(o.Aeq) {
		return true
	}

	return false
}

// SetAeq gets a reference to the given string and assigns it to the Aeq field.
func (o *AccountUpdate) SetAeq(v string) {
	o.Aeq = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *AccountUpdate) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *AccountUpdate) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *AccountUpdate) SetSmallb(v string) {
	o.Smallb = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *AccountUpdate) GetSmallm() string {
	if o == nil || common.IsNil(o.Smallm) {
		var ret string
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetSmallmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *AccountUpdate) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given string and assigns it to the M field.
func (o *AccountUpdate) SetSmallm(v string) {
	o.Smallm = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *AccountUpdate) GetSmallu() string {
	if o == nil || common.IsNil(o.Smallu) {
		var ret string
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetSmalluOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *AccountUpdate) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given string and assigns it to the U field.
func (o *AccountUpdate) SetSmallu(v string) {
	o.Smallu = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *AccountUpdate) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *AccountUpdate) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *AccountUpdate) SetSmalli(v string) {
	o.Smalli = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *AccountUpdate) GetM() string {
	if o == nil || common.IsNil(o.M) {
		var ret string
		return ret
	}
	return *o.M
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetMOk() (*string, bool) {
	if o == nil || common.IsNil(o.M) {
		return nil, false
	}
	return o.M, true
}

// HasM returns a boolean if a field has been set.
func (o *AccountUpdate) HasM() bool {
	if o != nil && !common.IsNil(o.M) {
		return true
	}

	return false
}

// SetM gets a reference to the given string and assigns it to the M field.
func (o *AccountUpdate) SetM(v string) {
	o.M = &v
}

func (o AccountUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smalleq) {
		toSerialize["eq"] = o.Smalleq
	}
	if !common.IsNil(o.Aeq) {
		toSerialize["aeq"] = o.Aeq
	}
	if !common.IsNil(o.Smallb) {
		toSerialize["b"] = o.Smallb
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.M) {
		toSerialize["M"] = o.M
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountUpdate) UnmarshalJSON(data []byte) (err error) {
	varAccountUpdate := _AccountUpdate{}

	err = json.Unmarshal(data, &varAccountUpdate)

	if err != nil {
		return err
	}

	*o = AccountUpdate(varAccountUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "eq")
		delete(additionalProperties, "aeq")
		delete(additionalProperties, "b")
		delete(additionalProperties, "m")
		delete(additionalProperties, "u")
		delete(additionalProperties, "i")
		delete(additionalProperties, "M")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUpdate struct {
	value *AccountUpdate
	isSet bool
}

func (v NullableAccountUpdate) Get() *AccountUpdate {
	return v.value
}

func (v *NullableAccountUpdate) Set(val *AccountUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdate(val *AccountUpdate) *NullableAccountUpdate {
	return &NullableAccountUpdate{value: val, isSet: true}
}

func (v NullableAccountUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
