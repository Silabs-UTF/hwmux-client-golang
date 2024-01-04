/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.26.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
	"fmt"
)

// checks if the ObjectPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectPermissions{}

// ObjectPermissions struct for ObjectPermissions
type ObjectPermissions struct {
	Users map[string]interface{} `json:"users"`
	UserGroups map[string]interface{} `json:"user_groups"`
}

type _ObjectPermissions ObjectPermissions

// NewObjectPermissions instantiates a new ObjectPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectPermissions(users map[string]interface{}, userGroups map[string]interface{}) *ObjectPermissions {
	this := ObjectPermissions{}
	this.Users = users
	this.UserGroups = userGroups
	return &this
}

// NewObjectPermissionsWithDefaults instantiates a new ObjectPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectPermissionsWithDefaults() *ObjectPermissions {
	this := ObjectPermissions{}
	return &this
}

// GetUsers returns the Users field value
func (o *ObjectPermissions) GetUsers() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ObjectPermissions) GetUsersOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ObjectPermissions) SetUsers(v map[string]interface{}) {
	o.Users = v
}

// GetUserGroups returns the UserGroups field value
func (o *ObjectPermissions) GetUserGroups() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserGroups
}

// GetUserGroupsOk returns a tuple with the UserGroups field value
// and a boolean to check if the value has been set.
func (o *ObjectPermissions) GetUserGroupsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.UserGroups, true
}

// SetUserGroups sets field value
func (o *ObjectPermissions) SetUserGroups(v map[string]interface{}) {
	o.UserGroups = v
}

func (o ObjectPermissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	toSerialize["user_groups"] = o.UserGroups
	return toSerialize, nil
}

func (o *ObjectPermissions) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
		"user_groups",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varObjectPermissions := _ObjectPermissions{}

	err = json.Unmarshal(bytes, &varObjectPermissions)

	if err != nil {
		return err
	}

	*o = ObjectPermissions(varObjectPermissions)

	return err
}

type NullableObjectPermissions struct {
	value *ObjectPermissions
	isSet bool
}

func (v NullableObjectPermissions) Get() *ObjectPermissions {
	return v.value
}

func (v *NullableObjectPermissions) Set(val *ObjectPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectPermissions(val *ObjectPermissions) *NullableObjectPermissions {
	return &NullableObjectPermissions{value: val, isSet: true}
}

func (v NullableObjectPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


