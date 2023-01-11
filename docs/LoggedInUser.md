# LoggedInUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**IsStaff** | Pointer to **bool** | Designates whether the user can log into this admin site. | [optional] 
**Groups** | **[]string** |  | [readonly] 
**IsSuperuser** | Pointer to **bool** | Designates that this user has all permissions without explicitly assigning them. | [optional] 
**Password** | **string** |  | 

## Methods

### NewLoggedInUser

`func NewLoggedInUser(id int32, username string, groups []string, password string, ) *LoggedInUser`

NewLoggedInUser instantiates a new LoggedInUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggedInUserWithDefaults

`func NewLoggedInUserWithDefaults() *LoggedInUser`

NewLoggedInUserWithDefaults instantiates a new LoggedInUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoggedInUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoggedInUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoggedInUser) SetId(v int32)`

SetId sets Id field to given value.


### GetUsername

`func (o *LoggedInUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LoggedInUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LoggedInUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *LoggedInUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LoggedInUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LoggedInUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LoggedInUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *LoggedInUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *LoggedInUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *LoggedInUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *LoggedInUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *LoggedInUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *LoggedInUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *LoggedInUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *LoggedInUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetIsStaff

`func (o *LoggedInUser) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *LoggedInUser) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *LoggedInUser) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *LoggedInUser) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetGroups

`func (o *LoggedInUser) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *LoggedInUser) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *LoggedInUser) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetIsSuperuser

`func (o *LoggedInUser) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *LoggedInUser) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *LoggedInUser) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *LoggedInUser) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetPassword

`func (o *LoggedInUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoggedInUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoggedInUser) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


