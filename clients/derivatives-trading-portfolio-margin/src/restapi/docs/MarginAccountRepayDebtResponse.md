# MarginAccountRepayDebtResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
**SpecifyRepayAssets** | Pointer to **[]string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewMarginAccountRepayDebtResponse

`func NewMarginAccountRepayDebtResponse() *MarginAccountRepayDebtResponse`

NewMarginAccountRepayDebtResponse instantiates a new MarginAccountRepayDebtResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginAccountRepayDebtResponseWithDefaults

`func NewMarginAccountRepayDebtResponseWithDefaults() *MarginAccountRepayDebtResponse`

NewMarginAccountRepayDebtResponseWithDefaults instantiates a new MarginAccountRepayDebtResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *MarginAccountRepayDebtResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MarginAccountRepayDebtResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MarginAccountRepayDebtResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *MarginAccountRepayDebtResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAsset

`func (o *MarginAccountRepayDebtResponse) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *MarginAccountRepayDebtResponse) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *MarginAccountRepayDebtResponse) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *MarginAccountRepayDebtResponse) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetSpecifyRepayAssets

`func (o *MarginAccountRepayDebtResponse) GetSpecifyRepayAssets() []string`

GetSpecifyRepayAssets returns the SpecifyRepayAssets field if non-nil, zero value otherwise.

### GetSpecifyRepayAssetsOk

`func (o *MarginAccountRepayDebtResponse) GetSpecifyRepayAssetsOk() (*[]string, bool)`

GetSpecifyRepayAssetsOk returns a tuple with the SpecifyRepayAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifyRepayAssets

`func (o *MarginAccountRepayDebtResponse) SetSpecifyRepayAssets(v []string)`

SetSpecifyRepayAssets sets SpecifyRepayAssets field to given value.

### HasSpecifyRepayAssets

`func (o *MarginAccountRepayDebtResponse) HasSpecifyRepayAssets() bool`

HasSpecifyRepayAssets returns a boolean if a field has been set.

### GetUpdateTime

`func (o *MarginAccountRepayDebtResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *MarginAccountRepayDebtResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *MarginAccountRepayDebtResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *MarginAccountRepayDebtResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetSuccess

`func (o *MarginAccountRepayDebtResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MarginAccountRepayDebtResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MarginAccountRepayDebtResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *MarginAccountRepayDebtResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to README]](../README.md)


