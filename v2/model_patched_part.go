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

// PatchedPart Serializer for the Part model
type PatchedPart struct {
	PartNo *string `json:"part_no,omitempty"`
	PartFamily *PartFamily `json:"part_family,omitempty"`
	BoardNo NullableString `json:"board_no,omitempty"`
	Variant *string `json:"variant,omitempty"`
	Revision *string `json:"revision,omitempty"`
	ChipNo NullableString `json:"chip_no,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPatchedPart instantiates a new PatchedPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPart() *PatchedPart {
	this := PatchedPart{}
	var variant string = ""
	this.Variant = &variant
	var revision string = ""
	this.Revision = &revision
	return &this
}

// NewPatchedPartWithDefaults instantiates a new PatchedPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPartWithDefaults() *PatchedPart {
	this := PatchedPart{}
	var variant string = ""
	this.Variant = &variant
	var revision string = ""
	this.Revision = &revision
	return &this
}

// GetPartNo returns the PartNo field value if set, zero value otherwise.
func (o *PatchedPart) GetPartNo() string {
	if o == nil || isNil(o.PartNo) {
		var ret string
		return ret
	}
	return *o.PartNo
}

// GetPartNoOk returns a tuple with the PartNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPart) GetPartNoOk() (*string, bool) {
	if o == nil || isNil(o.PartNo) {
    return nil, false
	}
	return o.PartNo, true
}

// HasPartNo returns a boolean if a field has been set.
func (o *PatchedPart) HasPartNo() bool {
	if o != nil && !isNil(o.PartNo) {
		return true
	}

	return false
}

// SetPartNo gets a reference to the given string and assigns it to the PartNo field.
func (o *PatchedPart) SetPartNo(v string) {
	o.PartNo = &v
}

// GetPartFamily returns the PartFamily field value if set, zero value otherwise.
func (o *PatchedPart) GetPartFamily() PartFamily {
	if o == nil || isNil(o.PartFamily) {
		var ret PartFamily
		return ret
	}
	return *o.PartFamily
}

// GetPartFamilyOk returns a tuple with the PartFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPart) GetPartFamilyOk() (*PartFamily, bool) {
	if o == nil || isNil(o.PartFamily) {
    return nil, false
	}
	return o.PartFamily, true
}

// HasPartFamily returns a boolean if a field has been set.
func (o *PatchedPart) HasPartFamily() bool {
	if o != nil && !isNil(o.PartFamily) {
		return true
	}

	return false
}

// SetPartFamily gets a reference to the given PartFamily and assigns it to the PartFamily field.
func (o *PatchedPart) SetPartFamily(v PartFamily) {
	o.PartFamily = &v
}

// GetBoardNo returns the BoardNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPart) GetBoardNo() string {
	if o == nil || isNil(o.BoardNo.Get()) {
		var ret string
		return ret
	}
	return *o.BoardNo.Get()
}

// GetBoardNoOk returns a tuple with the BoardNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPart) GetBoardNoOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BoardNo.Get(), o.BoardNo.IsSet()
}

// HasBoardNo returns a boolean if a field has been set.
func (o *PatchedPart) HasBoardNo() bool {
	if o != nil && o.BoardNo.IsSet() {
		return true
	}

	return false
}

// SetBoardNo gets a reference to the given NullableString and assigns it to the BoardNo field.
func (o *PatchedPart) SetBoardNo(v string) {
	o.BoardNo.Set(&v)
}
// SetBoardNoNil sets the value for BoardNo to be an explicit nil
func (o *PatchedPart) SetBoardNoNil() {
	o.BoardNo.Set(nil)
}

// UnsetBoardNo ensures that no value is present for BoardNo, not even an explicit nil
func (o *PatchedPart) UnsetBoardNo() {
	o.BoardNo.Unset()
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *PatchedPart) GetVariant() string {
	if o == nil || isNil(o.Variant) {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPart) GetVariantOk() (*string, bool) {
	if o == nil || isNil(o.Variant) {
    return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *PatchedPart) HasVariant() bool {
	if o != nil && !isNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *PatchedPart) SetVariant(v string) {
	o.Variant = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *PatchedPart) GetRevision() string {
	if o == nil || isNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPart) GetRevisionOk() (*string, bool) {
	if o == nil || isNil(o.Revision) {
    return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *PatchedPart) HasRevision() bool {
	if o != nil && !isNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *PatchedPart) SetRevision(v string) {
	o.Revision = &v
}

// GetChipNo returns the ChipNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPart) GetChipNo() string {
	if o == nil || isNil(o.ChipNo.Get()) {
		var ret string
		return ret
	}
	return *o.ChipNo.Get()
}

// GetChipNoOk returns a tuple with the ChipNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPart) GetChipNoOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChipNo.Get(), o.ChipNo.IsSet()
}

// HasChipNo returns a boolean if a field has been set.
func (o *PatchedPart) HasChipNo() bool {
	if o != nil && o.ChipNo.IsSet() {
		return true
	}

	return false
}

// SetChipNo gets a reference to the given NullableString and assigns it to the ChipNo field.
func (o *PatchedPart) SetChipNo(v string) {
	o.ChipNo.Set(&v)
}
// SetChipNoNil sets the value for ChipNo to be an explicit nil
func (o *PatchedPart) SetChipNoNil() {
	o.ChipNo.Set(nil)
}

// UnsetChipNo ensures that no value is present for ChipNo, not even an explicit nil
func (o *PatchedPart) UnsetChipNo() {
	o.ChipNo.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PatchedPart) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPart) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchedPart) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchedPart) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PatchedPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PartNo) {
		toSerialize["part_no"] = o.PartNo
	}
	if !isNil(o.PartFamily) {
		toSerialize["part_family"] = o.PartFamily
	}
	if o.BoardNo.IsSet() {
		toSerialize["board_no"] = o.BoardNo.Get()
	}
	if !isNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !isNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if o.ChipNo.IsSet() {
		toSerialize["chip_no"] = o.ChipNo.Get()
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedPart struct {
	value *PatchedPart
	isSet bool
}

func (v NullablePatchedPart) Get() *PatchedPart {
	return v.value
}

func (v *NullablePatchedPart) Set(val *PatchedPart) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPart) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPart(val *PatchedPart) *NullablePatchedPart {
	return &NullablePatchedPart{value: val, isSet: true}
}

func (v NullablePatchedPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


