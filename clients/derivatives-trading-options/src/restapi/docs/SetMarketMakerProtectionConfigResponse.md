# SetMarketMakerProtectionConfigResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**UnderlyingId** | Pointer to **int64** |  | [optional] 
**Underlying** | Pointer to **string** |  | [optional] 
**WindowTimeInMilliseconds** | Pointer to **int64** |  | [optional] 
**FrozenTimeInMilliseconds** | Pointer to **int64** |  | [optional] 
**QtyLimit** | Pointer to **string** |  | [optional] 
**DeltaLimit** | Pointer to **string** |  | [optional] 
**LastTriggerTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewSetMarketMakerProtectionConfigResponse

`func NewSetMarketMakerProtectionConfigResponse() *SetMarketMakerProtectionConfigResponse`

NewSetMarketMakerProtectionConfigResponse instantiates a new SetMarketMakerProtectionConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetMarketMakerProtectionConfigResponseWithDefaults

`func NewSetMarketMakerProtectionConfigResponseWithDefaults() *SetMarketMakerProtectionConfigResponse`

NewSetMarketMakerProtectionConfigResponseWithDefaults instantiates a new SetMarketMakerProtectionConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderlyingId

`func (o *SetMarketMakerProtectionConfigResponse) GetUnderlyingId() int64`

GetUnderlyingId returns the UnderlyingId field if non-nil, zero value otherwise.

### GetUnderlyingIdOk

`func (o *SetMarketMakerProtectionConfigResponse) GetUnderlyingIdOk() (*int64, bool)`

GetUnderlyingIdOk returns a tuple with the UnderlyingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingId

`func (o *SetMarketMakerProtectionConfigResponse) SetUnderlyingId(v int64)`

SetUnderlyingId sets UnderlyingId field to given value.

### HasUnderlyingId

`func (o *SetMarketMakerProtectionConfigResponse) HasUnderlyingId() bool`

HasUnderlyingId returns a boolean if a field has been set.

### GetUnderlying

`func (o *SetMarketMakerProtectionConfigResponse) GetUnderlying() string`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *SetMarketMakerProtectionConfigResponse) GetUnderlyingOk() (*string, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *SetMarketMakerProtectionConfigResponse) SetUnderlying(v string)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *SetMarketMakerProtectionConfigResponse) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.

### GetWindowTimeInMilliseconds

`func (o *SetMarketMakerProtectionConfigResponse) GetWindowTimeInMilliseconds() int64`

GetWindowTimeInMilliseconds returns the WindowTimeInMilliseconds field if non-nil, zero value otherwise.

### GetWindowTimeInMillisecondsOk

`func (o *SetMarketMakerProtectionConfigResponse) GetWindowTimeInMillisecondsOk() (*int64, bool)`

GetWindowTimeInMillisecondsOk returns a tuple with the WindowTimeInMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowTimeInMilliseconds

`func (o *SetMarketMakerProtectionConfigResponse) SetWindowTimeInMilliseconds(v int64)`

SetWindowTimeInMilliseconds sets WindowTimeInMilliseconds field to given value.

### HasWindowTimeInMilliseconds

`func (o *SetMarketMakerProtectionConfigResponse) HasWindowTimeInMilliseconds() bool`

HasWindowTimeInMilliseconds returns a boolean if a field has been set.

### GetFrozenTimeInMilliseconds

`func (o *SetMarketMakerProtectionConfigResponse) GetFrozenTimeInMilliseconds() int64`

GetFrozenTimeInMilliseconds returns the FrozenTimeInMilliseconds field if non-nil, zero value otherwise.

### GetFrozenTimeInMillisecondsOk

`func (o *SetMarketMakerProtectionConfigResponse) GetFrozenTimeInMillisecondsOk() (*int64, bool)`

GetFrozenTimeInMillisecondsOk returns a tuple with the FrozenTimeInMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenTimeInMilliseconds

`func (o *SetMarketMakerProtectionConfigResponse) SetFrozenTimeInMilliseconds(v int64)`

SetFrozenTimeInMilliseconds sets FrozenTimeInMilliseconds field to given value.

### HasFrozenTimeInMilliseconds

`func (o *SetMarketMakerProtectionConfigResponse) HasFrozenTimeInMilliseconds() bool`

HasFrozenTimeInMilliseconds returns a boolean if a field has been set.

### GetQtyLimit

`func (o *SetMarketMakerProtectionConfigResponse) GetQtyLimit() string`

GetQtyLimit returns the QtyLimit field if non-nil, zero value otherwise.

### GetQtyLimitOk

`func (o *SetMarketMakerProtectionConfigResponse) GetQtyLimitOk() (*string, bool)`

GetQtyLimitOk returns a tuple with the QtyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyLimit

`func (o *SetMarketMakerProtectionConfigResponse) SetQtyLimit(v string)`

SetQtyLimit sets QtyLimit field to given value.

### HasQtyLimit

`func (o *SetMarketMakerProtectionConfigResponse) HasQtyLimit() bool`

HasQtyLimit returns a boolean if a field has been set.

### GetDeltaLimit

`func (o *SetMarketMakerProtectionConfigResponse) GetDeltaLimit() string`

GetDeltaLimit returns the DeltaLimit field if non-nil, zero value otherwise.

### GetDeltaLimitOk

`func (o *SetMarketMakerProtectionConfigResponse) GetDeltaLimitOk() (*string, bool)`

GetDeltaLimitOk returns a tuple with the DeltaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaLimit

`func (o *SetMarketMakerProtectionConfigResponse) SetDeltaLimit(v string)`

SetDeltaLimit sets DeltaLimit field to given value.

### HasDeltaLimit

`func (o *SetMarketMakerProtectionConfigResponse) HasDeltaLimit() bool`

HasDeltaLimit returns a boolean if a field has been set.

### GetLastTriggerTime

`func (o *SetMarketMakerProtectionConfigResponse) GetLastTriggerTime() int64`

GetLastTriggerTime returns the LastTriggerTime field if non-nil, zero value otherwise.

### GetLastTriggerTimeOk

`func (o *SetMarketMakerProtectionConfigResponse) GetLastTriggerTimeOk() (*int64, bool)`

GetLastTriggerTimeOk returns a tuple with the LastTriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggerTime

`func (o *SetMarketMakerProtectionConfigResponse) SetLastTriggerTime(v int64)`

SetLastTriggerTime sets LastTriggerTime field to given value.

### HasLastTriggerTime

`func (o *SetMarketMakerProtectionConfigResponse) HasLastTriggerTime() bool`

HasLastTriggerTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


