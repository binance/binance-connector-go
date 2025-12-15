# StatisticListResponseData

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FifteenMinHashRate** | Pointer to **string** |  | [optional] 
**DayHashRate** | Pointer to **string** |  | [optional] 
**ValidNum** | Pointer to **int64** |  | [optional] 
**InvalidNum** | Pointer to **int64** |  | [optional] 
**ProfitToday** | Pointer to [**StatisticListResponseDataProfitToday**](StatisticListResponseDataProfitToday.md) |  | [optional] 
**ProfitYesterday** | Pointer to [**StatisticListResponseDataProfitToday**](StatisticListResponseDataProfitToday.md) |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Algo** | Pointer to **string** |  | [optional] 

## Methods

### NewStatisticListResponseData

`func NewStatisticListResponseData() *StatisticListResponseData`

NewStatisticListResponseData instantiates a new StatisticListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticListResponseDataWithDefaults

`func NewStatisticListResponseDataWithDefaults() *StatisticListResponseData`

NewStatisticListResponseDataWithDefaults instantiates a new StatisticListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFifteenMinHashRate

`func (o *StatisticListResponseData) GetFifteenMinHashRate() string`

GetFifteenMinHashRate returns the FifteenMinHashRate field if non-nil, zero value otherwise.

### GetFifteenMinHashRateOk

`func (o *StatisticListResponseData) GetFifteenMinHashRateOk() (*string, bool)`

GetFifteenMinHashRateOk returns a tuple with the FifteenMinHashRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifteenMinHashRate

`func (o *StatisticListResponseData) SetFifteenMinHashRate(v string)`

SetFifteenMinHashRate sets FifteenMinHashRate field to given value.

### HasFifteenMinHashRate

`func (o *StatisticListResponseData) HasFifteenMinHashRate() bool`

HasFifteenMinHashRate returns a boolean if a field has been set.

### GetDayHashRate

`func (o *StatisticListResponseData) GetDayHashRate() string`

GetDayHashRate returns the DayHashRate field if non-nil, zero value otherwise.

### GetDayHashRateOk

`func (o *StatisticListResponseData) GetDayHashRateOk() (*string, bool)`

GetDayHashRateOk returns a tuple with the DayHashRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayHashRate

`func (o *StatisticListResponseData) SetDayHashRate(v string)`

SetDayHashRate sets DayHashRate field to given value.

### HasDayHashRate

`func (o *StatisticListResponseData) HasDayHashRate() bool`

HasDayHashRate returns a boolean if a field has been set.

### GetValidNum

`func (o *StatisticListResponseData) GetValidNum() int64`

GetValidNum returns the ValidNum field if non-nil, zero value otherwise.

### GetValidNumOk

`func (o *StatisticListResponseData) GetValidNumOk() (*int64, bool)`

GetValidNumOk returns a tuple with the ValidNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNum

`func (o *StatisticListResponseData) SetValidNum(v int64)`

SetValidNum sets ValidNum field to given value.

### HasValidNum

`func (o *StatisticListResponseData) HasValidNum() bool`

HasValidNum returns a boolean if a field has been set.

### GetInvalidNum

`func (o *StatisticListResponseData) GetInvalidNum() int64`

GetInvalidNum returns the InvalidNum field if non-nil, zero value otherwise.

### GetInvalidNumOk

`func (o *StatisticListResponseData) GetInvalidNumOk() (*int64, bool)`

GetInvalidNumOk returns a tuple with the InvalidNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidNum

`func (o *StatisticListResponseData) SetInvalidNum(v int64)`

SetInvalidNum sets InvalidNum field to given value.

### HasInvalidNum

`func (o *StatisticListResponseData) HasInvalidNum() bool`

HasInvalidNum returns a boolean if a field has been set.

### GetProfitToday

`func (o *StatisticListResponseData) GetProfitToday() StatisticListResponseDataProfitToday`

GetProfitToday returns the ProfitToday field if non-nil, zero value otherwise.

### GetProfitTodayOk

`func (o *StatisticListResponseData) GetProfitTodayOk() (*StatisticListResponseDataProfitToday, bool)`

GetProfitTodayOk returns a tuple with the ProfitToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitToday

`func (o *StatisticListResponseData) SetProfitToday(v StatisticListResponseDataProfitToday)`

SetProfitToday sets ProfitToday field to given value.

### HasProfitToday

`func (o *StatisticListResponseData) HasProfitToday() bool`

HasProfitToday returns a boolean if a field has been set.

### GetProfitYesterday

`func (o *StatisticListResponseData) GetProfitYesterday() StatisticListResponseDataProfitToday`

GetProfitYesterday returns the ProfitYesterday field if non-nil, zero value otherwise.

### GetProfitYesterdayOk

`func (o *StatisticListResponseData) GetProfitYesterdayOk() (*StatisticListResponseDataProfitToday, bool)`

GetProfitYesterdayOk returns a tuple with the ProfitYesterday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitYesterday

`func (o *StatisticListResponseData) SetProfitYesterday(v StatisticListResponseDataProfitToday)`

SetProfitYesterday sets ProfitYesterday field to given value.

### HasProfitYesterday

`func (o *StatisticListResponseData) HasProfitYesterday() bool`

HasProfitYesterday returns a boolean if a field has been set.

### GetUserName

`func (o *StatisticListResponseData) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *StatisticListResponseData) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *StatisticListResponseData) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *StatisticListResponseData) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUnit

`func (o *StatisticListResponseData) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *StatisticListResponseData) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *StatisticListResponseData) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *StatisticListResponseData) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetAlgo

`func (o *StatisticListResponseData) GetAlgo() string`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *StatisticListResponseData) GetAlgoOk() (*string, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *StatisticListResponseData) SetAlgo(v string)`

SetAlgo sets Algo field to given value.

### HasAlgo

`func (o *StatisticListResponseData) HasAlgo() bool`

HasAlgo returns a boolean if a field has been set.


[[Back to README]](../README.md)


