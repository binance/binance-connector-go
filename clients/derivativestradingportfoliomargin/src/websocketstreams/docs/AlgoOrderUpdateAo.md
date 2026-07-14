# AlgoOrderUpdateAo

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Caid** | Pointer to **string** | Client Algo Id | [optional] 
**Aid** | Pointer to **int64** | Algo Id | [optional] 
**At** | Pointer to **string** | Algo Type | [optional] 
**O** | Pointer to **string** | Order Type | [optional] 
**S** | Pointer to **string** | Symbol | [optional] 
**S** | Pointer to **string** | Side | [optional] 
**Ps** | Pointer to **string** | Position Side | [optional] 
**F** | Pointer to **string** | Time in Force | [optional] 
**Q** | Pointer to **string** | Quantity | [optional] 
**X** | Pointer to **string** | Algo Status: NEW, CANCELED, TRIGGERING, TRIGGERED, FINISHED, REJECTED, EXPIRED | [optional] 
**Ai** | Pointer to **string** | Actual order ID in matching engine | [optional] 
**Ap** | Pointer to **string** | Avg fill price in matching engine | [optional] 
**Aq** | Pointer to **string** | Executed quantity in matching engine | [optional] 
**Act** | Pointer to **string** | Actual order type in matching engine | [optional] 
**Tp** | Pointer to **string** | Trigger Price | [optional] 
**P** | Pointer to **string** | Order Price | [optional] 
**V** | Pointer to **string** | Self Trade Prevention Mode | [optional] 
**Wt** | Pointer to **string** | Working Type | [optional] 
**Pm** | Pointer to **string** | Price Match | [optional] 
**Cp** | Pointer to **bool** | If Close-All | [optional] 
**PP** | Pointer to **bool** | If price protection is on | [optional] 
**R** | Pointer to **bool** | Is reduce only | [optional] 
**Tt** | Pointer to **int64** | Trigger Time | [optional] 
**Gtd** | Pointer to **int64** | Good Till Date | [optional] 
**Rm** | Pointer to **string** | Algo order failed reason | [optional] 

## Methods

### NewAlgoOrderUpdateAo

`func NewAlgoOrderUpdateAo() *AlgoOrderUpdateAo`

NewAlgoOrderUpdateAo instantiates a new AlgoOrderUpdateAo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlgoOrderUpdateAoWithDefaults

`func NewAlgoOrderUpdateAoWithDefaults() *AlgoOrderUpdateAo`

NewAlgoOrderUpdateAoWithDefaults instantiates a new AlgoOrderUpdateAo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaid

`func (o *AlgoOrderUpdateAo) GetCaid() string`

GetCaid returns the Caid field if non-nil, zero value otherwise.

### GetCaidOk

`func (o *AlgoOrderUpdateAo) GetCaidOk() (*string, bool)`

GetCaidOk returns a tuple with the Caid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaid

`func (o *AlgoOrderUpdateAo) SetCaid(v string)`

SetCaid sets Caid field to given value.

### HasCaid

`func (o *AlgoOrderUpdateAo) HasCaid() bool`

HasCaid returns a boolean if a field has been set.

### GetAid

`func (o *AlgoOrderUpdateAo) GetAid() int64`

GetAid returns the Aid field if non-nil, zero value otherwise.

### GetAidOk

`func (o *AlgoOrderUpdateAo) GetAidOk() (*int64, bool)`

GetAidOk returns a tuple with the Aid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAid

`func (o *AlgoOrderUpdateAo) SetAid(v int64)`

SetAid sets Aid field to given value.

### HasAid

`func (o *AlgoOrderUpdateAo) HasAid() bool`

HasAid returns a boolean if a field has been set.

### GetAt

`func (o *AlgoOrderUpdateAo) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *AlgoOrderUpdateAo) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *AlgoOrderUpdateAo) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *AlgoOrderUpdateAo) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetO

`func (o *AlgoOrderUpdateAo) GetO() string`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *AlgoOrderUpdateAo) GetOOk() (*string, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *AlgoOrderUpdateAo) SetO(v string)`

SetO sets O field to given value.

### HasO

`func (o *AlgoOrderUpdateAo) HasO() bool`

HasO returns a boolean if a field has been set.

### GetS

`func (o *AlgoOrderUpdateAo) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *AlgoOrderUpdateAo) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *AlgoOrderUpdateAo) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *AlgoOrderUpdateAo) HasS() bool`

HasS returns a boolean if a field has been set.

### GetS

`func (o *AlgoOrderUpdateAo) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *AlgoOrderUpdateAo) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *AlgoOrderUpdateAo) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *AlgoOrderUpdateAo) HasS() bool`

HasS returns a boolean if a field has been set.

### GetPs

`func (o *AlgoOrderUpdateAo) GetPs() string`

GetPs returns the Ps field if non-nil, zero value otherwise.

### GetPsOk

`func (o *AlgoOrderUpdateAo) GetPsOk() (*string, bool)`

GetPsOk returns a tuple with the Ps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPs

`func (o *AlgoOrderUpdateAo) SetPs(v string)`

SetPs sets Ps field to given value.

### HasPs

`func (o *AlgoOrderUpdateAo) HasPs() bool`

HasPs returns a boolean if a field has been set.

### GetF

`func (o *AlgoOrderUpdateAo) GetF() string`

GetF returns the F field if non-nil, zero value otherwise.

### GetFOk

`func (o *AlgoOrderUpdateAo) GetFOk() (*string, bool)`

GetFOk returns a tuple with the F field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF

`func (o *AlgoOrderUpdateAo) SetF(v string)`

SetF sets F field to given value.

### HasF

`func (o *AlgoOrderUpdateAo) HasF() bool`

HasF returns a boolean if a field has been set.

### GetQ

`func (o *AlgoOrderUpdateAo) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *AlgoOrderUpdateAo) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *AlgoOrderUpdateAo) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *AlgoOrderUpdateAo) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetX

`func (o *AlgoOrderUpdateAo) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *AlgoOrderUpdateAo) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *AlgoOrderUpdateAo) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *AlgoOrderUpdateAo) HasX() bool`

HasX returns a boolean if a field has been set.

### GetAi

`func (o *AlgoOrderUpdateAo) GetAi() string`

GetAi returns the Ai field if non-nil, zero value otherwise.

### GetAiOk

`func (o *AlgoOrderUpdateAo) GetAiOk() (*string, bool)`

GetAiOk returns a tuple with the Ai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAi

`func (o *AlgoOrderUpdateAo) SetAi(v string)`

SetAi sets Ai field to given value.

### HasAi

`func (o *AlgoOrderUpdateAo) HasAi() bool`

HasAi returns a boolean if a field has been set.

### GetAp

`func (o *AlgoOrderUpdateAo) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *AlgoOrderUpdateAo) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *AlgoOrderUpdateAo) SetAp(v string)`

SetAp sets Ap field to given value.

### HasAp

`func (o *AlgoOrderUpdateAo) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetAq

`func (o *AlgoOrderUpdateAo) GetAq() string`

GetAq returns the Aq field if non-nil, zero value otherwise.

### GetAqOk

`func (o *AlgoOrderUpdateAo) GetAqOk() (*string, bool)`

GetAqOk returns a tuple with the Aq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAq

`func (o *AlgoOrderUpdateAo) SetAq(v string)`

SetAq sets Aq field to given value.

### HasAq

`func (o *AlgoOrderUpdateAo) HasAq() bool`

HasAq returns a boolean if a field has been set.

### GetAct

`func (o *AlgoOrderUpdateAo) GetAct() string`

GetAct returns the Act field if non-nil, zero value otherwise.

### GetActOk

`func (o *AlgoOrderUpdateAo) GetActOk() (*string, bool)`

GetActOk returns a tuple with the Act field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAct

`func (o *AlgoOrderUpdateAo) SetAct(v string)`

SetAct sets Act field to given value.

### HasAct

`func (o *AlgoOrderUpdateAo) HasAct() bool`

HasAct returns a boolean if a field has been set.

### GetTp

`func (o *AlgoOrderUpdateAo) GetTp() string`

GetTp returns the Tp field if non-nil, zero value otherwise.

### GetTpOk

`func (o *AlgoOrderUpdateAo) GetTpOk() (*string, bool)`

GetTpOk returns a tuple with the Tp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTp

`func (o *AlgoOrderUpdateAo) SetTp(v string)`

SetTp sets Tp field to given value.

### HasTp

`func (o *AlgoOrderUpdateAo) HasTp() bool`

HasTp returns a boolean if a field has been set.

### GetP

`func (o *AlgoOrderUpdateAo) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *AlgoOrderUpdateAo) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *AlgoOrderUpdateAo) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *AlgoOrderUpdateAo) HasP() bool`

HasP returns a boolean if a field has been set.

### GetV

`func (o *AlgoOrderUpdateAo) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *AlgoOrderUpdateAo) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *AlgoOrderUpdateAo) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *AlgoOrderUpdateAo) HasV() bool`

HasV returns a boolean if a field has been set.

### GetWt

`func (o *AlgoOrderUpdateAo) GetWt() string`

GetWt returns the Wt field if non-nil, zero value otherwise.

### GetWtOk

`func (o *AlgoOrderUpdateAo) GetWtOk() (*string, bool)`

GetWtOk returns a tuple with the Wt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWt

`func (o *AlgoOrderUpdateAo) SetWt(v string)`

SetWt sets Wt field to given value.

### HasWt

`func (o *AlgoOrderUpdateAo) HasWt() bool`

HasWt returns a boolean if a field has been set.

### GetPm

`func (o *AlgoOrderUpdateAo) GetPm() string`

GetPm returns the Pm field if non-nil, zero value otherwise.

### GetPmOk

`func (o *AlgoOrderUpdateAo) GetPmOk() (*string, bool)`

GetPmOk returns a tuple with the Pm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm

`func (o *AlgoOrderUpdateAo) SetPm(v string)`

SetPm sets Pm field to given value.

### HasPm

`func (o *AlgoOrderUpdateAo) HasPm() bool`

HasPm returns a boolean if a field has been set.

### GetCp

`func (o *AlgoOrderUpdateAo) GetCp() bool`

GetCp returns the Cp field if non-nil, zero value otherwise.

### GetCpOk

`func (o *AlgoOrderUpdateAo) GetCpOk() (*bool, bool)`

GetCpOk returns a tuple with the Cp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCp

`func (o *AlgoOrderUpdateAo) SetCp(v bool)`

SetCp sets Cp field to given value.

### HasCp

`func (o *AlgoOrderUpdateAo) HasCp() bool`

HasCp returns a boolean if a field has been set.

### GetPP

`func (o *AlgoOrderUpdateAo) GetPP() bool`

GetPP returns the PP field if non-nil, zero value otherwise.

### GetPPOk

`func (o *AlgoOrderUpdateAo) GetPPOk() (*bool, bool)`

GetPPOk returns a tuple with the PP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPP

`func (o *AlgoOrderUpdateAo) SetPP(v bool)`

SetPP sets PP field to given value.

### HasPP

`func (o *AlgoOrderUpdateAo) HasPP() bool`

HasPP returns a boolean if a field has been set.

### GetR

`func (o *AlgoOrderUpdateAo) GetR() bool`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *AlgoOrderUpdateAo) GetROk() (*bool, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *AlgoOrderUpdateAo) SetR(v bool)`

SetR sets R field to given value.

### HasR

`func (o *AlgoOrderUpdateAo) HasR() bool`

HasR returns a boolean if a field has been set.

### GetTt

`func (o *AlgoOrderUpdateAo) GetTt() int64`

GetTt returns the Tt field if non-nil, zero value otherwise.

### GetTtOk

`func (o *AlgoOrderUpdateAo) GetTtOk() (*int64, bool)`

GetTtOk returns a tuple with the Tt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTt

`func (o *AlgoOrderUpdateAo) SetTt(v int64)`

SetTt sets Tt field to given value.

### HasTt

`func (o *AlgoOrderUpdateAo) HasTt() bool`

HasTt returns a boolean if a field has been set.

### GetGtd

`func (o *AlgoOrderUpdateAo) GetGtd() int64`

GetGtd returns the Gtd field if non-nil, zero value otherwise.

### GetGtdOk

`func (o *AlgoOrderUpdateAo) GetGtdOk() (*int64, bool)`

GetGtdOk returns a tuple with the Gtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtd

`func (o *AlgoOrderUpdateAo) SetGtd(v int64)`

SetGtd sets Gtd field to given value.

### HasGtd

`func (o *AlgoOrderUpdateAo) HasGtd() bool`

HasGtd returns a boolean if a field has been set.

### GetRm

`func (o *AlgoOrderUpdateAo) GetRm() string`

GetRm returns the Rm field if non-nil, zero value otherwise.

### GetRmOk

`func (o *AlgoOrderUpdateAo) GetRmOk() (*string, bool)`

GetRmOk returns a tuple with the Rm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRm

`func (o *AlgoOrderUpdateAo) SetRm(v string)`

SetRm sets Rm field to given value.

### HasRm

`func (o *AlgoOrderUpdateAo) HasRm() bool`

HasRm returns a boolean if a field has been set.


[[Back to README]](../README.md)


