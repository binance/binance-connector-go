# PredictionOrderBookAllResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**MsgType** | Pointer to **string** | Message type, fixed value &#x60;orderbook&#x60;; use it to dispatch message types. | [optional] 
**MarketId** | Pointer to **int64** | Binance internal market ID (mapped from the upstream vendor ID), aligned with the REST API. | [optional] 
**UpdateTimestampMs** | Pointer to **int64** | Quote update time, millisecond UTC timestamp. No strict ordering is guaranteed for the same &#x60;marketId&#x60;, so use this field for out-of-order detection and discard messages older than the current snapshot. | [optional] 
**Asks** | Pointer to **[][]string** | Ask levels sorted by price **ascending**; an empty array means no asks. Typically the top N levels (generally &lt;&#x3D; 50, depending on upstream policy). | [optional] 
**Bids** | Pointer to **[][]string** | Bid levels sorted by price **descending**; an empty array means no bids. Typically the top N levels (generally &lt;&#x3D; 50, depending on upstream policy). | [optional] 

## Methods

### NewPredictionOrderBookAllResponse

`func NewPredictionOrderBookAllResponse() *PredictionOrderBookAllResponse`

NewPredictionOrderBookAllResponse instantiates a new PredictionOrderBookAllResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictionOrderBookAllResponseWithDefaults

`func NewPredictionOrderBookAllResponseWithDefaults() *PredictionOrderBookAllResponse`

NewPredictionOrderBookAllResponseWithDefaults instantiates a new PredictionOrderBookAllResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgType

`func (o *PredictionOrderBookAllResponse) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *PredictionOrderBookAllResponse) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *PredictionOrderBookAllResponse) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *PredictionOrderBookAllResponse) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetMarketId

`func (o *PredictionOrderBookAllResponse) GetMarketId() int64`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *PredictionOrderBookAllResponse) GetMarketIdOk() (*int64, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *PredictionOrderBookAllResponse) SetMarketId(v int64)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *PredictionOrderBookAllResponse) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetUpdateTimestampMs

`func (o *PredictionOrderBookAllResponse) GetUpdateTimestampMs() int64`

GetUpdateTimestampMs returns the UpdateTimestampMs field if non-nil, zero value otherwise.

### GetUpdateTimestampMsOk

`func (o *PredictionOrderBookAllResponse) GetUpdateTimestampMsOk() (*int64, bool)`

GetUpdateTimestampMsOk returns a tuple with the UpdateTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestampMs

`func (o *PredictionOrderBookAllResponse) SetUpdateTimestampMs(v int64)`

SetUpdateTimestampMs sets UpdateTimestampMs field to given value.

### HasUpdateTimestampMs

`func (o *PredictionOrderBookAllResponse) HasUpdateTimestampMs() bool`

HasUpdateTimestampMs returns a boolean if a field has been set.

### GetAsks

`func (o *PredictionOrderBookAllResponse) GetAsks() [][]string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *PredictionOrderBookAllResponse) GetAsksOk() (*[][]string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *PredictionOrderBookAllResponse) SetAsks(v [][]string)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *PredictionOrderBookAllResponse) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetBids

`func (o *PredictionOrderBookAllResponse) GetBids() [][]string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *PredictionOrderBookAllResponse) GetBidsOk() (*[][]string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *PredictionOrderBookAllResponse) SetBids(v [][]string)`

SetBids sets Bids field to given value.

### HasBids

`func (o *PredictionOrderBookAllResponse) HasBids() bool`

HasBids returns a boolean if a field has been set.


[[Back to README]](../README.md)


