# MemberCertification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Certification** | [**CertificationReference**](CertificationReference.md) |  | 
**PercentComplete** | Pointer to **NullableInt32** |  | [optional] 
**DateReceived** | Pointer to **time.Time** |  | [optional] 
**DateExpires** | Pointer to **time.Time** |  | [optional] 
**CertificationNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMemberCertification

`func NewMemberCertification(certification CertificationReference, ) *MemberCertification`

NewMemberCertification instantiates a new MemberCertification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberCertificationWithDefaults

`func NewMemberCertificationWithDefaults() *MemberCertification`

NewMemberCertificationWithDefaults instantiates a new MemberCertification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberCertification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberCertification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberCertification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MemberCertification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCertification

`func (o *MemberCertification) GetCertification() CertificationReference`

GetCertification returns the Certification field if non-nil, zero value otherwise.

### GetCertificationOk

`func (o *MemberCertification) GetCertificationOk() (*CertificationReference, bool)`

GetCertificationOk returns a tuple with the Certification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertification

`func (o *MemberCertification) SetCertification(v CertificationReference)`

SetCertification sets Certification field to given value.


### GetPercentComplete

`func (o *MemberCertification) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *MemberCertification) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *MemberCertification) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *MemberCertification) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentCompleteNil

`func (o *MemberCertification) SetPercentCompleteNil(b bool)`

 SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil

### UnsetPercentComplete
`func (o *MemberCertification) UnsetPercentComplete()`

UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
### GetDateReceived

`func (o *MemberCertification) GetDateReceived() time.Time`

GetDateReceived returns the DateReceived field if non-nil, zero value otherwise.

### GetDateReceivedOk

`func (o *MemberCertification) GetDateReceivedOk() (*time.Time, bool)`

GetDateReceivedOk returns a tuple with the DateReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateReceived

`func (o *MemberCertification) SetDateReceived(v time.Time)`

SetDateReceived sets DateReceived field to given value.

### HasDateReceived

`func (o *MemberCertification) HasDateReceived() bool`

HasDateReceived returns a boolean if a field has been set.

### GetDateExpires

`func (o *MemberCertification) GetDateExpires() time.Time`

GetDateExpires returns the DateExpires field if non-nil, zero value otherwise.

### GetDateExpiresOk

`func (o *MemberCertification) GetDateExpiresOk() (*time.Time, bool)`

GetDateExpiresOk returns a tuple with the DateExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateExpires

`func (o *MemberCertification) SetDateExpires(v time.Time)`

SetDateExpires sets DateExpires field to given value.

### HasDateExpires

`func (o *MemberCertification) HasDateExpires() bool`

HasDateExpires returns a boolean if a field has been set.

### GetCertificationNumber

`func (o *MemberCertification) GetCertificationNumber() string`

GetCertificationNumber returns the CertificationNumber field if non-nil, zero value otherwise.

### GetCertificationNumberOk

`func (o *MemberCertification) GetCertificationNumberOk() (*string, bool)`

GetCertificationNumberOk returns a tuple with the CertificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationNumber

`func (o *MemberCertification) SetCertificationNumber(v string)`

SetCertificationNumber sets CertificationNumber field to given value.

### HasCertificationNumber

`func (o *MemberCertification) HasCertificationNumber() bool`

HasCertificationNumber returns a boolean if a field has been set.

### GetNotes

`func (o *MemberCertification) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MemberCertification) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MemberCertification) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MemberCertification) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetMember

`func (o *MemberCertification) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *MemberCertification) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *MemberCertification) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *MemberCertification) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetCompany

`func (o *MemberCertification) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *MemberCertification) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *MemberCertification) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *MemberCertification) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetInfo

`func (o *MemberCertification) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MemberCertification) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MemberCertification) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MemberCertification) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


