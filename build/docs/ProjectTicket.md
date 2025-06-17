# ProjectTicket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Summary** | **string** |  Max length: 100; | 
**IsIssueFlag** | Pointer to **NullableBool** |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Status** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**Phase** | [**ProjectPhaseReference**](ProjectPhaseReference.md) |  | 
**WbsCode** | Pointer to **string** |  Max length: 50; | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**SiteName** | Pointer to **string** |  Max length: 50; | [optional] 
**AddressLine1** | Pointer to **string** |  Max length: 50; | [optional] 
**AddressLine2** | Pointer to **string** |  Max length: 50; | [optional] 
**City** | Pointer to **string** |  Max length: 50; | [optional] 
**StateIdentifier** | Pointer to **string** |  Max length: 50; | [optional] 
**Zip** | Pointer to **string** |  Max length: 12; | [optional] 
**Country** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**ContactName** | Pointer to **string** |  Max length: 62; | [optional] 
**ContactPhoneNumber** | Pointer to **string** |  Max length: 20; | [optional] 
**ContactPhoneExtension** | Pointer to **string** |  Max length: 15; | [optional] 
**ContactEmailAddress** | Pointer to **string** |  Max length: 250; | [optional] 
**Type** | Pointer to [**ServiceTypeReference**](ServiceTypeReference.md) |  | [optional] 
**SubType** | Pointer to [**ServiceSubTypeReference**](ServiceSubTypeReference.md) |  | [optional] 
**Item** | Pointer to [**ServiceItemReference**](ServiceItemReference.md) |  | [optional] 
**Owner** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Priority** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**ServiceLocation** | Pointer to [**ServiceLocationReference**](ServiceLocationReference.md) |  | [optional] 
**Source** | Pointer to [**ServiceSourceReference**](ServiceSourceReference.md) |  | [optional] 
**RequiredDate** | Pointer to **time.Time** |  | [optional] 
**BudgetHours** | Pointer to **NullableFloat64** |  | [optional] 
**Opportunity** | Pointer to [**OpportunityReference**](OpportunityReference.md) |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**AgreementType** | Pointer to **string** |  | [optional] 
**KnowledgeBaseCategoryId** | Pointer to **NullableInt32** |  | [optional] 
**KnowledgeBaseSubCategoryId** | Pointer to **NullableInt32** |  | [optional] 
**KnowledgeBaseLinkId** | Pointer to **NullableInt32** |  | [optional] 
**KnowledgeBaseLinkType** | Pointer to **NullableString** |  | [optional] 
**AllowAllClientsPortalView** | Pointer to **NullableBool** |  | [optional] 
**CustomerUpdatedFlag** | Pointer to **NullableBool** |  | [optional] 
**AutomaticEmailContactFlag** | Pointer to **NullableBool** |  | [optional] 
**AutomaticEmailResourceFlag** | Pointer to **NullableBool** |  | [optional] 
**AutomaticEmailCcFlag** | Pointer to **NullableBool** |  | [optional] 
**AutomaticEmailCc** | Pointer to **string** |  Max length: 1000; | [optional] 
**ClosedDate** | Pointer to **string** |  | [optional] 
**ClosedBy** | Pointer to **string** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**ActualHours** | Pointer to **NullableFloat64** |  | [optional] 
**Approved** | Pointer to **NullableBool** |  | [optional] 
**SubBillingMethod** | Pointer to **NullableString** |  | [optional] 
**SubBillingAmount** | Pointer to **NullableFloat64** |  | [optional] 
**SubDateAccepted** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to **string** |  | [optional] 
**BillTime** | Pointer to **NullableString** |  | [optional] 
**BillExpenses** | Pointer to **NullableString** |  | [optional] 
**BillProducts** | Pointer to **NullableString** |  | [optional] 
**PredecessorType** | Pointer to **NullableString** |  | [optional] 
**PredecessorId** | Pointer to **NullableInt32** |  | [optional] 
**PredecessorClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**LagDays** | Pointer to **NullableInt32** |  | [optional] 
**LagNonworkingDaysFlag** | Pointer to **NullableBool** |  | [optional] 
**EstimatedStartDate** | Pointer to **time.Time** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**Duration** | Pointer to **NullableInt32** |  | [optional] 
**ScheduleStartDate** | Pointer to **time.Time** |  | [optional] 
**ScheduleEndDate** | Pointer to **time.Time** |  | [optional] 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**Tasks** | Pointer to [**[]TicketTask**](TicketTask.md) |  | [optional] 
**InitialDescription** | Pointer to **string** | Only available for POST, will not be returned in the response. | [optional] 
**InitialInternalAnalysis** | Pointer to **string** | Only available for POST, will not be returned in the response. | [optional] 
**InitialResolution** | Pointer to **string** | Only available for POST, will not be returned in the response. | [optional] 
**ContactEmailLookup** | Pointer to **string** |  | [optional] 
**ProcessNotifications** | Pointer to **NullableBool** | Can be set to false to skip notification processing when adding or updating a ticket (Defaults to True). | [optional] 
**SkipCallback** | Pointer to **NullableBool** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewProjectTicket

`func NewProjectTicket(summary string, phase ProjectPhaseReference, ) *ProjectTicket`

NewProjectTicket instantiates a new ProjectTicket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTicketWithDefaults

`func NewProjectTicketWithDefaults() *ProjectTicket`

NewProjectTicketWithDefaults instantiates a new ProjectTicket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTicket) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTicket) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTicket) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTicket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSummary

`func (o *ProjectTicket) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ProjectTicket) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ProjectTicket) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetIsIssueFlag

`func (o *ProjectTicket) GetIsIssueFlag() bool`

GetIsIssueFlag returns the IsIssueFlag field if non-nil, zero value otherwise.

### GetIsIssueFlagOk

`func (o *ProjectTicket) GetIsIssueFlagOk() (*bool, bool)`

GetIsIssueFlagOk returns a tuple with the IsIssueFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIssueFlag

`func (o *ProjectTicket) SetIsIssueFlag(v bool)`

SetIsIssueFlag sets IsIssueFlag field to given value.

### HasIsIssueFlag

`func (o *ProjectTicket) HasIsIssueFlag() bool`

HasIsIssueFlag returns a boolean if a field has been set.

### SetIsIssueFlagNil

`func (o *ProjectTicket) SetIsIssueFlagNil(b bool)`

 SetIsIssueFlagNil sets the value for IsIssueFlag to be an explicit nil

### UnsetIsIssueFlag
`func (o *ProjectTicket) UnsetIsIssueFlag()`

UnsetIsIssueFlag ensures that no value is present for IsIssueFlag, not even an explicit nil
### GetBoard

`func (o *ProjectTicket) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *ProjectTicket) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *ProjectTicket) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *ProjectTicket) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectTicket) GetStatus() ServiceStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectTicket) GetStatusOk() (*ServiceStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectTicket) SetStatus(v ServiceStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectTicket) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkRole

`func (o *ProjectTicket) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *ProjectTicket) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *ProjectTicket) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *ProjectTicket) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *ProjectTicket) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *ProjectTicket) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *ProjectTicket) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *ProjectTicket) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetProject

`func (o *ProjectTicket) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectTicket) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectTicket) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectTicket) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPhase

`func (o *ProjectTicket) GetPhase() ProjectPhaseReference`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ProjectTicket) GetPhaseOk() (*ProjectPhaseReference, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ProjectTicket) SetPhase(v ProjectPhaseReference)`

SetPhase sets Phase field to given value.


### GetWbsCode

`func (o *ProjectTicket) GetWbsCode() string`

GetWbsCode returns the WbsCode field if non-nil, zero value otherwise.

### GetWbsCodeOk

`func (o *ProjectTicket) GetWbsCodeOk() (*string, bool)`

GetWbsCodeOk returns a tuple with the WbsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbsCode

`func (o *ProjectTicket) SetWbsCode(v string)`

SetWbsCode sets WbsCode field to given value.

### HasWbsCode

`func (o *ProjectTicket) HasWbsCode() bool`

HasWbsCode returns a boolean if a field has been set.

### GetCompany

`func (o *ProjectTicket) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ProjectTicket) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ProjectTicket) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ProjectTicket) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetSite

`func (o *ProjectTicket) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ProjectTicket) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ProjectTicket) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *ProjectTicket) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSiteName

`func (o *ProjectTicket) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *ProjectTicket) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *ProjectTicket) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *ProjectTicket) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *ProjectTicket) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ProjectTicket) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ProjectTicket) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *ProjectTicket) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *ProjectTicket) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ProjectTicket) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ProjectTicket) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ProjectTicket) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *ProjectTicket) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ProjectTicket) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ProjectTicket) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ProjectTicket) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStateIdentifier

`func (o *ProjectTicket) GetStateIdentifier() string`

GetStateIdentifier returns the StateIdentifier field if non-nil, zero value otherwise.

### GetStateIdentifierOk

`func (o *ProjectTicket) GetStateIdentifierOk() (*string, bool)`

GetStateIdentifierOk returns a tuple with the StateIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateIdentifier

`func (o *ProjectTicket) SetStateIdentifier(v string)`

SetStateIdentifier sets StateIdentifier field to given value.

### HasStateIdentifier

`func (o *ProjectTicket) HasStateIdentifier() bool`

HasStateIdentifier returns a boolean if a field has been set.

### GetZip

`func (o *ProjectTicket) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *ProjectTicket) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *ProjectTicket) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *ProjectTicket) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetCountry

`func (o *ProjectTicket) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ProjectTicket) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ProjectTicket) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ProjectTicket) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetContact

`func (o *ProjectTicket) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ProjectTicket) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ProjectTicket) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ProjectTicket) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetContactName

`func (o *ProjectTicket) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *ProjectTicket) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *ProjectTicket) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *ProjectTicket) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactPhoneNumber

`func (o *ProjectTicket) GetContactPhoneNumber() string`

GetContactPhoneNumber returns the ContactPhoneNumber field if non-nil, zero value otherwise.

### GetContactPhoneNumberOk

`func (o *ProjectTicket) GetContactPhoneNumberOk() (*string, bool)`

GetContactPhoneNumberOk returns a tuple with the ContactPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhoneNumber

`func (o *ProjectTicket) SetContactPhoneNumber(v string)`

SetContactPhoneNumber sets ContactPhoneNumber field to given value.

### HasContactPhoneNumber

`func (o *ProjectTicket) HasContactPhoneNumber() bool`

HasContactPhoneNumber returns a boolean if a field has been set.

### GetContactPhoneExtension

`func (o *ProjectTicket) GetContactPhoneExtension() string`

GetContactPhoneExtension returns the ContactPhoneExtension field if non-nil, zero value otherwise.

### GetContactPhoneExtensionOk

`func (o *ProjectTicket) GetContactPhoneExtensionOk() (*string, bool)`

GetContactPhoneExtensionOk returns a tuple with the ContactPhoneExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhoneExtension

`func (o *ProjectTicket) SetContactPhoneExtension(v string)`

SetContactPhoneExtension sets ContactPhoneExtension field to given value.

### HasContactPhoneExtension

`func (o *ProjectTicket) HasContactPhoneExtension() bool`

HasContactPhoneExtension returns a boolean if a field has been set.

### GetContactEmailAddress

`func (o *ProjectTicket) GetContactEmailAddress() string`

GetContactEmailAddress returns the ContactEmailAddress field if non-nil, zero value otherwise.

### GetContactEmailAddressOk

`func (o *ProjectTicket) GetContactEmailAddressOk() (*string, bool)`

GetContactEmailAddressOk returns a tuple with the ContactEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmailAddress

`func (o *ProjectTicket) SetContactEmailAddress(v string)`

SetContactEmailAddress sets ContactEmailAddress field to given value.

### HasContactEmailAddress

`func (o *ProjectTicket) HasContactEmailAddress() bool`

HasContactEmailAddress returns a boolean if a field has been set.

### GetType

`func (o *ProjectTicket) GetType() ServiceTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectTicket) GetTypeOk() (*ServiceTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectTicket) SetType(v ServiceTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *ProjectTicket) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubType

`func (o *ProjectTicket) GetSubType() ServiceSubTypeReference`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *ProjectTicket) GetSubTypeOk() (*ServiceSubTypeReference, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *ProjectTicket) SetSubType(v ServiceSubTypeReference)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *ProjectTicket) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetItem

`func (o *ProjectTicket) GetItem() ServiceItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ProjectTicket) GetItemOk() (*ServiceItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ProjectTicket) SetItem(v ServiceItemReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *ProjectTicket) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetOwner

`func (o *ProjectTicket) GetOwner() MemberReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ProjectTicket) GetOwnerOk() (*MemberReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ProjectTicket) SetOwner(v MemberReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ProjectTicket) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPriority

`func (o *ProjectTicket) GetPriority() PriorityReference`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ProjectTicket) GetPriorityOk() (*PriorityReference, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ProjectTicket) SetPriority(v PriorityReference)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ProjectTicket) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetServiceLocation

`func (o *ProjectTicket) GetServiceLocation() ServiceLocationReference`

GetServiceLocation returns the ServiceLocation field if non-nil, zero value otherwise.

### GetServiceLocationOk

`func (o *ProjectTicket) GetServiceLocationOk() (*ServiceLocationReference, bool)`

GetServiceLocationOk returns a tuple with the ServiceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLocation

`func (o *ProjectTicket) SetServiceLocation(v ServiceLocationReference)`

SetServiceLocation sets ServiceLocation field to given value.

### HasServiceLocation

`func (o *ProjectTicket) HasServiceLocation() bool`

HasServiceLocation returns a boolean if a field has been set.

### GetSource

`func (o *ProjectTicket) GetSource() ServiceSourceReference`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProjectTicket) GetSourceOk() (*ServiceSourceReference, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProjectTicket) SetSource(v ServiceSourceReference)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProjectTicket) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetRequiredDate

`func (o *ProjectTicket) GetRequiredDate() time.Time`

GetRequiredDate returns the RequiredDate field if non-nil, zero value otherwise.

### GetRequiredDateOk

`func (o *ProjectTicket) GetRequiredDateOk() (*time.Time, bool)`

GetRequiredDateOk returns a tuple with the RequiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDate

`func (o *ProjectTicket) SetRequiredDate(v time.Time)`

SetRequiredDate sets RequiredDate field to given value.

### HasRequiredDate

`func (o *ProjectTicket) HasRequiredDate() bool`

HasRequiredDate returns a boolean if a field has been set.

### GetBudgetHours

`func (o *ProjectTicket) GetBudgetHours() float64`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *ProjectTicket) GetBudgetHoursOk() (*float64, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *ProjectTicket) SetBudgetHours(v float64)`

SetBudgetHours sets BudgetHours field to given value.

### HasBudgetHours

`func (o *ProjectTicket) HasBudgetHours() bool`

HasBudgetHours returns a boolean if a field has been set.

### SetBudgetHoursNil

`func (o *ProjectTicket) SetBudgetHoursNil(b bool)`

 SetBudgetHoursNil sets the value for BudgetHours to be an explicit nil

### UnsetBudgetHours
`func (o *ProjectTicket) UnsetBudgetHours()`

UnsetBudgetHours ensures that no value is present for BudgetHours, not even an explicit nil
### GetOpportunity

`func (o *ProjectTicket) GetOpportunity() OpportunityReference`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *ProjectTicket) GetOpportunityOk() (*OpportunityReference, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *ProjectTicket) SetOpportunity(v OpportunityReference)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *ProjectTicket) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### GetAgreement

`func (o *ProjectTicket) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *ProjectTicket) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *ProjectTicket) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *ProjectTicket) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetAgreementType

`func (o *ProjectTicket) GetAgreementType() string`

GetAgreementType returns the AgreementType field if non-nil, zero value otherwise.

### GetAgreementTypeOk

`func (o *ProjectTicket) GetAgreementTypeOk() (*string, bool)`

GetAgreementTypeOk returns a tuple with the AgreementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementType

`func (o *ProjectTicket) SetAgreementType(v string)`

SetAgreementType sets AgreementType field to given value.

### HasAgreementType

`func (o *ProjectTicket) HasAgreementType() bool`

HasAgreementType returns a boolean if a field has been set.

### GetKnowledgeBaseCategoryId

`func (o *ProjectTicket) GetKnowledgeBaseCategoryId() int32`

GetKnowledgeBaseCategoryId returns the KnowledgeBaseCategoryId field if non-nil, zero value otherwise.

### GetKnowledgeBaseCategoryIdOk

`func (o *ProjectTicket) GetKnowledgeBaseCategoryIdOk() (*int32, bool)`

GetKnowledgeBaseCategoryIdOk returns a tuple with the KnowledgeBaseCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBaseCategoryId

`func (o *ProjectTicket) SetKnowledgeBaseCategoryId(v int32)`

SetKnowledgeBaseCategoryId sets KnowledgeBaseCategoryId field to given value.

### HasKnowledgeBaseCategoryId

`func (o *ProjectTicket) HasKnowledgeBaseCategoryId() bool`

HasKnowledgeBaseCategoryId returns a boolean if a field has been set.

### SetKnowledgeBaseCategoryIdNil

`func (o *ProjectTicket) SetKnowledgeBaseCategoryIdNil(b bool)`

 SetKnowledgeBaseCategoryIdNil sets the value for KnowledgeBaseCategoryId to be an explicit nil

### UnsetKnowledgeBaseCategoryId
`func (o *ProjectTicket) UnsetKnowledgeBaseCategoryId()`

UnsetKnowledgeBaseCategoryId ensures that no value is present for KnowledgeBaseCategoryId, not even an explicit nil
### GetKnowledgeBaseSubCategoryId

`func (o *ProjectTicket) GetKnowledgeBaseSubCategoryId() int32`

GetKnowledgeBaseSubCategoryId returns the KnowledgeBaseSubCategoryId field if non-nil, zero value otherwise.

### GetKnowledgeBaseSubCategoryIdOk

`func (o *ProjectTicket) GetKnowledgeBaseSubCategoryIdOk() (*int32, bool)`

GetKnowledgeBaseSubCategoryIdOk returns a tuple with the KnowledgeBaseSubCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBaseSubCategoryId

`func (o *ProjectTicket) SetKnowledgeBaseSubCategoryId(v int32)`

SetKnowledgeBaseSubCategoryId sets KnowledgeBaseSubCategoryId field to given value.

### HasKnowledgeBaseSubCategoryId

`func (o *ProjectTicket) HasKnowledgeBaseSubCategoryId() bool`

HasKnowledgeBaseSubCategoryId returns a boolean if a field has been set.

### SetKnowledgeBaseSubCategoryIdNil

`func (o *ProjectTicket) SetKnowledgeBaseSubCategoryIdNil(b bool)`

 SetKnowledgeBaseSubCategoryIdNil sets the value for KnowledgeBaseSubCategoryId to be an explicit nil

### UnsetKnowledgeBaseSubCategoryId
`func (o *ProjectTicket) UnsetKnowledgeBaseSubCategoryId()`

UnsetKnowledgeBaseSubCategoryId ensures that no value is present for KnowledgeBaseSubCategoryId, not even an explicit nil
### GetKnowledgeBaseLinkId

`func (o *ProjectTicket) GetKnowledgeBaseLinkId() int32`

GetKnowledgeBaseLinkId returns the KnowledgeBaseLinkId field if non-nil, zero value otherwise.

### GetKnowledgeBaseLinkIdOk

`func (o *ProjectTicket) GetKnowledgeBaseLinkIdOk() (*int32, bool)`

GetKnowledgeBaseLinkIdOk returns a tuple with the KnowledgeBaseLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBaseLinkId

`func (o *ProjectTicket) SetKnowledgeBaseLinkId(v int32)`

SetKnowledgeBaseLinkId sets KnowledgeBaseLinkId field to given value.

### HasKnowledgeBaseLinkId

`func (o *ProjectTicket) HasKnowledgeBaseLinkId() bool`

HasKnowledgeBaseLinkId returns a boolean if a field has been set.

### SetKnowledgeBaseLinkIdNil

`func (o *ProjectTicket) SetKnowledgeBaseLinkIdNil(b bool)`

 SetKnowledgeBaseLinkIdNil sets the value for KnowledgeBaseLinkId to be an explicit nil

### UnsetKnowledgeBaseLinkId
`func (o *ProjectTicket) UnsetKnowledgeBaseLinkId()`

UnsetKnowledgeBaseLinkId ensures that no value is present for KnowledgeBaseLinkId, not even an explicit nil
### GetKnowledgeBaseLinkType

`func (o *ProjectTicket) GetKnowledgeBaseLinkType() string`

GetKnowledgeBaseLinkType returns the KnowledgeBaseLinkType field if non-nil, zero value otherwise.

### GetKnowledgeBaseLinkTypeOk

`func (o *ProjectTicket) GetKnowledgeBaseLinkTypeOk() (*string, bool)`

GetKnowledgeBaseLinkTypeOk returns a tuple with the KnowledgeBaseLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBaseLinkType

`func (o *ProjectTicket) SetKnowledgeBaseLinkType(v string)`

SetKnowledgeBaseLinkType sets KnowledgeBaseLinkType field to given value.

### HasKnowledgeBaseLinkType

`func (o *ProjectTicket) HasKnowledgeBaseLinkType() bool`

HasKnowledgeBaseLinkType returns a boolean if a field has been set.

### SetKnowledgeBaseLinkTypeNil

`func (o *ProjectTicket) SetKnowledgeBaseLinkTypeNil(b bool)`

 SetKnowledgeBaseLinkTypeNil sets the value for KnowledgeBaseLinkType to be an explicit nil

### UnsetKnowledgeBaseLinkType
`func (o *ProjectTicket) UnsetKnowledgeBaseLinkType()`

UnsetKnowledgeBaseLinkType ensures that no value is present for KnowledgeBaseLinkType, not even an explicit nil
### GetAllowAllClientsPortalView

`func (o *ProjectTicket) GetAllowAllClientsPortalView() bool`

GetAllowAllClientsPortalView returns the AllowAllClientsPortalView field if non-nil, zero value otherwise.

### GetAllowAllClientsPortalViewOk

`func (o *ProjectTicket) GetAllowAllClientsPortalViewOk() (*bool, bool)`

GetAllowAllClientsPortalViewOk returns a tuple with the AllowAllClientsPortalView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAllClientsPortalView

`func (o *ProjectTicket) SetAllowAllClientsPortalView(v bool)`

SetAllowAllClientsPortalView sets AllowAllClientsPortalView field to given value.

### HasAllowAllClientsPortalView

`func (o *ProjectTicket) HasAllowAllClientsPortalView() bool`

HasAllowAllClientsPortalView returns a boolean if a field has been set.

### SetAllowAllClientsPortalViewNil

`func (o *ProjectTicket) SetAllowAllClientsPortalViewNil(b bool)`

 SetAllowAllClientsPortalViewNil sets the value for AllowAllClientsPortalView to be an explicit nil

### UnsetAllowAllClientsPortalView
`func (o *ProjectTicket) UnsetAllowAllClientsPortalView()`

UnsetAllowAllClientsPortalView ensures that no value is present for AllowAllClientsPortalView, not even an explicit nil
### GetCustomerUpdatedFlag

`func (o *ProjectTicket) GetCustomerUpdatedFlag() bool`

GetCustomerUpdatedFlag returns the CustomerUpdatedFlag field if non-nil, zero value otherwise.

### GetCustomerUpdatedFlagOk

`func (o *ProjectTicket) GetCustomerUpdatedFlagOk() (*bool, bool)`

GetCustomerUpdatedFlagOk returns a tuple with the CustomerUpdatedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUpdatedFlag

`func (o *ProjectTicket) SetCustomerUpdatedFlag(v bool)`

SetCustomerUpdatedFlag sets CustomerUpdatedFlag field to given value.

### HasCustomerUpdatedFlag

`func (o *ProjectTicket) HasCustomerUpdatedFlag() bool`

HasCustomerUpdatedFlag returns a boolean if a field has been set.

### SetCustomerUpdatedFlagNil

`func (o *ProjectTicket) SetCustomerUpdatedFlagNil(b bool)`

 SetCustomerUpdatedFlagNil sets the value for CustomerUpdatedFlag to be an explicit nil

### UnsetCustomerUpdatedFlag
`func (o *ProjectTicket) UnsetCustomerUpdatedFlag()`

UnsetCustomerUpdatedFlag ensures that no value is present for CustomerUpdatedFlag, not even an explicit nil
### GetAutomaticEmailContactFlag

`func (o *ProjectTicket) GetAutomaticEmailContactFlag() bool`

GetAutomaticEmailContactFlag returns the AutomaticEmailContactFlag field if non-nil, zero value otherwise.

### GetAutomaticEmailContactFlagOk

`func (o *ProjectTicket) GetAutomaticEmailContactFlagOk() (*bool, bool)`

GetAutomaticEmailContactFlagOk returns a tuple with the AutomaticEmailContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticEmailContactFlag

`func (o *ProjectTicket) SetAutomaticEmailContactFlag(v bool)`

SetAutomaticEmailContactFlag sets AutomaticEmailContactFlag field to given value.

### HasAutomaticEmailContactFlag

`func (o *ProjectTicket) HasAutomaticEmailContactFlag() bool`

HasAutomaticEmailContactFlag returns a boolean if a field has been set.

### SetAutomaticEmailContactFlagNil

`func (o *ProjectTicket) SetAutomaticEmailContactFlagNil(b bool)`

 SetAutomaticEmailContactFlagNil sets the value for AutomaticEmailContactFlag to be an explicit nil

### UnsetAutomaticEmailContactFlag
`func (o *ProjectTicket) UnsetAutomaticEmailContactFlag()`

UnsetAutomaticEmailContactFlag ensures that no value is present for AutomaticEmailContactFlag, not even an explicit nil
### GetAutomaticEmailResourceFlag

`func (o *ProjectTicket) GetAutomaticEmailResourceFlag() bool`

GetAutomaticEmailResourceFlag returns the AutomaticEmailResourceFlag field if non-nil, zero value otherwise.

### GetAutomaticEmailResourceFlagOk

`func (o *ProjectTicket) GetAutomaticEmailResourceFlagOk() (*bool, bool)`

GetAutomaticEmailResourceFlagOk returns a tuple with the AutomaticEmailResourceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticEmailResourceFlag

`func (o *ProjectTicket) SetAutomaticEmailResourceFlag(v bool)`

SetAutomaticEmailResourceFlag sets AutomaticEmailResourceFlag field to given value.

### HasAutomaticEmailResourceFlag

`func (o *ProjectTicket) HasAutomaticEmailResourceFlag() bool`

HasAutomaticEmailResourceFlag returns a boolean if a field has been set.

### SetAutomaticEmailResourceFlagNil

`func (o *ProjectTicket) SetAutomaticEmailResourceFlagNil(b bool)`

 SetAutomaticEmailResourceFlagNil sets the value for AutomaticEmailResourceFlag to be an explicit nil

### UnsetAutomaticEmailResourceFlag
`func (o *ProjectTicket) UnsetAutomaticEmailResourceFlag()`

UnsetAutomaticEmailResourceFlag ensures that no value is present for AutomaticEmailResourceFlag, not even an explicit nil
### GetAutomaticEmailCcFlag

`func (o *ProjectTicket) GetAutomaticEmailCcFlag() bool`

GetAutomaticEmailCcFlag returns the AutomaticEmailCcFlag field if non-nil, zero value otherwise.

### GetAutomaticEmailCcFlagOk

`func (o *ProjectTicket) GetAutomaticEmailCcFlagOk() (*bool, bool)`

GetAutomaticEmailCcFlagOk returns a tuple with the AutomaticEmailCcFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticEmailCcFlag

`func (o *ProjectTicket) SetAutomaticEmailCcFlag(v bool)`

SetAutomaticEmailCcFlag sets AutomaticEmailCcFlag field to given value.

### HasAutomaticEmailCcFlag

`func (o *ProjectTicket) HasAutomaticEmailCcFlag() bool`

HasAutomaticEmailCcFlag returns a boolean if a field has been set.

### SetAutomaticEmailCcFlagNil

`func (o *ProjectTicket) SetAutomaticEmailCcFlagNil(b bool)`

 SetAutomaticEmailCcFlagNil sets the value for AutomaticEmailCcFlag to be an explicit nil

### UnsetAutomaticEmailCcFlag
`func (o *ProjectTicket) UnsetAutomaticEmailCcFlag()`

UnsetAutomaticEmailCcFlag ensures that no value is present for AutomaticEmailCcFlag, not even an explicit nil
### GetAutomaticEmailCc

`func (o *ProjectTicket) GetAutomaticEmailCc() string`

GetAutomaticEmailCc returns the AutomaticEmailCc field if non-nil, zero value otherwise.

### GetAutomaticEmailCcOk

`func (o *ProjectTicket) GetAutomaticEmailCcOk() (*string, bool)`

GetAutomaticEmailCcOk returns a tuple with the AutomaticEmailCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticEmailCc

`func (o *ProjectTicket) SetAutomaticEmailCc(v string)`

SetAutomaticEmailCc sets AutomaticEmailCc field to given value.

### HasAutomaticEmailCc

`func (o *ProjectTicket) HasAutomaticEmailCc() bool`

HasAutomaticEmailCc returns a boolean if a field has been set.

### GetClosedDate

`func (o *ProjectTicket) GetClosedDate() string`

GetClosedDate returns the ClosedDate field if non-nil, zero value otherwise.

### GetClosedDateOk

`func (o *ProjectTicket) GetClosedDateOk() (*string, bool)`

GetClosedDateOk returns a tuple with the ClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDate

`func (o *ProjectTicket) SetClosedDate(v string)`

SetClosedDate sets ClosedDate field to given value.

### HasClosedDate

`func (o *ProjectTicket) HasClosedDate() bool`

HasClosedDate returns a boolean if a field has been set.

### GetClosedBy

`func (o *ProjectTicket) GetClosedBy() string`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *ProjectTicket) GetClosedByOk() (*string, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *ProjectTicket) SetClosedBy(v string)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *ProjectTicket) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetClosedFlag

`func (o *ProjectTicket) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *ProjectTicket) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *ProjectTicket) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *ProjectTicket) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *ProjectTicket) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *ProjectTicket) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetActualHours

`func (o *ProjectTicket) GetActualHours() float64`

GetActualHours returns the ActualHours field if non-nil, zero value otherwise.

### GetActualHoursOk

`func (o *ProjectTicket) GetActualHoursOk() (*float64, bool)`

GetActualHoursOk returns a tuple with the ActualHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualHours

`func (o *ProjectTicket) SetActualHours(v float64)`

SetActualHours sets ActualHours field to given value.

### HasActualHours

`func (o *ProjectTicket) HasActualHours() bool`

HasActualHours returns a boolean if a field has been set.

### SetActualHoursNil

`func (o *ProjectTicket) SetActualHoursNil(b bool)`

 SetActualHoursNil sets the value for ActualHours to be an explicit nil

### UnsetActualHours
`func (o *ProjectTicket) UnsetActualHours()`

UnsetActualHours ensures that no value is present for ActualHours, not even an explicit nil
### GetApproved

`func (o *ProjectTicket) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *ProjectTicket) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *ProjectTicket) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *ProjectTicket) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### SetApprovedNil

`func (o *ProjectTicket) SetApprovedNil(b bool)`

 SetApprovedNil sets the value for Approved to be an explicit nil

### UnsetApproved
`func (o *ProjectTicket) UnsetApproved()`

UnsetApproved ensures that no value is present for Approved, not even an explicit nil
### GetSubBillingMethod

`func (o *ProjectTicket) GetSubBillingMethod() string`

GetSubBillingMethod returns the SubBillingMethod field if non-nil, zero value otherwise.

### GetSubBillingMethodOk

`func (o *ProjectTicket) GetSubBillingMethodOk() (*string, bool)`

GetSubBillingMethodOk returns a tuple with the SubBillingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubBillingMethod

`func (o *ProjectTicket) SetSubBillingMethod(v string)`

SetSubBillingMethod sets SubBillingMethod field to given value.

### HasSubBillingMethod

`func (o *ProjectTicket) HasSubBillingMethod() bool`

HasSubBillingMethod returns a boolean if a field has been set.

### SetSubBillingMethodNil

`func (o *ProjectTicket) SetSubBillingMethodNil(b bool)`

 SetSubBillingMethodNil sets the value for SubBillingMethod to be an explicit nil

### UnsetSubBillingMethod
`func (o *ProjectTicket) UnsetSubBillingMethod()`

UnsetSubBillingMethod ensures that no value is present for SubBillingMethod, not even an explicit nil
### GetSubBillingAmount

`func (o *ProjectTicket) GetSubBillingAmount() float64`

GetSubBillingAmount returns the SubBillingAmount field if non-nil, zero value otherwise.

### GetSubBillingAmountOk

`func (o *ProjectTicket) GetSubBillingAmountOk() (*float64, bool)`

GetSubBillingAmountOk returns a tuple with the SubBillingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubBillingAmount

`func (o *ProjectTicket) SetSubBillingAmount(v float64)`

SetSubBillingAmount sets SubBillingAmount field to given value.

### HasSubBillingAmount

`func (o *ProjectTicket) HasSubBillingAmount() bool`

HasSubBillingAmount returns a boolean if a field has been set.

### SetSubBillingAmountNil

`func (o *ProjectTicket) SetSubBillingAmountNil(b bool)`

 SetSubBillingAmountNil sets the value for SubBillingAmount to be an explicit nil

### UnsetSubBillingAmount
`func (o *ProjectTicket) UnsetSubBillingAmount()`

UnsetSubBillingAmount ensures that no value is present for SubBillingAmount, not even an explicit nil
### GetSubDateAccepted

`func (o *ProjectTicket) GetSubDateAccepted() string`

GetSubDateAccepted returns the SubDateAccepted field if non-nil, zero value otherwise.

### GetSubDateAcceptedOk

`func (o *ProjectTicket) GetSubDateAcceptedOk() (*string, bool)`

GetSubDateAcceptedOk returns a tuple with the SubDateAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDateAccepted

`func (o *ProjectTicket) SetSubDateAccepted(v string)`

SetSubDateAccepted sets SubDateAccepted field to given value.

### HasSubDateAccepted

`func (o *ProjectTicket) HasSubDateAccepted() bool`

HasSubDateAccepted returns a boolean if a field has been set.

### GetResources

`func (o *ProjectTicket) GetResources() string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ProjectTicket) GetResourcesOk() (*string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ProjectTicket) SetResources(v string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ProjectTicket) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetBillTime

`func (o *ProjectTicket) GetBillTime() string`

GetBillTime returns the BillTime field if non-nil, zero value otherwise.

### GetBillTimeOk

`func (o *ProjectTicket) GetBillTimeOk() (*string, bool)`

GetBillTimeOk returns a tuple with the BillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTime

`func (o *ProjectTicket) SetBillTime(v string)`

SetBillTime sets BillTime field to given value.

### HasBillTime

`func (o *ProjectTicket) HasBillTime() bool`

HasBillTime returns a boolean if a field has been set.

### SetBillTimeNil

`func (o *ProjectTicket) SetBillTimeNil(b bool)`

 SetBillTimeNil sets the value for BillTime to be an explicit nil

### UnsetBillTime
`func (o *ProjectTicket) UnsetBillTime()`

UnsetBillTime ensures that no value is present for BillTime, not even an explicit nil
### GetBillExpenses

`func (o *ProjectTicket) GetBillExpenses() string`

GetBillExpenses returns the BillExpenses field if non-nil, zero value otherwise.

### GetBillExpensesOk

`func (o *ProjectTicket) GetBillExpensesOk() (*string, bool)`

GetBillExpensesOk returns a tuple with the BillExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillExpenses

`func (o *ProjectTicket) SetBillExpenses(v string)`

SetBillExpenses sets BillExpenses field to given value.

### HasBillExpenses

`func (o *ProjectTicket) HasBillExpenses() bool`

HasBillExpenses returns a boolean if a field has been set.

### SetBillExpensesNil

`func (o *ProjectTicket) SetBillExpensesNil(b bool)`

 SetBillExpensesNil sets the value for BillExpenses to be an explicit nil

### UnsetBillExpenses
`func (o *ProjectTicket) UnsetBillExpenses()`

UnsetBillExpenses ensures that no value is present for BillExpenses, not even an explicit nil
### GetBillProducts

`func (o *ProjectTicket) GetBillProducts() string`

GetBillProducts returns the BillProducts field if non-nil, zero value otherwise.

### GetBillProductsOk

`func (o *ProjectTicket) GetBillProductsOk() (*string, bool)`

GetBillProductsOk returns a tuple with the BillProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProducts

`func (o *ProjectTicket) SetBillProducts(v string)`

SetBillProducts sets BillProducts field to given value.

### HasBillProducts

`func (o *ProjectTicket) HasBillProducts() bool`

HasBillProducts returns a boolean if a field has been set.

### SetBillProductsNil

`func (o *ProjectTicket) SetBillProductsNil(b bool)`

 SetBillProductsNil sets the value for BillProducts to be an explicit nil

### UnsetBillProducts
`func (o *ProjectTicket) UnsetBillProducts()`

UnsetBillProducts ensures that no value is present for BillProducts, not even an explicit nil
### GetPredecessorType

`func (o *ProjectTicket) GetPredecessorType() string`

GetPredecessorType returns the PredecessorType field if non-nil, zero value otherwise.

### GetPredecessorTypeOk

`func (o *ProjectTicket) GetPredecessorTypeOk() (*string, bool)`

GetPredecessorTypeOk returns a tuple with the PredecessorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorType

`func (o *ProjectTicket) SetPredecessorType(v string)`

SetPredecessorType sets PredecessorType field to given value.

### HasPredecessorType

`func (o *ProjectTicket) HasPredecessorType() bool`

HasPredecessorType returns a boolean if a field has been set.

### SetPredecessorTypeNil

`func (o *ProjectTicket) SetPredecessorTypeNil(b bool)`

 SetPredecessorTypeNil sets the value for PredecessorType to be an explicit nil

### UnsetPredecessorType
`func (o *ProjectTicket) UnsetPredecessorType()`

UnsetPredecessorType ensures that no value is present for PredecessorType, not even an explicit nil
### GetPredecessorId

`func (o *ProjectTicket) GetPredecessorId() int32`

GetPredecessorId returns the PredecessorId field if non-nil, zero value otherwise.

### GetPredecessorIdOk

`func (o *ProjectTicket) GetPredecessorIdOk() (*int32, bool)`

GetPredecessorIdOk returns a tuple with the PredecessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorId

`func (o *ProjectTicket) SetPredecessorId(v int32)`

SetPredecessorId sets PredecessorId field to given value.

### HasPredecessorId

`func (o *ProjectTicket) HasPredecessorId() bool`

HasPredecessorId returns a boolean if a field has been set.

### SetPredecessorIdNil

`func (o *ProjectTicket) SetPredecessorIdNil(b bool)`

 SetPredecessorIdNil sets the value for PredecessorId to be an explicit nil

### UnsetPredecessorId
`func (o *ProjectTicket) UnsetPredecessorId()`

UnsetPredecessorId ensures that no value is present for PredecessorId, not even an explicit nil
### GetPredecessorClosedFlag

`func (o *ProjectTicket) GetPredecessorClosedFlag() bool`

GetPredecessorClosedFlag returns the PredecessorClosedFlag field if non-nil, zero value otherwise.

### GetPredecessorClosedFlagOk

`func (o *ProjectTicket) GetPredecessorClosedFlagOk() (*bool, bool)`

GetPredecessorClosedFlagOk returns a tuple with the PredecessorClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorClosedFlag

`func (o *ProjectTicket) SetPredecessorClosedFlag(v bool)`

SetPredecessorClosedFlag sets PredecessorClosedFlag field to given value.

### HasPredecessorClosedFlag

`func (o *ProjectTicket) HasPredecessorClosedFlag() bool`

HasPredecessorClosedFlag returns a boolean if a field has been set.

### SetPredecessorClosedFlagNil

`func (o *ProjectTicket) SetPredecessorClosedFlagNil(b bool)`

 SetPredecessorClosedFlagNil sets the value for PredecessorClosedFlag to be an explicit nil

### UnsetPredecessorClosedFlag
`func (o *ProjectTicket) UnsetPredecessorClosedFlag()`

UnsetPredecessorClosedFlag ensures that no value is present for PredecessorClosedFlag, not even an explicit nil
### GetLagDays

`func (o *ProjectTicket) GetLagDays() int32`

GetLagDays returns the LagDays field if non-nil, zero value otherwise.

### GetLagDaysOk

`func (o *ProjectTicket) GetLagDaysOk() (*int32, bool)`

GetLagDaysOk returns a tuple with the LagDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagDays

`func (o *ProjectTicket) SetLagDays(v int32)`

SetLagDays sets LagDays field to given value.

### HasLagDays

`func (o *ProjectTicket) HasLagDays() bool`

HasLagDays returns a boolean if a field has been set.

### SetLagDaysNil

`func (o *ProjectTicket) SetLagDaysNil(b bool)`

 SetLagDaysNil sets the value for LagDays to be an explicit nil

### UnsetLagDays
`func (o *ProjectTicket) UnsetLagDays()`

UnsetLagDays ensures that no value is present for LagDays, not even an explicit nil
### GetLagNonworkingDaysFlag

`func (o *ProjectTicket) GetLagNonworkingDaysFlag() bool`

GetLagNonworkingDaysFlag returns the LagNonworkingDaysFlag field if non-nil, zero value otherwise.

### GetLagNonworkingDaysFlagOk

`func (o *ProjectTicket) GetLagNonworkingDaysFlagOk() (*bool, bool)`

GetLagNonworkingDaysFlagOk returns a tuple with the LagNonworkingDaysFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagNonworkingDaysFlag

`func (o *ProjectTicket) SetLagNonworkingDaysFlag(v bool)`

SetLagNonworkingDaysFlag sets LagNonworkingDaysFlag field to given value.

### HasLagNonworkingDaysFlag

`func (o *ProjectTicket) HasLagNonworkingDaysFlag() bool`

HasLagNonworkingDaysFlag returns a boolean if a field has been set.

### SetLagNonworkingDaysFlagNil

`func (o *ProjectTicket) SetLagNonworkingDaysFlagNil(b bool)`

 SetLagNonworkingDaysFlagNil sets the value for LagNonworkingDaysFlag to be an explicit nil

### UnsetLagNonworkingDaysFlag
`func (o *ProjectTicket) UnsetLagNonworkingDaysFlag()`

UnsetLagNonworkingDaysFlag ensures that no value is present for LagNonworkingDaysFlag, not even an explicit nil
### GetEstimatedStartDate

`func (o *ProjectTicket) GetEstimatedStartDate() time.Time`

GetEstimatedStartDate returns the EstimatedStartDate field if non-nil, zero value otherwise.

### GetEstimatedStartDateOk

`func (o *ProjectTicket) GetEstimatedStartDateOk() (*time.Time, bool)`

GetEstimatedStartDateOk returns a tuple with the EstimatedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStartDate

`func (o *ProjectTicket) SetEstimatedStartDate(v time.Time)`

SetEstimatedStartDate sets EstimatedStartDate field to given value.

### HasEstimatedStartDate

`func (o *ProjectTicket) HasEstimatedStartDate() bool`

HasEstimatedStartDate returns a boolean if a field has been set.

### GetLocation

`func (o *ProjectTicket) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProjectTicket) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProjectTicket) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ProjectTicket) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *ProjectTicket) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ProjectTicket) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ProjectTicket) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ProjectTicket) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDuration

`func (o *ProjectTicket) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ProjectTicket) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ProjectTicket) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ProjectTicket) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ProjectTicket) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ProjectTicket) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetScheduleStartDate

`func (o *ProjectTicket) GetScheduleStartDate() time.Time`

GetScheduleStartDate returns the ScheduleStartDate field if non-nil, zero value otherwise.

### GetScheduleStartDateOk

`func (o *ProjectTicket) GetScheduleStartDateOk() (*time.Time, bool)`

GetScheduleStartDateOk returns a tuple with the ScheduleStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStartDate

`func (o *ProjectTicket) SetScheduleStartDate(v time.Time)`

SetScheduleStartDate sets ScheduleStartDate field to given value.

### HasScheduleStartDate

`func (o *ProjectTicket) HasScheduleStartDate() bool`

HasScheduleStartDate returns a boolean if a field has been set.

### GetScheduleEndDate

`func (o *ProjectTicket) GetScheduleEndDate() time.Time`

GetScheduleEndDate returns the ScheduleEndDate field if non-nil, zero value otherwise.

### GetScheduleEndDateOk

`func (o *ProjectTicket) GetScheduleEndDateOk() (*time.Time, bool)`

GetScheduleEndDateOk returns a tuple with the ScheduleEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEndDate

`func (o *ProjectTicket) SetScheduleEndDate(v time.Time)`

SetScheduleEndDate sets ScheduleEndDate field to given value.

### HasScheduleEndDate

`func (o *ProjectTicket) HasScheduleEndDate() bool`

HasScheduleEndDate returns a boolean if a field has been set.

### GetMobileGuid

`func (o *ProjectTicket) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *ProjectTicket) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *ProjectTicket) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *ProjectTicket) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *ProjectTicket) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *ProjectTicket) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetCurrency

`func (o *ProjectTicket) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProjectTicket) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProjectTicket) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProjectTicket) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInfo

`func (o *ProjectTicket) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectTicket) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectTicket) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectTicket) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetTasks

`func (o *ProjectTicket) GetTasks() []TicketTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ProjectTicket) GetTasksOk() (*[]TicketTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ProjectTicket) SetTasks(v []TicketTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ProjectTicket) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetInitialDescription

`func (o *ProjectTicket) GetInitialDescription() string`

GetInitialDescription returns the InitialDescription field if non-nil, zero value otherwise.

### GetInitialDescriptionOk

`func (o *ProjectTicket) GetInitialDescriptionOk() (*string, bool)`

GetInitialDescriptionOk returns a tuple with the InitialDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDescription

`func (o *ProjectTicket) SetInitialDescription(v string)`

SetInitialDescription sets InitialDescription field to given value.

### HasInitialDescription

`func (o *ProjectTicket) HasInitialDescription() bool`

HasInitialDescription returns a boolean if a field has been set.

### GetInitialInternalAnalysis

`func (o *ProjectTicket) GetInitialInternalAnalysis() string`

GetInitialInternalAnalysis returns the InitialInternalAnalysis field if non-nil, zero value otherwise.

### GetInitialInternalAnalysisOk

`func (o *ProjectTicket) GetInitialInternalAnalysisOk() (*string, bool)`

GetInitialInternalAnalysisOk returns a tuple with the InitialInternalAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialInternalAnalysis

`func (o *ProjectTicket) SetInitialInternalAnalysis(v string)`

SetInitialInternalAnalysis sets InitialInternalAnalysis field to given value.

### HasInitialInternalAnalysis

`func (o *ProjectTicket) HasInitialInternalAnalysis() bool`

HasInitialInternalAnalysis returns a boolean if a field has been set.

### GetInitialResolution

`func (o *ProjectTicket) GetInitialResolution() string`

GetInitialResolution returns the InitialResolution field if non-nil, zero value otherwise.

### GetInitialResolutionOk

`func (o *ProjectTicket) GetInitialResolutionOk() (*string, bool)`

GetInitialResolutionOk returns a tuple with the InitialResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialResolution

`func (o *ProjectTicket) SetInitialResolution(v string)`

SetInitialResolution sets InitialResolution field to given value.

### HasInitialResolution

`func (o *ProjectTicket) HasInitialResolution() bool`

HasInitialResolution returns a boolean if a field has been set.

### GetContactEmailLookup

`func (o *ProjectTicket) GetContactEmailLookup() string`

GetContactEmailLookup returns the ContactEmailLookup field if non-nil, zero value otherwise.

### GetContactEmailLookupOk

`func (o *ProjectTicket) GetContactEmailLookupOk() (*string, bool)`

GetContactEmailLookupOk returns a tuple with the ContactEmailLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmailLookup

`func (o *ProjectTicket) SetContactEmailLookup(v string)`

SetContactEmailLookup sets ContactEmailLookup field to given value.

### HasContactEmailLookup

`func (o *ProjectTicket) HasContactEmailLookup() bool`

HasContactEmailLookup returns a boolean if a field has been set.

### GetProcessNotifications

`func (o *ProjectTicket) GetProcessNotifications() bool`

GetProcessNotifications returns the ProcessNotifications field if non-nil, zero value otherwise.

### GetProcessNotificationsOk

`func (o *ProjectTicket) GetProcessNotificationsOk() (*bool, bool)`

GetProcessNotificationsOk returns a tuple with the ProcessNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessNotifications

`func (o *ProjectTicket) SetProcessNotifications(v bool)`

SetProcessNotifications sets ProcessNotifications field to given value.

### HasProcessNotifications

`func (o *ProjectTicket) HasProcessNotifications() bool`

HasProcessNotifications returns a boolean if a field has been set.

### SetProcessNotificationsNil

`func (o *ProjectTicket) SetProcessNotificationsNil(b bool)`

 SetProcessNotificationsNil sets the value for ProcessNotifications to be an explicit nil

### UnsetProcessNotifications
`func (o *ProjectTicket) UnsetProcessNotifications()`

UnsetProcessNotifications ensures that no value is present for ProcessNotifications, not even an explicit nil
### GetSkipCallback

`func (o *ProjectTicket) GetSkipCallback() bool`

GetSkipCallback returns the SkipCallback field if non-nil, zero value otherwise.

### GetSkipCallbackOk

`func (o *ProjectTicket) GetSkipCallbackOk() (*bool, bool)`

GetSkipCallbackOk returns a tuple with the SkipCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCallback

`func (o *ProjectTicket) SetSkipCallback(v bool)`

SetSkipCallback sets SkipCallback field to given value.

### HasSkipCallback

`func (o *ProjectTicket) HasSkipCallback() bool`

HasSkipCallback returns a boolean if a field has been set.

### SetSkipCallbackNil

`func (o *ProjectTicket) SetSkipCallbackNil(b bool)`

 SetSkipCallbackNil sets the value for SkipCallback to be an explicit nil

### UnsetSkipCallback
`func (o *ProjectTicket) UnsetSkipCallback()`

UnsetSkipCallback ensures that no value is present for SkipCallback, not even an explicit nil
### GetCustomFields

`func (o *ProjectTicket) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProjectTicket) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProjectTicket) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProjectTicket) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


