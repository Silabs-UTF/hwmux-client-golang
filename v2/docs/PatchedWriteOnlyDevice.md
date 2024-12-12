# PatchedWriteOnlyDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**PermissionGroups** | Pointer to **[]string** |  | [optional] 
**Part** | Pointer to **string** |  | [optional] 
**LocDesc** | Pointer to **string** |  | [optional] [readonly] 
**IsReserved** | Pointer to **NullableBool** |  | [optional] [readonly] 
**Location** | Pointer to [**LocationSerializerWriteOnly**](LocationSerializerWriteOnly.md) |  | [optional] 
**SnOrName** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to [**SourceEnum**](SourceEnum.md) |  | [optional] 
**SocketedChip** | Pointer to **string** |  | [optional] 
**IsWstk** | Pointer to **bool** |  | [optional] 
**Uri** | Pointer to **NullableString** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to [**StatusEnum**](StatusEnum.md) |  | [optional] 
**LastUpdate** | Pointer to **time.Time** |  | [optional] [readonly] 
**DateCreated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**IsReservedFilter** | Pointer to **bool** |  | [optional] 
**WstkPart** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchedWriteOnlyDevice

`func NewPatchedWriteOnlyDevice() *PatchedWriteOnlyDevice`

NewPatchedWriteOnlyDevice instantiates a new PatchedWriteOnlyDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWriteOnlyDeviceWithDefaults

`func NewPatchedWriteOnlyDeviceWithDefaults() *PatchedWriteOnlyDevice`

NewPatchedWriteOnlyDeviceWithDefaults instantiates a new PatchedWriteOnlyDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWriteOnlyDevice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWriteOnlyDevice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWriteOnlyDevice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWriteOnlyDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermissionGroups

`func (o *PatchedWriteOnlyDevice) GetPermissionGroups() []string`

GetPermissionGroups returns the PermissionGroups field if non-nil, zero value otherwise.

### GetPermissionGroupsOk

`func (o *PatchedWriteOnlyDevice) GetPermissionGroupsOk() (*[]string, bool)`

GetPermissionGroupsOk returns a tuple with the PermissionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroups

`func (o *PatchedWriteOnlyDevice) SetPermissionGroups(v []string)`

SetPermissionGroups sets PermissionGroups field to given value.

### HasPermissionGroups

`func (o *PatchedWriteOnlyDevice) HasPermissionGroups() bool`

HasPermissionGroups returns a boolean if a field has been set.

### GetPart

`func (o *PatchedWriteOnlyDevice) GetPart() string`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *PatchedWriteOnlyDevice) GetPartOk() (*string, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *PatchedWriteOnlyDevice) SetPart(v string)`

SetPart sets Part field to given value.

### HasPart

`func (o *PatchedWriteOnlyDevice) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetLocDesc

`func (o *PatchedWriteOnlyDevice) GetLocDesc() string`

GetLocDesc returns the LocDesc field if non-nil, zero value otherwise.

### GetLocDescOk

`func (o *PatchedWriteOnlyDevice) GetLocDescOk() (*string, bool)`

GetLocDescOk returns a tuple with the LocDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocDesc

`func (o *PatchedWriteOnlyDevice) SetLocDesc(v string)`

SetLocDesc sets LocDesc field to given value.

### HasLocDesc

`func (o *PatchedWriteOnlyDevice) HasLocDesc() bool`

HasLocDesc returns a boolean if a field has been set.

### GetIsReserved

`func (o *PatchedWriteOnlyDevice) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *PatchedWriteOnlyDevice) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *PatchedWriteOnlyDevice) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.

### HasIsReserved

`func (o *PatchedWriteOnlyDevice) HasIsReserved() bool`

HasIsReserved returns a boolean if a field has been set.

### SetIsReservedNil

`func (o *PatchedWriteOnlyDevice) SetIsReservedNil(b bool)`

 SetIsReservedNil sets the value for IsReserved to be an explicit nil

### UnsetIsReserved
`func (o *PatchedWriteOnlyDevice) UnsetIsReserved()`

UnsetIsReserved ensures that no value is present for IsReserved, not even an explicit nil
### GetLocation

`func (o *PatchedWriteOnlyDevice) GetLocation() LocationSerializerWriteOnly`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedWriteOnlyDevice) GetLocationOk() (*LocationSerializerWriteOnly, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedWriteOnlyDevice) SetLocation(v LocationSerializerWriteOnly)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedWriteOnlyDevice) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSnOrName

`func (o *PatchedWriteOnlyDevice) GetSnOrName() string`

GetSnOrName returns the SnOrName field if non-nil, zero value otherwise.

### GetSnOrNameOk

`func (o *PatchedWriteOnlyDevice) GetSnOrNameOk() (*string, bool)`

GetSnOrNameOk returns a tuple with the SnOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnOrName

`func (o *PatchedWriteOnlyDevice) SetSnOrName(v string)`

SetSnOrName sets SnOrName field to given value.

### HasSnOrName

`func (o *PatchedWriteOnlyDevice) HasSnOrName() bool`

HasSnOrName returns a boolean if a field has been set.

### SetSnOrNameNil

`func (o *PatchedWriteOnlyDevice) SetSnOrNameNil(b bool)`

 SetSnOrNameNil sets the value for SnOrName to be an explicit nil

### UnsetSnOrName
`func (o *PatchedWriteOnlyDevice) UnsetSnOrName()`

UnsetSnOrName ensures that no value is present for SnOrName, not even an explicit nil
### GetSource

`func (o *PatchedWriteOnlyDevice) GetSource() SourceEnum`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchedWriteOnlyDevice) GetSourceOk() (*SourceEnum, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchedWriteOnlyDevice) SetSource(v SourceEnum)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchedWriteOnlyDevice) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSocketedChip

`func (o *PatchedWriteOnlyDevice) GetSocketedChip() string`

GetSocketedChip returns the SocketedChip field if non-nil, zero value otherwise.

### GetSocketedChipOk

`func (o *PatchedWriteOnlyDevice) GetSocketedChipOk() (*string, bool)`

GetSocketedChipOk returns a tuple with the SocketedChip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketedChip

`func (o *PatchedWriteOnlyDevice) SetSocketedChip(v string)`

SetSocketedChip sets SocketedChip field to given value.

### HasSocketedChip

`func (o *PatchedWriteOnlyDevice) HasSocketedChip() bool`

HasSocketedChip returns a boolean if a field has been set.

### GetIsWstk

`func (o *PatchedWriteOnlyDevice) GetIsWstk() bool`

GetIsWstk returns the IsWstk field if non-nil, zero value otherwise.

### GetIsWstkOk

`func (o *PatchedWriteOnlyDevice) GetIsWstkOk() (*bool, bool)`

GetIsWstkOk returns a tuple with the IsWstk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWstk

`func (o *PatchedWriteOnlyDevice) SetIsWstk(v bool)`

SetIsWstk sets IsWstk field to given value.

### HasIsWstk

`func (o *PatchedWriteOnlyDevice) HasIsWstk() bool`

HasIsWstk returns a boolean if a field has been set.

### GetUri

`func (o *PatchedWriteOnlyDevice) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PatchedWriteOnlyDevice) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PatchedWriteOnlyDevice) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PatchedWriteOnlyDevice) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *PatchedWriteOnlyDevice) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *PatchedWriteOnlyDevice) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetOnline

`func (o *PatchedWriteOnlyDevice) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *PatchedWriteOnlyDevice) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *PatchedWriteOnlyDevice) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *PatchedWriteOnlyDevice) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWriteOnlyDevice) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWriteOnlyDevice) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWriteOnlyDevice) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWriteOnlyDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastUpdate

`func (o *PatchedWriteOnlyDevice) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *PatchedWriteOnlyDevice) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *PatchedWriteOnlyDevice) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *PatchedWriteOnlyDevice) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetDateCreated

`func (o *PatchedWriteOnlyDevice) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *PatchedWriteOnlyDevice) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *PatchedWriteOnlyDevice) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *PatchedWriteOnlyDevice) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchedWriteOnlyDevice) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchedWriteOnlyDevice) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchedWriteOnlyDevice) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchedWriteOnlyDevice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIsReservedFilter

`func (o *PatchedWriteOnlyDevice) GetIsReservedFilter() bool`

GetIsReservedFilter returns the IsReservedFilter field if non-nil, zero value otherwise.

### GetIsReservedFilterOk

`func (o *PatchedWriteOnlyDevice) GetIsReservedFilterOk() (*bool, bool)`

GetIsReservedFilterOk returns a tuple with the IsReservedFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReservedFilter

`func (o *PatchedWriteOnlyDevice) SetIsReservedFilter(v bool)`

SetIsReservedFilter sets IsReservedFilter field to given value.

### HasIsReservedFilter

`func (o *PatchedWriteOnlyDevice) HasIsReservedFilter() bool`

HasIsReservedFilter returns a boolean if a field has been set.

### GetWstkPart

`func (o *PatchedWriteOnlyDevice) GetWstkPart() string`

GetWstkPart returns the WstkPart field if non-nil, zero value otherwise.

### GetWstkPartOk

`func (o *PatchedWriteOnlyDevice) GetWstkPartOk() (*string, bool)`

GetWstkPartOk returns a tuple with the WstkPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWstkPart

`func (o *PatchedWriteOnlyDevice) SetWstkPart(v string)`

SetWstkPart sets WstkPart field to given value.

### HasWstkPart

`func (o *PatchedWriteOnlyDevice) HasWstkPart() bool`

HasWstkPart returns a boolean if a field has been set.

### SetWstkPartNil

`func (o *PatchedWriteOnlyDevice) SetWstkPartNil(b bool)`

 SetWstkPartNil sets the value for WstkPart to be an explicit nil

### UnsetWstkPart
`func (o *PatchedWriteOnlyDevice) UnsetWstkPart()`

UnsetWstkPart ensures that no value is present for WstkPart, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


