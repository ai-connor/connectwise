# ProjectPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**Description** | **string** |  Max length: 100; | 
**Board** | Pointer to [**ProjectBoardReference**](ProjectBoardReference.md) |  | [optional] 
**Status** | Pointer to [**PhaseStatusReference**](PhaseStatusReference.md) |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**Opportunity** | Pointer to [**OpportunityReference**](OpportunityReference.md) |  | [optional] 
**Department** | Pointer to [**BillingUnitReference**](BillingUnitReference.md) |  | [optional] 
**ParentPhase** | Pointer to [**ProjectPhaseReference**](ProjectPhaseReference.md) |  | [optional] 
**WbsCode** | Pointer to **string** |  Max length: 50; | [optional] 
**BillTime** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**BillExpenses** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**BillProducts** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**MarkAsMilestoneFlag** | Pointer to **NullableBool** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**DeadlineDate** | Pointer to **time.Time** |  | [optional] 
**BillSeparatelyFlag** | Pointer to **NullableBool** |  | [optional] 
**BillingMethod** | Pointer to **NullableString** | billingMethod is required if the phase billSeparatelyFlag is true. | [optional] 
**ScheduledHours** | Pointer to **NullableFloat64** |  | [optional] 
**ScheduledStart** | Pointer to **string** |  | [optional] 
**ScheduledEnd** | Pointer to **string** |  | [optional] 
**ActualHours** | Pointer to **NullableFloat64** |  | [optional] 
**ActualStart** | Pointer to **string** |  | [optional] 
**ActualEnd** | Pointer to **string** |  | [optional] 
**BudgetHours** | Pointer to **NullableFloat64** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  | [optional] 
**HourlyRate** | Pointer to **NullableFloat64** |  | [optional] 
**BillingStartDate** | Pointer to **time.Time** |  | [optional] 
**BillPhaseClosedFlag** | Pointer to **NullableBool** | This phase can only be billed after it has been closed. | [optional] 
**BillProjectClosedFlag** | Pointer to **NullableBool** | This phase can only be billed after the project has been closed. | [optional] 
**Downpayment** | Pointer to **NullableFloat64** |  | [optional] 
**PoNumber** | Pointer to **string** |  Max length: 25; | [optional] 
**PoAmount** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedTimeCost** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedExpenseCost** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedProductCost** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedTimeRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedExpenseRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedProductRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**BillToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**BillToContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**BillToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**ShipToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ShipToContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**ShipToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewProjectPhase

`func NewProjectPhase(description string, ) *ProjectPhase`

NewProjectPhase instantiates a new ProjectPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectPhaseWithDefaults

`func NewProjectPhaseWithDefaults() *ProjectPhase`

NewProjectPhaseWithDefaults instantiates a new ProjectPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectPhase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectPhase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectPhase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectPhase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectPhase) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectPhase) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectPhase) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectPhase) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ProjectPhase) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ProjectPhase) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetDescription

`func (o *ProjectPhase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectPhase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectPhase) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetBoard

`func (o *ProjectPhase) GetBoard() ProjectBoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *ProjectPhase) GetBoardOk() (*ProjectBoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *ProjectPhase) SetBoard(v ProjectBoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *ProjectPhase) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectPhase) GetStatus() PhaseStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectPhase) GetStatusOk() (*PhaseStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectPhase) SetStatus(v PhaseStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectPhase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAgreement

`func (o *ProjectPhase) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *ProjectPhase) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *ProjectPhase) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *ProjectPhase) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetOpportunity

`func (o *ProjectPhase) GetOpportunity() OpportunityReference`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *ProjectPhase) GetOpportunityOk() (*OpportunityReference, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *ProjectPhase) SetOpportunity(v OpportunityReference)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *ProjectPhase) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### GetDepartment

`func (o *ProjectPhase) GetDepartment() BillingUnitReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ProjectPhase) GetDepartmentOk() (*BillingUnitReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ProjectPhase) SetDepartment(v BillingUnitReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ProjectPhase) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetParentPhase

`func (o *ProjectPhase) GetParentPhase() ProjectPhaseReference`

GetParentPhase returns the ParentPhase field if non-nil, zero value otherwise.

### GetParentPhaseOk

`func (o *ProjectPhase) GetParentPhaseOk() (*ProjectPhaseReference, bool)`

GetParentPhaseOk returns a tuple with the ParentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPhase

`func (o *ProjectPhase) SetParentPhase(v ProjectPhaseReference)`

SetParentPhase sets ParentPhase field to given value.

### HasParentPhase

`func (o *ProjectPhase) HasParentPhase() bool`

HasParentPhase returns a boolean if a field has been set.

### GetWbsCode

`func (o *ProjectPhase) GetWbsCode() string`

GetWbsCode returns the WbsCode field if non-nil, zero value otherwise.

### GetWbsCodeOk

`func (o *ProjectPhase) GetWbsCodeOk() (*string, bool)`

GetWbsCodeOk returns a tuple with the WbsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbsCode

`func (o *ProjectPhase) SetWbsCode(v string)`

SetWbsCode sets WbsCode field to given value.

### HasWbsCode

`func (o *ProjectPhase) HasWbsCode() bool`

HasWbsCode returns a boolean if a field has been set.

### GetBillTime

`func (o *ProjectPhase) GetBillTime() string`

GetBillTime returns the BillTime field if non-nil, zero value otherwise.

### GetBillTimeOk

`func (o *ProjectPhase) GetBillTimeOk() (*string, bool)`

GetBillTimeOk returns a tuple with the BillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTime

`func (o *ProjectPhase) SetBillTime(v string)`

SetBillTime sets BillTime field to given value.

### HasBillTime

`func (o *ProjectPhase) HasBillTime() bool`

HasBillTime returns a boolean if a field has been set.

### SetBillTimeNil

`func (o *ProjectPhase) SetBillTimeNil(b bool)`

 SetBillTimeNil sets the value for BillTime to be an explicit nil

### UnsetBillTime
`func (o *ProjectPhase) UnsetBillTime()`

UnsetBillTime ensures that no value is present for BillTime, not even an explicit nil
### GetBillExpenses

`func (o *ProjectPhase) GetBillExpenses() string`

GetBillExpenses returns the BillExpenses field if non-nil, zero value otherwise.

### GetBillExpensesOk

`func (o *ProjectPhase) GetBillExpensesOk() (*string, bool)`

GetBillExpensesOk returns a tuple with the BillExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillExpenses

`func (o *ProjectPhase) SetBillExpenses(v string)`

SetBillExpenses sets BillExpenses field to given value.

### HasBillExpenses

`func (o *ProjectPhase) HasBillExpenses() bool`

HasBillExpenses returns a boolean if a field has been set.

### SetBillExpensesNil

`func (o *ProjectPhase) SetBillExpensesNil(b bool)`

 SetBillExpensesNil sets the value for BillExpenses to be an explicit nil

### UnsetBillExpenses
`func (o *ProjectPhase) UnsetBillExpenses()`

UnsetBillExpenses ensures that no value is present for BillExpenses, not even an explicit nil
### GetBillProducts

`func (o *ProjectPhase) GetBillProducts() string`

GetBillProducts returns the BillProducts field if non-nil, zero value otherwise.

### GetBillProductsOk

`func (o *ProjectPhase) GetBillProductsOk() (*string, bool)`

GetBillProductsOk returns a tuple with the BillProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProducts

`func (o *ProjectPhase) SetBillProducts(v string)`

SetBillProducts sets BillProducts field to given value.

### HasBillProducts

`func (o *ProjectPhase) HasBillProducts() bool`

HasBillProducts returns a boolean if a field has been set.

### SetBillProductsNil

`func (o *ProjectPhase) SetBillProductsNil(b bool)`

 SetBillProductsNil sets the value for BillProducts to be an explicit nil

### UnsetBillProducts
`func (o *ProjectPhase) UnsetBillProducts()`

UnsetBillProducts ensures that no value is present for BillProducts, not even an explicit nil
### GetMarkAsMilestoneFlag

`func (o *ProjectPhase) GetMarkAsMilestoneFlag() bool`

GetMarkAsMilestoneFlag returns the MarkAsMilestoneFlag field if non-nil, zero value otherwise.

### GetMarkAsMilestoneFlagOk

`func (o *ProjectPhase) GetMarkAsMilestoneFlagOk() (*bool, bool)`

GetMarkAsMilestoneFlagOk returns a tuple with the MarkAsMilestoneFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAsMilestoneFlag

`func (o *ProjectPhase) SetMarkAsMilestoneFlag(v bool)`

SetMarkAsMilestoneFlag sets MarkAsMilestoneFlag field to given value.

### HasMarkAsMilestoneFlag

`func (o *ProjectPhase) HasMarkAsMilestoneFlag() bool`

HasMarkAsMilestoneFlag returns a boolean if a field has been set.

### SetMarkAsMilestoneFlagNil

`func (o *ProjectPhase) SetMarkAsMilestoneFlagNil(b bool)`

 SetMarkAsMilestoneFlagNil sets the value for MarkAsMilestoneFlag to be an explicit nil

### UnsetMarkAsMilestoneFlag
`func (o *ProjectPhase) UnsetMarkAsMilestoneFlag()`

UnsetMarkAsMilestoneFlag ensures that no value is present for MarkAsMilestoneFlag, not even an explicit nil
### GetNotes

`func (o *ProjectPhase) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ProjectPhase) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ProjectPhase) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ProjectPhase) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDeadlineDate

`func (o *ProjectPhase) GetDeadlineDate() time.Time`

GetDeadlineDate returns the DeadlineDate field if non-nil, zero value otherwise.

### GetDeadlineDateOk

`func (o *ProjectPhase) GetDeadlineDateOk() (*time.Time, bool)`

GetDeadlineDateOk returns a tuple with the DeadlineDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineDate

`func (o *ProjectPhase) SetDeadlineDate(v time.Time)`

SetDeadlineDate sets DeadlineDate field to given value.

### HasDeadlineDate

`func (o *ProjectPhase) HasDeadlineDate() bool`

HasDeadlineDate returns a boolean if a field has been set.

### GetBillSeparatelyFlag

`func (o *ProjectPhase) GetBillSeparatelyFlag() bool`

GetBillSeparatelyFlag returns the BillSeparatelyFlag field if non-nil, zero value otherwise.

### GetBillSeparatelyFlagOk

`func (o *ProjectPhase) GetBillSeparatelyFlagOk() (*bool, bool)`

GetBillSeparatelyFlagOk returns a tuple with the BillSeparatelyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillSeparatelyFlag

`func (o *ProjectPhase) SetBillSeparatelyFlag(v bool)`

SetBillSeparatelyFlag sets BillSeparatelyFlag field to given value.

### HasBillSeparatelyFlag

`func (o *ProjectPhase) HasBillSeparatelyFlag() bool`

HasBillSeparatelyFlag returns a boolean if a field has been set.

### SetBillSeparatelyFlagNil

`func (o *ProjectPhase) SetBillSeparatelyFlagNil(b bool)`

 SetBillSeparatelyFlagNil sets the value for BillSeparatelyFlag to be an explicit nil

### UnsetBillSeparatelyFlag
`func (o *ProjectPhase) UnsetBillSeparatelyFlag()`

UnsetBillSeparatelyFlag ensures that no value is present for BillSeparatelyFlag, not even an explicit nil
### GetBillingMethod

`func (o *ProjectPhase) GetBillingMethod() string`

GetBillingMethod returns the BillingMethod field if non-nil, zero value otherwise.

### GetBillingMethodOk

`func (o *ProjectPhase) GetBillingMethodOk() (*string, bool)`

GetBillingMethodOk returns a tuple with the BillingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMethod

`func (o *ProjectPhase) SetBillingMethod(v string)`

SetBillingMethod sets BillingMethod field to given value.

### HasBillingMethod

`func (o *ProjectPhase) HasBillingMethod() bool`

HasBillingMethod returns a boolean if a field has been set.

### SetBillingMethodNil

`func (o *ProjectPhase) SetBillingMethodNil(b bool)`

 SetBillingMethodNil sets the value for BillingMethod to be an explicit nil

### UnsetBillingMethod
`func (o *ProjectPhase) UnsetBillingMethod()`

UnsetBillingMethod ensures that no value is present for BillingMethod, not even an explicit nil
### GetScheduledHours

`func (o *ProjectPhase) GetScheduledHours() float64`

GetScheduledHours returns the ScheduledHours field if non-nil, zero value otherwise.

### GetScheduledHoursOk

`func (o *ProjectPhase) GetScheduledHoursOk() (*float64, bool)`

GetScheduledHoursOk returns a tuple with the ScheduledHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledHours

`func (o *ProjectPhase) SetScheduledHours(v float64)`

SetScheduledHours sets ScheduledHours field to given value.

### HasScheduledHours

`func (o *ProjectPhase) HasScheduledHours() bool`

HasScheduledHours returns a boolean if a field has been set.

### SetScheduledHoursNil

`func (o *ProjectPhase) SetScheduledHoursNil(b bool)`

 SetScheduledHoursNil sets the value for ScheduledHours to be an explicit nil

### UnsetScheduledHours
`func (o *ProjectPhase) UnsetScheduledHours()`

UnsetScheduledHours ensures that no value is present for ScheduledHours, not even an explicit nil
### GetScheduledStart

`func (o *ProjectPhase) GetScheduledStart() string`

GetScheduledStart returns the ScheduledStart field if non-nil, zero value otherwise.

### GetScheduledStartOk

`func (o *ProjectPhase) GetScheduledStartOk() (*string, bool)`

GetScheduledStartOk returns a tuple with the ScheduledStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStart

`func (o *ProjectPhase) SetScheduledStart(v string)`

SetScheduledStart sets ScheduledStart field to given value.

### HasScheduledStart

`func (o *ProjectPhase) HasScheduledStart() bool`

HasScheduledStart returns a boolean if a field has been set.

### GetScheduledEnd

`func (o *ProjectPhase) GetScheduledEnd() string`

GetScheduledEnd returns the ScheduledEnd field if non-nil, zero value otherwise.

### GetScheduledEndOk

`func (o *ProjectPhase) GetScheduledEndOk() (*string, bool)`

GetScheduledEndOk returns a tuple with the ScheduledEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEnd

`func (o *ProjectPhase) SetScheduledEnd(v string)`

SetScheduledEnd sets ScheduledEnd field to given value.

### HasScheduledEnd

`func (o *ProjectPhase) HasScheduledEnd() bool`

HasScheduledEnd returns a boolean if a field has been set.

### GetActualHours

`func (o *ProjectPhase) GetActualHours() float64`

GetActualHours returns the ActualHours field if non-nil, zero value otherwise.

### GetActualHoursOk

`func (o *ProjectPhase) GetActualHoursOk() (*float64, bool)`

GetActualHoursOk returns a tuple with the ActualHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualHours

`func (o *ProjectPhase) SetActualHours(v float64)`

SetActualHours sets ActualHours field to given value.

### HasActualHours

`func (o *ProjectPhase) HasActualHours() bool`

HasActualHours returns a boolean if a field has been set.

### SetActualHoursNil

`func (o *ProjectPhase) SetActualHoursNil(b bool)`

 SetActualHoursNil sets the value for ActualHours to be an explicit nil

### UnsetActualHours
`func (o *ProjectPhase) UnsetActualHours()`

UnsetActualHours ensures that no value is present for ActualHours, not even an explicit nil
### GetActualStart

`func (o *ProjectPhase) GetActualStart() string`

GetActualStart returns the ActualStart field if non-nil, zero value otherwise.

### GetActualStartOk

`func (o *ProjectPhase) GetActualStartOk() (*string, bool)`

GetActualStartOk returns a tuple with the ActualStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualStart

`func (o *ProjectPhase) SetActualStart(v string)`

SetActualStart sets ActualStart field to given value.

### HasActualStart

`func (o *ProjectPhase) HasActualStart() bool`

HasActualStart returns a boolean if a field has been set.

### GetActualEnd

`func (o *ProjectPhase) GetActualEnd() string`

GetActualEnd returns the ActualEnd field if non-nil, zero value otherwise.

### GetActualEndOk

`func (o *ProjectPhase) GetActualEndOk() (*string, bool)`

GetActualEndOk returns a tuple with the ActualEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualEnd

`func (o *ProjectPhase) SetActualEnd(v string)`

SetActualEnd sets ActualEnd field to given value.

### HasActualEnd

`func (o *ProjectPhase) HasActualEnd() bool`

HasActualEnd returns a boolean if a field has been set.

### GetBudgetHours

`func (o *ProjectPhase) GetBudgetHours() float64`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *ProjectPhase) GetBudgetHoursOk() (*float64, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *ProjectPhase) SetBudgetHours(v float64)`

SetBudgetHours sets BudgetHours field to given value.

### HasBudgetHours

`func (o *ProjectPhase) HasBudgetHours() bool`

HasBudgetHours returns a boolean if a field has been set.

### SetBudgetHoursNil

`func (o *ProjectPhase) SetBudgetHoursNil(b bool)`

 SetBudgetHoursNil sets the value for BudgetHours to be an explicit nil

### UnsetBudgetHours
`func (o *ProjectPhase) UnsetBudgetHours()`

UnsetBudgetHours ensures that no value is present for BudgetHours, not even an explicit nil
### GetStartDate

`func (o *ProjectPhase) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ProjectPhase) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ProjectPhase) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ProjectPhase) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ProjectPhase) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ProjectPhase) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ProjectPhase) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ProjectPhase) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLocationId

`func (o *ProjectPhase) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ProjectPhase) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ProjectPhase) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *ProjectPhase) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *ProjectPhase) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *ProjectPhase) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetBusinessUnitId

`func (o *ProjectPhase) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *ProjectPhase) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *ProjectPhase) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *ProjectPhase) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *ProjectPhase) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *ProjectPhase) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetHourlyRate

`func (o *ProjectPhase) GetHourlyRate() float64`

GetHourlyRate returns the HourlyRate field if non-nil, zero value otherwise.

### GetHourlyRateOk

`func (o *ProjectPhase) GetHourlyRateOk() (*float64, bool)`

GetHourlyRateOk returns a tuple with the HourlyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyRate

`func (o *ProjectPhase) SetHourlyRate(v float64)`

SetHourlyRate sets HourlyRate field to given value.

### HasHourlyRate

`func (o *ProjectPhase) HasHourlyRate() bool`

HasHourlyRate returns a boolean if a field has been set.

### SetHourlyRateNil

`func (o *ProjectPhase) SetHourlyRateNil(b bool)`

 SetHourlyRateNil sets the value for HourlyRate to be an explicit nil

### UnsetHourlyRate
`func (o *ProjectPhase) UnsetHourlyRate()`

UnsetHourlyRate ensures that no value is present for HourlyRate, not even an explicit nil
### GetBillingStartDate

`func (o *ProjectPhase) GetBillingStartDate() time.Time`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *ProjectPhase) GetBillingStartDateOk() (*time.Time, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *ProjectPhase) SetBillingStartDate(v time.Time)`

SetBillingStartDate sets BillingStartDate field to given value.

### HasBillingStartDate

`func (o *ProjectPhase) HasBillingStartDate() bool`

HasBillingStartDate returns a boolean if a field has been set.

### GetBillPhaseClosedFlag

`func (o *ProjectPhase) GetBillPhaseClosedFlag() bool`

GetBillPhaseClosedFlag returns the BillPhaseClosedFlag field if non-nil, zero value otherwise.

### GetBillPhaseClosedFlagOk

`func (o *ProjectPhase) GetBillPhaseClosedFlagOk() (*bool, bool)`

GetBillPhaseClosedFlagOk returns a tuple with the BillPhaseClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillPhaseClosedFlag

`func (o *ProjectPhase) SetBillPhaseClosedFlag(v bool)`

SetBillPhaseClosedFlag sets BillPhaseClosedFlag field to given value.

### HasBillPhaseClosedFlag

`func (o *ProjectPhase) HasBillPhaseClosedFlag() bool`

HasBillPhaseClosedFlag returns a boolean if a field has been set.

### SetBillPhaseClosedFlagNil

`func (o *ProjectPhase) SetBillPhaseClosedFlagNil(b bool)`

 SetBillPhaseClosedFlagNil sets the value for BillPhaseClosedFlag to be an explicit nil

### UnsetBillPhaseClosedFlag
`func (o *ProjectPhase) UnsetBillPhaseClosedFlag()`

UnsetBillPhaseClosedFlag ensures that no value is present for BillPhaseClosedFlag, not even an explicit nil
### GetBillProjectClosedFlag

`func (o *ProjectPhase) GetBillProjectClosedFlag() bool`

GetBillProjectClosedFlag returns the BillProjectClosedFlag field if non-nil, zero value otherwise.

### GetBillProjectClosedFlagOk

`func (o *ProjectPhase) GetBillProjectClosedFlagOk() (*bool, bool)`

GetBillProjectClosedFlagOk returns a tuple with the BillProjectClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProjectClosedFlag

`func (o *ProjectPhase) SetBillProjectClosedFlag(v bool)`

SetBillProjectClosedFlag sets BillProjectClosedFlag field to given value.

### HasBillProjectClosedFlag

`func (o *ProjectPhase) HasBillProjectClosedFlag() bool`

HasBillProjectClosedFlag returns a boolean if a field has been set.

### SetBillProjectClosedFlagNil

`func (o *ProjectPhase) SetBillProjectClosedFlagNil(b bool)`

 SetBillProjectClosedFlagNil sets the value for BillProjectClosedFlag to be an explicit nil

### UnsetBillProjectClosedFlag
`func (o *ProjectPhase) UnsetBillProjectClosedFlag()`

UnsetBillProjectClosedFlag ensures that no value is present for BillProjectClosedFlag, not even an explicit nil
### GetDownpayment

`func (o *ProjectPhase) GetDownpayment() float64`

GetDownpayment returns the Downpayment field if non-nil, zero value otherwise.

### GetDownpaymentOk

`func (o *ProjectPhase) GetDownpaymentOk() (*float64, bool)`

GetDownpaymentOk returns a tuple with the Downpayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownpayment

`func (o *ProjectPhase) SetDownpayment(v float64)`

SetDownpayment sets Downpayment field to given value.

### HasDownpayment

`func (o *ProjectPhase) HasDownpayment() bool`

HasDownpayment returns a boolean if a field has been set.

### SetDownpaymentNil

`func (o *ProjectPhase) SetDownpaymentNil(b bool)`

 SetDownpaymentNil sets the value for Downpayment to be an explicit nil

### UnsetDownpayment
`func (o *ProjectPhase) UnsetDownpayment()`

UnsetDownpayment ensures that no value is present for Downpayment, not even an explicit nil
### GetPoNumber

`func (o *ProjectPhase) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *ProjectPhase) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *ProjectPhase) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *ProjectPhase) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetPoAmount

`func (o *ProjectPhase) GetPoAmount() float64`

GetPoAmount returns the PoAmount field if non-nil, zero value otherwise.

### GetPoAmountOk

`func (o *ProjectPhase) GetPoAmountOk() (*float64, bool)`

GetPoAmountOk returns a tuple with the PoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoAmount

`func (o *ProjectPhase) SetPoAmount(v float64)`

SetPoAmount sets PoAmount field to given value.

### HasPoAmount

`func (o *ProjectPhase) HasPoAmount() bool`

HasPoAmount returns a boolean if a field has been set.

### SetPoAmountNil

`func (o *ProjectPhase) SetPoAmountNil(b bool)`

 SetPoAmountNil sets the value for PoAmount to be an explicit nil

### UnsetPoAmount
`func (o *ProjectPhase) UnsetPoAmount()`

UnsetPoAmount ensures that no value is present for PoAmount, not even an explicit nil
### GetEstimatedTimeCost

`func (o *ProjectPhase) GetEstimatedTimeCost() float64`

GetEstimatedTimeCost returns the EstimatedTimeCost field if non-nil, zero value otherwise.

### GetEstimatedTimeCostOk

`func (o *ProjectPhase) GetEstimatedTimeCostOk() (*float64, bool)`

GetEstimatedTimeCostOk returns a tuple with the EstimatedTimeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeCost

`func (o *ProjectPhase) SetEstimatedTimeCost(v float64)`

SetEstimatedTimeCost sets EstimatedTimeCost field to given value.

### HasEstimatedTimeCost

`func (o *ProjectPhase) HasEstimatedTimeCost() bool`

HasEstimatedTimeCost returns a boolean if a field has been set.

### SetEstimatedTimeCostNil

`func (o *ProjectPhase) SetEstimatedTimeCostNil(b bool)`

 SetEstimatedTimeCostNil sets the value for EstimatedTimeCost to be an explicit nil

### UnsetEstimatedTimeCost
`func (o *ProjectPhase) UnsetEstimatedTimeCost()`

UnsetEstimatedTimeCost ensures that no value is present for EstimatedTimeCost, not even an explicit nil
### GetEstimatedExpenseCost

`func (o *ProjectPhase) GetEstimatedExpenseCost() float64`

GetEstimatedExpenseCost returns the EstimatedExpenseCost field if non-nil, zero value otherwise.

### GetEstimatedExpenseCostOk

`func (o *ProjectPhase) GetEstimatedExpenseCostOk() (*float64, bool)`

GetEstimatedExpenseCostOk returns a tuple with the EstimatedExpenseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedExpenseCost

`func (o *ProjectPhase) SetEstimatedExpenseCost(v float64)`

SetEstimatedExpenseCost sets EstimatedExpenseCost field to given value.

### HasEstimatedExpenseCost

`func (o *ProjectPhase) HasEstimatedExpenseCost() bool`

HasEstimatedExpenseCost returns a boolean if a field has been set.

### SetEstimatedExpenseCostNil

`func (o *ProjectPhase) SetEstimatedExpenseCostNil(b bool)`

 SetEstimatedExpenseCostNil sets the value for EstimatedExpenseCost to be an explicit nil

### UnsetEstimatedExpenseCost
`func (o *ProjectPhase) UnsetEstimatedExpenseCost()`

UnsetEstimatedExpenseCost ensures that no value is present for EstimatedExpenseCost, not even an explicit nil
### GetEstimatedProductCost

`func (o *ProjectPhase) GetEstimatedProductCost() float64`

GetEstimatedProductCost returns the EstimatedProductCost field if non-nil, zero value otherwise.

### GetEstimatedProductCostOk

`func (o *ProjectPhase) GetEstimatedProductCostOk() (*float64, bool)`

GetEstimatedProductCostOk returns a tuple with the EstimatedProductCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedProductCost

`func (o *ProjectPhase) SetEstimatedProductCost(v float64)`

SetEstimatedProductCost sets EstimatedProductCost field to given value.

### HasEstimatedProductCost

`func (o *ProjectPhase) HasEstimatedProductCost() bool`

HasEstimatedProductCost returns a boolean if a field has been set.

### SetEstimatedProductCostNil

`func (o *ProjectPhase) SetEstimatedProductCostNil(b bool)`

 SetEstimatedProductCostNil sets the value for EstimatedProductCost to be an explicit nil

### UnsetEstimatedProductCost
`func (o *ProjectPhase) UnsetEstimatedProductCost()`

UnsetEstimatedProductCost ensures that no value is present for EstimatedProductCost, not even an explicit nil
### GetEstimatedTimeRevenue

`func (o *ProjectPhase) GetEstimatedTimeRevenue() float64`

GetEstimatedTimeRevenue returns the EstimatedTimeRevenue field if non-nil, zero value otherwise.

### GetEstimatedTimeRevenueOk

`func (o *ProjectPhase) GetEstimatedTimeRevenueOk() (*float64, bool)`

GetEstimatedTimeRevenueOk returns a tuple with the EstimatedTimeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeRevenue

`func (o *ProjectPhase) SetEstimatedTimeRevenue(v float64)`

SetEstimatedTimeRevenue sets EstimatedTimeRevenue field to given value.

### HasEstimatedTimeRevenue

`func (o *ProjectPhase) HasEstimatedTimeRevenue() bool`

HasEstimatedTimeRevenue returns a boolean if a field has been set.

### SetEstimatedTimeRevenueNil

`func (o *ProjectPhase) SetEstimatedTimeRevenueNil(b bool)`

 SetEstimatedTimeRevenueNil sets the value for EstimatedTimeRevenue to be an explicit nil

### UnsetEstimatedTimeRevenue
`func (o *ProjectPhase) UnsetEstimatedTimeRevenue()`

UnsetEstimatedTimeRevenue ensures that no value is present for EstimatedTimeRevenue, not even an explicit nil
### GetEstimatedExpenseRevenue

`func (o *ProjectPhase) GetEstimatedExpenseRevenue() float64`

GetEstimatedExpenseRevenue returns the EstimatedExpenseRevenue field if non-nil, zero value otherwise.

### GetEstimatedExpenseRevenueOk

`func (o *ProjectPhase) GetEstimatedExpenseRevenueOk() (*float64, bool)`

GetEstimatedExpenseRevenueOk returns a tuple with the EstimatedExpenseRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedExpenseRevenue

`func (o *ProjectPhase) SetEstimatedExpenseRevenue(v float64)`

SetEstimatedExpenseRevenue sets EstimatedExpenseRevenue field to given value.

### HasEstimatedExpenseRevenue

`func (o *ProjectPhase) HasEstimatedExpenseRevenue() bool`

HasEstimatedExpenseRevenue returns a boolean if a field has been set.

### SetEstimatedExpenseRevenueNil

`func (o *ProjectPhase) SetEstimatedExpenseRevenueNil(b bool)`

 SetEstimatedExpenseRevenueNil sets the value for EstimatedExpenseRevenue to be an explicit nil

### UnsetEstimatedExpenseRevenue
`func (o *ProjectPhase) UnsetEstimatedExpenseRevenue()`

UnsetEstimatedExpenseRevenue ensures that no value is present for EstimatedExpenseRevenue, not even an explicit nil
### GetEstimatedProductRevenue

`func (o *ProjectPhase) GetEstimatedProductRevenue() float64`

GetEstimatedProductRevenue returns the EstimatedProductRevenue field if non-nil, zero value otherwise.

### GetEstimatedProductRevenueOk

`func (o *ProjectPhase) GetEstimatedProductRevenueOk() (*float64, bool)`

GetEstimatedProductRevenueOk returns a tuple with the EstimatedProductRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedProductRevenue

`func (o *ProjectPhase) SetEstimatedProductRevenue(v float64)`

SetEstimatedProductRevenue sets EstimatedProductRevenue field to given value.

### HasEstimatedProductRevenue

`func (o *ProjectPhase) HasEstimatedProductRevenue() bool`

HasEstimatedProductRevenue returns a boolean if a field has been set.

### SetEstimatedProductRevenueNil

`func (o *ProjectPhase) SetEstimatedProductRevenueNil(b bool)`

 SetEstimatedProductRevenueNil sets the value for EstimatedProductRevenue to be an explicit nil

### UnsetEstimatedProductRevenue
`func (o *ProjectPhase) UnsetEstimatedProductRevenue()`

UnsetEstimatedProductRevenue ensures that no value is present for EstimatedProductRevenue, not even an explicit nil
### GetCurrency

`func (o *ProjectPhase) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProjectPhase) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProjectPhase) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProjectPhase) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBillToCompany

`func (o *ProjectPhase) GetBillToCompany() CompanyReference`

GetBillToCompany returns the BillToCompany field if non-nil, zero value otherwise.

### GetBillToCompanyOk

`func (o *ProjectPhase) GetBillToCompanyOk() (*CompanyReference, bool)`

GetBillToCompanyOk returns a tuple with the BillToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToCompany

`func (o *ProjectPhase) SetBillToCompany(v CompanyReference)`

SetBillToCompany sets BillToCompany field to given value.

### HasBillToCompany

`func (o *ProjectPhase) HasBillToCompany() bool`

HasBillToCompany returns a boolean if a field has been set.

### GetBillToContact

`func (o *ProjectPhase) GetBillToContact() ContactReference`

GetBillToContact returns the BillToContact field if non-nil, zero value otherwise.

### GetBillToContactOk

`func (o *ProjectPhase) GetBillToContactOk() (*ContactReference, bool)`

GetBillToContactOk returns a tuple with the BillToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToContact

`func (o *ProjectPhase) SetBillToContact(v ContactReference)`

SetBillToContact sets BillToContact field to given value.

### HasBillToContact

`func (o *ProjectPhase) HasBillToContact() bool`

HasBillToContact returns a boolean if a field has been set.

### GetBillToSite

`func (o *ProjectPhase) GetBillToSite() SiteReference`

GetBillToSite returns the BillToSite field if non-nil, zero value otherwise.

### GetBillToSiteOk

`func (o *ProjectPhase) GetBillToSiteOk() (*SiteReference, bool)`

GetBillToSiteOk returns a tuple with the BillToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToSite

`func (o *ProjectPhase) SetBillToSite(v SiteReference)`

SetBillToSite sets BillToSite field to given value.

### HasBillToSite

`func (o *ProjectPhase) HasBillToSite() bool`

HasBillToSite returns a boolean if a field has been set.

### GetShipToCompany

`func (o *ProjectPhase) GetShipToCompany() CompanyReference`

GetShipToCompany returns the ShipToCompany field if non-nil, zero value otherwise.

### GetShipToCompanyOk

`func (o *ProjectPhase) GetShipToCompanyOk() (*CompanyReference, bool)`

GetShipToCompanyOk returns a tuple with the ShipToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompany

`func (o *ProjectPhase) SetShipToCompany(v CompanyReference)`

SetShipToCompany sets ShipToCompany field to given value.

### HasShipToCompany

`func (o *ProjectPhase) HasShipToCompany() bool`

HasShipToCompany returns a boolean if a field has been set.

### GetShipToContact

`func (o *ProjectPhase) GetShipToContact() ContactReference`

GetShipToContact returns the ShipToContact field if non-nil, zero value otherwise.

### GetShipToContactOk

`func (o *ProjectPhase) GetShipToContactOk() (*ContactReference, bool)`

GetShipToContactOk returns a tuple with the ShipToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToContact

`func (o *ProjectPhase) SetShipToContact(v ContactReference)`

SetShipToContact sets ShipToContact field to given value.

### HasShipToContact

`func (o *ProjectPhase) HasShipToContact() bool`

HasShipToContact returns a boolean if a field has been set.

### GetShipToSite

`func (o *ProjectPhase) GetShipToSite() SiteReference`

GetShipToSite returns the ShipToSite field if non-nil, zero value otherwise.

### GetShipToSiteOk

`func (o *ProjectPhase) GetShipToSiteOk() (*SiteReference, bool)`

GetShipToSiteOk returns a tuple with the ShipToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToSite

`func (o *ProjectPhase) SetShipToSite(v SiteReference)`

SetShipToSite sets ShipToSite field to given value.

### HasShipToSite

`func (o *ProjectPhase) HasShipToSite() bool`

HasShipToSite returns a boolean if a field has been set.

### GetBillingTerms

`func (o *ProjectPhase) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *ProjectPhase) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *ProjectPhase) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *ProjectPhase) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetTaxCode

`func (o *ProjectPhase) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *ProjectPhase) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *ProjectPhase) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *ProjectPhase) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetInfo

`func (o *ProjectPhase) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectPhase) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectPhase) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectPhase) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *ProjectPhase) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProjectPhase) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProjectPhase) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProjectPhase) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


