/*
Spot REST API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ExchangeInfoPermissionsParameterInner the model 'ExchangeInfoPermissionsParameterInner'
type ExchangeInfoPermissionsParameterInner string

// List of exchangeInfo_permissions_parameter_inner
const (
	ExchangeInfoPermissionsParameterInnerSpot      ExchangeInfoPermissionsParameterInner = "SPOT"
	ExchangeInfoPermissionsParameterInnerMargin    ExchangeInfoPermissionsParameterInner = "MARGIN"
	ExchangeInfoPermissionsParameterInnerLeveraged ExchangeInfoPermissionsParameterInner = "LEVERAGED"
	ExchangeInfoPermissionsParameterInnerTrdGrp002 ExchangeInfoPermissionsParameterInner = "TRD_GRP_002"
	ExchangeInfoPermissionsParameterInnerTrdGrp003 ExchangeInfoPermissionsParameterInner = "TRD_GRP_003"
	ExchangeInfoPermissionsParameterInnerTrdGrp004 ExchangeInfoPermissionsParameterInner = "TRD_GRP_004"
	ExchangeInfoPermissionsParameterInnerTrdGrp005 ExchangeInfoPermissionsParameterInner = "TRD_GRP_005"
	ExchangeInfoPermissionsParameterInnerTrdGrp006 ExchangeInfoPermissionsParameterInner = "TRD_GRP_006"
	ExchangeInfoPermissionsParameterInnerTrdGrp007 ExchangeInfoPermissionsParameterInner = "TRD_GRP_007"
	ExchangeInfoPermissionsParameterInnerTrdGrp008 ExchangeInfoPermissionsParameterInner = "TRD_GRP_008"
	ExchangeInfoPermissionsParameterInnerTrdGrp009 ExchangeInfoPermissionsParameterInner = "TRD_GRP_009"
	ExchangeInfoPermissionsParameterInnerTrdGrp010 ExchangeInfoPermissionsParameterInner = "TRD_GRP_010"
	ExchangeInfoPermissionsParameterInnerTrdGrp011 ExchangeInfoPermissionsParameterInner = "TRD_GRP_011"
	ExchangeInfoPermissionsParameterInnerTrdGrp012 ExchangeInfoPermissionsParameterInner = "TRD_GRP_012"
	ExchangeInfoPermissionsParameterInnerTrdGrp013 ExchangeInfoPermissionsParameterInner = "TRD_GRP_013"
	ExchangeInfoPermissionsParameterInnerTrdGrp014 ExchangeInfoPermissionsParameterInner = "TRD_GRP_014"
	ExchangeInfoPermissionsParameterInnerTrdGrp015 ExchangeInfoPermissionsParameterInner = "TRD_GRP_015"
	ExchangeInfoPermissionsParameterInnerTrdGrp016 ExchangeInfoPermissionsParameterInner = "TRD_GRP_016"
	ExchangeInfoPermissionsParameterInnerTrdGrp017 ExchangeInfoPermissionsParameterInner = "TRD_GRP_017"
	ExchangeInfoPermissionsParameterInnerTrdGrp018 ExchangeInfoPermissionsParameterInner = "TRD_GRP_018"
	ExchangeInfoPermissionsParameterInnerTrdGrp019 ExchangeInfoPermissionsParameterInner = "TRD_GRP_019"
	ExchangeInfoPermissionsParameterInnerTrdGrp020 ExchangeInfoPermissionsParameterInner = "TRD_GRP_020"
	ExchangeInfoPermissionsParameterInnerTrdGrp021 ExchangeInfoPermissionsParameterInner = "TRD_GRP_021"
	ExchangeInfoPermissionsParameterInnerTrdGrp022 ExchangeInfoPermissionsParameterInner = "TRD_GRP_022"
	ExchangeInfoPermissionsParameterInnerTrdGrp023 ExchangeInfoPermissionsParameterInner = "TRD_GRP_023"
	ExchangeInfoPermissionsParameterInnerTrdGrp024 ExchangeInfoPermissionsParameterInner = "TRD_GRP_024"
	ExchangeInfoPermissionsParameterInnerTrdGrp025 ExchangeInfoPermissionsParameterInner = "TRD_GRP_025"
)

// All allowed values of ExchangeInfoPermissionsParameterInner enum
var AllowedExchangeInfoPermissionsParameterInnerEnumValues = []ExchangeInfoPermissionsParameterInner{
	"SPOT",
	"MARGIN",
	"LEVERAGED",
	"TRD_GRP_002",
	"TRD_GRP_003",
	"TRD_GRP_004",
	"TRD_GRP_005",
	"TRD_GRP_006",
	"TRD_GRP_007",
	"TRD_GRP_008",
	"TRD_GRP_009",
	"TRD_GRP_010",
	"TRD_GRP_011",
	"TRD_GRP_012",
	"TRD_GRP_013",
	"TRD_GRP_014",
	"TRD_GRP_015",
	"TRD_GRP_016",
	"TRD_GRP_017",
	"TRD_GRP_018",
	"TRD_GRP_019",
	"TRD_GRP_020",
	"TRD_GRP_021",
	"TRD_GRP_022",
	"TRD_GRP_023",
	"TRD_GRP_024",
	"TRD_GRP_025",
}

func (v *ExchangeInfoPermissionsParameterInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExchangeInfoPermissionsParameterInner(value)
	for _, existing := range AllowedExchangeInfoPermissionsParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExchangeInfoPermissionsParameterInner", value)
}

// NewExchangeInfoPermissionsParameterInnerFromValue returns a pointer to a valid ExchangeInfoPermissionsParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExchangeInfoPermissionsParameterInnerFromValue(v string) (*ExchangeInfoPermissionsParameterInner, error) {
	ev := ExchangeInfoPermissionsParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExchangeInfoPermissionsParameterInner: valid values are %v", v, AllowedExchangeInfoPermissionsParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExchangeInfoPermissionsParameterInner) IsValid() bool {
	for _, existing := range AllowedExchangeInfoPermissionsParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to exchangeInfo_permissions_parameter_inner value
func (v ExchangeInfoPermissionsParameterInner) Ptr() *ExchangeInfoPermissionsParameterInner {
	return &v
}

type NullableExchangeInfoPermissionsParameterInner struct {
	value *ExchangeInfoPermissionsParameterInner
	isSet bool
}

func (v NullableExchangeInfoPermissionsParameterInner) Get() *ExchangeInfoPermissionsParameterInner {
	return v.value
}

func (v *NullableExchangeInfoPermissionsParameterInner) Set(val *ExchangeInfoPermissionsParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInfoPermissionsParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInfoPermissionsParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInfoPermissionsParameterInner(val *ExchangeInfoPermissionsParameterInner) *NullableExchangeInfoPermissionsParameterInner {
	return &NullableExchangeInfoPermissionsParameterInner{value: val, isSet: true}
}

func (v NullableExchangeInfoPermissionsParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInfoPermissionsParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
