# GetRegionListResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | Echoed country code (lowercase). | [optional] 
**Regions** | Pointer to [**[]GetRegionListResponseRegionsInner**](GetRegionListResponseRegionsInner.md) |  | [optional] 
**LastUpdated** | Pointer to **int64** | Last data update timestamp (epoch milliseconds); 0 if empty. | [optional] 

## Methods

### NewGetRegionListResponse

`func NewGetRegionListResponse() *GetRegionListResponse`

NewGetRegionListResponse instantiates a new GetRegionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRegionListResponseWithDefaults

`func NewGetRegionListResponseWithDefaults() *GetRegionListResponse`

NewGetRegionListResponseWithDefaults instantiates a new GetRegionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *GetRegionListResponse) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GetRegionListResponse) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GetRegionListResponse) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GetRegionListResponse) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetRegions

`func (o *GetRegionListResponse) GetRegions() []GetRegionListResponseRegionsInner`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *GetRegionListResponse) GetRegionsOk() (*[]GetRegionListResponseRegionsInner, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *GetRegionListResponse) SetRegions(v []GetRegionListResponseRegionsInner)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *GetRegionListResponse) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetRegionListResponse) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetRegionListResponse) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetRegionListResponse) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetRegionListResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to README]](../README.md)


