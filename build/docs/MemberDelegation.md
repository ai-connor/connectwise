# MemberDelegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DelegationType** | **NullableString** |  | 
**DelegatedTo** | [**MemberReference**](MemberReference.md) |  | 
**DateStart** | **time.Time** |  | 
**DateEnd** | **time.Time** |  | 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMemberDelegation

`func NewMemberDelegation(delegationType NullableString, delegatedTo MemberReference, dateStart time.Time, dateEnd time.Time, ) *MemberDelegation`

NewMemberDelegation instantiates a new MemberDelegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDelegationWithDefaults

`func NewMemberDelegationWithDefaults() *MemberDelegation`

NewMemberDelegationWithDefaults instantiates a new MemberDelegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberDelegation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberDelegation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberDelegation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MemberDelegation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDelegationType

`func (o *MemberDelegation) GetDelegationType() string`

GetDelegationType returns the DelegationType field if non-nil, zero value otherwise.

### GetDelegationTypeOk

`func (o *MemberDelegation) GetDelegationTypeOk() (*string, bool)`

GetDelegationTypeOk returns a tuple with the DelegationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationType

`func (o *MemberDelegation) SetDelegationType(v string)`

SetDelegationType sets DelegationType field to given value.


### SetDelegationTypeNil

`func (o *MemberDelegation) SetDelegationTypeNil(b bool)`

 SetDelegationTypeNil sets the value for DelegationType to be an explicit nil

### UnsetDelegationType
`func (o *MemberDelegation) UnsetDelegationType()`

UnsetDelegationType ensures that no value is present for DelegationType, not even an explicit nil
### GetDelegatedTo

`func (o *MemberDelegation) GetDelegatedTo() MemberReference`

GetDelegatedTo returns the DelegatedTo field if non-nil, zero value otherwise.

### GetDelegatedToOk

`func (o *MemberDelegation) GetDelegatedToOk() (*MemberReference, bool)`

GetDelegatedToOk returns a tuple with the DelegatedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedTo

`func (o *MemberDelegation) SetDelegatedTo(v MemberReference)`

SetDelegatedTo sets DelegatedTo field to given value.


### GetDateStart

`func (o *MemberDelegation) GetDateStart() time.Time`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *MemberDelegation) GetDateStartOk() (*time.Time, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *MemberDelegation) SetDateStart(v time.Time)`

SetDateStart sets DateStart field to given value.


### GetDateEnd

`func (o *MemberDelegation) GetDateEnd() time.Time`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *MemberDelegation) GetDateEndOk() (*time.Time, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *MemberDelegation) SetDateEnd(v time.Time)`

SetDateEnd sets DateEnd field to given value.


### GetMember

`func (o *MemberDelegation) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *MemberDelegation) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *MemberDelegation) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *MemberDelegation) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetInfo

`func (o *MemberDelegation) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MemberDelegation) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MemberDelegation) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MemberDelegation) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


