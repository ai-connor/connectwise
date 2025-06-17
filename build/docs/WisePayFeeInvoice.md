# WisePayFeeInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**InvoiceNumber** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**InvoiceHref** | Pointer to **string** |  | [optional] 

## Methods

### NewWisePayFeeInvoice

`func NewWisePayFeeInvoice() *WisePayFeeInvoice`

NewWisePayFeeInvoice instantiates a new WisePayFeeInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWisePayFeeInvoiceWithDefaults

`func NewWisePayFeeInvoiceWithDefaults() *WisePayFeeInvoice`

NewWisePayFeeInvoiceWithDefaults instantiates a new WisePayFeeInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WisePayFeeInvoice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WisePayFeeInvoice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WisePayFeeInvoice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WisePayFeeInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WisePayFeeInvoice) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WisePayFeeInvoice) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInvoiceNumber

`func (o *WisePayFeeInvoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *WisePayFeeInvoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *WisePayFeeInvoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *WisePayFeeInvoice) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetAmount

`func (o *WisePayFeeInvoice) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WisePayFeeInvoice) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WisePayFeeInvoice) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *WisePayFeeInvoice) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *WisePayFeeInvoice) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *WisePayFeeInvoice) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetInvoiceHref

`func (o *WisePayFeeInvoice) GetInvoiceHref() string`

GetInvoiceHref returns the InvoiceHref field if non-nil, zero value otherwise.

### GetInvoiceHrefOk

`func (o *WisePayFeeInvoice) GetInvoiceHrefOk() (*string, bool)`

GetInvoiceHrefOk returns a tuple with the InvoiceHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceHref

`func (o *WisePayFeeInvoice) SetInvoiceHref(v string)`

SetInvoiceHref sets InvoiceHref field to given value.

### HasInvoiceHref

`func (o *WisePayFeeInvoice) HasInvoiceHref() bool`

HasInvoiceHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


