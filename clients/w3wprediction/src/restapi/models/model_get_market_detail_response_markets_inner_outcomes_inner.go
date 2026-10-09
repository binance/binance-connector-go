/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetMarketDetailResponseMarketsInnerOutcomesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMarketDetailResponseMarketsInnerOutcomesInner{}

// GetMarketDetailResponseMarketsInnerOutcomesInner struct for GetMarketDetailResponseMarketsInnerOutcomesInner
type GetMarketDetailResponseMarketsInnerOutcomesInner struct {
	Name    *string `json:"name,omitempty"`
	Price   *string `json:"price,omitempty"`
	Chance  *string `json:"chance,omitempty"`
	Index   *int32  `json:"index,omitempty"`
	TokenId *string `json:"tokenId,omitempty"`
	// Whether this outcome won. Returned only after the market is resolved. More than one outcome can be `true` when the payout is split
	Winner               *bool `json:"winner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMarketDetailResponseMarketsInnerOutcomesInner GetMarketDetailResponseMarketsInnerOutcomesInner

// NewGetMarketDetailResponseMarketsInnerOutcomesInner instantiates a new GetMarketDetailResponseMarketsInnerOutcomesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketDetailResponseMarketsInnerOutcomesInner() *GetMarketDetailResponseMarketsInnerOutcomesInner {
	this := GetMarketDetailResponseMarketsInnerOutcomesInner{}
	return &this
}

// NewGetMarketDetailResponseMarketsInnerOutcomesInnerWithDefaults instantiates a new GetMarketDetailResponseMarketsInnerOutcomesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketDetailResponseMarketsInnerOutcomesInnerWithDefaults() *GetMarketDetailResponseMarketsInnerOutcomesInner {
	this := GetMarketDetailResponseMarketsInnerOutcomesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) SetName(v string) {
	o.Name = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) SetPrice(v string) {
	o.Price = &v
}

// GetChance returns the Chance field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) GetChance() string {
	if o == nil || common.IsNil(o.Chance) {
		var ret string
		return ret
	}
	return *o.Chance
}

// GetChanceOk returns a tuple with the Chance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) GetChanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Chance) {
		return nil, false
	}
	return o.Chance, true
}

// HasChance returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) HasChance() bool {
	if o != nil && !common.IsNil(o.Chance) {
		return true
	}

	return false
}

// SetChance gets a reference to the given string and assigns it to the Chance field.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) SetChance(v string) {
	o.Chance = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) GetIndex() int32 {
	if o == nil || common.IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) GetIndexOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) HasIndex() bool {
	if o != nil && !common.IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) SetIndex(v int32) {
	o.Index = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) SetTokenId(v string) {
	o.TokenId = &v
}

// GetWinner returns the Winner field value if set, zero value otherwise.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) GetWinner() bool {
	if o == nil || common.IsNil(o.Winner) {
		var ret bool
		return ret
	}
	return *o.Winner
}

// GetWinnerOk returns a tuple with the Winner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) GetWinnerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Winner) {
		return nil, false
	}
	return o.Winner, true
}

// HasWinner returns a boolean if a field has been set.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) HasWinner() bool {
	if o != nil && !common.IsNil(o.Winner) {
		return true
	}

	return false
}

// SetWinner gets a reference to the given bool and assigns it to the Winner field.
func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) SetWinner(v bool) {
	o.Winner = &v
}

func (o GetMarketDetailResponseMarketsInnerOutcomesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketDetailResponseMarketsInnerOutcomesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Chance) {
		toSerialize["chance"] = o.Chance
	}
	if !common.IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !common.IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !common.IsNil(o.Winner) {
		toSerialize["winner"] = o.Winner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMarketDetailResponseMarketsInnerOutcomesInner) UnmarshalJSON(data []byte) (err error) {
	varGetMarketDetailResponseMarketsInnerOutcomesInner := _GetMarketDetailResponseMarketsInnerOutcomesInner{}

	err = json.Unmarshal(data, &varGetMarketDetailResponseMarketsInnerOutcomesInner)

	if err != nil {
		return err
	}

	*o = GetMarketDetailResponseMarketsInnerOutcomesInner(varGetMarketDetailResponseMarketsInnerOutcomesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "price")
		delete(additionalProperties, "chance")
		delete(additionalProperties, "index")
		delete(additionalProperties, "tokenId")
		delete(additionalProperties, "winner")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMarketDetailResponseMarketsInnerOutcomesInner struct {
	value *GetMarketDetailResponseMarketsInnerOutcomesInner
	isSet bool
}

func (v NullableGetMarketDetailResponseMarketsInnerOutcomesInner) Get() *GetMarketDetailResponseMarketsInnerOutcomesInner {
	return v.value
}

func (v *NullableGetMarketDetailResponseMarketsInnerOutcomesInner) Set(val *GetMarketDetailResponseMarketsInnerOutcomesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketDetailResponseMarketsInnerOutcomesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketDetailResponseMarketsInnerOutcomesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketDetailResponseMarketsInnerOutcomesInner(val *GetMarketDetailResponseMarketsInnerOutcomesInner) *NullableGetMarketDetailResponseMarketsInnerOutcomesInner {
	return &NullableGetMarketDetailResponseMarketsInnerOutcomesInner{value: val, isSet: true}
}

func (v NullableGetMarketDetailResponseMarketsInnerOutcomesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketDetailResponseMarketsInnerOutcomesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
