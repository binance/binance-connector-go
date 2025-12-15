# SorOrderTestResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**StandardCommissionForOrder** | Pointer to [**OrderTestResponseResultStandardCommissionForOrder**](OrderTestResponseResultStandardCommissionForOrder.md) |  | [optional] 
**TaxCommissionForOrder** | Pointer to [**OrderTestResponseResultStandardCommissionForOrder**](OrderTestResponseResultStandardCommissionForOrder.md) |  | [optional] 
**Discount** | Pointer to [**OrderTestResponseResultDiscount**](OrderTestResponseResultDiscount.md) |  | [optional] 

## Methods

### NewSorOrderTestResponseResult

`func NewSorOrderTestResponseResult() *SorOrderTestResponseResult`

NewSorOrderTestResponseResult instantiates a new SorOrderTestResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSorOrderTestResponseResultWithDefaults

`func NewSorOrderTestResponseResultWithDefaults() *SorOrderTestResponseResult`

NewSorOrderTestResponseResultWithDefaults instantiates a new SorOrderTestResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandardCommissionForOrder

`func (o *SorOrderTestResponseResult) GetStandardCommissionForOrder() OrderTestResponseResultStandardCommissionForOrder`

GetStandardCommissionForOrder returns the StandardCommissionForOrder field if non-nil, zero value otherwise.

### GetStandardCommissionForOrderOk

`func (o *SorOrderTestResponseResult) GetStandardCommissionForOrderOk() (*OrderTestResponseResultStandardCommissionForOrder, bool)`

GetStandardCommissionForOrderOk returns a tuple with the StandardCommissionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCommissionForOrder

`func (o *SorOrderTestResponseResult) SetStandardCommissionForOrder(v OrderTestResponseResultStandardCommissionForOrder)`

SetStandardCommissionForOrder sets StandardCommissionForOrder field to given value.

### HasStandardCommissionForOrder

`func (o *SorOrderTestResponseResult) HasStandardCommissionForOrder() bool`

HasStandardCommissionForOrder returns a boolean if a field has been set.

### GetTaxCommissionForOrder

`func (o *SorOrderTestResponseResult) GetTaxCommissionForOrder() OrderTestResponseResultStandardCommissionForOrder`

GetTaxCommissionForOrder returns the TaxCommissionForOrder field if non-nil, zero value otherwise.

### GetTaxCommissionForOrderOk

`func (o *SorOrderTestResponseResult) GetTaxCommissionForOrderOk() (*OrderTestResponseResultStandardCommissionForOrder, bool)`

GetTaxCommissionForOrderOk returns a tuple with the TaxCommissionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCommissionForOrder

`func (o *SorOrderTestResponseResult) SetTaxCommissionForOrder(v OrderTestResponseResultStandardCommissionForOrder)`

SetTaxCommissionForOrder sets TaxCommissionForOrder field to given value.

### HasTaxCommissionForOrder

`func (o *SorOrderTestResponseResult) HasTaxCommissionForOrder() bool`

HasTaxCommissionForOrder returns a boolean if a field has been set.

### GetDiscount

`func (o *SorOrderTestResponseResult) GetDiscount() OrderTestResponseResultDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *SorOrderTestResponseResult) GetDiscountOk() (*OrderTestResponseResultDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *SorOrderTestResponseResult) SetDiscount(v OrderTestResponseResultDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *SorOrderTestResponseResult) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.


[[Back to README]](../README.md)


