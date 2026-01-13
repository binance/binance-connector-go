/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryManagedSubAccountFuturesAssetDetailsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountFuturesAssetDetailsResponse{}

// QueryManagedSubAccountFuturesAssetDetailsResponse struct for QueryManagedSubAccountFuturesAssetDetailsResponse
type QueryManagedSubAccountFuturesAssetDetailsResponse struct {
	Code                 *string                                                             `json:"code,omitempty"`
	Message              *string                                                             `json:"message,omitempty"`
	SnapshotVos          []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner `json:"snapshotVos,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountFuturesAssetDetailsResponse QueryManagedSubAccountFuturesAssetDetailsResponse

// NewQueryManagedSubAccountFuturesAssetDetailsResponse instantiates a new QueryManagedSubAccountFuturesAssetDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountFuturesAssetDetailsResponse() *QueryManagedSubAccountFuturesAssetDetailsResponse {
	this := QueryManagedSubAccountFuturesAssetDetailsResponse{}
	return &this
}

// NewQueryManagedSubAccountFuturesAssetDetailsResponseWithDefaults instantiates a new QueryManagedSubAccountFuturesAssetDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountFuturesAssetDetailsResponseWithDefaults() *QueryManagedSubAccountFuturesAssetDetailsResponse {
	this := QueryManagedSubAccountFuturesAssetDetailsResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) SetMessage(v string) {
	o.Message = &v
}

// GetSnapshotVos returns the SnapshotVos field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) GetSnapshotVos() []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner {
	if o == nil || common.IsNil(o.SnapshotVos) {
		var ret []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner
		return ret
	}
	return o.SnapshotVos
}

// GetSnapshotVosOk returns a tuple with the SnapshotVos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) GetSnapshotVosOk() ([]QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner, bool) {
	if o == nil || common.IsNil(o.SnapshotVos) {
		return nil, false
	}
	return o.SnapshotVos, true
}

// HasSnapshotVos returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) HasSnapshotVos() bool {
	if o != nil && !common.IsNil(o.SnapshotVos) {
		return true
	}

	return false
}

// SetSnapshotVos gets a reference to the given []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner and assigns it to the SnapshotVos field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) SetSnapshotVos(v []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) {
	o.SnapshotVos = v
}

func (o QueryManagedSubAccountFuturesAssetDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountFuturesAssetDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !common.IsNil(o.SnapshotVos) {
		toSerialize["snapshotVos"] = o.SnapshotVos
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryManagedSubAccountFuturesAssetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountFuturesAssetDetailsResponse := _QueryManagedSubAccountFuturesAssetDetailsResponse{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountFuturesAssetDetailsResponse)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountFuturesAssetDetailsResponse(varQueryManagedSubAccountFuturesAssetDetailsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "snapshotVos")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountFuturesAssetDetailsResponse struct {
	value *QueryManagedSubAccountFuturesAssetDetailsResponse
	isSet bool
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponse) Get() *QueryManagedSubAccountFuturesAssetDetailsResponse {
	return v.value
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponse) Set(val *QueryManagedSubAccountFuturesAssetDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountFuturesAssetDetailsResponse(val *QueryManagedSubAccountFuturesAssetDetailsResponse) *NullableQueryManagedSubAccountFuturesAssetDetailsResponse {
	return &NullableQueryManagedSubAccountFuturesAssetDetailsResponse{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
