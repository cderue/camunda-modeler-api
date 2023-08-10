# MilestoneDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**MilestoneMetadataDto**](MilestoneMetadataDto.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewMilestoneDto

`func NewMilestoneDto() *MilestoneDto`

NewMilestoneDto instantiates a new MilestoneDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMilestoneDtoWithDefaults

`func NewMilestoneDtoWithDefaults() *MilestoneDto`

NewMilestoneDtoWithDefaults instantiates a new MilestoneDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *MilestoneDto) GetMetadata() MilestoneMetadataDto`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MilestoneDto) GetMetadataOk() (*MilestoneMetadataDto, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MilestoneDto) SetMetadata(v MilestoneMetadataDto)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MilestoneDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetContent

`func (o *MilestoneDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MilestoneDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MilestoneDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MilestoneDto) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


