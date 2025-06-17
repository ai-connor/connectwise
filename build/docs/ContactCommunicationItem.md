# ContactCommunicationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**CommunicationTypeReference**](CommunicationTypeReference.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**CommunicationType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContactCommunicationItem

`func NewContactCommunicationItem() *ContactCommunicationItem`

NewContactCommunicationItem instantiates a new ContactCommunicationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactCommunicationItemWithDefaults

`func NewContactCommunicationItemWithDefaults() *ContactCommunicationItem`

NewContactCommunicationItemWithDefaults instantiates a new ContactCommunicationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactCommunicationItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactCommunicationItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactCommunicationItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContactCommunicationItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ContactCommunicationItem) GetType() CommunicationTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactCommunicationItem) GetTypeOk() (*CommunicationTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactCommunicationItem) SetType(v CommunicationTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *ContactCommunicationItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ContactCommunicationItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContactCommunicationItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContactCommunicationItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContactCommunicationItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetExtension

`func (o *ContactCommunicationItem) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *ContactCommunicationItem) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *ContactCommunicationItem) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *ContactCommunicationItem) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *ContactCommunicationItem) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ContactCommunicationItem) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ContactCommunicationItem) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ContactCommunicationItem) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ContactCommunicationItem) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ContactCommunicationItem) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetDomain

`func (o *ContactCommunicationItem) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ContactCommunicationItem) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ContactCommunicationItem) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ContactCommunicationItem) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetCommunicationType

`func (o *ContactCommunicationItem) GetCommunicationType() string`

GetCommunicationType returns the CommunicationType field if non-nil, zero value otherwise.

### GetCommunicationTypeOk

`func (o *ContactCommunicationItem) GetCommunicationTypeOk() (*string, bool)`

GetCommunicationTypeOk returns a tuple with the CommunicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationType

`func (o *ContactCommunicationItem) SetCommunicationType(v string)`

SetCommunicationType sets CommunicationType field to given value.

### HasCommunicationType

`func (o *ContactCommunicationItem) HasCommunicationType() bool`

HasCommunicationType returns a boolean if a field has been set.

### SetCommunicationTypeNil

`func (o *ContactCommunicationItem) SetCommunicationTypeNil(b bool)`

 SetCommunicationTypeNil sets the value for CommunicationType to be an explicit nil

### UnsetCommunicationType
`func (o *ContactCommunicationItem) UnsetCommunicationType()`

UnsetCommunicationType ensures that no value is present for CommunicationType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


