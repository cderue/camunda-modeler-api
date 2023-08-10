# UpdateFileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Revision** | **int32** |  | 
**ProjectId** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateFileDto

`func NewUpdateFileDto(revision int32, ) *UpdateFileDto`

NewUpdateFileDto instantiates a new UpdateFileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFileDtoWithDefaults

`func NewUpdateFileDtoWithDefaults() *UpdateFileDto`

NewUpdateFileDtoWithDefaults instantiates a new UpdateFileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateFileDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFileDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFileDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFileDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContent

`func (o *UpdateFileDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UpdateFileDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UpdateFileDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UpdateFileDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRevision

`func (o *UpdateFileDto) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *UpdateFileDto) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *UpdateFileDto) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetProjectId

`func (o *UpdateFileDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateFileDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateFileDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *UpdateFileDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetFolderId

`func (o *UpdateFileDto) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *UpdateFileDto) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *UpdateFileDto) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *UpdateFileDto) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


