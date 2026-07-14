# AllMarketMiniTickersStreamResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event type | [optional] 
**E** | Pointer to **int64** | Event time | [optional] 
**S** | Pointer to **string** | Symbol | [optional] 
**C** | Pointer to **string** | Close price | [optional] 
**O** | Pointer to **string** | Open price | [optional] 
**H** | Pointer to **string** | High price | [optional] 
**L** | Pointer to **string** | Low price | [optional] 
**V** | Pointer to **string** | Total traded base asset volume | [optional] 
**Q** | Pointer to **string** | Total traded quote asset volume | [optional] 
**Ps** | Pointer to **string** | (After CM migration) Pair symbol | [optional] 
**St** | Pointer to **int32** | (After CM migration) Symbol type: 1 &#x3D; UM, 2 &#x3D; CM | [optional] 

## Methods

### NewAllMarketMiniTickersStreamResponseInner

`func NewAllMarketMiniTickersStreamResponseInner() *AllMarketMiniTickersStreamResponseInner`

NewAllMarketMiniTickersStreamResponseInner instantiates a new AllMarketMiniTickersStreamResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllMarketMiniTickersStreamResponseInnerWithDefaults

`func NewAllMarketMiniTickersStreamResponseInnerWithDefaults() *AllMarketMiniTickersStreamResponseInner`

NewAllMarketMiniTickersStreamResponseInnerWithDefaults instantiates a new AllMarketMiniTickersStreamResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *AllMarketMiniTickersStreamResponseInner) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AllMarketMiniTickersStreamResponseInner) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AllMarketMiniTickersStreamResponseInner) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *AllMarketMiniTickersStreamResponseInner) HasE() bool`

HasE returns a boolean if a field has been set.

### GetE

`func (o *AllMarketMiniTickersStreamResponseInner) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AllMarketMiniTickersStreamResponseInner) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AllMarketMiniTickersStreamResponseInner) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *AllMarketMiniTickersStreamResponseInner) HasE() bool`

HasE returns a boolean if a field has been set.

### GetS

`func (o *AllMarketMiniTickersStreamResponseInner) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *AllMarketMiniTickersStreamResponseInner) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *AllMarketMiniTickersStreamResponseInner) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *AllMarketMiniTickersStreamResponseInner) HasS() bool`

HasS returns a boolean if a field has been set.

### GetC

`func (o *AllMarketMiniTickersStreamResponseInner) GetC() string`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *AllMarketMiniTickersStreamResponseInner) GetCOk() (*string, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *AllMarketMiniTickersStreamResponseInner) SetC(v string)`

SetC sets C field to given value.

### HasC

`func (o *AllMarketMiniTickersStreamResponseInner) HasC() bool`

HasC returns a boolean if a field has been set.

### GetO

`func (o *AllMarketMiniTickersStreamResponseInner) GetO() string`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *AllMarketMiniTickersStreamResponseInner) GetOOk() (*string, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *AllMarketMiniTickersStreamResponseInner) SetO(v string)`

SetO sets O field to given value.

### HasO

`func (o *AllMarketMiniTickersStreamResponseInner) HasO() bool`

HasO returns a boolean if a field has been set.

### GetH

`func (o *AllMarketMiniTickersStreamResponseInner) GetH() string`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *AllMarketMiniTickersStreamResponseInner) GetHOk() (*string, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *AllMarketMiniTickersStreamResponseInner) SetH(v string)`

SetH sets H field to given value.

### HasH

`func (o *AllMarketMiniTickersStreamResponseInner) HasH() bool`

HasH returns a boolean if a field has been set.

### GetL

`func (o *AllMarketMiniTickersStreamResponseInner) GetL() string`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *AllMarketMiniTickersStreamResponseInner) GetLOk() (*string, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *AllMarketMiniTickersStreamResponseInner) SetL(v string)`

SetL sets L field to given value.

### HasL

`func (o *AllMarketMiniTickersStreamResponseInner) HasL() bool`

HasL returns a boolean if a field has been set.

### GetV

`func (o *AllMarketMiniTickersStreamResponseInner) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *AllMarketMiniTickersStreamResponseInner) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *AllMarketMiniTickersStreamResponseInner) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *AllMarketMiniTickersStreamResponseInner) HasV() bool`

HasV returns a boolean if a field has been set.

### GetQ

`func (o *AllMarketMiniTickersStreamResponseInner) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *AllMarketMiniTickersStreamResponseInner) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *AllMarketMiniTickersStreamResponseInner) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *AllMarketMiniTickersStreamResponseInner) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetPs

`func (o *AllMarketMiniTickersStreamResponseInner) GetPs() string`

GetPs returns the Ps field if non-nil, zero value otherwise.

### GetPsOk

`func (o *AllMarketMiniTickersStreamResponseInner) GetPsOk() (*string, bool)`

GetPsOk returns a tuple with the Ps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPs

`func (o *AllMarketMiniTickersStreamResponseInner) SetPs(v string)`

SetPs sets Ps field to given value.

### HasPs

`func (o *AllMarketMiniTickersStreamResponseInner) HasPs() bool`

HasPs returns a boolean if a field has been set.

### GetSt

`func (o *AllMarketMiniTickersStreamResponseInner) GetSt() int32`

GetSt returns the St field if non-nil, zero value otherwise.

### GetStOk

`func (o *AllMarketMiniTickersStreamResponseInner) GetStOk() (*int32, bool)`

GetStOk returns a tuple with the St field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSt

`func (o *AllMarketMiniTickersStreamResponseInner) SetSt(v int32)`

SetSt sets St field to given value.

### HasSt

`func (o *AllMarketMiniTickersStreamResponseInner) HasSt() bool`

HasSt returns a boolean if a field has been set.


[[Back to README]](../README.md)


