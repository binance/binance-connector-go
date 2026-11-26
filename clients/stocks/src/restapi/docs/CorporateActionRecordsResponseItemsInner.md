# CorporateActionRecordsResponseItemsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique record ID. | [optional] 
**UserId** | Pointer to **int64** | User ID. | [optional] 
**AssetCode** | Pointer to **string** | Asset code, the unique identifier of the stock in the system, for example &#x60;EQ_AAPL&#x60;. Currently all asset codes carry the &#x60;EQ_&#x60; prefix. | [optional] 
**EventType** | Pointer to **string** | Corporate action event type. Values include &#x60;CASH_DIVIDEND&#x60;, &#x60;STOCK_DIVIDEND&#x60;, &#x60;FORWARD_SPLIT&#x60;, &#x60;REVERSE_SPLIT&#x60;, &#x60;SPIN_OFF&#x60;, &#x60;CASH_MERGER&#x60;, &#x60;STOCK_MERGER&#x60;, &#x60;STOCK_AND_CASH_MERGER&#x60;, &#x60;WORTHLESS_REMOVAL&#x60;, &#x60;NAME_CHANGE&#x60; and &#x60;RIGHTS_DISTRIBUTION&#x60;. | [optional] 
**Amount** | Pointer to **string** | Signed change of the asset quantity, including free and frozen balances. | [optional] 
**DbModifyTime** | Pointer to **int64** | Last modification time (ms epoch). | [optional] 

## Methods

### NewCorporateActionRecordsResponseItemsInner

`func NewCorporateActionRecordsResponseItemsInner() *CorporateActionRecordsResponseItemsInner`

NewCorporateActionRecordsResponseItemsInner instantiates a new CorporateActionRecordsResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporateActionRecordsResponseItemsInnerWithDefaults

`func NewCorporateActionRecordsResponseItemsInnerWithDefaults() *CorporateActionRecordsResponseItemsInner`

NewCorporateActionRecordsResponseItemsInnerWithDefaults instantiates a new CorporateActionRecordsResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CorporateActionRecordsResponseItemsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CorporateActionRecordsResponseItemsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CorporateActionRecordsResponseItemsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CorporateActionRecordsResponseItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *CorporateActionRecordsResponseItemsInner) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CorporateActionRecordsResponseItemsInner) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CorporateActionRecordsResponseItemsInner) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CorporateActionRecordsResponseItemsInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAssetCode

`func (o *CorporateActionRecordsResponseItemsInner) GetAssetCode() string`

GetAssetCode returns the AssetCode field if non-nil, zero value otherwise.

### GetAssetCodeOk

`func (o *CorporateActionRecordsResponseItemsInner) GetAssetCodeOk() (*string, bool)`

GetAssetCodeOk returns a tuple with the AssetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCode

`func (o *CorporateActionRecordsResponseItemsInner) SetAssetCode(v string)`

SetAssetCode sets AssetCode field to given value.

### HasAssetCode

`func (o *CorporateActionRecordsResponseItemsInner) HasAssetCode() bool`

HasAssetCode returns a boolean if a field has been set.

### GetEventType

`func (o *CorporateActionRecordsResponseItemsInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CorporateActionRecordsResponseItemsInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CorporateActionRecordsResponseItemsInner) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CorporateActionRecordsResponseItemsInner) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetAmount

`func (o *CorporateActionRecordsResponseItemsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CorporateActionRecordsResponseItemsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CorporateActionRecordsResponseItemsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CorporateActionRecordsResponseItemsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDbModifyTime

`func (o *CorporateActionRecordsResponseItemsInner) GetDbModifyTime() int64`

GetDbModifyTime returns the DbModifyTime field if non-nil, zero value otherwise.

### GetDbModifyTimeOk

`func (o *CorporateActionRecordsResponseItemsInner) GetDbModifyTimeOk() (*int64, bool)`

GetDbModifyTimeOk returns a tuple with the DbModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbModifyTime

`func (o *CorporateActionRecordsResponseItemsInner) SetDbModifyTime(v int64)`

SetDbModifyTime sets DbModifyTime field to given value.

### HasDbModifyTime

`func (o *CorporateActionRecordsResponseItemsInner) HasDbModifyTime() bool`

HasDbModifyTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


