/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
	"time"
)

// PatchedWriteOnlyDevice Serializes Device querysets to JSON
type PatchedWriteOnlyDevice struct {
	Id *int32 `json:"id,omitempty"`
	PermissionGroups []string `json:"permission_groups,omitempty"`
	Part *string `json:"part,omitempty"`
	LocDesc *string `json:"loc_desc,omitempty"`
	IsReserved NullableBool `json:"is_reserved,omitempty"`
	Location *LocationSerializerWriteOnly `json:"location,omitempty"`
	SnOrName NullableString `json:"sn_or_name,omitempty"`
	Source *DeviceGroupSource `json:"source,omitempty"`
	IsWstk *bool `json:"is_wstk,omitempty"`
	Uri NullableString `json:"uri,omitempty"`
	Online *bool `json:"online,omitempty"`
	Status *StatusF39Enum `json:"status,omitempty"`
	LastUpdate *time.Time `json:"last_update,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	WstkPart NullableString `json:"wstk_part,omitempty"`
}

// NewPatchedWriteOnlyDevice instantiates a new PatchedWriteOnlyDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWriteOnlyDevice() *PatchedWriteOnlyDevice {
	this := PatchedWriteOnlyDevice{}
	return &this
}

// NewPatchedWriteOnlyDeviceWithDefaults instantiates a new PatchedWriteOnlyDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWriteOnlyDeviceWithDefaults() *PatchedWriteOnlyDevice {
	this := PatchedWriteOnlyDevice{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedWriteOnlyDevice) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWriteOnlyDevice) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PatchedWriteOnlyDevice) SetId(v int32) {
	o.Id = &v
}

// GetPermissionGroups returns the PermissionGroups field value if set, zero value otherwise.
func (o *PatchedWriteOnlyDevice) GetPermissionGroups() []string {
	if o == nil || isNil(o.PermissionGroups) {
		var ret []string
		return ret
	}
	return o.PermissionGroups
}

// GetPermissionGroupsOk returns a tuple with the PermissionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWriteOnlyDevice) GetPermissionGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.PermissionGroups) {
    return nil, false
	}
	return o.PermissionGroups, true
}

// HasPermissionGroups returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasPermissionGroups() bool {
	if o != nil && !isNil(o.PermissionGroups) {
		return true
	}

	return false
}

// SetPermissionGroups gets a reference to the given []string and assigns it to the PermissionGroups field.
func (o *PatchedWriteOnlyDevice) SetPermissionGroups(v []string) {
	o.PermissionGroups = v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *PatchedWriteOnlyDevice) GetPart() string {
	if o == nil || isNil(o.Part) {
		var ret string
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWriteOnlyDevice) GetPartOk() (*string, bool) {
	if o == nil || isNil(o.Part) {
    return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasPart() bool {
	if o != nil && !isNil(o.Part) {
		return true
	}

	return false
}

// SetPart gets a reference to the given string and assigns it to the Part field.
func (o *PatchedWriteOnlyDevice) SetPart(v string) {
	o.Part = &v
}

// GetLocDesc returns the LocDesc field value if set, zero value otherwise.
func (o *PatchedWriteOnlyDevice) GetLocDesc() string {
	if o == nil || isNil(o.LocDesc) {
		var ret string
		return ret
	}
	return *o.LocDesc
}

// GetLocDescOk returns a tuple with the LocDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWriteOnlyDevice) GetLocDescOk() (*string, bool) {
	if o == nil || isNil(o.LocDesc) {
    return nil, false
	}
	return o.LocDesc, true
}

// HasLocDesc returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasLocDesc() bool {
	if o != nil && !isNil(o.LocDesc) {
		return true
	}

	return false
}

// SetLocDesc gets a reference to the given string and assigns it to the LocDesc field.
func (o *PatchedWriteOnlyDevice) SetLocDesc(v string) {
	o.LocDesc = &v
}

// GetIsReserved returns the IsReserved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWriteOnlyDevice) GetIsReserved() bool {
	if o == nil || isNil(o.IsReserved.Get()) {
		var ret bool
		return ret
	}
	return *o.IsReserved.Get()
}

// GetIsReservedOk returns a tuple with the IsReserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWriteOnlyDevice) GetIsReservedOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsReserved.Get(), o.IsReserved.IsSet()
}

// HasIsReserved returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasIsReserved() bool {
	if o != nil && o.IsReserved.IsSet() {
		return true
	}

	return false
}

// SetIsReserved gets a reference to the given NullableBool and assigns it to the IsReserved field.
func (o *PatchedWriteOnlyDevice) SetIsReserved(v bool) {
	o.IsReserved.Set(&v)
}
// SetIsReservedNil sets the value for IsReserved to be an explicit nil
func (o *PatchedWriteOnlyDevice) SetIsReservedNil() {
	o.IsReserved.Set(nil)
}

// UnsetIsReserved ensures that no value is present for IsReserved, not even an explicit nil
func (o *PatchedWriteOnlyDevice) UnsetIsReserved() {
	o.IsReserved.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PatchedWriteOnlyDevice) GetLocation() LocationSerializerWriteOnly {
	if o == nil || isNil(o.Location) {
		var ret LocationSerializerWriteOnly
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWriteOnlyDevice) GetLocationOk() (*LocationSerializerWriteOnly, bool) {
	if o == nil || isNil(o.Location) {
    return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasLocation() bool {
	if o != nil && !isNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given LocationSerializerWriteOnly and assigns it to the Location field.
func (o *PatchedWriteOnlyDevice) SetLocation(v LocationSerializerWriteOnly) {
	o.Location = &v
}

// GetSnOrName returns the SnOrName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWriteOnlyDevice) GetSnOrName() string {
	if o == nil || isNil(o.SnOrName.Get()) {
		var ret string
		return ret
	}
	return *o.SnOrName.Get()
}

// GetSnOrNameOk returns a tuple with the SnOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWriteOnlyDevice) GetSnOrNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SnOrName.Get(), o.SnOrName.IsSet()
}

// HasSnOrName returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasSnOrName() bool {
	if o != nil && o.SnOrName.IsSet() {
		return true
	}

	return false
}

// SetSnOrName gets a reference to the given NullableString and assigns it to the SnOrName field.
func (o *PatchedWriteOnlyDevice) SetSnOrName(v string) {
	o.SnOrName.Set(&v)
}
// SetSnOrNameNil sets the value for SnOrName to be an explicit nil
func (o *PatchedWriteOnlyDevice) SetSnOrNameNil() {
	o.SnOrName.Set(nil)
}

// UnsetSnOrName ensures that no value is present for SnOrName, not even an explicit nil
func (o *PatchedWriteOnlyDevice) UnsetSnOrName() {
	o.SnOrName.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PatchedWriteOnlyDevice) GetSource() DeviceGroupSource {
	if o == nil || isNil(o.Source) {
		var ret DeviceGroupSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWriteOnlyDevice) GetSourceOk() (*DeviceGroupSource, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given DeviceGroupSource and assigns it to the Source field.
func (o *PatchedWriteOnlyDevice) SetSource(v DeviceGroupSource) {
	o.Source = &v
}

// GetIsWstk returns the IsWstk field value if set, zero value otherwise.
func (o *PatchedWriteOnlyDevice) GetIsWstk() bool {
	if o == nil || isNil(o.IsWstk) {
		var ret bool
		return ret
	}
	return *o.IsWstk
}

// GetIsWstkOk returns a tuple with the IsWstk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWriteOnlyDevice) GetIsWstkOk() (*bool, bool) {
	if o == nil || isNil(o.IsWstk) {
    return nil, false
	}
	return o.IsWstk, true
}

// HasIsWstk returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasIsWstk() bool {
	if o != nil && !isNil(o.IsWstk) {
		return true
	}

	return false
}

// SetIsWstk gets a reference to the given bool and assigns it to the IsWstk field.
func (o *PatchedWriteOnlyDevice) SetIsWstk(v bool) {
	o.IsWstk = &v
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWriteOnlyDevice) GetUri() string {
	if o == nil || isNil(o.Uri.Get()) {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWriteOnlyDevice) GetUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *PatchedWriteOnlyDevice) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *PatchedWriteOnlyDevice) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *PatchedWriteOnlyDevice) UnsetUri() {
	o.Uri.Unset()
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *PatchedWriteOnlyDevice) GetOnline() bool {
	if o == nil || isNil(o.Online) {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWriteOnlyDevice) GetOnlineOk() (*bool, bool) {
	if o == nil || isNil(o.Online) {
    return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasOnline() bool {
	if o != nil && !isNil(o.Online) {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *PatchedWriteOnlyDevice) SetOnline(v bool) {
	o.Online = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWriteOnlyDevice) GetStatus() StatusF39Enum {
	if o == nil || isNil(o.Status) {
		var ret StatusF39Enum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWriteOnlyDevice) GetStatusOk() (*StatusF39Enum, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StatusF39Enum and assigns it to the Status field.
func (o *PatchedWriteOnlyDevice) SetStatus(v StatusF39Enum) {
	o.Status = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *PatchedWriteOnlyDevice) GetLastUpdate() time.Time {
	if o == nil || isNil(o.LastUpdate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWriteOnlyDevice) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastUpdate) {
    return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasLastUpdate() bool {
	if o != nil && !isNil(o.LastUpdate) {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given time.Time and assigns it to the LastUpdate field.
func (o *PatchedWriteOnlyDevice) SetLastUpdate(v time.Time) {
	o.LastUpdate = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *PatchedWriteOnlyDevice) GetDateCreated() time.Time {
	if o == nil || isNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWriteOnlyDevice) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.DateCreated) {
    return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasDateCreated() bool {
	if o != nil && !isNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *PatchedWriteOnlyDevice) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWriteOnlyDevice) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWriteOnlyDevice) GetMetadataOk() (*interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasMetadata() bool {
	if o != nil && isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PatchedWriteOnlyDevice) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetWstkPart returns the WstkPart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWriteOnlyDevice) GetWstkPart() string {
	if o == nil || isNil(o.WstkPart.Get()) {
		var ret string
		return ret
	}
	return *o.WstkPart.Get()
}

// GetWstkPartOk returns a tuple with the WstkPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWriteOnlyDevice) GetWstkPartOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.WstkPart.Get(), o.WstkPart.IsSet()
}

// HasWstkPart returns a boolean if a field has been set.
func (o *PatchedWriteOnlyDevice) HasWstkPart() bool {
	if o != nil && o.WstkPart.IsSet() {
		return true
	}

	return false
}

// SetWstkPart gets a reference to the given NullableString and assigns it to the WstkPart field.
func (o *PatchedWriteOnlyDevice) SetWstkPart(v string) {
	o.WstkPart.Set(&v)
}
// SetWstkPartNil sets the value for WstkPart to be an explicit nil
func (o *PatchedWriteOnlyDevice) SetWstkPartNil() {
	o.WstkPart.Set(nil)
}

// UnsetWstkPart ensures that no value is present for WstkPart, not even an explicit nil
func (o *PatchedWriteOnlyDevice) UnsetWstkPart() {
	o.WstkPart.Unset()
}

func (o PatchedWriteOnlyDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.PermissionGroups) {
		toSerialize["permission_groups"] = o.PermissionGroups
	}
	if !isNil(o.Part) {
		toSerialize["part"] = o.Part
	}
	if !isNil(o.LocDesc) {
		toSerialize["loc_desc"] = o.LocDesc
	}
	if o.IsReserved.IsSet() {
		toSerialize["is_reserved"] = o.IsReserved.Get()
	}
	if !isNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if o.SnOrName.IsSet() {
		toSerialize["sn_or_name"] = o.SnOrName.Get()
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.IsWstk) {
		toSerialize["is_wstk"] = o.IsWstk
	}
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
	}
	if !isNil(o.Online) {
		toSerialize["online"] = o.Online
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.LastUpdate) {
		toSerialize["last_update"] = o.LastUpdate
	}
	if !isNil(o.DateCreated) {
		toSerialize["date_created"] = o.DateCreated
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.WstkPart.IsSet() {
		toSerialize["wstk_part"] = o.WstkPart.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedWriteOnlyDevice struct {
	value *PatchedWriteOnlyDevice
	isSet bool
}

func (v NullablePatchedWriteOnlyDevice) Get() *PatchedWriteOnlyDevice {
	return v.value
}

func (v *NullablePatchedWriteOnlyDevice) Set(val *PatchedWriteOnlyDevice) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWriteOnlyDevice) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWriteOnlyDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWriteOnlyDevice(val *PatchedWriteOnlyDevice) *NullablePatchedWriteOnlyDevice {
	return &NullablePatchedWriteOnlyDevice{value: val, isSet: true}
}

func (v NullablePatchedWriteOnlyDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWriteOnlyDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


