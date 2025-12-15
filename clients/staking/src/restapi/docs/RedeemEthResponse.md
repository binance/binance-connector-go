# RedeemEthResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**EthAmount** | Pointer to **string** |  | [optional] 
**ConversionRatio** | Pointer to **string** |  | [optional] 
**ArrivalTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewRedeemEthResponse

`func NewRedeemEthResponse() *RedeemEthResponse`

NewRedeemEthResponse instantiates a new RedeemEthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemEthResponseWithDefaults

`func NewRedeemEthResponseWithDefaults() *RedeemEthResponse`

NewRedeemEthResponseWithDefaults instantiates a new RedeemEthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RedeemEthResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RedeemEthResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RedeemEthResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *RedeemEthResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetEthAmount

`func (o *RedeemEthResponse) GetEthAmount() string`

GetEthAmount returns the EthAmount field if non-nil, zero value otherwise.

### GetEthAmountOk

`func (o *RedeemEthResponse) GetEthAmountOk() (*string, bool)`

GetEthAmountOk returns a tuple with the EthAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthAmount

`func (o *RedeemEthResponse) SetEthAmount(v string)`

SetEthAmount sets EthAmount field to given value.

### HasEthAmount

`func (o *RedeemEthResponse) HasEthAmount() bool`

HasEthAmount returns a boolean if a field has been set.

### GetConversionRatio

`func (o *RedeemEthResponse) GetConversionRatio() string`

GetConversionRatio returns the ConversionRatio field if non-nil, zero value otherwise.

### GetConversionRatioOk

`func (o *RedeemEthResponse) GetConversionRatioOk() (*string, bool)`

GetConversionRatioOk returns a tuple with the ConversionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRatio

`func (o *RedeemEthResponse) SetConversionRatio(v string)`

SetConversionRatio sets ConversionRatio field to given value.

### HasConversionRatio

`func (o *RedeemEthResponse) HasConversionRatio() bool`

HasConversionRatio returns a boolean if a field has been set.

### GetArrivalTime

`func (o *RedeemEthResponse) GetArrivalTime() int64`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *RedeemEthResponse) GetArrivalTimeOk() (*int64, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *RedeemEthResponse) SetArrivalTime(v int64)`

SetArrivalTime sets ArrivalTime field to given value.

### HasArrivalTime

`func (o *RedeemEthResponse) HasArrivalTime() bool`

HasArrivalTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


