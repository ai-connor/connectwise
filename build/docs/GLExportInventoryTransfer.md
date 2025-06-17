# GLExportInventoryTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 
**DocumentDate** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**GlClass** | Pointer to **string** |  | [optional] 
**GlTypeId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SalesCode** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**CostAccountNumber** | Pointer to **string** |  | [optional] 
**InventoryAccountNumber** | Pointer to **string** |  | [optional] 
**TransferId** | Pointer to **NullableInt32** |  | [optional] 
**Item** | Pointer to [**IvItemReference**](IvItemReference.md) |  | [optional] 
**GlItemId** | Pointer to **string** |  | [optional] 
**SalesDescription** | Pointer to **string** |  | [optional] 
**ItemDescription** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**ItemPrice** | Pointer to **NullableFloat64** |  | [optional] 
**Taxable** | Pointer to **NullableBool** |  | [optional] 
**UnitOfMeasure** | Pointer to [**UnitOfMeasureReference**](UnitOfMeasureReference.md) |  | [optional] 
**Quantity** | Pointer to **NullableFloat64** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**SubCategory** | Pointer to [**ProductSubCategoryReference**](ProductSubCategoryReference.md) |  | [optional] 
**SerializedFlag** | Pointer to **NullableBool** |  | [optional] 
**SerialNumbers** | Pointer to **string** |  | [optional] 
**Bin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**Warehouse** | Pointer to [**WarehouseReference**](WarehouseReference.md) |  | [optional] 
**TransferFromBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**TransferFromLocationXref** | Pointer to **string** |  | [optional] 
**TransferToBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**TransferToLocationXref** | Pointer to **string** |  | [optional] 
**LocationXref** | Pointer to **string** |  | [optional] 
**PriceLevelXref** | Pointer to **string** |  | [optional] 
**UomScheduleXref** | Pointer to **string** |  | [optional] 
**ItemTypeXref** | Pointer to **string** |  | [optional] 
**InventoryXref** | Pointer to **string** |  | [optional] 
**CogsXref** | Pointer to **string** |  | [optional] 
**TaxNote** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to [**GLExportInventoryTransferOffset**](GLExportInventoryTransferOffset.md) |  | [optional] 

## Methods

### NewGLExportInventoryTransfer

`func NewGLExportInventoryTransfer() *GLExportInventoryTransfer`

NewGLExportInventoryTransfer instantiates a new GLExportInventoryTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportInventoryTransferWithDefaults

`func NewGLExportInventoryTransferWithDefaults() *GLExportInventoryTransfer`

NewGLExportInventoryTransferWithDefaults instantiates a new GLExportInventoryTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLExportInventoryTransfer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLExportInventoryTransfer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLExportInventoryTransfer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GLExportInventoryTransfer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocumentType

`func (o *GLExportInventoryTransfer) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *GLExportInventoryTransfer) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *GLExportInventoryTransfer) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *GLExportInventoryTransfer) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetDocumentDate

`func (o *GLExportInventoryTransfer) GetDocumentDate() string`

GetDocumentDate returns the DocumentDate field if non-nil, zero value otherwise.

### GetDocumentDateOk

`func (o *GLExportInventoryTransfer) GetDocumentDateOk() (*string, bool)`

GetDocumentDateOk returns a tuple with the DocumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDate

`func (o *GLExportInventoryTransfer) SetDocumentDate(v string)`

SetDocumentDate sets DocumentDate field to given value.

### HasDocumentDate

`func (o *GLExportInventoryTransfer) HasDocumentDate() bool`

HasDocumentDate returns a boolean if a field has been set.

### GetAccountNumber

`func (o *GLExportInventoryTransfer) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GLExportInventoryTransfer) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GLExportInventoryTransfer) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GLExportInventoryTransfer) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetGlClass

`func (o *GLExportInventoryTransfer) GetGlClass() string`

GetGlClass returns the GlClass field if non-nil, zero value otherwise.

### GetGlClassOk

`func (o *GLExportInventoryTransfer) GetGlClassOk() (*string, bool)`

GetGlClassOk returns a tuple with the GlClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlClass

`func (o *GLExportInventoryTransfer) SetGlClass(v string)`

SetGlClass sets GlClass field to given value.

### HasGlClass

`func (o *GLExportInventoryTransfer) HasGlClass() bool`

HasGlClass returns a boolean if a field has been set.

### GetGlTypeId

`func (o *GLExportInventoryTransfer) GetGlTypeId() string`

GetGlTypeId returns the GlTypeId field if non-nil, zero value otherwise.

### GetGlTypeIdOk

`func (o *GLExportInventoryTransfer) GetGlTypeIdOk() (*string, bool)`

GetGlTypeIdOk returns a tuple with the GlTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlTypeId

`func (o *GLExportInventoryTransfer) SetGlTypeId(v string)`

SetGlTypeId sets GlTypeId field to given value.

### HasGlTypeId

`func (o *GLExportInventoryTransfer) HasGlTypeId() bool`

HasGlTypeId returns a boolean if a field has been set.

### GetDescription

`func (o *GLExportInventoryTransfer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GLExportInventoryTransfer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GLExportInventoryTransfer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GLExportInventoryTransfer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSalesCode

`func (o *GLExportInventoryTransfer) GetSalesCode() string`

GetSalesCode returns the SalesCode field if non-nil, zero value otherwise.

### GetSalesCodeOk

`func (o *GLExportInventoryTransfer) GetSalesCodeOk() (*string, bool)`

GetSalesCodeOk returns a tuple with the SalesCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesCode

`func (o *GLExportInventoryTransfer) SetSalesCode(v string)`

SetSalesCode sets SalesCode field to given value.

### HasSalesCode

`func (o *GLExportInventoryTransfer) HasSalesCode() bool`

HasSalesCode returns a boolean if a field has been set.

### GetMemo

`func (o *GLExportInventoryTransfer) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GLExportInventoryTransfer) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GLExportInventoryTransfer) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GLExportInventoryTransfer) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetCostAccountNumber

`func (o *GLExportInventoryTransfer) GetCostAccountNumber() string`

GetCostAccountNumber returns the CostAccountNumber field if non-nil, zero value otherwise.

### GetCostAccountNumberOk

`func (o *GLExportInventoryTransfer) GetCostAccountNumberOk() (*string, bool)`

GetCostAccountNumberOk returns a tuple with the CostAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAccountNumber

`func (o *GLExportInventoryTransfer) SetCostAccountNumber(v string)`

SetCostAccountNumber sets CostAccountNumber field to given value.

### HasCostAccountNumber

`func (o *GLExportInventoryTransfer) HasCostAccountNumber() bool`

HasCostAccountNumber returns a boolean if a field has been set.

### GetInventoryAccountNumber

`func (o *GLExportInventoryTransfer) GetInventoryAccountNumber() string`

GetInventoryAccountNumber returns the InventoryAccountNumber field if non-nil, zero value otherwise.

### GetInventoryAccountNumberOk

`func (o *GLExportInventoryTransfer) GetInventoryAccountNumberOk() (*string, bool)`

GetInventoryAccountNumberOk returns a tuple with the InventoryAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryAccountNumber

`func (o *GLExportInventoryTransfer) SetInventoryAccountNumber(v string)`

SetInventoryAccountNumber sets InventoryAccountNumber field to given value.

### HasInventoryAccountNumber

`func (o *GLExportInventoryTransfer) HasInventoryAccountNumber() bool`

HasInventoryAccountNumber returns a boolean if a field has been set.

### GetTransferId

`func (o *GLExportInventoryTransfer) GetTransferId() int32`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *GLExportInventoryTransfer) GetTransferIdOk() (*int32, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *GLExportInventoryTransfer) SetTransferId(v int32)`

SetTransferId sets TransferId field to given value.

### HasTransferId

`func (o *GLExportInventoryTransfer) HasTransferId() bool`

HasTransferId returns a boolean if a field has been set.

### SetTransferIdNil

`func (o *GLExportInventoryTransfer) SetTransferIdNil(b bool)`

 SetTransferIdNil sets the value for TransferId to be an explicit nil

### UnsetTransferId
`func (o *GLExportInventoryTransfer) UnsetTransferId()`

UnsetTransferId ensures that no value is present for TransferId, not even an explicit nil
### GetItem

`func (o *GLExportInventoryTransfer) GetItem() IvItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *GLExportInventoryTransfer) GetItemOk() (*IvItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *GLExportInventoryTransfer) SetItem(v IvItemReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *GLExportInventoryTransfer) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetGlItemId

`func (o *GLExportInventoryTransfer) GetGlItemId() string`

GetGlItemId returns the GlItemId field if non-nil, zero value otherwise.

### GetGlItemIdOk

`func (o *GLExportInventoryTransfer) GetGlItemIdOk() (*string, bool)`

GetGlItemIdOk returns a tuple with the GlItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlItemId

`func (o *GLExportInventoryTransfer) SetGlItemId(v string)`

SetGlItemId sets GlItemId field to given value.

### HasGlItemId

`func (o *GLExportInventoryTransfer) HasGlItemId() bool`

HasGlItemId returns a boolean if a field has been set.

### GetSalesDescription

`func (o *GLExportInventoryTransfer) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *GLExportInventoryTransfer) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *GLExportInventoryTransfer) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.

### HasSalesDescription

`func (o *GLExportInventoryTransfer) HasSalesDescription() bool`

HasSalesDescription returns a boolean if a field has been set.

### GetItemDescription

`func (o *GLExportInventoryTransfer) GetItemDescription() string`

GetItemDescription returns the ItemDescription field if non-nil, zero value otherwise.

### GetItemDescriptionOk

`func (o *GLExportInventoryTransfer) GetItemDescriptionOk() (*string, bool)`

GetItemDescriptionOk returns a tuple with the ItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDescription

`func (o *GLExportInventoryTransfer) SetItemDescription(v string)`

SetItemDescription sets ItemDescription field to given value.

### HasItemDescription

`func (o *GLExportInventoryTransfer) HasItemDescription() bool`

HasItemDescription returns a boolean if a field has been set.

### GetCurrency

`func (o *GLExportInventoryTransfer) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GLExportInventoryTransfer) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GLExportInventoryTransfer) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GLExportInventoryTransfer) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetItemPrice

`func (o *GLExportInventoryTransfer) GetItemPrice() float64`

GetItemPrice returns the ItemPrice field if non-nil, zero value otherwise.

### GetItemPriceOk

`func (o *GLExportInventoryTransfer) GetItemPriceOk() (*float64, bool)`

GetItemPriceOk returns a tuple with the ItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrice

`func (o *GLExportInventoryTransfer) SetItemPrice(v float64)`

SetItemPrice sets ItemPrice field to given value.

### HasItemPrice

`func (o *GLExportInventoryTransfer) HasItemPrice() bool`

HasItemPrice returns a boolean if a field has been set.

### SetItemPriceNil

`func (o *GLExportInventoryTransfer) SetItemPriceNil(b bool)`

 SetItemPriceNil sets the value for ItemPrice to be an explicit nil

### UnsetItemPrice
`func (o *GLExportInventoryTransfer) UnsetItemPrice()`

UnsetItemPrice ensures that no value is present for ItemPrice, not even an explicit nil
### GetTaxable

`func (o *GLExportInventoryTransfer) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *GLExportInventoryTransfer) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *GLExportInventoryTransfer) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *GLExportInventoryTransfer) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### SetTaxableNil

`func (o *GLExportInventoryTransfer) SetTaxableNil(b bool)`

 SetTaxableNil sets the value for Taxable to be an explicit nil

### UnsetTaxable
`func (o *GLExportInventoryTransfer) UnsetTaxable()`

UnsetTaxable ensures that no value is present for Taxable, not even an explicit nil
### GetUnitOfMeasure

`func (o *GLExportInventoryTransfer) GetUnitOfMeasure() UnitOfMeasureReference`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *GLExportInventoryTransfer) GetUnitOfMeasureOk() (*UnitOfMeasureReference, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *GLExportInventoryTransfer) SetUnitOfMeasure(v UnitOfMeasureReference)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *GLExportInventoryTransfer) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetQuantity

`func (o *GLExportInventoryTransfer) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GLExportInventoryTransfer) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GLExportInventoryTransfer) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GLExportInventoryTransfer) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *GLExportInventoryTransfer) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *GLExportInventoryTransfer) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetCost

`func (o *GLExportInventoryTransfer) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GLExportInventoryTransfer) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GLExportInventoryTransfer) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GLExportInventoryTransfer) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *GLExportInventoryTransfer) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *GLExportInventoryTransfer) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetTotal

`func (o *GLExportInventoryTransfer) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GLExportInventoryTransfer) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GLExportInventoryTransfer) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GLExportInventoryTransfer) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GLExportInventoryTransfer) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GLExportInventoryTransfer) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetSubCategory

`func (o *GLExportInventoryTransfer) GetSubCategory() ProductSubCategoryReference`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *GLExportInventoryTransfer) GetSubCategoryOk() (*ProductSubCategoryReference, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *GLExportInventoryTransfer) SetSubCategory(v ProductSubCategoryReference)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *GLExportInventoryTransfer) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetSerializedFlag

`func (o *GLExportInventoryTransfer) GetSerializedFlag() bool`

GetSerializedFlag returns the SerializedFlag field if non-nil, zero value otherwise.

### GetSerializedFlagOk

`func (o *GLExportInventoryTransfer) GetSerializedFlagOk() (*bool, bool)`

GetSerializedFlagOk returns a tuple with the SerializedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedFlag

`func (o *GLExportInventoryTransfer) SetSerializedFlag(v bool)`

SetSerializedFlag sets SerializedFlag field to given value.

### HasSerializedFlag

`func (o *GLExportInventoryTransfer) HasSerializedFlag() bool`

HasSerializedFlag returns a boolean if a field has been set.

### SetSerializedFlagNil

`func (o *GLExportInventoryTransfer) SetSerializedFlagNil(b bool)`

 SetSerializedFlagNil sets the value for SerializedFlag to be an explicit nil

### UnsetSerializedFlag
`func (o *GLExportInventoryTransfer) UnsetSerializedFlag()`

UnsetSerializedFlag ensures that no value is present for SerializedFlag, not even an explicit nil
### GetSerialNumbers

`func (o *GLExportInventoryTransfer) GetSerialNumbers() string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *GLExportInventoryTransfer) GetSerialNumbersOk() (*string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *GLExportInventoryTransfer) SetSerialNumbers(v string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *GLExportInventoryTransfer) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.

### GetBin

`func (o *GLExportInventoryTransfer) GetBin() WarehouseBinReference`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *GLExportInventoryTransfer) GetBinOk() (*WarehouseBinReference, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *GLExportInventoryTransfer) SetBin(v WarehouseBinReference)`

SetBin sets Bin field to given value.

### HasBin

`func (o *GLExportInventoryTransfer) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetWarehouse

`func (o *GLExportInventoryTransfer) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *GLExportInventoryTransfer) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *GLExportInventoryTransfer) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *GLExportInventoryTransfer) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetTransferFromBin

`func (o *GLExportInventoryTransfer) GetTransferFromBin() WarehouseBinReference`

GetTransferFromBin returns the TransferFromBin field if non-nil, zero value otherwise.

### GetTransferFromBinOk

`func (o *GLExportInventoryTransfer) GetTransferFromBinOk() (*WarehouseBinReference, bool)`

GetTransferFromBinOk returns a tuple with the TransferFromBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFromBin

`func (o *GLExportInventoryTransfer) SetTransferFromBin(v WarehouseBinReference)`

SetTransferFromBin sets TransferFromBin field to given value.

### HasTransferFromBin

`func (o *GLExportInventoryTransfer) HasTransferFromBin() bool`

HasTransferFromBin returns a boolean if a field has been set.

### GetTransferFromLocationXref

`func (o *GLExportInventoryTransfer) GetTransferFromLocationXref() string`

GetTransferFromLocationXref returns the TransferFromLocationXref field if non-nil, zero value otherwise.

### GetTransferFromLocationXrefOk

`func (o *GLExportInventoryTransfer) GetTransferFromLocationXrefOk() (*string, bool)`

GetTransferFromLocationXrefOk returns a tuple with the TransferFromLocationXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFromLocationXref

`func (o *GLExportInventoryTransfer) SetTransferFromLocationXref(v string)`

SetTransferFromLocationXref sets TransferFromLocationXref field to given value.

### HasTransferFromLocationXref

`func (o *GLExportInventoryTransfer) HasTransferFromLocationXref() bool`

HasTransferFromLocationXref returns a boolean if a field has been set.

### GetTransferToBin

`func (o *GLExportInventoryTransfer) GetTransferToBin() WarehouseBinReference`

GetTransferToBin returns the TransferToBin field if non-nil, zero value otherwise.

### GetTransferToBinOk

`func (o *GLExportInventoryTransfer) GetTransferToBinOk() (*WarehouseBinReference, bool)`

GetTransferToBinOk returns a tuple with the TransferToBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferToBin

`func (o *GLExportInventoryTransfer) SetTransferToBin(v WarehouseBinReference)`

SetTransferToBin sets TransferToBin field to given value.

### HasTransferToBin

`func (o *GLExportInventoryTransfer) HasTransferToBin() bool`

HasTransferToBin returns a boolean if a field has been set.

### GetTransferToLocationXref

`func (o *GLExportInventoryTransfer) GetTransferToLocationXref() string`

GetTransferToLocationXref returns the TransferToLocationXref field if non-nil, zero value otherwise.

### GetTransferToLocationXrefOk

`func (o *GLExportInventoryTransfer) GetTransferToLocationXrefOk() (*string, bool)`

GetTransferToLocationXrefOk returns a tuple with the TransferToLocationXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferToLocationXref

`func (o *GLExportInventoryTransfer) SetTransferToLocationXref(v string)`

SetTransferToLocationXref sets TransferToLocationXref field to given value.

### HasTransferToLocationXref

`func (o *GLExportInventoryTransfer) HasTransferToLocationXref() bool`

HasTransferToLocationXref returns a boolean if a field has been set.

### GetLocationXref

`func (o *GLExportInventoryTransfer) GetLocationXref() string`

GetLocationXref returns the LocationXref field if non-nil, zero value otherwise.

### GetLocationXrefOk

`func (o *GLExportInventoryTransfer) GetLocationXrefOk() (*string, bool)`

GetLocationXrefOk returns a tuple with the LocationXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationXref

`func (o *GLExportInventoryTransfer) SetLocationXref(v string)`

SetLocationXref sets LocationXref field to given value.

### HasLocationXref

`func (o *GLExportInventoryTransfer) HasLocationXref() bool`

HasLocationXref returns a boolean if a field has been set.

### GetPriceLevelXref

`func (o *GLExportInventoryTransfer) GetPriceLevelXref() string`

GetPriceLevelXref returns the PriceLevelXref field if non-nil, zero value otherwise.

### GetPriceLevelXrefOk

`func (o *GLExportInventoryTransfer) GetPriceLevelXrefOk() (*string, bool)`

GetPriceLevelXrefOk returns a tuple with the PriceLevelXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevelXref

`func (o *GLExportInventoryTransfer) SetPriceLevelXref(v string)`

SetPriceLevelXref sets PriceLevelXref field to given value.

### HasPriceLevelXref

`func (o *GLExportInventoryTransfer) HasPriceLevelXref() bool`

HasPriceLevelXref returns a boolean if a field has been set.

### GetUomScheduleXref

`func (o *GLExportInventoryTransfer) GetUomScheduleXref() string`

GetUomScheduleXref returns the UomScheduleXref field if non-nil, zero value otherwise.

### GetUomScheduleXrefOk

`func (o *GLExportInventoryTransfer) GetUomScheduleXrefOk() (*string, bool)`

GetUomScheduleXrefOk returns a tuple with the UomScheduleXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUomScheduleXref

`func (o *GLExportInventoryTransfer) SetUomScheduleXref(v string)`

SetUomScheduleXref sets UomScheduleXref field to given value.

### HasUomScheduleXref

`func (o *GLExportInventoryTransfer) HasUomScheduleXref() bool`

HasUomScheduleXref returns a boolean if a field has been set.

### GetItemTypeXref

`func (o *GLExportInventoryTransfer) GetItemTypeXref() string`

GetItemTypeXref returns the ItemTypeXref field if non-nil, zero value otherwise.

### GetItemTypeXrefOk

`func (o *GLExportInventoryTransfer) GetItemTypeXrefOk() (*string, bool)`

GetItemTypeXrefOk returns a tuple with the ItemTypeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTypeXref

`func (o *GLExportInventoryTransfer) SetItemTypeXref(v string)`

SetItemTypeXref sets ItemTypeXref field to given value.

### HasItemTypeXref

`func (o *GLExportInventoryTransfer) HasItemTypeXref() bool`

HasItemTypeXref returns a boolean if a field has been set.

### GetInventoryXref

`func (o *GLExportInventoryTransfer) GetInventoryXref() string`

GetInventoryXref returns the InventoryXref field if non-nil, zero value otherwise.

### GetInventoryXrefOk

`func (o *GLExportInventoryTransfer) GetInventoryXrefOk() (*string, bool)`

GetInventoryXrefOk returns a tuple with the InventoryXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryXref

`func (o *GLExportInventoryTransfer) SetInventoryXref(v string)`

SetInventoryXref sets InventoryXref field to given value.

### HasInventoryXref

`func (o *GLExportInventoryTransfer) HasInventoryXref() bool`

HasInventoryXref returns a boolean if a field has been set.

### GetCogsXref

`func (o *GLExportInventoryTransfer) GetCogsXref() string`

GetCogsXref returns the CogsXref field if non-nil, zero value otherwise.

### GetCogsXrefOk

`func (o *GLExportInventoryTransfer) GetCogsXrefOk() (*string, bool)`

GetCogsXrefOk returns a tuple with the CogsXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogsXref

`func (o *GLExportInventoryTransfer) SetCogsXref(v string)`

SetCogsXref sets CogsXref field to given value.

### HasCogsXref

`func (o *GLExportInventoryTransfer) HasCogsXref() bool`

HasCogsXref returns a boolean if a field has been set.

### GetTaxNote

`func (o *GLExportInventoryTransfer) GetTaxNote() string`

GetTaxNote returns the TaxNote field if non-nil, zero value otherwise.

### GetTaxNoteOk

`func (o *GLExportInventoryTransfer) GetTaxNoteOk() (*string, bool)`

GetTaxNoteOk returns a tuple with the TaxNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNote

`func (o *GLExportInventoryTransfer) SetTaxNote(v string)`

SetTaxNote sets TaxNote field to given value.

### HasTaxNote

`func (o *GLExportInventoryTransfer) HasTaxNote() bool`

HasTaxNote returns a boolean if a field has been set.

### GetOffset

`func (o *GLExportInventoryTransfer) GetOffset() GLExportInventoryTransferOffset`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GLExportInventoryTransfer) GetOffsetOk() (*GLExportInventoryTransferOffset, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GLExportInventoryTransfer) SetOffset(v GLExportInventoryTransferOffset)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GLExportInventoryTransfer) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


