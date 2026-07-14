/*
Spot WebSocket Market Streams

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AvgPriceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AvgPriceResponse{}

// AvgPriceResponse struct for AvgPriceResponse
type AvgPriceResponse struct {
	// Event type
	Smalle *string `json:"e,omitempty"`
	// Event time
	E *int64 `json:"E,omitempty"`
	// Symbol
	Smalls *string `json:"s,omitempty"`
	// Average price interval
	Smalli *string `json:"i,omitempty"`
	// Average price
	Smallw *string `json:"w,omitempty"`
	// Last trade time
	T                    *int64 `json:"T,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AvgPriceResponse AvgPriceResponse

// NewAvgPriceResponse instantiates a new AvgPriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvgPriceResponse() *AvgPriceResponse {
	this := AvgPriceResponse{}
	return &this
}

// NewAvgPriceResponseWithDefaults instantiates a new AvgPriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvgPriceResponseWithDefaults() *AvgPriceResponse {
	this := AvgPriceResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AvgPriceResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AvgPriceResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AvgPriceResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *AvgPriceResponse) SetSmalli(v string) {
	o.Smalli = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetSmallw() string {
	if o == nil || common.IsNil(o.Smallw) {
		var ret string
		return ret
	}
	return *o.Smallw
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetSmallwOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallw) {
		return nil, false
	}
	return o.Smallw, true
}

// HasW returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasSmallw() bool {
	if o != nil && !common.IsNil(o.Smallw) {
		return true
	}

	return false
}

// SetW gets a reference to the given string and assigns it to the W field.
func (o *AvgPriceResponse) SetSmallw(v string) {
	o.Smallw = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AvgPriceResponse) SetT(v int64) {
	o.T = &v
}

func (o AvgPriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvgPriceResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.Smallw) {
		toSerialize["w"] = o.Smallw
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AvgPriceResponse) UnmarshalJSON(data []byte) (err error) {
	varAvgPriceResponse := _AvgPriceResponse{}

	err = json.Unmarshal(data, &varAvgPriceResponse)

	if err != nil {
		return err
	}

	*o = AvgPriceResponse(varAvgPriceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "i")
		delete(additionalProperties, "w")
		delete(additionalProperties, "T")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAvgPriceResponse struct {
	value *AvgPriceResponse
	isSet bool
}

func (v NullableAvgPriceResponse) Get() *AvgPriceResponse {
	return v.value
}

func (v *NullableAvgPriceResponse) Set(val *AvgPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAvgPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAvgPriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvgPriceResponse(val *AvgPriceResponse) *NullableAvgPriceResponse {
	return &NullableAvgPriceResponse{value: val, isSet: true}
}

func (v NullableAvgPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvgPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
