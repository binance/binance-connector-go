# ReferencePriceCalculationResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**ReferencePriceCalculationResponseResult**](ReferencePriceCalculationResponseResult.md) |  | [optional] 

## Methods

### NewReferencePriceCalculationResponse

`func NewReferencePriceCalculationResponse() *ReferencePriceCalculationResponse`

NewReferencePriceCalculationResponse instantiates a new ReferencePriceCalculationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencePriceCalculationResponseWithDefaults

`func NewReferencePriceCalculationResponseWithDefaults() *ReferencePriceCalculationResponse`

NewReferencePriceCalculationResponseWithDefaults instantiates a new ReferencePriceCalculationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReferencePriceCalculationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReferencePriceCalculationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReferencePriceCalculationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReferencePriceCalculationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ReferencePriceCalculationResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReferencePriceCalculationResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReferencePriceCalculationResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReferencePriceCalculationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *ReferencePriceCalculationResponse) GetResult() ReferencePriceCalculationResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ReferencePriceCalculationResponse) GetResultOk() (*ReferencePriceCalculationResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ReferencePriceCalculationResponse) SetResult(v ReferencePriceCalculationResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ReferencePriceCalculationResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to README]](../README.md)


