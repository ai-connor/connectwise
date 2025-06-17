# ExpenseDetailReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Amount** | Pointer to **float64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewExpenseDetailReference

`func NewExpenseDetailReference() *ExpenseDetailReference`

NewExpenseDetailReference instantiates a new ExpenseDetailReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseDetailReferenceWithDefaults

`func NewExpenseDetailReferenceWithDefaults() *ExpenseDetailReference`

NewExpenseDetailReferenceWithDefaults instantiates a new ExpenseDetailReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseDetailReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseDetailReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseDetailReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseDetailReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExpenseDetailReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExpenseDetailReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAmount

`func (o *ExpenseDetailReference) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExpenseDetailReference) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExpenseDetailReference) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExpenseDetailReference) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetInfo

`func (o *ExpenseDetailReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ExpenseDetailReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ExpenseDetailReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ExpenseDetailReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


