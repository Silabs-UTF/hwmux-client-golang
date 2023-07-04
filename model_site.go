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

// Site Serializer for the Site model
type Site struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewSite instantiates a new Site object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSite(name string, lon float32, lat float32) *Site {
	this := Site{}
	this.Name = name
	this.Lon = lon
	this.Lat = lat
	return &this
}

// NewSiteWithDefaults instantiates a new Site object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteWithDefaults() *Site {
	this := Site{}
	return &this
}

// GetName returns the Name field value
func (o *Site) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Site) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Site) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Site) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Site) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Site) SetDescription(v string) {
	o.Description = &v
}

// GetLon returns the Lon field value
func (o *Site) GetLon() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Lon
}

// GetLonOk returns a tuple with the Lon field value
// and a boolean to check if the value has been set.
func (o *Site) GetLonOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Lon, true
}

// SetLon sets field value
func (o *Site) SetLon(v float32) {
	o.Lon = v
}

// GetLat returns the Lat field value
func (o *Site) GetLat() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Lat
}

// GetLatOk returns a tuple with the Lat field value
// and a boolean to check if the value has been set.
func (o *Site) GetLatOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Lat, true
}

// SetLat sets field value
func (o *Site) SetLat(v float32) {
	o.Lat = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Site) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Site) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Site) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o Site) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["lon"] = o.Lon
	}
	if true {
		toSerialize["lat"] = o.Lat
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableSite struct {
	value *Site
	isSet bool
}

func (v NullableSite) Get() *Site {
	return v.value
}

func (v *NullableSite) Set(val *Site) {
	v.value = val
	v.isSet = true
}

func (v NullableSite) IsSet() bool {
	return v.isSet
}

func (v *NullableSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSite(val *Site) *NullableSite {
	return &NullableSite{value: val, isSet: true}
}

func (v NullableSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


