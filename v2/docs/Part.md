# Part

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartNo** | **string** |  | [readonly] 
**PartFamily** | [**PartFamily**](PartFamily.md) |  | [readonly] 
**BoardNo** | **NullableString** |  | 
**Variant** | Pointer to **string** |  | [optional] [default to ""]
**Revision** | Pointer to **string** |  | [optional] [default to ""]
**ChipNo** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPart

`func NewPart(partNo string, partFamily PartFamily, boardNo NullableString, ) *Part`

NewPart instantiates a new Part object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartWithDefaults

`func NewPartWithDefaults() *Part`

NewPartWithDefaults instantiates a new Part object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartNo

`func (o *Part) GetPartNo() string`

GetPartNo returns the PartNo field if non-nil, zero value otherwise.

### GetPartNoOk

`func (o *Part) GetPartNoOk() (*string, bool)`

GetPartNoOk returns a tuple with the PartNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNo

`func (o *Part) SetPartNo(v string)`

SetPartNo sets PartNo field to given value.


### GetPartFamily

`func (o *Part) GetPartFamily() PartFamily`

GetPartFamily returns the PartFamily field if non-nil, zero value otherwise.

### GetPartFamilyOk

`func (o *Part) GetPartFamilyOk() (*PartFamily, bool)`

GetPartFamilyOk returns a tuple with the PartFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartFamily

`func (o *Part) SetPartFamily(v PartFamily)`

SetPartFamily sets PartFamily field to given value.


### GetBoardNo

`func (o *Part) GetBoardNo() string`

GetBoardNo returns the BoardNo field if non-nil, zero value otherwise.

### GetBoardNoOk

`func (o *Part) GetBoardNoOk() (*string, bool)`

GetBoardNoOk returns a tuple with the BoardNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardNo

`func (o *Part) SetBoardNo(v string)`

SetBoardNo sets BoardNo field to given value.


### SetBoardNoNil

`func (o *Part) SetBoardNoNil(b bool)`

 SetBoardNoNil sets the value for BoardNo to be an explicit nil

### UnsetBoardNo
`func (o *Part) UnsetBoardNo()`

UnsetBoardNo ensures that no value is present for BoardNo, not even an explicit nil
### GetVariant

`func (o *Part) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *Part) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *Part) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *Part) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetRevision

`func (o *Part) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Part) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Part) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *Part) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetChipNo

`func (o *Part) GetChipNo() string`

GetChipNo returns the ChipNo field if non-nil, zero value otherwise.

### GetChipNoOk

`func (o *Part) GetChipNoOk() (*string, bool)`

GetChipNoOk returns a tuple with the ChipNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChipNo

`func (o *Part) SetChipNo(v string)`

SetChipNo sets ChipNo field to given value.

### HasChipNo

`func (o *Part) HasChipNo() bool`

HasChipNo returns a boolean if a field has been set.

### SetChipNoNil

`func (o *Part) SetChipNoNil(b bool)`

 SetChipNoNil sets the value for ChipNo to be an explicit nil

### UnsetChipNo
`func (o *Part) UnsetChipNo()`

UnsetChipNo ensures that no value is present for ChipNo, not even an explicit nil
### GetMetadata

`func (o *Part) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Part) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Part) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Part) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Part) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Part) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


