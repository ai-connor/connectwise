# ExpenseTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**Type** | Pointer to [**ExpenseTaxTypeReference**](ExpenseTaxTypeReference.md) |  | [optional] 

## Methods

### NewExpenseTax

`func NewExpenseTax() *ExpenseTax`

NewExpenseTax instantiates a new ExpenseTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseTaxWithDefaults

`func NewExpenseTaxWithDefaults() *ExpenseTax`

NewExpenseTaxWithDefaults instantiates a new ExpenseTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseTax) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseTax) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseTax) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseTax) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAmount

`func (o *ExpenseTax) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExpenseTax) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExpenseTax) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExpenseTax) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *ExpenseTax) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *ExpenseTax) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetType

`func (o *ExpenseTax) GetType() ExpenseTaxTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExpenseTax) GetTypeOk() (*ExpenseTaxTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExpenseTax) SetType(v ExpenseTaxTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *ExpenseTax) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


