# BoardAutoAssignResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NotifyWho** | [**NotificationRecipientReference**](NotificationRecipientReference.md) |  | 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardAutoAssignResource

`func NewBoardAutoAssignResource(notifyWho NotificationRecipientReference, ) *BoardAutoAssignResource`

NewBoardAutoAssignResource instantiates a new BoardAutoAssignResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardAutoAssignResourceWithDefaults

`func NewBoardAutoAssignResourceWithDefaults() *BoardAutoAssignResource`

NewBoardAutoAssignResourceWithDefaults instantiates a new BoardAutoAssignResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardAutoAssignResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardAutoAssignResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardAutoAssignResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardAutoAssignResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotifyWho

`func (o *BoardAutoAssignResource) GetNotifyWho() NotificationRecipientReference`

GetNotifyWho returns the NotifyWho field if non-nil, zero value otherwise.

### GetNotifyWhoOk

`func (o *BoardAutoAssignResource) GetNotifyWhoOk() (*NotificationRecipientReference, bool)`

GetNotifyWhoOk returns a tuple with the NotifyWho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWho

`func (o *BoardAutoAssignResource) SetNotifyWho(v NotificationRecipientReference)`

SetNotifyWho sets NotifyWho field to given value.


### GetMember

`func (o *BoardAutoAssignResource) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *BoardAutoAssignResource) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *BoardAutoAssignResource) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *BoardAutoAssignResource) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetInfo

`func (o *BoardAutoAssignResource) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardAutoAssignResource) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardAutoAssignResource) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardAutoAssignResource) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


