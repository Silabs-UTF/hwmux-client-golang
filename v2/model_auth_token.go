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

// checks if the AuthToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthToken{}

// AuthToken struct for AuthToken
type AuthToken struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token string `json:"token"`
}

type _AuthToken AuthToken

// NewAuthToken instantiates a new AuthToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthToken(username string, password string, token string) *AuthToken {
	this := AuthToken{}
	this.Username = username
	this.Password = password
	this.Token = token
	return &this
}

// NewAuthTokenWithDefaults instantiates a new AuthToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenWithDefaults() *AuthToken {
	this := AuthToken{}
	return &this
}

// GetUsername returns the Username field value
func (o *AuthToken) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthToken) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AuthToken) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *AuthToken) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AuthToken) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AuthToken) SetPassword(v string) {
	o.Password = v
}

// GetToken returns the Token field value
func (o *AuthToken) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AuthToken) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AuthToken) SetToken(v string) {
	o.Token = v
}

func (o AuthToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *AuthToken) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
		"password",
		"token",
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

	varAuthToken := _AuthToken{}

	err = json.Unmarshal(bytes, &varAuthToken)

	if err != nil {
		return err
	}

	*o = AuthToken(varAuthToken)

	return err
}

type NullableAuthToken struct {
	value *AuthToken
	isSet bool
}

func (v NullableAuthToken) Get() *AuthToken {
	return v.value
}

func (v *NullableAuthToken) Set(val *AuthToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthToken(val *AuthToken) *NullableAuthToken {
	return &NullableAuthToken{value: val, isSet: true}
}

func (v NullableAuthToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


