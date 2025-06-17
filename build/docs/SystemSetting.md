# SystemSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Value** | **string** |  | 
**ValueType** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSystemSetting

`func NewSystemSetting(value string, ) *SystemSetting`

NewSystemSetting instantiates a new SystemSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemSettingWithDefaults

`func NewSystemSettingWithDefaults() *SystemSetting`

NewSystemSettingWithDefaults instantiates a new SystemSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemSetting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SystemSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *SystemSetting) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemSetting) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemSetting) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemSetting) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValue

`func (o *SystemSetting) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SystemSetting) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SystemSetting) SetValue(v string)`

SetValue sets Value field to given value.


### GetValueType

`func (o *SystemSetting) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *SystemSetting) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *SystemSetting) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *SystemSetting) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetInfo

`func (o *SystemSetting) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SystemSetting) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SystemSetting) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SystemSetting) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


