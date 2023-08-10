# FileMetadataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**SimplePath** | Pointer to **string** |  | [optional] 
**CanonicalPath** | Pointer to [**[]PathElementDto**](PathElementDto.md) |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**UserDto**](UserDto.md) |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to [**UserDto**](UserDto.md) |  | [optional] 

## Methods

### NewFileMetadataDto

`func NewFileMetadataDto() *FileMetadataDto`

NewFileMetadataDto instantiates a new FileMetadataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileMetadataDtoWithDefaults

`func NewFileMetadataDtoWithDefaults() *FileMetadataDto`

NewFileMetadataDtoWithDefaults instantiates a new FileMetadataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileMetadataDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileMetadataDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileMetadataDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FileMetadataDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FileMetadataDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileMetadataDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileMetadataDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileMetadataDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *FileMetadataDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FileMetadataDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FileMetadataDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *FileMetadataDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetFolderId

`func (o *FileMetadataDto) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *FileMetadataDto) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *FileMetadataDto) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *FileMetadataDto) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetSimplePath

`func (o *FileMetadataDto) GetSimplePath() string`

GetSimplePath returns the SimplePath field if non-nil, zero value otherwise.

### GetSimplePathOk

`func (o *FileMetadataDto) GetSimplePathOk() (*string, bool)`

GetSimplePathOk returns a tuple with the SimplePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplePath

`func (o *FileMetadataDto) SetSimplePath(v string)`

SetSimplePath sets SimplePath field to given value.

### HasSimplePath

`func (o *FileMetadataDto) HasSimplePath() bool`

HasSimplePath returns a boolean if a field has been set.

### GetCanonicalPath

`func (o *FileMetadataDto) GetCanonicalPath() []PathElementDto`

GetCanonicalPath returns the CanonicalPath field if non-nil, zero value otherwise.

### GetCanonicalPathOk

`func (o *FileMetadataDto) GetCanonicalPathOk() (*[]PathElementDto, bool)`

GetCanonicalPathOk returns a tuple with the CanonicalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalPath

`func (o *FileMetadataDto) SetCanonicalPath(v []PathElementDto)`

SetCanonicalPath sets CanonicalPath field to given value.

### HasCanonicalPath

`func (o *FileMetadataDto) HasCanonicalPath() bool`

HasCanonicalPath returns a boolean if a field has been set.

### GetRevision

`func (o *FileMetadataDto) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FileMetadataDto) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FileMetadataDto) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *FileMetadataDto) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetType

`func (o *FileMetadataDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileMetadataDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileMetadataDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FileMetadataDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *FileMetadataDto) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FileMetadataDto) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FileMetadataDto) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FileMetadataDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FileMetadataDto) GetCreatedBy() UserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FileMetadataDto) GetCreatedByOk() (*UserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FileMetadataDto) SetCreatedBy(v UserDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FileMetadataDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdated

`func (o *FileMetadataDto) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FileMetadataDto) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FileMetadataDto) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FileMetadataDto) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FileMetadataDto) GetUpdatedBy() UserDto`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FileMetadataDto) GetUpdatedByOk() (*UserDto, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FileMetadataDto) SetUpdatedBy(v UserDto)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FileMetadataDto) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


