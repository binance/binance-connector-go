# OptionMarginAccountInformationResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to [**[]OptionMarginAccountInformationResponseAssetInner**](OptionMarginAccountInformationResponseAssetInner.md) |  | [optional] 
**Greek** | Pointer to [**[]OptionAccountInformationResponseGreekInner**](OptionAccountInformationResponseGreekInner.md) |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewOptionMarginAccountInformationResponse

`func NewOptionMarginAccountInformationResponse() *OptionMarginAccountInformationResponse`

NewOptionMarginAccountInformationResponse instantiates a new OptionMarginAccountInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionMarginAccountInformationResponseWithDefaults

`func NewOptionMarginAccountInformationResponseWithDefaults() *OptionMarginAccountInformationResponse`

NewOptionMarginAccountInformationResponseWithDefaults instantiates a new OptionMarginAccountInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *OptionMarginAccountInformationResponse) GetAsset() []OptionMarginAccountInformationResponseAssetInner`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *OptionMarginAccountInformationResponse) GetAssetOk() (*[]OptionMarginAccountInformationResponseAssetInner, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *OptionMarginAccountInformationResponse) SetAsset(v []OptionMarginAccountInformationResponseAssetInner)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *OptionMarginAccountInformationResponse) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetGreek

`func (o *OptionMarginAccountInformationResponse) GetGreek() []OptionAccountInformationResponseGreekInner`

GetGreek returns the Greek field if non-nil, zero value otherwise.

### GetGreekOk

`func (o *OptionMarginAccountInformationResponse) GetGreekOk() (*[]OptionAccountInformationResponseGreekInner, bool)`

GetGreekOk returns a tuple with the Greek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreek

`func (o *OptionMarginAccountInformationResponse) SetGreek(v []OptionAccountInformationResponseGreekInner)`

SetGreek sets Greek field to given value.

### HasGreek

`func (o *OptionMarginAccountInformationResponse) HasGreek() bool`

HasGreek returns a boolean if a field has been set.

### GetTime

`func (o *OptionMarginAccountInformationResponse) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OptionMarginAccountInformationResponse) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OptionMarginAccountInformationResponse) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OptionMarginAccountInformationResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


