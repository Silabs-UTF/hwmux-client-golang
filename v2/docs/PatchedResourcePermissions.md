# PatchedResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Permissions** | Pointer to [**[]PermissionsEnum**](PermissionsEnum.md) |  | [optional] [default to ["view","add","delete","change"]]

## Methods

### NewPatchedResourcePermissions

`func NewPatchedResourcePermissions() *PatchedResourcePermissions`

NewPatchedResourcePermissions instantiates a new PatchedResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedResourcePermissionsWithDefaults

`func NewPatchedResourcePermissionsWithDefaults() *PatchedResourcePermissions`

NewPatchedResourcePermissionsWithDefaults instantiates a new PatchedResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedResourcePermissions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedResourcePermissions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedResourcePermissions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedResourcePermissions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermissions

`func (o *PatchedResourcePermissions) GetPermissions() []PermissionsEnum`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PatchedResourcePermissions) GetPermissionsOk() (*[]PermissionsEnum, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PatchedResourcePermissions) SetPermissions(v []PermissionsEnum)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PatchedResourcePermissions) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


