# UserDataStreamEventsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** | Event Time | [optional] 
**A** | Pointer to **string** | Asset | [optional] 
**D** | Pointer to **string** | Delta | [optional] 
**T** | Pointer to **int64** | Transaction Time | [optional] 
**S** | Pointer to **string** | Symbol | [optional] 
**C** | Pointer to **string** | Contingency Type | [optional] 
**S** | Pointer to **string** | Side | [optional] 
**O** | Pointer to **string** | Order type | [optional] 
**F** | Pointer to **string** | Time in force | [optional] 
**Q** | Pointer to **string** | Order quantity | [optional] 
**P** | Pointer to **string** | Order price | [optional] 
**P** | Pointer to **string** | Stop price | [optional] 
**F** | Pointer to **string** | Iceberg quantity | [optional] 
**G** | Pointer to **int64** | OrderListId | [optional] 
**C** | Pointer to **string** | List Client Order ID | [optional] 
**X** | Pointer to **string** | Current execution type | [optional] 
**X** | Pointer to **string** | Current order status | [optional] 
**R** | Pointer to **string** | List Reject Reason | [optional] 
**I** | Pointer to **int64** | Order ID | [optional] 
**L** | Pointer to **string** | List Status Type | [optional] 
**Z** | Pointer to **string** | Cumulative filled quantity | [optional] 
**L** | Pointer to **string** | List Order Status | [optional] 
**N** | Pointer to **string** | Commission amount | [optional] 
**N** | Pointer to **string** | Commission asset | [optional] 
**T** | Pointer to **int64** | Trade ID | [optional] 
**V** | Pointer to **int64** | Prevented Match Id; This is only visible if the order expired due to STP | [optional] 
**I** | Pointer to **int64** | Execution Id | [optional] 
**W** | Pointer to **bool** | Is the order on the book? | [optional] 
**M** | Pointer to **bool** | Is this trade the maker side? | [optional] 
**M** | Pointer to **bool** | Ignore | [optional] 
**O** | Pointer to [**[]ListStatusOInner**](ListStatusOInner.md) | An array of objects | [optional] 
**Z** | Pointer to **string** | Cumulative quote asset transacted quantity | [optional] 
**Y** | Pointer to **string** | Last quote asset transacted quantity (i.e. lastPrice * lastQty) | [optional] 
**Q** | Pointer to **string** | Quote Order Quantity | [optional] 
**W** | Pointer to **int64** | Working Time; This is only visible if the order has been placed on the book. | [optional] 
**V** | Pointer to **string** | SelfTradePreventionMode | [optional] 
**D** | Pointer to **int64** | Trailing Time | [optional] 
**J** | Pointer to **int64** | Strategy Id | [optional] 
**J** | Pointer to **int64** | Strategy Type | [optional] 
**A** | Pointer to **string** | Prevented Quantity | [optional] 
**B** | Pointer to [**[]OutboundAccountPositionBInner**](OutboundAccountPositionBInner.md) | Balances Array | [optional] 
**U** | Pointer to **int64** | Time of last account update | [optional] 
**U** | Pointer to **int64** | Counter Order Id | [optional] 
**Cs** | Pointer to **string** | Counter Symbol | [optional] 
**Pl** | Pointer to **string** | Prevented Execution Quantity | [optional] 
**PL** | Pointer to **string** | Prevented Execution Price | [optional] 
**PY** | Pointer to **string** | Prevented Execution Quote Qty | [optional] 
**B** | Pointer to **string** | Match Type | [optional] 
**K** | Pointer to **string** | Working Floor | [optional] 
**US** | Pointer to **bool** | UsedSor | [optional] 
**GP** | Pointer to **string** | Pegged Price Type | [optional] 
**GOT** | Pointer to **string** | Pegged Offset Type | [optional] 
**GOV** | Pointer to **int64** | Pegged Offset Value | [optional] 
**Gp** | Pointer to **string** | Pegged Price | [optional] 
**ER** | Pointer to **string** | Expiry Reason. Appears when the order has expired. | [optional] 

## Methods

### NewUserDataStreamEventsResponse

`func NewUserDataStreamEventsResponse() *UserDataStreamEventsResponse`

NewUserDataStreamEventsResponse instantiates a new UserDataStreamEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataStreamEventsResponseWithDefaults

`func NewUserDataStreamEventsResponseWithDefaults() *UserDataStreamEventsResponse`

NewUserDataStreamEventsResponseWithDefaults instantiates a new UserDataStreamEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *UserDataStreamEventsResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *UserDataStreamEventsResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *UserDataStreamEventsResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *UserDataStreamEventsResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetA

`func (o *UserDataStreamEventsResponse) GetA() string`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *UserDataStreamEventsResponse) GetAOk() (*string, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *UserDataStreamEventsResponse) SetA(v string)`

SetA sets A field to given value.

### HasA

`func (o *UserDataStreamEventsResponse) HasA() bool`

HasA returns a boolean if a field has been set.

### GetD

`func (o *UserDataStreamEventsResponse) GetD() string`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *UserDataStreamEventsResponse) GetDOk() (*string, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *UserDataStreamEventsResponse) SetD(v string)`

SetD sets D field to given value.

### HasD

`func (o *UserDataStreamEventsResponse) HasD() bool`

HasD returns a boolean if a field has been set.

### GetT

`func (o *UserDataStreamEventsResponse) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UserDataStreamEventsResponse) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UserDataStreamEventsResponse) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *UserDataStreamEventsResponse) HasT() bool`

HasT returns a boolean if a field has been set.

### GetS

`func (o *UserDataStreamEventsResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *UserDataStreamEventsResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *UserDataStreamEventsResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *UserDataStreamEventsResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetC

`func (o *UserDataStreamEventsResponse) GetC() string`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *UserDataStreamEventsResponse) GetCOk() (*string, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *UserDataStreamEventsResponse) SetC(v string)`

SetC sets C field to given value.

### HasC

`func (o *UserDataStreamEventsResponse) HasC() bool`

HasC returns a boolean if a field has been set.

### GetS

`func (o *UserDataStreamEventsResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *UserDataStreamEventsResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *UserDataStreamEventsResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *UserDataStreamEventsResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetO

`func (o *UserDataStreamEventsResponse) GetO() string`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *UserDataStreamEventsResponse) GetOOk() (*string, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *UserDataStreamEventsResponse) SetO(v string)`

SetO sets O field to given value.

### HasO

`func (o *UserDataStreamEventsResponse) HasO() bool`

HasO returns a boolean if a field has been set.

### GetF

`func (o *UserDataStreamEventsResponse) GetF() string`

GetF returns the F field if non-nil, zero value otherwise.

### GetFOk

`func (o *UserDataStreamEventsResponse) GetFOk() (*string, bool)`

GetFOk returns a tuple with the F field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF

`func (o *UserDataStreamEventsResponse) SetF(v string)`

SetF sets F field to given value.

### HasF

`func (o *UserDataStreamEventsResponse) HasF() bool`

HasF returns a boolean if a field has been set.

### GetQ

`func (o *UserDataStreamEventsResponse) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *UserDataStreamEventsResponse) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *UserDataStreamEventsResponse) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *UserDataStreamEventsResponse) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetP

`func (o *UserDataStreamEventsResponse) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *UserDataStreamEventsResponse) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *UserDataStreamEventsResponse) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *UserDataStreamEventsResponse) HasP() bool`

HasP returns a boolean if a field has been set.

### GetP

`func (o *UserDataStreamEventsResponse) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *UserDataStreamEventsResponse) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *UserDataStreamEventsResponse) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *UserDataStreamEventsResponse) HasP() bool`

HasP returns a boolean if a field has been set.

### GetF

`func (o *UserDataStreamEventsResponse) GetF() string`

GetF returns the F field if non-nil, zero value otherwise.

### GetFOk

`func (o *UserDataStreamEventsResponse) GetFOk() (*string, bool)`

GetFOk returns a tuple with the F field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF

`func (o *UserDataStreamEventsResponse) SetF(v string)`

SetF sets F field to given value.

### HasF

`func (o *UserDataStreamEventsResponse) HasF() bool`

HasF returns a boolean if a field has been set.

### GetG

`func (o *UserDataStreamEventsResponse) GetG() int64`

GetG returns the G field if non-nil, zero value otherwise.

### GetGOk

`func (o *UserDataStreamEventsResponse) GetGOk() (*int64, bool)`

GetGOk returns a tuple with the G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetG

`func (o *UserDataStreamEventsResponse) SetG(v int64)`

SetG sets G field to given value.

### HasG

`func (o *UserDataStreamEventsResponse) HasG() bool`

HasG returns a boolean if a field has been set.

### GetC

`func (o *UserDataStreamEventsResponse) GetC() string`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *UserDataStreamEventsResponse) GetCOk() (*string, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *UserDataStreamEventsResponse) SetC(v string)`

SetC sets C field to given value.

### HasC

`func (o *UserDataStreamEventsResponse) HasC() bool`

HasC returns a boolean if a field has been set.

### GetX

`func (o *UserDataStreamEventsResponse) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *UserDataStreamEventsResponse) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *UserDataStreamEventsResponse) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *UserDataStreamEventsResponse) HasX() bool`

HasX returns a boolean if a field has been set.

### GetX

`func (o *UserDataStreamEventsResponse) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *UserDataStreamEventsResponse) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *UserDataStreamEventsResponse) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *UserDataStreamEventsResponse) HasX() bool`

HasX returns a boolean if a field has been set.

### GetR

`func (o *UserDataStreamEventsResponse) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *UserDataStreamEventsResponse) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *UserDataStreamEventsResponse) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *UserDataStreamEventsResponse) HasR() bool`

HasR returns a boolean if a field has been set.

### GetI

`func (o *UserDataStreamEventsResponse) GetI() int64`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *UserDataStreamEventsResponse) GetIOk() (*int64, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *UserDataStreamEventsResponse) SetI(v int64)`

SetI sets I field to given value.

### HasI

`func (o *UserDataStreamEventsResponse) HasI() bool`

HasI returns a boolean if a field has been set.

### GetL

`func (o *UserDataStreamEventsResponse) GetL() string`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *UserDataStreamEventsResponse) GetLOk() (*string, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *UserDataStreamEventsResponse) SetL(v string)`

SetL sets L field to given value.

### HasL

`func (o *UserDataStreamEventsResponse) HasL() bool`

HasL returns a boolean if a field has been set.

### GetZ

`func (o *UserDataStreamEventsResponse) GetZ() string`

GetZ returns the Z field if non-nil, zero value otherwise.

### GetZOk

`func (o *UserDataStreamEventsResponse) GetZOk() (*string, bool)`

GetZOk returns a tuple with the Z field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZ

`func (o *UserDataStreamEventsResponse) SetZ(v string)`

SetZ sets Z field to given value.

### HasZ

`func (o *UserDataStreamEventsResponse) HasZ() bool`

HasZ returns a boolean if a field has been set.

### GetL

`func (o *UserDataStreamEventsResponse) GetL() string`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *UserDataStreamEventsResponse) GetLOk() (*string, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *UserDataStreamEventsResponse) SetL(v string)`

SetL sets L field to given value.

### HasL

`func (o *UserDataStreamEventsResponse) HasL() bool`

HasL returns a boolean if a field has been set.

### GetN

`func (o *UserDataStreamEventsResponse) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *UserDataStreamEventsResponse) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *UserDataStreamEventsResponse) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *UserDataStreamEventsResponse) HasN() bool`

HasN returns a boolean if a field has been set.

### GetN

`func (o *UserDataStreamEventsResponse) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *UserDataStreamEventsResponse) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *UserDataStreamEventsResponse) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *UserDataStreamEventsResponse) HasN() bool`

HasN returns a boolean if a field has been set.

### GetT

`func (o *UserDataStreamEventsResponse) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UserDataStreamEventsResponse) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UserDataStreamEventsResponse) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *UserDataStreamEventsResponse) HasT() bool`

HasT returns a boolean if a field has been set.

### GetV

`func (o *UserDataStreamEventsResponse) GetV() int64`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *UserDataStreamEventsResponse) GetVOk() (*int64, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *UserDataStreamEventsResponse) SetV(v int64)`

SetV sets V field to given value.

### HasV

`func (o *UserDataStreamEventsResponse) HasV() bool`

HasV returns a boolean if a field has been set.

### GetI

`func (o *UserDataStreamEventsResponse) GetI() int64`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *UserDataStreamEventsResponse) GetIOk() (*int64, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *UserDataStreamEventsResponse) SetI(v int64)`

SetI sets I field to given value.

### HasI

`func (o *UserDataStreamEventsResponse) HasI() bool`

HasI returns a boolean if a field has been set.

### GetW

`func (o *UserDataStreamEventsResponse) GetW() bool`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *UserDataStreamEventsResponse) GetWOk() (*bool, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *UserDataStreamEventsResponse) SetW(v bool)`

SetW sets W field to given value.

### HasW

`func (o *UserDataStreamEventsResponse) HasW() bool`

HasW returns a boolean if a field has been set.

### GetM

`func (o *UserDataStreamEventsResponse) GetM() bool`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *UserDataStreamEventsResponse) GetMOk() (*bool, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *UserDataStreamEventsResponse) SetM(v bool)`

SetM sets M field to given value.

### HasM

`func (o *UserDataStreamEventsResponse) HasM() bool`

HasM returns a boolean if a field has been set.

### GetM

`func (o *UserDataStreamEventsResponse) GetM() bool`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *UserDataStreamEventsResponse) GetMOk() (*bool, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *UserDataStreamEventsResponse) SetM(v bool)`

SetM sets M field to given value.

### HasM

`func (o *UserDataStreamEventsResponse) HasM() bool`

HasM returns a boolean if a field has been set.

### GetO

`func (o *UserDataStreamEventsResponse) GetO() []ListStatusOInner`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *UserDataStreamEventsResponse) GetOOk() (*[]ListStatusOInner, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *UserDataStreamEventsResponse) SetO(v []ListStatusOInner)`

SetO sets O field to given value.

### HasO

`func (o *UserDataStreamEventsResponse) HasO() bool`

HasO returns a boolean if a field has been set.

### GetZ

`func (o *UserDataStreamEventsResponse) GetZ() string`

GetZ returns the Z field if non-nil, zero value otherwise.

### GetZOk

`func (o *UserDataStreamEventsResponse) GetZOk() (*string, bool)`

GetZOk returns a tuple with the Z field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZ

`func (o *UserDataStreamEventsResponse) SetZ(v string)`

SetZ sets Z field to given value.

### HasZ

`func (o *UserDataStreamEventsResponse) HasZ() bool`

HasZ returns a boolean if a field has been set.

### GetY

`func (o *UserDataStreamEventsResponse) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *UserDataStreamEventsResponse) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *UserDataStreamEventsResponse) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *UserDataStreamEventsResponse) HasY() bool`

HasY returns a boolean if a field has been set.

### GetQ

`func (o *UserDataStreamEventsResponse) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *UserDataStreamEventsResponse) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *UserDataStreamEventsResponse) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *UserDataStreamEventsResponse) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetW

`func (o *UserDataStreamEventsResponse) GetW() int64`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *UserDataStreamEventsResponse) GetWOk() (*int64, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *UserDataStreamEventsResponse) SetW(v int64)`

SetW sets W field to given value.

### HasW

`func (o *UserDataStreamEventsResponse) HasW() bool`

HasW returns a boolean if a field has been set.

### GetV

`func (o *UserDataStreamEventsResponse) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *UserDataStreamEventsResponse) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *UserDataStreamEventsResponse) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *UserDataStreamEventsResponse) HasV() bool`

HasV returns a boolean if a field has been set.

### GetD

`func (o *UserDataStreamEventsResponse) GetD() int64`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *UserDataStreamEventsResponse) GetDOk() (*int64, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *UserDataStreamEventsResponse) SetD(v int64)`

SetD sets D field to given value.

### HasD

`func (o *UserDataStreamEventsResponse) HasD() bool`

HasD returns a boolean if a field has been set.

### GetJ

`func (o *UserDataStreamEventsResponse) GetJ() int64`

GetJ returns the J field if non-nil, zero value otherwise.

### GetJOk

`func (o *UserDataStreamEventsResponse) GetJOk() (*int64, bool)`

GetJOk returns a tuple with the J field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJ

`func (o *UserDataStreamEventsResponse) SetJ(v int64)`

SetJ sets J field to given value.

### HasJ

`func (o *UserDataStreamEventsResponse) HasJ() bool`

HasJ returns a boolean if a field has been set.

### GetJ

`func (o *UserDataStreamEventsResponse) GetJ() int64`

GetJ returns the J field if non-nil, zero value otherwise.

### GetJOk

`func (o *UserDataStreamEventsResponse) GetJOk() (*int64, bool)`

GetJOk returns a tuple with the J field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJ

`func (o *UserDataStreamEventsResponse) SetJ(v int64)`

SetJ sets J field to given value.

### HasJ

`func (o *UserDataStreamEventsResponse) HasJ() bool`

HasJ returns a boolean if a field has been set.

### GetA

`func (o *UserDataStreamEventsResponse) GetA() string`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *UserDataStreamEventsResponse) GetAOk() (*string, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *UserDataStreamEventsResponse) SetA(v string)`

SetA sets A field to given value.

### HasA

`func (o *UserDataStreamEventsResponse) HasA() bool`

HasA returns a boolean if a field has been set.

### GetB

`func (o *UserDataStreamEventsResponse) GetB() []OutboundAccountPositionBInner`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *UserDataStreamEventsResponse) GetBOk() (*[]OutboundAccountPositionBInner, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *UserDataStreamEventsResponse) SetB(v []OutboundAccountPositionBInner)`

SetB sets B field to given value.

### HasB

`func (o *UserDataStreamEventsResponse) HasB() bool`

HasB returns a boolean if a field has been set.

### GetU

`func (o *UserDataStreamEventsResponse) GetU() int64`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *UserDataStreamEventsResponse) GetUOk() (*int64, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *UserDataStreamEventsResponse) SetU(v int64)`

SetU sets U field to given value.

### HasU

`func (o *UserDataStreamEventsResponse) HasU() bool`

HasU returns a boolean if a field has been set.

### GetU

`func (o *UserDataStreamEventsResponse) GetU() int64`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *UserDataStreamEventsResponse) GetUOk() (*int64, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *UserDataStreamEventsResponse) SetU(v int64)`

SetU sets U field to given value.

### HasU

`func (o *UserDataStreamEventsResponse) HasU() bool`

HasU returns a boolean if a field has been set.

### GetCs

`func (o *UserDataStreamEventsResponse) GetCs() string`

GetCs returns the Cs field if non-nil, zero value otherwise.

### GetCsOk

`func (o *UserDataStreamEventsResponse) GetCsOk() (*string, bool)`

GetCsOk returns a tuple with the Cs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCs

`func (o *UserDataStreamEventsResponse) SetCs(v string)`

SetCs sets Cs field to given value.

### HasCs

`func (o *UserDataStreamEventsResponse) HasCs() bool`

HasCs returns a boolean if a field has been set.

### GetPl

`func (o *UserDataStreamEventsResponse) GetPl() string`

GetPl returns the Pl field if non-nil, zero value otherwise.

### GetPlOk

`func (o *UserDataStreamEventsResponse) GetPlOk() (*string, bool)`

GetPlOk returns a tuple with the Pl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPl

`func (o *UserDataStreamEventsResponse) SetPl(v string)`

SetPl sets Pl field to given value.

### HasPl

`func (o *UserDataStreamEventsResponse) HasPl() bool`

HasPl returns a boolean if a field has been set.

### GetPL

`func (o *UserDataStreamEventsResponse) GetPL() string`

GetPL returns the PL field if non-nil, zero value otherwise.

### GetPLOk

`func (o *UserDataStreamEventsResponse) GetPLOk() (*string, bool)`

GetPLOk returns a tuple with the PL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPL

`func (o *UserDataStreamEventsResponse) SetPL(v string)`

SetPL sets PL field to given value.

### HasPL

`func (o *UserDataStreamEventsResponse) HasPL() bool`

HasPL returns a boolean if a field has been set.

### GetPY

`func (o *UserDataStreamEventsResponse) GetPY() string`

GetPY returns the PY field if non-nil, zero value otherwise.

### GetPYOk

`func (o *UserDataStreamEventsResponse) GetPYOk() (*string, bool)`

GetPYOk returns a tuple with the PY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPY

`func (o *UserDataStreamEventsResponse) SetPY(v string)`

SetPY sets PY field to given value.

### HasPY

`func (o *UserDataStreamEventsResponse) HasPY() bool`

HasPY returns a boolean if a field has been set.

### GetB

`func (o *UserDataStreamEventsResponse) GetB() string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *UserDataStreamEventsResponse) GetBOk() (*string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *UserDataStreamEventsResponse) SetB(v string)`

SetB sets B field to given value.

### HasB

`func (o *UserDataStreamEventsResponse) HasB() bool`

HasB returns a boolean if a field has been set.

### GetK

`func (o *UserDataStreamEventsResponse) GetK() string`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *UserDataStreamEventsResponse) GetKOk() (*string, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *UserDataStreamEventsResponse) SetK(v string)`

SetK sets K field to given value.

### HasK

`func (o *UserDataStreamEventsResponse) HasK() bool`

HasK returns a boolean if a field has been set.

### GetUS

`func (o *UserDataStreamEventsResponse) GetUS() bool`

GetUS returns the US field if non-nil, zero value otherwise.

### GetUSOk

`func (o *UserDataStreamEventsResponse) GetUSOk() (*bool, bool)`

GetUSOk returns a tuple with the US field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUS

`func (o *UserDataStreamEventsResponse) SetUS(v bool)`

SetUS sets US field to given value.

### HasUS

`func (o *UserDataStreamEventsResponse) HasUS() bool`

HasUS returns a boolean if a field has been set.

### GetGP

`func (o *UserDataStreamEventsResponse) GetGP() string`

GetGP returns the GP field if non-nil, zero value otherwise.

### GetGPOk

`func (o *UserDataStreamEventsResponse) GetGPOk() (*string, bool)`

GetGPOk returns a tuple with the GP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGP

`func (o *UserDataStreamEventsResponse) SetGP(v string)`

SetGP sets GP field to given value.

### HasGP

`func (o *UserDataStreamEventsResponse) HasGP() bool`

HasGP returns a boolean if a field has been set.

### GetGOT

`func (o *UserDataStreamEventsResponse) GetGOT() string`

GetGOT returns the GOT field if non-nil, zero value otherwise.

### GetGOTOk

`func (o *UserDataStreamEventsResponse) GetGOTOk() (*string, bool)`

GetGOTOk returns a tuple with the GOT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGOT

`func (o *UserDataStreamEventsResponse) SetGOT(v string)`

SetGOT sets GOT field to given value.

### HasGOT

`func (o *UserDataStreamEventsResponse) HasGOT() bool`

HasGOT returns a boolean if a field has been set.

### GetGOV

`func (o *UserDataStreamEventsResponse) GetGOV() int64`

GetGOV returns the GOV field if non-nil, zero value otherwise.

### GetGOVOk

`func (o *UserDataStreamEventsResponse) GetGOVOk() (*int64, bool)`

GetGOVOk returns a tuple with the GOV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGOV

`func (o *UserDataStreamEventsResponse) SetGOV(v int64)`

SetGOV sets GOV field to given value.

### HasGOV

`func (o *UserDataStreamEventsResponse) HasGOV() bool`

HasGOV returns a boolean if a field has been set.

### GetGp

`func (o *UserDataStreamEventsResponse) GetGp() string`

GetGp returns the Gp field if non-nil, zero value otherwise.

### GetGpOk

`func (o *UserDataStreamEventsResponse) GetGpOk() (*string, bool)`

GetGpOk returns a tuple with the Gp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGp

`func (o *UserDataStreamEventsResponse) SetGp(v string)`

SetGp sets Gp field to given value.

### HasGp

`func (o *UserDataStreamEventsResponse) HasGp() bool`

HasGp returns a boolean if a field has been set.

### GetER

`func (o *UserDataStreamEventsResponse) GetER() string`

GetER returns the ER field if non-nil, zero value otherwise.

### GetEROk

`func (o *UserDataStreamEventsResponse) GetEROk() (*string, bool)`

GetEROk returns a tuple with the ER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetER

`func (o *UserDataStreamEventsResponse) SetER(v string)`

SetER sets ER field to given value.

### HasER

`func (o *UserDataStreamEventsResponse) HasER() bool`

HasER returns a boolean if a field has been set.


[[Back to README]](../README.md)


