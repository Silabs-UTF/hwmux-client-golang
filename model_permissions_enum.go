/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
	"fmt"
)

// PermissionsEnum the model 'PermissionsEnum'
type PermissionsEnum string

// List of PermissionsEnum
const (
	VIEW PermissionsEnum = "view"
	CHANGE PermissionsEnum = "change"
	ADD PermissionsEnum = "add"
	DELETE PermissionsEnum = "delete"
)

// All allowed values of PermissionsEnum enum
var AllowedPermissionsEnumEnumValues = []PermissionsEnum{
	"view",
	"change",
	"add",
	"delete",
}

func (v *PermissionsEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PermissionsEnum(value)
	for _, existing := range AllowedPermissionsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PermissionsEnum", value)
}

// NewPermissionsEnumFromValue returns a pointer to a valid PermissionsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPermissionsEnumFromValue(v string) (*PermissionsEnum, error) {
	ev := PermissionsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PermissionsEnum: valid values are %v", v, AllowedPermissionsEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PermissionsEnum) IsValid() bool {
	for _, existing := range AllowedPermissionsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PermissionsEnum value
func (v PermissionsEnum) Ptr() *PermissionsEnum {
	return &v
}

type NullablePermissionsEnum struct {
	value *PermissionsEnum
	isSet bool
}

func (v NullablePermissionsEnum) Get() *PermissionsEnum {
	return v.value
}

func (v *NullablePermissionsEnum) Set(val *PermissionsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsEnum(val *PermissionsEnum) *NullablePermissionsEnum {
	return &NullablePermissionsEnum{value: val, isSet: true}
}

func (v NullablePermissionsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

