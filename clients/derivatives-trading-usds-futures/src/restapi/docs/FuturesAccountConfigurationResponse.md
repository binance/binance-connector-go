# FuturesAccountConfigurationResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FeeTier** | Pointer to **int64** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**DualSidePosition** | Pointer to **bool** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**MultiAssetsMargin** | Pointer to **bool** |  | [optional] 
**TradeGroupId** | Pointer to **int64** |  | [optional] 

## Methods

### NewFuturesAccountConfigurationResponse

`func NewFuturesAccountConfigurationResponse() *FuturesAccountConfigurationResponse`

NewFuturesAccountConfigurationResponse instantiates a new FuturesAccountConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuturesAccountConfigurationResponseWithDefaults

`func NewFuturesAccountConfigurationResponseWithDefaults() *FuturesAccountConfigurationResponse`

NewFuturesAccountConfigurationResponseWithDefaults instantiates a new FuturesAccountConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTier

`func (o *FuturesAccountConfigurationResponse) GetFeeTier() int64`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *FuturesAccountConfigurationResponse) GetFeeTierOk() (*int64, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *FuturesAccountConfigurationResponse) SetFeeTier(v int64)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *FuturesAccountConfigurationResponse) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetCanTrade

`func (o *FuturesAccountConfigurationResponse) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *FuturesAccountConfigurationResponse) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *FuturesAccountConfigurationResponse) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *FuturesAccountConfigurationResponse) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanDeposit

`func (o *FuturesAccountConfigurationResponse) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *FuturesAccountConfigurationResponse) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *FuturesAccountConfigurationResponse) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *FuturesAccountConfigurationResponse) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *FuturesAccountConfigurationResponse) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *FuturesAccountConfigurationResponse) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *FuturesAccountConfigurationResponse) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *FuturesAccountConfigurationResponse) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetDualSidePosition

`func (o *FuturesAccountConfigurationResponse) GetDualSidePosition() bool`

GetDualSidePosition returns the DualSidePosition field if non-nil, zero value otherwise.

### GetDualSidePositionOk

`func (o *FuturesAccountConfigurationResponse) GetDualSidePositionOk() (*bool, bool)`

GetDualSidePositionOk returns a tuple with the DualSidePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualSidePosition

`func (o *FuturesAccountConfigurationResponse) SetDualSidePosition(v bool)`

SetDualSidePosition sets DualSidePosition field to given value.

### HasDualSidePosition

`func (o *FuturesAccountConfigurationResponse) HasDualSidePosition() bool`

HasDualSidePosition returns a boolean if a field has been set.

### GetUpdateTime

`func (o *FuturesAccountConfigurationResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *FuturesAccountConfigurationResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *FuturesAccountConfigurationResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *FuturesAccountConfigurationResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetMultiAssetsMargin

`func (o *FuturesAccountConfigurationResponse) GetMultiAssetsMargin() bool`

GetMultiAssetsMargin returns the MultiAssetsMargin field if non-nil, zero value otherwise.

### GetMultiAssetsMarginOk

`func (o *FuturesAccountConfigurationResponse) GetMultiAssetsMarginOk() (*bool, bool)`

GetMultiAssetsMarginOk returns a tuple with the MultiAssetsMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAssetsMargin

`func (o *FuturesAccountConfigurationResponse) SetMultiAssetsMargin(v bool)`

SetMultiAssetsMargin sets MultiAssetsMargin field to given value.

### HasMultiAssetsMargin

`func (o *FuturesAccountConfigurationResponse) HasMultiAssetsMargin() bool`

HasMultiAssetsMargin returns a boolean if a field has been set.

### GetTradeGroupId

`func (o *FuturesAccountConfigurationResponse) GetTradeGroupId() int64`

GetTradeGroupId returns the TradeGroupId field if non-nil, zero value otherwise.

### GetTradeGroupIdOk

`func (o *FuturesAccountConfigurationResponse) GetTradeGroupIdOk() (*int64, bool)`

GetTradeGroupIdOk returns a tuple with the TradeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeGroupId

`func (o *FuturesAccountConfigurationResponse) SetTradeGroupId(v int64)`

SetTradeGroupId sets TradeGroupId field to given value.

### HasTradeGroupId

`func (o *FuturesAccountConfigurationResponse) HasTradeGroupId() bool`

HasTradeGroupId returns a boolean if a field has been set.


[[Back to README]](../README.md)


