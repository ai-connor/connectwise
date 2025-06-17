# SubCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**IntegrationXref** | Pointer to **string** |  Max length: 50; | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Category** | [**ProductCategoryReference**](ProductCategoryReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSubCategory

`func NewSubCategory(name string, category ProductCategoryReference, ) *SubCategory`

NewSubCategory instantiates a new SubCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubCategoryWithDefaults

`func NewSubCategoryWithDefaults() *SubCategory`

NewSubCategoryWithDefaults instantiates a new SubCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SubCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubCategory) SetName(v string)`

SetName sets Name field to given value.


### GetInactiveFlag

`func (o *SubCategory) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *SubCategory) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *SubCategory) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *SubCategory) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *SubCategory) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *SubCategory) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetIntegrationXref

`func (o *SubCategory) GetIntegrationXref() string`

GetIntegrationXref returns the IntegrationXref field if non-nil, zero value otherwise.

### GetIntegrationXrefOk

`func (o *SubCategory) GetIntegrationXrefOk() (*string, bool)`

GetIntegrationXrefOk returns a tuple with the IntegrationXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationXref

`func (o *SubCategory) SetIntegrationXref(v string)`

SetIntegrationXref sets IntegrationXref field to given value.

### HasIntegrationXref

`func (o *SubCategory) HasIntegrationXref() bool`

HasIntegrationXref returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *SubCategory) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *SubCategory) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *SubCategory) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *SubCategory) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *SubCategory) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *SubCategory) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetCategory

`func (o *SubCategory) GetCategory() ProductCategoryReference`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SubCategory) GetCategoryOk() (*ProductCategoryReference, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SubCategory) SetCategory(v ProductCategoryReference)`

SetCategory sets Category field to given value.


### GetInfo

`func (o *SubCategory) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SubCategory) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SubCategory) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SubCategory) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


