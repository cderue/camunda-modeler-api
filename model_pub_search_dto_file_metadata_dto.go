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

// checks if the PubSearchDtoFileMetadataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PubSearchDtoFileMetadataDto{}

// PubSearchDtoFileMetadataDto struct for PubSearchDtoFileMetadataDto
type PubSearchDtoFileMetadataDto struct {
	Filter *FileMetadataDto `json:"filter,omitempty"`
	Sort []SortDto `json:"sort,omitempty"`
	Page *int32 `json:"page,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewPubSearchDtoFileMetadataDto instantiates a new PubSearchDtoFileMetadataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPubSearchDtoFileMetadataDto() *PubSearchDtoFileMetadataDto {
	this := PubSearchDtoFileMetadataDto{}
	return &this
}

// NewPubSearchDtoFileMetadataDtoWithDefaults instantiates a new PubSearchDtoFileMetadataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPubSearchDtoFileMetadataDtoWithDefaults() *PubSearchDtoFileMetadataDto {
	this := PubSearchDtoFileMetadataDto{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PubSearchDtoFileMetadataDto) GetFilter() FileMetadataDto {
	if o == nil || IsNil(o.Filter) {
		var ret FileMetadataDto
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PubSearchDtoFileMetadataDto) GetFilterOk() (*FileMetadataDto, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PubSearchDtoFileMetadataDto) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given FileMetadataDto and assigns it to the Filter field.
func (o *PubSearchDtoFileMetadataDto) SetFilter(v FileMetadataDto) {
	o.Filter = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *PubSearchDtoFileMetadataDto) GetSort() []SortDto {
	if o == nil || IsNil(o.Sort) {
		var ret []SortDto
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PubSearchDtoFileMetadataDto) GetSortOk() ([]SortDto, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *PubSearchDtoFileMetadataDto) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []SortDto and assigns it to the Sort field.
func (o *PubSearchDtoFileMetadataDto) SetSort(v []SortDto) {
	o.Sort = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PubSearchDtoFileMetadataDto) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PubSearchDtoFileMetadataDto) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PubSearchDtoFileMetadataDto) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *PubSearchDtoFileMetadataDto) SetPage(v int32) {
	o.Page = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PubSearchDtoFileMetadataDto) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PubSearchDtoFileMetadataDto) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PubSearchDtoFileMetadataDto) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *PubSearchDtoFileMetadataDto) SetSize(v int32) {
	o.Size = &v
}

func (o PubSearchDtoFileMetadataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PubSearchDtoFileMetadataDto) ToMap() (map[string]interface{}, error) {
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

type NullablePubSearchDtoFileMetadataDto struct {
	value *PubSearchDtoFileMetadataDto
	isSet bool
}

func (v NullablePubSearchDtoFileMetadataDto) Get() *PubSearchDtoFileMetadataDto {
	return v.value
}

func (v *NullablePubSearchDtoFileMetadataDto) Set(val *PubSearchDtoFileMetadataDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePubSearchDtoFileMetadataDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePubSearchDtoFileMetadataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePubSearchDtoFileMetadataDto(val *PubSearchDtoFileMetadataDto) *NullablePubSearchDtoFileMetadataDto {
	return &NullablePubSearchDtoFileMetadataDto{value: val, isSet: true}
}

func (v NullablePubSearchDtoFileMetadataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePubSearchDtoFileMetadataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

