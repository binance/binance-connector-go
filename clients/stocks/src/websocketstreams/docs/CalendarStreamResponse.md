# CalendarStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event type, always &#x60;\&quot;calendar\&quot;&#x60;. | [optional] 
**From** | Pointer to **string** | Previous market phase. | [optional] 
**To** | Pointer to **string** | New market phase. | [optional] 
**Ts** | Pointer to **int64** | Event timestamp (epoch milliseconds UTC) when the transition was detected. | [optional] 

## Methods

### NewCalendarStreamResponse

`func NewCalendarStreamResponse() *CalendarStreamResponse`

NewCalendarStreamResponse instantiates a new CalendarStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarStreamResponseWithDefaults

`func NewCalendarStreamResponseWithDefaults() *CalendarStreamResponse`

NewCalendarStreamResponseWithDefaults instantiates a new CalendarStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *CalendarStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *CalendarStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *CalendarStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *CalendarStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetFrom

`func (o *CalendarStreamResponse) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CalendarStreamResponse) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CalendarStreamResponse) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CalendarStreamResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CalendarStreamResponse) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CalendarStreamResponse) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CalendarStreamResponse) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CalendarStreamResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTs

`func (o *CalendarStreamResponse) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CalendarStreamResponse) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CalendarStreamResponse) SetTs(v int64)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CalendarStreamResponse) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to README]](../README.md)


