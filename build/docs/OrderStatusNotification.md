# OrderStatusNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NotifyWho** | [**NotificationRecipientReference**](NotificationRecipientReference.md) |  | 
**Status** | Pointer to [**OrderStatusReference**](OrderStatusReference.md) |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Email** | Pointer to **string** | Order Status Notification sendEmail must be entered if the notify type is \&quot;Email Address\&quot;. Max length: 50; | [optional] 
**WorkflowStep** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOrderStatusNotification

`func NewOrderStatusNotification(notifyWho NotificationRecipientReference, ) *OrderStatusNotification`

NewOrderStatusNotification instantiates a new OrderStatusNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusNotificationWithDefaults

`func NewOrderStatusNotificationWithDefaults() *OrderStatusNotification`

NewOrderStatusNotificationWithDefaults instantiates a new OrderStatusNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderStatusNotification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderStatusNotification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderStatusNotification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrderStatusNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotifyWho

`func (o *OrderStatusNotification) GetNotifyWho() NotificationRecipientReference`

GetNotifyWho returns the NotifyWho field if non-nil, zero value otherwise.

### GetNotifyWhoOk

`func (o *OrderStatusNotification) GetNotifyWhoOk() (*NotificationRecipientReference, bool)`

GetNotifyWhoOk returns a tuple with the NotifyWho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWho

`func (o *OrderStatusNotification) SetNotifyWho(v NotificationRecipientReference)`

SetNotifyWho sets NotifyWho field to given value.


### GetStatus

`func (o *OrderStatusNotification) GetStatus() OrderStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderStatusNotification) GetStatusOk() (*OrderStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderStatusNotification) SetStatus(v OrderStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderStatusNotification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMember

`func (o *OrderStatusNotification) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *OrderStatusNotification) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *OrderStatusNotification) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *OrderStatusNotification) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetEmail

`func (o *OrderStatusNotification) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderStatusNotification) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderStatusNotification) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderStatusNotification) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetWorkflowStep

`func (o *OrderStatusNotification) GetWorkflowStep() int32`

GetWorkflowStep returns the WorkflowStep field if non-nil, zero value otherwise.

### GetWorkflowStepOk

`func (o *OrderStatusNotification) GetWorkflowStepOk() (*int32, bool)`

GetWorkflowStepOk returns a tuple with the WorkflowStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStep

`func (o *OrderStatusNotification) SetWorkflowStep(v int32)`

SetWorkflowStep sets WorkflowStep field to given value.

### HasWorkflowStep

`func (o *OrderStatusNotification) HasWorkflowStep() bool`

HasWorkflowStep returns a boolean if a field has been set.

### SetWorkflowStepNil

`func (o *OrderStatusNotification) SetWorkflowStepNil(b bool)`

 SetWorkflowStepNil sets the value for WorkflowStep to be an explicit nil

### UnsetWorkflowStep
`func (o *OrderStatusNotification) UnsetWorkflowStep()`

UnsetWorkflowStep ensures that no value is present for WorkflowStep, not even an explicit nil
### GetInfo

`func (o *OrderStatusNotification) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OrderStatusNotification) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OrderStatusNotification) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OrderStatusNotification) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


