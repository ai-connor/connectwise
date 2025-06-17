# DeliveryMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailFlag** | Pointer to **NullableBool** |  | [optional] 
**IntegrationEmailFlag** | Pointer to **NullableBool** |  | [optional] 
**IntegrationPrintFlag** | Pointer to **NullableBool** |  | [optional] 
**IntegrationActiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDeliveryMethod

`func NewDeliveryMethod(name string, ) *DeliveryMethod`

NewDeliveryMethod instantiates a new DeliveryMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryMethodWithDefaults

`func NewDeliveryMethodWithDefaults() *DeliveryMethod`

NewDeliveryMethodWithDefaults instantiates a new DeliveryMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliveryMethod) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryMethod) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryMethod) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeliveryMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeliveryMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeliveryMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeliveryMethod) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *DeliveryMethod) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *DeliveryMethod) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *DeliveryMethod) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *DeliveryMethod) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *DeliveryMethod) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *DeliveryMethod) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetEmailFlag

`func (o *DeliveryMethod) GetEmailFlag() bool`

GetEmailFlag returns the EmailFlag field if non-nil, zero value otherwise.

### GetEmailFlagOk

`func (o *DeliveryMethod) GetEmailFlagOk() (*bool, bool)`

GetEmailFlagOk returns a tuple with the EmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFlag

`func (o *DeliveryMethod) SetEmailFlag(v bool)`

SetEmailFlag sets EmailFlag field to given value.

### HasEmailFlag

`func (o *DeliveryMethod) HasEmailFlag() bool`

HasEmailFlag returns a boolean if a field has been set.

### SetEmailFlagNil

`func (o *DeliveryMethod) SetEmailFlagNil(b bool)`

 SetEmailFlagNil sets the value for EmailFlag to be an explicit nil

### UnsetEmailFlag
`func (o *DeliveryMethod) UnsetEmailFlag()`

UnsetEmailFlag ensures that no value is present for EmailFlag, not even an explicit nil
### GetIntegrationEmailFlag

`func (o *DeliveryMethod) GetIntegrationEmailFlag() bool`

GetIntegrationEmailFlag returns the IntegrationEmailFlag field if non-nil, zero value otherwise.

### GetIntegrationEmailFlagOk

`func (o *DeliveryMethod) GetIntegrationEmailFlagOk() (*bool, bool)`

GetIntegrationEmailFlagOk returns a tuple with the IntegrationEmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationEmailFlag

`func (o *DeliveryMethod) SetIntegrationEmailFlag(v bool)`

SetIntegrationEmailFlag sets IntegrationEmailFlag field to given value.

### HasIntegrationEmailFlag

`func (o *DeliveryMethod) HasIntegrationEmailFlag() bool`

HasIntegrationEmailFlag returns a boolean if a field has been set.

### SetIntegrationEmailFlagNil

`func (o *DeliveryMethod) SetIntegrationEmailFlagNil(b bool)`

 SetIntegrationEmailFlagNil sets the value for IntegrationEmailFlag to be an explicit nil

### UnsetIntegrationEmailFlag
`func (o *DeliveryMethod) UnsetIntegrationEmailFlag()`

UnsetIntegrationEmailFlag ensures that no value is present for IntegrationEmailFlag, not even an explicit nil
### GetIntegrationPrintFlag

`func (o *DeliveryMethod) GetIntegrationPrintFlag() bool`

GetIntegrationPrintFlag returns the IntegrationPrintFlag field if non-nil, zero value otherwise.

### GetIntegrationPrintFlagOk

`func (o *DeliveryMethod) GetIntegrationPrintFlagOk() (*bool, bool)`

GetIntegrationPrintFlagOk returns a tuple with the IntegrationPrintFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationPrintFlag

`func (o *DeliveryMethod) SetIntegrationPrintFlag(v bool)`

SetIntegrationPrintFlag sets IntegrationPrintFlag field to given value.

### HasIntegrationPrintFlag

`func (o *DeliveryMethod) HasIntegrationPrintFlag() bool`

HasIntegrationPrintFlag returns a boolean if a field has been set.

### SetIntegrationPrintFlagNil

`func (o *DeliveryMethod) SetIntegrationPrintFlagNil(b bool)`

 SetIntegrationPrintFlagNil sets the value for IntegrationPrintFlag to be an explicit nil

### UnsetIntegrationPrintFlag
`func (o *DeliveryMethod) UnsetIntegrationPrintFlag()`

UnsetIntegrationPrintFlag ensures that no value is present for IntegrationPrintFlag, not even an explicit nil
### GetIntegrationActiveFlag

`func (o *DeliveryMethod) GetIntegrationActiveFlag() bool`

GetIntegrationActiveFlag returns the IntegrationActiveFlag field if non-nil, zero value otherwise.

### GetIntegrationActiveFlagOk

`func (o *DeliveryMethod) GetIntegrationActiveFlagOk() (*bool, bool)`

GetIntegrationActiveFlagOk returns a tuple with the IntegrationActiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationActiveFlag

`func (o *DeliveryMethod) SetIntegrationActiveFlag(v bool)`

SetIntegrationActiveFlag sets IntegrationActiveFlag field to given value.

### HasIntegrationActiveFlag

`func (o *DeliveryMethod) HasIntegrationActiveFlag() bool`

HasIntegrationActiveFlag returns a boolean if a field has been set.

### SetIntegrationActiveFlagNil

`func (o *DeliveryMethod) SetIntegrationActiveFlagNil(b bool)`

 SetIntegrationActiveFlagNil sets the value for IntegrationActiveFlag to be an explicit nil

### UnsetIntegrationActiveFlag
`func (o *DeliveryMethod) UnsetIntegrationActiveFlag()`

UnsetIntegrationActiveFlag ensures that no value is present for IntegrationActiveFlag, not even an explicit nil
### GetInfo

`func (o *DeliveryMethod) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DeliveryMethod) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DeliveryMethod) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *DeliveryMethod) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


