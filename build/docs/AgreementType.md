# AgreementType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**PrefixSuffixOption** | Pointer to **NullableString** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**PrePaymentFlag** | Pointer to **NullableBool** |  | [optional] 
**InvoicePreSuffix** | Pointer to **string** |  Max length: 5; | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**RestrictLocationFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictDepartmentFlag** | Pointer to **NullableBool** |  | [optional] 
**Sla** | Pointer to [**SLAReference**](SLAReference.md) |  | [optional] 
**ApplicationUnits** | Pointer to **NullableString** |  | [optional] 
**ApplicationLimit** | Pointer to **NullableFloat64** |  | [optional] 
**ApplicationCycle** | Pointer to **NullableString** |  | [optional] 
**ApplicationUnlimitedFlag** | Pointer to **NullableBool** |  | [optional] 
**OneTimeFlag** | Pointer to **NullableBool** |  | [optional] 
**CoverAgreementTimeFlag** | Pointer to **NullableBool** |  | [optional] 
**CoverAgreementProductFlag** | Pointer to **NullableBool** |  | [optional] 
**CoverAgreementExpenseFlag** | Pointer to **NullableBool** |  | [optional] 
**CoverSalesTaxFlag** | Pointer to **NullableBool** |  | [optional] 
**CarryOverUnusedFlag** | Pointer to **NullableBool** |  | [optional] 
**AllowOverrunsFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpiredDays** | Pointer to **NullableInt32** |  | [optional] 
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**ExpireWhenZero** | Pointer to **NullableBool** |  | [optional] 
**ChargeToFirmFlag** | Pointer to **NullableBool** |  | [optional] 
**EmployeeCompRate** | **NullableString** |  | 
**EmployeeCompNotExceed** | **NullableString** |  | 
**CompHourlyRate** | Pointer to **NullableFloat64** |  | [optional] 
**CompLimitAmount** | Pointer to **NullableFloat64** |  | [optional] 
**BillingCycle** | Pointer to [**BillingCycleReference**](BillingCycleReference.md) |  | [optional] 
**BillOneTimeFlag** | Pointer to **NullableBool** |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**InvoicingCycle** | **NullableString** |  | 
**BillAmount** | Pointer to **NullableFloat64** |  | [optional] 
**TaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictDownPaymentFlag** | Pointer to **NullableBool** |  | [optional] 
**InvoiceDescription** | Pointer to **string** |  Max length: 4000; | [optional] 
**TopCommentFlag** | Pointer to **NullableBool** |  | [optional] 
**BottomCommentFlag** | Pointer to **NullableBool** |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**ProjectType** | Pointer to [**ProjectTypeReference**](ProjectTypeReference.md) |  | [optional] 
**InvoiceTemplate** | Pointer to [**InvoiceTemplateReference**](InvoiceTemplateReference.md) |  | [optional] 
**BillTime** | **NullableString** |  | 
**BillExpenses** | **NullableString** |  | 
**BillProducts** | **NullableString** |  | 
**BillableTimeInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**BillableExpenseInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**BillableProductInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**CopyWorkRolesFlag** | Pointer to **NullableBool** |  | [optional] 
**CopyWorkTypesFlag** | Pointer to **NullableBool** |  | [optional] 
**ExclusionWorkRoleIds** | Pointer to **[]int32** |  | [optional] 
**AddAllWorkRoleExclusions** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllWorkRoleExclusions** | Pointer to **NullableBool** |  | [optional] 
**ExclusionWorkTypeIds** | Pointer to **[]int32** |  | [optional] 
**AddAllWorkTypeExclusions** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllWorkTypeExclusions** | Pointer to **NullableBool** |  | [optional] 
**IntegrationXRef** | Pointer to **string** |  Max length: 50; | [optional] 
**ProrateFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailTemplate** | Pointer to [**EmailTemplateReference**](EmailTemplateReference.md) |  | [optional] 
**AutoInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**InvoiceProratedAdditionsFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementType

`func NewAgreementType(name string, employeeCompRate NullableString, employeeCompNotExceed NullableString, invoicingCycle NullableString, billTime NullableString, billExpenses NullableString, billProducts NullableString, ) *AgreementType`

NewAgreementType instantiates a new AgreementType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementTypeWithDefaults

`func NewAgreementTypeWithDefaults() *AgreementType`

NewAgreementTypeWithDefaults instantiates a new AgreementType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AgreementType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgreementType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgreementType) SetName(v string)`

SetName sets Name field to given value.


### GetPrefixSuffixOption

`func (o *AgreementType) GetPrefixSuffixOption() string`

GetPrefixSuffixOption returns the PrefixSuffixOption field if non-nil, zero value otherwise.

### GetPrefixSuffixOptionOk

`func (o *AgreementType) GetPrefixSuffixOptionOk() (*string, bool)`

GetPrefixSuffixOptionOk returns a tuple with the PrefixSuffixOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSuffixOption

`func (o *AgreementType) SetPrefixSuffixOption(v string)`

SetPrefixSuffixOption sets PrefixSuffixOption field to given value.

### HasPrefixSuffixOption

`func (o *AgreementType) HasPrefixSuffixOption() bool`

HasPrefixSuffixOption returns a boolean if a field has been set.

### SetPrefixSuffixOptionNil

`func (o *AgreementType) SetPrefixSuffixOptionNil(b bool)`

 SetPrefixSuffixOptionNil sets the value for PrefixSuffixOption to be an explicit nil

### UnsetPrefixSuffixOption
`func (o *AgreementType) UnsetPrefixSuffixOption()`

UnsetPrefixSuffixOption ensures that no value is present for PrefixSuffixOption, not even an explicit nil
### GetDefaultFlag

`func (o *AgreementType) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *AgreementType) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *AgreementType) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *AgreementType) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *AgreementType) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *AgreementType) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *AgreementType) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *AgreementType) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *AgreementType) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *AgreementType) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *AgreementType) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *AgreementType) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetPrePaymentFlag

`func (o *AgreementType) GetPrePaymentFlag() bool`

GetPrePaymentFlag returns the PrePaymentFlag field if non-nil, zero value otherwise.

### GetPrePaymentFlagOk

`func (o *AgreementType) GetPrePaymentFlagOk() (*bool, bool)`

GetPrePaymentFlagOk returns a tuple with the PrePaymentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaymentFlag

`func (o *AgreementType) SetPrePaymentFlag(v bool)`

SetPrePaymentFlag sets PrePaymentFlag field to given value.

### HasPrePaymentFlag

`func (o *AgreementType) HasPrePaymentFlag() bool`

HasPrePaymentFlag returns a boolean if a field has been set.

### SetPrePaymentFlagNil

`func (o *AgreementType) SetPrePaymentFlagNil(b bool)`

 SetPrePaymentFlagNil sets the value for PrePaymentFlag to be an explicit nil

### UnsetPrePaymentFlag
`func (o *AgreementType) UnsetPrePaymentFlag()`

UnsetPrePaymentFlag ensures that no value is present for PrePaymentFlag, not even an explicit nil
### GetInvoicePreSuffix

`func (o *AgreementType) GetInvoicePreSuffix() string`

GetInvoicePreSuffix returns the InvoicePreSuffix field if non-nil, zero value otherwise.

### GetInvoicePreSuffixOk

`func (o *AgreementType) GetInvoicePreSuffixOk() (*string, bool)`

GetInvoicePreSuffixOk returns a tuple with the InvoicePreSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePreSuffix

`func (o *AgreementType) SetInvoicePreSuffix(v string)`

SetInvoicePreSuffix sets InvoicePreSuffix field to given value.

### HasInvoicePreSuffix

`func (o *AgreementType) HasInvoicePreSuffix() bool`

HasInvoicePreSuffix returns a boolean if a field has been set.

### GetLocation

`func (o *AgreementType) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AgreementType) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AgreementType) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AgreementType) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *AgreementType) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *AgreementType) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *AgreementType) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *AgreementType) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetRestrictLocationFlag

`func (o *AgreementType) GetRestrictLocationFlag() bool`

GetRestrictLocationFlag returns the RestrictLocationFlag field if non-nil, zero value otherwise.

### GetRestrictLocationFlagOk

`func (o *AgreementType) GetRestrictLocationFlagOk() (*bool, bool)`

GetRestrictLocationFlagOk returns a tuple with the RestrictLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictLocationFlag

`func (o *AgreementType) SetRestrictLocationFlag(v bool)`

SetRestrictLocationFlag sets RestrictLocationFlag field to given value.

### HasRestrictLocationFlag

`func (o *AgreementType) HasRestrictLocationFlag() bool`

HasRestrictLocationFlag returns a boolean if a field has been set.

### SetRestrictLocationFlagNil

`func (o *AgreementType) SetRestrictLocationFlagNil(b bool)`

 SetRestrictLocationFlagNil sets the value for RestrictLocationFlag to be an explicit nil

### UnsetRestrictLocationFlag
`func (o *AgreementType) UnsetRestrictLocationFlag()`

UnsetRestrictLocationFlag ensures that no value is present for RestrictLocationFlag, not even an explicit nil
### GetRestrictDepartmentFlag

`func (o *AgreementType) GetRestrictDepartmentFlag() bool`

GetRestrictDepartmentFlag returns the RestrictDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictDepartmentFlagOk

`func (o *AgreementType) GetRestrictDepartmentFlagOk() (*bool, bool)`

GetRestrictDepartmentFlagOk returns a tuple with the RestrictDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDepartmentFlag

`func (o *AgreementType) SetRestrictDepartmentFlag(v bool)`

SetRestrictDepartmentFlag sets RestrictDepartmentFlag field to given value.

### HasRestrictDepartmentFlag

`func (o *AgreementType) HasRestrictDepartmentFlag() bool`

HasRestrictDepartmentFlag returns a boolean if a field has been set.

### SetRestrictDepartmentFlagNil

`func (o *AgreementType) SetRestrictDepartmentFlagNil(b bool)`

 SetRestrictDepartmentFlagNil sets the value for RestrictDepartmentFlag to be an explicit nil

### UnsetRestrictDepartmentFlag
`func (o *AgreementType) UnsetRestrictDepartmentFlag()`

UnsetRestrictDepartmentFlag ensures that no value is present for RestrictDepartmentFlag, not even an explicit nil
### GetSla

`func (o *AgreementType) GetSla() SLAReference`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *AgreementType) GetSlaOk() (*SLAReference, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *AgreementType) SetSla(v SLAReference)`

SetSla sets Sla field to given value.

### HasSla

`func (o *AgreementType) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetApplicationUnits

`func (o *AgreementType) GetApplicationUnits() string`

GetApplicationUnits returns the ApplicationUnits field if non-nil, zero value otherwise.

### GetApplicationUnitsOk

`func (o *AgreementType) GetApplicationUnitsOk() (*string, bool)`

GetApplicationUnitsOk returns a tuple with the ApplicationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUnits

`func (o *AgreementType) SetApplicationUnits(v string)`

SetApplicationUnits sets ApplicationUnits field to given value.

### HasApplicationUnits

`func (o *AgreementType) HasApplicationUnits() bool`

HasApplicationUnits returns a boolean if a field has been set.

### SetApplicationUnitsNil

`func (o *AgreementType) SetApplicationUnitsNil(b bool)`

 SetApplicationUnitsNil sets the value for ApplicationUnits to be an explicit nil

### UnsetApplicationUnits
`func (o *AgreementType) UnsetApplicationUnits()`

UnsetApplicationUnits ensures that no value is present for ApplicationUnits, not even an explicit nil
### GetApplicationLimit

`func (o *AgreementType) GetApplicationLimit() float64`

GetApplicationLimit returns the ApplicationLimit field if non-nil, zero value otherwise.

### GetApplicationLimitOk

`func (o *AgreementType) GetApplicationLimitOk() (*float64, bool)`

GetApplicationLimitOk returns a tuple with the ApplicationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLimit

`func (o *AgreementType) SetApplicationLimit(v float64)`

SetApplicationLimit sets ApplicationLimit field to given value.

### HasApplicationLimit

`func (o *AgreementType) HasApplicationLimit() bool`

HasApplicationLimit returns a boolean if a field has been set.

### SetApplicationLimitNil

`func (o *AgreementType) SetApplicationLimitNil(b bool)`

 SetApplicationLimitNil sets the value for ApplicationLimit to be an explicit nil

### UnsetApplicationLimit
`func (o *AgreementType) UnsetApplicationLimit()`

UnsetApplicationLimit ensures that no value is present for ApplicationLimit, not even an explicit nil
### GetApplicationCycle

`func (o *AgreementType) GetApplicationCycle() string`

GetApplicationCycle returns the ApplicationCycle field if non-nil, zero value otherwise.

### GetApplicationCycleOk

`func (o *AgreementType) GetApplicationCycleOk() (*string, bool)`

GetApplicationCycleOk returns a tuple with the ApplicationCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCycle

`func (o *AgreementType) SetApplicationCycle(v string)`

SetApplicationCycle sets ApplicationCycle field to given value.

### HasApplicationCycle

`func (o *AgreementType) HasApplicationCycle() bool`

HasApplicationCycle returns a boolean if a field has been set.

### SetApplicationCycleNil

`func (o *AgreementType) SetApplicationCycleNil(b bool)`

 SetApplicationCycleNil sets the value for ApplicationCycle to be an explicit nil

### UnsetApplicationCycle
`func (o *AgreementType) UnsetApplicationCycle()`

UnsetApplicationCycle ensures that no value is present for ApplicationCycle, not even an explicit nil
### GetApplicationUnlimitedFlag

`func (o *AgreementType) GetApplicationUnlimitedFlag() bool`

GetApplicationUnlimitedFlag returns the ApplicationUnlimitedFlag field if non-nil, zero value otherwise.

### GetApplicationUnlimitedFlagOk

`func (o *AgreementType) GetApplicationUnlimitedFlagOk() (*bool, bool)`

GetApplicationUnlimitedFlagOk returns a tuple with the ApplicationUnlimitedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUnlimitedFlag

`func (o *AgreementType) SetApplicationUnlimitedFlag(v bool)`

SetApplicationUnlimitedFlag sets ApplicationUnlimitedFlag field to given value.

### HasApplicationUnlimitedFlag

`func (o *AgreementType) HasApplicationUnlimitedFlag() bool`

HasApplicationUnlimitedFlag returns a boolean if a field has been set.

### SetApplicationUnlimitedFlagNil

`func (o *AgreementType) SetApplicationUnlimitedFlagNil(b bool)`

 SetApplicationUnlimitedFlagNil sets the value for ApplicationUnlimitedFlag to be an explicit nil

### UnsetApplicationUnlimitedFlag
`func (o *AgreementType) UnsetApplicationUnlimitedFlag()`

UnsetApplicationUnlimitedFlag ensures that no value is present for ApplicationUnlimitedFlag, not even an explicit nil
### GetOneTimeFlag

`func (o *AgreementType) GetOneTimeFlag() bool`

GetOneTimeFlag returns the OneTimeFlag field if non-nil, zero value otherwise.

### GetOneTimeFlagOk

`func (o *AgreementType) GetOneTimeFlagOk() (*bool, bool)`

GetOneTimeFlagOk returns a tuple with the OneTimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimeFlag

`func (o *AgreementType) SetOneTimeFlag(v bool)`

SetOneTimeFlag sets OneTimeFlag field to given value.

### HasOneTimeFlag

`func (o *AgreementType) HasOneTimeFlag() bool`

HasOneTimeFlag returns a boolean if a field has been set.

### SetOneTimeFlagNil

`func (o *AgreementType) SetOneTimeFlagNil(b bool)`

 SetOneTimeFlagNil sets the value for OneTimeFlag to be an explicit nil

### UnsetOneTimeFlag
`func (o *AgreementType) UnsetOneTimeFlag()`

UnsetOneTimeFlag ensures that no value is present for OneTimeFlag, not even an explicit nil
### GetCoverAgreementTimeFlag

`func (o *AgreementType) GetCoverAgreementTimeFlag() bool`

GetCoverAgreementTimeFlag returns the CoverAgreementTimeFlag field if non-nil, zero value otherwise.

### GetCoverAgreementTimeFlagOk

`func (o *AgreementType) GetCoverAgreementTimeFlagOk() (*bool, bool)`

GetCoverAgreementTimeFlagOk returns a tuple with the CoverAgreementTimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverAgreementTimeFlag

`func (o *AgreementType) SetCoverAgreementTimeFlag(v bool)`

SetCoverAgreementTimeFlag sets CoverAgreementTimeFlag field to given value.

### HasCoverAgreementTimeFlag

`func (o *AgreementType) HasCoverAgreementTimeFlag() bool`

HasCoverAgreementTimeFlag returns a boolean if a field has been set.

### SetCoverAgreementTimeFlagNil

`func (o *AgreementType) SetCoverAgreementTimeFlagNil(b bool)`

 SetCoverAgreementTimeFlagNil sets the value for CoverAgreementTimeFlag to be an explicit nil

### UnsetCoverAgreementTimeFlag
`func (o *AgreementType) UnsetCoverAgreementTimeFlag()`

UnsetCoverAgreementTimeFlag ensures that no value is present for CoverAgreementTimeFlag, not even an explicit nil
### GetCoverAgreementProductFlag

`func (o *AgreementType) GetCoverAgreementProductFlag() bool`

GetCoverAgreementProductFlag returns the CoverAgreementProductFlag field if non-nil, zero value otherwise.

### GetCoverAgreementProductFlagOk

`func (o *AgreementType) GetCoverAgreementProductFlagOk() (*bool, bool)`

GetCoverAgreementProductFlagOk returns a tuple with the CoverAgreementProductFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverAgreementProductFlag

`func (o *AgreementType) SetCoverAgreementProductFlag(v bool)`

SetCoverAgreementProductFlag sets CoverAgreementProductFlag field to given value.

### HasCoverAgreementProductFlag

`func (o *AgreementType) HasCoverAgreementProductFlag() bool`

HasCoverAgreementProductFlag returns a boolean if a field has been set.

### SetCoverAgreementProductFlagNil

`func (o *AgreementType) SetCoverAgreementProductFlagNil(b bool)`

 SetCoverAgreementProductFlagNil sets the value for CoverAgreementProductFlag to be an explicit nil

### UnsetCoverAgreementProductFlag
`func (o *AgreementType) UnsetCoverAgreementProductFlag()`

UnsetCoverAgreementProductFlag ensures that no value is present for CoverAgreementProductFlag, not even an explicit nil
### GetCoverAgreementExpenseFlag

`func (o *AgreementType) GetCoverAgreementExpenseFlag() bool`

GetCoverAgreementExpenseFlag returns the CoverAgreementExpenseFlag field if non-nil, zero value otherwise.

### GetCoverAgreementExpenseFlagOk

`func (o *AgreementType) GetCoverAgreementExpenseFlagOk() (*bool, bool)`

GetCoverAgreementExpenseFlagOk returns a tuple with the CoverAgreementExpenseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverAgreementExpenseFlag

`func (o *AgreementType) SetCoverAgreementExpenseFlag(v bool)`

SetCoverAgreementExpenseFlag sets CoverAgreementExpenseFlag field to given value.

### HasCoverAgreementExpenseFlag

`func (o *AgreementType) HasCoverAgreementExpenseFlag() bool`

HasCoverAgreementExpenseFlag returns a boolean if a field has been set.

### SetCoverAgreementExpenseFlagNil

`func (o *AgreementType) SetCoverAgreementExpenseFlagNil(b bool)`

 SetCoverAgreementExpenseFlagNil sets the value for CoverAgreementExpenseFlag to be an explicit nil

### UnsetCoverAgreementExpenseFlag
`func (o *AgreementType) UnsetCoverAgreementExpenseFlag()`

UnsetCoverAgreementExpenseFlag ensures that no value is present for CoverAgreementExpenseFlag, not even an explicit nil
### GetCoverSalesTaxFlag

`func (o *AgreementType) GetCoverSalesTaxFlag() bool`

GetCoverSalesTaxFlag returns the CoverSalesTaxFlag field if non-nil, zero value otherwise.

### GetCoverSalesTaxFlagOk

`func (o *AgreementType) GetCoverSalesTaxFlagOk() (*bool, bool)`

GetCoverSalesTaxFlagOk returns a tuple with the CoverSalesTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverSalesTaxFlag

`func (o *AgreementType) SetCoverSalesTaxFlag(v bool)`

SetCoverSalesTaxFlag sets CoverSalesTaxFlag field to given value.

### HasCoverSalesTaxFlag

`func (o *AgreementType) HasCoverSalesTaxFlag() bool`

HasCoverSalesTaxFlag returns a boolean if a field has been set.

### SetCoverSalesTaxFlagNil

`func (o *AgreementType) SetCoverSalesTaxFlagNil(b bool)`

 SetCoverSalesTaxFlagNil sets the value for CoverSalesTaxFlag to be an explicit nil

### UnsetCoverSalesTaxFlag
`func (o *AgreementType) UnsetCoverSalesTaxFlag()`

UnsetCoverSalesTaxFlag ensures that no value is present for CoverSalesTaxFlag, not even an explicit nil
### GetCarryOverUnusedFlag

`func (o *AgreementType) GetCarryOverUnusedFlag() bool`

GetCarryOverUnusedFlag returns the CarryOverUnusedFlag field if non-nil, zero value otherwise.

### GetCarryOverUnusedFlagOk

`func (o *AgreementType) GetCarryOverUnusedFlagOk() (*bool, bool)`

GetCarryOverUnusedFlagOk returns a tuple with the CarryOverUnusedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarryOverUnusedFlag

`func (o *AgreementType) SetCarryOverUnusedFlag(v bool)`

SetCarryOverUnusedFlag sets CarryOverUnusedFlag field to given value.

### HasCarryOverUnusedFlag

`func (o *AgreementType) HasCarryOverUnusedFlag() bool`

HasCarryOverUnusedFlag returns a boolean if a field has been set.

### SetCarryOverUnusedFlagNil

`func (o *AgreementType) SetCarryOverUnusedFlagNil(b bool)`

 SetCarryOverUnusedFlagNil sets the value for CarryOverUnusedFlag to be an explicit nil

### UnsetCarryOverUnusedFlag
`func (o *AgreementType) UnsetCarryOverUnusedFlag()`

UnsetCarryOverUnusedFlag ensures that no value is present for CarryOverUnusedFlag, not even an explicit nil
### GetAllowOverrunsFlag

`func (o *AgreementType) GetAllowOverrunsFlag() bool`

GetAllowOverrunsFlag returns the AllowOverrunsFlag field if non-nil, zero value otherwise.

### GetAllowOverrunsFlagOk

`func (o *AgreementType) GetAllowOverrunsFlagOk() (*bool, bool)`

GetAllowOverrunsFlagOk returns a tuple with the AllowOverrunsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverrunsFlag

`func (o *AgreementType) SetAllowOverrunsFlag(v bool)`

SetAllowOverrunsFlag sets AllowOverrunsFlag field to given value.

### HasAllowOverrunsFlag

`func (o *AgreementType) HasAllowOverrunsFlag() bool`

HasAllowOverrunsFlag returns a boolean if a field has been set.

### SetAllowOverrunsFlagNil

`func (o *AgreementType) SetAllowOverrunsFlagNil(b bool)`

 SetAllowOverrunsFlagNil sets the value for AllowOverrunsFlag to be an explicit nil

### UnsetAllowOverrunsFlag
`func (o *AgreementType) UnsetAllowOverrunsFlag()`

UnsetAllowOverrunsFlag ensures that no value is present for AllowOverrunsFlag, not even an explicit nil
### GetExpiredDays

`func (o *AgreementType) GetExpiredDays() int32`

GetExpiredDays returns the ExpiredDays field if non-nil, zero value otherwise.

### GetExpiredDaysOk

`func (o *AgreementType) GetExpiredDaysOk() (*int32, bool)`

GetExpiredDaysOk returns a tuple with the ExpiredDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredDays

`func (o *AgreementType) SetExpiredDays(v int32)`

SetExpiredDays sets ExpiredDays field to given value.

### HasExpiredDays

`func (o *AgreementType) HasExpiredDays() bool`

HasExpiredDays returns a boolean if a field has been set.

### SetExpiredDaysNil

`func (o *AgreementType) SetExpiredDaysNil(b bool)`

 SetExpiredDaysNil sets the value for ExpiredDays to be an explicit nil

### UnsetExpiredDays
`func (o *AgreementType) UnsetExpiredDays()`

UnsetExpiredDays ensures that no value is present for ExpiredDays, not even an explicit nil
### GetLimit

`func (o *AgreementType) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AgreementType) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AgreementType) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AgreementType) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *AgreementType) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *AgreementType) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetExpireWhenZero

`func (o *AgreementType) GetExpireWhenZero() bool`

GetExpireWhenZero returns the ExpireWhenZero field if non-nil, zero value otherwise.

### GetExpireWhenZeroOk

`func (o *AgreementType) GetExpireWhenZeroOk() (*bool, bool)`

GetExpireWhenZeroOk returns a tuple with the ExpireWhenZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWhenZero

`func (o *AgreementType) SetExpireWhenZero(v bool)`

SetExpireWhenZero sets ExpireWhenZero field to given value.

### HasExpireWhenZero

`func (o *AgreementType) HasExpireWhenZero() bool`

HasExpireWhenZero returns a boolean if a field has been set.

### SetExpireWhenZeroNil

`func (o *AgreementType) SetExpireWhenZeroNil(b bool)`

 SetExpireWhenZeroNil sets the value for ExpireWhenZero to be an explicit nil

### UnsetExpireWhenZero
`func (o *AgreementType) UnsetExpireWhenZero()`

UnsetExpireWhenZero ensures that no value is present for ExpireWhenZero, not even an explicit nil
### GetChargeToFirmFlag

`func (o *AgreementType) GetChargeToFirmFlag() bool`

GetChargeToFirmFlag returns the ChargeToFirmFlag field if non-nil, zero value otherwise.

### GetChargeToFirmFlagOk

`func (o *AgreementType) GetChargeToFirmFlagOk() (*bool, bool)`

GetChargeToFirmFlagOk returns a tuple with the ChargeToFirmFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeToFirmFlag

`func (o *AgreementType) SetChargeToFirmFlag(v bool)`

SetChargeToFirmFlag sets ChargeToFirmFlag field to given value.

### HasChargeToFirmFlag

`func (o *AgreementType) HasChargeToFirmFlag() bool`

HasChargeToFirmFlag returns a boolean if a field has been set.

### SetChargeToFirmFlagNil

`func (o *AgreementType) SetChargeToFirmFlagNil(b bool)`

 SetChargeToFirmFlagNil sets the value for ChargeToFirmFlag to be an explicit nil

### UnsetChargeToFirmFlag
`func (o *AgreementType) UnsetChargeToFirmFlag()`

UnsetChargeToFirmFlag ensures that no value is present for ChargeToFirmFlag, not even an explicit nil
### GetEmployeeCompRate

`func (o *AgreementType) GetEmployeeCompRate() string`

GetEmployeeCompRate returns the EmployeeCompRate field if non-nil, zero value otherwise.

### GetEmployeeCompRateOk

`func (o *AgreementType) GetEmployeeCompRateOk() (*string, bool)`

GetEmployeeCompRateOk returns a tuple with the EmployeeCompRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCompRate

`func (o *AgreementType) SetEmployeeCompRate(v string)`

SetEmployeeCompRate sets EmployeeCompRate field to given value.


### SetEmployeeCompRateNil

`func (o *AgreementType) SetEmployeeCompRateNil(b bool)`

 SetEmployeeCompRateNil sets the value for EmployeeCompRate to be an explicit nil

### UnsetEmployeeCompRate
`func (o *AgreementType) UnsetEmployeeCompRate()`

UnsetEmployeeCompRate ensures that no value is present for EmployeeCompRate, not even an explicit nil
### GetEmployeeCompNotExceed

`func (o *AgreementType) GetEmployeeCompNotExceed() string`

GetEmployeeCompNotExceed returns the EmployeeCompNotExceed field if non-nil, zero value otherwise.

### GetEmployeeCompNotExceedOk

`func (o *AgreementType) GetEmployeeCompNotExceedOk() (*string, bool)`

GetEmployeeCompNotExceedOk returns a tuple with the EmployeeCompNotExceed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCompNotExceed

`func (o *AgreementType) SetEmployeeCompNotExceed(v string)`

SetEmployeeCompNotExceed sets EmployeeCompNotExceed field to given value.


### SetEmployeeCompNotExceedNil

`func (o *AgreementType) SetEmployeeCompNotExceedNil(b bool)`

 SetEmployeeCompNotExceedNil sets the value for EmployeeCompNotExceed to be an explicit nil

### UnsetEmployeeCompNotExceed
`func (o *AgreementType) UnsetEmployeeCompNotExceed()`

UnsetEmployeeCompNotExceed ensures that no value is present for EmployeeCompNotExceed, not even an explicit nil
### GetCompHourlyRate

`func (o *AgreementType) GetCompHourlyRate() float64`

GetCompHourlyRate returns the CompHourlyRate field if non-nil, zero value otherwise.

### GetCompHourlyRateOk

`func (o *AgreementType) GetCompHourlyRateOk() (*float64, bool)`

GetCompHourlyRateOk returns a tuple with the CompHourlyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompHourlyRate

`func (o *AgreementType) SetCompHourlyRate(v float64)`

SetCompHourlyRate sets CompHourlyRate field to given value.

### HasCompHourlyRate

`func (o *AgreementType) HasCompHourlyRate() bool`

HasCompHourlyRate returns a boolean if a field has been set.

### SetCompHourlyRateNil

`func (o *AgreementType) SetCompHourlyRateNil(b bool)`

 SetCompHourlyRateNil sets the value for CompHourlyRate to be an explicit nil

### UnsetCompHourlyRate
`func (o *AgreementType) UnsetCompHourlyRate()`

UnsetCompHourlyRate ensures that no value is present for CompHourlyRate, not even an explicit nil
### GetCompLimitAmount

`func (o *AgreementType) GetCompLimitAmount() float64`

GetCompLimitAmount returns the CompLimitAmount field if non-nil, zero value otherwise.

### GetCompLimitAmountOk

`func (o *AgreementType) GetCompLimitAmountOk() (*float64, bool)`

GetCompLimitAmountOk returns a tuple with the CompLimitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompLimitAmount

`func (o *AgreementType) SetCompLimitAmount(v float64)`

SetCompLimitAmount sets CompLimitAmount field to given value.

### HasCompLimitAmount

`func (o *AgreementType) HasCompLimitAmount() bool`

HasCompLimitAmount returns a boolean if a field has been set.

### SetCompLimitAmountNil

`func (o *AgreementType) SetCompLimitAmountNil(b bool)`

 SetCompLimitAmountNil sets the value for CompLimitAmount to be an explicit nil

### UnsetCompLimitAmount
`func (o *AgreementType) UnsetCompLimitAmount()`

UnsetCompLimitAmount ensures that no value is present for CompLimitAmount, not even an explicit nil
### GetBillingCycle

`func (o *AgreementType) GetBillingCycle() BillingCycleReference`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *AgreementType) GetBillingCycleOk() (*BillingCycleReference, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *AgreementType) SetBillingCycle(v BillingCycleReference)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *AgreementType) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetBillOneTimeFlag

`func (o *AgreementType) GetBillOneTimeFlag() bool`

GetBillOneTimeFlag returns the BillOneTimeFlag field if non-nil, zero value otherwise.

### GetBillOneTimeFlagOk

`func (o *AgreementType) GetBillOneTimeFlagOk() (*bool, bool)`

GetBillOneTimeFlagOk returns a tuple with the BillOneTimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOneTimeFlag

`func (o *AgreementType) SetBillOneTimeFlag(v bool)`

SetBillOneTimeFlag sets BillOneTimeFlag field to given value.

### HasBillOneTimeFlag

`func (o *AgreementType) HasBillOneTimeFlag() bool`

HasBillOneTimeFlag returns a boolean if a field has been set.

### SetBillOneTimeFlagNil

`func (o *AgreementType) SetBillOneTimeFlagNil(b bool)`

 SetBillOneTimeFlagNil sets the value for BillOneTimeFlag to be an explicit nil

### UnsetBillOneTimeFlag
`func (o *AgreementType) UnsetBillOneTimeFlag()`

UnsetBillOneTimeFlag ensures that no value is present for BillOneTimeFlag, not even an explicit nil
### GetBillingTerms

`func (o *AgreementType) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *AgreementType) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *AgreementType) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *AgreementType) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetInvoicingCycle

`func (o *AgreementType) GetInvoicingCycle() string`

GetInvoicingCycle returns the InvoicingCycle field if non-nil, zero value otherwise.

### GetInvoicingCycleOk

`func (o *AgreementType) GetInvoicingCycleOk() (*string, bool)`

GetInvoicingCycleOk returns a tuple with the InvoicingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingCycle

`func (o *AgreementType) SetInvoicingCycle(v string)`

SetInvoicingCycle sets InvoicingCycle field to given value.


### SetInvoicingCycleNil

`func (o *AgreementType) SetInvoicingCycleNil(b bool)`

 SetInvoicingCycleNil sets the value for InvoicingCycle to be an explicit nil

### UnsetInvoicingCycle
`func (o *AgreementType) UnsetInvoicingCycle()`

UnsetInvoicingCycle ensures that no value is present for InvoicingCycle, not even an explicit nil
### GetBillAmount

`func (o *AgreementType) GetBillAmount() float64`

GetBillAmount returns the BillAmount field if non-nil, zero value otherwise.

### GetBillAmountOk

`func (o *AgreementType) GetBillAmountOk() (*float64, bool)`

GetBillAmountOk returns a tuple with the BillAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillAmount

`func (o *AgreementType) SetBillAmount(v float64)`

SetBillAmount sets BillAmount field to given value.

### HasBillAmount

`func (o *AgreementType) HasBillAmount() bool`

HasBillAmount returns a boolean if a field has been set.

### SetBillAmountNil

`func (o *AgreementType) SetBillAmountNil(b bool)`

 SetBillAmountNil sets the value for BillAmount to be an explicit nil

### UnsetBillAmount
`func (o *AgreementType) UnsetBillAmount()`

UnsetBillAmount ensures that no value is present for BillAmount, not even an explicit nil
### GetTaxableFlag

`func (o *AgreementType) GetTaxableFlag() bool`

GetTaxableFlag returns the TaxableFlag field if non-nil, zero value otherwise.

### GetTaxableFlagOk

`func (o *AgreementType) GetTaxableFlagOk() (*bool, bool)`

GetTaxableFlagOk returns a tuple with the TaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableFlag

`func (o *AgreementType) SetTaxableFlag(v bool)`

SetTaxableFlag sets TaxableFlag field to given value.

### HasTaxableFlag

`func (o *AgreementType) HasTaxableFlag() bool`

HasTaxableFlag returns a boolean if a field has been set.

### SetTaxableFlagNil

`func (o *AgreementType) SetTaxableFlagNil(b bool)`

 SetTaxableFlagNil sets the value for TaxableFlag to be an explicit nil

### UnsetTaxableFlag
`func (o *AgreementType) UnsetTaxableFlag()`

UnsetTaxableFlag ensures that no value is present for TaxableFlag, not even an explicit nil
### GetRestrictDownPaymentFlag

`func (o *AgreementType) GetRestrictDownPaymentFlag() bool`

GetRestrictDownPaymentFlag returns the RestrictDownPaymentFlag field if non-nil, zero value otherwise.

### GetRestrictDownPaymentFlagOk

`func (o *AgreementType) GetRestrictDownPaymentFlagOk() (*bool, bool)`

GetRestrictDownPaymentFlagOk returns a tuple with the RestrictDownPaymentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDownPaymentFlag

`func (o *AgreementType) SetRestrictDownPaymentFlag(v bool)`

SetRestrictDownPaymentFlag sets RestrictDownPaymentFlag field to given value.

### HasRestrictDownPaymentFlag

`func (o *AgreementType) HasRestrictDownPaymentFlag() bool`

HasRestrictDownPaymentFlag returns a boolean if a field has been set.

### SetRestrictDownPaymentFlagNil

`func (o *AgreementType) SetRestrictDownPaymentFlagNil(b bool)`

 SetRestrictDownPaymentFlagNil sets the value for RestrictDownPaymentFlag to be an explicit nil

### UnsetRestrictDownPaymentFlag
`func (o *AgreementType) UnsetRestrictDownPaymentFlag()`

UnsetRestrictDownPaymentFlag ensures that no value is present for RestrictDownPaymentFlag, not even an explicit nil
### GetInvoiceDescription

`func (o *AgreementType) GetInvoiceDescription() string`

GetInvoiceDescription returns the InvoiceDescription field if non-nil, zero value otherwise.

### GetInvoiceDescriptionOk

`func (o *AgreementType) GetInvoiceDescriptionOk() (*string, bool)`

GetInvoiceDescriptionOk returns a tuple with the InvoiceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDescription

`func (o *AgreementType) SetInvoiceDescription(v string)`

SetInvoiceDescription sets InvoiceDescription field to given value.

### HasInvoiceDescription

`func (o *AgreementType) HasInvoiceDescription() bool`

HasInvoiceDescription returns a boolean if a field has been set.

### GetTopCommentFlag

`func (o *AgreementType) GetTopCommentFlag() bool`

GetTopCommentFlag returns the TopCommentFlag field if non-nil, zero value otherwise.

### GetTopCommentFlagOk

`func (o *AgreementType) GetTopCommentFlagOk() (*bool, bool)`

GetTopCommentFlagOk returns a tuple with the TopCommentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopCommentFlag

`func (o *AgreementType) SetTopCommentFlag(v bool)`

SetTopCommentFlag sets TopCommentFlag field to given value.

### HasTopCommentFlag

`func (o *AgreementType) HasTopCommentFlag() bool`

HasTopCommentFlag returns a boolean if a field has been set.

### SetTopCommentFlagNil

`func (o *AgreementType) SetTopCommentFlagNil(b bool)`

 SetTopCommentFlagNil sets the value for TopCommentFlag to be an explicit nil

### UnsetTopCommentFlag
`func (o *AgreementType) UnsetTopCommentFlag()`

UnsetTopCommentFlag ensures that no value is present for TopCommentFlag, not even an explicit nil
### GetBottomCommentFlag

`func (o *AgreementType) GetBottomCommentFlag() bool`

GetBottomCommentFlag returns the BottomCommentFlag field if non-nil, zero value otherwise.

### GetBottomCommentFlagOk

`func (o *AgreementType) GetBottomCommentFlagOk() (*bool, bool)`

GetBottomCommentFlagOk returns a tuple with the BottomCommentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomCommentFlag

`func (o *AgreementType) SetBottomCommentFlag(v bool)`

SetBottomCommentFlag sets BottomCommentFlag field to given value.

### HasBottomCommentFlag

`func (o *AgreementType) HasBottomCommentFlag() bool`

HasBottomCommentFlag returns a boolean if a field has been set.

### SetBottomCommentFlagNil

`func (o *AgreementType) SetBottomCommentFlagNil(b bool)`

 SetBottomCommentFlagNil sets the value for BottomCommentFlag to be an explicit nil

### UnsetBottomCommentFlag
`func (o *AgreementType) UnsetBottomCommentFlag()`

UnsetBottomCommentFlag ensures that no value is present for BottomCommentFlag, not even an explicit nil
### GetWorkRole

`func (o *AgreementType) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *AgreementType) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *AgreementType) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *AgreementType) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *AgreementType) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *AgreementType) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *AgreementType) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *AgreementType) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetProjectType

`func (o *AgreementType) GetProjectType() ProjectTypeReference`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *AgreementType) GetProjectTypeOk() (*ProjectTypeReference, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *AgreementType) SetProjectType(v ProjectTypeReference)`

SetProjectType sets ProjectType field to given value.

### HasProjectType

`func (o *AgreementType) HasProjectType() bool`

HasProjectType returns a boolean if a field has been set.

### GetInvoiceTemplate

`func (o *AgreementType) GetInvoiceTemplate() InvoiceTemplateReference`

GetInvoiceTemplate returns the InvoiceTemplate field if non-nil, zero value otherwise.

### GetInvoiceTemplateOk

`func (o *AgreementType) GetInvoiceTemplateOk() (*InvoiceTemplateReference, bool)`

GetInvoiceTemplateOk returns a tuple with the InvoiceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTemplate

`func (o *AgreementType) SetInvoiceTemplate(v InvoiceTemplateReference)`

SetInvoiceTemplate sets InvoiceTemplate field to given value.

### HasInvoiceTemplate

`func (o *AgreementType) HasInvoiceTemplate() bool`

HasInvoiceTemplate returns a boolean if a field has been set.

### GetBillTime

`func (o *AgreementType) GetBillTime() string`

GetBillTime returns the BillTime field if non-nil, zero value otherwise.

### GetBillTimeOk

`func (o *AgreementType) GetBillTimeOk() (*string, bool)`

GetBillTimeOk returns a tuple with the BillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTime

`func (o *AgreementType) SetBillTime(v string)`

SetBillTime sets BillTime field to given value.


### SetBillTimeNil

`func (o *AgreementType) SetBillTimeNil(b bool)`

 SetBillTimeNil sets the value for BillTime to be an explicit nil

### UnsetBillTime
`func (o *AgreementType) UnsetBillTime()`

UnsetBillTime ensures that no value is present for BillTime, not even an explicit nil
### GetBillExpenses

`func (o *AgreementType) GetBillExpenses() string`

GetBillExpenses returns the BillExpenses field if non-nil, zero value otherwise.

### GetBillExpensesOk

`func (o *AgreementType) GetBillExpensesOk() (*string, bool)`

GetBillExpensesOk returns a tuple with the BillExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillExpenses

`func (o *AgreementType) SetBillExpenses(v string)`

SetBillExpenses sets BillExpenses field to given value.


### SetBillExpensesNil

`func (o *AgreementType) SetBillExpensesNil(b bool)`

 SetBillExpensesNil sets the value for BillExpenses to be an explicit nil

### UnsetBillExpenses
`func (o *AgreementType) UnsetBillExpenses()`

UnsetBillExpenses ensures that no value is present for BillExpenses, not even an explicit nil
### GetBillProducts

`func (o *AgreementType) GetBillProducts() string`

GetBillProducts returns the BillProducts field if non-nil, zero value otherwise.

### GetBillProductsOk

`func (o *AgreementType) GetBillProductsOk() (*string, bool)`

GetBillProductsOk returns a tuple with the BillProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProducts

`func (o *AgreementType) SetBillProducts(v string)`

SetBillProducts sets BillProducts field to given value.


### SetBillProductsNil

`func (o *AgreementType) SetBillProductsNil(b bool)`

 SetBillProductsNil sets the value for BillProducts to be an explicit nil

### UnsetBillProducts
`func (o *AgreementType) UnsetBillProducts()`

UnsetBillProducts ensures that no value is present for BillProducts, not even an explicit nil
### GetBillableTimeInvoiceFlag

`func (o *AgreementType) GetBillableTimeInvoiceFlag() bool`

GetBillableTimeInvoiceFlag returns the BillableTimeInvoiceFlag field if non-nil, zero value otherwise.

### GetBillableTimeInvoiceFlagOk

`func (o *AgreementType) GetBillableTimeInvoiceFlagOk() (*bool, bool)`

GetBillableTimeInvoiceFlagOk returns a tuple with the BillableTimeInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableTimeInvoiceFlag

`func (o *AgreementType) SetBillableTimeInvoiceFlag(v bool)`

SetBillableTimeInvoiceFlag sets BillableTimeInvoiceFlag field to given value.

### HasBillableTimeInvoiceFlag

`func (o *AgreementType) HasBillableTimeInvoiceFlag() bool`

HasBillableTimeInvoiceFlag returns a boolean if a field has been set.

### SetBillableTimeInvoiceFlagNil

`func (o *AgreementType) SetBillableTimeInvoiceFlagNil(b bool)`

 SetBillableTimeInvoiceFlagNil sets the value for BillableTimeInvoiceFlag to be an explicit nil

### UnsetBillableTimeInvoiceFlag
`func (o *AgreementType) UnsetBillableTimeInvoiceFlag()`

UnsetBillableTimeInvoiceFlag ensures that no value is present for BillableTimeInvoiceFlag, not even an explicit nil
### GetBillableExpenseInvoiceFlag

`func (o *AgreementType) GetBillableExpenseInvoiceFlag() bool`

GetBillableExpenseInvoiceFlag returns the BillableExpenseInvoiceFlag field if non-nil, zero value otherwise.

### GetBillableExpenseInvoiceFlagOk

`func (o *AgreementType) GetBillableExpenseInvoiceFlagOk() (*bool, bool)`

GetBillableExpenseInvoiceFlagOk returns a tuple with the BillableExpenseInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableExpenseInvoiceFlag

`func (o *AgreementType) SetBillableExpenseInvoiceFlag(v bool)`

SetBillableExpenseInvoiceFlag sets BillableExpenseInvoiceFlag field to given value.

### HasBillableExpenseInvoiceFlag

`func (o *AgreementType) HasBillableExpenseInvoiceFlag() bool`

HasBillableExpenseInvoiceFlag returns a boolean if a field has been set.

### SetBillableExpenseInvoiceFlagNil

`func (o *AgreementType) SetBillableExpenseInvoiceFlagNil(b bool)`

 SetBillableExpenseInvoiceFlagNil sets the value for BillableExpenseInvoiceFlag to be an explicit nil

### UnsetBillableExpenseInvoiceFlag
`func (o *AgreementType) UnsetBillableExpenseInvoiceFlag()`

UnsetBillableExpenseInvoiceFlag ensures that no value is present for BillableExpenseInvoiceFlag, not even an explicit nil
### GetBillableProductInvoiceFlag

`func (o *AgreementType) GetBillableProductInvoiceFlag() bool`

GetBillableProductInvoiceFlag returns the BillableProductInvoiceFlag field if non-nil, zero value otherwise.

### GetBillableProductInvoiceFlagOk

`func (o *AgreementType) GetBillableProductInvoiceFlagOk() (*bool, bool)`

GetBillableProductInvoiceFlagOk returns a tuple with the BillableProductInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableProductInvoiceFlag

`func (o *AgreementType) SetBillableProductInvoiceFlag(v bool)`

SetBillableProductInvoiceFlag sets BillableProductInvoiceFlag field to given value.

### HasBillableProductInvoiceFlag

`func (o *AgreementType) HasBillableProductInvoiceFlag() bool`

HasBillableProductInvoiceFlag returns a boolean if a field has been set.

### SetBillableProductInvoiceFlagNil

`func (o *AgreementType) SetBillableProductInvoiceFlagNil(b bool)`

 SetBillableProductInvoiceFlagNil sets the value for BillableProductInvoiceFlag to be an explicit nil

### UnsetBillableProductInvoiceFlag
`func (o *AgreementType) UnsetBillableProductInvoiceFlag()`

UnsetBillableProductInvoiceFlag ensures that no value is present for BillableProductInvoiceFlag, not even an explicit nil
### GetCopyWorkRolesFlag

`func (o *AgreementType) GetCopyWorkRolesFlag() bool`

GetCopyWorkRolesFlag returns the CopyWorkRolesFlag field if non-nil, zero value otherwise.

### GetCopyWorkRolesFlagOk

`func (o *AgreementType) GetCopyWorkRolesFlagOk() (*bool, bool)`

GetCopyWorkRolesFlagOk returns a tuple with the CopyWorkRolesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyWorkRolesFlag

`func (o *AgreementType) SetCopyWorkRolesFlag(v bool)`

SetCopyWorkRolesFlag sets CopyWorkRolesFlag field to given value.

### HasCopyWorkRolesFlag

`func (o *AgreementType) HasCopyWorkRolesFlag() bool`

HasCopyWorkRolesFlag returns a boolean if a field has been set.

### SetCopyWorkRolesFlagNil

`func (o *AgreementType) SetCopyWorkRolesFlagNil(b bool)`

 SetCopyWorkRolesFlagNil sets the value for CopyWorkRolesFlag to be an explicit nil

### UnsetCopyWorkRolesFlag
`func (o *AgreementType) UnsetCopyWorkRolesFlag()`

UnsetCopyWorkRolesFlag ensures that no value is present for CopyWorkRolesFlag, not even an explicit nil
### GetCopyWorkTypesFlag

`func (o *AgreementType) GetCopyWorkTypesFlag() bool`

GetCopyWorkTypesFlag returns the CopyWorkTypesFlag field if non-nil, zero value otherwise.

### GetCopyWorkTypesFlagOk

`func (o *AgreementType) GetCopyWorkTypesFlagOk() (*bool, bool)`

GetCopyWorkTypesFlagOk returns a tuple with the CopyWorkTypesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyWorkTypesFlag

`func (o *AgreementType) SetCopyWorkTypesFlag(v bool)`

SetCopyWorkTypesFlag sets CopyWorkTypesFlag field to given value.

### HasCopyWorkTypesFlag

`func (o *AgreementType) HasCopyWorkTypesFlag() bool`

HasCopyWorkTypesFlag returns a boolean if a field has been set.

### SetCopyWorkTypesFlagNil

`func (o *AgreementType) SetCopyWorkTypesFlagNil(b bool)`

 SetCopyWorkTypesFlagNil sets the value for CopyWorkTypesFlag to be an explicit nil

### UnsetCopyWorkTypesFlag
`func (o *AgreementType) UnsetCopyWorkTypesFlag()`

UnsetCopyWorkTypesFlag ensures that no value is present for CopyWorkTypesFlag, not even an explicit nil
### GetExclusionWorkRoleIds

`func (o *AgreementType) GetExclusionWorkRoleIds() []int32`

GetExclusionWorkRoleIds returns the ExclusionWorkRoleIds field if non-nil, zero value otherwise.

### GetExclusionWorkRoleIdsOk

`func (o *AgreementType) GetExclusionWorkRoleIdsOk() (*[]int32, bool)`

GetExclusionWorkRoleIdsOk returns a tuple with the ExclusionWorkRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionWorkRoleIds

`func (o *AgreementType) SetExclusionWorkRoleIds(v []int32)`

SetExclusionWorkRoleIds sets ExclusionWorkRoleIds field to given value.

### HasExclusionWorkRoleIds

`func (o *AgreementType) HasExclusionWorkRoleIds() bool`

HasExclusionWorkRoleIds returns a boolean if a field has been set.

### GetAddAllWorkRoleExclusions

`func (o *AgreementType) GetAddAllWorkRoleExclusions() bool`

GetAddAllWorkRoleExclusions returns the AddAllWorkRoleExclusions field if non-nil, zero value otherwise.

### GetAddAllWorkRoleExclusionsOk

`func (o *AgreementType) GetAddAllWorkRoleExclusionsOk() (*bool, bool)`

GetAddAllWorkRoleExclusionsOk returns a tuple with the AddAllWorkRoleExclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllWorkRoleExclusions

`func (o *AgreementType) SetAddAllWorkRoleExclusions(v bool)`

SetAddAllWorkRoleExclusions sets AddAllWorkRoleExclusions field to given value.

### HasAddAllWorkRoleExclusions

`func (o *AgreementType) HasAddAllWorkRoleExclusions() bool`

HasAddAllWorkRoleExclusions returns a boolean if a field has been set.

### SetAddAllWorkRoleExclusionsNil

`func (o *AgreementType) SetAddAllWorkRoleExclusionsNil(b bool)`

 SetAddAllWorkRoleExclusionsNil sets the value for AddAllWorkRoleExclusions to be an explicit nil

### UnsetAddAllWorkRoleExclusions
`func (o *AgreementType) UnsetAddAllWorkRoleExclusions()`

UnsetAddAllWorkRoleExclusions ensures that no value is present for AddAllWorkRoleExclusions, not even an explicit nil
### GetRemoveAllWorkRoleExclusions

`func (o *AgreementType) GetRemoveAllWorkRoleExclusions() bool`

GetRemoveAllWorkRoleExclusions returns the RemoveAllWorkRoleExclusions field if non-nil, zero value otherwise.

### GetRemoveAllWorkRoleExclusionsOk

`func (o *AgreementType) GetRemoveAllWorkRoleExclusionsOk() (*bool, bool)`

GetRemoveAllWorkRoleExclusionsOk returns a tuple with the RemoveAllWorkRoleExclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllWorkRoleExclusions

`func (o *AgreementType) SetRemoveAllWorkRoleExclusions(v bool)`

SetRemoveAllWorkRoleExclusions sets RemoveAllWorkRoleExclusions field to given value.

### HasRemoveAllWorkRoleExclusions

`func (o *AgreementType) HasRemoveAllWorkRoleExclusions() bool`

HasRemoveAllWorkRoleExclusions returns a boolean if a field has been set.

### SetRemoveAllWorkRoleExclusionsNil

`func (o *AgreementType) SetRemoveAllWorkRoleExclusionsNil(b bool)`

 SetRemoveAllWorkRoleExclusionsNil sets the value for RemoveAllWorkRoleExclusions to be an explicit nil

### UnsetRemoveAllWorkRoleExclusions
`func (o *AgreementType) UnsetRemoveAllWorkRoleExclusions()`

UnsetRemoveAllWorkRoleExclusions ensures that no value is present for RemoveAllWorkRoleExclusions, not even an explicit nil
### GetExclusionWorkTypeIds

`func (o *AgreementType) GetExclusionWorkTypeIds() []int32`

GetExclusionWorkTypeIds returns the ExclusionWorkTypeIds field if non-nil, zero value otherwise.

### GetExclusionWorkTypeIdsOk

`func (o *AgreementType) GetExclusionWorkTypeIdsOk() (*[]int32, bool)`

GetExclusionWorkTypeIdsOk returns a tuple with the ExclusionWorkTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionWorkTypeIds

`func (o *AgreementType) SetExclusionWorkTypeIds(v []int32)`

SetExclusionWorkTypeIds sets ExclusionWorkTypeIds field to given value.

### HasExclusionWorkTypeIds

`func (o *AgreementType) HasExclusionWorkTypeIds() bool`

HasExclusionWorkTypeIds returns a boolean if a field has been set.

### GetAddAllWorkTypeExclusions

`func (o *AgreementType) GetAddAllWorkTypeExclusions() bool`

GetAddAllWorkTypeExclusions returns the AddAllWorkTypeExclusions field if non-nil, zero value otherwise.

### GetAddAllWorkTypeExclusionsOk

`func (o *AgreementType) GetAddAllWorkTypeExclusionsOk() (*bool, bool)`

GetAddAllWorkTypeExclusionsOk returns a tuple with the AddAllWorkTypeExclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllWorkTypeExclusions

`func (o *AgreementType) SetAddAllWorkTypeExclusions(v bool)`

SetAddAllWorkTypeExclusions sets AddAllWorkTypeExclusions field to given value.

### HasAddAllWorkTypeExclusions

`func (o *AgreementType) HasAddAllWorkTypeExclusions() bool`

HasAddAllWorkTypeExclusions returns a boolean if a field has been set.

### SetAddAllWorkTypeExclusionsNil

`func (o *AgreementType) SetAddAllWorkTypeExclusionsNil(b bool)`

 SetAddAllWorkTypeExclusionsNil sets the value for AddAllWorkTypeExclusions to be an explicit nil

### UnsetAddAllWorkTypeExclusions
`func (o *AgreementType) UnsetAddAllWorkTypeExclusions()`

UnsetAddAllWorkTypeExclusions ensures that no value is present for AddAllWorkTypeExclusions, not even an explicit nil
### GetRemoveAllWorkTypeExclusions

`func (o *AgreementType) GetRemoveAllWorkTypeExclusions() bool`

GetRemoveAllWorkTypeExclusions returns the RemoveAllWorkTypeExclusions field if non-nil, zero value otherwise.

### GetRemoveAllWorkTypeExclusionsOk

`func (o *AgreementType) GetRemoveAllWorkTypeExclusionsOk() (*bool, bool)`

GetRemoveAllWorkTypeExclusionsOk returns a tuple with the RemoveAllWorkTypeExclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllWorkTypeExclusions

`func (o *AgreementType) SetRemoveAllWorkTypeExclusions(v bool)`

SetRemoveAllWorkTypeExclusions sets RemoveAllWorkTypeExclusions field to given value.

### HasRemoveAllWorkTypeExclusions

`func (o *AgreementType) HasRemoveAllWorkTypeExclusions() bool`

HasRemoveAllWorkTypeExclusions returns a boolean if a field has been set.

### SetRemoveAllWorkTypeExclusionsNil

`func (o *AgreementType) SetRemoveAllWorkTypeExclusionsNil(b bool)`

 SetRemoveAllWorkTypeExclusionsNil sets the value for RemoveAllWorkTypeExclusions to be an explicit nil

### UnsetRemoveAllWorkTypeExclusions
`func (o *AgreementType) UnsetRemoveAllWorkTypeExclusions()`

UnsetRemoveAllWorkTypeExclusions ensures that no value is present for RemoveAllWorkTypeExclusions, not even an explicit nil
### GetIntegrationXRef

`func (o *AgreementType) GetIntegrationXRef() string`

GetIntegrationXRef returns the IntegrationXRef field if non-nil, zero value otherwise.

### GetIntegrationXRefOk

`func (o *AgreementType) GetIntegrationXRefOk() (*string, bool)`

GetIntegrationXRefOk returns a tuple with the IntegrationXRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationXRef

`func (o *AgreementType) SetIntegrationXRef(v string)`

SetIntegrationXRef sets IntegrationXRef field to given value.

### HasIntegrationXRef

`func (o *AgreementType) HasIntegrationXRef() bool`

HasIntegrationXRef returns a boolean if a field has been set.

### GetProrateFlag

`func (o *AgreementType) GetProrateFlag() bool`

GetProrateFlag returns the ProrateFlag field if non-nil, zero value otherwise.

### GetProrateFlagOk

`func (o *AgreementType) GetProrateFlagOk() (*bool, bool)`

GetProrateFlagOk returns a tuple with the ProrateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrateFlag

`func (o *AgreementType) SetProrateFlag(v bool)`

SetProrateFlag sets ProrateFlag field to given value.

### HasProrateFlag

`func (o *AgreementType) HasProrateFlag() bool`

HasProrateFlag returns a boolean if a field has been set.

### SetProrateFlagNil

`func (o *AgreementType) SetProrateFlagNil(b bool)`

 SetProrateFlagNil sets the value for ProrateFlag to be an explicit nil

### UnsetProrateFlag
`func (o *AgreementType) UnsetProrateFlag()`

UnsetProrateFlag ensures that no value is present for ProrateFlag, not even an explicit nil
### GetEmailTemplate

`func (o *AgreementType) GetEmailTemplate() EmailTemplateReference`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *AgreementType) GetEmailTemplateOk() (*EmailTemplateReference, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *AgreementType) SetEmailTemplate(v EmailTemplateReference)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *AgreementType) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### GetAutoInvoiceFlag

`func (o *AgreementType) GetAutoInvoiceFlag() bool`

GetAutoInvoiceFlag returns the AutoInvoiceFlag field if non-nil, zero value otherwise.

### GetAutoInvoiceFlagOk

`func (o *AgreementType) GetAutoInvoiceFlagOk() (*bool, bool)`

GetAutoInvoiceFlagOk returns a tuple with the AutoInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInvoiceFlag

`func (o *AgreementType) SetAutoInvoiceFlag(v bool)`

SetAutoInvoiceFlag sets AutoInvoiceFlag field to given value.

### HasAutoInvoiceFlag

`func (o *AgreementType) HasAutoInvoiceFlag() bool`

HasAutoInvoiceFlag returns a boolean if a field has been set.

### SetAutoInvoiceFlagNil

`func (o *AgreementType) SetAutoInvoiceFlagNil(b bool)`

 SetAutoInvoiceFlagNil sets the value for AutoInvoiceFlag to be an explicit nil

### UnsetAutoInvoiceFlag
`func (o *AgreementType) UnsetAutoInvoiceFlag()`

UnsetAutoInvoiceFlag ensures that no value is present for AutoInvoiceFlag, not even an explicit nil
### GetInvoiceProratedAdditionsFlag

`func (o *AgreementType) GetInvoiceProratedAdditionsFlag() bool`

GetInvoiceProratedAdditionsFlag returns the InvoiceProratedAdditionsFlag field if non-nil, zero value otherwise.

### GetInvoiceProratedAdditionsFlagOk

`func (o *AgreementType) GetInvoiceProratedAdditionsFlagOk() (*bool, bool)`

GetInvoiceProratedAdditionsFlagOk returns a tuple with the InvoiceProratedAdditionsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceProratedAdditionsFlag

`func (o *AgreementType) SetInvoiceProratedAdditionsFlag(v bool)`

SetInvoiceProratedAdditionsFlag sets InvoiceProratedAdditionsFlag field to given value.

### HasInvoiceProratedAdditionsFlag

`func (o *AgreementType) HasInvoiceProratedAdditionsFlag() bool`

HasInvoiceProratedAdditionsFlag returns a boolean if a field has been set.

### SetInvoiceProratedAdditionsFlagNil

`func (o *AgreementType) SetInvoiceProratedAdditionsFlagNil(b bool)`

 SetInvoiceProratedAdditionsFlagNil sets the value for InvoiceProratedAdditionsFlag to be an explicit nil

### UnsetInvoiceProratedAdditionsFlag
`func (o *AgreementType) UnsetInvoiceProratedAdditionsFlag()`

UnsetInvoiceProratedAdditionsFlag ensures that no value is present for InvoiceProratedAdditionsFlag, not even an explicit nil
### GetInfo

`func (o *AgreementType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


