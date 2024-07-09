# DeviceSerializerPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**PermissionGroups** | Pointer to **[]string** |  | [optional] 
**Part** | [**Part**](Part.md) |  | 
**LocDesc** | **string** |  | [readonly] 
**IsReserved** | **NullableBool** |  | [readonly] 
**SnOrName** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to [**DeviceGroupSource**](DeviceGroupSource.md) |  | [optional] 
**IsWstk** | Pointer to **bool** |  | [optional] 
**Uri** | Pointer to **NullableString** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to [**StatusF39Enum**](StatusF39Enum.md) |  | [optional] 
**LastUpdate** | **time.Time** |  | [readonly] 
**DateCreated** | **time.Time** |  | [readonly] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**WstkPart** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDeviceSerializerPublic

`func NewDeviceSerializerPublic(id int32, part Part, locDesc string, isReserved NullableBool, lastUpdate time.Time, dateCreated time.Time, ) *DeviceSerializerPublic`

NewDeviceSerializerPublic instantiates a new DeviceSerializerPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceSerializerPublicWithDefaults

`func NewDeviceSerializerPublicWithDefaults() *DeviceSerializerPublic`

NewDeviceSerializerPublicWithDefaults instantiates a new DeviceSerializerPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceSerializerPublic) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceSerializerPublic) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceSerializerPublic) SetId(v int32)`

SetId sets Id field to given value.


### GetPermissionGroups

`func (o *DeviceSerializerPublic) GetPermissionGroups() []string`

GetPermissionGroups returns the PermissionGroups field if non-nil, zero value otherwise.

### GetPermissionGroupsOk

`func (o *DeviceSerializerPublic) GetPermissionGroupsOk() (*[]string, bool)`

GetPermissionGroupsOk returns a tuple with the PermissionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroups

`func (o *DeviceSerializerPublic) SetPermissionGroups(v []string)`

SetPermissionGroups sets PermissionGroups field to given value.

### HasPermissionGroups

`func (o *DeviceSerializerPublic) HasPermissionGroups() bool`

HasPermissionGroups returns a boolean if a field has been set.

### SetPermissionGroupsNil

`func (o *DeviceSerializerPublic) SetPermissionGroupsNil(b bool)`

 SetPermissionGroupsNil sets the value for PermissionGroups to be an explicit nil

### UnsetPermissionGroups
`func (o *DeviceSerializerPublic) UnsetPermissionGroups()`

UnsetPermissionGroups ensures that no value is present for PermissionGroups, not even an explicit nil
### GetPart

`func (o *DeviceSerializerPublic) GetPart() Part`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *DeviceSerializerPublic) GetPartOk() (*Part, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *DeviceSerializerPublic) SetPart(v Part)`

SetPart sets Part field to given value.


### GetLocDesc

`func (o *DeviceSerializerPublic) GetLocDesc() string`

GetLocDesc returns the LocDesc field if non-nil, zero value otherwise.

### GetLocDescOk

`func (o *DeviceSerializerPublic) GetLocDescOk() (*string, bool)`

GetLocDescOk returns a tuple with the LocDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocDesc

`func (o *DeviceSerializerPublic) SetLocDesc(v string)`

SetLocDesc sets LocDesc field to given value.


### GetIsReserved

`func (o *DeviceSerializerPublic) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *DeviceSerializerPublic) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *DeviceSerializerPublic) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.


### SetIsReservedNil

`func (o *DeviceSerializerPublic) SetIsReservedNil(b bool)`

 SetIsReservedNil sets the value for IsReserved to be an explicit nil

### UnsetIsReserved
`func (o *DeviceSerializerPublic) UnsetIsReserved()`

UnsetIsReserved ensures that no value is present for IsReserved, not even an explicit nil
### GetSnOrName

`func (o *DeviceSerializerPublic) GetSnOrName() string`

GetSnOrName returns the SnOrName field if non-nil, zero value otherwise.

### GetSnOrNameOk

`func (o *DeviceSerializerPublic) GetSnOrNameOk() (*string, bool)`

GetSnOrNameOk returns a tuple with the SnOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnOrName

`func (o *DeviceSerializerPublic) SetSnOrName(v string)`

SetSnOrName sets SnOrName field to given value.

### HasSnOrName

`func (o *DeviceSerializerPublic) HasSnOrName() bool`

HasSnOrName returns a boolean if a field has been set.

### SetSnOrNameNil

`func (o *DeviceSerializerPublic) SetSnOrNameNil(b bool)`

 SetSnOrNameNil sets the value for SnOrName to be an explicit nil

### UnsetSnOrName
`func (o *DeviceSerializerPublic) UnsetSnOrName()`

UnsetSnOrName ensures that no value is present for SnOrName, not even an explicit nil
### GetSource

`func (o *DeviceSerializerPublic) GetSource() DeviceGroupSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeviceSerializerPublic) GetSourceOk() (*DeviceGroupSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeviceSerializerPublic) SetSource(v DeviceGroupSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *DeviceSerializerPublic) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetIsWstk

`func (o *DeviceSerializerPublic) GetIsWstk() bool`

GetIsWstk returns the IsWstk field if non-nil, zero value otherwise.

### GetIsWstkOk

`func (o *DeviceSerializerPublic) GetIsWstkOk() (*bool, bool)`

GetIsWstkOk returns a tuple with the IsWstk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWstk

`func (o *DeviceSerializerPublic) SetIsWstk(v bool)`

SetIsWstk sets IsWstk field to given value.

### HasIsWstk

`func (o *DeviceSerializerPublic) HasIsWstk() bool`

HasIsWstk returns a boolean if a field has been set.

### GetUri

`func (o *DeviceSerializerPublic) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *DeviceSerializerPublic) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *DeviceSerializerPublic) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *DeviceSerializerPublic) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *DeviceSerializerPublic) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *DeviceSerializerPublic) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetOnline

`func (o *DeviceSerializerPublic) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *DeviceSerializerPublic) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *DeviceSerializerPublic) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *DeviceSerializerPublic) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceSerializerPublic) GetStatus() StatusF39Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceSerializerPublic) GetStatusOk() (*StatusF39Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceSerializerPublic) SetStatus(v StatusF39Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceSerializerPublic) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastUpdate

`func (o *DeviceSerializerPublic) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *DeviceSerializerPublic) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *DeviceSerializerPublic) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.


### GetDateCreated

`func (o *DeviceSerializerPublic) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *DeviceSerializerPublic) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *DeviceSerializerPublic) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetMetadata

`func (o *DeviceSerializerPublic) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeviceSerializerPublic) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeviceSerializerPublic) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeviceSerializerPublic) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *DeviceSerializerPublic) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DeviceSerializerPublic) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetWstkPart

`func (o *DeviceSerializerPublic) GetWstkPart() string`

GetWstkPart returns the WstkPart field if non-nil, zero value otherwise.

### GetWstkPartOk

`func (o *DeviceSerializerPublic) GetWstkPartOk() (*string, bool)`

GetWstkPartOk returns a tuple with the WstkPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWstkPart

`func (o *DeviceSerializerPublic) SetWstkPart(v string)`

SetWstkPart sets WstkPart field to given value.

### HasWstkPart

`func (o *DeviceSerializerPublic) HasWstkPart() bool`

HasWstkPart returns a boolean if a field has been set.

### SetWstkPartNil

`func (o *DeviceSerializerPublic) SetWstkPartNil(b bool)`

 SetWstkPartNil sets the value for WstkPart to be an explicit nil

### UnsetWstkPart
`func (o *DeviceSerializerPublic) UnsetWstkPart()`

UnsetWstkPart ensures that no value is present for WstkPart, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


