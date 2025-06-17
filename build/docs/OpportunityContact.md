# OpportunityContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Contact** | [**ContactReference**](ContactReference.md) |  | 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Role** | Pointer to [**OpportunitySalesRoleReference**](OpportunitySalesRoleReference.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ReferralFlag** | Pointer to **NullableBool** |  | [optional] 
**OpportunityId** | Pointer to **NullableInt32** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOpportunityContact

`func NewOpportunityContact(contact ContactReference, ) *OpportunityContact`

NewOpportunityContact instantiates a new OpportunityContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityContactWithDefaults

`func NewOpportunityContactWithDefaults() *OpportunityContact`

NewOpportunityContactWithDefaults instantiates a new OpportunityContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpportunityContact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpportunityContact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpportunityContact) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpportunityContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContact

`func (o *OpportunityContact) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OpportunityContact) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OpportunityContact) SetContact(v ContactReference)`

SetContact sets Contact field to given value.


### GetCompany

`func (o *OpportunityContact) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *OpportunityContact) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *OpportunityContact) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *OpportunityContact) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetRole

`func (o *OpportunityContact) GetRole() OpportunitySalesRoleReference`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OpportunityContact) GetRoleOk() (*OpportunitySalesRoleReference, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OpportunityContact) SetRole(v OpportunitySalesRoleReference)`

SetRole sets Role field to given value.

### HasRole

`func (o *OpportunityContact) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetNotes

`func (o *OpportunityContact) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OpportunityContact) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OpportunityContact) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OpportunityContact) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetReferralFlag

`func (o *OpportunityContact) GetReferralFlag() bool`

GetReferralFlag returns the ReferralFlag field if non-nil, zero value otherwise.

### GetReferralFlagOk

`func (o *OpportunityContact) GetReferralFlagOk() (*bool, bool)`

GetReferralFlagOk returns a tuple with the ReferralFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralFlag

`func (o *OpportunityContact) SetReferralFlag(v bool)`

SetReferralFlag sets ReferralFlag field to given value.

### HasReferralFlag

`func (o *OpportunityContact) HasReferralFlag() bool`

HasReferralFlag returns a boolean if a field has been set.

### SetReferralFlagNil

`func (o *OpportunityContact) SetReferralFlagNil(b bool)`

 SetReferralFlagNil sets the value for ReferralFlag to be an explicit nil

### UnsetReferralFlag
`func (o *OpportunityContact) UnsetReferralFlag()`

UnsetReferralFlag ensures that no value is present for ReferralFlag, not even an explicit nil
### GetOpportunityId

`func (o *OpportunityContact) GetOpportunityId() int32`

GetOpportunityId returns the OpportunityId field if non-nil, zero value otherwise.

### GetOpportunityIdOk

`func (o *OpportunityContact) GetOpportunityIdOk() (*int32, bool)`

GetOpportunityIdOk returns a tuple with the OpportunityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityId

`func (o *OpportunityContact) SetOpportunityId(v int32)`

SetOpportunityId sets OpportunityId field to given value.

### HasOpportunityId

`func (o *OpportunityContact) HasOpportunityId() bool`

HasOpportunityId returns a boolean if a field has been set.

### SetOpportunityIdNil

`func (o *OpportunityContact) SetOpportunityIdNil(b bool)`

 SetOpportunityIdNil sets the value for OpportunityId to be an explicit nil

### UnsetOpportunityId
`func (o *OpportunityContact) UnsetOpportunityId()`

UnsetOpportunityId ensures that no value is present for OpportunityId, not even an explicit nil
### GetPhoneNumber

`func (o *OpportunityContact) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OpportunityContact) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OpportunityContact) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OpportunityContact) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmailAddress

`func (o *OpportunityContact) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *OpportunityContact) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *OpportunityContact) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *OpportunityContact) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetInfo

`func (o *OpportunityContact) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OpportunityContact) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OpportunityContact) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OpportunityContact) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


