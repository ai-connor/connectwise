# PurchaseOrderStatusNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NotifyWho** | [**NotificationRecipientReference**](NotificationRecipientReference.md) |  | 
**Status** | Pointer to [**PurchaseOrderStatusReference**](PurchaseOrderStatusReference.md) |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Email** | Pointer to **string** | Purchase Order Status Notification email must be entered if the notify type is \&quot;Email Address\&quot;. Max length: 50; | [optional] 
**WorkflowStep** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPurchaseOrderStatusNotification

`func NewPurchaseOrderStatusNotification(notifyWho NotificationRecipientReference, ) *PurchaseOrderStatusNotification`

NewPurchaseOrderStatusNotification instantiates a new PurchaseOrderStatusNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderStatusNotificationWithDefaults

`func NewPurchaseOrderStatusNotificationWithDefaults() *PurchaseOrderStatusNotification`

NewPurchaseOrderStatusNotificationWithDefaults instantiates a new PurchaseOrderStatusNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PurchaseOrderStatusNotification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseOrderStatusNotification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseOrderStatusNotification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseOrderStatusNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotifyWho

`func (o *PurchaseOrderStatusNotification) GetNotifyWho() NotificationRecipientReference`

GetNotifyWho returns the NotifyWho field if non-nil, zero value otherwise.

### GetNotifyWhoOk

`func (o *PurchaseOrderStatusNotification) GetNotifyWhoOk() (*NotificationRecipientReference, bool)`

GetNotifyWhoOk returns a tuple with the NotifyWho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWho

`func (o *PurchaseOrderStatusNotification) SetNotifyWho(v NotificationRecipientReference)`

SetNotifyWho sets NotifyWho field to given value.


### GetStatus

`func (o *PurchaseOrderStatusNotification) GetStatus() PurchaseOrderStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PurchaseOrderStatusNotification) GetStatusOk() (*PurchaseOrderStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PurchaseOrderStatusNotification) SetStatus(v PurchaseOrderStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PurchaseOrderStatusNotification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMember

`func (o *PurchaseOrderStatusNotification) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *PurchaseOrderStatusNotification) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *PurchaseOrderStatusNotification) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *PurchaseOrderStatusNotification) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetEmail

`func (o *PurchaseOrderStatusNotification) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PurchaseOrderStatusNotification) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PurchaseOrderStatusNotification) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PurchaseOrderStatusNotification) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetWorkflowStep

`func (o *PurchaseOrderStatusNotification) GetWorkflowStep() int32`

GetWorkflowStep returns the WorkflowStep field if non-nil, zero value otherwise.

### GetWorkflowStepOk

`func (o *PurchaseOrderStatusNotification) GetWorkflowStepOk() (*int32, bool)`

GetWorkflowStepOk returns a tuple with the WorkflowStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStep

`func (o *PurchaseOrderStatusNotification) SetWorkflowStep(v int32)`

SetWorkflowStep sets WorkflowStep field to given value.

### HasWorkflowStep

`func (o *PurchaseOrderStatusNotification) HasWorkflowStep() bool`

HasWorkflowStep returns a boolean if a field has been set.

### SetWorkflowStepNil

`func (o *PurchaseOrderStatusNotification) SetWorkflowStepNil(b bool)`

 SetWorkflowStepNil sets the value for WorkflowStep to be an explicit nil

### UnsetWorkflowStep
`func (o *PurchaseOrderStatusNotification) UnsetWorkflowStep()`

UnsetWorkflowStep ensures that no value is present for WorkflowStep, not even an explicit nil
### GetInfo

`func (o *PurchaseOrderStatusNotification) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PurchaseOrderStatusNotification) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PurchaseOrderStatusNotification) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PurchaseOrderStatusNotification) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


