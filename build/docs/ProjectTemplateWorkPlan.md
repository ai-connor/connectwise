# ProjectTemplateWorkPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | Pointer to **int32** |  | [optional] 
**Phases** | Pointer to [**[]TemplatePhase**](TemplatePhase.md) |  | [optional] 

## Methods

### NewProjectTemplateWorkPlan

`func NewProjectTemplateWorkPlan() *ProjectTemplateWorkPlan`

NewProjectTemplateWorkPlan instantiates a new ProjectTemplateWorkPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTemplateWorkPlanWithDefaults

`func NewProjectTemplateWorkPlanWithDefaults() *ProjectTemplateWorkPlan`

NewProjectTemplateWorkPlanWithDefaults instantiates a new ProjectTemplateWorkPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *ProjectTemplateWorkPlan) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ProjectTemplateWorkPlan) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ProjectTemplateWorkPlan) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ProjectTemplateWorkPlan) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetPhases

`func (o *ProjectTemplateWorkPlan) GetPhases() []TemplatePhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *ProjectTemplateWorkPlan) GetPhasesOk() (*[]TemplatePhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *ProjectTemplateWorkPlan) SetPhases(v []TemplatePhase)`

SetPhases sets Phases field to given value.

### HasPhases

`func (o *ProjectTemplateWorkPlan) HasPhases() bool`

HasPhases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


