# FullDepthResponseData

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateId** | Pointer to **int64** | Last order book update ID. | [optional] 
**Symbol** | Pointer to **string** | Trading pair symbol. | [optional] 
**Bids** | Pointer to **[][]string** | Bid orders. Each entry is [price, quantity]. | [optional] 
**Asks** | Pointer to **[][]string** | Ask orders. Each entry is [price, quantity]. | [optional] 
**E** | Pointer to **int64** | Event time in milliseconds. | [optional] 
**T** | Pointer to **int64** | Transaction time in milliseconds. | [optional] 

## Methods

### NewFullDepthResponseData

`func NewFullDepthResponseData() *FullDepthResponseData`

NewFullDepthResponseData instantiates a new FullDepthResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullDepthResponseDataWithDefaults

`func NewFullDepthResponseDataWithDefaults() *FullDepthResponseData`

NewFullDepthResponseDataWithDefaults instantiates a new FullDepthResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateId

`func (o *FullDepthResponseData) GetLastUpdateId() int64`

GetLastUpdateId returns the LastUpdateId field if non-nil, zero value otherwise.

### GetLastUpdateIdOk

`func (o *FullDepthResponseData) GetLastUpdateIdOk() (*int64, bool)`

GetLastUpdateIdOk returns a tuple with the LastUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateId

`func (o *FullDepthResponseData) SetLastUpdateId(v int64)`

SetLastUpdateId sets LastUpdateId field to given value.

### HasLastUpdateId

`func (o *FullDepthResponseData) HasLastUpdateId() bool`

HasLastUpdateId returns a boolean if a field has been set.

### GetSymbol

`func (o *FullDepthResponseData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FullDepthResponseData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FullDepthResponseData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FullDepthResponseData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetBids

`func (o *FullDepthResponseData) GetBids() [][]string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *FullDepthResponseData) GetBidsOk() (*[][]string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *FullDepthResponseData) SetBids(v [][]string)`

SetBids sets Bids field to given value.

### HasBids

`func (o *FullDepthResponseData) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetAsks

`func (o *FullDepthResponseData) GetAsks() [][]string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *FullDepthResponseData) GetAsksOk() (*[][]string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *FullDepthResponseData) SetAsks(v [][]string)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *FullDepthResponseData) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetE

`func (o *FullDepthResponseData) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *FullDepthResponseData) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *FullDepthResponseData) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *FullDepthResponseData) HasE() bool`

HasE returns a boolean if a field has been set.

### GetT

`func (o *FullDepthResponseData) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FullDepthResponseData) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FullDepthResponseData) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *FullDepthResponseData) HasT() bool`

HasT returns a boolean if a field has been set.


[[Back to README]](../README.md)


