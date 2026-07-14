# ModifySubAccountApiKeyPermissionResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ApiName** | Pointer to **string** |  | [optional] 
**Apikey** | Pointer to **string** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanMarginLoanRepay** | Pointer to **bool** |  | [optional] 
**CanFuturesTrade** | Pointer to **bool** |  | [optional] 
**CanUniversalTransfer** | Pointer to **bool** |  | [optional] 
**CanVanillaOptions** | Pointer to **bool** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewModifySubAccountApiKeyPermissionResponse

`func NewModifySubAccountApiKeyPermissionResponse() *ModifySubAccountApiKeyPermissionResponse`

NewModifySubAccountApiKeyPermissionResponse instantiates a new ModifySubAccountApiKeyPermissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifySubAccountApiKeyPermissionResponseWithDefaults

`func NewModifySubAccountApiKeyPermissionResponseWithDefaults() *ModifySubAccountApiKeyPermissionResponse`

NewModifySubAccountApiKeyPermissionResponseWithDefaults instantiates a new ModifySubAccountApiKeyPermissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiName

`func (o *ModifySubAccountApiKeyPermissionResponse) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *ModifySubAccountApiKeyPermissionResponse) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *ModifySubAccountApiKeyPermissionResponse) SetApiName(v string)`

SetApiName sets ApiName field to given value.

### HasApiName

`func (o *ModifySubAccountApiKeyPermissionResponse) HasApiName() bool`

HasApiName returns a boolean if a field has been set.

### GetApikey

`func (o *ModifySubAccountApiKeyPermissionResponse) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *ModifySubAccountApiKeyPermissionResponse) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *ModifySubAccountApiKeyPermissionResponse) SetApikey(v string)`

SetApikey sets Apikey field to given value.

### HasApikey

`func (o *ModifySubAccountApiKeyPermissionResponse) HasApikey() bool`

HasApikey returns a boolean if a field has been set.

### GetCanTrade

`func (o *ModifySubAccountApiKeyPermissionResponse) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *ModifySubAccountApiKeyPermissionResponse) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *ModifySubAccountApiKeyPermissionResponse) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *ModifySubAccountApiKeyPermissionResponse) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanMarginLoanRepay

`func (o *ModifySubAccountApiKeyPermissionResponse) GetCanMarginLoanRepay() bool`

GetCanMarginLoanRepay returns the CanMarginLoanRepay field if non-nil, zero value otherwise.

### GetCanMarginLoanRepayOk

`func (o *ModifySubAccountApiKeyPermissionResponse) GetCanMarginLoanRepayOk() (*bool, bool)`

GetCanMarginLoanRepayOk returns a tuple with the CanMarginLoanRepay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMarginLoanRepay

`func (o *ModifySubAccountApiKeyPermissionResponse) SetCanMarginLoanRepay(v bool)`

SetCanMarginLoanRepay sets CanMarginLoanRepay field to given value.

### HasCanMarginLoanRepay

`func (o *ModifySubAccountApiKeyPermissionResponse) HasCanMarginLoanRepay() bool`

HasCanMarginLoanRepay returns a boolean if a field has been set.

### GetCanFuturesTrade

`func (o *ModifySubAccountApiKeyPermissionResponse) GetCanFuturesTrade() bool`

GetCanFuturesTrade returns the CanFuturesTrade field if non-nil, zero value otherwise.

### GetCanFuturesTradeOk

`func (o *ModifySubAccountApiKeyPermissionResponse) GetCanFuturesTradeOk() (*bool, bool)`

GetCanFuturesTradeOk returns a tuple with the CanFuturesTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanFuturesTrade

`func (o *ModifySubAccountApiKeyPermissionResponse) SetCanFuturesTrade(v bool)`

SetCanFuturesTrade sets CanFuturesTrade field to given value.

### HasCanFuturesTrade

`func (o *ModifySubAccountApiKeyPermissionResponse) HasCanFuturesTrade() bool`

HasCanFuturesTrade returns a boolean if a field has been set.

### GetCanUniversalTransfer

`func (o *ModifySubAccountApiKeyPermissionResponse) GetCanUniversalTransfer() bool`

GetCanUniversalTransfer returns the CanUniversalTransfer field if non-nil, zero value otherwise.

### GetCanUniversalTransferOk

`func (o *ModifySubAccountApiKeyPermissionResponse) GetCanUniversalTransferOk() (*bool, bool)`

GetCanUniversalTransferOk returns a tuple with the CanUniversalTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUniversalTransfer

`func (o *ModifySubAccountApiKeyPermissionResponse) SetCanUniversalTransfer(v bool)`

SetCanUniversalTransfer sets CanUniversalTransfer field to given value.

### HasCanUniversalTransfer

`func (o *ModifySubAccountApiKeyPermissionResponse) HasCanUniversalTransfer() bool`

HasCanUniversalTransfer returns a boolean if a field has been set.

### GetCanVanillaOptions

`func (o *ModifySubAccountApiKeyPermissionResponse) GetCanVanillaOptions() bool`

GetCanVanillaOptions returns the CanVanillaOptions field if non-nil, zero value otherwise.

### GetCanVanillaOptionsOk

`func (o *ModifySubAccountApiKeyPermissionResponse) GetCanVanillaOptionsOk() (*bool, bool)`

GetCanVanillaOptionsOk returns a tuple with the CanVanillaOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanVanillaOptions

`func (o *ModifySubAccountApiKeyPermissionResponse) SetCanVanillaOptions(v bool)`

SetCanVanillaOptions sets CanVanillaOptions field to given value.

### HasCanVanillaOptions

`func (o *ModifySubAccountApiKeyPermissionResponse) HasCanVanillaOptions() bool`

HasCanVanillaOptions returns a boolean if a field has been set.

### GetTimestamp

`func (o *ModifySubAccountApiKeyPermissionResponse) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ModifySubAccountApiKeyPermissionResponse) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ModifySubAccountApiKeyPermissionResponse) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ModifySubAccountApiKeyPermissionResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to README]](../README.md)


