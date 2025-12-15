# StatisticListResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**StatisticListResponseData**](StatisticListResponseData.md) |  | [optional] 

## Methods

### NewStatisticListResponse

`func NewStatisticListResponse() *StatisticListResponse`

NewStatisticListResponse instantiates a new StatisticListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticListResponseWithDefaults

`func NewStatisticListResponseWithDefaults() *StatisticListResponse`

NewStatisticListResponseWithDefaults instantiates a new StatisticListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *StatisticListResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StatisticListResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StatisticListResponse) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *StatisticListResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *StatisticListResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *StatisticListResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *StatisticListResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *StatisticListResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *StatisticListResponse) GetData() StatisticListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StatisticListResponse) GetDataOk() (*StatisticListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StatisticListResponse) SetData(v StatisticListResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *StatisticListResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


