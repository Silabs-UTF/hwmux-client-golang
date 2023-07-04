# PatchedObjectPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to **map[string]interface{}** |  | [optional] 
**UserGroups** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedObjectPermissions

`func NewPatchedObjectPermissions() *PatchedObjectPermissions`

NewPatchedObjectPermissions instantiates a new PatchedObjectPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedObjectPermissionsWithDefaults

`func NewPatchedObjectPermissionsWithDefaults() *PatchedObjectPermissions`

NewPatchedObjectPermissionsWithDefaults instantiates a new PatchedObjectPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *PatchedObjectPermissions) GetUsers() map[string]interface{}`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PatchedObjectPermissions) GetUsersOk() (*map[string]interface{}, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PatchedObjectPermissions) SetUsers(v map[string]interface{})`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PatchedObjectPermissions) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetUserGroups

`func (o *PatchedObjectPermissions) GetUserGroups() map[string]interface{}`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *PatchedObjectPermissions) GetUserGroupsOk() (*map[string]interface{}, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *PatchedObjectPermissions) SetUserGroups(v map[string]interface{})`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *PatchedObjectPermissions) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


