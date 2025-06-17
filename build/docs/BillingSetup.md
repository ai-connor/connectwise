# BillingSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RemitName** | **string** |  Max length: 50; | 
**Location** | [**SystemLocationReference**](SystemLocationReference.md) |  | 
**AddressOne** | Pointer to **string** |  Max length: 50; | [optional] 
**AddressTwo** | Pointer to **string** |  Max length: 50; | [optional] 
**City** | Pointer to **string** |  Max length: 50; | [optional] 
**State** | Pointer to [**StateReference**](StateReference.md) |  | [optional] 
**Zip** | Pointer to **string** |  Max length: 12; | [optional] 
**Country** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**Phone** | Pointer to **string** |  Max length: 15; | [optional] 
**InvoiceTitle** | **string** |  Max length: 50; | 
**PayableName** | **string** |  Max length: 50; | 
**Topcomment** | Pointer to **string** |  Max length: 4000; | [optional] 
**InvoiceFooter** | Pointer to **string** |  Max length: 500; | [optional] 
**QuoteFooter** | Pointer to **string** |  Max length: 1000; | [optional] 
**OverallInvoiceDefault** | [**InvoiceTemplateReference**](InvoiceTemplateReference.md) |  | 
**StandardInvoiceActual** | Pointer to [**InvoiceTemplateReference**](InvoiceTemplateReference.md) |  | [optional] 
**StandardInvoiceFixed** | Pointer to [**InvoiceTemplateReference**](InvoiceTemplateReference.md) |  | [optional] 
**ProgressInvoice** | Pointer to [**InvoiceTemplateReference**](InvoiceTemplateReference.md) |  | [optional] 
**AgreementInvoice** | Pointer to [**InvoiceTemplateReference**](InvoiceTemplateReference.md) |  | [optional] 
**CreditMemoInvoice** | Pointer to [**InvoiceTemplateReference**](InvoiceTemplateReference.md) |  | [optional] 
**DownPaymentInvoice** | Pointer to [**InvoiceTemplateReference**](InvoiceTemplateReference.md) |  | [optional] 
**MiscInvoice** | Pointer to [**InvoiceTemplateReference**](InvoiceTemplateReference.md) |  | [optional] 
**SalesOrderInvoice** | Pointer to [**InvoiceTemplateReference**](InvoiceTemplateReference.md) |  | [optional] 
**ExcludeDoNotBillTimeFlag** | Pointer to **NullableBool** |  | [optional] 
**ExcludeDoNotBillExpenseFlag** | Pointer to **NullableBool** |  | [optional] 
**ExcludeDoNotBillProductFlag** | Pointer to **NullableBool** |  | [optional] 
**PrefixSuffixFlag** | Pointer to **NullableString** |  | [optional] 
**PrefixSuffixText** | Pointer to **string** |  Max length: 5; | [optional] 
**ChargeAdjToFirmFlag** | Pointer to **NullableBool** |  | [optional] 
**NoWatermarkFlag** | Pointer to **NullableBool** |  | [optional] 
**DisplayTaxFlag** | Pointer to **NullableBool** |  | [optional] 
**AllowRestrictedDeptOnRoutingFlag** | Pointer to **NullableBool** |  | [optional] 
**BillTicketSeparatelyFlag** | Pointer to **NullableBool** |  | [optional] 
**BillTicketCompleteFlag** | Pointer to **NullableBool** |  | [optional] 
**BillTicketUnapprovedFlag** | Pointer to **NullableBool** |  | [optional] 
**BillProjectCompleteFlag** | Pointer to **NullableBool** |  | [optional] 
**BillProjectUnapprovedFlag** | Pointer to **NullableBool** |  | [optional] 
**ProgressTimeFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictProjectDownpaymentFlag** | Pointer to **NullableBool** |  | [optional] 
**BillSalesOrderCompleteFlag** | Pointer to **NullableBool** |  | [optional] 
**BillProductAfterShipFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictDownpaymentFlag** | Pointer to **NullableBool** |  | [optional] 
**CopyNonServiceProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**CopyServiceProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**CopyAgreementProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**PrintLogoFlag** | Pointer to **NullableBool** |  | [optional] 
**ReadReceiptFlag** | Pointer to **NullableBool** |  | [optional] 
**DeliveryReceiptFlag** | Pointer to **NullableBool** |  | [optional] 
**AttachXmlInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**DisableRoutingEmailFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailTemplate** | [**EmailTemplateReference**](EmailTemplateReference.md) |  | 
**LocalizedCountry** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**BusinessNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**CustomLabel** | Pointer to **string** |  Max length: 50; | [optional] 
**CustomText** | Pointer to **string** |  Max length: 500; | [optional] 
**CompanyCode** | Pointer to **string** |  Max length: 250; | [optional] 
**ExcludeAvalaraFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBillingSetup

`func NewBillingSetup(remitName string, location SystemLocationReference, invoiceTitle string, payableName string, overallInvoiceDefault InvoiceTemplateReference, emailTemplate EmailTemplateReference, ) *BillingSetup`

NewBillingSetup instantiates a new BillingSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSetupWithDefaults

`func NewBillingSetupWithDefaults() *BillingSetup`

NewBillingSetupWithDefaults instantiates a new BillingSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BillingSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemitName

`func (o *BillingSetup) GetRemitName() string`

GetRemitName returns the RemitName field if non-nil, zero value otherwise.

### GetRemitNameOk

`func (o *BillingSetup) GetRemitNameOk() (*string, bool)`

GetRemitNameOk returns a tuple with the RemitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemitName

`func (o *BillingSetup) SetRemitName(v string)`

SetRemitName sets RemitName field to given value.


### GetLocation

`func (o *BillingSetup) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BillingSetup) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BillingSetup) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.


### GetAddressOne

`func (o *BillingSetup) GetAddressOne() string`

GetAddressOne returns the AddressOne field if non-nil, zero value otherwise.

### GetAddressOneOk

`func (o *BillingSetup) GetAddressOneOk() (*string, bool)`

GetAddressOneOk returns a tuple with the AddressOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressOne

`func (o *BillingSetup) SetAddressOne(v string)`

SetAddressOne sets AddressOne field to given value.

### HasAddressOne

`func (o *BillingSetup) HasAddressOne() bool`

HasAddressOne returns a boolean if a field has been set.

### GetAddressTwo

`func (o *BillingSetup) GetAddressTwo() string`

GetAddressTwo returns the AddressTwo field if non-nil, zero value otherwise.

### GetAddressTwoOk

`func (o *BillingSetup) GetAddressTwoOk() (*string, bool)`

GetAddressTwoOk returns a tuple with the AddressTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTwo

`func (o *BillingSetup) SetAddressTwo(v string)`

SetAddressTwo sets AddressTwo field to given value.

### HasAddressTwo

`func (o *BillingSetup) HasAddressTwo() bool`

HasAddressTwo returns a boolean if a field has been set.

### GetCity

`func (o *BillingSetup) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BillingSetup) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BillingSetup) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BillingSetup) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *BillingSetup) GetState() StateReference`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BillingSetup) GetStateOk() (*StateReference, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BillingSetup) SetState(v StateReference)`

SetState sets State field to given value.

### HasState

`func (o *BillingSetup) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *BillingSetup) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *BillingSetup) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *BillingSetup) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *BillingSetup) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetCountry

`func (o *BillingSetup) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BillingSetup) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BillingSetup) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BillingSetup) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPhone

`func (o *BillingSetup) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BillingSetup) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BillingSetup) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BillingSetup) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetInvoiceTitle

`func (o *BillingSetup) GetInvoiceTitle() string`

GetInvoiceTitle returns the InvoiceTitle field if non-nil, zero value otherwise.

### GetInvoiceTitleOk

`func (o *BillingSetup) GetInvoiceTitleOk() (*string, bool)`

GetInvoiceTitleOk returns a tuple with the InvoiceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTitle

`func (o *BillingSetup) SetInvoiceTitle(v string)`

SetInvoiceTitle sets InvoiceTitle field to given value.


### GetPayableName

`func (o *BillingSetup) GetPayableName() string`

GetPayableName returns the PayableName field if non-nil, zero value otherwise.

### GetPayableNameOk

`func (o *BillingSetup) GetPayableNameOk() (*string, bool)`

GetPayableNameOk returns a tuple with the PayableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableName

`func (o *BillingSetup) SetPayableName(v string)`

SetPayableName sets PayableName field to given value.


### GetTopcomment

`func (o *BillingSetup) GetTopcomment() string`

GetTopcomment returns the Topcomment field if non-nil, zero value otherwise.

### GetTopcommentOk

`func (o *BillingSetup) GetTopcommentOk() (*string, bool)`

GetTopcommentOk returns a tuple with the Topcomment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopcomment

`func (o *BillingSetup) SetTopcomment(v string)`

SetTopcomment sets Topcomment field to given value.

### HasTopcomment

`func (o *BillingSetup) HasTopcomment() bool`

HasTopcomment returns a boolean if a field has been set.

### GetInvoiceFooter

`func (o *BillingSetup) GetInvoiceFooter() string`

GetInvoiceFooter returns the InvoiceFooter field if non-nil, zero value otherwise.

### GetInvoiceFooterOk

`func (o *BillingSetup) GetInvoiceFooterOk() (*string, bool)`

GetInvoiceFooterOk returns a tuple with the InvoiceFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceFooter

`func (o *BillingSetup) SetInvoiceFooter(v string)`

SetInvoiceFooter sets InvoiceFooter field to given value.

### HasInvoiceFooter

`func (o *BillingSetup) HasInvoiceFooter() bool`

HasInvoiceFooter returns a boolean if a field has been set.

### GetQuoteFooter

`func (o *BillingSetup) GetQuoteFooter() string`

GetQuoteFooter returns the QuoteFooter field if non-nil, zero value otherwise.

### GetQuoteFooterOk

`func (o *BillingSetup) GetQuoteFooterOk() (*string, bool)`

GetQuoteFooterOk returns a tuple with the QuoteFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteFooter

`func (o *BillingSetup) SetQuoteFooter(v string)`

SetQuoteFooter sets QuoteFooter field to given value.

### HasQuoteFooter

`func (o *BillingSetup) HasQuoteFooter() bool`

HasQuoteFooter returns a boolean if a field has been set.

### GetOverallInvoiceDefault

`func (o *BillingSetup) GetOverallInvoiceDefault() InvoiceTemplateReference`

GetOverallInvoiceDefault returns the OverallInvoiceDefault field if non-nil, zero value otherwise.

### GetOverallInvoiceDefaultOk

`func (o *BillingSetup) GetOverallInvoiceDefaultOk() (*InvoiceTemplateReference, bool)`

GetOverallInvoiceDefaultOk returns a tuple with the OverallInvoiceDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallInvoiceDefault

`func (o *BillingSetup) SetOverallInvoiceDefault(v InvoiceTemplateReference)`

SetOverallInvoiceDefault sets OverallInvoiceDefault field to given value.


### GetStandardInvoiceActual

`func (o *BillingSetup) GetStandardInvoiceActual() InvoiceTemplateReference`

GetStandardInvoiceActual returns the StandardInvoiceActual field if non-nil, zero value otherwise.

### GetStandardInvoiceActualOk

`func (o *BillingSetup) GetStandardInvoiceActualOk() (*InvoiceTemplateReference, bool)`

GetStandardInvoiceActualOk returns a tuple with the StandardInvoiceActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardInvoiceActual

`func (o *BillingSetup) SetStandardInvoiceActual(v InvoiceTemplateReference)`

SetStandardInvoiceActual sets StandardInvoiceActual field to given value.

### HasStandardInvoiceActual

`func (o *BillingSetup) HasStandardInvoiceActual() bool`

HasStandardInvoiceActual returns a boolean if a field has been set.

### GetStandardInvoiceFixed

`func (o *BillingSetup) GetStandardInvoiceFixed() InvoiceTemplateReference`

GetStandardInvoiceFixed returns the StandardInvoiceFixed field if non-nil, zero value otherwise.

### GetStandardInvoiceFixedOk

`func (o *BillingSetup) GetStandardInvoiceFixedOk() (*InvoiceTemplateReference, bool)`

GetStandardInvoiceFixedOk returns a tuple with the StandardInvoiceFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardInvoiceFixed

`func (o *BillingSetup) SetStandardInvoiceFixed(v InvoiceTemplateReference)`

SetStandardInvoiceFixed sets StandardInvoiceFixed field to given value.

### HasStandardInvoiceFixed

`func (o *BillingSetup) HasStandardInvoiceFixed() bool`

HasStandardInvoiceFixed returns a boolean if a field has been set.

### GetProgressInvoice

`func (o *BillingSetup) GetProgressInvoice() InvoiceTemplateReference`

GetProgressInvoice returns the ProgressInvoice field if non-nil, zero value otherwise.

### GetProgressInvoiceOk

`func (o *BillingSetup) GetProgressInvoiceOk() (*InvoiceTemplateReference, bool)`

GetProgressInvoiceOk returns a tuple with the ProgressInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressInvoice

`func (o *BillingSetup) SetProgressInvoice(v InvoiceTemplateReference)`

SetProgressInvoice sets ProgressInvoice field to given value.

### HasProgressInvoice

`func (o *BillingSetup) HasProgressInvoice() bool`

HasProgressInvoice returns a boolean if a field has been set.

### GetAgreementInvoice

`func (o *BillingSetup) GetAgreementInvoice() InvoiceTemplateReference`

GetAgreementInvoice returns the AgreementInvoice field if non-nil, zero value otherwise.

### GetAgreementInvoiceOk

`func (o *BillingSetup) GetAgreementInvoiceOk() (*InvoiceTemplateReference, bool)`

GetAgreementInvoiceOk returns a tuple with the AgreementInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementInvoice

`func (o *BillingSetup) SetAgreementInvoice(v InvoiceTemplateReference)`

SetAgreementInvoice sets AgreementInvoice field to given value.

### HasAgreementInvoice

`func (o *BillingSetup) HasAgreementInvoice() bool`

HasAgreementInvoice returns a boolean if a field has been set.

### GetCreditMemoInvoice

`func (o *BillingSetup) GetCreditMemoInvoice() InvoiceTemplateReference`

GetCreditMemoInvoice returns the CreditMemoInvoice field if non-nil, zero value otherwise.

### GetCreditMemoInvoiceOk

`func (o *BillingSetup) GetCreditMemoInvoiceOk() (*InvoiceTemplateReference, bool)`

GetCreditMemoInvoiceOk returns a tuple with the CreditMemoInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditMemoInvoice

`func (o *BillingSetup) SetCreditMemoInvoice(v InvoiceTemplateReference)`

SetCreditMemoInvoice sets CreditMemoInvoice field to given value.

### HasCreditMemoInvoice

`func (o *BillingSetup) HasCreditMemoInvoice() bool`

HasCreditMemoInvoice returns a boolean if a field has been set.

### GetDownPaymentInvoice

`func (o *BillingSetup) GetDownPaymentInvoice() InvoiceTemplateReference`

GetDownPaymentInvoice returns the DownPaymentInvoice field if non-nil, zero value otherwise.

### GetDownPaymentInvoiceOk

`func (o *BillingSetup) GetDownPaymentInvoiceOk() (*InvoiceTemplateReference, bool)`

GetDownPaymentInvoiceOk returns a tuple with the DownPaymentInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownPaymentInvoice

`func (o *BillingSetup) SetDownPaymentInvoice(v InvoiceTemplateReference)`

SetDownPaymentInvoice sets DownPaymentInvoice field to given value.

### HasDownPaymentInvoice

`func (o *BillingSetup) HasDownPaymentInvoice() bool`

HasDownPaymentInvoice returns a boolean if a field has been set.

### GetMiscInvoice

`func (o *BillingSetup) GetMiscInvoice() InvoiceTemplateReference`

GetMiscInvoice returns the MiscInvoice field if non-nil, zero value otherwise.

### GetMiscInvoiceOk

`func (o *BillingSetup) GetMiscInvoiceOk() (*InvoiceTemplateReference, bool)`

GetMiscInvoiceOk returns a tuple with the MiscInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscInvoice

`func (o *BillingSetup) SetMiscInvoice(v InvoiceTemplateReference)`

SetMiscInvoice sets MiscInvoice field to given value.

### HasMiscInvoice

`func (o *BillingSetup) HasMiscInvoice() bool`

HasMiscInvoice returns a boolean if a field has been set.

### GetSalesOrderInvoice

`func (o *BillingSetup) GetSalesOrderInvoice() InvoiceTemplateReference`

GetSalesOrderInvoice returns the SalesOrderInvoice field if non-nil, zero value otherwise.

### GetSalesOrderInvoiceOk

`func (o *BillingSetup) GetSalesOrderInvoiceOk() (*InvoiceTemplateReference, bool)`

GetSalesOrderInvoiceOk returns a tuple with the SalesOrderInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrderInvoice

`func (o *BillingSetup) SetSalesOrderInvoice(v InvoiceTemplateReference)`

SetSalesOrderInvoice sets SalesOrderInvoice field to given value.

### HasSalesOrderInvoice

`func (o *BillingSetup) HasSalesOrderInvoice() bool`

HasSalesOrderInvoice returns a boolean if a field has been set.

### GetExcludeDoNotBillTimeFlag

`func (o *BillingSetup) GetExcludeDoNotBillTimeFlag() bool`

GetExcludeDoNotBillTimeFlag returns the ExcludeDoNotBillTimeFlag field if non-nil, zero value otherwise.

### GetExcludeDoNotBillTimeFlagOk

`func (o *BillingSetup) GetExcludeDoNotBillTimeFlagOk() (*bool, bool)`

GetExcludeDoNotBillTimeFlagOk returns a tuple with the ExcludeDoNotBillTimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDoNotBillTimeFlag

`func (o *BillingSetup) SetExcludeDoNotBillTimeFlag(v bool)`

SetExcludeDoNotBillTimeFlag sets ExcludeDoNotBillTimeFlag field to given value.

### HasExcludeDoNotBillTimeFlag

`func (o *BillingSetup) HasExcludeDoNotBillTimeFlag() bool`

HasExcludeDoNotBillTimeFlag returns a boolean if a field has been set.

### SetExcludeDoNotBillTimeFlagNil

`func (o *BillingSetup) SetExcludeDoNotBillTimeFlagNil(b bool)`

 SetExcludeDoNotBillTimeFlagNil sets the value for ExcludeDoNotBillTimeFlag to be an explicit nil

### UnsetExcludeDoNotBillTimeFlag
`func (o *BillingSetup) UnsetExcludeDoNotBillTimeFlag()`

UnsetExcludeDoNotBillTimeFlag ensures that no value is present for ExcludeDoNotBillTimeFlag, not even an explicit nil
### GetExcludeDoNotBillExpenseFlag

`func (o *BillingSetup) GetExcludeDoNotBillExpenseFlag() bool`

GetExcludeDoNotBillExpenseFlag returns the ExcludeDoNotBillExpenseFlag field if non-nil, zero value otherwise.

### GetExcludeDoNotBillExpenseFlagOk

`func (o *BillingSetup) GetExcludeDoNotBillExpenseFlagOk() (*bool, bool)`

GetExcludeDoNotBillExpenseFlagOk returns a tuple with the ExcludeDoNotBillExpenseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDoNotBillExpenseFlag

`func (o *BillingSetup) SetExcludeDoNotBillExpenseFlag(v bool)`

SetExcludeDoNotBillExpenseFlag sets ExcludeDoNotBillExpenseFlag field to given value.

### HasExcludeDoNotBillExpenseFlag

`func (o *BillingSetup) HasExcludeDoNotBillExpenseFlag() bool`

HasExcludeDoNotBillExpenseFlag returns a boolean if a field has been set.

### SetExcludeDoNotBillExpenseFlagNil

`func (o *BillingSetup) SetExcludeDoNotBillExpenseFlagNil(b bool)`

 SetExcludeDoNotBillExpenseFlagNil sets the value for ExcludeDoNotBillExpenseFlag to be an explicit nil

### UnsetExcludeDoNotBillExpenseFlag
`func (o *BillingSetup) UnsetExcludeDoNotBillExpenseFlag()`

UnsetExcludeDoNotBillExpenseFlag ensures that no value is present for ExcludeDoNotBillExpenseFlag, not even an explicit nil
### GetExcludeDoNotBillProductFlag

`func (o *BillingSetup) GetExcludeDoNotBillProductFlag() bool`

GetExcludeDoNotBillProductFlag returns the ExcludeDoNotBillProductFlag field if non-nil, zero value otherwise.

### GetExcludeDoNotBillProductFlagOk

`func (o *BillingSetup) GetExcludeDoNotBillProductFlagOk() (*bool, bool)`

GetExcludeDoNotBillProductFlagOk returns a tuple with the ExcludeDoNotBillProductFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDoNotBillProductFlag

`func (o *BillingSetup) SetExcludeDoNotBillProductFlag(v bool)`

SetExcludeDoNotBillProductFlag sets ExcludeDoNotBillProductFlag field to given value.

### HasExcludeDoNotBillProductFlag

`func (o *BillingSetup) HasExcludeDoNotBillProductFlag() bool`

HasExcludeDoNotBillProductFlag returns a boolean if a field has been set.

### SetExcludeDoNotBillProductFlagNil

`func (o *BillingSetup) SetExcludeDoNotBillProductFlagNil(b bool)`

 SetExcludeDoNotBillProductFlagNil sets the value for ExcludeDoNotBillProductFlag to be an explicit nil

### UnsetExcludeDoNotBillProductFlag
`func (o *BillingSetup) UnsetExcludeDoNotBillProductFlag()`

UnsetExcludeDoNotBillProductFlag ensures that no value is present for ExcludeDoNotBillProductFlag, not even an explicit nil
### GetPrefixSuffixFlag

`func (o *BillingSetup) GetPrefixSuffixFlag() string`

GetPrefixSuffixFlag returns the PrefixSuffixFlag field if non-nil, zero value otherwise.

### GetPrefixSuffixFlagOk

`func (o *BillingSetup) GetPrefixSuffixFlagOk() (*string, bool)`

GetPrefixSuffixFlagOk returns a tuple with the PrefixSuffixFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSuffixFlag

`func (o *BillingSetup) SetPrefixSuffixFlag(v string)`

SetPrefixSuffixFlag sets PrefixSuffixFlag field to given value.

### HasPrefixSuffixFlag

`func (o *BillingSetup) HasPrefixSuffixFlag() bool`

HasPrefixSuffixFlag returns a boolean if a field has been set.

### SetPrefixSuffixFlagNil

`func (o *BillingSetup) SetPrefixSuffixFlagNil(b bool)`

 SetPrefixSuffixFlagNil sets the value for PrefixSuffixFlag to be an explicit nil

### UnsetPrefixSuffixFlag
`func (o *BillingSetup) UnsetPrefixSuffixFlag()`

UnsetPrefixSuffixFlag ensures that no value is present for PrefixSuffixFlag, not even an explicit nil
### GetPrefixSuffixText

`func (o *BillingSetup) GetPrefixSuffixText() string`

GetPrefixSuffixText returns the PrefixSuffixText field if non-nil, zero value otherwise.

### GetPrefixSuffixTextOk

`func (o *BillingSetup) GetPrefixSuffixTextOk() (*string, bool)`

GetPrefixSuffixTextOk returns a tuple with the PrefixSuffixText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSuffixText

`func (o *BillingSetup) SetPrefixSuffixText(v string)`

SetPrefixSuffixText sets PrefixSuffixText field to given value.

### HasPrefixSuffixText

`func (o *BillingSetup) HasPrefixSuffixText() bool`

HasPrefixSuffixText returns a boolean if a field has been set.

### GetChargeAdjToFirmFlag

`func (o *BillingSetup) GetChargeAdjToFirmFlag() bool`

GetChargeAdjToFirmFlag returns the ChargeAdjToFirmFlag field if non-nil, zero value otherwise.

### GetChargeAdjToFirmFlagOk

`func (o *BillingSetup) GetChargeAdjToFirmFlagOk() (*bool, bool)`

GetChargeAdjToFirmFlagOk returns a tuple with the ChargeAdjToFirmFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAdjToFirmFlag

`func (o *BillingSetup) SetChargeAdjToFirmFlag(v bool)`

SetChargeAdjToFirmFlag sets ChargeAdjToFirmFlag field to given value.

### HasChargeAdjToFirmFlag

`func (o *BillingSetup) HasChargeAdjToFirmFlag() bool`

HasChargeAdjToFirmFlag returns a boolean if a field has been set.

### SetChargeAdjToFirmFlagNil

`func (o *BillingSetup) SetChargeAdjToFirmFlagNil(b bool)`

 SetChargeAdjToFirmFlagNil sets the value for ChargeAdjToFirmFlag to be an explicit nil

### UnsetChargeAdjToFirmFlag
`func (o *BillingSetup) UnsetChargeAdjToFirmFlag()`

UnsetChargeAdjToFirmFlag ensures that no value is present for ChargeAdjToFirmFlag, not even an explicit nil
### GetNoWatermarkFlag

`func (o *BillingSetup) GetNoWatermarkFlag() bool`

GetNoWatermarkFlag returns the NoWatermarkFlag field if non-nil, zero value otherwise.

### GetNoWatermarkFlagOk

`func (o *BillingSetup) GetNoWatermarkFlagOk() (*bool, bool)`

GetNoWatermarkFlagOk returns a tuple with the NoWatermarkFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoWatermarkFlag

`func (o *BillingSetup) SetNoWatermarkFlag(v bool)`

SetNoWatermarkFlag sets NoWatermarkFlag field to given value.

### HasNoWatermarkFlag

`func (o *BillingSetup) HasNoWatermarkFlag() bool`

HasNoWatermarkFlag returns a boolean if a field has been set.

### SetNoWatermarkFlagNil

`func (o *BillingSetup) SetNoWatermarkFlagNil(b bool)`

 SetNoWatermarkFlagNil sets the value for NoWatermarkFlag to be an explicit nil

### UnsetNoWatermarkFlag
`func (o *BillingSetup) UnsetNoWatermarkFlag()`

UnsetNoWatermarkFlag ensures that no value is present for NoWatermarkFlag, not even an explicit nil
### GetDisplayTaxFlag

`func (o *BillingSetup) GetDisplayTaxFlag() bool`

GetDisplayTaxFlag returns the DisplayTaxFlag field if non-nil, zero value otherwise.

### GetDisplayTaxFlagOk

`func (o *BillingSetup) GetDisplayTaxFlagOk() (*bool, bool)`

GetDisplayTaxFlagOk returns a tuple with the DisplayTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTaxFlag

`func (o *BillingSetup) SetDisplayTaxFlag(v bool)`

SetDisplayTaxFlag sets DisplayTaxFlag field to given value.

### HasDisplayTaxFlag

`func (o *BillingSetup) HasDisplayTaxFlag() bool`

HasDisplayTaxFlag returns a boolean if a field has been set.

### SetDisplayTaxFlagNil

`func (o *BillingSetup) SetDisplayTaxFlagNil(b bool)`

 SetDisplayTaxFlagNil sets the value for DisplayTaxFlag to be an explicit nil

### UnsetDisplayTaxFlag
`func (o *BillingSetup) UnsetDisplayTaxFlag()`

UnsetDisplayTaxFlag ensures that no value is present for DisplayTaxFlag, not even an explicit nil
### GetAllowRestrictedDeptOnRoutingFlag

`func (o *BillingSetup) GetAllowRestrictedDeptOnRoutingFlag() bool`

GetAllowRestrictedDeptOnRoutingFlag returns the AllowRestrictedDeptOnRoutingFlag field if non-nil, zero value otherwise.

### GetAllowRestrictedDeptOnRoutingFlagOk

`func (o *BillingSetup) GetAllowRestrictedDeptOnRoutingFlagOk() (*bool, bool)`

GetAllowRestrictedDeptOnRoutingFlagOk returns a tuple with the AllowRestrictedDeptOnRoutingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRestrictedDeptOnRoutingFlag

`func (o *BillingSetup) SetAllowRestrictedDeptOnRoutingFlag(v bool)`

SetAllowRestrictedDeptOnRoutingFlag sets AllowRestrictedDeptOnRoutingFlag field to given value.

### HasAllowRestrictedDeptOnRoutingFlag

`func (o *BillingSetup) HasAllowRestrictedDeptOnRoutingFlag() bool`

HasAllowRestrictedDeptOnRoutingFlag returns a boolean if a field has been set.

### SetAllowRestrictedDeptOnRoutingFlagNil

`func (o *BillingSetup) SetAllowRestrictedDeptOnRoutingFlagNil(b bool)`

 SetAllowRestrictedDeptOnRoutingFlagNil sets the value for AllowRestrictedDeptOnRoutingFlag to be an explicit nil

### UnsetAllowRestrictedDeptOnRoutingFlag
`func (o *BillingSetup) UnsetAllowRestrictedDeptOnRoutingFlag()`

UnsetAllowRestrictedDeptOnRoutingFlag ensures that no value is present for AllowRestrictedDeptOnRoutingFlag, not even an explicit nil
### GetBillTicketSeparatelyFlag

`func (o *BillingSetup) GetBillTicketSeparatelyFlag() bool`

GetBillTicketSeparatelyFlag returns the BillTicketSeparatelyFlag field if non-nil, zero value otherwise.

### GetBillTicketSeparatelyFlagOk

`func (o *BillingSetup) GetBillTicketSeparatelyFlagOk() (*bool, bool)`

GetBillTicketSeparatelyFlagOk returns a tuple with the BillTicketSeparatelyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTicketSeparatelyFlag

`func (o *BillingSetup) SetBillTicketSeparatelyFlag(v bool)`

SetBillTicketSeparatelyFlag sets BillTicketSeparatelyFlag field to given value.

### HasBillTicketSeparatelyFlag

`func (o *BillingSetup) HasBillTicketSeparatelyFlag() bool`

HasBillTicketSeparatelyFlag returns a boolean if a field has been set.

### SetBillTicketSeparatelyFlagNil

`func (o *BillingSetup) SetBillTicketSeparatelyFlagNil(b bool)`

 SetBillTicketSeparatelyFlagNil sets the value for BillTicketSeparatelyFlag to be an explicit nil

### UnsetBillTicketSeparatelyFlag
`func (o *BillingSetup) UnsetBillTicketSeparatelyFlag()`

UnsetBillTicketSeparatelyFlag ensures that no value is present for BillTicketSeparatelyFlag, not even an explicit nil
### GetBillTicketCompleteFlag

`func (o *BillingSetup) GetBillTicketCompleteFlag() bool`

GetBillTicketCompleteFlag returns the BillTicketCompleteFlag field if non-nil, zero value otherwise.

### GetBillTicketCompleteFlagOk

`func (o *BillingSetup) GetBillTicketCompleteFlagOk() (*bool, bool)`

GetBillTicketCompleteFlagOk returns a tuple with the BillTicketCompleteFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTicketCompleteFlag

`func (o *BillingSetup) SetBillTicketCompleteFlag(v bool)`

SetBillTicketCompleteFlag sets BillTicketCompleteFlag field to given value.

### HasBillTicketCompleteFlag

`func (o *BillingSetup) HasBillTicketCompleteFlag() bool`

HasBillTicketCompleteFlag returns a boolean if a field has been set.

### SetBillTicketCompleteFlagNil

`func (o *BillingSetup) SetBillTicketCompleteFlagNil(b bool)`

 SetBillTicketCompleteFlagNil sets the value for BillTicketCompleteFlag to be an explicit nil

### UnsetBillTicketCompleteFlag
`func (o *BillingSetup) UnsetBillTicketCompleteFlag()`

UnsetBillTicketCompleteFlag ensures that no value is present for BillTicketCompleteFlag, not even an explicit nil
### GetBillTicketUnapprovedFlag

`func (o *BillingSetup) GetBillTicketUnapprovedFlag() bool`

GetBillTicketUnapprovedFlag returns the BillTicketUnapprovedFlag field if non-nil, zero value otherwise.

### GetBillTicketUnapprovedFlagOk

`func (o *BillingSetup) GetBillTicketUnapprovedFlagOk() (*bool, bool)`

GetBillTicketUnapprovedFlagOk returns a tuple with the BillTicketUnapprovedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTicketUnapprovedFlag

`func (o *BillingSetup) SetBillTicketUnapprovedFlag(v bool)`

SetBillTicketUnapprovedFlag sets BillTicketUnapprovedFlag field to given value.

### HasBillTicketUnapprovedFlag

`func (o *BillingSetup) HasBillTicketUnapprovedFlag() bool`

HasBillTicketUnapprovedFlag returns a boolean if a field has been set.

### SetBillTicketUnapprovedFlagNil

`func (o *BillingSetup) SetBillTicketUnapprovedFlagNil(b bool)`

 SetBillTicketUnapprovedFlagNil sets the value for BillTicketUnapprovedFlag to be an explicit nil

### UnsetBillTicketUnapprovedFlag
`func (o *BillingSetup) UnsetBillTicketUnapprovedFlag()`

UnsetBillTicketUnapprovedFlag ensures that no value is present for BillTicketUnapprovedFlag, not even an explicit nil
### GetBillProjectCompleteFlag

`func (o *BillingSetup) GetBillProjectCompleteFlag() bool`

GetBillProjectCompleteFlag returns the BillProjectCompleteFlag field if non-nil, zero value otherwise.

### GetBillProjectCompleteFlagOk

`func (o *BillingSetup) GetBillProjectCompleteFlagOk() (*bool, bool)`

GetBillProjectCompleteFlagOk returns a tuple with the BillProjectCompleteFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProjectCompleteFlag

`func (o *BillingSetup) SetBillProjectCompleteFlag(v bool)`

SetBillProjectCompleteFlag sets BillProjectCompleteFlag field to given value.

### HasBillProjectCompleteFlag

`func (o *BillingSetup) HasBillProjectCompleteFlag() bool`

HasBillProjectCompleteFlag returns a boolean if a field has been set.

### SetBillProjectCompleteFlagNil

`func (o *BillingSetup) SetBillProjectCompleteFlagNil(b bool)`

 SetBillProjectCompleteFlagNil sets the value for BillProjectCompleteFlag to be an explicit nil

### UnsetBillProjectCompleteFlag
`func (o *BillingSetup) UnsetBillProjectCompleteFlag()`

UnsetBillProjectCompleteFlag ensures that no value is present for BillProjectCompleteFlag, not even an explicit nil
### GetBillProjectUnapprovedFlag

`func (o *BillingSetup) GetBillProjectUnapprovedFlag() bool`

GetBillProjectUnapprovedFlag returns the BillProjectUnapprovedFlag field if non-nil, zero value otherwise.

### GetBillProjectUnapprovedFlagOk

`func (o *BillingSetup) GetBillProjectUnapprovedFlagOk() (*bool, bool)`

GetBillProjectUnapprovedFlagOk returns a tuple with the BillProjectUnapprovedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProjectUnapprovedFlag

`func (o *BillingSetup) SetBillProjectUnapprovedFlag(v bool)`

SetBillProjectUnapprovedFlag sets BillProjectUnapprovedFlag field to given value.

### HasBillProjectUnapprovedFlag

`func (o *BillingSetup) HasBillProjectUnapprovedFlag() bool`

HasBillProjectUnapprovedFlag returns a boolean if a field has been set.

### SetBillProjectUnapprovedFlagNil

`func (o *BillingSetup) SetBillProjectUnapprovedFlagNil(b bool)`

 SetBillProjectUnapprovedFlagNil sets the value for BillProjectUnapprovedFlag to be an explicit nil

### UnsetBillProjectUnapprovedFlag
`func (o *BillingSetup) UnsetBillProjectUnapprovedFlag()`

UnsetBillProjectUnapprovedFlag ensures that no value is present for BillProjectUnapprovedFlag, not even an explicit nil
### GetProgressTimeFlag

`func (o *BillingSetup) GetProgressTimeFlag() bool`

GetProgressTimeFlag returns the ProgressTimeFlag field if non-nil, zero value otherwise.

### GetProgressTimeFlagOk

`func (o *BillingSetup) GetProgressTimeFlagOk() (*bool, bool)`

GetProgressTimeFlagOk returns a tuple with the ProgressTimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTimeFlag

`func (o *BillingSetup) SetProgressTimeFlag(v bool)`

SetProgressTimeFlag sets ProgressTimeFlag field to given value.

### HasProgressTimeFlag

`func (o *BillingSetup) HasProgressTimeFlag() bool`

HasProgressTimeFlag returns a boolean if a field has been set.

### SetProgressTimeFlagNil

`func (o *BillingSetup) SetProgressTimeFlagNil(b bool)`

 SetProgressTimeFlagNil sets the value for ProgressTimeFlag to be an explicit nil

### UnsetProgressTimeFlag
`func (o *BillingSetup) UnsetProgressTimeFlag()`

UnsetProgressTimeFlag ensures that no value is present for ProgressTimeFlag, not even an explicit nil
### GetRestrictProjectDownpaymentFlag

`func (o *BillingSetup) GetRestrictProjectDownpaymentFlag() bool`

GetRestrictProjectDownpaymentFlag returns the RestrictProjectDownpaymentFlag field if non-nil, zero value otherwise.

### GetRestrictProjectDownpaymentFlagOk

`func (o *BillingSetup) GetRestrictProjectDownpaymentFlagOk() (*bool, bool)`

GetRestrictProjectDownpaymentFlagOk returns a tuple with the RestrictProjectDownpaymentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictProjectDownpaymentFlag

`func (o *BillingSetup) SetRestrictProjectDownpaymentFlag(v bool)`

SetRestrictProjectDownpaymentFlag sets RestrictProjectDownpaymentFlag field to given value.

### HasRestrictProjectDownpaymentFlag

`func (o *BillingSetup) HasRestrictProjectDownpaymentFlag() bool`

HasRestrictProjectDownpaymentFlag returns a boolean if a field has been set.

### SetRestrictProjectDownpaymentFlagNil

`func (o *BillingSetup) SetRestrictProjectDownpaymentFlagNil(b bool)`

 SetRestrictProjectDownpaymentFlagNil sets the value for RestrictProjectDownpaymentFlag to be an explicit nil

### UnsetRestrictProjectDownpaymentFlag
`func (o *BillingSetup) UnsetRestrictProjectDownpaymentFlag()`

UnsetRestrictProjectDownpaymentFlag ensures that no value is present for RestrictProjectDownpaymentFlag, not even an explicit nil
### GetBillSalesOrderCompleteFlag

`func (o *BillingSetup) GetBillSalesOrderCompleteFlag() bool`

GetBillSalesOrderCompleteFlag returns the BillSalesOrderCompleteFlag field if non-nil, zero value otherwise.

### GetBillSalesOrderCompleteFlagOk

`func (o *BillingSetup) GetBillSalesOrderCompleteFlagOk() (*bool, bool)`

GetBillSalesOrderCompleteFlagOk returns a tuple with the BillSalesOrderCompleteFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillSalesOrderCompleteFlag

`func (o *BillingSetup) SetBillSalesOrderCompleteFlag(v bool)`

SetBillSalesOrderCompleteFlag sets BillSalesOrderCompleteFlag field to given value.

### HasBillSalesOrderCompleteFlag

`func (o *BillingSetup) HasBillSalesOrderCompleteFlag() bool`

HasBillSalesOrderCompleteFlag returns a boolean if a field has been set.

### SetBillSalesOrderCompleteFlagNil

`func (o *BillingSetup) SetBillSalesOrderCompleteFlagNil(b bool)`

 SetBillSalesOrderCompleteFlagNil sets the value for BillSalesOrderCompleteFlag to be an explicit nil

### UnsetBillSalesOrderCompleteFlag
`func (o *BillingSetup) UnsetBillSalesOrderCompleteFlag()`

UnsetBillSalesOrderCompleteFlag ensures that no value is present for BillSalesOrderCompleteFlag, not even an explicit nil
### GetBillProductAfterShipFlag

`func (o *BillingSetup) GetBillProductAfterShipFlag() bool`

GetBillProductAfterShipFlag returns the BillProductAfterShipFlag field if non-nil, zero value otherwise.

### GetBillProductAfterShipFlagOk

`func (o *BillingSetup) GetBillProductAfterShipFlagOk() (*bool, bool)`

GetBillProductAfterShipFlagOk returns a tuple with the BillProductAfterShipFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProductAfterShipFlag

`func (o *BillingSetup) SetBillProductAfterShipFlag(v bool)`

SetBillProductAfterShipFlag sets BillProductAfterShipFlag field to given value.

### HasBillProductAfterShipFlag

`func (o *BillingSetup) HasBillProductAfterShipFlag() bool`

HasBillProductAfterShipFlag returns a boolean if a field has been set.

### SetBillProductAfterShipFlagNil

`func (o *BillingSetup) SetBillProductAfterShipFlagNil(b bool)`

 SetBillProductAfterShipFlagNil sets the value for BillProductAfterShipFlag to be an explicit nil

### UnsetBillProductAfterShipFlag
`func (o *BillingSetup) UnsetBillProductAfterShipFlag()`

UnsetBillProductAfterShipFlag ensures that no value is present for BillProductAfterShipFlag, not even an explicit nil
### GetRestrictDownpaymentFlag

`func (o *BillingSetup) GetRestrictDownpaymentFlag() bool`

GetRestrictDownpaymentFlag returns the RestrictDownpaymentFlag field if non-nil, zero value otherwise.

### GetRestrictDownpaymentFlagOk

`func (o *BillingSetup) GetRestrictDownpaymentFlagOk() (*bool, bool)`

GetRestrictDownpaymentFlagOk returns a tuple with the RestrictDownpaymentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDownpaymentFlag

`func (o *BillingSetup) SetRestrictDownpaymentFlag(v bool)`

SetRestrictDownpaymentFlag sets RestrictDownpaymentFlag field to given value.

### HasRestrictDownpaymentFlag

`func (o *BillingSetup) HasRestrictDownpaymentFlag() bool`

HasRestrictDownpaymentFlag returns a boolean if a field has been set.

### SetRestrictDownpaymentFlagNil

`func (o *BillingSetup) SetRestrictDownpaymentFlagNil(b bool)`

 SetRestrictDownpaymentFlagNil sets the value for RestrictDownpaymentFlag to be an explicit nil

### UnsetRestrictDownpaymentFlag
`func (o *BillingSetup) UnsetRestrictDownpaymentFlag()`

UnsetRestrictDownpaymentFlag ensures that no value is present for RestrictDownpaymentFlag, not even an explicit nil
### GetCopyNonServiceProductsFlag

`func (o *BillingSetup) GetCopyNonServiceProductsFlag() bool`

GetCopyNonServiceProductsFlag returns the CopyNonServiceProductsFlag field if non-nil, zero value otherwise.

### GetCopyNonServiceProductsFlagOk

`func (o *BillingSetup) GetCopyNonServiceProductsFlagOk() (*bool, bool)`

GetCopyNonServiceProductsFlagOk returns a tuple with the CopyNonServiceProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyNonServiceProductsFlag

`func (o *BillingSetup) SetCopyNonServiceProductsFlag(v bool)`

SetCopyNonServiceProductsFlag sets CopyNonServiceProductsFlag field to given value.

### HasCopyNonServiceProductsFlag

`func (o *BillingSetup) HasCopyNonServiceProductsFlag() bool`

HasCopyNonServiceProductsFlag returns a boolean if a field has been set.

### SetCopyNonServiceProductsFlagNil

`func (o *BillingSetup) SetCopyNonServiceProductsFlagNil(b bool)`

 SetCopyNonServiceProductsFlagNil sets the value for CopyNonServiceProductsFlag to be an explicit nil

### UnsetCopyNonServiceProductsFlag
`func (o *BillingSetup) UnsetCopyNonServiceProductsFlag()`

UnsetCopyNonServiceProductsFlag ensures that no value is present for CopyNonServiceProductsFlag, not even an explicit nil
### GetCopyServiceProductsFlag

`func (o *BillingSetup) GetCopyServiceProductsFlag() bool`

GetCopyServiceProductsFlag returns the CopyServiceProductsFlag field if non-nil, zero value otherwise.

### GetCopyServiceProductsFlagOk

`func (o *BillingSetup) GetCopyServiceProductsFlagOk() (*bool, bool)`

GetCopyServiceProductsFlagOk returns a tuple with the CopyServiceProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyServiceProductsFlag

`func (o *BillingSetup) SetCopyServiceProductsFlag(v bool)`

SetCopyServiceProductsFlag sets CopyServiceProductsFlag field to given value.

### HasCopyServiceProductsFlag

`func (o *BillingSetup) HasCopyServiceProductsFlag() bool`

HasCopyServiceProductsFlag returns a boolean if a field has been set.

### SetCopyServiceProductsFlagNil

`func (o *BillingSetup) SetCopyServiceProductsFlagNil(b bool)`

 SetCopyServiceProductsFlagNil sets the value for CopyServiceProductsFlag to be an explicit nil

### UnsetCopyServiceProductsFlag
`func (o *BillingSetup) UnsetCopyServiceProductsFlag()`

UnsetCopyServiceProductsFlag ensures that no value is present for CopyServiceProductsFlag, not even an explicit nil
### GetCopyAgreementProductsFlag

`func (o *BillingSetup) GetCopyAgreementProductsFlag() bool`

GetCopyAgreementProductsFlag returns the CopyAgreementProductsFlag field if non-nil, zero value otherwise.

### GetCopyAgreementProductsFlagOk

`func (o *BillingSetup) GetCopyAgreementProductsFlagOk() (*bool, bool)`

GetCopyAgreementProductsFlagOk returns a tuple with the CopyAgreementProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyAgreementProductsFlag

`func (o *BillingSetup) SetCopyAgreementProductsFlag(v bool)`

SetCopyAgreementProductsFlag sets CopyAgreementProductsFlag field to given value.

### HasCopyAgreementProductsFlag

`func (o *BillingSetup) HasCopyAgreementProductsFlag() bool`

HasCopyAgreementProductsFlag returns a boolean if a field has been set.

### SetCopyAgreementProductsFlagNil

`func (o *BillingSetup) SetCopyAgreementProductsFlagNil(b bool)`

 SetCopyAgreementProductsFlagNil sets the value for CopyAgreementProductsFlag to be an explicit nil

### UnsetCopyAgreementProductsFlag
`func (o *BillingSetup) UnsetCopyAgreementProductsFlag()`

UnsetCopyAgreementProductsFlag ensures that no value is present for CopyAgreementProductsFlag, not even an explicit nil
### GetPrintLogoFlag

`func (o *BillingSetup) GetPrintLogoFlag() bool`

GetPrintLogoFlag returns the PrintLogoFlag field if non-nil, zero value otherwise.

### GetPrintLogoFlagOk

`func (o *BillingSetup) GetPrintLogoFlagOk() (*bool, bool)`

GetPrintLogoFlagOk returns a tuple with the PrintLogoFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintLogoFlag

`func (o *BillingSetup) SetPrintLogoFlag(v bool)`

SetPrintLogoFlag sets PrintLogoFlag field to given value.

### HasPrintLogoFlag

`func (o *BillingSetup) HasPrintLogoFlag() bool`

HasPrintLogoFlag returns a boolean if a field has been set.

### SetPrintLogoFlagNil

`func (o *BillingSetup) SetPrintLogoFlagNil(b bool)`

 SetPrintLogoFlagNil sets the value for PrintLogoFlag to be an explicit nil

### UnsetPrintLogoFlag
`func (o *BillingSetup) UnsetPrintLogoFlag()`

UnsetPrintLogoFlag ensures that no value is present for PrintLogoFlag, not even an explicit nil
### GetReadReceiptFlag

`func (o *BillingSetup) GetReadReceiptFlag() bool`

GetReadReceiptFlag returns the ReadReceiptFlag field if non-nil, zero value otherwise.

### GetReadReceiptFlagOk

`func (o *BillingSetup) GetReadReceiptFlagOk() (*bool, bool)`

GetReadReceiptFlagOk returns a tuple with the ReadReceiptFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadReceiptFlag

`func (o *BillingSetup) SetReadReceiptFlag(v bool)`

SetReadReceiptFlag sets ReadReceiptFlag field to given value.

### HasReadReceiptFlag

`func (o *BillingSetup) HasReadReceiptFlag() bool`

HasReadReceiptFlag returns a boolean if a field has been set.

### SetReadReceiptFlagNil

`func (o *BillingSetup) SetReadReceiptFlagNil(b bool)`

 SetReadReceiptFlagNil sets the value for ReadReceiptFlag to be an explicit nil

### UnsetReadReceiptFlag
`func (o *BillingSetup) UnsetReadReceiptFlag()`

UnsetReadReceiptFlag ensures that no value is present for ReadReceiptFlag, not even an explicit nil
### GetDeliveryReceiptFlag

`func (o *BillingSetup) GetDeliveryReceiptFlag() bool`

GetDeliveryReceiptFlag returns the DeliveryReceiptFlag field if non-nil, zero value otherwise.

### GetDeliveryReceiptFlagOk

`func (o *BillingSetup) GetDeliveryReceiptFlagOk() (*bool, bool)`

GetDeliveryReceiptFlagOk returns a tuple with the DeliveryReceiptFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryReceiptFlag

`func (o *BillingSetup) SetDeliveryReceiptFlag(v bool)`

SetDeliveryReceiptFlag sets DeliveryReceiptFlag field to given value.

### HasDeliveryReceiptFlag

`func (o *BillingSetup) HasDeliveryReceiptFlag() bool`

HasDeliveryReceiptFlag returns a boolean if a field has been set.

### SetDeliveryReceiptFlagNil

`func (o *BillingSetup) SetDeliveryReceiptFlagNil(b bool)`

 SetDeliveryReceiptFlagNil sets the value for DeliveryReceiptFlag to be an explicit nil

### UnsetDeliveryReceiptFlag
`func (o *BillingSetup) UnsetDeliveryReceiptFlag()`

UnsetDeliveryReceiptFlag ensures that no value is present for DeliveryReceiptFlag, not even an explicit nil
### GetAttachXmlInvoiceFlag

`func (o *BillingSetup) GetAttachXmlInvoiceFlag() bool`

GetAttachXmlInvoiceFlag returns the AttachXmlInvoiceFlag field if non-nil, zero value otherwise.

### GetAttachXmlInvoiceFlagOk

`func (o *BillingSetup) GetAttachXmlInvoiceFlagOk() (*bool, bool)`

GetAttachXmlInvoiceFlagOk returns a tuple with the AttachXmlInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachXmlInvoiceFlag

`func (o *BillingSetup) SetAttachXmlInvoiceFlag(v bool)`

SetAttachXmlInvoiceFlag sets AttachXmlInvoiceFlag field to given value.

### HasAttachXmlInvoiceFlag

`func (o *BillingSetup) HasAttachXmlInvoiceFlag() bool`

HasAttachXmlInvoiceFlag returns a boolean if a field has been set.

### SetAttachXmlInvoiceFlagNil

`func (o *BillingSetup) SetAttachXmlInvoiceFlagNil(b bool)`

 SetAttachXmlInvoiceFlagNil sets the value for AttachXmlInvoiceFlag to be an explicit nil

### UnsetAttachXmlInvoiceFlag
`func (o *BillingSetup) UnsetAttachXmlInvoiceFlag()`

UnsetAttachXmlInvoiceFlag ensures that no value is present for AttachXmlInvoiceFlag, not even an explicit nil
### GetDisableRoutingEmailFlag

`func (o *BillingSetup) GetDisableRoutingEmailFlag() bool`

GetDisableRoutingEmailFlag returns the DisableRoutingEmailFlag field if non-nil, zero value otherwise.

### GetDisableRoutingEmailFlagOk

`func (o *BillingSetup) GetDisableRoutingEmailFlagOk() (*bool, bool)`

GetDisableRoutingEmailFlagOk returns a tuple with the DisableRoutingEmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRoutingEmailFlag

`func (o *BillingSetup) SetDisableRoutingEmailFlag(v bool)`

SetDisableRoutingEmailFlag sets DisableRoutingEmailFlag field to given value.

### HasDisableRoutingEmailFlag

`func (o *BillingSetup) HasDisableRoutingEmailFlag() bool`

HasDisableRoutingEmailFlag returns a boolean if a field has been set.

### SetDisableRoutingEmailFlagNil

`func (o *BillingSetup) SetDisableRoutingEmailFlagNil(b bool)`

 SetDisableRoutingEmailFlagNil sets the value for DisableRoutingEmailFlag to be an explicit nil

### UnsetDisableRoutingEmailFlag
`func (o *BillingSetup) UnsetDisableRoutingEmailFlag()`

UnsetDisableRoutingEmailFlag ensures that no value is present for DisableRoutingEmailFlag, not even an explicit nil
### GetEmailTemplate

`func (o *BillingSetup) GetEmailTemplate() EmailTemplateReference`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *BillingSetup) GetEmailTemplateOk() (*EmailTemplateReference, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *BillingSetup) SetEmailTemplate(v EmailTemplateReference)`

SetEmailTemplate sets EmailTemplate field to given value.


### GetLocalizedCountry

`func (o *BillingSetup) GetLocalizedCountry() CountryReference`

GetLocalizedCountry returns the LocalizedCountry field if non-nil, zero value otherwise.

### GetLocalizedCountryOk

`func (o *BillingSetup) GetLocalizedCountryOk() (*CountryReference, bool)`

GetLocalizedCountryOk returns a tuple with the LocalizedCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedCountry

`func (o *BillingSetup) SetLocalizedCountry(v CountryReference)`

SetLocalizedCountry sets LocalizedCountry field to given value.

### HasLocalizedCountry

`func (o *BillingSetup) HasLocalizedCountry() bool`

HasLocalizedCountry returns a boolean if a field has been set.

### GetBusinessNumber

`func (o *BillingSetup) GetBusinessNumber() string`

GetBusinessNumber returns the BusinessNumber field if non-nil, zero value otherwise.

### GetBusinessNumberOk

`func (o *BillingSetup) GetBusinessNumberOk() (*string, bool)`

GetBusinessNumberOk returns a tuple with the BusinessNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessNumber

`func (o *BillingSetup) SetBusinessNumber(v string)`

SetBusinessNumber sets BusinessNumber field to given value.

### HasBusinessNumber

`func (o *BillingSetup) HasBusinessNumber() bool`

HasBusinessNumber returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingSetup) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingSetup) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingSetup) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingSetup) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomLabel

`func (o *BillingSetup) GetCustomLabel() string`

GetCustomLabel returns the CustomLabel field if non-nil, zero value otherwise.

### GetCustomLabelOk

`func (o *BillingSetup) GetCustomLabelOk() (*string, bool)`

GetCustomLabelOk returns a tuple with the CustomLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabel

`func (o *BillingSetup) SetCustomLabel(v string)`

SetCustomLabel sets CustomLabel field to given value.

### HasCustomLabel

`func (o *BillingSetup) HasCustomLabel() bool`

HasCustomLabel returns a boolean if a field has been set.

### GetCustomText

`func (o *BillingSetup) GetCustomText() string`

GetCustomText returns the CustomText field if non-nil, zero value otherwise.

### GetCustomTextOk

`func (o *BillingSetup) GetCustomTextOk() (*string, bool)`

GetCustomTextOk returns a tuple with the CustomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText

`func (o *BillingSetup) SetCustomText(v string)`

SetCustomText sets CustomText field to given value.

### HasCustomText

`func (o *BillingSetup) HasCustomText() bool`

HasCustomText returns a boolean if a field has been set.

### GetCompanyCode

`func (o *BillingSetup) GetCompanyCode() string`

GetCompanyCode returns the CompanyCode field if non-nil, zero value otherwise.

### GetCompanyCodeOk

`func (o *BillingSetup) GetCompanyCodeOk() (*string, bool)`

GetCompanyCodeOk returns a tuple with the CompanyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCode

`func (o *BillingSetup) SetCompanyCode(v string)`

SetCompanyCode sets CompanyCode field to given value.

### HasCompanyCode

`func (o *BillingSetup) HasCompanyCode() bool`

HasCompanyCode returns a boolean if a field has been set.

### GetExcludeAvalaraFlag

`func (o *BillingSetup) GetExcludeAvalaraFlag() bool`

GetExcludeAvalaraFlag returns the ExcludeAvalaraFlag field if non-nil, zero value otherwise.

### GetExcludeAvalaraFlagOk

`func (o *BillingSetup) GetExcludeAvalaraFlagOk() (*bool, bool)`

GetExcludeAvalaraFlagOk returns a tuple with the ExcludeAvalaraFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAvalaraFlag

`func (o *BillingSetup) SetExcludeAvalaraFlag(v bool)`

SetExcludeAvalaraFlag sets ExcludeAvalaraFlag field to given value.

### HasExcludeAvalaraFlag

`func (o *BillingSetup) HasExcludeAvalaraFlag() bool`

HasExcludeAvalaraFlag returns a boolean if a field has been set.

### SetExcludeAvalaraFlagNil

`func (o *BillingSetup) SetExcludeAvalaraFlagNil(b bool)`

 SetExcludeAvalaraFlagNil sets the value for ExcludeAvalaraFlag to be an explicit nil

### UnsetExcludeAvalaraFlag
`func (o *BillingSetup) UnsetExcludeAvalaraFlag()`

UnsetExcludeAvalaraFlag ensures that no value is present for ExcludeAvalaraFlag, not even an explicit nil
### GetInfo

`func (o *BillingSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BillingSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BillingSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BillingSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


