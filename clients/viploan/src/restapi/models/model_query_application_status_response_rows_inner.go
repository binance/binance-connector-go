/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryApplicationStatusResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryApplicationStatusResponseRowsInner{}

// QueryApplicationStatusResponseRowsInner struct for QueryApplicationStatusResponseRowsInner
type QueryApplicationStatusResponseRowsInner struct {
	LoanAccountId        *string `json:"loanAccountId,omitempty"`
	OrderId              *string `json:"orderId,omitempty"`
	RequestId            *string `json:"requestId,omitempty"`
	LoanCoin             *string `json:"loanCoin,omitempty"`
	LoanAmount           *string `json:"loanAmount,omitempty"`
	CollateralAccountId  *string `json:"collateralAccountId,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	LoanTerm             *string `json:"loanTerm,omitempty"`
	Status               *string `json:"status,omitempty"`
	LoanDate             *string `json:"loanDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryApplicationStatusResponseRowsInner QueryApplicationStatusResponseRowsInner

// NewQueryApplicationStatusResponseRowsInner instantiates a new QueryApplicationStatusResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryApplicationStatusResponseRowsInner() *QueryApplicationStatusResponseRowsInner {
	this := QueryApplicationStatusResponseRowsInner{}
	return &this
}

// NewQueryApplicationStatusResponseRowsInnerWithDefaults instantiates a new QueryApplicationStatusResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryApplicationStatusResponseRowsInnerWithDefaults() *QueryApplicationStatusResponseRowsInner {
	this := QueryApplicationStatusResponseRowsInner{}
	return &this
}

// GetLoanAccountId returns the LoanAccountId field value if set, zero value otherwise.
func (o *QueryApplicationStatusResponseRowsInner) GetLoanAccountId() string {
	if o == nil || common.IsNil(o.LoanAccountId) {
		var ret string
		return ret
	}
	return *o.LoanAccountId
}

// GetLoanAccountIdOk returns a tuple with the LoanAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApplicationStatusResponseRowsInner) GetLoanAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanAccountId) {
		return nil, false
	}
	return o.LoanAccountId, true
}

// HasLoanAccountId returns a boolean if a field has been set.
func (o *QueryApplicationStatusResponseRowsInner) HasLoanAccountId() bool {
	if o != nil && !common.IsNil(o.LoanAccountId) {
		return true
	}

	return false
}

// SetLoanAccountId gets a reference to the given string and assigns it to the LoanAccountId field.
func (o *QueryApplicationStatusResponseRowsInner) SetLoanAccountId(v string) {
	o.LoanAccountId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QueryApplicationStatusResponseRowsInner) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApplicationStatusResponseRowsInner) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryApplicationStatusResponseRowsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *QueryApplicationStatusResponseRowsInner) SetOrderId(v string) {
	o.OrderId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *QueryApplicationStatusResponseRowsInner) GetRequestId() string {
	if o == nil || common.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApplicationStatusResponseRowsInner) GetRequestIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *QueryApplicationStatusResponseRowsInner) HasRequestId() bool {
	if o != nil && !common.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *QueryApplicationStatusResponseRowsInner) SetRequestId(v string) {
	o.RequestId = &v
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *QueryApplicationStatusResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApplicationStatusResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *QueryApplicationStatusResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *QueryApplicationStatusResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetLoanAmount returns the LoanAmount field value if set, zero value otherwise.
func (o *QueryApplicationStatusResponseRowsInner) GetLoanAmount() string {
	if o == nil || common.IsNil(o.LoanAmount) {
		var ret string
		return ret
	}
	return *o.LoanAmount
}

// GetLoanAmountOk returns a tuple with the LoanAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApplicationStatusResponseRowsInner) GetLoanAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanAmount) {
		return nil, false
	}
	return o.LoanAmount, true
}

// HasLoanAmount returns a boolean if a field has been set.
func (o *QueryApplicationStatusResponseRowsInner) HasLoanAmount() bool {
	if o != nil && !common.IsNil(o.LoanAmount) {
		return true
	}

	return false
}

// SetLoanAmount gets a reference to the given string and assigns it to the LoanAmount field.
func (o *QueryApplicationStatusResponseRowsInner) SetLoanAmount(v string) {
	o.LoanAmount = &v
}

// GetCollateralAccountId returns the CollateralAccountId field value if set, zero value otherwise.
func (o *QueryApplicationStatusResponseRowsInner) GetCollateralAccountId() string {
	if o == nil || common.IsNil(o.CollateralAccountId) {
		var ret string
		return ret
	}
	return *o.CollateralAccountId
}

// GetCollateralAccountIdOk returns a tuple with the CollateralAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApplicationStatusResponseRowsInner) GetCollateralAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralAccountId) {
		return nil, false
	}
	return o.CollateralAccountId, true
}

// HasCollateralAccountId returns a boolean if a field has been set.
func (o *QueryApplicationStatusResponseRowsInner) HasCollateralAccountId() bool {
	if o != nil && !common.IsNil(o.CollateralAccountId) {
		return true
	}

	return false
}

// SetCollateralAccountId gets a reference to the given string and assigns it to the CollateralAccountId field.
func (o *QueryApplicationStatusResponseRowsInner) SetCollateralAccountId(v string) {
	o.CollateralAccountId = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *QueryApplicationStatusResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApplicationStatusResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *QueryApplicationStatusResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *QueryApplicationStatusResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetLoanTerm returns the LoanTerm field value if set, zero value otherwise.
func (o *QueryApplicationStatusResponseRowsInner) GetLoanTerm() string {
	if o == nil || common.IsNil(o.LoanTerm) {
		var ret string
		return ret
	}
	return *o.LoanTerm
}

// GetLoanTermOk returns a tuple with the LoanTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApplicationStatusResponseRowsInner) GetLoanTermOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanTerm) {
		return nil, false
	}
	return o.LoanTerm, true
}

// HasLoanTerm returns a boolean if a field has been set.
func (o *QueryApplicationStatusResponseRowsInner) HasLoanTerm() bool {
	if o != nil && !common.IsNil(o.LoanTerm) {
		return true
	}

	return false
}

// SetLoanTerm gets a reference to the given string and assigns it to the LoanTerm field.
func (o *QueryApplicationStatusResponseRowsInner) SetLoanTerm(v string) {
	o.LoanTerm = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryApplicationStatusResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApplicationStatusResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryApplicationStatusResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryApplicationStatusResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetLoanDate returns the LoanDate field value if set, zero value otherwise.
func (o *QueryApplicationStatusResponseRowsInner) GetLoanDate() string {
	if o == nil || common.IsNil(o.LoanDate) {
		var ret string
		return ret
	}
	return *o.LoanDate
}

// GetLoanDateOk returns a tuple with the LoanDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApplicationStatusResponseRowsInner) GetLoanDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanDate) {
		return nil, false
	}
	return o.LoanDate, true
}

// HasLoanDate returns a boolean if a field has been set.
func (o *QueryApplicationStatusResponseRowsInner) HasLoanDate() bool {
	if o != nil && !common.IsNil(o.LoanDate) {
		return true
	}

	return false
}

// SetLoanDate gets a reference to the given string and assigns it to the LoanDate field.
func (o *QueryApplicationStatusResponseRowsInner) SetLoanDate(v string) {
	o.LoanDate = &v
}

func (o QueryApplicationStatusResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryApplicationStatusResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanAccountId) {
		toSerialize["loanAccountId"] = o.LoanAccountId
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.LoanAmount) {
		toSerialize["loanAmount"] = o.LoanAmount
	}
	if !common.IsNil(o.CollateralAccountId) {
		toSerialize["collateralAccountId"] = o.CollateralAccountId
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.LoanTerm) {
		toSerialize["loanTerm"] = o.LoanTerm
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.LoanDate) {
		toSerialize["loanDate"] = o.LoanDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryApplicationStatusResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryApplicationStatusResponseRowsInner := _QueryApplicationStatusResponseRowsInner{}

	err = json.Unmarshal(data, &varQueryApplicationStatusResponseRowsInner)

	if err != nil {
		return err
	}

	*o = QueryApplicationStatusResponseRowsInner(varQueryApplicationStatusResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanAccountId")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "loanAmount")
		delete(additionalProperties, "collateralAccountId")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "loanTerm")
		delete(additionalProperties, "status")
		delete(additionalProperties, "loanDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryApplicationStatusResponseRowsInner struct {
	value *QueryApplicationStatusResponseRowsInner
	isSet bool
}

func (v NullableQueryApplicationStatusResponseRowsInner) Get() *QueryApplicationStatusResponseRowsInner {
	return v.value
}

func (v *NullableQueryApplicationStatusResponseRowsInner) Set(val *QueryApplicationStatusResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryApplicationStatusResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryApplicationStatusResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryApplicationStatusResponseRowsInner(val *QueryApplicationStatusResponseRowsInner) *NullableQueryApplicationStatusResponseRowsInner {
	return &NullableQueryApplicationStatusResponseRowsInner{value: val, isSet: true}
}

func (v NullableQueryApplicationStatusResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryApplicationStatusResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
