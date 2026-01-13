# QueryCrossMarginAccountDetailsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** |  | [optional] 
**BorrowEnabled** | Pointer to **bool** |  | [optional] 
**MarginLevel** | Pointer to **string** |  | [optional] 
**CollateralMarginLevel** | Pointer to **string** |  | [optional] 
**TotalAssetOfBtc** | Pointer to **string** |  | [optional] 
**TotalLiabilityOfBtc** | Pointer to **string** |  | [optional] 
**TotalNetAssetOfBtc** | Pointer to **string** |  | [optional] 
**TotalCollateralValueInUSDT** | Pointer to **string** |  | [optional] 
**TotalOpenOrderLossInUSDT** | Pointer to **string** |  | [optional] 
**TradeEnabled** | Pointer to **bool** |  | [optional] 
**TransferInEnabled** | Pointer to **bool** |  | [optional] 
**TransferOutEnabled** | Pointer to **bool** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**UserAssets** | Pointer to [**[]QueryCrossMarginAccountDetailsResponseUserAssetsInner**](QueryCrossMarginAccountDetailsResponseUserAssetsInner.md) |  | [optional] 

## Methods

### NewQueryCrossMarginAccountDetailsResponse

`func NewQueryCrossMarginAccountDetailsResponse() *QueryCrossMarginAccountDetailsResponse`

NewQueryCrossMarginAccountDetailsResponse instantiates a new QueryCrossMarginAccountDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCrossMarginAccountDetailsResponseWithDefaults

`func NewQueryCrossMarginAccountDetailsResponseWithDefaults() *QueryCrossMarginAccountDetailsResponse`

NewQueryCrossMarginAccountDetailsResponseWithDefaults instantiates a new QueryCrossMarginAccountDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *QueryCrossMarginAccountDetailsResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *QueryCrossMarginAccountDetailsResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *QueryCrossMarginAccountDetailsResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetBorrowEnabled

`func (o *QueryCrossMarginAccountDetailsResponse) GetBorrowEnabled() bool`

GetBorrowEnabled returns the BorrowEnabled field if non-nil, zero value otherwise.

### GetBorrowEnabledOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetBorrowEnabledOk() (*bool, bool)`

GetBorrowEnabledOk returns a tuple with the BorrowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowEnabled

`func (o *QueryCrossMarginAccountDetailsResponse) SetBorrowEnabled(v bool)`

SetBorrowEnabled sets BorrowEnabled field to given value.

### HasBorrowEnabled

`func (o *QueryCrossMarginAccountDetailsResponse) HasBorrowEnabled() bool`

HasBorrowEnabled returns a boolean if a field has been set.

### GetMarginLevel

`func (o *QueryCrossMarginAccountDetailsResponse) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *QueryCrossMarginAccountDetailsResponse) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.

### HasMarginLevel

`func (o *QueryCrossMarginAccountDetailsResponse) HasMarginLevel() bool`

HasMarginLevel returns a boolean if a field has been set.

### GetCollateralMarginLevel

`func (o *QueryCrossMarginAccountDetailsResponse) GetCollateralMarginLevel() string`

GetCollateralMarginLevel returns the CollateralMarginLevel field if non-nil, zero value otherwise.

### GetCollateralMarginLevelOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetCollateralMarginLevelOk() (*string, bool)`

GetCollateralMarginLevelOk returns a tuple with the CollateralMarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralMarginLevel

`func (o *QueryCrossMarginAccountDetailsResponse) SetCollateralMarginLevel(v string)`

SetCollateralMarginLevel sets CollateralMarginLevel field to given value.

### HasCollateralMarginLevel

`func (o *QueryCrossMarginAccountDetailsResponse) HasCollateralMarginLevel() bool`

HasCollateralMarginLevel returns a boolean if a field has been set.

### GetTotalAssetOfBtc

`func (o *QueryCrossMarginAccountDetailsResponse) GetTotalAssetOfBtc() string`

GetTotalAssetOfBtc returns the TotalAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalAssetOfBtcOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetTotalAssetOfBtcOk() (*string, bool)`

GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssetOfBtc

`func (o *QueryCrossMarginAccountDetailsResponse) SetTotalAssetOfBtc(v string)`

SetTotalAssetOfBtc sets TotalAssetOfBtc field to given value.

### HasTotalAssetOfBtc

`func (o *QueryCrossMarginAccountDetailsResponse) HasTotalAssetOfBtc() bool`

HasTotalAssetOfBtc returns a boolean if a field has been set.

### GetTotalLiabilityOfBtc

`func (o *QueryCrossMarginAccountDetailsResponse) GetTotalLiabilityOfBtc() string`

GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field if non-nil, zero value otherwise.

### GetTotalLiabilityOfBtcOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetTotalLiabilityOfBtcOk() (*string, bool)`

GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilityOfBtc

`func (o *QueryCrossMarginAccountDetailsResponse) SetTotalLiabilityOfBtc(v string)`

SetTotalLiabilityOfBtc sets TotalLiabilityOfBtc field to given value.

### HasTotalLiabilityOfBtc

`func (o *QueryCrossMarginAccountDetailsResponse) HasTotalLiabilityOfBtc() bool`

HasTotalLiabilityOfBtc returns a boolean if a field has been set.

### GetTotalNetAssetOfBtc

`func (o *QueryCrossMarginAccountDetailsResponse) GetTotalNetAssetOfBtc() string`

GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalNetAssetOfBtcOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetTotalNetAssetOfBtcOk() (*string, bool)`

GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAssetOfBtc

`func (o *QueryCrossMarginAccountDetailsResponse) SetTotalNetAssetOfBtc(v string)`

SetTotalNetAssetOfBtc sets TotalNetAssetOfBtc field to given value.

### HasTotalNetAssetOfBtc

`func (o *QueryCrossMarginAccountDetailsResponse) HasTotalNetAssetOfBtc() bool`

HasTotalNetAssetOfBtc returns a boolean if a field has been set.

### GetTotalCollateralValueInUSDT

`func (o *QueryCrossMarginAccountDetailsResponse) GetTotalCollateralValueInUSDT() string`

GetTotalCollateralValueInUSDT returns the TotalCollateralValueInUSDT field if non-nil, zero value otherwise.

### GetTotalCollateralValueInUSDTOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetTotalCollateralValueInUSDTOk() (*string, bool)`

GetTotalCollateralValueInUSDTOk returns a tuple with the TotalCollateralValueInUSDT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCollateralValueInUSDT

`func (o *QueryCrossMarginAccountDetailsResponse) SetTotalCollateralValueInUSDT(v string)`

SetTotalCollateralValueInUSDT sets TotalCollateralValueInUSDT field to given value.

### HasTotalCollateralValueInUSDT

`func (o *QueryCrossMarginAccountDetailsResponse) HasTotalCollateralValueInUSDT() bool`

HasTotalCollateralValueInUSDT returns a boolean if a field has been set.

### GetTotalOpenOrderLossInUSDT

`func (o *QueryCrossMarginAccountDetailsResponse) GetTotalOpenOrderLossInUSDT() string`

GetTotalOpenOrderLossInUSDT returns the TotalOpenOrderLossInUSDT field if non-nil, zero value otherwise.

### GetTotalOpenOrderLossInUSDTOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetTotalOpenOrderLossInUSDTOk() (*string, bool)`

GetTotalOpenOrderLossInUSDTOk returns a tuple with the TotalOpenOrderLossInUSDT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderLossInUSDT

`func (o *QueryCrossMarginAccountDetailsResponse) SetTotalOpenOrderLossInUSDT(v string)`

SetTotalOpenOrderLossInUSDT sets TotalOpenOrderLossInUSDT field to given value.

### HasTotalOpenOrderLossInUSDT

`func (o *QueryCrossMarginAccountDetailsResponse) HasTotalOpenOrderLossInUSDT() bool`

HasTotalOpenOrderLossInUSDT returns a boolean if a field has been set.

### GetTradeEnabled

`func (o *QueryCrossMarginAccountDetailsResponse) GetTradeEnabled() bool`

GetTradeEnabled returns the TradeEnabled field if non-nil, zero value otherwise.

### GetTradeEnabledOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetTradeEnabledOk() (*bool, bool)`

GetTradeEnabledOk returns a tuple with the TradeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeEnabled

`func (o *QueryCrossMarginAccountDetailsResponse) SetTradeEnabled(v bool)`

SetTradeEnabled sets TradeEnabled field to given value.

### HasTradeEnabled

`func (o *QueryCrossMarginAccountDetailsResponse) HasTradeEnabled() bool`

HasTradeEnabled returns a boolean if a field has been set.

### GetTransferInEnabled

`func (o *QueryCrossMarginAccountDetailsResponse) GetTransferInEnabled() bool`

GetTransferInEnabled returns the TransferInEnabled field if non-nil, zero value otherwise.

### GetTransferInEnabledOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetTransferInEnabledOk() (*bool, bool)`

GetTransferInEnabledOk returns a tuple with the TransferInEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferInEnabled

`func (o *QueryCrossMarginAccountDetailsResponse) SetTransferInEnabled(v bool)`

SetTransferInEnabled sets TransferInEnabled field to given value.

### HasTransferInEnabled

`func (o *QueryCrossMarginAccountDetailsResponse) HasTransferInEnabled() bool`

HasTransferInEnabled returns a boolean if a field has been set.

### GetTransferOutEnabled

`func (o *QueryCrossMarginAccountDetailsResponse) GetTransferOutEnabled() bool`

GetTransferOutEnabled returns the TransferOutEnabled field if non-nil, zero value otherwise.

### GetTransferOutEnabledOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetTransferOutEnabledOk() (*bool, bool)`

GetTransferOutEnabledOk returns a tuple with the TransferOutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferOutEnabled

`func (o *QueryCrossMarginAccountDetailsResponse) SetTransferOutEnabled(v bool)`

SetTransferOutEnabled sets TransferOutEnabled field to given value.

### HasTransferOutEnabled

`func (o *QueryCrossMarginAccountDetailsResponse) HasTransferOutEnabled() bool`

HasTransferOutEnabled returns a boolean if a field has been set.

### GetAccountType

`func (o *QueryCrossMarginAccountDetailsResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *QueryCrossMarginAccountDetailsResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *QueryCrossMarginAccountDetailsResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetUserAssets

`func (o *QueryCrossMarginAccountDetailsResponse) GetUserAssets() []QueryCrossMarginAccountDetailsResponseUserAssetsInner`

GetUserAssets returns the UserAssets field if non-nil, zero value otherwise.

### GetUserAssetsOk

`func (o *QueryCrossMarginAccountDetailsResponse) GetUserAssetsOk() (*[]QueryCrossMarginAccountDetailsResponseUserAssetsInner, bool)`

GetUserAssetsOk returns a tuple with the UserAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssets

`func (o *QueryCrossMarginAccountDetailsResponse) SetUserAssets(v []QueryCrossMarginAccountDetailsResponseUserAssetsInner)`

SetUserAssets sets UserAssets field to given value.

### HasUserAssets

`func (o *QueryCrossMarginAccountDetailsResponse) HasUserAssets() bool`

HasUserAssets returns a boolean if a field has been set.


[[Back to README]](../README.md)


