# PlaceEquityOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Acknowledgement code: &#x60;S&#x60; &#x3D; accepted, &#x60;F&#x60; &#x3D; failed. Not an order lifecycle status — to poll lifecycle, call &#x60;/order/detail&#x60; or &#x60;/order/history&#x60;. | [optional] 
**OrderId** | Pointer to **string** | Order id (UUID). | [optional] 
**ClientOrderId** | Pointer to **string** | Echoes the supplied or server-generated client order id. | [optional] 

## Methods

### NewPlaceEquityOrderResponse

`func NewPlaceEquityOrderResponse() *PlaceEquityOrderResponse`

NewPlaceEquityOrderResponse instantiates a new PlaceEquityOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceEquityOrderResponseWithDefaults

`func NewPlaceEquityOrderResponseWithDefaults() *PlaceEquityOrderResponse`

NewPlaceEquityOrderResponseWithDefaults instantiates a new PlaceEquityOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PlaceEquityOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlaceEquityOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlaceEquityOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlaceEquityOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrderId

`func (o *PlaceEquityOrderResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PlaceEquityOrderResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PlaceEquityOrderResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PlaceEquityOrderResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *PlaceEquityOrderResponse) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *PlaceEquityOrderResponse) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *PlaceEquityOrderResponse) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *PlaceEquityOrderResponse) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.


[[Back to README]](../README.md)


