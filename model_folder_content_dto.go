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

// checks if the FolderContentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolderContentDto{}

// FolderContentDto struct for FolderContentDto
type FolderContentDto struct {
	Folders []FolderMetadataDto `json:"folders,omitempty"`
	Files []FileMetadataDto `json:"files,omitempty"`
}

// NewFolderContentDto instantiates a new FolderContentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderContentDto() *FolderContentDto {
	this := FolderContentDto{}
	return &this
}

// NewFolderContentDtoWithDefaults instantiates a new FolderContentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderContentDtoWithDefaults() *FolderContentDto {
	this := FolderContentDto{}
	return &this
}

// GetFolders returns the Folders field value if set, zero value otherwise.
func (o *FolderContentDto) GetFolders() []FolderMetadataDto {
	if o == nil || IsNil(o.Folders) {
		var ret []FolderMetadataDto
		return ret
	}
	return o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderContentDto) GetFoldersOk() ([]FolderMetadataDto, bool) {
	if o == nil || IsNil(o.Folders) {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *FolderContentDto) HasFolders() bool {
	if o != nil && !IsNil(o.Folders) {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []FolderMetadataDto and assigns it to the Folders field.
func (o *FolderContentDto) SetFolders(v []FolderMetadataDto) {
	o.Folders = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *FolderContentDto) GetFiles() []FileMetadataDto {
	if o == nil || IsNil(o.Files) {
		var ret []FileMetadataDto
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderContentDto) GetFilesOk() ([]FileMetadataDto, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *FolderContentDto) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FileMetadataDto and assigns it to the Files field.
func (o *FolderContentDto) SetFiles(v []FileMetadataDto) {
	o.Files = v
}

func (o FolderContentDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolderContentDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Folders) {
		toSerialize["folders"] = o.Folders
	}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	return toSerialize, nil
}

type NullableFolderContentDto struct {
	value *FolderContentDto
	isSet bool
}

func (v NullableFolderContentDto) Get() *FolderContentDto {
	return v.value
}

func (v *NullableFolderContentDto) Set(val *FolderContentDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderContentDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderContentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderContentDto(val *FolderContentDto) *NullableFolderContentDto {
	return &NullableFolderContentDto{value: val, isSet: true}
}

func (v NullableFolderContentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderContentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

