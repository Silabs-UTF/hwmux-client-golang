/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
)

// PermissionGroup struct for PermissionGroup
type PermissionGroup struct {
	Id int32 `json:"id"`
	Permissions []string `json:"permissions"`
	Name string `json:"name"`
}

// NewPermissionGroup instantiates a new PermissionGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionGroup(id int32, permissions []string, name string) *PermissionGroup {
	this := PermissionGroup{}
	this.Id = id
	this.Permissions = permissions
	this.Name = name
	return &this
}

// NewPermissionGroupWithDefaults instantiates a new PermissionGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionGroupWithDefaults() *PermissionGroup {
	this := PermissionGroup{}
	return &this
}

// GetId returns the Id field value
func (o *PermissionGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PermissionGroup) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PermissionGroup) SetId(v int32) {
	o.Id = v
}

// GetPermissions returns the Permissions field value
func (o *PermissionGroup) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *PermissionGroup) GetPermissionsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *PermissionGroup) SetPermissions(v []string) {
	o.Permissions = v
}

// GetName returns the Name field value
func (o *PermissionGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PermissionGroup) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PermissionGroup) SetName(v string) {
	o.Name = v
}

func (o PermissionGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionGroup struct {
	value *PermissionGroup
	isSet bool
}

func (v NullablePermissionGroup) Get() *PermissionGroup {
	return v.value
}

func (v *NullablePermissionGroup) Set(val *PermissionGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionGroup(val *PermissionGroup) *NullablePermissionGroup {
	return &NullablePermissionGroup{value: val, isSet: true}
}

func (v NullablePermissionGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


