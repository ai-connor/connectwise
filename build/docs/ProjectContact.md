# ProjectContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Contact** | [**ContactReference**](ContactReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectContact

`func NewProjectContact(contact ContactReference, ) *ProjectContact`

NewProjectContact instantiates a new ProjectContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectContactWithDefaults

`func NewProjectContactWithDefaults() *ProjectContact`

NewProjectContactWithDefaults instantiates a new ProjectContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectContact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectContact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectContact) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectContact) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectContact) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectContact) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectContact) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetContact

`func (o *ProjectContact) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ProjectContact) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ProjectContact) SetContact(v ContactReference)`

SetContact sets Contact field to given value.


### GetInfo

`func (o *ProjectContact) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectContact) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectContact) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectContact) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


