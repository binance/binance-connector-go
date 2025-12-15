# ChangeInitialLeverageResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Leverage** | Pointer to **int64** |  | [optional] 
**MaxNotionalValue** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewChangeInitialLeverageResponse

`func NewChangeInitialLeverageResponse() *ChangeInitialLeverageResponse`

NewChangeInitialLeverageResponse instantiates a new ChangeInitialLeverageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeInitialLeverageResponseWithDefaults

`func NewChangeInitialLeverageResponseWithDefaults() *ChangeInitialLeverageResponse`

NewChangeInitialLeverageResponseWithDefaults instantiates a new ChangeInitialLeverageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeverage

`func (o *ChangeInitialLeverageResponse) GetLeverage() int64`

GetLeverage returns the Leverage field if non-nil, zero value otherwise.

### GetLeverageOk

`func (o *ChangeInitialLeverageResponse) GetLeverageOk() (*int64, bool)`

GetLeverageOk returns a tuple with the Leverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverage

`func (o *ChangeInitialLeverageResponse) SetLeverage(v int64)`

SetLeverage sets Leverage field to given value.

### HasLeverage

`func (o *ChangeInitialLeverageResponse) HasLeverage() bool`

HasLeverage returns a boolean if a field has been set.

### GetMaxNotionalValue

`func (o *ChangeInitialLeverageResponse) GetMaxNotionalValue() string`

GetMaxNotionalValue returns the MaxNotionalValue field if non-nil, zero value otherwise.

### GetMaxNotionalValueOk

`func (o *ChangeInitialLeverageResponse) GetMaxNotionalValueOk() (*string, bool)`

GetMaxNotionalValueOk returns a tuple with the MaxNotionalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNotionalValue

`func (o *ChangeInitialLeverageResponse) SetMaxNotionalValue(v string)`

SetMaxNotionalValue sets MaxNotionalValue field to given value.

### HasMaxNotionalValue

`func (o *ChangeInitialLeverageResponse) HasMaxNotionalValue() bool`

HasMaxNotionalValue returns a boolean if a field has been set.

### GetSymbol

`func (o *ChangeInitialLeverageResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ChangeInitialLeverageResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ChangeInitialLeverageResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ChangeInitialLeverageResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to README]](../README.md)


