# LabelSerializerWithPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**PermissionGroups** | Pointer to **[]string** |  | [optional] 
**DeviceGroups** | **[]int32** |  | 
**Name** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Source** | Pointer to [**DeviceGroupSource**](DeviceGroupSource.md) |  | [optional] 

## Methods

### NewLabelSerializerWithPermissions

`func NewLabelSerializerWithPermissions(id int32, deviceGroups []int32, name string, ) *LabelSerializerWithPermissions`

NewLabelSerializerWithPermissions instantiates a new LabelSerializerWithPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelSerializerWithPermissionsWithDefaults

`func NewLabelSerializerWithPermissionsWithDefaults() *LabelSerializerWithPermissions`

NewLabelSerializerWithPermissionsWithDefaults instantiates a new LabelSerializerWithPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LabelSerializerWithPermissions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabelSerializerWithPermissions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabelSerializerWithPermissions) SetId(v int32)`

SetId sets Id field to given value.


### GetPermissionGroups

`func (o *LabelSerializerWithPermissions) GetPermissionGroups() []string`

GetPermissionGroups returns the PermissionGroups field if non-nil, zero value otherwise.

### GetPermissionGroupsOk

`func (o *LabelSerializerWithPermissions) GetPermissionGroupsOk() (*[]string, bool)`

GetPermissionGroupsOk returns a tuple with the PermissionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroups

`func (o *LabelSerializerWithPermissions) SetPermissionGroups(v []string)`

SetPermissionGroups sets PermissionGroups field to given value.

### HasPermissionGroups

`func (o *LabelSerializerWithPermissions) HasPermissionGroups() bool`

HasPermissionGroups returns a boolean if a field has been set.

### SetPermissionGroupsNil

`func (o *LabelSerializerWithPermissions) SetPermissionGroupsNil(b bool)`

 SetPermissionGroupsNil sets the value for PermissionGroups to be an explicit nil

### UnsetPermissionGroups
`func (o *LabelSerializerWithPermissions) UnsetPermissionGroups()`

UnsetPermissionGroups ensures that no value is present for PermissionGroups, not even an explicit nil
### GetDeviceGroups

`func (o *LabelSerializerWithPermissions) GetDeviceGroups() []int32`

GetDeviceGroups returns the DeviceGroups field if non-nil, zero value otherwise.

### GetDeviceGroupsOk

`func (o *LabelSerializerWithPermissions) GetDeviceGroupsOk() (*[]int32, bool)`

GetDeviceGroupsOk returns a tuple with the DeviceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroups

`func (o *LabelSerializerWithPermissions) SetDeviceGroups(v []int32)`

SetDeviceGroups sets DeviceGroups field to given value.


### GetName

`func (o *LabelSerializerWithPermissions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LabelSerializerWithPermissions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LabelSerializerWithPermissions) SetName(v string)`

SetName sets Name field to given value.


### GetMetadata

`func (o *LabelSerializerWithPermissions) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LabelSerializerWithPermissions) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LabelSerializerWithPermissions) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LabelSerializerWithPermissions) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *LabelSerializerWithPermissions) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *LabelSerializerWithPermissions) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetSource

`func (o *LabelSerializerWithPermissions) GetSource() DeviceGroupSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LabelSerializerWithPermissions) GetSourceOk() (*DeviceGroupSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LabelSerializerWithPermissions) SetSource(v DeviceGroupSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *LabelSerializerWithPermissions) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


