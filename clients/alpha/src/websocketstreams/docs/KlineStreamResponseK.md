# KlineStreamResponseK

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **int64** | startTime | [optional] 
**T** | Pointer to **int64** | endTime | [optional] 
**S** | Pointer to **string** | symbol | [optional] 
**I** | Pointer to **string** | interval | [optional] 
**F** | Pointer to **int64** | firstTradeId | [optional] 
**L** | Pointer to **int64** | lastTradeId | [optional] 
**O** | Pointer to **string** | openPrice | [optional] 
**C** | Pointer to **string** | closePrice | [optional] 
**H** | Pointer to **string** | highPrice | [optional] 
**L** | Pointer to **string** | lowPrice | [optional] 
**V** | Pointer to **string** | volume | [optional] 
**N** | Pointer to **int64** | tradeNum | [optional] 
**X** | Pointer to **bool** | klineComplete | [optional] 
**Q** | Pointer to **string** | quoteAssetVolume | [optional] 
**V** | Pointer to **string** | takerBuyBaseAssetVolume | [optional] 
**Q** | Pointer to **string** | takerBuyQuoteAssetVolume | [optional] 
**B** | Pointer to **string** | ignore | [optional] 

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

### GetF

`func (o *KlineStreamResponseK) GetF() int64`

GetF returns the F field if non-nil, zero value otherwise.

### GetFOk

`func (o *KlineStreamResponseK) GetFOk() (*int64, bool)`

GetFOk returns a tuple with the F field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF

`func (o *KlineStreamResponseK) SetF(v int64)`

SetF sets F field to given value.

### HasF

`func (o *KlineStreamResponseK) HasF() bool`

HasF returns a boolean if a field has been set.

### GetL

`func (o *KlineStreamResponseK) GetL() int64`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *KlineStreamResponseK) GetLOk() (*int64, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *KlineStreamResponseK) SetL(v int64)`

SetL sets L field to given value.

### HasL

`func (o *KlineStreamResponseK) HasL() bool`

HasL returns a boolean if a field has been set.

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

`func (o *KlineStreamResponseK) GetN() int64`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *KlineStreamResponseK) GetNOk() (*int64, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *KlineStreamResponseK) SetN(v int64)`

SetN sets N field to given value.

### HasN

`func (o *KlineStreamResponseK) HasN() bool`

HasN returns a boolean if a field has been set.

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

### GetQ

`func (o *KlineStreamResponseK) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *KlineStreamResponseK) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *KlineStreamResponseK) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *KlineStreamResponseK) HasQ() bool`

HasQ returns a boolean if a field has been set.

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

### GetQ

`func (o *KlineStreamResponseK) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *KlineStreamResponseK) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *KlineStreamResponseK) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *KlineStreamResponseK) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetB

`func (o *KlineStreamResponseK) GetB() string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *KlineStreamResponseK) GetBOk() (*string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *KlineStreamResponseK) SetB(v string)`

SetB sets B field to given value.

### HasB

`func (o *KlineStreamResponseK) HasB() bool`

HasB returns a boolean if a field has been set.


[[Back to README]](../README.md)


