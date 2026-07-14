/*
Spot WebSocket Market Streams

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BlockTradeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BlockTradeResponse{}

// BlockTradeResponse struct for BlockTradeResponse
type BlockTradeResponse struct {
	// Event type
	Smalle *string `json:"e,omitempty"`
	// Event time
	E *int64 `json:"E,omitempty"`
	// Symbol
	Smalls *string `json:"s,omitempty"`
	// Block Trade ID
	Smallt *int64 `json:"t,omitempty"`
	// Price
	Smallp *string `json:"p,omitempty"`
	// Quantity
	Smallq *string `json:"q,omitempty"`
	// Trade time
	T *int64 `json:"T,omitempty"`
	// Is the buyer the maker?
	Smallm               *bool `json:"m,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BlockTradeResponse BlockTradeResponse

// NewBlockTradeResponse instantiates a new BlockTradeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockTradeResponse() *BlockTradeResponse {
	this := BlockTradeResponse{}
	return &this
}

// NewBlockTradeResponseWithDefaults instantiates a new BlockTradeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockTradeResponseWithDefaults() *BlockTradeResponse {
	this := BlockTradeResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *BlockTradeResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradeResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *BlockTradeResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *BlockTradeResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *BlockTradeResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradeResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *BlockTradeResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *BlockTradeResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *BlockTradeResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradeResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *BlockTradeResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *BlockTradeResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *BlockTradeResponse) GetSmallt() int64 {
	if o == nil || common.IsNil(o.Smallt) {
		var ret int64
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradeResponse) GetSmalltOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *BlockTradeResponse) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *BlockTradeResponse) SetSmallt(v int64) {
	o.Smallt = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *BlockTradeResponse) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradeResponse) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *BlockTradeResponse) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *BlockTradeResponse) SetSmallp(v string) {
	o.Smallp = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *BlockTradeResponse) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradeResponse) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *BlockTradeResponse) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *BlockTradeResponse) SetSmallq(v string) {
	o.Smallq = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *BlockTradeResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradeResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *BlockTradeResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *BlockTradeResponse) SetT(v int64) {
	o.T = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *BlockTradeResponse) GetSmallm() bool {
	if o == nil || common.IsNil(o.Smallm) {
		var ret bool
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradeResponse) GetSmallmOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *BlockTradeResponse) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *BlockTradeResponse) SetSmallm(v bool) {
	o.Smallm = &v
}

func (o BlockTradeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockTradeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BlockTradeResponse) UnmarshalJSON(data []byte) (err error) {
	varBlockTradeResponse := _BlockTradeResponse{}

	err = json.Unmarshal(data, &varBlockTradeResponse)

	if err != nil {
		return err
	}

	*o = BlockTradeResponse(varBlockTradeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "t")
		delete(additionalProperties, "p")
		delete(additionalProperties, "q")
		delete(additionalProperties, "T")
		delete(additionalProperties, "m")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBlockTradeResponse struct {
	value *BlockTradeResponse
	isSet bool
}

func (v NullableBlockTradeResponse) Get() *BlockTradeResponse {
	return v.value
}

func (v *NullableBlockTradeResponse) Set(val *BlockTradeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockTradeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockTradeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockTradeResponse(val *BlockTradeResponse) *NullableBlockTradeResponse {
	return &NullableBlockTradeResponse{value: val, isSet: true}
}

func (v NullableBlockTradeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockTradeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
