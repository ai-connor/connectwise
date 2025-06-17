# Board

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Location** | [**SystemLocationReference**](SystemLocationReference.md) |  | 
**Department** | [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**SignOffTemplate** | Pointer to [**ServiceSignoffReference**](ServiceSignoffReference.md) |  | [optional] 
**SendToContactFlag** | Pointer to **NullableBool** |  | [optional] 
**ContactTemplate** | Pointer to [**ServiceEmailTemplateReference**](ServiceEmailTemplateReference.md) |  | [optional] 
**SendToResourceFlag** | Pointer to **NullableBool** |  | [optional] 
**ResourceTemplate** | Pointer to [**ServiceEmailTemplateReference**](ServiceEmailTemplateReference.md) |  | [optional] 
**ProjectFlag** | Pointer to **NullableBool** |  | [optional] 
**ShowDependenciesFlag** | Pointer to **NullableBool** | This field only shows if it is Project Board. | [optional] 
**ShowEstimatesFlag** | Pointer to **NullableBool** | This field only shows if it is Project Board. | [optional] 
**BoardIcon** | Pointer to [**DocumentReference**](DocumentReference.md) |  | [optional] 
**BillTicketsAfterClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**BillTicketSeparatelyFlag** | Pointer to **NullableBool** |  | [optional] 
**BillUnapprovedTimeExpenseFlag** | Pointer to **NullableBool** |  | [optional] 
**OverrideBillingSetupFlag** | Pointer to **NullableBool** |  | [optional] 
**DispatchMember** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**ServiceManagerMember** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**DutyManagerMember** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**OncallMember** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**BillTime** | Pointer to **NullableString** |  | [optional] 
**BillExpense** | Pointer to **NullableString** |  | [optional] 
**BillProduct** | Pointer to **NullableString** |  | [optional] 
**AutoCloseStatus** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**AutoAssignNewTicketsFlag** | Pointer to **NullableBool** |  | [optional] 
**AutoAssignNewECTicketsFlag** | Pointer to **NullableBool** |  | [optional] 
**AutoAssignNewPortalTicketsFlag** | Pointer to **NullableBool** |  | [optional] 
**DiscussionsLockedFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeEntryLockedFlag** | Pointer to **NullableBool** |  | [optional] 
**NotifyEmailFrom** | Pointer to **string** |  Max length: 50; | [optional] 
**NotifyEmailFromName** | Pointer to **string** |  Max length: 60; | [optional] 
**ClosedLoopDiscussionsFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedLoopResolutionFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedLoopInternalAnalysisFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeEntryDiscussionFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeEntryResolutionFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeEntryInternalAnalysisFlag** | Pointer to **NullableBool** |  | [optional] 
**ProblemSort** | Pointer to **NullableString** |  | [optional] 
**ResolutionSort** | Pointer to **NullableString** |  | [optional] 
**InternalAnalysisSort** | Pointer to **NullableString** |  | [optional] 
**EmailConnectorAllowReopenClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailConnectorReopenStatus** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**EmailConnectorReopenResourcesFlag** | Pointer to **NullableBool** | This field can only be set when emailConnectorAllowReopenClosed is true. | [optional] 
**EmailConnectorNewTicketNoMatchFlag** | Pointer to **NullableBool** | This field can only be set when emailConnectorAllowReopenClosed is true. | [optional] 
**EmailConnectorNeverReopenByDaysFlag** | Pointer to **NullableBool** | This field can only be set when emailConnectorAllowReopenClosed is true. | [optional] 
**EmailConnectorReopenDaysLimit** | Pointer to **NullableInt32** | This field can only be set when emailConnectorNeverReopenByDaysFlag and emailConnectorAllowReopenClosed are both true             This field is required when emailConnectorNeverReopenByDaysFlag is true. | [optional] 
**EmailConnectorNeverReopenByDaysClosedFlag** | Pointer to **NullableBool** | This field can only be set when emailConnectorAllowReopenClosed is true. | [optional] 
**EmailConnectorReopenDaysClosedLimit** | Pointer to **NullableInt32** | This field can only be set when emailConnectorNeverReopenByDaysClosedFlag and emailConnectorAllowReopenClosed are both true             This field is required when emailConnectorNeverReopenByDaysClosedFlag is true. | [optional] 
**UseMemberDisplayNameFlag** | Pointer to **NullableBool** |  | [optional] 
**SendToCCFlag** | Pointer to **NullableBool** |  | [optional] 
**AutoAssignTicketOwnerFlag** | Pointer to **NullableBool** |  | [optional] 
**AutoAssignLimitFlag** | Pointer to **NullableBool** |  | [optional] 
**AutoAssignLimitAmount** | Pointer to **NullableInt32** | This field can only be set when autoAssignLimitFlag is true | [optional] 
**ClosedLoopAllFlag** | Pointer to **NullableBool** |  | [optional] 
**PercentageCalculation** | Pointer to **NullableString** |  | [optional] 
**AllSort** | Pointer to **NullableString** |  | [optional] 
**MarkFirstNoteIssueFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictBoardByDefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**SendToBundledFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoard

`func NewBoard(name string, location SystemLocationReference, department SystemDepartmentReference, ) *Board`

NewBoard instantiates a new Board object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardWithDefaults

`func NewBoardWithDefaults() *Board`

NewBoardWithDefaults instantiates a new Board object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Board) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Board) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Board) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Board) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Board) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Board) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Board) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *Board) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Board) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Board) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.


### GetDepartment

`func (o *Board) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Board) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Board) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.


### GetInactiveFlag

`func (o *Board) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *Board) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *Board) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *Board) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *Board) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *Board) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetSignOffTemplate

`func (o *Board) GetSignOffTemplate() ServiceSignoffReference`

GetSignOffTemplate returns the SignOffTemplate field if non-nil, zero value otherwise.

### GetSignOffTemplateOk

`func (o *Board) GetSignOffTemplateOk() (*ServiceSignoffReference, bool)`

GetSignOffTemplateOk returns a tuple with the SignOffTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOffTemplate

`func (o *Board) SetSignOffTemplate(v ServiceSignoffReference)`

SetSignOffTemplate sets SignOffTemplate field to given value.

### HasSignOffTemplate

`func (o *Board) HasSignOffTemplate() bool`

HasSignOffTemplate returns a boolean if a field has been set.

### GetSendToContactFlag

`func (o *Board) GetSendToContactFlag() bool`

GetSendToContactFlag returns the SendToContactFlag field if non-nil, zero value otherwise.

### GetSendToContactFlagOk

`func (o *Board) GetSendToContactFlagOk() (*bool, bool)`

GetSendToContactFlagOk returns a tuple with the SendToContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToContactFlag

`func (o *Board) SetSendToContactFlag(v bool)`

SetSendToContactFlag sets SendToContactFlag field to given value.

### HasSendToContactFlag

`func (o *Board) HasSendToContactFlag() bool`

HasSendToContactFlag returns a boolean if a field has been set.

### SetSendToContactFlagNil

`func (o *Board) SetSendToContactFlagNil(b bool)`

 SetSendToContactFlagNil sets the value for SendToContactFlag to be an explicit nil

### UnsetSendToContactFlag
`func (o *Board) UnsetSendToContactFlag()`

UnsetSendToContactFlag ensures that no value is present for SendToContactFlag, not even an explicit nil
### GetContactTemplate

`func (o *Board) GetContactTemplate() ServiceEmailTemplateReference`

GetContactTemplate returns the ContactTemplate field if non-nil, zero value otherwise.

### GetContactTemplateOk

`func (o *Board) GetContactTemplateOk() (*ServiceEmailTemplateReference, bool)`

GetContactTemplateOk returns a tuple with the ContactTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactTemplate

`func (o *Board) SetContactTemplate(v ServiceEmailTemplateReference)`

SetContactTemplate sets ContactTemplate field to given value.

### HasContactTemplate

`func (o *Board) HasContactTemplate() bool`

HasContactTemplate returns a boolean if a field has been set.

### GetSendToResourceFlag

`func (o *Board) GetSendToResourceFlag() bool`

GetSendToResourceFlag returns the SendToResourceFlag field if non-nil, zero value otherwise.

### GetSendToResourceFlagOk

`func (o *Board) GetSendToResourceFlagOk() (*bool, bool)`

GetSendToResourceFlagOk returns a tuple with the SendToResourceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToResourceFlag

`func (o *Board) SetSendToResourceFlag(v bool)`

SetSendToResourceFlag sets SendToResourceFlag field to given value.

### HasSendToResourceFlag

`func (o *Board) HasSendToResourceFlag() bool`

HasSendToResourceFlag returns a boolean if a field has been set.

### SetSendToResourceFlagNil

`func (o *Board) SetSendToResourceFlagNil(b bool)`

 SetSendToResourceFlagNil sets the value for SendToResourceFlag to be an explicit nil

### UnsetSendToResourceFlag
`func (o *Board) UnsetSendToResourceFlag()`

UnsetSendToResourceFlag ensures that no value is present for SendToResourceFlag, not even an explicit nil
### GetResourceTemplate

`func (o *Board) GetResourceTemplate() ServiceEmailTemplateReference`

GetResourceTemplate returns the ResourceTemplate field if non-nil, zero value otherwise.

### GetResourceTemplateOk

`func (o *Board) GetResourceTemplateOk() (*ServiceEmailTemplateReference, bool)`

GetResourceTemplateOk returns a tuple with the ResourceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTemplate

`func (o *Board) SetResourceTemplate(v ServiceEmailTemplateReference)`

SetResourceTemplate sets ResourceTemplate field to given value.

### HasResourceTemplate

`func (o *Board) HasResourceTemplate() bool`

HasResourceTemplate returns a boolean if a field has been set.

### GetProjectFlag

`func (o *Board) GetProjectFlag() bool`

GetProjectFlag returns the ProjectFlag field if non-nil, zero value otherwise.

### GetProjectFlagOk

`func (o *Board) GetProjectFlagOk() (*bool, bool)`

GetProjectFlagOk returns a tuple with the ProjectFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectFlag

`func (o *Board) SetProjectFlag(v bool)`

SetProjectFlag sets ProjectFlag field to given value.

### HasProjectFlag

`func (o *Board) HasProjectFlag() bool`

HasProjectFlag returns a boolean if a field has been set.

### SetProjectFlagNil

`func (o *Board) SetProjectFlagNil(b bool)`

 SetProjectFlagNil sets the value for ProjectFlag to be an explicit nil

### UnsetProjectFlag
`func (o *Board) UnsetProjectFlag()`

UnsetProjectFlag ensures that no value is present for ProjectFlag, not even an explicit nil
### GetShowDependenciesFlag

`func (o *Board) GetShowDependenciesFlag() bool`

GetShowDependenciesFlag returns the ShowDependenciesFlag field if non-nil, zero value otherwise.

### GetShowDependenciesFlagOk

`func (o *Board) GetShowDependenciesFlagOk() (*bool, bool)`

GetShowDependenciesFlagOk returns a tuple with the ShowDependenciesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDependenciesFlag

`func (o *Board) SetShowDependenciesFlag(v bool)`

SetShowDependenciesFlag sets ShowDependenciesFlag field to given value.

### HasShowDependenciesFlag

`func (o *Board) HasShowDependenciesFlag() bool`

HasShowDependenciesFlag returns a boolean if a field has been set.

### SetShowDependenciesFlagNil

`func (o *Board) SetShowDependenciesFlagNil(b bool)`

 SetShowDependenciesFlagNil sets the value for ShowDependenciesFlag to be an explicit nil

### UnsetShowDependenciesFlag
`func (o *Board) UnsetShowDependenciesFlag()`

UnsetShowDependenciesFlag ensures that no value is present for ShowDependenciesFlag, not even an explicit nil
### GetShowEstimatesFlag

`func (o *Board) GetShowEstimatesFlag() bool`

GetShowEstimatesFlag returns the ShowEstimatesFlag field if non-nil, zero value otherwise.

### GetShowEstimatesFlagOk

`func (o *Board) GetShowEstimatesFlagOk() (*bool, bool)`

GetShowEstimatesFlagOk returns a tuple with the ShowEstimatesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowEstimatesFlag

`func (o *Board) SetShowEstimatesFlag(v bool)`

SetShowEstimatesFlag sets ShowEstimatesFlag field to given value.

### HasShowEstimatesFlag

`func (o *Board) HasShowEstimatesFlag() bool`

HasShowEstimatesFlag returns a boolean if a field has been set.

### SetShowEstimatesFlagNil

`func (o *Board) SetShowEstimatesFlagNil(b bool)`

 SetShowEstimatesFlagNil sets the value for ShowEstimatesFlag to be an explicit nil

### UnsetShowEstimatesFlag
`func (o *Board) UnsetShowEstimatesFlag()`

UnsetShowEstimatesFlag ensures that no value is present for ShowEstimatesFlag, not even an explicit nil
### GetBoardIcon

`func (o *Board) GetBoardIcon() DocumentReference`

GetBoardIcon returns the BoardIcon field if non-nil, zero value otherwise.

### GetBoardIconOk

`func (o *Board) GetBoardIconOk() (*DocumentReference, bool)`

GetBoardIconOk returns a tuple with the BoardIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardIcon

`func (o *Board) SetBoardIcon(v DocumentReference)`

SetBoardIcon sets BoardIcon field to given value.

### HasBoardIcon

`func (o *Board) HasBoardIcon() bool`

HasBoardIcon returns a boolean if a field has been set.

### GetBillTicketsAfterClosedFlag

`func (o *Board) GetBillTicketsAfterClosedFlag() bool`

GetBillTicketsAfterClosedFlag returns the BillTicketsAfterClosedFlag field if non-nil, zero value otherwise.

### GetBillTicketsAfterClosedFlagOk

`func (o *Board) GetBillTicketsAfterClosedFlagOk() (*bool, bool)`

GetBillTicketsAfterClosedFlagOk returns a tuple with the BillTicketsAfterClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTicketsAfterClosedFlag

`func (o *Board) SetBillTicketsAfterClosedFlag(v bool)`

SetBillTicketsAfterClosedFlag sets BillTicketsAfterClosedFlag field to given value.

### HasBillTicketsAfterClosedFlag

`func (o *Board) HasBillTicketsAfterClosedFlag() bool`

HasBillTicketsAfterClosedFlag returns a boolean if a field has been set.

### SetBillTicketsAfterClosedFlagNil

`func (o *Board) SetBillTicketsAfterClosedFlagNil(b bool)`

 SetBillTicketsAfterClosedFlagNil sets the value for BillTicketsAfterClosedFlag to be an explicit nil

### UnsetBillTicketsAfterClosedFlag
`func (o *Board) UnsetBillTicketsAfterClosedFlag()`

UnsetBillTicketsAfterClosedFlag ensures that no value is present for BillTicketsAfterClosedFlag, not even an explicit nil
### GetBillTicketSeparatelyFlag

`func (o *Board) GetBillTicketSeparatelyFlag() bool`

GetBillTicketSeparatelyFlag returns the BillTicketSeparatelyFlag field if non-nil, zero value otherwise.

### GetBillTicketSeparatelyFlagOk

`func (o *Board) GetBillTicketSeparatelyFlagOk() (*bool, bool)`

GetBillTicketSeparatelyFlagOk returns a tuple with the BillTicketSeparatelyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTicketSeparatelyFlag

`func (o *Board) SetBillTicketSeparatelyFlag(v bool)`

SetBillTicketSeparatelyFlag sets BillTicketSeparatelyFlag field to given value.

### HasBillTicketSeparatelyFlag

`func (o *Board) HasBillTicketSeparatelyFlag() bool`

HasBillTicketSeparatelyFlag returns a boolean if a field has been set.

### SetBillTicketSeparatelyFlagNil

`func (o *Board) SetBillTicketSeparatelyFlagNil(b bool)`

 SetBillTicketSeparatelyFlagNil sets the value for BillTicketSeparatelyFlag to be an explicit nil

### UnsetBillTicketSeparatelyFlag
`func (o *Board) UnsetBillTicketSeparatelyFlag()`

UnsetBillTicketSeparatelyFlag ensures that no value is present for BillTicketSeparatelyFlag, not even an explicit nil
### GetBillUnapprovedTimeExpenseFlag

`func (o *Board) GetBillUnapprovedTimeExpenseFlag() bool`

GetBillUnapprovedTimeExpenseFlag returns the BillUnapprovedTimeExpenseFlag field if non-nil, zero value otherwise.

### GetBillUnapprovedTimeExpenseFlagOk

`func (o *Board) GetBillUnapprovedTimeExpenseFlagOk() (*bool, bool)`

GetBillUnapprovedTimeExpenseFlagOk returns a tuple with the BillUnapprovedTimeExpenseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillUnapprovedTimeExpenseFlag

`func (o *Board) SetBillUnapprovedTimeExpenseFlag(v bool)`

SetBillUnapprovedTimeExpenseFlag sets BillUnapprovedTimeExpenseFlag field to given value.

### HasBillUnapprovedTimeExpenseFlag

`func (o *Board) HasBillUnapprovedTimeExpenseFlag() bool`

HasBillUnapprovedTimeExpenseFlag returns a boolean if a field has been set.

### SetBillUnapprovedTimeExpenseFlagNil

`func (o *Board) SetBillUnapprovedTimeExpenseFlagNil(b bool)`

 SetBillUnapprovedTimeExpenseFlagNil sets the value for BillUnapprovedTimeExpenseFlag to be an explicit nil

### UnsetBillUnapprovedTimeExpenseFlag
`func (o *Board) UnsetBillUnapprovedTimeExpenseFlag()`

UnsetBillUnapprovedTimeExpenseFlag ensures that no value is present for BillUnapprovedTimeExpenseFlag, not even an explicit nil
### GetOverrideBillingSetupFlag

`func (o *Board) GetOverrideBillingSetupFlag() bool`

GetOverrideBillingSetupFlag returns the OverrideBillingSetupFlag field if non-nil, zero value otherwise.

### GetOverrideBillingSetupFlagOk

`func (o *Board) GetOverrideBillingSetupFlagOk() (*bool, bool)`

GetOverrideBillingSetupFlagOk returns a tuple with the OverrideBillingSetupFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideBillingSetupFlag

`func (o *Board) SetOverrideBillingSetupFlag(v bool)`

SetOverrideBillingSetupFlag sets OverrideBillingSetupFlag field to given value.

### HasOverrideBillingSetupFlag

`func (o *Board) HasOverrideBillingSetupFlag() bool`

HasOverrideBillingSetupFlag returns a boolean if a field has been set.

### SetOverrideBillingSetupFlagNil

`func (o *Board) SetOverrideBillingSetupFlagNil(b bool)`

 SetOverrideBillingSetupFlagNil sets the value for OverrideBillingSetupFlag to be an explicit nil

### UnsetOverrideBillingSetupFlag
`func (o *Board) UnsetOverrideBillingSetupFlag()`

UnsetOverrideBillingSetupFlag ensures that no value is present for OverrideBillingSetupFlag, not even an explicit nil
### GetDispatchMember

`func (o *Board) GetDispatchMember() MemberReference`

GetDispatchMember returns the DispatchMember field if non-nil, zero value otherwise.

### GetDispatchMemberOk

`func (o *Board) GetDispatchMemberOk() (*MemberReference, bool)`

GetDispatchMemberOk returns a tuple with the DispatchMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchMember

`func (o *Board) SetDispatchMember(v MemberReference)`

SetDispatchMember sets DispatchMember field to given value.

### HasDispatchMember

`func (o *Board) HasDispatchMember() bool`

HasDispatchMember returns a boolean if a field has been set.

### GetServiceManagerMember

`func (o *Board) GetServiceManagerMember() MemberReference`

GetServiceManagerMember returns the ServiceManagerMember field if non-nil, zero value otherwise.

### GetServiceManagerMemberOk

`func (o *Board) GetServiceManagerMemberOk() (*MemberReference, bool)`

GetServiceManagerMemberOk returns a tuple with the ServiceManagerMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceManagerMember

`func (o *Board) SetServiceManagerMember(v MemberReference)`

SetServiceManagerMember sets ServiceManagerMember field to given value.

### HasServiceManagerMember

`func (o *Board) HasServiceManagerMember() bool`

HasServiceManagerMember returns a boolean if a field has been set.

### GetDutyManagerMember

`func (o *Board) GetDutyManagerMember() MemberReference`

GetDutyManagerMember returns the DutyManagerMember field if non-nil, zero value otherwise.

### GetDutyManagerMemberOk

`func (o *Board) GetDutyManagerMemberOk() (*MemberReference, bool)`

GetDutyManagerMemberOk returns a tuple with the DutyManagerMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyManagerMember

`func (o *Board) SetDutyManagerMember(v MemberReference)`

SetDutyManagerMember sets DutyManagerMember field to given value.

### HasDutyManagerMember

`func (o *Board) HasDutyManagerMember() bool`

HasDutyManagerMember returns a boolean if a field has been set.

### GetOncallMember

`func (o *Board) GetOncallMember() MemberReference`

GetOncallMember returns the OncallMember field if non-nil, zero value otherwise.

### GetOncallMemberOk

`func (o *Board) GetOncallMemberOk() (*MemberReference, bool)`

GetOncallMemberOk returns a tuple with the OncallMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOncallMember

`func (o *Board) SetOncallMember(v MemberReference)`

SetOncallMember sets OncallMember field to given value.

### HasOncallMember

`func (o *Board) HasOncallMember() bool`

HasOncallMember returns a boolean if a field has been set.

### GetWorkRole

`func (o *Board) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *Board) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *Board) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *Board) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *Board) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *Board) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *Board) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *Board) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetBillTime

`func (o *Board) GetBillTime() string`

GetBillTime returns the BillTime field if non-nil, zero value otherwise.

### GetBillTimeOk

`func (o *Board) GetBillTimeOk() (*string, bool)`

GetBillTimeOk returns a tuple with the BillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTime

`func (o *Board) SetBillTime(v string)`

SetBillTime sets BillTime field to given value.

### HasBillTime

`func (o *Board) HasBillTime() bool`

HasBillTime returns a boolean if a field has been set.

### SetBillTimeNil

`func (o *Board) SetBillTimeNil(b bool)`

 SetBillTimeNil sets the value for BillTime to be an explicit nil

### UnsetBillTime
`func (o *Board) UnsetBillTime()`

UnsetBillTime ensures that no value is present for BillTime, not even an explicit nil
### GetBillExpense

`func (o *Board) GetBillExpense() string`

GetBillExpense returns the BillExpense field if non-nil, zero value otherwise.

### GetBillExpenseOk

`func (o *Board) GetBillExpenseOk() (*string, bool)`

GetBillExpenseOk returns a tuple with the BillExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillExpense

`func (o *Board) SetBillExpense(v string)`

SetBillExpense sets BillExpense field to given value.

### HasBillExpense

`func (o *Board) HasBillExpense() bool`

HasBillExpense returns a boolean if a field has been set.

### SetBillExpenseNil

`func (o *Board) SetBillExpenseNil(b bool)`

 SetBillExpenseNil sets the value for BillExpense to be an explicit nil

### UnsetBillExpense
`func (o *Board) UnsetBillExpense()`

UnsetBillExpense ensures that no value is present for BillExpense, not even an explicit nil
### GetBillProduct

`func (o *Board) GetBillProduct() string`

GetBillProduct returns the BillProduct field if non-nil, zero value otherwise.

### GetBillProductOk

`func (o *Board) GetBillProductOk() (*string, bool)`

GetBillProductOk returns a tuple with the BillProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProduct

`func (o *Board) SetBillProduct(v string)`

SetBillProduct sets BillProduct field to given value.

### HasBillProduct

`func (o *Board) HasBillProduct() bool`

HasBillProduct returns a boolean if a field has been set.

### SetBillProductNil

`func (o *Board) SetBillProductNil(b bool)`

 SetBillProductNil sets the value for BillProduct to be an explicit nil

### UnsetBillProduct
`func (o *Board) UnsetBillProduct()`

UnsetBillProduct ensures that no value is present for BillProduct, not even an explicit nil
### GetAutoCloseStatus

`func (o *Board) GetAutoCloseStatus() ServiceStatusReference`

GetAutoCloseStatus returns the AutoCloseStatus field if non-nil, zero value otherwise.

### GetAutoCloseStatusOk

`func (o *Board) GetAutoCloseStatusOk() (*ServiceStatusReference, bool)`

GetAutoCloseStatusOk returns a tuple with the AutoCloseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCloseStatus

`func (o *Board) SetAutoCloseStatus(v ServiceStatusReference)`

SetAutoCloseStatus sets AutoCloseStatus field to given value.

### HasAutoCloseStatus

`func (o *Board) HasAutoCloseStatus() bool`

HasAutoCloseStatus returns a boolean if a field has been set.

### GetAutoAssignNewTicketsFlag

`func (o *Board) GetAutoAssignNewTicketsFlag() bool`

GetAutoAssignNewTicketsFlag returns the AutoAssignNewTicketsFlag field if non-nil, zero value otherwise.

### GetAutoAssignNewTicketsFlagOk

`func (o *Board) GetAutoAssignNewTicketsFlagOk() (*bool, bool)`

GetAutoAssignNewTicketsFlagOk returns a tuple with the AutoAssignNewTicketsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAssignNewTicketsFlag

`func (o *Board) SetAutoAssignNewTicketsFlag(v bool)`

SetAutoAssignNewTicketsFlag sets AutoAssignNewTicketsFlag field to given value.

### HasAutoAssignNewTicketsFlag

`func (o *Board) HasAutoAssignNewTicketsFlag() bool`

HasAutoAssignNewTicketsFlag returns a boolean if a field has been set.

### SetAutoAssignNewTicketsFlagNil

`func (o *Board) SetAutoAssignNewTicketsFlagNil(b bool)`

 SetAutoAssignNewTicketsFlagNil sets the value for AutoAssignNewTicketsFlag to be an explicit nil

### UnsetAutoAssignNewTicketsFlag
`func (o *Board) UnsetAutoAssignNewTicketsFlag()`

UnsetAutoAssignNewTicketsFlag ensures that no value is present for AutoAssignNewTicketsFlag, not even an explicit nil
### GetAutoAssignNewECTicketsFlag

`func (o *Board) GetAutoAssignNewECTicketsFlag() bool`

GetAutoAssignNewECTicketsFlag returns the AutoAssignNewECTicketsFlag field if non-nil, zero value otherwise.

### GetAutoAssignNewECTicketsFlagOk

`func (o *Board) GetAutoAssignNewECTicketsFlagOk() (*bool, bool)`

GetAutoAssignNewECTicketsFlagOk returns a tuple with the AutoAssignNewECTicketsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAssignNewECTicketsFlag

`func (o *Board) SetAutoAssignNewECTicketsFlag(v bool)`

SetAutoAssignNewECTicketsFlag sets AutoAssignNewECTicketsFlag field to given value.

### HasAutoAssignNewECTicketsFlag

`func (o *Board) HasAutoAssignNewECTicketsFlag() bool`

HasAutoAssignNewECTicketsFlag returns a boolean if a field has been set.

### SetAutoAssignNewECTicketsFlagNil

`func (o *Board) SetAutoAssignNewECTicketsFlagNil(b bool)`

 SetAutoAssignNewECTicketsFlagNil sets the value for AutoAssignNewECTicketsFlag to be an explicit nil

### UnsetAutoAssignNewECTicketsFlag
`func (o *Board) UnsetAutoAssignNewECTicketsFlag()`

UnsetAutoAssignNewECTicketsFlag ensures that no value is present for AutoAssignNewECTicketsFlag, not even an explicit nil
### GetAutoAssignNewPortalTicketsFlag

`func (o *Board) GetAutoAssignNewPortalTicketsFlag() bool`

GetAutoAssignNewPortalTicketsFlag returns the AutoAssignNewPortalTicketsFlag field if non-nil, zero value otherwise.

### GetAutoAssignNewPortalTicketsFlagOk

`func (o *Board) GetAutoAssignNewPortalTicketsFlagOk() (*bool, bool)`

GetAutoAssignNewPortalTicketsFlagOk returns a tuple with the AutoAssignNewPortalTicketsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAssignNewPortalTicketsFlag

`func (o *Board) SetAutoAssignNewPortalTicketsFlag(v bool)`

SetAutoAssignNewPortalTicketsFlag sets AutoAssignNewPortalTicketsFlag field to given value.

### HasAutoAssignNewPortalTicketsFlag

`func (o *Board) HasAutoAssignNewPortalTicketsFlag() bool`

HasAutoAssignNewPortalTicketsFlag returns a boolean if a field has been set.

### SetAutoAssignNewPortalTicketsFlagNil

`func (o *Board) SetAutoAssignNewPortalTicketsFlagNil(b bool)`

 SetAutoAssignNewPortalTicketsFlagNil sets the value for AutoAssignNewPortalTicketsFlag to be an explicit nil

### UnsetAutoAssignNewPortalTicketsFlag
`func (o *Board) UnsetAutoAssignNewPortalTicketsFlag()`

UnsetAutoAssignNewPortalTicketsFlag ensures that no value is present for AutoAssignNewPortalTicketsFlag, not even an explicit nil
### GetDiscussionsLockedFlag

`func (o *Board) GetDiscussionsLockedFlag() bool`

GetDiscussionsLockedFlag returns the DiscussionsLockedFlag field if non-nil, zero value otherwise.

### GetDiscussionsLockedFlagOk

`func (o *Board) GetDiscussionsLockedFlagOk() (*bool, bool)`

GetDiscussionsLockedFlagOk returns a tuple with the DiscussionsLockedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscussionsLockedFlag

`func (o *Board) SetDiscussionsLockedFlag(v bool)`

SetDiscussionsLockedFlag sets DiscussionsLockedFlag field to given value.

### HasDiscussionsLockedFlag

`func (o *Board) HasDiscussionsLockedFlag() bool`

HasDiscussionsLockedFlag returns a boolean if a field has been set.

### SetDiscussionsLockedFlagNil

`func (o *Board) SetDiscussionsLockedFlagNil(b bool)`

 SetDiscussionsLockedFlagNil sets the value for DiscussionsLockedFlag to be an explicit nil

### UnsetDiscussionsLockedFlag
`func (o *Board) UnsetDiscussionsLockedFlag()`

UnsetDiscussionsLockedFlag ensures that no value is present for DiscussionsLockedFlag, not even an explicit nil
### GetTimeEntryLockedFlag

`func (o *Board) GetTimeEntryLockedFlag() bool`

GetTimeEntryLockedFlag returns the TimeEntryLockedFlag field if non-nil, zero value otherwise.

### GetTimeEntryLockedFlagOk

`func (o *Board) GetTimeEntryLockedFlagOk() (*bool, bool)`

GetTimeEntryLockedFlagOk returns a tuple with the TimeEntryLockedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntryLockedFlag

`func (o *Board) SetTimeEntryLockedFlag(v bool)`

SetTimeEntryLockedFlag sets TimeEntryLockedFlag field to given value.

### HasTimeEntryLockedFlag

`func (o *Board) HasTimeEntryLockedFlag() bool`

HasTimeEntryLockedFlag returns a boolean if a field has been set.

### SetTimeEntryLockedFlagNil

`func (o *Board) SetTimeEntryLockedFlagNil(b bool)`

 SetTimeEntryLockedFlagNil sets the value for TimeEntryLockedFlag to be an explicit nil

### UnsetTimeEntryLockedFlag
`func (o *Board) UnsetTimeEntryLockedFlag()`

UnsetTimeEntryLockedFlag ensures that no value is present for TimeEntryLockedFlag, not even an explicit nil
### GetNotifyEmailFrom

`func (o *Board) GetNotifyEmailFrom() string`

GetNotifyEmailFrom returns the NotifyEmailFrom field if non-nil, zero value otherwise.

### GetNotifyEmailFromOk

`func (o *Board) GetNotifyEmailFromOk() (*string, bool)`

GetNotifyEmailFromOk returns a tuple with the NotifyEmailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEmailFrom

`func (o *Board) SetNotifyEmailFrom(v string)`

SetNotifyEmailFrom sets NotifyEmailFrom field to given value.

### HasNotifyEmailFrom

`func (o *Board) HasNotifyEmailFrom() bool`

HasNotifyEmailFrom returns a boolean if a field has been set.

### GetNotifyEmailFromName

`func (o *Board) GetNotifyEmailFromName() string`

GetNotifyEmailFromName returns the NotifyEmailFromName field if non-nil, zero value otherwise.

### GetNotifyEmailFromNameOk

`func (o *Board) GetNotifyEmailFromNameOk() (*string, bool)`

GetNotifyEmailFromNameOk returns a tuple with the NotifyEmailFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEmailFromName

`func (o *Board) SetNotifyEmailFromName(v string)`

SetNotifyEmailFromName sets NotifyEmailFromName field to given value.

### HasNotifyEmailFromName

`func (o *Board) HasNotifyEmailFromName() bool`

HasNotifyEmailFromName returns a boolean if a field has been set.

### GetClosedLoopDiscussionsFlag

`func (o *Board) GetClosedLoopDiscussionsFlag() bool`

GetClosedLoopDiscussionsFlag returns the ClosedLoopDiscussionsFlag field if non-nil, zero value otherwise.

### GetClosedLoopDiscussionsFlagOk

`func (o *Board) GetClosedLoopDiscussionsFlagOk() (*bool, bool)`

GetClosedLoopDiscussionsFlagOk returns a tuple with the ClosedLoopDiscussionsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedLoopDiscussionsFlag

`func (o *Board) SetClosedLoopDiscussionsFlag(v bool)`

SetClosedLoopDiscussionsFlag sets ClosedLoopDiscussionsFlag field to given value.

### HasClosedLoopDiscussionsFlag

`func (o *Board) HasClosedLoopDiscussionsFlag() bool`

HasClosedLoopDiscussionsFlag returns a boolean if a field has been set.

### SetClosedLoopDiscussionsFlagNil

`func (o *Board) SetClosedLoopDiscussionsFlagNil(b bool)`

 SetClosedLoopDiscussionsFlagNil sets the value for ClosedLoopDiscussionsFlag to be an explicit nil

### UnsetClosedLoopDiscussionsFlag
`func (o *Board) UnsetClosedLoopDiscussionsFlag()`

UnsetClosedLoopDiscussionsFlag ensures that no value is present for ClosedLoopDiscussionsFlag, not even an explicit nil
### GetClosedLoopResolutionFlag

`func (o *Board) GetClosedLoopResolutionFlag() bool`

GetClosedLoopResolutionFlag returns the ClosedLoopResolutionFlag field if non-nil, zero value otherwise.

### GetClosedLoopResolutionFlagOk

`func (o *Board) GetClosedLoopResolutionFlagOk() (*bool, bool)`

GetClosedLoopResolutionFlagOk returns a tuple with the ClosedLoopResolutionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedLoopResolutionFlag

`func (o *Board) SetClosedLoopResolutionFlag(v bool)`

SetClosedLoopResolutionFlag sets ClosedLoopResolutionFlag field to given value.

### HasClosedLoopResolutionFlag

`func (o *Board) HasClosedLoopResolutionFlag() bool`

HasClosedLoopResolutionFlag returns a boolean if a field has been set.

### SetClosedLoopResolutionFlagNil

`func (o *Board) SetClosedLoopResolutionFlagNil(b bool)`

 SetClosedLoopResolutionFlagNil sets the value for ClosedLoopResolutionFlag to be an explicit nil

### UnsetClosedLoopResolutionFlag
`func (o *Board) UnsetClosedLoopResolutionFlag()`

UnsetClosedLoopResolutionFlag ensures that no value is present for ClosedLoopResolutionFlag, not even an explicit nil
### GetClosedLoopInternalAnalysisFlag

`func (o *Board) GetClosedLoopInternalAnalysisFlag() bool`

GetClosedLoopInternalAnalysisFlag returns the ClosedLoopInternalAnalysisFlag field if non-nil, zero value otherwise.

### GetClosedLoopInternalAnalysisFlagOk

`func (o *Board) GetClosedLoopInternalAnalysisFlagOk() (*bool, bool)`

GetClosedLoopInternalAnalysisFlagOk returns a tuple with the ClosedLoopInternalAnalysisFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedLoopInternalAnalysisFlag

`func (o *Board) SetClosedLoopInternalAnalysisFlag(v bool)`

SetClosedLoopInternalAnalysisFlag sets ClosedLoopInternalAnalysisFlag field to given value.

### HasClosedLoopInternalAnalysisFlag

`func (o *Board) HasClosedLoopInternalAnalysisFlag() bool`

HasClosedLoopInternalAnalysisFlag returns a boolean if a field has been set.

### SetClosedLoopInternalAnalysisFlagNil

`func (o *Board) SetClosedLoopInternalAnalysisFlagNil(b bool)`

 SetClosedLoopInternalAnalysisFlagNil sets the value for ClosedLoopInternalAnalysisFlag to be an explicit nil

### UnsetClosedLoopInternalAnalysisFlag
`func (o *Board) UnsetClosedLoopInternalAnalysisFlag()`

UnsetClosedLoopInternalAnalysisFlag ensures that no value is present for ClosedLoopInternalAnalysisFlag, not even an explicit nil
### GetTimeEntryDiscussionFlag

`func (o *Board) GetTimeEntryDiscussionFlag() bool`

GetTimeEntryDiscussionFlag returns the TimeEntryDiscussionFlag field if non-nil, zero value otherwise.

### GetTimeEntryDiscussionFlagOk

`func (o *Board) GetTimeEntryDiscussionFlagOk() (*bool, bool)`

GetTimeEntryDiscussionFlagOk returns a tuple with the TimeEntryDiscussionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntryDiscussionFlag

`func (o *Board) SetTimeEntryDiscussionFlag(v bool)`

SetTimeEntryDiscussionFlag sets TimeEntryDiscussionFlag field to given value.

### HasTimeEntryDiscussionFlag

`func (o *Board) HasTimeEntryDiscussionFlag() bool`

HasTimeEntryDiscussionFlag returns a boolean if a field has been set.

### SetTimeEntryDiscussionFlagNil

`func (o *Board) SetTimeEntryDiscussionFlagNil(b bool)`

 SetTimeEntryDiscussionFlagNil sets the value for TimeEntryDiscussionFlag to be an explicit nil

### UnsetTimeEntryDiscussionFlag
`func (o *Board) UnsetTimeEntryDiscussionFlag()`

UnsetTimeEntryDiscussionFlag ensures that no value is present for TimeEntryDiscussionFlag, not even an explicit nil
### GetTimeEntryResolutionFlag

`func (o *Board) GetTimeEntryResolutionFlag() bool`

GetTimeEntryResolutionFlag returns the TimeEntryResolutionFlag field if non-nil, zero value otherwise.

### GetTimeEntryResolutionFlagOk

`func (o *Board) GetTimeEntryResolutionFlagOk() (*bool, bool)`

GetTimeEntryResolutionFlagOk returns a tuple with the TimeEntryResolutionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntryResolutionFlag

`func (o *Board) SetTimeEntryResolutionFlag(v bool)`

SetTimeEntryResolutionFlag sets TimeEntryResolutionFlag field to given value.

### HasTimeEntryResolutionFlag

`func (o *Board) HasTimeEntryResolutionFlag() bool`

HasTimeEntryResolutionFlag returns a boolean if a field has been set.

### SetTimeEntryResolutionFlagNil

`func (o *Board) SetTimeEntryResolutionFlagNil(b bool)`

 SetTimeEntryResolutionFlagNil sets the value for TimeEntryResolutionFlag to be an explicit nil

### UnsetTimeEntryResolutionFlag
`func (o *Board) UnsetTimeEntryResolutionFlag()`

UnsetTimeEntryResolutionFlag ensures that no value is present for TimeEntryResolutionFlag, not even an explicit nil
### GetTimeEntryInternalAnalysisFlag

`func (o *Board) GetTimeEntryInternalAnalysisFlag() bool`

GetTimeEntryInternalAnalysisFlag returns the TimeEntryInternalAnalysisFlag field if non-nil, zero value otherwise.

### GetTimeEntryInternalAnalysisFlagOk

`func (o *Board) GetTimeEntryInternalAnalysisFlagOk() (*bool, bool)`

GetTimeEntryInternalAnalysisFlagOk returns a tuple with the TimeEntryInternalAnalysisFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntryInternalAnalysisFlag

`func (o *Board) SetTimeEntryInternalAnalysisFlag(v bool)`

SetTimeEntryInternalAnalysisFlag sets TimeEntryInternalAnalysisFlag field to given value.

### HasTimeEntryInternalAnalysisFlag

`func (o *Board) HasTimeEntryInternalAnalysisFlag() bool`

HasTimeEntryInternalAnalysisFlag returns a boolean if a field has been set.

### SetTimeEntryInternalAnalysisFlagNil

`func (o *Board) SetTimeEntryInternalAnalysisFlagNil(b bool)`

 SetTimeEntryInternalAnalysisFlagNil sets the value for TimeEntryInternalAnalysisFlag to be an explicit nil

### UnsetTimeEntryInternalAnalysisFlag
`func (o *Board) UnsetTimeEntryInternalAnalysisFlag()`

UnsetTimeEntryInternalAnalysisFlag ensures that no value is present for TimeEntryInternalAnalysisFlag, not even an explicit nil
### GetProblemSort

`func (o *Board) GetProblemSort() string`

GetProblemSort returns the ProblemSort field if non-nil, zero value otherwise.

### GetProblemSortOk

`func (o *Board) GetProblemSortOk() (*string, bool)`

GetProblemSortOk returns a tuple with the ProblemSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemSort

`func (o *Board) SetProblemSort(v string)`

SetProblemSort sets ProblemSort field to given value.

### HasProblemSort

`func (o *Board) HasProblemSort() bool`

HasProblemSort returns a boolean if a field has been set.

### SetProblemSortNil

`func (o *Board) SetProblemSortNil(b bool)`

 SetProblemSortNil sets the value for ProblemSort to be an explicit nil

### UnsetProblemSort
`func (o *Board) UnsetProblemSort()`

UnsetProblemSort ensures that no value is present for ProblemSort, not even an explicit nil
### GetResolutionSort

`func (o *Board) GetResolutionSort() string`

GetResolutionSort returns the ResolutionSort field if non-nil, zero value otherwise.

### GetResolutionSortOk

`func (o *Board) GetResolutionSortOk() (*string, bool)`

GetResolutionSortOk returns a tuple with the ResolutionSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionSort

`func (o *Board) SetResolutionSort(v string)`

SetResolutionSort sets ResolutionSort field to given value.

### HasResolutionSort

`func (o *Board) HasResolutionSort() bool`

HasResolutionSort returns a boolean if a field has been set.

### SetResolutionSortNil

`func (o *Board) SetResolutionSortNil(b bool)`

 SetResolutionSortNil sets the value for ResolutionSort to be an explicit nil

### UnsetResolutionSort
`func (o *Board) UnsetResolutionSort()`

UnsetResolutionSort ensures that no value is present for ResolutionSort, not even an explicit nil
### GetInternalAnalysisSort

`func (o *Board) GetInternalAnalysisSort() string`

GetInternalAnalysisSort returns the InternalAnalysisSort field if non-nil, zero value otherwise.

### GetInternalAnalysisSortOk

`func (o *Board) GetInternalAnalysisSortOk() (*string, bool)`

GetInternalAnalysisSortOk returns a tuple with the InternalAnalysisSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAnalysisSort

`func (o *Board) SetInternalAnalysisSort(v string)`

SetInternalAnalysisSort sets InternalAnalysisSort field to given value.

### HasInternalAnalysisSort

`func (o *Board) HasInternalAnalysisSort() bool`

HasInternalAnalysisSort returns a boolean if a field has been set.

### SetInternalAnalysisSortNil

`func (o *Board) SetInternalAnalysisSortNil(b bool)`

 SetInternalAnalysisSortNil sets the value for InternalAnalysisSort to be an explicit nil

### UnsetInternalAnalysisSort
`func (o *Board) UnsetInternalAnalysisSort()`

UnsetInternalAnalysisSort ensures that no value is present for InternalAnalysisSort, not even an explicit nil
### GetEmailConnectorAllowReopenClosedFlag

`func (o *Board) GetEmailConnectorAllowReopenClosedFlag() bool`

GetEmailConnectorAllowReopenClosedFlag returns the EmailConnectorAllowReopenClosedFlag field if non-nil, zero value otherwise.

### GetEmailConnectorAllowReopenClosedFlagOk

`func (o *Board) GetEmailConnectorAllowReopenClosedFlagOk() (*bool, bool)`

GetEmailConnectorAllowReopenClosedFlagOk returns a tuple with the EmailConnectorAllowReopenClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConnectorAllowReopenClosedFlag

`func (o *Board) SetEmailConnectorAllowReopenClosedFlag(v bool)`

SetEmailConnectorAllowReopenClosedFlag sets EmailConnectorAllowReopenClosedFlag field to given value.

### HasEmailConnectorAllowReopenClosedFlag

`func (o *Board) HasEmailConnectorAllowReopenClosedFlag() bool`

HasEmailConnectorAllowReopenClosedFlag returns a boolean if a field has been set.

### SetEmailConnectorAllowReopenClosedFlagNil

`func (o *Board) SetEmailConnectorAllowReopenClosedFlagNil(b bool)`

 SetEmailConnectorAllowReopenClosedFlagNil sets the value for EmailConnectorAllowReopenClosedFlag to be an explicit nil

### UnsetEmailConnectorAllowReopenClosedFlag
`func (o *Board) UnsetEmailConnectorAllowReopenClosedFlag()`

UnsetEmailConnectorAllowReopenClosedFlag ensures that no value is present for EmailConnectorAllowReopenClosedFlag, not even an explicit nil
### GetEmailConnectorReopenStatus

`func (o *Board) GetEmailConnectorReopenStatus() ServiceStatusReference`

GetEmailConnectorReopenStatus returns the EmailConnectorReopenStatus field if non-nil, zero value otherwise.

### GetEmailConnectorReopenStatusOk

`func (o *Board) GetEmailConnectorReopenStatusOk() (*ServiceStatusReference, bool)`

GetEmailConnectorReopenStatusOk returns a tuple with the EmailConnectorReopenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConnectorReopenStatus

`func (o *Board) SetEmailConnectorReopenStatus(v ServiceStatusReference)`

SetEmailConnectorReopenStatus sets EmailConnectorReopenStatus field to given value.

### HasEmailConnectorReopenStatus

`func (o *Board) HasEmailConnectorReopenStatus() bool`

HasEmailConnectorReopenStatus returns a boolean if a field has been set.

### GetEmailConnectorReopenResourcesFlag

`func (o *Board) GetEmailConnectorReopenResourcesFlag() bool`

GetEmailConnectorReopenResourcesFlag returns the EmailConnectorReopenResourcesFlag field if non-nil, zero value otherwise.

### GetEmailConnectorReopenResourcesFlagOk

`func (o *Board) GetEmailConnectorReopenResourcesFlagOk() (*bool, bool)`

GetEmailConnectorReopenResourcesFlagOk returns a tuple with the EmailConnectorReopenResourcesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConnectorReopenResourcesFlag

`func (o *Board) SetEmailConnectorReopenResourcesFlag(v bool)`

SetEmailConnectorReopenResourcesFlag sets EmailConnectorReopenResourcesFlag field to given value.

### HasEmailConnectorReopenResourcesFlag

`func (o *Board) HasEmailConnectorReopenResourcesFlag() bool`

HasEmailConnectorReopenResourcesFlag returns a boolean if a field has been set.

### SetEmailConnectorReopenResourcesFlagNil

`func (o *Board) SetEmailConnectorReopenResourcesFlagNil(b bool)`

 SetEmailConnectorReopenResourcesFlagNil sets the value for EmailConnectorReopenResourcesFlag to be an explicit nil

### UnsetEmailConnectorReopenResourcesFlag
`func (o *Board) UnsetEmailConnectorReopenResourcesFlag()`

UnsetEmailConnectorReopenResourcesFlag ensures that no value is present for EmailConnectorReopenResourcesFlag, not even an explicit nil
### GetEmailConnectorNewTicketNoMatchFlag

`func (o *Board) GetEmailConnectorNewTicketNoMatchFlag() bool`

GetEmailConnectorNewTicketNoMatchFlag returns the EmailConnectorNewTicketNoMatchFlag field if non-nil, zero value otherwise.

### GetEmailConnectorNewTicketNoMatchFlagOk

`func (o *Board) GetEmailConnectorNewTicketNoMatchFlagOk() (*bool, bool)`

GetEmailConnectorNewTicketNoMatchFlagOk returns a tuple with the EmailConnectorNewTicketNoMatchFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConnectorNewTicketNoMatchFlag

`func (o *Board) SetEmailConnectorNewTicketNoMatchFlag(v bool)`

SetEmailConnectorNewTicketNoMatchFlag sets EmailConnectorNewTicketNoMatchFlag field to given value.

### HasEmailConnectorNewTicketNoMatchFlag

`func (o *Board) HasEmailConnectorNewTicketNoMatchFlag() bool`

HasEmailConnectorNewTicketNoMatchFlag returns a boolean if a field has been set.

### SetEmailConnectorNewTicketNoMatchFlagNil

`func (o *Board) SetEmailConnectorNewTicketNoMatchFlagNil(b bool)`

 SetEmailConnectorNewTicketNoMatchFlagNil sets the value for EmailConnectorNewTicketNoMatchFlag to be an explicit nil

### UnsetEmailConnectorNewTicketNoMatchFlag
`func (o *Board) UnsetEmailConnectorNewTicketNoMatchFlag()`

UnsetEmailConnectorNewTicketNoMatchFlag ensures that no value is present for EmailConnectorNewTicketNoMatchFlag, not even an explicit nil
### GetEmailConnectorNeverReopenByDaysFlag

`func (o *Board) GetEmailConnectorNeverReopenByDaysFlag() bool`

GetEmailConnectorNeverReopenByDaysFlag returns the EmailConnectorNeverReopenByDaysFlag field if non-nil, zero value otherwise.

### GetEmailConnectorNeverReopenByDaysFlagOk

`func (o *Board) GetEmailConnectorNeverReopenByDaysFlagOk() (*bool, bool)`

GetEmailConnectorNeverReopenByDaysFlagOk returns a tuple with the EmailConnectorNeverReopenByDaysFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConnectorNeverReopenByDaysFlag

`func (o *Board) SetEmailConnectorNeverReopenByDaysFlag(v bool)`

SetEmailConnectorNeverReopenByDaysFlag sets EmailConnectorNeverReopenByDaysFlag field to given value.

### HasEmailConnectorNeverReopenByDaysFlag

`func (o *Board) HasEmailConnectorNeverReopenByDaysFlag() bool`

HasEmailConnectorNeverReopenByDaysFlag returns a boolean if a field has been set.

### SetEmailConnectorNeverReopenByDaysFlagNil

`func (o *Board) SetEmailConnectorNeverReopenByDaysFlagNil(b bool)`

 SetEmailConnectorNeverReopenByDaysFlagNil sets the value for EmailConnectorNeverReopenByDaysFlag to be an explicit nil

### UnsetEmailConnectorNeverReopenByDaysFlag
`func (o *Board) UnsetEmailConnectorNeverReopenByDaysFlag()`

UnsetEmailConnectorNeverReopenByDaysFlag ensures that no value is present for EmailConnectorNeverReopenByDaysFlag, not even an explicit nil
### GetEmailConnectorReopenDaysLimit

`func (o *Board) GetEmailConnectorReopenDaysLimit() int32`

GetEmailConnectorReopenDaysLimit returns the EmailConnectorReopenDaysLimit field if non-nil, zero value otherwise.

### GetEmailConnectorReopenDaysLimitOk

`func (o *Board) GetEmailConnectorReopenDaysLimitOk() (*int32, bool)`

GetEmailConnectorReopenDaysLimitOk returns a tuple with the EmailConnectorReopenDaysLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConnectorReopenDaysLimit

`func (o *Board) SetEmailConnectorReopenDaysLimit(v int32)`

SetEmailConnectorReopenDaysLimit sets EmailConnectorReopenDaysLimit field to given value.

### HasEmailConnectorReopenDaysLimit

`func (o *Board) HasEmailConnectorReopenDaysLimit() bool`

HasEmailConnectorReopenDaysLimit returns a boolean if a field has been set.

### SetEmailConnectorReopenDaysLimitNil

`func (o *Board) SetEmailConnectorReopenDaysLimitNil(b bool)`

 SetEmailConnectorReopenDaysLimitNil sets the value for EmailConnectorReopenDaysLimit to be an explicit nil

### UnsetEmailConnectorReopenDaysLimit
`func (o *Board) UnsetEmailConnectorReopenDaysLimit()`

UnsetEmailConnectorReopenDaysLimit ensures that no value is present for EmailConnectorReopenDaysLimit, not even an explicit nil
### GetEmailConnectorNeverReopenByDaysClosedFlag

`func (o *Board) GetEmailConnectorNeverReopenByDaysClosedFlag() bool`

GetEmailConnectorNeverReopenByDaysClosedFlag returns the EmailConnectorNeverReopenByDaysClosedFlag field if non-nil, zero value otherwise.

### GetEmailConnectorNeverReopenByDaysClosedFlagOk

`func (o *Board) GetEmailConnectorNeverReopenByDaysClosedFlagOk() (*bool, bool)`

GetEmailConnectorNeverReopenByDaysClosedFlagOk returns a tuple with the EmailConnectorNeverReopenByDaysClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConnectorNeverReopenByDaysClosedFlag

`func (o *Board) SetEmailConnectorNeverReopenByDaysClosedFlag(v bool)`

SetEmailConnectorNeverReopenByDaysClosedFlag sets EmailConnectorNeverReopenByDaysClosedFlag field to given value.

### HasEmailConnectorNeverReopenByDaysClosedFlag

`func (o *Board) HasEmailConnectorNeverReopenByDaysClosedFlag() bool`

HasEmailConnectorNeverReopenByDaysClosedFlag returns a boolean if a field has been set.

### SetEmailConnectorNeverReopenByDaysClosedFlagNil

`func (o *Board) SetEmailConnectorNeverReopenByDaysClosedFlagNil(b bool)`

 SetEmailConnectorNeverReopenByDaysClosedFlagNil sets the value for EmailConnectorNeverReopenByDaysClosedFlag to be an explicit nil

### UnsetEmailConnectorNeverReopenByDaysClosedFlag
`func (o *Board) UnsetEmailConnectorNeverReopenByDaysClosedFlag()`

UnsetEmailConnectorNeverReopenByDaysClosedFlag ensures that no value is present for EmailConnectorNeverReopenByDaysClosedFlag, not even an explicit nil
### GetEmailConnectorReopenDaysClosedLimit

`func (o *Board) GetEmailConnectorReopenDaysClosedLimit() int32`

GetEmailConnectorReopenDaysClosedLimit returns the EmailConnectorReopenDaysClosedLimit field if non-nil, zero value otherwise.

### GetEmailConnectorReopenDaysClosedLimitOk

`func (o *Board) GetEmailConnectorReopenDaysClosedLimitOk() (*int32, bool)`

GetEmailConnectorReopenDaysClosedLimitOk returns a tuple with the EmailConnectorReopenDaysClosedLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConnectorReopenDaysClosedLimit

`func (o *Board) SetEmailConnectorReopenDaysClosedLimit(v int32)`

SetEmailConnectorReopenDaysClosedLimit sets EmailConnectorReopenDaysClosedLimit field to given value.

### HasEmailConnectorReopenDaysClosedLimit

`func (o *Board) HasEmailConnectorReopenDaysClosedLimit() bool`

HasEmailConnectorReopenDaysClosedLimit returns a boolean if a field has been set.

### SetEmailConnectorReopenDaysClosedLimitNil

`func (o *Board) SetEmailConnectorReopenDaysClosedLimitNil(b bool)`

 SetEmailConnectorReopenDaysClosedLimitNil sets the value for EmailConnectorReopenDaysClosedLimit to be an explicit nil

### UnsetEmailConnectorReopenDaysClosedLimit
`func (o *Board) UnsetEmailConnectorReopenDaysClosedLimit()`

UnsetEmailConnectorReopenDaysClosedLimit ensures that no value is present for EmailConnectorReopenDaysClosedLimit, not even an explicit nil
### GetUseMemberDisplayNameFlag

`func (o *Board) GetUseMemberDisplayNameFlag() bool`

GetUseMemberDisplayNameFlag returns the UseMemberDisplayNameFlag field if non-nil, zero value otherwise.

### GetUseMemberDisplayNameFlagOk

`func (o *Board) GetUseMemberDisplayNameFlagOk() (*bool, bool)`

GetUseMemberDisplayNameFlagOk returns a tuple with the UseMemberDisplayNameFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMemberDisplayNameFlag

`func (o *Board) SetUseMemberDisplayNameFlag(v bool)`

SetUseMemberDisplayNameFlag sets UseMemberDisplayNameFlag field to given value.

### HasUseMemberDisplayNameFlag

`func (o *Board) HasUseMemberDisplayNameFlag() bool`

HasUseMemberDisplayNameFlag returns a boolean if a field has been set.

### SetUseMemberDisplayNameFlagNil

`func (o *Board) SetUseMemberDisplayNameFlagNil(b bool)`

 SetUseMemberDisplayNameFlagNil sets the value for UseMemberDisplayNameFlag to be an explicit nil

### UnsetUseMemberDisplayNameFlag
`func (o *Board) UnsetUseMemberDisplayNameFlag()`

UnsetUseMemberDisplayNameFlag ensures that no value is present for UseMemberDisplayNameFlag, not even an explicit nil
### GetSendToCCFlag

`func (o *Board) GetSendToCCFlag() bool`

GetSendToCCFlag returns the SendToCCFlag field if non-nil, zero value otherwise.

### GetSendToCCFlagOk

`func (o *Board) GetSendToCCFlagOk() (*bool, bool)`

GetSendToCCFlagOk returns a tuple with the SendToCCFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToCCFlag

`func (o *Board) SetSendToCCFlag(v bool)`

SetSendToCCFlag sets SendToCCFlag field to given value.

### HasSendToCCFlag

`func (o *Board) HasSendToCCFlag() bool`

HasSendToCCFlag returns a boolean if a field has been set.

### SetSendToCCFlagNil

`func (o *Board) SetSendToCCFlagNil(b bool)`

 SetSendToCCFlagNil sets the value for SendToCCFlag to be an explicit nil

### UnsetSendToCCFlag
`func (o *Board) UnsetSendToCCFlag()`

UnsetSendToCCFlag ensures that no value is present for SendToCCFlag, not even an explicit nil
### GetAutoAssignTicketOwnerFlag

`func (o *Board) GetAutoAssignTicketOwnerFlag() bool`

GetAutoAssignTicketOwnerFlag returns the AutoAssignTicketOwnerFlag field if non-nil, zero value otherwise.

### GetAutoAssignTicketOwnerFlagOk

`func (o *Board) GetAutoAssignTicketOwnerFlagOk() (*bool, bool)`

GetAutoAssignTicketOwnerFlagOk returns a tuple with the AutoAssignTicketOwnerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAssignTicketOwnerFlag

`func (o *Board) SetAutoAssignTicketOwnerFlag(v bool)`

SetAutoAssignTicketOwnerFlag sets AutoAssignTicketOwnerFlag field to given value.

### HasAutoAssignTicketOwnerFlag

`func (o *Board) HasAutoAssignTicketOwnerFlag() bool`

HasAutoAssignTicketOwnerFlag returns a boolean if a field has been set.

### SetAutoAssignTicketOwnerFlagNil

`func (o *Board) SetAutoAssignTicketOwnerFlagNil(b bool)`

 SetAutoAssignTicketOwnerFlagNil sets the value for AutoAssignTicketOwnerFlag to be an explicit nil

### UnsetAutoAssignTicketOwnerFlag
`func (o *Board) UnsetAutoAssignTicketOwnerFlag()`

UnsetAutoAssignTicketOwnerFlag ensures that no value is present for AutoAssignTicketOwnerFlag, not even an explicit nil
### GetAutoAssignLimitFlag

`func (o *Board) GetAutoAssignLimitFlag() bool`

GetAutoAssignLimitFlag returns the AutoAssignLimitFlag field if non-nil, zero value otherwise.

### GetAutoAssignLimitFlagOk

`func (o *Board) GetAutoAssignLimitFlagOk() (*bool, bool)`

GetAutoAssignLimitFlagOk returns a tuple with the AutoAssignLimitFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAssignLimitFlag

`func (o *Board) SetAutoAssignLimitFlag(v bool)`

SetAutoAssignLimitFlag sets AutoAssignLimitFlag field to given value.

### HasAutoAssignLimitFlag

`func (o *Board) HasAutoAssignLimitFlag() bool`

HasAutoAssignLimitFlag returns a boolean if a field has been set.

### SetAutoAssignLimitFlagNil

`func (o *Board) SetAutoAssignLimitFlagNil(b bool)`

 SetAutoAssignLimitFlagNil sets the value for AutoAssignLimitFlag to be an explicit nil

### UnsetAutoAssignLimitFlag
`func (o *Board) UnsetAutoAssignLimitFlag()`

UnsetAutoAssignLimitFlag ensures that no value is present for AutoAssignLimitFlag, not even an explicit nil
### GetAutoAssignLimitAmount

`func (o *Board) GetAutoAssignLimitAmount() int32`

GetAutoAssignLimitAmount returns the AutoAssignLimitAmount field if non-nil, zero value otherwise.

### GetAutoAssignLimitAmountOk

`func (o *Board) GetAutoAssignLimitAmountOk() (*int32, bool)`

GetAutoAssignLimitAmountOk returns a tuple with the AutoAssignLimitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAssignLimitAmount

`func (o *Board) SetAutoAssignLimitAmount(v int32)`

SetAutoAssignLimitAmount sets AutoAssignLimitAmount field to given value.

### HasAutoAssignLimitAmount

`func (o *Board) HasAutoAssignLimitAmount() bool`

HasAutoAssignLimitAmount returns a boolean if a field has been set.

### SetAutoAssignLimitAmountNil

`func (o *Board) SetAutoAssignLimitAmountNil(b bool)`

 SetAutoAssignLimitAmountNil sets the value for AutoAssignLimitAmount to be an explicit nil

### UnsetAutoAssignLimitAmount
`func (o *Board) UnsetAutoAssignLimitAmount()`

UnsetAutoAssignLimitAmount ensures that no value is present for AutoAssignLimitAmount, not even an explicit nil
### GetClosedLoopAllFlag

`func (o *Board) GetClosedLoopAllFlag() bool`

GetClosedLoopAllFlag returns the ClosedLoopAllFlag field if non-nil, zero value otherwise.

### GetClosedLoopAllFlagOk

`func (o *Board) GetClosedLoopAllFlagOk() (*bool, bool)`

GetClosedLoopAllFlagOk returns a tuple with the ClosedLoopAllFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedLoopAllFlag

`func (o *Board) SetClosedLoopAllFlag(v bool)`

SetClosedLoopAllFlag sets ClosedLoopAllFlag field to given value.

### HasClosedLoopAllFlag

`func (o *Board) HasClosedLoopAllFlag() bool`

HasClosedLoopAllFlag returns a boolean if a field has been set.

### SetClosedLoopAllFlagNil

`func (o *Board) SetClosedLoopAllFlagNil(b bool)`

 SetClosedLoopAllFlagNil sets the value for ClosedLoopAllFlag to be an explicit nil

### UnsetClosedLoopAllFlag
`func (o *Board) UnsetClosedLoopAllFlag()`

UnsetClosedLoopAllFlag ensures that no value is present for ClosedLoopAllFlag, not even an explicit nil
### GetPercentageCalculation

`func (o *Board) GetPercentageCalculation() string`

GetPercentageCalculation returns the PercentageCalculation field if non-nil, zero value otherwise.

### GetPercentageCalculationOk

`func (o *Board) GetPercentageCalculationOk() (*string, bool)`

GetPercentageCalculationOk returns a tuple with the PercentageCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCalculation

`func (o *Board) SetPercentageCalculation(v string)`

SetPercentageCalculation sets PercentageCalculation field to given value.

### HasPercentageCalculation

`func (o *Board) HasPercentageCalculation() bool`

HasPercentageCalculation returns a boolean if a field has been set.

### SetPercentageCalculationNil

`func (o *Board) SetPercentageCalculationNil(b bool)`

 SetPercentageCalculationNil sets the value for PercentageCalculation to be an explicit nil

### UnsetPercentageCalculation
`func (o *Board) UnsetPercentageCalculation()`

UnsetPercentageCalculation ensures that no value is present for PercentageCalculation, not even an explicit nil
### GetAllSort

`func (o *Board) GetAllSort() string`

GetAllSort returns the AllSort field if non-nil, zero value otherwise.

### GetAllSortOk

`func (o *Board) GetAllSortOk() (*string, bool)`

GetAllSortOk returns a tuple with the AllSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSort

`func (o *Board) SetAllSort(v string)`

SetAllSort sets AllSort field to given value.

### HasAllSort

`func (o *Board) HasAllSort() bool`

HasAllSort returns a boolean if a field has been set.

### SetAllSortNil

`func (o *Board) SetAllSortNil(b bool)`

 SetAllSortNil sets the value for AllSort to be an explicit nil

### UnsetAllSort
`func (o *Board) UnsetAllSort()`

UnsetAllSort ensures that no value is present for AllSort, not even an explicit nil
### GetMarkFirstNoteIssueFlag

`func (o *Board) GetMarkFirstNoteIssueFlag() bool`

GetMarkFirstNoteIssueFlag returns the MarkFirstNoteIssueFlag field if non-nil, zero value otherwise.

### GetMarkFirstNoteIssueFlagOk

`func (o *Board) GetMarkFirstNoteIssueFlagOk() (*bool, bool)`

GetMarkFirstNoteIssueFlagOk returns a tuple with the MarkFirstNoteIssueFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkFirstNoteIssueFlag

`func (o *Board) SetMarkFirstNoteIssueFlag(v bool)`

SetMarkFirstNoteIssueFlag sets MarkFirstNoteIssueFlag field to given value.

### HasMarkFirstNoteIssueFlag

`func (o *Board) HasMarkFirstNoteIssueFlag() bool`

HasMarkFirstNoteIssueFlag returns a boolean if a field has been set.

### SetMarkFirstNoteIssueFlagNil

`func (o *Board) SetMarkFirstNoteIssueFlagNil(b bool)`

 SetMarkFirstNoteIssueFlagNil sets the value for MarkFirstNoteIssueFlag to be an explicit nil

### UnsetMarkFirstNoteIssueFlag
`func (o *Board) UnsetMarkFirstNoteIssueFlag()`

UnsetMarkFirstNoteIssueFlag ensures that no value is present for MarkFirstNoteIssueFlag, not even an explicit nil
### GetRestrictBoardByDefaultFlag

`func (o *Board) GetRestrictBoardByDefaultFlag() bool`

GetRestrictBoardByDefaultFlag returns the RestrictBoardByDefaultFlag field if non-nil, zero value otherwise.

### GetRestrictBoardByDefaultFlagOk

`func (o *Board) GetRestrictBoardByDefaultFlagOk() (*bool, bool)`

GetRestrictBoardByDefaultFlagOk returns a tuple with the RestrictBoardByDefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictBoardByDefaultFlag

`func (o *Board) SetRestrictBoardByDefaultFlag(v bool)`

SetRestrictBoardByDefaultFlag sets RestrictBoardByDefaultFlag field to given value.

### HasRestrictBoardByDefaultFlag

`func (o *Board) HasRestrictBoardByDefaultFlag() bool`

HasRestrictBoardByDefaultFlag returns a boolean if a field has been set.

### SetRestrictBoardByDefaultFlagNil

`func (o *Board) SetRestrictBoardByDefaultFlagNil(b bool)`

 SetRestrictBoardByDefaultFlagNil sets the value for RestrictBoardByDefaultFlag to be an explicit nil

### UnsetRestrictBoardByDefaultFlag
`func (o *Board) UnsetRestrictBoardByDefaultFlag()`

UnsetRestrictBoardByDefaultFlag ensures that no value is present for RestrictBoardByDefaultFlag, not even an explicit nil
### GetSendToBundledFlag

`func (o *Board) GetSendToBundledFlag() bool`

GetSendToBundledFlag returns the SendToBundledFlag field if non-nil, zero value otherwise.

### GetSendToBundledFlagOk

`func (o *Board) GetSendToBundledFlagOk() (*bool, bool)`

GetSendToBundledFlagOk returns a tuple with the SendToBundledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToBundledFlag

`func (o *Board) SetSendToBundledFlag(v bool)`

SetSendToBundledFlag sets SendToBundledFlag field to given value.

### HasSendToBundledFlag

`func (o *Board) HasSendToBundledFlag() bool`

HasSendToBundledFlag returns a boolean if a field has been set.

### SetSendToBundledFlagNil

`func (o *Board) SetSendToBundledFlagNil(b bool)`

 SetSendToBundledFlagNil sets the value for SendToBundledFlag to be an explicit nil

### UnsetSendToBundledFlag
`func (o *Board) UnsetSendToBundledFlag()`

UnsetSendToBundledFlag ensures that no value is present for SendToBundledFlag, not even an explicit nil
### GetInfo

`func (o *Board) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Board) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Board) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Board) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


