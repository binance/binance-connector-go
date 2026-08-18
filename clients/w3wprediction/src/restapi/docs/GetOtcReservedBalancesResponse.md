# GetOtcReservedBalancesResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Balances** | Pointer to [**[]GetOtcReservedBalancesResponseBalancesInner**](GetOtcReservedBalancesResponseBalancesInner.md) |  | [optional] 

## Methods

### NewGetOtcReservedBalancesResponse

`func NewGetOtcReservedBalancesResponse() *GetOtcReservedBalancesResponse`

NewGetOtcReservedBalancesResponse instantiates a new GetOtcReservedBalancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOtcReservedBalancesResponseWithDefaults

`func NewGetOtcReservedBalancesResponseWithDefaults() *GetOtcReservedBalancesResponse`

NewGetOtcReservedBalancesResponseWithDefaults instantiates a new GetOtcReservedBalancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalances

`func (o *GetOtcReservedBalancesResponse) GetBalances() []GetOtcReservedBalancesResponseBalancesInner`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *GetOtcReservedBalancesResponse) GetBalancesOk() (*[]GetOtcReservedBalancesResponseBalancesInner, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *GetOtcReservedBalancesResponse) SetBalances(v []GetOtcReservedBalancesResponseBalancesInner)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *GetOtcReservedBalancesResponse) HasBalances() bool`

HasBalances returns a boolean if a field has been set.


[[Back to README]](../README.md)


