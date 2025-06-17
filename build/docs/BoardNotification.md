# BoardNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NotifyWho** | [**NotificationRecipientReference**](NotificationRecipientReference.md) |  | 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Email** | Pointer to **string** |  Max length: 50; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardNotification

`func NewBoardNotification(notifyWho NotificationRecipientReference, ) *BoardNotification`

NewBoardNotification instantiates a new BoardNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardNotificationWithDefaults

`func NewBoardNotificationWithDefaults() *BoardNotification`

NewBoardNotificationWithDefaults instantiates a new BoardNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardNotification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardNotification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardNotification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotifyWho

`func (o *BoardNotification) GetNotifyWho() NotificationRecipientReference`

GetNotifyWho returns the NotifyWho field if non-nil, zero value otherwise.

### GetNotifyWhoOk

`func (o *BoardNotification) GetNotifyWhoOk() (*NotificationRecipientReference, bool)`

GetNotifyWhoOk returns a tuple with the NotifyWho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWho

`func (o *BoardNotification) SetNotifyWho(v NotificationRecipientReference)`

SetNotifyWho sets NotifyWho field to given value.


### GetMember

`func (o *BoardNotification) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *BoardNotification) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *BoardNotification) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *BoardNotification) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetEmail

`func (o *BoardNotification) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BoardNotification) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BoardNotification) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BoardNotification) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetInfo

`func (o *BoardNotification) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardNotification) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardNotification) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardNotification) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


