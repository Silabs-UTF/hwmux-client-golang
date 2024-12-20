/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
)

// ReservationExtensionRequest struct for ReservationExtensionRequest
type ReservationExtensionRequest struct {
	ExtensionDurationS int32 `json:"extension_duration_s"`
	ExtendExisting *bool `json:"extend_existing,omitempty"`
}

// NewReservationExtensionRequest instantiates a new ReservationExtensionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationExtensionRequest(extensionDurationS int32) *ReservationExtensionRequest {
	this := ReservationExtensionRequest{}
	this.ExtensionDurationS = extensionDurationS
	var extendExisting bool = false
	this.ExtendExisting = &extendExisting
	return &this
}

// NewReservationExtensionRequestWithDefaults instantiates a new ReservationExtensionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationExtensionRequestWithDefaults() *ReservationExtensionRequest {
	this := ReservationExtensionRequest{}
	var extendExisting bool = false
	this.ExtendExisting = &extendExisting
	return &this
}

// GetExtensionDurationS returns the ExtensionDurationS field value
func (o *ReservationExtensionRequest) GetExtensionDurationS() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExtensionDurationS
}

// GetExtensionDurationSOk returns a tuple with the ExtensionDurationS field value
// and a boolean to check if the value has been set.
func (o *ReservationExtensionRequest) GetExtensionDurationSOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionDurationS, true
}

// SetExtensionDurationS sets field value
func (o *ReservationExtensionRequest) SetExtensionDurationS(v int32) {
	o.ExtensionDurationS = v
}

// GetExtendExisting returns the ExtendExisting field value if set, zero value otherwise.
func (o *ReservationExtensionRequest) GetExtendExisting() bool {
	if o == nil || isNil(o.ExtendExisting) {
		var ret bool
		return ret
	}
	return *o.ExtendExisting
}

// GetExtendExistingOk returns a tuple with the ExtendExisting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationExtensionRequest) GetExtendExistingOk() (*bool, bool) {
	if o == nil || isNil(o.ExtendExisting) {
    return nil, false
	}
	return o.ExtendExisting, true
}

// HasExtendExisting returns a boolean if a field has been set.
func (o *ReservationExtensionRequest) HasExtendExisting() bool {
	if o != nil && !isNil(o.ExtendExisting) {
		return true
	}

	return false
}

// SetExtendExisting gets a reference to the given bool and assigns it to the ExtendExisting field.
func (o *ReservationExtensionRequest) SetExtendExisting(v bool) {
	o.ExtendExisting = &v
}

func (o ReservationExtensionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["extension_duration_s"] = o.ExtensionDurationS
	}
	if !isNil(o.ExtendExisting) {
		toSerialize["extend_existing"] = o.ExtendExisting
	}
	return json.Marshal(toSerialize)
}

type NullableReservationExtensionRequest struct {
	value *ReservationExtensionRequest
	isSet bool
}

func (v NullableReservationExtensionRequest) Get() *ReservationExtensionRequest {
	return v.value
}

func (v *NullableReservationExtensionRequest) Set(val *ReservationExtensionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationExtensionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationExtensionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationExtensionRequest(val *ReservationExtensionRequest) *NullableReservationExtensionRequest {
	return &NullableReservationExtensionRequest{value: val, isSet: true}
}

func (v NullableReservationExtensionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationExtensionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


