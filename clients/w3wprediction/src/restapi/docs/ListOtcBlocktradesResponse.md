# ListOtcBlocktradesResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Blocktrades** | Pointer to [**[]ListOtcBlocktradesResponseBlocktradesInner**](ListOtcBlocktradesResponseBlocktradesInner.md) |  | [optional] 

## Methods

### NewListOtcBlocktradesResponse

`func NewListOtcBlocktradesResponse() *ListOtcBlocktradesResponse`

NewListOtcBlocktradesResponse instantiates a new ListOtcBlocktradesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOtcBlocktradesResponseWithDefaults

`func NewListOtcBlocktradesResponseWithDefaults() *ListOtcBlocktradesResponse`

NewListOtcBlocktradesResponseWithDefaults instantiates a new ListOtcBlocktradesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ListOtcBlocktradesResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListOtcBlocktradesResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListOtcBlocktradesResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListOtcBlocktradesResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetBlocktrades

`func (o *ListOtcBlocktradesResponse) GetBlocktrades() []ListOtcBlocktradesResponseBlocktradesInner`

GetBlocktrades returns the Blocktrades field if non-nil, zero value otherwise.

### GetBlocktradesOk

`func (o *ListOtcBlocktradesResponse) GetBlocktradesOk() (*[]ListOtcBlocktradesResponseBlocktradesInner, bool)`

GetBlocktradesOk returns a tuple with the Blocktrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocktrades

`func (o *ListOtcBlocktradesResponse) SetBlocktrades(v []ListOtcBlocktradesResponseBlocktradesInner)`

SetBlocktrades sets Blocktrades field to given value.

### HasBlocktrades

`func (o *ListOtcBlocktradesResponse) HasBlocktrades() bool`

HasBlocktrades returns a boolean if a field has been set.


[[Back to README]](../README.md)


