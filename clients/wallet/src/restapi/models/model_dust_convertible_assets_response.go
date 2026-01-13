/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DustConvertibleAssetsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DustConvertibleAssetsResponse{}

// DustConvertibleAssetsResponse struct for DustConvertibleAssetsResponse
type DustConvertibleAssetsResponse struct {
	DribbletPercentage             *string                                     `json:"dribbletPercentage,omitempty"`
	TotalTransferQuotaAssetAmount  *string                                     `json:"totalTransferQuotaAssetAmount,omitempty"`
	TotalTransferTargetAssetAmount *string                                     `json:"totalTransferTargetAssetAmount,omitempty"`
	DribbletBase                   *string                                     `json:"dribbletBase,omitempty"`
	Details                        []DustConvertibleAssetsResponseDetailsInner `json:"details,omitempty"`
	AdditionalProperties           map[string]interface{}
}

type _DustConvertibleAssetsResponse DustConvertibleAssetsResponse

// NewDustConvertibleAssetsResponse instantiates a new DustConvertibleAssetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDustConvertibleAssetsResponse() *DustConvertibleAssetsResponse {
	this := DustConvertibleAssetsResponse{}
	return &this
}

// NewDustConvertibleAssetsResponseWithDefaults instantiates a new DustConvertibleAssetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDustConvertibleAssetsResponseWithDefaults() *DustConvertibleAssetsResponse {
	this := DustConvertibleAssetsResponse{}
	return &this
}

// GetDribbletPercentage returns the DribbletPercentage field value if set, zero value otherwise.
func (o *DustConvertibleAssetsResponse) GetDribbletPercentage() string {
	if o == nil || common.IsNil(o.DribbletPercentage) {
		var ret string
		return ret
	}
	return *o.DribbletPercentage
}

// GetDribbletPercentageOk returns a tuple with the DribbletPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertibleAssetsResponse) GetDribbletPercentageOk() (*string, bool) {
	if o == nil || common.IsNil(o.DribbletPercentage) {
		return nil, false
	}
	return o.DribbletPercentage, true
}

// HasDribbletPercentage returns a boolean if a field has been set.
func (o *DustConvertibleAssetsResponse) HasDribbletPercentage() bool {
	if o != nil && !common.IsNil(o.DribbletPercentage) {
		return true
	}

	return false
}

// SetDribbletPercentage gets a reference to the given string and assigns it to the DribbletPercentage field.
func (o *DustConvertibleAssetsResponse) SetDribbletPercentage(v string) {
	o.DribbletPercentage = &v
}

// GetTotalTransferQuotaAssetAmount returns the TotalTransferQuotaAssetAmount field value if set, zero value otherwise.
func (o *DustConvertibleAssetsResponse) GetTotalTransferQuotaAssetAmount() string {
	if o == nil || common.IsNil(o.TotalTransferQuotaAssetAmount) {
		var ret string
		return ret
	}
	return *o.TotalTransferQuotaAssetAmount
}

// GetTotalTransferQuotaAssetAmountOk returns a tuple with the TotalTransferQuotaAssetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertibleAssetsResponse) GetTotalTransferQuotaAssetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalTransferQuotaAssetAmount) {
		return nil, false
	}
	return o.TotalTransferQuotaAssetAmount, true
}

// HasTotalTransferQuotaAssetAmount returns a boolean if a field has been set.
func (o *DustConvertibleAssetsResponse) HasTotalTransferQuotaAssetAmount() bool {
	if o != nil && !common.IsNil(o.TotalTransferQuotaAssetAmount) {
		return true
	}

	return false
}

// SetTotalTransferQuotaAssetAmount gets a reference to the given string and assigns it to the TotalTransferQuotaAssetAmount field.
func (o *DustConvertibleAssetsResponse) SetTotalTransferQuotaAssetAmount(v string) {
	o.TotalTransferQuotaAssetAmount = &v
}

// GetTotalTransferTargetAssetAmount returns the TotalTransferTargetAssetAmount field value if set, zero value otherwise.
func (o *DustConvertibleAssetsResponse) GetTotalTransferTargetAssetAmount() string {
	if o == nil || common.IsNil(o.TotalTransferTargetAssetAmount) {
		var ret string
		return ret
	}
	return *o.TotalTransferTargetAssetAmount
}

// GetTotalTransferTargetAssetAmountOk returns a tuple with the TotalTransferTargetAssetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertibleAssetsResponse) GetTotalTransferTargetAssetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalTransferTargetAssetAmount) {
		return nil, false
	}
	return o.TotalTransferTargetAssetAmount, true
}

// HasTotalTransferTargetAssetAmount returns a boolean if a field has been set.
func (o *DustConvertibleAssetsResponse) HasTotalTransferTargetAssetAmount() bool {
	if o != nil && !common.IsNil(o.TotalTransferTargetAssetAmount) {
		return true
	}

	return false
}

// SetTotalTransferTargetAssetAmount gets a reference to the given string and assigns it to the TotalTransferTargetAssetAmount field.
func (o *DustConvertibleAssetsResponse) SetTotalTransferTargetAssetAmount(v string) {
	o.TotalTransferTargetAssetAmount = &v
}

// GetDribbletBase returns the DribbletBase field value if set, zero value otherwise.
func (o *DustConvertibleAssetsResponse) GetDribbletBase() string {
	if o == nil || common.IsNil(o.DribbletBase) {
		var ret string
		return ret
	}
	return *o.DribbletBase
}

// GetDribbletBaseOk returns a tuple with the DribbletBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertibleAssetsResponse) GetDribbletBaseOk() (*string, bool) {
	if o == nil || common.IsNil(o.DribbletBase) {
		return nil, false
	}
	return o.DribbletBase, true
}

// HasDribbletBase returns a boolean if a field has been set.
func (o *DustConvertibleAssetsResponse) HasDribbletBase() bool {
	if o != nil && !common.IsNil(o.DribbletBase) {
		return true
	}

	return false
}

// SetDribbletBase gets a reference to the given string and assigns it to the DribbletBase field.
func (o *DustConvertibleAssetsResponse) SetDribbletBase(v string) {
	o.DribbletBase = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *DustConvertibleAssetsResponse) GetDetails() []DustConvertibleAssetsResponseDetailsInner {
	if o == nil || common.IsNil(o.Details) {
		var ret []DustConvertibleAssetsResponseDetailsInner
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertibleAssetsResponse) GetDetailsOk() ([]DustConvertibleAssetsResponseDetailsInner, bool) {
	if o == nil || common.IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *DustConvertibleAssetsResponse) HasDetails() bool {
	if o != nil && !common.IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []DustConvertibleAssetsResponseDetailsInner and assigns it to the Details field.
func (o *DustConvertibleAssetsResponse) SetDetails(v []DustConvertibleAssetsResponseDetailsInner) {
	o.Details = v
}

func (o DustConvertibleAssetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DustConvertibleAssetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DribbletPercentage) {
		toSerialize["dribbletPercentage"] = o.DribbletPercentage
	}
	if !common.IsNil(o.TotalTransferQuotaAssetAmount) {
		toSerialize["totalTransferQuotaAssetAmount"] = o.TotalTransferQuotaAssetAmount
	}
	if !common.IsNil(o.TotalTransferTargetAssetAmount) {
		toSerialize["totalTransferTargetAssetAmount"] = o.TotalTransferTargetAssetAmount
	}
	if !common.IsNil(o.DribbletBase) {
		toSerialize["dribbletBase"] = o.DribbletBase
	}
	if !common.IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DustConvertibleAssetsResponse) UnmarshalJSON(data []byte) (err error) {
	varDustConvertibleAssetsResponse := _DustConvertibleAssetsResponse{}

	err = json.Unmarshal(data, &varDustConvertibleAssetsResponse)

	if err != nil {
		return err
	}

	*o = DustConvertibleAssetsResponse(varDustConvertibleAssetsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dribbletPercentage")
		delete(additionalProperties, "totalTransferQuotaAssetAmount")
		delete(additionalProperties, "totalTransferTargetAssetAmount")
		delete(additionalProperties, "dribbletBase")
		delete(additionalProperties, "details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDustConvertibleAssetsResponse struct {
	value *DustConvertibleAssetsResponse
	isSet bool
}

func (v NullableDustConvertibleAssetsResponse) Get() *DustConvertibleAssetsResponse {
	return v.value
}

func (v *NullableDustConvertibleAssetsResponse) Set(val *DustConvertibleAssetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDustConvertibleAssetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDustConvertibleAssetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDustConvertibleAssetsResponse(val *DustConvertibleAssetsResponse) *NullableDustConvertibleAssetsResponse {
	return &NullableDustConvertibleAssetsResponse{value: val, isSet: true}
}

func (v NullableDustConvertibleAssetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDustConvertibleAssetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
