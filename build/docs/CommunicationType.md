# CommunicationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Description** | **string** |  | 
**PhoneFlag** | Pointer to **NullableBool** | Gets or sets at least one flag is required to be true -- phone, fax, or email. | [optional] 
**FaxFlag** | Pointer to **NullableBool** | Gets or sets at least one flag is required to be true -- phone, fax, or email. | [optional] 
**EmailFlag** | Pointer to **NullableBool** | Gets or sets at least one flag is required to be true -- phone, fax, or email. | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ExchangeXref** | Pointer to **string** |  Max length: 50; | [optional] 
**IphoneXref** | Pointer to **string** |  Max length: 50; | [optional] 
**AndroidXref** | Pointer to **string** |  Max length: 50; | [optional] 
**GoogleXref** | Pointer to **string** |  Max length: 50; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCommunicationType

`func NewCommunicationType(description string, ) *CommunicationType`

NewCommunicationType instantiates a new CommunicationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationTypeWithDefaults

`func NewCommunicationTypeWithDefaults() *CommunicationType`

NewCommunicationTypeWithDefaults instantiates a new CommunicationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommunicationType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommunicationType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommunicationType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CommunicationType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *CommunicationType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommunicationType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommunicationType) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPhoneFlag

`func (o *CommunicationType) GetPhoneFlag() bool`

GetPhoneFlag returns the PhoneFlag field if non-nil, zero value otherwise.

### GetPhoneFlagOk

`func (o *CommunicationType) GetPhoneFlagOk() (*bool, bool)`

GetPhoneFlagOk returns a tuple with the PhoneFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneFlag

`func (o *CommunicationType) SetPhoneFlag(v bool)`

SetPhoneFlag sets PhoneFlag field to given value.

### HasPhoneFlag

`func (o *CommunicationType) HasPhoneFlag() bool`

HasPhoneFlag returns a boolean if a field has been set.

### SetPhoneFlagNil

`func (o *CommunicationType) SetPhoneFlagNil(b bool)`

 SetPhoneFlagNil sets the value for PhoneFlag to be an explicit nil

### UnsetPhoneFlag
`func (o *CommunicationType) UnsetPhoneFlag()`

UnsetPhoneFlag ensures that no value is present for PhoneFlag, not even an explicit nil
### GetFaxFlag

`func (o *CommunicationType) GetFaxFlag() bool`

GetFaxFlag returns the FaxFlag field if non-nil, zero value otherwise.

### GetFaxFlagOk

`func (o *CommunicationType) GetFaxFlagOk() (*bool, bool)`

GetFaxFlagOk returns a tuple with the FaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxFlag

`func (o *CommunicationType) SetFaxFlag(v bool)`

SetFaxFlag sets FaxFlag field to given value.

### HasFaxFlag

`func (o *CommunicationType) HasFaxFlag() bool`

HasFaxFlag returns a boolean if a field has been set.

### SetFaxFlagNil

`func (o *CommunicationType) SetFaxFlagNil(b bool)`

 SetFaxFlagNil sets the value for FaxFlag to be an explicit nil

### UnsetFaxFlag
`func (o *CommunicationType) UnsetFaxFlag()`

UnsetFaxFlag ensures that no value is present for FaxFlag, not even an explicit nil
### GetEmailFlag

`func (o *CommunicationType) GetEmailFlag() bool`

GetEmailFlag returns the EmailFlag field if non-nil, zero value otherwise.

### GetEmailFlagOk

`func (o *CommunicationType) GetEmailFlagOk() (*bool, bool)`

GetEmailFlagOk returns a tuple with the EmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFlag

`func (o *CommunicationType) SetEmailFlag(v bool)`

SetEmailFlag sets EmailFlag field to given value.

### HasEmailFlag

`func (o *CommunicationType) HasEmailFlag() bool`

HasEmailFlag returns a boolean if a field has been set.

### SetEmailFlagNil

`func (o *CommunicationType) SetEmailFlagNil(b bool)`

 SetEmailFlagNil sets the value for EmailFlag to be an explicit nil

### UnsetEmailFlag
`func (o *CommunicationType) UnsetEmailFlag()`

UnsetEmailFlag ensures that no value is present for EmailFlag, not even an explicit nil
### GetDefaultFlag

`func (o *CommunicationType) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *CommunicationType) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *CommunicationType) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *CommunicationType) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *CommunicationType) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *CommunicationType) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetExchangeXref

`func (o *CommunicationType) GetExchangeXref() string`

GetExchangeXref returns the ExchangeXref field if non-nil, zero value otherwise.

### GetExchangeXrefOk

`func (o *CommunicationType) GetExchangeXrefOk() (*string, bool)`

GetExchangeXrefOk returns a tuple with the ExchangeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeXref

`func (o *CommunicationType) SetExchangeXref(v string)`

SetExchangeXref sets ExchangeXref field to given value.

### HasExchangeXref

`func (o *CommunicationType) HasExchangeXref() bool`

HasExchangeXref returns a boolean if a field has been set.

### GetIphoneXref

`func (o *CommunicationType) GetIphoneXref() string`

GetIphoneXref returns the IphoneXref field if non-nil, zero value otherwise.

### GetIphoneXrefOk

`func (o *CommunicationType) GetIphoneXrefOk() (*string, bool)`

GetIphoneXrefOk returns a tuple with the IphoneXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIphoneXref

`func (o *CommunicationType) SetIphoneXref(v string)`

SetIphoneXref sets IphoneXref field to given value.

### HasIphoneXref

`func (o *CommunicationType) HasIphoneXref() bool`

HasIphoneXref returns a boolean if a field has been set.

### GetAndroidXref

`func (o *CommunicationType) GetAndroidXref() string`

GetAndroidXref returns the AndroidXref field if non-nil, zero value otherwise.

### GetAndroidXrefOk

`func (o *CommunicationType) GetAndroidXrefOk() (*string, bool)`

GetAndroidXrefOk returns a tuple with the AndroidXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidXref

`func (o *CommunicationType) SetAndroidXref(v string)`

SetAndroidXref sets AndroidXref field to given value.

### HasAndroidXref

`func (o *CommunicationType) HasAndroidXref() bool`

HasAndroidXref returns a boolean if a field has been set.

### GetGoogleXref

`func (o *CommunicationType) GetGoogleXref() string`

GetGoogleXref returns the GoogleXref field if non-nil, zero value otherwise.

### GetGoogleXrefOk

`func (o *CommunicationType) GetGoogleXrefOk() (*string, bool)`

GetGoogleXrefOk returns a tuple with the GoogleXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleXref

`func (o *CommunicationType) SetGoogleXref(v string)`

SetGoogleXref sets GoogleXref field to given value.

### HasGoogleXref

`func (o *CommunicationType) HasGoogleXref() bool`

HasGoogleXref returns a boolean if a field has been set.

### GetInfo

`func (o *CommunicationType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CommunicationType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CommunicationType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CommunicationType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


