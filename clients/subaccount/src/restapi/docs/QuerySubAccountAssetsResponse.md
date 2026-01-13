# QuerySubAccountAssetsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Balances** | Pointer to [**[]QuerySubAccountAssetsResponseBalancesInner**](QuerySubAccountAssetsResponseBalancesInner.md) |  | [optional] 

## Methods

### NewQuerySubAccountAssetsResponse

`func NewQuerySubAccountAssetsResponse() *QuerySubAccountAssetsResponse`

NewQuerySubAccountAssetsResponse instantiates a new QuerySubAccountAssetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySubAccountAssetsResponseWithDefaults

`func NewQuerySubAccountAssetsResponseWithDefaults() *QuerySubAccountAssetsResponse`

NewQuerySubAccountAssetsResponseWithDefaults instantiates a new QuerySubAccountAssetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalances

`func (o *QuerySubAccountAssetsResponse) GetBalances() []QuerySubAccountAssetsResponseBalancesInner`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *QuerySubAccountAssetsResponse) GetBalancesOk() (*[]QuerySubAccountAssetsResponseBalancesInner, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *QuerySubAccountAssetsResponse) SetBalances(v []QuerySubAccountAssetsResponseBalancesInner)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *QuerySubAccountAssetsResponse) HasBalances() bool`

HasBalances returns a boolean if a field has been set.


[[Back to README]](../README.md)


