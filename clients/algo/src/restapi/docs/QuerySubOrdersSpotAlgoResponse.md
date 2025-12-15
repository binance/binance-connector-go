# QuerySubOrdersSpotAlgoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**ExecutedAmt** | Pointer to **string** |  | [optional] 
**SubOrders** | Pointer to [**[]QuerySubOrdersFutureAlgoResponseSubOrdersInner**](QuerySubOrdersFutureAlgoResponseSubOrdersInner.md) |  | [optional] 

## Methods

### NewQuerySubOrdersSpotAlgoResponse

`func NewQuerySubOrdersSpotAlgoResponse() *QuerySubOrdersSpotAlgoResponse`

NewQuerySubOrdersSpotAlgoResponse instantiates a new QuerySubOrdersSpotAlgoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySubOrdersSpotAlgoResponseWithDefaults

`func NewQuerySubOrdersSpotAlgoResponseWithDefaults() *QuerySubOrdersSpotAlgoResponse`

NewQuerySubOrdersSpotAlgoResponseWithDefaults instantiates a new QuerySubOrdersSpotAlgoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *QuerySubOrdersSpotAlgoResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QuerySubOrdersSpotAlgoResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QuerySubOrdersSpotAlgoResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QuerySubOrdersSpotAlgoResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetExecutedQty

`func (o *QuerySubOrdersSpotAlgoResponse) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *QuerySubOrdersSpotAlgoResponse) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *QuerySubOrdersSpotAlgoResponse) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *QuerySubOrdersSpotAlgoResponse) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetExecutedAmt

`func (o *QuerySubOrdersSpotAlgoResponse) GetExecutedAmt() string`

GetExecutedAmt returns the ExecutedAmt field if non-nil, zero value otherwise.

### GetExecutedAmtOk

`func (o *QuerySubOrdersSpotAlgoResponse) GetExecutedAmtOk() (*string, bool)`

GetExecutedAmtOk returns a tuple with the ExecutedAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAmt

`func (o *QuerySubOrdersSpotAlgoResponse) SetExecutedAmt(v string)`

SetExecutedAmt sets ExecutedAmt field to given value.

### HasExecutedAmt

`func (o *QuerySubOrdersSpotAlgoResponse) HasExecutedAmt() bool`

HasExecutedAmt returns a boolean if a field has been set.

### GetSubOrders

`func (o *QuerySubOrdersSpotAlgoResponse) GetSubOrders() []QuerySubOrdersFutureAlgoResponseSubOrdersInner`

GetSubOrders returns the SubOrders field if non-nil, zero value otherwise.

### GetSubOrdersOk

`func (o *QuerySubOrdersSpotAlgoResponse) GetSubOrdersOk() (*[]QuerySubOrdersFutureAlgoResponseSubOrdersInner, bool)`

GetSubOrdersOk returns a tuple with the SubOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrders

`func (o *QuerySubOrdersSpotAlgoResponse) SetSubOrders(v []QuerySubOrdersFutureAlgoResponseSubOrdersInner)`

SetSubOrders sets SubOrders field to given value.

### HasSubOrders

`func (o *QuerySubOrdersSpotAlgoResponse) HasSubOrders() bool`

HasSubOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


