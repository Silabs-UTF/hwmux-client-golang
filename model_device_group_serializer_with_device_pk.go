/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DeviceGroupSerializerWithDevicePk struct for DeviceGroupSerializerWithDevicePk
type DeviceGroupSerializerWithDevicePk struct {
	Id int32 `json:"id"`
	Devices []int32 `json:"devices,omitempty"`
	IsReserved NullableBool `json:"is_reserved"`
	PermissionGroups []string `json:"permission_groups,omitempty"`
	Name string `json:"name"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewDeviceGroupSerializerWithDevicePk instantiates a new DeviceGroupSerializerWithDevicePk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceGroupSerializerWithDevicePk(id int32, isReserved NullableBool, name string) *DeviceGroupSerializerWithDevicePk {
	this := DeviceGroupSerializerWithDevicePk{}
	this.Id = id
	this.IsReserved = isReserved
	this.Name = name
	return &this
}

// NewDeviceGroupSerializerWithDevicePkWithDefaults instantiates a new DeviceGroupSerializerWithDevicePk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceGroupSerializerWithDevicePkWithDefaults() *DeviceGroupSerializerWithDevicePk {
	this := DeviceGroupSerializerWithDevicePk{}
	return &this
}

// GetId returns the Id field value
func (o *DeviceGroupSerializerWithDevicePk) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupSerializerWithDevicePk) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceGroupSerializerWithDevicePk) SetId(v int32) {
	o.Id = v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *DeviceGroupSerializerWithDevicePk) GetDevices() []int32 {
	if o == nil || isNil(o.Devices) {
		var ret []int32
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGroupSerializerWithDevicePk) GetDevicesOk() ([]int32, bool) {
	if o == nil || isNil(o.Devices) {
    return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *DeviceGroupSerializerWithDevicePk) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []int32 and assigns it to the Devices field.
func (o *DeviceGroupSerializerWithDevicePk) SetDevices(v []int32) {
	o.Devices = v
}

// GetIsReserved returns the IsReserved field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *DeviceGroupSerializerWithDevicePk) GetIsReserved() bool {
	if o == nil || o.IsReserved.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsReserved.Get()
}

// GetIsReservedOk returns a tuple with the IsReserved field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceGroupSerializerWithDevicePk) GetIsReservedOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsReserved.Get(), o.IsReserved.IsSet()
}

// SetIsReserved sets field value
func (o *DeviceGroupSerializerWithDevicePk) SetIsReserved(v bool) {
	o.IsReserved.Set(&v)
}

// GetPermissionGroups returns the PermissionGroups field value if set, zero value otherwise.
func (o *DeviceGroupSerializerWithDevicePk) GetPermissionGroups() []string {
	if o == nil || isNil(o.PermissionGroups) {
		var ret []string
		return ret
	}
	return o.PermissionGroups
}

// GetPermissionGroupsOk returns a tuple with the PermissionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGroupSerializerWithDevicePk) GetPermissionGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.PermissionGroups) {
    return nil, false
	}
	return o.PermissionGroups, true
}

// HasPermissionGroups returns a boolean if a field has been set.
func (o *DeviceGroupSerializerWithDevicePk) HasPermissionGroups() bool {
	if o != nil && !isNil(o.PermissionGroups) {
		return true
	}

	return false
}

// SetPermissionGroups gets a reference to the given []string and assigns it to the PermissionGroups field.
func (o *DeviceGroupSerializerWithDevicePk) SetPermissionGroups(v []string) {
	o.PermissionGroups = v
}

// GetName returns the Name field value
func (o *DeviceGroupSerializerWithDevicePk) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupSerializerWithDevicePk) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceGroupSerializerWithDevicePk) SetName(v string) {
	o.Name = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceGroupSerializerWithDevicePk) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceGroupSerializerWithDevicePk) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DeviceGroupSerializerWithDevicePk) HasMetadata() bool {
	if o != nil && isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *DeviceGroupSerializerWithDevicePk) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o DeviceGroupSerializerWithDevicePk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if true {
		toSerialize["is_reserved"] = o.IsReserved.Get()
	}
	if !isNil(o.PermissionGroups) {
		toSerialize["permission_groups"] = o.PermissionGroups
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceGroupSerializerWithDevicePk struct {
	value *DeviceGroupSerializerWithDevicePk
	isSet bool
}

func (v NullableDeviceGroupSerializerWithDevicePk) Get() *DeviceGroupSerializerWithDevicePk {
	return v.value
}

func (v *NullableDeviceGroupSerializerWithDevicePk) Set(val *DeviceGroupSerializerWithDevicePk) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceGroupSerializerWithDevicePk) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceGroupSerializerWithDevicePk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceGroupSerializerWithDevicePk(val *DeviceGroupSerializerWithDevicePk) *NullableDeviceGroupSerializerWithDevicePk {
	return &NullableDeviceGroupSerializerWithDevicePk{value: val, isSet: true}
}

func (v NullableDeviceGroupSerializerWithDevicePk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceGroupSerializerWithDevicePk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


