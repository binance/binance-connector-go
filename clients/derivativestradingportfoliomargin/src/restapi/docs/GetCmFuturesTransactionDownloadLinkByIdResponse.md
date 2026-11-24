# GetCmFuturesTransactionDownloadLinkByIdResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**DownloadId** | Pointer to **string** | Download ID. | [optional] 
**Status** | Pointer to **string** | Enum：completed，processing | [optional] 
**Url** | Pointer to **string** | The link is mapped to download id | [optional] 
**S3Link** | Pointer to **string** | S3 Link. | [optional] 
**Notified** | Pointer to **bool** | ignore | [optional] 
**ExpirationTimestamp** | Pointer to **int64** | The link would expire after this timestamp | [optional] 
**IsExpired** | Pointer to **string** | Is Expired. | [optional] 

## Methods

### NewGetCmFuturesTransactionDownloadLinkByIdResponse

`func NewGetCmFuturesTransactionDownloadLinkByIdResponse() *GetCmFuturesTransactionDownloadLinkByIdResponse`

NewGetCmFuturesTransactionDownloadLinkByIdResponse instantiates a new GetCmFuturesTransactionDownloadLinkByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCmFuturesTransactionDownloadLinkByIdResponseWithDefaults

`func NewGetCmFuturesTransactionDownloadLinkByIdResponseWithDefaults() *GetCmFuturesTransactionDownloadLinkByIdResponse`

NewGetCmFuturesTransactionDownloadLinkByIdResponseWithDefaults instantiates a new GetCmFuturesTransactionDownloadLinkByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadId

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### GetStatus

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetS3Link

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetS3Link() string`

GetS3Link returns the S3Link field if non-nil, zero value otherwise.

### GetS3LinkOk

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetS3LinkOk() (*string, bool)`

GetS3LinkOk returns a tuple with the S3Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Link

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) SetS3Link(v string)`

SetS3Link sets S3Link field to given value.

### HasS3Link

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) HasS3Link() bool`

HasS3Link returns a boolean if a field has been set.

### GetNotified

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetNotified() bool`

GetNotified returns the Notified field if non-nil, zero value otherwise.

### GetNotifiedOk

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetNotifiedOk() (*bool, bool)`

GetNotifiedOk returns a tuple with the Notified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotified

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) SetNotified(v bool)`

SetNotified sets Notified field to given value.

### HasNotified

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) HasNotified() bool`

HasNotified returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetExpirationTimestamp() int64`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetExpirationTimestampOk() (*int64, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) SetExpirationTimestamp(v int64)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetIsExpired

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetIsExpired() string`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) GetIsExpiredOk() (*string, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) SetIsExpired(v string)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *GetCmFuturesTransactionDownloadLinkByIdResponse) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.


[[Back to README]](../README.md)


