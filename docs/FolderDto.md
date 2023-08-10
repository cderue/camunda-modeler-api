# FolderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**FolderMetadataDto**](FolderMetadataDto.md) |  | [optional] 
**Content** | Pointer to [**FolderContentDto**](FolderContentDto.md) |  | [optional] 

## Methods

### NewFolderDto

`func NewFolderDto() *FolderDto`

NewFolderDto instantiates a new FolderDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderDtoWithDefaults

`func NewFolderDtoWithDefaults() *FolderDto`

NewFolderDtoWithDefaults instantiates a new FolderDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *FolderDto) GetMetadata() FolderMetadataDto`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FolderDto) GetMetadataOk() (*FolderMetadataDto, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FolderDto) SetMetadata(v FolderMetadataDto)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FolderDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetContent

`func (o *FolderDto) GetContent() FolderContentDto`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FolderDto) GetContentOk() (*FolderContentDto, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FolderDto) SetContent(v FolderContentDto)`

SetContent sets Content field to given value.

### HasContent

`func (o *FolderDto) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


