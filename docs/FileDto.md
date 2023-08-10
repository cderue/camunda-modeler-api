# FileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**FileMetadataDto**](FileMetadataDto.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewFileDto

`func NewFileDto() *FileDto`

NewFileDto instantiates a new FileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDtoWithDefaults

`func NewFileDtoWithDefaults() *FileDto`

NewFileDtoWithDefaults instantiates a new FileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *FileDto) GetMetadata() FileMetadataDto`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FileDto) GetMetadataOk() (*FileMetadataDto, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FileDto) SetMetadata(v FileMetadataDto)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FileDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetContent

`func (o *FileDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FileDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FileDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *FileDto) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


