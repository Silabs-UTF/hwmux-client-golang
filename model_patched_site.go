/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
)

// PatchedSite Serializer for the Site model
type PatchedSite struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Lon *float32 `json:"lon,omitempty"`
	Lat *float32 `json:"lat,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPatchedSite instantiates a new PatchedSite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSite() *PatchedSite {
	this := PatchedSite{}
	return &this
}

// NewPatchedSiteWithDefaults instantiates a new PatchedSite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSiteWithDefaults() *PatchedSite {
	this := PatchedSite{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedSite) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSite) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedSite) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedSite) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedSite) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSite) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedSite) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedSite) SetDescription(v string) {
	o.Description = &v
}

// GetLon returns the Lon field value if set, zero value otherwise.
func (o *PatchedSite) GetLon() float32 {
	if o == nil || isNil(o.Lon) {
		var ret float32
		return ret
	}
	return *o.Lon
}

// GetLonOk returns a tuple with the Lon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSite) GetLonOk() (*float32, bool) {
	if o == nil || isNil(o.Lon) {
    return nil, false
	}
	return o.Lon, true
}

// HasLon returns a boolean if a field has been set.
func (o *PatchedSite) HasLon() bool {
	if o != nil && !isNil(o.Lon) {
		return true
	}

	return false
}

// SetLon gets a reference to the given float32 and assigns it to the Lon field.
func (o *PatchedSite) SetLon(v float32) {
	o.Lon = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *PatchedSite) GetLat() float32 {
	if o == nil || isNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSite) GetLatOk() (*float32, bool) {
	if o == nil || isNil(o.Lat) {
    return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *PatchedSite) HasLat() bool {
	if o != nil && !isNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *PatchedSite) SetLat(v float32) {
	o.Lat = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PatchedSite) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSite) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchedSite) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchedSite) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PatchedSite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Lon) {
		toSerialize["lon"] = o.Lon
	}
	if !isNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedSite struct {
	value *PatchedSite
	isSet bool
}

func (v NullablePatchedSite) Get() *PatchedSite {
	return v.value
}

func (v *NullablePatchedSite) Set(val *PatchedSite) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSite) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSite(val *PatchedSite) *NullablePatchedSite {
	return &NullablePatchedSite{value: val, isSet: true}
}

func (v NullablePatchedSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


