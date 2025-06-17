# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**PhoneExt** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**Status** | [**OrderStatusReference**](OrderStatusReference.md) |  | 
**Opportunity** | Pointer to [**OpportunityReference**](OpportunityReference.md) |  | [optional] 
**OrderDate** | Pointer to **time.Time** |  | [optional] 
**DueDate** | Pointer to **time.Time** |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**PoNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**SalesRep** | [**MemberReference**](MemberReference.md) |  | 
**Notes** | Pointer to **string** |  | [optional] 
**BillClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**BillShippedFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictDownpaymentFlag** | Pointer to **NullableBool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TopCommentFlag** | Pointer to **NullableBool** |  | [optional] 
**BottomCommentFlag** | Pointer to **NullableBool** |  | [optional] 
**ShipToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ShipToContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**ShipToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**BillToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**BillToContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**BillToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**ProductIds** | Pointer to **[]int32** |  | [optional] 
**DocumentIds** | Pointer to **[]int32** |  | [optional] 
**InvoiceIds** | Pointer to **[]int32** |  | [optional] 
**ConfigIds** | Pointer to **[]int32** |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**TaxTotal** | Pointer to **NullableFloat64** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**CompanyLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**SubTotal** | Pointer to **float64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewOrder

`func NewOrder(company CompanyReference, status OrderStatusReference, salesRep MemberReference, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Order) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompany

`func (o *Order) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Order) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Order) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetContact

`func (o *Order) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Order) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Order) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Order) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetPhone

`func (o *Order) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Order) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Order) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Order) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPhoneExt

`func (o *Order) GetPhoneExt() string`

GetPhoneExt returns the PhoneExt field if non-nil, zero value otherwise.

### GetPhoneExtOk

`func (o *Order) GetPhoneExtOk() (*string, bool)`

GetPhoneExtOk returns a tuple with the PhoneExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneExt

`func (o *Order) SetPhoneExt(v string)`

SetPhoneExt sets PhoneExt field to given value.

### HasPhoneExt

`func (o *Order) HasPhoneExt() bool`

HasPhoneExt returns a boolean if a field has been set.

### GetEmail

`func (o *Order) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Order) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Order) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Order) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSite

`func (o *Order) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Order) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Order) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *Order) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetStatus

`func (o *Order) GetStatus() OrderStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*OrderStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v OrderStatusReference)`

SetStatus sets Status field to given value.


### GetOpportunity

`func (o *Order) GetOpportunity() OpportunityReference`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *Order) GetOpportunityOk() (*OpportunityReference, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *Order) SetOpportunity(v OpportunityReference)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *Order) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### GetOrderDate

`func (o *Order) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *Order) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *Order) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.

### HasOrderDate

`func (o *Order) HasOrderDate() bool`

HasOrderDate returns a boolean if a field has been set.

### GetDueDate

`func (o *Order) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Order) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Order) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Order) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetBillingTerms

`func (o *Order) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *Order) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *Order) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *Order) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetTaxCode

`func (o *Order) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *Order) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *Order) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *Order) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetPoNumber

`func (o *Order) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *Order) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *Order) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *Order) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetLocation

`func (o *Order) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Order) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Order) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Order) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *Order) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Order) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Order) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Order) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetSalesRep

`func (o *Order) GetSalesRep() MemberReference`

GetSalesRep returns the SalesRep field if non-nil, zero value otherwise.

### GetSalesRepOk

`func (o *Order) GetSalesRepOk() (*MemberReference, bool)`

GetSalesRepOk returns a tuple with the SalesRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRep

`func (o *Order) SetSalesRep(v MemberReference)`

SetSalesRep sets SalesRep field to given value.


### GetNotes

`func (o *Order) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Order) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Order) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Order) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetBillClosedFlag

`func (o *Order) GetBillClosedFlag() bool`

GetBillClosedFlag returns the BillClosedFlag field if non-nil, zero value otherwise.

### GetBillClosedFlagOk

`func (o *Order) GetBillClosedFlagOk() (*bool, bool)`

GetBillClosedFlagOk returns a tuple with the BillClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillClosedFlag

`func (o *Order) SetBillClosedFlag(v bool)`

SetBillClosedFlag sets BillClosedFlag field to given value.

### HasBillClosedFlag

`func (o *Order) HasBillClosedFlag() bool`

HasBillClosedFlag returns a boolean if a field has been set.

### SetBillClosedFlagNil

`func (o *Order) SetBillClosedFlagNil(b bool)`

 SetBillClosedFlagNil sets the value for BillClosedFlag to be an explicit nil

### UnsetBillClosedFlag
`func (o *Order) UnsetBillClosedFlag()`

UnsetBillClosedFlag ensures that no value is present for BillClosedFlag, not even an explicit nil
### GetBillShippedFlag

`func (o *Order) GetBillShippedFlag() bool`

GetBillShippedFlag returns the BillShippedFlag field if non-nil, zero value otherwise.

### GetBillShippedFlagOk

`func (o *Order) GetBillShippedFlagOk() (*bool, bool)`

GetBillShippedFlagOk returns a tuple with the BillShippedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillShippedFlag

`func (o *Order) SetBillShippedFlag(v bool)`

SetBillShippedFlag sets BillShippedFlag field to given value.

### HasBillShippedFlag

`func (o *Order) HasBillShippedFlag() bool`

HasBillShippedFlag returns a boolean if a field has been set.

### SetBillShippedFlagNil

`func (o *Order) SetBillShippedFlagNil(b bool)`

 SetBillShippedFlagNil sets the value for BillShippedFlag to be an explicit nil

### UnsetBillShippedFlag
`func (o *Order) UnsetBillShippedFlag()`

UnsetBillShippedFlag ensures that no value is present for BillShippedFlag, not even an explicit nil
### GetRestrictDownpaymentFlag

`func (o *Order) GetRestrictDownpaymentFlag() bool`

GetRestrictDownpaymentFlag returns the RestrictDownpaymentFlag field if non-nil, zero value otherwise.

### GetRestrictDownpaymentFlagOk

`func (o *Order) GetRestrictDownpaymentFlagOk() (*bool, bool)`

GetRestrictDownpaymentFlagOk returns a tuple with the RestrictDownpaymentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDownpaymentFlag

`func (o *Order) SetRestrictDownpaymentFlag(v bool)`

SetRestrictDownpaymentFlag sets RestrictDownpaymentFlag field to given value.

### HasRestrictDownpaymentFlag

`func (o *Order) HasRestrictDownpaymentFlag() bool`

HasRestrictDownpaymentFlag returns a boolean if a field has been set.

### SetRestrictDownpaymentFlagNil

`func (o *Order) SetRestrictDownpaymentFlagNil(b bool)`

 SetRestrictDownpaymentFlagNil sets the value for RestrictDownpaymentFlag to be an explicit nil

### UnsetRestrictDownpaymentFlag
`func (o *Order) UnsetRestrictDownpaymentFlag()`

UnsetRestrictDownpaymentFlag ensures that no value is present for RestrictDownpaymentFlag, not even an explicit nil
### GetDescription

`func (o *Order) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Order) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Order) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Order) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTopCommentFlag

`func (o *Order) GetTopCommentFlag() bool`

GetTopCommentFlag returns the TopCommentFlag field if non-nil, zero value otherwise.

### GetTopCommentFlagOk

`func (o *Order) GetTopCommentFlagOk() (*bool, bool)`

GetTopCommentFlagOk returns a tuple with the TopCommentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopCommentFlag

`func (o *Order) SetTopCommentFlag(v bool)`

SetTopCommentFlag sets TopCommentFlag field to given value.

### HasTopCommentFlag

`func (o *Order) HasTopCommentFlag() bool`

HasTopCommentFlag returns a boolean if a field has been set.

### SetTopCommentFlagNil

`func (o *Order) SetTopCommentFlagNil(b bool)`

 SetTopCommentFlagNil sets the value for TopCommentFlag to be an explicit nil

### UnsetTopCommentFlag
`func (o *Order) UnsetTopCommentFlag()`

UnsetTopCommentFlag ensures that no value is present for TopCommentFlag, not even an explicit nil
### GetBottomCommentFlag

`func (o *Order) GetBottomCommentFlag() bool`

GetBottomCommentFlag returns the BottomCommentFlag field if non-nil, zero value otherwise.

### GetBottomCommentFlagOk

`func (o *Order) GetBottomCommentFlagOk() (*bool, bool)`

GetBottomCommentFlagOk returns a tuple with the BottomCommentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomCommentFlag

`func (o *Order) SetBottomCommentFlag(v bool)`

SetBottomCommentFlag sets BottomCommentFlag field to given value.

### HasBottomCommentFlag

`func (o *Order) HasBottomCommentFlag() bool`

HasBottomCommentFlag returns a boolean if a field has been set.

### SetBottomCommentFlagNil

`func (o *Order) SetBottomCommentFlagNil(b bool)`

 SetBottomCommentFlagNil sets the value for BottomCommentFlag to be an explicit nil

### UnsetBottomCommentFlag
`func (o *Order) UnsetBottomCommentFlag()`

UnsetBottomCommentFlag ensures that no value is present for BottomCommentFlag, not even an explicit nil
### GetShipToCompany

`func (o *Order) GetShipToCompany() CompanyReference`

GetShipToCompany returns the ShipToCompany field if non-nil, zero value otherwise.

### GetShipToCompanyOk

`func (o *Order) GetShipToCompanyOk() (*CompanyReference, bool)`

GetShipToCompanyOk returns a tuple with the ShipToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompany

`func (o *Order) SetShipToCompany(v CompanyReference)`

SetShipToCompany sets ShipToCompany field to given value.

### HasShipToCompany

`func (o *Order) HasShipToCompany() bool`

HasShipToCompany returns a boolean if a field has been set.

### GetShipToContact

`func (o *Order) GetShipToContact() ContactReference`

GetShipToContact returns the ShipToContact field if non-nil, zero value otherwise.

### GetShipToContactOk

`func (o *Order) GetShipToContactOk() (*ContactReference, bool)`

GetShipToContactOk returns a tuple with the ShipToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToContact

`func (o *Order) SetShipToContact(v ContactReference)`

SetShipToContact sets ShipToContact field to given value.

### HasShipToContact

`func (o *Order) HasShipToContact() bool`

HasShipToContact returns a boolean if a field has been set.

### GetShipToSite

`func (o *Order) GetShipToSite() SiteReference`

GetShipToSite returns the ShipToSite field if non-nil, zero value otherwise.

### GetShipToSiteOk

`func (o *Order) GetShipToSiteOk() (*SiteReference, bool)`

GetShipToSiteOk returns a tuple with the ShipToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToSite

`func (o *Order) SetShipToSite(v SiteReference)`

SetShipToSite sets ShipToSite field to given value.

### HasShipToSite

`func (o *Order) HasShipToSite() bool`

HasShipToSite returns a boolean if a field has been set.

### GetBillToCompany

`func (o *Order) GetBillToCompany() CompanyReference`

GetBillToCompany returns the BillToCompany field if non-nil, zero value otherwise.

### GetBillToCompanyOk

`func (o *Order) GetBillToCompanyOk() (*CompanyReference, bool)`

GetBillToCompanyOk returns a tuple with the BillToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToCompany

`func (o *Order) SetBillToCompany(v CompanyReference)`

SetBillToCompany sets BillToCompany field to given value.

### HasBillToCompany

`func (o *Order) HasBillToCompany() bool`

HasBillToCompany returns a boolean if a field has been set.

### GetBillToContact

`func (o *Order) GetBillToContact() ContactReference`

GetBillToContact returns the BillToContact field if non-nil, zero value otherwise.

### GetBillToContactOk

`func (o *Order) GetBillToContactOk() (*ContactReference, bool)`

GetBillToContactOk returns a tuple with the BillToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToContact

`func (o *Order) SetBillToContact(v ContactReference)`

SetBillToContact sets BillToContact field to given value.

### HasBillToContact

`func (o *Order) HasBillToContact() bool`

HasBillToContact returns a boolean if a field has been set.

### GetBillToSite

`func (o *Order) GetBillToSite() SiteReference`

GetBillToSite returns the BillToSite field if non-nil, zero value otherwise.

### GetBillToSiteOk

`func (o *Order) GetBillToSiteOk() (*SiteReference, bool)`

GetBillToSiteOk returns a tuple with the BillToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToSite

`func (o *Order) SetBillToSite(v SiteReference)`

SetBillToSite sets BillToSite field to given value.

### HasBillToSite

`func (o *Order) HasBillToSite() bool`

HasBillToSite returns a boolean if a field has been set.

### GetProductIds

`func (o *Order) GetProductIds() []int32`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *Order) GetProductIdsOk() (*[]int32, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *Order) SetProductIds(v []int32)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *Order) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetDocumentIds

`func (o *Order) GetDocumentIds() []int32`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *Order) GetDocumentIdsOk() (*[]int32, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *Order) SetDocumentIds(v []int32)`

SetDocumentIds sets DocumentIds field to given value.

### HasDocumentIds

`func (o *Order) HasDocumentIds() bool`

HasDocumentIds returns a boolean if a field has been set.

### GetInvoiceIds

`func (o *Order) GetInvoiceIds() []int32`

GetInvoiceIds returns the InvoiceIds field if non-nil, zero value otherwise.

### GetInvoiceIdsOk

`func (o *Order) GetInvoiceIdsOk() (*[]int32, bool)`

GetInvoiceIdsOk returns a tuple with the InvoiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceIds

`func (o *Order) SetInvoiceIds(v []int32)`

SetInvoiceIds sets InvoiceIds field to given value.

### HasInvoiceIds

`func (o *Order) HasInvoiceIds() bool`

HasInvoiceIds returns a boolean if a field has been set.

### GetConfigIds

`func (o *Order) GetConfigIds() []int32`

GetConfigIds returns the ConfigIds field if non-nil, zero value otherwise.

### GetConfigIdsOk

`func (o *Order) GetConfigIdsOk() (*[]int32, bool)`

GetConfigIdsOk returns a tuple with the ConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIds

`func (o *Order) SetConfigIds(v []int32)`

SetConfigIds sets ConfigIds field to given value.

### HasConfigIds

`func (o *Order) HasConfigIds() bool`

HasConfigIds returns a boolean if a field has been set.

### GetTotal

`func (o *Order) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Order) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Order) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Order) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *Order) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *Order) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetTaxTotal

`func (o *Order) GetTaxTotal() float64`

GetTaxTotal returns the TaxTotal field if non-nil, zero value otherwise.

### GetTaxTotalOk

`func (o *Order) GetTaxTotalOk() (*float64, bool)`

GetTaxTotalOk returns a tuple with the TaxTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxTotal

`func (o *Order) SetTaxTotal(v float64)`

SetTaxTotal sets TaxTotal field to given value.

### HasTaxTotal

`func (o *Order) HasTaxTotal() bool`

HasTaxTotal returns a boolean if a field has been set.

### SetTaxTotalNil

`func (o *Order) SetTaxTotalNil(b bool)`

 SetTaxTotalNil sets the value for TaxTotal to be an explicit nil

### UnsetTaxTotal
`func (o *Order) UnsetTaxTotal()`

UnsetTaxTotal ensures that no value is present for TaxTotal, not even an explicit nil
### GetCurrency

`func (o *Order) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Order) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Order) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Order) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCompanyLocation

`func (o *Order) GetCompanyLocation() SystemLocationReference`

GetCompanyLocation returns the CompanyLocation field if non-nil, zero value otherwise.

### GetCompanyLocationOk

`func (o *Order) GetCompanyLocationOk() (*SystemLocationReference, bool)`

GetCompanyLocationOk returns a tuple with the CompanyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLocation

`func (o *Order) SetCompanyLocation(v SystemLocationReference)`

SetCompanyLocation sets CompanyLocation field to given value.

### HasCompanyLocation

`func (o *Order) HasCompanyLocation() bool`

HasCompanyLocation returns a boolean if a field has been set.

### GetSubTotal

`func (o *Order) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *Order) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *Order) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *Order) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### GetInfo

`func (o *Order) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Order) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Order) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Order) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *Order) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Order) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Order) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Order) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


