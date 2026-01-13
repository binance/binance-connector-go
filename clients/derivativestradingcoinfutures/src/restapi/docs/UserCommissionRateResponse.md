# UserCommissionRateResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**MakerCommissionRate** | Pointer to **string** |  | [optional] 
**TakerCommissionRate** | Pointer to **string** |  | [optional] 

## Methods

### NewUserCommissionRateResponse

`func NewUserCommissionRateResponse() *UserCommissionRateResponse`

NewUserCommissionRateResponse instantiates a new UserCommissionRateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCommissionRateResponseWithDefaults

`func NewUserCommissionRateResponseWithDefaults() *UserCommissionRateResponse`

NewUserCommissionRateResponseWithDefaults instantiates a new UserCommissionRateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *UserCommissionRateResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UserCommissionRateResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UserCommissionRateResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UserCommissionRateResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetMakerCommissionRate

`func (o *UserCommissionRateResponse) GetMakerCommissionRate() string`

GetMakerCommissionRate returns the MakerCommissionRate field if non-nil, zero value otherwise.

### GetMakerCommissionRateOk

`func (o *UserCommissionRateResponse) GetMakerCommissionRateOk() (*string, bool)`

GetMakerCommissionRateOk returns a tuple with the MakerCommissionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerCommissionRate

`func (o *UserCommissionRateResponse) SetMakerCommissionRate(v string)`

SetMakerCommissionRate sets MakerCommissionRate field to given value.

### HasMakerCommissionRate

`func (o *UserCommissionRateResponse) HasMakerCommissionRate() bool`

HasMakerCommissionRate returns a boolean if a field has been set.

### GetTakerCommissionRate

`func (o *UserCommissionRateResponse) GetTakerCommissionRate() string`

GetTakerCommissionRate returns the TakerCommissionRate field if non-nil, zero value otherwise.

### GetTakerCommissionRateOk

`func (o *UserCommissionRateResponse) GetTakerCommissionRateOk() (*string, bool)`

GetTakerCommissionRateOk returns a tuple with the TakerCommissionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerCommissionRate

`func (o *UserCommissionRateResponse) SetTakerCommissionRate(v string)`

SetTakerCommissionRate sets TakerCommissionRate field to given value.

### HasTakerCommissionRate

`func (o *UserCommissionRateResponse) HasTakerCommissionRate() bool`

HasTakerCommissionRate returns a boolean if a field has been set.


[[Back to README]](../README.md)


