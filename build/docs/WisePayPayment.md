# WisePayPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentDateUtc** | Pointer to **string** |  | [optional] 
**WisePayReference** | Pointer to **string** |  | [optional] 
**BatchPayment** | Pointer to [**WisePayBatchPayment**](WisePayBatchPayment.md) |  | [optional] 
**FeeInvoice** | Pointer to [**WisePayFeeInvoice**](WisePayFeeInvoice.md) |  | [optional] 

## Methods

### NewWisePayPayment

`func NewWisePayPayment() *WisePayPayment`

NewWisePayPayment instantiates a new WisePayPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWisePayPaymentWithDefaults

`func NewWisePayPaymentWithDefaults() *WisePayPayment`

NewWisePayPaymentWithDefaults instantiates a new WisePayPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentDateUtc

`func (o *WisePayPayment) GetPaymentDateUtc() string`

GetPaymentDateUtc returns the PaymentDateUtc field if non-nil, zero value otherwise.

### GetPaymentDateUtcOk

`func (o *WisePayPayment) GetPaymentDateUtcOk() (*string, bool)`

GetPaymentDateUtcOk returns a tuple with the PaymentDateUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDateUtc

`func (o *WisePayPayment) SetPaymentDateUtc(v string)`

SetPaymentDateUtc sets PaymentDateUtc field to given value.

### HasPaymentDateUtc

`func (o *WisePayPayment) HasPaymentDateUtc() bool`

HasPaymentDateUtc returns a boolean if a field has been set.

### GetWisePayReference

`func (o *WisePayPayment) GetWisePayReference() string`

GetWisePayReference returns the WisePayReference field if non-nil, zero value otherwise.

### GetWisePayReferenceOk

`func (o *WisePayPayment) GetWisePayReferenceOk() (*string, bool)`

GetWisePayReferenceOk returns a tuple with the WisePayReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWisePayReference

`func (o *WisePayPayment) SetWisePayReference(v string)`

SetWisePayReference sets WisePayReference field to given value.

### HasWisePayReference

`func (o *WisePayPayment) HasWisePayReference() bool`

HasWisePayReference returns a boolean if a field has been set.

### GetBatchPayment

`func (o *WisePayPayment) GetBatchPayment() WisePayBatchPayment`

GetBatchPayment returns the BatchPayment field if non-nil, zero value otherwise.

### GetBatchPaymentOk

`func (o *WisePayPayment) GetBatchPaymentOk() (*WisePayBatchPayment, bool)`

GetBatchPaymentOk returns a tuple with the BatchPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchPayment

`func (o *WisePayPayment) SetBatchPayment(v WisePayBatchPayment)`

SetBatchPayment sets BatchPayment field to given value.

### HasBatchPayment

`func (o *WisePayPayment) HasBatchPayment() bool`

HasBatchPayment returns a boolean if a field has been set.

### GetFeeInvoice

`func (o *WisePayPayment) GetFeeInvoice() WisePayFeeInvoice`

GetFeeInvoice returns the FeeInvoice field if non-nil, zero value otherwise.

### GetFeeInvoiceOk

`func (o *WisePayPayment) GetFeeInvoiceOk() (*WisePayFeeInvoice, bool)`

GetFeeInvoiceOk returns a tuple with the FeeInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeInvoice

`func (o *WisePayPayment) SetFeeInvoice(v WisePayFeeInvoice)`

SetFeeInvoice sets FeeInvoice field to given value.

### HasFeeInvoice

`func (o *WisePayPayment) HasFeeInvoice() bool`

HasFeeInvoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


