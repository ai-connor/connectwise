# UnpostedProcurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**UnpostedProductId** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**DepartmentId** | Pointer to **NullableInt32** |  | [optional] 
**ProcurementType** | Pointer to **NullableString** |  | [optional] 
**PurchaseOrder** | Pointer to [**PurchaseOrderReference**](PurchaseOrderReference.md) |  | [optional] 
**PurchaseDate** | Pointer to **string** |  | [optional] 
**TrackingNumber** | Pointer to **string** |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**AvalaraTaxFlag** | Pointer to **NullableBool** | Used to determine if Avalara tax is enabled. | [optional] 
**ItemTaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**PurchaseOrderTaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**StateTaxFlag** | Pointer to **NullableBool** | Set to true if transaction is taxable at the state level. | [optional] 
**StateTaxXref** | Pointer to **string** |  | [optional] 
**StateTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**CountyTaxFlag** | Pointer to **NullableBool** | Set to true if transaction is taxable at the county level. | [optional] 
**CountyTaxXref** | Pointer to **string** |  | [optional] 
**CountyTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**CityTaxFlag** | Pointer to **NullableBool** | Set to true if transaction is taxable at the city level. | [optional] 
**CityTaxXref** | Pointer to **string** |  | [optional] 
**CityTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**CountryTaxFlag** | Pointer to **NullableBool** | Set to true if transaction is taxable at the country level. | [optional] 
**CountryTaxXref** | Pointer to **string** |  | [optional] 
**CountryTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**CompositeTaxFlag** | Pointer to **NullableBool** | Set to true if transaction is taxable at the composite level. | [optional] 
**CompositeTaxXref** | Pointer to **string** |  | [optional] 
**CompositeTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**LevelSixTaxFlag** | Pointer to **NullableBool** | Set to true if transaction is taxable at level six. | [optional] 
**LevelSixTaxXref** | Pointer to **string** |  | [optional] 
**LevelSixTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**TaxTotal** | Pointer to **NullableFloat64** |  | [optional] 
**Customer** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Vendor** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**VendorAccountNumber** | Pointer to **string** |  | [optional] 
**VendorInvoiceNumber** | Pointer to **string** |  | [optional] 
**VendorInvoiceDate** | Pointer to **string** |  | [optional] 
**TaxFreightFlag** | Pointer to **NullableBool** |  | [optional] 
**FreightTaxTotal** | Pointer to **NullableFloat64** |  | [optional] 
**FreightCost** | Pointer to **NullableFloat64** |  | [optional] 
**DateClosed** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUnpostedProcurement

`func NewUnpostedProcurement() *UnpostedProcurement`

NewUnpostedProcurement instantiates a new UnpostedProcurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnpostedProcurementWithDefaults

`func NewUnpostedProcurementWithDefaults() *UnpostedProcurement`

NewUnpostedProcurementWithDefaults instantiates a new UnpostedProcurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnpostedProcurement) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnpostedProcurement) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnpostedProcurement) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UnpostedProcurement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *UnpostedProcurement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnpostedProcurement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnpostedProcurement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnpostedProcurement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnpostedProductId

`func (o *UnpostedProcurement) GetUnpostedProductId() string`

GetUnpostedProductId returns the UnpostedProductId field if non-nil, zero value otherwise.

### GetUnpostedProductIdOk

`func (o *UnpostedProcurement) GetUnpostedProductIdOk() (*string, bool)`

GetUnpostedProductIdOk returns a tuple with the UnpostedProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpostedProductId

`func (o *UnpostedProcurement) SetUnpostedProductId(v string)`

SetUnpostedProductId sets UnpostedProductId field to given value.

### HasUnpostedProductId

`func (o *UnpostedProcurement) HasUnpostedProductId() bool`

HasUnpostedProductId returns a boolean if a field has been set.

### GetLocationId

`func (o *UnpostedProcurement) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *UnpostedProcurement) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *UnpostedProcurement) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *UnpostedProcurement) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *UnpostedProcurement) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *UnpostedProcurement) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetDepartmentId

`func (o *UnpostedProcurement) GetDepartmentId() int32`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *UnpostedProcurement) GetDepartmentIdOk() (*int32, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *UnpostedProcurement) SetDepartmentId(v int32)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *UnpostedProcurement) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### SetDepartmentIdNil

`func (o *UnpostedProcurement) SetDepartmentIdNil(b bool)`

 SetDepartmentIdNil sets the value for DepartmentId to be an explicit nil

### UnsetDepartmentId
`func (o *UnpostedProcurement) UnsetDepartmentId()`

UnsetDepartmentId ensures that no value is present for DepartmentId, not even an explicit nil
### GetProcurementType

`func (o *UnpostedProcurement) GetProcurementType() string`

GetProcurementType returns the ProcurementType field if non-nil, zero value otherwise.

### GetProcurementTypeOk

`func (o *UnpostedProcurement) GetProcurementTypeOk() (*string, bool)`

GetProcurementTypeOk returns a tuple with the ProcurementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcurementType

`func (o *UnpostedProcurement) SetProcurementType(v string)`

SetProcurementType sets ProcurementType field to given value.

### HasProcurementType

`func (o *UnpostedProcurement) HasProcurementType() bool`

HasProcurementType returns a boolean if a field has been set.

### SetProcurementTypeNil

`func (o *UnpostedProcurement) SetProcurementTypeNil(b bool)`

 SetProcurementTypeNil sets the value for ProcurementType to be an explicit nil

### UnsetProcurementType
`func (o *UnpostedProcurement) UnsetProcurementType()`

UnsetProcurementType ensures that no value is present for ProcurementType, not even an explicit nil
### GetPurchaseOrder

`func (o *UnpostedProcurement) GetPurchaseOrder() PurchaseOrderReference`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *UnpostedProcurement) GetPurchaseOrderOk() (*PurchaseOrderReference, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *UnpostedProcurement) SetPurchaseOrder(v PurchaseOrderReference)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *UnpostedProcurement) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.

### GetPurchaseDate

`func (o *UnpostedProcurement) GetPurchaseDate() string`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *UnpostedProcurement) GetPurchaseDateOk() (*string, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *UnpostedProcurement) SetPurchaseDate(v string)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *UnpostedProcurement) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *UnpostedProcurement) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *UnpostedProcurement) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *UnpostedProcurement) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *UnpostedProcurement) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetBillingTerms

`func (o *UnpostedProcurement) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *UnpostedProcurement) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *UnpostedProcurement) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *UnpostedProcurement) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetCurrency

`func (o *UnpostedProcurement) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnpostedProcurement) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnpostedProcurement) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnpostedProcurement) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotal

`func (o *UnpostedProcurement) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnpostedProcurement) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnpostedProcurement) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnpostedProcurement) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *UnpostedProcurement) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *UnpostedProcurement) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetTaxCode

`func (o *UnpostedProcurement) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *UnpostedProcurement) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *UnpostedProcurement) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *UnpostedProcurement) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetAvalaraTaxFlag

`func (o *UnpostedProcurement) GetAvalaraTaxFlag() bool`

GetAvalaraTaxFlag returns the AvalaraTaxFlag field if non-nil, zero value otherwise.

### GetAvalaraTaxFlagOk

`func (o *UnpostedProcurement) GetAvalaraTaxFlagOk() (*bool, bool)`

GetAvalaraTaxFlagOk returns a tuple with the AvalaraTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvalaraTaxFlag

`func (o *UnpostedProcurement) SetAvalaraTaxFlag(v bool)`

SetAvalaraTaxFlag sets AvalaraTaxFlag field to given value.

### HasAvalaraTaxFlag

`func (o *UnpostedProcurement) HasAvalaraTaxFlag() bool`

HasAvalaraTaxFlag returns a boolean if a field has been set.

### SetAvalaraTaxFlagNil

`func (o *UnpostedProcurement) SetAvalaraTaxFlagNil(b bool)`

 SetAvalaraTaxFlagNil sets the value for AvalaraTaxFlag to be an explicit nil

### UnsetAvalaraTaxFlag
`func (o *UnpostedProcurement) UnsetAvalaraTaxFlag()`

UnsetAvalaraTaxFlag ensures that no value is present for AvalaraTaxFlag, not even an explicit nil
### GetItemTaxableFlag

`func (o *UnpostedProcurement) GetItemTaxableFlag() bool`

GetItemTaxableFlag returns the ItemTaxableFlag field if non-nil, zero value otherwise.

### GetItemTaxableFlagOk

`func (o *UnpostedProcurement) GetItemTaxableFlagOk() (*bool, bool)`

GetItemTaxableFlagOk returns a tuple with the ItemTaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTaxableFlag

`func (o *UnpostedProcurement) SetItemTaxableFlag(v bool)`

SetItemTaxableFlag sets ItemTaxableFlag field to given value.

### HasItemTaxableFlag

`func (o *UnpostedProcurement) HasItemTaxableFlag() bool`

HasItemTaxableFlag returns a boolean if a field has been set.

### SetItemTaxableFlagNil

`func (o *UnpostedProcurement) SetItemTaxableFlagNil(b bool)`

 SetItemTaxableFlagNil sets the value for ItemTaxableFlag to be an explicit nil

### UnsetItemTaxableFlag
`func (o *UnpostedProcurement) UnsetItemTaxableFlag()`

UnsetItemTaxableFlag ensures that no value is present for ItemTaxableFlag, not even an explicit nil
### GetPurchaseOrderTaxableFlag

`func (o *UnpostedProcurement) GetPurchaseOrderTaxableFlag() bool`

GetPurchaseOrderTaxableFlag returns the PurchaseOrderTaxableFlag field if non-nil, zero value otherwise.

### GetPurchaseOrderTaxableFlagOk

`func (o *UnpostedProcurement) GetPurchaseOrderTaxableFlagOk() (*bool, bool)`

GetPurchaseOrderTaxableFlagOk returns a tuple with the PurchaseOrderTaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderTaxableFlag

`func (o *UnpostedProcurement) SetPurchaseOrderTaxableFlag(v bool)`

SetPurchaseOrderTaxableFlag sets PurchaseOrderTaxableFlag field to given value.

### HasPurchaseOrderTaxableFlag

`func (o *UnpostedProcurement) HasPurchaseOrderTaxableFlag() bool`

HasPurchaseOrderTaxableFlag returns a boolean if a field has been set.

### SetPurchaseOrderTaxableFlagNil

`func (o *UnpostedProcurement) SetPurchaseOrderTaxableFlagNil(b bool)`

 SetPurchaseOrderTaxableFlagNil sets the value for PurchaseOrderTaxableFlag to be an explicit nil

### UnsetPurchaseOrderTaxableFlag
`func (o *UnpostedProcurement) UnsetPurchaseOrderTaxableFlag()`

UnsetPurchaseOrderTaxableFlag ensures that no value is present for PurchaseOrderTaxableFlag, not even an explicit nil
### GetStateTaxFlag

`func (o *UnpostedProcurement) GetStateTaxFlag() bool`

GetStateTaxFlag returns the StateTaxFlag field if non-nil, zero value otherwise.

### GetStateTaxFlagOk

`func (o *UnpostedProcurement) GetStateTaxFlagOk() (*bool, bool)`

GetStateTaxFlagOk returns a tuple with the StateTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxFlag

`func (o *UnpostedProcurement) SetStateTaxFlag(v bool)`

SetStateTaxFlag sets StateTaxFlag field to given value.

### HasStateTaxFlag

`func (o *UnpostedProcurement) HasStateTaxFlag() bool`

HasStateTaxFlag returns a boolean if a field has been set.

### SetStateTaxFlagNil

`func (o *UnpostedProcurement) SetStateTaxFlagNil(b bool)`

 SetStateTaxFlagNil sets the value for StateTaxFlag to be an explicit nil

### UnsetStateTaxFlag
`func (o *UnpostedProcurement) UnsetStateTaxFlag()`

UnsetStateTaxFlag ensures that no value is present for StateTaxFlag, not even an explicit nil
### GetStateTaxXref

`func (o *UnpostedProcurement) GetStateTaxXref() string`

GetStateTaxXref returns the StateTaxXref field if non-nil, zero value otherwise.

### GetStateTaxXrefOk

`func (o *UnpostedProcurement) GetStateTaxXrefOk() (*string, bool)`

GetStateTaxXrefOk returns a tuple with the StateTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxXref

`func (o *UnpostedProcurement) SetStateTaxXref(v string)`

SetStateTaxXref sets StateTaxXref field to given value.

### HasStateTaxXref

`func (o *UnpostedProcurement) HasStateTaxXref() bool`

HasStateTaxXref returns a boolean if a field has been set.

### GetStateTaxAmount

`func (o *UnpostedProcurement) GetStateTaxAmount() float64`

GetStateTaxAmount returns the StateTaxAmount field if non-nil, zero value otherwise.

### GetStateTaxAmountOk

`func (o *UnpostedProcurement) GetStateTaxAmountOk() (*float64, bool)`

GetStateTaxAmountOk returns a tuple with the StateTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxAmount

`func (o *UnpostedProcurement) SetStateTaxAmount(v float64)`

SetStateTaxAmount sets StateTaxAmount field to given value.

### HasStateTaxAmount

`func (o *UnpostedProcurement) HasStateTaxAmount() bool`

HasStateTaxAmount returns a boolean if a field has been set.

### SetStateTaxAmountNil

`func (o *UnpostedProcurement) SetStateTaxAmountNil(b bool)`

 SetStateTaxAmountNil sets the value for StateTaxAmount to be an explicit nil

### UnsetStateTaxAmount
`func (o *UnpostedProcurement) UnsetStateTaxAmount()`

UnsetStateTaxAmount ensures that no value is present for StateTaxAmount, not even an explicit nil
### GetCountyTaxFlag

`func (o *UnpostedProcurement) GetCountyTaxFlag() bool`

GetCountyTaxFlag returns the CountyTaxFlag field if non-nil, zero value otherwise.

### GetCountyTaxFlagOk

`func (o *UnpostedProcurement) GetCountyTaxFlagOk() (*bool, bool)`

GetCountyTaxFlagOk returns a tuple with the CountyTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxFlag

`func (o *UnpostedProcurement) SetCountyTaxFlag(v bool)`

SetCountyTaxFlag sets CountyTaxFlag field to given value.

### HasCountyTaxFlag

`func (o *UnpostedProcurement) HasCountyTaxFlag() bool`

HasCountyTaxFlag returns a boolean if a field has been set.

### SetCountyTaxFlagNil

`func (o *UnpostedProcurement) SetCountyTaxFlagNil(b bool)`

 SetCountyTaxFlagNil sets the value for CountyTaxFlag to be an explicit nil

### UnsetCountyTaxFlag
`func (o *UnpostedProcurement) UnsetCountyTaxFlag()`

UnsetCountyTaxFlag ensures that no value is present for CountyTaxFlag, not even an explicit nil
### GetCountyTaxXref

`func (o *UnpostedProcurement) GetCountyTaxXref() string`

GetCountyTaxXref returns the CountyTaxXref field if non-nil, zero value otherwise.

### GetCountyTaxXrefOk

`func (o *UnpostedProcurement) GetCountyTaxXrefOk() (*string, bool)`

GetCountyTaxXrefOk returns a tuple with the CountyTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxXref

`func (o *UnpostedProcurement) SetCountyTaxXref(v string)`

SetCountyTaxXref sets CountyTaxXref field to given value.

### HasCountyTaxXref

`func (o *UnpostedProcurement) HasCountyTaxXref() bool`

HasCountyTaxXref returns a boolean if a field has been set.

### GetCountyTaxAmount

`func (o *UnpostedProcurement) GetCountyTaxAmount() float64`

GetCountyTaxAmount returns the CountyTaxAmount field if non-nil, zero value otherwise.

### GetCountyTaxAmountOk

`func (o *UnpostedProcurement) GetCountyTaxAmountOk() (*float64, bool)`

GetCountyTaxAmountOk returns a tuple with the CountyTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxAmount

`func (o *UnpostedProcurement) SetCountyTaxAmount(v float64)`

SetCountyTaxAmount sets CountyTaxAmount field to given value.

### HasCountyTaxAmount

`func (o *UnpostedProcurement) HasCountyTaxAmount() bool`

HasCountyTaxAmount returns a boolean if a field has been set.

### SetCountyTaxAmountNil

`func (o *UnpostedProcurement) SetCountyTaxAmountNil(b bool)`

 SetCountyTaxAmountNil sets the value for CountyTaxAmount to be an explicit nil

### UnsetCountyTaxAmount
`func (o *UnpostedProcurement) UnsetCountyTaxAmount()`

UnsetCountyTaxAmount ensures that no value is present for CountyTaxAmount, not even an explicit nil
### GetCityTaxFlag

`func (o *UnpostedProcurement) GetCityTaxFlag() bool`

GetCityTaxFlag returns the CityTaxFlag field if non-nil, zero value otherwise.

### GetCityTaxFlagOk

`func (o *UnpostedProcurement) GetCityTaxFlagOk() (*bool, bool)`

GetCityTaxFlagOk returns a tuple with the CityTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxFlag

`func (o *UnpostedProcurement) SetCityTaxFlag(v bool)`

SetCityTaxFlag sets CityTaxFlag field to given value.

### HasCityTaxFlag

`func (o *UnpostedProcurement) HasCityTaxFlag() bool`

HasCityTaxFlag returns a boolean if a field has been set.

### SetCityTaxFlagNil

`func (o *UnpostedProcurement) SetCityTaxFlagNil(b bool)`

 SetCityTaxFlagNil sets the value for CityTaxFlag to be an explicit nil

### UnsetCityTaxFlag
`func (o *UnpostedProcurement) UnsetCityTaxFlag()`

UnsetCityTaxFlag ensures that no value is present for CityTaxFlag, not even an explicit nil
### GetCityTaxXref

`func (o *UnpostedProcurement) GetCityTaxXref() string`

GetCityTaxXref returns the CityTaxXref field if non-nil, zero value otherwise.

### GetCityTaxXrefOk

`func (o *UnpostedProcurement) GetCityTaxXrefOk() (*string, bool)`

GetCityTaxXrefOk returns a tuple with the CityTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxXref

`func (o *UnpostedProcurement) SetCityTaxXref(v string)`

SetCityTaxXref sets CityTaxXref field to given value.

### HasCityTaxXref

`func (o *UnpostedProcurement) HasCityTaxXref() bool`

HasCityTaxXref returns a boolean if a field has been set.

### GetCityTaxAmount

`func (o *UnpostedProcurement) GetCityTaxAmount() float64`

GetCityTaxAmount returns the CityTaxAmount field if non-nil, zero value otherwise.

### GetCityTaxAmountOk

`func (o *UnpostedProcurement) GetCityTaxAmountOk() (*float64, bool)`

GetCityTaxAmountOk returns a tuple with the CityTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxAmount

`func (o *UnpostedProcurement) SetCityTaxAmount(v float64)`

SetCityTaxAmount sets CityTaxAmount field to given value.

### HasCityTaxAmount

`func (o *UnpostedProcurement) HasCityTaxAmount() bool`

HasCityTaxAmount returns a boolean if a field has been set.

### SetCityTaxAmountNil

`func (o *UnpostedProcurement) SetCityTaxAmountNil(b bool)`

 SetCityTaxAmountNil sets the value for CityTaxAmount to be an explicit nil

### UnsetCityTaxAmount
`func (o *UnpostedProcurement) UnsetCityTaxAmount()`

UnsetCityTaxAmount ensures that no value is present for CityTaxAmount, not even an explicit nil
### GetCountryTaxFlag

`func (o *UnpostedProcurement) GetCountryTaxFlag() bool`

GetCountryTaxFlag returns the CountryTaxFlag field if non-nil, zero value otherwise.

### GetCountryTaxFlagOk

`func (o *UnpostedProcurement) GetCountryTaxFlagOk() (*bool, bool)`

GetCountryTaxFlagOk returns a tuple with the CountryTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryTaxFlag

`func (o *UnpostedProcurement) SetCountryTaxFlag(v bool)`

SetCountryTaxFlag sets CountryTaxFlag field to given value.

### HasCountryTaxFlag

`func (o *UnpostedProcurement) HasCountryTaxFlag() bool`

HasCountryTaxFlag returns a boolean if a field has been set.

### SetCountryTaxFlagNil

`func (o *UnpostedProcurement) SetCountryTaxFlagNil(b bool)`

 SetCountryTaxFlagNil sets the value for CountryTaxFlag to be an explicit nil

### UnsetCountryTaxFlag
`func (o *UnpostedProcurement) UnsetCountryTaxFlag()`

UnsetCountryTaxFlag ensures that no value is present for CountryTaxFlag, not even an explicit nil
### GetCountryTaxXref

`func (o *UnpostedProcurement) GetCountryTaxXref() string`

GetCountryTaxXref returns the CountryTaxXref field if non-nil, zero value otherwise.

### GetCountryTaxXrefOk

`func (o *UnpostedProcurement) GetCountryTaxXrefOk() (*string, bool)`

GetCountryTaxXrefOk returns a tuple with the CountryTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryTaxXref

`func (o *UnpostedProcurement) SetCountryTaxXref(v string)`

SetCountryTaxXref sets CountryTaxXref field to given value.

### HasCountryTaxXref

`func (o *UnpostedProcurement) HasCountryTaxXref() bool`

HasCountryTaxXref returns a boolean if a field has been set.

### GetCountryTaxAmount

`func (o *UnpostedProcurement) GetCountryTaxAmount() float64`

GetCountryTaxAmount returns the CountryTaxAmount field if non-nil, zero value otherwise.

### GetCountryTaxAmountOk

`func (o *UnpostedProcurement) GetCountryTaxAmountOk() (*float64, bool)`

GetCountryTaxAmountOk returns a tuple with the CountryTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryTaxAmount

`func (o *UnpostedProcurement) SetCountryTaxAmount(v float64)`

SetCountryTaxAmount sets CountryTaxAmount field to given value.

### HasCountryTaxAmount

`func (o *UnpostedProcurement) HasCountryTaxAmount() bool`

HasCountryTaxAmount returns a boolean if a field has been set.

### SetCountryTaxAmountNil

`func (o *UnpostedProcurement) SetCountryTaxAmountNil(b bool)`

 SetCountryTaxAmountNil sets the value for CountryTaxAmount to be an explicit nil

### UnsetCountryTaxAmount
`func (o *UnpostedProcurement) UnsetCountryTaxAmount()`

UnsetCountryTaxAmount ensures that no value is present for CountryTaxAmount, not even an explicit nil
### GetCompositeTaxFlag

`func (o *UnpostedProcurement) GetCompositeTaxFlag() bool`

GetCompositeTaxFlag returns the CompositeTaxFlag field if non-nil, zero value otherwise.

### GetCompositeTaxFlagOk

`func (o *UnpostedProcurement) GetCompositeTaxFlagOk() (*bool, bool)`

GetCompositeTaxFlagOk returns a tuple with the CompositeTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeTaxFlag

`func (o *UnpostedProcurement) SetCompositeTaxFlag(v bool)`

SetCompositeTaxFlag sets CompositeTaxFlag field to given value.

### HasCompositeTaxFlag

`func (o *UnpostedProcurement) HasCompositeTaxFlag() bool`

HasCompositeTaxFlag returns a boolean if a field has been set.

### SetCompositeTaxFlagNil

`func (o *UnpostedProcurement) SetCompositeTaxFlagNil(b bool)`

 SetCompositeTaxFlagNil sets the value for CompositeTaxFlag to be an explicit nil

### UnsetCompositeTaxFlag
`func (o *UnpostedProcurement) UnsetCompositeTaxFlag()`

UnsetCompositeTaxFlag ensures that no value is present for CompositeTaxFlag, not even an explicit nil
### GetCompositeTaxXref

`func (o *UnpostedProcurement) GetCompositeTaxXref() string`

GetCompositeTaxXref returns the CompositeTaxXref field if non-nil, zero value otherwise.

### GetCompositeTaxXrefOk

`func (o *UnpostedProcurement) GetCompositeTaxXrefOk() (*string, bool)`

GetCompositeTaxXrefOk returns a tuple with the CompositeTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeTaxXref

`func (o *UnpostedProcurement) SetCompositeTaxXref(v string)`

SetCompositeTaxXref sets CompositeTaxXref field to given value.

### HasCompositeTaxXref

`func (o *UnpostedProcurement) HasCompositeTaxXref() bool`

HasCompositeTaxXref returns a boolean if a field has been set.

### GetCompositeTaxAmount

`func (o *UnpostedProcurement) GetCompositeTaxAmount() float64`

GetCompositeTaxAmount returns the CompositeTaxAmount field if non-nil, zero value otherwise.

### GetCompositeTaxAmountOk

`func (o *UnpostedProcurement) GetCompositeTaxAmountOk() (*float64, bool)`

GetCompositeTaxAmountOk returns a tuple with the CompositeTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeTaxAmount

`func (o *UnpostedProcurement) SetCompositeTaxAmount(v float64)`

SetCompositeTaxAmount sets CompositeTaxAmount field to given value.

### HasCompositeTaxAmount

`func (o *UnpostedProcurement) HasCompositeTaxAmount() bool`

HasCompositeTaxAmount returns a boolean if a field has been set.

### SetCompositeTaxAmountNil

`func (o *UnpostedProcurement) SetCompositeTaxAmountNil(b bool)`

 SetCompositeTaxAmountNil sets the value for CompositeTaxAmount to be an explicit nil

### UnsetCompositeTaxAmount
`func (o *UnpostedProcurement) UnsetCompositeTaxAmount()`

UnsetCompositeTaxAmount ensures that no value is present for CompositeTaxAmount, not even an explicit nil
### GetLevelSixTaxFlag

`func (o *UnpostedProcurement) GetLevelSixTaxFlag() bool`

GetLevelSixTaxFlag returns the LevelSixTaxFlag field if non-nil, zero value otherwise.

### GetLevelSixTaxFlagOk

`func (o *UnpostedProcurement) GetLevelSixTaxFlagOk() (*bool, bool)`

GetLevelSixTaxFlagOk returns a tuple with the LevelSixTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixTaxFlag

`func (o *UnpostedProcurement) SetLevelSixTaxFlag(v bool)`

SetLevelSixTaxFlag sets LevelSixTaxFlag field to given value.

### HasLevelSixTaxFlag

`func (o *UnpostedProcurement) HasLevelSixTaxFlag() bool`

HasLevelSixTaxFlag returns a boolean if a field has been set.

### SetLevelSixTaxFlagNil

`func (o *UnpostedProcurement) SetLevelSixTaxFlagNil(b bool)`

 SetLevelSixTaxFlagNil sets the value for LevelSixTaxFlag to be an explicit nil

### UnsetLevelSixTaxFlag
`func (o *UnpostedProcurement) UnsetLevelSixTaxFlag()`

UnsetLevelSixTaxFlag ensures that no value is present for LevelSixTaxFlag, not even an explicit nil
### GetLevelSixTaxXref

`func (o *UnpostedProcurement) GetLevelSixTaxXref() string`

GetLevelSixTaxXref returns the LevelSixTaxXref field if non-nil, zero value otherwise.

### GetLevelSixTaxXrefOk

`func (o *UnpostedProcurement) GetLevelSixTaxXrefOk() (*string, bool)`

GetLevelSixTaxXrefOk returns a tuple with the LevelSixTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixTaxXref

`func (o *UnpostedProcurement) SetLevelSixTaxXref(v string)`

SetLevelSixTaxXref sets LevelSixTaxXref field to given value.

### HasLevelSixTaxXref

`func (o *UnpostedProcurement) HasLevelSixTaxXref() bool`

HasLevelSixTaxXref returns a boolean if a field has been set.

### GetLevelSixTaxAmount

`func (o *UnpostedProcurement) GetLevelSixTaxAmount() float64`

GetLevelSixTaxAmount returns the LevelSixTaxAmount field if non-nil, zero value otherwise.

### GetLevelSixTaxAmountOk

`func (o *UnpostedProcurement) GetLevelSixTaxAmountOk() (*float64, bool)`

GetLevelSixTaxAmountOk returns a tuple with the LevelSixTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixTaxAmount

`func (o *UnpostedProcurement) SetLevelSixTaxAmount(v float64)`

SetLevelSixTaxAmount sets LevelSixTaxAmount field to given value.

### HasLevelSixTaxAmount

`func (o *UnpostedProcurement) HasLevelSixTaxAmount() bool`

HasLevelSixTaxAmount returns a boolean if a field has been set.

### SetLevelSixTaxAmountNil

`func (o *UnpostedProcurement) SetLevelSixTaxAmountNil(b bool)`

 SetLevelSixTaxAmountNil sets the value for LevelSixTaxAmount to be an explicit nil

### UnsetLevelSixTaxAmount
`func (o *UnpostedProcurement) UnsetLevelSixTaxAmount()`

UnsetLevelSixTaxAmount ensures that no value is present for LevelSixTaxAmount, not even an explicit nil
### GetTaxTotal

`func (o *UnpostedProcurement) GetTaxTotal() float64`

GetTaxTotal returns the TaxTotal field if non-nil, zero value otherwise.

### GetTaxTotalOk

`func (o *UnpostedProcurement) GetTaxTotalOk() (*float64, bool)`

GetTaxTotalOk returns a tuple with the TaxTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxTotal

`func (o *UnpostedProcurement) SetTaxTotal(v float64)`

SetTaxTotal sets TaxTotal field to given value.

### HasTaxTotal

`func (o *UnpostedProcurement) HasTaxTotal() bool`

HasTaxTotal returns a boolean if a field has been set.

### SetTaxTotalNil

`func (o *UnpostedProcurement) SetTaxTotalNil(b bool)`

 SetTaxTotalNil sets the value for TaxTotal to be an explicit nil

### UnsetTaxTotal
`func (o *UnpostedProcurement) UnsetTaxTotal()`

UnsetTaxTotal ensures that no value is present for TaxTotal, not even an explicit nil
### GetCustomer

`func (o *UnpostedProcurement) GetCustomer() CompanyReference`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *UnpostedProcurement) GetCustomerOk() (*CompanyReference, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *UnpostedProcurement) SetCustomer(v CompanyReference)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *UnpostedProcurement) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetVendor

`func (o *UnpostedProcurement) GetVendor() CompanyReference`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *UnpostedProcurement) GetVendorOk() (*CompanyReference, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *UnpostedProcurement) SetVendor(v CompanyReference)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *UnpostedProcurement) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVendorAccountNumber

`func (o *UnpostedProcurement) GetVendorAccountNumber() string`

GetVendorAccountNumber returns the VendorAccountNumber field if non-nil, zero value otherwise.

### GetVendorAccountNumberOk

`func (o *UnpostedProcurement) GetVendorAccountNumberOk() (*string, bool)`

GetVendorAccountNumberOk returns a tuple with the VendorAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAccountNumber

`func (o *UnpostedProcurement) SetVendorAccountNumber(v string)`

SetVendorAccountNumber sets VendorAccountNumber field to given value.

### HasVendorAccountNumber

`func (o *UnpostedProcurement) HasVendorAccountNumber() bool`

HasVendorAccountNumber returns a boolean if a field has been set.

### GetVendorInvoiceNumber

`func (o *UnpostedProcurement) GetVendorInvoiceNumber() string`

GetVendorInvoiceNumber returns the VendorInvoiceNumber field if non-nil, zero value otherwise.

### GetVendorInvoiceNumberOk

`func (o *UnpostedProcurement) GetVendorInvoiceNumberOk() (*string, bool)`

GetVendorInvoiceNumberOk returns a tuple with the VendorInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInvoiceNumber

`func (o *UnpostedProcurement) SetVendorInvoiceNumber(v string)`

SetVendorInvoiceNumber sets VendorInvoiceNumber field to given value.

### HasVendorInvoiceNumber

`func (o *UnpostedProcurement) HasVendorInvoiceNumber() bool`

HasVendorInvoiceNumber returns a boolean if a field has been set.

### GetVendorInvoiceDate

`func (o *UnpostedProcurement) GetVendorInvoiceDate() string`

GetVendorInvoiceDate returns the VendorInvoiceDate field if non-nil, zero value otherwise.

### GetVendorInvoiceDateOk

`func (o *UnpostedProcurement) GetVendorInvoiceDateOk() (*string, bool)`

GetVendorInvoiceDateOk returns a tuple with the VendorInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInvoiceDate

`func (o *UnpostedProcurement) SetVendorInvoiceDate(v string)`

SetVendorInvoiceDate sets VendorInvoiceDate field to given value.

### HasVendorInvoiceDate

`func (o *UnpostedProcurement) HasVendorInvoiceDate() bool`

HasVendorInvoiceDate returns a boolean if a field has been set.

### GetTaxFreightFlag

`func (o *UnpostedProcurement) GetTaxFreightFlag() bool`

GetTaxFreightFlag returns the TaxFreightFlag field if non-nil, zero value otherwise.

### GetTaxFreightFlagOk

`func (o *UnpostedProcurement) GetTaxFreightFlagOk() (*bool, bool)`

GetTaxFreightFlagOk returns a tuple with the TaxFreightFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFreightFlag

`func (o *UnpostedProcurement) SetTaxFreightFlag(v bool)`

SetTaxFreightFlag sets TaxFreightFlag field to given value.

### HasTaxFreightFlag

`func (o *UnpostedProcurement) HasTaxFreightFlag() bool`

HasTaxFreightFlag returns a boolean if a field has been set.

### SetTaxFreightFlagNil

`func (o *UnpostedProcurement) SetTaxFreightFlagNil(b bool)`

 SetTaxFreightFlagNil sets the value for TaxFreightFlag to be an explicit nil

### UnsetTaxFreightFlag
`func (o *UnpostedProcurement) UnsetTaxFreightFlag()`

UnsetTaxFreightFlag ensures that no value is present for TaxFreightFlag, not even an explicit nil
### GetFreightTaxTotal

`func (o *UnpostedProcurement) GetFreightTaxTotal() float64`

GetFreightTaxTotal returns the FreightTaxTotal field if non-nil, zero value otherwise.

### GetFreightTaxTotalOk

`func (o *UnpostedProcurement) GetFreightTaxTotalOk() (*float64, bool)`

GetFreightTaxTotalOk returns a tuple with the FreightTaxTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxTotal

`func (o *UnpostedProcurement) SetFreightTaxTotal(v float64)`

SetFreightTaxTotal sets FreightTaxTotal field to given value.

### HasFreightTaxTotal

`func (o *UnpostedProcurement) HasFreightTaxTotal() bool`

HasFreightTaxTotal returns a boolean if a field has been set.

### SetFreightTaxTotalNil

`func (o *UnpostedProcurement) SetFreightTaxTotalNil(b bool)`

 SetFreightTaxTotalNil sets the value for FreightTaxTotal to be an explicit nil

### UnsetFreightTaxTotal
`func (o *UnpostedProcurement) UnsetFreightTaxTotal()`

UnsetFreightTaxTotal ensures that no value is present for FreightTaxTotal, not even an explicit nil
### GetFreightCost

`func (o *UnpostedProcurement) GetFreightCost() float64`

GetFreightCost returns the FreightCost field if non-nil, zero value otherwise.

### GetFreightCostOk

`func (o *UnpostedProcurement) GetFreightCostOk() (*float64, bool)`

GetFreightCostOk returns a tuple with the FreightCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCost

`func (o *UnpostedProcurement) SetFreightCost(v float64)`

SetFreightCost sets FreightCost field to given value.

### HasFreightCost

`func (o *UnpostedProcurement) HasFreightCost() bool`

HasFreightCost returns a boolean if a field has been set.

### SetFreightCostNil

`func (o *UnpostedProcurement) SetFreightCostNil(b bool)`

 SetFreightCostNil sets the value for FreightCost to be an explicit nil

### UnsetFreightCost
`func (o *UnpostedProcurement) UnsetFreightCost()`

UnsetFreightCost ensures that no value is present for FreightCost, not even an explicit nil
### GetDateClosed

`func (o *UnpostedProcurement) GetDateClosed() string`

GetDateClosed returns the DateClosed field if non-nil, zero value otherwise.

### GetDateClosedOk

`func (o *UnpostedProcurement) GetDateClosedOk() (*string, bool)`

GetDateClosedOk returns a tuple with the DateClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClosed

`func (o *UnpostedProcurement) SetDateClosed(v string)`

SetDateClosed sets DateClosed field to given value.

### HasDateClosed

`func (o *UnpostedProcurement) HasDateClosed() bool`

HasDateClosed returns a boolean if a field has been set.

### GetInfo

`func (o *UnpostedProcurement) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UnpostedProcurement) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UnpostedProcurement) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *UnpostedProcurement) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


