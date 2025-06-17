# WorkRoleExemption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**WorkRole** | [**WorkRoleReference**](WorkRoleReference.md) |  | 
**TaxableLevels** | Pointer to **[]int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkRoleExemption

`func NewWorkRoleExemption(workRole WorkRoleReference, ) *WorkRoleExemption`

NewWorkRoleExemption instantiates a new WorkRoleExemption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkRoleExemptionWithDefaults

`func NewWorkRoleExemptionWithDefaults() *WorkRoleExemption`

NewWorkRoleExemptionWithDefaults instantiates a new WorkRoleExemption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkRoleExemption) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkRoleExemption) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkRoleExemption) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkRoleExemption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWorkRole

`func (o *WorkRoleExemption) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *WorkRoleExemption) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *WorkRoleExemption) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.


### GetTaxableLevels

`func (o *WorkRoleExemption) GetTaxableLevels() []int32`

GetTaxableLevels returns the TaxableLevels field if non-nil, zero value otherwise.

### GetTaxableLevelsOk

`func (o *WorkRoleExemption) GetTaxableLevelsOk() (*[]int32, bool)`

GetTaxableLevelsOk returns a tuple with the TaxableLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableLevels

`func (o *WorkRoleExemption) SetTaxableLevels(v []int32)`

SetTaxableLevels sets TaxableLevels field to given value.

### HasTaxableLevels

`func (o *WorkRoleExemption) HasTaxableLevels() bool`

HasTaxableLevels returns a boolean if a field has been set.

### GetInfo

`func (o *WorkRoleExemption) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkRoleExemption) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkRoleExemption) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkRoleExemption) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


