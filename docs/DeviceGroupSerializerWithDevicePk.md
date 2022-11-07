# DeviceGroupSerializerWithDevicePk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Devices** | Pointer to **[]int32** |  | [optional] 
**IsReserved** | **NullableBool** |  | [readonly] 
**PermissionGroups** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeviceGroupSerializerWithDevicePk

`func NewDeviceGroupSerializerWithDevicePk(id int32, isReserved NullableBool, name string, ) *DeviceGroupSerializerWithDevicePk`

NewDeviceGroupSerializerWithDevicePk instantiates a new DeviceGroupSerializerWithDevicePk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceGroupSerializerWithDevicePkWithDefaults

`func NewDeviceGroupSerializerWithDevicePkWithDefaults() *DeviceGroupSerializerWithDevicePk`

NewDeviceGroupSerializerWithDevicePkWithDefaults instantiates a new DeviceGroupSerializerWithDevicePk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceGroupSerializerWithDevicePk) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceGroupSerializerWithDevicePk) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceGroupSerializerWithDevicePk) SetId(v int32)`

SetId sets Id field to given value.


### GetDevices

`func (o *DeviceGroupSerializerWithDevicePk) GetDevices() []int32`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *DeviceGroupSerializerWithDevicePk) GetDevicesOk() (*[]int32, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *DeviceGroupSerializerWithDevicePk) SetDevices(v []int32)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *DeviceGroupSerializerWithDevicePk) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetIsReserved

`func (o *DeviceGroupSerializerWithDevicePk) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *DeviceGroupSerializerWithDevicePk) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *DeviceGroupSerializerWithDevicePk) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.


### SetIsReservedNil

`func (o *DeviceGroupSerializerWithDevicePk) SetIsReservedNil(b bool)`

 SetIsReservedNil sets the value for IsReserved to be an explicit nil

### UnsetIsReserved
`func (o *DeviceGroupSerializerWithDevicePk) UnsetIsReserved()`

UnsetIsReserved ensures that no value is present for IsReserved, not even an explicit nil
### GetPermissionGroups

`func (o *DeviceGroupSerializerWithDevicePk) GetPermissionGroups() []string`

GetPermissionGroups returns the PermissionGroups field if non-nil, zero value otherwise.

### GetPermissionGroupsOk

`func (o *DeviceGroupSerializerWithDevicePk) GetPermissionGroupsOk() (*[]string, bool)`

GetPermissionGroupsOk returns a tuple with the PermissionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroups

`func (o *DeviceGroupSerializerWithDevicePk) SetPermissionGroups(v []string)`

SetPermissionGroups sets PermissionGroups field to given value.

### HasPermissionGroups

`func (o *DeviceGroupSerializerWithDevicePk) HasPermissionGroups() bool`

HasPermissionGroups returns a boolean if a field has been set.

### GetName

`func (o *DeviceGroupSerializerWithDevicePk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceGroupSerializerWithDevicePk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceGroupSerializerWithDevicePk) SetName(v string)`

SetName sets Name field to given value.


### GetMetadata

`func (o *DeviceGroupSerializerWithDevicePk) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeviceGroupSerializerWithDevicePk) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeviceGroupSerializerWithDevicePk) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeviceGroupSerializerWithDevicePk) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *DeviceGroupSerializerWithDevicePk) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DeviceGroupSerializerWithDevicePk) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


