# AggregatedTradesResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**MessageDetail** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]AggregatedTradesResponseDataInner**](AggregatedTradesResponseDataInner.md) |  | [optional] 

## Methods

### NewAggregatedTradesResponse

`func NewAggregatedTradesResponse() *AggregatedTradesResponse`

NewAggregatedTradesResponse instantiates a new AggregatedTradesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedTradesResponseWithDefaults

`func NewAggregatedTradesResponseWithDefaults() *AggregatedTradesResponse`

NewAggregatedTradesResponseWithDefaults instantiates a new AggregatedTradesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AggregatedTradesResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AggregatedTradesResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AggregatedTradesResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AggregatedTradesResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *AggregatedTradesResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AggregatedTradesResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AggregatedTradesResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AggregatedTradesResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageDetail

`func (o *AggregatedTradesResponse) GetMessageDetail() string`

GetMessageDetail returns the MessageDetail field if non-nil, zero value otherwise.

### GetMessageDetailOk

`func (o *AggregatedTradesResponse) GetMessageDetailOk() (*string, bool)`

GetMessageDetailOk returns a tuple with the MessageDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDetail

`func (o *AggregatedTradesResponse) SetMessageDetail(v string)`

SetMessageDetail sets MessageDetail field to given value.

### HasMessageDetail

`func (o *AggregatedTradesResponse) HasMessageDetail() bool`

HasMessageDetail returns a boolean if a field has been set.

### GetData

`func (o *AggregatedTradesResponse) GetData() []AggregatedTradesResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AggregatedTradesResponse) GetDataOk() (*[]AggregatedTradesResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AggregatedTradesResponse) SetData(v []AggregatedTradesResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AggregatedTradesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


