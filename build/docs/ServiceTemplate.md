# ServiceTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Type** | Pointer to [**ServiceTypeReference**](ServiceTypeReference.md) |  | [optional] 
**Item** | Pointer to [**ServiceItemReference**](ServiceItemReference.md) |  | [optional] 
**Subtype** | Pointer to [**ServiceSubTypeReference**](ServiceSubTypeReference.md) |  | [optional] 
**ServiceLocation** | Pointer to [**ServiceLocationReference**](ServiceLocationReference.md) |  | [optional] 
**Status** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**Source** | Pointer to [**ServiceSourceReference**](ServiceSourceReference.md) |  | [optional] 
**Priority** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**Team** | Pointer to [**ServiceTeamReference**](ServiceTeamReference.md) |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**AssignedNotifyFlag** | Pointer to **NullableBool** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Problem** | Pointer to **string** |  | [optional] 
**HoursBudget** | Pointer to **NullableFloat64** |  | [optional] 
**InternalAnalysis** | Pointer to **string** |  | [optional] 
**TimeBillableFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseBillableFlag** | Pointer to **NullableBool** |  | [optional] 
**PurchaseOrderNumber** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**BillCompleteFlag** | Pointer to **NullableBool** |  | [optional] 
**BillServiceSeparatelyFlag** | Pointer to **NullableBool** |  | [optional] 
**BillingAmount** | Pointer to **NullableFloat64** |  | [optional] 
**BillUnapprovedTimeAndExpensesFlag** | Pointer to **NullableBool** |  | [optional] 
**OverrideFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**ProductInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**BillingMethod** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**AssignedBy** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**ScheduleDaysBefore** | Pointer to **NullableInt32** |  | [optional] 
**ServiceDaysBefore** | Pointer to **NullableInt32** |  | [optional] 
**AttachScheduleToNewServiceFlag** | Pointer to **NullableBool** |  | [optional] 
**TemplateFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailContactFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailResourceFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailCCFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailCC** | Pointer to **string** |  | [optional] 
**RestrictDownpaymentFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceTemplate

`func NewServiceTemplate() *ServiceTemplate`

NewServiceTemplate instantiates a new ServiceTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTemplateWithDefaults

`func NewServiceTemplateWithDefaults() *ServiceTemplate`

NewServiceTemplateWithDefaults instantiates a new ServiceTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBoard

`func (o *ServiceTemplate) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *ServiceTemplate) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *ServiceTemplate) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *ServiceTemplate) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetType

`func (o *ServiceTemplate) GetType() ServiceTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceTemplate) GetTypeOk() (*ServiceTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceTemplate) SetType(v ServiceTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItem

`func (o *ServiceTemplate) GetItem() ServiceItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ServiceTemplate) GetItemOk() (*ServiceItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ServiceTemplate) SetItem(v ServiceItemReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *ServiceTemplate) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetSubtype

`func (o *ServiceTemplate) GetSubtype() ServiceSubTypeReference`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ServiceTemplate) GetSubtypeOk() (*ServiceSubTypeReference, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ServiceTemplate) SetSubtype(v ServiceSubTypeReference)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *ServiceTemplate) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetServiceLocation

`func (o *ServiceTemplate) GetServiceLocation() ServiceLocationReference`

GetServiceLocation returns the ServiceLocation field if non-nil, zero value otherwise.

### GetServiceLocationOk

`func (o *ServiceTemplate) GetServiceLocationOk() (*ServiceLocationReference, bool)`

GetServiceLocationOk returns a tuple with the ServiceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLocation

`func (o *ServiceTemplate) SetServiceLocation(v ServiceLocationReference)`

SetServiceLocation sets ServiceLocation field to given value.

### HasServiceLocation

`func (o *ServiceTemplate) HasServiceLocation() bool`

HasServiceLocation returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceTemplate) GetStatus() ServiceStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceTemplate) GetStatusOk() (*ServiceStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceTemplate) SetStatus(v ServiceStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSource

`func (o *ServiceTemplate) GetSource() ServiceSourceReference`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ServiceTemplate) GetSourceOk() (*ServiceSourceReference, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ServiceTemplate) SetSource(v ServiceSourceReference)`

SetSource sets Source field to given value.

### HasSource

`func (o *ServiceTemplate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPriority

`func (o *ServiceTemplate) GetPriority() PriorityReference`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ServiceTemplate) GetPriorityOk() (*PriorityReference, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ServiceTemplate) SetPriority(v PriorityReference)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ServiceTemplate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTeam

`func (o *ServiceTemplate) GetTeam() ServiceTeamReference`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ServiceTemplate) GetTeamOk() (*ServiceTeamReference, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ServiceTemplate) SetTeam(v ServiceTeamReference)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ServiceTemplate) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetCompany

`func (o *ServiceTemplate) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ServiceTemplate) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ServiceTemplate) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ServiceTemplate) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetContact

`func (o *ServiceTemplate) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ServiceTemplate) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ServiceTemplate) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ServiceTemplate) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetSite

`func (o *ServiceTemplate) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ServiceTemplate) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ServiceTemplate) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *ServiceTemplate) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAssignedNotifyFlag

`func (o *ServiceTemplate) GetAssignedNotifyFlag() bool`

GetAssignedNotifyFlag returns the AssignedNotifyFlag field if non-nil, zero value otherwise.

### GetAssignedNotifyFlagOk

`func (o *ServiceTemplate) GetAssignedNotifyFlagOk() (*bool, bool)`

GetAssignedNotifyFlagOk returns a tuple with the AssignedNotifyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedNotifyFlag

`func (o *ServiceTemplate) SetAssignedNotifyFlag(v bool)`

SetAssignedNotifyFlag sets AssignedNotifyFlag field to given value.

### HasAssignedNotifyFlag

`func (o *ServiceTemplate) HasAssignedNotifyFlag() bool`

HasAssignedNotifyFlag returns a boolean if a field has been set.

### SetAssignedNotifyFlagNil

`func (o *ServiceTemplate) SetAssignedNotifyFlagNil(b bool)`

 SetAssignedNotifyFlagNil sets the value for AssignedNotifyFlag to be an explicit nil

### UnsetAssignedNotifyFlag
`func (o *ServiceTemplate) UnsetAssignedNotifyFlag()`

UnsetAssignedNotifyFlag ensures that no value is present for AssignedNotifyFlag, not even an explicit nil
### GetLocation

`func (o *ServiceTemplate) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ServiceTemplate) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ServiceTemplate) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ServiceTemplate) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *ServiceTemplate) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ServiceTemplate) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ServiceTemplate) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ServiceTemplate) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetSummary

`func (o *ServiceTemplate) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ServiceTemplate) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ServiceTemplate) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ServiceTemplate) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetProblem

`func (o *ServiceTemplate) GetProblem() string`

GetProblem returns the Problem field if non-nil, zero value otherwise.

### GetProblemOk

`func (o *ServiceTemplate) GetProblemOk() (*string, bool)`

GetProblemOk returns a tuple with the Problem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblem

`func (o *ServiceTemplate) SetProblem(v string)`

SetProblem sets Problem field to given value.

### HasProblem

`func (o *ServiceTemplate) HasProblem() bool`

HasProblem returns a boolean if a field has been set.

### GetHoursBudget

`func (o *ServiceTemplate) GetHoursBudget() float64`

GetHoursBudget returns the HoursBudget field if non-nil, zero value otherwise.

### GetHoursBudgetOk

`func (o *ServiceTemplate) GetHoursBudgetOk() (*float64, bool)`

GetHoursBudgetOk returns a tuple with the HoursBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursBudget

`func (o *ServiceTemplate) SetHoursBudget(v float64)`

SetHoursBudget sets HoursBudget field to given value.

### HasHoursBudget

`func (o *ServiceTemplate) HasHoursBudget() bool`

HasHoursBudget returns a boolean if a field has been set.

### SetHoursBudgetNil

`func (o *ServiceTemplate) SetHoursBudgetNil(b bool)`

 SetHoursBudgetNil sets the value for HoursBudget to be an explicit nil

### UnsetHoursBudget
`func (o *ServiceTemplate) UnsetHoursBudget()`

UnsetHoursBudget ensures that no value is present for HoursBudget, not even an explicit nil
### GetInternalAnalysis

`func (o *ServiceTemplate) GetInternalAnalysis() string`

GetInternalAnalysis returns the InternalAnalysis field if non-nil, zero value otherwise.

### GetInternalAnalysisOk

`func (o *ServiceTemplate) GetInternalAnalysisOk() (*string, bool)`

GetInternalAnalysisOk returns a tuple with the InternalAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAnalysis

`func (o *ServiceTemplate) SetInternalAnalysis(v string)`

SetInternalAnalysis sets InternalAnalysis field to given value.

### HasInternalAnalysis

`func (o *ServiceTemplate) HasInternalAnalysis() bool`

HasInternalAnalysis returns a boolean if a field has been set.

### GetTimeBillableFlag

`func (o *ServiceTemplate) GetTimeBillableFlag() bool`

GetTimeBillableFlag returns the TimeBillableFlag field if non-nil, zero value otherwise.

### GetTimeBillableFlagOk

`func (o *ServiceTemplate) GetTimeBillableFlagOk() (*bool, bool)`

GetTimeBillableFlagOk returns a tuple with the TimeBillableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBillableFlag

`func (o *ServiceTemplate) SetTimeBillableFlag(v bool)`

SetTimeBillableFlag sets TimeBillableFlag field to given value.

### HasTimeBillableFlag

`func (o *ServiceTemplate) HasTimeBillableFlag() bool`

HasTimeBillableFlag returns a boolean if a field has been set.

### SetTimeBillableFlagNil

`func (o *ServiceTemplate) SetTimeBillableFlagNil(b bool)`

 SetTimeBillableFlagNil sets the value for TimeBillableFlag to be an explicit nil

### UnsetTimeBillableFlag
`func (o *ServiceTemplate) UnsetTimeBillableFlag()`

UnsetTimeBillableFlag ensures that no value is present for TimeBillableFlag, not even an explicit nil
### GetExpenseBillableFlag

`func (o *ServiceTemplate) GetExpenseBillableFlag() bool`

GetExpenseBillableFlag returns the ExpenseBillableFlag field if non-nil, zero value otherwise.

### GetExpenseBillableFlagOk

`func (o *ServiceTemplate) GetExpenseBillableFlagOk() (*bool, bool)`

GetExpenseBillableFlagOk returns a tuple with the ExpenseBillableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseBillableFlag

`func (o *ServiceTemplate) SetExpenseBillableFlag(v bool)`

SetExpenseBillableFlag sets ExpenseBillableFlag field to given value.

### HasExpenseBillableFlag

`func (o *ServiceTemplate) HasExpenseBillableFlag() bool`

HasExpenseBillableFlag returns a boolean if a field has been set.

### SetExpenseBillableFlagNil

`func (o *ServiceTemplate) SetExpenseBillableFlagNil(b bool)`

 SetExpenseBillableFlagNil sets the value for ExpenseBillableFlag to be an explicit nil

### UnsetExpenseBillableFlag
`func (o *ServiceTemplate) UnsetExpenseBillableFlag()`

UnsetExpenseBillableFlag ensures that no value is present for ExpenseBillableFlag, not even an explicit nil
### GetPurchaseOrderNumber

`func (o *ServiceTemplate) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *ServiceTemplate) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *ServiceTemplate) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *ServiceTemplate) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetReference

`func (o *ServiceTemplate) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ServiceTemplate) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ServiceTemplate) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ServiceTemplate) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetBillCompleteFlag

`func (o *ServiceTemplate) GetBillCompleteFlag() bool`

GetBillCompleteFlag returns the BillCompleteFlag field if non-nil, zero value otherwise.

### GetBillCompleteFlagOk

`func (o *ServiceTemplate) GetBillCompleteFlagOk() (*bool, bool)`

GetBillCompleteFlagOk returns a tuple with the BillCompleteFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCompleteFlag

`func (o *ServiceTemplate) SetBillCompleteFlag(v bool)`

SetBillCompleteFlag sets BillCompleteFlag field to given value.

### HasBillCompleteFlag

`func (o *ServiceTemplate) HasBillCompleteFlag() bool`

HasBillCompleteFlag returns a boolean if a field has been set.

### SetBillCompleteFlagNil

`func (o *ServiceTemplate) SetBillCompleteFlagNil(b bool)`

 SetBillCompleteFlagNil sets the value for BillCompleteFlag to be an explicit nil

### UnsetBillCompleteFlag
`func (o *ServiceTemplate) UnsetBillCompleteFlag()`

UnsetBillCompleteFlag ensures that no value is present for BillCompleteFlag, not even an explicit nil
### GetBillServiceSeparatelyFlag

`func (o *ServiceTemplate) GetBillServiceSeparatelyFlag() bool`

GetBillServiceSeparatelyFlag returns the BillServiceSeparatelyFlag field if non-nil, zero value otherwise.

### GetBillServiceSeparatelyFlagOk

`func (o *ServiceTemplate) GetBillServiceSeparatelyFlagOk() (*bool, bool)`

GetBillServiceSeparatelyFlagOk returns a tuple with the BillServiceSeparatelyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillServiceSeparatelyFlag

`func (o *ServiceTemplate) SetBillServiceSeparatelyFlag(v bool)`

SetBillServiceSeparatelyFlag sets BillServiceSeparatelyFlag field to given value.

### HasBillServiceSeparatelyFlag

`func (o *ServiceTemplate) HasBillServiceSeparatelyFlag() bool`

HasBillServiceSeparatelyFlag returns a boolean if a field has been set.

### SetBillServiceSeparatelyFlagNil

`func (o *ServiceTemplate) SetBillServiceSeparatelyFlagNil(b bool)`

 SetBillServiceSeparatelyFlagNil sets the value for BillServiceSeparatelyFlag to be an explicit nil

### UnsetBillServiceSeparatelyFlag
`func (o *ServiceTemplate) UnsetBillServiceSeparatelyFlag()`

UnsetBillServiceSeparatelyFlag ensures that no value is present for BillServiceSeparatelyFlag, not even an explicit nil
### GetBillingAmount

`func (o *ServiceTemplate) GetBillingAmount() float64`

GetBillingAmount returns the BillingAmount field if non-nil, zero value otherwise.

### GetBillingAmountOk

`func (o *ServiceTemplate) GetBillingAmountOk() (*float64, bool)`

GetBillingAmountOk returns a tuple with the BillingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAmount

`func (o *ServiceTemplate) SetBillingAmount(v float64)`

SetBillingAmount sets BillingAmount field to given value.

### HasBillingAmount

`func (o *ServiceTemplate) HasBillingAmount() bool`

HasBillingAmount returns a boolean if a field has been set.

### SetBillingAmountNil

`func (o *ServiceTemplate) SetBillingAmountNil(b bool)`

 SetBillingAmountNil sets the value for BillingAmount to be an explicit nil

### UnsetBillingAmount
`func (o *ServiceTemplate) UnsetBillingAmount()`

UnsetBillingAmount ensures that no value is present for BillingAmount, not even an explicit nil
### GetBillUnapprovedTimeAndExpensesFlag

`func (o *ServiceTemplate) GetBillUnapprovedTimeAndExpensesFlag() bool`

GetBillUnapprovedTimeAndExpensesFlag returns the BillUnapprovedTimeAndExpensesFlag field if non-nil, zero value otherwise.

### GetBillUnapprovedTimeAndExpensesFlagOk

`func (o *ServiceTemplate) GetBillUnapprovedTimeAndExpensesFlagOk() (*bool, bool)`

GetBillUnapprovedTimeAndExpensesFlagOk returns a tuple with the BillUnapprovedTimeAndExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillUnapprovedTimeAndExpensesFlag

`func (o *ServiceTemplate) SetBillUnapprovedTimeAndExpensesFlag(v bool)`

SetBillUnapprovedTimeAndExpensesFlag sets BillUnapprovedTimeAndExpensesFlag field to given value.

### HasBillUnapprovedTimeAndExpensesFlag

`func (o *ServiceTemplate) HasBillUnapprovedTimeAndExpensesFlag() bool`

HasBillUnapprovedTimeAndExpensesFlag returns a boolean if a field has been set.

### SetBillUnapprovedTimeAndExpensesFlagNil

`func (o *ServiceTemplate) SetBillUnapprovedTimeAndExpensesFlagNil(b bool)`

 SetBillUnapprovedTimeAndExpensesFlagNil sets the value for BillUnapprovedTimeAndExpensesFlag to be an explicit nil

### UnsetBillUnapprovedTimeAndExpensesFlag
`func (o *ServiceTemplate) UnsetBillUnapprovedTimeAndExpensesFlag()`

UnsetBillUnapprovedTimeAndExpensesFlag ensures that no value is present for BillUnapprovedTimeAndExpensesFlag, not even an explicit nil
### GetOverrideFlag

`func (o *ServiceTemplate) GetOverrideFlag() bool`

GetOverrideFlag returns the OverrideFlag field if non-nil, zero value otherwise.

### GetOverrideFlagOk

`func (o *ServiceTemplate) GetOverrideFlagOk() (*bool, bool)`

GetOverrideFlagOk returns a tuple with the OverrideFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideFlag

`func (o *ServiceTemplate) SetOverrideFlag(v bool)`

SetOverrideFlag sets OverrideFlag field to given value.

### HasOverrideFlag

`func (o *ServiceTemplate) HasOverrideFlag() bool`

HasOverrideFlag returns a boolean if a field has been set.

### SetOverrideFlagNil

`func (o *ServiceTemplate) SetOverrideFlagNil(b bool)`

 SetOverrideFlagNil sets the value for OverrideFlag to be an explicit nil

### UnsetOverrideFlag
`func (o *ServiceTemplate) UnsetOverrideFlag()`

UnsetOverrideFlag ensures that no value is present for OverrideFlag, not even an explicit nil
### GetTimeInvoiceFlag

`func (o *ServiceTemplate) GetTimeInvoiceFlag() bool`

GetTimeInvoiceFlag returns the TimeInvoiceFlag field if non-nil, zero value otherwise.

### GetTimeInvoiceFlagOk

`func (o *ServiceTemplate) GetTimeInvoiceFlagOk() (*bool, bool)`

GetTimeInvoiceFlagOk returns a tuple with the TimeInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInvoiceFlag

`func (o *ServiceTemplate) SetTimeInvoiceFlag(v bool)`

SetTimeInvoiceFlag sets TimeInvoiceFlag field to given value.

### HasTimeInvoiceFlag

`func (o *ServiceTemplate) HasTimeInvoiceFlag() bool`

HasTimeInvoiceFlag returns a boolean if a field has been set.

### SetTimeInvoiceFlagNil

`func (o *ServiceTemplate) SetTimeInvoiceFlagNil(b bool)`

 SetTimeInvoiceFlagNil sets the value for TimeInvoiceFlag to be an explicit nil

### UnsetTimeInvoiceFlag
`func (o *ServiceTemplate) UnsetTimeInvoiceFlag()`

UnsetTimeInvoiceFlag ensures that no value is present for TimeInvoiceFlag, not even an explicit nil
### GetExpenseInvoiceFlag

`func (o *ServiceTemplate) GetExpenseInvoiceFlag() bool`

GetExpenseInvoiceFlag returns the ExpenseInvoiceFlag field if non-nil, zero value otherwise.

### GetExpenseInvoiceFlagOk

`func (o *ServiceTemplate) GetExpenseInvoiceFlagOk() (*bool, bool)`

GetExpenseInvoiceFlagOk returns a tuple with the ExpenseInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseInvoiceFlag

`func (o *ServiceTemplate) SetExpenseInvoiceFlag(v bool)`

SetExpenseInvoiceFlag sets ExpenseInvoiceFlag field to given value.

### HasExpenseInvoiceFlag

`func (o *ServiceTemplate) HasExpenseInvoiceFlag() bool`

HasExpenseInvoiceFlag returns a boolean if a field has been set.

### SetExpenseInvoiceFlagNil

`func (o *ServiceTemplate) SetExpenseInvoiceFlagNil(b bool)`

 SetExpenseInvoiceFlagNil sets the value for ExpenseInvoiceFlag to be an explicit nil

### UnsetExpenseInvoiceFlag
`func (o *ServiceTemplate) UnsetExpenseInvoiceFlag()`

UnsetExpenseInvoiceFlag ensures that no value is present for ExpenseInvoiceFlag, not even an explicit nil
### GetProductInvoiceFlag

`func (o *ServiceTemplate) GetProductInvoiceFlag() bool`

GetProductInvoiceFlag returns the ProductInvoiceFlag field if non-nil, zero value otherwise.

### GetProductInvoiceFlagOk

`func (o *ServiceTemplate) GetProductInvoiceFlagOk() (*bool, bool)`

GetProductInvoiceFlagOk returns a tuple with the ProductInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInvoiceFlag

`func (o *ServiceTemplate) SetProductInvoiceFlag(v bool)`

SetProductInvoiceFlag sets ProductInvoiceFlag field to given value.

### HasProductInvoiceFlag

`func (o *ServiceTemplate) HasProductInvoiceFlag() bool`

HasProductInvoiceFlag returns a boolean if a field has been set.

### SetProductInvoiceFlagNil

`func (o *ServiceTemplate) SetProductInvoiceFlagNil(b bool)`

 SetProductInvoiceFlagNil sets the value for ProductInvoiceFlag to be an explicit nil

### UnsetProductInvoiceFlag
`func (o *ServiceTemplate) UnsetProductInvoiceFlag()`

UnsetProductInvoiceFlag ensures that no value is present for ProductInvoiceFlag, not even an explicit nil
### GetAgreement

`func (o *ServiceTemplate) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *ServiceTemplate) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *ServiceTemplate) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *ServiceTemplate) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetBillingMethod

`func (o *ServiceTemplate) GetBillingMethod() string`

GetBillingMethod returns the BillingMethod field if non-nil, zero value otherwise.

### GetBillingMethodOk

`func (o *ServiceTemplate) GetBillingMethodOk() (*string, bool)`

GetBillingMethodOk returns a tuple with the BillingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMethod

`func (o *ServiceTemplate) SetBillingMethod(v string)`

SetBillingMethod sets BillingMethod field to given value.

### HasBillingMethod

`func (o *ServiceTemplate) HasBillingMethod() bool`

HasBillingMethod returns a boolean if a field has been set.

### GetSeverity

`func (o *ServiceTemplate) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ServiceTemplate) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ServiceTemplate) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ServiceTemplate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetImpact

`func (o *ServiceTemplate) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *ServiceTemplate) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *ServiceTemplate) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *ServiceTemplate) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetAssignedBy

`func (o *ServiceTemplate) GetAssignedBy() MemberReference`

GetAssignedBy returns the AssignedBy field if non-nil, zero value otherwise.

### GetAssignedByOk

`func (o *ServiceTemplate) GetAssignedByOk() (*MemberReference, bool)`

GetAssignedByOk returns a tuple with the AssignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBy

`func (o *ServiceTemplate) SetAssignedBy(v MemberReference)`

SetAssignedBy sets AssignedBy field to given value.

### HasAssignedBy

`func (o *ServiceTemplate) HasAssignedBy() bool`

HasAssignedBy returns a boolean if a field has been set.

### GetScheduleDaysBefore

`func (o *ServiceTemplate) GetScheduleDaysBefore() int32`

GetScheduleDaysBefore returns the ScheduleDaysBefore field if non-nil, zero value otherwise.

### GetScheduleDaysBeforeOk

`func (o *ServiceTemplate) GetScheduleDaysBeforeOk() (*int32, bool)`

GetScheduleDaysBeforeOk returns a tuple with the ScheduleDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDaysBefore

`func (o *ServiceTemplate) SetScheduleDaysBefore(v int32)`

SetScheduleDaysBefore sets ScheduleDaysBefore field to given value.

### HasScheduleDaysBefore

`func (o *ServiceTemplate) HasScheduleDaysBefore() bool`

HasScheduleDaysBefore returns a boolean if a field has been set.

### SetScheduleDaysBeforeNil

`func (o *ServiceTemplate) SetScheduleDaysBeforeNil(b bool)`

 SetScheduleDaysBeforeNil sets the value for ScheduleDaysBefore to be an explicit nil

### UnsetScheduleDaysBefore
`func (o *ServiceTemplate) UnsetScheduleDaysBefore()`

UnsetScheduleDaysBefore ensures that no value is present for ScheduleDaysBefore, not even an explicit nil
### GetServiceDaysBefore

`func (o *ServiceTemplate) GetServiceDaysBefore() int32`

GetServiceDaysBefore returns the ServiceDaysBefore field if non-nil, zero value otherwise.

### GetServiceDaysBeforeOk

`func (o *ServiceTemplate) GetServiceDaysBeforeOk() (*int32, bool)`

GetServiceDaysBeforeOk returns a tuple with the ServiceDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDaysBefore

`func (o *ServiceTemplate) SetServiceDaysBefore(v int32)`

SetServiceDaysBefore sets ServiceDaysBefore field to given value.

### HasServiceDaysBefore

`func (o *ServiceTemplate) HasServiceDaysBefore() bool`

HasServiceDaysBefore returns a boolean if a field has been set.

### SetServiceDaysBeforeNil

`func (o *ServiceTemplate) SetServiceDaysBeforeNil(b bool)`

 SetServiceDaysBeforeNil sets the value for ServiceDaysBefore to be an explicit nil

### UnsetServiceDaysBefore
`func (o *ServiceTemplate) UnsetServiceDaysBefore()`

UnsetServiceDaysBefore ensures that no value is present for ServiceDaysBefore, not even an explicit nil
### GetAttachScheduleToNewServiceFlag

`func (o *ServiceTemplate) GetAttachScheduleToNewServiceFlag() bool`

GetAttachScheduleToNewServiceFlag returns the AttachScheduleToNewServiceFlag field if non-nil, zero value otherwise.

### GetAttachScheduleToNewServiceFlagOk

`func (o *ServiceTemplate) GetAttachScheduleToNewServiceFlagOk() (*bool, bool)`

GetAttachScheduleToNewServiceFlagOk returns a tuple with the AttachScheduleToNewServiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachScheduleToNewServiceFlag

`func (o *ServiceTemplate) SetAttachScheduleToNewServiceFlag(v bool)`

SetAttachScheduleToNewServiceFlag sets AttachScheduleToNewServiceFlag field to given value.

### HasAttachScheduleToNewServiceFlag

`func (o *ServiceTemplate) HasAttachScheduleToNewServiceFlag() bool`

HasAttachScheduleToNewServiceFlag returns a boolean if a field has been set.

### SetAttachScheduleToNewServiceFlagNil

`func (o *ServiceTemplate) SetAttachScheduleToNewServiceFlagNil(b bool)`

 SetAttachScheduleToNewServiceFlagNil sets the value for AttachScheduleToNewServiceFlag to be an explicit nil

### UnsetAttachScheduleToNewServiceFlag
`func (o *ServiceTemplate) UnsetAttachScheduleToNewServiceFlag()`

UnsetAttachScheduleToNewServiceFlag ensures that no value is present for AttachScheduleToNewServiceFlag, not even an explicit nil
### GetTemplateFlag

`func (o *ServiceTemplate) GetTemplateFlag() bool`

GetTemplateFlag returns the TemplateFlag field if non-nil, zero value otherwise.

### GetTemplateFlagOk

`func (o *ServiceTemplate) GetTemplateFlagOk() (*bool, bool)`

GetTemplateFlagOk returns a tuple with the TemplateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateFlag

`func (o *ServiceTemplate) SetTemplateFlag(v bool)`

SetTemplateFlag sets TemplateFlag field to given value.

### HasTemplateFlag

`func (o *ServiceTemplate) HasTemplateFlag() bool`

HasTemplateFlag returns a boolean if a field has been set.

### SetTemplateFlagNil

`func (o *ServiceTemplate) SetTemplateFlagNil(b bool)`

 SetTemplateFlagNil sets the value for TemplateFlag to be an explicit nil

### UnsetTemplateFlag
`func (o *ServiceTemplate) UnsetTemplateFlag()`

UnsetTemplateFlag ensures that no value is present for TemplateFlag, not even an explicit nil
### GetEmailContactFlag

`func (o *ServiceTemplate) GetEmailContactFlag() bool`

GetEmailContactFlag returns the EmailContactFlag field if non-nil, zero value otherwise.

### GetEmailContactFlagOk

`func (o *ServiceTemplate) GetEmailContactFlagOk() (*bool, bool)`

GetEmailContactFlagOk returns a tuple with the EmailContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailContactFlag

`func (o *ServiceTemplate) SetEmailContactFlag(v bool)`

SetEmailContactFlag sets EmailContactFlag field to given value.

### HasEmailContactFlag

`func (o *ServiceTemplate) HasEmailContactFlag() bool`

HasEmailContactFlag returns a boolean if a field has been set.

### SetEmailContactFlagNil

`func (o *ServiceTemplate) SetEmailContactFlagNil(b bool)`

 SetEmailContactFlagNil sets the value for EmailContactFlag to be an explicit nil

### UnsetEmailContactFlag
`func (o *ServiceTemplate) UnsetEmailContactFlag()`

UnsetEmailContactFlag ensures that no value is present for EmailContactFlag, not even an explicit nil
### GetEmailResourceFlag

`func (o *ServiceTemplate) GetEmailResourceFlag() bool`

GetEmailResourceFlag returns the EmailResourceFlag field if non-nil, zero value otherwise.

### GetEmailResourceFlagOk

`func (o *ServiceTemplate) GetEmailResourceFlagOk() (*bool, bool)`

GetEmailResourceFlagOk returns a tuple with the EmailResourceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailResourceFlag

`func (o *ServiceTemplate) SetEmailResourceFlag(v bool)`

SetEmailResourceFlag sets EmailResourceFlag field to given value.

### HasEmailResourceFlag

`func (o *ServiceTemplate) HasEmailResourceFlag() bool`

HasEmailResourceFlag returns a boolean if a field has been set.

### SetEmailResourceFlagNil

`func (o *ServiceTemplate) SetEmailResourceFlagNil(b bool)`

 SetEmailResourceFlagNil sets the value for EmailResourceFlag to be an explicit nil

### UnsetEmailResourceFlag
`func (o *ServiceTemplate) UnsetEmailResourceFlag()`

UnsetEmailResourceFlag ensures that no value is present for EmailResourceFlag, not even an explicit nil
### GetEmailCCFlag

`func (o *ServiceTemplate) GetEmailCCFlag() bool`

GetEmailCCFlag returns the EmailCCFlag field if non-nil, zero value otherwise.

### GetEmailCCFlagOk

`func (o *ServiceTemplate) GetEmailCCFlagOk() (*bool, bool)`

GetEmailCCFlagOk returns a tuple with the EmailCCFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCCFlag

`func (o *ServiceTemplate) SetEmailCCFlag(v bool)`

SetEmailCCFlag sets EmailCCFlag field to given value.

### HasEmailCCFlag

`func (o *ServiceTemplate) HasEmailCCFlag() bool`

HasEmailCCFlag returns a boolean if a field has been set.

### SetEmailCCFlagNil

`func (o *ServiceTemplate) SetEmailCCFlagNil(b bool)`

 SetEmailCCFlagNil sets the value for EmailCCFlag to be an explicit nil

### UnsetEmailCCFlag
`func (o *ServiceTemplate) UnsetEmailCCFlag()`

UnsetEmailCCFlag ensures that no value is present for EmailCCFlag, not even an explicit nil
### GetEmailCC

`func (o *ServiceTemplate) GetEmailCC() string`

GetEmailCC returns the EmailCC field if non-nil, zero value otherwise.

### GetEmailCCOk

`func (o *ServiceTemplate) GetEmailCCOk() (*string, bool)`

GetEmailCCOk returns a tuple with the EmailCC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCC

`func (o *ServiceTemplate) SetEmailCC(v string)`

SetEmailCC sets EmailCC field to given value.

### HasEmailCC

`func (o *ServiceTemplate) HasEmailCC() bool`

HasEmailCC returns a boolean if a field has been set.

### GetRestrictDownpaymentFlag

`func (o *ServiceTemplate) GetRestrictDownpaymentFlag() bool`

GetRestrictDownpaymentFlag returns the RestrictDownpaymentFlag field if non-nil, zero value otherwise.

### GetRestrictDownpaymentFlagOk

`func (o *ServiceTemplate) GetRestrictDownpaymentFlagOk() (*bool, bool)`

GetRestrictDownpaymentFlagOk returns a tuple with the RestrictDownpaymentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDownpaymentFlag

`func (o *ServiceTemplate) SetRestrictDownpaymentFlag(v bool)`

SetRestrictDownpaymentFlag sets RestrictDownpaymentFlag field to given value.

### HasRestrictDownpaymentFlag

`func (o *ServiceTemplate) HasRestrictDownpaymentFlag() bool`

HasRestrictDownpaymentFlag returns a boolean if a field has been set.

### SetRestrictDownpaymentFlagNil

`func (o *ServiceTemplate) SetRestrictDownpaymentFlagNil(b bool)`

 SetRestrictDownpaymentFlagNil sets the value for RestrictDownpaymentFlag to be an explicit nil

### UnsetRestrictDownpaymentFlag
`func (o *ServiceTemplate) UnsetRestrictDownpaymentFlag()`

UnsetRestrictDownpaymentFlag ensures that no value is present for RestrictDownpaymentFlag, not even an explicit nil
### GetInfo

`func (o *ServiceTemplate) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceTemplate) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceTemplate) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceTemplate) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


