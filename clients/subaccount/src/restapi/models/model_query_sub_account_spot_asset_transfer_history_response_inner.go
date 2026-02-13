/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySubAccountSpotAssetTransferHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountSpotAssetTransferHistoryResponseInner{}

// QuerySubAccountSpotAssetTransferHistoryResponseInner struct for QuerySubAccountSpotAssetTransferHistoryResponseInner
type QuerySubAccountSpotAssetTransferHistoryResponseInner struct {
	From                 *string `json:"from,omitempty"`
	To                   *string `json:"to,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	Status               *string `json:"status,omitempty"`
	TranId               *int64  `json:"tranId,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubAccountSpotAssetTransferHistoryResponseInner QuerySubAccountSpotAssetTransferHistoryResponseInner

// NewQuerySubAccountSpotAssetTransferHistoryResponseInner instantiates a new QuerySubAccountSpotAssetTransferHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountSpotAssetTransferHistoryResponseInner() *QuerySubAccountSpotAssetTransferHistoryResponseInner {
	this := QuerySubAccountSpotAssetTransferHistoryResponseInner{}
	return &this
}

// NewQuerySubAccountSpotAssetTransferHistoryResponseInnerWithDefaults instantiates a new QuerySubAccountSpotAssetTransferHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountSpotAssetTransferHistoryResponseInnerWithDefaults() *QuerySubAccountSpotAssetTransferHistoryResponseInner {
	this := QuerySubAccountSpotAssetTransferHistoryResponseInner{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetFrom() string {
	if o == nil || common.IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetFromOk() (*string, bool) {
	if o == nil || common.IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) HasFrom() bool {
	if o != nil && !common.IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetTo() string {
	if o == nil || common.IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetToOk() (*string, bool) {
	if o == nil || common.IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) HasTo() bool {
	if o != nil && !common.IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) SetTo(v string) {
	o.To = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) SetQty(v string) {
	o.Qty = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) SetTranId(v int64) {
	o.TranId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) SetTime(v int64) {
	o.Time = &v
}

func (o QuerySubAccountSpotAssetTransferHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountSpotAssetTransferHistoryResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !common.IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubAccountSpotAssetTransferHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountSpotAssetTransferHistoryResponseInner := _QuerySubAccountSpotAssetTransferHistoryResponseInner{}

	err = json.Unmarshal(data, &varQuerySubAccountSpotAssetTransferHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = QuerySubAccountSpotAssetTransferHistoryResponseInner(varQuerySubAccountSpotAssetTransferHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tranId")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountSpotAssetTransferHistoryResponseInner struct {
	value *QuerySubAccountSpotAssetTransferHistoryResponseInner
	isSet bool
}

func (v NullableQuerySubAccountSpotAssetTransferHistoryResponseInner) Get() *QuerySubAccountSpotAssetTransferHistoryResponseInner {
	return v.value
}

func (v *NullableQuerySubAccountSpotAssetTransferHistoryResponseInner) Set(val *QuerySubAccountSpotAssetTransferHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountSpotAssetTransferHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountSpotAssetTransferHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountSpotAssetTransferHistoryResponseInner(val *QuerySubAccountSpotAssetTransferHistoryResponseInner) *NullableQuerySubAccountSpotAssetTransferHistoryResponseInner {
	return &NullableQuerySubAccountSpotAssetTransferHistoryResponseInner{value: val, isSet: true}
}

func (v NullableQuerySubAccountSpotAssetTransferHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountSpotAssetTransferHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
