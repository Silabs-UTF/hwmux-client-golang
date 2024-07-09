/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
)

// LabelSerializerWithPermissions struct for LabelSerializerWithPermissions
type LabelSerializerWithPermissions struct {
	Id int32 `json:"id"`
	PermissionGroups []string `json:"permission_groups,omitempty"`
	DeviceGroups []int32 `json:"device_groups"`
	Name string `json:"name"`
	Metadata interface{} `json:"metadata,omitempty"`
	Source *DeviceGroupSource `json:"source,omitempty"`
}

// NewLabelSerializerWithPermissions instantiates a new LabelSerializerWithPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelSerializerWithPermissions(id int32, deviceGroups []int32, name string) *LabelSerializerWithPermissions {
	this := LabelSerializerWithPermissions{}
	this.Id = id
	this.DeviceGroups = deviceGroups
	this.Name = name
	return &this
}

// NewLabelSerializerWithPermissionsWithDefaults instantiates a new LabelSerializerWithPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelSerializerWithPermissionsWithDefaults() *LabelSerializerWithPermissions {
	this := LabelSerializerWithPermissions{}
	return &this
}

// GetId returns the Id field value
func (o *LabelSerializerWithPermissions) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LabelSerializerWithPermissions) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LabelSerializerWithPermissions) SetId(v int32) {
	o.Id = v
}

// GetPermissionGroups returns the PermissionGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LabelSerializerWithPermissions) GetPermissionGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PermissionGroups
}

// GetPermissionGroupsOk returns a tuple with the PermissionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabelSerializerWithPermissions) GetPermissionGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.PermissionGroups) {
    return nil, false
	}
	return o.PermissionGroups, true
}

// HasPermissionGroups returns a boolean if a field has been set.
func (o *LabelSerializerWithPermissions) HasPermissionGroups() bool {
	if o != nil && isNil(o.PermissionGroups) {
		return true
	}

	return false
}

// SetPermissionGroups gets a reference to the given []string and assigns it to the PermissionGroups field.
func (o *LabelSerializerWithPermissions) SetPermissionGroups(v []string) {
	o.PermissionGroups = v
}

// GetDeviceGroups returns the DeviceGroups field value
func (o *LabelSerializerWithPermissions) GetDeviceGroups() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.DeviceGroups
}

// GetDeviceGroupsOk returns a tuple with the DeviceGroups field value
// and a boolean to check if the value has been set.
func (o *LabelSerializerWithPermissions) GetDeviceGroupsOk() ([]int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.DeviceGroups, true
}

// SetDeviceGroups sets field value
func (o *LabelSerializerWithPermissions) SetDeviceGroups(v []int32) {
	o.DeviceGroups = v
}

// GetName returns the Name field value
func (o *LabelSerializerWithPermissions) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LabelSerializerWithPermissions) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LabelSerializerWithPermissions) SetName(v string) {
	o.Name = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LabelSerializerWithPermissions) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabelSerializerWithPermissions) GetMetadataOk() (*interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LabelSerializerWithPermissions) HasMetadata() bool {
	if o != nil && isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *LabelSerializerWithPermissions) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *LabelSerializerWithPermissions) GetSource() DeviceGroupSource {
	if o == nil || isNil(o.Source) {
		var ret DeviceGroupSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelSerializerWithPermissions) GetSourceOk() (*DeviceGroupSource, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *LabelSerializerWithPermissions) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given DeviceGroupSource and assigns it to the Source field.
func (o *LabelSerializerWithPermissions) SetSource(v DeviceGroupSource) {
	o.Source = &v
}

func (o LabelSerializerWithPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.PermissionGroups != nil {
		toSerialize["permission_groups"] = o.PermissionGroups
	}
	if true {
		toSerialize["device_groups"] = o.DeviceGroups
	}
	if true {
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

type NullableLabelSerializerWithPermissions struct {
	value *LabelSerializerWithPermissions
	isSet bool
}

func (v NullableLabelSerializerWithPermissions) Get() *LabelSerializerWithPermissions {
	return v.value
}

func (v *NullableLabelSerializerWithPermissions) Set(val *LabelSerializerWithPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelSerializerWithPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelSerializerWithPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelSerializerWithPermissions(val *LabelSerializerWithPermissions) *NullableLabelSerializerWithPermissions {
	return &NullableLabelSerializerWithPermissions{value: val, isSet: true}
}

func (v NullableLabelSerializerWithPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelSerializerWithPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


