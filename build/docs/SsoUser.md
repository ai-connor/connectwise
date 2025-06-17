# SsoUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SsoUserId** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmailConfirmed** | Pointer to **NullableBool** |  | [optional] 
**DisabledFlag** | Pointer to **NullableBool** |  | [optional] 
**LinkedFlag** | Pointer to **NullableBool** |  | [optional] 
**DateEntered** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**LinkedMember** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 

## Methods

### NewSsoUser

`func NewSsoUser() *SsoUser`

NewSsoUser instantiates a new SsoUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoUserWithDefaults

`func NewSsoUserWithDefaults() *SsoUser`

NewSsoUserWithDefaults instantiates a new SsoUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SsoUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SsoUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SsoUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SsoUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSsoUserId

`func (o *SsoUser) GetSsoUserId() string`

GetSsoUserId returns the SsoUserId field if non-nil, zero value otherwise.

### GetSsoUserIdOk

`func (o *SsoUser) GetSsoUserIdOk() (*string, bool)`

GetSsoUserIdOk returns a tuple with the SsoUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUserId

`func (o *SsoUser) SetSsoUserId(v string)`

SetSsoUserId sets SsoUserId field to given value.

### HasSsoUserId

`func (o *SsoUser) HasSsoUserId() bool`

HasSsoUserId returns a boolean if a field has been set.

### GetUserName

`func (o *SsoUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SsoUser) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SsoUser) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SsoUser) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetFirstName

`func (o *SsoUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SsoUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SsoUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SsoUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *SsoUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SsoUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SsoUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SsoUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *SsoUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SsoUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SsoUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SsoUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailConfirmed

`func (o *SsoUser) GetEmailConfirmed() bool`

GetEmailConfirmed returns the EmailConfirmed field if non-nil, zero value otherwise.

### GetEmailConfirmedOk

`func (o *SsoUser) GetEmailConfirmedOk() (*bool, bool)`

GetEmailConfirmedOk returns a tuple with the EmailConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfirmed

`func (o *SsoUser) SetEmailConfirmed(v bool)`

SetEmailConfirmed sets EmailConfirmed field to given value.

### HasEmailConfirmed

`func (o *SsoUser) HasEmailConfirmed() bool`

HasEmailConfirmed returns a boolean if a field has been set.

### SetEmailConfirmedNil

`func (o *SsoUser) SetEmailConfirmedNil(b bool)`

 SetEmailConfirmedNil sets the value for EmailConfirmed to be an explicit nil

### UnsetEmailConfirmed
`func (o *SsoUser) UnsetEmailConfirmed()`

UnsetEmailConfirmed ensures that no value is present for EmailConfirmed, not even an explicit nil
### GetDisabledFlag

`func (o *SsoUser) GetDisabledFlag() bool`

GetDisabledFlag returns the DisabledFlag field if non-nil, zero value otherwise.

### GetDisabledFlagOk

`func (o *SsoUser) GetDisabledFlagOk() (*bool, bool)`

GetDisabledFlagOk returns a tuple with the DisabledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledFlag

`func (o *SsoUser) SetDisabledFlag(v bool)`

SetDisabledFlag sets DisabledFlag field to given value.

### HasDisabledFlag

`func (o *SsoUser) HasDisabledFlag() bool`

HasDisabledFlag returns a boolean if a field has been set.

### SetDisabledFlagNil

`func (o *SsoUser) SetDisabledFlagNil(b bool)`

 SetDisabledFlagNil sets the value for DisabledFlag to be an explicit nil

### UnsetDisabledFlag
`func (o *SsoUser) UnsetDisabledFlag()`

UnsetDisabledFlag ensures that no value is present for DisabledFlag, not even an explicit nil
### GetLinkedFlag

`func (o *SsoUser) GetLinkedFlag() bool`

GetLinkedFlag returns the LinkedFlag field if non-nil, zero value otherwise.

### GetLinkedFlagOk

`func (o *SsoUser) GetLinkedFlagOk() (*bool, bool)`

GetLinkedFlagOk returns a tuple with the LinkedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFlag

`func (o *SsoUser) SetLinkedFlag(v bool)`

SetLinkedFlag sets LinkedFlag field to given value.

### HasLinkedFlag

`func (o *SsoUser) HasLinkedFlag() bool`

HasLinkedFlag returns a boolean if a field has been set.

### SetLinkedFlagNil

`func (o *SsoUser) SetLinkedFlagNil(b bool)`

 SetLinkedFlagNil sets the value for LinkedFlag to be an explicit nil

### UnsetLinkedFlag
`func (o *SsoUser) UnsetLinkedFlag()`

UnsetLinkedFlag ensures that no value is present for LinkedFlag, not even an explicit nil
### GetDateEntered

`func (o *SsoUser) GetDateEntered() string`

GetDateEntered returns the DateEntered field if non-nil, zero value otherwise.

### GetDateEnteredOk

`func (o *SsoUser) GetDateEnteredOk() (*string, bool)`

GetDateEnteredOk returns a tuple with the DateEntered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEntered

`func (o *SsoUser) SetDateEntered(v string)`

SetDateEntered sets DateEntered field to given value.

### HasDateEntered

`func (o *SsoUser) HasDateEntered() bool`

HasDateEntered returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SsoUser) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SsoUser) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SsoUser) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SsoUser) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinkedMember

`func (o *SsoUser) GetLinkedMember() MemberReference`

GetLinkedMember returns the LinkedMember field if non-nil, zero value otherwise.

### GetLinkedMemberOk

`func (o *SsoUser) GetLinkedMemberOk() (*MemberReference, bool)`

GetLinkedMemberOk returns a tuple with the LinkedMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedMember

`func (o *SsoUser) SetLinkedMember(v MemberReference)`

SetLinkedMember sets LinkedMember field to given value.

### HasLinkedMember

`func (o *SsoUser) HasLinkedMember() bool`

HasLinkedMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


