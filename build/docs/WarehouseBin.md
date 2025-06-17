# WarehouseBin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Warehouse** | [**WarehouseReference**](WarehouseReference.md) |  | 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**MinQuantity** | Pointer to **NullableFloat64** |  | [optional] 
**MaxQuantity** | Pointer to **NullableFloat64** |  | [optional] 
**OverflowBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**Manager** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Length** | Pointer to **NullableFloat64** |  | [optional] 
**Width** | Pointer to **NullableFloat64** |  | [optional] 
**Height** | Pointer to **NullableFloat64** |  | [optional] 
**Weight** | Pointer to **NullableFloat64** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**QuantityOnHand** | Pointer to **NullableInt32** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**TransferBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWarehouseBin

`func NewWarehouseBin(name string, warehouse WarehouseReference, ) *WarehouseBin`

NewWarehouseBin instantiates a new WarehouseBin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseBinWithDefaults

`func NewWarehouseBinWithDefaults() *WarehouseBin`

NewWarehouseBinWithDefaults instantiates a new WarehouseBin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WarehouseBin) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarehouseBin) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarehouseBin) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WarehouseBin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WarehouseBin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WarehouseBin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WarehouseBin) SetName(v string)`

SetName sets Name field to given value.


### GetWarehouse

`func (o *WarehouseBin) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *WarehouseBin) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *WarehouseBin) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.


### GetLocation

`func (o *WarehouseBin) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WarehouseBin) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WarehouseBin) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *WarehouseBin) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *WarehouseBin) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *WarehouseBin) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *WarehouseBin) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *WarehouseBin) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetMinQuantity

`func (o *WarehouseBin) GetMinQuantity() float64`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *WarehouseBin) GetMinQuantityOk() (*float64, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *WarehouseBin) SetMinQuantity(v float64)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *WarehouseBin) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### SetMinQuantityNil

`func (o *WarehouseBin) SetMinQuantityNil(b bool)`

 SetMinQuantityNil sets the value for MinQuantity to be an explicit nil

### UnsetMinQuantity
`func (o *WarehouseBin) UnsetMinQuantity()`

UnsetMinQuantity ensures that no value is present for MinQuantity, not even an explicit nil
### GetMaxQuantity

`func (o *WarehouseBin) GetMaxQuantity() float64`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *WarehouseBin) GetMaxQuantityOk() (*float64, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *WarehouseBin) SetMaxQuantity(v float64)`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *WarehouseBin) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### SetMaxQuantityNil

`func (o *WarehouseBin) SetMaxQuantityNil(b bool)`

 SetMaxQuantityNil sets the value for MaxQuantity to be an explicit nil

### UnsetMaxQuantity
`func (o *WarehouseBin) UnsetMaxQuantity()`

UnsetMaxQuantity ensures that no value is present for MaxQuantity, not even an explicit nil
### GetOverflowBin

`func (o *WarehouseBin) GetOverflowBin() WarehouseBinReference`

GetOverflowBin returns the OverflowBin field if non-nil, zero value otherwise.

### GetOverflowBinOk

`func (o *WarehouseBin) GetOverflowBinOk() (*WarehouseBinReference, bool)`

GetOverflowBinOk returns a tuple with the OverflowBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflowBin

`func (o *WarehouseBin) SetOverflowBin(v WarehouseBinReference)`

SetOverflowBin sets OverflowBin field to given value.

### HasOverflowBin

`func (o *WarehouseBin) HasOverflowBin() bool`

HasOverflowBin returns a boolean if a field has been set.

### GetManager

`func (o *WarehouseBin) GetManager() MemberReference`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *WarehouseBin) GetManagerOk() (*MemberReference, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *WarehouseBin) SetManager(v MemberReference)`

SetManager sets Manager field to given value.

### HasManager

`func (o *WarehouseBin) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetLength

`func (o *WarehouseBin) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *WarehouseBin) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *WarehouseBin) SetLength(v float64)`

SetLength sets Length field to given value.

### HasLength

`func (o *WarehouseBin) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *WarehouseBin) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *WarehouseBin) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetWidth

`func (o *WarehouseBin) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *WarehouseBin) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *WarehouseBin) SetWidth(v float64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *WarehouseBin) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *WarehouseBin) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *WarehouseBin) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *WarehouseBin) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *WarehouseBin) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *WarehouseBin) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *WarehouseBin) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *WarehouseBin) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *WarehouseBin) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetWeight

`func (o *WarehouseBin) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WarehouseBin) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WarehouseBin) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WarehouseBin) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *WarehouseBin) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *WarehouseBin) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetDefaultFlag

`func (o *WarehouseBin) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *WarehouseBin) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *WarehouseBin) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *WarehouseBin) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *WarehouseBin) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *WarehouseBin) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *WarehouseBin) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *WarehouseBin) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *WarehouseBin) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *WarehouseBin) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *WarehouseBin) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *WarehouseBin) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetQuantityOnHand

`func (o *WarehouseBin) GetQuantityOnHand() int32`

GetQuantityOnHand returns the QuantityOnHand field if non-nil, zero value otherwise.

### GetQuantityOnHandOk

`func (o *WarehouseBin) GetQuantityOnHandOk() (*int32, bool)`

GetQuantityOnHandOk returns a tuple with the QuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnHand

`func (o *WarehouseBin) SetQuantityOnHand(v int32)`

SetQuantityOnHand sets QuantityOnHand field to given value.

### HasQuantityOnHand

`func (o *WarehouseBin) HasQuantityOnHand() bool`

HasQuantityOnHand returns a boolean if a field has been set.

### SetQuantityOnHandNil

`func (o *WarehouseBin) SetQuantityOnHandNil(b bool)`

 SetQuantityOnHandNil sets the value for QuantityOnHand to be an explicit nil

### UnsetQuantityOnHand
`func (o *WarehouseBin) UnsetQuantityOnHand()`

UnsetQuantityOnHand ensures that no value is present for QuantityOnHand, not even an explicit nil
### GetCompany

`func (o *WarehouseBin) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *WarehouseBin) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *WarehouseBin) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *WarehouseBin) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetTransferBin

`func (o *WarehouseBin) GetTransferBin() WarehouseBinReference`

GetTransferBin returns the TransferBin field if non-nil, zero value otherwise.

### GetTransferBinOk

`func (o *WarehouseBin) GetTransferBinOk() (*WarehouseBinReference, bool)`

GetTransferBinOk returns a tuple with the TransferBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferBin

`func (o *WarehouseBin) SetTransferBin(v WarehouseBinReference)`

SetTransferBin sets TransferBin field to given value.

### HasTransferBin

`func (o *WarehouseBin) HasTransferBin() bool`

HasTransferBin returns a boolean if a field has been set.

### GetInfo

`func (o *WarehouseBin) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WarehouseBin) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WarehouseBin) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WarehouseBin) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


