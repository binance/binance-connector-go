# AllMarketTickersStreamsResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event type | [optional] 
**E** | Pointer to **int64** | Event time | [optional] 
**S** | Pointer to **string** | Symbol | [optional] 
**P** | Pointer to **string** | Price change | [optional] 
**P** | Pointer to **string** | Price change percent | [optional] 
**W** | Pointer to **string** | Weighted average price | [optional] 
**C** | Pointer to **string** | Last price | [optional] 
**Q** | Pointer to **string** | Last quantity | [optional] 
**O** | Pointer to **string** | Open price | [optional] 
**H** | Pointer to **string** | High price | [optional] 
**L** | Pointer to **string** | Low price | [optional] 
**V** | Pointer to **string** | Total traded base asset volume | [optional] 
**Q** | Pointer to **string** | Total traded quote asset volume | [optional] 
**O** | Pointer to **int64** | Statistics open time | [optional] 
**C** | Pointer to **int64** | Statistics close time | [optional] 
**F** | Pointer to **int64** | First trade ID | [optional] 
**L** | Pointer to **int64** | Last trade Id | [optional] 
**N** | Pointer to **int64** | Total number of trades | [optional] 
**Ps** | Pointer to **string** | (After CM migration) Pair symbol | [optional] 
**St** | Pointer to **int32** | (After CM migration) Symbol type: 1 &#x3D; UM, 2 &#x3D; CM | [optional] 

## Methods

### NewAllMarketTickersStreamsResponseInner

`func NewAllMarketTickersStreamsResponseInner() *AllMarketTickersStreamsResponseInner`

NewAllMarketTickersStreamsResponseInner instantiates a new AllMarketTickersStreamsResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllMarketTickersStreamsResponseInnerWithDefaults

`func NewAllMarketTickersStreamsResponseInnerWithDefaults() *AllMarketTickersStreamsResponseInner`

NewAllMarketTickersStreamsResponseInnerWithDefaults instantiates a new AllMarketTickersStreamsResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *AllMarketTickersStreamsResponseInner) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AllMarketTickersStreamsResponseInner) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AllMarketTickersStreamsResponseInner) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *AllMarketTickersStreamsResponseInner) HasE() bool`

HasE returns a boolean if a field has been set.

### GetE

`func (o *AllMarketTickersStreamsResponseInner) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AllMarketTickersStreamsResponseInner) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AllMarketTickersStreamsResponseInner) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *AllMarketTickersStreamsResponseInner) HasE() bool`

HasE returns a boolean if a field has been set.

### GetS

`func (o *AllMarketTickersStreamsResponseInner) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *AllMarketTickersStreamsResponseInner) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *AllMarketTickersStreamsResponseInner) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *AllMarketTickersStreamsResponseInner) HasS() bool`

HasS returns a boolean if a field has been set.

### GetP

`func (o *AllMarketTickersStreamsResponseInner) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *AllMarketTickersStreamsResponseInner) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *AllMarketTickersStreamsResponseInner) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *AllMarketTickersStreamsResponseInner) HasP() bool`

HasP returns a boolean if a field has been set.

### GetP

`func (o *AllMarketTickersStreamsResponseInner) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *AllMarketTickersStreamsResponseInner) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *AllMarketTickersStreamsResponseInner) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *AllMarketTickersStreamsResponseInner) HasP() bool`

HasP returns a boolean if a field has been set.

### GetW

`func (o *AllMarketTickersStreamsResponseInner) GetW() string`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *AllMarketTickersStreamsResponseInner) GetWOk() (*string, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *AllMarketTickersStreamsResponseInner) SetW(v string)`

SetW sets W field to given value.

### HasW

`func (o *AllMarketTickersStreamsResponseInner) HasW() bool`

HasW returns a boolean if a field has been set.

### GetC

`func (o *AllMarketTickersStreamsResponseInner) GetC() string`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *AllMarketTickersStreamsResponseInner) GetCOk() (*string, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *AllMarketTickersStreamsResponseInner) SetC(v string)`

SetC sets C field to given value.

### HasC

`func (o *AllMarketTickersStreamsResponseInner) HasC() bool`

HasC returns a boolean if a field has been set.

### GetQ

`func (o *AllMarketTickersStreamsResponseInner) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *AllMarketTickersStreamsResponseInner) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *AllMarketTickersStreamsResponseInner) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *AllMarketTickersStreamsResponseInner) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetO

`func (o *AllMarketTickersStreamsResponseInner) GetO() string`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *AllMarketTickersStreamsResponseInner) GetOOk() (*string, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *AllMarketTickersStreamsResponseInner) SetO(v string)`

SetO sets O field to given value.

### HasO

`func (o *AllMarketTickersStreamsResponseInner) HasO() bool`

HasO returns a boolean if a field has been set.

### GetH

`func (o *AllMarketTickersStreamsResponseInner) GetH() string`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *AllMarketTickersStreamsResponseInner) GetHOk() (*string, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *AllMarketTickersStreamsResponseInner) SetH(v string)`

SetH sets H field to given value.

### HasH

`func (o *AllMarketTickersStreamsResponseInner) HasH() bool`

HasH returns a boolean if a field has been set.

### GetL

`func (o *AllMarketTickersStreamsResponseInner) GetL() string`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *AllMarketTickersStreamsResponseInner) GetLOk() (*string, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *AllMarketTickersStreamsResponseInner) SetL(v string)`

SetL sets L field to given value.

### HasL

`func (o *AllMarketTickersStreamsResponseInner) HasL() bool`

HasL returns a boolean if a field has been set.

### GetV

`func (o *AllMarketTickersStreamsResponseInner) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *AllMarketTickersStreamsResponseInner) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *AllMarketTickersStreamsResponseInner) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *AllMarketTickersStreamsResponseInner) HasV() bool`

HasV returns a boolean if a field has been set.

### GetQ

`func (o *AllMarketTickersStreamsResponseInner) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *AllMarketTickersStreamsResponseInner) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *AllMarketTickersStreamsResponseInner) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *AllMarketTickersStreamsResponseInner) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetO

`func (o *AllMarketTickersStreamsResponseInner) GetO() int64`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *AllMarketTickersStreamsResponseInner) GetOOk() (*int64, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *AllMarketTickersStreamsResponseInner) SetO(v int64)`

SetO sets O field to given value.

### HasO

`func (o *AllMarketTickersStreamsResponseInner) HasO() bool`

HasO returns a boolean if a field has been set.

### GetC

`func (o *AllMarketTickersStreamsResponseInner) GetC() int64`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *AllMarketTickersStreamsResponseInner) GetCOk() (*int64, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *AllMarketTickersStreamsResponseInner) SetC(v int64)`

SetC sets C field to given value.

### HasC

`func (o *AllMarketTickersStreamsResponseInner) HasC() bool`

HasC returns a boolean if a field has been set.

### GetF

`func (o *AllMarketTickersStreamsResponseInner) GetF() int64`

GetF returns the F field if non-nil, zero value otherwise.

### GetFOk

`func (o *AllMarketTickersStreamsResponseInner) GetFOk() (*int64, bool)`

GetFOk returns a tuple with the F field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF

`func (o *AllMarketTickersStreamsResponseInner) SetF(v int64)`

SetF sets F field to given value.

### HasF

`func (o *AllMarketTickersStreamsResponseInner) HasF() bool`

HasF returns a boolean if a field has been set.

### GetL

`func (o *AllMarketTickersStreamsResponseInner) GetL() int64`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *AllMarketTickersStreamsResponseInner) GetLOk() (*int64, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *AllMarketTickersStreamsResponseInner) SetL(v int64)`

SetL sets L field to given value.

### HasL

`func (o *AllMarketTickersStreamsResponseInner) HasL() bool`

HasL returns a boolean if a field has been set.

### GetN

`func (o *AllMarketTickersStreamsResponseInner) GetN() int64`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AllMarketTickersStreamsResponseInner) GetNOk() (*int64, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AllMarketTickersStreamsResponseInner) SetN(v int64)`

SetN sets N field to given value.

### HasN

`func (o *AllMarketTickersStreamsResponseInner) HasN() bool`

HasN returns a boolean if a field has been set.

### GetPs

`func (o *AllMarketTickersStreamsResponseInner) GetPs() string`

GetPs returns the Ps field if non-nil, zero value otherwise.

### GetPsOk

`func (o *AllMarketTickersStreamsResponseInner) GetPsOk() (*string, bool)`

GetPsOk returns a tuple with the Ps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPs

`func (o *AllMarketTickersStreamsResponseInner) SetPs(v string)`

SetPs sets Ps field to given value.

### HasPs

`func (o *AllMarketTickersStreamsResponseInner) HasPs() bool`

HasPs returns a boolean if a field has been set.

### GetSt

`func (o *AllMarketTickersStreamsResponseInner) GetSt() int32`

GetSt returns the St field if non-nil, zero value otherwise.

### GetStOk

`func (o *AllMarketTickersStreamsResponseInner) GetStOk() (*int32, bool)`

GetStOk returns a tuple with the St field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSt

`func (o *AllMarketTickersStreamsResponseInner) SetSt(v int32)`

SetSt sets St field to given value.

### HasSt

`func (o *AllMarketTickersStreamsResponseInner) HasSt() bool`

HasSt returns a boolean if a field has been set.


[[Back to README]](../README.md)


