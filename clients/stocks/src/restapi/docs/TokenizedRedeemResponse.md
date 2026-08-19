# TokenizedRedeemResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**IssuerRequestId** | Pointer to **string** | Server-assigned convert request id. Use it to poll &#x60;/tokenized/convert-status&#x60;. | [optional] 
**Status** | Pointer to **string** | Convert status code: &#x60;P&#x60; &#x3D; processing, &#x60;S&#x60; &#x3D; success, &#x60;F&#x60; &#x3D; failed. | [optional] 

## Methods

### NewTokenizedRedeemResponse

`func NewTokenizedRedeemResponse() *TokenizedRedeemResponse`

NewTokenizedRedeemResponse instantiates a new TokenizedRedeemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizedRedeemResponseWithDefaults

`func NewTokenizedRedeemResponseWithDefaults() *TokenizedRedeemResponse`

NewTokenizedRedeemResponseWithDefaults instantiates a new TokenizedRedeemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerRequestId

`func (o *TokenizedRedeemResponse) GetIssuerRequestId() string`

GetIssuerRequestId returns the IssuerRequestId field if non-nil, zero value otherwise.

### GetIssuerRequestIdOk

`func (o *TokenizedRedeemResponse) GetIssuerRequestIdOk() (*string, bool)`

GetIssuerRequestIdOk returns a tuple with the IssuerRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerRequestId

`func (o *TokenizedRedeemResponse) SetIssuerRequestId(v string)`

SetIssuerRequestId sets IssuerRequestId field to given value.

### HasIssuerRequestId

`func (o *TokenizedRedeemResponse) HasIssuerRequestId() bool`

HasIssuerRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *TokenizedRedeemResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenizedRedeemResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenizedRedeemResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TokenizedRedeemResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to README]](../README.md)


