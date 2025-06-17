# ExportAccountingBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchIdentifier** | Pointer to **string** |  Max length: 50; | [optional] 
**GlInterfaceIdentifier** | Pointer to **string** |  | [optional] 
**ThruDate** | Pointer to **time.Time** |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**SummarizeInvoices** | Pointer to **string** |  | [optional] 
**ExportInvoicesFlag** | Pointer to **NullableBool** | Batch export must include invoices, expenses, or products (procurement). | [optional] 
**IncludedInvoiceIds** | Pointer to **[]int32** |  | [optional] 
**ExcludedInvoiceIds** | Pointer to **[]int32** |  | [optional] 
**ExportExpensesFlag** | Pointer to **NullableBool** | Batch export must include invoices, expenses, or products (procurement). | [optional] 
**IncludedExpenseIds** | Pointer to **[]int32** |  | [optional] 
**ExcludedExpenseIds** | Pointer to **[]int32** |  | [optional] 
**ExportPaymentsFlag** | Pointer to **NullableBool** | Batch export must include invoices, expenses, or products (procurement). | [optional] 
**IncludedPaymentIds** | Pointer to **[]int32** |  | [optional] 
**ExportProductsFlag** | Pointer to **NullableBool** | Batch export must include invoices, expenses, or products (procurement). | [optional] 
**IncludedProductIds** | Pointer to **[]string** |  | [optional] 
**ExcludedProductIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewExportAccountingBatchRequest

`func NewExportAccountingBatchRequest() *ExportAccountingBatchRequest`

NewExportAccountingBatchRequest instantiates a new ExportAccountingBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportAccountingBatchRequestWithDefaults

`func NewExportAccountingBatchRequestWithDefaults() *ExportAccountingBatchRequest`

NewExportAccountingBatchRequestWithDefaults instantiates a new ExportAccountingBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchIdentifier

`func (o *ExportAccountingBatchRequest) GetBatchIdentifier() string`

GetBatchIdentifier returns the BatchIdentifier field if non-nil, zero value otherwise.

### GetBatchIdentifierOk

`func (o *ExportAccountingBatchRequest) GetBatchIdentifierOk() (*string, bool)`

GetBatchIdentifierOk returns a tuple with the BatchIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchIdentifier

`func (o *ExportAccountingBatchRequest) SetBatchIdentifier(v string)`

SetBatchIdentifier sets BatchIdentifier field to given value.

### HasBatchIdentifier

`func (o *ExportAccountingBatchRequest) HasBatchIdentifier() bool`

HasBatchIdentifier returns a boolean if a field has been set.

### GetGlInterfaceIdentifier

`func (o *ExportAccountingBatchRequest) GetGlInterfaceIdentifier() string`

GetGlInterfaceIdentifier returns the GlInterfaceIdentifier field if non-nil, zero value otherwise.

### GetGlInterfaceIdentifierOk

`func (o *ExportAccountingBatchRequest) GetGlInterfaceIdentifierOk() (*string, bool)`

GetGlInterfaceIdentifierOk returns a tuple with the GlInterfaceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlInterfaceIdentifier

`func (o *ExportAccountingBatchRequest) SetGlInterfaceIdentifier(v string)`

SetGlInterfaceIdentifier sets GlInterfaceIdentifier field to given value.

### HasGlInterfaceIdentifier

`func (o *ExportAccountingBatchRequest) HasGlInterfaceIdentifier() bool`

HasGlInterfaceIdentifier returns a boolean if a field has been set.

### GetThruDate

`func (o *ExportAccountingBatchRequest) GetThruDate() time.Time`

GetThruDate returns the ThruDate field if non-nil, zero value otherwise.

### GetThruDateOk

`func (o *ExportAccountingBatchRequest) GetThruDateOk() (*time.Time, bool)`

GetThruDateOk returns a tuple with the ThruDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThruDate

`func (o *ExportAccountingBatchRequest) SetThruDate(v time.Time)`

SetThruDate sets ThruDate field to given value.

### HasThruDate

`func (o *ExportAccountingBatchRequest) HasThruDate() bool`

HasThruDate returns a boolean if a field has been set.

### GetLocationId

`func (o *ExportAccountingBatchRequest) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ExportAccountingBatchRequest) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ExportAccountingBatchRequest) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *ExportAccountingBatchRequest) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *ExportAccountingBatchRequest) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *ExportAccountingBatchRequest) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetSummarizeInvoices

`func (o *ExportAccountingBatchRequest) GetSummarizeInvoices() string`

GetSummarizeInvoices returns the SummarizeInvoices field if non-nil, zero value otherwise.

### GetSummarizeInvoicesOk

`func (o *ExportAccountingBatchRequest) GetSummarizeInvoicesOk() (*string, bool)`

GetSummarizeInvoicesOk returns a tuple with the SummarizeInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummarizeInvoices

`func (o *ExportAccountingBatchRequest) SetSummarizeInvoices(v string)`

SetSummarizeInvoices sets SummarizeInvoices field to given value.

### HasSummarizeInvoices

`func (o *ExportAccountingBatchRequest) HasSummarizeInvoices() bool`

HasSummarizeInvoices returns a boolean if a field has been set.

### GetExportInvoicesFlag

`func (o *ExportAccountingBatchRequest) GetExportInvoicesFlag() bool`

GetExportInvoicesFlag returns the ExportInvoicesFlag field if non-nil, zero value otherwise.

### GetExportInvoicesFlagOk

`func (o *ExportAccountingBatchRequest) GetExportInvoicesFlagOk() (*bool, bool)`

GetExportInvoicesFlagOk returns a tuple with the ExportInvoicesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportInvoicesFlag

`func (o *ExportAccountingBatchRequest) SetExportInvoicesFlag(v bool)`

SetExportInvoicesFlag sets ExportInvoicesFlag field to given value.

### HasExportInvoicesFlag

`func (o *ExportAccountingBatchRequest) HasExportInvoicesFlag() bool`

HasExportInvoicesFlag returns a boolean if a field has been set.

### SetExportInvoicesFlagNil

`func (o *ExportAccountingBatchRequest) SetExportInvoicesFlagNil(b bool)`

 SetExportInvoicesFlagNil sets the value for ExportInvoicesFlag to be an explicit nil

### UnsetExportInvoicesFlag
`func (o *ExportAccountingBatchRequest) UnsetExportInvoicesFlag()`

UnsetExportInvoicesFlag ensures that no value is present for ExportInvoicesFlag, not even an explicit nil
### GetIncludedInvoiceIds

`func (o *ExportAccountingBatchRequest) GetIncludedInvoiceIds() []*int32`

GetIncludedInvoiceIds returns the IncludedInvoiceIds field if non-nil, zero value otherwise.

### GetIncludedInvoiceIdsOk

`func (o *ExportAccountingBatchRequest) GetIncludedInvoiceIdsOk() (*[]*int32, bool)`

GetIncludedInvoiceIdsOk returns a tuple with the IncludedInvoiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedInvoiceIds

`func (o *ExportAccountingBatchRequest) SetIncludedInvoiceIds(v []*int32)`

SetIncludedInvoiceIds sets IncludedInvoiceIds field to given value.

### HasIncludedInvoiceIds

`func (o *ExportAccountingBatchRequest) HasIncludedInvoiceIds() bool`

HasIncludedInvoiceIds returns a boolean if a field has been set.

### GetExcludedInvoiceIds

`func (o *ExportAccountingBatchRequest) GetExcludedInvoiceIds() []*int32`

GetExcludedInvoiceIds returns the ExcludedInvoiceIds field if non-nil, zero value otherwise.

### GetExcludedInvoiceIdsOk

`func (o *ExportAccountingBatchRequest) GetExcludedInvoiceIdsOk() (*[]*int32, bool)`

GetExcludedInvoiceIdsOk returns a tuple with the ExcludedInvoiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedInvoiceIds

`func (o *ExportAccountingBatchRequest) SetExcludedInvoiceIds(v []*int32)`

SetExcludedInvoiceIds sets ExcludedInvoiceIds field to given value.

### HasExcludedInvoiceIds

`func (o *ExportAccountingBatchRequest) HasExcludedInvoiceIds() bool`

HasExcludedInvoiceIds returns a boolean if a field has been set.

### GetExportExpensesFlag

`func (o *ExportAccountingBatchRequest) GetExportExpensesFlag() bool`

GetExportExpensesFlag returns the ExportExpensesFlag field if non-nil, zero value otherwise.

### GetExportExpensesFlagOk

`func (o *ExportAccountingBatchRequest) GetExportExpensesFlagOk() (*bool, bool)`

GetExportExpensesFlagOk returns a tuple with the ExportExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportExpensesFlag

`func (o *ExportAccountingBatchRequest) SetExportExpensesFlag(v bool)`

SetExportExpensesFlag sets ExportExpensesFlag field to given value.

### HasExportExpensesFlag

`func (o *ExportAccountingBatchRequest) HasExportExpensesFlag() bool`

HasExportExpensesFlag returns a boolean if a field has been set.

### SetExportExpensesFlagNil

`func (o *ExportAccountingBatchRequest) SetExportExpensesFlagNil(b bool)`

 SetExportExpensesFlagNil sets the value for ExportExpensesFlag to be an explicit nil

### UnsetExportExpensesFlag
`func (o *ExportAccountingBatchRequest) UnsetExportExpensesFlag()`

UnsetExportExpensesFlag ensures that no value is present for ExportExpensesFlag, not even an explicit nil
### GetIncludedExpenseIds

`func (o *ExportAccountingBatchRequest) GetIncludedExpenseIds() []*int32`

GetIncludedExpenseIds returns the IncludedExpenseIds field if non-nil, zero value otherwise.

### GetIncludedExpenseIdsOk

`func (o *ExportAccountingBatchRequest) GetIncludedExpenseIdsOk() (*[]*int32, bool)`

GetIncludedExpenseIdsOk returns a tuple with the IncludedExpenseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedExpenseIds

`func (o *ExportAccountingBatchRequest) SetIncludedExpenseIds(v []*int32)`

SetIncludedExpenseIds sets IncludedExpenseIds field to given value.

### HasIncludedExpenseIds

`func (o *ExportAccountingBatchRequest) HasIncludedExpenseIds() bool`

HasIncludedExpenseIds returns a boolean if a field has been set.

### GetExcludedExpenseIds

`func (o *ExportAccountingBatchRequest) GetExcludedExpenseIds() []*int32`

GetExcludedExpenseIds returns the ExcludedExpenseIds field if non-nil, zero value otherwise.

### GetExcludedExpenseIdsOk

`func (o *ExportAccountingBatchRequest) GetExcludedExpenseIdsOk() (*[]*int32, bool)`

GetExcludedExpenseIdsOk returns a tuple with the ExcludedExpenseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedExpenseIds

`func (o *ExportAccountingBatchRequest) SetExcludedExpenseIds(v []*int32)`

SetExcludedExpenseIds sets ExcludedExpenseIds field to given value.

### HasExcludedExpenseIds

`func (o *ExportAccountingBatchRequest) HasExcludedExpenseIds() bool`

HasExcludedExpenseIds returns a boolean if a field has been set.

### GetExportPaymentsFlag

`func (o *ExportAccountingBatchRequest) GetExportPaymentsFlag() bool`

GetExportPaymentsFlag returns the ExportPaymentsFlag field if non-nil, zero value otherwise.

### GetExportPaymentsFlagOk

`func (o *ExportAccountingBatchRequest) GetExportPaymentsFlagOk() (*bool, bool)`

GetExportPaymentsFlagOk returns a tuple with the ExportPaymentsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPaymentsFlag

`func (o *ExportAccountingBatchRequest) SetExportPaymentsFlag(v bool)`

SetExportPaymentsFlag sets ExportPaymentsFlag field to given value.

### HasExportPaymentsFlag

`func (o *ExportAccountingBatchRequest) HasExportPaymentsFlag() bool`

HasExportPaymentsFlag returns a boolean if a field has been set.

### SetExportPaymentsFlagNil

`func (o *ExportAccountingBatchRequest) SetExportPaymentsFlagNil(b bool)`

 SetExportPaymentsFlagNil sets the value for ExportPaymentsFlag to be an explicit nil

### UnsetExportPaymentsFlag
`func (o *ExportAccountingBatchRequest) UnsetExportPaymentsFlag()`

UnsetExportPaymentsFlag ensures that no value is present for ExportPaymentsFlag, not even an explicit nil
### GetIncludedPaymentIds

`func (o *ExportAccountingBatchRequest) GetIncludedPaymentIds() []int32`

GetIncludedPaymentIds returns the IncludedPaymentIds field if non-nil, zero value otherwise.

### GetIncludedPaymentIdsOk

`func (o *ExportAccountingBatchRequest) GetIncludedPaymentIdsOk() (*[]int32, bool)`

GetIncludedPaymentIdsOk returns a tuple with the IncludedPaymentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedPaymentIds

`func (o *ExportAccountingBatchRequest) SetIncludedPaymentIds(v []int32)`

SetIncludedPaymentIds sets IncludedPaymentIds field to given value.

### HasIncludedPaymentIds

`func (o *ExportAccountingBatchRequest) HasIncludedPaymentIds() bool`

HasIncludedPaymentIds returns a boolean if a field has been set.

### GetExportProductsFlag

`func (o *ExportAccountingBatchRequest) GetExportProductsFlag() bool`

GetExportProductsFlag returns the ExportProductsFlag field if non-nil, zero value otherwise.

### GetExportProductsFlagOk

`func (o *ExportAccountingBatchRequest) GetExportProductsFlagOk() (*bool, bool)`

GetExportProductsFlagOk returns a tuple with the ExportProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportProductsFlag

`func (o *ExportAccountingBatchRequest) SetExportProductsFlag(v bool)`

SetExportProductsFlag sets ExportProductsFlag field to given value.

### HasExportProductsFlag

`func (o *ExportAccountingBatchRequest) HasExportProductsFlag() bool`

HasExportProductsFlag returns a boolean if a field has been set.

### SetExportProductsFlagNil

`func (o *ExportAccountingBatchRequest) SetExportProductsFlagNil(b bool)`

 SetExportProductsFlagNil sets the value for ExportProductsFlag to be an explicit nil

### UnsetExportProductsFlag
`func (o *ExportAccountingBatchRequest) UnsetExportProductsFlag()`

UnsetExportProductsFlag ensures that no value is present for ExportProductsFlag, not even an explicit nil
### GetIncludedProductIds

`func (o *ExportAccountingBatchRequest) GetIncludedProductIds() []string`

GetIncludedProductIds returns the IncludedProductIds field if non-nil, zero value otherwise.

### GetIncludedProductIdsOk

`func (o *ExportAccountingBatchRequest) GetIncludedProductIdsOk() (*[]string, bool)`

GetIncludedProductIdsOk returns a tuple with the IncludedProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedProductIds

`func (o *ExportAccountingBatchRequest) SetIncludedProductIds(v []string)`

SetIncludedProductIds sets IncludedProductIds field to given value.

### HasIncludedProductIds

`func (o *ExportAccountingBatchRequest) HasIncludedProductIds() bool`

HasIncludedProductIds returns a boolean if a field has been set.

### GetExcludedProductIds

`func (o *ExportAccountingBatchRequest) GetExcludedProductIds() []string`

GetExcludedProductIds returns the ExcludedProductIds field if non-nil, zero value otherwise.

### GetExcludedProductIdsOk

`func (o *ExportAccountingBatchRequest) GetExcludedProductIdsOk() (*[]string, bool)`

GetExcludedProductIdsOk returns a tuple with the ExcludedProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProductIds

`func (o *ExportAccountingBatchRequest) SetExcludedProductIds(v []string)`

SetExcludedProductIds sets ExcludedProductIds field to given value.

### HasExcludedProductIds

`func (o *ExportAccountingBatchRequest) HasExcludedProductIds() bool`

HasExcludedProductIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


