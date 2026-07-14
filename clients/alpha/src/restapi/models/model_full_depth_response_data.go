/*
Alpha Trading REST API

APIs for Binance Alpha Trading.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FullDepthResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FullDepthResponseData{}

// FullDepthResponseData Order book data.
type FullDepthResponseData struct {
	// Last order book update ID.
	LastUpdateId *int64 `json:"lastUpdateId,omitempty"`
	// Trading pair symbol.
	Symbol *string `json:"symbol,omitempty"`
	// Bid orders. Each entry is [price, quantity].
	Bids [][]string `json:"bids,omitempty"`
	// Ask orders. Each entry is [price, quantity].
	Asks [][]string `json:"asks,omitempty"`
	// Event time in milliseconds.
	E *int64 `json:"E,omitempty"`
	// Transaction time in milliseconds.
	T                    *int64 `json:"T,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FullDepthResponseData FullDepthResponseData

// NewFullDepthResponseData instantiates a new FullDepthResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullDepthResponseData() *FullDepthResponseData {
	this := FullDepthResponseData{}
	return &this
}

// NewFullDepthResponseDataWithDefaults instantiates a new FullDepthResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullDepthResponseDataWithDefaults() *FullDepthResponseData {
	this := FullDepthResponseData{}
	return &this
}

// GetLastUpdateId returns the LastUpdateId field value if set, zero value otherwise.
func (o *FullDepthResponseData) GetLastUpdateId() int64 {
	if o == nil || common.IsNil(o.LastUpdateId) {
		var ret int64
		return ret
	}
	return *o.LastUpdateId
}

// GetLastUpdateIdOk returns a tuple with the LastUpdateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthResponseData) GetLastUpdateIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastUpdateId) {
		return nil, false
	}
	return o.LastUpdateId, true
}

// HasLastUpdateId returns a boolean if a field has been set.
func (o *FullDepthResponseData) HasLastUpdateId() bool {
	if o != nil && !common.IsNil(o.LastUpdateId) {
		return true
	}

	return false
}

// SetLastUpdateId gets a reference to the given int64 and assigns it to the LastUpdateId field.
func (o *FullDepthResponseData) SetLastUpdateId(v int64) {
	o.LastUpdateId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *FullDepthResponseData) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthResponseData) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *FullDepthResponseData) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *FullDepthResponseData) SetSymbol(v string) {
	o.Symbol = &v
}

// GetBids returns the Bids field value if set, zero value otherwise.
func (o *FullDepthResponseData) GetBids() [][]string {
	if o == nil || common.IsNil(o.Bids) {
		var ret [][]string
		return ret
	}
	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthResponseData) GetBidsOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Bids) {
		return nil, false
	}
	return o.Bids, true
}

// HasBids returns a boolean if a field has been set.
func (o *FullDepthResponseData) HasBids() bool {
	if o != nil && !common.IsNil(o.Bids) {
		return true
	}

	return false
}

// SetBids gets a reference to the given [][]string and assigns it to the Bids field.
func (o *FullDepthResponseData) SetBids(v [][]string) {
	o.Bids = v
}

// GetAsks returns the Asks field value if set, zero value otherwise.
func (o *FullDepthResponseData) GetAsks() [][]string {
	if o == nil || common.IsNil(o.Asks) {
		var ret [][]string
		return ret
	}
	return o.Asks
}

// GetAsksOk returns a tuple with the Asks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthResponseData) GetAsksOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Asks) {
		return nil, false
	}
	return o.Asks, true
}

// HasAsks returns a boolean if a field has been set.
func (o *FullDepthResponseData) HasAsks() bool {
	if o != nil && !common.IsNil(o.Asks) {
		return true
	}

	return false
}

// SetAsks gets a reference to the given [][]string and assigns it to the Asks field.
func (o *FullDepthResponseData) SetAsks(v [][]string) {
	o.Asks = v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *FullDepthResponseData) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthResponseData) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *FullDepthResponseData) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *FullDepthResponseData) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *FullDepthResponseData) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthResponseData) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *FullDepthResponseData) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *FullDepthResponseData) SetT(v int64) {
	o.T = &v
}

func (o FullDepthResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FullDepthResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LastUpdateId) {
		toSerialize["lastUpdateId"] = o.LastUpdateId
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Bids) {
		toSerialize["bids"] = o.Bids
	}
	if !common.IsNil(o.Asks) {
		toSerialize["asks"] = o.Asks
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FullDepthResponseData) UnmarshalJSON(data []byte) (err error) {
	varFullDepthResponseData := _FullDepthResponseData{}

	err = json.Unmarshal(data, &varFullDepthResponseData)

	if err != nil {
		return err
	}

	*o = FullDepthResponseData(varFullDepthResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lastUpdateId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "bids")
		delete(additionalProperties, "asks")
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFullDepthResponseData struct {
	value *FullDepthResponseData
	isSet bool
}

func (v NullableFullDepthResponseData) Get() *FullDepthResponseData {
	return v.value
}

func (v *NullableFullDepthResponseData) Set(val *FullDepthResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableFullDepthResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableFullDepthResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullDepthResponseData(val *FullDepthResponseData) *NullableFullDepthResponseData {
	return &NullableFullDepthResponseData{value: val, isSet: true}
}

func (v NullableFullDepthResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullDepthResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
