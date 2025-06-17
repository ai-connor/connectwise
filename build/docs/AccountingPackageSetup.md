# AccountingPackageSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AccountingPackage** | [**AccountingPackageReference**](AccountingPackageReference.md) |  | 
**DirectTransferFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeInvoicesFlag** | Pointer to **NullableBool** |  | [optional] 
**InvoiceFormat** | Pointer to **NullableString** |  | [optional] 
**IncludeExpensesFlag** | Pointer to **NullableBool** |  | [optional] 
**TransferExpensesAsBillFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseFormat** | Pointer to **NullableString** |  | [optional] 
**SuppressMemoFlag** | Pointer to **NullableBool** |  | [optional] 
**SyncPaymentInfoFlag** | Pointer to **NullableBool** |  | [optional] 
**SyncWisePayPaymentInfoFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeSalesTaxFlag** | Pointer to **NullableBool** |  | [optional] 
**EnableTaxGroupsFlag** | Pointer to **NullableBool** |  | [optional] 
**ZeroDollarTaxAmountsFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeItemsFlag** | Pointer to **NullableBool** |  | [optional] 
**InventorySOHFlag** | Pointer to **NullableBool** |  | [optional] 
**SendComponentAmountFlag** | Pointer to **NullableBool** |  | [optional] 
**SendUomFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeCogsEntriesFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeCogsDropShipFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAccountingPackageSetup

`func NewAccountingPackageSetup(accountingPackage AccountingPackageReference, ) *AccountingPackageSetup`

NewAccountingPackageSetup instantiates a new AccountingPackageSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingPackageSetupWithDefaults

`func NewAccountingPackageSetupWithDefaults() *AccountingPackageSetup`

NewAccountingPackageSetupWithDefaults instantiates a new AccountingPackageSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountingPackageSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountingPackageSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountingPackageSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccountingPackageSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountingPackage

`func (o *AccountingPackageSetup) GetAccountingPackage() AccountingPackageReference`

GetAccountingPackage returns the AccountingPackage field if non-nil, zero value otherwise.

### GetAccountingPackageOk

`func (o *AccountingPackageSetup) GetAccountingPackageOk() (*AccountingPackageReference, bool)`

GetAccountingPackageOk returns a tuple with the AccountingPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingPackage

`func (o *AccountingPackageSetup) SetAccountingPackage(v AccountingPackageReference)`

SetAccountingPackage sets AccountingPackage field to given value.


### GetDirectTransferFlag

`func (o *AccountingPackageSetup) GetDirectTransferFlag() bool`

GetDirectTransferFlag returns the DirectTransferFlag field if non-nil, zero value otherwise.

### GetDirectTransferFlagOk

`func (o *AccountingPackageSetup) GetDirectTransferFlagOk() (*bool, bool)`

GetDirectTransferFlagOk returns a tuple with the DirectTransferFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectTransferFlag

`func (o *AccountingPackageSetup) SetDirectTransferFlag(v bool)`

SetDirectTransferFlag sets DirectTransferFlag field to given value.

### HasDirectTransferFlag

`func (o *AccountingPackageSetup) HasDirectTransferFlag() bool`

HasDirectTransferFlag returns a boolean if a field has been set.

### SetDirectTransferFlagNil

`func (o *AccountingPackageSetup) SetDirectTransferFlagNil(b bool)`

 SetDirectTransferFlagNil sets the value for DirectTransferFlag to be an explicit nil

### UnsetDirectTransferFlag
`func (o *AccountingPackageSetup) UnsetDirectTransferFlag()`

UnsetDirectTransferFlag ensures that no value is present for DirectTransferFlag, not even an explicit nil
### GetIncludeInvoicesFlag

`func (o *AccountingPackageSetup) GetIncludeInvoicesFlag() bool`

GetIncludeInvoicesFlag returns the IncludeInvoicesFlag field if non-nil, zero value otherwise.

### GetIncludeInvoicesFlagOk

`func (o *AccountingPackageSetup) GetIncludeInvoicesFlagOk() (*bool, bool)`

GetIncludeInvoicesFlagOk returns a tuple with the IncludeInvoicesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInvoicesFlag

`func (o *AccountingPackageSetup) SetIncludeInvoicesFlag(v bool)`

SetIncludeInvoicesFlag sets IncludeInvoicesFlag field to given value.

### HasIncludeInvoicesFlag

`func (o *AccountingPackageSetup) HasIncludeInvoicesFlag() bool`

HasIncludeInvoicesFlag returns a boolean if a field has been set.

### SetIncludeInvoicesFlagNil

`func (o *AccountingPackageSetup) SetIncludeInvoicesFlagNil(b bool)`

 SetIncludeInvoicesFlagNil sets the value for IncludeInvoicesFlag to be an explicit nil

### UnsetIncludeInvoicesFlag
`func (o *AccountingPackageSetup) UnsetIncludeInvoicesFlag()`

UnsetIncludeInvoicesFlag ensures that no value is present for IncludeInvoicesFlag, not even an explicit nil
### GetInvoiceFormat

`func (o *AccountingPackageSetup) GetInvoiceFormat() string`

GetInvoiceFormat returns the InvoiceFormat field if non-nil, zero value otherwise.

### GetInvoiceFormatOk

`func (o *AccountingPackageSetup) GetInvoiceFormatOk() (*string, bool)`

GetInvoiceFormatOk returns a tuple with the InvoiceFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceFormat

`func (o *AccountingPackageSetup) SetInvoiceFormat(v string)`

SetInvoiceFormat sets InvoiceFormat field to given value.

### HasInvoiceFormat

`func (o *AccountingPackageSetup) HasInvoiceFormat() bool`

HasInvoiceFormat returns a boolean if a field has been set.

### SetInvoiceFormatNil

`func (o *AccountingPackageSetup) SetInvoiceFormatNil(b bool)`

 SetInvoiceFormatNil sets the value for InvoiceFormat to be an explicit nil

### UnsetInvoiceFormat
`func (o *AccountingPackageSetup) UnsetInvoiceFormat()`

UnsetInvoiceFormat ensures that no value is present for InvoiceFormat, not even an explicit nil
### GetIncludeExpensesFlag

`func (o *AccountingPackageSetup) GetIncludeExpensesFlag() bool`

GetIncludeExpensesFlag returns the IncludeExpensesFlag field if non-nil, zero value otherwise.

### GetIncludeExpensesFlagOk

`func (o *AccountingPackageSetup) GetIncludeExpensesFlagOk() (*bool, bool)`

GetIncludeExpensesFlagOk returns a tuple with the IncludeExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExpensesFlag

`func (o *AccountingPackageSetup) SetIncludeExpensesFlag(v bool)`

SetIncludeExpensesFlag sets IncludeExpensesFlag field to given value.

### HasIncludeExpensesFlag

`func (o *AccountingPackageSetup) HasIncludeExpensesFlag() bool`

HasIncludeExpensesFlag returns a boolean if a field has been set.

### SetIncludeExpensesFlagNil

`func (o *AccountingPackageSetup) SetIncludeExpensesFlagNil(b bool)`

 SetIncludeExpensesFlagNil sets the value for IncludeExpensesFlag to be an explicit nil

### UnsetIncludeExpensesFlag
`func (o *AccountingPackageSetup) UnsetIncludeExpensesFlag()`

UnsetIncludeExpensesFlag ensures that no value is present for IncludeExpensesFlag, not even an explicit nil
### GetTransferExpensesAsBillFlag

`func (o *AccountingPackageSetup) GetTransferExpensesAsBillFlag() bool`

GetTransferExpensesAsBillFlag returns the TransferExpensesAsBillFlag field if non-nil, zero value otherwise.

### GetTransferExpensesAsBillFlagOk

`func (o *AccountingPackageSetup) GetTransferExpensesAsBillFlagOk() (*bool, bool)`

GetTransferExpensesAsBillFlagOk returns a tuple with the TransferExpensesAsBillFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferExpensesAsBillFlag

`func (o *AccountingPackageSetup) SetTransferExpensesAsBillFlag(v bool)`

SetTransferExpensesAsBillFlag sets TransferExpensesAsBillFlag field to given value.

### HasTransferExpensesAsBillFlag

`func (o *AccountingPackageSetup) HasTransferExpensesAsBillFlag() bool`

HasTransferExpensesAsBillFlag returns a boolean if a field has been set.

### SetTransferExpensesAsBillFlagNil

`func (o *AccountingPackageSetup) SetTransferExpensesAsBillFlagNil(b bool)`

 SetTransferExpensesAsBillFlagNil sets the value for TransferExpensesAsBillFlag to be an explicit nil

### UnsetTransferExpensesAsBillFlag
`func (o *AccountingPackageSetup) UnsetTransferExpensesAsBillFlag()`

UnsetTransferExpensesAsBillFlag ensures that no value is present for TransferExpensesAsBillFlag, not even an explicit nil
### GetExpenseFormat

`func (o *AccountingPackageSetup) GetExpenseFormat() string`

GetExpenseFormat returns the ExpenseFormat field if non-nil, zero value otherwise.

### GetExpenseFormatOk

`func (o *AccountingPackageSetup) GetExpenseFormatOk() (*string, bool)`

GetExpenseFormatOk returns a tuple with the ExpenseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseFormat

`func (o *AccountingPackageSetup) SetExpenseFormat(v string)`

SetExpenseFormat sets ExpenseFormat field to given value.

### HasExpenseFormat

`func (o *AccountingPackageSetup) HasExpenseFormat() bool`

HasExpenseFormat returns a boolean if a field has been set.

### SetExpenseFormatNil

`func (o *AccountingPackageSetup) SetExpenseFormatNil(b bool)`

 SetExpenseFormatNil sets the value for ExpenseFormat to be an explicit nil

### UnsetExpenseFormat
`func (o *AccountingPackageSetup) UnsetExpenseFormat()`

UnsetExpenseFormat ensures that no value is present for ExpenseFormat, not even an explicit nil
### GetSuppressMemoFlag

`func (o *AccountingPackageSetup) GetSuppressMemoFlag() bool`

GetSuppressMemoFlag returns the SuppressMemoFlag field if non-nil, zero value otherwise.

### GetSuppressMemoFlagOk

`func (o *AccountingPackageSetup) GetSuppressMemoFlagOk() (*bool, bool)`

GetSuppressMemoFlagOk returns a tuple with the SuppressMemoFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressMemoFlag

`func (o *AccountingPackageSetup) SetSuppressMemoFlag(v bool)`

SetSuppressMemoFlag sets SuppressMemoFlag field to given value.

### HasSuppressMemoFlag

`func (o *AccountingPackageSetup) HasSuppressMemoFlag() bool`

HasSuppressMemoFlag returns a boolean if a field has been set.

### SetSuppressMemoFlagNil

`func (o *AccountingPackageSetup) SetSuppressMemoFlagNil(b bool)`

 SetSuppressMemoFlagNil sets the value for SuppressMemoFlag to be an explicit nil

### UnsetSuppressMemoFlag
`func (o *AccountingPackageSetup) UnsetSuppressMemoFlag()`

UnsetSuppressMemoFlag ensures that no value is present for SuppressMemoFlag, not even an explicit nil
### GetSyncPaymentInfoFlag

`func (o *AccountingPackageSetup) GetSyncPaymentInfoFlag() bool`

GetSyncPaymentInfoFlag returns the SyncPaymentInfoFlag field if non-nil, zero value otherwise.

### GetSyncPaymentInfoFlagOk

`func (o *AccountingPackageSetup) GetSyncPaymentInfoFlagOk() (*bool, bool)`

GetSyncPaymentInfoFlagOk returns a tuple with the SyncPaymentInfoFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPaymentInfoFlag

`func (o *AccountingPackageSetup) SetSyncPaymentInfoFlag(v bool)`

SetSyncPaymentInfoFlag sets SyncPaymentInfoFlag field to given value.

### HasSyncPaymentInfoFlag

`func (o *AccountingPackageSetup) HasSyncPaymentInfoFlag() bool`

HasSyncPaymentInfoFlag returns a boolean if a field has been set.

### SetSyncPaymentInfoFlagNil

`func (o *AccountingPackageSetup) SetSyncPaymentInfoFlagNil(b bool)`

 SetSyncPaymentInfoFlagNil sets the value for SyncPaymentInfoFlag to be an explicit nil

### UnsetSyncPaymentInfoFlag
`func (o *AccountingPackageSetup) UnsetSyncPaymentInfoFlag()`

UnsetSyncPaymentInfoFlag ensures that no value is present for SyncPaymentInfoFlag, not even an explicit nil
### GetSyncWisePayPaymentInfoFlag

`func (o *AccountingPackageSetup) GetSyncWisePayPaymentInfoFlag() bool`

GetSyncWisePayPaymentInfoFlag returns the SyncWisePayPaymentInfoFlag field if non-nil, zero value otherwise.

### GetSyncWisePayPaymentInfoFlagOk

`func (o *AccountingPackageSetup) GetSyncWisePayPaymentInfoFlagOk() (*bool, bool)`

GetSyncWisePayPaymentInfoFlagOk returns a tuple with the SyncWisePayPaymentInfoFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncWisePayPaymentInfoFlag

`func (o *AccountingPackageSetup) SetSyncWisePayPaymentInfoFlag(v bool)`

SetSyncWisePayPaymentInfoFlag sets SyncWisePayPaymentInfoFlag field to given value.

### HasSyncWisePayPaymentInfoFlag

`func (o *AccountingPackageSetup) HasSyncWisePayPaymentInfoFlag() bool`

HasSyncWisePayPaymentInfoFlag returns a boolean if a field has been set.

### SetSyncWisePayPaymentInfoFlagNil

`func (o *AccountingPackageSetup) SetSyncWisePayPaymentInfoFlagNil(b bool)`

 SetSyncWisePayPaymentInfoFlagNil sets the value for SyncWisePayPaymentInfoFlag to be an explicit nil

### UnsetSyncWisePayPaymentInfoFlag
`func (o *AccountingPackageSetup) UnsetSyncWisePayPaymentInfoFlag()`

UnsetSyncWisePayPaymentInfoFlag ensures that no value is present for SyncWisePayPaymentInfoFlag, not even an explicit nil
### GetIncludeSalesTaxFlag

`func (o *AccountingPackageSetup) GetIncludeSalesTaxFlag() bool`

GetIncludeSalesTaxFlag returns the IncludeSalesTaxFlag field if non-nil, zero value otherwise.

### GetIncludeSalesTaxFlagOk

`func (o *AccountingPackageSetup) GetIncludeSalesTaxFlagOk() (*bool, bool)`

GetIncludeSalesTaxFlagOk returns a tuple with the IncludeSalesTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSalesTaxFlag

`func (o *AccountingPackageSetup) SetIncludeSalesTaxFlag(v bool)`

SetIncludeSalesTaxFlag sets IncludeSalesTaxFlag field to given value.

### HasIncludeSalesTaxFlag

`func (o *AccountingPackageSetup) HasIncludeSalesTaxFlag() bool`

HasIncludeSalesTaxFlag returns a boolean if a field has been set.

### SetIncludeSalesTaxFlagNil

`func (o *AccountingPackageSetup) SetIncludeSalesTaxFlagNil(b bool)`

 SetIncludeSalesTaxFlagNil sets the value for IncludeSalesTaxFlag to be an explicit nil

### UnsetIncludeSalesTaxFlag
`func (o *AccountingPackageSetup) UnsetIncludeSalesTaxFlag()`

UnsetIncludeSalesTaxFlag ensures that no value is present for IncludeSalesTaxFlag, not even an explicit nil
### GetEnableTaxGroupsFlag

`func (o *AccountingPackageSetup) GetEnableTaxGroupsFlag() bool`

GetEnableTaxGroupsFlag returns the EnableTaxGroupsFlag field if non-nil, zero value otherwise.

### GetEnableTaxGroupsFlagOk

`func (o *AccountingPackageSetup) GetEnableTaxGroupsFlagOk() (*bool, bool)`

GetEnableTaxGroupsFlagOk returns a tuple with the EnableTaxGroupsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaxGroupsFlag

`func (o *AccountingPackageSetup) SetEnableTaxGroupsFlag(v bool)`

SetEnableTaxGroupsFlag sets EnableTaxGroupsFlag field to given value.

### HasEnableTaxGroupsFlag

`func (o *AccountingPackageSetup) HasEnableTaxGroupsFlag() bool`

HasEnableTaxGroupsFlag returns a boolean if a field has been set.

### SetEnableTaxGroupsFlagNil

`func (o *AccountingPackageSetup) SetEnableTaxGroupsFlagNil(b bool)`

 SetEnableTaxGroupsFlagNil sets the value for EnableTaxGroupsFlag to be an explicit nil

### UnsetEnableTaxGroupsFlag
`func (o *AccountingPackageSetup) UnsetEnableTaxGroupsFlag()`

UnsetEnableTaxGroupsFlag ensures that no value is present for EnableTaxGroupsFlag, not even an explicit nil
### GetZeroDollarTaxAmountsFlag

`func (o *AccountingPackageSetup) GetZeroDollarTaxAmountsFlag() bool`

GetZeroDollarTaxAmountsFlag returns the ZeroDollarTaxAmountsFlag field if non-nil, zero value otherwise.

### GetZeroDollarTaxAmountsFlagOk

`func (o *AccountingPackageSetup) GetZeroDollarTaxAmountsFlagOk() (*bool, bool)`

GetZeroDollarTaxAmountsFlagOk returns a tuple with the ZeroDollarTaxAmountsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroDollarTaxAmountsFlag

`func (o *AccountingPackageSetup) SetZeroDollarTaxAmountsFlag(v bool)`

SetZeroDollarTaxAmountsFlag sets ZeroDollarTaxAmountsFlag field to given value.

### HasZeroDollarTaxAmountsFlag

`func (o *AccountingPackageSetup) HasZeroDollarTaxAmountsFlag() bool`

HasZeroDollarTaxAmountsFlag returns a boolean if a field has been set.

### SetZeroDollarTaxAmountsFlagNil

`func (o *AccountingPackageSetup) SetZeroDollarTaxAmountsFlagNil(b bool)`

 SetZeroDollarTaxAmountsFlagNil sets the value for ZeroDollarTaxAmountsFlag to be an explicit nil

### UnsetZeroDollarTaxAmountsFlag
`func (o *AccountingPackageSetup) UnsetZeroDollarTaxAmountsFlag()`

UnsetZeroDollarTaxAmountsFlag ensures that no value is present for ZeroDollarTaxAmountsFlag, not even an explicit nil
### GetIncludeItemsFlag

`func (o *AccountingPackageSetup) GetIncludeItemsFlag() bool`

GetIncludeItemsFlag returns the IncludeItemsFlag field if non-nil, zero value otherwise.

### GetIncludeItemsFlagOk

`func (o *AccountingPackageSetup) GetIncludeItemsFlagOk() (*bool, bool)`

GetIncludeItemsFlagOk returns a tuple with the IncludeItemsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeItemsFlag

`func (o *AccountingPackageSetup) SetIncludeItemsFlag(v bool)`

SetIncludeItemsFlag sets IncludeItemsFlag field to given value.

### HasIncludeItemsFlag

`func (o *AccountingPackageSetup) HasIncludeItemsFlag() bool`

HasIncludeItemsFlag returns a boolean if a field has been set.

### SetIncludeItemsFlagNil

`func (o *AccountingPackageSetup) SetIncludeItemsFlagNil(b bool)`

 SetIncludeItemsFlagNil sets the value for IncludeItemsFlag to be an explicit nil

### UnsetIncludeItemsFlag
`func (o *AccountingPackageSetup) UnsetIncludeItemsFlag()`

UnsetIncludeItemsFlag ensures that no value is present for IncludeItemsFlag, not even an explicit nil
### GetInventorySOHFlag

`func (o *AccountingPackageSetup) GetInventorySOHFlag() bool`

GetInventorySOHFlag returns the InventorySOHFlag field if non-nil, zero value otherwise.

### GetInventorySOHFlagOk

`func (o *AccountingPackageSetup) GetInventorySOHFlagOk() (*bool, bool)`

GetInventorySOHFlagOk returns a tuple with the InventorySOHFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySOHFlag

`func (o *AccountingPackageSetup) SetInventorySOHFlag(v bool)`

SetInventorySOHFlag sets InventorySOHFlag field to given value.

### HasInventorySOHFlag

`func (o *AccountingPackageSetup) HasInventorySOHFlag() bool`

HasInventorySOHFlag returns a boolean if a field has been set.

### SetInventorySOHFlagNil

`func (o *AccountingPackageSetup) SetInventorySOHFlagNil(b bool)`

 SetInventorySOHFlagNil sets the value for InventorySOHFlag to be an explicit nil

### UnsetInventorySOHFlag
`func (o *AccountingPackageSetup) UnsetInventorySOHFlag()`

UnsetInventorySOHFlag ensures that no value is present for InventorySOHFlag, not even an explicit nil
### GetSendComponentAmountFlag

`func (o *AccountingPackageSetup) GetSendComponentAmountFlag() bool`

GetSendComponentAmountFlag returns the SendComponentAmountFlag field if non-nil, zero value otherwise.

### GetSendComponentAmountFlagOk

`func (o *AccountingPackageSetup) GetSendComponentAmountFlagOk() (*bool, bool)`

GetSendComponentAmountFlagOk returns a tuple with the SendComponentAmountFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendComponentAmountFlag

`func (o *AccountingPackageSetup) SetSendComponentAmountFlag(v bool)`

SetSendComponentAmountFlag sets SendComponentAmountFlag field to given value.

### HasSendComponentAmountFlag

`func (o *AccountingPackageSetup) HasSendComponentAmountFlag() bool`

HasSendComponentAmountFlag returns a boolean if a field has been set.

### SetSendComponentAmountFlagNil

`func (o *AccountingPackageSetup) SetSendComponentAmountFlagNil(b bool)`

 SetSendComponentAmountFlagNil sets the value for SendComponentAmountFlag to be an explicit nil

### UnsetSendComponentAmountFlag
`func (o *AccountingPackageSetup) UnsetSendComponentAmountFlag()`

UnsetSendComponentAmountFlag ensures that no value is present for SendComponentAmountFlag, not even an explicit nil
### GetSendUomFlag

`func (o *AccountingPackageSetup) GetSendUomFlag() bool`

GetSendUomFlag returns the SendUomFlag field if non-nil, zero value otherwise.

### GetSendUomFlagOk

`func (o *AccountingPackageSetup) GetSendUomFlagOk() (*bool, bool)`

GetSendUomFlagOk returns a tuple with the SendUomFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendUomFlag

`func (o *AccountingPackageSetup) SetSendUomFlag(v bool)`

SetSendUomFlag sets SendUomFlag field to given value.

### HasSendUomFlag

`func (o *AccountingPackageSetup) HasSendUomFlag() bool`

HasSendUomFlag returns a boolean if a field has been set.

### SetSendUomFlagNil

`func (o *AccountingPackageSetup) SetSendUomFlagNil(b bool)`

 SetSendUomFlagNil sets the value for SendUomFlag to be an explicit nil

### UnsetSendUomFlag
`func (o *AccountingPackageSetup) UnsetSendUomFlag()`

UnsetSendUomFlag ensures that no value is present for SendUomFlag, not even an explicit nil
### GetIncludeCogsEntriesFlag

`func (o *AccountingPackageSetup) GetIncludeCogsEntriesFlag() bool`

GetIncludeCogsEntriesFlag returns the IncludeCogsEntriesFlag field if non-nil, zero value otherwise.

### GetIncludeCogsEntriesFlagOk

`func (o *AccountingPackageSetup) GetIncludeCogsEntriesFlagOk() (*bool, bool)`

GetIncludeCogsEntriesFlagOk returns a tuple with the IncludeCogsEntriesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCogsEntriesFlag

`func (o *AccountingPackageSetup) SetIncludeCogsEntriesFlag(v bool)`

SetIncludeCogsEntriesFlag sets IncludeCogsEntriesFlag field to given value.

### HasIncludeCogsEntriesFlag

`func (o *AccountingPackageSetup) HasIncludeCogsEntriesFlag() bool`

HasIncludeCogsEntriesFlag returns a boolean if a field has been set.

### SetIncludeCogsEntriesFlagNil

`func (o *AccountingPackageSetup) SetIncludeCogsEntriesFlagNil(b bool)`

 SetIncludeCogsEntriesFlagNil sets the value for IncludeCogsEntriesFlag to be an explicit nil

### UnsetIncludeCogsEntriesFlag
`func (o *AccountingPackageSetup) UnsetIncludeCogsEntriesFlag()`

UnsetIncludeCogsEntriesFlag ensures that no value is present for IncludeCogsEntriesFlag, not even an explicit nil
### GetIncludeCogsDropShipFlag

`func (o *AccountingPackageSetup) GetIncludeCogsDropShipFlag() bool`

GetIncludeCogsDropShipFlag returns the IncludeCogsDropShipFlag field if non-nil, zero value otherwise.

### GetIncludeCogsDropShipFlagOk

`func (o *AccountingPackageSetup) GetIncludeCogsDropShipFlagOk() (*bool, bool)`

GetIncludeCogsDropShipFlagOk returns a tuple with the IncludeCogsDropShipFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCogsDropShipFlag

`func (o *AccountingPackageSetup) SetIncludeCogsDropShipFlag(v bool)`

SetIncludeCogsDropShipFlag sets IncludeCogsDropShipFlag field to given value.

### HasIncludeCogsDropShipFlag

`func (o *AccountingPackageSetup) HasIncludeCogsDropShipFlag() bool`

HasIncludeCogsDropShipFlag returns a boolean if a field has been set.

### SetIncludeCogsDropShipFlagNil

`func (o *AccountingPackageSetup) SetIncludeCogsDropShipFlagNil(b bool)`

 SetIncludeCogsDropShipFlagNil sets the value for IncludeCogsDropShipFlag to be an explicit nil

### UnsetIncludeCogsDropShipFlag
`func (o *AccountingPackageSetup) UnsetIncludeCogsDropShipFlag()`

UnsetIncludeCogsDropShipFlag ensures that no value is present for IncludeCogsDropShipFlag, not even an explicit nil
### GetInfo

`func (o *AccountingPackageSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AccountingPackageSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AccountingPackageSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AccountingPackageSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


