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

// PaginatedDeviceGroupList struct for PaginatedDeviceGroupList
type PaginatedDeviceGroupList struct {
	Count int32 `json:"count"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []DeviceGroup `json:"results"`
}

// NewPaginatedDeviceGroupList instantiates a new PaginatedDeviceGroupList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedDeviceGroupList(count int32, results []DeviceGroup) *PaginatedDeviceGroupList {
	this := PaginatedDeviceGroupList{}
	this.Count = count
	this.Results = results
	return &this
}

// NewPaginatedDeviceGroupListWithDefaults instantiates a new PaginatedDeviceGroupList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedDeviceGroupListWithDefaults() *PaginatedDeviceGroupList {
	this := PaginatedDeviceGroupList{}
	return &this
}

// GetCount returns the Count field value
func (o *PaginatedDeviceGroupList) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginatedDeviceGroupList) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PaginatedDeviceGroupList) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedDeviceGroupList) GetNext() string {
	if o == nil || isNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedDeviceGroupList) GetNextOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedDeviceGroupList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedDeviceGroupList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedDeviceGroupList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedDeviceGroupList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedDeviceGroupList) GetPrevious() string {
	if o == nil || isNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedDeviceGroupList) GetPreviousOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedDeviceGroupList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedDeviceGroupList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedDeviceGroupList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedDeviceGroupList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value
func (o *PaginatedDeviceGroupList) GetResults() []DeviceGroup {
	if o == nil {
		var ret []DeviceGroup
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedDeviceGroupList) GetResultsOk() ([]DeviceGroup, bool) {
	if o == nil {
    return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedDeviceGroupList) SetResults(v []DeviceGroup) {
	o.Results = v
}

func (o PaginatedDeviceGroupList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedDeviceGroupList struct {
	value *PaginatedDeviceGroupList
	isSet bool
}

func (v NullablePaginatedDeviceGroupList) Get() *PaginatedDeviceGroupList {
	return v.value
}

func (v *NullablePaginatedDeviceGroupList) Set(val *PaginatedDeviceGroupList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedDeviceGroupList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedDeviceGroupList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedDeviceGroupList(val *PaginatedDeviceGroupList) *NullablePaginatedDeviceGroupList {
	return &NullablePaginatedDeviceGroupList{value: val, isSet: true}
}

func (v NullablePaginatedDeviceGroupList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedDeviceGroupList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


