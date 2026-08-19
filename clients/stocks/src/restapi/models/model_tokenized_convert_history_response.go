/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TokenizedConvertHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TokenizedConvertHistoryResponse{}

// TokenizedConvertHistoryResponse struct for TokenizedConvertHistoryResponse
type TokenizedConvertHistoryResponse struct {
	// Convert history rows on this page. Empty array if nothing matches.
	Rows []TokenizedConvertHistoryResponseRowsInner `json:"rows,omitempty"`
	// `true` when more pages exist — pass `nextLastId` as `lastId` on the next request.
	HasMore *bool `json:"hasMore,omitempty"`
	// Pass this value as `lastId` on the next request to get the following page.
	NextLastId           *int64 `json:"nextLastId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenizedConvertHistoryResponse TokenizedConvertHistoryResponse

// NewTokenizedConvertHistoryResponse instantiates a new TokenizedConvertHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizedConvertHistoryResponse() *TokenizedConvertHistoryResponse {
	this := TokenizedConvertHistoryResponse{}
	return &this
}

// NewTokenizedConvertHistoryResponseWithDefaults instantiates a new TokenizedConvertHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizedConvertHistoryResponseWithDefaults() *TokenizedConvertHistoryResponse {
	this := TokenizedConvertHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *TokenizedConvertHistoryResponse) GetRows() []TokenizedConvertHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []TokenizedConvertHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertHistoryResponse) GetRowsOk() ([]TokenizedConvertHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *TokenizedConvertHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []TokenizedConvertHistoryResponseRowsInner and assigns it to the Rows field.
func (o *TokenizedConvertHistoryResponse) SetRows(v []TokenizedConvertHistoryResponseRowsInner) {
	o.Rows = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *TokenizedConvertHistoryResponse) GetHasMore() bool {
	if o == nil || common.IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertHistoryResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil || common.IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *TokenizedConvertHistoryResponse) HasHasMore() bool {
	if o != nil && !common.IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *TokenizedConvertHistoryResponse) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetNextLastId returns the NextLastId field value if set, zero value otherwise.
func (o *TokenizedConvertHistoryResponse) GetNextLastId() int64 {
	if o == nil || common.IsNil(o.NextLastId) {
		var ret int64
		return ret
	}
	return *o.NextLastId
}

// GetNextLastIdOk returns a tuple with the NextLastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertHistoryResponse) GetNextLastIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.NextLastId) {
		return nil, false
	}
	return o.NextLastId, true
}

// HasNextLastId returns a boolean if a field has been set.
func (o *TokenizedConvertHistoryResponse) HasNextLastId() bool {
	if o != nil && !common.IsNil(o.NextLastId) {
		return true
	}

	return false
}

// SetNextLastId gets a reference to the given int64 and assigns it to the NextLastId field.
func (o *TokenizedConvertHistoryResponse) SetNextLastId(v int64) {
	o.NextLastId = &v
}

func (o TokenizedConvertHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizedConvertHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !common.IsNil(o.HasMore) {
		toSerialize["hasMore"] = o.HasMore
	}
	if !common.IsNil(o.NextLastId) {
		toSerialize["nextLastId"] = o.NextLastId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenizedConvertHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varTokenizedConvertHistoryResponse := _TokenizedConvertHistoryResponse{}

	err = json.Unmarshal(data, &varTokenizedConvertHistoryResponse)

	if err != nil {
		return err
	}

	*o = TokenizedConvertHistoryResponse(varTokenizedConvertHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "hasMore")
		delete(additionalProperties, "nextLastId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenizedConvertHistoryResponse struct {
	value *TokenizedConvertHistoryResponse
	isSet bool
}

func (v NullableTokenizedConvertHistoryResponse) Get() *TokenizedConvertHistoryResponse {
	return v.value
}

func (v *NullableTokenizedConvertHistoryResponse) Set(val *TokenizedConvertHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizedConvertHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizedConvertHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizedConvertHistoryResponse(val *TokenizedConvertHistoryResponse) *NullableTokenizedConvertHistoryResponse {
	return &NullableTokenizedConvertHistoryResponse{value: val, isSet: true}
}

func (v NullableTokenizedConvertHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizedConvertHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
