# GLExportVendor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Vendor** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**VendorNumber** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**DueDays** | Pointer to **NullableInt32** |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 

## Methods

### NewGLExportVendor

`func NewGLExportVendor() *GLExportVendor`

NewGLExportVendor instantiates a new GLExportVendor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportVendorWithDefaults

`func NewGLExportVendorWithDefaults() *GLExportVendor`

NewGLExportVendorWithDefaults instantiates a new GLExportVendor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *GLExportVendor) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GLExportVendor) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GLExportVendor) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *GLExportVendor) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetVendor

`func (o *GLExportVendor) GetVendor() CompanyReference`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GLExportVendor) GetVendorOk() (*CompanyReference, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GLExportVendor) SetVendor(v CompanyReference)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *GLExportVendor) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVendorNumber

`func (o *GLExportVendor) GetVendorNumber() string`

GetVendorNumber returns the VendorNumber field if non-nil, zero value otherwise.

### GetVendorNumberOk

`func (o *GLExportVendor) GetVendorNumberOk() (*string, bool)`

GetVendorNumberOk returns a tuple with the VendorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumber

`func (o *GLExportVendor) SetVendorNumber(v string)`

SetVendorNumber sets VendorNumber field to given value.

### HasVendorNumber

`func (o *GLExportVendor) HasVendorNumber() bool`

HasVendorNumber returns a boolean if a field has been set.

### GetCompany

`func (o *GLExportVendor) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GLExportVendor) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GLExportVendor) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GLExportVendor) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetContact

`func (o *GLExportVendor) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *GLExportVendor) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *GLExportVendor) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *GLExportVendor) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetAccountNumber

`func (o *GLExportVendor) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GLExportVendor) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GLExportVendor) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GLExportVendor) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetBillingTerms

`func (o *GLExportVendor) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *GLExportVendor) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *GLExportVendor) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *GLExportVendor) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetDueDays

`func (o *GLExportVendor) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *GLExportVendor) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *GLExportVendor) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *GLExportVendor) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### SetDueDaysNil

`func (o *GLExportVendor) SetDueDaysNil(b bool)`

 SetDueDaysNil sets the value for DueDays to be an explicit nil

### UnsetDueDays
`func (o *GLExportVendor) UnsetDueDays()`

UnsetDueDays ensures that no value is present for DueDays, not even an explicit nil
### GetSite

`func (o *GLExportVendor) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GLExportVendor) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GLExportVendor) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *GLExportVendor) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetTaxCode

`func (o *GLExportVendor) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *GLExportVendor) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *GLExportVendor) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *GLExportVendor) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


