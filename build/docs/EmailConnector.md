# EmailConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**EmailServerType** | Pointer to **NullableString** |  | [optional] 
**ImapSetup** | Pointer to [**ImapSetupReference**](ImapSetupReference.md) |  | [optional] 
**Office365EmailSetup** | Pointer to [**Office365EmailSetupReference**](Office365EmailSetupReference.md) |  | [optional] 
**Asio365EmailSetup** | Pointer to [**Office365EmailSetupReference**](Office365EmailSetupReference.md) |  | [optional] 
**GoogleEmailSetup** | Pointer to [**GoogleEmailSetupReference**](GoogleEmailSetupReference.md) |  | [optional] 
**ServiceBoard** | [**BoardReference**](BoardReference.md) |  | 
**DefaultCompany** | [**CompanyReference**](CompanyReference.md) |  | 
**DefaultMember** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**EmailNotifyFrom** | Pointer to **string** |  Max length: 50; | [optional] 
**BccEmailTo** | Pointer to **string** |  Max length: 250; | [optional] 
**EmailErrorsTo** | **string** |  Max length: 50; | 
**SetEmailToDefaultContactFlag** | Pointer to **NullableBool** |  | [optional] 
**NoResponseFlag** | Pointer to **NullableBool** |  | [optional] 
**NeverRespondFlag** | Pointer to **NullableBool** |  | [optional] 
**DiscardDuplicatesFlag** | Pointer to **NullableBool** |  | [optional] 
**PostRepliesToTicketFlag** | Pointer to **NullableBool** |  | [optional] 
**CreateContactFlag** | Pointer to **NullableBool** |  | [optional] 
**ResponseEmailForNew** | Pointer to **string** |  | [optional] 
**ResponseEmailForExisting** | Pointer to **string** |  | [optional] 
**SourceOverride** | Pointer to [**ServiceSourceReference**](ServiceSourceReference.md) |  | [optional] 
**PriorityOverride** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**TypeOverride** | Pointer to [**ServiceTypeReference**](ServiceTypeReference.md) |  | [optional] 
**SubTypeOverride** | Pointer to [**ServiceSubTypeReference**](ServiceSubTypeReference.md) |  | [optional] 
**ItemOverride** | Pointer to [**ServiceItemReference**](ServiceItemReference.md) |  | [optional] 
**StatusOverride** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**AddCcFlag** | Pointer to **NullableBool** |  | [optional] 
**InboundTicketMailboxId** | Pointer to **string** |  | [optional] 
**UseEmailMessageIdFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEmailConnector

`func NewEmailConnector(serviceBoard BoardReference, defaultCompany CompanyReference, emailErrorsTo string, ) *EmailConnector`

NewEmailConnector instantiates a new EmailConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConnectorWithDefaults

`func NewEmailConnectorWithDefaults() *EmailConnector`

NewEmailConnectorWithDefaults instantiates a new EmailConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailConnector) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailConnector) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailConnector) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EmailConnector) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmailServerType

`func (o *EmailConnector) GetEmailServerType() string`

GetEmailServerType returns the EmailServerType field if non-nil, zero value otherwise.

### GetEmailServerTypeOk

`func (o *EmailConnector) GetEmailServerTypeOk() (*string, bool)`

GetEmailServerTypeOk returns a tuple with the EmailServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailServerType

`func (o *EmailConnector) SetEmailServerType(v string)`

SetEmailServerType sets EmailServerType field to given value.

### HasEmailServerType

`func (o *EmailConnector) HasEmailServerType() bool`

HasEmailServerType returns a boolean if a field has been set.

### SetEmailServerTypeNil

`func (o *EmailConnector) SetEmailServerTypeNil(b bool)`

 SetEmailServerTypeNil sets the value for EmailServerType to be an explicit nil

### UnsetEmailServerType
`func (o *EmailConnector) UnsetEmailServerType()`

UnsetEmailServerType ensures that no value is present for EmailServerType, not even an explicit nil
### GetImapSetup

`func (o *EmailConnector) GetImapSetup() ImapSetupReference`

GetImapSetup returns the ImapSetup field if non-nil, zero value otherwise.

### GetImapSetupOk

`func (o *EmailConnector) GetImapSetupOk() (*ImapSetupReference, bool)`

GetImapSetupOk returns a tuple with the ImapSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapSetup

`func (o *EmailConnector) SetImapSetup(v ImapSetupReference)`

SetImapSetup sets ImapSetup field to given value.

### HasImapSetup

`func (o *EmailConnector) HasImapSetup() bool`

HasImapSetup returns a boolean if a field has been set.

### GetOffice365EmailSetup

`func (o *EmailConnector) GetOffice365EmailSetup() Office365EmailSetupReference`

GetOffice365EmailSetup returns the Office365EmailSetup field if non-nil, zero value otherwise.

### GetOffice365EmailSetupOk

`func (o *EmailConnector) GetOffice365EmailSetupOk() (*Office365EmailSetupReference, bool)`

GetOffice365EmailSetupOk returns a tuple with the Office365EmailSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365EmailSetup

`func (o *EmailConnector) SetOffice365EmailSetup(v Office365EmailSetupReference)`

SetOffice365EmailSetup sets Office365EmailSetup field to given value.

### HasOffice365EmailSetup

`func (o *EmailConnector) HasOffice365EmailSetup() bool`

HasOffice365EmailSetup returns a boolean if a field has been set.

### GetAsio365EmailSetup

`func (o *EmailConnector) GetAsio365EmailSetup() Office365EmailSetupReference`

GetAsio365EmailSetup returns the Asio365EmailSetup field if non-nil, zero value otherwise.

### GetAsio365EmailSetupOk

`func (o *EmailConnector) GetAsio365EmailSetupOk() (*Office365EmailSetupReference, bool)`

GetAsio365EmailSetupOk returns a tuple with the Asio365EmailSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsio365EmailSetup

`func (o *EmailConnector) SetAsio365EmailSetup(v Office365EmailSetupReference)`

SetAsio365EmailSetup sets Asio365EmailSetup field to given value.

### HasAsio365EmailSetup

`func (o *EmailConnector) HasAsio365EmailSetup() bool`

HasAsio365EmailSetup returns a boolean if a field has been set.

### GetGoogleEmailSetup

`func (o *EmailConnector) GetGoogleEmailSetup() GoogleEmailSetupReference`

GetGoogleEmailSetup returns the GoogleEmailSetup field if non-nil, zero value otherwise.

### GetGoogleEmailSetupOk

`func (o *EmailConnector) GetGoogleEmailSetupOk() (*GoogleEmailSetupReference, bool)`

GetGoogleEmailSetupOk returns a tuple with the GoogleEmailSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleEmailSetup

`func (o *EmailConnector) SetGoogleEmailSetup(v GoogleEmailSetupReference)`

SetGoogleEmailSetup sets GoogleEmailSetup field to given value.

### HasGoogleEmailSetup

`func (o *EmailConnector) HasGoogleEmailSetup() bool`

HasGoogleEmailSetup returns a boolean if a field has been set.

### GetServiceBoard

`func (o *EmailConnector) GetServiceBoard() BoardReference`

GetServiceBoard returns the ServiceBoard field if non-nil, zero value otherwise.

### GetServiceBoardOk

`func (o *EmailConnector) GetServiceBoardOk() (*BoardReference, bool)`

GetServiceBoardOk returns a tuple with the ServiceBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBoard

`func (o *EmailConnector) SetServiceBoard(v BoardReference)`

SetServiceBoard sets ServiceBoard field to given value.


### GetDefaultCompany

`func (o *EmailConnector) GetDefaultCompany() CompanyReference`

GetDefaultCompany returns the DefaultCompany field if non-nil, zero value otherwise.

### GetDefaultCompanyOk

`func (o *EmailConnector) GetDefaultCompanyOk() (*CompanyReference, bool)`

GetDefaultCompanyOk returns a tuple with the DefaultCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCompany

`func (o *EmailConnector) SetDefaultCompany(v CompanyReference)`

SetDefaultCompany sets DefaultCompany field to given value.


### GetDefaultMember

`func (o *EmailConnector) GetDefaultMember() MemberReference`

GetDefaultMember returns the DefaultMember field if non-nil, zero value otherwise.

### GetDefaultMemberOk

`func (o *EmailConnector) GetDefaultMemberOk() (*MemberReference, bool)`

GetDefaultMemberOk returns a tuple with the DefaultMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMember

`func (o *EmailConnector) SetDefaultMember(v MemberReference)`

SetDefaultMember sets DefaultMember field to given value.

### HasDefaultMember

`func (o *EmailConnector) HasDefaultMember() bool`

HasDefaultMember returns a boolean if a field has been set.

### GetLocation

`func (o *EmailConnector) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EmailConnector) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EmailConnector) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EmailConnector) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *EmailConnector) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *EmailConnector) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *EmailConnector) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *EmailConnector) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetEmailNotifyFrom

`func (o *EmailConnector) GetEmailNotifyFrom() string`

GetEmailNotifyFrom returns the EmailNotifyFrom field if non-nil, zero value otherwise.

### GetEmailNotifyFromOk

`func (o *EmailConnector) GetEmailNotifyFromOk() (*string, bool)`

GetEmailNotifyFromOk returns a tuple with the EmailNotifyFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifyFrom

`func (o *EmailConnector) SetEmailNotifyFrom(v string)`

SetEmailNotifyFrom sets EmailNotifyFrom field to given value.

### HasEmailNotifyFrom

`func (o *EmailConnector) HasEmailNotifyFrom() bool`

HasEmailNotifyFrom returns a boolean if a field has been set.

### GetBccEmailTo

`func (o *EmailConnector) GetBccEmailTo() string`

GetBccEmailTo returns the BccEmailTo field if non-nil, zero value otherwise.

### GetBccEmailToOk

`func (o *EmailConnector) GetBccEmailToOk() (*string, bool)`

GetBccEmailToOk returns a tuple with the BccEmailTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccEmailTo

`func (o *EmailConnector) SetBccEmailTo(v string)`

SetBccEmailTo sets BccEmailTo field to given value.

### HasBccEmailTo

`func (o *EmailConnector) HasBccEmailTo() bool`

HasBccEmailTo returns a boolean if a field has been set.

### GetEmailErrorsTo

`func (o *EmailConnector) GetEmailErrorsTo() string`

GetEmailErrorsTo returns the EmailErrorsTo field if non-nil, zero value otherwise.

### GetEmailErrorsToOk

`func (o *EmailConnector) GetEmailErrorsToOk() (*string, bool)`

GetEmailErrorsToOk returns a tuple with the EmailErrorsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailErrorsTo

`func (o *EmailConnector) SetEmailErrorsTo(v string)`

SetEmailErrorsTo sets EmailErrorsTo field to given value.


### GetSetEmailToDefaultContactFlag

`func (o *EmailConnector) GetSetEmailToDefaultContactFlag() bool`

GetSetEmailToDefaultContactFlag returns the SetEmailToDefaultContactFlag field if non-nil, zero value otherwise.

### GetSetEmailToDefaultContactFlagOk

`func (o *EmailConnector) GetSetEmailToDefaultContactFlagOk() (*bool, bool)`

GetSetEmailToDefaultContactFlagOk returns a tuple with the SetEmailToDefaultContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetEmailToDefaultContactFlag

`func (o *EmailConnector) SetSetEmailToDefaultContactFlag(v bool)`

SetSetEmailToDefaultContactFlag sets SetEmailToDefaultContactFlag field to given value.

### HasSetEmailToDefaultContactFlag

`func (o *EmailConnector) HasSetEmailToDefaultContactFlag() bool`

HasSetEmailToDefaultContactFlag returns a boolean if a field has been set.

### SetSetEmailToDefaultContactFlagNil

`func (o *EmailConnector) SetSetEmailToDefaultContactFlagNil(b bool)`

 SetSetEmailToDefaultContactFlagNil sets the value for SetEmailToDefaultContactFlag to be an explicit nil

### UnsetSetEmailToDefaultContactFlag
`func (o *EmailConnector) UnsetSetEmailToDefaultContactFlag()`

UnsetSetEmailToDefaultContactFlag ensures that no value is present for SetEmailToDefaultContactFlag, not even an explicit nil
### GetNoResponseFlag

`func (o *EmailConnector) GetNoResponseFlag() bool`

GetNoResponseFlag returns the NoResponseFlag field if non-nil, zero value otherwise.

### GetNoResponseFlagOk

`func (o *EmailConnector) GetNoResponseFlagOk() (*bool, bool)`

GetNoResponseFlagOk returns a tuple with the NoResponseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoResponseFlag

`func (o *EmailConnector) SetNoResponseFlag(v bool)`

SetNoResponseFlag sets NoResponseFlag field to given value.

### HasNoResponseFlag

`func (o *EmailConnector) HasNoResponseFlag() bool`

HasNoResponseFlag returns a boolean if a field has been set.

### SetNoResponseFlagNil

`func (o *EmailConnector) SetNoResponseFlagNil(b bool)`

 SetNoResponseFlagNil sets the value for NoResponseFlag to be an explicit nil

### UnsetNoResponseFlag
`func (o *EmailConnector) UnsetNoResponseFlag()`

UnsetNoResponseFlag ensures that no value is present for NoResponseFlag, not even an explicit nil
### GetNeverRespondFlag

`func (o *EmailConnector) GetNeverRespondFlag() bool`

GetNeverRespondFlag returns the NeverRespondFlag field if non-nil, zero value otherwise.

### GetNeverRespondFlagOk

`func (o *EmailConnector) GetNeverRespondFlagOk() (*bool, bool)`

GetNeverRespondFlagOk returns a tuple with the NeverRespondFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverRespondFlag

`func (o *EmailConnector) SetNeverRespondFlag(v bool)`

SetNeverRespondFlag sets NeverRespondFlag field to given value.

### HasNeverRespondFlag

`func (o *EmailConnector) HasNeverRespondFlag() bool`

HasNeverRespondFlag returns a boolean if a field has been set.

### SetNeverRespondFlagNil

`func (o *EmailConnector) SetNeverRespondFlagNil(b bool)`

 SetNeverRespondFlagNil sets the value for NeverRespondFlag to be an explicit nil

### UnsetNeverRespondFlag
`func (o *EmailConnector) UnsetNeverRespondFlag()`

UnsetNeverRespondFlag ensures that no value is present for NeverRespondFlag, not even an explicit nil
### GetDiscardDuplicatesFlag

`func (o *EmailConnector) GetDiscardDuplicatesFlag() bool`

GetDiscardDuplicatesFlag returns the DiscardDuplicatesFlag field if non-nil, zero value otherwise.

### GetDiscardDuplicatesFlagOk

`func (o *EmailConnector) GetDiscardDuplicatesFlagOk() (*bool, bool)`

GetDiscardDuplicatesFlagOk returns a tuple with the DiscardDuplicatesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardDuplicatesFlag

`func (o *EmailConnector) SetDiscardDuplicatesFlag(v bool)`

SetDiscardDuplicatesFlag sets DiscardDuplicatesFlag field to given value.

### HasDiscardDuplicatesFlag

`func (o *EmailConnector) HasDiscardDuplicatesFlag() bool`

HasDiscardDuplicatesFlag returns a boolean if a field has been set.

### SetDiscardDuplicatesFlagNil

`func (o *EmailConnector) SetDiscardDuplicatesFlagNil(b bool)`

 SetDiscardDuplicatesFlagNil sets the value for DiscardDuplicatesFlag to be an explicit nil

### UnsetDiscardDuplicatesFlag
`func (o *EmailConnector) UnsetDiscardDuplicatesFlag()`

UnsetDiscardDuplicatesFlag ensures that no value is present for DiscardDuplicatesFlag, not even an explicit nil
### GetPostRepliesToTicketFlag

`func (o *EmailConnector) GetPostRepliesToTicketFlag() bool`

GetPostRepliesToTicketFlag returns the PostRepliesToTicketFlag field if non-nil, zero value otherwise.

### GetPostRepliesToTicketFlagOk

`func (o *EmailConnector) GetPostRepliesToTicketFlagOk() (*bool, bool)`

GetPostRepliesToTicketFlagOk returns a tuple with the PostRepliesToTicketFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRepliesToTicketFlag

`func (o *EmailConnector) SetPostRepliesToTicketFlag(v bool)`

SetPostRepliesToTicketFlag sets PostRepliesToTicketFlag field to given value.

### HasPostRepliesToTicketFlag

`func (o *EmailConnector) HasPostRepliesToTicketFlag() bool`

HasPostRepliesToTicketFlag returns a boolean if a field has been set.

### SetPostRepliesToTicketFlagNil

`func (o *EmailConnector) SetPostRepliesToTicketFlagNil(b bool)`

 SetPostRepliesToTicketFlagNil sets the value for PostRepliesToTicketFlag to be an explicit nil

### UnsetPostRepliesToTicketFlag
`func (o *EmailConnector) UnsetPostRepliesToTicketFlag()`

UnsetPostRepliesToTicketFlag ensures that no value is present for PostRepliesToTicketFlag, not even an explicit nil
### GetCreateContactFlag

`func (o *EmailConnector) GetCreateContactFlag() bool`

GetCreateContactFlag returns the CreateContactFlag field if non-nil, zero value otherwise.

### GetCreateContactFlagOk

`func (o *EmailConnector) GetCreateContactFlagOk() (*bool, bool)`

GetCreateContactFlagOk returns a tuple with the CreateContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateContactFlag

`func (o *EmailConnector) SetCreateContactFlag(v bool)`

SetCreateContactFlag sets CreateContactFlag field to given value.

### HasCreateContactFlag

`func (o *EmailConnector) HasCreateContactFlag() bool`

HasCreateContactFlag returns a boolean if a field has been set.

### SetCreateContactFlagNil

`func (o *EmailConnector) SetCreateContactFlagNil(b bool)`

 SetCreateContactFlagNil sets the value for CreateContactFlag to be an explicit nil

### UnsetCreateContactFlag
`func (o *EmailConnector) UnsetCreateContactFlag()`

UnsetCreateContactFlag ensures that no value is present for CreateContactFlag, not even an explicit nil
### GetResponseEmailForNew

`func (o *EmailConnector) GetResponseEmailForNew() string`

GetResponseEmailForNew returns the ResponseEmailForNew field if non-nil, zero value otherwise.

### GetResponseEmailForNewOk

`func (o *EmailConnector) GetResponseEmailForNewOk() (*string, bool)`

GetResponseEmailForNewOk returns a tuple with the ResponseEmailForNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseEmailForNew

`func (o *EmailConnector) SetResponseEmailForNew(v string)`

SetResponseEmailForNew sets ResponseEmailForNew field to given value.

### HasResponseEmailForNew

`func (o *EmailConnector) HasResponseEmailForNew() bool`

HasResponseEmailForNew returns a boolean if a field has been set.

### GetResponseEmailForExisting

`func (o *EmailConnector) GetResponseEmailForExisting() string`

GetResponseEmailForExisting returns the ResponseEmailForExisting field if non-nil, zero value otherwise.

### GetResponseEmailForExistingOk

`func (o *EmailConnector) GetResponseEmailForExistingOk() (*string, bool)`

GetResponseEmailForExistingOk returns a tuple with the ResponseEmailForExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseEmailForExisting

`func (o *EmailConnector) SetResponseEmailForExisting(v string)`

SetResponseEmailForExisting sets ResponseEmailForExisting field to given value.

### HasResponseEmailForExisting

`func (o *EmailConnector) HasResponseEmailForExisting() bool`

HasResponseEmailForExisting returns a boolean if a field has been set.

### GetSourceOverride

`func (o *EmailConnector) GetSourceOverride() ServiceSourceReference`

GetSourceOverride returns the SourceOverride field if non-nil, zero value otherwise.

### GetSourceOverrideOk

`func (o *EmailConnector) GetSourceOverrideOk() (*ServiceSourceReference, bool)`

GetSourceOverrideOk returns a tuple with the SourceOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOverride

`func (o *EmailConnector) SetSourceOverride(v ServiceSourceReference)`

SetSourceOverride sets SourceOverride field to given value.

### HasSourceOverride

`func (o *EmailConnector) HasSourceOverride() bool`

HasSourceOverride returns a boolean if a field has been set.

### GetPriorityOverride

`func (o *EmailConnector) GetPriorityOverride() PriorityReference`

GetPriorityOverride returns the PriorityOverride field if non-nil, zero value otherwise.

### GetPriorityOverrideOk

`func (o *EmailConnector) GetPriorityOverrideOk() (*PriorityReference, bool)`

GetPriorityOverrideOk returns a tuple with the PriorityOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityOverride

`func (o *EmailConnector) SetPriorityOverride(v PriorityReference)`

SetPriorityOverride sets PriorityOverride field to given value.

### HasPriorityOverride

`func (o *EmailConnector) HasPriorityOverride() bool`

HasPriorityOverride returns a boolean if a field has been set.

### GetTypeOverride

`func (o *EmailConnector) GetTypeOverride() ServiceTypeReference`

GetTypeOverride returns the TypeOverride field if non-nil, zero value otherwise.

### GetTypeOverrideOk

`func (o *EmailConnector) GetTypeOverrideOk() (*ServiceTypeReference, bool)`

GetTypeOverrideOk returns a tuple with the TypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOverride

`func (o *EmailConnector) SetTypeOverride(v ServiceTypeReference)`

SetTypeOverride sets TypeOverride field to given value.

### HasTypeOverride

`func (o *EmailConnector) HasTypeOverride() bool`

HasTypeOverride returns a boolean if a field has been set.

### GetSubTypeOverride

`func (o *EmailConnector) GetSubTypeOverride() ServiceSubTypeReference`

GetSubTypeOverride returns the SubTypeOverride field if non-nil, zero value otherwise.

### GetSubTypeOverrideOk

`func (o *EmailConnector) GetSubTypeOverrideOk() (*ServiceSubTypeReference, bool)`

GetSubTypeOverrideOk returns a tuple with the SubTypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTypeOverride

`func (o *EmailConnector) SetSubTypeOverride(v ServiceSubTypeReference)`

SetSubTypeOverride sets SubTypeOverride field to given value.

### HasSubTypeOverride

`func (o *EmailConnector) HasSubTypeOverride() bool`

HasSubTypeOverride returns a boolean if a field has been set.

### GetItemOverride

`func (o *EmailConnector) GetItemOverride() ServiceItemReference`

GetItemOverride returns the ItemOverride field if non-nil, zero value otherwise.

### GetItemOverrideOk

`func (o *EmailConnector) GetItemOverrideOk() (*ServiceItemReference, bool)`

GetItemOverrideOk returns a tuple with the ItemOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemOverride

`func (o *EmailConnector) SetItemOverride(v ServiceItemReference)`

SetItemOverride sets ItemOverride field to given value.

### HasItemOverride

`func (o *EmailConnector) HasItemOverride() bool`

HasItemOverride returns a boolean if a field has been set.

### GetStatusOverride

`func (o *EmailConnector) GetStatusOverride() ServiceStatusReference`

GetStatusOverride returns the StatusOverride field if non-nil, zero value otherwise.

### GetStatusOverrideOk

`func (o *EmailConnector) GetStatusOverrideOk() (*ServiceStatusReference, bool)`

GetStatusOverrideOk returns a tuple with the StatusOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusOverride

`func (o *EmailConnector) SetStatusOverride(v ServiceStatusReference)`

SetStatusOverride sets StatusOverride field to given value.

### HasStatusOverride

`func (o *EmailConnector) HasStatusOverride() bool`

HasStatusOverride returns a boolean if a field has been set.

### GetAddCcFlag

`func (o *EmailConnector) GetAddCcFlag() bool`

GetAddCcFlag returns the AddCcFlag field if non-nil, zero value otherwise.

### GetAddCcFlagOk

`func (o *EmailConnector) GetAddCcFlagOk() (*bool, bool)`

GetAddCcFlagOk returns a tuple with the AddCcFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCcFlag

`func (o *EmailConnector) SetAddCcFlag(v bool)`

SetAddCcFlag sets AddCcFlag field to given value.

### HasAddCcFlag

`func (o *EmailConnector) HasAddCcFlag() bool`

HasAddCcFlag returns a boolean if a field has been set.

### SetAddCcFlagNil

`func (o *EmailConnector) SetAddCcFlagNil(b bool)`

 SetAddCcFlagNil sets the value for AddCcFlag to be an explicit nil

### UnsetAddCcFlag
`func (o *EmailConnector) UnsetAddCcFlag()`

UnsetAddCcFlag ensures that no value is present for AddCcFlag, not even an explicit nil
### GetInboundTicketMailboxId

`func (o *EmailConnector) GetInboundTicketMailboxId() string`

GetInboundTicketMailboxId returns the InboundTicketMailboxId field if non-nil, zero value otherwise.

### GetInboundTicketMailboxIdOk

`func (o *EmailConnector) GetInboundTicketMailboxIdOk() (*string, bool)`

GetInboundTicketMailboxIdOk returns a tuple with the InboundTicketMailboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundTicketMailboxId

`func (o *EmailConnector) SetInboundTicketMailboxId(v string)`

SetInboundTicketMailboxId sets InboundTicketMailboxId field to given value.

### HasInboundTicketMailboxId

`func (o *EmailConnector) HasInboundTicketMailboxId() bool`

HasInboundTicketMailboxId returns a boolean if a field has been set.

### GetUseEmailMessageIdFlag

`func (o *EmailConnector) GetUseEmailMessageIdFlag() bool`

GetUseEmailMessageIdFlag returns the UseEmailMessageIdFlag field if non-nil, zero value otherwise.

### GetUseEmailMessageIdFlagOk

`func (o *EmailConnector) GetUseEmailMessageIdFlagOk() (*bool, bool)`

GetUseEmailMessageIdFlagOk returns a tuple with the UseEmailMessageIdFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailMessageIdFlag

`func (o *EmailConnector) SetUseEmailMessageIdFlag(v bool)`

SetUseEmailMessageIdFlag sets UseEmailMessageIdFlag field to given value.

### HasUseEmailMessageIdFlag

`func (o *EmailConnector) HasUseEmailMessageIdFlag() bool`

HasUseEmailMessageIdFlag returns a boolean if a field has been set.

### SetUseEmailMessageIdFlagNil

`func (o *EmailConnector) SetUseEmailMessageIdFlagNil(b bool)`

 SetUseEmailMessageIdFlagNil sets the value for UseEmailMessageIdFlag to be an explicit nil

### UnsetUseEmailMessageIdFlag
`func (o *EmailConnector) UnsetUseEmailMessageIdFlag()`

UnsetUseEmailMessageIdFlag ensures that no value is present for UseEmailMessageIdFlag, not even an explicit nil
### GetInfo

`func (o *EmailConnector) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *EmailConnector) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *EmailConnector) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *EmailConnector) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


