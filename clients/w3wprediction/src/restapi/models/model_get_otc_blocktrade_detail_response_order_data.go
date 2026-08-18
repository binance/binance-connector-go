/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOtcBlocktradeDetailResponseOrderData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOtcBlocktradeDetailResponseOrderData{}

// GetOtcBlocktradeDetailResponseOrderData struct for GetOtcBlocktradeDetailResponseOrderData
type GetOtcBlocktradeDetailResponseOrderData struct {
	Status   *string `json:"status,omitempty"`
	MarketId *string `json:"marketId,omitempty"`
	TokenId  *string `json:"tokenId,omitempty"`
	// 0 = BUY (Bid), 1 = SELL (Ask) — see Create OTC Blocktrade for the side/quoteType mapping
	Side                 *int32   `json:"side,omitempty"`
	Maker                *string  `json:"maker,omitempty"`
	Taker                *string  `json:"taker,omitempty"`
	MakerAmount          *string  `json:"makerAmount,omitempty"`
	TakerAmount          *string  `json:"takerAmount,omitempty"`
	Price                *float32 `json:"price,omitempty"`
	OrderType            *string  `json:"orderType,omitempty"`
	TimeInForce          *string  `json:"timeInForce,omitempty"`
	Expiration           *int64   `json:"expiration,omitempty"`
	FilledAmount         *string  `json:"filledAmount,omitempty"`
	QuoteType            *string  `json:"quoteType,omitempty"`
	CreatedAt            *string  `json:"createdAt,omitempty"`
	IsNegRisk            *bool    `json:"isNegRisk,omitempty"`
	IsYieldBearing       *bool    `json:"isYieldBearing,omitempty"`
	SecretToken          *string  `json:"secretToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOtcBlocktradeDetailResponseOrderData GetOtcBlocktradeDetailResponseOrderData

// NewGetOtcBlocktradeDetailResponseOrderData instantiates a new GetOtcBlocktradeDetailResponseOrderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOtcBlocktradeDetailResponseOrderData() *GetOtcBlocktradeDetailResponseOrderData {
	this := GetOtcBlocktradeDetailResponseOrderData{}
	return &this
}

// NewGetOtcBlocktradeDetailResponseOrderDataWithDefaults instantiates a new GetOtcBlocktradeDetailResponseOrderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOtcBlocktradeDetailResponseOrderDataWithDefaults() *GetOtcBlocktradeDetailResponseOrderData {
	this := GetOtcBlocktradeDetailResponseOrderData{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetStatus(v string) {
	o.Status = &v
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetMarketId() string {
	if o == nil || common.IsNil(o.MarketId) {
		var ret string
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetMarketIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketId) {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasMarketId() bool {
	if o != nil && !common.IsNil(o.MarketId) {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given string and assigns it to the MarketId field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetMarketId(v string) {
	o.MarketId = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetTokenId(v string) {
	o.TokenId = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetSide() int32 {
	if o == nil || common.IsNil(o.Side) {
		var ret int32
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetSideOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given int32 and assigns it to the Side field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetSide(v int32) {
	o.Side = &v
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetMaker() string {
	if o == nil || common.IsNil(o.Maker) {
		var ret string
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetMakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasMaker() bool {
	if o != nil && !common.IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given string and assigns it to the Maker field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetMaker(v string) {
	o.Maker = &v
}

// GetTaker returns the Taker field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetOtcBlocktradeDetailResponseOrderData) GetTaker() string {
	if o == nil || common.IsNil(o.Taker) {
		var ret string
		return ret
	}
	return *o.Taker
}

// GetTakerOk returns a tuple with the Taker field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOtcBlocktradeDetailResponseOrderData) GetTakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Taker) {
		return nil, false
	}
	return o.Taker, true
}

// HasTaker returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasTaker() bool {
	if o != nil && !common.IsNil(o.Taker) {
		return true
	}

	return false
}

// SetTaker gets a reference to the given NullableString and assigns it to the Taker field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetTaker(v string) {
	o.Taker = &v
}

// GetMakerAmount returns the MakerAmount field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetMakerAmount() string {
	if o == nil || common.IsNil(o.MakerAmount) {
		var ret string
		return ret
	}
	return *o.MakerAmount
}

// GetMakerAmountOk returns a tuple with the MakerAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetMakerAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerAmount) {
		return nil, false
	}
	return o.MakerAmount, true
}

// HasMakerAmount returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasMakerAmount() bool {
	if o != nil && !common.IsNil(o.MakerAmount) {
		return true
	}

	return false
}

// SetMakerAmount gets a reference to the given string and assigns it to the MakerAmount field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetMakerAmount(v string) {
	o.MakerAmount = &v
}

// GetTakerAmount returns the TakerAmount field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetTakerAmount() string {
	if o == nil || common.IsNil(o.TakerAmount) {
		var ret string
		return ret
	}
	return *o.TakerAmount
}

// GetTakerAmountOk returns a tuple with the TakerAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetTakerAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerAmount) {
		return nil, false
	}
	return o.TakerAmount, true
}

// HasTakerAmount returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasTakerAmount() bool {
	if o != nil && !common.IsNil(o.TakerAmount) {
		return true
	}

	return false
}

// SetTakerAmount gets a reference to the given string and assigns it to the TakerAmount field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetTakerAmount(v string) {
	o.TakerAmount = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetPrice() float32 {
	if o == nil || common.IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetPriceOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetPrice(v float32) {
	o.Price = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetOrderType() string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetOrderType(v string) {
	o.OrderType = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetExpiration() int64 {
	if o == nil || common.IsNil(o.Expiration) {
		var ret int64
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetExpirationOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasExpiration() bool {
	if o != nil && !common.IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given int64 and assigns it to the Expiration field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetExpiration(v int64) {
	o.Expiration = &v
}

// GetFilledAmount returns the FilledAmount field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetFilledAmount() string {
	if o == nil || common.IsNil(o.FilledAmount) {
		var ret string
		return ret
	}
	return *o.FilledAmount
}

// GetFilledAmountOk returns a tuple with the FilledAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetFilledAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilledAmount) {
		return nil, false
	}
	return o.FilledAmount, true
}

// HasFilledAmount returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasFilledAmount() bool {
	if o != nil && !common.IsNil(o.FilledAmount) {
		return true
	}

	return false
}

// SetFilledAmount gets a reference to the given string and assigns it to the FilledAmount field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetFilledAmount(v string) {
	o.FilledAmount = &v
}

// GetQuoteType returns the QuoteType field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetQuoteType() string {
	if o == nil || common.IsNil(o.QuoteType) {
		var ret string
		return ret
	}
	return *o.QuoteType
}

// GetQuoteTypeOk returns a tuple with the QuoteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetQuoteTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteType) {
		return nil, false
	}
	return o.QuoteType, true
}

// HasQuoteType returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasQuoteType() bool {
	if o != nil && !common.IsNil(o.QuoteType) {
		return true
	}

	return false
}

// SetQuoteType gets a reference to the given string and assigns it to the QuoteType field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetQuoteType(v string) {
	o.QuoteType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetCreatedAt() string {
	if o == nil || common.IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetCreatedAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasCreatedAt() bool {
	if o != nil && !common.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetIsNegRisk returns the IsNegRisk field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetIsNegRisk() bool {
	if o == nil || common.IsNil(o.IsNegRisk) {
		var ret bool
		return ret
	}
	return *o.IsNegRisk
}

// GetIsNegRiskOk returns a tuple with the IsNegRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetIsNegRiskOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsNegRisk) {
		return nil, false
	}
	return o.IsNegRisk, true
}

// HasIsNegRisk returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasIsNegRisk() bool {
	if o != nil && !common.IsNil(o.IsNegRisk) {
		return true
	}

	return false
}

// SetIsNegRisk gets a reference to the given bool and assigns it to the IsNegRisk field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetIsNegRisk(v bool) {
	o.IsNegRisk = &v
}

// GetIsYieldBearing returns the IsYieldBearing field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetIsYieldBearing() bool {
	if o == nil || common.IsNil(o.IsYieldBearing) {
		var ret bool
		return ret
	}
	return *o.IsYieldBearing
}

// GetIsYieldBearingOk returns a tuple with the IsYieldBearing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetIsYieldBearingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsYieldBearing) {
		return nil, false
	}
	return o.IsYieldBearing, true
}

// HasIsYieldBearing returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasIsYieldBearing() bool {
	if o != nil && !common.IsNil(o.IsYieldBearing) {
		return true
	}

	return false
}

// SetIsYieldBearing gets a reference to the given bool and assigns it to the IsYieldBearing field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetIsYieldBearing(v bool) {
	o.IsYieldBearing = &v
}

// GetSecretToken returns the SecretToken field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetSecretToken() string {
	if o == nil || common.IsNil(o.SecretToken) {
		var ret string
		return ret
	}
	return *o.SecretToken
}

// GetSecretTokenOk returns a tuple with the SecretToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) GetSecretTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.SecretToken) {
		return nil, false
	}
	return o.SecretToken, true
}

// HasSecretToken returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponseOrderData) HasSecretToken() bool {
	if o != nil && !common.IsNil(o.SecretToken) {
		return true
	}

	return false
}

// SetSecretToken gets a reference to the given string and assigns it to the SecretToken field.
func (o *GetOtcBlocktradeDetailResponseOrderData) SetSecretToken(v string) {
	o.SecretToken = &v
}

func (o GetOtcBlocktradeDetailResponseOrderData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOtcBlocktradeDetailResponseOrderData) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Taker) {
		toSerialize["taker"] = o.Taker
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
	if !common.IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !common.IsNil(o.FilledAmount) {
		toSerialize["filledAmount"] = o.FilledAmount
	}
	if !common.IsNil(o.QuoteType) {
		toSerialize["quoteType"] = o.QuoteType
	}
	if !common.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !common.IsNil(o.IsNegRisk) {
		toSerialize["isNegRisk"] = o.IsNegRisk
	}
	if !common.IsNil(o.IsYieldBearing) {
		toSerialize["isYieldBearing"] = o.IsYieldBearing
	}
	if !common.IsNil(o.SecretToken) {
		toSerialize["secretToken"] = o.SecretToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOtcBlocktradeDetailResponseOrderData) UnmarshalJSON(data []byte) (err error) {
	varGetOtcBlocktradeDetailResponseOrderData := _GetOtcBlocktradeDetailResponseOrderData{}

	err = json.Unmarshal(data, &varGetOtcBlocktradeDetailResponseOrderData)

	if err != nil {
		return err
	}

	*o = GetOtcBlocktradeDetailResponseOrderData(varGetOtcBlocktradeDetailResponseOrderData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "marketId")
		delete(additionalProperties, "tokenId")
		delete(additionalProperties, "side")
		delete(additionalProperties, "maker")
		delete(additionalProperties, "taker")
		delete(additionalProperties, "makerAmount")
		delete(additionalProperties, "takerAmount")
		delete(additionalProperties, "price")
		delete(additionalProperties, "orderType")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "filledAmount")
		delete(additionalProperties, "quoteType")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "isNegRisk")
		delete(additionalProperties, "isYieldBearing")
		delete(additionalProperties, "secretToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOtcBlocktradeDetailResponseOrderData struct {
	value *GetOtcBlocktradeDetailResponseOrderData
	isSet bool
}

func (v NullableGetOtcBlocktradeDetailResponseOrderData) Get() *GetOtcBlocktradeDetailResponseOrderData {
	return v.value
}

func (v *NullableGetOtcBlocktradeDetailResponseOrderData) Set(val *GetOtcBlocktradeDetailResponseOrderData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOtcBlocktradeDetailResponseOrderData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOtcBlocktradeDetailResponseOrderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOtcBlocktradeDetailResponseOrderData(val *GetOtcBlocktradeDetailResponseOrderData) *NullableGetOtcBlocktradeDetailResponseOrderData {
	return &NullableGetOtcBlocktradeDetailResponseOrderData{value: val, isSet: true}
}

func (v NullableGetOtcBlocktradeDetailResponseOrderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOtcBlocktradeDetailResponseOrderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
