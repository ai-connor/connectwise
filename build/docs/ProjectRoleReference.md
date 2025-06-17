# ProjectRoleReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectRoleReference

`func NewProjectRoleReference() *ProjectRoleReference`

NewProjectRoleReference instantiates a new ProjectRoleReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRoleReferenceWithDefaults

`func NewProjectRoleReferenceWithDefaults() *ProjectRoleReference`

NewProjectRoleReferenceWithDefaults instantiates a new ProjectRoleReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectRoleReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectRoleReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectRoleReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectRoleReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProjectRoleReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProjectRoleReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentifier

`func (o *ProjectRoleReference) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ProjectRoleReference) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ProjectRoleReference) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ProjectRoleReference) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetInfo

`func (o *ProjectRoleReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectRoleReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectRoleReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectRoleReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


