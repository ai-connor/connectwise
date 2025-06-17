# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**PriceLevelXref** | Pointer to **string** |  Max length: 10; | [optional] 
**IntegrationXref** | Pointer to **string** |  Max length: 50; | [optional] 
**LocationIds** | Pointer to **[]int32** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**AddAllLocations** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllLocations** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCategory

`func NewCategory(name string, ) *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Category) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Category) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Category) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Category) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Category) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Category) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Category) SetName(v string)`

SetName sets Name field to given value.


### GetInactiveFlag

`func (o *Category) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *Category) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *Category) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *Category) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *Category) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *Category) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetPriceLevelXref

`func (o *Category) GetPriceLevelXref() string`

GetPriceLevelXref returns the PriceLevelXref field if non-nil, zero value otherwise.

### GetPriceLevelXrefOk

`func (o *Category) GetPriceLevelXrefOk() (*string, bool)`

GetPriceLevelXrefOk returns a tuple with the PriceLevelXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevelXref

`func (o *Category) SetPriceLevelXref(v string)`

SetPriceLevelXref sets PriceLevelXref field to given value.

### HasPriceLevelXref

`func (o *Category) HasPriceLevelXref() bool`

HasPriceLevelXref returns a boolean if a field has been set.

### GetIntegrationXref

`func (o *Category) GetIntegrationXref() string`

GetIntegrationXref returns the IntegrationXref field if non-nil, zero value otherwise.

### GetIntegrationXrefOk

`func (o *Category) GetIntegrationXrefOk() (*string, bool)`

GetIntegrationXrefOk returns a tuple with the IntegrationXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationXref

`func (o *Category) SetIntegrationXref(v string)`

SetIntegrationXref sets IntegrationXref field to given value.

### HasIntegrationXref

`func (o *Category) HasIntegrationXref() bool`

HasIntegrationXref returns a boolean if a field has been set.

### GetLocationIds

`func (o *Category) GetLocationIds() []int32`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *Category) GetLocationIdsOk() (*[]int32, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *Category) SetLocationIds(v []int32)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *Category) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *Category) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *Category) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *Category) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *Category) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *Category) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *Category) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetAddAllLocations

`func (o *Category) GetAddAllLocations() bool`

GetAddAllLocations returns the AddAllLocations field if non-nil, zero value otherwise.

### GetAddAllLocationsOk

`func (o *Category) GetAddAllLocationsOk() (*bool, bool)`

GetAddAllLocationsOk returns a tuple with the AddAllLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllLocations

`func (o *Category) SetAddAllLocations(v bool)`

SetAddAllLocations sets AddAllLocations field to given value.

### HasAddAllLocations

`func (o *Category) HasAddAllLocations() bool`

HasAddAllLocations returns a boolean if a field has been set.

### SetAddAllLocationsNil

`func (o *Category) SetAddAllLocationsNil(b bool)`

 SetAddAllLocationsNil sets the value for AddAllLocations to be an explicit nil

### UnsetAddAllLocations
`func (o *Category) UnsetAddAllLocations()`

UnsetAddAllLocations ensures that no value is present for AddAllLocations, not even an explicit nil
### GetRemoveAllLocations

`func (o *Category) GetRemoveAllLocations() bool`

GetRemoveAllLocations returns the RemoveAllLocations field if non-nil, zero value otherwise.

### GetRemoveAllLocationsOk

`func (o *Category) GetRemoveAllLocationsOk() (*bool, bool)`

GetRemoveAllLocationsOk returns a tuple with the RemoveAllLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllLocations

`func (o *Category) SetRemoveAllLocations(v bool)`

SetRemoveAllLocations sets RemoveAllLocations field to given value.

### HasRemoveAllLocations

`func (o *Category) HasRemoveAllLocations() bool`

HasRemoveAllLocations returns a boolean if a field has been set.

### SetRemoveAllLocationsNil

`func (o *Category) SetRemoveAllLocationsNil(b bool)`

 SetRemoveAllLocationsNil sets the value for RemoveAllLocations to be an explicit nil

### UnsetRemoveAllLocations
`func (o *Category) UnsetRemoveAllLocations()`

UnsetRemoveAllLocations ensures that no value is present for RemoveAllLocations, not even an explicit nil
### GetInfo

`func (o *Category) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Category) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Category) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Category) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


