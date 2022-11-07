# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Owner** | **string** |  | [readonly] 
**Device** | **NullableString** |  | [readonly] 
**DeviceGroup** | **NullableString** |  | [readonly] 
**Event** | [**EventEnum**](EventEnum.md) |  | 
**Datetime** | **time.Time** |  | [readonly] 
**Details** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**DeviceGroupLabel** | Pointer to **NullableInt32** |  | [optional] 
**CausedBy** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewLog

`func NewLog(id int32, owner string, device NullableString, deviceGroup NullableString, event EventEnum, datetime time.Time, ) *Log`

NewLog instantiates a new Log object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogWithDefaults

`func NewLogWithDefaults() *Log`

NewLogWithDefaults instantiates a new Log object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Log) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Log) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Log) SetId(v int32)`

SetId sets Id field to given value.


### GetOwner

`func (o *Log) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Log) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Log) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetDevice

`func (o *Log) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Log) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Log) SetDevice(v string)`

SetDevice sets Device field to given value.


### SetDeviceNil

`func (o *Log) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *Log) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetDeviceGroup

`func (o *Log) GetDeviceGroup() string`

GetDeviceGroup returns the DeviceGroup field if non-nil, zero value otherwise.

### GetDeviceGroupOk

`func (o *Log) GetDeviceGroupOk() (*string, bool)`

GetDeviceGroupOk returns a tuple with the DeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroup

`func (o *Log) SetDeviceGroup(v string)`

SetDeviceGroup sets DeviceGroup field to given value.


### SetDeviceGroupNil

`func (o *Log) SetDeviceGroupNil(b bool)`

 SetDeviceGroupNil sets the value for DeviceGroup to be an explicit nil

### UnsetDeviceGroup
`func (o *Log) UnsetDeviceGroup()`

UnsetDeviceGroup ensures that no value is present for DeviceGroup, not even an explicit nil
### GetEvent

`func (o *Log) GetEvent() EventEnum`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Log) GetEventOk() (*EventEnum, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Log) SetEvent(v EventEnum)`

SetEvent sets Event field to given value.


### GetDatetime

`func (o *Log) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *Log) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *Log) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### GetDetails

`func (o *Log) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Log) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Log) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Log) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMetadata

`func (o *Log) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Log) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Log) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Log) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDeviceGroupLabel

`func (o *Log) GetDeviceGroupLabel() int32`

GetDeviceGroupLabel returns the DeviceGroupLabel field if non-nil, zero value otherwise.

### GetDeviceGroupLabelOk

`func (o *Log) GetDeviceGroupLabelOk() (*int32, bool)`

GetDeviceGroupLabelOk returns a tuple with the DeviceGroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroupLabel

`func (o *Log) SetDeviceGroupLabel(v int32)`

SetDeviceGroupLabel sets DeviceGroupLabel field to given value.

### HasDeviceGroupLabel

`func (o *Log) HasDeviceGroupLabel() bool`

HasDeviceGroupLabel returns a boolean if a field has been set.

### SetDeviceGroupLabelNil

`func (o *Log) SetDeviceGroupLabelNil(b bool)`

 SetDeviceGroupLabelNil sets the value for DeviceGroupLabel to be an explicit nil

### UnsetDeviceGroupLabel
`func (o *Log) UnsetDeviceGroupLabel()`

UnsetDeviceGroupLabel ensures that no value is present for DeviceGroupLabel, not even an explicit nil
### GetCausedBy

`func (o *Log) GetCausedBy() int32`

GetCausedBy returns the CausedBy field if non-nil, zero value otherwise.

### GetCausedByOk

`func (o *Log) GetCausedByOk() (*int32, bool)`

GetCausedByOk returns a tuple with the CausedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCausedBy

`func (o *Log) SetCausedBy(v int32)`

SetCausedBy sets CausedBy field to given value.

### HasCausedBy

`func (o *Log) HasCausedBy() bool`

HasCausedBy returns a boolean if a field has been set.

### SetCausedByNil

`func (o *Log) SetCausedByNil(b bool)`

 SetCausedByNil sets the value for CausedBy to be an explicit nil

### UnsetCausedBy
`func (o *Log) UnsetCausedBy()`

UnsetCausedBy ensures that no value is present for CausedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


