# CatalogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  Max length: 75; | 
**Description** | **string** |  Max length: 60; | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Subcategory** | [**ProductSubCategoryReference**](ProductSubCategoryReference.md) |  | 
**Type** | [**ProductTypeReference**](ProductTypeReference.md) |  | 
**ProductClass** | Pointer to **NullableString** | Defaults to Non-Inventory. | [optional] 
**SerializedFlag** | Pointer to **NullableBool** |  | [optional] 
**SerializedCostFlag** | Pointer to **NullableBool** |  | [optional] 
**PhaseProductFlag** | Pointer to **NullableBool** |  | [optional] 
**UnitOfMeasure** | Pointer to [**UnitOfMeasureReference**](UnitOfMeasureReference.md) |  | [optional] 
**MinStockLevel** | Pointer to **NullableInt32** |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**PriceAttribute** | Pointer to **NullableString** |  | [optional] 
**TaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**DropShipFlag** | Pointer to **NullableBool** |  | [optional] 
**SpecialOrderFlag** | Pointer to **NullableBool** |  | [optional] 
**CustomerDescription** | **string** |  Max length: 6000; | 
**Manufacturer** | Pointer to [**ManufacturerReference**](ManufacturerReference.md) |  | [optional] 
**ManufacturerPartNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**Vendor** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**VendorSku** | Pointer to **string** |  Max length: 50; | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**IntegrationXRef** | Pointer to **string** |  Max length: 50; | [optional] 
**Sla** | Pointer to [**SLAReference**](SLAReference.md) |  | [optional] 
**EntityType** | Pointer to [**EntityTypeReference**](EntityTypeReference.md) |  | [optional] 
**RecurringFlag** | Pointer to **NullableBool** |  | [optional] 
**RecurringRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**RecurringCost** | Pointer to **NullableFloat64** |  | [optional] 
**RecurringOneTimeFlag** | Pointer to **NullableBool** |  | [optional] 
**RecurringBillCycle** | Pointer to [**BillingCycleReference**](BillingCycleReference.md) |  | [optional] 
**RecurringCycleType** | Pointer to **NullableString** |  | [optional] 
**CalculatedPriceFlag** | Pointer to **NullableBool** |  | [optional] 
**CalculatedCostFlag** | Pointer to **NullableBool** |  | [optional] 
**Category** | Pointer to [**ProductCategoryReference**](ProductCategoryReference.md) |  | [optional] 
**CalculatedPrice** | Pointer to **NullableFloat64** |  | [optional] 
**CalculatedCost** | Pointer to **NullableFloat64** |  | [optional] 
**BillableOption** | Pointer to **NullableString** |  | [optional] 
**ConnectWiseID** | Pointer to **string** |  | [optional] 
**AgreementType** | Pointer to [**AgreementTypeReference**](AgreementTypeReference.md) |  | [optional] 
**MarkupPercentage** | Pointer to **NullableFloat64** |  | [optional] 
**MarkupFlag** | Pointer to **NullableBool** |  | [optional] 
**AutoUpdateUnitCostFlag** | Pointer to **NullableBool** |  | [optional] 
**AutoUpdateUnitPriceFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewCatalogItem

`func NewCatalogItem(identifier string, description string, subcategory ProductSubCategoryReference, type_ ProductTypeReference, customerDescription string, ) *CatalogItem`

NewCatalogItem instantiates a new CatalogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemWithDefaults

`func NewCatalogItemWithDefaults() *CatalogItem`

NewCatalogItemWithDefaults instantiates a new CatalogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *CatalogItem) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CatalogItem) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CatalogItem) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetDescription

`func (o *CatalogItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInactiveFlag

`func (o *CatalogItem) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *CatalogItem) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *CatalogItem) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *CatalogItem) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *CatalogItem) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *CatalogItem) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetSubcategory

`func (o *CatalogItem) GetSubcategory() ProductSubCategoryReference`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *CatalogItem) GetSubcategoryOk() (*ProductSubCategoryReference, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *CatalogItem) SetSubcategory(v ProductSubCategoryReference)`

SetSubcategory sets Subcategory field to given value.


### GetType

`func (o *CatalogItem) GetType() ProductTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogItem) GetTypeOk() (*ProductTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogItem) SetType(v ProductTypeReference)`

SetType sets Type field to given value.


### GetProductClass

`func (o *CatalogItem) GetProductClass() string`

GetProductClass returns the ProductClass field if non-nil, zero value otherwise.

### GetProductClassOk

`func (o *CatalogItem) GetProductClassOk() (*string, bool)`

GetProductClassOk returns a tuple with the ProductClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductClass

`func (o *CatalogItem) SetProductClass(v string)`

SetProductClass sets ProductClass field to given value.

### HasProductClass

`func (o *CatalogItem) HasProductClass() bool`

HasProductClass returns a boolean if a field has been set.

### SetProductClassNil

`func (o *CatalogItem) SetProductClassNil(b bool)`

 SetProductClassNil sets the value for ProductClass to be an explicit nil

### UnsetProductClass
`func (o *CatalogItem) UnsetProductClass()`

UnsetProductClass ensures that no value is present for ProductClass, not even an explicit nil
### GetSerializedFlag

`func (o *CatalogItem) GetSerializedFlag() bool`

GetSerializedFlag returns the SerializedFlag field if non-nil, zero value otherwise.

### GetSerializedFlagOk

`func (o *CatalogItem) GetSerializedFlagOk() (*bool, bool)`

GetSerializedFlagOk returns a tuple with the SerializedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedFlag

`func (o *CatalogItem) SetSerializedFlag(v bool)`

SetSerializedFlag sets SerializedFlag field to given value.

### HasSerializedFlag

`func (o *CatalogItem) HasSerializedFlag() bool`

HasSerializedFlag returns a boolean if a field has been set.

### SetSerializedFlagNil

`func (o *CatalogItem) SetSerializedFlagNil(b bool)`

 SetSerializedFlagNil sets the value for SerializedFlag to be an explicit nil

### UnsetSerializedFlag
`func (o *CatalogItem) UnsetSerializedFlag()`

UnsetSerializedFlag ensures that no value is present for SerializedFlag, not even an explicit nil
### GetSerializedCostFlag

`func (o *CatalogItem) GetSerializedCostFlag() bool`

GetSerializedCostFlag returns the SerializedCostFlag field if non-nil, zero value otherwise.

### GetSerializedCostFlagOk

`func (o *CatalogItem) GetSerializedCostFlagOk() (*bool, bool)`

GetSerializedCostFlagOk returns a tuple with the SerializedCostFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedCostFlag

`func (o *CatalogItem) SetSerializedCostFlag(v bool)`

SetSerializedCostFlag sets SerializedCostFlag field to given value.

### HasSerializedCostFlag

`func (o *CatalogItem) HasSerializedCostFlag() bool`

HasSerializedCostFlag returns a boolean if a field has been set.

### SetSerializedCostFlagNil

`func (o *CatalogItem) SetSerializedCostFlagNil(b bool)`

 SetSerializedCostFlagNil sets the value for SerializedCostFlag to be an explicit nil

### UnsetSerializedCostFlag
`func (o *CatalogItem) UnsetSerializedCostFlag()`

UnsetSerializedCostFlag ensures that no value is present for SerializedCostFlag, not even an explicit nil
### GetPhaseProductFlag

`func (o *CatalogItem) GetPhaseProductFlag() bool`

GetPhaseProductFlag returns the PhaseProductFlag field if non-nil, zero value otherwise.

### GetPhaseProductFlagOk

`func (o *CatalogItem) GetPhaseProductFlagOk() (*bool, bool)`

GetPhaseProductFlagOk returns a tuple with the PhaseProductFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseProductFlag

`func (o *CatalogItem) SetPhaseProductFlag(v bool)`

SetPhaseProductFlag sets PhaseProductFlag field to given value.

### HasPhaseProductFlag

`func (o *CatalogItem) HasPhaseProductFlag() bool`

HasPhaseProductFlag returns a boolean if a field has been set.

### SetPhaseProductFlagNil

`func (o *CatalogItem) SetPhaseProductFlagNil(b bool)`

 SetPhaseProductFlagNil sets the value for PhaseProductFlag to be an explicit nil

### UnsetPhaseProductFlag
`func (o *CatalogItem) UnsetPhaseProductFlag()`

UnsetPhaseProductFlag ensures that no value is present for PhaseProductFlag, not even an explicit nil
### GetUnitOfMeasure

`func (o *CatalogItem) GetUnitOfMeasure() UnitOfMeasureReference`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *CatalogItem) GetUnitOfMeasureOk() (*UnitOfMeasureReference, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *CatalogItem) SetUnitOfMeasure(v UnitOfMeasureReference)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *CatalogItem) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetMinStockLevel

`func (o *CatalogItem) GetMinStockLevel() int32`

GetMinStockLevel returns the MinStockLevel field if non-nil, zero value otherwise.

### GetMinStockLevelOk

`func (o *CatalogItem) GetMinStockLevelOk() (*int32, bool)`

GetMinStockLevelOk returns a tuple with the MinStockLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStockLevel

`func (o *CatalogItem) SetMinStockLevel(v int32)`

SetMinStockLevel sets MinStockLevel field to given value.

### HasMinStockLevel

`func (o *CatalogItem) HasMinStockLevel() bool`

HasMinStockLevel returns a boolean if a field has been set.

### SetMinStockLevelNil

`func (o *CatalogItem) SetMinStockLevelNil(b bool)`

 SetMinStockLevelNil sets the value for MinStockLevel to be an explicit nil

### UnsetMinStockLevel
`func (o *CatalogItem) UnsetMinStockLevel()`

UnsetMinStockLevel ensures that no value is present for MinStockLevel, not even an explicit nil
### GetPrice

`func (o *CatalogItem) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogItem) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogItem) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CatalogItem) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CatalogItem) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCost

`func (o *CatalogItem) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CatalogItem) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CatalogItem) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CatalogItem) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *CatalogItem) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *CatalogItem) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetPriceAttribute

`func (o *CatalogItem) GetPriceAttribute() string`

GetPriceAttribute returns the PriceAttribute field if non-nil, zero value otherwise.

### GetPriceAttributeOk

`func (o *CatalogItem) GetPriceAttributeOk() (*string, bool)`

GetPriceAttributeOk returns a tuple with the PriceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAttribute

`func (o *CatalogItem) SetPriceAttribute(v string)`

SetPriceAttribute sets PriceAttribute field to given value.

### HasPriceAttribute

`func (o *CatalogItem) HasPriceAttribute() bool`

HasPriceAttribute returns a boolean if a field has been set.

### SetPriceAttributeNil

`func (o *CatalogItem) SetPriceAttributeNil(b bool)`

 SetPriceAttributeNil sets the value for PriceAttribute to be an explicit nil

### UnsetPriceAttribute
`func (o *CatalogItem) UnsetPriceAttribute()`

UnsetPriceAttribute ensures that no value is present for PriceAttribute, not even an explicit nil
### GetTaxableFlag

`func (o *CatalogItem) GetTaxableFlag() bool`

GetTaxableFlag returns the TaxableFlag field if non-nil, zero value otherwise.

### GetTaxableFlagOk

`func (o *CatalogItem) GetTaxableFlagOk() (*bool, bool)`

GetTaxableFlagOk returns a tuple with the TaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableFlag

`func (o *CatalogItem) SetTaxableFlag(v bool)`

SetTaxableFlag sets TaxableFlag field to given value.

### HasTaxableFlag

`func (o *CatalogItem) HasTaxableFlag() bool`

HasTaxableFlag returns a boolean if a field has been set.

### SetTaxableFlagNil

`func (o *CatalogItem) SetTaxableFlagNil(b bool)`

 SetTaxableFlagNil sets the value for TaxableFlag to be an explicit nil

### UnsetTaxableFlag
`func (o *CatalogItem) UnsetTaxableFlag()`

UnsetTaxableFlag ensures that no value is present for TaxableFlag, not even an explicit nil
### GetDropShipFlag

`func (o *CatalogItem) GetDropShipFlag() bool`

GetDropShipFlag returns the DropShipFlag field if non-nil, zero value otherwise.

### GetDropShipFlagOk

`func (o *CatalogItem) GetDropShipFlagOk() (*bool, bool)`

GetDropShipFlagOk returns a tuple with the DropShipFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropShipFlag

`func (o *CatalogItem) SetDropShipFlag(v bool)`

SetDropShipFlag sets DropShipFlag field to given value.

### HasDropShipFlag

`func (o *CatalogItem) HasDropShipFlag() bool`

HasDropShipFlag returns a boolean if a field has been set.

### SetDropShipFlagNil

`func (o *CatalogItem) SetDropShipFlagNil(b bool)`

 SetDropShipFlagNil sets the value for DropShipFlag to be an explicit nil

### UnsetDropShipFlag
`func (o *CatalogItem) UnsetDropShipFlag()`

UnsetDropShipFlag ensures that no value is present for DropShipFlag, not even an explicit nil
### GetSpecialOrderFlag

`func (o *CatalogItem) GetSpecialOrderFlag() bool`

GetSpecialOrderFlag returns the SpecialOrderFlag field if non-nil, zero value otherwise.

### GetSpecialOrderFlagOk

`func (o *CatalogItem) GetSpecialOrderFlagOk() (*bool, bool)`

GetSpecialOrderFlagOk returns a tuple with the SpecialOrderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialOrderFlag

`func (o *CatalogItem) SetSpecialOrderFlag(v bool)`

SetSpecialOrderFlag sets SpecialOrderFlag field to given value.

### HasSpecialOrderFlag

`func (o *CatalogItem) HasSpecialOrderFlag() bool`

HasSpecialOrderFlag returns a boolean if a field has been set.

### SetSpecialOrderFlagNil

`func (o *CatalogItem) SetSpecialOrderFlagNil(b bool)`

 SetSpecialOrderFlagNil sets the value for SpecialOrderFlag to be an explicit nil

### UnsetSpecialOrderFlag
`func (o *CatalogItem) UnsetSpecialOrderFlag()`

UnsetSpecialOrderFlag ensures that no value is present for SpecialOrderFlag, not even an explicit nil
### GetCustomerDescription

`func (o *CatalogItem) GetCustomerDescription() string`

GetCustomerDescription returns the CustomerDescription field if non-nil, zero value otherwise.

### GetCustomerDescriptionOk

`func (o *CatalogItem) GetCustomerDescriptionOk() (*string, bool)`

GetCustomerDescriptionOk returns a tuple with the CustomerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDescription

`func (o *CatalogItem) SetCustomerDescription(v string)`

SetCustomerDescription sets CustomerDescription field to given value.


### GetManufacturer

`func (o *CatalogItem) GetManufacturer() ManufacturerReference`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *CatalogItem) GetManufacturerOk() (*ManufacturerReference, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *CatalogItem) SetManufacturer(v ManufacturerReference)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *CatalogItem) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetManufacturerPartNumber

`func (o *CatalogItem) GetManufacturerPartNumber() string`

GetManufacturerPartNumber returns the ManufacturerPartNumber field if non-nil, zero value otherwise.

### GetManufacturerPartNumberOk

`func (o *CatalogItem) GetManufacturerPartNumberOk() (*string, bool)`

GetManufacturerPartNumberOk returns a tuple with the ManufacturerPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerPartNumber

`func (o *CatalogItem) SetManufacturerPartNumber(v string)`

SetManufacturerPartNumber sets ManufacturerPartNumber field to given value.

### HasManufacturerPartNumber

`func (o *CatalogItem) HasManufacturerPartNumber() bool`

HasManufacturerPartNumber returns a boolean if a field has been set.

### GetVendor

`func (o *CatalogItem) GetVendor() CompanyReference`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CatalogItem) GetVendorOk() (*CompanyReference, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CatalogItem) SetVendor(v CompanyReference)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CatalogItem) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVendorSku

`func (o *CatalogItem) GetVendorSku() string`

GetVendorSku returns the VendorSku field if non-nil, zero value otherwise.

### GetVendorSkuOk

`func (o *CatalogItem) GetVendorSkuOk() (*string, bool)`

GetVendorSkuOk returns a tuple with the VendorSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSku

`func (o *CatalogItem) SetVendorSku(v string)`

SetVendorSku sets VendorSku field to given value.

### HasVendorSku

`func (o *CatalogItem) HasVendorSku() bool`

HasVendorSku returns a boolean if a field has been set.

### GetNotes

`func (o *CatalogItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CatalogItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CatalogItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CatalogItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIntegrationXRef

`func (o *CatalogItem) GetIntegrationXRef() string`

GetIntegrationXRef returns the IntegrationXRef field if non-nil, zero value otherwise.

### GetIntegrationXRefOk

`func (o *CatalogItem) GetIntegrationXRefOk() (*string, bool)`

GetIntegrationXRefOk returns a tuple with the IntegrationXRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationXRef

`func (o *CatalogItem) SetIntegrationXRef(v string)`

SetIntegrationXRef sets IntegrationXRef field to given value.

### HasIntegrationXRef

`func (o *CatalogItem) HasIntegrationXRef() bool`

HasIntegrationXRef returns a boolean if a field has been set.

### GetSla

`func (o *CatalogItem) GetSla() SLAReference`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *CatalogItem) GetSlaOk() (*SLAReference, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *CatalogItem) SetSla(v SLAReference)`

SetSla sets Sla field to given value.

### HasSla

`func (o *CatalogItem) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetEntityType

`func (o *CatalogItem) GetEntityType() EntityTypeReference`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CatalogItem) GetEntityTypeOk() (*EntityTypeReference, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CatalogItem) SetEntityType(v EntityTypeReference)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *CatalogItem) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetRecurringFlag

`func (o *CatalogItem) GetRecurringFlag() bool`

GetRecurringFlag returns the RecurringFlag field if non-nil, zero value otherwise.

### GetRecurringFlagOk

`func (o *CatalogItem) GetRecurringFlagOk() (*bool, bool)`

GetRecurringFlagOk returns a tuple with the RecurringFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringFlag

`func (o *CatalogItem) SetRecurringFlag(v bool)`

SetRecurringFlag sets RecurringFlag field to given value.

### HasRecurringFlag

`func (o *CatalogItem) HasRecurringFlag() bool`

HasRecurringFlag returns a boolean if a field has been set.

### SetRecurringFlagNil

`func (o *CatalogItem) SetRecurringFlagNil(b bool)`

 SetRecurringFlagNil sets the value for RecurringFlag to be an explicit nil

### UnsetRecurringFlag
`func (o *CatalogItem) UnsetRecurringFlag()`

UnsetRecurringFlag ensures that no value is present for RecurringFlag, not even an explicit nil
### GetRecurringRevenue

`func (o *CatalogItem) GetRecurringRevenue() float64`

GetRecurringRevenue returns the RecurringRevenue field if non-nil, zero value otherwise.

### GetRecurringRevenueOk

`func (o *CatalogItem) GetRecurringRevenueOk() (*float64, bool)`

GetRecurringRevenueOk returns a tuple with the RecurringRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringRevenue

`func (o *CatalogItem) SetRecurringRevenue(v float64)`

SetRecurringRevenue sets RecurringRevenue field to given value.

### HasRecurringRevenue

`func (o *CatalogItem) HasRecurringRevenue() bool`

HasRecurringRevenue returns a boolean if a field has been set.

### SetRecurringRevenueNil

`func (o *CatalogItem) SetRecurringRevenueNil(b bool)`

 SetRecurringRevenueNil sets the value for RecurringRevenue to be an explicit nil

### UnsetRecurringRevenue
`func (o *CatalogItem) UnsetRecurringRevenue()`

UnsetRecurringRevenue ensures that no value is present for RecurringRevenue, not even an explicit nil
### GetRecurringCost

`func (o *CatalogItem) GetRecurringCost() float64`

GetRecurringCost returns the RecurringCost field if non-nil, zero value otherwise.

### GetRecurringCostOk

`func (o *CatalogItem) GetRecurringCostOk() (*float64, bool)`

GetRecurringCostOk returns a tuple with the RecurringCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringCost

`func (o *CatalogItem) SetRecurringCost(v float64)`

SetRecurringCost sets RecurringCost field to given value.

### HasRecurringCost

`func (o *CatalogItem) HasRecurringCost() bool`

HasRecurringCost returns a boolean if a field has been set.

### SetRecurringCostNil

`func (o *CatalogItem) SetRecurringCostNil(b bool)`

 SetRecurringCostNil sets the value for RecurringCost to be an explicit nil

### UnsetRecurringCost
`func (o *CatalogItem) UnsetRecurringCost()`

UnsetRecurringCost ensures that no value is present for RecurringCost, not even an explicit nil
### GetRecurringOneTimeFlag

`func (o *CatalogItem) GetRecurringOneTimeFlag() bool`

GetRecurringOneTimeFlag returns the RecurringOneTimeFlag field if non-nil, zero value otherwise.

### GetRecurringOneTimeFlagOk

`func (o *CatalogItem) GetRecurringOneTimeFlagOk() (*bool, bool)`

GetRecurringOneTimeFlagOk returns a tuple with the RecurringOneTimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringOneTimeFlag

`func (o *CatalogItem) SetRecurringOneTimeFlag(v bool)`

SetRecurringOneTimeFlag sets RecurringOneTimeFlag field to given value.

### HasRecurringOneTimeFlag

`func (o *CatalogItem) HasRecurringOneTimeFlag() bool`

HasRecurringOneTimeFlag returns a boolean if a field has been set.

### SetRecurringOneTimeFlagNil

`func (o *CatalogItem) SetRecurringOneTimeFlagNil(b bool)`

 SetRecurringOneTimeFlagNil sets the value for RecurringOneTimeFlag to be an explicit nil

### UnsetRecurringOneTimeFlag
`func (o *CatalogItem) UnsetRecurringOneTimeFlag()`

UnsetRecurringOneTimeFlag ensures that no value is present for RecurringOneTimeFlag, not even an explicit nil
### GetRecurringBillCycle

`func (o *CatalogItem) GetRecurringBillCycle() BillingCycleReference`

GetRecurringBillCycle returns the RecurringBillCycle field if non-nil, zero value otherwise.

### GetRecurringBillCycleOk

`func (o *CatalogItem) GetRecurringBillCycleOk() (*BillingCycleReference, bool)`

GetRecurringBillCycleOk returns a tuple with the RecurringBillCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringBillCycle

`func (o *CatalogItem) SetRecurringBillCycle(v BillingCycleReference)`

SetRecurringBillCycle sets RecurringBillCycle field to given value.

### HasRecurringBillCycle

`func (o *CatalogItem) HasRecurringBillCycle() bool`

HasRecurringBillCycle returns a boolean if a field has been set.

### GetRecurringCycleType

`func (o *CatalogItem) GetRecurringCycleType() string`

GetRecurringCycleType returns the RecurringCycleType field if non-nil, zero value otherwise.

### GetRecurringCycleTypeOk

`func (o *CatalogItem) GetRecurringCycleTypeOk() (*string, bool)`

GetRecurringCycleTypeOk returns a tuple with the RecurringCycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringCycleType

`func (o *CatalogItem) SetRecurringCycleType(v string)`

SetRecurringCycleType sets RecurringCycleType field to given value.

### HasRecurringCycleType

`func (o *CatalogItem) HasRecurringCycleType() bool`

HasRecurringCycleType returns a boolean if a field has been set.

### SetRecurringCycleTypeNil

`func (o *CatalogItem) SetRecurringCycleTypeNil(b bool)`

 SetRecurringCycleTypeNil sets the value for RecurringCycleType to be an explicit nil

### UnsetRecurringCycleType
`func (o *CatalogItem) UnsetRecurringCycleType()`

UnsetRecurringCycleType ensures that no value is present for RecurringCycleType, not even an explicit nil
### GetCalculatedPriceFlag

`func (o *CatalogItem) GetCalculatedPriceFlag() bool`

GetCalculatedPriceFlag returns the CalculatedPriceFlag field if non-nil, zero value otherwise.

### GetCalculatedPriceFlagOk

`func (o *CatalogItem) GetCalculatedPriceFlagOk() (*bool, bool)`

GetCalculatedPriceFlagOk returns a tuple with the CalculatedPriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedPriceFlag

`func (o *CatalogItem) SetCalculatedPriceFlag(v bool)`

SetCalculatedPriceFlag sets CalculatedPriceFlag field to given value.

### HasCalculatedPriceFlag

`func (o *CatalogItem) HasCalculatedPriceFlag() bool`

HasCalculatedPriceFlag returns a boolean if a field has been set.

### SetCalculatedPriceFlagNil

`func (o *CatalogItem) SetCalculatedPriceFlagNil(b bool)`

 SetCalculatedPriceFlagNil sets the value for CalculatedPriceFlag to be an explicit nil

### UnsetCalculatedPriceFlag
`func (o *CatalogItem) UnsetCalculatedPriceFlag()`

UnsetCalculatedPriceFlag ensures that no value is present for CalculatedPriceFlag, not even an explicit nil
### GetCalculatedCostFlag

`func (o *CatalogItem) GetCalculatedCostFlag() bool`

GetCalculatedCostFlag returns the CalculatedCostFlag field if non-nil, zero value otherwise.

### GetCalculatedCostFlagOk

`func (o *CatalogItem) GetCalculatedCostFlagOk() (*bool, bool)`

GetCalculatedCostFlagOk returns a tuple with the CalculatedCostFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedCostFlag

`func (o *CatalogItem) SetCalculatedCostFlag(v bool)`

SetCalculatedCostFlag sets CalculatedCostFlag field to given value.

### HasCalculatedCostFlag

`func (o *CatalogItem) HasCalculatedCostFlag() bool`

HasCalculatedCostFlag returns a boolean if a field has been set.

### SetCalculatedCostFlagNil

`func (o *CatalogItem) SetCalculatedCostFlagNil(b bool)`

 SetCalculatedCostFlagNil sets the value for CalculatedCostFlag to be an explicit nil

### UnsetCalculatedCostFlag
`func (o *CatalogItem) UnsetCalculatedCostFlag()`

UnsetCalculatedCostFlag ensures that no value is present for CalculatedCostFlag, not even an explicit nil
### GetCategory

`func (o *CatalogItem) GetCategory() ProductCategoryReference`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CatalogItem) GetCategoryOk() (*ProductCategoryReference, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CatalogItem) SetCategory(v ProductCategoryReference)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CatalogItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCalculatedPrice

`func (o *CatalogItem) GetCalculatedPrice() float64`

GetCalculatedPrice returns the CalculatedPrice field if non-nil, zero value otherwise.

### GetCalculatedPriceOk

`func (o *CatalogItem) GetCalculatedPriceOk() (*float64, bool)`

GetCalculatedPriceOk returns a tuple with the CalculatedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedPrice

`func (o *CatalogItem) SetCalculatedPrice(v float64)`

SetCalculatedPrice sets CalculatedPrice field to given value.

### HasCalculatedPrice

`func (o *CatalogItem) HasCalculatedPrice() bool`

HasCalculatedPrice returns a boolean if a field has been set.

### SetCalculatedPriceNil

`func (o *CatalogItem) SetCalculatedPriceNil(b bool)`

 SetCalculatedPriceNil sets the value for CalculatedPrice to be an explicit nil

### UnsetCalculatedPrice
`func (o *CatalogItem) UnsetCalculatedPrice()`

UnsetCalculatedPrice ensures that no value is present for CalculatedPrice, not even an explicit nil
### GetCalculatedCost

`func (o *CatalogItem) GetCalculatedCost() float64`

GetCalculatedCost returns the CalculatedCost field if non-nil, zero value otherwise.

### GetCalculatedCostOk

`func (o *CatalogItem) GetCalculatedCostOk() (*float64, bool)`

GetCalculatedCostOk returns a tuple with the CalculatedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedCost

`func (o *CatalogItem) SetCalculatedCost(v float64)`

SetCalculatedCost sets CalculatedCost field to given value.

### HasCalculatedCost

`func (o *CatalogItem) HasCalculatedCost() bool`

HasCalculatedCost returns a boolean if a field has been set.

### SetCalculatedCostNil

`func (o *CatalogItem) SetCalculatedCostNil(b bool)`

 SetCalculatedCostNil sets the value for CalculatedCost to be an explicit nil

### UnsetCalculatedCost
`func (o *CatalogItem) UnsetCalculatedCost()`

UnsetCalculatedCost ensures that no value is present for CalculatedCost, not even an explicit nil
### GetBillableOption

`func (o *CatalogItem) GetBillableOption() string`

GetBillableOption returns the BillableOption field if non-nil, zero value otherwise.

### GetBillableOptionOk

`func (o *CatalogItem) GetBillableOptionOk() (*string, bool)`

GetBillableOptionOk returns a tuple with the BillableOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableOption

`func (o *CatalogItem) SetBillableOption(v string)`

SetBillableOption sets BillableOption field to given value.

### HasBillableOption

`func (o *CatalogItem) HasBillableOption() bool`

HasBillableOption returns a boolean if a field has been set.

### SetBillableOptionNil

`func (o *CatalogItem) SetBillableOptionNil(b bool)`

 SetBillableOptionNil sets the value for BillableOption to be an explicit nil

### UnsetBillableOption
`func (o *CatalogItem) UnsetBillableOption()`

UnsetBillableOption ensures that no value is present for BillableOption, not even an explicit nil
### GetConnectWiseID

`func (o *CatalogItem) GetConnectWiseID() string`

GetConnectWiseID returns the ConnectWiseID field if non-nil, zero value otherwise.

### GetConnectWiseIDOk

`func (o *CatalogItem) GetConnectWiseIDOk() (*string, bool)`

GetConnectWiseIDOk returns a tuple with the ConnectWiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseID

`func (o *CatalogItem) SetConnectWiseID(v string)`

SetConnectWiseID sets ConnectWiseID field to given value.

### HasConnectWiseID

`func (o *CatalogItem) HasConnectWiseID() bool`

HasConnectWiseID returns a boolean if a field has been set.

### GetAgreementType

`func (o *CatalogItem) GetAgreementType() AgreementTypeReference`

GetAgreementType returns the AgreementType field if non-nil, zero value otherwise.

### GetAgreementTypeOk

`func (o *CatalogItem) GetAgreementTypeOk() (*AgreementTypeReference, bool)`

GetAgreementTypeOk returns a tuple with the AgreementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementType

`func (o *CatalogItem) SetAgreementType(v AgreementTypeReference)`

SetAgreementType sets AgreementType field to given value.

### HasAgreementType

`func (o *CatalogItem) HasAgreementType() bool`

HasAgreementType returns a boolean if a field has been set.

### GetMarkupPercentage

`func (o *CatalogItem) GetMarkupPercentage() float64`

GetMarkupPercentage returns the MarkupPercentage field if non-nil, zero value otherwise.

### GetMarkupPercentageOk

`func (o *CatalogItem) GetMarkupPercentageOk() (*float64, bool)`

GetMarkupPercentageOk returns a tuple with the MarkupPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupPercentage

`func (o *CatalogItem) SetMarkupPercentage(v float64)`

SetMarkupPercentage sets MarkupPercentage field to given value.

### HasMarkupPercentage

`func (o *CatalogItem) HasMarkupPercentage() bool`

HasMarkupPercentage returns a boolean if a field has been set.

### SetMarkupPercentageNil

`func (o *CatalogItem) SetMarkupPercentageNil(b bool)`

 SetMarkupPercentageNil sets the value for MarkupPercentage to be an explicit nil

### UnsetMarkupPercentage
`func (o *CatalogItem) UnsetMarkupPercentage()`

UnsetMarkupPercentage ensures that no value is present for MarkupPercentage, not even an explicit nil
### GetMarkupFlag

`func (o *CatalogItem) GetMarkupFlag() bool`

GetMarkupFlag returns the MarkupFlag field if non-nil, zero value otherwise.

### GetMarkupFlagOk

`func (o *CatalogItem) GetMarkupFlagOk() (*bool, bool)`

GetMarkupFlagOk returns a tuple with the MarkupFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupFlag

`func (o *CatalogItem) SetMarkupFlag(v bool)`

SetMarkupFlag sets MarkupFlag field to given value.

### HasMarkupFlag

`func (o *CatalogItem) HasMarkupFlag() bool`

HasMarkupFlag returns a boolean if a field has been set.

### SetMarkupFlagNil

`func (o *CatalogItem) SetMarkupFlagNil(b bool)`

 SetMarkupFlagNil sets the value for MarkupFlag to be an explicit nil

### UnsetMarkupFlag
`func (o *CatalogItem) UnsetMarkupFlag()`

UnsetMarkupFlag ensures that no value is present for MarkupFlag, not even an explicit nil
### GetAutoUpdateUnitCostFlag

`func (o *CatalogItem) GetAutoUpdateUnitCostFlag() bool`

GetAutoUpdateUnitCostFlag returns the AutoUpdateUnitCostFlag field if non-nil, zero value otherwise.

### GetAutoUpdateUnitCostFlagOk

`func (o *CatalogItem) GetAutoUpdateUnitCostFlagOk() (*bool, bool)`

GetAutoUpdateUnitCostFlagOk returns a tuple with the AutoUpdateUnitCostFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateUnitCostFlag

`func (o *CatalogItem) SetAutoUpdateUnitCostFlag(v bool)`

SetAutoUpdateUnitCostFlag sets AutoUpdateUnitCostFlag field to given value.

### HasAutoUpdateUnitCostFlag

`func (o *CatalogItem) HasAutoUpdateUnitCostFlag() bool`

HasAutoUpdateUnitCostFlag returns a boolean if a field has been set.

### SetAutoUpdateUnitCostFlagNil

`func (o *CatalogItem) SetAutoUpdateUnitCostFlagNil(b bool)`

 SetAutoUpdateUnitCostFlagNil sets the value for AutoUpdateUnitCostFlag to be an explicit nil

### UnsetAutoUpdateUnitCostFlag
`func (o *CatalogItem) UnsetAutoUpdateUnitCostFlag()`

UnsetAutoUpdateUnitCostFlag ensures that no value is present for AutoUpdateUnitCostFlag, not even an explicit nil
### GetAutoUpdateUnitPriceFlag

`func (o *CatalogItem) GetAutoUpdateUnitPriceFlag() bool`

GetAutoUpdateUnitPriceFlag returns the AutoUpdateUnitPriceFlag field if non-nil, zero value otherwise.

### GetAutoUpdateUnitPriceFlagOk

`func (o *CatalogItem) GetAutoUpdateUnitPriceFlagOk() (*bool, bool)`

GetAutoUpdateUnitPriceFlagOk returns a tuple with the AutoUpdateUnitPriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateUnitPriceFlag

`func (o *CatalogItem) SetAutoUpdateUnitPriceFlag(v bool)`

SetAutoUpdateUnitPriceFlag sets AutoUpdateUnitPriceFlag field to given value.

### HasAutoUpdateUnitPriceFlag

`func (o *CatalogItem) HasAutoUpdateUnitPriceFlag() bool`

HasAutoUpdateUnitPriceFlag returns a boolean if a field has been set.

### SetAutoUpdateUnitPriceFlagNil

`func (o *CatalogItem) SetAutoUpdateUnitPriceFlagNil(b bool)`

 SetAutoUpdateUnitPriceFlagNil sets the value for AutoUpdateUnitPriceFlag to be an explicit nil

### UnsetAutoUpdateUnitPriceFlag
`func (o *CatalogItem) UnsetAutoUpdateUnitPriceFlag()`

UnsetAutoUpdateUnitPriceFlag ensures that no value is present for AutoUpdateUnitPriceFlag, not even an explicit nil
### GetInfo

`func (o *CatalogItem) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CatalogItem) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CatalogItem) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CatalogItem) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *CatalogItem) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CatalogItem) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CatalogItem) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CatalogItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


