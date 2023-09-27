# PatchedLabelSerializerWithPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**PermissionGroups** | Pointer to **[]string** |  | [optional] 
**DeviceGroups** | Pointer to **[]int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedLabelSerializerWithPermissions

`func NewPatchedLabelSerializerWithPermissions() *PatchedLabelSerializerWithPermissions`

NewPatchedLabelSerializerWithPermissions instantiates a new PatchedLabelSerializerWithPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLabelSerializerWithPermissionsWithDefaults

`func NewPatchedLabelSerializerWithPermissionsWithDefaults() *PatchedLabelSerializerWithPermissions`

NewPatchedLabelSerializerWithPermissionsWithDefaults instantiates a new PatchedLabelSerializerWithPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedLabelSerializerWithPermissions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedLabelSerializerWithPermissions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedLabelSerializerWithPermissions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedLabelSerializerWithPermissions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermissionGroups

`func (o *PatchedLabelSerializerWithPermissions) GetPermissionGroups() []string`

GetPermissionGroups returns the PermissionGroups field if non-nil, zero value otherwise.

### GetPermissionGroupsOk

`func (o *PatchedLabelSerializerWithPermissions) GetPermissionGroupsOk() (*[]string, bool)`

GetPermissionGroupsOk returns a tuple with the PermissionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroups

`func (o *PatchedLabelSerializerWithPermissions) SetPermissionGroups(v []string)`

SetPermissionGroups sets PermissionGroups field to given value.

### HasPermissionGroups

`func (o *PatchedLabelSerializerWithPermissions) HasPermissionGroups() bool`

HasPermissionGroups returns a boolean if a field has been set.

### SetPermissionGroupsNil

`func (o *PatchedLabelSerializerWithPermissions) SetPermissionGroupsNil(b bool)`

 SetPermissionGroupsNil sets the value for PermissionGroups to be an explicit nil

### UnsetPermissionGroups
`func (o *PatchedLabelSerializerWithPermissions) UnsetPermissionGroups()`

UnsetPermissionGroups ensures that no value is present for PermissionGroups, not even an explicit nil
### GetDeviceGroups

`func (o *PatchedLabelSerializerWithPermissions) GetDeviceGroups() []int32`

GetDeviceGroups returns the DeviceGroups field if non-nil, zero value otherwise.

### GetDeviceGroupsOk

`func (o *PatchedLabelSerializerWithPermissions) GetDeviceGroupsOk() (*[]int32, bool)`

GetDeviceGroupsOk returns a tuple with the DeviceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroups

`func (o *PatchedLabelSerializerWithPermissions) SetDeviceGroups(v []int32)`

SetDeviceGroups sets DeviceGroups field to given value.

### HasDeviceGroups

`func (o *PatchedLabelSerializerWithPermissions) HasDeviceGroups() bool`

HasDeviceGroups returns a boolean if a field has been set.

### GetName

`func (o *PatchedLabelSerializerWithPermissions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedLabelSerializerWithPermissions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedLabelSerializerWithPermissions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedLabelSerializerWithPermissions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchedLabelSerializerWithPermissions) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchedLabelSerializerWithPermissions) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchedLabelSerializerWithPermissions) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchedLabelSerializerWithPermissions) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PatchedLabelSerializerWithPermissions) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PatchedLabelSerializerWithPermissions) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


