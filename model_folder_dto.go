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

// checks if the FolderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolderDto{}

// FolderDto struct for FolderDto
type FolderDto struct {
	Metadata *FolderMetadataDto `json:"metadata,omitempty"`
	Content *FolderContentDto `json:"content,omitempty"`
}

// NewFolderDto instantiates a new FolderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderDto() *FolderDto {
	this := FolderDto{}
	return &this
}

// NewFolderDtoWithDefaults instantiates a new FolderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderDtoWithDefaults() *FolderDto {
	this := FolderDto{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *FolderDto) GetMetadata() FolderMetadataDto {
	if o == nil || IsNil(o.Metadata) {
		var ret FolderMetadataDto
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDto) GetMetadataOk() (*FolderMetadataDto, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *FolderDto) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given FolderMetadataDto and assigns it to the Metadata field.
func (o *FolderDto) SetMetadata(v FolderMetadataDto) {
	o.Metadata = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *FolderDto) GetContent() FolderContentDto {
	if o == nil || IsNil(o.Content) {
		var ret FolderContentDto
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDto) GetContentOk() (*FolderContentDto, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *FolderDto) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given FolderContentDto and assigns it to the Content field.
func (o *FolderDto) SetContent(v FolderContentDto) {
	o.Content = &v
}

func (o FolderDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableFolderDto struct {
	value *FolderDto
	isSet bool
}

func (v NullableFolderDto) Get() *FolderDto {
	return v.value
}

func (v *NullableFolderDto) Set(val *FolderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderDto(val *FolderDto) *NullableFolderDto {
	return &NullableFolderDto{value: val, isSet: true}
}

func (v NullableFolderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

