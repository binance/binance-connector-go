# PriceStreamResponseRatesInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**S** | Pointer to **string** | Symbol (UPPERCASE ticker), e.g. &#x60;\&quot;AAPL\&quot;&#x60;. | [optional] 
**Ac** | Pointer to **string** | Internal asset code, &#x60;EQ_{symbol}&#x60;. Reference only; do not use as a trading identifier. | [optional] 
**P** | Pointer to **string** | Latest price, trailing-zero stripped. | [optional] 
**T** | Pointer to **NullableInt64** | Price time (epoch milliseconds UTC); may be null. | [optional] 
**Pc** | Pointer to **string** | Previous day&#39;s RTH close price (reference); absent when unknown. | [optional] 
**Tc** | Pointer to **string** | Today&#39;s RTH close price; absent before After-Hours. | [optional] 
**Mp** | Pointer to **string** | Per-symbol market phase: &#x60;C&#x60; Closed, &#x60;ON&#x60; Overnight, &#x60;PRE&#x60; Pre-Market, &#x60;O&#x60; Market Open, &#x60;POST&#x60; Post-Market. | [optional] 

## Methods

### NewPriceStreamResponseRatesInner

`func NewPriceStreamResponseRatesInner() *PriceStreamResponseRatesInner`

NewPriceStreamResponseRatesInner instantiates a new PriceStreamResponseRatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceStreamResponseRatesInnerWithDefaults

`func NewPriceStreamResponseRatesInnerWithDefaults() *PriceStreamResponseRatesInner`

NewPriceStreamResponseRatesInnerWithDefaults instantiates a new PriceStreamResponseRatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS

`func (o *PriceStreamResponseRatesInner) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *PriceStreamResponseRatesInner) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *PriceStreamResponseRatesInner) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *PriceStreamResponseRatesInner) HasS() bool`

HasS returns a boolean if a field has been set.

### GetAc

`func (o *PriceStreamResponseRatesInner) GetAc() string`

GetAc returns the Ac field if non-nil, zero value otherwise.

### GetAcOk

`func (o *PriceStreamResponseRatesInner) GetAcOk() (*string, bool)`

GetAcOk returns a tuple with the Ac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAc

`func (o *PriceStreamResponseRatesInner) SetAc(v string)`

SetAc sets Ac field to given value.

### HasAc

`func (o *PriceStreamResponseRatesInner) HasAc() bool`

HasAc returns a boolean if a field has been set.

### GetP

`func (o *PriceStreamResponseRatesInner) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *PriceStreamResponseRatesInner) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *PriceStreamResponseRatesInner) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *PriceStreamResponseRatesInner) HasP() bool`

HasP returns a boolean if a field has been set.

### GetT

`func (o *PriceStreamResponseRatesInner) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *PriceStreamResponseRatesInner) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *PriceStreamResponseRatesInner) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *PriceStreamResponseRatesInner) HasT() bool`

HasT returns a boolean if a field has been set.

### SetTNil

`func (o *PriceStreamResponseRatesInner) SetTNil(b bool)`

 SetTNil sets the value for T to be an explicit nil

### UnsetT
`func (o *PriceStreamResponseRatesInner) UnsetT()`

UnsetT ensures that no value is present for T, not even an explicit nil
### GetPc

`func (o *PriceStreamResponseRatesInner) GetPc() string`

GetPc returns the Pc field if non-nil, zero value otherwise.

### GetPcOk

`func (o *PriceStreamResponseRatesInner) GetPcOk() (*string, bool)`

GetPcOk returns a tuple with the Pc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPc

`func (o *PriceStreamResponseRatesInner) SetPc(v string)`

SetPc sets Pc field to given value.

### HasPc

`func (o *PriceStreamResponseRatesInner) HasPc() bool`

HasPc returns a boolean if a field has been set.

### GetTc

`func (o *PriceStreamResponseRatesInner) GetTc() string`

GetTc returns the Tc field if non-nil, zero value otherwise.

### GetTcOk

`func (o *PriceStreamResponseRatesInner) GetTcOk() (*string, bool)`

GetTcOk returns a tuple with the Tc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTc

`func (o *PriceStreamResponseRatesInner) SetTc(v string)`

SetTc sets Tc field to given value.

### HasTc

`func (o *PriceStreamResponseRatesInner) HasTc() bool`

HasTc returns a boolean if a field has been set.

### GetMp

`func (o *PriceStreamResponseRatesInner) GetMp() string`

GetMp returns the Mp field if non-nil, zero value otherwise.

### GetMpOk

`func (o *PriceStreamResponseRatesInner) GetMpOk() (*string, bool)`

GetMpOk returns a tuple with the Mp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp

`func (o *PriceStreamResponseRatesInner) SetMp(v string)`

SetMp sets Mp field to given value.

### HasMp

`func (o *PriceStreamResponseRatesInner) HasMp() bool`

HasMp returns a boolean if a field has been set.


[[Back to README]](../README.md)


