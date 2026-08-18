# PreviewOtcBlocktradeResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**OrderData** | Pointer to [**PreviewOtcBlocktradeResponseOrderData**](PreviewOtcBlocktradeResponseOrderData.md) |  | [optional] 

## Methods

### NewPreviewOtcBlocktradeResponse

`func NewPreviewOtcBlocktradeResponse() *PreviewOtcBlocktradeResponse`

NewPreviewOtcBlocktradeResponse instantiates a new PreviewOtcBlocktradeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewOtcBlocktradeResponseWithDefaults

`func NewPreviewOtcBlocktradeResponseWithDefaults() *PreviewOtcBlocktradeResponse`

NewPreviewOtcBlocktradeResponseWithDefaults instantiates a new PreviewOtcBlocktradeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *PreviewOtcBlocktradeResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PreviewOtcBlocktradeResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PreviewOtcBlocktradeResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PreviewOtcBlocktradeResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *PreviewOtcBlocktradeResponse) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *PreviewOtcBlocktradeResponse) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetStatus

`func (o *PreviewOtcBlocktradeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PreviewOtcBlocktradeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PreviewOtcBlocktradeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PreviewOtcBlocktradeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrderData

`func (o *PreviewOtcBlocktradeResponse) GetOrderData() PreviewOtcBlocktradeResponseOrderData`

GetOrderData returns the OrderData field if non-nil, zero value otherwise.

### GetOrderDataOk

`func (o *PreviewOtcBlocktradeResponse) GetOrderDataOk() (*PreviewOtcBlocktradeResponseOrderData, bool)`

GetOrderDataOk returns a tuple with the OrderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderData

`func (o *PreviewOtcBlocktradeResponse) SetOrderData(v PreviewOtcBlocktradeResponseOrderData)`

SetOrderData sets OrderData field to given value.

### HasOrderData

`func (o *PreviewOtcBlocktradeResponse) HasOrderData() bool`

HasOrderData returns a boolean if a field has been set.


[[Back to README]](../README.md)


