# InvoicePayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Invoice** | Pointer to [**InvoiceReference**](InvoiceReference.md) |  | [optional] 
**Credit** | Pointer to [**InvoiceReference**](InvoiceReference.md) |  | [optional] 
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**Balance** | Pointer to **NullableFloat64** |  | [optional] 
**InvoiceBalance** | Pointer to **NullableFloat64** |  | [optional] 
**InvoiceTotal** | Pointer to **NullableFloat64** |  | [optional] 
**PaymentDate** | Pointer to **time.Time** |  | [optional] 
**AppliedBy** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**WisePayPayment** | Pointer to [**WisePayPayment**](WisePayPayment.md) |  | [optional] 
**PaymentSyncStatus** | Pointer to **string** |  | [optional] 
**GlBatchID** | Pointer to **string** |  Max length: 50; | [optional] 
**PaymentSyncDate** | Pointer to **time.Time** |  | [optional] 
**PaymentAccount** | Pointer to **string** |  | [optional] 
**ARPaymentAccount** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoicePayment

`func NewInvoicePayment() *InvoicePayment`

NewInvoicePayment instantiates a new InvoicePayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePaymentWithDefaults

`func NewInvoicePaymentWithDefaults() *InvoicePayment`

NewInvoicePaymentWithDefaults instantiates a new InvoicePayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoicePayment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoicePayment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoicePayment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoicePayment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *InvoicePayment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoicePayment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoicePayment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvoicePayment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSource

`func (o *InvoicePayment) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InvoicePayment) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InvoicePayment) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InvoicePayment) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetInvoice

`func (o *InvoicePayment) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *InvoicePayment) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *InvoicePayment) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *InvoicePayment) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetCredit

`func (o *InvoicePayment) GetCredit() InvoiceReference`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *InvoicePayment) GetCreditOk() (*InvoiceReference, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *InvoicePayment) SetCredit(v InvoiceReference)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *InvoicePayment) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetAmount

`func (o *InvoicePayment) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoicePayment) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoicePayment) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InvoicePayment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *InvoicePayment) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *InvoicePayment) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetBalance

`func (o *InvoicePayment) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *InvoicePayment) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *InvoicePayment) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *InvoicePayment) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalanceNil

`func (o *InvoicePayment) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *InvoicePayment) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetInvoiceBalance

`func (o *InvoicePayment) GetInvoiceBalance() float64`

GetInvoiceBalance returns the InvoiceBalance field if non-nil, zero value otherwise.

### GetInvoiceBalanceOk

`func (o *InvoicePayment) GetInvoiceBalanceOk() (*float64, bool)`

GetInvoiceBalanceOk returns a tuple with the InvoiceBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceBalance

`func (o *InvoicePayment) SetInvoiceBalance(v float64)`

SetInvoiceBalance sets InvoiceBalance field to given value.

### HasInvoiceBalance

`func (o *InvoicePayment) HasInvoiceBalance() bool`

HasInvoiceBalance returns a boolean if a field has been set.

### SetInvoiceBalanceNil

`func (o *InvoicePayment) SetInvoiceBalanceNil(b bool)`

 SetInvoiceBalanceNil sets the value for InvoiceBalance to be an explicit nil

### UnsetInvoiceBalance
`func (o *InvoicePayment) UnsetInvoiceBalance()`

UnsetInvoiceBalance ensures that no value is present for InvoiceBalance, not even an explicit nil
### GetInvoiceTotal

`func (o *InvoicePayment) GetInvoiceTotal() float64`

GetInvoiceTotal returns the InvoiceTotal field if non-nil, zero value otherwise.

### GetInvoiceTotalOk

`func (o *InvoicePayment) GetInvoiceTotalOk() (*float64, bool)`

GetInvoiceTotalOk returns a tuple with the InvoiceTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTotal

`func (o *InvoicePayment) SetInvoiceTotal(v float64)`

SetInvoiceTotal sets InvoiceTotal field to given value.

### HasInvoiceTotal

`func (o *InvoicePayment) HasInvoiceTotal() bool`

HasInvoiceTotal returns a boolean if a field has been set.

### SetInvoiceTotalNil

`func (o *InvoicePayment) SetInvoiceTotalNil(b bool)`

 SetInvoiceTotalNil sets the value for InvoiceTotal to be an explicit nil

### UnsetInvoiceTotal
`func (o *InvoicePayment) UnsetInvoiceTotal()`

UnsetInvoiceTotal ensures that no value is present for InvoiceTotal, not even an explicit nil
### GetPaymentDate

`func (o *InvoicePayment) GetPaymentDate() time.Time`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *InvoicePayment) GetPaymentDateOk() (*time.Time, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *InvoicePayment) SetPaymentDate(v time.Time)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *InvoicePayment) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetAppliedBy

`func (o *InvoicePayment) GetAppliedBy() string`

GetAppliedBy returns the AppliedBy field if non-nil, zero value otherwise.

### GetAppliedByOk

`func (o *InvoicePayment) GetAppliedByOk() (*string, bool)`

GetAppliedByOk returns a tuple with the AppliedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedBy

`func (o *InvoicePayment) SetAppliedBy(v string)`

SetAppliedBy sets AppliedBy field to given value.

### HasAppliedBy

`func (o *InvoicePayment) HasAppliedBy() bool`

HasAppliedBy returns a boolean if a field has been set.

### GetInfo

`func (o *InvoicePayment) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InvoicePayment) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InvoicePayment) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InvoicePayment) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetWisePayPayment

`func (o *InvoicePayment) GetWisePayPayment() WisePayPayment`

GetWisePayPayment returns the WisePayPayment field if non-nil, zero value otherwise.

### GetWisePayPaymentOk

`func (o *InvoicePayment) GetWisePayPaymentOk() (*WisePayPayment, bool)`

GetWisePayPaymentOk returns a tuple with the WisePayPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWisePayPayment

`func (o *InvoicePayment) SetWisePayPayment(v WisePayPayment)`

SetWisePayPayment sets WisePayPayment field to given value.

### HasWisePayPayment

`func (o *InvoicePayment) HasWisePayPayment() bool`

HasWisePayPayment returns a boolean if a field has been set.

### GetPaymentSyncStatus

`func (o *InvoicePayment) GetPaymentSyncStatus() string`

GetPaymentSyncStatus returns the PaymentSyncStatus field if non-nil, zero value otherwise.

### GetPaymentSyncStatusOk

`func (o *InvoicePayment) GetPaymentSyncStatusOk() (*string, bool)`

GetPaymentSyncStatusOk returns a tuple with the PaymentSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSyncStatus

`func (o *InvoicePayment) SetPaymentSyncStatus(v string)`

SetPaymentSyncStatus sets PaymentSyncStatus field to given value.

### HasPaymentSyncStatus

`func (o *InvoicePayment) HasPaymentSyncStatus() bool`

HasPaymentSyncStatus returns a boolean if a field has been set.

### GetGlBatchID

`func (o *InvoicePayment) GetGlBatchID() string`

GetGlBatchID returns the GlBatchID field if non-nil, zero value otherwise.

### GetGlBatchIDOk

`func (o *InvoicePayment) GetGlBatchIDOk() (*string, bool)`

GetGlBatchIDOk returns a tuple with the GlBatchID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlBatchID

`func (o *InvoicePayment) SetGlBatchID(v string)`

SetGlBatchID sets GlBatchID field to given value.

### HasGlBatchID

`func (o *InvoicePayment) HasGlBatchID() bool`

HasGlBatchID returns a boolean if a field has been set.

### GetPaymentSyncDate

`func (o *InvoicePayment) GetPaymentSyncDate() time.Time`

GetPaymentSyncDate returns the PaymentSyncDate field if non-nil, zero value otherwise.

### GetPaymentSyncDateOk

`func (o *InvoicePayment) GetPaymentSyncDateOk() (*time.Time, bool)`

GetPaymentSyncDateOk returns a tuple with the PaymentSyncDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSyncDate

`func (o *InvoicePayment) SetPaymentSyncDate(v time.Time)`

SetPaymentSyncDate sets PaymentSyncDate field to given value.

### HasPaymentSyncDate

`func (o *InvoicePayment) HasPaymentSyncDate() bool`

HasPaymentSyncDate returns a boolean if a field has been set.

### GetPaymentAccount

`func (o *InvoicePayment) GetPaymentAccount() string`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *InvoicePayment) GetPaymentAccountOk() (*string, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *InvoicePayment) SetPaymentAccount(v string)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *InvoicePayment) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### GetARPaymentAccount

`func (o *InvoicePayment) GetARPaymentAccount() string`

GetARPaymentAccount returns the ARPaymentAccount field if non-nil, zero value otherwise.

### GetARPaymentAccountOk

`func (o *InvoicePayment) GetARPaymentAccountOk() (*string, bool)`

GetARPaymentAccountOk returns a tuple with the ARPaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetARPaymentAccount

`func (o *InvoicePayment) SetARPaymentAccount(v string)`

SetARPaymentAccount sets ARPaymentAccount field to given value.

### HasARPaymentAccount

`func (o *InvoicePayment) HasARPaymentAccount() bool`

HasARPaymentAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


