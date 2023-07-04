# ObjectPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | **map[string]interface{}** |  | 
**UserGroups** | **map[string]interface{}** |  | 

## Methods

### NewObjectPermissions

`func NewObjectPermissions(users map[string]interface{}, userGroups map[string]interface{}, ) *ObjectPermissions`

NewObjectPermissions instantiates a new ObjectPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectPermissionsWithDefaults

`func NewObjectPermissionsWithDefaults() *ObjectPermissions`

NewObjectPermissionsWithDefaults instantiates a new ObjectPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ObjectPermissions) GetUsers() map[string]interface{}`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ObjectPermissions) GetUsersOk() (*map[string]interface{}, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ObjectPermissions) SetUsers(v map[string]interface{})`

SetUsers sets Users field to given value.


### GetUserGroups

`func (o *ObjectPermissions) GetUserGroups() map[string]interface{}`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *ObjectPermissions) GetUserGroupsOk() (*map[string]interface{}, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *ObjectPermissions) SetUserGroups(v map[string]interface{})`

SetUserGroups sets UserGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


