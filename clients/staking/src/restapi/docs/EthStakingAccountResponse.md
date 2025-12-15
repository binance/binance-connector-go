# EthStakingAccountResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**HoldingInETH** | Pointer to **string** |  | [optional] 
**Holdings** | Pointer to [**EthStakingAccountResponseHoldings**](EthStakingAccountResponseHoldings.md) |  | [optional] 
**ThirtyDaysProfitInETH** | Pointer to **string** |  | [optional] 
**Profit** | Pointer to [**EthStakingAccountResponseProfit**](EthStakingAccountResponseProfit.md) |  | [optional] 

## Methods

### NewEthStakingAccountResponse

`func NewEthStakingAccountResponse() *EthStakingAccountResponse`

NewEthStakingAccountResponse instantiates a new EthStakingAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthStakingAccountResponseWithDefaults

`func NewEthStakingAccountResponseWithDefaults() *EthStakingAccountResponse`

NewEthStakingAccountResponseWithDefaults instantiates a new EthStakingAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoldingInETH

`func (o *EthStakingAccountResponse) GetHoldingInETH() string`

GetHoldingInETH returns the HoldingInETH field if non-nil, zero value otherwise.

### GetHoldingInETHOk

`func (o *EthStakingAccountResponse) GetHoldingInETHOk() (*string, bool)`

GetHoldingInETHOk returns a tuple with the HoldingInETH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingInETH

`func (o *EthStakingAccountResponse) SetHoldingInETH(v string)`

SetHoldingInETH sets HoldingInETH field to given value.

### HasHoldingInETH

`func (o *EthStakingAccountResponse) HasHoldingInETH() bool`

HasHoldingInETH returns a boolean if a field has been set.

### GetHoldings

`func (o *EthStakingAccountResponse) GetHoldings() EthStakingAccountResponseHoldings`

GetHoldings returns the Holdings field if non-nil, zero value otherwise.

### GetHoldingsOk

`func (o *EthStakingAccountResponse) GetHoldingsOk() (*EthStakingAccountResponseHoldings, bool)`

GetHoldingsOk returns a tuple with the Holdings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldings

`func (o *EthStakingAccountResponse) SetHoldings(v EthStakingAccountResponseHoldings)`

SetHoldings sets Holdings field to given value.

### HasHoldings

`func (o *EthStakingAccountResponse) HasHoldings() bool`

HasHoldings returns a boolean if a field has been set.

### GetThirtyDaysProfitInETH

`func (o *EthStakingAccountResponse) GetThirtyDaysProfitInETH() string`

GetThirtyDaysProfitInETH returns the ThirtyDaysProfitInETH field if non-nil, zero value otherwise.

### GetThirtyDaysProfitInETHOk

`func (o *EthStakingAccountResponse) GetThirtyDaysProfitInETHOk() (*string, bool)`

GetThirtyDaysProfitInETHOk returns a tuple with the ThirtyDaysProfitInETH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirtyDaysProfitInETH

`func (o *EthStakingAccountResponse) SetThirtyDaysProfitInETH(v string)`

SetThirtyDaysProfitInETH sets ThirtyDaysProfitInETH field to given value.

### HasThirtyDaysProfitInETH

`func (o *EthStakingAccountResponse) HasThirtyDaysProfitInETH() bool`

HasThirtyDaysProfitInETH returns a boolean if a field has been set.

### GetProfit

`func (o *EthStakingAccountResponse) GetProfit() EthStakingAccountResponseProfit`

GetProfit returns the Profit field if non-nil, zero value otherwise.

### GetProfitOk

`func (o *EthStakingAccountResponse) GetProfitOk() (*EthStakingAccountResponseProfit, bool)`

GetProfitOk returns a tuple with the Profit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfit

`func (o *EthStakingAccountResponse) SetProfit(v EthStakingAccountResponseProfit)`

SetProfit sets Profit field to given value.

### HasProfit

`func (o *EthStakingAccountResponse) HasProfit() bool`

HasProfit returns a boolean if a field has been set.


[[Back to README]](../README.md)


