# GroupLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Rooms** | [**[]Room**](Room.md) |  | [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**DeviceGroup** | **int32** |  | 

## Methods

### NewGroupLocation

`func NewGroupLocation(id int32, rooms []Room, deviceGroup int32, ) *GroupLocation`

NewGroupLocation instantiates a new GroupLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupLocationWithDefaults

`func NewGroupLocationWithDefaults() *GroupLocation`

NewGroupLocationWithDefaults instantiates a new GroupLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupLocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupLocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupLocation) SetId(v int32)`

SetId sets Id field to given value.


### GetRooms

`func (o *GroupLocation) GetRooms() []Room`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *GroupLocation) GetRoomsOk() (*[]Room, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *GroupLocation) SetRooms(v []Room)`

SetRooms sets Rooms field to given value.


### GetDescription

`func (o *GroupLocation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupLocation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupLocation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupLocation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *GroupLocation) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GroupLocation) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GroupLocation) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GroupLocation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDeviceGroup

`func (o *GroupLocation) GetDeviceGroup() int32`

GetDeviceGroup returns the DeviceGroup field if non-nil, zero value otherwise.

### GetDeviceGroupOk

`func (o *GroupLocation) GetDeviceGroupOk() (*int32, bool)`

GetDeviceGroupOk returns a tuple with the DeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroup

`func (o *GroupLocation) SetDeviceGroup(v int32)`

SetDeviceGroup sets DeviceGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


