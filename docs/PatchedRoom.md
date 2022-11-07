# PatchedRoom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedRoom

`func NewPatchedRoom() *PatchedRoom`

NewPatchedRoom instantiates a new PatchedRoom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRoomWithDefaults

`func NewPatchedRoomWithDefaults() *PatchedRoom`

NewPatchedRoomWithDefaults instantiates a new PatchedRoom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedRoom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedRoom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedRoom) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedRoom) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSite

`func (o *PatchedRoom) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedRoom) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedRoom) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedRoom) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedRoom) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedRoom) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedRoom) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedRoom) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchedRoom) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchedRoom) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchedRoom) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchedRoom) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


