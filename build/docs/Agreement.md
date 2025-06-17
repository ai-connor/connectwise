# Agreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 100; | 
**Type** | [**AgreementTypeReference**](AgreementTypeReference.md) |  | 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
**Contact** | [**ContactReference**](ContactReference.md) |  | 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**SubContractCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**SubContractContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**ParentAgreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**CustomerPO** | Pointer to **string** |  Max length: 50; | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**RestrictLocationFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictDepartmentFlag** | Pointer to **NullableBool** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**NoEndingDateFlag** | Pointer to **NullableBool** |  | [optional] 
**Opportunity** | Pointer to [**OpportunityReference**](OpportunityReference.md) |  | [optional] 
**CancelledFlag** | Pointer to **NullableBool** |  | [optional] 
**DateCancelled** | Pointer to **time.Time** |  | [optional] 
**ReasonCancelled** | Pointer to **string** |  Max length: 100; | [optional] 
**Sla** | Pointer to [**SLAReference**](SLAReference.md) |  | [optional] 
**WorkOrder** | Pointer to **string** |  Max length: 20; | [optional] 
**InternalNotes** | Pointer to **string** |  | [optional] 
**ApplicationUnits** | Pointer to **NullableString** |  | [optional] 
**ApplicationLimit** | Pointer to **NullableFloat64** |  | [optional] 
**ApplicationCycle** | Pointer to **NullableString** |  | [optional] 
**ApplicationUnlimitedFlag** | Pointer to **NullableBool** |  | [optional] 
**OneTimeFlag** | Pointer to **NullableBool** |  | [optional] 
**CoverAgreementTime** | Pointer to **NullableBool** |  | [optional] 
**CoverAgreementProduct** | Pointer to **NullableBool** |  | [optional] 
**CoverAgreementExpense** | Pointer to **NullableBool** |  | [optional] 
**CoverSalesTax** | Pointer to **NullableBool** |  | [optional] 
**CarryOverUnused** | Pointer to **NullableBool** |  | [optional] 
**AllowOverruns** | Pointer to **NullableBool** |  | [optional] 
**ExpiredDays** | Pointer to **NullableInt32** |  | [optional] 
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**ExpireWhenZero** | Pointer to **NullableBool** |  | [optional] 
**ChargeToFirm** | Pointer to **NullableBool** |  | [optional] 
**EmployeeCompRate** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**EmployeeCompNotExceed** | Pointer to **NullableString** |  | [optional] 
**CompHourlyRate** | Pointer to **NullableFloat64** |  | [optional] 
**CompLimitAmount** | Pointer to **NullableFloat64** |  | [optional] 
**BillingCycle** | Pointer to [**BillingCycleReference**](BillingCycleReference.md) |  | [optional] 
**BillOneTimeFlag** | Pointer to **NullableBool** |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**InvoicingCycle** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**BillToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**BillToContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**BillToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**BillAmount** | Pointer to **NullableFloat64** |  | [optional] 
**Taxable** | Pointer to **NullableBool** |  | [optional] 
**ProrateFirstBill** | Pointer to **NullableFloat64** |  | [optional] 
**BillStartDate** | Pointer to **time.Time** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**RestrictDownPayment** | Pointer to **NullableBool** |  | [optional] 
**ProrateFlag** | Pointer to **NullableBool** |  | [optional] 
**InvoiceProratedAdditionsFlag** | Pointer to **NullableBool** |  | [optional] 
**InvoiceDescription** | Pointer to **string** |  | [optional] 
**TopComment** | Pointer to **NullableBool** |  | [optional] 
**BottomComment** | Pointer to **NullableBool** |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**ProjectType** | Pointer to [**ProjectTypeReference**](ProjectTypeReference.md) |  | [optional] 
**InvoiceTemplate** | Pointer to [**InvoiceTemplateReference**](InvoiceTemplateReference.md) |  | [optional] 
**BillTime** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**BillExpenses** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**BillProducts** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**BillableTimeInvoice** | Pointer to **NullableBool** |  | [optional] 
**BillableExpenseInvoice** | Pointer to **NullableBool** |  | [optional] 
**BillableProductInvoice** | Pointer to **NullableBool** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**PeriodType** | Pointer to **NullableString** |  | [optional] 
**AutoInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**NextInvoiceDate** | Pointer to **string** |  | [optional] 
**CompanyLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**ShipToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ShipToContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**ShipToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**AgreementStatus** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewAgreement

`func NewAgreement(name string, type_ AgreementTypeReference, company CompanyReference, contact ContactReference, ) *Agreement`

NewAgreement instantiates a new Agreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementWithDefaults

`func NewAgreementWithDefaults() *Agreement`

NewAgreementWithDefaults instantiates a new Agreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Agreement) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agreement) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agreement) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Agreement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Agreement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agreement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agreement) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Agreement) GetType() AgreementTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Agreement) GetTypeOk() (*AgreementTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Agreement) SetType(v AgreementTypeReference)`

SetType sets Type field to given value.


### GetCompany

`func (o *Agreement) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Agreement) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Agreement) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetContact

`func (o *Agreement) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Agreement) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Agreement) SetContact(v ContactReference)`

SetContact sets Contact field to given value.


### GetSite

`func (o *Agreement) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Agreement) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Agreement) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *Agreement) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSubContractCompany

`func (o *Agreement) GetSubContractCompany() CompanyReference`

GetSubContractCompany returns the SubContractCompany field if non-nil, zero value otherwise.

### GetSubContractCompanyOk

`func (o *Agreement) GetSubContractCompanyOk() (*CompanyReference, bool)`

GetSubContractCompanyOk returns a tuple with the SubContractCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubContractCompany

`func (o *Agreement) SetSubContractCompany(v CompanyReference)`

SetSubContractCompany sets SubContractCompany field to given value.

### HasSubContractCompany

`func (o *Agreement) HasSubContractCompany() bool`

HasSubContractCompany returns a boolean if a field has been set.

### GetSubContractContact

`func (o *Agreement) GetSubContractContact() ContactReference`

GetSubContractContact returns the SubContractContact field if non-nil, zero value otherwise.

### GetSubContractContactOk

`func (o *Agreement) GetSubContractContactOk() (*ContactReference, bool)`

GetSubContractContactOk returns a tuple with the SubContractContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubContractContact

`func (o *Agreement) SetSubContractContact(v ContactReference)`

SetSubContractContact sets SubContractContact field to given value.

### HasSubContractContact

`func (o *Agreement) HasSubContractContact() bool`

HasSubContractContact returns a boolean if a field has been set.

### GetParentAgreement

`func (o *Agreement) GetParentAgreement() AgreementReference`

GetParentAgreement returns the ParentAgreement field if non-nil, zero value otherwise.

### GetParentAgreementOk

`func (o *Agreement) GetParentAgreementOk() (*AgreementReference, bool)`

GetParentAgreementOk returns a tuple with the ParentAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAgreement

`func (o *Agreement) SetParentAgreement(v AgreementReference)`

SetParentAgreement sets ParentAgreement field to given value.

### HasParentAgreement

`func (o *Agreement) HasParentAgreement() bool`

HasParentAgreement returns a boolean if a field has been set.

### GetCustomerPO

`func (o *Agreement) GetCustomerPO() string`

GetCustomerPO returns the CustomerPO field if non-nil, zero value otherwise.

### GetCustomerPOOk

`func (o *Agreement) GetCustomerPOOk() (*string, bool)`

GetCustomerPOOk returns a tuple with the CustomerPO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPO

`func (o *Agreement) SetCustomerPO(v string)`

SetCustomerPO sets CustomerPO field to given value.

### HasCustomerPO

`func (o *Agreement) HasCustomerPO() bool`

HasCustomerPO returns a boolean if a field has been set.

### GetLocation

`func (o *Agreement) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Agreement) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Agreement) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Agreement) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *Agreement) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Agreement) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Agreement) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Agreement) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetRestrictLocationFlag

`func (o *Agreement) GetRestrictLocationFlag() bool`

GetRestrictLocationFlag returns the RestrictLocationFlag field if non-nil, zero value otherwise.

### GetRestrictLocationFlagOk

`func (o *Agreement) GetRestrictLocationFlagOk() (*bool, bool)`

GetRestrictLocationFlagOk returns a tuple with the RestrictLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictLocationFlag

`func (o *Agreement) SetRestrictLocationFlag(v bool)`

SetRestrictLocationFlag sets RestrictLocationFlag field to given value.

### HasRestrictLocationFlag

`func (o *Agreement) HasRestrictLocationFlag() bool`

HasRestrictLocationFlag returns a boolean if a field has been set.

### SetRestrictLocationFlagNil

`func (o *Agreement) SetRestrictLocationFlagNil(b bool)`

 SetRestrictLocationFlagNil sets the value for RestrictLocationFlag to be an explicit nil

### UnsetRestrictLocationFlag
`func (o *Agreement) UnsetRestrictLocationFlag()`

UnsetRestrictLocationFlag ensures that no value is present for RestrictLocationFlag, not even an explicit nil
### GetRestrictDepartmentFlag

`func (o *Agreement) GetRestrictDepartmentFlag() bool`

GetRestrictDepartmentFlag returns the RestrictDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictDepartmentFlagOk

`func (o *Agreement) GetRestrictDepartmentFlagOk() (*bool, bool)`

GetRestrictDepartmentFlagOk returns a tuple with the RestrictDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDepartmentFlag

`func (o *Agreement) SetRestrictDepartmentFlag(v bool)`

SetRestrictDepartmentFlag sets RestrictDepartmentFlag field to given value.

### HasRestrictDepartmentFlag

`func (o *Agreement) HasRestrictDepartmentFlag() bool`

HasRestrictDepartmentFlag returns a boolean if a field has been set.

### SetRestrictDepartmentFlagNil

`func (o *Agreement) SetRestrictDepartmentFlagNil(b bool)`

 SetRestrictDepartmentFlagNil sets the value for RestrictDepartmentFlag to be an explicit nil

### UnsetRestrictDepartmentFlag
`func (o *Agreement) UnsetRestrictDepartmentFlag()`

UnsetRestrictDepartmentFlag ensures that no value is present for RestrictDepartmentFlag, not even an explicit nil
### GetStartDate

`func (o *Agreement) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Agreement) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Agreement) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Agreement) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Agreement) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Agreement) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Agreement) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Agreement) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNoEndingDateFlag

`func (o *Agreement) GetNoEndingDateFlag() bool`

GetNoEndingDateFlag returns the NoEndingDateFlag field if non-nil, zero value otherwise.

### GetNoEndingDateFlagOk

`func (o *Agreement) GetNoEndingDateFlagOk() (*bool, bool)`

GetNoEndingDateFlagOk returns a tuple with the NoEndingDateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoEndingDateFlag

`func (o *Agreement) SetNoEndingDateFlag(v bool)`

SetNoEndingDateFlag sets NoEndingDateFlag field to given value.

### HasNoEndingDateFlag

`func (o *Agreement) HasNoEndingDateFlag() bool`

HasNoEndingDateFlag returns a boolean if a field has been set.

### SetNoEndingDateFlagNil

`func (o *Agreement) SetNoEndingDateFlagNil(b bool)`

 SetNoEndingDateFlagNil sets the value for NoEndingDateFlag to be an explicit nil

### UnsetNoEndingDateFlag
`func (o *Agreement) UnsetNoEndingDateFlag()`

UnsetNoEndingDateFlag ensures that no value is present for NoEndingDateFlag, not even an explicit nil
### GetOpportunity

`func (o *Agreement) GetOpportunity() OpportunityReference`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *Agreement) GetOpportunityOk() (*OpportunityReference, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *Agreement) SetOpportunity(v OpportunityReference)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *Agreement) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### GetCancelledFlag

`func (o *Agreement) GetCancelledFlag() bool`

GetCancelledFlag returns the CancelledFlag field if non-nil, zero value otherwise.

### GetCancelledFlagOk

`func (o *Agreement) GetCancelledFlagOk() (*bool, bool)`

GetCancelledFlagOk returns a tuple with the CancelledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledFlag

`func (o *Agreement) SetCancelledFlag(v bool)`

SetCancelledFlag sets CancelledFlag field to given value.

### HasCancelledFlag

`func (o *Agreement) HasCancelledFlag() bool`

HasCancelledFlag returns a boolean if a field has been set.

### SetCancelledFlagNil

`func (o *Agreement) SetCancelledFlagNil(b bool)`

 SetCancelledFlagNil sets the value for CancelledFlag to be an explicit nil

### UnsetCancelledFlag
`func (o *Agreement) UnsetCancelledFlag()`

UnsetCancelledFlag ensures that no value is present for CancelledFlag, not even an explicit nil
### GetDateCancelled

`func (o *Agreement) GetDateCancelled() time.Time`

GetDateCancelled returns the DateCancelled field if non-nil, zero value otherwise.

### GetDateCancelledOk

`func (o *Agreement) GetDateCancelledOk() (*time.Time, bool)`

GetDateCancelledOk returns a tuple with the DateCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCancelled

`func (o *Agreement) SetDateCancelled(v time.Time)`

SetDateCancelled sets DateCancelled field to given value.

### HasDateCancelled

`func (o *Agreement) HasDateCancelled() bool`

HasDateCancelled returns a boolean if a field has been set.

### GetReasonCancelled

`func (o *Agreement) GetReasonCancelled() string`

GetReasonCancelled returns the ReasonCancelled field if non-nil, zero value otherwise.

### GetReasonCancelledOk

`func (o *Agreement) GetReasonCancelledOk() (*string, bool)`

GetReasonCancelledOk returns a tuple with the ReasonCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCancelled

`func (o *Agreement) SetReasonCancelled(v string)`

SetReasonCancelled sets ReasonCancelled field to given value.

### HasReasonCancelled

`func (o *Agreement) HasReasonCancelled() bool`

HasReasonCancelled returns a boolean if a field has been set.

### GetSla

`func (o *Agreement) GetSla() SLAReference`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *Agreement) GetSlaOk() (*SLAReference, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *Agreement) SetSla(v SLAReference)`

SetSla sets Sla field to given value.

### HasSla

`func (o *Agreement) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetWorkOrder

`func (o *Agreement) GetWorkOrder() string`

GetWorkOrder returns the WorkOrder field if non-nil, zero value otherwise.

### GetWorkOrderOk

`func (o *Agreement) GetWorkOrderOk() (*string, bool)`

GetWorkOrderOk returns a tuple with the WorkOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkOrder

`func (o *Agreement) SetWorkOrder(v string)`

SetWorkOrder sets WorkOrder field to given value.

### HasWorkOrder

`func (o *Agreement) HasWorkOrder() bool`

HasWorkOrder returns a boolean if a field has been set.

### GetInternalNotes

`func (o *Agreement) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *Agreement) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *Agreement) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *Agreement) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetApplicationUnits

`func (o *Agreement) GetApplicationUnits() string`

GetApplicationUnits returns the ApplicationUnits field if non-nil, zero value otherwise.

### GetApplicationUnitsOk

`func (o *Agreement) GetApplicationUnitsOk() (*string, bool)`

GetApplicationUnitsOk returns a tuple with the ApplicationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUnits

`func (o *Agreement) SetApplicationUnits(v string)`

SetApplicationUnits sets ApplicationUnits field to given value.

### HasApplicationUnits

`func (o *Agreement) HasApplicationUnits() bool`

HasApplicationUnits returns a boolean if a field has been set.

### SetApplicationUnitsNil

`func (o *Agreement) SetApplicationUnitsNil(b bool)`

 SetApplicationUnitsNil sets the value for ApplicationUnits to be an explicit nil

### UnsetApplicationUnits
`func (o *Agreement) UnsetApplicationUnits()`

UnsetApplicationUnits ensures that no value is present for ApplicationUnits, not even an explicit nil
### GetApplicationLimit

`func (o *Agreement) GetApplicationLimit() float64`

GetApplicationLimit returns the ApplicationLimit field if non-nil, zero value otherwise.

### GetApplicationLimitOk

`func (o *Agreement) GetApplicationLimitOk() (*float64, bool)`

GetApplicationLimitOk returns a tuple with the ApplicationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLimit

`func (o *Agreement) SetApplicationLimit(v float64)`

SetApplicationLimit sets ApplicationLimit field to given value.

### HasApplicationLimit

`func (o *Agreement) HasApplicationLimit() bool`

HasApplicationLimit returns a boolean if a field has been set.

### SetApplicationLimitNil

`func (o *Agreement) SetApplicationLimitNil(b bool)`

 SetApplicationLimitNil sets the value for ApplicationLimit to be an explicit nil

### UnsetApplicationLimit
`func (o *Agreement) UnsetApplicationLimit()`

UnsetApplicationLimit ensures that no value is present for ApplicationLimit, not even an explicit nil
### GetApplicationCycle

`func (o *Agreement) GetApplicationCycle() string`

GetApplicationCycle returns the ApplicationCycle field if non-nil, zero value otherwise.

### GetApplicationCycleOk

`func (o *Agreement) GetApplicationCycleOk() (*string, bool)`

GetApplicationCycleOk returns a tuple with the ApplicationCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCycle

`func (o *Agreement) SetApplicationCycle(v string)`

SetApplicationCycle sets ApplicationCycle field to given value.

### HasApplicationCycle

`func (o *Agreement) HasApplicationCycle() bool`

HasApplicationCycle returns a boolean if a field has been set.

### SetApplicationCycleNil

`func (o *Agreement) SetApplicationCycleNil(b bool)`

 SetApplicationCycleNil sets the value for ApplicationCycle to be an explicit nil

### UnsetApplicationCycle
`func (o *Agreement) UnsetApplicationCycle()`

UnsetApplicationCycle ensures that no value is present for ApplicationCycle, not even an explicit nil
### GetApplicationUnlimitedFlag

`func (o *Agreement) GetApplicationUnlimitedFlag() bool`

GetApplicationUnlimitedFlag returns the ApplicationUnlimitedFlag field if non-nil, zero value otherwise.

### GetApplicationUnlimitedFlagOk

`func (o *Agreement) GetApplicationUnlimitedFlagOk() (*bool, bool)`

GetApplicationUnlimitedFlagOk returns a tuple with the ApplicationUnlimitedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUnlimitedFlag

`func (o *Agreement) SetApplicationUnlimitedFlag(v bool)`

SetApplicationUnlimitedFlag sets ApplicationUnlimitedFlag field to given value.

### HasApplicationUnlimitedFlag

`func (o *Agreement) HasApplicationUnlimitedFlag() bool`

HasApplicationUnlimitedFlag returns a boolean if a field has been set.

### SetApplicationUnlimitedFlagNil

`func (o *Agreement) SetApplicationUnlimitedFlagNil(b bool)`

 SetApplicationUnlimitedFlagNil sets the value for ApplicationUnlimitedFlag to be an explicit nil

### UnsetApplicationUnlimitedFlag
`func (o *Agreement) UnsetApplicationUnlimitedFlag()`

UnsetApplicationUnlimitedFlag ensures that no value is present for ApplicationUnlimitedFlag, not even an explicit nil
### GetOneTimeFlag

`func (o *Agreement) GetOneTimeFlag() bool`

GetOneTimeFlag returns the OneTimeFlag field if non-nil, zero value otherwise.

### GetOneTimeFlagOk

`func (o *Agreement) GetOneTimeFlagOk() (*bool, bool)`

GetOneTimeFlagOk returns a tuple with the OneTimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimeFlag

`func (o *Agreement) SetOneTimeFlag(v bool)`

SetOneTimeFlag sets OneTimeFlag field to given value.

### HasOneTimeFlag

`func (o *Agreement) HasOneTimeFlag() bool`

HasOneTimeFlag returns a boolean if a field has been set.

### SetOneTimeFlagNil

`func (o *Agreement) SetOneTimeFlagNil(b bool)`

 SetOneTimeFlagNil sets the value for OneTimeFlag to be an explicit nil

### UnsetOneTimeFlag
`func (o *Agreement) UnsetOneTimeFlag()`

UnsetOneTimeFlag ensures that no value is present for OneTimeFlag, not even an explicit nil
### GetCoverAgreementTime

`func (o *Agreement) GetCoverAgreementTime() bool`

GetCoverAgreementTime returns the CoverAgreementTime field if non-nil, zero value otherwise.

### GetCoverAgreementTimeOk

`func (o *Agreement) GetCoverAgreementTimeOk() (*bool, bool)`

GetCoverAgreementTimeOk returns a tuple with the CoverAgreementTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverAgreementTime

`func (o *Agreement) SetCoverAgreementTime(v bool)`

SetCoverAgreementTime sets CoverAgreementTime field to given value.

### HasCoverAgreementTime

`func (o *Agreement) HasCoverAgreementTime() bool`

HasCoverAgreementTime returns a boolean if a field has been set.

### SetCoverAgreementTimeNil

`func (o *Agreement) SetCoverAgreementTimeNil(b bool)`

 SetCoverAgreementTimeNil sets the value for CoverAgreementTime to be an explicit nil

### UnsetCoverAgreementTime
`func (o *Agreement) UnsetCoverAgreementTime()`

UnsetCoverAgreementTime ensures that no value is present for CoverAgreementTime, not even an explicit nil
### GetCoverAgreementProduct

`func (o *Agreement) GetCoverAgreementProduct() bool`

GetCoverAgreementProduct returns the CoverAgreementProduct field if non-nil, zero value otherwise.

### GetCoverAgreementProductOk

`func (o *Agreement) GetCoverAgreementProductOk() (*bool, bool)`

GetCoverAgreementProductOk returns a tuple with the CoverAgreementProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverAgreementProduct

`func (o *Agreement) SetCoverAgreementProduct(v bool)`

SetCoverAgreementProduct sets CoverAgreementProduct field to given value.

### HasCoverAgreementProduct

`func (o *Agreement) HasCoverAgreementProduct() bool`

HasCoverAgreementProduct returns a boolean if a field has been set.

### SetCoverAgreementProductNil

`func (o *Agreement) SetCoverAgreementProductNil(b bool)`

 SetCoverAgreementProductNil sets the value for CoverAgreementProduct to be an explicit nil

### UnsetCoverAgreementProduct
`func (o *Agreement) UnsetCoverAgreementProduct()`

UnsetCoverAgreementProduct ensures that no value is present for CoverAgreementProduct, not even an explicit nil
### GetCoverAgreementExpense

`func (o *Agreement) GetCoverAgreementExpense() bool`

GetCoverAgreementExpense returns the CoverAgreementExpense field if non-nil, zero value otherwise.

### GetCoverAgreementExpenseOk

`func (o *Agreement) GetCoverAgreementExpenseOk() (*bool, bool)`

GetCoverAgreementExpenseOk returns a tuple with the CoverAgreementExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverAgreementExpense

`func (o *Agreement) SetCoverAgreementExpense(v bool)`

SetCoverAgreementExpense sets CoverAgreementExpense field to given value.

### HasCoverAgreementExpense

`func (o *Agreement) HasCoverAgreementExpense() bool`

HasCoverAgreementExpense returns a boolean if a field has been set.

### SetCoverAgreementExpenseNil

`func (o *Agreement) SetCoverAgreementExpenseNil(b bool)`

 SetCoverAgreementExpenseNil sets the value for CoverAgreementExpense to be an explicit nil

### UnsetCoverAgreementExpense
`func (o *Agreement) UnsetCoverAgreementExpense()`

UnsetCoverAgreementExpense ensures that no value is present for CoverAgreementExpense, not even an explicit nil
### GetCoverSalesTax

`func (o *Agreement) GetCoverSalesTax() bool`

GetCoverSalesTax returns the CoverSalesTax field if non-nil, zero value otherwise.

### GetCoverSalesTaxOk

`func (o *Agreement) GetCoverSalesTaxOk() (*bool, bool)`

GetCoverSalesTaxOk returns a tuple with the CoverSalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverSalesTax

`func (o *Agreement) SetCoverSalesTax(v bool)`

SetCoverSalesTax sets CoverSalesTax field to given value.

### HasCoverSalesTax

`func (o *Agreement) HasCoverSalesTax() bool`

HasCoverSalesTax returns a boolean if a field has been set.

### SetCoverSalesTaxNil

`func (o *Agreement) SetCoverSalesTaxNil(b bool)`

 SetCoverSalesTaxNil sets the value for CoverSalesTax to be an explicit nil

### UnsetCoverSalesTax
`func (o *Agreement) UnsetCoverSalesTax()`

UnsetCoverSalesTax ensures that no value is present for CoverSalesTax, not even an explicit nil
### GetCarryOverUnused

`func (o *Agreement) GetCarryOverUnused() bool`

GetCarryOverUnused returns the CarryOverUnused field if non-nil, zero value otherwise.

### GetCarryOverUnusedOk

`func (o *Agreement) GetCarryOverUnusedOk() (*bool, bool)`

GetCarryOverUnusedOk returns a tuple with the CarryOverUnused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarryOverUnused

`func (o *Agreement) SetCarryOverUnused(v bool)`

SetCarryOverUnused sets CarryOverUnused field to given value.

### HasCarryOverUnused

`func (o *Agreement) HasCarryOverUnused() bool`

HasCarryOverUnused returns a boolean if a field has been set.

### SetCarryOverUnusedNil

`func (o *Agreement) SetCarryOverUnusedNil(b bool)`

 SetCarryOverUnusedNil sets the value for CarryOverUnused to be an explicit nil

### UnsetCarryOverUnused
`func (o *Agreement) UnsetCarryOverUnused()`

UnsetCarryOverUnused ensures that no value is present for CarryOverUnused, not even an explicit nil
### GetAllowOverruns

`func (o *Agreement) GetAllowOverruns() bool`

GetAllowOverruns returns the AllowOverruns field if non-nil, zero value otherwise.

### GetAllowOverrunsOk

`func (o *Agreement) GetAllowOverrunsOk() (*bool, bool)`

GetAllowOverrunsOk returns a tuple with the AllowOverruns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverruns

`func (o *Agreement) SetAllowOverruns(v bool)`

SetAllowOverruns sets AllowOverruns field to given value.

### HasAllowOverruns

`func (o *Agreement) HasAllowOverruns() bool`

HasAllowOverruns returns a boolean if a field has been set.

### SetAllowOverrunsNil

`func (o *Agreement) SetAllowOverrunsNil(b bool)`

 SetAllowOverrunsNil sets the value for AllowOverruns to be an explicit nil

### UnsetAllowOverruns
`func (o *Agreement) UnsetAllowOverruns()`

UnsetAllowOverruns ensures that no value is present for AllowOverruns, not even an explicit nil
### GetExpiredDays

`func (o *Agreement) GetExpiredDays() int32`

GetExpiredDays returns the ExpiredDays field if non-nil, zero value otherwise.

### GetExpiredDaysOk

`func (o *Agreement) GetExpiredDaysOk() (*int32, bool)`

GetExpiredDaysOk returns a tuple with the ExpiredDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredDays

`func (o *Agreement) SetExpiredDays(v int32)`

SetExpiredDays sets ExpiredDays field to given value.

### HasExpiredDays

`func (o *Agreement) HasExpiredDays() bool`

HasExpiredDays returns a boolean if a field has been set.

### SetExpiredDaysNil

`func (o *Agreement) SetExpiredDaysNil(b bool)`

 SetExpiredDaysNil sets the value for ExpiredDays to be an explicit nil

### UnsetExpiredDays
`func (o *Agreement) UnsetExpiredDays()`

UnsetExpiredDays ensures that no value is present for ExpiredDays, not even an explicit nil
### GetLimit

`func (o *Agreement) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Agreement) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Agreement) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Agreement) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *Agreement) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *Agreement) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetExpireWhenZero

`func (o *Agreement) GetExpireWhenZero() bool`

GetExpireWhenZero returns the ExpireWhenZero field if non-nil, zero value otherwise.

### GetExpireWhenZeroOk

`func (o *Agreement) GetExpireWhenZeroOk() (*bool, bool)`

GetExpireWhenZeroOk returns a tuple with the ExpireWhenZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWhenZero

`func (o *Agreement) SetExpireWhenZero(v bool)`

SetExpireWhenZero sets ExpireWhenZero field to given value.

### HasExpireWhenZero

`func (o *Agreement) HasExpireWhenZero() bool`

HasExpireWhenZero returns a boolean if a field has been set.

### SetExpireWhenZeroNil

`func (o *Agreement) SetExpireWhenZeroNil(b bool)`

 SetExpireWhenZeroNil sets the value for ExpireWhenZero to be an explicit nil

### UnsetExpireWhenZero
`func (o *Agreement) UnsetExpireWhenZero()`

UnsetExpireWhenZero ensures that no value is present for ExpireWhenZero, not even an explicit nil
### GetChargeToFirm

`func (o *Agreement) GetChargeToFirm() bool`

GetChargeToFirm returns the ChargeToFirm field if non-nil, zero value otherwise.

### GetChargeToFirmOk

`func (o *Agreement) GetChargeToFirmOk() (*bool, bool)`

GetChargeToFirmOk returns a tuple with the ChargeToFirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeToFirm

`func (o *Agreement) SetChargeToFirm(v bool)`

SetChargeToFirm sets ChargeToFirm field to given value.

### HasChargeToFirm

`func (o *Agreement) HasChargeToFirm() bool`

HasChargeToFirm returns a boolean if a field has been set.

### SetChargeToFirmNil

`func (o *Agreement) SetChargeToFirmNil(b bool)`

 SetChargeToFirmNil sets the value for ChargeToFirm to be an explicit nil

### UnsetChargeToFirm
`func (o *Agreement) UnsetChargeToFirm()`

UnsetChargeToFirm ensures that no value is present for ChargeToFirm, not even an explicit nil
### GetEmployeeCompRate

`func (o *Agreement) GetEmployeeCompRate() string`

GetEmployeeCompRate returns the EmployeeCompRate field if non-nil, zero value otherwise.

### GetEmployeeCompRateOk

`func (o *Agreement) GetEmployeeCompRateOk() (*string, bool)`

GetEmployeeCompRateOk returns a tuple with the EmployeeCompRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCompRate

`func (o *Agreement) SetEmployeeCompRate(v string)`

SetEmployeeCompRate sets EmployeeCompRate field to given value.

### HasEmployeeCompRate

`func (o *Agreement) HasEmployeeCompRate() bool`

HasEmployeeCompRate returns a boolean if a field has been set.

### SetEmployeeCompRateNil

`func (o *Agreement) SetEmployeeCompRateNil(b bool)`

 SetEmployeeCompRateNil sets the value for EmployeeCompRate to be an explicit nil

### UnsetEmployeeCompRate
`func (o *Agreement) UnsetEmployeeCompRate()`

UnsetEmployeeCompRate ensures that no value is present for EmployeeCompRate, not even an explicit nil
### GetEmployeeCompNotExceed

`func (o *Agreement) GetEmployeeCompNotExceed() string`

GetEmployeeCompNotExceed returns the EmployeeCompNotExceed field if non-nil, zero value otherwise.

### GetEmployeeCompNotExceedOk

`func (o *Agreement) GetEmployeeCompNotExceedOk() (*string, bool)`

GetEmployeeCompNotExceedOk returns a tuple with the EmployeeCompNotExceed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCompNotExceed

`func (o *Agreement) SetEmployeeCompNotExceed(v string)`

SetEmployeeCompNotExceed sets EmployeeCompNotExceed field to given value.

### HasEmployeeCompNotExceed

`func (o *Agreement) HasEmployeeCompNotExceed() bool`

HasEmployeeCompNotExceed returns a boolean if a field has been set.

### SetEmployeeCompNotExceedNil

`func (o *Agreement) SetEmployeeCompNotExceedNil(b bool)`

 SetEmployeeCompNotExceedNil sets the value for EmployeeCompNotExceed to be an explicit nil

### UnsetEmployeeCompNotExceed
`func (o *Agreement) UnsetEmployeeCompNotExceed()`

UnsetEmployeeCompNotExceed ensures that no value is present for EmployeeCompNotExceed, not even an explicit nil
### GetCompHourlyRate

`func (o *Agreement) GetCompHourlyRate() float64`

GetCompHourlyRate returns the CompHourlyRate field if non-nil, zero value otherwise.

### GetCompHourlyRateOk

`func (o *Agreement) GetCompHourlyRateOk() (*float64, bool)`

GetCompHourlyRateOk returns a tuple with the CompHourlyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompHourlyRate

`func (o *Agreement) SetCompHourlyRate(v float64)`

SetCompHourlyRate sets CompHourlyRate field to given value.

### HasCompHourlyRate

`func (o *Agreement) HasCompHourlyRate() bool`

HasCompHourlyRate returns a boolean if a field has been set.

### SetCompHourlyRateNil

`func (o *Agreement) SetCompHourlyRateNil(b bool)`

 SetCompHourlyRateNil sets the value for CompHourlyRate to be an explicit nil

### UnsetCompHourlyRate
`func (o *Agreement) UnsetCompHourlyRate()`

UnsetCompHourlyRate ensures that no value is present for CompHourlyRate, not even an explicit nil
### GetCompLimitAmount

`func (o *Agreement) GetCompLimitAmount() float64`

GetCompLimitAmount returns the CompLimitAmount field if non-nil, zero value otherwise.

### GetCompLimitAmountOk

`func (o *Agreement) GetCompLimitAmountOk() (*float64, bool)`

GetCompLimitAmountOk returns a tuple with the CompLimitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompLimitAmount

`func (o *Agreement) SetCompLimitAmount(v float64)`

SetCompLimitAmount sets CompLimitAmount field to given value.

### HasCompLimitAmount

`func (o *Agreement) HasCompLimitAmount() bool`

HasCompLimitAmount returns a boolean if a field has been set.

### SetCompLimitAmountNil

`func (o *Agreement) SetCompLimitAmountNil(b bool)`

 SetCompLimitAmountNil sets the value for CompLimitAmount to be an explicit nil

### UnsetCompLimitAmount
`func (o *Agreement) UnsetCompLimitAmount()`

UnsetCompLimitAmount ensures that no value is present for CompLimitAmount, not even an explicit nil
### GetBillingCycle

`func (o *Agreement) GetBillingCycle() BillingCycleReference`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *Agreement) GetBillingCycleOk() (*BillingCycleReference, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *Agreement) SetBillingCycle(v BillingCycleReference)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *Agreement) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetBillOneTimeFlag

`func (o *Agreement) GetBillOneTimeFlag() bool`

GetBillOneTimeFlag returns the BillOneTimeFlag field if non-nil, zero value otherwise.

### GetBillOneTimeFlagOk

`func (o *Agreement) GetBillOneTimeFlagOk() (*bool, bool)`

GetBillOneTimeFlagOk returns a tuple with the BillOneTimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOneTimeFlag

`func (o *Agreement) SetBillOneTimeFlag(v bool)`

SetBillOneTimeFlag sets BillOneTimeFlag field to given value.

### HasBillOneTimeFlag

`func (o *Agreement) HasBillOneTimeFlag() bool`

HasBillOneTimeFlag returns a boolean if a field has been set.

### SetBillOneTimeFlagNil

`func (o *Agreement) SetBillOneTimeFlagNil(b bool)`

 SetBillOneTimeFlagNil sets the value for BillOneTimeFlag to be an explicit nil

### UnsetBillOneTimeFlag
`func (o *Agreement) UnsetBillOneTimeFlag()`

UnsetBillOneTimeFlag ensures that no value is present for BillOneTimeFlag, not even an explicit nil
### GetBillingTerms

`func (o *Agreement) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *Agreement) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *Agreement) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *Agreement) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetInvoicingCycle

`func (o *Agreement) GetInvoicingCycle() string`

GetInvoicingCycle returns the InvoicingCycle field if non-nil, zero value otherwise.

### GetInvoicingCycleOk

`func (o *Agreement) GetInvoicingCycleOk() (*string, bool)`

GetInvoicingCycleOk returns a tuple with the InvoicingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingCycle

`func (o *Agreement) SetInvoicingCycle(v string)`

SetInvoicingCycle sets InvoicingCycle field to given value.

### HasInvoicingCycle

`func (o *Agreement) HasInvoicingCycle() bool`

HasInvoicingCycle returns a boolean if a field has been set.

### SetInvoicingCycleNil

`func (o *Agreement) SetInvoicingCycleNil(b bool)`

 SetInvoicingCycleNil sets the value for InvoicingCycle to be an explicit nil

### UnsetInvoicingCycle
`func (o *Agreement) UnsetInvoicingCycle()`

UnsetInvoicingCycle ensures that no value is present for InvoicingCycle, not even an explicit nil
### GetBillToCompany

`func (o *Agreement) GetBillToCompany() CompanyReference`

GetBillToCompany returns the BillToCompany field if non-nil, zero value otherwise.

### GetBillToCompanyOk

`func (o *Agreement) GetBillToCompanyOk() (*CompanyReference, bool)`

GetBillToCompanyOk returns a tuple with the BillToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToCompany

`func (o *Agreement) SetBillToCompany(v CompanyReference)`

SetBillToCompany sets BillToCompany field to given value.

### HasBillToCompany

`func (o *Agreement) HasBillToCompany() bool`

HasBillToCompany returns a boolean if a field has been set.

### GetBillToContact

`func (o *Agreement) GetBillToContact() ContactReference`

GetBillToContact returns the BillToContact field if non-nil, zero value otherwise.

### GetBillToContactOk

`func (o *Agreement) GetBillToContactOk() (*ContactReference, bool)`

GetBillToContactOk returns a tuple with the BillToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToContact

`func (o *Agreement) SetBillToContact(v ContactReference)`

SetBillToContact sets BillToContact field to given value.

### HasBillToContact

`func (o *Agreement) HasBillToContact() bool`

HasBillToContact returns a boolean if a field has been set.

### GetBillToSite

`func (o *Agreement) GetBillToSite() SiteReference`

GetBillToSite returns the BillToSite field if non-nil, zero value otherwise.

### GetBillToSiteOk

`func (o *Agreement) GetBillToSiteOk() (*SiteReference, bool)`

GetBillToSiteOk returns a tuple with the BillToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToSite

`func (o *Agreement) SetBillToSite(v SiteReference)`

SetBillToSite sets BillToSite field to given value.

### HasBillToSite

`func (o *Agreement) HasBillToSite() bool`

HasBillToSite returns a boolean if a field has been set.

### GetBillAmount

`func (o *Agreement) GetBillAmount() float64`

GetBillAmount returns the BillAmount field if non-nil, zero value otherwise.

### GetBillAmountOk

`func (o *Agreement) GetBillAmountOk() (*float64, bool)`

GetBillAmountOk returns a tuple with the BillAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillAmount

`func (o *Agreement) SetBillAmount(v float64)`

SetBillAmount sets BillAmount field to given value.

### HasBillAmount

`func (o *Agreement) HasBillAmount() bool`

HasBillAmount returns a boolean if a field has been set.

### SetBillAmountNil

`func (o *Agreement) SetBillAmountNil(b bool)`

 SetBillAmountNil sets the value for BillAmount to be an explicit nil

### UnsetBillAmount
`func (o *Agreement) UnsetBillAmount()`

UnsetBillAmount ensures that no value is present for BillAmount, not even an explicit nil
### GetTaxable

`func (o *Agreement) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *Agreement) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *Agreement) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *Agreement) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### SetTaxableNil

`func (o *Agreement) SetTaxableNil(b bool)`

 SetTaxableNil sets the value for Taxable to be an explicit nil

### UnsetTaxable
`func (o *Agreement) UnsetTaxable()`

UnsetTaxable ensures that no value is present for Taxable, not even an explicit nil
### GetProrateFirstBill

`func (o *Agreement) GetProrateFirstBill() float64`

GetProrateFirstBill returns the ProrateFirstBill field if non-nil, zero value otherwise.

### GetProrateFirstBillOk

`func (o *Agreement) GetProrateFirstBillOk() (*float64, bool)`

GetProrateFirstBillOk returns a tuple with the ProrateFirstBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrateFirstBill

`func (o *Agreement) SetProrateFirstBill(v float64)`

SetProrateFirstBill sets ProrateFirstBill field to given value.

### HasProrateFirstBill

`func (o *Agreement) HasProrateFirstBill() bool`

HasProrateFirstBill returns a boolean if a field has been set.

### SetProrateFirstBillNil

`func (o *Agreement) SetProrateFirstBillNil(b bool)`

 SetProrateFirstBillNil sets the value for ProrateFirstBill to be an explicit nil

### UnsetProrateFirstBill
`func (o *Agreement) UnsetProrateFirstBill()`

UnsetProrateFirstBill ensures that no value is present for ProrateFirstBill, not even an explicit nil
### GetBillStartDate

`func (o *Agreement) GetBillStartDate() time.Time`

GetBillStartDate returns the BillStartDate field if non-nil, zero value otherwise.

### GetBillStartDateOk

`func (o *Agreement) GetBillStartDateOk() (*time.Time, bool)`

GetBillStartDateOk returns a tuple with the BillStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillStartDate

`func (o *Agreement) SetBillStartDate(v time.Time)`

SetBillStartDate sets BillStartDate field to given value.

### HasBillStartDate

`func (o *Agreement) HasBillStartDate() bool`

HasBillStartDate returns a boolean if a field has been set.

### GetTaxCode

`func (o *Agreement) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *Agreement) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *Agreement) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *Agreement) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetRestrictDownPayment

`func (o *Agreement) GetRestrictDownPayment() bool`

GetRestrictDownPayment returns the RestrictDownPayment field if non-nil, zero value otherwise.

### GetRestrictDownPaymentOk

`func (o *Agreement) GetRestrictDownPaymentOk() (*bool, bool)`

GetRestrictDownPaymentOk returns a tuple with the RestrictDownPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDownPayment

`func (o *Agreement) SetRestrictDownPayment(v bool)`

SetRestrictDownPayment sets RestrictDownPayment field to given value.

### HasRestrictDownPayment

`func (o *Agreement) HasRestrictDownPayment() bool`

HasRestrictDownPayment returns a boolean if a field has been set.

### SetRestrictDownPaymentNil

`func (o *Agreement) SetRestrictDownPaymentNil(b bool)`

 SetRestrictDownPaymentNil sets the value for RestrictDownPayment to be an explicit nil

### UnsetRestrictDownPayment
`func (o *Agreement) UnsetRestrictDownPayment()`

UnsetRestrictDownPayment ensures that no value is present for RestrictDownPayment, not even an explicit nil
### GetProrateFlag

`func (o *Agreement) GetProrateFlag() bool`

GetProrateFlag returns the ProrateFlag field if non-nil, zero value otherwise.

### GetProrateFlagOk

`func (o *Agreement) GetProrateFlagOk() (*bool, bool)`

GetProrateFlagOk returns a tuple with the ProrateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrateFlag

`func (o *Agreement) SetProrateFlag(v bool)`

SetProrateFlag sets ProrateFlag field to given value.

### HasProrateFlag

`func (o *Agreement) HasProrateFlag() bool`

HasProrateFlag returns a boolean if a field has been set.

### SetProrateFlagNil

`func (o *Agreement) SetProrateFlagNil(b bool)`

 SetProrateFlagNil sets the value for ProrateFlag to be an explicit nil

### UnsetProrateFlag
`func (o *Agreement) UnsetProrateFlag()`

UnsetProrateFlag ensures that no value is present for ProrateFlag, not even an explicit nil
### GetInvoiceProratedAdditionsFlag

`func (o *Agreement) GetInvoiceProratedAdditionsFlag() bool`

GetInvoiceProratedAdditionsFlag returns the InvoiceProratedAdditionsFlag field if non-nil, zero value otherwise.

### GetInvoiceProratedAdditionsFlagOk

`func (o *Agreement) GetInvoiceProratedAdditionsFlagOk() (*bool, bool)`

GetInvoiceProratedAdditionsFlagOk returns a tuple with the InvoiceProratedAdditionsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceProratedAdditionsFlag

`func (o *Agreement) SetInvoiceProratedAdditionsFlag(v bool)`

SetInvoiceProratedAdditionsFlag sets InvoiceProratedAdditionsFlag field to given value.

### HasInvoiceProratedAdditionsFlag

`func (o *Agreement) HasInvoiceProratedAdditionsFlag() bool`

HasInvoiceProratedAdditionsFlag returns a boolean if a field has been set.

### SetInvoiceProratedAdditionsFlagNil

`func (o *Agreement) SetInvoiceProratedAdditionsFlagNil(b bool)`

 SetInvoiceProratedAdditionsFlagNil sets the value for InvoiceProratedAdditionsFlag to be an explicit nil

### UnsetInvoiceProratedAdditionsFlag
`func (o *Agreement) UnsetInvoiceProratedAdditionsFlag()`

UnsetInvoiceProratedAdditionsFlag ensures that no value is present for InvoiceProratedAdditionsFlag, not even an explicit nil
### GetInvoiceDescription

`func (o *Agreement) GetInvoiceDescription() string`

GetInvoiceDescription returns the InvoiceDescription field if non-nil, zero value otherwise.

### GetInvoiceDescriptionOk

`func (o *Agreement) GetInvoiceDescriptionOk() (*string, bool)`

GetInvoiceDescriptionOk returns a tuple with the InvoiceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDescription

`func (o *Agreement) SetInvoiceDescription(v string)`

SetInvoiceDescription sets InvoiceDescription field to given value.

### HasInvoiceDescription

`func (o *Agreement) HasInvoiceDescription() bool`

HasInvoiceDescription returns a boolean if a field has been set.

### GetTopComment

`func (o *Agreement) GetTopComment() bool`

GetTopComment returns the TopComment field if non-nil, zero value otherwise.

### GetTopCommentOk

`func (o *Agreement) GetTopCommentOk() (*bool, bool)`

GetTopCommentOk returns a tuple with the TopComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopComment

`func (o *Agreement) SetTopComment(v bool)`

SetTopComment sets TopComment field to given value.

### HasTopComment

`func (o *Agreement) HasTopComment() bool`

HasTopComment returns a boolean if a field has been set.

### SetTopCommentNil

`func (o *Agreement) SetTopCommentNil(b bool)`

 SetTopCommentNil sets the value for TopComment to be an explicit nil

### UnsetTopComment
`func (o *Agreement) UnsetTopComment()`

UnsetTopComment ensures that no value is present for TopComment, not even an explicit nil
### GetBottomComment

`func (o *Agreement) GetBottomComment() bool`

GetBottomComment returns the BottomComment field if non-nil, zero value otherwise.

### GetBottomCommentOk

`func (o *Agreement) GetBottomCommentOk() (*bool, bool)`

GetBottomCommentOk returns a tuple with the BottomComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomComment

`func (o *Agreement) SetBottomComment(v bool)`

SetBottomComment sets BottomComment field to given value.

### HasBottomComment

`func (o *Agreement) HasBottomComment() bool`

HasBottomComment returns a boolean if a field has been set.

### SetBottomCommentNil

`func (o *Agreement) SetBottomCommentNil(b bool)`

 SetBottomCommentNil sets the value for BottomComment to be an explicit nil

### UnsetBottomComment
`func (o *Agreement) UnsetBottomComment()`

UnsetBottomComment ensures that no value is present for BottomComment, not even an explicit nil
### GetWorkRole

`func (o *Agreement) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *Agreement) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *Agreement) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *Agreement) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *Agreement) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *Agreement) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *Agreement) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *Agreement) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetProjectType

`func (o *Agreement) GetProjectType() ProjectTypeReference`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *Agreement) GetProjectTypeOk() (*ProjectTypeReference, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *Agreement) SetProjectType(v ProjectTypeReference)`

SetProjectType sets ProjectType field to given value.

### HasProjectType

`func (o *Agreement) HasProjectType() bool`

HasProjectType returns a boolean if a field has been set.

### GetInvoiceTemplate

`func (o *Agreement) GetInvoiceTemplate() InvoiceTemplateReference`

GetInvoiceTemplate returns the InvoiceTemplate field if non-nil, zero value otherwise.

### GetInvoiceTemplateOk

`func (o *Agreement) GetInvoiceTemplateOk() (*InvoiceTemplateReference, bool)`

GetInvoiceTemplateOk returns a tuple with the InvoiceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTemplate

`func (o *Agreement) SetInvoiceTemplate(v InvoiceTemplateReference)`

SetInvoiceTemplate sets InvoiceTemplate field to given value.

### HasInvoiceTemplate

`func (o *Agreement) HasInvoiceTemplate() bool`

HasInvoiceTemplate returns a boolean if a field has been set.

### GetBillTime

`func (o *Agreement) GetBillTime() string`

GetBillTime returns the BillTime field if non-nil, zero value otherwise.

### GetBillTimeOk

`func (o *Agreement) GetBillTimeOk() (*string, bool)`

GetBillTimeOk returns a tuple with the BillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTime

`func (o *Agreement) SetBillTime(v string)`

SetBillTime sets BillTime field to given value.

### HasBillTime

`func (o *Agreement) HasBillTime() bool`

HasBillTime returns a boolean if a field has been set.

### SetBillTimeNil

`func (o *Agreement) SetBillTimeNil(b bool)`

 SetBillTimeNil sets the value for BillTime to be an explicit nil

### UnsetBillTime
`func (o *Agreement) UnsetBillTime()`

UnsetBillTime ensures that no value is present for BillTime, not even an explicit nil
### GetBillExpenses

`func (o *Agreement) GetBillExpenses() string`

GetBillExpenses returns the BillExpenses field if non-nil, zero value otherwise.

### GetBillExpensesOk

`func (o *Agreement) GetBillExpensesOk() (*string, bool)`

GetBillExpensesOk returns a tuple with the BillExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillExpenses

`func (o *Agreement) SetBillExpenses(v string)`

SetBillExpenses sets BillExpenses field to given value.

### HasBillExpenses

`func (o *Agreement) HasBillExpenses() bool`

HasBillExpenses returns a boolean if a field has been set.

### SetBillExpensesNil

`func (o *Agreement) SetBillExpensesNil(b bool)`

 SetBillExpensesNil sets the value for BillExpenses to be an explicit nil

### UnsetBillExpenses
`func (o *Agreement) UnsetBillExpenses()`

UnsetBillExpenses ensures that no value is present for BillExpenses, not even an explicit nil
### GetBillProducts

`func (o *Agreement) GetBillProducts() string`

GetBillProducts returns the BillProducts field if non-nil, zero value otherwise.

### GetBillProductsOk

`func (o *Agreement) GetBillProductsOk() (*string, bool)`

GetBillProductsOk returns a tuple with the BillProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProducts

`func (o *Agreement) SetBillProducts(v string)`

SetBillProducts sets BillProducts field to given value.

### HasBillProducts

`func (o *Agreement) HasBillProducts() bool`

HasBillProducts returns a boolean if a field has been set.

### SetBillProductsNil

`func (o *Agreement) SetBillProductsNil(b bool)`

 SetBillProductsNil sets the value for BillProducts to be an explicit nil

### UnsetBillProducts
`func (o *Agreement) UnsetBillProducts()`

UnsetBillProducts ensures that no value is present for BillProducts, not even an explicit nil
### GetBillableTimeInvoice

`func (o *Agreement) GetBillableTimeInvoice() bool`

GetBillableTimeInvoice returns the BillableTimeInvoice field if non-nil, zero value otherwise.

### GetBillableTimeInvoiceOk

`func (o *Agreement) GetBillableTimeInvoiceOk() (*bool, bool)`

GetBillableTimeInvoiceOk returns a tuple with the BillableTimeInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableTimeInvoice

`func (o *Agreement) SetBillableTimeInvoice(v bool)`

SetBillableTimeInvoice sets BillableTimeInvoice field to given value.

### HasBillableTimeInvoice

`func (o *Agreement) HasBillableTimeInvoice() bool`

HasBillableTimeInvoice returns a boolean if a field has been set.

### SetBillableTimeInvoiceNil

`func (o *Agreement) SetBillableTimeInvoiceNil(b bool)`

 SetBillableTimeInvoiceNil sets the value for BillableTimeInvoice to be an explicit nil

### UnsetBillableTimeInvoice
`func (o *Agreement) UnsetBillableTimeInvoice()`

UnsetBillableTimeInvoice ensures that no value is present for BillableTimeInvoice, not even an explicit nil
### GetBillableExpenseInvoice

`func (o *Agreement) GetBillableExpenseInvoice() bool`

GetBillableExpenseInvoice returns the BillableExpenseInvoice field if non-nil, zero value otherwise.

### GetBillableExpenseInvoiceOk

`func (o *Agreement) GetBillableExpenseInvoiceOk() (*bool, bool)`

GetBillableExpenseInvoiceOk returns a tuple with the BillableExpenseInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableExpenseInvoice

`func (o *Agreement) SetBillableExpenseInvoice(v bool)`

SetBillableExpenseInvoice sets BillableExpenseInvoice field to given value.

### HasBillableExpenseInvoice

`func (o *Agreement) HasBillableExpenseInvoice() bool`

HasBillableExpenseInvoice returns a boolean if a field has been set.

### SetBillableExpenseInvoiceNil

`func (o *Agreement) SetBillableExpenseInvoiceNil(b bool)`

 SetBillableExpenseInvoiceNil sets the value for BillableExpenseInvoice to be an explicit nil

### UnsetBillableExpenseInvoice
`func (o *Agreement) UnsetBillableExpenseInvoice()`

UnsetBillableExpenseInvoice ensures that no value is present for BillableExpenseInvoice, not even an explicit nil
### GetBillableProductInvoice

`func (o *Agreement) GetBillableProductInvoice() bool`

GetBillableProductInvoice returns the BillableProductInvoice field if non-nil, zero value otherwise.

### GetBillableProductInvoiceOk

`func (o *Agreement) GetBillableProductInvoiceOk() (*bool, bool)`

GetBillableProductInvoiceOk returns a tuple with the BillableProductInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableProductInvoice

`func (o *Agreement) SetBillableProductInvoice(v bool)`

SetBillableProductInvoice sets BillableProductInvoice field to given value.

### HasBillableProductInvoice

`func (o *Agreement) HasBillableProductInvoice() bool`

HasBillableProductInvoice returns a boolean if a field has been set.

### SetBillableProductInvoiceNil

`func (o *Agreement) SetBillableProductInvoiceNil(b bool)`

 SetBillableProductInvoiceNil sets the value for BillableProductInvoice to be an explicit nil

### UnsetBillableProductInvoice
`func (o *Agreement) UnsetBillableProductInvoice()`

UnsetBillableProductInvoice ensures that no value is present for BillableProductInvoice, not even an explicit nil
### GetCurrency

`func (o *Agreement) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Agreement) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Agreement) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Agreement) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPeriodType

`func (o *Agreement) GetPeriodType() string`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *Agreement) GetPeriodTypeOk() (*string, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *Agreement) SetPeriodType(v string)`

SetPeriodType sets PeriodType field to given value.

### HasPeriodType

`func (o *Agreement) HasPeriodType() bool`

HasPeriodType returns a boolean if a field has been set.

### SetPeriodTypeNil

`func (o *Agreement) SetPeriodTypeNil(b bool)`

 SetPeriodTypeNil sets the value for PeriodType to be an explicit nil

### UnsetPeriodType
`func (o *Agreement) UnsetPeriodType()`

UnsetPeriodType ensures that no value is present for PeriodType, not even an explicit nil
### GetAutoInvoiceFlag

`func (o *Agreement) GetAutoInvoiceFlag() bool`

GetAutoInvoiceFlag returns the AutoInvoiceFlag field if non-nil, zero value otherwise.

### GetAutoInvoiceFlagOk

`func (o *Agreement) GetAutoInvoiceFlagOk() (*bool, bool)`

GetAutoInvoiceFlagOk returns a tuple with the AutoInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInvoiceFlag

`func (o *Agreement) SetAutoInvoiceFlag(v bool)`

SetAutoInvoiceFlag sets AutoInvoiceFlag field to given value.

### HasAutoInvoiceFlag

`func (o *Agreement) HasAutoInvoiceFlag() bool`

HasAutoInvoiceFlag returns a boolean if a field has been set.

### SetAutoInvoiceFlagNil

`func (o *Agreement) SetAutoInvoiceFlagNil(b bool)`

 SetAutoInvoiceFlagNil sets the value for AutoInvoiceFlag to be an explicit nil

### UnsetAutoInvoiceFlag
`func (o *Agreement) UnsetAutoInvoiceFlag()`

UnsetAutoInvoiceFlag ensures that no value is present for AutoInvoiceFlag, not even an explicit nil
### GetNextInvoiceDate

`func (o *Agreement) GetNextInvoiceDate() string`

GetNextInvoiceDate returns the NextInvoiceDate field if non-nil, zero value otherwise.

### GetNextInvoiceDateOk

`func (o *Agreement) GetNextInvoiceDateOk() (*string, bool)`

GetNextInvoiceDateOk returns a tuple with the NextInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceDate

`func (o *Agreement) SetNextInvoiceDate(v string)`

SetNextInvoiceDate sets NextInvoiceDate field to given value.

### HasNextInvoiceDate

`func (o *Agreement) HasNextInvoiceDate() bool`

HasNextInvoiceDate returns a boolean if a field has been set.

### GetCompanyLocation

`func (o *Agreement) GetCompanyLocation() SystemLocationReference`

GetCompanyLocation returns the CompanyLocation field if non-nil, zero value otherwise.

### GetCompanyLocationOk

`func (o *Agreement) GetCompanyLocationOk() (*SystemLocationReference, bool)`

GetCompanyLocationOk returns a tuple with the CompanyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLocation

`func (o *Agreement) SetCompanyLocation(v SystemLocationReference)`

SetCompanyLocation sets CompanyLocation field to given value.

### HasCompanyLocation

`func (o *Agreement) HasCompanyLocation() bool`

HasCompanyLocation returns a boolean if a field has been set.

### GetShipToCompany

`func (o *Agreement) GetShipToCompany() CompanyReference`

GetShipToCompany returns the ShipToCompany field if non-nil, zero value otherwise.

### GetShipToCompanyOk

`func (o *Agreement) GetShipToCompanyOk() (*CompanyReference, bool)`

GetShipToCompanyOk returns a tuple with the ShipToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompany

`func (o *Agreement) SetShipToCompany(v CompanyReference)`

SetShipToCompany sets ShipToCompany field to given value.

### HasShipToCompany

`func (o *Agreement) HasShipToCompany() bool`

HasShipToCompany returns a boolean if a field has been set.

### GetShipToContact

`func (o *Agreement) GetShipToContact() ContactReference`

GetShipToContact returns the ShipToContact field if non-nil, zero value otherwise.

### GetShipToContactOk

`func (o *Agreement) GetShipToContactOk() (*ContactReference, bool)`

GetShipToContactOk returns a tuple with the ShipToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToContact

`func (o *Agreement) SetShipToContact(v ContactReference)`

SetShipToContact sets ShipToContact field to given value.

### HasShipToContact

`func (o *Agreement) HasShipToContact() bool`

HasShipToContact returns a boolean if a field has been set.

### GetShipToSite

`func (o *Agreement) GetShipToSite() SiteReference`

GetShipToSite returns the ShipToSite field if non-nil, zero value otherwise.

### GetShipToSiteOk

`func (o *Agreement) GetShipToSiteOk() (*SiteReference, bool)`

GetShipToSiteOk returns a tuple with the ShipToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToSite

`func (o *Agreement) SetShipToSite(v SiteReference)`

SetShipToSite sets ShipToSite field to given value.

### HasShipToSite

`func (o *Agreement) HasShipToSite() bool`

HasShipToSite returns a boolean if a field has been set.

### GetAgreementStatus

`func (o *Agreement) GetAgreementStatus() string`

GetAgreementStatus returns the AgreementStatus field if non-nil, zero value otherwise.

### GetAgreementStatusOk

`func (o *Agreement) GetAgreementStatusOk() (*string, bool)`

GetAgreementStatusOk returns a tuple with the AgreementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementStatus

`func (o *Agreement) SetAgreementStatus(v string)`

SetAgreementStatus sets AgreementStatus field to given value.

### HasAgreementStatus

`func (o *Agreement) HasAgreementStatus() bool`

HasAgreementStatus returns a boolean if a field has been set.

### SetAgreementStatusNil

`func (o *Agreement) SetAgreementStatusNil(b bool)`

 SetAgreementStatusNil sets the value for AgreementStatus to be an explicit nil

### UnsetAgreementStatus
`func (o *Agreement) UnsetAgreementStatus()`

UnsetAgreementStatus ensures that no value is present for AgreementStatus, not even an explicit nil
### GetInfo

`func (o *Agreement) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Agreement) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Agreement) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Agreement) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *Agreement) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Agreement) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Agreement) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Agreement) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


