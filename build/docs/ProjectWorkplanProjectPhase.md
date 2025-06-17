# ProjectWorkplanProjectPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PhaseStatusReference**](PhaseStatusReference.md) |  | [optional] 
**ParentPhase** | Pointer to [**ProjectPhaseReference**](ProjectPhaseReference.md) |  | [optional] 
**WbsCode** | Pointer to **string** |  | [optional] 
**MarkAsMilestoneFlag** | Pointer to **NullableBool** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**BudgetHours** | Pointer to **NullableFloat64** |  | [optional] 
**ActualHours** | Pointer to **NullableFloat64** |  | [optional] 
**BillableHours** | Pointer to **NullableFloat64** |  | [optional] 
**ScheduledHours** | Pointer to **NullableFloat64** |  | [optional] 
**ScheduledStart** | Pointer to **string** |  | [optional] 
**ScheduledEnd** | Pointer to **string** |  | [optional] 
**ScheduledDuration** | Pointer to **NullableInt32** |  | [optional] 
**BillPhaseSeparately** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewProjectWorkplanProjectPhase

`func NewProjectWorkplanProjectPhase() *ProjectWorkplanProjectPhase`

NewProjectWorkplanProjectPhase instantiates a new ProjectWorkplanProjectPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWorkplanProjectPhaseWithDefaults

`func NewProjectWorkplanProjectPhaseWithDefaults() *ProjectWorkplanProjectPhase`

NewProjectWorkplanProjectPhaseWithDefaults instantiates a new ProjectWorkplanProjectPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectWorkplanProjectPhase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectWorkplanProjectPhase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectWorkplanProjectPhase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectWorkplanProjectPhase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectWorkplanProjectPhase) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectWorkplanProjectPhase) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectWorkplanProjectPhase) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectWorkplanProjectPhase) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ProjectWorkplanProjectPhase) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ProjectWorkplanProjectPhase) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetDescription

`func (o *ProjectWorkplanProjectPhase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectWorkplanProjectPhase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectWorkplanProjectPhase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectWorkplanProjectPhase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectWorkplanProjectPhase) GetStatus() PhaseStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectWorkplanProjectPhase) GetStatusOk() (*PhaseStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectWorkplanProjectPhase) SetStatus(v PhaseStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectWorkplanProjectPhase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParentPhase

`func (o *ProjectWorkplanProjectPhase) GetParentPhase() ProjectPhaseReference`

GetParentPhase returns the ParentPhase field if non-nil, zero value otherwise.

### GetParentPhaseOk

`func (o *ProjectWorkplanProjectPhase) GetParentPhaseOk() (*ProjectPhaseReference, bool)`

GetParentPhaseOk returns a tuple with the ParentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPhase

`func (o *ProjectWorkplanProjectPhase) SetParentPhase(v ProjectPhaseReference)`

SetParentPhase sets ParentPhase field to given value.

### HasParentPhase

`func (o *ProjectWorkplanProjectPhase) HasParentPhase() bool`

HasParentPhase returns a boolean if a field has been set.

### GetWbsCode

`func (o *ProjectWorkplanProjectPhase) GetWbsCode() string`

GetWbsCode returns the WbsCode field if non-nil, zero value otherwise.

### GetWbsCodeOk

`func (o *ProjectWorkplanProjectPhase) GetWbsCodeOk() (*string, bool)`

GetWbsCodeOk returns a tuple with the WbsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbsCode

`func (o *ProjectWorkplanProjectPhase) SetWbsCode(v string)`

SetWbsCode sets WbsCode field to given value.

### HasWbsCode

`func (o *ProjectWorkplanProjectPhase) HasWbsCode() bool`

HasWbsCode returns a boolean if a field has been set.

### GetMarkAsMilestoneFlag

`func (o *ProjectWorkplanProjectPhase) GetMarkAsMilestoneFlag() bool`

GetMarkAsMilestoneFlag returns the MarkAsMilestoneFlag field if non-nil, zero value otherwise.

### GetMarkAsMilestoneFlagOk

`func (o *ProjectWorkplanProjectPhase) GetMarkAsMilestoneFlagOk() (*bool, bool)`

GetMarkAsMilestoneFlagOk returns a tuple with the MarkAsMilestoneFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAsMilestoneFlag

`func (o *ProjectWorkplanProjectPhase) SetMarkAsMilestoneFlag(v bool)`

SetMarkAsMilestoneFlag sets MarkAsMilestoneFlag field to given value.

### HasMarkAsMilestoneFlag

`func (o *ProjectWorkplanProjectPhase) HasMarkAsMilestoneFlag() bool`

HasMarkAsMilestoneFlag returns a boolean if a field has been set.

### SetMarkAsMilestoneFlagNil

`func (o *ProjectWorkplanProjectPhase) SetMarkAsMilestoneFlagNil(b bool)`

 SetMarkAsMilestoneFlagNil sets the value for MarkAsMilestoneFlag to be an explicit nil

### UnsetMarkAsMilestoneFlag
`func (o *ProjectWorkplanProjectPhase) UnsetMarkAsMilestoneFlag()`

UnsetMarkAsMilestoneFlag ensures that no value is present for MarkAsMilestoneFlag, not even an explicit nil
### GetNotes

`func (o *ProjectWorkplanProjectPhase) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ProjectWorkplanProjectPhase) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ProjectWorkplanProjectPhase) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ProjectWorkplanProjectPhase) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetStartDate

`func (o *ProjectWorkplanProjectPhase) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ProjectWorkplanProjectPhase) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ProjectWorkplanProjectPhase) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ProjectWorkplanProjectPhase) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ProjectWorkplanProjectPhase) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ProjectWorkplanProjectPhase) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ProjectWorkplanProjectPhase) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ProjectWorkplanProjectPhase) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetBudgetHours

`func (o *ProjectWorkplanProjectPhase) GetBudgetHours() float64`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *ProjectWorkplanProjectPhase) GetBudgetHoursOk() (*float64, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *ProjectWorkplanProjectPhase) SetBudgetHours(v float64)`

SetBudgetHours sets BudgetHours field to given value.

### HasBudgetHours

`func (o *ProjectWorkplanProjectPhase) HasBudgetHours() bool`

HasBudgetHours returns a boolean if a field has been set.

### SetBudgetHoursNil

`func (o *ProjectWorkplanProjectPhase) SetBudgetHoursNil(b bool)`

 SetBudgetHoursNil sets the value for BudgetHours to be an explicit nil

### UnsetBudgetHours
`func (o *ProjectWorkplanProjectPhase) UnsetBudgetHours()`

UnsetBudgetHours ensures that no value is present for BudgetHours, not even an explicit nil
### GetActualHours

`func (o *ProjectWorkplanProjectPhase) GetActualHours() float64`

GetActualHours returns the ActualHours field if non-nil, zero value otherwise.

### GetActualHoursOk

`func (o *ProjectWorkplanProjectPhase) GetActualHoursOk() (*float64, bool)`

GetActualHoursOk returns a tuple with the ActualHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualHours

`func (o *ProjectWorkplanProjectPhase) SetActualHours(v float64)`

SetActualHours sets ActualHours field to given value.

### HasActualHours

`func (o *ProjectWorkplanProjectPhase) HasActualHours() bool`

HasActualHours returns a boolean if a field has been set.

### SetActualHoursNil

`func (o *ProjectWorkplanProjectPhase) SetActualHoursNil(b bool)`

 SetActualHoursNil sets the value for ActualHours to be an explicit nil

### UnsetActualHours
`func (o *ProjectWorkplanProjectPhase) UnsetActualHours()`

UnsetActualHours ensures that no value is present for ActualHours, not even an explicit nil
### GetBillableHours

`func (o *ProjectWorkplanProjectPhase) GetBillableHours() float64`

GetBillableHours returns the BillableHours field if non-nil, zero value otherwise.

### GetBillableHoursOk

`func (o *ProjectWorkplanProjectPhase) GetBillableHoursOk() (*float64, bool)`

GetBillableHoursOk returns a tuple with the BillableHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableHours

`func (o *ProjectWorkplanProjectPhase) SetBillableHours(v float64)`

SetBillableHours sets BillableHours field to given value.

### HasBillableHours

`func (o *ProjectWorkplanProjectPhase) HasBillableHours() bool`

HasBillableHours returns a boolean if a field has been set.

### SetBillableHoursNil

`func (o *ProjectWorkplanProjectPhase) SetBillableHoursNil(b bool)`

 SetBillableHoursNil sets the value for BillableHours to be an explicit nil

### UnsetBillableHours
`func (o *ProjectWorkplanProjectPhase) UnsetBillableHours()`

UnsetBillableHours ensures that no value is present for BillableHours, not even an explicit nil
### GetScheduledHours

`func (o *ProjectWorkplanProjectPhase) GetScheduledHours() float64`

GetScheduledHours returns the ScheduledHours field if non-nil, zero value otherwise.

### GetScheduledHoursOk

`func (o *ProjectWorkplanProjectPhase) GetScheduledHoursOk() (*float64, bool)`

GetScheduledHoursOk returns a tuple with the ScheduledHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledHours

`func (o *ProjectWorkplanProjectPhase) SetScheduledHours(v float64)`

SetScheduledHours sets ScheduledHours field to given value.

### HasScheduledHours

`func (o *ProjectWorkplanProjectPhase) HasScheduledHours() bool`

HasScheduledHours returns a boolean if a field has been set.

### SetScheduledHoursNil

`func (o *ProjectWorkplanProjectPhase) SetScheduledHoursNil(b bool)`

 SetScheduledHoursNil sets the value for ScheduledHours to be an explicit nil

### UnsetScheduledHours
`func (o *ProjectWorkplanProjectPhase) UnsetScheduledHours()`

UnsetScheduledHours ensures that no value is present for ScheduledHours, not even an explicit nil
### GetScheduledStart

`func (o *ProjectWorkplanProjectPhase) GetScheduledStart() string`

GetScheduledStart returns the ScheduledStart field if non-nil, zero value otherwise.

### GetScheduledStartOk

`func (o *ProjectWorkplanProjectPhase) GetScheduledStartOk() (*string, bool)`

GetScheduledStartOk returns a tuple with the ScheduledStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStart

`func (o *ProjectWorkplanProjectPhase) SetScheduledStart(v string)`

SetScheduledStart sets ScheduledStart field to given value.

### HasScheduledStart

`func (o *ProjectWorkplanProjectPhase) HasScheduledStart() bool`

HasScheduledStart returns a boolean if a field has been set.

### GetScheduledEnd

`func (o *ProjectWorkplanProjectPhase) GetScheduledEnd() string`

GetScheduledEnd returns the ScheduledEnd field if non-nil, zero value otherwise.

### GetScheduledEndOk

`func (o *ProjectWorkplanProjectPhase) GetScheduledEndOk() (*string, bool)`

GetScheduledEndOk returns a tuple with the ScheduledEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEnd

`func (o *ProjectWorkplanProjectPhase) SetScheduledEnd(v string)`

SetScheduledEnd sets ScheduledEnd field to given value.

### HasScheduledEnd

`func (o *ProjectWorkplanProjectPhase) HasScheduledEnd() bool`

HasScheduledEnd returns a boolean if a field has been set.

### GetScheduledDuration

`func (o *ProjectWorkplanProjectPhase) GetScheduledDuration() int32`

GetScheduledDuration returns the ScheduledDuration field if non-nil, zero value otherwise.

### GetScheduledDurationOk

`func (o *ProjectWorkplanProjectPhase) GetScheduledDurationOk() (*int32, bool)`

GetScheduledDurationOk returns a tuple with the ScheduledDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDuration

`func (o *ProjectWorkplanProjectPhase) SetScheduledDuration(v int32)`

SetScheduledDuration sets ScheduledDuration field to given value.

### HasScheduledDuration

`func (o *ProjectWorkplanProjectPhase) HasScheduledDuration() bool`

HasScheduledDuration returns a boolean if a field has been set.

### SetScheduledDurationNil

`func (o *ProjectWorkplanProjectPhase) SetScheduledDurationNil(b bool)`

 SetScheduledDurationNil sets the value for ScheduledDuration to be an explicit nil

### UnsetScheduledDuration
`func (o *ProjectWorkplanProjectPhase) UnsetScheduledDuration()`

UnsetScheduledDuration ensures that no value is present for ScheduledDuration, not even an explicit nil
### GetBillPhaseSeparately

`func (o *ProjectWorkplanProjectPhase) GetBillPhaseSeparately() bool`

GetBillPhaseSeparately returns the BillPhaseSeparately field if non-nil, zero value otherwise.

### GetBillPhaseSeparatelyOk

`func (o *ProjectWorkplanProjectPhase) GetBillPhaseSeparatelyOk() (*bool, bool)`

GetBillPhaseSeparatelyOk returns a tuple with the BillPhaseSeparately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillPhaseSeparately

`func (o *ProjectWorkplanProjectPhase) SetBillPhaseSeparately(v bool)`

SetBillPhaseSeparately sets BillPhaseSeparately field to given value.

### HasBillPhaseSeparately

`func (o *ProjectWorkplanProjectPhase) HasBillPhaseSeparately() bool`

HasBillPhaseSeparately returns a boolean if a field has been set.

### SetBillPhaseSeparatelyNil

`func (o *ProjectWorkplanProjectPhase) SetBillPhaseSeparatelyNil(b bool)`

 SetBillPhaseSeparatelyNil sets the value for BillPhaseSeparately to be an explicit nil

### UnsetBillPhaseSeparately
`func (o *ProjectWorkplanProjectPhase) UnsetBillPhaseSeparately()`

UnsetBillPhaseSeparately ensures that no value is present for BillPhaseSeparately, not even an explicit nil
### GetInfo

`func (o *ProjectWorkplanProjectPhase) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectWorkplanProjectPhase) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectWorkplanProjectPhase) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectWorkplanProjectPhase) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *ProjectWorkplanProjectPhase) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProjectWorkplanProjectPhase) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProjectWorkplanProjectPhase) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProjectWorkplanProjectPhase) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


