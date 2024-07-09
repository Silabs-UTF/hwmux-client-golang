# PatchedPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartNo** | Pointer to **string** |  | [optional] [readonly] 
**PartFamily** | Pointer to [**PartFamily**](PartFamily.md) |  | [optional] [readonly] 
**BoardNo** | Pointer to **NullableString** |  | [optional] 
**Variant** | Pointer to **string** |  | [optional] [default to ""]
**Revision** | Pointer to **string** |  | [optional] [default to ""]
**ChipNo** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchedPart

`func NewPatchedPart() *PatchedPart`

NewPatchedPart instantiates a new PatchedPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPartWithDefaults

`func NewPatchedPartWithDefaults() *PatchedPart`

NewPatchedPartWithDefaults instantiates a new PatchedPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartNo

`func (o *PatchedPart) GetPartNo() string`

GetPartNo returns the PartNo field if non-nil, zero value otherwise.

### GetPartNoOk

`func (o *PatchedPart) GetPartNoOk() (*string, bool)`

GetPartNoOk returns a tuple with the PartNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNo

`func (o *PatchedPart) SetPartNo(v string)`

SetPartNo sets PartNo field to given value.

### HasPartNo

`func (o *PatchedPart) HasPartNo() bool`

HasPartNo returns a boolean if a field has been set.

### GetPartFamily

`func (o *PatchedPart) GetPartFamily() PartFamily`

GetPartFamily returns the PartFamily field if non-nil, zero value otherwise.

### GetPartFamilyOk

`func (o *PatchedPart) GetPartFamilyOk() (*PartFamily, bool)`

GetPartFamilyOk returns a tuple with the PartFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartFamily

`func (o *PatchedPart) SetPartFamily(v PartFamily)`

SetPartFamily sets PartFamily field to given value.

### HasPartFamily

`func (o *PatchedPart) HasPartFamily() bool`

HasPartFamily returns a boolean if a field has been set.

### GetBoardNo

`func (o *PatchedPart) GetBoardNo() string`

GetBoardNo returns the BoardNo field if non-nil, zero value otherwise.

### GetBoardNoOk

`func (o *PatchedPart) GetBoardNoOk() (*string, bool)`

GetBoardNoOk returns a tuple with the BoardNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardNo

`func (o *PatchedPart) SetBoardNo(v string)`

SetBoardNo sets BoardNo field to given value.

### HasBoardNo

`func (o *PatchedPart) HasBoardNo() bool`

HasBoardNo returns a boolean if a field has been set.

### SetBoardNoNil

`func (o *PatchedPart) SetBoardNoNil(b bool)`

 SetBoardNoNil sets the value for BoardNo to be an explicit nil

### UnsetBoardNo
`func (o *PatchedPart) UnsetBoardNo()`

UnsetBoardNo ensures that no value is present for BoardNo, not even an explicit nil
### GetVariant

`func (o *PatchedPart) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *PatchedPart) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *PatchedPart) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *PatchedPart) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetRevision

`func (o *PatchedPart) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PatchedPart) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PatchedPart) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PatchedPart) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetChipNo

`func (o *PatchedPart) GetChipNo() string`

GetChipNo returns the ChipNo field if non-nil, zero value otherwise.

### GetChipNoOk

`func (o *PatchedPart) GetChipNoOk() (*string, bool)`

GetChipNoOk returns a tuple with the ChipNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChipNo

`func (o *PatchedPart) SetChipNo(v string)`

SetChipNo sets ChipNo field to given value.

### HasChipNo

`func (o *PatchedPart) HasChipNo() bool`

HasChipNo returns a boolean if a field has been set.

### SetChipNoNil

`func (o *PatchedPart) SetChipNoNil(b bool)`

 SetChipNoNil sets the value for ChipNo to be an explicit nil

### UnsetChipNo
`func (o *PatchedPart) UnsetChipNo()`

UnsetChipNo ensures that no value is present for ChipNo, not even an explicit nil
### GetMetadata

`func (o *PatchedPart) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchedPart) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchedPart) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchedPart) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PatchedPart) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PatchedPart) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


