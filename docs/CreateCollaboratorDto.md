# CreateCollaboratorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**ProjectId** | **string** |  | 
**Role** | **string** | Allowed values are: \&quot;project_admin\&quot;, \&quot;editor\&quot;, \&quot;viewer\&quot; and \&quot;commenter\&quot; | 

## Methods

### NewCreateCollaboratorDto

`func NewCreateCollaboratorDto(email string, projectId string, role string, ) *CreateCollaboratorDto`

NewCreateCollaboratorDto instantiates a new CreateCollaboratorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCollaboratorDtoWithDefaults

`func NewCreateCollaboratorDtoWithDefaults() *CreateCollaboratorDto`

NewCreateCollaboratorDtoWithDefaults instantiates a new CreateCollaboratorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateCollaboratorDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateCollaboratorDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateCollaboratorDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProjectId

`func (o *CreateCollaboratorDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateCollaboratorDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateCollaboratorDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetRole

`func (o *CreateCollaboratorDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateCollaboratorDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateCollaboratorDto) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


