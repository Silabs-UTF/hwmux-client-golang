# PartFamily

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**RegexPattern** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPartFamily

`func NewPartFamily(name string, ) *PartFamily`

NewPartFamily instantiates a new PartFamily object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartFamilyWithDefaults

`func NewPartFamilyWithDefaults() *PartFamily`

NewPartFamilyWithDefaults instantiates a new PartFamily object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PartFamily) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartFamily) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartFamily) SetName(v string)`

SetName sets Name field to given value.


### GetRegexPattern

`func (o *PartFamily) GetRegexPattern() string`

GetRegexPattern returns the RegexPattern field if non-nil, zero value otherwise.

### GetRegexPatternOk

`func (o *PartFamily) GetRegexPatternOk() (*string, bool)`

GetRegexPatternOk returns a tuple with the RegexPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexPattern

`func (o *PartFamily) SetRegexPattern(v string)`

SetRegexPattern sets RegexPattern field to given value.

### HasRegexPattern

`func (o *PartFamily) HasRegexPattern() bool`

HasRegexPattern returns a boolean if a field has been set.

### SetRegexPatternNil

`func (o *PartFamily) SetRegexPatternNil(b bool)`

 SetRegexPatternNil sets the value for RegexPattern to be an explicit nil

### UnsetRegexPattern
`func (o *PartFamily) UnsetRegexPattern()`

UnsetRegexPattern ensures that no value is present for RegexPattern, not even an explicit nil
### GetDescription

`func (o *PartFamily) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PartFamily) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PartFamily) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PartFamily) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *PartFamily) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PartFamily) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PartFamily) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PartFamily) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


