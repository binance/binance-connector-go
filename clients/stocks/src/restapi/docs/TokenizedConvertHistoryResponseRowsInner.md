# TokenizedConvertHistoryResponseRowsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**UnderlyingAsset** | Pointer to **string** | Underlying US-equity ticker, e.g. &#x60;AAPL&#x60;. | [optional] 
**UnderlyingAssetAmount** | Pointer to **string** | Quantity of the underlying asset involved. | [optional] 
**TokenizedAsset** | Pointer to **string** | Tokenized asset, e.g. &#x60;AAPLB&#x60;. | [optional] 
**TokenizedAssetAmount** | Pointer to **string** | Quantity of the tokenized asset involved. | [optional] 
**IssuerRequestId** | Pointer to **string** | Convert request id. | [optional] 
**ConvertType** | Pointer to **string** | &#x60;MINT&#x60; or &#x60;REDEEM&#x60;. | [optional] 
**Status** | Pointer to **string** | Convert status: &#x60;P&#x60; &#x3D; processing, &#x60;S&#x60; &#x3D; success, &#x60;F&#x60; &#x3D; failed. | [optional] 
**CreatedAt** | Pointer to **int64** | Creation time (ms epoch). | [optional] 
**UpdatedAt** | Pointer to **int64** | Last update time (ms epoch). | [optional] 

## Methods

### NewTokenizedConvertHistoryResponseRowsInner

`func NewTokenizedConvertHistoryResponseRowsInner() *TokenizedConvertHistoryResponseRowsInner`

NewTokenizedConvertHistoryResponseRowsInner instantiates a new TokenizedConvertHistoryResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizedConvertHistoryResponseRowsInnerWithDefaults

`func NewTokenizedConvertHistoryResponseRowsInnerWithDefaults() *TokenizedConvertHistoryResponseRowsInner`

NewTokenizedConvertHistoryResponseRowsInnerWithDefaults instantiates a new TokenizedConvertHistoryResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderlyingAsset

`func (o *TokenizedConvertHistoryResponseRowsInner) GetUnderlyingAsset() string`

GetUnderlyingAsset returns the UnderlyingAsset field if non-nil, zero value otherwise.

### GetUnderlyingAssetOk

`func (o *TokenizedConvertHistoryResponseRowsInner) GetUnderlyingAssetOk() (*string, bool)`

GetUnderlyingAssetOk returns a tuple with the UnderlyingAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingAsset

`func (o *TokenizedConvertHistoryResponseRowsInner) SetUnderlyingAsset(v string)`

SetUnderlyingAsset sets UnderlyingAsset field to given value.

### HasUnderlyingAsset

`func (o *TokenizedConvertHistoryResponseRowsInner) HasUnderlyingAsset() bool`

HasUnderlyingAsset returns a boolean if a field has been set.

### GetUnderlyingAssetAmount

`func (o *TokenizedConvertHistoryResponseRowsInner) GetUnderlyingAssetAmount() string`

GetUnderlyingAssetAmount returns the UnderlyingAssetAmount field if non-nil, zero value otherwise.

### GetUnderlyingAssetAmountOk

`func (o *TokenizedConvertHistoryResponseRowsInner) GetUnderlyingAssetAmountOk() (*string, bool)`

GetUnderlyingAssetAmountOk returns a tuple with the UnderlyingAssetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingAssetAmount

`func (o *TokenizedConvertHistoryResponseRowsInner) SetUnderlyingAssetAmount(v string)`

SetUnderlyingAssetAmount sets UnderlyingAssetAmount field to given value.

### HasUnderlyingAssetAmount

`func (o *TokenizedConvertHistoryResponseRowsInner) HasUnderlyingAssetAmount() bool`

HasUnderlyingAssetAmount returns a boolean if a field has been set.

### GetTokenizedAsset

`func (o *TokenizedConvertHistoryResponseRowsInner) GetTokenizedAsset() string`

GetTokenizedAsset returns the TokenizedAsset field if non-nil, zero value otherwise.

### GetTokenizedAssetOk

`func (o *TokenizedConvertHistoryResponseRowsInner) GetTokenizedAssetOk() (*string, bool)`

GetTokenizedAssetOk returns a tuple with the TokenizedAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizedAsset

`func (o *TokenizedConvertHistoryResponseRowsInner) SetTokenizedAsset(v string)`

SetTokenizedAsset sets TokenizedAsset field to given value.

### HasTokenizedAsset

`func (o *TokenizedConvertHistoryResponseRowsInner) HasTokenizedAsset() bool`

HasTokenizedAsset returns a boolean if a field has been set.

### GetTokenizedAssetAmount

`func (o *TokenizedConvertHistoryResponseRowsInner) GetTokenizedAssetAmount() string`

GetTokenizedAssetAmount returns the TokenizedAssetAmount field if non-nil, zero value otherwise.

### GetTokenizedAssetAmountOk

`func (o *TokenizedConvertHistoryResponseRowsInner) GetTokenizedAssetAmountOk() (*string, bool)`

GetTokenizedAssetAmountOk returns a tuple with the TokenizedAssetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizedAssetAmount

`func (o *TokenizedConvertHistoryResponseRowsInner) SetTokenizedAssetAmount(v string)`

SetTokenizedAssetAmount sets TokenizedAssetAmount field to given value.

### HasTokenizedAssetAmount

`func (o *TokenizedConvertHistoryResponseRowsInner) HasTokenizedAssetAmount() bool`

HasTokenizedAssetAmount returns a boolean if a field has been set.

### GetIssuerRequestId

`func (o *TokenizedConvertHistoryResponseRowsInner) GetIssuerRequestId() string`

GetIssuerRequestId returns the IssuerRequestId field if non-nil, zero value otherwise.

### GetIssuerRequestIdOk

`func (o *TokenizedConvertHistoryResponseRowsInner) GetIssuerRequestIdOk() (*string, bool)`

GetIssuerRequestIdOk returns a tuple with the IssuerRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerRequestId

`func (o *TokenizedConvertHistoryResponseRowsInner) SetIssuerRequestId(v string)`

SetIssuerRequestId sets IssuerRequestId field to given value.

### HasIssuerRequestId

`func (o *TokenizedConvertHistoryResponseRowsInner) HasIssuerRequestId() bool`

HasIssuerRequestId returns a boolean if a field has been set.

### GetConvertType

`func (o *TokenizedConvertHistoryResponseRowsInner) GetConvertType() string`

GetConvertType returns the ConvertType field if non-nil, zero value otherwise.

### GetConvertTypeOk

`func (o *TokenizedConvertHistoryResponseRowsInner) GetConvertTypeOk() (*string, bool)`

GetConvertTypeOk returns a tuple with the ConvertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertType

`func (o *TokenizedConvertHistoryResponseRowsInner) SetConvertType(v string)`

SetConvertType sets ConvertType field to given value.

### HasConvertType

`func (o *TokenizedConvertHistoryResponseRowsInner) HasConvertType() bool`

HasConvertType returns a boolean if a field has been set.

### GetStatus

`func (o *TokenizedConvertHistoryResponseRowsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenizedConvertHistoryResponseRowsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenizedConvertHistoryResponseRowsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TokenizedConvertHistoryResponseRowsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TokenizedConvertHistoryResponseRowsInner) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TokenizedConvertHistoryResponseRowsInner) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TokenizedConvertHistoryResponseRowsInner) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TokenizedConvertHistoryResponseRowsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TokenizedConvertHistoryResponseRowsInner) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TokenizedConvertHistoryResponseRowsInner) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TokenizedConvertHistoryResponseRowsInner) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TokenizedConvertHistoryResponseRowsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to README]](../README.md)


