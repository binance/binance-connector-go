/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RedeemBfusdResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RedeemBfusdResponse{}

// RedeemBfusdResponse struct for RedeemBfusdResponse
type RedeemBfusdResponse struct {
	Success              *bool   `json:"success,omitempty"`
	ReceiveAmount        *string `json:"receiveAmount,omitempty"`
	Fee                  *string `json:"fee,omitempty"`
	ArrivalTime          *int64  `json:"arrivalTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RedeemBfusdResponse RedeemBfusdResponse

// NewRedeemBfusdResponse instantiates a new RedeemBfusdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedeemBfusdResponse() *RedeemBfusdResponse {
	this := RedeemBfusdResponse{}
	return &this
}

// NewRedeemBfusdResponseWithDefaults instantiates a new RedeemBfusdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedeemBfusdResponseWithDefaults() *RedeemBfusdResponse {
	this := RedeemBfusdResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *RedeemBfusdResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemBfusdResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *RedeemBfusdResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *RedeemBfusdResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetReceiveAmount returns the ReceiveAmount field value if set, zero value otherwise.
func (o *RedeemBfusdResponse) GetReceiveAmount() string {
	if o == nil || common.IsNil(o.ReceiveAmount) {
		var ret string
		return ret
	}
	return *o.ReceiveAmount
}

// GetReceiveAmountOk returns a tuple with the ReceiveAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemBfusdResponse) GetReceiveAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReceiveAmount) {
		return nil, false
	}
	return o.ReceiveAmount, true
}

// HasReceiveAmount returns a boolean if a field has been set.
func (o *RedeemBfusdResponse) HasReceiveAmount() bool {
	if o != nil && !common.IsNil(o.ReceiveAmount) {
		return true
	}

	return false
}

// SetReceiveAmount gets a reference to the given string and assigns it to the ReceiveAmount field.
func (o *RedeemBfusdResponse) SetReceiveAmount(v string) {
	o.ReceiveAmount = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *RedeemBfusdResponse) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemBfusdResponse) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *RedeemBfusdResponse) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *RedeemBfusdResponse) SetFee(v string) {
	o.Fee = &v
}

// GetArrivalTime returns the ArrivalTime field value if set, zero value otherwise.
func (o *RedeemBfusdResponse) GetArrivalTime() int64 {
	if o == nil || common.IsNil(o.ArrivalTime) {
		var ret int64
		return ret
	}
	return *o.ArrivalTime
}

// GetArrivalTimeOk returns a tuple with the ArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemBfusdResponse) GetArrivalTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ArrivalTime) {
		return nil, false
	}
	return o.ArrivalTime, true
}

// HasArrivalTime returns a boolean if a field has been set.
func (o *RedeemBfusdResponse) HasArrivalTime() bool {
	if o != nil && !common.IsNil(o.ArrivalTime) {
		return true
	}

	return false
}

// SetArrivalTime gets a reference to the given int64 and assigns it to the ArrivalTime field.
func (o *RedeemBfusdResponse) SetArrivalTime(v int64) {
	o.ArrivalTime = &v
}

func (o RedeemBfusdResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedeemBfusdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.ReceiveAmount) {
		toSerialize["receiveAmount"] = o.ReceiveAmount
	}
	if !common.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !common.IsNil(o.ArrivalTime) {
		toSerialize["arrivalTime"] = o.ArrivalTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RedeemBfusdResponse) UnmarshalJSON(data []byte) (err error) {
	varRedeemBfusdResponse := _RedeemBfusdResponse{}

	err = json.Unmarshal(data, &varRedeemBfusdResponse)

	if err != nil {
		return err
	}

	*o = RedeemBfusdResponse(varRedeemBfusdResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "receiveAmount")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "arrivalTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRedeemBfusdResponse struct {
	value *RedeemBfusdResponse
	isSet bool
}

func (v NullableRedeemBfusdResponse) Get() *RedeemBfusdResponse {
	return v.value
}

func (v *NullableRedeemBfusdResponse) Set(val *RedeemBfusdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRedeemBfusdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRedeemBfusdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedeemBfusdResponse(val *RedeemBfusdResponse) *NullableRedeemBfusdResponse {
	return &NullableRedeemBfusdResponse{value: val, isSet: true}
}

func (v NullableRedeemBfusdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedeemBfusdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
