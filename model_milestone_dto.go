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

// checks if the MilestoneDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MilestoneDto{}

// MilestoneDto struct for MilestoneDto
type MilestoneDto struct {
	Metadata *MilestoneMetadataDto `json:"metadata,omitempty"`
	Content *string `json:"content,omitempty"`
}

// NewMilestoneDto instantiates a new MilestoneDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMilestoneDto() *MilestoneDto {
	this := MilestoneDto{}
	return &this
}

// NewMilestoneDtoWithDefaults instantiates a new MilestoneDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMilestoneDtoWithDefaults() *MilestoneDto {
	this := MilestoneDto{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *MilestoneDto) GetMetadata() MilestoneMetadataDto {
	if o == nil || IsNil(o.Metadata) {
		var ret MilestoneMetadataDto
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilestoneDto) GetMetadataOk() (*MilestoneMetadataDto, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *MilestoneDto) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given MilestoneMetadataDto and assigns it to the Metadata field.
func (o *MilestoneDto) SetMetadata(v MilestoneMetadataDto) {
	o.Metadata = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *MilestoneDto) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilestoneDto) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *MilestoneDto) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *MilestoneDto) SetContent(v string) {
	o.Content = &v
}

func (o MilestoneDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MilestoneDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableMilestoneDto struct {
	value *MilestoneDto
	isSet bool
}

func (v NullableMilestoneDto) Get() *MilestoneDto {
	return v.value
}

func (v *NullableMilestoneDto) Set(val *MilestoneDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMilestoneDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMilestoneDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMilestoneDto(val *MilestoneDto) *NullableMilestoneDto {
	return &NullableMilestoneDto{value: val, isSet: true}
}

func (v NullableMilestoneDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMilestoneDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

