# AutoCancelAllOpenOrdersResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**CountdownTime** | Pointer to **string** |  | [optional] 

## Methods

### NewAutoCancelAllOpenOrdersResponse

`func NewAutoCancelAllOpenOrdersResponse() *AutoCancelAllOpenOrdersResponse`

NewAutoCancelAllOpenOrdersResponse instantiates a new AutoCancelAllOpenOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCancelAllOpenOrdersResponseWithDefaults

`func NewAutoCancelAllOpenOrdersResponseWithDefaults() *AutoCancelAllOpenOrdersResponse`

NewAutoCancelAllOpenOrdersResponseWithDefaults instantiates a new AutoCancelAllOpenOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *AutoCancelAllOpenOrdersResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AutoCancelAllOpenOrdersResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AutoCancelAllOpenOrdersResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AutoCancelAllOpenOrdersResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCountdownTime

`func (o *AutoCancelAllOpenOrdersResponse) GetCountdownTime() string`

GetCountdownTime returns the CountdownTime field if non-nil, zero value otherwise.

### GetCountdownTimeOk

`func (o *AutoCancelAllOpenOrdersResponse) GetCountdownTimeOk() (*string, bool)`

GetCountdownTimeOk returns a tuple with the CountdownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountdownTime

`func (o *AutoCancelAllOpenOrdersResponse) SetCountdownTime(v string)`

SetCountdownTime sets CountdownTime field to given value.

### HasCountdownTime

`func (o *AutoCancelAllOpenOrdersResponse) HasCountdownTime() bool`

HasCountdownTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


