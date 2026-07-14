/*
Portfolio Margin Pro WebSocket Market Streams

Access advanced account management and high-frequency trading with Binance Portfolio Margin Pro.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PmProAccountUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PmProAccountUpdate{}

// PmProAccountUpdate Pushed every 5 seconds with account asset status.
type PmProAccountUpdate struct {
	// Event Time
	E *int64 `json:"E,omitempty"`
	// uniMMR level
	Smallu *string `json:"u,omitempty"`
	// Account equity in USD
	Smalleq *string `json:"eq,omitempty"`
	// Actual equity without collateral rate in USD
	Smallae *string `json:"ae,omitempty"`
	// Total initial margin in USD
	Smallim *string `json:"im,omitempty"`
	// Total maintenance margin in USD
	Smallmm *string `json:"mm,omitempty"`
	// Total available balance in USD
	Avb *string `json:"avb,omitempty"`
	// Virtual maxWithdraw amount in USD
	Vmw                  *string `json:"vmw,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PmProAccountUpdate PmProAccountUpdate

// NewPmProAccountUpdate instantiates a new PmProAccountUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmProAccountUpdate() *PmProAccountUpdate {
	this := PmProAccountUpdate{}
	return &this
}

// NewPmProAccountUpdateWithDefaults instantiates a new PmProAccountUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmProAccountUpdateWithDefaults() *PmProAccountUpdate {
	this := PmProAccountUpdate{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *PmProAccountUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmProAccountUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *PmProAccountUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *PmProAccountUpdate) SetE(v int64) {
	o.E = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *PmProAccountUpdate) GetSmallu() string {
	if o == nil || common.IsNil(o.Smallu) {
		var ret string
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmProAccountUpdate) GetSmalluOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *PmProAccountUpdate) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given string and assigns it to the U field.
func (o *PmProAccountUpdate) SetSmallu(v string) {
	o.Smallu = &v
}

// GetEq returns the Eq field value if set, zero value otherwise.
func (o *PmProAccountUpdate) GetSmalleq() string {
	if o == nil || common.IsNil(o.Smalleq) {
		var ret string
		return ret
	}
	return *o.Smalleq
}

// GetEqOk returns a tuple with the Eq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmProAccountUpdate) GetSmalleqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalleq) {
		return nil, false
	}
	return o.Smalleq, true
}

// HasEq returns a boolean if a field has been set.
func (o *PmProAccountUpdate) HasSmalleq() bool {
	if o != nil && !common.IsNil(o.Smalleq) {
		return true
	}

	return false
}

// SetEq gets a reference to the given string and assigns it to the Eq field.
func (o *PmProAccountUpdate) SetSmalleq(v string) {
	o.Smalleq = &v
}

// GetAe returns the Ae field value if set, zero value otherwise.
func (o *PmProAccountUpdate) GetSmallae() string {
	if o == nil || common.IsNil(o.Smallae) {
		var ret string
		return ret
	}
	return *o.Smallae
}

// GetAeOk returns a tuple with the Ae field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmProAccountUpdate) GetSmallaeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallae) {
		return nil, false
	}
	return o.Smallae, true
}

// HasAe returns a boolean if a field has been set.
func (o *PmProAccountUpdate) HasSmallae() bool {
	if o != nil && !common.IsNil(o.Smallae) {
		return true
	}

	return false
}

// SetAe gets a reference to the given string and assigns it to the Ae field.
func (o *PmProAccountUpdate) SetSmallae(v string) {
	o.Smallae = &v
}

// GetIm returns the Im field value if set, zero value otherwise.
func (o *PmProAccountUpdate) GetSmallim() string {
	if o == nil || common.IsNil(o.Smallim) {
		var ret string
		return ret
	}
	return *o.Smallim
}

// GetImOk returns a tuple with the Im field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmProAccountUpdate) GetSmallimOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallim) {
		return nil, false
	}
	return o.Smallim, true
}

// HasIm returns a boolean if a field has been set.
func (o *PmProAccountUpdate) HasSmallim() bool {
	if o != nil && !common.IsNil(o.Smallim) {
		return true
	}

	return false
}

// SetIm gets a reference to the given string and assigns it to the Im field.
func (o *PmProAccountUpdate) SetSmallim(v string) {
	o.Smallim = &v
}

// GetMm returns the Mm field value if set, zero value otherwise.
func (o *PmProAccountUpdate) GetSmallmm() string {
	if o == nil || common.IsNil(o.Smallmm) {
		var ret string
		return ret
	}
	return *o.Smallmm
}

// GetMmOk returns a tuple with the Mm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmProAccountUpdate) GetSmallmmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallmm) {
		return nil, false
	}
	return o.Smallmm, true
}

// HasMm returns a boolean if a field has been set.
func (o *PmProAccountUpdate) HasSmallmm() bool {
	if o != nil && !common.IsNil(o.Smallmm) {
		return true
	}

	return false
}

// SetMm gets a reference to the given string and assigns it to the Mm field.
func (o *PmProAccountUpdate) SetSmallmm(v string) {
	o.Smallmm = &v
}

// GetAvb returns the Avb field value if set, zero value otherwise.
func (o *PmProAccountUpdate) GetAvb() string {
	if o == nil || common.IsNil(o.Avb) {
		var ret string
		return ret
	}
	return *o.Avb
}

// GetAvbOk returns a tuple with the Avb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmProAccountUpdate) GetAvbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Avb) {
		return nil, false
	}
	return o.Avb, true
}

// HasAvb returns a boolean if a field has been set.
func (o *PmProAccountUpdate) HasAvb() bool {
	if o != nil && !common.IsNil(o.Avb) {
		return true
	}

	return false
}

// SetAvb gets a reference to the given string and assigns it to the Avb field.
func (o *PmProAccountUpdate) SetAvb(v string) {
	o.Avb = &v
}

// GetVmw returns the Vmw field value if set, zero value otherwise.
func (o *PmProAccountUpdate) GetVmw() string {
	if o == nil || common.IsNil(o.Vmw) {
		var ret string
		return ret
	}
	return *o.Vmw
}

// GetVmwOk returns a tuple with the Vmw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmProAccountUpdate) GetVmwOk() (*string, bool) {
	if o == nil || common.IsNil(o.Vmw) {
		return nil, false
	}
	return o.Vmw, true
}

// HasVmw returns a boolean if a field has been set.
func (o *PmProAccountUpdate) HasVmw() bool {
	if o != nil && !common.IsNil(o.Vmw) {
		return true
	}

	return false
}

// SetVmw gets a reference to the given string and assigns it to the Vmw field.
func (o *PmProAccountUpdate) SetVmw(v string) {
	o.Vmw = &v
}

func (o PmProAccountUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmProAccountUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.Smalleq) {
		toSerialize["eq"] = o.Smalleq
	}
	if !common.IsNil(o.Smallae) {
		toSerialize["ae"] = o.Smallae
	}
	if !common.IsNil(o.Smallim) {
		toSerialize["im"] = o.Smallim
	}
	if !common.IsNil(o.Smallmm) {
		toSerialize["mm"] = o.Smallmm
	}
	if !common.IsNil(o.Avb) {
		toSerialize["avb"] = o.Avb
	}
	if !common.IsNil(o.Vmw) {
		toSerialize["vmw"] = o.Vmw
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PmProAccountUpdate) UnmarshalJSON(data []byte) (err error) {
	varPmProAccountUpdate := _PmProAccountUpdate{}

	err = json.Unmarshal(data, &varPmProAccountUpdate)

	if err != nil {
		return err
	}

	*o = PmProAccountUpdate(varPmProAccountUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "u")
		delete(additionalProperties, "eq")
		delete(additionalProperties, "ae")
		delete(additionalProperties, "im")
		delete(additionalProperties, "mm")
		delete(additionalProperties, "avb")
		delete(additionalProperties, "vmw")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePmProAccountUpdate struct {
	value *PmProAccountUpdate
	isSet bool
}

func (v NullablePmProAccountUpdate) Get() *PmProAccountUpdate {
	return v.value
}

func (v *NullablePmProAccountUpdate) Set(val *PmProAccountUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePmProAccountUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePmProAccountUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmProAccountUpdate(val *PmProAccountUpdate) *NullablePmProAccountUpdate {
	return &NullablePmProAccountUpdate{value: val, isSet: true}
}

func (v NullablePmProAccountUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmProAccountUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
