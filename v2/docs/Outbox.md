# Outbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Version** | Pointer to **int32** | record revision number | [optional] 
**TCreated** | **time.Time** |  | [readonly] 
**TUpdated** | **time.Time** |  | [readonly] 
**TProcessed** | Pointer to **NullableTime** |  | [optional] 
**Payload** | **map[string]interface{}** |  | 
**OutboxStatus** | Pointer to [**OutboxStatusEnum**](OutboxStatusEnum.md) |  | [optional] 
**SagaId** | **string** |  | 
**SagaStatus** | Pointer to [**SagaStatusEnum**](SagaStatusEnum.md) |  | [optional] 
**ReservationId** | **string** |  | 
**ReservationStatus** | Pointer to [**OutboxReservationStatus**](OutboxReservationStatus.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOutbox

`func NewOutbox(id string, tCreated time.Time, tUpdated time.Time, payload map[string]interface{}, sagaId string, reservationId string, ) *Outbox`

NewOutbox instantiates a new Outbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboxWithDefaults

`func NewOutboxWithDefaults() *Outbox`

NewOutboxWithDefaults instantiates a new Outbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Outbox) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Outbox) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Outbox) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *Outbox) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Outbox) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Outbox) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Outbox) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTCreated

`func (o *Outbox) GetTCreated() time.Time`

GetTCreated returns the TCreated field if non-nil, zero value otherwise.

### GetTCreatedOk

`func (o *Outbox) GetTCreatedOk() (*time.Time, bool)`

GetTCreatedOk returns a tuple with the TCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTCreated

`func (o *Outbox) SetTCreated(v time.Time)`

SetTCreated sets TCreated field to given value.


### GetTUpdated

`func (o *Outbox) GetTUpdated() time.Time`

GetTUpdated returns the TUpdated field if non-nil, zero value otherwise.

### GetTUpdatedOk

`func (o *Outbox) GetTUpdatedOk() (*time.Time, bool)`

GetTUpdatedOk returns a tuple with the TUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTUpdated

`func (o *Outbox) SetTUpdated(v time.Time)`

SetTUpdated sets TUpdated field to given value.


### GetTProcessed

`func (o *Outbox) GetTProcessed() time.Time`

GetTProcessed returns the TProcessed field if non-nil, zero value otherwise.

### GetTProcessedOk

`func (o *Outbox) GetTProcessedOk() (*time.Time, bool)`

GetTProcessedOk returns a tuple with the TProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTProcessed

`func (o *Outbox) SetTProcessed(v time.Time)`

SetTProcessed sets TProcessed field to given value.

### HasTProcessed

`func (o *Outbox) HasTProcessed() bool`

HasTProcessed returns a boolean if a field has been set.

### SetTProcessedNil

`func (o *Outbox) SetTProcessedNil(b bool)`

 SetTProcessedNil sets the value for TProcessed to be an explicit nil

### UnsetTProcessed
`func (o *Outbox) UnsetTProcessed()`

UnsetTProcessed ensures that no value is present for TProcessed, not even an explicit nil
### GetPayload

`func (o *Outbox) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Outbox) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Outbox) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### GetOutboxStatus

`func (o *Outbox) GetOutboxStatus() OutboxStatusEnum`

GetOutboxStatus returns the OutboxStatus field if non-nil, zero value otherwise.

### GetOutboxStatusOk

`func (o *Outbox) GetOutboxStatusOk() (*OutboxStatusEnum, bool)`

GetOutboxStatusOk returns a tuple with the OutboxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboxStatus

`func (o *Outbox) SetOutboxStatus(v OutboxStatusEnum)`

SetOutboxStatus sets OutboxStatus field to given value.

### HasOutboxStatus

`func (o *Outbox) HasOutboxStatus() bool`

HasOutboxStatus returns a boolean if a field has been set.

### GetSagaId

`func (o *Outbox) GetSagaId() string`

GetSagaId returns the SagaId field if non-nil, zero value otherwise.

### GetSagaIdOk

`func (o *Outbox) GetSagaIdOk() (*string, bool)`

GetSagaIdOk returns a tuple with the SagaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSagaId

`func (o *Outbox) SetSagaId(v string)`

SetSagaId sets SagaId field to given value.


### GetSagaStatus

`func (o *Outbox) GetSagaStatus() SagaStatusEnum`

GetSagaStatus returns the SagaStatus field if non-nil, zero value otherwise.

### GetSagaStatusOk

`func (o *Outbox) GetSagaStatusOk() (*SagaStatusEnum, bool)`

GetSagaStatusOk returns a tuple with the SagaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSagaStatus

`func (o *Outbox) SetSagaStatus(v SagaStatusEnum)`

SetSagaStatus sets SagaStatus field to given value.

### HasSagaStatus

`func (o *Outbox) HasSagaStatus() bool`

HasSagaStatus returns a boolean if a field has been set.

### GetReservationId

`func (o *Outbox) GetReservationId() string`

GetReservationId returns the ReservationId field if non-nil, zero value otherwise.

### GetReservationIdOk

`func (o *Outbox) GetReservationIdOk() (*string, bool)`

GetReservationIdOk returns a tuple with the ReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationId

`func (o *Outbox) SetReservationId(v string)`

SetReservationId sets ReservationId field to given value.


### GetReservationStatus

`func (o *Outbox) GetReservationStatus() OutboxReservationStatus`

GetReservationStatus returns the ReservationStatus field if non-nil, zero value otherwise.

### GetReservationStatusOk

`func (o *Outbox) GetReservationStatusOk() (*OutboxReservationStatus, bool)`

GetReservationStatusOk returns a tuple with the ReservationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationStatus

`func (o *Outbox) SetReservationStatus(v OutboxReservationStatus)`

SetReservationStatus sets ReservationStatus field to given value.

### HasReservationStatus

`func (o *Outbox) HasReservationStatus() bool`

HasReservationStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *Outbox) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Outbox) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Outbox) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Outbox) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


