# ProductItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CatalogItem** | [**CatalogItemReference**](CatalogItemReference.md) |  | 
**Description** | Pointer to **string** |  Max length: 2000; | [optional] 
**SequenceNumber** | Pointer to **NullableFloat64** |  | [optional] 
**Quantity** | Pointer to **NullableFloat64** |  | [optional] 
**UnitOfMeasure** | Pointer to [**UnitOfMeasureReference**](UnitOfMeasureReference.md) |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**ExtPrice** | Pointer to **NullableFloat64** |  | [optional] 
**ExtCost** | Pointer to **NullableFloat64** |  | [optional] 
**Discount** | Pointer to **NullableFloat64** |  | [optional] 
**Margin** | Pointer to **NullableFloat64** |  | [optional] 
**AgreementAmount** | Pointer to **NullableFloat64** |  | [optional] 
**PriceMethod** | Pointer to **NullableString** |  | [optional] 
**BillableOption** | **NullableString** |  | 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  Required On Updates; | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  Required On Updates; | [optional] 
**BusinessUnit** | Pointer to [**BillingUnitReference**](BillingUnitReference.md) |  | [optional] 
**Vendor** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**VendorSku** | Pointer to **string** |  Max length: 50; | [optional] 
**TaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**DropshipFlag** | Pointer to **NullableBool** |  | [optional] 
**SpecialOrderFlag** | Pointer to **NullableBool** |  | [optional] 
**PhaseProductFlag** | Pointer to **NullableBool** |  | [optional] 
**CancelledFlag** | Pointer to **NullableBool** |  | [optional] 
**QuantityCancelled** | Pointer to **NullableFloat64** |  | [optional] 
**CancelledReason** | Pointer to **string** |  Max length: 100; | [optional] 
**CustomerDescription** | Pointer to **string** |  Max length: 6000; Required On Updates; | [optional] 
**InternalNotes** | Pointer to **string** |  Max length: 1000; | [optional] 
**ProductSuppliedFlag** | Pointer to **NullableBool** |  | [optional] 
**SubContractorShipToId** | Pointer to **NullableInt32** |  | [optional] 
**SubContractorAmountLimit** | Pointer to **NullableFloat64** |  | [optional] 
**Recurring** | Pointer to [**ProductRecurring**](ProductRecurring.md) |  | [optional] 
**Sla** | Pointer to [**SLAReference**](SLAReference.md) |  | [optional] 
**EntityType** | Pointer to [**EntityTypeReference**](EntityTypeReference.md) |  | [optional] 
**Ticket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**Phase** | Pointer to [**ProjectPhaseReference**](ProjectPhaseReference.md) |  | [optional] 
**SalesOrder** | Pointer to [**SalesOrderReference**](SalesOrderReference.md) |  | [optional] 
**Opportunity** | Pointer to [**OpportunityReference**](OpportunityReference.md) |  | [optional] 
**Invoice** | Pointer to [**InvoiceReference**](InvoiceReference.md) |  | [optional] 
**WarehouseId** | Pointer to **NullableInt32** |  | [optional] 
**WarehouseIdObject** | Pointer to [**WarehouseReference**](WarehouseReference.md) |  | [optional] 
**WarehouseBinId** | Pointer to **NullableInt32** |  | [optional] 
**WarehouseBinIdObject** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**CalculatedPriceFlag** | Pointer to **NullableBool** |  | [optional] 
**CalculatedCostFlag** | Pointer to **NullableBool** |  | [optional] 
**ForecastDetailId** | Pointer to **NullableInt32** |  | [optional] 
**CancelledBy** | Pointer to **NullableInt32** |  | [optional] 
**CancelledDate** | Pointer to **time.Time** |  | [optional] 
**Warehouse** | Pointer to **string** |  | [optional] 
**WarehouseBin** | Pointer to **string** |  | [optional] 
**PurchaseDate** | Pointer to **time.Time** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**IntegrationXRef** | Pointer to **string** |  | [optional] 
**ListPrice** | Pointer to **NullableFloat64** |  | [optional] 
**SerialNumberIds** | Pointer to **[]int32** |  | [optional] 
**SerialNumbers** | Pointer to **[]string** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ForecastStatus** | Pointer to [**OpportunityStatusReference**](OpportunityStatusReference.md) |  | [optional] 
**ProductClass** | Pointer to **NullableString** |  | [optional] 
**NeedToPurchaseFlag** | Pointer to **NullableBool** |  | [optional] 
**NeedToOrderQuantity** | Pointer to **NullableInt32** |  | [optional] 
**MinimumStockFlag** | Pointer to **NullableBool** |  | [optional] 
**ShipSet** | Pointer to **string** |  Max length: 10; | [optional] 
**CalculatedPrice** | Pointer to **NullableFloat64** |  | [optional] 
**CalculatedCost** | Pointer to **NullableFloat64** |  | [optional] 
**InvoiceGrouping** | Pointer to [**InvoiceGroupingReference**](InvoiceGroupingReference.md) |  | [optional] 
**PoApprovedFlag** | Pointer to **NullableBool** |  | [optional] 
**Uom** | Pointer to **string** |  | [optional] 
**AddComponentsFlag** | Pointer to **NullableBool** |  | [optional] 
**IgnorePricingSchedulesFlag** | Pointer to **NullableBool** |  | [optional] 
**AsioSubscriptionsID** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**BypassForecastUpdate** | Pointer to **NullableBool** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewProductItem

`func NewProductItem(catalogItem CatalogItemReference, billableOption NullableString, ) *ProductItem`

NewProductItem instantiates a new ProductItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductItemWithDefaults

`func NewProductItemWithDefaults() *ProductItem`

NewProductItemWithDefaults instantiates a new ProductItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCatalogItem

`func (o *ProductItem) GetCatalogItem() CatalogItemReference`

GetCatalogItem returns the CatalogItem field if non-nil, zero value otherwise.

### GetCatalogItemOk

`func (o *ProductItem) GetCatalogItemOk() (*CatalogItemReference, bool)`

GetCatalogItemOk returns a tuple with the CatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItem

`func (o *ProductItem) SetCatalogItem(v CatalogItemReference)`

SetCatalogItem sets CatalogItem field to given value.


### GetDescription

`func (o *ProductItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *ProductItem) GetSequenceNumber() float64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ProductItem) GetSequenceNumberOk() (*float64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ProductItem) SetSequenceNumber(v float64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *ProductItem) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### SetSequenceNumberNil

`func (o *ProductItem) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *ProductItem) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetQuantity

`func (o *ProductItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductItem) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *ProductItem) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *ProductItem) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetUnitOfMeasure

`func (o *ProductItem) GetUnitOfMeasure() UnitOfMeasureReference`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *ProductItem) GetUnitOfMeasureOk() (*UnitOfMeasureReference, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *ProductItem) SetUnitOfMeasure(v UnitOfMeasureReference)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *ProductItem) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetPrice

`func (o *ProductItem) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductItem) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductItem) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ProductItem) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ProductItem) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCost

`func (o *ProductItem) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProductItem) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProductItem) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ProductItem) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *ProductItem) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *ProductItem) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetExtPrice

`func (o *ProductItem) GetExtPrice() float64`

GetExtPrice returns the ExtPrice field if non-nil, zero value otherwise.

### GetExtPriceOk

`func (o *ProductItem) GetExtPriceOk() (*float64, bool)`

GetExtPriceOk returns a tuple with the ExtPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtPrice

`func (o *ProductItem) SetExtPrice(v float64)`

SetExtPrice sets ExtPrice field to given value.

### HasExtPrice

`func (o *ProductItem) HasExtPrice() bool`

HasExtPrice returns a boolean if a field has been set.

### SetExtPriceNil

`func (o *ProductItem) SetExtPriceNil(b bool)`

 SetExtPriceNil sets the value for ExtPrice to be an explicit nil

### UnsetExtPrice
`func (o *ProductItem) UnsetExtPrice()`

UnsetExtPrice ensures that no value is present for ExtPrice, not even an explicit nil
### GetExtCost

`func (o *ProductItem) GetExtCost() float64`

GetExtCost returns the ExtCost field if non-nil, zero value otherwise.

### GetExtCostOk

`func (o *ProductItem) GetExtCostOk() (*float64, bool)`

GetExtCostOk returns a tuple with the ExtCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtCost

`func (o *ProductItem) SetExtCost(v float64)`

SetExtCost sets ExtCost field to given value.

### HasExtCost

`func (o *ProductItem) HasExtCost() bool`

HasExtCost returns a boolean if a field has been set.

### SetExtCostNil

`func (o *ProductItem) SetExtCostNil(b bool)`

 SetExtCostNil sets the value for ExtCost to be an explicit nil

### UnsetExtCost
`func (o *ProductItem) UnsetExtCost()`

UnsetExtCost ensures that no value is present for ExtCost, not even an explicit nil
### GetDiscount

`func (o *ProductItem) GetDiscount() float64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ProductItem) GetDiscountOk() (*float64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ProductItem) SetDiscount(v float64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *ProductItem) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *ProductItem) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *ProductItem) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
### GetMargin

`func (o *ProductItem) GetMargin() float64`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *ProductItem) GetMarginOk() (*float64, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *ProductItem) SetMargin(v float64)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *ProductItem) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### SetMarginNil

`func (o *ProductItem) SetMarginNil(b bool)`

 SetMarginNil sets the value for Margin to be an explicit nil

### UnsetMargin
`func (o *ProductItem) UnsetMargin()`

UnsetMargin ensures that no value is present for Margin, not even an explicit nil
### GetAgreementAmount

`func (o *ProductItem) GetAgreementAmount() float64`

GetAgreementAmount returns the AgreementAmount field if non-nil, zero value otherwise.

### GetAgreementAmountOk

`func (o *ProductItem) GetAgreementAmountOk() (*float64, bool)`

GetAgreementAmountOk returns a tuple with the AgreementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementAmount

`func (o *ProductItem) SetAgreementAmount(v float64)`

SetAgreementAmount sets AgreementAmount field to given value.

### HasAgreementAmount

`func (o *ProductItem) HasAgreementAmount() bool`

HasAgreementAmount returns a boolean if a field has been set.

### SetAgreementAmountNil

`func (o *ProductItem) SetAgreementAmountNil(b bool)`

 SetAgreementAmountNil sets the value for AgreementAmount to be an explicit nil

### UnsetAgreementAmount
`func (o *ProductItem) UnsetAgreementAmount()`

UnsetAgreementAmount ensures that no value is present for AgreementAmount, not even an explicit nil
### GetPriceMethod

`func (o *ProductItem) GetPriceMethod() string`

GetPriceMethod returns the PriceMethod field if non-nil, zero value otherwise.

### GetPriceMethodOk

`func (o *ProductItem) GetPriceMethodOk() (*string, bool)`

GetPriceMethodOk returns a tuple with the PriceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMethod

`func (o *ProductItem) SetPriceMethod(v string)`

SetPriceMethod sets PriceMethod field to given value.

### HasPriceMethod

`func (o *ProductItem) HasPriceMethod() bool`

HasPriceMethod returns a boolean if a field has been set.

### SetPriceMethodNil

`func (o *ProductItem) SetPriceMethodNil(b bool)`

 SetPriceMethodNil sets the value for PriceMethod to be an explicit nil

### UnsetPriceMethod
`func (o *ProductItem) UnsetPriceMethod()`

UnsetPriceMethod ensures that no value is present for PriceMethod, not even an explicit nil
### GetBillableOption

`func (o *ProductItem) GetBillableOption() string`

GetBillableOption returns the BillableOption field if non-nil, zero value otherwise.

### GetBillableOptionOk

`func (o *ProductItem) GetBillableOptionOk() (*string, bool)`

GetBillableOptionOk returns a tuple with the BillableOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableOption

`func (o *ProductItem) SetBillableOption(v string)`

SetBillableOption sets BillableOption field to given value.


### SetBillableOptionNil

`func (o *ProductItem) SetBillableOptionNil(b bool)`

 SetBillableOptionNil sets the value for BillableOption to be an explicit nil

### UnsetBillableOption
`func (o *ProductItem) UnsetBillableOption()`

UnsetBillableOption ensures that no value is present for BillableOption, not even an explicit nil
### GetAgreement

`func (o *ProductItem) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *ProductItem) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *ProductItem) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *ProductItem) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetLocationId

`func (o *ProductItem) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ProductItem) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ProductItem) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *ProductItem) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *ProductItem) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *ProductItem) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetLocation

`func (o *ProductItem) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProductItem) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProductItem) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ProductItem) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *ProductItem) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *ProductItem) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *ProductItem) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *ProductItem) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *ProductItem) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *ProductItem) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetBusinessUnit

`func (o *ProductItem) GetBusinessUnit() BillingUnitReference`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *ProductItem) GetBusinessUnitOk() (*BillingUnitReference, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *ProductItem) SetBusinessUnit(v BillingUnitReference)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *ProductItem) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetVendor

`func (o *ProductItem) GetVendor() CompanyReference`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ProductItem) GetVendorOk() (*CompanyReference, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ProductItem) SetVendor(v CompanyReference)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ProductItem) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVendorSku

`func (o *ProductItem) GetVendorSku() string`

GetVendorSku returns the VendorSku field if non-nil, zero value otherwise.

### GetVendorSkuOk

`func (o *ProductItem) GetVendorSkuOk() (*string, bool)`

GetVendorSkuOk returns a tuple with the VendorSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSku

`func (o *ProductItem) SetVendorSku(v string)`

SetVendorSku sets VendorSku field to given value.

### HasVendorSku

`func (o *ProductItem) HasVendorSku() bool`

HasVendorSku returns a boolean if a field has been set.

### GetTaxableFlag

`func (o *ProductItem) GetTaxableFlag() bool`

GetTaxableFlag returns the TaxableFlag field if non-nil, zero value otherwise.

### GetTaxableFlagOk

`func (o *ProductItem) GetTaxableFlagOk() (*bool, bool)`

GetTaxableFlagOk returns a tuple with the TaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableFlag

`func (o *ProductItem) SetTaxableFlag(v bool)`

SetTaxableFlag sets TaxableFlag field to given value.

### HasTaxableFlag

`func (o *ProductItem) HasTaxableFlag() bool`

HasTaxableFlag returns a boolean if a field has been set.

### SetTaxableFlagNil

`func (o *ProductItem) SetTaxableFlagNil(b bool)`

 SetTaxableFlagNil sets the value for TaxableFlag to be an explicit nil

### UnsetTaxableFlag
`func (o *ProductItem) UnsetTaxableFlag()`

UnsetTaxableFlag ensures that no value is present for TaxableFlag, not even an explicit nil
### GetDropshipFlag

`func (o *ProductItem) GetDropshipFlag() bool`

GetDropshipFlag returns the DropshipFlag field if non-nil, zero value otherwise.

### GetDropshipFlagOk

`func (o *ProductItem) GetDropshipFlagOk() (*bool, bool)`

GetDropshipFlagOk returns a tuple with the DropshipFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropshipFlag

`func (o *ProductItem) SetDropshipFlag(v bool)`

SetDropshipFlag sets DropshipFlag field to given value.

### HasDropshipFlag

`func (o *ProductItem) HasDropshipFlag() bool`

HasDropshipFlag returns a boolean if a field has been set.

### SetDropshipFlagNil

`func (o *ProductItem) SetDropshipFlagNil(b bool)`

 SetDropshipFlagNil sets the value for DropshipFlag to be an explicit nil

### UnsetDropshipFlag
`func (o *ProductItem) UnsetDropshipFlag()`

UnsetDropshipFlag ensures that no value is present for DropshipFlag, not even an explicit nil
### GetSpecialOrderFlag

`func (o *ProductItem) GetSpecialOrderFlag() bool`

GetSpecialOrderFlag returns the SpecialOrderFlag field if non-nil, zero value otherwise.

### GetSpecialOrderFlagOk

`func (o *ProductItem) GetSpecialOrderFlagOk() (*bool, bool)`

GetSpecialOrderFlagOk returns a tuple with the SpecialOrderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialOrderFlag

`func (o *ProductItem) SetSpecialOrderFlag(v bool)`

SetSpecialOrderFlag sets SpecialOrderFlag field to given value.

### HasSpecialOrderFlag

`func (o *ProductItem) HasSpecialOrderFlag() bool`

HasSpecialOrderFlag returns a boolean if a field has been set.

### SetSpecialOrderFlagNil

`func (o *ProductItem) SetSpecialOrderFlagNil(b bool)`

 SetSpecialOrderFlagNil sets the value for SpecialOrderFlag to be an explicit nil

### UnsetSpecialOrderFlag
`func (o *ProductItem) UnsetSpecialOrderFlag()`

UnsetSpecialOrderFlag ensures that no value is present for SpecialOrderFlag, not even an explicit nil
### GetPhaseProductFlag

`func (o *ProductItem) GetPhaseProductFlag() bool`

GetPhaseProductFlag returns the PhaseProductFlag field if non-nil, zero value otherwise.

### GetPhaseProductFlagOk

`func (o *ProductItem) GetPhaseProductFlagOk() (*bool, bool)`

GetPhaseProductFlagOk returns a tuple with the PhaseProductFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseProductFlag

`func (o *ProductItem) SetPhaseProductFlag(v bool)`

SetPhaseProductFlag sets PhaseProductFlag field to given value.

### HasPhaseProductFlag

`func (o *ProductItem) HasPhaseProductFlag() bool`

HasPhaseProductFlag returns a boolean if a field has been set.

### SetPhaseProductFlagNil

`func (o *ProductItem) SetPhaseProductFlagNil(b bool)`

 SetPhaseProductFlagNil sets the value for PhaseProductFlag to be an explicit nil

### UnsetPhaseProductFlag
`func (o *ProductItem) UnsetPhaseProductFlag()`

UnsetPhaseProductFlag ensures that no value is present for PhaseProductFlag, not even an explicit nil
### GetCancelledFlag

`func (o *ProductItem) GetCancelledFlag() bool`

GetCancelledFlag returns the CancelledFlag field if non-nil, zero value otherwise.

### GetCancelledFlagOk

`func (o *ProductItem) GetCancelledFlagOk() (*bool, bool)`

GetCancelledFlagOk returns a tuple with the CancelledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledFlag

`func (o *ProductItem) SetCancelledFlag(v bool)`

SetCancelledFlag sets CancelledFlag field to given value.

### HasCancelledFlag

`func (o *ProductItem) HasCancelledFlag() bool`

HasCancelledFlag returns a boolean if a field has been set.

### SetCancelledFlagNil

`func (o *ProductItem) SetCancelledFlagNil(b bool)`

 SetCancelledFlagNil sets the value for CancelledFlag to be an explicit nil

### UnsetCancelledFlag
`func (o *ProductItem) UnsetCancelledFlag()`

UnsetCancelledFlag ensures that no value is present for CancelledFlag, not even an explicit nil
### GetQuantityCancelled

`func (o *ProductItem) GetQuantityCancelled() float64`

GetQuantityCancelled returns the QuantityCancelled field if non-nil, zero value otherwise.

### GetQuantityCancelledOk

`func (o *ProductItem) GetQuantityCancelledOk() (*float64, bool)`

GetQuantityCancelledOk returns a tuple with the QuantityCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityCancelled

`func (o *ProductItem) SetQuantityCancelled(v float64)`

SetQuantityCancelled sets QuantityCancelled field to given value.

### HasQuantityCancelled

`func (o *ProductItem) HasQuantityCancelled() bool`

HasQuantityCancelled returns a boolean if a field has been set.

### SetQuantityCancelledNil

`func (o *ProductItem) SetQuantityCancelledNil(b bool)`

 SetQuantityCancelledNil sets the value for QuantityCancelled to be an explicit nil

### UnsetQuantityCancelled
`func (o *ProductItem) UnsetQuantityCancelled()`

UnsetQuantityCancelled ensures that no value is present for QuantityCancelled, not even an explicit nil
### GetCancelledReason

`func (o *ProductItem) GetCancelledReason() string`

GetCancelledReason returns the CancelledReason field if non-nil, zero value otherwise.

### GetCancelledReasonOk

`func (o *ProductItem) GetCancelledReasonOk() (*string, bool)`

GetCancelledReasonOk returns a tuple with the CancelledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledReason

`func (o *ProductItem) SetCancelledReason(v string)`

SetCancelledReason sets CancelledReason field to given value.

### HasCancelledReason

`func (o *ProductItem) HasCancelledReason() bool`

HasCancelledReason returns a boolean if a field has been set.

### GetCustomerDescription

`func (o *ProductItem) GetCustomerDescription() string`

GetCustomerDescription returns the CustomerDescription field if non-nil, zero value otherwise.

### GetCustomerDescriptionOk

`func (o *ProductItem) GetCustomerDescriptionOk() (*string, bool)`

GetCustomerDescriptionOk returns a tuple with the CustomerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDescription

`func (o *ProductItem) SetCustomerDescription(v string)`

SetCustomerDescription sets CustomerDescription field to given value.

### HasCustomerDescription

`func (o *ProductItem) HasCustomerDescription() bool`

HasCustomerDescription returns a boolean if a field has been set.

### GetInternalNotes

`func (o *ProductItem) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *ProductItem) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *ProductItem) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *ProductItem) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetProductSuppliedFlag

`func (o *ProductItem) GetProductSuppliedFlag() bool`

GetProductSuppliedFlag returns the ProductSuppliedFlag field if non-nil, zero value otherwise.

### GetProductSuppliedFlagOk

`func (o *ProductItem) GetProductSuppliedFlagOk() (*bool, bool)`

GetProductSuppliedFlagOk returns a tuple with the ProductSuppliedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSuppliedFlag

`func (o *ProductItem) SetProductSuppliedFlag(v bool)`

SetProductSuppliedFlag sets ProductSuppliedFlag field to given value.

### HasProductSuppliedFlag

`func (o *ProductItem) HasProductSuppliedFlag() bool`

HasProductSuppliedFlag returns a boolean if a field has been set.

### SetProductSuppliedFlagNil

`func (o *ProductItem) SetProductSuppliedFlagNil(b bool)`

 SetProductSuppliedFlagNil sets the value for ProductSuppliedFlag to be an explicit nil

### UnsetProductSuppliedFlag
`func (o *ProductItem) UnsetProductSuppliedFlag()`

UnsetProductSuppliedFlag ensures that no value is present for ProductSuppliedFlag, not even an explicit nil
### GetSubContractorShipToId

`func (o *ProductItem) GetSubContractorShipToId() int32`

GetSubContractorShipToId returns the SubContractorShipToId field if non-nil, zero value otherwise.

### GetSubContractorShipToIdOk

`func (o *ProductItem) GetSubContractorShipToIdOk() (*int32, bool)`

GetSubContractorShipToIdOk returns a tuple with the SubContractorShipToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubContractorShipToId

`func (o *ProductItem) SetSubContractorShipToId(v int32)`

SetSubContractorShipToId sets SubContractorShipToId field to given value.

### HasSubContractorShipToId

`func (o *ProductItem) HasSubContractorShipToId() bool`

HasSubContractorShipToId returns a boolean if a field has been set.

### SetSubContractorShipToIdNil

`func (o *ProductItem) SetSubContractorShipToIdNil(b bool)`

 SetSubContractorShipToIdNil sets the value for SubContractorShipToId to be an explicit nil

### UnsetSubContractorShipToId
`func (o *ProductItem) UnsetSubContractorShipToId()`

UnsetSubContractorShipToId ensures that no value is present for SubContractorShipToId, not even an explicit nil
### GetSubContractorAmountLimit

`func (o *ProductItem) GetSubContractorAmountLimit() float64`

GetSubContractorAmountLimit returns the SubContractorAmountLimit field if non-nil, zero value otherwise.

### GetSubContractorAmountLimitOk

`func (o *ProductItem) GetSubContractorAmountLimitOk() (*float64, bool)`

GetSubContractorAmountLimitOk returns a tuple with the SubContractorAmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubContractorAmountLimit

`func (o *ProductItem) SetSubContractorAmountLimit(v float64)`

SetSubContractorAmountLimit sets SubContractorAmountLimit field to given value.

### HasSubContractorAmountLimit

`func (o *ProductItem) HasSubContractorAmountLimit() bool`

HasSubContractorAmountLimit returns a boolean if a field has been set.

### SetSubContractorAmountLimitNil

`func (o *ProductItem) SetSubContractorAmountLimitNil(b bool)`

 SetSubContractorAmountLimitNil sets the value for SubContractorAmountLimit to be an explicit nil

### UnsetSubContractorAmountLimit
`func (o *ProductItem) UnsetSubContractorAmountLimit()`

UnsetSubContractorAmountLimit ensures that no value is present for SubContractorAmountLimit, not even an explicit nil
### GetRecurring

`func (o *ProductItem) GetRecurring() ProductRecurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *ProductItem) GetRecurringOk() (*ProductRecurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *ProductItem) SetRecurring(v ProductRecurring)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *ProductItem) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetSla

`func (o *ProductItem) GetSla() SLAReference`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *ProductItem) GetSlaOk() (*SLAReference, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *ProductItem) SetSla(v SLAReference)`

SetSla sets Sla field to given value.

### HasSla

`func (o *ProductItem) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetEntityType

`func (o *ProductItem) GetEntityType() EntityTypeReference`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ProductItem) GetEntityTypeOk() (*EntityTypeReference, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ProductItem) SetEntityType(v EntityTypeReference)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *ProductItem) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetTicket

`func (o *ProductItem) GetTicket() TicketReference`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *ProductItem) GetTicketOk() (*TicketReference, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *ProductItem) SetTicket(v TicketReference)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *ProductItem) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetProject

`func (o *ProductItem) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProductItem) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProductItem) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProductItem) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPhase

`func (o *ProductItem) GetPhase() ProjectPhaseReference`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ProductItem) GetPhaseOk() (*ProjectPhaseReference, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ProductItem) SetPhase(v ProjectPhaseReference)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *ProductItem) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetSalesOrder

`func (o *ProductItem) GetSalesOrder() SalesOrderReference`

GetSalesOrder returns the SalesOrder field if non-nil, zero value otherwise.

### GetSalesOrderOk

`func (o *ProductItem) GetSalesOrderOk() (*SalesOrderReference, bool)`

GetSalesOrderOk returns a tuple with the SalesOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrder

`func (o *ProductItem) SetSalesOrder(v SalesOrderReference)`

SetSalesOrder sets SalesOrder field to given value.

### HasSalesOrder

`func (o *ProductItem) HasSalesOrder() bool`

HasSalesOrder returns a boolean if a field has been set.

### GetOpportunity

`func (o *ProductItem) GetOpportunity() OpportunityReference`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *ProductItem) GetOpportunityOk() (*OpportunityReference, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *ProductItem) SetOpportunity(v OpportunityReference)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *ProductItem) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### GetInvoice

`func (o *ProductItem) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *ProductItem) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *ProductItem) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *ProductItem) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetWarehouseId

`func (o *ProductItem) GetWarehouseId() int32`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ProductItem) GetWarehouseIdOk() (*int32, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ProductItem) SetWarehouseId(v int32)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ProductItem) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### SetWarehouseIdNil

`func (o *ProductItem) SetWarehouseIdNil(b bool)`

 SetWarehouseIdNil sets the value for WarehouseId to be an explicit nil

### UnsetWarehouseId
`func (o *ProductItem) UnsetWarehouseId()`

UnsetWarehouseId ensures that no value is present for WarehouseId, not even an explicit nil
### GetWarehouseIdObject

`func (o *ProductItem) GetWarehouseIdObject() WarehouseReference`

GetWarehouseIdObject returns the WarehouseIdObject field if non-nil, zero value otherwise.

### GetWarehouseIdObjectOk

`func (o *ProductItem) GetWarehouseIdObjectOk() (*WarehouseReference, bool)`

GetWarehouseIdObjectOk returns a tuple with the WarehouseIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseIdObject

`func (o *ProductItem) SetWarehouseIdObject(v WarehouseReference)`

SetWarehouseIdObject sets WarehouseIdObject field to given value.

### HasWarehouseIdObject

`func (o *ProductItem) HasWarehouseIdObject() bool`

HasWarehouseIdObject returns a boolean if a field has been set.

### GetWarehouseBinId

`func (o *ProductItem) GetWarehouseBinId() int32`

GetWarehouseBinId returns the WarehouseBinId field if non-nil, zero value otherwise.

### GetWarehouseBinIdOk

`func (o *ProductItem) GetWarehouseBinIdOk() (*int32, bool)`

GetWarehouseBinIdOk returns a tuple with the WarehouseBinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBinId

`func (o *ProductItem) SetWarehouseBinId(v int32)`

SetWarehouseBinId sets WarehouseBinId field to given value.

### HasWarehouseBinId

`func (o *ProductItem) HasWarehouseBinId() bool`

HasWarehouseBinId returns a boolean if a field has been set.

### SetWarehouseBinIdNil

`func (o *ProductItem) SetWarehouseBinIdNil(b bool)`

 SetWarehouseBinIdNil sets the value for WarehouseBinId to be an explicit nil

### UnsetWarehouseBinId
`func (o *ProductItem) UnsetWarehouseBinId()`

UnsetWarehouseBinId ensures that no value is present for WarehouseBinId, not even an explicit nil
### GetWarehouseBinIdObject

`func (o *ProductItem) GetWarehouseBinIdObject() WarehouseBinReference`

GetWarehouseBinIdObject returns the WarehouseBinIdObject field if non-nil, zero value otherwise.

### GetWarehouseBinIdObjectOk

`func (o *ProductItem) GetWarehouseBinIdObjectOk() (*WarehouseBinReference, bool)`

GetWarehouseBinIdObjectOk returns a tuple with the WarehouseBinIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBinIdObject

`func (o *ProductItem) SetWarehouseBinIdObject(v WarehouseBinReference)`

SetWarehouseBinIdObject sets WarehouseBinIdObject field to given value.

### HasWarehouseBinIdObject

`func (o *ProductItem) HasWarehouseBinIdObject() bool`

HasWarehouseBinIdObject returns a boolean if a field has been set.

### GetCalculatedPriceFlag

`func (o *ProductItem) GetCalculatedPriceFlag() bool`

GetCalculatedPriceFlag returns the CalculatedPriceFlag field if non-nil, zero value otherwise.

### GetCalculatedPriceFlagOk

`func (o *ProductItem) GetCalculatedPriceFlagOk() (*bool, bool)`

GetCalculatedPriceFlagOk returns a tuple with the CalculatedPriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedPriceFlag

`func (o *ProductItem) SetCalculatedPriceFlag(v bool)`

SetCalculatedPriceFlag sets CalculatedPriceFlag field to given value.

### HasCalculatedPriceFlag

`func (o *ProductItem) HasCalculatedPriceFlag() bool`

HasCalculatedPriceFlag returns a boolean if a field has been set.

### SetCalculatedPriceFlagNil

`func (o *ProductItem) SetCalculatedPriceFlagNil(b bool)`

 SetCalculatedPriceFlagNil sets the value for CalculatedPriceFlag to be an explicit nil

### UnsetCalculatedPriceFlag
`func (o *ProductItem) UnsetCalculatedPriceFlag()`

UnsetCalculatedPriceFlag ensures that no value is present for CalculatedPriceFlag, not even an explicit nil
### GetCalculatedCostFlag

`func (o *ProductItem) GetCalculatedCostFlag() bool`

GetCalculatedCostFlag returns the CalculatedCostFlag field if non-nil, zero value otherwise.

### GetCalculatedCostFlagOk

`func (o *ProductItem) GetCalculatedCostFlagOk() (*bool, bool)`

GetCalculatedCostFlagOk returns a tuple with the CalculatedCostFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedCostFlag

`func (o *ProductItem) SetCalculatedCostFlag(v bool)`

SetCalculatedCostFlag sets CalculatedCostFlag field to given value.

### HasCalculatedCostFlag

`func (o *ProductItem) HasCalculatedCostFlag() bool`

HasCalculatedCostFlag returns a boolean if a field has been set.

### SetCalculatedCostFlagNil

`func (o *ProductItem) SetCalculatedCostFlagNil(b bool)`

 SetCalculatedCostFlagNil sets the value for CalculatedCostFlag to be an explicit nil

### UnsetCalculatedCostFlag
`func (o *ProductItem) UnsetCalculatedCostFlag()`

UnsetCalculatedCostFlag ensures that no value is present for CalculatedCostFlag, not even an explicit nil
### GetForecastDetailId

`func (o *ProductItem) GetForecastDetailId() int32`

GetForecastDetailId returns the ForecastDetailId field if non-nil, zero value otherwise.

### GetForecastDetailIdOk

`func (o *ProductItem) GetForecastDetailIdOk() (*int32, bool)`

GetForecastDetailIdOk returns a tuple with the ForecastDetailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastDetailId

`func (o *ProductItem) SetForecastDetailId(v int32)`

SetForecastDetailId sets ForecastDetailId field to given value.

### HasForecastDetailId

`func (o *ProductItem) HasForecastDetailId() bool`

HasForecastDetailId returns a boolean if a field has been set.

### SetForecastDetailIdNil

`func (o *ProductItem) SetForecastDetailIdNil(b bool)`

 SetForecastDetailIdNil sets the value for ForecastDetailId to be an explicit nil

### UnsetForecastDetailId
`func (o *ProductItem) UnsetForecastDetailId()`

UnsetForecastDetailId ensures that no value is present for ForecastDetailId, not even an explicit nil
### GetCancelledBy

`func (o *ProductItem) GetCancelledBy() int32`

GetCancelledBy returns the CancelledBy field if non-nil, zero value otherwise.

### GetCancelledByOk

`func (o *ProductItem) GetCancelledByOk() (*int32, bool)`

GetCancelledByOk returns a tuple with the CancelledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledBy

`func (o *ProductItem) SetCancelledBy(v int32)`

SetCancelledBy sets CancelledBy field to given value.

### HasCancelledBy

`func (o *ProductItem) HasCancelledBy() bool`

HasCancelledBy returns a boolean if a field has been set.

### SetCancelledByNil

`func (o *ProductItem) SetCancelledByNil(b bool)`

 SetCancelledByNil sets the value for CancelledBy to be an explicit nil

### UnsetCancelledBy
`func (o *ProductItem) UnsetCancelledBy()`

UnsetCancelledBy ensures that no value is present for CancelledBy, not even an explicit nil
### GetCancelledDate

`func (o *ProductItem) GetCancelledDate() time.Time`

GetCancelledDate returns the CancelledDate field if non-nil, zero value otherwise.

### GetCancelledDateOk

`func (o *ProductItem) GetCancelledDateOk() (*time.Time, bool)`

GetCancelledDateOk returns a tuple with the CancelledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledDate

`func (o *ProductItem) SetCancelledDate(v time.Time)`

SetCancelledDate sets CancelledDate field to given value.

### HasCancelledDate

`func (o *ProductItem) HasCancelledDate() bool`

HasCancelledDate returns a boolean if a field has been set.

### GetWarehouse

`func (o *ProductItem) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *ProductItem) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *ProductItem) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *ProductItem) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetWarehouseBin

`func (o *ProductItem) GetWarehouseBin() string`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *ProductItem) GetWarehouseBinOk() (*string, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *ProductItem) SetWarehouseBin(v string)`

SetWarehouseBin sets WarehouseBin field to given value.

### HasWarehouseBin

`func (o *ProductItem) HasWarehouseBin() bool`

HasWarehouseBin returns a boolean if a field has been set.

### GetPurchaseDate

`func (o *ProductItem) GetPurchaseDate() time.Time`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *ProductItem) GetPurchaseDateOk() (*time.Time, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *ProductItem) SetPurchaseDate(v time.Time)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *ProductItem) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### GetTaxCode

`func (o *ProductItem) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *ProductItem) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *ProductItem) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *ProductItem) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetIntegrationXRef

`func (o *ProductItem) GetIntegrationXRef() string`

GetIntegrationXRef returns the IntegrationXRef field if non-nil, zero value otherwise.

### GetIntegrationXRefOk

`func (o *ProductItem) GetIntegrationXRefOk() (*string, bool)`

GetIntegrationXRefOk returns a tuple with the IntegrationXRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationXRef

`func (o *ProductItem) SetIntegrationXRef(v string)`

SetIntegrationXRef sets IntegrationXRef field to given value.

### HasIntegrationXRef

`func (o *ProductItem) HasIntegrationXRef() bool`

HasIntegrationXRef returns a boolean if a field has been set.

### GetListPrice

`func (o *ProductItem) GetListPrice() float64`

GetListPrice returns the ListPrice field if non-nil, zero value otherwise.

### GetListPriceOk

`func (o *ProductItem) GetListPriceOk() (*float64, bool)`

GetListPriceOk returns a tuple with the ListPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListPrice

`func (o *ProductItem) SetListPrice(v float64)`

SetListPrice sets ListPrice field to given value.

### HasListPrice

`func (o *ProductItem) HasListPrice() bool`

HasListPrice returns a boolean if a field has been set.

### SetListPriceNil

`func (o *ProductItem) SetListPriceNil(b bool)`

 SetListPriceNil sets the value for ListPrice to be an explicit nil

### UnsetListPrice
`func (o *ProductItem) UnsetListPrice()`

UnsetListPrice ensures that no value is present for ListPrice, not even an explicit nil
### GetSerialNumberIds

`func (o *ProductItem) GetSerialNumberIds() []int32`

GetSerialNumberIds returns the SerialNumberIds field if non-nil, zero value otherwise.

### GetSerialNumberIdsOk

`func (o *ProductItem) GetSerialNumberIdsOk() (*[]int32, bool)`

GetSerialNumberIdsOk returns a tuple with the SerialNumberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumberIds

`func (o *ProductItem) SetSerialNumberIds(v []int32)`

SetSerialNumberIds sets SerialNumberIds field to given value.

### HasSerialNumberIds

`func (o *ProductItem) HasSerialNumberIds() bool`

HasSerialNumberIds returns a boolean if a field has been set.

### GetSerialNumbers

`func (o *ProductItem) GetSerialNumbers() []string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *ProductItem) GetSerialNumbersOk() (*[]string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *ProductItem) SetSerialNumbers(v []string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *ProductItem) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.

### GetCompany

`func (o *ProductItem) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ProductItem) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ProductItem) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ProductItem) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetForecastStatus

`func (o *ProductItem) GetForecastStatus() OpportunityStatusReference`

GetForecastStatus returns the ForecastStatus field if non-nil, zero value otherwise.

### GetForecastStatusOk

`func (o *ProductItem) GetForecastStatusOk() (*OpportunityStatusReference, bool)`

GetForecastStatusOk returns a tuple with the ForecastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastStatus

`func (o *ProductItem) SetForecastStatus(v OpportunityStatusReference)`

SetForecastStatus sets ForecastStatus field to given value.

### HasForecastStatus

`func (o *ProductItem) HasForecastStatus() bool`

HasForecastStatus returns a boolean if a field has been set.

### GetProductClass

`func (o *ProductItem) GetProductClass() string`

GetProductClass returns the ProductClass field if non-nil, zero value otherwise.

### GetProductClassOk

`func (o *ProductItem) GetProductClassOk() (*string, bool)`

GetProductClassOk returns a tuple with the ProductClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductClass

`func (o *ProductItem) SetProductClass(v string)`

SetProductClass sets ProductClass field to given value.

### HasProductClass

`func (o *ProductItem) HasProductClass() bool`

HasProductClass returns a boolean if a field has been set.

### SetProductClassNil

`func (o *ProductItem) SetProductClassNil(b bool)`

 SetProductClassNil sets the value for ProductClass to be an explicit nil

### UnsetProductClass
`func (o *ProductItem) UnsetProductClass()`

UnsetProductClass ensures that no value is present for ProductClass, not even an explicit nil
### GetNeedToPurchaseFlag

`func (o *ProductItem) GetNeedToPurchaseFlag() bool`

GetNeedToPurchaseFlag returns the NeedToPurchaseFlag field if non-nil, zero value otherwise.

### GetNeedToPurchaseFlagOk

`func (o *ProductItem) GetNeedToPurchaseFlagOk() (*bool, bool)`

GetNeedToPurchaseFlagOk returns a tuple with the NeedToPurchaseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedToPurchaseFlag

`func (o *ProductItem) SetNeedToPurchaseFlag(v bool)`

SetNeedToPurchaseFlag sets NeedToPurchaseFlag field to given value.

### HasNeedToPurchaseFlag

`func (o *ProductItem) HasNeedToPurchaseFlag() bool`

HasNeedToPurchaseFlag returns a boolean if a field has been set.

### SetNeedToPurchaseFlagNil

`func (o *ProductItem) SetNeedToPurchaseFlagNil(b bool)`

 SetNeedToPurchaseFlagNil sets the value for NeedToPurchaseFlag to be an explicit nil

### UnsetNeedToPurchaseFlag
`func (o *ProductItem) UnsetNeedToPurchaseFlag()`

UnsetNeedToPurchaseFlag ensures that no value is present for NeedToPurchaseFlag, not even an explicit nil
### GetNeedToOrderQuantity

`func (o *ProductItem) GetNeedToOrderQuantity() int32`

GetNeedToOrderQuantity returns the NeedToOrderQuantity field if non-nil, zero value otherwise.

### GetNeedToOrderQuantityOk

`func (o *ProductItem) GetNeedToOrderQuantityOk() (*int32, bool)`

GetNeedToOrderQuantityOk returns a tuple with the NeedToOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedToOrderQuantity

`func (o *ProductItem) SetNeedToOrderQuantity(v int32)`

SetNeedToOrderQuantity sets NeedToOrderQuantity field to given value.

### HasNeedToOrderQuantity

`func (o *ProductItem) HasNeedToOrderQuantity() bool`

HasNeedToOrderQuantity returns a boolean if a field has been set.

### SetNeedToOrderQuantityNil

`func (o *ProductItem) SetNeedToOrderQuantityNil(b bool)`

 SetNeedToOrderQuantityNil sets the value for NeedToOrderQuantity to be an explicit nil

### UnsetNeedToOrderQuantity
`func (o *ProductItem) UnsetNeedToOrderQuantity()`

UnsetNeedToOrderQuantity ensures that no value is present for NeedToOrderQuantity, not even an explicit nil
### GetMinimumStockFlag

`func (o *ProductItem) GetMinimumStockFlag() bool`

GetMinimumStockFlag returns the MinimumStockFlag field if non-nil, zero value otherwise.

### GetMinimumStockFlagOk

`func (o *ProductItem) GetMinimumStockFlagOk() (*bool, bool)`

GetMinimumStockFlagOk returns a tuple with the MinimumStockFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumStockFlag

`func (o *ProductItem) SetMinimumStockFlag(v bool)`

SetMinimumStockFlag sets MinimumStockFlag field to given value.

### HasMinimumStockFlag

`func (o *ProductItem) HasMinimumStockFlag() bool`

HasMinimumStockFlag returns a boolean if a field has been set.

### SetMinimumStockFlagNil

`func (o *ProductItem) SetMinimumStockFlagNil(b bool)`

 SetMinimumStockFlagNil sets the value for MinimumStockFlag to be an explicit nil

### UnsetMinimumStockFlag
`func (o *ProductItem) UnsetMinimumStockFlag()`

UnsetMinimumStockFlag ensures that no value is present for MinimumStockFlag, not even an explicit nil
### GetShipSet

`func (o *ProductItem) GetShipSet() string`

GetShipSet returns the ShipSet field if non-nil, zero value otherwise.

### GetShipSetOk

`func (o *ProductItem) GetShipSetOk() (*string, bool)`

GetShipSetOk returns a tuple with the ShipSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipSet

`func (o *ProductItem) SetShipSet(v string)`

SetShipSet sets ShipSet field to given value.

### HasShipSet

`func (o *ProductItem) HasShipSet() bool`

HasShipSet returns a boolean if a field has been set.

### GetCalculatedPrice

`func (o *ProductItem) GetCalculatedPrice() float64`

GetCalculatedPrice returns the CalculatedPrice field if non-nil, zero value otherwise.

### GetCalculatedPriceOk

`func (o *ProductItem) GetCalculatedPriceOk() (*float64, bool)`

GetCalculatedPriceOk returns a tuple with the CalculatedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedPrice

`func (o *ProductItem) SetCalculatedPrice(v float64)`

SetCalculatedPrice sets CalculatedPrice field to given value.

### HasCalculatedPrice

`func (o *ProductItem) HasCalculatedPrice() bool`

HasCalculatedPrice returns a boolean if a field has been set.

### SetCalculatedPriceNil

`func (o *ProductItem) SetCalculatedPriceNil(b bool)`

 SetCalculatedPriceNil sets the value for CalculatedPrice to be an explicit nil

### UnsetCalculatedPrice
`func (o *ProductItem) UnsetCalculatedPrice()`

UnsetCalculatedPrice ensures that no value is present for CalculatedPrice, not even an explicit nil
### GetCalculatedCost

`func (o *ProductItem) GetCalculatedCost() float64`

GetCalculatedCost returns the CalculatedCost field if non-nil, zero value otherwise.

### GetCalculatedCostOk

`func (o *ProductItem) GetCalculatedCostOk() (*float64, bool)`

GetCalculatedCostOk returns a tuple with the CalculatedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedCost

`func (o *ProductItem) SetCalculatedCost(v float64)`

SetCalculatedCost sets CalculatedCost field to given value.

### HasCalculatedCost

`func (o *ProductItem) HasCalculatedCost() bool`

HasCalculatedCost returns a boolean if a field has been set.

### SetCalculatedCostNil

`func (o *ProductItem) SetCalculatedCostNil(b bool)`

 SetCalculatedCostNil sets the value for CalculatedCost to be an explicit nil

### UnsetCalculatedCost
`func (o *ProductItem) UnsetCalculatedCost()`

UnsetCalculatedCost ensures that no value is present for CalculatedCost, not even an explicit nil
### GetInvoiceGrouping

`func (o *ProductItem) GetInvoiceGrouping() InvoiceGroupingReference`

GetInvoiceGrouping returns the InvoiceGrouping field if non-nil, zero value otherwise.

### GetInvoiceGroupingOk

`func (o *ProductItem) GetInvoiceGroupingOk() (*InvoiceGroupingReference, bool)`

GetInvoiceGroupingOk returns a tuple with the InvoiceGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceGrouping

`func (o *ProductItem) SetInvoiceGrouping(v InvoiceGroupingReference)`

SetInvoiceGrouping sets InvoiceGrouping field to given value.

### HasInvoiceGrouping

`func (o *ProductItem) HasInvoiceGrouping() bool`

HasInvoiceGrouping returns a boolean if a field has been set.

### GetPoApprovedFlag

`func (o *ProductItem) GetPoApprovedFlag() bool`

GetPoApprovedFlag returns the PoApprovedFlag field if non-nil, zero value otherwise.

### GetPoApprovedFlagOk

`func (o *ProductItem) GetPoApprovedFlagOk() (*bool, bool)`

GetPoApprovedFlagOk returns a tuple with the PoApprovedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoApprovedFlag

`func (o *ProductItem) SetPoApprovedFlag(v bool)`

SetPoApprovedFlag sets PoApprovedFlag field to given value.

### HasPoApprovedFlag

`func (o *ProductItem) HasPoApprovedFlag() bool`

HasPoApprovedFlag returns a boolean if a field has been set.

### SetPoApprovedFlagNil

`func (o *ProductItem) SetPoApprovedFlagNil(b bool)`

 SetPoApprovedFlagNil sets the value for PoApprovedFlag to be an explicit nil

### UnsetPoApprovedFlag
`func (o *ProductItem) UnsetPoApprovedFlag()`

UnsetPoApprovedFlag ensures that no value is present for PoApprovedFlag, not even an explicit nil
### GetUom

`func (o *ProductItem) GetUom() string`

GetUom returns the Uom field if non-nil, zero value otherwise.

### GetUomOk

`func (o *ProductItem) GetUomOk() (*string, bool)`

GetUomOk returns a tuple with the Uom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUom

`func (o *ProductItem) SetUom(v string)`

SetUom sets Uom field to given value.

### HasUom

`func (o *ProductItem) HasUom() bool`

HasUom returns a boolean if a field has been set.

### GetAddComponentsFlag

`func (o *ProductItem) GetAddComponentsFlag() bool`

GetAddComponentsFlag returns the AddComponentsFlag field if non-nil, zero value otherwise.

### GetAddComponentsFlagOk

`func (o *ProductItem) GetAddComponentsFlagOk() (*bool, bool)`

GetAddComponentsFlagOk returns a tuple with the AddComponentsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddComponentsFlag

`func (o *ProductItem) SetAddComponentsFlag(v bool)`

SetAddComponentsFlag sets AddComponentsFlag field to given value.

### HasAddComponentsFlag

`func (o *ProductItem) HasAddComponentsFlag() bool`

HasAddComponentsFlag returns a boolean if a field has been set.

### SetAddComponentsFlagNil

`func (o *ProductItem) SetAddComponentsFlagNil(b bool)`

 SetAddComponentsFlagNil sets the value for AddComponentsFlag to be an explicit nil

### UnsetAddComponentsFlag
`func (o *ProductItem) UnsetAddComponentsFlag()`

UnsetAddComponentsFlag ensures that no value is present for AddComponentsFlag, not even an explicit nil
### GetIgnorePricingSchedulesFlag

`func (o *ProductItem) GetIgnorePricingSchedulesFlag() bool`

GetIgnorePricingSchedulesFlag returns the IgnorePricingSchedulesFlag field if non-nil, zero value otherwise.

### GetIgnorePricingSchedulesFlagOk

`func (o *ProductItem) GetIgnorePricingSchedulesFlagOk() (*bool, bool)`

GetIgnorePricingSchedulesFlagOk returns a tuple with the IgnorePricingSchedulesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnorePricingSchedulesFlag

`func (o *ProductItem) SetIgnorePricingSchedulesFlag(v bool)`

SetIgnorePricingSchedulesFlag sets IgnorePricingSchedulesFlag field to given value.

### HasIgnorePricingSchedulesFlag

`func (o *ProductItem) HasIgnorePricingSchedulesFlag() bool`

HasIgnorePricingSchedulesFlag returns a boolean if a field has been set.

### SetIgnorePricingSchedulesFlagNil

`func (o *ProductItem) SetIgnorePricingSchedulesFlagNil(b bool)`

 SetIgnorePricingSchedulesFlagNil sets the value for IgnorePricingSchedulesFlag to be an explicit nil

### UnsetIgnorePricingSchedulesFlag
`func (o *ProductItem) UnsetIgnorePricingSchedulesFlag()`

UnsetIgnorePricingSchedulesFlag ensures that no value is present for IgnorePricingSchedulesFlag, not even an explicit nil
### GetAsioSubscriptionsID

`func (o *ProductItem) GetAsioSubscriptionsID() string`

GetAsioSubscriptionsID returns the AsioSubscriptionsID field if non-nil, zero value otherwise.

### GetAsioSubscriptionsIDOk

`func (o *ProductItem) GetAsioSubscriptionsIDOk() (*string, bool)`

GetAsioSubscriptionsIDOk returns a tuple with the AsioSubscriptionsID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsioSubscriptionsID

`func (o *ProductItem) SetAsioSubscriptionsID(v string)`

SetAsioSubscriptionsID sets AsioSubscriptionsID field to given value.

### HasAsioSubscriptionsID

`func (o *ProductItem) HasAsioSubscriptionsID() bool`

HasAsioSubscriptionsID returns a boolean if a field has been set.

### SetAsioSubscriptionsIDNil

`func (o *ProductItem) SetAsioSubscriptionsIDNil(b bool)`

 SetAsioSubscriptionsIDNil sets the value for AsioSubscriptionsID to be an explicit nil

### UnsetAsioSubscriptionsID
`func (o *ProductItem) UnsetAsioSubscriptionsID()`

UnsetAsioSubscriptionsID ensures that no value is present for AsioSubscriptionsID, not even an explicit nil
### GetInfo

`func (o *ProductItem) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProductItem) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProductItem) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProductItem) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetBypassForecastUpdate

`func (o *ProductItem) GetBypassForecastUpdate() bool`

GetBypassForecastUpdate returns the BypassForecastUpdate field if non-nil, zero value otherwise.

### GetBypassForecastUpdateOk

`func (o *ProductItem) GetBypassForecastUpdateOk() (*bool, bool)`

GetBypassForecastUpdateOk returns a tuple with the BypassForecastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassForecastUpdate

`func (o *ProductItem) SetBypassForecastUpdate(v bool)`

SetBypassForecastUpdate sets BypassForecastUpdate field to given value.

### HasBypassForecastUpdate

`func (o *ProductItem) HasBypassForecastUpdate() bool`

HasBypassForecastUpdate returns a boolean if a field has been set.

### SetBypassForecastUpdateNil

`func (o *ProductItem) SetBypassForecastUpdateNil(b bool)`

 SetBypassForecastUpdateNil sets the value for BypassForecastUpdate to be an explicit nil

### UnsetBypassForecastUpdate
`func (o *ProductItem) UnsetBypassForecastUpdate()`

UnsetBypassForecastUpdate ensures that no value is present for BypassForecastUpdate, not even an explicit nil
### GetCustomFields

`func (o *ProductItem) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProductItem) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProductItem) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProductItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


