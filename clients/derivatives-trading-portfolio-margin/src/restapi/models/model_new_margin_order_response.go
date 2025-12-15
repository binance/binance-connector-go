/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the NewMarginOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NewMarginOrderResponse{}

// NewMarginOrderResponse struct for NewMarginOrderResponse
type NewMarginOrderResponse struct {
	Symbol                *string                            `json:"symbol,omitempty"`
	OrderId               *int64                             `json:"orderId,omitempty"`
	ClientOrderId         *string                            `json:"clientOrderId,omitempty"`
	TransactTime          *int64                             `json:"transactTime,omitempty"`
	Price                 *string                            `json:"price,omitempty"`
	OrigQty               *string                            `json:"origQty,omitempty"`
	ExecutedQty           *string                            `json:"executedQty,omitempty"`
	CummulativeQuoteQty   *string                            `json:"cummulativeQuoteQty,omitempty"`
	Status                *string                            `json:"status,omitempty"`
	TimeInForce           *string                            `json:"timeInForce,omitempty"`
	Type                  *string                            `json:"type,omitempty"`
	Side                  *string                            `json:"side,omitempty"`
	MarginBuyBorrowAmount *string                            `json:"marginBuyBorrowAmount,omitempty"`
	MarginBuyBorrowAsset  *string                            `json:"marginBuyBorrowAsset,omitempty"`
	Fills                 []NewMarginOrderResponseFillsInner `json:"fills,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _NewMarginOrderResponse NewMarginOrderResponse

// NewNewMarginOrderResponse instantiates a new NewMarginOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewMarginOrderResponse() *NewMarginOrderResponse {
	this := NewMarginOrderResponse{}
	return &this
}

// NewNewMarginOrderResponseWithDefaults instantiates a new NewMarginOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewMarginOrderResponseWithDefaults() *NewMarginOrderResponse {
	this := NewMarginOrderResponse{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *NewMarginOrderResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *NewMarginOrderResponse) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *NewMarginOrderResponse) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetTransactTime returns the TransactTime field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetTransactTime() int64 {
	if o == nil || common.IsNil(o.TransactTime) {
		var ret int64
		return ret
	}
	return *o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetTransactTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactTime) {
		return nil, false
	}
	return o.TransactTime, true
}

// HasTransactTime returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasTransactTime() bool {
	if o != nil && !common.IsNil(o.TransactTime) {
		return true
	}

	return false
}

// SetTransactTime gets a reference to the given int64 and assigns it to the TransactTime field.
func (o *NewMarginOrderResponse) SetTransactTime(v int64) {
	o.TransactTime = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *NewMarginOrderResponse) SetPrice(v string) {
	o.Price = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *NewMarginOrderResponse) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *NewMarginOrderResponse) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetCummulativeQuoteQty() string {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasCummulativeQuoteQty() bool {
	if o != nil && !common.IsNil(o.CummulativeQuoteQty) {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given string and assigns it to the CummulativeQuoteQty field.
func (o *NewMarginOrderResponse) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NewMarginOrderResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *NewMarginOrderResponse) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NewMarginOrderResponse) SetType(v string) {
	o.Type = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *NewMarginOrderResponse) SetSide(v string) {
	o.Side = &v
}

// GetMarginBuyBorrowAmount returns the MarginBuyBorrowAmount field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetMarginBuyBorrowAmount() string {
	if o == nil || common.IsNil(o.MarginBuyBorrowAmount) {
		var ret string
		return ret
	}
	return *o.MarginBuyBorrowAmount
}

// GetMarginBuyBorrowAmountOk returns a tuple with the MarginBuyBorrowAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetMarginBuyBorrowAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginBuyBorrowAmount) {
		return nil, false
	}
	return o.MarginBuyBorrowAmount, true
}

// HasMarginBuyBorrowAmount returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasMarginBuyBorrowAmount() bool {
	if o != nil && !common.IsNil(o.MarginBuyBorrowAmount) {
		return true
	}

	return false
}

// SetMarginBuyBorrowAmount gets a reference to the given string and assigns it to the MarginBuyBorrowAmount field.
func (o *NewMarginOrderResponse) SetMarginBuyBorrowAmount(v string) {
	o.MarginBuyBorrowAmount = &v
}

// GetMarginBuyBorrowAsset returns the MarginBuyBorrowAsset field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetMarginBuyBorrowAsset() string {
	if o == nil || common.IsNil(o.MarginBuyBorrowAsset) {
		var ret string
		return ret
	}
	return *o.MarginBuyBorrowAsset
}

// GetMarginBuyBorrowAssetOk returns a tuple with the MarginBuyBorrowAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetMarginBuyBorrowAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginBuyBorrowAsset) {
		return nil, false
	}
	return o.MarginBuyBorrowAsset, true
}

// HasMarginBuyBorrowAsset returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasMarginBuyBorrowAsset() bool {
	if o != nil && !common.IsNil(o.MarginBuyBorrowAsset) {
		return true
	}

	return false
}

// SetMarginBuyBorrowAsset gets a reference to the given string and assigns it to the MarginBuyBorrowAsset field.
func (o *NewMarginOrderResponse) SetMarginBuyBorrowAsset(v string) {
	o.MarginBuyBorrowAsset = &v
}

// GetFills returns the Fills field value if set, zero value otherwise.
func (o *NewMarginOrderResponse) GetFills() []NewMarginOrderResponseFillsInner {
	if o == nil || common.IsNil(o.Fills) {
		var ret []NewMarginOrderResponseFillsInner
		return ret
	}
	return o.Fills
}

// GetFillsOk returns a tuple with the Fills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponse) GetFillsOk() ([]NewMarginOrderResponseFillsInner, bool) {
	if o == nil || common.IsNil(o.Fills) {
		return nil, false
	}
	return o.Fills, true
}

// HasFills returns a boolean if a field has been set.
func (o *NewMarginOrderResponse) HasFills() bool {
	if o != nil && !common.IsNil(o.Fills) {
		return true
	}

	return false
}

// SetFills gets a reference to the given []NewMarginOrderResponseFillsInner and assigns it to the Fills field.
func (o *NewMarginOrderResponse) SetFills(v []NewMarginOrderResponseFillsInner) {
	o.Fills = v
}

func (o NewMarginOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewMarginOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.TransactTime) {
		toSerialize["transactTime"] = o.TransactTime
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.CummulativeQuoteQty) {
		toSerialize["cummulativeQuoteQty"] = o.CummulativeQuoteQty
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.MarginBuyBorrowAmount) {
		toSerialize["marginBuyBorrowAmount"] = o.MarginBuyBorrowAmount
	}
	if !common.IsNil(o.MarginBuyBorrowAsset) {
		toSerialize["marginBuyBorrowAsset"] = o.MarginBuyBorrowAsset
	}
	if !common.IsNil(o.Fills) {
		toSerialize["fills"] = o.Fills
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NewMarginOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varNewMarginOrderResponse := _NewMarginOrderResponse{}

	err = json.Unmarshal(data, &varNewMarginOrderResponse)

	if err != nil {
		return err
	}

	*o = NewMarginOrderResponse(varNewMarginOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "transactTime")
		delete(additionalProperties, "price")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "cummulativeQuoteQty")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "type")
		delete(additionalProperties, "side")
		delete(additionalProperties, "marginBuyBorrowAmount")
		delete(additionalProperties, "marginBuyBorrowAsset")
		delete(additionalProperties, "fills")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNewMarginOrderResponse struct {
	value *NewMarginOrderResponse
	isSet bool
}

func (v NullableNewMarginOrderResponse) Get() *NewMarginOrderResponse {
	return v.value
}

func (v *NullableNewMarginOrderResponse) Set(val *NewMarginOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNewMarginOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNewMarginOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewMarginOrderResponse(val *NewMarginOrderResponse) *NullableNewMarginOrderResponse {
	return &NullableNewMarginOrderResponse{value: val, isSet: true}
}

func (v NullableNewMarginOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewMarginOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
