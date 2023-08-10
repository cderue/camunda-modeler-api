# PubSearchDtoMilestoneMetadataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**MilestoneMetadataDto**](MilestoneMetadataDto.md) |  | [optional] 
**Sort** | Pointer to [**[]SortDto**](SortDto.md) |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewPubSearchDtoMilestoneMetadataDto

`func NewPubSearchDtoMilestoneMetadataDto() *PubSearchDtoMilestoneMetadataDto`

NewPubSearchDtoMilestoneMetadataDto instantiates a new PubSearchDtoMilestoneMetadataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubSearchDtoMilestoneMetadataDtoWithDefaults

`func NewPubSearchDtoMilestoneMetadataDtoWithDefaults() *PubSearchDtoMilestoneMetadataDto`

NewPubSearchDtoMilestoneMetadataDtoWithDefaults instantiates a new PubSearchDtoMilestoneMetadataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *PubSearchDtoMilestoneMetadataDto) GetFilter() MilestoneMetadataDto`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PubSearchDtoMilestoneMetadataDto) GetFilterOk() (*MilestoneMetadataDto, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PubSearchDtoMilestoneMetadataDto) SetFilter(v MilestoneMetadataDto)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PubSearchDtoMilestoneMetadataDto) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *PubSearchDtoMilestoneMetadataDto) GetSort() []SortDto`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PubSearchDtoMilestoneMetadataDto) GetSortOk() (*[]SortDto, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PubSearchDtoMilestoneMetadataDto) SetSort(v []SortDto)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PubSearchDtoMilestoneMetadataDto) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetPage

`func (o *PubSearchDtoMilestoneMetadataDto) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PubSearchDtoMilestoneMetadataDto) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PubSearchDtoMilestoneMetadataDto) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PubSearchDtoMilestoneMetadataDto) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *PubSearchDtoMilestoneMetadataDto) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PubSearchDtoMilestoneMetadataDto) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PubSearchDtoMilestoneMetadataDto) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PubSearchDtoMilestoneMetadataDto) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


