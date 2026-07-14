# LiquidationLoanRepayResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**RepayId** | Pointer to **int64** | Unique identifier for this repayment transaction | [optional] 
**Asset** | Pointer to **string** | Asset used for repayment | [optional] 
**Amount** | Pointer to **string** | Actual repayment amount | [optional] 
**Status** | Pointer to **string** | Repayment status: &#x60;SUCCESS&#x60; (completed) or &#x60;PENDING&#x60; (processing) | [optional] 
**CreateTime** | Pointer to **int64** | Unix timestamp (milliseconds) when the repayment was created | [optional] 

## Methods

### NewLiquidationLoanRepayResponse

`func NewLiquidationLoanRepayResponse() *LiquidationLoanRepayResponse`

NewLiquidationLoanRepayResponse instantiates a new LiquidationLoanRepayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiquidationLoanRepayResponseWithDefaults

`func NewLiquidationLoanRepayResponseWithDefaults() *LiquidationLoanRepayResponse`

NewLiquidationLoanRepayResponseWithDefaults instantiates a new LiquidationLoanRepayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepayId

`func (o *LiquidationLoanRepayResponse) GetRepayId() int64`

GetRepayId returns the RepayId field if non-nil, zero value otherwise.

### GetRepayIdOk

`func (o *LiquidationLoanRepayResponse) GetRepayIdOk() (*int64, bool)`

GetRepayIdOk returns a tuple with the RepayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayId

`func (o *LiquidationLoanRepayResponse) SetRepayId(v int64)`

SetRepayId sets RepayId field to given value.

### HasRepayId

`func (o *LiquidationLoanRepayResponse) HasRepayId() bool`

HasRepayId returns a boolean if a field has been set.

### GetAsset

`func (o *LiquidationLoanRepayResponse) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *LiquidationLoanRepayResponse) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *LiquidationLoanRepayResponse) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *LiquidationLoanRepayResponse) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAmount

`func (o *LiquidationLoanRepayResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LiquidationLoanRepayResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LiquidationLoanRepayResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *LiquidationLoanRepayResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetStatus

`func (o *LiquidationLoanRepayResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LiquidationLoanRepayResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LiquidationLoanRepayResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LiquidationLoanRepayResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreateTime

`func (o *LiquidationLoanRepayResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *LiquidationLoanRepayResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *LiquidationLoanRepayResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *LiquidationLoanRepayResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


