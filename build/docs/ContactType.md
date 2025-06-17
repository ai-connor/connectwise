# ContactType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Description** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceAlertFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceAlertMessage** | Pointer to **string** |  Max length: 150; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewContactType

`func NewContactType(description string, ) *ContactType`

NewContactType instantiates a new ContactType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactTypeWithDefaults

`func NewContactTypeWithDefaults() *ContactType`

NewContactTypeWithDefaults instantiates a new ContactType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContactType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *ContactType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContactType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContactType) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDefaultFlag

`func (o *ContactType) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ContactType) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ContactType) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ContactType) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ContactType) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ContactType) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetServiceAlertFlag

`func (o *ContactType) GetServiceAlertFlag() bool`

GetServiceAlertFlag returns the ServiceAlertFlag field if non-nil, zero value otherwise.

### GetServiceAlertFlagOk

`func (o *ContactType) GetServiceAlertFlagOk() (*bool, bool)`

GetServiceAlertFlagOk returns a tuple with the ServiceAlertFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAlertFlag

`func (o *ContactType) SetServiceAlertFlag(v bool)`

SetServiceAlertFlag sets ServiceAlertFlag field to given value.

### HasServiceAlertFlag

`func (o *ContactType) HasServiceAlertFlag() bool`

HasServiceAlertFlag returns a boolean if a field has been set.

### SetServiceAlertFlagNil

`func (o *ContactType) SetServiceAlertFlagNil(b bool)`

 SetServiceAlertFlagNil sets the value for ServiceAlertFlag to be an explicit nil

### UnsetServiceAlertFlag
`func (o *ContactType) UnsetServiceAlertFlag()`

UnsetServiceAlertFlag ensures that no value is present for ServiceAlertFlag, not even an explicit nil
### GetServiceAlertMessage

`func (o *ContactType) GetServiceAlertMessage() string`

GetServiceAlertMessage returns the ServiceAlertMessage field if non-nil, zero value otherwise.

### GetServiceAlertMessageOk

`func (o *ContactType) GetServiceAlertMessageOk() (*string, bool)`

GetServiceAlertMessageOk returns a tuple with the ServiceAlertMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAlertMessage

`func (o *ContactType) SetServiceAlertMessage(v string)`

SetServiceAlertMessage sets ServiceAlertMessage field to given value.

### HasServiceAlertMessage

`func (o *ContactType) HasServiceAlertMessage() bool`

HasServiceAlertMessage returns a boolean if a field has been set.

### GetInfo

`func (o *ContactType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ContactType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ContactType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ContactType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


