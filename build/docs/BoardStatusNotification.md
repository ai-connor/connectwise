# BoardStatusNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NotifyWho** | [**NotificationRecipientReference**](NotificationRecipientReference.md) |  | 
**Status** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Email** | Pointer to **string** | Service Status Notification email must be entered if the notify type is \&quot;Email Address\&quot;. Max length: 255; | [optional] 
**WorkflowStep** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardStatusNotification

`func NewBoardStatusNotification(notifyWho NotificationRecipientReference, ) *BoardStatusNotification`

NewBoardStatusNotification instantiates a new BoardStatusNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardStatusNotificationWithDefaults

`func NewBoardStatusNotificationWithDefaults() *BoardStatusNotification`

NewBoardStatusNotificationWithDefaults instantiates a new BoardStatusNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardStatusNotification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardStatusNotification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardStatusNotification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardStatusNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotifyWho

`func (o *BoardStatusNotification) GetNotifyWho() NotificationRecipientReference`

GetNotifyWho returns the NotifyWho field if non-nil, zero value otherwise.

### GetNotifyWhoOk

`func (o *BoardStatusNotification) GetNotifyWhoOk() (*NotificationRecipientReference, bool)`

GetNotifyWhoOk returns a tuple with the NotifyWho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWho

`func (o *BoardStatusNotification) SetNotifyWho(v NotificationRecipientReference)`

SetNotifyWho sets NotifyWho field to given value.


### GetStatus

`func (o *BoardStatusNotification) GetStatus() ServiceStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BoardStatusNotification) GetStatusOk() (*ServiceStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BoardStatusNotification) SetStatus(v ServiceStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BoardStatusNotification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMember

`func (o *BoardStatusNotification) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *BoardStatusNotification) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *BoardStatusNotification) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *BoardStatusNotification) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetEmail

`func (o *BoardStatusNotification) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BoardStatusNotification) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BoardStatusNotification) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BoardStatusNotification) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetWorkflowStep

`func (o *BoardStatusNotification) GetWorkflowStep() int32`

GetWorkflowStep returns the WorkflowStep field if non-nil, zero value otherwise.

### GetWorkflowStepOk

`func (o *BoardStatusNotification) GetWorkflowStepOk() (*int32, bool)`

GetWorkflowStepOk returns a tuple with the WorkflowStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStep

`func (o *BoardStatusNotification) SetWorkflowStep(v int32)`

SetWorkflowStep sets WorkflowStep field to given value.

### HasWorkflowStep

`func (o *BoardStatusNotification) HasWorkflowStep() bool`

HasWorkflowStep returns a boolean if a field has been set.

### SetWorkflowStepNil

`func (o *BoardStatusNotification) SetWorkflowStepNil(b bool)`

 SetWorkflowStepNil sets the value for WorkflowStep to be an explicit nil

### UnsetWorkflowStep
`func (o *BoardStatusNotification) UnsetWorkflowStep()`

UnsetWorkflowStep ensures that no value is present for WorkflowStep, not even an explicit nil
### GetInfo

`func (o *BoardStatusNotification) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardStatusNotification) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardStatusNotification) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardStatusNotification) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


