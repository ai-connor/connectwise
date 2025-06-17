# Country

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Currency** | [**CurrencyReference**](CurrencyReference.md) |  | 
**CityCaption** | Pointer to **string** |  Max length: 25; | [optional] 
**StateCaption** | Pointer to **string** |  Max length: 25; | [optional] 
**ZipCaption** | Pointer to **string** |  Max length: 25; | [optional] 
**ZipMinimumLength** | Pointer to **NullableInt32** |  | [optional] 
**DialingPrefix** | Pointer to **string** |  Max length: 5; | [optional] 
**AddressFormat** | Pointer to [**AddressFormatReference**](AddressFormatReference.md) |  | [optional] 
**CountryCode** | Pointer to **string** |  Max length: 2; | [optional] 
**CoreEntityCountryCode** | Pointer to **NullableString** |  | [optional] 
**LocalizationCaptionOne** | Pointer to **string** |  Max length: 25; | [optional] 
**LocalizationValueOne** | Pointer to **string** |  Max length: 50; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCountry

`func NewCountry(name string, currency CurrencyReference, ) *Country`

NewCountry instantiates a new Country object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryWithDefaults

`func NewCountryWithDefaults() *Country`

NewCountryWithDefaults instantiates a new Country object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Country) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Country) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Country) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Country) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Country) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Country) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Country) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *Country) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *Country) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *Country) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *Country) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *Country) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *Country) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetCurrency

`func (o *Country) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Country) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Country) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.


### GetCityCaption

`func (o *Country) GetCityCaption() string`

GetCityCaption returns the CityCaption field if non-nil, zero value otherwise.

### GetCityCaptionOk

`func (o *Country) GetCityCaptionOk() (*string, bool)`

GetCityCaptionOk returns a tuple with the CityCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityCaption

`func (o *Country) SetCityCaption(v string)`

SetCityCaption sets CityCaption field to given value.

### HasCityCaption

`func (o *Country) HasCityCaption() bool`

HasCityCaption returns a boolean if a field has been set.

### GetStateCaption

`func (o *Country) GetStateCaption() string`

GetStateCaption returns the StateCaption field if non-nil, zero value otherwise.

### GetStateCaptionOk

`func (o *Country) GetStateCaptionOk() (*string, bool)`

GetStateCaptionOk returns a tuple with the StateCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCaption

`func (o *Country) SetStateCaption(v string)`

SetStateCaption sets StateCaption field to given value.

### HasStateCaption

`func (o *Country) HasStateCaption() bool`

HasStateCaption returns a boolean if a field has been set.

### GetZipCaption

`func (o *Country) GetZipCaption() string`

GetZipCaption returns the ZipCaption field if non-nil, zero value otherwise.

### GetZipCaptionOk

`func (o *Country) GetZipCaptionOk() (*string, bool)`

GetZipCaptionOk returns a tuple with the ZipCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCaption

`func (o *Country) SetZipCaption(v string)`

SetZipCaption sets ZipCaption field to given value.

### HasZipCaption

`func (o *Country) HasZipCaption() bool`

HasZipCaption returns a boolean if a field has been set.

### GetZipMinimumLength

`func (o *Country) GetZipMinimumLength() int32`

GetZipMinimumLength returns the ZipMinimumLength field if non-nil, zero value otherwise.

### GetZipMinimumLengthOk

`func (o *Country) GetZipMinimumLengthOk() (*int32, bool)`

GetZipMinimumLengthOk returns a tuple with the ZipMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipMinimumLength

`func (o *Country) SetZipMinimumLength(v int32)`

SetZipMinimumLength sets ZipMinimumLength field to given value.

### HasZipMinimumLength

`func (o *Country) HasZipMinimumLength() bool`

HasZipMinimumLength returns a boolean if a field has been set.

### SetZipMinimumLengthNil

`func (o *Country) SetZipMinimumLengthNil(b bool)`

 SetZipMinimumLengthNil sets the value for ZipMinimumLength to be an explicit nil

### UnsetZipMinimumLength
`func (o *Country) UnsetZipMinimumLength()`

UnsetZipMinimumLength ensures that no value is present for ZipMinimumLength, not even an explicit nil
### GetDialingPrefix

`func (o *Country) GetDialingPrefix() string`

GetDialingPrefix returns the DialingPrefix field if non-nil, zero value otherwise.

### GetDialingPrefixOk

`func (o *Country) GetDialingPrefixOk() (*string, bool)`

GetDialingPrefixOk returns a tuple with the DialingPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialingPrefix

`func (o *Country) SetDialingPrefix(v string)`

SetDialingPrefix sets DialingPrefix field to given value.

### HasDialingPrefix

`func (o *Country) HasDialingPrefix() bool`

HasDialingPrefix returns a boolean if a field has been set.

### GetAddressFormat

`func (o *Country) GetAddressFormat() AddressFormatReference`

GetAddressFormat returns the AddressFormat field if non-nil, zero value otherwise.

### GetAddressFormatOk

`func (o *Country) GetAddressFormatOk() (*AddressFormatReference, bool)`

GetAddressFormatOk returns a tuple with the AddressFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFormat

`func (o *Country) SetAddressFormat(v AddressFormatReference)`

SetAddressFormat sets AddressFormat field to given value.

### HasAddressFormat

`func (o *Country) HasAddressFormat() bool`

HasAddressFormat returns a boolean if a field has been set.

### GetCountryCode

`func (o *Country) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Country) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Country) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Country) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCoreEntityCountryCode

`func (o *Country) GetCoreEntityCountryCode() string`

GetCoreEntityCountryCode returns the CoreEntityCountryCode field if non-nil, zero value otherwise.

### GetCoreEntityCountryCodeOk

`func (o *Country) GetCoreEntityCountryCodeOk() (*string, bool)`

GetCoreEntityCountryCodeOk returns a tuple with the CoreEntityCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreEntityCountryCode

`func (o *Country) SetCoreEntityCountryCode(v string)`

SetCoreEntityCountryCode sets CoreEntityCountryCode field to given value.

### HasCoreEntityCountryCode

`func (o *Country) HasCoreEntityCountryCode() bool`

HasCoreEntityCountryCode returns a boolean if a field has been set.

### SetCoreEntityCountryCodeNil

`func (o *Country) SetCoreEntityCountryCodeNil(b bool)`

 SetCoreEntityCountryCodeNil sets the value for CoreEntityCountryCode to be an explicit nil

### UnsetCoreEntityCountryCode
`func (o *Country) UnsetCoreEntityCountryCode()`

UnsetCoreEntityCountryCode ensures that no value is present for CoreEntityCountryCode, not even an explicit nil
### GetLocalizationCaptionOne

`func (o *Country) GetLocalizationCaptionOne() string`

GetLocalizationCaptionOne returns the LocalizationCaptionOne field if non-nil, zero value otherwise.

### GetLocalizationCaptionOneOk

`func (o *Country) GetLocalizationCaptionOneOk() (*string, bool)`

GetLocalizationCaptionOneOk returns a tuple with the LocalizationCaptionOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizationCaptionOne

`func (o *Country) SetLocalizationCaptionOne(v string)`

SetLocalizationCaptionOne sets LocalizationCaptionOne field to given value.

### HasLocalizationCaptionOne

`func (o *Country) HasLocalizationCaptionOne() bool`

HasLocalizationCaptionOne returns a boolean if a field has been set.

### GetLocalizationValueOne

`func (o *Country) GetLocalizationValueOne() string`

GetLocalizationValueOne returns the LocalizationValueOne field if non-nil, zero value otherwise.

### GetLocalizationValueOneOk

`func (o *Country) GetLocalizationValueOneOk() (*string, bool)`

GetLocalizationValueOneOk returns a tuple with the LocalizationValueOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizationValueOne

`func (o *Country) SetLocalizationValueOne(v string)`

SetLocalizationValueOne sets LocalizationValueOne field to given value.

### HasLocalizationValueOne

`func (o *Country) HasLocalizationValueOne() bool`

HasLocalizationValueOne returns a boolean if a field has been set.

### GetInfo

`func (o *Country) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Country) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Country) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Country) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


