/*
Portfolio Margin Pro REST API

Access advanced account management and high-frequency trading with Binance Portfolio Margin Pro.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DeleteMarginCallLevelResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DeleteMarginCallLevelResponse{}

// DeleteMarginCallLevelResponse struct for DeleteMarginCallLevelResponse
type DeleteMarginCallLevelResponse struct {
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteMarginCallLevelResponse DeleteMarginCallLevelResponse

// NewDeleteMarginCallLevelResponse instantiates a new DeleteMarginCallLevelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteMarginCallLevelResponse() *DeleteMarginCallLevelResponse {
	this := DeleteMarginCallLevelResponse{}
	return &this
}

// NewDeleteMarginCallLevelResponseWithDefaults instantiates a new DeleteMarginCallLevelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteMarginCallLevelResponseWithDefaults() *DeleteMarginCallLevelResponse {
	this := DeleteMarginCallLevelResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *DeleteMarginCallLevelResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteMarginCallLevelResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *DeleteMarginCallLevelResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *DeleteMarginCallLevelResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o DeleteMarginCallLevelResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteMarginCallLevelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteMarginCallLevelResponse) UnmarshalJSON(data []byte) (err error) {
	varDeleteMarginCallLevelResponse := _DeleteMarginCallLevelResponse{}

	err = json.Unmarshal(data, &varDeleteMarginCallLevelResponse)

	if err != nil {
		return err
	}

	*o = DeleteMarginCallLevelResponse(varDeleteMarginCallLevelResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteMarginCallLevelResponse struct {
	value *DeleteMarginCallLevelResponse
	isSet bool
}

func (v NullableDeleteMarginCallLevelResponse) Get() *DeleteMarginCallLevelResponse {
	return v.value
}

func (v *NullableDeleteMarginCallLevelResponse) Set(val *DeleteMarginCallLevelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteMarginCallLevelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteMarginCallLevelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteMarginCallLevelResponse(val *DeleteMarginCallLevelResponse) *NullableDeleteMarginCallLevelResponse {
	return &NullableDeleteMarginCallLevelResponse{value: val, isSet: true}
}

func (v NullableDeleteMarginCallLevelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteMarginCallLevelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
