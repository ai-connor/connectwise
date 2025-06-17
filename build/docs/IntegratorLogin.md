# IntegratorLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Username** | **string** |  Max length: 50; | 
**Password** | Pointer to **string** | The password will never be returned in response. Max length: 50; | [optional] 
**CanAccessAllRecordsFlag** | Pointer to **NullableBool** | This flag controls whether the integrator can access only the db records it created, or all system records. | [optional] 
**CanAccessAllApisFlag** | Pointer to **NullableBool** | Setting this flag to true will create an integrator that can access all of the available apis in the system.             If this field is set to true, both the member and board fields are required. | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**DateInactivated** | Pointer to **time.Time** |  | [optional] 
**InactivatedBy** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**ServiceTicketApiFlag** | Pointer to **NullableBool** |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**ServiceBoardCallbackUrl** | Pointer to **string** |  Max length: 1000; | [optional] 
**ServiceBoardLegacyCallbackFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeEntryApiFlag** | Pointer to **NullableBool** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**TimeEntryCallbackUrl** | Pointer to **string** |  Max length: 1000; | [optional] 
**TimeEntryLegacyCallbackFlag** | Pointer to **NullableBool** |  | [optional] 
**ManagedServicesApiFlag** | Pointer to **NullableBool** |  | [optional] 
**ManagedServicesAutoChildFlag** | Pointer to **NullableBool** |  | [optional] 
**ManagedServicesChildingFlag** | Pointer to **NullableBool** | True if integrator is allowed to child configurations. | [optional] 
**ContactApiFlag** | Pointer to **NullableBool** |  | [optional] 
**ContactCallbackUrl** | Pointer to **string** |  Max length: 1000; | [optional] 
**ContactLegacyCallbackFlag** | Pointer to **NullableBool** |  | [optional] 
**CompanyApiFlag** | Pointer to **NullableBool** |  | [optional] 
**CompanyCallbackUrl** | Pointer to **string** |  Max length: 1000; | [optional] 
**CompanyLegacyCallbackFlag** | Pointer to **NullableBool** |  | [optional] 
**ActivityApiFlag** | Pointer to **NullableBool** |  | [optional] 
**ActivityCallbackUrl** | Pointer to **string** |  Max length: 1000; | [optional] 
**ActivityLegacyCallbackFlag** | Pointer to **NullableBool** |  | [optional] 
**InvoiceApiFlag** | Pointer to **NullableBool** |  | [optional] 
**ProductApiFlag** | Pointer to **NullableBool** |  | [optional] 
**ProductCallbackUrl** | Pointer to **string** |  Max length: 1000; | [optional] 
**ProductLegacyCallbackFlag** | Pointer to **NullableBool** |  | [optional] 
**OpportunityApiFlag** | Pointer to **NullableBool** |  | [optional] 
**OpportunityCallbackUrl** | Pointer to **string** |  Max length: 1000; | [optional] 
**OpportunityLegacyCallbackFlag** | Pointer to **NullableBool** |  | [optional] 
**OpportunityConversionApiFlag** | Pointer to **NullableBool** | True if the member has access to the Opportunity Conversion Api. | [optional] 
**MemberApiFlag** | Pointer to **NullableBool** |  | [optional] 
**MarketingApiFlag** | Pointer to **NullableBool** |  | [optional] 
**PurchasingApiFlag** | Pointer to **NullableBool** |  | [optional] 
**PurchasingCallbackUrl** | Pointer to **string** |  Max length: 1000; | [optional] 
**PurchasingLegacyCallbackFlag** | Pointer to **NullableBool** |  | [optional] 
**ReportingApiFlag** | Pointer to **NullableBool** |  | [optional] 
**SystemApiFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectApiFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectCallbackUrl** | Pointer to **string** |  Max length: 1000; | [optional] 
**ProjectLegacyCallbackFlag** | Pointer to **NullableBool** |  | [optional] 
**ConfigurationApiFlag** | Pointer to **NullableBool** |  | [optional] 
**ConfigurationAutoChildFlag** | Pointer to **NullableBool** |  | [optional] 
**ConfigurationChildlingFlag** | Pointer to **NullableBool** | True if integrator is allowed to child configurations. | [optional] 
**ConfigurationCallbackUrl** | Pointer to **string** |  Max length: 1000; | [optional] 
**ConfigurationLegacyCallbackFlag** | Pointer to **NullableBool** |  | [optional] 
**ScheduleApiFlag** | Pointer to **NullableBool** |  | [optional] 
**ScheduleCallbackUrl** | Pointer to **string** |  Max length: 1000; | [optional] 
**ScheduleLegacyCallbackFlag** | Pointer to **NullableBool** |  | [optional] 
**AgreementApiFlag** | Pointer to **NullableBool** |  | [optional] 
**AgreementCallbackUrl** | Pointer to **string** |  Max length: 1000; | [optional] 
**AgreementCallbackLegacyFlag** | Pointer to **NullableBool** |  | [optional] 
**DocumentApiFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewIntegratorLogin

`func NewIntegratorLogin(username string, ) *IntegratorLogin`

NewIntegratorLogin instantiates a new IntegratorLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegratorLoginWithDefaults

`func NewIntegratorLoginWithDefaults() *IntegratorLogin`

NewIntegratorLoginWithDefaults instantiates a new IntegratorLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegratorLogin) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegratorLogin) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegratorLogin) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IntegratorLogin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *IntegratorLogin) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IntegratorLogin) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IntegratorLogin) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *IntegratorLogin) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IntegratorLogin) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IntegratorLogin) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IntegratorLogin) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCanAccessAllRecordsFlag

`func (o *IntegratorLogin) GetCanAccessAllRecordsFlag() bool`

GetCanAccessAllRecordsFlag returns the CanAccessAllRecordsFlag field if non-nil, zero value otherwise.

### GetCanAccessAllRecordsFlagOk

`func (o *IntegratorLogin) GetCanAccessAllRecordsFlagOk() (*bool, bool)`

GetCanAccessAllRecordsFlagOk returns a tuple with the CanAccessAllRecordsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessAllRecordsFlag

`func (o *IntegratorLogin) SetCanAccessAllRecordsFlag(v bool)`

SetCanAccessAllRecordsFlag sets CanAccessAllRecordsFlag field to given value.

### HasCanAccessAllRecordsFlag

`func (o *IntegratorLogin) HasCanAccessAllRecordsFlag() bool`

HasCanAccessAllRecordsFlag returns a boolean if a field has been set.

### SetCanAccessAllRecordsFlagNil

`func (o *IntegratorLogin) SetCanAccessAllRecordsFlagNil(b bool)`

 SetCanAccessAllRecordsFlagNil sets the value for CanAccessAllRecordsFlag to be an explicit nil

### UnsetCanAccessAllRecordsFlag
`func (o *IntegratorLogin) UnsetCanAccessAllRecordsFlag()`

UnsetCanAccessAllRecordsFlag ensures that no value is present for CanAccessAllRecordsFlag, not even an explicit nil
### GetCanAccessAllApisFlag

`func (o *IntegratorLogin) GetCanAccessAllApisFlag() bool`

GetCanAccessAllApisFlag returns the CanAccessAllApisFlag field if non-nil, zero value otherwise.

### GetCanAccessAllApisFlagOk

`func (o *IntegratorLogin) GetCanAccessAllApisFlagOk() (*bool, bool)`

GetCanAccessAllApisFlagOk returns a tuple with the CanAccessAllApisFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessAllApisFlag

`func (o *IntegratorLogin) SetCanAccessAllApisFlag(v bool)`

SetCanAccessAllApisFlag sets CanAccessAllApisFlag field to given value.

### HasCanAccessAllApisFlag

`func (o *IntegratorLogin) HasCanAccessAllApisFlag() bool`

HasCanAccessAllApisFlag returns a boolean if a field has been set.

### SetCanAccessAllApisFlagNil

`func (o *IntegratorLogin) SetCanAccessAllApisFlagNil(b bool)`

 SetCanAccessAllApisFlagNil sets the value for CanAccessAllApisFlag to be an explicit nil

### UnsetCanAccessAllApisFlag
`func (o *IntegratorLogin) UnsetCanAccessAllApisFlag()`

UnsetCanAccessAllApisFlag ensures that no value is present for CanAccessAllApisFlag, not even an explicit nil
### GetInactiveFlag

`func (o *IntegratorLogin) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *IntegratorLogin) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *IntegratorLogin) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *IntegratorLogin) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *IntegratorLogin) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *IntegratorLogin) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetDateInactivated

`func (o *IntegratorLogin) GetDateInactivated() time.Time`

GetDateInactivated returns the DateInactivated field if non-nil, zero value otherwise.

### GetDateInactivatedOk

`func (o *IntegratorLogin) GetDateInactivatedOk() (*time.Time, bool)`

GetDateInactivatedOk returns a tuple with the DateInactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateInactivated

`func (o *IntegratorLogin) SetDateInactivated(v time.Time)`

SetDateInactivated sets DateInactivated field to given value.

### HasDateInactivated

`func (o *IntegratorLogin) HasDateInactivated() bool`

HasDateInactivated returns a boolean if a field has been set.

### GetInactivatedBy

`func (o *IntegratorLogin) GetInactivatedBy() MemberReference`

GetInactivatedBy returns the InactivatedBy field if non-nil, zero value otherwise.

### GetInactivatedByOk

`func (o *IntegratorLogin) GetInactivatedByOk() (*MemberReference, bool)`

GetInactivatedByOk returns a tuple with the InactivatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivatedBy

`func (o *IntegratorLogin) SetInactivatedBy(v MemberReference)`

SetInactivatedBy sets InactivatedBy field to given value.

### HasInactivatedBy

`func (o *IntegratorLogin) HasInactivatedBy() bool`

HasInactivatedBy returns a boolean if a field has been set.

### GetServiceTicketApiFlag

`func (o *IntegratorLogin) GetServiceTicketApiFlag() bool`

GetServiceTicketApiFlag returns the ServiceTicketApiFlag field if non-nil, zero value otherwise.

### GetServiceTicketApiFlagOk

`func (o *IntegratorLogin) GetServiceTicketApiFlagOk() (*bool, bool)`

GetServiceTicketApiFlagOk returns a tuple with the ServiceTicketApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTicketApiFlag

`func (o *IntegratorLogin) SetServiceTicketApiFlag(v bool)`

SetServiceTicketApiFlag sets ServiceTicketApiFlag field to given value.

### HasServiceTicketApiFlag

`func (o *IntegratorLogin) HasServiceTicketApiFlag() bool`

HasServiceTicketApiFlag returns a boolean if a field has been set.

### SetServiceTicketApiFlagNil

`func (o *IntegratorLogin) SetServiceTicketApiFlagNil(b bool)`

 SetServiceTicketApiFlagNil sets the value for ServiceTicketApiFlag to be an explicit nil

### UnsetServiceTicketApiFlag
`func (o *IntegratorLogin) UnsetServiceTicketApiFlag()`

UnsetServiceTicketApiFlag ensures that no value is present for ServiceTicketApiFlag, not even an explicit nil
### GetBoard

`func (o *IntegratorLogin) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *IntegratorLogin) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *IntegratorLogin) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *IntegratorLogin) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetServiceBoardCallbackUrl

`func (o *IntegratorLogin) GetServiceBoardCallbackUrl() string`

GetServiceBoardCallbackUrl returns the ServiceBoardCallbackUrl field if non-nil, zero value otherwise.

### GetServiceBoardCallbackUrlOk

`func (o *IntegratorLogin) GetServiceBoardCallbackUrlOk() (*string, bool)`

GetServiceBoardCallbackUrlOk returns a tuple with the ServiceBoardCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBoardCallbackUrl

`func (o *IntegratorLogin) SetServiceBoardCallbackUrl(v string)`

SetServiceBoardCallbackUrl sets ServiceBoardCallbackUrl field to given value.

### HasServiceBoardCallbackUrl

`func (o *IntegratorLogin) HasServiceBoardCallbackUrl() bool`

HasServiceBoardCallbackUrl returns a boolean if a field has been set.

### GetServiceBoardLegacyCallbackFlag

`func (o *IntegratorLogin) GetServiceBoardLegacyCallbackFlag() bool`

GetServiceBoardLegacyCallbackFlag returns the ServiceBoardLegacyCallbackFlag field if non-nil, zero value otherwise.

### GetServiceBoardLegacyCallbackFlagOk

`func (o *IntegratorLogin) GetServiceBoardLegacyCallbackFlagOk() (*bool, bool)`

GetServiceBoardLegacyCallbackFlagOk returns a tuple with the ServiceBoardLegacyCallbackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBoardLegacyCallbackFlag

`func (o *IntegratorLogin) SetServiceBoardLegacyCallbackFlag(v bool)`

SetServiceBoardLegacyCallbackFlag sets ServiceBoardLegacyCallbackFlag field to given value.

### HasServiceBoardLegacyCallbackFlag

`func (o *IntegratorLogin) HasServiceBoardLegacyCallbackFlag() bool`

HasServiceBoardLegacyCallbackFlag returns a boolean if a field has been set.

### SetServiceBoardLegacyCallbackFlagNil

`func (o *IntegratorLogin) SetServiceBoardLegacyCallbackFlagNil(b bool)`

 SetServiceBoardLegacyCallbackFlagNil sets the value for ServiceBoardLegacyCallbackFlag to be an explicit nil

### UnsetServiceBoardLegacyCallbackFlag
`func (o *IntegratorLogin) UnsetServiceBoardLegacyCallbackFlag()`

UnsetServiceBoardLegacyCallbackFlag ensures that no value is present for ServiceBoardLegacyCallbackFlag, not even an explicit nil
### GetTimeEntryApiFlag

`func (o *IntegratorLogin) GetTimeEntryApiFlag() bool`

GetTimeEntryApiFlag returns the TimeEntryApiFlag field if non-nil, zero value otherwise.

### GetTimeEntryApiFlagOk

`func (o *IntegratorLogin) GetTimeEntryApiFlagOk() (*bool, bool)`

GetTimeEntryApiFlagOk returns a tuple with the TimeEntryApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntryApiFlag

`func (o *IntegratorLogin) SetTimeEntryApiFlag(v bool)`

SetTimeEntryApiFlag sets TimeEntryApiFlag field to given value.

### HasTimeEntryApiFlag

`func (o *IntegratorLogin) HasTimeEntryApiFlag() bool`

HasTimeEntryApiFlag returns a boolean if a field has been set.

### SetTimeEntryApiFlagNil

`func (o *IntegratorLogin) SetTimeEntryApiFlagNil(b bool)`

 SetTimeEntryApiFlagNil sets the value for TimeEntryApiFlag to be an explicit nil

### UnsetTimeEntryApiFlag
`func (o *IntegratorLogin) UnsetTimeEntryApiFlag()`

UnsetTimeEntryApiFlag ensures that no value is present for TimeEntryApiFlag, not even an explicit nil
### GetMember

`func (o *IntegratorLogin) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *IntegratorLogin) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *IntegratorLogin) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *IntegratorLogin) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetTimeEntryCallbackUrl

`func (o *IntegratorLogin) GetTimeEntryCallbackUrl() string`

GetTimeEntryCallbackUrl returns the TimeEntryCallbackUrl field if non-nil, zero value otherwise.

### GetTimeEntryCallbackUrlOk

`func (o *IntegratorLogin) GetTimeEntryCallbackUrlOk() (*string, bool)`

GetTimeEntryCallbackUrlOk returns a tuple with the TimeEntryCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntryCallbackUrl

`func (o *IntegratorLogin) SetTimeEntryCallbackUrl(v string)`

SetTimeEntryCallbackUrl sets TimeEntryCallbackUrl field to given value.

### HasTimeEntryCallbackUrl

`func (o *IntegratorLogin) HasTimeEntryCallbackUrl() bool`

HasTimeEntryCallbackUrl returns a boolean if a field has been set.

### GetTimeEntryLegacyCallbackFlag

`func (o *IntegratorLogin) GetTimeEntryLegacyCallbackFlag() bool`

GetTimeEntryLegacyCallbackFlag returns the TimeEntryLegacyCallbackFlag field if non-nil, zero value otherwise.

### GetTimeEntryLegacyCallbackFlagOk

`func (o *IntegratorLogin) GetTimeEntryLegacyCallbackFlagOk() (*bool, bool)`

GetTimeEntryLegacyCallbackFlagOk returns a tuple with the TimeEntryLegacyCallbackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntryLegacyCallbackFlag

`func (o *IntegratorLogin) SetTimeEntryLegacyCallbackFlag(v bool)`

SetTimeEntryLegacyCallbackFlag sets TimeEntryLegacyCallbackFlag field to given value.

### HasTimeEntryLegacyCallbackFlag

`func (o *IntegratorLogin) HasTimeEntryLegacyCallbackFlag() bool`

HasTimeEntryLegacyCallbackFlag returns a boolean if a field has been set.

### SetTimeEntryLegacyCallbackFlagNil

`func (o *IntegratorLogin) SetTimeEntryLegacyCallbackFlagNil(b bool)`

 SetTimeEntryLegacyCallbackFlagNil sets the value for TimeEntryLegacyCallbackFlag to be an explicit nil

### UnsetTimeEntryLegacyCallbackFlag
`func (o *IntegratorLogin) UnsetTimeEntryLegacyCallbackFlag()`

UnsetTimeEntryLegacyCallbackFlag ensures that no value is present for TimeEntryLegacyCallbackFlag, not even an explicit nil
### GetManagedServicesApiFlag

`func (o *IntegratorLogin) GetManagedServicesApiFlag() bool`

GetManagedServicesApiFlag returns the ManagedServicesApiFlag field if non-nil, zero value otherwise.

### GetManagedServicesApiFlagOk

`func (o *IntegratorLogin) GetManagedServicesApiFlagOk() (*bool, bool)`

GetManagedServicesApiFlagOk returns a tuple with the ManagedServicesApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedServicesApiFlag

`func (o *IntegratorLogin) SetManagedServicesApiFlag(v bool)`

SetManagedServicesApiFlag sets ManagedServicesApiFlag field to given value.

### HasManagedServicesApiFlag

`func (o *IntegratorLogin) HasManagedServicesApiFlag() bool`

HasManagedServicesApiFlag returns a boolean if a field has been set.

### SetManagedServicesApiFlagNil

`func (o *IntegratorLogin) SetManagedServicesApiFlagNil(b bool)`

 SetManagedServicesApiFlagNil sets the value for ManagedServicesApiFlag to be an explicit nil

### UnsetManagedServicesApiFlag
`func (o *IntegratorLogin) UnsetManagedServicesApiFlag()`

UnsetManagedServicesApiFlag ensures that no value is present for ManagedServicesApiFlag, not even an explicit nil
### GetManagedServicesAutoChildFlag

`func (o *IntegratorLogin) GetManagedServicesAutoChildFlag() bool`

GetManagedServicesAutoChildFlag returns the ManagedServicesAutoChildFlag field if non-nil, zero value otherwise.

### GetManagedServicesAutoChildFlagOk

`func (o *IntegratorLogin) GetManagedServicesAutoChildFlagOk() (*bool, bool)`

GetManagedServicesAutoChildFlagOk returns a tuple with the ManagedServicesAutoChildFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedServicesAutoChildFlag

`func (o *IntegratorLogin) SetManagedServicesAutoChildFlag(v bool)`

SetManagedServicesAutoChildFlag sets ManagedServicesAutoChildFlag field to given value.

### HasManagedServicesAutoChildFlag

`func (o *IntegratorLogin) HasManagedServicesAutoChildFlag() bool`

HasManagedServicesAutoChildFlag returns a boolean if a field has been set.

### SetManagedServicesAutoChildFlagNil

`func (o *IntegratorLogin) SetManagedServicesAutoChildFlagNil(b bool)`

 SetManagedServicesAutoChildFlagNil sets the value for ManagedServicesAutoChildFlag to be an explicit nil

### UnsetManagedServicesAutoChildFlag
`func (o *IntegratorLogin) UnsetManagedServicesAutoChildFlag()`

UnsetManagedServicesAutoChildFlag ensures that no value is present for ManagedServicesAutoChildFlag, not even an explicit nil
### GetManagedServicesChildingFlag

`func (o *IntegratorLogin) GetManagedServicesChildingFlag() bool`

GetManagedServicesChildingFlag returns the ManagedServicesChildingFlag field if non-nil, zero value otherwise.

### GetManagedServicesChildingFlagOk

`func (o *IntegratorLogin) GetManagedServicesChildingFlagOk() (*bool, bool)`

GetManagedServicesChildingFlagOk returns a tuple with the ManagedServicesChildingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedServicesChildingFlag

`func (o *IntegratorLogin) SetManagedServicesChildingFlag(v bool)`

SetManagedServicesChildingFlag sets ManagedServicesChildingFlag field to given value.

### HasManagedServicesChildingFlag

`func (o *IntegratorLogin) HasManagedServicesChildingFlag() bool`

HasManagedServicesChildingFlag returns a boolean if a field has been set.

### SetManagedServicesChildingFlagNil

`func (o *IntegratorLogin) SetManagedServicesChildingFlagNil(b bool)`

 SetManagedServicesChildingFlagNil sets the value for ManagedServicesChildingFlag to be an explicit nil

### UnsetManagedServicesChildingFlag
`func (o *IntegratorLogin) UnsetManagedServicesChildingFlag()`

UnsetManagedServicesChildingFlag ensures that no value is present for ManagedServicesChildingFlag, not even an explicit nil
### GetContactApiFlag

`func (o *IntegratorLogin) GetContactApiFlag() bool`

GetContactApiFlag returns the ContactApiFlag field if non-nil, zero value otherwise.

### GetContactApiFlagOk

`func (o *IntegratorLogin) GetContactApiFlagOk() (*bool, bool)`

GetContactApiFlagOk returns a tuple with the ContactApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactApiFlag

`func (o *IntegratorLogin) SetContactApiFlag(v bool)`

SetContactApiFlag sets ContactApiFlag field to given value.

### HasContactApiFlag

`func (o *IntegratorLogin) HasContactApiFlag() bool`

HasContactApiFlag returns a boolean if a field has been set.

### SetContactApiFlagNil

`func (o *IntegratorLogin) SetContactApiFlagNil(b bool)`

 SetContactApiFlagNil sets the value for ContactApiFlag to be an explicit nil

### UnsetContactApiFlag
`func (o *IntegratorLogin) UnsetContactApiFlag()`

UnsetContactApiFlag ensures that no value is present for ContactApiFlag, not even an explicit nil
### GetContactCallbackUrl

`func (o *IntegratorLogin) GetContactCallbackUrl() string`

GetContactCallbackUrl returns the ContactCallbackUrl field if non-nil, zero value otherwise.

### GetContactCallbackUrlOk

`func (o *IntegratorLogin) GetContactCallbackUrlOk() (*string, bool)`

GetContactCallbackUrlOk returns a tuple with the ContactCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactCallbackUrl

`func (o *IntegratorLogin) SetContactCallbackUrl(v string)`

SetContactCallbackUrl sets ContactCallbackUrl field to given value.

### HasContactCallbackUrl

`func (o *IntegratorLogin) HasContactCallbackUrl() bool`

HasContactCallbackUrl returns a boolean if a field has been set.

### GetContactLegacyCallbackFlag

`func (o *IntegratorLogin) GetContactLegacyCallbackFlag() bool`

GetContactLegacyCallbackFlag returns the ContactLegacyCallbackFlag field if non-nil, zero value otherwise.

### GetContactLegacyCallbackFlagOk

`func (o *IntegratorLogin) GetContactLegacyCallbackFlagOk() (*bool, bool)`

GetContactLegacyCallbackFlagOk returns a tuple with the ContactLegacyCallbackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactLegacyCallbackFlag

`func (o *IntegratorLogin) SetContactLegacyCallbackFlag(v bool)`

SetContactLegacyCallbackFlag sets ContactLegacyCallbackFlag field to given value.

### HasContactLegacyCallbackFlag

`func (o *IntegratorLogin) HasContactLegacyCallbackFlag() bool`

HasContactLegacyCallbackFlag returns a boolean if a field has been set.

### SetContactLegacyCallbackFlagNil

`func (o *IntegratorLogin) SetContactLegacyCallbackFlagNil(b bool)`

 SetContactLegacyCallbackFlagNil sets the value for ContactLegacyCallbackFlag to be an explicit nil

### UnsetContactLegacyCallbackFlag
`func (o *IntegratorLogin) UnsetContactLegacyCallbackFlag()`

UnsetContactLegacyCallbackFlag ensures that no value is present for ContactLegacyCallbackFlag, not even an explicit nil
### GetCompanyApiFlag

`func (o *IntegratorLogin) GetCompanyApiFlag() bool`

GetCompanyApiFlag returns the CompanyApiFlag field if non-nil, zero value otherwise.

### GetCompanyApiFlagOk

`func (o *IntegratorLogin) GetCompanyApiFlagOk() (*bool, bool)`

GetCompanyApiFlagOk returns a tuple with the CompanyApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyApiFlag

`func (o *IntegratorLogin) SetCompanyApiFlag(v bool)`

SetCompanyApiFlag sets CompanyApiFlag field to given value.

### HasCompanyApiFlag

`func (o *IntegratorLogin) HasCompanyApiFlag() bool`

HasCompanyApiFlag returns a boolean if a field has been set.

### SetCompanyApiFlagNil

`func (o *IntegratorLogin) SetCompanyApiFlagNil(b bool)`

 SetCompanyApiFlagNil sets the value for CompanyApiFlag to be an explicit nil

### UnsetCompanyApiFlag
`func (o *IntegratorLogin) UnsetCompanyApiFlag()`

UnsetCompanyApiFlag ensures that no value is present for CompanyApiFlag, not even an explicit nil
### GetCompanyCallbackUrl

`func (o *IntegratorLogin) GetCompanyCallbackUrl() string`

GetCompanyCallbackUrl returns the CompanyCallbackUrl field if non-nil, zero value otherwise.

### GetCompanyCallbackUrlOk

`func (o *IntegratorLogin) GetCompanyCallbackUrlOk() (*string, bool)`

GetCompanyCallbackUrlOk returns a tuple with the CompanyCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCallbackUrl

`func (o *IntegratorLogin) SetCompanyCallbackUrl(v string)`

SetCompanyCallbackUrl sets CompanyCallbackUrl field to given value.

### HasCompanyCallbackUrl

`func (o *IntegratorLogin) HasCompanyCallbackUrl() bool`

HasCompanyCallbackUrl returns a boolean if a field has been set.

### GetCompanyLegacyCallbackFlag

`func (o *IntegratorLogin) GetCompanyLegacyCallbackFlag() bool`

GetCompanyLegacyCallbackFlag returns the CompanyLegacyCallbackFlag field if non-nil, zero value otherwise.

### GetCompanyLegacyCallbackFlagOk

`func (o *IntegratorLogin) GetCompanyLegacyCallbackFlagOk() (*bool, bool)`

GetCompanyLegacyCallbackFlagOk returns a tuple with the CompanyLegacyCallbackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLegacyCallbackFlag

`func (o *IntegratorLogin) SetCompanyLegacyCallbackFlag(v bool)`

SetCompanyLegacyCallbackFlag sets CompanyLegacyCallbackFlag field to given value.

### HasCompanyLegacyCallbackFlag

`func (o *IntegratorLogin) HasCompanyLegacyCallbackFlag() bool`

HasCompanyLegacyCallbackFlag returns a boolean if a field has been set.

### SetCompanyLegacyCallbackFlagNil

`func (o *IntegratorLogin) SetCompanyLegacyCallbackFlagNil(b bool)`

 SetCompanyLegacyCallbackFlagNil sets the value for CompanyLegacyCallbackFlag to be an explicit nil

### UnsetCompanyLegacyCallbackFlag
`func (o *IntegratorLogin) UnsetCompanyLegacyCallbackFlag()`

UnsetCompanyLegacyCallbackFlag ensures that no value is present for CompanyLegacyCallbackFlag, not even an explicit nil
### GetActivityApiFlag

`func (o *IntegratorLogin) GetActivityApiFlag() bool`

GetActivityApiFlag returns the ActivityApiFlag field if non-nil, zero value otherwise.

### GetActivityApiFlagOk

`func (o *IntegratorLogin) GetActivityApiFlagOk() (*bool, bool)`

GetActivityApiFlagOk returns a tuple with the ActivityApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityApiFlag

`func (o *IntegratorLogin) SetActivityApiFlag(v bool)`

SetActivityApiFlag sets ActivityApiFlag field to given value.

### HasActivityApiFlag

`func (o *IntegratorLogin) HasActivityApiFlag() bool`

HasActivityApiFlag returns a boolean if a field has been set.

### SetActivityApiFlagNil

`func (o *IntegratorLogin) SetActivityApiFlagNil(b bool)`

 SetActivityApiFlagNil sets the value for ActivityApiFlag to be an explicit nil

### UnsetActivityApiFlag
`func (o *IntegratorLogin) UnsetActivityApiFlag()`

UnsetActivityApiFlag ensures that no value is present for ActivityApiFlag, not even an explicit nil
### GetActivityCallbackUrl

`func (o *IntegratorLogin) GetActivityCallbackUrl() string`

GetActivityCallbackUrl returns the ActivityCallbackUrl field if non-nil, zero value otherwise.

### GetActivityCallbackUrlOk

`func (o *IntegratorLogin) GetActivityCallbackUrlOk() (*string, bool)`

GetActivityCallbackUrlOk returns a tuple with the ActivityCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCallbackUrl

`func (o *IntegratorLogin) SetActivityCallbackUrl(v string)`

SetActivityCallbackUrl sets ActivityCallbackUrl field to given value.

### HasActivityCallbackUrl

`func (o *IntegratorLogin) HasActivityCallbackUrl() bool`

HasActivityCallbackUrl returns a boolean if a field has been set.

### GetActivityLegacyCallbackFlag

`func (o *IntegratorLogin) GetActivityLegacyCallbackFlag() bool`

GetActivityLegacyCallbackFlag returns the ActivityLegacyCallbackFlag field if non-nil, zero value otherwise.

### GetActivityLegacyCallbackFlagOk

`func (o *IntegratorLogin) GetActivityLegacyCallbackFlagOk() (*bool, bool)`

GetActivityLegacyCallbackFlagOk returns a tuple with the ActivityLegacyCallbackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityLegacyCallbackFlag

`func (o *IntegratorLogin) SetActivityLegacyCallbackFlag(v bool)`

SetActivityLegacyCallbackFlag sets ActivityLegacyCallbackFlag field to given value.

### HasActivityLegacyCallbackFlag

`func (o *IntegratorLogin) HasActivityLegacyCallbackFlag() bool`

HasActivityLegacyCallbackFlag returns a boolean if a field has been set.

### SetActivityLegacyCallbackFlagNil

`func (o *IntegratorLogin) SetActivityLegacyCallbackFlagNil(b bool)`

 SetActivityLegacyCallbackFlagNil sets the value for ActivityLegacyCallbackFlag to be an explicit nil

### UnsetActivityLegacyCallbackFlag
`func (o *IntegratorLogin) UnsetActivityLegacyCallbackFlag()`

UnsetActivityLegacyCallbackFlag ensures that no value is present for ActivityLegacyCallbackFlag, not even an explicit nil
### GetInvoiceApiFlag

`func (o *IntegratorLogin) GetInvoiceApiFlag() bool`

GetInvoiceApiFlag returns the InvoiceApiFlag field if non-nil, zero value otherwise.

### GetInvoiceApiFlagOk

`func (o *IntegratorLogin) GetInvoiceApiFlagOk() (*bool, bool)`

GetInvoiceApiFlagOk returns a tuple with the InvoiceApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceApiFlag

`func (o *IntegratorLogin) SetInvoiceApiFlag(v bool)`

SetInvoiceApiFlag sets InvoiceApiFlag field to given value.

### HasInvoiceApiFlag

`func (o *IntegratorLogin) HasInvoiceApiFlag() bool`

HasInvoiceApiFlag returns a boolean if a field has been set.

### SetInvoiceApiFlagNil

`func (o *IntegratorLogin) SetInvoiceApiFlagNil(b bool)`

 SetInvoiceApiFlagNil sets the value for InvoiceApiFlag to be an explicit nil

### UnsetInvoiceApiFlag
`func (o *IntegratorLogin) UnsetInvoiceApiFlag()`

UnsetInvoiceApiFlag ensures that no value is present for InvoiceApiFlag, not even an explicit nil
### GetProductApiFlag

`func (o *IntegratorLogin) GetProductApiFlag() bool`

GetProductApiFlag returns the ProductApiFlag field if non-nil, zero value otherwise.

### GetProductApiFlagOk

`func (o *IntegratorLogin) GetProductApiFlagOk() (*bool, bool)`

GetProductApiFlagOk returns a tuple with the ProductApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductApiFlag

`func (o *IntegratorLogin) SetProductApiFlag(v bool)`

SetProductApiFlag sets ProductApiFlag field to given value.

### HasProductApiFlag

`func (o *IntegratorLogin) HasProductApiFlag() bool`

HasProductApiFlag returns a boolean if a field has been set.

### SetProductApiFlagNil

`func (o *IntegratorLogin) SetProductApiFlagNil(b bool)`

 SetProductApiFlagNil sets the value for ProductApiFlag to be an explicit nil

### UnsetProductApiFlag
`func (o *IntegratorLogin) UnsetProductApiFlag()`

UnsetProductApiFlag ensures that no value is present for ProductApiFlag, not even an explicit nil
### GetProductCallbackUrl

`func (o *IntegratorLogin) GetProductCallbackUrl() string`

GetProductCallbackUrl returns the ProductCallbackUrl field if non-nil, zero value otherwise.

### GetProductCallbackUrlOk

`func (o *IntegratorLogin) GetProductCallbackUrlOk() (*string, bool)`

GetProductCallbackUrlOk returns a tuple with the ProductCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCallbackUrl

`func (o *IntegratorLogin) SetProductCallbackUrl(v string)`

SetProductCallbackUrl sets ProductCallbackUrl field to given value.

### HasProductCallbackUrl

`func (o *IntegratorLogin) HasProductCallbackUrl() bool`

HasProductCallbackUrl returns a boolean if a field has been set.

### GetProductLegacyCallbackFlag

`func (o *IntegratorLogin) GetProductLegacyCallbackFlag() bool`

GetProductLegacyCallbackFlag returns the ProductLegacyCallbackFlag field if non-nil, zero value otherwise.

### GetProductLegacyCallbackFlagOk

`func (o *IntegratorLogin) GetProductLegacyCallbackFlagOk() (*bool, bool)`

GetProductLegacyCallbackFlagOk returns a tuple with the ProductLegacyCallbackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductLegacyCallbackFlag

`func (o *IntegratorLogin) SetProductLegacyCallbackFlag(v bool)`

SetProductLegacyCallbackFlag sets ProductLegacyCallbackFlag field to given value.

### HasProductLegacyCallbackFlag

`func (o *IntegratorLogin) HasProductLegacyCallbackFlag() bool`

HasProductLegacyCallbackFlag returns a boolean if a field has been set.

### SetProductLegacyCallbackFlagNil

`func (o *IntegratorLogin) SetProductLegacyCallbackFlagNil(b bool)`

 SetProductLegacyCallbackFlagNil sets the value for ProductLegacyCallbackFlag to be an explicit nil

### UnsetProductLegacyCallbackFlag
`func (o *IntegratorLogin) UnsetProductLegacyCallbackFlag()`

UnsetProductLegacyCallbackFlag ensures that no value is present for ProductLegacyCallbackFlag, not even an explicit nil
### GetOpportunityApiFlag

`func (o *IntegratorLogin) GetOpportunityApiFlag() bool`

GetOpportunityApiFlag returns the OpportunityApiFlag field if non-nil, zero value otherwise.

### GetOpportunityApiFlagOk

`func (o *IntegratorLogin) GetOpportunityApiFlagOk() (*bool, bool)`

GetOpportunityApiFlagOk returns a tuple with the OpportunityApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityApiFlag

`func (o *IntegratorLogin) SetOpportunityApiFlag(v bool)`

SetOpportunityApiFlag sets OpportunityApiFlag field to given value.

### HasOpportunityApiFlag

`func (o *IntegratorLogin) HasOpportunityApiFlag() bool`

HasOpportunityApiFlag returns a boolean if a field has been set.

### SetOpportunityApiFlagNil

`func (o *IntegratorLogin) SetOpportunityApiFlagNil(b bool)`

 SetOpportunityApiFlagNil sets the value for OpportunityApiFlag to be an explicit nil

### UnsetOpportunityApiFlag
`func (o *IntegratorLogin) UnsetOpportunityApiFlag()`

UnsetOpportunityApiFlag ensures that no value is present for OpportunityApiFlag, not even an explicit nil
### GetOpportunityCallbackUrl

`func (o *IntegratorLogin) GetOpportunityCallbackUrl() string`

GetOpportunityCallbackUrl returns the OpportunityCallbackUrl field if non-nil, zero value otherwise.

### GetOpportunityCallbackUrlOk

`func (o *IntegratorLogin) GetOpportunityCallbackUrlOk() (*string, bool)`

GetOpportunityCallbackUrlOk returns a tuple with the OpportunityCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityCallbackUrl

`func (o *IntegratorLogin) SetOpportunityCallbackUrl(v string)`

SetOpportunityCallbackUrl sets OpportunityCallbackUrl field to given value.

### HasOpportunityCallbackUrl

`func (o *IntegratorLogin) HasOpportunityCallbackUrl() bool`

HasOpportunityCallbackUrl returns a boolean if a field has been set.

### GetOpportunityLegacyCallbackFlag

`func (o *IntegratorLogin) GetOpportunityLegacyCallbackFlag() bool`

GetOpportunityLegacyCallbackFlag returns the OpportunityLegacyCallbackFlag field if non-nil, zero value otherwise.

### GetOpportunityLegacyCallbackFlagOk

`func (o *IntegratorLogin) GetOpportunityLegacyCallbackFlagOk() (*bool, bool)`

GetOpportunityLegacyCallbackFlagOk returns a tuple with the OpportunityLegacyCallbackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityLegacyCallbackFlag

`func (o *IntegratorLogin) SetOpportunityLegacyCallbackFlag(v bool)`

SetOpportunityLegacyCallbackFlag sets OpportunityLegacyCallbackFlag field to given value.

### HasOpportunityLegacyCallbackFlag

`func (o *IntegratorLogin) HasOpportunityLegacyCallbackFlag() bool`

HasOpportunityLegacyCallbackFlag returns a boolean if a field has been set.

### SetOpportunityLegacyCallbackFlagNil

`func (o *IntegratorLogin) SetOpportunityLegacyCallbackFlagNil(b bool)`

 SetOpportunityLegacyCallbackFlagNil sets the value for OpportunityLegacyCallbackFlag to be an explicit nil

### UnsetOpportunityLegacyCallbackFlag
`func (o *IntegratorLogin) UnsetOpportunityLegacyCallbackFlag()`

UnsetOpportunityLegacyCallbackFlag ensures that no value is present for OpportunityLegacyCallbackFlag, not even an explicit nil
### GetOpportunityConversionApiFlag

`func (o *IntegratorLogin) GetOpportunityConversionApiFlag() bool`

GetOpportunityConversionApiFlag returns the OpportunityConversionApiFlag field if non-nil, zero value otherwise.

### GetOpportunityConversionApiFlagOk

`func (o *IntegratorLogin) GetOpportunityConversionApiFlagOk() (*bool, bool)`

GetOpportunityConversionApiFlagOk returns a tuple with the OpportunityConversionApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityConversionApiFlag

`func (o *IntegratorLogin) SetOpportunityConversionApiFlag(v bool)`

SetOpportunityConversionApiFlag sets OpportunityConversionApiFlag field to given value.

### HasOpportunityConversionApiFlag

`func (o *IntegratorLogin) HasOpportunityConversionApiFlag() bool`

HasOpportunityConversionApiFlag returns a boolean if a field has been set.

### SetOpportunityConversionApiFlagNil

`func (o *IntegratorLogin) SetOpportunityConversionApiFlagNil(b bool)`

 SetOpportunityConversionApiFlagNil sets the value for OpportunityConversionApiFlag to be an explicit nil

### UnsetOpportunityConversionApiFlag
`func (o *IntegratorLogin) UnsetOpportunityConversionApiFlag()`

UnsetOpportunityConversionApiFlag ensures that no value is present for OpportunityConversionApiFlag, not even an explicit nil
### GetMemberApiFlag

`func (o *IntegratorLogin) GetMemberApiFlag() bool`

GetMemberApiFlag returns the MemberApiFlag field if non-nil, zero value otherwise.

### GetMemberApiFlagOk

`func (o *IntegratorLogin) GetMemberApiFlagOk() (*bool, bool)`

GetMemberApiFlagOk returns a tuple with the MemberApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberApiFlag

`func (o *IntegratorLogin) SetMemberApiFlag(v bool)`

SetMemberApiFlag sets MemberApiFlag field to given value.

### HasMemberApiFlag

`func (o *IntegratorLogin) HasMemberApiFlag() bool`

HasMemberApiFlag returns a boolean if a field has been set.

### SetMemberApiFlagNil

`func (o *IntegratorLogin) SetMemberApiFlagNil(b bool)`

 SetMemberApiFlagNil sets the value for MemberApiFlag to be an explicit nil

### UnsetMemberApiFlag
`func (o *IntegratorLogin) UnsetMemberApiFlag()`

UnsetMemberApiFlag ensures that no value is present for MemberApiFlag, not even an explicit nil
### GetMarketingApiFlag

`func (o *IntegratorLogin) GetMarketingApiFlag() bool`

GetMarketingApiFlag returns the MarketingApiFlag field if non-nil, zero value otherwise.

### GetMarketingApiFlagOk

`func (o *IntegratorLogin) GetMarketingApiFlagOk() (*bool, bool)`

GetMarketingApiFlagOk returns a tuple with the MarketingApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingApiFlag

`func (o *IntegratorLogin) SetMarketingApiFlag(v bool)`

SetMarketingApiFlag sets MarketingApiFlag field to given value.

### HasMarketingApiFlag

`func (o *IntegratorLogin) HasMarketingApiFlag() bool`

HasMarketingApiFlag returns a boolean if a field has been set.

### SetMarketingApiFlagNil

`func (o *IntegratorLogin) SetMarketingApiFlagNil(b bool)`

 SetMarketingApiFlagNil sets the value for MarketingApiFlag to be an explicit nil

### UnsetMarketingApiFlag
`func (o *IntegratorLogin) UnsetMarketingApiFlag()`

UnsetMarketingApiFlag ensures that no value is present for MarketingApiFlag, not even an explicit nil
### GetPurchasingApiFlag

`func (o *IntegratorLogin) GetPurchasingApiFlag() bool`

GetPurchasingApiFlag returns the PurchasingApiFlag field if non-nil, zero value otherwise.

### GetPurchasingApiFlagOk

`func (o *IntegratorLogin) GetPurchasingApiFlagOk() (*bool, bool)`

GetPurchasingApiFlagOk returns a tuple with the PurchasingApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingApiFlag

`func (o *IntegratorLogin) SetPurchasingApiFlag(v bool)`

SetPurchasingApiFlag sets PurchasingApiFlag field to given value.

### HasPurchasingApiFlag

`func (o *IntegratorLogin) HasPurchasingApiFlag() bool`

HasPurchasingApiFlag returns a boolean if a field has been set.

### SetPurchasingApiFlagNil

`func (o *IntegratorLogin) SetPurchasingApiFlagNil(b bool)`

 SetPurchasingApiFlagNil sets the value for PurchasingApiFlag to be an explicit nil

### UnsetPurchasingApiFlag
`func (o *IntegratorLogin) UnsetPurchasingApiFlag()`

UnsetPurchasingApiFlag ensures that no value is present for PurchasingApiFlag, not even an explicit nil
### GetPurchasingCallbackUrl

`func (o *IntegratorLogin) GetPurchasingCallbackUrl() string`

GetPurchasingCallbackUrl returns the PurchasingCallbackUrl field if non-nil, zero value otherwise.

### GetPurchasingCallbackUrlOk

`func (o *IntegratorLogin) GetPurchasingCallbackUrlOk() (*string, bool)`

GetPurchasingCallbackUrlOk returns a tuple with the PurchasingCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingCallbackUrl

`func (o *IntegratorLogin) SetPurchasingCallbackUrl(v string)`

SetPurchasingCallbackUrl sets PurchasingCallbackUrl field to given value.

### HasPurchasingCallbackUrl

`func (o *IntegratorLogin) HasPurchasingCallbackUrl() bool`

HasPurchasingCallbackUrl returns a boolean if a field has been set.

### GetPurchasingLegacyCallbackFlag

`func (o *IntegratorLogin) GetPurchasingLegacyCallbackFlag() bool`

GetPurchasingLegacyCallbackFlag returns the PurchasingLegacyCallbackFlag field if non-nil, zero value otherwise.

### GetPurchasingLegacyCallbackFlagOk

`func (o *IntegratorLogin) GetPurchasingLegacyCallbackFlagOk() (*bool, bool)`

GetPurchasingLegacyCallbackFlagOk returns a tuple with the PurchasingLegacyCallbackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingLegacyCallbackFlag

`func (o *IntegratorLogin) SetPurchasingLegacyCallbackFlag(v bool)`

SetPurchasingLegacyCallbackFlag sets PurchasingLegacyCallbackFlag field to given value.

### HasPurchasingLegacyCallbackFlag

`func (o *IntegratorLogin) HasPurchasingLegacyCallbackFlag() bool`

HasPurchasingLegacyCallbackFlag returns a boolean if a field has been set.

### SetPurchasingLegacyCallbackFlagNil

`func (o *IntegratorLogin) SetPurchasingLegacyCallbackFlagNil(b bool)`

 SetPurchasingLegacyCallbackFlagNil sets the value for PurchasingLegacyCallbackFlag to be an explicit nil

### UnsetPurchasingLegacyCallbackFlag
`func (o *IntegratorLogin) UnsetPurchasingLegacyCallbackFlag()`

UnsetPurchasingLegacyCallbackFlag ensures that no value is present for PurchasingLegacyCallbackFlag, not even an explicit nil
### GetReportingApiFlag

`func (o *IntegratorLogin) GetReportingApiFlag() bool`

GetReportingApiFlag returns the ReportingApiFlag field if non-nil, zero value otherwise.

### GetReportingApiFlagOk

`func (o *IntegratorLogin) GetReportingApiFlagOk() (*bool, bool)`

GetReportingApiFlagOk returns a tuple with the ReportingApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingApiFlag

`func (o *IntegratorLogin) SetReportingApiFlag(v bool)`

SetReportingApiFlag sets ReportingApiFlag field to given value.

### HasReportingApiFlag

`func (o *IntegratorLogin) HasReportingApiFlag() bool`

HasReportingApiFlag returns a boolean if a field has been set.

### SetReportingApiFlagNil

`func (o *IntegratorLogin) SetReportingApiFlagNil(b bool)`

 SetReportingApiFlagNil sets the value for ReportingApiFlag to be an explicit nil

### UnsetReportingApiFlag
`func (o *IntegratorLogin) UnsetReportingApiFlag()`

UnsetReportingApiFlag ensures that no value is present for ReportingApiFlag, not even an explicit nil
### GetSystemApiFlag

`func (o *IntegratorLogin) GetSystemApiFlag() bool`

GetSystemApiFlag returns the SystemApiFlag field if non-nil, zero value otherwise.

### GetSystemApiFlagOk

`func (o *IntegratorLogin) GetSystemApiFlagOk() (*bool, bool)`

GetSystemApiFlagOk returns a tuple with the SystemApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemApiFlag

`func (o *IntegratorLogin) SetSystemApiFlag(v bool)`

SetSystemApiFlag sets SystemApiFlag field to given value.

### HasSystemApiFlag

`func (o *IntegratorLogin) HasSystemApiFlag() bool`

HasSystemApiFlag returns a boolean if a field has been set.

### SetSystemApiFlagNil

`func (o *IntegratorLogin) SetSystemApiFlagNil(b bool)`

 SetSystemApiFlagNil sets the value for SystemApiFlag to be an explicit nil

### UnsetSystemApiFlag
`func (o *IntegratorLogin) UnsetSystemApiFlag()`

UnsetSystemApiFlag ensures that no value is present for SystemApiFlag, not even an explicit nil
### GetProjectApiFlag

`func (o *IntegratorLogin) GetProjectApiFlag() bool`

GetProjectApiFlag returns the ProjectApiFlag field if non-nil, zero value otherwise.

### GetProjectApiFlagOk

`func (o *IntegratorLogin) GetProjectApiFlagOk() (*bool, bool)`

GetProjectApiFlagOk returns a tuple with the ProjectApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectApiFlag

`func (o *IntegratorLogin) SetProjectApiFlag(v bool)`

SetProjectApiFlag sets ProjectApiFlag field to given value.

### HasProjectApiFlag

`func (o *IntegratorLogin) HasProjectApiFlag() bool`

HasProjectApiFlag returns a boolean if a field has been set.

### SetProjectApiFlagNil

`func (o *IntegratorLogin) SetProjectApiFlagNil(b bool)`

 SetProjectApiFlagNil sets the value for ProjectApiFlag to be an explicit nil

### UnsetProjectApiFlag
`func (o *IntegratorLogin) UnsetProjectApiFlag()`

UnsetProjectApiFlag ensures that no value is present for ProjectApiFlag, not even an explicit nil
### GetProjectCallbackUrl

`func (o *IntegratorLogin) GetProjectCallbackUrl() string`

GetProjectCallbackUrl returns the ProjectCallbackUrl field if non-nil, zero value otherwise.

### GetProjectCallbackUrlOk

`func (o *IntegratorLogin) GetProjectCallbackUrlOk() (*string, bool)`

GetProjectCallbackUrlOk returns a tuple with the ProjectCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCallbackUrl

`func (o *IntegratorLogin) SetProjectCallbackUrl(v string)`

SetProjectCallbackUrl sets ProjectCallbackUrl field to given value.

### HasProjectCallbackUrl

`func (o *IntegratorLogin) HasProjectCallbackUrl() bool`

HasProjectCallbackUrl returns a boolean if a field has been set.

### GetProjectLegacyCallbackFlag

`func (o *IntegratorLogin) GetProjectLegacyCallbackFlag() bool`

GetProjectLegacyCallbackFlag returns the ProjectLegacyCallbackFlag field if non-nil, zero value otherwise.

### GetProjectLegacyCallbackFlagOk

`func (o *IntegratorLogin) GetProjectLegacyCallbackFlagOk() (*bool, bool)`

GetProjectLegacyCallbackFlagOk returns a tuple with the ProjectLegacyCallbackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLegacyCallbackFlag

`func (o *IntegratorLogin) SetProjectLegacyCallbackFlag(v bool)`

SetProjectLegacyCallbackFlag sets ProjectLegacyCallbackFlag field to given value.

### HasProjectLegacyCallbackFlag

`func (o *IntegratorLogin) HasProjectLegacyCallbackFlag() bool`

HasProjectLegacyCallbackFlag returns a boolean if a field has been set.

### SetProjectLegacyCallbackFlagNil

`func (o *IntegratorLogin) SetProjectLegacyCallbackFlagNil(b bool)`

 SetProjectLegacyCallbackFlagNil sets the value for ProjectLegacyCallbackFlag to be an explicit nil

### UnsetProjectLegacyCallbackFlag
`func (o *IntegratorLogin) UnsetProjectLegacyCallbackFlag()`

UnsetProjectLegacyCallbackFlag ensures that no value is present for ProjectLegacyCallbackFlag, not even an explicit nil
### GetConfigurationApiFlag

`func (o *IntegratorLogin) GetConfigurationApiFlag() bool`

GetConfigurationApiFlag returns the ConfigurationApiFlag field if non-nil, zero value otherwise.

### GetConfigurationApiFlagOk

`func (o *IntegratorLogin) GetConfigurationApiFlagOk() (*bool, bool)`

GetConfigurationApiFlagOk returns a tuple with the ConfigurationApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationApiFlag

`func (o *IntegratorLogin) SetConfigurationApiFlag(v bool)`

SetConfigurationApiFlag sets ConfigurationApiFlag field to given value.

### HasConfigurationApiFlag

`func (o *IntegratorLogin) HasConfigurationApiFlag() bool`

HasConfigurationApiFlag returns a boolean if a field has been set.

### SetConfigurationApiFlagNil

`func (o *IntegratorLogin) SetConfigurationApiFlagNil(b bool)`

 SetConfigurationApiFlagNil sets the value for ConfigurationApiFlag to be an explicit nil

### UnsetConfigurationApiFlag
`func (o *IntegratorLogin) UnsetConfigurationApiFlag()`

UnsetConfigurationApiFlag ensures that no value is present for ConfigurationApiFlag, not even an explicit nil
### GetConfigurationAutoChildFlag

`func (o *IntegratorLogin) GetConfigurationAutoChildFlag() bool`

GetConfigurationAutoChildFlag returns the ConfigurationAutoChildFlag field if non-nil, zero value otherwise.

### GetConfigurationAutoChildFlagOk

`func (o *IntegratorLogin) GetConfigurationAutoChildFlagOk() (*bool, bool)`

GetConfigurationAutoChildFlagOk returns a tuple with the ConfigurationAutoChildFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationAutoChildFlag

`func (o *IntegratorLogin) SetConfigurationAutoChildFlag(v bool)`

SetConfigurationAutoChildFlag sets ConfigurationAutoChildFlag field to given value.

### HasConfigurationAutoChildFlag

`func (o *IntegratorLogin) HasConfigurationAutoChildFlag() bool`

HasConfigurationAutoChildFlag returns a boolean if a field has been set.

### SetConfigurationAutoChildFlagNil

`func (o *IntegratorLogin) SetConfigurationAutoChildFlagNil(b bool)`

 SetConfigurationAutoChildFlagNil sets the value for ConfigurationAutoChildFlag to be an explicit nil

### UnsetConfigurationAutoChildFlag
`func (o *IntegratorLogin) UnsetConfigurationAutoChildFlag()`

UnsetConfigurationAutoChildFlag ensures that no value is present for ConfigurationAutoChildFlag, not even an explicit nil
### GetConfigurationChildlingFlag

`func (o *IntegratorLogin) GetConfigurationChildlingFlag() bool`

GetConfigurationChildlingFlag returns the ConfigurationChildlingFlag field if non-nil, zero value otherwise.

### GetConfigurationChildlingFlagOk

`func (o *IntegratorLogin) GetConfigurationChildlingFlagOk() (*bool, bool)`

GetConfigurationChildlingFlagOk returns a tuple with the ConfigurationChildlingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationChildlingFlag

`func (o *IntegratorLogin) SetConfigurationChildlingFlag(v bool)`

SetConfigurationChildlingFlag sets ConfigurationChildlingFlag field to given value.

### HasConfigurationChildlingFlag

`func (o *IntegratorLogin) HasConfigurationChildlingFlag() bool`

HasConfigurationChildlingFlag returns a boolean if a field has been set.

### SetConfigurationChildlingFlagNil

`func (o *IntegratorLogin) SetConfigurationChildlingFlagNil(b bool)`

 SetConfigurationChildlingFlagNil sets the value for ConfigurationChildlingFlag to be an explicit nil

### UnsetConfigurationChildlingFlag
`func (o *IntegratorLogin) UnsetConfigurationChildlingFlag()`

UnsetConfigurationChildlingFlag ensures that no value is present for ConfigurationChildlingFlag, not even an explicit nil
### GetConfigurationCallbackUrl

`func (o *IntegratorLogin) GetConfigurationCallbackUrl() string`

GetConfigurationCallbackUrl returns the ConfigurationCallbackUrl field if non-nil, zero value otherwise.

### GetConfigurationCallbackUrlOk

`func (o *IntegratorLogin) GetConfigurationCallbackUrlOk() (*string, bool)`

GetConfigurationCallbackUrlOk returns a tuple with the ConfigurationCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationCallbackUrl

`func (o *IntegratorLogin) SetConfigurationCallbackUrl(v string)`

SetConfigurationCallbackUrl sets ConfigurationCallbackUrl field to given value.

### HasConfigurationCallbackUrl

`func (o *IntegratorLogin) HasConfigurationCallbackUrl() bool`

HasConfigurationCallbackUrl returns a boolean if a field has been set.

### GetConfigurationLegacyCallbackFlag

`func (o *IntegratorLogin) GetConfigurationLegacyCallbackFlag() bool`

GetConfigurationLegacyCallbackFlag returns the ConfigurationLegacyCallbackFlag field if non-nil, zero value otherwise.

### GetConfigurationLegacyCallbackFlagOk

`func (o *IntegratorLogin) GetConfigurationLegacyCallbackFlagOk() (*bool, bool)`

GetConfigurationLegacyCallbackFlagOk returns a tuple with the ConfigurationLegacyCallbackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationLegacyCallbackFlag

`func (o *IntegratorLogin) SetConfigurationLegacyCallbackFlag(v bool)`

SetConfigurationLegacyCallbackFlag sets ConfigurationLegacyCallbackFlag field to given value.

### HasConfigurationLegacyCallbackFlag

`func (o *IntegratorLogin) HasConfigurationLegacyCallbackFlag() bool`

HasConfigurationLegacyCallbackFlag returns a boolean if a field has been set.

### SetConfigurationLegacyCallbackFlagNil

`func (o *IntegratorLogin) SetConfigurationLegacyCallbackFlagNil(b bool)`

 SetConfigurationLegacyCallbackFlagNil sets the value for ConfigurationLegacyCallbackFlag to be an explicit nil

### UnsetConfigurationLegacyCallbackFlag
`func (o *IntegratorLogin) UnsetConfigurationLegacyCallbackFlag()`

UnsetConfigurationLegacyCallbackFlag ensures that no value is present for ConfigurationLegacyCallbackFlag, not even an explicit nil
### GetScheduleApiFlag

`func (o *IntegratorLogin) GetScheduleApiFlag() bool`

GetScheduleApiFlag returns the ScheduleApiFlag field if non-nil, zero value otherwise.

### GetScheduleApiFlagOk

`func (o *IntegratorLogin) GetScheduleApiFlagOk() (*bool, bool)`

GetScheduleApiFlagOk returns a tuple with the ScheduleApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleApiFlag

`func (o *IntegratorLogin) SetScheduleApiFlag(v bool)`

SetScheduleApiFlag sets ScheduleApiFlag field to given value.

### HasScheduleApiFlag

`func (o *IntegratorLogin) HasScheduleApiFlag() bool`

HasScheduleApiFlag returns a boolean if a field has been set.

### SetScheduleApiFlagNil

`func (o *IntegratorLogin) SetScheduleApiFlagNil(b bool)`

 SetScheduleApiFlagNil sets the value for ScheduleApiFlag to be an explicit nil

### UnsetScheduleApiFlag
`func (o *IntegratorLogin) UnsetScheduleApiFlag()`

UnsetScheduleApiFlag ensures that no value is present for ScheduleApiFlag, not even an explicit nil
### GetScheduleCallbackUrl

`func (o *IntegratorLogin) GetScheduleCallbackUrl() string`

GetScheduleCallbackUrl returns the ScheduleCallbackUrl field if non-nil, zero value otherwise.

### GetScheduleCallbackUrlOk

`func (o *IntegratorLogin) GetScheduleCallbackUrlOk() (*string, bool)`

GetScheduleCallbackUrlOk returns a tuple with the ScheduleCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCallbackUrl

`func (o *IntegratorLogin) SetScheduleCallbackUrl(v string)`

SetScheduleCallbackUrl sets ScheduleCallbackUrl field to given value.

### HasScheduleCallbackUrl

`func (o *IntegratorLogin) HasScheduleCallbackUrl() bool`

HasScheduleCallbackUrl returns a boolean if a field has been set.

### GetScheduleLegacyCallbackFlag

`func (o *IntegratorLogin) GetScheduleLegacyCallbackFlag() bool`

GetScheduleLegacyCallbackFlag returns the ScheduleLegacyCallbackFlag field if non-nil, zero value otherwise.

### GetScheduleLegacyCallbackFlagOk

`func (o *IntegratorLogin) GetScheduleLegacyCallbackFlagOk() (*bool, bool)`

GetScheduleLegacyCallbackFlagOk returns a tuple with the ScheduleLegacyCallbackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleLegacyCallbackFlag

`func (o *IntegratorLogin) SetScheduleLegacyCallbackFlag(v bool)`

SetScheduleLegacyCallbackFlag sets ScheduleLegacyCallbackFlag field to given value.

### HasScheduleLegacyCallbackFlag

`func (o *IntegratorLogin) HasScheduleLegacyCallbackFlag() bool`

HasScheduleLegacyCallbackFlag returns a boolean if a field has been set.

### SetScheduleLegacyCallbackFlagNil

`func (o *IntegratorLogin) SetScheduleLegacyCallbackFlagNil(b bool)`

 SetScheduleLegacyCallbackFlagNil sets the value for ScheduleLegacyCallbackFlag to be an explicit nil

### UnsetScheduleLegacyCallbackFlag
`func (o *IntegratorLogin) UnsetScheduleLegacyCallbackFlag()`

UnsetScheduleLegacyCallbackFlag ensures that no value is present for ScheduleLegacyCallbackFlag, not even an explicit nil
### GetAgreementApiFlag

`func (o *IntegratorLogin) GetAgreementApiFlag() bool`

GetAgreementApiFlag returns the AgreementApiFlag field if non-nil, zero value otherwise.

### GetAgreementApiFlagOk

`func (o *IntegratorLogin) GetAgreementApiFlagOk() (*bool, bool)`

GetAgreementApiFlagOk returns a tuple with the AgreementApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementApiFlag

`func (o *IntegratorLogin) SetAgreementApiFlag(v bool)`

SetAgreementApiFlag sets AgreementApiFlag field to given value.

### HasAgreementApiFlag

`func (o *IntegratorLogin) HasAgreementApiFlag() bool`

HasAgreementApiFlag returns a boolean if a field has been set.

### SetAgreementApiFlagNil

`func (o *IntegratorLogin) SetAgreementApiFlagNil(b bool)`

 SetAgreementApiFlagNil sets the value for AgreementApiFlag to be an explicit nil

### UnsetAgreementApiFlag
`func (o *IntegratorLogin) UnsetAgreementApiFlag()`

UnsetAgreementApiFlag ensures that no value is present for AgreementApiFlag, not even an explicit nil
### GetAgreementCallbackUrl

`func (o *IntegratorLogin) GetAgreementCallbackUrl() string`

GetAgreementCallbackUrl returns the AgreementCallbackUrl field if non-nil, zero value otherwise.

### GetAgreementCallbackUrlOk

`func (o *IntegratorLogin) GetAgreementCallbackUrlOk() (*string, bool)`

GetAgreementCallbackUrlOk returns a tuple with the AgreementCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementCallbackUrl

`func (o *IntegratorLogin) SetAgreementCallbackUrl(v string)`

SetAgreementCallbackUrl sets AgreementCallbackUrl field to given value.

### HasAgreementCallbackUrl

`func (o *IntegratorLogin) HasAgreementCallbackUrl() bool`

HasAgreementCallbackUrl returns a boolean if a field has been set.

### GetAgreementCallbackLegacyFlag

`func (o *IntegratorLogin) GetAgreementCallbackLegacyFlag() bool`

GetAgreementCallbackLegacyFlag returns the AgreementCallbackLegacyFlag field if non-nil, zero value otherwise.

### GetAgreementCallbackLegacyFlagOk

`func (o *IntegratorLogin) GetAgreementCallbackLegacyFlagOk() (*bool, bool)`

GetAgreementCallbackLegacyFlagOk returns a tuple with the AgreementCallbackLegacyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementCallbackLegacyFlag

`func (o *IntegratorLogin) SetAgreementCallbackLegacyFlag(v bool)`

SetAgreementCallbackLegacyFlag sets AgreementCallbackLegacyFlag field to given value.

### HasAgreementCallbackLegacyFlag

`func (o *IntegratorLogin) HasAgreementCallbackLegacyFlag() bool`

HasAgreementCallbackLegacyFlag returns a boolean if a field has been set.

### SetAgreementCallbackLegacyFlagNil

`func (o *IntegratorLogin) SetAgreementCallbackLegacyFlagNil(b bool)`

 SetAgreementCallbackLegacyFlagNil sets the value for AgreementCallbackLegacyFlag to be an explicit nil

### UnsetAgreementCallbackLegacyFlag
`func (o *IntegratorLogin) UnsetAgreementCallbackLegacyFlag()`

UnsetAgreementCallbackLegacyFlag ensures that no value is present for AgreementCallbackLegacyFlag, not even an explicit nil
### GetDocumentApiFlag

`func (o *IntegratorLogin) GetDocumentApiFlag() bool`

GetDocumentApiFlag returns the DocumentApiFlag field if non-nil, zero value otherwise.

### GetDocumentApiFlagOk

`func (o *IntegratorLogin) GetDocumentApiFlagOk() (*bool, bool)`

GetDocumentApiFlagOk returns a tuple with the DocumentApiFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentApiFlag

`func (o *IntegratorLogin) SetDocumentApiFlag(v bool)`

SetDocumentApiFlag sets DocumentApiFlag field to given value.

### HasDocumentApiFlag

`func (o *IntegratorLogin) HasDocumentApiFlag() bool`

HasDocumentApiFlag returns a boolean if a field has been set.

### SetDocumentApiFlagNil

`func (o *IntegratorLogin) SetDocumentApiFlagNil(b bool)`

 SetDocumentApiFlagNil sets the value for DocumentApiFlag to be an explicit nil

### UnsetDocumentApiFlag
`func (o *IntegratorLogin) UnsetDocumentApiFlag()`

UnsetDocumentApiFlag ensures that no value is present for DocumentApiFlag, not even an explicit nil
### GetInfo

`func (o *IntegratorLogin) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *IntegratorLogin) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *IntegratorLogin) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *IntegratorLogin) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


