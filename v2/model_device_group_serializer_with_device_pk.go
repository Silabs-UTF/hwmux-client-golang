/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
)

// DeviceGroupSerializerWithDevicePk struct for DeviceGroupSerializerWithDevicePk
type DeviceGroupSerializerWithDevicePk struct {
	Id int32 `json:"id"`
	PermissionGroups []string `json:"permission_groups,omitempty"`
	Devices []int32 `json:"devices,omitempty"`
	Online bool `json:"online"`
	Status string `json:"status"`
	LocDesc []string `json:"loc_desc"`
	IsReserved NullableBool `json:"is_reserved"`
	Name string `json:"name"`
	EnableAhs *bool `json:"enable_ahs,omitempty"`
	EnableAhsActions *bool `json:"enable_ahs_actions,omitempty"`
	EnableAhsCas *bool `json:"enable_ahs_cas,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Source *SourceEnum `json:"source,omitempty"`
	IsReservedFilter *bool `json:"is_reserved_filter,omitempty"`
}

// NewDeviceGroupSerializerWithDevicePk instantiates a new DeviceGroupSerializerWithDevicePk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceGroupSerializerWithDevicePk(id int32, online bool, status string, locDesc []string, isReserved NullableBool, name string) *DeviceGroupSerializerWithDevicePk {
	this := DeviceGroupSerializerWithDevicePk{}
	this.Id = id
	this.Online = online
	this.Status = status
	this.LocDesc = locDesc
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

// GetPermissionGroups returns the PermissionGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceGroupSerializerWithDevicePk) GetPermissionGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PermissionGroups
}

// GetPermissionGroupsOk returns a tuple with the PermissionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceGroupSerializerWithDevicePk) GetPermissionGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.PermissionGroups) {
    return nil, false
	}
	return o.PermissionGroups, true
}

// HasPermissionGroups returns a boolean if a field has been set.
func (o *DeviceGroupSerializerWithDevicePk) HasPermissionGroups() bool {
	if o != nil && isNil(o.PermissionGroups) {
		return true
	}

	return false
}

// SetPermissionGroups gets a reference to the given []string and assigns it to the PermissionGroups field.
func (o *DeviceGroupSerializerWithDevicePk) SetPermissionGroups(v []string) {
	o.PermissionGroups = v
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

// GetOnline returns the Online field value
func (o *DeviceGroupSerializerWithDevicePk) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupSerializerWithDevicePk) GetOnlineOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *DeviceGroupSerializerWithDevicePk) SetOnline(v bool) {
	o.Online = v
}

// GetStatus returns the Status field value
func (o *DeviceGroupSerializerWithDevicePk) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupSerializerWithDevicePk) GetStatusOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeviceGroupSerializerWithDevicePk) SetStatus(v string) {
	o.Status = v
}

// GetLocDesc returns the LocDesc field value
func (o *DeviceGroupSerializerWithDevicePk) GetLocDesc() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LocDesc
}

// GetLocDescOk returns a tuple with the LocDesc field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupSerializerWithDevicePk) GetLocDescOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.LocDesc, true
}

// SetLocDesc sets field value
func (o *DeviceGroupSerializerWithDevicePk) SetLocDesc(v []string) {
	o.LocDesc = v
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

// GetEnableAhs returns the EnableAhs field value if set, zero value otherwise.
func (o *DeviceGroupSerializerWithDevicePk) GetEnableAhs() bool {
	if o == nil || isNil(o.EnableAhs) {
		var ret bool
		return ret
	}
	return *o.EnableAhs
}

// GetEnableAhsOk returns a tuple with the EnableAhs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGroupSerializerWithDevicePk) GetEnableAhsOk() (*bool, bool) {
	if o == nil || isNil(o.EnableAhs) {
    return nil, false
	}
	return o.EnableAhs, true
}

// HasEnableAhs returns a boolean if a field has been set.
func (o *DeviceGroupSerializerWithDevicePk) HasEnableAhs() bool {
	if o != nil && !isNil(o.EnableAhs) {
		return true
	}

	return false
}

// SetEnableAhs gets a reference to the given bool and assigns it to the EnableAhs field.
func (o *DeviceGroupSerializerWithDevicePk) SetEnableAhs(v bool) {
	o.EnableAhs = &v
}

// GetEnableAhsActions returns the EnableAhsActions field value if set, zero value otherwise.
func (o *DeviceGroupSerializerWithDevicePk) GetEnableAhsActions() bool {
	if o == nil || isNil(o.EnableAhsActions) {
		var ret bool
		return ret
	}
	return *o.EnableAhsActions
}

// GetEnableAhsActionsOk returns a tuple with the EnableAhsActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGroupSerializerWithDevicePk) GetEnableAhsActionsOk() (*bool, bool) {
	if o == nil || isNil(o.EnableAhsActions) {
    return nil, false
	}
	return o.EnableAhsActions, true
}

// HasEnableAhsActions returns a boolean if a field has been set.
func (o *DeviceGroupSerializerWithDevicePk) HasEnableAhsActions() bool {
	if o != nil && !isNil(o.EnableAhsActions) {
		return true
	}

	return false
}

// SetEnableAhsActions gets a reference to the given bool and assigns it to the EnableAhsActions field.
func (o *DeviceGroupSerializerWithDevicePk) SetEnableAhsActions(v bool) {
	o.EnableAhsActions = &v
}

// GetEnableAhsCas returns the EnableAhsCas field value if set, zero value otherwise.
func (o *DeviceGroupSerializerWithDevicePk) GetEnableAhsCas() bool {
	if o == nil || isNil(o.EnableAhsCas) {
		var ret bool
		return ret
	}
	return *o.EnableAhsCas
}

// GetEnableAhsCasOk returns a tuple with the EnableAhsCas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGroupSerializerWithDevicePk) GetEnableAhsCasOk() (*bool, bool) {
	if o == nil || isNil(o.EnableAhsCas) {
    return nil, false
	}
	return o.EnableAhsCas, true
}

// HasEnableAhsCas returns a boolean if a field has been set.
func (o *DeviceGroupSerializerWithDevicePk) HasEnableAhsCas() bool {
	if o != nil && !isNil(o.EnableAhsCas) {
		return true
	}

	return false
}

// SetEnableAhsCas gets a reference to the given bool and assigns it to the EnableAhsCas field.
func (o *DeviceGroupSerializerWithDevicePk) SetEnableAhsCas(v bool) {
	o.EnableAhsCas = &v
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

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DeviceGroupSerializerWithDevicePk) GetSource() SourceEnum {
	if o == nil || isNil(o.Source) {
		var ret SourceEnum
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGroupSerializerWithDevicePk) GetSourceOk() (*SourceEnum, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DeviceGroupSerializerWithDevicePk) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given SourceEnum and assigns it to the Source field.
func (o *DeviceGroupSerializerWithDevicePk) SetSource(v SourceEnum) {
	o.Source = &v
}

// GetIsReservedFilter returns the IsReservedFilter field value if set, zero value otherwise.
func (o *DeviceGroupSerializerWithDevicePk) GetIsReservedFilter() bool {
	if o == nil || isNil(o.IsReservedFilter) {
		var ret bool
		return ret
	}
	return *o.IsReservedFilter
}

// GetIsReservedFilterOk returns a tuple with the IsReservedFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGroupSerializerWithDevicePk) GetIsReservedFilterOk() (*bool, bool) {
	if o == nil || isNil(o.IsReservedFilter) {
    return nil, false
	}
	return o.IsReservedFilter, true
}

// HasIsReservedFilter returns a boolean if a field has been set.
func (o *DeviceGroupSerializerWithDevicePk) HasIsReservedFilter() bool {
	if o != nil && !isNil(o.IsReservedFilter) {
		return true
	}

	return false
}

// SetIsReservedFilter gets a reference to the given bool and assigns it to the IsReservedFilter field.
func (o *DeviceGroupSerializerWithDevicePk) SetIsReservedFilter(v bool) {
	o.IsReservedFilter = &v
}

func (o DeviceGroupSerializerWithDevicePk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.PermissionGroups != nil {
		toSerialize["permission_groups"] = o.PermissionGroups
	}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if true {
		toSerialize["online"] = o.Online
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["loc_desc"] = o.LocDesc
	}
	if true {
		toSerialize["is_reserved"] = o.IsReserved.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.EnableAhs) {
		toSerialize["enable_ahs"] = o.EnableAhs
	}
	if !isNil(o.EnableAhsActions) {
		toSerialize["enable_ahs_actions"] = o.EnableAhsActions
	}
	if !isNil(o.EnableAhsCas) {
		toSerialize["enable_ahs_cas"] = o.EnableAhsCas
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.IsReservedFilter) {
		toSerialize["is_reserved_filter"] = o.IsReservedFilter
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


