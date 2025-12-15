# OptionAccountInformationResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to [**[]OptionAccountInformationResponseAssetInner**](OptionAccountInformationResponseAssetInner.md) |  | [optional] 
**Greek** | Pointer to [**[]OptionAccountInformationResponseGreekInner**](OptionAccountInformationResponseGreekInner.md) |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**RiskLevel** | Pointer to **string** |  | [optional] 

## Methods

### NewOptionAccountInformationResponse

`func NewOptionAccountInformationResponse() *OptionAccountInformationResponse`

NewOptionAccountInformationResponse instantiates a new OptionAccountInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionAccountInformationResponseWithDefaults

`func NewOptionAccountInformationResponseWithDefaults() *OptionAccountInformationResponse`

NewOptionAccountInformationResponseWithDefaults instantiates a new OptionAccountInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *OptionAccountInformationResponse) GetAsset() []OptionAccountInformationResponseAssetInner`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *OptionAccountInformationResponse) GetAssetOk() (*[]OptionAccountInformationResponseAssetInner, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *OptionAccountInformationResponse) SetAsset(v []OptionAccountInformationResponseAssetInner)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *OptionAccountInformationResponse) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetGreek

`func (o *OptionAccountInformationResponse) GetGreek() []OptionAccountInformationResponseGreekInner`

GetGreek returns the Greek field if non-nil, zero value otherwise.

### GetGreekOk

`func (o *OptionAccountInformationResponse) GetGreekOk() (*[]OptionAccountInformationResponseGreekInner, bool)`

GetGreekOk returns a tuple with the Greek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreek

`func (o *OptionAccountInformationResponse) SetGreek(v []OptionAccountInformationResponseGreekInner)`

SetGreek sets Greek field to given value.

### HasGreek

`func (o *OptionAccountInformationResponse) HasGreek() bool`

HasGreek returns a boolean if a field has been set.

### GetTime

`func (o *OptionAccountInformationResponse) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OptionAccountInformationResponse) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OptionAccountInformationResponse) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OptionAccountInformationResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetRiskLevel

`func (o *OptionAccountInformationResponse) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *OptionAccountInformationResponse) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *OptionAccountInformationResponse) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *OptionAccountInformationResponse) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.


[[Back to README]](../README.md)


