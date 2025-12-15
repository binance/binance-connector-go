# SorOrderTestResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**StandardCommissionForOrder** | Pointer to [**OrderTestResponseStandardCommissionForOrder**](OrderTestResponseStandardCommissionForOrder.md) |  | [optional] 
**TaxCommissionForOrder** | Pointer to [**OrderTestResponseStandardCommissionForOrder**](OrderTestResponseStandardCommissionForOrder.md) |  | [optional] 
**Discount** | Pointer to [**OrderTestResponseDiscount**](OrderTestResponseDiscount.md) |  | [optional] 

## Methods

### NewSorOrderTestResponse

`func NewSorOrderTestResponse() *SorOrderTestResponse`

NewSorOrderTestResponse instantiates a new SorOrderTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSorOrderTestResponseWithDefaults

`func NewSorOrderTestResponseWithDefaults() *SorOrderTestResponse`

NewSorOrderTestResponseWithDefaults instantiates a new SorOrderTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandardCommissionForOrder

`func (o *SorOrderTestResponse) GetStandardCommissionForOrder() OrderTestResponseStandardCommissionForOrder`

GetStandardCommissionForOrder returns the StandardCommissionForOrder field if non-nil, zero value otherwise.

### GetStandardCommissionForOrderOk

`func (o *SorOrderTestResponse) GetStandardCommissionForOrderOk() (*OrderTestResponseStandardCommissionForOrder, bool)`

GetStandardCommissionForOrderOk returns a tuple with the StandardCommissionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCommissionForOrder

`func (o *SorOrderTestResponse) SetStandardCommissionForOrder(v OrderTestResponseStandardCommissionForOrder)`

SetStandardCommissionForOrder sets StandardCommissionForOrder field to given value.

### HasStandardCommissionForOrder

`func (o *SorOrderTestResponse) HasStandardCommissionForOrder() bool`

HasStandardCommissionForOrder returns a boolean if a field has been set.

### GetTaxCommissionForOrder

`func (o *SorOrderTestResponse) GetTaxCommissionForOrder() OrderTestResponseStandardCommissionForOrder`

GetTaxCommissionForOrder returns the TaxCommissionForOrder field if non-nil, zero value otherwise.

### GetTaxCommissionForOrderOk

`func (o *SorOrderTestResponse) GetTaxCommissionForOrderOk() (*OrderTestResponseStandardCommissionForOrder, bool)`

GetTaxCommissionForOrderOk returns a tuple with the TaxCommissionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCommissionForOrder

`func (o *SorOrderTestResponse) SetTaxCommissionForOrder(v OrderTestResponseStandardCommissionForOrder)`

SetTaxCommissionForOrder sets TaxCommissionForOrder field to given value.

### HasTaxCommissionForOrder

`func (o *SorOrderTestResponse) HasTaxCommissionForOrder() bool`

HasTaxCommissionForOrder returns a boolean if a field has been set.

### GetDiscount

`func (o *SorOrderTestResponse) GetDiscount() OrderTestResponseDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *SorOrderTestResponse) GetDiscountOk() (*OrderTestResponseDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *SorOrderTestResponse) SetDiscount(v OrderTestResponseDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *SorOrderTestResponse) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.


[[Back to README]](../README.md)


