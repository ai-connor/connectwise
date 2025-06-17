# Warehouse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Location** | [**SystemLocationReference**](SystemLocationReference.md) |  | 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Department** | [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | 
**Manager** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**LocationXref** | Pointer to **string** |  Max length: 10; | [optional] 
**LocationDefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**OverallDefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**LockedFlag** | Pointer to **NullableBool** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWarehouse

`func NewWarehouse(name string, location SystemLocationReference, department SystemDepartmentReference, ) *Warehouse`

NewWarehouse instantiates a new Warehouse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseWithDefaults

`func NewWarehouseWithDefaults() *Warehouse`

NewWarehouseWithDefaults instantiates a new Warehouse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Warehouse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Warehouse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Warehouse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Warehouse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Warehouse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Warehouse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Warehouse) SetName(v string)`

SetName sets Name field to given value.


### GetCompany

`func (o *Warehouse) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Warehouse) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Warehouse) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Warehouse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLocation

`func (o *Warehouse) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Warehouse) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Warehouse) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.


### GetContact

`func (o *Warehouse) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Warehouse) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Warehouse) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Warehouse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDepartment

`func (o *Warehouse) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Warehouse) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Warehouse) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.


### GetManager

`func (o *Warehouse) GetManager() MemberReference`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *Warehouse) GetManagerOk() (*MemberReference, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *Warehouse) SetManager(v MemberReference)`

SetManager sets Manager field to given value.

### HasManager

`func (o *Warehouse) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetSite

`func (o *Warehouse) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Warehouse) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Warehouse) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *Warehouse) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetLocationXref

`func (o *Warehouse) GetLocationXref() string`

GetLocationXref returns the LocationXref field if non-nil, zero value otherwise.

### GetLocationXrefOk

`func (o *Warehouse) GetLocationXrefOk() (*string, bool)`

GetLocationXrefOk returns a tuple with the LocationXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationXref

`func (o *Warehouse) SetLocationXref(v string)`

SetLocationXref sets LocationXref field to given value.

### HasLocationXref

`func (o *Warehouse) HasLocationXref() bool`

HasLocationXref returns a boolean if a field has been set.

### GetLocationDefaultFlag

`func (o *Warehouse) GetLocationDefaultFlag() bool`

GetLocationDefaultFlag returns the LocationDefaultFlag field if non-nil, zero value otherwise.

### GetLocationDefaultFlagOk

`func (o *Warehouse) GetLocationDefaultFlagOk() (*bool, bool)`

GetLocationDefaultFlagOk returns a tuple with the LocationDefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationDefaultFlag

`func (o *Warehouse) SetLocationDefaultFlag(v bool)`

SetLocationDefaultFlag sets LocationDefaultFlag field to given value.

### HasLocationDefaultFlag

`func (o *Warehouse) HasLocationDefaultFlag() bool`

HasLocationDefaultFlag returns a boolean if a field has been set.

### SetLocationDefaultFlagNil

`func (o *Warehouse) SetLocationDefaultFlagNil(b bool)`

 SetLocationDefaultFlagNil sets the value for LocationDefaultFlag to be an explicit nil

### UnsetLocationDefaultFlag
`func (o *Warehouse) UnsetLocationDefaultFlag()`

UnsetLocationDefaultFlag ensures that no value is present for LocationDefaultFlag, not even an explicit nil
### GetOverallDefaultFlag

`func (o *Warehouse) GetOverallDefaultFlag() bool`

GetOverallDefaultFlag returns the OverallDefaultFlag field if non-nil, zero value otherwise.

### GetOverallDefaultFlagOk

`func (o *Warehouse) GetOverallDefaultFlagOk() (*bool, bool)`

GetOverallDefaultFlagOk returns a tuple with the OverallDefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallDefaultFlag

`func (o *Warehouse) SetOverallDefaultFlag(v bool)`

SetOverallDefaultFlag sets OverallDefaultFlag field to given value.

### HasOverallDefaultFlag

`func (o *Warehouse) HasOverallDefaultFlag() bool`

HasOverallDefaultFlag returns a boolean if a field has been set.

### SetOverallDefaultFlagNil

`func (o *Warehouse) SetOverallDefaultFlagNil(b bool)`

 SetOverallDefaultFlagNil sets the value for OverallDefaultFlag to be an explicit nil

### UnsetOverallDefaultFlag
`func (o *Warehouse) UnsetOverallDefaultFlag()`

UnsetOverallDefaultFlag ensures that no value is present for OverallDefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *Warehouse) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *Warehouse) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *Warehouse) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *Warehouse) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *Warehouse) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *Warehouse) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetLockedFlag

`func (o *Warehouse) GetLockedFlag() bool`

GetLockedFlag returns the LockedFlag field if non-nil, zero value otherwise.

### GetLockedFlagOk

`func (o *Warehouse) GetLockedFlagOk() (*bool, bool)`

GetLockedFlagOk returns a tuple with the LockedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedFlag

`func (o *Warehouse) SetLockedFlag(v bool)`

SetLockedFlag sets LockedFlag field to given value.

### HasLockedFlag

`func (o *Warehouse) HasLockedFlag() bool`

HasLockedFlag returns a boolean if a field has been set.

### SetLockedFlagNil

`func (o *Warehouse) SetLockedFlagNil(b bool)`

 SetLockedFlagNil sets the value for LockedFlag to be an explicit nil

### UnsetLockedFlag
`func (o *Warehouse) UnsetLockedFlag()`

UnsetLockedFlag ensures that no value is present for LockedFlag, not even an explicit nil
### GetCurrency

`func (o *Warehouse) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Warehouse) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Warehouse) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Warehouse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInfo

`func (o *Warehouse) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Warehouse) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Warehouse) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Warehouse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


