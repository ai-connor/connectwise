# ChargeCodeExpenseType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | [**ExpenseTypeReference**](ExpenseTypeReference.md) |  | 
**ChargeCode** | Pointer to [**ChargeCodeReference**](ChargeCodeReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewChargeCodeExpenseType

`func NewChargeCodeExpenseType(type_ ExpenseTypeReference, ) *ChargeCodeExpenseType`

NewChargeCodeExpenseType instantiates a new ChargeCodeExpenseType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeCodeExpenseTypeWithDefaults

`func NewChargeCodeExpenseTypeWithDefaults() *ChargeCodeExpenseType`

NewChargeCodeExpenseTypeWithDefaults instantiates a new ChargeCodeExpenseType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChargeCodeExpenseType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargeCodeExpenseType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargeCodeExpenseType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChargeCodeExpenseType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ChargeCodeExpenseType) GetType() ExpenseTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChargeCodeExpenseType) GetTypeOk() (*ExpenseTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChargeCodeExpenseType) SetType(v ExpenseTypeReference)`

SetType sets Type field to given value.


### GetChargeCode

`func (o *ChargeCodeExpenseType) GetChargeCode() ChargeCodeReference`

GetChargeCode returns the ChargeCode field if non-nil, zero value otherwise.

### GetChargeCodeOk

`func (o *ChargeCodeExpenseType) GetChargeCodeOk() (*ChargeCodeReference, bool)`

GetChargeCodeOk returns a tuple with the ChargeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeCode

`func (o *ChargeCodeExpenseType) SetChargeCode(v ChargeCodeReference)`

SetChargeCode sets ChargeCode field to given value.

### HasChargeCode

`func (o *ChargeCodeExpenseType) HasChargeCode() bool`

HasChargeCode returns a boolean if a field has been set.

### GetInfo

`func (o *ChargeCodeExpenseType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ChargeCodeExpenseType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ChargeCodeExpenseType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ChargeCodeExpenseType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


