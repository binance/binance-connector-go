# GetCountryListResponseCountriesInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | ISO 2-digit country code, lowercase. | [optional] 
**CountryName** | Pointer to **string** | Country display name. | [optional] 
**BlockType** | Pointer to **string** | &#x60;supported&#x60;, &#x60;limited&#x60;, or &#x60;blocked&#x60;. | [optional] 
**DepositAllowed** | Pointer to **bool** | Whether deposit is allowed for this country. | [optional] 
**WithdrawalAllowed** | Pointer to **bool** | Whether withdrawal is allowed for this country. | [optional] 
**HasRegionRestrictions** | Pointer to **bool** | Whether this country has region-level restrictions. | [optional] 

## Methods

### NewGetCountryListResponseCountriesInner

`func NewGetCountryListResponseCountriesInner() *GetCountryListResponseCountriesInner`

NewGetCountryListResponseCountriesInner instantiates a new GetCountryListResponseCountriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCountryListResponseCountriesInnerWithDefaults

`func NewGetCountryListResponseCountriesInnerWithDefaults() *GetCountryListResponseCountriesInner`

NewGetCountryListResponseCountriesInnerWithDefaults instantiates a new GetCountryListResponseCountriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *GetCountryListResponseCountriesInner) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GetCountryListResponseCountriesInner) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GetCountryListResponseCountriesInner) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GetCountryListResponseCountriesInner) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *GetCountryListResponseCountriesInner) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *GetCountryListResponseCountriesInner) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *GetCountryListResponseCountriesInner) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *GetCountryListResponseCountriesInner) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetBlockType

`func (o *GetCountryListResponseCountriesInner) GetBlockType() string`

GetBlockType returns the BlockType field if non-nil, zero value otherwise.

### GetBlockTypeOk

`func (o *GetCountryListResponseCountriesInner) GetBlockTypeOk() (*string, bool)`

GetBlockTypeOk returns a tuple with the BlockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockType

`func (o *GetCountryListResponseCountriesInner) SetBlockType(v string)`

SetBlockType sets BlockType field to given value.

### HasBlockType

`func (o *GetCountryListResponseCountriesInner) HasBlockType() bool`

HasBlockType returns a boolean if a field has been set.

### GetDepositAllowed

`func (o *GetCountryListResponseCountriesInner) GetDepositAllowed() bool`

GetDepositAllowed returns the DepositAllowed field if non-nil, zero value otherwise.

### GetDepositAllowedOk

`func (o *GetCountryListResponseCountriesInner) GetDepositAllowedOk() (*bool, bool)`

GetDepositAllowedOk returns a tuple with the DepositAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAllowed

`func (o *GetCountryListResponseCountriesInner) SetDepositAllowed(v bool)`

SetDepositAllowed sets DepositAllowed field to given value.

### HasDepositAllowed

`func (o *GetCountryListResponseCountriesInner) HasDepositAllowed() bool`

HasDepositAllowed returns a boolean if a field has been set.

### GetWithdrawalAllowed

`func (o *GetCountryListResponseCountriesInner) GetWithdrawalAllowed() bool`

GetWithdrawalAllowed returns the WithdrawalAllowed field if non-nil, zero value otherwise.

### GetWithdrawalAllowedOk

`func (o *GetCountryListResponseCountriesInner) GetWithdrawalAllowedOk() (*bool, bool)`

GetWithdrawalAllowedOk returns a tuple with the WithdrawalAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalAllowed

`func (o *GetCountryListResponseCountriesInner) SetWithdrawalAllowed(v bool)`

SetWithdrawalAllowed sets WithdrawalAllowed field to given value.

### HasWithdrawalAllowed

`func (o *GetCountryListResponseCountriesInner) HasWithdrawalAllowed() bool`

HasWithdrawalAllowed returns a boolean if a field has been set.

### GetHasRegionRestrictions

`func (o *GetCountryListResponseCountriesInner) GetHasRegionRestrictions() bool`

GetHasRegionRestrictions returns the HasRegionRestrictions field if non-nil, zero value otherwise.

### GetHasRegionRestrictionsOk

`func (o *GetCountryListResponseCountriesInner) GetHasRegionRestrictionsOk() (*bool, bool)`

GetHasRegionRestrictionsOk returns a tuple with the HasRegionRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRegionRestrictions

`func (o *GetCountryListResponseCountriesInner) SetHasRegionRestrictions(v bool)`

SetHasRegionRestrictions sets HasRegionRestrictions field to given value.

### HasHasRegionRestrictions

`func (o *GetCountryListResponseCountriesInner) HasHasRegionRestrictions() bool`

HasHasRegionRestrictions returns a boolean if a field has been set.


[[Back to README]](../README.md)


