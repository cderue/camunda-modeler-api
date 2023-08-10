# InfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**AuthorizedOrganization** | Pointer to **string** |  | [optional] 
**CreatePermission** | Pointer to **bool** |  | [optional] 
**ReadPermission** | Pointer to **bool** |  | [optional] 
**UpdatePermission** | Pointer to **bool** |  | [optional] 
**DeletePermission** | Pointer to **bool** |  | [optional] 

## Methods

### NewInfoDto

`func NewInfoDto() *InfoDto`

NewInfoDto instantiates a new InfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoDtoWithDefaults

`func NewInfoDtoWithDefaults() *InfoDto`

NewInfoDtoWithDefaults instantiates a new InfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *InfoDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InfoDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InfoDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InfoDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAuthorizedOrganization

`func (o *InfoDto) GetAuthorizedOrganization() string`

GetAuthorizedOrganization returns the AuthorizedOrganization field if non-nil, zero value otherwise.

### GetAuthorizedOrganizationOk

`func (o *InfoDto) GetAuthorizedOrganizationOk() (*string, bool)`

GetAuthorizedOrganizationOk returns a tuple with the AuthorizedOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedOrganization

`func (o *InfoDto) SetAuthorizedOrganization(v string)`

SetAuthorizedOrganization sets AuthorizedOrganization field to given value.

### HasAuthorizedOrganization

`func (o *InfoDto) HasAuthorizedOrganization() bool`

HasAuthorizedOrganization returns a boolean if a field has been set.

### GetCreatePermission

`func (o *InfoDto) GetCreatePermission() bool`

GetCreatePermission returns the CreatePermission field if non-nil, zero value otherwise.

### GetCreatePermissionOk

`func (o *InfoDto) GetCreatePermissionOk() (*bool, bool)`

GetCreatePermissionOk returns a tuple with the CreatePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePermission

`func (o *InfoDto) SetCreatePermission(v bool)`

SetCreatePermission sets CreatePermission field to given value.

### HasCreatePermission

`func (o *InfoDto) HasCreatePermission() bool`

HasCreatePermission returns a boolean if a field has been set.

### GetReadPermission

`func (o *InfoDto) GetReadPermission() bool`

GetReadPermission returns the ReadPermission field if non-nil, zero value otherwise.

### GetReadPermissionOk

`func (o *InfoDto) GetReadPermissionOk() (*bool, bool)`

GetReadPermissionOk returns a tuple with the ReadPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPermission

`func (o *InfoDto) SetReadPermission(v bool)`

SetReadPermission sets ReadPermission field to given value.

### HasReadPermission

`func (o *InfoDto) HasReadPermission() bool`

HasReadPermission returns a boolean if a field has been set.

### GetUpdatePermission

`func (o *InfoDto) GetUpdatePermission() bool`

GetUpdatePermission returns the UpdatePermission field if non-nil, zero value otherwise.

### GetUpdatePermissionOk

`func (o *InfoDto) GetUpdatePermissionOk() (*bool, bool)`

GetUpdatePermissionOk returns a tuple with the UpdatePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePermission

`func (o *InfoDto) SetUpdatePermission(v bool)`

SetUpdatePermission sets UpdatePermission field to given value.

### HasUpdatePermission

`func (o *InfoDto) HasUpdatePermission() bool`

HasUpdatePermission returns a boolean if a field has been set.

### GetDeletePermission

`func (o *InfoDto) GetDeletePermission() bool`

GetDeletePermission returns the DeletePermission field if non-nil, zero value otherwise.

### GetDeletePermissionOk

`func (o *InfoDto) GetDeletePermissionOk() (*bool, bool)`

GetDeletePermissionOk returns a tuple with the DeletePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletePermission

`func (o *InfoDto) SetDeletePermission(v bool)`

SetDeletePermission sets DeletePermission field to given value.

### HasDeletePermission

`func (o *InfoDto) HasDeletePermission() bool`

HasDeletePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


