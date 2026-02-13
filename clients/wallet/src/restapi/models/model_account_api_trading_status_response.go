/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountApiTradingStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountApiTradingStatusResponse{}

// AccountApiTradingStatusResponse struct for AccountApiTradingStatusResponse
type AccountApiTradingStatusResponse struct {
	Data                 *AccountApiTradingStatusResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountApiTradingStatusResponse AccountApiTradingStatusResponse

// NewAccountApiTradingStatusResponse instantiates a new AccountApiTradingStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountApiTradingStatusResponse() *AccountApiTradingStatusResponse {
	this := AccountApiTradingStatusResponse{}
	return &this
}

// NewAccountApiTradingStatusResponseWithDefaults instantiates a new AccountApiTradingStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountApiTradingStatusResponseWithDefaults() *AccountApiTradingStatusResponse {
	this := AccountApiTradingStatusResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AccountApiTradingStatusResponse) GetData() AccountApiTradingStatusResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret AccountApiTradingStatusResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountApiTradingStatusResponse) GetDataOk() (*AccountApiTradingStatusResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AccountApiTradingStatusResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AccountApiTradingStatusResponseData and assigns it to the Data field.
func (o *AccountApiTradingStatusResponse) SetData(v AccountApiTradingStatusResponseData) {
	o.Data = &v
}

func (o AccountApiTradingStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountApiTradingStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountApiTradingStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varAccountApiTradingStatusResponse := _AccountApiTradingStatusResponse{}

	err = json.Unmarshal(data, &varAccountApiTradingStatusResponse)

	if err != nil {
		return err
	}

	*o = AccountApiTradingStatusResponse(varAccountApiTradingStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountApiTradingStatusResponse struct {
	value *AccountApiTradingStatusResponse
	isSet bool
}

func (v NullableAccountApiTradingStatusResponse) Get() *AccountApiTradingStatusResponse {
	return v.value
}

func (v *NullableAccountApiTradingStatusResponse) Set(val *AccountApiTradingStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountApiTradingStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountApiTradingStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountApiTradingStatusResponse(val *AccountApiTradingStatusResponse) *NullableAccountApiTradingStatusResponse {
	return &NullableAccountApiTradingStatusResponse{value: val, isSet: true}
}

func (v NullableAccountApiTradingStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountApiTradingStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
