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

// PaginatedPartList struct for PaginatedPartList
type PaginatedPartList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []Part `json:"results,omitempty"`
}

// NewPaginatedPartList instantiates a new PaginatedPartList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPartList() *PaginatedPartList {
	this := PaginatedPartList{}
	return &this
}

// NewPaginatedPartListWithDefaults instantiates a new PaginatedPartList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPartListWithDefaults() *PaginatedPartList {
	this := PaginatedPartList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedPartList) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedPartList) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedPartList) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedPartList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedPartList) GetNext() string {
	if o == nil || isNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedPartList) GetNextOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedPartList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedPartList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedPartList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedPartList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedPartList) GetPrevious() string {
	if o == nil || isNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedPartList) GetPreviousOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedPartList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedPartList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedPartList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedPartList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedPartList) GetResults() []Part {
	if o == nil || isNil(o.Results) {
		var ret []Part
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedPartList) GetResultsOk() ([]Part, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedPartList) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Part and assigns it to the Results field.
func (o *PaginatedPartList) SetResults(v []Part) {
	o.Results = v
}

func (o PaginatedPartList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedPartList struct {
	value *PaginatedPartList
	isSet bool
}

func (v NullablePaginatedPartList) Get() *PaginatedPartList {
	return v.value
}

func (v *NullablePaginatedPartList) Set(val *PaginatedPartList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPartList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPartList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPartList(val *PaginatedPartList) *NullablePaginatedPartList {
	return &NullablePaginatedPartList{value: val, isSet: true}
}

func (v NullablePaginatedPartList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPartList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


