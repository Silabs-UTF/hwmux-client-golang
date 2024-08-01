/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
)

// PatchedPartFamily Serializer for the PartFamily model
type PatchedPartFamily struct {
	Name *string `json:"name,omitempty"`
	RegexPattern NullableString `json:"regex_pattern,omitempty"`
	Description *string `json:"description,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPatchedPartFamily instantiates a new PatchedPartFamily object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPartFamily() *PatchedPartFamily {
	this := PatchedPartFamily{}
	return &this
}

// NewPatchedPartFamilyWithDefaults instantiates a new PatchedPartFamily object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPartFamilyWithDefaults() *PatchedPartFamily {
	this := PatchedPartFamily{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedPartFamily) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPartFamily) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedPartFamily) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedPartFamily) SetName(v string) {
	o.Name = &v
}

// GetRegexPattern returns the RegexPattern field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPartFamily) GetRegexPattern() string {
	if o == nil || isNil(o.RegexPattern.Get()) {
		var ret string
		return ret
	}
	return *o.RegexPattern.Get()
}

// GetRegexPatternOk returns a tuple with the RegexPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPartFamily) GetRegexPatternOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RegexPattern.Get(), o.RegexPattern.IsSet()
}

// HasRegexPattern returns a boolean if a field has been set.
func (o *PatchedPartFamily) HasRegexPattern() bool {
	if o != nil && o.RegexPattern.IsSet() {
		return true
	}

	return false
}

// SetRegexPattern gets a reference to the given NullableString and assigns it to the RegexPattern field.
func (o *PatchedPartFamily) SetRegexPattern(v string) {
	o.RegexPattern.Set(&v)
}
// SetRegexPatternNil sets the value for RegexPattern to be an explicit nil
func (o *PatchedPartFamily) SetRegexPatternNil() {
	o.RegexPattern.Set(nil)
}

// UnsetRegexPattern ensures that no value is present for RegexPattern, not even an explicit nil
func (o *PatchedPartFamily) UnsetRegexPattern() {
	o.RegexPattern.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedPartFamily) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPartFamily) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedPartFamily) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedPartFamily) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PatchedPartFamily) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPartFamily) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchedPartFamily) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchedPartFamily) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PatchedPartFamily) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.RegexPattern.IsSet() {
		toSerialize["regex_pattern"] = o.RegexPattern.Get()
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedPartFamily struct {
	value *PatchedPartFamily
	isSet bool
}

func (v NullablePatchedPartFamily) Get() *PatchedPartFamily {
	return v.value
}

func (v *NullablePatchedPartFamily) Set(val *PatchedPartFamily) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPartFamily) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPartFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPartFamily(val *PatchedPartFamily) *NullablePatchedPartFamily {
	return &NullablePatchedPartFamily{value: val, isSet: true}
}

func (v NullablePatchedPartFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPartFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


