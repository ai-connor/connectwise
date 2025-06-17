# TeamRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 20; | 
**AccountManagerFlag** | Pointer to **NullableBool** |  | [optional] 
**TechFlag** | Pointer to **NullableBool** |  | [optional] 
**SalesFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTeamRole

`func NewTeamRole(name string, ) *TeamRole`

NewTeamRole instantiates a new TeamRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamRoleWithDefaults

`func NewTeamRoleWithDefaults() *TeamRole`

NewTeamRoleWithDefaults instantiates a new TeamRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamRole) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TeamRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TeamRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamRole) SetName(v string)`

SetName sets Name field to given value.


### GetAccountManagerFlag

`func (o *TeamRole) GetAccountManagerFlag() bool`

GetAccountManagerFlag returns the AccountManagerFlag field if non-nil, zero value otherwise.

### GetAccountManagerFlagOk

`func (o *TeamRole) GetAccountManagerFlagOk() (*bool, bool)`

GetAccountManagerFlagOk returns a tuple with the AccountManagerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountManagerFlag

`func (o *TeamRole) SetAccountManagerFlag(v bool)`

SetAccountManagerFlag sets AccountManagerFlag field to given value.

### HasAccountManagerFlag

`func (o *TeamRole) HasAccountManagerFlag() bool`

HasAccountManagerFlag returns a boolean if a field has been set.

### SetAccountManagerFlagNil

`func (o *TeamRole) SetAccountManagerFlagNil(b bool)`

 SetAccountManagerFlagNil sets the value for AccountManagerFlag to be an explicit nil

### UnsetAccountManagerFlag
`func (o *TeamRole) UnsetAccountManagerFlag()`

UnsetAccountManagerFlag ensures that no value is present for AccountManagerFlag, not even an explicit nil
### GetTechFlag

`func (o *TeamRole) GetTechFlag() bool`

GetTechFlag returns the TechFlag field if non-nil, zero value otherwise.

### GetTechFlagOk

`func (o *TeamRole) GetTechFlagOk() (*bool, bool)`

GetTechFlagOk returns a tuple with the TechFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechFlag

`func (o *TeamRole) SetTechFlag(v bool)`

SetTechFlag sets TechFlag field to given value.

### HasTechFlag

`func (o *TeamRole) HasTechFlag() bool`

HasTechFlag returns a boolean if a field has been set.

### SetTechFlagNil

`func (o *TeamRole) SetTechFlagNil(b bool)`

 SetTechFlagNil sets the value for TechFlag to be an explicit nil

### UnsetTechFlag
`func (o *TeamRole) UnsetTechFlag()`

UnsetTechFlag ensures that no value is present for TechFlag, not even an explicit nil
### GetSalesFlag

`func (o *TeamRole) GetSalesFlag() bool`

GetSalesFlag returns the SalesFlag field if non-nil, zero value otherwise.

### GetSalesFlagOk

`func (o *TeamRole) GetSalesFlagOk() (*bool, bool)`

GetSalesFlagOk returns a tuple with the SalesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesFlag

`func (o *TeamRole) SetSalesFlag(v bool)`

SetSalesFlag sets SalesFlag field to given value.

### HasSalesFlag

`func (o *TeamRole) HasSalesFlag() bool`

HasSalesFlag returns a boolean if a field has been set.

### SetSalesFlagNil

`func (o *TeamRole) SetSalesFlagNil(b bool)`

 SetSalesFlagNil sets the value for SalesFlag to be an explicit nil

### UnsetSalesFlag
`func (o *TeamRole) UnsetSalesFlag()`

UnsetSalesFlag ensures that no value is present for SalesFlag, not even an explicit nil
### GetInfo

`func (o *TeamRole) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TeamRole) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TeamRole) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TeamRole) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


