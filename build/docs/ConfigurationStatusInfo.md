# ConfigurationStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConfigurationStatusInfo

`func NewConfigurationStatusInfo() *ConfigurationStatusInfo`

NewConfigurationStatusInfo instantiates a new ConfigurationStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationStatusInfoWithDefaults

`func NewConfigurationStatusInfoWithDefaults() *ConfigurationStatusInfo`

NewConfigurationStatusInfoWithDefaults instantiates a new ConfigurationStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationStatusInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationStatusInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationStatusInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationStatusInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigurationStatusInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigurationStatusInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigurationStatusInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigurationStatusInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClosedFlag

`func (o *ConfigurationStatusInfo) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *ConfigurationStatusInfo) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *ConfigurationStatusInfo) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *ConfigurationStatusInfo) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *ConfigurationStatusInfo) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *ConfigurationStatusInfo) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetDefaultFlag

`func (o *ConfigurationStatusInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ConfigurationStatusInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ConfigurationStatusInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ConfigurationStatusInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ConfigurationStatusInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ConfigurationStatusInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInfo

`func (o *ConfigurationStatusInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConfigurationStatusInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConfigurationStatusInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConfigurationStatusInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


