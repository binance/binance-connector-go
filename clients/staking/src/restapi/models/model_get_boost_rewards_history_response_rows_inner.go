/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetBoostRewardsHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBoostRewardsHistoryResponseRowsInner{}

// GetBoostRewardsHistoryResponseRowsInner struct for GetBoostRewardsHistoryResponseRowsInner
type GetBoostRewardsHistoryResponseRowsInner struct {
	Time                 *int64  `json:"time,omitempty"`
	Token                *string `json:"token,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	BnsolHolding         *string `json:"bnsolHolding,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBoostRewardsHistoryResponseRowsInner GetBoostRewardsHistoryResponseRowsInner

// NewGetBoostRewardsHistoryResponseRowsInner instantiates a new GetBoostRewardsHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBoostRewardsHistoryResponseRowsInner() *GetBoostRewardsHistoryResponseRowsInner {
	this := GetBoostRewardsHistoryResponseRowsInner{}
	return &this
}

// NewGetBoostRewardsHistoryResponseRowsInnerWithDefaults instantiates a new GetBoostRewardsHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBoostRewardsHistoryResponseRowsInnerWithDefaults() *GetBoostRewardsHistoryResponseRowsInner {
	this := GetBoostRewardsHistoryResponseRowsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetBoostRewardsHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBoostRewardsHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetBoostRewardsHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetBoostRewardsHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetBoostRewardsHistoryResponseRowsInner) GetToken() string {
	if o == nil || common.IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBoostRewardsHistoryResponseRowsInner) GetTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetBoostRewardsHistoryResponseRowsInner) HasToken() bool {
	if o != nil && !common.IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetBoostRewardsHistoryResponseRowsInner) SetToken(v string) {
	o.Token = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetBoostRewardsHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBoostRewardsHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetBoostRewardsHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetBoostRewardsHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetBnsolHolding returns the BnsolHolding field value if set, zero value otherwise.
func (o *GetBoostRewardsHistoryResponseRowsInner) GetBnsolHolding() string {
	if o == nil || common.IsNil(o.BnsolHolding) {
		var ret string
		return ret
	}
	return *o.BnsolHolding
}

// GetBnsolHoldingOk returns a tuple with the BnsolHolding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBoostRewardsHistoryResponseRowsInner) GetBnsolHoldingOk() (*string, bool) {
	if o == nil || common.IsNil(o.BnsolHolding) {
		return nil, false
	}
	return o.BnsolHolding, true
}

// HasBnsolHolding returns a boolean if a field has been set.
func (o *GetBoostRewardsHistoryResponseRowsInner) HasBnsolHolding() bool {
	if o != nil && !common.IsNil(o.BnsolHolding) {
		return true
	}

	return false
}

// SetBnsolHolding gets a reference to the given string and assigns it to the BnsolHolding field.
func (o *GetBoostRewardsHistoryResponseRowsInner) SetBnsolHolding(v string) {
	o.BnsolHolding = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetBoostRewardsHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBoostRewardsHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetBoostRewardsHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetBoostRewardsHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetBoostRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBoostRewardsHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.BnsolHolding) {
		toSerialize["bnsolHolding"] = o.BnsolHolding
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBoostRewardsHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetBoostRewardsHistoryResponseRowsInner := _GetBoostRewardsHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetBoostRewardsHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetBoostRewardsHistoryResponseRowsInner(varGetBoostRewardsHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "token")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "bnsolHolding")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBoostRewardsHistoryResponseRowsInner struct {
	value *GetBoostRewardsHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetBoostRewardsHistoryResponseRowsInner) Get() *GetBoostRewardsHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetBoostRewardsHistoryResponseRowsInner) Set(val *GetBoostRewardsHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBoostRewardsHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBoostRewardsHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBoostRewardsHistoryResponseRowsInner(val *GetBoostRewardsHistoryResponseRowsInner) *NullableGetBoostRewardsHistoryResponseRowsInner {
	return &NullableGetBoostRewardsHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetBoostRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBoostRewardsHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
