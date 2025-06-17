# AccountingBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BatchIdentifier** | Pointer to **string** |  | [optional] 
**ExportInvoicesFlag** | Pointer to **NullableBool** |  | [optional] 
**ExportExpensesFlag** | Pointer to **NullableBool** |  | [optional] 
**ExportProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAccountingBatch

`func NewAccountingBatch() *AccountingBatch`

NewAccountingBatch instantiates a new AccountingBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingBatchWithDefaults

`func NewAccountingBatchWithDefaults() *AccountingBatch`

NewAccountingBatchWithDefaults instantiates a new AccountingBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountingBatch) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountingBatch) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountingBatch) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccountingBatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBatchIdentifier

`func (o *AccountingBatch) GetBatchIdentifier() string`

GetBatchIdentifier returns the BatchIdentifier field if non-nil, zero value otherwise.

### GetBatchIdentifierOk

`func (o *AccountingBatch) GetBatchIdentifierOk() (*string, bool)`

GetBatchIdentifierOk returns a tuple with the BatchIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchIdentifier

`func (o *AccountingBatch) SetBatchIdentifier(v string)`

SetBatchIdentifier sets BatchIdentifier field to given value.

### HasBatchIdentifier

`func (o *AccountingBatch) HasBatchIdentifier() bool`

HasBatchIdentifier returns a boolean if a field has been set.

### GetExportInvoicesFlag

`func (o *AccountingBatch) GetExportInvoicesFlag() bool`

GetExportInvoicesFlag returns the ExportInvoicesFlag field if non-nil, zero value otherwise.

### GetExportInvoicesFlagOk

`func (o *AccountingBatch) GetExportInvoicesFlagOk() (*bool, bool)`

GetExportInvoicesFlagOk returns a tuple with the ExportInvoicesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportInvoicesFlag

`func (o *AccountingBatch) SetExportInvoicesFlag(v bool)`

SetExportInvoicesFlag sets ExportInvoicesFlag field to given value.

### HasExportInvoicesFlag

`func (o *AccountingBatch) HasExportInvoicesFlag() bool`

HasExportInvoicesFlag returns a boolean if a field has been set.

### SetExportInvoicesFlagNil

`func (o *AccountingBatch) SetExportInvoicesFlagNil(b bool)`

 SetExportInvoicesFlagNil sets the value for ExportInvoicesFlag to be an explicit nil

### UnsetExportInvoicesFlag
`func (o *AccountingBatch) UnsetExportInvoicesFlag()`

UnsetExportInvoicesFlag ensures that no value is present for ExportInvoicesFlag, not even an explicit nil
### GetExportExpensesFlag

`func (o *AccountingBatch) GetExportExpensesFlag() bool`

GetExportExpensesFlag returns the ExportExpensesFlag field if non-nil, zero value otherwise.

### GetExportExpensesFlagOk

`func (o *AccountingBatch) GetExportExpensesFlagOk() (*bool, bool)`

GetExportExpensesFlagOk returns a tuple with the ExportExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportExpensesFlag

`func (o *AccountingBatch) SetExportExpensesFlag(v bool)`

SetExportExpensesFlag sets ExportExpensesFlag field to given value.

### HasExportExpensesFlag

`func (o *AccountingBatch) HasExportExpensesFlag() bool`

HasExportExpensesFlag returns a boolean if a field has been set.

### SetExportExpensesFlagNil

`func (o *AccountingBatch) SetExportExpensesFlagNil(b bool)`

 SetExportExpensesFlagNil sets the value for ExportExpensesFlag to be an explicit nil

### UnsetExportExpensesFlag
`func (o *AccountingBatch) UnsetExportExpensesFlag()`

UnsetExportExpensesFlag ensures that no value is present for ExportExpensesFlag, not even an explicit nil
### GetExportProductsFlag

`func (o *AccountingBatch) GetExportProductsFlag() bool`

GetExportProductsFlag returns the ExportProductsFlag field if non-nil, zero value otherwise.

### GetExportProductsFlagOk

`func (o *AccountingBatch) GetExportProductsFlagOk() (*bool, bool)`

GetExportProductsFlagOk returns a tuple with the ExportProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportProductsFlag

`func (o *AccountingBatch) SetExportProductsFlag(v bool)`

SetExportProductsFlag sets ExportProductsFlag field to given value.

### HasExportProductsFlag

`func (o *AccountingBatch) HasExportProductsFlag() bool`

HasExportProductsFlag returns a boolean if a field has been set.

### SetExportProductsFlagNil

`func (o *AccountingBatch) SetExportProductsFlagNil(b bool)`

 SetExportProductsFlagNil sets the value for ExportProductsFlag to be an explicit nil

### UnsetExportProductsFlag
`func (o *AccountingBatch) UnsetExportProductsFlag()`

UnsetExportProductsFlag ensures that no value is present for ExportProductsFlag, not even an explicit nil
### GetClosedFlag

`func (o *AccountingBatch) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *AccountingBatch) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *AccountingBatch) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *AccountingBatch) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *AccountingBatch) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *AccountingBatch) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetInfo

`func (o *AccountingBatch) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AccountingBatch) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AccountingBatch) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AccountingBatch) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


