# UserDataStreamEventsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **int64** |  | [optional] 
**E** | Pointer to **int64** |  | [optional] 
**Fs** | Pointer to **string** |  | [optional] 
**So** | Pointer to [**ConditionalOrderTradeUpdateSo**](ConditionalOrderTradeUpdateSo.md) |  | [optional] 
**Ac** | Pointer to [**AccountConfigUpdateAc**](AccountConfigUpdateAc.md) |  | [optional] 
**I** | Pointer to **int64** |  | [optional] 
**A** | Pointer to **string** |  | [optional] 
**O** | Pointer to **string** |  | [optional] 
**T** | Pointer to **int64** |  | [optional] 
**P** | Pointer to **string** |  | [optional] 
**L** | Pointer to **string** |  | [optional] 
**U** | Pointer to **string** |  | [optional] 
**U** | Pointer to **int64** |  | [optional] 
**B** | Pointer to **string** |  | [optional] 
**D** | Pointer to **int64** |  | [optional] 
**S** | Pointer to **string** |  | [optional] 
**C** | Pointer to **string** |  | [optional] 
**S** | Pointer to **string** |  | [optional] 
**F** | Pointer to **string** |  | [optional] 
**Q** | Pointer to **string** |  | [optional] 
**P** | Pointer to **string** |  | [optional] 
**F** | Pointer to **string** |  | [optional] 
**G** | Pointer to **int64** |  | [optional] 
**C** | Pointer to **string** |  | [optional] 
**X** | Pointer to **string** |  | [optional] 
**X** | Pointer to **string** |  | [optional] 
**R** | Pointer to **string** |  | [optional] 
**Z** | Pointer to **string** |  | [optional] 
**L** | Pointer to **string** |  | [optional] 
**N** | Pointer to **string** |  | [optional] 
**N** | Pointer to **string** |  | [optional] 
**V** | Pointer to **int64** |  | [optional] 
**I** | Pointer to **int64** |  | [optional] 
**W** | Pointer to **bool** |  | [optional] 
**M** | Pointer to **string** |  | [optional] 
**O** | Pointer to [**[]OpenorderlossOInner**](OpenorderlossOInner.md) |  | [optional] 
**Z** | Pointer to **string** |  | [optional] 
**Y** | Pointer to **string** |  | [optional] 
**Q** | Pointer to **string** |  | [optional] 
**D** | Pointer to **int64** |  | [optional] 
**J** | Pointer to **int64** |  | [optional] 
**J** | Pointer to **int64** |  | [optional] 
**W** | Pointer to **int64** |  | [optional] 
**V** | Pointer to **string** |  | [optional] 
**A** | Pointer to **string** |  | [optional] 
**Eq** | Pointer to **string** |  | [optional] 
**Ae** | Pointer to **string** |  | [optional] 

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

### GetFs

`func (o *UserDataStreamEventsResponse) GetFs() string`

GetFs returns the Fs field if non-nil, zero value otherwise.

### GetFsOk

`func (o *UserDataStreamEventsResponse) GetFsOk() (*string, bool)`

GetFsOk returns a tuple with the Fs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFs

`func (o *UserDataStreamEventsResponse) SetFs(v string)`

SetFs sets Fs field to given value.

### HasFs

`func (o *UserDataStreamEventsResponse) HasFs() bool`

HasFs returns a boolean if a field has been set.

### GetSo

`func (o *UserDataStreamEventsResponse) GetSo() ConditionalOrderTradeUpdateSo`

GetSo returns the So field if non-nil, zero value otherwise.

### GetSoOk

`func (o *UserDataStreamEventsResponse) GetSoOk() (*ConditionalOrderTradeUpdateSo, bool)`

GetSoOk returns a tuple with the So field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSo

`func (o *UserDataStreamEventsResponse) SetSo(v ConditionalOrderTradeUpdateSo)`

SetSo sets So field to given value.

### HasSo

`func (o *UserDataStreamEventsResponse) HasSo() bool`

HasSo returns a boolean if a field has been set.

### GetAc

`func (o *UserDataStreamEventsResponse) GetAc() AccountConfigUpdateAc`

GetAc returns the Ac field if non-nil, zero value otherwise.

### GetAcOk

`func (o *UserDataStreamEventsResponse) GetAcOk() (*AccountConfigUpdateAc, bool)`

GetAcOk returns a tuple with the Ac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAc

`func (o *UserDataStreamEventsResponse) SetAc(v AccountConfigUpdateAc)`

SetAc sets Ac field to given value.

### HasAc

`func (o *UserDataStreamEventsResponse) HasAc() bool`

HasAc returns a boolean if a field has been set.

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

### GetU

`func (o *UserDataStreamEventsResponse) GetU() string`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *UserDataStreamEventsResponse) GetUOk() (*string, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *UserDataStreamEventsResponse) SetU(v string)`

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

`func (o *UserDataStreamEventsResponse) GetM() string`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *UserDataStreamEventsResponse) GetMOk() (*string, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *UserDataStreamEventsResponse) SetM(v string)`

SetM sets M field to given value.

### HasM

`func (o *UserDataStreamEventsResponse) HasM() bool`

HasM returns a boolean if a field has been set.

### GetO

`func (o *UserDataStreamEventsResponse) GetO() []OpenorderlossOInner`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *UserDataStreamEventsResponse) GetOOk() (*[]OpenorderlossOInner, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *UserDataStreamEventsResponse) SetO(v []OpenorderlossOInner)`

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

### GetEq

`func (o *UserDataStreamEventsResponse) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *UserDataStreamEventsResponse) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *UserDataStreamEventsResponse) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *UserDataStreamEventsResponse) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetAe

`func (o *UserDataStreamEventsResponse) GetAe() string`

GetAe returns the Ae field if non-nil, zero value otherwise.

### GetAeOk

`func (o *UserDataStreamEventsResponse) GetAeOk() (*string, bool)`

GetAeOk returns a tuple with the Ae field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAe

`func (o *UserDataStreamEventsResponse) SetAe(v string)`

SetAe sets Ae field to given value.

### HasAe

`func (o *UserDataStreamEventsResponse) HasAe() bool`

HasAe returns a boolean if a field has been set.


[[Back to README]](../README.md)


