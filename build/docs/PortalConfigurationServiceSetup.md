# PortalConfigurationServiceSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ServiceTypeFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceSubTypeFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceSubTypeItemFlag** | Pointer to **NullableBool** |  | [optional] 
**StatusFlag** | Pointer to **NullableBool** |  | [optional] 
**SiteNameFlag** | Pointer to **NullableBool** |  | [optional] 
**EnteredDateFlag** | Pointer to **NullableBool** |  | [optional] 
**LastUpdateFlag** | Pointer to **NullableBool** |  | [optional] 
**RequiredDateFlag** | Pointer to **NullableBool** |  | [optional] 
**ContactFlag** | Pointer to **NullableBool** |  | [optional] 
**AssignedResourcesFlag** | Pointer to **NullableBool** |  | [optional] 
**SlaInfoFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceBoardFlag** | Pointer to **NullableBool** |  | [optional] 
**BudgetHoursFlag** | Pointer to **NullableBool** |  | [optional] 
**ActualHoursFlag** | Pointer to **NullableBool** |  | [optional] 
**ApprovalStatusFlag** | Pointer to **NullableBool** |  | [optional] 
**OpenTasksFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedTasksFlag** | Pointer to **NullableBool** |  | [optional] 
**EnableChatAssistFlag** | Pointer to **NullableBool** |  | [optional] 
**DisplayClosedTicketsOption** | **NullableString** |  | 
**TimeMaterialsTicketTemplate** | [**ServiceSignoffReference**](ServiceSignoffReference.md) |  | 
**FixedFeeTicketTemplate** | [**ServiceSignoffReference**](ServiceSignoffReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPortalConfigurationServiceSetup

`func NewPortalConfigurationServiceSetup(displayClosedTicketsOption NullableString, timeMaterialsTicketTemplate ServiceSignoffReference, fixedFeeTicketTemplate ServiceSignoffReference, ) *PortalConfigurationServiceSetup`

NewPortalConfigurationServiceSetup instantiates a new PortalConfigurationServiceSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalConfigurationServiceSetupWithDefaults

`func NewPortalConfigurationServiceSetupWithDefaults() *PortalConfigurationServiceSetup`

NewPortalConfigurationServiceSetupWithDefaults instantiates a new PortalConfigurationServiceSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalConfigurationServiceSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalConfigurationServiceSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalConfigurationServiceSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PortalConfigurationServiceSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServiceTypeFlag

`func (o *PortalConfigurationServiceSetup) GetServiceTypeFlag() bool`

GetServiceTypeFlag returns the ServiceTypeFlag field if non-nil, zero value otherwise.

### GetServiceTypeFlagOk

`func (o *PortalConfigurationServiceSetup) GetServiceTypeFlagOk() (*bool, bool)`

GetServiceTypeFlagOk returns a tuple with the ServiceTypeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypeFlag

`func (o *PortalConfigurationServiceSetup) SetServiceTypeFlag(v bool)`

SetServiceTypeFlag sets ServiceTypeFlag field to given value.

### HasServiceTypeFlag

`func (o *PortalConfigurationServiceSetup) HasServiceTypeFlag() bool`

HasServiceTypeFlag returns a boolean if a field has been set.

### SetServiceTypeFlagNil

`func (o *PortalConfigurationServiceSetup) SetServiceTypeFlagNil(b bool)`

 SetServiceTypeFlagNil sets the value for ServiceTypeFlag to be an explicit nil

### UnsetServiceTypeFlag
`func (o *PortalConfigurationServiceSetup) UnsetServiceTypeFlag()`

UnsetServiceTypeFlag ensures that no value is present for ServiceTypeFlag, not even an explicit nil
### GetServiceSubTypeFlag

`func (o *PortalConfigurationServiceSetup) GetServiceSubTypeFlag() bool`

GetServiceSubTypeFlag returns the ServiceSubTypeFlag field if non-nil, zero value otherwise.

### GetServiceSubTypeFlagOk

`func (o *PortalConfigurationServiceSetup) GetServiceSubTypeFlagOk() (*bool, bool)`

GetServiceSubTypeFlagOk returns a tuple with the ServiceSubTypeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSubTypeFlag

`func (o *PortalConfigurationServiceSetup) SetServiceSubTypeFlag(v bool)`

SetServiceSubTypeFlag sets ServiceSubTypeFlag field to given value.

### HasServiceSubTypeFlag

`func (o *PortalConfigurationServiceSetup) HasServiceSubTypeFlag() bool`

HasServiceSubTypeFlag returns a boolean if a field has been set.

### SetServiceSubTypeFlagNil

`func (o *PortalConfigurationServiceSetup) SetServiceSubTypeFlagNil(b bool)`

 SetServiceSubTypeFlagNil sets the value for ServiceSubTypeFlag to be an explicit nil

### UnsetServiceSubTypeFlag
`func (o *PortalConfigurationServiceSetup) UnsetServiceSubTypeFlag()`

UnsetServiceSubTypeFlag ensures that no value is present for ServiceSubTypeFlag, not even an explicit nil
### GetServiceSubTypeItemFlag

`func (o *PortalConfigurationServiceSetup) GetServiceSubTypeItemFlag() bool`

GetServiceSubTypeItemFlag returns the ServiceSubTypeItemFlag field if non-nil, zero value otherwise.

### GetServiceSubTypeItemFlagOk

`func (o *PortalConfigurationServiceSetup) GetServiceSubTypeItemFlagOk() (*bool, bool)`

GetServiceSubTypeItemFlagOk returns a tuple with the ServiceSubTypeItemFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSubTypeItemFlag

`func (o *PortalConfigurationServiceSetup) SetServiceSubTypeItemFlag(v bool)`

SetServiceSubTypeItemFlag sets ServiceSubTypeItemFlag field to given value.

### HasServiceSubTypeItemFlag

`func (o *PortalConfigurationServiceSetup) HasServiceSubTypeItemFlag() bool`

HasServiceSubTypeItemFlag returns a boolean if a field has been set.

### SetServiceSubTypeItemFlagNil

`func (o *PortalConfigurationServiceSetup) SetServiceSubTypeItemFlagNil(b bool)`

 SetServiceSubTypeItemFlagNil sets the value for ServiceSubTypeItemFlag to be an explicit nil

### UnsetServiceSubTypeItemFlag
`func (o *PortalConfigurationServiceSetup) UnsetServiceSubTypeItemFlag()`

UnsetServiceSubTypeItemFlag ensures that no value is present for ServiceSubTypeItemFlag, not even an explicit nil
### GetStatusFlag

`func (o *PortalConfigurationServiceSetup) GetStatusFlag() bool`

GetStatusFlag returns the StatusFlag field if non-nil, zero value otherwise.

### GetStatusFlagOk

`func (o *PortalConfigurationServiceSetup) GetStatusFlagOk() (*bool, bool)`

GetStatusFlagOk returns a tuple with the StatusFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusFlag

`func (o *PortalConfigurationServiceSetup) SetStatusFlag(v bool)`

SetStatusFlag sets StatusFlag field to given value.

### HasStatusFlag

`func (o *PortalConfigurationServiceSetup) HasStatusFlag() bool`

HasStatusFlag returns a boolean if a field has been set.

### SetStatusFlagNil

`func (o *PortalConfigurationServiceSetup) SetStatusFlagNil(b bool)`

 SetStatusFlagNil sets the value for StatusFlag to be an explicit nil

### UnsetStatusFlag
`func (o *PortalConfigurationServiceSetup) UnsetStatusFlag()`

UnsetStatusFlag ensures that no value is present for StatusFlag, not even an explicit nil
### GetSiteNameFlag

`func (o *PortalConfigurationServiceSetup) GetSiteNameFlag() bool`

GetSiteNameFlag returns the SiteNameFlag field if non-nil, zero value otherwise.

### GetSiteNameFlagOk

`func (o *PortalConfigurationServiceSetup) GetSiteNameFlagOk() (*bool, bool)`

GetSiteNameFlagOk returns a tuple with the SiteNameFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNameFlag

`func (o *PortalConfigurationServiceSetup) SetSiteNameFlag(v bool)`

SetSiteNameFlag sets SiteNameFlag field to given value.

### HasSiteNameFlag

`func (o *PortalConfigurationServiceSetup) HasSiteNameFlag() bool`

HasSiteNameFlag returns a boolean if a field has been set.

### SetSiteNameFlagNil

`func (o *PortalConfigurationServiceSetup) SetSiteNameFlagNil(b bool)`

 SetSiteNameFlagNil sets the value for SiteNameFlag to be an explicit nil

### UnsetSiteNameFlag
`func (o *PortalConfigurationServiceSetup) UnsetSiteNameFlag()`

UnsetSiteNameFlag ensures that no value is present for SiteNameFlag, not even an explicit nil
### GetEnteredDateFlag

`func (o *PortalConfigurationServiceSetup) GetEnteredDateFlag() bool`

GetEnteredDateFlag returns the EnteredDateFlag field if non-nil, zero value otherwise.

### GetEnteredDateFlagOk

`func (o *PortalConfigurationServiceSetup) GetEnteredDateFlagOk() (*bool, bool)`

GetEnteredDateFlagOk returns a tuple with the EnteredDateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredDateFlag

`func (o *PortalConfigurationServiceSetup) SetEnteredDateFlag(v bool)`

SetEnteredDateFlag sets EnteredDateFlag field to given value.

### HasEnteredDateFlag

`func (o *PortalConfigurationServiceSetup) HasEnteredDateFlag() bool`

HasEnteredDateFlag returns a boolean if a field has been set.

### SetEnteredDateFlagNil

`func (o *PortalConfigurationServiceSetup) SetEnteredDateFlagNil(b bool)`

 SetEnteredDateFlagNil sets the value for EnteredDateFlag to be an explicit nil

### UnsetEnteredDateFlag
`func (o *PortalConfigurationServiceSetup) UnsetEnteredDateFlag()`

UnsetEnteredDateFlag ensures that no value is present for EnteredDateFlag, not even an explicit nil
### GetLastUpdateFlag

`func (o *PortalConfigurationServiceSetup) GetLastUpdateFlag() bool`

GetLastUpdateFlag returns the LastUpdateFlag field if non-nil, zero value otherwise.

### GetLastUpdateFlagOk

`func (o *PortalConfigurationServiceSetup) GetLastUpdateFlagOk() (*bool, bool)`

GetLastUpdateFlagOk returns a tuple with the LastUpdateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateFlag

`func (o *PortalConfigurationServiceSetup) SetLastUpdateFlag(v bool)`

SetLastUpdateFlag sets LastUpdateFlag field to given value.

### HasLastUpdateFlag

`func (o *PortalConfigurationServiceSetup) HasLastUpdateFlag() bool`

HasLastUpdateFlag returns a boolean if a field has been set.

### SetLastUpdateFlagNil

`func (o *PortalConfigurationServiceSetup) SetLastUpdateFlagNil(b bool)`

 SetLastUpdateFlagNil sets the value for LastUpdateFlag to be an explicit nil

### UnsetLastUpdateFlag
`func (o *PortalConfigurationServiceSetup) UnsetLastUpdateFlag()`

UnsetLastUpdateFlag ensures that no value is present for LastUpdateFlag, not even an explicit nil
### GetRequiredDateFlag

`func (o *PortalConfigurationServiceSetup) GetRequiredDateFlag() bool`

GetRequiredDateFlag returns the RequiredDateFlag field if non-nil, zero value otherwise.

### GetRequiredDateFlagOk

`func (o *PortalConfigurationServiceSetup) GetRequiredDateFlagOk() (*bool, bool)`

GetRequiredDateFlagOk returns a tuple with the RequiredDateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDateFlag

`func (o *PortalConfigurationServiceSetup) SetRequiredDateFlag(v bool)`

SetRequiredDateFlag sets RequiredDateFlag field to given value.

### HasRequiredDateFlag

`func (o *PortalConfigurationServiceSetup) HasRequiredDateFlag() bool`

HasRequiredDateFlag returns a boolean if a field has been set.

### SetRequiredDateFlagNil

`func (o *PortalConfigurationServiceSetup) SetRequiredDateFlagNil(b bool)`

 SetRequiredDateFlagNil sets the value for RequiredDateFlag to be an explicit nil

### UnsetRequiredDateFlag
`func (o *PortalConfigurationServiceSetup) UnsetRequiredDateFlag()`

UnsetRequiredDateFlag ensures that no value is present for RequiredDateFlag, not even an explicit nil
### GetContactFlag

`func (o *PortalConfigurationServiceSetup) GetContactFlag() bool`

GetContactFlag returns the ContactFlag field if non-nil, zero value otherwise.

### GetContactFlagOk

`func (o *PortalConfigurationServiceSetup) GetContactFlagOk() (*bool, bool)`

GetContactFlagOk returns a tuple with the ContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactFlag

`func (o *PortalConfigurationServiceSetup) SetContactFlag(v bool)`

SetContactFlag sets ContactFlag field to given value.

### HasContactFlag

`func (o *PortalConfigurationServiceSetup) HasContactFlag() bool`

HasContactFlag returns a boolean if a field has been set.

### SetContactFlagNil

`func (o *PortalConfigurationServiceSetup) SetContactFlagNil(b bool)`

 SetContactFlagNil sets the value for ContactFlag to be an explicit nil

### UnsetContactFlag
`func (o *PortalConfigurationServiceSetup) UnsetContactFlag()`

UnsetContactFlag ensures that no value is present for ContactFlag, not even an explicit nil
### GetAssignedResourcesFlag

`func (o *PortalConfigurationServiceSetup) GetAssignedResourcesFlag() bool`

GetAssignedResourcesFlag returns the AssignedResourcesFlag field if non-nil, zero value otherwise.

### GetAssignedResourcesFlagOk

`func (o *PortalConfigurationServiceSetup) GetAssignedResourcesFlagOk() (*bool, bool)`

GetAssignedResourcesFlagOk returns a tuple with the AssignedResourcesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedResourcesFlag

`func (o *PortalConfigurationServiceSetup) SetAssignedResourcesFlag(v bool)`

SetAssignedResourcesFlag sets AssignedResourcesFlag field to given value.

### HasAssignedResourcesFlag

`func (o *PortalConfigurationServiceSetup) HasAssignedResourcesFlag() bool`

HasAssignedResourcesFlag returns a boolean if a field has been set.

### SetAssignedResourcesFlagNil

`func (o *PortalConfigurationServiceSetup) SetAssignedResourcesFlagNil(b bool)`

 SetAssignedResourcesFlagNil sets the value for AssignedResourcesFlag to be an explicit nil

### UnsetAssignedResourcesFlag
`func (o *PortalConfigurationServiceSetup) UnsetAssignedResourcesFlag()`

UnsetAssignedResourcesFlag ensures that no value is present for AssignedResourcesFlag, not even an explicit nil
### GetSlaInfoFlag

`func (o *PortalConfigurationServiceSetup) GetSlaInfoFlag() bool`

GetSlaInfoFlag returns the SlaInfoFlag field if non-nil, zero value otherwise.

### GetSlaInfoFlagOk

`func (o *PortalConfigurationServiceSetup) GetSlaInfoFlagOk() (*bool, bool)`

GetSlaInfoFlagOk returns a tuple with the SlaInfoFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaInfoFlag

`func (o *PortalConfigurationServiceSetup) SetSlaInfoFlag(v bool)`

SetSlaInfoFlag sets SlaInfoFlag field to given value.

### HasSlaInfoFlag

`func (o *PortalConfigurationServiceSetup) HasSlaInfoFlag() bool`

HasSlaInfoFlag returns a boolean if a field has been set.

### SetSlaInfoFlagNil

`func (o *PortalConfigurationServiceSetup) SetSlaInfoFlagNil(b bool)`

 SetSlaInfoFlagNil sets the value for SlaInfoFlag to be an explicit nil

### UnsetSlaInfoFlag
`func (o *PortalConfigurationServiceSetup) UnsetSlaInfoFlag()`

UnsetSlaInfoFlag ensures that no value is present for SlaInfoFlag, not even an explicit nil
### GetServiceBoardFlag

`func (o *PortalConfigurationServiceSetup) GetServiceBoardFlag() bool`

GetServiceBoardFlag returns the ServiceBoardFlag field if non-nil, zero value otherwise.

### GetServiceBoardFlagOk

`func (o *PortalConfigurationServiceSetup) GetServiceBoardFlagOk() (*bool, bool)`

GetServiceBoardFlagOk returns a tuple with the ServiceBoardFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBoardFlag

`func (o *PortalConfigurationServiceSetup) SetServiceBoardFlag(v bool)`

SetServiceBoardFlag sets ServiceBoardFlag field to given value.

### HasServiceBoardFlag

`func (o *PortalConfigurationServiceSetup) HasServiceBoardFlag() bool`

HasServiceBoardFlag returns a boolean if a field has been set.

### SetServiceBoardFlagNil

`func (o *PortalConfigurationServiceSetup) SetServiceBoardFlagNil(b bool)`

 SetServiceBoardFlagNil sets the value for ServiceBoardFlag to be an explicit nil

### UnsetServiceBoardFlag
`func (o *PortalConfigurationServiceSetup) UnsetServiceBoardFlag()`

UnsetServiceBoardFlag ensures that no value is present for ServiceBoardFlag, not even an explicit nil
### GetBudgetHoursFlag

`func (o *PortalConfigurationServiceSetup) GetBudgetHoursFlag() bool`

GetBudgetHoursFlag returns the BudgetHoursFlag field if non-nil, zero value otherwise.

### GetBudgetHoursFlagOk

`func (o *PortalConfigurationServiceSetup) GetBudgetHoursFlagOk() (*bool, bool)`

GetBudgetHoursFlagOk returns a tuple with the BudgetHoursFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHoursFlag

`func (o *PortalConfigurationServiceSetup) SetBudgetHoursFlag(v bool)`

SetBudgetHoursFlag sets BudgetHoursFlag field to given value.

### HasBudgetHoursFlag

`func (o *PortalConfigurationServiceSetup) HasBudgetHoursFlag() bool`

HasBudgetHoursFlag returns a boolean if a field has been set.

### SetBudgetHoursFlagNil

`func (o *PortalConfigurationServiceSetup) SetBudgetHoursFlagNil(b bool)`

 SetBudgetHoursFlagNil sets the value for BudgetHoursFlag to be an explicit nil

### UnsetBudgetHoursFlag
`func (o *PortalConfigurationServiceSetup) UnsetBudgetHoursFlag()`

UnsetBudgetHoursFlag ensures that no value is present for BudgetHoursFlag, not even an explicit nil
### GetActualHoursFlag

`func (o *PortalConfigurationServiceSetup) GetActualHoursFlag() bool`

GetActualHoursFlag returns the ActualHoursFlag field if non-nil, zero value otherwise.

### GetActualHoursFlagOk

`func (o *PortalConfigurationServiceSetup) GetActualHoursFlagOk() (*bool, bool)`

GetActualHoursFlagOk returns a tuple with the ActualHoursFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualHoursFlag

`func (o *PortalConfigurationServiceSetup) SetActualHoursFlag(v bool)`

SetActualHoursFlag sets ActualHoursFlag field to given value.

### HasActualHoursFlag

`func (o *PortalConfigurationServiceSetup) HasActualHoursFlag() bool`

HasActualHoursFlag returns a boolean if a field has been set.

### SetActualHoursFlagNil

`func (o *PortalConfigurationServiceSetup) SetActualHoursFlagNil(b bool)`

 SetActualHoursFlagNil sets the value for ActualHoursFlag to be an explicit nil

### UnsetActualHoursFlag
`func (o *PortalConfigurationServiceSetup) UnsetActualHoursFlag()`

UnsetActualHoursFlag ensures that no value is present for ActualHoursFlag, not even an explicit nil
### GetApprovalStatusFlag

`func (o *PortalConfigurationServiceSetup) GetApprovalStatusFlag() bool`

GetApprovalStatusFlag returns the ApprovalStatusFlag field if non-nil, zero value otherwise.

### GetApprovalStatusFlagOk

`func (o *PortalConfigurationServiceSetup) GetApprovalStatusFlagOk() (*bool, bool)`

GetApprovalStatusFlagOk returns a tuple with the ApprovalStatusFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatusFlag

`func (o *PortalConfigurationServiceSetup) SetApprovalStatusFlag(v bool)`

SetApprovalStatusFlag sets ApprovalStatusFlag field to given value.

### HasApprovalStatusFlag

`func (o *PortalConfigurationServiceSetup) HasApprovalStatusFlag() bool`

HasApprovalStatusFlag returns a boolean if a field has been set.

### SetApprovalStatusFlagNil

`func (o *PortalConfigurationServiceSetup) SetApprovalStatusFlagNil(b bool)`

 SetApprovalStatusFlagNil sets the value for ApprovalStatusFlag to be an explicit nil

### UnsetApprovalStatusFlag
`func (o *PortalConfigurationServiceSetup) UnsetApprovalStatusFlag()`

UnsetApprovalStatusFlag ensures that no value is present for ApprovalStatusFlag, not even an explicit nil
### GetOpenTasksFlag

`func (o *PortalConfigurationServiceSetup) GetOpenTasksFlag() bool`

GetOpenTasksFlag returns the OpenTasksFlag field if non-nil, zero value otherwise.

### GetOpenTasksFlagOk

`func (o *PortalConfigurationServiceSetup) GetOpenTasksFlagOk() (*bool, bool)`

GetOpenTasksFlagOk returns a tuple with the OpenTasksFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTasksFlag

`func (o *PortalConfigurationServiceSetup) SetOpenTasksFlag(v bool)`

SetOpenTasksFlag sets OpenTasksFlag field to given value.

### HasOpenTasksFlag

`func (o *PortalConfigurationServiceSetup) HasOpenTasksFlag() bool`

HasOpenTasksFlag returns a boolean if a field has been set.

### SetOpenTasksFlagNil

`func (o *PortalConfigurationServiceSetup) SetOpenTasksFlagNil(b bool)`

 SetOpenTasksFlagNil sets the value for OpenTasksFlag to be an explicit nil

### UnsetOpenTasksFlag
`func (o *PortalConfigurationServiceSetup) UnsetOpenTasksFlag()`

UnsetOpenTasksFlag ensures that no value is present for OpenTasksFlag, not even an explicit nil
### GetClosedTasksFlag

`func (o *PortalConfigurationServiceSetup) GetClosedTasksFlag() bool`

GetClosedTasksFlag returns the ClosedTasksFlag field if non-nil, zero value otherwise.

### GetClosedTasksFlagOk

`func (o *PortalConfigurationServiceSetup) GetClosedTasksFlagOk() (*bool, bool)`

GetClosedTasksFlagOk returns a tuple with the ClosedTasksFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedTasksFlag

`func (o *PortalConfigurationServiceSetup) SetClosedTasksFlag(v bool)`

SetClosedTasksFlag sets ClosedTasksFlag field to given value.

### HasClosedTasksFlag

`func (o *PortalConfigurationServiceSetup) HasClosedTasksFlag() bool`

HasClosedTasksFlag returns a boolean if a field has been set.

### SetClosedTasksFlagNil

`func (o *PortalConfigurationServiceSetup) SetClosedTasksFlagNil(b bool)`

 SetClosedTasksFlagNil sets the value for ClosedTasksFlag to be an explicit nil

### UnsetClosedTasksFlag
`func (o *PortalConfigurationServiceSetup) UnsetClosedTasksFlag()`

UnsetClosedTasksFlag ensures that no value is present for ClosedTasksFlag, not even an explicit nil
### GetEnableChatAssistFlag

`func (o *PortalConfigurationServiceSetup) GetEnableChatAssistFlag() bool`

GetEnableChatAssistFlag returns the EnableChatAssistFlag field if non-nil, zero value otherwise.

### GetEnableChatAssistFlagOk

`func (o *PortalConfigurationServiceSetup) GetEnableChatAssistFlagOk() (*bool, bool)`

GetEnableChatAssistFlagOk returns a tuple with the EnableChatAssistFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChatAssistFlag

`func (o *PortalConfigurationServiceSetup) SetEnableChatAssistFlag(v bool)`

SetEnableChatAssistFlag sets EnableChatAssistFlag field to given value.

### HasEnableChatAssistFlag

`func (o *PortalConfigurationServiceSetup) HasEnableChatAssistFlag() bool`

HasEnableChatAssistFlag returns a boolean if a field has been set.

### SetEnableChatAssistFlagNil

`func (o *PortalConfigurationServiceSetup) SetEnableChatAssistFlagNil(b bool)`

 SetEnableChatAssistFlagNil sets the value for EnableChatAssistFlag to be an explicit nil

### UnsetEnableChatAssistFlag
`func (o *PortalConfigurationServiceSetup) UnsetEnableChatAssistFlag()`

UnsetEnableChatAssistFlag ensures that no value is present for EnableChatAssistFlag, not even an explicit nil
### GetDisplayClosedTicketsOption

`func (o *PortalConfigurationServiceSetup) GetDisplayClosedTicketsOption() string`

GetDisplayClosedTicketsOption returns the DisplayClosedTicketsOption field if non-nil, zero value otherwise.

### GetDisplayClosedTicketsOptionOk

`func (o *PortalConfigurationServiceSetup) GetDisplayClosedTicketsOptionOk() (*string, bool)`

GetDisplayClosedTicketsOptionOk returns a tuple with the DisplayClosedTicketsOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayClosedTicketsOption

`func (o *PortalConfigurationServiceSetup) SetDisplayClosedTicketsOption(v string)`

SetDisplayClosedTicketsOption sets DisplayClosedTicketsOption field to given value.


### SetDisplayClosedTicketsOptionNil

`func (o *PortalConfigurationServiceSetup) SetDisplayClosedTicketsOptionNil(b bool)`

 SetDisplayClosedTicketsOptionNil sets the value for DisplayClosedTicketsOption to be an explicit nil

### UnsetDisplayClosedTicketsOption
`func (o *PortalConfigurationServiceSetup) UnsetDisplayClosedTicketsOption()`

UnsetDisplayClosedTicketsOption ensures that no value is present for DisplayClosedTicketsOption, not even an explicit nil
### GetTimeMaterialsTicketTemplate

`func (o *PortalConfigurationServiceSetup) GetTimeMaterialsTicketTemplate() ServiceSignoffReference`

GetTimeMaterialsTicketTemplate returns the TimeMaterialsTicketTemplate field if non-nil, zero value otherwise.

### GetTimeMaterialsTicketTemplateOk

`func (o *PortalConfigurationServiceSetup) GetTimeMaterialsTicketTemplateOk() (*ServiceSignoffReference, bool)`

GetTimeMaterialsTicketTemplateOk returns a tuple with the TimeMaterialsTicketTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMaterialsTicketTemplate

`func (o *PortalConfigurationServiceSetup) SetTimeMaterialsTicketTemplate(v ServiceSignoffReference)`

SetTimeMaterialsTicketTemplate sets TimeMaterialsTicketTemplate field to given value.


### GetFixedFeeTicketTemplate

`func (o *PortalConfigurationServiceSetup) GetFixedFeeTicketTemplate() ServiceSignoffReference`

GetFixedFeeTicketTemplate returns the FixedFeeTicketTemplate field if non-nil, zero value otherwise.

### GetFixedFeeTicketTemplateOk

`func (o *PortalConfigurationServiceSetup) GetFixedFeeTicketTemplateOk() (*ServiceSignoffReference, bool)`

GetFixedFeeTicketTemplateOk returns a tuple with the FixedFeeTicketTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedFeeTicketTemplate

`func (o *PortalConfigurationServiceSetup) SetFixedFeeTicketTemplate(v ServiceSignoffReference)`

SetFixedFeeTicketTemplate sets FixedFeeTicketTemplate field to given value.


### GetInfo

`func (o *PortalConfigurationServiceSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PortalConfigurationServiceSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PortalConfigurationServiceSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PortalConfigurationServiceSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


