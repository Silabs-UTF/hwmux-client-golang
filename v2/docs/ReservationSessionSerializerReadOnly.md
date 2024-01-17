# ReservationSessionSerializerReadOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Owner** | [**NullableReservationSessionSerializerReadOnlyOwner**](ReservationSessionSerializerReadOnlyOwner.md) |  | 
**IsSatisfied** | **bool** |  | [readonly] 
**ResponseMessage** | **string** |  | [readonly] 
**Status** | **NullableString** |  | [readonly] 
**QueuePosition** | **NullableInt32** |  | [readonly] 
**Queue** | **[]string** |  | [readonly] 
**ResourceData** | **map[string]interface{}** |  | [readonly] 
**Details** | **string** |  | [readonly] 
**UseWatchdog** | Pointer to **bool** |  | [optional] 
**InvertPriority** | **bool** |  | [readonly] 
**CancelExisting** | **NullableBool** |  | [readonly] 
**TRequested** | **time.Time** |  | [readonly] 
**TSatisfied** | **NullableTime** |  | [readonly] 
**TCompleted** | **NullableTime** |  | [readonly] 
**TLeaseExpires** | **NullableTime** |  | [readonly] 
**Metadata** | **map[string]interface{}** |  | [readonly] 
**RDevices** | **[]int32** |  | [readonly] 
**RDeviceGroups** | **[]int32** |  | [readonly] 
**RDeviceGroupLabels** | **[]int32** |  | [readonly] 
**ADevices** | **[]int32** |  | [readonly] 
**ADeviceGroups** | **[]int32** |  | [readonly] 

## Methods

### NewReservationSessionSerializerReadOnly

`func NewReservationSessionSerializerReadOnly(id string, owner NullableReservationSessionSerializerReadOnlyOwner, isSatisfied bool, responseMessage string, status NullableString, queuePosition NullableInt32, queue []string, resourceData map[string]interface{}, details string, invertPriority bool, cancelExisting NullableBool, tRequested time.Time, tSatisfied NullableTime, tCompleted NullableTime, tLeaseExpires NullableTime, metadata map[string]interface{}, rDevices []int32, rDeviceGroups []int32, rDeviceGroupLabels []int32, aDevices []int32, aDeviceGroups []int32, ) *ReservationSessionSerializerReadOnly`

NewReservationSessionSerializerReadOnly instantiates a new ReservationSessionSerializerReadOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationSessionSerializerReadOnlyWithDefaults

`func NewReservationSessionSerializerReadOnlyWithDefaults() *ReservationSessionSerializerReadOnly`

NewReservationSessionSerializerReadOnlyWithDefaults instantiates a new ReservationSessionSerializerReadOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReservationSessionSerializerReadOnly) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReservationSessionSerializerReadOnly) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReservationSessionSerializerReadOnly) SetId(v string)`

SetId sets Id field to given value.


### GetOwner

`func (o *ReservationSessionSerializerReadOnly) GetOwner() ReservationSessionSerializerReadOnlyOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ReservationSessionSerializerReadOnly) GetOwnerOk() (*ReservationSessionSerializerReadOnlyOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ReservationSessionSerializerReadOnly) SetOwner(v ReservationSessionSerializerReadOnlyOwner)`

SetOwner sets Owner field to given value.


### SetOwnerNil

`func (o *ReservationSessionSerializerReadOnly) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ReservationSessionSerializerReadOnly) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetIsSatisfied

`func (o *ReservationSessionSerializerReadOnly) GetIsSatisfied() bool`

GetIsSatisfied returns the IsSatisfied field if non-nil, zero value otherwise.

### GetIsSatisfiedOk

`func (o *ReservationSessionSerializerReadOnly) GetIsSatisfiedOk() (*bool, bool)`

GetIsSatisfiedOk returns a tuple with the IsSatisfied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSatisfied

`func (o *ReservationSessionSerializerReadOnly) SetIsSatisfied(v bool)`

SetIsSatisfied sets IsSatisfied field to given value.


### GetResponseMessage

`func (o *ReservationSessionSerializerReadOnly) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *ReservationSessionSerializerReadOnly) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *ReservationSessionSerializerReadOnly) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetStatus

`func (o *ReservationSessionSerializerReadOnly) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReservationSessionSerializerReadOnly) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReservationSessionSerializerReadOnly) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *ReservationSessionSerializerReadOnly) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ReservationSessionSerializerReadOnly) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetQueuePosition

`func (o *ReservationSessionSerializerReadOnly) GetQueuePosition() int32`

GetQueuePosition returns the QueuePosition field if non-nil, zero value otherwise.

### GetQueuePositionOk

`func (o *ReservationSessionSerializerReadOnly) GetQueuePositionOk() (*int32, bool)`

GetQueuePositionOk returns a tuple with the QueuePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePosition

`func (o *ReservationSessionSerializerReadOnly) SetQueuePosition(v int32)`

SetQueuePosition sets QueuePosition field to given value.


### SetQueuePositionNil

`func (o *ReservationSessionSerializerReadOnly) SetQueuePositionNil(b bool)`

 SetQueuePositionNil sets the value for QueuePosition to be an explicit nil

### UnsetQueuePosition
`func (o *ReservationSessionSerializerReadOnly) UnsetQueuePosition()`

UnsetQueuePosition ensures that no value is present for QueuePosition, not even an explicit nil
### GetQueue

`func (o *ReservationSessionSerializerReadOnly) GetQueue() []string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *ReservationSessionSerializerReadOnly) GetQueueOk() (*[]string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *ReservationSessionSerializerReadOnly) SetQueue(v []string)`

SetQueue sets Queue field to given value.


### SetQueueNil

`func (o *ReservationSessionSerializerReadOnly) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *ReservationSessionSerializerReadOnly) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetResourceData

`func (o *ReservationSessionSerializerReadOnly) GetResourceData() map[string]interface{}`

GetResourceData returns the ResourceData field if non-nil, zero value otherwise.

### GetResourceDataOk

`func (o *ReservationSessionSerializerReadOnly) GetResourceDataOk() (*map[string]interface{}, bool)`

GetResourceDataOk returns a tuple with the ResourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceData

`func (o *ReservationSessionSerializerReadOnly) SetResourceData(v map[string]interface{})`

SetResourceData sets ResourceData field to given value.


### SetResourceDataNil

`func (o *ReservationSessionSerializerReadOnly) SetResourceDataNil(b bool)`

 SetResourceDataNil sets the value for ResourceData to be an explicit nil

### UnsetResourceData
`func (o *ReservationSessionSerializerReadOnly) UnsetResourceData()`

UnsetResourceData ensures that no value is present for ResourceData, not even an explicit nil
### GetDetails

`func (o *ReservationSessionSerializerReadOnly) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ReservationSessionSerializerReadOnly) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ReservationSessionSerializerReadOnly) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetUseWatchdog

`func (o *ReservationSessionSerializerReadOnly) GetUseWatchdog() bool`

GetUseWatchdog returns the UseWatchdog field if non-nil, zero value otherwise.

### GetUseWatchdogOk

`func (o *ReservationSessionSerializerReadOnly) GetUseWatchdogOk() (*bool, bool)`

GetUseWatchdogOk returns a tuple with the UseWatchdog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWatchdog

`func (o *ReservationSessionSerializerReadOnly) SetUseWatchdog(v bool)`

SetUseWatchdog sets UseWatchdog field to given value.

### HasUseWatchdog

`func (o *ReservationSessionSerializerReadOnly) HasUseWatchdog() bool`

HasUseWatchdog returns a boolean if a field has been set.

### GetInvertPriority

`func (o *ReservationSessionSerializerReadOnly) GetInvertPriority() bool`

GetInvertPriority returns the InvertPriority field if non-nil, zero value otherwise.

### GetInvertPriorityOk

`func (o *ReservationSessionSerializerReadOnly) GetInvertPriorityOk() (*bool, bool)`

GetInvertPriorityOk returns a tuple with the InvertPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertPriority

`func (o *ReservationSessionSerializerReadOnly) SetInvertPriority(v bool)`

SetInvertPriority sets InvertPriority field to given value.


### GetCancelExisting

`func (o *ReservationSessionSerializerReadOnly) GetCancelExisting() bool`

GetCancelExisting returns the CancelExisting field if non-nil, zero value otherwise.

### GetCancelExistingOk

`func (o *ReservationSessionSerializerReadOnly) GetCancelExistingOk() (*bool, bool)`

GetCancelExistingOk returns a tuple with the CancelExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelExisting

`func (o *ReservationSessionSerializerReadOnly) SetCancelExisting(v bool)`

SetCancelExisting sets CancelExisting field to given value.


### SetCancelExistingNil

`func (o *ReservationSessionSerializerReadOnly) SetCancelExistingNil(b bool)`

 SetCancelExistingNil sets the value for CancelExisting to be an explicit nil

### UnsetCancelExisting
`func (o *ReservationSessionSerializerReadOnly) UnsetCancelExisting()`

UnsetCancelExisting ensures that no value is present for CancelExisting, not even an explicit nil
### GetTRequested

`func (o *ReservationSessionSerializerReadOnly) GetTRequested() time.Time`

GetTRequested returns the TRequested field if non-nil, zero value otherwise.

### GetTRequestedOk

`func (o *ReservationSessionSerializerReadOnly) GetTRequestedOk() (*time.Time, bool)`

GetTRequestedOk returns a tuple with the TRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTRequested

`func (o *ReservationSessionSerializerReadOnly) SetTRequested(v time.Time)`

SetTRequested sets TRequested field to given value.


### GetTSatisfied

`func (o *ReservationSessionSerializerReadOnly) GetTSatisfied() time.Time`

GetTSatisfied returns the TSatisfied field if non-nil, zero value otherwise.

### GetTSatisfiedOk

`func (o *ReservationSessionSerializerReadOnly) GetTSatisfiedOk() (*time.Time, bool)`

GetTSatisfiedOk returns a tuple with the TSatisfied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSatisfied

`func (o *ReservationSessionSerializerReadOnly) SetTSatisfied(v time.Time)`

SetTSatisfied sets TSatisfied field to given value.


### SetTSatisfiedNil

`func (o *ReservationSessionSerializerReadOnly) SetTSatisfiedNil(b bool)`

 SetTSatisfiedNil sets the value for TSatisfied to be an explicit nil

### UnsetTSatisfied
`func (o *ReservationSessionSerializerReadOnly) UnsetTSatisfied()`

UnsetTSatisfied ensures that no value is present for TSatisfied, not even an explicit nil
### GetTCompleted

`func (o *ReservationSessionSerializerReadOnly) GetTCompleted() time.Time`

GetTCompleted returns the TCompleted field if non-nil, zero value otherwise.

### GetTCompletedOk

`func (o *ReservationSessionSerializerReadOnly) GetTCompletedOk() (*time.Time, bool)`

GetTCompletedOk returns a tuple with the TCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTCompleted

`func (o *ReservationSessionSerializerReadOnly) SetTCompleted(v time.Time)`

SetTCompleted sets TCompleted field to given value.


### SetTCompletedNil

`func (o *ReservationSessionSerializerReadOnly) SetTCompletedNil(b bool)`

 SetTCompletedNil sets the value for TCompleted to be an explicit nil

### UnsetTCompleted
`func (o *ReservationSessionSerializerReadOnly) UnsetTCompleted()`

UnsetTCompleted ensures that no value is present for TCompleted, not even an explicit nil
### GetTLeaseExpires

`func (o *ReservationSessionSerializerReadOnly) GetTLeaseExpires() time.Time`

GetTLeaseExpires returns the TLeaseExpires field if non-nil, zero value otherwise.

### GetTLeaseExpiresOk

`func (o *ReservationSessionSerializerReadOnly) GetTLeaseExpiresOk() (*time.Time, bool)`

GetTLeaseExpiresOk returns a tuple with the TLeaseExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLeaseExpires

`func (o *ReservationSessionSerializerReadOnly) SetTLeaseExpires(v time.Time)`

SetTLeaseExpires sets TLeaseExpires field to given value.


### SetTLeaseExpiresNil

`func (o *ReservationSessionSerializerReadOnly) SetTLeaseExpiresNil(b bool)`

 SetTLeaseExpiresNil sets the value for TLeaseExpires to be an explicit nil

### UnsetTLeaseExpires
`func (o *ReservationSessionSerializerReadOnly) UnsetTLeaseExpires()`

UnsetTLeaseExpires ensures that no value is present for TLeaseExpires, not even an explicit nil
### GetMetadata

`func (o *ReservationSessionSerializerReadOnly) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReservationSessionSerializerReadOnly) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReservationSessionSerializerReadOnly) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetRDevices

`func (o *ReservationSessionSerializerReadOnly) GetRDevices() []int32`

GetRDevices returns the RDevices field if non-nil, zero value otherwise.

### GetRDevicesOk

`func (o *ReservationSessionSerializerReadOnly) GetRDevicesOk() (*[]int32, bool)`

GetRDevicesOk returns a tuple with the RDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRDevices

`func (o *ReservationSessionSerializerReadOnly) SetRDevices(v []int32)`

SetRDevices sets RDevices field to given value.


### GetRDeviceGroups

`func (o *ReservationSessionSerializerReadOnly) GetRDeviceGroups() []int32`

GetRDeviceGroups returns the RDeviceGroups field if non-nil, zero value otherwise.

### GetRDeviceGroupsOk

`func (o *ReservationSessionSerializerReadOnly) GetRDeviceGroupsOk() (*[]int32, bool)`

GetRDeviceGroupsOk returns a tuple with the RDeviceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRDeviceGroups

`func (o *ReservationSessionSerializerReadOnly) SetRDeviceGroups(v []int32)`

SetRDeviceGroups sets RDeviceGroups field to given value.


### GetRDeviceGroupLabels

`func (o *ReservationSessionSerializerReadOnly) GetRDeviceGroupLabels() []int32`

GetRDeviceGroupLabels returns the RDeviceGroupLabels field if non-nil, zero value otherwise.

### GetRDeviceGroupLabelsOk

`func (o *ReservationSessionSerializerReadOnly) GetRDeviceGroupLabelsOk() (*[]int32, bool)`

GetRDeviceGroupLabelsOk returns a tuple with the RDeviceGroupLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRDeviceGroupLabels

`func (o *ReservationSessionSerializerReadOnly) SetRDeviceGroupLabels(v []int32)`

SetRDeviceGroupLabels sets RDeviceGroupLabels field to given value.


### GetADevices

`func (o *ReservationSessionSerializerReadOnly) GetADevices() []int32`

GetADevices returns the ADevices field if non-nil, zero value otherwise.

### GetADevicesOk

`func (o *ReservationSessionSerializerReadOnly) GetADevicesOk() (*[]int32, bool)`

GetADevicesOk returns a tuple with the ADevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADevices

`func (o *ReservationSessionSerializerReadOnly) SetADevices(v []int32)`

SetADevices sets ADevices field to given value.


### GetADeviceGroups

`func (o *ReservationSessionSerializerReadOnly) GetADeviceGroups() []int32`

GetADeviceGroups returns the ADeviceGroups field if non-nil, zero value otherwise.

### GetADeviceGroupsOk

`func (o *ReservationSessionSerializerReadOnly) GetADeviceGroupsOk() (*[]int32, bool)`

GetADeviceGroupsOk returns a tuple with the ADeviceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADeviceGroups

`func (o *ReservationSessionSerializerReadOnly) SetADeviceGroups(v []int32)`

SetADeviceGroups sets ADeviceGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


