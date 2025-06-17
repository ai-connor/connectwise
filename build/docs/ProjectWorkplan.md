# ProjectWorkplan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Phases** | Pointer to [**[]ProjectWorkplanProjectPhase**](ProjectWorkplanProjectPhase.md) |  | [optional] 

## Methods

### NewProjectWorkplan

`func NewProjectWorkplan() *ProjectWorkplan`

NewProjectWorkplan instantiates a new ProjectWorkplan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWorkplanWithDefaults

`func NewProjectWorkplanWithDefaults() *ProjectWorkplan`

NewProjectWorkplanWithDefaults instantiates a new ProjectWorkplan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ProjectWorkplan) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectWorkplan) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectWorkplan) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectWorkplan) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPhases

`func (o *ProjectWorkplan) GetPhases() []ProjectWorkplanProjectPhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *ProjectWorkplan) GetPhasesOk() (*[]ProjectWorkplanProjectPhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *ProjectWorkplan) SetPhases(v []ProjectWorkplanProjectPhase)`

SetPhases sets Phases field to given value.

### HasPhases

`func (o *ProjectWorkplan) HasPhases() bool`

HasPhases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


