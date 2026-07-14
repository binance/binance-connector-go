# QueryVIPLoanFixedRateMarketResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** | Total number of records | [optional] 
**Rows** | Pointer to [**[]QueryVIPLoanFixedRateMarketResponseRowsInner**](QueryVIPLoanFixedRateMarketResponseRowsInner.md) | Current page data | [optional] 

## Methods

### NewQueryVIPLoanFixedRateMarketResponse

`func NewQueryVIPLoanFixedRateMarketResponse() *QueryVIPLoanFixedRateMarketResponse`

NewQueryVIPLoanFixedRateMarketResponse instantiates a new QueryVIPLoanFixedRateMarketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryVIPLoanFixedRateMarketResponseWithDefaults

`func NewQueryVIPLoanFixedRateMarketResponseWithDefaults() *QueryVIPLoanFixedRateMarketResponse`

NewQueryVIPLoanFixedRateMarketResponseWithDefaults instantiates a new QueryVIPLoanFixedRateMarketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *QueryVIPLoanFixedRateMarketResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryVIPLoanFixedRateMarketResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryVIPLoanFixedRateMarketResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryVIPLoanFixedRateMarketResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRows

`func (o *QueryVIPLoanFixedRateMarketResponse) GetRows() []QueryVIPLoanFixedRateMarketResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *QueryVIPLoanFixedRateMarketResponse) GetRowsOk() (*[]QueryVIPLoanFixedRateMarketResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *QueryVIPLoanFixedRateMarketResponse) SetRows(v []QueryVIPLoanFixedRateMarketResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *QueryVIPLoanFixedRateMarketResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to README]](../README.md)


