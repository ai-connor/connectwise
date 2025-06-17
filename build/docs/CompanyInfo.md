# CompanyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Territory** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**DefaultContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**IsVendorFlag** | Pointer to **NullableBool** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**BillToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**BillingSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**BillingContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**DeletedFlag** | Pointer to **NullableBool** |  | [optional] 
**Types** | Pointer to [**[]CompanyTypeReference**](CompanyTypeReference.md) |  | [optional] 
**Status** | Pointer to [**CompanyStatusReference**](CompanyStatusReference.md) |  | [optional] 
**NoServiceFlag** | Pointer to **NullableBool** |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Country** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**Zip** | Pointer to **string** |  | [optional] 
**LeadFlag** | Pointer to **NullableBool** |  | [optional] 
**FaxNumber** | Pointer to **string** |  | [optional] 
**VendorIdentifier** | Pointer to **string** |  | [optional] 
**TaxIdentifier** | Pointer to **string** |  | [optional] 
**FacebookUrl** | Pointer to **string** |  | [optional] 
**TwitterUrl** | Pointer to **string** |  | [optional] 
**LinkedInUrl** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyInfo

`func NewCompanyInfo() *CompanyInfo`

NewCompanyInfo instantiates a new CompanyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyInfoWithDefaults

`func NewCompanyInfoWithDefaults() *CompanyInfo`

NewCompanyInfoWithDefaults instantiates a new CompanyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *CompanyInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CompanyInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CompanyInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CompanyInfo) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *CompanyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTerritory

`func (o *CompanyInfo) GetTerritory() SystemLocationReference`

GetTerritory returns the Territory field if non-nil, zero value otherwise.

### GetTerritoryOk

`func (o *CompanyInfo) GetTerritoryOk() (*SystemLocationReference, bool)`

GetTerritoryOk returns a tuple with the Territory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritory

`func (o *CompanyInfo) SetTerritory(v SystemLocationReference)`

SetTerritory sets Territory field to given value.

### HasTerritory

`func (o *CompanyInfo) HasTerritory() bool`

HasTerritory returns a boolean if a field has been set.

### GetDefaultContact

`func (o *CompanyInfo) GetDefaultContact() ContactReference`

GetDefaultContact returns the DefaultContact field if non-nil, zero value otherwise.

### GetDefaultContactOk

`func (o *CompanyInfo) GetDefaultContactOk() (*ContactReference, bool)`

GetDefaultContactOk returns a tuple with the DefaultContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContact

`func (o *CompanyInfo) SetDefaultContact(v ContactReference)`

SetDefaultContact sets DefaultContact field to given value.

### HasDefaultContact

`func (o *CompanyInfo) HasDefaultContact() bool`

HasDefaultContact returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CompanyInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CompanyInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CompanyInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CompanyInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCity

`func (o *CompanyInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CompanyInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CompanyInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CompanyInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetSite

`func (o *CompanyInfo) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CompanyInfo) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CompanyInfo) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *CompanyInfo) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetIsVendorFlag

`func (o *CompanyInfo) GetIsVendorFlag() bool`

GetIsVendorFlag returns the IsVendorFlag field if non-nil, zero value otherwise.

### GetIsVendorFlagOk

`func (o *CompanyInfo) GetIsVendorFlagOk() (*bool, bool)`

GetIsVendorFlagOk returns a tuple with the IsVendorFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVendorFlag

`func (o *CompanyInfo) SetIsVendorFlag(v bool)`

SetIsVendorFlag sets IsVendorFlag field to given value.

### HasIsVendorFlag

`func (o *CompanyInfo) HasIsVendorFlag() bool`

HasIsVendorFlag returns a boolean if a field has been set.

### SetIsVendorFlagNil

`func (o *CompanyInfo) SetIsVendorFlagNil(b bool)`

 SetIsVendorFlagNil sets the value for IsVendorFlag to be an explicit nil

### UnsetIsVendorFlag
`func (o *CompanyInfo) UnsetIsVendorFlag()`

UnsetIsVendorFlag ensures that no value is present for IsVendorFlag, not even an explicit nil
### GetCurrency

`func (o *CompanyInfo) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CompanyInfo) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CompanyInfo) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CompanyInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBillToCompany

`func (o *CompanyInfo) GetBillToCompany() CompanyReference`

GetBillToCompany returns the BillToCompany field if non-nil, zero value otherwise.

### GetBillToCompanyOk

`func (o *CompanyInfo) GetBillToCompanyOk() (*CompanyReference, bool)`

GetBillToCompanyOk returns a tuple with the BillToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToCompany

`func (o *CompanyInfo) SetBillToCompany(v CompanyReference)`

SetBillToCompany sets BillToCompany field to given value.

### HasBillToCompany

`func (o *CompanyInfo) HasBillToCompany() bool`

HasBillToCompany returns a boolean if a field has been set.

### GetBillingSite

`func (o *CompanyInfo) GetBillingSite() SiteReference`

GetBillingSite returns the BillingSite field if non-nil, zero value otherwise.

### GetBillingSiteOk

`func (o *CompanyInfo) GetBillingSiteOk() (*SiteReference, bool)`

GetBillingSiteOk returns a tuple with the BillingSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSite

`func (o *CompanyInfo) SetBillingSite(v SiteReference)`

SetBillingSite sets BillingSite field to given value.

### HasBillingSite

`func (o *CompanyInfo) HasBillingSite() bool`

HasBillingSite returns a boolean if a field has been set.

### GetBillingContact

`func (o *CompanyInfo) GetBillingContact() ContactReference`

GetBillingContact returns the BillingContact field if non-nil, zero value otherwise.

### GetBillingContactOk

`func (o *CompanyInfo) GetBillingContactOk() (*ContactReference, bool)`

GetBillingContactOk returns a tuple with the BillingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContact

`func (o *CompanyInfo) SetBillingContact(v ContactReference)`

SetBillingContact sets BillingContact field to given value.

### HasBillingContact

`func (o *CompanyInfo) HasBillingContact() bool`

HasBillingContact returns a boolean if a field has been set.

### GetBillingTerms

`func (o *CompanyInfo) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *CompanyInfo) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *CompanyInfo) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *CompanyInfo) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetTaxCode

`func (o *CompanyInfo) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *CompanyInfo) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *CompanyInfo) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *CompanyInfo) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetDeletedFlag

`func (o *CompanyInfo) GetDeletedFlag() bool`

GetDeletedFlag returns the DeletedFlag field if non-nil, zero value otherwise.

### GetDeletedFlagOk

`func (o *CompanyInfo) GetDeletedFlagOk() (*bool, bool)`

GetDeletedFlagOk returns a tuple with the DeletedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedFlag

`func (o *CompanyInfo) SetDeletedFlag(v bool)`

SetDeletedFlag sets DeletedFlag field to given value.

### HasDeletedFlag

`func (o *CompanyInfo) HasDeletedFlag() bool`

HasDeletedFlag returns a boolean if a field has been set.

### SetDeletedFlagNil

`func (o *CompanyInfo) SetDeletedFlagNil(b bool)`

 SetDeletedFlagNil sets the value for DeletedFlag to be an explicit nil

### UnsetDeletedFlag
`func (o *CompanyInfo) UnsetDeletedFlag()`

UnsetDeletedFlag ensures that no value is present for DeletedFlag, not even an explicit nil
### GetTypes

`func (o *CompanyInfo) GetTypes() []CompanyTypeReference`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *CompanyInfo) GetTypesOk() (*[]CompanyTypeReference, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *CompanyInfo) SetTypes(v []CompanyTypeReference)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *CompanyInfo) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetStatus

`func (o *CompanyInfo) GetStatus() CompanyStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompanyInfo) GetStatusOk() (*CompanyStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompanyInfo) SetStatus(v CompanyStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompanyInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNoServiceFlag

`func (o *CompanyInfo) GetNoServiceFlag() bool`

GetNoServiceFlag returns the NoServiceFlag field if non-nil, zero value otherwise.

### GetNoServiceFlagOk

`func (o *CompanyInfo) GetNoServiceFlagOk() (*bool, bool)`

GetNoServiceFlagOk returns a tuple with the NoServiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoServiceFlag

`func (o *CompanyInfo) SetNoServiceFlag(v bool)`

SetNoServiceFlag sets NoServiceFlag field to given value.

### HasNoServiceFlag

`func (o *CompanyInfo) HasNoServiceFlag() bool`

HasNoServiceFlag returns a boolean if a field has been set.

### SetNoServiceFlagNil

`func (o *CompanyInfo) SetNoServiceFlagNil(b bool)`

 SetNoServiceFlagNil sets the value for NoServiceFlag to be an explicit nil

### UnsetNoServiceFlag
`func (o *CompanyInfo) UnsetNoServiceFlag()`

UnsetNoServiceFlag ensures that no value is present for NoServiceFlag, not even an explicit nil
### GetAddressLine1

`func (o *CompanyInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CompanyInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CompanyInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CompanyInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *CompanyInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CompanyInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CompanyInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CompanyInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetState

`func (o *CompanyInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CompanyInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CompanyInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CompanyInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *CompanyInfo) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CompanyInfo) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CompanyInfo) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CompanyInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetZip

`func (o *CompanyInfo) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *CompanyInfo) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *CompanyInfo) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *CompanyInfo) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetLeadFlag

`func (o *CompanyInfo) GetLeadFlag() bool`

GetLeadFlag returns the LeadFlag field if non-nil, zero value otherwise.

### GetLeadFlagOk

`func (o *CompanyInfo) GetLeadFlagOk() (*bool, bool)`

GetLeadFlagOk returns a tuple with the LeadFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadFlag

`func (o *CompanyInfo) SetLeadFlag(v bool)`

SetLeadFlag sets LeadFlag field to given value.

### HasLeadFlag

`func (o *CompanyInfo) HasLeadFlag() bool`

HasLeadFlag returns a boolean if a field has been set.

### SetLeadFlagNil

`func (o *CompanyInfo) SetLeadFlagNil(b bool)`

 SetLeadFlagNil sets the value for LeadFlag to be an explicit nil

### UnsetLeadFlag
`func (o *CompanyInfo) UnsetLeadFlag()`

UnsetLeadFlag ensures that no value is present for LeadFlag, not even an explicit nil
### GetFaxNumber

`func (o *CompanyInfo) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *CompanyInfo) GetFaxNumberOk() (*string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNumber

`func (o *CompanyInfo) SetFaxNumber(v string)`

SetFaxNumber sets FaxNumber field to given value.

### HasFaxNumber

`func (o *CompanyInfo) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *CompanyInfo) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *CompanyInfo) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *CompanyInfo) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.

### HasVendorIdentifier

`func (o *CompanyInfo) HasVendorIdentifier() bool`

HasVendorIdentifier returns a boolean if a field has been set.

### GetTaxIdentifier

`func (o *CompanyInfo) GetTaxIdentifier() string`

GetTaxIdentifier returns the TaxIdentifier field if non-nil, zero value otherwise.

### GetTaxIdentifierOk

`func (o *CompanyInfo) GetTaxIdentifierOk() (*string, bool)`

GetTaxIdentifierOk returns a tuple with the TaxIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentifier

`func (o *CompanyInfo) SetTaxIdentifier(v string)`

SetTaxIdentifier sets TaxIdentifier field to given value.

### HasTaxIdentifier

`func (o *CompanyInfo) HasTaxIdentifier() bool`

HasTaxIdentifier returns a boolean if a field has been set.

### GetFacebookUrl

`func (o *CompanyInfo) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *CompanyInfo) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *CompanyInfo) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *CompanyInfo) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### GetTwitterUrl

`func (o *CompanyInfo) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *CompanyInfo) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *CompanyInfo) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *CompanyInfo) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### GetLinkedInUrl

`func (o *CompanyInfo) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *CompanyInfo) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *CompanyInfo) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *CompanyInfo) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### GetInfo

`func (o *CompanyInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


