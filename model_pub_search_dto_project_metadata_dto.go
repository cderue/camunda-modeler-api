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

// checks if the PubSearchDtoProjectMetadataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PubSearchDtoProjectMetadataDto{}

// PubSearchDtoProjectMetadataDto struct for PubSearchDtoProjectMetadataDto
type PubSearchDtoProjectMetadataDto struct {
	Filter *ProjectMetadataDto `json:"filter,omitempty"`
	Sort []SortDto `json:"sort,omitempty"`
	Page *int32 `json:"page,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewPubSearchDtoProjectMetadataDto instantiates a new PubSearchDtoProjectMetadataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPubSearchDtoProjectMetadataDto() *PubSearchDtoProjectMetadataDto {
	this := PubSearchDtoProjectMetadataDto{}
	return &this
}

// NewPubSearchDtoProjectMetadataDtoWithDefaults instantiates a new PubSearchDtoProjectMetadataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPubSearchDtoProjectMetadataDtoWithDefaults() *PubSearchDtoProjectMetadataDto {
	this := PubSearchDtoProjectMetadataDto{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PubSearchDtoProjectMetadataDto) GetFilter() ProjectMetadataDto {
	if o == nil || IsNil(o.Filter) {
		var ret ProjectMetadataDto
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PubSearchDtoProjectMetadataDto) GetFilterOk() (*ProjectMetadataDto, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PubSearchDtoProjectMetadataDto) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given ProjectMetadataDto and assigns it to the Filter field.
func (o *PubSearchDtoProjectMetadataDto) SetFilter(v ProjectMetadataDto) {
	o.Filter = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *PubSearchDtoProjectMetadataDto) GetSort() []SortDto {
	if o == nil || IsNil(o.Sort) {
		var ret []SortDto
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PubSearchDtoProjectMetadataDto) GetSortOk() ([]SortDto, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *PubSearchDtoProjectMetadataDto) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []SortDto and assigns it to the Sort field.
func (o *PubSearchDtoProjectMetadataDto) SetSort(v []SortDto) {
	o.Sort = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PubSearchDtoProjectMetadataDto) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PubSearchDtoProjectMetadataDto) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PubSearchDtoProjectMetadataDto) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *PubSearchDtoProjectMetadataDto) SetPage(v int32) {
	o.Page = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PubSearchDtoProjectMetadataDto) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PubSearchDtoProjectMetadataDto) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PubSearchDtoProjectMetadataDto) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *PubSearchDtoProjectMetadataDto) SetSize(v int32) {
	o.Size = &v
}

func (o PubSearchDtoProjectMetadataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PubSearchDtoProjectMetadataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullablePubSearchDtoProjectMetadataDto struct {
	value *PubSearchDtoProjectMetadataDto
	isSet bool
}

func (v NullablePubSearchDtoProjectMetadataDto) Get() *PubSearchDtoProjectMetadataDto {
	return v.value
}

func (v *NullablePubSearchDtoProjectMetadataDto) Set(val *PubSearchDtoProjectMetadataDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePubSearchDtoProjectMetadataDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePubSearchDtoProjectMetadataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePubSearchDtoProjectMetadataDto(val *PubSearchDtoProjectMetadataDto) *NullablePubSearchDtoProjectMetadataDto {
	return &NullablePubSearchDtoProjectMetadataDto{value: val, isSet: true}
}

func (v NullablePubSearchDtoProjectMetadataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePubSearchDtoProjectMetadataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


