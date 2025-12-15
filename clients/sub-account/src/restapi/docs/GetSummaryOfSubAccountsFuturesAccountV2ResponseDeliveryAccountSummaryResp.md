# GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TotalMarginBalanceOfBTC** | Pointer to **string** |  | [optional] 
**TotalUnrealizedProfitOfBTC** | Pointer to **string** |  | [optional] 
**TotalWalletBalanceOfBTC** | Pointer to **string** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
**SubAccountList** | Pointer to [**[]GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner**](GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner.md) |  | [optional] 

## Methods

### NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp

`func NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp() *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp`

NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp instantiates a new GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespWithDefaults

`func NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespWithDefaults() *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp`

NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespWithDefaults instantiates a new GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalMarginBalanceOfBTC

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetTotalMarginBalanceOfBTC() string`

GetTotalMarginBalanceOfBTC returns the TotalMarginBalanceOfBTC field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOfBTCOk

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetTotalMarginBalanceOfBTCOk() (*string, bool)`

GetTotalMarginBalanceOfBTCOk returns a tuple with the TotalMarginBalanceOfBTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalanceOfBTC

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) SetTotalMarginBalanceOfBTC(v string)`

SetTotalMarginBalanceOfBTC sets TotalMarginBalanceOfBTC field to given value.

### HasTotalMarginBalanceOfBTC

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) HasTotalMarginBalanceOfBTC() bool`

HasTotalMarginBalanceOfBTC returns a boolean if a field has been set.

### GetTotalUnrealizedProfitOfBTC

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetTotalUnrealizedProfitOfBTC() string`

GetTotalUnrealizedProfitOfBTC returns the TotalUnrealizedProfitOfBTC field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOfBTCOk

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetTotalUnrealizedProfitOfBTCOk() (*string, bool)`

GetTotalUnrealizedProfitOfBTCOk returns a tuple with the TotalUnrealizedProfitOfBTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfitOfBTC

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) SetTotalUnrealizedProfitOfBTC(v string)`

SetTotalUnrealizedProfitOfBTC sets TotalUnrealizedProfitOfBTC field to given value.

### HasTotalUnrealizedProfitOfBTC

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) HasTotalUnrealizedProfitOfBTC() bool`

HasTotalUnrealizedProfitOfBTC returns a boolean if a field has been set.

### GetTotalWalletBalanceOfBTC

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetTotalWalletBalanceOfBTC() string`

GetTotalWalletBalanceOfBTC returns the TotalWalletBalanceOfBTC field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOfBTCOk

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetTotalWalletBalanceOfBTCOk() (*string, bool)`

GetTotalWalletBalanceOfBTCOk returns a tuple with the TotalWalletBalanceOfBTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalanceOfBTC

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) SetTotalWalletBalanceOfBTC(v string)`

SetTotalWalletBalanceOfBTC sets TotalWalletBalanceOfBTC field to given value.

### HasTotalWalletBalanceOfBTC

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) HasTotalWalletBalanceOfBTC() bool`

HasTotalWalletBalanceOfBTC returns a boolean if a field has been set.

### GetAsset

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetSubAccountList

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetSubAccountList() []GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner`

GetSubAccountList returns the SubAccountList field if non-nil, zero value otherwise.

### GetSubAccountListOk

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetSubAccountListOk() (*[]GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner, bool)`

GetSubAccountListOk returns a tuple with the SubAccountList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountList

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) SetSubAccountList(v []GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner)`

SetSubAccountList sets SubAccountList field to given value.

### HasSubAccountList

`func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) HasSubAccountList() bool`

HasSubAccountList returns a boolean if a field has been set.


[[Back to README]](../README.md)


