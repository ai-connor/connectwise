# PurchaseOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  | [optional] 
**BusinessUnit** | Pointer to [**BillingUnitReference**](BillingUnitReference.md) |  | [optional] 
**CancelReason** | Pointer to **string** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** | The closed flag can only be updated via updating the purchase order status to a closed/open status. | [optional] 
**ClosedBy** | Pointer to **string** |  | [optional] 
**CustomerCity** | Pointer to **string** |  | [optional] 
**CustomerCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**CustomerContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**CustomerCountry** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**CustomerExtension** | Pointer to **string** |  | [optional] 
**CustomerName** | Pointer to **string** |  | [optional] 
**CustomerPhone** | Pointer to **string** |  | [optional] 
**CustomerSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**CustomerSiteName** | Pointer to **string** |  | [optional] 
**CustomerState** | Pointer to **string** |  | [optional] 
**CustomerStreetLine1** | Pointer to **string** |  | [optional] 
**CustomerStreetLine2** | Pointer to **string** |  | [optional] 
**CustomerZip** | Pointer to **string** |  | [optional] 
**DateClosed** | Pointer to **time.Time** |  | [optional] 
**DropShipCustomerFlag** | Pointer to **NullableBool** |  | [optional] 
**EnteredBy** | Pointer to **string** |  | [optional] 
**FreightCost** | Pointer to **NullableFloat64** |  | [optional] 
**FreightPackingSlip** | Pointer to **string** |  | [optional] 
**FreightTaxTotal** | Pointer to **NullableFloat64** |  | [optional] 
**InternalNotes** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**PoDate** | Pointer to **time.Time** |  Required On Updates; | [optional] 
**PoNumber** | Pointer to **string** |  Required On Updates; Max length: 50; | [optional] 
**SalesTax** | Pointer to **NullableFloat64** |  | [optional] 
**ShipmentDate** | Pointer to **time.Time** |  | [optional] 
**ShipmentMethod** | Pointer to [**ShipmentMethodReference**](ShipmentMethodReference.md) |  | [optional] 
**ShippingInstructions** | Pointer to **string** |  | [optional] 
**Status** | [**PurchaseOrderStatusReference**](PurchaseOrderStatusReference.md) |  | 
**SubTotal** | Pointer to **NullableFloat64** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**TaxFreightFlag** | Pointer to **NullableBool** |  | [optional] 
**TaxPoFlag** | Pointer to **NullableBool** |  | [optional] 
**Terms** | [**BillingTermsReference**](BillingTermsReference.md) |  | 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**TrackingNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**UpdateShipmentInfo** | Pointer to **NullableBool** | Determines whether or not to update all of the shipment info for each associated line item when new shipment info is passed in. | [optional] 
**UpdateVendorOrderNumber** | Pointer to **NullableBool** | Determines whether or not to update vendor order number for each associated line item when new vendor order number is passed in. | [optional] 
**VendorCompany** | [**CompanyReference**](CompanyReference.md) |  | 
**VendorContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**VendorInvoiceDate** | Pointer to **time.Time** |  | [optional] 
**VendorInvoiceNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**VendorOrderNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**VendorSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**Warehouse** | Pointer to [**WarehouseReference**](WarehouseReference.md) |  | [optional] 
**WarehouseContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewPurchaseOrder

`func NewPurchaseOrder(status PurchaseOrderStatusReference, terms BillingTermsReference, vendorCompany CompanyReference, ) *PurchaseOrder`

NewPurchaseOrder instantiates a new PurchaseOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderWithDefaults

`func NewPurchaseOrderWithDefaults() *PurchaseOrder`

NewPurchaseOrderWithDefaults instantiates a new PurchaseOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PurchaseOrder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseOrder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseOrder) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *PurchaseOrder) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *PurchaseOrder) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *PurchaseOrder) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *PurchaseOrder) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *PurchaseOrder) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *PurchaseOrder) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetBusinessUnit

`func (o *PurchaseOrder) GetBusinessUnit() BillingUnitReference`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *PurchaseOrder) GetBusinessUnitOk() (*BillingUnitReference, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *PurchaseOrder) SetBusinessUnit(v BillingUnitReference)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *PurchaseOrder) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetCancelReason

`func (o *PurchaseOrder) GetCancelReason() string`

GetCancelReason returns the CancelReason field if non-nil, zero value otherwise.

### GetCancelReasonOk

`func (o *PurchaseOrder) GetCancelReasonOk() (*string, bool)`

GetCancelReasonOk returns a tuple with the CancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReason

`func (o *PurchaseOrder) SetCancelReason(v string)`

SetCancelReason sets CancelReason field to given value.

### HasCancelReason

`func (o *PurchaseOrder) HasCancelReason() bool`

HasCancelReason returns a boolean if a field has been set.

### GetClosedFlag

`func (o *PurchaseOrder) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *PurchaseOrder) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *PurchaseOrder) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *PurchaseOrder) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *PurchaseOrder) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *PurchaseOrder) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetClosedBy

`func (o *PurchaseOrder) GetClosedBy() string`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *PurchaseOrder) GetClosedByOk() (*string, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *PurchaseOrder) SetClosedBy(v string)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *PurchaseOrder) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetCustomerCity

`func (o *PurchaseOrder) GetCustomerCity() string`

GetCustomerCity returns the CustomerCity field if non-nil, zero value otherwise.

### GetCustomerCityOk

`func (o *PurchaseOrder) GetCustomerCityOk() (*string, bool)`

GetCustomerCityOk returns a tuple with the CustomerCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCity

`func (o *PurchaseOrder) SetCustomerCity(v string)`

SetCustomerCity sets CustomerCity field to given value.

### HasCustomerCity

`func (o *PurchaseOrder) HasCustomerCity() bool`

HasCustomerCity returns a boolean if a field has been set.

### GetCustomerCompany

`func (o *PurchaseOrder) GetCustomerCompany() CompanyReference`

GetCustomerCompany returns the CustomerCompany field if non-nil, zero value otherwise.

### GetCustomerCompanyOk

`func (o *PurchaseOrder) GetCustomerCompanyOk() (*CompanyReference, bool)`

GetCustomerCompanyOk returns a tuple with the CustomerCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCompany

`func (o *PurchaseOrder) SetCustomerCompany(v CompanyReference)`

SetCustomerCompany sets CustomerCompany field to given value.

### HasCustomerCompany

`func (o *PurchaseOrder) HasCustomerCompany() bool`

HasCustomerCompany returns a boolean if a field has been set.

### GetCustomerContact

`func (o *PurchaseOrder) GetCustomerContact() ContactReference`

GetCustomerContact returns the CustomerContact field if non-nil, zero value otherwise.

### GetCustomerContactOk

`func (o *PurchaseOrder) GetCustomerContactOk() (*ContactReference, bool)`

GetCustomerContactOk returns a tuple with the CustomerContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerContact

`func (o *PurchaseOrder) SetCustomerContact(v ContactReference)`

SetCustomerContact sets CustomerContact field to given value.

### HasCustomerContact

`func (o *PurchaseOrder) HasCustomerContact() bool`

HasCustomerContact returns a boolean if a field has been set.

### GetCustomerCountry

`func (o *PurchaseOrder) GetCustomerCountry() CountryReference`

GetCustomerCountry returns the CustomerCountry field if non-nil, zero value otherwise.

### GetCustomerCountryOk

`func (o *PurchaseOrder) GetCustomerCountryOk() (*CountryReference, bool)`

GetCustomerCountryOk returns a tuple with the CustomerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCountry

`func (o *PurchaseOrder) SetCustomerCountry(v CountryReference)`

SetCustomerCountry sets CustomerCountry field to given value.

### HasCustomerCountry

`func (o *PurchaseOrder) HasCustomerCountry() bool`

HasCustomerCountry returns a boolean if a field has been set.

### GetCustomerExtension

`func (o *PurchaseOrder) GetCustomerExtension() string`

GetCustomerExtension returns the CustomerExtension field if non-nil, zero value otherwise.

### GetCustomerExtensionOk

`func (o *PurchaseOrder) GetCustomerExtensionOk() (*string, bool)`

GetCustomerExtensionOk returns a tuple with the CustomerExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExtension

`func (o *PurchaseOrder) SetCustomerExtension(v string)`

SetCustomerExtension sets CustomerExtension field to given value.

### HasCustomerExtension

`func (o *PurchaseOrder) HasCustomerExtension() bool`

HasCustomerExtension returns a boolean if a field has been set.

### GetCustomerName

`func (o *PurchaseOrder) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *PurchaseOrder) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *PurchaseOrder) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *PurchaseOrder) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetCustomerPhone

`func (o *PurchaseOrder) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *PurchaseOrder) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *PurchaseOrder) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.

### HasCustomerPhone

`func (o *PurchaseOrder) HasCustomerPhone() bool`

HasCustomerPhone returns a boolean if a field has been set.

### GetCustomerSite

`func (o *PurchaseOrder) GetCustomerSite() SiteReference`

GetCustomerSite returns the CustomerSite field if non-nil, zero value otherwise.

### GetCustomerSiteOk

`func (o *PurchaseOrder) GetCustomerSiteOk() (*SiteReference, bool)`

GetCustomerSiteOk returns a tuple with the CustomerSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSite

`func (o *PurchaseOrder) SetCustomerSite(v SiteReference)`

SetCustomerSite sets CustomerSite field to given value.

### HasCustomerSite

`func (o *PurchaseOrder) HasCustomerSite() bool`

HasCustomerSite returns a boolean if a field has been set.

### GetCustomerSiteName

`func (o *PurchaseOrder) GetCustomerSiteName() string`

GetCustomerSiteName returns the CustomerSiteName field if non-nil, zero value otherwise.

### GetCustomerSiteNameOk

`func (o *PurchaseOrder) GetCustomerSiteNameOk() (*string, bool)`

GetCustomerSiteNameOk returns a tuple with the CustomerSiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSiteName

`func (o *PurchaseOrder) SetCustomerSiteName(v string)`

SetCustomerSiteName sets CustomerSiteName field to given value.

### HasCustomerSiteName

`func (o *PurchaseOrder) HasCustomerSiteName() bool`

HasCustomerSiteName returns a boolean if a field has been set.

### GetCustomerState

`func (o *PurchaseOrder) GetCustomerState() string`

GetCustomerState returns the CustomerState field if non-nil, zero value otherwise.

### GetCustomerStateOk

`func (o *PurchaseOrder) GetCustomerStateOk() (*string, bool)`

GetCustomerStateOk returns a tuple with the CustomerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerState

`func (o *PurchaseOrder) SetCustomerState(v string)`

SetCustomerState sets CustomerState field to given value.

### HasCustomerState

`func (o *PurchaseOrder) HasCustomerState() bool`

HasCustomerState returns a boolean if a field has been set.

### GetCustomerStreetLine1

`func (o *PurchaseOrder) GetCustomerStreetLine1() string`

GetCustomerStreetLine1 returns the CustomerStreetLine1 field if non-nil, zero value otherwise.

### GetCustomerStreetLine1Ok

`func (o *PurchaseOrder) GetCustomerStreetLine1Ok() (*string, bool)`

GetCustomerStreetLine1Ok returns a tuple with the CustomerStreetLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerStreetLine1

`func (o *PurchaseOrder) SetCustomerStreetLine1(v string)`

SetCustomerStreetLine1 sets CustomerStreetLine1 field to given value.

### HasCustomerStreetLine1

`func (o *PurchaseOrder) HasCustomerStreetLine1() bool`

HasCustomerStreetLine1 returns a boolean if a field has been set.

### GetCustomerStreetLine2

`func (o *PurchaseOrder) GetCustomerStreetLine2() string`

GetCustomerStreetLine2 returns the CustomerStreetLine2 field if non-nil, zero value otherwise.

### GetCustomerStreetLine2Ok

`func (o *PurchaseOrder) GetCustomerStreetLine2Ok() (*string, bool)`

GetCustomerStreetLine2Ok returns a tuple with the CustomerStreetLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerStreetLine2

`func (o *PurchaseOrder) SetCustomerStreetLine2(v string)`

SetCustomerStreetLine2 sets CustomerStreetLine2 field to given value.

### HasCustomerStreetLine2

`func (o *PurchaseOrder) HasCustomerStreetLine2() bool`

HasCustomerStreetLine2 returns a boolean if a field has been set.

### GetCustomerZip

`func (o *PurchaseOrder) GetCustomerZip() string`

GetCustomerZip returns the CustomerZip field if non-nil, zero value otherwise.

### GetCustomerZipOk

`func (o *PurchaseOrder) GetCustomerZipOk() (*string, bool)`

GetCustomerZipOk returns a tuple with the CustomerZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerZip

`func (o *PurchaseOrder) SetCustomerZip(v string)`

SetCustomerZip sets CustomerZip field to given value.

### HasCustomerZip

`func (o *PurchaseOrder) HasCustomerZip() bool`

HasCustomerZip returns a boolean if a field has been set.

### GetDateClosed

`func (o *PurchaseOrder) GetDateClosed() time.Time`

GetDateClosed returns the DateClosed field if non-nil, zero value otherwise.

### GetDateClosedOk

`func (o *PurchaseOrder) GetDateClosedOk() (*time.Time, bool)`

GetDateClosedOk returns a tuple with the DateClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClosed

`func (o *PurchaseOrder) SetDateClosed(v time.Time)`

SetDateClosed sets DateClosed field to given value.

### HasDateClosed

`func (o *PurchaseOrder) HasDateClosed() bool`

HasDateClosed returns a boolean if a field has been set.

### GetDropShipCustomerFlag

`func (o *PurchaseOrder) GetDropShipCustomerFlag() bool`

GetDropShipCustomerFlag returns the DropShipCustomerFlag field if non-nil, zero value otherwise.

### GetDropShipCustomerFlagOk

`func (o *PurchaseOrder) GetDropShipCustomerFlagOk() (*bool, bool)`

GetDropShipCustomerFlagOk returns a tuple with the DropShipCustomerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropShipCustomerFlag

`func (o *PurchaseOrder) SetDropShipCustomerFlag(v bool)`

SetDropShipCustomerFlag sets DropShipCustomerFlag field to given value.

### HasDropShipCustomerFlag

`func (o *PurchaseOrder) HasDropShipCustomerFlag() bool`

HasDropShipCustomerFlag returns a boolean if a field has been set.

### SetDropShipCustomerFlagNil

`func (o *PurchaseOrder) SetDropShipCustomerFlagNil(b bool)`

 SetDropShipCustomerFlagNil sets the value for DropShipCustomerFlag to be an explicit nil

### UnsetDropShipCustomerFlag
`func (o *PurchaseOrder) UnsetDropShipCustomerFlag()`

UnsetDropShipCustomerFlag ensures that no value is present for DropShipCustomerFlag, not even an explicit nil
### GetEnteredBy

`func (o *PurchaseOrder) GetEnteredBy() string`

GetEnteredBy returns the EnteredBy field if non-nil, zero value otherwise.

### GetEnteredByOk

`func (o *PurchaseOrder) GetEnteredByOk() (*string, bool)`

GetEnteredByOk returns a tuple with the EnteredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredBy

`func (o *PurchaseOrder) SetEnteredBy(v string)`

SetEnteredBy sets EnteredBy field to given value.

### HasEnteredBy

`func (o *PurchaseOrder) HasEnteredBy() bool`

HasEnteredBy returns a boolean if a field has been set.

### GetFreightCost

`func (o *PurchaseOrder) GetFreightCost() float64`

GetFreightCost returns the FreightCost field if non-nil, zero value otherwise.

### GetFreightCostOk

`func (o *PurchaseOrder) GetFreightCostOk() (*float64, bool)`

GetFreightCostOk returns a tuple with the FreightCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCost

`func (o *PurchaseOrder) SetFreightCost(v float64)`

SetFreightCost sets FreightCost field to given value.

### HasFreightCost

`func (o *PurchaseOrder) HasFreightCost() bool`

HasFreightCost returns a boolean if a field has been set.

### SetFreightCostNil

`func (o *PurchaseOrder) SetFreightCostNil(b bool)`

 SetFreightCostNil sets the value for FreightCost to be an explicit nil

### UnsetFreightCost
`func (o *PurchaseOrder) UnsetFreightCost()`

UnsetFreightCost ensures that no value is present for FreightCost, not even an explicit nil
### GetFreightPackingSlip

`func (o *PurchaseOrder) GetFreightPackingSlip() string`

GetFreightPackingSlip returns the FreightPackingSlip field if non-nil, zero value otherwise.

### GetFreightPackingSlipOk

`func (o *PurchaseOrder) GetFreightPackingSlipOk() (*string, bool)`

GetFreightPackingSlipOk returns a tuple with the FreightPackingSlip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightPackingSlip

`func (o *PurchaseOrder) SetFreightPackingSlip(v string)`

SetFreightPackingSlip sets FreightPackingSlip field to given value.

### HasFreightPackingSlip

`func (o *PurchaseOrder) HasFreightPackingSlip() bool`

HasFreightPackingSlip returns a boolean if a field has been set.

### GetFreightTaxTotal

`func (o *PurchaseOrder) GetFreightTaxTotal() float64`

GetFreightTaxTotal returns the FreightTaxTotal field if non-nil, zero value otherwise.

### GetFreightTaxTotalOk

`func (o *PurchaseOrder) GetFreightTaxTotalOk() (*float64, bool)`

GetFreightTaxTotalOk returns a tuple with the FreightTaxTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxTotal

`func (o *PurchaseOrder) SetFreightTaxTotal(v float64)`

SetFreightTaxTotal sets FreightTaxTotal field to given value.

### HasFreightTaxTotal

`func (o *PurchaseOrder) HasFreightTaxTotal() bool`

HasFreightTaxTotal returns a boolean if a field has been set.

### SetFreightTaxTotalNil

`func (o *PurchaseOrder) SetFreightTaxTotalNil(b bool)`

 SetFreightTaxTotalNil sets the value for FreightTaxTotal to be an explicit nil

### UnsetFreightTaxTotal
`func (o *PurchaseOrder) UnsetFreightTaxTotal()`

UnsetFreightTaxTotal ensures that no value is present for FreightTaxTotal, not even an explicit nil
### GetInternalNotes

`func (o *PurchaseOrder) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *PurchaseOrder) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *PurchaseOrder) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *PurchaseOrder) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetLocationId

`func (o *PurchaseOrder) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *PurchaseOrder) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *PurchaseOrder) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *PurchaseOrder) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *PurchaseOrder) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *PurchaseOrder) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetLocation

`func (o *PurchaseOrder) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PurchaseOrder) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PurchaseOrder) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PurchaseOrder) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPoDate

`func (o *PurchaseOrder) GetPoDate() time.Time`

GetPoDate returns the PoDate field if non-nil, zero value otherwise.

### GetPoDateOk

`func (o *PurchaseOrder) GetPoDateOk() (*time.Time, bool)`

GetPoDateOk returns a tuple with the PoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoDate

`func (o *PurchaseOrder) SetPoDate(v time.Time)`

SetPoDate sets PoDate field to given value.

### HasPoDate

`func (o *PurchaseOrder) HasPoDate() bool`

HasPoDate returns a boolean if a field has been set.

### GetPoNumber

`func (o *PurchaseOrder) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *PurchaseOrder) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *PurchaseOrder) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *PurchaseOrder) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetSalesTax

`func (o *PurchaseOrder) GetSalesTax() float64`

GetSalesTax returns the SalesTax field if non-nil, zero value otherwise.

### GetSalesTaxOk

`func (o *PurchaseOrder) GetSalesTaxOk() (*float64, bool)`

GetSalesTaxOk returns a tuple with the SalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTax

`func (o *PurchaseOrder) SetSalesTax(v float64)`

SetSalesTax sets SalesTax field to given value.

### HasSalesTax

`func (o *PurchaseOrder) HasSalesTax() bool`

HasSalesTax returns a boolean if a field has been set.

### SetSalesTaxNil

`func (o *PurchaseOrder) SetSalesTaxNil(b bool)`

 SetSalesTaxNil sets the value for SalesTax to be an explicit nil

### UnsetSalesTax
`func (o *PurchaseOrder) UnsetSalesTax()`

UnsetSalesTax ensures that no value is present for SalesTax, not even an explicit nil
### GetShipmentDate

`func (o *PurchaseOrder) GetShipmentDate() time.Time`

GetShipmentDate returns the ShipmentDate field if non-nil, zero value otherwise.

### GetShipmentDateOk

`func (o *PurchaseOrder) GetShipmentDateOk() (*time.Time, bool)`

GetShipmentDateOk returns a tuple with the ShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDate

`func (o *PurchaseOrder) SetShipmentDate(v time.Time)`

SetShipmentDate sets ShipmentDate field to given value.

### HasShipmentDate

`func (o *PurchaseOrder) HasShipmentDate() bool`

HasShipmentDate returns a boolean if a field has been set.

### GetShipmentMethod

`func (o *PurchaseOrder) GetShipmentMethod() ShipmentMethodReference`

GetShipmentMethod returns the ShipmentMethod field if non-nil, zero value otherwise.

### GetShipmentMethodOk

`func (o *PurchaseOrder) GetShipmentMethodOk() (*ShipmentMethodReference, bool)`

GetShipmentMethodOk returns a tuple with the ShipmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentMethod

`func (o *PurchaseOrder) SetShipmentMethod(v ShipmentMethodReference)`

SetShipmentMethod sets ShipmentMethod field to given value.

### HasShipmentMethod

`func (o *PurchaseOrder) HasShipmentMethod() bool`

HasShipmentMethod returns a boolean if a field has been set.

### GetShippingInstructions

`func (o *PurchaseOrder) GetShippingInstructions() string`

GetShippingInstructions returns the ShippingInstructions field if non-nil, zero value otherwise.

### GetShippingInstructionsOk

`func (o *PurchaseOrder) GetShippingInstructionsOk() (*string, bool)`

GetShippingInstructionsOk returns a tuple with the ShippingInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInstructions

`func (o *PurchaseOrder) SetShippingInstructions(v string)`

SetShippingInstructions sets ShippingInstructions field to given value.

### HasShippingInstructions

`func (o *PurchaseOrder) HasShippingInstructions() bool`

HasShippingInstructions returns a boolean if a field has been set.

### GetStatus

`func (o *PurchaseOrder) GetStatus() PurchaseOrderStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PurchaseOrder) GetStatusOk() (*PurchaseOrderStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PurchaseOrder) SetStatus(v PurchaseOrderStatusReference)`

SetStatus sets Status field to given value.


### GetSubTotal

`func (o *PurchaseOrder) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *PurchaseOrder) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *PurchaseOrder) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *PurchaseOrder) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### SetSubTotalNil

`func (o *PurchaseOrder) SetSubTotalNil(b bool)`

 SetSubTotalNil sets the value for SubTotal to be an explicit nil

### UnsetSubTotal
`func (o *PurchaseOrder) UnsetSubTotal()`

UnsetSubTotal ensures that no value is present for SubTotal, not even an explicit nil
### GetTaxCode

`func (o *PurchaseOrder) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *PurchaseOrder) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *PurchaseOrder) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *PurchaseOrder) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxFreightFlag

`func (o *PurchaseOrder) GetTaxFreightFlag() bool`

GetTaxFreightFlag returns the TaxFreightFlag field if non-nil, zero value otherwise.

### GetTaxFreightFlagOk

`func (o *PurchaseOrder) GetTaxFreightFlagOk() (*bool, bool)`

GetTaxFreightFlagOk returns a tuple with the TaxFreightFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFreightFlag

`func (o *PurchaseOrder) SetTaxFreightFlag(v bool)`

SetTaxFreightFlag sets TaxFreightFlag field to given value.

### HasTaxFreightFlag

`func (o *PurchaseOrder) HasTaxFreightFlag() bool`

HasTaxFreightFlag returns a boolean if a field has been set.

### SetTaxFreightFlagNil

`func (o *PurchaseOrder) SetTaxFreightFlagNil(b bool)`

 SetTaxFreightFlagNil sets the value for TaxFreightFlag to be an explicit nil

### UnsetTaxFreightFlag
`func (o *PurchaseOrder) UnsetTaxFreightFlag()`

UnsetTaxFreightFlag ensures that no value is present for TaxFreightFlag, not even an explicit nil
### GetTaxPoFlag

`func (o *PurchaseOrder) GetTaxPoFlag() bool`

GetTaxPoFlag returns the TaxPoFlag field if non-nil, zero value otherwise.

### GetTaxPoFlagOk

`func (o *PurchaseOrder) GetTaxPoFlagOk() (*bool, bool)`

GetTaxPoFlagOk returns a tuple with the TaxPoFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPoFlag

`func (o *PurchaseOrder) SetTaxPoFlag(v bool)`

SetTaxPoFlag sets TaxPoFlag field to given value.

### HasTaxPoFlag

`func (o *PurchaseOrder) HasTaxPoFlag() bool`

HasTaxPoFlag returns a boolean if a field has been set.

### SetTaxPoFlagNil

`func (o *PurchaseOrder) SetTaxPoFlagNil(b bool)`

 SetTaxPoFlagNil sets the value for TaxPoFlag to be an explicit nil

### UnsetTaxPoFlag
`func (o *PurchaseOrder) UnsetTaxPoFlag()`

UnsetTaxPoFlag ensures that no value is present for TaxPoFlag, not even an explicit nil
### GetTerms

`func (o *PurchaseOrder) GetTerms() BillingTermsReference`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *PurchaseOrder) GetTermsOk() (*BillingTermsReference, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *PurchaseOrder) SetTerms(v BillingTermsReference)`

SetTerms sets Terms field to given value.


### GetTotal

`func (o *PurchaseOrder) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PurchaseOrder) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PurchaseOrder) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PurchaseOrder) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *PurchaseOrder) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *PurchaseOrder) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetTrackingNumber

`func (o *PurchaseOrder) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *PurchaseOrder) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *PurchaseOrder) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *PurchaseOrder) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetUpdateShipmentInfo

`func (o *PurchaseOrder) GetUpdateShipmentInfo() bool`

GetUpdateShipmentInfo returns the UpdateShipmentInfo field if non-nil, zero value otherwise.

### GetUpdateShipmentInfoOk

`func (o *PurchaseOrder) GetUpdateShipmentInfoOk() (*bool, bool)`

GetUpdateShipmentInfoOk returns a tuple with the UpdateShipmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateShipmentInfo

`func (o *PurchaseOrder) SetUpdateShipmentInfo(v bool)`

SetUpdateShipmentInfo sets UpdateShipmentInfo field to given value.

### HasUpdateShipmentInfo

`func (o *PurchaseOrder) HasUpdateShipmentInfo() bool`

HasUpdateShipmentInfo returns a boolean if a field has been set.

### SetUpdateShipmentInfoNil

`func (o *PurchaseOrder) SetUpdateShipmentInfoNil(b bool)`

 SetUpdateShipmentInfoNil sets the value for UpdateShipmentInfo to be an explicit nil

### UnsetUpdateShipmentInfo
`func (o *PurchaseOrder) UnsetUpdateShipmentInfo()`

UnsetUpdateShipmentInfo ensures that no value is present for UpdateShipmentInfo, not even an explicit nil
### GetUpdateVendorOrderNumber

`func (o *PurchaseOrder) GetUpdateVendorOrderNumber() bool`

GetUpdateVendorOrderNumber returns the UpdateVendorOrderNumber field if non-nil, zero value otherwise.

### GetUpdateVendorOrderNumberOk

`func (o *PurchaseOrder) GetUpdateVendorOrderNumberOk() (*bool, bool)`

GetUpdateVendorOrderNumberOk returns a tuple with the UpdateVendorOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVendorOrderNumber

`func (o *PurchaseOrder) SetUpdateVendorOrderNumber(v bool)`

SetUpdateVendorOrderNumber sets UpdateVendorOrderNumber field to given value.

### HasUpdateVendorOrderNumber

`func (o *PurchaseOrder) HasUpdateVendorOrderNumber() bool`

HasUpdateVendorOrderNumber returns a boolean if a field has been set.

### SetUpdateVendorOrderNumberNil

`func (o *PurchaseOrder) SetUpdateVendorOrderNumberNil(b bool)`

 SetUpdateVendorOrderNumberNil sets the value for UpdateVendorOrderNumber to be an explicit nil

### UnsetUpdateVendorOrderNumber
`func (o *PurchaseOrder) UnsetUpdateVendorOrderNumber()`

UnsetUpdateVendorOrderNumber ensures that no value is present for UpdateVendorOrderNumber, not even an explicit nil
### GetVendorCompany

`func (o *PurchaseOrder) GetVendorCompany() CompanyReference`

GetVendorCompany returns the VendorCompany field if non-nil, zero value otherwise.

### GetVendorCompanyOk

`func (o *PurchaseOrder) GetVendorCompanyOk() (*CompanyReference, bool)`

GetVendorCompanyOk returns a tuple with the VendorCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCompany

`func (o *PurchaseOrder) SetVendorCompany(v CompanyReference)`

SetVendorCompany sets VendorCompany field to given value.


### GetVendorContact

`func (o *PurchaseOrder) GetVendorContact() ContactReference`

GetVendorContact returns the VendorContact field if non-nil, zero value otherwise.

### GetVendorContactOk

`func (o *PurchaseOrder) GetVendorContactOk() (*ContactReference, bool)`

GetVendorContactOk returns a tuple with the VendorContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorContact

`func (o *PurchaseOrder) SetVendorContact(v ContactReference)`

SetVendorContact sets VendorContact field to given value.

### HasVendorContact

`func (o *PurchaseOrder) HasVendorContact() bool`

HasVendorContact returns a boolean if a field has been set.

### GetVendorInvoiceDate

`func (o *PurchaseOrder) GetVendorInvoiceDate() time.Time`

GetVendorInvoiceDate returns the VendorInvoiceDate field if non-nil, zero value otherwise.

### GetVendorInvoiceDateOk

`func (o *PurchaseOrder) GetVendorInvoiceDateOk() (*time.Time, bool)`

GetVendorInvoiceDateOk returns a tuple with the VendorInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInvoiceDate

`func (o *PurchaseOrder) SetVendorInvoiceDate(v time.Time)`

SetVendorInvoiceDate sets VendorInvoiceDate field to given value.

### HasVendorInvoiceDate

`func (o *PurchaseOrder) HasVendorInvoiceDate() bool`

HasVendorInvoiceDate returns a boolean if a field has been set.

### GetVendorInvoiceNumber

`func (o *PurchaseOrder) GetVendorInvoiceNumber() string`

GetVendorInvoiceNumber returns the VendorInvoiceNumber field if non-nil, zero value otherwise.

### GetVendorInvoiceNumberOk

`func (o *PurchaseOrder) GetVendorInvoiceNumberOk() (*string, bool)`

GetVendorInvoiceNumberOk returns a tuple with the VendorInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInvoiceNumber

`func (o *PurchaseOrder) SetVendorInvoiceNumber(v string)`

SetVendorInvoiceNumber sets VendorInvoiceNumber field to given value.

### HasVendorInvoiceNumber

`func (o *PurchaseOrder) HasVendorInvoiceNumber() bool`

HasVendorInvoiceNumber returns a boolean if a field has been set.

### GetVendorOrderNumber

`func (o *PurchaseOrder) GetVendorOrderNumber() string`

GetVendorOrderNumber returns the VendorOrderNumber field if non-nil, zero value otherwise.

### GetVendorOrderNumberOk

`func (o *PurchaseOrder) GetVendorOrderNumberOk() (*string, bool)`

GetVendorOrderNumberOk returns a tuple with the VendorOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorOrderNumber

`func (o *PurchaseOrder) SetVendorOrderNumber(v string)`

SetVendorOrderNumber sets VendorOrderNumber field to given value.

### HasVendorOrderNumber

`func (o *PurchaseOrder) HasVendorOrderNumber() bool`

HasVendorOrderNumber returns a boolean if a field has been set.

### GetVendorSite

`func (o *PurchaseOrder) GetVendorSite() SiteReference`

GetVendorSite returns the VendorSite field if non-nil, zero value otherwise.

### GetVendorSiteOk

`func (o *PurchaseOrder) GetVendorSiteOk() (*SiteReference, bool)`

GetVendorSiteOk returns a tuple with the VendorSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSite

`func (o *PurchaseOrder) SetVendorSite(v SiteReference)`

SetVendorSite sets VendorSite field to given value.

### HasVendorSite

`func (o *PurchaseOrder) HasVendorSite() bool`

HasVendorSite returns a boolean if a field has been set.

### GetWarehouse

`func (o *PurchaseOrder) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *PurchaseOrder) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *PurchaseOrder) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *PurchaseOrder) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetWarehouseContact

`func (o *PurchaseOrder) GetWarehouseContact() ContactReference`

GetWarehouseContact returns the WarehouseContact field if non-nil, zero value otherwise.

### GetWarehouseContactOk

`func (o *PurchaseOrder) GetWarehouseContactOk() (*ContactReference, bool)`

GetWarehouseContactOk returns a tuple with the WarehouseContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseContact

`func (o *PurchaseOrder) SetWarehouseContact(v ContactReference)`

SetWarehouseContact sets WarehouseContact field to given value.

### HasWarehouseContact

`func (o *PurchaseOrder) HasWarehouseContact() bool`

HasWarehouseContact returns a boolean if a field has been set.

### GetCurrency

`func (o *PurchaseOrder) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PurchaseOrder) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PurchaseOrder) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PurchaseOrder) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInfo

`func (o *PurchaseOrder) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PurchaseOrder) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PurchaseOrder) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PurchaseOrder) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *PurchaseOrder) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PurchaseOrder) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PurchaseOrder) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PurchaseOrder) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


