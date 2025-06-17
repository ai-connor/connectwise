# TicketChangeLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Ticket Change Log ID | [optional] 
**PartnerId** | Pointer to **string** | Partner ID. | [optional] 
**ProductInstanceId** | Pointer to **string** | Product Instance ID. | [optional] 
**Action** | Pointer to **string** | Action. | [optional] 
**BoardId** | Pointer to **NullableInt32** | Board ID. | [optional] 
**BoardName** | Pointer to **string** | Board Name. | [optional] 
**CompanyIdentifier** | Pointer to **NullableInt32** | Company Identifier. | [optional] 
**CompanyName** | Pointer to **string** | Company Name. | [optional] 
**ContactId** | Pointer to **NullableInt32** | Contact ID. | [optional] 
**ContactName** | Pointer to **string** | Contact Name. | [optional] 
**Impact** | Pointer to **string** | Impact. | [optional] 
**OwnerIdentifier** | Pointer to **NullableInt32** | Owner Identifier. | [optional] 
**PriorityId** | Pointer to **NullableInt32** | Priority ID. | [optional] 
**PriorityLevel** | Pointer to **string** | Priority Level. | [optional] 
**PriorityName** | Pointer to **string** | Priority Name. | [optional] 
**PrioritySort** | Pointer to **NullableInt32** | Priority Sort. | [optional] 
**ResourceList** | Pointer to **string** | Resource List. | [optional] 
**Severity** | Pointer to **string** | Severity. | [optional] 
**SlaName** | Pointer to **string** | SLA Name. | [optional] 
**SlaStatus** | Pointer to **string** | SLA Status. | [optional] 
**Status** | Pointer to **string** | Status. | [optional] 
**Summary** | Pointer to **string** | Summary. | [optional] 
**TeamName** | Pointer to **string** | Team Name. | [optional] 
**TicketNumber** | Pointer to **NullableInt32** | Ticket Number. | [optional] 
**RecordType** | Pointer to **string** | Record Type. | [optional] 
**TicketOwner** | Pointer to **string** | Ticket Owner. | [optional] 
**ClosedFlag** | Pointer to **NullableBool** | Closed Flag. | [optional] 
**CustomerUpdatedFlag** | Pointer to **NullableBool** | Customer Updated Flag. | [optional] 
**ProcessingStatus** | Pointer to **string** | Processing Status. | [optional] 
**ParentTicketId** | Pointer to **NullableInt32** | Parent Ticket ID. | [optional] 
**MergedParentTicketId** | Pointer to **NullableInt32** | Merged Parent Ticket ID. | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTicketChangeLog

`func NewTicketChangeLog() *TicketChangeLog`

NewTicketChangeLog instantiates a new TicketChangeLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketChangeLogWithDefaults

`func NewTicketChangeLogWithDefaults() *TicketChangeLog`

NewTicketChangeLogWithDefaults instantiates a new TicketChangeLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TicketChangeLog) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TicketChangeLog) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TicketChangeLog) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TicketChangeLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartnerId

`func (o *TicketChangeLog) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *TicketChangeLog) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *TicketChangeLog) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *TicketChangeLog) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetProductInstanceId

`func (o *TicketChangeLog) GetProductInstanceId() string`

GetProductInstanceId returns the ProductInstanceId field if non-nil, zero value otherwise.

### GetProductInstanceIdOk

`func (o *TicketChangeLog) GetProductInstanceIdOk() (*string, bool)`

GetProductInstanceIdOk returns a tuple with the ProductInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInstanceId

`func (o *TicketChangeLog) SetProductInstanceId(v string)`

SetProductInstanceId sets ProductInstanceId field to given value.

### HasProductInstanceId

`func (o *TicketChangeLog) HasProductInstanceId() bool`

HasProductInstanceId returns a boolean if a field has been set.

### GetAction

`func (o *TicketChangeLog) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TicketChangeLog) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TicketChangeLog) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TicketChangeLog) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBoardId

`func (o *TicketChangeLog) GetBoardId() int32`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *TicketChangeLog) GetBoardIdOk() (*int32, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *TicketChangeLog) SetBoardId(v int32)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *TicketChangeLog) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### SetBoardIdNil

`func (o *TicketChangeLog) SetBoardIdNil(b bool)`

 SetBoardIdNil sets the value for BoardId to be an explicit nil

### UnsetBoardId
`func (o *TicketChangeLog) UnsetBoardId()`

UnsetBoardId ensures that no value is present for BoardId, not even an explicit nil
### GetBoardName

`func (o *TicketChangeLog) GetBoardName() string`

GetBoardName returns the BoardName field if non-nil, zero value otherwise.

### GetBoardNameOk

`func (o *TicketChangeLog) GetBoardNameOk() (*string, bool)`

GetBoardNameOk returns a tuple with the BoardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardName

`func (o *TicketChangeLog) SetBoardName(v string)`

SetBoardName sets BoardName field to given value.

### HasBoardName

`func (o *TicketChangeLog) HasBoardName() bool`

HasBoardName returns a boolean if a field has been set.

### GetCompanyIdentifier

`func (o *TicketChangeLog) GetCompanyIdentifier() int32`

GetCompanyIdentifier returns the CompanyIdentifier field if non-nil, zero value otherwise.

### GetCompanyIdentifierOk

`func (o *TicketChangeLog) GetCompanyIdentifierOk() (*int32, bool)`

GetCompanyIdentifierOk returns a tuple with the CompanyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentifier

`func (o *TicketChangeLog) SetCompanyIdentifier(v int32)`

SetCompanyIdentifier sets CompanyIdentifier field to given value.

### HasCompanyIdentifier

`func (o *TicketChangeLog) HasCompanyIdentifier() bool`

HasCompanyIdentifier returns a boolean if a field has been set.

### SetCompanyIdentifierNil

`func (o *TicketChangeLog) SetCompanyIdentifierNil(b bool)`

 SetCompanyIdentifierNil sets the value for CompanyIdentifier to be an explicit nil

### UnsetCompanyIdentifier
`func (o *TicketChangeLog) UnsetCompanyIdentifier()`

UnsetCompanyIdentifier ensures that no value is present for CompanyIdentifier, not even an explicit nil
### GetCompanyName

`func (o *TicketChangeLog) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *TicketChangeLog) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *TicketChangeLog) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *TicketChangeLog) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetContactId

`func (o *TicketChangeLog) GetContactId() int32`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *TicketChangeLog) GetContactIdOk() (*int32, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *TicketChangeLog) SetContactId(v int32)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *TicketChangeLog) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *TicketChangeLog) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *TicketChangeLog) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetContactName

`func (o *TicketChangeLog) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *TicketChangeLog) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *TicketChangeLog) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *TicketChangeLog) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetImpact

`func (o *TicketChangeLog) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *TicketChangeLog) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *TicketChangeLog) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *TicketChangeLog) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetOwnerIdentifier

`func (o *TicketChangeLog) GetOwnerIdentifier() int32`

GetOwnerIdentifier returns the OwnerIdentifier field if non-nil, zero value otherwise.

### GetOwnerIdentifierOk

`func (o *TicketChangeLog) GetOwnerIdentifierOk() (*int32, bool)`

GetOwnerIdentifierOk returns a tuple with the OwnerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdentifier

`func (o *TicketChangeLog) SetOwnerIdentifier(v int32)`

SetOwnerIdentifier sets OwnerIdentifier field to given value.

### HasOwnerIdentifier

`func (o *TicketChangeLog) HasOwnerIdentifier() bool`

HasOwnerIdentifier returns a boolean if a field has been set.

### SetOwnerIdentifierNil

`func (o *TicketChangeLog) SetOwnerIdentifierNil(b bool)`

 SetOwnerIdentifierNil sets the value for OwnerIdentifier to be an explicit nil

### UnsetOwnerIdentifier
`func (o *TicketChangeLog) UnsetOwnerIdentifier()`

UnsetOwnerIdentifier ensures that no value is present for OwnerIdentifier, not even an explicit nil
### GetPriorityId

`func (o *TicketChangeLog) GetPriorityId() int32`

GetPriorityId returns the PriorityId field if non-nil, zero value otherwise.

### GetPriorityIdOk

`func (o *TicketChangeLog) GetPriorityIdOk() (*int32, bool)`

GetPriorityIdOk returns a tuple with the PriorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityId

`func (o *TicketChangeLog) SetPriorityId(v int32)`

SetPriorityId sets PriorityId field to given value.

### HasPriorityId

`func (o *TicketChangeLog) HasPriorityId() bool`

HasPriorityId returns a boolean if a field has been set.

### SetPriorityIdNil

`func (o *TicketChangeLog) SetPriorityIdNil(b bool)`

 SetPriorityIdNil sets the value for PriorityId to be an explicit nil

### UnsetPriorityId
`func (o *TicketChangeLog) UnsetPriorityId()`

UnsetPriorityId ensures that no value is present for PriorityId, not even an explicit nil
### GetPriorityLevel

`func (o *TicketChangeLog) GetPriorityLevel() string`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *TicketChangeLog) GetPriorityLevelOk() (*string, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *TicketChangeLog) SetPriorityLevel(v string)`

SetPriorityLevel sets PriorityLevel field to given value.

### HasPriorityLevel

`func (o *TicketChangeLog) HasPriorityLevel() bool`

HasPriorityLevel returns a boolean if a field has been set.

### GetPriorityName

`func (o *TicketChangeLog) GetPriorityName() string`

GetPriorityName returns the PriorityName field if non-nil, zero value otherwise.

### GetPriorityNameOk

`func (o *TicketChangeLog) GetPriorityNameOk() (*string, bool)`

GetPriorityNameOk returns a tuple with the PriorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityName

`func (o *TicketChangeLog) SetPriorityName(v string)`

SetPriorityName sets PriorityName field to given value.

### HasPriorityName

`func (o *TicketChangeLog) HasPriorityName() bool`

HasPriorityName returns a boolean if a field has been set.

### GetPrioritySort

`func (o *TicketChangeLog) GetPrioritySort() int32`

GetPrioritySort returns the PrioritySort field if non-nil, zero value otherwise.

### GetPrioritySortOk

`func (o *TicketChangeLog) GetPrioritySortOk() (*int32, bool)`

GetPrioritySortOk returns a tuple with the PrioritySort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritySort

`func (o *TicketChangeLog) SetPrioritySort(v int32)`

SetPrioritySort sets PrioritySort field to given value.

### HasPrioritySort

`func (o *TicketChangeLog) HasPrioritySort() bool`

HasPrioritySort returns a boolean if a field has been set.

### SetPrioritySortNil

`func (o *TicketChangeLog) SetPrioritySortNil(b bool)`

 SetPrioritySortNil sets the value for PrioritySort to be an explicit nil

### UnsetPrioritySort
`func (o *TicketChangeLog) UnsetPrioritySort()`

UnsetPrioritySort ensures that no value is present for PrioritySort, not even an explicit nil
### GetResourceList

`func (o *TicketChangeLog) GetResourceList() string`

GetResourceList returns the ResourceList field if non-nil, zero value otherwise.

### GetResourceListOk

`func (o *TicketChangeLog) GetResourceListOk() (*string, bool)`

GetResourceListOk returns a tuple with the ResourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceList

`func (o *TicketChangeLog) SetResourceList(v string)`

SetResourceList sets ResourceList field to given value.

### HasResourceList

`func (o *TicketChangeLog) HasResourceList() bool`

HasResourceList returns a boolean if a field has been set.

### GetSeverity

`func (o *TicketChangeLog) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *TicketChangeLog) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *TicketChangeLog) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *TicketChangeLog) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSlaName

`func (o *TicketChangeLog) GetSlaName() string`

GetSlaName returns the SlaName field if non-nil, zero value otherwise.

### GetSlaNameOk

`func (o *TicketChangeLog) GetSlaNameOk() (*string, bool)`

GetSlaNameOk returns a tuple with the SlaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaName

`func (o *TicketChangeLog) SetSlaName(v string)`

SetSlaName sets SlaName field to given value.

### HasSlaName

`func (o *TicketChangeLog) HasSlaName() bool`

HasSlaName returns a boolean if a field has been set.

### GetSlaStatus

`func (o *TicketChangeLog) GetSlaStatus() string`

GetSlaStatus returns the SlaStatus field if non-nil, zero value otherwise.

### GetSlaStatusOk

`func (o *TicketChangeLog) GetSlaStatusOk() (*string, bool)`

GetSlaStatusOk returns a tuple with the SlaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaStatus

`func (o *TicketChangeLog) SetSlaStatus(v string)`

SetSlaStatus sets SlaStatus field to given value.

### HasSlaStatus

`func (o *TicketChangeLog) HasSlaStatus() bool`

HasSlaStatus returns a boolean if a field has been set.

### GetStatus

`func (o *TicketChangeLog) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TicketChangeLog) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TicketChangeLog) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TicketChangeLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *TicketChangeLog) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TicketChangeLog) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TicketChangeLog) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *TicketChangeLog) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTeamName

`func (o *TicketChangeLog) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *TicketChangeLog) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *TicketChangeLog) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *TicketChangeLog) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetTicketNumber

`func (o *TicketChangeLog) GetTicketNumber() int32`

GetTicketNumber returns the TicketNumber field if non-nil, zero value otherwise.

### GetTicketNumberOk

`func (o *TicketChangeLog) GetTicketNumberOk() (*int32, bool)`

GetTicketNumberOk returns a tuple with the TicketNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketNumber

`func (o *TicketChangeLog) SetTicketNumber(v int32)`

SetTicketNumber sets TicketNumber field to given value.

### HasTicketNumber

`func (o *TicketChangeLog) HasTicketNumber() bool`

HasTicketNumber returns a boolean if a field has been set.

### SetTicketNumberNil

`func (o *TicketChangeLog) SetTicketNumberNil(b bool)`

 SetTicketNumberNil sets the value for TicketNumber to be an explicit nil

### UnsetTicketNumber
`func (o *TicketChangeLog) UnsetTicketNumber()`

UnsetTicketNumber ensures that no value is present for TicketNumber, not even an explicit nil
### GetRecordType

`func (o *TicketChangeLog) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *TicketChangeLog) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *TicketChangeLog) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *TicketChangeLog) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetTicketOwner

`func (o *TicketChangeLog) GetTicketOwner() string`

GetTicketOwner returns the TicketOwner field if non-nil, zero value otherwise.

### GetTicketOwnerOk

`func (o *TicketChangeLog) GetTicketOwnerOk() (*string, bool)`

GetTicketOwnerOk returns a tuple with the TicketOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketOwner

`func (o *TicketChangeLog) SetTicketOwner(v string)`

SetTicketOwner sets TicketOwner field to given value.

### HasTicketOwner

`func (o *TicketChangeLog) HasTicketOwner() bool`

HasTicketOwner returns a boolean if a field has been set.

### GetClosedFlag

`func (o *TicketChangeLog) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *TicketChangeLog) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *TicketChangeLog) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *TicketChangeLog) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *TicketChangeLog) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *TicketChangeLog) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetCustomerUpdatedFlag

`func (o *TicketChangeLog) GetCustomerUpdatedFlag() bool`

GetCustomerUpdatedFlag returns the CustomerUpdatedFlag field if non-nil, zero value otherwise.

### GetCustomerUpdatedFlagOk

`func (o *TicketChangeLog) GetCustomerUpdatedFlagOk() (*bool, bool)`

GetCustomerUpdatedFlagOk returns a tuple with the CustomerUpdatedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUpdatedFlag

`func (o *TicketChangeLog) SetCustomerUpdatedFlag(v bool)`

SetCustomerUpdatedFlag sets CustomerUpdatedFlag field to given value.

### HasCustomerUpdatedFlag

`func (o *TicketChangeLog) HasCustomerUpdatedFlag() bool`

HasCustomerUpdatedFlag returns a boolean if a field has been set.

### SetCustomerUpdatedFlagNil

`func (o *TicketChangeLog) SetCustomerUpdatedFlagNil(b bool)`

 SetCustomerUpdatedFlagNil sets the value for CustomerUpdatedFlag to be an explicit nil

### UnsetCustomerUpdatedFlag
`func (o *TicketChangeLog) UnsetCustomerUpdatedFlag()`

UnsetCustomerUpdatedFlag ensures that no value is present for CustomerUpdatedFlag, not even an explicit nil
### GetProcessingStatus

`func (o *TicketChangeLog) GetProcessingStatus() string`

GetProcessingStatus returns the ProcessingStatus field if non-nil, zero value otherwise.

### GetProcessingStatusOk

`func (o *TicketChangeLog) GetProcessingStatusOk() (*string, bool)`

GetProcessingStatusOk returns a tuple with the ProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStatus

`func (o *TicketChangeLog) SetProcessingStatus(v string)`

SetProcessingStatus sets ProcessingStatus field to given value.

### HasProcessingStatus

`func (o *TicketChangeLog) HasProcessingStatus() bool`

HasProcessingStatus returns a boolean if a field has been set.

### GetParentTicketId

`func (o *TicketChangeLog) GetParentTicketId() int32`

GetParentTicketId returns the ParentTicketId field if non-nil, zero value otherwise.

### GetParentTicketIdOk

`func (o *TicketChangeLog) GetParentTicketIdOk() (*int32, bool)`

GetParentTicketIdOk returns a tuple with the ParentTicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTicketId

`func (o *TicketChangeLog) SetParentTicketId(v int32)`

SetParentTicketId sets ParentTicketId field to given value.

### HasParentTicketId

`func (o *TicketChangeLog) HasParentTicketId() bool`

HasParentTicketId returns a boolean if a field has been set.

### SetParentTicketIdNil

`func (o *TicketChangeLog) SetParentTicketIdNil(b bool)`

 SetParentTicketIdNil sets the value for ParentTicketId to be an explicit nil

### UnsetParentTicketId
`func (o *TicketChangeLog) UnsetParentTicketId()`

UnsetParentTicketId ensures that no value is present for ParentTicketId, not even an explicit nil
### GetMergedParentTicketId

`func (o *TicketChangeLog) GetMergedParentTicketId() int32`

GetMergedParentTicketId returns the MergedParentTicketId field if non-nil, zero value otherwise.

### GetMergedParentTicketIdOk

`func (o *TicketChangeLog) GetMergedParentTicketIdOk() (*int32, bool)`

GetMergedParentTicketIdOk returns a tuple with the MergedParentTicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedParentTicketId

`func (o *TicketChangeLog) SetMergedParentTicketId(v int32)`

SetMergedParentTicketId sets MergedParentTicketId field to given value.

### HasMergedParentTicketId

`func (o *TicketChangeLog) HasMergedParentTicketId() bool`

HasMergedParentTicketId returns a boolean if a field has been set.

### SetMergedParentTicketIdNil

`func (o *TicketChangeLog) SetMergedParentTicketIdNil(b bool)`

 SetMergedParentTicketIdNil sets the value for MergedParentTicketId to be an explicit nil

### UnsetMergedParentTicketId
`func (o *TicketChangeLog) UnsetMergedParentTicketId()`

UnsetMergedParentTicketId ensures that no value is present for MergedParentTicketId, not even an explicit nil
### GetInfo

`func (o *TicketChangeLog) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TicketChangeLog) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TicketChangeLog) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TicketChangeLog) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


