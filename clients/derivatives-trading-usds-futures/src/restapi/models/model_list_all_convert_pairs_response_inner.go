/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ListAllConvertPairsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListAllConvertPairsResponseInner{}

// ListAllConvertPairsResponseInner struct for ListAllConvertPairsResponseInner
type ListAllConvertPairsResponseInner struct {
	FromAsset            *string `json:"fromAsset,omitempty"`
	ToAsset              *string `json:"toAsset,omitempty"`
	FromAssetMinAmount   *string `json:"fromAssetMinAmount,omitempty"`
	FromAssetMaxAmount   *string `json:"fromAssetMaxAmount,omitempty"`
	ToAssetMinAmount     *string `json:"toAssetMinAmount,omitempty"`
	ToAssetMaxAmount     *string `json:"toAssetMaxAmount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListAllConvertPairsResponseInner ListAllConvertPairsResponseInner

// NewListAllConvertPairsResponseInner instantiates a new ListAllConvertPairsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllConvertPairsResponseInner() *ListAllConvertPairsResponseInner {
	this := ListAllConvertPairsResponseInner{}
	return &this
}

// NewListAllConvertPairsResponseInnerWithDefaults instantiates a new ListAllConvertPairsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllConvertPairsResponseInnerWithDefaults() *ListAllConvertPairsResponseInner {
	this := ListAllConvertPairsResponseInner{}
	return &this
}

// GetFromAsset returns the FromAsset field value if set, zero value otherwise.
func (o *ListAllConvertPairsResponseInner) GetFromAsset() string {
	if o == nil || common.IsNil(o.FromAsset) {
		var ret string
		return ret
	}
	return *o.FromAsset
}

// GetFromAssetOk returns a tuple with the FromAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllConvertPairsResponseInner) GetFromAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAsset) {
		return nil, false
	}
	return o.FromAsset, true
}

// HasFromAsset returns a boolean if a field has been set.
func (o *ListAllConvertPairsResponseInner) HasFromAsset() bool {
	if o != nil && !common.IsNil(o.FromAsset) {
		return true
	}

	return false
}

// SetFromAsset gets a reference to the given string and assigns it to the FromAsset field.
func (o *ListAllConvertPairsResponseInner) SetFromAsset(v string) {
	o.FromAsset = &v
}

// GetToAsset returns the ToAsset field value if set, zero value otherwise.
func (o *ListAllConvertPairsResponseInner) GetToAsset() string {
	if o == nil || common.IsNil(o.ToAsset) {
		var ret string
		return ret
	}
	return *o.ToAsset
}

// GetToAssetOk returns a tuple with the ToAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllConvertPairsResponseInner) GetToAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAsset) {
		return nil, false
	}
	return o.ToAsset, true
}

// HasToAsset returns a boolean if a field has been set.
func (o *ListAllConvertPairsResponseInner) HasToAsset() bool {
	if o != nil && !common.IsNil(o.ToAsset) {
		return true
	}

	return false
}

// SetToAsset gets a reference to the given string and assigns it to the ToAsset field.
func (o *ListAllConvertPairsResponseInner) SetToAsset(v string) {
	o.ToAsset = &v
}

// GetFromAssetMinAmount returns the FromAssetMinAmount field value if set, zero value otherwise.
func (o *ListAllConvertPairsResponseInner) GetFromAssetMinAmount() string {
	if o == nil || common.IsNil(o.FromAssetMinAmount) {
		var ret string
		return ret
	}
	return *o.FromAssetMinAmount
}

// GetFromAssetMinAmountOk returns a tuple with the FromAssetMinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllConvertPairsResponseInner) GetFromAssetMinAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAssetMinAmount) {
		return nil, false
	}
	return o.FromAssetMinAmount, true
}

// HasFromAssetMinAmount returns a boolean if a field has been set.
func (o *ListAllConvertPairsResponseInner) HasFromAssetMinAmount() bool {
	if o != nil && !common.IsNil(o.FromAssetMinAmount) {
		return true
	}

	return false
}

// SetFromAssetMinAmount gets a reference to the given string and assigns it to the FromAssetMinAmount field.
func (o *ListAllConvertPairsResponseInner) SetFromAssetMinAmount(v string) {
	o.FromAssetMinAmount = &v
}

// GetFromAssetMaxAmount returns the FromAssetMaxAmount field value if set, zero value otherwise.
func (o *ListAllConvertPairsResponseInner) GetFromAssetMaxAmount() string {
	if o == nil || common.IsNil(o.FromAssetMaxAmount) {
		var ret string
		return ret
	}
	return *o.FromAssetMaxAmount
}

// GetFromAssetMaxAmountOk returns a tuple with the FromAssetMaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllConvertPairsResponseInner) GetFromAssetMaxAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAssetMaxAmount) {
		return nil, false
	}
	return o.FromAssetMaxAmount, true
}

// HasFromAssetMaxAmount returns a boolean if a field has been set.
func (o *ListAllConvertPairsResponseInner) HasFromAssetMaxAmount() bool {
	if o != nil && !common.IsNil(o.FromAssetMaxAmount) {
		return true
	}

	return false
}

// SetFromAssetMaxAmount gets a reference to the given string and assigns it to the FromAssetMaxAmount field.
func (o *ListAllConvertPairsResponseInner) SetFromAssetMaxAmount(v string) {
	o.FromAssetMaxAmount = &v
}

// GetToAssetMinAmount returns the ToAssetMinAmount field value if set, zero value otherwise.
func (o *ListAllConvertPairsResponseInner) GetToAssetMinAmount() string {
	if o == nil || common.IsNil(o.ToAssetMinAmount) {
		var ret string
		return ret
	}
	return *o.ToAssetMinAmount
}

// GetToAssetMinAmountOk returns a tuple with the ToAssetMinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllConvertPairsResponseInner) GetToAssetMinAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAssetMinAmount) {
		return nil, false
	}
	return o.ToAssetMinAmount, true
}

// HasToAssetMinAmount returns a boolean if a field has been set.
func (o *ListAllConvertPairsResponseInner) HasToAssetMinAmount() bool {
	if o != nil && !common.IsNil(o.ToAssetMinAmount) {
		return true
	}

	return false
}

// SetToAssetMinAmount gets a reference to the given string and assigns it to the ToAssetMinAmount field.
func (o *ListAllConvertPairsResponseInner) SetToAssetMinAmount(v string) {
	o.ToAssetMinAmount = &v
}

// GetToAssetMaxAmount returns the ToAssetMaxAmount field value if set, zero value otherwise.
func (o *ListAllConvertPairsResponseInner) GetToAssetMaxAmount() string {
	if o == nil || common.IsNil(o.ToAssetMaxAmount) {
		var ret string
		return ret
	}
	return *o.ToAssetMaxAmount
}

// GetToAssetMaxAmountOk returns a tuple with the ToAssetMaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllConvertPairsResponseInner) GetToAssetMaxAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAssetMaxAmount) {
		return nil, false
	}
	return o.ToAssetMaxAmount, true
}

// HasToAssetMaxAmount returns a boolean if a field has been set.
func (o *ListAllConvertPairsResponseInner) HasToAssetMaxAmount() bool {
	if o != nil && !common.IsNil(o.ToAssetMaxAmount) {
		return true
	}

	return false
}

// SetToAssetMaxAmount gets a reference to the given string and assigns it to the ToAssetMaxAmount field.
func (o *ListAllConvertPairsResponseInner) SetToAssetMaxAmount(v string) {
	o.ToAssetMaxAmount = &v
}

func (o ListAllConvertPairsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAllConvertPairsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FromAsset) {
		toSerialize["fromAsset"] = o.FromAsset
	}
	if !common.IsNil(o.ToAsset) {
		toSerialize["toAsset"] = o.ToAsset
	}
	if !common.IsNil(o.FromAssetMinAmount) {
		toSerialize["fromAssetMinAmount"] = o.FromAssetMinAmount
	}
	if !common.IsNil(o.FromAssetMaxAmount) {
		toSerialize["fromAssetMaxAmount"] = o.FromAssetMaxAmount
	}
	if !common.IsNil(o.ToAssetMinAmount) {
		toSerialize["toAssetMinAmount"] = o.ToAssetMinAmount
	}
	if !common.IsNil(o.ToAssetMaxAmount) {
		toSerialize["toAssetMaxAmount"] = o.ToAssetMaxAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListAllConvertPairsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varListAllConvertPairsResponseInner := _ListAllConvertPairsResponseInner{}

	err = json.Unmarshal(data, &varListAllConvertPairsResponseInner)

	if err != nil {
		return err
	}

	*o = ListAllConvertPairsResponseInner(varListAllConvertPairsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fromAsset")
		delete(additionalProperties, "toAsset")
		delete(additionalProperties, "fromAssetMinAmount")
		delete(additionalProperties, "fromAssetMaxAmount")
		delete(additionalProperties, "toAssetMinAmount")
		delete(additionalProperties, "toAssetMaxAmount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAllConvertPairsResponseInner struct {
	value *ListAllConvertPairsResponseInner
	isSet bool
}

func (v NullableListAllConvertPairsResponseInner) Get() *ListAllConvertPairsResponseInner {
	return v.value
}

func (v *NullableListAllConvertPairsResponseInner) Set(val *ListAllConvertPairsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllConvertPairsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllConvertPairsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllConvertPairsResponseInner(val *ListAllConvertPairsResponseInner) *NullableListAllConvertPairsResponseInner {
	return &NullableListAllConvertPairsResponseInner{value: val, isSet: true}
}

func (v NullableListAllConvertPairsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllConvertPairsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
