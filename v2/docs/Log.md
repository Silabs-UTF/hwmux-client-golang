# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Owner** | **string** |  | [readonly] 
**Event** | [**EventEnum**](EventEnum.md) |  | 
**Datetime** | **time.Time** |  | [readonly] 
**ResourceName** | Pointer to **NullableString** |  | [optional] 
**ResourceNameList** | Pointer to **[]string** |  | [optional] 
**ResourceType** | Pointer to [**NullableLogResourceType**](LogResourceType.md) |  | [optional] 
**Status** | Pointer to [**LogStatusEnum**](LogStatusEnum.md) |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CausedBy** | Pointer to **NullableInt32** |  | [optional] 
**Device** | Pointer to **[]int32** |  | [optional] 
**DeviceGroup** | Pointer to **[]int32** |  | [optional] 
**DeviceGroupLabel** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewLog

`func NewLog(id int32, owner string, event EventEnum, datetime time.Time, ) *Log`

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


### GetResourceName

`func (o *Log) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Log) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Log) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *Log) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *Log) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *Log) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceNameList

`func (o *Log) GetResourceNameList() []string`

GetResourceNameList returns the ResourceNameList field if non-nil, zero value otherwise.

### GetResourceNameListOk

`func (o *Log) GetResourceNameListOk() (*[]string, bool)`

GetResourceNameListOk returns a tuple with the ResourceNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceNameList

`func (o *Log) SetResourceNameList(v []string)`

SetResourceNameList sets ResourceNameList field to given value.

### HasResourceNameList

`func (o *Log) HasResourceNameList() bool`

HasResourceNameList returns a boolean if a field has been set.

### GetResourceType

`func (o *Log) GetResourceType() LogResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Log) GetResourceTypeOk() (*LogResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Log) SetResourceType(v LogResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Log) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *Log) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *Log) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetStatus

`func (o *Log) GetStatus() LogStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Log) GetStatusOk() (*LogStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Log) SetStatus(v LogStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Log) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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
### GetDevice

`func (o *Log) GetDevice() []int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Log) GetDeviceOk() (*[]int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Log) SetDevice(v []int32)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *Log) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDeviceGroup

`func (o *Log) GetDeviceGroup() []int32`

GetDeviceGroup returns the DeviceGroup field if non-nil, zero value otherwise.

### GetDeviceGroupOk

`func (o *Log) GetDeviceGroupOk() (*[]int32, bool)`

GetDeviceGroupOk returns a tuple with the DeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroup

`func (o *Log) SetDeviceGroup(v []int32)`

SetDeviceGroup sets DeviceGroup field to given value.

### HasDeviceGroup

`func (o *Log) HasDeviceGroup() bool`

HasDeviceGroup returns a boolean if a field has been set.

### GetDeviceGroupLabel

`func (o *Log) GetDeviceGroupLabel() []int32`

GetDeviceGroupLabel returns the DeviceGroupLabel field if non-nil, zero value otherwise.

### GetDeviceGroupLabelOk

`func (o *Log) GetDeviceGroupLabelOk() (*[]int32, bool)`

GetDeviceGroupLabelOk returns a tuple with the DeviceGroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroupLabel

`func (o *Log) SetDeviceGroupLabel(v []int32)`

SetDeviceGroupLabel sets DeviceGroupLabel field to given value.

### HasDeviceGroupLabel

`func (o *Log) HasDeviceGroupLabel() bool`

HasDeviceGroupLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


