# GetApiKeyPermissionResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**IpRestrict** | Pointer to **bool** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**EnableReading** | Pointer to **bool** |  | [optional] 
**EnableWithdrawals** | Pointer to **bool** |  | [optional] 
**EnableInternalTransfer** | Pointer to **bool** |  | [optional] 
**EnableMargin** | Pointer to **bool** |  | [optional] 
**EnableFutures** | Pointer to **bool** |  | [optional] 
**PermitsUniversalTransfer** | Pointer to **bool** |  | [optional] 
**EnableVanillaOptions** | Pointer to **bool** |  | [optional] 
**EnableFixApiTrade** | Pointer to **bool** |  | [optional] 
**EnableFixReadOnly** | Pointer to **bool** |  | [optional] 
**EnableSpotAndMarginTrading** | Pointer to **bool** |  | [optional] 
**EnablePortfolioMarginTrading** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetApiKeyPermissionResponse

`func NewGetApiKeyPermissionResponse() *GetApiKeyPermissionResponse`

NewGetApiKeyPermissionResponse instantiates a new GetApiKeyPermissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiKeyPermissionResponseWithDefaults

`func NewGetApiKeyPermissionResponseWithDefaults() *GetApiKeyPermissionResponse`

NewGetApiKeyPermissionResponseWithDefaults instantiates a new GetApiKeyPermissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpRestrict

`func (o *GetApiKeyPermissionResponse) GetIpRestrict() bool`

GetIpRestrict returns the IpRestrict field if non-nil, zero value otherwise.

### GetIpRestrictOk

`func (o *GetApiKeyPermissionResponse) GetIpRestrictOk() (*bool, bool)`

GetIpRestrictOk returns a tuple with the IpRestrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRestrict

`func (o *GetApiKeyPermissionResponse) SetIpRestrict(v bool)`

SetIpRestrict sets IpRestrict field to given value.

### HasIpRestrict

`func (o *GetApiKeyPermissionResponse) HasIpRestrict() bool`

HasIpRestrict returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetApiKeyPermissionResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetApiKeyPermissionResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetApiKeyPermissionResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetApiKeyPermissionResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetEnableReading

`func (o *GetApiKeyPermissionResponse) GetEnableReading() bool`

GetEnableReading returns the EnableReading field if non-nil, zero value otherwise.

### GetEnableReadingOk

`func (o *GetApiKeyPermissionResponse) GetEnableReadingOk() (*bool, bool)`

GetEnableReadingOk returns a tuple with the EnableReading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReading

`func (o *GetApiKeyPermissionResponse) SetEnableReading(v bool)`

SetEnableReading sets EnableReading field to given value.

### HasEnableReading

`func (o *GetApiKeyPermissionResponse) HasEnableReading() bool`

HasEnableReading returns a boolean if a field has been set.

### GetEnableWithdrawals

`func (o *GetApiKeyPermissionResponse) GetEnableWithdrawals() bool`

GetEnableWithdrawals returns the EnableWithdrawals field if non-nil, zero value otherwise.

### GetEnableWithdrawalsOk

`func (o *GetApiKeyPermissionResponse) GetEnableWithdrawalsOk() (*bool, bool)`

GetEnableWithdrawalsOk returns a tuple with the EnableWithdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWithdrawals

`func (o *GetApiKeyPermissionResponse) SetEnableWithdrawals(v bool)`

SetEnableWithdrawals sets EnableWithdrawals field to given value.

### HasEnableWithdrawals

`func (o *GetApiKeyPermissionResponse) HasEnableWithdrawals() bool`

HasEnableWithdrawals returns a boolean if a field has been set.

### GetEnableInternalTransfer

`func (o *GetApiKeyPermissionResponse) GetEnableInternalTransfer() bool`

GetEnableInternalTransfer returns the EnableInternalTransfer field if non-nil, zero value otherwise.

### GetEnableInternalTransferOk

`func (o *GetApiKeyPermissionResponse) GetEnableInternalTransferOk() (*bool, bool)`

GetEnableInternalTransferOk returns a tuple with the EnableInternalTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternalTransfer

`func (o *GetApiKeyPermissionResponse) SetEnableInternalTransfer(v bool)`

SetEnableInternalTransfer sets EnableInternalTransfer field to given value.

### HasEnableInternalTransfer

`func (o *GetApiKeyPermissionResponse) HasEnableInternalTransfer() bool`

HasEnableInternalTransfer returns a boolean if a field has been set.

### GetEnableMargin

`func (o *GetApiKeyPermissionResponse) GetEnableMargin() bool`

GetEnableMargin returns the EnableMargin field if non-nil, zero value otherwise.

### GetEnableMarginOk

`func (o *GetApiKeyPermissionResponse) GetEnableMarginOk() (*bool, bool)`

GetEnableMarginOk returns a tuple with the EnableMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMargin

`func (o *GetApiKeyPermissionResponse) SetEnableMargin(v bool)`

SetEnableMargin sets EnableMargin field to given value.

### HasEnableMargin

`func (o *GetApiKeyPermissionResponse) HasEnableMargin() bool`

HasEnableMargin returns a boolean if a field has been set.

### GetEnableFutures

`func (o *GetApiKeyPermissionResponse) GetEnableFutures() bool`

GetEnableFutures returns the EnableFutures field if non-nil, zero value otherwise.

### GetEnableFuturesOk

`func (o *GetApiKeyPermissionResponse) GetEnableFuturesOk() (*bool, bool)`

GetEnableFuturesOk returns a tuple with the EnableFutures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFutures

`func (o *GetApiKeyPermissionResponse) SetEnableFutures(v bool)`

SetEnableFutures sets EnableFutures field to given value.

### HasEnableFutures

`func (o *GetApiKeyPermissionResponse) HasEnableFutures() bool`

HasEnableFutures returns a boolean if a field has been set.

### GetPermitsUniversalTransfer

`func (o *GetApiKeyPermissionResponse) GetPermitsUniversalTransfer() bool`

GetPermitsUniversalTransfer returns the PermitsUniversalTransfer field if non-nil, zero value otherwise.

### GetPermitsUniversalTransferOk

`func (o *GetApiKeyPermissionResponse) GetPermitsUniversalTransferOk() (*bool, bool)`

GetPermitsUniversalTransferOk returns a tuple with the PermitsUniversalTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitsUniversalTransfer

`func (o *GetApiKeyPermissionResponse) SetPermitsUniversalTransfer(v bool)`

SetPermitsUniversalTransfer sets PermitsUniversalTransfer field to given value.

### HasPermitsUniversalTransfer

`func (o *GetApiKeyPermissionResponse) HasPermitsUniversalTransfer() bool`

HasPermitsUniversalTransfer returns a boolean if a field has been set.

### GetEnableVanillaOptions

`func (o *GetApiKeyPermissionResponse) GetEnableVanillaOptions() bool`

GetEnableVanillaOptions returns the EnableVanillaOptions field if non-nil, zero value otherwise.

### GetEnableVanillaOptionsOk

`func (o *GetApiKeyPermissionResponse) GetEnableVanillaOptionsOk() (*bool, bool)`

GetEnableVanillaOptionsOk returns a tuple with the EnableVanillaOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVanillaOptions

`func (o *GetApiKeyPermissionResponse) SetEnableVanillaOptions(v bool)`

SetEnableVanillaOptions sets EnableVanillaOptions field to given value.

### HasEnableVanillaOptions

`func (o *GetApiKeyPermissionResponse) HasEnableVanillaOptions() bool`

HasEnableVanillaOptions returns a boolean if a field has been set.

### GetEnableFixApiTrade

`func (o *GetApiKeyPermissionResponse) GetEnableFixApiTrade() bool`

GetEnableFixApiTrade returns the EnableFixApiTrade field if non-nil, zero value otherwise.

### GetEnableFixApiTradeOk

`func (o *GetApiKeyPermissionResponse) GetEnableFixApiTradeOk() (*bool, bool)`

GetEnableFixApiTradeOk returns a tuple with the EnableFixApiTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFixApiTrade

`func (o *GetApiKeyPermissionResponse) SetEnableFixApiTrade(v bool)`

SetEnableFixApiTrade sets EnableFixApiTrade field to given value.

### HasEnableFixApiTrade

`func (o *GetApiKeyPermissionResponse) HasEnableFixApiTrade() bool`

HasEnableFixApiTrade returns a boolean if a field has been set.

### GetEnableFixReadOnly

`func (o *GetApiKeyPermissionResponse) GetEnableFixReadOnly() bool`

GetEnableFixReadOnly returns the EnableFixReadOnly field if non-nil, zero value otherwise.

### GetEnableFixReadOnlyOk

`func (o *GetApiKeyPermissionResponse) GetEnableFixReadOnlyOk() (*bool, bool)`

GetEnableFixReadOnlyOk returns a tuple with the EnableFixReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFixReadOnly

`func (o *GetApiKeyPermissionResponse) SetEnableFixReadOnly(v bool)`

SetEnableFixReadOnly sets EnableFixReadOnly field to given value.

### HasEnableFixReadOnly

`func (o *GetApiKeyPermissionResponse) HasEnableFixReadOnly() bool`

HasEnableFixReadOnly returns a boolean if a field has been set.

### GetEnableSpotAndMarginTrading

`func (o *GetApiKeyPermissionResponse) GetEnableSpotAndMarginTrading() bool`

GetEnableSpotAndMarginTrading returns the EnableSpotAndMarginTrading field if non-nil, zero value otherwise.

### GetEnableSpotAndMarginTradingOk

`func (o *GetApiKeyPermissionResponse) GetEnableSpotAndMarginTradingOk() (*bool, bool)`

GetEnableSpotAndMarginTradingOk returns a tuple with the EnableSpotAndMarginTrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpotAndMarginTrading

`func (o *GetApiKeyPermissionResponse) SetEnableSpotAndMarginTrading(v bool)`

SetEnableSpotAndMarginTrading sets EnableSpotAndMarginTrading field to given value.

### HasEnableSpotAndMarginTrading

`func (o *GetApiKeyPermissionResponse) HasEnableSpotAndMarginTrading() bool`

HasEnableSpotAndMarginTrading returns a boolean if a field has been set.

### GetEnablePortfolioMarginTrading

`func (o *GetApiKeyPermissionResponse) GetEnablePortfolioMarginTrading() bool`

GetEnablePortfolioMarginTrading returns the EnablePortfolioMarginTrading field if non-nil, zero value otherwise.

### GetEnablePortfolioMarginTradingOk

`func (o *GetApiKeyPermissionResponse) GetEnablePortfolioMarginTradingOk() (*bool, bool)`

GetEnablePortfolioMarginTradingOk returns a tuple with the EnablePortfolioMarginTrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePortfolioMarginTrading

`func (o *GetApiKeyPermissionResponse) SetEnablePortfolioMarginTrading(v bool)`

SetEnablePortfolioMarginTrading sets EnablePortfolioMarginTrading field to given value.

### HasEnablePortfolioMarginTrading

`func (o *GetApiKeyPermissionResponse) HasEnablePortfolioMarginTrading() bool`

HasEnablePortfolioMarginTrading returns a boolean if a field has been set.


[[Back to README]](../README.md)


