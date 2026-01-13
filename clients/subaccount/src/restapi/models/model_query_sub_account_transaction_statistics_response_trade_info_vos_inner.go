/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner{}

// QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner struct for QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner
type QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner struct {
	UserId               *int64 `json:"userId,omitempty"`
	Btc                  *int64 `json:"btc,omitempty"`
	BtcFutures           *int64 `json:"btcFutures,omitempty"`
	BtcMargin            *int64 `json:"btcMargin,omitempty"`
	Busd                 *int64 `json:"busd,omitempty"`
	BusdFutures          *int64 `json:"busdFutures,omitempty"`
	BusdMargin           *int64 `json:"busdMargin,omitempty"`
	Date                 *int64 `json:"date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner

// NewQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner instantiates a new QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner() *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner {
	this := QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner{}
	return &this
}

// NewQuerySubAccountTransactionStatisticsResponseTradeInfoVosInnerWithDefaults instantiates a new QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountTransactionStatisticsResponseTradeInfoVosInnerWithDefaults() *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner {
	this := QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetUserId() int64 {
	if o == nil || common.IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetUserIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) HasUserId() bool {
	if o != nil && !common.IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) SetUserId(v int64) {
	o.UserId = &v
}

// GetBtc returns the Btc field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetBtc() int64 {
	if o == nil || common.IsNil(o.Btc) {
		var ret int64
		return ret
	}
	return *o.Btc
}

// GetBtcOk returns a tuple with the Btc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetBtcOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Btc) {
		return nil, false
	}
	return o.Btc, true
}

// HasBtc returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) HasBtc() bool {
	if o != nil && !common.IsNil(o.Btc) {
		return true
	}

	return false
}

// SetBtc gets a reference to the given int64 and assigns it to the Btc field.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) SetBtc(v int64) {
	o.Btc = &v
}

// GetBtcFutures returns the BtcFutures field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetBtcFutures() int64 {
	if o == nil || common.IsNil(o.BtcFutures) {
		var ret int64
		return ret
	}
	return *o.BtcFutures
}

// GetBtcFuturesOk returns a tuple with the BtcFutures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetBtcFuturesOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BtcFutures) {
		return nil, false
	}
	return o.BtcFutures, true
}

// HasBtcFutures returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) HasBtcFutures() bool {
	if o != nil && !common.IsNil(o.BtcFutures) {
		return true
	}

	return false
}

// SetBtcFutures gets a reference to the given int64 and assigns it to the BtcFutures field.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) SetBtcFutures(v int64) {
	o.BtcFutures = &v
}

// GetBtcMargin returns the BtcMargin field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetBtcMargin() int64 {
	if o == nil || common.IsNil(o.BtcMargin) {
		var ret int64
		return ret
	}
	return *o.BtcMargin
}

// GetBtcMarginOk returns a tuple with the BtcMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetBtcMarginOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BtcMargin) {
		return nil, false
	}
	return o.BtcMargin, true
}

// HasBtcMargin returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) HasBtcMargin() bool {
	if o != nil && !common.IsNil(o.BtcMargin) {
		return true
	}

	return false
}

// SetBtcMargin gets a reference to the given int64 and assigns it to the BtcMargin field.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) SetBtcMargin(v int64) {
	o.BtcMargin = &v
}

// GetBusd returns the Busd field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetBusd() int64 {
	if o == nil || common.IsNil(o.Busd) {
		var ret int64
		return ret
	}
	return *o.Busd
}

// GetBusdOk returns a tuple with the Busd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetBusdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Busd) {
		return nil, false
	}
	return o.Busd, true
}

// HasBusd returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) HasBusd() bool {
	if o != nil && !common.IsNil(o.Busd) {
		return true
	}

	return false
}

// SetBusd gets a reference to the given int64 and assigns it to the Busd field.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) SetBusd(v int64) {
	o.Busd = &v
}

// GetBusdFutures returns the BusdFutures field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetBusdFutures() int64 {
	if o == nil || common.IsNil(o.BusdFutures) {
		var ret int64
		return ret
	}
	return *o.BusdFutures
}

// GetBusdFuturesOk returns a tuple with the BusdFutures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetBusdFuturesOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BusdFutures) {
		return nil, false
	}
	return o.BusdFutures, true
}

// HasBusdFutures returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) HasBusdFutures() bool {
	if o != nil && !common.IsNil(o.BusdFutures) {
		return true
	}

	return false
}

// SetBusdFutures gets a reference to the given int64 and assigns it to the BusdFutures field.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) SetBusdFutures(v int64) {
	o.BusdFutures = &v
}

// GetBusdMargin returns the BusdMargin field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetBusdMargin() int64 {
	if o == nil || common.IsNil(o.BusdMargin) {
		var ret int64
		return ret
	}
	return *o.BusdMargin
}

// GetBusdMarginOk returns a tuple with the BusdMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetBusdMarginOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BusdMargin) {
		return nil, false
	}
	return o.BusdMargin, true
}

// HasBusdMargin returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) HasBusdMargin() bool {
	if o != nil && !common.IsNil(o.BusdMargin) {
		return true
	}

	return false
}

// SetBusdMargin gets a reference to the given int64 and assigns it to the BusdMargin field.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) SetBusdMargin(v int64) {
	o.BusdMargin = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetDate() int64 {
	if o == nil || common.IsNil(o.Date) {
		var ret int64
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) GetDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) HasDate() bool {
	if o != nil && !common.IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given int64 and assigns it to the Date field.
func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) SetDate(v int64) {
	o.Date = &v
}

func (o QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !common.IsNil(o.Btc) {
		toSerialize["btc"] = o.Btc
	}
	if !common.IsNil(o.BtcFutures) {
		toSerialize["btcFutures"] = o.BtcFutures
	}
	if !common.IsNil(o.BtcMargin) {
		toSerialize["btcMargin"] = o.BtcMargin
	}
	if !common.IsNil(o.Busd) {
		toSerialize["busd"] = o.Busd
	}
	if !common.IsNil(o.BusdFutures) {
		toSerialize["busdFutures"] = o.BusdFutures
	}
	if !common.IsNil(o.BusdMargin) {
		toSerialize["busdMargin"] = o.BusdMargin
	}
	if !common.IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner := _QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner{}

	err = json.Unmarshal(data, &varQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner)

	if err != nil {
		return err
	}

	*o = QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner(varQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "userId")
		delete(additionalProperties, "btc")
		delete(additionalProperties, "btcFutures")
		delete(additionalProperties, "btcMargin")
		delete(additionalProperties, "busd")
		delete(additionalProperties, "busdFutures")
		delete(additionalProperties, "busdMargin")
		delete(additionalProperties, "date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner struct {
	value *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner
	isSet bool
}

func (v NullableQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) Get() *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner {
	return v.value
}

func (v *NullableQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) Set(val *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner(val *QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) *NullableQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner {
	return &NullableQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner{value: val, isSet: true}
}

func (v NullableQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
