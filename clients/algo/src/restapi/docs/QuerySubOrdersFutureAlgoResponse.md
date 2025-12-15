# QuerySubOrdersFutureAlgoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**ExecutedAmt** | Pointer to **string** |  | [optional] 
**SubOrders** | Pointer to [**[]QuerySubOrdersFutureAlgoResponseSubOrdersInner**](QuerySubOrdersFutureAlgoResponseSubOrdersInner.md) |  | [optional] 

## Methods

### NewQuerySubOrdersFutureAlgoResponse

`func NewQuerySubOrdersFutureAlgoResponse() *QuerySubOrdersFutureAlgoResponse`

NewQuerySubOrdersFutureAlgoResponse instantiates a new QuerySubOrdersFutureAlgoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySubOrdersFutureAlgoResponseWithDefaults

`func NewQuerySubOrdersFutureAlgoResponseWithDefaults() *QuerySubOrdersFutureAlgoResponse`

NewQuerySubOrdersFutureAlgoResponseWithDefaults instantiates a new QuerySubOrdersFutureAlgoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *QuerySubOrdersFutureAlgoResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QuerySubOrdersFutureAlgoResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QuerySubOrdersFutureAlgoResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QuerySubOrdersFutureAlgoResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetExecutedQty

`func (o *QuerySubOrdersFutureAlgoResponse) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *QuerySubOrdersFutureAlgoResponse) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *QuerySubOrdersFutureAlgoResponse) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *QuerySubOrdersFutureAlgoResponse) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetExecutedAmt

`func (o *QuerySubOrdersFutureAlgoResponse) GetExecutedAmt() string`

GetExecutedAmt returns the ExecutedAmt field if non-nil, zero value otherwise.

### GetExecutedAmtOk

`func (o *QuerySubOrdersFutureAlgoResponse) GetExecutedAmtOk() (*string, bool)`

GetExecutedAmtOk returns a tuple with the ExecutedAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAmt

`func (o *QuerySubOrdersFutureAlgoResponse) SetExecutedAmt(v string)`

SetExecutedAmt sets ExecutedAmt field to given value.

### HasExecutedAmt

`func (o *QuerySubOrdersFutureAlgoResponse) HasExecutedAmt() bool`

HasExecutedAmt returns a boolean if a field has been set.

### GetSubOrders

`func (o *QuerySubOrdersFutureAlgoResponse) GetSubOrders() []QuerySubOrdersFutureAlgoResponseSubOrdersInner`

GetSubOrders returns the SubOrders field if non-nil, zero value otherwise.

### GetSubOrdersOk

`func (o *QuerySubOrdersFutureAlgoResponse) GetSubOrdersOk() (*[]QuerySubOrdersFutureAlgoResponseSubOrdersInner, bool)`

GetSubOrdersOk returns a tuple with the SubOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrders

`func (o *QuerySubOrdersFutureAlgoResponse) SetSubOrders(v []QuerySubOrdersFutureAlgoResponseSubOrdersInner)`

SetSubOrders sets SubOrders field to given value.

### HasSubOrders

`func (o *QuerySubOrdersFutureAlgoResponse) HasSubOrders() bool`

HasSubOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


