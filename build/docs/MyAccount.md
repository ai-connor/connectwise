# MyAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  Max length: 15; | 
**Password** | Pointer to **string** | ConditionallyRequired. API Member will get random password generated Max length: 60; | [optional] 
**FirstName** | **string** |  Max length: 30; | 
**MiddleInitial** | Pointer to **string** |  Max length: 1; | [optional] 
**LastName** | **string** |  Max length: 30; | 
**Title** | Pointer to **string** |  Max length: 50; | [optional] 
**ReportCard** | Pointer to [**ReportCardReference**](ReportCardReference.md) |  | [optional] 
**LicenseClass** | **NullableString** | F &#x3D; Full Member, A &#x3D; API Member, C &#x3D; StreamlineIT Member, X &#x3D; Subcontractor Member | 
**DisableOnlineFlag** | Pointer to **NullableBool** |  | [optional] 
**EnableMobileFlag** | Pointer to **NullableBool** |  | [optional] 
**Type** | Pointer to [**MemberTypeReference**](MemberTypeReference.md) |  | [optional] 
**EmployeeIdentifer** | Pointer to **string** |  Max length: 10; | [optional] 
**VendorNumber** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**TimeZone** | [**TimeZoneSetupReference**](TimeZoneSetupReference.md) |  | 
**Country** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**ServiceBoardTeamIds** | Pointer to **[]int32** |  | [optional] 
**EnableMobileGpsFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveDate** | Pointer to **time.Time** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**LastLogin** | Pointer to **string** |  | [optional] 
**Photo** | Pointer to [**DocumentReference**](DocumentReference.md) |  | [optional] 
**PartnerPortalFlag** | Pointer to **NullableBool** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**StsUserAdminUrl** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**ToastNotificationFlag** | Pointer to **NullableBool** |  | [optional] 
**MemberPersonas** | Pointer to **[]int32** |  | [optional] 
**Office365** | Pointer to [**MemberOffice365**](MemberOffice365.md) |  | [optional] 
**OfficeEmail** | Pointer to **string** |  Max length: 250; | [optional] 
**OfficePhone** | Pointer to **string** |  Max length: 15; | [optional] 
**OfficeExtension** | Pointer to **string** |  Max length: 10; | [optional] 
**MobileEmail** | Pointer to **string** |  Max length: 250; | [optional] 
**MobilePhone** | Pointer to **string** |  Max length: 15; | [optional] 
**MobileExtension** | Pointer to **string** |  Max length: 10; | [optional] 
**HomeEmail** | Pointer to **string** |  Max length: 250; | [optional] 
**HomePhone** | Pointer to **string** |  Max length: 15; | [optional] 
**HomeExtension** | Pointer to **string** |  Max length: 10; | [optional] 
**DefaultEmail** | **NullableString** |  | 
**PrimaryEmail** | Pointer to **string** |  Max length: 250; | [optional] 
**DefaultPhone** | **NullableString** |  | 
**DefaultLocation** | [**SystemLocationReference**](SystemLocationReference.md) |  | 
**DefaultDepartment** | [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | 
**ReportsTo** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**WorkRole** | [**WorkRoleReference**](WorkRoleReference.md) |  | 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**TimeApprover** | [**MemberReference**](MemberReference.md) |  | 
**ExpenseApprover** | [**MemberReference**](MemberReference.md) |  | 
**BillableForecast** | Pointer to **NullableFloat64** |  | [optional] 
**DailyCapacity** | Pointer to **NullableFloat64** |  | [optional] 
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
**TimeSheetStartDate** | Pointer to **time.Time** |  | [optional] 
**HireDate** | **time.Time** |  | 
**ServiceDefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**ServiceDefaultDepartment** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**ServiceDefaultBoard** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**ProjectDefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**ProjectDefaultDepartment** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**ProjectDefaultBoard** | Pointer to [**ProjectBoardReference**](ProjectBoardReference.md) |  | [optional] 
**ScheduleDefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**ScheduleDefaultDepartment** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**ScheduleCapacity** | Pointer to **NullableFloat64** |  | [optional] 
**ServiceLocation** | Pointer to [**ServiceLocationReference**](ServiceLocationReference.md) |  | [optional] 
**HideMemberInDispatchPortalFlag** | Pointer to **NullableBool** |  | [optional] 
**Calendar** | Pointer to [**CalendarReference**](CalendarReference.md) |  | [optional] 
**SalesDefaultLocation** | [**SystemLocationReference**](SystemLocationReference.md) |  | 
**Warehouse** | Pointer to [**WarehouseReference**](WarehouseReference.md) |  | [optional] 
**WarehouseBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**MapiName** | Pointer to **string** |  | [optional] 
**CalendarSyncIntegrationFlag** | Pointer to **NullableBool** |  | [optional] 
**CompanyActivityTabFormat** | **NullableString** |  | 
**InvoiceTimeTabFormat** | **NullableString** |  | 
**InvoiceScreenDefaultTabFormat** | **NullableString** |  | 
**InvoicingDisplayOptions** | **NullableString** |  | 
**AgreementInvoicingDisplayOptions** | **NullableString** |  | 
**AuthenticationServiceType** | Pointer to **NullableString** |  | [optional] 
**TimebasedOneTimePasswordActivated** | Pointer to **NullableBool** |  | [optional] 
**DirectionalSync** | Pointer to [**DirectionalSyncReference**](DirectionalSyncReference.md) |  | [optional] 
**AutoStartStopwatch** | Pointer to **NullableBool** |  | [optional] 
**AutoPopupQuickNotesWithStopwatch** | Pointer to **NullableBool** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**GlobalSearchDefaultTicketFilter** | Pointer to **NullableString** |  | [optional] 
**GlobalSearchDefaultSort** | Pointer to **NullableString** |  | [optional] 
**PhoneSource** | Pointer to **string** |  | [optional] 
**PhoneIntegrationType** | Pointer to **NullableString** |  | [optional] 
**UseBrowserLanguageFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CopyPodLayouts** | Pointer to **bool** |  | [optional] 
**CopySharedDefaultViews** | Pointer to **bool** |  | [optional] 
**CopyColumnLayoutsAndFilters** | Pointer to **bool** |  | [optional] 
**FromMemberRecId** | Pointer to **int32** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewMyAccount

`func NewMyAccount(identifier string, firstName string, lastName string, licenseClass NullableString, timeZone TimeZoneSetupReference, defaultEmail NullableString, defaultPhone NullableString, defaultLocation SystemLocationReference, defaultDepartment SystemDepartmentReference, workRole WorkRoleReference, timeApprover MemberReference, expenseApprover MemberReference, hireDate time.Time, salesDefaultLocation SystemLocationReference, companyActivityTabFormat NullableString, invoiceTimeTabFormat NullableString, invoiceScreenDefaultTabFormat NullableString, invoicingDisplayOptions NullableString, agreementInvoicingDisplayOptions NullableString, ) *MyAccount`

NewMyAccount instantiates a new MyAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyAccountWithDefaults

`func NewMyAccountWithDefaults() *MyAccount`

NewMyAccountWithDefaults instantiates a new MyAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MyAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MyAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *MyAccount) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MyAccount) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MyAccount) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetPassword

`func (o *MyAccount) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MyAccount) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MyAccount) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MyAccount) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetFirstName

`func (o *MyAccount) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MyAccount) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MyAccount) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleInitial

`func (o *MyAccount) GetMiddleInitial() string`

GetMiddleInitial returns the MiddleInitial field if non-nil, zero value otherwise.

### GetMiddleInitialOk

`func (o *MyAccount) GetMiddleInitialOk() (*string, bool)`

GetMiddleInitialOk returns a tuple with the MiddleInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleInitial

`func (o *MyAccount) SetMiddleInitial(v string)`

SetMiddleInitial sets MiddleInitial field to given value.

### HasMiddleInitial

`func (o *MyAccount) HasMiddleInitial() bool`

HasMiddleInitial returns a boolean if a field has been set.

### GetLastName

`func (o *MyAccount) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MyAccount) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MyAccount) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetTitle

`func (o *MyAccount) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MyAccount) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MyAccount) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MyAccount) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetReportCard

`func (o *MyAccount) GetReportCard() ReportCardReference`

GetReportCard returns the ReportCard field if non-nil, zero value otherwise.

### GetReportCardOk

`func (o *MyAccount) GetReportCardOk() (*ReportCardReference, bool)`

GetReportCardOk returns a tuple with the ReportCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCard

`func (o *MyAccount) SetReportCard(v ReportCardReference)`

SetReportCard sets ReportCard field to given value.

### HasReportCard

`func (o *MyAccount) HasReportCard() bool`

HasReportCard returns a boolean if a field has been set.

### GetLicenseClass

`func (o *MyAccount) GetLicenseClass() string`

GetLicenseClass returns the LicenseClass field if non-nil, zero value otherwise.

### GetLicenseClassOk

`func (o *MyAccount) GetLicenseClassOk() (*string, bool)`

GetLicenseClassOk returns a tuple with the LicenseClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseClass

`func (o *MyAccount) SetLicenseClass(v string)`

SetLicenseClass sets LicenseClass field to given value.


### SetLicenseClassNil

`func (o *MyAccount) SetLicenseClassNil(b bool)`

 SetLicenseClassNil sets the value for LicenseClass to be an explicit nil

### UnsetLicenseClass
`func (o *MyAccount) UnsetLicenseClass()`

UnsetLicenseClass ensures that no value is present for LicenseClass, not even an explicit nil
### GetDisableOnlineFlag

`func (o *MyAccount) GetDisableOnlineFlag() bool`

GetDisableOnlineFlag returns the DisableOnlineFlag field if non-nil, zero value otherwise.

### GetDisableOnlineFlagOk

`func (o *MyAccount) GetDisableOnlineFlagOk() (*bool, bool)`

GetDisableOnlineFlagOk returns a tuple with the DisableOnlineFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOnlineFlag

`func (o *MyAccount) SetDisableOnlineFlag(v bool)`

SetDisableOnlineFlag sets DisableOnlineFlag field to given value.

### HasDisableOnlineFlag

`func (o *MyAccount) HasDisableOnlineFlag() bool`

HasDisableOnlineFlag returns a boolean if a field has been set.

### SetDisableOnlineFlagNil

`func (o *MyAccount) SetDisableOnlineFlagNil(b bool)`

 SetDisableOnlineFlagNil sets the value for DisableOnlineFlag to be an explicit nil

### UnsetDisableOnlineFlag
`func (o *MyAccount) UnsetDisableOnlineFlag()`

UnsetDisableOnlineFlag ensures that no value is present for DisableOnlineFlag, not even an explicit nil
### GetEnableMobileFlag

`func (o *MyAccount) GetEnableMobileFlag() bool`

GetEnableMobileFlag returns the EnableMobileFlag field if non-nil, zero value otherwise.

### GetEnableMobileFlagOk

`func (o *MyAccount) GetEnableMobileFlagOk() (*bool, bool)`

GetEnableMobileFlagOk returns a tuple with the EnableMobileFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMobileFlag

`func (o *MyAccount) SetEnableMobileFlag(v bool)`

SetEnableMobileFlag sets EnableMobileFlag field to given value.

### HasEnableMobileFlag

`func (o *MyAccount) HasEnableMobileFlag() bool`

HasEnableMobileFlag returns a boolean if a field has been set.

### SetEnableMobileFlagNil

`func (o *MyAccount) SetEnableMobileFlagNil(b bool)`

 SetEnableMobileFlagNil sets the value for EnableMobileFlag to be an explicit nil

### UnsetEnableMobileFlag
`func (o *MyAccount) UnsetEnableMobileFlag()`

UnsetEnableMobileFlag ensures that no value is present for EnableMobileFlag, not even an explicit nil
### GetType

`func (o *MyAccount) GetType() MemberTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MyAccount) GetTypeOk() (*MemberTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MyAccount) SetType(v MemberTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *MyAccount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmployeeIdentifer

`func (o *MyAccount) GetEmployeeIdentifer() string`

GetEmployeeIdentifer returns the EmployeeIdentifer field if non-nil, zero value otherwise.

### GetEmployeeIdentiferOk

`func (o *MyAccount) GetEmployeeIdentiferOk() (*string, bool)`

GetEmployeeIdentiferOk returns a tuple with the EmployeeIdentifer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeIdentifer

`func (o *MyAccount) SetEmployeeIdentifer(v string)`

SetEmployeeIdentifer sets EmployeeIdentifer field to given value.

### HasEmployeeIdentifer

`func (o *MyAccount) HasEmployeeIdentifer() bool`

HasEmployeeIdentifer returns a boolean if a field has been set.

### GetVendorNumber

`func (o *MyAccount) GetVendorNumber() string`

GetVendorNumber returns the VendorNumber field if non-nil, zero value otherwise.

### GetVendorNumberOk

`func (o *MyAccount) GetVendorNumberOk() (*string, bool)`

GetVendorNumberOk returns a tuple with the VendorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumber

`func (o *MyAccount) SetVendorNumber(v string)`

SetVendorNumber sets VendorNumber field to given value.

### HasVendorNumber

`func (o *MyAccount) HasVendorNumber() bool`

HasVendorNumber returns a boolean if a field has been set.

### GetNotes

`func (o *MyAccount) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MyAccount) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MyAccount) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MyAccount) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTimeZone

`func (o *MyAccount) GetTimeZone() TimeZoneSetupReference`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MyAccount) GetTimeZoneOk() (*TimeZoneSetupReference, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MyAccount) SetTimeZone(v TimeZoneSetupReference)`

SetTimeZone sets TimeZone field to given value.


### GetCountry

`func (o *MyAccount) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MyAccount) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MyAccount) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *MyAccount) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetServiceBoardTeamIds

`func (o *MyAccount) GetServiceBoardTeamIds() []int32`

GetServiceBoardTeamIds returns the ServiceBoardTeamIds field if non-nil, zero value otherwise.

### GetServiceBoardTeamIdsOk

`func (o *MyAccount) GetServiceBoardTeamIdsOk() (*[]int32, bool)`

GetServiceBoardTeamIdsOk returns a tuple with the ServiceBoardTeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBoardTeamIds

`func (o *MyAccount) SetServiceBoardTeamIds(v []int32)`

SetServiceBoardTeamIds sets ServiceBoardTeamIds field to given value.

### HasServiceBoardTeamIds

`func (o *MyAccount) HasServiceBoardTeamIds() bool`

HasServiceBoardTeamIds returns a boolean if a field has been set.

### GetEnableMobileGpsFlag

`func (o *MyAccount) GetEnableMobileGpsFlag() bool`

GetEnableMobileGpsFlag returns the EnableMobileGpsFlag field if non-nil, zero value otherwise.

### GetEnableMobileGpsFlagOk

`func (o *MyAccount) GetEnableMobileGpsFlagOk() (*bool, bool)`

GetEnableMobileGpsFlagOk returns a tuple with the EnableMobileGpsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMobileGpsFlag

`func (o *MyAccount) SetEnableMobileGpsFlag(v bool)`

SetEnableMobileGpsFlag sets EnableMobileGpsFlag field to given value.

### HasEnableMobileGpsFlag

`func (o *MyAccount) HasEnableMobileGpsFlag() bool`

HasEnableMobileGpsFlag returns a boolean if a field has been set.

### SetEnableMobileGpsFlagNil

`func (o *MyAccount) SetEnableMobileGpsFlagNil(b bool)`

 SetEnableMobileGpsFlagNil sets the value for EnableMobileGpsFlag to be an explicit nil

### UnsetEnableMobileGpsFlag
`func (o *MyAccount) UnsetEnableMobileGpsFlag()`

UnsetEnableMobileGpsFlag ensures that no value is present for EnableMobileGpsFlag, not even an explicit nil
### GetInactiveDate

`func (o *MyAccount) GetInactiveDate() time.Time`

GetInactiveDate returns the InactiveDate field if non-nil, zero value otherwise.

### GetInactiveDateOk

`func (o *MyAccount) GetInactiveDateOk() (*time.Time, bool)`

GetInactiveDateOk returns a tuple with the InactiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveDate

`func (o *MyAccount) SetInactiveDate(v time.Time)`

SetInactiveDate sets InactiveDate field to given value.

### HasInactiveDate

`func (o *MyAccount) HasInactiveDate() bool`

HasInactiveDate returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *MyAccount) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *MyAccount) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *MyAccount) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *MyAccount) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *MyAccount) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *MyAccount) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetLastLogin

`func (o *MyAccount) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *MyAccount) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *MyAccount) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *MyAccount) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetPhoto

`func (o *MyAccount) GetPhoto() DocumentReference`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MyAccount) GetPhotoOk() (*DocumentReference, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *MyAccount) SetPhoto(v DocumentReference)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *MyAccount) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetPartnerPortalFlag

`func (o *MyAccount) GetPartnerPortalFlag() bool`

GetPartnerPortalFlag returns the PartnerPortalFlag field if non-nil, zero value otherwise.

### GetPartnerPortalFlagOk

`func (o *MyAccount) GetPartnerPortalFlagOk() (*bool, bool)`

GetPartnerPortalFlagOk returns a tuple with the PartnerPortalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPortalFlag

`func (o *MyAccount) SetPartnerPortalFlag(v bool)`

SetPartnerPortalFlag sets PartnerPortalFlag field to given value.

### HasPartnerPortalFlag

`func (o *MyAccount) HasPartnerPortalFlag() bool`

HasPartnerPortalFlag returns a boolean if a field has been set.

### SetPartnerPortalFlagNil

`func (o *MyAccount) SetPartnerPortalFlagNil(b bool)`

 SetPartnerPortalFlagNil sets the value for PartnerPortalFlag to be an explicit nil

### UnsetPartnerPortalFlag
`func (o *MyAccount) UnsetPartnerPortalFlag()`

UnsetPartnerPortalFlag ensures that no value is present for PartnerPortalFlag, not even an explicit nil
### GetClientId

`func (o *MyAccount) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MyAccount) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MyAccount) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *MyAccount) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetStsUserAdminUrl

`func (o *MyAccount) GetStsUserAdminUrl() string`

GetStsUserAdminUrl returns the StsUserAdminUrl field if non-nil, zero value otherwise.

### GetStsUserAdminUrlOk

`func (o *MyAccount) GetStsUserAdminUrlOk() (*string, bool)`

GetStsUserAdminUrlOk returns a tuple with the StsUserAdminUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsUserAdminUrl

`func (o *MyAccount) SetStsUserAdminUrl(v string)`

SetStsUserAdminUrl sets StsUserAdminUrl field to given value.

### HasStsUserAdminUrl

`func (o *MyAccount) HasStsUserAdminUrl() bool`

HasStsUserAdminUrl returns a boolean if a field has been set.

### GetToken

`func (o *MyAccount) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MyAccount) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MyAccount) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MyAccount) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetToastNotificationFlag

`func (o *MyAccount) GetToastNotificationFlag() bool`

GetToastNotificationFlag returns the ToastNotificationFlag field if non-nil, zero value otherwise.

### GetToastNotificationFlagOk

`func (o *MyAccount) GetToastNotificationFlagOk() (*bool, bool)`

GetToastNotificationFlagOk returns a tuple with the ToastNotificationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToastNotificationFlag

`func (o *MyAccount) SetToastNotificationFlag(v bool)`

SetToastNotificationFlag sets ToastNotificationFlag field to given value.

### HasToastNotificationFlag

`func (o *MyAccount) HasToastNotificationFlag() bool`

HasToastNotificationFlag returns a boolean if a field has been set.

### SetToastNotificationFlagNil

`func (o *MyAccount) SetToastNotificationFlagNil(b bool)`

 SetToastNotificationFlagNil sets the value for ToastNotificationFlag to be an explicit nil

### UnsetToastNotificationFlag
`func (o *MyAccount) UnsetToastNotificationFlag()`

UnsetToastNotificationFlag ensures that no value is present for ToastNotificationFlag, not even an explicit nil
### GetMemberPersonas

`func (o *MyAccount) GetMemberPersonas() []int32`

GetMemberPersonas returns the MemberPersonas field if non-nil, zero value otherwise.

### GetMemberPersonasOk

`func (o *MyAccount) GetMemberPersonasOk() (*[]int32, bool)`

GetMemberPersonasOk returns a tuple with the MemberPersonas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPersonas

`func (o *MyAccount) SetMemberPersonas(v []int32)`

SetMemberPersonas sets MemberPersonas field to given value.

### HasMemberPersonas

`func (o *MyAccount) HasMemberPersonas() bool`

HasMemberPersonas returns a boolean if a field has been set.

### GetOffice365

`func (o *MyAccount) GetOffice365() MemberOffice365`

GetOffice365 returns the Office365 field if non-nil, zero value otherwise.

### GetOffice365Ok

`func (o *MyAccount) GetOffice365Ok() (*MemberOffice365, bool)`

GetOffice365Ok returns a tuple with the Office365 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365

`func (o *MyAccount) SetOffice365(v MemberOffice365)`

SetOffice365 sets Office365 field to given value.

### HasOffice365

`func (o *MyAccount) HasOffice365() bool`

HasOffice365 returns a boolean if a field has been set.

### GetOfficeEmail

`func (o *MyAccount) GetOfficeEmail() string`

GetOfficeEmail returns the OfficeEmail field if non-nil, zero value otherwise.

### GetOfficeEmailOk

`func (o *MyAccount) GetOfficeEmailOk() (*string, bool)`

GetOfficeEmailOk returns a tuple with the OfficeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeEmail

`func (o *MyAccount) SetOfficeEmail(v string)`

SetOfficeEmail sets OfficeEmail field to given value.

### HasOfficeEmail

`func (o *MyAccount) HasOfficeEmail() bool`

HasOfficeEmail returns a boolean if a field has been set.

### GetOfficePhone

`func (o *MyAccount) GetOfficePhone() string`

GetOfficePhone returns the OfficePhone field if non-nil, zero value otherwise.

### GetOfficePhoneOk

`func (o *MyAccount) GetOfficePhoneOk() (*string, bool)`

GetOfficePhoneOk returns a tuple with the OfficePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficePhone

`func (o *MyAccount) SetOfficePhone(v string)`

SetOfficePhone sets OfficePhone field to given value.

### HasOfficePhone

`func (o *MyAccount) HasOfficePhone() bool`

HasOfficePhone returns a boolean if a field has been set.

### GetOfficeExtension

`func (o *MyAccount) GetOfficeExtension() string`

GetOfficeExtension returns the OfficeExtension field if non-nil, zero value otherwise.

### GetOfficeExtensionOk

`func (o *MyAccount) GetOfficeExtensionOk() (*string, bool)`

GetOfficeExtensionOk returns a tuple with the OfficeExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeExtension

`func (o *MyAccount) SetOfficeExtension(v string)`

SetOfficeExtension sets OfficeExtension field to given value.

### HasOfficeExtension

`func (o *MyAccount) HasOfficeExtension() bool`

HasOfficeExtension returns a boolean if a field has been set.

### GetMobileEmail

`func (o *MyAccount) GetMobileEmail() string`

GetMobileEmail returns the MobileEmail field if non-nil, zero value otherwise.

### GetMobileEmailOk

`func (o *MyAccount) GetMobileEmailOk() (*string, bool)`

GetMobileEmailOk returns a tuple with the MobileEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileEmail

`func (o *MyAccount) SetMobileEmail(v string)`

SetMobileEmail sets MobileEmail field to given value.

### HasMobileEmail

`func (o *MyAccount) HasMobileEmail() bool`

HasMobileEmail returns a boolean if a field has been set.

### GetMobilePhone

`func (o *MyAccount) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *MyAccount) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *MyAccount) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *MyAccount) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetMobileExtension

`func (o *MyAccount) GetMobileExtension() string`

GetMobileExtension returns the MobileExtension field if non-nil, zero value otherwise.

### GetMobileExtensionOk

`func (o *MyAccount) GetMobileExtensionOk() (*string, bool)`

GetMobileExtensionOk returns a tuple with the MobileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileExtension

`func (o *MyAccount) SetMobileExtension(v string)`

SetMobileExtension sets MobileExtension field to given value.

### HasMobileExtension

`func (o *MyAccount) HasMobileExtension() bool`

HasMobileExtension returns a boolean if a field has been set.

### GetHomeEmail

`func (o *MyAccount) GetHomeEmail() string`

GetHomeEmail returns the HomeEmail field if non-nil, zero value otherwise.

### GetHomeEmailOk

`func (o *MyAccount) GetHomeEmailOk() (*string, bool)`

GetHomeEmailOk returns a tuple with the HomeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeEmail

`func (o *MyAccount) SetHomeEmail(v string)`

SetHomeEmail sets HomeEmail field to given value.

### HasHomeEmail

`func (o *MyAccount) HasHomeEmail() bool`

HasHomeEmail returns a boolean if a field has been set.

### GetHomePhone

`func (o *MyAccount) GetHomePhone() string`

GetHomePhone returns the HomePhone field if non-nil, zero value otherwise.

### GetHomePhoneOk

`func (o *MyAccount) GetHomePhoneOk() (*string, bool)`

GetHomePhoneOk returns a tuple with the HomePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePhone

`func (o *MyAccount) SetHomePhone(v string)`

SetHomePhone sets HomePhone field to given value.

### HasHomePhone

`func (o *MyAccount) HasHomePhone() bool`

HasHomePhone returns a boolean if a field has been set.

### GetHomeExtension

`func (o *MyAccount) GetHomeExtension() string`

GetHomeExtension returns the HomeExtension field if non-nil, zero value otherwise.

### GetHomeExtensionOk

`func (o *MyAccount) GetHomeExtensionOk() (*string, bool)`

GetHomeExtensionOk returns a tuple with the HomeExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeExtension

`func (o *MyAccount) SetHomeExtension(v string)`

SetHomeExtension sets HomeExtension field to given value.

### HasHomeExtension

`func (o *MyAccount) HasHomeExtension() bool`

HasHomeExtension returns a boolean if a field has been set.

### GetDefaultEmail

`func (o *MyAccount) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *MyAccount) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *MyAccount) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.


### SetDefaultEmailNil

`func (o *MyAccount) SetDefaultEmailNil(b bool)`

 SetDefaultEmailNil sets the value for DefaultEmail to be an explicit nil

### UnsetDefaultEmail
`func (o *MyAccount) UnsetDefaultEmail()`

UnsetDefaultEmail ensures that no value is present for DefaultEmail, not even an explicit nil
### GetPrimaryEmail

`func (o *MyAccount) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *MyAccount) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *MyAccount) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *MyAccount) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### GetDefaultPhone

`func (o *MyAccount) GetDefaultPhone() string`

GetDefaultPhone returns the DefaultPhone field if non-nil, zero value otherwise.

### GetDefaultPhoneOk

`func (o *MyAccount) GetDefaultPhoneOk() (*string, bool)`

GetDefaultPhoneOk returns a tuple with the DefaultPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhone

`func (o *MyAccount) SetDefaultPhone(v string)`

SetDefaultPhone sets DefaultPhone field to given value.


### SetDefaultPhoneNil

`func (o *MyAccount) SetDefaultPhoneNil(b bool)`

 SetDefaultPhoneNil sets the value for DefaultPhone to be an explicit nil

### UnsetDefaultPhone
`func (o *MyAccount) UnsetDefaultPhone()`

UnsetDefaultPhone ensures that no value is present for DefaultPhone, not even an explicit nil
### GetDefaultLocation

`func (o *MyAccount) GetDefaultLocation() SystemLocationReference`

GetDefaultLocation returns the DefaultLocation field if non-nil, zero value otherwise.

### GetDefaultLocationOk

`func (o *MyAccount) GetDefaultLocationOk() (*SystemLocationReference, bool)`

GetDefaultLocationOk returns a tuple with the DefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocation

`func (o *MyAccount) SetDefaultLocation(v SystemLocationReference)`

SetDefaultLocation sets DefaultLocation field to given value.


### GetDefaultDepartment

`func (o *MyAccount) GetDefaultDepartment() SystemDepartmentReference`

GetDefaultDepartment returns the DefaultDepartment field if non-nil, zero value otherwise.

### GetDefaultDepartmentOk

`func (o *MyAccount) GetDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetDefaultDepartmentOk returns a tuple with the DefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDepartment

`func (o *MyAccount) SetDefaultDepartment(v SystemDepartmentReference)`

SetDefaultDepartment sets DefaultDepartment field to given value.


### GetReportsTo

`func (o *MyAccount) GetReportsTo() MemberReference`

GetReportsTo returns the ReportsTo field if non-nil, zero value otherwise.

### GetReportsToOk

`func (o *MyAccount) GetReportsToOk() (*MemberReference, bool)`

GetReportsToOk returns a tuple with the ReportsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsTo

`func (o *MyAccount) SetReportsTo(v MemberReference)`

SetReportsTo sets ReportsTo field to given value.

### HasReportsTo

`func (o *MyAccount) HasReportsTo() bool`

HasReportsTo returns a boolean if a field has been set.

### GetWorkRole

`func (o *MyAccount) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *MyAccount) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *MyAccount) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.


### GetWorkType

`func (o *MyAccount) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *MyAccount) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *MyAccount) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *MyAccount) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetTimeApprover

`func (o *MyAccount) GetTimeApprover() MemberReference`

GetTimeApprover returns the TimeApprover field if non-nil, zero value otherwise.

### GetTimeApproverOk

`func (o *MyAccount) GetTimeApproverOk() (*MemberReference, bool)`

GetTimeApproverOk returns a tuple with the TimeApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeApprover

`func (o *MyAccount) SetTimeApprover(v MemberReference)`

SetTimeApprover sets TimeApprover field to given value.


### GetExpenseApprover

`func (o *MyAccount) GetExpenseApprover() MemberReference`

GetExpenseApprover returns the ExpenseApprover field if non-nil, zero value otherwise.

### GetExpenseApproverOk

`func (o *MyAccount) GetExpenseApproverOk() (*MemberReference, bool)`

GetExpenseApproverOk returns a tuple with the ExpenseApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApprover

`func (o *MyAccount) SetExpenseApprover(v MemberReference)`

SetExpenseApprover sets ExpenseApprover field to given value.


### GetBillableForecast

`func (o *MyAccount) GetBillableForecast() float64`

GetBillableForecast returns the BillableForecast field if non-nil, zero value otherwise.

### GetBillableForecastOk

`func (o *MyAccount) GetBillableForecastOk() (*float64, bool)`

GetBillableForecastOk returns a tuple with the BillableForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableForecast

`func (o *MyAccount) SetBillableForecast(v float64)`

SetBillableForecast sets BillableForecast field to given value.

### HasBillableForecast

`func (o *MyAccount) HasBillableForecast() bool`

HasBillableForecast returns a boolean if a field has been set.

### SetBillableForecastNil

`func (o *MyAccount) SetBillableForecastNil(b bool)`

 SetBillableForecastNil sets the value for BillableForecast to be an explicit nil

### UnsetBillableForecast
`func (o *MyAccount) UnsetBillableForecast()`

UnsetBillableForecast ensures that no value is present for BillableForecast, not even an explicit nil
### GetDailyCapacity

`func (o *MyAccount) GetDailyCapacity() float64`

GetDailyCapacity returns the DailyCapacity field if non-nil, zero value otherwise.

### GetDailyCapacityOk

`func (o *MyAccount) GetDailyCapacityOk() (*float64, bool)`

GetDailyCapacityOk returns a tuple with the DailyCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyCapacity

`func (o *MyAccount) SetDailyCapacity(v float64)`

SetDailyCapacity sets DailyCapacity field to given value.

### HasDailyCapacity

`func (o *MyAccount) HasDailyCapacity() bool`

HasDailyCapacity returns a boolean if a field has been set.

### SetDailyCapacityNil

`func (o *MyAccount) SetDailyCapacityNil(b bool)`

 SetDailyCapacityNil sets the value for DailyCapacity to be an explicit nil

### UnsetDailyCapacity
`func (o *MyAccount) UnsetDailyCapacity()`

UnsetDailyCapacity ensures that no value is present for DailyCapacity, not even an explicit nil
### GetIncludeInUtilizationReportingFlag

`func (o *MyAccount) GetIncludeInUtilizationReportingFlag() bool`

GetIncludeInUtilizationReportingFlag returns the IncludeInUtilizationReportingFlag field if non-nil, zero value otherwise.

### GetIncludeInUtilizationReportingFlagOk

`func (o *MyAccount) GetIncludeInUtilizationReportingFlagOk() (*bool, bool)`

GetIncludeInUtilizationReportingFlagOk returns a tuple with the IncludeInUtilizationReportingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInUtilizationReportingFlag

`func (o *MyAccount) SetIncludeInUtilizationReportingFlag(v bool)`

SetIncludeInUtilizationReportingFlag sets IncludeInUtilizationReportingFlag field to given value.

### HasIncludeInUtilizationReportingFlag

`func (o *MyAccount) HasIncludeInUtilizationReportingFlag() bool`

HasIncludeInUtilizationReportingFlag returns a boolean if a field has been set.

### SetIncludeInUtilizationReportingFlagNil

`func (o *MyAccount) SetIncludeInUtilizationReportingFlagNil(b bool)`

 SetIncludeInUtilizationReportingFlagNil sets the value for IncludeInUtilizationReportingFlag to be an explicit nil

### UnsetIncludeInUtilizationReportingFlag
`func (o *MyAccount) UnsetIncludeInUtilizationReportingFlag()`

UnsetIncludeInUtilizationReportingFlag ensures that no value is present for IncludeInUtilizationReportingFlag, not even an explicit nil
### GetRequireExpenseEntryFlag

`func (o *MyAccount) GetRequireExpenseEntryFlag() bool`

GetRequireExpenseEntryFlag returns the RequireExpenseEntryFlag field if non-nil, zero value otherwise.

### GetRequireExpenseEntryFlagOk

`func (o *MyAccount) GetRequireExpenseEntryFlagOk() (*bool, bool)`

GetRequireExpenseEntryFlagOk returns a tuple with the RequireExpenseEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExpenseEntryFlag

`func (o *MyAccount) SetRequireExpenseEntryFlag(v bool)`

SetRequireExpenseEntryFlag sets RequireExpenseEntryFlag field to given value.

### HasRequireExpenseEntryFlag

`func (o *MyAccount) HasRequireExpenseEntryFlag() bool`

HasRequireExpenseEntryFlag returns a boolean if a field has been set.

### SetRequireExpenseEntryFlagNil

`func (o *MyAccount) SetRequireExpenseEntryFlagNil(b bool)`

 SetRequireExpenseEntryFlagNil sets the value for RequireExpenseEntryFlag to be an explicit nil

### UnsetRequireExpenseEntryFlag
`func (o *MyAccount) UnsetRequireExpenseEntryFlag()`

UnsetRequireExpenseEntryFlag ensures that no value is present for RequireExpenseEntryFlag, not even an explicit nil
### GetRequireTimeSheetEntryFlag

`func (o *MyAccount) GetRequireTimeSheetEntryFlag() bool`

GetRequireTimeSheetEntryFlag returns the RequireTimeSheetEntryFlag field if non-nil, zero value otherwise.

### GetRequireTimeSheetEntryFlagOk

`func (o *MyAccount) GetRequireTimeSheetEntryFlagOk() (*bool, bool)`

GetRequireTimeSheetEntryFlagOk returns a tuple with the RequireTimeSheetEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTimeSheetEntryFlag

`func (o *MyAccount) SetRequireTimeSheetEntryFlag(v bool)`

SetRequireTimeSheetEntryFlag sets RequireTimeSheetEntryFlag field to given value.

### HasRequireTimeSheetEntryFlag

`func (o *MyAccount) HasRequireTimeSheetEntryFlag() bool`

HasRequireTimeSheetEntryFlag returns a boolean if a field has been set.

### SetRequireTimeSheetEntryFlagNil

`func (o *MyAccount) SetRequireTimeSheetEntryFlagNil(b bool)`

 SetRequireTimeSheetEntryFlagNil sets the value for RequireTimeSheetEntryFlag to be an explicit nil

### UnsetRequireTimeSheetEntryFlag
`func (o *MyAccount) UnsetRequireTimeSheetEntryFlag()`

UnsetRequireTimeSheetEntryFlag ensures that no value is present for RequireTimeSheetEntryFlag, not even an explicit nil
### GetRequireStartAndEndTimeOnTimeEntryFlag

`func (o *MyAccount) GetRequireStartAndEndTimeOnTimeEntryFlag() bool`

GetRequireStartAndEndTimeOnTimeEntryFlag returns the RequireStartAndEndTimeOnTimeEntryFlag field if non-nil, zero value otherwise.

### GetRequireStartAndEndTimeOnTimeEntryFlagOk

`func (o *MyAccount) GetRequireStartAndEndTimeOnTimeEntryFlagOk() (*bool, bool)`

GetRequireStartAndEndTimeOnTimeEntryFlagOk returns a tuple with the RequireStartAndEndTimeOnTimeEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireStartAndEndTimeOnTimeEntryFlag

`func (o *MyAccount) SetRequireStartAndEndTimeOnTimeEntryFlag(v bool)`

SetRequireStartAndEndTimeOnTimeEntryFlag sets RequireStartAndEndTimeOnTimeEntryFlag field to given value.

### HasRequireStartAndEndTimeOnTimeEntryFlag

`func (o *MyAccount) HasRequireStartAndEndTimeOnTimeEntryFlag() bool`

HasRequireStartAndEndTimeOnTimeEntryFlag returns a boolean if a field has been set.

### SetRequireStartAndEndTimeOnTimeEntryFlagNil

`func (o *MyAccount) SetRequireStartAndEndTimeOnTimeEntryFlagNil(b bool)`

 SetRequireStartAndEndTimeOnTimeEntryFlagNil sets the value for RequireStartAndEndTimeOnTimeEntryFlag to be an explicit nil

### UnsetRequireStartAndEndTimeOnTimeEntryFlag
`func (o *MyAccount) UnsetRequireStartAndEndTimeOnTimeEntryFlag()`

UnsetRequireStartAndEndTimeOnTimeEntryFlag ensures that no value is present for RequireStartAndEndTimeOnTimeEntryFlag, not even an explicit nil
### GetAllowInCellEntryOnTimeSheet

`func (o *MyAccount) GetAllowInCellEntryOnTimeSheet() bool`

GetAllowInCellEntryOnTimeSheet returns the AllowInCellEntryOnTimeSheet field if non-nil, zero value otherwise.

### GetAllowInCellEntryOnTimeSheetOk

`func (o *MyAccount) GetAllowInCellEntryOnTimeSheetOk() (*bool, bool)`

GetAllowInCellEntryOnTimeSheetOk returns a tuple with the AllowInCellEntryOnTimeSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInCellEntryOnTimeSheet

`func (o *MyAccount) SetAllowInCellEntryOnTimeSheet(v bool)`

SetAllowInCellEntryOnTimeSheet sets AllowInCellEntryOnTimeSheet field to given value.

### HasAllowInCellEntryOnTimeSheet

`func (o *MyAccount) HasAllowInCellEntryOnTimeSheet() bool`

HasAllowInCellEntryOnTimeSheet returns a boolean if a field has been set.

### SetAllowInCellEntryOnTimeSheetNil

`func (o *MyAccount) SetAllowInCellEntryOnTimeSheetNil(b bool)`

 SetAllowInCellEntryOnTimeSheetNil sets the value for AllowInCellEntryOnTimeSheet to be an explicit nil

### UnsetAllowInCellEntryOnTimeSheet
`func (o *MyAccount) UnsetAllowInCellEntryOnTimeSheet()`

UnsetAllowInCellEntryOnTimeSheet ensures that no value is present for AllowInCellEntryOnTimeSheet, not even an explicit nil
### GetEnterTimeAgainstCompanyFlag

`func (o *MyAccount) GetEnterTimeAgainstCompanyFlag() bool`

GetEnterTimeAgainstCompanyFlag returns the EnterTimeAgainstCompanyFlag field if non-nil, zero value otherwise.

### GetEnterTimeAgainstCompanyFlagOk

`func (o *MyAccount) GetEnterTimeAgainstCompanyFlagOk() (*bool, bool)`

GetEnterTimeAgainstCompanyFlagOk returns a tuple with the EnterTimeAgainstCompanyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterTimeAgainstCompanyFlag

`func (o *MyAccount) SetEnterTimeAgainstCompanyFlag(v bool)`

SetEnterTimeAgainstCompanyFlag sets EnterTimeAgainstCompanyFlag field to given value.

### HasEnterTimeAgainstCompanyFlag

`func (o *MyAccount) HasEnterTimeAgainstCompanyFlag() bool`

HasEnterTimeAgainstCompanyFlag returns a boolean if a field has been set.

### SetEnterTimeAgainstCompanyFlagNil

`func (o *MyAccount) SetEnterTimeAgainstCompanyFlagNil(b bool)`

 SetEnterTimeAgainstCompanyFlagNil sets the value for EnterTimeAgainstCompanyFlag to be an explicit nil

### UnsetEnterTimeAgainstCompanyFlag
`func (o *MyAccount) UnsetEnterTimeAgainstCompanyFlag()`

UnsetEnterTimeAgainstCompanyFlag ensures that no value is present for EnterTimeAgainstCompanyFlag, not even an explicit nil
### GetAllowExpensesEnteredAgainstCompaniesFlag

`func (o *MyAccount) GetAllowExpensesEnteredAgainstCompaniesFlag() bool`

GetAllowExpensesEnteredAgainstCompaniesFlag returns the AllowExpensesEnteredAgainstCompaniesFlag field if non-nil, zero value otherwise.

### GetAllowExpensesEnteredAgainstCompaniesFlagOk

`func (o *MyAccount) GetAllowExpensesEnteredAgainstCompaniesFlagOk() (*bool, bool)`

GetAllowExpensesEnteredAgainstCompaniesFlagOk returns a tuple with the AllowExpensesEnteredAgainstCompaniesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExpensesEnteredAgainstCompaniesFlag

`func (o *MyAccount) SetAllowExpensesEnteredAgainstCompaniesFlag(v bool)`

SetAllowExpensesEnteredAgainstCompaniesFlag sets AllowExpensesEnteredAgainstCompaniesFlag field to given value.

### HasAllowExpensesEnteredAgainstCompaniesFlag

`func (o *MyAccount) HasAllowExpensesEnteredAgainstCompaniesFlag() bool`

HasAllowExpensesEnteredAgainstCompaniesFlag returns a boolean if a field has been set.

### SetAllowExpensesEnteredAgainstCompaniesFlagNil

`func (o *MyAccount) SetAllowExpensesEnteredAgainstCompaniesFlagNil(b bool)`

 SetAllowExpensesEnteredAgainstCompaniesFlagNil sets the value for AllowExpensesEnteredAgainstCompaniesFlag to be an explicit nil

### UnsetAllowExpensesEnteredAgainstCompaniesFlag
`func (o *MyAccount) UnsetAllowExpensesEnteredAgainstCompaniesFlag()`

UnsetAllowExpensesEnteredAgainstCompaniesFlag ensures that no value is present for AllowExpensesEnteredAgainstCompaniesFlag, not even an explicit nil
### GetTimeReminderEmailFlag

`func (o *MyAccount) GetTimeReminderEmailFlag() bool`

GetTimeReminderEmailFlag returns the TimeReminderEmailFlag field if non-nil, zero value otherwise.

### GetTimeReminderEmailFlagOk

`func (o *MyAccount) GetTimeReminderEmailFlagOk() (*bool, bool)`

GetTimeReminderEmailFlagOk returns a tuple with the TimeReminderEmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeReminderEmailFlag

`func (o *MyAccount) SetTimeReminderEmailFlag(v bool)`

SetTimeReminderEmailFlag sets TimeReminderEmailFlag field to given value.

### HasTimeReminderEmailFlag

`func (o *MyAccount) HasTimeReminderEmailFlag() bool`

HasTimeReminderEmailFlag returns a boolean if a field has been set.

### SetTimeReminderEmailFlagNil

`func (o *MyAccount) SetTimeReminderEmailFlagNil(b bool)`

 SetTimeReminderEmailFlagNil sets the value for TimeReminderEmailFlag to be an explicit nil

### UnsetTimeReminderEmailFlag
`func (o *MyAccount) UnsetTimeReminderEmailFlag()`

UnsetTimeReminderEmailFlag ensures that no value is present for TimeReminderEmailFlag, not even an explicit nil
### GetDaysTolerance

`func (o *MyAccount) GetDaysTolerance() int32`

GetDaysTolerance returns the DaysTolerance field if non-nil, zero value otherwise.

### GetDaysToleranceOk

`func (o *MyAccount) GetDaysToleranceOk() (*int32, bool)`

GetDaysToleranceOk returns a tuple with the DaysTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysTolerance

`func (o *MyAccount) SetDaysTolerance(v int32)`

SetDaysTolerance sets DaysTolerance field to given value.

### HasDaysTolerance

`func (o *MyAccount) HasDaysTolerance() bool`

HasDaysTolerance returns a boolean if a field has been set.

### SetDaysToleranceNil

`func (o *MyAccount) SetDaysToleranceNil(b bool)`

 SetDaysToleranceNil sets the value for DaysTolerance to be an explicit nil

### UnsetDaysTolerance
`func (o *MyAccount) UnsetDaysTolerance()`

UnsetDaysTolerance ensures that no value is present for DaysTolerance, not even an explicit nil
### GetMinimumHours

`func (o *MyAccount) GetMinimumHours() float64`

GetMinimumHours returns the MinimumHours field if non-nil, zero value otherwise.

### GetMinimumHoursOk

`func (o *MyAccount) GetMinimumHoursOk() (*float64, bool)`

GetMinimumHoursOk returns a tuple with the MinimumHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumHours

`func (o *MyAccount) SetMinimumHours(v float64)`

SetMinimumHours sets MinimumHours field to given value.

### HasMinimumHours

`func (o *MyAccount) HasMinimumHours() bool`

HasMinimumHours returns a boolean if a field has been set.

### SetMinimumHoursNil

`func (o *MyAccount) SetMinimumHoursNil(b bool)`

 SetMinimumHoursNil sets the value for MinimumHours to be an explicit nil

### UnsetMinimumHours
`func (o *MyAccount) UnsetMinimumHours()`

UnsetMinimumHours ensures that no value is present for MinimumHours, not even an explicit nil
### GetTimeSheetStartDate

`func (o *MyAccount) GetTimeSheetStartDate() time.Time`

GetTimeSheetStartDate returns the TimeSheetStartDate field if non-nil, zero value otherwise.

### GetTimeSheetStartDateOk

`func (o *MyAccount) GetTimeSheetStartDateOk() (*time.Time, bool)`

GetTimeSheetStartDateOk returns a tuple with the TimeSheetStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSheetStartDate

`func (o *MyAccount) SetTimeSheetStartDate(v time.Time)`

SetTimeSheetStartDate sets TimeSheetStartDate field to given value.

### HasTimeSheetStartDate

`func (o *MyAccount) HasTimeSheetStartDate() bool`

HasTimeSheetStartDate returns a boolean if a field has been set.

### GetHireDate

`func (o *MyAccount) GetHireDate() time.Time`

GetHireDate returns the HireDate field if non-nil, zero value otherwise.

### GetHireDateOk

`func (o *MyAccount) GetHireDateOk() (*time.Time, bool)`

GetHireDateOk returns a tuple with the HireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireDate

`func (o *MyAccount) SetHireDate(v time.Time)`

SetHireDate sets HireDate field to given value.


### GetServiceDefaultLocation

`func (o *MyAccount) GetServiceDefaultLocation() SystemLocationReference`

GetServiceDefaultLocation returns the ServiceDefaultLocation field if non-nil, zero value otherwise.

### GetServiceDefaultLocationOk

`func (o *MyAccount) GetServiceDefaultLocationOk() (*SystemLocationReference, bool)`

GetServiceDefaultLocationOk returns a tuple with the ServiceDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultLocation

`func (o *MyAccount) SetServiceDefaultLocation(v SystemLocationReference)`

SetServiceDefaultLocation sets ServiceDefaultLocation field to given value.

### HasServiceDefaultLocation

`func (o *MyAccount) HasServiceDefaultLocation() bool`

HasServiceDefaultLocation returns a boolean if a field has been set.

### GetServiceDefaultDepartment

`func (o *MyAccount) GetServiceDefaultDepartment() SystemDepartmentReference`

GetServiceDefaultDepartment returns the ServiceDefaultDepartment field if non-nil, zero value otherwise.

### GetServiceDefaultDepartmentOk

`func (o *MyAccount) GetServiceDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetServiceDefaultDepartmentOk returns a tuple with the ServiceDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultDepartment

`func (o *MyAccount) SetServiceDefaultDepartment(v SystemDepartmentReference)`

SetServiceDefaultDepartment sets ServiceDefaultDepartment field to given value.

### HasServiceDefaultDepartment

`func (o *MyAccount) HasServiceDefaultDepartment() bool`

HasServiceDefaultDepartment returns a boolean if a field has been set.

### GetServiceDefaultBoard

`func (o *MyAccount) GetServiceDefaultBoard() BoardReference`

GetServiceDefaultBoard returns the ServiceDefaultBoard field if non-nil, zero value otherwise.

### GetServiceDefaultBoardOk

`func (o *MyAccount) GetServiceDefaultBoardOk() (*BoardReference, bool)`

GetServiceDefaultBoardOk returns a tuple with the ServiceDefaultBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultBoard

`func (o *MyAccount) SetServiceDefaultBoard(v BoardReference)`

SetServiceDefaultBoard sets ServiceDefaultBoard field to given value.

### HasServiceDefaultBoard

`func (o *MyAccount) HasServiceDefaultBoard() bool`

HasServiceDefaultBoard returns a boolean if a field has been set.

### GetProjectDefaultLocation

`func (o *MyAccount) GetProjectDefaultLocation() SystemLocationReference`

GetProjectDefaultLocation returns the ProjectDefaultLocation field if non-nil, zero value otherwise.

### GetProjectDefaultLocationOk

`func (o *MyAccount) GetProjectDefaultLocationOk() (*SystemLocationReference, bool)`

GetProjectDefaultLocationOk returns a tuple with the ProjectDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultLocation

`func (o *MyAccount) SetProjectDefaultLocation(v SystemLocationReference)`

SetProjectDefaultLocation sets ProjectDefaultLocation field to given value.

### HasProjectDefaultLocation

`func (o *MyAccount) HasProjectDefaultLocation() bool`

HasProjectDefaultLocation returns a boolean if a field has been set.

### GetProjectDefaultDepartment

`func (o *MyAccount) GetProjectDefaultDepartment() SystemDepartmentReference`

GetProjectDefaultDepartment returns the ProjectDefaultDepartment field if non-nil, zero value otherwise.

### GetProjectDefaultDepartmentOk

`func (o *MyAccount) GetProjectDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetProjectDefaultDepartmentOk returns a tuple with the ProjectDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultDepartment

`func (o *MyAccount) SetProjectDefaultDepartment(v SystemDepartmentReference)`

SetProjectDefaultDepartment sets ProjectDefaultDepartment field to given value.

### HasProjectDefaultDepartment

`func (o *MyAccount) HasProjectDefaultDepartment() bool`

HasProjectDefaultDepartment returns a boolean if a field has been set.

### GetProjectDefaultBoard

`func (o *MyAccount) GetProjectDefaultBoard() ProjectBoardReference`

GetProjectDefaultBoard returns the ProjectDefaultBoard field if non-nil, zero value otherwise.

### GetProjectDefaultBoardOk

`func (o *MyAccount) GetProjectDefaultBoardOk() (*ProjectBoardReference, bool)`

GetProjectDefaultBoardOk returns a tuple with the ProjectDefaultBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultBoard

`func (o *MyAccount) SetProjectDefaultBoard(v ProjectBoardReference)`

SetProjectDefaultBoard sets ProjectDefaultBoard field to given value.

### HasProjectDefaultBoard

`func (o *MyAccount) HasProjectDefaultBoard() bool`

HasProjectDefaultBoard returns a boolean if a field has been set.

### GetScheduleDefaultLocation

`func (o *MyAccount) GetScheduleDefaultLocation() SystemLocationReference`

GetScheduleDefaultLocation returns the ScheduleDefaultLocation field if non-nil, zero value otherwise.

### GetScheduleDefaultLocationOk

`func (o *MyAccount) GetScheduleDefaultLocationOk() (*SystemLocationReference, bool)`

GetScheduleDefaultLocationOk returns a tuple with the ScheduleDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDefaultLocation

`func (o *MyAccount) SetScheduleDefaultLocation(v SystemLocationReference)`

SetScheduleDefaultLocation sets ScheduleDefaultLocation field to given value.

### HasScheduleDefaultLocation

`func (o *MyAccount) HasScheduleDefaultLocation() bool`

HasScheduleDefaultLocation returns a boolean if a field has been set.

### GetScheduleDefaultDepartment

`func (o *MyAccount) GetScheduleDefaultDepartment() SystemDepartmentReference`

GetScheduleDefaultDepartment returns the ScheduleDefaultDepartment field if non-nil, zero value otherwise.

### GetScheduleDefaultDepartmentOk

`func (o *MyAccount) GetScheduleDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetScheduleDefaultDepartmentOk returns a tuple with the ScheduleDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDefaultDepartment

`func (o *MyAccount) SetScheduleDefaultDepartment(v SystemDepartmentReference)`

SetScheduleDefaultDepartment sets ScheduleDefaultDepartment field to given value.

### HasScheduleDefaultDepartment

`func (o *MyAccount) HasScheduleDefaultDepartment() bool`

HasScheduleDefaultDepartment returns a boolean if a field has been set.

### GetScheduleCapacity

`func (o *MyAccount) GetScheduleCapacity() float64`

GetScheduleCapacity returns the ScheduleCapacity field if non-nil, zero value otherwise.

### GetScheduleCapacityOk

`func (o *MyAccount) GetScheduleCapacityOk() (*float64, bool)`

GetScheduleCapacityOk returns a tuple with the ScheduleCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCapacity

`func (o *MyAccount) SetScheduleCapacity(v float64)`

SetScheduleCapacity sets ScheduleCapacity field to given value.

### HasScheduleCapacity

`func (o *MyAccount) HasScheduleCapacity() bool`

HasScheduleCapacity returns a boolean if a field has been set.

### SetScheduleCapacityNil

`func (o *MyAccount) SetScheduleCapacityNil(b bool)`

 SetScheduleCapacityNil sets the value for ScheduleCapacity to be an explicit nil

### UnsetScheduleCapacity
`func (o *MyAccount) UnsetScheduleCapacity()`

UnsetScheduleCapacity ensures that no value is present for ScheduleCapacity, not even an explicit nil
### GetServiceLocation

`func (o *MyAccount) GetServiceLocation() ServiceLocationReference`

GetServiceLocation returns the ServiceLocation field if non-nil, zero value otherwise.

### GetServiceLocationOk

`func (o *MyAccount) GetServiceLocationOk() (*ServiceLocationReference, bool)`

GetServiceLocationOk returns a tuple with the ServiceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLocation

`func (o *MyAccount) SetServiceLocation(v ServiceLocationReference)`

SetServiceLocation sets ServiceLocation field to given value.

### HasServiceLocation

`func (o *MyAccount) HasServiceLocation() bool`

HasServiceLocation returns a boolean if a field has been set.

### GetHideMemberInDispatchPortalFlag

`func (o *MyAccount) GetHideMemberInDispatchPortalFlag() bool`

GetHideMemberInDispatchPortalFlag returns the HideMemberInDispatchPortalFlag field if non-nil, zero value otherwise.

### GetHideMemberInDispatchPortalFlagOk

`func (o *MyAccount) GetHideMemberInDispatchPortalFlagOk() (*bool, bool)`

GetHideMemberInDispatchPortalFlagOk returns a tuple with the HideMemberInDispatchPortalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideMemberInDispatchPortalFlag

`func (o *MyAccount) SetHideMemberInDispatchPortalFlag(v bool)`

SetHideMemberInDispatchPortalFlag sets HideMemberInDispatchPortalFlag field to given value.

### HasHideMemberInDispatchPortalFlag

`func (o *MyAccount) HasHideMemberInDispatchPortalFlag() bool`

HasHideMemberInDispatchPortalFlag returns a boolean if a field has been set.

### SetHideMemberInDispatchPortalFlagNil

`func (o *MyAccount) SetHideMemberInDispatchPortalFlagNil(b bool)`

 SetHideMemberInDispatchPortalFlagNil sets the value for HideMemberInDispatchPortalFlag to be an explicit nil

### UnsetHideMemberInDispatchPortalFlag
`func (o *MyAccount) UnsetHideMemberInDispatchPortalFlag()`

UnsetHideMemberInDispatchPortalFlag ensures that no value is present for HideMemberInDispatchPortalFlag, not even an explicit nil
### GetCalendar

`func (o *MyAccount) GetCalendar() CalendarReference`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *MyAccount) GetCalendarOk() (*CalendarReference, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *MyAccount) SetCalendar(v CalendarReference)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *MyAccount) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### GetSalesDefaultLocation

`func (o *MyAccount) GetSalesDefaultLocation() SystemLocationReference`

GetSalesDefaultLocation returns the SalesDefaultLocation field if non-nil, zero value otherwise.

### GetSalesDefaultLocationOk

`func (o *MyAccount) GetSalesDefaultLocationOk() (*SystemLocationReference, bool)`

GetSalesDefaultLocationOk returns a tuple with the SalesDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDefaultLocation

`func (o *MyAccount) SetSalesDefaultLocation(v SystemLocationReference)`

SetSalesDefaultLocation sets SalesDefaultLocation field to given value.


### GetWarehouse

`func (o *MyAccount) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *MyAccount) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *MyAccount) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *MyAccount) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetWarehouseBin

`func (o *MyAccount) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *MyAccount) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *MyAccount) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.

### HasWarehouseBin

`func (o *MyAccount) HasWarehouseBin() bool`

HasWarehouseBin returns a boolean if a field has been set.

### GetMapiName

`func (o *MyAccount) GetMapiName() string`

GetMapiName returns the MapiName field if non-nil, zero value otherwise.

### GetMapiNameOk

`func (o *MyAccount) GetMapiNameOk() (*string, bool)`

GetMapiNameOk returns a tuple with the MapiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapiName

`func (o *MyAccount) SetMapiName(v string)`

SetMapiName sets MapiName field to given value.

### HasMapiName

`func (o *MyAccount) HasMapiName() bool`

HasMapiName returns a boolean if a field has been set.

### GetCalendarSyncIntegrationFlag

`func (o *MyAccount) GetCalendarSyncIntegrationFlag() bool`

GetCalendarSyncIntegrationFlag returns the CalendarSyncIntegrationFlag field if non-nil, zero value otherwise.

### GetCalendarSyncIntegrationFlagOk

`func (o *MyAccount) GetCalendarSyncIntegrationFlagOk() (*bool, bool)`

GetCalendarSyncIntegrationFlagOk returns a tuple with the CalendarSyncIntegrationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarSyncIntegrationFlag

`func (o *MyAccount) SetCalendarSyncIntegrationFlag(v bool)`

SetCalendarSyncIntegrationFlag sets CalendarSyncIntegrationFlag field to given value.

### HasCalendarSyncIntegrationFlag

`func (o *MyAccount) HasCalendarSyncIntegrationFlag() bool`

HasCalendarSyncIntegrationFlag returns a boolean if a field has been set.

### SetCalendarSyncIntegrationFlagNil

`func (o *MyAccount) SetCalendarSyncIntegrationFlagNil(b bool)`

 SetCalendarSyncIntegrationFlagNil sets the value for CalendarSyncIntegrationFlag to be an explicit nil

### UnsetCalendarSyncIntegrationFlag
`func (o *MyAccount) UnsetCalendarSyncIntegrationFlag()`

UnsetCalendarSyncIntegrationFlag ensures that no value is present for CalendarSyncIntegrationFlag, not even an explicit nil
### GetCompanyActivityTabFormat

`func (o *MyAccount) GetCompanyActivityTabFormat() string`

GetCompanyActivityTabFormat returns the CompanyActivityTabFormat field if non-nil, zero value otherwise.

### GetCompanyActivityTabFormatOk

`func (o *MyAccount) GetCompanyActivityTabFormatOk() (*string, bool)`

GetCompanyActivityTabFormatOk returns a tuple with the CompanyActivityTabFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyActivityTabFormat

`func (o *MyAccount) SetCompanyActivityTabFormat(v string)`

SetCompanyActivityTabFormat sets CompanyActivityTabFormat field to given value.


### SetCompanyActivityTabFormatNil

`func (o *MyAccount) SetCompanyActivityTabFormatNil(b bool)`

 SetCompanyActivityTabFormatNil sets the value for CompanyActivityTabFormat to be an explicit nil

### UnsetCompanyActivityTabFormat
`func (o *MyAccount) UnsetCompanyActivityTabFormat()`

UnsetCompanyActivityTabFormat ensures that no value is present for CompanyActivityTabFormat, not even an explicit nil
### GetInvoiceTimeTabFormat

`func (o *MyAccount) GetInvoiceTimeTabFormat() string`

GetInvoiceTimeTabFormat returns the InvoiceTimeTabFormat field if non-nil, zero value otherwise.

### GetInvoiceTimeTabFormatOk

`func (o *MyAccount) GetInvoiceTimeTabFormatOk() (*string, bool)`

GetInvoiceTimeTabFormatOk returns a tuple with the InvoiceTimeTabFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTimeTabFormat

`func (o *MyAccount) SetInvoiceTimeTabFormat(v string)`

SetInvoiceTimeTabFormat sets InvoiceTimeTabFormat field to given value.


### SetInvoiceTimeTabFormatNil

`func (o *MyAccount) SetInvoiceTimeTabFormatNil(b bool)`

 SetInvoiceTimeTabFormatNil sets the value for InvoiceTimeTabFormat to be an explicit nil

### UnsetInvoiceTimeTabFormat
`func (o *MyAccount) UnsetInvoiceTimeTabFormat()`

UnsetInvoiceTimeTabFormat ensures that no value is present for InvoiceTimeTabFormat, not even an explicit nil
### GetInvoiceScreenDefaultTabFormat

`func (o *MyAccount) GetInvoiceScreenDefaultTabFormat() string`

GetInvoiceScreenDefaultTabFormat returns the InvoiceScreenDefaultTabFormat field if non-nil, zero value otherwise.

### GetInvoiceScreenDefaultTabFormatOk

`func (o *MyAccount) GetInvoiceScreenDefaultTabFormatOk() (*string, bool)`

GetInvoiceScreenDefaultTabFormatOk returns a tuple with the InvoiceScreenDefaultTabFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceScreenDefaultTabFormat

`func (o *MyAccount) SetInvoiceScreenDefaultTabFormat(v string)`

SetInvoiceScreenDefaultTabFormat sets InvoiceScreenDefaultTabFormat field to given value.


### SetInvoiceScreenDefaultTabFormatNil

`func (o *MyAccount) SetInvoiceScreenDefaultTabFormatNil(b bool)`

 SetInvoiceScreenDefaultTabFormatNil sets the value for InvoiceScreenDefaultTabFormat to be an explicit nil

### UnsetInvoiceScreenDefaultTabFormat
`func (o *MyAccount) UnsetInvoiceScreenDefaultTabFormat()`

UnsetInvoiceScreenDefaultTabFormat ensures that no value is present for InvoiceScreenDefaultTabFormat, not even an explicit nil
### GetInvoicingDisplayOptions

`func (o *MyAccount) GetInvoicingDisplayOptions() string`

GetInvoicingDisplayOptions returns the InvoicingDisplayOptions field if non-nil, zero value otherwise.

### GetInvoicingDisplayOptionsOk

`func (o *MyAccount) GetInvoicingDisplayOptionsOk() (*string, bool)`

GetInvoicingDisplayOptionsOk returns a tuple with the InvoicingDisplayOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingDisplayOptions

`func (o *MyAccount) SetInvoicingDisplayOptions(v string)`

SetInvoicingDisplayOptions sets InvoicingDisplayOptions field to given value.


### SetInvoicingDisplayOptionsNil

`func (o *MyAccount) SetInvoicingDisplayOptionsNil(b bool)`

 SetInvoicingDisplayOptionsNil sets the value for InvoicingDisplayOptions to be an explicit nil

### UnsetInvoicingDisplayOptions
`func (o *MyAccount) UnsetInvoicingDisplayOptions()`

UnsetInvoicingDisplayOptions ensures that no value is present for InvoicingDisplayOptions, not even an explicit nil
### GetAgreementInvoicingDisplayOptions

`func (o *MyAccount) GetAgreementInvoicingDisplayOptions() string`

GetAgreementInvoicingDisplayOptions returns the AgreementInvoicingDisplayOptions field if non-nil, zero value otherwise.

### GetAgreementInvoicingDisplayOptionsOk

`func (o *MyAccount) GetAgreementInvoicingDisplayOptionsOk() (*string, bool)`

GetAgreementInvoicingDisplayOptionsOk returns a tuple with the AgreementInvoicingDisplayOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementInvoicingDisplayOptions

`func (o *MyAccount) SetAgreementInvoicingDisplayOptions(v string)`

SetAgreementInvoicingDisplayOptions sets AgreementInvoicingDisplayOptions field to given value.


### SetAgreementInvoicingDisplayOptionsNil

`func (o *MyAccount) SetAgreementInvoicingDisplayOptionsNil(b bool)`

 SetAgreementInvoicingDisplayOptionsNil sets the value for AgreementInvoicingDisplayOptions to be an explicit nil

### UnsetAgreementInvoicingDisplayOptions
`func (o *MyAccount) UnsetAgreementInvoicingDisplayOptions()`

UnsetAgreementInvoicingDisplayOptions ensures that no value is present for AgreementInvoicingDisplayOptions, not even an explicit nil
### GetAuthenticationServiceType

`func (o *MyAccount) GetAuthenticationServiceType() string`

GetAuthenticationServiceType returns the AuthenticationServiceType field if non-nil, zero value otherwise.

### GetAuthenticationServiceTypeOk

`func (o *MyAccount) GetAuthenticationServiceTypeOk() (*string, bool)`

GetAuthenticationServiceTypeOk returns a tuple with the AuthenticationServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationServiceType

`func (o *MyAccount) SetAuthenticationServiceType(v string)`

SetAuthenticationServiceType sets AuthenticationServiceType field to given value.

### HasAuthenticationServiceType

`func (o *MyAccount) HasAuthenticationServiceType() bool`

HasAuthenticationServiceType returns a boolean if a field has been set.

### SetAuthenticationServiceTypeNil

`func (o *MyAccount) SetAuthenticationServiceTypeNil(b bool)`

 SetAuthenticationServiceTypeNil sets the value for AuthenticationServiceType to be an explicit nil

### UnsetAuthenticationServiceType
`func (o *MyAccount) UnsetAuthenticationServiceType()`

UnsetAuthenticationServiceType ensures that no value is present for AuthenticationServiceType, not even an explicit nil
### GetTimebasedOneTimePasswordActivated

`func (o *MyAccount) GetTimebasedOneTimePasswordActivated() bool`

GetTimebasedOneTimePasswordActivated returns the TimebasedOneTimePasswordActivated field if non-nil, zero value otherwise.

### GetTimebasedOneTimePasswordActivatedOk

`func (o *MyAccount) GetTimebasedOneTimePasswordActivatedOk() (*bool, bool)`

GetTimebasedOneTimePasswordActivatedOk returns a tuple with the TimebasedOneTimePasswordActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimebasedOneTimePasswordActivated

`func (o *MyAccount) SetTimebasedOneTimePasswordActivated(v bool)`

SetTimebasedOneTimePasswordActivated sets TimebasedOneTimePasswordActivated field to given value.

### HasTimebasedOneTimePasswordActivated

`func (o *MyAccount) HasTimebasedOneTimePasswordActivated() bool`

HasTimebasedOneTimePasswordActivated returns a boolean if a field has been set.

### SetTimebasedOneTimePasswordActivatedNil

`func (o *MyAccount) SetTimebasedOneTimePasswordActivatedNil(b bool)`

 SetTimebasedOneTimePasswordActivatedNil sets the value for TimebasedOneTimePasswordActivated to be an explicit nil

### UnsetTimebasedOneTimePasswordActivated
`func (o *MyAccount) UnsetTimebasedOneTimePasswordActivated()`

UnsetTimebasedOneTimePasswordActivated ensures that no value is present for TimebasedOneTimePasswordActivated, not even an explicit nil
### GetDirectionalSync

`func (o *MyAccount) GetDirectionalSync() DirectionalSyncReference`

GetDirectionalSync returns the DirectionalSync field if non-nil, zero value otherwise.

### GetDirectionalSyncOk

`func (o *MyAccount) GetDirectionalSyncOk() (*DirectionalSyncReference, bool)`

GetDirectionalSyncOk returns a tuple with the DirectionalSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionalSync

`func (o *MyAccount) SetDirectionalSync(v DirectionalSyncReference)`

SetDirectionalSync sets DirectionalSync field to given value.

### HasDirectionalSync

`func (o *MyAccount) HasDirectionalSync() bool`

HasDirectionalSync returns a boolean if a field has been set.

### GetAutoStartStopwatch

`func (o *MyAccount) GetAutoStartStopwatch() bool`

GetAutoStartStopwatch returns the AutoStartStopwatch field if non-nil, zero value otherwise.

### GetAutoStartStopwatchOk

`func (o *MyAccount) GetAutoStartStopwatchOk() (*bool, bool)`

GetAutoStartStopwatchOk returns a tuple with the AutoStartStopwatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStartStopwatch

`func (o *MyAccount) SetAutoStartStopwatch(v bool)`

SetAutoStartStopwatch sets AutoStartStopwatch field to given value.

### HasAutoStartStopwatch

`func (o *MyAccount) HasAutoStartStopwatch() bool`

HasAutoStartStopwatch returns a boolean if a field has been set.

### SetAutoStartStopwatchNil

`func (o *MyAccount) SetAutoStartStopwatchNil(b bool)`

 SetAutoStartStopwatchNil sets the value for AutoStartStopwatch to be an explicit nil

### UnsetAutoStartStopwatch
`func (o *MyAccount) UnsetAutoStartStopwatch()`

UnsetAutoStartStopwatch ensures that no value is present for AutoStartStopwatch, not even an explicit nil
### GetAutoPopupQuickNotesWithStopwatch

`func (o *MyAccount) GetAutoPopupQuickNotesWithStopwatch() bool`

GetAutoPopupQuickNotesWithStopwatch returns the AutoPopupQuickNotesWithStopwatch field if non-nil, zero value otherwise.

### GetAutoPopupQuickNotesWithStopwatchOk

`func (o *MyAccount) GetAutoPopupQuickNotesWithStopwatchOk() (*bool, bool)`

GetAutoPopupQuickNotesWithStopwatchOk returns a tuple with the AutoPopupQuickNotesWithStopwatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPopupQuickNotesWithStopwatch

`func (o *MyAccount) SetAutoPopupQuickNotesWithStopwatch(v bool)`

SetAutoPopupQuickNotesWithStopwatch sets AutoPopupQuickNotesWithStopwatch field to given value.

### HasAutoPopupQuickNotesWithStopwatch

`func (o *MyAccount) HasAutoPopupQuickNotesWithStopwatch() bool`

HasAutoPopupQuickNotesWithStopwatch returns a boolean if a field has been set.

### SetAutoPopupQuickNotesWithStopwatchNil

`func (o *MyAccount) SetAutoPopupQuickNotesWithStopwatchNil(b bool)`

 SetAutoPopupQuickNotesWithStopwatchNil sets the value for AutoPopupQuickNotesWithStopwatch to be an explicit nil

### UnsetAutoPopupQuickNotesWithStopwatch
`func (o *MyAccount) UnsetAutoPopupQuickNotesWithStopwatch()`

UnsetAutoPopupQuickNotesWithStopwatch ensures that no value is present for AutoPopupQuickNotesWithStopwatch, not even an explicit nil
### GetSignature

`func (o *MyAccount) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *MyAccount) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *MyAccount) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *MyAccount) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetGlobalSearchDefaultTicketFilter

`func (o *MyAccount) GetGlobalSearchDefaultTicketFilter() string`

GetGlobalSearchDefaultTicketFilter returns the GlobalSearchDefaultTicketFilter field if non-nil, zero value otherwise.

### GetGlobalSearchDefaultTicketFilterOk

`func (o *MyAccount) GetGlobalSearchDefaultTicketFilterOk() (*string, bool)`

GetGlobalSearchDefaultTicketFilterOk returns a tuple with the GlobalSearchDefaultTicketFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSearchDefaultTicketFilter

`func (o *MyAccount) SetGlobalSearchDefaultTicketFilter(v string)`

SetGlobalSearchDefaultTicketFilter sets GlobalSearchDefaultTicketFilter field to given value.

### HasGlobalSearchDefaultTicketFilter

`func (o *MyAccount) HasGlobalSearchDefaultTicketFilter() bool`

HasGlobalSearchDefaultTicketFilter returns a boolean if a field has been set.

### SetGlobalSearchDefaultTicketFilterNil

`func (o *MyAccount) SetGlobalSearchDefaultTicketFilterNil(b bool)`

 SetGlobalSearchDefaultTicketFilterNil sets the value for GlobalSearchDefaultTicketFilter to be an explicit nil

### UnsetGlobalSearchDefaultTicketFilter
`func (o *MyAccount) UnsetGlobalSearchDefaultTicketFilter()`

UnsetGlobalSearchDefaultTicketFilter ensures that no value is present for GlobalSearchDefaultTicketFilter, not even an explicit nil
### GetGlobalSearchDefaultSort

`func (o *MyAccount) GetGlobalSearchDefaultSort() string`

GetGlobalSearchDefaultSort returns the GlobalSearchDefaultSort field if non-nil, zero value otherwise.

### GetGlobalSearchDefaultSortOk

`func (o *MyAccount) GetGlobalSearchDefaultSortOk() (*string, bool)`

GetGlobalSearchDefaultSortOk returns a tuple with the GlobalSearchDefaultSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSearchDefaultSort

`func (o *MyAccount) SetGlobalSearchDefaultSort(v string)`

SetGlobalSearchDefaultSort sets GlobalSearchDefaultSort field to given value.

### HasGlobalSearchDefaultSort

`func (o *MyAccount) HasGlobalSearchDefaultSort() bool`

HasGlobalSearchDefaultSort returns a boolean if a field has been set.

### SetGlobalSearchDefaultSortNil

`func (o *MyAccount) SetGlobalSearchDefaultSortNil(b bool)`

 SetGlobalSearchDefaultSortNil sets the value for GlobalSearchDefaultSort to be an explicit nil

### UnsetGlobalSearchDefaultSort
`func (o *MyAccount) UnsetGlobalSearchDefaultSort()`

UnsetGlobalSearchDefaultSort ensures that no value is present for GlobalSearchDefaultSort, not even an explicit nil
### GetPhoneSource

`func (o *MyAccount) GetPhoneSource() string`

GetPhoneSource returns the PhoneSource field if non-nil, zero value otherwise.

### GetPhoneSourceOk

`func (o *MyAccount) GetPhoneSourceOk() (*string, bool)`

GetPhoneSourceOk returns a tuple with the PhoneSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneSource

`func (o *MyAccount) SetPhoneSource(v string)`

SetPhoneSource sets PhoneSource field to given value.

### HasPhoneSource

`func (o *MyAccount) HasPhoneSource() bool`

HasPhoneSource returns a boolean if a field has been set.

### GetPhoneIntegrationType

`func (o *MyAccount) GetPhoneIntegrationType() string`

GetPhoneIntegrationType returns the PhoneIntegrationType field if non-nil, zero value otherwise.

### GetPhoneIntegrationTypeOk

`func (o *MyAccount) GetPhoneIntegrationTypeOk() (*string, bool)`

GetPhoneIntegrationTypeOk returns a tuple with the PhoneIntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneIntegrationType

`func (o *MyAccount) SetPhoneIntegrationType(v string)`

SetPhoneIntegrationType sets PhoneIntegrationType field to given value.

### HasPhoneIntegrationType

`func (o *MyAccount) HasPhoneIntegrationType() bool`

HasPhoneIntegrationType returns a boolean if a field has been set.

### SetPhoneIntegrationTypeNil

`func (o *MyAccount) SetPhoneIntegrationTypeNil(b bool)`

 SetPhoneIntegrationTypeNil sets the value for PhoneIntegrationType to be an explicit nil

### UnsetPhoneIntegrationType
`func (o *MyAccount) UnsetPhoneIntegrationType()`

UnsetPhoneIntegrationType ensures that no value is present for PhoneIntegrationType, not even an explicit nil
### GetUseBrowserLanguageFlag

`func (o *MyAccount) GetUseBrowserLanguageFlag() bool`

GetUseBrowserLanguageFlag returns the UseBrowserLanguageFlag field if non-nil, zero value otherwise.

### GetUseBrowserLanguageFlagOk

`func (o *MyAccount) GetUseBrowserLanguageFlagOk() (*bool, bool)`

GetUseBrowserLanguageFlagOk returns a tuple with the UseBrowserLanguageFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBrowserLanguageFlag

`func (o *MyAccount) SetUseBrowserLanguageFlag(v bool)`

SetUseBrowserLanguageFlag sets UseBrowserLanguageFlag field to given value.

### HasUseBrowserLanguageFlag

`func (o *MyAccount) HasUseBrowserLanguageFlag() bool`

HasUseBrowserLanguageFlag returns a boolean if a field has been set.

### SetUseBrowserLanguageFlagNil

`func (o *MyAccount) SetUseBrowserLanguageFlagNil(b bool)`

 SetUseBrowserLanguageFlagNil sets the value for UseBrowserLanguageFlag to be an explicit nil

### UnsetUseBrowserLanguageFlag
`func (o *MyAccount) UnsetUseBrowserLanguageFlag()`

UnsetUseBrowserLanguageFlag ensures that no value is present for UseBrowserLanguageFlag, not even an explicit nil
### GetInfo

`func (o *MyAccount) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MyAccount) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MyAccount) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MyAccount) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCopyPodLayouts

`func (o *MyAccount) GetCopyPodLayouts() bool`

GetCopyPodLayouts returns the CopyPodLayouts field if non-nil, zero value otherwise.

### GetCopyPodLayoutsOk

`func (o *MyAccount) GetCopyPodLayoutsOk() (*bool, bool)`

GetCopyPodLayoutsOk returns a tuple with the CopyPodLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyPodLayouts

`func (o *MyAccount) SetCopyPodLayouts(v bool)`

SetCopyPodLayouts sets CopyPodLayouts field to given value.

### HasCopyPodLayouts

`func (o *MyAccount) HasCopyPodLayouts() bool`

HasCopyPodLayouts returns a boolean if a field has been set.

### GetCopySharedDefaultViews

`func (o *MyAccount) GetCopySharedDefaultViews() bool`

GetCopySharedDefaultViews returns the CopySharedDefaultViews field if non-nil, zero value otherwise.

### GetCopySharedDefaultViewsOk

`func (o *MyAccount) GetCopySharedDefaultViewsOk() (*bool, bool)`

GetCopySharedDefaultViewsOk returns a tuple with the CopySharedDefaultViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySharedDefaultViews

`func (o *MyAccount) SetCopySharedDefaultViews(v bool)`

SetCopySharedDefaultViews sets CopySharedDefaultViews field to given value.

### HasCopySharedDefaultViews

`func (o *MyAccount) HasCopySharedDefaultViews() bool`

HasCopySharedDefaultViews returns a boolean if a field has been set.

### GetCopyColumnLayoutsAndFilters

`func (o *MyAccount) GetCopyColumnLayoutsAndFilters() bool`

GetCopyColumnLayoutsAndFilters returns the CopyColumnLayoutsAndFilters field if non-nil, zero value otherwise.

### GetCopyColumnLayoutsAndFiltersOk

`func (o *MyAccount) GetCopyColumnLayoutsAndFiltersOk() (*bool, bool)`

GetCopyColumnLayoutsAndFiltersOk returns a tuple with the CopyColumnLayoutsAndFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyColumnLayoutsAndFilters

`func (o *MyAccount) SetCopyColumnLayoutsAndFilters(v bool)`

SetCopyColumnLayoutsAndFilters sets CopyColumnLayoutsAndFilters field to given value.

### HasCopyColumnLayoutsAndFilters

`func (o *MyAccount) HasCopyColumnLayoutsAndFilters() bool`

HasCopyColumnLayoutsAndFilters returns a boolean if a field has been set.

### GetFromMemberRecId

`func (o *MyAccount) GetFromMemberRecId() int32`

GetFromMemberRecId returns the FromMemberRecId field if non-nil, zero value otherwise.

### GetFromMemberRecIdOk

`func (o *MyAccount) GetFromMemberRecIdOk() (*int32, bool)`

GetFromMemberRecIdOk returns a tuple with the FromMemberRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromMemberRecId

`func (o *MyAccount) SetFromMemberRecId(v int32)`

SetFromMemberRecId sets FromMemberRecId field to given value.

### HasFromMemberRecId

`func (o *MyAccount) HasFromMemberRecId() bool`

HasFromMemberRecId returns a boolean if a field has been set.

### GetCustomFields

`func (o *MyAccount) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MyAccount) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MyAccount) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MyAccount) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


