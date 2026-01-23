# OptionMarginAccountInformationResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to [**[]OptionMarginAccountInformationResponseAssetInner**](OptionMarginAccountInformationResponseAssetInner.md) |  | [optional] 
**Greek** | Pointer to [**[]OptionMarginAccountInformationResponseGreekInner**](OptionMarginAccountInformationResponseGreekInner.md) |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewOptionMarginAccountInformationResponse

`func NewOptionMarginAccountInformationResponse() *OptionMarginAccountInformationResponse`

NewOptionMarginAccountInformationResponse instantiates a new OptionMarginAccountInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionMarginAccountInformationResponseWithDefaults

`func NewOptionMarginAccountInformationResponseWithDefaults() *OptionMarginAccountInformationResponse`

NewOptionMarginAccountInformationResponseWithDefaults instantiates a new OptionMarginAccountInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *OptionMarginAccountInformationResponse) GetAsset() []OptionMarginAccountInformationResponseAssetInner`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *OptionMarginAccountInformationResponse) GetAssetOk() (*[]OptionMarginAccountInformationResponseAssetInner, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *OptionMarginAccountInformationResponse) SetAsset(v []OptionMarginAccountInformationResponseAssetInner)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *OptionMarginAccountInformationResponse) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetGreek

`func (o *OptionMarginAccountInformationResponse) GetGreek() []OptionMarginAccountInformationResponseGreekInner`

GetGreek returns the Greek field if non-nil, zero value otherwise.

### GetGreekOk

`func (o *OptionMarginAccountInformationResponse) GetGreekOk() (*[]OptionMarginAccountInformationResponseGreekInner, bool)`

GetGreekOk returns a tuple with the Greek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreek

`func (o *OptionMarginAccountInformationResponse) SetGreek(v []OptionMarginAccountInformationResponseGreekInner)`

SetGreek sets Greek field to given value.

### HasGreek

`func (o *OptionMarginAccountInformationResponse) HasGreek() bool`

HasGreek returns a boolean if a field has been set.

### GetTime

`func (o *OptionMarginAccountInformationResponse) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OptionMarginAccountInformationResponse) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OptionMarginAccountInformationResponse) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OptionMarginAccountInformationResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetCanTrade

`func (o *OptionMarginAccountInformationResponse) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *OptionMarginAccountInformationResponse) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *OptionMarginAccountInformationResponse) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *OptionMarginAccountInformationResponse) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanDeposit

`func (o *OptionMarginAccountInformationResponse) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *OptionMarginAccountInformationResponse) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *OptionMarginAccountInformationResponse) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *OptionMarginAccountInformationResponse) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *OptionMarginAccountInformationResponse) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *OptionMarginAccountInformationResponse) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *OptionMarginAccountInformationResponse) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *OptionMarginAccountInformationResponse) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetReduceOnly

`func (o *OptionMarginAccountInformationResponse) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *OptionMarginAccountInformationResponse) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *OptionMarginAccountInformationResponse) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *OptionMarginAccountInformationResponse) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.


[[Back to README]](../README.md)


