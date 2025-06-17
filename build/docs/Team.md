# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | **NullableString** |  | 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**SalesTeam** | Pointer to [**SalesTeamReference**](SalesTeamReference.md) |  | [optional] 
**CommissionPercent** | Pointer to **NullableInt32** |  | [optional] 
**ReferralFlag** | Pointer to **NullableBool** |  | [optional] 
**OpportunityId** | Pointer to **NullableInt32** |  | [optional] 
**ResponsibleFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTeam

`func NewTeam(type_ NullableString, ) *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Team) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Team) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Team) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Team) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Team) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Team) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Team) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *Team) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Team) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMember

`func (o *Team) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *Team) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *Team) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *Team) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetSalesTeam

`func (o *Team) GetSalesTeam() SalesTeamReference`

GetSalesTeam returns the SalesTeam field if non-nil, zero value otherwise.

### GetSalesTeamOk

`func (o *Team) GetSalesTeamOk() (*SalesTeamReference, bool)`

GetSalesTeamOk returns a tuple with the SalesTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTeam

`func (o *Team) SetSalesTeam(v SalesTeamReference)`

SetSalesTeam sets SalesTeam field to given value.

### HasSalesTeam

`func (o *Team) HasSalesTeam() bool`

HasSalesTeam returns a boolean if a field has been set.

### GetCommissionPercent

`func (o *Team) GetCommissionPercent() int32`

GetCommissionPercent returns the CommissionPercent field if non-nil, zero value otherwise.

### GetCommissionPercentOk

`func (o *Team) GetCommissionPercentOk() (*int32, bool)`

GetCommissionPercentOk returns a tuple with the CommissionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionPercent

`func (o *Team) SetCommissionPercent(v int32)`

SetCommissionPercent sets CommissionPercent field to given value.

### HasCommissionPercent

`func (o *Team) HasCommissionPercent() bool`

HasCommissionPercent returns a boolean if a field has been set.

### SetCommissionPercentNil

`func (o *Team) SetCommissionPercentNil(b bool)`

 SetCommissionPercentNil sets the value for CommissionPercent to be an explicit nil

### UnsetCommissionPercent
`func (o *Team) UnsetCommissionPercent()`

UnsetCommissionPercent ensures that no value is present for CommissionPercent, not even an explicit nil
### GetReferralFlag

`func (o *Team) GetReferralFlag() bool`

GetReferralFlag returns the ReferralFlag field if non-nil, zero value otherwise.

### GetReferralFlagOk

`func (o *Team) GetReferralFlagOk() (*bool, bool)`

GetReferralFlagOk returns a tuple with the ReferralFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralFlag

`func (o *Team) SetReferralFlag(v bool)`

SetReferralFlag sets ReferralFlag field to given value.

### HasReferralFlag

`func (o *Team) HasReferralFlag() bool`

HasReferralFlag returns a boolean if a field has been set.

### SetReferralFlagNil

`func (o *Team) SetReferralFlagNil(b bool)`

 SetReferralFlagNil sets the value for ReferralFlag to be an explicit nil

### UnsetReferralFlag
`func (o *Team) UnsetReferralFlag()`

UnsetReferralFlag ensures that no value is present for ReferralFlag, not even an explicit nil
### GetOpportunityId

`func (o *Team) GetOpportunityId() int32`

GetOpportunityId returns the OpportunityId field if non-nil, zero value otherwise.

### GetOpportunityIdOk

`func (o *Team) GetOpportunityIdOk() (*int32, bool)`

GetOpportunityIdOk returns a tuple with the OpportunityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityId

`func (o *Team) SetOpportunityId(v int32)`

SetOpportunityId sets OpportunityId field to given value.

### HasOpportunityId

`func (o *Team) HasOpportunityId() bool`

HasOpportunityId returns a boolean if a field has been set.

### SetOpportunityIdNil

`func (o *Team) SetOpportunityIdNil(b bool)`

 SetOpportunityIdNil sets the value for OpportunityId to be an explicit nil

### UnsetOpportunityId
`func (o *Team) UnsetOpportunityId()`

UnsetOpportunityId ensures that no value is present for OpportunityId, not even an explicit nil
### GetResponsibleFlag

`func (o *Team) GetResponsibleFlag() bool`

GetResponsibleFlag returns the ResponsibleFlag field if non-nil, zero value otherwise.

### GetResponsibleFlagOk

`func (o *Team) GetResponsibleFlagOk() (*bool, bool)`

GetResponsibleFlagOk returns a tuple with the ResponsibleFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleFlag

`func (o *Team) SetResponsibleFlag(v bool)`

SetResponsibleFlag sets ResponsibleFlag field to given value.

### HasResponsibleFlag

`func (o *Team) HasResponsibleFlag() bool`

HasResponsibleFlag returns a boolean if a field has been set.

### SetResponsibleFlagNil

`func (o *Team) SetResponsibleFlagNil(b bool)`

 SetResponsibleFlagNil sets the value for ResponsibleFlag to be an explicit nil

### UnsetResponsibleFlag
`func (o *Team) UnsetResponsibleFlag()`

UnsetResponsibleFlag ensures that no value is present for ResponsibleFlag, not even an explicit nil
### GetInfo

`func (o *Team) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Team) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Team) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Team) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


