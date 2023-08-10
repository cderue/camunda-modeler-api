# FolderContentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folders** | Pointer to [**[]FolderMetadataDto**](FolderMetadataDto.md) |  | [optional] 
**Files** | Pointer to [**[]FileMetadataDto**](FileMetadataDto.md) |  | [optional] 

## Methods

### NewFolderContentDto

`func NewFolderContentDto() *FolderContentDto`

NewFolderContentDto instantiates a new FolderContentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderContentDtoWithDefaults

`func NewFolderContentDtoWithDefaults() *FolderContentDto`

NewFolderContentDtoWithDefaults instantiates a new FolderContentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolders

`func (o *FolderContentDto) GetFolders() []FolderMetadataDto`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *FolderContentDto) GetFoldersOk() (*[]FolderMetadataDto, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *FolderContentDto) SetFolders(v []FolderMetadataDto)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *FolderContentDto) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetFiles

`func (o *FolderContentDto) GetFiles() []FileMetadataDto`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FolderContentDto) GetFilesOk() (*[]FileMetadataDto, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FolderContentDto) SetFiles(v []FileMetadataDto)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FolderContentDto) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


