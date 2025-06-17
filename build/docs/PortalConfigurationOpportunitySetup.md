# PortalConfigurationOpportunitySetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OpportunityStatusRecIDs** | Pointer to **[]int32** |  | [optional] 
**AddAllOpportunityStatuses** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllOpportunityStatuses** | Pointer to **NullableBool** |  | [optional] 
**OpportunityTypeRecIDs** | Pointer to **[]int32** |  | [optional] 
**AddAllOpportunityTypes** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllOpportunityTypes** | Pointer to **NullableBool** |  | [optional] 
**RestrictViewByOpportunityStatusFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictViewByOpportunityTypeFlag** | Pointer to **NullableBool** |  | [optional] 
**AcceptanceChangeStatusFlag** | Pointer to **NullableBool** |  | [optional] 
**AcceptanceCreateActivityFlag** | Pointer to **NullableBool** |  | [optional] 
**AcceptanceOpportunityStatus** | Pointer to [**OpportunityStatusReference**](OpportunityStatusReference.md) |  | [optional] 
**AcceptanceSendEmailFlag** | Pointer to **NullableBool** |  | [optional] 
**AcceptanceEmailFromFirstName** | Pointer to **string** |  | [optional] 
**AcceptanceEmailFromLastName** | Pointer to **string** |  | [optional] 
**AcceptanceEmailSubject** | Pointer to **string** |  | [optional] 
**AcceptanceEmailBody** | Pointer to **string** |  | [optional] 
**AcceptanceFromEmail** | Pointer to **string** | Gets or sets             required when acceptanceSendEmailFlag is true. | [optional] 
**AcceptanceEmailActivityType** | Pointer to [**ActivityTypeReference**](ActivityTypeReference.md) |  | [optional] 
**AcceptanceEmailAssignedByMember** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**RejectionChangeStatusFlag** | Pointer to **NullableBool** |  | [optional] 
**RejectionCreateActivityFlag** | Pointer to **NullableBool** |  | [optional] 
**RejectionOpportunityStatus** | Pointer to [**OpportunityStatusReference**](OpportunityStatusReference.md) |  | [optional] 
**RejectionSendEmailFlag** | Pointer to **NullableBool** |  | [optional] 
**RejectionEmailFromFirstName** | Pointer to **string** |  | [optional] 
**RejectionEmailFromLastName** | Pointer to **string** |  | [optional] 
**RejectionFromEmail** | Pointer to **string** | Gets or sets             required when rejectionSendEmailFlag is true. | [optional] 
**RejectionEmailSubject** | Pointer to **string** |  | [optional] 
**RejectionEmailBody** | Pointer to **string** |  | [optional] 
**RejectionEmailActivityType** | Pointer to [**ActivityTypeReference**](ActivityTypeReference.md) |  | [optional] 
**RejectionEmailAssignedByMember** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**ConfirmationSendEmailFlag** | Pointer to **NullableBool** |  | [optional] 
**ConfirmationEmailUseDefaultCompanyEmailAddressFlag** | Pointer to **NullableBool** |  | [optional] 
**ConfirmationEmailFromFirstName** | Pointer to **string** |  | [optional] 
**ConfirmationEmailFromLastName** | Pointer to **string** |  | [optional] 
**ConfirmationFromEmail** | Pointer to **string** | Gets or sets             required when confirmationSendEmailFlag is true. | [optional] 
**ConfirmationEmailSubject** | Pointer to **string** |  | [optional] 
**ConfirmationEmailBody** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPortalConfigurationOpportunitySetup

`func NewPortalConfigurationOpportunitySetup() *PortalConfigurationOpportunitySetup`

NewPortalConfigurationOpportunitySetup instantiates a new PortalConfigurationOpportunitySetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalConfigurationOpportunitySetupWithDefaults

`func NewPortalConfigurationOpportunitySetupWithDefaults() *PortalConfigurationOpportunitySetup`

NewPortalConfigurationOpportunitySetupWithDefaults instantiates a new PortalConfigurationOpportunitySetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalConfigurationOpportunitySetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalConfigurationOpportunitySetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalConfigurationOpportunitySetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PortalConfigurationOpportunitySetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOpportunityStatusRecIDs

`func (o *PortalConfigurationOpportunitySetup) GetOpportunityStatusRecIDs() []int32`

GetOpportunityStatusRecIDs returns the OpportunityStatusRecIDs field if non-nil, zero value otherwise.

### GetOpportunityStatusRecIDsOk

`func (o *PortalConfigurationOpportunitySetup) GetOpportunityStatusRecIDsOk() (*[]int32, bool)`

GetOpportunityStatusRecIDsOk returns a tuple with the OpportunityStatusRecIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityStatusRecIDs

`func (o *PortalConfigurationOpportunitySetup) SetOpportunityStatusRecIDs(v []int32)`

SetOpportunityStatusRecIDs sets OpportunityStatusRecIDs field to given value.

### HasOpportunityStatusRecIDs

`func (o *PortalConfigurationOpportunitySetup) HasOpportunityStatusRecIDs() bool`

HasOpportunityStatusRecIDs returns a boolean if a field has been set.

### GetAddAllOpportunityStatuses

`func (o *PortalConfigurationOpportunitySetup) GetAddAllOpportunityStatuses() bool`

GetAddAllOpportunityStatuses returns the AddAllOpportunityStatuses field if non-nil, zero value otherwise.

### GetAddAllOpportunityStatusesOk

`func (o *PortalConfigurationOpportunitySetup) GetAddAllOpportunityStatusesOk() (*bool, bool)`

GetAddAllOpportunityStatusesOk returns a tuple with the AddAllOpportunityStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllOpportunityStatuses

`func (o *PortalConfigurationOpportunitySetup) SetAddAllOpportunityStatuses(v bool)`

SetAddAllOpportunityStatuses sets AddAllOpportunityStatuses field to given value.

### HasAddAllOpportunityStatuses

`func (o *PortalConfigurationOpportunitySetup) HasAddAllOpportunityStatuses() bool`

HasAddAllOpportunityStatuses returns a boolean if a field has been set.

### SetAddAllOpportunityStatusesNil

`func (o *PortalConfigurationOpportunitySetup) SetAddAllOpportunityStatusesNil(b bool)`

 SetAddAllOpportunityStatusesNil sets the value for AddAllOpportunityStatuses to be an explicit nil

### UnsetAddAllOpportunityStatuses
`func (o *PortalConfigurationOpportunitySetup) UnsetAddAllOpportunityStatuses()`

UnsetAddAllOpportunityStatuses ensures that no value is present for AddAllOpportunityStatuses, not even an explicit nil
### GetRemoveAllOpportunityStatuses

`func (o *PortalConfigurationOpportunitySetup) GetRemoveAllOpportunityStatuses() bool`

GetRemoveAllOpportunityStatuses returns the RemoveAllOpportunityStatuses field if non-nil, zero value otherwise.

### GetRemoveAllOpportunityStatusesOk

`func (o *PortalConfigurationOpportunitySetup) GetRemoveAllOpportunityStatusesOk() (*bool, bool)`

GetRemoveAllOpportunityStatusesOk returns a tuple with the RemoveAllOpportunityStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllOpportunityStatuses

`func (o *PortalConfigurationOpportunitySetup) SetRemoveAllOpportunityStatuses(v bool)`

SetRemoveAllOpportunityStatuses sets RemoveAllOpportunityStatuses field to given value.

### HasRemoveAllOpportunityStatuses

`func (o *PortalConfigurationOpportunitySetup) HasRemoveAllOpportunityStatuses() bool`

HasRemoveAllOpportunityStatuses returns a boolean if a field has been set.

### SetRemoveAllOpportunityStatusesNil

`func (o *PortalConfigurationOpportunitySetup) SetRemoveAllOpportunityStatusesNil(b bool)`

 SetRemoveAllOpportunityStatusesNil sets the value for RemoveAllOpportunityStatuses to be an explicit nil

### UnsetRemoveAllOpportunityStatuses
`func (o *PortalConfigurationOpportunitySetup) UnsetRemoveAllOpportunityStatuses()`

UnsetRemoveAllOpportunityStatuses ensures that no value is present for RemoveAllOpportunityStatuses, not even an explicit nil
### GetOpportunityTypeRecIDs

`func (o *PortalConfigurationOpportunitySetup) GetOpportunityTypeRecIDs() []int32`

GetOpportunityTypeRecIDs returns the OpportunityTypeRecIDs field if non-nil, zero value otherwise.

### GetOpportunityTypeRecIDsOk

`func (o *PortalConfigurationOpportunitySetup) GetOpportunityTypeRecIDsOk() (*[]int32, bool)`

GetOpportunityTypeRecIDsOk returns a tuple with the OpportunityTypeRecIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityTypeRecIDs

`func (o *PortalConfigurationOpportunitySetup) SetOpportunityTypeRecIDs(v []int32)`

SetOpportunityTypeRecIDs sets OpportunityTypeRecIDs field to given value.

### HasOpportunityTypeRecIDs

`func (o *PortalConfigurationOpportunitySetup) HasOpportunityTypeRecIDs() bool`

HasOpportunityTypeRecIDs returns a boolean if a field has been set.

### GetAddAllOpportunityTypes

`func (o *PortalConfigurationOpportunitySetup) GetAddAllOpportunityTypes() bool`

GetAddAllOpportunityTypes returns the AddAllOpportunityTypes field if non-nil, zero value otherwise.

### GetAddAllOpportunityTypesOk

`func (o *PortalConfigurationOpportunitySetup) GetAddAllOpportunityTypesOk() (*bool, bool)`

GetAddAllOpportunityTypesOk returns a tuple with the AddAllOpportunityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllOpportunityTypes

`func (o *PortalConfigurationOpportunitySetup) SetAddAllOpportunityTypes(v bool)`

SetAddAllOpportunityTypes sets AddAllOpportunityTypes field to given value.

### HasAddAllOpportunityTypes

`func (o *PortalConfigurationOpportunitySetup) HasAddAllOpportunityTypes() bool`

HasAddAllOpportunityTypes returns a boolean if a field has been set.

### SetAddAllOpportunityTypesNil

`func (o *PortalConfigurationOpportunitySetup) SetAddAllOpportunityTypesNil(b bool)`

 SetAddAllOpportunityTypesNil sets the value for AddAllOpportunityTypes to be an explicit nil

### UnsetAddAllOpportunityTypes
`func (o *PortalConfigurationOpportunitySetup) UnsetAddAllOpportunityTypes()`

UnsetAddAllOpportunityTypes ensures that no value is present for AddAllOpportunityTypes, not even an explicit nil
### GetRemoveAllOpportunityTypes

`func (o *PortalConfigurationOpportunitySetup) GetRemoveAllOpportunityTypes() bool`

GetRemoveAllOpportunityTypes returns the RemoveAllOpportunityTypes field if non-nil, zero value otherwise.

### GetRemoveAllOpportunityTypesOk

`func (o *PortalConfigurationOpportunitySetup) GetRemoveAllOpportunityTypesOk() (*bool, bool)`

GetRemoveAllOpportunityTypesOk returns a tuple with the RemoveAllOpportunityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllOpportunityTypes

`func (o *PortalConfigurationOpportunitySetup) SetRemoveAllOpportunityTypes(v bool)`

SetRemoveAllOpportunityTypes sets RemoveAllOpportunityTypes field to given value.

### HasRemoveAllOpportunityTypes

`func (o *PortalConfigurationOpportunitySetup) HasRemoveAllOpportunityTypes() bool`

HasRemoveAllOpportunityTypes returns a boolean if a field has been set.

### SetRemoveAllOpportunityTypesNil

`func (o *PortalConfigurationOpportunitySetup) SetRemoveAllOpportunityTypesNil(b bool)`

 SetRemoveAllOpportunityTypesNil sets the value for RemoveAllOpportunityTypes to be an explicit nil

### UnsetRemoveAllOpportunityTypes
`func (o *PortalConfigurationOpportunitySetup) UnsetRemoveAllOpportunityTypes()`

UnsetRemoveAllOpportunityTypes ensures that no value is present for RemoveAllOpportunityTypes, not even an explicit nil
### GetRestrictViewByOpportunityStatusFlag

`func (o *PortalConfigurationOpportunitySetup) GetRestrictViewByOpportunityStatusFlag() bool`

GetRestrictViewByOpportunityStatusFlag returns the RestrictViewByOpportunityStatusFlag field if non-nil, zero value otherwise.

### GetRestrictViewByOpportunityStatusFlagOk

`func (o *PortalConfigurationOpportunitySetup) GetRestrictViewByOpportunityStatusFlagOk() (*bool, bool)`

GetRestrictViewByOpportunityStatusFlagOk returns a tuple with the RestrictViewByOpportunityStatusFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictViewByOpportunityStatusFlag

`func (o *PortalConfigurationOpportunitySetup) SetRestrictViewByOpportunityStatusFlag(v bool)`

SetRestrictViewByOpportunityStatusFlag sets RestrictViewByOpportunityStatusFlag field to given value.

### HasRestrictViewByOpportunityStatusFlag

`func (o *PortalConfigurationOpportunitySetup) HasRestrictViewByOpportunityStatusFlag() bool`

HasRestrictViewByOpportunityStatusFlag returns a boolean if a field has been set.

### SetRestrictViewByOpportunityStatusFlagNil

`func (o *PortalConfigurationOpportunitySetup) SetRestrictViewByOpportunityStatusFlagNil(b bool)`

 SetRestrictViewByOpportunityStatusFlagNil sets the value for RestrictViewByOpportunityStatusFlag to be an explicit nil

### UnsetRestrictViewByOpportunityStatusFlag
`func (o *PortalConfigurationOpportunitySetup) UnsetRestrictViewByOpportunityStatusFlag()`

UnsetRestrictViewByOpportunityStatusFlag ensures that no value is present for RestrictViewByOpportunityStatusFlag, not even an explicit nil
### GetRestrictViewByOpportunityTypeFlag

`func (o *PortalConfigurationOpportunitySetup) GetRestrictViewByOpportunityTypeFlag() bool`

GetRestrictViewByOpportunityTypeFlag returns the RestrictViewByOpportunityTypeFlag field if non-nil, zero value otherwise.

### GetRestrictViewByOpportunityTypeFlagOk

`func (o *PortalConfigurationOpportunitySetup) GetRestrictViewByOpportunityTypeFlagOk() (*bool, bool)`

GetRestrictViewByOpportunityTypeFlagOk returns a tuple with the RestrictViewByOpportunityTypeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictViewByOpportunityTypeFlag

`func (o *PortalConfigurationOpportunitySetup) SetRestrictViewByOpportunityTypeFlag(v bool)`

SetRestrictViewByOpportunityTypeFlag sets RestrictViewByOpportunityTypeFlag field to given value.

### HasRestrictViewByOpportunityTypeFlag

`func (o *PortalConfigurationOpportunitySetup) HasRestrictViewByOpportunityTypeFlag() bool`

HasRestrictViewByOpportunityTypeFlag returns a boolean if a field has been set.

### SetRestrictViewByOpportunityTypeFlagNil

`func (o *PortalConfigurationOpportunitySetup) SetRestrictViewByOpportunityTypeFlagNil(b bool)`

 SetRestrictViewByOpportunityTypeFlagNil sets the value for RestrictViewByOpportunityTypeFlag to be an explicit nil

### UnsetRestrictViewByOpportunityTypeFlag
`func (o *PortalConfigurationOpportunitySetup) UnsetRestrictViewByOpportunityTypeFlag()`

UnsetRestrictViewByOpportunityTypeFlag ensures that no value is present for RestrictViewByOpportunityTypeFlag, not even an explicit nil
### GetAcceptanceChangeStatusFlag

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceChangeStatusFlag() bool`

GetAcceptanceChangeStatusFlag returns the AcceptanceChangeStatusFlag field if non-nil, zero value otherwise.

### GetAcceptanceChangeStatusFlagOk

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceChangeStatusFlagOk() (*bool, bool)`

GetAcceptanceChangeStatusFlagOk returns a tuple with the AcceptanceChangeStatusFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceChangeStatusFlag

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceChangeStatusFlag(v bool)`

SetAcceptanceChangeStatusFlag sets AcceptanceChangeStatusFlag field to given value.

### HasAcceptanceChangeStatusFlag

`func (o *PortalConfigurationOpportunitySetup) HasAcceptanceChangeStatusFlag() bool`

HasAcceptanceChangeStatusFlag returns a boolean if a field has been set.

### SetAcceptanceChangeStatusFlagNil

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceChangeStatusFlagNil(b bool)`

 SetAcceptanceChangeStatusFlagNil sets the value for AcceptanceChangeStatusFlag to be an explicit nil

### UnsetAcceptanceChangeStatusFlag
`func (o *PortalConfigurationOpportunitySetup) UnsetAcceptanceChangeStatusFlag()`

UnsetAcceptanceChangeStatusFlag ensures that no value is present for AcceptanceChangeStatusFlag, not even an explicit nil
### GetAcceptanceCreateActivityFlag

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceCreateActivityFlag() bool`

GetAcceptanceCreateActivityFlag returns the AcceptanceCreateActivityFlag field if non-nil, zero value otherwise.

### GetAcceptanceCreateActivityFlagOk

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceCreateActivityFlagOk() (*bool, bool)`

GetAcceptanceCreateActivityFlagOk returns a tuple with the AcceptanceCreateActivityFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceCreateActivityFlag

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceCreateActivityFlag(v bool)`

SetAcceptanceCreateActivityFlag sets AcceptanceCreateActivityFlag field to given value.

### HasAcceptanceCreateActivityFlag

`func (o *PortalConfigurationOpportunitySetup) HasAcceptanceCreateActivityFlag() bool`

HasAcceptanceCreateActivityFlag returns a boolean if a field has been set.

### SetAcceptanceCreateActivityFlagNil

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceCreateActivityFlagNil(b bool)`

 SetAcceptanceCreateActivityFlagNil sets the value for AcceptanceCreateActivityFlag to be an explicit nil

### UnsetAcceptanceCreateActivityFlag
`func (o *PortalConfigurationOpportunitySetup) UnsetAcceptanceCreateActivityFlag()`

UnsetAcceptanceCreateActivityFlag ensures that no value is present for AcceptanceCreateActivityFlag, not even an explicit nil
### GetAcceptanceOpportunityStatus

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceOpportunityStatus() OpportunityStatusReference`

GetAcceptanceOpportunityStatus returns the AcceptanceOpportunityStatus field if non-nil, zero value otherwise.

### GetAcceptanceOpportunityStatusOk

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceOpportunityStatusOk() (*OpportunityStatusReference, bool)`

GetAcceptanceOpportunityStatusOk returns a tuple with the AcceptanceOpportunityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceOpportunityStatus

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceOpportunityStatus(v OpportunityStatusReference)`

SetAcceptanceOpportunityStatus sets AcceptanceOpportunityStatus field to given value.

### HasAcceptanceOpportunityStatus

`func (o *PortalConfigurationOpportunitySetup) HasAcceptanceOpportunityStatus() bool`

HasAcceptanceOpportunityStatus returns a boolean if a field has been set.

### GetAcceptanceSendEmailFlag

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceSendEmailFlag() bool`

GetAcceptanceSendEmailFlag returns the AcceptanceSendEmailFlag field if non-nil, zero value otherwise.

### GetAcceptanceSendEmailFlagOk

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceSendEmailFlagOk() (*bool, bool)`

GetAcceptanceSendEmailFlagOk returns a tuple with the AcceptanceSendEmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceSendEmailFlag

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceSendEmailFlag(v bool)`

SetAcceptanceSendEmailFlag sets AcceptanceSendEmailFlag field to given value.

### HasAcceptanceSendEmailFlag

`func (o *PortalConfigurationOpportunitySetup) HasAcceptanceSendEmailFlag() bool`

HasAcceptanceSendEmailFlag returns a boolean if a field has been set.

### SetAcceptanceSendEmailFlagNil

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceSendEmailFlagNil(b bool)`

 SetAcceptanceSendEmailFlagNil sets the value for AcceptanceSendEmailFlag to be an explicit nil

### UnsetAcceptanceSendEmailFlag
`func (o *PortalConfigurationOpportunitySetup) UnsetAcceptanceSendEmailFlag()`

UnsetAcceptanceSendEmailFlag ensures that no value is present for AcceptanceSendEmailFlag, not even an explicit nil
### GetAcceptanceEmailFromFirstName

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceEmailFromFirstName() string`

GetAcceptanceEmailFromFirstName returns the AcceptanceEmailFromFirstName field if non-nil, zero value otherwise.

### GetAcceptanceEmailFromFirstNameOk

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceEmailFromFirstNameOk() (*string, bool)`

GetAcceptanceEmailFromFirstNameOk returns a tuple with the AcceptanceEmailFromFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceEmailFromFirstName

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceEmailFromFirstName(v string)`

SetAcceptanceEmailFromFirstName sets AcceptanceEmailFromFirstName field to given value.

### HasAcceptanceEmailFromFirstName

`func (o *PortalConfigurationOpportunitySetup) HasAcceptanceEmailFromFirstName() bool`

HasAcceptanceEmailFromFirstName returns a boolean if a field has been set.

### GetAcceptanceEmailFromLastName

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceEmailFromLastName() string`

GetAcceptanceEmailFromLastName returns the AcceptanceEmailFromLastName field if non-nil, zero value otherwise.

### GetAcceptanceEmailFromLastNameOk

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceEmailFromLastNameOk() (*string, bool)`

GetAcceptanceEmailFromLastNameOk returns a tuple with the AcceptanceEmailFromLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceEmailFromLastName

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceEmailFromLastName(v string)`

SetAcceptanceEmailFromLastName sets AcceptanceEmailFromLastName field to given value.

### HasAcceptanceEmailFromLastName

`func (o *PortalConfigurationOpportunitySetup) HasAcceptanceEmailFromLastName() bool`

HasAcceptanceEmailFromLastName returns a boolean if a field has been set.

### GetAcceptanceEmailSubject

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceEmailSubject() string`

GetAcceptanceEmailSubject returns the AcceptanceEmailSubject field if non-nil, zero value otherwise.

### GetAcceptanceEmailSubjectOk

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceEmailSubjectOk() (*string, bool)`

GetAcceptanceEmailSubjectOk returns a tuple with the AcceptanceEmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceEmailSubject

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceEmailSubject(v string)`

SetAcceptanceEmailSubject sets AcceptanceEmailSubject field to given value.

### HasAcceptanceEmailSubject

`func (o *PortalConfigurationOpportunitySetup) HasAcceptanceEmailSubject() bool`

HasAcceptanceEmailSubject returns a boolean if a field has been set.

### GetAcceptanceEmailBody

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceEmailBody() string`

GetAcceptanceEmailBody returns the AcceptanceEmailBody field if non-nil, zero value otherwise.

### GetAcceptanceEmailBodyOk

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceEmailBodyOk() (*string, bool)`

GetAcceptanceEmailBodyOk returns a tuple with the AcceptanceEmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceEmailBody

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceEmailBody(v string)`

SetAcceptanceEmailBody sets AcceptanceEmailBody field to given value.

### HasAcceptanceEmailBody

`func (o *PortalConfigurationOpportunitySetup) HasAcceptanceEmailBody() bool`

HasAcceptanceEmailBody returns a boolean if a field has been set.

### GetAcceptanceFromEmail

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceFromEmail() string`

GetAcceptanceFromEmail returns the AcceptanceFromEmail field if non-nil, zero value otherwise.

### GetAcceptanceFromEmailOk

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceFromEmailOk() (*string, bool)`

GetAcceptanceFromEmailOk returns a tuple with the AcceptanceFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceFromEmail

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceFromEmail(v string)`

SetAcceptanceFromEmail sets AcceptanceFromEmail field to given value.

### HasAcceptanceFromEmail

`func (o *PortalConfigurationOpportunitySetup) HasAcceptanceFromEmail() bool`

HasAcceptanceFromEmail returns a boolean if a field has been set.

### GetAcceptanceEmailActivityType

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceEmailActivityType() ActivityTypeReference`

GetAcceptanceEmailActivityType returns the AcceptanceEmailActivityType field if non-nil, zero value otherwise.

### GetAcceptanceEmailActivityTypeOk

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceEmailActivityTypeOk() (*ActivityTypeReference, bool)`

GetAcceptanceEmailActivityTypeOk returns a tuple with the AcceptanceEmailActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceEmailActivityType

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceEmailActivityType(v ActivityTypeReference)`

SetAcceptanceEmailActivityType sets AcceptanceEmailActivityType field to given value.

### HasAcceptanceEmailActivityType

`func (o *PortalConfigurationOpportunitySetup) HasAcceptanceEmailActivityType() bool`

HasAcceptanceEmailActivityType returns a boolean if a field has been set.

### GetAcceptanceEmailAssignedByMember

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceEmailAssignedByMember() MemberReference`

GetAcceptanceEmailAssignedByMember returns the AcceptanceEmailAssignedByMember field if non-nil, zero value otherwise.

### GetAcceptanceEmailAssignedByMemberOk

`func (o *PortalConfigurationOpportunitySetup) GetAcceptanceEmailAssignedByMemberOk() (*MemberReference, bool)`

GetAcceptanceEmailAssignedByMemberOk returns a tuple with the AcceptanceEmailAssignedByMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceEmailAssignedByMember

`func (o *PortalConfigurationOpportunitySetup) SetAcceptanceEmailAssignedByMember(v MemberReference)`

SetAcceptanceEmailAssignedByMember sets AcceptanceEmailAssignedByMember field to given value.

### HasAcceptanceEmailAssignedByMember

`func (o *PortalConfigurationOpportunitySetup) HasAcceptanceEmailAssignedByMember() bool`

HasAcceptanceEmailAssignedByMember returns a boolean if a field has been set.

### GetRejectionChangeStatusFlag

`func (o *PortalConfigurationOpportunitySetup) GetRejectionChangeStatusFlag() bool`

GetRejectionChangeStatusFlag returns the RejectionChangeStatusFlag field if non-nil, zero value otherwise.

### GetRejectionChangeStatusFlagOk

`func (o *PortalConfigurationOpportunitySetup) GetRejectionChangeStatusFlagOk() (*bool, bool)`

GetRejectionChangeStatusFlagOk returns a tuple with the RejectionChangeStatusFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionChangeStatusFlag

`func (o *PortalConfigurationOpportunitySetup) SetRejectionChangeStatusFlag(v bool)`

SetRejectionChangeStatusFlag sets RejectionChangeStatusFlag field to given value.

### HasRejectionChangeStatusFlag

`func (o *PortalConfigurationOpportunitySetup) HasRejectionChangeStatusFlag() bool`

HasRejectionChangeStatusFlag returns a boolean if a field has been set.

### SetRejectionChangeStatusFlagNil

`func (o *PortalConfigurationOpportunitySetup) SetRejectionChangeStatusFlagNil(b bool)`

 SetRejectionChangeStatusFlagNil sets the value for RejectionChangeStatusFlag to be an explicit nil

### UnsetRejectionChangeStatusFlag
`func (o *PortalConfigurationOpportunitySetup) UnsetRejectionChangeStatusFlag()`

UnsetRejectionChangeStatusFlag ensures that no value is present for RejectionChangeStatusFlag, not even an explicit nil
### GetRejectionCreateActivityFlag

`func (o *PortalConfigurationOpportunitySetup) GetRejectionCreateActivityFlag() bool`

GetRejectionCreateActivityFlag returns the RejectionCreateActivityFlag field if non-nil, zero value otherwise.

### GetRejectionCreateActivityFlagOk

`func (o *PortalConfigurationOpportunitySetup) GetRejectionCreateActivityFlagOk() (*bool, bool)`

GetRejectionCreateActivityFlagOk returns a tuple with the RejectionCreateActivityFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionCreateActivityFlag

`func (o *PortalConfigurationOpportunitySetup) SetRejectionCreateActivityFlag(v bool)`

SetRejectionCreateActivityFlag sets RejectionCreateActivityFlag field to given value.

### HasRejectionCreateActivityFlag

`func (o *PortalConfigurationOpportunitySetup) HasRejectionCreateActivityFlag() bool`

HasRejectionCreateActivityFlag returns a boolean if a field has been set.

### SetRejectionCreateActivityFlagNil

`func (o *PortalConfigurationOpportunitySetup) SetRejectionCreateActivityFlagNil(b bool)`

 SetRejectionCreateActivityFlagNil sets the value for RejectionCreateActivityFlag to be an explicit nil

### UnsetRejectionCreateActivityFlag
`func (o *PortalConfigurationOpportunitySetup) UnsetRejectionCreateActivityFlag()`

UnsetRejectionCreateActivityFlag ensures that no value is present for RejectionCreateActivityFlag, not even an explicit nil
### GetRejectionOpportunityStatus

`func (o *PortalConfigurationOpportunitySetup) GetRejectionOpportunityStatus() OpportunityStatusReference`

GetRejectionOpportunityStatus returns the RejectionOpportunityStatus field if non-nil, zero value otherwise.

### GetRejectionOpportunityStatusOk

`func (o *PortalConfigurationOpportunitySetup) GetRejectionOpportunityStatusOk() (*OpportunityStatusReference, bool)`

GetRejectionOpportunityStatusOk returns a tuple with the RejectionOpportunityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionOpportunityStatus

`func (o *PortalConfigurationOpportunitySetup) SetRejectionOpportunityStatus(v OpportunityStatusReference)`

SetRejectionOpportunityStatus sets RejectionOpportunityStatus field to given value.

### HasRejectionOpportunityStatus

`func (o *PortalConfigurationOpportunitySetup) HasRejectionOpportunityStatus() bool`

HasRejectionOpportunityStatus returns a boolean if a field has been set.

### GetRejectionSendEmailFlag

`func (o *PortalConfigurationOpportunitySetup) GetRejectionSendEmailFlag() bool`

GetRejectionSendEmailFlag returns the RejectionSendEmailFlag field if non-nil, zero value otherwise.

### GetRejectionSendEmailFlagOk

`func (o *PortalConfigurationOpportunitySetup) GetRejectionSendEmailFlagOk() (*bool, bool)`

GetRejectionSendEmailFlagOk returns a tuple with the RejectionSendEmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionSendEmailFlag

`func (o *PortalConfigurationOpportunitySetup) SetRejectionSendEmailFlag(v bool)`

SetRejectionSendEmailFlag sets RejectionSendEmailFlag field to given value.

### HasRejectionSendEmailFlag

`func (o *PortalConfigurationOpportunitySetup) HasRejectionSendEmailFlag() bool`

HasRejectionSendEmailFlag returns a boolean if a field has been set.

### SetRejectionSendEmailFlagNil

`func (o *PortalConfigurationOpportunitySetup) SetRejectionSendEmailFlagNil(b bool)`

 SetRejectionSendEmailFlagNil sets the value for RejectionSendEmailFlag to be an explicit nil

### UnsetRejectionSendEmailFlag
`func (o *PortalConfigurationOpportunitySetup) UnsetRejectionSendEmailFlag()`

UnsetRejectionSendEmailFlag ensures that no value is present for RejectionSendEmailFlag, not even an explicit nil
### GetRejectionEmailFromFirstName

`func (o *PortalConfigurationOpportunitySetup) GetRejectionEmailFromFirstName() string`

GetRejectionEmailFromFirstName returns the RejectionEmailFromFirstName field if non-nil, zero value otherwise.

### GetRejectionEmailFromFirstNameOk

`func (o *PortalConfigurationOpportunitySetup) GetRejectionEmailFromFirstNameOk() (*string, bool)`

GetRejectionEmailFromFirstNameOk returns a tuple with the RejectionEmailFromFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionEmailFromFirstName

`func (o *PortalConfigurationOpportunitySetup) SetRejectionEmailFromFirstName(v string)`

SetRejectionEmailFromFirstName sets RejectionEmailFromFirstName field to given value.

### HasRejectionEmailFromFirstName

`func (o *PortalConfigurationOpportunitySetup) HasRejectionEmailFromFirstName() bool`

HasRejectionEmailFromFirstName returns a boolean if a field has been set.

### GetRejectionEmailFromLastName

`func (o *PortalConfigurationOpportunitySetup) GetRejectionEmailFromLastName() string`

GetRejectionEmailFromLastName returns the RejectionEmailFromLastName field if non-nil, zero value otherwise.

### GetRejectionEmailFromLastNameOk

`func (o *PortalConfigurationOpportunitySetup) GetRejectionEmailFromLastNameOk() (*string, bool)`

GetRejectionEmailFromLastNameOk returns a tuple with the RejectionEmailFromLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionEmailFromLastName

`func (o *PortalConfigurationOpportunitySetup) SetRejectionEmailFromLastName(v string)`

SetRejectionEmailFromLastName sets RejectionEmailFromLastName field to given value.

### HasRejectionEmailFromLastName

`func (o *PortalConfigurationOpportunitySetup) HasRejectionEmailFromLastName() bool`

HasRejectionEmailFromLastName returns a boolean if a field has been set.

### GetRejectionFromEmail

`func (o *PortalConfigurationOpportunitySetup) GetRejectionFromEmail() string`

GetRejectionFromEmail returns the RejectionFromEmail field if non-nil, zero value otherwise.

### GetRejectionFromEmailOk

`func (o *PortalConfigurationOpportunitySetup) GetRejectionFromEmailOk() (*string, bool)`

GetRejectionFromEmailOk returns a tuple with the RejectionFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionFromEmail

`func (o *PortalConfigurationOpportunitySetup) SetRejectionFromEmail(v string)`

SetRejectionFromEmail sets RejectionFromEmail field to given value.

### HasRejectionFromEmail

`func (o *PortalConfigurationOpportunitySetup) HasRejectionFromEmail() bool`

HasRejectionFromEmail returns a boolean if a field has been set.

### GetRejectionEmailSubject

`func (o *PortalConfigurationOpportunitySetup) GetRejectionEmailSubject() string`

GetRejectionEmailSubject returns the RejectionEmailSubject field if non-nil, zero value otherwise.

### GetRejectionEmailSubjectOk

`func (o *PortalConfigurationOpportunitySetup) GetRejectionEmailSubjectOk() (*string, bool)`

GetRejectionEmailSubjectOk returns a tuple with the RejectionEmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionEmailSubject

`func (o *PortalConfigurationOpportunitySetup) SetRejectionEmailSubject(v string)`

SetRejectionEmailSubject sets RejectionEmailSubject field to given value.

### HasRejectionEmailSubject

`func (o *PortalConfigurationOpportunitySetup) HasRejectionEmailSubject() bool`

HasRejectionEmailSubject returns a boolean if a field has been set.

### GetRejectionEmailBody

`func (o *PortalConfigurationOpportunitySetup) GetRejectionEmailBody() string`

GetRejectionEmailBody returns the RejectionEmailBody field if non-nil, zero value otherwise.

### GetRejectionEmailBodyOk

`func (o *PortalConfigurationOpportunitySetup) GetRejectionEmailBodyOk() (*string, bool)`

GetRejectionEmailBodyOk returns a tuple with the RejectionEmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionEmailBody

`func (o *PortalConfigurationOpportunitySetup) SetRejectionEmailBody(v string)`

SetRejectionEmailBody sets RejectionEmailBody field to given value.

### HasRejectionEmailBody

`func (o *PortalConfigurationOpportunitySetup) HasRejectionEmailBody() bool`

HasRejectionEmailBody returns a boolean if a field has been set.

### GetRejectionEmailActivityType

`func (o *PortalConfigurationOpportunitySetup) GetRejectionEmailActivityType() ActivityTypeReference`

GetRejectionEmailActivityType returns the RejectionEmailActivityType field if non-nil, zero value otherwise.

### GetRejectionEmailActivityTypeOk

`func (o *PortalConfigurationOpportunitySetup) GetRejectionEmailActivityTypeOk() (*ActivityTypeReference, bool)`

GetRejectionEmailActivityTypeOk returns a tuple with the RejectionEmailActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionEmailActivityType

`func (o *PortalConfigurationOpportunitySetup) SetRejectionEmailActivityType(v ActivityTypeReference)`

SetRejectionEmailActivityType sets RejectionEmailActivityType field to given value.

### HasRejectionEmailActivityType

`func (o *PortalConfigurationOpportunitySetup) HasRejectionEmailActivityType() bool`

HasRejectionEmailActivityType returns a boolean if a field has been set.

### GetRejectionEmailAssignedByMember

`func (o *PortalConfigurationOpportunitySetup) GetRejectionEmailAssignedByMember() MemberReference`

GetRejectionEmailAssignedByMember returns the RejectionEmailAssignedByMember field if non-nil, zero value otherwise.

### GetRejectionEmailAssignedByMemberOk

`func (o *PortalConfigurationOpportunitySetup) GetRejectionEmailAssignedByMemberOk() (*MemberReference, bool)`

GetRejectionEmailAssignedByMemberOk returns a tuple with the RejectionEmailAssignedByMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionEmailAssignedByMember

`func (o *PortalConfigurationOpportunitySetup) SetRejectionEmailAssignedByMember(v MemberReference)`

SetRejectionEmailAssignedByMember sets RejectionEmailAssignedByMember field to given value.

### HasRejectionEmailAssignedByMember

`func (o *PortalConfigurationOpportunitySetup) HasRejectionEmailAssignedByMember() bool`

HasRejectionEmailAssignedByMember returns a boolean if a field has been set.

### GetConfirmationSendEmailFlag

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationSendEmailFlag() bool`

GetConfirmationSendEmailFlag returns the ConfirmationSendEmailFlag field if non-nil, zero value otherwise.

### GetConfirmationSendEmailFlagOk

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationSendEmailFlagOk() (*bool, bool)`

GetConfirmationSendEmailFlagOk returns a tuple with the ConfirmationSendEmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationSendEmailFlag

`func (o *PortalConfigurationOpportunitySetup) SetConfirmationSendEmailFlag(v bool)`

SetConfirmationSendEmailFlag sets ConfirmationSendEmailFlag field to given value.

### HasConfirmationSendEmailFlag

`func (o *PortalConfigurationOpportunitySetup) HasConfirmationSendEmailFlag() bool`

HasConfirmationSendEmailFlag returns a boolean if a field has been set.

### SetConfirmationSendEmailFlagNil

`func (o *PortalConfigurationOpportunitySetup) SetConfirmationSendEmailFlagNil(b bool)`

 SetConfirmationSendEmailFlagNil sets the value for ConfirmationSendEmailFlag to be an explicit nil

### UnsetConfirmationSendEmailFlag
`func (o *PortalConfigurationOpportunitySetup) UnsetConfirmationSendEmailFlag()`

UnsetConfirmationSendEmailFlag ensures that no value is present for ConfirmationSendEmailFlag, not even an explicit nil
### GetConfirmationEmailUseDefaultCompanyEmailAddressFlag

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationEmailUseDefaultCompanyEmailAddressFlag() bool`

GetConfirmationEmailUseDefaultCompanyEmailAddressFlag returns the ConfirmationEmailUseDefaultCompanyEmailAddressFlag field if non-nil, zero value otherwise.

### GetConfirmationEmailUseDefaultCompanyEmailAddressFlagOk

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationEmailUseDefaultCompanyEmailAddressFlagOk() (*bool, bool)`

GetConfirmationEmailUseDefaultCompanyEmailAddressFlagOk returns a tuple with the ConfirmationEmailUseDefaultCompanyEmailAddressFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationEmailUseDefaultCompanyEmailAddressFlag

`func (o *PortalConfigurationOpportunitySetup) SetConfirmationEmailUseDefaultCompanyEmailAddressFlag(v bool)`

SetConfirmationEmailUseDefaultCompanyEmailAddressFlag sets ConfirmationEmailUseDefaultCompanyEmailAddressFlag field to given value.

### HasConfirmationEmailUseDefaultCompanyEmailAddressFlag

`func (o *PortalConfigurationOpportunitySetup) HasConfirmationEmailUseDefaultCompanyEmailAddressFlag() bool`

HasConfirmationEmailUseDefaultCompanyEmailAddressFlag returns a boolean if a field has been set.

### SetConfirmationEmailUseDefaultCompanyEmailAddressFlagNil

`func (o *PortalConfigurationOpportunitySetup) SetConfirmationEmailUseDefaultCompanyEmailAddressFlagNil(b bool)`

 SetConfirmationEmailUseDefaultCompanyEmailAddressFlagNil sets the value for ConfirmationEmailUseDefaultCompanyEmailAddressFlag to be an explicit nil

### UnsetConfirmationEmailUseDefaultCompanyEmailAddressFlag
`func (o *PortalConfigurationOpportunitySetup) UnsetConfirmationEmailUseDefaultCompanyEmailAddressFlag()`

UnsetConfirmationEmailUseDefaultCompanyEmailAddressFlag ensures that no value is present for ConfirmationEmailUseDefaultCompanyEmailAddressFlag, not even an explicit nil
### GetConfirmationEmailFromFirstName

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationEmailFromFirstName() string`

GetConfirmationEmailFromFirstName returns the ConfirmationEmailFromFirstName field if non-nil, zero value otherwise.

### GetConfirmationEmailFromFirstNameOk

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationEmailFromFirstNameOk() (*string, bool)`

GetConfirmationEmailFromFirstNameOk returns a tuple with the ConfirmationEmailFromFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationEmailFromFirstName

`func (o *PortalConfigurationOpportunitySetup) SetConfirmationEmailFromFirstName(v string)`

SetConfirmationEmailFromFirstName sets ConfirmationEmailFromFirstName field to given value.

### HasConfirmationEmailFromFirstName

`func (o *PortalConfigurationOpportunitySetup) HasConfirmationEmailFromFirstName() bool`

HasConfirmationEmailFromFirstName returns a boolean if a field has been set.

### GetConfirmationEmailFromLastName

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationEmailFromLastName() string`

GetConfirmationEmailFromLastName returns the ConfirmationEmailFromLastName field if non-nil, zero value otherwise.

### GetConfirmationEmailFromLastNameOk

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationEmailFromLastNameOk() (*string, bool)`

GetConfirmationEmailFromLastNameOk returns a tuple with the ConfirmationEmailFromLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationEmailFromLastName

`func (o *PortalConfigurationOpportunitySetup) SetConfirmationEmailFromLastName(v string)`

SetConfirmationEmailFromLastName sets ConfirmationEmailFromLastName field to given value.

### HasConfirmationEmailFromLastName

`func (o *PortalConfigurationOpportunitySetup) HasConfirmationEmailFromLastName() bool`

HasConfirmationEmailFromLastName returns a boolean if a field has been set.

### GetConfirmationFromEmail

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationFromEmail() string`

GetConfirmationFromEmail returns the ConfirmationFromEmail field if non-nil, zero value otherwise.

### GetConfirmationFromEmailOk

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationFromEmailOk() (*string, bool)`

GetConfirmationFromEmailOk returns a tuple with the ConfirmationFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationFromEmail

`func (o *PortalConfigurationOpportunitySetup) SetConfirmationFromEmail(v string)`

SetConfirmationFromEmail sets ConfirmationFromEmail field to given value.

### HasConfirmationFromEmail

`func (o *PortalConfigurationOpportunitySetup) HasConfirmationFromEmail() bool`

HasConfirmationFromEmail returns a boolean if a field has been set.

### GetConfirmationEmailSubject

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationEmailSubject() string`

GetConfirmationEmailSubject returns the ConfirmationEmailSubject field if non-nil, zero value otherwise.

### GetConfirmationEmailSubjectOk

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationEmailSubjectOk() (*string, bool)`

GetConfirmationEmailSubjectOk returns a tuple with the ConfirmationEmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationEmailSubject

`func (o *PortalConfigurationOpportunitySetup) SetConfirmationEmailSubject(v string)`

SetConfirmationEmailSubject sets ConfirmationEmailSubject field to given value.

### HasConfirmationEmailSubject

`func (o *PortalConfigurationOpportunitySetup) HasConfirmationEmailSubject() bool`

HasConfirmationEmailSubject returns a boolean if a field has been set.

### GetConfirmationEmailBody

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationEmailBody() string`

GetConfirmationEmailBody returns the ConfirmationEmailBody field if non-nil, zero value otherwise.

### GetConfirmationEmailBodyOk

`func (o *PortalConfigurationOpportunitySetup) GetConfirmationEmailBodyOk() (*string, bool)`

GetConfirmationEmailBodyOk returns a tuple with the ConfirmationEmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationEmailBody

`func (o *PortalConfigurationOpportunitySetup) SetConfirmationEmailBody(v string)`

SetConfirmationEmailBody sets ConfirmationEmailBody field to given value.

### HasConfirmationEmailBody

`func (o *PortalConfigurationOpportunitySetup) HasConfirmationEmailBody() bool`

HasConfirmationEmailBody returns a boolean if a field has been set.

### GetInfo

`func (o *PortalConfigurationOpportunitySetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PortalConfigurationOpportunitySetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PortalConfigurationOpportunitySetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PortalConfigurationOpportunitySetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


