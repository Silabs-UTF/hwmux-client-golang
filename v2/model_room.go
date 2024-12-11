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

// Room Serializer for the Room model
type Room struct {
	Name string `json:"name"`
	Site string `json:"site"`
	Description *string `json:"description,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewRoom instantiates a new Room object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoom(name string, site string) *Room {
	this := Room{}
	this.Name = name
	this.Site = site
	return &this
}

// NewRoomWithDefaults instantiates a new Room object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomWithDefaults() *Room {
	this := Room{}
	return &this
}

// GetName returns the Name field value
func (o *Room) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Room) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Room) SetName(v string) {
	o.Name = v
}

// GetSite returns the Site field value
func (o *Room) GetSite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *Room) GetSiteOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Site, true
}

// SetSite sets field value
func (o *Room) SetSite(v string) {
	o.Site = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Room) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Room) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Room) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Room) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Room) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Room) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Room) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Room) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o Room) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["site"] = o.Site
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableRoom struct {
	value *Room
	isSet bool
}

func (v NullableRoom) Get() *Room {
	return v.value
}

func (v *NullableRoom) Set(val *Room) {
	v.value = val
	v.isSet = true
}

func (v NullableRoom) IsSet() bool {
	return v.isSet
}

func (v *NullableRoom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoom(val *Room) *NullableRoom {
	return &NullableRoom{value: val, isSet: true}
}

func (v NullableRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


