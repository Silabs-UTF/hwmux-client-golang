# ResourceStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**StatusF44Enum**](StatusF44Enum.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewResourceStatusRequest

`func NewResourceStatusRequest(status StatusF44Enum, ) *ResourceStatusRequest`

NewResourceStatusRequest instantiates a new ResourceStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceStatusRequestWithDefaults

`func NewResourceStatusRequestWithDefaults() *ResourceStatusRequest`

NewResourceStatusRequestWithDefaults instantiates a new ResourceStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResourceStatusRequest) GetStatus() StatusF44Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceStatusRequest) GetStatusOk() (*StatusF44Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceStatusRequest) SetStatus(v StatusF44Enum)`

SetStatus sets Status field to given value.


### GetMetadata

`func (o *ResourceStatusRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceStatusRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceStatusRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceStatusRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetComment

`func (o *ResourceStatusRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ResourceStatusRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ResourceStatusRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ResourceStatusRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


