/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
	"fmt"
)

// LogStatusEnum * `SUCCESS` - Success * `ERROR` - Error
type LogStatusEnum string

// List of LogStatusEnum
const (
	SUCCESS LogStatusEnum = "SUCCESS"
	ERROR LogStatusEnum = "ERROR"
)

// All allowed values of LogStatusEnum enum
var AllowedLogStatusEnumEnumValues = []LogStatusEnum{
	"SUCCESS",
	"ERROR",
}

func (v *LogStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogStatusEnum(value)
	for _, existing := range AllowedLogStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogStatusEnum", value)
}

// NewLogStatusEnumFromValue returns a pointer to a valid LogStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogStatusEnumFromValue(v string) (*LogStatusEnum, error) {
	ev := LogStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogStatusEnum: valid values are %v", v, AllowedLogStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogStatusEnum) IsValid() bool {
	for _, existing := range AllowedLogStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogStatusEnum value
func (v LogStatusEnum) Ptr() *LogStatusEnum {
	return &v
}

type NullableLogStatusEnum struct {
	value *LogStatusEnum
	isSet bool
}

func (v NullableLogStatusEnum) Get() *LogStatusEnum {
	return v.value
}

func (v *NullableLogStatusEnum) Set(val *LogStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStatusEnum(val *LogStatusEnum) *NullableLogStatusEnum {
	return &NullableLogStatusEnum{value: val, isSet: true}
}

func (v NullableLogStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

