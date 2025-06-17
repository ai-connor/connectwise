# CompanyType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**VendorFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceAlertFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceAlertMessage** | Pointer to **string** |  Max length: 150; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyType

`func NewCompanyType(name string, ) *CompanyType`

NewCompanyType instantiates a new CompanyType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyTypeWithDefaults

`func NewCompanyTypeWithDefaults() *CompanyType`

NewCompanyTypeWithDefaults instantiates a new CompanyType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CompanyType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyType) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *CompanyType) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *CompanyType) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *CompanyType) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *CompanyType) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *CompanyType) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *CompanyType) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetVendorFlag

`func (o *CompanyType) GetVendorFlag() bool`

GetVendorFlag returns the VendorFlag field if non-nil, zero value otherwise.

### GetVendorFlagOk

`func (o *CompanyType) GetVendorFlagOk() (*bool, bool)`

GetVendorFlagOk returns a tuple with the VendorFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorFlag

`func (o *CompanyType) SetVendorFlag(v bool)`

SetVendorFlag sets VendorFlag field to given value.

### HasVendorFlag

`func (o *CompanyType) HasVendorFlag() bool`

HasVendorFlag returns a boolean if a field has been set.

### SetVendorFlagNil

`func (o *CompanyType) SetVendorFlagNil(b bool)`

 SetVendorFlagNil sets the value for VendorFlag to be an explicit nil

### UnsetVendorFlag
`func (o *CompanyType) UnsetVendorFlag()`

UnsetVendorFlag ensures that no value is present for VendorFlag, not even an explicit nil
### GetServiceAlertFlag

`func (o *CompanyType) GetServiceAlertFlag() bool`

GetServiceAlertFlag returns the ServiceAlertFlag field if non-nil, zero value otherwise.

### GetServiceAlertFlagOk

`func (o *CompanyType) GetServiceAlertFlagOk() (*bool, bool)`

GetServiceAlertFlagOk returns a tuple with the ServiceAlertFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAlertFlag

`func (o *CompanyType) SetServiceAlertFlag(v bool)`

SetServiceAlertFlag sets ServiceAlertFlag field to given value.

### HasServiceAlertFlag

`func (o *CompanyType) HasServiceAlertFlag() bool`

HasServiceAlertFlag returns a boolean if a field has been set.

### SetServiceAlertFlagNil

`func (o *CompanyType) SetServiceAlertFlagNil(b bool)`

 SetServiceAlertFlagNil sets the value for ServiceAlertFlag to be an explicit nil

### UnsetServiceAlertFlag
`func (o *CompanyType) UnsetServiceAlertFlag()`

UnsetServiceAlertFlag ensures that no value is present for ServiceAlertFlag, not even an explicit nil
### GetServiceAlertMessage

`func (o *CompanyType) GetServiceAlertMessage() string`

GetServiceAlertMessage returns the ServiceAlertMessage field if non-nil, zero value otherwise.

### GetServiceAlertMessageOk

`func (o *CompanyType) GetServiceAlertMessageOk() (*string, bool)`

GetServiceAlertMessageOk returns a tuple with the ServiceAlertMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAlertMessage

`func (o *CompanyType) SetServiceAlertMessage(v string)`

SetServiceAlertMessage sets ServiceAlertMessage field to given value.

### HasServiceAlertMessage

`func (o *CompanyType) HasServiceAlertMessage() bool`

HasServiceAlertMessage returns a boolean if a field has been set.

### GetInfo

`func (o *CompanyType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


