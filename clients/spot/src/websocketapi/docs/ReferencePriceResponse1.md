# ReferencePriceResponse1

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**ReferencePriceResponse1Result**](ReferencePriceResponse1Result.md) |  | [optional] 
**RateLimits** | Pointer to [**[]AvgPriceResponseRateLimitsInner**](AvgPriceResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewReferencePriceResponse1

`func NewReferencePriceResponse1() *ReferencePriceResponse1`

NewReferencePriceResponse1 instantiates a new ReferencePriceResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencePriceResponse1WithDefaults

`func NewReferencePriceResponse1WithDefaults() *ReferencePriceResponse1`

NewReferencePriceResponse1WithDefaults instantiates a new ReferencePriceResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReferencePriceResponse1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReferencePriceResponse1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReferencePriceResponse1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReferencePriceResponse1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ReferencePriceResponse1) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReferencePriceResponse1) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReferencePriceResponse1) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReferencePriceResponse1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *ReferencePriceResponse1) GetResult() ReferencePriceResponse1Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ReferencePriceResponse1) GetResultOk() (*ReferencePriceResponse1Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ReferencePriceResponse1) SetResult(v ReferencePriceResponse1Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *ReferencePriceResponse1) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *ReferencePriceResponse1) GetRateLimits() []AvgPriceResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *ReferencePriceResponse1) GetRateLimitsOk() (*[]AvgPriceResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *ReferencePriceResponse1) SetRateLimits(v []AvgPriceResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *ReferencePriceResponse1) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


