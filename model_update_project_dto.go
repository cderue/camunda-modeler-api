/*
Web Modeler REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 41f4f56
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateProjectDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProjectDto{}

// UpdateProjectDto struct for UpdateProjectDto
type UpdateProjectDto struct {
	Name string `json:"name"`
}

type _UpdateProjectDto UpdateProjectDto

// NewUpdateProjectDto instantiates a new UpdateProjectDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProjectDto(name string) *UpdateProjectDto {
	this := UpdateProjectDto{}
	this.Name = name
	return &this
}

// NewUpdateProjectDtoWithDefaults instantiates a new UpdateProjectDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProjectDtoWithDefaults() *UpdateProjectDto {
	this := UpdateProjectDto{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateProjectDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateProjectDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateProjectDto) SetName(v string) {
	o.Name = v
}

func (o UpdateProjectDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProjectDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *UpdateProjectDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateProjectDto := _UpdateProjectDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateProjectDto)

	if err != nil {
		return err
	}

	*o = UpdateProjectDto(varUpdateProjectDto)

	return err
}

type NullableUpdateProjectDto struct {
	value *UpdateProjectDto
	isSet bool
}

func (v NullableUpdateProjectDto) Get() *UpdateProjectDto {
	return v.value
}

func (v *NullableUpdateProjectDto) Set(val *UpdateProjectDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProjectDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProjectDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProjectDto(val *UpdateProjectDto) *NullableUpdateProjectDto {
	return &NullableUpdateProjectDto{value: val, isSet: true}
}

func (v NullableUpdateProjectDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProjectDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

