/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
)

// PatchedObjectPermissions struct for PatchedObjectPermissions
type PatchedObjectPermissions struct {
	Users map[string]interface{} `json:"users,omitempty"`
	UserGroups map[string]interface{} `json:"user_groups,omitempty"`
}

// NewPatchedObjectPermissions instantiates a new PatchedObjectPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedObjectPermissions() *PatchedObjectPermissions {
	this := PatchedObjectPermissions{}
	return &this
}

// NewPatchedObjectPermissionsWithDefaults instantiates a new PatchedObjectPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedObjectPermissionsWithDefaults() *PatchedObjectPermissions {
	this := PatchedObjectPermissions{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *PatchedObjectPermissions) GetUsers() map[string]interface{} {
	if o == nil || isNil(o.Users) {
		var ret map[string]interface{}
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedObjectPermissions) GetUsersOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Users) {
    return map[string]interface{}{}, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *PatchedObjectPermissions) HasUsers() bool {
	if o != nil && !isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given map[string]interface{} and assigns it to the Users field.
func (o *PatchedObjectPermissions) SetUsers(v map[string]interface{}) {
	o.Users = v
}

// GetUserGroups returns the UserGroups field value if set, zero value otherwise.
func (o *PatchedObjectPermissions) GetUserGroups() map[string]interface{} {
	if o == nil || isNil(o.UserGroups) {
		var ret map[string]interface{}
		return ret
	}
	return o.UserGroups
}

// GetUserGroupsOk returns a tuple with the UserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedObjectPermissions) GetUserGroupsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.UserGroups) {
    return map[string]interface{}{}, false
	}
	return o.UserGroups, true
}

// HasUserGroups returns a boolean if a field has been set.
func (o *PatchedObjectPermissions) HasUserGroups() bool {
	if o != nil && !isNil(o.UserGroups) {
		return true
	}

	return false
}

// SetUserGroups gets a reference to the given map[string]interface{} and assigns it to the UserGroups field.
func (o *PatchedObjectPermissions) SetUserGroups(v map[string]interface{}) {
	o.UserGroups = v
}

func (o PatchedObjectPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !isNil(o.UserGroups) {
		toSerialize["user_groups"] = o.UserGroups
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedObjectPermissions struct {
	value *PatchedObjectPermissions
	isSet bool
}

func (v NullablePatchedObjectPermissions) Get() *PatchedObjectPermissions {
	return v.value
}

func (v *NullablePatchedObjectPermissions) Set(val *PatchedObjectPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedObjectPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedObjectPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedObjectPermissions(val *PatchedObjectPermissions) *NullablePatchedObjectPermissions {
	return &NullablePatchedObjectPermissions{value: val, isSet: true}
}

func (v NullablePatchedObjectPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedObjectPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


