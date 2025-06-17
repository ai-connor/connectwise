# AddressFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Format** | **string** |  Max length: 250; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**CountryIds** | Pointer to **[]int32** |  | [optional] 
**AddAllCountries** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllCountries** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAddressFormat

`func NewAddressFormat(name string, format string, ) *AddressFormat`

NewAddressFormat instantiates a new AddressFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressFormatWithDefaults

`func NewAddressFormatWithDefaults() *AddressFormat`

NewAddressFormatWithDefaults instantiates a new AddressFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddressFormat) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressFormat) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressFormat) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AddressFormat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddressFormat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressFormat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressFormat) SetName(v string)`

SetName sets Name field to given value.


### GetFormat

`func (o *AddressFormat) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AddressFormat) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AddressFormat) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetDefaultFlag

`func (o *AddressFormat) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *AddressFormat) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *AddressFormat) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *AddressFormat) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *AddressFormat) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *AddressFormat) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetCountryIds

`func (o *AddressFormat) GetCountryIds() []int32`

GetCountryIds returns the CountryIds field if non-nil, zero value otherwise.

### GetCountryIdsOk

`func (o *AddressFormat) GetCountryIdsOk() (*[]int32, bool)`

GetCountryIdsOk returns a tuple with the CountryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIds

`func (o *AddressFormat) SetCountryIds(v []int32)`

SetCountryIds sets CountryIds field to given value.

### HasCountryIds

`func (o *AddressFormat) HasCountryIds() bool`

HasCountryIds returns a boolean if a field has been set.

### GetAddAllCountries

`func (o *AddressFormat) GetAddAllCountries() bool`

GetAddAllCountries returns the AddAllCountries field if non-nil, zero value otherwise.

### GetAddAllCountriesOk

`func (o *AddressFormat) GetAddAllCountriesOk() (*bool, bool)`

GetAddAllCountriesOk returns a tuple with the AddAllCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllCountries

`func (o *AddressFormat) SetAddAllCountries(v bool)`

SetAddAllCountries sets AddAllCountries field to given value.

### HasAddAllCountries

`func (o *AddressFormat) HasAddAllCountries() bool`

HasAddAllCountries returns a boolean if a field has been set.

### SetAddAllCountriesNil

`func (o *AddressFormat) SetAddAllCountriesNil(b bool)`

 SetAddAllCountriesNil sets the value for AddAllCountries to be an explicit nil

### UnsetAddAllCountries
`func (o *AddressFormat) UnsetAddAllCountries()`

UnsetAddAllCountries ensures that no value is present for AddAllCountries, not even an explicit nil
### GetRemoveAllCountries

`func (o *AddressFormat) GetRemoveAllCountries() bool`

GetRemoveAllCountries returns the RemoveAllCountries field if non-nil, zero value otherwise.

### GetRemoveAllCountriesOk

`func (o *AddressFormat) GetRemoveAllCountriesOk() (*bool, bool)`

GetRemoveAllCountriesOk returns a tuple with the RemoveAllCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllCountries

`func (o *AddressFormat) SetRemoveAllCountries(v bool)`

SetRemoveAllCountries sets RemoveAllCountries field to given value.

### HasRemoveAllCountries

`func (o *AddressFormat) HasRemoveAllCountries() bool`

HasRemoveAllCountries returns a boolean if a field has been set.

### SetRemoveAllCountriesNil

`func (o *AddressFormat) SetRemoveAllCountriesNil(b bool)`

 SetRemoveAllCountriesNil sets the value for RemoveAllCountries to be an explicit nil

### UnsetRemoveAllCountries
`func (o *AddressFormat) UnsetRemoveAllCountries()`

UnsetRemoveAllCountries ensures that no value is present for RemoveAllCountries, not even an explicit nil
### GetInfo

`func (o *AddressFormat) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AddressFormat) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AddressFormat) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AddressFormat) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


