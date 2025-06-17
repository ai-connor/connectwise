# GLExportPurchaseTransactionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**DocumentDate** | Pointer to **string** |  | [optional] 
**GlClass** | Pointer to **string** |  | [optional] 
**GlTypeId** | Pointer to **string** |  | [optional] 
**GlItemId** | Pointer to **string** |  | [optional] 
**SalesCode** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**TaxNote** | Pointer to **string** |  | [optional] 
**VendorNumber** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**CostAccountNumber** | Pointer to **string** |  | [optional] 
**InventoryAccountNumber** | Pointer to **string** |  | [optional] 
**VendorAccountNumber** | Pointer to **string** |  | [optional] 
**Item** | Pointer to [**IvItemReference**](IvItemReference.md) |  | [optional] 
**ItemDescription** | Pointer to **string** |  | [optional] 
**SalesDescription** | Pointer to **string** |  | [optional] 
**Taxable** | Pointer to **NullableBool** |  | [optional] 
**ItemPrice** | Pointer to **NullableFloat64** |  | [optional] 
**ItemCost** | Pointer to **NullableFloat64** |  | [optional] 
**UnitOfMeasure** | Pointer to [**UnitOfMeasureReference**](UnitOfMeasureReference.md) |  | [optional] 
**Quantity** | Pointer to **NullableFloat64** |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**SerializedFlag** | Pointer to **NullableBool** |  | [optional] 
**SerialNumbers** | Pointer to **string** |  | [optional] 
**DropShippedFlag** | Pointer to **NullableBool** |  | [optional] 
**LineNumber** | Pointer to **int32** |  | [optional] 
**WarehouseBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**WarehouseSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**SubCategory** | Pointer to [**ProductSubCategoryReference**](ProductSubCategoryReference.md) |  | [optional] 
**ShipmentMethod** | Pointer to [**ShipmentMethodReference**](ShipmentMethodReference.md) |  | [optional] 
**ItemTypeXref** | Pointer to **string** |  | [optional] 
**InventoryXref** | Pointer to **string** |  | [optional] 
**CogsXref** | Pointer to **string** |  | [optional] 
**UomScheduleXref** | Pointer to **string** |  | [optional] 
**PriceLevelXref** | Pointer to **string** |  | [optional] 
**LocationXref** | Pointer to **string** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**PurchaseHeaderTaxGroup** | Pointer to **string** |  | [optional] 
**TaxCodeXref** | Pointer to **string** |  | [optional] 
**TaxRate** | Pointer to **float64** |  | [optional] 
**TaxAgencyXref** | Pointer to **string** |  | [optional] 

## Methods

### NewGLExportPurchaseTransactionDetail

`func NewGLExportPurchaseTransactionDetail() *GLExportPurchaseTransactionDetail`

NewGLExportPurchaseTransactionDetail instantiates a new GLExportPurchaseTransactionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportPurchaseTransactionDetailWithDefaults

`func NewGLExportPurchaseTransactionDetailWithDefaults() *GLExportPurchaseTransactionDetail`

NewGLExportPurchaseTransactionDetailWithDefaults instantiates a new GLExportPurchaseTransactionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLExportPurchaseTransactionDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLExportPurchaseTransactionDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLExportPurchaseTransactionDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GLExportPurchaseTransactionDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GLExportPurchaseTransactionDetail) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GLExportPurchaseTransactionDetail) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDocumentDate

`func (o *GLExportPurchaseTransactionDetail) GetDocumentDate() string`

GetDocumentDate returns the DocumentDate field if non-nil, zero value otherwise.

### GetDocumentDateOk

`func (o *GLExportPurchaseTransactionDetail) GetDocumentDateOk() (*string, bool)`

GetDocumentDateOk returns a tuple with the DocumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDate

`func (o *GLExportPurchaseTransactionDetail) SetDocumentDate(v string)`

SetDocumentDate sets DocumentDate field to given value.

### HasDocumentDate

`func (o *GLExportPurchaseTransactionDetail) HasDocumentDate() bool`

HasDocumentDate returns a boolean if a field has been set.

### GetGlClass

`func (o *GLExportPurchaseTransactionDetail) GetGlClass() string`

GetGlClass returns the GlClass field if non-nil, zero value otherwise.

### GetGlClassOk

`func (o *GLExportPurchaseTransactionDetail) GetGlClassOk() (*string, bool)`

GetGlClassOk returns a tuple with the GlClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlClass

`func (o *GLExportPurchaseTransactionDetail) SetGlClass(v string)`

SetGlClass sets GlClass field to given value.

### HasGlClass

`func (o *GLExportPurchaseTransactionDetail) HasGlClass() bool`

HasGlClass returns a boolean if a field has been set.

### GetGlTypeId

`func (o *GLExportPurchaseTransactionDetail) GetGlTypeId() string`

GetGlTypeId returns the GlTypeId field if non-nil, zero value otherwise.

### GetGlTypeIdOk

`func (o *GLExportPurchaseTransactionDetail) GetGlTypeIdOk() (*string, bool)`

GetGlTypeIdOk returns a tuple with the GlTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlTypeId

`func (o *GLExportPurchaseTransactionDetail) SetGlTypeId(v string)`

SetGlTypeId sets GlTypeId field to given value.

### HasGlTypeId

`func (o *GLExportPurchaseTransactionDetail) HasGlTypeId() bool`

HasGlTypeId returns a boolean if a field has been set.

### GetGlItemId

`func (o *GLExportPurchaseTransactionDetail) GetGlItemId() string`

GetGlItemId returns the GlItemId field if non-nil, zero value otherwise.

### GetGlItemIdOk

`func (o *GLExportPurchaseTransactionDetail) GetGlItemIdOk() (*string, bool)`

GetGlItemIdOk returns a tuple with the GlItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlItemId

`func (o *GLExportPurchaseTransactionDetail) SetGlItemId(v string)`

SetGlItemId sets GlItemId field to given value.

### HasGlItemId

`func (o *GLExportPurchaseTransactionDetail) HasGlItemId() bool`

HasGlItemId returns a boolean if a field has been set.

### GetSalesCode

`func (o *GLExportPurchaseTransactionDetail) GetSalesCode() string`

GetSalesCode returns the SalesCode field if non-nil, zero value otherwise.

### GetSalesCodeOk

`func (o *GLExportPurchaseTransactionDetail) GetSalesCodeOk() (*string, bool)`

GetSalesCodeOk returns a tuple with the SalesCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesCode

`func (o *GLExportPurchaseTransactionDetail) SetSalesCode(v string)`

SetSalesCode sets SalesCode field to given value.

### HasSalesCode

`func (o *GLExportPurchaseTransactionDetail) HasSalesCode() bool`

HasSalesCode returns a boolean if a field has been set.

### GetDescription

`func (o *GLExportPurchaseTransactionDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GLExportPurchaseTransactionDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GLExportPurchaseTransactionDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GLExportPurchaseTransactionDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCost

`func (o *GLExportPurchaseTransactionDetail) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GLExportPurchaseTransactionDetail) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GLExportPurchaseTransactionDetail) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GLExportPurchaseTransactionDetail) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *GLExportPurchaseTransactionDetail) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *GLExportPurchaseTransactionDetail) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetMemo

`func (o *GLExportPurchaseTransactionDetail) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GLExportPurchaseTransactionDetail) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GLExportPurchaseTransactionDetail) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GLExportPurchaseTransactionDetail) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetTaxNote

`func (o *GLExportPurchaseTransactionDetail) GetTaxNote() string`

GetTaxNote returns the TaxNote field if non-nil, zero value otherwise.

### GetTaxNoteOk

`func (o *GLExportPurchaseTransactionDetail) GetTaxNoteOk() (*string, bool)`

GetTaxNoteOk returns a tuple with the TaxNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNote

`func (o *GLExportPurchaseTransactionDetail) SetTaxNote(v string)`

SetTaxNote sets TaxNote field to given value.

### HasTaxNote

`func (o *GLExportPurchaseTransactionDetail) HasTaxNote() bool`

HasTaxNote returns a boolean if a field has been set.

### GetVendorNumber

`func (o *GLExportPurchaseTransactionDetail) GetVendorNumber() string`

GetVendorNumber returns the VendorNumber field if non-nil, zero value otherwise.

### GetVendorNumberOk

`func (o *GLExportPurchaseTransactionDetail) GetVendorNumberOk() (*string, bool)`

GetVendorNumberOk returns a tuple with the VendorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumber

`func (o *GLExportPurchaseTransactionDetail) SetVendorNumber(v string)`

SetVendorNumber sets VendorNumber field to given value.

### HasVendorNumber

`func (o *GLExportPurchaseTransactionDetail) HasVendorNumber() bool`

HasVendorNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *GLExportPurchaseTransactionDetail) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GLExportPurchaseTransactionDetail) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GLExportPurchaseTransactionDetail) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GLExportPurchaseTransactionDetail) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetCostAccountNumber

`func (o *GLExportPurchaseTransactionDetail) GetCostAccountNumber() string`

GetCostAccountNumber returns the CostAccountNumber field if non-nil, zero value otherwise.

### GetCostAccountNumberOk

`func (o *GLExportPurchaseTransactionDetail) GetCostAccountNumberOk() (*string, bool)`

GetCostAccountNumberOk returns a tuple with the CostAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAccountNumber

`func (o *GLExportPurchaseTransactionDetail) SetCostAccountNumber(v string)`

SetCostAccountNumber sets CostAccountNumber field to given value.

### HasCostAccountNumber

`func (o *GLExportPurchaseTransactionDetail) HasCostAccountNumber() bool`

HasCostAccountNumber returns a boolean if a field has been set.

### GetInventoryAccountNumber

`func (o *GLExportPurchaseTransactionDetail) GetInventoryAccountNumber() string`

GetInventoryAccountNumber returns the InventoryAccountNumber field if non-nil, zero value otherwise.

### GetInventoryAccountNumberOk

`func (o *GLExportPurchaseTransactionDetail) GetInventoryAccountNumberOk() (*string, bool)`

GetInventoryAccountNumberOk returns a tuple with the InventoryAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryAccountNumber

`func (o *GLExportPurchaseTransactionDetail) SetInventoryAccountNumber(v string)`

SetInventoryAccountNumber sets InventoryAccountNumber field to given value.

### HasInventoryAccountNumber

`func (o *GLExportPurchaseTransactionDetail) HasInventoryAccountNumber() bool`

HasInventoryAccountNumber returns a boolean if a field has been set.

### GetVendorAccountNumber

`func (o *GLExportPurchaseTransactionDetail) GetVendorAccountNumber() string`

GetVendorAccountNumber returns the VendorAccountNumber field if non-nil, zero value otherwise.

### GetVendorAccountNumberOk

`func (o *GLExportPurchaseTransactionDetail) GetVendorAccountNumberOk() (*string, bool)`

GetVendorAccountNumberOk returns a tuple with the VendorAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAccountNumber

`func (o *GLExportPurchaseTransactionDetail) SetVendorAccountNumber(v string)`

SetVendorAccountNumber sets VendorAccountNumber field to given value.

### HasVendorAccountNumber

`func (o *GLExportPurchaseTransactionDetail) HasVendorAccountNumber() bool`

HasVendorAccountNumber returns a boolean if a field has been set.

### GetItem

`func (o *GLExportPurchaseTransactionDetail) GetItem() IvItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *GLExportPurchaseTransactionDetail) GetItemOk() (*IvItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *GLExportPurchaseTransactionDetail) SetItem(v IvItemReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *GLExportPurchaseTransactionDetail) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetItemDescription

`func (o *GLExportPurchaseTransactionDetail) GetItemDescription() string`

GetItemDescription returns the ItemDescription field if non-nil, zero value otherwise.

### GetItemDescriptionOk

`func (o *GLExportPurchaseTransactionDetail) GetItemDescriptionOk() (*string, bool)`

GetItemDescriptionOk returns a tuple with the ItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDescription

`func (o *GLExportPurchaseTransactionDetail) SetItemDescription(v string)`

SetItemDescription sets ItemDescription field to given value.

### HasItemDescription

`func (o *GLExportPurchaseTransactionDetail) HasItemDescription() bool`

HasItemDescription returns a boolean if a field has been set.

### GetSalesDescription

`func (o *GLExportPurchaseTransactionDetail) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *GLExportPurchaseTransactionDetail) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *GLExportPurchaseTransactionDetail) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.

### HasSalesDescription

`func (o *GLExportPurchaseTransactionDetail) HasSalesDescription() bool`

HasSalesDescription returns a boolean if a field has been set.

### GetTaxable

`func (o *GLExportPurchaseTransactionDetail) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *GLExportPurchaseTransactionDetail) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *GLExportPurchaseTransactionDetail) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *GLExportPurchaseTransactionDetail) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### SetTaxableNil

`func (o *GLExportPurchaseTransactionDetail) SetTaxableNil(b bool)`

 SetTaxableNil sets the value for Taxable to be an explicit nil

### UnsetTaxable
`func (o *GLExportPurchaseTransactionDetail) UnsetTaxable()`

UnsetTaxable ensures that no value is present for Taxable, not even an explicit nil
### GetItemPrice

`func (o *GLExportPurchaseTransactionDetail) GetItemPrice() float64`

GetItemPrice returns the ItemPrice field if non-nil, zero value otherwise.

### GetItemPriceOk

`func (o *GLExportPurchaseTransactionDetail) GetItemPriceOk() (*float64, bool)`

GetItemPriceOk returns a tuple with the ItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrice

`func (o *GLExportPurchaseTransactionDetail) SetItemPrice(v float64)`

SetItemPrice sets ItemPrice field to given value.

### HasItemPrice

`func (o *GLExportPurchaseTransactionDetail) HasItemPrice() bool`

HasItemPrice returns a boolean if a field has been set.

### SetItemPriceNil

`func (o *GLExportPurchaseTransactionDetail) SetItemPriceNil(b bool)`

 SetItemPriceNil sets the value for ItemPrice to be an explicit nil

### UnsetItemPrice
`func (o *GLExportPurchaseTransactionDetail) UnsetItemPrice()`

UnsetItemPrice ensures that no value is present for ItemPrice, not even an explicit nil
### GetItemCost

`func (o *GLExportPurchaseTransactionDetail) GetItemCost() float64`

GetItemCost returns the ItemCost field if non-nil, zero value otherwise.

### GetItemCostOk

`func (o *GLExportPurchaseTransactionDetail) GetItemCostOk() (*float64, bool)`

GetItemCostOk returns a tuple with the ItemCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCost

`func (o *GLExportPurchaseTransactionDetail) SetItemCost(v float64)`

SetItemCost sets ItemCost field to given value.

### HasItemCost

`func (o *GLExportPurchaseTransactionDetail) HasItemCost() bool`

HasItemCost returns a boolean if a field has been set.

### SetItemCostNil

`func (o *GLExportPurchaseTransactionDetail) SetItemCostNil(b bool)`

 SetItemCostNil sets the value for ItemCost to be an explicit nil

### UnsetItemCost
`func (o *GLExportPurchaseTransactionDetail) UnsetItemCost()`

UnsetItemCost ensures that no value is present for ItemCost, not even an explicit nil
### GetUnitOfMeasure

`func (o *GLExportPurchaseTransactionDetail) GetUnitOfMeasure() UnitOfMeasureReference`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *GLExportPurchaseTransactionDetail) GetUnitOfMeasureOk() (*UnitOfMeasureReference, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *GLExportPurchaseTransactionDetail) SetUnitOfMeasure(v UnitOfMeasureReference)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *GLExportPurchaseTransactionDetail) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetQuantity

`func (o *GLExportPurchaseTransactionDetail) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GLExportPurchaseTransactionDetail) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GLExportPurchaseTransactionDetail) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GLExportPurchaseTransactionDetail) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *GLExportPurchaseTransactionDetail) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *GLExportPurchaseTransactionDetail) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetTotal

`func (o *GLExportPurchaseTransactionDetail) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GLExportPurchaseTransactionDetail) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GLExportPurchaseTransactionDetail) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GLExportPurchaseTransactionDetail) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GLExportPurchaseTransactionDetail) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GLExportPurchaseTransactionDetail) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetCurrency

`func (o *GLExportPurchaseTransactionDetail) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GLExportPurchaseTransactionDetail) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GLExportPurchaseTransactionDetail) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GLExportPurchaseTransactionDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSerializedFlag

`func (o *GLExportPurchaseTransactionDetail) GetSerializedFlag() bool`

GetSerializedFlag returns the SerializedFlag field if non-nil, zero value otherwise.

### GetSerializedFlagOk

`func (o *GLExportPurchaseTransactionDetail) GetSerializedFlagOk() (*bool, bool)`

GetSerializedFlagOk returns a tuple with the SerializedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedFlag

`func (o *GLExportPurchaseTransactionDetail) SetSerializedFlag(v bool)`

SetSerializedFlag sets SerializedFlag field to given value.

### HasSerializedFlag

`func (o *GLExportPurchaseTransactionDetail) HasSerializedFlag() bool`

HasSerializedFlag returns a boolean if a field has been set.

### SetSerializedFlagNil

`func (o *GLExportPurchaseTransactionDetail) SetSerializedFlagNil(b bool)`

 SetSerializedFlagNil sets the value for SerializedFlag to be an explicit nil

### UnsetSerializedFlag
`func (o *GLExportPurchaseTransactionDetail) UnsetSerializedFlag()`

UnsetSerializedFlag ensures that no value is present for SerializedFlag, not even an explicit nil
### GetSerialNumbers

`func (o *GLExportPurchaseTransactionDetail) GetSerialNumbers() string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *GLExportPurchaseTransactionDetail) GetSerialNumbersOk() (*string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *GLExportPurchaseTransactionDetail) SetSerialNumbers(v string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *GLExportPurchaseTransactionDetail) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.

### GetDropShippedFlag

`func (o *GLExportPurchaseTransactionDetail) GetDropShippedFlag() bool`

GetDropShippedFlag returns the DropShippedFlag field if non-nil, zero value otherwise.

### GetDropShippedFlagOk

`func (o *GLExportPurchaseTransactionDetail) GetDropShippedFlagOk() (*bool, bool)`

GetDropShippedFlagOk returns a tuple with the DropShippedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropShippedFlag

`func (o *GLExportPurchaseTransactionDetail) SetDropShippedFlag(v bool)`

SetDropShippedFlag sets DropShippedFlag field to given value.

### HasDropShippedFlag

`func (o *GLExportPurchaseTransactionDetail) HasDropShippedFlag() bool`

HasDropShippedFlag returns a boolean if a field has been set.

### SetDropShippedFlagNil

`func (o *GLExportPurchaseTransactionDetail) SetDropShippedFlagNil(b bool)`

 SetDropShippedFlagNil sets the value for DropShippedFlag to be an explicit nil

### UnsetDropShippedFlag
`func (o *GLExportPurchaseTransactionDetail) UnsetDropShippedFlag()`

UnsetDropShippedFlag ensures that no value is present for DropShippedFlag, not even an explicit nil
### GetLineNumber

`func (o *GLExportPurchaseTransactionDetail) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *GLExportPurchaseTransactionDetail) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *GLExportPurchaseTransactionDetail) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *GLExportPurchaseTransactionDetail) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetWarehouseBin

`func (o *GLExportPurchaseTransactionDetail) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *GLExportPurchaseTransactionDetail) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *GLExportPurchaseTransactionDetail) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.

### HasWarehouseBin

`func (o *GLExportPurchaseTransactionDetail) HasWarehouseBin() bool`

HasWarehouseBin returns a boolean if a field has been set.

### GetWarehouseSite

`func (o *GLExportPurchaseTransactionDetail) GetWarehouseSite() SiteReference`

GetWarehouseSite returns the WarehouseSite field if non-nil, zero value otherwise.

### GetWarehouseSiteOk

`func (o *GLExportPurchaseTransactionDetail) GetWarehouseSiteOk() (*SiteReference, bool)`

GetWarehouseSiteOk returns a tuple with the WarehouseSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseSite

`func (o *GLExportPurchaseTransactionDetail) SetWarehouseSite(v SiteReference)`

SetWarehouseSite sets WarehouseSite field to given value.

### HasWarehouseSite

`func (o *GLExportPurchaseTransactionDetail) HasWarehouseSite() bool`

HasWarehouseSite returns a boolean if a field has been set.

### GetSubCategory

`func (o *GLExportPurchaseTransactionDetail) GetSubCategory() ProductSubCategoryReference`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *GLExportPurchaseTransactionDetail) GetSubCategoryOk() (*ProductSubCategoryReference, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *GLExportPurchaseTransactionDetail) SetSubCategory(v ProductSubCategoryReference)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *GLExportPurchaseTransactionDetail) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetShipmentMethod

`func (o *GLExportPurchaseTransactionDetail) GetShipmentMethod() ShipmentMethodReference`

GetShipmentMethod returns the ShipmentMethod field if non-nil, zero value otherwise.

### GetShipmentMethodOk

`func (o *GLExportPurchaseTransactionDetail) GetShipmentMethodOk() (*ShipmentMethodReference, bool)`

GetShipmentMethodOk returns a tuple with the ShipmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentMethod

`func (o *GLExportPurchaseTransactionDetail) SetShipmentMethod(v ShipmentMethodReference)`

SetShipmentMethod sets ShipmentMethod field to given value.

### HasShipmentMethod

`func (o *GLExportPurchaseTransactionDetail) HasShipmentMethod() bool`

HasShipmentMethod returns a boolean if a field has been set.

### GetItemTypeXref

`func (o *GLExportPurchaseTransactionDetail) GetItemTypeXref() string`

GetItemTypeXref returns the ItemTypeXref field if non-nil, zero value otherwise.

### GetItemTypeXrefOk

`func (o *GLExportPurchaseTransactionDetail) GetItemTypeXrefOk() (*string, bool)`

GetItemTypeXrefOk returns a tuple with the ItemTypeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTypeXref

`func (o *GLExportPurchaseTransactionDetail) SetItemTypeXref(v string)`

SetItemTypeXref sets ItemTypeXref field to given value.

### HasItemTypeXref

`func (o *GLExportPurchaseTransactionDetail) HasItemTypeXref() bool`

HasItemTypeXref returns a boolean if a field has been set.

### GetInventoryXref

`func (o *GLExportPurchaseTransactionDetail) GetInventoryXref() string`

GetInventoryXref returns the InventoryXref field if non-nil, zero value otherwise.

### GetInventoryXrefOk

`func (o *GLExportPurchaseTransactionDetail) GetInventoryXrefOk() (*string, bool)`

GetInventoryXrefOk returns a tuple with the InventoryXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryXref

`func (o *GLExportPurchaseTransactionDetail) SetInventoryXref(v string)`

SetInventoryXref sets InventoryXref field to given value.

### HasInventoryXref

`func (o *GLExportPurchaseTransactionDetail) HasInventoryXref() bool`

HasInventoryXref returns a boolean if a field has been set.

### GetCogsXref

`func (o *GLExportPurchaseTransactionDetail) GetCogsXref() string`

GetCogsXref returns the CogsXref field if non-nil, zero value otherwise.

### GetCogsXrefOk

`func (o *GLExportPurchaseTransactionDetail) GetCogsXrefOk() (*string, bool)`

GetCogsXrefOk returns a tuple with the CogsXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogsXref

`func (o *GLExportPurchaseTransactionDetail) SetCogsXref(v string)`

SetCogsXref sets CogsXref field to given value.

### HasCogsXref

`func (o *GLExportPurchaseTransactionDetail) HasCogsXref() bool`

HasCogsXref returns a boolean if a field has been set.

### GetUomScheduleXref

`func (o *GLExportPurchaseTransactionDetail) GetUomScheduleXref() string`

GetUomScheduleXref returns the UomScheduleXref field if non-nil, zero value otherwise.

### GetUomScheduleXrefOk

`func (o *GLExportPurchaseTransactionDetail) GetUomScheduleXrefOk() (*string, bool)`

GetUomScheduleXrefOk returns a tuple with the UomScheduleXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUomScheduleXref

`func (o *GLExportPurchaseTransactionDetail) SetUomScheduleXref(v string)`

SetUomScheduleXref sets UomScheduleXref field to given value.

### HasUomScheduleXref

`func (o *GLExportPurchaseTransactionDetail) HasUomScheduleXref() bool`

HasUomScheduleXref returns a boolean if a field has been set.

### GetPriceLevelXref

`func (o *GLExportPurchaseTransactionDetail) GetPriceLevelXref() string`

GetPriceLevelXref returns the PriceLevelXref field if non-nil, zero value otherwise.

### GetPriceLevelXrefOk

`func (o *GLExportPurchaseTransactionDetail) GetPriceLevelXrefOk() (*string, bool)`

GetPriceLevelXrefOk returns a tuple with the PriceLevelXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevelXref

`func (o *GLExportPurchaseTransactionDetail) SetPriceLevelXref(v string)`

SetPriceLevelXref sets PriceLevelXref field to given value.

### HasPriceLevelXref

`func (o *GLExportPurchaseTransactionDetail) HasPriceLevelXref() bool`

HasPriceLevelXref returns a boolean if a field has been set.

### GetLocationXref

`func (o *GLExportPurchaseTransactionDetail) GetLocationXref() string`

GetLocationXref returns the LocationXref field if non-nil, zero value otherwise.

### GetLocationXrefOk

`func (o *GLExportPurchaseTransactionDetail) GetLocationXrefOk() (*string, bool)`

GetLocationXrefOk returns a tuple with the LocationXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationXref

`func (o *GLExportPurchaseTransactionDetail) SetLocationXref(v string)`

SetLocationXref sets LocationXref field to given value.

### HasLocationXref

`func (o *GLExportPurchaseTransactionDetail) HasLocationXref() bool`

HasLocationXref returns a boolean if a field has been set.

### GetTaxCode

`func (o *GLExportPurchaseTransactionDetail) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *GLExportPurchaseTransactionDetail) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *GLExportPurchaseTransactionDetail) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *GLExportPurchaseTransactionDetail) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetPurchaseHeaderTaxGroup

`func (o *GLExportPurchaseTransactionDetail) GetPurchaseHeaderTaxGroup() string`

GetPurchaseHeaderTaxGroup returns the PurchaseHeaderTaxGroup field if non-nil, zero value otherwise.

### GetPurchaseHeaderTaxGroupOk

`func (o *GLExportPurchaseTransactionDetail) GetPurchaseHeaderTaxGroupOk() (*string, bool)`

GetPurchaseHeaderTaxGroupOk returns a tuple with the PurchaseHeaderTaxGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseHeaderTaxGroup

`func (o *GLExportPurchaseTransactionDetail) SetPurchaseHeaderTaxGroup(v string)`

SetPurchaseHeaderTaxGroup sets PurchaseHeaderTaxGroup field to given value.

### HasPurchaseHeaderTaxGroup

`func (o *GLExportPurchaseTransactionDetail) HasPurchaseHeaderTaxGroup() bool`

HasPurchaseHeaderTaxGroup returns a boolean if a field has been set.

### GetTaxCodeXref

`func (o *GLExportPurchaseTransactionDetail) GetTaxCodeXref() string`

GetTaxCodeXref returns the TaxCodeXref field if non-nil, zero value otherwise.

### GetTaxCodeXrefOk

`func (o *GLExportPurchaseTransactionDetail) GetTaxCodeXrefOk() (*string, bool)`

GetTaxCodeXrefOk returns a tuple with the TaxCodeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodeXref

`func (o *GLExportPurchaseTransactionDetail) SetTaxCodeXref(v string)`

SetTaxCodeXref sets TaxCodeXref field to given value.

### HasTaxCodeXref

`func (o *GLExportPurchaseTransactionDetail) HasTaxCodeXref() bool`

HasTaxCodeXref returns a boolean if a field has been set.

### GetTaxRate

`func (o *GLExportPurchaseTransactionDetail) GetTaxRate() float64`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *GLExportPurchaseTransactionDetail) GetTaxRateOk() (*float64, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *GLExportPurchaseTransactionDetail) SetTaxRate(v float64)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *GLExportPurchaseTransactionDetail) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxAgencyXref

`func (o *GLExportPurchaseTransactionDetail) GetTaxAgencyXref() string`

GetTaxAgencyXref returns the TaxAgencyXref field if non-nil, zero value otherwise.

### GetTaxAgencyXrefOk

`func (o *GLExportPurchaseTransactionDetail) GetTaxAgencyXrefOk() (*string, bool)`

GetTaxAgencyXrefOk returns a tuple with the TaxAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAgencyXref

`func (o *GLExportPurchaseTransactionDetail) SetTaxAgencyXref(v string)`

SetTaxAgencyXref sets TaxAgencyXref field to given value.

### HasTaxAgencyXref

`func (o *GLExportPurchaseTransactionDetail) HasTaxAgencyXref() bool`

HasTaxAgencyXref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


