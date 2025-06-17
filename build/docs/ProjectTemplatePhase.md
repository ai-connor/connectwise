# ProjectTemplatePhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TemplateRecId** | Pointer to **int32** |  | [optional] 
**ParentPhase** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**BudgetHours** | Pointer to **NullableFloat64** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**MarkAsMilestone** | Pointer to **bool** |  | [optional] 
**PhaseBilledSeparately** | Pointer to **bool** |  | [optional] 
**WbsCode** | Pointer to **string** |  | [optional] 
**ConnectWiseId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableInt32** |  | [optional] 
**ParentConnectWiseId** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectTemplatePhase

`func NewProjectTemplatePhase() *ProjectTemplatePhase`

NewProjectTemplatePhase instantiates a new ProjectTemplatePhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTemplatePhaseWithDefaults

`func NewProjectTemplatePhaseWithDefaults() *ProjectTemplatePhase`

NewProjectTemplatePhaseWithDefaults instantiates a new ProjectTemplatePhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTemplatePhase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTemplatePhase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTemplatePhase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTemplatePhase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTemplateRecId

`func (o *ProjectTemplatePhase) GetTemplateRecId() int32`

GetTemplateRecId returns the TemplateRecId field if non-nil, zero value otherwise.

### GetTemplateRecIdOk

`func (o *ProjectTemplatePhase) GetTemplateRecIdOk() (*int32, bool)`

GetTemplateRecIdOk returns a tuple with the TemplateRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRecId

`func (o *ProjectTemplatePhase) SetTemplateRecId(v int32)`

SetTemplateRecId sets TemplateRecId field to given value.

### HasTemplateRecId

`func (o *ProjectTemplatePhase) HasTemplateRecId() bool`

HasTemplateRecId returns a boolean if a field has been set.

### GetParentPhase

`func (o *ProjectTemplatePhase) GetParentPhase() int32`

GetParentPhase returns the ParentPhase field if non-nil, zero value otherwise.

### GetParentPhaseOk

`func (o *ProjectTemplatePhase) GetParentPhaseOk() (*int32, bool)`

GetParentPhaseOk returns a tuple with the ParentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPhase

`func (o *ProjectTemplatePhase) SetParentPhase(v int32)`

SetParentPhase sets ParentPhase field to given value.

### HasParentPhase

`func (o *ProjectTemplatePhase) HasParentPhase() bool`

HasParentPhase returns a boolean if a field has been set.

### SetParentPhaseNil

`func (o *ProjectTemplatePhase) SetParentPhaseNil(b bool)`

 SetParentPhaseNil sets the value for ParentPhase to be an explicit nil

### UnsetParentPhase
`func (o *ProjectTemplatePhase) UnsetParentPhase()`

UnsetParentPhase ensures that no value is present for ParentPhase, not even an explicit nil
### GetDescription

`func (o *ProjectTemplatePhase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectTemplatePhase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectTemplatePhase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectTemplatePhase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBudgetHours

`func (o *ProjectTemplatePhase) GetBudgetHours() float64`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *ProjectTemplatePhase) GetBudgetHoursOk() (*float64, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *ProjectTemplatePhase) SetBudgetHours(v float64)`

SetBudgetHours sets BudgetHours field to given value.

### HasBudgetHours

`func (o *ProjectTemplatePhase) HasBudgetHours() bool`

HasBudgetHours returns a boolean if a field has been set.

### SetBudgetHoursNil

`func (o *ProjectTemplatePhase) SetBudgetHoursNil(b bool)`

 SetBudgetHoursNil sets the value for BudgetHours to be an explicit nil

### UnsetBudgetHours
`func (o *ProjectTemplatePhase) UnsetBudgetHours()`

UnsetBudgetHours ensures that no value is present for BudgetHours, not even an explicit nil
### GetNotes

`func (o *ProjectTemplatePhase) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ProjectTemplatePhase) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ProjectTemplatePhase) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ProjectTemplatePhase) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetMarkAsMilestone

`func (o *ProjectTemplatePhase) GetMarkAsMilestone() bool`

GetMarkAsMilestone returns the MarkAsMilestone field if non-nil, zero value otherwise.

### GetMarkAsMilestoneOk

`func (o *ProjectTemplatePhase) GetMarkAsMilestoneOk() (*bool, bool)`

GetMarkAsMilestoneOk returns a tuple with the MarkAsMilestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAsMilestone

`func (o *ProjectTemplatePhase) SetMarkAsMilestone(v bool)`

SetMarkAsMilestone sets MarkAsMilestone field to given value.

### HasMarkAsMilestone

`func (o *ProjectTemplatePhase) HasMarkAsMilestone() bool`

HasMarkAsMilestone returns a boolean if a field has been set.

### GetPhaseBilledSeparately

`func (o *ProjectTemplatePhase) GetPhaseBilledSeparately() bool`

GetPhaseBilledSeparately returns the PhaseBilledSeparately field if non-nil, zero value otherwise.

### GetPhaseBilledSeparatelyOk

`func (o *ProjectTemplatePhase) GetPhaseBilledSeparatelyOk() (*bool, bool)`

GetPhaseBilledSeparatelyOk returns a tuple with the PhaseBilledSeparately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseBilledSeparately

`func (o *ProjectTemplatePhase) SetPhaseBilledSeparately(v bool)`

SetPhaseBilledSeparately sets PhaseBilledSeparately field to given value.

### HasPhaseBilledSeparately

`func (o *ProjectTemplatePhase) HasPhaseBilledSeparately() bool`

HasPhaseBilledSeparately returns a boolean if a field has been set.

### GetWbsCode

`func (o *ProjectTemplatePhase) GetWbsCode() string`

GetWbsCode returns the WbsCode field if non-nil, zero value otherwise.

### GetWbsCodeOk

`func (o *ProjectTemplatePhase) GetWbsCodeOk() (*string, bool)`

GetWbsCodeOk returns a tuple with the WbsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbsCode

`func (o *ProjectTemplatePhase) SetWbsCode(v string)`

SetWbsCode sets WbsCode field to given value.

### HasWbsCode

`func (o *ProjectTemplatePhase) HasWbsCode() bool`

HasWbsCode returns a boolean if a field has been set.

### GetConnectWiseId

`func (o *ProjectTemplatePhase) GetConnectWiseId() string`

GetConnectWiseId returns the ConnectWiseId field if non-nil, zero value otherwise.

### GetConnectWiseIdOk

`func (o *ProjectTemplatePhase) GetConnectWiseIdOk() (*string, bool)`

GetConnectWiseIdOk returns a tuple with the ConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseId

`func (o *ProjectTemplatePhase) SetConnectWiseId(v string)`

SetConnectWiseId sets ConnectWiseId field to given value.

### HasConnectWiseId

`func (o *ProjectTemplatePhase) HasConnectWiseId() bool`

HasConnectWiseId returns a boolean if a field has been set.

### GetParentId

`func (o *ProjectTemplatePhase) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ProjectTemplatePhase) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ProjectTemplatePhase) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ProjectTemplatePhase) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ProjectTemplatePhase) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ProjectTemplatePhase) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetParentConnectWiseId

`func (o *ProjectTemplatePhase) GetParentConnectWiseId() string`

GetParentConnectWiseId returns the ParentConnectWiseId field if non-nil, zero value otherwise.

### GetParentConnectWiseIdOk

`func (o *ProjectTemplatePhase) GetParentConnectWiseIdOk() (*string, bool)`

GetParentConnectWiseIdOk returns a tuple with the ParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectWiseId

`func (o *ProjectTemplatePhase) SetParentConnectWiseId(v string)`

SetParentConnectWiseId sets ParentConnectWiseId field to given value.

### HasParentConnectWiseId

`func (o *ProjectTemplatePhase) HasParentConnectWiseId() bool`

HasParentConnectWiseId returns a boolean if a field has been set.

### GetInfo

`func (o *ProjectTemplatePhase) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectTemplatePhase) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectTemplatePhase) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectTemplatePhase) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


