# LocationSerializerWriteOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Room** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewLocationSerializerWriteOnly

`func NewLocationSerializerWriteOnly(id int32, room string, ) *LocationSerializerWriteOnly`

NewLocationSerializerWriteOnly instantiates a new LocationSerializerWriteOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationSerializerWriteOnlyWithDefaults

`func NewLocationSerializerWriteOnlyWithDefaults() *LocationSerializerWriteOnly`

NewLocationSerializerWriteOnlyWithDefaults instantiates a new LocationSerializerWriteOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocationSerializerWriteOnly) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationSerializerWriteOnly) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationSerializerWriteOnly) SetId(v int32)`

SetId sets Id field to given value.


### GetRoom

`func (o *LocationSerializerWriteOnly) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *LocationSerializerWriteOnly) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *LocationSerializerWriteOnly) SetRoom(v string)`

SetRoom sets Room field to given value.


### GetDescription

`func (o *LocationSerializerWriteOnly) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LocationSerializerWriteOnly) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LocationSerializerWriteOnly) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LocationSerializerWriteOnly) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *LocationSerializerWriteOnly) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LocationSerializerWriteOnly) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LocationSerializerWriteOnly) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LocationSerializerWriteOnly) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


