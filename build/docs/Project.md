# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ActualEnd** | Pointer to **time.Time** |  | [optional] 
**ActualHours** | Pointer to **NullableFloat64** |  | [optional] 
**ActualStart** | Pointer to **time.Time** |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**BillExpenses** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**BillingAmount** | Pointer to **NullableFloat64** |  | [optional] 
**BillingAttention** | Pointer to **string** |  Max length: 50; | [optional] 
**BillingMethod** | **NullableString** |  | 
**BillingRateType** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**BillProducts** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**BillProjectAfterClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**BillTime** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**BillToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**BillToContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**BillToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**BillUnapprovedTimeAndExpense** | Pointer to **NullableBool** |  | [optional] 
**Board** | [**ProjectBoardReference**](ProjectBoardReference.md) |  | 
**BudgetAnalysis** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**BudgetFlag** | Pointer to **NullableBool** |  | [optional] 
**BudgetHours** | Pointer to **NullableFloat64** |  | [optional] 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**CustomerPO** | Pointer to **string** |  Max length: 50; | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Downpayment** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedEnd** | **time.Time** |  | 
**PercentComplete** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedExpenseRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedHours** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedProductRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedStart** | **time.Time** |  | 
**EstimatedTimeRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**ExpenseApprover** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**IncludeDependenciesFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeEstimatesFlag** | Pointer to **NullableBool** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**Manager** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Name** | **string** |  Max length: 100; | 
**Opportunity** | Pointer to [**OpportunityReference**](OpportunityReference.md) |  | [optional] 
**ProjectTemplateId** | Pointer to **NullableInt32** |  | [optional] 
**RestrictDownPaymentFlag** | Pointer to **NullableBool** |  | [optional] 
**ScheduledEnd** | Pointer to **time.Time** |  | [optional] 
**ScheduledHours** | Pointer to **NullableFloat64** |  | [optional] 
**ScheduledStart** | Pointer to **time.Time** |  | [optional] 
**ShipToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ShipToContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**ShipToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**Status** | Pointer to [**ProjectStatusReference**](ProjectStatusReference.md) |  | [optional] 
**ClosedFlag** | Pointer to **bool** |  | [optional] 
**TimeApprover** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Type** | Pointer to [**ProjectTypeReference**](ProjectTypeReference.md) |  | [optional] 
**DoNotDisplayInPortalFlag** | Pointer to **NullableBool** |  | [optional] 
**BillingStartDate** | Pointer to **time.Time** |  | [optional] 
**PoAmount** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedTimeCost** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedExpenseCost** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedProductCost** | Pointer to **NullableFloat64** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**CompanyLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewProject

`func NewProject(billingMethod NullableString, board ProjectBoardReference, company CompanyReference, estimatedEnd time.Time, estimatedStart time.Time, name string, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Project) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Project) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActualEnd

`func (o *Project) GetActualEnd() time.Time`

GetActualEnd returns the ActualEnd field if non-nil, zero value otherwise.

### GetActualEndOk

`func (o *Project) GetActualEndOk() (*time.Time, bool)`

GetActualEndOk returns a tuple with the ActualEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualEnd

`func (o *Project) SetActualEnd(v time.Time)`

SetActualEnd sets ActualEnd field to given value.

### HasActualEnd

`func (o *Project) HasActualEnd() bool`

HasActualEnd returns a boolean if a field has been set.

### GetActualHours

`func (o *Project) GetActualHours() float64`

GetActualHours returns the ActualHours field if non-nil, zero value otherwise.

### GetActualHoursOk

`func (o *Project) GetActualHoursOk() (*float64, bool)`

GetActualHoursOk returns a tuple with the ActualHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualHours

`func (o *Project) SetActualHours(v float64)`

SetActualHours sets ActualHours field to given value.

### HasActualHours

`func (o *Project) HasActualHours() bool`

HasActualHours returns a boolean if a field has been set.

### SetActualHoursNil

`func (o *Project) SetActualHoursNil(b bool)`

 SetActualHoursNil sets the value for ActualHours to be an explicit nil

### UnsetActualHours
`func (o *Project) UnsetActualHours()`

UnsetActualHours ensures that no value is present for ActualHours, not even an explicit nil
### GetActualStart

`func (o *Project) GetActualStart() time.Time`

GetActualStart returns the ActualStart field if non-nil, zero value otherwise.

### GetActualStartOk

`func (o *Project) GetActualStartOk() (*time.Time, bool)`

GetActualStartOk returns a tuple with the ActualStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualStart

`func (o *Project) SetActualStart(v time.Time)`

SetActualStart sets ActualStart field to given value.

### HasActualStart

`func (o *Project) HasActualStart() bool`

HasActualStart returns a boolean if a field has been set.

### GetAgreement

`func (o *Project) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *Project) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *Project) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *Project) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetBillExpenses

`func (o *Project) GetBillExpenses() string`

GetBillExpenses returns the BillExpenses field if non-nil, zero value otherwise.

### GetBillExpensesOk

`func (o *Project) GetBillExpensesOk() (*string, bool)`

GetBillExpensesOk returns a tuple with the BillExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillExpenses

`func (o *Project) SetBillExpenses(v string)`

SetBillExpenses sets BillExpenses field to given value.

### HasBillExpenses

`func (o *Project) HasBillExpenses() bool`

HasBillExpenses returns a boolean if a field has been set.

### SetBillExpensesNil

`func (o *Project) SetBillExpensesNil(b bool)`

 SetBillExpensesNil sets the value for BillExpenses to be an explicit nil

### UnsetBillExpenses
`func (o *Project) UnsetBillExpenses()`

UnsetBillExpenses ensures that no value is present for BillExpenses, not even an explicit nil
### GetBillingAmount

`func (o *Project) GetBillingAmount() float64`

GetBillingAmount returns the BillingAmount field if non-nil, zero value otherwise.

### GetBillingAmountOk

`func (o *Project) GetBillingAmountOk() (*float64, bool)`

GetBillingAmountOk returns a tuple with the BillingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAmount

`func (o *Project) SetBillingAmount(v float64)`

SetBillingAmount sets BillingAmount field to given value.

### HasBillingAmount

`func (o *Project) HasBillingAmount() bool`

HasBillingAmount returns a boolean if a field has been set.

### SetBillingAmountNil

`func (o *Project) SetBillingAmountNil(b bool)`

 SetBillingAmountNil sets the value for BillingAmount to be an explicit nil

### UnsetBillingAmount
`func (o *Project) UnsetBillingAmount()`

UnsetBillingAmount ensures that no value is present for BillingAmount, not even an explicit nil
### GetBillingAttention

`func (o *Project) GetBillingAttention() string`

GetBillingAttention returns the BillingAttention field if non-nil, zero value otherwise.

### GetBillingAttentionOk

`func (o *Project) GetBillingAttentionOk() (*string, bool)`

GetBillingAttentionOk returns a tuple with the BillingAttention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAttention

`func (o *Project) SetBillingAttention(v string)`

SetBillingAttention sets BillingAttention field to given value.

### HasBillingAttention

`func (o *Project) HasBillingAttention() bool`

HasBillingAttention returns a boolean if a field has been set.

### GetBillingMethod

`func (o *Project) GetBillingMethod() string`

GetBillingMethod returns the BillingMethod field if non-nil, zero value otherwise.

### GetBillingMethodOk

`func (o *Project) GetBillingMethodOk() (*string, bool)`

GetBillingMethodOk returns a tuple with the BillingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMethod

`func (o *Project) SetBillingMethod(v string)`

SetBillingMethod sets BillingMethod field to given value.


### SetBillingMethodNil

`func (o *Project) SetBillingMethodNil(b bool)`

 SetBillingMethodNil sets the value for BillingMethod to be an explicit nil

### UnsetBillingMethod
`func (o *Project) UnsetBillingMethod()`

UnsetBillingMethod ensures that no value is present for BillingMethod, not even an explicit nil
### GetBillingRateType

`func (o *Project) GetBillingRateType() string`

GetBillingRateType returns the BillingRateType field if non-nil, zero value otherwise.

### GetBillingRateTypeOk

`func (o *Project) GetBillingRateTypeOk() (*string, bool)`

GetBillingRateTypeOk returns a tuple with the BillingRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRateType

`func (o *Project) SetBillingRateType(v string)`

SetBillingRateType sets BillingRateType field to given value.

### HasBillingRateType

`func (o *Project) HasBillingRateType() bool`

HasBillingRateType returns a boolean if a field has been set.

### SetBillingRateTypeNil

`func (o *Project) SetBillingRateTypeNil(b bool)`

 SetBillingRateTypeNil sets the value for BillingRateType to be an explicit nil

### UnsetBillingRateType
`func (o *Project) UnsetBillingRateType()`

UnsetBillingRateType ensures that no value is present for BillingRateType, not even an explicit nil
### GetBillingTerms

`func (o *Project) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *Project) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *Project) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *Project) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetBillProducts

`func (o *Project) GetBillProducts() string`

GetBillProducts returns the BillProducts field if non-nil, zero value otherwise.

### GetBillProductsOk

`func (o *Project) GetBillProductsOk() (*string, bool)`

GetBillProductsOk returns a tuple with the BillProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProducts

`func (o *Project) SetBillProducts(v string)`

SetBillProducts sets BillProducts field to given value.

### HasBillProducts

`func (o *Project) HasBillProducts() bool`

HasBillProducts returns a boolean if a field has been set.

### SetBillProductsNil

`func (o *Project) SetBillProductsNil(b bool)`

 SetBillProductsNil sets the value for BillProducts to be an explicit nil

### UnsetBillProducts
`func (o *Project) UnsetBillProducts()`

UnsetBillProducts ensures that no value is present for BillProducts, not even an explicit nil
### GetBillProjectAfterClosedFlag

`func (o *Project) GetBillProjectAfterClosedFlag() bool`

GetBillProjectAfterClosedFlag returns the BillProjectAfterClosedFlag field if non-nil, zero value otherwise.

### GetBillProjectAfterClosedFlagOk

`func (o *Project) GetBillProjectAfterClosedFlagOk() (*bool, bool)`

GetBillProjectAfterClosedFlagOk returns a tuple with the BillProjectAfterClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProjectAfterClosedFlag

`func (o *Project) SetBillProjectAfterClosedFlag(v bool)`

SetBillProjectAfterClosedFlag sets BillProjectAfterClosedFlag field to given value.

### HasBillProjectAfterClosedFlag

`func (o *Project) HasBillProjectAfterClosedFlag() bool`

HasBillProjectAfterClosedFlag returns a boolean if a field has been set.

### SetBillProjectAfterClosedFlagNil

`func (o *Project) SetBillProjectAfterClosedFlagNil(b bool)`

 SetBillProjectAfterClosedFlagNil sets the value for BillProjectAfterClosedFlag to be an explicit nil

### UnsetBillProjectAfterClosedFlag
`func (o *Project) UnsetBillProjectAfterClosedFlag()`

UnsetBillProjectAfterClosedFlag ensures that no value is present for BillProjectAfterClosedFlag, not even an explicit nil
### GetBillTime

`func (o *Project) GetBillTime() string`

GetBillTime returns the BillTime field if non-nil, zero value otherwise.

### GetBillTimeOk

`func (o *Project) GetBillTimeOk() (*string, bool)`

GetBillTimeOk returns a tuple with the BillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTime

`func (o *Project) SetBillTime(v string)`

SetBillTime sets BillTime field to given value.

### HasBillTime

`func (o *Project) HasBillTime() bool`

HasBillTime returns a boolean if a field has been set.

### SetBillTimeNil

`func (o *Project) SetBillTimeNil(b bool)`

 SetBillTimeNil sets the value for BillTime to be an explicit nil

### UnsetBillTime
`func (o *Project) UnsetBillTime()`

UnsetBillTime ensures that no value is present for BillTime, not even an explicit nil
### GetBillToCompany

`func (o *Project) GetBillToCompany() CompanyReference`

GetBillToCompany returns the BillToCompany field if non-nil, zero value otherwise.

### GetBillToCompanyOk

`func (o *Project) GetBillToCompanyOk() (*CompanyReference, bool)`

GetBillToCompanyOk returns a tuple with the BillToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToCompany

`func (o *Project) SetBillToCompany(v CompanyReference)`

SetBillToCompany sets BillToCompany field to given value.

### HasBillToCompany

`func (o *Project) HasBillToCompany() bool`

HasBillToCompany returns a boolean if a field has been set.

### GetBillToContact

`func (o *Project) GetBillToContact() ContactReference`

GetBillToContact returns the BillToContact field if non-nil, zero value otherwise.

### GetBillToContactOk

`func (o *Project) GetBillToContactOk() (*ContactReference, bool)`

GetBillToContactOk returns a tuple with the BillToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToContact

`func (o *Project) SetBillToContact(v ContactReference)`

SetBillToContact sets BillToContact field to given value.

### HasBillToContact

`func (o *Project) HasBillToContact() bool`

HasBillToContact returns a boolean if a field has been set.

### GetBillToSite

`func (o *Project) GetBillToSite() SiteReference`

GetBillToSite returns the BillToSite field if non-nil, zero value otherwise.

### GetBillToSiteOk

`func (o *Project) GetBillToSiteOk() (*SiteReference, bool)`

GetBillToSiteOk returns a tuple with the BillToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToSite

`func (o *Project) SetBillToSite(v SiteReference)`

SetBillToSite sets BillToSite field to given value.

### HasBillToSite

`func (o *Project) HasBillToSite() bool`

HasBillToSite returns a boolean if a field has been set.

### GetBillUnapprovedTimeAndExpense

`func (o *Project) GetBillUnapprovedTimeAndExpense() bool`

GetBillUnapprovedTimeAndExpense returns the BillUnapprovedTimeAndExpense field if non-nil, zero value otherwise.

### GetBillUnapprovedTimeAndExpenseOk

`func (o *Project) GetBillUnapprovedTimeAndExpenseOk() (*bool, bool)`

GetBillUnapprovedTimeAndExpenseOk returns a tuple with the BillUnapprovedTimeAndExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillUnapprovedTimeAndExpense

`func (o *Project) SetBillUnapprovedTimeAndExpense(v bool)`

SetBillUnapprovedTimeAndExpense sets BillUnapprovedTimeAndExpense field to given value.

### HasBillUnapprovedTimeAndExpense

`func (o *Project) HasBillUnapprovedTimeAndExpense() bool`

HasBillUnapprovedTimeAndExpense returns a boolean if a field has been set.

### SetBillUnapprovedTimeAndExpenseNil

`func (o *Project) SetBillUnapprovedTimeAndExpenseNil(b bool)`

 SetBillUnapprovedTimeAndExpenseNil sets the value for BillUnapprovedTimeAndExpense to be an explicit nil

### UnsetBillUnapprovedTimeAndExpense
`func (o *Project) UnsetBillUnapprovedTimeAndExpense()`

UnsetBillUnapprovedTimeAndExpense ensures that no value is present for BillUnapprovedTimeAndExpense, not even an explicit nil
### GetBoard

`func (o *Project) GetBoard() ProjectBoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *Project) GetBoardOk() (*ProjectBoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *Project) SetBoard(v ProjectBoardReference)`

SetBoard sets Board field to given value.


### GetBudgetAnalysis

`func (o *Project) GetBudgetAnalysis() string`

GetBudgetAnalysis returns the BudgetAnalysis field if non-nil, zero value otherwise.

### GetBudgetAnalysisOk

`func (o *Project) GetBudgetAnalysisOk() (*string, bool)`

GetBudgetAnalysisOk returns a tuple with the BudgetAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetAnalysis

`func (o *Project) SetBudgetAnalysis(v string)`

SetBudgetAnalysis sets BudgetAnalysis field to given value.

### HasBudgetAnalysis

`func (o *Project) HasBudgetAnalysis() bool`

HasBudgetAnalysis returns a boolean if a field has been set.

### SetBudgetAnalysisNil

`func (o *Project) SetBudgetAnalysisNil(b bool)`

 SetBudgetAnalysisNil sets the value for BudgetAnalysis to be an explicit nil

### UnsetBudgetAnalysis
`func (o *Project) UnsetBudgetAnalysis()`

UnsetBudgetAnalysis ensures that no value is present for BudgetAnalysis, not even an explicit nil
### GetBudgetFlag

`func (o *Project) GetBudgetFlag() bool`

GetBudgetFlag returns the BudgetFlag field if non-nil, zero value otherwise.

### GetBudgetFlagOk

`func (o *Project) GetBudgetFlagOk() (*bool, bool)`

GetBudgetFlagOk returns a tuple with the BudgetFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetFlag

`func (o *Project) SetBudgetFlag(v bool)`

SetBudgetFlag sets BudgetFlag field to given value.

### HasBudgetFlag

`func (o *Project) HasBudgetFlag() bool`

HasBudgetFlag returns a boolean if a field has been set.

### SetBudgetFlagNil

`func (o *Project) SetBudgetFlagNil(b bool)`

 SetBudgetFlagNil sets the value for BudgetFlag to be an explicit nil

### UnsetBudgetFlag
`func (o *Project) UnsetBudgetFlag()`

UnsetBudgetFlag ensures that no value is present for BudgetFlag, not even an explicit nil
### GetBudgetHours

`func (o *Project) GetBudgetHours() float64`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *Project) GetBudgetHoursOk() (*float64, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *Project) SetBudgetHours(v float64)`

SetBudgetHours sets BudgetHours field to given value.

### HasBudgetHours

`func (o *Project) HasBudgetHours() bool`

HasBudgetHours returns a boolean if a field has been set.

### SetBudgetHoursNil

`func (o *Project) SetBudgetHoursNil(b bool)`

 SetBudgetHoursNil sets the value for BudgetHours to be an explicit nil

### UnsetBudgetHours
`func (o *Project) UnsetBudgetHours()`

UnsetBudgetHours ensures that no value is present for BudgetHours, not even an explicit nil
### GetCompany

`func (o *Project) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Project) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Project) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetContact

`func (o *Project) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Project) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Project) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Project) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCustomerPO

`func (o *Project) GetCustomerPO() string`

GetCustomerPO returns the CustomerPO field if non-nil, zero value otherwise.

### GetCustomerPOOk

`func (o *Project) GetCustomerPOOk() (*string, bool)`

GetCustomerPOOk returns a tuple with the CustomerPO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPO

`func (o *Project) SetCustomerPO(v string)`

SetCustomerPO sets CustomerPO field to given value.

### HasCustomerPO

`func (o *Project) HasCustomerPO() bool`

HasCustomerPO returns a boolean if a field has been set.

### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Project) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCurrency

`func (o *Project) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Project) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Project) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Project) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDownpayment

`func (o *Project) GetDownpayment() float64`

GetDownpayment returns the Downpayment field if non-nil, zero value otherwise.

### GetDownpaymentOk

`func (o *Project) GetDownpaymentOk() (*float64, bool)`

GetDownpaymentOk returns a tuple with the Downpayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownpayment

`func (o *Project) SetDownpayment(v float64)`

SetDownpayment sets Downpayment field to given value.

### HasDownpayment

`func (o *Project) HasDownpayment() bool`

HasDownpayment returns a boolean if a field has been set.

### SetDownpaymentNil

`func (o *Project) SetDownpaymentNil(b bool)`

 SetDownpaymentNil sets the value for Downpayment to be an explicit nil

### UnsetDownpayment
`func (o *Project) UnsetDownpayment()`

UnsetDownpayment ensures that no value is present for Downpayment, not even an explicit nil
### GetEstimatedEnd

`func (o *Project) GetEstimatedEnd() time.Time`

GetEstimatedEnd returns the EstimatedEnd field if non-nil, zero value otherwise.

### GetEstimatedEndOk

`func (o *Project) GetEstimatedEndOk() (*time.Time, bool)`

GetEstimatedEndOk returns a tuple with the EstimatedEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedEnd

`func (o *Project) SetEstimatedEnd(v time.Time)`

SetEstimatedEnd sets EstimatedEnd field to given value.


### GetPercentComplete

`func (o *Project) GetPercentComplete() float64`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *Project) GetPercentCompleteOk() (*float64, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *Project) SetPercentComplete(v float64)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *Project) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentCompleteNil

`func (o *Project) SetPercentCompleteNil(b bool)`

 SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil

### UnsetPercentComplete
`func (o *Project) UnsetPercentComplete()`

UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
### GetEstimatedExpenseRevenue

`func (o *Project) GetEstimatedExpenseRevenue() float64`

GetEstimatedExpenseRevenue returns the EstimatedExpenseRevenue field if non-nil, zero value otherwise.

### GetEstimatedExpenseRevenueOk

`func (o *Project) GetEstimatedExpenseRevenueOk() (*float64, bool)`

GetEstimatedExpenseRevenueOk returns a tuple with the EstimatedExpenseRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedExpenseRevenue

`func (o *Project) SetEstimatedExpenseRevenue(v float64)`

SetEstimatedExpenseRevenue sets EstimatedExpenseRevenue field to given value.

### HasEstimatedExpenseRevenue

`func (o *Project) HasEstimatedExpenseRevenue() bool`

HasEstimatedExpenseRevenue returns a boolean if a field has been set.

### SetEstimatedExpenseRevenueNil

`func (o *Project) SetEstimatedExpenseRevenueNil(b bool)`

 SetEstimatedExpenseRevenueNil sets the value for EstimatedExpenseRevenue to be an explicit nil

### UnsetEstimatedExpenseRevenue
`func (o *Project) UnsetEstimatedExpenseRevenue()`

UnsetEstimatedExpenseRevenue ensures that no value is present for EstimatedExpenseRevenue, not even an explicit nil
### GetEstimatedHours

`func (o *Project) GetEstimatedHours() float64`

GetEstimatedHours returns the EstimatedHours field if non-nil, zero value otherwise.

### GetEstimatedHoursOk

`func (o *Project) GetEstimatedHoursOk() (*float64, bool)`

GetEstimatedHoursOk returns a tuple with the EstimatedHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedHours

`func (o *Project) SetEstimatedHours(v float64)`

SetEstimatedHours sets EstimatedHours field to given value.

### HasEstimatedHours

`func (o *Project) HasEstimatedHours() bool`

HasEstimatedHours returns a boolean if a field has been set.

### SetEstimatedHoursNil

`func (o *Project) SetEstimatedHoursNil(b bool)`

 SetEstimatedHoursNil sets the value for EstimatedHours to be an explicit nil

### UnsetEstimatedHours
`func (o *Project) UnsetEstimatedHours()`

UnsetEstimatedHours ensures that no value is present for EstimatedHours, not even an explicit nil
### GetEstimatedProductRevenue

`func (o *Project) GetEstimatedProductRevenue() float64`

GetEstimatedProductRevenue returns the EstimatedProductRevenue field if non-nil, zero value otherwise.

### GetEstimatedProductRevenueOk

`func (o *Project) GetEstimatedProductRevenueOk() (*float64, bool)`

GetEstimatedProductRevenueOk returns a tuple with the EstimatedProductRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedProductRevenue

`func (o *Project) SetEstimatedProductRevenue(v float64)`

SetEstimatedProductRevenue sets EstimatedProductRevenue field to given value.

### HasEstimatedProductRevenue

`func (o *Project) HasEstimatedProductRevenue() bool`

HasEstimatedProductRevenue returns a boolean if a field has been set.

### SetEstimatedProductRevenueNil

`func (o *Project) SetEstimatedProductRevenueNil(b bool)`

 SetEstimatedProductRevenueNil sets the value for EstimatedProductRevenue to be an explicit nil

### UnsetEstimatedProductRevenue
`func (o *Project) UnsetEstimatedProductRevenue()`

UnsetEstimatedProductRevenue ensures that no value is present for EstimatedProductRevenue, not even an explicit nil
### GetEstimatedStart

`func (o *Project) GetEstimatedStart() time.Time`

GetEstimatedStart returns the EstimatedStart field if non-nil, zero value otherwise.

### GetEstimatedStartOk

`func (o *Project) GetEstimatedStartOk() (*time.Time, bool)`

GetEstimatedStartOk returns a tuple with the EstimatedStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStart

`func (o *Project) SetEstimatedStart(v time.Time)`

SetEstimatedStart sets EstimatedStart field to given value.


### GetEstimatedTimeRevenue

`func (o *Project) GetEstimatedTimeRevenue() float64`

GetEstimatedTimeRevenue returns the EstimatedTimeRevenue field if non-nil, zero value otherwise.

### GetEstimatedTimeRevenueOk

`func (o *Project) GetEstimatedTimeRevenueOk() (*float64, bool)`

GetEstimatedTimeRevenueOk returns a tuple with the EstimatedTimeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeRevenue

`func (o *Project) SetEstimatedTimeRevenue(v float64)`

SetEstimatedTimeRevenue sets EstimatedTimeRevenue field to given value.

### HasEstimatedTimeRevenue

`func (o *Project) HasEstimatedTimeRevenue() bool`

HasEstimatedTimeRevenue returns a boolean if a field has been set.

### SetEstimatedTimeRevenueNil

`func (o *Project) SetEstimatedTimeRevenueNil(b bool)`

 SetEstimatedTimeRevenueNil sets the value for EstimatedTimeRevenue to be an explicit nil

### UnsetEstimatedTimeRevenue
`func (o *Project) UnsetEstimatedTimeRevenue()`

UnsetEstimatedTimeRevenue ensures that no value is present for EstimatedTimeRevenue, not even an explicit nil
### GetExpenseApprover

`func (o *Project) GetExpenseApprover() MemberReference`

GetExpenseApprover returns the ExpenseApprover field if non-nil, zero value otherwise.

### GetExpenseApproverOk

`func (o *Project) GetExpenseApproverOk() (*MemberReference, bool)`

GetExpenseApproverOk returns a tuple with the ExpenseApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApprover

`func (o *Project) SetExpenseApprover(v MemberReference)`

SetExpenseApprover sets ExpenseApprover field to given value.

### HasExpenseApprover

`func (o *Project) HasExpenseApprover() bool`

HasExpenseApprover returns a boolean if a field has been set.

### GetIncludeDependenciesFlag

`func (o *Project) GetIncludeDependenciesFlag() bool`

GetIncludeDependenciesFlag returns the IncludeDependenciesFlag field if non-nil, zero value otherwise.

### GetIncludeDependenciesFlagOk

`func (o *Project) GetIncludeDependenciesFlagOk() (*bool, bool)`

GetIncludeDependenciesFlagOk returns a tuple with the IncludeDependenciesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDependenciesFlag

`func (o *Project) SetIncludeDependenciesFlag(v bool)`

SetIncludeDependenciesFlag sets IncludeDependenciesFlag field to given value.

### HasIncludeDependenciesFlag

`func (o *Project) HasIncludeDependenciesFlag() bool`

HasIncludeDependenciesFlag returns a boolean if a field has been set.

### SetIncludeDependenciesFlagNil

`func (o *Project) SetIncludeDependenciesFlagNil(b bool)`

 SetIncludeDependenciesFlagNil sets the value for IncludeDependenciesFlag to be an explicit nil

### UnsetIncludeDependenciesFlag
`func (o *Project) UnsetIncludeDependenciesFlag()`

UnsetIncludeDependenciesFlag ensures that no value is present for IncludeDependenciesFlag, not even an explicit nil
### GetIncludeEstimatesFlag

`func (o *Project) GetIncludeEstimatesFlag() bool`

GetIncludeEstimatesFlag returns the IncludeEstimatesFlag field if non-nil, zero value otherwise.

### GetIncludeEstimatesFlagOk

`func (o *Project) GetIncludeEstimatesFlagOk() (*bool, bool)`

GetIncludeEstimatesFlagOk returns a tuple with the IncludeEstimatesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEstimatesFlag

`func (o *Project) SetIncludeEstimatesFlag(v bool)`

SetIncludeEstimatesFlag sets IncludeEstimatesFlag field to given value.

### HasIncludeEstimatesFlag

`func (o *Project) HasIncludeEstimatesFlag() bool`

HasIncludeEstimatesFlag returns a boolean if a field has been set.

### SetIncludeEstimatesFlagNil

`func (o *Project) SetIncludeEstimatesFlagNil(b bool)`

 SetIncludeEstimatesFlagNil sets the value for IncludeEstimatesFlag to be an explicit nil

### UnsetIncludeEstimatesFlag
`func (o *Project) UnsetIncludeEstimatesFlag()`

UnsetIncludeEstimatesFlag ensures that no value is present for IncludeEstimatesFlag, not even an explicit nil
### GetLocation

`func (o *Project) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Project) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Project) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Project) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *Project) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Project) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Project) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Project) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetManager

`func (o *Project) GetManager() MemberReference`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *Project) GetManagerOk() (*MemberReference, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *Project) SetManager(v MemberReference)`

SetManager sets Manager field to given value.

### HasManager

`func (o *Project) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.


### GetOpportunity

`func (o *Project) GetOpportunity() OpportunityReference`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *Project) GetOpportunityOk() (*OpportunityReference, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *Project) SetOpportunity(v OpportunityReference)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *Project) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### GetProjectTemplateId

`func (o *Project) GetProjectTemplateId() int32`

GetProjectTemplateId returns the ProjectTemplateId field if non-nil, zero value otherwise.

### GetProjectTemplateIdOk

`func (o *Project) GetProjectTemplateIdOk() (*int32, bool)`

GetProjectTemplateIdOk returns a tuple with the ProjectTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTemplateId

`func (o *Project) SetProjectTemplateId(v int32)`

SetProjectTemplateId sets ProjectTemplateId field to given value.

### HasProjectTemplateId

`func (o *Project) HasProjectTemplateId() bool`

HasProjectTemplateId returns a boolean if a field has been set.

### SetProjectTemplateIdNil

`func (o *Project) SetProjectTemplateIdNil(b bool)`

 SetProjectTemplateIdNil sets the value for ProjectTemplateId to be an explicit nil

### UnsetProjectTemplateId
`func (o *Project) UnsetProjectTemplateId()`

UnsetProjectTemplateId ensures that no value is present for ProjectTemplateId, not even an explicit nil
### GetRestrictDownPaymentFlag

`func (o *Project) GetRestrictDownPaymentFlag() bool`

GetRestrictDownPaymentFlag returns the RestrictDownPaymentFlag field if non-nil, zero value otherwise.

### GetRestrictDownPaymentFlagOk

`func (o *Project) GetRestrictDownPaymentFlagOk() (*bool, bool)`

GetRestrictDownPaymentFlagOk returns a tuple with the RestrictDownPaymentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDownPaymentFlag

`func (o *Project) SetRestrictDownPaymentFlag(v bool)`

SetRestrictDownPaymentFlag sets RestrictDownPaymentFlag field to given value.

### HasRestrictDownPaymentFlag

`func (o *Project) HasRestrictDownPaymentFlag() bool`

HasRestrictDownPaymentFlag returns a boolean if a field has been set.

### SetRestrictDownPaymentFlagNil

`func (o *Project) SetRestrictDownPaymentFlagNil(b bool)`

 SetRestrictDownPaymentFlagNil sets the value for RestrictDownPaymentFlag to be an explicit nil

### UnsetRestrictDownPaymentFlag
`func (o *Project) UnsetRestrictDownPaymentFlag()`

UnsetRestrictDownPaymentFlag ensures that no value is present for RestrictDownPaymentFlag, not even an explicit nil
### GetScheduledEnd

`func (o *Project) GetScheduledEnd() time.Time`

GetScheduledEnd returns the ScheduledEnd field if non-nil, zero value otherwise.

### GetScheduledEndOk

`func (o *Project) GetScheduledEndOk() (*time.Time, bool)`

GetScheduledEndOk returns a tuple with the ScheduledEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEnd

`func (o *Project) SetScheduledEnd(v time.Time)`

SetScheduledEnd sets ScheduledEnd field to given value.

### HasScheduledEnd

`func (o *Project) HasScheduledEnd() bool`

HasScheduledEnd returns a boolean if a field has been set.

### GetScheduledHours

`func (o *Project) GetScheduledHours() float64`

GetScheduledHours returns the ScheduledHours field if non-nil, zero value otherwise.

### GetScheduledHoursOk

`func (o *Project) GetScheduledHoursOk() (*float64, bool)`

GetScheduledHoursOk returns a tuple with the ScheduledHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledHours

`func (o *Project) SetScheduledHours(v float64)`

SetScheduledHours sets ScheduledHours field to given value.

### HasScheduledHours

`func (o *Project) HasScheduledHours() bool`

HasScheduledHours returns a boolean if a field has been set.

### SetScheduledHoursNil

`func (o *Project) SetScheduledHoursNil(b bool)`

 SetScheduledHoursNil sets the value for ScheduledHours to be an explicit nil

### UnsetScheduledHours
`func (o *Project) UnsetScheduledHours()`

UnsetScheduledHours ensures that no value is present for ScheduledHours, not even an explicit nil
### GetScheduledStart

`func (o *Project) GetScheduledStart() time.Time`

GetScheduledStart returns the ScheduledStart field if non-nil, zero value otherwise.

### GetScheduledStartOk

`func (o *Project) GetScheduledStartOk() (*time.Time, bool)`

GetScheduledStartOk returns a tuple with the ScheduledStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStart

`func (o *Project) SetScheduledStart(v time.Time)`

SetScheduledStart sets ScheduledStart field to given value.

### HasScheduledStart

`func (o *Project) HasScheduledStart() bool`

HasScheduledStart returns a boolean if a field has been set.

### GetShipToCompany

`func (o *Project) GetShipToCompany() CompanyReference`

GetShipToCompany returns the ShipToCompany field if non-nil, zero value otherwise.

### GetShipToCompanyOk

`func (o *Project) GetShipToCompanyOk() (*CompanyReference, bool)`

GetShipToCompanyOk returns a tuple with the ShipToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompany

`func (o *Project) SetShipToCompany(v CompanyReference)`

SetShipToCompany sets ShipToCompany field to given value.

### HasShipToCompany

`func (o *Project) HasShipToCompany() bool`

HasShipToCompany returns a boolean if a field has been set.

### GetShipToContact

`func (o *Project) GetShipToContact() ContactReference`

GetShipToContact returns the ShipToContact field if non-nil, zero value otherwise.

### GetShipToContactOk

`func (o *Project) GetShipToContactOk() (*ContactReference, bool)`

GetShipToContactOk returns a tuple with the ShipToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToContact

`func (o *Project) SetShipToContact(v ContactReference)`

SetShipToContact sets ShipToContact field to given value.

### HasShipToContact

`func (o *Project) HasShipToContact() bool`

HasShipToContact returns a boolean if a field has been set.

### GetShipToSite

`func (o *Project) GetShipToSite() SiteReference`

GetShipToSite returns the ShipToSite field if non-nil, zero value otherwise.

### GetShipToSiteOk

`func (o *Project) GetShipToSiteOk() (*SiteReference, bool)`

GetShipToSiteOk returns a tuple with the ShipToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToSite

`func (o *Project) SetShipToSite(v SiteReference)`

SetShipToSite sets ShipToSite field to given value.

### HasShipToSite

`func (o *Project) HasShipToSite() bool`

HasShipToSite returns a boolean if a field has been set.

### GetSite

`func (o *Project) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Project) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Project) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *Project) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetStatus

`func (o *Project) GetStatus() ProjectStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Project) GetStatusOk() (*ProjectStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Project) SetStatus(v ProjectStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Project) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClosedFlag

`func (o *Project) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *Project) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *Project) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *Project) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### GetTimeApprover

`func (o *Project) GetTimeApprover() MemberReference`

GetTimeApprover returns the TimeApprover field if non-nil, zero value otherwise.

### GetTimeApproverOk

`func (o *Project) GetTimeApproverOk() (*MemberReference, bool)`

GetTimeApproverOk returns a tuple with the TimeApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeApprover

`func (o *Project) SetTimeApprover(v MemberReference)`

SetTimeApprover sets TimeApprover field to given value.

### HasTimeApprover

`func (o *Project) HasTimeApprover() bool`

HasTimeApprover returns a boolean if a field has been set.

### GetType

`func (o *Project) GetType() ProjectTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Project) GetTypeOk() (*ProjectTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Project) SetType(v ProjectTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *Project) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDoNotDisplayInPortalFlag

`func (o *Project) GetDoNotDisplayInPortalFlag() bool`

GetDoNotDisplayInPortalFlag returns the DoNotDisplayInPortalFlag field if non-nil, zero value otherwise.

### GetDoNotDisplayInPortalFlagOk

`func (o *Project) GetDoNotDisplayInPortalFlagOk() (*bool, bool)`

GetDoNotDisplayInPortalFlagOk returns a tuple with the DoNotDisplayInPortalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisplayInPortalFlag

`func (o *Project) SetDoNotDisplayInPortalFlag(v bool)`

SetDoNotDisplayInPortalFlag sets DoNotDisplayInPortalFlag field to given value.

### HasDoNotDisplayInPortalFlag

`func (o *Project) HasDoNotDisplayInPortalFlag() bool`

HasDoNotDisplayInPortalFlag returns a boolean if a field has been set.

### SetDoNotDisplayInPortalFlagNil

`func (o *Project) SetDoNotDisplayInPortalFlagNil(b bool)`

 SetDoNotDisplayInPortalFlagNil sets the value for DoNotDisplayInPortalFlag to be an explicit nil

### UnsetDoNotDisplayInPortalFlag
`func (o *Project) UnsetDoNotDisplayInPortalFlag()`

UnsetDoNotDisplayInPortalFlag ensures that no value is present for DoNotDisplayInPortalFlag, not even an explicit nil
### GetBillingStartDate

`func (o *Project) GetBillingStartDate() time.Time`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *Project) GetBillingStartDateOk() (*time.Time, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *Project) SetBillingStartDate(v time.Time)`

SetBillingStartDate sets BillingStartDate field to given value.

### HasBillingStartDate

`func (o *Project) HasBillingStartDate() bool`

HasBillingStartDate returns a boolean if a field has been set.

### GetPoAmount

`func (o *Project) GetPoAmount() float64`

GetPoAmount returns the PoAmount field if non-nil, zero value otherwise.

### GetPoAmountOk

`func (o *Project) GetPoAmountOk() (*float64, bool)`

GetPoAmountOk returns a tuple with the PoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoAmount

`func (o *Project) SetPoAmount(v float64)`

SetPoAmount sets PoAmount field to given value.

### HasPoAmount

`func (o *Project) HasPoAmount() bool`

HasPoAmount returns a boolean if a field has been set.

### SetPoAmountNil

`func (o *Project) SetPoAmountNil(b bool)`

 SetPoAmountNil sets the value for PoAmount to be an explicit nil

### UnsetPoAmount
`func (o *Project) UnsetPoAmount()`

UnsetPoAmount ensures that no value is present for PoAmount, not even an explicit nil
### GetEstimatedTimeCost

`func (o *Project) GetEstimatedTimeCost() float64`

GetEstimatedTimeCost returns the EstimatedTimeCost field if non-nil, zero value otherwise.

### GetEstimatedTimeCostOk

`func (o *Project) GetEstimatedTimeCostOk() (*float64, bool)`

GetEstimatedTimeCostOk returns a tuple with the EstimatedTimeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeCost

`func (o *Project) SetEstimatedTimeCost(v float64)`

SetEstimatedTimeCost sets EstimatedTimeCost field to given value.

### HasEstimatedTimeCost

`func (o *Project) HasEstimatedTimeCost() bool`

HasEstimatedTimeCost returns a boolean if a field has been set.

### SetEstimatedTimeCostNil

`func (o *Project) SetEstimatedTimeCostNil(b bool)`

 SetEstimatedTimeCostNil sets the value for EstimatedTimeCost to be an explicit nil

### UnsetEstimatedTimeCost
`func (o *Project) UnsetEstimatedTimeCost()`

UnsetEstimatedTimeCost ensures that no value is present for EstimatedTimeCost, not even an explicit nil
### GetEstimatedExpenseCost

`func (o *Project) GetEstimatedExpenseCost() float64`

GetEstimatedExpenseCost returns the EstimatedExpenseCost field if non-nil, zero value otherwise.

### GetEstimatedExpenseCostOk

`func (o *Project) GetEstimatedExpenseCostOk() (*float64, bool)`

GetEstimatedExpenseCostOk returns a tuple with the EstimatedExpenseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedExpenseCost

`func (o *Project) SetEstimatedExpenseCost(v float64)`

SetEstimatedExpenseCost sets EstimatedExpenseCost field to given value.

### HasEstimatedExpenseCost

`func (o *Project) HasEstimatedExpenseCost() bool`

HasEstimatedExpenseCost returns a boolean if a field has been set.

### SetEstimatedExpenseCostNil

`func (o *Project) SetEstimatedExpenseCostNil(b bool)`

 SetEstimatedExpenseCostNil sets the value for EstimatedExpenseCost to be an explicit nil

### UnsetEstimatedExpenseCost
`func (o *Project) UnsetEstimatedExpenseCost()`

UnsetEstimatedExpenseCost ensures that no value is present for EstimatedExpenseCost, not even an explicit nil
### GetEstimatedProductCost

`func (o *Project) GetEstimatedProductCost() float64`

GetEstimatedProductCost returns the EstimatedProductCost field if non-nil, zero value otherwise.

### GetEstimatedProductCostOk

`func (o *Project) GetEstimatedProductCostOk() (*float64, bool)`

GetEstimatedProductCostOk returns a tuple with the EstimatedProductCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedProductCost

`func (o *Project) SetEstimatedProductCost(v float64)`

SetEstimatedProductCost sets EstimatedProductCost field to given value.

### HasEstimatedProductCost

`func (o *Project) HasEstimatedProductCost() bool`

HasEstimatedProductCost returns a boolean if a field has been set.

### SetEstimatedProductCostNil

`func (o *Project) SetEstimatedProductCostNil(b bool)`

 SetEstimatedProductCostNil sets the value for EstimatedProductCost to be an explicit nil

### UnsetEstimatedProductCost
`func (o *Project) UnsetEstimatedProductCost()`

UnsetEstimatedProductCost ensures that no value is present for EstimatedProductCost, not even an explicit nil
### GetTaxCode

`func (o *Project) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *Project) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *Project) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *Project) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetCompanyLocation

`func (o *Project) GetCompanyLocation() SystemLocationReference`

GetCompanyLocation returns the CompanyLocation field if non-nil, zero value otherwise.

### GetCompanyLocationOk

`func (o *Project) GetCompanyLocationOk() (*SystemLocationReference, bool)`

GetCompanyLocationOk returns a tuple with the CompanyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLocation

`func (o *Project) SetCompanyLocation(v SystemLocationReference)`

SetCompanyLocation sets CompanyLocation field to given value.

### HasCompanyLocation

`func (o *Project) HasCompanyLocation() bool`

HasCompanyLocation returns a boolean if a field has been set.

### GetInfo

`func (o *Project) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Project) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Project) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Project) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *Project) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Project) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Project) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Project) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


