# BoardAutoTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | [**ServiceTypeReference**](ServiceTypeReference.md) |  | 
**Subtype** | [**ServiceSubTypeReference**](ServiceSubTypeReference.md) |  | 
**Item** | [**ServiceItemReference**](ServiceItemReference.md) |  | 
**ServiceTemplate** | [**ServiceTemplateReference**](ServiceTemplateReference.md) |  | 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**SummarySetting** | Pointer to **NullableString** |  | [optional] 
**DiscussionSetting** | Pointer to **NullableString** |  | [optional] 
**InternalAnalysisSetting** | Pointer to **NullableString** |  | [optional] 
**ResolutionSetting** | Pointer to **NullableString** |  | [optional] 
**TasksSetting** | Pointer to **NullableString** |  | [optional] 
**DocumentsSetting** | Pointer to **NullableString** |  | [optional] 
**ResourcesSetting** | Pointer to **NullableString** |  | [optional] 
**BudgetHoursSetting** | Pointer to **NullableString** |  | [optional] 
**FinanceInformationSetting** | Pointer to **NullableString** |  | [optional] 
**SendNotesAsEmailSetting** | Pointer to **NullableString** |  | [optional] 
**ImpactUrgencySetting** | Pointer to **NullableString** |  | [optional] 
**TemplatePrioritySetting** | Pointer to **NullableString** |  | [optional] 
**AutoApplyFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardAutoTemplate

`func NewBoardAutoTemplate(type_ ServiceTypeReference, subtype ServiceSubTypeReference, item ServiceItemReference, serviceTemplate ServiceTemplateReference, ) *BoardAutoTemplate`

NewBoardAutoTemplate instantiates a new BoardAutoTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardAutoTemplateWithDefaults

`func NewBoardAutoTemplateWithDefaults() *BoardAutoTemplate`

NewBoardAutoTemplateWithDefaults instantiates a new BoardAutoTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardAutoTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardAutoTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardAutoTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardAutoTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BoardAutoTemplate) GetType() ServiceTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BoardAutoTemplate) GetTypeOk() (*ServiceTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BoardAutoTemplate) SetType(v ServiceTypeReference)`

SetType sets Type field to given value.


### GetSubtype

`func (o *BoardAutoTemplate) GetSubtype() ServiceSubTypeReference`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BoardAutoTemplate) GetSubtypeOk() (*ServiceSubTypeReference, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BoardAutoTemplate) SetSubtype(v ServiceSubTypeReference)`

SetSubtype sets Subtype field to given value.


### GetItem

`func (o *BoardAutoTemplate) GetItem() ServiceItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *BoardAutoTemplate) GetItemOk() (*ServiceItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *BoardAutoTemplate) SetItem(v ServiceItemReference)`

SetItem sets Item field to given value.


### GetServiceTemplate

`func (o *BoardAutoTemplate) GetServiceTemplate() ServiceTemplateReference`

GetServiceTemplate returns the ServiceTemplate field if non-nil, zero value otherwise.

### GetServiceTemplateOk

`func (o *BoardAutoTemplate) GetServiceTemplateOk() (*ServiceTemplateReference, bool)`

GetServiceTemplateOk returns a tuple with the ServiceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTemplate

`func (o *BoardAutoTemplate) SetServiceTemplate(v ServiceTemplateReference)`

SetServiceTemplate sets ServiceTemplate field to given value.


### GetBoard

`func (o *BoardAutoTemplate) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *BoardAutoTemplate) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *BoardAutoTemplate) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *BoardAutoTemplate) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetSummarySetting

`func (o *BoardAutoTemplate) GetSummarySetting() string`

GetSummarySetting returns the SummarySetting field if non-nil, zero value otherwise.

### GetSummarySettingOk

`func (o *BoardAutoTemplate) GetSummarySettingOk() (*string, bool)`

GetSummarySettingOk returns a tuple with the SummarySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummarySetting

`func (o *BoardAutoTemplate) SetSummarySetting(v string)`

SetSummarySetting sets SummarySetting field to given value.

### HasSummarySetting

`func (o *BoardAutoTemplate) HasSummarySetting() bool`

HasSummarySetting returns a boolean if a field has been set.

### SetSummarySettingNil

`func (o *BoardAutoTemplate) SetSummarySettingNil(b bool)`

 SetSummarySettingNil sets the value for SummarySetting to be an explicit nil

### UnsetSummarySetting
`func (o *BoardAutoTemplate) UnsetSummarySetting()`

UnsetSummarySetting ensures that no value is present for SummarySetting, not even an explicit nil
### GetDiscussionSetting

`func (o *BoardAutoTemplate) GetDiscussionSetting() string`

GetDiscussionSetting returns the DiscussionSetting field if non-nil, zero value otherwise.

### GetDiscussionSettingOk

`func (o *BoardAutoTemplate) GetDiscussionSettingOk() (*string, bool)`

GetDiscussionSettingOk returns a tuple with the DiscussionSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscussionSetting

`func (o *BoardAutoTemplate) SetDiscussionSetting(v string)`

SetDiscussionSetting sets DiscussionSetting field to given value.

### HasDiscussionSetting

`func (o *BoardAutoTemplate) HasDiscussionSetting() bool`

HasDiscussionSetting returns a boolean if a field has been set.

### SetDiscussionSettingNil

`func (o *BoardAutoTemplate) SetDiscussionSettingNil(b bool)`

 SetDiscussionSettingNil sets the value for DiscussionSetting to be an explicit nil

### UnsetDiscussionSetting
`func (o *BoardAutoTemplate) UnsetDiscussionSetting()`

UnsetDiscussionSetting ensures that no value is present for DiscussionSetting, not even an explicit nil
### GetInternalAnalysisSetting

`func (o *BoardAutoTemplate) GetInternalAnalysisSetting() string`

GetInternalAnalysisSetting returns the InternalAnalysisSetting field if non-nil, zero value otherwise.

### GetInternalAnalysisSettingOk

`func (o *BoardAutoTemplate) GetInternalAnalysisSettingOk() (*string, bool)`

GetInternalAnalysisSettingOk returns a tuple with the InternalAnalysisSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAnalysisSetting

`func (o *BoardAutoTemplate) SetInternalAnalysisSetting(v string)`

SetInternalAnalysisSetting sets InternalAnalysisSetting field to given value.

### HasInternalAnalysisSetting

`func (o *BoardAutoTemplate) HasInternalAnalysisSetting() bool`

HasInternalAnalysisSetting returns a boolean if a field has been set.

### SetInternalAnalysisSettingNil

`func (o *BoardAutoTemplate) SetInternalAnalysisSettingNil(b bool)`

 SetInternalAnalysisSettingNil sets the value for InternalAnalysisSetting to be an explicit nil

### UnsetInternalAnalysisSetting
`func (o *BoardAutoTemplate) UnsetInternalAnalysisSetting()`

UnsetInternalAnalysisSetting ensures that no value is present for InternalAnalysisSetting, not even an explicit nil
### GetResolutionSetting

`func (o *BoardAutoTemplate) GetResolutionSetting() string`

GetResolutionSetting returns the ResolutionSetting field if non-nil, zero value otherwise.

### GetResolutionSettingOk

`func (o *BoardAutoTemplate) GetResolutionSettingOk() (*string, bool)`

GetResolutionSettingOk returns a tuple with the ResolutionSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionSetting

`func (o *BoardAutoTemplate) SetResolutionSetting(v string)`

SetResolutionSetting sets ResolutionSetting field to given value.

### HasResolutionSetting

`func (o *BoardAutoTemplate) HasResolutionSetting() bool`

HasResolutionSetting returns a boolean if a field has been set.

### SetResolutionSettingNil

`func (o *BoardAutoTemplate) SetResolutionSettingNil(b bool)`

 SetResolutionSettingNil sets the value for ResolutionSetting to be an explicit nil

### UnsetResolutionSetting
`func (o *BoardAutoTemplate) UnsetResolutionSetting()`

UnsetResolutionSetting ensures that no value is present for ResolutionSetting, not even an explicit nil
### GetTasksSetting

`func (o *BoardAutoTemplate) GetTasksSetting() string`

GetTasksSetting returns the TasksSetting field if non-nil, zero value otherwise.

### GetTasksSettingOk

`func (o *BoardAutoTemplate) GetTasksSettingOk() (*string, bool)`

GetTasksSettingOk returns a tuple with the TasksSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksSetting

`func (o *BoardAutoTemplate) SetTasksSetting(v string)`

SetTasksSetting sets TasksSetting field to given value.

### HasTasksSetting

`func (o *BoardAutoTemplate) HasTasksSetting() bool`

HasTasksSetting returns a boolean if a field has been set.

### SetTasksSettingNil

`func (o *BoardAutoTemplate) SetTasksSettingNil(b bool)`

 SetTasksSettingNil sets the value for TasksSetting to be an explicit nil

### UnsetTasksSetting
`func (o *BoardAutoTemplate) UnsetTasksSetting()`

UnsetTasksSetting ensures that no value is present for TasksSetting, not even an explicit nil
### GetDocumentsSetting

`func (o *BoardAutoTemplate) GetDocumentsSetting() string`

GetDocumentsSetting returns the DocumentsSetting field if non-nil, zero value otherwise.

### GetDocumentsSettingOk

`func (o *BoardAutoTemplate) GetDocumentsSettingOk() (*string, bool)`

GetDocumentsSettingOk returns a tuple with the DocumentsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentsSetting

`func (o *BoardAutoTemplate) SetDocumentsSetting(v string)`

SetDocumentsSetting sets DocumentsSetting field to given value.

### HasDocumentsSetting

`func (o *BoardAutoTemplate) HasDocumentsSetting() bool`

HasDocumentsSetting returns a boolean if a field has been set.

### SetDocumentsSettingNil

`func (o *BoardAutoTemplate) SetDocumentsSettingNil(b bool)`

 SetDocumentsSettingNil sets the value for DocumentsSetting to be an explicit nil

### UnsetDocumentsSetting
`func (o *BoardAutoTemplate) UnsetDocumentsSetting()`

UnsetDocumentsSetting ensures that no value is present for DocumentsSetting, not even an explicit nil
### GetResourcesSetting

`func (o *BoardAutoTemplate) GetResourcesSetting() string`

GetResourcesSetting returns the ResourcesSetting field if non-nil, zero value otherwise.

### GetResourcesSettingOk

`func (o *BoardAutoTemplate) GetResourcesSettingOk() (*string, bool)`

GetResourcesSettingOk returns a tuple with the ResourcesSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesSetting

`func (o *BoardAutoTemplate) SetResourcesSetting(v string)`

SetResourcesSetting sets ResourcesSetting field to given value.

### HasResourcesSetting

`func (o *BoardAutoTemplate) HasResourcesSetting() bool`

HasResourcesSetting returns a boolean if a field has been set.

### SetResourcesSettingNil

`func (o *BoardAutoTemplate) SetResourcesSettingNil(b bool)`

 SetResourcesSettingNil sets the value for ResourcesSetting to be an explicit nil

### UnsetResourcesSetting
`func (o *BoardAutoTemplate) UnsetResourcesSetting()`

UnsetResourcesSetting ensures that no value is present for ResourcesSetting, not even an explicit nil
### GetBudgetHoursSetting

`func (o *BoardAutoTemplate) GetBudgetHoursSetting() string`

GetBudgetHoursSetting returns the BudgetHoursSetting field if non-nil, zero value otherwise.

### GetBudgetHoursSettingOk

`func (o *BoardAutoTemplate) GetBudgetHoursSettingOk() (*string, bool)`

GetBudgetHoursSettingOk returns a tuple with the BudgetHoursSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHoursSetting

`func (o *BoardAutoTemplate) SetBudgetHoursSetting(v string)`

SetBudgetHoursSetting sets BudgetHoursSetting field to given value.

### HasBudgetHoursSetting

`func (o *BoardAutoTemplate) HasBudgetHoursSetting() bool`

HasBudgetHoursSetting returns a boolean if a field has been set.

### SetBudgetHoursSettingNil

`func (o *BoardAutoTemplate) SetBudgetHoursSettingNil(b bool)`

 SetBudgetHoursSettingNil sets the value for BudgetHoursSetting to be an explicit nil

### UnsetBudgetHoursSetting
`func (o *BoardAutoTemplate) UnsetBudgetHoursSetting()`

UnsetBudgetHoursSetting ensures that no value is present for BudgetHoursSetting, not even an explicit nil
### GetFinanceInformationSetting

`func (o *BoardAutoTemplate) GetFinanceInformationSetting() string`

GetFinanceInformationSetting returns the FinanceInformationSetting field if non-nil, zero value otherwise.

### GetFinanceInformationSettingOk

`func (o *BoardAutoTemplate) GetFinanceInformationSettingOk() (*string, bool)`

GetFinanceInformationSettingOk returns a tuple with the FinanceInformationSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinanceInformationSetting

`func (o *BoardAutoTemplate) SetFinanceInformationSetting(v string)`

SetFinanceInformationSetting sets FinanceInformationSetting field to given value.

### HasFinanceInformationSetting

`func (o *BoardAutoTemplate) HasFinanceInformationSetting() bool`

HasFinanceInformationSetting returns a boolean if a field has been set.

### SetFinanceInformationSettingNil

`func (o *BoardAutoTemplate) SetFinanceInformationSettingNil(b bool)`

 SetFinanceInformationSettingNil sets the value for FinanceInformationSetting to be an explicit nil

### UnsetFinanceInformationSetting
`func (o *BoardAutoTemplate) UnsetFinanceInformationSetting()`

UnsetFinanceInformationSetting ensures that no value is present for FinanceInformationSetting, not even an explicit nil
### GetSendNotesAsEmailSetting

`func (o *BoardAutoTemplate) GetSendNotesAsEmailSetting() string`

GetSendNotesAsEmailSetting returns the SendNotesAsEmailSetting field if non-nil, zero value otherwise.

### GetSendNotesAsEmailSettingOk

`func (o *BoardAutoTemplate) GetSendNotesAsEmailSettingOk() (*string, bool)`

GetSendNotesAsEmailSettingOk returns a tuple with the SendNotesAsEmailSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotesAsEmailSetting

`func (o *BoardAutoTemplate) SetSendNotesAsEmailSetting(v string)`

SetSendNotesAsEmailSetting sets SendNotesAsEmailSetting field to given value.

### HasSendNotesAsEmailSetting

`func (o *BoardAutoTemplate) HasSendNotesAsEmailSetting() bool`

HasSendNotesAsEmailSetting returns a boolean if a field has been set.

### SetSendNotesAsEmailSettingNil

`func (o *BoardAutoTemplate) SetSendNotesAsEmailSettingNil(b bool)`

 SetSendNotesAsEmailSettingNil sets the value for SendNotesAsEmailSetting to be an explicit nil

### UnsetSendNotesAsEmailSetting
`func (o *BoardAutoTemplate) UnsetSendNotesAsEmailSetting()`

UnsetSendNotesAsEmailSetting ensures that no value is present for SendNotesAsEmailSetting, not even an explicit nil
### GetImpactUrgencySetting

`func (o *BoardAutoTemplate) GetImpactUrgencySetting() string`

GetImpactUrgencySetting returns the ImpactUrgencySetting field if non-nil, zero value otherwise.

### GetImpactUrgencySettingOk

`func (o *BoardAutoTemplate) GetImpactUrgencySettingOk() (*string, bool)`

GetImpactUrgencySettingOk returns a tuple with the ImpactUrgencySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactUrgencySetting

`func (o *BoardAutoTemplate) SetImpactUrgencySetting(v string)`

SetImpactUrgencySetting sets ImpactUrgencySetting field to given value.

### HasImpactUrgencySetting

`func (o *BoardAutoTemplate) HasImpactUrgencySetting() bool`

HasImpactUrgencySetting returns a boolean if a field has been set.

### SetImpactUrgencySettingNil

`func (o *BoardAutoTemplate) SetImpactUrgencySettingNil(b bool)`

 SetImpactUrgencySettingNil sets the value for ImpactUrgencySetting to be an explicit nil

### UnsetImpactUrgencySetting
`func (o *BoardAutoTemplate) UnsetImpactUrgencySetting()`

UnsetImpactUrgencySetting ensures that no value is present for ImpactUrgencySetting, not even an explicit nil
### GetTemplatePrioritySetting

`func (o *BoardAutoTemplate) GetTemplatePrioritySetting() string`

GetTemplatePrioritySetting returns the TemplatePrioritySetting field if non-nil, zero value otherwise.

### GetTemplatePrioritySettingOk

`func (o *BoardAutoTemplate) GetTemplatePrioritySettingOk() (*string, bool)`

GetTemplatePrioritySettingOk returns a tuple with the TemplatePrioritySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePrioritySetting

`func (o *BoardAutoTemplate) SetTemplatePrioritySetting(v string)`

SetTemplatePrioritySetting sets TemplatePrioritySetting field to given value.

### HasTemplatePrioritySetting

`func (o *BoardAutoTemplate) HasTemplatePrioritySetting() bool`

HasTemplatePrioritySetting returns a boolean if a field has been set.

### SetTemplatePrioritySettingNil

`func (o *BoardAutoTemplate) SetTemplatePrioritySettingNil(b bool)`

 SetTemplatePrioritySettingNil sets the value for TemplatePrioritySetting to be an explicit nil

### UnsetTemplatePrioritySetting
`func (o *BoardAutoTemplate) UnsetTemplatePrioritySetting()`

UnsetTemplatePrioritySetting ensures that no value is present for TemplatePrioritySetting, not even an explicit nil
### GetAutoApplyFlag

`func (o *BoardAutoTemplate) GetAutoApplyFlag() bool`

GetAutoApplyFlag returns the AutoApplyFlag field if non-nil, zero value otherwise.

### GetAutoApplyFlagOk

`func (o *BoardAutoTemplate) GetAutoApplyFlagOk() (*bool, bool)`

GetAutoApplyFlagOk returns a tuple with the AutoApplyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApplyFlag

`func (o *BoardAutoTemplate) SetAutoApplyFlag(v bool)`

SetAutoApplyFlag sets AutoApplyFlag field to given value.

### HasAutoApplyFlag

`func (o *BoardAutoTemplate) HasAutoApplyFlag() bool`

HasAutoApplyFlag returns a boolean if a field has been set.

### SetAutoApplyFlagNil

`func (o *BoardAutoTemplate) SetAutoApplyFlagNil(b bool)`

 SetAutoApplyFlagNil sets the value for AutoApplyFlag to be an explicit nil

### UnsetAutoApplyFlag
`func (o *BoardAutoTemplate) UnsetAutoApplyFlag()`

UnsetAutoApplyFlag ensures that no value is present for AutoApplyFlag, not even an explicit nil
### GetInfo

`func (o *BoardAutoTemplate) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardAutoTemplate) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardAutoTemplate) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardAutoTemplate) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


