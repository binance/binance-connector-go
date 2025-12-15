# UmFuturesAccountConfigurationResponse

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

### NewUmFuturesAccountConfigurationResponse

`func NewUmFuturesAccountConfigurationResponse() *UmFuturesAccountConfigurationResponse`

NewUmFuturesAccountConfigurationResponse instantiates a new UmFuturesAccountConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmFuturesAccountConfigurationResponseWithDefaults

`func NewUmFuturesAccountConfigurationResponseWithDefaults() *UmFuturesAccountConfigurationResponse`

NewUmFuturesAccountConfigurationResponseWithDefaults instantiates a new UmFuturesAccountConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTier

`func (o *UmFuturesAccountConfigurationResponse) GetFeeTier() int64`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *UmFuturesAccountConfigurationResponse) GetFeeTierOk() (*int64, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *UmFuturesAccountConfigurationResponse) SetFeeTier(v int64)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *UmFuturesAccountConfigurationResponse) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetCanTrade

`func (o *UmFuturesAccountConfigurationResponse) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *UmFuturesAccountConfigurationResponse) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *UmFuturesAccountConfigurationResponse) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *UmFuturesAccountConfigurationResponse) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanDeposit

`func (o *UmFuturesAccountConfigurationResponse) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *UmFuturesAccountConfigurationResponse) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *UmFuturesAccountConfigurationResponse) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *UmFuturesAccountConfigurationResponse) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *UmFuturesAccountConfigurationResponse) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *UmFuturesAccountConfigurationResponse) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *UmFuturesAccountConfigurationResponse) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *UmFuturesAccountConfigurationResponse) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetDualSidePosition

`func (o *UmFuturesAccountConfigurationResponse) GetDualSidePosition() bool`

GetDualSidePosition returns the DualSidePosition field if non-nil, zero value otherwise.

### GetDualSidePositionOk

`func (o *UmFuturesAccountConfigurationResponse) GetDualSidePositionOk() (*bool, bool)`

GetDualSidePositionOk returns a tuple with the DualSidePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualSidePosition

`func (o *UmFuturesAccountConfigurationResponse) SetDualSidePosition(v bool)`

SetDualSidePosition sets DualSidePosition field to given value.

### HasDualSidePosition

`func (o *UmFuturesAccountConfigurationResponse) HasDualSidePosition() bool`

HasDualSidePosition returns a boolean if a field has been set.

### GetUpdateTime

`func (o *UmFuturesAccountConfigurationResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *UmFuturesAccountConfigurationResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *UmFuturesAccountConfigurationResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *UmFuturesAccountConfigurationResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetMultiAssetsMargin

`func (o *UmFuturesAccountConfigurationResponse) GetMultiAssetsMargin() bool`

GetMultiAssetsMargin returns the MultiAssetsMargin field if non-nil, zero value otherwise.

### GetMultiAssetsMarginOk

`func (o *UmFuturesAccountConfigurationResponse) GetMultiAssetsMarginOk() (*bool, bool)`

GetMultiAssetsMarginOk returns a tuple with the MultiAssetsMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAssetsMargin

`func (o *UmFuturesAccountConfigurationResponse) SetMultiAssetsMargin(v bool)`

SetMultiAssetsMargin sets MultiAssetsMargin field to given value.

### HasMultiAssetsMargin

`func (o *UmFuturesAccountConfigurationResponse) HasMultiAssetsMargin() bool`

HasMultiAssetsMargin returns a boolean if a field has been set.

### GetTradeGroupId

`func (o *UmFuturesAccountConfigurationResponse) GetTradeGroupId() int64`

GetTradeGroupId returns the TradeGroupId field if non-nil, zero value otherwise.

### GetTradeGroupIdOk

`func (o *UmFuturesAccountConfigurationResponse) GetTradeGroupIdOk() (*int64, bool)`

GetTradeGroupIdOk returns a tuple with the TradeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeGroupId

`func (o *UmFuturesAccountConfigurationResponse) SetTradeGroupId(v int64)`

SetTradeGroupId sets TradeGroupId field to given value.

### HasTradeGroupId

`func (o *UmFuturesAccountConfigurationResponse) HasTradeGroupId() bool`

HasTradeGroupId returns a boolean if a field has been set.


[[Back to README]](../README.md)


