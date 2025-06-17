# MemberDeactivation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**ServiceTeam** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**CompanyTeam** | Pointer to [**[]MemberDeactivationCompanyTeam**](MemberDeactivationCompanyTeam.md) | A list of customers for which the member holds a team role | [optional] 
**WorkflowEmail** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**ServiceStatusWorkflow** | Pointer to [**[]MemberDeactivationStatusWorkflow**](MemberDeactivationStatusWorkflow.md) |  | [optional] 
**TicketTemplate** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**Opportunity** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**SalesTeam** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**ProjectManager** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**ProjectTimeApprover** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**ProjectExpenseApprover** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**KnowledgeBaseArticle** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**MyCompanyPresident** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**MyCompanyCOO** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**MyCompanyController** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**MyCompanyDispatch** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**MyCompanyServiceManager** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**MyCompanyDutyManagerRole** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**DepartmentManager** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**DispatchMember** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**ServiceManager** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**DutyManager** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**SendFromEmailNotify** | Pointer to [**MemberDeactivationItem**](MemberDeactivationItem.md) |  | [optional] 
**DeleteOpenTimeSheetsFlag** | Pointer to **NullableBool** | By default, this is set to false             If there is any open timesheets, system will return error message             that there is open timesheets still attached to this member             If user would like to delete member with open timesheets, they can set this boolean to TRUE             System will delete member and any associated open timesheets | [optional] 

## Methods

### NewMemberDeactivation

`func NewMemberDeactivation() *MemberDeactivation`

NewMemberDeactivation instantiates a new MemberDeactivation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDeactivationWithDefaults

`func NewMemberDeactivationWithDefaults() *MemberDeactivation`

NewMemberDeactivationWithDefaults instantiates a new MemberDeactivation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *MemberDeactivation) GetActivity() MemberDeactivationItem`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *MemberDeactivation) GetActivityOk() (*MemberDeactivationItem, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *MemberDeactivation) SetActivity(v MemberDeactivationItem)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *MemberDeactivation) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetServiceTeam

`func (o *MemberDeactivation) GetServiceTeam() MemberDeactivationItem`

GetServiceTeam returns the ServiceTeam field if non-nil, zero value otherwise.

### GetServiceTeamOk

`func (o *MemberDeactivation) GetServiceTeamOk() (*MemberDeactivationItem, bool)`

GetServiceTeamOk returns a tuple with the ServiceTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTeam

`func (o *MemberDeactivation) SetServiceTeam(v MemberDeactivationItem)`

SetServiceTeam sets ServiceTeam field to given value.

### HasServiceTeam

`func (o *MemberDeactivation) HasServiceTeam() bool`

HasServiceTeam returns a boolean if a field has been set.

### GetCompanyTeam

`func (o *MemberDeactivation) GetCompanyTeam() []MemberDeactivationCompanyTeam`

GetCompanyTeam returns the CompanyTeam field if non-nil, zero value otherwise.

### GetCompanyTeamOk

`func (o *MemberDeactivation) GetCompanyTeamOk() (*[]MemberDeactivationCompanyTeam, bool)`

GetCompanyTeamOk returns a tuple with the CompanyTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyTeam

`func (o *MemberDeactivation) SetCompanyTeam(v []MemberDeactivationCompanyTeam)`

SetCompanyTeam sets CompanyTeam field to given value.

### HasCompanyTeam

`func (o *MemberDeactivation) HasCompanyTeam() bool`

HasCompanyTeam returns a boolean if a field has been set.

### GetWorkflowEmail

`func (o *MemberDeactivation) GetWorkflowEmail() MemberDeactivationItem`

GetWorkflowEmail returns the WorkflowEmail field if non-nil, zero value otherwise.

### GetWorkflowEmailOk

`func (o *MemberDeactivation) GetWorkflowEmailOk() (*MemberDeactivationItem, bool)`

GetWorkflowEmailOk returns a tuple with the WorkflowEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowEmail

`func (o *MemberDeactivation) SetWorkflowEmail(v MemberDeactivationItem)`

SetWorkflowEmail sets WorkflowEmail field to given value.

### HasWorkflowEmail

`func (o *MemberDeactivation) HasWorkflowEmail() bool`

HasWorkflowEmail returns a boolean if a field has been set.

### GetServiceStatusWorkflow

`func (o *MemberDeactivation) GetServiceStatusWorkflow() []MemberDeactivationStatusWorkflow`

GetServiceStatusWorkflow returns the ServiceStatusWorkflow field if non-nil, zero value otherwise.

### GetServiceStatusWorkflowOk

`func (o *MemberDeactivation) GetServiceStatusWorkflowOk() (*[]MemberDeactivationStatusWorkflow, bool)`

GetServiceStatusWorkflowOk returns a tuple with the ServiceStatusWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatusWorkflow

`func (o *MemberDeactivation) SetServiceStatusWorkflow(v []MemberDeactivationStatusWorkflow)`

SetServiceStatusWorkflow sets ServiceStatusWorkflow field to given value.

### HasServiceStatusWorkflow

`func (o *MemberDeactivation) HasServiceStatusWorkflow() bool`

HasServiceStatusWorkflow returns a boolean if a field has been set.

### GetTicketTemplate

`func (o *MemberDeactivation) GetTicketTemplate() MemberDeactivationItem`

GetTicketTemplate returns the TicketTemplate field if non-nil, zero value otherwise.

### GetTicketTemplateOk

`func (o *MemberDeactivation) GetTicketTemplateOk() (*MemberDeactivationItem, bool)`

GetTicketTemplateOk returns a tuple with the TicketTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketTemplate

`func (o *MemberDeactivation) SetTicketTemplate(v MemberDeactivationItem)`

SetTicketTemplate sets TicketTemplate field to given value.

### HasTicketTemplate

`func (o *MemberDeactivation) HasTicketTemplate() bool`

HasTicketTemplate returns a boolean if a field has been set.

### GetOpportunity

`func (o *MemberDeactivation) GetOpportunity() MemberDeactivationItem`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *MemberDeactivation) GetOpportunityOk() (*MemberDeactivationItem, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *MemberDeactivation) SetOpportunity(v MemberDeactivationItem)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *MemberDeactivation) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### GetSalesTeam

`func (o *MemberDeactivation) GetSalesTeam() MemberDeactivationItem`

GetSalesTeam returns the SalesTeam field if non-nil, zero value otherwise.

### GetSalesTeamOk

`func (o *MemberDeactivation) GetSalesTeamOk() (*MemberDeactivationItem, bool)`

GetSalesTeamOk returns a tuple with the SalesTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTeam

`func (o *MemberDeactivation) SetSalesTeam(v MemberDeactivationItem)`

SetSalesTeam sets SalesTeam field to given value.

### HasSalesTeam

`func (o *MemberDeactivation) HasSalesTeam() bool`

HasSalesTeam returns a boolean if a field has been set.

### GetProjectManager

`func (o *MemberDeactivation) GetProjectManager() MemberDeactivationItem`

GetProjectManager returns the ProjectManager field if non-nil, zero value otherwise.

### GetProjectManagerOk

`func (o *MemberDeactivation) GetProjectManagerOk() (*MemberDeactivationItem, bool)`

GetProjectManagerOk returns a tuple with the ProjectManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectManager

`func (o *MemberDeactivation) SetProjectManager(v MemberDeactivationItem)`

SetProjectManager sets ProjectManager field to given value.

### HasProjectManager

`func (o *MemberDeactivation) HasProjectManager() bool`

HasProjectManager returns a boolean if a field has been set.

### GetProjectTimeApprover

`func (o *MemberDeactivation) GetProjectTimeApprover() MemberDeactivationItem`

GetProjectTimeApprover returns the ProjectTimeApprover field if non-nil, zero value otherwise.

### GetProjectTimeApproverOk

`func (o *MemberDeactivation) GetProjectTimeApproverOk() (*MemberDeactivationItem, bool)`

GetProjectTimeApproverOk returns a tuple with the ProjectTimeApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTimeApprover

`func (o *MemberDeactivation) SetProjectTimeApprover(v MemberDeactivationItem)`

SetProjectTimeApprover sets ProjectTimeApprover field to given value.

### HasProjectTimeApprover

`func (o *MemberDeactivation) HasProjectTimeApprover() bool`

HasProjectTimeApprover returns a boolean if a field has been set.

### GetProjectExpenseApprover

`func (o *MemberDeactivation) GetProjectExpenseApprover() MemberDeactivationItem`

GetProjectExpenseApprover returns the ProjectExpenseApprover field if non-nil, zero value otherwise.

### GetProjectExpenseApproverOk

`func (o *MemberDeactivation) GetProjectExpenseApproverOk() (*MemberDeactivationItem, bool)`

GetProjectExpenseApproverOk returns a tuple with the ProjectExpenseApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectExpenseApprover

`func (o *MemberDeactivation) SetProjectExpenseApprover(v MemberDeactivationItem)`

SetProjectExpenseApprover sets ProjectExpenseApprover field to given value.

### HasProjectExpenseApprover

`func (o *MemberDeactivation) HasProjectExpenseApprover() bool`

HasProjectExpenseApprover returns a boolean if a field has been set.

### GetKnowledgeBaseArticle

`func (o *MemberDeactivation) GetKnowledgeBaseArticle() MemberDeactivationItem`

GetKnowledgeBaseArticle returns the KnowledgeBaseArticle field if non-nil, zero value otherwise.

### GetKnowledgeBaseArticleOk

`func (o *MemberDeactivation) GetKnowledgeBaseArticleOk() (*MemberDeactivationItem, bool)`

GetKnowledgeBaseArticleOk returns a tuple with the KnowledgeBaseArticle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBaseArticle

`func (o *MemberDeactivation) SetKnowledgeBaseArticle(v MemberDeactivationItem)`

SetKnowledgeBaseArticle sets KnowledgeBaseArticle field to given value.

### HasKnowledgeBaseArticle

`func (o *MemberDeactivation) HasKnowledgeBaseArticle() bool`

HasKnowledgeBaseArticle returns a boolean if a field has been set.

### GetMyCompanyPresident

`func (o *MemberDeactivation) GetMyCompanyPresident() MemberDeactivationItem`

GetMyCompanyPresident returns the MyCompanyPresident field if non-nil, zero value otherwise.

### GetMyCompanyPresidentOk

`func (o *MemberDeactivation) GetMyCompanyPresidentOk() (*MemberDeactivationItem, bool)`

GetMyCompanyPresidentOk returns a tuple with the MyCompanyPresident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyCompanyPresident

`func (o *MemberDeactivation) SetMyCompanyPresident(v MemberDeactivationItem)`

SetMyCompanyPresident sets MyCompanyPresident field to given value.

### HasMyCompanyPresident

`func (o *MemberDeactivation) HasMyCompanyPresident() bool`

HasMyCompanyPresident returns a boolean if a field has been set.

### GetMyCompanyCOO

`func (o *MemberDeactivation) GetMyCompanyCOO() MemberDeactivationItem`

GetMyCompanyCOO returns the MyCompanyCOO field if non-nil, zero value otherwise.

### GetMyCompanyCOOOk

`func (o *MemberDeactivation) GetMyCompanyCOOOk() (*MemberDeactivationItem, bool)`

GetMyCompanyCOOOk returns a tuple with the MyCompanyCOO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyCompanyCOO

`func (o *MemberDeactivation) SetMyCompanyCOO(v MemberDeactivationItem)`

SetMyCompanyCOO sets MyCompanyCOO field to given value.

### HasMyCompanyCOO

`func (o *MemberDeactivation) HasMyCompanyCOO() bool`

HasMyCompanyCOO returns a boolean if a field has been set.

### GetMyCompanyController

`func (o *MemberDeactivation) GetMyCompanyController() MemberDeactivationItem`

GetMyCompanyController returns the MyCompanyController field if non-nil, zero value otherwise.

### GetMyCompanyControllerOk

`func (o *MemberDeactivation) GetMyCompanyControllerOk() (*MemberDeactivationItem, bool)`

GetMyCompanyControllerOk returns a tuple with the MyCompanyController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyCompanyController

`func (o *MemberDeactivation) SetMyCompanyController(v MemberDeactivationItem)`

SetMyCompanyController sets MyCompanyController field to given value.

### HasMyCompanyController

`func (o *MemberDeactivation) HasMyCompanyController() bool`

HasMyCompanyController returns a boolean if a field has been set.

### GetMyCompanyDispatch

`func (o *MemberDeactivation) GetMyCompanyDispatch() MemberDeactivationItem`

GetMyCompanyDispatch returns the MyCompanyDispatch field if non-nil, zero value otherwise.

### GetMyCompanyDispatchOk

`func (o *MemberDeactivation) GetMyCompanyDispatchOk() (*MemberDeactivationItem, bool)`

GetMyCompanyDispatchOk returns a tuple with the MyCompanyDispatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyCompanyDispatch

`func (o *MemberDeactivation) SetMyCompanyDispatch(v MemberDeactivationItem)`

SetMyCompanyDispatch sets MyCompanyDispatch field to given value.

### HasMyCompanyDispatch

`func (o *MemberDeactivation) HasMyCompanyDispatch() bool`

HasMyCompanyDispatch returns a boolean if a field has been set.

### GetMyCompanyServiceManager

`func (o *MemberDeactivation) GetMyCompanyServiceManager() MemberDeactivationItem`

GetMyCompanyServiceManager returns the MyCompanyServiceManager field if non-nil, zero value otherwise.

### GetMyCompanyServiceManagerOk

`func (o *MemberDeactivation) GetMyCompanyServiceManagerOk() (*MemberDeactivationItem, bool)`

GetMyCompanyServiceManagerOk returns a tuple with the MyCompanyServiceManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyCompanyServiceManager

`func (o *MemberDeactivation) SetMyCompanyServiceManager(v MemberDeactivationItem)`

SetMyCompanyServiceManager sets MyCompanyServiceManager field to given value.

### HasMyCompanyServiceManager

`func (o *MemberDeactivation) HasMyCompanyServiceManager() bool`

HasMyCompanyServiceManager returns a boolean if a field has been set.

### GetMyCompanyDutyManagerRole

`func (o *MemberDeactivation) GetMyCompanyDutyManagerRole() MemberDeactivationItem`

GetMyCompanyDutyManagerRole returns the MyCompanyDutyManagerRole field if non-nil, zero value otherwise.

### GetMyCompanyDutyManagerRoleOk

`func (o *MemberDeactivation) GetMyCompanyDutyManagerRoleOk() (*MemberDeactivationItem, bool)`

GetMyCompanyDutyManagerRoleOk returns a tuple with the MyCompanyDutyManagerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyCompanyDutyManagerRole

`func (o *MemberDeactivation) SetMyCompanyDutyManagerRole(v MemberDeactivationItem)`

SetMyCompanyDutyManagerRole sets MyCompanyDutyManagerRole field to given value.

### HasMyCompanyDutyManagerRole

`func (o *MemberDeactivation) HasMyCompanyDutyManagerRole() bool`

HasMyCompanyDutyManagerRole returns a boolean if a field has been set.

### GetDepartmentManager

`func (o *MemberDeactivation) GetDepartmentManager() MemberDeactivationItem`

GetDepartmentManager returns the DepartmentManager field if non-nil, zero value otherwise.

### GetDepartmentManagerOk

`func (o *MemberDeactivation) GetDepartmentManagerOk() (*MemberDeactivationItem, bool)`

GetDepartmentManagerOk returns a tuple with the DepartmentManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentManager

`func (o *MemberDeactivation) SetDepartmentManager(v MemberDeactivationItem)`

SetDepartmentManager sets DepartmentManager field to given value.

### HasDepartmentManager

`func (o *MemberDeactivation) HasDepartmentManager() bool`

HasDepartmentManager returns a boolean if a field has been set.

### GetDispatchMember

`func (o *MemberDeactivation) GetDispatchMember() MemberDeactivationItem`

GetDispatchMember returns the DispatchMember field if non-nil, zero value otherwise.

### GetDispatchMemberOk

`func (o *MemberDeactivation) GetDispatchMemberOk() (*MemberDeactivationItem, bool)`

GetDispatchMemberOk returns a tuple with the DispatchMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchMember

`func (o *MemberDeactivation) SetDispatchMember(v MemberDeactivationItem)`

SetDispatchMember sets DispatchMember field to given value.

### HasDispatchMember

`func (o *MemberDeactivation) HasDispatchMember() bool`

HasDispatchMember returns a boolean if a field has been set.

### GetServiceManager

`func (o *MemberDeactivation) GetServiceManager() MemberDeactivationItem`

GetServiceManager returns the ServiceManager field if non-nil, zero value otherwise.

### GetServiceManagerOk

`func (o *MemberDeactivation) GetServiceManagerOk() (*MemberDeactivationItem, bool)`

GetServiceManagerOk returns a tuple with the ServiceManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceManager

`func (o *MemberDeactivation) SetServiceManager(v MemberDeactivationItem)`

SetServiceManager sets ServiceManager field to given value.

### HasServiceManager

`func (o *MemberDeactivation) HasServiceManager() bool`

HasServiceManager returns a boolean if a field has been set.

### GetDutyManager

`func (o *MemberDeactivation) GetDutyManager() MemberDeactivationItem`

GetDutyManager returns the DutyManager field if non-nil, zero value otherwise.

### GetDutyManagerOk

`func (o *MemberDeactivation) GetDutyManagerOk() (*MemberDeactivationItem, bool)`

GetDutyManagerOk returns a tuple with the DutyManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyManager

`func (o *MemberDeactivation) SetDutyManager(v MemberDeactivationItem)`

SetDutyManager sets DutyManager field to given value.

### HasDutyManager

`func (o *MemberDeactivation) HasDutyManager() bool`

HasDutyManager returns a boolean if a field has been set.

### GetSendFromEmailNotify

`func (o *MemberDeactivation) GetSendFromEmailNotify() MemberDeactivationItem`

GetSendFromEmailNotify returns the SendFromEmailNotify field if non-nil, zero value otherwise.

### GetSendFromEmailNotifyOk

`func (o *MemberDeactivation) GetSendFromEmailNotifyOk() (*MemberDeactivationItem, bool)`

GetSendFromEmailNotifyOk returns a tuple with the SendFromEmailNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendFromEmailNotify

`func (o *MemberDeactivation) SetSendFromEmailNotify(v MemberDeactivationItem)`

SetSendFromEmailNotify sets SendFromEmailNotify field to given value.

### HasSendFromEmailNotify

`func (o *MemberDeactivation) HasSendFromEmailNotify() bool`

HasSendFromEmailNotify returns a boolean if a field has been set.

### GetDeleteOpenTimeSheetsFlag

`func (o *MemberDeactivation) GetDeleteOpenTimeSheetsFlag() bool`

GetDeleteOpenTimeSheetsFlag returns the DeleteOpenTimeSheetsFlag field if non-nil, zero value otherwise.

### GetDeleteOpenTimeSheetsFlagOk

`func (o *MemberDeactivation) GetDeleteOpenTimeSheetsFlagOk() (*bool, bool)`

GetDeleteOpenTimeSheetsFlagOk returns a tuple with the DeleteOpenTimeSheetsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOpenTimeSheetsFlag

`func (o *MemberDeactivation) SetDeleteOpenTimeSheetsFlag(v bool)`

SetDeleteOpenTimeSheetsFlag sets DeleteOpenTimeSheetsFlag field to given value.

### HasDeleteOpenTimeSheetsFlag

`func (o *MemberDeactivation) HasDeleteOpenTimeSheetsFlag() bool`

HasDeleteOpenTimeSheetsFlag returns a boolean if a field has been set.

### SetDeleteOpenTimeSheetsFlagNil

`func (o *MemberDeactivation) SetDeleteOpenTimeSheetsFlagNil(b bool)`

 SetDeleteOpenTimeSheetsFlagNil sets the value for DeleteOpenTimeSheetsFlag to be an explicit nil

### UnsetDeleteOpenTimeSheetsFlag
`func (o *MemberDeactivation) UnsetDeleteOpenTimeSheetsFlag()`

UnsetDeleteOpenTimeSheetsFlag ensures that no value is present for DeleteOpenTimeSheetsFlag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


