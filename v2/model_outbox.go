/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
	"time"
)

// Outbox struct for Outbox
type Outbox struct {
	Id string `json:"id"`
	// record revision number
	Version *int32 `json:"version,omitempty"`
	TCreated time.Time `json:"t_created"`
	TUpdated time.Time `json:"t_updated"`
	TProcessed NullableTime `json:"t_processed,omitempty"`
	Payload map[string]interface{} `json:"payload"`
	OutboxStatus *OutboxStatusEnum `json:"outbox_status,omitempty"`
	SagaId string `json:"saga_id"`
	SagaStatus *SagaStatusEnum `json:"saga_status,omitempty"`
	ReservationId string `json:"reservation_id"`
	ReservationStatus *OutboxReservationStatus `json:"reservation_status,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewOutbox instantiates a new Outbox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutbox(id string, tCreated time.Time, tUpdated time.Time, payload map[string]interface{}, sagaId string, reservationId string) *Outbox {
	this := Outbox{}
	this.Id = id
	this.TCreated = tCreated
	this.TUpdated = tUpdated
	this.Payload = payload
	this.SagaId = sagaId
	this.ReservationId = reservationId
	return &this
}

// NewOutboxWithDefaults instantiates a new Outbox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutboxWithDefaults() *Outbox {
	this := Outbox{}
	return &this
}

// GetId returns the Id field value
func (o *Outbox) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Outbox) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Outbox) SetId(v string) {
	o.Id = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Outbox) GetVersion() int32 {
	if o == nil || isNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outbox) GetVersionOk() (*int32, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Outbox) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Outbox) SetVersion(v int32) {
	o.Version = &v
}

// GetTCreated returns the TCreated field value
func (o *Outbox) GetTCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TCreated
}

// GetTCreatedOk returns a tuple with the TCreated field value
// and a boolean to check if the value has been set.
func (o *Outbox) GetTCreatedOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TCreated, true
}

// SetTCreated sets field value
func (o *Outbox) SetTCreated(v time.Time) {
	o.TCreated = v
}

// GetTUpdated returns the TUpdated field value
func (o *Outbox) GetTUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TUpdated
}

// GetTUpdatedOk returns a tuple with the TUpdated field value
// and a boolean to check if the value has been set.
func (o *Outbox) GetTUpdatedOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TUpdated, true
}

// SetTUpdated sets field value
func (o *Outbox) SetTUpdated(v time.Time) {
	o.TUpdated = v
}

// GetTProcessed returns the TProcessed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Outbox) GetTProcessed() time.Time {
	if o == nil || isNil(o.TProcessed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.TProcessed.Get()
}

// GetTProcessedOk returns a tuple with the TProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Outbox) GetTProcessedOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.TProcessed.Get(), o.TProcessed.IsSet()
}

// HasTProcessed returns a boolean if a field has been set.
func (o *Outbox) HasTProcessed() bool {
	if o != nil && o.TProcessed.IsSet() {
		return true
	}

	return false
}

// SetTProcessed gets a reference to the given NullableTime and assigns it to the TProcessed field.
func (o *Outbox) SetTProcessed(v time.Time) {
	o.TProcessed.Set(&v)
}
// SetTProcessedNil sets the value for TProcessed to be an explicit nil
func (o *Outbox) SetTProcessedNil() {
	o.TProcessed.Set(nil)
}

// UnsetTProcessed ensures that no value is present for TProcessed, not even an explicit nil
func (o *Outbox) UnsetTProcessed() {
	o.TProcessed.Unset()
}

// GetPayload returns the Payload field value
func (o *Outbox) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *Outbox) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *Outbox) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetOutboxStatus returns the OutboxStatus field value if set, zero value otherwise.
func (o *Outbox) GetOutboxStatus() OutboxStatusEnum {
	if o == nil || isNil(o.OutboxStatus) {
		var ret OutboxStatusEnum
		return ret
	}
	return *o.OutboxStatus
}

// GetOutboxStatusOk returns a tuple with the OutboxStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outbox) GetOutboxStatusOk() (*OutboxStatusEnum, bool) {
	if o == nil || isNil(o.OutboxStatus) {
    return nil, false
	}
	return o.OutboxStatus, true
}

// HasOutboxStatus returns a boolean if a field has been set.
func (o *Outbox) HasOutboxStatus() bool {
	if o != nil && !isNil(o.OutboxStatus) {
		return true
	}

	return false
}

// SetOutboxStatus gets a reference to the given OutboxStatusEnum and assigns it to the OutboxStatus field.
func (o *Outbox) SetOutboxStatus(v OutboxStatusEnum) {
	o.OutboxStatus = &v
}

// GetSagaId returns the SagaId field value
func (o *Outbox) GetSagaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SagaId
}

// GetSagaIdOk returns a tuple with the SagaId field value
// and a boolean to check if the value has been set.
func (o *Outbox) GetSagaIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SagaId, true
}

// SetSagaId sets field value
func (o *Outbox) SetSagaId(v string) {
	o.SagaId = v
}

// GetSagaStatus returns the SagaStatus field value if set, zero value otherwise.
func (o *Outbox) GetSagaStatus() SagaStatusEnum {
	if o == nil || isNil(o.SagaStatus) {
		var ret SagaStatusEnum
		return ret
	}
	return *o.SagaStatus
}

// GetSagaStatusOk returns a tuple with the SagaStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outbox) GetSagaStatusOk() (*SagaStatusEnum, bool) {
	if o == nil || isNil(o.SagaStatus) {
    return nil, false
	}
	return o.SagaStatus, true
}

// HasSagaStatus returns a boolean if a field has been set.
func (o *Outbox) HasSagaStatus() bool {
	if o != nil && !isNil(o.SagaStatus) {
		return true
	}

	return false
}

// SetSagaStatus gets a reference to the given SagaStatusEnum and assigns it to the SagaStatus field.
func (o *Outbox) SetSagaStatus(v SagaStatusEnum) {
	o.SagaStatus = &v
}

// GetReservationId returns the ReservationId field value
func (o *Outbox) GetReservationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value
// and a boolean to check if the value has been set.
func (o *Outbox) GetReservationIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReservationId, true
}

// SetReservationId sets field value
func (o *Outbox) SetReservationId(v string) {
	o.ReservationId = v
}

// GetReservationStatus returns the ReservationStatus field value if set, zero value otherwise.
func (o *Outbox) GetReservationStatus() OutboxReservationStatus {
	if o == nil || isNil(o.ReservationStatus) {
		var ret OutboxReservationStatus
		return ret
	}
	return *o.ReservationStatus
}

// GetReservationStatusOk returns a tuple with the ReservationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outbox) GetReservationStatusOk() (*OutboxReservationStatus, bool) {
	if o == nil || isNil(o.ReservationStatus) {
    return nil, false
	}
	return o.ReservationStatus, true
}

// HasReservationStatus returns a boolean if a field has been set.
func (o *Outbox) HasReservationStatus() bool {
	if o != nil && !isNil(o.ReservationStatus) {
		return true
	}

	return false
}

// SetReservationStatus gets a reference to the given OutboxReservationStatus and assigns it to the ReservationStatus field.
func (o *Outbox) SetReservationStatus(v OutboxReservationStatus) {
	o.ReservationStatus = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Outbox) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outbox) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Outbox) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Outbox) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o Outbox) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["t_created"] = o.TCreated
	}
	if true {
		toSerialize["t_updated"] = o.TUpdated
	}
	if o.TProcessed.IsSet() {
		toSerialize["t_processed"] = o.TProcessed.Get()
	}
	if true {
		toSerialize["payload"] = o.Payload
	}
	if !isNil(o.OutboxStatus) {
		toSerialize["outbox_status"] = o.OutboxStatus
	}
	if true {
		toSerialize["saga_id"] = o.SagaId
	}
	if !isNil(o.SagaStatus) {
		toSerialize["saga_status"] = o.SagaStatus
	}
	if true {
		toSerialize["reservation_id"] = o.ReservationId
	}
	if !isNil(o.ReservationStatus) {
		toSerialize["reservation_status"] = o.ReservationStatus
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableOutbox struct {
	value *Outbox
	isSet bool
}

func (v NullableOutbox) Get() *Outbox {
	return v.value
}

func (v *NullableOutbox) Set(val *Outbox) {
	v.value = val
	v.isSet = true
}

func (v NullableOutbox) IsSet() bool {
	return v.isSet
}

func (v *NullableOutbox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutbox(val *Outbox) *NullableOutbox {
	return &NullableOutbox{value: val, isSet: true}
}

func (v NullableOutbox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutbox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


