# TimeEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**CompanyType** | Pointer to **string** |  | [optional] 
**ChargeToId** | Pointer to **NullableInt32** | If chargeToId is not specified, we asume you enter time against the company specified | [optional] 
**ChargeToType** | Pointer to **NullableString** | If chargeToId is not specified, we asume you enter time against the company specified | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  | [optional] 
**BusinessGroupDesc** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**OwnerLevelReference**](OwnerLevelReference.md) |  | [optional] 
**Department** | Pointer to [**BillingUnitReference**](BillingUnitReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**AgreementType** | Pointer to **string** |  | [optional] 
**Activity** | Pointer to [**ActivityReference**](ActivityReference.md) |  | [optional] 
**OpportunityRecid** | Pointer to **NullableInt32** |  | [optional] 
**ProjectActivity** | Pointer to **string** |  | [optional] 
**Territory** | Pointer to **string** |  | [optional] 
**TimeStart** | **time.Time** |  | 
**TimeEnd** | Pointer to **time.Time** |  | [optional] 
**HoursDeduct** | Pointer to **NullableFloat64** |  | [optional] 
**ActualHours** | Pointer to **NullableFloat64** |  | [optional] 
**BillableOption** | Pointer to **NullableString** |  Required On Updates; | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**InternalNotes** | Pointer to **string** |  | [optional] 
**AddToDetailDescriptionFlag** | Pointer to **NullableBool** |  | [optional] 
**AddToInternalAnalysisFlag** | Pointer to **NullableBool** |  | [optional] 
**AddToResolutionFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailResourceFlag** | Pointer to **NullableBool** | This is an action flag. To update this value use the /service/tickets endpoint automaticEmailResourceFlag field | [optional] 
**EmailContactFlag** | Pointer to **NullableBool** | This is an action flag. To update this value use the /service/tickets endpoint automaticEmailContactFlag field | [optional] 
**EmailCcFlag** | Pointer to **NullableBool** | This is an action flag. To update this value use the /service/tickets endpoint automaticEmailCcFlag field | [optional] 
**EmailCc** | Pointer to **string** | To update this value use the /service/tickets endpoint automaticEmailCc field | [optional] 
**HoursBilled** | Pointer to **NullableFloat64** |  | [optional] 
**InvoiceHours** | Pointer to **NullableFloat64** |  | [optional] 
**HourlyCost** | Pointer to **string** |  | [optional] 
**EnteredBy** | Pointer to **string** |  | [optional] 
**DateEntered** | Pointer to **time.Time** |  | [optional] 
**Invoice** | Pointer to [**InvoiceReference**](InvoiceReference.md) |  | [optional] 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**HourlyRate** | Pointer to **NullableFloat64** | This field may only be Updated, it is defaulted on Create | [optional] 
**OverageRate** | Pointer to **NullableFloat64** |  | [optional] 
**AgreementHours** | Pointer to **NullableFloat64** |  | [optional] 
**AgreementAmount** | Pointer to **NullableFloat64** |  | [optional] 
**AgreementAdjustment** | Pointer to **NullableFloat64** |  | [optional] 
**Adjustment** | Pointer to **NullableFloat64** |  | [optional] 
**InvoiceReady** | Pointer to **NullableInt32** |  | [optional] 
**TimeSheet** | Pointer to [**TimeSheetReference**](TimeSheetReference.md) |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Ticket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**Phase** | Pointer to [**ProjectPhaseReference**](ProjectPhaseReference.md) |  | [optional] 
**TicketBoard** | Pointer to **string** |  | [optional] 
**TicketStatus** | Pointer to **string** |  | [optional] 
**TicketType** | Pointer to **string** |  | [optional] 
**TicketSubType** | Pointer to **string** |  | [optional] 
**InvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**ExtendedInvoiceAmount** | Pointer to **NullableFloat64** |  | [optional] 
**LocationName** | Pointer to **string** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewTimeEntry

`func NewTimeEntry(timeStart time.Time, ) *TimeEntry`

NewTimeEntry instantiates a new TimeEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeEntryWithDefaults

`func NewTimeEntryWithDefaults() *TimeEntry`

NewTimeEntryWithDefaults instantiates a new TimeEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimeEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompany

`func (o *TimeEntry) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *TimeEntry) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *TimeEntry) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *TimeEntry) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyType

`func (o *TimeEntry) GetCompanyType() string`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *TimeEntry) GetCompanyTypeOk() (*string, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *TimeEntry) SetCompanyType(v string)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *TimeEntry) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### GetChargeToId

`func (o *TimeEntry) GetChargeToId() int32`

GetChargeToId returns the ChargeToId field if non-nil, zero value otherwise.

### GetChargeToIdOk

`func (o *TimeEntry) GetChargeToIdOk() (*int32, bool)`

GetChargeToIdOk returns a tuple with the ChargeToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeToId

`func (o *TimeEntry) SetChargeToId(v int32)`

SetChargeToId sets ChargeToId field to given value.

### HasChargeToId

`func (o *TimeEntry) HasChargeToId() bool`

HasChargeToId returns a boolean if a field has been set.

### SetChargeToIdNil

`func (o *TimeEntry) SetChargeToIdNil(b bool)`

 SetChargeToIdNil sets the value for ChargeToId to be an explicit nil

### UnsetChargeToId
`func (o *TimeEntry) UnsetChargeToId()`

UnsetChargeToId ensures that no value is present for ChargeToId, not even an explicit nil
### GetChargeToType

`func (o *TimeEntry) GetChargeToType() string`

GetChargeToType returns the ChargeToType field if non-nil, zero value otherwise.

### GetChargeToTypeOk

`func (o *TimeEntry) GetChargeToTypeOk() (*string, bool)`

GetChargeToTypeOk returns a tuple with the ChargeToType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeToType

`func (o *TimeEntry) SetChargeToType(v string)`

SetChargeToType sets ChargeToType field to given value.

### HasChargeToType

`func (o *TimeEntry) HasChargeToType() bool`

HasChargeToType returns a boolean if a field has been set.

### SetChargeToTypeNil

`func (o *TimeEntry) SetChargeToTypeNil(b bool)`

 SetChargeToTypeNil sets the value for ChargeToType to be an explicit nil

### UnsetChargeToType
`func (o *TimeEntry) UnsetChargeToType()`

UnsetChargeToType ensures that no value is present for ChargeToType, not even an explicit nil
### GetMember

`func (o *TimeEntry) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *TimeEntry) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *TimeEntry) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *TimeEntry) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetLocationId

`func (o *TimeEntry) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *TimeEntry) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *TimeEntry) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *TimeEntry) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *TimeEntry) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *TimeEntry) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetBusinessUnitId

`func (o *TimeEntry) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *TimeEntry) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *TimeEntry) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *TimeEntry) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *TimeEntry) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *TimeEntry) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetBusinessGroupDesc

`func (o *TimeEntry) GetBusinessGroupDesc() string`

GetBusinessGroupDesc returns the BusinessGroupDesc field if non-nil, zero value otherwise.

### GetBusinessGroupDescOk

`func (o *TimeEntry) GetBusinessGroupDescOk() (*string, bool)`

GetBusinessGroupDescOk returns a tuple with the BusinessGroupDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessGroupDesc

`func (o *TimeEntry) SetBusinessGroupDesc(v string)`

SetBusinessGroupDesc sets BusinessGroupDesc field to given value.

### HasBusinessGroupDesc

`func (o *TimeEntry) HasBusinessGroupDesc() bool`

HasBusinessGroupDesc returns a boolean if a field has been set.

### GetLocation

`func (o *TimeEntry) GetLocation() OwnerLevelReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TimeEntry) GetLocationOk() (*OwnerLevelReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TimeEntry) SetLocation(v OwnerLevelReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TimeEntry) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *TimeEntry) GetDepartment() BillingUnitReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *TimeEntry) GetDepartmentOk() (*BillingUnitReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *TimeEntry) SetDepartment(v BillingUnitReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *TimeEntry) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetWorkType

`func (o *TimeEntry) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *TimeEntry) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *TimeEntry) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *TimeEntry) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetWorkRole

`func (o *TimeEntry) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *TimeEntry) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *TimeEntry) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *TimeEntry) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetAgreement

`func (o *TimeEntry) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *TimeEntry) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *TimeEntry) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *TimeEntry) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetAgreementType

`func (o *TimeEntry) GetAgreementType() string`

GetAgreementType returns the AgreementType field if non-nil, zero value otherwise.

### GetAgreementTypeOk

`func (o *TimeEntry) GetAgreementTypeOk() (*string, bool)`

GetAgreementTypeOk returns a tuple with the AgreementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementType

`func (o *TimeEntry) SetAgreementType(v string)`

SetAgreementType sets AgreementType field to given value.

### HasAgreementType

`func (o *TimeEntry) HasAgreementType() bool`

HasAgreementType returns a boolean if a field has been set.

### GetActivity

`func (o *TimeEntry) GetActivity() ActivityReference`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *TimeEntry) GetActivityOk() (*ActivityReference, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *TimeEntry) SetActivity(v ActivityReference)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *TimeEntry) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetOpportunityRecid

`func (o *TimeEntry) GetOpportunityRecid() int32`

GetOpportunityRecid returns the OpportunityRecid field if non-nil, zero value otherwise.

### GetOpportunityRecidOk

`func (o *TimeEntry) GetOpportunityRecidOk() (*int32, bool)`

GetOpportunityRecidOk returns a tuple with the OpportunityRecid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityRecid

`func (o *TimeEntry) SetOpportunityRecid(v int32)`

SetOpportunityRecid sets OpportunityRecid field to given value.

### HasOpportunityRecid

`func (o *TimeEntry) HasOpportunityRecid() bool`

HasOpportunityRecid returns a boolean if a field has been set.

### SetOpportunityRecidNil

`func (o *TimeEntry) SetOpportunityRecidNil(b bool)`

 SetOpportunityRecidNil sets the value for OpportunityRecid to be an explicit nil

### UnsetOpportunityRecid
`func (o *TimeEntry) UnsetOpportunityRecid()`

UnsetOpportunityRecid ensures that no value is present for OpportunityRecid, not even an explicit nil
### GetProjectActivity

`func (o *TimeEntry) GetProjectActivity() string`

GetProjectActivity returns the ProjectActivity field if non-nil, zero value otherwise.

### GetProjectActivityOk

`func (o *TimeEntry) GetProjectActivityOk() (*string, bool)`

GetProjectActivityOk returns a tuple with the ProjectActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectActivity

`func (o *TimeEntry) SetProjectActivity(v string)`

SetProjectActivity sets ProjectActivity field to given value.

### HasProjectActivity

`func (o *TimeEntry) HasProjectActivity() bool`

HasProjectActivity returns a boolean if a field has been set.

### GetTerritory

`func (o *TimeEntry) GetTerritory() string`

GetTerritory returns the Territory field if non-nil, zero value otherwise.

### GetTerritoryOk

`func (o *TimeEntry) GetTerritoryOk() (*string, bool)`

GetTerritoryOk returns a tuple with the Territory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritory

`func (o *TimeEntry) SetTerritory(v string)`

SetTerritory sets Territory field to given value.

### HasTerritory

`func (o *TimeEntry) HasTerritory() bool`

HasTerritory returns a boolean if a field has been set.

### GetTimeStart

`func (o *TimeEntry) GetTimeStart() time.Time`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *TimeEntry) GetTimeStartOk() (*time.Time, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *TimeEntry) SetTimeStart(v time.Time)`

SetTimeStart sets TimeStart field to given value.


### GetTimeEnd

`func (o *TimeEntry) GetTimeEnd() time.Time`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *TimeEntry) GetTimeEndOk() (*time.Time, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *TimeEntry) SetTimeEnd(v time.Time)`

SetTimeEnd sets TimeEnd field to given value.

### HasTimeEnd

`func (o *TimeEntry) HasTimeEnd() bool`

HasTimeEnd returns a boolean if a field has been set.

### GetHoursDeduct

`func (o *TimeEntry) GetHoursDeduct() float64`

GetHoursDeduct returns the HoursDeduct field if non-nil, zero value otherwise.

### GetHoursDeductOk

`func (o *TimeEntry) GetHoursDeductOk() (*float64, bool)`

GetHoursDeductOk returns a tuple with the HoursDeduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursDeduct

`func (o *TimeEntry) SetHoursDeduct(v float64)`

SetHoursDeduct sets HoursDeduct field to given value.

### HasHoursDeduct

`func (o *TimeEntry) HasHoursDeduct() bool`

HasHoursDeduct returns a boolean if a field has been set.

### SetHoursDeductNil

`func (o *TimeEntry) SetHoursDeductNil(b bool)`

 SetHoursDeductNil sets the value for HoursDeduct to be an explicit nil

### UnsetHoursDeduct
`func (o *TimeEntry) UnsetHoursDeduct()`

UnsetHoursDeduct ensures that no value is present for HoursDeduct, not even an explicit nil
### GetActualHours

`func (o *TimeEntry) GetActualHours() float64`

GetActualHours returns the ActualHours field if non-nil, zero value otherwise.

### GetActualHoursOk

`func (o *TimeEntry) GetActualHoursOk() (*float64, bool)`

GetActualHoursOk returns a tuple with the ActualHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualHours

`func (o *TimeEntry) SetActualHours(v float64)`

SetActualHours sets ActualHours field to given value.

### HasActualHours

`func (o *TimeEntry) HasActualHours() bool`

HasActualHours returns a boolean if a field has been set.

### SetActualHoursNil

`func (o *TimeEntry) SetActualHoursNil(b bool)`

 SetActualHoursNil sets the value for ActualHours to be an explicit nil

### UnsetActualHours
`func (o *TimeEntry) UnsetActualHours()`

UnsetActualHours ensures that no value is present for ActualHours, not even an explicit nil
### GetBillableOption

`func (o *TimeEntry) GetBillableOption() string`

GetBillableOption returns the BillableOption field if non-nil, zero value otherwise.

### GetBillableOptionOk

`func (o *TimeEntry) GetBillableOptionOk() (*string, bool)`

GetBillableOptionOk returns a tuple with the BillableOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableOption

`func (o *TimeEntry) SetBillableOption(v string)`

SetBillableOption sets BillableOption field to given value.

### HasBillableOption

`func (o *TimeEntry) HasBillableOption() bool`

HasBillableOption returns a boolean if a field has been set.

### SetBillableOptionNil

`func (o *TimeEntry) SetBillableOptionNil(b bool)`

 SetBillableOptionNil sets the value for BillableOption to be an explicit nil

### UnsetBillableOption
`func (o *TimeEntry) UnsetBillableOption()`

UnsetBillableOption ensures that no value is present for BillableOption, not even an explicit nil
### GetNotes

`func (o *TimeEntry) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TimeEntry) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TimeEntry) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TimeEntry) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetInternalNotes

`func (o *TimeEntry) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *TimeEntry) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *TimeEntry) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *TimeEntry) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetAddToDetailDescriptionFlag

`func (o *TimeEntry) GetAddToDetailDescriptionFlag() bool`

GetAddToDetailDescriptionFlag returns the AddToDetailDescriptionFlag field if non-nil, zero value otherwise.

### GetAddToDetailDescriptionFlagOk

`func (o *TimeEntry) GetAddToDetailDescriptionFlagOk() (*bool, bool)`

GetAddToDetailDescriptionFlagOk returns a tuple with the AddToDetailDescriptionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToDetailDescriptionFlag

`func (o *TimeEntry) SetAddToDetailDescriptionFlag(v bool)`

SetAddToDetailDescriptionFlag sets AddToDetailDescriptionFlag field to given value.

### HasAddToDetailDescriptionFlag

`func (o *TimeEntry) HasAddToDetailDescriptionFlag() bool`

HasAddToDetailDescriptionFlag returns a boolean if a field has been set.

### SetAddToDetailDescriptionFlagNil

`func (o *TimeEntry) SetAddToDetailDescriptionFlagNil(b bool)`

 SetAddToDetailDescriptionFlagNil sets the value for AddToDetailDescriptionFlag to be an explicit nil

### UnsetAddToDetailDescriptionFlag
`func (o *TimeEntry) UnsetAddToDetailDescriptionFlag()`

UnsetAddToDetailDescriptionFlag ensures that no value is present for AddToDetailDescriptionFlag, not even an explicit nil
### GetAddToInternalAnalysisFlag

`func (o *TimeEntry) GetAddToInternalAnalysisFlag() bool`

GetAddToInternalAnalysisFlag returns the AddToInternalAnalysisFlag field if non-nil, zero value otherwise.

### GetAddToInternalAnalysisFlagOk

`func (o *TimeEntry) GetAddToInternalAnalysisFlagOk() (*bool, bool)`

GetAddToInternalAnalysisFlagOk returns a tuple with the AddToInternalAnalysisFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToInternalAnalysisFlag

`func (o *TimeEntry) SetAddToInternalAnalysisFlag(v bool)`

SetAddToInternalAnalysisFlag sets AddToInternalAnalysisFlag field to given value.

### HasAddToInternalAnalysisFlag

`func (o *TimeEntry) HasAddToInternalAnalysisFlag() bool`

HasAddToInternalAnalysisFlag returns a boolean if a field has been set.

### SetAddToInternalAnalysisFlagNil

`func (o *TimeEntry) SetAddToInternalAnalysisFlagNil(b bool)`

 SetAddToInternalAnalysisFlagNil sets the value for AddToInternalAnalysisFlag to be an explicit nil

### UnsetAddToInternalAnalysisFlag
`func (o *TimeEntry) UnsetAddToInternalAnalysisFlag()`

UnsetAddToInternalAnalysisFlag ensures that no value is present for AddToInternalAnalysisFlag, not even an explicit nil
### GetAddToResolutionFlag

`func (o *TimeEntry) GetAddToResolutionFlag() bool`

GetAddToResolutionFlag returns the AddToResolutionFlag field if non-nil, zero value otherwise.

### GetAddToResolutionFlagOk

`func (o *TimeEntry) GetAddToResolutionFlagOk() (*bool, bool)`

GetAddToResolutionFlagOk returns a tuple with the AddToResolutionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToResolutionFlag

`func (o *TimeEntry) SetAddToResolutionFlag(v bool)`

SetAddToResolutionFlag sets AddToResolutionFlag field to given value.

### HasAddToResolutionFlag

`func (o *TimeEntry) HasAddToResolutionFlag() bool`

HasAddToResolutionFlag returns a boolean if a field has been set.

### SetAddToResolutionFlagNil

`func (o *TimeEntry) SetAddToResolutionFlagNil(b bool)`

 SetAddToResolutionFlagNil sets the value for AddToResolutionFlag to be an explicit nil

### UnsetAddToResolutionFlag
`func (o *TimeEntry) UnsetAddToResolutionFlag()`

UnsetAddToResolutionFlag ensures that no value is present for AddToResolutionFlag, not even an explicit nil
### GetEmailResourceFlag

`func (o *TimeEntry) GetEmailResourceFlag() bool`

GetEmailResourceFlag returns the EmailResourceFlag field if non-nil, zero value otherwise.

### GetEmailResourceFlagOk

`func (o *TimeEntry) GetEmailResourceFlagOk() (*bool, bool)`

GetEmailResourceFlagOk returns a tuple with the EmailResourceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailResourceFlag

`func (o *TimeEntry) SetEmailResourceFlag(v bool)`

SetEmailResourceFlag sets EmailResourceFlag field to given value.

### HasEmailResourceFlag

`func (o *TimeEntry) HasEmailResourceFlag() bool`

HasEmailResourceFlag returns a boolean if a field has been set.

### SetEmailResourceFlagNil

`func (o *TimeEntry) SetEmailResourceFlagNil(b bool)`

 SetEmailResourceFlagNil sets the value for EmailResourceFlag to be an explicit nil

### UnsetEmailResourceFlag
`func (o *TimeEntry) UnsetEmailResourceFlag()`

UnsetEmailResourceFlag ensures that no value is present for EmailResourceFlag, not even an explicit nil
### GetEmailContactFlag

`func (o *TimeEntry) GetEmailContactFlag() bool`

GetEmailContactFlag returns the EmailContactFlag field if non-nil, zero value otherwise.

### GetEmailContactFlagOk

`func (o *TimeEntry) GetEmailContactFlagOk() (*bool, bool)`

GetEmailContactFlagOk returns a tuple with the EmailContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailContactFlag

`func (o *TimeEntry) SetEmailContactFlag(v bool)`

SetEmailContactFlag sets EmailContactFlag field to given value.

### HasEmailContactFlag

`func (o *TimeEntry) HasEmailContactFlag() bool`

HasEmailContactFlag returns a boolean if a field has been set.

### SetEmailContactFlagNil

`func (o *TimeEntry) SetEmailContactFlagNil(b bool)`

 SetEmailContactFlagNil sets the value for EmailContactFlag to be an explicit nil

### UnsetEmailContactFlag
`func (o *TimeEntry) UnsetEmailContactFlag()`

UnsetEmailContactFlag ensures that no value is present for EmailContactFlag, not even an explicit nil
### GetEmailCcFlag

`func (o *TimeEntry) GetEmailCcFlag() bool`

GetEmailCcFlag returns the EmailCcFlag field if non-nil, zero value otherwise.

### GetEmailCcFlagOk

`func (o *TimeEntry) GetEmailCcFlagOk() (*bool, bool)`

GetEmailCcFlagOk returns a tuple with the EmailCcFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCcFlag

`func (o *TimeEntry) SetEmailCcFlag(v bool)`

SetEmailCcFlag sets EmailCcFlag field to given value.

### HasEmailCcFlag

`func (o *TimeEntry) HasEmailCcFlag() bool`

HasEmailCcFlag returns a boolean if a field has been set.

### SetEmailCcFlagNil

`func (o *TimeEntry) SetEmailCcFlagNil(b bool)`

 SetEmailCcFlagNil sets the value for EmailCcFlag to be an explicit nil

### UnsetEmailCcFlag
`func (o *TimeEntry) UnsetEmailCcFlag()`

UnsetEmailCcFlag ensures that no value is present for EmailCcFlag, not even an explicit nil
### GetEmailCc

`func (o *TimeEntry) GetEmailCc() string`

GetEmailCc returns the EmailCc field if non-nil, zero value otherwise.

### GetEmailCcOk

`func (o *TimeEntry) GetEmailCcOk() (*string, bool)`

GetEmailCcOk returns a tuple with the EmailCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCc

`func (o *TimeEntry) SetEmailCc(v string)`

SetEmailCc sets EmailCc field to given value.

### HasEmailCc

`func (o *TimeEntry) HasEmailCc() bool`

HasEmailCc returns a boolean if a field has been set.

### GetHoursBilled

`func (o *TimeEntry) GetHoursBilled() float64`

GetHoursBilled returns the HoursBilled field if non-nil, zero value otherwise.

### GetHoursBilledOk

`func (o *TimeEntry) GetHoursBilledOk() (*float64, bool)`

GetHoursBilledOk returns a tuple with the HoursBilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursBilled

`func (o *TimeEntry) SetHoursBilled(v float64)`

SetHoursBilled sets HoursBilled field to given value.

### HasHoursBilled

`func (o *TimeEntry) HasHoursBilled() bool`

HasHoursBilled returns a boolean if a field has been set.

### SetHoursBilledNil

`func (o *TimeEntry) SetHoursBilledNil(b bool)`

 SetHoursBilledNil sets the value for HoursBilled to be an explicit nil

### UnsetHoursBilled
`func (o *TimeEntry) UnsetHoursBilled()`

UnsetHoursBilled ensures that no value is present for HoursBilled, not even an explicit nil
### GetInvoiceHours

`func (o *TimeEntry) GetInvoiceHours() float64`

GetInvoiceHours returns the InvoiceHours field if non-nil, zero value otherwise.

### GetInvoiceHoursOk

`func (o *TimeEntry) GetInvoiceHoursOk() (*float64, bool)`

GetInvoiceHoursOk returns a tuple with the InvoiceHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceHours

`func (o *TimeEntry) SetInvoiceHours(v float64)`

SetInvoiceHours sets InvoiceHours field to given value.

### HasInvoiceHours

`func (o *TimeEntry) HasInvoiceHours() bool`

HasInvoiceHours returns a boolean if a field has been set.

### SetInvoiceHoursNil

`func (o *TimeEntry) SetInvoiceHoursNil(b bool)`

 SetInvoiceHoursNil sets the value for InvoiceHours to be an explicit nil

### UnsetInvoiceHours
`func (o *TimeEntry) UnsetInvoiceHours()`

UnsetInvoiceHours ensures that no value is present for InvoiceHours, not even an explicit nil
### GetHourlyCost

`func (o *TimeEntry) GetHourlyCost() string`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *TimeEntry) GetHourlyCostOk() (*string, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *TimeEntry) SetHourlyCost(v string)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *TimeEntry) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetEnteredBy

`func (o *TimeEntry) GetEnteredBy() string`

GetEnteredBy returns the EnteredBy field if non-nil, zero value otherwise.

### GetEnteredByOk

`func (o *TimeEntry) GetEnteredByOk() (*string, bool)`

GetEnteredByOk returns a tuple with the EnteredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredBy

`func (o *TimeEntry) SetEnteredBy(v string)`

SetEnteredBy sets EnteredBy field to given value.

### HasEnteredBy

`func (o *TimeEntry) HasEnteredBy() bool`

HasEnteredBy returns a boolean if a field has been set.

### GetDateEntered

`func (o *TimeEntry) GetDateEntered() time.Time`

GetDateEntered returns the DateEntered field if non-nil, zero value otherwise.

### GetDateEnteredOk

`func (o *TimeEntry) GetDateEnteredOk() (*time.Time, bool)`

GetDateEnteredOk returns a tuple with the DateEntered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEntered

`func (o *TimeEntry) SetDateEntered(v time.Time)`

SetDateEntered sets DateEntered field to given value.

### HasDateEntered

`func (o *TimeEntry) HasDateEntered() bool`

HasDateEntered returns a boolean if a field has been set.

### GetInvoice

`func (o *TimeEntry) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *TimeEntry) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *TimeEntry) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *TimeEntry) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetMobileGuid

`func (o *TimeEntry) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *TimeEntry) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *TimeEntry) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *TimeEntry) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *TimeEntry) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *TimeEntry) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetHourlyRate

`func (o *TimeEntry) GetHourlyRate() float64`

GetHourlyRate returns the HourlyRate field if non-nil, zero value otherwise.

### GetHourlyRateOk

`func (o *TimeEntry) GetHourlyRateOk() (*float64, bool)`

GetHourlyRateOk returns a tuple with the HourlyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyRate

`func (o *TimeEntry) SetHourlyRate(v float64)`

SetHourlyRate sets HourlyRate field to given value.

### HasHourlyRate

`func (o *TimeEntry) HasHourlyRate() bool`

HasHourlyRate returns a boolean if a field has been set.

### SetHourlyRateNil

`func (o *TimeEntry) SetHourlyRateNil(b bool)`

 SetHourlyRateNil sets the value for HourlyRate to be an explicit nil

### UnsetHourlyRate
`func (o *TimeEntry) UnsetHourlyRate()`

UnsetHourlyRate ensures that no value is present for HourlyRate, not even an explicit nil
### GetOverageRate

`func (o *TimeEntry) GetOverageRate() float64`

GetOverageRate returns the OverageRate field if non-nil, zero value otherwise.

### GetOverageRateOk

`func (o *TimeEntry) GetOverageRateOk() (*float64, bool)`

GetOverageRateOk returns a tuple with the OverageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageRate

`func (o *TimeEntry) SetOverageRate(v float64)`

SetOverageRate sets OverageRate field to given value.

### HasOverageRate

`func (o *TimeEntry) HasOverageRate() bool`

HasOverageRate returns a boolean if a field has been set.

### SetOverageRateNil

`func (o *TimeEntry) SetOverageRateNil(b bool)`

 SetOverageRateNil sets the value for OverageRate to be an explicit nil

### UnsetOverageRate
`func (o *TimeEntry) UnsetOverageRate()`

UnsetOverageRate ensures that no value is present for OverageRate, not even an explicit nil
### GetAgreementHours

`func (o *TimeEntry) GetAgreementHours() float64`

GetAgreementHours returns the AgreementHours field if non-nil, zero value otherwise.

### GetAgreementHoursOk

`func (o *TimeEntry) GetAgreementHoursOk() (*float64, bool)`

GetAgreementHoursOk returns a tuple with the AgreementHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementHours

`func (o *TimeEntry) SetAgreementHours(v float64)`

SetAgreementHours sets AgreementHours field to given value.

### HasAgreementHours

`func (o *TimeEntry) HasAgreementHours() bool`

HasAgreementHours returns a boolean if a field has been set.

### SetAgreementHoursNil

`func (o *TimeEntry) SetAgreementHoursNil(b bool)`

 SetAgreementHoursNil sets the value for AgreementHours to be an explicit nil

### UnsetAgreementHours
`func (o *TimeEntry) UnsetAgreementHours()`

UnsetAgreementHours ensures that no value is present for AgreementHours, not even an explicit nil
### GetAgreementAmount

`func (o *TimeEntry) GetAgreementAmount() float64`

GetAgreementAmount returns the AgreementAmount field if non-nil, zero value otherwise.

### GetAgreementAmountOk

`func (o *TimeEntry) GetAgreementAmountOk() (*float64, bool)`

GetAgreementAmountOk returns a tuple with the AgreementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementAmount

`func (o *TimeEntry) SetAgreementAmount(v float64)`

SetAgreementAmount sets AgreementAmount field to given value.

### HasAgreementAmount

`func (o *TimeEntry) HasAgreementAmount() bool`

HasAgreementAmount returns a boolean if a field has been set.

### SetAgreementAmountNil

`func (o *TimeEntry) SetAgreementAmountNil(b bool)`

 SetAgreementAmountNil sets the value for AgreementAmount to be an explicit nil

### UnsetAgreementAmount
`func (o *TimeEntry) UnsetAgreementAmount()`

UnsetAgreementAmount ensures that no value is present for AgreementAmount, not even an explicit nil
### GetAgreementAdjustment

`func (o *TimeEntry) GetAgreementAdjustment() float64`

GetAgreementAdjustment returns the AgreementAdjustment field if non-nil, zero value otherwise.

### GetAgreementAdjustmentOk

`func (o *TimeEntry) GetAgreementAdjustmentOk() (*float64, bool)`

GetAgreementAdjustmentOk returns a tuple with the AgreementAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementAdjustment

`func (o *TimeEntry) SetAgreementAdjustment(v float64)`

SetAgreementAdjustment sets AgreementAdjustment field to given value.

### HasAgreementAdjustment

`func (o *TimeEntry) HasAgreementAdjustment() bool`

HasAgreementAdjustment returns a boolean if a field has been set.

### SetAgreementAdjustmentNil

`func (o *TimeEntry) SetAgreementAdjustmentNil(b bool)`

 SetAgreementAdjustmentNil sets the value for AgreementAdjustment to be an explicit nil

### UnsetAgreementAdjustment
`func (o *TimeEntry) UnsetAgreementAdjustment()`

UnsetAgreementAdjustment ensures that no value is present for AgreementAdjustment, not even an explicit nil
### GetAdjustment

`func (o *TimeEntry) GetAdjustment() float64`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *TimeEntry) GetAdjustmentOk() (*float64, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *TimeEntry) SetAdjustment(v float64)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *TimeEntry) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### SetAdjustmentNil

`func (o *TimeEntry) SetAdjustmentNil(b bool)`

 SetAdjustmentNil sets the value for Adjustment to be an explicit nil

### UnsetAdjustment
`func (o *TimeEntry) UnsetAdjustment()`

UnsetAdjustment ensures that no value is present for Adjustment, not even an explicit nil
### GetInvoiceReady

`func (o *TimeEntry) GetInvoiceReady() int32`

GetInvoiceReady returns the InvoiceReady field if non-nil, zero value otherwise.

### GetInvoiceReadyOk

`func (o *TimeEntry) GetInvoiceReadyOk() (*int32, bool)`

GetInvoiceReadyOk returns a tuple with the InvoiceReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceReady

`func (o *TimeEntry) SetInvoiceReady(v int32)`

SetInvoiceReady sets InvoiceReady field to given value.

### HasInvoiceReady

`func (o *TimeEntry) HasInvoiceReady() bool`

HasInvoiceReady returns a boolean if a field has been set.

### SetInvoiceReadyNil

`func (o *TimeEntry) SetInvoiceReadyNil(b bool)`

 SetInvoiceReadyNil sets the value for InvoiceReady to be an explicit nil

### UnsetInvoiceReady
`func (o *TimeEntry) UnsetInvoiceReady()`

UnsetInvoiceReady ensures that no value is present for InvoiceReady, not even an explicit nil
### GetTimeSheet

`func (o *TimeEntry) GetTimeSheet() TimeSheetReference`

GetTimeSheet returns the TimeSheet field if non-nil, zero value otherwise.

### GetTimeSheetOk

`func (o *TimeEntry) GetTimeSheetOk() (*TimeSheetReference, bool)`

GetTimeSheetOk returns a tuple with the TimeSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSheet

`func (o *TimeEntry) SetTimeSheet(v TimeSheetReference)`

SetTimeSheet sets TimeSheet field to given value.

### HasTimeSheet

`func (o *TimeEntry) HasTimeSheet() bool`

HasTimeSheet returns a boolean if a field has been set.

### GetStatus

`func (o *TimeEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TimeEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TimeEntry) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TimeEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TimeEntry) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TimeEntry) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTicket

`func (o *TimeEntry) GetTicket() TicketReference`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *TimeEntry) GetTicketOk() (*TicketReference, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *TimeEntry) SetTicket(v TicketReference)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *TimeEntry) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetProject

`func (o *TimeEntry) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TimeEntry) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TimeEntry) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *TimeEntry) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPhase

`func (o *TimeEntry) GetPhase() ProjectPhaseReference`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *TimeEntry) GetPhaseOk() (*ProjectPhaseReference, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *TimeEntry) SetPhase(v ProjectPhaseReference)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *TimeEntry) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetTicketBoard

`func (o *TimeEntry) GetTicketBoard() string`

GetTicketBoard returns the TicketBoard field if non-nil, zero value otherwise.

### GetTicketBoardOk

`func (o *TimeEntry) GetTicketBoardOk() (*string, bool)`

GetTicketBoardOk returns a tuple with the TicketBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketBoard

`func (o *TimeEntry) SetTicketBoard(v string)`

SetTicketBoard sets TicketBoard field to given value.

### HasTicketBoard

`func (o *TimeEntry) HasTicketBoard() bool`

HasTicketBoard returns a boolean if a field has been set.

### GetTicketStatus

`func (o *TimeEntry) GetTicketStatus() string`

GetTicketStatus returns the TicketStatus field if non-nil, zero value otherwise.

### GetTicketStatusOk

`func (o *TimeEntry) GetTicketStatusOk() (*string, bool)`

GetTicketStatusOk returns a tuple with the TicketStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketStatus

`func (o *TimeEntry) SetTicketStatus(v string)`

SetTicketStatus sets TicketStatus field to given value.

### HasTicketStatus

`func (o *TimeEntry) HasTicketStatus() bool`

HasTicketStatus returns a boolean if a field has been set.

### GetTicketType

`func (o *TimeEntry) GetTicketType() string`

GetTicketType returns the TicketType field if non-nil, zero value otherwise.

### GetTicketTypeOk

`func (o *TimeEntry) GetTicketTypeOk() (*string, bool)`

GetTicketTypeOk returns a tuple with the TicketType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketType

`func (o *TimeEntry) SetTicketType(v string)`

SetTicketType sets TicketType field to given value.

### HasTicketType

`func (o *TimeEntry) HasTicketType() bool`

HasTicketType returns a boolean if a field has been set.

### GetTicketSubType

`func (o *TimeEntry) GetTicketSubType() string`

GetTicketSubType returns the TicketSubType field if non-nil, zero value otherwise.

### GetTicketSubTypeOk

`func (o *TimeEntry) GetTicketSubTypeOk() (*string, bool)`

GetTicketSubTypeOk returns a tuple with the TicketSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketSubType

`func (o *TimeEntry) SetTicketSubType(v string)`

SetTicketSubType sets TicketSubType field to given value.

### HasTicketSubType

`func (o *TimeEntry) HasTicketSubType() bool`

HasTicketSubType returns a boolean if a field has been set.

### GetInvoiceFlag

`func (o *TimeEntry) GetInvoiceFlag() bool`

GetInvoiceFlag returns the InvoiceFlag field if non-nil, zero value otherwise.

### GetInvoiceFlagOk

`func (o *TimeEntry) GetInvoiceFlagOk() (*bool, bool)`

GetInvoiceFlagOk returns a tuple with the InvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceFlag

`func (o *TimeEntry) SetInvoiceFlag(v bool)`

SetInvoiceFlag sets InvoiceFlag field to given value.

### HasInvoiceFlag

`func (o *TimeEntry) HasInvoiceFlag() bool`

HasInvoiceFlag returns a boolean if a field has been set.

### SetInvoiceFlagNil

`func (o *TimeEntry) SetInvoiceFlagNil(b bool)`

 SetInvoiceFlagNil sets the value for InvoiceFlag to be an explicit nil

### UnsetInvoiceFlag
`func (o *TimeEntry) UnsetInvoiceFlag()`

UnsetInvoiceFlag ensures that no value is present for InvoiceFlag, not even an explicit nil
### GetExtendedInvoiceAmount

`func (o *TimeEntry) GetExtendedInvoiceAmount() float64`

GetExtendedInvoiceAmount returns the ExtendedInvoiceAmount field if non-nil, zero value otherwise.

### GetExtendedInvoiceAmountOk

`func (o *TimeEntry) GetExtendedInvoiceAmountOk() (*float64, bool)`

GetExtendedInvoiceAmountOk returns a tuple with the ExtendedInvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedInvoiceAmount

`func (o *TimeEntry) SetExtendedInvoiceAmount(v float64)`

SetExtendedInvoiceAmount sets ExtendedInvoiceAmount field to given value.

### HasExtendedInvoiceAmount

`func (o *TimeEntry) HasExtendedInvoiceAmount() bool`

HasExtendedInvoiceAmount returns a boolean if a field has been set.

### SetExtendedInvoiceAmountNil

`func (o *TimeEntry) SetExtendedInvoiceAmountNil(b bool)`

 SetExtendedInvoiceAmountNil sets the value for ExtendedInvoiceAmount to be an explicit nil

### UnsetExtendedInvoiceAmount
`func (o *TimeEntry) UnsetExtendedInvoiceAmount()`

UnsetExtendedInvoiceAmount ensures that no value is present for ExtendedInvoiceAmount, not even an explicit nil
### GetLocationName

`func (o *TimeEntry) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *TimeEntry) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *TimeEntry) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *TimeEntry) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetTaxCode

`func (o *TimeEntry) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *TimeEntry) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *TimeEntry) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *TimeEntry) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetInfo

`func (o *TimeEntry) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TimeEntry) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TimeEntry) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TimeEntry) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *TimeEntry) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TimeEntry) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TimeEntry) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TimeEntry) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


