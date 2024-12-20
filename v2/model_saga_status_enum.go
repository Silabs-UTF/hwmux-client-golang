/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
	"fmt"
)

// SagaStatusEnum * `started` - started * `failed` - failed * `succeeded` - succeeded * `processing` - processing * `compensating` - compensating * `compensated` - compensated
type SagaStatusEnum string

// List of SagaStatusEnum
const (
	SAGASTATUSENUM_STARTED SagaStatusEnum = "started"
	SAGASTATUSENUM_FAILED SagaStatusEnum = "failed"
	SAGASTATUSENUM_SUCCEEDED SagaStatusEnum = "succeeded"
	SAGASTATUSENUM_PROCESSING SagaStatusEnum = "processing"
	SAGASTATUSENUM_COMPENSATING SagaStatusEnum = "compensating"
	SAGASTATUSENUM_COMPENSATED SagaStatusEnum = "compensated"
)

// All allowed values of SagaStatusEnum enum
var AllowedSagaStatusEnumEnumValues = []SagaStatusEnum{
	"started",
	"failed",
	"succeeded",
	"processing",
	"compensating",
	"compensated",
}

func (v *SagaStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SagaStatusEnum(value)
	for _, existing := range AllowedSagaStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SagaStatusEnum", value)
}

// NewSagaStatusEnumFromValue returns a pointer to a valid SagaStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSagaStatusEnumFromValue(v string) (*SagaStatusEnum, error) {
	ev := SagaStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SagaStatusEnum: valid values are %v", v, AllowedSagaStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SagaStatusEnum) IsValid() bool {
	for _, existing := range AllowedSagaStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SagaStatusEnum value
func (v SagaStatusEnum) Ptr() *SagaStatusEnum {
	return &v
}

type NullableSagaStatusEnum struct {
	value *SagaStatusEnum
	isSet bool
}

func (v NullableSagaStatusEnum) Get() *SagaStatusEnum {
	return v.value
}

func (v *NullableSagaStatusEnum) Set(val *SagaStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSagaStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSagaStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSagaStatusEnum(val *SagaStatusEnum) *NullableSagaStatusEnum {
	return &NullableSagaStatusEnum{value: val, isSet: true}
}

func (v NullableSagaStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSagaStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

