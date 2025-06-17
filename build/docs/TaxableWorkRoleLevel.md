# TaxableWorkRoleLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TaxCodeLevel** | [**TaxCodeLevelReference**](TaxCodeLevelReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTaxableWorkRoleLevel

`func NewTaxableWorkRoleLevel(taxCodeLevel TaxCodeLevelReference, ) *TaxableWorkRoleLevel`

NewTaxableWorkRoleLevel instantiates a new TaxableWorkRoleLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxableWorkRoleLevelWithDefaults

`func NewTaxableWorkRoleLevelWithDefaults() *TaxableWorkRoleLevel`

NewTaxableWorkRoleLevelWithDefaults instantiates a new TaxableWorkRoleLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxableWorkRoleLevel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxableWorkRoleLevel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxableWorkRoleLevel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaxableWorkRoleLevel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTaxCodeLevel

`func (o *TaxableWorkRoleLevel) GetTaxCodeLevel() TaxCodeLevelReference`

GetTaxCodeLevel returns the TaxCodeLevel field if non-nil, zero value otherwise.

### GetTaxCodeLevelOk

`func (o *TaxableWorkRoleLevel) GetTaxCodeLevelOk() (*TaxCodeLevelReference, bool)`

GetTaxCodeLevelOk returns a tuple with the TaxCodeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodeLevel

`func (o *TaxableWorkRoleLevel) SetTaxCodeLevel(v TaxCodeLevelReference)`

SetTaxCodeLevel sets TaxCodeLevel field to given value.


### GetInfo

`func (o *TaxableWorkRoleLevel) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TaxableWorkRoleLevel) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TaxableWorkRoleLevel) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TaxableWorkRoleLevel) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


