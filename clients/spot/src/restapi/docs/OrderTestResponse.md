# OrderTestResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**StandardCommissionForOrder** | Pointer to [**OrderTestResponseStandardCommissionForOrder**](OrderTestResponseStandardCommissionForOrder.md) |  | [optional] 
**SpecialCommissionForOrder** | Pointer to [**OrderTestResponseSpecialCommissionForOrder**](OrderTestResponseSpecialCommissionForOrder.md) |  | [optional] 
**TaxCommissionForOrder** | Pointer to [**OrderTestResponseStandardCommissionForOrder**](OrderTestResponseStandardCommissionForOrder.md) |  | [optional] 
**Discount** | Pointer to [**OrderTestResponseDiscount**](OrderTestResponseDiscount.md) |  | [optional] 

## Methods

### NewOrderTestResponse

`func NewOrderTestResponse() *OrderTestResponse`

NewOrderTestResponse instantiates a new OrderTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTestResponseWithDefaults

`func NewOrderTestResponseWithDefaults() *OrderTestResponse`

NewOrderTestResponseWithDefaults instantiates a new OrderTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandardCommissionForOrder

`func (o *OrderTestResponse) GetStandardCommissionForOrder() OrderTestResponseStandardCommissionForOrder`

GetStandardCommissionForOrder returns the StandardCommissionForOrder field if non-nil, zero value otherwise.

### GetStandardCommissionForOrderOk

`func (o *OrderTestResponse) GetStandardCommissionForOrderOk() (*OrderTestResponseStandardCommissionForOrder, bool)`

GetStandardCommissionForOrderOk returns a tuple with the StandardCommissionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCommissionForOrder

`func (o *OrderTestResponse) SetStandardCommissionForOrder(v OrderTestResponseStandardCommissionForOrder)`

SetStandardCommissionForOrder sets StandardCommissionForOrder field to given value.

### HasStandardCommissionForOrder

`func (o *OrderTestResponse) HasStandardCommissionForOrder() bool`

HasStandardCommissionForOrder returns a boolean if a field has been set.

### GetSpecialCommissionForOrder

`func (o *OrderTestResponse) GetSpecialCommissionForOrder() OrderTestResponseSpecialCommissionForOrder`

GetSpecialCommissionForOrder returns the SpecialCommissionForOrder field if non-nil, zero value otherwise.

### GetSpecialCommissionForOrderOk

`func (o *OrderTestResponse) GetSpecialCommissionForOrderOk() (*OrderTestResponseSpecialCommissionForOrder, bool)`

GetSpecialCommissionForOrderOk returns a tuple with the SpecialCommissionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialCommissionForOrder

`func (o *OrderTestResponse) SetSpecialCommissionForOrder(v OrderTestResponseSpecialCommissionForOrder)`

SetSpecialCommissionForOrder sets SpecialCommissionForOrder field to given value.

### HasSpecialCommissionForOrder

`func (o *OrderTestResponse) HasSpecialCommissionForOrder() bool`

HasSpecialCommissionForOrder returns a boolean if a field has been set.

### GetTaxCommissionForOrder

`func (o *OrderTestResponse) GetTaxCommissionForOrder() OrderTestResponseStandardCommissionForOrder`

GetTaxCommissionForOrder returns the TaxCommissionForOrder field if non-nil, zero value otherwise.

### GetTaxCommissionForOrderOk

`func (o *OrderTestResponse) GetTaxCommissionForOrderOk() (*OrderTestResponseStandardCommissionForOrder, bool)`

GetTaxCommissionForOrderOk returns a tuple with the TaxCommissionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCommissionForOrder

`func (o *OrderTestResponse) SetTaxCommissionForOrder(v OrderTestResponseStandardCommissionForOrder)`

SetTaxCommissionForOrder sets TaxCommissionForOrder field to given value.

### HasTaxCommissionForOrder

`func (o *OrderTestResponse) HasTaxCommissionForOrder() bool`

HasTaxCommissionForOrder returns a boolean if a field has been set.

### GetDiscount

`func (o *OrderTestResponse) GetDiscount() OrderTestResponseDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *OrderTestResponse) GetDiscountOk() (*OrderTestResponseDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *OrderTestResponse) SetDiscount(v OrderTestResponseDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *OrderTestResponse) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.


[[Back to README]](../README.md)


