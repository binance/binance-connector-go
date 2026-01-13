/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarkPriceStreamForAllMarketResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarkPriceStreamForAllMarketResponseInner{}

// MarkPriceStreamForAllMarketResponseInner struct for MarkPriceStreamForAllMarketResponseInner
type MarkPriceStreamForAllMarketResponseInner struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	Smalli               *string `json:"i,omitempty"`
	P                    *string `json:"P,omitempty"`
	Smallr               *string `json:"r,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarkPriceStreamForAllMarketResponseInner MarkPriceStreamForAllMarketResponseInner

// NewMarkPriceStreamForAllMarketResponseInner instantiates a new MarkPriceStreamForAllMarketResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkPriceStreamForAllMarketResponseInner() *MarkPriceStreamForAllMarketResponseInner {
	this := MarkPriceStreamForAllMarketResponseInner{}
	return &this
}

// NewMarkPriceStreamForAllMarketResponseInnerWithDefaults instantiates a new MarkPriceStreamForAllMarketResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkPriceStreamForAllMarketResponseInnerWithDefaults() *MarkPriceStreamForAllMarketResponseInner {
	this := MarkPriceStreamForAllMarketResponseInner{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MarkPriceStreamForAllMarketResponseInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *MarkPriceStreamForAllMarketResponseInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MarkPriceStreamForAllMarketResponseInner) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *MarkPriceStreamForAllMarketResponseInner) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *MarkPriceStreamForAllMarketResponseInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *MarkPriceStreamForAllMarketResponseInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *MarkPriceStreamForAllMarketResponseInner) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *MarkPriceStreamForAllMarketResponseInner) SetSmallp(v string) {
	o.Smallp = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *MarkPriceStreamForAllMarketResponseInner) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *MarkPriceStreamForAllMarketResponseInner) SetSmalli(v string) {
	o.Smalli = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *MarkPriceStreamForAllMarketResponseInner) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *MarkPriceStreamForAllMarketResponseInner) SetP(v string) {
	o.P = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *MarkPriceStreamForAllMarketResponseInner) GetSmallr() string {
	if o == nil || common.IsNil(o.Smallr) {
		var ret string
		return ret
	}
	return *o.Smallr
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) GetSmallrOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallr) {
		return nil, false
	}
	return o.Smallr, true
}

// HasR returns a boolean if a field has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) HasSmallr() bool {
	if o != nil && !common.IsNil(o.Smallr) {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *MarkPriceStreamForAllMarketResponseInner) SetSmallr(v string) {
	o.Smallr = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *MarkPriceStreamForAllMarketResponseInner) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *MarkPriceStreamForAllMarketResponseInner) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *MarkPriceStreamForAllMarketResponseInner) SetT(v int64) {
	o.T = &v
}

func (o MarkPriceStreamForAllMarketResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkPriceStreamForAllMarketResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.P) {
		toSerialize["P"] = o.P
	}
	if !common.IsNil(o.Smallr) {
		toSerialize["r"] = o.Smallr
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarkPriceStreamForAllMarketResponseInner) UnmarshalJSON(data []byte) (err error) {
	varMarkPriceStreamForAllMarketResponseInner := _MarkPriceStreamForAllMarketResponseInner{}

	err = json.Unmarshal(data, &varMarkPriceStreamForAllMarketResponseInner)

	if err != nil {
		return err
	}

	*o = MarkPriceStreamForAllMarketResponseInner(varMarkPriceStreamForAllMarketResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "p")
		delete(additionalProperties, "i")
		delete(additionalProperties, "P")
		delete(additionalProperties, "r")
		delete(additionalProperties, "T")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarkPriceStreamForAllMarketResponseInner struct {
	value *MarkPriceStreamForAllMarketResponseInner
	isSet bool
}

func (v NullableMarkPriceStreamForAllMarketResponseInner) Get() *MarkPriceStreamForAllMarketResponseInner {
	return v.value
}

func (v *NullableMarkPriceStreamForAllMarketResponseInner) Set(val *MarkPriceStreamForAllMarketResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceStreamForAllMarketResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceStreamForAllMarketResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkPriceStreamForAllMarketResponseInner(val *MarkPriceStreamForAllMarketResponseInner) *NullableMarkPriceStreamForAllMarketResponseInner {
	return &NullableMarkPriceStreamForAllMarketResponseInner{value: val, isSet: true}
}

func (v NullableMarkPriceStreamForAllMarketResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceStreamForAllMarketResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
