# MemberDeactivationStatusWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**ReAssignToMember** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 

## Methods

### NewMemberDeactivationStatusWorkflow

`func NewMemberDeactivationStatusWorkflow() *MemberDeactivationStatusWorkflow`

NewMemberDeactivationStatusWorkflow instantiates a new MemberDeactivationStatusWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDeactivationStatusWorkflowWithDefaults

`func NewMemberDeactivationStatusWorkflowWithDefaults() *MemberDeactivationStatusWorkflow`

NewMemberDeactivationStatusWorkflowWithDefaults instantiates a new MemberDeactivationStatusWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberDeactivationStatusWorkflow) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberDeactivationStatusWorkflow) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberDeactivationStatusWorkflow) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MemberDeactivationStatusWorkflow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MemberDeactivationStatusWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberDeactivationStatusWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberDeactivationStatusWorkflow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemberDeactivationStatusWorkflow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCount

`func (o *MemberDeactivationStatusWorkflow) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MemberDeactivationStatusWorkflow) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MemberDeactivationStatusWorkflow) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *MemberDeactivationStatusWorkflow) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetReAssignToMember

`func (o *MemberDeactivationStatusWorkflow) GetReAssignToMember() MemberReference`

GetReAssignToMember returns the ReAssignToMember field if non-nil, zero value otherwise.

### GetReAssignToMemberOk

`func (o *MemberDeactivationStatusWorkflow) GetReAssignToMemberOk() (*MemberReference, bool)`

GetReAssignToMemberOk returns a tuple with the ReAssignToMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReAssignToMember

`func (o *MemberDeactivationStatusWorkflow) SetReAssignToMember(v MemberReference)`

SetReAssignToMember sets ReAssignToMember field to given value.

### HasReAssignToMember

`func (o *MemberDeactivationStatusWorkflow) HasReAssignToMember() bool`

HasReAssignToMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


