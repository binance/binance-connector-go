/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySubAccountTransactionStatisticsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountTransactionStatisticsResponse{}

// QuerySubAccountTransactionStatisticsResponse struct for QuerySubAccountTransactionStatisticsResponse
type QuerySubAccountTransactionStatisticsResponse struct {
	Recent30BtcTotal         *string                                                         `json:"recent30BtcTotal,omitempty"`
	Recent30BtcFuturesTotal  *string                                                         `json:"recent30BtcFuturesTotal,omitempty"`
	Recent30BtcMarginTotal   *string                                                         `json:"recent30BtcMarginTotal,omitempty"`
	Recent30BusdTotal        *string                                                         `json:"recent30BusdTotal,omitempty"`
	Recent30BusdFuturesTotal *string                                                         `json:"recent30BusdFuturesTotal,omitempty"`
	Recent30BusdMarginTotal  *string                                                         `json:"recent30BusdMarginTotal,omitempty"`
	TradeInfoVos             []QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner `json:"tradeInfoVos,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _QuerySubAccountTransactionStatisticsResponse QuerySubAccountTransactionStatisticsResponse

// NewQuerySubAccountTransactionStatisticsResponse instantiates a new QuerySubAccountTransactionStatisticsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountTransactionStatisticsResponse() *QuerySubAccountTransactionStatisticsResponse {
	this := QuerySubAccountTransactionStatisticsResponse{}
	return &this
}

// NewQuerySubAccountTransactionStatisticsResponseWithDefaults instantiates a new QuerySubAccountTransactionStatisticsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountTransactionStatisticsResponseWithDefaults() *QuerySubAccountTransactionStatisticsResponse {
	this := QuerySubAccountTransactionStatisticsResponse{}
	return &this
}

// GetRecent30BtcTotal returns the Recent30BtcTotal field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponse) GetRecent30BtcTotal() string {
	if o == nil || common.IsNil(o.Recent30BtcTotal) {
		var ret string
		return ret
	}
	return *o.Recent30BtcTotal
}

// GetRecent30BtcTotalOk returns a tuple with the Recent30BtcTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) GetRecent30BtcTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Recent30BtcTotal) {
		return nil, false
	}
	return o.Recent30BtcTotal, true
}

// HasRecent30BtcTotal returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) HasRecent30BtcTotal() bool {
	if o != nil && !common.IsNil(o.Recent30BtcTotal) {
		return true
	}

	return false
}

// SetRecent30BtcTotal gets a reference to the given string and assigns it to the Recent30BtcTotal field.
func (o *QuerySubAccountTransactionStatisticsResponse) SetRecent30BtcTotal(v string) {
	o.Recent30BtcTotal = &v
}

// GetRecent30BtcFuturesTotal returns the Recent30BtcFuturesTotal field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponse) GetRecent30BtcFuturesTotal() string {
	if o == nil || common.IsNil(o.Recent30BtcFuturesTotal) {
		var ret string
		return ret
	}
	return *o.Recent30BtcFuturesTotal
}

// GetRecent30BtcFuturesTotalOk returns a tuple with the Recent30BtcFuturesTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) GetRecent30BtcFuturesTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Recent30BtcFuturesTotal) {
		return nil, false
	}
	return o.Recent30BtcFuturesTotal, true
}

// HasRecent30BtcFuturesTotal returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) HasRecent30BtcFuturesTotal() bool {
	if o != nil && !common.IsNil(o.Recent30BtcFuturesTotal) {
		return true
	}

	return false
}

// SetRecent30BtcFuturesTotal gets a reference to the given string and assigns it to the Recent30BtcFuturesTotal field.
func (o *QuerySubAccountTransactionStatisticsResponse) SetRecent30BtcFuturesTotal(v string) {
	o.Recent30BtcFuturesTotal = &v
}

// GetRecent30BtcMarginTotal returns the Recent30BtcMarginTotal field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponse) GetRecent30BtcMarginTotal() string {
	if o == nil || common.IsNil(o.Recent30BtcMarginTotal) {
		var ret string
		return ret
	}
	return *o.Recent30BtcMarginTotal
}

// GetRecent30BtcMarginTotalOk returns a tuple with the Recent30BtcMarginTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) GetRecent30BtcMarginTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Recent30BtcMarginTotal) {
		return nil, false
	}
	return o.Recent30BtcMarginTotal, true
}

// HasRecent30BtcMarginTotal returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) HasRecent30BtcMarginTotal() bool {
	if o != nil && !common.IsNil(o.Recent30BtcMarginTotal) {
		return true
	}

	return false
}

// SetRecent30BtcMarginTotal gets a reference to the given string and assigns it to the Recent30BtcMarginTotal field.
func (o *QuerySubAccountTransactionStatisticsResponse) SetRecent30BtcMarginTotal(v string) {
	o.Recent30BtcMarginTotal = &v
}

// GetRecent30BusdTotal returns the Recent30BusdTotal field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponse) GetRecent30BusdTotal() string {
	if o == nil || common.IsNil(o.Recent30BusdTotal) {
		var ret string
		return ret
	}
	return *o.Recent30BusdTotal
}

// GetRecent30BusdTotalOk returns a tuple with the Recent30BusdTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) GetRecent30BusdTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Recent30BusdTotal) {
		return nil, false
	}
	return o.Recent30BusdTotal, true
}

// HasRecent30BusdTotal returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) HasRecent30BusdTotal() bool {
	if o != nil && !common.IsNil(o.Recent30BusdTotal) {
		return true
	}

	return false
}

// SetRecent30BusdTotal gets a reference to the given string and assigns it to the Recent30BusdTotal field.
func (o *QuerySubAccountTransactionStatisticsResponse) SetRecent30BusdTotal(v string) {
	o.Recent30BusdTotal = &v
}

// GetRecent30BusdFuturesTotal returns the Recent30BusdFuturesTotal field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponse) GetRecent30BusdFuturesTotal() string {
	if o == nil || common.IsNil(o.Recent30BusdFuturesTotal) {
		var ret string
		return ret
	}
	return *o.Recent30BusdFuturesTotal
}

// GetRecent30BusdFuturesTotalOk returns a tuple with the Recent30BusdFuturesTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) GetRecent30BusdFuturesTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Recent30BusdFuturesTotal) {
		return nil, false
	}
	return o.Recent30BusdFuturesTotal, true
}

// HasRecent30BusdFuturesTotal returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) HasRecent30BusdFuturesTotal() bool {
	if o != nil && !common.IsNil(o.Recent30BusdFuturesTotal) {
		return true
	}

	return false
}

// SetRecent30BusdFuturesTotal gets a reference to the given string and assigns it to the Recent30BusdFuturesTotal field.
func (o *QuerySubAccountTransactionStatisticsResponse) SetRecent30BusdFuturesTotal(v string) {
	o.Recent30BusdFuturesTotal = &v
}

// GetRecent30BusdMarginTotal returns the Recent30BusdMarginTotal field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponse) GetRecent30BusdMarginTotal() string {
	if o == nil || common.IsNil(o.Recent30BusdMarginTotal) {
		var ret string
		return ret
	}
	return *o.Recent30BusdMarginTotal
}

// GetRecent30BusdMarginTotalOk returns a tuple with the Recent30BusdMarginTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) GetRecent30BusdMarginTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Recent30BusdMarginTotal) {
		return nil, false
	}
	return o.Recent30BusdMarginTotal, true
}

// HasRecent30BusdMarginTotal returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) HasRecent30BusdMarginTotal() bool {
	if o != nil && !common.IsNil(o.Recent30BusdMarginTotal) {
		return true
	}

	return false
}

// SetRecent30BusdMarginTotal gets a reference to the given string and assigns it to the Recent30BusdMarginTotal field.
func (o *QuerySubAccountTransactionStatisticsResponse) SetRecent30BusdMarginTotal(v string) {
	o.Recent30BusdMarginTotal = &v
}

// GetTradeInfoVos returns the TradeInfoVos field value if set, zero value otherwise.
func (o *QuerySubAccountTransactionStatisticsResponse) GetTradeInfoVos() []QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner {
	if o == nil || common.IsNil(o.TradeInfoVos) {
		var ret []QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner
		return ret
	}
	return o.TradeInfoVos
}

// GetTradeInfoVosOk returns a tuple with the TradeInfoVos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) GetTradeInfoVosOk() ([]QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner, bool) {
	if o == nil || common.IsNil(o.TradeInfoVos) {
		return nil, false
	}
	return o.TradeInfoVos, true
}

// HasTradeInfoVos returns a boolean if a field has been set.
func (o *QuerySubAccountTransactionStatisticsResponse) HasTradeInfoVos() bool {
	if o != nil && !common.IsNil(o.TradeInfoVos) {
		return true
	}

	return false
}

// SetTradeInfoVos gets a reference to the given []QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner and assigns it to the TradeInfoVos field.
func (o *QuerySubAccountTransactionStatisticsResponse) SetTradeInfoVos(v []QuerySubAccountTransactionStatisticsResponseTradeInfoVosInner) {
	o.TradeInfoVos = v
}

func (o QuerySubAccountTransactionStatisticsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountTransactionStatisticsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Recent30BtcTotal) {
		toSerialize["recent30BtcTotal"] = o.Recent30BtcTotal
	}
	if !common.IsNil(o.Recent30BtcFuturesTotal) {
		toSerialize["recent30BtcFuturesTotal"] = o.Recent30BtcFuturesTotal
	}
	if !common.IsNil(o.Recent30BtcMarginTotal) {
		toSerialize["recent30BtcMarginTotal"] = o.Recent30BtcMarginTotal
	}
	if !common.IsNil(o.Recent30BusdTotal) {
		toSerialize["recent30BusdTotal"] = o.Recent30BusdTotal
	}
	if !common.IsNil(o.Recent30BusdFuturesTotal) {
		toSerialize["recent30BusdFuturesTotal"] = o.Recent30BusdFuturesTotal
	}
	if !common.IsNil(o.Recent30BusdMarginTotal) {
		toSerialize["recent30BusdMarginTotal"] = o.Recent30BusdMarginTotal
	}
	if !common.IsNil(o.TradeInfoVos) {
		toSerialize["tradeInfoVos"] = o.TradeInfoVos
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubAccountTransactionStatisticsResponse) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountTransactionStatisticsResponse := _QuerySubAccountTransactionStatisticsResponse{}

	err = json.Unmarshal(data, &varQuerySubAccountTransactionStatisticsResponse)

	if err != nil {
		return err
	}

	*o = QuerySubAccountTransactionStatisticsResponse(varQuerySubAccountTransactionStatisticsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "recent30BtcTotal")
		delete(additionalProperties, "recent30BtcFuturesTotal")
		delete(additionalProperties, "recent30BtcMarginTotal")
		delete(additionalProperties, "recent30BusdTotal")
		delete(additionalProperties, "recent30BusdFuturesTotal")
		delete(additionalProperties, "recent30BusdMarginTotal")
		delete(additionalProperties, "tradeInfoVos")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountTransactionStatisticsResponse struct {
	value *QuerySubAccountTransactionStatisticsResponse
	isSet bool
}

func (v NullableQuerySubAccountTransactionStatisticsResponse) Get() *QuerySubAccountTransactionStatisticsResponse {
	return v.value
}

func (v *NullableQuerySubAccountTransactionStatisticsResponse) Set(val *QuerySubAccountTransactionStatisticsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountTransactionStatisticsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountTransactionStatisticsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountTransactionStatisticsResponse(val *QuerySubAccountTransactionStatisticsResponse) *NullableQuerySubAccountTransactionStatisticsResponse {
	return &NullableQuerySubAccountTransactionStatisticsResponse{value: val, isSet: true}
}

func (v NullableQuerySubAccountTransactionStatisticsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountTransactionStatisticsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
