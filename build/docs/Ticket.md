# Ticket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Summary** | **string** |  Max length: 100; | 
**RecordType** | Pointer to **NullableString** |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Status** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
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
**Team** | Pointer to [**ServiceTeamReference**](ServiceTeamReference.md) |  | [optional] 
**Owner** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Priority** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**ServiceLocation** | Pointer to [**ServiceLocationReference**](ServiceLocationReference.md) |  | [optional] 
**Source** | Pointer to [**ServiceSourceReference**](ServiceSourceReference.md) |  | [optional] 
**RequiredDate** | Pointer to **time.Time** |  | [optional] 
**BudgetHours** | Pointer to **NullableFloat64** |  | [optional] 
**Opportunity** | Pointer to [**OpportunityReference**](OpportunityReference.md) |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**AgreementType** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**Impact** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**ExternalXRef** | Pointer to **string** |  Max length: 100; | [optional] 
**PoNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**KnowledgeBaseCategoryId** | Pointer to **NullableInt32** |  | [optional] 
**KnowledgeBaseSubCategoryId** | Pointer to **NullableInt32** |  | [optional] 
**AllowAllClientsPortalView** | Pointer to **NullableBool** |  | [optional] 
**CustomerUpdatedFlag** | Pointer to **NullableBool** |  | [optional] 
**AutomaticEmailContactFlag** | Pointer to **NullableBool** |  | [optional] 
**AutomaticEmailResourceFlag** | Pointer to **NullableBool** |  | [optional] 
**AutomaticEmailCcFlag** | Pointer to **NullableBool** |  | [optional] 
**AutomaticEmailCc** | Pointer to **string** |  Max length: 1000; | [optional] 
**InitialDescription** | Pointer to **string** | Only available for POST, will not be returned in the response. | [optional] 
**InitialInternalAnalysis** | Pointer to **string** | Only available for POST, will not be returned in the response. | [optional] 
**InitialResolution** | Pointer to **string** | Only available for POST, will not be returned in the response. | [optional] 
**InitialDescriptionFrom** | Pointer to **string** |  | [optional] 
**ContactEmailLookup** | Pointer to **string** |  | [optional] 
**ProcessNotifications** | Pointer to **NullableBool** | Can be set to false to skip notification processing when adding or updating a ticket (Defaults to True). | [optional] 
**SkipCallback** | Pointer to **NullableBool** |  | [optional] 
**ClosedDate** | Pointer to **string** |  | [optional] 
**ClosedBy** | Pointer to **string** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**ActualHours** | Pointer to **NullableFloat64** |  | [optional] 
**Approved** | Pointer to **NullableBool** |  | [optional] 
**EstimatedExpenseCost** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedExpenseRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedProductCost** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedProductRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedTimeCost** | Pointer to **NullableFloat64** |  | [optional] 
**EstimatedTimeRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**BillingMethod** | Pointer to **NullableString** |  | [optional] 
**BillingAmount** | Pointer to **NullableFloat64** |  | [optional] 
**HourlyRate** | Pointer to **NullableFloat64** |  | [optional] 
**SubBillingMethod** | Pointer to **NullableString** |  | [optional] 
**SubBillingAmount** | Pointer to **NullableFloat64** |  | [optional] 
**SubDateAccepted** | Pointer to **string** |  | [optional] 
**DateResolved** | Pointer to **string** |  | [optional] 
**DateResplan** | Pointer to **string** |  | [optional] 
**DateResponded** | Pointer to **string** |  | [optional] 
**ResolveMinutes** | Pointer to **NullableInt32** |  | [optional] 
**ResPlanMinutes** | Pointer to **NullableInt32** |  | [optional] 
**RespondMinutes** | Pointer to **NullableInt32** |  | [optional] 
**IsInSla** | Pointer to **NullableBool** |  | [optional] 
**KnowledgeBaseLinkId** | Pointer to **NullableInt32** |  | [optional] 
**Resources** | Pointer to **string** |  | [optional] 
**ParentTicketId** | Pointer to **NullableInt32** |  | [optional] 
**HasChildTicket** | Pointer to **NullableBool** |  | [optional] 
**HasMergedChildTicketFlag** | Pointer to **NullableBool** |  | [optional] 
**KnowledgeBaseLinkType** | Pointer to **NullableString** |  | [optional] 
**BillTime** | Pointer to **NullableString** |  | [optional] 
**BillExpenses** | Pointer to **NullableString** |  | [optional] 
**BillProducts** | Pointer to **NullableString** |  | [optional] 
**PredecessorType** | Pointer to **NullableString** |  | [optional] 
**PredecessorId** | Pointer to **NullableInt32** |  | [optional] 
**PredecessorClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**LagDays** | Pointer to **NullableInt32** |  | [optional] 
**LagNonworkingDaysFlag** | Pointer to **NullableBool** |  | [optional] 
**EstimatedStartDate** | Pointer to **time.Time** |  | [optional] 
**Duration** | Pointer to **NullableInt32** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**Sla** | Pointer to [**SLAReference**](SLAReference.md) |  | [optional] 
**SlaStatus** | Pointer to **string** |  | [optional] 
**RequestForChangeFlag** | Pointer to **NullableBool** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**MergedParentTicket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**IntegratorTags** | Pointer to **[]string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**EscalationStartDateUTC** | Pointer to **string** |  | [optional] 
**EscalationLevel** | Pointer to **NullableInt32** |  | [optional] 
**MinutesBeforeWaiting** | Pointer to **NullableInt32** |  | [optional] 
**RespondedSkippedMinutes** | Pointer to **NullableInt32** |  | [optional] 
**ResplanSkippedMinutes** | Pointer to **NullableInt32** |  | [optional] 
**RespondedHours** | Pointer to **NullableFloat64** |  | [optional] 
**RespondedBy** | Pointer to **string** |  | [optional] 
**ResplanHours** | Pointer to **NullableFloat64** |  | [optional] 
**ResplanBy** | Pointer to **string** |  | [optional] 
**ResolutionHours** | Pointer to **NullableFloat64** |  | [optional] 
**ResolvedBy** | Pointer to **string** |  | [optional] 
**MinutesWaiting** | Pointer to **NullableInt32** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewTicket

`func NewTicket(summary string, company CompanyReference, ) *Ticket`

NewTicket instantiates a new Ticket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketWithDefaults

`func NewTicketWithDefaults() *Ticket`

NewTicketWithDefaults instantiates a new Ticket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ticket) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ticket) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ticket) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Ticket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSummary

`func (o *Ticket) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Ticket) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Ticket) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetRecordType

`func (o *Ticket) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *Ticket) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *Ticket) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *Ticket) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### SetRecordTypeNil

`func (o *Ticket) SetRecordTypeNil(b bool)`

 SetRecordTypeNil sets the value for RecordType to be an explicit nil

### UnsetRecordType
`func (o *Ticket) UnsetRecordType()`

UnsetRecordType ensures that no value is present for RecordType, not even an explicit nil
### GetBoard

`func (o *Ticket) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *Ticket) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *Ticket) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *Ticket) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetStatus

`func (o *Ticket) GetStatus() ServiceStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ticket) GetStatusOk() (*ServiceStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ticket) SetStatus(v ServiceStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Ticket) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkRole

`func (o *Ticket) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *Ticket) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *Ticket) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *Ticket) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *Ticket) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *Ticket) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *Ticket) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *Ticket) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetCompany

`func (o *Ticket) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Ticket) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Ticket) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetSite

`func (o *Ticket) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Ticket) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Ticket) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *Ticket) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSiteName

`func (o *Ticket) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *Ticket) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *Ticket) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *Ticket) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *Ticket) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *Ticket) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *Ticket) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *Ticket) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *Ticket) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *Ticket) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *Ticket) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *Ticket) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *Ticket) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Ticket) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Ticket) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Ticket) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStateIdentifier

`func (o *Ticket) GetStateIdentifier() string`

GetStateIdentifier returns the StateIdentifier field if non-nil, zero value otherwise.

### GetStateIdentifierOk

`func (o *Ticket) GetStateIdentifierOk() (*string, bool)`

GetStateIdentifierOk returns a tuple with the StateIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateIdentifier

`func (o *Ticket) SetStateIdentifier(v string)`

SetStateIdentifier sets StateIdentifier field to given value.

### HasStateIdentifier

`func (o *Ticket) HasStateIdentifier() bool`

HasStateIdentifier returns a boolean if a field has been set.

### GetZip

`func (o *Ticket) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *Ticket) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *Ticket) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *Ticket) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetCountry

`func (o *Ticket) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Ticket) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Ticket) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Ticket) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetContact

`func (o *Ticket) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Ticket) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Ticket) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Ticket) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetContactName

`func (o *Ticket) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *Ticket) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *Ticket) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *Ticket) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactPhoneNumber

`func (o *Ticket) GetContactPhoneNumber() string`

GetContactPhoneNumber returns the ContactPhoneNumber field if non-nil, zero value otherwise.

### GetContactPhoneNumberOk

`func (o *Ticket) GetContactPhoneNumberOk() (*string, bool)`

GetContactPhoneNumberOk returns a tuple with the ContactPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhoneNumber

`func (o *Ticket) SetContactPhoneNumber(v string)`

SetContactPhoneNumber sets ContactPhoneNumber field to given value.

### HasContactPhoneNumber

`func (o *Ticket) HasContactPhoneNumber() bool`

HasContactPhoneNumber returns a boolean if a field has been set.

### GetContactPhoneExtension

`func (o *Ticket) GetContactPhoneExtension() string`

GetContactPhoneExtension returns the ContactPhoneExtension field if non-nil, zero value otherwise.

### GetContactPhoneExtensionOk

`func (o *Ticket) GetContactPhoneExtensionOk() (*string, bool)`

GetContactPhoneExtensionOk returns a tuple with the ContactPhoneExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhoneExtension

`func (o *Ticket) SetContactPhoneExtension(v string)`

SetContactPhoneExtension sets ContactPhoneExtension field to given value.

### HasContactPhoneExtension

`func (o *Ticket) HasContactPhoneExtension() bool`

HasContactPhoneExtension returns a boolean if a field has been set.

### GetContactEmailAddress

`func (o *Ticket) GetContactEmailAddress() string`

GetContactEmailAddress returns the ContactEmailAddress field if non-nil, zero value otherwise.

### GetContactEmailAddressOk

`func (o *Ticket) GetContactEmailAddressOk() (*string, bool)`

GetContactEmailAddressOk returns a tuple with the ContactEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmailAddress

`func (o *Ticket) SetContactEmailAddress(v string)`

SetContactEmailAddress sets ContactEmailAddress field to given value.

### HasContactEmailAddress

`func (o *Ticket) HasContactEmailAddress() bool`

HasContactEmailAddress returns a boolean if a field has been set.

### GetType

`func (o *Ticket) GetType() ServiceTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ticket) GetTypeOk() (*ServiceTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Ticket) SetType(v ServiceTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *Ticket) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubType

`func (o *Ticket) GetSubType() ServiceSubTypeReference`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Ticket) GetSubTypeOk() (*ServiceSubTypeReference, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Ticket) SetSubType(v ServiceSubTypeReference)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *Ticket) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetItem

`func (o *Ticket) GetItem() ServiceItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *Ticket) GetItemOk() (*ServiceItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *Ticket) SetItem(v ServiceItemReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *Ticket) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetTeam

`func (o *Ticket) GetTeam() ServiceTeamReference`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *Ticket) GetTeamOk() (*ServiceTeamReference, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *Ticket) SetTeam(v ServiceTeamReference)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *Ticket) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetOwner

`func (o *Ticket) GetOwner() MemberReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Ticket) GetOwnerOk() (*MemberReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Ticket) SetOwner(v MemberReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Ticket) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPriority

`func (o *Ticket) GetPriority() PriorityReference`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Ticket) GetPriorityOk() (*PriorityReference, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Ticket) SetPriority(v PriorityReference)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Ticket) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetServiceLocation

`func (o *Ticket) GetServiceLocation() ServiceLocationReference`

GetServiceLocation returns the ServiceLocation field if non-nil, zero value otherwise.

### GetServiceLocationOk

`func (o *Ticket) GetServiceLocationOk() (*ServiceLocationReference, bool)`

GetServiceLocationOk returns a tuple with the ServiceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLocation

`func (o *Ticket) SetServiceLocation(v ServiceLocationReference)`

SetServiceLocation sets ServiceLocation field to given value.

### HasServiceLocation

`func (o *Ticket) HasServiceLocation() bool`

HasServiceLocation returns a boolean if a field has been set.

### GetSource

`func (o *Ticket) GetSource() ServiceSourceReference`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Ticket) GetSourceOk() (*ServiceSourceReference, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Ticket) SetSource(v ServiceSourceReference)`

SetSource sets Source field to given value.

### HasSource

`func (o *Ticket) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetRequiredDate

`func (o *Ticket) GetRequiredDate() time.Time`

GetRequiredDate returns the RequiredDate field if non-nil, zero value otherwise.

### GetRequiredDateOk

`func (o *Ticket) GetRequiredDateOk() (*time.Time, bool)`

GetRequiredDateOk returns a tuple with the RequiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDate

`func (o *Ticket) SetRequiredDate(v time.Time)`

SetRequiredDate sets RequiredDate field to given value.

### HasRequiredDate

`func (o *Ticket) HasRequiredDate() bool`

HasRequiredDate returns a boolean if a field has been set.

### GetBudgetHours

`func (o *Ticket) GetBudgetHours() float64`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *Ticket) GetBudgetHoursOk() (*float64, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *Ticket) SetBudgetHours(v float64)`

SetBudgetHours sets BudgetHours field to given value.

### HasBudgetHours

`func (o *Ticket) HasBudgetHours() bool`

HasBudgetHours returns a boolean if a field has been set.

### SetBudgetHoursNil

`func (o *Ticket) SetBudgetHoursNil(b bool)`

 SetBudgetHoursNil sets the value for BudgetHours to be an explicit nil

### UnsetBudgetHours
`func (o *Ticket) UnsetBudgetHours()`

UnsetBudgetHours ensures that no value is present for BudgetHours, not even an explicit nil
### GetOpportunity

`func (o *Ticket) GetOpportunity() OpportunityReference`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *Ticket) GetOpportunityOk() (*OpportunityReference, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *Ticket) SetOpportunity(v OpportunityReference)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *Ticket) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### GetAgreement

`func (o *Ticket) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *Ticket) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *Ticket) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *Ticket) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetAgreementType

`func (o *Ticket) GetAgreementType() string`

GetAgreementType returns the AgreementType field if non-nil, zero value otherwise.

### GetAgreementTypeOk

`func (o *Ticket) GetAgreementTypeOk() (*string, bool)`

GetAgreementTypeOk returns a tuple with the AgreementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementType

`func (o *Ticket) SetAgreementType(v string)`

SetAgreementType sets AgreementType field to given value.

### HasAgreementType

`func (o *Ticket) HasAgreementType() bool`

HasAgreementType returns a boolean if a field has been set.

### GetSeverity

`func (o *Ticket) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Ticket) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Ticket) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Ticket) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *Ticket) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *Ticket) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetImpact

`func (o *Ticket) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *Ticket) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *Ticket) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *Ticket) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### SetImpactNil

`func (o *Ticket) SetImpactNil(b bool)`

 SetImpactNil sets the value for Impact to be an explicit nil

### UnsetImpact
`func (o *Ticket) UnsetImpact()`

UnsetImpact ensures that no value is present for Impact, not even an explicit nil
### GetExternalXRef

`func (o *Ticket) GetExternalXRef() string`

GetExternalXRef returns the ExternalXRef field if non-nil, zero value otherwise.

### GetExternalXRefOk

`func (o *Ticket) GetExternalXRefOk() (*string, bool)`

GetExternalXRefOk returns a tuple with the ExternalXRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalXRef

`func (o *Ticket) SetExternalXRef(v string)`

SetExternalXRef sets ExternalXRef field to given value.

### HasExternalXRef

`func (o *Ticket) HasExternalXRef() bool`

HasExternalXRef returns a boolean if a field has been set.

### GetPoNumber

`func (o *Ticket) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *Ticket) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *Ticket) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *Ticket) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetKnowledgeBaseCategoryId

`func (o *Ticket) GetKnowledgeBaseCategoryId() int32`

GetKnowledgeBaseCategoryId returns the KnowledgeBaseCategoryId field if non-nil, zero value otherwise.

### GetKnowledgeBaseCategoryIdOk

`func (o *Ticket) GetKnowledgeBaseCategoryIdOk() (*int32, bool)`

GetKnowledgeBaseCategoryIdOk returns a tuple with the KnowledgeBaseCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBaseCategoryId

`func (o *Ticket) SetKnowledgeBaseCategoryId(v int32)`

SetKnowledgeBaseCategoryId sets KnowledgeBaseCategoryId field to given value.

### HasKnowledgeBaseCategoryId

`func (o *Ticket) HasKnowledgeBaseCategoryId() bool`

HasKnowledgeBaseCategoryId returns a boolean if a field has been set.

### SetKnowledgeBaseCategoryIdNil

`func (o *Ticket) SetKnowledgeBaseCategoryIdNil(b bool)`

 SetKnowledgeBaseCategoryIdNil sets the value for KnowledgeBaseCategoryId to be an explicit nil

### UnsetKnowledgeBaseCategoryId
`func (o *Ticket) UnsetKnowledgeBaseCategoryId()`

UnsetKnowledgeBaseCategoryId ensures that no value is present for KnowledgeBaseCategoryId, not even an explicit nil
### GetKnowledgeBaseSubCategoryId

`func (o *Ticket) GetKnowledgeBaseSubCategoryId() int32`

GetKnowledgeBaseSubCategoryId returns the KnowledgeBaseSubCategoryId field if non-nil, zero value otherwise.

### GetKnowledgeBaseSubCategoryIdOk

`func (o *Ticket) GetKnowledgeBaseSubCategoryIdOk() (*int32, bool)`

GetKnowledgeBaseSubCategoryIdOk returns a tuple with the KnowledgeBaseSubCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBaseSubCategoryId

`func (o *Ticket) SetKnowledgeBaseSubCategoryId(v int32)`

SetKnowledgeBaseSubCategoryId sets KnowledgeBaseSubCategoryId field to given value.

### HasKnowledgeBaseSubCategoryId

`func (o *Ticket) HasKnowledgeBaseSubCategoryId() bool`

HasKnowledgeBaseSubCategoryId returns a boolean if a field has been set.

### SetKnowledgeBaseSubCategoryIdNil

`func (o *Ticket) SetKnowledgeBaseSubCategoryIdNil(b bool)`

 SetKnowledgeBaseSubCategoryIdNil sets the value for KnowledgeBaseSubCategoryId to be an explicit nil

### UnsetKnowledgeBaseSubCategoryId
`func (o *Ticket) UnsetKnowledgeBaseSubCategoryId()`

UnsetKnowledgeBaseSubCategoryId ensures that no value is present for KnowledgeBaseSubCategoryId, not even an explicit nil
### GetAllowAllClientsPortalView

`func (o *Ticket) GetAllowAllClientsPortalView() bool`

GetAllowAllClientsPortalView returns the AllowAllClientsPortalView field if non-nil, zero value otherwise.

### GetAllowAllClientsPortalViewOk

`func (o *Ticket) GetAllowAllClientsPortalViewOk() (*bool, bool)`

GetAllowAllClientsPortalViewOk returns a tuple with the AllowAllClientsPortalView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAllClientsPortalView

`func (o *Ticket) SetAllowAllClientsPortalView(v bool)`

SetAllowAllClientsPortalView sets AllowAllClientsPortalView field to given value.

### HasAllowAllClientsPortalView

`func (o *Ticket) HasAllowAllClientsPortalView() bool`

HasAllowAllClientsPortalView returns a boolean if a field has been set.

### SetAllowAllClientsPortalViewNil

`func (o *Ticket) SetAllowAllClientsPortalViewNil(b bool)`

 SetAllowAllClientsPortalViewNil sets the value for AllowAllClientsPortalView to be an explicit nil

### UnsetAllowAllClientsPortalView
`func (o *Ticket) UnsetAllowAllClientsPortalView()`

UnsetAllowAllClientsPortalView ensures that no value is present for AllowAllClientsPortalView, not even an explicit nil
### GetCustomerUpdatedFlag

`func (o *Ticket) GetCustomerUpdatedFlag() bool`

GetCustomerUpdatedFlag returns the CustomerUpdatedFlag field if non-nil, zero value otherwise.

### GetCustomerUpdatedFlagOk

`func (o *Ticket) GetCustomerUpdatedFlagOk() (*bool, bool)`

GetCustomerUpdatedFlagOk returns a tuple with the CustomerUpdatedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUpdatedFlag

`func (o *Ticket) SetCustomerUpdatedFlag(v bool)`

SetCustomerUpdatedFlag sets CustomerUpdatedFlag field to given value.

### HasCustomerUpdatedFlag

`func (o *Ticket) HasCustomerUpdatedFlag() bool`

HasCustomerUpdatedFlag returns a boolean if a field has been set.

### SetCustomerUpdatedFlagNil

`func (o *Ticket) SetCustomerUpdatedFlagNil(b bool)`

 SetCustomerUpdatedFlagNil sets the value for CustomerUpdatedFlag to be an explicit nil

### UnsetCustomerUpdatedFlag
`func (o *Ticket) UnsetCustomerUpdatedFlag()`

UnsetCustomerUpdatedFlag ensures that no value is present for CustomerUpdatedFlag, not even an explicit nil
### GetAutomaticEmailContactFlag

`func (o *Ticket) GetAutomaticEmailContactFlag() bool`

GetAutomaticEmailContactFlag returns the AutomaticEmailContactFlag field if non-nil, zero value otherwise.

### GetAutomaticEmailContactFlagOk

`func (o *Ticket) GetAutomaticEmailContactFlagOk() (*bool, bool)`

GetAutomaticEmailContactFlagOk returns a tuple with the AutomaticEmailContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticEmailContactFlag

`func (o *Ticket) SetAutomaticEmailContactFlag(v bool)`

SetAutomaticEmailContactFlag sets AutomaticEmailContactFlag field to given value.

### HasAutomaticEmailContactFlag

`func (o *Ticket) HasAutomaticEmailContactFlag() bool`

HasAutomaticEmailContactFlag returns a boolean if a field has been set.

### SetAutomaticEmailContactFlagNil

`func (o *Ticket) SetAutomaticEmailContactFlagNil(b bool)`

 SetAutomaticEmailContactFlagNil sets the value for AutomaticEmailContactFlag to be an explicit nil

### UnsetAutomaticEmailContactFlag
`func (o *Ticket) UnsetAutomaticEmailContactFlag()`

UnsetAutomaticEmailContactFlag ensures that no value is present for AutomaticEmailContactFlag, not even an explicit nil
### GetAutomaticEmailResourceFlag

`func (o *Ticket) GetAutomaticEmailResourceFlag() bool`

GetAutomaticEmailResourceFlag returns the AutomaticEmailResourceFlag field if non-nil, zero value otherwise.

### GetAutomaticEmailResourceFlagOk

`func (o *Ticket) GetAutomaticEmailResourceFlagOk() (*bool, bool)`

GetAutomaticEmailResourceFlagOk returns a tuple with the AutomaticEmailResourceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticEmailResourceFlag

`func (o *Ticket) SetAutomaticEmailResourceFlag(v bool)`

SetAutomaticEmailResourceFlag sets AutomaticEmailResourceFlag field to given value.

### HasAutomaticEmailResourceFlag

`func (o *Ticket) HasAutomaticEmailResourceFlag() bool`

HasAutomaticEmailResourceFlag returns a boolean if a field has been set.

### SetAutomaticEmailResourceFlagNil

`func (o *Ticket) SetAutomaticEmailResourceFlagNil(b bool)`

 SetAutomaticEmailResourceFlagNil sets the value for AutomaticEmailResourceFlag to be an explicit nil

### UnsetAutomaticEmailResourceFlag
`func (o *Ticket) UnsetAutomaticEmailResourceFlag()`

UnsetAutomaticEmailResourceFlag ensures that no value is present for AutomaticEmailResourceFlag, not even an explicit nil
### GetAutomaticEmailCcFlag

`func (o *Ticket) GetAutomaticEmailCcFlag() bool`

GetAutomaticEmailCcFlag returns the AutomaticEmailCcFlag field if non-nil, zero value otherwise.

### GetAutomaticEmailCcFlagOk

`func (o *Ticket) GetAutomaticEmailCcFlagOk() (*bool, bool)`

GetAutomaticEmailCcFlagOk returns a tuple with the AutomaticEmailCcFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticEmailCcFlag

`func (o *Ticket) SetAutomaticEmailCcFlag(v bool)`

SetAutomaticEmailCcFlag sets AutomaticEmailCcFlag field to given value.

### HasAutomaticEmailCcFlag

`func (o *Ticket) HasAutomaticEmailCcFlag() bool`

HasAutomaticEmailCcFlag returns a boolean if a field has been set.

### SetAutomaticEmailCcFlagNil

`func (o *Ticket) SetAutomaticEmailCcFlagNil(b bool)`

 SetAutomaticEmailCcFlagNil sets the value for AutomaticEmailCcFlag to be an explicit nil

### UnsetAutomaticEmailCcFlag
`func (o *Ticket) UnsetAutomaticEmailCcFlag()`

UnsetAutomaticEmailCcFlag ensures that no value is present for AutomaticEmailCcFlag, not even an explicit nil
### GetAutomaticEmailCc

`func (o *Ticket) GetAutomaticEmailCc() string`

GetAutomaticEmailCc returns the AutomaticEmailCc field if non-nil, zero value otherwise.

### GetAutomaticEmailCcOk

`func (o *Ticket) GetAutomaticEmailCcOk() (*string, bool)`

GetAutomaticEmailCcOk returns a tuple with the AutomaticEmailCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticEmailCc

`func (o *Ticket) SetAutomaticEmailCc(v string)`

SetAutomaticEmailCc sets AutomaticEmailCc field to given value.

### HasAutomaticEmailCc

`func (o *Ticket) HasAutomaticEmailCc() bool`

HasAutomaticEmailCc returns a boolean if a field has been set.

### GetInitialDescription

`func (o *Ticket) GetInitialDescription() string`

GetInitialDescription returns the InitialDescription field if non-nil, zero value otherwise.

### GetInitialDescriptionOk

`func (o *Ticket) GetInitialDescriptionOk() (*string, bool)`

GetInitialDescriptionOk returns a tuple with the InitialDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDescription

`func (o *Ticket) SetInitialDescription(v string)`

SetInitialDescription sets InitialDescription field to given value.

### HasInitialDescription

`func (o *Ticket) HasInitialDescription() bool`

HasInitialDescription returns a boolean if a field has been set.

### GetInitialInternalAnalysis

`func (o *Ticket) GetInitialInternalAnalysis() string`

GetInitialInternalAnalysis returns the InitialInternalAnalysis field if non-nil, zero value otherwise.

### GetInitialInternalAnalysisOk

`func (o *Ticket) GetInitialInternalAnalysisOk() (*string, bool)`

GetInitialInternalAnalysisOk returns a tuple with the InitialInternalAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialInternalAnalysis

`func (o *Ticket) SetInitialInternalAnalysis(v string)`

SetInitialInternalAnalysis sets InitialInternalAnalysis field to given value.

### HasInitialInternalAnalysis

`func (o *Ticket) HasInitialInternalAnalysis() bool`

HasInitialInternalAnalysis returns a boolean if a field has been set.

### GetInitialResolution

`func (o *Ticket) GetInitialResolution() string`

GetInitialResolution returns the InitialResolution field if non-nil, zero value otherwise.

### GetInitialResolutionOk

`func (o *Ticket) GetInitialResolutionOk() (*string, bool)`

GetInitialResolutionOk returns a tuple with the InitialResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialResolution

`func (o *Ticket) SetInitialResolution(v string)`

SetInitialResolution sets InitialResolution field to given value.

### HasInitialResolution

`func (o *Ticket) HasInitialResolution() bool`

HasInitialResolution returns a boolean if a field has been set.

### GetInitialDescriptionFrom

`func (o *Ticket) GetInitialDescriptionFrom() string`

GetInitialDescriptionFrom returns the InitialDescriptionFrom field if non-nil, zero value otherwise.

### GetInitialDescriptionFromOk

`func (o *Ticket) GetInitialDescriptionFromOk() (*string, bool)`

GetInitialDescriptionFromOk returns a tuple with the InitialDescriptionFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDescriptionFrom

`func (o *Ticket) SetInitialDescriptionFrom(v string)`

SetInitialDescriptionFrom sets InitialDescriptionFrom field to given value.

### HasInitialDescriptionFrom

`func (o *Ticket) HasInitialDescriptionFrom() bool`

HasInitialDescriptionFrom returns a boolean if a field has been set.

### GetContactEmailLookup

`func (o *Ticket) GetContactEmailLookup() string`

GetContactEmailLookup returns the ContactEmailLookup field if non-nil, zero value otherwise.

### GetContactEmailLookupOk

`func (o *Ticket) GetContactEmailLookupOk() (*string, bool)`

GetContactEmailLookupOk returns a tuple with the ContactEmailLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmailLookup

`func (o *Ticket) SetContactEmailLookup(v string)`

SetContactEmailLookup sets ContactEmailLookup field to given value.

### HasContactEmailLookup

`func (o *Ticket) HasContactEmailLookup() bool`

HasContactEmailLookup returns a boolean if a field has been set.

### GetProcessNotifications

`func (o *Ticket) GetProcessNotifications() bool`

GetProcessNotifications returns the ProcessNotifications field if non-nil, zero value otherwise.

### GetProcessNotificationsOk

`func (o *Ticket) GetProcessNotificationsOk() (*bool, bool)`

GetProcessNotificationsOk returns a tuple with the ProcessNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessNotifications

`func (o *Ticket) SetProcessNotifications(v bool)`

SetProcessNotifications sets ProcessNotifications field to given value.

### HasProcessNotifications

`func (o *Ticket) HasProcessNotifications() bool`

HasProcessNotifications returns a boolean if a field has been set.

### SetProcessNotificationsNil

`func (o *Ticket) SetProcessNotificationsNil(b bool)`

 SetProcessNotificationsNil sets the value for ProcessNotifications to be an explicit nil

### UnsetProcessNotifications
`func (o *Ticket) UnsetProcessNotifications()`

UnsetProcessNotifications ensures that no value is present for ProcessNotifications, not even an explicit nil
### GetSkipCallback

`func (o *Ticket) GetSkipCallback() bool`

GetSkipCallback returns the SkipCallback field if non-nil, zero value otherwise.

### GetSkipCallbackOk

`func (o *Ticket) GetSkipCallbackOk() (*bool, bool)`

GetSkipCallbackOk returns a tuple with the SkipCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCallback

`func (o *Ticket) SetSkipCallback(v bool)`

SetSkipCallback sets SkipCallback field to given value.

### HasSkipCallback

`func (o *Ticket) HasSkipCallback() bool`

HasSkipCallback returns a boolean if a field has been set.

### SetSkipCallbackNil

`func (o *Ticket) SetSkipCallbackNil(b bool)`

 SetSkipCallbackNil sets the value for SkipCallback to be an explicit nil

### UnsetSkipCallback
`func (o *Ticket) UnsetSkipCallback()`

UnsetSkipCallback ensures that no value is present for SkipCallback, not even an explicit nil
### GetClosedDate

`func (o *Ticket) GetClosedDate() string`

GetClosedDate returns the ClosedDate field if non-nil, zero value otherwise.

### GetClosedDateOk

`func (o *Ticket) GetClosedDateOk() (*string, bool)`

GetClosedDateOk returns a tuple with the ClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDate

`func (o *Ticket) SetClosedDate(v string)`

SetClosedDate sets ClosedDate field to given value.

### HasClosedDate

`func (o *Ticket) HasClosedDate() bool`

HasClosedDate returns a boolean if a field has been set.

### GetClosedBy

`func (o *Ticket) GetClosedBy() string`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *Ticket) GetClosedByOk() (*string, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *Ticket) SetClosedBy(v string)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *Ticket) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetClosedFlag

`func (o *Ticket) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *Ticket) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *Ticket) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *Ticket) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *Ticket) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *Ticket) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetActualHours

`func (o *Ticket) GetActualHours() float64`

GetActualHours returns the ActualHours field if non-nil, zero value otherwise.

### GetActualHoursOk

`func (o *Ticket) GetActualHoursOk() (*float64, bool)`

GetActualHoursOk returns a tuple with the ActualHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualHours

`func (o *Ticket) SetActualHours(v float64)`

SetActualHours sets ActualHours field to given value.

### HasActualHours

`func (o *Ticket) HasActualHours() bool`

HasActualHours returns a boolean if a field has been set.

### SetActualHoursNil

`func (o *Ticket) SetActualHoursNil(b bool)`

 SetActualHoursNil sets the value for ActualHours to be an explicit nil

### UnsetActualHours
`func (o *Ticket) UnsetActualHours()`

UnsetActualHours ensures that no value is present for ActualHours, not even an explicit nil
### GetApproved

`func (o *Ticket) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *Ticket) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *Ticket) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *Ticket) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### SetApprovedNil

`func (o *Ticket) SetApprovedNil(b bool)`

 SetApprovedNil sets the value for Approved to be an explicit nil

### UnsetApproved
`func (o *Ticket) UnsetApproved()`

UnsetApproved ensures that no value is present for Approved, not even an explicit nil
### GetEstimatedExpenseCost

`func (o *Ticket) GetEstimatedExpenseCost() float64`

GetEstimatedExpenseCost returns the EstimatedExpenseCost field if non-nil, zero value otherwise.

### GetEstimatedExpenseCostOk

`func (o *Ticket) GetEstimatedExpenseCostOk() (*float64, bool)`

GetEstimatedExpenseCostOk returns a tuple with the EstimatedExpenseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedExpenseCost

`func (o *Ticket) SetEstimatedExpenseCost(v float64)`

SetEstimatedExpenseCost sets EstimatedExpenseCost field to given value.

### HasEstimatedExpenseCost

`func (o *Ticket) HasEstimatedExpenseCost() bool`

HasEstimatedExpenseCost returns a boolean if a field has been set.

### SetEstimatedExpenseCostNil

`func (o *Ticket) SetEstimatedExpenseCostNil(b bool)`

 SetEstimatedExpenseCostNil sets the value for EstimatedExpenseCost to be an explicit nil

### UnsetEstimatedExpenseCost
`func (o *Ticket) UnsetEstimatedExpenseCost()`

UnsetEstimatedExpenseCost ensures that no value is present for EstimatedExpenseCost, not even an explicit nil
### GetEstimatedExpenseRevenue

`func (o *Ticket) GetEstimatedExpenseRevenue() float64`

GetEstimatedExpenseRevenue returns the EstimatedExpenseRevenue field if non-nil, zero value otherwise.

### GetEstimatedExpenseRevenueOk

`func (o *Ticket) GetEstimatedExpenseRevenueOk() (*float64, bool)`

GetEstimatedExpenseRevenueOk returns a tuple with the EstimatedExpenseRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedExpenseRevenue

`func (o *Ticket) SetEstimatedExpenseRevenue(v float64)`

SetEstimatedExpenseRevenue sets EstimatedExpenseRevenue field to given value.

### HasEstimatedExpenseRevenue

`func (o *Ticket) HasEstimatedExpenseRevenue() bool`

HasEstimatedExpenseRevenue returns a boolean if a field has been set.

### SetEstimatedExpenseRevenueNil

`func (o *Ticket) SetEstimatedExpenseRevenueNil(b bool)`

 SetEstimatedExpenseRevenueNil sets the value for EstimatedExpenseRevenue to be an explicit nil

### UnsetEstimatedExpenseRevenue
`func (o *Ticket) UnsetEstimatedExpenseRevenue()`

UnsetEstimatedExpenseRevenue ensures that no value is present for EstimatedExpenseRevenue, not even an explicit nil
### GetEstimatedProductCost

`func (o *Ticket) GetEstimatedProductCost() float64`

GetEstimatedProductCost returns the EstimatedProductCost field if non-nil, zero value otherwise.

### GetEstimatedProductCostOk

`func (o *Ticket) GetEstimatedProductCostOk() (*float64, bool)`

GetEstimatedProductCostOk returns a tuple with the EstimatedProductCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedProductCost

`func (o *Ticket) SetEstimatedProductCost(v float64)`

SetEstimatedProductCost sets EstimatedProductCost field to given value.

### HasEstimatedProductCost

`func (o *Ticket) HasEstimatedProductCost() bool`

HasEstimatedProductCost returns a boolean if a field has been set.

### SetEstimatedProductCostNil

`func (o *Ticket) SetEstimatedProductCostNil(b bool)`

 SetEstimatedProductCostNil sets the value for EstimatedProductCost to be an explicit nil

### UnsetEstimatedProductCost
`func (o *Ticket) UnsetEstimatedProductCost()`

UnsetEstimatedProductCost ensures that no value is present for EstimatedProductCost, not even an explicit nil
### GetEstimatedProductRevenue

`func (o *Ticket) GetEstimatedProductRevenue() float64`

GetEstimatedProductRevenue returns the EstimatedProductRevenue field if non-nil, zero value otherwise.

### GetEstimatedProductRevenueOk

`func (o *Ticket) GetEstimatedProductRevenueOk() (*float64, bool)`

GetEstimatedProductRevenueOk returns a tuple with the EstimatedProductRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedProductRevenue

`func (o *Ticket) SetEstimatedProductRevenue(v float64)`

SetEstimatedProductRevenue sets EstimatedProductRevenue field to given value.

### HasEstimatedProductRevenue

`func (o *Ticket) HasEstimatedProductRevenue() bool`

HasEstimatedProductRevenue returns a boolean if a field has been set.

### SetEstimatedProductRevenueNil

`func (o *Ticket) SetEstimatedProductRevenueNil(b bool)`

 SetEstimatedProductRevenueNil sets the value for EstimatedProductRevenue to be an explicit nil

### UnsetEstimatedProductRevenue
`func (o *Ticket) UnsetEstimatedProductRevenue()`

UnsetEstimatedProductRevenue ensures that no value is present for EstimatedProductRevenue, not even an explicit nil
### GetEstimatedTimeCost

`func (o *Ticket) GetEstimatedTimeCost() float64`

GetEstimatedTimeCost returns the EstimatedTimeCost field if non-nil, zero value otherwise.

### GetEstimatedTimeCostOk

`func (o *Ticket) GetEstimatedTimeCostOk() (*float64, bool)`

GetEstimatedTimeCostOk returns a tuple with the EstimatedTimeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeCost

`func (o *Ticket) SetEstimatedTimeCost(v float64)`

SetEstimatedTimeCost sets EstimatedTimeCost field to given value.

### HasEstimatedTimeCost

`func (o *Ticket) HasEstimatedTimeCost() bool`

HasEstimatedTimeCost returns a boolean if a field has been set.

### SetEstimatedTimeCostNil

`func (o *Ticket) SetEstimatedTimeCostNil(b bool)`

 SetEstimatedTimeCostNil sets the value for EstimatedTimeCost to be an explicit nil

### UnsetEstimatedTimeCost
`func (o *Ticket) UnsetEstimatedTimeCost()`

UnsetEstimatedTimeCost ensures that no value is present for EstimatedTimeCost, not even an explicit nil
### GetEstimatedTimeRevenue

`func (o *Ticket) GetEstimatedTimeRevenue() float64`

GetEstimatedTimeRevenue returns the EstimatedTimeRevenue field if non-nil, zero value otherwise.

### GetEstimatedTimeRevenueOk

`func (o *Ticket) GetEstimatedTimeRevenueOk() (*float64, bool)`

GetEstimatedTimeRevenueOk returns a tuple with the EstimatedTimeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeRevenue

`func (o *Ticket) SetEstimatedTimeRevenue(v float64)`

SetEstimatedTimeRevenue sets EstimatedTimeRevenue field to given value.

### HasEstimatedTimeRevenue

`func (o *Ticket) HasEstimatedTimeRevenue() bool`

HasEstimatedTimeRevenue returns a boolean if a field has been set.

### SetEstimatedTimeRevenueNil

`func (o *Ticket) SetEstimatedTimeRevenueNil(b bool)`

 SetEstimatedTimeRevenueNil sets the value for EstimatedTimeRevenue to be an explicit nil

### UnsetEstimatedTimeRevenue
`func (o *Ticket) UnsetEstimatedTimeRevenue()`

UnsetEstimatedTimeRevenue ensures that no value is present for EstimatedTimeRevenue, not even an explicit nil
### GetBillingMethod

`func (o *Ticket) GetBillingMethod() string`

GetBillingMethod returns the BillingMethod field if non-nil, zero value otherwise.

### GetBillingMethodOk

`func (o *Ticket) GetBillingMethodOk() (*string, bool)`

GetBillingMethodOk returns a tuple with the BillingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMethod

`func (o *Ticket) SetBillingMethod(v string)`

SetBillingMethod sets BillingMethod field to given value.

### HasBillingMethod

`func (o *Ticket) HasBillingMethod() bool`

HasBillingMethod returns a boolean if a field has been set.

### SetBillingMethodNil

`func (o *Ticket) SetBillingMethodNil(b bool)`

 SetBillingMethodNil sets the value for BillingMethod to be an explicit nil

### UnsetBillingMethod
`func (o *Ticket) UnsetBillingMethod()`

UnsetBillingMethod ensures that no value is present for BillingMethod, not even an explicit nil
### GetBillingAmount

`func (o *Ticket) GetBillingAmount() float64`

GetBillingAmount returns the BillingAmount field if non-nil, zero value otherwise.

### GetBillingAmountOk

`func (o *Ticket) GetBillingAmountOk() (*float64, bool)`

GetBillingAmountOk returns a tuple with the BillingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAmount

`func (o *Ticket) SetBillingAmount(v float64)`

SetBillingAmount sets BillingAmount field to given value.

### HasBillingAmount

`func (o *Ticket) HasBillingAmount() bool`

HasBillingAmount returns a boolean if a field has been set.

### SetBillingAmountNil

`func (o *Ticket) SetBillingAmountNil(b bool)`

 SetBillingAmountNil sets the value for BillingAmount to be an explicit nil

### UnsetBillingAmount
`func (o *Ticket) UnsetBillingAmount()`

UnsetBillingAmount ensures that no value is present for BillingAmount, not even an explicit nil
### GetHourlyRate

`func (o *Ticket) GetHourlyRate() float64`

GetHourlyRate returns the HourlyRate field if non-nil, zero value otherwise.

### GetHourlyRateOk

`func (o *Ticket) GetHourlyRateOk() (*float64, bool)`

GetHourlyRateOk returns a tuple with the HourlyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyRate

`func (o *Ticket) SetHourlyRate(v float64)`

SetHourlyRate sets HourlyRate field to given value.

### HasHourlyRate

`func (o *Ticket) HasHourlyRate() bool`

HasHourlyRate returns a boolean if a field has been set.

### SetHourlyRateNil

`func (o *Ticket) SetHourlyRateNil(b bool)`

 SetHourlyRateNil sets the value for HourlyRate to be an explicit nil

### UnsetHourlyRate
`func (o *Ticket) UnsetHourlyRate()`

UnsetHourlyRate ensures that no value is present for HourlyRate, not even an explicit nil
### GetSubBillingMethod

`func (o *Ticket) GetSubBillingMethod() string`

GetSubBillingMethod returns the SubBillingMethod field if non-nil, zero value otherwise.

### GetSubBillingMethodOk

`func (o *Ticket) GetSubBillingMethodOk() (*string, bool)`

GetSubBillingMethodOk returns a tuple with the SubBillingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubBillingMethod

`func (o *Ticket) SetSubBillingMethod(v string)`

SetSubBillingMethod sets SubBillingMethod field to given value.

### HasSubBillingMethod

`func (o *Ticket) HasSubBillingMethod() bool`

HasSubBillingMethod returns a boolean if a field has been set.

### SetSubBillingMethodNil

`func (o *Ticket) SetSubBillingMethodNil(b bool)`

 SetSubBillingMethodNil sets the value for SubBillingMethod to be an explicit nil

### UnsetSubBillingMethod
`func (o *Ticket) UnsetSubBillingMethod()`

UnsetSubBillingMethod ensures that no value is present for SubBillingMethod, not even an explicit nil
### GetSubBillingAmount

`func (o *Ticket) GetSubBillingAmount() float64`

GetSubBillingAmount returns the SubBillingAmount field if non-nil, zero value otherwise.

### GetSubBillingAmountOk

`func (o *Ticket) GetSubBillingAmountOk() (*float64, bool)`

GetSubBillingAmountOk returns a tuple with the SubBillingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubBillingAmount

`func (o *Ticket) SetSubBillingAmount(v float64)`

SetSubBillingAmount sets SubBillingAmount field to given value.

### HasSubBillingAmount

`func (o *Ticket) HasSubBillingAmount() bool`

HasSubBillingAmount returns a boolean if a field has been set.

### SetSubBillingAmountNil

`func (o *Ticket) SetSubBillingAmountNil(b bool)`

 SetSubBillingAmountNil sets the value for SubBillingAmount to be an explicit nil

### UnsetSubBillingAmount
`func (o *Ticket) UnsetSubBillingAmount()`

UnsetSubBillingAmount ensures that no value is present for SubBillingAmount, not even an explicit nil
### GetSubDateAccepted

`func (o *Ticket) GetSubDateAccepted() string`

GetSubDateAccepted returns the SubDateAccepted field if non-nil, zero value otherwise.

### GetSubDateAcceptedOk

`func (o *Ticket) GetSubDateAcceptedOk() (*string, bool)`

GetSubDateAcceptedOk returns a tuple with the SubDateAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDateAccepted

`func (o *Ticket) SetSubDateAccepted(v string)`

SetSubDateAccepted sets SubDateAccepted field to given value.

### HasSubDateAccepted

`func (o *Ticket) HasSubDateAccepted() bool`

HasSubDateAccepted returns a boolean if a field has been set.

### GetDateResolved

`func (o *Ticket) GetDateResolved() string`

GetDateResolved returns the DateResolved field if non-nil, zero value otherwise.

### GetDateResolvedOk

`func (o *Ticket) GetDateResolvedOk() (*string, bool)`

GetDateResolvedOk returns a tuple with the DateResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateResolved

`func (o *Ticket) SetDateResolved(v string)`

SetDateResolved sets DateResolved field to given value.

### HasDateResolved

`func (o *Ticket) HasDateResolved() bool`

HasDateResolved returns a boolean if a field has been set.

### GetDateResplan

`func (o *Ticket) GetDateResplan() string`

GetDateResplan returns the DateResplan field if non-nil, zero value otherwise.

### GetDateResplanOk

`func (o *Ticket) GetDateResplanOk() (*string, bool)`

GetDateResplanOk returns a tuple with the DateResplan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateResplan

`func (o *Ticket) SetDateResplan(v string)`

SetDateResplan sets DateResplan field to given value.

### HasDateResplan

`func (o *Ticket) HasDateResplan() bool`

HasDateResplan returns a boolean if a field has been set.

### GetDateResponded

`func (o *Ticket) GetDateResponded() string`

GetDateResponded returns the DateResponded field if non-nil, zero value otherwise.

### GetDateRespondedOk

`func (o *Ticket) GetDateRespondedOk() (*string, bool)`

GetDateRespondedOk returns a tuple with the DateResponded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateResponded

`func (o *Ticket) SetDateResponded(v string)`

SetDateResponded sets DateResponded field to given value.

### HasDateResponded

`func (o *Ticket) HasDateResponded() bool`

HasDateResponded returns a boolean if a field has been set.

### GetResolveMinutes

`func (o *Ticket) GetResolveMinutes() int32`

GetResolveMinutes returns the ResolveMinutes field if non-nil, zero value otherwise.

### GetResolveMinutesOk

`func (o *Ticket) GetResolveMinutesOk() (*int32, bool)`

GetResolveMinutesOk returns a tuple with the ResolveMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveMinutes

`func (o *Ticket) SetResolveMinutes(v int32)`

SetResolveMinutes sets ResolveMinutes field to given value.

### HasResolveMinutes

`func (o *Ticket) HasResolveMinutes() bool`

HasResolveMinutes returns a boolean if a field has been set.

### SetResolveMinutesNil

`func (o *Ticket) SetResolveMinutesNil(b bool)`

 SetResolveMinutesNil sets the value for ResolveMinutes to be an explicit nil

### UnsetResolveMinutes
`func (o *Ticket) UnsetResolveMinutes()`

UnsetResolveMinutes ensures that no value is present for ResolveMinutes, not even an explicit nil
### GetResPlanMinutes

`func (o *Ticket) GetResPlanMinutes() int32`

GetResPlanMinutes returns the ResPlanMinutes field if non-nil, zero value otherwise.

### GetResPlanMinutesOk

`func (o *Ticket) GetResPlanMinutesOk() (*int32, bool)`

GetResPlanMinutesOk returns a tuple with the ResPlanMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResPlanMinutes

`func (o *Ticket) SetResPlanMinutes(v int32)`

SetResPlanMinutes sets ResPlanMinutes field to given value.

### HasResPlanMinutes

`func (o *Ticket) HasResPlanMinutes() bool`

HasResPlanMinutes returns a boolean if a field has been set.

### SetResPlanMinutesNil

`func (o *Ticket) SetResPlanMinutesNil(b bool)`

 SetResPlanMinutesNil sets the value for ResPlanMinutes to be an explicit nil

### UnsetResPlanMinutes
`func (o *Ticket) UnsetResPlanMinutes()`

UnsetResPlanMinutes ensures that no value is present for ResPlanMinutes, not even an explicit nil
### GetRespondMinutes

`func (o *Ticket) GetRespondMinutes() int32`

GetRespondMinutes returns the RespondMinutes field if non-nil, zero value otherwise.

### GetRespondMinutesOk

`func (o *Ticket) GetRespondMinutesOk() (*int32, bool)`

GetRespondMinutesOk returns a tuple with the RespondMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondMinutes

`func (o *Ticket) SetRespondMinutes(v int32)`

SetRespondMinutes sets RespondMinutes field to given value.

### HasRespondMinutes

`func (o *Ticket) HasRespondMinutes() bool`

HasRespondMinutes returns a boolean if a field has been set.

### SetRespondMinutesNil

`func (o *Ticket) SetRespondMinutesNil(b bool)`

 SetRespondMinutesNil sets the value for RespondMinutes to be an explicit nil

### UnsetRespondMinutes
`func (o *Ticket) UnsetRespondMinutes()`

UnsetRespondMinutes ensures that no value is present for RespondMinutes, not even an explicit nil
### GetIsInSla

`func (o *Ticket) GetIsInSla() bool`

GetIsInSla returns the IsInSla field if non-nil, zero value otherwise.

### GetIsInSlaOk

`func (o *Ticket) GetIsInSlaOk() (*bool, bool)`

GetIsInSlaOk returns a tuple with the IsInSla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInSla

`func (o *Ticket) SetIsInSla(v bool)`

SetIsInSla sets IsInSla field to given value.

### HasIsInSla

`func (o *Ticket) HasIsInSla() bool`

HasIsInSla returns a boolean if a field has been set.

### SetIsInSlaNil

`func (o *Ticket) SetIsInSlaNil(b bool)`

 SetIsInSlaNil sets the value for IsInSla to be an explicit nil

### UnsetIsInSla
`func (o *Ticket) UnsetIsInSla()`

UnsetIsInSla ensures that no value is present for IsInSla, not even an explicit nil
### GetKnowledgeBaseLinkId

`func (o *Ticket) GetKnowledgeBaseLinkId() int32`

GetKnowledgeBaseLinkId returns the KnowledgeBaseLinkId field if non-nil, zero value otherwise.

### GetKnowledgeBaseLinkIdOk

`func (o *Ticket) GetKnowledgeBaseLinkIdOk() (*int32, bool)`

GetKnowledgeBaseLinkIdOk returns a tuple with the KnowledgeBaseLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBaseLinkId

`func (o *Ticket) SetKnowledgeBaseLinkId(v int32)`

SetKnowledgeBaseLinkId sets KnowledgeBaseLinkId field to given value.

### HasKnowledgeBaseLinkId

`func (o *Ticket) HasKnowledgeBaseLinkId() bool`

HasKnowledgeBaseLinkId returns a boolean if a field has been set.

### SetKnowledgeBaseLinkIdNil

`func (o *Ticket) SetKnowledgeBaseLinkIdNil(b bool)`

 SetKnowledgeBaseLinkIdNil sets the value for KnowledgeBaseLinkId to be an explicit nil

### UnsetKnowledgeBaseLinkId
`func (o *Ticket) UnsetKnowledgeBaseLinkId()`

UnsetKnowledgeBaseLinkId ensures that no value is present for KnowledgeBaseLinkId, not even an explicit nil
### GetResources

`func (o *Ticket) GetResources() string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Ticket) GetResourcesOk() (*string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Ticket) SetResources(v string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Ticket) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetParentTicketId

`func (o *Ticket) GetParentTicketId() int32`

GetParentTicketId returns the ParentTicketId field if non-nil, zero value otherwise.

### GetParentTicketIdOk

`func (o *Ticket) GetParentTicketIdOk() (*int32, bool)`

GetParentTicketIdOk returns a tuple with the ParentTicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTicketId

`func (o *Ticket) SetParentTicketId(v int32)`

SetParentTicketId sets ParentTicketId field to given value.

### HasParentTicketId

`func (o *Ticket) HasParentTicketId() bool`

HasParentTicketId returns a boolean if a field has been set.

### SetParentTicketIdNil

`func (o *Ticket) SetParentTicketIdNil(b bool)`

 SetParentTicketIdNil sets the value for ParentTicketId to be an explicit nil

### UnsetParentTicketId
`func (o *Ticket) UnsetParentTicketId()`

UnsetParentTicketId ensures that no value is present for ParentTicketId, not even an explicit nil
### GetHasChildTicket

`func (o *Ticket) GetHasChildTicket() bool`

GetHasChildTicket returns the HasChildTicket field if non-nil, zero value otherwise.

### GetHasChildTicketOk

`func (o *Ticket) GetHasChildTicketOk() (*bool, bool)`

GetHasChildTicketOk returns a tuple with the HasChildTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildTicket

`func (o *Ticket) SetHasChildTicket(v bool)`

SetHasChildTicket sets HasChildTicket field to given value.

### HasHasChildTicket

`func (o *Ticket) HasHasChildTicket() bool`

HasHasChildTicket returns a boolean if a field has been set.

### SetHasChildTicketNil

`func (o *Ticket) SetHasChildTicketNil(b bool)`

 SetHasChildTicketNil sets the value for HasChildTicket to be an explicit nil

### UnsetHasChildTicket
`func (o *Ticket) UnsetHasChildTicket()`

UnsetHasChildTicket ensures that no value is present for HasChildTicket, not even an explicit nil
### GetHasMergedChildTicketFlag

`func (o *Ticket) GetHasMergedChildTicketFlag() bool`

GetHasMergedChildTicketFlag returns the HasMergedChildTicketFlag field if non-nil, zero value otherwise.

### GetHasMergedChildTicketFlagOk

`func (o *Ticket) GetHasMergedChildTicketFlagOk() (*bool, bool)`

GetHasMergedChildTicketFlagOk returns a tuple with the HasMergedChildTicketFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMergedChildTicketFlag

`func (o *Ticket) SetHasMergedChildTicketFlag(v bool)`

SetHasMergedChildTicketFlag sets HasMergedChildTicketFlag field to given value.

### HasHasMergedChildTicketFlag

`func (o *Ticket) HasHasMergedChildTicketFlag() bool`

HasHasMergedChildTicketFlag returns a boolean if a field has been set.

### SetHasMergedChildTicketFlagNil

`func (o *Ticket) SetHasMergedChildTicketFlagNil(b bool)`

 SetHasMergedChildTicketFlagNil sets the value for HasMergedChildTicketFlag to be an explicit nil

### UnsetHasMergedChildTicketFlag
`func (o *Ticket) UnsetHasMergedChildTicketFlag()`

UnsetHasMergedChildTicketFlag ensures that no value is present for HasMergedChildTicketFlag, not even an explicit nil
### GetKnowledgeBaseLinkType

`func (o *Ticket) GetKnowledgeBaseLinkType() string`

GetKnowledgeBaseLinkType returns the KnowledgeBaseLinkType field if non-nil, zero value otherwise.

### GetKnowledgeBaseLinkTypeOk

`func (o *Ticket) GetKnowledgeBaseLinkTypeOk() (*string, bool)`

GetKnowledgeBaseLinkTypeOk returns a tuple with the KnowledgeBaseLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBaseLinkType

`func (o *Ticket) SetKnowledgeBaseLinkType(v string)`

SetKnowledgeBaseLinkType sets KnowledgeBaseLinkType field to given value.

### HasKnowledgeBaseLinkType

`func (o *Ticket) HasKnowledgeBaseLinkType() bool`

HasKnowledgeBaseLinkType returns a boolean if a field has been set.

### SetKnowledgeBaseLinkTypeNil

`func (o *Ticket) SetKnowledgeBaseLinkTypeNil(b bool)`

 SetKnowledgeBaseLinkTypeNil sets the value for KnowledgeBaseLinkType to be an explicit nil

### UnsetKnowledgeBaseLinkType
`func (o *Ticket) UnsetKnowledgeBaseLinkType()`

UnsetKnowledgeBaseLinkType ensures that no value is present for KnowledgeBaseLinkType, not even an explicit nil
### GetBillTime

`func (o *Ticket) GetBillTime() string`

GetBillTime returns the BillTime field if non-nil, zero value otherwise.

### GetBillTimeOk

`func (o *Ticket) GetBillTimeOk() (*string, bool)`

GetBillTimeOk returns a tuple with the BillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTime

`func (o *Ticket) SetBillTime(v string)`

SetBillTime sets BillTime field to given value.

### HasBillTime

`func (o *Ticket) HasBillTime() bool`

HasBillTime returns a boolean if a field has been set.

### SetBillTimeNil

`func (o *Ticket) SetBillTimeNil(b bool)`

 SetBillTimeNil sets the value for BillTime to be an explicit nil

### UnsetBillTime
`func (o *Ticket) UnsetBillTime()`

UnsetBillTime ensures that no value is present for BillTime, not even an explicit nil
### GetBillExpenses

`func (o *Ticket) GetBillExpenses() string`

GetBillExpenses returns the BillExpenses field if non-nil, zero value otherwise.

### GetBillExpensesOk

`func (o *Ticket) GetBillExpensesOk() (*string, bool)`

GetBillExpensesOk returns a tuple with the BillExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillExpenses

`func (o *Ticket) SetBillExpenses(v string)`

SetBillExpenses sets BillExpenses field to given value.

### HasBillExpenses

`func (o *Ticket) HasBillExpenses() bool`

HasBillExpenses returns a boolean if a field has been set.

### SetBillExpensesNil

`func (o *Ticket) SetBillExpensesNil(b bool)`

 SetBillExpensesNil sets the value for BillExpenses to be an explicit nil

### UnsetBillExpenses
`func (o *Ticket) UnsetBillExpenses()`

UnsetBillExpenses ensures that no value is present for BillExpenses, not even an explicit nil
### GetBillProducts

`func (o *Ticket) GetBillProducts() string`

GetBillProducts returns the BillProducts field if non-nil, zero value otherwise.

### GetBillProductsOk

`func (o *Ticket) GetBillProductsOk() (*string, bool)`

GetBillProductsOk returns a tuple with the BillProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProducts

`func (o *Ticket) SetBillProducts(v string)`

SetBillProducts sets BillProducts field to given value.

### HasBillProducts

`func (o *Ticket) HasBillProducts() bool`

HasBillProducts returns a boolean if a field has been set.

### SetBillProductsNil

`func (o *Ticket) SetBillProductsNil(b bool)`

 SetBillProductsNil sets the value for BillProducts to be an explicit nil

### UnsetBillProducts
`func (o *Ticket) UnsetBillProducts()`

UnsetBillProducts ensures that no value is present for BillProducts, not even an explicit nil
### GetPredecessorType

`func (o *Ticket) GetPredecessorType() string`

GetPredecessorType returns the PredecessorType field if non-nil, zero value otherwise.

### GetPredecessorTypeOk

`func (o *Ticket) GetPredecessorTypeOk() (*string, bool)`

GetPredecessorTypeOk returns a tuple with the PredecessorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorType

`func (o *Ticket) SetPredecessorType(v string)`

SetPredecessorType sets PredecessorType field to given value.

### HasPredecessorType

`func (o *Ticket) HasPredecessorType() bool`

HasPredecessorType returns a boolean if a field has been set.

### SetPredecessorTypeNil

`func (o *Ticket) SetPredecessorTypeNil(b bool)`

 SetPredecessorTypeNil sets the value for PredecessorType to be an explicit nil

### UnsetPredecessorType
`func (o *Ticket) UnsetPredecessorType()`

UnsetPredecessorType ensures that no value is present for PredecessorType, not even an explicit nil
### GetPredecessorId

`func (o *Ticket) GetPredecessorId() int32`

GetPredecessorId returns the PredecessorId field if non-nil, zero value otherwise.

### GetPredecessorIdOk

`func (o *Ticket) GetPredecessorIdOk() (*int32, bool)`

GetPredecessorIdOk returns a tuple with the PredecessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorId

`func (o *Ticket) SetPredecessorId(v int32)`

SetPredecessorId sets PredecessorId field to given value.

### HasPredecessorId

`func (o *Ticket) HasPredecessorId() bool`

HasPredecessorId returns a boolean if a field has been set.

### SetPredecessorIdNil

`func (o *Ticket) SetPredecessorIdNil(b bool)`

 SetPredecessorIdNil sets the value for PredecessorId to be an explicit nil

### UnsetPredecessorId
`func (o *Ticket) UnsetPredecessorId()`

UnsetPredecessorId ensures that no value is present for PredecessorId, not even an explicit nil
### GetPredecessorClosedFlag

`func (o *Ticket) GetPredecessorClosedFlag() bool`

GetPredecessorClosedFlag returns the PredecessorClosedFlag field if non-nil, zero value otherwise.

### GetPredecessorClosedFlagOk

`func (o *Ticket) GetPredecessorClosedFlagOk() (*bool, bool)`

GetPredecessorClosedFlagOk returns a tuple with the PredecessorClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorClosedFlag

`func (o *Ticket) SetPredecessorClosedFlag(v bool)`

SetPredecessorClosedFlag sets PredecessorClosedFlag field to given value.

### HasPredecessorClosedFlag

`func (o *Ticket) HasPredecessorClosedFlag() bool`

HasPredecessorClosedFlag returns a boolean if a field has been set.

### SetPredecessorClosedFlagNil

`func (o *Ticket) SetPredecessorClosedFlagNil(b bool)`

 SetPredecessorClosedFlagNil sets the value for PredecessorClosedFlag to be an explicit nil

### UnsetPredecessorClosedFlag
`func (o *Ticket) UnsetPredecessorClosedFlag()`

UnsetPredecessorClosedFlag ensures that no value is present for PredecessorClosedFlag, not even an explicit nil
### GetLagDays

`func (o *Ticket) GetLagDays() int32`

GetLagDays returns the LagDays field if non-nil, zero value otherwise.

### GetLagDaysOk

`func (o *Ticket) GetLagDaysOk() (*int32, bool)`

GetLagDaysOk returns a tuple with the LagDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagDays

`func (o *Ticket) SetLagDays(v int32)`

SetLagDays sets LagDays field to given value.

### HasLagDays

`func (o *Ticket) HasLagDays() bool`

HasLagDays returns a boolean if a field has been set.

### SetLagDaysNil

`func (o *Ticket) SetLagDaysNil(b bool)`

 SetLagDaysNil sets the value for LagDays to be an explicit nil

### UnsetLagDays
`func (o *Ticket) UnsetLagDays()`

UnsetLagDays ensures that no value is present for LagDays, not even an explicit nil
### GetLagNonworkingDaysFlag

`func (o *Ticket) GetLagNonworkingDaysFlag() bool`

GetLagNonworkingDaysFlag returns the LagNonworkingDaysFlag field if non-nil, zero value otherwise.

### GetLagNonworkingDaysFlagOk

`func (o *Ticket) GetLagNonworkingDaysFlagOk() (*bool, bool)`

GetLagNonworkingDaysFlagOk returns a tuple with the LagNonworkingDaysFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagNonworkingDaysFlag

`func (o *Ticket) SetLagNonworkingDaysFlag(v bool)`

SetLagNonworkingDaysFlag sets LagNonworkingDaysFlag field to given value.

### HasLagNonworkingDaysFlag

`func (o *Ticket) HasLagNonworkingDaysFlag() bool`

HasLagNonworkingDaysFlag returns a boolean if a field has been set.

### SetLagNonworkingDaysFlagNil

`func (o *Ticket) SetLagNonworkingDaysFlagNil(b bool)`

 SetLagNonworkingDaysFlagNil sets the value for LagNonworkingDaysFlag to be an explicit nil

### UnsetLagNonworkingDaysFlag
`func (o *Ticket) UnsetLagNonworkingDaysFlag()`

UnsetLagNonworkingDaysFlag ensures that no value is present for LagNonworkingDaysFlag, not even an explicit nil
### GetEstimatedStartDate

`func (o *Ticket) GetEstimatedStartDate() time.Time`

GetEstimatedStartDate returns the EstimatedStartDate field if non-nil, zero value otherwise.

### GetEstimatedStartDateOk

`func (o *Ticket) GetEstimatedStartDateOk() (*time.Time, bool)`

GetEstimatedStartDateOk returns a tuple with the EstimatedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStartDate

`func (o *Ticket) SetEstimatedStartDate(v time.Time)`

SetEstimatedStartDate sets EstimatedStartDate field to given value.

### HasEstimatedStartDate

`func (o *Ticket) HasEstimatedStartDate() bool`

HasEstimatedStartDate returns a boolean if a field has been set.

### GetDuration

`func (o *Ticket) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Ticket) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Ticket) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Ticket) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *Ticket) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *Ticket) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetLocation

`func (o *Ticket) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Ticket) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Ticket) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Ticket) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *Ticket) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Ticket) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Ticket) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Ticket) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetMobileGuid

`func (o *Ticket) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *Ticket) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *Ticket) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *Ticket) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *Ticket) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *Ticket) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetSla

`func (o *Ticket) GetSla() SLAReference`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *Ticket) GetSlaOk() (*SLAReference, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *Ticket) SetSla(v SLAReference)`

SetSla sets Sla field to given value.

### HasSla

`func (o *Ticket) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetSlaStatus

`func (o *Ticket) GetSlaStatus() string`

GetSlaStatus returns the SlaStatus field if non-nil, zero value otherwise.

### GetSlaStatusOk

`func (o *Ticket) GetSlaStatusOk() (*string, bool)`

GetSlaStatusOk returns a tuple with the SlaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaStatus

`func (o *Ticket) SetSlaStatus(v string)`

SetSlaStatus sets SlaStatus field to given value.

### HasSlaStatus

`func (o *Ticket) HasSlaStatus() bool`

HasSlaStatus returns a boolean if a field has been set.

### GetRequestForChangeFlag

`func (o *Ticket) GetRequestForChangeFlag() bool`

GetRequestForChangeFlag returns the RequestForChangeFlag field if non-nil, zero value otherwise.

### GetRequestForChangeFlagOk

`func (o *Ticket) GetRequestForChangeFlagOk() (*bool, bool)`

GetRequestForChangeFlagOk returns a tuple with the RequestForChangeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestForChangeFlag

`func (o *Ticket) SetRequestForChangeFlag(v bool)`

SetRequestForChangeFlag sets RequestForChangeFlag field to given value.

### HasRequestForChangeFlag

`func (o *Ticket) HasRequestForChangeFlag() bool`

HasRequestForChangeFlag returns a boolean if a field has been set.

### SetRequestForChangeFlagNil

`func (o *Ticket) SetRequestForChangeFlagNil(b bool)`

 SetRequestForChangeFlagNil sets the value for RequestForChangeFlag to be an explicit nil

### UnsetRequestForChangeFlag
`func (o *Ticket) UnsetRequestForChangeFlag()`

UnsetRequestForChangeFlag ensures that no value is present for RequestForChangeFlag, not even an explicit nil
### GetCurrency

`func (o *Ticket) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Ticket) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Ticket) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Ticket) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMergedParentTicket

`func (o *Ticket) GetMergedParentTicket() TicketReference`

GetMergedParentTicket returns the MergedParentTicket field if non-nil, zero value otherwise.

### GetMergedParentTicketOk

`func (o *Ticket) GetMergedParentTicketOk() (*TicketReference, bool)`

GetMergedParentTicketOk returns a tuple with the MergedParentTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedParentTicket

`func (o *Ticket) SetMergedParentTicket(v TicketReference)`

SetMergedParentTicket sets MergedParentTicket field to given value.

### HasMergedParentTicket

`func (o *Ticket) HasMergedParentTicket() bool`

HasMergedParentTicket returns a boolean if a field has been set.

### GetIntegratorTags

`func (o *Ticket) GetIntegratorTags() []string`

GetIntegratorTags returns the IntegratorTags field if non-nil, zero value otherwise.

### GetIntegratorTagsOk

`func (o *Ticket) GetIntegratorTagsOk() (*[]string, bool)`

GetIntegratorTagsOk returns a tuple with the IntegratorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratorTags

`func (o *Ticket) SetIntegratorTags(v []string)`

SetIntegratorTags sets IntegratorTags field to given value.

### HasIntegratorTags

`func (o *Ticket) HasIntegratorTags() bool`

HasIntegratorTags returns a boolean if a field has been set.

### GetInfo

`func (o *Ticket) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Ticket) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Ticket) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Ticket) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetEscalationStartDateUTC

`func (o *Ticket) GetEscalationStartDateUTC() string`

GetEscalationStartDateUTC returns the EscalationStartDateUTC field if non-nil, zero value otherwise.

### GetEscalationStartDateUTCOk

`func (o *Ticket) GetEscalationStartDateUTCOk() (*string, bool)`

GetEscalationStartDateUTCOk returns a tuple with the EscalationStartDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalationStartDateUTC

`func (o *Ticket) SetEscalationStartDateUTC(v string)`

SetEscalationStartDateUTC sets EscalationStartDateUTC field to given value.

### HasEscalationStartDateUTC

`func (o *Ticket) HasEscalationStartDateUTC() bool`

HasEscalationStartDateUTC returns a boolean if a field has been set.

### GetEscalationLevel

`func (o *Ticket) GetEscalationLevel() int32`

GetEscalationLevel returns the EscalationLevel field if non-nil, zero value otherwise.

### GetEscalationLevelOk

`func (o *Ticket) GetEscalationLevelOk() (*int32, bool)`

GetEscalationLevelOk returns a tuple with the EscalationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalationLevel

`func (o *Ticket) SetEscalationLevel(v int32)`

SetEscalationLevel sets EscalationLevel field to given value.

### HasEscalationLevel

`func (o *Ticket) HasEscalationLevel() bool`

HasEscalationLevel returns a boolean if a field has been set.

### SetEscalationLevelNil

`func (o *Ticket) SetEscalationLevelNil(b bool)`

 SetEscalationLevelNil sets the value for EscalationLevel to be an explicit nil

### UnsetEscalationLevel
`func (o *Ticket) UnsetEscalationLevel()`

UnsetEscalationLevel ensures that no value is present for EscalationLevel, not even an explicit nil
### GetMinutesBeforeWaiting

`func (o *Ticket) GetMinutesBeforeWaiting() int32`

GetMinutesBeforeWaiting returns the MinutesBeforeWaiting field if non-nil, zero value otherwise.

### GetMinutesBeforeWaitingOk

`func (o *Ticket) GetMinutesBeforeWaitingOk() (*int32, bool)`

GetMinutesBeforeWaitingOk returns a tuple with the MinutesBeforeWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesBeforeWaiting

`func (o *Ticket) SetMinutesBeforeWaiting(v int32)`

SetMinutesBeforeWaiting sets MinutesBeforeWaiting field to given value.

### HasMinutesBeforeWaiting

`func (o *Ticket) HasMinutesBeforeWaiting() bool`

HasMinutesBeforeWaiting returns a boolean if a field has been set.

### SetMinutesBeforeWaitingNil

`func (o *Ticket) SetMinutesBeforeWaitingNil(b bool)`

 SetMinutesBeforeWaitingNil sets the value for MinutesBeforeWaiting to be an explicit nil

### UnsetMinutesBeforeWaiting
`func (o *Ticket) UnsetMinutesBeforeWaiting()`

UnsetMinutesBeforeWaiting ensures that no value is present for MinutesBeforeWaiting, not even an explicit nil
### GetRespondedSkippedMinutes

`func (o *Ticket) GetRespondedSkippedMinutes() int32`

GetRespondedSkippedMinutes returns the RespondedSkippedMinutes field if non-nil, zero value otherwise.

### GetRespondedSkippedMinutesOk

`func (o *Ticket) GetRespondedSkippedMinutesOk() (*int32, bool)`

GetRespondedSkippedMinutesOk returns a tuple with the RespondedSkippedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondedSkippedMinutes

`func (o *Ticket) SetRespondedSkippedMinutes(v int32)`

SetRespondedSkippedMinutes sets RespondedSkippedMinutes field to given value.

### HasRespondedSkippedMinutes

`func (o *Ticket) HasRespondedSkippedMinutes() bool`

HasRespondedSkippedMinutes returns a boolean if a field has been set.

### SetRespondedSkippedMinutesNil

`func (o *Ticket) SetRespondedSkippedMinutesNil(b bool)`

 SetRespondedSkippedMinutesNil sets the value for RespondedSkippedMinutes to be an explicit nil

### UnsetRespondedSkippedMinutes
`func (o *Ticket) UnsetRespondedSkippedMinutes()`

UnsetRespondedSkippedMinutes ensures that no value is present for RespondedSkippedMinutes, not even an explicit nil
### GetResplanSkippedMinutes

`func (o *Ticket) GetResplanSkippedMinutes() int32`

GetResplanSkippedMinutes returns the ResplanSkippedMinutes field if non-nil, zero value otherwise.

### GetResplanSkippedMinutesOk

`func (o *Ticket) GetResplanSkippedMinutesOk() (*int32, bool)`

GetResplanSkippedMinutesOk returns a tuple with the ResplanSkippedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResplanSkippedMinutes

`func (o *Ticket) SetResplanSkippedMinutes(v int32)`

SetResplanSkippedMinutes sets ResplanSkippedMinutes field to given value.

### HasResplanSkippedMinutes

`func (o *Ticket) HasResplanSkippedMinutes() bool`

HasResplanSkippedMinutes returns a boolean if a field has been set.

### SetResplanSkippedMinutesNil

`func (o *Ticket) SetResplanSkippedMinutesNil(b bool)`

 SetResplanSkippedMinutesNil sets the value for ResplanSkippedMinutes to be an explicit nil

### UnsetResplanSkippedMinutes
`func (o *Ticket) UnsetResplanSkippedMinutes()`

UnsetResplanSkippedMinutes ensures that no value is present for ResplanSkippedMinutes, not even an explicit nil
### GetRespondedHours

`func (o *Ticket) GetRespondedHours() float64`

GetRespondedHours returns the RespondedHours field if non-nil, zero value otherwise.

### GetRespondedHoursOk

`func (o *Ticket) GetRespondedHoursOk() (*float64, bool)`

GetRespondedHoursOk returns a tuple with the RespondedHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondedHours

`func (o *Ticket) SetRespondedHours(v float64)`

SetRespondedHours sets RespondedHours field to given value.

### HasRespondedHours

`func (o *Ticket) HasRespondedHours() bool`

HasRespondedHours returns a boolean if a field has been set.

### SetRespondedHoursNil

`func (o *Ticket) SetRespondedHoursNil(b bool)`

 SetRespondedHoursNil sets the value for RespondedHours to be an explicit nil

### UnsetRespondedHours
`func (o *Ticket) UnsetRespondedHours()`

UnsetRespondedHours ensures that no value is present for RespondedHours, not even an explicit nil
### GetRespondedBy

`func (o *Ticket) GetRespondedBy() string`

GetRespondedBy returns the RespondedBy field if non-nil, zero value otherwise.

### GetRespondedByOk

`func (o *Ticket) GetRespondedByOk() (*string, bool)`

GetRespondedByOk returns a tuple with the RespondedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondedBy

`func (o *Ticket) SetRespondedBy(v string)`

SetRespondedBy sets RespondedBy field to given value.

### HasRespondedBy

`func (o *Ticket) HasRespondedBy() bool`

HasRespondedBy returns a boolean if a field has been set.

### GetResplanHours

`func (o *Ticket) GetResplanHours() float64`

GetResplanHours returns the ResplanHours field if non-nil, zero value otherwise.

### GetResplanHoursOk

`func (o *Ticket) GetResplanHoursOk() (*float64, bool)`

GetResplanHoursOk returns a tuple with the ResplanHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResplanHours

`func (o *Ticket) SetResplanHours(v float64)`

SetResplanHours sets ResplanHours field to given value.

### HasResplanHours

`func (o *Ticket) HasResplanHours() bool`

HasResplanHours returns a boolean if a field has been set.

### SetResplanHoursNil

`func (o *Ticket) SetResplanHoursNil(b bool)`

 SetResplanHoursNil sets the value for ResplanHours to be an explicit nil

### UnsetResplanHours
`func (o *Ticket) UnsetResplanHours()`

UnsetResplanHours ensures that no value is present for ResplanHours, not even an explicit nil
### GetResplanBy

`func (o *Ticket) GetResplanBy() string`

GetResplanBy returns the ResplanBy field if non-nil, zero value otherwise.

### GetResplanByOk

`func (o *Ticket) GetResplanByOk() (*string, bool)`

GetResplanByOk returns a tuple with the ResplanBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResplanBy

`func (o *Ticket) SetResplanBy(v string)`

SetResplanBy sets ResplanBy field to given value.

### HasResplanBy

`func (o *Ticket) HasResplanBy() bool`

HasResplanBy returns a boolean if a field has been set.

### GetResolutionHours

`func (o *Ticket) GetResolutionHours() float64`

GetResolutionHours returns the ResolutionHours field if non-nil, zero value otherwise.

### GetResolutionHoursOk

`func (o *Ticket) GetResolutionHoursOk() (*float64, bool)`

GetResolutionHoursOk returns a tuple with the ResolutionHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionHours

`func (o *Ticket) SetResolutionHours(v float64)`

SetResolutionHours sets ResolutionHours field to given value.

### HasResolutionHours

`func (o *Ticket) HasResolutionHours() bool`

HasResolutionHours returns a boolean if a field has been set.

### SetResolutionHoursNil

`func (o *Ticket) SetResolutionHoursNil(b bool)`

 SetResolutionHoursNil sets the value for ResolutionHours to be an explicit nil

### UnsetResolutionHours
`func (o *Ticket) UnsetResolutionHours()`

UnsetResolutionHours ensures that no value is present for ResolutionHours, not even an explicit nil
### GetResolvedBy

`func (o *Ticket) GetResolvedBy() string`

GetResolvedBy returns the ResolvedBy field if non-nil, zero value otherwise.

### GetResolvedByOk

`func (o *Ticket) GetResolvedByOk() (*string, bool)`

GetResolvedByOk returns a tuple with the ResolvedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedBy

`func (o *Ticket) SetResolvedBy(v string)`

SetResolvedBy sets ResolvedBy field to given value.

### HasResolvedBy

`func (o *Ticket) HasResolvedBy() bool`

HasResolvedBy returns a boolean if a field has been set.

### GetMinutesWaiting

`func (o *Ticket) GetMinutesWaiting() int32`

GetMinutesWaiting returns the MinutesWaiting field if non-nil, zero value otherwise.

### GetMinutesWaitingOk

`func (o *Ticket) GetMinutesWaitingOk() (*int32, bool)`

GetMinutesWaitingOk returns a tuple with the MinutesWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesWaiting

`func (o *Ticket) SetMinutesWaiting(v int32)`

SetMinutesWaiting sets MinutesWaiting field to given value.

### HasMinutesWaiting

`func (o *Ticket) HasMinutesWaiting() bool`

HasMinutesWaiting returns a boolean if a field has been set.

### SetMinutesWaitingNil

`func (o *Ticket) SetMinutesWaitingNil(b bool)`

 SetMinutesWaitingNil sets the value for MinutesWaiting to be an explicit nil

### UnsetMinutesWaiting
`func (o *Ticket) UnsetMinutesWaiting()`

UnsetMinutesWaiting ensures that no value is present for MinutesWaiting, not even an explicit nil
### GetCustomFields

`func (o *Ticket) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Ticket) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Ticket) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Ticket) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


