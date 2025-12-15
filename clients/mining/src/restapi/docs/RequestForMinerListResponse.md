# RequestForMinerListResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**RequestForMinerListResponseData**](RequestForMinerListResponseData.md) |  | [optional] 

## Methods

### NewRequestForMinerListResponse

`func NewRequestForMinerListResponse() *RequestForMinerListResponse`

NewRequestForMinerListResponse instantiates a new RequestForMinerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestForMinerListResponseWithDefaults

`func NewRequestForMinerListResponseWithDefaults() *RequestForMinerListResponse`

NewRequestForMinerListResponseWithDefaults instantiates a new RequestForMinerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *RequestForMinerListResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RequestForMinerListResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RequestForMinerListResponse) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *RequestForMinerListResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *RequestForMinerListResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *RequestForMinerListResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *RequestForMinerListResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *RequestForMinerListResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *RequestForMinerListResponse) GetData() RequestForMinerListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestForMinerListResponse) GetDataOk() (*RequestForMinerListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestForMinerListResponse) SetData(v RequestForMinerListResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *RequestForMinerListResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


