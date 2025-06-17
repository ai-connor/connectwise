# UnpostedInvoiceTaxableLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TaxLevel** | Pointer to **int32** |  | [optional] 
**TaxCodeXref** | Pointer to **string** |  | [optional] 
**TaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUnpostedInvoiceTaxableLevel

`func NewUnpostedInvoiceTaxableLevel() *UnpostedInvoiceTaxableLevel`

NewUnpostedInvoiceTaxableLevel instantiates a new UnpostedInvoiceTaxableLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnpostedInvoiceTaxableLevelWithDefaults

`func NewUnpostedInvoiceTaxableLevelWithDefaults() *UnpostedInvoiceTaxableLevel`

NewUnpostedInvoiceTaxableLevelWithDefaults instantiates a new UnpostedInvoiceTaxableLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnpostedInvoiceTaxableLevel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnpostedInvoiceTaxableLevel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnpostedInvoiceTaxableLevel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UnpostedInvoiceTaxableLevel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTaxLevel

`func (o *UnpostedInvoiceTaxableLevel) GetTaxLevel() int32`

GetTaxLevel returns the TaxLevel field if non-nil, zero value otherwise.

### GetTaxLevelOk

`func (o *UnpostedInvoiceTaxableLevel) GetTaxLevelOk() (*int32, bool)`

GetTaxLevelOk returns a tuple with the TaxLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLevel

`func (o *UnpostedInvoiceTaxableLevel) SetTaxLevel(v int32)`

SetTaxLevel sets TaxLevel field to given value.

### HasTaxLevel

`func (o *UnpostedInvoiceTaxableLevel) HasTaxLevel() bool`

HasTaxLevel returns a boolean if a field has been set.

### GetTaxCodeXref

`func (o *UnpostedInvoiceTaxableLevel) GetTaxCodeXref() string`

GetTaxCodeXref returns the TaxCodeXref field if non-nil, zero value otherwise.

### GetTaxCodeXrefOk

`func (o *UnpostedInvoiceTaxableLevel) GetTaxCodeXrefOk() (*string, bool)`

GetTaxCodeXrefOk returns a tuple with the TaxCodeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodeXref

`func (o *UnpostedInvoiceTaxableLevel) SetTaxCodeXref(v string)`

SetTaxCodeXref sets TaxCodeXref field to given value.

### HasTaxCodeXref

`func (o *UnpostedInvoiceTaxableLevel) HasTaxCodeXref() bool`

HasTaxCodeXref returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnpostedInvoiceTaxableLevel) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnpostedInvoiceTaxableLevel) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnpostedInvoiceTaxableLevel) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnpostedInvoiceTaxableLevel) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### SetTaxAmountNil

`func (o *UnpostedInvoiceTaxableLevel) SetTaxAmountNil(b bool)`

 SetTaxAmountNil sets the value for TaxAmount to be an explicit nil

### UnsetTaxAmount
`func (o *UnpostedInvoiceTaxableLevel) UnsetTaxAmount()`

UnsetTaxAmount ensures that no value is present for TaxAmount, not even an explicit nil
### GetInfo

`func (o *UnpostedInvoiceTaxableLevel) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UnpostedInvoiceTaxableLevel) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UnpostedInvoiceTaxableLevel) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *UnpostedInvoiceTaxableLevel) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


