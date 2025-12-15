# GetPortfolioMarginProSpanAccountInfoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**UniMMR** | Pointer to **string** |  | [optional] 
**AccountEquity** | Pointer to **string** |  | [optional] 
**ActualEquity** | Pointer to **string** |  | [optional] 
**AccountMaintMargin** | Pointer to **string** |  | [optional] 
**RiskUnitMMList** | Pointer to [**[]GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner**](GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner.md) |  | [optional] 
**MarginMM** | Pointer to **string** |  | [optional] 
**OtherMM** | Pointer to **string** |  | [optional] 
**AccountStatus** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 

## Methods

### NewGetPortfolioMarginProSpanAccountInfoResponse

`func NewGetPortfolioMarginProSpanAccountInfoResponse() *GetPortfolioMarginProSpanAccountInfoResponse`

NewGetPortfolioMarginProSpanAccountInfoResponse instantiates a new GetPortfolioMarginProSpanAccountInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPortfolioMarginProSpanAccountInfoResponseWithDefaults

`func NewGetPortfolioMarginProSpanAccountInfoResponseWithDefaults() *GetPortfolioMarginProSpanAccountInfoResponse`

NewGetPortfolioMarginProSpanAccountInfoResponseWithDefaults instantiates a new GetPortfolioMarginProSpanAccountInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniMMR

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetUniMMR() string`

GetUniMMR returns the UniMMR field if non-nil, zero value otherwise.

### GetUniMMROk

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetUniMMROk() (*string, bool)`

GetUniMMROk returns a tuple with the UniMMR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniMMR

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetUniMMR(v string)`

SetUniMMR sets UniMMR field to given value.

### HasUniMMR

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasUniMMR() bool`

HasUniMMR returns a boolean if a field has been set.

### GetAccountEquity

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountEquity() string`

GetAccountEquity returns the AccountEquity field if non-nil, zero value otherwise.

### GetAccountEquityOk

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountEquityOk() (*string, bool)`

GetAccountEquityOk returns a tuple with the AccountEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEquity

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetAccountEquity(v string)`

SetAccountEquity sets AccountEquity field to given value.

### HasAccountEquity

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasAccountEquity() bool`

HasAccountEquity returns a boolean if a field has been set.

### GetActualEquity

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetActualEquity() string`

GetActualEquity returns the ActualEquity field if non-nil, zero value otherwise.

### GetActualEquityOk

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetActualEquityOk() (*string, bool)`

GetActualEquityOk returns a tuple with the ActualEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualEquity

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetActualEquity(v string)`

SetActualEquity sets ActualEquity field to given value.

### HasActualEquity

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasActualEquity() bool`

HasActualEquity returns a boolean if a field has been set.

### GetAccountMaintMargin

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountMaintMargin() string`

GetAccountMaintMargin returns the AccountMaintMargin field if non-nil, zero value otherwise.

### GetAccountMaintMarginOk

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountMaintMarginOk() (*string, bool)`

GetAccountMaintMarginOk returns a tuple with the AccountMaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMaintMargin

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetAccountMaintMargin(v string)`

SetAccountMaintMargin sets AccountMaintMargin field to given value.

### HasAccountMaintMargin

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasAccountMaintMargin() bool`

HasAccountMaintMargin returns a boolean if a field has been set.

### GetRiskUnitMMList

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetRiskUnitMMList() []GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner`

GetRiskUnitMMList returns the RiskUnitMMList field if non-nil, zero value otherwise.

### GetRiskUnitMMListOk

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetRiskUnitMMListOk() (*[]GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner, bool)`

GetRiskUnitMMListOk returns a tuple with the RiskUnitMMList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskUnitMMList

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetRiskUnitMMList(v []GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner)`

SetRiskUnitMMList sets RiskUnitMMList field to given value.

### HasRiskUnitMMList

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasRiskUnitMMList() bool`

HasRiskUnitMMList returns a boolean if a field has been set.

### GetMarginMM

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetMarginMM() string`

GetMarginMM returns the MarginMM field if non-nil, zero value otherwise.

### GetMarginMMOk

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetMarginMMOk() (*string, bool)`

GetMarginMMOk returns a tuple with the MarginMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginMM

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetMarginMM(v string)`

SetMarginMM sets MarginMM field to given value.

### HasMarginMM

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasMarginMM() bool`

HasMarginMM returns a boolean if a field has been set.

### GetOtherMM

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetOtherMM() string`

GetOtherMM returns the OtherMM field if non-nil, zero value otherwise.

### GetOtherMMOk

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetOtherMMOk() (*string, bool)`

GetOtherMMOk returns a tuple with the OtherMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMM

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetOtherMM(v string)`

SetOtherMM sets OtherMM field to given value.

### HasOtherMM

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasOtherMM() bool`

HasOtherMM returns a boolean if a field has been set.

### GetAccountStatus

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetAccountType

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.


[[Back to README]](../README.md)


