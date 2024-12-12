/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
	"time"
)

// WriteOnlyDevice Serializes Device querysets to JSON
type WriteOnlyDevice struct {
	Id int32 `json:"id"`
	PermissionGroups []string `json:"permission_groups,omitempty"`
	Part string `json:"part"`
	LocDesc string `json:"loc_desc"`
	IsReserved NullableBool `json:"is_reserved"`
	Location LocationSerializerWriteOnly `json:"location"`
	SnOrName NullableString `json:"sn_or_name,omitempty"`
	Source *SourceEnum `json:"source,omitempty"`
	SocketedChip *string `json:"socketed_chip,omitempty"`
	IsWstk *bool `json:"is_wstk,omitempty"`
	Uri NullableString `json:"uri,omitempty"`
	Online *bool `json:"online,omitempty"`
	Status *StatusEnum `json:"status,omitempty"`
	LastUpdate time.Time `json:"last_update"`
	DateCreated time.Time `json:"date_created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	IsReservedFilter *bool `json:"is_reserved_filter,omitempty"`
	WstkPart NullableString `json:"wstk_part,omitempty"`
}

// NewWriteOnlyDevice instantiates a new WriteOnlyDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteOnlyDevice(id int32, part string, locDesc string, isReserved NullableBool, location LocationSerializerWriteOnly, lastUpdate time.Time, dateCreated time.Time) *WriteOnlyDevice {
	this := WriteOnlyDevice{}
	this.Id = id
	this.Part = part
	this.LocDesc = locDesc
	this.IsReserved = isReserved
	this.Location = location
	this.LastUpdate = lastUpdate
	this.DateCreated = dateCreated
	return &this
}

// NewWriteOnlyDeviceWithDefaults instantiates a new WriteOnlyDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteOnlyDeviceWithDefaults() *WriteOnlyDevice {
	this := WriteOnlyDevice{}
	return &this
}

// GetId returns the Id field value
func (o *WriteOnlyDevice) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WriteOnlyDevice) SetId(v int32) {
	o.Id = v
}

// GetPermissionGroups returns the PermissionGroups field value if set, zero value otherwise.
func (o *WriteOnlyDevice) GetPermissionGroups() []string {
	if o == nil || isNil(o.PermissionGroups) {
		var ret []string
		return ret
	}
	return o.PermissionGroups
}

// GetPermissionGroupsOk returns a tuple with the PermissionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetPermissionGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.PermissionGroups) {
    return nil, false
	}
	return o.PermissionGroups, true
}

// HasPermissionGroups returns a boolean if a field has been set.
func (o *WriteOnlyDevice) HasPermissionGroups() bool {
	if o != nil && !isNil(o.PermissionGroups) {
		return true
	}

	return false
}

// SetPermissionGroups gets a reference to the given []string and assigns it to the PermissionGroups field.
func (o *WriteOnlyDevice) SetPermissionGroups(v []string) {
	o.PermissionGroups = v
}

// GetPart returns the Part field value
func (o *WriteOnlyDevice) GetPart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Part
}

// GetPartOk returns a tuple with the Part field value
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetPartOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Part, true
}

// SetPart sets field value
func (o *WriteOnlyDevice) SetPart(v string) {
	o.Part = v
}

// GetLocDesc returns the LocDesc field value
func (o *WriteOnlyDevice) GetLocDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocDesc
}

// GetLocDescOk returns a tuple with the LocDesc field value
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetLocDescOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocDesc, true
}

// SetLocDesc sets field value
func (o *WriteOnlyDevice) SetLocDesc(v string) {
	o.LocDesc = v
}

// GetIsReserved returns the IsReserved field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *WriteOnlyDevice) GetIsReserved() bool {
	if o == nil || o.IsReserved.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsReserved.Get()
}

// GetIsReservedOk returns a tuple with the IsReserved field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WriteOnlyDevice) GetIsReservedOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsReserved.Get(), o.IsReserved.IsSet()
}

// SetIsReserved sets field value
func (o *WriteOnlyDevice) SetIsReserved(v bool) {
	o.IsReserved.Set(&v)
}

// GetLocation returns the Location field value
func (o *WriteOnlyDevice) GetLocation() LocationSerializerWriteOnly {
	if o == nil {
		var ret LocationSerializerWriteOnly
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetLocationOk() (*LocationSerializerWriteOnly, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *WriteOnlyDevice) SetLocation(v LocationSerializerWriteOnly) {
	o.Location = v
}

// GetSnOrName returns the SnOrName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WriteOnlyDevice) GetSnOrName() string {
	if o == nil || isNil(o.SnOrName.Get()) {
		var ret string
		return ret
	}
	return *o.SnOrName.Get()
}

// GetSnOrNameOk returns a tuple with the SnOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WriteOnlyDevice) GetSnOrNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SnOrName.Get(), o.SnOrName.IsSet()
}

// HasSnOrName returns a boolean if a field has been set.
func (o *WriteOnlyDevice) HasSnOrName() bool {
	if o != nil && o.SnOrName.IsSet() {
		return true
	}

	return false
}

// SetSnOrName gets a reference to the given NullableString and assigns it to the SnOrName field.
func (o *WriteOnlyDevice) SetSnOrName(v string) {
	o.SnOrName.Set(&v)
}
// SetSnOrNameNil sets the value for SnOrName to be an explicit nil
func (o *WriteOnlyDevice) SetSnOrNameNil() {
	o.SnOrName.Set(nil)
}

// UnsetSnOrName ensures that no value is present for SnOrName, not even an explicit nil
func (o *WriteOnlyDevice) UnsetSnOrName() {
	o.SnOrName.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *WriteOnlyDevice) GetSource() SourceEnum {
	if o == nil || isNil(o.Source) {
		var ret SourceEnum
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetSourceOk() (*SourceEnum, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *WriteOnlyDevice) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given SourceEnum and assigns it to the Source field.
func (o *WriteOnlyDevice) SetSource(v SourceEnum) {
	o.Source = &v
}

// GetSocketedChip returns the SocketedChip field value if set, zero value otherwise.
func (o *WriteOnlyDevice) GetSocketedChip() string {
	if o == nil || isNil(o.SocketedChip) {
		var ret string
		return ret
	}
	return *o.SocketedChip
}

// GetSocketedChipOk returns a tuple with the SocketedChip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetSocketedChipOk() (*string, bool) {
	if o == nil || isNil(o.SocketedChip) {
    return nil, false
	}
	return o.SocketedChip, true
}

// HasSocketedChip returns a boolean if a field has been set.
func (o *WriteOnlyDevice) HasSocketedChip() bool {
	if o != nil && !isNil(o.SocketedChip) {
		return true
	}

	return false
}

// SetSocketedChip gets a reference to the given string and assigns it to the SocketedChip field.
func (o *WriteOnlyDevice) SetSocketedChip(v string) {
	o.SocketedChip = &v
}

// GetIsWstk returns the IsWstk field value if set, zero value otherwise.
func (o *WriteOnlyDevice) GetIsWstk() bool {
	if o == nil || isNil(o.IsWstk) {
		var ret bool
		return ret
	}
	return *o.IsWstk
}

// GetIsWstkOk returns a tuple with the IsWstk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetIsWstkOk() (*bool, bool) {
	if o == nil || isNil(o.IsWstk) {
    return nil, false
	}
	return o.IsWstk, true
}

// HasIsWstk returns a boolean if a field has been set.
func (o *WriteOnlyDevice) HasIsWstk() bool {
	if o != nil && !isNil(o.IsWstk) {
		return true
	}

	return false
}

// SetIsWstk gets a reference to the given bool and assigns it to the IsWstk field.
func (o *WriteOnlyDevice) SetIsWstk(v bool) {
	o.IsWstk = &v
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WriteOnlyDevice) GetUri() string {
	if o == nil || isNil(o.Uri.Get()) {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WriteOnlyDevice) GetUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *WriteOnlyDevice) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *WriteOnlyDevice) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *WriteOnlyDevice) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *WriteOnlyDevice) UnsetUri() {
	o.Uri.Unset()
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *WriteOnlyDevice) GetOnline() bool {
	if o == nil || isNil(o.Online) {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetOnlineOk() (*bool, bool) {
	if o == nil || isNil(o.Online) {
    return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *WriteOnlyDevice) HasOnline() bool {
	if o != nil && !isNil(o.Online) {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *WriteOnlyDevice) SetOnline(v bool) {
	o.Online = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WriteOnlyDevice) GetStatus() StatusEnum {
	if o == nil || isNil(o.Status) {
		var ret StatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetStatusOk() (*StatusEnum, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WriteOnlyDevice) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StatusEnum and assigns it to the Status field.
func (o *WriteOnlyDevice) SetStatus(v StatusEnum) {
	o.Status = &v
}

// GetLastUpdate returns the LastUpdate field value
func (o *WriteOnlyDevice) GetLastUpdate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LastUpdate, true
}

// SetLastUpdate sets field value
func (o *WriteOnlyDevice) SetLastUpdate(v time.Time) {
	o.LastUpdate = v
}

// GetDateCreated returns the DateCreated field value
func (o *WriteOnlyDevice) GetDateCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DateCreated, true
}

// SetDateCreated sets field value
func (o *WriteOnlyDevice) SetDateCreated(v time.Time) {
	o.DateCreated = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *WriteOnlyDevice) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WriteOnlyDevice) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *WriteOnlyDevice) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetIsReservedFilter returns the IsReservedFilter field value if set, zero value otherwise.
func (o *WriteOnlyDevice) GetIsReservedFilter() bool {
	if o == nil || isNil(o.IsReservedFilter) {
		var ret bool
		return ret
	}
	return *o.IsReservedFilter
}

// GetIsReservedFilterOk returns a tuple with the IsReservedFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteOnlyDevice) GetIsReservedFilterOk() (*bool, bool) {
	if o == nil || isNil(o.IsReservedFilter) {
    return nil, false
	}
	return o.IsReservedFilter, true
}

// HasIsReservedFilter returns a boolean if a field has been set.
func (o *WriteOnlyDevice) HasIsReservedFilter() bool {
	if o != nil && !isNil(o.IsReservedFilter) {
		return true
	}

	return false
}

// SetIsReservedFilter gets a reference to the given bool and assigns it to the IsReservedFilter field.
func (o *WriteOnlyDevice) SetIsReservedFilter(v bool) {
	o.IsReservedFilter = &v
}

// GetWstkPart returns the WstkPart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WriteOnlyDevice) GetWstkPart() string {
	if o == nil || isNil(o.WstkPart.Get()) {
		var ret string
		return ret
	}
	return *o.WstkPart.Get()
}

// GetWstkPartOk returns a tuple with the WstkPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WriteOnlyDevice) GetWstkPartOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.WstkPart.Get(), o.WstkPart.IsSet()
}

// HasWstkPart returns a boolean if a field has been set.
func (o *WriteOnlyDevice) HasWstkPart() bool {
	if o != nil && o.WstkPart.IsSet() {
		return true
	}

	return false
}

// SetWstkPart gets a reference to the given NullableString and assigns it to the WstkPart field.
func (o *WriteOnlyDevice) SetWstkPart(v string) {
	o.WstkPart.Set(&v)
}
// SetWstkPartNil sets the value for WstkPart to be an explicit nil
func (o *WriteOnlyDevice) SetWstkPartNil() {
	o.WstkPart.Set(nil)
}

// UnsetWstkPart ensures that no value is present for WstkPart, not even an explicit nil
func (o *WriteOnlyDevice) UnsetWstkPart() {
	o.WstkPart.Unset()
}

func (o WriteOnlyDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.PermissionGroups) {
		toSerialize["permission_groups"] = o.PermissionGroups
	}
	if true {
		toSerialize["part"] = o.Part
	}
	if true {
		toSerialize["loc_desc"] = o.LocDesc
	}
	if true {
		toSerialize["is_reserved"] = o.IsReserved.Get()
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if o.SnOrName.IsSet() {
		toSerialize["sn_or_name"] = o.SnOrName.Get()
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.SocketedChip) {
		toSerialize["socketed_chip"] = o.SocketedChip
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
	if true {
		toSerialize["last_update"] = o.LastUpdate
	}
	if true {
		toSerialize["date_created"] = o.DateCreated
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.IsReservedFilter) {
		toSerialize["is_reserved_filter"] = o.IsReservedFilter
	}
	if o.WstkPart.IsSet() {
		toSerialize["wstk_part"] = o.WstkPart.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableWriteOnlyDevice struct {
	value *WriteOnlyDevice
	isSet bool
}

func (v NullableWriteOnlyDevice) Get() *WriteOnlyDevice {
	return v.value
}

func (v *NullableWriteOnlyDevice) Set(val *WriteOnlyDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteOnlyDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteOnlyDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteOnlyDevice(val *WriteOnlyDevice) *NullableWriteOnlyDevice {
	return &NullableWriteOnlyDevice{value: val, isSet: true}
}

func (v NullableWriteOnlyDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteOnlyDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


