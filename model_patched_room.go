/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.19.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
)

// PatchedRoom Serializer for the Room model
type PatchedRoom struct {
	Name *string `json:"name,omitempty"`
	Site *string `json:"site,omitempty"`
	Description *string `json:"description,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPatchedRoom instantiates a new PatchedRoom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedRoom() *PatchedRoom {
	this := PatchedRoom{}
	return &this
}

// NewPatchedRoomWithDefaults instantiates a new PatchedRoom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedRoomWithDefaults() *PatchedRoom {
	this := PatchedRoom{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedRoom) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoom) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedRoom) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedRoom) SetName(v string) {
	o.Name = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *PatchedRoom) GetSite() string {
	if o == nil || isNil(o.Site) {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoom) GetSiteOk() (*string, bool) {
	if o == nil || isNil(o.Site) {
    return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *PatchedRoom) HasSite() bool {
	if o != nil && !isNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *PatchedRoom) SetSite(v string) {
	o.Site = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedRoom) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoom) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedRoom) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedRoom) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PatchedRoom) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoom) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchedRoom) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchedRoom) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PatchedRoom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Site) {
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

type NullablePatchedRoom struct {
	value *PatchedRoom
	isSet bool
}

func (v NullablePatchedRoom) Get() *PatchedRoom {
	return v.value
}

func (v *NullablePatchedRoom) Set(val *PatchedRoom) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedRoom) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedRoom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedRoom(val *PatchedRoom) *NullablePatchedRoom {
	return &NullablePatchedRoom{value: val, isSet: true}
}

func (v NullablePatchedRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedRoom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


