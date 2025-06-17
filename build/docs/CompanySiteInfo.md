# CompanySiteInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**StateReference** | Pointer to [**StateReference**](StateReference.md) |  | [optional] 
**Country** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**Zip** | Pointer to **string** |  | [optional] 
**AddressFormat** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultShippingFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultBillingFlag** | Pointer to **NullableBool** |  | [optional] 
**PrimaryAddressFlag** | Pointer to **NullableBool** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**PhoneNumberExt** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanySiteInfo

`func NewCompanySiteInfo() *CompanySiteInfo`

NewCompanySiteInfo instantiates a new CompanySiteInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanySiteInfoWithDefaults

`func NewCompanySiteInfoWithDefaults() *CompanySiteInfo`

NewCompanySiteInfoWithDefaults instantiates a new CompanySiteInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanySiteInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanySiteInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanySiteInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanySiteInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CompanySiteInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanySiteInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanySiteInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanySiteInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *CompanySiteInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CompanySiteInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CompanySiteInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CompanySiteInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *CompanySiteInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CompanySiteInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CompanySiteInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CompanySiteInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *CompanySiteInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CompanySiteInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CompanySiteInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CompanySiteInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStateReference

`func (o *CompanySiteInfo) GetStateReference() StateReference`

GetStateReference returns the StateReference field if non-nil, zero value otherwise.

### GetStateReferenceOk

`func (o *CompanySiteInfo) GetStateReferenceOk() (*StateReference, bool)`

GetStateReferenceOk returns a tuple with the StateReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReference

`func (o *CompanySiteInfo) SetStateReference(v StateReference)`

SetStateReference sets StateReference field to given value.

### HasStateReference

`func (o *CompanySiteInfo) HasStateReference() bool`

HasStateReference returns a boolean if a field has been set.

### GetCountry

`func (o *CompanySiteInfo) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CompanySiteInfo) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CompanySiteInfo) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CompanySiteInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetZip

`func (o *CompanySiteInfo) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *CompanySiteInfo) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *CompanySiteInfo) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *CompanySiteInfo) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetAddressFormat

`func (o *CompanySiteInfo) GetAddressFormat() string`

GetAddressFormat returns the AddressFormat field if non-nil, zero value otherwise.

### GetAddressFormatOk

`func (o *CompanySiteInfo) GetAddressFormatOk() (*string, bool)`

GetAddressFormatOk returns a tuple with the AddressFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFormat

`func (o *CompanySiteInfo) SetAddressFormat(v string)`

SetAddressFormat sets AddressFormat field to given value.

### HasAddressFormat

`func (o *CompanySiteInfo) HasAddressFormat() bool`

HasAddressFormat returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CompanySiteInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CompanySiteInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CompanySiteInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CompanySiteInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *CompanySiteInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *CompanySiteInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *CompanySiteInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *CompanySiteInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *CompanySiteInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *CompanySiteInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetDefaultShippingFlag

`func (o *CompanySiteInfo) GetDefaultShippingFlag() bool`

GetDefaultShippingFlag returns the DefaultShippingFlag field if non-nil, zero value otherwise.

### GetDefaultShippingFlagOk

`func (o *CompanySiteInfo) GetDefaultShippingFlagOk() (*bool, bool)`

GetDefaultShippingFlagOk returns a tuple with the DefaultShippingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShippingFlag

`func (o *CompanySiteInfo) SetDefaultShippingFlag(v bool)`

SetDefaultShippingFlag sets DefaultShippingFlag field to given value.

### HasDefaultShippingFlag

`func (o *CompanySiteInfo) HasDefaultShippingFlag() bool`

HasDefaultShippingFlag returns a boolean if a field has been set.

### SetDefaultShippingFlagNil

`func (o *CompanySiteInfo) SetDefaultShippingFlagNil(b bool)`

 SetDefaultShippingFlagNil sets the value for DefaultShippingFlag to be an explicit nil

### UnsetDefaultShippingFlag
`func (o *CompanySiteInfo) UnsetDefaultShippingFlag()`

UnsetDefaultShippingFlag ensures that no value is present for DefaultShippingFlag, not even an explicit nil
### GetDefaultBillingFlag

`func (o *CompanySiteInfo) GetDefaultBillingFlag() bool`

GetDefaultBillingFlag returns the DefaultBillingFlag field if non-nil, zero value otherwise.

### GetDefaultBillingFlagOk

`func (o *CompanySiteInfo) GetDefaultBillingFlagOk() (*bool, bool)`

GetDefaultBillingFlagOk returns a tuple with the DefaultBillingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBillingFlag

`func (o *CompanySiteInfo) SetDefaultBillingFlag(v bool)`

SetDefaultBillingFlag sets DefaultBillingFlag field to given value.

### HasDefaultBillingFlag

`func (o *CompanySiteInfo) HasDefaultBillingFlag() bool`

HasDefaultBillingFlag returns a boolean if a field has been set.

### SetDefaultBillingFlagNil

`func (o *CompanySiteInfo) SetDefaultBillingFlagNil(b bool)`

 SetDefaultBillingFlagNil sets the value for DefaultBillingFlag to be an explicit nil

### UnsetDefaultBillingFlag
`func (o *CompanySiteInfo) UnsetDefaultBillingFlag()`

UnsetDefaultBillingFlag ensures that no value is present for DefaultBillingFlag, not even an explicit nil
### GetPrimaryAddressFlag

`func (o *CompanySiteInfo) GetPrimaryAddressFlag() bool`

GetPrimaryAddressFlag returns the PrimaryAddressFlag field if non-nil, zero value otherwise.

### GetPrimaryAddressFlagOk

`func (o *CompanySiteInfo) GetPrimaryAddressFlagOk() (*bool, bool)`

GetPrimaryAddressFlagOk returns a tuple with the PrimaryAddressFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAddressFlag

`func (o *CompanySiteInfo) SetPrimaryAddressFlag(v bool)`

SetPrimaryAddressFlag sets PrimaryAddressFlag field to given value.

### HasPrimaryAddressFlag

`func (o *CompanySiteInfo) HasPrimaryAddressFlag() bool`

HasPrimaryAddressFlag returns a boolean if a field has been set.

### SetPrimaryAddressFlagNil

`func (o *CompanySiteInfo) SetPrimaryAddressFlagNil(b bool)`

 SetPrimaryAddressFlagNil sets the value for PrimaryAddressFlag to be an explicit nil

### UnsetPrimaryAddressFlag
`func (o *CompanySiteInfo) UnsetPrimaryAddressFlag()`

UnsetPrimaryAddressFlag ensures that no value is present for PrimaryAddressFlag, not even an explicit nil
### GetTaxCode

`func (o *CompanySiteInfo) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *CompanySiteInfo) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *CompanySiteInfo) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *CompanySiteInfo) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetPhoneNumberExt

`func (o *CompanySiteInfo) GetPhoneNumberExt() string`

GetPhoneNumberExt returns the PhoneNumberExt field if non-nil, zero value otherwise.

### GetPhoneNumberExtOk

`func (o *CompanySiteInfo) GetPhoneNumberExtOk() (*string, bool)`

GetPhoneNumberExtOk returns a tuple with the PhoneNumberExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberExt

`func (o *CompanySiteInfo) SetPhoneNumberExt(v string)`

SetPhoneNumberExt sets PhoneNumberExt field to given value.

### HasPhoneNumberExt

`func (o *CompanySiteInfo) HasPhoneNumberExt() bool`

HasPhoneNumberExt returns a boolean if a field has been set.

### GetInfo

`func (o *CompanySiteInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanySiteInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanySiteInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanySiteInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


