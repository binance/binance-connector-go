# EquityTradeHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** | Total number of rows matching the filter. | [optional] 
**Page** | Pointer to **int32** | Current page. | [optional] 
**Size** | Pointer to **int32** | Current page size. | [optional] 
**Rows** | Pointer to [**[]EquityTradeHistoryResponseRowsInner**](EquityTradeHistoryResponseRowsInner.md) | Trade rows on this page. Empty when nothing matches. | [optional] 

## Methods

### NewEquityTradeHistoryResponse

`func NewEquityTradeHistoryResponse() *EquityTradeHistoryResponse`

NewEquityTradeHistoryResponse instantiates a new EquityTradeHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityTradeHistoryResponseWithDefaults

`func NewEquityTradeHistoryResponseWithDefaults() *EquityTradeHistoryResponse`

NewEquityTradeHistoryResponseWithDefaults instantiates a new EquityTradeHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *EquityTradeHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EquityTradeHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EquityTradeHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *EquityTradeHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *EquityTradeHistoryResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EquityTradeHistoryResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EquityTradeHistoryResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *EquityTradeHistoryResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *EquityTradeHistoryResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *EquityTradeHistoryResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *EquityTradeHistoryResponse) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *EquityTradeHistoryResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetRows

`func (o *EquityTradeHistoryResponse) GetRows() []EquityTradeHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *EquityTradeHistoryResponse) GetRowsOk() (*[]EquityTradeHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *EquityTradeHistoryResponse) SetRows(v []EquityTradeHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *EquityTradeHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to README]](../README.md)


