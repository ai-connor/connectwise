# ServiceSurvey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**HeaderIncludeLogoFlag** | Pointer to **NullableBool** |  | [optional] 
**HeaderText** | Pointer to **string** |  Max length: 4000; | [optional] 
**HeaderTextVisibleFlag** | Pointer to **NullableBool** |  | [optional] 
**FooterText** | Pointer to **string** |  Max length: 500; | [optional] 
**FooterTextVisibleFlag** | Pointer to **NullableBool** |  | [optional] 
**ThankYouText** | Pointer to **string** |  Max length: 4000; | [optional] 
**NotifyWho** | Pointer to [**GenericIdIdentifierReference**](GenericIdIdentifierReference.md) |  | [optional] 
**NotifyWhoVisibleFlag** | Pointer to **NullableBool** |  | [optional] 
**NotifyMember** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceSurvey

`func NewServiceSurvey(name string, ) *ServiceSurvey`

NewServiceSurvey instantiates a new ServiceSurvey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSurveyWithDefaults

`func NewServiceSurveyWithDefaults() *ServiceSurvey`

NewServiceSurveyWithDefaults instantiates a new ServiceSurvey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceSurvey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceSurvey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceSurvey) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceSurvey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceSurvey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceSurvey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceSurvey) SetName(v string)`

SetName sets Name field to given value.


### GetInactiveFlag

`func (o *ServiceSurvey) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ServiceSurvey) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ServiceSurvey) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ServiceSurvey) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ServiceSurvey) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ServiceSurvey) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetHeaderIncludeLogoFlag

`func (o *ServiceSurvey) GetHeaderIncludeLogoFlag() bool`

GetHeaderIncludeLogoFlag returns the HeaderIncludeLogoFlag field if non-nil, zero value otherwise.

### GetHeaderIncludeLogoFlagOk

`func (o *ServiceSurvey) GetHeaderIncludeLogoFlagOk() (*bool, bool)`

GetHeaderIncludeLogoFlagOk returns a tuple with the HeaderIncludeLogoFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderIncludeLogoFlag

`func (o *ServiceSurvey) SetHeaderIncludeLogoFlag(v bool)`

SetHeaderIncludeLogoFlag sets HeaderIncludeLogoFlag field to given value.

### HasHeaderIncludeLogoFlag

`func (o *ServiceSurvey) HasHeaderIncludeLogoFlag() bool`

HasHeaderIncludeLogoFlag returns a boolean if a field has been set.

### SetHeaderIncludeLogoFlagNil

`func (o *ServiceSurvey) SetHeaderIncludeLogoFlagNil(b bool)`

 SetHeaderIncludeLogoFlagNil sets the value for HeaderIncludeLogoFlag to be an explicit nil

### UnsetHeaderIncludeLogoFlag
`func (o *ServiceSurvey) UnsetHeaderIncludeLogoFlag()`

UnsetHeaderIncludeLogoFlag ensures that no value is present for HeaderIncludeLogoFlag, not even an explicit nil
### GetHeaderText

`func (o *ServiceSurvey) GetHeaderText() string`

GetHeaderText returns the HeaderText field if non-nil, zero value otherwise.

### GetHeaderTextOk

`func (o *ServiceSurvey) GetHeaderTextOk() (*string, bool)`

GetHeaderTextOk returns a tuple with the HeaderText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderText

`func (o *ServiceSurvey) SetHeaderText(v string)`

SetHeaderText sets HeaderText field to given value.

### HasHeaderText

`func (o *ServiceSurvey) HasHeaderText() bool`

HasHeaderText returns a boolean if a field has been set.

### GetHeaderTextVisibleFlag

`func (o *ServiceSurvey) GetHeaderTextVisibleFlag() bool`

GetHeaderTextVisibleFlag returns the HeaderTextVisibleFlag field if non-nil, zero value otherwise.

### GetHeaderTextVisibleFlagOk

`func (o *ServiceSurvey) GetHeaderTextVisibleFlagOk() (*bool, bool)`

GetHeaderTextVisibleFlagOk returns a tuple with the HeaderTextVisibleFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderTextVisibleFlag

`func (o *ServiceSurvey) SetHeaderTextVisibleFlag(v bool)`

SetHeaderTextVisibleFlag sets HeaderTextVisibleFlag field to given value.

### HasHeaderTextVisibleFlag

`func (o *ServiceSurvey) HasHeaderTextVisibleFlag() bool`

HasHeaderTextVisibleFlag returns a boolean if a field has been set.

### SetHeaderTextVisibleFlagNil

`func (o *ServiceSurvey) SetHeaderTextVisibleFlagNil(b bool)`

 SetHeaderTextVisibleFlagNil sets the value for HeaderTextVisibleFlag to be an explicit nil

### UnsetHeaderTextVisibleFlag
`func (o *ServiceSurvey) UnsetHeaderTextVisibleFlag()`

UnsetHeaderTextVisibleFlag ensures that no value is present for HeaderTextVisibleFlag, not even an explicit nil
### GetFooterText

`func (o *ServiceSurvey) GetFooterText() string`

GetFooterText returns the FooterText field if non-nil, zero value otherwise.

### GetFooterTextOk

`func (o *ServiceSurvey) GetFooterTextOk() (*string, bool)`

GetFooterTextOk returns a tuple with the FooterText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterText

`func (o *ServiceSurvey) SetFooterText(v string)`

SetFooterText sets FooterText field to given value.

### HasFooterText

`func (o *ServiceSurvey) HasFooterText() bool`

HasFooterText returns a boolean if a field has been set.

### GetFooterTextVisibleFlag

`func (o *ServiceSurvey) GetFooterTextVisibleFlag() bool`

GetFooterTextVisibleFlag returns the FooterTextVisibleFlag field if non-nil, zero value otherwise.

### GetFooterTextVisibleFlagOk

`func (o *ServiceSurvey) GetFooterTextVisibleFlagOk() (*bool, bool)`

GetFooterTextVisibleFlagOk returns a tuple with the FooterTextVisibleFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterTextVisibleFlag

`func (o *ServiceSurvey) SetFooterTextVisibleFlag(v bool)`

SetFooterTextVisibleFlag sets FooterTextVisibleFlag field to given value.

### HasFooterTextVisibleFlag

`func (o *ServiceSurvey) HasFooterTextVisibleFlag() bool`

HasFooterTextVisibleFlag returns a boolean if a field has been set.

### SetFooterTextVisibleFlagNil

`func (o *ServiceSurvey) SetFooterTextVisibleFlagNil(b bool)`

 SetFooterTextVisibleFlagNil sets the value for FooterTextVisibleFlag to be an explicit nil

### UnsetFooterTextVisibleFlag
`func (o *ServiceSurvey) UnsetFooterTextVisibleFlag()`

UnsetFooterTextVisibleFlag ensures that no value is present for FooterTextVisibleFlag, not even an explicit nil
### GetThankYouText

`func (o *ServiceSurvey) GetThankYouText() string`

GetThankYouText returns the ThankYouText field if non-nil, zero value otherwise.

### GetThankYouTextOk

`func (o *ServiceSurvey) GetThankYouTextOk() (*string, bool)`

GetThankYouTextOk returns a tuple with the ThankYouText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouText

`func (o *ServiceSurvey) SetThankYouText(v string)`

SetThankYouText sets ThankYouText field to given value.

### HasThankYouText

`func (o *ServiceSurvey) HasThankYouText() bool`

HasThankYouText returns a boolean if a field has been set.

### GetNotifyWho

`func (o *ServiceSurvey) GetNotifyWho() GenericIdIdentifierReference`

GetNotifyWho returns the NotifyWho field if non-nil, zero value otherwise.

### GetNotifyWhoOk

`func (o *ServiceSurvey) GetNotifyWhoOk() (*GenericIdIdentifierReference, bool)`

GetNotifyWhoOk returns a tuple with the NotifyWho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWho

`func (o *ServiceSurvey) SetNotifyWho(v GenericIdIdentifierReference)`

SetNotifyWho sets NotifyWho field to given value.

### HasNotifyWho

`func (o *ServiceSurvey) HasNotifyWho() bool`

HasNotifyWho returns a boolean if a field has been set.

### GetNotifyWhoVisibleFlag

`func (o *ServiceSurvey) GetNotifyWhoVisibleFlag() bool`

GetNotifyWhoVisibleFlag returns the NotifyWhoVisibleFlag field if non-nil, zero value otherwise.

### GetNotifyWhoVisibleFlagOk

`func (o *ServiceSurvey) GetNotifyWhoVisibleFlagOk() (*bool, bool)`

GetNotifyWhoVisibleFlagOk returns a tuple with the NotifyWhoVisibleFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWhoVisibleFlag

`func (o *ServiceSurvey) SetNotifyWhoVisibleFlag(v bool)`

SetNotifyWhoVisibleFlag sets NotifyWhoVisibleFlag field to given value.

### HasNotifyWhoVisibleFlag

`func (o *ServiceSurvey) HasNotifyWhoVisibleFlag() bool`

HasNotifyWhoVisibleFlag returns a boolean if a field has been set.

### SetNotifyWhoVisibleFlagNil

`func (o *ServiceSurvey) SetNotifyWhoVisibleFlagNil(b bool)`

 SetNotifyWhoVisibleFlagNil sets the value for NotifyWhoVisibleFlag to be an explicit nil

### UnsetNotifyWhoVisibleFlag
`func (o *ServiceSurvey) UnsetNotifyWhoVisibleFlag()`

UnsetNotifyWhoVisibleFlag ensures that no value is present for NotifyWhoVisibleFlag, not even an explicit nil
### GetNotifyMember

`func (o *ServiceSurvey) GetNotifyMember() MemberReference`

GetNotifyMember returns the NotifyMember field if non-nil, zero value otherwise.

### GetNotifyMemberOk

`func (o *ServiceSurvey) GetNotifyMemberOk() (*MemberReference, bool)`

GetNotifyMemberOk returns a tuple with the NotifyMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMember

`func (o *ServiceSurvey) SetNotifyMember(v MemberReference)`

SetNotifyMember sets NotifyMember field to given value.

### HasNotifyMember

`func (o *ServiceSurvey) HasNotifyMember() bool`

HasNotifyMember returns a boolean if a field has been set.

### GetInfo

`func (o *ServiceSurvey) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceSurvey) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceSurvey) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceSurvey) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


