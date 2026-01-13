# GetOrderModifyHistoryResponseInnerAmendment

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to [**GetOrderModifyHistoryResponseInnerAmendmentPrice**](GetOrderModifyHistoryResponseInnerAmendmentPrice.md) |  | [optional] 
**OrigQty** | Pointer to [**GetOrderModifyHistoryResponseInnerAmendmentOrigQty**](GetOrderModifyHistoryResponseInnerAmendmentOrigQty.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetOrderModifyHistoryResponseInnerAmendment

`func NewGetOrderModifyHistoryResponseInnerAmendment() *GetOrderModifyHistoryResponseInnerAmendment`

NewGetOrderModifyHistoryResponseInnerAmendment instantiates a new GetOrderModifyHistoryResponseInnerAmendment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderModifyHistoryResponseInnerAmendmentWithDefaults

`func NewGetOrderModifyHistoryResponseInnerAmendmentWithDefaults() *GetOrderModifyHistoryResponseInnerAmendment`

NewGetOrderModifyHistoryResponseInnerAmendmentWithDefaults instantiates a new GetOrderModifyHistoryResponseInnerAmendment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *GetOrderModifyHistoryResponseInnerAmendment) GetPrice() GetOrderModifyHistoryResponseInnerAmendmentPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetOrderModifyHistoryResponseInnerAmendment) GetPriceOk() (*GetOrderModifyHistoryResponseInnerAmendmentPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetOrderModifyHistoryResponseInnerAmendment) SetPrice(v GetOrderModifyHistoryResponseInnerAmendmentPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetOrderModifyHistoryResponseInnerAmendment) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *GetOrderModifyHistoryResponseInnerAmendment) GetOrigQty() GetOrderModifyHistoryResponseInnerAmendmentOrigQty`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *GetOrderModifyHistoryResponseInnerAmendment) GetOrigQtyOk() (*GetOrderModifyHistoryResponseInnerAmendmentOrigQty, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *GetOrderModifyHistoryResponseInnerAmendment) SetOrigQty(v GetOrderModifyHistoryResponseInnerAmendmentOrigQty)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *GetOrderModifyHistoryResponseInnerAmendment) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetCount

`func (o *GetOrderModifyHistoryResponseInnerAmendment) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetOrderModifyHistoryResponseInnerAmendment) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetOrderModifyHistoryResponseInnerAmendment) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetOrderModifyHistoryResponseInnerAmendment) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to README]](../README.md)


