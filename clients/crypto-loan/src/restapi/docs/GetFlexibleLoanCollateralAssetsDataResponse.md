# GetFlexibleLoanCollateralAssetsDataResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetFlexibleLoanCollateralAssetsDataResponseRowsInner**](GetFlexibleLoanCollateralAssetsDataResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetFlexibleLoanCollateralAssetsDataResponse

`func NewGetFlexibleLoanCollateralAssetsDataResponse() *GetFlexibleLoanCollateralAssetsDataResponse`

NewGetFlexibleLoanCollateralAssetsDataResponse instantiates a new GetFlexibleLoanCollateralAssetsDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFlexibleLoanCollateralAssetsDataResponseWithDefaults

`func NewGetFlexibleLoanCollateralAssetsDataResponseWithDefaults() *GetFlexibleLoanCollateralAssetsDataResponse`

NewGetFlexibleLoanCollateralAssetsDataResponseWithDefaults instantiates a new GetFlexibleLoanCollateralAssetsDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetFlexibleLoanCollateralAssetsDataResponse) GetRows() []GetFlexibleLoanCollateralAssetsDataResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetFlexibleLoanCollateralAssetsDataResponse) GetRowsOk() (*[]GetFlexibleLoanCollateralAssetsDataResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetFlexibleLoanCollateralAssetsDataResponse) SetRows(v []GetFlexibleLoanCollateralAssetsDataResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetFlexibleLoanCollateralAssetsDataResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetFlexibleLoanCollateralAssetsDataResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetFlexibleLoanCollateralAssetsDataResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetFlexibleLoanCollateralAssetsDataResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetFlexibleLoanCollateralAssetsDataResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


