# MemberNotificationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NotificationType** | **NullableString** |  | 
**NotificationTrigger** | **NullableString** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMemberNotificationSetting

`func NewMemberNotificationSetting(notificationType NullableString, notificationTrigger NullableString, ) *MemberNotificationSetting`

NewMemberNotificationSetting instantiates a new MemberNotificationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberNotificationSettingWithDefaults

`func NewMemberNotificationSettingWithDefaults() *MemberNotificationSetting`

NewMemberNotificationSettingWithDefaults instantiates a new MemberNotificationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberNotificationSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberNotificationSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberNotificationSetting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MemberNotificationSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotificationType

`func (o *MemberNotificationSetting) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *MemberNotificationSetting) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *MemberNotificationSetting) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### SetNotificationTypeNil

`func (o *MemberNotificationSetting) SetNotificationTypeNil(b bool)`

 SetNotificationTypeNil sets the value for NotificationType to be an explicit nil

### UnsetNotificationType
`func (o *MemberNotificationSetting) UnsetNotificationType()`

UnsetNotificationType ensures that no value is present for NotificationType, not even an explicit nil
### GetNotificationTrigger

`func (o *MemberNotificationSetting) GetNotificationTrigger() string`

GetNotificationTrigger returns the NotificationTrigger field if non-nil, zero value otherwise.

### GetNotificationTriggerOk

`func (o *MemberNotificationSetting) GetNotificationTriggerOk() (*string, bool)`

GetNotificationTriggerOk returns a tuple with the NotificationTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTrigger

`func (o *MemberNotificationSetting) SetNotificationTrigger(v string)`

SetNotificationTrigger sets NotificationTrigger field to given value.


### SetNotificationTriggerNil

`func (o *MemberNotificationSetting) SetNotificationTriggerNil(b bool)`

 SetNotificationTriggerNil sets the value for NotificationTrigger to be an explicit nil

### UnsetNotificationTrigger
`func (o *MemberNotificationSetting) UnsetNotificationTrigger()`

UnsetNotificationTrigger ensures that no value is present for NotificationTrigger, not even an explicit nil
### GetInfo

`func (o *MemberNotificationSetting) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MemberNotificationSetting) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MemberNotificationSetting) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MemberNotificationSetting) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


