# ProjectNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**Text** | **string** |  | 
**Type** | Pointer to [**NoteTypeReference**](NoteTypeReference.md) |  | [optional] 
**Flagged** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectNote

`func NewProjectNote(text string, ) *ProjectNote`

NewProjectNote instantiates a new ProjectNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectNoteWithDefaults

`func NewProjectNoteWithDefaults() *ProjectNote`

NewProjectNoteWithDefaults instantiates a new ProjectNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectNote) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectNote) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectNote) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectNote) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectNote) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ProjectNote) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ProjectNote) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetText

`func (o *ProjectNote) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ProjectNote) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ProjectNote) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *ProjectNote) GetType() NoteTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectNote) GetTypeOk() (*NoteTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectNote) SetType(v NoteTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *ProjectNote) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlagged

`func (o *ProjectNote) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *ProjectNote) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *ProjectNote) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *ProjectNote) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### SetFlaggedNil

`func (o *ProjectNote) SetFlaggedNil(b bool)`

 SetFlaggedNil sets the value for Flagged to be an explicit nil

### UnsetFlagged
`func (o *ProjectNote) UnsetFlagged()`

UnsetFlagged ensures that no value is present for Flagged, not even an explicit nil
### GetInfo

`func (o *ProjectNote) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectNote) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectNote) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectNote) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


