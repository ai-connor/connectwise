# RmaStatusNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NotifyWho** | [**NotificationRecipientReference**](NotificationRecipientReference.md) |  | 
**Status** | Pointer to [**RmaStatusReference**](RmaStatusReference.md) |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Email** | Pointer to **string** | RMA Status Notification sendEmail must be entered if the notify type is \&quot;Email Address\&quot;. Max length: 50; | [optional] 
**WorkflowStep** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewRmaStatusNotification

`func NewRmaStatusNotification(notifyWho NotificationRecipientReference, ) *RmaStatusNotification`

NewRmaStatusNotification instantiates a new RmaStatusNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRmaStatusNotificationWithDefaults

`func NewRmaStatusNotificationWithDefaults() *RmaStatusNotification`

NewRmaStatusNotificationWithDefaults instantiates a new RmaStatusNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RmaStatusNotification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RmaStatusNotification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RmaStatusNotification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RmaStatusNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotifyWho

`func (o *RmaStatusNotification) GetNotifyWho() NotificationRecipientReference`

GetNotifyWho returns the NotifyWho field if non-nil, zero value otherwise.

### GetNotifyWhoOk

`func (o *RmaStatusNotification) GetNotifyWhoOk() (*NotificationRecipientReference, bool)`

GetNotifyWhoOk returns a tuple with the NotifyWho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWho

`func (o *RmaStatusNotification) SetNotifyWho(v NotificationRecipientReference)`

SetNotifyWho sets NotifyWho field to given value.


### GetStatus

`func (o *RmaStatusNotification) GetStatus() RmaStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RmaStatusNotification) GetStatusOk() (*RmaStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RmaStatusNotification) SetStatus(v RmaStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RmaStatusNotification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMember

`func (o *RmaStatusNotification) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *RmaStatusNotification) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *RmaStatusNotification) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *RmaStatusNotification) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetEmail

`func (o *RmaStatusNotification) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RmaStatusNotification) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RmaStatusNotification) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RmaStatusNotification) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetWorkflowStep

`func (o *RmaStatusNotification) GetWorkflowStep() int32`

GetWorkflowStep returns the WorkflowStep field if non-nil, zero value otherwise.

### GetWorkflowStepOk

`func (o *RmaStatusNotification) GetWorkflowStepOk() (*int32, bool)`

GetWorkflowStepOk returns a tuple with the WorkflowStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStep

`func (o *RmaStatusNotification) SetWorkflowStep(v int32)`

SetWorkflowStep sets WorkflowStep field to given value.

### HasWorkflowStep

`func (o *RmaStatusNotification) HasWorkflowStep() bool`

HasWorkflowStep returns a boolean if a field has been set.

### SetWorkflowStepNil

`func (o *RmaStatusNotification) SetWorkflowStepNil(b bool)`

 SetWorkflowStepNil sets the value for WorkflowStep to be an explicit nil

### UnsetWorkflowStep
`func (o *RmaStatusNotification) UnsetWorkflowStep()`

UnsetWorkflowStep ensures that no value is present for WorkflowStep, not even an explicit nil
### GetInfo

`func (o *RmaStatusNotification) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *RmaStatusNotification) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *RmaStatusNotification) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *RmaStatusNotification) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


