# OptionMarkPriceResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**S** | Pointer to **string** | Symbol | [optional] 
**Mp** | Pointer to **string** | Mark price | [optional] 
**E** | Pointer to **int64** | Event time | [optional] 
**E** | Pointer to **string** | Event type | [optional] 
**I** | Pointer to **string** | Index price | [optional] 
**P** | Pointer to **string** | Estimated Settle Price, only useful in the 0.5 hour before the settlement starts | [optional] 
**Bo** | Pointer to **string** | The best buy price | [optional] 
**Ao** | Pointer to **string** | The best sell price | [optional] 
**Bq** | Pointer to **string** | The best buy quantity | [optional] 
**Aq** | Pointer to **string** | The best sell quantity | [optional] 
**B** | Pointer to **string** | BuyImplied volatility | [optional] 
**A** | Pointer to **string** | SellImplied volatility | [optional] 
**Hl** | Pointer to **string** | Buy Maximum price | [optional] 
**Ll** | Pointer to **string** | Sell Minimum price | [optional] 
**Vo** | Pointer to **string** | volatility | [optional] 
**Rf** | Pointer to **string** | risk free rate | [optional] 
**D** | Pointer to **string** | delta | [optional] 
**T** | Pointer to **string** | theta | [optional] 
**G** | Pointer to **string** | gamma | [optional] 
**V** | Pointer to **string** | vega | [optional] 

## Methods

### NewOptionMarkPriceResponseInner

`func NewOptionMarkPriceResponseInner() *OptionMarkPriceResponseInner`

NewOptionMarkPriceResponseInner instantiates a new OptionMarkPriceResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionMarkPriceResponseInnerWithDefaults

`func NewOptionMarkPriceResponseInnerWithDefaults() *OptionMarkPriceResponseInner`

NewOptionMarkPriceResponseInnerWithDefaults instantiates a new OptionMarkPriceResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS

`func (o *OptionMarkPriceResponseInner) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *OptionMarkPriceResponseInner) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *OptionMarkPriceResponseInner) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *OptionMarkPriceResponseInner) HasS() bool`

HasS returns a boolean if a field has been set.

### GetMp

`func (o *OptionMarkPriceResponseInner) GetMp() string`

GetMp returns the Mp field if non-nil, zero value otherwise.

### GetMpOk

`func (o *OptionMarkPriceResponseInner) GetMpOk() (*string, bool)`

GetMpOk returns a tuple with the Mp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp

`func (o *OptionMarkPriceResponseInner) SetMp(v string)`

SetMp sets Mp field to given value.

### HasMp

`func (o *OptionMarkPriceResponseInner) HasMp() bool`

HasMp returns a boolean if a field has been set.

### GetE

`func (o *OptionMarkPriceResponseInner) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OptionMarkPriceResponseInner) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OptionMarkPriceResponseInner) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *OptionMarkPriceResponseInner) HasE() bool`

HasE returns a boolean if a field has been set.

### GetE

`func (o *OptionMarkPriceResponseInner) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OptionMarkPriceResponseInner) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OptionMarkPriceResponseInner) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *OptionMarkPriceResponseInner) HasE() bool`

HasE returns a boolean if a field has been set.

### GetI

`func (o *OptionMarkPriceResponseInner) GetI() string`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *OptionMarkPriceResponseInner) GetIOk() (*string, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *OptionMarkPriceResponseInner) SetI(v string)`

SetI sets I field to given value.

### HasI

`func (o *OptionMarkPriceResponseInner) HasI() bool`

HasI returns a boolean if a field has been set.

### GetP

`func (o *OptionMarkPriceResponseInner) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *OptionMarkPriceResponseInner) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *OptionMarkPriceResponseInner) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *OptionMarkPriceResponseInner) HasP() bool`

HasP returns a boolean if a field has been set.

### GetBo

`func (o *OptionMarkPriceResponseInner) GetBo() string`

GetBo returns the Bo field if non-nil, zero value otherwise.

### GetBoOk

`func (o *OptionMarkPriceResponseInner) GetBoOk() (*string, bool)`

GetBoOk returns a tuple with the Bo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBo

`func (o *OptionMarkPriceResponseInner) SetBo(v string)`

SetBo sets Bo field to given value.

### HasBo

`func (o *OptionMarkPriceResponseInner) HasBo() bool`

HasBo returns a boolean if a field has been set.

### GetAo

`func (o *OptionMarkPriceResponseInner) GetAo() string`

GetAo returns the Ao field if non-nil, zero value otherwise.

### GetAoOk

`func (o *OptionMarkPriceResponseInner) GetAoOk() (*string, bool)`

GetAoOk returns a tuple with the Ao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAo

`func (o *OptionMarkPriceResponseInner) SetAo(v string)`

SetAo sets Ao field to given value.

### HasAo

`func (o *OptionMarkPriceResponseInner) HasAo() bool`

HasAo returns a boolean if a field has been set.

### GetBq

`func (o *OptionMarkPriceResponseInner) GetBq() string`

GetBq returns the Bq field if non-nil, zero value otherwise.

### GetBqOk

`func (o *OptionMarkPriceResponseInner) GetBqOk() (*string, bool)`

GetBqOk returns a tuple with the Bq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBq

`func (o *OptionMarkPriceResponseInner) SetBq(v string)`

SetBq sets Bq field to given value.

### HasBq

`func (o *OptionMarkPriceResponseInner) HasBq() bool`

HasBq returns a boolean if a field has been set.

### GetAq

`func (o *OptionMarkPriceResponseInner) GetAq() string`

GetAq returns the Aq field if non-nil, zero value otherwise.

### GetAqOk

`func (o *OptionMarkPriceResponseInner) GetAqOk() (*string, bool)`

GetAqOk returns a tuple with the Aq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAq

`func (o *OptionMarkPriceResponseInner) SetAq(v string)`

SetAq sets Aq field to given value.

### HasAq

`func (o *OptionMarkPriceResponseInner) HasAq() bool`

HasAq returns a boolean if a field has been set.

### GetB

`func (o *OptionMarkPriceResponseInner) GetB() string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *OptionMarkPriceResponseInner) GetBOk() (*string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *OptionMarkPriceResponseInner) SetB(v string)`

SetB sets B field to given value.

### HasB

`func (o *OptionMarkPriceResponseInner) HasB() bool`

HasB returns a boolean if a field has been set.

### GetA

`func (o *OptionMarkPriceResponseInner) GetA() string`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *OptionMarkPriceResponseInner) GetAOk() (*string, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *OptionMarkPriceResponseInner) SetA(v string)`

SetA sets A field to given value.

### HasA

`func (o *OptionMarkPriceResponseInner) HasA() bool`

HasA returns a boolean if a field has been set.

### GetHl

`func (o *OptionMarkPriceResponseInner) GetHl() string`

GetHl returns the Hl field if non-nil, zero value otherwise.

### GetHlOk

`func (o *OptionMarkPriceResponseInner) GetHlOk() (*string, bool)`

GetHlOk returns a tuple with the Hl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHl

`func (o *OptionMarkPriceResponseInner) SetHl(v string)`

SetHl sets Hl field to given value.

### HasHl

`func (o *OptionMarkPriceResponseInner) HasHl() bool`

HasHl returns a boolean if a field has been set.

### GetLl

`func (o *OptionMarkPriceResponseInner) GetLl() string`

GetLl returns the Ll field if non-nil, zero value otherwise.

### GetLlOk

`func (o *OptionMarkPriceResponseInner) GetLlOk() (*string, bool)`

GetLlOk returns a tuple with the Ll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLl

`func (o *OptionMarkPriceResponseInner) SetLl(v string)`

SetLl sets Ll field to given value.

### HasLl

`func (o *OptionMarkPriceResponseInner) HasLl() bool`

HasLl returns a boolean if a field has been set.

### GetVo

`func (o *OptionMarkPriceResponseInner) GetVo() string`

GetVo returns the Vo field if non-nil, zero value otherwise.

### GetVoOk

`func (o *OptionMarkPriceResponseInner) GetVoOk() (*string, bool)`

GetVoOk returns a tuple with the Vo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVo

`func (o *OptionMarkPriceResponseInner) SetVo(v string)`

SetVo sets Vo field to given value.

### HasVo

`func (o *OptionMarkPriceResponseInner) HasVo() bool`

HasVo returns a boolean if a field has been set.

### GetRf

`func (o *OptionMarkPriceResponseInner) GetRf() string`

GetRf returns the Rf field if non-nil, zero value otherwise.

### GetRfOk

`func (o *OptionMarkPriceResponseInner) GetRfOk() (*string, bool)`

GetRfOk returns a tuple with the Rf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRf

`func (o *OptionMarkPriceResponseInner) SetRf(v string)`

SetRf sets Rf field to given value.

### HasRf

`func (o *OptionMarkPriceResponseInner) HasRf() bool`

HasRf returns a boolean if a field has been set.

### GetD

`func (o *OptionMarkPriceResponseInner) GetD() string`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *OptionMarkPriceResponseInner) GetDOk() (*string, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *OptionMarkPriceResponseInner) SetD(v string)`

SetD sets D field to given value.

### HasD

`func (o *OptionMarkPriceResponseInner) HasD() bool`

HasD returns a boolean if a field has been set.

### GetT

`func (o *OptionMarkPriceResponseInner) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *OptionMarkPriceResponseInner) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *OptionMarkPriceResponseInner) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *OptionMarkPriceResponseInner) HasT() bool`

HasT returns a boolean if a field has been set.

### GetG

`func (o *OptionMarkPriceResponseInner) GetG() string`

GetG returns the G field if non-nil, zero value otherwise.

### GetGOk

`func (o *OptionMarkPriceResponseInner) GetGOk() (*string, bool)`

GetGOk returns a tuple with the G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetG

`func (o *OptionMarkPriceResponseInner) SetG(v string)`

SetG sets G field to given value.

### HasG

`func (o *OptionMarkPriceResponseInner) HasG() bool`

HasG returns a boolean if a field has been set.

### GetV

`func (o *OptionMarkPriceResponseInner) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *OptionMarkPriceResponseInner) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *OptionMarkPriceResponseInner) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *OptionMarkPriceResponseInner) HasV() bool`

HasV returns a boolean if a field has been set.


[[Back to README]](../README.md)


