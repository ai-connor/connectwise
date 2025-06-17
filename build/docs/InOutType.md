# InOutType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Description** | **string** |  Max length: 30; | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewInOutType

`func NewInOutType(description string, ) *InOutType`

NewInOutType instantiates a new InOutType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInOutTypeWithDefaults

`func NewInOutTypeWithDefaults() *InOutType`

NewInOutTypeWithDefaults instantiates a new InOutType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InOutType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InOutType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InOutType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InOutType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *InOutType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InOutType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InOutType) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInfo

`func (o *InOutType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InOutType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InOutType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InOutType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


