# CommunicationTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PhoneFlag** | Pointer to **NullableBool** |  | [optional] 
**FaxFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCommunicationTypeInfo

`func NewCommunicationTypeInfo() *CommunicationTypeInfo`

NewCommunicationTypeInfo instantiates a new CommunicationTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationTypeInfoWithDefaults

`func NewCommunicationTypeInfoWithDefaults() *CommunicationTypeInfo`

NewCommunicationTypeInfoWithDefaults instantiates a new CommunicationTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommunicationTypeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommunicationTypeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommunicationTypeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CommunicationTypeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *CommunicationTypeInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommunicationTypeInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommunicationTypeInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommunicationTypeInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPhoneFlag

`func (o *CommunicationTypeInfo) GetPhoneFlag() bool`

GetPhoneFlag returns the PhoneFlag field if non-nil, zero value otherwise.

### GetPhoneFlagOk

`func (o *CommunicationTypeInfo) GetPhoneFlagOk() (*bool, bool)`

GetPhoneFlagOk returns a tuple with the PhoneFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneFlag

`func (o *CommunicationTypeInfo) SetPhoneFlag(v bool)`

SetPhoneFlag sets PhoneFlag field to given value.

### HasPhoneFlag

`func (o *CommunicationTypeInfo) HasPhoneFlag() bool`

HasPhoneFlag returns a boolean if a field has been set.

### SetPhoneFlagNil

`func (o *CommunicationTypeInfo) SetPhoneFlagNil(b bool)`

 SetPhoneFlagNil sets the value for PhoneFlag to be an explicit nil

### UnsetPhoneFlag
`func (o *CommunicationTypeInfo) UnsetPhoneFlag()`

UnsetPhoneFlag ensures that no value is present for PhoneFlag, not even an explicit nil
### GetFaxFlag

`func (o *CommunicationTypeInfo) GetFaxFlag() bool`

GetFaxFlag returns the FaxFlag field if non-nil, zero value otherwise.

### GetFaxFlagOk

`func (o *CommunicationTypeInfo) GetFaxFlagOk() (*bool, bool)`

GetFaxFlagOk returns a tuple with the FaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxFlag

`func (o *CommunicationTypeInfo) SetFaxFlag(v bool)`

SetFaxFlag sets FaxFlag field to given value.

### HasFaxFlag

`func (o *CommunicationTypeInfo) HasFaxFlag() bool`

HasFaxFlag returns a boolean if a field has been set.

### SetFaxFlagNil

`func (o *CommunicationTypeInfo) SetFaxFlagNil(b bool)`

 SetFaxFlagNil sets the value for FaxFlag to be an explicit nil

### UnsetFaxFlag
`func (o *CommunicationTypeInfo) UnsetFaxFlag()`

UnsetFaxFlag ensures that no value is present for FaxFlag, not even an explicit nil
### GetEmailFlag

`func (o *CommunicationTypeInfo) GetEmailFlag() bool`

GetEmailFlag returns the EmailFlag field if non-nil, zero value otherwise.

### GetEmailFlagOk

`func (o *CommunicationTypeInfo) GetEmailFlagOk() (*bool, bool)`

GetEmailFlagOk returns a tuple with the EmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFlag

`func (o *CommunicationTypeInfo) SetEmailFlag(v bool)`

SetEmailFlag sets EmailFlag field to given value.

### HasEmailFlag

`func (o *CommunicationTypeInfo) HasEmailFlag() bool`

HasEmailFlag returns a boolean if a field has been set.

### SetEmailFlagNil

`func (o *CommunicationTypeInfo) SetEmailFlagNil(b bool)`

 SetEmailFlagNil sets the value for EmailFlag to be an explicit nil

### UnsetEmailFlag
`func (o *CommunicationTypeInfo) UnsetEmailFlag()`

UnsetEmailFlag ensures that no value is present for EmailFlag, not even an explicit nil
### GetDefaultFlag

`func (o *CommunicationTypeInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *CommunicationTypeInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *CommunicationTypeInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *CommunicationTypeInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *CommunicationTypeInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *CommunicationTypeInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInfo

`func (o *CommunicationTypeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CommunicationTypeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CommunicationTypeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CommunicationTypeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


