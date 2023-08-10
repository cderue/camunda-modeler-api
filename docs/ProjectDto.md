# ProjectDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ProjectMetadataDto**](ProjectMetadataDto.md) |  | [optional] 
**Content** | Pointer to [**ProjectContent**](ProjectContent.md) |  | [optional] 

## Methods

### NewProjectDto

`func NewProjectDto() *ProjectDto`

NewProjectDto instantiates a new ProjectDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDtoWithDefaults

`func NewProjectDtoWithDefaults() *ProjectDto`

NewProjectDtoWithDefaults instantiates a new ProjectDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ProjectDto) GetMetadata() ProjectMetadataDto`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProjectDto) GetMetadataOk() (*ProjectMetadataDto, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProjectDto) SetMetadata(v ProjectMetadataDto)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProjectDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetContent

`func (o *ProjectDto) GetContent() ProjectContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ProjectDto) GetContentOk() (*ProjectContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ProjectDto) SetContent(v ProjectContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *ProjectDto) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


