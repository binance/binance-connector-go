/*
Portfolio Margin Pro REST API

Access advanced account management and high-frequency trading with Binance Portfolio Margin Pro.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SetMarginCallLevelResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SetMarginCallLevelResponse{}

// SetMarginCallLevelResponse struct for SetMarginCallLevelResponse
type SetMarginCallLevelResponse struct {
	// The margin call level that was set
	MarginCallLevel      *string `json:"marginCallLevel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetMarginCallLevelResponse SetMarginCallLevelResponse

// NewSetMarginCallLevelResponse instantiates a new SetMarginCallLevelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetMarginCallLevelResponse() *SetMarginCallLevelResponse {
	this := SetMarginCallLevelResponse{}
	return &this
}

// NewSetMarginCallLevelResponseWithDefaults instantiates a new SetMarginCallLevelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetMarginCallLevelResponseWithDefaults() *SetMarginCallLevelResponse {
	this := SetMarginCallLevelResponse{}
	return &this
}

// GetMarginCallLevel returns the MarginCallLevel field value if set, zero value otherwise.
func (o *SetMarginCallLevelResponse) GetMarginCallLevel() string {
	if o == nil || common.IsNil(o.MarginCallLevel) {
		var ret string
		return ret
	}
	return *o.MarginCallLevel
}

// GetMarginCallLevelOk returns a tuple with the MarginCallLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetMarginCallLevelResponse) GetMarginCallLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginCallLevel) {
		return nil, false
	}
	return o.MarginCallLevel, true
}

// HasMarginCallLevel returns a boolean if a field has been set.
func (o *SetMarginCallLevelResponse) HasMarginCallLevel() bool {
	if o != nil && !common.IsNil(o.MarginCallLevel) {
		return true
	}

	return false
}

// SetMarginCallLevel gets a reference to the given string and assigns it to the MarginCallLevel field.
func (o *SetMarginCallLevelResponse) SetMarginCallLevel(v string) {
	o.MarginCallLevel = &v
}

func (o SetMarginCallLevelResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetMarginCallLevelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MarginCallLevel) {
		toSerialize["marginCallLevel"] = o.MarginCallLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetMarginCallLevelResponse) UnmarshalJSON(data []byte) (err error) {
	varSetMarginCallLevelResponse := _SetMarginCallLevelResponse{}

	err = json.Unmarshal(data, &varSetMarginCallLevelResponse)

	if err != nil {
		return err
	}

	*o = SetMarginCallLevelResponse(varSetMarginCallLevelResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "marginCallLevel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetMarginCallLevelResponse struct {
	value *SetMarginCallLevelResponse
	isSet bool
}

func (v NullableSetMarginCallLevelResponse) Get() *SetMarginCallLevelResponse {
	return v.value
}

func (v *NullableSetMarginCallLevelResponse) Set(val *SetMarginCallLevelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSetMarginCallLevelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSetMarginCallLevelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetMarginCallLevelResponse(val *SetMarginCallLevelResponse) *NullableSetMarginCallLevelResponse {
	return &NullableSetMarginCallLevelResponse{value: val, isSet: true}
}

func (v NullableSetMarginCallLevelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetMarginCallLevelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
