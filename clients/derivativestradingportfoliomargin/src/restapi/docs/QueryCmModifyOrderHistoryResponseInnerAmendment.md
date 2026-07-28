# QueryCmModifyOrderHistoryResponseInnerAmendment

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to [**QueryCmModifyOrderHistoryResponseInnerAmendmentPrice**](QueryCmModifyOrderHistoryResponseInnerAmendmentPrice.md) |  | [optional] 
**OrigQty** | Pointer to [**QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty**](QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty.md) |  | [optional] 
**Count** | Pointer to **int64** | Order modification count, representing the number of times the order has been modified | [optional] 
**ModifyId** | Pointer to **int64** | user-defined modification identifier, only returned if provided in the request | [optional] 

## Methods

### NewQueryCmModifyOrderHistoryResponseInnerAmendment

`func NewQueryCmModifyOrderHistoryResponseInnerAmendment() *QueryCmModifyOrderHistoryResponseInnerAmendment`

NewQueryCmModifyOrderHistoryResponseInnerAmendment instantiates a new QueryCmModifyOrderHistoryResponseInnerAmendment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCmModifyOrderHistoryResponseInnerAmendmentWithDefaults

`func NewQueryCmModifyOrderHistoryResponseInnerAmendmentWithDefaults() *QueryCmModifyOrderHistoryResponseInnerAmendment`

NewQueryCmModifyOrderHistoryResponseInnerAmendmentWithDefaults instantiates a new QueryCmModifyOrderHistoryResponseInnerAmendment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetPrice() QueryCmModifyOrderHistoryResponseInnerAmendmentPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetPriceOk() (*QueryCmModifyOrderHistoryResponseInnerAmendmentPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) SetPrice(v QueryCmModifyOrderHistoryResponseInnerAmendmentPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetOrigQty() QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetOrigQtyOk() (*QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) SetOrigQty(v QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetCount

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetModifyId

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetModifyId() int64`

GetModifyId returns the ModifyId field if non-nil, zero value otherwise.

### GetModifyIdOk

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetModifyIdOk() (*int64, bool)`

GetModifyIdOk returns a tuple with the ModifyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyId

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) SetModifyId(v int64)`

SetModifyId sets ModifyId field to given value.

### HasModifyId

`func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) HasModifyId() bool`

HasModifyId returns a boolean if a field has been set.


[[Back to README]](../README.md)


