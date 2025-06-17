# ExpenseTaxTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Inactive** | Pointer to **bool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewExpenseTaxTypeInfo

`func NewExpenseTaxTypeInfo() *ExpenseTaxTypeInfo`

NewExpenseTaxTypeInfo instantiates a new ExpenseTaxTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseTaxTypeInfoWithDefaults

`func NewExpenseTaxTypeInfoWithDefaults() *ExpenseTaxTypeInfo`

NewExpenseTaxTypeInfoWithDefaults instantiates a new ExpenseTaxTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseTaxTypeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseTaxTypeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseTaxTypeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseTaxTypeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ExpenseTaxTypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpenseTaxTypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpenseTaxTypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExpenseTaxTypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInactive

`func (o *ExpenseTaxTypeInfo) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *ExpenseTaxTypeInfo) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *ExpenseTaxTypeInfo) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *ExpenseTaxTypeInfo) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetInfo

`func (o *ExpenseTaxTypeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ExpenseTaxTypeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ExpenseTaxTypeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ExpenseTaxTypeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


