# GLExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**Vendors** | Pointer to [**[]GLExportVendor**](GLExportVendor.md) |  | [optional] 
**Customers** | Pointer to [**[]GLExportCustomer**](GLExportCustomer.md) |  | [optional] 
**Transactions** | Pointer to [**[]GLExportTransaction**](GLExportTransaction.md) |  | [optional] 
**Expenses** | Pointer to [**[]GLExportExpense**](GLExportExpense.md) |  | [optional] 
**ExpenseBills** | Pointer to [**[]GLExportExpenseBill**](GLExportExpenseBill.md) |  | [optional] 
**PurchaseTransactions** | Pointer to [**[]GLExportPurchaseTransaction**](GLExportPurchaseTransaction.md) |  | [optional] 
**AdjustmentTransactions** | Pointer to [**[]GLExportAdjustmentTransaction**](GLExportAdjustmentTransaction.md) |  | [optional] 
**InventoryTransfers** | Pointer to [**[]GLExportInventoryTransfer**](GLExportInventoryTransfer.md) |  | [optional] 

## Methods

### NewGLExport

`func NewGLExport() *GLExport`

NewGLExport instantiates a new GLExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportWithDefaults

`func NewGLExportWithDefaults() *GLExport`

NewGLExportWithDefaults instantiates a new GLExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportSettings

`func (o *GLExport) GetExportSettings() map[string]interface{}`

GetExportSettings returns the ExportSettings field if non-nil, zero value otherwise.

### GetExportSettingsOk

`func (o *GLExport) GetExportSettingsOk() (*map[string]interface{}, bool)`

GetExportSettingsOk returns a tuple with the ExportSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportSettings

`func (o *GLExport) SetExportSettings(v map[string]interface{})`

SetExportSettings sets ExportSettings field to given value.

### HasExportSettings

`func (o *GLExport) HasExportSettings() bool`

HasExportSettings returns a boolean if a field has been set.

### GetVendors

`func (o *GLExport) GetVendors() []GLExportVendor`

GetVendors returns the Vendors field if non-nil, zero value otherwise.

### GetVendorsOk

`func (o *GLExport) GetVendorsOk() (*[]GLExportVendor, bool)`

GetVendorsOk returns a tuple with the Vendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendors

`func (o *GLExport) SetVendors(v []GLExportVendor)`

SetVendors sets Vendors field to given value.

### HasVendors

`func (o *GLExport) HasVendors() bool`

HasVendors returns a boolean if a field has been set.

### GetCustomers

`func (o *GLExport) GetCustomers() []GLExportCustomer`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *GLExport) GetCustomersOk() (*[]GLExportCustomer, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *GLExport) SetCustomers(v []GLExportCustomer)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *GLExport) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetTransactions

`func (o *GLExport) GetTransactions() []GLExportTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *GLExport) GetTransactionsOk() (*[]GLExportTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *GLExport) SetTransactions(v []GLExportTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *GLExport) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetExpenses

`func (o *GLExport) GetExpenses() []GLExportExpense`

GetExpenses returns the Expenses field if non-nil, zero value otherwise.

### GetExpensesOk

`func (o *GLExport) GetExpensesOk() (*[]GLExportExpense, bool)`

GetExpensesOk returns a tuple with the Expenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenses

`func (o *GLExport) SetExpenses(v []GLExportExpense)`

SetExpenses sets Expenses field to given value.

### HasExpenses

`func (o *GLExport) HasExpenses() bool`

HasExpenses returns a boolean if a field has been set.

### GetExpenseBills

`func (o *GLExport) GetExpenseBills() []GLExportExpenseBill`

GetExpenseBills returns the ExpenseBills field if non-nil, zero value otherwise.

### GetExpenseBillsOk

`func (o *GLExport) GetExpenseBillsOk() (*[]GLExportExpenseBill, bool)`

GetExpenseBillsOk returns a tuple with the ExpenseBills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseBills

`func (o *GLExport) SetExpenseBills(v []GLExportExpenseBill)`

SetExpenseBills sets ExpenseBills field to given value.

### HasExpenseBills

`func (o *GLExport) HasExpenseBills() bool`

HasExpenseBills returns a boolean if a field has been set.

### GetPurchaseTransactions

`func (o *GLExport) GetPurchaseTransactions() []GLExportPurchaseTransaction`

GetPurchaseTransactions returns the PurchaseTransactions field if non-nil, zero value otherwise.

### GetPurchaseTransactionsOk

`func (o *GLExport) GetPurchaseTransactionsOk() (*[]GLExportPurchaseTransaction, bool)`

GetPurchaseTransactionsOk returns a tuple with the PurchaseTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTransactions

`func (o *GLExport) SetPurchaseTransactions(v []GLExportPurchaseTransaction)`

SetPurchaseTransactions sets PurchaseTransactions field to given value.

### HasPurchaseTransactions

`func (o *GLExport) HasPurchaseTransactions() bool`

HasPurchaseTransactions returns a boolean if a field has been set.

### GetAdjustmentTransactions

`func (o *GLExport) GetAdjustmentTransactions() []GLExportAdjustmentTransaction`

GetAdjustmentTransactions returns the AdjustmentTransactions field if non-nil, zero value otherwise.

### GetAdjustmentTransactionsOk

`func (o *GLExport) GetAdjustmentTransactionsOk() (*[]GLExportAdjustmentTransaction, bool)`

GetAdjustmentTransactionsOk returns a tuple with the AdjustmentTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTransactions

`func (o *GLExport) SetAdjustmentTransactions(v []GLExportAdjustmentTransaction)`

SetAdjustmentTransactions sets AdjustmentTransactions field to given value.

### HasAdjustmentTransactions

`func (o *GLExport) HasAdjustmentTransactions() bool`

HasAdjustmentTransactions returns a boolean if a field has been set.

### GetInventoryTransfers

`func (o *GLExport) GetInventoryTransfers() []GLExportInventoryTransfer`

GetInventoryTransfers returns the InventoryTransfers field if non-nil, zero value otherwise.

### GetInventoryTransfersOk

`func (o *GLExport) GetInventoryTransfersOk() (*[]GLExportInventoryTransfer, bool)`

GetInventoryTransfersOk returns a tuple with the InventoryTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryTransfers

`func (o *GLExport) SetInventoryTransfers(v []GLExportInventoryTransfer)`

SetInventoryTransfers sets InventoryTransfers field to given value.

### HasInventoryTransfers

`func (o *GLExport) HasInventoryTransfers() bool`

HasInventoryTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


