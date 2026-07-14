# GetRegionListResponseRegionsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**RegionName** | Pointer to **string** | Region/city display name (use this value in questionnaire answers). | [optional] 
**BlockType** | Pointer to **string** | &#x60;supported&#x60;, &#x60;limited&#x60;, or &#x60;blocked&#x60;. | [optional] 
**DepositAllowed** | Pointer to **bool** | Whether deposit is allowed for this region. | [optional] 
**WithdrawalAllowed** | Pointer to **bool** | Whether withdrawal is allowed for this region. | [optional] 

## Methods

### NewGetRegionListResponseRegionsInner

`func NewGetRegionListResponseRegionsInner() *GetRegionListResponseRegionsInner`

NewGetRegionListResponseRegionsInner instantiates a new GetRegionListResponseRegionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRegionListResponseRegionsInnerWithDefaults

`func NewGetRegionListResponseRegionsInnerWithDefaults() *GetRegionListResponseRegionsInner`

NewGetRegionListResponseRegionsInnerWithDefaults instantiates a new GetRegionListResponseRegionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionName

`func (o *GetRegionListResponseRegionsInner) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *GetRegionListResponseRegionsInner) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *GetRegionListResponseRegionsInner) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *GetRegionListResponseRegionsInner) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetBlockType

`func (o *GetRegionListResponseRegionsInner) GetBlockType() string`

GetBlockType returns the BlockType field if non-nil, zero value otherwise.

### GetBlockTypeOk

`func (o *GetRegionListResponseRegionsInner) GetBlockTypeOk() (*string, bool)`

GetBlockTypeOk returns a tuple with the BlockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockType

`func (o *GetRegionListResponseRegionsInner) SetBlockType(v string)`

SetBlockType sets BlockType field to given value.

### HasBlockType

`func (o *GetRegionListResponseRegionsInner) HasBlockType() bool`

HasBlockType returns a boolean if a field has been set.

### GetDepositAllowed

`func (o *GetRegionListResponseRegionsInner) GetDepositAllowed() bool`

GetDepositAllowed returns the DepositAllowed field if non-nil, zero value otherwise.

### GetDepositAllowedOk

`func (o *GetRegionListResponseRegionsInner) GetDepositAllowedOk() (*bool, bool)`

GetDepositAllowedOk returns a tuple with the DepositAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAllowed

`func (o *GetRegionListResponseRegionsInner) SetDepositAllowed(v bool)`

SetDepositAllowed sets DepositAllowed field to given value.

### HasDepositAllowed

`func (o *GetRegionListResponseRegionsInner) HasDepositAllowed() bool`

HasDepositAllowed returns a boolean if a field has been set.

### GetWithdrawalAllowed

`func (o *GetRegionListResponseRegionsInner) GetWithdrawalAllowed() bool`

GetWithdrawalAllowed returns the WithdrawalAllowed field if non-nil, zero value otherwise.

### GetWithdrawalAllowedOk

`func (o *GetRegionListResponseRegionsInner) GetWithdrawalAllowedOk() (*bool, bool)`

GetWithdrawalAllowedOk returns a tuple with the WithdrawalAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalAllowed

`func (o *GetRegionListResponseRegionsInner) SetWithdrawalAllowed(v bool)`

SetWithdrawalAllowed sets WithdrawalAllowed field to given value.

### HasWithdrawalAllowed

`func (o *GetRegionListResponseRegionsInner) HasWithdrawalAllowed() bool`

HasWithdrawalAllowed returns a boolean if a field has been set.


[[Back to README]](../README.md)


