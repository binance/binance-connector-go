# OrderAmendmentsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]OrderAmendmentsResponseResultInner**](OrderAmendmentsResponseResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewOrderAmendmentsResponse

`func NewOrderAmendmentsResponse() *OrderAmendmentsResponse`

NewOrderAmendmentsResponse instantiates a new OrderAmendmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmendmentsResponseWithDefaults

`func NewOrderAmendmentsResponseWithDefaults() *OrderAmendmentsResponse`

NewOrderAmendmentsResponseWithDefaults instantiates a new OrderAmendmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderAmendmentsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderAmendmentsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderAmendmentsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderAmendmentsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *OrderAmendmentsResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderAmendmentsResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderAmendmentsResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderAmendmentsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *OrderAmendmentsResponse) GetResult() []OrderAmendmentsResponseResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OrderAmendmentsResponse) GetResultOk() (*[]OrderAmendmentsResponseResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OrderAmendmentsResponse) SetResult(v []OrderAmendmentsResponseResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *OrderAmendmentsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *OrderAmendmentsResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *OrderAmendmentsResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *OrderAmendmentsResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *OrderAmendmentsResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


