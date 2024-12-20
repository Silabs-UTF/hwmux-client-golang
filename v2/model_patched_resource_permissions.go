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

// PatchedResourcePermissions struct for PatchedResourcePermissions
type PatchedResourcePermissions struct {
	Id *int32 `json:"id,omitempty"`
	Permissions []PermissionsEnum `json:"permissions,omitempty"`
}

// NewPatchedResourcePermissions instantiates a new PatchedResourcePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedResourcePermissions() *PatchedResourcePermissions {
	this := PatchedResourcePermissions{}
	return &this
}

// NewPatchedResourcePermissionsWithDefaults instantiates a new PatchedResourcePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedResourcePermissionsWithDefaults() *PatchedResourcePermissions {
	this := PatchedResourcePermissions{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedResourcePermissions) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedResourcePermissions) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedResourcePermissions) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PatchedResourcePermissions) SetId(v int32) {
	o.Id = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *PatchedResourcePermissions) GetPermissions() []PermissionsEnum {
	if o == nil || isNil(o.Permissions) {
		var ret []PermissionsEnum
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedResourcePermissions) GetPermissionsOk() ([]PermissionsEnum, bool) {
	if o == nil || isNil(o.Permissions) {
    return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PatchedResourcePermissions) HasPermissions() bool {
	if o != nil && !isNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []PermissionsEnum and assigns it to the Permissions field.
func (o *PatchedResourcePermissions) SetPermissions(v []PermissionsEnum) {
	o.Permissions = v
}

func (o PatchedResourcePermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedResourcePermissions struct {
	value *PatchedResourcePermissions
	isSet bool
}

func (v NullablePatchedResourcePermissions) Get() *PatchedResourcePermissions {
	return v.value
}

func (v *NullablePatchedResourcePermissions) Set(val *PatchedResourcePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedResourcePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedResourcePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedResourcePermissions(val *PatchedResourcePermissions) *NullablePatchedResourcePermissions {
	return &NullablePatchedResourcePermissions{value: val, isSet: true}
}

func (v NullablePatchedResourcePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedResourcePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


