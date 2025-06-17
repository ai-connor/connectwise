# ExpenseTypeExemption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ExpenseType** | [**ExpenseTypeReference**](ExpenseTypeReference.md) |  | 
**TaxableLevels** | Pointer to **[]int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewExpenseTypeExemption

`func NewExpenseTypeExemption(expenseType ExpenseTypeReference, ) *ExpenseTypeExemption`

NewExpenseTypeExemption instantiates a new ExpenseTypeExemption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseTypeExemptionWithDefaults

`func NewExpenseTypeExemptionWithDefaults() *ExpenseTypeExemption`

NewExpenseTypeExemptionWithDefaults instantiates a new ExpenseTypeExemption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseTypeExemption) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseTypeExemption) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseTypeExemption) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseTypeExemption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExpenseType

`func (o *ExpenseTypeExemption) GetExpenseType() ExpenseTypeReference`

GetExpenseType returns the ExpenseType field if non-nil, zero value otherwise.

### GetExpenseTypeOk

`func (o *ExpenseTypeExemption) GetExpenseTypeOk() (*ExpenseTypeReference, bool)`

GetExpenseTypeOk returns a tuple with the ExpenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseType

`func (o *ExpenseTypeExemption) SetExpenseType(v ExpenseTypeReference)`

SetExpenseType sets ExpenseType field to given value.


### GetTaxableLevels

`func (o *ExpenseTypeExemption) GetTaxableLevels() []int32`

GetTaxableLevels returns the TaxableLevels field if non-nil, zero value otherwise.

### GetTaxableLevelsOk

`func (o *ExpenseTypeExemption) GetTaxableLevelsOk() (*[]int32, bool)`

GetTaxableLevelsOk returns a tuple with the TaxableLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableLevels

`func (o *ExpenseTypeExemption) SetTaxableLevels(v []int32)`

SetTaxableLevels sets TaxableLevels field to given value.

### HasTaxableLevels

`func (o *ExpenseTypeExemption) HasTaxableLevels() bool`

HasTaxableLevels returns a boolean if a field has been set.

### GetInfo

`func (o *ExpenseTypeExemption) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ExpenseTypeExemption) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ExpenseTypeExemption) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ExpenseTypeExemption) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


