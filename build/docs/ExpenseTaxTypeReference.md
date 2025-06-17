# ExpenseTaxTypeReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewExpenseTaxTypeReference

`func NewExpenseTaxTypeReference() *ExpenseTaxTypeReference`

NewExpenseTaxTypeReference instantiates a new ExpenseTaxTypeReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseTaxTypeReferenceWithDefaults

`func NewExpenseTaxTypeReferenceWithDefaults() *ExpenseTaxTypeReference`

NewExpenseTaxTypeReferenceWithDefaults instantiates a new ExpenseTaxTypeReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseTaxTypeReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseTaxTypeReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseTaxTypeReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseTaxTypeReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExpenseTaxTypeReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExpenseTaxTypeReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ExpenseTaxTypeReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpenseTaxTypeReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpenseTaxTypeReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExpenseTaxTypeReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInfo

`func (o *ExpenseTaxTypeReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ExpenseTaxTypeReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ExpenseTaxTypeReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ExpenseTaxTypeReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


