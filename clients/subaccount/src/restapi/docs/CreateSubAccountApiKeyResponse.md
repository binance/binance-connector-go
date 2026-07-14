# CreateSubAccountApiKeyResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ApiName** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** | Secret Key. Returned only once on creation, please keep it safe. | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanMarginLoanRepay** | Pointer to **bool** |  | [optional] 
**CanFuturesTrade** | Pointer to **bool** |  | [optional] 
**CanUniversalTransfer** | Pointer to **bool** |  | [optional] 
**CanVanillaOptions** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**IpList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateSubAccountApiKeyResponse

`func NewCreateSubAccountApiKeyResponse() *CreateSubAccountApiKeyResponse`

NewCreateSubAccountApiKeyResponse instantiates a new CreateSubAccountApiKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubAccountApiKeyResponseWithDefaults

`func NewCreateSubAccountApiKeyResponseWithDefaults() *CreateSubAccountApiKeyResponse`

NewCreateSubAccountApiKeyResponseWithDefaults instantiates a new CreateSubAccountApiKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiName

`func (o *CreateSubAccountApiKeyResponse) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *CreateSubAccountApiKeyResponse) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *CreateSubAccountApiKeyResponse) SetApiName(v string)`

SetApiName sets ApiName field to given value.

### HasApiName

`func (o *CreateSubAccountApiKeyResponse) HasApiName() bool`

HasApiName returns a boolean if a field has been set.

### GetApiKey

`func (o *CreateSubAccountApiKeyResponse) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateSubAccountApiKeyResponse) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateSubAccountApiKeyResponse) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CreateSubAccountApiKeyResponse) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *CreateSubAccountApiKeyResponse) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *CreateSubAccountApiKeyResponse) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *CreateSubAccountApiKeyResponse) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *CreateSubAccountApiKeyResponse) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetCanTrade

`func (o *CreateSubAccountApiKeyResponse) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *CreateSubAccountApiKeyResponse) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *CreateSubAccountApiKeyResponse) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *CreateSubAccountApiKeyResponse) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanMarginLoanRepay

`func (o *CreateSubAccountApiKeyResponse) GetCanMarginLoanRepay() bool`

GetCanMarginLoanRepay returns the CanMarginLoanRepay field if non-nil, zero value otherwise.

### GetCanMarginLoanRepayOk

`func (o *CreateSubAccountApiKeyResponse) GetCanMarginLoanRepayOk() (*bool, bool)`

GetCanMarginLoanRepayOk returns a tuple with the CanMarginLoanRepay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMarginLoanRepay

`func (o *CreateSubAccountApiKeyResponse) SetCanMarginLoanRepay(v bool)`

SetCanMarginLoanRepay sets CanMarginLoanRepay field to given value.

### HasCanMarginLoanRepay

`func (o *CreateSubAccountApiKeyResponse) HasCanMarginLoanRepay() bool`

HasCanMarginLoanRepay returns a boolean if a field has been set.

### GetCanFuturesTrade

`func (o *CreateSubAccountApiKeyResponse) GetCanFuturesTrade() bool`

GetCanFuturesTrade returns the CanFuturesTrade field if non-nil, zero value otherwise.

### GetCanFuturesTradeOk

`func (o *CreateSubAccountApiKeyResponse) GetCanFuturesTradeOk() (*bool, bool)`

GetCanFuturesTradeOk returns a tuple with the CanFuturesTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanFuturesTrade

`func (o *CreateSubAccountApiKeyResponse) SetCanFuturesTrade(v bool)`

SetCanFuturesTrade sets CanFuturesTrade field to given value.

### HasCanFuturesTrade

`func (o *CreateSubAccountApiKeyResponse) HasCanFuturesTrade() bool`

HasCanFuturesTrade returns a boolean if a field has been set.

### GetCanUniversalTransfer

`func (o *CreateSubAccountApiKeyResponse) GetCanUniversalTransfer() bool`

GetCanUniversalTransfer returns the CanUniversalTransfer field if non-nil, zero value otherwise.

### GetCanUniversalTransferOk

`func (o *CreateSubAccountApiKeyResponse) GetCanUniversalTransferOk() (*bool, bool)`

GetCanUniversalTransferOk returns a tuple with the CanUniversalTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUniversalTransfer

`func (o *CreateSubAccountApiKeyResponse) SetCanUniversalTransfer(v bool)`

SetCanUniversalTransfer sets CanUniversalTransfer field to given value.

### HasCanUniversalTransfer

`func (o *CreateSubAccountApiKeyResponse) HasCanUniversalTransfer() bool`

HasCanUniversalTransfer returns a boolean if a field has been set.

### GetCanVanillaOptions

`func (o *CreateSubAccountApiKeyResponse) GetCanVanillaOptions() bool`

GetCanVanillaOptions returns the CanVanillaOptions field if non-nil, zero value otherwise.

### GetCanVanillaOptionsOk

`func (o *CreateSubAccountApiKeyResponse) GetCanVanillaOptionsOk() (*bool, bool)`

GetCanVanillaOptionsOk returns a tuple with the CanVanillaOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanVanillaOptions

`func (o *CreateSubAccountApiKeyResponse) SetCanVanillaOptions(v bool)`

SetCanVanillaOptions sets CanVanillaOptions field to given value.

### HasCanVanillaOptions

`func (o *CreateSubAccountApiKeyResponse) HasCanVanillaOptions() bool`

HasCanVanillaOptions returns a boolean if a field has been set.

### GetStatus

`func (o *CreateSubAccountApiKeyResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateSubAccountApiKeyResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateSubAccountApiKeyResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateSubAccountApiKeyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIpList

`func (o *CreateSubAccountApiKeyResponse) GetIpList() []string`

GetIpList returns the IpList field if non-nil, zero value otherwise.

### GetIpListOk

`func (o *CreateSubAccountApiKeyResponse) GetIpListOk() (*[]string, bool)`

GetIpListOk returns a tuple with the IpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpList

`func (o *CreateSubAccountApiKeyResponse) SetIpList(v []string)`

SetIpList sets IpList field to given value.

### HasIpList

`func (o *CreateSubAccountApiKeyResponse) HasIpList() bool`

HasIpList returns a boolean if a field has been set.


[[Back to README]](../README.md)


