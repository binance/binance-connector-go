/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetAssetsThatCanBeConvertedIntoBnbResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetAssetsThatCanBeConvertedIntoBnbResponse{}

// GetAssetsThatCanBeConvertedIntoBnbResponse struct for GetAssetsThatCanBeConvertedIntoBnbResponse
type GetAssetsThatCanBeConvertedIntoBnbResponse struct {
	Details              []GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner `json:"details,omitempty"`
	TotalTransferBtc     *string                                                  `json:"totalTransferBtc,omitempty"`
	TotalTransferBNB     *string                                                  `json:"totalTransferBNB,omitempty"`
	DribbletPercentage   *string                                                  `json:"dribbletPercentage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAssetsThatCanBeConvertedIntoBnbResponse GetAssetsThatCanBeConvertedIntoBnbResponse

// NewGetAssetsThatCanBeConvertedIntoBnbResponse instantiates a new GetAssetsThatCanBeConvertedIntoBnbResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetsThatCanBeConvertedIntoBnbResponse() *GetAssetsThatCanBeConvertedIntoBnbResponse {
	this := GetAssetsThatCanBeConvertedIntoBnbResponse{}
	return &this
}

// NewGetAssetsThatCanBeConvertedIntoBnbResponseWithDefaults instantiates a new GetAssetsThatCanBeConvertedIntoBnbResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetsThatCanBeConvertedIntoBnbResponseWithDefaults() *GetAssetsThatCanBeConvertedIntoBnbResponse {
	this := GetAssetsThatCanBeConvertedIntoBnbResponse{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) GetDetails() []GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner {
	if o == nil || common.IsNil(o.Details) {
		var ret []GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) GetDetailsOk() ([]GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner, bool) {
	if o == nil || common.IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) HasDetails() bool {
	if o != nil && !common.IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner and assigns it to the Details field.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) SetDetails(v []GetAssetsThatCanBeConvertedIntoBnbResponseDetailsInner) {
	o.Details = v
}

// GetTotalTransferBtc returns the TotalTransferBtc field value if set, zero value otherwise.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) GetTotalTransferBtc() string {
	if o == nil || common.IsNil(o.TotalTransferBtc) {
		var ret string
		return ret
	}
	return *o.TotalTransferBtc
}

// GetTotalTransferBtcOk returns a tuple with the TotalTransferBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) GetTotalTransferBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalTransferBtc) {
		return nil, false
	}
	return o.TotalTransferBtc, true
}

// HasTotalTransferBtc returns a boolean if a field has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) HasTotalTransferBtc() bool {
	if o != nil && !common.IsNil(o.TotalTransferBtc) {
		return true
	}

	return false
}

// SetTotalTransferBtc gets a reference to the given string and assigns it to the TotalTransferBtc field.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) SetTotalTransferBtc(v string) {
	o.TotalTransferBtc = &v
}

// GetTotalTransferBNB returns the TotalTransferBNB field value if set, zero value otherwise.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) GetTotalTransferBNB() string {
	if o == nil || common.IsNil(o.TotalTransferBNB) {
		var ret string
		return ret
	}
	return *o.TotalTransferBNB
}

// GetTotalTransferBNBOk returns a tuple with the TotalTransferBNB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) GetTotalTransferBNBOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalTransferBNB) {
		return nil, false
	}
	return o.TotalTransferBNB, true
}

// HasTotalTransferBNB returns a boolean if a field has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) HasTotalTransferBNB() bool {
	if o != nil && !common.IsNil(o.TotalTransferBNB) {
		return true
	}

	return false
}

// SetTotalTransferBNB gets a reference to the given string and assigns it to the TotalTransferBNB field.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) SetTotalTransferBNB(v string) {
	o.TotalTransferBNB = &v
}

// GetDribbletPercentage returns the DribbletPercentage field value if set, zero value otherwise.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) GetDribbletPercentage() string {
	if o == nil || common.IsNil(o.DribbletPercentage) {
		var ret string
		return ret
	}
	return *o.DribbletPercentage
}

// GetDribbletPercentageOk returns a tuple with the DribbletPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) GetDribbletPercentageOk() (*string, bool) {
	if o == nil || common.IsNil(o.DribbletPercentage) {
		return nil, false
	}
	return o.DribbletPercentage, true
}

// HasDribbletPercentage returns a boolean if a field has been set.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) HasDribbletPercentage() bool {
	if o != nil && !common.IsNil(o.DribbletPercentage) {
		return true
	}

	return false
}

// SetDribbletPercentage gets a reference to the given string and assigns it to the DribbletPercentage field.
func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) SetDribbletPercentage(v string) {
	o.DribbletPercentage = &v
}

func (o GetAssetsThatCanBeConvertedIntoBnbResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssetsThatCanBeConvertedIntoBnbResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !common.IsNil(o.TotalTransferBtc) {
		toSerialize["totalTransferBtc"] = o.TotalTransferBtc
	}
	if !common.IsNil(o.TotalTransferBNB) {
		toSerialize["totalTransferBNB"] = o.TotalTransferBNB
	}
	if !common.IsNil(o.DribbletPercentage) {
		toSerialize["dribbletPercentage"] = o.DribbletPercentage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAssetsThatCanBeConvertedIntoBnbResponse) UnmarshalJSON(data []byte) (err error) {
	varGetAssetsThatCanBeConvertedIntoBnbResponse := _GetAssetsThatCanBeConvertedIntoBnbResponse{}

	err = json.Unmarshal(data, &varGetAssetsThatCanBeConvertedIntoBnbResponse)

	if err != nil {
		return err
	}

	*o = GetAssetsThatCanBeConvertedIntoBnbResponse(varGetAssetsThatCanBeConvertedIntoBnbResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "details")
		delete(additionalProperties, "totalTransferBtc")
		delete(additionalProperties, "totalTransferBNB")
		delete(additionalProperties, "dribbletPercentage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAssetsThatCanBeConvertedIntoBnbResponse struct {
	value *GetAssetsThatCanBeConvertedIntoBnbResponse
	isSet bool
}

func (v NullableGetAssetsThatCanBeConvertedIntoBnbResponse) Get() *GetAssetsThatCanBeConvertedIntoBnbResponse {
	return v.value
}

func (v *NullableGetAssetsThatCanBeConvertedIntoBnbResponse) Set(val *GetAssetsThatCanBeConvertedIntoBnbResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetsThatCanBeConvertedIntoBnbResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetsThatCanBeConvertedIntoBnbResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetsThatCanBeConvertedIntoBnbResponse(val *GetAssetsThatCanBeConvertedIntoBnbResponse) *NullableGetAssetsThatCanBeConvertedIntoBnbResponse {
	return &NullableGetAssetsThatCanBeConvertedIntoBnbResponse{value: val, isSet: true}
}

func (v NullableGetAssetsThatCanBeConvertedIntoBnbResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetsThatCanBeConvertedIntoBnbResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
