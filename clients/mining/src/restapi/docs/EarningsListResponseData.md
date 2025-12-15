# EarningsListResponseData

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AccountProfits** | Pointer to [**[]EarningsListResponseDataAccountProfitsInner**](EarningsListResponseDataAccountProfitsInner.md) |  | [optional] 
**TotalNum** | Pointer to **int64** |  | [optional] 
**PageSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewEarningsListResponseData

`func NewEarningsListResponseData() *EarningsListResponseData`

NewEarningsListResponseData instantiates a new EarningsListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarningsListResponseDataWithDefaults

`func NewEarningsListResponseDataWithDefaults() *EarningsListResponseData`

NewEarningsListResponseDataWithDefaults instantiates a new EarningsListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountProfits

`func (o *EarningsListResponseData) GetAccountProfits() []EarningsListResponseDataAccountProfitsInner`

GetAccountProfits returns the AccountProfits field if non-nil, zero value otherwise.

### GetAccountProfitsOk

`func (o *EarningsListResponseData) GetAccountProfitsOk() (*[]EarningsListResponseDataAccountProfitsInner, bool)`

GetAccountProfitsOk returns a tuple with the AccountProfits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountProfits

`func (o *EarningsListResponseData) SetAccountProfits(v []EarningsListResponseDataAccountProfitsInner)`

SetAccountProfits sets AccountProfits field to given value.

### HasAccountProfits

`func (o *EarningsListResponseData) HasAccountProfits() bool`

HasAccountProfits returns a boolean if a field has been set.

### GetTotalNum

`func (o *EarningsListResponseData) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *EarningsListResponseData) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *EarningsListResponseData) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.

### HasTotalNum

`func (o *EarningsListResponseData) HasTotalNum() bool`

HasTotalNum returns a boolean if a field has been set.

### GetPageSize

`func (o *EarningsListResponseData) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EarningsListResponseData) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EarningsListResponseData) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *EarningsListResponseData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to README]](../README.md)


