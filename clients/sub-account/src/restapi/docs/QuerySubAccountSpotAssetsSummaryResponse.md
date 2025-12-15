# QuerySubAccountSpotAssetsSummaryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**MasterAccountTotalAsset** | Pointer to **string** |  | [optional] 
**SpotSubUserAssetBtcVoList** | Pointer to [**[]QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner**](QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner.md) |  | [optional] 

## Methods

### NewQuerySubAccountSpotAssetsSummaryResponse

`func NewQuerySubAccountSpotAssetsSummaryResponse() *QuerySubAccountSpotAssetsSummaryResponse`

NewQuerySubAccountSpotAssetsSummaryResponse instantiates a new QuerySubAccountSpotAssetsSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySubAccountSpotAssetsSummaryResponseWithDefaults

`func NewQuerySubAccountSpotAssetsSummaryResponseWithDefaults() *QuerySubAccountSpotAssetsSummaryResponse`

NewQuerySubAccountSpotAssetsSummaryResponseWithDefaults instantiates a new QuerySubAccountSpotAssetsSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *QuerySubAccountSpotAssetsSummaryResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *QuerySubAccountSpotAssetsSummaryResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *QuerySubAccountSpotAssetsSummaryResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *QuerySubAccountSpotAssetsSummaryResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetMasterAccountTotalAsset

`func (o *QuerySubAccountSpotAssetsSummaryResponse) GetMasterAccountTotalAsset() string`

GetMasterAccountTotalAsset returns the MasterAccountTotalAsset field if non-nil, zero value otherwise.

### GetMasterAccountTotalAssetOk

`func (o *QuerySubAccountSpotAssetsSummaryResponse) GetMasterAccountTotalAssetOk() (*string, bool)`

GetMasterAccountTotalAssetOk returns a tuple with the MasterAccountTotalAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterAccountTotalAsset

`func (o *QuerySubAccountSpotAssetsSummaryResponse) SetMasterAccountTotalAsset(v string)`

SetMasterAccountTotalAsset sets MasterAccountTotalAsset field to given value.

### HasMasterAccountTotalAsset

`func (o *QuerySubAccountSpotAssetsSummaryResponse) HasMasterAccountTotalAsset() bool`

HasMasterAccountTotalAsset returns a boolean if a field has been set.

### GetSpotSubUserAssetBtcVoList

`func (o *QuerySubAccountSpotAssetsSummaryResponse) GetSpotSubUserAssetBtcVoList() []QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner`

GetSpotSubUserAssetBtcVoList returns the SpotSubUserAssetBtcVoList field if non-nil, zero value otherwise.

### GetSpotSubUserAssetBtcVoListOk

`func (o *QuerySubAccountSpotAssetsSummaryResponse) GetSpotSubUserAssetBtcVoListOk() (*[]QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner, bool)`

GetSpotSubUserAssetBtcVoListOk returns a tuple with the SpotSubUserAssetBtcVoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotSubUserAssetBtcVoList

`func (o *QuerySubAccountSpotAssetsSummaryResponse) SetSpotSubUserAssetBtcVoList(v []QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner)`

SetSpotSubUserAssetBtcVoList sets SpotSubUserAssetBtcVoList field to given value.

### HasSpotSubUserAssetBtcVoList

`func (o *QuerySubAccountSpotAssetsSummaryResponse) HasSpotSubUserAssetBtcVoList() bool`

HasSpotSubUserAssetBtcVoList returns a boolean if a field has been set.


[[Back to README]](../README.md)


