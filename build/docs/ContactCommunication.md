# ContactCommunication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ContactId** | Pointer to **NullableInt32** |  | [optional] 
**Type** | [**CommunicationTypeReference**](CommunicationTypeReference.md) |  | 
**Value** | **string** |  Max length: 250; | 
**Extension** | Pointer to **string** |  Max length: 15; | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**CommunicationType** | Pointer to **NullableString** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewContactCommunication

`func NewContactCommunication(type_ CommunicationTypeReference, value string, ) *ContactCommunication`

NewContactCommunication instantiates a new ContactCommunication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactCommunicationWithDefaults

`func NewContactCommunicationWithDefaults() *ContactCommunication`

NewContactCommunicationWithDefaults instantiates a new ContactCommunication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactCommunication) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactCommunication) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactCommunication) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContactCommunication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContactId

`func (o *ContactCommunication) GetContactId() int32`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ContactCommunication) GetContactIdOk() (*int32, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ContactCommunication) SetContactId(v int32)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *ContactCommunication) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *ContactCommunication) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *ContactCommunication) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetType

`func (o *ContactCommunication) GetType() CommunicationTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactCommunication) GetTypeOk() (*CommunicationTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactCommunication) SetType(v CommunicationTypeReference)`

SetType sets Type field to given value.


### GetValue

`func (o *ContactCommunication) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContactCommunication) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContactCommunication) SetValue(v string)`

SetValue sets Value field to given value.


### GetExtension

`func (o *ContactCommunication) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *ContactCommunication) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *ContactCommunication) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *ContactCommunication) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *ContactCommunication) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ContactCommunication) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ContactCommunication) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ContactCommunication) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ContactCommunication) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ContactCommunication) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetMobileGuid

`func (o *ContactCommunication) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *ContactCommunication) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *ContactCommunication) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *ContactCommunication) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *ContactCommunication) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *ContactCommunication) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetCommunicationType

`func (o *ContactCommunication) GetCommunicationType() string`

GetCommunicationType returns the CommunicationType field if non-nil, zero value otherwise.

### GetCommunicationTypeOk

`func (o *ContactCommunication) GetCommunicationTypeOk() (*string, bool)`

GetCommunicationTypeOk returns a tuple with the CommunicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationType

`func (o *ContactCommunication) SetCommunicationType(v string)`

SetCommunicationType sets CommunicationType field to given value.

### HasCommunicationType

`func (o *ContactCommunication) HasCommunicationType() bool`

HasCommunicationType returns a boolean if a field has been set.

### SetCommunicationTypeNil

`func (o *ContactCommunication) SetCommunicationTypeNil(b bool)`

 SetCommunicationTypeNil sets the value for CommunicationType to be an explicit nil

### UnsetCommunicationType
`func (o *ContactCommunication) UnsetCommunicationType()`

UnsetCommunicationType ensures that no value is present for CommunicationType, not even an explicit nil
### GetDomain

`func (o *ContactCommunication) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ContactCommunication) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ContactCommunication) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ContactCommunication) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetInfo

`func (o *ContactCommunication) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ContactCommunication) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ContactCommunication) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ContactCommunication) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


