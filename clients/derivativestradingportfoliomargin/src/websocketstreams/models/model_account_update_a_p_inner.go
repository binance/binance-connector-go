/*
Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountUpdateAPInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountUpdateAPInner{}

// AccountUpdateAPInner struct for AccountUpdateAPInner
type AccountUpdateAPInner struct {
	Smalls               *string `json:"s,omitempty"`
	Smallpa              *string `json:"pa,omitempty"`
	Smallep              *string `json:"ep,omitempty"`
	Smallcr              *string `json:"cr,omitempty"`
	Smallup              *string `json:"up,omitempty"`
	Smallps              *string `json:"ps,omitempty"`
	Bep                  *string `json:"bep,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUpdateAPInner AccountUpdateAPInner

// NewAccountUpdateAPInner instantiates a new AccountUpdateAPInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdateAPInner() *AccountUpdateAPInner {
	this := AccountUpdateAPInner{}
	return &this
}

// NewAccountUpdateAPInnerWithDefaults instantiates a new AccountUpdateAPInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateAPInnerWithDefaults() *AccountUpdateAPInner {
	this := AccountUpdateAPInner{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AccountUpdateAPInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateAPInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AccountUpdateAPInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AccountUpdateAPInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetPa returns the Pa field value if set, zero value otherwise.
func (o *AccountUpdateAPInner) GetSmallpa() string {
	if o == nil || common.IsNil(o.Smallpa) {
		var ret string
		return ret
	}
	return *o.Smallpa
}

// GetPaOk returns a tuple with the Pa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateAPInner) GetSmallpaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallpa) {
		return nil, false
	}
	return o.Smallpa, true
}

// HasPa returns a boolean if a field has been set.
func (o *AccountUpdateAPInner) HasSmallpa() bool {
	if o != nil && !common.IsNil(o.Smallpa) {
		return true
	}

	return false
}

// SetPa gets a reference to the given string and assigns it to the Pa field.
func (o *AccountUpdateAPInner) SetSmallpa(v string) {
	o.Smallpa = &v
}

// GetEp returns the Ep field value if set, zero value otherwise.
func (o *AccountUpdateAPInner) GetSmallep() string {
	if o == nil || common.IsNil(o.Smallep) {
		var ret string
		return ret
	}
	return *o.Smallep
}

// GetEpOk returns a tuple with the Ep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateAPInner) GetSmallepOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallep) {
		return nil, false
	}
	return o.Smallep, true
}

// HasEp returns a boolean if a field has been set.
func (o *AccountUpdateAPInner) HasSmallep() bool {
	if o != nil && !common.IsNil(o.Smallep) {
		return true
	}

	return false
}

// SetEp gets a reference to the given string and assigns it to the Ep field.
func (o *AccountUpdateAPInner) SetSmallep(v string) {
	o.Smallep = &v
}

// GetCr returns the Cr field value if set, zero value otherwise.
func (o *AccountUpdateAPInner) GetSmallcr() string {
	if o == nil || common.IsNil(o.Smallcr) {
		var ret string
		return ret
	}
	return *o.Smallcr
}

// GetCrOk returns a tuple with the Cr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateAPInner) GetSmallcrOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallcr) {
		return nil, false
	}
	return o.Smallcr, true
}

// HasCr returns a boolean if a field has been set.
func (o *AccountUpdateAPInner) HasSmallcr() bool {
	if o != nil && !common.IsNil(o.Smallcr) {
		return true
	}

	return false
}

// SetCr gets a reference to the given string and assigns it to the Cr field.
func (o *AccountUpdateAPInner) SetSmallcr(v string) {
	o.Smallcr = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *AccountUpdateAPInner) GetSmallup() string {
	if o == nil || common.IsNil(o.Smallup) {
		var ret string
		return ret
	}
	return *o.Smallup
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateAPInner) GetSmallupOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallup) {
		return nil, false
	}
	return o.Smallup, true
}

// HasUp returns a boolean if a field has been set.
func (o *AccountUpdateAPInner) HasSmallup() bool {
	if o != nil && !common.IsNil(o.Smallup) {
		return true
	}

	return false
}

// SetUp gets a reference to the given string and assigns it to the Up field.
func (o *AccountUpdateAPInner) SetSmallup(v string) {
	o.Smallup = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *AccountUpdateAPInner) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateAPInner) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *AccountUpdateAPInner) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *AccountUpdateAPInner) SetSmallps(v string) {
	o.Smallps = &v
}

// GetBep returns the Bep field value if set, zero value otherwise.
func (o *AccountUpdateAPInner) GetBep() string {
	if o == nil || common.IsNil(o.Bep) {
		var ret string
		return ret
	}
	return *o.Bep
}

// GetBepOk returns a tuple with the Bep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateAPInner) GetBepOk() (*string, bool) {
	if o == nil || common.IsNil(o.Bep) {
		return nil, false
	}
	return o.Bep, true
}

// HasBep returns a boolean if a field has been set.
func (o *AccountUpdateAPInner) HasBep() bool {
	if o != nil && !common.IsNil(o.Bep) {
		return true
	}

	return false
}

// SetBep gets a reference to the given string and assigns it to the Bep field.
func (o *AccountUpdateAPInner) SetBep(v string) {
	o.Bep = &v
}

func (o AccountUpdateAPInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdateAPInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallpa) {
		toSerialize["pa"] = o.Smallpa
	}
	if !common.IsNil(o.Smallep) {
		toSerialize["ep"] = o.Smallep
	}
	if !common.IsNil(o.Smallcr) {
		toSerialize["cr"] = o.Smallcr
	}
	if !common.IsNil(o.Smallup) {
		toSerialize["up"] = o.Smallup
	}
	if !common.IsNil(o.Smallps) {
		toSerialize["ps"] = o.Smallps
	}
	if !common.IsNil(o.Bep) {
		toSerialize["bep"] = o.Bep
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountUpdateAPInner) UnmarshalJSON(data []byte) (err error) {
	varAccountUpdateAPInner := _AccountUpdateAPInner{}

	err = json.Unmarshal(data, &varAccountUpdateAPInner)

	if err != nil {
		return err
	}

	*o = AccountUpdateAPInner(varAccountUpdateAPInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "pa")
		delete(additionalProperties, "ep")
		delete(additionalProperties, "cr")
		delete(additionalProperties, "up")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "bep")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUpdateAPInner struct {
	value *AccountUpdateAPInner
	isSet bool
}

func (v NullableAccountUpdateAPInner) Get() *AccountUpdateAPInner {
	return v.value
}

func (v *NullableAccountUpdateAPInner) Set(val *AccountUpdateAPInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdateAPInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdateAPInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdateAPInner(val *AccountUpdateAPInner) *NullableAccountUpdateAPInner {
	return &NullableAccountUpdateAPInner{value: val, isSet: true}
}

func (v NullableAccountUpdateAPInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdateAPInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
