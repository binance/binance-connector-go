# AssetDividendRecordResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]AssetDividendRecordResponseRowsInner**](AssetDividendRecordResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewAssetDividendRecordResponse

`func NewAssetDividendRecordResponse() *AssetDividendRecordResponse`

NewAssetDividendRecordResponse instantiates a new AssetDividendRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDividendRecordResponseWithDefaults

`func NewAssetDividendRecordResponseWithDefaults() *AssetDividendRecordResponse`

NewAssetDividendRecordResponseWithDefaults instantiates a new AssetDividendRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *AssetDividendRecordResponse) GetRows() []AssetDividendRecordResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *AssetDividendRecordResponse) GetRowsOk() (*[]AssetDividendRecordResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *AssetDividendRecordResponse) SetRows(v []AssetDividendRecordResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *AssetDividendRecordResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *AssetDividendRecordResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AssetDividendRecordResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AssetDividendRecordResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AssetDividendRecordResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


