# DeviceGroupSerializerWithDevicePk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**PermissionGroups** | Pointer to **[]string** |  | [optional] 
**Devices** | Pointer to **[]int32** |  | [optional] 
**Online** | **bool** |  | [readonly] 
**Status** | **string** |  | [readonly] 
**LocDesc** | **[]string** |  | [readonly] 
**IsReserved** | **NullableBool** |  | [readonly] 
**Name** | **string** |  | 
**EnableAhs** | Pointer to **bool** |  | [optional] 
**EnableAhsActions** | Pointer to **bool** |  | [optional] 
**EnableAhsCas** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**Source** | Pointer to [**DeviceGroupSource**](DeviceGroupSource.md) |  | [optional] 

## Methods

### NewDeviceGroupSerializerWithDevicePk

`func NewDeviceGroupSerializerWithDevicePk(id int32, online bool, status string, locDesc []string, isReserved NullableBool, name string, ) *DeviceGroupSerializerWithDevicePk`

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

### SetPermissionGroupsNil

`func (o *DeviceGroupSerializerWithDevicePk) SetPermissionGroupsNil(b bool)`

 SetPermissionGroupsNil sets the value for PermissionGroups to be an explicit nil

### UnsetPermissionGroups
`func (o *DeviceGroupSerializerWithDevicePk) UnsetPermissionGroups()`

UnsetPermissionGroups ensures that no value is present for PermissionGroups, not even an explicit nil
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

### GetOnline

`func (o *DeviceGroupSerializerWithDevicePk) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *DeviceGroupSerializerWithDevicePk) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *DeviceGroupSerializerWithDevicePk) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStatus

`func (o *DeviceGroupSerializerWithDevicePk) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceGroupSerializerWithDevicePk) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceGroupSerializerWithDevicePk) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLocDesc

`func (o *DeviceGroupSerializerWithDevicePk) GetLocDesc() []string`

GetLocDesc returns the LocDesc field if non-nil, zero value otherwise.

### GetLocDescOk

`func (o *DeviceGroupSerializerWithDevicePk) GetLocDescOk() (*[]string, bool)`

GetLocDescOk returns a tuple with the LocDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocDesc

`func (o *DeviceGroupSerializerWithDevicePk) SetLocDesc(v []string)`

SetLocDesc sets LocDesc field to given value.


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


### GetEnableAhs

`func (o *DeviceGroupSerializerWithDevicePk) GetEnableAhs() bool`

GetEnableAhs returns the EnableAhs field if non-nil, zero value otherwise.

### GetEnableAhsOk

`func (o *DeviceGroupSerializerWithDevicePk) GetEnableAhsOk() (*bool, bool)`

GetEnableAhsOk returns a tuple with the EnableAhs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAhs

`func (o *DeviceGroupSerializerWithDevicePk) SetEnableAhs(v bool)`

SetEnableAhs sets EnableAhs field to given value.

### HasEnableAhs

`func (o *DeviceGroupSerializerWithDevicePk) HasEnableAhs() bool`

HasEnableAhs returns a boolean if a field has been set.

### GetEnableAhsActions

`func (o *DeviceGroupSerializerWithDevicePk) GetEnableAhsActions() bool`

GetEnableAhsActions returns the EnableAhsActions field if non-nil, zero value otherwise.

### GetEnableAhsActionsOk

`func (o *DeviceGroupSerializerWithDevicePk) GetEnableAhsActionsOk() (*bool, bool)`

GetEnableAhsActionsOk returns a tuple with the EnableAhsActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAhsActions

`func (o *DeviceGroupSerializerWithDevicePk) SetEnableAhsActions(v bool)`

SetEnableAhsActions sets EnableAhsActions field to given value.

### HasEnableAhsActions

`func (o *DeviceGroupSerializerWithDevicePk) HasEnableAhsActions() bool`

HasEnableAhsActions returns a boolean if a field has been set.

### GetEnableAhsCas

`func (o *DeviceGroupSerializerWithDevicePk) GetEnableAhsCas() bool`

GetEnableAhsCas returns the EnableAhsCas field if non-nil, zero value otherwise.

### GetEnableAhsCasOk

`func (o *DeviceGroupSerializerWithDevicePk) GetEnableAhsCasOk() (*bool, bool)`

GetEnableAhsCasOk returns a tuple with the EnableAhsCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAhsCas

`func (o *DeviceGroupSerializerWithDevicePk) SetEnableAhsCas(v bool)`

SetEnableAhsCas sets EnableAhsCas field to given value.

### HasEnableAhsCas

`func (o *DeviceGroupSerializerWithDevicePk) HasEnableAhsCas() bool`

HasEnableAhsCas returns a boolean if a field has been set.

### GetMetadata

`func (o *DeviceGroupSerializerWithDevicePk) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeviceGroupSerializerWithDevicePk) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeviceGroupSerializerWithDevicePk) SetMetadata(v interface{})`

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
### GetSource

`func (o *DeviceGroupSerializerWithDevicePk) GetSource() DeviceGroupSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeviceGroupSerializerWithDevicePk) GetSourceOk() (*DeviceGroupSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeviceGroupSerializerWithDevicePk) SetSource(v DeviceGroupSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *DeviceGroupSerializerWithDevicePk) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


