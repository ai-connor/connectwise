# BatchEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**Debit** | Pointer to **NullableFloat64** |  | [optional] 
**Credit** | Pointer to **NullableFloat64** |  | [optional] 
**Cost** | Pointer to **float64** |  | [optional] 
**Item** | Pointer to **string** |  | [optional] 
**SalesCode** | Pointer to **string** |  | [optional] 
**CostOfGoodsSoldAccountNumber** | Pointer to **string** |  | [optional] 
**Invoice** | Pointer to [**InvoiceReference**](InvoiceReference.md) |  | [optional] 
**PurchaseOrder** | Pointer to [**PurchaseOrderReference**](PurchaseOrderReference.md) |  | [optional] 
**LineItem** | Pointer to [**PurchaseOrderLineItemReference**](PurchaseOrderLineItemReference.md) |  | [optional] 
**Transfer** | Pointer to **string** |  | [optional] 
**Expense** | Pointer to [**ExpenseDetailReference**](ExpenseDetailReference.md) |  | [optional] 
**AdjustmentDetail** | Pointer to [**AdjustmentDetailReference**](AdjustmentDetailReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBatchEntry

`func NewBatchEntry() *BatchEntry`

NewBatchEntry instantiates a new BatchEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchEntryWithDefaults

`func NewBatchEntryWithDefaults() *BatchEntry`

NewBatchEntryWithDefaults instantiates a new BatchEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BatchEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountType

`func (o *BatchEntry) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BatchEntry) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BatchEntry) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *BatchEntry) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetName

`func (o *BatchEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BatchEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BatchEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BatchEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountNumber

`func (o *BatchEntry) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BatchEntry) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BatchEntry) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *BatchEntry) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetDebit

`func (o *BatchEntry) GetDebit() float64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *BatchEntry) GetDebitOk() (*float64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *BatchEntry) SetDebit(v float64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *BatchEntry) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### SetDebitNil

`func (o *BatchEntry) SetDebitNil(b bool)`

 SetDebitNil sets the value for Debit to be an explicit nil

### UnsetDebit
`func (o *BatchEntry) UnsetDebit()`

UnsetDebit ensures that no value is present for Debit, not even an explicit nil
### GetCredit

`func (o *BatchEntry) GetCredit() float64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *BatchEntry) GetCreditOk() (*float64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *BatchEntry) SetCredit(v float64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *BatchEntry) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### SetCreditNil

`func (o *BatchEntry) SetCreditNil(b bool)`

 SetCreditNil sets the value for Credit to be an explicit nil

### UnsetCredit
`func (o *BatchEntry) UnsetCredit()`

UnsetCredit ensures that no value is present for Credit, not even an explicit nil
### GetCost

`func (o *BatchEntry) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BatchEntry) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BatchEntry) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BatchEntry) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetItem

`func (o *BatchEntry) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *BatchEntry) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *BatchEntry) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *BatchEntry) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetSalesCode

`func (o *BatchEntry) GetSalesCode() string`

GetSalesCode returns the SalesCode field if non-nil, zero value otherwise.

### GetSalesCodeOk

`func (o *BatchEntry) GetSalesCodeOk() (*string, bool)`

GetSalesCodeOk returns a tuple with the SalesCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesCode

`func (o *BatchEntry) SetSalesCode(v string)`

SetSalesCode sets SalesCode field to given value.

### HasSalesCode

`func (o *BatchEntry) HasSalesCode() bool`

HasSalesCode returns a boolean if a field has been set.

### GetCostOfGoodsSoldAccountNumber

`func (o *BatchEntry) GetCostOfGoodsSoldAccountNumber() string`

GetCostOfGoodsSoldAccountNumber returns the CostOfGoodsSoldAccountNumber field if non-nil, zero value otherwise.

### GetCostOfGoodsSoldAccountNumberOk

`func (o *BatchEntry) GetCostOfGoodsSoldAccountNumberOk() (*string, bool)`

GetCostOfGoodsSoldAccountNumberOk returns a tuple with the CostOfGoodsSoldAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOfGoodsSoldAccountNumber

`func (o *BatchEntry) SetCostOfGoodsSoldAccountNumber(v string)`

SetCostOfGoodsSoldAccountNumber sets CostOfGoodsSoldAccountNumber field to given value.

### HasCostOfGoodsSoldAccountNumber

`func (o *BatchEntry) HasCostOfGoodsSoldAccountNumber() bool`

HasCostOfGoodsSoldAccountNumber returns a boolean if a field has been set.

### GetInvoice

`func (o *BatchEntry) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *BatchEntry) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *BatchEntry) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *BatchEntry) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetPurchaseOrder

`func (o *BatchEntry) GetPurchaseOrder() PurchaseOrderReference`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *BatchEntry) GetPurchaseOrderOk() (*PurchaseOrderReference, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *BatchEntry) SetPurchaseOrder(v PurchaseOrderReference)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *BatchEntry) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.

### GetLineItem

`func (o *BatchEntry) GetLineItem() PurchaseOrderLineItemReference`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *BatchEntry) GetLineItemOk() (*PurchaseOrderLineItemReference, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *BatchEntry) SetLineItem(v PurchaseOrderLineItemReference)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *BatchEntry) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.

### GetTransfer

`func (o *BatchEntry) GetTransfer() string`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *BatchEntry) GetTransferOk() (*string, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *BatchEntry) SetTransfer(v string)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *BatchEntry) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.

### GetExpense

`func (o *BatchEntry) GetExpense() ExpenseDetailReference`

GetExpense returns the Expense field if non-nil, zero value otherwise.

### GetExpenseOk

`func (o *BatchEntry) GetExpenseOk() (*ExpenseDetailReference, bool)`

GetExpenseOk returns a tuple with the Expense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpense

`func (o *BatchEntry) SetExpense(v ExpenseDetailReference)`

SetExpense sets Expense field to given value.

### HasExpense

`func (o *BatchEntry) HasExpense() bool`

HasExpense returns a boolean if a field has been set.

### GetAdjustmentDetail

`func (o *BatchEntry) GetAdjustmentDetail() AdjustmentDetailReference`

GetAdjustmentDetail returns the AdjustmentDetail field if non-nil, zero value otherwise.

### GetAdjustmentDetailOk

`func (o *BatchEntry) GetAdjustmentDetailOk() (*AdjustmentDetailReference, bool)`

GetAdjustmentDetailOk returns a tuple with the AdjustmentDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentDetail

`func (o *BatchEntry) SetAdjustmentDetail(v AdjustmentDetailReference)`

SetAdjustmentDetail sets AdjustmentDetail field to given value.

### HasAdjustmentDetail

`func (o *BatchEntry) HasAdjustmentDetail() bool`

HasAdjustmentDetail returns a boolean if a field has been set.

### GetInfo

`func (o *BatchEntry) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BatchEntry) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BatchEntry) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BatchEntry) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


