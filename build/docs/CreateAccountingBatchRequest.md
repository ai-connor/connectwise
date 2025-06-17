# CreateAccountingBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BatchIdentifier** | Pointer to **string** |  Max length: 50; | [optional] 
**GlInterfaceIdentifier** | Pointer to **string** |  | [optional] 
**ExportInvoicesFlag** | Pointer to **NullableBool** | Batch must export Invoices, Expenses or Products. | [optional] 
**ExportExpensesFlag** | Pointer to **NullableBool** | Batch must export Invoices, Expenses or Products. | [optional] 
**ExportProductsFlag** | Pointer to **NullableBool** | Batch must export Invoices, Expenses or Products. | [optional] 
**ProcessedRecordIds** | **[]int32** | GL Entry RecIDs. | 
**SummarizeExpenses** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCreateAccountingBatchRequest

`func NewCreateAccountingBatchRequest(processedRecordIds []*int32, ) *CreateAccountingBatchRequest`

NewCreateAccountingBatchRequest instantiates a new CreateAccountingBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountingBatchRequestWithDefaults

`func NewCreateAccountingBatchRequestWithDefaults() *CreateAccountingBatchRequest`

NewCreateAccountingBatchRequestWithDefaults instantiates a new CreateAccountingBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateAccountingBatchRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAccountingBatchRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAccountingBatchRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateAccountingBatchRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBatchIdentifier

`func (o *CreateAccountingBatchRequest) GetBatchIdentifier() string`

GetBatchIdentifier returns the BatchIdentifier field if non-nil, zero value otherwise.

### GetBatchIdentifierOk

`func (o *CreateAccountingBatchRequest) GetBatchIdentifierOk() (*string, bool)`

GetBatchIdentifierOk returns a tuple with the BatchIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchIdentifier

`func (o *CreateAccountingBatchRequest) SetBatchIdentifier(v string)`

SetBatchIdentifier sets BatchIdentifier field to given value.

### HasBatchIdentifier

`func (o *CreateAccountingBatchRequest) HasBatchIdentifier() bool`

HasBatchIdentifier returns a boolean if a field has been set.

### GetGlInterfaceIdentifier

`func (o *CreateAccountingBatchRequest) GetGlInterfaceIdentifier() string`

GetGlInterfaceIdentifier returns the GlInterfaceIdentifier field if non-nil, zero value otherwise.

### GetGlInterfaceIdentifierOk

`func (o *CreateAccountingBatchRequest) GetGlInterfaceIdentifierOk() (*string, bool)`

GetGlInterfaceIdentifierOk returns a tuple with the GlInterfaceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlInterfaceIdentifier

`func (o *CreateAccountingBatchRequest) SetGlInterfaceIdentifier(v string)`

SetGlInterfaceIdentifier sets GlInterfaceIdentifier field to given value.

### HasGlInterfaceIdentifier

`func (o *CreateAccountingBatchRequest) HasGlInterfaceIdentifier() bool`

HasGlInterfaceIdentifier returns a boolean if a field has been set.

### GetExportInvoicesFlag

`func (o *CreateAccountingBatchRequest) GetExportInvoicesFlag() bool`

GetExportInvoicesFlag returns the ExportInvoicesFlag field if non-nil, zero value otherwise.

### GetExportInvoicesFlagOk

`func (o *CreateAccountingBatchRequest) GetExportInvoicesFlagOk() (*bool, bool)`

GetExportInvoicesFlagOk returns a tuple with the ExportInvoicesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportInvoicesFlag

`func (o *CreateAccountingBatchRequest) SetExportInvoicesFlag(v bool)`

SetExportInvoicesFlag sets ExportInvoicesFlag field to given value.

### HasExportInvoicesFlag

`func (o *CreateAccountingBatchRequest) HasExportInvoicesFlag() bool`

HasExportInvoicesFlag returns a boolean if a field has been set.

### SetExportInvoicesFlagNil

`func (o *CreateAccountingBatchRequest) SetExportInvoicesFlagNil(b bool)`

 SetExportInvoicesFlagNil sets the value for ExportInvoicesFlag to be an explicit nil

### UnsetExportInvoicesFlag
`func (o *CreateAccountingBatchRequest) UnsetExportInvoicesFlag()`

UnsetExportInvoicesFlag ensures that no value is present for ExportInvoicesFlag, not even an explicit nil
### GetExportExpensesFlag

`func (o *CreateAccountingBatchRequest) GetExportExpensesFlag() bool`

GetExportExpensesFlag returns the ExportExpensesFlag field if non-nil, zero value otherwise.

### GetExportExpensesFlagOk

`func (o *CreateAccountingBatchRequest) GetExportExpensesFlagOk() (*bool, bool)`

GetExportExpensesFlagOk returns a tuple with the ExportExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportExpensesFlag

`func (o *CreateAccountingBatchRequest) SetExportExpensesFlag(v bool)`

SetExportExpensesFlag sets ExportExpensesFlag field to given value.

### HasExportExpensesFlag

`func (o *CreateAccountingBatchRequest) HasExportExpensesFlag() bool`

HasExportExpensesFlag returns a boolean if a field has been set.

### SetExportExpensesFlagNil

`func (o *CreateAccountingBatchRequest) SetExportExpensesFlagNil(b bool)`

 SetExportExpensesFlagNil sets the value for ExportExpensesFlag to be an explicit nil

### UnsetExportExpensesFlag
`func (o *CreateAccountingBatchRequest) UnsetExportExpensesFlag()`

UnsetExportExpensesFlag ensures that no value is present for ExportExpensesFlag, not even an explicit nil
### GetExportProductsFlag

`func (o *CreateAccountingBatchRequest) GetExportProductsFlag() bool`

GetExportProductsFlag returns the ExportProductsFlag field if non-nil, zero value otherwise.

### GetExportProductsFlagOk

`func (o *CreateAccountingBatchRequest) GetExportProductsFlagOk() (*bool, bool)`

GetExportProductsFlagOk returns a tuple with the ExportProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportProductsFlag

`func (o *CreateAccountingBatchRequest) SetExportProductsFlag(v bool)`

SetExportProductsFlag sets ExportProductsFlag field to given value.

### HasExportProductsFlag

`func (o *CreateAccountingBatchRequest) HasExportProductsFlag() bool`

HasExportProductsFlag returns a boolean if a field has been set.

### SetExportProductsFlagNil

`func (o *CreateAccountingBatchRequest) SetExportProductsFlagNil(b bool)`

 SetExportProductsFlagNil sets the value for ExportProductsFlag to be an explicit nil

### UnsetExportProductsFlag
`func (o *CreateAccountingBatchRequest) UnsetExportProductsFlag()`

UnsetExportProductsFlag ensures that no value is present for ExportProductsFlag, not even an explicit nil
### GetProcessedRecordIds

`func (o *CreateAccountingBatchRequest) GetProcessedRecordIds() []*int32`

GetProcessedRecordIds returns the ProcessedRecordIds field if non-nil, zero value otherwise.

### GetProcessedRecordIdsOk

`func (o *CreateAccountingBatchRequest) GetProcessedRecordIdsOk() (*[]*int32, bool)`

GetProcessedRecordIdsOk returns a tuple with the ProcessedRecordIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedRecordIds

`func (o *CreateAccountingBatchRequest) SetProcessedRecordIds(v []*int32)`

SetProcessedRecordIds sets ProcessedRecordIds field to given value.


### GetSummarizeExpenses

`func (o *CreateAccountingBatchRequest) GetSummarizeExpenses() bool`

GetSummarizeExpenses returns the SummarizeExpenses field if non-nil, zero value otherwise.

### GetSummarizeExpensesOk

`func (o *CreateAccountingBatchRequest) GetSummarizeExpensesOk() (*bool, bool)`

GetSummarizeExpensesOk returns a tuple with the SummarizeExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummarizeExpenses

`func (o *CreateAccountingBatchRequest) SetSummarizeExpenses(v bool)`

SetSummarizeExpenses sets SummarizeExpenses field to given value.

### HasSummarizeExpenses

`func (o *CreateAccountingBatchRequest) HasSummarizeExpenses() bool`

HasSummarizeExpenses returns a boolean if a field has been set.

### SetSummarizeExpensesNil

`func (o *CreateAccountingBatchRequest) SetSummarizeExpensesNil(b bool)`

 SetSummarizeExpensesNil sets the value for SummarizeExpenses to be an explicit nil

### UnsetSummarizeExpenses
`func (o *CreateAccountingBatchRequest) UnsetSummarizeExpenses()`

UnsetSummarizeExpenses ensures that no value is present for SummarizeExpenses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


