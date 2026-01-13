/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetApiKeyPermissionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetApiKeyPermissionResponse{}

// GetApiKeyPermissionResponse struct for GetApiKeyPermissionResponse
type GetApiKeyPermissionResponse struct {
	IpRestrict                   *bool  `json:"ipRestrict,omitempty"`
	CreateTime                   *int64 `json:"createTime,omitempty"`
	EnableReading                *bool  `json:"enableReading,omitempty"`
	EnableWithdrawals            *bool  `json:"enableWithdrawals,omitempty"`
	EnableInternalTransfer       *bool  `json:"enableInternalTransfer,omitempty"`
	EnableMargin                 *bool  `json:"enableMargin,omitempty"`
	EnableFutures                *bool  `json:"enableFutures,omitempty"`
	PermitsUniversalTransfer     *bool  `json:"permitsUniversalTransfer,omitempty"`
	EnableVanillaOptions         *bool  `json:"enableVanillaOptions,omitempty"`
	EnableFixApiTrade            *bool  `json:"enableFixApiTrade,omitempty"`
	EnableFixReadOnly            *bool  `json:"enableFixReadOnly,omitempty"`
	EnableSpotAndMarginTrading   *bool  `json:"enableSpotAndMarginTrading,omitempty"`
	EnablePortfolioMarginTrading *bool  `json:"enablePortfolioMarginTrading,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _GetApiKeyPermissionResponse GetApiKeyPermissionResponse

// NewGetApiKeyPermissionResponse instantiates a new GetApiKeyPermissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApiKeyPermissionResponse() *GetApiKeyPermissionResponse {
	this := GetApiKeyPermissionResponse{}
	return &this
}

// NewGetApiKeyPermissionResponseWithDefaults instantiates a new GetApiKeyPermissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApiKeyPermissionResponseWithDefaults() *GetApiKeyPermissionResponse {
	this := GetApiKeyPermissionResponse{}
	return &this
}

// GetIpRestrict returns the IpRestrict field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetIpRestrict() bool {
	if o == nil || common.IsNil(o.IpRestrict) {
		var ret bool
		return ret
	}
	return *o.IpRestrict
}

// GetIpRestrictOk returns a tuple with the IpRestrict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetIpRestrictOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IpRestrict) {
		return nil, false
	}
	return o.IpRestrict, true
}

// HasIpRestrict returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasIpRestrict() bool {
	if o != nil && !common.IsNil(o.IpRestrict) {
		return true
	}

	return false
}

// SetIpRestrict gets a reference to the given bool and assigns it to the IpRestrict field.
func (o *GetApiKeyPermissionResponse) SetIpRestrict(v bool) {
	o.IpRestrict = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *GetApiKeyPermissionResponse) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetEnableReading returns the EnableReading field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetEnableReading() bool {
	if o == nil || common.IsNil(o.EnableReading) {
		var ret bool
		return ret
	}
	return *o.EnableReading
}

// GetEnableReadingOk returns a tuple with the EnableReading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetEnableReadingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnableReading) {
		return nil, false
	}
	return o.EnableReading, true
}

// HasEnableReading returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasEnableReading() bool {
	if o != nil && !common.IsNil(o.EnableReading) {
		return true
	}

	return false
}

// SetEnableReading gets a reference to the given bool and assigns it to the EnableReading field.
func (o *GetApiKeyPermissionResponse) SetEnableReading(v bool) {
	o.EnableReading = &v
}

// GetEnableWithdrawals returns the EnableWithdrawals field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetEnableWithdrawals() bool {
	if o == nil || common.IsNil(o.EnableWithdrawals) {
		var ret bool
		return ret
	}
	return *o.EnableWithdrawals
}

// GetEnableWithdrawalsOk returns a tuple with the EnableWithdrawals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetEnableWithdrawalsOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnableWithdrawals) {
		return nil, false
	}
	return o.EnableWithdrawals, true
}

// HasEnableWithdrawals returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasEnableWithdrawals() bool {
	if o != nil && !common.IsNil(o.EnableWithdrawals) {
		return true
	}

	return false
}

// SetEnableWithdrawals gets a reference to the given bool and assigns it to the EnableWithdrawals field.
func (o *GetApiKeyPermissionResponse) SetEnableWithdrawals(v bool) {
	o.EnableWithdrawals = &v
}

// GetEnableInternalTransfer returns the EnableInternalTransfer field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetEnableInternalTransfer() bool {
	if o == nil || common.IsNil(o.EnableInternalTransfer) {
		var ret bool
		return ret
	}
	return *o.EnableInternalTransfer
}

// GetEnableInternalTransferOk returns a tuple with the EnableInternalTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetEnableInternalTransferOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnableInternalTransfer) {
		return nil, false
	}
	return o.EnableInternalTransfer, true
}

// HasEnableInternalTransfer returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasEnableInternalTransfer() bool {
	if o != nil && !common.IsNil(o.EnableInternalTransfer) {
		return true
	}

	return false
}

// SetEnableInternalTransfer gets a reference to the given bool and assigns it to the EnableInternalTransfer field.
func (o *GetApiKeyPermissionResponse) SetEnableInternalTransfer(v bool) {
	o.EnableInternalTransfer = &v
}

// GetEnableMargin returns the EnableMargin field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetEnableMargin() bool {
	if o == nil || common.IsNil(o.EnableMargin) {
		var ret bool
		return ret
	}
	return *o.EnableMargin
}

// GetEnableMarginOk returns a tuple with the EnableMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetEnableMarginOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnableMargin) {
		return nil, false
	}
	return o.EnableMargin, true
}

// HasEnableMargin returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasEnableMargin() bool {
	if o != nil && !common.IsNil(o.EnableMargin) {
		return true
	}

	return false
}

// SetEnableMargin gets a reference to the given bool and assigns it to the EnableMargin field.
func (o *GetApiKeyPermissionResponse) SetEnableMargin(v bool) {
	o.EnableMargin = &v
}

// GetEnableFutures returns the EnableFutures field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetEnableFutures() bool {
	if o == nil || common.IsNil(o.EnableFutures) {
		var ret bool
		return ret
	}
	return *o.EnableFutures
}

// GetEnableFuturesOk returns a tuple with the EnableFutures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetEnableFuturesOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnableFutures) {
		return nil, false
	}
	return o.EnableFutures, true
}

// HasEnableFutures returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasEnableFutures() bool {
	if o != nil && !common.IsNil(o.EnableFutures) {
		return true
	}

	return false
}

// SetEnableFutures gets a reference to the given bool and assigns it to the EnableFutures field.
func (o *GetApiKeyPermissionResponse) SetEnableFutures(v bool) {
	o.EnableFutures = &v
}

// GetPermitsUniversalTransfer returns the PermitsUniversalTransfer field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetPermitsUniversalTransfer() bool {
	if o == nil || common.IsNil(o.PermitsUniversalTransfer) {
		var ret bool
		return ret
	}
	return *o.PermitsUniversalTransfer
}

// GetPermitsUniversalTransferOk returns a tuple with the PermitsUniversalTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetPermitsUniversalTransferOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PermitsUniversalTransfer) {
		return nil, false
	}
	return o.PermitsUniversalTransfer, true
}

// HasPermitsUniversalTransfer returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasPermitsUniversalTransfer() bool {
	if o != nil && !common.IsNil(o.PermitsUniversalTransfer) {
		return true
	}

	return false
}

// SetPermitsUniversalTransfer gets a reference to the given bool and assigns it to the PermitsUniversalTransfer field.
func (o *GetApiKeyPermissionResponse) SetPermitsUniversalTransfer(v bool) {
	o.PermitsUniversalTransfer = &v
}

// GetEnableVanillaOptions returns the EnableVanillaOptions field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetEnableVanillaOptions() bool {
	if o == nil || common.IsNil(o.EnableVanillaOptions) {
		var ret bool
		return ret
	}
	return *o.EnableVanillaOptions
}

// GetEnableVanillaOptionsOk returns a tuple with the EnableVanillaOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetEnableVanillaOptionsOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnableVanillaOptions) {
		return nil, false
	}
	return o.EnableVanillaOptions, true
}

// HasEnableVanillaOptions returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasEnableVanillaOptions() bool {
	if o != nil && !common.IsNil(o.EnableVanillaOptions) {
		return true
	}

	return false
}

// SetEnableVanillaOptions gets a reference to the given bool and assigns it to the EnableVanillaOptions field.
func (o *GetApiKeyPermissionResponse) SetEnableVanillaOptions(v bool) {
	o.EnableVanillaOptions = &v
}

// GetEnableFixApiTrade returns the EnableFixApiTrade field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetEnableFixApiTrade() bool {
	if o == nil || common.IsNil(o.EnableFixApiTrade) {
		var ret bool
		return ret
	}
	return *o.EnableFixApiTrade
}

// GetEnableFixApiTradeOk returns a tuple with the EnableFixApiTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetEnableFixApiTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnableFixApiTrade) {
		return nil, false
	}
	return o.EnableFixApiTrade, true
}

// HasEnableFixApiTrade returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasEnableFixApiTrade() bool {
	if o != nil && !common.IsNil(o.EnableFixApiTrade) {
		return true
	}

	return false
}

// SetEnableFixApiTrade gets a reference to the given bool and assigns it to the EnableFixApiTrade field.
func (o *GetApiKeyPermissionResponse) SetEnableFixApiTrade(v bool) {
	o.EnableFixApiTrade = &v
}

// GetEnableFixReadOnly returns the EnableFixReadOnly field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetEnableFixReadOnly() bool {
	if o == nil || common.IsNil(o.EnableFixReadOnly) {
		var ret bool
		return ret
	}
	return *o.EnableFixReadOnly
}

// GetEnableFixReadOnlyOk returns a tuple with the EnableFixReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetEnableFixReadOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnableFixReadOnly) {
		return nil, false
	}
	return o.EnableFixReadOnly, true
}

// HasEnableFixReadOnly returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasEnableFixReadOnly() bool {
	if o != nil && !common.IsNil(o.EnableFixReadOnly) {
		return true
	}

	return false
}

// SetEnableFixReadOnly gets a reference to the given bool and assigns it to the EnableFixReadOnly field.
func (o *GetApiKeyPermissionResponse) SetEnableFixReadOnly(v bool) {
	o.EnableFixReadOnly = &v
}

// GetEnableSpotAndMarginTrading returns the EnableSpotAndMarginTrading field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetEnableSpotAndMarginTrading() bool {
	if o == nil || common.IsNil(o.EnableSpotAndMarginTrading) {
		var ret bool
		return ret
	}
	return *o.EnableSpotAndMarginTrading
}

// GetEnableSpotAndMarginTradingOk returns a tuple with the EnableSpotAndMarginTrading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetEnableSpotAndMarginTradingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnableSpotAndMarginTrading) {
		return nil, false
	}
	return o.EnableSpotAndMarginTrading, true
}

// HasEnableSpotAndMarginTrading returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasEnableSpotAndMarginTrading() bool {
	if o != nil && !common.IsNil(o.EnableSpotAndMarginTrading) {
		return true
	}

	return false
}

// SetEnableSpotAndMarginTrading gets a reference to the given bool and assigns it to the EnableSpotAndMarginTrading field.
func (o *GetApiKeyPermissionResponse) SetEnableSpotAndMarginTrading(v bool) {
	o.EnableSpotAndMarginTrading = &v
}

// GetEnablePortfolioMarginTrading returns the EnablePortfolioMarginTrading field value if set, zero value otherwise.
func (o *GetApiKeyPermissionResponse) GetEnablePortfolioMarginTrading() bool {
	if o == nil || common.IsNil(o.EnablePortfolioMarginTrading) {
		var ret bool
		return ret
	}
	return *o.EnablePortfolioMarginTrading
}

// GetEnablePortfolioMarginTradingOk returns a tuple with the EnablePortfolioMarginTrading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyPermissionResponse) GetEnablePortfolioMarginTradingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnablePortfolioMarginTrading) {
		return nil, false
	}
	return o.EnablePortfolioMarginTrading, true
}

// HasEnablePortfolioMarginTrading returns a boolean if a field has been set.
func (o *GetApiKeyPermissionResponse) HasEnablePortfolioMarginTrading() bool {
	if o != nil && !common.IsNil(o.EnablePortfolioMarginTrading) {
		return true
	}

	return false
}

// SetEnablePortfolioMarginTrading gets a reference to the given bool and assigns it to the EnablePortfolioMarginTrading field.
func (o *GetApiKeyPermissionResponse) SetEnablePortfolioMarginTrading(v bool) {
	o.EnablePortfolioMarginTrading = &v
}

func (o GetApiKeyPermissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApiKeyPermissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IpRestrict) {
		toSerialize["ipRestrict"] = o.IpRestrict
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.EnableReading) {
		toSerialize["enableReading"] = o.EnableReading
	}
	if !common.IsNil(o.EnableWithdrawals) {
		toSerialize["enableWithdrawals"] = o.EnableWithdrawals
	}
	if !common.IsNil(o.EnableInternalTransfer) {
		toSerialize["enableInternalTransfer"] = o.EnableInternalTransfer
	}
	if !common.IsNil(o.EnableMargin) {
		toSerialize["enableMargin"] = o.EnableMargin
	}
	if !common.IsNil(o.EnableFutures) {
		toSerialize["enableFutures"] = o.EnableFutures
	}
	if !common.IsNil(o.PermitsUniversalTransfer) {
		toSerialize["permitsUniversalTransfer"] = o.PermitsUniversalTransfer
	}
	if !common.IsNil(o.EnableVanillaOptions) {
		toSerialize["enableVanillaOptions"] = o.EnableVanillaOptions
	}
	if !common.IsNil(o.EnableFixApiTrade) {
		toSerialize["enableFixApiTrade"] = o.EnableFixApiTrade
	}
	if !common.IsNil(o.EnableFixReadOnly) {
		toSerialize["enableFixReadOnly"] = o.EnableFixReadOnly
	}
	if !common.IsNil(o.EnableSpotAndMarginTrading) {
		toSerialize["enableSpotAndMarginTrading"] = o.EnableSpotAndMarginTrading
	}
	if !common.IsNil(o.EnablePortfolioMarginTrading) {
		toSerialize["enablePortfolioMarginTrading"] = o.EnablePortfolioMarginTrading
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetApiKeyPermissionResponse) UnmarshalJSON(data []byte) (err error) {
	varGetApiKeyPermissionResponse := _GetApiKeyPermissionResponse{}

	err = json.Unmarshal(data, &varGetApiKeyPermissionResponse)

	if err != nil {
		return err
	}

	*o = GetApiKeyPermissionResponse(varGetApiKeyPermissionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ipRestrict")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "enableReading")
		delete(additionalProperties, "enableWithdrawals")
		delete(additionalProperties, "enableInternalTransfer")
		delete(additionalProperties, "enableMargin")
		delete(additionalProperties, "enableFutures")
		delete(additionalProperties, "permitsUniversalTransfer")
		delete(additionalProperties, "enableVanillaOptions")
		delete(additionalProperties, "enableFixApiTrade")
		delete(additionalProperties, "enableFixReadOnly")
		delete(additionalProperties, "enableSpotAndMarginTrading")
		delete(additionalProperties, "enablePortfolioMarginTrading")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetApiKeyPermissionResponse struct {
	value *GetApiKeyPermissionResponse
	isSet bool
}

func (v NullableGetApiKeyPermissionResponse) Get() *GetApiKeyPermissionResponse {
	return v.value
}

func (v *NullableGetApiKeyPermissionResponse) Set(val *GetApiKeyPermissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApiKeyPermissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApiKeyPermissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApiKeyPermissionResponse(val *GetApiKeyPermissionResponse) *NullableGetApiKeyPermissionResponse {
	return &NullableGetApiKeyPermissionResponse{value: val, isSet: true}
}

func (v NullableGetApiKeyPermissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApiKeyPermissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
