# PricingSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Companies** | Pointer to **[]int32** |  | [optional] 
**SetAllCompaniesFlag** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllCompaniesFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPricingSchedule

`func NewPricingSchedule(name string, ) *PricingSchedule`

NewPricingSchedule instantiates a new PricingSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingScheduleWithDefaults

`func NewPricingScheduleWithDefaults() *PricingSchedule`

NewPricingScheduleWithDefaults instantiates a new PricingSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PricingSchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PricingSchedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PricingSchedule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PricingSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PricingSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PricingSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PricingSchedule) SetName(v string)`

SetName sets Name field to given value.


### GetInactiveFlag

`func (o *PricingSchedule) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *PricingSchedule) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *PricingSchedule) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *PricingSchedule) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *PricingSchedule) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *PricingSchedule) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetDefaultFlag

`func (o *PricingSchedule) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *PricingSchedule) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *PricingSchedule) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *PricingSchedule) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *PricingSchedule) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *PricingSchedule) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetCurrency

`func (o *PricingSchedule) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PricingSchedule) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PricingSchedule) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PricingSchedule) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCompanies

`func (o *PricingSchedule) GetCompanies() []int32`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *PricingSchedule) GetCompaniesOk() (*[]int32, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *PricingSchedule) SetCompanies(v []int32)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *PricingSchedule) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### GetSetAllCompaniesFlag

`func (o *PricingSchedule) GetSetAllCompaniesFlag() bool`

GetSetAllCompaniesFlag returns the SetAllCompaniesFlag field if non-nil, zero value otherwise.

### GetSetAllCompaniesFlagOk

`func (o *PricingSchedule) GetSetAllCompaniesFlagOk() (*bool, bool)`

GetSetAllCompaniesFlagOk returns a tuple with the SetAllCompaniesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetAllCompaniesFlag

`func (o *PricingSchedule) SetSetAllCompaniesFlag(v bool)`

SetSetAllCompaniesFlag sets SetAllCompaniesFlag field to given value.

### HasSetAllCompaniesFlag

`func (o *PricingSchedule) HasSetAllCompaniesFlag() bool`

HasSetAllCompaniesFlag returns a boolean if a field has been set.

### SetSetAllCompaniesFlagNil

`func (o *PricingSchedule) SetSetAllCompaniesFlagNil(b bool)`

 SetSetAllCompaniesFlagNil sets the value for SetAllCompaniesFlag to be an explicit nil

### UnsetSetAllCompaniesFlag
`func (o *PricingSchedule) UnsetSetAllCompaniesFlag()`

UnsetSetAllCompaniesFlag ensures that no value is present for SetAllCompaniesFlag, not even an explicit nil
### GetRemoveAllCompaniesFlag

`func (o *PricingSchedule) GetRemoveAllCompaniesFlag() bool`

GetRemoveAllCompaniesFlag returns the RemoveAllCompaniesFlag field if non-nil, zero value otherwise.

### GetRemoveAllCompaniesFlagOk

`func (o *PricingSchedule) GetRemoveAllCompaniesFlagOk() (*bool, bool)`

GetRemoveAllCompaniesFlagOk returns a tuple with the RemoveAllCompaniesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllCompaniesFlag

`func (o *PricingSchedule) SetRemoveAllCompaniesFlag(v bool)`

SetRemoveAllCompaniesFlag sets RemoveAllCompaniesFlag field to given value.

### HasRemoveAllCompaniesFlag

`func (o *PricingSchedule) HasRemoveAllCompaniesFlag() bool`

HasRemoveAllCompaniesFlag returns a boolean if a field has been set.

### SetRemoveAllCompaniesFlagNil

`func (o *PricingSchedule) SetRemoveAllCompaniesFlagNil(b bool)`

 SetRemoveAllCompaniesFlagNil sets the value for RemoveAllCompaniesFlag to be an explicit nil

### UnsetRemoveAllCompaniesFlag
`func (o *PricingSchedule) UnsetRemoveAllCompaniesFlag()`

UnsetRemoveAllCompaniesFlag ensures that no value is present for RemoveAllCompaniesFlag, not even an explicit nil
### GetInfo

`func (o *PricingSchedule) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PricingSchedule) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PricingSchedule) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PricingSchedule) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


