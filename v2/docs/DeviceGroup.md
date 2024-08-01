# DeviceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**PermissionGroups** | Pointer to **[]string** |  | [optional] 
**Devices** | [**[]LightDevice**](LightDevice.md) |  | [readonly] 
**Online** | **bool** |  | [readonly] 
**Status** | **string** |  | [readonly] 
**LocDesc** | **[]string** |  | [readonly] 
**IsReserved** | **NullableBool** |  | [readonly] 
**Name** | **string** |  | 
**EnableAhs** | Pointer to **bool** |  | [optional] 
**EnableAhsActions** | Pointer to **bool** |  | [optional] 
**EnableAhsCas** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Source** | Pointer to [**SourceEnum**](SourceEnum.md) |  | [optional] 

## Methods

### NewDeviceGroup

`func NewDeviceGroup(id int32, devices []LightDevice, online bool, status string, locDesc []string, isReserved NullableBool, name string, ) *DeviceGroup`

NewDeviceGroup instantiates a new DeviceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceGroupWithDefaults

`func NewDeviceGroupWithDefaults() *DeviceGroup`

NewDeviceGroupWithDefaults instantiates a new DeviceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetPermissionGroups

`func (o *DeviceGroup) GetPermissionGroups() []string`

GetPermissionGroups returns the PermissionGroups field if non-nil, zero value otherwise.

### GetPermissionGroupsOk

`func (o *DeviceGroup) GetPermissionGroupsOk() (*[]string, bool)`

GetPermissionGroupsOk returns a tuple with the PermissionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroups

`func (o *DeviceGroup) SetPermissionGroups(v []string)`

SetPermissionGroups sets PermissionGroups field to given value.

### HasPermissionGroups

`func (o *DeviceGroup) HasPermissionGroups() bool`

HasPermissionGroups returns a boolean if a field has been set.

### SetPermissionGroupsNil

`func (o *DeviceGroup) SetPermissionGroupsNil(b bool)`

 SetPermissionGroupsNil sets the value for PermissionGroups to be an explicit nil

### UnsetPermissionGroups
`func (o *DeviceGroup) UnsetPermissionGroups()`

UnsetPermissionGroups ensures that no value is present for PermissionGroups, not even an explicit nil
### GetDevices

`func (o *DeviceGroup) GetDevices() []LightDevice`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *DeviceGroup) GetDevicesOk() (*[]LightDevice, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *DeviceGroup) SetDevices(v []LightDevice)`

SetDevices sets Devices field to given value.


### GetOnline

`func (o *DeviceGroup) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *DeviceGroup) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *DeviceGroup) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStatus

`func (o *DeviceGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceGroup) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLocDesc

`func (o *DeviceGroup) GetLocDesc() []string`

GetLocDesc returns the LocDesc field if non-nil, zero value otherwise.

### GetLocDescOk

`func (o *DeviceGroup) GetLocDescOk() (*[]string, bool)`

GetLocDescOk returns a tuple with the LocDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocDesc

`func (o *DeviceGroup) SetLocDesc(v []string)`

SetLocDesc sets LocDesc field to given value.


### GetIsReserved

`func (o *DeviceGroup) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *DeviceGroup) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *DeviceGroup) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.


### SetIsReservedNil

`func (o *DeviceGroup) SetIsReservedNil(b bool)`

 SetIsReservedNil sets the value for IsReserved to be an explicit nil

### UnsetIsReserved
`func (o *DeviceGroup) UnsetIsReserved()`

UnsetIsReserved ensures that no value is present for IsReserved, not even an explicit nil
### GetName

`func (o *DeviceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceGroup) SetName(v string)`

SetName sets Name field to given value.


### GetEnableAhs

`func (o *DeviceGroup) GetEnableAhs() bool`

GetEnableAhs returns the EnableAhs field if non-nil, zero value otherwise.

### GetEnableAhsOk

`func (o *DeviceGroup) GetEnableAhsOk() (*bool, bool)`

GetEnableAhsOk returns a tuple with the EnableAhs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAhs

`func (o *DeviceGroup) SetEnableAhs(v bool)`

SetEnableAhs sets EnableAhs field to given value.

### HasEnableAhs

`func (o *DeviceGroup) HasEnableAhs() bool`

HasEnableAhs returns a boolean if a field has been set.

### GetEnableAhsActions

`func (o *DeviceGroup) GetEnableAhsActions() bool`

GetEnableAhsActions returns the EnableAhsActions field if non-nil, zero value otherwise.

### GetEnableAhsActionsOk

`func (o *DeviceGroup) GetEnableAhsActionsOk() (*bool, bool)`

GetEnableAhsActionsOk returns a tuple with the EnableAhsActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAhsActions

`func (o *DeviceGroup) SetEnableAhsActions(v bool)`

SetEnableAhsActions sets EnableAhsActions field to given value.

### HasEnableAhsActions

`func (o *DeviceGroup) HasEnableAhsActions() bool`

HasEnableAhsActions returns a boolean if a field has been set.

### GetEnableAhsCas

`func (o *DeviceGroup) GetEnableAhsCas() bool`

GetEnableAhsCas returns the EnableAhsCas field if non-nil, zero value otherwise.

### GetEnableAhsCasOk

`func (o *DeviceGroup) GetEnableAhsCasOk() (*bool, bool)`

GetEnableAhsCasOk returns a tuple with the EnableAhsCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAhsCas

`func (o *DeviceGroup) SetEnableAhsCas(v bool)`

SetEnableAhsCas sets EnableAhsCas field to given value.

### HasEnableAhsCas

`func (o *DeviceGroup) HasEnableAhsCas() bool`

HasEnableAhsCas returns a boolean if a field has been set.

### GetMetadata

`func (o *DeviceGroup) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeviceGroup) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeviceGroup) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeviceGroup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *DeviceGroup) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DeviceGroup) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetSource

`func (o *DeviceGroup) GetSource() SourceEnum`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeviceGroup) GetSourceOk() (*SourceEnum, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeviceGroup) SetSource(v SourceEnum)`

SetSource sets Source field to given value.

### HasSource

`func (o *DeviceGroup) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


