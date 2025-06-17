# MyMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** | ConditionallyRequired. API Member will get random password generated | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**MiddleInitial** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ReportCard** | Pointer to [**ReportCardReference**](ReportCardReference.md) |  | [optional] 
**LicenseClass** | Pointer to **NullableString** | F &#x3D; Full Member, A &#x3D; API Member, C &#x3D; StreamlineIT Member, X &#x3D; Subcontractor Member | [optional] 
**DisableOnlineFlag** | Pointer to **NullableBool** |  | [optional] 
**EnableMobileFlag** | Pointer to **NullableBool** |  | [optional] 
**Type** | Pointer to [**MemberTypeReference**](MemberTypeReference.md) |  | [optional] 
**EmployeeIdentifer** | Pointer to **string** |  | [optional] 
**VendorNumber** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to [**TimeZoneSetupReference**](TimeZoneSetupReference.md) |  | [optional] 
**Country** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**ServiceBoardTeamIds** | Pointer to **[]int32** |  | [optional] 
**EnableMobileGpsFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveDate** | Pointer to **string** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**LastLogin** | Pointer to **string** |  | [optional] 
**Photo** | Pointer to [**DocumentReference**](DocumentReference.md) |  | [optional] 
**ToastNotificationFlag** | Pointer to **NullableBool** |  | [optional] 
**OfficeEmail** | Pointer to **string** |  | [optional] 
**OfficePhone** | Pointer to **string** |  | [optional] 
**OfficeExtension** | Pointer to **string** |  | [optional] 
**MobileEmail** | Pointer to **string** |  | [optional] 
**MobilePhone** | Pointer to **string** |  | [optional] 
**MobileExtension** | Pointer to **string** |  | [optional] 
**HomeEmail** | Pointer to **string** |  | [optional] 
**HomePhone** | Pointer to **string** |  | [optional] 
**HomeExtension** | Pointer to **string** |  | [optional] 
**DefaultEmail** | Pointer to **NullableString** |  | [optional] 
**DefaultPhone** | Pointer to **NullableString** |  | [optional] 
**SecurityRole** | Pointer to [**SecurityRoleReference**](SecurityRoleReference.md) |  | [optional] 
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
**HireDate** | Pointer to **string** |  | [optional] 
**ServiceDefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**ServiceDefaultDepartment** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**ServiceDefaultBoard** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**RestrictServiceDefaultLocationFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictServiceDefaultDepartmentFlag** | Pointer to **NullableBool** |  | [optional] 
**ExcludedServiceBoardIds** | Pointer to **[]int32** |  | [optional] 
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
**MapiName** | Pointer to **string** |  | [optional] 
**CalendarSyncIntegrationFlag** | Pointer to **NullableBool** |  | [optional] 
**EnableLdapAuthenticationFlag** | Pointer to **NullableBool** |  | [optional] 
**LdapConfiguration** | Pointer to [**LdapConfigurationReference**](LdapConfigurationReference.md) |  | [optional] 
**LdapUserName** | Pointer to **string** |  | [optional] 
**CompanyActivityTabFormat** | Pointer to **NullableString** |  | [optional] 
**InvoiceTimeTabFormat** | Pointer to **NullableString** |  | [optional] 
**InvoiceScreenDefaultTabFormat** | Pointer to **NullableString** |  | [optional] 
**InvoicingDisplayOptions** | Pointer to **NullableString** |  | [optional] 
**AgreementInvoicingDisplayOptions** | Pointer to **NullableString** |  | [optional] 
**CorelyticsUsername** | Pointer to **string** |  | [optional] 
**CorelyticsPassword** | Pointer to **string** |  | [optional] 
**AuthenticationServiceType** | Pointer to **NullableString** |  | [optional] 
**TimebasedOneTimePasswordActivated** | Pointer to **NullableBool** |  | [optional] 
**DirectionalSync** | Pointer to [**DirectionalSyncReference**](DirectionalSyncReference.md) |  | [optional] 
**SsoSessionFlag** | Pointer to **NullableBool** |  | [optional] 
**SsoClientId** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMyMember

`func NewMyMember() *MyMember`

NewMyMember instantiates a new MyMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyMemberWithDefaults

`func NewMyMemberWithDefaults() *MyMember`

NewMyMemberWithDefaults instantiates a new MyMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MyMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyMember) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MyMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *MyMember) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MyMember) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MyMember) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MyMember) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetPassword

`func (o *MyMember) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MyMember) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MyMember) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MyMember) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetFirstName

`func (o *MyMember) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MyMember) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MyMember) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MyMember) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleInitial

`func (o *MyMember) GetMiddleInitial() string`

GetMiddleInitial returns the MiddleInitial field if non-nil, zero value otherwise.

### GetMiddleInitialOk

`func (o *MyMember) GetMiddleInitialOk() (*string, bool)`

GetMiddleInitialOk returns a tuple with the MiddleInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleInitial

`func (o *MyMember) SetMiddleInitial(v string)`

SetMiddleInitial sets MiddleInitial field to given value.

### HasMiddleInitial

`func (o *MyMember) HasMiddleInitial() bool`

HasMiddleInitial returns a boolean if a field has been set.

### GetLastName

`func (o *MyMember) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MyMember) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MyMember) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MyMember) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetTitle

`func (o *MyMember) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MyMember) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MyMember) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MyMember) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetReportCard

`func (o *MyMember) GetReportCard() ReportCardReference`

GetReportCard returns the ReportCard field if non-nil, zero value otherwise.

### GetReportCardOk

`func (o *MyMember) GetReportCardOk() (*ReportCardReference, bool)`

GetReportCardOk returns a tuple with the ReportCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCard

`func (o *MyMember) SetReportCard(v ReportCardReference)`

SetReportCard sets ReportCard field to given value.

### HasReportCard

`func (o *MyMember) HasReportCard() bool`

HasReportCard returns a boolean if a field has been set.

### GetLicenseClass

`func (o *MyMember) GetLicenseClass() string`

GetLicenseClass returns the LicenseClass field if non-nil, zero value otherwise.

### GetLicenseClassOk

`func (o *MyMember) GetLicenseClassOk() (*string, bool)`

GetLicenseClassOk returns a tuple with the LicenseClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseClass

`func (o *MyMember) SetLicenseClass(v string)`

SetLicenseClass sets LicenseClass field to given value.

### HasLicenseClass

`func (o *MyMember) HasLicenseClass() bool`

HasLicenseClass returns a boolean if a field has been set.

### SetLicenseClassNil

`func (o *MyMember) SetLicenseClassNil(b bool)`

 SetLicenseClassNil sets the value for LicenseClass to be an explicit nil

### UnsetLicenseClass
`func (o *MyMember) UnsetLicenseClass()`

UnsetLicenseClass ensures that no value is present for LicenseClass, not even an explicit nil
### GetDisableOnlineFlag

`func (o *MyMember) GetDisableOnlineFlag() bool`

GetDisableOnlineFlag returns the DisableOnlineFlag field if non-nil, zero value otherwise.

### GetDisableOnlineFlagOk

`func (o *MyMember) GetDisableOnlineFlagOk() (*bool, bool)`

GetDisableOnlineFlagOk returns a tuple with the DisableOnlineFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOnlineFlag

`func (o *MyMember) SetDisableOnlineFlag(v bool)`

SetDisableOnlineFlag sets DisableOnlineFlag field to given value.

### HasDisableOnlineFlag

`func (o *MyMember) HasDisableOnlineFlag() bool`

HasDisableOnlineFlag returns a boolean if a field has been set.

### SetDisableOnlineFlagNil

`func (o *MyMember) SetDisableOnlineFlagNil(b bool)`

 SetDisableOnlineFlagNil sets the value for DisableOnlineFlag to be an explicit nil

### UnsetDisableOnlineFlag
`func (o *MyMember) UnsetDisableOnlineFlag()`

UnsetDisableOnlineFlag ensures that no value is present for DisableOnlineFlag, not even an explicit nil
### GetEnableMobileFlag

`func (o *MyMember) GetEnableMobileFlag() bool`

GetEnableMobileFlag returns the EnableMobileFlag field if non-nil, zero value otherwise.

### GetEnableMobileFlagOk

`func (o *MyMember) GetEnableMobileFlagOk() (*bool, bool)`

GetEnableMobileFlagOk returns a tuple with the EnableMobileFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMobileFlag

`func (o *MyMember) SetEnableMobileFlag(v bool)`

SetEnableMobileFlag sets EnableMobileFlag field to given value.

### HasEnableMobileFlag

`func (o *MyMember) HasEnableMobileFlag() bool`

HasEnableMobileFlag returns a boolean if a field has been set.

### SetEnableMobileFlagNil

`func (o *MyMember) SetEnableMobileFlagNil(b bool)`

 SetEnableMobileFlagNil sets the value for EnableMobileFlag to be an explicit nil

### UnsetEnableMobileFlag
`func (o *MyMember) UnsetEnableMobileFlag()`

UnsetEnableMobileFlag ensures that no value is present for EnableMobileFlag, not even an explicit nil
### GetType

`func (o *MyMember) GetType() MemberTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MyMember) GetTypeOk() (*MemberTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MyMember) SetType(v MemberTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *MyMember) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmployeeIdentifer

`func (o *MyMember) GetEmployeeIdentifer() string`

GetEmployeeIdentifer returns the EmployeeIdentifer field if non-nil, zero value otherwise.

### GetEmployeeIdentiferOk

`func (o *MyMember) GetEmployeeIdentiferOk() (*string, bool)`

GetEmployeeIdentiferOk returns a tuple with the EmployeeIdentifer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeIdentifer

`func (o *MyMember) SetEmployeeIdentifer(v string)`

SetEmployeeIdentifer sets EmployeeIdentifer field to given value.

### HasEmployeeIdentifer

`func (o *MyMember) HasEmployeeIdentifer() bool`

HasEmployeeIdentifer returns a boolean if a field has been set.

### GetVendorNumber

`func (o *MyMember) GetVendorNumber() string`

GetVendorNumber returns the VendorNumber field if non-nil, zero value otherwise.

### GetVendorNumberOk

`func (o *MyMember) GetVendorNumberOk() (*string, bool)`

GetVendorNumberOk returns a tuple with the VendorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumber

`func (o *MyMember) SetVendorNumber(v string)`

SetVendorNumber sets VendorNumber field to given value.

### HasVendorNumber

`func (o *MyMember) HasVendorNumber() bool`

HasVendorNumber returns a boolean if a field has been set.

### GetNotes

`func (o *MyMember) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MyMember) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MyMember) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MyMember) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTimeZone

`func (o *MyMember) GetTimeZone() TimeZoneSetupReference`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MyMember) GetTimeZoneOk() (*TimeZoneSetupReference, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MyMember) SetTimeZone(v TimeZoneSetupReference)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MyMember) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetCountry

`func (o *MyMember) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MyMember) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MyMember) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *MyMember) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetServiceBoardTeamIds

`func (o *MyMember) GetServiceBoardTeamIds() []int32`

GetServiceBoardTeamIds returns the ServiceBoardTeamIds field if non-nil, zero value otherwise.

### GetServiceBoardTeamIdsOk

`func (o *MyMember) GetServiceBoardTeamIdsOk() (*[]int32, bool)`

GetServiceBoardTeamIdsOk returns a tuple with the ServiceBoardTeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBoardTeamIds

`func (o *MyMember) SetServiceBoardTeamIds(v []int32)`

SetServiceBoardTeamIds sets ServiceBoardTeamIds field to given value.

### HasServiceBoardTeamIds

`func (o *MyMember) HasServiceBoardTeamIds() bool`

HasServiceBoardTeamIds returns a boolean if a field has been set.

### GetEnableMobileGpsFlag

`func (o *MyMember) GetEnableMobileGpsFlag() bool`

GetEnableMobileGpsFlag returns the EnableMobileGpsFlag field if non-nil, zero value otherwise.

### GetEnableMobileGpsFlagOk

`func (o *MyMember) GetEnableMobileGpsFlagOk() (*bool, bool)`

GetEnableMobileGpsFlagOk returns a tuple with the EnableMobileGpsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMobileGpsFlag

`func (o *MyMember) SetEnableMobileGpsFlag(v bool)`

SetEnableMobileGpsFlag sets EnableMobileGpsFlag field to given value.

### HasEnableMobileGpsFlag

`func (o *MyMember) HasEnableMobileGpsFlag() bool`

HasEnableMobileGpsFlag returns a boolean if a field has been set.

### SetEnableMobileGpsFlagNil

`func (o *MyMember) SetEnableMobileGpsFlagNil(b bool)`

 SetEnableMobileGpsFlagNil sets the value for EnableMobileGpsFlag to be an explicit nil

### UnsetEnableMobileGpsFlag
`func (o *MyMember) UnsetEnableMobileGpsFlag()`

UnsetEnableMobileGpsFlag ensures that no value is present for EnableMobileGpsFlag, not even an explicit nil
### GetInactiveDate

`func (o *MyMember) GetInactiveDate() string`

GetInactiveDate returns the InactiveDate field if non-nil, zero value otherwise.

### GetInactiveDateOk

`func (o *MyMember) GetInactiveDateOk() (*string, bool)`

GetInactiveDateOk returns a tuple with the InactiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveDate

`func (o *MyMember) SetInactiveDate(v string)`

SetInactiveDate sets InactiveDate field to given value.

### HasInactiveDate

`func (o *MyMember) HasInactiveDate() bool`

HasInactiveDate returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *MyMember) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *MyMember) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *MyMember) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *MyMember) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *MyMember) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *MyMember) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetLastLogin

`func (o *MyMember) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *MyMember) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *MyMember) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *MyMember) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetPhoto

`func (o *MyMember) GetPhoto() DocumentReference`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MyMember) GetPhotoOk() (*DocumentReference, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *MyMember) SetPhoto(v DocumentReference)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *MyMember) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetToastNotificationFlag

`func (o *MyMember) GetToastNotificationFlag() bool`

GetToastNotificationFlag returns the ToastNotificationFlag field if non-nil, zero value otherwise.

### GetToastNotificationFlagOk

`func (o *MyMember) GetToastNotificationFlagOk() (*bool, bool)`

GetToastNotificationFlagOk returns a tuple with the ToastNotificationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToastNotificationFlag

`func (o *MyMember) SetToastNotificationFlag(v bool)`

SetToastNotificationFlag sets ToastNotificationFlag field to given value.

### HasToastNotificationFlag

`func (o *MyMember) HasToastNotificationFlag() bool`

HasToastNotificationFlag returns a boolean if a field has been set.

### SetToastNotificationFlagNil

`func (o *MyMember) SetToastNotificationFlagNil(b bool)`

 SetToastNotificationFlagNil sets the value for ToastNotificationFlag to be an explicit nil

### UnsetToastNotificationFlag
`func (o *MyMember) UnsetToastNotificationFlag()`

UnsetToastNotificationFlag ensures that no value is present for ToastNotificationFlag, not even an explicit nil
### GetOfficeEmail

`func (o *MyMember) GetOfficeEmail() string`

GetOfficeEmail returns the OfficeEmail field if non-nil, zero value otherwise.

### GetOfficeEmailOk

`func (o *MyMember) GetOfficeEmailOk() (*string, bool)`

GetOfficeEmailOk returns a tuple with the OfficeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeEmail

`func (o *MyMember) SetOfficeEmail(v string)`

SetOfficeEmail sets OfficeEmail field to given value.

### HasOfficeEmail

`func (o *MyMember) HasOfficeEmail() bool`

HasOfficeEmail returns a boolean if a field has been set.

### GetOfficePhone

`func (o *MyMember) GetOfficePhone() string`

GetOfficePhone returns the OfficePhone field if non-nil, zero value otherwise.

### GetOfficePhoneOk

`func (o *MyMember) GetOfficePhoneOk() (*string, bool)`

GetOfficePhoneOk returns a tuple with the OfficePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficePhone

`func (o *MyMember) SetOfficePhone(v string)`

SetOfficePhone sets OfficePhone field to given value.

### HasOfficePhone

`func (o *MyMember) HasOfficePhone() bool`

HasOfficePhone returns a boolean if a field has been set.

### GetOfficeExtension

`func (o *MyMember) GetOfficeExtension() string`

GetOfficeExtension returns the OfficeExtension field if non-nil, zero value otherwise.

### GetOfficeExtensionOk

`func (o *MyMember) GetOfficeExtensionOk() (*string, bool)`

GetOfficeExtensionOk returns a tuple with the OfficeExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeExtension

`func (o *MyMember) SetOfficeExtension(v string)`

SetOfficeExtension sets OfficeExtension field to given value.

### HasOfficeExtension

`func (o *MyMember) HasOfficeExtension() bool`

HasOfficeExtension returns a boolean if a field has been set.

### GetMobileEmail

`func (o *MyMember) GetMobileEmail() string`

GetMobileEmail returns the MobileEmail field if non-nil, zero value otherwise.

### GetMobileEmailOk

`func (o *MyMember) GetMobileEmailOk() (*string, bool)`

GetMobileEmailOk returns a tuple with the MobileEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileEmail

`func (o *MyMember) SetMobileEmail(v string)`

SetMobileEmail sets MobileEmail field to given value.

### HasMobileEmail

`func (o *MyMember) HasMobileEmail() bool`

HasMobileEmail returns a boolean if a field has been set.

### GetMobilePhone

`func (o *MyMember) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *MyMember) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *MyMember) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *MyMember) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetMobileExtension

`func (o *MyMember) GetMobileExtension() string`

GetMobileExtension returns the MobileExtension field if non-nil, zero value otherwise.

### GetMobileExtensionOk

`func (o *MyMember) GetMobileExtensionOk() (*string, bool)`

GetMobileExtensionOk returns a tuple with the MobileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileExtension

`func (o *MyMember) SetMobileExtension(v string)`

SetMobileExtension sets MobileExtension field to given value.

### HasMobileExtension

`func (o *MyMember) HasMobileExtension() bool`

HasMobileExtension returns a boolean if a field has been set.

### GetHomeEmail

`func (o *MyMember) GetHomeEmail() string`

GetHomeEmail returns the HomeEmail field if non-nil, zero value otherwise.

### GetHomeEmailOk

`func (o *MyMember) GetHomeEmailOk() (*string, bool)`

GetHomeEmailOk returns a tuple with the HomeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeEmail

`func (o *MyMember) SetHomeEmail(v string)`

SetHomeEmail sets HomeEmail field to given value.

### HasHomeEmail

`func (o *MyMember) HasHomeEmail() bool`

HasHomeEmail returns a boolean if a field has been set.

### GetHomePhone

`func (o *MyMember) GetHomePhone() string`

GetHomePhone returns the HomePhone field if non-nil, zero value otherwise.

### GetHomePhoneOk

`func (o *MyMember) GetHomePhoneOk() (*string, bool)`

GetHomePhoneOk returns a tuple with the HomePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePhone

`func (o *MyMember) SetHomePhone(v string)`

SetHomePhone sets HomePhone field to given value.

### HasHomePhone

`func (o *MyMember) HasHomePhone() bool`

HasHomePhone returns a boolean if a field has been set.

### GetHomeExtension

`func (o *MyMember) GetHomeExtension() string`

GetHomeExtension returns the HomeExtension field if non-nil, zero value otherwise.

### GetHomeExtensionOk

`func (o *MyMember) GetHomeExtensionOk() (*string, bool)`

GetHomeExtensionOk returns a tuple with the HomeExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeExtension

`func (o *MyMember) SetHomeExtension(v string)`

SetHomeExtension sets HomeExtension field to given value.

### HasHomeExtension

`func (o *MyMember) HasHomeExtension() bool`

HasHomeExtension returns a boolean if a field has been set.

### GetDefaultEmail

`func (o *MyMember) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *MyMember) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *MyMember) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.

### HasDefaultEmail

`func (o *MyMember) HasDefaultEmail() bool`

HasDefaultEmail returns a boolean if a field has been set.

### SetDefaultEmailNil

`func (o *MyMember) SetDefaultEmailNil(b bool)`

 SetDefaultEmailNil sets the value for DefaultEmail to be an explicit nil

### UnsetDefaultEmail
`func (o *MyMember) UnsetDefaultEmail()`

UnsetDefaultEmail ensures that no value is present for DefaultEmail, not even an explicit nil
### GetDefaultPhone

`func (o *MyMember) GetDefaultPhone() string`

GetDefaultPhone returns the DefaultPhone field if non-nil, zero value otherwise.

### GetDefaultPhoneOk

`func (o *MyMember) GetDefaultPhoneOk() (*string, bool)`

GetDefaultPhoneOk returns a tuple with the DefaultPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhone

`func (o *MyMember) SetDefaultPhone(v string)`

SetDefaultPhone sets DefaultPhone field to given value.

### HasDefaultPhone

`func (o *MyMember) HasDefaultPhone() bool`

HasDefaultPhone returns a boolean if a field has been set.

### SetDefaultPhoneNil

`func (o *MyMember) SetDefaultPhoneNil(b bool)`

 SetDefaultPhoneNil sets the value for DefaultPhone to be an explicit nil

### UnsetDefaultPhone
`func (o *MyMember) UnsetDefaultPhone()`

UnsetDefaultPhone ensures that no value is present for DefaultPhone, not even an explicit nil
### GetSecurityRole

`func (o *MyMember) GetSecurityRole() SecurityRoleReference`

GetSecurityRole returns the SecurityRole field if non-nil, zero value otherwise.

### GetSecurityRoleOk

`func (o *MyMember) GetSecurityRoleOk() (*SecurityRoleReference, bool)`

GetSecurityRoleOk returns a tuple with the SecurityRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityRole

`func (o *MyMember) SetSecurityRole(v SecurityRoleReference)`

SetSecurityRole sets SecurityRole field to given value.

### HasSecurityRole

`func (o *MyMember) HasSecurityRole() bool`

HasSecurityRole returns a boolean if a field has been set.

### GetAdminFlag

`func (o *MyMember) GetAdminFlag() bool`

GetAdminFlag returns the AdminFlag field if non-nil, zero value otherwise.

### GetAdminFlagOk

`func (o *MyMember) GetAdminFlagOk() (*bool, bool)`

GetAdminFlagOk returns a tuple with the AdminFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFlag

`func (o *MyMember) SetAdminFlag(v bool)`

SetAdminFlag sets AdminFlag field to given value.

### HasAdminFlag

`func (o *MyMember) HasAdminFlag() bool`

HasAdminFlag returns a boolean if a field has been set.

### SetAdminFlagNil

`func (o *MyMember) SetAdminFlagNil(b bool)`

 SetAdminFlagNil sets the value for AdminFlag to be an explicit nil

### UnsetAdminFlag
`func (o *MyMember) UnsetAdminFlag()`

UnsetAdminFlag ensures that no value is present for AdminFlag, not even an explicit nil
### GetStructureLevel

`func (o *MyMember) GetStructureLevel() StructureReference`

GetStructureLevel returns the StructureLevel field if non-nil, zero value otherwise.

### GetStructureLevelOk

`func (o *MyMember) GetStructureLevelOk() (*StructureReference, bool)`

GetStructureLevelOk returns a tuple with the StructureLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureLevel

`func (o *MyMember) SetStructureLevel(v StructureReference)`

SetStructureLevel sets StructureLevel field to given value.

### HasStructureLevel

`func (o *MyMember) HasStructureLevel() bool`

HasStructureLevel returns a boolean if a field has been set.

### GetSecurityLocation

`func (o *MyMember) GetSecurityLocation() SystemLocationReference`

GetSecurityLocation returns the SecurityLocation field if non-nil, zero value otherwise.

### GetSecurityLocationOk

`func (o *MyMember) GetSecurityLocationOk() (*SystemLocationReference, bool)`

GetSecurityLocationOk returns a tuple with the SecurityLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLocation

`func (o *MyMember) SetSecurityLocation(v SystemLocationReference)`

SetSecurityLocation sets SecurityLocation field to given value.

### HasSecurityLocation

`func (o *MyMember) HasSecurityLocation() bool`

HasSecurityLocation returns a boolean if a field has been set.

### GetDefaultLocation

`func (o *MyMember) GetDefaultLocation() SystemLocationReference`

GetDefaultLocation returns the DefaultLocation field if non-nil, zero value otherwise.

### GetDefaultLocationOk

`func (o *MyMember) GetDefaultLocationOk() (*SystemLocationReference, bool)`

GetDefaultLocationOk returns a tuple with the DefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocation

`func (o *MyMember) SetDefaultLocation(v SystemLocationReference)`

SetDefaultLocation sets DefaultLocation field to given value.

### HasDefaultLocation

`func (o *MyMember) HasDefaultLocation() bool`

HasDefaultLocation returns a boolean if a field has been set.

### GetDefaultDepartment

`func (o *MyMember) GetDefaultDepartment() SystemDepartmentReference`

GetDefaultDepartment returns the DefaultDepartment field if non-nil, zero value otherwise.

### GetDefaultDepartmentOk

`func (o *MyMember) GetDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetDefaultDepartmentOk returns a tuple with the DefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDepartment

`func (o *MyMember) SetDefaultDepartment(v SystemDepartmentReference)`

SetDefaultDepartment sets DefaultDepartment field to given value.

### HasDefaultDepartment

`func (o *MyMember) HasDefaultDepartment() bool`

HasDefaultDepartment returns a boolean if a field has been set.

### GetReportsTo

`func (o *MyMember) GetReportsTo() MemberReference`

GetReportsTo returns the ReportsTo field if non-nil, zero value otherwise.

### GetReportsToOk

`func (o *MyMember) GetReportsToOk() (*MemberReference, bool)`

GetReportsToOk returns a tuple with the ReportsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsTo

`func (o *MyMember) SetReportsTo(v MemberReference)`

SetReportsTo sets ReportsTo field to given value.

### HasReportsTo

`func (o *MyMember) HasReportsTo() bool`

HasReportsTo returns a boolean if a field has been set.

### GetRestrictLocationFlag

`func (o *MyMember) GetRestrictLocationFlag() bool`

GetRestrictLocationFlag returns the RestrictLocationFlag field if non-nil, zero value otherwise.

### GetRestrictLocationFlagOk

`func (o *MyMember) GetRestrictLocationFlagOk() (*bool, bool)`

GetRestrictLocationFlagOk returns a tuple with the RestrictLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictLocationFlag

`func (o *MyMember) SetRestrictLocationFlag(v bool)`

SetRestrictLocationFlag sets RestrictLocationFlag field to given value.

### HasRestrictLocationFlag

`func (o *MyMember) HasRestrictLocationFlag() bool`

HasRestrictLocationFlag returns a boolean if a field has been set.

### SetRestrictLocationFlagNil

`func (o *MyMember) SetRestrictLocationFlagNil(b bool)`

 SetRestrictLocationFlagNil sets the value for RestrictLocationFlag to be an explicit nil

### UnsetRestrictLocationFlag
`func (o *MyMember) UnsetRestrictLocationFlag()`

UnsetRestrictLocationFlag ensures that no value is present for RestrictLocationFlag, not even an explicit nil
### GetRestrictDepartmentFlag

`func (o *MyMember) GetRestrictDepartmentFlag() bool`

GetRestrictDepartmentFlag returns the RestrictDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictDepartmentFlagOk

`func (o *MyMember) GetRestrictDepartmentFlagOk() (*bool, bool)`

GetRestrictDepartmentFlagOk returns a tuple with the RestrictDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDepartmentFlag

`func (o *MyMember) SetRestrictDepartmentFlag(v bool)`

SetRestrictDepartmentFlag sets RestrictDepartmentFlag field to given value.

### HasRestrictDepartmentFlag

`func (o *MyMember) HasRestrictDepartmentFlag() bool`

HasRestrictDepartmentFlag returns a boolean if a field has been set.

### SetRestrictDepartmentFlagNil

`func (o *MyMember) SetRestrictDepartmentFlagNil(b bool)`

 SetRestrictDepartmentFlagNil sets the value for RestrictDepartmentFlag to be an explicit nil

### UnsetRestrictDepartmentFlag
`func (o *MyMember) UnsetRestrictDepartmentFlag()`

UnsetRestrictDepartmentFlag ensures that no value is present for RestrictDepartmentFlag, not even an explicit nil
### GetWorkRole

`func (o *MyMember) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *MyMember) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *MyMember) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *MyMember) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *MyMember) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *MyMember) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *MyMember) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *MyMember) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetTimeApprover

`func (o *MyMember) GetTimeApprover() MemberReference`

GetTimeApprover returns the TimeApprover field if non-nil, zero value otherwise.

### GetTimeApproverOk

`func (o *MyMember) GetTimeApproverOk() (*MemberReference, bool)`

GetTimeApproverOk returns a tuple with the TimeApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeApprover

`func (o *MyMember) SetTimeApprover(v MemberReference)`

SetTimeApprover sets TimeApprover field to given value.

### HasTimeApprover

`func (o *MyMember) HasTimeApprover() bool`

HasTimeApprover returns a boolean if a field has been set.

### GetExpenseApprover

`func (o *MyMember) GetExpenseApprover() MemberReference`

GetExpenseApprover returns the ExpenseApprover field if non-nil, zero value otherwise.

### GetExpenseApproverOk

`func (o *MyMember) GetExpenseApproverOk() (*MemberReference, bool)`

GetExpenseApproverOk returns a tuple with the ExpenseApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApprover

`func (o *MyMember) SetExpenseApprover(v MemberReference)`

SetExpenseApprover sets ExpenseApprover field to given value.

### HasExpenseApprover

`func (o *MyMember) HasExpenseApprover() bool`

HasExpenseApprover returns a boolean if a field has been set.

### GetBillableForecast

`func (o *MyMember) GetBillableForecast() float64`

GetBillableForecast returns the BillableForecast field if non-nil, zero value otherwise.

### GetBillableForecastOk

`func (o *MyMember) GetBillableForecastOk() (*float64, bool)`

GetBillableForecastOk returns a tuple with the BillableForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableForecast

`func (o *MyMember) SetBillableForecast(v float64)`

SetBillableForecast sets BillableForecast field to given value.

### HasBillableForecast

`func (o *MyMember) HasBillableForecast() bool`

HasBillableForecast returns a boolean if a field has been set.

### SetBillableForecastNil

`func (o *MyMember) SetBillableForecastNil(b bool)`

 SetBillableForecastNil sets the value for BillableForecast to be an explicit nil

### UnsetBillableForecast
`func (o *MyMember) UnsetBillableForecast()`

UnsetBillableForecast ensures that no value is present for BillableForecast, not even an explicit nil
### GetDailyCapacity

`func (o *MyMember) GetDailyCapacity() float64`

GetDailyCapacity returns the DailyCapacity field if non-nil, zero value otherwise.

### GetDailyCapacityOk

`func (o *MyMember) GetDailyCapacityOk() (*float64, bool)`

GetDailyCapacityOk returns a tuple with the DailyCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyCapacity

`func (o *MyMember) SetDailyCapacity(v float64)`

SetDailyCapacity sets DailyCapacity field to given value.

### HasDailyCapacity

`func (o *MyMember) HasDailyCapacity() bool`

HasDailyCapacity returns a boolean if a field has been set.

### SetDailyCapacityNil

`func (o *MyMember) SetDailyCapacityNil(b bool)`

 SetDailyCapacityNil sets the value for DailyCapacity to be an explicit nil

### UnsetDailyCapacity
`func (o *MyMember) UnsetDailyCapacity()`

UnsetDailyCapacity ensures that no value is present for DailyCapacity, not even an explicit nil
### GetHourlyCost

`func (o *MyMember) GetHourlyCost() float64`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *MyMember) GetHourlyCostOk() (*float64, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *MyMember) SetHourlyCost(v float64)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *MyMember) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### SetHourlyCostNil

`func (o *MyMember) SetHourlyCostNil(b bool)`

 SetHourlyCostNil sets the value for HourlyCost to be an explicit nil

### UnsetHourlyCost
`func (o *MyMember) UnsetHourlyCost()`

UnsetHourlyCost ensures that no value is present for HourlyCost, not even an explicit nil
### GetHourlyRate

`func (o *MyMember) GetHourlyRate() float64`

GetHourlyRate returns the HourlyRate field if non-nil, zero value otherwise.

### GetHourlyRateOk

`func (o *MyMember) GetHourlyRateOk() (*float64, bool)`

GetHourlyRateOk returns a tuple with the HourlyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyRate

`func (o *MyMember) SetHourlyRate(v float64)`

SetHourlyRate sets HourlyRate field to given value.

### HasHourlyRate

`func (o *MyMember) HasHourlyRate() bool`

HasHourlyRate returns a boolean if a field has been set.

### SetHourlyRateNil

`func (o *MyMember) SetHourlyRateNil(b bool)`

 SetHourlyRateNil sets the value for HourlyRate to be an explicit nil

### UnsetHourlyRate
`func (o *MyMember) UnsetHourlyRate()`

UnsetHourlyRate ensures that no value is present for HourlyRate, not even an explicit nil
### GetIncludeInUtilizationReportingFlag

`func (o *MyMember) GetIncludeInUtilizationReportingFlag() bool`

GetIncludeInUtilizationReportingFlag returns the IncludeInUtilizationReportingFlag field if non-nil, zero value otherwise.

### GetIncludeInUtilizationReportingFlagOk

`func (o *MyMember) GetIncludeInUtilizationReportingFlagOk() (*bool, bool)`

GetIncludeInUtilizationReportingFlagOk returns a tuple with the IncludeInUtilizationReportingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInUtilizationReportingFlag

`func (o *MyMember) SetIncludeInUtilizationReportingFlag(v bool)`

SetIncludeInUtilizationReportingFlag sets IncludeInUtilizationReportingFlag field to given value.

### HasIncludeInUtilizationReportingFlag

`func (o *MyMember) HasIncludeInUtilizationReportingFlag() bool`

HasIncludeInUtilizationReportingFlag returns a boolean if a field has been set.

### SetIncludeInUtilizationReportingFlagNil

`func (o *MyMember) SetIncludeInUtilizationReportingFlagNil(b bool)`

 SetIncludeInUtilizationReportingFlagNil sets the value for IncludeInUtilizationReportingFlag to be an explicit nil

### UnsetIncludeInUtilizationReportingFlag
`func (o *MyMember) UnsetIncludeInUtilizationReportingFlag()`

UnsetIncludeInUtilizationReportingFlag ensures that no value is present for IncludeInUtilizationReportingFlag, not even an explicit nil
### GetRequireExpenseEntryFlag

`func (o *MyMember) GetRequireExpenseEntryFlag() bool`

GetRequireExpenseEntryFlag returns the RequireExpenseEntryFlag field if non-nil, zero value otherwise.

### GetRequireExpenseEntryFlagOk

`func (o *MyMember) GetRequireExpenseEntryFlagOk() (*bool, bool)`

GetRequireExpenseEntryFlagOk returns a tuple with the RequireExpenseEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExpenseEntryFlag

`func (o *MyMember) SetRequireExpenseEntryFlag(v bool)`

SetRequireExpenseEntryFlag sets RequireExpenseEntryFlag field to given value.

### HasRequireExpenseEntryFlag

`func (o *MyMember) HasRequireExpenseEntryFlag() bool`

HasRequireExpenseEntryFlag returns a boolean if a field has been set.

### SetRequireExpenseEntryFlagNil

`func (o *MyMember) SetRequireExpenseEntryFlagNil(b bool)`

 SetRequireExpenseEntryFlagNil sets the value for RequireExpenseEntryFlag to be an explicit nil

### UnsetRequireExpenseEntryFlag
`func (o *MyMember) UnsetRequireExpenseEntryFlag()`

UnsetRequireExpenseEntryFlag ensures that no value is present for RequireExpenseEntryFlag, not even an explicit nil
### GetRequireTimeSheetEntryFlag

`func (o *MyMember) GetRequireTimeSheetEntryFlag() bool`

GetRequireTimeSheetEntryFlag returns the RequireTimeSheetEntryFlag field if non-nil, zero value otherwise.

### GetRequireTimeSheetEntryFlagOk

`func (o *MyMember) GetRequireTimeSheetEntryFlagOk() (*bool, bool)`

GetRequireTimeSheetEntryFlagOk returns a tuple with the RequireTimeSheetEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTimeSheetEntryFlag

`func (o *MyMember) SetRequireTimeSheetEntryFlag(v bool)`

SetRequireTimeSheetEntryFlag sets RequireTimeSheetEntryFlag field to given value.

### HasRequireTimeSheetEntryFlag

`func (o *MyMember) HasRequireTimeSheetEntryFlag() bool`

HasRequireTimeSheetEntryFlag returns a boolean if a field has been set.

### SetRequireTimeSheetEntryFlagNil

`func (o *MyMember) SetRequireTimeSheetEntryFlagNil(b bool)`

 SetRequireTimeSheetEntryFlagNil sets the value for RequireTimeSheetEntryFlag to be an explicit nil

### UnsetRequireTimeSheetEntryFlag
`func (o *MyMember) UnsetRequireTimeSheetEntryFlag()`

UnsetRequireTimeSheetEntryFlag ensures that no value is present for RequireTimeSheetEntryFlag, not even an explicit nil
### GetRequireStartAndEndTimeOnTimeEntryFlag

`func (o *MyMember) GetRequireStartAndEndTimeOnTimeEntryFlag() bool`

GetRequireStartAndEndTimeOnTimeEntryFlag returns the RequireStartAndEndTimeOnTimeEntryFlag field if non-nil, zero value otherwise.

### GetRequireStartAndEndTimeOnTimeEntryFlagOk

`func (o *MyMember) GetRequireStartAndEndTimeOnTimeEntryFlagOk() (*bool, bool)`

GetRequireStartAndEndTimeOnTimeEntryFlagOk returns a tuple with the RequireStartAndEndTimeOnTimeEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireStartAndEndTimeOnTimeEntryFlag

`func (o *MyMember) SetRequireStartAndEndTimeOnTimeEntryFlag(v bool)`

SetRequireStartAndEndTimeOnTimeEntryFlag sets RequireStartAndEndTimeOnTimeEntryFlag field to given value.

### HasRequireStartAndEndTimeOnTimeEntryFlag

`func (o *MyMember) HasRequireStartAndEndTimeOnTimeEntryFlag() bool`

HasRequireStartAndEndTimeOnTimeEntryFlag returns a boolean if a field has been set.

### SetRequireStartAndEndTimeOnTimeEntryFlagNil

`func (o *MyMember) SetRequireStartAndEndTimeOnTimeEntryFlagNil(b bool)`

 SetRequireStartAndEndTimeOnTimeEntryFlagNil sets the value for RequireStartAndEndTimeOnTimeEntryFlag to be an explicit nil

### UnsetRequireStartAndEndTimeOnTimeEntryFlag
`func (o *MyMember) UnsetRequireStartAndEndTimeOnTimeEntryFlag()`

UnsetRequireStartAndEndTimeOnTimeEntryFlag ensures that no value is present for RequireStartAndEndTimeOnTimeEntryFlag, not even an explicit nil
### GetAllowInCellEntryOnTimeSheet

`func (o *MyMember) GetAllowInCellEntryOnTimeSheet() bool`

GetAllowInCellEntryOnTimeSheet returns the AllowInCellEntryOnTimeSheet field if non-nil, zero value otherwise.

### GetAllowInCellEntryOnTimeSheetOk

`func (o *MyMember) GetAllowInCellEntryOnTimeSheetOk() (*bool, bool)`

GetAllowInCellEntryOnTimeSheetOk returns a tuple with the AllowInCellEntryOnTimeSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInCellEntryOnTimeSheet

`func (o *MyMember) SetAllowInCellEntryOnTimeSheet(v bool)`

SetAllowInCellEntryOnTimeSheet sets AllowInCellEntryOnTimeSheet field to given value.

### HasAllowInCellEntryOnTimeSheet

`func (o *MyMember) HasAllowInCellEntryOnTimeSheet() bool`

HasAllowInCellEntryOnTimeSheet returns a boolean if a field has been set.

### SetAllowInCellEntryOnTimeSheetNil

`func (o *MyMember) SetAllowInCellEntryOnTimeSheetNil(b bool)`

 SetAllowInCellEntryOnTimeSheetNil sets the value for AllowInCellEntryOnTimeSheet to be an explicit nil

### UnsetAllowInCellEntryOnTimeSheet
`func (o *MyMember) UnsetAllowInCellEntryOnTimeSheet()`

UnsetAllowInCellEntryOnTimeSheet ensures that no value is present for AllowInCellEntryOnTimeSheet, not even an explicit nil
### GetEnterTimeAgainstCompanyFlag

`func (o *MyMember) GetEnterTimeAgainstCompanyFlag() bool`

GetEnterTimeAgainstCompanyFlag returns the EnterTimeAgainstCompanyFlag field if non-nil, zero value otherwise.

### GetEnterTimeAgainstCompanyFlagOk

`func (o *MyMember) GetEnterTimeAgainstCompanyFlagOk() (*bool, bool)`

GetEnterTimeAgainstCompanyFlagOk returns a tuple with the EnterTimeAgainstCompanyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterTimeAgainstCompanyFlag

`func (o *MyMember) SetEnterTimeAgainstCompanyFlag(v bool)`

SetEnterTimeAgainstCompanyFlag sets EnterTimeAgainstCompanyFlag field to given value.

### HasEnterTimeAgainstCompanyFlag

`func (o *MyMember) HasEnterTimeAgainstCompanyFlag() bool`

HasEnterTimeAgainstCompanyFlag returns a boolean if a field has been set.

### SetEnterTimeAgainstCompanyFlagNil

`func (o *MyMember) SetEnterTimeAgainstCompanyFlagNil(b bool)`

 SetEnterTimeAgainstCompanyFlagNil sets the value for EnterTimeAgainstCompanyFlag to be an explicit nil

### UnsetEnterTimeAgainstCompanyFlag
`func (o *MyMember) UnsetEnterTimeAgainstCompanyFlag()`

UnsetEnterTimeAgainstCompanyFlag ensures that no value is present for EnterTimeAgainstCompanyFlag, not even an explicit nil
### GetAllowExpensesEnteredAgainstCompaniesFlag

`func (o *MyMember) GetAllowExpensesEnteredAgainstCompaniesFlag() bool`

GetAllowExpensesEnteredAgainstCompaniesFlag returns the AllowExpensesEnteredAgainstCompaniesFlag field if non-nil, zero value otherwise.

### GetAllowExpensesEnteredAgainstCompaniesFlagOk

`func (o *MyMember) GetAllowExpensesEnteredAgainstCompaniesFlagOk() (*bool, bool)`

GetAllowExpensesEnteredAgainstCompaniesFlagOk returns a tuple with the AllowExpensesEnteredAgainstCompaniesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExpensesEnteredAgainstCompaniesFlag

`func (o *MyMember) SetAllowExpensesEnteredAgainstCompaniesFlag(v bool)`

SetAllowExpensesEnteredAgainstCompaniesFlag sets AllowExpensesEnteredAgainstCompaniesFlag field to given value.

### HasAllowExpensesEnteredAgainstCompaniesFlag

`func (o *MyMember) HasAllowExpensesEnteredAgainstCompaniesFlag() bool`

HasAllowExpensesEnteredAgainstCompaniesFlag returns a boolean if a field has been set.

### SetAllowExpensesEnteredAgainstCompaniesFlagNil

`func (o *MyMember) SetAllowExpensesEnteredAgainstCompaniesFlagNil(b bool)`

 SetAllowExpensesEnteredAgainstCompaniesFlagNil sets the value for AllowExpensesEnteredAgainstCompaniesFlag to be an explicit nil

### UnsetAllowExpensesEnteredAgainstCompaniesFlag
`func (o *MyMember) UnsetAllowExpensesEnteredAgainstCompaniesFlag()`

UnsetAllowExpensesEnteredAgainstCompaniesFlag ensures that no value is present for AllowExpensesEnteredAgainstCompaniesFlag, not even an explicit nil
### GetTimeReminderEmailFlag

`func (o *MyMember) GetTimeReminderEmailFlag() bool`

GetTimeReminderEmailFlag returns the TimeReminderEmailFlag field if non-nil, zero value otherwise.

### GetTimeReminderEmailFlagOk

`func (o *MyMember) GetTimeReminderEmailFlagOk() (*bool, bool)`

GetTimeReminderEmailFlagOk returns a tuple with the TimeReminderEmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeReminderEmailFlag

`func (o *MyMember) SetTimeReminderEmailFlag(v bool)`

SetTimeReminderEmailFlag sets TimeReminderEmailFlag field to given value.

### HasTimeReminderEmailFlag

`func (o *MyMember) HasTimeReminderEmailFlag() bool`

HasTimeReminderEmailFlag returns a boolean if a field has been set.

### SetTimeReminderEmailFlagNil

`func (o *MyMember) SetTimeReminderEmailFlagNil(b bool)`

 SetTimeReminderEmailFlagNil sets the value for TimeReminderEmailFlag to be an explicit nil

### UnsetTimeReminderEmailFlag
`func (o *MyMember) UnsetTimeReminderEmailFlag()`

UnsetTimeReminderEmailFlag ensures that no value is present for TimeReminderEmailFlag, not even an explicit nil
### GetDaysTolerance

`func (o *MyMember) GetDaysTolerance() int32`

GetDaysTolerance returns the DaysTolerance field if non-nil, zero value otherwise.

### GetDaysToleranceOk

`func (o *MyMember) GetDaysToleranceOk() (*int32, bool)`

GetDaysToleranceOk returns a tuple with the DaysTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysTolerance

`func (o *MyMember) SetDaysTolerance(v int32)`

SetDaysTolerance sets DaysTolerance field to given value.

### HasDaysTolerance

`func (o *MyMember) HasDaysTolerance() bool`

HasDaysTolerance returns a boolean if a field has been set.

### SetDaysToleranceNil

`func (o *MyMember) SetDaysToleranceNil(b bool)`

 SetDaysToleranceNil sets the value for DaysTolerance to be an explicit nil

### UnsetDaysTolerance
`func (o *MyMember) UnsetDaysTolerance()`

UnsetDaysTolerance ensures that no value is present for DaysTolerance, not even an explicit nil
### GetMinimumHours

`func (o *MyMember) GetMinimumHours() float64`

GetMinimumHours returns the MinimumHours field if non-nil, zero value otherwise.

### GetMinimumHoursOk

`func (o *MyMember) GetMinimumHoursOk() (*float64, bool)`

GetMinimumHoursOk returns a tuple with the MinimumHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumHours

`func (o *MyMember) SetMinimumHours(v float64)`

SetMinimumHours sets MinimumHours field to given value.

### HasMinimumHours

`func (o *MyMember) HasMinimumHours() bool`

HasMinimumHours returns a boolean if a field has been set.

### SetMinimumHoursNil

`func (o *MyMember) SetMinimumHoursNil(b bool)`

 SetMinimumHoursNil sets the value for MinimumHours to be an explicit nil

### UnsetMinimumHours
`func (o *MyMember) UnsetMinimumHours()`

UnsetMinimumHours ensures that no value is present for MinimumHours, not even an explicit nil
### GetTimeSheetStartDate

`func (o *MyMember) GetTimeSheetStartDate() string`

GetTimeSheetStartDate returns the TimeSheetStartDate field if non-nil, zero value otherwise.

### GetTimeSheetStartDateOk

`func (o *MyMember) GetTimeSheetStartDateOk() (*string, bool)`

GetTimeSheetStartDateOk returns a tuple with the TimeSheetStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSheetStartDate

`func (o *MyMember) SetTimeSheetStartDate(v string)`

SetTimeSheetStartDate sets TimeSheetStartDate field to given value.

### HasTimeSheetStartDate

`func (o *MyMember) HasTimeSheetStartDate() bool`

HasTimeSheetStartDate returns a boolean if a field has been set.

### GetHireDate

`func (o *MyMember) GetHireDate() string`

GetHireDate returns the HireDate field if non-nil, zero value otherwise.

### GetHireDateOk

`func (o *MyMember) GetHireDateOk() (*string, bool)`

GetHireDateOk returns a tuple with the HireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireDate

`func (o *MyMember) SetHireDate(v string)`

SetHireDate sets HireDate field to given value.

### HasHireDate

`func (o *MyMember) HasHireDate() bool`

HasHireDate returns a boolean if a field has been set.

### GetServiceDefaultLocation

`func (o *MyMember) GetServiceDefaultLocation() SystemLocationReference`

GetServiceDefaultLocation returns the ServiceDefaultLocation field if non-nil, zero value otherwise.

### GetServiceDefaultLocationOk

`func (o *MyMember) GetServiceDefaultLocationOk() (*SystemLocationReference, bool)`

GetServiceDefaultLocationOk returns a tuple with the ServiceDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultLocation

`func (o *MyMember) SetServiceDefaultLocation(v SystemLocationReference)`

SetServiceDefaultLocation sets ServiceDefaultLocation field to given value.

### HasServiceDefaultLocation

`func (o *MyMember) HasServiceDefaultLocation() bool`

HasServiceDefaultLocation returns a boolean if a field has been set.

### GetServiceDefaultDepartment

`func (o *MyMember) GetServiceDefaultDepartment() SystemDepartmentReference`

GetServiceDefaultDepartment returns the ServiceDefaultDepartment field if non-nil, zero value otherwise.

### GetServiceDefaultDepartmentOk

`func (o *MyMember) GetServiceDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetServiceDefaultDepartmentOk returns a tuple with the ServiceDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultDepartment

`func (o *MyMember) SetServiceDefaultDepartment(v SystemDepartmentReference)`

SetServiceDefaultDepartment sets ServiceDefaultDepartment field to given value.

### HasServiceDefaultDepartment

`func (o *MyMember) HasServiceDefaultDepartment() bool`

HasServiceDefaultDepartment returns a boolean if a field has been set.

### GetServiceDefaultBoard

`func (o *MyMember) GetServiceDefaultBoard() BoardReference`

GetServiceDefaultBoard returns the ServiceDefaultBoard field if non-nil, zero value otherwise.

### GetServiceDefaultBoardOk

`func (o *MyMember) GetServiceDefaultBoardOk() (*BoardReference, bool)`

GetServiceDefaultBoardOk returns a tuple with the ServiceDefaultBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultBoard

`func (o *MyMember) SetServiceDefaultBoard(v BoardReference)`

SetServiceDefaultBoard sets ServiceDefaultBoard field to given value.

### HasServiceDefaultBoard

`func (o *MyMember) HasServiceDefaultBoard() bool`

HasServiceDefaultBoard returns a boolean if a field has been set.

### GetRestrictServiceDefaultLocationFlag

`func (o *MyMember) GetRestrictServiceDefaultLocationFlag() bool`

GetRestrictServiceDefaultLocationFlag returns the RestrictServiceDefaultLocationFlag field if non-nil, zero value otherwise.

### GetRestrictServiceDefaultLocationFlagOk

`func (o *MyMember) GetRestrictServiceDefaultLocationFlagOk() (*bool, bool)`

GetRestrictServiceDefaultLocationFlagOk returns a tuple with the RestrictServiceDefaultLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictServiceDefaultLocationFlag

`func (o *MyMember) SetRestrictServiceDefaultLocationFlag(v bool)`

SetRestrictServiceDefaultLocationFlag sets RestrictServiceDefaultLocationFlag field to given value.

### HasRestrictServiceDefaultLocationFlag

`func (o *MyMember) HasRestrictServiceDefaultLocationFlag() bool`

HasRestrictServiceDefaultLocationFlag returns a boolean if a field has been set.

### SetRestrictServiceDefaultLocationFlagNil

`func (o *MyMember) SetRestrictServiceDefaultLocationFlagNil(b bool)`

 SetRestrictServiceDefaultLocationFlagNil sets the value for RestrictServiceDefaultLocationFlag to be an explicit nil

### UnsetRestrictServiceDefaultLocationFlag
`func (o *MyMember) UnsetRestrictServiceDefaultLocationFlag()`

UnsetRestrictServiceDefaultLocationFlag ensures that no value is present for RestrictServiceDefaultLocationFlag, not even an explicit nil
### GetRestrictServiceDefaultDepartmentFlag

`func (o *MyMember) GetRestrictServiceDefaultDepartmentFlag() bool`

GetRestrictServiceDefaultDepartmentFlag returns the RestrictServiceDefaultDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictServiceDefaultDepartmentFlagOk

`func (o *MyMember) GetRestrictServiceDefaultDepartmentFlagOk() (*bool, bool)`

GetRestrictServiceDefaultDepartmentFlagOk returns a tuple with the RestrictServiceDefaultDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictServiceDefaultDepartmentFlag

`func (o *MyMember) SetRestrictServiceDefaultDepartmentFlag(v bool)`

SetRestrictServiceDefaultDepartmentFlag sets RestrictServiceDefaultDepartmentFlag field to given value.

### HasRestrictServiceDefaultDepartmentFlag

`func (o *MyMember) HasRestrictServiceDefaultDepartmentFlag() bool`

HasRestrictServiceDefaultDepartmentFlag returns a boolean if a field has been set.

### SetRestrictServiceDefaultDepartmentFlagNil

`func (o *MyMember) SetRestrictServiceDefaultDepartmentFlagNil(b bool)`

 SetRestrictServiceDefaultDepartmentFlagNil sets the value for RestrictServiceDefaultDepartmentFlag to be an explicit nil

### UnsetRestrictServiceDefaultDepartmentFlag
`func (o *MyMember) UnsetRestrictServiceDefaultDepartmentFlag()`

UnsetRestrictServiceDefaultDepartmentFlag ensures that no value is present for RestrictServiceDefaultDepartmentFlag, not even an explicit nil
### GetExcludedServiceBoardIds

`func (o *MyMember) GetExcludedServiceBoardIds() []int32`

GetExcludedServiceBoardIds returns the ExcludedServiceBoardIds field if non-nil, zero value otherwise.

### GetExcludedServiceBoardIdsOk

`func (o *MyMember) GetExcludedServiceBoardIdsOk() (*[]int32, bool)`

GetExcludedServiceBoardIdsOk returns a tuple with the ExcludedServiceBoardIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedServiceBoardIds

`func (o *MyMember) SetExcludedServiceBoardIds(v []int32)`

SetExcludedServiceBoardIds sets ExcludedServiceBoardIds field to given value.

### HasExcludedServiceBoardIds

`func (o *MyMember) HasExcludedServiceBoardIds() bool`

HasExcludedServiceBoardIds returns a boolean if a field has been set.

### GetProjectDefaultLocation

`func (o *MyMember) GetProjectDefaultLocation() SystemLocationReference`

GetProjectDefaultLocation returns the ProjectDefaultLocation field if non-nil, zero value otherwise.

### GetProjectDefaultLocationOk

`func (o *MyMember) GetProjectDefaultLocationOk() (*SystemLocationReference, bool)`

GetProjectDefaultLocationOk returns a tuple with the ProjectDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultLocation

`func (o *MyMember) SetProjectDefaultLocation(v SystemLocationReference)`

SetProjectDefaultLocation sets ProjectDefaultLocation field to given value.

### HasProjectDefaultLocation

`func (o *MyMember) HasProjectDefaultLocation() bool`

HasProjectDefaultLocation returns a boolean if a field has been set.

### GetProjectDefaultDepartment

`func (o *MyMember) GetProjectDefaultDepartment() SystemDepartmentReference`

GetProjectDefaultDepartment returns the ProjectDefaultDepartment field if non-nil, zero value otherwise.

### GetProjectDefaultDepartmentOk

`func (o *MyMember) GetProjectDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetProjectDefaultDepartmentOk returns a tuple with the ProjectDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultDepartment

`func (o *MyMember) SetProjectDefaultDepartment(v SystemDepartmentReference)`

SetProjectDefaultDepartment sets ProjectDefaultDepartment field to given value.

### HasProjectDefaultDepartment

`func (o *MyMember) HasProjectDefaultDepartment() bool`

HasProjectDefaultDepartment returns a boolean if a field has been set.

### GetProjectDefaultBoard

`func (o *MyMember) GetProjectDefaultBoard() ProjectBoardReference`

GetProjectDefaultBoard returns the ProjectDefaultBoard field if non-nil, zero value otherwise.

### GetProjectDefaultBoardOk

`func (o *MyMember) GetProjectDefaultBoardOk() (*ProjectBoardReference, bool)`

GetProjectDefaultBoardOk returns a tuple with the ProjectDefaultBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultBoard

`func (o *MyMember) SetProjectDefaultBoard(v ProjectBoardReference)`

SetProjectDefaultBoard sets ProjectDefaultBoard field to given value.

### HasProjectDefaultBoard

`func (o *MyMember) HasProjectDefaultBoard() bool`

HasProjectDefaultBoard returns a boolean if a field has been set.

### GetRestrictProjectDefaultLocationFlag

`func (o *MyMember) GetRestrictProjectDefaultLocationFlag() bool`

GetRestrictProjectDefaultLocationFlag returns the RestrictProjectDefaultLocationFlag field if non-nil, zero value otherwise.

### GetRestrictProjectDefaultLocationFlagOk

`func (o *MyMember) GetRestrictProjectDefaultLocationFlagOk() (*bool, bool)`

GetRestrictProjectDefaultLocationFlagOk returns a tuple with the RestrictProjectDefaultLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictProjectDefaultLocationFlag

`func (o *MyMember) SetRestrictProjectDefaultLocationFlag(v bool)`

SetRestrictProjectDefaultLocationFlag sets RestrictProjectDefaultLocationFlag field to given value.

### HasRestrictProjectDefaultLocationFlag

`func (o *MyMember) HasRestrictProjectDefaultLocationFlag() bool`

HasRestrictProjectDefaultLocationFlag returns a boolean if a field has been set.

### SetRestrictProjectDefaultLocationFlagNil

`func (o *MyMember) SetRestrictProjectDefaultLocationFlagNil(b bool)`

 SetRestrictProjectDefaultLocationFlagNil sets the value for RestrictProjectDefaultLocationFlag to be an explicit nil

### UnsetRestrictProjectDefaultLocationFlag
`func (o *MyMember) UnsetRestrictProjectDefaultLocationFlag()`

UnsetRestrictProjectDefaultLocationFlag ensures that no value is present for RestrictProjectDefaultLocationFlag, not even an explicit nil
### GetRestrictProjectDefaultDepartmentFlag

`func (o *MyMember) GetRestrictProjectDefaultDepartmentFlag() bool`

GetRestrictProjectDefaultDepartmentFlag returns the RestrictProjectDefaultDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictProjectDefaultDepartmentFlagOk

`func (o *MyMember) GetRestrictProjectDefaultDepartmentFlagOk() (*bool, bool)`

GetRestrictProjectDefaultDepartmentFlagOk returns a tuple with the RestrictProjectDefaultDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictProjectDefaultDepartmentFlag

`func (o *MyMember) SetRestrictProjectDefaultDepartmentFlag(v bool)`

SetRestrictProjectDefaultDepartmentFlag sets RestrictProjectDefaultDepartmentFlag field to given value.

### HasRestrictProjectDefaultDepartmentFlag

`func (o *MyMember) HasRestrictProjectDefaultDepartmentFlag() bool`

HasRestrictProjectDefaultDepartmentFlag returns a boolean if a field has been set.

### SetRestrictProjectDefaultDepartmentFlagNil

`func (o *MyMember) SetRestrictProjectDefaultDepartmentFlagNil(b bool)`

 SetRestrictProjectDefaultDepartmentFlagNil sets the value for RestrictProjectDefaultDepartmentFlag to be an explicit nil

### UnsetRestrictProjectDefaultDepartmentFlag
`func (o *MyMember) UnsetRestrictProjectDefaultDepartmentFlag()`

UnsetRestrictProjectDefaultDepartmentFlag ensures that no value is present for RestrictProjectDefaultDepartmentFlag, not even an explicit nil
### GetExcludedProjectBoardIds

`func (o *MyMember) GetExcludedProjectBoardIds() []int32`

GetExcludedProjectBoardIds returns the ExcludedProjectBoardIds field if non-nil, zero value otherwise.

### GetExcludedProjectBoardIdsOk

`func (o *MyMember) GetExcludedProjectBoardIdsOk() (*[]int32, bool)`

GetExcludedProjectBoardIdsOk returns a tuple with the ExcludedProjectBoardIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProjectBoardIds

`func (o *MyMember) SetExcludedProjectBoardIds(v []int32)`

SetExcludedProjectBoardIds sets ExcludedProjectBoardIds field to given value.

### HasExcludedProjectBoardIds

`func (o *MyMember) HasExcludedProjectBoardIds() bool`

HasExcludedProjectBoardIds returns a boolean if a field has been set.

### GetScheduleDefaultLocation

`func (o *MyMember) GetScheduleDefaultLocation() SystemLocationReference`

GetScheduleDefaultLocation returns the ScheduleDefaultLocation field if non-nil, zero value otherwise.

### GetScheduleDefaultLocationOk

`func (o *MyMember) GetScheduleDefaultLocationOk() (*SystemLocationReference, bool)`

GetScheduleDefaultLocationOk returns a tuple with the ScheduleDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDefaultLocation

`func (o *MyMember) SetScheduleDefaultLocation(v SystemLocationReference)`

SetScheduleDefaultLocation sets ScheduleDefaultLocation field to given value.

### HasScheduleDefaultLocation

`func (o *MyMember) HasScheduleDefaultLocation() bool`

HasScheduleDefaultLocation returns a boolean if a field has been set.

### GetScheduleDefaultDepartment

`func (o *MyMember) GetScheduleDefaultDepartment() SystemDepartmentReference`

GetScheduleDefaultDepartment returns the ScheduleDefaultDepartment field if non-nil, zero value otherwise.

### GetScheduleDefaultDepartmentOk

`func (o *MyMember) GetScheduleDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetScheduleDefaultDepartmentOk returns a tuple with the ScheduleDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDefaultDepartment

`func (o *MyMember) SetScheduleDefaultDepartment(v SystemDepartmentReference)`

SetScheduleDefaultDepartment sets ScheduleDefaultDepartment field to given value.

### HasScheduleDefaultDepartment

`func (o *MyMember) HasScheduleDefaultDepartment() bool`

HasScheduleDefaultDepartment returns a boolean if a field has been set.

### GetScheduleCapacity

`func (o *MyMember) GetScheduleCapacity() float64`

GetScheduleCapacity returns the ScheduleCapacity field if non-nil, zero value otherwise.

### GetScheduleCapacityOk

`func (o *MyMember) GetScheduleCapacityOk() (*float64, bool)`

GetScheduleCapacityOk returns a tuple with the ScheduleCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCapacity

`func (o *MyMember) SetScheduleCapacity(v float64)`

SetScheduleCapacity sets ScheduleCapacity field to given value.

### HasScheduleCapacity

`func (o *MyMember) HasScheduleCapacity() bool`

HasScheduleCapacity returns a boolean if a field has been set.

### SetScheduleCapacityNil

`func (o *MyMember) SetScheduleCapacityNil(b bool)`

 SetScheduleCapacityNil sets the value for ScheduleCapacity to be an explicit nil

### UnsetScheduleCapacity
`func (o *MyMember) UnsetScheduleCapacity()`

UnsetScheduleCapacity ensures that no value is present for ScheduleCapacity, not even an explicit nil
### GetServiceLocation

`func (o *MyMember) GetServiceLocation() ServiceLocationReference`

GetServiceLocation returns the ServiceLocation field if non-nil, zero value otherwise.

### GetServiceLocationOk

`func (o *MyMember) GetServiceLocationOk() (*ServiceLocationReference, bool)`

GetServiceLocationOk returns a tuple with the ServiceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLocation

`func (o *MyMember) SetServiceLocation(v ServiceLocationReference)`

SetServiceLocation sets ServiceLocation field to given value.

### HasServiceLocation

`func (o *MyMember) HasServiceLocation() bool`

HasServiceLocation returns a boolean if a field has been set.

### GetRestrictScheduleFlag

`func (o *MyMember) GetRestrictScheduleFlag() bool`

GetRestrictScheduleFlag returns the RestrictScheduleFlag field if non-nil, zero value otherwise.

### GetRestrictScheduleFlagOk

`func (o *MyMember) GetRestrictScheduleFlagOk() (*bool, bool)`

GetRestrictScheduleFlagOk returns a tuple with the RestrictScheduleFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictScheduleFlag

`func (o *MyMember) SetRestrictScheduleFlag(v bool)`

SetRestrictScheduleFlag sets RestrictScheduleFlag field to given value.

### HasRestrictScheduleFlag

`func (o *MyMember) HasRestrictScheduleFlag() bool`

HasRestrictScheduleFlag returns a boolean if a field has been set.

### SetRestrictScheduleFlagNil

`func (o *MyMember) SetRestrictScheduleFlagNil(b bool)`

 SetRestrictScheduleFlagNil sets the value for RestrictScheduleFlag to be an explicit nil

### UnsetRestrictScheduleFlag
`func (o *MyMember) UnsetRestrictScheduleFlag()`

UnsetRestrictScheduleFlag ensures that no value is present for RestrictScheduleFlag, not even an explicit nil
### GetHideMemberInDispatchPortalFlag

`func (o *MyMember) GetHideMemberInDispatchPortalFlag() bool`

GetHideMemberInDispatchPortalFlag returns the HideMemberInDispatchPortalFlag field if non-nil, zero value otherwise.

### GetHideMemberInDispatchPortalFlagOk

`func (o *MyMember) GetHideMemberInDispatchPortalFlagOk() (*bool, bool)`

GetHideMemberInDispatchPortalFlagOk returns a tuple with the HideMemberInDispatchPortalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideMemberInDispatchPortalFlag

`func (o *MyMember) SetHideMemberInDispatchPortalFlag(v bool)`

SetHideMemberInDispatchPortalFlag sets HideMemberInDispatchPortalFlag field to given value.

### HasHideMemberInDispatchPortalFlag

`func (o *MyMember) HasHideMemberInDispatchPortalFlag() bool`

HasHideMemberInDispatchPortalFlag returns a boolean if a field has been set.

### SetHideMemberInDispatchPortalFlagNil

`func (o *MyMember) SetHideMemberInDispatchPortalFlagNil(b bool)`

 SetHideMemberInDispatchPortalFlagNil sets the value for HideMemberInDispatchPortalFlag to be an explicit nil

### UnsetHideMemberInDispatchPortalFlag
`func (o *MyMember) UnsetHideMemberInDispatchPortalFlag()`

UnsetHideMemberInDispatchPortalFlag ensures that no value is present for HideMemberInDispatchPortalFlag, not even an explicit nil
### GetCalendar

`func (o *MyMember) GetCalendar() CalendarReference`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *MyMember) GetCalendarOk() (*CalendarReference, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *MyMember) SetCalendar(v CalendarReference)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *MyMember) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### GetSalesDefaultLocation

`func (o *MyMember) GetSalesDefaultLocation() SystemLocationReference`

GetSalesDefaultLocation returns the SalesDefaultLocation field if non-nil, zero value otherwise.

### GetSalesDefaultLocationOk

`func (o *MyMember) GetSalesDefaultLocationOk() (*SystemLocationReference, bool)`

GetSalesDefaultLocationOk returns a tuple with the SalesDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDefaultLocation

`func (o *MyMember) SetSalesDefaultLocation(v SystemLocationReference)`

SetSalesDefaultLocation sets SalesDefaultLocation field to given value.

### HasSalesDefaultLocation

`func (o *MyMember) HasSalesDefaultLocation() bool`

HasSalesDefaultLocation returns a boolean if a field has been set.

### GetRestrictDefaultSalesTerritoryFlag

`func (o *MyMember) GetRestrictDefaultSalesTerritoryFlag() bool`

GetRestrictDefaultSalesTerritoryFlag returns the RestrictDefaultSalesTerritoryFlag field if non-nil, zero value otherwise.

### GetRestrictDefaultSalesTerritoryFlagOk

`func (o *MyMember) GetRestrictDefaultSalesTerritoryFlagOk() (*bool, bool)`

GetRestrictDefaultSalesTerritoryFlagOk returns a tuple with the RestrictDefaultSalesTerritoryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDefaultSalesTerritoryFlag

`func (o *MyMember) SetRestrictDefaultSalesTerritoryFlag(v bool)`

SetRestrictDefaultSalesTerritoryFlag sets RestrictDefaultSalesTerritoryFlag field to given value.

### HasRestrictDefaultSalesTerritoryFlag

`func (o *MyMember) HasRestrictDefaultSalesTerritoryFlag() bool`

HasRestrictDefaultSalesTerritoryFlag returns a boolean if a field has been set.

### SetRestrictDefaultSalesTerritoryFlagNil

`func (o *MyMember) SetRestrictDefaultSalesTerritoryFlagNil(b bool)`

 SetRestrictDefaultSalesTerritoryFlagNil sets the value for RestrictDefaultSalesTerritoryFlag to be an explicit nil

### UnsetRestrictDefaultSalesTerritoryFlag
`func (o *MyMember) UnsetRestrictDefaultSalesTerritoryFlag()`

UnsetRestrictDefaultSalesTerritoryFlag ensures that no value is present for RestrictDefaultSalesTerritoryFlag, not even an explicit nil
### GetWarehouse

`func (o *MyMember) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *MyMember) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *MyMember) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *MyMember) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetWarehouseBin

`func (o *MyMember) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *MyMember) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *MyMember) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.

### HasWarehouseBin

`func (o *MyMember) HasWarehouseBin() bool`

HasWarehouseBin returns a boolean if a field has been set.

### GetRestrictDefaultWarehouseFlag

`func (o *MyMember) GetRestrictDefaultWarehouseFlag() bool`

GetRestrictDefaultWarehouseFlag returns the RestrictDefaultWarehouseFlag field if non-nil, zero value otherwise.

### GetRestrictDefaultWarehouseFlagOk

`func (o *MyMember) GetRestrictDefaultWarehouseFlagOk() (*bool, bool)`

GetRestrictDefaultWarehouseFlagOk returns a tuple with the RestrictDefaultWarehouseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDefaultWarehouseFlag

`func (o *MyMember) SetRestrictDefaultWarehouseFlag(v bool)`

SetRestrictDefaultWarehouseFlag sets RestrictDefaultWarehouseFlag field to given value.

### HasRestrictDefaultWarehouseFlag

`func (o *MyMember) HasRestrictDefaultWarehouseFlag() bool`

HasRestrictDefaultWarehouseFlag returns a boolean if a field has been set.

### SetRestrictDefaultWarehouseFlagNil

`func (o *MyMember) SetRestrictDefaultWarehouseFlagNil(b bool)`

 SetRestrictDefaultWarehouseFlagNil sets the value for RestrictDefaultWarehouseFlag to be an explicit nil

### UnsetRestrictDefaultWarehouseFlag
`func (o *MyMember) UnsetRestrictDefaultWarehouseFlag()`

UnsetRestrictDefaultWarehouseFlag ensures that no value is present for RestrictDefaultWarehouseFlag, not even an explicit nil
### GetRestrictDefaultWarehouseBinFlag

`func (o *MyMember) GetRestrictDefaultWarehouseBinFlag() bool`

GetRestrictDefaultWarehouseBinFlag returns the RestrictDefaultWarehouseBinFlag field if non-nil, zero value otherwise.

### GetRestrictDefaultWarehouseBinFlagOk

`func (o *MyMember) GetRestrictDefaultWarehouseBinFlagOk() (*bool, bool)`

GetRestrictDefaultWarehouseBinFlagOk returns a tuple with the RestrictDefaultWarehouseBinFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDefaultWarehouseBinFlag

`func (o *MyMember) SetRestrictDefaultWarehouseBinFlag(v bool)`

SetRestrictDefaultWarehouseBinFlag sets RestrictDefaultWarehouseBinFlag field to given value.

### HasRestrictDefaultWarehouseBinFlag

`func (o *MyMember) HasRestrictDefaultWarehouseBinFlag() bool`

HasRestrictDefaultWarehouseBinFlag returns a boolean if a field has been set.

### SetRestrictDefaultWarehouseBinFlagNil

`func (o *MyMember) SetRestrictDefaultWarehouseBinFlagNil(b bool)`

 SetRestrictDefaultWarehouseBinFlagNil sets the value for RestrictDefaultWarehouseBinFlag to be an explicit nil

### UnsetRestrictDefaultWarehouseBinFlag
`func (o *MyMember) UnsetRestrictDefaultWarehouseBinFlag()`

UnsetRestrictDefaultWarehouseBinFlag ensures that no value is present for RestrictDefaultWarehouseBinFlag, not even an explicit nil
### GetMapiName

`func (o *MyMember) GetMapiName() string`

GetMapiName returns the MapiName field if non-nil, zero value otherwise.

### GetMapiNameOk

`func (o *MyMember) GetMapiNameOk() (*string, bool)`

GetMapiNameOk returns a tuple with the MapiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapiName

`func (o *MyMember) SetMapiName(v string)`

SetMapiName sets MapiName field to given value.

### HasMapiName

`func (o *MyMember) HasMapiName() bool`

HasMapiName returns a boolean if a field has been set.

### GetCalendarSyncIntegrationFlag

`func (o *MyMember) GetCalendarSyncIntegrationFlag() bool`

GetCalendarSyncIntegrationFlag returns the CalendarSyncIntegrationFlag field if non-nil, zero value otherwise.

### GetCalendarSyncIntegrationFlagOk

`func (o *MyMember) GetCalendarSyncIntegrationFlagOk() (*bool, bool)`

GetCalendarSyncIntegrationFlagOk returns a tuple with the CalendarSyncIntegrationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarSyncIntegrationFlag

`func (o *MyMember) SetCalendarSyncIntegrationFlag(v bool)`

SetCalendarSyncIntegrationFlag sets CalendarSyncIntegrationFlag field to given value.

### HasCalendarSyncIntegrationFlag

`func (o *MyMember) HasCalendarSyncIntegrationFlag() bool`

HasCalendarSyncIntegrationFlag returns a boolean if a field has been set.

### SetCalendarSyncIntegrationFlagNil

`func (o *MyMember) SetCalendarSyncIntegrationFlagNil(b bool)`

 SetCalendarSyncIntegrationFlagNil sets the value for CalendarSyncIntegrationFlag to be an explicit nil

### UnsetCalendarSyncIntegrationFlag
`func (o *MyMember) UnsetCalendarSyncIntegrationFlag()`

UnsetCalendarSyncIntegrationFlag ensures that no value is present for CalendarSyncIntegrationFlag, not even an explicit nil
### GetEnableLdapAuthenticationFlag

`func (o *MyMember) GetEnableLdapAuthenticationFlag() bool`

GetEnableLdapAuthenticationFlag returns the EnableLdapAuthenticationFlag field if non-nil, zero value otherwise.

### GetEnableLdapAuthenticationFlagOk

`func (o *MyMember) GetEnableLdapAuthenticationFlagOk() (*bool, bool)`

GetEnableLdapAuthenticationFlagOk returns a tuple with the EnableLdapAuthenticationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLdapAuthenticationFlag

`func (o *MyMember) SetEnableLdapAuthenticationFlag(v bool)`

SetEnableLdapAuthenticationFlag sets EnableLdapAuthenticationFlag field to given value.

### HasEnableLdapAuthenticationFlag

`func (o *MyMember) HasEnableLdapAuthenticationFlag() bool`

HasEnableLdapAuthenticationFlag returns a boolean if a field has been set.

### SetEnableLdapAuthenticationFlagNil

`func (o *MyMember) SetEnableLdapAuthenticationFlagNil(b bool)`

 SetEnableLdapAuthenticationFlagNil sets the value for EnableLdapAuthenticationFlag to be an explicit nil

### UnsetEnableLdapAuthenticationFlag
`func (o *MyMember) UnsetEnableLdapAuthenticationFlag()`

UnsetEnableLdapAuthenticationFlag ensures that no value is present for EnableLdapAuthenticationFlag, not even an explicit nil
### GetLdapConfiguration

`func (o *MyMember) GetLdapConfiguration() LdapConfigurationReference`

GetLdapConfiguration returns the LdapConfiguration field if non-nil, zero value otherwise.

### GetLdapConfigurationOk

`func (o *MyMember) GetLdapConfigurationOk() (*LdapConfigurationReference, bool)`

GetLdapConfigurationOk returns a tuple with the LdapConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapConfiguration

`func (o *MyMember) SetLdapConfiguration(v LdapConfigurationReference)`

SetLdapConfiguration sets LdapConfiguration field to given value.

### HasLdapConfiguration

`func (o *MyMember) HasLdapConfiguration() bool`

HasLdapConfiguration returns a boolean if a field has been set.

### GetLdapUserName

`func (o *MyMember) GetLdapUserName() string`

GetLdapUserName returns the LdapUserName field if non-nil, zero value otherwise.

### GetLdapUserNameOk

`func (o *MyMember) GetLdapUserNameOk() (*string, bool)`

GetLdapUserNameOk returns a tuple with the LdapUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserName

`func (o *MyMember) SetLdapUserName(v string)`

SetLdapUserName sets LdapUserName field to given value.

### HasLdapUserName

`func (o *MyMember) HasLdapUserName() bool`

HasLdapUserName returns a boolean if a field has been set.

### GetCompanyActivityTabFormat

`func (o *MyMember) GetCompanyActivityTabFormat() string`

GetCompanyActivityTabFormat returns the CompanyActivityTabFormat field if non-nil, zero value otherwise.

### GetCompanyActivityTabFormatOk

`func (o *MyMember) GetCompanyActivityTabFormatOk() (*string, bool)`

GetCompanyActivityTabFormatOk returns a tuple with the CompanyActivityTabFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyActivityTabFormat

`func (o *MyMember) SetCompanyActivityTabFormat(v string)`

SetCompanyActivityTabFormat sets CompanyActivityTabFormat field to given value.

### HasCompanyActivityTabFormat

`func (o *MyMember) HasCompanyActivityTabFormat() bool`

HasCompanyActivityTabFormat returns a boolean if a field has been set.

### SetCompanyActivityTabFormatNil

`func (o *MyMember) SetCompanyActivityTabFormatNil(b bool)`

 SetCompanyActivityTabFormatNil sets the value for CompanyActivityTabFormat to be an explicit nil

### UnsetCompanyActivityTabFormat
`func (o *MyMember) UnsetCompanyActivityTabFormat()`

UnsetCompanyActivityTabFormat ensures that no value is present for CompanyActivityTabFormat, not even an explicit nil
### GetInvoiceTimeTabFormat

`func (o *MyMember) GetInvoiceTimeTabFormat() string`

GetInvoiceTimeTabFormat returns the InvoiceTimeTabFormat field if non-nil, zero value otherwise.

### GetInvoiceTimeTabFormatOk

`func (o *MyMember) GetInvoiceTimeTabFormatOk() (*string, bool)`

GetInvoiceTimeTabFormatOk returns a tuple with the InvoiceTimeTabFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTimeTabFormat

`func (o *MyMember) SetInvoiceTimeTabFormat(v string)`

SetInvoiceTimeTabFormat sets InvoiceTimeTabFormat field to given value.

### HasInvoiceTimeTabFormat

`func (o *MyMember) HasInvoiceTimeTabFormat() bool`

HasInvoiceTimeTabFormat returns a boolean if a field has been set.

### SetInvoiceTimeTabFormatNil

`func (o *MyMember) SetInvoiceTimeTabFormatNil(b bool)`

 SetInvoiceTimeTabFormatNil sets the value for InvoiceTimeTabFormat to be an explicit nil

### UnsetInvoiceTimeTabFormat
`func (o *MyMember) UnsetInvoiceTimeTabFormat()`

UnsetInvoiceTimeTabFormat ensures that no value is present for InvoiceTimeTabFormat, not even an explicit nil
### GetInvoiceScreenDefaultTabFormat

`func (o *MyMember) GetInvoiceScreenDefaultTabFormat() string`

GetInvoiceScreenDefaultTabFormat returns the InvoiceScreenDefaultTabFormat field if non-nil, zero value otherwise.

### GetInvoiceScreenDefaultTabFormatOk

`func (o *MyMember) GetInvoiceScreenDefaultTabFormatOk() (*string, bool)`

GetInvoiceScreenDefaultTabFormatOk returns a tuple with the InvoiceScreenDefaultTabFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceScreenDefaultTabFormat

`func (o *MyMember) SetInvoiceScreenDefaultTabFormat(v string)`

SetInvoiceScreenDefaultTabFormat sets InvoiceScreenDefaultTabFormat field to given value.

### HasInvoiceScreenDefaultTabFormat

`func (o *MyMember) HasInvoiceScreenDefaultTabFormat() bool`

HasInvoiceScreenDefaultTabFormat returns a boolean if a field has been set.

### SetInvoiceScreenDefaultTabFormatNil

`func (o *MyMember) SetInvoiceScreenDefaultTabFormatNil(b bool)`

 SetInvoiceScreenDefaultTabFormatNil sets the value for InvoiceScreenDefaultTabFormat to be an explicit nil

### UnsetInvoiceScreenDefaultTabFormat
`func (o *MyMember) UnsetInvoiceScreenDefaultTabFormat()`

UnsetInvoiceScreenDefaultTabFormat ensures that no value is present for InvoiceScreenDefaultTabFormat, not even an explicit nil
### GetInvoicingDisplayOptions

`func (o *MyMember) GetInvoicingDisplayOptions() string`

GetInvoicingDisplayOptions returns the InvoicingDisplayOptions field if non-nil, zero value otherwise.

### GetInvoicingDisplayOptionsOk

`func (o *MyMember) GetInvoicingDisplayOptionsOk() (*string, bool)`

GetInvoicingDisplayOptionsOk returns a tuple with the InvoicingDisplayOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingDisplayOptions

`func (o *MyMember) SetInvoicingDisplayOptions(v string)`

SetInvoicingDisplayOptions sets InvoicingDisplayOptions field to given value.

### HasInvoicingDisplayOptions

`func (o *MyMember) HasInvoicingDisplayOptions() bool`

HasInvoicingDisplayOptions returns a boolean if a field has been set.

### SetInvoicingDisplayOptionsNil

`func (o *MyMember) SetInvoicingDisplayOptionsNil(b bool)`

 SetInvoicingDisplayOptionsNil sets the value for InvoicingDisplayOptions to be an explicit nil

### UnsetInvoicingDisplayOptions
`func (o *MyMember) UnsetInvoicingDisplayOptions()`

UnsetInvoicingDisplayOptions ensures that no value is present for InvoicingDisplayOptions, not even an explicit nil
### GetAgreementInvoicingDisplayOptions

`func (o *MyMember) GetAgreementInvoicingDisplayOptions() string`

GetAgreementInvoicingDisplayOptions returns the AgreementInvoicingDisplayOptions field if non-nil, zero value otherwise.

### GetAgreementInvoicingDisplayOptionsOk

`func (o *MyMember) GetAgreementInvoicingDisplayOptionsOk() (*string, bool)`

GetAgreementInvoicingDisplayOptionsOk returns a tuple with the AgreementInvoicingDisplayOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementInvoicingDisplayOptions

`func (o *MyMember) SetAgreementInvoicingDisplayOptions(v string)`

SetAgreementInvoicingDisplayOptions sets AgreementInvoicingDisplayOptions field to given value.

### HasAgreementInvoicingDisplayOptions

`func (o *MyMember) HasAgreementInvoicingDisplayOptions() bool`

HasAgreementInvoicingDisplayOptions returns a boolean if a field has been set.

### SetAgreementInvoicingDisplayOptionsNil

`func (o *MyMember) SetAgreementInvoicingDisplayOptionsNil(b bool)`

 SetAgreementInvoicingDisplayOptionsNil sets the value for AgreementInvoicingDisplayOptions to be an explicit nil

### UnsetAgreementInvoicingDisplayOptions
`func (o *MyMember) UnsetAgreementInvoicingDisplayOptions()`

UnsetAgreementInvoicingDisplayOptions ensures that no value is present for AgreementInvoicingDisplayOptions, not even an explicit nil
### GetCorelyticsUsername

`func (o *MyMember) GetCorelyticsUsername() string`

GetCorelyticsUsername returns the CorelyticsUsername field if non-nil, zero value otherwise.

### GetCorelyticsUsernameOk

`func (o *MyMember) GetCorelyticsUsernameOk() (*string, bool)`

GetCorelyticsUsernameOk returns a tuple with the CorelyticsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorelyticsUsername

`func (o *MyMember) SetCorelyticsUsername(v string)`

SetCorelyticsUsername sets CorelyticsUsername field to given value.

### HasCorelyticsUsername

`func (o *MyMember) HasCorelyticsUsername() bool`

HasCorelyticsUsername returns a boolean if a field has been set.

### GetCorelyticsPassword

`func (o *MyMember) GetCorelyticsPassword() string`

GetCorelyticsPassword returns the CorelyticsPassword field if non-nil, zero value otherwise.

### GetCorelyticsPasswordOk

`func (o *MyMember) GetCorelyticsPasswordOk() (*string, bool)`

GetCorelyticsPasswordOk returns a tuple with the CorelyticsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorelyticsPassword

`func (o *MyMember) SetCorelyticsPassword(v string)`

SetCorelyticsPassword sets CorelyticsPassword field to given value.

### HasCorelyticsPassword

`func (o *MyMember) HasCorelyticsPassword() bool`

HasCorelyticsPassword returns a boolean if a field has been set.

### GetAuthenticationServiceType

`func (o *MyMember) GetAuthenticationServiceType() string`

GetAuthenticationServiceType returns the AuthenticationServiceType field if non-nil, zero value otherwise.

### GetAuthenticationServiceTypeOk

`func (o *MyMember) GetAuthenticationServiceTypeOk() (*string, bool)`

GetAuthenticationServiceTypeOk returns a tuple with the AuthenticationServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationServiceType

`func (o *MyMember) SetAuthenticationServiceType(v string)`

SetAuthenticationServiceType sets AuthenticationServiceType field to given value.

### HasAuthenticationServiceType

`func (o *MyMember) HasAuthenticationServiceType() bool`

HasAuthenticationServiceType returns a boolean if a field has been set.

### SetAuthenticationServiceTypeNil

`func (o *MyMember) SetAuthenticationServiceTypeNil(b bool)`

 SetAuthenticationServiceTypeNil sets the value for AuthenticationServiceType to be an explicit nil

### UnsetAuthenticationServiceType
`func (o *MyMember) UnsetAuthenticationServiceType()`

UnsetAuthenticationServiceType ensures that no value is present for AuthenticationServiceType, not even an explicit nil
### GetTimebasedOneTimePasswordActivated

`func (o *MyMember) GetTimebasedOneTimePasswordActivated() bool`

GetTimebasedOneTimePasswordActivated returns the TimebasedOneTimePasswordActivated field if non-nil, zero value otherwise.

### GetTimebasedOneTimePasswordActivatedOk

`func (o *MyMember) GetTimebasedOneTimePasswordActivatedOk() (*bool, bool)`

GetTimebasedOneTimePasswordActivatedOk returns a tuple with the TimebasedOneTimePasswordActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimebasedOneTimePasswordActivated

`func (o *MyMember) SetTimebasedOneTimePasswordActivated(v bool)`

SetTimebasedOneTimePasswordActivated sets TimebasedOneTimePasswordActivated field to given value.

### HasTimebasedOneTimePasswordActivated

`func (o *MyMember) HasTimebasedOneTimePasswordActivated() bool`

HasTimebasedOneTimePasswordActivated returns a boolean if a field has been set.

### SetTimebasedOneTimePasswordActivatedNil

`func (o *MyMember) SetTimebasedOneTimePasswordActivatedNil(b bool)`

 SetTimebasedOneTimePasswordActivatedNil sets the value for TimebasedOneTimePasswordActivated to be an explicit nil

### UnsetTimebasedOneTimePasswordActivated
`func (o *MyMember) UnsetTimebasedOneTimePasswordActivated()`

UnsetTimebasedOneTimePasswordActivated ensures that no value is present for TimebasedOneTimePasswordActivated, not even an explicit nil
### GetDirectionalSync

`func (o *MyMember) GetDirectionalSync() DirectionalSyncReference`

GetDirectionalSync returns the DirectionalSync field if non-nil, zero value otherwise.

### GetDirectionalSyncOk

`func (o *MyMember) GetDirectionalSyncOk() (*DirectionalSyncReference, bool)`

GetDirectionalSyncOk returns a tuple with the DirectionalSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionalSync

`func (o *MyMember) SetDirectionalSync(v DirectionalSyncReference)`

SetDirectionalSync sets DirectionalSync field to given value.

### HasDirectionalSync

`func (o *MyMember) HasDirectionalSync() bool`

HasDirectionalSync returns a boolean if a field has been set.

### GetSsoSessionFlag

`func (o *MyMember) GetSsoSessionFlag() bool`

GetSsoSessionFlag returns the SsoSessionFlag field if non-nil, zero value otherwise.

### GetSsoSessionFlagOk

`func (o *MyMember) GetSsoSessionFlagOk() (*bool, bool)`

GetSsoSessionFlagOk returns a tuple with the SsoSessionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSessionFlag

`func (o *MyMember) SetSsoSessionFlag(v bool)`

SetSsoSessionFlag sets SsoSessionFlag field to given value.

### HasSsoSessionFlag

`func (o *MyMember) HasSsoSessionFlag() bool`

HasSsoSessionFlag returns a boolean if a field has been set.

### SetSsoSessionFlagNil

`func (o *MyMember) SetSsoSessionFlagNil(b bool)`

 SetSsoSessionFlagNil sets the value for SsoSessionFlag to be an explicit nil

### UnsetSsoSessionFlag
`func (o *MyMember) UnsetSsoSessionFlag()`

UnsetSsoSessionFlag ensures that no value is present for SsoSessionFlag, not even an explicit nil
### GetSsoClientId

`func (o *MyMember) GetSsoClientId() string`

GetSsoClientId returns the SsoClientId field if non-nil, zero value otherwise.

### GetSsoClientIdOk

`func (o *MyMember) GetSsoClientIdOk() (*string, bool)`

GetSsoClientIdOk returns a tuple with the SsoClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoClientId

`func (o *MyMember) SetSsoClientId(v string)`

SetSsoClientId sets SsoClientId field to given value.

### HasSsoClientId

`func (o *MyMember) HasSsoClientId() bool`

HasSsoClientId returns a boolean if a field has been set.

### GetInfo

`func (o *MyMember) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MyMember) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MyMember) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MyMember) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


