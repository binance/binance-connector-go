# GetDetailOnSubAccountsMarginAccountResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**MarginLevel** | Pointer to **string** |  | [optional] 
**TotalAssetOfBtc** | Pointer to **string** |  | [optional] 
**TotalLiabilityOfBtc** | Pointer to **string** |  | [optional] 
**TotalNetAssetOfBtc** | Pointer to **string** |  | [optional] 
**MarginTradeCoeffVo** | Pointer to [**GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo**](GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo.md) |  | [optional] 
**MarginUserAssetVoList** | Pointer to [**[]GetDetailOnSubAccountsMarginAccountResponseMarginUserAssetVoListInner**](GetDetailOnSubAccountsMarginAccountResponseMarginUserAssetVoListInner.md) |  | [optional] 

## Methods

### NewGetDetailOnSubAccountsMarginAccountResponse

`func NewGetDetailOnSubAccountsMarginAccountResponse() *GetDetailOnSubAccountsMarginAccountResponse`

NewGetDetailOnSubAccountsMarginAccountResponse instantiates a new GetDetailOnSubAccountsMarginAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDetailOnSubAccountsMarginAccountResponseWithDefaults

`func NewGetDetailOnSubAccountsMarginAccountResponseWithDefaults() *GetDetailOnSubAccountsMarginAccountResponse`

NewGetDetailOnSubAccountsMarginAccountResponseWithDefaults instantiates a new GetDetailOnSubAccountsMarginAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetDetailOnSubAccountsMarginAccountResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetDetailOnSubAccountsMarginAccountResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMarginLevel

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *GetDetailOnSubAccountsMarginAccountResponse) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.

### HasMarginLevel

`func (o *GetDetailOnSubAccountsMarginAccountResponse) HasMarginLevel() bool`

HasMarginLevel returns a boolean if a field has been set.

### GetTotalAssetOfBtc

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetTotalAssetOfBtc() string`

GetTotalAssetOfBtc returns the TotalAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalAssetOfBtcOk

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetTotalAssetOfBtcOk() (*string, bool)`

GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssetOfBtc

`func (o *GetDetailOnSubAccountsMarginAccountResponse) SetTotalAssetOfBtc(v string)`

SetTotalAssetOfBtc sets TotalAssetOfBtc field to given value.

### HasTotalAssetOfBtc

`func (o *GetDetailOnSubAccountsMarginAccountResponse) HasTotalAssetOfBtc() bool`

HasTotalAssetOfBtc returns a boolean if a field has been set.

### GetTotalLiabilityOfBtc

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetTotalLiabilityOfBtc() string`

GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field if non-nil, zero value otherwise.

### GetTotalLiabilityOfBtcOk

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetTotalLiabilityOfBtcOk() (*string, bool)`

GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilityOfBtc

`func (o *GetDetailOnSubAccountsMarginAccountResponse) SetTotalLiabilityOfBtc(v string)`

SetTotalLiabilityOfBtc sets TotalLiabilityOfBtc field to given value.

### HasTotalLiabilityOfBtc

`func (o *GetDetailOnSubAccountsMarginAccountResponse) HasTotalLiabilityOfBtc() bool`

HasTotalLiabilityOfBtc returns a boolean if a field has been set.

### GetTotalNetAssetOfBtc

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetTotalNetAssetOfBtc() string`

GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalNetAssetOfBtcOk

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetTotalNetAssetOfBtcOk() (*string, bool)`

GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAssetOfBtc

`func (o *GetDetailOnSubAccountsMarginAccountResponse) SetTotalNetAssetOfBtc(v string)`

SetTotalNetAssetOfBtc sets TotalNetAssetOfBtc field to given value.

### HasTotalNetAssetOfBtc

`func (o *GetDetailOnSubAccountsMarginAccountResponse) HasTotalNetAssetOfBtc() bool`

HasTotalNetAssetOfBtc returns a boolean if a field has been set.

### GetMarginTradeCoeffVo

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetMarginTradeCoeffVo() GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo`

GetMarginTradeCoeffVo returns the MarginTradeCoeffVo field if non-nil, zero value otherwise.

### GetMarginTradeCoeffVoOk

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetMarginTradeCoeffVoOk() (*GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo, bool)`

GetMarginTradeCoeffVoOk returns a tuple with the MarginTradeCoeffVo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginTradeCoeffVo

`func (o *GetDetailOnSubAccountsMarginAccountResponse) SetMarginTradeCoeffVo(v GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo)`

SetMarginTradeCoeffVo sets MarginTradeCoeffVo field to given value.

### HasMarginTradeCoeffVo

`func (o *GetDetailOnSubAccountsMarginAccountResponse) HasMarginTradeCoeffVo() bool`

HasMarginTradeCoeffVo returns a boolean if a field has been set.

### GetMarginUserAssetVoList

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetMarginUserAssetVoList() []GetDetailOnSubAccountsMarginAccountResponseMarginUserAssetVoListInner`

GetMarginUserAssetVoList returns the MarginUserAssetVoList field if non-nil, zero value otherwise.

### GetMarginUserAssetVoListOk

`func (o *GetDetailOnSubAccountsMarginAccountResponse) GetMarginUserAssetVoListOk() (*[]GetDetailOnSubAccountsMarginAccountResponseMarginUserAssetVoListInner, bool)`

GetMarginUserAssetVoListOk returns a tuple with the MarginUserAssetVoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginUserAssetVoList

`func (o *GetDetailOnSubAccountsMarginAccountResponse) SetMarginUserAssetVoList(v []GetDetailOnSubAccountsMarginAccountResponseMarginUserAssetVoListInner)`

SetMarginUserAssetVoList sets MarginUserAssetVoList field to given value.

### HasMarginUserAssetVoList

`func (o *GetDetailOnSubAccountsMarginAccountResponse) HasMarginUserAssetVoList() bool`

HasMarginUserAssetVoList returns a boolean if a field has been set.


[[Back to README]](../README.md)


