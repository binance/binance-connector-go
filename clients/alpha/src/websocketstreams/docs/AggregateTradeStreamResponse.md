# AggregateTradeStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event type | [optional] 
**E** | Pointer to **int64** | Event time (ms) | [optional] 
**T** | Pointer to **int64** | Trade time (ms) | [optional] 
**A** | Pointer to **int64** | Aggregated trade ID | [optional] 
**F** | Pointer to **int64** | First trade ID in the aggregation | [optional] 
**L** | Pointer to **int64** | Last trade ID in the aggregation | [optional] 
**M** | Pointer to **bool** | Is the buyer the market maker | [optional] 
**P** | Pointer to **string** | Price | [optional] 
**Q** | Pointer to **string** | Quantity | [optional] 
**S** | Pointer to **string** | Symbol | [optional] 

## Methods

### NewAggregateTradeStreamResponse

`func NewAggregateTradeStreamResponse() *AggregateTradeStreamResponse`

NewAggregateTradeStreamResponse instantiates a new AggregateTradeStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateTradeStreamResponseWithDefaults

`func NewAggregateTradeStreamResponseWithDefaults() *AggregateTradeStreamResponse`

NewAggregateTradeStreamResponseWithDefaults instantiates a new AggregateTradeStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *AggregateTradeStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AggregateTradeStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AggregateTradeStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *AggregateTradeStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetE

`func (o *AggregateTradeStreamResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AggregateTradeStreamResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AggregateTradeStreamResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *AggregateTradeStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetT

`func (o *AggregateTradeStreamResponse) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AggregateTradeStreamResponse) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AggregateTradeStreamResponse) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *AggregateTradeStreamResponse) HasT() bool`

HasT returns a boolean if a field has been set.

### GetA

`func (o *AggregateTradeStreamResponse) GetA() int64`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *AggregateTradeStreamResponse) GetAOk() (*int64, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *AggregateTradeStreamResponse) SetA(v int64)`

SetA sets A field to given value.

### HasA

`func (o *AggregateTradeStreamResponse) HasA() bool`

HasA returns a boolean if a field has been set.

### GetF

`func (o *AggregateTradeStreamResponse) GetF() int64`

GetF returns the F field if non-nil, zero value otherwise.

### GetFOk

`func (o *AggregateTradeStreamResponse) GetFOk() (*int64, bool)`

GetFOk returns a tuple with the F field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF

`func (o *AggregateTradeStreamResponse) SetF(v int64)`

SetF sets F field to given value.

### HasF

`func (o *AggregateTradeStreamResponse) HasF() bool`

HasF returns a boolean if a field has been set.

### GetL

`func (o *AggregateTradeStreamResponse) GetL() int64`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *AggregateTradeStreamResponse) GetLOk() (*int64, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *AggregateTradeStreamResponse) SetL(v int64)`

SetL sets L field to given value.

### HasL

`func (o *AggregateTradeStreamResponse) HasL() bool`

HasL returns a boolean if a field has been set.

### GetM

`func (o *AggregateTradeStreamResponse) GetM() bool`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *AggregateTradeStreamResponse) GetMOk() (*bool, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *AggregateTradeStreamResponse) SetM(v bool)`

SetM sets M field to given value.

### HasM

`func (o *AggregateTradeStreamResponse) HasM() bool`

HasM returns a boolean if a field has been set.

### GetP

`func (o *AggregateTradeStreamResponse) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *AggregateTradeStreamResponse) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *AggregateTradeStreamResponse) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *AggregateTradeStreamResponse) HasP() bool`

HasP returns a boolean if a field has been set.

### GetQ

`func (o *AggregateTradeStreamResponse) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *AggregateTradeStreamResponse) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *AggregateTradeStreamResponse) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *AggregateTradeStreamResponse) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetS

`func (o *AggregateTradeStreamResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *AggregateTradeStreamResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *AggregateTradeStreamResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *AggregateTradeStreamResponse) HasS() bool`

HasS returns a boolean if a field has been set.


[[Back to README]](../README.md)


