# PredictionOrderBookByMarketIdResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**MsgType** | Pointer to **string** | Message type, always &#x60;orderbook&#x60;. | [optional] 
**MarketId** | Pointer to **int64** | Market ID of this order book, the same &#x60;markets[].marketId&#x60; used by the REST API. | [optional] 
**UpdateTimestampMs** | Pointer to **int64** | Quote update time in milliseconds (UTC). Messages for the same market may arrive out of order, so discard anything older than the snapshot you already have. | [optional] 
**Asks** | Pointer to **[][]float64** | Ask levels sorted by price **ascending**, usually the top 50 or fewer; an empty array means no asks. Values may carry floating-point noise (e.g. &#x60;6865.999999999999&#x60;), so round them before display. | [optional] 
**Bids** | Pointer to **[][]float64** | Bid levels sorted by price **descending**, usually the top 50 or fewer; an empty array means no bids. Values may carry floating-point noise (e.g. &#x60;6865.999999999999&#x60;), so round them before display. | [optional] 

## Methods

### NewPredictionOrderBookByMarketIdResponse

`func NewPredictionOrderBookByMarketIdResponse() *PredictionOrderBookByMarketIdResponse`

NewPredictionOrderBookByMarketIdResponse instantiates a new PredictionOrderBookByMarketIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictionOrderBookByMarketIdResponseWithDefaults

`func NewPredictionOrderBookByMarketIdResponseWithDefaults() *PredictionOrderBookByMarketIdResponse`

NewPredictionOrderBookByMarketIdResponseWithDefaults instantiates a new PredictionOrderBookByMarketIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgType

`func (o *PredictionOrderBookByMarketIdResponse) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *PredictionOrderBookByMarketIdResponse) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *PredictionOrderBookByMarketIdResponse) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *PredictionOrderBookByMarketIdResponse) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetMarketId

`func (o *PredictionOrderBookByMarketIdResponse) GetMarketId() int64`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *PredictionOrderBookByMarketIdResponse) GetMarketIdOk() (*int64, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *PredictionOrderBookByMarketIdResponse) SetMarketId(v int64)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *PredictionOrderBookByMarketIdResponse) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetUpdateTimestampMs

`func (o *PredictionOrderBookByMarketIdResponse) GetUpdateTimestampMs() int64`

GetUpdateTimestampMs returns the UpdateTimestampMs field if non-nil, zero value otherwise.

### GetUpdateTimestampMsOk

`func (o *PredictionOrderBookByMarketIdResponse) GetUpdateTimestampMsOk() (*int64, bool)`

GetUpdateTimestampMsOk returns a tuple with the UpdateTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestampMs

`func (o *PredictionOrderBookByMarketIdResponse) SetUpdateTimestampMs(v int64)`

SetUpdateTimestampMs sets UpdateTimestampMs field to given value.

### HasUpdateTimestampMs

`func (o *PredictionOrderBookByMarketIdResponse) HasUpdateTimestampMs() bool`

HasUpdateTimestampMs returns a boolean if a field has been set.

### GetAsks

`func (o *PredictionOrderBookByMarketIdResponse) GetAsks() [][]float64`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *PredictionOrderBookByMarketIdResponse) GetAsksOk() (*[][]float64, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *PredictionOrderBookByMarketIdResponse) SetAsks(v [][]float64)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *PredictionOrderBookByMarketIdResponse) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetBids

`func (o *PredictionOrderBookByMarketIdResponse) GetBids() [][]float64`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *PredictionOrderBookByMarketIdResponse) GetBidsOk() (*[][]float64, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *PredictionOrderBookByMarketIdResponse) SetBids(v [][]float64)`

SetBids sets Bids field to given value.

### HasBids

`func (o *PredictionOrderBookByMarketIdResponse) HasBids() bool`

HasBids returns a boolean if a field has been set.


[[Back to README]](../README.md)


