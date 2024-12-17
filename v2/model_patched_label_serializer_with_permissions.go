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

// PatchedLabelSerializerWithPermissions struct for PatchedLabelSerializerWithPermissions
type PatchedLabelSerializerWithPermissions struct {
	Id *int32 `json:"id,omitempty"`
	PermissionGroups []string `json:"permission_groups,omitempty"`
	DeviceGroups []int32 `json:"device_groups,omitempty"`
	Name *string `json:"name,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Source *SourceEnum `json:"source,omitempty"`
}

// NewPatchedLabelSerializerWithPermissions instantiates a new PatchedLabelSerializerWithPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedLabelSerializerWithPermissions() *PatchedLabelSerializerWithPermissions {
	this := PatchedLabelSerializerWithPermissions{}
	return &this
}

// NewPatchedLabelSerializerWithPermissionsWithDefaults instantiates a new PatchedLabelSerializerWithPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedLabelSerializerWithPermissionsWithDefaults() *PatchedLabelSerializerWithPermissions {
	this := PatchedLabelSerializerWithPermissions{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedLabelSerializerWithPermissions) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLabelSerializerWithPermissions) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedLabelSerializerWithPermissions) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PatchedLabelSerializerWithPermissions) SetId(v int32) {
	o.Id = &v
}

// GetPermissionGroups returns the PermissionGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedLabelSerializerWithPermissions) GetPermissionGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PermissionGroups
}

// GetPermissionGroupsOk returns a tuple with the PermissionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedLabelSerializerWithPermissions) GetPermissionGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.PermissionGroups) {
    return nil, false
	}
	return o.PermissionGroups, true
}

// HasPermissionGroups returns a boolean if a field has been set.
func (o *PatchedLabelSerializerWithPermissions) HasPermissionGroups() bool {
	if o != nil && isNil(o.PermissionGroups) {
		return true
	}

	return false
}

// SetPermissionGroups gets a reference to the given []string and assigns it to the PermissionGroups field.
func (o *PatchedLabelSerializerWithPermissions) SetPermissionGroups(v []string) {
	o.PermissionGroups = v
}

// GetDeviceGroups returns the DeviceGroups field value if set, zero value otherwise.
func (o *PatchedLabelSerializerWithPermissions) GetDeviceGroups() []int32 {
	if o == nil || isNil(o.DeviceGroups) {
		var ret []int32
		return ret
	}
	return o.DeviceGroups
}

// GetDeviceGroupsOk returns a tuple with the DeviceGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLabelSerializerWithPermissions) GetDeviceGroupsOk() ([]int32, bool) {
	if o == nil || isNil(o.DeviceGroups) {
    return nil, false
	}
	return o.DeviceGroups, true
}

// HasDeviceGroups returns a boolean if a field has been set.
func (o *PatchedLabelSerializerWithPermissions) HasDeviceGroups() bool {
	if o != nil && !isNil(o.DeviceGroups) {
		return true
	}

	return false
}

// SetDeviceGroups gets a reference to the given []int32 and assigns it to the DeviceGroups field.
func (o *PatchedLabelSerializerWithPermissions) SetDeviceGroups(v []int32) {
	o.DeviceGroups = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedLabelSerializerWithPermissions) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLabelSerializerWithPermissions) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedLabelSerializerWithPermissions) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedLabelSerializerWithPermissions) SetName(v string) {
	o.Name = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedLabelSerializerWithPermissions) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedLabelSerializerWithPermissions) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchedLabelSerializerWithPermissions) HasMetadata() bool {
	if o != nil && isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchedLabelSerializerWithPermissions) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PatchedLabelSerializerWithPermissions) GetSource() SourceEnum {
	if o == nil || isNil(o.Source) {
		var ret SourceEnum
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLabelSerializerWithPermissions) GetSourceOk() (*SourceEnum, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PatchedLabelSerializerWithPermissions) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given SourceEnum and assigns it to the Source field.
func (o *PatchedLabelSerializerWithPermissions) SetSource(v SourceEnum) {
	o.Source = &v
}

func (o PatchedLabelSerializerWithPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.PermissionGroups != nil {
		toSerialize["permission_groups"] = o.PermissionGroups
	}
	if !isNil(o.DeviceGroups) {
		toSerialize["device_groups"] = o.DeviceGroups
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedLabelSerializerWithPermissions struct {
	value *PatchedLabelSerializerWithPermissions
	isSet bool
}

func (v NullablePatchedLabelSerializerWithPermissions) Get() *PatchedLabelSerializerWithPermissions {
	return v.value
}

func (v *NullablePatchedLabelSerializerWithPermissions) Set(val *PatchedLabelSerializerWithPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedLabelSerializerWithPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedLabelSerializerWithPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedLabelSerializerWithPermissions(val *PatchedLabelSerializerWithPermissions) *NullablePatchedLabelSerializerWithPermissions {
	return &NullablePatchedLabelSerializerWithPermissions{value: val, isSet: true}
}

func (v NullablePatchedLabelSerializerWithPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedLabelSerializerWithPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


