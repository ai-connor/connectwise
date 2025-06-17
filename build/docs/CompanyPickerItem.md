# CompanyPickerItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
**CompanyStatus** | Pointer to [**CompanyStatusReference**](CompanyStatusReference.md) |  | [optional] 
**CompanyType** | Pointer to [**CompanyTypeReference**](CompanyTypeReference.md) |  | [optional] 
**CompanySite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**CompanyLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**CompanyCountry** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**VendorPickerFlag** | Pointer to **NullableBool** | Gets or sets if true, this record was created by the vendor picker component. Otherwise, the record was created by the company picker component. | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyPickerItem

`func NewCompanyPickerItem(company CompanyReference, ) *CompanyPickerItem`

NewCompanyPickerItem instantiates a new CompanyPickerItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyPickerItemWithDefaults

`func NewCompanyPickerItemWithDefaults() *CompanyPickerItem`

NewCompanyPickerItemWithDefaults instantiates a new CompanyPickerItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyPickerItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyPickerItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyPickerItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyPickerItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *CompanyPickerItem) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *CompanyPickerItem) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *CompanyPickerItem) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *CompanyPickerItem) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetCompany

`func (o *CompanyPickerItem) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyPickerItem) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyPickerItem) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetCompanyStatus

`func (o *CompanyPickerItem) GetCompanyStatus() CompanyStatusReference`

GetCompanyStatus returns the CompanyStatus field if non-nil, zero value otherwise.

### GetCompanyStatusOk

`func (o *CompanyPickerItem) GetCompanyStatusOk() (*CompanyStatusReference, bool)`

GetCompanyStatusOk returns a tuple with the CompanyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyStatus

`func (o *CompanyPickerItem) SetCompanyStatus(v CompanyStatusReference)`

SetCompanyStatus sets CompanyStatus field to given value.

### HasCompanyStatus

`func (o *CompanyPickerItem) HasCompanyStatus() bool`

HasCompanyStatus returns a boolean if a field has been set.

### GetCompanyType

`func (o *CompanyPickerItem) GetCompanyType() CompanyTypeReference`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *CompanyPickerItem) GetCompanyTypeOk() (*CompanyTypeReference, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *CompanyPickerItem) SetCompanyType(v CompanyTypeReference)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *CompanyPickerItem) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### GetCompanySite

`func (o *CompanyPickerItem) GetCompanySite() SiteReference`

GetCompanySite returns the CompanySite field if non-nil, zero value otherwise.

### GetCompanySiteOk

`func (o *CompanyPickerItem) GetCompanySiteOk() (*SiteReference, bool)`

GetCompanySiteOk returns a tuple with the CompanySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySite

`func (o *CompanyPickerItem) SetCompanySite(v SiteReference)`

SetCompanySite sets CompanySite field to given value.

### HasCompanySite

`func (o *CompanyPickerItem) HasCompanySite() bool`

HasCompanySite returns a boolean if a field has been set.

### GetCompanyLocation

`func (o *CompanyPickerItem) GetCompanyLocation() SystemLocationReference`

GetCompanyLocation returns the CompanyLocation field if non-nil, zero value otherwise.

### GetCompanyLocationOk

`func (o *CompanyPickerItem) GetCompanyLocationOk() (*SystemLocationReference, bool)`

GetCompanyLocationOk returns a tuple with the CompanyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLocation

`func (o *CompanyPickerItem) SetCompanyLocation(v SystemLocationReference)`

SetCompanyLocation sets CompanyLocation field to given value.

### HasCompanyLocation

`func (o *CompanyPickerItem) HasCompanyLocation() bool`

HasCompanyLocation returns a boolean if a field has been set.

### GetCompanyCountry

`func (o *CompanyPickerItem) GetCompanyCountry() CountryReference`

GetCompanyCountry returns the CompanyCountry field if non-nil, zero value otherwise.

### GetCompanyCountryOk

`func (o *CompanyPickerItem) GetCompanyCountryOk() (*CountryReference, bool)`

GetCompanyCountryOk returns a tuple with the CompanyCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCountry

`func (o *CompanyPickerItem) SetCompanyCountry(v CountryReference)`

SetCompanyCountry sets CompanyCountry field to given value.

### HasCompanyCountry

`func (o *CompanyPickerItem) HasCompanyCountry() bool`

HasCompanyCountry returns a boolean if a field has been set.

### GetVendorPickerFlag

`func (o *CompanyPickerItem) GetVendorPickerFlag() bool`

GetVendorPickerFlag returns the VendorPickerFlag field if non-nil, zero value otherwise.

### GetVendorPickerFlagOk

`func (o *CompanyPickerItem) GetVendorPickerFlagOk() (*bool, bool)`

GetVendorPickerFlagOk returns a tuple with the VendorPickerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPickerFlag

`func (o *CompanyPickerItem) SetVendorPickerFlag(v bool)`

SetVendorPickerFlag sets VendorPickerFlag field to given value.

### HasVendorPickerFlag

`func (o *CompanyPickerItem) HasVendorPickerFlag() bool`

HasVendorPickerFlag returns a boolean if a field has been set.

### SetVendorPickerFlagNil

`func (o *CompanyPickerItem) SetVendorPickerFlagNil(b bool)`

 SetVendorPickerFlagNil sets the value for VendorPickerFlag to be an explicit nil

### UnsetVendorPickerFlag
`func (o *CompanyPickerItem) UnsetVendorPickerFlag()`

UnsetVendorPickerFlag ensures that no value is present for VendorPickerFlag, not even an explicit nil
### GetInfo

`func (o *CompanyPickerItem) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyPickerItem) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyPickerItem) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyPickerItem) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


