# AssetDetailResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**CTR** | Pointer to [**AssetDetailResponseCTR**](AssetDetailResponseCTR.md) |  | [optional] 
**SKY** | Pointer to [**AssetDetailResponseSKY**](AssetDetailResponseSKY.md) |  | [optional] 

## Methods

### NewAssetDetailResponse

`func NewAssetDetailResponse() *AssetDetailResponse`

NewAssetDetailResponse instantiates a new AssetDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDetailResponseWithDefaults

`func NewAssetDetailResponseWithDefaults() *AssetDetailResponse`

NewAssetDetailResponseWithDefaults instantiates a new AssetDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCTR

`func (o *AssetDetailResponse) GetCTR() AssetDetailResponseCTR`

GetCTR returns the CTR field if non-nil, zero value otherwise.

### GetCTROk

`func (o *AssetDetailResponse) GetCTROk() (*AssetDetailResponseCTR, bool)`

GetCTROk returns a tuple with the CTR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTR

`func (o *AssetDetailResponse) SetCTR(v AssetDetailResponseCTR)`

SetCTR sets CTR field to given value.

### HasCTR

`func (o *AssetDetailResponse) HasCTR() bool`

HasCTR returns a boolean if a field has been set.

### GetSKY

`func (o *AssetDetailResponse) GetSKY() AssetDetailResponseSKY`

GetSKY returns the SKY field if non-nil, zero value otherwise.

### GetSKYOk

`func (o *AssetDetailResponse) GetSKYOk() (*AssetDetailResponseSKY, bool)`

GetSKYOk returns a tuple with the SKY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSKY

`func (o *AssetDetailResponse) SetSKY(v AssetDetailResponseSKY)`

SetSKY sets SKY field to given value.

### HasSKY

`func (o *AssetDetailResponse) HasSKY() bool`

HasSKY returns a boolean if a field has been set.


[[Back to README]](../README.md)


