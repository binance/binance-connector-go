/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the StatisticListResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StatisticListResponseData{}

// StatisticListResponseData struct for StatisticListResponseData
type StatisticListResponseData struct {
	FifteenMinHashRate   *string                               `json:"fifteenMinHashRate,omitempty"`
	DayHashRate          *string                               `json:"dayHashRate,omitempty"`
	ValidNum             *int64                                `json:"validNum,omitempty"`
	InvalidNum           *int64                                `json:"invalidNum,omitempty"`
	ProfitToday          *StatisticListResponseDataProfitToday `json:"profitToday,omitempty"`
	ProfitYesterday      *StatisticListResponseDataProfitToday `json:"profitYesterday,omitempty"`
	UserName             *string                               `json:"userName,omitempty"`
	Unit                 *string                               `json:"unit,omitempty"`
	Algo                 *string                               `json:"algo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StatisticListResponseData StatisticListResponseData

// NewStatisticListResponseData instantiates a new StatisticListResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticListResponseData() *StatisticListResponseData {
	this := StatisticListResponseData{}
	return &this
}

// NewStatisticListResponseDataWithDefaults instantiates a new StatisticListResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticListResponseDataWithDefaults() *StatisticListResponseData {
	this := StatisticListResponseData{}
	return &this
}

// GetFifteenMinHashRate returns the FifteenMinHashRate field value if set, zero value otherwise.
func (o *StatisticListResponseData) GetFifteenMinHashRate() string {
	if o == nil || common.IsNil(o.FifteenMinHashRate) {
		var ret string
		return ret
	}
	return *o.FifteenMinHashRate
}

// GetFifteenMinHashRateOk returns a tuple with the FifteenMinHashRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponseData) GetFifteenMinHashRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.FifteenMinHashRate) {
		return nil, false
	}
	return o.FifteenMinHashRate, true
}

// HasFifteenMinHashRate returns a boolean if a field has been set.
func (o *StatisticListResponseData) HasFifteenMinHashRate() bool {
	if o != nil && !common.IsNil(o.FifteenMinHashRate) {
		return true
	}

	return false
}

// SetFifteenMinHashRate gets a reference to the given string and assigns it to the FifteenMinHashRate field.
func (o *StatisticListResponseData) SetFifteenMinHashRate(v string) {
	o.FifteenMinHashRate = &v
}

// GetDayHashRate returns the DayHashRate field value if set, zero value otherwise.
func (o *StatisticListResponseData) GetDayHashRate() string {
	if o == nil || common.IsNil(o.DayHashRate) {
		var ret string
		return ret
	}
	return *o.DayHashRate
}

// GetDayHashRateOk returns a tuple with the DayHashRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponseData) GetDayHashRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.DayHashRate) {
		return nil, false
	}
	return o.DayHashRate, true
}

// HasDayHashRate returns a boolean if a field has been set.
func (o *StatisticListResponseData) HasDayHashRate() bool {
	if o != nil && !common.IsNil(o.DayHashRate) {
		return true
	}

	return false
}

// SetDayHashRate gets a reference to the given string and assigns it to the DayHashRate field.
func (o *StatisticListResponseData) SetDayHashRate(v string) {
	o.DayHashRate = &v
}

// GetValidNum returns the ValidNum field value if set, zero value otherwise.
func (o *StatisticListResponseData) GetValidNum() int64 {
	if o == nil || common.IsNil(o.ValidNum) {
		var ret int64
		return ret
	}
	return *o.ValidNum
}

// GetValidNumOk returns a tuple with the ValidNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponseData) GetValidNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ValidNum) {
		return nil, false
	}
	return o.ValidNum, true
}

// HasValidNum returns a boolean if a field has been set.
func (o *StatisticListResponseData) HasValidNum() bool {
	if o != nil && !common.IsNil(o.ValidNum) {
		return true
	}

	return false
}

// SetValidNum gets a reference to the given int64 and assigns it to the ValidNum field.
func (o *StatisticListResponseData) SetValidNum(v int64) {
	o.ValidNum = &v
}

// GetInvalidNum returns the InvalidNum field value if set, zero value otherwise.
func (o *StatisticListResponseData) GetInvalidNum() int64 {
	if o == nil || common.IsNil(o.InvalidNum) {
		var ret int64
		return ret
	}
	return *o.InvalidNum
}

// GetInvalidNumOk returns a tuple with the InvalidNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponseData) GetInvalidNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InvalidNum) {
		return nil, false
	}
	return o.InvalidNum, true
}

// HasInvalidNum returns a boolean if a field has been set.
func (o *StatisticListResponseData) HasInvalidNum() bool {
	if o != nil && !common.IsNil(o.InvalidNum) {
		return true
	}

	return false
}

// SetInvalidNum gets a reference to the given int64 and assigns it to the InvalidNum field.
func (o *StatisticListResponseData) SetInvalidNum(v int64) {
	o.InvalidNum = &v
}

// GetProfitToday returns the ProfitToday field value if set, zero value otherwise.
func (o *StatisticListResponseData) GetProfitToday() StatisticListResponseDataProfitToday {
	if o == nil || common.IsNil(o.ProfitToday) {
		var ret StatisticListResponseDataProfitToday
		return ret
	}
	return *o.ProfitToday
}

// GetProfitTodayOk returns a tuple with the ProfitToday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponseData) GetProfitTodayOk() (*StatisticListResponseDataProfitToday, bool) {
	if o == nil || common.IsNil(o.ProfitToday) {
		return nil, false
	}
	return o.ProfitToday, true
}

// HasProfitToday returns a boolean if a field has been set.
func (o *StatisticListResponseData) HasProfitToday() bool {
	if o != nil && !common.IsNil(o.ProfitToday) {
		return true
	}

	return false
}

// SetProfitToday gets a reference to the given StatisticListResponseDataProfitToday and assigns it to the ProfitToday field.
func (o *StatisticListResponseData) SetProfitToday(v StatisticListResponseDataProfitToday) {
	o.ProfitToday = &v
}

// GetProfitYesterday returns the ProfitYesterday field value if set, zero value otherwise.
func (o *StatisticListResponseData) GetProfitYesterday() StatisticListResponseDataProfitToday {
	if o == nil || common.IsNil(o.ProfitYesterday) {
		var ret StatisticListResponseDataProfitToday
		return ret
	}
	return *o.ProfitYesterday
}

// GetProfitYesterdayOk returns a tuple with the ProfitYesterday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponseData) GetProfitYesterdayOk() (*StatisticListResponseDataProfitToday, bool) {
	if o == nil || common.IsNil(o.ProfitYesterday) {
		return nil, false
	}
	return o.ProfitYesterday, true
}

// HasProfitYesterday returns a boolean if a field has been set.
func (o *StatisticListResponseData) HasProfitYesterday() bool {
	if o != nil && !common.IsNil(o.ProfitYesterday) {
		return true
	}

	return false
}

// SetProfitYesterday gets a reference to the given StatisticListResponseDataProfitToday and assigns it to the ProfitYesterday field.
func (o *StatisticListResponseData) SetProfitYesterday(v StatisticListResponseDataProfitToday) {
	o.ProfitYesterday = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *StatisticListResponseData) GetUserName() string {
	if o == nil || common.IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponseData) GetUserNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *StatisticListResponseData) HasUserName() bool {
	if o != nil && !common.IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *StatisticListResponseData) SetUserName(v string) {
	o.UserName = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *StatisticListResponseData) GetUnit() string {
	if o == nil || common.IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponseData) GetUnitOk() (*string, bool) {
	if o == nil || common.IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *StatisticListResponseData) HasUnit() bool {
	if o != nil && !common.IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *StatisticListResponseData) SetUnit(v string) {
	o.Unit = &v
}

// GetAlgo returns the Algo field value if set, zero value otherwise.
func (o *StatisticListResponseData) GetAlgo() string {
	if o == nil || common.IsNil(o.Algo) {
		var ret string
		return ret
	}
	return *o.Algo
}

// GetAlgoOk returns a tuple with the Algo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponseData) GetAlgoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Algo) {
		return nil, false
	}
	return o.Algo, true
}

// HasAlgo returns a boolean if a field has been set.
func (o *StatisticListResponseData) HasAlgo() bool {
	if o != nil && !common.IsNil(o.Algo) {
		return true
	}

	return false
}

// SetAlgo gets a reference to the given string and assigns it to the Algo field.
func (o *StatisticListResponseData) SetAlgo(v string) {
	o.Algo = &v
}

func (o StatisticListResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatisticListResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FifteenMinHashRate) {
		toSerialize["fifteenMinHashRate"] = o.FifteenMinHashRate
	}
	if !common.IsNil(o.DayHashRate) {
		toSerialize["dayHashRate"] = o.DayHashRate
	}
	if !common.IsNil(o.ValidNum) {
		toSerialize["validNum"] = o.ValidNum
	}
	if !common.IsNil(o.InvalidNum) {
		toSerialize["invalidNum"] = o.InvalidNum
	}
	if !common.IsNil(o.ProfitToday) {
		toSerialize["profitToday"] = o.ProfitToday
	}
	if !common.IsNil(o.ProfitYesterday) {
		toSerialize["profitYesterday"] = o.ProfitYesterday
	}
	if !common.IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !common.IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !common.IsNil(o.Algo) {
		toSerialize["algo"] = o.Algo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StatisticListResponseData) UnmarshalJSON(data []byte) (err error) {
	varStatisticListResponseData := _StatisticListResponseData{}

	err = json.Unmarshal(data, &varStatisticListResponseData)

	if err != nil {
		return err
	}

	*o = StatisticListResponseData(varStatisticListResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fifteenMinHashRate")
		delete(additionalProperties, "dayHashRate")
		delete(additionalProperties, "validNum")
		delete(additionalProperties, "invalidNum")
		delete(additionalProperties, "profitToday")
		delete(additionalProperties, "profitYesterday")
		delete(additionalProperties, "userName")
		delete(additionalProperties, "unit")
		delete(additionalProperties, "algo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStatisticListResponseData struct {
	value *StatisticListResponseData
	isSet bool
}

func (v NullableStatisticListResponseData) Get() *StatisticListResponseData {
	return v.value
}

func (v *NullableStatisticListResponseData) Set(val *StatisticListResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticListResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticListResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticListResponseData(val *StatisticListResponseData) *NullableStatisticListResponseData {
	return &NullableStatisticListResponseData{value: val, isSet: true}
}

func (v NullableStatisticListResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticListResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
