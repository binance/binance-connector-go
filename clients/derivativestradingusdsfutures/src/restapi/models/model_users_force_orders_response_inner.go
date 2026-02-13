/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the UsersForceOrdersResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UsersForceOrdersResponseInner{}

// UsersForceOrdersResponseInner struct for UsersForceOrdersResponseInner
type UsersForceOrdersResponseInner struct {
	OrderId              *int64  `json:"orderId,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Status               *string `json:"status,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	Price                *string `json:"price,omitempty"`
	AvgPrice             *string `json:"avgPrice,omitempty"`
	OrigQty              *string `json:"origQty,omitempty"`
	ExecutedQty          *string `json:"executedQty,omitempty"`
	CumQuote             *string `json:"cumQuote,omitempty"`
	TimeInForce          *string `json:"timeInForce,omitempty"`
	Type                 *string `json:"type,omitempty"`
	ReduceOnly           *bool   `json:"reduceOnly,omitempty"`
	ClosePosition        *bool   `json:"closePosition,omitempty"`
	Side                 *string `json:"side,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	StopPrice            *string `json:"stopPrice,omitempty"`
	WorkingType          *string `json:"workingType,omitempty"`
	OrigType             *string `json:"origType,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UsersForceOrdersResponseInner UsersForceOrdersResponseInner

// NewUsersForceOrdersResponseInner instantiates a new UsersForceOrdersResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersForceOrdersResponseInner() *UsersForceOrdersResponseInner {
	this := UsersForceOrdersResponseInner{}
	return &this
}

// NewUsersForceOrdersResponseInnerWithDefaults instantiates a new UsersForceOrdersResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersForceOrdersResponseInnerWithDefaults() *UsersForceOrdersResponseInner {
	this := UsersForceOrdersResponseInner{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *UsersForceOrdersResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UsersForceOrdersResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UsersForceOrdersResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *UsersForceOrdersResponseInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *UsersForceOrdersResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetAvgPrice() string {
	if o == nil || common.IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasAvgPrice() bool {
	if o != nil && !common.IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *UsersForceOrdersResponseInner) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *UsersForceOrdersResponseInner) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *UsersForceOrdersResponseInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetCumQuote returns the CumQuote field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetCumQuote() string {
	if o == nil || common.IsNil(o.CumQuote) {
		var ret string
		return ret
	}
	return *o.CumQuote
}

// GetCumQuoteOk returns a tuple with the CumQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetCumQuoteOk() (*string, bool) {
	if o == nil || common.IsNil(o.CumQuote) {
		return nil, false
	}
	return o.CumQuote, true
}

// HasCumQuote returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasCumQuote() bool {
	if o != nil && !common.IsNil(o.CumQuote) {
		return true
	}

	return false
}

// SetCumQuote gets a reference to the given string and assigns it to the CumQuote field.
func (o *UsersForceOrdersResponseInner) SetCumQuote(v string) {
	o.CumQuote = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *UsersForceOrdersResponseInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UsersForceOrdersResponseInner) SetType(v string) {
	o.Type = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *UsersForceOrdersResponseInner) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetClosePosition returns the ClosePosition field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetClosePosition() bool {
	if o == nil || common.IsNil(o.ClosePosition) {
		var ret bool
		return ret
	}
	return *o.ClosePosition
}

// GetClosePositionOk returns a tuple with the ClosePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetClosePositionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ClosePosition) {
		return nil, false
	}
	return o.ClosePosition, true
}

// HasClosePosition returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasClosePosition() bool {
	if o != nil && !common.IsNil(o.ClosePosition) {
		return true
	}

	return false
}

// SetClosePosition gets a reference to the given bool and assigns it to the ClosePosition field.
func (o *UsersForceOrdersResponseInner) SetClosePosition(v bool) {
	o.ClosePosition = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *UsersForceOrdersResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *UsersForceOrdersResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetStopPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *UsersForceOrdersResponseInner) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetWorkingType() string {
	if o == nil || common.IsNil(o.WorkingType) {
		var ret string
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetWorkingTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasWorkingType() bool {
	if o != nil && !common.IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given string and assigns it to the WorkingType field.
func (o *UsersForceOrdersResponseInner) SetWorkingType(v string) {
	o.WorkingType = &v
}

// GetOrigType returns the OrigType field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetOrigType() string {
	if o == nil || common.IsNil(o.OrigType) {
		var ret string
		return ret
	}
	return *o.OrigType
}

// GetOrigTypeOk returns a tuple with the OrigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetOrigTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigType) {
		return nil, false
	}
	return o.OrigType, true
}

// HasOrigType returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasOrigType() bool {
	if o != nil && !common.IsNil(o.OrigType) {
		return true
	}

	return false
}

// SetOrigType gets a reference to the given string and assigns it to the OrigType field.
func (o *UsersForceOrdersResponseInner) SetOrigType(v string) {
	o.OrigType = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *UsersForceOrdersResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *UsersForceOrdersResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersForceOrdersResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *UsersForceOrdersResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *UsersForceOrdersResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o UsersForceOrdersResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersForceOrdersResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.CumQuote) {
		toSerialize["cumQuote"] = o.CumQuote
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}
	if !common.IsNil(o.ClosePosition) {
		toSerialize["closePosition"] = o.ClosePosition
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !common.IsNil(o.WorkingType) {
		toSerialize["workingType"] = o.WorkingType
	}
	if !common.IsNil(o.OrigType) {
		toSerialize["origType"] = o.OrigType
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsersForceOrdersResponseInner) UnmarshalJSON(data []byte) (err error) {
	varUsersForceOrdersResponseInner := _UsersForceOrdersResponseInner{}

	err = json.Unmarshal(data, &varUsersForceOrdersResponseInner)

	if err != nil {
		return err
	}

	*o = UsersForceOrdersResponseInner(varUsersForceOrdersResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "status")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "price")
		delete(additionalProperties, "avgPrice")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "cumQuote")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "type")
		delete(additionalProperties, "reduceOnly")
		delete(additionalProperties, "closePosition")
		delete(additionalProperties, "side")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "stopPrice")
		delete(additionalProperties, "workingType")
		delete(additionalProperties, "origType")
		delete(additionalProperties, "time")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsersForceOrdersResponseInner struct {
	value *UsersForceOrdersResponseInner
	isSet bool
}

func (v NullableUsersForceOrdersResponseInner) Get() *UsersForceOrdersResponseInner {
	return v.value
}

func (v *NullableUsersForceOrdersResponseInner) Set(val *UsersForceOrdersResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersForceOrdersResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersForceOrdersResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersForceOrdersResponseInner(val *UsersForceOrdersResponseInner) *NullableUsersForceOrdersResponseInner {
	return &NullableUsersForceOrdersResponseInner{value: val, isSet: true}
}

func (v NullableUsersForceOrdersResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersForceOrdersResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
