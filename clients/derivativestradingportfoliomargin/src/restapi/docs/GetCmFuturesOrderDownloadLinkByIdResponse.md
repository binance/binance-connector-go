# GetCmFuturesOrderDownloadLinkByIdResponse

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

### NewGetCmFuturesOrderDownloadLinkByIdResponse

`func NewGetCmFuturesOrderDownloadLinkByIdResponse() *GetCmFuturesOrderDownloadLinkByIdResponse`

NewGetCmFuturesOrderDownloadLinkByIdResponse instantiates a new GetCmFuturesOrderDownloadLinkByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCmFuturesOrderDownloadLinkByIdResponseWithDefaults

`func NewGetCmFuturesOrderDownloadLinkByIdResponseWithDefaults() *GetCmFuturesOrderDownloadLinkByIdResponse`

NewGetCmFuturesOrderDownloadLinkByIdResponseWithDefaults instantiates a new GetCmFuturesOrderDownloadLinkByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadId

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### GetStatus

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetS3Link

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetS3Link() string`

GetS3Link returns the S3Link field if non-nil, zero value otherwise.

### GetS3LinkOk

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetS3LinkOk() (*string, bool)`

GetS3LinkOk returns a tuple with the S3Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Link

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) SetS3Link(v string)`

SetS3Link sets S3Link field to given value.

### HasS3Link

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) HasS3Link() bool`

HasS3Link returns a boolean if a field has been set.

### GetNotified

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetNotified() bool`

GetNotified returns the Notified field if non-nil, zero value otherwise.

### GetNotifiedOk

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetNotifiedOk() (*bool, bool)`

GetNotifiedOk returns a tuple with the Notified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotified

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) SetNotified(v bool)`

SetNotified sets Notified field to given value.

### HasNotified

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) HasNotified() bool`

HasNotified returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetExpirationTimestamp() int64`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetExpirationTimestampOk() (*int64, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) SetExpirationTimestamp(v int64)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetIsExpired

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetIsExpired() string`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) GetIsExpiredOk() (*string, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) SetIsExpired(v string)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *GetCmFuturesOrderDownloadLinkByIdResponse) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.


[[Back to README]](../README.md)


