/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ListOtcBlocktradesResponseBlocktradesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListOtcBlocktradesResponseBlocktradesInner{}

// ListOtcBlocktradesResponseBlocktradesInner struct for ListOtcBlocktradesResponseBlocktradesInner
type ListOtcBlocktradesResponseBlocktradesInner struct {
	Status   *string `json:"status,omitempty"`
	MarketId *string `json:"marketId,omitempty"`
	TokenId  *string `json:"tokenId,omitempty"`
	// 0 = BUY (Bid), 1 = SELL (Ask) — see Create OTC Blocktrade for the side/quoteType mapping
	Side                 *int32   `json:"side,omitempty"`
	Maker                *string  `json:"maker,omitempty"`
	MakerAmount          *string  `json:"makerAmount,omitempty"`
	TakerAmount          *string  `json:"takerAmount,omitempty"`
	Price                *float64 `json:"price,omitempty"`
	QuoteType            *string  `json:"quoteType,omitempty"`
	CreatedAt            *string  `json:"createdAt,omitempty"`
	SecretToken          *string  `json:"secretToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListOtcBlocktradesResponseBlocktradesInner ListOtcBlocktradesResponseBlocktradesInner

// NewListOtcBlocktradesResponseBlocktradesInner instantiates a new ListOtcBlocktradesResponseBlocktradesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOtcBlocktradesResponseBlocktradesInner() *ListOtcBlocktradesResponseBlocktradesInner {
	this := ListOtcBlocktradesResponseBlocktradesInner{}
	return &this
}

// NewListOtcBlocktradesResponseBlocktradesInnerWithDefaults instantiates a new ListOtcBlocktradesResponseBlocktradesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOtcBlocktradesResponseBlocktradesInnerWithDefaults() *ListOtcBlocktradesResponseBlocktradesInner {
	this := ListOtcBlocktradesResponseBlocktradesInner{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ListOtcBlocktradesResponseBlocktradesInner) SetStatus(v string) {
	o.Status = &v
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetMarketId() string {
	if o == nil || common.IsNil(o.MarketId) {
		var ret string
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetMarketIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketId) {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) HasMarketId() bool {
	if o != nil && !common.IsNil(o.MarketId) {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given string and assigns it to the MarketId field.
func (o *ListOtcBlocktradesResponseBlocktradesInner) SetMarketId(v string) {
	o.MarketId = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *ListOtcBlocktradesResponseBlocktradesInner) SetTokenId(v string) {
	o.TokenId = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetSide() int32 {
	if o == nil || common.IsNil(o.Side) {
		var ret int32
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetSideOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given int32 and assigns it to the Side field.
func (o *ListOtcBlocktradesResponseBlocktradesInner) SetSide(v int32) {
	o.Side = &v
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetMaker() string {
	if o == nil || common.IsNil(o.Maker) {
		var ret string
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetMakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) HasMaker() bool {
	if o != nil && !common.IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given string and assigns it to the Maker field.
func (o *ListOtcBlocktradesResponseBlocktradesInner) SetMaker(v string) {
	o.Maker = &v
}

// GetMakerAmount returns the MakerAmount field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetMakerAmount() string {
	if o == nil || common.IsNil(o.MakerAmount) {
		var ret string
		return ret
	}
	return *o.MakerAmount
}

// GetMakerAmountOk returns a tuple with the MakerAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetMakerAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerAmount) {
		return nil, false
	}
	return o.MakerAmount, true
}

// HasMakerAmount returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) HasMakerAmount() bool {
	if o != nil && !common.IsNil(o.MakerAmount) {
		return true
	}

	return false
}

// SetMakerAmount gets a reference to the given string and assigns it to the MakerAmount field.
func (o *ListOtcBlocktradesResponseBlocktradesInner) SetMakerAmount(v string) {
	o.MakerAmount = &v
}

// GetTakerAmount returns the TakerAmount field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetTakerAmount() string {
	if o == nil || common.IsNil(o.TakerAmount) {
		var ret string
		return ret
	}
	return *o.TakerAmount
}

// GetTakerAmountOk returns a tuple with the TakerAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetTakerAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerAmount) {
		return nil, false
	}
	return o.TakerAmount, true
}

// HasTakerAmount returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) HasTakerAmount() bool {
	if o != nil && !common.IsNil(o.TakerAmount) {
		return true
	}

	return false
}

// SetTakerAmount gets a reference to the given string and assigns it to the TakerAmount field.
func (o *ListOtcBlocktradesResponseBlocktradesInner) SetTakerAmount(v string) {
	o.TakerAmount = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetPrice() float64 {
	if o == nil || common.IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetPriceOk() (*float64, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *ListOtcBlocktradesResponseBlocktradesInner) SetPrice(v float64) {
	o.Price = &v
}

// GetQuoteType returns the QuoteType field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetQuoteType() string {
	if o == nil || common.IsNil(o.QuoteType) {
		var ret string
		return ret
	}
	return *o.QuoteType
}

// GetQuoteTypeOk returns a tuple with the QuoteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetQuoteTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteType) {
		return nil, false
	}
	return o.QuoteType, true
}

// HasQuoteType returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) HasQuoteType() bool {
	if o != nil && !common.IsNil(o.QuoteType) {
		return true
	}

	return false
}

// SetQuoteType gets a reference to the given string and assigns it to the QuoteType field.
func (o *ListOtcBlocktradesResponseBlocktradesInner) SetQuoteType(v string) {
	o.QuoteType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetCreatedAt() string {
	if o == nil || common.IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetCreatedAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) HasCreatedAt() bool {
	if o != nil && !common.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ListOtcBlocktradesResponseBlocktradesInner) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetSecretToken returns the SecretToken field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetSecretToken() string {
	if o == nil || common.IsNil(o.SecretToken) {
		var ret string
		return ret
	}
	return *o.SecretToken
}

// GetSecretTokenOk returns a tuple with the SecretToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) GetSecretTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.SecretToken) {
		return nil, false
	}
	return o.SecretToken, true
}

// HasSecretToken returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponseBlocktradesInner) HasSecretToken() bool {
	if o != nil && !common.IsNil(o.SecretToken) {
		return true
	}

	return false
}

// SetSecretToken gets a reference to the given string and assigns it to the SecretToken field.
func (o *ListOtcBlocktradesResponseBlocktradesInner) SetSecretToken(v string) {
	o.SecretToken = &v
}

func (o ListOtcBlocktradesResponseBlocktradesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOtcBlocktradesResponseBlocktradesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.MarketId) {
		toSerialize["marketId"] = o.MarketId
	}
	if !common.IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Maker) {
		toSerialize["maker"] = o.Maker
	}
	if !common.IsNil(o.MakerAmount) {
		toSerialize["makerAmount"] = o.MakerAmount
	}
	if !common.IsNil(o.TakerAmount) {
		toSerialize["takerAmount"] = o.TakerAmount
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.QuoteType) {
		toSerialize["quoteType"] = o.QuoteType
	}
	if !common.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !common.IsNil(o.SecretToken) {
		toSerialize["secretToken"] = o.SecretToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListOtcBlocktradesResponseBlocktradesInner) UnmarshalJSON(data []byte) (err error) {
	varListOtcBlocktradesResponseBlocktradesInner := _ListOtcBlocktradesResponseBlocktradesInner{}

	err = json.Unmarshal(data, &varListOtcBlocktradesResponseBlocktradesInner)

	if err != nil {
		return err
	}

	*o = ListOtcBlocktradesResponseBlocktradesInner(varListOtcBlocktradesResponseBlocktradesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "marketId")
		delete(additionalProperties, "tokenId")
		delete(additionalProperties, "side")
		delete(additionalProperties, "maker")
		delete(additionalProperties, "makerAmount")
		delete(additionalProperties, "takerAmount")
		delete(additionalProperties, "price")
		delete(additionalProperties, "quoteType")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "secretToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListOtcBlocktradesResponseBlocktradesInner struct {
	value *ListOtcBlocktradesResponseBlocktradesInner
	isSet bool
}

func (v NullableListOtcBlocktradesResponseBlocktradesInner) Get() *ListOtcBlocktradesResponseBlocktradesInner {
	return v.value
}

func (v *NullableListOtcBlocktradesResponseBlocktradesInner) Set(val *ListOtcBlocktradesResponseBlocktradesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListOtcBlocktradesResponseBlocktradesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListOtcBlocktradesResponseBlocktradesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOtcBlocktradesResponseBlocktradesInner(val *ListOtcBlocktradesResponseBlocktradesInner) *NullableListOtcBlocktradesResponseBlocktradesInner {
	return &NullableListOtcBlocktradesResponseBlocktradesInner{value: val, isSet: true}
}

func (v NullableListOtcBlocktradesResponseBlocktradesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOtcBlocktradesResponseBlocktradesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
