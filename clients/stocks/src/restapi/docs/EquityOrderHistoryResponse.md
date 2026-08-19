# EquityOrderHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** | Total number of rows matching the filter. | [optional] 
**Page** | Pointer to **int32** | Current page (echoes &#x60;current&#x60;). | [optional] 
**Size** | Pointer to **int32** | Current page size (echoes &#x60;size&#x60;). | [optional] 
**Rows** | Pointer to [**[]EquityOrderHistoryResponseRowsInner**](EquityOrderHistoryResponseRowsInner.md) | Order rows on this page. Empty array if nothing matches. | [optional] 

## Methods

### NewEquityOrderHistoryResponse

`func NewEquityOrderHistoryResponse() *EquityOrderHistoryResponse`

NewEquityOrderHistoryResponse instantiates a new EquityOrderHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityOrderHistoryResponseWithDefaults

`func NewEquityOrderHistoryResponseWithDefaults() *EquityOrderHistoryResponse`

NewEquityOrderHistoryResponseWithDefaults instantiates a new EquityOrderHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *EquityOrderHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EquityOrderHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EquityOrderHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *EquityOrderHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *EquityOrderHistoryResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EquityOrderHistoryResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EquityOrderHistoryResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *EquityOrderHistoryResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *EquityOrderHistoryResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *EquityOrderHistoryResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *EquityOrderHistoryResponse) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *EquityOrderHistoryResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetRows

`func (o *EquityOrderHistoryResponse) GetRows() []EquityOrderHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *EquityOrderHistoryResponse) GetRowsOk() (*[]EquityOrderHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *EquityOrderHistoryResponse) SetRows(v []EquityOrderHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *EquityOrderHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to README]](../README.md)


