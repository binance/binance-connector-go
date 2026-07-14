/*
Convert REST API

Request quotes and execute cryptocurrency conversions via the Convert REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// SendQuoteRequestValidTimeParameter the model 'SendQuoteRequestValidTimeParameter'
type SendQuoteRequestValidTimeParameter string

// List of sendQuoteRequest_validTime_parameter
const (
	SendQuoteRequestValidTimeParameterValidTime10s SendQuoteRequestValidTimeParameter = "10s"
	SendQuoteRequestValidTimeParameterValidTime30s SendQuoteRequestValidTimeParameter = "30s"
	SendQuoteRequestValidTimeParameterValidTime1m  SendQuoteRequestValidTimeParameter = "1m"
)

// All allowed values of SendQuoteRequestValidTimeParameter enum
var AllowedSendQuoteRequestValidTimeParameterEnumValues = []SendQuoteRequestValidTimeParameter{
	"10s",
	"30s",
	"1m",
}

func (v *SendQuoteRequestValidTimeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SendQuoteRequestValidTimeParameter(value)
	for _, existing := range AllowedSendQuoteRequestValidTimeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SendQuoteRequestValidTimeParameter", value)
}

// NewSendQuoteRequestValidTimeParameterFromValue returns a pointer to a valid SendQuoteRequestValidTimeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSendQuoteRequestValidTimeParameterFromValue(v string) (*SendQuoteRequestValidTimeParameter, error) {
	ev := SendQuoteRequestValidTimeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SendQuoteRequestValidTimeParameter: valid values are %v", v, AllowedSendQuoteRequestValidTimeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SendQuoteRequestValidTimeParameter) IsValid() bool {
	for _, existing := range AllowedSendQuoteRequestValidTimeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to sendQuoteRequest_validTime_parameter value
func (v SendQuoteRequestValidTimeParameter) Ptr() *SendQuoteRequestValidTimeParameter {
	return &v
}

type NullableSendQuoteRequestValidTimeParameter struct {
	value *SendQuoteRequestValidTimeParameter
	isSet bool
}

func (v NullableSendQuoteRequestValidTimeParameter) Get() *SendQuoteRequestValidTimeParameter {
	return v.value
}

func (v *NullableSendQuoteRequestValidTimeParameter) Set(val *SendQuoteRequestValidTimeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableSendQuoteRequestValidTimeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableSendQuoteRequestValidTimeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendQuoteRequestValidTimeParameter(val *SendQuoteRequestValidTimeParameter) *NullableSendQuoteRequestValidTimeParameter {
	return &NullableSendQuoteRequestValidTimeParameter{value: val, isSet: true}
}

func (v NullableSendQuoteRequestValidTimeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendQuoteRequestValidTimeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
