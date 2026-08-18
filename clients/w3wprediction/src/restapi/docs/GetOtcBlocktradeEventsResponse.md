# GetOtcBlocktradeEventsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Events** | Pointer to [**[]GetOtcBlocktradeEventsResponseEventsInner**](GetOtcBlocktradeEventsResponseEventsInner.md) |  | [optional] 

## Methods

### NewGetOtcBlocktradeEventsResponse

`func NewGetOtcBlocktradeEventsResponse() *GetOtcBlocktradeEventsResponse`

NewGetOtcBlocktradeEventsResponse instantiates a new GetOtcBlocktradeEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOtcBlocktradeEventsResponseWithDefaults

`func NewGetOtcBlocktradeEventsResponseWithDefaults() *GetOtcBlocktradeEventsResponse`

NewGetOtcBlocktradeEventsResponseWithDefaults instantiates a new GetOtcBlocktradeEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *GetOtcBlocktradeEventsResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetOtcBlocktradeEventsResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetOtcBlocktradeEventsResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetOtcBlocktradeEventsResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetEvents

`func (o *GetOtcBlocktradeEventsResponse) GetEvents() []GetOtcBlocktradeEventsResponseEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetOtcBlocktradeEventsResponse) GetEventsOk() (*[]GetOtcBlocktradeEventsResponseEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetOtcBlocktradeEventsResponse) SetEvents(v []GetOtcBlocktradeEventsResponseEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GetOtcBlocktradeEventsResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to README]](../README.md)


