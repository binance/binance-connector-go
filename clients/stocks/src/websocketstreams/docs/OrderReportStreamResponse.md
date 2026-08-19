# OrderReportStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event type, always &#x60;\&quot;orderReport\&quot;&#x60;. | [optional] 
**E** | Pointer to **int64** | Event time (epoch milliseconds); server push time. | [optional] 
**X** | Pointer to **string** | Execution type: &#x60;\&quot;ORDER_UPDATE\&quot;&#x60; (still open) or &#x60;\&quot;ORDER_TERMINAL\&quot;&#x60; (reached terminal state). | [optional] 
**I** | Pointer to **string** | Order ID (UUID). | [optional] 
**Ai** | Pointer to **string** | Asset ID (internal identifier). | [optional] 
**B** | Pointer to **string** | Base asset — the internal asset code with an &#x60;EQ_&#x60; prefix (e.g. &#x60;\&quot;EQ_AAPL\&quot;&#x60;), not the bare ticker used in REST responses / order input. Strip the &#x60;EQ_&#x60; prefix to match a symbol used elsewhere. | [optional] 
**Q** | Pointer to **string** | Quote currency, e.g. &#x60;\&quot;USD\&quot;&#x60;. | [optional] 
**S** | Pointer to **string** | Order side: &#x60;\&quot;buy\&quot;&#x60; or &#x60;\&quot;sell\&quot;&#x60;. Note: lowercase, unlike REST responses. | [optional] 
**O** | Pointer to **string** | Order type: &#x60;\&quot;market\&quot;&#x60; / &#x60;\&quot;limit\&quot;&#x60; / &#x60;\&quot;stop\&quot;&#x60; / &#x60;\&quot;stop_limit\&quot;&#x60; / &#x60;\&quot;trailing_stop\&quot;&#x60;. Note: lowercase. | [optional] 
**P** | Pointer to **NullableFloat32** | Limit price; null for market orders. | [optional] 
**Q** | Pointer to **float32** | Order quantity (shares); &#x60;0&#x60; when the order was submitted as notional. | [optional] 
**N** | Pointer to **NullableFloat32** | Order notional; set when the order was submitted as notional (market buy), null when submitted as quantity. | [optional] 
**Fq** | Pointer to **float32** | Filled quantity. | [optional] 
**FN** | Pointer to **float32** | Filled notional (&#x3D; filledQty × filledAvgPrice). | [optional] 
**Tc** | Pointer to **float32** | Total cost — cumulative buy-in cost including the commission fee. | [optional] 
**Z** | Pointer to **float32** | Fill progress percentage (0–100, 2 dp). By notional: &#x60;FN / N × 100&#x60;. By qty: &#x60;fq / Q × 100&#x60;. | [optional] 
**N** | Pointer to **string** | Trading session label, e.g. &#x60;\&quot;Regular\&quot;&#x60;, &#x60;\&quot;24 Hours Trading\&quot;&#x60;. | [optional] 
**S** | Pointer to **string** | Order status, e.g. &#x60;\&quot;accepted\&quot;&#x60;, &#x60;\&quot;partially_filled\&quot;&#x60;, &#x60;\&quot;filled\&quot;&#x60;, &#x60;\&quot;canceled\&quot;&#x60;. Note: lowercase, unlike REST responses. | [optional] 
**T** | Pointer to **int64** | Order create time (epoch milliseconds). | [optional] 
**U** | Pointer to **int64** | Order update time (epoch milliseconds). | [optional] 

## Methods

### NewOrderReportStreamResponse

`func NewOrderReportStreamResponse() *OrderReportStreamResponse`

NewOrderReportStreamResponse instantiates a new OrderReportStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderReportStreamResponseWithDefaults

`func NewOrderReportStreamResponseWithDefaults() *OrderReportStreamResponse`

NewOrderReportStreamResponseWithDefaults instantiates a new OrderReportStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *OrderReportStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OrderReportStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OrderReportStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *OrderReportStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetE

`func (o *OrderReportStreamResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OrderReportStreamResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OrderReportStreamResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *OrderReportStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetX

`func (o *OrderReportStreamResponse) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *OrderReportStreamResponse) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *OrderReportStreamResponse) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *OrderReportStreamResponse) HasX() bool`

HasX returns a boolean if a field has been set.

### GetI

`func (o *OrderReportStreamResponse) GetI() string`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *OrderReportStreamResponse) GetIOk() (*string, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *OrderReportStreamResponse) SetI(v string)`

SetI sets I field to given value.

### HasI

`func (o *OrderReportStreamResponse) HasI() bool`

HasI returns a boolean if a field has been set.

### GetAi

`func (o *OrderReportStreamResponse) GetAi() string`

GetAi returns the Ai field if non-nil, zero value otherwise.

### GetAiOk

`func (o *OrderReportStreamResponse) GetAiOk() (*string, bool)`

GetAiOk returns a tuple with the Ai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAi

`func (o *OrderReportStreamResponse) SetAi(v string)`

SetAi sets Ai field to given value.

### HasAi

`func (o *OrderReportStreamResponse) HasAi() bool`

HasAi returns a boolean if a field has been set.

### GetB

`func (o *OrderReportStreamResponse) GetB() string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *OrderReportStreamResponse) GetBOk() (*string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *OrderReportStreamResponse) SetB(v string)`

SetB sets B field to given value.

### HasB

`func (o *OrderReportStreamResponse) HasB() bool`

HasB returns a boolean if a field has been set.

### GetQ

`func (o *OrderReportStreamResponse) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *OrderReportStreamResponse) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *OrderReportStreamResponse) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *OrderReportStreamResponse) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetS

`func (o *OrderReportStreamResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *OrderReportStreamResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *OrderReportStreamResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *OrderReportStreamResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetO

`func (o *OrderReportStreamResponse) GetO() string`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *OrderReportStreamResponse) GetOOk() (*string, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *OrderReportStreamResponse) SetO(v string)`

SetO sets O field to given value.

### HasO

`func (o *OrderReportStreamResponse) HasO() bool`

HasO returns a boolean if a field has been set.

### GetP

`func (o *OrderReportStreamResponse) GetP() float32`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *OrderReportStreamResponse) GetPOk() (*float32, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *OrderReportStreamResponse) SetP(v float32)`

SetP sets P field to given value.

### HasP

`func (o *OrderReportStreamResponse) HasP() bool`

HasP returns a boolean if a field has been set.

### SetPNil

`func (o *OrderReportStreamResponse) SetPNil(b bool)`

 SetPNil sets the value for P to be an explicit nil

### UnsetP
`func (o *OrderReportStreamResponse) UnsetP()`

UnsetP ensures that no value is present for P, not even an explicit nil
### GetQ

`func (o *OrderReportStreamResponse) GetQ() float32`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *OrderReportStreamResponse) GetQOk() (*float32, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *OrderReportStreamResponse) SetQ(v float32)`

SetQ sets Q field to given value.

### HasQ

`func (o *OrderReportStreamResponse) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetN

`func (o *OrderReportStreamResponse) GetN() float32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *OrderReportStreamResponse) GetNOk() (*float32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *OrderReportStreamResponse) SetN(v float32)`

SetN sets N field to given value.

### HasN

`func (o *OrderReportStreamResponse) HasN() bool`

HasN returns a boolean if a field has been set.

### SetNNil

`func (o *OrderReportStreamResponse) SetNNil(b bool)`

 SetNNil sets the value for N to be an explicit nil

### UnsetN
`func (o *OrderReportStreamResponse) UnsetN()`

UnsetN ensures that no value is present for N, not even an explicit nil
### GetFq

`func (o *OrderReportStreamResponse) GetFq() float32`

GetFq returns the Fq field if non-nil, zero value otherwise.

### GetFqOk

`func (o *OrderReportStreamResponse) GetFqOk() (*float32, bool)`

GetFqOk returns a tuple with the Fq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFq

`func (o *OrderReportStreamResponse) SetFq(v float32)`

SetFq sets Fq field to given value.

### HasFq

`func (o *OrderReportStreamResponse) HasFq() bool`

HasFq returns a boolean if a field has been set.

### GetFN

`func (o *OrderReportStreamResponse) GetFN() float32`

GetFN returns the FN field if non-nil, zero value otherwise.

### GetFNOk

`func (o *OrderReportStreamResponse) GetFNOk() (*float32, bool)`

GetFNOk returns a tuple with the FN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFN

`func (o *OrderReportStreamResponse) SetFN(v float32)`

SetFN sets FN field to given value.

### HasFN

`func (o *OrderReportStreamResponse) HasFN() bool`

HasFN returns a boolean if a field has been set.

### GetTc

`func (o *OrderReportStreamResponse) GetTc() float32`

GetTc returns the Tc field if non-nil, zero value otherwise.

### GetTcOk

`func (o *OrderReportStreamResponse) GetTcOk() (*float32, bool)`

GetTcOk returns a tuple with the Tc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTc

`func (o *OrderReportStreamResponse) SetTc(v float32)`

SetTc sets Tc field to given value.

### HasTc

`func (o *OrderReportStreamResponse) HasTc() bool`

HasTc returns a boolean if a field has been set.

### GetZ

`func (o *OrderReportStreamResponse) GetZ() float32`

GetZ returns the Z field if non-nil, zero value otherwise.

### GetZOk

`func (o *OrderReportStreamResponse) GetZOk() (*float32, bool)`

GetZOk returns a tuple with the Z field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZ

`func (o *OrderReportStreamResponse) SetZ(v float32)`

SetZ sets Z field to given value.

### HasZ

`func (o *OrderReportStreamResponse) HasZ() bool`

HasZ returns a boolean if a field has been set.

### GetN

`func (o *OrderReportStreamResponse) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *OrderReportStreamResponse) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *OrderReportStreamResponse) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *OrderReportStreamResponse) HasN() bool`

HasN returns a boolean if a field has been set.

### GetS

`func (o *OrderReportStreamResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *OrderReportStreamResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *OrderReportStreamResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *OrderReportStreamResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetT

`func (o *OrderReportStreamResponse) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *OrderReportStreamResponse) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *OrderReportStreamResponse) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *OrderReportStreamResponse) HasT() bool`

HasT returns a boolean if a field has been set.

### GetU

`func (o *OrderReportStreamResponse) GetU() int64`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *OrderReportStreamResponse) GetUOk() (*int64, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *OrderReportStreamResponse) SetU(v int64)`

SetU sets U field to given value.

### HasU

`func (o *OrderReportStreamResponse) HasU() bool`

HasU returns a boolean if a field has been set.


[[Back to README]](../README.md)


