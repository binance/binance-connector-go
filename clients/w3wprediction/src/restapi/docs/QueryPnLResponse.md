# QueryPnLResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **string** |  | [optional] 
**WalletAddress** | Pointer to **string** |  | [optional] 
**Pnl** | Pointer to **map[string]interface{}** |  | [optional] 
**PnlList** | Pointer to [**[]GetPortfolioResponsePositionsInner**](GetPortfolioResponsePositionsInner.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**TotalRealizedPnl** | Pointer to **string** |  | [optional] 
**TotalUnrealizedPnl** | Pointer to **string** |  | [optional] 
**TotalPnl** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryPnLResponse

`func NewQueryPnLResponse() *QueryPnLResponse`

NewQueryPnLResponse instantiates a new QueryPnLResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPnLResponseWithDefaults

`func NewQueryPnLResponseWithDefaults() *QueryPnLResponse`

NewQueryPnLResponseWithDefaults instantiates a new QueryPnLResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *QueryPnLResponse) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *QueryPnLResponse) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *QueryPnLResponse) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *QueryPnLResponse) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetWalletAddress

`func (o *QueryPnLResponse) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *QueryPnLResponse) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *QueryPnLResponse) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *QueryPnLResponse) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.

### GetPnl

`func (o *QueryPnLResponse) GetPnl() map[string]interface{}`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *QueryPnLResponse) GetPnlOk() (*map[string]interface{}, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *QueryPnLResponse) SetPnl(v map[string]interface{})`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *QueryPnLResponse) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### SetPnlNil

`func (o *QueryPnLResponse) SetPnlNil(b bool)`

 SetPnlNil sets the value for Pnl to be an explicit nil

### UnsetPnl
`func (o *QueryPnLResponse) UnsetPnl()`

UnsetPnl ensures that no value is present for Pnl, not even an explicit nil
### GetPnlList

`func (o *QueryPnLResponse) GetPnlList() []GetPortfolioResponsePositionsInner`

GetPnlList returns the PnlList field if non-nil, zero value otherwise.

### GetPnlListOk

`func (o *QueryPnLResponse) GetPnlListOk() (*[]GetPortfolioResponsePositionsInner, bool)`

GetPnlListOk returns a tuple with the PnlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlList

`func (o *QueryPnLResponse) SetPnlList(v []GetPortfolioResponsePositionsInner)`

SetPnlList sets PnlList field to given value.

### HasPnlList

`func (o *QueryPnLResponse) HasPnlList() bool`

HasPnlList returns a boolean if a field has been set.

### GetTotalCount

`func (o *QueryPnLResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *QueryPnLResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *QueryPnLResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *QueryPnLResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalRealizedPnl

`func (o *QueryPnLResponse) GetTotalRealizedPnl() string`

GetTotalRealizedPnl returns the TotalRealizedPnl field if non-nil, zero value otherwise.

### GetTotalRealizedPnlOk

`func (o *QueryPnLResponse) GetTotalRealizedPnlOk() (*string, bool)`

GetTotalRealizedPnlOk returns a tuple with the TotalRealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRealizedPnl

`func (o *QueryPnLResponse) SetTotalRealizedPnl(v string)`

SetTotalRealizedPnl sets TotalRealizedPnl field to given value.

### HasTotalRealizedPnl

`func (o *QueryPnLResponse) HasTotalRealizedPnl() bool`

HasTotalRealizedPnl returns a boolean if a field has been set.

### GetTotalUnrealizedPnl

`func (o *QueryPnLResponse) GetTotalUnrealizedPnl() string`

GetTotalUnrealizedPnl returns the TotalUnrealizedPnl field if non-nil, zero value otherwise.

### GetTotalUnrealizedPnlOk

`func (o *QueryPnLResponse) GetTotalUnrealizedPnlOk() (*string, bool)`

GetTotalUnrealizedPnlOk returns a tuple with the TotalUnrealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedPnl

`func (o *QueryPnLResponse) SetTotalUnrealizedPnl(v string)`

SetTotalUnrealizedPnl sets TotalUnrealizedPnl field to given value.

### HasTotalUnrealizedPnl

`func (o *QueryPnLResponse) HasTotalUnrealizedPnl() bool`

HasTotalUnrealizedPnl returns a boolean if a field has been set.

### GetTotalPnl

`func (o *QueryPnLResponse) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *QueryPnLResponse) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *QueryPnLResponse) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *QueryPnLResponse) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.


[[Back to README]](../README.md)


