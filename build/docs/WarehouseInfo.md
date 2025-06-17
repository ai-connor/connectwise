# WarehouseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**OverallDefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWarehouseInfo

`func NewWarehouseInfo() *WarehouseInfo`

NewWarehouseInfo instantiates a new WarehouseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseInfoWithDefaults

`func NewWarehouseInfoWithDefaults() *WarehouseInfo`

NewWarehouseInfoWithDefaults instantiates a new WarehouseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WarehouseInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarehouseInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarehouseInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WarehouseInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WarehouseInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WarehouseInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WarehouseInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WarehouseInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *WarehouseInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *WarehouseInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *WarehouseInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *WarehouseInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *WarehouseInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *WarehouseInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetOverallDefaultFlag

`func (o *WarehouseInfo) GetOverallDefaultFlag() bool`

GetOverallDefaultFlag returns the OverallDefaultFlag field if non-nil, zero value otherwise.

### GetOverallDefaultFlagOk

`func (o *WarehouseInfo) GetOverallDefaultFlagOk() (*bool, bool)`

GetOverallDefaultFlagOk returns a tuple with the OverallDefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallDefaultFlag

`func (o *WarehouseInfo) SetOverallDefaultFlag(v bool)`

SetOverallDefaultFlag sets OverallDefaultFlag field to given value.

### HasOverallDefaultFlag

`func (o *WarehouseInfo) HasOverallDefaultFlag() bool`

HasOverallDefaultFlag returns a boolean if a field has been set.

### SetOverallDefaultFlagNil

`func (o *WarehouseInfo) SetOverallDefaultFlagNil(b bool)`

 SetOverallDefaultFlagNil sets the value for OverallDefaultFlag to be an explicit nil

### UnsetOverallDefaultFlag
`func (o *WarehouseInfo) UnsetOverallDefaultFlag()`

UnsetOverallDefaultFlag ensures that no value is present for OverallDefaultFlag, not even an explicit nil
### GetCompany

`func (o *WarehouseInfo) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *WarehouseInfo) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *WarehouseInfo) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *WarehouseInfo) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetInfo

`func (o *WarehouseInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WarehouseInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WarehouseInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WarehouseInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


