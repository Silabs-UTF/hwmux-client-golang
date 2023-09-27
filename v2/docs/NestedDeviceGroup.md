# NestedDeviceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Devices** | Pointer to [**[]WriteOnlyDevice**](WriteOnlyDevice.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**PermissionGroups** | Pointer to **[]string** |  | [optional] 
**EnableAhs** | Pointer to **bool** |  | [optional] 
**EnableAhsActions** | Pointer to **bool** |  | [optional] 

## Methods

### NewNestedDeviceGroup

`func NewNestedDeviceGroup(name string, ) *NestedDeviceGroup`

NewNestedDeviceGroup instantiates a new NestedDeviceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedDeviceGroupWithDefaults

`func NewNestedDeviceGroupWithDefaults() *NestedDeviceGroup`

NewNestedDeviceGroupWithDefaults instantiates a new NestedDeviceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NestedDeviceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedDeviceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedDeviceGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDevices

`func (o *NestedDeviceGroup) GetDevices() []WriteOnlyDevice`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *NestedDeviceGroup) GetDevicesOk() (*[]WriteOnlyDevice, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *NestedDeviceGroup) SetDevices(v []WriteOnlyDevice)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *NestedDeviceGroup) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetMetadata

`func (o *NestedDeviceGroup) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NestedDeviceGroup) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NestedDeviceGroup) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NestedDeviceGroup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *NestedDeviceGroup) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *NestedDeviceGroup) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPermissionGroups

`func (o *NestedDeviceGroup) GetPermissionGroups() []string`

GetPermissionGroups returns the PermissionGroups field if non-nil, zero value otherwise.

### GetPermissionGroupsOk

`func (o *NestedDeviceGroup) GetPermissionGroupsOk() (*[]string, bool)`

GetPermissionGroupsOk returns a tuple with the PermissionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroups

`func (o *NestedDeviceGroup) SetPermissionGroups(v []string)`

SetPermissionGroups sets PermissionGroups field to given value.

### HasPermissionGroups

`func (o *NestedDeviceGroup) HasPermissionGroups() bool`

HasPermissionGroups returns a boolean if a field has been set.

### GetEnableAhs

`func (o *NestedDeviceGroup) GetEnableAhs() bool`

GetEnableAhs returns the EnableAhs field if non-nil, zero value otherwise.

### GetEnableAhsOk

`func (o *NestedDeviceGroup) GetEnableAhsOk() (*bool, bool)`

GetEnableAhsOk returns a tuple with the EnableAhs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAhs

`func (o *NestedDeviceGroup) SetEnableAhs(v bool)`

SetEnableAhs sets EnableAhs field to given value.

### HasEnableAhs

`func (o *NestedDeviceGroup) HasEnableAhs() bool`

HasEnableAhs returns a boolean if a field has been set.

### GetEnableAhsActions

`func (o *NestedDeviceGroup) GetEnableAhsActions() bool`

GetEnableAhsActions returns the EnableAhsActions field if non-nil, zero value otherwise.

### GetEnableAhsActionsOk

`func (o *NestedDeviceGroup) GetEnableAhsActionsOk() (*bool, bool)`

GetEnableAhsActionsOk returns a tuple with the EnableAhsActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAhsActions

`func (o *NestedDeviceGroup) SetEnableAhsActions(v bool)`

SetEnableAhsActions sets EnableAhsActions field to given value.

### HasEnableAhsActions

`func (o *NestedDeviceGroup) HasEnableAhsActions() bool`

HasEnableAhsActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


