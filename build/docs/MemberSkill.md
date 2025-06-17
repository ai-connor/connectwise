# MemberSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Skill** | [**SkillReference**](SkillReference.md) |  | 
**SkillLevel** | **NullableString** |  | 
**CertifiedFlag** | Pointer to **NullableBool** |  | [optional] 
**YearsExperience** | Pointer to **NullableInt32** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMemberSkill

`func NewMemberSkill(skill SkillReference, skillLevel NullableString, ) *MemberSkill`

NewMemberSkill instantiates a new MemberSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberSkillWithDefaults

`func NewMemberSkillWithDefaults() *MemberSkill`

NewMemberSkillWithDefaults instantiates a new MemberSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberSkill) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberSkill) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberSkill) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MemberSkill) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSkill

`func (o *MemberSkill) GetSkill() SkillReference`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *MemberSkill) GetSkillOk() (*SkillReference, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *MemberSkill) SetSkill(v SkillReference)`

SetSkill sets Skill field to given value.


### GetSkillLevel

`func (o *MemberSkill) GetSkillLevel() string`

GetSkillLevel returns the SkillLevel field if non-nil, zero value otherwise.

### GetSkillLevelOk

`func (o *MemberSkill) GetSkillLevelOk() (*string, bool)`

GetSkillLevelOk returns a tuple with the SkillLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillLevel

`func (o *MemberSkill) SetSkillLevel(v string)`

SetSkillLevel sets SkillLevel field to given value.


### SetSkillLevelNil

`func (o *MemberSkill) SetSkillLevelNil(b bool)`

 SetSkillLevelNil sets the value for SkillLevel to be an explicit nil

### UnsetSkillLevel
`func (o *MemberSkill) UnsetSkillLevel()`

UnsetSkillLevel ensures that no value is present for SkillLevel, not even an explicit nil
### GetCertifiedFlag

`func (o *MemberSkill) GetCertifiedFlag() bool`

GetCertifiedFlag returns the CertifiedFlag field if non-nil, zero value otherwise.

### GetCertifiedFlagOk

`func (o *MemberSkill) GetCertifiedFlagOk() (*bool, bool)`

GetCertifiedFlagOk returns a tuple with the CertifiedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedFlag

`func (o *MemberSkill) SetCertifiedFlag(v bool)`

SetCertifiedFlag sets CertifiedFlag field to given value.

### HasCertifiedFlag

`func (o *MemberSkill) HasCertifiedFlag() bool`

HasCertifiedFlag returns a boolean if a field has been set.

### SetCertifiedFlagNil

`func (o *MemberSkill) SetCertifiedFlagNil(b bool)`

 SetCertifiedFlagNil sets the value for CertifiedFlag to be an explicit nil

### UnsetCertifiedFlag
`func (o *MemberSkill) UnsetCertifiedFlag()`

UnsetCertifiedFlag ensures that no value is present for CertifiedFlag, not even an explicit nil
### GetYearsExperience

`func (o *MemberSkill) GetYearsExperience() int32`

GetYearsExperience returns the YearsExperience field if non-nil, zero value otherwise.

### GetYearsExperienceOk

`func (o *MemberSkill) GetYearsExperienceOk() (*int32, bool)`

GetYearsExperienceOk returns a tuple with the YearsExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearsExperience

`func (o *MemberSkill) SetYearsExperience(v int32)`

SetYearsExperience sets YearsExperience field to given value.

### HasYearsExperience

`func (o *MemberSkill) HasYearsExperience() bool`

HasYearsExperience returns a boolean if a field has been set.

### SetYearsExperienceNil

`func (o *MemberSkill) SetYearsExperienceNil(b bool)`

 SetYearsExperienceNil sets the value for YearsExperience to be an explicit nil

### UnsetYearsExperience
`func (o *MemberSkill) UnsetYearsExperience()`

UnsetYearsExperience ensures that no value is present for YearsExperience, not even an explicit nil
### GetNotes

`func (o *MemberSkill) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MemberSkill) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MemberSkill) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MemberSkill) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetMember

`func (o *MemberSkill) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *MemberSkill) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *MemberSkill) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *MemberSkill) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetInfo

`func (o *MemberSkill) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MemberSkill) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MemberSkill) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MemberSkill) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


