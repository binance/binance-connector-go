# CancelAllEquityOrdersResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | &#x60;true&#x60; when the cancel-all request was accepted by upstream. Per-order outcomes live in &#x60;/order/history&#x60;. | [optional] 

## Methods

### NewCancelAllEquityOrdersResponse

`func NewCancelAllEquityOrdersResponse() *CancelAllEquityOrdersResponse`

NewCancelAllEquityOrdersResponse instantiates a new CancelAllEquityOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelAllEquityOrdersResponseWithDefaults

`func NewCancelAllEquityOrdersResponseWithDefaults() *CancelAllEquityOrdersResponse`

NewCancelAllEquityOrdersResponseWithDefaults instantiates a new CancelAllEquityOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CancelAllEquityOrdersResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CancelAllEquityOrdersResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CancelAllEquityOrdersResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CancelAllEquityOrdersResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to README]](../README.md)


