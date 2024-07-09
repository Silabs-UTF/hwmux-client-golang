# LightDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**SnOrName** | Pointer to **NullableString** |  | [optional] 
**Uri** | Pointer to **NullableString** |  | [optional] 
**IsWstk** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Part** | [**Part**](Part.md) |  | 
**Location** | **int32** |  | 
**WstkPart** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**StatusF39Enum**](StatusF39Enum.md) |  | [optional] 
**LocDesc** | **string** |  | [readonly] 

## Methods

### NewLightDevice

`func NewLightDevice(id int32, part Part, location int32, locDesc string, ) *LightDevice`

NewLightDevice instantiates a new LightDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLightDeviceWithDefaults

`func NewLightDeviceWithDefaults() *LightDevice`

NewLightDeviceWithDefaults instantiates a new LightDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LightDevice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LightDevice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LightDevice) SetId(v int32)`

SetId sets Id field to given value.


### GetSnOrName

`func (o *LightDevice) GetSnOrName() string`

GetSnOrName returns the SnOrName field if non-nil, zero value otherwise.

### GetSnOrNameOk

`func (o *LightDevice) GetSnOrNameOk() (*string, bool)`

GetSnOrNameOk returns a tuple with the SnOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnOrName

`func (o *LightDevice) SetSnOrName(v string)`

SetSnOrName sets SnOrName field to given value.

### HasSnOrName

`func (o *LightDevice) HasSnOrName() bool`

HasSnOrName returns a boolean if a field has been set.

### SetSnOrNameNil

`func (o *LightDevice) SetSnOrNameNil(b bool)`

 SetSnOrNameNil sets the value for SnOrName to be an explicit nil

### UnsetSnOrName
`func (o *LightDevice) UnsetSnOrName()`

UnsetSnOrName ensures that no value is present for SnOrName, not even an explicit nil
### GetUri

`func (o *LightDevice) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *LightDevice) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *LightDevice) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *LightDevice) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *LightDevice) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *LightDevice) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetIsWstk

`func (o *LightDevice) GetIsWstk() bool`

GetIsWstk returns the IsWstk field if non-nil, zero value otherwise.

### GetIsWstkOk

`func (o *LightDevice) GetIsWstkOk() (*bool, bool)`

GetIsWstkOk returns a tuple with the IsWstk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWstk

`func (o *LightDevice) SetIsWstk(v bool)`

SetIsWstk sets IsWstk field to given value.

### HasIsWstk

`func (o *LightDevice) HasIsWstk() bool`

HasIsWstk returns a boolean if a field has been set.

### GetMetadata

`func (o *LightDevice) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LightDevice) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LightDevice) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LightDevice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *LightDevice) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *LightDevice) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOnline

`func (o *LightDevice) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *LightDevice) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *LightDevice) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *LightDevice) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetPart

`func (o *LightDevice) GetPart() Part`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *LightDevice) GetPartOk() (*Part, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *LightDevice) SetPart(v Part)`

SetPart sets Part field to given value.


### GetLocation

`func (o *LightDevice) GetLocation() int32`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LightDevice) GetLocationOk() (*int32, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LightDevice) SetLocation(v int32)`

SetLocation sets Location field to given value.


### GetWstkPart

`func (o *LightDevice) GetWstkPart() string`

GetWstkPart returns the WstkPart field if non-nil, zero value otherwise.

### GetWstkPartOk

`func (o *LightDevice) GetWstkPartOk() (*string, bool)`

GetWstkPartOk returns a tuple with the WstkPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWstkPart

`func (o *LightDevice) SetWstkPart(v string)`

SetWstkPart sets WstkPart field to given value.

### HasWstkPart

`func (o *LightDevice) HasWstkPart() bool`

HasWstkPart returns a boolean if a field has been set.

### SetWstkPartNil

`func (o *LightDevice) SetWstkPartNil(b bool)`

 SetWstkPartNil sets the value for WstkPart to be an explicit nil

### UnsetWstkPart
`func (o *LightDevice) UnsetWstkPart()`

UnsetWstkPart ensures that no value is present for WstkPart, not even an explicit nil
### GetStatus

`func (o *LightDevice) GetStatus() StatusF39Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LightDevice) GetStatusOk() (*StatusF39Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LightDevice) SetStatus(v StatusF39Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LightDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLocDesc

`func (o *LightDevice) GetLocDesc() string`

GetLocDesc returns the LocDesc field if non-nil, zero value otherwise.

### GetLocDescOk

`func (o *LightDevice) GetLocDescOk() (*string, bool)`

GetLocDescOk returns a tuple with the LocDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocDesc

`func (o *LightDevice) SetLocDesc(v string)`

SetLocDesc sets LocDesc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


