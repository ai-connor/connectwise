# WisePayBatchPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**WisePayHref** | Pointer to **string** |  | [optional] 

## Methods

### NewWisePayBatchPayment

`func NewWisePayBatchPayment() *WisePayBatchPayment`

NewWisePayBatchPayment instantiates a new WisePayBatchPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWisePayBatchPaymentWithDefaults

`func NewWisePayBatchPaymentWithDefaults() *WisePayBatchPayment`

NewWisePayBatchPaymentWithDefaults instantiates a new WisePayBatchPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *WisePayBatchPayment) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WisePayBatchPayment) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WisePayBatchPayment) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *WisePayBatchPayment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *WisePayBatchPayment) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *WisePayBatchPayment) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetWisePayHref

`func (o *WisePayBatchPayment) GetWisePayHref() string`

GetWisePayHref returns the WisePayHref field if non-nil, zero value otherwise.

### GetWisePayHrefOk

`func (o *WisePayBatchPayment) GetWisePayHrefOk() (*string, bool)`

GetWisePayHrefOk returns a tuple with the WisePayHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWisePayHref

`func (o *WisePayBatchPayment) SetWisePayHref(v string)`

SetWisePayHref sets WisePayHref field to given value.

### HasWisePayHref

`func (o *WisePayBatchPayment) HasWisePayHref() bool`

HasWisePayHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


