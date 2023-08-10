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

// checks if the ProjectMetadataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMetadataDto{}

// ProjectMetadataDto struct for ProjectMetadataDto
type ProjectMetadataDto struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Created *string `json:"created,omitempty"`
	CreatedBy *UserDto `json:"createdBy,omitempty"`
	Updated *string `json:"updated,omitempty"`
	UpdatedBy *UserDto `json:"updatedBy,omitempty"`
}

// NewProjectMetadataDto instantiates a new ProjectMetadataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMetadataDto() *ProjectMetadataDto {
	this := ProjectMetadataDto{}
	return &this
}

// NewProjectMetadataDtoWithDefaults instantiates a new ProjectMetadataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMetadataDtoWithDefaults() *ProjectMetadataDto {
	this := ProjectMetadataDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectMetadataDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMetadataDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectMetadataDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProjectMetadataDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectMetadataDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMetadataDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectMetadataDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectMetadataDto) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ProjectMetadataDto) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMetadataDto) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ProjectMetadataDto) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *ProjectMetadataDto) SetCreated(v string) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ProjectMetadataDto) GetCreatedBy() UserDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret UserDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMetadataDto) GetCreatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ProjectMetadataDto) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given UserDto and assigns it to the CreatedBy field.
func (o *ProjectMetadataDto) SetCreatedBy(v UserDto) {
	o.CreatedBy = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ProjectMetadataDto) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMetadataDto) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ProjectMetadataDto) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *ProjectMetadataDto) SetUpdated(v string) {
	o.Updated = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ProjectMetadataDto) GetUpdatedBy() UserDto {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret UserDto
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMetadataDto) GetUpdatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ProjectMetadataDto) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given UserDto and assigns it to the UpdatedBy field.
func (o *ProjectMetadataDto) SetUpdatedBy(v UserDto) {
	o.UpdatedBy = &v
}

func (o ProjectMetadataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMetadataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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

type NullableProjectMetadataDto struct {
	value *ProjectMetadataDto
	isSet bool
}

func (v NullableProjectMetadataDto) Get() *ProjectMetadataDto {
	return v.value
}

func (v *NullableProjectMetadataDto) Set(val *ProjectMetadataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMetadataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMetadataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMetadataDto(val *ProjectMetadataDto) *NullableProjectMetadataDto {
	return &NullableProjectMetadataDto{value: val, isSet: true}
}

func (v NullableProjectMetadataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMetadataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

