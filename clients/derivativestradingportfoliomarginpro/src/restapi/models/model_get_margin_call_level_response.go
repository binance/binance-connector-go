/*
Portfolio Margin Pro REST API

Access advanced account management and high-frequency trading with Binance Portfolio Margin Pro.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetMarginCallLevelResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMarginCallLevelResponse{}

// GetMarginCallLevelResponse struct for GetMarginCallLevelResponse
type GetMarginCallLevelResponse struct {
	// The margin call level value. Empty object returned if not set.
	MarginCallLevel      *string `json:"marginCallLevel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMarginCallLevelResponse GetMarginCallLevelResponse

// NewGetMarginCallLevelResponse instantiates a new GetMarginCallLevelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarginCallLevelResponse() *GetMarginCallLevelResponse {
	this := GetMarginCallLevelResponse{}
	return &this
}

// NewGetMarginCallLevelResponseWithDefaults instantiates a new GetMarginCallLevelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarginCallLevelResponseWithDefaults() *GetMarginCallLevelResponse {
	this := GetMarginCallLevelResponse{}
	return &this
}

// GetMarginCallLevel returns the MarginCallLevel field value if set, zero value otherwise.
func (o *GetMarginCallLevelResponse) GetMarginCallLevel() string {
	if o == nil || common.IsNil(o.MarginCallLevel) {
		var ret string
		return ret
	}
	return *o.MarginCallLevel
}

// GetMarginCallLevelOk returns a tuple with the MarginCallLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginCallLevelResponse) GetMarginCallLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginCallLevel) {
		return nil, false
	}
	return o.MarginCallLevel, true
}

// HasMarginCallLevel returns a boolean if a field has been set.
func (o *GetMarginCallLevelResponse) HasMarginCallLevel() bool {
	if o != nil && !common.IsNil(o.MarginCallLevel) {
		return true
	}

	return false
}

// SetMarginCallLevel gets a reference to the given string and assigns it to the MarginCallLevel field.
func (o *GetMarginCallLevelResponse) SetMarginCallLevel(v string) {
	o.MarginCallLevel = &v
}

func (o GetMarginCallLevelResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarginCallLevelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MarginCallLevel) {
		toSerialize["marginCallLevel"] = o.MarginCallLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMarginCallLevelResponse) UnmarshalJSON(data []byte) (err error) {
	varGetMarginCallLevelResponse := _GetMarginCallLevelResponse{}

	err = json.Unmarshal(data, &varGetMarginCallLevelResponse)

	if err != nil {
		return err
	}

	*o = GetMarginCallLevelResponse(varGetMarginCallLevelResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "marginCallLevel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMarginCallLevelResponse struct {
	value *GetMarginCallLevelResponse
	isSet bool
}

func (v NullableGetMarginCallLevelResponse) Get() *GetMarginCallLevelResponse {
	return v.value
}

func (v *NullableGetMarginCallLevelResponse) Set(val *GetMarginCallLevelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarginCallLevelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarginCallLevelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarginCallLevelResponse(val *GetMarginCallLevelResponse) *NullableGetMarginCallLevelResponse {
	return &NullableGetMarginCallLevelResponse{value: val, isSet: true}
}

func (v NullableGetMarginCallLevelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarginCallLevelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
