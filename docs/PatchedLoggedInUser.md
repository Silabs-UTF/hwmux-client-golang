# PatchedLoggedInUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**IsStaff** | Pointer to **bool** | Designates whether the user can log into this admin site. | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] [readonly] 
**IsSuperuser** | Pointer to **bool** | Designates that this user has all permissions without explicitly assigning them. | [optional] 

## Methods

### NewPatchedLoggedInUser

`func NewPatchedLoggedInUser() *PatchedLoggedInUser`

NewPatchedLoggedInUser instantiates a new PatchedLoggedInUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLoggedInUserWithDefaults

`func NewPatchedLoggedInUserWithDefaults() *PatchedLoggedInUser`

NewPatchedLoggedInUserWithDefaults instantiates a new PatchedLoggedInUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *PatchedLoggedInUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedLoggedInUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedLoggedInUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedLoggedInUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *PatchedLoggedInUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedLoggedInUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedLoggedInUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedLoggedInUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *PatchedLoggedInUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PatchedLoggedInUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PatchedLoggedInUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PatchedLoggedInUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PatchedLoggedInUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PatchedLoggedInUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PatchedLoggedInUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PatchedLoggedInUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetIsStaff

`func (o *PatchedLoggedInUser) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *PatchedLoggedInUser) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *PatchedLoggedInUser) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *PatchedLoggedInUser) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetGroups

`func (o *PatchedLoggedInUser) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PatchedLoggedInUser) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PatchedLoggedInUser) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PatchedLoggedInUser) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIsSuperuser

`func (o *PatchedLoggedInUser) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *PatchedLoggedInUser) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *PatchedLoggedInUser) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *PatchedLoggedInUser) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


