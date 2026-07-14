# GetCountryListResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Countries** | Pointer to [**[]GetCountryListResponseCountriesInner**](GetCountryListResponseCountriesInner.md) |  | [optional] 
**LastUpdated** | Pointer to **int64** | Last data update timestamp (epoch milliseconds); 0 if empty. | [optional] 

## Methods

### NewGetCountryListResponse

`func NewGetCountryListResponse() *GetCountryListResponse`

NewGetCountryListResponse instantiates a new GetCountryListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCountryListResponseWithDefaults

`func NewGetCountryListResponseWithDefaults() *GetCountryListResponse`

NewGetCountryListResponseWithDefaults instantiates a new GetCountryListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *GetCountryListResponse) GetCountries() []GetCountryListResponseCountriesInner`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *GetCountryListResponse) GetCountriesOk() (*[]GetCountryListResponseCountriesInner, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *GetCountryListResponse) SetCountries(v []GetCountryListResponseCountriesInner)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *GetCountryListResponse) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetCountryListResponse) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetCountryListResponse) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetCountryListResponse) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetCountryListResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to README]](../README.md)


