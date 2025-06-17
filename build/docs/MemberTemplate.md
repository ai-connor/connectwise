# MemberTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  Max length: 50; | 
**TemplateDescription** | Pointer to **string** |  Max length: 1024; | [optional] 
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

### NewMemberTemplate

`func NewMemberTemplate(identifier string, ) *MemberTemplate`

NewMemberTemplate instantiates a new MemberTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberTemplateWithDefaults

`func NewMemberTemplateWithDefaults() *MemberTemplate`

NewMemberTemplateWithDefaults instantiates a new MemberTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MemberTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *MemberTemplate) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MemberTemplate) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MemberTemplate) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetTemplateDescription

`func (o *MemberTemplate) GetTemplateDescription() string`

GetTemplateDescription returns the TemplateDescription field if non-nil, zero value otherwise.

### GetTemplateDescriptionOk

`func (o *MemberTemplate) GetTemplateDescriptionOk() (*string, bool)`

GetTemplateDescriptionOk returns a tuple with the TemplateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDescription

`func (o *MemberTemplate) SetTemplateDescription(v string)`

SetTemplateDescription sets TemplateDescription field to given value.

### HasTemplateDescription

`func (o *MemberTemplate) HasTemplateDescription() bool`

HasTemplateDescription returns a boolean if a field has been set.

### GetTitle

`func (o *MemberTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MemberTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MemberTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MemberTemplate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetReportCard

`func (o *MemberTemplate) GetReportCard() ReportCardReference`

GetReportCard returns the ReportCard field if non-nil, zero value otherwise.

### GetReportCardOk

`func (o *MemberTemplate) GetReportCardOk() (*ReportCardReference, bool)`

GetReportCardOk returns a tuple with the ReportCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCard

`func (o *MemberTemplate) SetReportCard(v ReportCardReference)`

SetReportCard sets ReportCard field to given value.

### HasReportCard

`func (o *MemberTemplate) HasReportCard() bool`

HasReportCard returns a boolean if a field has been set.

### GetEnableMobileFlag

`func (o *MemberTemplate) GetEnableMobileFlag() bool`

GetEnableMobileFlag returns the EnableMobileFlag field if non-nil, zero value otherwise.

### GetEnableMobileFlagOk

`func (o *MemberTemplate) GetEnableMobileFlagOk() (*bool, bool)`

GetEnableMobileFlagOk returns a tuple with the EnableMobileFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMobileFlag

`func (o *MemberTemplate) SetEnableMobileFlag(v bool)`

SetEnableMobileFlag sets EnableMobileFlag field to given value.

### HasEnableMobileFlag

`func (o *MemberTemplate) HasEnableMobileFlag() bool`

HasEnableMobileFlag returns a boolean if a field has been set.

### SetEnableMobileFlagNil

`func (o *MemberTemplate) SetEnableMobileFlagNil(b bool)`

 SetEnableMobileFlagNil sets the value for EnableMobileFlag to be an explicit nil

### UnsetEnableMobileFlag
`func (o *MemberTemplate) UnsetEnableMobileFlag()`

UnsetEnableMobileFlag ensures that no value is present for EnableMobileFlag, not even an explicit nil
### GetType

`func (o *MemberTemplate) GetType() MemberTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemberTemplate) GetTypeOk() (*MemberTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemberTemplate) SetType(v MemberTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *MemberTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeZone

`func (o *MemberTemplate) GetTimeZone() TimeZoneSetupReference`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MemberTemplate) GetTimeZoneOk() (*TimeZoneSetupReference, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MemberTemplate) SetTimeZone(v TimeZoneSetupReference)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MemberTemplate) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetPartnerPortalFlag

`func (o *MemberTemplate) GetPartnerPortalFlag() bool`

GetPartnerPortalFlag returns the PartnerPortalFlag field if non-nil, zero value otherwise.

### GetPartnerPortalFlagOk

`func (o *MemberTemplate) GetPartnerPortalFlagOk() (*bool, bool)`

GetPartnerPortalFlagOk returns a tuple with the PartnerPortalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPortalFlag

`func (o *MemberTemplate) SetPartnerPortalFlag(v bool)`

SetPartnerPortalFlag sets PartnerPortalFlag field to given value.

### HasPartnerPortalFlag

`func (o *MemberTemplate) HasPartnerPortalFlag() bool`

HasPartnerPortalFlag returns a boolean if a field has been set.

### SetPartnerPortalFlagNil

`func (o *MemberTemplate) SetPartnerPortalFlagNil(b bool)`

 SetPartnerPortalFlagNil sets the value for PartnerPortalFlag to be an explicit nil

### UnsetPartnerPortalFlag
`func (o *MemberTemplate) UnsetPartnerPortalFlag()`

UnsetPartnerPortalFlag ensures that no value is present for PartnerPortalFlag, not even an explicit nil
### GetStsUserAdminUrl

`func (o *MemberTemplate) GetStsUserAdminUrl() string`

GetStsUserAdminUrl returns the StsUserAdminUrl field if non-nil, zero value otherwise.

### GetStsUserAdminUrlOk

`func (o *MemberTemplate) GetStsUserAdminUrlOk() (*string, bool)`

GetStsUserAdminUrlOk returns a tuple with the StsUserAdminUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsUserAdminUrl

`func (o *MemberTemplate) SetStsUserAdminUrl(v string)`

SetStsUserAdminUrl sets StsUserAdminUrl field to given value.

### HasStsUserAdminUrl

`func (o *MemberTemplate) HasStsUserAdminUrl() bool`

HasStsUserAdminUrl returns a boolean if a field has been set.

### GetToastNotificationFlag

`func (o *MemberTemplate) GetToastNotificationFlag() bool`

GetToastNotificationFlag returns the ToastNotificationFlag field if non-nil, zero value otherwise.

### GetToastNotificationFlagOk

`func (o *MemberTemplate) GetToastNotificationFlagOk() (*bool, bool)`

GetToastNotificationFlagOk returns a tuple with the ToastNotificationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToastNotificationFlag

`func (o *MemberTemplate) SetToastNotificationFlag(v bool)`

SetToastNotificationFlag sets ToastNotificationFlag field to given value.

### HasToastNotificationFlag

`func (o *MemberTemplate) HasToastNotificationFlag() bool`

HasToastNotificationFlag returns a boolean if a field has been set.

### SetToastNotificationFlagNil

`func (o *MemberTemplate) SetToastNotificationFlagNil(b bool)`

 SetToastNotificationFlagNil sets the value for ToastNotificationFlag to be an explicit nil

### UnsetToastNotificationFlag
`func (o *MemberTemplate) UnsetToastNotificationFlag()`

UnsetToastNotificationFlag ensures that no value is present for ToastNotificationFlag, not even an explicit nil
### GetMemberPersonas

`func (o *MemberTemplate) GetMemberPersonas() []int32`

GetMemberPersonas returns the MemberPersonas field if non-nil, zero value otherwise.

### GetMemberPersonasOk

`func (o *MemberTemplate) GetMemberPersonasOk() (*[]int32, bool)`

GetMemberPersonasOk returns a tuple with the MemberPersonas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPersonas

`func (o *MemberTemplate) SetMemberPersonas(v []int32)`

SetMemberPersonas sets MemberPersonas field to given value.

### HasMemberPersonas

`func (o *MemberTemplate) HasMemberPersonas() bool`

HasMemberPersonas returns a boolean if a field has been set.

### GetAdminFlag

`func (o *MemberTemplate) GetAdminFlag() bool`

GetAdminFlag returns the AdminFlag field if non-nil, zero value otherwise.

### GetAdminFlagOk

`func (o *MemberTemplate) GetAdminFlagOk() (*bool, bool)`

GetAdminFlagOk returns a tuple with the AdminFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFlag

`func (o *MemberTemplate) SetAdminFlag(v bool)`

SetAdminFlag sets AdminFlag field to given value.

### HasAdminFlag

`func (o *MemberTemplate) HasAdminFlag() bool`

HasAdminFlag returns a boolean if a field has been set.

### SetAdminFlagNil

`func (o *MemberTemplate) SetAdminFlagNil(b bool)`

 SetAdminFlagNil sets the value for AdminFlag to be an explicit nil

### UnsetAdminFlag
`func (o *MemberTemplate) UnsetAdminFlag()`

UnsetAdminFlag ensures that no value is present for AdminFlag, not even an explicit nil
### GetStructureLevel

`func (o *MemberTemplate) GetStructureLevel() StructureReference`

GetStructureLevel returns the StructureLevel field if non-nil, zero value otherwise.

### GetStructureLevelOk

`func (o *MemberTemplate) GetStructureLevelOk() (*StructureReference, bool)`

GetStructureLevelOk returns a tuple with the StructureLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureLevel

`func (o *MemberTemplate) SetStructureLevel(v StructureReference)`

SetStructureLevel sets StructureLevel field to given value.

### HasStructureLevel

`func (o *MemberTemplate) HasStructureLevel() bool`

HasStructureLevel returns a boolean if a field has been set.

### GetSecurityLocation

`func (o *MemberTemplate) GetSecurityLocation() SystemLocationReference`

GetSecurityLocation returns the SecurityLocation field if non-nil, zero value otherwise.

### GetSecurityLocationOk

`func (o *MemberTemplate) GetSecurityLocationOk() (*SystemLocationReference, bool)`

GetSecurityLocationOk returns a tuple with the SecurityLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLocation

`func (o *MemberTemplate) SetSecurityLocation(v SystemLocationReference)`

SetSecurityLocation sets SecurityLocation field to given value.

### HasSecurityLocation

`func (o *MemberTemplate) HasSecurityLocation() bool`

HasSecurityLocation returns a boolean if a field has been set.

### GetDefaultLocation

`func (o *MemberTemplate) GetDefaultLocation() SystemLocationReference`

GetDefaultLocation returns the DefaultLocation field if non-nil, zero value otherwise.

### GetDefaultLocationOk

`func (o *MemberTemplate) GetDefaultLocationOk() (*SystemLocationReference, bool)`

GetDefaultLocationOk returns a tuple with the DefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocation

`func (o *MemberTemplate) SetDefaultLocation(v SystemLocationReference)`

SetDefaultLocation sets DefaultLocation field to given value.

### HasDefaultLocation

`func (o *MemberTemplate) HasDefaultLocation() bool`

HasDefaultLocation returns a boolean if a field has been set.

### GetDefaultDepartment

`func (o *MemberTemplate) GetDefaultDepartment() SystemDepartmentReference`

GetDefaultDepartment returns the DefaultDepartment field if non-nil, zero value otherwise.

### GetDefaultDepartmentOk

`func (o *MemberTemplate) GetDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetDefaultDepartmentOk returns a tuple with the DefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDepartment

`func (o *MemberTemplate) SetDefaultDepartment(v SystemDepartmentReference)`

SetDefaultDepartment sets DefaultDepartment field to given value.

### HasDefaultDepartment

`func (o *MemberTemplate) HasDefaultDepartment() bool`

HasDefaultDepartment returns a boolean if a field has been set.

### GetReportsTo

`func (o *MemberTemplate) GetReportsTo() MemberReference`

GetReportsTo returns the ReportsTo field if non-nil, zero value otherwise.

### GetReportsToOk

`func (o *MemberTemplate) GetReportsToOk() (*MemberReference, bool)`

GetReportsToOk returns a tuple with the ReportsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsTo

`func (o *MemberTemplate) SetReportsTo(v MemberReference)`

SetReportsTo sets ReportsTo field to given value.

### HasReportsTo

`func (o *MemberTemplate) HasReportsTo() bool`

HasReportsTo returns a boolean if a field has been set.

### GetRestrictLocationFlag

`func (o *MemberTemplate) GetRestrictLocationFlag() bool`

GetRestrictLocationFlag returns the RestrictLocationFlag field if non-nil, zero value otherwise.

### GetRestrictLocationFlagOk

`func (o *MemberTemplate) GetRestrictLocationFlagOk() (*bool, bool)`

GetRestrictLocationFlagOk returns a tuple with the RestrictLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictLocationFlag

`func (o *MemberTemplate) SetRestrictLocationFlag(v bool)`

SetRestrictLocationFlag sets RestrictLocationFlag field to given value.

### HasRestrictLocationFlag

`func (o *MemberTemplate) HasRestrictLocationFlag() bool`

HasRestrictLocationFlag returns a boolean if a field has been set.

### SetRestrictLocationFlagNil

`func (o *MemberTemplate) SetRestrictLocationFlagNil(b bool)`

 SetRestrictLocationFlagNil sets the value for RestrictLocationFlag to be an explicit nil

### UnsetRestrictLocationFlag
`func (o *MemberTemplate) UnsetRestrictLocationFlag()`

UnsetRestrictLocationFlag ensures that no value is present for RestrictLocationFlag, not even an explicit nil
### GetRestrictDepartmentFlag

`func (o *MemberTemplate) GetRestrictDepartmentFlag() bool`

GetRestrictDepartmentFlag returns the RestrictDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictDepartmentFlagOk

`func (o *MemberTemplate) GetRestrictDepartmentFlagOk() (*bool, bool)`

GetRestrictDepartmentFlagOk returns a tuple with the RestrictDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDepartmentFlag

`func (o *MemberTemplate) SetRestrictDepartmentFlag(v bool)`

SetRestrictDepartmentFlag sets RestrictDepartmentFlag field to given value.

### HasRestrictDepartmentFlag

`func (o *MemberTemplate) HasRestrictDepartmentFlag() bool`

HasRestrictDepartmentFlag returns a boolean if a field has been set.

### SetRestrictDepartmentFlagNil

`func (o *MemberTemplate) SetRestrictDepartmentFlagNil(b bool)`

 SetRestrictDepartmentFlagNil sets the value for RestrictDepartmentFlag to be an explicit nil

### UnsetRestrictDepartmentFlag
`func (o *MemberTemplate) UnsetRestrictDepartmentFlag()`

UnsetRestrictDepartmentFlag ensures that no value is present for RestrictDepartmentFlag, not even an explicit nil
### GetWorkRole

`func (o *MemberTemplate) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *MemberTemplate) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *MemberTemplate) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *MemberTemplate) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *MemberTemplate) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *MemberTemplate) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *MemberTemplate) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *MemberTemplate) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetTimeApprover

`func (o *MemberTemplate) GetTimeApprover() MemberReference`

GetTimeApprover returns the TimeApprover field if non-nil, zero value otherwise.

### GetTimeApproverOk

`func (o *MemberTemplate) GetTimeApproverOk() (*MemberReference, bool)`

GetTimeApproverOk returns a tuple with the TimeApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeApprover

`func (o *MemberTemplate) SetTimeApprover(v MemberReference)`

SetTimeApprover sets TimeApprover field to given value.

### HasTimeApprover

`func (o *MemberTemplate) HasTimeApprover() bool`

HasTimeApprover returns a boolean if a field has been set.

### GetExpenseApprover

`func (o *MemberTemplate) GetExpenseApprover() MemberReference`

GetExpenseApprover returns the ExpenseApprover field if non-nil, zero value otherwise.

### GetExpenseApproverOk

`func (o *MemberTemplate) GetExpenseApproverOk() (*MemberReference, bool)`

GetExpenseApproverOk returns a tuple with the ExpenseApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApprover

`func (o *MemberTemplate) SetExpenseApprover(v MemberReference)`

SetExpenseApprover sets ExpenseApprover field to given value.

### HasExpenseApprover

`func (o *MemberTemplate) HasExpenseApprover() bool`

HasExpenseApprover returns a boolean if a field has been set.

### GetBillableForecast

`func (o *MemberTemplate) GetBillableForecast() float64`

GetBillableForecast returns the BillableForecast field if non-nil, zero value otherwise.

### GetBillableForecastOk

`func (o *MemberTemplate) GetBillableForecastOk() (*float64, bool)`

GetBillableForecastOk returns a tuple with the BillableForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableForecast

`func (o *MemberTemplate) SetBillableForecast(v float64)`

SetBillableForecast sets BillableForecast field to given value.

### HasBillableForecast

`func (o *MemberTemplate) HasBillableForecast() bool`

HasBillableForecast returns a boolean if a field has been set.

### SetBillableForecastNil

`func (o *MemberTemplate) SetBillableForecastNil(b bool)`

 SetBillableForecastNil sets the value for BillableForecast to be an explicit nil

### UnsetBillableForecast
`func (o *MemberTemplate) UnsetBillableForecast()`

UnsetBillableForecast ensures that no value is present for BillableForecast, not even an explicit nil
### GetDailyCapacity

`func (o *MemberTemplate) GetDailyCapacity() float64`

GetDailyCapacity returns the DailyCapacity field if non-nil, zero value otherwise.

### GetDailyCapacityOk

`func (o *MemberTemplate) GetDailyCapacityOk() (*float64, bool)`

GetDailyCapacityOk returns a tuple with the DailyCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyCapacity

`func (o *MemberTemplate) SetDailyCapacity(v float64)`

SetDailyCapacity sets DailyCapacity field to given value.

### HasDailyCapacity

`func (o *MemberTemplate) HasDailyCapacity() bool`

HasDailyCapacity returns a boolean if a field has been set.

### SetDailyCapacityNil

`func (o *MemberTemplate) SetDailyCapacityNil(b bool)`

 SetDailyCapacityNil sets the value for DailyCapacity to be an explicit nil

### UnsetDailyCapacity
`func (o *MemberTemplate) UnsetDailyCapacity()`

UnsetDailyCapacity ensures that no value is present for DailyCapacity, not even an explicit nil
### GetHourlyCost

`func (o *MemberTemplate) GetHourlyCost() float64`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *MemberTemplate) GetHourlyCostOk() (*float64, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *MemberTemplate) SetHourlyCost(v float64)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *MemberTemplate) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### SetHourlyCostNil

`func (o *MemberTemplate) SetHourlyCostNil(b bool)`

 SetHourlyCostNil sets the value for HourlyCost to be an explicit nil

### UnsetHourlyCost
`func (o *MemberTemplate) UnsetHourlyCost()`

UnsetHourlyCost ensures that no value is present for HourlyCost, not even an explicit nil
### GetHourlyRate

`func (o *MemberTemplate) GetHourlyRate() float64`

GetHourlyRate returns the HourlyRate field if non-nil, zero value otherwise.

### GetHourlyRateOk

`func (o *MemberTemplate) GetHourlyRateOk() (*float64, bool)`

GetHourlyRateOk returns a tuple with the HourlyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyRate

`func (o *MemberTemplate) SetHourlyRate(v float64)`

SetHourlyRate sets HourlyRate field to given value.

### HasHourlyRate

`func (o *MemberTemplate) HasHourlyRate() bool`

HasHourlyRate returns a boolean if a field has been set.

### SetHourlyRateNil

`func (o *MemberTemplate) SetHourlyRateNil(b bool)`

 SetHourlyRateNil sets the value for HourlyRate to be an explicit nil

### UnsetHourlyRate
`func (o *MemberTemplate) UnsetHourlyRate()`

UnsetHourlyRate ensures that no value is present for HourlyRate, not even an explicit nil
### GetIncludeInUtilizationReportingFlag

`func (o *MemberTemplate) GetIncludeInUtilizationReportingFlag() bool`

GetIncludeInUtilizationReportingFlag returns the IncludeInUtilizationReportingFlag field if non-nil, zero value otherwise.

### GetIncludeInUtilizationReportingFlagOk

`func (o *MemberTemplate) GetIncludeInUtilizationReportingFlagOk() (*bool, bool)`

GetIncludeInUtilizationReportingFlagOk returns a tuple with the IncludeInUtilizationReportingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInUtilizationReportingFlag

`func (o *MemberTemplate) SetIncludeInUtilizationReportingFlag(v bool)`

SetIncludeInUtilizationReportingFlag sets IncludeInUtilizationReportingFlag field to given value.

### HasIncludeInUtilizationReportingFlag

`func (o *MemberTemplate) HasIncludeInUtilizationReportingFlag() bool`

HasIncludeInUtilizationReportingFlag returns a boolean if a field has been set.

### SetIncludeInUtilizationReportingFlagNil

`func (o *MemberTemplate) SetIncludeInUtilizationReportingFlagNil(b bool)`

 SetIncludeInUtilizationReportingFlagNil sets the value for IncludeInUtilizationReportingFlag to be an explicit nil

### UnsetIncludeInUtilizationReportingFlag
`func (o *MemberTemplate) UnsetIncludeInUtilizationReportingFlag()`

UnsetIncludeInUtilizationReportingFlag ensures that no value is present for IncludeInUtilizationReportingFlag, not even an explicit nil
### GetRequireExpenseEntryFlag

`func (o *MemberTemplate) GetRequireExpenseEntryFlag() bool`

GetRequireExpenseEntryFlag returns the RequireExpenseEntryFlag field if non-nil, zero value otherwise.

### GetRequireExpenseEntryFlagOk

`func (o *MemberTemplate) GetRequireExpenseEntryFlagOk() (*bool, bool)`

GetRequireExpenseEntryFlagOk returns a tuple with the RequireExpenseEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExpenseEntryFlag

`func (o *MemberTemplate) SetRequireExpenseEntryFlag(v bool)`

SetRequireExpenseEntryFlag sets RequireExpenseEntryFlag field to given value.

### HasRequireExpenseEntryFlag

`func (o *MemberTemplate) HasRequireExpenseEntryFlag() bool`

HasRequireExpenseEntryFlag returns a boolean if a field has been set.

### SetRequireExpenseEntryFlagNil

`func (o *MemberTemplate) SetRequireExpenseEntryFlagNil(b bool)`

 SetRequireExpenseEntryFlagNil sets the value for RequireExpenseEntryFlag to be an explicit nil

### UnsetRequireExpenseEntryFlag
`func (o *MemberTemplate) UnsetRequireExpenseEntryFlag()`

UnsetRequireExpenseEntryFlag ensures that no value is present for RequireExpenseEntryFlag, not even an explicit nil
### GetRequireTimeSheetEntryFlag

`func (o *MemberTemplate) GetRequireTimeSheetEntryFlag() bool`

GetRequireTimeSheetEntryFlag returns the RequireTimeSheetEntryFlag field if non-nil, zero value otherwise.

### GetRequireTimeSheetEntryFlagOk

`func (o *MemberTemplate) GetRequireTimeSheetEntryFlagOk() (*bool, bool)`

GetRequireTimeSheetEntryFlagOk returns a tuple with the RequireTimeSheetEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTimeSheetEntryFlag

`func (o *MemberTemplate) SetRequireTimeSheetEntryFlag(v bool)`

SetRequireTimeSheetEntryFlag sets RequireTimeSheetEntryFlag field to given value.

### HasRequireTimeSheetEntryFlag

`func (o *MemberTemplate) HasRequireTimeSheetEntryFlag() bool`

HasRequireTimeSheetEntryFlag returns a boolean if a field has been set.

### SetRequireTimeSheetEntryFlagNil

`func (o *MemberTemplate) SetRequireTimeSheetEntryFlagNil(b bool)`

 SetRequireTimeSheetEntryFlagNil sets the value for RequireTimeSheetEntryFlag to be an explicit nil

### UnsetRequireTimeSheetEntryFlag
`func (o *MemberTemplate) UnsetRequireTimeSheetEntryFlag()`

UnsetRequireTimeSheetEntryFlag ensures that no value is present for RequireTimeSheetEntryFlag, not even an explicit nil
### GetRequireStartAndEndTimeOnTimeEntryFlag

`func (o *MemberTemplate) GetRequireStartAndEndTimeOnTimeEntryFlag() bool`

GetRequireStartAndEndTimeOnTimeEntryFlag returns the RequireStartAndEndTimeOnTimeEntryFlag field if non-nil, zero value otherwise.

### GetRequireStartAndEndTimeOnTimeEntryFlagOk

`func (o *MemberTemplate) GetRequireStartAndEndTimeOnTimeEntryFlagOk() (*bool, bool)`

GetRequireStartAndEndTimeOnTimeEntryFlagOk returns a tuple with the RequireStartAndEndTimeOnTimeEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireStartAndEndTimeOnTimeEntryFlag

`func (o *MemberTemplate) SetRequireStartAndEndTimeOnTimeEntryFlag(v bool)`

SetRequireStartAndEndTimeOnTimeEntryFlag sets RequireStartAndEndTimeOnTimeEntryFlag field to given value.

### HasRequireStartAndEndTimeOnTimeEntryFlag

`func (o *MemberTemplate) HasRequireStartAndEndTimeOnTimeEntryFlag() bool`

HasRequireStartAndEndTimeOnTimeEntryFlag returns a boolean if a field has been set.

### SetRequireStartAndEndTimeOnTimeEntryFlagNil

`func (o *MemberTemplate) SetRequireStartAndEndTimeOnTimeEntryFlagNil(b bool)`

 SetRequireStartAndEndTimeOnTimeEntryFlagNil sets the value for RequireStartAndEndTimeOnTimeEntryFlag to be an explicit nil

### UnsetRequireStartAndEndTimeOnTimeEntryFlag
`func (o *MemberTemplate) UnsetRequireStartAndEndTimeOnTimeEntryFlag()`

UnsetRequireStartAndEndTimeOnTimeEntryFlag ensures that no value is present for RequireStartAndEndTimeOnTimeEntryFlag, not even an explicit nil
### GetAllowInCellEntryOnTimeSheet

`func (o *MemberTemplate) GetAllowInCellEntryOnTimeSheet() bool`

GetAllowInCellEntryOnTimeSheet returns the AllowInCellEntryOnTimeSheet field if non-nil, zero value otherwise.

### GetAllowInCellEntryOnTimeSheetOk

`func (o *MemberTemplate) GetAllowInCellEntryOnTimeSheetOk() (*bool, bool)`

GetAllowInCellEntryOnTimeSheetOk returns a tuple with the AllowInCellEntryOnTimeSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInCellEntryOnTimeSheet

`func (o *MemberTemplate) SetAllowInCellEntryOnTimeSheet(v bool)`

SetAllowInCellEntryOnTimeSheet sets AllowInCellEntryOnTimeSheet field to given value.

### HasAllowInCellEntryOnTimeSheet

`func (o *MemberTemplate) HasAllowInCellEntryOnTimeSheet() bool`

HasAllowInCellEntryOnTimeSheet returns a boolean if a field has been set.

### SetAllowInCellEntryOnTimeSheetNil

`func (o *MemberTemplate) SetAllowInCellEntryOnTimeSheetNil(b bool)`

 SetAllowInCellEntryOnTimeSheetNil sets the value for AllowInCellEntryOnTimeSheet to be an explicit nil

### UnsetAllowInCellEntryOnTimeSheet
`func (o *MemberTemplate) UnsetAllowInCellEntryOnTimeSheet()`

UnsetAllowInCellEntryOnTimeSheet ensures that no value is present for AllowInCellEntryOnTimeSheet, not even an explicit nil
### GetEnterTimeAgainstCompanyFlag

`func (o *MemberTemplate) GetEnterTimeAgainstCompanyFlag() bool`

GetEnterTimeAgainstCompanyFlag returns the EnterTimeAgainstCompanyFlag field if non-nil, zero value otherwise.

### GetEnterTimeAgainstCompanyFlagOk

`func (o *MemberTemplate) GetEnterTimeAgainstCompanyFlagOk() (*bool, bool)`

GetEnterTimeAgainstCompanyFlagOk returns a tuple with the EnterTimeAgainstCompanyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterTimeAgainstCompanyFlag

`func (o *MemberTemplate) SetEnterTimeAgainstCompanyFlag(v bool)`

SetEnterTimeAgainstCompanyFlag sets EnterTimeAgainstCompanyFlag field to given value.

### HasEnterTimeAgainstCompanyFlag

`func (o *MemberTemplate) HasEnterTimeAgainstCompanyFlag() bool`

HasEnterTimeAgainstCompanyFlag returns a boolean if a field has been set.

### SetEnterTimeAgainstCompanyFlagNil

`func (o *MemberTemplate) SetEnterTimeAgainstCompanyFlagNil(b bool)`

 SetEnterTimeAgainstCompanyFlagNil sets the value for EnterTimeAgainstCompanyFlag to be an explicit nil

### UnsetEnterTimeAgainstCompanyFlag
`func (o *MemberTemplate) UnsetEnterTimeAgainstCompanyFlag()`

UnsetEnterTimeAgainstCompanyFlag ensures that no value is present for EnterTimeAgainstCompanyFlag, not even an explicit nil
### GetAllowExpensesEnteredAgainstCompaniesFlag

`func (o *MemberTemplate) GetAllowExpensesEnteredAgainstCompaniesFlag() bool`

GetAllowExpensesEnteredAgainstCompaniesFlag returns the AllowExpensesEnteredAgainstCompaniesFlag field if non-nil, zero value otherwise.

### GetAllowExpensesEnteredAgainstCompaniesFlagOk

`func (o *MemberTemplate) GetAllowExpensesEnteredAgainstCompaniesFlagOk() (*bool, bool)`

GetAllowExpensesEnteredAgainstCompaniesFlagOk returns a tuple with the AllowExpensesEnteredAgainstCompaniesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExpensesEnteredAgainstCompaniesFlag

`func (o *MemberTemplate) SetAllowExpensesEnteredAgainstCompaniesFlag(v bool)`

SetAllowExpensesEnteredAgainstCompaniesFlag sets AllowExpensesEnteredAgainstCompaniesFlag field to given value.

### HasAllowExpensesEnteredAgainstCompaniesFlag

`func (o *MemberTemplate) HasAllowExpensesEnteredAgainstCompaniesFlag() bool`

HasAllowExpensesEnteredAgainstCompaniesFlag returns a boolean if a field has been set.

### SetAllowExpensesEnteredAgainstCompaniesFlagNil

`func (o *MemberTemplate) SetAllowExpensesEnteredAgainstCompaniesFlagNil(b bool)`

 SetAllowExpensesEnteredAgainstCompaniesFlagNil sets the value for AllowExpensesEnteredAgainstCompaniesFlag to be an explicit nil

### UnsetAllowExpensesEnteredAgainstCompaniesFlag
`func (o *MemberTemplate) UnsetAllowExpensesEnteredAgainstCompaniesFlag()`

UnsetAllowExpensesEnteredAgainstCompaniesFlag ensures that no value is present for AllowExpensesEnteredAgainstCompaniesFlag, not even an explicit nil
### GetTimeReminderEmailFlag

`func (o *MemberTemplate) GetTimeReminderEmailFlag() bool`

GetTimeReminderEmailFlag returns the TimeReminderEmailFlag field if non-nil, zero value otherwise.

### GetTimeReminderEmailFlagOk

`func (o *MemberTemplate) GetTimeReminderEmailFlagOk() (*bool, bool)`

GetTimeReminderEmailFlagOk returns a tuple with the TimeReminderEmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeReminderEmailFlag

`func (o *MemberTemplate) SetTimeReminderEmailFlag(v bool)`

SetTimeReminderEmailFlag sets TimeReminderEmailFlag field to given value.

### HasTimeReminderEmailFlag

`func (o *MemberTemplate) HasTimeReminderEmailFlag() bool`

HasTimeReminderEmailFlag returns a boolean if a field has been set.

### SetTimeReminderEmailFlagNil

`func (o *MemberTemplate) SetTimeReminderEmailFlagNil(b bool)`

 SetTimeReminderEmailFlagNil sets the value for TimeReminderEmailFlag to be an explicit nil

### UnsetTimeReminderEmailFlag
`func (o *MemberTemplate) UnsetTimeReminderEmailFlag()`

UnsetTimeReminderEmailFlag ensures that no value is present for TimeReminderEmailFlag, not even an explicit nil
### GetDaysTolerance

`func (o *MemberTemplate) GetDaysTolerance() int32`

GetDaysTolerance returns the DaysTolerance field if non-nil, zero value otherwise.

### GetDaysToleranceOk

`func (o *MemberTemplate) GetDaysToleranceOk() (*int32, bool)`

GetDaysToleranceOk returns a tuple with the DaysTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysTolerance

`func (o *MemberTemplate) SetDaysTolerance(v int32)`

SetDaysTolerance sets DaysTolerance field to given value.

### HasDaysTolerance

`func (o *MemberTemplate) HasDaysTolerance() bool`

HasDaysTolerance returns a boolean if a field has been set.

### SetDaysToleranceNil

`func (o *MemberTemplate) SetDaysToleranceNil(b bool)`

 SetDaysToleranceNil sets the value for DaysTolerance to be an explicit nil

### UnsetDaysTolerance
`func (o *MemberTemplate) UnsetDaysTolerance()`

UnsetDaysTolerance ensures that no value is present for DaysTolerance, not even an explicit nil
### GetMinimumHours

`func (o *MemberTemplate) GetMinimumHours() float64`

GetMinimumHours returns the MinimumHours field if non-nil, zero value otherwise.

### GetMinimumHoursOk

`func (o *MemberTemplate) GetMinimumHoursOk() (*float64, bool)`

GetMinimumHoursOk returns a tuple with the MinimumHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumHours

`func (o *MemberTemplate) SetMinimumHours(v float64)`

SetMinimumHours sets MinimumHours field to given value.

### HasMinimumHours

`func (o *MemberTemplate) HasMinimumHours() bool`

HasMinimumHours returns a boolean if a field has been set.

### SetMinimumHoursNil

`func (o *MemberTemplate) SetMinimumHoursNil(b bool)`

 SetMinimumHoursNil sets the value for MinimumHours to be an explicit nil

### UnsetMinimumHours
`func (o *MemberTemplate) UnsetMinimumHours()`

UnsetMinimumHours ensures that no value is present for MinimumHours, not even an explicit nil
### GetTimeSheetStartDate

`func (o *MemberTemplate) GetTimeSheetStartDate() string`

GetTimeSheetStartDate returns the TimeSheetStartDate field if non-nil, zero value otherwise.

### GetTimeSheetStartDateOk

`func (o *MemberTemplate) GetTimeSheetStartDateOk() (*string, bool)`

GetTimeSheetStartDateOk returns a tuple with the TimeSheetStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSheetStartDate

`func (o *MemberTemplate) SetTimeSheetStartDate(v string)`

SetTimeSheetStartDate sets TimeSheetStartDate field to given value.

### HasTimeSheetStartDate

`func (o *MemberTemplate) HasTimeSheetStartDate() bool`

HasTimeSheetStartDate returns a boolean if a field has been set.

### GetServiceDefaultLocation

`func (o *MemberTemplate) GetServiceDefaultLocation() SystemLocationReference`

GetServiceDefaultLocation returns the ServiceDefaultLocation field if non-nil, zero value otherwise.

### GetServiceDefaultLocationOk

`func (o *MemberTemplate) GetServiceDefaultLocationOk() (*SystemLocationReference, bool)`

GetServiceDefaultLocationOk returns a tuple with the ServiceDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultLocation

`func (o *MemberTemplate) SetServiceDefaultLocation(v SystemLocationReference)`

SetServiceDefaultLocation sets ServiceDefaultLocation field to given value.

### HasServiceDefaultLocation

`func (o *MemberTemplate) HasServiceDefaultLocation() bool`

HasServiceDefaultLocation returns a boolean if a field has been set.

### GetServiceDefaultDepartment

`func (o *MemberTemplate) GetServiceDefaultDepartment() SystemDepartmentReference`

GetServiceDefaultDepartment returns the ServiceDefaultDepartment field if non-nil, zero value otherwise.

### GetServiceDefaultDepartmentOk

`func (o *MemberTemplate) GetServiceDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetServiceDefaultDepartmentOk returns a tuple with the ServiceDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultDepartment

`func (o *MemberTemplate) SetServiceDefaultDepartment(v SystemDepartmentReference)`

SetServiceDefaultDepartment sets ServiceDefaultDepartment field to given value.

### HasServiceDefaultDepartment

`func (o *MemberTemplate) HasServiceDefaultDepartment() bool`

HasServiceDefaultDepartment returns a boolean if a field has been set.

### GetServiceDefaultBoard

`func (o *MemberTemplate) GetServiceDefaultBoard() BoardReference`

GetServiceDefaultBoard returns the ServiceDefaultBoard field if non-nil, zero value otherwise.

### GetServiceDefaultBoardOk

`func (o *MemberTemplate) GetServiceDefaultBoardOk() (*BoardReference, bool)`

GetServiceDefaultBoardOk returns a tuple with the ServiceDefaultBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultBoard

`func (o *MemberTemplate) SetServiceDefaultBoard(v BoardReference)`

SetServiceDefaultBoard sets ServiceDefaultBoard field to given value.

### HasServiceDefaultBoard

`func (o *MemberTemplate) HasServiceDefaultBoard() bool`

HasServiceDefaultBoard returns a boolean if a field has been set.

### GetRestrictServiceDefaultLocationFlag

`func (o *MemberTemplate) GetRestrictServiceDefaultLocationFlag() bool`

GetRestrictServiceDefaultLocationFlag returns the RestrictServiceDefaultLocationFlag field if non-nil, zero value otherwise.

### GetRestrictServiceDefaultLocationFlagOk

`func (o *MemberTemplate) GetRestrictServiceDefaultLocationFlagOk() (*bool, bool)`

GetRestrictServiceDefaultLocationFlagOk returns a tuple with the RestrictServiceDefaultLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictServiceDefaultLocationFlag

`func (o *MemberTemplate) SetRestrictServiceDefaultLocationFlag(v bool)`

SetRestrictServiceDefaultLocationFlag sets RestrictServiceDefaultLocationFlag field to given value.

### HasRestrictServiceDefaultLocationFlag

`func (o *MemberTemplate) HasRestrictServiceDefaultLocationFlag() bool`

HasRestrictServiceDefaultLocationFlag returns a boolean if a field has been set.

### SetRestrictServiceDefaultLocationFlagNil

`func (o *MemberTemplate) SetRestrictServiceDefaultLocationFlagNil(b bool)`

 SetRestrictServiceDefaultLocationFlagNil sets the value for RestrictServiceDefaultLocationFlag to be an explicit nil

### UnsetRestrictServiceDefaultLocationFlag
`func (o *MemberTemplate) UnsetRestrictServiceDefaultLocationFlag()`

UnsetRestrictServiceDefaultLocationFlag ensures that no value is present for RestrictServiceDefaultLocationFlag, not even an explicit nil
### GetRestrictServiceDefaultDepartmentFlag

`func (o *MemberTemplate) GetRestrictServiceDefaultDepartmentFlag() bool`

GetRestrictServiceDefaultDepartmentFlag returns the RestrictServiceDefaultDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictServiceDefaultDepartmentFlagOk

`func (o *MemberTemplate) GetRestrictServiceDefaultDepartmentFlagOk() (*bool, bool)`

GetRestrictServiceDefaultDepartmentFlagOk returns a tuple with the RestrictServiceDefaultDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictServiceDefaultDepartmentFlag

`func (o *MemberTemplate) SetRestrictServiceDefaultDepartmentFlag(v bool)`

SetRestrictServiceDefaultDepartmentFlag sets RestrictServiceDefaultDepartmentFlag field to given value.

### HasRestrictServiceDefaultDepartmentFlag

`func (o *MemberTemplate) HasRestrictServiceDefaultDepartmentFlag() bool`

HasRestrictServiceDefaultDepartmentFlag returns a boolean if a field has been set.

### SetRestrictServiceDefaultDepartmentFlagNil

`func (o *MemberTemplate) SetRestrictServiceDefaultDepartmentFlagNil(b bool)`

 SetRestrictServiceDefaultDepartmentFlagNil sets the value for RestrictServiceDefaultDepartmentFlag to be an explicit nil

### UnsetRestrictServiceDefaultDepartmentFlag
`func (o *MemberTemplate) UnsetRestrictServiceDefaultDepartmentFlag()`

UnsetRestrictServiceDefaultDepartmentFlag ensures that no value is present for RestrictServiceDefaultDepartmentFlag, not even an explicit nil
### GetExcludedServiceBoardIds

`func (o *MemberTemplate) GetExcludedServiceBoardIds() []int32`

GetExcludedServiceBoardIds returns the ExcludedServiceBoardIds field if non-nil, zero value otherwise.

### GetExcludedServiceBoardIdsOk

`func (o *MemberTemplate) GetExcludedServiceBoardIdsOk() (*[]int32, bool)`

GetExcludedServiceBoardIdsOk returns a tuple with the ExcludedServiceBoardIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedServiceBoardIds

`func (o *MemberTemplate) SetExcludedServiceBoardIds(v []int32)`

SetExcludedServiceBoardIds sets ExcludedServiceBoardIds field to given value.

### HasExcludedServiceBoardIds

`func (o *MemberTemplate) HasExcludedServiceBoardIds() bool`

HasExcludedServiceBoardIds returns a boolean if a field has been set.

### GetTeams

`func (o *MemberTemplate) GetTeams() []int32`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *MemberTemplate) GetTeamsOk() (*[]int32, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *MemberTemplate) SetTeams(v []int32)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *MemberTemplate) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetServiceBoardTeamIds

`func (o *MemberTemplate) GetServiceBoardTeamIds() []int32`

GetServiceBoardTeamIds returns the ServiceBoardTeamIds field if non-nil, zero value otherwise.

### GetServiceBoardTeamIdsOk

`func (o *MemberTemplate) GetServiceBoardTeamIdsOk() (*[]int32, bool)`

GetServiceBoardTeamIdsOk returns a tuple with the ServiceBoardTeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBoardTeamIds

`func (o *MemberTemplate) SetServiceBoardTeamIds(v []int32)`

SetServiceBoardTeamIds sets ServiceBoardTeamIds field to given value.

### HasServiceBoardTeamIds

`func (o *MemberTemplate) HasServiceBoardTeamIds() bool`

HasServiceBoardTeamIds returns a boolean if a field has been set.

### GetProjectDefaultLocation

`func (o *MemberTemplate) GetProjectDefaultLocation() SystemLocationReference`

GetProjectDefaultLocation returns the ProjectDefaultLocation field if non-nil, zero value otherwise.

### GetProjectDefaultLocationOk

`func (o *MemberTemplate) GetProjectDefaultLocationOk() (*SystemLocationReference, bool)`

GetProjectDefaultLocationOk returns a tuple with the ProjectDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultLocation

`func (o *MemberTemplate) SetProjectDefaultLocation(v SystemLocationReference)`

SetProjectDefaultLocation sets ProjectDefaultLocation field to given value.

### HasProjectDefaultLocation

`func (o *MemberTemplate) HasProjectDefaultLocation() bool`

HasProjectDefaultLocation returns a boolean if a field has been set.

### GetProjectDefaultDepartment

`func (o *MemberTemplate) GetProjectDefaultDepartment() SystemDepartmentReference`

GetProjectDefaultDepartment returns the ProjectDefaultDepartment field if non-nil, zero value otherwise.

### GetProjectDefaultDepartmentOk

`func (o *MemberTemplate) GetProjectDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetProjectDefaultDepartmentOk returns a tuple with the ProjectDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultDepartment

`func (o *MemberTemplate) SetProjectDefaultDepartment(v SystemDepartmentReference)`

SetProjectDefaultDepartment sets ProjectDefaultDepartment field to given value.

### HasProjectDefaultDepartment

`func (o *MemberTemplate) HasProjectDefaultDepartment() bool`

HasProjectDefaultDepartment returns a boolean if a field has been set.

### GetProjectDefaultBoard

`func (o *MemberTemplate) GetProjectDefaultBoard() ProjectBoardReference`

GetProjectDefaultBoard returns the ProjectDefaultBoard field if non-nil, zero value otherwise.

### GetProjectDefaultBoardOk

`func (o *MemberTemplate) GetProjectDefaultBoardOk() (*ProjectBoardReference, bool)`

GetProjectDefaultBoardOk returns a tuple with the ProjectDefaultBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDefaultBoard

`func (o *MemberTemplate) SetProjectDefaultBoard(v ProjectBoardReference)`

SetProjectDefaultBoard sets ProjectDefaultBoard field to given value.

### HasProjectDefaultBoard

`func (o *MemberTemplate) HasProjectDefaultBoard() bool`

HasProjectDefaultBoard returns a boolean if a field has been set.

### GetRestrictProjectDefaultLocationFlag

`func (o *MemberTemplate) GetRestrictProjectDefaultLocationFlag() bool`

GetRestrictProjectDefaultLocationFlag returns the RestrictProjectDefaultLocationFlag field if non-nil, zero value otherwise.

### GetRestrictProjectDefaultLocationFlagOk

`func (o *MemberTemplate) GetRestrictProjectDefaultLocationFlagOk() (*bool, bool)`

GetRestrictProjectDefaultLocationFlagOk returns a tuple with the RestrictProjectDefaultLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictProjectDefaultLocationFlag

`func (o *MemberTemplate) SetRestrictProjectDefaultLocationFlag(v bool)`

SetRestrictProjectDefaultLocationFlag sets RestrictProjectDefaultLocationFlag field to given value.

### HasRestrictProjectDefaultLocationFlag

`func (o *MemberTemplate) HasRestrictProjectDefaultLocationFlag() bool`

HasRestrictProjectDefaultLocationFlag returns a boolean if a field has been set.

### SetRestrictProjectDefaultLocationFlagNil

`func (o *MemberTemplate) SetRestrictProjectDefaultLocationFlagNil(b bool)`

 SetRestrictProjectDefaultLocationFlagNil sets the value for RestrictProjectDefaultLocationFlag to be an explicit nil

### UnsetRestrictProjectDefaultLocationFlag
`func (o *MemberTemplate) UnsetRestrictProjectDefaultLocationFlag()`

UnsetRestrictProjectDefaultLocationFlag ensures that no value is present for RestrictProjectDefaultLocationFlag, not even an explicit nil
### GetRestrictProjectDefaultDepartmentFlag

`func (o *MemberTemplate) GetRestrictProjectDefaultDepartmentFlag() bool`

GetRestrictProjectDefaultDepartmentFlag returns the RestrictProjectDefaultDepartmentFlag field if non-nil, zero value otherwise.

### GetRestrictProjectDefaultDepartmentFlagOk

`func (o *MemberTemplate) GetRestrictProjectDefaultDepartmentFlagOk() (*bool, bool)`

GetRestrictProjectDefaultDepartmentFlagOk returns a tuple with the RestrictProjectDefaultDepartmentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictProjectDefaultDepartmentFlag

`func (o *MemberTemplate) SetRestrictProjectDefaultDepartmentFlag(v bool)`

SetRestrictProjectDefaultDepartmentFlag sets RestrictProjectDefaultDepartmentFlag field to given value.

### HasRestrictProjectDefaultDepartmentFlag

`func (o *MemberTemplate) HasRestrictProjectDefaultDepartmentFlag() bool`

HasRestrictProjectDefaultDepartmentFlag returns a boolean if a field has been set.

### SetRestrictProjectDefaultDepartmentFlagNil

`func (o *MemberTemplate) SetRestrictProjectDefaultDepartmentFlagNil(b bool)`

 SetRestrictProjectDefaultDepartmentFlagNil sets the value for RestrictProjectDefaultDepartmentFlag to be an explicit nil

### UnsetRestrictProjectDefaultDepartmentFlag
`func (o *MemberTemplate) UnsetRestrictProjectDefaultDepartmentFlag()`

UnsetRestrictProjectDefaultDepartmentFlag ensures that no value is present for RestrictProjectDefaultDepartmentFlag, not even an explicit nil
### GetExcludedProjectBoardIds

`func (o *MemberTemplate) GetExcludedProjectBoardIds() []int32`

GetExcludedProjectBoardIds returns the ExcludedProjectBoardIds field if non-nil, zero value otherwise.

### GetExcludedProjectBoardIdsOk

`func (o *MemberTemplate) GetExcludedProjectBoardIdsOk() (*[]int32, bool)`

GetExcludedProjectBoardIdsOk returns a tuple with the ExcludedProjectBoardIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProjectBoardIds

`func (o *MemberTemplate) SetExcludedProjectBoardIds(v []int32)`

SetExcludedProjectBoardIds sets ExcludedProjectBoardIds field to given value.

### HasExcludedProjectBoardIds

`func (o *MemberTemplate) HasExcludedProjectBoardIds() bool`

HasExcludedProjectBoardIds returns a boolean if a field has been set.

### GetScheduleDefaultLocation

`func (o *MemberTemplate) GetScheduleDefaultLocation() SystemLocationReference`

GetScheduleDefaultLocation returns the ScheduleDefaultLocation field if non-nil, zero value otherwise.

### GetScheduleDefaultLocationOk

`func (o *MemberTemplate) GetScheduleDefaultLocationOk() (*SystemLocationReference, bool)`

GetScheduleDefaultLocationOk returns a tuple with the ScheduleDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDefaultLocation

`func (o *MemberTemplate) SetScheduleDefaultLocation(v SystemLocationReference)`

SetScheduleDefaultLocation sets ScheduleDefaultLocation field to given value.

### HasScheduleDefaultLocation

`func (o *MemberTemplate) HasScheduleDefaultLocation() bool`

HasScheduleDefaultLocation returns a boolean if a field has been set.

### GetScheduleDefaultDepartment

`func (o *MemberTemplate) GetScheduleDefaultDepartment() SystemDepartmentReference`

GetScheduleDefaultDepartment returns the ScheduleDefaultDepartment field if non-nil, zero value otherwise.

### GetScheduleDefaultDepartmentOk

`func (o *MemberTemplate) GetScheduleDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetScheduleDefaultDepartmentOk returns a tuple with the ScheduleDefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDefaultDepartment

`func (o *MemberTemplate) SetScheduleDefaultDepartment(v SystemDepartmentReference)`

SetScheduleDefaultDepartment sets ScheduleDefaultDepartment field to given value.

### HasScheduleDefaultDepartment

`func (o *MemberTemplate) HasScheduleDefaultDepartment() bool`

HasScheduleDefaultDepartment returns a boolean if a field has been set.

### GetScheduleCapacity

`func (o *MemberTemplate) GetScheduleCapacity() float64`

GetScheduleCapacity returns the ScheduleCapacity field if non-nil, zero value otherwise.

### GetScheduleCapacityOk

`func (o *MemberTemplate) GetScheduleCapacityOk() (*float64, bool)`

GetScheduleCapacityOk returns a tuple with the ScheduleCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCapacity

`func (o *MemberTemplate) SetScheduleCapacity(v float64)`

SetScheduleCapacity sets ScheduleCapacity field to given value.

### HasScheduleCapacity

`func (o *MemberTemplate) HasScheduleCapacity() bool`

HasScheduleCapacity returns a boolean if a field has been set.

### SetScheduleCapacityNil

`func (o *MemberTemplate) SetScheduleCapacityNil(b bool)`

 SetScheduleCapacityNil sets the value for ScheduleCapacity to be an explicit nil

### UnsetScheduleCapacity
`func (o *MemberTemplate) UnsetScheduleCapacity()`

UnsetScheduleCapacity ensures that no value is present for ScheduleCapacity, not even an explicit nil
### GetServiceLocation

`func (o *MemberTemplate) GetServiceLocation() ServiceLocationReference`

GetServiceLocation returns the ServiceLocation field if non-nil, zero value otherwise.

### GetServiceLocationOk

`func (o *MemberTemplate) GetServiceLocationOk() (*ServiceLocationReference, bool)`

GetServiceLocationOk returns a tuple with the ServiceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLocation

`func (o *MemberTemplate) SetServiceLocation(v ServiceLocationReference)`

SetServiceLocation sets ServiceLocation field to given value.

### HasServiceLocation

`func (o *MemberTemplate) HasServiceLocation() bool`

HasServiceLocation returns a boolean if a field has been set.

### GetRestrictScheduleFlag

`func (o *MemberTemplate) GetRestrictScheduleFlag() bool`

GetRestrictScheduleFlag returns the RestrictScheduleFlag field if non-nil, zero value otherwise.

### GetRestrictScheduleFlagOk

`func (o *MemberTemplate) GetRestrictScheduleFlagOk() (*bool, bool)`

GetRestrictScheduleFlagOk returns a tuple with the RestrictScheduleFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictScheduleFlag

`func (o *MemberTemplate) SetRestrictScheduleFlag(v bool)`

SetRestrictScheduleFlag sets RestrictScheduleFlag field to given value.

### HasRestrictScheduleFlag

`func (o *MemberTemplate) HasRestrictScheduleFlag() bool`

HasRestrictScheduleFlag returns a boolean if a field has been set.

### SetRestrictScheduleFlagNil

`func (o *MemberTemplate) SetRestrictScheduleFlagNil(b bool)`

 SetRestrictScheduleFlagNil sets the value for RestrictScheduleFlag to be an explicit nil

### UnsetRestrictScheduleFlag
`func (o *MemberTemplate) UnsetRestrictScheduleFlag()`

UnsetRestrictScheduleFlag ensures that no value is present for RestrictScheduleFlag, not even an explicit nil
### GetHideMemberInDispatchPortalFlag

`func (o *MemberTemplate) GetHideMemberInDispatchPortalFlag() bool`

GetHideMemberInDispatchPortalFlag returns the HideMemberInDispatchPortalFlag field if non-nil, zero value otherwise.

### GetHideMemberInDispatchPortalFlagOk

`func (o *MemberTemplate) GetHideMemberInDispatchPortalFlagOk() (*bool, bool)`

GetHideMemberInDispatchPortalFlagOk returns a tuple with the HideMemberInDispatchPortalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideMemberInDispatchPortalFlag

`func (o *MemberTemplate) SetHideMemberInDispatchPortalFlag(v bool)`

SetHideMemberInDispatchPortalFlag sets HideMemberInDispatchPortalFlag field to given value.

### HasHideMemberInDispatchPortalFlag

`func (o *MemberTemplate) HasHideMemberInDispatchPortalFlag() bool`

HasHideMemberInDispatchPortalFlag returns a boolean if a field has been set.

### SetHideMemberInDispatchPortalFlagNil

`func (o *MemberTemplate) SetHideMemberInDispatchPortalFlagNil(b bool)`

 SetHideMemberInDispatchPortalFlagNil sets the value for HideMemberInDispatchPortalFlag to be an explicit nil

### UnsetHideMemberInDispatchPortalFlag
`func (o *MemberTemplate) UnsetHideMemberInDispatchPortalFlag()`

UnsetHideMemberInDispatchPortalFlag ensures that no value is present for HideMemberInDispatchPortalFlag, not even an explicit nil
### GetCalendar

`func (o *MemberTemplate) GetCalendar() CalendarReference`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *MemberTemplate) GetCalendarOk() (*CalendarReference, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *MemberTemplate) SetCalendar(v CalendarReference)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *MemberTemplate) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### GetSalesDefaultLocation

`func (o *MemberTemplate) GetSalesDefaultLocation() SystemLocationReference`

GetSalesDefaultLocation returns the SalesDefaultLocation field if non-nil, zero value otherwise.

### GetSalesDefaultLocationOk

`func (o *MemberTemplate) GetSalesDefaultLocationOk() (*SystemLocationReference, bool)`

GetSalesDefaultLocationOk returns a tuple with the SalesDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDefaultLocation

`func (o *MemberTemplate) SetSalesDefaultLocation(v SystemLocationReference)`

SetSalesDefaultLocation sets SalesDefaultLocation field to given value.

### HasSalesDefaultLocation

`func (o *MemberTemplate) HasSalesDefaultLocation() bool`

HasSalesDefaultLocation returns a boolean if a field has been set.

### GetRestrictDefaultSalesTerritoryFlag

`func (o *MemberTemplate) GetRestrictDefaultSalesTerritoryFlag() bool`

GetRestrictDefaultSalesTerritoryFlag returns the RestrictDefaultSalesTerritoryFlag field if non-nil, zero value otherwise.

### GetRestrictDefaultSalesTerritoryFlagOk

`func (o *MemberTemplate) GetRestrictDefaultSalesTerritoryFlagOk() (*bool, bool)`

GetRestrictDefaultSalesTerritoryFlagOk returns a tuple with the RestrictDefaultSalesTerritoryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDefaultSalesTerritoryFlag

`func (o *MemberTemplate) SetRestrictDefaultSalesTerritoryFlag(v bool)`

SetRestrictDefaultSalesTerritoryFlag sets RestrictDefaultSalesTerritoryFlag field to given value.

### HasRestrictDefaultSalesTerritoryFlag

`func (o *MemberTemplate) HasRestrictDefaultSalesTerritoryFlag() bool`

HasRestrictDefaultSalesTerritoryFlag returns a boolean if a field has been set.

### SetRestrictDefaultSalesTerritoryFlagNil

`func (o *MemberTemplate) SetRestrictDefaultSalesTerritoryFlagNil(b bool)`

 SetRestrictDefaultSalesTerritoryFlagNil sets the value for RestrictDefaultSalesTerritoryFlag to be an explicit nil

### UnsetRestrictDefaultSalesTerritoryFlag
`func (o *MemberTemplate) UnsetRestrictDefaultSalesTerritoryFlag()`

UnsetRestrictDefaultSalesTerritoryFlag ensures that no value is present for RestrictDefaultSalesTerritoryFlag, not even an explicit nil
### GetWarehouse

`func (o *MemberTemplate) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *MemberTemplate) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *MemberTemplate) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *MemberTemplate) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetWarehouseBin

`func (o *MemberTemplate) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *MemberTemplate) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *MemberTemplate) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.

### HasWarehouseBin

`func (o *MemberTemplate) HasWarehouseBin() bool`

HasWarehouseBin returns a boolean if a field has been set.

### GetRestrictDefaultWarehouseFlag

`func (o *MemberTemplate) GetRestrictDefaultWarehouseFlag() bool`

GetRestrictDefaultWarehouseFlag returns the RestrictDefaultWarehouseFlag field if non-nil, zero value otherwise.

### GetRestrictDefaultWarehouseFlagOk

`func (o *MemberTemplate) GetRestrictDefaultWarehouseFlagOk() (*bool, bool)`

GetRestrictDefaultWarehouseFlagOk returns a tuple with the RestrictDefaultWarehouseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDefaultWarehouseFlag

`func (o *MemberTemplate) SetRestrictDefaultWarehouseFlag(v bool)`

SetRestrictDefaultWarehouseFlag sets RestrictDefaultWarehouseFlag field to given value.

### HasRestrictDefaultWarehouseFlag

`func (o *MemberTemplate) HasRestrictDefaultWarehouseFlag() bool`

HasRestrictDefaultWarehouseFlag returns a boolean if a field has been set.

### SetRestrictDefaultWarehouseFlagNil

`func (o *MemberTemplate) SetRestrictDefaultWarehouseFlagNil(b bool)`

 SetRestrictDefaultWarehouseFlagNil sets the value for RestrictDefaultWarehouseFlag to be an explicit nil

### UnsetRestrictDefaultWarehouseFlag
`func (o *MemberTemplate) UnsetRestrictDefaultWarehouseFlag()`

UnsetRestrictDefaultWarehouseFlag ensures that no value is present for RestrictDefaultWarehouseFlag, not even an explicit nil
### GetRestrictDefaultWarehouseBinFlag

`func (o *MemberTemplate) GetRestrictDefaultWarehouseBinFlag() bool`

GetRestrictDefaultWarehouseBinFlag returns the RestrictDefaultWarehouseBinFlag field if non-nil, zero value otherwise.

### GetRestrictDefaultWarehouseBinFlagOk

`func (o *MemberTemplate) GetRestrictDefaultWarehouseBinFlagOk() (*bool, bool)`

GetRestrictDefaultWarehouseBinFlagOk returns a tuple with the RestrictDefaultWarehouseBinFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDefaultWarehouseBinFlag

`func (o *MemberTemplate) SetRestrictDefaultWarehouseBinFlag(v bool)`

SetRestrictDefaultWarehouseBinFlag sets RestrictDefaultWarehouseBinFlag field to given value.

### HasRestrictDefaultWarehouseBinFlag

`func (o *MemberTemplate) HasRestrictDefaultWarehouseBinFlag() bool`

HasRestrictDefaultWarehouseBinFlag returns a boolean if a field has been set.

### SetRestrictDefaultWarehouseBinFlagNil

`func (o *MemberTemplate) SetRestrictDefaultWarehouseBinFlagNil(b bool)`

 SetRestrictDefaultWarehouseBinFlagNil sets the value for RestrictDefaultWarehouseBinFlag to be an explicit nil

### UnsetRestrictDefaultWarehouseBinFlag
`func (o *MemberTemplate) UnsetRestrictDefaultWarehouseBinFlag()`

UnsetRestrictDefaultWarehouseBinFlag ensures that no value is present for RestrictDefaultWarehouseBinFlag, not even an explicit nil
### GetCompanyActivityTabFormat

`func (o *MemberTemplate) GetCompanyActivityTabFormat() string`

GetCompanyActivityTabFormat returns the CompanyActivityTabFormat field if non-nil, zero value otherwise.

### GetCompanyActivityTabFormatOk

`func (o *MemberTemplate) GetCompanyActivityTabFormatOk() (*string, bool)`

GetCompanyActivityTabFormatOk returns a tuple with the CompanyActivityTabFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyActivityTabFormat

`func (o *MemberTemplate) SetCompanyActivityTabFormat(v string)`

SetCompanyActivityTabFormat sets CompanyActivityTabFormat field to given value.

### HasCompanyActivityTabFormat

`func (o *MemberTemplate) HasCompanyActivityTabFormat() bool`

HasCompanyActivityTabFormat returns a boolean if a field has been set.

### SetCompanyActivityTabFormatNil

`func (o *MemberTemplate) SetCompanyActivityTabFormatNil(b bool)`

 SetCompanyActivityTabFormatNil sets the value for CompanyActivityTabFormat to be an explicit nil

### UnsetCompanyActivityTabFormat
`func (o *MemberTemplate) UnsetCompanyActivityTabFormat()`

UnsetCompanyActivityTabFormat ensures that no value is present for CompanyActivityTabFormat, not even an explicit nil
### GetInvoiceTimeTabFormat

`func (o *MemberTemplate) GetInvoiceTimeTabFormat() string`

GetInvoiceTimeTabFormat returns the InvoiceTimeTabFormat field if non-nil, zero value otherwise.

### GetInvoiceTimeTabFormatOk

`func (o *MemberTemplate) GetInvoiceTimeTabFormatOk() (*string, bool)`

GetInvoiceTimeTabFormatOk returns a tuple with the InvoiceTimeTabFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTimeTabFormat

`func (o *MemberTemplate) SetInvoiceTimeTabFormat(v string)`

SetInvoiceTimeTabFormat sets InvoiceTimeTabFormat field to given value.

### HasInvoiceTimeTabFormat

`func (o *MemberTemplate) HasInvoiceTimeTabFormat() bool`

HasInvoiceTimeTabFormat returns a boolean if a field has been set.

### SetInvoiceTimeTabFormatNil

`func (o *MemberTemplate) SetInvoiceTimeTabFormatNil(b bool)`

 SetInvoiceTimeTabFormatNil sets the value for InvoiceTimeTabFormat to be an explicit nil

### UnsetInvoiceTimeTabFormat
`func (o *MemberTemplate) UnsetInvoiceTimeTabFormat()`

UnsetInvoiceTimeTabFormat ensures that no value is present for InvoiceTimeTabFormat, not even an explicit nil
### GetInvoiceScreenDefaultTabFormat

`func (o *MemberTemplate) GetInvoiceScreenDefaultTabFormat() string`

GetInvoiceScreenDefaultTabFormat returns the InvoiceScreenDefaultTabFormat field if non-nil, zero value otherwise.

### GetInvoiceScreenDefaultTabFormatOk

`func (o *MemberTemplate) GetInvoiceScreenDefaultTabFormatOk() (*string, bool)`

GetInvoiceScreenDefaultTabFormatOk returns a tuple with the InvoiceScreenDefaultTabFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceScreenDefaultTabFormat

`func (o *MemberTemplate) SetInvoiceScreenDefaultTabFormat(v string)`

SetInvoiceScreenDefaultTabFormat sets InvoiceScreenDefaultTabFormat field to given value.

### HasInvoiceScreenDefaultTabFormat

`func (o *MemberTemplate) HasInvoiceScreenDefaultTabFormat() bool`

HasInvoiceScreenDefaultTabFormat returns a boolean if a field has been set.

### SetInvoiceScreenDefaultTabFormatNil

`func (o *MemberTemplate) SetInvoiceScreenDefaultTabFormatNil(b bool)`

 SetInvoiceScreenDefaultTabFormatNil sets the value for InvoiceScreenDefaultTabFormat to be an explicit nil

### UnsetInvoiceScreenDefaultTabFormat
`func (o *MemberTemplate) UnsetInvoiceScreenDefaultTabFormat()`

UnsetInvoiceScreenDefaultTabFormat ensures that no value is present for InvoiceScreenDefaultTabFormat, not even an explicit nil
### GetInvoicingDisplayOptions

`func (o *MemberTemplate) GetInvoicingDisplayOptions() string`

GetInvoicingDisplayOptions returns the InvoicingDisplayOptions field if non-nil, zero value otherwise.

### GetInvoicingDisplayOptionsOk

`func (o *MemberTemplate) GetInvoicingDisplayOptionsOk() (*string, bool)`

GetInvoicingDisplayOptionsOk returns a tuple with the InvoicingDisplayOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingDisplayOptions

`func (o *MemberTemplate) SetInvoicingDisplayOptions(v string)`

SetInvoicingDisplayOptions sets InvoicingDisplayOptions field to given value.

### HasInvoicingDisplayOptions

`func (o *MemberTemplate) HasInvoicingDisplayOptions() bool`

HasInvoicingDisplayOptions returns a boolean if a field has been set.

### SetInvoicingDisplayOptionsNil

`func (o *MemberTemplate) SetInvoicingDisplayOptionsNil(b bool)`

 SetInvoicingDisplayOptionsNil sets the value for InvoicingDisplayOptions to be an explicit nil

### UnsetInvoicingDisplayOptions
`func (o *MemberTemplate) UnsetInvoicingDisplayOptions()`

UnsetInvoicingDisplayOptions ensures that no value is present for InvoicingDisplayOptions, not even an explicit nil
### GetAgreementInvoicingDisplayOptions

`func (o *MemberTemplate) GetAgreementInvoicingDisplayOptions() string`

GetAgreementInvoicingDisplayOptions returns the AgreementInvoicingDisplayOptions field if non-nil, zero value otherwise.

### GetAgreementInvoicingDisplayOptionsOk

`func (o *MemberTemplate) GetAgreementInvoicingDisplayOptionsOk() (*string, bool)`

GetAgreementInvoicingDisplayOptionsOk returns a tuple with the AgreementInvoicingDisplayOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementInvoicingDisplayOptions

`func (o *MemberTemplate) SetAgreementInvoicingDisplayOptions(v string)`

SetAgreementInvoicingDisplayOptions sets AgreementInvoicingDisplayOptions field to given value.

### HasAgreementInvoicingDisplayOptions

`func (o *MemberTemplate) HasAgreementInvoicingDisplayOptions() bool`

HasAgreementInvoicingDisplayOptions returns a boolean if a field has been set.

### SetAgreementInvoicingDisplayOptionsNil

`func (o *MemberTemplate) SetAgreementInvoicingDisplayOptionsNil(b bool)`

 SetAgreementInvoicingDisplayOptionsNil sets the value for AgreementInvoicingDisplayOptions to be an explicit nil

### UnsetAgreementInvoicingDisplayOptions
`func (o *MemberTemplate) UnsetAgreementInvoicingDisplayOptions()`

UnsetAgreementInvoicingDisplayOptions ensures that no value is present for AgreementInvoicingDisplayOptions, not even an explicit nil
### GetAutoStartStopwatch

`func (o *MemberTemplate) GetAutoStartStopwatch() bool`

GetAutoStartStopwatch returns the AutoStartStopwatch field if non-nil, zero value otherwise.

### GetAutoStartStopwatchOk

`func (o *MemberTemplate) GetAutoStartStopwatchOk() (*bool, bool)`

GetAutoStartStopwatchOk returns a tuple with the AutoStartStopwatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStartStopwatch

`func (o *MemberTemplate) SetAutoStartStopwatch(v bool)`

SetAutoStartStopwatch sets AutoStartStopwatch field to given value.

### HasAutoStartStopwatch

`func (o *MemberTemplate) HasAutoStartStopwatch() bool`

HasAutoStartStopwatch returns a boolean if a field has been set.

### SetAutoStartStopwatchNil

`func (o *MemberTemplate) SetAutoStartStopwatchNil(b bool)`

 SetAutoStartStopwatchNil sets the value for AutoStartStopwatch to be an explicit nil

### UnsetAutoStartStopwatch
`func (o *MemberTemplate) UnsetAutoStartStopwatch()`

UnsetAutoStartStopwatch ensures that no value is present for AutoStartStopwatch, not even an explicit nil
### GetAutoPopupQuickNotesWithStopwatch

`func (o *MemberTemplate) GetAutoPopupQuickNotesWithStopwatch() bool`

GetAutoPopupQuickNotesWithStopwatch returns the AutoPopupQuickNotesWithStopwatch field if non-nil, zero value otherwise.

### GetAutoPopupQuickNotesWithStopwatchOk

`func (o *MemberTemplate) GetAutoPopupQuickNotesWithStopwatchOk() (*bool, bool)`

GetAutoPopupQuickNotesWithStopwatchOk returns a tuple with the AutoPopupQuickNotesWithStopwatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPopupQuickNotesWithStopwatch

`func (o *MemberTemplate) SetAutoPopupQuickNotesWithStopwatch(v bool)`

SetAutoPopupQuickNotesWithStopwatch sets AutoPopupQuickNotesWithStopwatch field to given value.

### HasAutoPopupQuickNotesWithStopwatch

`func (o *MemberTemplate) HasAutoPopupQuickNotesWithStopwatch() bool`

HasAutoPopupQuickNotesWithStopwatch returns a boolean if a field has been set.

### SetAutoPopupQuickNotesWithStopwatchNil

`func (o *MemberTemplate) SetAutoPopupQuickNotesWithStopwatchNil(b bool)`

 SetAutoPopupQuickNotesWithStopwatchNil sets the value for AutoPopupQuickNotesWithStopwatch to be an explicit nil

### UnsetAutoPopupQuickNotesWithStopwatch
`func (o *MemberTemplate) UnsetAutoPopupQuickNotesWithStopwatch()`

UnsetAutoPopupQuickNotesWithStopwatch ensures that no value is present for AutoPopupQuickNotesWithStopwatch, not even an explicit nil
### GetGlobalSearchDefaultTicketFilter

`func (o *MemberTemplate) GetGlobalSearchDefaultTicketFilter() string`

GetGlobalSearchDefaultTicketFilter returns the GlobalSearchDefaultTicketFilter field if non-nil, zero value otherwise.

### GetGlobalSearchDefaultTicketFilterOk

`func (o *MemberTemplate) GetGlobalSearchDefaultTicketFilterOk() (*string, bool)`

GetGlobalSearchDefaultTicketFilterOk returns a tuple with the GlobalSearchDefaultTicketFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSearchDefaultTicketFilter

`func (o *MemberTemplate) SetGlobalSearchDefaultTicketFilter(v string)`

SetGlobalSearchDefaultTicketFilter sets GlobalSearchDefaultTicketFilter field to given value.

### HasGlobalSearchDefaultTicketFilter

`func (o *MemberTemplate) HasGlobalSearchDefaultTicketFilter() bool`

HasGlobalSearchDefaultTicketFilter returns a boolean if a field has been set.

### SetGlobalSearchDefaultTicketFilterNil

`func (o *MemberTemplate) SetGlobalSearchDefaultTicketFilterNil(b bool)`

 SetGlobalSearchDefaultTicketFilterNil sets the value for GlobalSearchDefaultTicketFilter to be an explicit nil

### UnsetGlobalSearchDefaultTicketFilter
`func (o *MemberTemplate) UnsetGlobalSearchDefaultTicketFilter()`

UnsetGlobalSearchDefaultTicketFilter ensures that no value is present for GlobalSearchDefaultTicketFilter, not even an explicit nil
### GetGlobalSearchDefaultSort

`func (o *MemberTemplate) GetGlobalSearchDefaultSort() string`

GetGlobalSearchDefaultSort returns the GlobalSearchDefaultSort field if non-nil, zero value otherwise.

### GetGlobalSearchDefaultSortOk

`func (o *MemberTemplate) GetGlobalSearchDefaultSortOk() (*string, bool)`

GetGlobalSearchDefaultSortOk returns a tuple with the GlobalSearchDefaultSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSearchDefaultSort

`func (o *MemberTemplate) SetGlobalSearchDefaultSort(v string)`

SetGlobalSearchDefaultSort sets GlobalSearchDefaultSort field to given value.

### HasGlobalSearchDefaultSort

`func (o *MemberTemplate) HasGlobalSearchDefaultSort() bool`

HasGlobalSearchDefaultSort returns a boolean if a field has been set.

### SetGlobalSearchDefaultSortNil

`func (o *MemberTemplate) SetGlobalSearchDefaultSortNil(b bool)`

 SetGlobalSearchDefaultSortNil sets the value for GlobalSearchDefaultSort to be an explicit nil

### UnsetGlobalSearchDefaultSort
`func (o *MemberTemplate) UnsetGlobalSearchDefaultSort()`

UnsetGlobalSearchDefaultSort ensures that no value is present for GlobalSearchDefaultSort, not even an explicit nil
### GetPhoneSource

`func (o *MemberTemplate) GetPhoneSource() string`

GetPhoneSource returns the PhoneSource field if non-nil, zero value otherwise.

### GetPhoneSourceOk

`func (o *MemberTemplate) GetPhoneSourceOk() (*string, bool)`

GetPhoneSourceOk returns a tuple with the PhoneSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneSource

`func (o *MemberTemplate) SetPhoneSource(v string)`

SetPhoneSource sets PhoneSource field to given value.

### HasPhoneSource

`func (o *MemberTemplate) HasPhoneSource() bool`

HasPhoneSource returns a boolean if a field has been set.

### GetInfo

`func (o *MemberTemplate) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MemberTemplate) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MemberTemplate) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MemberTemplate) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCopyPodLayouts

`func (o *MemberTemplate) GetCopyPodLayouts() bool`

GetCopyPodLayouts returns the CopyPodLayouts field if non-nil, zero value otherwise.

### GetCopyPodLayoutsOk

`func (o *MemberTemplate) GetCopyPodLayoutsOk() (*bool, bool)`

GetCopyPodLayoutsOk returns a tuple with the CopyPodLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyPodLayouts

`func (o *MemberTemplate) SetCopyPodLayouts(v bool)`

SetCopyPodLayouts sets CopyPodLayouts field to given value.

### HasCopyPodLayouts

`func (o *MemberTemplate) HasCopyPodLayouts() bool`

HasCopyPodLayouts returns a boolean if a field has been set.

### GetCopySharedDefaultViews

`func (o *MemberTemplate) GetCopySharedDefaultViews() bool`

GetCopySharedDefaultViews returns the CopySharedDefaultViews field if non-nil, zero value otherwise.

### GetCopySharedDefaultViewsOk

`func (o *MemberTemplate) GetCopySharedDefaultViewsOk() (*bool, bool)`

GetCopySharedDefaultViewsOk returns a tuple with the CopySharedDefaultViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySharedDefaultViews

`func (o *MemberTemplate) SetCopySharedDefaultViews(v bool)`

SetCopySharedDefaultViews sets CopySharedDefaultViews field to given value.

### HasCopySharedDefaultViews

`func (o *MemberTemplate) HasCopySharedDefaultViews() bool`

HasCopySharedDefaultViews returns a boolean if a field has been set.

### GetCopyColumnLayoutsAndFilters

`func (o *MemberTemplate) GetCopyColumnLayoutsAndFilters() bool`

GetCopyColumnLayoutsAndFilters returns the CopyColumnLayoutsAndFilters field if non-nil, zero value otherwise.

### GetCopyColumnLayoutsAndFiltersOk

`func (o *MemberTemplate) GetCopyColumnLayoutsAndFiltersOk() (*bool, bool)`

GetCopyColumnLayoutsAndFiltersOk returns a tuple with the CopyColumnLayoutsAndFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyColumnLayoutsAndFilters

`func (o *MemberTemplate) SetCopyColumnLayoutsAndFilters(v bool)`

SetCopyColumnLayoutsAndFilters sets CopyColumnLayoutsAndFilters field to given value.

### HasCopyColumnLayoutsAndFilters

`func (o *MemberTemplate) HasCopyColumnLayoutsAndFilters() bool`

HasCopyColumnLayoutsAndFilters returns a boolean if a field has been set.

### GetFromMemberRecId

`func (o *MemberTemplate) GetFromMemberRecId() int32`

GetFromMemberRecId returns the FromMemberRecId field if non-nil, zero value otherwise.

### GetFromMemberRecIdOk

`func (o *MemberTemplate) GetFromMemberRecIdOk() (*int32, bool)`

GetFromMemberRecIdOk returns a tuple with the FromMemberRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromMemberRecId

`func (o *MemberTemplate) SetFromMemberRecId(v int32)`

SetFromMemberRecId sets FromMemberRecId field to given value.

### HasFromMemberRecId

`func (o *MemberTemplate) HasFromMemberRecId() bool`

HasFromMemberRecId returns a boolean if a field has been set.

### GetFromMemberTemplateRecId

`func (o *MemberTemplate) GetFromMemberTemplateRecId() int32`

GetFromMemberTemplateRecId returns the FromMemberTemplateRecId field if non-nil, zero value otherwise.

### GetFromMemberTemplateRecIdOk

`func (o *MemberTemplate) GetFromMemberTemplateRecIdOk() (*int32, bool)`

GetFromMemberTemplateRecIdOk returns a tuple with the FromMemberTemplateRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromMemberTemplateRecId

`func (o *MemberTemplate) SetFromMemberTemplateRecId(v int32)`

SetFromMemberTemplateRecId sets FromMemberTemplateRecId field to given value.

### HasFromMemberTemplateRecId

`func (o *MemberTemplate) HasFromMemberTemplateRecId() bool`

HasFromMemberTemplateRecId returns a boolean if a field has been set.

### GetCustomFields

`func (o *MemberTemplate) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MemberTemplate) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MemberTemplate) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MemberTemplate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


