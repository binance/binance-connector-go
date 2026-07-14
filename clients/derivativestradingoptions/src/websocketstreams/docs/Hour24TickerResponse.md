# Hour24TickerResponse

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
**V** | Pointer to **string** | Trading volume(in contracts) | [optional] 
**Q** | Pointer to **string** | trade amount(in quote asset) | [optional] 
**O** | Pointer to **int64** | Statistics open time | [optional] 
**C** | Pointer to **int64** | Statistics close time | [optional] 
**F** | Pointer to **int64** | First trade ID | [optional] 
**L** | Pointer to **int64** | Last trade Id | [optional] 
**N** | Pointer to **int64** | Total number of trade | [optional] 

## Methods

### NewHour24TickerResponse

`func NewHour24TickerResponse() *Hour24TickerResponse`

NewHour24TickerResponse instantiates a new Hour24TickerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHour24TickerResponseWithDefaults

`func NewHour24TickerResponseWithDefaults() *Hour24TickerResponse`

NewHour24TickerResponseWithDefaults instantiates a new Hour24TickerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *Hour24TickerResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *Hour24TickerResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *Hour24TickerResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *Hour24TickerResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetE

`func (o *Hour24TickerResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *Hour24TickerResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *Hour24TickerResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *Hour24TickerResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetS

`func (o *Hour24TickerResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *Hour24TickerResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *Hour24TickerResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *Hour24TickerResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetP

`func (o *Hour24TickerResponse) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *Hour24TickerResponse) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *Hour24TickerResponse) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *Hour24TickerResponse) HasP() bool`

HasP returns a boolean if a field has been set.

### GetP

`func (o *Hour24TickerResponse) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *Hour24TickerResponse) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *Hour24TickerResponse) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *Hour24TickerResponse) HasP() bool`

HasP returns a boolean if a field has been set.

### GetW

`func (o *Hour24TickerResponse) GetW() string`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *Hour24TickerResponse) GetWOk() (*string, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *Hour24TickerResponse) SetW(v string)`

SetW sets W field to given value.

### HasW

`func (o *Hour24TickerResponse) HasW() bool`

HasW returns a boolean if a field has been set.

### GetC

`func (o *Hour24TickerResponse) GetC() string`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *Hour24TickerResponse) GetCOk() (*string, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *Hour24TickerResponse) SetC(v string)`

SetC sets C field to given value.

### HasC

`func (o *Hour24TickerResponse) HasC() bool`

HasC returns a boolean if a field has been set.

### GetQ

`func (o *Hour24TickerResponse) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *Hour24TickerResponse) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *Hour24TickerResponse) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *Hour24TickerResponse) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetO

`func (o *Hour24TickerResponse) GetO() string`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *Hour24TickerResponse) GetOOk() (*string, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *Hour24TickerResponse) SetO(v string)`

SetO sets O field to given value.

### HasO

`func (o *Hour24TickerResponse) HasO() bool`

HasO returns a boolean if a field has been set.

### GetH

`func (o *Hour24TickerResponse) GetH() string`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *Hour24TickerResponse) GetHOk() (*string, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *Hour24TickerResponse) SetH(v string)`

SetH sets H field to given value.

### HasH

`func (o *Hour24TickerResponse) HasH() bool`

HasH returns a boolean if a field has been set.

### GetL

`func (o *Hour24TickerResponse) GetL() string`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *Hour24TickerResponse) GetLOk() (*string, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *Hour24TickerResponse) SetL(v string)`

SetL sets L field to given value.

### HasL

`func (o *Hour24TickerResponse) HasL() bool`

HasL returns a boolean if a field has been set.

### GetV

`func (o *Hour24TickerResponse) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *Hour24TickerResponse) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *Hour24TickerResponse) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *Hour24TickerResponse) HasV() bool`

HasV returns a boolean if a field has been set.

### GetQ

`func (o *Hour24TickerResponse) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *Hour24TickerResponse) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *Hour24TickerResponse) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *Hour24TickerResponse) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetO

`func (o *Hour24TickerResponse) GetO() int64`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *Hour24TickerResponse) GetOOk() (*int64, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *Hour24TickerResponse) SetO(v int64)`

SetO sets O field to given value.

### HasO

`func (o *Hour24TickerResponse) HasO() bool`

HasO returns a boolean if a field has been set.

### GetC

`func (o *Hour24TickerResponse) GetC() int64`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *Hour24TickerResponse) GetCOk() (*int64, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *Hour24TickerResponse) SetC(v int64)`

SetC sets C field to given value.

### HasC

`func (o *Hour24TickerResponse) HasC() bool`

HasC returns a boolean if a field has been set.

### GetF

`func (o *Hour24TickerResponse) GetF() int64`

GetF returns the F field if non-nil, zero value otherwise.

### GetFOk

`func (o *Hour24TickerResponse) GetFOk() (*int64, bool)`

GetFOk returns a tuple with the F field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF

`func (o *Hour24TickerResponse) SetF(v int64)`

SetF sets F field to given value.

### HasF

`func (o *Hour24TickerResponse) HasF() bool`

HasF returns a boolean if a field has been set.

### GetL

`func (o *Hour24TickerResponse) GetL() int64`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *Hour24TickerResponse) GetLOk() (*int64, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *Hour24TickerResponse) SetL(v int64)`

SetL sets L field to given value.

### HasL

`func (o *Hour24TickerResponse) HasL() bool`

HasL returns a boolean if a field has been set.

### GetN

`func (o *Hour24TickerResponse) GetN() int64`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *Hour24TickerResponse) GetNOk() (*int64, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *Hour24TickerResponse) SetN(v int64)`

SetN sets N field to given value.

### HasN

`func (o *Hour24TickerResponse) HasN() bool`

HasN returns a boolean if a field has been set.


[[Back to README]](../README.md)


