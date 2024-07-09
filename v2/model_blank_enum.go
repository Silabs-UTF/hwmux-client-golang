/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
	"fmt"
)

// BlankEnum the model 'BlankEnum'
type BlankEnum string

// List of BlankEnum
const (
	EMPTY BlankEnum = ""
)

// All allowed values of BlankEnum enum
var AllowedBlankEnumEnumValues = []BlankEnum{
	"",
}

func (v *BlankEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BlankEnum(value)
	for _, existing := range AllowedBlankEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BlankEnum", value)
}

// NewBlankEnumFromValue returns a pointer to a valid BlankEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBlankEnumFromValue(v string) (*BlankEnum, error) {
	ev := BlankEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BlankEnum: valid values are %v", v, AllowedBlankEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BlankEnum) IsValid() bool {
	for _, existing := range AllowedBlankEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BlankEnum value
func (v BlankEnum) Ptr() *BlankEnum {
	return &v
}

type NullableBlankEnum struct {
	value *BlankEnum
	isSet bool
}

func (v NullableBlankEnum) Get() *BlankEnum {
	return v.value
}

func (v *NullableBlankEnum) Set(val *BlankEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBlankEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBlankEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlankEnum(val *BlankEnum) *NullableBlankEnum {
	return &NullableBlankEnum{value: val, isSet: true}
}

func (v NullableBlankEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlankEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

