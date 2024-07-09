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

// User Used for showing username of the person currently owning the device instead of showing user id.
type User struct {
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(username string) *User {
	this := User{}
	this.Username = username
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetUsername returns the Username field value
func (o *User) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *User) SetUsername(v string) {
	o.Username = v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


