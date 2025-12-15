# RedeemSolResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**SolAmount** | Pointer to **string** |  | [optional] 
**ExchangeRate** | Pointer to **string** |  | [optional] 
**ArrivalTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewRedeemSolResponse

`func NewRedeemSolResponse() *RedeemSolResponse`

NewRedeemSolResponse instantiates a new RedeemSolResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemSolResponseWithDefaults

`func NewRedeemSolResponseWithDefaults() *RedeemSolResponse`

NewRedeemSolResponseWithDefaults instantiates a new RedeemSolResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RedeemSolResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RedeemSolResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RedeemSolResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *RedeemSolResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSolAmount

`func (o *RedeemSolResponse) GetSolAmount() string`

GetSolAmount returns the SolAmount field if non-nil, zero value otherwise.

### GetSolAmountOk

`func (o *RedeemSolResponse) GetSolAmountOk() (*string, bool)`

GetSolAmountOk returns a tuple with the SolAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolAmount

`func (o *RedeemSolResponse) SetSolAmount(v string)`

SetSolAmount sets SolAmount field to given value.

### HasSolAmount

`func (o *RedeemSolResponse) HasSolAmount() bool`

HasSolAmount returns a boolean if a field has been set.

### GetExchangeRate

`func (o *RedeemSolResponse) GetExchangeRate() string`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *RedeemSolResponse) GetExchangeRateOk() (*string, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *RedeemSolResponse) SetExchangeRate(v string)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *RedeemSolResponse) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetArrivalTime

`func (o *RedeemSolResponse) GetArrivalTime() int64`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *RedeemSolResponse) GetArrivalTimeOk() (*int64, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *RedeemSolResponse) SetArrivalTime(v int64)`

SetArrivalTime sets ArrivalTime field to given value.

### HasArrivalTime

`func (o *RedeemSolResponse) HasArrivalTime() bool`

HasArrivalTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


