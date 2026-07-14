/*
VIP Loan REST API

Access over-collateralized loan services, manage positions, and monitor collateral via the VIP Loan API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryVIPLoanFixedRateMarketResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryVIPLoanFixedRateMarketResponseRowsInner{}

// QueryVIPLoanFixedRateMarketResponseRowsInner struct for QueryVIPLoanFixedRateMarketResponseRowsInner
type QueryVIPLoanFixedRateMarketResponseRowsInner struct {
	// Supply request ID
	RequestId *int64 `json:"requestId,omitempty"`
	// Request number
	RequestNo *int64 `json:"requestNo,omitempty"`
	// Coin
	Coin *string `json:"coin,omitempty"`
	// Annual interest rate
	InterestRate *string `json:"interestRate,omitempty"`
	// Duration in days
	Duration *int64 `json:"duration,omitempty"`
	// Minimum borrow amount
	MinimumAmount *string `json:"minimumAmount,omitempty"`
	// Maximum available borrow amount
	AvailableAmount *string `json:"availableAmount,omitempty"`
	// Estimated interest
	EstimatedInterest    *string `json:"estimatedInterest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryVIPLoanFixedRateMarketResponseRowsInner QueryVIPLoanFixedRateMarketResponseRowsInner

// NewQueryVIPLoanFixedRateMarketResponseRowsInner instantiates a new QueryVIPLoanFixedRateMarketResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryVIPLoanFixedRateMarketResponseRowsInner() *QueryVIPLoanFixedRateMarketResponseRowsInner {
	this := QueryVIPLoanFixedRateMarketResponseRowsInner{}
	return &this
}

// NewQueryVIPLoanFixedRateMarketResponseRowsInnerWithDefaults instantiates a new QueryVIPLoanFixedRateMarketResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryVIPLoanFixedRateMarketResponseRowsInnerWithDefaults() *QueryVIPLoanFixedRateMarketResponseRowsInner {
	this := QueryVIPLoanFixedRateMarketResponseRowsInner{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetRequestId() int64 {
	if o == nil || common.IsNil(o.RequestId) {
		var ret int64
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetRequestIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) HasRequestId() bool {
	if o != nil && !common.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given int64 and assigns it to the RequestId field.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) SetRequestId(v int64) {
	o.RequestId = &v
}

// GetRequestNo returns the RequestNo field value if set, zero value otherwise.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetRequestNo() int64 {
	if o == nil || common.IsNil(o.RequestNo) {
		var ret int64
		return ret
	}
	return *o.RequestNo
}

// GetRequestNoOk returns a tuple with the RequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetRequestNoOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RequestNo) {
		return nil, false
	}
	return o.RequestNo, true
}

// HasRequestNo returns a boolean if a field has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) HasRequestNo() bool {
	if o != nil && !common.IsNil(o.RequestNo) {
		return true
	}

	return false
}

// SetRequestNo gets a reference to the given int64 and assigns it to the RequestNo field.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) SetRequestNo(v int64) {
	o.RequestNo = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) SetCoin(v string) {
	o.Coin = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetInterestRate() string {
	if o == nil || common.IsNil(o.InterestRate) {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) HasInterestRate() bool {
	if o != nil && !common.IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) SetInterestRate(v string) {
	o.InterestRate = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetDuration() int64 {
	if o == nil || common.IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetDurationOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) HasDuration() bool {
	if o != nil && !common.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) SetDuration(v int64) {
	o.Duration = &v
}

// GetMinimumAmount returns the MinimumAmount field value if set, zero value otherwise.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetMinimumAmount() string {
	if o == nil || common.IsNil(o.MinimumAmount) {
		var ret string
		return ret
	}
	return *o.MinimumAmount
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetMinimumAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinimumAmount) {
		return nil, false
	}
	return o.MinimumAmount, true
}

// HasMinimumAmount returns a boolean if a field has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) HasMinimumAmount() bool {
	if o != nil && !common.IsNil(o.MinimumAmount) {
		return true
	}

	return false
}

// SetMinimumAmount gets a reference to the given string and assigns it to the MinimumAmount field.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) SetMinimumAmount(v string) {
	o.MinimumAmount = &v
}

// GetAvailableAmount returns the AvailableAmount field value if set, zero value otherwise.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetAvailableAmount() string {
	if o == nil || common.IsNil(o.AvailableAmount) {
		var ret string
		return ret
	}
	return *o.AvailableAmount
}

// GetAvailableAmountOk returns a tuple with the AvailableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetAvailableAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvailableAmount) {
		return nil, false
	}
	return o.AvailableAmount, true
}

// HasAvailableAmount returns a boolean if a field has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) HasAvailableAmount() bool {
	if o != nil && !common.IsNil(o.AvailableAmount) {
		return true
	}

	return false
}

// SetAvailableAmount gets a reference to the given string and assigns it to the AvailableAmount field.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) SetAvailableAmount(v string) {
	o.AvailableAmount = &v
}

// GetEstimatedInterest returns the EstimatedInterest field value if set, zero value otherwise.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetEstimatedInterest() string {
	if o == nil || common.IsNil(o.EstimatedInterest) {
		var ret string
		return ret
	}
	return *o.EstimatedInterest
}

// GetEstimatedInterestOk returns a tuple with the EstimatedInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) GetEstimatedInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.EstimatedInterest) {
		return nil, false
	}
	return o.EstimatedInterest, true
}

// HasEstimatedInterest returns a boolean if a field has been set.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) HasEstimatedInterest() bool {
	if o != nil && !common.IsNil(o.EstimatedInterest) {
		return true
	}

	return false
}

// SetEstimatedInterest gets a reference to the given string and assigns it to the EstimatedInterest field.
func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) SetEstimatedInterest(v string) {
	o.EstimatedInterest = &v
}

func (o QueryVIPLoanFixedRateMarketResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryVIPLoanFixedRateMarketResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !common.IsNil(o.RequestNo) {
		toSerialize["requestNo"] = o.RequestNo
	}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.InterestRate) {
		toSerialize["interestRate"] = o.InterestRate
	}
	if !common.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !common.IsNil(o.MinimumAmount) {
		toSerialize["minimumAmount"] = o.MinimumAmount
	}
	if !common.IsNil(o.AvailableAmount) {
		toSerialize["availableAmount"] = o.AvailableAmount
	}
	if !common.IsNil(o.EstimatedInterest) {
		toSerialize["estimatedInterest"] = o.EstimatedInterest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryVIPLoanFixedRateMarketResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryVIPLoanFixedRateMarketResponseRowsInner := _QueryVIPLoanFixedRateMarketResponseRowsInner{}

	err = json.Unmarshal(data, &varQueryVIPLoanFixedRateMarketResponseRowsInner)

	if err != nil {
		return err
	}

	*o = QueryVIPLoanFixedRateMarketResponseRowsInner(varQueryVIPLoanFixedRateMarketResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "requestNo")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "interestRate")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "minimumAmount")
		delete(additionalProperties, "availableAmount")
		delete(additionalProperties, "estimatedInterest")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryVIPLoanFixedRateMarketResponseRowsInner struct {
	value *QueryVIPLoanFixedRateMarketResponseRowsInner
	isSet bool
}

func (v NullableQueryVIPLoanFixedRateMarketResponseRowsInner) Get() *QueryVIPLoanFixedRateMarketResponseRowsInner {
	return v.value
}

func (v *NullableQueryVIPLoanFixedRateMarketResponseRowsInner) Set(val *QueryVIPLoanFixedRateMarketResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryVIPLoanFixedRateMarketResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryVIPLoanFixedRateMarketResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryVIPLoanFixedRateMarketResponseRowsInner(val *QueryVIPLoanFixedRateMarketResponseRowsInner) *NullableQueryVIPLoanFixedRateMarketResponseRowsInner {
	return &NullableQueryVIPLoanFixedRateMarketResponseRowsInner{value: val, isSet: true}
}

func (v NullableQueryVIPLoanFixedRateMarketResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryVIPLoanFixedRateMarketResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
