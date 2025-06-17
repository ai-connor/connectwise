# CompanySite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**AddressLine1** | Pointer to **string** |  Max length: 50; | [optional] 
**AddressLine2** | Pointer to **string** |  Max length: 50; | [optional] 
**City** | Pointer to **string** |  Max length: 50; | [optional] 
**StateReference** | Pointer to [**StateReference**](StateReference.md) |  | [optional] 
**Zip** | Pointer to **string** |  Max length: 12; | [optional] 
**Country** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**AddressFormat** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  Max length: 30; | [optional] 
**PhoneNumberExt** | Pointer to **string** |  Max length: 30; | [optional] 
**FaxNumber** | Pointer to **string** |  Max length: 30; | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**EntityType** | Pointer to [**EntityTypeReference**](EntityTypeReference.md) |  | [optional] 
**ExpenseReimbursement** | Pointer to **NullableFloat64** |  | [optional] 
**PrimaryAddressFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultShippingFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultBillingFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultMailingFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**BillSeparateFlag** | Pointer to **NullableBool** |  | [optional] 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**Calendar** | Pointer to [**CalendarReference**](CalendarReference.md) |  | [optional] 
**TimeZone** | Pointer to [**TimeZoneSetupReference**](TimeZoneSetupReference.md) |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewCompanySite

`func NewCompanySite(name string, ) *CompanySite`

NewCompanySite instantiates a new CompanySite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanySiteWithDefaults

`func NewCompanySiteWithDefaults() *CompanySite`

NewCompanySiteWithDefaults instantiates a new CompanySite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanySite) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanySite) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanySite) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanySite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CompanySite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanySite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanySite) SetName(v string)`

SetName sets Name field to given value.


### GetAddressLine1

`func (o *CompanySite) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CompanySite) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CompanySite) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CompanySite) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *CompanySite) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CompanySite) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CompanySite) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CompanySite) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *CompanySite) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CompanySite) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CompanySite) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CompanySite) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStateReference

`func (o *CompanySite) GetStateReference() StateReference`

GetStateReference returns the StateReference field if non-nil, zero value otherwise.

### GetStateReferenceOk

`func (o *CompanySite) GetStateReferenceOk() (*StateReference, bool)`

GetStateReferenceOk returns a tuple with the StateReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReference

`func (o *CompanySite) SetStateReference(v StateReference)`

SetStateReference sets StateReference field to given value.

### HasStateReference

`func (o *CompanySite) HasStateReference() bool`

HasStateReference returns a boolean if a field has been set.

### GetZip

`func (o *CompanySite) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *CompanySite) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *CompanySite) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *CompanySite) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetCountry

`func (o *CompanySite) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CompanySite) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CompanySite) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CompanySite) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAddressFormat

`func (o *CompanySite) GetAddressFormat() string`

GetAddressFormat returns the AddressFormat field if non-nil, zero value otherwise.

### GetAddressFormatOk

`func (o *CompanySite) GetAddressFormatOk() (*string, bool)`

GetAddressFormatOk returns a tuple with the AddressFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFormat

`func (o *CompanySite) SetAddressFormat(v string)`

SetAddressFormat sets AddressFormat field to given value.

### HasAddressFormat

`func (o *CompanySite) HasAddressFormat() bool`

HasAddressFormat returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CompanySite) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CompanySite) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CompanySite) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CompanySite) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumberExt

`func (o *CompanySite) GetPhoneNumberExt() string`

GetPhoneNumberExt returns the PhoneNumberExt field if non-nil, zero value otherwise.

### GetPhoneNumberExtOk

`func (o *CompanySite) GetPhoneNumberExtOk() (*string, bool)`

GetPhoneNumberExtOk returns a tuple with the PhoneNumberExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberExt

`func (o *CompanySite) SetPhoneNumberExt(v string)`

SetPhoneNumberExt sets PhoneNumberExt field to given value.

### HasPhoneNumberExt

`func (o *CompanySite) HasPhoneNumberExt() bool`

HasPhoneNumberExt returns a boolean if a field has been set.

### GetFaxNumber

`func (o *CompanySite) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *CompanySite) GetFaxNumberOk() (*string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNumber

`func (o *CompanySite) SetFaxNumber(v string)`

SetFaxNumber sets FaxNumber field to given value.

### HasFaxNumber

`func (o *CompanySite) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

### GetTaxCode

`func (o *CompanySite) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *CompanySite) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *CompanySite) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *CompanySite) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetEntityType

`func (o *CompanySite) GetEntityType() EntityTypeReference`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CompanySite) GetEntityTypeOk() (*EntityTypeReference, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CompanySite) SetEntityType(v EntityTypeReference)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *CompanySite) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetExpenseReimbursement

`func (o *CompanySite) GetExpenseReimbursement() float64`

GetExpenseReimbursement returns the ExpenseReimbursement field if non-nil, zero value otherwise.

### GetExpenseReimbursementOk

`func (o *CompanySite) GetExpenseReimbursementOk() (*float64, bool)`

GetExpenseReimbursementOk returns a tuple with the ExpenseReimbursement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseReimbursement

`func (o *CompanySite) SetExpenseReimbursement(v float64)`

SetExpenseReimbursement sets ExpenseReimbursement field to given value.

### HasExpenseReimbursement

`func (o *CompanySite) HasExpenseReimbursement() bool`

HasExpenseReimbursement returns a boolean if a field has been set.

### SetExpenseReimbursementNil

`func (o *CompanySite) SetExpenseReimbursementNil(b bool)`

 SetExpenseReimbursementNil sets the value for ExpenseReimbursement to be an explicit nil

### UnsetExpenseReimbursement
`func (o *CompanySite) UnsetExpenseReimbursement()`

UnsetExpenseReimbursement ensures that no value is present for ExpenseReimbursement, not even an explicit nil
### GetPrimaryAddressFlag

`func (o *CompanySite) GetPrimaryAddressFlag() bool`

GetPrimaryAddressFlag returns the PrimaryAddressFlag field if non-nil, zero value otherwise.

### GetPrimaryAddressFlagOk

`func (o *CompanySite) GetPrimaryAddressFlagOk() (*bool, bool)`

GetPrimaryAddressFlagOk returns a tuple with the PrimaryAddressFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAddressFlag

`func (o *CompanySite) SetPrimaryAddressFlag(v bool)`

SetPrimaryAddressFlag sets PrimaryAddressFlag field to given value.

### HasPrimaryAddressFlag

`func (o *CompanySite) HasPrimaryAddressFlag() bool`

HasPrimaryAddressFlag returns a boolean if a field has been set.

### SetPrimaryAddressFlagNil

`func (o *CompanySite) SetPrimaryAddressFlagNil(b bool)`

 SetPrimaryAddressFlagNil sets the value for PrimaryAddressFlag to be an explicit nil

### UnsetPrimaryAddressFlag
`func (o *CompanySite) UnsetPrimaryAddressFlag()`

UnsetPrimaryAddressFlag ensures that no value is present for PrimaryAddressFlag, not even an explicit nil
### GetDefaultShippingFlag

`func (o *CompanySite) GetDefaultShippingFlag() bool`

GetDefaultShippingFlag returns the DefaultShippingFlag field if non-nil, zero value otherwise.

### GetDefaultShippingFlagOk

`func (o *CompanySite) GetDefaultShippingFlagOk() (*bool, bool)`

GetDefaultShippingFlagOk returns a tuple with the DefaultShippingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShippingFlag

`func (o *CompanySite) SetDefaultShippingFlag(v bool)`

SetDefaultShippingFlag sets DefaultShippingFlag field to given value.

### HasDefaultShippingFlag

`func (o *CompanySite) HasDefaultShippingFlag() bool`

HasDefaultShippingFlag returns a boolean if a field has been set.

### SetDefaultShippingFlagNil

`func (o *CompanySite) SetDefaultShippingFlagNil(b bool)`

 SetDefaultShippingFlagNil sets the value for DefaultShippingFlag to be an explicit nil

### UnsetDefaultShippingFlag
`func (o *CompanySite) UnsetDefaultShippingFlag()`

UnsetDefaultShippingFlag ensures that no value is present for DefaultShippingFlag, not even an explicit nil
### GetDefaultBillingFlag

`func (o *CompanySite) GetDefaultBillingFlag() bool`

GetDefaultBillingFlag returns the DefaultBillingFlag field if non-nil, zero value otherwise.

### GetDefaultBillingFlagOk

`func (o *CompanySite) GetDefaultBillingFlagOk() (*bool, bool)`

GetDefaultBillingFlagOk returns a tuple with the DefaultBillingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBillingFlag

`func (o *CompanySite) SetDefaultBillingFlag(v bool)`

SetDefaultBillingFlag sets DefaultBillingFlag field to given value.

### HasDefaultBillingFlag

`func (o *CompanySite) HasDefaultBillingFlag() bool`

HasDefaultBillingFlag returns a boolean if a field has been set.

### SetDefaultBillingFlagNil

`func (o *CompanySite) SetDefaultBillingFlagNil(b bool)`

 SetDefaultBillingFlagNil sets the value for DefaultBillingFlag to be an explicit nil

### UnsetDefaultBillingFlag
`func (o *CompanySite) UnsetDefaultBillingFlag()`

UnsetDefaultBillingFlag ensures that no value is present for DefaultBillingFlag, not even an explicit nil
### GetDefaultMailingFlag

`func (o *CompanySite) GetDefaultMailingFlag() bool`

GetDefaultMailingFlag returns the DefaultMailingFlag field if non-nil, zero value otherwise.

### GetDefaultMailingFlagOk

`func (o *CompanySite) GetDefaultMailingFlagOk() (*bool, bool)`

GetDefaultMailingFlagOk returns a tuple with the DefaultMailingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMailingFlag

`func (o *CompanySite) SetDefaultMailingFlag(v bool)`

SetDefaultMailingFlag sets DefaultMailingFlag field to given value.

### HasDefaultMailingFlag

`func (o *CompanySite) HasDefaultMailingFlag() bool`

HasDefaultMailingFlag returns a boolean if a field has been set.

### SetDefaultMailingFlagNil

`func (o *CompanySite) SetDefaultMailingFlagNil(b bool)`

 SetDefaultMailingFlagNil sets the value for DefaultMailingFlag to be an explicit nil

### UnsetDefaultMailingFlag
`func (o *CompanySite) UnsetDefaultMailingFlag()`

UnsetDefaultMailingFlag ensures that no value is present for DefaultMailingFlag, not even an explicit nil
### GetInactiveFlag

`func (o *CompanySite) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *CompanySite) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *CompanySite) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *CompanySite) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *CompanySite) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *CompanySite) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetBillSeparateFlag

`func (o *CompanySite) GetBillSeparateFlag() bool`

GetBillSeparateFlag returns the BillSeparateFlag field if non-nil, zero value otherwise.

### GetBillSeparateFlagOk

`func (o *CompanySite) GetBillSeparateFlagOk() (*bool, bool)`

GetBillSeparateFlagOk returns a tuple with the BillSeparateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillSeparateFlag

`func (o *CompanySite) SetBillSeparateFlag(v bool)`

SetBillSeparateFlag sets BillSeparateFlag field to given value.

### HasBillSeparateFlag

`func (o *CompanySite) HasBillSeparateFlag() bool`

HasBillSeparateFlag returns a boolean if a field has been set.

### SetBillSeparateFlagNil

`func (o *CompanySite) SetBillSeparateFlagNil(b bool)`

 SetBillSeparateFlagNil sets the value for BillSeparateFlag to be an explicit nil

### UnsetBillSeparateFlag
`func (o *CompanySite) UnsetBillSeparateFlag()`

UnsetBillSeparateFlag ensures that no value is present for BillSeparateFlag, not even an explicit nil
### GetMobileGuid

`func (o *CompanySite) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *CompanySite) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *CompanySite) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *CompanySite) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *CompanySite) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *CompanySite) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetCalendar

`func (o *CompanySite) GetCalendar() CalendarReference`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *CompanySite) GetCalendarOk() (*CalendarReference, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *CompanySite) SetCalendar(v CalendarReference)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *CompanySite) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### GetTimeZone

`func (o *CompanySite) GetTimeZone() TimeZoneSetupReference`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CompanySite) GetTimeZoneOk() (*TimeZoneSetupReference, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CompanySite) SetTimeZone(v TimeZoneSetupReference)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *CompanySite) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetCompany

`func (o *CompanySite) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanySite) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanySite) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CompanySite) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetInfo

`func (o *CompanySite) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanySite) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanySite) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanySite) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *CompanySite) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CompanySite) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CompanySite) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CompanySite) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


