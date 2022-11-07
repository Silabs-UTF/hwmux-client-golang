/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LocationSerializerWriteOnly Serializer for the Location model
type LocationSerializerWriteOnly struct {
	Id int32 `json:"id"`
	Room string `json:"room"`
	Description *string `json:"description,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewLocationSerializerWriteOnly instantiates a new LocationSerializerWriteOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationSerializerWriteOnly(id int32, room string) *LocationSerializerWriteOnly {
	this := LocationSerializerWriteOnly{}
	this.Id = id
	this.Room = room
	return &this
}

// NewLocationSerializerWriteOnlyWithDefaults instantiates a new LocationSerializerWriteOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationSerializerWriteOnlyWithDefaults() *LocationSerializerWriteOnly {
	this := LocationSerializerWriteOnly{}
	return &this
}

// GetId returns the Id field value
func (o *LocationSerializerWriteOnly) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LocationSerializerWriteOnly) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LocationSerializerWriteOnly) SetId(v int32) {
	o.Id = v
}

// GetRoom returns the Room field value
func (o *LocationSerializerWriteOnly) GetRoom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Room
}

// GetRoomOk returns a tuple with the Room field value
// and a boolean to check if the value has been set.
func (o *LocationSerializerWriteOnly) GetRoomOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Room, true
}

// SetRoom sets field value
func (o *LocationSerializerWriteOnly) SetRoom(v string) {
	o.Room = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LocationSerializerWriteOnly) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationSerializerWriteOnly) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LocationSerializerWriteOnly) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LocationSerializerWriteOnly) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LocationSerializerWriteOnly) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationSerializerWriteOnly) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LocationSerializerWriteOnly) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LocationSerializerWriteOnly) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o LocationSerializerWriteOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["room"] = o.Room
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableLocationSerializerWriteOnly struct {
	value *LocationSerializerWriteOnly
	isSet bool
}

func (v NullableLocationSerializerWriteOnly) Get() *LocationSerializerWriteOnly {
	return v.value
}

func (v *NullableLocationSerializerWriteOnly) Set(val *LocationSerializerWriteOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationSerializerWriteOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationSerializerWriteOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationSerializerWriteOnly(val *LocationSerializerWriteOnly) *NullableLocationSerializerWriteOnly {
	return &NullableLocationSerializerWriteOnly{value: val, isSet: true}
}

func (v NullableLocationSerializerWriteOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationSerializerWriteOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


