# QueryUniversalTransferHistoryResponseResultInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TranId** | Pointer to **int64** |  | [optional] 
**FromEmail** | Pointer to **string** |  | [optional] 
**ToEmail** | Pointer to **string** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**CreateTimeStamp** | Pointer to **int64** |  | [optional] 
**FromAccountType** | Pointer to **string** |  | [optional] 
**ToAccountType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ClientTranId** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryUniversalTransferHistoryResponseResultInner

`func NewQueryUniversalTransferHistoryResponseResultInner() *QueryUniversalTransferHistoryResponseResultInner`

NewQueryUniversalTransferHistoryResponseResultInner instantiates a new QueryUniversalTransferHistoryResponseResultInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryUniversalTransferHistoryResponseResultInnerWithDefaults

`func NewQueryUniversalTransferHistoryResponseResultInnerWithDefaults() *QueryUniversalTransferHistoryResponseResultInner`

NewQueryUniversalTransferHistoryResponseResultInnerWithDefaults instantiates a new QueryUniversalTransferHistoryResponseResultInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTranId

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetTranId() int64`

GetTranId returns the TranId field if non-nil, zero value otherwise.

### GetTranIdOk

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetTranIdOk() (*int64, bool)`

GetTranIdOk returns a tuple with the TranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranId

`func (o *QueryUniversalTransferHistoryResponseResultInner) SetTranId(v int64)`

SetTranId sets TranId field to given value.

### HasTranId

`func (o *QueryUniversalTransferHistoryResponseResultInner) HasTranId() bool`

HasTranId returns a boolean if a field has been set.

### GetFromEmail

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *QueryUniversalTransferHistoryResponseResultInner) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *QueryUniversalTransferHistoryResponseResultInner) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetToEmail

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetToEmail() string`

GetToEmail returns the ToEmail field if non-nil, zero value otherwise.

### GetToEmailOk

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetToEmailOk() (*string, bool)`

GetToEmailOk returns a tuple with the ToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEmail

`func (o *QueryUniversalTransferHistoryResponseResultInner) SetToEmail(v string)`

SetToEmail sets ToEmail field to given value.

### HasToEmail

`func (o *QueryUniversalTransferHistoryResponseResultInner) HasToEmail() bool`

HasToEmail returns a boolean if a field has been set.

### GetAsset

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *QueryUniversalTransferHistoryResponseResultInner) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *QueryUniversalTransferHistoryResponseResultInner) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAmount

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QueryUniversalTransferHistoryResponseResultInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QueryUniversalTransferHistoryResponseResultInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreateTimeStamp

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetCreateTimeStamp() int64`

GetCreateTimeStamp returns the CreateTimeStamp field if non-nil, zero value otherwise.

### GetCreateTimeStampOk

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetCreateTimeStampOk() (*int64, bool)`

GetCreateTimeStampOk returns a tuple with the CreateTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStamp

`func (o *QueryUniversalTransferHistoryResponseResultInner) SetCreateTimeStamp(v int64)`

SetCreateTimeStamp sets CreateTimeStamp field to given value.

### HasCreateTimeStamp

`func (o *QueryUniversalTransferHistoryResponseResultInner) HasCreateTimeStamp() bool`

HasCreateTimeStamp returns a boolean if a field has been set.

### GetFromAccountType

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetFromAccountType() string`

GetFromAccountType returns the FromAccountType field if non-nil, zero value otherwise.

### GetFromAccountTypeOk

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetFromAccountTypeOk() (*string, bool)`

GetFromAccountTypeOk returns a tuple with the FromAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccountType

`func (o *QueryUniversalTransferHistoryResponseResultInner) SetFromAccountType(v string)`

SetFromAccountType sets FromAccountType field to given value.

### HasFromAccountType

`func (o *QueryUniversalTransferHistoryResponseResultInner) HasFromAccountType() bool`

HasFromAccountType returns a boolean if a field has been set.

### GetToAccountType

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetToAccountType() string`

GetToAccountType returns the ToAccountType field if non-nil, zero value otherwise.

### GetToAccountTypeOk

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetToAccountTypeOk() (*string, bool)`

GetToAccountTypeOk returns a tuple with the ToAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccountType

`func (o *QueryUniversalTransferHistoryResponseResultInner) SetToAccountType(v string)`

SetToAccountType sets ToAccountType field to given value.

### HasToAccountType

`func (o *QueryUniversalTransferHistoryResponseResultInner) HasToAccountType() bool`

HasToAccountType returns a boolean if a field has been set.

### GetStatus

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryUniversalTransferHistoryResponseResultInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryUniversalTransferHistoryResponseResultInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClientTranId

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetClientTranId() string`

GetClientTranId returns the ClientTranId field if non-nil, zero value otherwise.

### GetClientTranIdOk

`func (o *QueryUniversalTransferHistoryResponseResultInner) GetClientTranIdOk() (*string, bool)`

GetClientTranIdOk returns a tuple with the ClientTranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTranId

`func (o *QueryUniversalTransferHistoryResponseResultInner) SetClientTranId(v string)`

SetClientTranId sets ClientTranId field to given value.

### HasClientTranId

`func (o *QueryUniversalTransferHistoryResponseResultInner) HasClientTranId() bool`

HasClientTranId returns a boolean if a field has been set.


[[Back to README]](../README.md)


