# ExecutionReport

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** | Event time | [optional] 
**S** | Pointer to **string** | Symbol | [optional] 
**C** | Pointer to **string** | Client order ID | [optional] 
**S** | Pointer to **string** | Side | [optional] 
**O** | Pointer to **string** | Order type | [optional] 
**F** | Pointer to **string** | Time in force | [optional] 
**Q** | Pointer to **string** | Order quantity | [optional] 
**P** | Pointer to **string** | Order price | [optional] 
**P** | Pointer to **string** | Stop price | [optional] 
**F** | Pointer to **string** | Iceberg quantity | [optional] 
**G** | Pointer to **int64** | OrderListId | [optional] 
**C** | Pointer to **string** | Original client order ID; This is the ID of the order being canceled | [optional] 
**X** | Pointer to **string** | Current execution type | [optional] 
**X** | Pointer to **string** | Current order status | [optional] 
**R** | Pointer to **string** | Order reject reason; will be an error code. | [optional] 
**I** | Pointer to **int64** | Order ID | [optional] 
**L** | Pointer to **string** | Last executed quantity | [optional] 
**Z** | Pointer to **string** | Cumulative filled quantity | [optional] 
**L** | Pointer to **string** | Last executed price | [optional] 
**N** | Pointer to **string** | Commission amount | [optional] 
**N** | Pointer to **string** | Commission asset | [optional] 
**T** | Pointer to **int64** | Transaction time | [optional] 
**T** | Pointer to **int64** | Trade ID | [optional] 
**I** | Pointer to **int64** | Ignore | [optional] 
**W** | Pointer to **bool** | Is the order on the book? | [optional] 
**M** | Pointer to **bool** | Is this trade the maker side? | [optional] 
**M** | Pointer to **bool** | Ignore | [optional] 
**O** | Pointer to **int64** | Order creation time | [optional] 
**Z** | Pointer to **string** | Cumulative quote asset transacted quantity | [optional] 
**Y** | Pointer to **string** | Last quote asset transacted quantity (i.e. lastPrice * lastQty) | [optional] 
**Q** | Pointer to **string** | Quote Order Quantity | [optional] 
**W** | Pointer to **int64** | Working Time; This is only visible if the order has been placed on the book. | [optional] 
**V** | Pointer to **string** | selfTradePreventionMode | [optional] 
**D** | Pointer to **string** | Trailing Delta; This is only visible if the order was a trailing stop order. | [optional] 
**D** | Pointer to **string** | Trailing Time | [optional] 
**J** | Pointer to **string** | Strategy Id | [optional] 
**J** | Pointer to **string** | Strategy Type | [optional] 
**V** | Pointer to **string** | Prevented Match Id | [optional] 
**A** | Pointer to **string** | Prevented Quantity | [optional] 
**B** | Pointer to **string** | Last Prevented Quantity | [optional] 
**U** | Pointer to **string** | Trade Group Id | [optional] 
**U** | Pointer to **string** | Counter Order Id | [optional] 
**Cs** | Pointer to **string** | Counter Symbol | [optional] 
**Pl** | Pointer to **string** | Prevented Execution Quantity | [optional] 
**PL** | Pointer to **string** | Prevented Execution Price | [optional] 
**PY** | Pointer to **string** | Prevented Execution Quote Qty | [optional] 
**B** | Pointer to **string** | Match Type | [optional] 
**A** | Pointer to **string** | Allocation ID | [optional] 
**K** | Pointer to **string** | Working Floor | [optional] 
**US** | Pointer to **bool** | UsedSor | [optional] 

## Methods

### NewExecutionReport

`func NewExecutionReport() *ExecutionReport`

NewExecutionReport instantiates a new ExecutionReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionReportWithDefaults

`func NewExecutionReportWithDefaults() *ExecutionReport`

NewExecutionReportWithDefaults instantiates a new ExecutionReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *ExecutionReport) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *ExecutionReport) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *ExecutionReport) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *ExecutionReport) HasE() bool`

HasE returns a boolean if a field has been set.

### GetS

`func (o *ExecutionReport) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *ExecutionReport) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *ExecutionReport) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *ExecutionReport) HasS() bool`

HasS returns a boolean if a field has been set.

### GetC

`func (o *ExecutionReport) GetC() string`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *ExecutionReport) GetCOk() (*string, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *ExecutionReport) SetC(v string)`

SetC sets C field to given value.

### HasC

`func (o *ExecutionReport) HasC() bool`

HasC returns a boolean if a field has been set.

### GetS

`func (o *ExecutionReport) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *ExecutionReport) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *ExecutionReport) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *ExecutionReport) HasS() bool`

HasS returns a boolean if a field has been set.

### GetO

`func (o *ExecutionReport) GetO() string`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *ExecutionReport) GetOOk() (*string, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *ExecutionReport) SetO(v string)`

SetO sets O field to given value.

### HasO

`func (o *ExecutionReport) HasO() bool`

HasO returns a boolean if a field has been set.

### GetF

`func (o *ExecutionReport) GetF() string`

GetF returns the F field if non-nil, zero value otherwise.

### GetFOk

`func (o *ExecutionReport) GetFOk() (*string, bool)`

GetFOk returns a tuple with the F field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF

`func (o *ExecutionReport) SetF(v string)`

SetF sets F field to given value.

### HasF

`func (o *ExecutionReport) HasF() bool`

HasF returns a boolean if a field has been set.

### GetQ

`func (o *ExecutionReport) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *ExecutionReport) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *ExecutionReport) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *ExecutionReport) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetP

`func (o *ExecutionReport) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *ExecutionReport) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *ExecutionReport) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *ExecutionReport) HasP() bool`

HasP returns a boolean if a field has been set.

### GetP

`func (o *ExecutionReport) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *ExecutionReport) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *ExecutionReport) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *ExecutionReport) HasP() bool`

HasP returns a boolean if a field has been set.

### GetF

`func (o *ExecutionReport) GetF() string`

GetF returns the F field if non-nil, zero value otherwise.

### GetFOk

`func (o *ExecutionReport) GetFOk() (*string, bool)`

GetFOk returns a tuple with the F field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF

`func (o *ExecutionReport) SetF(v string)`

SetF sets F field to given value.

### HasF

`func (o *ExecutionReport) HasF() bool`

HasF returns a boolean if a field has been set.

### GetG

`func (o *ExecutionReport) GetG() int64`

GetG returns the G field if non-nil, zero value otherwise.

### GetGOk

`func (o *ExecutionReport) GetGOk() (*int64, bool)`

GetGOk returns a tuple with the G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetG

`func (o *ExecutionReport) SetG(v int64)`

SetG sets G field to given value.

### HasG

`func (o *ExecutionReport) HasG() bool`

HasG returns a boolean if a field has been set.

### GetC

`func (o *ExecutionReport) GetC() string`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *ExecutionReport) GetCOk() (*string, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *ExecutionReport) SetC(v string)`

SetC sets C field to given value.

### HasC

`func (o *ExecutionReport) HasC() bool`

HasC returns a boolean if a field has been set.

### GetX

`func (o *ExecutionReport) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ExecutionReport) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ExecutionReport) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *ExecutionReport) HasX() bool`

HasX returns a boolean if a field has been set.

### GetX

`func (o *ExecutionReport) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ExecutionReport) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ExecutionReport) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *ExecutionReport) HasX() bool`

HasX returns a boolean if a field has been set.

### GetR

`func (o *ExecutionReport) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *ExecutionReport) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *ExecutionReport) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *ExecutionReport) HasR() bool`

HasR returns a boolean if a field has been set.

### GetI

`func (o *ExecutionReport) GetI() int64`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *ExecutionReport) GetIOk() (*int64, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *ExecutionReport) SetI(v int64)`

SetI sets I field to given value.

### HasI

`func (o *ExecutionReport) HasI() bool`

HasI returns a boolean if a field has been set.

### GetL

`func (o *ExecutionReport) GetL() string`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *ExecutionReport) GetLOk() (*string, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *ExecutionReport) SetL(v string)`

SetL sets L field to given value.

### HasL

`func (o *ExecutionReport) HasL() bool`

HasL returns a boolean if a field has been set.

### GetZ

`func (o *ExecutionReport) GetZ() string`

GetZ returns the Z field if non-nil, zero value otherwise.

### GetZOk

`func (o *ExecutionReport) GetZOk() (*string, bool)`

GetZOk returns a tuple with the Z field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZ

`func (o *ExecutionReport) SetZ(v string)`

SetZ sets Z field to given value.

### HasZ

`func (o *ExecutionReport) HasZ() bool`

HasZ returns a boolean if a field has been set.

### GetL

`func (o *ExecutionReport) GetL() string`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *ExecutionReport) GetLOk() (*string, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *ExecutionReport) SetL(v string)`

SetL sets L field to given value.

### HasL

`func (o *ExecutionReport) HasL() bool`

HasL returns a boolean if a field has been set.

### GetN

`func (o *ExecutionReport) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *ExecutionReport) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *ExecutionReport) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *ExecutionReport) HasN() bool`

HasN returns a boolean if a field has been set.

### GetN

`func (o *ExecutionReport) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *ExecutionReport) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *ExecutionReport) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *ExecutionReport) HasN() bool`

HasN returns a boolean if a field has been set.

### GetT

`func (o *ExecutionReport) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *ExecutionReport) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *ExecutionReport) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *ExecutionReport) HasT() bool`

HasT returns a boolean if a field has been set.

### GetT

`func (o *ExecutionReport) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *ExecutionReport) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *ExecutionReport) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *ExecutionReport) HasT() bool`

HasT returns a boolean if a field has been set.

### GetI

`func (o *ExecutionReport) GetI() int64`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *ExecutionReport) GetIOk() (*int64, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *ExecutionReport) SetI(v int64)`

SetI sets I field to given value.

### HasI

`func (o *ExecutionReport) HasI() bool`

HasI returns a boolean if a field has been set.

### GetW

`func (o *ExecutionReport) GetW() bool`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *ExecutionReport) GetWOk() (*bool, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *ExecutionReport) SetW(v bool)`

SetW sets W field to given value.

### HasW

`func (o *ExecutionReport) HasW() bool`

HasW returns a boolean if a field has been set.

### GetM

`func (o *ExecutionReport) GetM() bool`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *ExecutionReport) GetMOk() (*bool, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *ExecutionReport) SetM(v bool)`

SetM sets M field to given value.

### HasM

`func (o *ExecutionReport) HasM() bool`

HasM returns a boolean if a field has been set.

### GetM

`func (o *ExecutionReport) GetM() bool`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *ExecutionReport) GetMOk() (*bool, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *ExecutionReport) SetM(v bool)`

SetM sets M field to given value.

### HasM

`func (o *ExecutionReport) HasM() bool`

HasM returns a boolean if a field has been set.

### GetO

`func (o *ExecutionReport) GetO() int64`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *ExecutionReport) GetOOk() (*int64, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *ExecutionReport) SetO(v int64)`

SetO sets O field to given value.

### HasO

`func (o *ExecutionReport) HasO() bool`

HasO returns a boolean if a field has been set.

### GetZ

`func (o *ExecutionReport) GetZ() string`

GetZ returns the Z field if non-nil, zero value otherwise.

### GetZOk

`func (o *ExecutionReport) GetZOk() (*string, bool)`

GetZOk returns a tuple with the Z field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZ

`func (o *ExecutionReport) SetZ(v string)`

SetZ sets Z field to given value.

### HasZ

`func (o *ExecutionReport) HasZ() bool`

HasZ returns a boolean if a field has been set.

### GetY

`func (o *ExecutionReport) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ExecutionReport) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ExecutionReport) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *ExecutionReport) HasY() bool`

HasY returns a boolean if a field has been set.

### GetQ

`func (o *ExecutionReport) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *ExecutionReport) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *ExecutionReport) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *ExecutionReport) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetW

`func (o *ExecutionReport) GetW() int64`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *ExecutionReport) GetWOk() (*int64, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *ExecutionReport) SetW(v int64)`

SetW sets W field to given value.

### HasW

`func (o *ExecutionReport) HasW() bool`

HasW returns a boolean if a field has been set.

### GetV

`func (o *ExecutionReport) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ExecutionReport) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ExecutionReport) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *ExecutionReport) HasV() bool`

HasV returns a boolean if a field has been set.

### GetD

`func (o *ExecutionReport) GetD() string`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *ExecutionReport) GetDOk() (*string, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *ExecutionReport) SetD(v string)`

SetD sets D field to given value.

### HasD

`func (o *ExecutionReport) HasD() bool`

HasD returns a boolean if a field has been set.

### GetD

`func (o *ExecutionReport) GetD() string`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *ExecutionReport) GetDOk() (*string, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *ExecutionReport) SetD(v string)`

SetD sets D field to given value.

### HasD

`func (o *ExecutionReport) HasD() bool`

HasD returns a boolean if a field has been set.

### GetJ

`func (o *ExecutionReport) GetJ() string`

GetJ returns the J field if non-nil, zero value otherwise.

### GetJOk

`func (o *ExecutionReport) GetJOk() (*string, bool)`

GetJOk returns a tuple with the J field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJ

`func (o *ExecutionReport) SetJ(v string)`

SetJ sets J field to given value.

### HasJ

`func (o *ExecutionReport) HasJ() bool`

HasJ returns a boolean if a field has been set.

### GetJ

`func (o *ExecutionReport) GetJ() string`

GetJ returns the J field if non-nil, zero value otherwise.

### GetJOk

`func (o *ExecutionReport) GetJOk() (*string, bool)`

GetJOk returns a tuple with the J field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJ

`func (o *ExecutionReport) SetJ(v string)`

SetJ sets J field to given value.

### HasJ

`func (o *ExecutionReport) HasJ() bool`

HasJ returns a boolean if a field has been set.

### GetV

`func (o *ExecutionReport) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ExecutionReport) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ExecutionReport) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *ExecutionReport) HasV() bool`

HasV returns a boolean if a field has been set.

### GetA

`func (o *ExecutionReport) GetA() string`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *ExecutionReport) GetAOk() (*string, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *ExecutionReport) SetA(v string)`

SetA sets A field to given value.

### HasA

`func (o *ExecutionReport) HasA() bool`

HasA returns a boolean if a field has been set.

### GetB

`func (o *ExecutionReport) GetB() string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *ExecutionReport) GetBOk() (*string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *ExecutionReport) SetB(v string)`

SetB sets B field to given value.

### HasB

`func (o *ExecutionReport) HasB() bool`

HasB returns a boolean if a field has been set.

### GetU

`func (o *ExecutionReport) GetU() string`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *ExecutionReport) GetUOk() (*string, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *ExecutionReport) SetU(v string)`

SetU sets U field to given value.

### HasU

`func (o *ExecutionReport) HasU() bool`

HasU returns a boolean if a field has been set.

### GetU

`func (o *ExecutionReport) GetU() string`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *ExecutionReport) GetUOk() (*string, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *ExecutionReport) SetU(v string)`

SetU sets U field to given value.

### HasU

`func (o *ExecutionReport) HasU() bool`

HasU returns a boolean if a field has been set.

### GetCs

`func (o *ExecutionReport) GetCs() string`

GetCs returns the Cs field if non-nil, zero value otherwise.

### GetCsOk

`func (o *ExecutionReport) GetCsOk() (*string, bool)`

GetCsOk returns a tuple with the Cs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCs

`func (o *ExecutionReport) SetCs(v string)`

SetCs sets Cs field to given value.

### HasCs

`func (o *ExecutionReport) HasCs() bool`

HasCs returns a boolean if a field has been set.

### GetPl

`func (o *ExecutionReport) GetPl() string`

GetPl returns the Pl field if non-nil, zero value otherwise.

### GetPlOk

`func (o *ExecutionReport) GetPlOk() (*string, bool)`

GetPlOk returns a tuple with the Pl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPl

`func (o *ExecutionReport) SetPl(v string)`

SetPl sets Pl field to given value.

### HasPl

`func (o *ExecutionReport) HasPl() bool`

HasPl returns a boolean if a field has been set.

### GetPL

`func (o *ExecutionReport) GetPL() string`

GetPL returns the PL field if non-nil, zero value otherwise.

### GetPLOk

`func (o *ExecutionReport) GetPLOk() (*string, bool)`

GetPLOk returns a tuple with the PL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPL

`func (o *ExecutionReport) SetPL(v string)`

SetPL sets PL field to given value.

### HasPL

`func (o *ExecutionReport) HasPL() bool`

HasPL returns a boolean if a field has been set.

### GetPY

`func (o *ExecutionReport) GetPY() string`

GetPY returns the PY field if non-nil, zero value otherwise.

### GetPYOk

`func (o *ExecutionReport) GetPYOk() (*string, bool)`

GetPYOk returns a tuple with the PY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPY

`func (o *ExecutionReport) SetPY(v string)`

SetPY sets PY field to given value.

### HasPY

`func (o *ExecutionReport) HasPY() bool`

HasPY returns a boolean if a field has been set.

### GetB

`func (o *ExecutionReport) GetB() string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *ExecutionReport) GetBOk() (*string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *ExecutionReport) SetB(v string)`

SetB sets B field to given value.

### HasB

`func (o *ExecutionReport) HasB() bool`

HasB returns a boolean if a field has been set.

### GetA

`func (o *ExecutionReport) GetA() string`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *ExecutionReport) GetAOk() (*string, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *ExecutionReport) SetA(v string)`

SetA sets A field to given value.

### HasA

`func (o *ExecutionReport) HasA() bool`

HasA returns a boolean if a field has been set.

### GetK

`func (o *ExecutionReport) GetK() string`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *ExecutionReport) GetKOk() (*string, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *ExecutionReport) SetK(v string)`

SetK sets K field to given value.

### HasK

`func (o *ExecutionReport) HasK() bool`

HasK returns a boolean if a field has been set.

### GetUS

`func (o *ExecutionReport) GetUS() bool`

GetUS returns the US field if non-nil, zero value otherwise.

### GetUSOk

`func (o *ExecutionReport) GetUSOk() (*bool, bool)`

GetUSOk returns a tuple with the US field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUS

`func (o *ExecutionReport) SetUS(v bool)`

SetUS sets US field to given value.

### HasUS

`func (o *ExecutionReport) HasUS() bool`

HasUS returns a boolean if a field has been set.


[[Back to README]](../README.md)


