# AccountCommissionResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**StandardCommission** | Pointer to [**AccountCommissionResponseStandardCommission**](AccountCommissionResponseStandardCommission.md) |  | [optional] 
**SpecialCommission** | Pointer to [**AccountCommissionResponseSpecialCommission**](AccountCommissionResponseSpecialCommission.md) |  | [optional] 
**TaxCommission** | Pointer to [**AccountCommissionResponseTaxCommission**](AccountCommissionResponseTaxCommission.md) |  | [optional] 
**Discount** | Pointer to [**AccountCommissionResponseDiscount**](AccountCommissionResponseDiscount.md) |  | [optional] 

## Methods

### NewAccountCommissionResponse

`func NewAccountCommissionResponse() *AccountCommissionResponse`

NewAccountCommissionResponse instantiates a new AccountCommissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCommissionResponseWithDefaults

`func NewAccountCommissionResponseWithDefaults() *AccountCommissionResponse`

NewAccountCommissionResponseWithDefaults instantiates a new AccountCommissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *AccountCommissionResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AccountCommissionResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AccountCommissionResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AccountCommissionResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetStandardCommission

`func (o *AccountCommissionResponse) GetStandardCommission() AccountCommissionResponseStandardCommission`

GetStandardCommission returns the StandardCommission field if non-nil, zero value otherwise.

### GetStandardCommissionOk

`func (o *AccountCommissionResponse) GetStandardCommissionOk() (*AccountCommissionResponseStandardCommission, bool)`

GetStandardCommissionOk returns a tuple with the StandardCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCommission

`func (o *AccountCommissionResponse) SetStandardCommission(v AccountCommissionResponseStandardCommission)`

SetStandardCommission sets StandardCommission field to given value.

### HasStandardCommission

`func (o *AccountCommissionResponse) HasStandardCommission() bool`

HasStandardCommission returns a boolean if a field has been set.

### GetSpecialCommission

`func (o *AccountCommissionResponse) GetSpecialCommission() AccountCommissionResponseSpecialCommission`

GetSpecialCommission returns the SpecialCommission field if non-nil, zero value otherwise.

### GetSpecialCommissionOk

`func (o *AccountCommissionResponse) GetSpecialCommissionOk() (*AccountCommissionResponseSpecialCommission, bool)`

GetSpecialCommissionOk returns a tuple with the SpecialCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialCommission

`func (o *AccountCommissionResponse) SetSpecialCommission(v AccountCommissionResponseSpecialCommission)`

SetSpecialCommission sets SpecialCommission field to given value.

### HasSpecialCommission

`func (o *AccountCommissionResponse) HasSpecialCommission() bool`

HasSpecialCommission returns a boolean if a field has been set.

### GetTaxCommission

`func (o *AccountCommissionResponse) GetTaxCommission() AccountCommissionResponseTaxCommission`

GetTaxCommission returns the TaxCommission field if non-nil, zero value otherwise.

### GetTaxCommissionOk

`func (o *AccountCommissionResponse) GetTaxCommissionOk() (*AccountCommissionResponseTaxCommission, bool)`

GetTaxCommissionOk returns a tuple with the TaxCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCommission

`func (o *AccountCommissionResponse) SetTaxCommission(v AccountCommissionResponseTaxCommission)`

SetTaxCommission sets TaxCommission field to given value.

### HasTaxCommission

`func (o *AccountCommissionResponse) HasTaxCommission() bool`

HasTaxCommission returns a boolean if a field has been set.

### GetDiscount

`func (o *AccountCommissionResponse) GetDiscount() AccountCommissionResponseDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *AccountCommissionResponse) GetDiscountOk() (*AccountCommissionResponseDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *AccountCommissionResponse) SetDiscount(v AccountCommissionResponseDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *AccountCommissionResponse) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.


[[Back to README]](../README.md)


