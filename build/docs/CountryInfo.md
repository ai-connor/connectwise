# CountryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**CityCaption** | Pointer to **string** |  | [optional] 
**StateCaption** | Pointer to **string** |  | [optional] 
**ZipCaption** | Pointer to **string** |  | [optional] 
**DialingPrefix** | Pointer to **string** |  | [optional] 
**LocalizationCaptionOne** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCountryInfo

`func NewCountryInfo() *CountryInfo`

NewCountryInfo instantiates a new CountryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryInfoWithDefaults

`func NewCountryInfoWithDefaults() *CountryInfo`

NewCountryInfoWithDefaults instantiates a new CountryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CountryInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CountryInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CountryInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CountryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CountryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CountryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CountryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CountryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *CountryInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *CountryInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *CountryInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *CountryInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *CountryInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *CountryInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetCityCaption

`func (o *CountryInfo) GetCityCaption() string`

GetCityCaption returns the CityCaption field if non-nil, zero value otherwise.

### GetCityCaptionOk

`func (o *CountryInfo) GetCityCaptionOk() (*string, bool)`

GetCityCaptionOk returns a tuple with the CityCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityCaption

`func (o *CountryInfo) SetCityCaption(v string)`

SetCityCaption sets CityCaption field to given value.

### HasCityCaption

`func (o *CountryInfo) HasCityCaption() bool`

HasCityCaption returns a boolean if a field has been set.

### GetStateCaption

`func (o *CountryInfo) GetStateCaption() string`

GetStateCaption returns the StateCaption field if non-nil, zero value otherwise.

### GetStateCaptionOk

`func (o *CountryInfo) GetStateCaptionOk() (*string, bool)`

GetStateCaptionOk returns a tuple with the StateCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCaption

`func (o *CountryInfo) SetStateCaption(v string)`

SetStateCaption sets StateCaption field to given value.

### HasStateCaption

`func (o *CountryInfo) HasStateCaption() bool`

HasStateCaption returns a boolean if a field has been set.

### GetZipCaption

`func (o *CountryInfo) GetZipCaption() string`

GetZipCaption returns the ZipCaption field if non-nil, zero value otherwise.

### GetZipCaptionOk

`func (o *CountryInfo) GetZipCaptionOk() (*string, bool)`

GetZipCaptionOk returns a tuple with the ZipCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCaption

`func (o *CountryInfo) SetZipCaption(v string)`

SetZipCaption sets ZipCaption field to given value.

### HasZipCaption

`func (o *CountryInfo) HasZipCaption() bool`

HasZipCaption returns a boolean if a field has been set.

### GetDialingPrefix

`func (o *CountryInfo) GetDialingPrefix() string`

GetDialingPrefix returns the DialingPrefix field if non-nil, zero value otherwise.

### GetDialingPrefixOk

`func (o *CountryInfo) GetDialingPrefixOk() (*string, bool)`

GetDialingPrefixOk returns a tuple with the DialingPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialingPrefix

`func (o *CountryInfo) SetDialingPrefix(v string)`

SetDialingPrefix sets DialingPrefix field to given value.

### HasDialingPrefix

`func (o *CountryInfo) HasDialingPrefix() bool`

HasDialingPrefix returns a boolean if a field has been set.

### GetLocalizationCaptionOne

`func (o *CountryInfo) GetLocalizationCaptionOne() string`

GetLocalizationCaptionOne returns the LocalizationCaptionOne field if non-nil, zero value otherwise.

### GetLocalizationCaptionOneOk

`func (o *CountryInfo) GetLocalizationCaptionOneOk() (*string, bool)`

GetLocalizationCaptionOneOk returns a tuple with the LocalizationCaptionOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizationCaptionOne

`func (o *CountryInfo) SetLocalizationCaptionOne(v string)`

SetLocalizationCaptionOne sets LocalizationCaptionOne field to given value.

### HasLocalizationCaptionOne

`func (o *CountryInfo) HasLocalizationCaptionOne() bool`

HasLocalizationCaptionOne returns a boolean if a field has been set.

### GetCurrency

`func (o *CountryInfo) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CountryInfo) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CountryInfo) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CountryInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInfo

`func (o *CountryInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CountryInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CountryInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CountryInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


