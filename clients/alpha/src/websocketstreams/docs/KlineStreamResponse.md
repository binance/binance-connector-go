# KlineStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | eventType | [optional] 
**E** | Pointer to **int64** | eventTime | [optional] 
**S** | Pointer to **string** | symbol | [optional] 
**K** | Pointer to [**KlineStreamResponseK**](KlineStreamResponseK.md) |  | [optional] 

## Methods

### NewKlineStreamResponse

`func NewKlineStreamResponse() *KlineStreamResponse`

NewKlineStreamResponse instantiates a new KlineStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlineStreamResponseWithDefaults

`func NewKlineStreamResponseWithDefaults() *KlineStreamResponse`

NewKlineStreamResponseWithDefaults instantiates a new KlineStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *KlineStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *KlineStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *KlineStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *KlineStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetE

`func (o *KlineStreamResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *KlineStreamResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *KlineStreamResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *KlineStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetS

`func (o *KlineStreamResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *KlineStreamResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *KlineStreamResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *KlineStreamResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetK

`func (o *KlineStreamResponse) GetK() KlineStreamResponseK`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *KlineStreamResponse) GetKOk() (*KlineStreamResponseK, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *KlineStreamResponse) SetK(v KlineStreamResponseK)`

SetK sets K field to given value.

### HasK

`func (o *KlineStreamResponse) HasK() bool`

HasK returns a boolean if a field has been set.


[[Back to README]](../README.md)


