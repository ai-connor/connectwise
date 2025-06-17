# Classification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Multiplier** | Pointer to **NullableFloat64** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**CompanyFlag** | Pointer to **NullableBool** |  | [optional] 
**EmployeeFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewClassification

`func NewClassification() *Classification`

NewClassification instantiates a new Classification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationWithDefaults

`func NewClassificationWithDefaults() *Classification`

NewClassificationWithDefaults instantiates a new Classification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Classification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Classification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Classification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Classification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Classification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Classification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Classification) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Classification) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMultiplier

`func (o *Classification) GetMultiplier() float64`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *Classification) GetMultiplierOk() (*float64, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *Classification) SetMultiplier(v float64)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *Classification) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### SetMultiplierNil

`func (o *Classification) SetMultiplierNil(b bool)`

 SetMultiplierNil sets the value for Multiplier to be an explicit nil

### UnsetMultiplier
`func (o *Classification) UnsetMultiplier()`

UnsetMultiplier ensures that no value is present for Multiplier, not even an explicit nil
### GetDefaultFlag

`func (o *Classification) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *Classification) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *Classification) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *Classification) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *Classification) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *Classification) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetCompanyFlag

`func (o *Classification) GetCompanyFlag() bool`

GetCompanyFlag returns the CompanyFlag field if non-nil, zero value otherwise.

### GetCompanyFlagOk

`func (o *Classification) GetCompanyFlagOk() (*bool, bool)`

GetCompanyFlagOk returns a tuple with the CompanyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyFlag

`func (o *Classification) SetCompanyFlag(v bool)`

SetCompanyFlag sets CompanyFlag field to given value.

### HasCompanyFlag

`func (o *Classification) HasCompanyFlag() bool`

HasCompanyFlag returns a boolean if a field has been set.

### SetCompanyFlagNil

`func (o *Classification) SetCompanyFlagNil(b bool)`

 SetCompanyFlagNil sets the value for CompanyFlag to be an explicit nil

### UnsetCompanyFlag
`func (o *Classification) UnsetCompanyFlag()`

UnsetCompanyFlag ensures that no value is present for CompanyFlag, not even an explicit nil
### GetEmployeeFlag

`func (o *Classification) GetEmployeeFlag() bool`

GetEmployeeFlag returns the EmployeeFlag field if non-nil, zero value otherwise.

### GetEmployeeFlagOk

`func (o *Classification) GetEmployeeFlagOk() (*bool, bool)`

GetEmployeeFlagOk returns a tuple with the EmployeeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeFlag

`func (o *Classification) SetEmployeeFlag(v bool)`

SetEmployeeFlag sets EmployeeFlag field to given value.

### HasEmployeeFlag

`func (o *Classification) HasEmployeeFlag() bool`

HasEmployeeFlag returns a boolean if a field has been set.

### SetEmployeeFlagNil

`func (o *Classification) SetEmployeeFlagNil(b bool)`

 SetEmployeeFlagNil sets the value for EmployeeFlag to be an explicit nil

### UnsetEmployeeFlag
`func (o *Classification) UnsetEmployeeFlag()`

UnsetEmployeeFlag ensures that no value is present for EmployeeFlag, not even an explicit nil
### GetInfo

`func (o *Classification) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Classification) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Classification) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Classification) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


