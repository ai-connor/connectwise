# MyMemberInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**MiddleInitial** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**DefaultEmail** | Pointer to **string** |  | [optional] 
**Photo** | Pointer to [**DocumentReference**](DocumentReference.md) |  | [optional] 
**LicenseClass** | Pointer to **NullableString** | F &#x3D; Full Member, A &#x3D; API Member, C &#x3D; StreamlineIT Member, X &#x3D; Subcontractor Member | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeZone** | Pointer to [**TimeZoneSetupReference**](TimeZoneSetupReference.md) |  | [optional] 
**UseBrowserLanguageFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**DefaultDepartment** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**DailyCapacity** | Pointer to **NullableFloat64** |  | [optional] 
**RequireExpenseEntryFlag** | Pointer to **NullableBool** |  | [optional] 
**RequireTimeSheetEntryFlag** | Pointer to **NullableBool** |  | [optional] 
**RequireStartAndEndTimeOnTimeEntryFlag** | Pointer to **NullableBool** |  | [optional] 
**EnterTimeAgainstCompanyFlag** | Pointer to **NullableBool** |  | [optional] 
**AllowExpensesEnteredAgainstCompaniesFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceDefaultBoard** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**ServiceDefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**ServiceDefaultDepartment** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
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
**SalesDefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Warehouse** | Pointer to [**WarehouseReference**](WarehouseReference.md) |  | [optional] 
**WarehouseBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**RestrictDefaultWarehouseFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictDefaultWarehouseBinFlag** | Pointer to **NullableBool** |  | [optional] 
**SsoSessionFlag** | Pointer to **NullableBool** |  | [optional] 
**SsoClientId** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMyMemberInfo

`func NewMyMemberInfo() *MyMemberInfo`

NewMyMemberInfo instantiates a new MyMemberInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyMemberInfoWithDefaults

`func NewMyMemberInfoWithDefaults() *MyMemberInfo`

NewMyMemberInfoWithDefaults instantiates a new MyMemberInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MyMemberInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyMemberInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyMemberInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MyMemberInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *MyMemberInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MyMemberInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MyMemberInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MyMemberInfo) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetFirstName

`func (o *MyMemberInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MyMemberInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MyMemberInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MyMemberInfo) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleInitial

`func (o *MyMemberInfo) GetMiddleInitial() string`

GetMiddleInitial returns the MiddleInitial field if non-nil, zero value otherwise.

### GetMiddleInitialOk

`func (o *MyMemberInfo) GetMiddleInitialOk() (*string, bool)`

GetMiddleInitialOk returns a tuple with the MiddleInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleInitial

`func (o *MyMemberInfo) SetMiddleInitial(v string)`

SetMiddleInitial sets MiddleInitial field to given value.

### HasMiddleInitial

`func (o *MyMemberInfo) HasMiddleInitial() bool`

HasMiddleInitial returns a boolean if a field has been set.

### GetLastName

`func (o *MyMemberInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MyMemberInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MyMemberInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MyMemberInfo) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetFullName

`func (o *MyMemberInfo) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *MyMemberInfo) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *MyMemberInfo) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *MyMemberInfo) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDefaultEmail

`func (o *MyMemberInfo) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *MyMemberInfo) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *MyMemberInfo) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.

### HasDefaultEmail

`func (o *MyMemberInfo) HasDefaultEmail() bool`

HasDefaultEmail returns a boolean if a field has been set.

### GetPhoto

`func (o *MyMemberInfo) GetPhoto() DocumentReference`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MyMemberInfo) GetPhotoOk() (*DocumentReference, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *MyMemberInfo) SetPhoto(v DocumentReference)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *MyMemberInfo) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetLicenseClass

`func (o *MyMemberInfo) GetLicenseClass() string`

GetLicenseClass returns the LicenseClass field if non-nil, zero value otherwise.

### GetLicenseClassOk

`func (o *MyMemberInfo) GetLicenseClassOk() (*string, bool)`

GetLicenseClassOk returns a tuple with the LicenseClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseClass

`func (o *MyMemberInfo) SetLicenseClass(v string)`

SetLicenseClass sets LicenseClass field to given value.

### HasLicenseClass

`func (o *MyMemberInfo) HasLicenseClass() bool`

HasLicenseClass returns a boolean if a field has been set.

### SetLicenseClassNil

`func (o *MyMemberInfo) SetLicenseClassNil(b bool)`

 SetLicenseClassNil sets the value for LicenseClass to be an explicit nil

### UnsetLicenseClass
`func (o *MyMemberInfo) UnsetLicenseClass()`

UnsetLicenseClass ensures that no value is present for LicenseClass, not even an explicit nil
### GetInactiveFlag

`func (o *MyMemberInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *MyMemberInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *MyMemberInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *MyMemberInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *MyMemberInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *MyMemberInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetTimeZone

`func (o *MyMemberInfo) GetTimeZone() TimeZoneSetupReference`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MyMemberInfo) GetTimeZoneOk() (*TimeZoneSetupReference, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MyMemberInfo) SetTimeZone(v TimeZoneSetupReference)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MyMemberInfo) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUseBrowserLanguageFlag

`func (o *MyMemberInfo) GetUseBrowserLanguageFlag() bool`

GetUseBrowserLanguageFlag returns the UseBrowserLanguageFlag field if non-nil, zero value otherwise.

### GetUseBrowserLanguageFlagOk

`func (o *MyMemberInfo) GetUseBrowserLanguageFlagOk() (*bool, bool)`

GetUseBrowserLanguageFlagOk returns a tuple with the UseBrowserLanguageFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBrowserLanguageFlag

`func (o *MyMemberInfo) SetUseBrowserLanguageFlag(v bool)`

SetUseBrowserLanguageFlag sets UseBrowserLanguageFlag field to given value.

### HasUseBrowserLanguageFlag

`func (o *MyMemberInfo) HasUseBrowserLanguageFlag() bool`

HasUseBrowserLanguageFlag returns a boolean if a field has been set.

### SetUseBrowserLanguageFlagNil

`func (o *MyMemberInfo) SetUseBrowserLanguageFlagNil(b bool)`

 SetUseBrowserLanguageFlagNil sets the value for UseBrowserLanguageFlag to be an explicit nil

### UnsetUseBrowserLanguageFlag
`func (o *MyMemberInfo) UnsetUseBrowserLanguageFlag()`

UnsetUseBrowserLanguageFlag ensures that no value is present for UseBrowserLanguageFlag, not even an explicit nil
### GetDefaultLocation

`func (o *MyMemberInfo) GetDefaultLocation() SystemLocationReference`

GetDefaultLocation returns the DefaultLocation field if non-nil, zero value otherwise.

### GetDefaultLocationOk

`func (o *MyMemberInfo) GetDefaultLocationOk() (*SystemLocationReference, bool)`

GetDefaultLocationOk returns a tuple with the DefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocation

`func (o *MyMemberInfo) SetDefaultLocation(v SystemLocationReference)`

SetDefaultLocation sets DefaultLocation field to given value.

### HasDefaultLocation

`func (o *MyMemberInfo) HasDefaultLocation() bool`

HasDefaultLocation returns a boolean if a field has been set.

### GetDefaultDepartment

`func (o *MyMemberInfo) GetDefaultDepartment() SystemDepartmentReference`

GetDefaultDepartment returns the DefaultDepartment field if non-nil, zero value otherwise.

### GetDefaultDepartmentOk

`func (o *MyMemberInfo) GetDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetDefaultDepartmentOk returns a tuple with the DefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDepartment

`func (o *MyMemberInfo) SetDefaultDepartment(v SystemDepartmentReference)`

SetDefaultDepartment sets DefaultDepartment field to given value.

### HasDefaultDepartment

`func (o *MyMemberInfo) HasDefaultDepartment() bool`

HasDefaultDepartment returns a boolean if a field has been set.

### GetWorkRole

`func (o *MyMemberInfo) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *MyMemberInfo) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *MyMemberInfo) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *MyMemberInfo) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *MyMemberInfo) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *MyMemberInfo) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *MyMemberInfo) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *MyMemberInfo) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetDailyCapacity

`func (o *MyMemberInfo) GetDailyCapacity() float64`

GetDailyCapacity returns the DailyCapacity field if non-nil, zero value otherwise.

### GetDailyCapacityOk

`func (o *MyMemberInfo) GetDailyCapacityOk() (*float64, bool)`

GetDailyCapacityOk returns a tuple with the DailyCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyCapacity

`func (o *MyMemberInfo) SetDailyCapacity(v float64)`

SetDailyCapacity sets DailyCapacity field to given value.

### HasDailyCapacity

`func (o *MyMemberInfo) HasDailyCapacity() bool`

HasDailyCapacity returns a boolean if a field has been set.

### SetDailyCapacityNil

`func (o *MyMemberInfo) SetDailyCapacityNil(b bool)`

 SetDailyCapacityNil sets the value for DailyCapacity to be an explicit nil

### UnsetDailyCapacity
`func (o *MyMemberInfo) UnsetDailyCapacity()`

UnsetDailyCapacity ensures that no value is present for DailyCapacity, not even an explicit nil
### GetRequireExpenseEntryFlag

`func (o *MyMemberInfo) GetRequireExpenseEntryFlag() bool`

GetRequireExpenseEntryFlag returns the RequireExpenseEntryFlag field if non-nil, zero value otherwise.

### GetRequireExpenseEntryFlagOk

`func (o *MyMemberInfo) GetRequireExpenseEntryFlagOk() (*bool, bool)`

GetRequireExpenseEntryFlagOk returns a tuple with the RequireExpenseEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExpenseEntryFlag

`func (o *MyMemberInfo) SetRequireExpenseEntryFlag(v bool)`

SetRequireExpenseEntryFlag sets RequireExpenseEntryFlag field to given value.

### HasRequireExpenseEntryFlag

`func (o *MyMemberInfo) HasRequireExpenseEntryFlag() bool`

HasRequireExpenseEntryFlag returns a boolean if a field has been set.

### SetRequireExpenseEntryFlagNil

`func (o *MyMemberInfo) SetRequireExpenseEntryFlagNil(b bool)`

 SetRequireExpenseEntryFlagNil sets the value for RequireExpenseEntryFlag to be an explicit nil

### UnsetRequireExpenseEntryFlag
`func (o *MyMemberInfo) UnsetRequireExpenseEntryFlag()`

UnsetRequireExpenseEntryFlag ensures that no value is present for RequireExpenseEntryFlag, not even an explicit nil
### GetRequireTimeSheetEntryFlag

`func (o *MyMemberInfo) GetRequireTimeSheetEntryFlag() bool`

GetRequireTimeSheetEntryFlag returns the RequireTimeSheetEntryFlag field if non-nil, zero value otherwise.

### GetRequireTimeSheetEntryFlagOk

`func (o *MyMemberInfo) GetRequireTimeSheetEntryFlagOk() (*bool, bool)`

GetRequireTimeSheetEntryFlagOk returns a tuple with the RequireTimeSheetEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTimeSheetEntryFlag

`func (o *MyMemberInfo) SetRequireTimeSheetEntryFlag(v bool)`

SetRequireTimeSheetEntryFlag sets RequireTimeSheetEntryFlag field to given value.

### HasRequireTimeSheetEntryFlag

`func (o *MyMemberInfo) HasRequireTimeSheetEntryFlag() bool`

HasRequireTimeSheetEntryFlag returns a boolean if a field has been set.

### SetRequireTimeSheetEntryFlagNil

`func (o *MyMemberInfo) SetRequireTimeSheetEntryFlagNil(b bool)`

 SetRequireTimeSheetEntryFlagNil sets the value for RequireTimeSheetEntryFlag to be an explicit nil

### UnsetRequireTimeSheetEntryFlag
`func (o *MyMemberInfo) UnsetRequireTimeSheetEntryFlag()`

UnsetRequireTimeSheetEntryFlag ensures that no value is present for RequireTimeSheetEntryFlag, not even an explicit nil
### GetRequireStartAndEndTimeOnTimeEntryFlag

`func (o *MyMemberInfo) GetRequireStartAndEndTimeOnTimeEntryFlag() bool`

GetRequireStartAndEndTimeOnTimeEntryFlag returns the RequireStartAndEndTimeOnTimeEntryFlag field if non-nil, zero value otherwise.

### GetRequireStartAndEndTimeOnTimeEntryFlagOk

`func (o *MyMemberInfo) GetRequireStartAndEndTimeOnTimeEntryFlagOk() (*bool, bool)`

GetRequireStartAndEndTimeOnTimeEntryFlagOk returns a tuple with the RequireStartAndEndTimeOnTimeEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireStartAndEndTimeOnTimeEntryFlag

`func (o *MyMemberInfo) SetRequireStartAndEndTimeOnTimeEntryFlag(v bool)`

SetRequireStartAndEndTimeOnTimeEntryFlag sets RequireStartAndEndTimeOnTimeEntryFlag field to given value.

### HasRequireStartAndEndTimeOnTimeEntryFlag

`func (o *MyMemberInfo) HasRequireStartAndEndTimeOnTimeEntryFlag() bool`

HasRequireStartAndEndTimeOnTimeEntryFlag returns a boolean if a field has been set.

### SetRequireStartAndEndTimeOnTimeEntryFlagNil

`func (o *MyMemberInfo) SetRequireStartAndEndTimeOnTimeEntryFlagNil(b bool)`

 SetRequireStartAndEndTimeOnTimeEntryFlagNil sets the value for RequireStartAndEndTimeOnTimeEntryFlag to be an explicit nil

### UnsetRequireStartAndEndTimeOnTimeEntryFlag
`func (o *MyMemberInfo) UnsetRequireStartAndEndTimeOnTimeEntryFlag()`

UnsetRequireStartAndEndTimeOnTimeEntryFlag ensures that no value is present for RequireStartAndEndTimeOnTimeEntryFlag, not even an explicit nil
### GetEnterTimeAgainstCompanyFlag

`func (o *MyMemberInfo) GetEnterTimeAgainstCompanyFlag() bool`

GetEnterTimeAgainstCompanyFlag returns the EnterTimeAgainstCompanyFlag field if non-nil, zero value otherwise.

### GetEnterTimeAgainstCompanyFlagOk

`func (o *MyMemberInfo) GetEnterTimeAgainstCompanyFlagOk() (*bool, bool)`

GetEnterTimeAgainstCompanyFlagOk returns a tuple with the EnterTimeAgainstCompanyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterTimeAgainstCompanyFlag

`func (o *MyMemberInfo) SetEnterTimeAgainstCompanyFlag(v bool)`

SetEnterTimeAgainstCompanyFlag sets EnterTimeAgainstCompanyFlag field to given value.

### HasEnterTimeAgainstCompanyFlag

`func (o *MyMemberInfo) HasEnterTimeAgainstCompanyFlag() bool`

HasEnterTimeAgainstCompanyFlag returns a boolean if a field has been set.

### SetEnterTimeAgainstCompanyFlagNil

`func (o *MyMemberInfo) SetEnterTimeAgainstCompanyFlagNil(b bool)`

 SetEnterTimeAgainstCompanyFlagNil sets the value for EnterTimeAgainstCompanyFlag to be an explicit nil

### UnsetEnterTimeAgainstCompanyFlag
`func (o *MyMemberInfo) UnsetEnterTimeAgainstCompanyFlag()`

UnsetEnterTimeAgainstCompanyFlag ensures that no value is present for EnterTimeAgainstCompanyFlag, not even an explicit nil
### GetAllowExpensesEnteredAgainstCompaniesFlag

`func (o *MyMemberInfo) GetAllowExpensesEnteredAgainstCompaniesFlag() bool`

GetAllowExpensesEnteredAgainstCompaniesFlag returns the AllowExpensesEnteredAgainstCompaniesFlag field if non-nil, zero value otherwise.

### GetAllowExpensesEnteredAgainstCompaniesFlagOk

`func (o *MyMemberInfo) GetAllowExpensesEnteredAgainstCompaniesFlagOk() (*bool, bool)`

GetAllowExpensesEnteredAgainstCompaniesFlagOk returns a tuple with the AllowExpensesEnteredAgainstCompaniesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExpensesEnteredAgainstCompaniesFlag

`func (o *MyMemberInfo) SetAllowExpensesEnteredAgainstCompaniesFlag(v bool)`

SetAllowExpensesEnteredAgainstCompaniesFlag sets AllowExpensesEnteredAgainstCompaniesFlag field to given value.

### HasAllowExpensesEnteredAgainstCompaniesFlag

`func (o *MyMemberInfo) HasAllowExpensesEnteredAgainstCompaniesFlag() bool`

HasAllowExpensesEnteredAgainstCompaniesFlag returns a boolean if a field has been set.

### SetAllowExpensesEnteredAgainstCompaniesFlagNil

`func (o *MyMemberInfo) SetAllowExpensesEnteredAgainstCompaniesFlagNil(b bool)`

 SetAllowExpensesEnteredAgainstCompaniesFlagNil sets the value for AllowExpensesEnteredAgainstCompaniesFlag to be an explicit nil

### UnsetAllowExpensesEnteredAgainstCompaniesFlag
`func (o *MyMemberInfo) UnsetAllowExpensesEnteredAgainstCompaniesFlag()`

UnsetAllowExpensesEnteredAgainstCompaniesFlag ensures that no value is present for AllowExpensesEnteredAgainstCompaniesFlag, not even an explicit nil
### GetServiceDefaultBoard

`func (o *MyMemberInfo) GetServiceDefaultBoard() BoardReference`

GetServiceDefaultBoard returns the ServiceDefaultBoard field if non-nil, zero value otherwise.

### GetServiceDefaultBoardOk

`func (o *MyMemberInfo) GetServiceDefaultBoardOk() (*BoardReference, bool)`

GetServiceDefaultBoardOk returns a tuple with the ServiceDefaultBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultBoard

`func (o *MyMemberInfo) SetServiceDefaultBoard(v BoardReference)`

SetServiceDefaultBoard sets ServiceDefaultBoard field to given value.

### HasServiceDefaultBoard

`func (o *MyMemberInfo) HasServiceDefaultBoard() bool`

HasServiceDefaultBoard returns a boolean if a field has been set.

### GetServiceDefaultLocation

`func (o *MyMemberInfo) GetServiceDefaultLocation() SystemLocationReference`

GetServiceDefaultLocation returns the ServiceDefaultLocation field if non-nil, zero value otherwise.

### GetServiceDefaultLocationOk

`func (o *MyMemberInfo) GetServiceDefaultLocationOk() (*SystemLocationReference, bool)`

GetServiceDefaultLocationOk returns a tuple with the ServiceDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultLocation

`func (o *MyMemberInfo) SetServiceDefaultLocation(v SystemLocationReference)`

SetServiceDefaultLocation sets ServiceDefaultLocation field to given value.

### HasServiceDefaultLocation

`func (o *MyMemberInfo) HasServiceDefaultLocation() bool`

HasServiceDefaultLocation returns a boolean if a field has been set.

### GetServiceDefaultDepartment

`func (o *MyMemberInfo) GetServiceDefaultDepartment() SystemDepartmentReference`

GetServiceDefaultDepartment returns the ServiceDefaultDepartment field if non-nil, zero value otherwise.

### GetServiceDefaultDepartmentOk

`func (o *MyMemberInfo) GetServiceDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetServiceDefaultDepartmentOk returns a tuple with the ServiceDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultDepartment

`func (o *MyMemberInfo) SetServiceDefaultDepartment(v SystemDepartmentReference)`

SetServiceDefaultDepartment sets ServiceDefaultDepartment field to given value.

### HasServiceDefaultDepartment

`func (o *MyMemberInfo) HasServiceDefaultDepartment() bool`

HasServiceDefaultDepartment returns a boolean if a field has been set.

### GetRestrictServiceDefaultLocationFlag

`func (o *MyMemberInfo) GetRestrictServiceDefaultLocationFlag() bool`

GetRestrictServiceDefaultLocationFlag returns the RestrictServiceDefaultLocationFlag field if non-nil, zero value otherwise.

### GetRestrictServiceDefaultLocationFlagOk

`func (o *MyMemberInfo) GetRestrictServiceDefaultLocationFlagOk() (*bool, bool)`

GetRestrictServiceDefaultLocationFlagOk returns a tuple with the RestrictServiceDefaultLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictServiceDefaultLocationFlag

`func (o *MyMemberInfo) SetRestrictServiceDefaultLocationFlag(v bool)`

SetRestrictServiceDefaultLocationFlag sets RestrictServiceDefaultLocationFlag field to given value.

### HasRestrictServiceDefaultLocationFlag

`func (o *MyMemberInfo) HasRestrictServiceDefaultLocationFlag() bool`

HasRestrictServiceDefaultLocationFlag returns a boolean if a field has been set.

### SetRestrictServiceDefaultLocationFlagNil

`func (o *MyMemberInfo) SetRestrictServiceDefaultLocationFlagNil(b bool)`

 SetRestrictServiceDefaultLocationFlagNil sets the value for RestrictServiceDefaultLocationFlag to be an explicit nil

### UnsetRestrictServiceDefaultLocationFlag
`func (o *MyMemberInfo) UnsetRestrictServiceDefaultLocationFlag()`

UnsetRestrictServiceDefaultLocationFlag ensures that no value is present for RestrictServiceDefaultLocationFlag, not even an explicit nil
### GetRestrictServiceDefaultDepartmentFlag

`func (o *MyMemberInfo) GetRestrictServiceDefaultDepartmentFlag() bool`

GetRestrictServiceDefaultDepartmentFlag returns the RestrictServiceDefaultDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictServiceDefaultDepartmentFlagOk

`func (o *MyMemberInfo) GetRestrictServiceDefaultDepartmentFlagOk() (*bool, bool)`

GetRestrictServiceDefaultDepartmentFlagOk returns a tuple with the RestrictServiceDefaultDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictServiceDefaultDepartmentFlag

`func (o *MyMemberInfo) SetRestrictServiceDefaultDepartmentFlag(v bool)`

SetRestrictServiceDefaultDepartmentFlag sets RestrictServiceDefaultDepartmentFlag field to given value.

### HasRestrictServiceDefaultDepartmentFlag

`func (o *MyMemberInfo) HasRestrictServiceDefaultDepartmentFlag() bool`

HasRestrictServiceDefaultDepartmentFlag returns a boolean if a field has been set.

### SetRestrictServiceDefaultDepartmentFlagNil

`func (o *MyMemberInfo) SetRestrictServiceDefaultDepartmentFlagNil(b bool)`

 SetRestrictServiceDefaultDepartmentFlagNil sets the value for RestrictServiceDefaultDepartmentFlag to be an explicit nil

### UnsetRestrictServiceDefaultDepartmentFlag
`func (o *MyMemberInfo) UnsetRestrictServiceDefaultDepartmentFlag()`

UnsetRestrictServiceDefaultDepartmentFlag ensures that no value is present for RestrictServiceDefaultDepartmentFlag, not even an explicit nil
### GetExcludedServiceBoardIds

`func (o *MyMemberInfo) GetExcludedServiceBoardIds() []int32`

GetExcludedServiceBoardIds returns the ExcludedServiceBoardIds field if non-nil, zero value otherwise.

### GetExcludedServiceBoardIdsOk

`func (o *MyMemberInfo) GetExcludedServiceBoardIdsOk() (*[]int32, bool)`

GetExcludedServiceBoardIdsOk returns a tuple with the ExcludedServiceBoardIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedServiceBoardIds

`func (o *MyMemberInfo) SetExcludedServiceBoardIds(v []int32)`

SetExcludedServiceBoardIds sets ExcludedServiceBoardIds field to given value.

### HasExcludedServiceBoardIds

`func (o *MyMemberInfo) HasExcludedServiceBoardIds() bool`

HasExcludedServiceBoardIds returns a boolean if a field has been set.

### GetProjectDefaultLocation

`func (o *MyMemberInfo) GetProjectDefaultLocation() SystemLocationReference`

GetProjectDefaultLocation returns the ProjectDefaultLocation field if non-nil, zero value otherwise.

### GetProjectDefaultLocationOk

`func (o *MyMemberInfo) GetProjectDefaultLocationOk() (*SystemLocationReference, bool)`

GetProjectDefaultLocationOk returns a tuple with the ProjectDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultLocation

`func (o *MyMemberInfo) SetProjectDefaultLocation(v SystemLocationReference)`

SetProjectDefaultLocation sets ProjectDefaultLocation field to given value.

### HasProjectDefaultLocation

`func (o *MyMemberInfo) HasProjectDefaultLocation() bool`

HasProjectDefaultLocation returns a boolean if a field has been set.

### GetProjectDefaultDepartment

`func (o *MyMemberInfo) GetProjectDefaultDepartment() SystemDepartmentReference`

GetProjectDefaultDepartment returns the ProjectDefaultDepartment field if non-nil, zero value otherwise.

### GetProjectDefaultDepartmentOk

`func (o *MyMemberInfo) GetProjectDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetProjectDefaultDepartmentOk returns a tuple with the ProjectDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultDepartment

`func (o *MyMemberInfo) SetProjectDefaultDepartment(v SystemDepartmentReference)`

SetProjectDefaultDepartment sets ProjectDefaultDepartment field to given value.

### HasProjectDefaultDepartment

`func (o *MyMemberInfo) HasProjectDefaultDepartment() bool`

HasProjectDefaultDepartment returns a boolean if a field has been set.

### GetProjectDefaultBoard

`func (o *MyMemberInfo) GetProjectDefaultBoard() ProjectBoardReference`

GetProjectDefaultBoard returns the ProjectDefaultBoard field if non-nil, zero value otherwise.

### GetProjectDefaultBoardOk

`func (o *MyMemberInfo) GetProjectDefaultBoardOk() (*ProjectBoardReference, bool)`

GetProjectDefaultBoardOk returns a tuple with the ProjectDefaultBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultBoard

`func (o *MyMemberInfo) SetProjectDefaultBoard(v ProjectBoardReference)`

SetProjectDefaultBoard sets ProjectDefaultBoard field to given value.

### HasProjectDefaultBoard

`func (o *MyMemberInfo) HasProjectDefaultBoard() bool`

HasProjectDefaultBoard returns a boolean if a field has been set.

### GetRestrictProjectDefaultLocationFlag

`func (o *MyMemberInfo) GetRestrictProjectDefaultLocationFlag() bool`

GetRestrictProjectDefaultLocationFlag returns the RestrictProjectDefaultLocationFlag field if non-nil, zero value otherwise.

### GetRestrictProjectDefaultLocationFlagOk

`func (o *MyMemberInfo) GetRestrictProjectDefaultLocationFlagOk() (*bool, bool)`

GetRestrictProjectDefaultLocationFlagOk returns a tuple with the RestrictProjectDefaultLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictProjectDefaultLocationFlag

`func (o *MyMemberInfo) SetRestrictProjectDefaultLocationFlag(v bool)`

SetRestrictProjectDefaultLocationFlag sets RestrictProjectDefaultLocationFlag field to given value.

### HasRestrictProjectDefaultLocationFlag

`func (o *MyMemberInfo) HasRestrictProjectDefaultLocationFlag() bool`

HasRestrictProjectDefaultLocationFlag returns a boolean if a field has been set.

### SetRestrictProjectDefaultLocationFlagNil

`func (o *MyMemberInfo) SetRestrictProjectDefaultLocationFlagNil(b bool)`

 SetRestrictProjectDefaultLocationFlagNil sets the value for RestrictProjectDefaultLocationFlag to be an explicit nil

### UnsetRestrictProjectDefaultLocationFlag
`func (o *MyMemberInfo) UnsetRestrictProjectDefaultLocationFlag()`

UnsetRestrictProjectDefaultLocationFlag ensures that no value is present for RestrictProjectDefaultLocationFlag, not even an explicit nil
### GetRestrictProjectDefaultDepartmentFlag

`func (o *MyMemberInfo) GetRestrictProjectDefaultDepartmentFlag() bool`

GetRestrictProjectDefaultDepartmentFlag returns the RestrictProjectDefaultDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictProjectDefaultDepartmentFlagOk

`func (o *MyMemberInfo) GetRestrictProjectDefaultDepartmentFlagOk() (*bool, bool)`

GetRestrictProjectDefaultDepartmentFlagOk returns a tuple with the RestrictProjectDefaultDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictProjectDefaultDepartmentFlag

`func (o *MyMemberInfo) SetRestrictProjectDefaultDepartmentFlag(v bool)`

SetRestrictProjectDefaultDepartmentFlag sets RestrictProjectDefaultDepartmentFlag field to given value.

### HasRestrictProjectDefaultDepartmentFlag

`func (o *MyMemberInfo) HasRestrictProjectDefaultDepartmentFlag() bool`

HasRestrictProjectDefaultDepartmentFlag returns a boolean if a field has been set.

### SetRestrictProjectDefaultDepartmentFlagNil

`func (o *MyMemberInfo) SetRestrictProjectDefaultDepartmentFlagNil(b bool)`

 SetRestrictProjectDefaultDepartmentFlagNil sets the value for RestrictProjectDefaultDepartmentFlag to be an explicit nil

### UnsetRestrictProjectDefaultDepartmentFlag
`func (o *MyMemberInfo) UnsetRestrictProjectDefaultDepartmentFlag()`

UnsetRestrictProjectDefaultDepartmentFlag ensures that no value is present for RestrictProjectDefaultDepartmentFlag, not even an explicit nil
### GetExcludedProjectBoardIds

`func (o *MyMemberInfo) GetExcludedProjectBoardIds() []int32`

GetExcludedProjectBoardIds returns the ExcludedProjectBoardIds field if non-nil, zero value otherwise.

### GetExcludedProjectBoardIdsOk

`func (o *MyMemberInfo) GetExcludedProjectBoardIdsOk() (*[]int32, bool)`

GetExcludedProjectBoardIdsOk returns a tuple with the ExcludedProjectBoardIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProjectBoardIds

`func (o *MyMemberInfo) SetExcludedProjectBoardIds(v []int32)`

SetExcludedProjectBoardIds sets ExcludedProjectBoardIds field to given value.

### HasExcludedProjectBoardIds

`func (o *MyMemberInfo) HasExcludedProjectBoardIds() bool`

HasExcludedProjectBoardIds returns a boolean if a field has been set.

### GetScheduleDefaultLocation

`func (o *MyMemberInfo) GetScheduleDefaultLocation() SystemLocationReference`

GetScheduleDefaultLocation returns the ScheduleDefaultLocation field if non-nil, zero value otherwise.

### GetScheduleDefaultLocationOk

`func (o *MyMemberInfo) GetScheduleDefaultLocationOk() (*SystemLocationReference, bool)`

GetScheduleDefaultLocationOk returns a tuple with the ScheduleDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDefaultLocation

`func (o *MyMemberInfo) SetScheduleDefaultLocation(v SystemLocationReference)`

SetScheduleDefaultLocation sets ScheduleDefaultLocation field to given value.

### HasScheduleDefaultLocation

`func (o *MyMemberInfo) HasScheduleDefaultLocation() bool`

HasScheduleDefaultLocation returns a boolean if a field has been set.

### GetScheduleDefaultDepartment

`func (o *MyMemberInfo) GetScheduleDefaultDepartment() SystemDepartmentReference`

GetScheduleDefaultDepartment returns the ScheduleDefaultDepartment field if non-nil, zero value otherwise.

### GetScheduleDefaultDepartmentOk

`func (o *MyMemberInfo) GetScheduleDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetScheduleDefaultDepartmentOk returns a tuple with the ScheduleDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDefaultDepartment

`func (o *MyMemberInfo) SetScheduleDefaultDepartment(v SystemDepartmentReference)`

SetScheduleDefaultDepartment sets ScheduleDefaultDepartment field to given value.

### HasScheduleDefaultDepartment

`func (o *MyMemberInfo) HasScheduleDefaultDepartment() bool`

HasScheduleDefaultDepartment returns a boolean if a field has been set.

### GetScheduleCapacity

`func (o *MyMemberInfo) GetScheduleCapacity() float64`

GetScheduleCapacity returns the ScheduleCapacity field if non-nil, zero value otherwise.

### GetScheduleCapacityOk

`func (o *MyMemberInfo) GetScheduleCapacityOk() (*float64, bool)`

GetScheduleCapacityOk returns a tuple with the ScheduleCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCapacity

`func (o *MyMemberInfo) SetScheduleCapacity(v float64)`

SetScheduleCapacity sets ScheduleCapacity field to given value.

### HasScheduleCapacity

`func (o *MyMemberInfo) HasScheduleCapacity() bool`

HasScheduleCapacity returns a boolean if a field has been set.

### SetScheduleCapacityNil

`func (o *MyMemberInfo) SetScheduleCapacityNil(b bool)`

 SetScheduleCapacityNil sets the value for ScheduleCapacity to be an explicit nil

### UnsetScheduleCapacity
`func (o *MyMemberInfo) UnsetScheduleCapacity()`

UnsetScheduleCapacity ensures that no value is present for ScheduleCapacity, not even an explicit nil
### GetServiceLocation

`func (o *MyMemberInfo) GetServiceLocation() ServiceLocationReference`

GetServiceLocation returns the ServiceLocation field if non-nil, zero value otherwise.

### GetServiceLocationOk

`func (o *MyMemberInfo) GetServiceLocationOk() (*ServiceLocationReference, bool)`

GetServiceLocationOk returns a tuple with the ServiceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLocation

`func (o *MyMemberInfo) SetServiceLocation(v ServiceLocationReference)`

SetServiceLocation sets ServiceLocation field to given value.

### HasServiceLocation

`func (o *MyMemberInfo) HasServiceLocation() bool`

HasServiceLocation returns a boolean if a field has been set.

### GetSalesDefaultLocation

`func (o *MyMemberInfo) GetSalesDefaultLocation() SystemLocationReference`

GetSalesDefaultLocation returns the SalesDefaultLocation field if non-nil, zero value otherwise.

### GetSalesDefaultLocationOk

`func (o *MyMemberInfo) GetSalesDefaultLocationOk() (*SystemLocationReference, bool)`

GetSalesDefaultLocationOk returns a tuple with the SalesDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDefaultLocation

`func (o *MyMemberInfo) SetSalesDefaultLocation(v SystemLocationReference)`

SetSalesDefaultLocation sets SalesDefaultLocation field to given value.

### HasSalesDefaultLocation

`func (o *MyMemberInfo) HasSalesDefaultLocation() bool`

HasSalesDefaultLocation returns a boolean if a field has been set.

### GetWarehouse

`func (o *MyMemberInfo) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *MyMemberInfo) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *MyMemberInfo) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *MyMemberInfo) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetWarehouseBin

`func (o *MyMemberInfo) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *MyMemberInfo) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *MyMemberInfo) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.

### HasWarehouseBin

`func (o *MyMemberInfo) HasWarehouseBin() bool`

HasWarehouseBin returns a boolean if a field has been set.

### GetRestrictDefaultWarehouseFlag

`func (o *MyMemberInfo) GetRestrictDefaultWarehouseFlag() bool`

GetRestrictDefaultWarehouseFlag returns the RestrictDefaultWarehouseFlag field if non-nil, zero value otherwise.

### GetRestrictDefaultWarehouseFlagOk

`func (o *MyMemberInfo) GetRestrictDefaultWarehouseFlagOk() (*bool, bool)`

GetRestrictDefaultWarehouseFlagOk returns a tuple with the RestrictDefaultWarehouseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDefaultWarehouseFlag

`func (o *MyMemberInfo) SetRestrictDefaultWarehouseFlag(v bool)`

SetRestrictDefaultWarehouseFlag sets RestrictDefaultWarehouseFlag field to given value.

### HasRestrictDefaultWarehouseFlag

`func (o *MyMemberInfo) HasRestrictDefaultWarehouseFlag() bool`

HasRestrictDefaultWarehouseFlag returns a boolean if a field has been set.

### SetRestrictDefaultWarehouseFlagNil

`func (o *MyMemberInfo) SetRestrictDefaultWarehouseFlagNil(b bool)`

 SetRestrictDefaultWarehouseFlagNil sets the value for RestrictDefaultWarehouseFlag to be an explicit nil

### UnsetRestrictDefaultWarehouseFlag
`func (o *MyMemberInfo) UnsetRestrictDefaultWarehouseFlag()`

UnsetRestrictDefaultWarehouseFlag ensures that no value is present for RestrictDefaultWarehouseFlag, not even an explicit nil
### GetRestrictDefaultWarehouseBinFlag

`func (o *MyMemberInfo) GetRestrictDefaultWarehouseBinFlag() bool`

GetRestrictDefaultWarehouseBinFlag returns the RestrictDefaultWarehouseBinFlag field if non-nil, zero value otherwise.

### GetRestrictDefaultWarehouseBinFlagOk

`func (o *MyMemberInfo) GetRestrictDefaultWarehouseBinFlagOk() (*bool, bool)`

GetRestrictDefaultWarehouseBinFlagOk returns a tuple with the RestrictDefaultWarehouseBinFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDefaultWarehouseBinFlag

`func (o *MyMemberInfo) SetRestrictDefaultWarehouseBinFlag(v bool)`

SetRestrictDefaultWarehouseBinFlag sets RestrictDefaultWarehouseBinFlag field to given value.

### HasRestrictDefaultWarehouseBinFlag

`func (o *MyMemberInfo) HasRestrictDefaultWarehouseBinFlag() bool`

HasRestrictDefaultWarehouseBinFlag returns a boolean if a field has been set.

### SetRestrictDefaultWarehouseBinFlagNil

`func (o *MyMemberInfo) SetRestrictDefaultWarehouseBinFlagNil(b bool)`

 SetRestrictDefaultWarehouseBinFlagNil sets the value for RestrictDefaultWarehouseBinFlag to be an explicit nil

### UnsetRestrictDefaultWarehouseBinFlag
`func (o *MyMemberInfo) UnsetRestrictDefaultWarehouseBinFlag()`

UnsetRestrictDefaultWarehouseBinFlag ensures that no value is present for RestrictDefaultWarehouseBinFlag, not even an explicit nil
### GetSsoSessionFlag

`func (o *MyMemberInfo) GetSsoSessionFlag() bool`

GetSsoSessionFlag returns the SsoSessionFlag field if non-nil, zero value otherwise.

### GetSsoSessionFlagOk

`func (o *MyMemberInfo) GetSsoSessionFlagOk() (*bool, bool)`

GetSsoSessionFlagOk returns a tuple with the SsoSessionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSessionFlag

`func (o *MyMemberInfo) SetSsoSessionFlag(v bool)`

SetSsoSessionFlag sets SsoSessionFlag field to given value.

### HasSsoSessionFlag

`func (o *MyMemberInfo) HasSsoSessionFlag() bool`

HasSsoSessionFlag returns a boolean if a field has been set.

### SetSsoSessionFlagNil

`func (o *MyMemberInfo) SetSsoSessionFlagNil(b bool)`

 SetSsoSessionFlagNil sets the value for SsoSessionFlag to be an explicit nil

### UnsetSsoSessionFlag
`func (o *MyMemberInfo) UnsetSsoSessionFlag()`

UnsetSsoSessionFlag ensures that no value is present for SsoSessionFlag, not even an explicit nil
### GetSsoClientId

`func (o *MyMemberInfo) GetSsoClientId() string`

GetSsoClientId returns the SsoClientId field if non-nil, zero value otherwise.

### GetSsoClientIdOk

`func (o *MyMemberInfo) GetSsoClientIdOk() (*string, bool)`

GetSsoClientIdOk returns a tuple with the SsoClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoClientId

`func (o *MyMemberInfo) SetSsoClientId(v string)`

SetSsoClientId sets SsoClientId field to given value.

### HasSsoClientId

`func (o *MyMemberInfo) HasSsoClientId() bool`

HasSsoClientId returns a boolean if a field has been set.

### GetInfo

`func (o *MyMemberInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MyMemberInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MyMemberInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MyMemberInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


