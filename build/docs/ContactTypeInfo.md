# ContactTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceAlertFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceAlertMessage** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewContactTypeInfo

`func NewContactTypeInfo() *ContactTypeInfo`

NewContactTypeInfo instantiates a new ContactTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactTypeInfoWithDefaults

`func NewContactTypeInfoWithDefaults() *ContactTypeInfo`

NewContactTypeInfoWithDefaults instantiates a new ContactTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactTypeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactTypeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactTypeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContactTypeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *ContactTypeInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContactTypeInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContactTypeInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContactTypeInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *ContactTypeInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ContactTypeInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ContactTypeInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ContactTypeInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ContactTypeInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ContactTypeInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetServiceAlertFlag

`func (o *ContactTypeInfo) GetServiceAlertFlag() bool`

GetServiceAlertFlag returns the ServiceAlertFlag field if non-nil, zero value otherwise.

### GetServiceAlertFlagOk

`func (o *ContactTypeInfo) GetServiceAlertFlagOk() (*bool, bool)`

GetServiceAlertFlagOk returns a tuple with the ServiceAlertFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAlertFlag

`func (o *ContactTypeInfo) SetServiceAlertFlag(v bool)`

SetServiceAlertFlag sets ServiceAlertFlag field to given value.

### HasServiceAlertFlag

`func (o *ContactTypeInfo) HasServiceAlertFlag() bool`

HasServiceAlertFlag returns a boolean if a field has been set.

### SetServiceAlertFlagNil

`func (o *ContactTypeInfo) SetServiceAlertFlagNil(b bool)`

 SetServiceAlertFlagNil sets the value for ServiceAlertFlag to be an explicit nil

### UnsetServiceAlertFlag
`func (o *ContactTypeInfo) UnsetServiceAlertFlag()`

UnsetServiceAlertFlag ensures that no value is present for ServiceAlertFlag, not even an explicit nil
### GetServiceAlertMessage

`func (o *ContactTypeInfo) GetServiceAlertMessage() string`

GetServiceAlertMessage returns the ServiceAlertMessage field if non-nil, zero value otherwise.

### GetServiceAlertMessageOk

`func (o *ContactTypeInfo) GetServiceAlertMessageOk() (*string, bool)`

GetServiceAlertMessageOk returns a tuple with the ServiceAlertMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAlertMessage

`func (o *ContactTypeInfo) SetServiceAlertMessage(v string)`

SetServiceAlertMessage sets ServiceAlertMessage field to given value.

### HasServiceAlertMessage

`func (o *ContactTypeInfo) HasServiceAlertMessage() bool`

HasServiceAlertMessage returns a boolean if a field has been set.

### GetInfo

`func (o *ContactTypeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ContactTypeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ContactTypeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ContactTypeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


