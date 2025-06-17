# ServiceTemplateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TemplateFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceTemplateInfo

`func NewServiceTemplateInfo() *ServiceTemplateInfo`

NewServiceTemplateInfo instantiates a new ServiceTemplateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTemplateInfoWithDefaults

`func NewServiceTemplateInfoWithDefaults() *ServiceTemplateInfo`

NewServiceTemplateInfoWithDefaults instantiates a new ServiceTemplateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceTemplateInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceTemplateInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceTemplateInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceTemplateInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceTemplateInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceTemplateInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceTemplateInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceTemplateInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateFlag

`func (o *ServiceTemplateInfo) GetTemplateFlag() bool`

GetTemplateFlag returns the TemplateFlag field if non-nil, zero value otherwise.

### GetTemplateFlagOk

`func (o *ServiceTemplateInfo) GetTemplateFlagOk() (*bool, bool)`

GetTemplateFlagOk returns a tuple with the TemplateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateFlag

`func (o *ServiceTemplateInfo) SetTemplateFlag(v bool)`

SetTemplateFlag sets TemplateFlag field to given value.

### HasTemplateFlag

`func (o *ServiceTemplateInfo) HasTemplateFlag() bool`

HasTemplateFlag returns a boolean if a field has been set.

### SetTemplateFlagNil

`func (o *ServiceTemplateInfo) SetTemplateFlagNil(b bool)`

 SetTemplateFlagNil sets the value for TemplateFlag to be an explicit nil

### UnsetTemplateFlag
`func (o *ServiceTemplateInfo) UnsetTemplateFlag()`

UnsetTemplateFlag ensures that no value is present for TemplateFlag, not even an explicit nil
### GetInfo

`func (o *ServiceTemplateInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceTemplateInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceTemplateInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceTemplateInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


