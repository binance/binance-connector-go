/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the BasisResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BasisResponseInner{}

// BasisResponseInner struct for BasisResponseInner
type BasisResponseInner struct {
	IndexPrice           *string `json:"indexPrice,omitempty"`
	ContractType         *string `json:"contractType,omitempty"`
	BasisRate            *string `json:"basisRate,omitempty"`
	FuturesPrice         *string `json:"futuresPrice,omitempty"`
	AnnualizedBasisRate  *string `json:"annualizedBasisRate,omitempty"`
	Basis                *string `json:"basis,omitempty"`
	Pair                 *string `json:"pair,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BasisResponseInner BasisResponseInner

// NewBasisResponseInner instantiates a new BasisResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasisResponseInner() *BasisResponseInner {
	this := BasisResponseInner{}
	return &this
}

// NewBasisResponseInnerWithDefaults instantiates a new BasisResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasisResponseInnerWithDefaults() *BasisResponseInner {
	this := BasisResponseInner{}
	return &this
}

// GetIndexPrice returns the IndexPrice field value if set, zero value otherwise.
func (o *BasisResponseInner) GetIndexPrice() string {
	if o == nil || common.IsNil(o.IndexPrice) {
		var ret string
		return ret
	}
	return *o.IndexPrice
}

// GetIndexPriceOk returns a tuple with the IndexPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasisResponseInner) GetIndexPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.IndexPrice) {
		return nil, false
	}
	return o.IndexPrice, true
}

// HasIndexPrice returns a boolean if a field has been set.
func (o *BasisResponseInner) HasIndexPrice() bool {
	if o != nil && !common.IsNil(o.IndexPrice) {
		return true
	}

	return false
}

// SetIndexPrice gets a reference to the given string and assigns it to the IndexPrice field.
func (o *BasisResponseInner) SetIndexPrice(v string) {
	o.IndexPrice = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *BasisResponseInner) GetContractType() string {
	if o == nil || common.IsNil(o.ContractType) {
		var ret string
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasisResponseInner) GetContractTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *BasisResponseInner) HasContractType() bool {
	if o != nil && !common.IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given string and assigns it to the ContractType field.
func (o *BasisResponseInner) SetContractType(v string) {
	o.ContractType = &v
}

// GetBasisRate returns the BasisRate field value if set, zero value otherwise.
func (o *BasisResponseInner) GetBasisRate() string {
	if o == nil || common.IsNil(o.BasisRate) {
		var ret string
		return ret
	}
	return *o.BasisRate
}

// GetBasisRateOk returns a tuple with the BasisRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasisResponseInner) GetBasisRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.BasisRate) {
		return nil, false
	}
	return o.BasisRate, true
}

// HasBasisRate returns a boolean if a field has been set.
func (o *BasisResponseInner) HasBasisRate() bool {
	if o != nil && !common.IsNil(o.BasisRate) {
		return true
	}

	return false
}

// SetBasisRate gets a reference to the given string and assigns it to the BasisRate field.
func (o *BasisResponseInner) SetBasisRate(v string) {
	o.BasisRate = &v
}

// GetFuturesPrice returns the FuturesPrice field value if set, zero value otherwise.
func (o *BasisResponseInner) GetFuturesPrice() string {
	if o == nil || common.IsNil(o.FuturesPrice) {
		var ret string
		return ret
	}
	return *o.FuturesPrice
}

// GetFuturesPriceOk returns a tuple with the FuturesPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasisResponseInner) GetFuturesPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.FuturesPrice) {
		return nil, false
	}
	return o.FuturesPrice, true
}

// HasFuturesPrice returns a boolean if a field has been set.
func (o *BasisResponseInner) HasFuturesPrice() bool {
	if o != nil && !common.IsNil(o.FuturesPrice) {
		return true
	}

	return false
}

// SetFuturesPrice gets a reference to the given string and assigns it to the FuturesPrice field.
func (o *BasisResponseInner) SetFuturesPrice(v string) {
	o.FuturesPrice = &v
}

// GetAnnualizedBasisRate returns the AnnualizedBasisRate field value if set, zero value otherwise.
func (o *BasisResponseInner) GetAnnualizedBasisRate() string {
	if o == nil || common.IsNil(o.AnnualizedBasisRate) {
		var ret string
		return ret
	}
	return *o.AnnualizedBasisRate
}

// GetAnnualizedBasisRateOk returns a tuple with the AnnualizedBasisRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasisResponseInner) GetAnnualizedBasisRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AnnualizedBasisRate) {
		return nil, false
	}
	return o.AnnualizedBasisRate, true
}

// HasAnnualizedBasisRate returns a boolean if a field has been set.
func (o *BasisResponseInner) HasAnnualizedBasisRate() bool {
	if o != nil && !common.IsNil(o.AnnualizedBasisRate) {
		return true
	}

	return false
}

// SetAnnualizedBasisRate gets a reference to the given string and assigns it to the AnnualizedBasisRate field.
func (o *BasisResponseInner) SetAnnualizedBasisRate(v string) {
	o.AnnualizedBasisRate = &v
}

// GetBasis returns the Basis field value if set, zero value otherwise.
func (o *BasisResponseInner) GetBasis() string {
	if o == nil || common.IsNil(o.Basis) {
		var ret string
		return ret
	}
	return *o.Basis
}

// GetBasisOk returns a tuple with the Basis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasisResponseInner) GetBasisOk() (*string, bool) {
	if o == nil || common.IsNil(o.Basis) {
		return nil, false
	}
	return o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *BasisResponseInner) HasBasis() bool {
	if o != nil && !common.IsNil(o.Basis) {
		return true
	}

	return false
}

// SetBasis gets a reference to the given string and assigns it to the Basis field.
func (o *BasisResponseInner) SetBasis(v string) {
	o.Basis = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *BasisResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasisResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *BasisResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *BasisResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BasisResponseInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasisResponseInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BasisResponseInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *BasisResponseInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o BasisResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasisResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IndexPrice) {
		toSerialize["indexPrice"] = o.IndexPrice
	}
	if !common.IsNil(o.ContractType) {
		toSerialize["contractType"] = o.ContractType
	}
	if !common.IsNil(o.BasisRate) {
		toSerialize["basisRate"] = o.BasisRate
	}
	if !common.IsNil(o.FuturesPrice) {
		toSerialize["futuresPrice"] = o.FuturesPrice
	}
	if !common.IsNil(o.AnnualizedBasisRate) {
		toSerialize["annualizedBasisRate"] = o.AnnualizedBasisRate
	}
	if !common.IsNil(o.Basis) {
		toSerialize["basis"] = o.Basis
	}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BasisResponseInner) UnmarshalJSON(data []byte) (err error) {
	varBasisResponseInner := _BasisResponseInner{}

	err = json.Unmarshal(data, &varBasisResponseInner)

	if err != nil {
		return err
	}

	*o = BasisResponseInner(varBasisResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "indexPrice")
		delete(additionalProperties, "contractType")
		delete(additionalProperties, "basisRate")
		delete(additionalProperties, "futuresPrice")
		delete(additionalProperties, "annualizedBasisRate")
		delete(additionalProperties, "basis")
		delete(additionalProperties, "pair")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBasisResponseInner struct {
	value *BasisResponseInner
	isSet bool
}

func (v NullableBasisResponseInner) Get() *BasisResponseInner {
	return v.value
}

func (v *NullableBasisResponseInner) Set(val *BasisResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBasisResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBasisResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasisResponseInner(val *BasisResponseInner) *NullableBasisResponseInner {
	return &NullableBasisResponseInner{value: val, isSet: true}
}

func (v NullableBasisResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasisResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
