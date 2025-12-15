# AccountCommissionResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**StandardCommission** | Pointer to [**AccountCommissionResponseResultStandardCommission**](AccountCommissionResponseResultStandardCommission.md) |  | [optional] 
**SpecialCommission** | Pointer to [**AccountCommissionResponseResultSpecialCommission**](AccountCommissionResponseResultSpecialCommission.md) |  | [optional] 
**TaxCommission** | Pointer to [**AccountCommissionResponseResultTaxCommission**](AccountCommissionResponseResultTaxCommission.md) |  | [optional] 
**Discount** | Pointer to [**AccountCommissionResponseResultDiscount**](AccountCommissionResponseResultDiscount.md) |  | [optional] 

## Methods

### NewAccountCommissionResponseResult

`func NewAccountCommissionResponseResult() *AccountCommissionResponseResult`

NewAccountCommissionResponseResult instantiates a new AccountCommissionResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCommissionResponseResultWithDefaults

`func NewAccountCommissionResponseResultWithDefaults() *AccountCommissionResponseResult`

NewAccountCommissionResponseResultWithDefaults instantiates a new AccountCommissionResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *AccountCommissionResponseResult) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AccountCommissionResponseResult) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AccountCommissionResponseResult) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AccountCommissionResponseResult) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetStandardCommission

`func (o *AccountCommissionResponseResult) GetStandardCommission() AccountCommissionResponseResultStandardCommission`

GetStandardCommission returns the StandardCommission field if non-nil, zero value otherwise.

### GetStandardCommissionOk

`func (o *AccountCommissionResponseResult) GetStandardCommissionOk() (*AccountCommissionResponseResultStandardCommission, bool)`

GetStandardCommissionOk returns a tuple with the StandardCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCommission

`func (o *AccountCommissionResponseResult) SetStandardCommission(v AccountCommissionResponseResultStandardCommission)`

SetStandardCommission sets StandardCommission field to given value.

### HasStandardCommission

`func (o *AccountCommissionResponseResult) HasStandardCommission() bool`

HasStandardCommission returns a boolean if a field has been set.

### GetSpecialCommission

`func (o *AccountCommissionResponseResult) GetSpecialCommission() AccountCommissionResponseResultSpecialCommission`

GetSpecialCommission returns the SpecialCommission field if non-nil, zero value otherwise.

### GetSpecialCommissionOk

`func (o *AccountCommissionResponseResult) GetSpecialCommissionOk() (*AccountCommissionResponseResultSpecialCommission, bool)`

GetSpecialCommissionOk returns a tuple with the SpecialCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialCommission

`func (o *AccountCommissionResponseResult) SetSpecialCommission(v AccountCommissionResponseResultSpecialCommission)`

SetSpecialCommission sets SpecialCommission field to given value.

### HasSpecialCommission

`func (o *AccountCommissionResponseResult) HasSpecialCommission() bool`

HasSpecialCommission returns a boolean if a field has been set.

### GetTaxCommission

`func (o *AccountCommissionResponseResult) GetTaxCommission() AccountCommissionResponseResultTaxCommission`

GetTaxCommission returns the TaxCommission field if non-nil, zero value otherwise.

### GetTaxCommissionOk

`func (o *AccountCommissionResponseResult) GetTaxCommissionOk() (*AccountCommissionResponseResultTaxCommission, bool)`

GetTaxCommissionOk returns a tuple with the TaxCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCommission

`func (o *AccountCommissionResponseResult) SetTaxCommission(v AccountCommissionResponseResultTaxCommission)`

SetTaxCommission sets TaxCommission field to given value.

### HasTaxCommission

`func (o *AccountCommissionResponseResult) HasTaxCommission() bool`

HasTaxCommission returns a boolean if a field has been set.

### GetDiscount

`func (o *AccountCommissionResponseResult) GetDiscount() AccountCommissionResponseResultDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *AccountCommissionResponseResult) GetDiscountOk() (*AccountCommissionResponseResultDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *AccountCommissionResponseResult) SetDiscount(v AccountCommissionResponseResultDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *AccountCommissionResponseResult) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.


[[Back to README]](../README.md)


