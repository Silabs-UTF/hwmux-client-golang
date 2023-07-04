# ReservationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **string** |  | [optional] 
**RDevices** | Pointer to **[]int32** |  | [optional] 
**RDeviceGroups** | Pointer to **[]int32** |  | [optional] 
**RDeviceGroupLabels** | Pointer to **[]int32** |  | [optional] 
**UseWatchdog** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**InvertPriority** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewReservationRequest

`func NewReservationRequest() *ReservationRequest`

NewReservationRequest instantiates a new ReservationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationRequestWithDefaults

`func NewReservationRequestWithDefaults() *ReservationRequest`

NewReservationRequestWithDefaults instantiates a new ReservationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ReservationRequest) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ReservationRequest) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ReservationRequest) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ReservationRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRDevices

`func (o *ReservationRequest) GetRDevices() []int32`

GetRDevices returns the RDevices field if non-nil, zero value otherwise.

### GetRDevicesOk

`func (o *ReservationRequest) GetRDevicesOk() (*[]int32, bool)`

GetRDevicesOk returns a tuple with the RDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRDevices

`func (o *ReservationRequest) SetRDevices(v []int32)`

SetRDevices sets RDevices field to given value.

### HasRDevices

`func (o *ReservationRequest) HasRDevices() bool`

HasRDevices returns a boolean if a field has been set.

### GetRDeviceGroups

`func (o *ReservationRequest) GetRDeviceGroups() []int32`

GetRDeviceGroups returns the RDeviceGroups field if non-nil, zero value otherwise.

### GetRDeviceGroupsOk

`func (o *ReservationRequest) GetRDeviceGroupsOk() (*[]int32, bool)`

GetRDeviceGroupsOk returns a tuple with the RDeviceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRDeviceGroups

`func (o *ReservationRequest) SetRDeviceGroups(v []int32)`

SetRDeviceGroups sets RDeviceGroups field to given value.

### HasRDeviceGroups

`func (o *ReservationRequest) HasRDeviceGroups() bool`

HasRDeviceGroups returns a boolean if a field has been set.

### GetRDeviceGroupLabels

`func (o *ReservationRequest) GetRDeviceGroupLabels() []int32`

GetRDeviceGroupLabels returns the RDeviceGroupLabels field if non-nil, zero value otherwise.

### GetRDeviceGroupLabelsOk

`func (o *ReservationRequest) GetRDeviceGroupLabelsOk() (*[]int32, bool)`

GetRDeviceGroupLabelsOk returns a tuple with the RDeviceGroupLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRDeviceGroupLabels

`func (o *ReservationRequest) SetRDeviceGroupLabels(v []int32)`

SetRDeviceGroupLabels sets RDeviceGroupLabels field to given value.

### HasRDeviceGroupLabels

`func (o *ReservationRequest) HasRDeviceGroupLabels() bool`

HasRDeviceGroupLabels returns a boolean if a field has been set.

### GetUseWatchdog

`func (o *ReservationRequest) GetUseWatchdog() bool`

GetUseWatchdog returns the UseWatchdog field if non-nil, zero value otherwise.

### GetUseWatchdogOk

`func (o *ReservationRequest) GetUseWatchdogOk() (*bool, bool)`

GetUseWatchdogOk returns a tuple with the UseWatchdog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWatchdog

`func (o *ReservationRequest) SetUseWatchdog(v bool)`

SetUseWatchdog sets UseWatchdog field to given value.

### HasUseWatchdog

`func (o *ReservationRequest) HasUseWatchdog() bool`

HasUseWatchdog returns a boolean if a field has been set.

### GetMetadata

`func (o *ReservationRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReservationRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReservationRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ReservationRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetInvertPriority

`func (o *ReservationRequest) GetInvertPriority() bool`

GetInvertPriority returns the InvertPriority field if non-nil, zero value otherwise.

### GetInvertPriorityOk

`func (o *ReservationRequest) GetInvertPriorityOk() (*bool, bool)`

GetInvertPriorityOk returns a tuple with the InvertPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertPriority

`func (o *ReservationRequest) SetInvertPriority(v bool)`

SetInvertPriority sets InvertPriority field to given value.

### HasInvertPriority

`func (o *ReservationRequest) HasInvertPriority() bool`

HasInvertPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


