/*
Binance Rebate REST API

OpenAPI Specification for the Binance Rebate REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSpotRebateHistoryRecordsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSpotRebateHistoryRecordsResponse{}

// GetSpotRebateHistoryRecordsResponse struct for GetSpotRebateHistoryRecordsResponse
type GetSpotRebateHistoryRecordsResponse struct {
	Status               *string                                  `json:"status,omitempty"`
	Type                 *string                                  `json:"type,omitempty"`
	Code                 *string                                  `json:"code,omitempty"`
	Data                 *GetSpotRebateHistoryRecordsResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSpotRebateHistoryRecordsResponse GetSpotRebateHistoryRecordsResponse

// NewGetSpotRebateHistoryRecordsResponse instantiates a new GetSpotRebateHistoryRecordsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSpotRebateHistoryRecordsResponse() *GetSpotRebateHistoryRecordsResponse {
	this := GetSpotRebateHistoryRecordsResponse{}
	return &this
}

// NewGetSpotRebateHistoryRecordsResponseWithDefaults instantiates a new GetSpotRebateHistoryRecordsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSpotRebateHistoryRecordsResponseWithDefaults() *GetSpotRebateHistoryRecordsResponse {
	this := GetSpotRebateHistoryRecordsResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetSpotRebateHistoryRecordsResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotRebateHistoryRecordsResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetSpotRebateHistoryRecordsResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetSpotRebateHistoryRecordsResponse) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetSpotRebateHistoryRecordsResponse) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotRebateHistoryRecordsResponse) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetSpotRebateHistoryRecordsResponse) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetSpotRebateHistoryRecordsResponse) SetType(v string) {
	o.Type = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetSpotRebateHistoryRecordsResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotRebateHistoryRecordsResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetSpotRebateHistoryRecordsResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetSpotRebateHistoryRecordsResponse) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSpotRebateHistoryRecordsResponse) GetData() GetSpotRebateHistoryRecordsResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret GetSpotRebateHistoryRecordsResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotRebateHistoryRecordsResponse) GetDataOk() (*GetSpotRebateHistoryRecordsResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSpotRebateHistoryRecordsResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetSpotRebateHistoryRecordsResponseData and assigns it to the Data field.
func (o *GetSpotRebateHistoryRecordsResponse) SetData(v GetSpotRebateHistoryRecordsResponseData) {
	o.Data = &v
}

func (o GetSpotRebateHistoryRecordsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSpotRebateHistoryRecordsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSpotRebateHistoryRecordsResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSpotRebateHistoryRecordsResponse := _GetSpotRebateHistoryRecordsResponse{}

	err = json.Unmarshal(data, &varGetSpotRebateHistoryRecordsResponse)

	if err != nil {
		return err
	}

	*o = GetSpotRebateHistoryRecordsResponse(varGetSpotRebateHistoryRecordsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "code")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSpotRebateHistoryRecordsResponse struct {
	value *GetSpotRebateHistoryRecordsResponse
	isSet bool
}

func (v NullableGetSpotRebateHistoryRecordsResponse) Get() *GetSpotRebateHistoryRecordsResponse {
	return v.value
}

func (v *NullableGetSpotRebateHistoryRecordsResponse) Set(val *GetSpotRebateHistoryRecordsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSpotRebateHistoryRecordsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSpotRebateHistoryRecordsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSpotRebateHistoryRecordsResponse(val *GetSpotRebateHistoryRecordsResponse) *NullableGetSpotRebateHistoryRecordsResponse {
	return &NullableGetSpotRebateHistoryRecordsResponse{value: val, isSet: true}
}

func (v NullableGetSpotRebateHistoryRecordsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSpotRebateHistoryRecordsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
