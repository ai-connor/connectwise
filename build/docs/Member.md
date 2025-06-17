# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  Max length: 15; | 
**Password** | Pointer to **string** | ConditionallyRequired. API Member will get random password generated Max length: 60; | [optional] 
**DisableOnlineFlag** | Pointer to **NullableBool** |  | [optional] 
**LicenseClass** | **NullableString** | F &#x3D; Full Member, A &#x3D; API Member, C &#x3D; StreamlineIT Member, X &#x3D; Subcontractor Member | 
**Notes** | Pointer to **string** |  | [optional] 
**EmployeeIdentifer** | Pointer to **string** |  Max length: 10; | [optional] 
**VendorNumber** | Pointer to **string** |  | [optional] 
**EnableMobileGpsFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveDate** | Pointer to **time.Time** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**LastLogin** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**FirstName** | **string** |  Max length: 30; | 
**MiddleInitial** | Pointer to **string** |  Max length: 1; | [optional] 
**LastName** | **string** |  Max length: 30; | 
**HireDate** | **time.Time** |  | 
**Country** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**Photo** | Pointer to [**DocumentReference**](DocumentReference.md) |  | [optional] 
**OfficeEmail** | Pointer to **string** |  Max length: 250; | [optional] 
**MobileEmail** | Pointer to **string** |  Max length: 250; | [optional] 
**HomeEmail** | Pointer to **string** |  Max length: 250; | [optional] 
**DefaultEmail** | **NullableString** |  | 
**PrimaryEmail** | Pointer to **string** |  Max length: 250; | [optional] 
**OfficePhone** | Pointer to **string** |  Max length: 15; | [optional] 
**OfficeExtension** | Pointer to **string** |  Max length: 10; | [optional] 
**MobilePhone** | Pointer to **string** |  Max length: 15; | [optional] 
**MobileExtension** | Pointer to **string** |  Max length: 10; | [optional] 
**HomePhone** | Pointer to **string** |  Max length: 15; | [optional] 
**HomeExtension** | Pointer to **string** |  Max length: 10; | [optional] 
**DefaultPhone** | **NullableString** |  | 
**SecurityRole** | [**SecurityRoleReference**](SecurityRoleReference.md) |  | 
**Office365** | Pointer to [**MemberOffice365**](MemberOffice365.md) |  | [optional] 
**MapiName** | Pointer to **string** |  | [optional] 
**CalendarSyncIntegrationFlag** | Pointer to **NullableBool** |  | [optional] 
**AuthenticationServiceType** | Pointer to **NullableString** |  | [optional] 
**TimebasedOneTimePasswordActivated** | Pointer to **NullableBool** |  | [optional] 
**EnableLdapAuthenticationFlag** | Pointer to **NullableBool** |  | [optional] 
**LdapConfiguration** | Pointer to [**LdapConfigurationReference**](LdapConfigurationReference.md) |  | [optional] 
**LdapUserName** | Pointer to **string** |  Max length: 50; | [optional] 
**DirectionalSync** | Pointer to [**DirectionalSyncReference**](DirectionalSyncReference.md) |  | [optional] 
**SsoSettings** | Pointer to [**MemberSsoSettingsReference**](MemberSsoSettingsReference.md) |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**PhoneIntegrationType** | Pointer to **NullableString** |  | [optional] 
**UseBrowserLanguageFlag** | Pointer to **NullableBool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ReportCard** | Pointer to [**ReportCardReference**](ReportCardReference.md) |  | [optional] 
**EnableMobileFlag** | Pointer to **NullableBool** |  | [optional] 
**Type** | Pointer to [**MemberTypeReference**](MemberTypeReference.md) |  | [optional] 
**TimeZone** | Pointer to [**TimeZoneSetupReference**](TimeZoneSetupReference.md) |  | [optional] 
**PartnerPortalFlag** | Pointer to **NullableBool** |  | [optional] 
**StsUserAdminUrl** | Pointer to **string** |  | [optional] 
**ToastNotificationFlag** | Pointer to **NullableBool** |  | [optional] 
**MemberPersonas** | Pointer to **[]int32** |  | [optional] 
**AdminFlag** | Pointer to **NullableBool** |  | [optional] 
**StructureLevel** | Pointer to [**StructureReference**](StructureReference.md) |  | [optional] 
**SecurityLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**DefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**DefaultDepartment** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**ReportsTo** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**RestrictLocationFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictDepartmentFlag** | Pointer to **NullableBool** |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**TimeApprover** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**ExpenseApprover** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**BillableForecast** | Pointer to **NullableFloat64** |  | [optional] 
**DailyCapacity** | Pointer to **NullableFloat64** |  | [optional] 
**HourlyCost** | Pointer to **NullableFloat64** |  | [optional] 
**HourlyRate** | Pointer to **NullableFloat64** |  | [optional] 
**IncludeInUtilizationReportingFlag** | Pointer to **NullableBool** |  | [optional] 
**RequireExpenseEntryFlag** | Pointer to **NullableBool** |  | [optional] 
**RequireTimeSheetEntryFlag** | Pointer to **NullableBool** |  | [optional] 
**RequireStartAndEndTimeOnTimeEntryFlag** | Pointer to **NullableBool** |  | [optional] 
**AllowInCellEntryOnTimeSheet** | Pointer to **NullableBool** |  | [optional] 
**EnterTimeAgainstCompanyFlag** | Pointer to **NullableBool** |  | [optional] 
**AllowExpensesEnteredAgainstCompaniesFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeReminderEmailFlag** | Pointer to **NullableBool** |  | [optional] 
**DaysTolerance** | Pointer to **NullableInt32** |  | [optional] 
**MinimumHours** | Pointer to **NullableFloat64** |  | [optional] 
**TimeSheetStartDate** | Pointer to **string** |  | [optional] 
**ServiceDefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**ServiceDefaultDepartment** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**ServiceDefaultBoard** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**RestrictServiceDefaultLocationFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictServiceDefaultDepartmentFlag** | Pointer to **NullableBool** |  | [optional] 
**ExcludedServiceBoardIds** | Pointer to **[]int32** |  | [optional] 
**Teams** | Pointer to **[]int32** |  | [optional] 
**ServiceBoardTeamIds** | Pointer to **[]int32** |  | [optional] 
**ProjectDefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**ProjectDefaultDepartment** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**ProjectDefaultBoard** | Pointer to [**ProjectBoardReference**](ProjectBoardReference.md) |  | [optional] 
**RestrictProjectDefaultLocationFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictProjectDefaultDepartmentFlag** | Pointer to **NullableBool** |  | [optional] 
**ExcludedProjectBoardIds** | Pointer to **[]int32** |  | [optional] 
**ScheduleDefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**ScheduleDefaultDepartment** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**ScheduleCapacity** | Pointer to **NullableFloat64** |  | [optional] 
**ServiceLocation** | Pointer to [**ServiceLocationReference**](ServiceLocationReference.md) |  | [optional] 
**RestrictScheduleFlag** | Pointer to **NullableBool** |  | [optional] 
**HideMemberInDispatchPortalFlag** | Pointer to **NullableBool** |  | [optional] 
**Calendar** | Pointer to [**CalendarReference**](CalendarReference.md) |  | [optional] 
**SalesDefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**RestrictDefaultSalesTerritoryFlag** | Pointer to **NullableBool** |  | [optional] 
**Warehouse** | Pointer to [**WarehouseReference**](WarehouseReference.md) |  | [optional] 
**WarehouseBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**RestrictDefaultWarehouseFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictDefaultWarehouseBinFlag** | Pointer to **NullableBool** |  | [optional] 
**CompanyActivityTabFormat** | Pointer to **NullableString** |  | [optional] 
**InvoiceTimeTabFormat** | Pointer to **NullableString** |  | [optional] 
**InvoiceScreenDefaultTabFormat** | Pointer to **NullableString** |  | [optional] 
**InvoicingDisplayOptions** | Pointer to **NullableString** |  | [optional] 
**AgreementInvoicingDisplayOptions** | Pointer to **NullableString** |  | [optional] 
**AutoStartStopwatch** | Pointer to **NullableBool** |  | [optional] 
**AutoPopupQuickNotesWithStopwatch** | Pointer to **NullableBool** |  | [optional] 
**GlobalSearchDefaultTicketFilter** | Pointer to **NullableString** |  | [optional] 
**GlobalSearchDefaultSort** | Pointer to **NullableString** |  | [optional] 
**PhoneSource** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CopyPodLayouts** | Pointer to **bool** |  | [optional] 
**CopySharedDefaultViews** | Pointer to **bool** |  | [optional] 
**CopyColumnLayoutsAndFilters** | Pointer to **bool** |  | [optional] 
**FromMemberRecId** | Pointer to **int32** |  | [optional] 
**FromMemberTemplateRecId** | Pointer to **int32** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewMember

`func NewMember(identifier string, licenseClass NullableString, firstName string, lastName string, hireDate time.Time, defaultEmail NullableString, defaultPhone NullableString, securityRole SecurityRoleReference, ) *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Member) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Member) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Member) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Member) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *Member) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Member) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Member) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetPassword

`func (o *Member) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Member) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Member) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Member) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDisableOnlineFlag

`func (o *Member) GetDisableOnlineFlag() bool`

GetDisableOnlineFlag returns the DisableOnlineFlag field if non-nil, zero value otherwise.

### GetDisableOnlineFlagOk

`func (o *Member) GetDisableOnlineFlagOk() (*bool, bool)`

GetDisableOnlineFlagOk returns a tuple with the DisableOnlineFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOnlineFlag

`func (o *Member) SetDisableOnlineFlag(v bool)`

SetDisableOnlineFlag sets DisableOnlineFlag field to given value.

### HasDisableOnlineFlag

`func (o *Member) HasDisableOnlineFlag() bool`

HasDisableOnlineFlag returns a boolean if a field has been set.

### SetDisableOnlineFlagNil

`func (o *Member) SetDisableOnlineFlagNil(b bool)`

 SetDisableOnlineFlagNil sets the value for DisableOnlineFlag to be an explicit nil

### UnsetDisableOnlineFlag
`func (o *Member) UnsetDisableOnlineFlag()`

UnsetDisableOnlineFlag ensures that no value is present for DisableOnlineFlag, not even an explicit nil
### GetLicenseClass

`func (o *Member) GetLicenseClass() string`

GetLicenseClass returns the LicenseClass field if non-nil, zero value otherwise.

### GetLicenseClassOk

`func (o *Member) GetLicenseClassOk() (*string, bool)`

GetLicenseClassOk returns a tuple with the LicenseClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseClass

`func (o *Member) SetLicenseClass(v string)`

SetLicenseClass sets LicenseClass field to given value.


### SetLicenseClassNil

`func (o *Member) SetLicenseClassNil(b bool)`

 SetLicenseClassNil sets the value for LicenseClass to be an explicit nil

### UnsetLicenseClass
`func (o *Member) UnsetLicenseClass()`

UnsetLicenseClass ensures that no value is present for LicenseClass, not even an explicit nil
### GetNotes

`func (o *Member) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Member) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Member) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Member) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetEmployeeIdentifer

`func (o *Member) GetEmployeeIdentifer() string`

GetEmployeeIdentifer returns the EmployeeIdentifer field if non-nil, zero value otherwise.

### GetEmployeeIdentiferOk

`func (o *Member) GetEmployeeIdentiferOk() (*string, bool)`

GetEmployeeIdentiferOk returns a tuple with the EmployeeIdentifer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeIdentifer

`func (o *Member) SetEmployeeIdentifer(v string)`

SetEmployeeIdentifer sets EmployeeIdentifer field to given value.

### HasEmployeeIdentifer

`func (o *Member) HasEmployeeIdentifer() bool`

HasEmployeeIdentifer returns a boolean if a field has been set.

### GetVendorNumber

`func (o *Member) GetVendorNumber() string`

GetVendorNumber returns the VendorNumber field if non-nil, zero value otherwise.

### GetVendorNumberOk

`func (o *Member) GetVendorNumberOk() (*string, bool)`

GetVendorNumberOk returns a tuple with the VendorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumber

`func (o *Member) SetVendorNumber(v string)`

SetVendorNumber sets VendorNumber field to given value.

### HasVendorNumber

`func (o *Member) HasVendorNumber() bool`

HasVendorNumber returns a boolean if a field has been set.

### GetEnableMobileGpsFlag

`func (o *Member) GetEnableMobileGpsFlag() bool`

GetEnableMobileGpsFlag returns the EnableMobileGpsFlag field if non-nil, zero value otherwise.

### GetEnableMobileGpsFlagOk

`func (o *Member) GetEnableMobileGpsFlagOk() (*bool, bool)`

GetEnableMobileGpsFlagOk returns a tuple with the EnableMobileGpsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMobileGpsFlag

`func (o *Member) SetEnableMobileGpsFlag(v bool)`

SetEnableMobileGpsFlag sets EnableMobileGpsFlag field to given value.

### HasEnableMobileGpsFlag

`func (o *Member) HasEnableMobileGpsFlag() bool`

HasEnableMobileGpsFlag returns a boolean if a field has been set.

### SetEnableMobileGpsFlagNil

`func (o *Member) SetEnableMobileGpsFlagNil(b bool)`

 SetEnableMobileGpsFlagNil sets the value for EnableMobileGpsFlag to be an explicit nil

### UnsetEnableMobileGpsFlag
`func (o *Member) UnsetEnableMobileGpsFlag()`

UnsetEnableMobileGpsFlag ensures that no value is present for EnableMobileGpsFlag, not even an explicit nil
### GetInactiveDate

`func (o *Member) GetInactiveDate() time.Time`

GetInactiveDate returns the InactiveDate field if non-nil, zero value otherwise.

### GetInactiveDateOk

`func (o *Member) GetInactiveDateOk() (*time.Time, bool)`

GetInactiveDateOk returns a tuple with the InactiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveDate

`func (o *Member) SetInactiveDate(v time.Time)`

SetInactiveDate sets InactiveDate field to given value.

### HasInactiveDate

`func (o *Member) HasInactiveDate() bool`

HasInactiveDate returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *Member) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *Member) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *Member) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *Member) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *Member) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *Member) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetLastLogin

`func (o *Member) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *Member) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *Member) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *Member) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetClientId

`func (o *Member) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Member) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Member) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Member) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetToken

`func (o *Member) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Member) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Member) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Member) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetFirstName

`func (o *Member) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Member) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Member) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleInitial

`func (o *Member) GetMiddleInitial() string`

GetMiddleInitial returns the MiddleInitial field if non-nil, zero value otherwise.

### GetMiddleInitialOk

`func (o *Member) GetMiddleInitialOk() (*string, bool)`

GetMiddleInitialOk returns a tuple with the MiddleInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleInitial

`func (o *Member) SetMiddleInitial(v string)`

SetMiddleInitial sets MiddleInitial field to given value.

### HasMiddleInitial

`func (o *Member) HasMiddleInitial() bool`

HasMiddleInitial returns a boolean if a field has been set.

### GetLastName

`func (o *Member) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Member) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Member) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetHireDate

`func (o *Member) GetHireDate() time.Time`

GetHireDate returns the HireDate field if non-nil, zero value otherwise.

### GetHireDateOk

`func (o *Member) GetHireDateOk() (*time.Time, bool)`

GetHireDateOk returns a tuple with the HireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireDate

`func (o *Member) SetHireDate(v time.Time)`

SetHireDate sets HireDate field to given value.


### GetCountry

`func (o *Member) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Member) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Member) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Member) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPhoto

`func (o *Member) GetPhoto() DocumentReference`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *Member) GetPhotoOk() (*DocumentReference, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *Member) SetPhoto(v DocumentReference)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *Member) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetOfficeEmail

`func (o *Member) GetOfficeEmail() string`

GetOfficeEmail returns the OfficeEmail field if non-nil, zero value otherwise.

### GetOfficeEmailOk

`func (o *Member) GetOfficeEmailOk() (*string, bool)`

GetOfficeEmailOk returns a tuple with the OfficeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeEmail

`func (o *Member) SetOfficeEmail(v string)`

SetOfficeEmail sets OfficeEmail field to given value.

### HasOfficeEmail

`func (o *Member) HasOfficeEmail() bool`

HasOfficeEmail returns a boolean if a field has been set.

### GetMobileEmail

`func (o *Member) GetMobileEmail() string`

GetMobileEmail returns the MobileEmail field if non-nil, zero value otherwise.

### GetMobileEmailOk

`func (o *Member) GetMobileEmailOk() (*string, bool)`

GetMobileEmailOk returns a tuple with the MobileEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileEmail

`func (o *Member) SetMobileEmail(v string)`

SetMobileEmail sets MobileEmail field to given value.

### HasMobileEmail

`func (o *Member) HasMobileEmail() bool`

HasMobileEmail returns a boolean if a field has been set.

### GetHomeEmail

`func (o *Member) GetHomeEmail() string`

GetHomeEmail returns the HomeEmail field if non-nil, zero value otherwise.

### GetHomeEmailOk

`func (o *Member) GetHomeEmailOk() (*string, bool)`

GetHomeEmailOk returns a tuple with the HomeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeEmail

`func (o *Member) SetHomeEmail(v string)`

SetHomeEmail sets HomeEmail field to given value.

### HasHomeEmail

`func (o *Member) HasHomeEmail() bool`

HasHomeEmail returns a boolean if a field has been set.

### GetDefaultEmail

`func (o *Member) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *Member) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *Member) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.


### SetDefaultEmailNil

`func (o *Member) SetDefaultEmailNil(b bool)`

 SetDefaultEmailNil sets the value for DefaultEmail to be an explicit nil

### UnsetDefaultEmail
`func (o *Member) UnsetDefaultEmail()`

UnsetDefaultEmail ensures that no value is present for DefaultEmail, not even an explicit nil
### GetPrimaryEmail

`func (o *Member) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *Member) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *Member) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *Member) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### GetOfficePhone

`func (o *Member) GetOfficePhone() string`

GetOfficePhone returns the OfficePhone field if non-nil, zero value otherwise.

### GetOfficePhoneOk

`func (o *Member) GetOfficePhoneOk() (*string, bool)`

GetOfficePhoneOk returns a tuple with the OfficePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficePhone

`func (o *Member) SetOfficePhone(v string)`

SetOfficePhone sets OfficePhone field to given value.

### HasOfficePhone

`func (o *Member) HasOfficePhone() bool`

HasOfficePhone returns a boolean if a field has been set.

### GetOfficeExtension

`func (o *Member) GetOfficeExtension() string`

GetOfficeExtension returns the OfficeExtension field if non-nil, zero value otherwise.

### GetOfficeExtensionOk

`func (o *Member) GetOfficeExtensionOk() (*string, bool)`

GetOfficeExtensionOk returns a tuple with the OfficeExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeExtension

`func (o *Member) SetOfficeExtension(v string)`

SetOfficeExtension sets OfficeExtension field to given value.

### HasOfficeExtension

`func (o *Member) HasOfficeExtension() bool`

HasOfficeExtension returns a boolean if a field has been set.

### GetMobilePhone

`func (o *Member) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *Member) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *Member) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *Member) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetMobileExtension

`func (o *Member) GetMobileExtension() string`

GetMobileExtension returns the MobileExtension field if non-nil, zero value otherwise.

### GetMobileExtensionOk

`func (o *Member) GetMobileExtensionOk() (*string, bool)`

GetMobileExtensionOk returns a tuple with the MobileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileExtension

`func (o *Member) SetMobileExtension(v string)`

SetMobileExtension sets MobileExtension field to given value.

### HasMobileExtension

`func (o *Member) HasMobileExtension() bool`

HasMobileExtension returns a boolean if a field has been set.

### GetHomePhone

`func (o *Member) GetHomePhone() string`

GetHomePhone returns the HomePhone field if non-nil, zero value otherwise.

### GetHomePhoneOk

`func (o *Member) GetHomePhoneOk() (*string, bool)`

GetHomePhoneOk returns a tuple with the HomePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePhone

`func (o *Member) SetHomePhone(v string)`

SetHomePhone sets HomePhone field to given value.

### HasHomePhone

`func (o *Member) HasHomePhone() bool`

HasHomePhone returns a boolean if a field has been set.

### GetHomeExtension

`func (o *Member) GetHomeExtension() string`

GetHomeExtension returns the HomeExtension field if non-nil, zero value otherwise.

### GetHomeExtensionOk

`func (o *Member) GetHomeExtensionOk() (*string, bool)`

GetHomeExtensionOk returns a tuple with the HomeExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeExtension

`func (o *Member) SetHomeExtension(v string)`

SetHomeExtension sets HomeExtension field to given value.

### HasHomeExtension

`func (o *Member) HasHomeExtension() bool`

HasHomeExtension returns a boolean if a field has been set.

### GetDefaultPhone

`func (o *Member) GetDefaultPhone() string`

GetDefaultPhone returns the DefaultPhone field if non-nil, zero value otherwise.

### GetDefaultPhoneOk

`func (o *Member) GetDefaultPhoneOk() (*string, bool)`

GetDefaultPhoneOk returns a tuple with the DefaultPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhone

`func (o *Member) SetDefaultPhone(v string)`

SetDefaultPhone sets DefaultPhone field to given value.


### SetDefaultPhoneNil

`func (o *Member) SetDefaultPhoneNil(b bool)`

 SetDefaultPhoneNil sets the value for DefaultPhone to be an explicit nil

### UnsetDefaultPhone
`func (o *Member) UnsetDefaultPhone()`

UnsetDefaultPhone ensures that no value is present for DefaultPhone, not even an explicit nil
### GetSecurityRole

`func (o *Member) GetSecurityRole() SecurityRoleReference`

GetSecurityRole returns the SecurityRole field if non-nil, zero value otherwise.

### GetSecurityRoleOk

`func (o *Member) GetSecurityRoleOk() (*SecurityRoleReference, bool)`

GetSecurityRoleOk returns a tuple with the SecurityRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityRole

`func (o *Member) SetSecurityRole(v SecurityRoleReference)`

SetSecurityRole sets SecurityRole field to given value.


### GetOffice365

`func (o *Member) GetOffice365() MemberOffice365`

GetOffice365 returns the Office365 field if non-nil, zero value otherwise.

### GetOffice365Ok

`func (o *Member) GetOffice365Ok() (*MemberOffice365, bool)`

GetOffice365Ok returns a tuple with the Office365 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365

`func (o *Member) SetOffice365(v MemberOffice365)`

SetOffice365 sets Office365 field to given value.

### HasOffice365

`func (o *Member) HasOffice365() bool`

HasOffice365 returns a boolean if a field has been set.

### GetMapiName

`func (o *Member) GetMapiName() string`

GetMapiName returns the MapiName field if non-nil, zero value otherwise.

### GetMapiNameOk

`func (o *Member) GetMapiNameOk() (*string, bool)`

GetMapiNameOk returns a tuple with the MapiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapiName

`func (o *Member) SetMapiName(v string)`

SetMapiName sets MapiName field to given value.

### HasMapiName

`func (o *Member) HasMapiName() bool`

HasMapiName returns a boolean if a field has been set.

### GetCalendarSyncIntegrationFlag

`func (o *Member) GetCalendarSyncIntegrationFlag() bool`

GetCalendarSyncIntegrationFlag returns the CalendarSyncIntegrationFlag field if non-nil, zero value otherwise.

### GetCalendarSyncIntegrationFlagOk

`func (o *Member) GetCalendarSyncIntegrationFlagOk() (*bool, bool)`

GetCalendarSyncIntegrationFlagOk returns a tuple with the CalendarSyncIntegrationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarSyncIntegrationFlag

`func (o *Member) SetCalendarSyncIntegrationFlag(v bool)`

SetCalendarSyncIntegrationFlag sets CalendarSyncIntegrationFlag field to given value.

### HasCalendarSyncIntegrationFlag

`func (o *Member) HasCalendarSyncIntegrationFlag() bool`

HasCalendarSyncIntegrationFlag returns a boolean if a field has been set.

### SetCalendarSyncIntegrationFlagNil

`func (o *Member) SetCalendarSyncIntegrationFlagNil(b bool)`

 SetCalendarSyncIntegrationFlagNil sets the value for CalendarSyncIntegrationFlag to be an explicit nil

### UnsetCalendarSyncIntegrationFlag
`func (o *Member) UnsetCalendarSyncIntegrationFlag()`

UnsetCalendarSyncIntegrationFlag ensures that no value is present for CalendarSyncIntegrationFlag, not even an explicit nil
### GetAuthenticationServiceType

`func (o *Member) GetAuthenticationServiceType() string`

GetAuthenticationServiceType returns the AuthenticationServiceType field if non-nil, zero value otherwise.

### GetAuthenticationServiceTypeOk

`func (o *Member) GetAuthenticationServiceTypeOk() (*string, bool)`

GetAuthenticationServiceTypeOk returns a tuple with the AuthenticationServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationServiceType

`func (o *Member) SetAuthenticationServiceType(v string)`

SetAuthenticationServiceType sets AuthenticationServiceType field to given value.

### HasAuthenticationServiceType

`func (o *Member) HasAuthenticationServiceType() bool`

HasAuthenticationServiceType returns a boolean if a field has been set.

### SetAuthenticationServiceTypeNil

`func (o *Member) SetAuthenticationServiceTypeNil(b bool)`

 SetAuthenticationServiceTypeNil sets the value for AuthenticationServiceType to be an explicit nil

### UnsetAuthenticationServiceType
`func (o *Member) UnsetAuthenticationServiceType()`

UnsetAuthenticationServiceType ensures that no value is present for AuthenticationServiceType, not even an explicit nil
### GetTimebasedOneTimePasswordActivated

`func (o *Member) GetTimebasedOneTimePasswordActivated() bool`

GetTimebasedOneTimePasswordActivated returns the TimebasedOneTimePasswordActivated field if non-nil, zero value otherwise.

### GetTimebasedOneTimePasswordActivatedOk

`func (o *Member) GetTimebasedOneTimePasswordActivatedOk() (*bool, bool)`

GetTimebasedOneTimePasswordActivatedOk returns a tuple with the TimebasedOneTimePasswordActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimebasedOneTimePasswordActivated

`func (o *Member) SetTimebasedOneTimePasswordActivated(v bool)`

SetTimebasedOneTimePasswordActivated sets TimebasedOneTimePasswordActivated field to given value.

### HasTimebasedOneTimePasswordActivated

`func (o *Member) HasTimebasedOneTimePasswordActivated() bool`

HasTimebasedOneTimePasswordActivated returns a boolean if a field has been set.

### SetTimebasedOneTimePasswordActivatedNil

`func (o *Member) SetTimebasedOneTimePasswordActivatedNil(b bool)`

 SetTimebasedOneTimePasswordActivatedNil sets the value for TimebasedOneTimePasswordActivated to be an explicit nil

### UnsetTimebasedOneTimePasswordActivated
`func (o *Member) UnsetTimebasedOneTimePasswordActivated()`

UnsetTimebasedOneTimePasswordActivated ensures that no value is present for TimebasedOneTimePasswordActivated, not even an explicit nil
### GetEnableLdapAuthenticationFlag

`func (o *Member) GetEnableLdapAuthenticationFlag() bool`

GetEnableLdapAuthenticationFlag returns the EnableLdapAuthenticationFlag field if non-nil, zero value otherwise.

### GetEnableLdapAuthenticationFlagOk

`func (o *Member) GetEnableLdapAuthenticationFlagOk() (*bool, bool)`

GetEnableLdapAuthenticationFlagOk returns a tuple with the EnableLdapAuthenticationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLdapAuthenticationFlag

`func (o *Member) SetEnableLdapAuthenticationFlag(v bool)`

SetEnableLdapAuthenticationFlag sets EnableLdapAuthenticationFlag field to given value.

### HasEnableLdapAuthenticationFlag

`func (o *Member) HasEnableLdapAuthenticationFlag() bool`

HasEnableLdapAuthenticationFlag returns a boolean if a field has been set.

### SetEnableLdapAuthenticationFlagNil

`func (o *Member) SetEnableLdapAuthenticationFlagNil(b bool)`

 SetEnableLdapAuthenticationFlagNil sets the value for EnableLdapAuthenticationFlag to be an explicit nil

### UnsetEnableLdapAuthenticationFlag
`func (o *Member) UnsetEnableLdapAuthenticationFlag()`

UnsetEnableLdapAuthenticationFlag ensures that no value is present for EnableLdapAuthenticationFlag, not even an explicit nil
### GetLdapConfiguration

`func (o *Member) GetLdapConfiguration() LdapConfigurationReference`

GetLdapConfiguration returns the LdapConfiguration field if non-nil, zero value otherwise.

### GetLdapConfigurationOk

`func (o *Member) GetLdapConfigurationOk() (*LdapConfigurationReference, bool)`

GetLdapConfigurationOk returns a tuple with the LdapConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapConfiguration

`func (o *Member) SetLdapConfiguration(v LdapConfigurationReference)`

SetLdapConfiguration sets LdapConfiguration field to given value.

### HasLdapConfiguration

`func (o *Member) HasLdapConfiguration() bool`

HasLdapConfiguration returns a boolean if a field has been set.

### GetLdapUserName

`func (o *Member) GetLdapUserName() string`

GetLdapUserName returns the LdapUserName field if non-nil, zero value otherwise.

### GetLdapUserNameOk

`func (o *Member) GetLdapUserNameOk() (*string, bool)`

GetLdapUserNameOk returns a tuple with the LdapUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserName

`func (o *Member) SetLdapUserName(v string)`

SetLdapUserName sets LdapUserName field to given value.

### HasLdapUserName

`func (o *Member) HasLdapUserName() bool`

HasLdapUserName returns a boolean if a field has been set.

### GetDirectionalSync

`func (o *Member) GetDirectionalSync() DirectionalSyncReference`

GetDirectionalSync returns the DirectionalSync field if non-nil, zero value otherwise.

### GetDirectionalSyncOk

`func (o *Member) GetDirectionalSyncOk() (*DirectionalSyncReference, bool)`

GetDirectionalSyncOk returns a tuple with the DirectionalSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionalSync

`func (o *Member) SetDirectionalSync(v DirectionalSyncReference)`

SetDirectionalSync sets DirectionalSync field to given value.

### HasDirectionalSync

`func (o *Member) HasDirectionalSync() bool`

HasDirectionalSync returns a boolean if a field has been set.

### GetSsoSettings

`func (o *Member) GetSsoSettings() MemberSsoSettingsReference`

GetSsoSettings returns the SsoSettings field if non-nil, zero value otherwise.

### GetSsoSettingsOk

`func (o *Member) GetSsoSettingsOk() (*MemberSsoSettingsReference, bool)`

GetSsoSettingsOk returns a tuple with the SsoSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSettings

`func (o *Member) SetSsoSettings(v MemberSsoSettingsReference)`

SetSsoSettings sets SsoSettings field to given value.

### HasSsoSettings

`func (o *Member) HasSsoSettings() bool`

HasSsoSettings returns a boolean if a field has been set.

### GetSignature

`func (o *Member) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Member) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Member) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *Member) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetPhoneIntegrationType

`func (o *Member) GetPhoneIntegrationType() string`

GetPhoneIntegrationType returns the PhoneIntegrationType field if non-nil, zero value otherwise.

### GetPhoneIntegrationTypeOk

`func (o *Member) GetPhoneIntegrationTypeOk() (*string, bool)`

GetPhoneIntegrationTypeOk returns a tuple with the PhoneIntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneIntegrationType

`func (o *Member) SetPhoneIntegrationType(v string)`

SetPhoneIntegrationType sets PhoneIntegrationType field to given value.

### HasPhoneIntegrationType

`func (o *Member) HasPhoneIntegrationType() bool`

HasPhoneIntegrationType returns a boolean if a field has been set.

### SetPhoneIntegrationTypeNil

`func (o *Member) SetPhoneIntegrationTypeNil(b bool)`

 SetPhoneIntegrationTypeNil sets the value for PhoneIntegrationType to be an explicit nil

### UnsetPhoneIntegrationType
`func (o *Member) UnsetPhoneIntegrationType()`

UnsetPhoneIntegrationType ensures that no value is present for PhoneIntegrationType, not even an explicit nil
### GetUseBrowserLanguageFlag

`func (o *Member) GetUseBrowserLanguageFlag() bool`

GetUseBrowserLanguageFlag returns the UseBrowserLanguageFlag field if non-nil, zero value otherwise.

### GetUseBrowserLanguageFlagOk

`func (o *Member) GetUseBrowserLanguageFlagOk() (*bool, bool)`

GetUseBrowserLanguageFlagOk returns a tuple with the UseBrowserLanguageFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBrowserLanguageFlag

`func (o *Member) SetUseBrowserLanguageFlag(v bool)`

SetUseBrowserLanguageFlag sets UseBrowserLanguageFlag field to given value.

### HasUseBrowserLanguageFlag

`func (o *Member) HasUseBrowserLanguageFlag() bool`

HasUseBrowserLanguageFlag returns a boolean if a field has been set.

### SetUseBrowserLanguageFlagNil

`func (o *Member) SetUseBrowserLanguageFlagNil(b bool)`

 SetUseBrowserLanguageFlagNil sets the value for UseBrowserLanguageFlag to be an explicit nil

### UnsetUseBrowserLanguageFlag
`func (o *Member) UnsetUseBrowserLanguageFlag()`

UnsetUseBrowserLanguageFlag ensures that no value is present for UseBrowserLanguageFlag, not even an explicit nil
### GetTitle

`func (o *Member) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Member) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Member) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Member) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetReportCard

`func (o *Member) GetReportCard() ReportCardReference`

GetReportCard returns the ReportCard field if non-nil, zero value otherwise.

### GetReportCardOk

`func (o *Member) GetReportCardOk() (*ReportCardReference, bool)`

GetReportCardOk returns a tuple with the ReportCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCard

`func (o *Member) SetReportCard(v ReportCardReference)`

SetReportCard sets ReportCard field to given value.

### HasReportCard

`func (o *Member) HasReportCard() bool`

HasReportCard returns a boolean if a field has been set.

### GetEnableMobileFlag

`func (o *Member) GetEnableMobileFlag() bool`

GetEnableMobileFlag returns the EnableMobileFlag field if non-nil, zero value otherwise.

### GetEnableMobileFlagOk

`func (o *Member) GetEnableMobileFlagOk() (*bool, bool)`

GetEnableMobileFlagOk returns a tuple with the EnableMobileFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMobileFlag

`func (o *Member) SetEnableMobileFlag(v bool)`

SetEnableMobileFlag sets EnableMobileFlag field to given value.

### HasEnableMobileFlag

`func (o *Member) HasEnableMobileFlag() bool`

HasEnableMobileFlag returns a boolean if a field has been set.

### SetEnableMobileFlagNil

`func (o *Member) SetEnableMobileFlagNil(b bool)`

 SetEnableMobileFlagNil sets the value for EnableMobileFlag to be an explicit nil

### UnsetEnableMobileFlag
`func (o *Member) UnsetEnableMobileFlag()`

UnsetEnableMobileFlag ensures that no value is present for EnableMobileFlag, not even an explicit nil
### GetType

`func (o *Member) GetType() MemberTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Member) GetTypeOk() (*MemberTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Member) SetType(v MemberTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *Member) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeZone

`func (o *Member) GetTimeZone() TimeZoneSetupReference`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Member) GetTimeZoneOk() (*TimeZoneSetupReference, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Member) SetTimeZone(v TimeZoneSetupReference)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Member) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetPartnerPortalFlag

`func (o *Member) GetPartnerPortalFlag() bool`

GetPartnerPortalFlag returns the PartnerPortalFlag field if non-nil, zero value otherwise.

### GetPartnerPortalFlagOk

`func (o *Member) GetPartnerPortalFlagOk() (*bool, bool)`

GetPartnerPortalFlagOk returns a tuple with the PartnerPortalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPortalFlag

`func (o *Member) SetPartnerPortalFlag(v bool)`

SetPartnerPortalFlag sets PartnerPortalFlag field to given value.

### HasPartnerPortalFlag

`func (o *Member) HasPartnerPortalFlag() bool`

HasPartnerPortalFlag returns a boolean if a field has been set.

### SetPartnerPortalFlagNil

`func (o *Member) SetPartnerPortalFlagNil(b bool)`

 SetPartnerPortalFlagNil sets the value for PartnerPortalFlag to be an explicit nil

### UnsetPartnerPortalFlag
`func (o *Member) UnsetPartnerPortalFlag()`

UnsetPartnerPortalFlag ensures that no value is present for PartnerPortalFlag, not even an explicit nil
### GetStsUserAdminUrl

`func (o *Member) GetStsUserAdminUrl() string`

GetStsUserAdminUrl returns the StsUserAdminUrl field if non-nil, zero value otherwise.

### GetStsUserAdminUrlOk

`func (o *Member) GetStsUserAdminUrlOk() (*string, bool)`

GetStsUserAdminUrlOk returns a tuple with the StsUserAdminUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsUserAdminUrl

`func (o *Member) SetStsUserAdminUrl(v string)`

SetStsUserAdminUrl sets StsUserAdminUrl field to given value.

### HasStsUserAdminUrl

`func (o *Member) HasStsUserAdminUrl() bool`

HasStsUserAdminUrl returns a boolean if a field has been set.

### GetToastNotificationFlag

`func (o *Member) GetToastNotificationFlag() bool`

GetToastNotificationFlag returns the ToastNotificationFlag field if non-nil, zero value otherwise.

### GetToastNotificationFlagOk

`func (o *Member) GetToastNotificationFlagOk() (*bool, bool)`

GetToastNotificationFlagOk returns a tuple with the ToastNotificationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToastNotificationFlag

`func (o *Member) SetToastNotificationFlag(v bool)`

SetToastNotificationFlag sets ToastNotificationFlag field to given value.

### HasToastNotificationFlag

`func (o *Member) HasToastNotificationFlag() bool`

HasToastNotificationFlag returns a boolean if a field has been set.

### SetToastNotificationFlagNil

`func (o *Member) SetToastNotificationFlagNil(b bool)`

 SetToastNotificationFlagNil sets the value for ToastNotificationFlag to be an explicit nil

### UnsetToastNotificationFlag
`func (o *Member) UnsetToastNotificationFlag()`

UnsetToastNotificationFlag ensures that no value is present for ToastNotificationFlag, not even an explicit nil
### GetMemberPersonas

`func (o *Member) GetMemberPersonas() []int32`

GetMemberPersonas returns the MemberPersonas field if non-nil, zero value otherwise.

### GetMemberPersonasOk

`func (o *Member) GetMemberPersonasOk() (*[]int32, bool)`

GetMemberPersonasOk returns a tuple with the MemberPersonas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPersonas

`func (o *Member) SetMemberPersonas(v []int32)`

SetMemberPersonas sets MemberPersonas field to given value.

### HasMemberPersonas

`func (o *Member) HasMemberPersonas() bool`

HasMemberPersonas returns a boolean if a field has been set.

### GetAdminFlag

`func (o *Member) GetAdminFlag() bool`

GetAdminFlag returns the AdminFlag field if non-nil, zero value otherwise.

### GetAdminFlagOk

`func (o *Member) GetAdminFlagOk() (*bool, bool)`

GetAdminFlagOk returns a tuple with the AdminFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFlag

`func (o *Member) SetAdminFlag(v bool)`

SetAdminFlag sets AdminFlag field to given value.

### HasAdminFlag

`func (o *Member) HasAdminFlag() bool`

HasAdminFlag returns a boolean if a field has been set.

### SetAdminFlagNil

`func (o *Member) SetAdminFlagNil(b bool)`

 SetAdminFlagNil sets the value for AdminFlag to be an explicit nil

### UnsetAdminFlag
`func (o *Member) UnsetAdminFlag()`

UnsetAdminFlag ensures that no value is present for AdminFlag, not even an explicit nil
### GetStructureLevel

`func (o *Member) GetStructureLevel() StructureReference`

GetStructureLevel returns the StructureLevel field if non-nil, zero value otherwise.

### GetStructureLevelOk

`func (o *Member) GetStructureLevelOk() (*StructureReference, bool)`

GetStructureLevelOk returns a tuple with the StructureLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureLevel

`func (o *Member) SetStructureLevel(v StructureReference)`

SetStructureLevel sets StructureLevel field to given value.

### HasStructureLevel

`func (o *Member) HasStructureLevel() bool`

HasStructureLevel returns a boolean if a field has been set.

### GetSecurityLocation

`func (o *Member) GetSecurityLocation() SystemLocationReference`

GetSecurityLocation returns the SecurityLocation field if non-nil, zero value otherwise.

### GetSecurityLocationOk

`func (o *Member) GetSecurityLocationOk() (*SystemLocationReference, bool)`

GetSecurityLocationOk returns a tuple with the SecurityLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLocation

`func (o *Member) SetSecurityLocation(v SystemLocationReference)`

SetSecurityLocation sets SecurityLocation field to given value.

### HasSecurityLocation

`func (o *Member) HasSecurityLocation() bool`

HasSecurityLocation returns a boolean if a field has been set.

### GetDefaultLocation

`func (o *Member) GetDefaultLocation() SystemLocationReference`

GetDefaultLocation returns the DefaultLocation field if non-nil, zero value otherwise.

### GetDefaultLocationOk

`func (o *Member) GetDefaultLocationOk() (*SystemLocationReference, bool)`

GetDefaultLocationOk returns a tuple with the DefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocation

`func (o *Member) SetDefaultLocation(v SystemLocationReference)`

SetDefaultLocation sets DefaultLocation field to given value.

### HasDefaultLocation

`func (o *Member) HasDefaultLocation() bool`

HasDefaultLocation returns a boolean if a field has been set.

### GetDefaultDepartment

`func (o *Member) GetDefaultDepartment() SystemDepartmentReference`

GetDefaultDepartment returns the DefaultDepartment field if non-nil, zero value otherwise.

### GetDefaultDepartmentOk

`func (o *Member) GetDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetDefaultDepartmentOk returns a tuple with the DefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDepartment

`func (o *Member) SetDefaultDepartment(v SystemDepartmentReference)`

SetDefaultDepartment sets DefaultDepartment field to given value.

### HasDefaultDepartment

`func (o *Member) HasDefaultDepartment() bool`

HasDefaultDepartment returns a boolean if a field has been set.

### GetReportsTo

`func (o *Member) GetReportsTo() MemberReference`

GetReportsTo returns the ReportsTo field if non-nil, zero value otherwise.

### GetReportsToOk

`func (o *Member) GetReportsToOk() (*MemberReference, bool)`

GetReportsToOk returns a tuple with the ReportsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsTo

`func (o *Member) SetReportsTo(v MemberReference)`

SetReportsTo sets ReportsTo field to given value.

### HasReportsTo

`func (o *Member) HasReportsTo() bool`

HasReportsTo returns a boolean if a field has been set.

### GetRestrictLocationFlag

`func (o *Member) GetRestrictLocationFlag() bool`

GetRestrictLocationFlag returns the RestrictLocationFlag field if non-nil, zero value otherwise.

### GetRestrictLocationFlagOk

`func (o *Member) GetRestrictLocationFlagOk() (*bool, bool)`

GetRestrictLocationFlagOk returns a tuple with the RestrictLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictLocationFlag

`func (o *Member) SetRestrictLocationFlag(v bool)`

SetRestrictLocationFlag sets RestrictLocationFlag field to given value.

### HasRestrictLocationFlag

`func (o *Member) HasRestrictLocationFlag() bool`

HasRestrictLocationFlag returns a boolean if a field has been set.

### SetRestrictLocationFlagNil

`func (o *Member) SetRestrictLocationFlagNil(b bool)`

 SetRestrictLocationFlagNil sets the value for RestrictLocationFlag to be an explicit nil

### UnsetRestrictLocationFlag
`func (o *Member) UnsetRestrictLocationFlag()`

UnsetRestrictLocationFlag ensures that no value is present for RestrictLocationFlag, not even an explicit nil
### GetRestrictDepartmentFlag

`func (o *Member) GetRestrictDepartmentFlag() bool`

GetRestrictDepartmentFlag returns the RestrictDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictDepartmentFlagOk

`func (o *Member) GetRestrictDepartmentFlagOk() (*bool, bool)`

GetRestrictDepartmentFlagOk returns a tuple with the RestrictDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDepartmentFlag

`func (o *Member) SetRestrictDepartmentFlag(v bool)`

SetRestrictDepartmentFlag sets RestrictDepartmentFlag field to given value.

### HasRestrictDepartmentFlag

`func (o *Member) HasRestrictDepartmentFlag() bool`

HasRestrictDepartmentFlag returns a boolean if a field has been set.

### SetRestrictDepartmentFlagNil

`func (o *Member) SetRestrictDepartmentFlagNil(b bool)`

 SetRestrictDepartmentFlagNil sets the value for RestrictDepartmentFlag to be an explicit nil

### UnsetRestrictDepartmentFlag
`func (o *Member) UnsetRestrictDepartmentFlag()`

UnsetRestrictDepartmentFlag ensures that no value is present for RestrictDepartmentFlag, not even an explicit nil
### GetWorkRole

`func (o *Member) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *Member) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *Member) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *Member) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *Member) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *Member) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *Member) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *Member) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetTimeApprover

`func (o *Member) GetTimeApprover() MemberReference`

GetTimeApprover returns the TimeApprover field if non-nil, zero value otherwise.

### GetTimeApproverOk

`func (o *Member) GetTimeApproverOk() (*MemberReference, bool)`

GetTimeApproverOk returns a tuple with the TimeApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeApprover

`func (o *Member) SetTimeApprover(v MemberReference)`

SetTimeApprover sets TimeApprover field to given value.

### HasTimeApprover

`func (o *Member) HasTimeApprover() bool`

HasTimeApprover returns a boolean if a field has been set.

### GetExpenseApprover

`func (o *Member) GetExpenseApprover() MemberReference`

GetExpenseApprover returns the ExpenseApprover field if non-nil, zero value otherwise.

### GetExpenseApproverOk

`func (o *Member) GetExpenseApproverOk() (*MemberReference, bool)`

GetExpenseApproverOk returns a tuple with the ExpenseApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApprover

`func (o *Member) SetExpenseApprover(v MemberReference)`

SetExpenseApprover sets ExpenseApprover field to given value.

### HasExpenseApprover

`func (o *Member) HasExpenseApprover() bool`

HasExpenseApprover returns a boolean if a field has been set.

### GetBillableForecast

`func (o *Member) GetBillableForecast() float64`

GetBillableForecast returns the BillableForecast field if non-nil, zero value otherwise.

### GetBillableForecastOk

`func (o *Member) GetBillableForecastOk() (*float64, bool)`

GetBillableForecastOk returns a tuple with the BillableForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableForecast

`func (o *Member) SetBillableForecast(v float64)`

SetBillableForecast sets BillableForecast field to given value.

### HasBillableForecast

`func (o *Member) HasBillableForecast() bool`

HasBillableForecast returns a boolean if a field has been set.

### SetBillableForecastNil

`func (o *Member) SetBillableForecastNil(b bool)`

 SetBillableForecastNil sets the value for BillableForecast to be an explicit nil

### UnsetBillableForecast
`func (o *Member) UnsetBillableForecast()`

UnsetBillableForecast ensures that no value is present for BillableForecast, not even an explicit nil
### GetDailyCapacity

`func (o *Member) GetDailyCapacity() float64`

GetDailyCapacity returns the DailyCapacity field if non-nil, zero value otherwise.

### GetDailyCapacityOk

`func (o *Member) GetDailyCapacityOk() (*float64, bool)`

GetDailyCapacityOk returns a tuple with the DailyCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyCapacity

`func (o *Member) SetDailyCapacity(v float64)`

SetDailyCapacity sets DailyCapacity field to given value.

### HasDailyCapacity

`func (o *Member) HasDailyCapacity() bool`

HasDailyCapacity returns a boolean if a field has been set.

### SetDailyCapacityNil

`func (o *Member) SetDailyCapacityNil(b bool)`

 SetDailyCapacityNil sets the value for DailyCapacity to be an explicit nil

### UnsetDailyCapacity
`func (o *Member) UnsetDailyCapacity()`

UnsetDailyCapacity ensures that no value is present for DailyCapacity, not even an explicit nil
### GetHourlyCost

`func (o *Member) GetHourlyCost() float64`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *Member) GetHourlyCostOk() (*float64, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *Member) SetHourlyCost(v float64)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *Member) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### SetHourlyCostNil

`func (o *Member) SetHourlyCostNil(b bool)`

 SetHourlyCostNil sets the value for HourlyCost to be an explicit nil

### UnsetHourlyCost
`func (o *Member) UnsetHourlyCost()`

UnsetHourlyCost ensures that no value is present for HourlyCost, not even an explicit nil
### GetHourlyRate

`func (o *Member) GetHourlyRate() float64`

GetHourlyRate returns the HourlyRate field if non-nil, zero value otherwise.

### GetHourlyRateOk

`func (o *Member) GetHourlyRateOk() (*float64, bool)`

GetHourlyRateOk returns a tuple with the HourlyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyRate

`func (o *Member) SetHourlyRate(v float64)`

SetHourlyRate sets HourlyRate field to given value.

### HasHourlyRate

`func (o *Member) HasHourlyRate() bool`

HasHourlyRate returns a boolean if a field has been set.

### SetHourlyRateNil

`func (o *Member) SetHourlyRateNil(b bool)`

 SetHourlyRateNil sets the value for HourlyRate to be an explicit nil

### UnsetHourlyRate
`func (o *Member) UnsetHourlyRate()`

UnsetHourlyRate ensures that no value is present for HourlyRate, not even an explicit nil
### GetIncludeInUtilizationReportingFlag

`func (o *Member) GetIncludeInUtilizationReportingFlag() bool`

GetIncludeInUtilizationReportingFlag returns the IncludeInUtilizationReportingFlag field if non-nil, zero value otherwise.

### GetIncludeInUtilizationReportingFlagOk

`func (o *Member) GetIncludeInUtilizationReportingFlagOk() (*bool, bool)`

GetIncludeInUtilizationReportingFlagOk returns a tuple with the IncludeInUtilizationReportingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInUtilizationReportingFlag

`func (o *Member) SetIncludeInUtilizationReportingFlag(v bool)`

SetIncludeInUtilizationReportingFlag sets IncludeInUtilizationReportingFlag field to given value.

### HasIncludeInUtilizationReportingFlag

`func (o *Member) HasIncludeInUtilizationReportingFlag() bool`

HasIncludeInUtilizationReportingFlag returns a boolean if a field has been set.

### SetIncludeInUtilizationReportingFlagNil

`func (o *Member) SetIncludeInUtilizationReportingFlagNil(b bool)`

 SetIncludeInUtilizationReportingFlagNil sets the value for IncludeInUtilizationReportingFlag to be an explicit nil

### UnsetIncludeInUtilizationReportingFlag
`func (o *Member) UnsetIncludeInUtilizationReportingFlag()`

UnsetIncludeInUtilizationReportingFlag ensures that no value is present for IncludeInUtilizationReportingFlag, not even an explicit nil
### GetRequireExpenseEntryFlag

`func (o *Member) GetRequireExpenseEntryFlag() bool`

GetRequireExpenseEntryFlag returns the RequireExpenseEntryFlag field if non-nil, zero value otherwise.

### GetRequireExpenseEntryFlagOk

`func (o *Member) GetRequireExpenseEntryFlagOk() (*bool, bool)`

GetRequireExpenseEntryFlagOk returns a tuple with the RequireExpenseEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExpenseEntryFlag

`func (o *Member) SetRequireExpenseEntryFlag(v bool)`

SetRequireExpenseEntryFlag sets RequireExpenseEntryFlag field to given value.

### HasRequireExpenseEntryFlag

`func (o *Member) HasRequireExpenseEntryFlag() bool`

HasRequireExpenseEntryFlag returns a boolean if a field has been set.

### SetRequireExpenseEntryFlagNil

`func (o *Member) SetRequireExpenseEntryFlagNil(b bool)`

 SetRequireExpenseEntryFlagNil sets the value for RequireExpenseEntryFlag to be an explicit nil

### UnsetRequireExpenseEntryFlag
`func (o *Member) UnsetRequireExpenseEntryFlag()`

UnsetRequireExpenseEntryFlag ensures that no value is present for RequireExpenseEntryFlag, not even an explicit nil
### GetRequireTimeSheetEntryFlag

`func (o *Member) GetRequireTimeSheetEntryFlag() bool`

GetRequireTimeSheetEntryFlag returns the RequireTimeSheetEntryFlag field if non-nil, zero value otherwise.

### GetRequireTimeSheetEntryFlagOk

`func (o *Member) GetRequireTimeSheetEntryFlagOk() (*bool, bool)`

GetRequireTimeSheetEntryFlagOk returns a tuple with the RequireTimeSheetEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTimeSheetEntryFlag

`func (o *Member) SetRequireTimeSheetEntryFlag(v bool)`

SetRequireTimeSheetEntryFlag sets RequireTimeSheetEntryFlag field to given value.

### HasRequireTimeSheetEntryFlag

`func (o *Member) HasRequireTimeSheetEntryFlag() bool`

HasRequireTimeSheetEntryFlag returns a boolean if a field has been set.

### SetRequireTimeSheetEntryFlagNil

`func (o *Member) SetRequireTimeSheetEntryFlagNil(b bool)`

 SetRequireTimeSheetEntryFlagNil sets the value for RequireTimeSheetEntryFlag to be an explicit nil

### UnsetRequireTimeSheetEntryFlag
`func (o *Member) UnsetRequireTimeSheetEntryFlag()`

UnsetRequireTimeSheetEntryFlag ensures that no value is present for RequireTimeSheetEntryFlag, not even an explicit nil
### GetRequireStartAndEndTimeOnTimeEntryFlag

`func (o *Member) GetRequireStartAndEndTimeOnTimeEntryFlag() bool`

GetRequireStartAndEndTimeOnTimeEntryFlag returns the RequireStartAndEndTimeOnTimeEntryFlag field if non-nil, zero value otherwise.

### GetRequireStartAndEndTimeOnTimeEntryFlagOk

`func (o *Member) GetRequireStartAndEndTimeOnTimeEntryFlagOk() (*bool, bool)`

GetRequireStartAndEndTimeOnTimeEntryFlagOk returns a tuple with the RequireStartAndEndTimeOnTimeEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireStartAndEndTimeOnTimeEntryFlag

`func (o *Member) SetRequireStartAndEndTimeOnTimeEntryFlag(v bool)`

SetRequireStartAndEndTimeOnTimeEntryFlag sets RequireStartAndEndTimeOnTimeEntryFlag field to given value.

### HasRequireStartAndEndTimeOnTimeEntryFlag

`func (o *Member) HasRequireStartAndEndTimeOnTimeEntryFlag() bool`

HasRequireStartAndEndTimeOnTimeEntryFlag returns a boolean if a field has been set.

### SetRequireStartAndEndTimeOnTimeEntryFlagNil

`func (o *Member) SetRequireStartAndEndTimeOnTimeEntryFlagNil(b bool)`

 SetRequireStartAndEndTimeOnTimeEntryFlagNil sets the value for RequireStartAndEndTimeOnTimeEntryFlag to be an explicit nil

### UnsetRequireStartAndEndTimeOnTimeEntryFlag
`func (o *Member) UnsetRequireStartAndEndTimeOnTimeEntryFlag()`

UnsetRequireStartAndEndTimeOnTimeEntryFlag ensures that no value is present for RequireStartAndEndTimeOnTimeEntryFlag, not even an explicit nil
### GetAllowInCellEntryOnTimeSheet

`func (o *Member) GetAllowInCellEntryOnTimeSheet() bool`

GetAllowInCellEntryOnTimeSheet returns the AllowInCellEntryOnTimeSheet field if non-nil, zero value otherwise.

### GetAllowInCellEntryOnTimeSheetOk

`func (o *Member) GetAllowInCellEntryOnTimeSheetOk() (*bool, bool)`

GetAllowInCellEntryOnTimeSheetOk returns a tuple with the AllowInCellEntryOnTimeSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInCellEntryOnTimeSheet

`func (o *Member) SetAllowInCellEntryOnTimeSheet(v bool)`

SetAllowInCellEntryOnTimeSheet sets AllowInCellEntryOnTimeSheet field to given value.

### HasAllowInCellEntryOnTimeSheet

`func (o *Member) HasAllowInCellEntryOnTimeSheet() bool`

HasAllowInCellEntryOnTimeSheet returns a boolean if a field has been set.

### SetAllowInCellEntryOnTimeSheetNil

`func (o *Member) SetAllowInCellEntryOnTimeSheetNil(b bool)`

 SetAllowInCellEntryOnTimeSheetNil sets the value for AllowInCellEntryOnTimeSheet to be an explicit nil

### UnsetAllowInCellEntryOnTimeSheet
`func (o *Member) UnsetAllowInCellEntryOnTimeSheet()`

UnsetAllowInCellEntryOnTimeSheet ensures that no value is present for AllowInCellEntryOnTimeSheet, not even an explicit nil
### GetEnterTimeAgainstCompanyFlag

`func (o *Member) GetEnterTimeAgainstCompanyFlag() bool`

GetEnterTimeAgainstCompanyFlag returns the EnterTimeAgainstCompanyFlag field if non-nil, zero value otherwise.

### GetEnterTimeAgainstCompanyFlagOk

`func (o *Member) GetEnterTimeAgainstCompanyFlagOk() (*bool, bool)`

GetEnterTimeAgainstCompanyFlagOk returns a tuple with the EnterTimeAgainstCompanyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterTimeAgainstCompanyFlag

`func (o *Member) SetEnterTimeAgainstCompanyFlag(v bool)`

SetEnterTimeAgainstCompanyFlag sets EnterTimeAgainstCompanyFlag field to given value.

### HasEnterTimeAgainstCompanyFlag

`func (o *Member) HasEnterTimeAgainstCompanyFlag() bool`

HasEnterTimeAgainstCompanyFlag returns a boolean if a field has been set.

### SetEnterTimeAgainstCompanyFlagNil

`func (o *Member) SetEnterTimeAgainstCompanyFlagNil(b bool)`

 SetEnterTimeAgainstCompanyFlagNil sets the value for EnterTimeAgainstCompanyFlag to be an explicit nil

### UnsetEnterTimeAgainstCompanyFlag
`func (o *Member) UnsetEnterTimeAgainstCompanyFlag()`

UnsetEnterTimeAgainstCompanyFlag ensures that no value is present for EnterTimeAgainstCompanyFlag, not even an explicit nil
### GetAllowExpensesEnteredAgainstCompaniesFlag

`func (o *Member) GetAllowExpensesEnteredAgainstCompaniesFlag() bool`

GetAllowExpensesEnteredAgainstCompaniesFlag returns the AllowExpensesEnteredAgainstCompaniesFlag field if non-nil, zero value otherwise.

### GetAllowExpensesEnteredAgainstCompaniesFlagOk

`func (o *Member) GetAllowExpensesEnteredAgainstCompaniesFlagOk() (*bool, bool)`

GetAllowExpensesEnteredAgainstCompaniesFlagOk returns a tuple with the AllowExpensesEnteredAgainstCompaniesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExpensesEnteredAgainstCompaniesFlag

`func (o *Member) SetAllowExpensesEnteredAgainstCompaniesFlag(v bool)`

SetAllowExpensesEnteredAgainstCompaniesFlag sets AllowExpensesEnteredAgainstCompaniesFlag field to given value.

### HasAllowExpensesEnteredAgainstCompaniesFlag

`func (o *Member) HasAllowExpensesEnteredAgainstCompaniesFlag() bool`

HasAllowExpensesEnteredAgainstCompaniesFlag returns a boolean if a field has been set.

### SetAllowExpensesEnteredAgainstCompaniesFlagNil

`func (o *Member) SetAllowExpensesEnteredAgainstCompaniesFlagNil(b bool)`

 SetAllowExpensesEnteredAgainstCompaniesFlagNil sets the value for AllowExpensesEnteredAgainstCompaniesFlag to be an explicit nil

### UnsetAllowExpensesEnteredAgainstCompaniesFlag
`func (o *Member) UnsetAllowExpensesEnteredAgainstCompaniesFlag()`

UnsetAllowExpensesEnteredAgainstCompaniesFlag ensures that no value is present for AllowExpensesEnteredAgainstCompaniesFlag, not even an explicit nil
### GetTimeReminderEmailFlag

`func (o *Member) GetTimeReminderEmailFlag() bool`

GetTimeReminderEmailFlag returns the TimeReminderEmailFlag field if non-nil, zero value otherwise.

### GetTimeReminderEmailFlagOk

`func (o *Member) GetTimeReminderEmailFlagOk() (*bool, bool)`

GetTimeReminderEmailFlagOk returns a tuple with the TimeReminderEmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeReminderEmailFlag

`func (o *Member) SetTimeReminderEmailFlag(v bool)`

SetTimeReminderEmailFlag sets TimeReminderEmailFlag field to given value.

### HasTimeReminderEmailFlag

`func (o *Member) HasTimeReminderEmailFlag() bool`

HasTimeReminderEmailFlag returns a boolean if a field has been set.

### SetTimeReminderEmailFlagNil

`func (o *Member) SetTimeReminderEmailFlagNil(b bool)`

 SetTimeReminderEmailFlagNil sets the value for TimeReminderEmailFlag to be an explicit nil

### UnsetTimeReminderEmailFlag
`func (o *Member) UnsetTimeReminderEmailFlag()`

UnsetTimeReminderEmailFlag ensures that no value is present for TimeReminderEmailFlag, not even an explicit nil
### GetDaysTolerance

`func (o *Member) GetDaysTolerance() int32`

GetDaysTolerance returns the DaysTolerance field if non-nil, zero value otherwise.

### GetDaysToleranceOk

`func (o *Member) GetDaysToleranceOk() (*int32, bool)`

GetDaysToleranceOk returns a tuple with the DaysTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysTolerance

`func (o *Member) SetDaysTolerance(v int32)`

SetDaysTolerance sets DaysTolerance field to given value.

### HasDaysTolerance

`func (o *Member) HasDaysTolerance() bool`

HasDaysTolerance returns a boolean if a field has been set.

### SetDaysToleranceNil

`func (o *Member) SetDaysToleranceNil(b bool)`

 SetDaysToleranceNil sets the value for DaysTolerance to be an explicit nil

### UnsetDaysTolerance
`func (o *Member) UnsetDaysTolerance()`

UnsetDaysTolerance ensures that no value is present for DaysTolerance, not even an explicit nil
### GetMinimumHours

`func (o *Member) GetMinimumHours() float64`

GetMinimumHours returns the MinimumHours field if non-nil, zero value otherwise.

### GetMinimumHoursOk

`func (o *Member) GetMinimumHoursOk() (*float64, bool)`

GetMinimumHoursOk returns a tuple with the MinimumHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumHours

`func (o *Member) SetMinimumHours(v float64)`

SetMinimumHours sets MinimumHours field to given value.

### HasMinimumHours

`func (o *Member) HasMinimumHours() bool`

HasMinimumHours returns a boolean if a field has been set.

### SetMinimumHoursNil

`func (o *Member) SetMinimumHoursNil(b bool)`

 SetMinimumHoursNil sets the value for MinimumHours to be an explicit nil

### UnsetMinimumHours
`func (o *Member) UnsetMinimumHours()`

UnsetMinimumHours ensures that no value is present for MinimumHours, not even an explicit nil
### GetTimeSheetStartDate

`func (o *Member) GetTimeSheetStartDate() string`

GetTimeSheetStartDate returns the TimeSheetStartDate field if non-nil, zero value otherwise.

### GetTimeSheetStartDateOk

`func (o *Member) GetTimeSheetStartDateOk() (*string, bool)`

GetTimeSheetStartDateOk returns a tuple with the TimeSheetStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSheetStartDate

`func (o *Member) SetTimeSheetStartDate(v string)`

SetTimeSheetStartDate sets TimeSheetStartDate field to given value.

### HasTimeSheetStartDate

`func (o *Member) HasTimeSheetStartDate() bool`

HasTimeSheetStartDate returns a boolean if a field has been set.

### GetServiceDefaultLocation

`func (o *Member) GetServiceDefaultLocation() SystemLocationReference`

GetServiceDefaultLocation returns the ServiceDefaultLocation field if non-nil, zero value otherwise.

### GetServiceDefaultLocationOk

`func (o *Member) GetServiceDefaultLocationOk() (*SystemLocationReference, bool)`

GetServiceDefaultLocationOk returns a tuple with the ServiceDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultLocation

`func (o *Member) SetServiceDefaultLocation(v SystemLocationReference)`

SetServiceDefaultLocation sets ServiceDefaultLocation field to given value.

### HasServiceDefaultLocation

`func (o *Member) HasServiceDefaultLocation() bool`

HasServiceDefaultLocation returns a boolean if a field has been set.

### GetServiceDefaultDepartment

`func (o *Member) GetServiceDefaultDepartment() SystemDepartmentReference`

GetServiceDefaultDepartment returns the ServiceDefaultDepartment field if non-nil, zero value otherwise.

### GetServiceDefaultDepartmentOk

`func (o *Member) GetServiceDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetServiceDefaultDepartmentOk returns a tuple with the ServiceDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultDepartment

`func (o *Member) SetServiceDefaultDepartment(v SystemDepartmentReference)`

SetServiceDefaultDepartment sets ServiceDefaultDepartment field to given value.

### HasServiceDefaultDepartment

`func (o *Member) HasServiceDefaultDepartment() bool`

HasServiceDefaultDepartment returns a boolean if a field has been set.

### GetServiceDefaultBoard

`func (o *Member) GetServiceDefaultBoard() BoardReference`

GetServiceDefaultBoard returns the ServiceDefaultBoard field if non-nil, zero value otherwise.

### GetServiceDefaultBoardOk

`func (o *Member) GetServiceDefaultBoardOk() (*BoardReference, bool)`

GetServiceDefaultBoardOk returns a tuple with the ServiceDefaultBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultBoard

`func (o *Member) SetServiceDefaultBoard(v BoardReference)`

SetServiceDefaultBoard sets ServiceDefaultBoard field to given value.

### HasServiceDefaultBoard

`func (o *Member) HasServiceDefaultBoard() bool`

HasServiceDefaultBoard returns a boolean if a field has been set.

### GetRestrictServiceDefaultLocationFlag

`func (o *Member) GetRestrictServiceDefaultLocationFlag() bool`

GetRestrictServiceDefaultLocationFlag returns the RestrictServiceDefaultLocationFlag field if non-nil, zero value otherwise.

### GetRestrictServiceDefaultLocationFlagOk

`func (o *Member) GetRestrictServiceDefaultLocationFlagOk() (*bool, bool)`

GetRestrictServiceDefaultLocationFlagOk returns a tuple with the RestrictServiceDefaultLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictServiceDefaultLocationFlag

`func (o *Member) SetRestrictServiceDefaultLocationFlag(v bool)`

SetRestrictServiceDefaultLocationFlag sets RestrictServiceDefaultLocationFlag field to given value.

### HasRestrictServiceDefaultLocationFlag

`func (o *Member) HasRestrictServiceDefaultLocationFlag() bool`

HasRestrictServiceDefaultLocationFlag returns a boolean if a field has been set.

### SetRestrictServiceDefaultLocationFlagNil

`func (o *Member) SetRestrictServiceDefaultLocationFlagNil(b bool)`

 SetRestrictServiceDefaultLocationFlagNil sets the value for RestrictServiceDefaultLocationFlag to be an explicit nil

### UnsetRestrictServiceDefaultLocationFlag
`func (o *Member) UnsetRestrictServiceDefaultLocationFlag()`

UnsetRestrictServiceDefaultLocationFlag ensures that no value is present for RestrictServiceDefaultLocationFlag, not even an explicit nil
### GetRestrictServiceDefaultDepartmentFlag

`func (o *Member) GetRestrictServiceDefaultDepartmentFlag() bool`

GetRestrictServiceDefaultDepartmentFlag returns the RestrictServiceDefaultDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictServiceDefaultDepartmentFlagOk

`func (o *Member) GetRestrictServiceDefaultDepartmentFlagOk() (*bool, bool)`

GetRestrictServiceDefaultDepartmentFlagOk returns a tuple with the RestrictServiceDefaultDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictServiceDefaultDepartmentFlag

`func (o *Member) SetRestrictServiceDefaultDepartmentFlag(v bool)`

SetRestrictServiceDefaultDepartmentFlag sets RestrictServiceDefaultDepartmentFlag field to given value.

### HasRestrictServiceDefaultDepartmentFlag

`func (o *Member) HasRestrictServiceDefaultDepartmentFlag() bool`

HasRestrictServiceDefaultDepartmentFlag returns a boolean if a field has been set.

### SetRestrictServiceDefaultDepartmentFlagNil

`func (o *Member) SetRestrictServiceDefaultDepartmentFlagNil(b bool)`

 SetRestrictServiceDefaultDepartmentFlagNil sets the value for RestrictServiceDefaultDepartmentFlag to be an explicit nil

### UnsetRestrictServiceDefaultDepartmentFlag
`func (o *Member) UnsetRestrictServiceDefaultDepartmentFlag()`

UnsetRestrictServiceDefaultDepartmentFlag ensures that no value is present for RestrictServiceDefaultDepartmentFlag, not even an explicit nil
### GetExcludedServiceBoardIds

`func (o *Member) GetExcludedServiceBoardIds() []int32`

GetExcludedServiceBoardIds returns the ExcludedServiceBoardIds field if non-nil, zero value otherwise.

### GetExcludedServiceBoardIdsOk

`func (o *Member) GetExcludedServiceBoardIdsOk() (*[]int32, bool)`

GetExcludedServiceBoardIdsOk returns a tuple with the ExcludedServiceBoardIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedServiceBoardIds

`func (o *Member) SetExcludedServiceBoardIds(v []int32)`

SetExcludedServiceBoardIds sets ExcludedServiceBoardIds field to given value.

### HasExcludedServiceBoardIds

`func (o *Member) HasExcludedServiceBoardIds() bool`

HasExcludedServiceBoardIds returns a boolean if a field has been set.

### GetTeams

`func (o *Member) GetTeams() []int32`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *Member) GetTeamsOk() (*[]int32, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *Member) SetTeams(v []int32)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *Member) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetServiceBoardTeamIds

`func (o *Member) GetServiceBoardTeamIds() []int32`

GetServiceBoardTeamIds returns the ServiceBoardTeamIds field if non-nil, zero value otherwise.

### GetServiceBoardTeamIdsOk

`func (o *Member) GetServiceBoardTeamIdsOk() (*[]int32, bool)`

GetServiceBoardTeamIdsOk returns a tuple with the ServiceBoardTeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBoardTeamIds

`func (o *Member) SetServiceBoardTeamIds(v []int32)`

SetServiceBoardTeamIds sets ServiceBoardTeamIds field to given value.

### HasServiceBoardTeamIds

`func (o *Member) HasServiceBoardTeamIds() bool`

HasServiceBoardTeamIds returns a boolean if a field has been set.

### GetProjectDefaultLocation

`func (o *Member) GetProjectDefaultLocation() SystemLocationReference`

GetProjectDefaultLocation returns the ProjectDefaultLocation field if non-nil, zero value otherwise.

### GetProjectDefaultLocationOk

`func (o *Member) GetProjectDefaultLocationOk() (*SystemLocationReference, bool)`

GetProjectDefaultLocationOk returns a tuple with the ProjectDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultLocation

`func (o *Member) SetProjectDefaultLocation(v SystemLocationReference)`

SetProjectDefaultLocation sets ProjectDefaultLocation field to given value.

### HasProjectDefaultLocation

`func (o *Member) HasProjectDefaultLocation() bool`

HasProjectDefaultLocation returns a boolean if a field has been set.

### GetProjectDefaultDepartment

`func (o *Member) GetProjectDefaultDepartment() SystemDepartmentReference`

GetProjectDefaultDepartment returns the ProjectDefaultDepartment field if non-nil, zero value otherwise.

### GetProjectDefaultDepartmentOk

`func (o *Member) GetProjectDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetProjectDefaultDepartmentOk returns a tuple with the ProjectDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultDepartment

`func (o *Member) SetProjectDefaultDepartment(v SystemDepartmentReference)`

SetProjectDefaultDepartment sets ProjectDefaultDepartment field to given value.

### HasProjectDefaultDepartment

`func (o *Member) HasProjectDefaultDepartment() bool`

HasProjectDefaultDepartment returns a boolean if a field has been set.

### GetProjectDefaultBoard

`func (o *Member) GetProjectDefaultBoard() ProjectBoardReference`

GetProjectDefaultBoard returns the ProjectDefaultBoard field if non-nil, zero value otherwise.

### GetProjectDefaultBoardOk

`func (o *Member) GetProjectDefaultBoardOk() (*ProjectBoardReference, bool)`

GetProjectDefaultBoardOk returns a tuple with the ProjectDefaultBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultBoard

`func (o *Member) SetProjectDefaultBoard(v ProjectBoardReference)`

SetProjectDefaultBoard sets ProjectDefaultBoard field to given value.

### HasProjectDefaultBoard

`func (o *Member) HasProjectDefaultBoard() bool`

HasProjectDefaultBoard returns a boolean if a field has been set.

### GetRestrictProjectDefaultLocationFlag

`func (o *Member) GetRestrictProjectDefaultLocationFlag() bool`

GetRestrictProjectDefaultLocationFlag returns the RestrictProjectDefaultLocationFlag field if non-nil, zero value otherwise.

### GetRestrictProjectDefaultLocationFlagOk

`func (o *Member) GetRestrictProjectDefaultLocationFlagOk() (*bool, bool)`

GetRestrictProjectDefaultLocationFlagOk returns a tuple with the RestrictProjectDefaultLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictProjectDefaultLocationFlag

`func (o *Member) SetRestrictProjectDefaultLocationFlag(v bool)`

SetRestrictProjectDefaultLocationFlag sets RestrictProjectDefaultLocationFlag field to given value.

### HasRestrictProjectDefaultLocationFlag

`func (o *Member) HasRestrictProjectDefaultLocationFlag() bool`

HasRestrictProjectDefaultLocationFlag returns a boolean if a field has been set.

### SetRestrictProjectDefaultLocationFlagNil

`func (o *Member) SetRestrictProjectDefaultLocationFlagNil(b bool)`

 SetRestrictProjectDefaultLocationFlagNil sets the value for RestrictProjectDefaultLocationFlag to be an explicit nil

### UnsetRestrictProjectDefaultLocationFlag
`func (o *Member) UnsetRestrictProjectDefaultLocationFlag()`

UnsetRestrictProjectDefaultLocationFlag ensures that no value is present for RestrictProjectDefaultLocationFlag, not even an explicit nil
### GetRestrictProjectDefaultDepartmentFlag

`func (o *Member) GetRestrictProjectDefaultDepartmentFlag() bool`

GetRestrictProjectDefaultDepartmentFlag returns the RestrictProjectDefaultDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictProjectDefaultDepartmentFlagOk

`func (o *Member) GetRestrictProjectDefaultDepartmentFlagOk() (*bool, bool)`

GetRestrictProjectDefaultDepartmentFlagOk returns a tuple with the RestrictProjectDefaultDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictProjectDefaultDepartmentFlag

`func (o *Member) SetRestrictProjectDefaultDepartmentFlag(v bool)`

SetRestrictProjectDefaultDepartmentFlag sets RestrictProjectDefaultDepartmentFlag field to given value.

### HasRestrictProjectDefaultDepartmentFlag

`func (o *Member) HasRestrictProjectDefaultDepartmentFlag() bool`

HasRestrictProjectDefaultDepartmentFlag returns a boolean if a field has been set.

### SetRestrictProjectDefaultDepartmentFlagNil

`func (o *Member) SetRestrictProjectDefaultDepartmentFlagNil(b bool)`

 SetRestrictProjectDefaultDepartmentFlagNil sets the value for RestrictProjectDefaultDepartmentFlag to be an explicit nil

### UnsetRestrictProjectDefaultDepartmentFlag
`func (o *Member) UnsetRestrictProjectDefaultDepartmentFlag()`

UnsetRestrictProjectDefaultDepartmentFlag ensures that no value is present for RestrictProjectDefaultDepartmentFlag, not even an explicit nil
### GetExcludedProjectBoardIds

`func (o *Member) GetExcludedProjectBoardIds() []int32`

GetExcludedProjectBoardIds returns the ExcludedProjectBoardIds field if non-nil, zero value otherwise.

### GetExcludedProjectBoardIdsOk

`func (o *Member) GetExcludedProjectBoardIdsOk() (*[]int32, bool)`

GetExcludedProjectBoardIdsOk returns a tuple with the ExcludedProjectBoardIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProjectBoardIds

`func (o *Member) SetExcludedProjectBoardIds(v []int32)`

SetExcludedProjectBoardIds sets ExcludedProjectBoardIds field to given value.

### HasExcludedProjectBoardIds

`func (o *Member) HasExcludedProjectBoardIds() bool`

HasExcludedProjectBoardIds returns a boolean if a field has been set.

### GetScheduleDefaultLocation

`func (o *Member) GetScheduleDefaultLocation() SystemLocationReference`

GetScheduleDefaultLocation returns the ScheduleDefaultLocation field if non-nil, zero value otherwise.

### GetScheduleDefaultLocationOk

`func (o *Member) GetScheduleDefaultLocationOk() (*SystemLocationReference, bool)`

GetScheduleDefaultLocationOk returns a tuple with the ScheduleDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDefaultLocation

`func (o *Member) SetScheduleDefaultLocation(v SystemLocationReference)`

SetScheduleDefaultLocation sets ScheduleDefaultLocation field to given value.

### HasScheduleDefaultLocation

`func (o *Member) HasScheduleDefaultLocation() bool`

HasScheduleDefaultLocation returns a boolean if a field has been set.

### GetScheduleDefaultDepartment

`func (o *Member) GetScheduleDefaultDepartment() SystemDepartmentReference`

GetScheduleDefaultDepartment returns the ScheduleDefaultDepartment field if non-nil, zero value otherwise.

### GetScheduleDefaultDepartmentOk

`func (o *Member) GetScheduleDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetScheduleDefaultDepartmentOk returns a tuple with the ScheduleDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDefaultDepartment

`func (o *Member) SetScheduleDefaultDepartment(v SystemDepartmentReference)`

SetScheduleDefaultDepartment sets ScheduleDefaultDepartment field to given value.

### HasScheduleDefaultDepartment

`func (o *Member) HasScheduleDefaultDepartment() bool`

HasScheduleDefaultDepartment returns a boolean if a field has been set.

### GetScheduleCapacity

`func (o *Member) GetScheduleCapacity() float64`

GetScheduleCapacity returns the ScheduleCapacity field if non-nil, zero value otherwise.

### GetScheduleCapacityOk

`func (o *Member) GetScheduleCapacityOk() (*float64, bool)`

GetScheduleCapacityOk returns a tuple with the ScheduleCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCapacity

`func (o *Member) SetScheduleCapacity(v float64)`

SetScheduleCapacity sets ScheduleCapacity field to given value.

### HasScheduleCapacity

`func (o *Member) HasScheduleCapacity() bool`

HasScheduleCapacity returns a boolean if a field has been set.

### SetScheduleCapacityNil

`func (o *Member) SetScheduleCapacityNil(b bool)`

 SetScheduleCapacityNil sets the value for ScheduleCapacity to be an explicit nil

### UnsetScheduleCapacity
`func (o *Member) UnsetScheduleCapacity()`

UnsetScheduleCapacity ensures that no value is present for ScheduleCapacity, not even an explicit nil
### GetServiceLocation

`func (o *Member) GetServiceLocation() ServiceLocationReference`

GetServiceLocation returns the ServiceLocation field if non-nil, zero value otherwise.

### GetServiceLocationOk

`func (o *Member) GetServiceLocationOk() (*ServiceLocationReference, bool)`

GetServiceLocationOk returns a tuple with the ServiceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLocation

`func (o *Member) SetServiceLocation(v ServiceLocationReference)`

SetServiceLocation sets ServiceLocation field to given value.

### HasServiceLocation

`func (o *Member) HasServiceLocation() bool`

HasServiceLocation returns a boolean if a field has been set.

### GetRestrictScheduleFlag

`func (o *Member) GetRestrictScheduleFlag() bool`

GetRestrictScheduleFlag returns the RestrictScheduleFlag field if non-nil, zero value otherwise.

### GetRestrictScheduleFlagOk

`func (o *Member) GetRestrictScheduleFlagOk() (*bool, bool)`

GetRestrictScheduleFlagOk returns a tuple with the RestrictScheduleFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictScheduleFlag

`func (o *Member) SetRestrictScheduleFlag(v bool)`

SetRestrictScheduleFlag sets RestrictScheduleFlag field to given value.

### HasRestrictScheduleFlag

`func (o *Member) HasRestrictScheduleFlag() bool`

HasRestrictScheduleFlag returns a boolean if a field has been set.

### SetRestrictScheduleFlagNil

`func (o *Member) SetRestrictScheduleFlagNil(b bool)`

 SetRestrictScheduleFlagNil sets the value for RestrictScheduleFlag to be an explicit nil

### UnsetRestrictScheduleFlag
`func (o *Member) UnsetRestrictScheduleFlag()`

UnsetRestrictScheduleFlag ensures that no value is present for RestrictScheduleFlag, not even an explicit nil
### GetHideMemberInDispatchPortalFlag

`func (o *Member) GetHideMemberInDispatchPortalFlag() bool`

GetHideMemberInDispatchPortalFlag returns the HideMemberInDispatchPortalFlag field if non-nil, zero value otherwise.

### GetHideMemberInDispatchPortalFlagOk

`func (o *Member) GetHideMemberInDispatchPortalFlagOk() (*bool, bool)`

GetHideMemberInDispatchPortalFlagOk returns a tuple with the HideMemberInDispatchPortalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideMemberInDispatchPortalFlag

`func (o *Member) SetHideMemberInDispatchPortalFlag(v bool)`

SetHideMemberInDispatchPortalFlag sets HideMemberInDispatchPortalFlag field to given value.

### HasHideMemberInDispatchPortalFlag

`func (o *Member) HasHideMemberInDispatchPortalFlag() bool`

HasHideMemberInDispatchPortalFlag returns a boolean if a field has been set.

### SetHideMemberInDispatchPortalFlagNil

`func (o *Member) SetHideMemberInDispatchPortalFlagNil(b bool)`

 SetHideMemberInDispatchPortalFlagNil sets the value for HideMemberInDispatchPortalFlag to be an explicit nil

### UnsetHideMemberInDispatchPortalFlag
`func (o *Member) UnsetHideMemberInDispatchPortalFlag()`

UnsetHideMemberInDispatchPortalFlag ensures that no value is present for HideMemberInDispatchPortalFlag, not even an explicit nil
### GetCalendar

`func (o *Member) GetCalendar() CalendarReference`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *Member) GetCalendarOk() (*CalendarReference, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *Member) SetCalendar(v CalendarReference)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *Member) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### GetSalesDefaultLocation

`func (o *Member) GetSalesDefaultLocation() SystemLocationReference`

GetSalesDefaultLocation returns the SalesDefaultLocation field if non-nil, zero value otherwise.

### GetSalesDefaultLocationOk

`func (o *Member) GetSalesDefaultLocationOk() (*SystemLocationReference, bool)`

GetSalesDefaultLocationOk returns a tuple with the SalesDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDefaultLocation

`func (o *Member) SetSalesDefaultLocation(v SystemLocationReference)`

SetSalesDefaultLocation sets SalesDefaultLocation field to given value.

### HasSalesDefaultLocation

`func (o *Member) HasSalesDefaultLocation() bool`

HasSalesDefaultLocation returns a boolean if a field has been set.

### GetRestrictDefaultSalesTerritoryFlag

`func (o *Member) GetRestrictDefaultSalesTerritoryFlag() bool`

GetRestrictDefaultSalesTerritoryFlag returns the RestrictDefaultSalesTerritoryFlag field if non-nil, zero value otherwise.

### GetRestrictDefaultSalesTerritoryFlagOk

`func (o *Member) GetRestrictDefaultSalesTerritoryFlagOk() (*bool, bool)`

GetRestrictDefaultSalesTerritoryFlagOk returns a tuple with the RestrictDefaultSalesTerritoryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDefaultSalesTerritoryFlag

`func (o *Member) SetRestrictDefaultSalesTerritoryFlag(v bool)`

SetRestrictDefaultSalesTerritoryFlag sets RestrictDefaultSalesTerritoryFlag field to given value.

### HasRestrictDefaultSalesTerritoryFlag

`func (o *Member) HasRestrictDefaultSalesTerritoryFlag() bool`

HasRestrictDefaultSalesTerritoryFlag returns a boolean if a field has been set.

### SetRestrictDefaultSalesTerritoryFlagNil

`func (o *Member) SetRestrictDefaultSalesTerritoryFlagNil(b bool)`

 SetRestrictDefaultSalesTerritoryFlagNil sets the value for RestrictDefaultSalesTerritoryFlag to be an explicit nil

### UnsetRestrictDefaultSalesTerritoryFlag
`func (o *Member) UnsetRestrictDefaultSalesTerritoryFlag()`

UnsetRestrictDefaultSalesTerritoryFlag ensures that no value is present for RestrictDefaultSalesTerritoryFlag, not even an explicit nil
### GetWarehouse

`func (o *Member) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *Member) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *Member) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *Member) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetWarehouseBin

`func (o *Member) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *Member) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *Member) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.

### HasWarehouseBin

`func (o *Member) HasWarehouseBin() bool`

HasWarehouseBin returns a boolean if a field has been set.

### GetRestrictDefaultWarehouseFlag

`func (o *Member) GetRestrictDefaultWarehouseFlag() bool`

GetRestrictDefaultWarehouseFlag returns the RestrictDefaultWarehouseFlag field if non-nil, zero value otherwise.

### GetRestrictDefaultWarehouseFlagOk

`func (o *Member) GetRestrictDefaultWarehouseFlagOk() (*bool, bool)`

GetRestrictDefaultWarehouseFlagOk returns a tuple with the RestrictDefaultWarehouseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDefaultWarehouseFlag

`func (o *Member) SetRestrictDefaultWarehouseFlag(v bool)`

SetRestrictDefaultWarehouseFlag sets RestrictDefaultWarehouseFlag field to given value.

### HasRestrictDefaultWarehouseFlag

`func (o *Member) HasRestrictDefaultWarehouseFlag() bool`

HasRestrictDefaultWarehouseFlag returns a boolean if a field has been set.

### SetRestrictDefaultWarehouseFlagNil

`func (o *Member) SetRestrictDefaultWarehouseFlagNil(b bool)`

 SetRestrictDefaultWarehouseFlagNil sets the value for RestrictDefaultWarehouseFlag to be an explicit nil

### UnsetRestrictDefaultWarehouseFlag
`func (o *Member) UnsetRestrictDefaultWarehouseFlag()`

UnsetRestrictDefaultWarehouseFlag ensures that no value is present for RestrictDefaultWarehouseFlag, not even an explicit nil
### GetRestrictDefaultWarehouseBinFlag

`func (o *Member) GetRestrictDefaultWarehouseBinFlag() bool`

GetRestrictDefaultWarehouseBinFlag returns the RestrictDefaultWarehouseBinFlag field if non-nil, zero value otherwise.

### GetRestrictDefaultWarehouseBinFlagOk

`func (o *Member) GetRestrictDefaultWarehouseBinFlagOk() (*bool, bool)`

GetRestrictDefaultWarehouseBinFlagOk returns a tuple with the RestrictDefaultWarehouseBinFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDefaultWarehouseBinFlag

`func (o *Member) SetRestrictDefaultWarehouseBinFlag(v bool)`

SetRestrictDefaultWarehouseBinFlag sets RestrictDefaultWarehouseBinFlag field to given value.

### HasRestrictDefaultWarehouseBinFlag

`func (o *Member) HasRestrictDefaultWarehouseBinFlag() bool`

HasRestrictDefaultWarehouseBinFlag returns a boolean if a field has been set.

### SetRestrictDefaultWarehouseBinFlagNil

`func (o *Member) SetRestrictDefaultWarehouseBinFlagNil(b bool)`

 SetRestrictDefaultWarehouseBinFlagNil sets the value for RestrictDefaultWarehouseBinFlag to be an explicit nil

### UnsetRestrictDefaultWarehouseBinFlag
`func (o *Member) UnsetRestrictDefaultWarehouseBinFlag()`

UnsetRestrictDefaultWarehouseBinFlag ensures that no value is present for RestrictDefaultWarehouseBinFlag, not even an explicit nil
### GetCompanyActivityTabFormat

`func (o *Member) GetCompanyActivityTabFormat() string`

GetCompanyActivityTabFormat returns the CompanyActivityTabFormat field if non-nil, zero value otherwise.

### GetCompanyActivityTabFormatOk

`func (o *Member) GetCompanyActivityTabFormatOk() (*string, bool)`

GetCompanyActivityTabFormatOk returns a tuple with the CompanyActivityTabFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyActivityTabFormat

`func (o *Member) SetCompanyActivityTabFormat(v string)`

SetCompanyActivityTabFormat sets CompanyActivityTabFormat field to given value.

### HasCompanyActivityTabFormat

`func (o *Member) HasCompanyActivityTabFormat() bool`

HasCompanyActivityTabFormat returns a boolean if a field has been set.

### SetCompanyActivityTabFormatNil

`func (o *Member) SetCompanyActivityTabFormatNil(b bool)`

 SetCompanyActivityTabFormatNil sets the value for CompanyActivityTabFormat to be an explicit nil

### UnsetCompanyActivityTabFormat
`func (o *Member) UnsetCompanyActivityTabFormat()`

UnsetCompanyActivityTabFormat ensures that no value is present for CompanyActivityTabFormat, not even an explicit nil
### GetInvoiceTimeTabFormat

`func (o *Member) GetInvoiceTimeTabFormat() string`

GetInvoiceTimeTabFormat returns the InvoiceTimeTabFormat field if non-nil, zero value otherwise.

### GetInvoiceTimeTabFormatOk

`func (o *Member) GetInvoiceTimeTabFormatOk() (*string, bool)`

GetInvoiceTimeTabFormatOk returns a tuple with the InvoiceTimeTabFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTimeTabFormat

`func (o *Member) SetInvoiceTimeTabFormat(v string)`

SetInvoiceTimeTabFormat sets InvoiceTimeTabFormat field to given value.

### HasInvoiceTimeTabFormat

`func (o *Member) HasInvoiceTimeTabFormat() bool`

HasInvoiceTimeTabFormat returns a boolean if a field has been set.

### SetInvoiceTimeTabFormatNil

`func (o *Member) SetInvoiceTimeTabFormatNil(b bool)`

 SetInvoiceTimeTabFormatNil sets the value for InvoiceTimeTabFormat to be an explicit nil

### UnsetInvoiceTimeTabFormat
`func (o *Member) UnsetInvoiceTimeTabFormat()`

UnsetInvoiceTimeTabFormat ensures that no value is present for InvoiceTimeTabFormat, not even an explicit nil
### GetInvoiceScreenDefaultTabFormat

`func (o *Member) GetInvoiceScreenDefaultTabFormat() string`

GetInvoiceScreenDefaultTabFormat returns the InvoiceScreenDefaultTabFormat field if non-nil, zero value otherwise.

### GetInvoiceScreenDefaultTabFormatOk

`func (o *Member) GetInvoiceScreenDefaultTabFormatOk() (*string, bool)`

GetInvoiceScreenDefaultTabFormatOk returns a tuple with the InvoiceScreenDefaultTabFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceScreenDefaultTabFormat

`func (o *Member) SetInvoiceScreenDefaultTabFormat(v string)`

SetInvoiceScreenDefaultTabFormat sets InvoiceScreenDefaultTabFormat field to given value.

### HasInvoiceScreenDefaultTabFormat

`func (o *Member) HasInvoiceScreenDefaultTabFormat() bool`

HasInvoiceScreenDefaultTabFormat returns a boolean if a field has been set.

### SetInvoiceScreenDefaultTabFormatNil

`func (o *Member) SetInvoiceScreenDefaultTabFormatNil(b bool)`

 SetInvoiceScreenDefaultTabFormatNil sets the value for InvoiceScreenDefaultTabFormat to be an explicit nil

### UnsetInvoiceScreenDefaultTabFormat
`func (o *Member) UnsetInvoiceScreenDefaultTabFormat()`

UnsetInvoiceScreenDefaultTabFormat ensures that no value is present for InvoiceScreenDefaultTabFormat, not even an explicit nil
### GetInvoicingDisplayOptions

`func (o *Member) GetInvoicingDisplayOptions() string`

GetInvoicingDisplayOptions returns the InvoicingDisplayOptions field if non-nil, zero value otherwise.

### GetInvoicingDisplayOptionsOk

`func (o *Member) GetInvoicingDisplayOptionsOk() (*string, bool)`

GetInvoicingDisplayOptionsOk returns a tuple with the InvoicingDisplayOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingDisplayOptions

`func (o *Member) SetInvoicingDisplayOptions(v string)`

SetInvoicingDisplayOptions sets InvoicingDisplayOptions field to given value.

### HasInvoicingDisplayOptions

`func (o *Member) HasInvoicingDisplayOptions() bool`

HasInvoicingDisplayOptions returns a boolean if a field has been set.

### SetInvoicingDisplayOptionsNil

`func (o *Member) SetInvoicingDisplayOptionsNil(b bool)`

 SetInvoicingDisplayOptionsNil sets the value for InvoicingDisplayOptions to be an explicit nil

### UnsetInvoicingDisplayOptions
`func (o *Member) UnsetInvoicingDisplayOptions()`

UnsetInvoicingDisplayOptions ensures that no value is present for InvoicingDisplayOptions, not even an explicit nil
### GetAgreementInvoicingDisplayOptions

`func (o *Member) GetAgreementInvoicingDisplayOptions() string`

GetAgreementInvoicingDisplayOptions returns the AgreementInvoicingDisplayOptions field if non-nil, zero value otherwise.

### GetAgreementInvoicingDisplayOptionsOk

`func (o *Member) GetAgreementInvoicingDisplayOptionsOk() (*string, bool)`

GetAgreementInvoicingDisplayOptionsOk returns a tuple with the AgreementInvoicingDisplayOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementInvoicingDisplayOptions

`func (o *Member) SetAgreementInvoicingDisplayOptions(v string)`

SetAgreementInvoicingDisplayOptions sets AgreementInvoicingDisplayOptions field to given value.

### HasAgreementInvoicingDisplayOptions

`func (o *Member) HasAgreementInvoicingDisplayOptions() bool`

HasAgreementInvoicingDisplayOptions returns a boolean if a field has been set.

### SetAgreementInvoicingDisplayOptionsNil

`func (o *Member) SetAgreementInvoicingDisplayOptionsNil(b bool)`

 SetAgreementInvoicingDisplayOptionsNil sets the value for AgreementInvoicingDisplayOptions to be an explicit nil

### UnsetAgreementInvoicingDisplayOptions
`func (o *Member) UnsetAgreementInvoicingDisplayOptions()`

UnsetAgreementInvoicingDisplayOptions ensures that no value is present for AgreementInvoicingDisplayOptions, not even an explicit nil
### GetAutoStartStopwatch

`func (o *Member) GetAutoStartStopwatch() bool`

GetAutoStartStopwatch returns the AutoStartStopwatch field if non-nil, zero value otherwise.

### GetAutoStartStopwatchOk

`func (o *Member) GetAutoStartStopwatchOk() (*bool, bool)`

GetAutoStartStopwatchOk returns a tuple with the AutoStartStopwatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStartStopwatch

`func (o *Member) SetAutoStartStopwatch(v bool)`

SetAutoStartStopwatch sets AutoStartStopwatch field to given value.

### HasAutoStartStopwatch

`func (o *Member) HasAutoStartStopwatch() bool`

HasAutoStartStopwatch returns a boolean if a field has been set.

### SetAutoStartStopwatchNil

`func (o *Member) SetAutoStartStopwatchNil(b bool)`

 SetAutoStartStopwatchNil sets the value for AutoStartStopwatch to be an explicit nil

### UnsetAutoStartStopwatch
`func (o *Member) UnsetAutoStartStopwatch()`

UnsetAutoStartStopwatch ensures that no value is present for AutoStartStopwatch, not even an explicit nil
### GetAutoPopupQuickNotesWithStopwatch

`func (o *Member) GetAutoPopupQuickNotesWithStopwatch() bool`

GetAutoPopupQuickNotesWithStopwatch returns the AutoPopupQuickNotesWithStopwatch field if non-nil, zero value otherwise.

### GetAutoPopupQuickNotesWithStopwatchOk

`func (o *Member) GetAutoPopupQuickNotesWithStopwatchOk() (*bool, bool)`

GetAutoPopupQuickNotesWithStopwatchOk returns a tuple with the AutoPopupQuickNotesWithStopwatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPopupQuickNotesWithStopwatch

`func (o *Member) SetAutoPopupQuickNotesWithStopwatch(v bool)`

SetAutoPopupQuickNotesWithStopwatch sets AutoPopupQuickNotesWithStopwatch field to given value.

### HasAutoPopupQuickNotesWithStopwatch

`func (o *Member) HasAutoPopupQuickNotesWithStopwatch() bool`

HasAutoPopupQuickNotesWithStopwatch returns a boolean if a field has been set.

### SetAutoPopupQuickNotesWithStopwatchNil

`func (o *Member) SetAutoPopupQuickNotesWithStopwatchNil(b bool)`

 SetAutoPopupQuickNotesWithStopwatchNil sets the value for AutoPopupQuickNotesWithStopwatch to be an explicit nil

### UnsetAutoPopupQuickNotesWithStopwatch
`func (o *Member) UnsetAutoPopupQuickNotesWithStopwatch()`

UnsetAutoPopupQuickNotesWithStopwatch ensures that no value is present for AutoPopupQuickNotesWithStopwatch, not even an explicit nil
### GetGlobalSearchDefaultTicketFilter

`func (o *Member) GetGlobalSearchDefaultTicketFilter() string`

GetGlobalSearchDefaultTicketFilter returns the GlobalSearchDefaultTicketFilter field if non-nil, zero value otherwise.

### GetGlobalSearchDefaultTicketFilterOk

`func (o *Member) GetGlobalSearchDefaultTicketFilterOk() (*string, bool)`

GetGlobalSearchDefaultTicketFilterOk returns a tuple with the GlobalSearchDefaultTicketFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSearchDefaultTicketFilter

`func (o *Member) SetGlobalSearchDefaultTicketFilter(v string)`

SetGlobalSearchDefaultTicketFilter sets GlobalSearchDefaultTicketFilter field to given value.

### HasGlobalSearchDefaultTicketFilter

`func (o *Member) HasGlobalSearchDefaultTicketFilter() bool`

HasGlobalSearchDefaultTicketFilter returns a boolean if a field has been set.

### SetGlobalSearchDefaultTicketFilterNil

`func (o *Member) SetGlobalSearchDefaultTicketFilterNil(b bool)`

 SetGlobalSearchDefaultTicketFilterNil sets the value for GlobalSearchDefaultTicketFilter to be an explicit nil

### UnsetGlobalSearchDefaultTicketFilter
`func (o *Member) UnsetGlobalSearchDefaultTicketFilter()`

UnsetGlobalSearchDefaultTicketFilter ensures that no value is present for GlobalSearchDefaultTicketFilter, not even an explicit nil
### GetGlobalSearchDefaultSort

`func (o *Member) GetGlobalSearchDefaultSort() string`

GetGlobalSearchDefaultSort returns the GlobalSearchDefaultSort field if non-nil, zero value otherwise.

### GetGlobalSearchDefaultSortOk

`func (o *Member) GetGlobalSearchDefaultSortOk() (*string, bool)`

GetGlobalSearchDefaultSortOk returns a tuple with the GlobalSearchDefaultSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSearchDefaultSort

`func (o *Member) SetGlobalSearchDefaultSort(v string)`

SetGlobalSearchDefaultSort sets GlobalSearchDefaultSort field to given value.

### HasGlobalSearchDefaultSort

`func (o *Member) HasGlobalSearchDefaultSort() bool`

HasGlobalSearchDefaultSort returns a boolean if a field has been set.

### SetGlobalSearchDefaultSortNil

`func (o *Member) SetGlobalSearchDefaultSortNil(b bool)`

 SetGlobalSearchDefaultSortNil sets the value for GlobalSearchDefaultSort to be an explicit nil

### UnsetGlobalSearchDefaultSort
`func (o *Member) UnsetGlobalSearchDefaultSort()`

UnsetGlobalSearchDefaultSort ensures that no value is present for GlobalSearchDefaultSort, not even an explicit nil
### GetPhoneSource

`func (o *Member) GetPhoneSource() string`

GetPhoneSource returns the PhoneSource field if non-nil, zero value otherwise.

### GetPhoneSourceOk

`func (o *Member) GetPhoneSourceOk() (*string, bool)`

GetPhoneSourceOk returns a tuple with the PhoneSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneSource

`func (o *Member) SetPhoneSource(v string)`

SetPhoneSource sets PhoneSource field to given value.

### HasPhoneSource

`func (o *Member) HasPhoneSource() bool`

HasPhoneSource returns a boolean if a field has been set.

### GetInfo

`func (o *Member) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Member) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Member) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Member) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCopyPodLayouts

`func (o *Member) GetCopyPodLayouts() bool`

GetCopyPodLayouts returns the CopyPodLayouts field if non-nil, zero value otherwise.

### GetCopyPodLayoutsOk

`func (o *Member) GetCopyPodLayoutsOk() (*bool, bool)`

GetCopyPodLayoutsOk returns a tuple with the CopyPodLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyPodLayouts

`func (o *Member) SetCopyPodLayouts(v bool)`

SetCopyPodLayouts sets CopyPodLayouts field to given value.

### HasCopyPodLayouts

`func (o *Member) HasCopyPodLayouts() bool`

HasCopyPodLayouts returns a boolean if a field has been set.

### GetCopySharedDefaultViews

`func (o *Member) GetCopySharedDefaultViews() bool`

GetCopySharedDefaultViews returns the CopySharedDefaultViews field if non-nil, zero value otherwise.

### GetCopySharedDefaultViewsOk

`func (o *Member) GetCopySharedDefaultViewsOk() (*bool, bool)`

GetCopySharedDefaultViewsOk returns a tuple with the CopySharedDefaultViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySharedDefaultViews

`func (o *Member) SetCopySharedDefaultViews(v bool)`

SetCopySharedDefaultViews sets CopySharedDefaultViews field to given value.

### HasCopySharedDefaultViews

`func (o *Member) HasCopySharedDefaultViews() bool`

HasCopySharedDefaultViews returns a boolean if a field has been set.

### GetCopyColumnLayoutsAndFilters

`func (o *Member) GetCopyColumnLayoutsAndFilters() bool`

GetCopyColumnLayoutsAndFilters returns the CopyColumnLayoutsAndFilters field if non-nil, zero value otherwise.

### GetCopyColumnLayoutsAndFiltersOk

`func (o *Member) GetCopyColumnLayoutsAndFiltersOk() (*bool, bool)`

GetCopyColumnLayoutsAndFiltersOk returns a tuple with the CopyColumnLayoutsAndFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyColumnLayoutsAndFilters

`func (o *Member) SetCopyColumnLayoutsAndFilters(v bool)`

SetCopyColumnLayoutsAndFilters sets CopyColumnLayoutsAndFilters field to given value.

### HasCopyColumnLayoutsAndFilters

`func (o *Member) HasCopyColumnLayoutsAndFilters() bool`

HasCopyColumnLayoutsAndFilters returns a boolean if a field has been set.

### GetFromMemberRecId

`func (o *Member) GetFromMemberRecId() int32`

GetFromMemberRecId returns the FromMemberRecId field if non-nil, zero value otherwise.

### GetFromMemberRecIdOk

`func (o *Member) GetFromMemberRecIdOk() (*int32, bool)`

GetFromMemberRecIdOk returns a tuple with the FromMemberRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromMemberRecId

`func (o *Member) SetFromMemberRecId(v int32)`

SetFromMemberRecId sets FromMemberRecId field to given value.

### HasFromMemberRecId

`func (o *Member) HasFromMemberRecId() bool`

HasFromMemberRecId returns a boolean if a field has been set.

### GetFromMemberTemplateRecId

`func (o *Member) GetFromMemberTemplateRecId() int32`

GetFromMemberTemplateRecId returns the FromMemberTemplateRecId field if non-nil, zero value otherwise.

### GetFromMemberTemplateRecIdOk

`func (o *Member) GetFromMemberTemplateRecIdOk() (*int32, bool)`

GetFromMemberTemplateRecIdOk returns a tuple with the FromMemberTemplateRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromMemberTemplateRecId

`func (o *Member) SetFromMemberTemplateRecId(v int32)`

SetFromMemberTemplateRecId sets FromMemberTemplateRecId field to given value.

### HasFromMemberTemplateRecId

`func (o *Member) HasFromMemberTemplateRecId() bool`

HasFromMemberTemplateRecId returns a boolean if a field has been set.

### GetCustomFields

`func (o *Member) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Member) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Member) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Member) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


