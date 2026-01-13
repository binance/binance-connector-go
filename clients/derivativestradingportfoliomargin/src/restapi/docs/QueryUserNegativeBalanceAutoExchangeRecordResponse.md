# QueryUserNegativeBalanceAutoExchangeRecordResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Rows** | Pointer to [**[]QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner**](QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner.md) |  | [optional] 

## Methods

### NewQueryUserNegativeBalanceAutoExchangeRecordResponse

`func NewQueryUserNegativeBalanceAutoExchangeRecordResponse() *QueryUserNegativeBalanceAutoExchangeRecordResponse`

NewQueryUserNegativeBalanceAutoExchangeRecordResponse instantiates a new QueryUserNegativeBalanceAutoExchangeRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryUserNegativeBalanceAutoExchangeRecordResponseWithDefaults

`func NewQueryUserNegativeBalanceAutoExchangeRecordResponseWithDefaults() *QueryUserNegativeBalanceAutoExchangeRecordResponse`

NewQueryUserNegativeBalanceAutoExchangeRecordResponseWithDefaults instantiates a new QueryUserNegativeBalanceAutoExchangeRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRows

`func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) GetRows() []QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) GetRowsOk() (*[]QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) SetRows(v []QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to README]](../README.md)


