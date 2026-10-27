# CreateSubAccountApiKeyV2Response

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
**CanFixApiTrading** | Pointer to **bool** |  | [optional] 
**CanFixApiReading** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**IpList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateSubAccountApiKeyV2Response

`func NewCreateSubAccountApiKeyV2Response() *CreateSubAccountApiKeyV2Response`

NewCreateSubAccountApiKeyV2Response instantiates a new CreateSubAccountApiKeyV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubAccountApiKeyV2ResponseWithDefaults

`func NewCreateSubAccountApiKeyV2ResponseWithDefaults() *CreateSubAccountApiKeyV2Response`

NewCreateSubAccountApiKeyV2ResponseWithDefaults instantiates a new CreateSubAccountApiKeyV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiName

`func (o *CreateSubAccountApiKeyV2Response) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *CreateSubAccountApiKeyV2Response) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *CreateSubAccountApiKeyV2Response) SetApiName(v string)`

SetApiName sets ApiName field to given value.

### HasApiName

`func (o *CreateSubAccountApiKeyV2Response) HasApiName() bool`

HasApiName returns a boolean if a field has been set.

### GetApiKey

`func (o *CreateSubAccountApiKeyV2Response) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateSubAccountApiKeyV2Response) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateSubAccountApiKeyV2Response) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CreateSubAccountApiKeyV2Response) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *CreateSubAccountApiKeyV2Response) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *CreateSubAccountApiKeyV2Response) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *CreateSubAccountApiKeyV2Response) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *CreateSubAccountApiKeyV2Response) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetCanTrade

`func (o *CreateSubAccountApiKeyV2Response) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *CreateSubAccountApiKeyV2Response) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *CreateSubAccountApiKeyV2Response) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *CreateSubAccountApiKeyV2Response) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanMarginLoanRepay

`func (o *CreateSubAccountApiKeyV2Response) GetCanMarginLoanRepay() bool`

GetCanMarginLoanRepay returns the CanMarginLoanRepay field if non-nil, zero value otherwise.

### GetCanMarginLoanRepayOk

`func (o *CreateSubAccountApiKeyV2Response) GetCanMarginLoanRepayOk() (*bool, bool)`

GetCanMarginLoanRepayOk returns a tuple with the CanMarginLoanRepay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMarginLoanRepay

`func (o *CreateSubAccountApiKeyV2Response) SetCanMarginLoanRepay(v bool)`

SetCanMarginLoanRepay sets CanMarginLoanRepay field to given value.

### HasCanMarginLoanRepay

`func (o *CreateSubAccountApiKeyV2Response) HasCanMarginLoanRepay() bool`

HasCanMarginLoanRepay returns a boolean if a field has been set.

### GetCanFuturesTrade

`func (o *CreateSubAccountApiKeyV2Response) GetCanFuturesTrade() bool`

GetCanFuturesTrade returns the CanFuturesTrade field if non-nil, zero value otherwise.

### GetCanFuturesTradeOk

`func (o *CreateSubAccountApiKeyV2Response) GetCanFuturesTradeOk() (*bool, bool)`

GetCanFuturesTradeOk returns a tuple with the CanFuturesTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanFuturesTrade

`func (o *CreateSubAccountApiKeyV2Response) SetCanFuturesTrade(v bool)`

SetCanFuturesTrade sets CanFuturesTrade field to given value.

### HasCanFuturesTrade

`func (o *CreateSubAccountApiKeyV2Response) HasCanFuturesTrade() bool`

HasCanFuturesTrade returns a boolean if a field has been set.

### GetCanUniversalTransfer

`func (o *CreateSubAccountApiKeyV2Response) GetCanUniversalTransfer() bool`

GetCanUniversalTransfer returns the CanUniversalTransfer field if non-nil, zero value otherwise.

### GetCanUniversalTransferOk

`func (o *CreateSubAccountApiKeyV2Response) GetCanUniversalTransferOk() (*bool, bool)`

GetCanUniversalTransferOk returns a tuple with the CanUniversalTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUniversalTransfer

`func (o *CreateSubAccountApiKeyV2Response) SetCanUniversalTransfer(v bool)`

SetCanUniversalTransfer sets CanUniversalTransfer field to given value.

### HasCanUniversalTransfer

`func (o *CreateSubAccountApiKeyV2Response) HasCanUniversalTransfer() bool`

HasCanUniversalTransfer returns a boolean if a field has been set.

### GetCanVanillaOptions

`func (o *CreateSubAccountApiKeyV2Response) GetCanVanillaOptions() bool`

GetCanVanillaOptions returns the CanVanillaOptions field if non-nil, zero value otherwise.

### GetCanVanillaOptionsOk

`func (o *CreateSubAccountApiKeyV2Response) GetCanVanillaOptionsOk() (*bool, bool)`

GetCanVanillaOptionsOk returns a tuple with the CanVanillaOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanVanillaOptions

`func (o *CreateSubAccountApiKeyV2Response) SetCanVanillaOptions(v bool)`

SetCanVanillaOptions sets CanVanillaOptions field to given value.

### HasCanVanillaOptions

`func (o *CreateSubAccountApiKeyV2Response) HasCanVanillaOptions() bool`

HasCanVanillaOptions returns a boolean if a field has been set.

### GetCanFixApiTrading

`func (o *CreateSubAccountApiKeyV2Response) GetCanFixApiTrading() bool`

GetCanFixApiTrading returns the CanFixApiTrading field if non-nil, zero value otherwise.

### GetCanFixApiTradingOk

`func (o *CreateSubAccountApiKeyV2Response) GetCanFixApiTradingOk() (*bool, bool)`

GetCanFixApiTradingOk returns a tuple with the CanFixApiTrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanFixApiTrading

`func (o *CreateSubAccountApiKeyV2Response) SetCanFixApiTrading(v bool)`

SetCanFixApiTrading sets CanFixApiTrading field to given value.

### HasCanFixApiTrading

`func (o *CreateSubAccountApiKeyV2Response) HasCanFixApiTrading() bool`

HasCanFixApiTrading returns a boolean if a field has been set.

### GetCanFixApiReading

`func (o *CreateSubAccountApiKeyV2Response) GetCanFixApiReading() bool`

GetCanFixApiReading returns the CanFixApiReading field if non-nil, zero value otherwise.

### GetCanFixApiReadingOk

`func (o *CreateSubAccountApiKeyV2Response) GetCanFixApiReadingOk() (*bool, bool)`

GetCanFixApiReadingOk returns a tuple with the CanFixApiReading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanFixApiReading

`func (o *CreateSubAccountApiKeyV2Response) SetCanFixApiReading(v bool)`

SetCanFixApiReading sets CanFixApiReading field to given value.

### HasCanFixApiReading

`func (o *CreateSubAccountApiKeyV2Response) HasCanFixApiReading() bool`

HasCanFixApiReading returns a boolean if a field has been set.

### GetStatus

`func (o *CreateSubAccountApiKeyV2Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateSubAccountApiKeyV2Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateSubAccountApiKeyV2Response) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateSubAccountApiKeyV2Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIpList

`func (o *CreateSubAccountApiKeyV2Response) GetIpList() []string`

GetIpList returns the IpList field if non-nil, zero value otherwise.

### GetIpListOk

`func (o *CreateSubAccountApiKeyV2Response) GetIpListOk() (*[]string, bool)`

GetIpListOk returns a tuple with the IpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpList

`func (o *CreateSubAccountApiKeyV2Response) SetIpList(v []string)`

SetIpList sets IpList field to given value.

### HasIpList

`func (o *CreateSubAccountApiKeyV2Response) HasIpList() bool`

HasIpList returns a boolean if a field has been set.


[[Back to README]](../README.md)


