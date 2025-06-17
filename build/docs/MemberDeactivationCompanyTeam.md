# MemberDeactivationCompanyTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReAssignToContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**ReAssignToMember** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 

## Methods

### NewMemberDeactivationCompanyTeam

`func NewMemberDeactivationCompanyTeam() *MemberDeactivationCompanyTeam`

NewMemberDeactivationCompanyTeam instantiates a new MemberDeactivationCompanyTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDeactivationCompanyTeamWithDefaults

`func NewMemberDeactivationCompanyTeamWithDefaults() *MemberDeactivationCompanyTeam`

NewMemberDeactivationCompanyTeamWithDefaults instantiates a new MemberDeactivationCompanyTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberDeactivationCompanyTeam) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberDeactivationCompanyTeam) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberDeactivationCompanyTeam) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MemberDeactivationCompanyTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MemberDeactivationCompanyTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberDeactivationCompanyTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberDeactivationCompanyTeam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemberDeactivationCompanyTeam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReAssignToContact

`func (o *MemberDeactivationCompanyTeam) GetReAssignToContact() ContactReference`

GetReAssignToContact returns the ReAssignToContact field if non-nil, zero value otherwise.

### GetReAssignToContactOk

`func (o *MemberDeactivationCompanyTeam) GetReAssignToContactOk() (*ContactReference, bool)`

GetReAssignToContactOk returns a tuple with the ReAssignToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReAssignToContact

`func (o *MemberDeactivationCompanyTeam) SetReAssignToContact(v ContactReference)`

SetReAssignToContact sets ReAssignToContact field to given value.

### HasReAssignToContact

`func (o *MemberDeactivationCompanyTeam) HasReAssignToContact() bool`

HasReAssignToContact returns a boolean if a field has been set.

### GetCount

`func (o *MemberDeactivationCompanyTeam) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MemberDeactivationCompanyTeam) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MemberDeactivationCompanyTeam) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *MemberDeactivationCompanyTeam) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetReAssignToMember

`func (o *MemberDeactivationCompanyTeam) GetReAssignToMember() MemberReference`

GetReAssignToMember returns the ReAssignToMember field if non-nil, zero value otherwise.

### GetReAssignToMemberOk

`func (o *MemberDeactivationCompanyTeam) GetReAssignToMemberOk() (*MemberReference, bool)`

GetReAssignToMemberOk returns a tuple with the ReAssignToMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReAssignToMember

`func (o *MemberDeactivationCompanyTeam) SetReAssignToMember(v MemberReference)`

SetReAssignToMember sets ReAssignToMember field to given value.

### HasReAssignToMember

`func (o *MemberDeactivationCompanyTeam) HasReAssignToMember() bool`

HasReAssignToMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


