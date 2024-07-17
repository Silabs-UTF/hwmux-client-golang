# PatchedDeviceGroupSerializerWithDevicePk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**PermissionGroups** | Pointer to **[]string** |  | [optional] 
**Devices** | Pointer to **[]int32** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**LocDesc** | Pointer to **[]string** |  | [optional] [readonly] 
**IsReserved** | Pointer to **NullableBool** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**EnableAhs** | Pointer to **bool** |  | [optional] 
**EnableAhsActions** | Pointer to **bool** |  | [optional] 
**EnableAhsCas** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Source** | Pointer to [**DeviceGroupSource**](DeviceGroupSource.md) |  | [optional] 

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

### SetPermissionGroupsNil

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetPermissionGroupsNil(b bool)`

 SetPermissionGroupsNil sets the value for PermissionGroups to be an explicit nil

### UnsetPermissionGroups
`func (o *PatchedDeviceGroupSerializerWithDevicePk) UnsetPermissionGroups()`

UnsetPermissionGroups ensures that no value is present for PermissionGroups, not even an explicit nil
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

### GetStatus

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLocDesc

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetLocDesc() []string`

GetLocDesc returns the LocDesc field if non-nil, zero value otherwise.

### GetLocDescOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetLocDescOk() (*[]string, bool)`

GetLocDescOk returns a tuple with the LocDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocDesc

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetLocDesc(v []string)`

SetLocDesc sets LocDesc field to given value.

### HasLocDesc

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasLocDesc() bool`

HasLocDesc returns a boolean if a field has been set.

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

### GetEnableAhs

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetEnableAhs() bool`

GetEnableAhs returns the EnableAhs field if non-nil, zero value otherwise.

### GetEnableAhsOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetEnableAhsOk() (*bool, bool)`

GetEnableAhsOk returns a tuple with the EnableAhs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAhs

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetEnableAhs(v bool)`

SetEnableAhs sets EnableAhs field to given value.

### HasEnableAhs

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasEnableAhs() bool`

HasEnableAhs returns a boolean if a field has been set.

### GetEnableAhsActions

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetEnableAhsActions() bool`

GetEnableAhsActions returns the EnableAhsActions field if non-nil, zero value otherwise.

### GetEnableAhsActionsOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetEnableAhsActionsOk() (*bool, bool)`

GetEnableAhsActionsOk returns a tuple with the EnableAhsActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAhsActions

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetEnableAhsActions(v bool)`

SetEnableAhsActions sets EnableAhsActions field to given value.

### HasEnableAhsActions

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasEnableAhsActions() bool`

HasEnableAhsActions returns a boolean if a field has been set.

### GetEnableAhsCas

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetEnableAhsCas() bool`

GetEnableAhsCas returns the EnableAhsCas field if non-nil, zero value otherwise.

### GetEnableAhsCasOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetEnableAhsCasOk() (*bool, bool)`

GetEnableAhsCasOk returns a tuple with the EnableAhsCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAhsCas

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetEnableAhsCas(v bool)`

SetEnableAhsCas sets EnableAhsCas field to given value.

### HasEnableAhsCas

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasEnableAhsCas() bool`

HasEnableAhsCas returns a boolean if a field has been set.

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
### GetSource

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetSource() DeviceGroupSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchedDeviceGroupSerializerWithDevicePk) GetSourceOk() (*DeviceGroupSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchedDeviceGroupSerializerWithDevicePk) SetSource(v DeviceGroupSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchedDeviceGroupSerializerWithDevicePk) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


