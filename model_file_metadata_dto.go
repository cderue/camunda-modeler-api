/*
Web Modeler REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 41f4f56
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FileMetadataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileMetadataDto{}

// FileMetadataDto struct for FileMetadataDto
type FileMetadataDto struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
	FolderId *string `json:"folderId,omitempty"`
	SimplePath *string `json:"simplePath,omitempty"`
	CanonicalPath []PathElementDto `json:"canonicalPath,omitempty"`
	Revision *int32 `json:"revision,omitempty"`
	Type *string `json:"type,omitempty"`
	Created *string `json:"created,omitempty"`
	CreatedBy *UserDto `json:"createdBy,omitempty"`
	Updated *string `json:"updated,omitempty"`
	UpdatedBy *UserDto `json:"updatedBy,omitempty"`
}

// NewFileMetadataDto instantiates a new FileMetadataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileMetadataDto() *FileMetadataDto {
	this := FileMetadataDto{}
	return &this
}

// NewFileMetadataDtoWithDefaults instantiates a new FileMetadataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileMetadataDtoWithDefaults() *FileMetadataDto {
	this := FileMetadataDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileMetadataDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadataDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileMetadataDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FileMetadataDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FileMetadataDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadataDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FileMetadataDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FileMetadataDto) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *FileMetadataDto) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadataDto) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *FileMetadataDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *FileMetadataDto) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *FileMetadataDto) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadataDto) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *FileMetadataDto) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *FileMetadataDto) SetFolderId(v string) {
	o.FolderId = &v
}

// GetSimplePath returns the SimplePath field value if set, zero value otherwise.
func (o *FileMetadataDto) GetSimplePath() string {
	if o == nil || IsNil(o.SimplePath) {
		var ret string
		return ret
	}
	return *o.SimplePath
}

// GetSimplePathOk returns a tuple with the SimplePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadataDto) GetSimplePathOk() (*string, bool) {
	if o == nil || IsNil(o.SimplePath) {
		return nil, false
	}
	return o.SimplePath, true
}

// HasSimplePath returns a boolean if a field has been set.
func (o *FileMetadataDto) HasSimplePath() bool {
	if o != nil && !IsNil(o.SimplePath) {
		return true
	}

	return false
}

// SetSimplePath gets a reference to the given string and assigns it to the SimplePath field.
func (o *FileMetadataDto) SetSimplePath(v string) {
	o.SimplePath = &v
}

// GetCanonicalPath returns the CanonicalPath field value if set, zero value otherwise.
func (o *FileMetadataDto) GetCanonicalPath() []PathElementDto {
	if o == nil || IsNil(o.CanonicalPath) {
		var ret []PathElementDto
		return ret
	}
	return o.CanonicalPath
}

// GetCanonicalPathOk returns a tuple with the CanonicalPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadataDto) GetCanonicalPathOk() ([]PathElementDto, bool) {
	if o == nil || IsNil(o.CanonicalPath) {
		return nil, false
	}
	return o.CanonicalPath, true
}

// HasCanonicalPath returns a boolean if a field has been set.
func (o *FileMetadataDto) HasCanonicalPath() bool {
	if o != nil && !IsNil(o.CanonicalPath) {
		return true
	}

	return false
}

// SetCanonicalPath gets a reference to the given []PathElementDto and assigns it to the CanonicalPath field.
func (o *FileMetadataDto) SetCanonicalPath(v []PathElementDto) {
	o.CanonicalPath = v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *FileMetadataDto) GetRevision() int32 {
	if o == nil || IsNil(o.Revision) {
		var ret int32
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadataDto) GetRevisionOk() (*int32, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *FileMetadataDto) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given int32 and assigns it to the Revision field.
func (o *FileMetadataDto) SetRevision(v int32) {
	o.Revision = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FileMetadataDto) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadataDto) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FileMetadataDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FileMetadataDto) SetType(v string) {
	o.Type = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FileMetadataDto) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadataDto) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FileMetadataDto) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *FileMetadataDto) SetCreated(v string) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FileMetadataDto) GetCreatedBy() UserDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret UserDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadataDto) GetCreatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FileMetadataDto) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given UserDto and assigns it to the CreatedBy field.
func (o *FileMetadataDto) SetCreatedBy(v UserDto) {
	o.CreatedBy = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *FileMetadataDto) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadataDto) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *FileMetadataDto) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *FileMetadataDto) SetUpdated(v string) {
	o.Updated = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *FileMetadataDto) GetUpdatedBy() UserDto {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret UserDto
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMetadataDto) GetUpdatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *FileMetadataDto) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given UserDto and assigns it to the UpdatedBy field.
func (o *FileMetadataDto) SetUpdatedBy(v UserDto) {
	o.UpdatedBy = &v
}

func (o FileMetadataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileMetadataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if !IsNil(o.SimplePath) {
		toSerialize["simplePath"] = o.SimplePath
	}
	if !IsNil(o.CanonicalPath) {
		toSerialize["canonicalPath"] = o.CanonicalPath
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	return toSerialize, nil
}

type NullableFileMetadataDto struct {
	value *FileMetadataDto
	isSet bool
}

func (v NullableFileMetadataDto) Get() *FileMetadataDto {
	return v.value
}

func (v *NullableFileMetadataDto) Set(val *FileMetadataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFileMetadataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFileMetadataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileMetadataDto(val *FileMetadataDto) *NullableFileMetadataDto {
	return &NullableFileMetadataDto{value: val, isSet: true}
}

func (v NullableFileMetadataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileMetadataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


