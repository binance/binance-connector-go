# CancelAlgoOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | Pointer to **int64** |  | [optional] 
**ClientAlgoId** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewCancelAlgoOrderResponse

`func NewCancelAlgoOrderResponse() *CancelAlgoOrderResponse`

NewCancelAlgoOrderResponse instantiates a new CancelAlgoOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelAlgoOrderResponseWithDefaults

`func NewCancelAlgoOrderResponseWithDefaults() *CancelAlgoOrderResponse`

NewCancelAlgoOrderResponseWithDefaults instantiates a new CancelAlgoOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CancelAlgoOrderResponse) GetAlgoId() int64`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CancelAlgoOrderResponse) GetAlgoIdOk() (*int64, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CancelAlgoOrderResponse) SetAlgoId(v int64)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *CancelAlgoOrderResponse) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetClientAlgoId

`func (o *CancelAlgoOrderResponse) GetClientAlgoId() string`

GetClientAlgoId returns the ClientAlgoId field if non-nil, zero value otherwise.

### GetClientAlgoIdOk

`func (o *CancelAlgoOrderResponse) GetClientAlgoIdOk() (*string, bool)`

GetClientAlgoIdOk returns a tuple with the ClientAlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAlgoId

`func (o *CancelAlgoOrderResponse) SetClientAlgoId(v string)`

SetClientAlgoId sets ClientAlgoId field to given value.

### HasClientAlgoId

`func (o *CancelAlgoOrderResponse) HasClientAlgoId() bool`

HasClientAlgoId returns a boolean if a field has been set.

### GetCode

`func (o *CancelAlgoOrderResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CancelAlgoOrderResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CancelAlgoOrderResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CancelAlgoOrderResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *CancelAlgoOrderResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CancelAlgoOrderResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CancelAlgoOrderResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CancelAlgoOrderResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to README]](../README.md)


