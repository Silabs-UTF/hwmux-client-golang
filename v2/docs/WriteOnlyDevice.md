# WriteOnlyDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**PermissionGroups** | Pointer to **[]string** |  | [optional] 
**Part** | **string** |  | 
**LocDesc** | **string** |  | [readonly] 
**IsReserved** | **NullableBool** |  | [readonly] 
**Location** | [**LocationSerializerWriteOnly**](LocationSerializerWriteOnly.md) |  | 
**SnOrName** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to [**DeviceGroupSource**](DeviceGroupSource.md) |  | [optional] 
**IsWstk** | Pointer to **bool** |  | [optional] 
**Uri** | Pointer to **NullableString** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to [**StatusEnum**](StatusEnum.md) |  | [optional] 
**LastUpdate** | **time.Time** |  | [readonly] 
**DateCreated** | **time.Time** |  | [readonly] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**WstkPart** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWriteOnlyDevice

`func NewWriteOnlyDevice(id int32, part string, locDesc string, isReserved NullableBool, location LocationSerializerWriteOnly, lastUpdate time.Time, dateCreated time.Time, ) *WriteOnlyDevice`

NewWriteOnlyDevice instantiates a new WriteOnlyDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteOnlyDeviceWithDefaults

`func NewWriteOnlyDeviceWithDefaults() *WriteOnlyDevice`

NewWriteOnlyDeviceWithDefaults instantiates a new WriteOnlyDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WriteOnlyDevice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WriteOnlyDevice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WriteOnlyDevice) SetId(v int32)`

SetId sets Id field to given value.


### GetPermissionGroups

`func (o *WriteOnlyDevice) GetPermissionGroups() []string`

GetPermissionGroups returns the PermissionGroups field if non-nil, zero value otherwise.

### GetPermissionGroupsOk

`func (o *WriteOnlyDevice) GetPermissionGroupsOk() (*[]string, bool)`

GetPermissionGroupsOk returns a tuple with the PermissionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroups

`func (o *WriteOnlyDevice) SetPermissionGroups(v []string)`

SetPermissionGroups sets PermissionGroups field to given value.

### HasPermissionGroups

`func (o *WriteOnlyDevice) HasPermissionGroups() bool`

HasPermissionGroups returns a boolean if a field has been set.

### GetPart

`func (o *WriteOnlyDevice) GetPart() string`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *WriteOnlyDevice) GetPartOk() (*string, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *WriteOnlyDevice) SetPart(v string)`

SetPart sets Part field to given value.


### GetLocDesc

`func (o *WriteOnlyDevice) GetLocDesc() string`

GetLocDesc returns the LocDesc field if non-nil, zero value otherwise.

### GetLocDescOk

`func (o *WriteOnlyDevice) GetLocDescOk() (*string, bool)`

GetLocDescOk returns a tuple with the LocDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocDesc

`func (o *WriteOnlyDevice) SetLocDesc(v string)`

SetLocDesc sets LocDesc field to given value.


### GetIsReserved

`func (o *WriteOnlyDevice) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *WriteOnlyDevice) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *WriteOnlyDevice) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.


### SetIsReservedNil

`func (o *WriteOnlyDevice) SetIsReservedNil(b bool)`

 SetIsReservedNil sets the value for IsReserved to be an explicit nil

### UnsetIsReserved
`func (o *WriteOnlyDevice) UnsetIsReserved()`

UnsetIsReserved ensures that no value is present for IsReserved, not even an explicit nil
### GetLocation

`func (o *WriteOnlyDevice) GetLocation() LocationSerializerWriteOnly`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WriteOnlyDevice) GetLocationOk() (*LocationSerializerWriteOnly, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WriteOnlyDevice) SetLocation(v LocationSerializerWriteOnly)`

SetLocation sets Location field to given value.


### GetSnOrName

`func (o *WriteOnlyDevice) GetSnOrName() string`

GetSnOrName returns the SnOrName field if non-nil, zero value otherwise.

### GetSnOrNameOk

`func (o *WriteOnlyDevice) GetSnOrNameOk() (*string, bool)`

GetSnOrNameOk returns a tuple with the SnOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnOrName

`func (o *WriteOnlyDevice) SetSnOrName(v string)`

SetSnOrName sets SnOrName field to given value.

### HasSnOrName

`func (o *WriteOnlyDevice) HasSnOrName() bool`

HasSnOrName returns a boolean if a field has been set.

### SetSnOrNameNil

`func (o *WriteOnlyDevice) SetSnOrNameNil(b bool)`

 SetSnOrNameNil sets the value for SnOrName to be an explicit nil

### UnsetSnOrName
`func (o *WriteOnlyDevice) UnsetSnOrName()`

UnsetSnOrName ensures that no value is present for SnOrName, not even an explicit nil
### GetSource

`func (o *WriteOnlyDevice) GetSource() DeviceGroupSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WriteOnlyDevice) GetSourceOk() (*DeviceGroupSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WriteOnlyDevice) SetSource(v DeviceGroupSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *WriteOnlyDevice) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetIsWstk

`func (o *WriteOnlyDevice) GetIsWstk() bool`

GetIsWstk returns the IsWstk field if non-nil, zero value otherwise.

### GetIsWstkOk

`func (o *WriteOnlyDevice) GetIsWstkOk() (*bool, bool)`

GetIsWstkOk returns a tuple with the IsWstk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWstk

`func (o *WriteOnlyDevice) SetIsWstk(v bool)`

SetIsWstk sets IsWstk field to given value.

### HasIsWstk

`func (o *WriteOnlyDevice) HasIsWstk() bool`

HasIsWstk returns a boolean if a field has been set.

### GetUri

`func (o *WriteOnlyDevice) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *WriteOnlyDevice) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *WriteOnlyDevice) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *WriteOnlyDevice) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *WriteOnlyDevice) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *WriteOnlyDevice) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetOnline

`func (o *WriteOnlyDevice) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *WriteOnlyDevice) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *WriteOnlyDevice) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *WriteOnlyDevice) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetStatus

`func (o *WriteOnlyDevice) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WriteOnlyDevice) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WriteOnlyDevice) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WriteOnlyDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastUpdate

`func (o *WriteOnlyDevice) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *WriteOnlyDevice) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *WriteOnlyDevice) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.


### GetDateCreated

`func (o *WriteOnlyDevice) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *WriteOnlyDevice) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *WriteOnlyDevice) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetMetadata

`func (o *WriteOnlyDevice) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WriteOnlyDevice) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WriteOnlyDevice) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WriteOnlyDevice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WriteOnlyDevice) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WriteOnlyDevice) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetWstkPart

`func (o *WriteOnlyDevice) GetWstkPart() string`

GetWstkPart returns the WstkPart field if non-nil, zero value otherwise.

### GetWstkPartOk

`func (o *WriteOnlyDevice) GetWstkPartOk() (*string, bool)`

GetWstkPartOk returns a tuple with the WstkPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWstkPart

`func (o *WriteOnlyDevice) SetWstkPart(v string)`

SetWstkPart sets WstkPart field to given value.

### HasWstkPart

`func (o *WriteOnlyDevice) HasWstkPart() bool`

HasWstkPart returns a boolean if a field has been set.

### SetWstkPartNil

`func (o *WriteOnlyDevice) SetWstkPartNil(b bool)`

 SetWstkPartNil sets the value for WstkPart to be an explicit nil

### UnsetWstkPart
`func (o *WriteOnlyDevice) UnsetWstkPart()`

UnsetWstkPart ensures that no value is present for WstkPart, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


