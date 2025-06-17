# GLExportPurchaseTransactionDetailTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**DocumentDate** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**GlClass** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**SalesCode** | Pointer to **string** |  | [optional] 
**GlTypeId** | Pointer to **string** |  | [optional] 
**GlItemId** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**VendorNumber** | Pointer to **string** |  | [optional] 
**VendorAccountNumber** | Pointer to **string** |  | [optional] 
**CostAccountNumber** | Pointer to **string** |  | [optional] 
**InventoryAccountNumber** | Pointer to **string** |  | [optional] 
**ItemTypeXref** | Pointer to **string** |  | [optional] 
**InventoryXref** | Pointer to **string** |  | [optional] 
**CogsXref** | Pointer to **string** |  | [optional] 
**UomScheduleXref** | Pointer to **string** |  | [optional] 
**PriceLevelXref** | Pointer to **string** |  | [optional] 
**LocationXref** | Pointer to **string** |  | [optional] 
**Item** | Pointer to [**IvItemReference**](IvItemReference.md) |  | [optional] 
**TaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**SalesDescription** | Pointer to **string** |  | [optional] 
**ItemDescription** | Pointer to **string** |  | [optional] 
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
**WarehouseSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**WarehouseBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**ShipmentMethod** | Pointer to [**ShipmentMethodReference**](ShipmentMethodReference.md) |  | [optional] 
**SubCategory** | Pointer to [**ProductSubCategoryReference**](ProductSubCategoryReference.md) |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**TaxRate** | Pointer to **NullableFloat64** |  | [optional] 
**TaxRatePercent** | Pointer to **NullableFloat64** |  | [optional] 
**TaxAgencyXref** | Pointer to **string** |  | [optional] 
**TaxNote** | Pointer to **string** |  | [optional] 
**PurchaseHeaderTaxGroup** | Pointer to **string** |  | [optional] 

## Methods

### NewGLExportPurchaseTransactionDetailTax

`func NewGLExportPurchaseTransactionDetailTax() *GLExportPurchaseTransactionDetailTax`

NewGLExportPurchaseTransactionDetailTax instantiates a new GLExportPurchaseTransactionDetailTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportPurchaseTransactionDetailTaxWithDefaults

`func NewGLExportPurchaseTransactionDetailTaxWithDefaults() *GLExportPurchaseTransactionDetailTax`

NewGLExportPurchaseTransactionDetailTaxWithDefaults instantiates a new GLExportPurchaseTransactionDetailTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLExportPurchaseTransactionDetailTax) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLExportPurchaseTransactionDetailTax) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLExportPurchaseTransactionDetailTax) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GLExportPurchaseTransactionDetailTax) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GLExportPurchaseTransactionDetailTax) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GLExportPurchaseTransactionDetailTax) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDocumentDate

`func (o *GLExportPurchaseTransactionDetailTax) GetDocumentDate() string`

GetDocumentDate returns the DocumentDate field if non-nil, zero value otherwise.

### GetDocumentDateOk

`func (o *GLExportPurchaseTransactionDetailTax) GetDocumentDateOk() (*string, bool)`

GetDocumentDateOk returns a tuple with the DocumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDate

`func (o *GLExportPurchaseTransactionDetailTax) SetDocumentDate(v string)`

SetDocumentDate sets DocumentDate field to given value.

### HasDocumentDate

`func (o *GLExportPurchaseTransactionDetailTax) HasDocumentDate() bool`

HasDocumentDate returns a boolean if a field has been set.

### GetAccountNumber

`func (o *GLExportPurchaseTransactionDetailTax) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GLExportPurchaseTransactionDetailTax) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GLExportPurchaseTransactionDetailTax) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GLExportPurchaseTransactionDetailTax) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetGlClass

`func (o *GLExportPurchaseTransactionDetailTax) GetGlClass() string`

GetGlClass returns the GlClass field if non-nil, zero value otherwise.

### GetGlClassOk

`func (o *GLExportPurchaseTransactionDetailTax) GetGlClassOk() (*string, bool)`

GetGlClassOk returns a tuple with the GlClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlClass

`func (o *GLExportPurchaseTransactionDetailTax) SetGlClass(v string)`

SetGlClass sets GlClass field to given value.

### HasGlClass

`func (o *GLExportPurchaseTransactionDetailTax) HasGlClass() bool`

HasGlClass returns a boolean if a field has been set.

### GetCost

`func (o *GLExportPurchaseTransactionDetailTax) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GLExportPurchaseTransactionDetailTax) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GLExportPurchaseTransactionDetailTax) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GLExportPurchaseTransactionDetailTax) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *GLExportPurchaseTransactionDetailTax) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *GLExportPurchaseTransactionDetailTax) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetSalesCode

`func (o *GLExportPurchaseTransactionDetailTax) GetSalesCode() string`

GetSalesCode returns the SalesCode field if non-nil, zero value otherwise.

### GetSalesCodeOk

`func (o *GLExportPurchaseTransactionDetailTax) GetSalesCodeOk() (*string, bool)`

GetSalesCodeOk returns a tuple with the SalesCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesCode

`func (o *GLExportPurchaseTransactionDetailTax) SetSalesCode(v string)`

SetSalesCode sets SalesCode field to given value.

### HasSalesCode

`func (o *GLExportPurchaseTransactionDetailTax) HasSalesCode() bool`

HasSalesCode returns a boolean if a field has been set.

### GetGlTypeId

`func (o *GLExportPurchaseTransactionDetailTax) GetGlTypeId() string`

GetGlTypeId returns the GlTypeId field if non-nil, zero value otherwise.

### GetGlTypeIdOk

`func (o *GLExportPurchaseTransactionDetailTax) GetGlTypeIdOk() (*string, bool)`

GetGlTypeIdOk returns a tuple with the GlTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlTypeId

`func (o *GLExportPurchaseTransactionDetailTax) SetGlTypeId(v string)`

SetGlTypeId sets GlTypeId field to given value.

### HasGlTypeId

`func (o *GLExportPurchaseTransactionDetailTax) HasGlTypeId() bool`

HasGlTypeId returns a boolean if a field has been set.

### GetGlItemId

`func (o *GLExportPurchaseTransactionDetailTax) GetGlItemId() string`

GetGlItemId returns the GlItemId field if non-nil, zero value otherwise.

### GetGlItemIdOk

`func (o *GLExportPurchaseTransactionDetailTax) GetGlItemIdOk() (*string, bool)`

GetGlItemIdOk returns a tuple with the GlItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlItemId

`func (o *GLExportPurchaseTransactionDetailTax) SetGlItemId(v string)`

SetGlItemId sets GlItemId field to given value.

### HasGlItemId

`func (o *GLExportPurchaseTransactionDetailTax) HasGlItemId() bool`

HasGlItemId returns a boolean if a field has been set.

### GetMemo

`func (o *GLExportPurchaseTransactionDetailTax) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GLExportPurchaseTransactionDetailTax) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GLExportPurchaseTransactionDetailTax) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GLExportPurchaseTransactionDetailTax) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetVendorNumber

`func (o *GLExportPurchaseTransactionDetailTax) GetVendorNumber() string`

GetVendorNumber returns the VendorNumber field if non-nil, zero value otherwise.

### GetVendorNumberOk

`func (o *GLExportPurchaseTransactionDetailTax) GetVendorNumberOk() (*string, bool)`

GetVendorNumberOk returns a tuple with the VendorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumber

`func (o *GLExportPurchaseTransactionDetailTax) SetVendorNumber(v string)`

SetVendorNumber sets VendorNumber field to given value.

### HasVendorNumber

`func (o *GLExportPurchaseTransactionDetailTax) HasVendorNumber() bool`

HasVendorNumber returns a boolean if a field has been set.

### GetVendorAccountNumber

`func (o *GLExportPurchaseTransactionDetailTax) GetVendorAccountNumber() string`

GetVendorAccountNumber returns the VendorAccountNumber field if non-nil, zero value otherwise.

### GetVendorAccountNumberOk

`func (o *GLExportPurchaseTransactionDetailTax) GetVendorAccountNumberOk() (*string, bool)`

GetVendorAccountNumberOk returns a tuple with the VendorAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAccountNumber

`func (o *GLExportPurchaseTransactionDetailTax) SetVendorAccountNumber(v string)`

SetVendorAccountNumber sets VendorAccountNumber field to given value.

### HasVendorAccountNumber

`func (o *GLExportPurchaseTransactionDetailTax) HasVendorAccountNumber() bool`

HasVendorAccountNumber returns a boolean if a field has been set.

### GetCostAccountNumber

`func (o *GLExportPurchaseTransactionDetailTax) GetCostAccountNumber() string`

GetCostAccountNumber returns the CostAccountNumber field if non-nil, zero value otherwise.

### GetCostAccountNumberOk

`func (o *GLExportPurchaseTransactionDetailTax) GetCostAccountNumberOk() (*string, bool)`

GetCostAccountNumberOk returns a tuple with the CostAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAccountNumber

`func (o *GLExportPurchaseTransactionDetailTax) SetCostAccountNumber(v string)`

SetCostAccountNumber sets CostAccountNumber field to given value.

### HasCostAccountNumber

`func (o *GLExportPurchaseTransactionDetailTax) HasCostAccountNumber() bool`

HasCostAccountNumber returns a boolean if a field has been set.

### GetInventoryAccountNumber

`func (o *GLExportPurchaseTransactionDetailTax) GetInventoryAccountNumber() string`

GetInventoryAccountNumber returns the InventoryAccountNumber field if non-nil, zero value otherwise.

### GetInventoryAccountNumberOk

`func (o *GLExportPurchaseTransactionDetailTax) GetInventoryAccountNumberOk() (*string, bool)`

GetInventoryAccountNumberOk returns a tuple with the InventoryAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryAccountNumber

`func (o *GLExportPurchaseTransactionDetailTax) SetInventoryAccountNumber(v string)`

SetInventoryAccountNumber sets InventoryAccountNumber field to given value.

### HasInventoryAccountNumber

`func (o *GLExportPurchaseTransactionDetailTax) HasInventoryAccountNumber() bool`

HasInventoryAccountNumber returns a boolean if a field has been set.

### GetItemTypeXref

`func (o *GLExportPurchaseTransactionDetailTax) GetItemTypeXref() string`

GetItemTypeXref returns the ItemTypeXref field if non-nil, zero value otherwise.

### GetItemTypeXrefOk

`func (o *GLExportPurchaseTransactionDetailTax) GetItemTypeXrefOk() (*string, bool)`

GetItemTypeXrefOk returns a tuple with the ItemTypeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTypeXref

`func (o *GLExportPurchaseTransactionDetailTax) SetItemTypeXref(v string)`

SetItemTypeXref sets ItemTypeXref field to given value.

### HasItemTypeXref

`func (o *GLExportPurchaseTransactionDetailTax) HasItemTypeXref() bool`

HasItemTypeXref returns a boolean if a field has been set.

### GetInventoryXref

`func (o *GLExportPurchaseTransactionDetailTax) GetInventoryXref() string`

GetInventoryXref returns the InventoryXref field if non-nil, zero value otherwise.

### GetInventoryXrefOk

`func (o *GLExportPurchaseTransactionDetailTax) GetInventoryXrefOk() (*string, bool)`

GetInventoryXrefOk returns a tuple with the InventoryXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryXref

`func (o *GLExportPurchaseTransactionDetailTax) SetInventoryXref(v string)`

SetInventoryXref sets InventoryXref field to given value.

### HasInventoryXref

`func (o *GLExportPurchaseTransactionDetailTax) HasInventoryXref() bool`

HasInventoryXref returns a boolean if a field has been set.

### GetCogsXref

`func (o *GLExportPurchaseTransactionDetailTax) GetCogsXref() string`

GetCogsXref returns the CogsXref field if non-nil, zero value otherwise.

### GetCogsXrefOk

`func (o *GLExportPurchaseTransactionDetailTax) GetCogsXrefOk() (*string, bool)`

GetCogsXrefOk returns a tuple with the CogsXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogsXref

`func (o *GLExportPurchaseTransactionDetailTax) SetCogsXref(v string)`

SetCogsXref sets CogsXref field to given value.

### HasCogsXref

`func (o *GLExportPurchaseTransactionDetailTax) HasCogsXref() bool`

HasCogsXref returns a boolean if a field has been set.

### GetUomScheduleXref

`func (o *GLExportPurchaseTransactionDetailTax) GetUomScheduleXref() string`

GetUomScheduleXref returns the UomScheduleXref field if non-nil, zero value otherwise.

### GetUomScheduleXrefOk

`func (o *GLExportPurchaseTransactionDetailTax) GetUomScheduleXrefOk() (*string, bool)`

GetUomScheduleXrefOk returns a tuple with the UomScheduleXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUomScheduleXref

`func (o *GLExportPurchaseTransactionDetailTax) SetUomScheduleXref(v string)`

SetUomScheduleXref sets UomScheduleXref field to given value.

### HasUomScheduleXref

`func (o *GLExportPurchaseTransactionDetailTax) HasUomScheduleXref() bool`

HasUomScheduleXref returns a boolean if a field has been set.

### GetPriceLevelXref

`func (o *GLExportPurchaseTransactionDetailTax) GetPriceLevelXref() string`

GetPriceLevelXref returns the PriceLevelXref field if non-nil, zero value otherwise.

### GetPriceLevelXrefOk

`func (o *GLExportPurchaseTransactionDetailTax) GetPriceLevelXrefOk() (*string, bool)`

GetPriceLevelXrefOk returns a tuple with the PriceLevelXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevelXref

`func (o *GLExportPurchaseTransactionDetailTax) SetPriceLevelXref(v string)`

SetPriceLevelXref sets PriceLevelXref field to given value.

### HasPriceLevelXref

`func (o *GLExportPurchaseTransactionDetailTax) HasPriceLevelXref() bool`

HasPriceLevelXref returns a boolean if a field has been set.

### GetLocationXref

`func (o *GLExportPurchaseTransactionDetailTax) GetLocationXref() string`

GetLocationXref returns the LocationXref field if non-nil, zero value otherwise.

### GetLocationXrefOk

`func (o *GLExportPurchaseTransactionDetailTax) GetLocationXrefOk() (*string, bool)`

GetLocationXrefOk returns a tuple with the LocationXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationXref

`func (o *GLExportPurchaseTransactionDetailTax) SetLocationXref(v string)`

SetLocationXref sets LocationXref field to given value.

### HasLocationXref

`func (o *GLExportPurchaseTransactionDetailTax) HasLocationXref() bool`

HasLocationXref returns a boolean if a field has been set.

### GetItem

`func (o *GLExportPurchaseTransactionDetailTax) GetItem() IvItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *GLExportPurchaseTransactionDetailTax) GetItemOk() (*IvItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *GLExportPurchaseTransactionDetailTax) SetItem(v IvItemReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *GLExportPurchaseTransactionDetailTax) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetTaxableFlag

`func (o *GLExportPurchaseTransactionDetailTax) GetTaxableFlag() bool`

GetTaxableFlag returns the TaxableFlag field if non-nil, zero value otherwise.

### GetTaxableFlagOk

`func (o *GLExportPurchaseTransactionDetailTax) GetTaxableFlagOk() (*bool, bool)`

GetTaxableFlagOk returns a tuple with the TaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableFlag

`func (o *GLExportPurchaseTransactionDetailTax) SetTaxableFlag(v bool)`

SetTaxableFlag sets TaxableFlag field to given value.

### HasTaxableFlag

`func (o *GLExportPurchaseTransactionDetailTax) HasTaxableFlag() bool`

HasTaxableFlag returns a boolean if a field has been set.

### SetTaxableFlagNil

`func (o *GLExportPurchaseTransactionDetailTax) SetTaxableFlagNil(b bool)`

 SetTaxableFlagNil sets the value for TaxableFlag to be an explicit nil

### UnsetTaxableFlag
`func (o *GLExportPurchaseTransactionDetailTax) UnsetTaxableFlag()`

UnsetTaxableFlag ensures that no value is present for TaxableFlag, not even an explicit nil
### GetSalesDescription

`func (o *GLExportPurchaseTransactionDetailTax) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *GLExportPurchaseTransactionDetailTax) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *GLExportPurchaseTransactionDetailTax) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.

### HasSalesDescription

`func (o *GLExportPurchaseTransactionDetailTax) HasSalesDescription() bool`

HasSalesDescription returns a boolean if a field has been set.

### GetItemDescription

`func (o *GLExportPurchaseTransactionDetailTax) GetItemDescription() string`

GetItemDescription returns the ItemDescription field if non-nil, zero value otherwise.

### GetItemDescriptionOk

`func (o *GLExportPurchaseTransactionDetailTax) GetItemDescriptionOk() (*string, bool)`

GetItemDescriptionOk returns a tuple with the ItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDescription

`func (o *GLExportPurchaseTransactionDetailTax) SetItemDescription(v string)`

SetItemDescription sets ItemDescription field to given value.

### HasItemDescription

`func (o *GLExportPurchaseTransactionDetailTax) HasItemDescription() bool`

HasItemDescription returns a boolean if a field has been set.

### GetItemPrice

`func (o *GLExportPurchaseTransactionDetailTax) GetItemPrice() float64`

GetItemPrice returns the ItemPrice field if non-nil, zero value otherwise.

### GetItemPriceOk

`func (o *GLExportPurchaseTransactionDetailTax) GetItemPriceOk() (*float64, bool)`

GetItemPriceOk returns a tuple with the ItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrice

`func (o *GLExportPurchaseTransactionDetailTax) SetItemPrice(v float64)`

SetItemPrice sets ItemPrice field to given value.

### HasItemPrice

`func (o *GLExportPurchaseTransactionDetailTax) HasItemPrice() bool`

HasItemPrice returns a boolean if a field has been set.

### SetItemPriceNil

`func (o *GLExportPurchaseTransactionDetailTax) SetItemPriceNil(b bool)`

 SetItemPriceNil sets the value for ItemPrice to be an explicit nil

### UnsetItemPrice
`func (o *GLExportPurchaseTransactionDetailTax) UnsetItemPrice()`

UnsetItemPrice ensures that no value is present for ItemPrice, not even an explicit nil
### GetItemCost

`func (o *GLExportPurchaseTransactionDetailTax) GetItemCost() float64`

GetItemCost returns the ItemCost field if non-nil, zero value otherwise.

### GetItemCostOk

`func (o *GLExportPurchaseTransactionDetailTax) GetItemCostOk() (*float64, bool)`

GetItemCostOk returns a tuple with the ItemCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCost

`func (o *GLExportPurchaseTransactionDetailTax) SetItemCost(v float64)`

SetItemCost sets ItemCost field to given value.

### HasItemCost

`func (o *GLExportPurchaseTransactionDetailTax) HasItemCost() bool`

HasItemCost returns a boolean if a field has been set.

### SetItemCostNil

`func (o *GLExportPurchaseTransactionDetailTax) SetItemCostNil(b bool)`

 SetItemCostNil sets the value for ItemCost to be an explicit nil

### UnsetItemCost
`func (o *GLExportPurchaseTransactionDetailTax) UnsetItemCost()`

UnsetItemCost ensures that no value is present for ItemCost, not even an explicit nil
### GetUnitOfMeasure

`func (o *GLExportPurchaseTransactionDetailTax) GetUnitOfMeasure() UnitOfMeasureReference`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *GLExportPurchaseTransactionDetailTax) GetUnitOfMeasureOk() (*UnitOfMeasureReference, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *GLExportPurchaseTransactionDetailTax) SetUnitOfMeasure(v UnitOfMeasureReference)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *GLExportPurchaseTransactionDetailTax) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetQuantity

`func (o *GLExportPurchaseTransactionDetailTax) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GLExportPurchaseTransactionDetailTax) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GLExportPurchaseTransactionDetailTax) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GLExportPurchaseTransactionDetailTax) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *GLExportPurchaseTransactionDetailTax) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *GLExportPurchaseTransactionDetailTax) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetTotal

`func (o *GLExportPurchaseTransactionDetailTax) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GLExportPurchaseTransactionDetailTax) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GLExportPurchaseTransactionDetailTax) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GLExportPurchaseTransactionDetailTax) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GLExportPurchaseTransactionDetailTax) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GLExportPurchaseTransactionDetailTax) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetCurrency

`func (o *GLExportPurchaseTransactionDetailTax) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GLExportPurchaseTransactionDetailTax) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GLExportPurchaseTransactionDetailTax) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GLExportPurchaseTransactionDetailTax) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSerializedFlag

`func (o *GLExportPurchaseTransactionDetailTax) GetSerializedFlag() bool`

GetSerializedFlag returns the SerializedFlag field if non-nil, zero value otherwise.

### GetSerializedFlagOk

`func (o *GLExportPurchaseTransactionDetailTax) GetSerializedFlagOk() (*bool, bool)`

GetSerializedFlagOk returns a tuple with the SerializedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedFlag

`func (o *GLExportPurchaseTransactionDetailTax) SetSerializedFlag(v bool)`

SetSerializedFlag sets SerializedFlag field to given value.

### HasSerializedFlag

`func (o *GLExportPurchaseTransactionDetailTax) HasSerializedFlag() bool`

HasSerializedFlag returns a boolean if a field has been set.

### SetSerializedFlagNil

`func (o *GLExportPurchaseTransactionDetailTax) SetSerializedFlagNil(b bool)`

 SetSerializedFlagNil sets the value for SerializedFlag to be an explicit nil

### UnsetSerializedFlag
`func (o *GLExportPurchaseTransactionDetailTax) UnsetSerializedFlag()`

UnsetSerializedFlag ensures that no value is present for SerializedFlag, not even an explicit nil
### GetSerialNumbers

`func (o *GLExportPurchaseTransactionDetailTax) GetSerialNumbers() string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *GLExportPurchaseTransactionDetailTax) GetSerialNumbersOk() (*string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *GLExportPurchaseTransactionDetailTax) SetSerialNumbers(v string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *GLExportPurchaseTransactionDetailTax) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.

### GetDropShippedFlag

`func (o *GLExportPurchaseTransactionDetailTax) GetDropShippedFlag() bool`

GetDropShippedFlag returns the DropShippedFlag field if non-nil, zero value otherwise.

### GetDropShippedFlagOk

`func (o *GLExportPurchaseTransactionDetailTax) GetDropShippedFlagOk() (*bool, bool)`

GetDropShippedFlagOk returns a tuple with the DropShippedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropShippedFlag

`func (o *GLExportPurchaseTransactionDetailTax) SetDropShippedFlag(v bool)`

SetDropShippedFlag sets DropShippedFlag field to given value.

### HasDropShippedFlag

`func (o *GLExportPurchaseTransactionDetailTax) HasDropShippedFlag() bool`

HasDropShippedFlag returns a boolean if a field has been set.

### SetDropShippedFlagNil

`func (o *GLExportPurchaseTransactionDetailTax) SetDropShippedFlagNil(b bool)`

 SetDropShippedFlagNil sets the value for DropShippedFlag to be an explicit nil

### UnsetDropShippedFlag
`func (o *GLExportPurchaseTransactionDetailTax) UnsetDropShippedFlag()`

UnsetDropShippedFlag ensures that no value is present for DropShippedFlag, not even an explicit nil
### GetLineNumber

`func (o *GLExportPurchaseTransactionDetailTax) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *GLExportPurchaseTransactionDetailTax) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *GLExportPurchaseTransactionDetailTax) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *GLExportPurchaseTransactionDetailTax) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetWarehouseSite

`func (o *GLExportPurchaseTransactionDetailTax) GetWarehouseSite() SiteReference`

GetWarehouseSite returns the WarehouseSite field if non-nil, zero value otherwise.

### GetWarehouseSiteOk

`func (o *GLExportPurchaseTransactionDetailTax) GetWarehouseSiteOk() (*SiteReference, bool)`

GetWarehouseSiteOk returns a tuple with the WarehouseSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseSite

`func (o *GLExportPurchaseTransactionDetailTax) SetWarehouseSite(v SiteReference)`

SetWarehouseSite sets WarehouseSite field to given value.

### HasWarehouseSite

`func (o *GLExportPurchaseTransactionDetailTax) HasWarehouseSite() bool`

HasWarehouseSite returns a boolean if a field has been set.

### GetWarehouseBin

`func (o *GLExportPurchaseTransactionDetailTax) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *GLExportPurchaseTransactionDetailTax) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *GLExportPurchaseTransactionDetailTax) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.

### HasWarehouseBin

`func (o *GLExportPurchaseTransactionDetailTax) HasWarehouseBin() bool`

HasWarehouseBin returns a boolean if a field has been set.

### GetShipmentMethod

`func (o *GLExportPurchaseTransactionDetailTax) GetShipmentMethod() ShipmentMethodReference`

GetShipmentMethod returns the ShipmentMethod field if non-nil, zero value otherwise.

### GetShipmentMethodOk

`func (o *GLExportPurchaseTransactionDetailTax) GetShipmentMethodOk() (*ShipmentMethodReference, bool)`

GetShipmentMethodOk returns a tuple with the ShipmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentMethod

`func (o *GLExportPurchaseTransactionDetailTax) SetShipmentMethod(v ShipmentMethodReference)`

SetShipmentMethod sets ShipmentMethod field to given value.

### HasShipmentMethod

`func (o *GLExportPurchaseTransactionDetailTax) HasShipmentMethod() bool`

HasShipmentMethod returns a boolean if a field has been set.

### GetSubCategory

`func (o *GLExportPurchaseTransactionDetailTax) GetSubCategory() ProductSubCategoryReference`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *GLExportPurchaseTransactionDetailTax) GetSubCategoryOk() (*ProductSubCategoryReference, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *GLExportPurchaseTransactionDetailTax) SetSubCategory(v ProductSubCategoryReference)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *GLExportPurchaseTransactionDetailTax) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetTaxCode

`func (o *GLExportPurchaseTransactionDetailTax) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *GLExportPurchaseTransactionDetailTax) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *GLExportPurchaseTransactionDetailTax) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *GLExportPurchaseTransactionDetailTax) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxRate

`func (o *GLExportPurchaseTransactionDetailTax) GetTaxRate() float64`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *GLExportPurchaseTransactionDetailTax) GetTaxRateOk() (*float64, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *GLExportPurchaseTransactionDetailTax) SetTaxRate(v float64)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *GLExportPurchaseTransactionDetailTax) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *GLExportPurchaseTransactionDetailTax) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *GLExportPurchaseTransactionDetailTax) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetTaxRatePercent

`func (o *GLExportPurchaseTransactionDetailTax) GetTaxRatePercent() float64`

GetTaxRatePercent returns the TaxRatePercent field if non-nil, zero value otherwise.

### GetTaxRatePercentOk

`func (o *GLExportPurchaseTransactionDetailTax) GetTaxRatePercentOk() (*float64, bool)`

GetTaxRatePercentOk returns a tuple with the TaxRatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRatePercent

`func (o *GLExportPurchaseTransactionDetailTax) SetTaxRatePercent(v float64)`

SetTaxRatePercent sets TaxRatePercent field to given value.

### HasTaxRatePercent

`func (o *GLExportPurchaseTransactionDetailTax) HasTaxRatePercent() bool`

HasTaxRatePercent returns a boolean if a field has been set.

### SetTaxRatePercentNil

`func (o *GLExportPurchaseTransactionDetailTax) SetTaxRatePercentNil(b bool)`

 SetTaxRatePercentNil sets the value for TaxRatePercent to be an explicit nil

### UnsetTaxRatePercent
`func (o *GLExportPurchaseTransactionDetailTax) UnsetTaxRatePercent()`

UnsetTaxRatePercent ensures that no value is present for TaxRatePercent, not even an explicit nil
### GetTaxAgencyXref

`func (o *GLExportPurchaseTransactionDetailTax) GetTaxAgencyXref() string`

GetTaxAgencyXref returns the TaxAgencyXref field if non-nil, zero value otherwise.

### GetTaxAgencyXrefOk

`func (o *GLExportPurchaseTransactionDetailTax) GetTaxAgencyXrefOk() (*string, bool)`

GetTaxAgencyXrefOk returns a tuple with the TaxAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAgencyXref

`func (o *GLExportPurchaseTransactionDetailTax) SetTaxAgencyXref(v string)`

SetTaxAgencyXref sets TaxAgencyXref field to given value.

### HasTaxAgencyXref

`func (o *GLExportPurchaseTransactionDetailTax) HasTaxAgencyXref() bool`

HasTaxAgencyXref returns a boolean if a field has been set.

### GetTaxNote

`func (o *GLExportPurchaseTransactionDetailTax) GetTaxNote() string`

GetTaxNote returns the TaxNote field if non-nil, zero value otherwise.

### GetTaxNoteOk

`func (o *GLExportPurchaseTransactionDetailTax) GetTaxNoteOk() (*string, bool)`

GetTaxNoteOk returns a tuple with the TaxNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNote

`func (o *GLExportPurchaseTransactionDetailTax) SetTaxNote(v string)`

SetTaxNote sets TaxNote field to given value.

### HasTaxNote

`func (o *GLExportPurchaseTransactionDetailTax) HasTaxNote() bool`

HasTaxNote returns a boolean if a field has been set.

### GetPurchaseHeaderTaxGroup

`func (o *GLExportPurchaseTransactionDetailTax) GetPurchaseHeaderTaxGroup() string`

GetPurchaseHeaderTaxGroup returns the PurchaseHeaderTaxGroup field if non-nil, zero value otherwise.

### GetPurchaseHeaderTaxGroupOk

`func (o *GLExportPurchaseTransactionDetailTax) GetPurchaseHeaderTaxGroupOk() (*string, bool)`

GetPurchaseHeaderTaxGroupOk returns a tuple with the PurchaseHeaderTaxGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseHeaderTaxGroup

`func (o *GLExportPurchaseTransactionDetailTax) SetPurchaseHeaderTaxGroup(v string)`

SetPurchaseHeaderTaxGroup sets PurchaseHeaderTaxGroup field to given value.

### HasPurchaseHeaderTaxGroup

`func (o *GLExportPurchaseTransactionDetailTax) HasPurchaseHeaderTaxGroup() bool`

HasPurchaseHeaderTaxGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


