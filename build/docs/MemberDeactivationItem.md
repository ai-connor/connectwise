# MemberDeactivationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**ReAssignToMember** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 

## Methods

### NewMemberDeactivationItem

`func NewMemberDeactivationItem() *MemberDeactivationItem`

NewMemberDeactivationItem instantiates a new MemberDeactivationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDeactivationItemWithDefaults

`func NewMemberDeactivationItemWithDefaults() *MemberDeactivationItem`

NewMemberDeactivationItemWithDefaults instantiates a new MemberDeactivationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *MemberDeactivationItem) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MemberDeactivationItem) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MemberDeactivationItem) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *MemberDeactivationItem) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetReAssignToMember

`func (o *MemberDeactivationItem) GetReAssignToMember() MemberReference`

GetReAssignToMember returns the ReAssignToMember field if non-nil, zero value otherwise.

### GetReAssignToMemberOk

`func (o *MemberDeactivationItem) GetReAssignToMemberOk() (*MemberReference, bool)`

GetReAssignToMemberOk returns a tuple with the ReAssignToMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReAssignToMember

`func (o *MemberDeactivationItem) SetReAssignToMember(v MemberReference)`

SetReAssignToMember sets ReAssignToMember field to given value.

### HasReAssignToMember

`func (o *MemberDeactivationItem) HasReAssignToMember() bool`

HasReAssignToMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


