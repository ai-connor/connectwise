# UnpostedPayments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Invoice** | Pointer to [**InvoiceReference**](InvoiceReference.md) |  | [optional] 
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**PaymentDate** | Pointer to **string** |  | [optional] 
**AppliedBy** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**WisePayPayment** | Pointer to [**WisePayPayment**](WisePayPayment.md) |  | [optional] 
**PaymentSyncStatus** | Pointer to **string** |  | [optional] 
**PaymentSyncDate** | Pointer to **string** |  | [optional] 
**PaymentAccount** | Pointer to **string** |  | [optional] 
**ARPaymentAccount** | Pointer to **string** |  | [optional] 

## Methods

### NewUnpostedPayments

`func NewUnpostedPayments() *UnpostedPayments`

NewUnpostedPayments instantiates a new UnpostedPayments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnpostedPaymentsWithDefaults

`func NewUnpostedPaymentsWithDefaults() *UnpostedPayments`

NewUnpostedPaymentsWithDefaults instantiates a new UnpostedPayments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnpostedPayments) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnpostedPayments) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnpostedPayments) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UnpostedPayments) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UnpostedPayments) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnpostedPayments) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnpostedPayments) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UnpostedPayments) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSource

`func (o *UnpostedPayments) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UnpostedPayments) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UnpostedPayments) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UnpostedPayments) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetInvoice

`func (o *UnpostedPayments) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *UnpostedPayments) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *UnpostedPayments) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *UnpostedPayments) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetAmount

`func (o *UnpostedPayments) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnpostedPayments) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnpostedPayments) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnpostedPayments) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *UnpostedPayments) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *UnpostedPayments) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetPaymentDate

`func (o *UnpostedPayments) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *UnpostedPayments) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *UnpostedPayments) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *UnpostedPayments) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetAppliedBy

`func (o *UnpostedPayments) GetAppliedBy() string`

GetAppliedBy returns the AppliedBy field if non-nil, zero value otherwise.

### GetAppliedByOk

`func (o *UnpostedPayments) GetAppliedByOk() (*string, bool)`

GetAppliedByOk returns a tuple with the AppliedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedBy

`func (o *UnpostedPayments) SetAppliedBy(v string)`

SetAppliedBy sets AppliedBy field to given value.

### HasAppliedBy

`func (o *UnpostedPayments) HasAppliedBy() bool`

HasAppliedBy returns a boolean if a field has been set.

### GetInfo

`func (o *UnpostedPayments) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UnpostedPayments) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UnpostedPayments) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *UnpostedPayments) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetWisePayPayment

`func (o *UnpostedPayments) GetWisePayPayment() WisePayPayment`

GetWisePayPayment returns the WisePayPayment field if non-nil, zero value otherwise.

### GetWisePayPaymentOk

`func (o *UnpostedPayments) GetWisePayPaymentOk() (*WisePayPayment, bool)`

GetWisePayPaymentOk returns a tuple with the WisePayPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWisePayPayment

`func (o *UnpostedPayments) SetWisePayPayment(v WisePayPayment)`

SetWisePayPayment sets WisePayPayment field to given value.

### HasWisePayPayment

`func (o *UnpostedPayments) HasWisePayPayment() bool`

HasWisePayPayment returns a boolean if a field has been set.

### GetPaymentSyncStatus

`func (o *UnpostedPayments) GetPaymentSyncStatus() string`

GetPaymentSyncStatus returns the PaymentSyncStatus field if non-nil, zero value otherwise.

### GetPaymentSyncStatusOk

`func (o *UnpostedPayments) GetPaymentSyncStatusOk() (*string, bool)`

GetPaymentSyncStatusOk returns a tuple with the PaymentSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSyncStatus

`func (o *UnpostedPayments) SetPaymentSyncStatus(v string)`

SetPaymentSyncStatus sets PaymentSyncStatus field to given value.

### HasPaymentSyncStatus

`func (o *UnpostedPayments) HasPaymentSyncStatus() bool`

HasPaymentSyncStatus returns a boolean if a field has been set.

### GetPaymentSyncDate

`func (o *UnpostedPayments) GetPaymentSyncDate() string`

GetPaymentSyncDate returns the PaymentSyncDate field if non-nil, zero value otherwise.

### GetPaymentSyncDateOk

`func (o *UnpostedPayments) GetPaymentSyncDateOk() (*string, bool)`

GetPaymentSyncDateOk returns a tuple with the PaymentSyncDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSyncDate

`func (o *UnpostedPayments) SetPaymentSyncDate(v string)`

SetPaymentSyncDate sets PaymentSyncDate field to given value.

### HasPaymentSyncDate

`func (o *UnpostedPayments) HasPaymentSyncDate() bool`

HasPaymentSyncDate returns a boolean if a field has been set.

### GetPaymentAccount

`func (o *UnpostedPayments) GetPaymentAccount() string`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *UnpostedPayments) GetPaymentAccountOk() (*string, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *UnpostedPayments) SetPaymentAccount(v string)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *UnpostedPayments) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### GetARPaymentAccount

`func (o *UnpostedPayments) GetARPaymentAccount() string`

GetARPaymentAccount returns the ARPaymentAccount field if non-nil, zero value otherwise.

### GetARPaymentAccountOk

`func (o *UnpostedPayments) GetARPaymentAccountOk() (*string, bool)`

GetARPaymentAccountOk returns a tuple with the ARPaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetARPaymentAccount

`func (o *UnpostedPayments) SetARPaymentAccount(v string)`

SetARPaymentAccount sets ARPaymentAccount field to given value.

### HasARPaymentAccount

`func (o *UnpostedPayments) HasARPaymentAccount() bool`

HasARPaymentAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


