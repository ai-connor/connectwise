# ContactGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Group** | [**GroupReference**](GroupReference.md) |  | 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Description** | Pointer to **string** |  Max length: 50; | [optional] 
**UnsubscribeFlag** | Pointer to **NullableBool** |  | [optional] 
**CompanyUnsubcribedEmailMessage** | Pointer to **string** |  | [optional] 
**CompanyGroupUnsubscribedEmailMessage** | Pointer to **string** |  | [optional] 
**ContactUnsubscribedEmailMessage** | Pointer to **string** |  | [optional] 
**ContactGroupUnsubscribedEmailMessage** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewContactGroup

`func NewContactGroup(group GroupReference, ) *ContactGroup`

NewContactGroup instantiates a new ContactGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactGroupWithDefaults

`func NewContactGroupWithDefaults() *ContactGroup`

NewContactGroupWithDefaults instantiates a new ContactGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContactGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroup

`func (o *ContactGroup) GetGroup() GroupReference`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ContactGroup) GetGroupOk() (*GroupReference, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ContactGroup) SetGroup(v GroupReference)`

SetGroup sets Group field to given value.


### GetContact

`func (o *ContactGroup) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ContactGroup) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ContactGroup) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ContactGroup) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDescription

`func (o *ContactGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContactGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContactGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContactGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnsubscribeFlag

`func (o *ContactGroup) GetUnsubscribeFlag() bool`

GetUnsubscribeFlag returns the UnsubscribeFlag field if non-nil, zero value otherwise.

### GetUnsubscribeFlagOk

`func (o *ContactGroup) GetUnsubscribeFlagOk() (*bool, bool)`

GetUnsubscribeFlagOk returns a tuple with the UnsubscribeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribeFlag

`func (o *ContactGroup) SetUnsubscribeFlag(v bool)`

SetUnsubscribeFlag sets UnsubscribeFlag field to given value.

### HasUnsubscribeFlag

`func (o *ContactGroup) HasUnsubscribeFlag() bool`

HasUnsubscribeFlag returns a boolean if a field has been set.

### SetUnsubscribeFlagNil

`func (o *ContactGroup) SetUnsubscribeFlagNil(b bool)`

 SetUnsubscribeFlagNil sets the value for UnsubscribeFlag to be an explicit nil

### UnsetUnsubscribeFlag
`func (o *ContactGroup) UnsetUnsubscribeFlag()`

UnsetUnsubscribeFlag ensures that no value is present for UnsubscribeFlag, not even an explicit nil
### GetCompanyUnsubcribedEmailMessage

`func (o *ContactGroup) GetCompanyUnsubcribedEmailMessage() string`

GetCompanyUnsubcribedEmailMessage returns the CompanyUnsubcribedEmailMessage field if non-nil, zero value otherwise.

### GetCompanyUnsubcribedEmailMessageOk

`func (o *ContactGroup) GetCompanyUnsubcribedEmailMessageOk() (*string, bool)`

GetCompanyUnsubcribedEmailMessageOk returns a tuple with the CompanyUnsubcribedEmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUnsubcribedEmailMessage

`func (o *ContactGroup) SetCompanyUnsubcribedEmailMessage(v string)`

SetCompanyUnsubcribedEmailMessage sets CompanyUnsubcribedEmailMessage field to given value.

### HasCompanyUnsubcribedEmailMessage

`func (o *ContactGroup) HasCompanyUnsubcribedEmailMessage() bool`

HasCompanyUnsubcribedEmailMessage returns a boolean if a field has been set.

### GetCompanyGroupUnsubscribedEmailMessage

`func (o *ContactGroup) GetCompanyGroupUnsubscribedEmailMessage() string`

GetCompanyGroupUnsubscribedEmailMessage returns the CompanyGroupUnsubscribedEmailMessage field if non-nil, zero value otherwise.

### GetCompanyGroupUnsubscribedEmailMessageOk

`func (o *ContactGroup) GetCompanyGroupUnsubscribedEmailMessageOk() (*string, bool)`

GetCompanyGroupUnsubscribedEmailMessageOk returns a tuple with the CompanyGroupUnsubscribedEmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyGroupUnsubscribedEmailMessage

`func (o *ContactGroup) SetCompanyGroupUnsubscribedEmailMessage(v string)`

SetCompanyGroupUnsubscribedEmailMessage sets CompanyGroupUnsubscribedEmailMessage field to given value.

### HasCompanyGroupUnsubscribedEmailMessage

`func (o *ContactGroup) HasCompanyGroupUnsubscribedEmailMessage() bool`

HasCompanyGroupUnsubscribedEmailMessage returns a boolean if a field has been set.

### GetContactUnsubscribedEmailMessage

`func (o *ContactGroup) GetContactUnsubscribedEmailMessage() string`

GetContactUnsubscribedEmailMessage returns the ContactUnsubscribedEmailMessage field if non-nil, zero value otherwise.

### GetContactUnsubscribedEmailMessageOk

`func (o *ContactGroup) GetContactUnsubscribedEmailMessageOk() (*string, bool)`

GetContactUnsubscribedEmailMessageOk returns a tuple with the ContactUnsubscribedEmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactUnsubscribedEmailMessage

`func (o *ContactGroup) SetContactUnsubscribedEmailMessage(v string)`

SetContactUnsubscribedEmailMessage sets ContactUnsubscribedEmailMessage field to given value.

### HasContactUnsubscribedEmailMessage

`func (o *ContactGroup) HasContactUnsubscribedEmailMessage() bool`

HasContactUnsubscribedEmailMessage returns a boolean if a field has been set.

### GetContactGroupUnsubscribedEmailMessage

`func (o *ContactGroup) GetContactGroupUnsubscribedEmailMessage() string`

GetContactGroupUnsubscribedEmailMessage returns the ContactGroupUnsubscribedEmailMessage field if non-nil, zero value otherwise.

### GetContactGroupUnsubscribedEmailMessageOk

`func (o *ContactGroup) GetContactGroupUnsubscribedEmailMessageOk() (*string, bool)`

GetContactGroupUnsubscribedEmailMessageOk returns a tuple with the ContactGroupUnsubscribedEmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactGroupUnsubscribedEmailMessage

`func (o *ContactGroup) SetContactGroupUnsubscribedEmailMessage(v string)`

SetContactGroupUnsubscribedEmailMessage sets ContactGroupUnsubscribedEmailMessage field to given value.

### HasContactGroupUnsubscribedEmailMessage

`func (o *ContactGroup) HasContactGroupUnsubscribedEmailMessage() bool`

HasContactGroupUnsubscribedEmailMessage returns a boolean if a field has been set.

### GetInfo

`func (o *ContactGroup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ContactGroup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ContactGroup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ContactGroup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


