# GetOrderDetailResponseData

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** |  | [optional] 
**OrderStatus** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**Fee** | Pointer to **string** |  | [optional] 
**FiatCurrency** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Ext** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetOrderDetailResponseData

`func NewGetOrderDetailResponseData() *GetOrderDetailResponseData`

NewGetOrderDetailResponseData instantiates a new GetOrderDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderDetailResponseDataWithDefaults

`func NewGetOrderDetailResponseDataWithDefaults() *GetOrderDetailResponseData`

NewGetOrderDetailResponseDataWithDefaults instantiates a new GetOrderDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetOrderDetailResponseData) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOrderDetailResponseData) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOrderDetailResponseData) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetOrderDetailResponseData) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderStatus

`func (o *GetOrderDetailResponseData) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *GetOrderDetailResponseData) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *GetOrderDetailResponseData) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *GetOrderDetailResponseData) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetAmount

`func (o *GetOrderDetailResponseData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetOrderDetailResponseData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetOrderDetailResponseData) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetOrderDetailResponseData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFee

`func (o *GetOrderDetailResponseData) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetOrderDetailResponseData) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetOrderDetailResponseData) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetOrderDetailResponseData) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFiatCurrency

`func (o *GetOrderDetailResponseData) GetFiatCurrency() string`

GetFiatCurrency returns the FiatCurrency field if non-nil, zero value otherwise.

### GetFiatCurrencyOk

`func (o *GetOrderDetailResponseData) GetFiatCurrencyOk() (*string, bool)`

GetFiatCurrencyOk returns a tuple with the FiatCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatCurrency

`func (o *GetOrderDetailResponseData) SetFiatCurrency(v string)`

SetFiatCurrency sets FiatCurrency field to given value.

### HasFiatCurrency

`func (o *GetOrderDetailResponseData) HasFiatCurrency() bool`

HasFiatCurrency returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetOrderDetailResponseData) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetOrderDetailResponseData) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetOrderDetailResponseData) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetOrderDetailResponseData) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetOrderDetailResponseData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetOrderDetailResponseData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetOrderDetailResponseData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetOrderDetailResponseData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetExt

`func (o *GetOrderDetailResponseData) GetExt() map[string]interface{}`

GetExt returns the Ext field if non-nil, zero value otherwise.

### GetExtOk

`func (o *GetOrderDetailResponseData) GetExtOk() (*map[string]interface{}, bool)`

GetExtOk returns a tuple with the Ext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExt

`func (o *GetOrderDetailResponseData) SetExt(v map[string]interface{})`

SetExt sets Ext field to given value.

### HasExt

`func (o *GetOrderDetailResponseData) HasExt() bool`

HasExt returns a boolean if a field has been set.


[[Back to README]](../README.md)


