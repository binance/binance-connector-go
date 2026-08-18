# GetOtcBlocktradeDetailResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**OrderData** | Pointer to [**GetOtcBlocktradeDetailResponseOrderData**](GetOtcBlocktradeDetailResponseOrderData.md) |  | [optional] 

## Methods

### NewGetOtcBlocktradeDetailResponse

`func NewGetOtcBlocktradeDetailResponse() *GetOtcBlocktradeDetailResponse`

NewGetOtcBlocktradeDetailResponse instantiates a new GetOtcBlocktradeDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOtcBlocktradeDetailResponseWithDefaults

`func NewGetOtcBlocktradeDetailResponseWithDefaults() *GetOtcBlocktradeDetailResponse`

NewGetOtcBlocktradeDetailResponseWithDefaults instantiates a new GetOtcBlocktradeDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetOtcBlocktradeDetailResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOtcBlocktradeDetailResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOtcBlocktradeDetailResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetOtcBlocktradeDetailResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *GetOtcBlocktradeDetailResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOtcBlocktradeDetailResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOtcBlocktradeDetailResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOtcBlocktradeDetailResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrderData

`func (o *GetOtcBlocktradeDetailResponse) GetOrderData() GetOtcBlocktradeDetailResponseOrderData`

GetOrderData returns the OrderData field if non-nil, zero value otherwise.

### GetOrderDataOk

`func (o *GetOtcBlocktradeDetailResponse) GetOrderDataOk() (*GetOtcBlocktradeDetailResponseOrderData, bool)`

GetOrderDataOk returns a tuple with the OrderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderData

`func (o *GetOtcBlocktradeDetailResponse) SetOrderData(v GetOtcBlocktradeDetailResponseOrderData)`

SetOrderData sets OrderData field to given value.

### HasOrderData

`func (o *GetOtcBlocktradeDetailResponse) HasOrderData() bool`

HasOrderData returns a boolean if a field has been set.


[[Back to README]](../README.md)


