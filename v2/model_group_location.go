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

// GroupLocation Serializer for the GroupLocation model
type GroupLocation struct {
	Id int32 `json:"id"`
	Rooms []Room `json:"rooms"`
	Description *string `json:"description,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	DeviceGroup int32 `json:"device_group"`
}

// NewGroupLocation instantiates a new GroupLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupLocation(id int32, rooms []Room, deviceGroup int32) *GroupLocation {
	this := GroupLocation{}
	this.Id = id
	this.Rooms = rooms
	this.DeviceGroup = deviceGroup
	return &this
}

// NewGroupLocationWithDefaults instantiates a new GroupLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupLocationWithDefaults() *GroupLocation {
	this := GroupLocation{}
	return &this
}

// GetId returns the Id field value
func (o *GroupLocation) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupLocation) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupLocation) SetId(v int32) {
	o.Id = v
}

// GetRooms returns the Rooms field value
func (o *GroupLocation) GetRooms() []Room {
	if o == nil {
		var ret []Room
		return ret
	}

	return o.Rooms
}

// GetRoomsOk returns a tuple with the Rooms field value
// and a boolean to check if the value has been set.
func (o *GroupLocation) GetRoomsOk() ([]Room, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rooms, true
}

// SetRooms sets field value
func (o *GroupLocation) SetRooms(v []Room) {
	o.Rooms = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupLocation) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLocation) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupLocation) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupLocation) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GroupLocation) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLocation) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GroupLocation) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GroupLocation) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetDeviceGroup returns the DeviceGroup field value
func (o *GroupLocation) GetDeviceGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceGroup
}

// GetDeviceGroupOk returns a tuple with the DeviceGroup field value
// and a boolean to check if the value has been set.
func (o *GroupLocation) GetDeviceGroupOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DeviceGroup, true
}

// SetDeviceGroup sets field value
func (o *GroupLocation) SetDeviceGroup(v int32) {
	o.DeviceGroup = v
}

func (o GroupLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["rooms"] = o.Rooms
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["device_group"] = o.DeviceGroup
	}
	return json.Marshal(toSerialize)
}

type NullableGroupLocation struct {
	value *GroupLocation
	isSet bool
}

func (v NullableGroupLocation) Get() *GroupLocation {
	return v.value
}

func (v *NullableGroupLocation) Set(val *GroupLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupLocation(val *GroupLocation) *NullableGroupLocation {
	return &NullableGroupLocation{value: val, isSet: true}
}

func (v NullableGroupLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


