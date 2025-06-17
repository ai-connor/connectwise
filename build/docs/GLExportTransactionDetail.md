# GLExportTransactionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**DocumentDate** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**GlClass** | Pointer to **string** |  | [optional] 
**GlTypeId** | Pointer to **string** |  | [optional] 
**GlItemId** | Pointer to **string** |  | [optional] 
**InvoiceSummaryOption** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**SalesCode** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **NullableFloat64** |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**TimeEntry** | Pointer to [**TimeEntryReference**](TimeEntryReference.md) |  | [optional] 
**CostAccountNumber** | Pointer to **string** |  | [optional] 
**InventoryAccountNumber** | Pointer to **string** |  | [optional] 
**ProductAccountNumber** | Pointer to **string** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**TaxCodeXref** | Pointer to **string** |  | [optional] 
**TaxAgencyXref** | Pointer to **string** |  | [optional] 
**TaxNote** | Pointer to **string** |  | [optional] 
**TaxRate** | Pointer to **NullableFloat64** |  | [optional] 
**TaxRatePercent** | Pointer to **NullableFloat64** |  | [optional] 
**TaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**Taxable2Flag** | Pointer to **NullableBool** |  | [optional] 
**Taxable3Flag** | Pointer to **NullableBool** |  | [optional] 
**Taxable4Flag** | Pointer to **NullableBool** |  | [optional] 
**Taxable5Flag** | Pointer to **NullableBool** |  | [optional] 
**Item** | Pointer to [**IvItemReference**](IvItemReference.md) |  | [optional] 
**Product** | Pointer to [**ProductReference**](ProductReference.md) |  | [optional] 
**ItemTaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**ItemPrice** | Pointer to **NullableFloat64** |  | [optional] 
**ItemCost** | Pointer to **NullableFloat64** |  | [optional] 
**ItemDescription** | Pointer to **string** |  | [optional] 
**SalesDescription** | Pointer to **string** |  | [optional] 
**UnitOfMeasure** | Pointer to [**UnitOfMeasureReference**](UnitOfMeasureReference.md) |  | [optional] 
**SubCategory** | Pointer to [**ProductSubCategoryReference**](ProductSubCategoryReference.md) |  | [optional] 
**SerializedFlag** | Pointer to **NullableBool** |  | [optional] 
**SerialNumbers** | Pointer to **string** |  | [optional] 
**WarehouseSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**WarehouseBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**ShipmentMethod** | Pointer to [**ShipmentMethodReference**](ShipmentMethodReference.md) |  | [optional] 
**DropShippedFlag** | Pointer to **NullableBool** |  | [optional] 
**ItemTypeXref** | Pointer to **string** |  | [optional] 
**InventoryXref** | Pointer to **string** |  | [optional] 
**CogsXref** | Pointer to **string** |  | [optional] 
**UomScheduleXref** | Pointer to **string** |  | [optional] 
**PriceLevelXref** | Pointer to **string** |  | [optional] 
**LocationXref** | Pointer to **string** |  | [optional] 
**TaxLevels** | Pointer to [**[]GLExportTransactionDetailTaxLevel**](GLExportTransactionDetailTaxLevel.md) |  | [optional] 

## Methods

### NewGLExportTransactionDetail

`func NewGLExportTransactionDetail() *GLExportTransactionDetail`

NewGLExportTransactionDetail instantiates a new GLExportTransactionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportTransactionDetailWithDefaults

`func NewGLExportTransactionDetailWithDefaults() *GLExportTransactionDetail`

NewGLExportTransactionDetailWithDefaults instantiates a new GLExportTransactionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLExportTransactionDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLExportTransactionDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLExportTransactionDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GLExportTransactionDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GLExportTransactionDetail) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GLExportTransactionDetail) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDocumentDate

`func (o *GLExportTransactionDetail) GetDocumentDate() string`

GetDocumentDate returns the DocumentDate field if non-nil, zero value otherwise.

### GetDocumentDateOk

`func (o *GLExportTransactionDetail) GetDocumentDateOk() (*string, bool)`

GetDocumentDateOk returns a tuple with the DocumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDate

`func (o *GLExportTransactionDetail) SetDocumentDate(v string)`

SetDocumentDate sets DocumentDate field to given value.

### HasDocumentDate

`func (o *GLExportTransactionDetail) HasDocumentDate() bool`

HasDocumentDate returns a boolean if a field has been set.

### GetDocumentType

`func (o *GLExportTransactionDetail) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *GLExportTransactionDetail) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *GLExportTransactionDetail) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *GLExportTransactionDetail) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetAccountNumber

`func (o *GLExportTransactionDetail) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GLExportTransactionDetail) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GLExportTransactionDetail) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GLExportTransactionDetail) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetGlClass

`func (o *GLExportTransactionDetail) GetGlClass() string`

GetGlClass returns the GlClass field if non-nil, zero value otherwise.

### GetGlClassOk

`func (o *GLExportTransactionDetail) GetGlClassOk() (*string, bool)`

GetGlClassOk returns a tuple with the GlClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlClass

`func (o *GLExportTransactionDetail) SetGlClass(v string)`

SetGlClass sets GlClass field to given value.

### HasGlClass

`func (o *GLExportTransactionDetail) HasGlClass() bool`

HasGlClass returns a boolean if a field has been set.

### GetGlTypeId

`func (o *GLExportTransactionDetail) GetGlTypeId() string`

GetGlTypeId returns the GlTypeId field if non-nil, zero value otherwise.

### GetGlTypeIdOk

`func (o *GLExportTransactionDetail) GetGlTypeIdOk() (*string, bool)`

GetGlTypeIdOk returns a tuple with the GlTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlTypeId

`func (o *GLExportTransactionDetail) SetGlTypeId(v string)`

SetGlTypeId sets GlTypeId field to given value.

### HasGlTypeId

`func (o *GLExportTransactionDetail) HasGlTypeId() bool`

HasGlTypeId returns a boolean if a field has been set.

### GetGlItemId

`func (o *GLExportTransactionDetail) GetGlItemId() string`

GetGlItemId returns the GlItemId field if non-nil, zero value otherwise.

### GetGlItemIdOk

`func (o *GLExportTransactionDetail) GetGlItemIdOk() (*string, bool)`

GetGlItemIdOk returns a tuple with the GlItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlItemId

`func (o *GLExportTransactionDetail) SetGlItemId(v string)`

SetGlItemId sets GlItemId field to given value.

### HasGlItemId

`func (o *GLExportTransactionDetail) HasGlItemId() bool`

HasGlItemId returns a boolean if a field has been set.

### GetInvoiceSummaryOption

`func (o *GLExportTransactionDetail) GetInvoiceSummaryOption() string`

GetInvoiceSummaryOption returns the InvoiceSummaryOption field if non-nil, zero value otherwise.

### GetInvoiceSummaryOptionOk

`func (o *GLExportTransactionDetail) GetInvoiceSummaryOptionOk() (*string, bool)`

GetInvoiceSummaryOptionOk returns a tuple with the InvoiceSummaryOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceSummaryOption

`func (o *GLExportTransactionDetail) SetInvoiceSummaryOption(v string)`

SetInvoiceSummaryOption sets InvoiceSummaryOption field to given value.

### HasInvoiceSummaryOption

`func (o *GLExportTransactionDetail) HasInvoiceSummaryOption() bool`

HasInvoiceSummaryOption returns a boolean if a field has been set.

### GetCost

`func (o *GLExportTransactionDetail) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GLExportTransactionDetail) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GLExportTransactionDetail) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GLExportTransactionDetail) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *GLExportTransactionDetail) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *GLExportTransactionDetail) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetSalesCode

`func (o *GLExportTransactionDetail) GetSalesCode() string`

GetSalesCode returns the SalesCode field if non-nil, zero value otherwise.

### GetSalesCodeOk

`func (o *GLExportTransactionDetail) GetSalesCodeOk() (*string, bool)`

GetSalesCodeOk returns a tuple with the SalesCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesCode

`func (o *GLExportTransactionDetail) SetSalesCode(v string)`

SetSalesCode sets SalesCode field to given value.

### HasSalesCode

`func (o *GLExportTransactionDetail) HasSalesCode() bool`

HasSalesCode returns a boolean if a field has been set.

### GetMemo

`func (o *GLExportTransactionDetail) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GLExportTransactionDetail) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GLExportTransactionDetail) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GLExportTransactionDetail) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetDescription

`func (o *GLExportTransactionDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GLExportTransactionDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GLExportTransactionDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GLExportTransactionDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *GLExportTransactionDetail) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GLExportTransactionDetail) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GLExportTransactionDetail) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GLExportTransactionDetail) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *GLExportTransactionDetail) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *GLExportTransactionDetail) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetTotal

`func (o *GLExportTransactionDetail) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GLExportTransactionDetail) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GLExportTransactionDetail) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GLExportTransactionDetail) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GLExportTransactionDetail) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GLExportTransactionDetail) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetCurrency

`func (o *GLExportTransactionDetail) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GLExportTransactionDetail) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GLExportTransactionDetail) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GLExportTransactionDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTimeEntry

`func (o *GLExportTransactionDetail) GetTimeEntry() TimeEntryReference`

GetTimeEntry returns the TimeEntry field if non-nil, zero value otherwise.

### GetTimeEntryOk

`func (o *GLExportTransactionDetail) GetTimeEntryOk() (*TimeEntryReference, bool)`

GetTimeEntryOk returns a tuple with the TimeEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntry

`func (o *GLExportTransactionDetail) SetTimeEntry(v TimeEntryReference)`

SetTimeEntry sets TimeEntry field to given value.

### HasTimeEntry

`func (o *GLExportTransactionDetail) HasTimeEntry() bool`

HasTimeEntry returns a boolean if a field has been set.

### GetCostAccountNumber

`func (o *GLExportTransactionDetail) GetCostAccountNumber() string`

GetCostAccountNumber returns the CostAccountNumber field if non-nil, zero value otherwise.

### GetCostAccountNumberOk

`func (o *GLExportTransactionDetail) GetCostAccountNumberOk() (*string, bool)`

GetCostAccountNumberOk returns a tuple with the CostAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAccountNumber

`func (o *GLExportTransactionDetail) SetCostAccountNumber(v string)`

SetCostAccountNumber sets CostAccountNumber field to given value.

### HasCostAccountNumber

`func (o *GLExportTransactionDetail) HasCostAccountNumber() bool`

HasCostAccountNumber returns a boolean if a field has been set.

### GetInventoryAccountNumber

`func (o *GLExportTransactionDetail) GetInventoryAccountNumber() string`

GetInventoryAccountNumber returns the InventoryAccountNumber field if non-nil, zero value otherwise.

### GetInventoryAccountNumberOk

`func (o *GLExportTransactionDetail) GetInventoryAccountNumberOk() (*string, bool)`

GetInventoryAccountNumberOk returns a tuple with the InventoryAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryAccountNumber

`func (o *GLExportTransactionDetail) SetInventoryAccountNumber(v string)`

SetInventoryAccountNumber sets InventoryAccountNumber field to given value.

### HasInventoryAccountNumber

`func (o *GLExportTransactionDetail) HasInventoryAccountNumber() bool`

HasInventoryAccountNumber returns a boolean if a field has been set.

### GetProductAccountNumber

`func (o *GLExportTransactionDetail) GetProductAccountNumber() string`

GetProductAccountNumber returns the ProductAccountNumber field if non-nil, zero value otherwise.

### GetProductAccountNumberOk

`func (o *GLExportTransactionDetail) GetProductAccountNumberOk() (*string, bool)`

GetProductAccountNumberOk returns a tuple with the ProductAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAccountNumber

`func (o *GLExportTransactionDetail) SetProductAccountNumber(v string)`

SetProductAccountNumber sets ProductAccountNumber field to given value.

### HasProductAccountNumber

`func (o *GLExportTransactionDetail) HasProductAccountNumber() bool`

HasProductAccountNumber returns a boolean if a field has been set.

### GetTaxCode

`func (o *GLExportTransactionDetail) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *GLExportTransactionDetail) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *GLExportTransactionDetail) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *GLExportTransactionDetail) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxCodeXref

`func (o *GLExportTransactionDetail) GetTaxCodeXref() string`

GetTaxCodeXref returns the TaxCodeXref field if non-nil, zero value otherwise.

### GetTaxCodeXrefOk

`func (o *GLExportTransactionDetail) GetTaxCodeXrefOk() (*string, bool)`

GetTaxCodeXrefOk returns a tuple with the TaxCodeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodeXref

`func (o *GLExportTransactionDetail) SetTaxCodeXref(v string)`

SetTaxCodeXref sets TaxCodeXref field to given value.

### HasTaxCodeXref

`func (o *GLExportTransactionDetail) HasTaxCodeXref() bool`

HasTaxCodeXref returns a boolean if a field has been set.

### GetTaxAgencyXref

`func (o *GLExportTransactionDetail) GetTaxAgencyXref() string`

GetTaxAgencyXref returns the TaxAgencyXref field if non-nil, zero value otherwise.

### GetTaxAgencyXrefOk

`func (o *GLExportTransactionDetail) GetTaxAgencyXrefOk() (*string, bool)`

GetTaxAgencyXrefOk returns a tuple with the TaxAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAgencyXref

`func (o *GLExportTransactionDetail) SetTaxAgencyXref(v string)`

SetTaxAgencyXref sets TaxAgencyXref field to given value.

### HasTaxAgencyXref

`func (o *GLExportTransactionDetail) HasTaxAgencyXref() bool`

HasTaxAgencyXref returns a boolean if a field has been set.

### GetTaxNote

`func (o *GLExportTransactionDetail) GetTaxNote() string`

GetTaxNote returns the TaxNote field if non-nil, zero value otherwise.

### GetTaxNoteOk

`func (o *GLExportTransactionDetail) GetTaxNoteOk() (*string, bool)`

GetTaxNoteOk returns a tuple with the TaxNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNote

`func (o *GLExportTransactionDetail) SetTaxNote(v string)`

SetTaxNote sets TaxNote field to given value.

### HasTaxNote

`func (o *GLExportTransactionDetail) HasTaxNote() bool`

HasTaxNote returns a boolean if a field has been set.

### GetTaxRate

`func (o *GLExportTransactionDetail) GetTaxRate() float64`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *GLExportTransactionDetail) GetTaxRateOk() (*float64, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *GLExportTransactionDetail) SetTaxRate(v float64)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *GLExportTransactionDetail) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *GLExportTransactionDetail) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *GLExportTransactionDetail) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetTaxRatePercent

`func (o *GLExportTransactionDetail) GetTaxRatePercent() float64`

GetTaxRatePercent returns the TaxRatePercent field if non-nil, zero value otherwise.

### GetTaxRatePercentOk

`func (o *GLExportTransactionDetail) GetTaxRatePercentOk() (*float64, bool)`

GetTaxRatePercentOk returns a tuple with the TaxRatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRatePercent

`func (o *GLExportTransactionDetail) SetTaxRatePercent(v float64)`

SetTaxRatePercent sets TaxRatePercent field to given value.

### HasTaxRatePercent

`func (o *GLExportTransactionDetail) HasTaxRatePercent() bool`

HasTaxRatePercent returns a boolean if a field has been set.

### SetTaxRatePercentNil

`func (o *GLExportTransactionDetail) SetTaxRatePercentNil(b bool)`

 SetTaxRatePercentNil sets the value for TaxRatePercent to be an explicit nil

### UnsetTaxRatePercent
`func (o *GLExportTransactionDetail) UnsetTaxRatePercent()`

UnsetTaxRatePercent ensures that no value is present for TaxRatePercent, not even an explicit nil
### GetTaxableFlag

`func (o *GLExportTransactionDetail) GetTaxableFlag() bool`

GetTaxableFlag returns the TaxableFlag field if non-nil, zero value otherwise.

### GetTaxableFlagOk

`func (o *GLExportTransactionDetail) GetTaxableFlagOk() (*bool, bool)`

GetTaxableFlagOk returns a tuple with the TaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableFlag

`func (o *GLExportTransactionDetail) SetTaxableFlag(v bool)`

SetTaxableFlag sets TaxableFlag field to given value.

### HasTaxableFlag

`func (o *GLExportTransactionDetail) HasTaxableFlag() bool`

HasTaxableFlag returns a boolean if a field has been set.

### SetTaxableFlagNil

`func (o *GLExportTransactionDetail) SetTaxableFlagNil(b bool)`

 SetTaxableFlagNil sets the value for TaxableFlag to be an explicit nil

### UnsetTaxableFlag
`func (o *GLExportTransactionDetail) UnsetTaxableFlag()`

UnsetTaxableFlag ensures that no value is present for TaxableFlag, not even an explicit nil
### GetTaxable2Flag

`func (o *GLExportTransactionDetail) GetTaxable2Flag() bool`

GetTaxable2Flag returns the Taxable2Flag field if non-nil, zero value otherwise.

### GetTaxable2FlagOk

`func (o *GLExportTransactionDetail) GetTaxable2FlagOk() (*bool, bool)`

GetTaxable2FlagOk returns a tuple with the Taxable2Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable2Flag

`func (o *GLExportTransactionDetail) SetTaxable2Flag(v bool)`

SetTaxable2Flag sets Taxable2Flag field to given value.

### HasTaxable2Flag

`func (o *GLExportTransactionDetail) HasTaxable2Flag() bool`

HasTaxable2Flag returns a boolean if a field has been set.

### SetTaxable2FlagNil

`func (o *GLExportTransactionDetail) SetTaxable2FlagNil(b bool)`

 SetTaxable2FlagNil sets the value for Taxable2Flag to be an explicit nil

### UnsetTaxable2Flag
`func (o *GLExportTransactionDetail) UnsetTaxable2Flag()`

UnsetTaxable2Flag ensures that no value is present for Taxable2Flag, not even an explicit nil
### GetTaxable3Flag

`func (o *GLExportTransactionDetail) GetTaxable3Flag() bool`

GetTaxable3Flag returns the Taxable3Flag field if non-nil, zero value otherwise.

### GetTaxable3FlagOk

`func (o *GLExportTransactionDetail) GetTaxable3FlagOk() (*bool, bool)`

GetTaxable3FlagOk returns a tuple with the Taxable3Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable3Flag

`func (o *GLExportTransactionDetail) SetTaxable3Flag(v bool)`

SetTaxable3Flag sets Taxable3Flag field to given value.

### HasTaxable3Flag

`func (o *GLExportTransactionDetail) HasTaxable3Flag() bool`

HasTaxable3Flag returns a boolean if a field has been set.

### SetTaxable3FlagNil

`func (o *GLExportTransactionDetail) SetTaxable3FlagNil(b bool)`

 SetTaxable3FlagNil sets the value for Taxable3Flag to be an explicit nil

### UnsetTaxable3Flag
`func (o *GLExportTransactionDetail) UnsetTaxable3Flag()`

UnsetTaxable3Flag ensures that no value is present for Taxable3Flag, not even an explicit nil
### GetTaxable4Flag

`func (o *GLExportTransactionDetail) GetTaxable4Flag() bool`

GetTaxable4Flag returns the Taxable4Flag field if non-nil, zero value otherwise.

### GetTaxable4FlagOk

`func (o *GLExportTransactionDetail) GetTaxable4FlagOk() (*bool, bool)`

GetTaxable4FlagOk returns a tuple with the Taxable4Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable4Flag

`func (o *GLExportTransactionDetail) SetTaxable4Flag(v bool)`

SetTaxable4Flag sets Taxable4Flag field to given value.

### HasTaxable4Flag

`func (o *GLExportTransactionDetail) HasTaxable4Flag() bool`

HasTaxable4Flag returns a boolean if a field has been set.

### SetTaxable4FlagNil

`func (o *GLExportTransactionDetail) SetTaxable4FlagNil(b bool)`

 SetTaxable4FlagNil sets the value for Taxable4Flag to be an explicit nil

### UnsetTaxable4Flag
`func (o *GLExportTransactionDetail) UnsetTaxable4Flag()`

UnsetTaxable4Flag ensures that no value is present for Taxable4Flag, not even an explicit nil
### GetTaxable5Flag

`func (o *GLExportTransactionDetail) GetTaxable5Flag() bool`

GetTaxable5Flag returns the Taxable5Flag field if non-nil, zero value otherwise.

### GetTaxable5FlagOk

`func (o *GLExportTransactionDetail) GetTaxable5FlagOk() (*bool, bool)`

GetTaxable5FlagOk returns a tuple with the Taxable5Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable5Flag

`func (o *GLExportTransactionDetail) SetTaxable5Flag(v bool)`

SetTaxable5Flag sets Taxable5Flag field to given value.

### HasTaxable5Flag

`func (o *GLExportTransactionDetail) HasTaxable5Flag() bool`

HasTaxable5Flag returns a boolean if a field has been set.

### SetTaxable5FlagNil

`func (o *GLExportTransactionDetail) SetTaxable5FlagNil(b bool)`

 SetTaxable5FlagNil sets the value for Taxable5Flag to be an explicit nil

### UnsetTaxable5Flag
`func (o *GLExportTransactionDetail) UnsetTaxable5Flag()`

UnsetTaxable5Flag ensures that no value is present for Taxable5Flag, not even an explicit nil
### GetItem

`func (o *GLExportTransactionDetail) GetItem() IvItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *GLExportTransactionDetail) GetItemOk() (*IvItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *GLExportTransactionDetail) SetItem(v IvItemReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *GLExportTransactionDetail) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetProduct

`func (o *GLExportTransactionDetail) GetProduct() ProductReference`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *GLExportTransactionDetail) GetProductOk() (*ProductReference, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *GLExportTransactionDetail) SetProduct(v ProductReference)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *GLExportTransactionDetail) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetItemTaxableFlag

`func (o *GLExportTransactionDetail) GetItemTaxableFlag() bool`

GetItemTaxableFlag returns the ItemTaxableFlag field if non-nil, zero value otherwise.

### GetItemTaxableFlagOk

`func (o *GLExportTransactionDetail) GetItemTaxableFlagOk() (*bool, bool)`

GetItemTaxableFlagOk returns a tuple with the ItemTaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTaxableFlag

`func (o *GLExportTransactionDetail) SetItemTaxableFlag(v bool)`

SetItemTaxableFlag sets ItemTaxableFlag field to given value.

### HasItemTaxableFlag

`func (o *GLExportTransactionDetail) HasItemTaxableFlag() bool`

HasItemTaxableFlag returns a boolean if a field has been set.

### SetItemTaxableFlagNil

`func (o *GLExportTransactionDetail) SetItemTaxableFlagNil(b bool)`

 SetItemTaxableFlagNil sets the value for ItemTaxableFlag to be an explicit nil

### UnsetItemTaxableFlag
`func (o *GLExportTransactionDetail) UnsetItemTaxableFlag()`

UnsetItemTaxableFlag ensures that no value is present for ItemTaxableFlag, not even an explicit nil
### GetItemPrice

`func (o *GLExportTransactionDetail) GetItemPrice() float64`

GetItemPrice returns the ItemPrice field if non-nil, zero value otherwise.

### GetItemPriceOk

`func (o *GLExportTransactionDetail) GetItemPriceOk() (*float64, bool)`

GetItemPriceOk returns a tuple with the ItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrice

`func (o *GLExportTransactionDetail) SetItemPrice(v float64)`

SetItemPrice sets ItemPrice field to given value.

### HasItemPrice

`func (o *GLExportTransactionDetail) HasItemPrice() bool`

HasItemPrice returns a boolean if a field has been set.

### SetItemPriceNil

`func (o *GLExportTransactionDetail) SetItemPriceNil(b bool)`

 SetItemPriceNil sets the value for ItemPrice to be an explicit nil

### UnsetItemPrice
`func (o *GLExportTransactionDetail) UnsetItemPrice()`

UnsetItemPrice ensures that no value is present for ItemPrice, not even an explicit nil
### GetItemCost

`func (o *GLExportTransactionDetail) GetItemCost() float64`

GetItemCost returns the ItemCost field if non-nil, zero value otherwise.

### GetItemCostOk

`func (o *GLExportTransactionDetail) GetItemCostOk() (*float64, bool)`

GetItemCostOk returns a tuple with the ItemCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCost

`func (o *GLExportTransactionDetail) SetItemCost(v float64)`

SetItemCost sets ItemCost field to given value.

### HasItemCost

`func (o *GLExportTransactionDetail) HasItemCost() bool`

HasItemCost returns a boolean if a field has been set.

### SetItemCostNil

`func (o *GLExportTransactionDetail) SetItemCostNil(b bool)`

 SetItemCostNil sets the value for ItemCost to be an explicit nil

### UnsetItemCost
`func (o *GLExportTransactionDetail) UnsetItemCost()`

UnsetItemCost ensures that no value is present for ItemCost, not even an explicit nil
### GetItemDescription

`func (o *GLExportTransactionDetail) GetItemDescription() string`

GetItemDescription returns the ItemDescription field if non-nil, zero value otherwise.

### GetItemDescriptionOk

`func (o *GLExportTransactionDetail) GetItemDescriptionOk() (*string, bool)`

GetItemDescriptionOk returns a tuple with the ItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDescription

`func (o *GLExportTransactionDetail) SetItemDescription(v string)`

SetItemDescription sets ItemDescription field to given value.

### HasItemDescription

`func (o *GLExportTransactionDetail) HasItemDescription() bool`

HasItemDescription returns a boolean if a field has been set.

### GetSalesDescription

`func (o *GLExportTransactionDetail) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *GLExportTransactionDetail) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *GLExportTransactionDetail) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.

### HasSalesDescription

`func (o *GLExportTransactionDetail) HasSalesDescription() bool`

HasSalesDescription returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *GLExportTransactionDetail) GetUnitOfMeasure() UnitOfMeasureReference`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *GLExportTransactionDetail) GetUnitOfMeasureOk() (*UnitOfMeasureReference, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *GLExportTransactionDetail) SetUnitOfMeasure(v UnitOfMeasureReference)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *GLExportTransactionDetail) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetSubCategory

`func (o *GLExportTransactionDetail) GetSubCategory() ProductSubCategoryReference`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *GLExportTransactionDetail) GetSubCategoryOk() (*ProductSubCategoryReference, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *GLExportTransactionDetail) SetSubCategory(v ProductSubCategoryReference)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *GLExportTransactionDetail) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetSerializedFlag

`func (o *GLExportTransactionDetail) GetSerializedFlag() bool`

GetSerializedFlag returns the SerializedFlag field if non-nil, zero value otherwise.

### GetSerializedFlagOk

`func (o *GLExportTransactionDetail) GetSerializedFlagOk() (*bool, bool)`

GetSerializedFlagOk returns a tuple with the SerializedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedFlag

`func (o *GLExportTransactionDetail) SetSerializedFlag(v bool)`

SetSerializedFlag sets SerializedFlag field to given value.

### HasSerializedFlag

`func (o *GLExportTransactionDetail) HasSerializedFlag() bool`

HasSerializedFlag returns a boolean if a field has been set.

### SetSerializedFlagNil

`func (o *GLExportTransactionDetail) SetSerializedFlagNil(b bool)`

 SetSerializedFlagNil sets the value for SerializedFlag to be an explicit nil

### UnsetSerializedFlag
`func (o *GLExportTransactionDetail) UnsetSerializedFlag()`

UnsetSerializedFlag ensures that no value is present for SerializedFlag, not even an explicit nil
### GetSerialNumbers

`func (o *GLExportTransactionDetail) GetSerialNumbers() string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *GLExportTransactionDetail) GetSerialNumbersOk() (*string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *GLExportTransactionDetail) SetSerialNumbers(v string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *GLExportTransactionDetail) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.

### GetWarehouseSite

`func (o *GLExportTransactionDetail) GetWarehouseSite() SiteReference`

GetWarehouseSite returns the WarehouseSite field if non-nil, zero value otherwise.

### GetWarehouseSiteOk

`func (o *GLExportTransactionDetail) GetWarehouseSiteOk() (*SiteReference, bool)`

GetWarehouseSiteOk returns a tuple with the WarehouseSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseSite

`func (o *GLExportTransactionDetail) SetWarehouseSite(v SiteReference)`

SetWarehouseSite sets WarehouseSite field to given value.

### HasWarehouseSite

`func (o *GLExportTransactionDetail) HasWarehouseSite() bool`

HasWarehouseSite returns a boolean if a field has been set.

### GetWarehouseBin

`func (o *GLExportTransactionDetail) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *GLExportTransactionDetail) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *GLExportTransactionDetail) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.

### HasWarehouseBin

`func (o *GLExportTransactionDetail) HasWarehouseBin() bool`

HasWarehouseBin returns a boolean if a field has been set.

### GetShipmentMethod

`func (o *GLExportTransactionDetail) GetShipmentMethod() ShipmentMethodReference`

GetShipmentMethod returns the ShipmentMethod field if non-nil, zero value otherwise.

### GetShipmentMethodOk

`func (o *GLExportTransactionDetail) GetShipmentMethodOk() (*ShipmentMethodReference, bool)`

GetShipmentMethodOk returns a tuple with the ShipmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentMethod

`func (o *GLExportTransactionDetail) SetShipmentMethod(v ShipmentMethodReference)`

SetShipmentMethod sets ShipmentMethod field to given value.

### HasShipmentMethod

`func (o *GLExportTransactionDetail) HasShipmentMethod() bool`

HasShipmentMethod returns a boolean if a field has been set.

### GetDropShippedFlag

`func (o *GLExportTransactionDetail) GetDropShippedFlag() bool`

GetDropShippedFlag returns the DropShippedFlag field if non-nil, zero value otherwise.

### GetDropShippedFlagOk

`func (o *GLExportTransactionDetail) GetDropShippedFlagOk() (*bool, bool)`

GetDropShippedFlagOk returns a tuple with the DropShippedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropShippedFlag

`func (o *GLExportTransactionDetail) SetDropShippedFlag(v bool)`

SetDropShippedFlag sets DropShippedFlag field to given value.

### HasDropShippedFlag

`func (o *GLExportTransactionDetail) HasDropShippedFlag() bool`

HasDropShippedFlag returns a boolean if a field has been set.

### SetDropShippedFlagNil

`func (o *GLExportTransactionDetail) SetDropShippedFlagNil(b bool)`

 SetDropShippedFlagNil sets the value for DropShippedFlag to be an explicit nil

### UnsetDropShippedFlag
`func (o *GLExportTransactionDetail) UnsetDropShippedFlag()`

UnsetDropShippedFlag ensures that no value is present for DropShippedFlag, not even an explicit nil
### GetItemTypeXref

`func (o *GLExportTransactionDetail) GetItemTypeXref() string`

GetItemTypeXref returns the ItemTypeXref field if non-nil, zero value otherwise.

### GetItemTypeXrefOk

`func (o *GLExportTransactionDetail) GetItemTypeXrefOk() (*string, bool)`

GetItemTypeXrefOk returns a tuple with the ItemTypeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTypeXref

`func (o *GLExportTransactionDetail) SetItemTypeXref(v string)`

SetItemTypeXref sets ItemTypeXref field to given value.

### HasItemTypeXref

`func (o *GLExportTransactionDetail) HasItemTypeXref() bool`

HasItemTypeXref returns a boolean if a field has been set.

### GetInventoryXref

`func (o *GLExportTransactionDetail) GetInventoryXref() string`

GetInventoryXref returns the InventoryXref field if non-nil, zero value otherwise.

### GetInventoryXrefOk

`func (o *GLExportTransactionDetail) GetInventoryXrefOk() (*string, bool)`

GetInventoryXrefOk returns a tuple with the InventoryXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryXref

`func (o *GLExportTransactionDetail) SetInventoryXref(v string)`

SetInventoryXref sets InventoryXref field to given value.

### HasInventoryXref

`func (o *GLExportTransactionDetail) HasInventoryXref() bool`

HasInventoryXref returns a boolean if a field has been set.

### GetCogsXref

`func (o *GLExportTransactionDetail) GetCogsXref() string`

GetCogsXref returns the CogsXref field if non-nil, zero value otherwise.

### GetCogsXrefOk

`func (o *GLExportTransactionDetail) GetCogsXrefOk() (*string, bool)`

GetCogsXrefOk returns a tuple with the CogsXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogsXref

`func (o *GLExportTransactionDetail) SetCogsXref(v string)`

SetCogsXref sets CogsXref field to given value.

### HasCogsXref

`func (o *GLExportTransactionDetail) HasCogsXref() bool`

HasCogsXref returns a boolean if a field has been set.

### GetUomScheduleXref

`func (o *GLExportTransactionDetail) GetUomScheduleXref() string`

GetUomScheduleXref returns the UomScheduleXref field if non-nil, zero value otherwise.

### GetUomScheduleXrefOk

`func (o *GLExportTransactionDetail) GetUomScheduleXrefOk() (*string, bool)`

GetUomScheduleXrefOk returns a tuple with the UomScheduleXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUomScheduleXref

`func (o *GLExportTransactionDetail) SetUomScheduleXref(v string)`

SetUomScheduleXref sets UomScheduleXref field to given value.

### HasUomScheduleXref

`func (o *GLExportTransactionDetail) HasUomScheduleXref() bool`

HasUomScheduleXref returns a boolean if a field has been set.

### GetPriceLevelXref

`func (o *GLExportTransactionDetail) GetPriceLevelXref() string`

GetPriceLevelXref returns the PriceLevelXref field if non-nil, zero value otherwise.

### GetPriceLevelXrefOk

`func (o *GLExportTransactionDetail) GetPriceLevelXrefOk() (*string, bool)`

GetPriceLevelXrefOk returns a tuple with the PriceLevelXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevelXref

`func (o *GLExportTransactionDetail) SetPriceLevelXref(v string)`

SetPriceLevelXref sets PriceLevelXref field to given value.

### HasPriceLevelXref

`func (o *GLExportTransactionDetail) HasPriceLevelXref() bool`

HasPriceLevelXref returns a boolean if a field has been set.

### GetLocationXref

`func (o *GLExportTransactionDetail) GetLocationXref() string`

GetLocationXref returns the LocationXref field if non-nil, zero value otherwise.

### GetLocationXrefOk

`func (o *GLExportTransactionDetail) GetLocationXrefOk() (*string, bool)`

GetLocationXrefOk returns a tuple with the LocationXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationXref

`func (o *GLExportTransactionDetail) SetLocationXref(v string)`

SetLocationXref sets LocationXref field to given value.

### HasLocationXref

`func (o *GLExportTransactionDetail) HasLocationXref() bool`

HasLocationXref returns a boolean if a field has been set.

### GetTaxLevels

`func (o *GLExportTransactionDetail) GetTaxLevels() []GLExportTransactionDetailTaxLevel`

GetTaxLevels returns the TaxLevels field if non-nil, zero value otherwise.

### GetTaxLevelsOk

`func (o *GLExportTransactionDetail) GetTaxLevelsOk() (*[]GLExportTransactionDetailTaxLevel, bool)`

GetTaxLevelsOk returns a tuple with the TaxLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLevels

`func (o *GLExportTransactionDetail) SetTaxLevels(v []GLExportTransactionDetailTaxLevel)`

SetTaxLevels sets TaxLevels field to given value.

### HasTaxLevels

`func (o *GLExportTransactionDetail) HasTaxLevels() bool`

HasTaxLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


