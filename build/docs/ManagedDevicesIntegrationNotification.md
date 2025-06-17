# ManagedDevicesIntegrationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ManagedDevicesIntegration** | Pointer to [**ManagedDevicesIntegrationReference**](ManagedDevicesIntegrationReference.md) |  | [optional] 
**NotifyWho** | [**NotificationRecipientReference**](NotificationRecipientReference.md) |  | 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**LogType** | **NullableString** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagedDevicesIntegrationNotification

`func NewManagedDevicesIntegrationNotification(notifyWho NotificationRecipientReference, logType NullableString, ) *ManagedDevicesIntegrationNotification`

NewManagedDevicesIntegrationNotification instantiates a new ManagedDevicesIntegrationNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedDevicesIntegrationNotificationWithDefaults

`func NewManagedDevicesIntegrationNotificationWithDefaults() *ManagedDevicesIntegrationNotification`

NewManagedDevicesIntegrationNotificationWithDefaults instantiates a new ManagedDevicesIntegrationNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedDevicesIntegrationNotification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedDevicesIntegrationNotification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedDevicesIntegrationNotification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedDevicesIntegrationNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManagedDevicesIntegration

`func (o *ManagedDevicesIntegrationNotification) GetManagedDevicesIntegration() ManagedDevicesIntegrationReference`

GetManagedDevicesIntegration returns the ManagedDevicesIntegration field if non-nil, zero value otherwise.

### GetManagedDevicesIntegrationOk

`func (o *ManagedDevicesIntegrationNotification) GetManagedDevicesIntegrationOk() (*ManagedDevicesIntegrationReference, bool)`

GetManagedDevicesIntegrationOk returns a tuple with the ManagedDevicesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDevicesIntegration

`func (o *ManagedDevicesIntegrationNotification) SetManagedDevicesIntegration(v ManagedDevicesIntegrationReference)`

SetManagedDevicesIntegration sets ManagedDevicesIntegration field to given value.

### HasManagedDevicesIntegration

`func (o *ManagedDevicesIntegrationNotification) HasManagedDevicesIntegration() bool`

HasManagedDevicesIntegration returns a boolean if a field has been set.

### GetNotifyWho

`func (o *ManagedDevicesIntegrationNotification) GetNotifyWho() NotificationRecipientReference`

GetNotifyWho returns the NotifyWho field if non-nil, zero value otherwise.

### GetNotifyWhoOk

`func (o *ManagedDevicesIntegrationNotification) GetNotifyWhoOk() (*NotificationRecipientReference, bool)`

GetNotifyWhoOk returns a tuple with the NotifyWho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWho

`func (o *ManagedDevicesIntegrationNotification) SetNotifyWho(v NotificationRecipientReference)`

SetNotifyWho sets NotifyWho field to given value.


### GetMember

`func (o *ManagedDevicesIntegrationNotification) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ManagedDevicesIntegrationNotification) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ManagedDevicesIntegrationNotification) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *ManagedDevicesIntegrationNotification) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetLogType

`func (o *ManagedDevicesIntegrationNotification) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *ManagedDevicesIntegrationNotification) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *ManagedDevicesIntegrationNotification) SetLogType(v string)`

SetLogType sets LogType field to given value.


### SetLogTypeNil

`func (o *ManagedDevicesIntegrationNotification) SetLogTypeNil(b bool)`

 SetLogTypeNil sets the value for LogType to be an explicit nil

### UnsetLogType
`func (o *ManagedDevicesIntegrationNotification) UnsetLogType()`

UnsetLogType ensures that no value is present for LogType, not even an explicit nil
### GetInfo

`func (o *ManagedDevicesIntegrationNotification) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ManagedDevicesIntegrationNotification) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ManagedDevicesIntegrationNotification) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ManagedDevicesIntegrationNotification) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


