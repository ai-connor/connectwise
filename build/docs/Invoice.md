# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**InvoiceNumber** | Pointer to **string** |  Max length: 15; Required On Updates; | [optional] 
**Type** | **NullableString** |  | 
**Status** | Pointer to [**BillingStatusReference**](BillingStatusReference.md) |  | [optional] 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
**BillToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ShipToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**ApplyToType** | Pointer to **NullableString** |  | [optional] 
**ApplyToId** | Pointer to **NullableInt32** |  | [optional] 
**Attention** | Pointer to **string** |  Max length: 60; | [optional] 
**ShipToAttention** | Pointer to **string** |  Max length: 60; | [optional] 
**BillingSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**BillingSiteAddressLine1** | Pointer to **string** |  | [optional] 
**BillingSiteAddressLine2** | Pointer to **string** |  | [optional] 
**BillingSiteCity** | Pointer to **string** |  | [optional] 
**BillingSiteState** | Pointer to **string** |  | [optional] 
**BillingSiteZip** | Pointer to **string** |  | [optional] 
**BillingSiteCountry** | Pointer to **string** |  | [optional] 
**ShippingSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**ShippingSiteAddressLine1** | Pointer to **string** |  | [optional] 
**ShippingSiteAddressLine2** | Pointer to **string** |  | [optional] 
**ShippingSiteCity** | Pointer to **string** |  | [optional] 
**ShippingSiteState** | Pointer to **string** |  | [optional] 
**ShippingSiteZip** | Pointer to **string** |  | [optional] 
**ShippingSiteCountry** | Pointer to **string** |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**Reference** | Pointer to **string** |  Max length: 50; | [optional] 
**CustomerPO** | Pointer to **string** |  Max length: 50; | [optional] 
**TemplateSetupId** | Pointer to **NullableInt32** | Can be obtained via InvoiceTemplate report. | [optional] 
**InvoiceTemplate** | Pointer to [**InvoiceTemplateDetailReference**](InvoiceTemplateDetailReference.md) |  | [optional] 
**EmailTemplateId** | Pointer to **NullableInt32** | Can be obtained via InvoiceEmailTemplate report. | [optional] 
**AddToBatchEmailList** | Pointer to **NullableBool** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**RestrictDownpaymentFlag** | Pointer to **NullableBool** |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  Required On Updates; | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**DepartmentId** | Pointer to **NullableInt32** | departmentId is only required for special invoices. | [optional] 
**Department** | Pointer to [**BillingUnitReference**](BillingUnitReference.md) |  | [optional] 
**TerritoryId** | Pointer to **NullableInt32** |  | [optional] 
**Territory** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**TopComment** | Pointer to **string** |  | [optional] 
**BottomComment** | Pointer to **string** |  | [optional] 
**TaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**InternalNotes** | Pointer to **string** |  | [optional] 
**DownpaymentPreviouslyTaxedFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceTotal** | Pointer to **NullableFloat64** |  | [optional] 
**OverrideDownPaymentAmountFlag** | Pointer to **NullableBool** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**DueDate** | Pointer to **time.Time** |  | [optional] 
**ExpenseTotal** | Pointer to **NullableFloat64** |  | [optional] 
**ProductTotal** | Pointer to **NullableFloat64** |  | [optional] 
**PreviousProgressApplied** | Pointer to **NullableFloat64** |  | [optional] 
**ServiceAdjustmentAmount** | Pointer to **NullableFloat64** |  | [optional] 
**AgreementAmount** | Pointer to **NullableFloat64** |  | [optional] 
**DownpaymentApplied** | Pointer to **NullableFloat64** |  | [optional] 
**Subtotal** | Pointer to **NullableFloat64** |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**RemainingDownpayment** | Pointer to **NullableFloat64** |  | [optional] 
**SalesTax** | Pointer to **NullableFloat64** |  | [optional] 
**AdjustmentReason** | Pointer to **string** |  | [optional] 
**AdjustedBy** | Pointer to **string** |  | [optional] 
**ClosedBy** | Pointer to **string** |  | [optional] 
**Payments** | Pointer to **NullableFloat64** |  | [optional] 
**Credits** | Pointer to **NullableFloat64** |  | [optional] 
**Balance** | Pointer to **NullableFloat64** |  | [optional] 
**SpecialInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**BillingSetupReference** | Pointer to [**BillingSetupReference**](BillingSetupReference.md) |  | [optional] 
**Ticket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**Phase** | Pointer to [**ProjectPhaseReference**](ProjectPhaseReference.md) |  | [optional] 
**SalesOrder** | Pointer to [**SalesOrderReference**](SalesOrderReference.md) |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**GlBatch** | Pointer to [**BatchReference**](BatchReference.md) |  | [optional] 
**UnbatchedBatch** | Pointer to [**BatchReference**](BatchReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewInvoice

`func NewInvoice(type_ NullableString, company CompanyReference, ) *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *Invoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *Invoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *Invoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *Invoice) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetType

`func (o *Invoice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Invoice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Invoice) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *Invoice) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Invoice) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStatus

`func (o *Invoice) GetStatus() BillingStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*BillingStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v BillingStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Invoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCompany

`func (o *Invoice) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Invoice) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Invoice) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetBillToCompany

`func (o *Invoice) GetBillToCompany() CompanyReference`

GetBillToCompany returns the BillToCompany field if non-nil, zero value otherwise.

### GetBillToCompanyOk

`func (o *Invoice) GetBillToCompanyOk() (*CompanyReference, bool)`

GetBillToCompanyOk returns a tuple with the BillToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToCompany

`func (o *Invoice) SetBillToCompany(v CompanyReference)`

SetBillToCompany sets BillToCompany field to given value.

### HasBillToCompany

`func (o *Invoice) HasBillToCompany() bool`

HasBillToCompany returns a boolean if a field has been set.

### GetShipToCompany

`func (o *Invoice) GetShipToCompany() CompanyReference`

GetShipToCompany returns the ShipToCompany field if non-nil, zero value otherwise.

### GetShipToCompanyOk

`func (o *Invoice) GetShipToCompanyOk() (*CompanyReference, bool)`

GetShipToCompanyOk returns a tuple with the ShipToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompany

`func (o *Invoice) SetShipToCompany(v CompanyReference)`

SetShipToCompany sets ShipToCompany field to given value.

### HasShipToCompany

`func (o *Invoice) HasShipToCompany() bool`

HasShipToCompany returns a boolean if a field has been set.

### GetAccountNumber

`func (o *Invoice) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Invoice) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Invoice) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Invoice) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetApplyToType

`func (o *Invoice) GetApplyToType() string`

GetApplyToType returns the ApplyToType field if non-nil, zero value otherwise.

### GetApplyToTypeOk

`func (o *Invoice) GetApplyToTypeOk() (*string, bool)`

GetApplyToTypeOk returns a tuple with the ApplyToType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToType

`func (o *Invoice) SetApplyToType(v string)`

SetApplyToType sets ApplyToType field to given value.

### HasApplyToType

`func (o *Invoice) HasApplyToType() bool`

HasApplyToType returns a boolean if a field has been set.

### SetApplyToTypeNil

`func (o *Invoice) SetApplyToTypeNil(b bool)`

 SetApplyToTypeNil sets the value for ApplyToType to be an explicit nil

### UnsetApplyToType
`func (o *Invoice) UnsetApplyToType()`

UnsetApplyToType ensures that no value is present for ApplyToType, not even an explicit nil
### GetApplyToId

`func (o *Invoice) GetApplyToId() int32`

GetApplyToId returns the ApplyToId field if non-nil, zero value otherwise.

### GetApplyToIdOk

`func (o *Invoice) GetApplyToIdOk() (*int32, bool)`

GetApplyToIdOk returns a tuple with the ApplyToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToId

`func (o *Invoice) SetApplyToId(v int32)`

SetApplyToId sets ApplyToId field to given value.

### HasApplyToId

`func (o *Invoice) HasApplyToId() bool`

HasApplyToId returns a boolean if a field has been set.

### SetApplyToIdNil

`func (o *Invoice) SetApplyToIdNil(b bool)`

 SetApplyToIdNil sets the value for ApplyToId to be an explicit nil

### UnsetApplyToId
`func (o *Invoice) UnsetApplyToId()`

UnsetApplyToId ensures that no value is present for ApplyToId, not even an explicit nil
### GetAttention

`func (o *Invoice) GetAttention() string`

GetAttention returns the Attention field if non-nil, zero value otherwise.

### GetAttentionOk

`func (o *Invoice) GetAttentionOk() (*string, bool)`

GetAttentionOk returns a tuple with the Attention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttention

`func (o *Invoice) SetAttention(v string)`

SetAttention sets Attention field to given value.

### HasAttention

`func (o *Invoice) HasAttention() bool`

HasAttention returns a boolean if a field has been set.

### GetShipToAttention

`func (o *Invoice) GetShipToAttention() string`

GetShipToAttention returns the ShipToAttention field if non-nil, zero value otherwise.

### GetShipToAttentionOk

`func (o *Invoice) GetShipToAttentionOk() (*string, bool)`

GetShipToAttentionOk returns a tuple with the ShipToAttention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToAttention

`func (o *Invoice) SetShipToAttention(v string)`

SetShipToAttention sets ShipToAttention field to given value.

### HasShipToAttention

`func (o *Invoice) HasShipToAttention() bool`

HasShipToAttention returns a boolean if a field has been set.

### GetBillingSite

`func (o *Invoice) GetBillingSite() SiteReference`

GetBillingSite returns the BillingSite field if non-nil, zero value otherwise.

### GetBillingSiteOk

`func (o *Invoice) GetBillingSiteOk() (*SiteReference, bool)`

GetBillingSiteOk returns a tuple with the BillingSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSite

`func (o *Invoice) SetBillingSite(v SiteReference)`

SetBillingSite sets BillingSite field to given value.

### HasBillingSite

`func (o *Invoice) HasBillingSite() bool`

HasBillingSite returns a boolean if a field has been set.

### GetBillingSiteAddressLine1

`func (o *Invoice) GetBillingSiteAddressLine1() string`

GetBillingSiteAddressLine1 returns the BillingSiteAddressLine1 field if non-nil, zero value otherwise.

### GetBillingSiteAddressLine1Ok

`func (o *Invoice) GetBillingSiteAddressLine1Ok() (*string, bool)`

GetBillingSiteAddressLine1Ok returns a tuple with the BillingSiteAddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSiteAddressLine1

`func (o *Invoice) SetBillingSiteAddressLine1(v string)`

SetBillingSiteAddressLine1 sets BillingSiteAddressLine1 field to given value.

### HasBillingSiteAddressLine1

`func (o *Invoice) HasBillingSiteAddressLine1() bool`

HasBillingSiteAddressLine1 returns a boolean if a field has been set.

### GetBillingSiteAddressLine2

`func (o *Invoice) GetBillingSiteAddressLine2() string`

GetBillingSiteAddressLine2 returns the BillingSiteAddressLine2 field if non-nil, zero value otherwise.

### GetBillingSiteAddressLine2Ok

`func (o *Invoice) GetBillingSiteAddressLine2Ok() (*string, bool)`

GetBillingSiteAddressLine2Ok returns a tuple with the BillingSiteAddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSiteAddressLine2

`func (o *Invoice) SetBillingSiteAddressLine2(v string)`

SetBillingSiteAddressLine2 sets BillingSiteAddressLine2 field to given value.

### HasBillingSiteAddressLine2

`func (o *Invoice) HasBillingSiteAddressLine2() bool`

HasBillingSiteAddressLine2 returns a boolean if a field has been set.

### GetBillingSiteCity

`func (o *Invoice) GetBillingSiteCity() string`

GetBillingSiteCity returns the BillingSiteCity field if non-nil, zero value otherwise.

### GetBillingSiteCityOk

`func (o *Invoice) GetBillingSiteCityOk() (*string, bool)`

GetBillingSiteCityOk returns a tuple with the BillingSiteCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSiteCity

`func (o *Invoice) SetBillingSiteCity(v string)`

SetBillingSiteCity sets BillingSiteCity field to given value.

### HasBillingSiteCity

`func (o *Invoice) HasBillingSiteCity() bool`

HasBillingSiteCity returns a boolean if a field has been set.

### GetBillingSiteState

`func (o *Invoice) GetBillingSiteState() string`

GetBillingSiteState returns the BillingSiteState field if non-nil, zero value otherwise.

### GetBillingSiteStateOk

`func (o *Invoice) GetBillingSiteStateOk() (*string, bool)`

GetBillingSiteStateOk returns a tuple with the BillingSiteState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSiteState

`func (o *Invoice) SetBillingSiteState(v string)`

SetBillingSiteState sets BillingSiteState field to given value.

### HasBillingSiteState

`func (o *Invoice) HasBillingSiteState() bool`

HasBillingSiteState returns a boolean if a field has been set.

### GetBillingSiteZip

`func (o *Invoice) GetBillingSiteZip() string`

GetBillingSiteZip returns the BillingSiteZip field if non-nil, zero value otherwise.

### GetBillingSiteZipOk

`func (o *Invoice) GetBillingSiteZipOk() (*string, bool)`

GetBillingSiteZipOk returns a tuple with the BillingSiteZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSiteZip

`func (o *Invoice) SetBillingSiteZip(v string)`

SetBillingSiteZip sets BillingSiteZip field to given value.

### HasBillingSiteZip

`func (o *Invoice) HasBillingSiteZip() bool`

HasBillingSiteZip returns a boolean if a field has been set.

### GetBillingSiteCountry

`func (o *Invoice) GetBillingSiteCountry() string`

GetBillingSiteCountry returns the BillingSiteCountry field if non-nil, zero value otherwise.

### GetBillingSiteCountryOk

`func (o *Invoice) GetBillingSiteCountryOk() (*string, bool)`

GetBillingSiteCountryOk returns a tuple with the BillingSiteCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSiteCountry

`func (o *Invoice) SetBillingSiteCountry(v string)`

SetBillingSiteCountry sets BillingSiteCountry field to given value.

### HasBillingSiteCountry

`func (o *Invoice) HasBillingSiteCountry() bool`

HasBillingSiteCountry returns a boolean if a field has been set.

### GetShippingSite

`func (o *Invoice) GetShippingSite() SiteReference`

GetShippingSite returns the ShippingSite field if non-nil, zero value otherwise.

### GetShippingSiteOk

`func (o *Invoice) GetShippingSiteOk() (*SiteReference, bool)`

GetShippingSiteOk returns a tuple with the ShippingSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingSite

`func (o *Invoice) SetShippingSite(v SiteReference)`

SetShippingSite sets ShippingSite field to given value.

### HasShippingSite

`func (o *Invoice) HasShippingSite() bool`

HasShippingSite returns a boolean if a field has been set.

### GetShippingSiteAddressLine1

`func (o *Invoice) GetShippingSiteAddressLine1() string`

GetShippingSiteAddressLine1 returns the ShippingSiteAddressLine1 field if non-nil, zero value otherwise.

### GetShippingSiteAddressLine1Ok

`func (o *Invoice) GetShippingSiteAddressLine1Ok() (*string, bool)`

GetShippingSiteAddressLine1Ok returns a tuple with the ShippingSiteAddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingSiteAddressLine1

`func (o *Invoice) SetShippingSiteAddressLine1(v string)`

SetShippingSiteAddressLine1 sets ShippingSiteAddressLine1 field to given value.

### HasShippingSiteAddressLine1

`func (o *Invoice) HasShippingSiteAddressLine1() bool`

HasShippingSiteAddressLine1 returns a boolean if a field has been set.

### GetShippingSiteAddressLine2

`func (o *Invoice) GetShippingSiteAddressLine2() string`

GetShippingSiteAddressLine2 returns the ShippingSiteAddressLine2 field if non-nil, zero value otherwise.

### GetShippingSiteAddressLine2Ok

`func (o *Invoice) GetShippingSiteAddressLine2Ok() (*string, bool)`

GetShippingSiteAddressLine2Ok returns a tuple with the ShippingSiteAddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingSiteAddressLine2

`func (o *Invoice) SetShippingSiteAddressLine2(v string)`

SetShippingSiteAddressLine2 sets ShippingSiteAddressLine2 field to given value.

### HasShippingSiteAddressLine2

`func (o *Invoice) HasShippingSiteAddressLine2() bool`

HasShippingSiteAddressLine2 returns a boolean if a field has been set.

### GetShippingSiteCity

`func (o *Invoice) GetShippingSiteCity() string`

GetShippingSiteCity returns the ShippingSiteCity field if non-nil, zero value otherwise.

### GetShippingSiteCityOk

`func (o *Invoice) GetShippingSiteCityOk() (*string, bool)`

GetShippingSiteCityOk returns a tuple with the ShippingSiteCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingSiteCity

`func (o *Invoice) SetShippingSiteCity(v string)`

SetShippingSiteCity sets ShippingSiteCity field to given value.

### HasShippingSiteCity

`func (o *Invoice) HasShippingSiteCity() bool`

HasShippingSiteCity returns a boolean if a field has been set.

### GetShippingSiteState

`func (o *Invoice) GetShippingSiteState() string`

GetShippingSiteState returns the ShippingSiteState field if non-nil, zero value otherwise.

### GetShippingSiteStateOk

`func (o *Invoice) GetShippingSiteStateOk() (*string, bool)`

GetShippingSiteStateOk returns a tuple with the ShippingSiteState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingSiteState

`func (o *Invoice) SetShippingSiteState(v string)`

SetShippingSiteState sets ShippingSiteState field to given value.

### HasShippingSiteState

`func (o *Invoice) HasShippingSiteState() bool`

HasShippingSiteState returns a boolean if a field has been set.

### GetShippingSiteZip

`func (o *Invoice) GetShippingSiteZip() string`

GetShippingSiteZip returns the ShippingSiteZip field if non-nil, zero value otherwise.

### GetShippingSiteZipOk

`func (o *Invoice) GetShippingSiteZipOk() (*string, bool)`

GetShippingSiteZipOk returns a tuple with the ShippingSiteZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingSiteZip

`func (o *Invoice) SetShippingSiteZip(v string)`

SetShippingSiteZip sets ShippingSiteZip field to given value.

### HasShippingSiteZip

`func (o *Invoice) HasShippingSiteZip() bool`

HasShippingSiteZip returns a boolean if a field has been set.

### GetShippingSiteCountry

`func (o *Invoice) GetShippingSiteCountry() string`

GetShippingSiteCountry returns the ShippingSiteCountry field if non-nil, zero value otherwise.

### GetShippingSiteCountryOk

`func (o *Invoice) GetShippingSiteCountryOk() (*string, bool)`

GetShippingSiteCountryOk returns a tuple with the ShippingSiteCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingSiteCountry

`func (o *Invoice) SetShippingSiteCountry(v string)`

SetShippingSiteCountry sets ShippingSiteCountry field to given value.

### HasShippingSiteCountry

`func (o *Invoice) HasShippingSiteCountry() bool`

HasShippingSiteCountry returns a boolean if a field has been set.

### GetBillingTerms

`func (o *Invoice) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *Invoice) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *Invoice) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *Invoice) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetReference

`func (o *Invoice) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Invoice) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Invoice) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Invoice) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetCustomerPO

`func (o *Invoice) GetCustomerPO() string`

GetCustomerPO returns the CustomerPO field if non-nil, zero value otherwise.

### GetCustomerPOOk

`func (o *Invoice) GetCustomerPOOk() (*string, bool)`

GetCustomerPOOk returns a tuple with the CustomerPO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPO

`func (o *Invoice) SetCustomerPO(v string)`

SetCustomerPO sets CustomerPO field to given value.

### HasCustomerPO

`func (o *Invoice) HasCustomerPO() bool`

HasCustomerPO returns a boolean if a field has been set.

### GetTemplateSetupId

`func (o *Invoice) GetTemplateSetupId() int32`

GetTemplateSetupId returns the TemplateSetupId field if non-nil, zero value otherwise.

### GetTemplateSetupIdOk

`func (o *Invoice) GetTemplateSetupIdOk() (*int32, bool)`

GetTemplateSetupIdOk returns a tuple with the TemplateSetupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSetupId

`func (o *Invoice) SetTemplateSetupId(v int32)`

SetTemplateSetupId sets TemplateSetupId field to given value.

### HasTemplateSetupId

`func (o *Invoice) HasTemplateSetupId() bool`

HasTemplateSetupId returns a boolean if a field has been set.

### SetTemplateSetupIdNil

`func (o *Invoice) SetTemplateSetupIdNil(b bool)`

 SetTemplateSetupIdNil sets the value for TemplateSetupId to be an explicit nil

### UnsetTemplateSetupId
`func (o *Invoice) UnsetTemplateSetupId()`

UnsetTemplateSetupId ensures that no value is present for TemplateSetupId, not even an explicit nil
### GetInvoiceTemplate

`func (o *Invoice) GetInvoiceTemplate() InvoiceTemplateDetailReference`

GetInvoiceTemplate returns the InvoiceTemplate field if non-nil, zero value otherwise.

### GetInvoiceTemplateOk

`func (o *Invoice) GetInvoiceTemplateOk() (*InvoiceTemplateDetailReference, bool)`

GetInvoiceTemplateOk returns a tuple with the InvoiceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTemplate

`func (o *Invoice) SetInvoiceTemplate(v InvoiceTemplateDetailReference)`

SetInvoiceTemplate sets InvoiceTemplate field to given value.

### HasInvoiceTemplate

`func (o *Invoice) HasInvoiceTemplate() bool`

HasInvoiceTemplate returns a boolean if a field has been set.

### GetEmailTemplateId

`func (o *Invoice) GetEmailTemplateId() int32`

GetEmailTemplateId returns the EmailTemplateId field if non-nil, zero value otherwise.

### GetEmailTemplateIdOk

`func (o *Invoice) GetEmailTemplateIdOk() (*int32, bool)`

GetEmailTemplateIdOk returns a tuple with the EmailTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateId

`func (o *Invoice) SetEmailTemplateId(v int32)`

SetEmailTemplateId sets EmailTemplateId field to given value.

### HasEmailTemplateId

`func (o *Invoice) HasEmailTemplateId() bool`

HasEmailTemplateId returns a boolean if a field has been set.

### SetEmailTemplateIdNil

`func (o *Invoice) SetEmailTemplateIdNil(b bool)`

 SetEmailTemplateIdNil sets the value for EmailTemplateId to be an explicit nil

### UnsetEmailTemplateId
`func (o *Invoice) UnsetEmailTemplateId()`

UnsetEmailTemplateId ensures that no value is present for EmailTemplateId, not even an explicit nil
### GetAddToBatchEmailList

`func (o *Invoice) GetAddToBatchEmailList() bool`

GetAddToBatchEmailList returns the AddToBatchEmailList field if non-nil, zero value otherwise.

### GetAddToBatchEmailListOk

`func (o *Invoice) GetAddToBatchEmailListOk() (*bool, bool)`

GetAddToBatchEmailListOk returns a tuple with the AddToBatchEmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToBatchEmailList

`func (o *Invoice) SetAddToBatchEmailList(v bool)`

SetAddToBatchEmailList sets AddToBatchEmailList field to given value.

### HasAddToBatchEmailList

`func (o *Invoice) HasAddToBatchEmailList() bool`

HasAddToBatchEmailList returns a boolean if a field has been set.

### SetAddToBatchEmailListNil

`func (o *Invoice) SetAddToBatchEmailListNil(b bool)`

 SetAddToBatchEmailListNil sets the value for AddToBatchEmailList to be an explicit nil

### UnsetAddToBatchEmailList
`func (o *Invoice) UnsetAddToBatchEmailList()`

UnsetAddToBatchEmailList ensures that no value is present for AddToBatchEmailList, not even an explicit nil
### GetDate

`func (o *Invoice) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Invoice) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Invoice) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *Invoice) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetRestrictDownpaymentFlag

`func (o *Invoice) GetRestrictDownpaymentFlag() bool`

GetRestrictDownpaymentFlag returns the RestrictDownpaymentFlag field if non-nil, zero value otherwise.

### GetRestrictDownpaymentFlagOk

`func (o *Invoice) GetRestrictDownpaymentFlagOk() (*bool, bool)`

GetRestrictDownpaymentFlagOk returns a tuple with the RestrictDownpaymentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDownpaymentFlag

`func (o *Invoice) SetRestrictDownpaymentFlag(v bool)`

SetRestrictDownpaymentFlag sets RestrictDownpaymentFlag field to given value.

### HasRestrictDownpaymentFlag

`func (o *Invoice) HasRestrictDownpaymentFlag() bool`

HasRestrictDownpaymentFlag returns a boolean if a field has been set.

### SetRestrictDownpaymentFlagNil

`func (o *Invoice) SetRestrictDownpaymentFlagNil(b bool)`

 SetRestrictDownpaymentFlagNil sets the value for RestrictDownpaymentFlag to be an explicit nil

### UnsetRestrictDownpaymentFlag
`func (o *Invoice) UnsetRestrictDownpaymentFlag()`

UnsetRestrictDownpaymentFlag ensures that no value is present for RestrictDownpaymentFlag, not even an explicit nil
### GetLocationId

`func (o *Invoice) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Invoice) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Invoice) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *Invoice) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *Invoice) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *Invoice) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetLocation

`func (o *Invoice) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Invoice) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Invoice) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Invoice) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartmentId

`func (o *Invoice) GetDepartmentId() int32`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *Invoice) GetDepartmentIdOk() (*int32, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *Invoice) SetDepartmentId(v int32)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *Invoice) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### SetDepartmentIdNil

`func (o *Invoice) SetDepartmentIdNil(b bool)`

 SetDepartmentIdNil sets the value for DepartmentId to be an explicit nil

### UnsetDepartmentId
`func (o *Invoice) UnsetDepartmentId()`

UnsetDepartmentId ensures that no value is present for DepartmentId, not even an explicit nil
### GetDepartment

`func (o *Invoice) GetDepartment() BillingUnitReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Invoice) GetDepartmentOk() (*BillingUnitReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Invoice) SetDepartment(v BillingUnitReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Invoice) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetTerritoryId

`func (o *Invoice) GetTerritoryId() int32`

GetTerritoryId returns the TerritoryId field if non-nil, zero value otherwise.

### GetTerritoryIdOk

`func (o *Invoice) GetTerritoryIdOk() (*int32, bool)`

GetTerritoryIdOk returns a tuple with the TerritoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritoryId

`func (o *Invoice) SetTerritoryId(v int32)`

SetTerritoryId sets TerritoryId field to given value.

### HasTerritoryId

`func (o *Invoice) HasTerritoryId() bool`

HasTerritoryId returns a boolean if a field has been set.

### SetTerritoryIdNil

`func (o *Invoice) SetTerritoryIdNil(b bool)`

 SetTerritoryIdNil sets the value for TerritoryId to be an explicit nil

### UnsetTerritoryId
`func (o *Invoice) UnsetTerritoryId()`

UnsetTerritoryId ensures that no value is present for TerritoryId, not even an explicit nil
### GetTerritory

`func (o *Invoice) GetTerritory() SystemLocationReference`

GetTerritory returns the Territory field if non-nil, zero value otherwise.

### GetTerritoryOk

`func (o *Invoice) GetTerritoryOk() (*SystemLocationReference, bool)`

GetTerritoryOk returns a tuple with the Territory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritory

`func (o *Invoice) SetTerritory(v SystemLocationReference)`

SetTerritory sets Territory field to given value.

### HasTerritory

`func (o *Invoice) HasTerritory() bool`

HasTerritory returns a boolean if a field has been set.

### GetTopComment

`func (o *Invoice) GetTopComment() string`

GetTopComment returns the TopComment field if non-nil, zero value otherwise.

### GetTopCommentOk

`func (o *Invoice) GetTopCommentOk() (*string, bool)`

GetTopCommentOk returns a tuple with the TopComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopComment

`func (o *Invoice) SetTopComment(v string)`

SetTopComment sets TopComment field to given value.

### HasTopComment

`func (o *Invoice) HasTopComment() bool`

HasTopComment returns a boolean if a field has been set.

### GetBottomComment

`func (o *Invoice) GetBottomComment() string`

GetBottomComment returns the BottomComment field if non-nil, zero value otherwise.

### GetBottomCommentOk

`func (o *Invoice) GetBottomCommentOk() (*string, bool)`

GetBottomCommentOk returns a tuple with the BottomComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomComment

`func (o *Invoice) SetBottomComment(v string)`

SetBottomComment sets BottomComment field to given value.

### HasBottomComment

`func (o *Invoice) HasBottomComment() bool`

HasBottomComment returns a boolean if a field has been set.

### GetTaxableFlag

`func (o *Invoice) GetTaxableFlag() bool`

GetTaxableFlag returns the TaxableFlag field if non-nil, zero value otherwise.

### GetTaxableFlagOk

`func (o *Invoice) GetTaxableFlagOk() (*bool, bool)`

GetTaxableFlagOk returns a tuple with the TaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableFlag

`func (o *Invoice) SetTaxableFlag(v bool)`

SetTaxableFlag sets TaxableFlag field to given value.

### HasTaxableFlag

`func (o *Invoice) HasTaxableFlag() bool`

HasTaxableFlag returns a boolean if a field has been set.

### SetTaxableFlagNil

`func (o *Invoice) SetTaxableFlagNil(b bool)`

 SetTaxableFlagNil sets the value for TaxableFlag to be an explicit nil

### UnsetTaxableFlag
`func (o *Invoice) UnsetTaxableFlag()`

UnsetTaxableFlag ensures that no value is present for TaxableFlag, not even an explicit nil
### GetTaxCode

`func (o *Invoice) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *Invoice) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *Invoice) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *Invoice) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetInternalNotes

`func (o *Invoice) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *Invoice) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *Invoice) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *Invoice) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetDownpaymentPreviouslyTaxedFlag

`func (o *Invoice) GetDownpaymentPreviouslyTaxedFlag() bool`

GetDownpaymentPreviouslyTaxedFlag returns the DownpaymentPreviouslyTaxedFlag field if non-nil, zero value otherwise.

### GetDownpaymentPreviouslyTaxedFlagOk

`func (o *Invoice) GetDownpaymentPreviouslyTaxedFlagOk() (*bool, bool)`

GetDownpaymentPreviouslyTaxedFlagOk returns a tuple with the DownpaymentPreviouslyTaxedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownpaymentPreviouslyTaxedFlag

`func (o *Invoice) SetDownpaymentPreviouslyTaxedFlag(v bool)`

SetDownpaymentPreviouslyTaxedFlag sets DownpaymentPreviouslyTaxedFlag field to given value.

### HasDownpaymentPreviouslyTaxedFlag

`func (o *Invoice) HasDownpaymentPreviouslyTaxedFlag() bool`

HasDownpaymentPreviouslyTaxedFlag returns a boolean if a field has been set.

### SetDownpaymentPreviouslyTaxedFlagNil

`func (o *Invoice) SetDownpaymentPreviouslyTaxedFlagNil(b bool)`

 SetDownpaymentPreviouslyTaxedFlagNil sets the value for DownpaymentPreviouslyTaxedFlag to be an explicit nil

### UnsetDownpaymentPreviouslyTaxedFlag
`func (o *Invoice) UnsetDownpaymentPreviouslyTaxedFlag()`

UnsetDownpaymentPreviouslyTaxedFlag ensures that no value is present for DownpaymentPreviouslyTaxedFlag, not even an explicit nil
### GetServiceTotal

`func (o *Invoice) GetServiceTotal() float64`

GetServiceTotal returns the ServiceTotal field if non-nil, zero value otherwise.

### GetServiceTotalOk

`func (o *Invoice) GetServiceTotalOk() (*float64, bool)`

GetServiceTotalOk returns a tuple with the ServiceTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTotal

`func (o *Invoice) SetServiceTotal(v float64)`

SetServiceTotal sets ServiceTotal field to given value.

### HasServiceTotal

`func (o *Invoice) HasServiceTotal() bool`

HasServiceTotal returns a boolean if a field has been set.

### SetServiceTotalNil

`func (o *Invoice) SetServiceTotalNil(b bool)`

 SetServiceTotalNil sets the value for ServiceTotal to be an explicit nil

### UnsetServiceTotal
`func (o *Invoice) UnsetServiceTotal()`

UnsetServiceTotal ensures that no value is present for ServiceTotal, not even an explicit nil
### GetOverrideDownPaymentAmountFlag

`func (o *Invoice) GetOverrideDownPaymentAmountFlag() bool`

GetOverrideDownPaymentAmountFlag returns the OverrideDownPaymentAmountFlag field if non-nil, zero value otherwise.

### GetOverrideDownPaymentAmountFlagOk

`func (o *Invoice) GetOverrideDownPaymentAmountFlagOk() (*bool, bool)`

GetOverrideDownPaymentAmountFlagOk returns a tuple with the OverrideDownPaymentAmountFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDownPaymentAmountFlag

`func (o *Invoice) SetOverrideDownPaymentAmountFlag(v bool)`

SetOverrideDownPaymentAmountFlag sets OverrideDownPaymentAmountFlag field to given value.

### HasOverrideDownPaymentAmountFlag

`func (o *Invoice) HasOverrideDownPaymentAmountFlag() bool`

HasOverrideDownPaymentAmountFlag returns a boolean if a field has been set.

### SetOverrideDownPaymentAmountFlagNil

`func (o *Invoice) SetOverrideDownPaymentAmountFlagNil(b bool)`

 SetOverrideDownPaymentAmountFlagNil sets the value for OverrideDownPaymentAmountFlag to be an explicit nil

### UnsetOverrideDownPaymentAmountFlag
`func (o *Invoice) UnsetOverrideDownPaymentAmountFlag()`

UnsetOverrideDownPaymentAmountFlag ensures that no value is present for OverrideDownPaymentAmountFlag, not even an explicit nil
### GetCurrency

`func (o *Invoice) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Invoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDueDate

`func (o *Invoice) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Invoice) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Invoice) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Invoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetExpenseTotal

`func (o *Invoice) GetExpenseTotal() float64`

GetExpenseTotal returns the ExpenseTotal field if non-nil, zero value otherwise.

### GetExpenseTotalOk

`func (o *Invoice) GetExpenseTotalOk() (*float64, bool)`

GetExpenseTotalOk returns a tuple with the ExpenseTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseTotal

`func (o *Invoice) SetExpenseTotal(v float64)`

SetExpenseTotal sets ExpenseTotal field to given value.

### HasExpenseTotal

`func (o *Invoice) HasExpenseTotal() bool`

HasExpenseTotal returns a boolean if a field has been set.

### SetExpenseTotalNil

`func (o *Invoice) SetExpenseTotalNil(b bool)`

 SetExpenseTotalNil sets the value for ExpenseTotal to be an explicit nil

### UnsetExpenseTotal
`func (o *Invoice) UnsetExpenseTotal()`

UnsetExpenseTotal ensures that no value is present for ExpenseTotal, not even an explicit nil
### GetProductTotal

`func (o *Invoice) GetProductTotal() float64`

GetProductTotal returns the ProductTotal field if non-nil, zero value otherwise.

### GetProductTotalOk

`func (o *Invoice) GetProductTotalOk() (*float64, bool)`

GetProductTotalOk returns a tuple with the ProductTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTotal

`func (o *Invoice) SetProductTotal(v float64)`

SetProductTotal sets ProductTotal field to given value.

### HasProductTotal

`func (o *Invoice) HasProductTotal() bool`

HasProductTotal returns a boolean if a field has been set.

### SetProductTotalNil

`func (o *Invoice) SetProductTotalNil(b bool)`

 SetProductTotalNil sets the value for ProductTotal to be an explicit nil

### UnsetProductTotal
`func (o *Invoice) UnsetProductTotal()`

UnsetProductTotal ensures that no value is present for ProductTotal, not even an explicit nil
### GetPreviousProgressApplied

`func (o *Invoice) GetPreviousProgressApplied() float64`

GetPreviousProgressApplied returns the PreviousProgressApplied field if non-nil, zero value otherwise.

### GetPreviousProgressAppliedOk

`func (o *Invoice) GetPreviousProgressAppliedOk() (*float64, bool)`

GetPreviousProgressAppliedOk returns a tuple with the PreviousProgressApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousProgressApplied

`func (o *Invoice) SetPreviousProgressApplied(v float64)`

SetPreviousProgressApplied sets PreviousProgressApplied field to given value.

### HasPreviousProgressApplied

`func (o *Invoice) HasPreviousProgressApplied() bool`

HasPreviousProgressApplied returns a boolean if a field has been set.

### SetPreviousProgressAppliedNil

`func (o *Invoice) SetPreviousProgressAppliedNil(b bool)`

 SetPreviousProgressAppliedNil sets the value for PreviousProgressApplied to be an explicit nil

### UnsetPreviousProgressApplied
`func (o *Invoice) UnsetPreviousProgressApplied()`

UnsetPreviousProgressApplied ensures that no value is present for PreviousProgressApplied, not even an explicit nil
### GetServiceAdjustmentAmount

`func (o *Invoice) GetServiceAdjustmentAmount() float64`

GetServiceAdjustmentAmount returns the ServiceAdjustmentAmount field if non-nil, zero value otherwise.

### GetServiceAdjustmentAmountOk

`func (o *Invoice) GetServiceAdjustmentAmountOk() (*float64, bool)`

GetServiceAdjustmentAmountOk returns a tuple with the ServiceAdjustmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAdjustmentAmount

`func (o *Invoice) SetServiceAdjustmentAmount(v float64)`

SetServiceAdjustmentAmount sets ServiceAdjustmentAmount field to given value.

### HasServiceAdjustmentAmount

`func (o *Invoice) HasServiceAdjustmentAmount() bool`

HasServiceAdjustmentAmount returns a boolean if a field has been set.

### SetServiceAdjustmentAmountNil

`func (o *Invoice) SetServiceAdjustmentAmountNil(b bool)`

 SetServiceAdjustmentAmountNil sets the value for ServiceAdjustmentAmount to be an explicit nil

### UnsetServiceAdjustmentAmount
`func (o *Invoice) UnsetServiceAdjustmentAmount()`

UnsetServiceAdjustmentAmount ensures that no value is present for ServiceAdjustmentAmount, not even an explicit nil
### GetAgreementAmount

`func (o *Invoice) GetAgreementAmount() float64`

GetAgreementAmount returns the AgreementAmount field if non-nil, zero value otherwise.

### GetAgreementAmountOk

`func (o *Invoice) GetAgreementAmountOk() (*float64, bool)`

GetAgreementAmountOk returns a tuple with the AgreementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementAmount

`func (o *Invoice) SetAgreementAmount(v float64)`

SetAgreementAmount sets AgreementAmount field to given value.

### HasAgreementAmount

`func (o *Invoice) HasAgreementAmount() bool`

HasAgreementAmount returns a boolean if a field has been set.

### SetAgreementAmountNil

`func (o *Invoice) SetAgreementAmountNil(b bool)`

 SetAgreementAmountNil sets the value for AgreementAmount to be an explicit nil

### UnsetAgreementAmount
`func (o *Invoice) UnsetAgreementAmount()`

UnsetAgreementAmount ensures that no value is present for AgreementAmount, not even an explicit nil
### GetDownpaymentApplied

`func (o *Invoice) GetDownpaymentApplied() float64`

GetDownpaymentApplied returns the DownpaymentApplied field if non-nil, zero value otherwise.

### GetDownpaymentAppliedOk

`func (o *Invoice) GetDownpaymentAppliedOk() (*float64, bool)`

GetDownpaymentAppliedOk returns a tuple with the DownpaymentApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownpaymentApplied

`func (o *Invoice) SetDownpaymentApplied(v float64)`

SetDownpaymentApplied sets DownpaymentApplied field to given value.

### HasDownpaymentApplied

`func (o *Invoice) HasDownpaymentApplied() bool`

HasDownpaymentApplied returns a boolean if a field has been set.

### SetDownpaymentAppliedNil

`func (o *Invoice) SetDownpaymentAppliedNil(b bool)`

 SetDownpaymentAppliedNil sets the value for DownpaymentApplied to be an explicit nil

### UnsetDownpaymentApplied
`func (o *Invoice) UnsetDownpaymentApplied()`

UnsetDownpaymentApplied ensures that no value is present for DownpaymentApplied, not even an explicit nil
### GetSubtotal

`func (o *Invoice) GetSubtotal() float64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *Invoice) GetSubtotalOk() (*float64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *Invoice) SetSubtotal(v float64)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *Invoice) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### SetSubtotalNil

`func (o *Invoice) SetSubtotalNil(b bool)`

 SetSubtotalNil sets the value for Subtotal to be an explicit nil

### UnsetSubtotal
`func (o *Invoice) UnsetSubtotal()`

UnsetSubtotal ensures that no value is present for Subtotal, not even an explicit nil
### GetTotal

`func (o *Invoice) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Invoice) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Invoice) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Invoice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *Invoice) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *Invoice) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetRemainingDownpayment

`func (o *Invoice) GetRemainingDownpayment() float64`

GetRemainingDownpayment returns the RemainingDownpayment field if non-nil, zero value otherwise.

### GetRemainingDownpaymentOk

`func (o *Invoice) GetRemainingDownpaymentOk() (*float64, bool)`

GetRemainingDownpaymentOk returns a tuple with the RemainingDownpayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingDownpayment

`func (o *Invoice) SetRemainingDownpayment(v float64)`

SetRemainingDownpayment sets RemainingDownpayment field to given value.

### HasRemainingDownpayment

`func (o *Invoice) HasRemainingDownpayment() bool`

HasRemainingDownpayment returns a boolean if a field has been set.

### SetRemainingDownpaymentNil

`func (o *Invoice) SetRemainingDownpaymentNil(b bool)`

 SetRemainingDownpaymentNil sets the value for RemainingDownpayment to be an explicit nil

### UnsetRemainingDownpayment
`func (o *Invoice) UnsetRemainingDownpayment()`

UnsetRemainingDownpayment ensures that no value is present for RemainingDownpayment, not even an explicit nil
### GetSalesTax

`func (o *Invoice) GetSalesTax() float64`

GetSalesTax returns the SalesTax field if non-nil, zero value otherwise.

### GetSalesTaxOk

`func (o *Invoice) GetSalesTaxOk() (*float64, bool)`

GetSalesTaxOk returns a tuple with the SalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTax

`func (o *Invoice) SetSalesTax(v float64)`

SetSalesTax sets SalesTax field to given value.

### HasSalesTax

`func (o *Invoice) HasSalesTax() bool`

HasSalesTax returns a boolean if a field has been set.

### SetSalesTaxNil

`func (o *Invoice) SetSalesTaxNil(b bool)`

 SetSalesTaxNil sets the value for SalesTax to be an explicit nil

### UnsetSalesTax
`func (o *Invoice) UnsetSalesTax()`

UnsetSalesTax ensures that no value is present for SalesTax, not even an explicit nil
### GetAdjustmentReason

`func (o *Invoice) GetAdjustmentReason() string`

GetAdjustmentReason returns the AdjustmentReason field if non-nil, zero value otherwise.

### GetAdjustmentReasonOk

`func (o *Invoice) GetAdjustmentReasonOk() (*string, bool)`

GetAdjustmentReasonOk returns a tuple with the AdjustmentReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentReason

`func (o *Invoice) SetAdjustmentReason(v string)`

SetAdjustmentReason sets AdjustmentReason field to given value.

### HasAdjustmentReason

`func (o *Invoice) HasAdjustmentReason() bool`

HasAdjustmentReason returns a boolean if a field has been set.

### GetAdjustedBy

`func (o *Invoice) GetAdjustedBy() string`

GetAdjustedBy returns the AdjustedBy field if non-nil, zero value otherwise.

### GetAdjustedByOk

`func (o *Invoice) GetAdjustedByOk() (*string, bool)`

GetAdjustedByOk returns a tuple with the AdjustedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedBy

`func (o *Invoice) SetAdjustedBy(v string)`

SetAdjustedBy sets AdjustedBy field to given value.

### HasAdjustedBy

`func (o *Invoice) HasAdjustedBy() bool`

HasAdjustedBy returns a boolean if a field has been set.

### GetClosedBy

`func (o *Invoice) GetClosedBy() string`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *Invoice) GetClosedByOk() (*string, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *Invoice) SetClosedBy(v string)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *Invoice) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetPayments

`func (o *Invoice) GetPayments() float64`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *Invoice) GetPaymentsOk() (*float64, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *Invoice) SetPayments(v float64)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *Invoice) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### SetPaymentsNil

`func (o *Invoice) SetPaymentsNil(b bool)`

 SetPaymentsNil sets the value for Payments to be an explicit nil

### UnsetPayments
`func (o *Invoice) UnsetPayments()`

UnsetPayments ensures that no value is present for Payments, not even an explicit nil
### GetCredits

`func (o *Invoice) GetCredits() float64`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *Invoice) GetCreditsOk() (*float64, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *Invoice) SetCredits(v float64)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *Invoice) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### SetCreditsNil

`func (o *Invoice) SetCreditsNil(b bool)`

 SetCreditsNil sets the value for Credits to be an explicit nil

### UnsetCredits
`func (o *Invoice) UnsetCredits()`

UnsetCredits ensures that no value is present for Credits, not even an explicit nil
### GetBalance

`func (o *Invoice) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Invoice) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Invoice) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *Invoice) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalanceNil

`func (o *Invoice) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *Invoice) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetSpecialInvoiceFlag

`func (o *Invoice) GetSpecialInvoiceFlag() bool`

GetSpecialInvoiceFlag returns the SpecialInvoiceFlag field if non-nil, zero value otherwise.

### GetSpecialInvoiceFlagOk

`func (o *Invoice) GetSpecialInvoiceFlagOk() (*bool, bool)`

GetSpecialInvoiceFlagOk returns a tuple with the SpecialInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInvoiceFlag

`func (o *Invoice) SetSpecialInvoiceFlag(v bool)`

SetSpecialInvoiceFlag sets SpecialInvoiceFlag field to given value.

### HasSpecialInvoiceFlag

`func (o *Invoice) HasSpecialInvoiceFlag() bool`

HasSpecialInvoiceFlag returns a boolean if a field has been set.

### SetSpecialInvoiceFlagNil

`func (o *Invoice) SetSpecialInvoiceFlagNil(b bool)`

 SetSpecialInvoiceFlagNil sets the value for SpecialInvoiceFlag to be an explicit nil

### UnsetSpecialInvoiceFlag
`func (o *Invoice) UnsetSpecialInvoiceFlag()`

UnsetSpecialInvoiceFlag ensures that no value is present for SpecialInvoiceFlag, not even an explicit nil
### GetBillingSetupReference

`func (o *Invoice) GetBillingSetupReference() BillingSetupReference`

GetBillingSetupReference returns the BillingSetupReference field if non-nil, zero value otherwise.

### GetBillingSetupReferenceOk

`func (o *Invoice) GetBillingSetupReferenceOk() (*BillingSetupReference, bool)`

GetBillingSetupReferenceOk returns a tuple with the BillingSetupReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSetupReference

`func (o *Invoice) SetBillingSetupReference(v BillingSetupReference)`

SetBillingSetupReference sets BillingSetupReference field to given value.

### HasBillingSetupReference

`func (o *Invoice) HasBillingSetupReference() bool`

HasBillingSetupReference returns a boolean if a field has been set.

### GetTicket

`func (o *Invoice) GetTicket() TicketReference`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *Invoice) GetTicketOk() (*TicketReference, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *Invoice) SetTicket(v TicketReference)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *Invoice) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetProject

`func (o *Invoice) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Invoice) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Invoice) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *Invoice) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPhase

`func (o *Invoice) GetPhase() ProjectPhaseReference`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Invoice) GetPhaseOk() (*ProjectPhaseReference, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Invoice) SetPhase(v ProjectPhaseReference)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Invoice) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetSalesOrder

`func (o *Invoice) GetSalesOrder() SalesOrderReference`

GetSalesOrder returns the SalesOrder field if non-nil, zero value otherwise.

### GetSalesOrderOk

`func (o *Invoice) GetSalesOrderOk() (*SalesOrderReference, bool)`

GetSalesOrderOk returns a tuple with the SalesOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrder

`func (o *Invoice) SetSalesOrder(v SalesOrderReference)`

SetSalesOrder sets SalesOrder field to given value.

### HasSalesOrder

`func (o *Invoice) HasSalesOrder() bool`

HasSalesOrder returns a boolean if a field has been set.

### GetAgreement

`func (o *Invoice) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *Invoice) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *Invoice) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *Invoice) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetGlBatch

`func (o *Invoice) GetGlBatch() BatchReference`

GetGlBatch returns the GlBatch field if non-nil, zero value otherwise.

### GetGlBatchOk

`func (o *Invoice) GetGlBatchOk() (*BatchReference, bool)`

GetGlBatchOk returns a tuple with the GlBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlBatch

`func (o *Invoice) SetGlBatch(v BatchReference)`

SetGlBatch sets GlBatch field to given value.

### HasGlBatch

`func (o *Invoice) HasGlBatch() bool`

HasGlBatch returns a boolean if a field has been set.

### GetUnbatchedBatch

`func (o *Invoice) GetUnbatchedBatch() BatchReference`

GetUnbatchedBatch returns the UnbatchedBatch field if non-nil, zero value otherwise.

### GetUnbatchedBatchOk

`func (o *Invoice) GetUnbatchedBatchOk() (*BatchReference, bool)`

GetUnbatchedBatchOk returns a tuple with the UnbatchedBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbatchedBatch

`func (o *Invoice) SetUnbatchedBatch(v BatchReference)`

SetUnbatchedBatch sets UnbatchedBatch field to given value.

### HasUnbatchedBatch

`func (o *Invoice) HasUnbatchedBatch() bool`

HasUnbatchedBatch returns a boolean if a field has been set.

### GetInfo

`func (o *Invoice) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Invoice) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Invoice) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Invoice) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *Invoice) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Invoice) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Invoice) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Invoice) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


