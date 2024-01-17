/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
)

// ReservationRequest struct for ReservationRequest
type ReservationRequest struct {
	Details *string `json:"details,omitempty"`
	RDevices []int32 `json:"r_devices,omitempty"`
	RDeviceGroups []int32 `json:"r_device_groups,omitempty"`
	RDeviceGroupLabels []int32 `json:"r_device_group_labels,omitempty"`
	UseWatchdog *bool `json:"use_watchdog,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	InvertPriority *bool `json:"invert_priority,omitempty"`
}

// NewReservationRequest instantiates a new ReservationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationRequest() *ReservationRequest {
	this := ReservationRequest{}
	var invertPriority bool = false
	this.InvertPriority = &invertPriority
	return &this
}

// NewReservationRequestWithDefaults instantiates a new ReservationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationRequestWithDefaults() *ReservationRequest {
	this := ReservationRequest{}
	var invertPriority bool = false
	this.InvertPriority = &invertPriority
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ReservationRequest) GetDetails() string {
	if o == nil || isNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationRequest) GetDetailsOk() (*string, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ReservationRequest) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ReservationRequest) SetDetails(v string) {
	o.Details = &v
}

// GetRDevices returns the RDevices field value if set, zero value otherwise.
func (o *ReservationRequest) GetRDevices() []int32 {
	if o == nil || isNil(o.RDevices) {
		var ret []int32
		return ret
	}
	return o.RDevices
}

// GetRDevicesOk returns a tuple with the RDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationRequest) GetRDevicesOk() ([]int32, bool) {
	if o == nil || isNil(o.RDevices) {
    return nil, false
	}
	return o.RDevices, true
}

// HasRDevices returns a boolean if a field has been set.
func (o *ReservationRequest) HasRDevices() bool {
	if o != nil && !isNil(o.RDevices) {
		return true
	}

	return false
}

// SetRDevices gets a reference to the given []int32 and assigns it to the RDevices field.
func (o *ReservationRequest) SetRDevices(v []int32) {
	o.RDevices = v
}

// GetRDeviceGroups returns the RDeviceGroups field value if set, zero value otherwise.
func (o *ReservationRequest) GetRDeviceGroups() []int32 {
	if o == nil || isNil(o.RDeviceGroups) {
		var ret []int32
		return ret
	}
	return o.RDeviceGroups
}

// GetRDeviceGroupsOk returns a tuple with the RDeviceGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationRequest) GetRDeviceGroupsOk() ([]int32, bool) {
	if o == nil || isNil(o.RDeviceGroups) {
    return nil, false
	}
	return o.RDeviceGroups, true
}

// HasRDeviceGroups returns a boolean if a field has been set.
func (o *ReservationRequest) HasRDeviceGroups() bool {
	if o != nil && !isNil(o.RDeviceGroups) {
		return true
	}

	return false
}

// SetRDeviceGroups gets a reference to the given []int32 and assigns it to the RDeviceGroups field.
func (o *ReservationRequest) SetRDeviceGroups(v []int32) {
	o.RDeviceGroups = v
}

// GetRDeviceGroupLabels returns the RDeviceGroupLabels field value if set, zero value otherwise.
func (o *ReservationRequest) GetRDeviceGroupLabels() []int32 {
	if o == nil || isNil(o.RDeviceGroupLabels) {
		var ret []int32
		return ret
	}
	return o.RDeviceGroupLabels
}

// GetRDeviceGroupLabelsOk returns a tuple with the RDeviceGroupLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationRequest) GetRDeviceGroupLabelsOk() ([]int32, bool) {
	if o == nil || isNil(o.RDeviceGroupLabels) {
    return nil, false
	}
	return o.RDeviceGroupLabels, true
}

// HasRDeviceGroupLabels returns a boolean if a field has been set.
func (o *ReservationRequest) HasRDeviceGroupLabels() bool {
	if o != nil && !isNil(o.RDeviceGroupLabels) {
		return true
	}

	return false
}

// SetRDeviceGroupLabels gets a reference to the given []int32 and assigns it to the RDeviceGroupLabels field.
func (o *ReservationRequest) SetRDeviceGroupLabels(v []int32) {
	o.RDeviceGroupLabels = v
}

// GetUseWatchdog returns the UseWatchdog field value if set, zero value otherwise.
func (o *ReservationRequest) GetUseWatchdog() bool {
	if o == nil || isNil(o.UseWatchdog) {
		var ret bool
		return ret
	}
	return *o.UseWatchdog
}

// GetUseWatchdogOk returns a tuple with the UseWatchdog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationRequest) GetUseWatchdogOk() (*bool, bool) {
	if o == nil || isNil(o.UseWatchdog) {
    return nil, false
	}
	return o.UseWatchdog, true
}

// HasUseWatchdog returns a boolean if a field has been set.
func (o *ReservationRequest) HasUseWatchdog() bool {
	if o != nil && !isNil(o.UseWatchdog) {
		return true
	}

	return false
}

// SetUseWatchdog gets a reference to the given bool and assigns it to the UseWatchdog field.
func (o *ReservationRequest) SetUseWatchdog(v bool) {
	o.UseWatchdog = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ReservationRequest) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ReservationRequest) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ReservationRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetInvertPriority returns the InvertPriority field value if set, zero value otherwise.
func (o *ReservationRequest) GetInvertPriority() bool {
	if o == nil || isNil(o.InvertPriority) {
		var ret bool
		return ret
	}
	return *o.InvertPriority
}

// GetInvertPriorityOk returns a tuple with the InvertPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationRequest) GetInvertPriorityOk() (*bool, bool) {
	if o == nil || isNil(o.InvertPriority) {
    return nil, false
	}
	return o.InvertPriority, true
}

// HasInvertPriority returns a boolean if a field has been set.
func (o *ReservationRequest) HasInvertPriority() bool {
	if o != nil && !isNil(o.InvertPriority) {
		return true
	}

	return false
}

// SetInvertPriority gets a reference to the given bool and assigns it to the InvertPriority field.
func (o *ReservationRequest) SetInvertPriority(v bool) {
	o.InvertPriority = &v
}

func (o ReservationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !isNil(o.RDevices) {
		toSerialize["r_devices"] = o.RDevices
	}
	if !isNil(o.RDeviceGroups) {
		toSerialize["r_device_groups"] = o.RDeviceGroups
	}
	if !isNil(o.RDeviceGroupLabels) {
		toSerialize["r_device_group_labels"] = o.RDeviceGroupLabels
	}
	if !isNil(o.UseWatchdog) {
		toSerialize["use_watchdog"] = o.UseWatchdog
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.InvertPriority) {
		toSerialize["invert_priority"] = o.InvertPriority
	}
	return json.Marshal(toSerialize)
}

type NullableReservationRequest struct {
	value *ReservationRequest
	isSet bool
}

func (v NullableReservationRequest) Get() *ReservationRequest {
	return v.value
}

func (v *NullableReservationRequest) Set(val *ReservationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationRequest(val *ReservationRequest) *NullableReservationRequest {
	return &NullableReservationRequest{value: val, isSet: true}
}

func (v NullableReservationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


