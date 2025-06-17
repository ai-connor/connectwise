# ManagementReportNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NotifyWho** | [**NotificationRecipientReference**](NotificationRecipientReference.md) |  | 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Email** | Pointer to **string** |  Max length: 50; | [optional] 
**GlobalFlag** | Pointer to **NullableBool** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagementReportNotification

`func NewManagementReportNotification(notifyWho NotificationRecipientReference, ) *ManagementReportNotification`

NewManagementReportNotification instantiates a new ManagementReportNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementReportNotificationWithDefaults

`func NewManagementReportNotificationWithDefaults() *ManagementReportNotification`

NewManagementReportNotificationWithDefaults instantiates a new ManagementReportNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagementReportNotification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagementReportNotification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagementReportNotification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagementReportNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotifyWho

`func (o *ManagementReportNotification) GetNotifyWho() NotificationRecipientReference`

GetNotifyWho returns the NotifyWho field if non-nil, zero value otherwise.

### GetNotifyWhoOk

`func (o *ManagementReportNotification) GetNotifyWhoOk() (*NotificationRecipientReference, bool)`

GetNotifyWhoOk returns a tuple with the NotifyWho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWho

`func (o *ManagementReportNotification) SetNotifyWho(v NotificationRecipientReference)`

SetNotifyWho sets NotifyWho field to given value.


### GetMember

`func (o *ManagementReportNotification) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ManagementReportNotification) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ManagementReportNotification) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *ManagementReportNotification) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetEmail

`func (o *ManagementReportNotification) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ManagementReportNotification) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ManagementReportNotification) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ManagementReportNotification) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGlobalFlag

`func (o *ManagementReportNotification) GetGlobalFlag() bool`

GetGlobalFlag returns the GlobalFlag field if non-nil, zero value otherwise.

### GetGlobalFlagOk

`func (o *ManagementReportNotification) GetGlobalFlagOk() (*bool, bool)`

GetGlobalFlagOk returns a tuple with the GlobalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalFlag

`func (o *ManagementReportNotification) SetGlobalFlag(v bool)`

SetGlobalFlag sets GlobalFlag field to given value.

### HasGlobalFlag

`func (o *ManagementReportNotification) HasGlobalFlag() bool`

HasGlobalFlag returns a boolean if a field has been set.

### SetGlobalFlagNil

`func (o *ManagementReportNotification) SetGlobalFlagNil(b bool)`

 SetGlobalFlagNil sets the value for GlobalFlag to be an explicit nil

### UnsetGlobalFlag
`func (o *ManagementReportNotification) UnsetGlobalFlag()`

UnsetGlobalFlag ensures that no value is present for GlobalFlag, not even an explicit nil
### GetCompany

`func (o *ManagementReportNotification) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ManagementReportNotification) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ManagementReportNotification) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ManagementReportNotification) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetInfo

`func (o *ManagementReportNotification) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ManagementReportNotification) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ManagementReportNotification) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ManagementReportNotification) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


