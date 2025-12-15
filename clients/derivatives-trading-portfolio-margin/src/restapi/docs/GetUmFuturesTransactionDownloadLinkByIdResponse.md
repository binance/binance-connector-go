# GetUmFuturesTransactionDownloadLinkByIdResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**DownloadId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**S3Link** | Pointer to **string** |  | [optional] 
**Notified** | Pointer to **bool** |  | [optional] 
**ExpirationTimestamp** | Pointer to **int64** |  | [optional] 
**IsExpired** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUmFuturesTransactionDownloadLinkByIdResponse

`func NewGetUmFuturesTransactionDownloadLinkByIdResponse() *GetUmFuturesTransactionDownloadLinkByIdResponse`

NewGetUmFuturesTransactionDownloadLinkByIdResponse instantiates a new GetUmFuturesTransactionDownloadLinkByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUmFuturesTransactionDownloadLinkByIdResponseWithDefaults

`func NewGetUmFuturesTransactionDownloadLinkByIdResponseWithDefaults() *GetUmFuturesTransactionDownloadLinkByIdResponse`

NewGetUmFuturesTransactionDownloadLinkByIdResponseWithDefaults instantiates a new GetUmFuturesTransactionDownloadLinkByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadId

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### GetStatus

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetS3Link

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetS3Link() string`

GetS3Link returns the S3Link field if non-nil, zero value otherwise.

### GetS3LinkOk

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetS3LinkOk() (*string, bool)`

GetS3LinkOk returns a tuple with the S3Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Link

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetS3Link(v string)`

SetS3Link sets S3Link field to given value.

### HasS3Link

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasS3Link() bool`

HasS3Link returns a boolean if a field has been set.

### GetNotified

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetNotified() bool`

GetNotified returns the Notified field if non-nil, zero value otherwise.

### GetNotifiedOk

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetNotifiedOk() (*bool, bool)`

GetNotifiedOk returns a tuple with the Notified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotified

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetNotified(v bool)`

SetNotified sets Notified field to given value.

### HasNotified

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasNotified() bool`

HasNotified returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetExpirationTimestamp() int64`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetExpirationTimestampOk() (*int64, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetExpirationTimestamp(v int64)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetIsExpired

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetIsExpired() string`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) GetIsExpiredOk() (*string, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) SetIsExpired(v string)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *GetUmFuturesTransactionDownloadLinkByIdResponse) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.


[[Back to README]](../README.md)


