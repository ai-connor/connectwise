# ProjectTemplateTicket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectTemplateId** | Pointer to **NullableInt32** |  | [optional] 
**ProjectTemplatePhaseId** | Pointer to **NullableInt32** |  | [optional] 
**LineNumber** | Pointer to **NullableFloat64** |  | [optional] 
**Description** | **string** |  Max length: 100; | 
**ConnectWiseId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableInt32** |  | [optional] 
**ParentConnectWiseId** | Pointer to **string** |  | [optional] 
**ProjectTemplatePhaseCwId** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**InternalAnalysis** | Pointer to **string** |  | [optional] 
**Resolution** | Pointer to **string** |  | [optional] 
**BudgetHours** | Pointer to **NullableFloat64** |  | [optional] 
**Duration** | Pointer to **NullableInt32** |  | [optional] 
**WbsCode** | Pointer to **string** |  Max length: 50; | [optional] 
**BillSeparatelyFlag** | Pointer to **NullableBool** |  | [optional] 
**MarkAsMilestoneFlag** | Pointer to **NullableBool** |  | [optional] 
**RecordType** | Pointer to **string** |  Max length: 1; | [optional] 
**PmTmpProjectRecID** | Pointer to **NullableInt32** |  | [optional] 
**PredecessorType** | Pointer to **NullableString** |  | [optional] 
**PredecessorId** | Pointer to **NullableInt32** |  | [optional] 
**PredecessorClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**LagDays** | Pointer to **NullableInt32** |  | [optional] 
**LagNonworkingDaysFlag** | Pointer to **NullableBool** |  | [optional] 
**Priority** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**Source** | Pointer to [**ServiceSourceReference**](ServiceSourceReference.md) |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectTemplateTicket

`func NewProjectTemplateTicket(description string, ) *ProjectTemplateTicket`

NewProjectTemplateTicket instantiates a new ProjectTemplateTicket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTemplateTicketWithDefaults

`func NewProjectTemplateTicketWithDefaults() *ProjectTemplateTicket`

NewProjectTemplateTicketWithDefaults instantiates a new ProjectTemplateTicket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTemplateTicket) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTemplateTicket) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTemplateTicket) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTemplateTicket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectTemplateId

`func (o *ProjectTemplateTicket) GetProjectTemplateId() int32`

GetProjectTemplateId returns the ProjectTemplateId field if non-nil, zero value otherwise.

### GetProjectTemplateIdOk

`func (o *ProjectTemplateTicket) GetProjectTemplateIdOk() (*int32, bool)`

GetProjectTemplateIdOk returns a tuple with the ProjectTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTemplateId

`func (o *ProjectTemplateTicket) SetProjectTemplateId(v int32)`

SetProjectTemplateId sets ProjectTemplateId field to given value.

### HasProjectTemplateId

`func (o *ProjectTemplateTicket) HasProjectTemplateId() bool`

HasProjectTemplateId returns a boolean if a field has been set.

### SetProjectTemplateIdNil

`func (o *ProjectTemplateTicket) SetProjectTemplateIdNil(b bool)`

 SetProjectTemplateIdNil sets the value for ProjectTemplateId to be an explicit nil

### UnsetProjectTemplateId
`func (o *ProjectTemplateTicket) UnsetProjectTemplateId()`

UnsetProjectTemplateId ensures that no value is present for ProjectTemplateId, not even an explicit nil
### GetProjectTemplatePhaseId

`func (o *ProjectTemplateTicket) GetProjectTemplatePhaseId() int32`

GetProjectTemplatePhaseId returns the ProjectTemplatePhaseId field if non-nil, zero value otherwise.

### GetProjectTemplatePhaseIdOk

`func (o *ProjectTemplateTicket) GetProjectTemplatePhaseIdOk() (*int32, bool)`

GetProjectTemplatePhaseIdOk returns a tuple with the ProjectTemplatePhaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTemplatePhaseId

`func (o *ProjectTemplateTicket) SetProjectTemplatePhaseId(v int32)`

SetProjectTemplatePhaseId sets ProjectTemplatePhaseId field to given value.

### HasProjectTemplatePhaseId

`func (o *ProjectTemplateTicket) HasProjectTemplatePhaseId() bool`

HasProjectTemplatePhaseId returns a boolean if a field has been set.

### SetProjectTemplatePhaseIdNil

`func (o *ProjectTemplateTicket) SetProjectTemplatePhaseIdNil(b bool)`

 SetProjectTemplatePhaseIdNil sets the value for ProjectTemplatePhaseId to be an explicit nil

### UnsetProjectTemplatePhaseId
`func (o *ProjectTemplateTicket) UnsetProjectTemplatePhaseId()`

UnsetProjectTemplatePhaseId ensures that no value is present for ProjectTemplatePhaseId, not even an explicit nil
### GetLineNumber

`func (o *ProjectTemplateTicket) GetLineNumber() float64`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *ProjectTemplateTicket) GetLineNumberOk() (*float64, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *ProjectTemplateTicket) SetLineNumber(v float64)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *ProjectTemplateTicket) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### SetLineNumberNil

`func (o *ProjectTemplateTicket) SetLineNumberNil(b bool)`

 SetLineNumberNil sets the value for LineNumber to be an explicit nil

### UnsetLineNumber
`func (o *ProjectTemplateTicket) UnsetLineNumber()`

UnsetLineNumber ensures that no value is present for LineNumber, not even an explicit nil
### GetDescription

`func (o *ProjectTemplateTicket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectTemplateTicket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectTemplateTicket) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetConnectWiseId

`func (o *ProjectTemplateTicket) GetConnectWiseId() string`

GetConnectWiseId returns the ConnectWiseId field if non-nil, zero value otherwise.

### GetConnectWiseIdOk

`func (o *ProjectTemplateTicket) GetConnectWiseIdOk() (*string, bool)`

GetConnectWiseIdOk returns a tuple with the ConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseId

`func (o *ProjectTemplateTicket) SetConnectWiseId(v string)`

SetConnectWiseId sets ConnectWiseId field to given value.

### HasConnectWiseId

`func (o *ProjectTemplateTicket) HasConnectWiseId() bool`

HasConnectWiseId returns a boolean if a field has been set.

### GetParentId

`func (o *ProjectTemplateTicket) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ProjectTemplateTicket) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ProjectTemplateTicket) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ProjectTemplateTicket) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ProjectTemplateTicket) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ProjectTemplateTicket) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetParentConnectWiseId

`func (o *ProjectTemplateTicket) GetParentConnectWiseId() string`

GetParentConnectWiseId returns the ParentConnectWiseId field if non-nil, zero value otherwise.

### GetParentConnectWiseIdOk

`func (o *ProjectTemplateTicket) GetParentConnectWiseIdOk() (*string, bool)`

GetParentConnectWiseIdOk returns a tuple with the ParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectWiseId

`func (o *ProjectTemplateTicket) SetParentConnectWiseId(v string)`

SetParentConnectWiseId sets ParentConnectWiseId field to given value.

### HasParentConnectWiseId

`func (o *ProjectTemplateTicket) HasParentConnectWiseId() bool`

HasParentConnectWiseId returns a boolean if a field has been set.

### GetProjectTemplatePhaseCwId

`func (o *ProjectTemplateTicket) GetProjectTemplatePhaseCwId() string`

GetProjectTemplatePhaseCwId returns the ProjectTemplatePhaseCwId field if non-nil, zero value otherwise.

### GetProjectTemplatePhaseCwIdOk

`func (o *ProjectTemplateTicket) GetProjectTemplatePhaseCwIdOk() (*string, bool)`

GetProjectTemplatePhaseCwIdOk returns a tuple with the ProjectTemplatePhaseCwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTemplatePhaseCwId

`func (o *ProjectTemplateTicket) SetProjectTemplatePhaseCwId(v string)`

SetProjectTemplatePhaseCwId sets ProjectTemplatePhaseCwId field to given value.

### HasProjectTemplatePhaseCwId

`func (o *ProjectTemplateTicket) HasProjectTemplatePhaseCwId() bool`

HasProjectTemplatePhaseCwId returns a boolean if a field has been set.

### GetNotes

`func (o *ProjectTemplateTicket) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ProjectTemplateTicket) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ProjectTemplateTicket) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ProjectTemplateTicket) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetInternalAnalysis

`func (o *ProjectTemplateTicket) GetInternalAnalysis() string`

GetInternalAnalysis returns the InternalAnalysis field if non-nil, zero value otherwise.

### GetInternalAnalysisOk

`func (o *ProjectTemplateTicket) GetInternalAnalysisOk() (*string, bool)`

GetInternalAnalysisOk returns a tuple with the InternalAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAnalysis

`func (o *ProjectTemplateTicket) SetInternalAnalysis(v string)`

SetInternalAnalysis sets InternalAnalysis field to given value.

### HasInternalAnalysis

`func (o *ProjectTemplateTicket) HasInternalAnalysis() bool`

HasInternalAnalysis returns a boolean if a field has been set.

### GetResolution

`func (o *ProjectTemplateTicket) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ProjectTemplateTicket) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ProjectTemplateTicket) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ProjectTemplateTicket) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetBudgetHours

`func (o *ProjectTemplateTicket) GetBudgetHours() float64`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *ProjectTemplateTicket) GetBudgetHoursOk() (*float64, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *ProjectTemplateTicket) SetBudgetHours(v float64)`

SetBudgetHours sets BudgetHours field to given value.

### HasBudgetHours

`func (o *ProjectTemplateTicket) HasBudgetHours() bool`

HasBudgetHours returns a boolean if a field has been set.

### SetBudgetHoursNil

`func (o *ProjectTemplateTicket) SetBudgetHoursNil(b bool)`

 SetBudgetHoursNil sets the value for BudgetHours to be an explicit nil

### UnsetBudgetHours
`func (o *ProjectTemplateTicket) UnsetBudgetHours()`

UnsetBudgetHours ensures that no value is present for BudgetHours, not even an explicit nil
### GetDuration

`func (o *ProjectTemplateTicket) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ProjectTemplateTicket) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ProjectTemplateTicket) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ProjectTemplateTicket) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ProjectTemplateTicket) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ProjectTemplateTicket) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetWbsCode

`func (o *ProjectTemplateTicket) GetWbsCode() string`

GetWbsCode returns the WbsCode field if non-nil, zero value otherwise.

### GetWbsCodeOk

`func (o *ProjectTemplateTicket) GetWbsCodeOk() (*string, bool)`

GetWbsCodeOk returns a tuple with the WbsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbsCode

`func (o *ProjectTemplateTicket) SetWbsCode(v string)`

SetWbsCode sets WbsCode field to given value.

### HasWbsCode

`func (o *ProjectTemplateTicket) HasWbsCode() bool`

HasWbsCode returns a boolean if a field has been set.

### GetBillSeparatelyFlag

`func (o *ProjectTemplateTicket) GetBillSeparatelyFlag() bool`

GetBillSeparatelyFlag returns the BillSeparatelyFlag field if non-nil, zero value otherwise.

### GetBillSeparatelyFlagOk

`func (o *ProjectTemplateTicket) GetBillSeparatelyFlagOk() (*bool, bool)`

GetBillSeparatelyFlagOk returns a tuple with the BillSeparatelyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillSeparatelyFlag

`func (o *ProjectTemplateTicket) SetBillSeparatelyFlag(v bool)`

SetBillSeparatelyFlag sets BillSeparatelyFlag field to given value.

### HasBillSeparatelyFlag

`func (o *ProjectTemplateTicket) HasBillSeparatelyFlag() bool`

HasBillSeparatelyFlag returns a boolean if a field has been set.

### SetBillSeparatelyFlagNil

`func (o *ProjectTemplateTicket) SetBillSeparatelyFlagNil(b bool)`

 SetBillSeparatelyFlagNil sets the value for BillSeparatelyFlag to be an explicit nil

### UnsetBillSeparatelyFlag
`func (o *ProjectTemplateTicket) UnsetBillSeparatelyFlag()`

UnsetBillSeparatelyFlag ensures that no value is present for BillSeparatelyFlag, not even an explicit nil
### GetMarkAsMilestoneFlag

`func (o *ProjectTemplateTicket) GetMarkAsMilestoneFlag() bool`

GetMarkAsMilestoneFlag returns the MarkAsMilestoneFlag field if non-nil, zero value otherwise.

### GetMarkAsMilestoneFlagOk

`func (o *ProjectTemplateTicket) GetMarkAsMilestoneFlagOk() (*bool, bool)`

GetMarkAsMilestoneFlagOk returns a tuple with the MarkAsMilestoneFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAsMilestoneFlag

`func (o *ProjectTemplateTicket) SetMarkAsMilestoneFlag(v bool)`

SetMarkAsMilestoneFlag sets MarkAsMilestoneFlag field to given value.

### HasMarkAsMilestoneFlag

`func (o *ProjectTemplateTicket) HasMarkAsMilestoneFlag() bool`

HasMarkAsMilestoneFlag returns a boolean if a field has been set.

### SetMarkAsMilestoneFlagNil

`func (o *ProjectTemplateTicket) SetMarkAsMilestoneFlagNil(b bool)`

 SetMarkAsMilestoneFlagNil sets the value for MarkAsMilestoneFlag to be an explicit nil

### UnsetMarkAsMilestoneFlag
`func (o *ProjectTemplateTicket) UnsetMarkAsMilestoneFlag()`

UnsetMarkAsMilestoneFlag ensures that no value is present for MarkAsMilestoneFlag, not even an explicit nil
### GetRecordType

`func (o *ProjectTemplateTicket) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ProjectTemplateTicket) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ProjectTemplateTicket) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ProjectTemplateTicket) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPmTmpProjectRecID

`func (o *ProjectTemplateTicket) GetPmTmpProjectRecID() int32`

GetPmTmpProjectRecID returns the PmTmpProjectRecID field if non-nil, zero value otherwise.

### GetPmTmpProjectRecIDOk

`func (o *ProjectTemplateTicket) GetPmTmpProjectRecIDOk() (*int32, bool)`

GetPmTmpProjectRecIDOk returns a tuple with the PmTmpProjectRecID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmTmpProjectRecID

`func (o *ProjectTemplateTicket) SetPmTmpProjectRecID(v int32)`

SetPmTmpProjectRecID sets PmTmpProjectRecID field to given value.

### HasPmTmpProjectRecID

`func (o *ProjectTemplateTicket) HasPmTmpProjectRecID() bool`

HasPmTmpProjectRecID returns a boolean if a field has been set.

### SetPmTmpProjectRecIDNil

`func (o *ProjectTemplateTicket) SetPmTmpProjectRecIDNil(b bool)`

 SetPmTmpProjectRecIDNil sets the value for PmTmpProjectRecID to be an explicit nil

### UnsetPmTmpProjectRecID
`func (o *ProjectTemplateTicket) UnsetPmTmpProjectRecID()`

UnsetPmTmpProjectRecID ensures that no value is present for PmTmpProjectRecID, not even an explicit nil
### GetPredecessorType

`func (o *ProjectTemplateTicket) GetPredecessorType() string`

GetPredecessorType returns the PredecessorType field if non-nil, zero value otherwise.

### GetPredecessorTypeOk

`func (o *ProjectTemplateTicket) GetPredecessorTypeOk() (*string, bool)`

GetPredecessorTypeOk returns a tuple with the PredecessorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorType

`func (o *ProjectTemplateTicket) SetPredecessorType(v string)`

SetPredecessorType sets PredecessorType field to given value.

### HasPredecessorType

`func (o *ProjectTemplateTicket) HasPredecessorType() bool`

HasPredecessorType returns a boolean if a field has been set.

### SetPredecessorTypeNil

`func (o *ProjectTemplateTicket) SetPredecessorTypeNil(b bool)`

 SetPredecessorTypeNil sets the value for PredecessorType to be an explicit nil

### UnsetPredecessorType
`func (o *ProjectTemplateTicket) UnsetPredecessorType()`

UnsetPredecessorType ensures that no value is present for PredecessorType, not even an explicit nil
### GetPredecessorId

`func (o *ProjectTemplateTicket) GetPredecessorId() int32`

GetPredecessorId returns the PredecessorId field if non-nil, zero value otherwise.

### GetPredecessorIdOk

`func (o *ProjectTemplateTicket) GetPredecessorIdOk() (*int32, bool)`

GetPredecessorIdOk returns a tuple with the PredecessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorId

`func (o *ProjectTemplateTicket) SetPredecessorId(v int32)`

SetPredecessorId sets PredecessorId field to given value.

### HasPredecessorId

`func (o *ProjectTemplateTicket) HasPredecessorId() bool`

HasPredecessorId returns a boolean if a field has been set.

### SetPredecessorIdNil

`func (o *ProjectTemplateTicket) SetPredecessorIdNil(b bool)`

 SetPredecessorIdNil sets the value for PredecessorId to be an explicit nil

### UnsetPredecessorId
`func (o *ProjectTemplateTicket) UnsetPredecessorId()`

UnsetPredecessorId ensures that no value is present for PredecessorId, not even an explicit nil
### GetPredecessorClosedFlag

`func (o *ProjectTemplateTicket) GetPredecessorClosedFlag() bool`

GetPredecessorClosedFlag returns the PredecessorClosedFlag field if non-nil, zero value otherwise.

### GetPredecessorClosedFlagOk

`func (o *ProjectTemplateTicket) GetPredecessorClosedFlagOk() (*bool, bool)`

GetPredecessorClosedFlagOk returns a tuple with the PredecessorClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorClosedFlag

`func (o *ProjectTemplateTicket) SetPredecessorClosedFlag(v bool)`

SetPredecessorClosedFlag sets PredecessorClosedFlag field to given value.

### HasPredecessorClosedFlag

`func (o *ProjectTemplateTicket) HasPredecessorClosedFlag() bool`

HasPredecessorClosedFlag returns a boolean if a field has been set.

### SetPredecessorClosedFlagNil

`func (o *ProjectTemplateTicket) SetPredecessorClosedFlagNil(b bool)`

 SetPredecessorClosedFlagNil sets the value for PredecessorClosedFlag to be an explicit nil

### UnsetPredecessorClosedFlag
`func (o *ProjectTemplateTicket) UnsetPredecessorClosedFlag()`

UnsetPredecessorClosedFlag ensures that no value is present for PredecessorClosedFlag, not even an explicit nil
### GetLagDays

`func (o *ProjectTemplateTicket) GetLagDays() int32`

GetLagDays returns the LagDays field if non-nil, zero value otherwise.

### GetLagDaysOk

`func (o *ProjectTemplateTicket) GetLagDaysOk() (*int32, bool)`

GetLagDaysOk returns a tuple with the LagDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagDays

`func (o *ProjectTemplateTicket) SetLagDays(v int32)`

SetLagDays sets LagDays field to given value.

### HasLagDays

`func (o *ProjectTemplateTicket) HasLagDays() bool`

HasLagDays returns a boolean if a field has been set.

### SetLagDaysNil

`func (o *ProjectTemplateTicket) SetLagDaysNil(b bool)`

 SetLagDaysNil sets the value for LagDays to be an explicit nil

### UnsetLagDays
`func (o *ProjectTemplateTicket) UnsetLagDays()`

UnsetLagDays ensures that no value is present for LagDays, not even an explicit nil
### GetLagNonworkingDaysFlag

`func (o *ProjectTemplateTicket) GetLagNonworkingDaysFlag() bool`

GetLagNonworkingDaysFlag returns the LagNonworkingDaysFlag field if non-nil, zero value otherwise.

### GetLagNonworkingDaysFlagOk

`func (o *ProjectTemplateTicket) GetLagNonworkingDaysFlagOk() (*bool, bool)`

GetLagNonworkingDaysFlagOk returns a tuple with the LagNonworkingDaysFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagNonworkingDaysFlag

`func (o *ProjectTemplateTicket) SetLagNonworkingDaysFlag(v bool)`

SetLagNonworkingDaysFlag sets LagNonworkingDaysFlag field to given value.

### HasLagNonworkingDaysFlag

`func (o *ProjectTemplateTicket) HasLagNonworkingDaysFlag() bool`

HasLagNonworkingDaysFlag returns a boolean if a field has been set.

### SetLagNonworkingDaysFlagNil

`func (o *ProjectTemplateTicket) SetLagNonworkingDaysFlagNil(b bool)`

 SetLagNonworkingDaysFlagNil sets the value for LagNonworkingDaysFlag to be an explicit nil

### UnsetLagNonworkingDaysFlag
`func (o *ProjectTemplateTicket) UnsetLagNonworkingDaysFlag()`

UnsetLagNonworkingDaysFlag ensures that no value is present for LagNonworkingDaysFlag, not even an explicit nil
### GetPriority

`func (o *ProjectTemplateTicket) GetPriority() PriorityReference`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ProjectTemplateTicket) GetPriorityOk() (*PriorityReference, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ProjectTemplateTicket) SetPriority(v PriorityReference)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ProjectTemplateTicket) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSource

`func (o *ProjectTemplateTicket) GetSource() ServiceSourceReference`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProjectTemplateTicket) GetSourceOk() (*ServiceSourceReference, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProjectTemplateTicket) SetSource(v ServiceSourceReference)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProjectTemplateTicket) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetWorkRole

`func (o *ProjectTemplateTicket) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *ProjectTemplateTicket) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *ProjectTemplateTicket) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *ProjectTemplateTicket) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *ProjectTemplateTicket) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *ProjectTemplateTicket) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *ProjectTemplateTicket) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *ProjectTemplateTicket) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetInfo

`func (o *ProjectTemplateTicket) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectTemplateTicket) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectTemplateTicket) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectTemplateTicket) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


