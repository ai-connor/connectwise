# ProductType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**TypeXref** | Pointer to **NullableString** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProductType

`func NewProductType(name string, ) *ProductType`

NewProductType instantiates a new ProductType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTypeWithDefaults

`func NewProductTypeWithDefaults() *ProductType`

NewProductTypeWithDefaults instantiates a new ProductType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProductType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductType) SetName(v string)`

SetName sets Name field to given value.


### GetInactiveFlag

`func (o *ProductType) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ProductType) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ProductType) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ProductType) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ProductType) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ProductType) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetTypeXref

`func (o *ProductType) GetTypeXref() string`

GetTypeXref returns the TypeXref field if non-nil, zero value otherwise.

### GetTypeXrefOk

`func (o *ProductType) GetTypeXrefOk() (*string, bool)`

GetTypeXrefOk returns a tuple with the TypeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeXref

`func (o *ProductType) SetTypeXref(v string)`

SetTypeXref sets TypeXref field to given value.

### HasTypeXref

`func (o *ProductType) HasTypeXref() bool`

HasTypeXref returns a boolean if a field has been set.

### SetTypeXrefNil

`func (o *ProductType) SetTypeXrefNil(b bool)`

 SetTypeXrefNil sets the value for TypeXref to be an explicit nil

### UnsetTypeXref
`func (o *ProductType) UnsetTypeXref()`

UnsetTypeXref ensures that no value is present for TypeXref, not even an explicit nil
### GetDefaultFlag

`func (o *ProductType) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ProductType) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ProductType) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ProductType) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ProductType) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ProductType) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInfo

`func (o *ProductType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProductType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProductType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProductType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


