# TradingScheduleResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**UpdateTime** | Pointer to **int64** |  | [optional] 
**MarketSchedules** | Pointer to [**TradingScheduleResponseMarketSchedules**](TradingScheduleResponseMarketSchedules.md) |  | [optional] 

## Methods

### NewTradingScheduleResponse

`func NewTradingScheduleResponse() *TradingScheduleResponse`

NewTradingScheduleResponse instantiates a new TradingScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradingScheduleResponseWithDefaults

`func NewTradingScheduleResponseWithDefaults() *TradingScheduleResponse`

NewTradingScheduleResponseWithDefaults instantiates a new TradingScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateTime

`func (o *TradingScheduleResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *TradingScheduleResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *TradingScheduleResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *TradingScheduleResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetMarketSchedules

`func (o *TradingScheduleResponse) GetMarketSchedules() TradingScheduleResponseMarketSchedules`

GetMarketSchedules returns the MarketSchedules field if non-nil, zero value otherwise.

### GetMarketSchedulesOk

`func (o *TradingScheduleResponse) GetMarketSchedulesOk() (*TradingScheduleResponseMarketSchedules, bool)`

GetMarketSchedulesOk returns a tuple with the MarketSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketSchedules

`func (o *TradingScheduleResponse) SetMarketSchedules(v TradingScheduleResponseMarketSchedules)`

SetMarketSchedules sets MarketSchedules field to given value.

### HasMarketSchedules

`func (o *TradingScheduleResponse) HasMarketSchedules() bool`

HasMarketSchedules returns a boolean if a field has been set.


[[Back to README]](../README.md)


