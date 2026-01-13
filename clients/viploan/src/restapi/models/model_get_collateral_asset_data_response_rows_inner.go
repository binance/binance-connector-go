/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetCollateralAssetDataResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCollateralAssetDataResponseRowsInner{}

// GetCollateralAssetDataResponseRowsInner struct for GetCollateralAssetDataResponseRowsInner
type GetCollateralAssetDataResponseRowsInner struct {
	CollateralCoin        *string `json:"collateralCoin,omitempty"`
	Var1stCollateralRatio *string `json:"_1stCollateralRatio,omitempty"`
	Var1stCollateralRange *string `json:"_1stCollateralRange,omitempty"`
	Var2ndCollateralRatio *string `json:"_2ndCollateralRatio,omitempty"`
	Var2ndCollateralRange *string `json:"_2ndCollateralRange,omitempty"`
	Var3rdCollateralRatio *string `json:"_3rdCollateralRatio,omitempty"`
	Var3rdCollateralRange *string `json:"_3rdCollateralRange,omitempty"`
	Var4thCollateralRatio *string `json:"_4thCollateralRatio,omitempty"`
	Var4thCollateralRange *string `json:"_4thCollateralRange,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _GetCollateralAssetDataResponseRowsInner GetCollateralAssetDataResponseRowsInner

// NewGetCollateralAssetDataResponseRowsInner instantiates a new GetCollateralAssetDataResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCollateralAssetDataResponseRowsInner() *GetCollateralAssetDataResponseRowsInner {
	this := GetCollateralAssetDataResponseRowsInner{}
	return &this
}

// NewGetCollateralAssetDataResponseRowsInnerWithDefaults instantiates a new GetCollateralAssetDataResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCollateralAssetDataResponseRowsInnerWithDefaults() *GetCollateralAssetDataResponseRowsInner {
	this := GetCollateralAssetDataResponseRowsInner{}
	return &this
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetCollateralAssetDataResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralAssetDataResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetCollateralAssetDataResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetCollateralAssetDataResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetVar1stCollateralRatio returns the Var1stCollateralRatio field value if set, zero value otherwise.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar1stCollateralRatio() string {
	if o == nil || common.IsNil(o.Var1stCollateralRatio) {
		var ret string
		return ret
	}
	return *o.Var1stCollateralRatio
}

// GetVar1stCollateralRatioOk returns a tuple with the Var1stCollateralRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar1stCollateralRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.Var1stCollateralRatio) {
		return nil, false
	}
	return o.Var1stCollateralRatio, true
}

// HasVar1stCollateralRatio returns a boolean if a field has been set.
func (o *GetCollateralAssetDataResponseRowsInner) HasVar1stCollateralRatio() bool {
	if o != nil && !common.IsNil(o.Var1stCollateralRatio) {
		return true
	}

	return false
}

// SetVar1stCollateralRatio gets a reference to the given string and assigns it to the Var1stCollateralRatio field.
func (o *GetCollateralAssetDataResponseRowsInner) SetVar1stCollateralRatio(v string) {
	o.Var1stCollateralRatio = &v
}

// GetVar1stCollateralRange returns the Var1stCollateralRange field value if set, zero value otherwise.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar1stCollateralRange() string {
	if o == nil || common.IsNil(o.Var1stCollateralRange) {
		var ret string
		return ret
	}
	return *o.Var1stCollateralRange
}

// GetVar1stCollateralRangeOk returns a tuple with the Var1stCollateralRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar1stCollateralRangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Var1stCollateralRange) {
		return nil, false
	}
	return o.Var1stCollateralRange, true
}

// HasVar1stCollateralRange returns a boolean if a field has been set.
func (o *GetCollateralAssetDataResponseRowsInner) HasVar1stCollateralRange() bool {
	if o != nil && !common.IsNil(o.Var1stCollateralRange) {
		return true
	}

	return false
}

// SetVar1stCollateralRange gets a reference to the given string and assigns it to the Var1stCollateralRange field.
func (o *GetCollateralAssetDataResponseRowsInner) SetVar1stCollateralRange(v string) {
	o.Var1stCollateralRange = &v
}

// GetVar2ndCollateralRatio returns the Var2ndCollateralRatio field value if set, zero value otherwise.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar2ndCollateralRatio() string {
	if o == nil || common.IsNil(o.Var2ndCollateralRatio) {
		var ret string
		return ret
	}
	return *o.Var2ndCollateralRatio
}

// GetVar2ndCollateralRatioOk returns a tuple with the Var2ndCollateralRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar2ndCollateralRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.Var2ndCollateralRatio) {
		return nil, false
	}
	return o.Var2ndCollateralRatio, true
}

// HasVar2ndCollateralRatio returns a boolean if a field has been set.
func (o *GetCollateralAssetDataResponseRowsInner) HasVar2ndCollateralRatio() bool {
	if o != nil && !common.IsNil(o.Var2ndCollateralRatio) {
		return true
	}

	return false
}

// SetVar2ndCollateralRatio gets a reference to the given string and assigns it to the Var2ndCollateralRatio field.
func (o *GetCollateralAssetDataResponseRowsInner) SetVar2ndCollateralRatio(v string) {
	o.Var2ndCollateralRatio = &v
}

// GetVar2ndCollateralRange returns the Var2ndCollateralRange field value if set, zero value otherwise.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar2ndCollateralRange() string {
	if o == nil || common.IsNil(o.Var2ndCollateralRange) {
		var ret string
		return ret
	}
	return *o.Var2ndCollateralRange
}

// GetVar2ndCollateralRangeOk returns a tuple with the Var2ndCollateralRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar2ndCollateralRangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Var2ndCollateralRange) {
		return nil, false
	}
	return o.Var2ndCollateralRange, true
}

// HasVar2ndCollateralRange returns a boolean if a field has been set.
func (o *GetCollateralAssetDataResponseRowsInner) HasVar2ndCollateralRange() bool {
	if o != nil && !common.IsNil(o.Var2ndCollateralRange) {
		return true
	}

	return false
}

// SetVar2ndCollateralRange gets a reference to the given string and assigns it to the Var2ndCollateralRange field.
func (o *GetCollateralAssetDataResponseRowsInner) SetVar2ndCollateralRange(v string) {
	o.Var2ndCollateralRange = &v
}

// GetVar3rdCollateralRatio returns the Var3rdCollateralRatio field value if set, zero value otherwise.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar3rdCollateralRatio() string {
	if o == nil || common.IsNil(o.Var3rdCollateralRatio) {
		var ret string
		return ret
	}
	return *o.Var3rdCollateralRatio
}

// GetVar3rdCollateralRatioOk returns a tuple with the Var3rdCollateralRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar3rdCollateralRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.Var3rdCollateralRatio) {
		return nil, false
	}
	return o.Var3rdCollateralRatio, true
}

// HasVar3rdCollateralRatio returns a boolean if a field has been set.
func (o *GetCollateralAssetDataResponseRowsInner) HasVar3rdCollateralRatio() bool {
	if o != nil && !common.IsNil(o.Var3rdCollateralRatio) {
		return true
	}

	return false
}

// SetVar3rdCollateralRatio gets a reference to the given string and assigns it to the Var3rdCollateralRatio field.
func (o *GetCollateralAssetDataResponseRowsInner) SetVar3rdCollateralRatio(v string) {
	o.Var3rdCollateralRatio = &v
}

// GetVar3rdCollateralRange returns the Var3rdCollateralRange field value if set, zero value otherwise.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar3rdCollateralRange() string {
	if o == nil || common.IsNil(o.Var3rdCollateralRange) {
		var ret string
		return ret
	}
	return *o.Var3rdCollateralRange
}

// GetVar3rdCollateralRangeOk returns a tuple with the Var3rdCollateralRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar3rdCollateralRangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Var3rdCollateralRange) {
		return nil, false
	}
	return o.Var3rdCollateralRange, true
}

// HasVar3rdCollateralRange returns a boolean if a field has been set.
func (o *GetCollateralAssetDataResponseRowsInner) HasVar3rdCollateralRange() bool {
	if o != nil && !common.IsNil(o.Var3rdCollateralRange) {
		return true
	}

	return false
}

// SetVar3rdCollateralRange gets a reference to the given string and assigns it to the Var3rdCollateralRange field.
func (o *GetCollateralAssetDataResponseRowsInner) SetVar3rdCollateralRange(v string) {
	o.Var3rdCollateralRange = &v
}

// GetVar4thCollateralRatio returns the Var4thCollateralRatio field value if set, zero value otherwise.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar4thCollateralRatio() string {
	if o == nil || common.IsNil(o.Var4thCollateralRatio) {
		var ret string
		return ret
	}
	return *o.Var4thCollateralRatio
}

// GetVar4thCollateralRatioOk returns a tuple with the Var4thCollateralRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar4thCollateralRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.Var4thCollateralRatio) {
		return nil, false
	}
	return o.Var4thCollateralRatio, true
}

// HasVar4thCollateralRatio returns a boolean if a field has been set.
func (o *GetCollateralAssetDataResponseRowsInner) HasVar4thCollateralRatio() bool {
	if o != nil && !common.IsNil(o.Var4thCollateralRatio) {
		return true
	}

	return false
}

// SetVar4thCollateralRatio gets a reference to the given string and assigns it to the Var4thCollateralRatio field.
func (o *GetCollateralAssetDataResponseRowsInner) SetVar4thCollateralRatio(v string) {
	o.Var4thCollateralRatio = &v
}

// GetVar4thCollateralRange returns the Var4thCollateralRange field value if set, zero value otherwise.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar4thCollateralRange() string {
	if o == nil || common.IsNil(o.Var4thCollateralRange) {
		var ret string
		return ret
	}
	return *o.Var4thCollateralRange
}

// GetVar4thCollateralRangeOk returns a tuple with the Var4thCollateralRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralAssetDataResponseRowsInner) GetVar4thCollateralRangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Var4thCollateralRange) {
		return nil, false
	}
	return o.Var4thCollateralRange, true
}

// HasVar4thCollateralRange returns a boolean if a field has been set.
func (o *GetCollateralAssetDataResponseRowsInner) HasVar4thCollateralRange() bool {
	if o != nil && !common.IsNil(o.Var4thCollateralRange) {
		return true
	}

	return false
}

// SetVar4thCollateralRange gets a reference to the given string and assigns it to the Var4thCollateralRange field.
func (o *GetCollateralAssetDataResponseRowsInner) SetVar4thCollateralRange(v string) {
	o.Var4thCollateralRange = &v
}

func (o GetCollateralAssetDataResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCollateralAssetDataResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.Var1stCollateralRatio) {
		toSerialize["_1stCollateralRatio"] = o.Var1stCollateralRatio
	}
	if !common.IsNil(o.Var1stCollateralRange) {
		toSerialize["_1stCollateralRange"] = o.Var1stCollateralRange
	}
	if !common.IsNil(o.Var2ndCollateralRatio) {
		toSerialize["_2ndCollateralRatio"] = o.Var2ndCollateralRatio
	}
	if !common.IsNil(o.Var2ndCollateralRange) {
		toSerialize["_2ndCollateralRange"] = o.Var2ndCollateralRange
	}
	if !common.IsNil(o.Var3rdCollateralRatio) {
		toSerialize["_3rdCollateralRatio"] = o.Var3rdCollateralRatio
	}
	if !common.IsNil(o.Var3rdCollateralRange) {
		toSerialize["_3rdCollateralRange"] = o.Var3rdCollateralRange
	}
	if !common.IsNil(o.Var4thCollateralRatio) {
		toSerialize["_4thCollateralRatio"] = o.Var4thCollateralRatio
	}
	if !common.IsNil(o.Var4thCollateralRange) {
		toSerialize["_4thCollateralRange"] = o.Var4thCollateralRange
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCollateralAssetDataResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetCollateralAssetDataResponseRowsInner := _GetCollateralAssetDataResponseRowsInner{}

	err = json.Unmarshal(data, &varGetCollateralAssetDataResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetCollateralAssetDataResponseRowsInner(varGetCollateralAssetDataResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "_1stCollateralRatio")
		delete(additionalProperties, "_1stCollateralRange")
		delete(additionalProperties, "_2ndCollateralRatio")
		delete(additionalProperties, "_2ndCollateralRange")
		delete(additionalProperties, "_3rdCollateralRatio")
		delete(additionalProperties, "_3rdCollateralRange")
		delete(additionalProperties, "_4thCollateralRatio")
		delete(additionalProperties, "_4thCollateralRange")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCollateralAssetDataResponseRowsInner struct {
	value *GetCollateralAssetDataResponseRowsInner
	isSet bool
}

func (v NullableGetCollateralAssetDataResponseRowsInner) Get() *GetCollateralAssetDataResponseRowsInner {
	return v.value
}

func (v *NullableGetCollateralAssetDataResponseRowsInner) Set(val *GetCollateralAssetDataResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCollateralAssetDataResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCollateralAssetDataResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCollateralAssetDataResponseRowsInner(val *GetCollateralAssetDataResponseRowsInner) *NullableGetCollateralAssetDataResponseRowsInner {
	return &NullableGetCollateralAssetDataResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetCollateralAssetDataResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCollateralAssetDataResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
