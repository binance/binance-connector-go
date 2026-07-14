# QueryLiquidationLoanResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** | The asset of the liquidation loan (USDC by default) | [optional] 
**Amount** | Pointer to **string** | Total liquidation loan amount | [optional] 
**RepaidAmount** | Pointer to **string** | Amount that has been repaid | [optional] 
**RemainingAmount** | Pointer to **string** | Outstanding amount remaining to be repaid | [optional] 

## Methods

### NewQueryLiquidationLoanResponse

`func NewQueryLiquidationLoanResponse() *QueryLiquidationLoanResponse`

NewQueryLiquidationLoanResponse instantiates a new QueryLiquidationLoanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLiquidationLoanResponseWithDefaults

`func NewQueryLiquidationLoanResponseWithDefaults() *QueryLiquidationLoanResponse`

NewQueryLiquidationLoanResponseWithDefaults instantiates a new QueryLiquidationLoanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *QueryLiquidationLoanResponse) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *QueryLiquidationLoanResponse) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *QueryLiquidationLoanResponse) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *QueryLiquidationLoanResponse) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAmount

`func (o *QueryLiquidationLoanResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QueryLiquidationLoanResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QueryLiquidationLoanResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QueryLiquidationLoanResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetRepaidAmount

`func (o *QueryLiquidationLoanResponse) GetRepaidAmount() string`

GetRepaidAmount returns the RepaidAmount field if non-nil, zero value otherwise.

### GetRepaidAmountOk

`func (o *QueryLiquidationLoanResponse) GetRepaidAmountOk() (*string, bool)`

GetRepaidAmountOk returns a tuple with the RepaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepaidAmount

`func (o *QueryLiquidationLoanResponse) SetRepaidAmount(v string)`

SetRepaidAmount sets RepaidAmount field to given value.

### HasRepaidAmount

`func (o *QueryLiquidationLoanResponse) HasRepaidAmount() bool`

HasRepaidAmount returns a boolean if a field has been set.

### GetRemainingAmount

`func (o *QueryLiquidationLoanResponse) GetRemainingAmount() string`

GetRemainingAmount returns the RemainingAmount field if non-nil, zero value otherwise.

### GetRemainingAmountOk

`func (o *QueryLiquidationLoanResponse) GetRemainingAmountOk() (*string, bool)`

GetRemainingAmountOk returns a tuple with the RemainingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingAmount

`func (o *QueryLiquidationLoanResponse) SetRemainingAmount(v string)`

SetRemainingAmount sets RemainingAmount field to given value.

### HasRemainingAmount

`func (o *QueryLiquidationLoanResponse) HasRemainingAmount() bool`

HasRemainingAmount returns a boolean if a field has been set.


[[Back to README]](../README.md)


