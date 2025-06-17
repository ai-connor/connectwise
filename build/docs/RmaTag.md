# RmaTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ServiceTicket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**SalesOrder** | Pointer to [**SalesOrderReference**](SalesOrderReference.md) |  | [optional] 
**Invoice** | Pointer to [**InvoiceReference**](InvoiceReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**Summary** | Pointer to **string** |  Max length: 150; | [optional] 
**Product** | [**IvItemReference**](IvItemReference.md) |  | 
**IvDescription** | Pointer to **string** |  | [optional] 
**ProductDescription** | **string** |  Max length: 200; | 
**SerialNumber** | Pointer to **string** |  | [optional] 
**MfgItemID** | Pointer to **string** |  Max length: 100; | [optional] 
**Status** | [**RmaStatusReference**](RmaStatusReference.md) |  | 
**ListPrice** | Pointer to **float64** |  | [optional] 
**UnitPrice** | Pointer to **NullableFloat64** |  | [optional] 
**Location** | [**SystemLocationReference**](SystemLocationReference.md) |  | 
**Department** | [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | 
**ProblemDescription** | Pointer to **string** |  Max length: 1000; | [optional] 
**ReturnedCompany** | [**CompanyReference**](CompanyReference.md) |  | 
**ReturnedContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**ReturnedContactType** | Pointer to **string** |  | [optional] 
**ReturnedContactPhone** | Pointer to **string** |  | [optional] 
**ReturnedContactExtension** | Pointer to **string** |  | [optional] 
**ReturnedContactEmail** | Pointer to **string** |  | [optional] 
**ReturnedContactAddressLine1** | Pointer to **string** |  Max length: 50; | [optional] 
**ReturnedContactAddressLine2** | Pointer to **string** |  Max length: 50; | [optional] 
**ReturnedContactCity** | Pointer to **string** |  Max length: 50; | [optional] 
**ReturnedContactState** | Pointer to **string** |  Max length: 50; | [optional] 
**ReturnedContactZip** | Pointer to **string** |  Max length: 12; | [optional] 
**ReturnedContactCountry** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**RmaDisposition** | [**RmaDispositionReference**](RmaDispositionReference.md) |  | 
**ReturnedSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**PurchasedCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**PurchasedContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**PurchasedContactType** | Pointer to **string** |  | [optional] 
**PurchasedContactPhone** | Pointer to **string** |  | [optional] 
**PurchasedContactExtension** | Pointer to **string** |  | [optional] 
**PurchasedContactEmail** | Pointer to **string** |  | [optional] 
**PurchasedContactAddressLine1** | Pointer to **string** |  Max length: 50; | [optional] 
**PurchasedContactAddressLine2** | Pointer to **string** |  Max length: 50; | [optional] 
**PurchasedContactCity** | Pointer to **string** |  Max length: 50; | [optional] 
**PurchasedContactState** | Pointer to **string** |  Max length: 50; | [optional] 
**PurchasedContactZip** | Pointer to **string** |  Max length: 12; | [optional] 
**PurchasedContactCountry** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**PurchasedInvoiceNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**PurchasedInvoiceDate** | Pointer to **string** |  | [optional] 
**PurchasedOrderNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**PurchasedVendorAction** | Pointer to [**RmaActionReference**](RmaActionReference.md) |  | [optional] 
**PurchasedVendorRmaNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**PurchasedSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**PurchasedNotes** | Pointer to **string** |  Max length: 1000; | [optional] 
**WarrantyCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**WarrantyContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**WarrantyContactType** | Pointer to **string** |  | [optional] 
**WarrantyContactPhone** | Pointer to **string** |  | [optional] 
**WarrantyContactEmail** | Pointer to **string** |  | [optional] 
**WarrantyContactExtension** | Pointer to **string** |  | [optional] 
**WarrantyContactAddressLine1** | Pointer to **string** |  Max length: 50; | [optional] 
**WarrantyContactAddressLine2** | Pointer to **string** |  Max length: 50; | [optional] 
**WarrantyContactCity** | Pointer to **string** |  Max length: 50; | [optional] 
**WarrantyContactState** | Pointer to **string** |  Max length: 50; | [optional] 
**WarrantyContactZip** | Pointer to **string** |  Max length: 12; | [optional] 
**WarrantyContactCountry** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**WarrantySite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**WarrantyNotes** | Pointer to **string** |  Max length: 1000; | [optional] 
**RepairCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**RepairContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**RepairContactType** | Pointer to **string** |  | [optional] 
**RepairContactPhone** | Pointer to **string** |  | [optional] 
**RepairContactExtension** | Pointer to **string** |  | [optional] 
**RepairContactEmail** | Pointer to **string** |  | [optional] 
**RepairContactAddressLine1** | Pointer to **string** |  Max length: 50; | [optional] 
**RepairContactAddressLine2** | Pointer to **string** |  Max length: 50; | [optional] 
**RepairContactCity** | Pointer to **string** |  Max length: 50; | [optional] 
**RepairContactState** | Pointer to **string** |  Max length: 50; | [optional] 
**RepairContactZip** | Pointer to **string** |  Max length: 12; | [optional] 
**RepairContactCountry** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**RepairOrderNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**RepairSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**RepairNotes** | Pointer to **string** |  Max length: 1000; | [optional] 
**DropShipFlag** | Pointer to **NullableBool** |  | [optional] 
**ShipMethod** | Pointer to [**ShipmentMethodReference**](ShipmentMethodReference.md) |  | [optional] 
**ShippingDate** | Pointer to **string** |  | [optional] 
**ShippingTrackingNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**InternalNotes** | Pointer to **string** |  Max length: 1000; | [optional] 
**ClosingNotes** | Pointer to **string** |  Max length: 1000; | [optional] 
**DateClosed** | Pointer to **string** |  | [optional] 
**AccountManager** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**TechnicalContact** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**ClosedBy** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewRmaTag

`func NewRmaTag(product IvItemReference, productDescription string, status RmaStatusReference, location SystemLocationReference, department SystemDepartmentReference, returnedCompany CompanyReference, rmaDisposition RmaDispositionReference, ) *RmaTag`

NewRmaTag instantiates a new RmaTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRmaTagWithDefaults

`func NewRmaTagWithDefaults() *RmaTag`

NewRmaTagWithDefaults instantiates a new RmaTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RmaTag) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RmaTag) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RmaTag) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RmaTag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServiceTicket

`func (o *RmaTag) GetServiceTicket() TicketReference`

GetServiceTicket returns the ServiceTicket field if non-nil, zero value otherwise.

### GetServiceTicketOk

`func (o *RmaTag) GetServiceTicketOk() (*TicketReference, bool)`

GetServiceTicketOk returns a tuple with the ServiceTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTicket

`func (o *RmaTag) SetServiceTicket(v TicketReference)`

SetServiceTicket sets ServiceTicket field to given value.

### HasServiceTicket

`func (o *RmaTag) HasServiceTicket() bool`

HasServiceTicket returns a boolean if a field has been set.

### GetSalesOrder

`func (o *RmaTag) GetSalesOrder() SalesOrderReference`

GetSalesOrder returns the SalesOrder field if non-nil, zero value otherwise.

### GetSalesOrderOk

`func (o *RmaTag) GetSalesOrderOk() (*SalesOrderReference, bool)`

GetSalesOrderOk returns a tuple with the SalesOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrder

`func (o *RmaTag) SetSalesOrder(v SalesOrderReference)`

SetSalesOrder sets SalesOrder field to given value.

### HasSalesOrder

`func (o *RmaTag) HasSalesOrder() bool`

HasSalesOrder returns a boolean if a field has been set.

### GetInvoice

`func (o *RmaTag) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *RmaTag) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *RmaTag) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *RmaTag) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetProject

`func (o *RmaTag) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RmaTag) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RmaTag) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *RmaTag) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSummary

`func (o *RmaTag) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RmaTag) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RmaTag) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RmaTag) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetProduct

`func (o *RmaTag) GetProduct() IvItemReference`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *RmaTag) GetProductOk() (*IvItemReference, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *RmaTag) SetProduct(v IvItemReference)`

SetProduct sets Product field to given value.


### GetIvDescription

`func (o *RmaTag) GetIvDescription() string`

GetIvDescription returns the IvDescription field if non-nil, zero value otherwise.

### GetIvDescriptionOk

`func (o *RmaTag) GetIvDescriptionOk() (*string, bool)`

GetIvDescriptionOk returns a tuple with the IvDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIvDescription

`func (o *RmaTag) SetIvDescription(v string)`

SetIvDescription sets IvDescription field to given value.

### HasIvDescription

`func (o *RmaTag) HasIvDescription() bool`

HasIvDescription returns a boolean if a field has been set.

### GetProductDescription

`func (o *RmaTag) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *RmaTag) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *RmaTag) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.


### GetSerialNumber

`func (o *RmaTag) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *RmaTag) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *RmaTag) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *RmaTag) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetMfgItemID

`func (o *RmaTag) GetMfgItemID() string`

GetMfgItemID returns the MfgItemID field if non-nil, zero value otherwise.

### GetMfgItemIDOk

`func (o *RmaTag) GetMfgItemIDOk() (*string, bool)`

GetMfgItemIDOk returns a tuple with the MfgItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgItemID

`func (o *RmaTag) SetMfgItemID(v string)`

SetMfgItemID sets MfgItemID field to given value.

### HasMfgItemID

`func (o *RmaTag) HasMfgItemID() bool`

HasMfgItemID returns a boolean if a field has been set.

### GetStatus

`func (o *RmaTag) GetStatus() RmaStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RmaTag) GetStatusOk() (*RmaStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RmaTag) SetStatus(v RmaStatusReference)`

SetStatus sets Status field to given value.


### GetListPrice

`func (o *RmaTag) GetListPrice() float64`

GetListPrice returns the ListPrice field if non-nil, zero value otherwise.

### GetListPriceOk

`func (o *RmaTag) GetListPriceOk() (*float64, bool)`

GetListPriceOk returns a tuple with the ListPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListPrice

`func (o *RmaTag) SetListPrice(v float64)`

SetListPrice sets ListPrice field to given value.

### HasListPrice

`func (o *RmaTag) HasListPrice() bool`

HasListPrice returns a boolean if a field has been set.

### GetUnitPrice

`func (o *RmaTag) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *RmaTag) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *RmaTag) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *RmaTag) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *RmaTag) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *RmaTag) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetLocation

`func (o *RmaTag) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *RmaTag) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *RmaTag) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.


### GetDepartment

`func (o *RmaTag) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *RmaTag) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *RmaTag) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.


### GetProblemDescription

`func (o *RmaTag) GetProblemDescription() string`

GetProblemDescription returns the ProblemDescription field if non-nil, zero value otherwise.

### GetProblemDescriptionOk

`func (o *RmaTag) GetProblemDescriptionOk() (*string, bool)`

GetProblemDescriptionOk returns a tuple with the ProblemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemDescription

`func (o *RmaTag) SetProblemDescription(v string)`

SetProblemDescription sets ProblemDescription field to given value.

### HasProblemDescription

`func (o *RmaTag) HasProblemDescription() bool`

HasProblemDescription returns a boolean if a field has been set.

### GetReturnedCompany

`func (o *RmaTag) GetReturnedCompany() CompanyReference`

GetReturnedCompany returns the ReturnedCompany field if non-nil, zero value otherwise.

### GetReturnedCompanyOk

`func (o *RmaTag) GetReturnedCompanyOk() (*CompanyReference, bool)`

GetReturnedCompanyOk returns a tuple with the ReturnedCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedCompany

`func (o *RmaTag) SetReturnedCompany(v CompanyReference)`

SetReturnedCompany sets ReturnedCompany field to given value.


### GetReturnedContact

`func (o *RmaTag) GetReturnedContact() ContactReference`

GetReturnedContact returns the ReturnedContact field if non-nil, zero value otherwise.

### GetReturnedContactOk

`func (o *RmaTag) GetReturnedContactOk() (*ContactReference, bool)`

GetReturnedContactOk returns a tuple with the ReturnedContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContact

`func (o *RmaTag) SetReturnedContact(v ContactReference)`

SetReturnedContact sets ReturnedContact field to given value.

### HasReturnedContact

`func (o *RmaTag) HasReturnedContact() bool`

HasReturnedContact returns a boolean if a field has been set.

### GetReturnedContactType

`func (o *RmaTag) GetReturnedContactType() string`

GetReturnedContactType returns the ReturnedContactType field if non-nil, zero value otherwise.

### GetReturnedContactTypeOk

`func (o *RmaTag) GetReturnedContactTypeOk() (*string, bool)`

GetReturnedContactTypeOk returns a tuple with the ReturnedContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContactType

`func (o *RmaTag) SetReturnedContactType(v string)`

SetReturnedContactType sets ReturnedContactType field to given value.

### HasReturnedContactType

`func (o *RmaTag) HasReturnedContactType() bool`

HasReturnedContactType returns a boolean if a field has been set.

### GetReturnedContactPhone

`func (o *RmaTag) GetReturnedContactPhone() string`

GetReturnedContactPhone returns the ReturnedContactPhone field if non-nil, zero value otherwise.

### GetReturnedContactPhoneOk

`func (o *RmaTag) GetReturnedContactPhoneOk() (*string, bool)`

GetReturnedContactPhoneOk returns a tuple with the ReturnedContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContactPhone

`func (o *RmaTag) SetReturnedContactPhone(v string)`

SetReturnedContactPhone sets ReturnedContactPhone field to given value.

### HasReturnedContactPhone

`func (o *RmaTag) HasReturnedContactPhone() bool`

HasReturnedContactPhone returns a boolean if a field has been set.

### GetReturnedContactExtension

`func (o *RmaTag) GetReturnedContactExtension() string`

GetReturnedContactExtension returns the ReturnedContactExtension field if non-nil, zero value otherwise.

### GetReturnedContactExtensionOk

`func (o *RmaTag) GetReturnedContactExtensionOk() (*string, bool)`

GetReturnedContactExtensionOk returns a tuple with the ReturnedContactExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContactExtension

`func (o *RmaTag) SetReturnedContactExtension(v string)`

SetReturnedContactExtension sets ReturnedContactExtension field to given value.

### HasReturnedContactExtension

`func (o *RmaTag) HasReturnedContactExtension() bool`

HasReturnedContactExtension returns a boolean if a field has been set.

### GetReturnedContactEmail

`func (o *RmaTag) GetReturnedContactEmail() string`

GetReturnedContactEmail returns the ReturnedContactEmail field if non-nil, zero value otherwise.

### GetReturnedContactEmailOk

`func (o *RmaTag) GetReturnedContactEmailOk() (*string, bool)`

GetReturnedContactEmailOk returns a tuple with the ReturnedContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContactEmail

`func (o *RmaTag) SetReturnedContactEmail(v string)`

SetReturnedContactEmail sets ReturnedContactEmail field to given value.

### HasReturnedContactEmail

`func (o *RmaTag) HasReturnedContactEmail() bool`

HasReturnedContactEmail returns a boolean if a field has been set.

### GetReturnedContactAddressLine1

`func (o *RmaTag) GetReturnedContactAddressLine1() string`

GetReturnedContactAddressLine1 returns the ReturnedContactAddressLine1 field if non-nil, zero value otherwise.

### GetReturnedContactAddressLine1Ok

`func (o *RmaTag) GetReturnedContactAddressLine1Ok() (*string, bool)`

GetReturnedContactAddressLine1Ok returns a tuple with the ReturnedContactAddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContactAddressLine1

`func (o *RmaTag) SetReturnedContactAddressLine1(v string)`

SetReturnedContactAddressLine1 sets ReturnedContactAddressLine1 field to given value.

### HasReturnedContactAddressLine1

`func (o *RmaTag) HasReturnedContactAddressLine1() bool`

HasReturnedContactAddressLine1 returns a boolean if a field has been set.

### GetReturnedContactAddressLine2

`func (o *RmaTag) GetReturnedContactAddressLine2() string`

GetReturnedContactAddressLine2 returns the ReturnedContactAddressLine2 field if non-nil, zero value otherwise.

### GetReturnedContactAddressLine2Ok

`func (o *RmaTag) GetReturnedContactAddressLine2Ok() (*string, bool)`

GetReturnedContactAddressLine2Ok returns a tuple with the ReturnedContactAddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContactAddressLine2

`func (o *RmaTag) SetReturnedContactAddressLine2(v string)`

SetReturnedContactAddressLine2 sets ReturnedContactAddressLine2 field to given value.

### HasReturnedContactAddressLine2

`func (o *RmaTag) HasReturnedContactAddressLine2() bool`

HasReturnedContactAddressLine2 returns a boolean if a field has been set.

### GetReturnedContactCity

`func (o *RmaTag) GetReturnedContactCity() string`

GetReturnedContactCity returns the ReturnedContactCity field if non-nil, zero value otherwise.

### GetReturnedContactCityOk

`func (o *RmaTag) GetReturnedContactCityOk() (*string, bool)`

GetReturnedContactCityOk returns a tuple with the ReturnedContactCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContactCity

`func (o *RmaTag) SetReturnedContactCity(v string)`

SetReturnedContactCity sets ReturnedContactCity field to given value.

### HasReturnedContactCity

`func (o *RmaTag) HasReturnedContactCity() bool`

HasReturnedContactCity returns a boolean if a field has been set.

### GetReturnedContactState

`func (o *RmaTag) GetReturnedContactState() string`

GetReturnedContactState returns the ReturnedContactState field if non-nil, zero value otherwise.

### GetReturnedContactStateOk

`func (o *RmaTag) GetReturnedContactStateOk() (*string, bool)`

GetReturnedContactStateOk returns a tuple with the ReturnedContactState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContactState

`func (o *RmaTag) SetReturnedContactState(v string)`

SetReturnedContactState sets ReturnedContactState field to given value.

### HasReturnedContactState

`func (o *RmaTag) HasReturnedContactState() bool`

HasReturnedContactState returns a boolean if a field has been set.

### GetReturnedContactZip

`func (o *RmaTag) GetReturnedContactZip() string`

GetReturnedContactZip returns the ReturnedContactZip field if non-nil, zero value otherwise.

### GetReturnedContactZipOk

`func (o *RmaTag) GetReturnedContactZipOk() (*string, bool)`

GetReturnedContactZipOk returns a tuple with the ReturnedContactZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContactZip

`func (o *RmaTag) SetReturnedContactZip(v string)`

SetReturnedContactZip sets ReturnedContactZip field to given value.

### HasReturnedContactZip

`func (o *RmaTag) HasReturnedContactZip() bool`

HasReturnedContactZip returns a boolean if a field has been set.

### GetReturnedContactCountry

`func (o *RmaTag) GetReturnedContactCountry() CountryReference`

GetReturnedContactCountry returns the ReturnedContactCountry field if non-nil, zero value otherwise.

### GetReturnedContactCountryOk

`func (o *RmaTag) GetReturnedContactCountryOk() (*CountryReference, bool)`

GetReturnedContactCountryOk returns a tuple with the ReturnedContactCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContactCountry

`func (o *RmaTag) SetReturnedContactCountry(v CountryReference)`

SetReturnedContactCountry sets ReturnedContactCountry field to given value.

### HasReturnedContactCountry

`func (o *RmaTag) HasReturnedContactCountry() bool`

HasReturnedContactCountry returns a boolean if a field has been set.

### GetRmaDisposition

`func (o *RmaTag) GetRmaDisposition() RmaDispositionReference`

GetRmaDisposition returns the RmaDisposition field if non-nil, zero value otherwise.

### GetRmaDispositionOk

`func (o *RmaTag) GetRmaDispositionOk() (*RmaDispositionReference, bool)`

GetRmaDispositionOk returns a tuple with the RmaDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmaDisposition

`func (o *RmaTag) SetRmaDisposition(v RmaDispositionReference)`

SetRmaDisposition sets RmaDisposition field to given value.


### GetReturnedSite

`func (o *RmaTag) GetReturnedSite() SiteReference`

GetReturnedSite returns the ReturnedSite field if non-nil, zero value otherwise.

### GetReturnedSiteOk

`func (o *RmaTag) GetReturnedSiteOk() (*SiteReference, bool)`

GetReturnedSiteOk returns a tuple with the ReturnedSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedSite

`func (o *RmaTag) SetReturnedSite(v SiteReference)`

SetReturnedSite sets ReturnedSite field to given value.

### HasReturnedSite

`func (o *RmaTag) HasReturnedSite() bool`

HasReturnedSite returns a boolean if a field has been set.

### GetPurchasedCompany

`func (o *RmaTag) GetPurchasedCompany() CompanyReference`

GetPurchasedCompany returns the PurchasedCompany field if non-nil, zero value otherwise.

### GetPurchasedCompanyOk

`func (o *RmaTag) GetPurchasedCompanyOk() (*CompanyReference, bool)`

GetPurchasedCompanyOk returns a tuple with the PurchasedCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedCompany

`func (o *RmaTag) SetPurchasedCompany(v CompanyReference)`

SetPurchasedCompany sets PurchasedCompany field to given value.

### HasPurchasedCompany

`func (o *RmaTag) HasPurchasedCompany() bool`

HasPurchasedCompany returns a boolean if a field has been set.

### GetPurchasedContact

`func (o *RmaTag) GetPurchasedContact() ContactReference`

GetPurchasedContact returns the PurchasedContact field if non-nil, zero value otherwise.

### GetPurchasedContactOk

`func (o *RmaTag) GetPurchasedContactOk() (*ContactReference, bool)`

GetPurchasedContactOk returns a tuple with the PurchasedContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedContact

`func (o *RmaTag) SetPurchasedContact(v ContactReference)`

SetPurchasedContact sets PurchasedContact field to given value.

### HasPurchasedContact

`func (o *RmaTag) HasPurchasedContact() bool`

HasPurchasedContact returns a boolean if a field has been set.

### GetPurchasedContactType

`func (o *RmaTag) GetPurchasedContactType() string`

GetPurchasedContactType returns the PurchasedContactType field if non-nil, zero value otherwise.

### GetPurchasedContactTypeOk

`func (o *RmaTag) GetPurchasedContactTypeOk() (*string, bool)`

GetPurchasedContactTypeOk returns a tuple with the PurchasedContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedContactType

`func (o *RmaTag) SetPurchasedContactType(v string)`

SetPurchasedContactType sets PurchasedContactType field to given value.

### HasPurchasedContactType

`func (o *RmaTag) HasPurchasedContactType() bool`

HasPurchasedContactType returns a boolean if a field has been set.

### GetPurchasedContactPhone

`func (o *RmaTag) GetPurchasedContactPhone() string`

GetPurchasedContactPhone returns the PurchasedContactPhone field if non-nil, zero value otherwise.

### GetPurchasedContactPhoneOk

`func (o *RmaTag) GetPurchasedContactPhoneOk() (*string, bool)`

GetPurchasedContactPhoneOk returns a tuple with the PurchasedContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedContactPhone

`func (o *RmaTag) SetPurchasedContactPhone(v string)`

SetPurchasedContactPhone sets PurchasedContactPhone field to given value.

### HasPurchasedContactPhone

`func (o *RmaTag) HasPurchasedContactPhone() bool`

HasPurchasedContactPhone returns a boolean if a field has been set.

### GetPurchasedContactExtension

`func (o *RmaTag) GetPurchasedContactExtension() string`

GetPurchasedContactExtension returns the PurchasedContactExtension field if non-nil, zero value otherwise.

### GetPurchasedContactExtensionOk

`func (o *RmaTag) GetPurchasedContactExtensionOk() (*string, bool)`

GetPurchasedContactExtensionOk returns a tuple with the PurchasedContactExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedContactExtension

`func (o *RmaTag) SetPurchasedContactExtension(v string)`

SetPurchasedContactExtension sets PurchasedContactExtension field to given value.

### HasPurchasedContactExtension

`func (o *RmaTag) HasPurchasedContactExtension() bool`

HasPurchasedContactExtension returns a boolean if a field has been set.

### GetPurchasedContactEmail

`func (o *RmaTag) GetPurchasedContactEmail() string`

GetPurchasedContactEmail returns the PurchasedContactEmail field if non-nil, zero value otherwise.

### GetPurchasedContactEmailOk

`func (o *RmaTag) GetPurchasedContactEmailOk() (*string, bool)`

GetPurchasedContactEmailOk returns a tuple with the PurchasedContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedContactEmail

`func (o *RmaTag) SetPurchasedContactEmail(v string)`

SetPurchasedContactEmail sets PurchasedContactEmail field to given value.

### HasPurchasedContactEmail

`func (o *RmaTag) HasPurchasedContactEmail() bool`

HasPurchasedContactEmail returns a boolean if a field has been set.

### GetPurchasedContactAddressLine1

`func (o *RmaTag) GetPurchasedContactAddressLine1() string`

GetPurchasedContactAddressLine1 returns the PurchasedContactAddressLine1 field if non-nil, zero value otherwise.

### GetPurchasedContactAddressLine1Ok

`func (o *RmaTag) GetPurchasedContactAddressLine1Ok() (*string, bool)`

GetPurchasedContactAddressLine1Ok returns a tuple with the PurchasedContactAddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedContactAddressLine1

`func (o *RmaTag) SetPurchasedContactAddressLine1(v string)`

SetPurchasedContactAddressLine1 sets PurchasedContactAddressLine1 field to given value.

### HasPurchasedContactAddressLine1

`func (o *RmaTag) HasPurchasedContactAddressLine1() bool`

HasPurchasedContactAddressLine1 returns a boolean if a field has been set.

### GetPurchasedContactAddressLine2

`func (o *RmaTag) GetPurchasedContactAddressLine2() string`

GetPurchasedContactAddressLine2 returns the PurchasedContactAddressLine2 field if non-nil, zero value otherwise.

### GetPurchasedContactAddressLine2Ok

`func (o *RmaTag) GetPurchasedContactAddressLine2Ok() (*string, bool)`

GetPurchasedContactAddressLine2Ok returns a tuple with the PurchasedContactAddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedContactAddressLine2

`func (o *RmaTag) SetPurchasedContactAddressLine2(v string)`

SetPurchasedContactAddressLine2 sets PurchasedContactAddressLine2 field to given value.

### HasPurchasedContactAddressLine2

`func (o *RmaTag) HasPurchasedContactAddressLine2() bool`

HasPurchasedContactAddressLine2 returns a boolean if a field has been set.

### GetPurchasedContactCity

`func (o *RmaTag) GetPurchasedContactCity() string`

GetPurchasedContactCity returns the PurchasedContactCity field if non-nil, zero value otherwise.

### GetPurchasedContactCityOk

`func (o *RmaTag) GetPurchasedContactCityOk() (*string, bool)`

GetPurchasedContactCityOk returns a tuple with the PurchasedContactCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedContactCity

`func (o *RmaTag) SetPurchasedContactCity(v string)`

SetPurchasedContactCity sets PurchasedContactCity field to given value.

### HasPurchasedContactCity

`func (o *RmaTag) HasPurchasedContactCity() bool`

HasPurchasedContactCity returns a boolean if a field has been set.

### GetPurchasedContactState

`func (o *RmaTag) GetPurchasedContactState() string`

GetPurchasedContactState returns the PurchasedContactState field if non-nil, zero value otherwise.

### GetPurchasedContactStateOk

`func (o *RmaTag) GetPurchasedContactStateOk() (*string, bool)`

GetPurchasedContactStateOk returns a tuple with the PurchasedContactState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedContactState

`func (o *RmaTag) SetPurchasedContactState(v string)`

SetPurchasedContactState sets PurchasedContactState field to given value.

### HasPurchasedContactState

`func (o *RmaTag) HasPurchasedContactState() bool`

HasPurchasedContactState returns a boolean if a field has been set.

### GetPurchasedContactZip

`func (o *RmaTag) GetPurchasedContactZip() string`

GetPurchasedContactZip returns the PurchasedContactZip field if non-nil, zero value otherwise.

### GetPurchasedContactZipOk

`func (o *RmaTag) GetPurchasedContactZipOk() (*string, bool)`

GetPurchasedContactZipOk returns a tuple with the PurchasedContactZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedContactZip

`func (o *RmaTag) SetPurchasedContactZip(v string)`

SetPurchasedContactZip sets PurchasedContactZip field to given value.

### HasPurchasedContactZip

`func (o *RmaTag) HasPurchasedContactZip() bool`

HasPurchasedContactZip returns a boolean if a field has been set.

### GetPurchasedContactCountry

`func (o *RmaTag) GetPurchasedContactCountry() CountryReference`

GetPurchasedContactCountry returns the PurchasedContactCountry field if non-nil, zero value otherwise.

### GetPurchasedContactCountryOk

`func (o *RmaTag) GetPurchasedContactCountryOk() (*CountryReference, bool)`

GetPurchasedContactCountryOk returns a tuple with the PurchasedContactCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedContactCountry

`func (o *RmaTag) SetPurchasedContactCountry(v CountryReference)`

SetPurchasedContactCountry sets PurchasedContactCountry field to given value.

### HasPurchasedContactCountry

`func (o *RmaTag) HasPurchasedContactCountry() bool`

HasPurchasedContactCountry returns a boolean if a field has been set.

### GetPurchasedInvoiceNumber

`func (o *RmaTag) GetPurchasedInvoiceNumber() string`

GetPurchasedInvoiceNumber returns the PurchasedInvoiceNumber field if non-nil, zero value otherwise.

### GetPurchasedInvoiceNumberOk

`func (o *RmaTag) GetPurchasedInvoiceNumberOk() (*string, bool)`

GetPurchasedInvoiceNumberOk returns a tuple with the PurchasedInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedInvoiceNumber

`func (o *RmaTag) SetPurchasedInvoiceNumber(v string)`

SetPurchasedInvoiceNumber sets PurchasedInvoiceNumber field to given value.

### HasPurchasedInvoiceNumber

`func (o *RmaTag) HasPurchasedInvoiceNumber() bool`

HasPurchasedInvoiceNumber returns a boolean if a field has been set.

### GetPurchasedInvoiceDate

`func (o *RmaTag) GetPurchasedInvoiceDate() string`

GetPurchasedInvoiceDate returns the PurchasedInvoiceDate field if non-nil, zero value otherwise.

### GetPurchasedInvoiceDateOk

`func (o *RmaTag) GetPurchasedInvoiceDateOk() (*string, bool)`

GetPurchasedInvoiceDateOk returns a tuple with the PurchasedInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedInvoiceDate

`func (o *RmaTag) SetPurchasedInvoiceDate(v string)`

SetPurchasedInvoiceDate sets PurchasedInvoiceDate field to given value.

### HasPurchasedInvoiceDate

`func (o *RmaTag) HasPurchasedInvoiceDate() bool`

HasPurchasedInvoiceDate returns a boolean if a field has been set.

### GetPurchasedOrderNumber

`func (o *RmaTag) GetPurchasedOrderNumber() string`

GetPurchasedOrderNumber returns the PurchasedOrderNumber field if non-nil, zero value otherwise.

### GetPurchasedOrderNumberOk

`func (o *RmaTag) GetPurchasedOrderNumberOk() (*string, bool)`

GetPurchasedOrderNumberOk returns a tuple with the PurchasedOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedOrderNumber

`func (o *RmaTag) SetPurchasedOrderNumber(v string)`

SetPurchasedOrderNumber sets PurchasedOrderNumber field to given value.

### HasPurchasedOrderNumber

`func (o *RmaTag) HasPurchasedOrderNumber() bool`

HasPurchasedOrderNumber returns a boolean if a field has been set.

### GetPurchasedVendorAction

`func (o *RmaTag) GetPurchasedVendorAction() RmaActionReference`

GetPurchasedVendorAction returns the PurchasedVendorAction field if non-nil, zero value otherwise.

### GetPurchasedVendorActionOk

`func (o *RmaTag) GetPurchasedVendorActionOk() (*RmaActionReference, bool)`

GetPurchasedVendorActionOk returns a tuple with the PurchasedVendorAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedVendorAction

`func (o *RmaTag) SetPurchasedVendorAction(v RmaActionReference)`

SetPurchasedVendorAction sets PurchasedVendorAction field to given value.

### HasPurchasedVendorAction

`func (o *RmaTag) HasPurchasedVendorAction() bool`

HasPurchasedVendorAction returns a boolean if a field has been set.

### GetPurchasedVendorRmaNumber

`func (o *RmaTag) GetPurchasedVendorRmaNumber() string`

GetPurchasedVendorRmaNumber returns the PurchasedVendorRmaNumber field if non-nil, zero value otherwise.

### GetPurchasedVendorRmaNumberOk

`func (o *RmaTag) GetPurchasedVendorRmaNumberOk() (*string, bool)`

GetPurchasedVendorRmaNumberOk returns a tuple with the PurchasedVendorRmaNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedVendorRmaNumber

`func (o *RmaTag) SetPurchasedVendorRmaNumber(v string)`

SetPurchasedVendorRmaNumber sets PurchasedVendorRmaNumber field to given value.

### HasPurchasedVendorRmaNumber

`func (o *RmaTag) HasPurchasedVendorRmaNumber() bool`

HasPurchasedVendorRmaNumber returns a boolean if a field has been set.

### GetPurchasedSite

`func (o *RmaTag) GetPurchasedSite() SiteReference`

GetPurchasedSite returns the PurchasedSite field if non-nil, zero value otherwise.

### GetPurchasedSiteOk

`func (o *RmaTag) GetPurchasedSiteOk() (*SiteReference, bool)`

GetPurchasedSiteOk returns a tuple with the PurchasedSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedSite

`func (o *RmaTag) SetPurchasedSite(v SiteReference)`

SetPurchasedSite sets PurchasedSite field to given value.

### HasPurchasedSite

`func (o *RmaTag) HasPurchasedSite() bool`

HasPurchasedSite returns a boolean if a field has been set.

### GetPurchasedNotes

`func (o *RmaTag) GetPurchasedNotes() string`

GetPurchasedNotes returns the PurchasedNotes field if non-nil, zero value otherwise.

### GetPurchasedNotesOk

`func (o *RmaTag) GetPurchasedNotesOk() (*string, bool)`

GetPurchasedNotesOk returns a tuple with the PurchasedNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedNotes

`func (o *RmaTag) SetPurchasedNotes(v string)`

SetPurchasedNotes sets PurchasedNotes field to given value.

### HasPurchasedNotes

`func (o *RmaTag) HasPurchasedNotes() bool`

HasPurchasedNotes returns a boolean if a field has been set.

### GetWarrantyCompany

`func (o *RmaTag) GetWarrantyCompany() CompanyReference`

GetWarrantyCompany returns the WarrantyCompany field if non-nil, zero value otherwise.

### GetWarrantyCompanyOk

`func (o *RmaTag) GetWarrantyCompanyOk() (*CompanyReference, bool)`

GetWarrantyCompanyOk returns a tuple with the WarrantyCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyCompany

`func (o *RmaTag) SetWarrantyCompany(v CompanyReference)`

SetWarrantyCompany sets WarrantyCompany field to given value.

### HasWarrantyCompany

`func (o *RmaTag) HasWarrantyCompany() bool`

HasWarrantyCompany returns a boolean if a field has been set.

### GetWarrantyContact

`func (o *RmaTag) GetWarrantyContact() ContactReference`

GetWarrantyContact returns the WarrantyContact field if non-nil, zero value otherwise.

### GetWarrantyContactOk

`func (o *RmaTag) GetWarrantyContactOk() (*ContactReference, bool)`

GetWarrantyContactOk returns a tuple with the WarrantyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyContact

`func (o *RmaTag) SetWarrantyContact(v ContactReference)`

SetWarrantyContact sets WarrantyContact field to given value.

### HasWarrantyContact

`func (o *RmaTag) HasWarrantyContact() bool`

HasWarrantyContact returns a boolean if a field has been set.

### GetWarrantyContactType

`func (o *RmaTag) GetWarrantyContactType() string`

GetWarrantyContactType returns the WarrantyContactType field if non-nil, zero value otherwise.

### GetWarrantyContactTypeOk

`func (o *RmaTag) GetWarrantyContactTypeOk() (*string, bool)`

GetWarrantyContactTypeOk returns a tuple with the WarrantyContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyContactType

`func (o *RmaTag) SetWarrantyContactType(v string)`

SetWarrantyContactType sets WarrantyContactType field to given value.

### HasWarrantyContactType

`func (o *RmaTag) HasWarrantyContactType() bool`

HasWarrantyContactType returns a boolean if a field has been set.

### GetWarrantyContactPhone

`func (o *RmaTag) GetWarrantyContactPhone() string`

GetWarrantyContactPhone returns the WarrantyContactPhone field if non-nil, zero value otherwise.

### GetWarrantyContactPhoneOk

`func (o *RmaTag) GetWarrantyContactPhoneOk() (*string, bool)`

GetWarrantyContactPhoneOk returns a tuple with the WarrantyContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyContactPhone

`func (o *RmaTag) SetWarrantyContactPhone(v string)`

SetWarrantyContactPhone sets WarrantyContactPhone field to given value.

### HasWarrantyContactPhone

`func (o *RmaTag) HasWarrantyContactPhone() bool`

HasWarrantyContactPhone returns a boolean if a field has been set.

### GetWarrantyContactEmail

`func (o *RmaTag) GetWarrantyContactEmail() string`

GetWarrantyContactEmail returns the WarrantyContactEmail field if non-nil, zero value otherwise.

### GetWarrantyContactEmailOk

`func (o *RmaTag) GetWarrantyContactEmailOk() (*string, bool)`

GetWarrantyContactEmailOk returns a tuple with the WarrantyContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyContactEmail

`func (o *RmaTag) SetWarrantyContactEmail(v string)`

SetWarrantyContactEmail sets WarrantyContactEmail field to given value.

### HasWarrantyContactEmail

`func (o *RmaTag) HasWarrantyContactEmail() bool`

HasWarrantyContactEmail returns a boolean if a field has been set.

### GetWarrantyContactExtension

`func (o *RmaTag) GetWarrantyContactExtension() string`

GetWarrantyContactExtension returns the WarrantyContactExtension field if non-nil, zero value otherwise.

### GetWarrantyContactExtensionOk

`func (o *RmaTag) GetWarrantyContactExtensionOk() (*string, bool)`

GetWarrantyContactExtensionOk returns a tuple with the WarrantyContactExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyContactExtension

`func (o *RmaTag) SetWarrantyContactExtension(v string)`

SetWarrantyContactExtension sets WarrantyContactExtension field to given value.

### HasWarrantyContactExtension

`func (o *RmaTag) HasWarrantyContactExtension() bool`

HasWarrantyContactExtension returns a boolean if a field has been set.

### GetWarrantyContactAddressLine1

`func (o *RmaTag) GetWarrantyContactAddressLine1() string`

GetWarrantyContactAddressLine1 returns the WarrantyContactAddressLine1 field if non-nil, zero value otherwise.

### GetWarrantyContactAddressLine1Ok

`func (o *RmaTag) GetWarrantyContactAddressLine1Ok() (*string, bool)`

GetWarrantyContactAddressLine1Ok returns a tuple with the WarrantyContactAddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyContactAddressLine1

`func (o *RmaTag) SetWarrantyContactAddressLine1(v string)`

SetWarrantyContactAddressLine1 sets WarrantyContactAddressLine1 field to given value.

### HasWarrantyContactAddressLine1

`func (o *RmaTag) HasWarrantyContactAddressLine1() bool`

HasWarrantyContactAddressLine1 returns a boolean if a field has been set.

### GetWarrantyContactAddressLine2

`func (o *RmaTag) GetWarrantyContactAddressLine2() string`

GetWarrantyContactAddressLine2 returns the WarrantyContactAddressLine2 field if non-nil, zero value otherwise.

### GetWarrantyContactAddressLine2Ok

`func (o *RmaTag) GetWarrantyContactAddressLine2Ok() (*string, bool)`

GetWarrantyContactAddressLine2Ok returns a tuple with the WarrantyContactAddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyContactAddressLine2

`func (o *RmaTag) SetWarrantyContactAddressLine2(v string)`

SetWarrantyContactAddressLine2 sets WarrantyContactAddressLine2 field to given value.

### HasWarrantyContactAddressLine2

`func (o *RmaTag) HasWarrantyContactAddressLine2() bool`

HasWarrantyContactAddressLine2 returns a boolean if a field has been set.

### GetWarrantyContactCity

`func (o *RmaTag) GetWarrantyContactCity() string`

GetWarrantyContactCity returns the WarrantyContactCity field if non-nil, zero value otherwise.

### GetWarrantyContactCityOk

`func (o *RmaTag) GetWarrantyContactCityOk() (*string, bool)`

GetWarrantyContactCityOk returns a tuple with the WarrantyContactCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyContactCity

`func (o *RmaTag) SetWarrantyContactCity(v string)`

SetWarrantyContactCity sets WarrantyContactCity field to given value.

### HasWarrantyContactCity

`func (o *RmaTag) HasWarrantyContactCity() bool`

HasWarrantyContactCity returns a boolean if a field has been set.

### GetWarrantyContactState

`func (o *RmaTag) GetWarrantyContactState() string`

GetWarrantyContactState returns the WarrantyContactState field if non-nil, zero value otherwise.

### GetWarrantyContactStateOk

`func (o *RmaTag) GetWarrantyContactStateOk() (*string, bool)`

GetWarrantyContactStateOk returns a tuple with the WarrantyContactState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyContactState

`func (o *RmaTag) SetWarrantyContactState(v string)`

SetWarrantyContactState sets WarrantyContactState field to given value.

### HasWarrantyContactState

`func (o *RmaTag) HasWarrantyContactState() bool`

HasWarrantyContactState returns a boolean if a field has been set.

### GetWarrantyContactZip

`func (o *RmaTag) GetWarrantyContactZip() string`

GetWarrantyContactZip returns the WarrantyContactZip field if non-nil, zero value otherwise.

### GetWarrantyContactZipOk

`func (o *RmaTag) GetWarrantyContactZipOk() (*string, bool)`

GetWarrantyContactZipOk returns a tuple with the WarrantyContactZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyContactZip

`func (o *RmaTag) SetWarrantyContactZip(v string)`

SetWarrantyContactZip sets WarrantyContactZip field to given value.

### HasWarrantyContactZip

`func (o *RmaTag) HasWarrantyContactZip() bool`

HasWarrantyContactZip returns a boolean if a field has been set.

### GetWarrantyContactCountry

`func (o *RmaTag) GetWarrantyContactCountry() CountryReference`

GetWarrantyContactCountry returns the WarrantyContactCountry field if non-nil, zero value otherwise.

### GetWarrantyContactCountryOk

`func (o *RmaTag) GetWarrantyContactCountryOk() (*CountryReference, bool)`

GetWarrantyContactCountryOk returns a tuple with the WarrantyContactCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyContactCountry

`func (o *RmaTag) SetWarrantyContactCountry(v CountryReference)`

SetWarrantyContactCountry sets WarrantyContactCountry field to given value.

### HasWarrantyContactCountry

`func (o *RmaTag) HasWarrantyContactCountry() bool`

HasWarrantyContactCountry returns a boolean if a field has been set.

### GetWarrantySite

`func (o *RmaTag) GetWarrantySite() SiteReference`

GetWarrantySite returns the WarrantySite field if non-nil, zero value otherwise.

### GetWarrantySiteOk

`func (o *RmaTag) GetWarrantySiteOk() (*SiteReference, bool)`

GetWarrantySiteOk returns a tuple with the WarrantySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantySite

`func (o *RmaTag) SetWarrantySite(v SiteReference)`

SetWarrantySite sets WarrantySite field to given value.

### HasWarrantySite

`func (o *RmaTag) HasWarrantySite() bool`

HasWarrantySite returns a boolean if a field has been set.

### GetWarrantyNotes

`func (o *RmaTag) GetWarrantyNotes() string`

GetWarrantyNotes returns the WarrantyNotes field if non-nil, zero value otherwise.

### GetWarrantyNotesOk

`func (o *RmaTag) GetWarrantyNotesOk() (*string, bool)`

GetWarrantyNotesOk returns a tuple with the WarrantyNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyNotes

`func (o *RmaTag) SetWarrantyNotes(v string)`

SetWarrantyNotes sets WarrantyNotes field to given value.

### HasWarrantyNotes

`func (o *RmaTag) HasWarrantyNotes() bool`

HasWarrantyNotes returns a boolean if a field has been set.

### GetRepairCompany

`func (o *RmaTag) GetRepairCompany() CompanyReference`

GetRepairCompany returns the RepairCompany field if non-nil, zero value otherwise.

### GetRepairCompanyOk

`func (o *RmaTag) GetRepairCompanyOk() (*CompanyReference, bool)`

GetRepairCompanyOk returns a tuple with the RepairCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairCompany

`func (o *RmaTag) SetRepairCompany(v CompanyReference)`

SetRepairCompany sets RepairCompany field to given value.

### HasRepairCompany

`func (o *RmaTag) HasRepairCompany() bool`

HasRepairCompany returns a boolean if a field has been set.

### GetRepairContact

`func (o *RmaTag) GetRepairContact() ContactReference`

GetRepairContact returns the RepairContact field if non-nil, zero value otherwise.

### GetRepairContactOk

`func (o *RmaTag) GetRepairContactOk() (*ContactReference, bool)`

GetRepairContactOk returns a tuple with the RepairContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairContact

`func (o *RmaTag) SetRepairContact(v ContactReference)`

SetRepairContact sets RepairContact field to given value.

### HasRepairContact

`func (o *RmaTag) HasRepairContact() bool`

HasRepairContact returns a boolean if a field has been set.

### GetRepairContactType

`func (o *RmaTag) GetRepairContactType() string`

GetRepairContactType returns the RepairContactType field if non-nil, zero value otherwise.

### GetRepairContactTypeOk

`func (o *RmaTag) GetRepairContactTypeOk() (*string, bool)`

GetRepairContactTypeOk returns a tuple with the RepairContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairContactType

`func (o *RmaTag) SetRepairContactType(v string)`

SetRepairContactType sets RepairContactType field to given value.

### HasRepairContactType

`func (o *RmaTag) HasRepairContactType() bool`

HasRepairContactType returns a boolean if a field has been set.

### GetRepairContactPhone

`func (o *RmaTag) GetRepairContactPhone() string`

GetRepairContactPhone returns the RepairContactPhone field if non-nil, zero value otherwise.

### GetRepairContactPhoneOk

`func (o *RmaTag) GetRepairContactPhoneOk() (*string, bool)`

GetRepairContactPhoneOk returns a tuple with the RepairContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairContactPhone

`func (o *RmaTag) SetRepairContactPhone(v string)`

SetRepairContactPhone sets RepairContactPhone field to given value.

### HasRepairContactPhone

`func (o *RmaTag) HasRepairContactPhone() bool`

HasRepairContactPhone returns a boolean if a field has been set.

### GetRepairContactExtension

`func (o *RmaTag) GetRepairContactExtension() string`

GetRepairContactExtension returns the RepairContactExtension field if non-nil, zero value otherwise.

### GetRepairContactExtensionOk

`func (o *RmaTag) GetRepairContactExtensionOk() (*string, bool)`

GetRepairContactExtensionOk returns a tuple with the RepairContactExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairContactExtension

`func (o *RmaTag) SetRepairContactExtension(v string)`

SetRepairContactExtension sets RepairContactExtension field to given value.

### HasRepairContactExtension

`func (o *RmaTag) HasRepairContactExtension() bool`

HasRepairContactExtension returns a boolean if a field has been set.

### GetRepairContactEmail

`func (o *RmaTag) GetRepairContactEmail() string`

GetRepairContactEmail returns the RepairContactEmail field if non-nil, zero value otherwise.

### GetRepairContactEmailOk

`func (o *RmaTag) GetRepairContactEmailOk() (*string, bool)`

GetRepairContactEmailOk returns a tuple with the RepairContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairContactEmail

`func (o *RmaTag) SetRepairContactEmail(v string)`

SetRepairContactEmail sets RepairContactEmail field to given value.

### HasRepairContactEmail

`func (o *RmaTag) HasRepairContactEmail() bool`

HasRepairContactEmail returns a boolean if a field has been set.

### GetRepairContactAddressLine1

`func (o *RmaTag) GetRepairContactAddressLine1() string`

GetRepairContactAddressLine1 returns the RepairContactAddressLine1 field if non-nil, zero value otherwise.

### GetRepairContactAddressLine1Ok

`func (o *RmaTag) GetRepairContactAddressLine1Ok() (*string, bool)`

GetRepairContactAddressLine1Ok returns a tuple with the RepairContactAddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairContactAddressLine1

`func (o *RmaTag) SetRepairContactAddressLine1(v string)`

SetRepairContactAddressLine1 sets RepairContactAddressLine1 field to given value.

### HasRepairContactAddressLine1

`func (o *RmaTag) HasRepairContactAddressLine1() bool`

HasRepairContactAddressLine1 returns a boolean if a field has been set.

### GetRepairContactAddressLine2

`func (o *RmaTag) GetRepairContactAddressLine2() string`

GetRepairContactAddressLine2 returns the RepairContactAddressLine2 field if non-nil, zero value otherwise.

### GetRepairContactAddressLine2Ok

`func (o *RmaTag) GetRepairContactAddressLine2Ok() (*string, bool)`

GetRepairContactAddressLine2Ok returns a tuple with the RepairContactAddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairContactAddressLine2

`func (o *RmaTag) SetRepairContactAddressLine2(v string)`

SetRepairContactAddressLine2 sets RepairContactAddressLine2 field to given value.

### HasRepairContactAddressLine2

`func (o *RmaTag) HasRepairContactAddressLine2() bool`

HasRepairContactAddressLine2 returns a boolean if a field has been set.

### GetRepairContactCity

`func (o *RmaTag) GetRepairContactCity() string`

GetRepairContactCity returns the RepairContactCity field if non-nil, zero value otherwise.

### GetRepairContactCityOk

`func (o *RmaTag) GetRepairContactCityOk() (*string, bool)`

GetRepairContactCityOk returns a tuple with the RepairContactCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairContactCity

`func (o *RmaTag) SetRepairContactCity(v string)`

SetRepairContactCity sets RepairContactCity field to given value.

### HasRepairContactCity

`func (o *RmaTag) HasRepairContactCity() bool`

HasRepairContactCity returns a boolean if a field has been set.

### GetRepairContactState

`func (o *RmaTag) GetRepairContactState() string`

GetRepairContactState returns the RepairContactState field if non-nil, zero value otherwise.

### GetRepairContactStateOk

`func (o *RmaTag) GetRepairContactStateOk() (*string, bool)`

GetRepairContactStateOk returns a tuple with the RepairContactState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairContactState

`func (o *RmaTag) SetRepairContactState(v string)`

SetRepairContactState sets RepairContactState field to given value.

### HasRepairContactState

`func (o *RmaTag) HasRepairContactState() bool`

HasRepairContactState returns a boolean if a field has been set.

### GetRepairContactZip

`func (o *RmaTag) GetRepairContactZip() string`

GetRepairContactZip returns the RepairContactZip field if non-nil, zero value otherwise.

### GetRepairContactZipOk

`func (o *RmaTag) GetRepairContactZipOk() (*string, bool)`

GetRepairContactZipOk returns a tuple with the RepairContactZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairContactZip

`func (o *RmaTag) SetRepairContactZip(v string)`

SetRepairContactZip sets RepairContactZip field to given value.

### HasRepairContactZip

`func (o *RmaTag) HasRepairContactZip() bool`

HasRepairContactZip returns a boolean if a field has been set.

### GetRepairContactCountry

`func (o *RmaTag) GetRepairContactCountry() CountryReference`

GetRepairContactCountry returns the RepairContactCountry field if non-nil, zero value otherwise.

### GetRepairContactCountryOk

`func (o *RmaTag) GetRepairContactCountryOk() (*CountryReference, bool)`

GetRepairContactCountryOk returns a tuple with the RepairContactCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairContactCountry

`func (o *RmaTag) SetRepairContactCountry(v CountryReference)`

SetRepairContactCountry sets RepairContactCountry field to given value.

### HasRepairContactCountry

`func (o *RmaTag) HasRepairContactCountry() bool`

HasRepairContactCountry returns a boolean if a field has been set.

### GetRepairOrderNumber

`func (o *RmaTag) GetRepairOrderNumber() string`

GetRepairOrderNumber returns the RepairOrderNumber field if non-nil, zero value otherwise.

### GetRepairOrderNumberOk

`func (o *RmaTag) GetRepairOrderNumberOk() (*string, bool)`

GetRepairOrderNumberOk returns a tuple with the RepairOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairOrderNumber

`func (o *RmaTag) SetRepairOrderNumber(v string)`

SetRepairOrderNumber sets RepairOrderNumber field to given value.

### HasRepairOrderNumber

`func (o *RmaTag) HasRepairOrderNumber() bool`

HasRepairOrderNumber returns a boolean if a field has been set.

### GetRepairSite

`func (o *RmaTag) GetRepairSite() SiteReference`

GetRepairSite returns the RepairSite field if non-nil, zero value otherwise.

### GetRepairSiteOk

`func (o *RmaTag) GetRepairSiteOk() (*SiteReference, bool)`

GetRepairSiteOk returns a tuple with the RepairSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairSite

`func (o *RmaTag) SetRepairSite(v SiteReference)`

SetRepairSite sets RepairSite field to given value.

### HasRepairSite

`func (o *RmaTag) HasRepairSite() bool`

HasRepairSite returns a boolean if a field has been set.

### GetRepairNotes

`func (o *RmaTag) GetRepairNotes() string`

GetRepairNotes returns the RepairNotes field if non-nil, zero value otherwise.

### GetRepairNotesOk

`func (o *RmaTag) GetRepairNotesOk() (*string, bool)`

GetRepairNotesOk returns a tuple with the RepairNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairNotes

`func (o *RmaTag) SetRepairNotes(v string)`

SetRepairNotes sets RepairNotes field to given value.

### HasRepairNotes

`func (o *RmaTag) HasRepairNotes() bool`

HasRepairNotes returns a boolean if a field has been set.

### GetDropShipFlag

`func (o *RmaTag) GetDropShipFlag() bool`

GetDropShipFlag returns the DropShipFlag field if non-nil, zero value otherwise.

### GetDropShipFlagOk

`func (o *RmaTag) GetDropShipFlagOk() (*bool, bool)`

GetDropShipFlagOk returns a tuple with the DropShipFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropShipFlag

`func (o *RmaTag) SetDropShipFlag(v bool)`

SetDropShipFlag sets DropShipFlag field to given value.

### HasDropShipFlag

`func (o *RmaTag) HasDropShipFlag() bool`

HasDropShipFlag returns a boolean if a field has been set.

### SetDropShipFlagNil

`func (o *RmaTag) SetDropShipFlagNil(b bool)`

 SetDropShipFlagNil sets the value for DropShipFlag to be an explicit nil

### UnsetDropShipFlag
`func (o *RmaTag) UnsetDropShipFlag()`

UnsetDropShipFlag ensures that no value is present for DropShipFlag, not even an explicit nil
### GetShipMethod

`func (o *RmaTag) GetShipMethod() ShipmentMethodReference`

GetShipMethod returns the ShipMethod field if non-nil, zero value otherwise.

### GetShipMethodOk

`func (o *RmaTag) GetShipMethodOk() (*ShipmentMethodReference, bool)`

GetShipMethodOk returns a tuple with the ShipMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipMethod

`func (o *RmaTag) SetShipMethod(v ShipmentMethodReference)`

SetShipMethod sets ShipMethod field to given value.

### HasShipMethod

`func (o *RmaTag) HasShipMethod() bool`

HasShipMethod returns a boolean if a field has been set.

### GetShippingDate

`func (o *RmaTag) GetShippingDate() string`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *RmaTag) GetShippingDateOk() (*string, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *RmaTag) SetShippingDate(v string)`

SetShippingDate sets ShippingDate field to given value.

### HasShippingDate

`func (o *RmaTag) HasShippingDate() bool`

HasShippingDate returns a boolean if a field has been set.

### GetShippingTrackingNumber

`func (o *RmaTag) GetShippingTrackingNumber() string`

GetShippingTrackingNumber returns the ShippingTrackingNumber field if non-nil, zero value otherwise.

### GetShippingTrackingNumberOk

`func (o *RmaTag) GetShippingTrackingNumberOk() (*string, bool)`

GetShippingTrackingNumberOk returns a tuple with the ShippingTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTrackingNumber

`func (o *RmaTag) SetShippingTrackingNumber(v string)`

SetShippingTrackingNumber sets ShippingTrackingNumber field to given value.

### HasShippingTrackingNumber

`func (o *RmaTag) HasShippingTrackingNumber() bool`

HasShippingTrackingNumber returns a boolean if a field has been set.

### GetInternalNotes

`func (o *RmaTag) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *RmaTag) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *RmaTag) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *RmaTag) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetClosingNotes

`func (o *RmaTag) GetClosingNotes() string`

GetClosingNotes returns the ClosingNotes field if non-nil, zero value otherwise.

### GetClosingNotesOk

`func (o *RmaTag) GetClosingNotesOk() (*string, bool)`

GetClosingNotesOk returns a tuple with the ClosingNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingNotes

`func (o *RmaTag) SetClosingNotes(v string)`

SetClosingNotes sets ClosingNotes field to given value.

### HasClosingNotes

`func (o *RmaTag) HasClosingNotes() bool`

HasClosingNotes returns a boolean if a field has been set.

### GetDateClosed

`func (o *RmaTag) GetDateClosed() string`

GetDateClosed returns the DateClosed field if non-nil, zero value otherwise.

### GetDateClosedOk

`func (o *RmaTag) GetDateClosedOk() (*string, bool)`

GetDateClosedOk returns a tuple with the DateClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClosed

`func (o *RmaTag) SetDateClosed(v string)`

SetDateClosed sets DateClosed field to given value.

### HasDateClosed

`func (o *RmaTag) HasDateClosed() bool`

HasDateClosed returns a boolean if a field has been set.

### GetAccountManager

`func (o *RmaTag) GetAccountManager() MemberReference`

GetAccountManager returns the AccountManager field if non-nil, zero value otherwise.

### GetAccountManagerOk

`func (o *RmaTag) GetAccountManagerOk() (*MemberReference, bool)`

GetAccountManagerOk returns a tuple with the AccountManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountManager

`func (o *RmaTag) SetAccountManager(v MemberReference)`

SetAccountManager sets AccountManager field to given value.

### HasAccountManager

`func (o *RmaTag) HasAccountManager() bool`

HasAccountManager returns a boolean if a field has been set.

### GetTechnicalContact

`func (o *RmaTag) GetTechnicalContact() MemberReference`

GetTechnicalContact returns the TechnicalContact field if non-nil, zero value otherwise.

### GetTechnicalContactOk

`func (o *RmaTag) GetTechnicalContactOk() (*MemberReference, bool)`

GetTechnicalContactOk returns a tuple with the TechnicalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContact

`func (o *RmaTag) SetTechnicalContact(v MemberReference)`

SetTechnicalContact sets TechnicalContact field to given value.

### HasTechnicalContact

`func (o *RmaTag) HasTechnicalContact() bool`

HasTechnicalContact returns a boolean if a field has been set.

### GetCurrency

`func (o *RmaTag) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RmaTag) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RmaTag) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *RmaTag) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetClosedBy

`func (o *RmaTag) GetClosedBy() MemberReference`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *RmaTag) GetClosedByOk() (*MemberReference, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *RmaTag) SetClosedBy(v MemberReference)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *RmaTag) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetInfo

`func (o *RmaTag) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *RmaTag) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *RmaTag) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *RmaTag) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *RmaTag) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RmaTag) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RmaTag) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RmaTag) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


