# PatchedDeviceGroupSerializerWithDevicePk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Devices** | Pointer to **[]int32** |  | [optional] 
**IsReserved** | Pointer to **NullableBool** |  | [optional] [readonly] 
**Online** | Pointer to **bool** |  | [optional] [readonly] 
**PermissionGroups** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedDeviceGroupSerializerWithDevicePk

`func NewPatchedDeviceGroupSerializerWithDevicePk() *PatchedDeviceGroupSerializerWithDevicePk`

NewPatchedDeviceGroupSerializerWithDevicePk instantiates a new PatchedDeviceGroupSerializerWithDevicePk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDeviceGroupSerializerWithDevicePkWithDefaults

`func NewPatchedDeviceGroupSerializerWithDevicePkWithDefaults() *PatchedDeviceGroupSerializerWithDevicePk`

NewPatchedDeviceGroupSerializerWithDevicePkWithDefaults instantiates a new PatchedDeviceGroupSerializerWithDevicePk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDevices

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetDevices() []int32`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetDevicesOk() (*[]int32, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetDevices(v []int32)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetIsReserved

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.

### HasIsReserved

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasIsReserved() bool`

HasIsReserved returns a boolean if a field has been set.

### SetIsReservedNil

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetIsReservedNil(b bool)`

 SetIsReservedNil sets the value for IsReserved to be an explicit nil

### UnsetIsReserved
`func (o *PatchedDeviceGroupSerializerWithDevicePk) UnsetIsReserved()`

UnsetIsReserved ensures that no value is present for IsReserved, not even an explicit nil
### GetOnline

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetPermissionGroups

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetPermissionGroups() []string`

GetPermissionGroups returns the PermissionGroups field if non-nil, zero value otherwise.

### GetPermissionGroupsOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetPermissionGroupsOk() (*[]string, bool)`

GetPermissionGroupsOk returns a tuple with the PermissionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroups

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetPermissionGroups(v []string)`

SetPermissionGroups sets PermissionGroups field to given value.

### HasPermissionGroups

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasPermissionGroups() bool`

HasPermissionGroups returns a boolean if a field has been set.

### GetName

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PatchedDeviceGroupSerializerWithDevicePk) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


