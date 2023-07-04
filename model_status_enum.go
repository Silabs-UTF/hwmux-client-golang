/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.19.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
	"fmt"
)

// StatusEnum the model 'StatusEnum'
type StatusEnum string

// List of StatusEnum
const (
	ACTIVE StatusEnum = "ACTIVE"
	DISABLED StatusEnum = "DISABLED"
	OFFLINE StatusEnum = "OFFLINE"
)

// All allowed values of StatusEnum enum
var AllowedStatusEnumEnumValues = []StatusEnum{
	"ACTIVE",
	"DISABLED",
	"OFFLINE",
}

func (v *StatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StatusEnum(value)
	for _, existing := range AllowedStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StatusEnum", value)
}

// NewStatusEnumFromValue returns a pointer to a valid StatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatusEnumFromValue(v string) (*StatusEnum, error) {
	ev := StatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StatusEnum: valid values are %v", v, AllowedStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatusEnum) IsValid() bool {
	for _, existing := range AllowedStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusEnum value
func (v StatusEnum) Ptr() *StatusEnum {
	return &v
}

type NullableStatusEnum struct {
	value *StatusEnum
	isSet bool
}

func (v NullableStatusEnum) Get() *StatusEnum {
	return v.value
}

func (v *NullableStatusEnum) Set(val *StatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusEnum(val *StatusEnum) *NullableStatusEnum {
	return &NullableStatusEnum{value: val, isSet: true}
}

func (v NullableStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

