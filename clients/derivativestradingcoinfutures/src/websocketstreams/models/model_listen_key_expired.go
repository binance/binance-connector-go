/*
Futures (COIN-M) WebSocket Market Streams

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ListenKeyExpired type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListenKeyExpired{}

// ListenKeyExpired struct for ListenKeyExpired
type ListenKeyExpired struct {
	// Event Time
	E                    *int64  `json:"E,omitempty"`
	ListenKey            *string `json:"listenKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListenKeyExpired ListenKeyExpired

// NewListenKeyExpired instantiates a new ListenKeyExpired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListenKeyExpired() *ListenKeyExpired {
	this := ListenKeyExpired{}
	return &this
}

// NewListenKeyExpiredWithDefaults instantiates a new ListenKeyExpired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListenKeyExpiredWithDefaults() *ListenKeyExpired {
	this := ListenKeyExpired{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *ListenKeyExpired) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenKeyExpired) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *ListenKeyExpired) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *ListenKeyExpired) SetE(v int64) {
	o.E = &v
}

// GetListenKey returns the ListenKey field value if set, zero value otherwise.
func (o *ListenKeyExpired) GetListenKey() string {
	if o == nil || common.IsNil(o.ListenKey) {
		var ret string
		return ret
	}
	return *o.ListenKey
}

// GetListenKeyOk returns a tuple with the ListenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenKeyExpired) GetListenKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListenKey) {
		return nil, false
	}
	return o.ListenKey, true
}

// HasListenKey returns a boolean if a field has been set.
func (o *ListenKeyExpired) HasListenKey() bool {
	if o != nil && !common.IsNil(o.ListenKey) {
		return true
	}

	return false
}

// SetListenKey gets a reference to the given string and assigns it to the ListenKey field.
func (o *ListenKeyExpired) SetListenKey(v string) {
	o.ListenKey = &v
}

func (o ListenKeyExpired) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListenKeyExpired) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.ListenKey) {
		toSerialize["listenKey"] = o.ListenKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListenKeyExpired) UnmarshalJSON(data []byte) (err error) {
	varListenKeyExpired := _ListenKeyExpired{}

	err = json.Unmarshal(data, &varListenKeyExpired)

	if err != nil {
		return err
	}

	*o = ListenKeyExpired(varListenKeyExpired)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "listenKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListenKeyExpired struct {
	value *ListenKeyExpired
	isSet bool
}

func (v NullableListenKeyExpired) Get() *ListenKeyExpired {
	return v.value
}

func (v *NullableListenKeyExpired) Set(val *ListenKeyExpired) {
	v.value = val
	v.isSet = true
}

func (v NullableListenKeyExpired) IsSet() bool {
	return v.isSet
}

func (v *NullableListenKeyExpired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListenKeyExpired(val *ListenKeyExpired) *NullableListenKeyExpired {
	return &NullableListenKeyExpired{value: val, isSet: true}
}

func (v NullableListenKeyExpired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListenKeyExpired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
