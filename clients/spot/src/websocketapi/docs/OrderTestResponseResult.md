# OrderTestResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**StandardCommissionForOrder** | Pointer to [**OrderTestResponseResultStandardCommissionForOrder**](OrderTestResponseResultStandardCommissionForOrder.md) |  | [optional] 
**SpecialCommissionForOrder** | Pointer to [**OrderTestResponseResultSpecialCommissionForOrder**](OrderTestResponseResultSpecialCommissionForOrder.md) |  | [optional] 
**TaxCommissionForOrder** | Pointer to [**OrderTestResponseResultStandardCommissionForOrder**](OrderTestResponseResultStandardCommissionForOrder.md) |  | [optional] 
**Discount** | Pointer to [**OrderTestResponseResultDiscount**](OrderTestResponseResultDiscount.md) |  | [optional] 

## Methods

### NewOrderTestResponseResult

`func NewOrderTestResponseResult() *OrderTestResponseResult`

NewOrderTestResponseResult instantiates a new OrderTestResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTestResponseResultWithDefaults

`func NewOrderTestResponseResultWithDefaults() *OrderTestResponseResult`

NewOrderTestResponseResultWithDefaults instantiates a new OrderTestResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandardCommissionForOrder

`func (o *OrderTestResponseResult) GetStandardCommissionForOrder() OrderTestResponseResultStandardCommissionForOrder`

GetStandardCommissionForOrder returns the StandardCommissionForOrder field if non-nil, zero value otherwise.

### GetStandardCommissionForOrderOk

`func (o *OrderTestResponseResult) GetStandardCommissionForOrderOk() (*OrderTestResponseResultStandardCommissionForOrder, bool)`

GetStandardCommissionForOrderOk returns a tuple with the StandardCommissionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCommissionForOrder

`func (o *OrderTestResponseResult) SetStandardCommissionForOrder(v OrderTestResponseResultStandardCommissionForOrder)`

SetStandardCommissionForOrder sets StandardCommissionForOrder field to given value.

### HasStandardCommissionForOrder

`func (o *OrderTestResponseResult) HasStandardCommissionForOrder() bool`

HasStandardCommissionForOrder returns a boolean if a field has been set.

### GetSpecialCommissionForOrder

`func (o *OrderTestResponseResult) GetSpecialCommissionForOrder() OrderTestResponseResultSpecialCommissionForOrder`

GetSpecialCommissionForOrder returns the SpecialCommissionForOrder field if non-nil, zero value otherwise.

### GetSpecialCommissionForOrderOk

`func (o *OrderTestResponseResult) GetSpecialCommissionForOrderOk() (*OrderTestResponseResultSpecialCommissionForOrder, bool)`

GetSpecialCommissionForOrderOk returns a tuple with the SpecialCommissionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialCommissionForOrder

`func (o *OrderTestResponseResult) SetSpecialCommissionForOrder(v OrderTestResponseResultSpecialCommissionForOrder)`

SetSpecialCommissionForOrder sets SpecialCommissionForOrder field to given value.

### HasSpecialCommissionForOrder

`func (o *OrderTestResponseResult) HasSpecialCommissionForOrder() bool`

HasSpecialCommissionForOrder returns a boolean if a field has been set.

### GetTaxCommissionForOrder

`func (o *OrderTestResponseResult) GetTaxCommissionForOrder() OrderTestResponseResultStandardCommissionForOrder`

GetTaxCommissionForOrder returns the TaxCommissionForOrder field if non-nil, zero value otherwise.

### GetTaxCommissionForOrderOk

`func (o *OrderTestResponseResult) GetTaxCommissionForOrderOk() (*OrderTestResponseResultStandardCommissionForOrder, bool)`

GetTaxCommissionForOrderOk returns a tuple with the TaxCommissionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCommissionForOrder

`func (o *OrderTestResponseResult) SetTaxCommissionForOrder(v OrderTestResponseResultStandardCommissionForOrder)`

SetTaxCommissionForOrder sets TaxCommissionForOrder field to given value.

### HasTaxCommissionForOrder

`func (o *OrderTestResponseResult) HasTaxCommissionForOrder() bool`

HasTaxCommissionForOrder returns a boolean if a field has been set.

### GetDiscount

`func (o *OrderTestResponseResult) GetDiscount() OrderTestResponseResultDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *OrderTestResponseResult) GetDiscountOk() (*OrderTestResponseResultDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *OrderTestResponseResult) SetDiscount(v OrderTestResponseResultDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *OrderTestResponseResult) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.


[[Back to README]](../README.md)


