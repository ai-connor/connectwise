# ContactInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**CommunicationItems** | Pointer to [**[]ContactCommunicationItem**](ContactCommunicationItem.md) |  | [optional] 
**DefaultPhoneNbr** | Pointer to **string** |  | [optional] 
**DefaultPhoneType** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**CompanyLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Types** | Pointer to [**[]ContactTypeReference**](ContactTypeReference.md) |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Country** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**Zip** | Pointer to **string** |  | [optional] 
**Department** | Pointer to [**ContactDepartmentReference**](ContactDepartmentReference.md) |  | [optional] 
**DefaultBillingFlag** | Pointer to **NullableBool** |  | [optional] 
**FacebookUrl** | Pointer to **string** |  | [optional] 
**TwitterUrl** | Pointer to **string** |  | [optional] 
**LinkedInUrl** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewContactInfo

`func NewContactInfo() *ContactInfo`

NewContactInfo instantiates a new ContactInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactInfoWithDefaults

`func NewContactInfoWithDefaults() *ContactInfo`

NewContactInfoWithDefaults instantiates a new ContactInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContactInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *ContactInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ContactInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ContactInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ContactInfo) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ContactInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ContactInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ContactInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ContactInfo) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCommunicationItems

`func (o *ContactInfo) GetCommunicationItems() []ContactCommunicationItem`

GetCommunicationItems returns the CommunicationItems field if non-nil, zero value otherwise.

### GetCommunicationItemsOk

`func (o *ContactInfo) GetCommunicationItemsOk() (*[]ContactCommunicationItem, bool)`

GetCommunicationItemsOk returns a tuple with the CommunicationItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationItems

`func (o *ContactInfo) SetCommunicationItems(v []ContactCommunicationItem)`

SetCommunicationItems sets CommunicationItems field to given value.

### HasCommunicationItems

`func (o *ContactInfo) HasCommunicationItems() bool`

HasCommunicationItems returns a boolean if a field has been set.

### GetDefaultPhoneNbr

`func (o *ContactInfo) GetDefaultPhoneNbr() string`

GetDefaultPhoneNbr returns the DefaultPhoneNbr field if non-nil, zero value otherwise.

### GetDefaultPhoneNbrOk

`func (o *ContactInfo) GetDefaultPhoneNbrOk() (*string, bool)`

GetDefaultPhoneNbrOk returns a tuple with the DefaultPhoneNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhoneNbr

`func (o *ContactInfo) SetDefaultPhoneNbr(v string)`

SetDefaultPhoneNbr sets DefaultPhoneNbr field to given value.

### HasDefaultPhoneNbr

`func (o *ContactInfo) HasDefaultPhoneNbr() bool`

HasDefaultPhoneNbr returns a boolean if a field has been set.

### GetDefaultPhoneType

`func (o *ContactInfo) GetDefaultPhoneType() string`

GetDefaultPhoneType returns the DefaultPhoneType field if non-nil, zero value otherwise.

### GetDefaultPhoneTypeOk

`func (o *ContactInfo) GetDefaultPhoneTypeOk() (*string, bool)`

GetDefaultPhoneTypeOk returns a tuple with the DefaultPhoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhoneType

`func (o *ContactInfo) SetDefaultPhoneType(v string)`

SetDefaultPhoneType sets DefaultPhoneType field to given value.

### HasDefaultPhoneType

`func (o *ContactInfo) HasDefaultPhoneType() bool`

HasDefaultPhoneType returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *ContactInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ContactInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ContactInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ContactInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ContactInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ContactInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetCompany

`func (o *ContactInfo) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ContactInfo) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ContactInfo) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ContactInfo) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyLocation

`func (o *ContactInfo) GetCompanyLocation() SystemLocationReference`

GetCompanyLocation returns the CompanyLocation field if non-nil, zero value otherwise.

### GetCompanyLocationOk

`func (o *ContactInfo) GetCompanyLocationOk() (*SystemLocationReference, bool)`

GetCompanyLocationOk returns a tuple with the CompanyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLocation

`func (o *ContactInfo) SetCompanyLocation(v SystemLocationReference)`

SetCompanyLocation sets CompanyLocation field to given value.

### HasCompanyLocation

`func (o *ContactInfo) HasCompanyLocation() bool`

HasCompanyLocation returns a boolean if a field has been set.

### GetSite

`func (o *ContactInfo) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ContactInfo) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ContactInfo) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *ContactInfo) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *ContactInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ContactInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ContactInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ContactInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ContactInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ContactInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetTitle

`func (o *ContactInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContactInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContactInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContactInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTypes

`func (o *ContactInfo) GetTypes() []ContactTypeReference`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ContactInfo) GetTypesOk() (*[]ContactTypeReference, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ContactInfo) SetTypes(v []ContactTypeReference)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ContactInfo) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetAddressLine1

`func (o *ContactInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ContactInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ContactInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *ContactInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *ContactInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ContactInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ContactInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ContactInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *ContactInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ContactInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ContactInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ContactInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *ContactInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContactInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContactInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ContactInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *ContactInfo) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ContactInfo) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ContactInfo) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ContactInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetZip

`func (o *ContactInfo) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *ContactInfo) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *ContactInfo) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *ContactInfo) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetDepartment

`func (o *ContactInfo) GetDepartment() ContactDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ContactInfo) GetDepartmentOk() (*ContactDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ContactInfo) SetDepartment(v ContactDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ContactInfo) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDefaultBillingFlag

`func (o *ContactInfo) GetDefaultBillingFlag() bool`

GetDefaultBillingFlag returns the DefaultBillingFlag field if non-nil, zero value otherwise.

### GetDefaultBillingFlagOk

`func (o *ContactInfo) GetDefaultBillingFlagOk() (*bool, bool)`

GetDefaultBillingFlagOk returns a tuple with the DefaultBillingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBillingFlag

`func (o *ContactInfo) SetDefaultBillingFlag(v bool)`

SetDefaultBillingFlag sets DefaultBillingFlag field to given value.

### HasDefaultBillingFlag

`func (o *ContactInfo) HasDefaultBillingFlag() bool`

HasDefaultBillingFlag returns a boolean if a field has been set.

### SetDefaultBillingFlagNil

`func (o *ContactInfo) SetDefaultBillingFlagNil(b bool)`

 SetDefaultBillingFlagNil sets the value for DefaultBillingFlag to be an explicit nil

### UnsetDefaultBillingFlag
`func (o *ContactInfo) UnsetDefaultBillingFlag()`

UnsetDefaultBillingFlag ensures that no value is present for DefaultBillingFlag, not even an explicit nil
### GetFacebookUrl

`func (o *ContactInfo) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *ContactInfo) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *ContactInfo) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *ContactInfo) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### GetTwitterUrl

`func (o *ContactInfo) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *ContactInfo) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *ContactInfo) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *ContactInfo) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### GetLinkedInUrl

`func (o *ContactInfo) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *ContactInfo) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *ContactInfo) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *ContactInfo) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### GetInfo

`func (o *ContactInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ContactInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ContactInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ContactInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


