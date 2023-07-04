# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Room** | [**Room**](Room.md) |  | [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Device** | **int32** |  | 

## Methods

### NewLocation

`func NewLocation(id int32, room Room, device int32, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Location) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location) SetId(v int32)`

SetId sets Id field to given value.


### GetRoom

`func (o *Location) GetRoom() Room`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *Location) GetRoomOk() (*Room, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *Location) SetRoom(v Room)`

SetRoom sets Room field to given value.


### GetDescription

`func (o *Location) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Location) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Location) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Location) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *Location) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Location) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Location) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Location) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDevice

`func (o *Location) GetDevice() int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Location) GetDeviceOk() (*int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Location) SetDevice(v int32)`

SetDevice sets Device field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


