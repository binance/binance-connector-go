# GetMovePositionHistoryForSubAccountResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**FutureMovePositionOrderVoList** | Pointer to [**[]GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner**](GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner.md) |  | [optional] 

## Methods

### NewGetMovePositionHistoryForSubAccountResponse

`func NewGetMovePositionHistoryForSubAccountResponse() *GetMovePositionHistoryForSubAccountResponse`

NewGetMovePositionHistoryForSubAccountResponse instantiates a new GetMovePositionHistoryForSubAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMovePositionHistoryForSubAccountResponseWithDefaults

`func NewGetMovePositionHistoryForSubAccountResponseWithDefaults() *GetMovePositionHistoryForSubAccountResponse`

NewGetMovePositionHistoryForSubAccountResponseWithDefaults instantiates a new GetMovePositionHistoryForSubAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetMovePositionHistoryForSubAccountResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetMovePositionHistoryForSubAccountResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetMovePositionHistoryForSubAccountResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetMovePositionHistoryForSubAccountResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetFutureMovePositionOrderVoList

`func (o *GetMovePositionHistoryForSubAccountResponse) GetFutureMovePositionOrderVoList() []GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner`

GetFutureMovePositionOrderVoList returns the FutureMovePositionOrderVoList field if non-nil, zero value otherwise.

### GetFutureMovePositionOrderVoListOk

`func (o *GetMovePositionHistoryForSubAccountResponse) GetFutureMovePositionOrderVoListOk() (*[]GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner, bool)`

GetFutureMovePositionOrderVoListOk returns a tuple with the FutureMovePositionOrderVoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureMovePositionOrderVoList

`func (o *GetMovePositionHistoryForSubAccountResponse) SetFutureMovePositionOrderVoList(v []GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner)`

SetFutureMovePositionOrderVoList sets FutureMovePositionOrderVoList field to given value.

### HasFutureMovePositionOrderVoList

`func (o *GetMovePositionHistoryForSubAccountResponse) HasFutureMovePositionOrderVoList() bool`

HasFutureMovePositionOrderVoList returns a boolean if a field has been set.


[[Back to README]](../README.md)


