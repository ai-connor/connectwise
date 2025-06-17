# GLExportCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**CompanyType** | Pointer to [**CompanyTypeReference**](CompanyTypeReference.md) |  | [optional] 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**BillingTermsXref** | Pointer to **string** |  | [optional] 
**DueDays** | Pointer to **NullableInt32** |  | [optional] 
**Taxable** | Pointer to **NullableBool** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**StateTaxXref** | Pointer to **string** |  | [optional] 
**CountyTaxXref** | Pointer to **string** |  | [optional] 
**CityTaxXref** | Pointer to **string** |  | [optional] 
**CountryTaxXref** | Pointer to **string** |  | [optional] 
**CompositeTaxXref** | Pointer to **string** |  | [optional] 
**StateTaxRate** | Pointer to **NullableFloat64** |  | [optional] 
**CountyTaxRate** | Pointer to **NullableFloat64** |  | [optional] 
**CityTaxRate** | Pointer to **NullableFloat64** |  | [optional] 
**CountryTaxRate** | Pointer to **NullableFloat64** |  | [optional] 
**CompositeTaxRate** | Pointer to **NullableFloat64** |  | [optional] 
**TaxGroupRate** | Pointer to **NullableFloat64** |  | [optional] 
**TaxAgencyXref** | Pointer to **string** |  | [optional] 
**StateTaxAgencyXref** | Pointer to **string** |  | [optional] 
**CountyTaxAgencyXref** | Pointer to **string** |  | [optional] 
**CityTaxAgencyXref** | Pointer to **string** |  | [optional] 
**CountryTaxAgencyXref** | Pointer to **string** |  | [optional] 
**CompositeTaxAgencyXref** | Pointer to **string** |  | [optional] 
**TaxLevels** | Pointer to [**[]GLExportCustomerTaxLevel**](GLExportCustomerTaxLevel.md) |  | [optional] 

## Methods

### NewGLExportCustomer

`func NewGLExportCustomer() *GLExportCustomer`

NewGLExportCustomer instantiates a new GLExportCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportCustomerWithDefaults

`func NewGLExportCustomerWithDefaults() *GLExportCustomer`

NewGLExportCustomerWithDefaults instantiates a new GLExportCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *GLExportCustomer) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GLExportCustomer) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GLExportCustomer) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GLExportCustomer) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyType

`func (o *GLExportCustomer) GetCompanyType() CompanyTypeReference`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *GLExportCustomer) GetCompanyTypeOk() (*CompanyTypeReference, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *GLExportCustomer) SetCompanyType(v CompanyTypeReference)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *GLExportCustomer) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### GetContact

`func (o *GLExportCustomer) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *GLExportCustomer) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *GLExportCustomer) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *GLExportCustomer) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetSite

`func (o *GLExportCustomer) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GLExportCustomer) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GLExportCustomer) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *GLExportCustomer) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAccountNumber

`func (o *GLExportCustomer) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GLExportCustomer) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GLExportCustomer) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GLExportCustomer) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetBillingTerms

`func (o *GLExportCustomer) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *GLExportCustomer) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *GLExportCustomer) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *GLExportCustomer) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetBillingTermsXref

`func (o *GLExportCustomer) GetBillingTermsXref() string`

GetBillingTermsXref returns the BillingTermsXref field if non-nil, zero value otherwise.

### GetBillingTermsXrefOk

`func (o *GLExportCustomer) GetBillingTermsXrefOk() (*string, bool)`

GetBillingTermsXrefOk returns a tuple with the BillingTermsXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTermsXref

`func (o *GLExportCustomer) SetBillingTermsXref(v string)`

SetBillingTermsXref sets BillingTermsXref field to given value.

### HasBillingTermsXref

`func (o *GLExportCustomer) HasBillingTermsXref() bool`

HasBillingTermsXref returns a boolean if a field has been set.

### GetDueDays

`func (o *GLExportCustomer) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *GLExportCustomer) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *GLExportCustomer) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *GLExportCustomer) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### SetDueDaysNil

`func (o *GLExportCustomer) SetDueDaysNil(b bool)`

 SetDueDaysNil sets the value for DueDays to be an explicit nil

### UnsetDueDays
`func (o *GLExportCustomer) UnsetDueDays()`

UnsetDueDays ensures that no value is present for DueDays, not even an explicit nil
### GetTaxable

`func (o *GLExportCustomer) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *GLExportCustomer) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *GLExportCustomer) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *GLExportCustomer) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### SetTaxableNil

`func (o *GLExportCustomer) SetTaxableNil(b bool)`

 SetTaxableNil sets the value for Taxable to be an explicit nil

### UnsetTaxable
`func (o *GLExportCustomer) UnsetTaxable()`

UnsetTaxable ensures that no value is present for Taxable, not even an explicit nil
### GetTaxCode

`func (o *GLExportCustomer) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *GLExportCustomer) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *GLExportCustomer) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *GLExportCustomer) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetCurrency

`func (o *GLExportCustomer) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GLExportCustomer) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GLExportCustomer) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GLExportCustomer) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStateTaxXref

`func (o *GLExportCustomer) GetStateTaxXref() string`

GetStateTaxXref returns the StateTaxXref field if non-nil, zero value otherwise.

### GetStateTaxXrefOk

`func (o *GLExportCustomer) GetStateTaxXrefOk() (*string, bool)`

GetStateTaxXrefOk returns a tuple with the StateTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxXref

`func (o *GLExportCustomer) SetStateTaxXref(v string)`

SetStateTaxXref sets StateTaxXref field to given value.

### HasStateTaxXref

`func (o *GLExportCustomer) HasStateTaxXref() bool`

HasStateTaxXref returns a boolean if a field has been set.

### GetCountyTaxXref

`func (o *GLExportCustomer) GetCountyTaxXref() string`

GetCountyTaxXref returns the CountyTaxXref field if non-nil, zero value otherwise.

### GetCountyTaxXrefOk

`func (o *GLExportCustomer) GetCountyTaxXrefOk() (*string, bool)`

GetCountyTaxXrefOk returns a tuple with the CountyTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxXref

`func (o *GLExportCustomer) SetCountyTaxXref(v string)`

SetCountyTaxXref sets CountyTaxXref field to given value.

### HasCountyTaxXref

`func (o *GLExportCustomer) HasCountyTaxXref() bool`

HasCountyTaxXref returns a boolean if a field has been set.

### GetCityTaxXref

`func (o *GLExportCustomer) GetCityTaxXref() string`

GetCityTaxXref returns the CityTaxXref field if non-nil, zero value otherwise.

### GetCityTaxXrefOk

`func (o *GLExportCustomer) GetCityTaxXrefOk() (*string, bool)`

GetCityTaxXrefOk returns a tuple with the CityTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxXref

`func (o *GLExportCustomer) SetCityTaxXref(v string)`

SetCityTaxXref sets CityTaxXref field to given value.

### HasCityTaxXref

`func (o *GLExportCustomer) HasCityTaxXref() bool`

HasCityTaxXref returns a boolean if a field has been set.

### GetCountryTaxXref

`func (o *GLExportCustomer) GetCountryTaxXref() string`

GetCountryTaxXref returns the CountryTaxXref field if non-nil, zero value otherwise.

### GetCountryTaxXrefOk

`func (o *GLExportCustomer) GetCountryTaxXrefOk() (*string, bool)`

GetCountryTaxXrefOk returns a tuple with the CountryTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryTaxXref

`func (o *GLExportCustomer) SetCountryTaxXref(v string)`

SetCountryTaxXref sets CountryTaxXref field to given value.

### HasCountryTaxXref

`func (o *GLExportCustomer) HasCountryTaxXref() bool`

HasCountryTaxXref returns a boolean if a field has been set.

### GetCompositeTaxXref

`func (o *GLExportCustomer) GetCompositeTaxXref() string`

GetCompositeTaxXref returns the CompositeTaxXref field if non-nil, zero value otherwise.

### GetCompositeTaxXrefOk

`func (o *GLExportCustomer) GetCompositeTaxXrefOk() (*string, bool)`

GetCompositeTaxXrefOk returns a tuple with the CompositeTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeTaxXref

`func (o *GLExportCustomer) SetCompositeTaxXref(v string)`

SetCompositeTaxXref sets CompositeTaxXref field to given value.

### HasCompositeTaxXref

`func (o *GLExportCustomer) HasCompositeTaxXref() bool`

HasCompositeTaxXref returns a boolean if a field has been set.

### GetStateTaxRate

`func (o *GLExportCustomer) GetStateTaxRate() float64`

GetStateTaxRate returns the StateTaxRate field if non-nil, zero value otherwise.

### GetStateTaxRateOk

`func (o *GLExportCustomer) GetStateTaxRateOk() (*float64, bool)`

GetStateTaxRateOk returns a tuple with the StateTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxRate

`func (o *GLExportCustomer) SetStateTaxRate(v float64)`

SetStateTaxRate sets StateTaxRate field to given value.

### HasStateTaxRate

`func (o *GLExportCustomer) HasStateTaxRate() bool`

HasStateTaxRate returns a boolean if a field has been set.

### SetStateTaxRateNil

`func (o *GLExportCustomer) SetStateTaxRateNil(b bool)`

 SetStateTaxRateNil sets the value for StateTaxRate to be an explicit nil

### UnsetStateTaxRate
`func (o *GLExportCustomer) UnsetStateTaxRate()`

UnsetStateTaxRate ensures that no value is present for StateTaxRate, not even an explicit nil
### GetCountyTaxRate

`func (o *GLExportCustomer) GetCountyTaxRate() float64`

GetCountyTaxRate returns the CountyTaxRate field if non-nil, zero value otherwise.

### GetCountyTaxRateOk

`func (o *GLExportCustomer) GetCountyTaxRateOk() (*float64, bool)`

GetCountyTaxRateOk returns a tuple with the CountyTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxRate

`func (o *GLExportCustomer) SetCountyTaxRate(v float64)`

SetCountyTaxRate sets CountyTaxRate field to given value.

### HasCountyTaxRate

`func (o *GLExportCustomer) HasCountyTaxRate() bool`

HasCountyTaxRate returns a boolean if a field has been set.

### SetCountyTaxRateNil

`func (o *GLExportCustomer) SetCountyTaxRateNil(b bool)`

 SetCountyTaxRateNil sets the value for CountyTaxRate to be an explicit nil

### UnsetCountyTaxRate
`func (o *GLExportCustomer) UnsetCountyTaxRate()`

UnsetCountyTaxRate ensures that no value is present for CountyTaxRate, not even an explicit nil
### GetCityTaxRate

`func (o *GLExportCustomer) GetCityTaxRate() float64`

GetCityTaxRate returns the CityTaxRate field if non-nil, zero value otherwise.

### GetCityTaxRateOk

`func (o *GLExportCustomer) GetCityTaxRateOk() (*float64, bool)`

GetCityTaxRateOk returns a tuple with the CityTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxRate

`func (o *GLExportCustomer) SetCityTaxRate(v float64)`

SetCityTaxRate sets CityTaxRate field to given value.

### HasCityTaxRate

`func (o *GLExportCustomer) HasCityTaxRate() bool`

HasCityTaxRate returns a boolean if a field has been set.

### SetCityTaxRateNil

`func (o *GLExportCustomer) SetCityTaxRateNil(b bool)`

 SetCityTaxRateNil sets the value for CityTaxRate to be an explicit nil

### UnsetCityTaxRate
`func (o *GLExportCustomer) UnsetCityTaxRate()`

UnsetCityTaxRate ensures that no value is present for CityTaxRate, not even an explicit nil
### GetCountryTaxRate

`func (o *GLExportCustomer) GetCountryTaxRate() float64`

GetCountryTaxRate returns the CountryTaxRate field if non-nil, zero value otherwise.

### GetCountryTaxRateOk

`func (o *GLExportCustomer) GetCountryTaxRateOk() (*float64, bool)`

GetCountryTaxRateOk returns a tuple with the CountryTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryTaxRate

`func (o *GLExportCustomer) SetCountryTaxRate(v float64)`

SetCountryTaxRate sets CountryTaxRate field to given value.

### HasCountryTaxRate

`func (o *GLExportCustomer) HasCountryTaxRate() bool`

HasCountryTaxRate returns a boolean if a field has been set.

### SetCountryTaxRateNil

`func (o *GLExportCustomer) SetCountryTaxRateNil(b bool)`

 SetCountryTaxRateNil sets the value for CountryTaxRate to be an explicit nil

### UnsetCountryTaxRate
`func (o *GLExportCustomer) UnsetCountryTaxRate()`

UnsetCountryTaxRate ensures that no value is present for CountryTaxRate, not even an explicit nil
### GetCompositeTaxRate

`func (o *GLExportCustomer) GetCompositeTaxRate() float64`

GetCompositeTaxRate returns the CompositeTaxRate field if non-nil, zero value otherwise.

### GetCompositeTaxRateOk

`func (o *GLExportCustomer) GetCompositeTaxRateOk() (*float64, bool)`

GetCompositeTaxRateOk returns a tuple with the CompositeTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeTaxRate

`func (o *GLExportCustomer) SetCompositeTaxRate(v float64)`

SetCompositeTaxRate sets CompositeTaxRate field to given value.

### HasCompositeTaxRate

`func (o *GLExportCustomer) HasCompositeTaxRate() bool`

HasCompositeTaxRate returns a boolean if a field has been set.

### SetCompositeTaxRateNil

`func (o *GLExportCustomer) SetCompositeTaxRateNil(b bool)`

 SetCompositeTaxRateNil sets the value for CompositeTaxRate to be an explicit nil

### UnsetCompositeTaxRate
`func (o *GLExportCustomer) UnsetCompositeTaxRate()`

UnsetCompositeTaxRate ensures that no value is present for CompositeTaxRate, not even an explicit nil
### GetTaxGroupRate

`func (o *GLExportCustomer) GetTaxGroupRate() float64`

GetTaxGroupRate returns the TaxGroupRate field if non-nil, zero value otherwise.

### GetTaxGroupRateOk

`func (o *GLExportCustomer) GetTaxGroupRateOk() (*float64, bool)`

GetTaxGroupRateOk returns a tuple with the TaxGroupRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxGroupRate

`func (o *GLExportCustomer) SetTaxGroupRate(v float64)`

SetTaxGroupRate sets TaxGroupRate field to given value.

### HasTaxGroupRate

`func (o *GLExportCustomer) HasTaxGroupRate() bool`

HasTaxGroupRate returns a boolean if a field has been set.

### SetTaxGroupRateNil

`func (o *GLExportCustomer) SetTaxGroupRateNil(b bool)`

 SetTaxGroupRateNil sets the value for TaxGroupRate to be an explicit nil

### UnsetTaxGroupRate
`func (o *GLExportCustomer) UnsetTaxGroupRate()`

UnsetTaxGroupRate ensures that no value is present for TaxGroupRate, not even an explicit nil
### GetTaxAgencyXref

`func (o *GLExportCustomer) GetTaxAgencyXref() string`

GetTaxAgencyXref returns the TaxAgencyXref field if non-nil, zero value otherwise.

### GetTaxAgencyXrefOk

`func (o *GLExportCustomer) GetTaxAgencyXrefOk() (*string, bool)`

GetTaxAgencyXrefOk returns a tuple with the TaxAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAgencyXref

`func (o *GLExportCustomer) SetTaxAgencyXref(v string)`

SetTaxAgencyXref sets TaxAgencyXref field to given value.

### HasTaxAgencyXref

`func (o *GLExportCustomer) HasTaxAgencyXref() bool`

HasTaxAgencyXref returns a boolean if a field has been set.

### GetStateTaxAgencyXref

`func (o *GLExportCustomer) GetStateTaxAgencyXref() string`

GetStateTaxAgencyXref returns the StateTaxAgencyXref field if non-nil, zero value otherwise.

### GetStateTaxAgencyXrefOk

`func (o *GLExportCustomer) GetStateTaxAgencyXrefOk() (*string, bool)`

GetStateTaxAgencyXrefOk returns a tuple with the StateTaxAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxAgencyXref

`func (o *GLExportCustomer) SetStateTaxAgencyXref(v string)`

SetStateTaxAgencyXref sets StateTaxAgencyXref field to given value.

### HasStateTaxAgencyXref

`func (o *GLExportCustomer) HasStateTaxAgencyXref() bool`

HasStateTaxAgencyXref returns a boolean if a field has been set.

### GetCountyTaxAgencyXref

`func (o *GLExportCustomer) GetCountyTaxAgencyXref() string`

GetCountyTaxAgencyXref returns the CountyTaxAgencyXref field if non-nil, zero value otherwise.

### GetCountyTaxAgencyXrefOk

`func (o *GLExportCustomer) GetCountyTaxAgencyXrefOk() (*string, bool)`

GetCountyTaxAgencyXrefOk returns a tuple with the CountyTaxAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxAgencyXref

`func (o *GLExportCustomer) SetCountyTaxAgencyXref(v string)`

SetCountyTaxAgencyXref sets CountyTaxAgencyXref field to given value.

### HasCountyTaxAgencyXref

`func (o *GLExportCustomer) HasCountyTaxAgencyXref() bool`

HasCountyTaxAgencyXref returns a boolean if a field has been set.

### GetCityTaxAgencyXref

`func (o *GLExportCustomer) GetCityTaxAgencyXref() string`

GetCityTaxAgencyXref returns the CityTaxAgencyXref field if non-nil, zero value otherwise.

### GetCityTaxAgencyXrefOk

`func (o *GLExportCustomer) GetCityTaxAgencyXrefOk() (*string, bool)`

GetCityTaxAgencyXrefOk returns a tuple with the CityTaxAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxAgencyXref

`func (o *GLExportCustomer) SetCityTaxAgencyXref(v string)`

SetCityTaxAgencyXref sets CityTaxAgencyXref field to given value.

### HasCityTaxAgencyXref

`func (o *GLExportCustomer) HasCityTaxAgencyXref() bool`

HasCityTaxAgencyXref returns a boolean if a field has been set.

### GetCountryTaxAgencyXref

`func (o *GLExportCustomer) GetCountryTaxAgencyXref() string`

GetCountryTaxAgencyXref returns the CountryTaxAgencyXref field if non-nil, zero value otherwise.

### GetCountryTaxAgencyXrefOk

`func (o *GLExportCustomer) GetCountryTaxAgencyXrefOk() (*string, bool)`

GetCountryTaxAgencyXrefOk returns a tuple with the CountryTaxAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryTaxAgencyXref

`func (o *GLExportCustomer) SetCountryTaxAgencyXref(v string)`

SetCountryTaxAgencyXref sets CountryTaxAgencyXref field to given value.

### HasCountryTaxAgencyXref

`func (o *GLExportCustomer) HasCountryTaxAgencyXref() bool`

HasCountryTaxAgencyXref returns a boolean if a field has been set.

### GetCompositeTaxAgencyXref

`func (o *GLExportCustomer) GetCompositeTaxAgencyXref() string`

GetCompositeTaxAgencyXref returns the CompositeTaxAgencyXref field if non-nil, zero value otherwise.

### GetCompositeTaxAgencyXrefOk

`func (o *GLExportCustomer) GetCompositeTaxAgencyXrefOk() (*string, bool)`

GetCompositeTaxAgencyXrefOk returns a tuple with the CompositeTaxAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeTaxAgencyXref

`func (o *GLExportCustomer) SetCompositeTaxAgencyXref(v string)`

SetCompositeTaxAgencyXref sets CompositeTaxAgencyXref field to given value.

### HasCompositeTaxAgencyXref

`func (o *GLExportCustomer) HasCompositeTaxAgencyXref() bool`

HasCompositeTaxAgencyXref returns a boolean if a field has been set.

### GetTaxLevels

`func (o *GLExportCustomer) GetTaxLevels() []GLExportCustomerTaxLevel`

GetTaxLevels returns the TaxLevels field if non-nil, zero value otherwise.

### GetTaxLevelsOk

`func (o *GLExportCustomer) GetTaxLevelsOk() (*[]GLExportCustomerTaxLevel, bool)`

GetTaxLevelsOk returns a tuple with the TaxLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLevels

`func (o *GLExportCustomer) SetTaxLevels(v []GLExportCustomerTaxLevel)`

SetTaxLevels sets TaxLevels field to given value.

### HasTaxLevels

`func (o *GLExportCustomer) HasTaxLevels() bool`

HasTaxLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


