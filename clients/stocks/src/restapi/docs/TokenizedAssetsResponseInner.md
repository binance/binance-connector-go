# TokenizedAssetsResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AssetCode** | Pointer to **string** | Tokenized asset code, e.g. &#x60;AAPLB&#x60;. | [optional] 
**AssetName** | Pointer to **string** | Human-readable display name of the tokenized asset. | [optional] 
**UnderlyingEquitySymbol** | Pointer to **string** | Underlying US-equity ticker, e.g. &#x60;AAPL&#x60;. Use this when calling &#x60;/tokenized/mint&#x60; or &#x60;/tokenized/redeem&#x60;. | [optional] 
**Multiplier** | Pointer to **string** | Conversion multiplier between the underlying equity and the tokenized asset. | [optional] 
**MultiplierValid** | Pointer to **bool** | Whether the multiplier is currently valid and applicable. | [optional] 

## Methods

### NewTokenizedAssetsResponseInner

`func NewTokenizedAssetsResponseInner() *TokenizedAssetsResponseInner`

NewTokenizedAssetsResponseInner instantiates a new TokenizedAssetsResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizedAssetsResponseInnerWithDefaults

`func NewTokenizedAssetsResponseInnerWithDefaults() *TokenizedAssetsResponseInner`

NewTokenizedAssetsResponseInnerWithDefaults instantiates a new TokenizedAssetsResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetCode

`func (o *TokenizedAssetsResponseInner) GetAssetCode() string`

GetAssetCode returns the AssetCode field if non-nil, zero value otherwise.

### GetAssetCodeOk

`func (o *TokenizedAssetsResponseInner) GetAssetCodeOk() (*string, bool)`

GetAssetCodeOk returns a tuple with the AssetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCode

`func (o *TokenizedAssetsResponseInner) SetAssetCode(v string)`

SetAssetCode sets AssetCode field to given value.

### HasAssetCode

`func (o *TokenizedAssetsResponseInner) HasAssetCode() bool`

HasAssetCode returns a boolean if a field has been set.

### GetAssetName

`func (o *TokenizedAssetsResponseInner) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *TokenizedAssetsResponseInner) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *TokenizedAssetsResponseInner) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.

### HasAssetName

`func (o *TokenizedAssetsResponseInner) HasAssetName() bool`

HasAssetName returns a boolean if a field has been set.

### GetUnderlyingEquitySymbol

`func (o *TokenizedAssetsResponseInner) GetUnderlyingEquitySymbol() string`

GetUnderlyingEquitySymbol returns the UnderlyingEquitySymbol field if non-nil, zero value otherwise.

### GetUnderlyingEquitySymbolOk

`func (o *TokenizedAssetsResponseInner) GetUnderlyingEquitySymbolOk() (*string, bool)`

GetUnderlyingEquitySymbolOk returns a tuple with the UnderlyingEquitySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingEquitySymbol

`func (o *TokenizedAssetsResponseInner) SetUnderlyingEquitySymbol(v string)`

SetUnderlyingEquitySymbol sets UnderlyingEquitySymbol field to given value.

### HasUnderlyingEquitySymbol

`func (o *TokenizedAssetsResponseInner) HasUnderlyingEquitySymbol() bool`

HasUnderlyingEquitySymbol returns a boolean if a field has been set.

### GetMultiplier

`func (o *TokenizedAssetsResponseInner) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *TokenizedAssetsResponseInner) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *TokenizedAssetsResponseInner) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *TokenizedAssetsResponseInner) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetMultiplierValid

`func (o *TokenizedAssetsResponseInner) GetMultiplierValid() bool`

GetMultiplierValid returns the MultiplierValid field if non-nil, zero value otherwise.

### GetMultiplierValidOk

`func (o *TokenizedAssetsResponseInner) GetMultiplierValidOk() (*bool, bool)`

GetMultiplierValidOk returns a tuple with the MultiplierValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierValid

`func (o *TokenizedAssetsResponseInner) SetMultiplierValid(v bool)`

SetMultiplierValid sets MultiplierValid field to given value.

### HasMultiplierValid

`func (o *TokenizedAssetsResponseInner) HasMultiplierValid() bool`

HasMultiplierValid returns a boolean if a field has been set.


[[Back to README]](../README.md)


