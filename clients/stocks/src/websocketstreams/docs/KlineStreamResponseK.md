# KlineStreamResponseK

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **int64** | Open time (epoch milliseconds). | [optional] 
**Ct** | Pointer to **int64** | Close time (epoch milliseconds). For fixed intervals: &#x60;openTime + duration − 1ms&#x60;. For &#x60;1M&#x60;: last day of month at 23:59:59.999 UTC. | [optional] 
**S** | Pointer to **string** | Symbol (UPPERCASE ticker). | [optional] 
**I** | Pointer to **string** | Interval — &#x60;\&quot;5m\&quot;&#x60; / &#x60;\&quot;1h\&quot;&#x60; / &#x60;\&quot;1d\&quot;&#x60; / &#x60;\&quot;1w\&quot;&#x60; / &#x60;\&quot;1M\&quot;&#x60;. | [optional] 
**O** | Pointer to **string** | Open price. | [optional] 
**C** | Pointer to **string** | Close price (latest print within the interval). | [optional] 
**H** | Pointer to **string** | High price. | [optional] 
**L** | Pointer to **string** | Low price. | [optional] 
**V** | Pointer to **string** | Volume (shares), default &#x60;\&quot;0\&quot;&#x60;. | [optional] 
**N** | Pointer to **NullableInt32** | Trade count; absent when upstream did not populate. | [optional] 
**X** | Pointer to **bool** | Candle is closed — &#x60;true&#x60; once the interval has ended. | [optional] 
**Vw** | Pointer to **NullableString** | VWAP (volume-weighted average price); absent when upstream did not populate. | [optional] 

## Methods

### NewKlineStreamResponseK

`func NewKlineStreamResponseK() *KlineStreamResponseK`

NewKlineStreamResponseK instantiates a new KlineStreamResponseK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlineStreamResponseKWithDefaults

`func NewKlineStreamResponseKWithDefaults() *KlineStreamResponseK`

NewKlineStreamResponseKWithDefaults instantiates a new KlineStreamResponseK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *KlineStreamResponseK) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *KlineStreamResponseK) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *KlineStreamResponseK) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *KlineStreamResponseK) HasT() bool`

HasT returns a boolean if a field has been set.

### GetCt

`func (o *KlineStreamResponseK) GetCt() int64`

GetCt returns the Ct field if non-nil, zero value otherwise.

### GetCtOk

`func (o *KlineStreamResponseK) GetCtOk() (*int64, bool)`

GetCtOk returns a tuple with the Ct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCt

`func (o *KlineStreamResponseK) SetCt(v int64)`

SetCt sets Ct field to given value.

### HasCt

`func (o *KlineStreamResponseK) HasCt() bool`

HasCt returns a boolean if a field has been set.

### GetS

`func (o *KlineStreamResponseK) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *KlineStreamResponseK) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *KlineStreamResponseK) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *KlineStreamResponseK) HasS() bool`

HasS returns a boolean if a field has been set.

### GetI

`func (o *KlineStreamResponseK) GetI() string`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *KlineStreamResponseK) GetIOk() (*string, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *KlineStreamResponseK) SetI(v string)`

SetI sets I field to given value.

### HasI

`func (o *KlineStreamResponseK) HasI() bool`

HasI returns a boolean if a field has been set.

### GetO

`func (o *KlineStreamResponseK) GetO() string`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *KlineStreamResponseK) GetOOk() (*string, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *KlineStreamResponseK) SetO(v string)`

SetO sets O field to given value.

### HasO

`func (o *KlineStreamResponseK) HasO() bool`

HasO returns a boolean if a field has been set.

### GetC

`func (o *KlineStreamResponseK) GetC() string`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *KlineStreamResponseK) GetCOk() (*string, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *KlineStreamResponseK) SetC(v string)`

SetC sets C field to given value.

### HasC

`func (o *KlineStreamResponseK) HasC() bool`

HasC returns a boolean if a field has been set.

### GetH

`func (o *KlineStreamResponseK) GetH() string`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *KlineStreamResponseK) GetHOk() (*string, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *KlineStreamResponseK) SetH(v string)`

SetH sets H field to given value.

### HasH

`func (o *KlineStreamResponseK) HasH() bool`

HasH returns a boolean if a field has been set.

### GetL

`func (o *KlineStreamResponseK) GetL() string`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *KlineStreamResponseK) GetLOk() (*string, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *KlineStreamResponseK) SetL(v string)`

SetL sets L field to given value.

### HasL

`func (o *KlineStreamResponseK) HasL() bool`

HasL returns a boolean if a field has been set.

### GetV

`func (o *KlineStreamResponseK) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *KlineStreamResponseK) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *KlineStreamResponseK) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *KlineStreamResponseK) HasV() bool`

HasV returns a boolean if a field has been set.

### GetN

`func (o *KlineStreamResponseK) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *KlineStreamResponseK) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *KlineStreamResponseK) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *KlineStreamResponseK) HasN() bool`

HasN returns a boolean if a field has been set.

### SetNNil

`func (o *KlineStreamResponseK) SetNNil(b bool)`

 SetNNil sets the value for N to be an explicit nil

### UnsetN
`func (o *KlineStreamResponseK) UnsetN()`

UnsetN ensures that no value is present for N, not even an explicit nil
### GetX

`func (o *KlineStreamResponseK) GetX() bool`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *KlineStreamResponseK) GetXOk() (*bool, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *KlineStreamResponseK) SetX(v bool)`

SetX sets X field to given value.

### HasX

`func (o *KlineStreamResponseK) HasX() bool`

HasX returns a boolean if a field has been set.

### GetVw

`func (o *KlineStreamResponseK) GetVw() string`

GetVw returns the Vw field if non-nil, zero value otherwise.

### GetVwOk

`func (o *KlineStreamResponseK) GetVwOk() (*string, bool)`

GetVwOk returns a tuple with the Vw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVw

`func (o *KlineStreamResponseK) SetVw(v string)`

SetVw sets Vw field to given value.

### HasVw

`func (o *KlineStreamResponseK) HasVw() bool`

HasVw returns a boolean if a field has been set.

### SetVwNil

`func (o *KlineStreamResponseK) SetVwNil(b bool)`

 SetVwNil sets the value for Vw to be an explicit nil

### UnsetVw
`func (o *KlineStreamResponseK) UnsetVw()`

UnsetVw ensures that no value is present for Vw, not even an explicit nil

[[Back to README]](../README.md)


