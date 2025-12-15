# GetFiatPaymentsHistoryResponseDataInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderNo** | Pointer to **string** |  | [optional] 
**SourceAmount** | Pointer to **string** |  | [optional] 
**FiatCurrency** | Pointer to **string** |  | [optional] 
**ObtainAmount** | Pointer to **string** |  | [optional] 
**CryptoCurrency** | Pointer to **string** |  | [optional] 
**TotalFee** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PaymentMethod** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetFiatPaymentsHistoryResponseDataInner

`func NewGetFiatPaymentsHistoryResponseDataInner() *GetFiatPaymentsHistoryResponseDataInner`

NewGetFiatPaymentsHistoryResponseDataInner instantiates a new GetFiatPaymentsHistoryResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatPaymentsHistoryResponseDataInnerWithDefaults

`func NewGetFiatPaymentsHistoryResponseDataInnerWithDefaults() *GetFiatPaymentsHistoryResponseDataInner`

NewGetFiatPaymentsHistoryResponseDataInnerWithDefaults instantiates a new GetFiatPaymentsHistoryResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNo

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetOrderNo() string`

GetOrderNo returns the OrderNo field if non-nil, zero value otherwise.

### GetOrderNoOk

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetOrderNoOk() (*string, bool)`

GetOrderNoOk returns a tuple with the OrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNo

`func (o *GetFiatPaymentsHistoryResponseDataInner) SetOrderNo(v string)`

SetOrderNo sets OrderNo field to given value.

### HasOrderNo

`func (o *GetFiatPaymentsHistoryResponseDataInner) HasOrderNo() bool`

HasOrderNo returns a boolean if a field has been set.

### GetSourceAmount

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetSourceAmount() string`

GetSourceAmount returns the SourceAmount field if non-nil, zero value otherwise.

### GetSourceAmountOk

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetSourceAmountOk() (*string, bool)`

GetSourceAmountOk returns a tuple with the SourceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAmount

`func (o *GetFiatPaymentsHistoryResponseDataInner) SetSourceAmount(v string)`

SetSourceAmount sets SourceAmount field to given value.

### HasSourceAmount

`func (o *GetFiatPaymentsHistoryResponseDataInner) HasSourceAmount() bool`

HasSourceAmount returns a boolean if a field has been set.

### GetFiatCurrency

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetFiatCurrency() string`

GetFiatCurrency returns the FiatCurrency field if non-nil, zero value otherwise.

### GetFiatCurrencyOk

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetFiatCurrencyOk() (*string, bool)`

GetFiatCurrencyOk returns a tuple with the FiatCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatCurrency

`func (o *GetFiatPaymentsHistoryResponseDataInner) SetFiatCurrency(v string)`

SetFiatCurrency sets FiatCurrency field to given value.

### HasFiatCurrency

`func (o *GetFiatPaymentsHistoryResponseDataInner) HasFiatCurrency() bool`

HasFiatCurrency returns a boolean if a field has been set.

### GetObtainAmount

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetObtainAmount() string`

GetObtainAmount returns the ObtainAmount field if non-nil, zero value otherwise.

### GetObtainAmountOk

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetObtainAmountOk() (*string, bool)`

GetObtainAmountOk returns a tuple with the ObtainAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObtainAmount

`func (o *GetFiatPaymentsHistoryResponseDataInner) SetObtainAmount(v string)`

SetObtainAmount sets ObtainAmount field to given value.

### HasObtainAmount

`func (o *GetFiatPaymentsHistoryResponseDataInner) HasObtainAmount() bool`

HasObtainAmount returns a boolean if a field has been set.

### GetCryptoCurrency

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetCryptoCurrency() string`

GetCryptoCurrency returns the CryptoCurrency field if non-nil, zero value otherwise.

### GetCryptoCurrencyOk

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetCryptoCurrencyOk() (*string, bool)`

GetCryptoCurrencyOk returns a tuple with the CryptoCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoCurrency

`func (o *GetFiatPaymentsHistoryResponseDataInner) SetCryptoCurrency(v string)`

SetCryptoCurrency sets CryptoCurrency field to given value.

### HasCryptoCurrency

`func (o *GetFiatPaymentsHistoryResponseDataInner) HasCryptoCurrency() bool`

HasCryptoCurrency returns a boolean if a field has been set.

### GetTotalFee

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetTotalFee() string`

GetTotalFee returns the TotalFee field if non-nil, zero value otherwise.

### GetTotalFeeOk

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetTotalFeeOk() (*string, bool)`

GetTotalFeeOk returns a tuple with the TotalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFee

`func (o *GetFiatPaymentsHistoryResponseDataInner) SetTotalFee(v string)`

SetTotalFee sets TotalFee field to given value.

### HasTotalFee

`func (o *GetFiatPaymentsHistoryResponseDataInner) HasTotalFee() bool`

HasTotalFee returns a boolean if a field has been set.

### GetPrice

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetFiatPaymentsHistoryResponseDataInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetFiatPaymentsHistoryResponseDataInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStatus

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetFiatPaymentsHistoryResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetFiatPaymentsHistoryResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GetFiatPaymentsHistoryResponseDataInner) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GetFiatPaymentsHistoryResponseDataInner) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetFiatPaymentsHistoryResponseDataInner) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetFiatPaymentsHistoryResponseDataInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetFiatPaymentsHistoryResponseDataInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetFiatPaymentsHistoryResponseDataInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetFiatPaymentsHistoryResponseDataInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


