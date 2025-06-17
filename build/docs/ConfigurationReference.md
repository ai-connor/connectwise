# ConfigurationReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**DeviceIdentifier** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConfigurationReference

`func NewConfigurationReference() *ConfigurationReference`

NewConfigurationReference instantiates a new ConfigurationReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationReferenceWithDefaults

`func NewConfigurationReferenceWithDefaults() *ConfigurationReference`

NewConfigurationReferenceWithDefaults instantiates a new ConfigurationReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ConfigurationReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ConfigurationReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDeviceIdentifier

`func (o *ConfigurationReference) GetDeviceIdentifier() string`

GetDeviceIdentifier returns the DeviceIdentifier field if non-nil, zero value otherwise.

### GetDeviceIdentifierOk

`func (o *ConfigurationReference) GetDeviceIdentifierOk() (*string, bool)`

GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdentifier

`func (o *ConfigurationReference) SetDeviceIdentifier(v string)`

SetDeviceIdentifier sets DeviceIdentifier field to given value.

### HasDeviceIdentifier

`func (o *ConfigurationReference) HasDeviceIdentifier() bool`

HasDeviceIdentifier returns a boolean if a field has been set.

### GetInfo

`func (o *ConfigurationReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConfigurationReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConfigurationReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConfigurationReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


