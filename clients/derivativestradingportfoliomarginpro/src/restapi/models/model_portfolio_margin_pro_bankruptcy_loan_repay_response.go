/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PortfolioMarginProBankruptcyLoanRepayResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PortfolioMarginProBankruptcyLoanRepayResponse{}

// PortfolioMarginProBankruptcyLoanRepayResponse struct for PortfolioMarginProBankruptcyLoanRepayResponse
type PortfolioMarginProBankruptcyLoanRepayResponse struct {
	TranId               *int64 `json:"tranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortfolioMarginProBankruptcyLoanRepayResponse PortfolioMarginProBankruptcyLoanRepayResponse

// NewPortfolioMarginProBankruptcyLoanRepayResponse instantiates a new PortfolioMarginProBankruptcyLoanRepayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMarginProBankruptcyLoanRepayResponse() *PortfolioMarginProBankruptcyLoanRepayResponse {
	this := PortfolioMarginProBankruptcyLoanRepayResponse{}
	return &this
}

// NewPortfolioMarginProBankruptcyLoanRepayResponseWithDefaults instantiates a new PortfolioMarginProBankruptcyLoanRepayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMarginProBankruptcyLoanRepayResponseWithDefaults() *PortfolioMarginProBankruptcyLoanRepayResponse {
	this := PortfolioMarginProBankruptcyLoanRepayResponse{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *PortfolioMarginProBankruptcyLoanRepayResponse) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginProBankruptcyLoanRepayResponse) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *PortfolioMarginProBankruptcyLoanRepayResponse) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *PortfolioMarginProBankruptcyLoanRepayResponse) SetTranId(v int64) {
	o.TranId = &v
}

func (o PortfolioMarginProBankruptcyLoanRepayResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioMarginProBankruptcyLoanRepayResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortfolioMarginProBankruptcyLoanRepayResponse) UnmarshalJSON(data []byte) (err error) {
	varPortfolioMarginProBankruptcyLoanRepayResponse := _PortfolioMarginProBankruptcyLoanRepayResponse{}

	err = json.Unmarshal(data, &varPortfolioMarginProBankruptcyLoanRepayResponse)

	if err != nil {
		return err
	}

	*o = PortfolioMarginProBankruptcyLoanRepayResponse(varPortfolioMarginProBankruptcyLoanRepayResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortfolioMarginProBankruptcyLoanRepayResponse struct {
	value *PortfolioMarginProBankruptcyLoanRepayResponse
	isSet bool
}

func (v NullablePortfolioMarginProBankruptcyLoanRepayResponse) Get() *PortfolioMarginProBankruptcyLoanRepayResponse {
	return v.value
}

func (v *NullablePortfolioMarginProBankruptcyLoanRepayResponse) Set(val *PortfolioMarginProBankruptcyLoanRepayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMarginProBankruptcyLoanRepayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMarginProBankruptcyLoanRepayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMarginProBankruptcyLoanRepayResponse(val *PortfolioMarginProBankruptcyLoanRepayResponse) *NullablePortfolioMarginProBankruptcyLoanRepayResponse {
	return &NullablePortfolioMarginProBankruptcyLoanRepayResponse{value: val, isSet: true}
}

func (v NullablePortfolioMarginProBankruptcyLoanRepayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioMarginProBankruptcyLoanRepayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
