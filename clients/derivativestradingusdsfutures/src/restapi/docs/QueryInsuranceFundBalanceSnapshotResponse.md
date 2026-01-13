# QueryInsuranceFundBalanceSnapshotResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbols** | Pointer to **[]string** |  | [optional] 
**Assets** | Pointer to [**[]QueryInsuranceFundBalanceSnapshotResponse1AssetsInner**](QueryInsuranceFundBalanceSnapshotResponse1AssetsInner.md) |  | [optional] 

## Methods

### NewQueryInsuranceFundBalanceSnapshotResponse

`func NewQueryInsuranceFundBalanceSnapshotResponse() *QueryInsuranceFundBalanceSnapshotResponse`

NewQueryInsuranceFundBalanceSnapshotResponse instantiates a new QueryInsuranceFundBalanceSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInsuranceFundBalanceSnapshotResponseWithDefaults

`func NewQueryInsuranceFundBalanceSnapshotResponseWithDefaults() *QueryInsuranceFundBalanceSnapshotResponse`

NewQueryInsuranceFundBalanceSnapshotResponseWithDefaults instantiates a new QueryInsuranceFundBalanceSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbols

`func (o *QueryInsuranceFundBalanceSnapshotResponse) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *QueryInsuranceFundBalanceSnapshotResponse) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *QueryInsuranceFundBalanceSnapshotResponse) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *QueryInsuranceFundBalanceSnapshotResponse) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetAssets

`func (o *QueryInsuranceFundBalanceSnapshotResponse) GetAssets() []QueryInsuranceFundBalanceSnapshotResponse1AssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *QueryInsuranceFundBalanceSnapshotResponse) GetAssetsOk() (*[]QueryInsuranceFundBalanceSnapshotResponse1AssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *QueryInsuranceFundBalanceSnapshotResponse) SetAssets(v []QueryInsuranceFundBalanceSnapshotResponse1AssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *QueryInsuranceFundBalanceSnapshotResponse) HasAssets() bool`

HasAssets returns a boolean if a field has been set.


[[Back to README]](../README.md)


