# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  Max length: 30; | 
**Name** | **string** |  Max length: 50; | 
**Status** | Pointer to [**CompanyStatusReference**](CompanyStatusReference.md) |  | [optional] 
**AddressLine1** | Pointer to **string** | Gets or sets at least one address field is required -- addressLine1, addressLine2, city, state, zip and/or country. Max length: 50; | [optional] 
**AddressLine2** | Pointer to **string** | Gets or sets at least one address field is required -- addressLine1, addressLine2, city, state, zip and/or country. Max length: 50; | [optional] 
**City** | Pointer to **string** | Gets or sets at least one address field is required -- addressLine1, addressLine2, city, state, zip and/or country. Max length: 50; | [optional] 
**State** | Pointer to **string** | Gets or sets at least one address field is required -- addressLine1, addressLine2, city, state, zip and/or country. Max length: 50; | [optional] 
**Zip** | Pointer to **string** | Gets or sets at least one address field is required -- addressLine1, addressLine2, city, state, zip and/or country. Max length: 12; | [optional] 
**Country** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** |  Max length: 30; | [optional] 
**FaxNumber** | Pointer to **string** |  Max length: 30; | [optional] 
**Website** | Pointer to **string** |  Max length: 255; | [optional] 
**Territory** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Market** | Pointer to [**MarketDescriptionReference**](MarketDescriptionReference.md) |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**DefaultContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**DateAcquired** | Pointer to **time.Time** |  | [optional] 
**SicCode** | Pointer to [**SicCodeReference**](SicCodeReference.md) |  | [optional] 
**ParentCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**AnnualRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**NumberOfEmployees** | Pointer to **NullableInt32** |  | [optional] 
**YearEstablished** | Pointer to **NullableInt32** |  | [optional] 
**RevenueYear** | Pointer to **NullableInt32** |  | [optional] 
**OwnershipType** | Pointer to [**OwnershipTypeReference**](OwnershipTypeReference.md) |  | [optional] 
**TimeZoneSetup** | Pointer to [**TimeZoneSetupReference**](TimeZoneSetupReference.md) |  | [optional] 
**LeadSource** | Pointer to **string** |  Max length: 50; | [optional] 
**LeadFlag** | Pointer to **NullableBool** |  | [optional] 
**UnsubscribeFlag** | Pointer to **NullableBool** |  | [optional] 
**Calendar** | Pointer to [**CalendarReference**](CalendarReference.md) |  | [optional] 
**UserDefinedField1** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField2** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField3** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField4** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField5** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField6** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField7** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField8** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField9** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField10** | Pointer to **string** |  Max length: 50; | [optional] 
**VendorIdentifier** | Pointer to **string** |  | [optional] 
**TaxIdentifier** | Pointer to **string** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**InvoiceTemplate** | Pointer to [**InvoiceTemplateReference**](InvoiceTemplateReference.md) |  | [optional] 
**PricingSchedule** | Pointer to [**PricingScheduleReference**](PricingScheduleReference.md) |  | [optional] 
**CompanyEntityType** | Pointer to [**EntityTypeReference**](EntityTypeReference.md) |  | [optional] 
**BillToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**BillingSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**BillingContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**InvoiceDeliveryMethod** | Pointer to [**BillingDeliveryReference**](BillingDeliveryReference.md) |  | [optional] 
**InvoiceToEmailAddress** | Pointer to **string** |  | [optional] 
**InvoiceCCEmailAddress** | Pointer to **string** |  | [optional] 
**DeletedFlag** | Pointer to **NullableBool** |  | [optional] 
**DateDeleted** | Pointer to **time.Time** |  | [optional] 
**DeletedBy** | Pointer to **string** |  | [optional] 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**FacebookUrl** | Pointer to **string** |  | [optional] 
**TwitterUrl** | Pointer to **string** |  | [optional] 
**LinkedInUrl** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**TerritoryManager** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**ResellerIdentifier** | Pointer to **string** |  | [optional] 
**IsVendorFlag** | Pointer to **NullableBool** |  | [optional] 
**Types** | Pointer to [**[]CompanyTypeReference**](CompanyTypeReference.md) | Gets or sets integrer array of Company_Type_Recids to be assigned to company that can be passed in only during new company creation (post)             To update existing companies type, use the /company/companyTypeAssociations or /company/companies/{ID}/typeAssociations endpoints. | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**IntegratorTags** | Pointer to **[]string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewCompany

`func NewCompany(identifier string, name string, ) *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Company) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Company) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Company) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Company) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *Company) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Company) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Company) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *Company) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Company) GetStatus() CompanyStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Company) GetStatusOk() (*CompanyStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Company) SetStatus(v CompanyStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Company) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAddressLine1

`func (o *Company) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *Company) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *Company) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *Company) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *Company) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *Company) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *Company) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *Company) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *Company) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Company) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Company) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Company) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *Company) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Company) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Company) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Company) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *Company) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *Company) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *Company) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *Company) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetCountry

`func (o *Company) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Company) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Company) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Company) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Company) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Company) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Company) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Company) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetFaxNumber

`func (o *Company) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *Company) GetFaxNumberOk() (*string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNumber

`func (o *Company) SetFaxNumber(v string)`

SetFaxNumber sets FaxNumber field to given value.

### HasFaxNumber

`func (o *Company) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

### GetWebsite

`func (o *Company) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Company) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Company) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Company) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetTerritory

`func (o *Company) GetTerritory() SystemLocationReference`

GetTerritory returns the Territory field if non-nil, zero value otherwise.

### GetTerritoryOk

`func (o *Company) GetTerritoryOk() (*SystemLocationReference, bool)`

GetTerritoryOk returns a tuple with the Territory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritory

`func (o *Company) SetTerritory(v SystemLocationReference)`

SetTerritory sets Territory field to given value.

### HasTerritory

`func (o *Company) HasTerritory() bool`

HasTerritory returns a boolean if a field has been set.

### GetMarket

`func (o *Company) GetMarket() MarketDescriptionReference`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *Company) GetMarketOk() (*MarketDescriptionReference, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *Company) SetMarket(v MarketDescriptionReference)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *Company) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetAccountNumber

`func (o *Company) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Company) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Company) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Company) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetDefaultContact

`func (o *Company) GetDefaultContact() ContactReference`

GetDefaultContact returns the DefaultContact field if non-nil, zero value otherwise.

### GetDefaultContactOk

`func (o *Company) GetDefaultContactOk() (*ContactReference, bool)`

GetDefaultContactOk returns a tuple with the DefaultContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContact

`func (o *Company) SetDefaultContact(v ContactReference)`

SetDefaultContact sets DefaultContact field to given value.

### HasDefaultContact

`func (o *Company) HasDefaultContact() bool`

HasDefaultContact returns a boolean if a field has been set.

### GetDateAcquired

`func (o *Company) GetDateAcquired() time.Time`

GetDateAcquired returns the DateAcquired field if non-nil, zero value otherwise.

### GetDateAcquiredOk

`func (o *Company) GetDateAcquiredOk() (*time.Time, bool)`

GetDateAcquiredOk returns a tuple with the DateAcquired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAcquired

`func (o *Company) SetDateAcquired(v time.Time)`

SetDateAcquired sets DateAcquired field to given value.

### HasDateAcquired

`func (o *Company) HasDateAcquired() bool`

HasDateAcquired returns a boolean if a field has been set.

### GetSicCode

`func (o *Company) GetSicCode() SicCodeReference`

GetSicCode returns the SicCode field if non-nil, zero value otherwise.

### GetSicCodeOk

`func (o *Company) GetSicCodeOk() (*SicCodeReference, bool)`

GetSicCodeOk returns a tuple with the SicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSicCode

`func (o *Company) SetSicCode(v SicCodeReference)`

SetSicCode sets SicCode field to given value.

### HasSicCode

`func (o *Company) HasSicCode() bool`

HasSicCode returns a boolean if a field has been set.

### GetParentCompany

`func (o *Company) GetParentCompany() CompanyReference`

GetParentCompany returns the ParentCompany field if non-nil, zero value otherwise.

### GetParentCompanyOk

`func (o *Company) GetParentCompanyOk() (*CompanyReference, bool)`

GetParentCompanyOk returns a tuple with the ParentCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCompany

`func (o *Company) SetParentCompany(v CompanyReference)`

SetParentCompany sets ParentCompany field to given value.

### HasParentCompany

`func (o *Company) HasParentCompany() bool`

HasParentCompany returns a boolean if a field has been set.

### GetAnnualRevenue

`func (o *Company) GetAnnualRevenue() float64`

GetAnnualRevenue returns the AnnualRevenue field if non-nil, zero value otherwise.

### GetAnnualRevenueOk

`func (o *Company) GetAnnualRevenueOk() (*float64, bool)`

GetAnnualRevenueOk returns a tuple with the AnnualRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualRevenue

`func (o *Company) SetAnnualRevenue(v float64)`

SetAnnualRevenue sets AnnualRevenue field to given value.

### HasAnnualRevenue

`func (o *Company) HasAnnualRevenue() bool`

HasAnnualRevenue returns a boolean if a field has been set.

### SetAnnualRevenueNil

`func (o *Company) SetAnnualRevenueNil(b bool)`

 SetAnnualRevenueNil sets the value for AnnualRevenue to be an explicit nil

### UnsetAnnualRevenue
`func (o *Company) UnsetAnnualRevenue()`

UnsetAnnualRevenue ensures that no value is present for AnnualRevenue, not even an explicit nil
### GetNumberOfEmployees

`func (o *Company) GetNumberOfEmployees() int32`

GetNumberOfEmployees returns the NumberOfEmployees field if non-nil, zero value otherwise.

### GetNumberOfEmployeesOk

`func (o *Company) GetNumberOfEmployeesOk() (*int32, bool)`

GetNumberOfEmployeesOk returns a tuple with the NumberOfEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEmployees

`func (o *Company) SetNumberOfEmployees(v int32)`

SetNumberOfEmployees sets NumberOfEmployees field to given value.

### HasNumberOfEmployees

`func (o *Company) HasNumberOfEmployees() bool`

HasNumberOfEmployees returns a boolean if a field has been set.

### SetNumberOfEmployeesNil

`func (o *Company) SetNumberOfEmployeesNil(b bool)`

 SetNumberOfEmployeesNil sets the value for NumberOfEmployees to be an explicit nil

### UnsetNumberOfEmployees
`func (o *Company) UnsetNumberOfEmployees()`

UnsetNumberOfEmployees ensures that no value is present for NumberOfEmployees, not even an explicit nil
### GetYearEstablished

`func (o *Company) GetYearEstablished() int32`

GetYearEstablished returns the YearEstablished field if non-nil, zero value otherwise.

### GetYearEstablishedOk

`func (o *Company) GetYearEstablishedOk() (*int32, bool)`

GetYearEstablishedOk returns a tuple with the YearEstablished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearEstablished

`func (o *Company) SetYearEstablished(v int32)`

SetYearEstablished sets YearEstablished field to given value.

### HasYearEstablished

`func (o *Company) HasYearEstablished() bool`

HasYearEstablished returns a boolean if a field has been set.

### SetYearEstablishedNil

`func (o *Company) SetYearEstablishedNil(b bool)`

 SetYearEstablishedNil sets the value for YearEstablished to be an explicit nil

### UnsetYearEstablished
`func (o *Company) UnsetYearEstablished()`

UnsetYearEstablished ensures that no value is present for YearEstablished, not even an explicit nil
### GetRevenueYear

`func (o *Company) GetRevenueYear() int32`

GetRevenueYear returns the RevenueYear field if non-nil, zero value otherwise.

### GetRevenueYearOk

`func (o *Company) GetRevenueYearOk() (*int32, bool)`

GetRevenueYearOk returns a tuple with the RevenueYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueYear

`func (o *Company) SetRevenueYear(v int32)`

SetRevenueYear sets RevenueYear field to given value.

### HasRevenueYear

`func (o *Company) HasRevenueYear() bool`

HasRevenueYear returns a boolean if a field has been set.

### SetRevenueYearNil

`func (o *Company) SetRevenueYearNil(b bool)`

 SetRevenueYearNil sets the value for RevenueYear to be an explicit nil

### UnsetRevenueYear
`func (o *Company) UnsetRevenueYear()`

UnsetRevenueYear ensures that no value is present for RevenueYear, not even an explicit nil
### GetOwnershipType

`func (o *Company) GetOwnershipType() OwnershipTypeReference`

GetOwnershipType returns the OwnershipType field if non-nil, zero value otherwise.

### GetOwnershipTypeOk

`func (o *Company) GetOwnershipTypeOk() (*OwnershipTypeReference, bool)`

GetOwnershipTypeOk returns a tuple with the OwnershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipType

`func (o *Company) SetOwnershipType(v OwnershipTypeReference)`

SetOwnershipType sets OwnershipType field to given value.

### HasOwnershipType

`func (o *Company) HasOwnershipType() bool`

HasOwnershipType returns a boolean if a field has been set.

### GetTimeZoneSetup

`func (o *Company) GetTimeZoneSetup() TimeZoneSetupReference`

GetTimeZoneSetup returns the TimeZoneSetup field if non-nil, zero value otherwise.

### GetTimeZoneSetupOk

`func (o *Company) GetTimeZoneSetupOk() (*TimeZoneSetupReference, bool)`

GetTimeZoneSetupOk returns a tuple with the TimeZoneSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneSetup

`func (o *Company) SetTimeZoneSetup(v TimeZoneSetupReference)`

SetTimeZoneSetup sets TimeZoneSetup field to given value.

### HasTimeZoneSetup

`func (o *Company) HasTimeZoneSetup() bool`

HasTimeZoneSetup returns a boolean if a field has been set.

### GetLeadSource

`func (o *Company) GetLeadSource() string`

GetLeadSource returns the LeadSource field if non-nil, zero value otherwise.

### GetLeadSourceOk

`func (o *Company) GetLeadSourceOk() (*string, bool)`

GetLeadSourceOk returns a tuple with the LeadSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSource

`func (o *Company) SetLeadSource(v string)`

SetLeadSource sets LeadSource field to given value.

### HasLeadSource

`func (o *Company) HasLeadSource() bool`

HasLeadSource returns a boolean if a field has been set.

### GetLeadFlag

`func (o *Company) GetLeadFlag() bool`

GetLeadFlag returns the LeadFlag field if non-nil, zero value otherwise.

### GetLeadFlagOk

`func (o *Company) GetLeadFlagOk() (*bool, bool)`

GetLeadFlagOk returns a tuple with the LeadFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadFlag

`func (o *Company) SetLeadFlag(v bool)`

SetLeadFlag sets LeadFlag field to given value.

### HasLeadFlag

`func (o *Company) HasLeadFlag() bool`

HasLeadFlag returns a boolean if a field has been set.

### SetLeadFlagNil

`func (o *Company) SetLeadFlagNil(b bool)`

 SetLeadFlagNil sets the value for LeadFlag to be an explicit nil

### UnsetLeadFlag
`func (o *Company) UnsetLeadFlag()`

UnsetLeadFlag ensures that no value is present for LeadFlag, not even an explicit nil
### GetUnsubscribeFlag

`func (o *Company) GetUnsubscribeFlag() bool`

GetUnsubscribeFlag returns the UnsubscribeFlag field if non-nil, zero value otherwise.

### GetUnsubscribeFlagOk

`func (o *Company) GetUnsubscribeFlagOk() (*bool, bool)`

GetUnsubscribeFlagOk returns a tuple with the UnsubscribeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribeFlag

`func (o *Company) SetUnsubscribeFlag(v bool)`

SetUnsubscribeFlag sets UnsubscribeFlag field to given value.

### HasUnsubscribeFlag

`func (o *Company) HasUnsubscribeFlag() bool`

HasUnsubscribeFlag returns a boolean if a field has been set.

### SetUnsubscribeFlagNil

`func (o *Company) SetUnsubscribeFlagNil(b bool)`

 SetUnsubscribeFlagNil sets the value for UnsubscribeFlag to be an explicit nil

### UnsetUnsubscribeFlag
`func (o *Company) UnsetUnsubscribeFlag()`

UnsetUnsubscribeFlag ensures that no value is present for UnsubscribeFlag, not even an explicit nil
### GetCalendar

`func (o *Company) GetCalendar() CalendarReference`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *Company) GetCalendarOk() (*CalendarReference, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *Company) SetCalendar(v CalendarReference)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *Company) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### GetUserDefinedField1

`func (o *Company) GetUserDefinedField1() string`

GetUserDefinedField1 returns the UserDefinedField1 field if non-nil, zero value otherwise.

### GetUserDefinedField1Ok

`func (o *Company) GetUserDefinedField1Ok() (*string, bool)`

GetUserDefinedField1Ok returns a tuple with the UserDefinedField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField1

`func (o *Company) SetUserDefinedField1(v string)`

SetUserDefinedField1 sets UserDefinedField1 field to given value.

### HasUserDefinedField1

`func (o *Company) HasUserDefinedField1() bool`

HasUserDefinedField1 returns a boolean if a field has been set.

### GetUserDefinedField2

`func (o *Company) GetUserDefinedField2() string`

GetUserDefinedField2 returns the UserDefinedField2 field if non-nil, zero value otherwise.

### GetUserDefinedField2Ok

`func (o *Company) GetUserDefinedField2Ok() (*string, bool)`

GetUserDefinedField2Ok returns a tuple with the UserDefinedField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField2

`func (o *Company) SetUserDefinedField2(v string)`

SetUserDefinedField2 sets UserDefinedField2 field to given value.

### HasUserDefinedField2

`func (o *Company) HasUserDefinedField2() bool`

HasUserDefinedField2 returns a boolean if a field has been set.

### GetUserDefinedField3

`func (o *Company) GetUserDefinedField3() string`

GetUserDefinedField3 returns the UserDefinedField3 field if non-nil, zero value otherwise.

### GetUserDefinedField3Ok

`func (o *Company) GetUserDefinedField3Ok() (*string, bool)`

GetUserDefinedField3Ok returns a tuple with the UserDefinedField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField3

`func (o *Company) SetUserDefinedField3(v string)`

SetUserDefinedField3 sets UserDefinedField3 field to given value.

### HasUserDefinedField3

`func (o *Company) HasUserDefinedField3() bool`

HasUserDefinedField3 returns a boolean if a field has been set.

### GetUserDefinedField4

`func (o *Company) GetUserDefinedField4() string`

GetUserDefinedField4 returns the UserDefinedField4 field if non-nil, zero value otherwise.

### GetUserDefinedField4Ok

`func (o *Company) GetUserDefinedField4Ok() (*string, bool)`

GetUserDefinedField4Ok returns a tuple with the UserDefinedField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField4

`func (o *Company) SetUserDefinedField4(v string)`

SetUserDefinedField4 sets UserDefinedField4 field to given value.

### HasUserDefinedField4

`func (o *Company) HasUserDefinedField4() bool`

HasUserDefinedField4 returns a boolean if a field has been set.

### GetUserDefinedField5

`func (o *Company) GetUserDefinedField5() string`

GetUserDefinedField5 returns the UserDefinedField5 field if non-nil, zero value otherwise.

### GetUserDefinedField5Ok

`func (o *Company) GetUserDefinedField5Ok() (*string, bool)`

GetUserDefinedField5Ok returns a tuple with the UserDefinedField5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField5

`func (o *Company) SetUserDefinedField5(v string)`

SetUserDefinedField5 sets UserDefinedField5 field to given value.

### HasUserDefinedField5

`func (o *Company) HasUserDefinedField5() bool`

HasUserDefinedField5 returns a boolean if a field has been set.

### GetUserDefinedField6

`func (o *Company) GetUserDefinedField6() string`

GetUserDefinedField6 returns the UserDefinedField6 field if non-nil, zero value otherwise.

### GetUserDefinedField6Ok

`func (o *Company) GetUserDefinedField6Ok() (*string, bool)`

GetUserDefinedField6Ok returns a tuple with the UserDefinedField6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField6

`func (o *Company) SetUserDefinedField6(v string)`

SetUserDefinedField6 sets UserDefinedField6 field to given value.

### HasUserDefinedField6

`func (o *Company) HasUserDefinedField6() bool`

HasUserDefinedField6 returns a boolean if a field has been set.

### GetUserDefinedField7

`func (o *Company) GetUserDefinedField7() string`

GetUserDefinedField7 returns the UserDefinedField7 field if non-nil, zero value otherwise.

### GetUserDefinedField7Ok

`func (o *Company) GetUserDefinedField7Ok() (*string, bool)`

GetUserDefinedField7Ok returns a tuple with the UserDefinedField7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField7

`func (o *Company) SetUserDefinedField7(v string)`

SetUserDefinedField7 sets UserDefinedField7 field to given value.

### HasUserDefinedField7

`func (o *Company) HasUserDefinedField7() bool`

HasUserDefinedField7 returns a boolean if a field has been set.

### GetUserDefinedField8

`func (o *Company) GetUserDefinedField8() string`

GetUserDefinedField8 returns the UserDefinedField8 field if non-nil, zero value otherwise.

### GetUserDefinedField8Ok

`func (o *Company) GetUserDefinedField8Ok() (*string, bool)`

GetUserDefinedField8Ok returns a tuple with the UserDefinedField8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField8

`func (o *Company) SetUserDefinedField8(v string)`

SetUserDefinedField8 sets UserDefinedField8 field to given value.

### HasUserDefinedField8

`func (o *Company) HasUserDefinedField8() bool`

HasUserDefinedField8 returns a boolean if a field has been set.

### GetUserDefinedField9

`func (o *Company) GetUserDefinedField9() string`

GetUserDefinedField9 returns the UserDefinedField9 field if non-nil, zero value otherwise.

### GetUserDefinedField9Ok

`func (o *Company) GetUserDefinedField9Ok() (*string, bool)`

GetUserDefinedField9Ok returns a tuple with the UserDefinedField9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField9

`func (o *Company) SetUserDefinedField9(v string)`

SetUserDefinedField9 sets UserDefinedField9 field to given value.

### HasUserDefinedField9

`func (o *Company) HasUserDefinedField9() bool`

HasUserDefinedField9 returns a boolean if a field has been set.

### GetUserDefinedField10

`func (o *Company) GetUserDefinedField10() string`

GetUserDefinedField10 returns the UserDefinedField10 field if non-nil, zero value otherwise.

### GetUserDefinedField10Ok

`func (o *Company) GetUserDefinedField10Ok() (*string, bool)`

GetUserDefinedField10Ok returns a tuple with the UserDefinedField10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField10

`func (o *Company) SetUserDefinedField10(v string)`

SetUserDefinedField10 sets UserDefinedField10 field to given value.

### HasUserDefinedField10

`func (o *Company) HasUserDefinedField10() bool`

HasUserDefinedField10 returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *Company) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *Company) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *Company) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.

### HasVendorIdentifier

`func (o *Company) HasVendorIdentifier() bool`

HasVendorIdentifier returns a boolean if a field has been set.

### GetTaxIdentifier

`func (o *Company) GetTaxIdentifier() string`

GetTaxIdentifier returns the TaxIdentifier field if non-nil, zero value otherwise.

### GetTaxIdentifierOk

`func (o *Company) GetTaxIdentifierOk() (*string, bool)`

GetTaxIdentifierOk returns a tuple with the TaxIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentifier

`func (o *Company) SetTaxIdentifier(v string)`

SetTaxIdentifier sets TaxIdentifier field to given value.

### HasTaxIdentifier

`func (o *Company) HasTaxIdentifier() bool`

HasTaxIdentifier returns a boolean if a field has been set.

### GetTaxCode

`func (o *Company) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *Company) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *Company) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *Company) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetBillingTerms

`func (o *Company) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *Company) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *Company) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *Company) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetInvoiceTemplate

`func (o *Company) GetInvoiceTemplate() InvoiceTemplateReference`

GetInvoiceTemplate returns the InvoiceTemplate field if non-nil, zero value otherwise.

### GetInvoiceTemplateOk

`func (o *Company) GetInvoiceTemplateOk() (*InvoiceTemplateReference, bool)`

GetInvoiceTemplateOk returns a tuple with the InvoiceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTemplate

`func (o *Company) SetInvoiceTemplate(v InvoiceTemplateReference)`

SetInvoiceTemplate sets InvoiceTemplate field to given value.

### HasInvoiceTemplate

`func (o *Company) HasInvoiceTemplate() bool`

HasInvoiceTemplate returns a boolean if a field has been set.

### GetPricingSchedule

`func (o *Company) GetPricingSchedule() PricingScheduleReference`

GetPricingSchedule returns the PricingSchedule field if non-nil, zero value otherwise.

### GetPricingScheduleOk

`func (o *Company) GetPricingScheduleOk() (*PricingScheduleReference, bool)`

GetPricingScheduleOk returns a tuple with the PricingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingSchedule

`func (o *Company) SetPricingSchedule(v PricingScheduleReference)`

SetPricingSchedule sets PricingSchedule field to given value.

### HasPricingSchedule

`func (o *Company) HasPricingSchedule() bool`

HasPricingSchedule returns a boolean if a field has been set.

### GetCompanyEntityType

`func (o *Company) GetCompanyEntityType() EntityTypeReference`

GetCompanyEntityType returns the CompanyEntityType field if non-nil, zero value otherwise.

### GetCompanyEntityTypeOk

`func (o *Company) GetCompanyEntityTypeOk() (*EntityTypeReference, bool)`

GetCompanyEntityTypeOk returns a tuple with the CompanyEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyEntityType

`func (o *Company) SetCompanyEntityType(v EntityTypeReference)`

SetCompanyEntityType sets CompanyEntityType field to given value.

### HasCompanyEntityType

`func (o *Company) HasCompanyEntityType() bool`

HasCompanyEntityType returns a boolean if a field has been set.

### GetBillToCompany

`func (o *Company) GetBillToCompany() CompanyReference`

GetBillToCompany returns the BillToCompany field if non-nil, zero value otherwise.

### GetBillToCompanyOk

`func (o *Company) GetBillToCompanyOk() (*CompanyReference, bool)`

GetBillToCompanyOk returns a tuple with the BillToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToCompany

`func (o *Company) SetBillToCompany(v CompanyReference)`

SetBillToCompany sets BillToCompany field to given value.

### HasBillToCompany

`func (o *Company) HasBillToCompany() bool`

HasBillToCompany returns a boolean if a field has been set.

### GetBillingSite

`func (o *Company) GetBillingSite() SiteReference`

GetBillingSite returns the BillingSite field if non-nil, zero value otherwise.

### GetBillingSiteOk

`func (o *Company) GetBillingSiteOk() (*SiteReference, bool)`

GetBillingSiteOk returns a tuple with the BillingSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSite

`func (o *Company) SetBillingSite(v SiteReference)`

SetBillingSite sets BillingSite field to given value.

### HasBillingSite

`func (o *Company) HasBillingSite() bool`

HasBillingSite returns a boolean if a field has been set.

### GetBillingContact

`func (o *Company) GetBillingContact() ContactReference`

GetBillingContact returns the BillingContact field if non-nil, zero value otherwise.

### GetBillingContactOk

`func (o *Company) GetBillingContactOk() (*ContactReference, bool)`

GetBillingContactOk returns a tuple with the BillingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContact

`func (o *Company) SetBillingContact(v ContactReference)`

SetBillingContact sets BillingContact field to given value.

### HasBillingContact

`func (o *Company) HasBillingContact() bool`

HasBillingContact returns a boolean if a field has been set.

### GetInvoiceDeliveryMethod

`func (o *Company) GetInvoiceDeliveryMethod() BillingDeliveryReference`

GetInvoiceDeliveryMethod returns the InvoiceDeliveryMethod field if non-nil, zero value otherwise.

### GetInvoiceDeliveryMethodOk

`func (o *Company) GetInvoiceDeliveryMethodOk() (*BillingDeliveryReference, bool)`

GetInvoiceDeliveryMethodOk returns a tuple with the InvoiceDeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDeliveryMethod

`func (o *Company) SetInvoiceDeliveryMethod(v BillingDeliveryReference)`

SetInvoiceDeliveryMethod sets InvoiceDeliveryMethod field to given value.

### HasInvoiceDeliveryMethod

`func (o *Company) HasInvoiceDeliveryMethod() bool`

HasInvoiceDeliveryMethod returns a boolean if a field has been set.

### GetInvoiceToEmailAddress

`func (o *Company) GetInvoiceToEmailAddress() string`

GetInvoiceToEmailAddress returns the InvoiceToEmailAddress field if non-nil, zero value otherwise.

### GetInvoiceToEmailAddressOk

`func (o *Company) GetInvoiceToEmailAddressOk() (*string, bool)`

GetInvoiceToEmailAddressOk returns a tuple with the InvoiceToEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceToEmailAddress

`func (o *Company) SetInvoiceToEmailAddress(v string)`

SetInvoiceToEmailAddress sets InvoiceToEmailAddress field to given value.

### HasInvoiceToEmailAddress

`func (o *Company) HasInvoiceToEmailAddress() bool`

HasInvoiceToEmailAddress returns a boolean if a field has been set.

### GetInvoiceCCEmailAddress

`func (o *Company) GetInvoiceCCEmailAddress() string`

GetInvoiceCCEmailAddress returns the InvoiceCCEmailAddress field if non-nil, zero value otherwise.

### GetInvoiceCCEmailAddressOk

`func (o *Company) GetInvoiceCCEmailAddressOk() (*string, bool)`

GetInvoiceCCEmailAddressOk returns a tuple with the InvoiceCCEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCCEmailAddress

`func (o *Company) SetInvoiceCCEmailAddress(v string)`

SetInvoiceCCEmailAddress sets InvoiceCCEmailAddress field to given value.

### HasInvoiceCCEmailAddress

`func (o *Company) HasInvoiceCCEmailAddress() bool`

HasInvoiceCCEmailAddress returns a boolean if a field has been set.

### GetDeletedFlag

`func (o *Company) GetDeletedFlag() bool`

GetDeletedFlag returns the DeletedFlag field if non-nil, zero value otherwise.

### GetDeletedFlagOk

`func (o *Company) GetDeletedFlagOk() (*bool, bool)`

GetDeletedFlagOk returns a tuple with the DeletedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedFlag

`func (o *Company) SetDeletedFlag(v bool)`

SetDeletedFlag sets DeletedFlag field to given value.

### HasDeletedFlag

`func (o *Company) HasDeletedFlag() bool`

HasDeletedFlag returns a boolean if a field has been set.

### SetDeletedFlagNil

`func (o *Company) SetDeletedFlagNil(b bool)`

 SetDeletedFlagNil sets the value for DeletedFlag to be an explicit nil

### UnsetDeletedFlag
`func (o *Company) UnsetDeletedFlag()`

UnsetDeletedFlag ensures that no value is present for DeletedFlag, not even an explicit nil
### GetDateDeleted

`func (o *Company) GetDateDeleted() time.Time`

GetDateDeleted returns the DateDeleted field if non-nil, zero value otherwise.

### GetDateDeletedOk

`func (o *Company) GetDateDeletedOk() (*time.Time, bool)`

GetDateDeletedOk returns a tuple with the DateDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeleted

`func (o *Company) SetDateDeleted(v time.Time)`

SetDateDeleted sets DateDeleted field to given value.

### HasDateDeleted

`func (o *Company) HasDateDeleted() bool`

HasDateDeleted returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Company) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Company) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Company) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Company) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetMobileGuid

`func (o *Company) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *Company) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *Company) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *Company) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *Company) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *Company) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetFacebookUrl

`func (o *Company) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *Company) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *Company) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *Company) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### GetTwitterUrl

`func (o *Company) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *Company) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *Company) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *Company) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### GetLinkedInUrl

`func (o *Company) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *Company) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *Company) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *Company) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### GetCurrency

`func (o *Company) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Company) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Company) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Company) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTerritoryManager

`func (o *Company) GetTerritoryManager() MemberReference`

GetTerritoryManager returns the TerritoryManager field if non-nil, zero value otherwise.

### GetTerritoryManagerOk

`func (o *Company) GetTerritoryManagerOk() (*MemberReference, bool)`

GetTerritoryManagerOk returns a tuple with the TerritoryManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritoryManager

`func (o *Company) SetTerritoryManager(v MemberReference)`

SetTerritoryManager sets TerritoryManager field to given value.

### HasTerritoryManager

`func (o *Company) HasTerritoryManager() bool`

HasTerritoryManager returns a boolean if a field has been set.

### GetResellerIdentifier

`func (o *Company) GetResellerIdentifier() string`

GetResellerIdentifier returns the ResellerIdentifier field if non-nil, zero value otherwise.

### GetResellerIdentifierOk

`func (o *Company) GetResellerIdentifierOk() (*string, bool)`

GetResellerIdentifierOk returns a tuple with the ResellerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerIdentifier

`func (o *Company) SetResellerIdentifier(v string)`

SetResellerIdentifier sets ResellerIdentifier field to given value.

### HasResellerIdentifier

`func (o *Company) HasResellerIdentifier() bool`

HasResellerIdentifier returns a boolean if a field has been set.

### GetIsVendorFlag

`func (o *Company) GetIsVendorFlag() bool`

GetIsVendorFlag returns the IsVendorFlag field if non-nil, zero value otherwise.

### GetIsVendorFlagOk

`func (o *Company) GetIsVendorFlagOk() (*bool, bool)`

GetIsVendorFlagOk returns a tuple with the IsVendorFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVendorFlag

`func (o *Company) SetIsVendorFlag(v bool)`

SetIsVendorFlag sets IsVendorFlag field to given value.

### HasIsVendorFlag

`func (o *Company) HasIsVendorFlag() bool`

HasIsVendorFlag returns a boolean if a field has been set.

### SetIsVendorFlagNil

`func (o *Company) SetIsVendorFlagNil(b bool)`

 SetIsVendorFlagNil sets the value for IsVendorFlag to be an explicit nil

### UnsetIsVendorFlag
`func (o *Company) UnsetIsVendorFlag()`

UnsetIsVendorFlag ensures that no value is present for IsVendorFlag, not even an explicit nil
### GetTypes

`func (o *Company) GetTypes() []CompanyTypeReference`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *Company) GetTypesOk() (*[]CompanyTypeReference, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *Company) SetTypes(v []CompanyTypeReference)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *Company) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetSite

`func (o *Company) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Company) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Company) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *Company) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetIntegratorTags

`func (o *Company) GetIntegratorTags() []string`

GetIntegratorTags returns the IntegratorTags field if non-nil, zero value otherwise.

### GetIntegratorTagsOk

`func (o *Company) GetIntegratorTagsOk() (*[]string, bool)`

GetIntegratorTagsOk returns a tuple with the IntegratorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratorTags

`func (o *Company) SetIntegratorTags(v []string)`

SetIntegratorTags sets IntegratorTags field to given value.

### HasIntegratorTags

`func (o *Company) HasIntegratorTags() bool`

HasIntegratorTags returns a boolean if a field has been set.

### GetInfo

`func (o *Company) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Company) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Company) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Company) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *Company) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Company) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Company) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Company) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


