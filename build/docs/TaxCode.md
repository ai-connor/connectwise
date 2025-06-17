# TaxCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  Max length: 8; | 
**Description** | **string** |  Max length: 50; | 
**InvoiceCaption** | **string** |  Max length: 25; | 
**Country** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**EffectiveDate** | **time.Time** |  | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**DisplayOnInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**CanadaCalculateGSTFlag** | Pointer to **NullableBool** |  | [optional] 
**CancelDate** | Pointer to **time.Time** |  | [optional] 
**LevelOneRate** | Pointer to **NullableFloat64** |  | [optional] 
**LevelOneRateType** | Pointer to **NullableString** |  | [optional] 
**LevelOneTaxableMax** | Pointer to **NullableFloat64** |  | [optional] 
**LevelOneCaption** | Pointer to **string** |  Max length: 25; | [optional] 
**LevelOneTaxCodeXref** | Pointer to **string** |  Max length: 50; | [optional] 
**LevelOneAgencyXref** | Pointer to **string** |  Max length: 100; | [optional] 
**LevelOneServicesFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelOneExpensesFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelOneProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelOneApplySingleUnitFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelOneApplySingleUnitMin** | Pointer to **NullableFloat64** |  | [optional] 
**LevelOneApplySingleUnitMax** | Pointer to **NullableFloat64** |  | [optional] 
**LevelTwoRate** | Pointer to **NullableFloat64** |  | [optional] 
**LevelTwoRateType** | Pointer to **NullableString** |  | [optional] 
**LevelTwoTaxableMax** | Pointer to **NullableFloat64** |  | [optional] 
**LevelTwoCaption** | Pointer to **string** |  Max length: 25; | [optional] 
**LevelTwoTaxCodeXref** | Pointer to **string** |  Max length: 50; | [optional] 
**LevelTwoAgencyXref** | Pointer to **string** |  Max length: 100; | [optional] 
**LevelTwoServicesFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelTwoExpensesFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelTwoProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelTwoApplySingleUnitFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelTwoApplySingleUnitMin** | Pointer to **NullableFloat64** |  | [optional] 
**LevelTwoApplySingleUnitMax** | Pointer to **NullableFloat64** |  | [optional] 
**LevelThreeRate** | Pointer to **NullableFloat64** |  | [optional] 
**LevelThreeRateType** | Pointer to **NullableString** |  | [optional] 
**LevelThreeTaxableMax** | Pointer to **NullableFloat64** |  | [optional] 
**LevelThreeCaption** | Pointer to **string** |  Max length: 25; | [optional] 
**LevelThreeTaxCodeXref** | Pointer to **string** |  Max length: 50; | [optional] 
**LevelThreeAgencyXref** | Pointer to **string** |  Max length: 100; | [optional] 
**LevelThreeServicesFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelThreeExpensesFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelThreeProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelThreeApplySingleUnitFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelThreeApplySingleUnitMin** | Pointer to **NullableFloat64** |  | [optional] 
**LevelThreeApplySingleUnitMax** | Pointer to **NullableFloat64** |  | [optional] 
**LevelFourRate** | Pointer to **NullableFloat64** |  | [optional] 
**LevelFourRateType** | Pointer to **NullableString** |  | [optional] 
**LevelFourTaxableMax** | Pointer to **NullableFloat64** |  | [optional] 
**LevelFourCaption** | Pointer to **string** |  Max length: 25; | [optional] 
**LevelFourTaxCodeXref** | Pointer to **string** |  Max length: 50; | [optional] 
**LevelFourAgencyXref** | Pointer to **string** |  Max length: 100; | [optional] 
**LevelFourServicesFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelFourExpensesFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelFourProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelFourApplySingleUnitFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelFourApplySingleUnitMin** | Pointer to **NullableFloat64** |  | [optional] 
**LevelFourApplySingleUnitMax** | Pointer to **NullableFloat64** |  | [optional] 
**LevelFiveRate** | Pointer to **NullableFloat64** |  | [optional] 
**LevelFiveRateType** | Pointer to **NullableString** |  | [optional] 
**LevelFiveTaxableMax** | Pointer to **NullableFloat64** |  | [optional] 
**LevelFiveCaption** | Pointer to **string** |  Max length: 25; | [optional] 
**LevelFiveTaxCodeXref** | Pointer to **string** |  Max length: 50; | [optional] 
**LevelFiveAgencyXref** | Pointer to **string** |  Max length: 100; | [optional] 
**LevelFiveServicesFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelFiveExpensesFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelFiveProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelFiveApplySingleUnitFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelFiveApplySingleUnitMin** | Pointer to **NullableFloat64** |  | [optional] 
**LevelFiveApplySingleUnitMax** | Pointer to **NullableFloat64** |  | [optional] 
**LevelSixRate** | Pointer to **NullableFloat64** |  | [optional] 
**LevelSixRateType** | Pointer to **NullableString** |  | [optional] 
**LevelSixTaxableMax** | Pointer to **NullableFloat64** |  | [optional] 
**LevelSixCaption** | Pointer to **string** |  Max length: 25; | [optional] 
**LevelSixTaxCodeXref** | Pointer to **string** |  Max length: 50; | [optional] 
**LevelSixAgencyXref** | Pointer to **string** |  Max length: 100; | [optional] 
**LevelSixServicesFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelSixExpensesFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelSixProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelSixApplySingleUnitFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelSixApplySingleUnitMin** | Pointer to **NullableFloat64** |  | [optional] 
**LevelSixApplySingleUnitMax** | Pointer to **NullableFloat64** |  | [optional] 
**WorkRoleIds** | Pointer to **[]int32** | Array of work role exemptions for the tax code. | [optional] 
**AddAllWorkRoles** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllWorkRoles** | Pointer to **NullableBool** |  | [optional] 
**ExpenseTypeIds** | Pointer to **[]int32** | Array of expense type exemptions for the tax code. | [optional] 
**AddAllExpenseTypes** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllExpenseTypes** | Pointer to **NullableBool** |  | [optional] 
**ProductTypeIds** | Pointer to **[]int32** | Array of product type exemptions for the tax code. | [optional] 
**AddAllProductTypes** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllProductTypes** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTaxCode

`func NewTaxCode(identifier string, description string, invoiceCaption string, effectiveDate time.Time, ) *TaxCode`

NewTaxCode instantiates a new TaxCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCodeWithDefaults

`func NewTaxCodeWithDefaults() *TaxCode`

NewTaxCodeWithDefaults instantiates a new TaxCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxCode) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxCode) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxCode) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaxCode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *TaxCode) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *TaxCode) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *TaxCode) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetDescription

`func (o *TaxCode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxCode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxCode) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInvoiceCaption

`func (o *TaxCode) GetInvoiceCaption() string`

GetInvoiceCaption returns the InvoiceCaption field if non-nil, zero value otherwise.

### GetInvoiceCaptionOk

`func (o *TaxCode) GetInvoiceCaptionOk() (*string, bool)`

GetInvoiceCaptionOk returns a tuple with the InvoiceCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCaption

`func (o *TaxCode) SetInvoiceCaption(v string)`

SetInvoiceCaption sets InvoiceCaption field to given value.


### GetCountry

`func (o *TaxCode) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TaxCode) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TaxCode) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TaxCode) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *TaxCode) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *TaxCode) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *TaxCode) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetDefaultFlag

`func (o *TaxCode) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *TaxCode) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *TaxCode) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *TaxCode) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *TaxCode) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *TaxCode) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetDisplayOnInvoiceFlag

`func (o *TaxCode) GetDisplayOnInvoiceFlag() bool`

GetDisplayOnInvoiceFlag returns the DisplayOnInvoiceFlag field if non-nil, zero value otherwise.

### GetDisplayOnInvoiceFlagOk

`func (o *TaxCode) GetDisplayOnInvoiceFlagOk() (*bool, bool)`

GetDisplayOnInvoiceFlagOk returns a tuple with the DisplayOnInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOnInvoiceFlag

`func (o *TaxCode) SetDisplayOnInvoiceFlag(v bool)`

SetDisplayOnInvoiceFlag sets DisplayOnInvoiceFlag field to given value.

### HasDisplayOnInvoiceFlag

`func (o *TaxCode) HasDisplayOnInvoiceFlag() bool`

HasDisplayOnInvoiceFlag returns a boolean if a field has been set.

### SetDisplayOnInvoiceFlagNil

`func (o *TaxCode) SetDisplayOnInvoiceFlagNil(b bool)`

 SetDisplayOnInvoiceFlagNil sets the value for DisplayOnInvoiceFlag to be an explicit nil

### UnsetDisplayOnInvoiceFlag
`func (o *TaxCode) UnsetDisplayOnInvoiceFlag()`

UnsetDisplayOnInvoiceFlag ensures that no value is present for DisplayOnInvoiceFlag, not even an explicit nil
### GetCanadaCalculateGSTFlag

`func (o *TaxCode) GetCanadaCalculateGSTFlag() bool`

GetCanadaCalculateGSTFlag returns the CanadaCalculateGSTFlag field if non-nil, zero value otherwise.

### GetCanadaCalculateGSTFlagOk

`func (o *TaxCode) GetCanadaCalculateGSTFlagOk() (*bool, bool)`

GetCanadaCalculateGSTFlagOk returns a tuple with the CanadaCalculateGSTFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanadaCalculateGSTFlag

`func (o *TaxCode) SetCanadaCalculateGSTFlag(v bool)`

SetCanadaCalculateGSTFlag sets CanadaCalculateGSTFlag field to given value.

### HasCanadaCalculateGSTFlag

`func (o *TaxCode) HasCanadaCalculateGSTFlag() bool`

HasCanadaCalculateGSTFlag returns a boolean if a field has been set.

### SetCanadaCalculateGSTFlagNil

`func (o *TaxCode) SetCanadaCalculateGSTFlagNil(b bool)`

 SetCanadaCalculateGSTFlagNil sets the value for CanadaCalculateGSTFlag to be an explicit nil

### UnsetCanadaCalculateGSTFlag
`func (o *TaxCode) UnsetCanadaCalculateGSTFlag()`

UnsetCanadaCalculateGSTFlag ensures that no value is present for CanadaCalculateGSTFlag, not even an explicit nil
### GetCancelDate

`func (o *TaxCode) GetCancelDate() time.Time`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *TaxCode) GetCancelDateOk() (*time.Time, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *TaxCode) SetCancelDate(v time.Time)`

SetCancelDate sets CancelDate field to given value.

### HasCancelDate

`func (o *TaxCode) HasCancelDate() bool`

HasCancelDate returns a boolean if a field has been set.

### GetLevelOneRate

`func (o *TaxCode) GetLevelOneRate() float64`

GetLevelOneRate returns the LevelOneRate field if non-nil, zero value otherwise.

### GetLevelOneRateOk

`func (o *TaxCode) GetLevelOneRateOk() (*float64, bool)`

GetLevelOneRateOk returns a tuple with the LevelOneRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneRate

`func (o *TaxCode) SetLevelOneRate(v float64)`

SetLevelOneRate sets LevelOneRate field to given value.

### HasLevelOneRate

`func (o *TaxCode) HasLevelOneRate() bool`

HasLevelOneRate returns a boolean if a field has been set.

### SetLevelOneRateNil

`func (o *TaxCode) SetLevelOneRateNil(b bool)`

 SetLevelOneRateNil sets the value for LevelOneRate to be an explicit nil

### UnsetLevelOneRate
`func (o *TaxCode) UnsetLevelOneRate()`

UnsetLevelOneRate ensures that no value is present for LevelOneRate, not even an explicit nil
### GetLevelOneRateType

`func (o *TaxCode) GetLevelOneRateType() string`

GetLevelOneRateType returns the LevelOneRateType field if non-nil, zero value otherwise.

### GetLevelOneRateTypeOk

`func (o *TaxCode) GetLevelOneRateTypeOk() (*string, bool)`

GetLevelOneRateTypeOk returns a tuple with the LevelOneRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneRateType

`func (o *TaxCode) SetLevelOneRateType(v string)`

SetLevelOneRateType sets LevelOneRateType field to given value.

### HasLevelOneRateType

`func (o *TaxCode) HasLevelOneRateType() bool`

HasLevelOneRateType returns a boolean if a field has been set.

### SetLevelOneRateTypeNil

`func (o *TaxCode) SetLevelOneRateTypeNil(b bool)`

 SetLevelOneRateTypeNil sets the value for LevelOneRateType to be an explicit nil

### UnsetLevelOneRateType
`func (o *TaxCode) UnsetLevelOneRateType()`

UnsetLevelOneRateType ensures that no value is present for LevelOneRateType, not even an explicit nil
### GetLevelOneTaxableMax

`func (o *TaxCode) GetLevelOneTaxableMax() float64`

GetLevelOneTaxableMax returns the LevelOneTaxableMax field if non-nil, zero value otherwise.

### GetLevelOneTaxableMaxOk

`func (o *TaxCode) GetLevelOneTaxableMaxOk() (*float64, bool)`

GetLevelOneTaxableMaxOk returns a tuple with the LevelOneTaxableMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneTaxableMax

`func (o *TaxCode) SetLevelOneTaxableMax(v float64)`

SetLevelOneTaxableMax sets LevelOneTaxableMax field to given value.

### HasLevelOneTaxableMax

`func (o *TaxCode) HasLevelOneTaxableMax() bool`

HasLevelOneTaxableMax returns a boolean if a field has been set.

### SetLevelOneTaxableMaxNil

`func (o *TaxCode) SetLevelOneTaxableMaxNil(b bool)`

 SetLevelOneTaxableMaxNil sets the value for LevelOneTaxableMax to be an explicit nil

### UnsetLevelOneTaxableMax
`func (o *TaxCode) UnsetLevelOneTaxableMax()`

UnsetLevelOneTaxableMax ensures that no value is present for LevelOneTaxableMax, not even an explicit nil
### GetLevelOneCaption

`func (o *TaxCode) GetLevelOneCaption() string`

GetLevelOneCaption returns the LevelOneCaption field if non-nil, zero value otherwise.

### GetLevelOneCaptionOk

`func (o *TaxCode) GetLevelOneCaptionOk() (*string, bool)`

GetLevelOneCaptionOk returns a tuple with the LevelOneCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneCaption

`func (o *TaxCode) SetLevelOneCaption(v string)`

SetLevelOneCaption sets LevelOneCaption field to given value.

### HasLevelOneCaption

`func (o *TaxCode) HasLevelOneCaption() bool`

HasLevelOneCaption returns a boolean if a field has been set.

### GetLevelOneTaxCodeXref

`func (o *TaxCode) GetLevelOneTaxCodeXref() string`

GetLevelOneTaxCodeXref returns the LevelOneTaxCodeXref field if non-nil, zero value otherwise.

### GetLevelOneTaxCodeXrefOk

`func (o *TaxCode) GetLevelOneTaxCodeXrefOk() (*string, bool)`

GetLevelOneTaxCodeXrefOk returns a tuple with the LevelOneTaxCodeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneTaxCodeXref

`func (o *TaxCode) SetLevelOneTaxCodeXref(v string)`

SetLevelOneTaxCodeXref sets LevelOneTaxCodeXref field to given value.

### HasLevelOneTaxCodeXref

`func (o *TaxCode) HasLevelOneTaxCodeXref() bool`

HasLevelOneTaxCodeXref returns a boolean if a field has been set.

### GetLevelOneAgencyXref

`func (o *TaxCode) GetLevelOneAgencyXref() string`

GetLevelOneAgencyXref returns the LevelOneAgencyXref field if non-nil, zero value otherwise.

### GetLevelOneAgencyXrefOk

`func (o *TaxCode) GetLevelOneAgencyXrefOk() (*string, bool)`

GetLevelOneAgencyXrefOk returns a tuple with the LevelOneAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneAgencyXref

`func (o *TaxCode) SetLevelOneAgencyXref(v string)`

SetLevelOneAgencyXref sets LevelOneAgencyXref field to given value.

### HasLevelOneAgencyXref

`func (o *TaxCode) HasLevelOneAgencyXref() bool`

HasLevelOneAgencyXref returns a boolean if a field has been set.

### GetLevelOneServicesFlag

`func (o *TaxCode) GetLevelOneServicesFlag() bool`

GetLevelOneServicesFlag returns the LevelOneServicesFlag field if non-nil, zero value otherwise.

### GetLevelOneServicesFlagOk

`func (o *TaxCode) GetLevelOneServicesFlagOk() (*bool, bool)`

GetLevelOneServicesFlagOk returns a tuple with the LevelOneServicesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneServicesFlag

`func (o *TaxCode) SetLevelOneServicesFlag(v bool)`

SetLevelOneServicesFlag sets LevelOneServicesFlag field to given value.

### HasLevelOneServicesFlag

`func (o *TaxCode) HasLevelOneServicesFlag() bool`

HasLevelOneServicesFlag returns a boolean if a field has been set.

### SetLevelOneServicesFlagNil

`func (o *TaxCode) SetLevelOneServicesFlagNil(b bool)`

 SetLevelOneServicesFlagNil sets the value for LevelOneServicesFlag to be an explicit nil

### UnsetLevelOneServicesFlag
`func (o *TaxCode) UnsetLevelOneServicesFlag()`

UnsetLevelOneServicesFlag ensures that no value is present for LevelOneServicesFlag, not even an explicit nil
### GetLevelOneExpensesFlag

`func (o *TaxCode) GetLevelOneExpensesFlag() bool`

GetLevelOneExpensesFlag returns the LevelOneExpensesFlag field if non-nil, zero value otherwise.

### GetLevelOneExpensesFlagOk

`func (o *TaxCode) GetLevelOneExpensesFlagOk() (*bool, bool)`

GetLevelOneExpensesFlagOk returns a tuple with the LevelOneExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneExpensesFlag

`func (o *TaxCode) SetLevelOneExpensesFlag(v bool)`

SetLevelOneExpensesFlag sets LevelOneExpensesFlag field to given value.

### HasLevelOneExpensesFlag

`func (o *TaxCode) HasLevelOneExpensesFlag() bool`

HasLevelOneExpensesFlag returns a boolean if a field has been set.

### SetLevelOneExpensesFlagNil

`func (o *TaxCode) SetLevelOneExpensesFlagNil(b bool)`

 SetLevelOneExpensesFlagNil sets the value for LevelOneExpensesFlag to be an explicit nil

### UnsetLevelOneExpensesFlag
`func (o *TaxCode) UnsetLevelOneExpensesFlag()`

UnsetLevelOneExpensesFlag ensures that no value is present for LevelOneExpensesFlag, not even an explicit nil
### GetLevelOneProductsFlag

`func (o *TaxCode) GetLevelOneProductsFlag() bool`

GetLevelOneProductsFlag returns the LevelOneProductsFlag field if non-nil, zero value otherwise.

### GetLevelOneProductsFlagOk

`func (o *TaxCode) GetLevelOneProductsFlagOk() (*bool, bool)`

GetLevelOneProductsFlagOk returns a tuple with the LevelOneProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneProductsFlag

`func (o *TaxCode) SetLevelOneProductsFlag(v bool)`

SetLevelOneProductsFlag sets LevelOneProductsFlag field to given value.

### HasLevelOneProductsFlag

`func (o *TaxCode) HasLevelOneProductsFlag() bool`

HasLevelOneProductsFlag returns a boolean if a field has been set.

### SetLevelOneProductsFlagNil

`func (o *TaxCode) SetLevelOneProductsFlagNil(b bool)`

 SetLevelOneProductsFlagNil sets the value for LevelOneProductsFlag to be an explicit nil

### UnsetLevelOneProductsFlag
`func (o *TaxCode) UnsetLevelOneProductsFlag()`

UnsetLevelOneProductsFlag ensures that no value is present for LevelOneProductsFlag, not even an explicit nil
### GetLevelOneApplySingleUnitFlag

`func (o *TaxCode) GetLevelOneApplySingleUnitFlag() bool`

GetLevelOneApplySingleUnitFlag returns the LevelOneApplySingleUnitFlag field if non-nil, zero value otherwise.

### GetLevelOneApplySingleUnitFlagOk

`func (o *TaxCode) GetLevelOneApplySingleUnitFlagOk() (*bool, bool)`

GetLevelOneApplySingleUnitFlagOk returns a tuple with the LevelOneApplySingleUnitFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneApplySingleUnitFlag

`func (o *TaxCode) SetLevelOneApplySingleUnitFlag(v bool)`

SetLevelOneApplySingleUnitFlag sets LevelOneApplySingleUnitFlag field to given value.

### HasLevelOneApplySingleUnitFlag

`func (o *TaxCode) HasLevelOneApplySingleUnitFlag() bool`

HasLevelOneApplySingleUnitFlag returns a boolean if a field has been set.

### SetLevelOneApplySingleUnitFlagNil

`func (o *TaxCode) SetLevelOneApplySingleUnitFlagNil(b bool)`

 SetLevelOneApplySingleUnitFlagNil sets the value for LevelOneApplySingleUnitFlag to be an explicit nil

### UnsetLevelOneApplySingleUnitFlag
`func (o *TaxCode) UnsetLevelOneApplySingleUnitFlag()`

UnsetLevelOneApplySingleUnitFlag ensures that no value is present for LevelOneApplySingleUnitFlag, not even an explicit nil
### GetLevelOneApplySingleUnitMin

`func (o *TaxCode) GetLevelOneApplySingleUnitMin() float64`

GetLevelOneApplySingleUnitMin returns the LevelOneApplySingleUnitMin field if non-nil, zero value otherwise.

### GetLevelOneApplySingleUnitMinOk

`func (o *TaxCode) GetLevelOneApplySingleUnitMinOk() (*float64, bool)`

GetLevelOneApplySingleUnitMinOk returns a tuple with the LevelOneApplySingleUnitMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneApplySingleUnitMin

`func (o *TaxCode) SetLevelOneApplySingleUnitMin(v float64)`

SetLevelOneApplySingleUnitMin sets LevelOneApplySingleUnitMin field to given value.

### HasLevelOneApplySingleUnitMin

`func (o *TaxCode) HasLevelOneApplySingleUnitMin() bool`

HasLevelOneApplySingleUnitMin returns a boolean if a field has been set.

### SetLevelOneApplySingleUnitMinNil

`func (o *TaxCode) SetLevelOneApplySingleUnitMinNil(b bool)`

 SetLevelOneApplySingleUnitMinNil sets the value for LevelOneApplySingleUnitMin to be an explicit nil

### UnsetLevelOneApplySingleUnitMin
`func (o *TaxCode) UnsetLevelOneApplySingleUnitMin()`

UnsetLevelOneApplySingleUnitMin ensures that no value is present for LevelOneApplySingleUnitMin, not even an explicit nil
### GetLevelOneApplySingleUnitMax

`func (o *TaxCode) GetLevelOneApplySingleUnitMax() float64`

GetLevelOneApplySingleUnitMax returns the LevelOneApplySingleUnitMax field if non-nil, zero value otherwise.

### GetLevelOneApplySingleUnitMaxOk

`func (o *TaxCode) GetLevelOneApplySingleUnitMaxOk() (*float64, bool)`

GetLevelOneApplySingleUnitMaxOk returns a tuple with the LevelOneApplySingleUnitMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneApplySingleUnitMax

`func (o *TaxCode) SetLevelOneApplySingleUnitMax(v float64)`

SetLevelOneApplySingleUnitMax sets LevelOneApplySingleUnitMax field to given value.

### HasLevelOneApplySingleUnitMax

`func (o *TaxCode) HasLevelOneApplySingleUnitMax() bool`

HasLevelOneApplySingleUnitMax returns a boolean if a field has been set.

### SetLevelOneApplySingleUnitMaxNil

`func (o *TaxCode) SetLevelOneApplySingleUnitMaxNil(b bool)`

 SetLevelOneApplySingleUnitMaxNil sets the value for LevelOneApplySingleUnitMax to be an explicit nil

### UnsetLevelOneApplySingleUnitMax
`func (o *TaxCode) UnsetLevelOneApplySingleUnitMax()`

UnsetLevelOneApplySingleUnitMax ensures that no value is present for LevelOneApplySingleUnitMax, not even an explicit nil
### GetLevelTwoRate

`func (o *TaxCode) GetLevelTwoRate() float64`

GetLevelTwoRate returns the LevelTwoRate field if non-nil, zero value otherwise.

### GetLevelTwoRateOk

`func (o *TaxCode) GetLevelTwoRateOk() (*float64, bool)`

GetLevelTwoRateOk returns a tuple with the LevelTwoRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoRate

`func (o *TaxCode) SetLevelTwoRate(v float64)`

SetLevelTwoRate sets LevelTwoRate field to given value.

### HasLevelTwoRate

`func (o *TaxCode) HasLevelTwoRate() bool`

HasLevelTwoRate returns a boolean if a field has been set.

### SetLevelTwoRateNil

`func (o *TaxCode) SetLevelTwoRateNil(b bool)`

 SetLevelTwoRateNil sets the value for LevelTwoRate to be an explicit nil

### UnsetLevelTwoRate
`func (o *TaxCode) UnsetLevelTwoRate()`

UnsetLevelTwoRate ensures that no value is present for LevelTwoRate, not even an explicit nil
### GetLevelTwoRateType

`func (o *TaxCode) GetLevelTwoRateType() string`

GetLevelTwoRateType returns the LevelTwoRateType field if non-nil, zero value otherwise.

### GetLevelTwoRateTypeOk

`func (o *TaxCode) GetLevelTwoRateTypeOk() (*string, bool)`

GetLevelTwoRateTypeOk returns a tuple with the LevelTwoRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoRateType

`func (o *TaxCode) SetLevelTwoRateType(v string)`

SetLevelTwoRateType sets LevelTwoRateType field to given value.

### HasLevelTwoRateType

`func (o *TaxCode) HasLevelTwoRateType() bool`

HasLevelTwoRateType returns a boolean if a field has been set.

### SetLevelTwoRateTypeNil

`func (o *TaxCode) SetLevelTwoRateTypeNil(b bool)`

 SetLevelTwoRateTypeNil sets the value for LevelTwoRateType to be an explicit nil

### UnsetLevelTwoRateType
`func (o *TaxCode) UnsetLevelTwoRateType()`

UnsetLevelTwoRateType ensures that no value is present for LevelTwoRateType, not even an explicit nil
### GetLevelTwoTaxableMax

`func (o *TaxCode) GetLevelTwoTaxableMax() float64`

GetLevelTwoTaxableMax returns the LevelTwoTaxableMax field if non-nil, zero value otherwise.

### GetLevelTwoTaxableMaxOk

`func (o *TaxCode) GetLevelTwoTaxableMaxOk() (*float64, bool)`

GetLevelTwoTaxableMaxOk returns a tuple with the LevelTwoTaxableMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoTaxableMax

`func (o *TaxCode) SetLevelTwoTaxableMax(v float64)`

SetLevelTwoTaxableMax sets LevelTwoTaxableMax field to given value.

### HasLevelTwoTaxableMax

`func (o *TaxCode) HasLevelTwoTaxableMax() bool`

HasLevelTwoTaxableMax returns a boolean if a field has been set.

### SetLevelTwoTaxableMaxNil

`func (o *TaxCode) SetLevelTwoTaxableMaxNil(b bool)`

 SetLevelTwoTaxableMaxNil sets the value for LevelTwoTaxableMax to be an explicit nil

### UnsetLevelTwoTaxableMax
`func (o *TaxCode) UnsetLevelTwoTaxableMax()`

UnsetLevelTwoTaxableMax ensures that no value is present for LevelTwoTaxableMax, not even an explicit nil
### GetLevelTwoCaption

`func (o *TaxCode) GetLevelTwoCaption() string`

GetLevelTwoCaption returns the LevelTwoCaption field if non-nil, zero value otherwise.

### GetLevelTwoCaptionOk

`func (o *TaxCode) GetLevelTwoCaptionOk() (*string, bool)`

GetLevelTwoCaptionOk returns a tuple with the LevelTwoCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoCaption

`func (o *TaxCode) SetLevelTwoCaption(v string)`

SetLevelTwoCaption sets LevelTwoCaption field to given value.

### HasLevelTwoCaption

`func (o *TaxCode) HasLevelTwoCaption() bool`

HasLevelTwoCaption returns a boolean if a field has been set.

### GetLevelTwoTaxCodeXref

`func (o *TaxCode) GetLevelTwoTaxCodeXref() string`

GetLevelTwoTaxCodeXref returns the LevelTwoTaxCodeXref field if non-nil, zero value otherwise.

### GetLevelTwoTaxCodeXrefOk

`func (o *TaxCode) GetLevelTwoTaxCodeXrefOk() (*string, bool)`

GetLevelTwoTaxCodeXrefOk returns a tuple with the LevelTwoTaxCodeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoTaxCodeXref

`func (o *TaxCode) SetLevelTwoTaxCodeXref(v string)`

SetLevelTwoTaxCodeXref sets LevelTwoTaxCodeXref field to given value.

### HasLevelTwoTaxCodeXref

`func (o *TaxCode) HasLevelTwoTaxCodeXref() bool`

HasLevelTwoTaxCodeXref returns a boolean if a field has been set.

### GetLevelTwoAgencyXref

`func (o *TaxCode) GetLevelTwoAgencyXref() string`

GetLevelTwoAgencyXref returns the LevelTwoAgencyXref field if non-nil, zero value otherwise.

### GetLevelTwoAgencyXrefOk

`func (o *TaxCode) GetLevelTwoAgencyXrefOk() (*string, bool)`

GetLevelTwoAgencyXrefOk returns a tuple with the LevelTwoAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoAgencyXref

`func (o *TaxCode) SetLevelTwoAgencyXref(v string)`

SetLevelTwoAgencyXref sets LevelTwoAgencyXref field to given value.

### HasLevelTwoAgencyXref

`func (o *TaxCode) HasLevelTwoAgencyXref() bool`

HasLevelTwoAgencyXref returns a boolean if a field has been set.

### GetLevelTwoServicesFlag

`func (o *TaxCode) GetLevelTwoServicesFlag() bool`

GetLevelTwoServicesFlag returns the LevelTwoServicesFlag field if non-nil, zero value otherwise.

### GetLevelTwoServicesFlagOk

`func (o *TaxCode) GetLevelTwoServicesFlagOk() (*bool, bool)`

GetLevelTwoServicesFlagOk returns a tuple with the LevelTwoServicesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoServicesFlag

`func (o *TaxCode) SetLevelTwoServicesFlag(v bool)`

SetLevelTwoServicesFlag sets LevelTwoServicesFlag field to given value.

### HasLevelTwoServicesFlag

`func (o *TaxCode) HasLevelTwoServicesFlag() bool`

HasLevelTwoServicesFlag returns a boolean if a field has been set.

### SetLevelTwoServicesFlagNil

`func (o *TaxCode) SetLevelTwoServicesFlagNil(b bool)`

 SetLevelTwoServicesFlagNil sets the value for LevelTwoServicesFlag to be an explicit nil

### UnsetLevelTwoServicesFlag
`func (o *TaxCode) UnsetLevelTwoServicesFlag()`

UnsetLevelTwoServicesFlag ensures that no value is present for LevelTwoServicesFlag, not even an explicit nil
### GetLevelTwoExpensesFlag

`func (o *TaxCode) GetLevelTwoExpensesFlag() bool`

GetLevelTwoExpensesFlag returns the LevelTwoExpensesFlag field if non-nil, zero value otherwise.

### GetLevelTwoExpensesFlagOk

`func (o *TaxCode) GetLevelTwoExpensesFlagOk() (*bool, bool)`

GetLevelTwoExpensesFlagOk returns a tuple with the LevelTwoExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoExpensesFlag

`func (o *TaxCode) SetLevelTwoExpensesFlag(v bool)`

SetLevelTwoExpensesFlag sets LevelTwoExpensesFlag field to given value.

### HasLevelTwoExpensesFlag

`func (o *TaxCode) HasLevelTwoExpensesFlag() bool`

HasLevelTwoExpensesFlag returns a boolean if a field has been set.

### SetLevelTwoExpensesFlagNil

`func (o *TaxCode) SetLevelTwoExpensesFlagNil(b bool)`

 SetLevelTwoExpensesFlagNil sets the value for LevelTwoExpensesFlag to be an explicit nil

### UnsetLevelTwoExpensesFlag
`func (o *TaxCode) UnsetLevelTwoExpensesFlag()`

UnsetLevelTwoExpensesFlag ensures that no value is present for LevelTwoExpensesFlag, not even an explicit nil
### GetLevelTwoProductsFlag

`func (o *TaxCode) GetLevelTwoProductsFlag() bool`

GetLevelTwoProductsFlag returns the LevelTwoProductsFlag field if non-nil, zero value otherwise.

### GetLevelTwoProductsFlagOk

`func (o *TaxCode) GetLevelTwoProductsFlagOk() (*bool, bool)`

GetLevelTwoProductsFlagOk returns a tuple with the LevelTwoProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoProductsFlag

`func (o *TaxCode) SetLevelTwoProductsFlag(v bool)`

SetLevelTwoProductsFlag sets LevelTwoProductsFlag field to given value.

### HasLevelTwoProductsFlag

`func (o *TaxCode) HasLevelTwoProductsFlag() bool`

HasLevelTwoProductsFlag returns a boolean if a field has been set.

### SetLevelTwoProductsFlagNil

`func (o *TaxCode) SetLevelTwoProductsFlagNil(b bool)`

 SetLevelTwoProductsFlagNil sets the value for LevelTwoProductsFlag to be an explicit nil

### UnsetLevelTwoProductsFlag
`func (o *TaxCode) UnsetLevelTwoProductsFlag()`

UnsetLevelTwoProductsFlag ensures that no value is present for LevelTwoProductsFlag, not even an explicit nil
### GetLevelTwoApplySingleUnitFlag

`func (o *TaxCode) GetLevelTwoApplySingleUnitFlag() bool`

GetLevelTwoApplySingleUnitFlag returns the LevelTwoApplySingleUnitFlag field if non-nil, zero value otherwise.

### GetLevelTwoApplySingleUnitFlagOk

`func (o *TaxCode) GetLevelTwoApplySingleUnitFlagOk() (*bool, bool)`

GetLevelTwoApplySingleUnitFlagOk returns a tuple with the LevelTwoApplySingleUnitFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoApplySingleUnitFlag

`func (o *TaxCode) SetLevelTwoApplySingleUnitFlag(v bool)`

SetLevelTwoApplySingleUnitFlag sets LevelTwoApplySingleUnitFlag field to given value.

### HasLevelTwoApplySingleUnitFlag

`func (o *TaxCode) HasLevelTwoApplySingleUnitFlag() bool`

HasLevelTwoApplySingleUnitFlag returns a boolean if a field has been set.

### SetLevelTwoApplySingleUnitFlagNil

`func (o *TaxCode) SetLevelTwoApplySingleUnitFlagNil(b bool)`

 SetLevelTwoApplySingleUnitFlagNil sets the value for LevelTwoApplySingleUnitFlag to be an explicit nil

### UnsetLevelTwoApplySingleUnitFlag
`func (o *TaxCode) UnsetLevelTwoApplySingleUnitFlag()`

UnsetLevelTwoApplySingleUnitFlag ensures that no value is present for LevelTwoApplySingleUnitFlag, not even an explicit nil
### GetLevelTwoApplySingleUnitMin

`func (o *TaxCode) GetLevelTwoApplySingleUnitMin() float64`

GetLevelTwoApplySingleUnitMin returns the LevelTwoApplySingleUnitMin field if non-nil, zero value otherwise.

### GetLevelTwoApplySingleUnitMinOk

`func (o *TaxCode) GetLevelTwoApplySingleUnitMinOk() (*float64, bool)`

GetLevelTwoApplySingleUnitMinOk returns a tuple with the LevelTwoApplySingleUnitMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoApplySingleUnitMin

`func (o *TaxCode) SetLevelTwoApplySingleUnitMin(v float64)`

SetLevelTwoApplySingleUnitMin sets LevelTwoApplySingleUnitMin field to given value.

### HasLevelTwoApplySingleUnitMin

`func (o *TaxCode) HasLevelTwoApplySingleUnitMin() bool`

HasLevelTwoApplySingleUnitMin returns a boolean if a field has been set.

### SetLevelTwoApplySingleUnitMinNil

`func (o *TaxCode) SetLevelTwoApplySingleUnitMinNil(b bool)`

 SetLevelTwoApplySingleUnitMinNil sets the value for LevelTwoApplySingleUnitMin to be an explicit nil

### UnsetLevelTwoApplySingleUnitMin
`func (o *TaxCode) UnsetLevelTwoApplySingleUnitMin()`

UnsetLevelTwoApplySingleUnitMin ensures that no value is present for LevelTwoApplySingleUnitMin, not even an explicit nil
### GetLevelTwoApplySingleUnitMax

`func (o *TaxCode) GetLevelTwoApplySingleUnitMax() float64`

GetLevelTwoApplySingleUnitMax returns the LevelTwoApplySingleUnitMax field if non-nil, zero value otherwise.

### GetLevelTwoApplySingleUnitMaxOk

`func (o *TaxCode) GetLevelTwoApplySingleUnitMaxOk() (*float64, bool)`

GetLevelTwoApplySingleUnitMaxOk returns a tuple with the LevelTwoApplySingleUnitMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoApplySingleUnitMax

`func (o *TaxCode) SetLevelTwoApplySingleUnitMax(v float64)`

SetLevelTwoApplySingleUnitMax sets LevelTwoApplySingleUnitMax field to given value.

### HasLevelTwoApplySingleUnitMax

`func (o *TaxCode) HasLevelTwoApplySingleUnitMax() bool`

HasLevelTwoApplySingleUnitMax returns a boolean if a field has been set.

### SetLevelTwoApplySingleUnitMaxNil

`func (o *TaxCode) SetLevelTwoApplySingleUnitMaxNil(b bool)`

 SetLevelTwoApplySingleUnitMaxNil sets the value for LevelTwoApplySingleUnitMax to be an explicit nil

### UnsetLevelTwoApplySingleUnitMax
`func (o *TaxCode) UnsetLevelTwoApplySingleUnitMax()`

UnsetLevelTwoApplySingleUnitMax ensures that no value is present for LevelTwoApplySingleUnitMax, not even an explicit nil
### GetLevelThreeRate

`func (o *TaxCode) GetLevelThreeRate() float64`

GetLevelThreeRate returns the LevelThreeRate field if non-nil, zero value otherwise.

### GetLevelThreeRateOk

`func (o *TaxCode) GetLevelThreeRateOk() (*float64, bool)`

GetLevelThreeRateOk returns a tuple with the LevelThreeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeRate

`func (o *TaxCode) SetLevelThreeRate(v float64)`

SetLevelThreeRate sets LevelThreeRate field to given value.

### HasLevelThreeRate

`func (o *TaxCode) HasLevelThreeRate() bool`

HasLevelThreeRate returns a boolean if a field has been set.

### SetLevelThreeRateNil

`func (o *TaxCode) SetLevelThreeRateNil(b bool)`

 SetLevelThreeRateNil sets the value for LevelThreeRate to be an explicit nil

### UnsetLevelThreeRate
`func (o *TaxCode) UnsetLevelThreeRate()`

UnsetLevelThreeRate ensures that no value is present for LevelThreeRate, not even an explicit nil
### GetLevelThreeRateType

`func (o *TaxCode) GetLevelThreeRateType() string`

GetLevelThreeRateType returns the LevelThreeRateType field if non-nil, zero value otherwise.

### GetLevelThreeRateTypeOk

`func (o *TaxCode) GetLevelThreeRateTypeOk() (*string, bool)`

GetLevelThreeRateTypeOk returns a tuple with the LevelThreeRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeRateType

`func (o *TaxCode) SetLevelThreeRateType(v string)`

SetLevelThreeRateType sets LevelThreeRateType field to given value.

### HasLevelThreeRateType

`func (o *TaxCode) HasLevelThreeRateType() bool`

HasLevelThreeRateType returns a boolean if a field has been set.

### SetLevelThreeRateTypeNil

`func (o *TaxCode) SetLevelThreeRateTypeNil(b bool)`

 SetLevelThreeRateTypeNil sets the value for LevelThreeRateType to be an explicit nil

### UnsetLevelThreeRateType
`func (o *TaxCode) UnsetLevelThreeRateType()`

UnsetLevelThreeRateType ensures that no value is present for LevelThreeRateType, not even an explicit nil
### GetLevelThreeTaxableMax

`func (o *TaxCode) GetLevelThreeTaxableMax() float64`

GetLevelThreeTaxableMax returns the LevelThreeTaxableMax field if non-nil, zero value otherwise.

### GetLevelThreeTaxableMaxOk

`func (o *TaxCode) GetLevelThreeTaxableMaxOk() (*float64, bool)`

GetLevelThreeTaxableMaxOk returns a tuple with the LevelThreeTaxableMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeTaxableMax

`func (o *TaxCode) SetLevelThreeTaxableMax(v float64)`

SetLevelThreeTaxableMax sets LevelThreeTaxableMax field to given value.

### HasLevelThreeTaxableMax

`func (o *TaxCode) HasLevelThreeTaxableMax() bool`

HasLevelThreeTaxableMax returns a boolean if a field has been set.

### SetLevelThreeTaxableMaxNil

`func (o *TaxCode) SetLevelThreeTaxableMaxNil(b bool)`

 SetLevelThreeTaxableMaxNil sets the value for LevelThreeTaxableMax to be an explicit nil

### UnsetLevelThreeTaxableMax
`func (o *TaxCode) UnsetLevelThreeTaxableMax()`

UnsetLevelThreeTaxableMax ensures that no value is present for LevelThreeTaxableMax, not even an explicit nil
### GetLevelThreeCaption

`func (o *TaxCode) GetLevelThreeCaption() string`

GetLevelThreeCaption returns the LevelThreeCaption field if non-nil, zero value otherwise.

### GetLevelThreeCaptionOk

`func (o *TaxCode) GetLevelThreeCaptionOk() (*string, bool)`

GetLevelThreeCaptionOk returns a tuple with the LevelThreeCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeCaption

`func (o *TaxCode) SetLevelThreeCaption(v string)`

SetLevelThreeCaption sets LevelThreeCaption field to given value.

### HasLevelThreeCaption

`func (o *TaxCode) HasLevelThreeCaption() bool`

HasLevelThreeCaption returns a boolean if a field has been set.

### GetLevelThreeTaxCodeXref

`func (o *TaxCode) GetLevelThreeTaxCodeXref() string`

GetLevelThreeTaxCodeXref returns the LevelThreeTaxCodeXref field if non-nil, zero value otherwise.

### GetLevelThreeTaxCodeXrefOk

`func (o *TaxCode) GetLevelThreeTaxCodeXrefOk() (*string, bool)`

GetLevelThreeTaxCodeXrefOk returns a tuple with the LevelThreeTaxCodeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeTaxCodeXref

`func (o *TaxCode) SetLevelThreeTaxCodeXref(v string)`

SetLevelThreeTaxCodeXref sets LevelThreeTaxCodeXref field to given value.

### HasLevelThreeTaxCodeXref

`func (o *TaxCode) HasLevelThreeTaxCodeXref() bool`

HasLevelThreeTaxCodeXref returns a boolean if a field has been set.

### GetLevelThreeAgencyXref

`func (o *TaxCode) GetLevelThreeAgencyXref() string`

GetLevelThreeAgencyXref returns the LevelThreeAgencyXref field if non-nil, zero value otherwise.

### GetLevelThreeAgencyXrefOk

`func (o *TaxCode) GetLevelThreeAgencyXrefOk() (*string, bool)`

GetLevelThreeAgencyXrefOk returns a tuple with the LevelThreeAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeAgencyXref

`func (o *TaxCode) SetLevelThreeAgencyXref(v string)`

SetLevelThreeAgencyXref sets LevelThreeAgencyXref field to given value.

### HasLevelThreeAgencyXref

`func (o *TaxCode) HasLevelThreeAgencyXref() bool`

HasLevelThreeAgencyXref returns a boolean if a field has been set.

### GetLevelThreeServicesFlag

`func (o *TaxCode) GetLevelThreeServicesFlag() bool`

GetLevelThreeServicesFlag returns the LevelThreeServicesFlag field if non-nil, zero value otherwise.

### GetLevelThreeServicesFlagOk

`func (o *TaxCode) GetLevelThreeServicesFlagOk() (*bool, bool)`

GetLevelThreeServicesFlagOk returns a tuple with the LevelThreeServicesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeServicesFlag

`func (o *TaxCode) SetLevelThreeServicesFlag(v bool)`

SetLevelThreeServicesFlag sets LevelThreeServicesFlag field to given value.

### HasLevelThreeServicesFlag

`func (o *TaxCode) HasLevelThreeServicesFlag() bool`

HasLevelThreeServicesFlag returns a boolean if a field has been set.

### SetLevelThreeServicesFlagNil

`func (o *TaxCode) SetLevelThreeServicesFlagNil(b bool)`

 SetLevelThreeServicesFlagNil sets the value for LevelThreeServicesFlag to be an explicit nil

### UnsetLevelThreeServicesFlag
`func (o *TaxCode) UnsetLevelThreeServicesFlag()`

UnsetLevelThreeServicesFlag ensures that no value is present for LevelThreeServicesFlag, not even an explicit nil
### GetLevelThreeExpensesFlag

`func (o *TaxCode) GetLevelThreeExpensesFlag() bool`

GetLevelThreeExpensesFlag returns the LevelThreeExpensesFlag field if non-nil, zero value otherwise.

### GetLevelThreeExpensesFlagOk

`func (o *TaxCode) GetLevelThreeExpensesFlagOk() (*bool, bool)`

GetLevelThreeExpensesFlagOk returns a tuple with the LevelThreeExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeExpensesFlag

`func (o *TaxCode) SetLevelThreeExpensesFlag(v bool)`

SetLevelThreeExpensesFlag sets LevelThreeExpensesFlag field to given value.

### HasLevelThreeExpensesFlag

`func (o *TaxCode) HasLevelThreeExpensesFlag() bool`

HasLevelThreeExpensesFlag returns a boolean if a field has been set.

### SetLevelThreeExpensesFlagNil

`func (o *TaxCode) SetLevelThreeExpensesFlagNil(b bool)`

 SetLevelThreeExpensesFlagNil sets the value for LevelThreeExpensesFlag to be an explicit nil

### UnsetLevelThreeExpensesFlag
`func (o *TaxCode) UnsetLevelThreeExpensesFlag()`

UnsetLevelThreeExpensesFlag ensures that no value is present for LevelThreeExpensesFlag, not even an explicit nil
### GetLevelThreeProductsFlag

`func (o *TaxCode) GetLevelThreeProductsFlag() bool`

GetLevelThreeProductsFlag returns the LevelThreeProductsFlag field if non-nil, zero value otherwise.

### GetLevelThreeProductsFlagOk

`func (o *TaxCode) GetLevelThreeProductsFlagOk() (*bool, bool)`

GetLevelThreeProductsFlagOk returns a tuple with the LevelThreeProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeProductsFlag

`func (o *TaxCode) SetLevelThreeProductsFlag(v bool)`

SetLevelThreeProductsFlag sets LevelThreeProductsFlag field to given value.

### HasLevelThreeProductsFlag

`func (o *TaxCode) HasLevelThreeProductsFlag() bool`

HasLevelThreeProductsFlag returns a boolean if a field has been set.

### SetLevelThreeProductsFlagNil

`func (o *TaxCode) SetLevelThreeProductsFlagNil(b bool)`

 SetLevelThreeProductsFlagNil sets the value for LevelThreeProductsFlag to be an explicit nil

### UnsetLevelThreeProductsFlag
`func (o *TaxCode) UnsetLevelThreeProductsFlag()`

UnsetLevelThreeProductsFlag ensures that no value is present for LevelThreeProductsFlag, not even an explicit nil
### GetLevelThreeApplySingleUnitFlag

`func (o *TaxCode) GetLevelThreeApplySingleUnitFlag() bool`

GetLevelThreeApplySingleUnitFlag returns the LevelThreeApplySingleUnitFlag field if non-nil, zero value otherwise.

### GetLevelThreeApplySingleUnitFlagOk

`func (o *TaxCode) GetLevelThreeApplySingleUnitFlagOk() (*bool, bool)`

GetLevelThreeApplySingleUnitFlagOk returns a tuple with the LevelThreeApplySingleUnitFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeApplySingleUnitFlag

`func (o *TaxCode) SetLevelThreeApplySingleUnitFlag(v bool)`

SetLevelThreeApplySingleUnitFlag sets LevelThreeApplySingleUnitFlag field to given value.

### HasLevelThreeApplySingleUnitFlag

`func (o *TaxCode) HasLevelThreeApplySingleUnitFlag() bool`

HasLevelThreeApplySingleUnitFlag returns a boolean if a field has been set.

### SetLevelThreeApplySingleUnitFlagNil

`func (o *TaxCode) SetLevelThreeApplySingleUnitFlagNil(b bool)`

 SetLevelThreeApplySingleUnitFlagNil sets the value for LevelThreeApplySingleUnitFlag to be an explicit nil

### UnsetLevelThreeApplySingleUnitFlag
`func (o *TaxCode) UnsetLevelThreeApplySingleUnitFlag()`

UnsetLevelThreeApplySingleUnitFlag ensures that no value is present for LevelThreeApplySingleUnitFlag, not even an explicit nil
### GetLevelThreeApplySingleUnitMin

`func (o *TaxCode) GetLevelThreeApplySingleUnitMin() float64`

GetLevelThreeApplySingleUnitMin returns the LevelThreeApplySingleUnitMin field if non-nil, zero value otherwise.

### GetLevelThreeApplySingleUnitMinOk

`func (o *TaxCode) GetLevelThreeApplySingleUnitMinOk() (*float64, bool)`

GetLevelThreeApplySingleUnitMinOk returns a tuple with the LevelThreeApplySingleUnitMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeApplySingleUnitMin

`func (o *TaxCode) SetLevelThreeApplySingleUnitMin(v float64)`

SetLevelThreeApplySingleUnitMin sets LevelThreeApplySingleUnitMin field to given value.

### HasLevelThreeApplySingleUnitMin

`func (o *TaxCode) HasLevelThreeApplySingleUnitMin() bool`

HasLevelThreeApplySingleUnitMin returns a boolean if a field has been set.

### SetLevelThreeApplySingleUnitMinNil

`func (o *TaxCode) SetLevelThreeApplySingleUnitMinNil(b bool)`

 SetLevelThreeApplySingleUnitMinNil sets the value for LevelThreeApplySingleUnitMin to be an explicit nil

### UnsetLevelThreeApplySingleUnitMin
`func (o *TaxCode) UnsetLevelThreeApplySingleUnitMin()`

UnsetLevelThreeApplySingleUnitMin ensures that no value is present for LevelThreeApplySingleUnitMin, not even an explicit nil
### GetLevelThreeApplySingleUnitMax

`func (o *TaxCode) GetLevelThreeApplySingleUnitMax() float64`

GetLevelThreeApplySingleUnitMax returns the LevelThreeApplySingleUnitMax field if non-nil, zero value otherwise.

### GetLevelThreeApplySingleUnitMaxOk

`func (o *TaxCode) GetLevelThreeApplySingleUnitMaxOk() (*float64, bool)`

GetLevelThreeApplySingleUnitMaxOk returns a tuple with the LevelThreeApplySingleUnitMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeApplySingleUnitMax

`func (o *TaxCode) SetLevelThreeApplySingleUnitMax(v float64)`

SetLevelThreeApplySingleUnitMax sets LevelThreeApplySingleUnitMax field to given value.

### HasLevelThreeApplySingleUnitMax

`func (o *TaxCode) HasLevelThreeApplySingleUnitMax() bool`

HasLevelThreeApplySingleUnitMax returns a boolean if a field has been set.

### SetLevelThreeApplySingleUnitMaxNil

`func (o *TaxCode) SetLevelThreeApplySingleUnitMaxNil(b bool)`

 SetLevelThreeApplySingleUnitMaxNil sets the value for LevelThreeApplySingleUnitMax to be an explicit nil

### UnsetLevelThreeApplySingleUnitMax
`func (o *TaxCode) UnsetLevelThreeApplySingleUnitMax()`

UnsetLevelThreeApplySingleUnitMax ensures that no value is present for LevelThreeApplySingleUnitMax, not even an explicit nil
### GetLevelFourRate

`func (o *TaxCode) GetLevelFourRate() float64`

GetLevelFourRate returns the LevelFourRate field if non-nil, zero value otherwise.

### GetLevelFourRateOk

`func (o *TaxCode) GetLevelFourRateOk() (*float64, bool)`

GetLevelFourRateOk returns a tuple with the LevelFourRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourRate

`func (o *TaxCode) SetLevelFourRate(v float64)`

SetLevelFourRate sets LevelFourRate field to given value.

### HasLevelFourRate

`func (o *TaxCode) HasLevelFourRate() bool`

HasLevelFourRate returns a boolean if a field has been set.

### SetLevelFourRateNil

`func (o *TaxCode) SetLevelFourRateNil(b bool)`

 SetLevelFourRateNil sets the value for LevelFourRate to be an explicit nil

### UnsetLevelFourRate
`func (o *TaxCode) UnsetLevelFourRate()`

UnsetLevelFourRate ensures that no value is present for LevelFourRate, not even an explicit nil
### GetLevelFourRateType

`func (o *TaxCode) GetLevelFourRateType() string`

GetLevelFourRateType returns the LevelFourRateType field if non-nil, zero value otherwise.

### GetLevelFourRateTypeOk

`func (o *TaxCode) GetLevelFourRateTypeOk() (*string, bool)`

GetLevelFourRateTypeOk returns a tuple with the LevelFourRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourRateType

`func (o *TaxCode) SetLevelFourRateType(v string)`

SetLevelFourRateType sets LevelFourRateType field to given value.

### HasLevelFourRateType

`func (o *TaxCode) HasLevelFourRateType() bool`

HasLevelFourRateType returns a boolean if a field has been set.

### SetLevelFourRateTypeNil

`func (o *TaxCode) SetLevelFourRateTypeNil(b bool)`

 SetLevelFourRateTypeNil sets the value for LevelFourRateType to be an explicit nil

### UnsetLevelFourRateType
`func (o *TaxCode) UnsetLevelFourRateType()`

UnsetLevelFourRateType ensures that no value is present for LevelFourRateType, not even an explicit nil
### GetLevelFourTaxableMax

`func (o *TaxCode) GetLevelFourTaxableMax() float64`

GetLevelFourTaxableMax returns the LevelFourTaxableMax field if non-nil, zero value otherwise.

### GetLevelFourTaxableMaxOk

`func (o *TaxCode) GetLevelFourTaxableMaxOk() (*float64, bool)`

GetLevelFourTaxableMaxOk returns a tuple with the LevelFourTaxableMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourTaxableMax

`func (o *TaxCode) SetLevelFourTaxableMax(v float64)`

SetLevelFourTaxableMax sets LevelFourTaxableMax field to given value.

### HasLevelFourTaxableMax

`func (o *TaxCode) HasLevelFourTaxableMax() bool`

HasLevelFourTaxableMax returns a boolean if a field has been set.

### SetLevelFourTaxableMaxNil

`func (o *TaxCode) SetLevelFourTaxableMaxNil(b bool)`

 SetLevelFourTaxableMaxNil sets the value for LevelFourTaxableMax to be an explicit nil

### UnsetLevelFourTaxableMax
`func (o *TaxCode) UnsetLevelFourTaxableMax()`

UnsetLevelFourTaxableMax ensures that no value is present for LevelFourTaxableMax, not even an explicit nil
### GetLevelFourCaption

`func (o *TaxCode) GetLevelFourCaption() string`

GetLevelFourCaption returns the LevelFourCaption field if non-nil, zero value otherwise.

### GetLevelFourCaptionOk

`func (o *TaxCode) GetLevelFourCaptionOk() (*string, bool)`

GetLevelFourCaptionOk returns a tuple with the LevelFourCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourCaption

`func (o *TaxCode) SetLevelFourCaption(v string)`

SetLevelFourCaption sets LevelFourCaption field to given value.

### HasLevelFourCaption

`func (o *TaxCode) HasLevelFourCaption() bool`

HasLevelFourCaption returns a boolean if a field has been set.

### GetLevelFourTaxCodeXref

`func (o *TaxCode) GetLevelFourTaxCodeXref() string`

GetLevelFourTaxCodeXref returns the LevelFourTaxCodeXref field if non-nil, zero value otherwise.

### GetLevelFourTaxCodeXrefOk

`func (o *TaxCode) GetLevelFourTaxCodeXrefOk() (*string, bool)`

GetLevelFourTaxCodeXrefOk returns a tuple with the LevelFourTaxCodeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourTaxCodeXref

`func (o *TaxCode) SetLevelFourTaxCodeXref(v string)`

SetLevelFourTaxCodeXref sets LevelFourTaxCodeXref field to given value.

### HasLevelFourTaxCodeXref

`func (o *TaxCode) HasLevelFourTaxCodeXref() bool`

HasLevelFourTaxCodeXref returns a boolean if a field has been set.

### GetLevelFourAgencyXref

`func (o *TaxCode) GetLevelFourAgencyXref() string`

GetLevelFourAgencyXref returns the LevelFourAgencyXref field if non-nil, zero value otherwise.

### GetLevelFourAgencyXrefOk

`func (o *TaxCode) GetLevelFourAgencyXrefOk() (*string, bool)`

GetLevelFourAgencyXrefOk returns a tuple with the LevelFourAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourAgencyXref

`func (o *TaxCode) SetLevelFourAgencyXref(v string)`

SetLevelFourAgencyXref sets LevelFourAgencyXref field to given value.

### HasLevelFourAgencyXref

`func (o *TaxCode) HasLevelFourAgencyXref() bool`

HasLevelFourAgencyXref returns a boolean if a field has been set.

### GetLevelFourServicesFlag

`func (o *TaxCode) GetLevelFourServicesFlag() bool`

GetLevelFourServicesFlag returns the LevelFourServicesFlag field if non-nil, zero value otherwise.

### GetLevelFourServicesFlagOk

`func (o *TaxCode) GetLevelFourServicesFlagOk() (*bool, bool)`

GetLevelFourServicesFlagOk returns a tuple with the LevelFourServicesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourServicesFlag

`func (o *TaxCode) SetLevelFourServicesFlag(v bool)`

SetLevelFourServicesFlag sets LevelFourServicesFlag field to given value.

### HasLevelFourServicesFlag

`func (o *TaxCode) HasLevelFourServicesFlag() bool`

HasLevelFourServicesFlag returns a boolean if a field has been set.

### SetLevelFourServicesFlagNil

`func (o *TaxCode) SetLevelFourServicesFlagNil(b bool)`

 SetLevelFourServicesFlagNil sets the value for LevelFourServicesFlag to be an explicit nil

### UnsetLevelFourServicesFlag
`func (o *TaxCode) UnsetLevelFourServicesFlag()`

UnsetLevelFourServicesFlag ensures that no value is present for LevelFourServicesFlag, not even an explicit nil
### GetLevelFourExpensesFlag

`func (o *TaxCode) GetLevelFourExpensesFlag() bool`

GetLevelFourExpensesFlag returns the LevelFourExpensesFlag field if non-nil, zero value otherwise.

### GetLevelFourExpensesFlagOk

`func (o *TaxCode) GetLevelFourExpensesFlagOk() (*bool, bool)`

GetLevelFourExpensesFlagOk returns a tuple with the LevelFourExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourExpensesFlag

`func (o *TaxCode) SetLevelFourExpensesFlag(v bool)`

SetLevelFourExpensesFlag sets LevelFourExpensesFlag field to given value.

### HasLevelFourExpensesFlag

`func (o *TaxCode) HasLevelFourExpensesFlag() bool`

HasLevelFourExpensesFlag returns a boolean if a field has been set.

### SetLevelFourExpensesFlagNil

`func (o *TaxCode) SetLevelFourExpensesFlagNil(b bool)`

 SetLevelFourExpensesFlagNil sets the value for LevelFourExpensesFlag to be an explicit nil

### UnsetLevelFourExpensesFlag
`func (o *TaxCode) UnsetLevelFourExpensesFlag()`

UnsetLevelFourExpensesFlag ensures that no value is present for LevelFourExpensesFlag, not even an explicit nil
### GetLevelFourProductsFlag

`func (o *TaxCode) GetLevelFourProductsFlag() bool`

GetLevelFourProductsFlag returns the LevelFourProductsFlag field if non-nil, zero value otherwise.

### GetLevelFourProductsFlagOk

`func (o *TaxCode) GetLevelFourProductsFlagOk() (*bool, bool)`

GetLevelFourProductsFlagOk returns a tuple with the LevelFourProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourProductsFlag

`func (o *TaxCode) SetLevelFourProductsFlag(v bool)`

SetLevelFourProductsFlag sets LevelFourProductsFlag field to given value.

### HasLevelFourProductsFlag

`func (o *TaxCode) HasLevelFourProductsFlag() bool`

HasLevelFourProductsFlag returns a boolean if a field has been set.

### SetLevelFourProductsFlagNil

`func (o *TaxCode) SetLevelFourProductsFlagNil(b bool)`

 SetLevelFourProductsFlagNil sets the value for LevelFourProductsFlag to be an explicit nil

### UnsetLevelFourProductsFlag
`func (o *TaxCode) UnsetLevelFourProductsFlag()`

UnsetLevelFourProductsFlag ensures that no value is present for LevelFourProductsFlag, not even an explicit nil
### GetLevelFourApplySingleUnitFlag

`func (o *TaxCode) GetLevelFourApplySingleUnitFlag() bool`

GetLevelFourApplySingleUnitFlag returns the LevelFourApplySingleUnitFlag field if non-nil, zero value otherwise.

### GetLevelFourApplySingleUnitFlagOk

`func (o *TaxCode) GetLevelFourApplySingleUnitFlagOk() (*bool, bool)`

GetLevelFourApplySingleUnitFlagOk returns a tuple with the LevelFourApplySingleUnitFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourApplySingleUnitFlag

`func (o *TaxCode) SetLevelFourApplySingleUnitFlag(v bool)`

SetLevelFourApplySingleUnitFlag sets LevelFourApplySingleUnitFlag field to given value.

### HasLevelFourApplySingleUnitFlag

`func (o *TaxCode) HasLevelFourApplySingleUnitFlag() bool`

HasLevelFourApplySingleUnitFlag returns a boolean if a field has been set.

### SetLevelFourApplySingleUnitFlagNil

`func (o *TaxCode) SetLevelFourApplySingleUnitFlagNil(b bool)`

 SetLevelFourApplySingleUnitFlagNil sets the value for LevelFourApplySingleUnitFlag to be an explicit nil

### UnsetLevelFourApplySingleUnitFlag
`func (o *TaxCode) UnsetLevelFourApplySingleUnitFlag()`

UnsetLevelFourApplySingleUnitFlag ensures that no value is present for LevelFourApplySingleUnitFlag, not even an explicit nil
### GetLevelFourApplySingleUnitMin

`func (o *TaxCode) GetLevelFourApplySingleUnitMin() float64`

GetLevelFourApplySingleUnitMin returns the LevelFourApplySingleUnitMin field if non-nil, zero value otherwise.

### GetLevelFourApplySingleUnitMinOk

`func (o *TaxCode) GetLevelFourApplySingleUnitMinOk() (*float64, bool)`

GetLevelFourApplySingleUnitMinOk returns a tuple with the LevelFourApplySingleUnitMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourApplySingleUnitMin

`func (o *TaxCode) SetLevelFourApplySingleUnitMin(v float64)`

SetLevelFourApplySingleUnitMin sets LevelFourApplySingleUnitMin field to given value.

### HasLevelFourApplySingleUnitMin

`func (o *TaxCode) HasLevelFourApplySingleUnitMin() bool`

HasLevelFourApplySingleUnitMin returns a boolean if a field has been set.

### SetLevelFourApplySingleUnitMinNil

`func (o *TaxCode) SetLevelFourApplySingleUnitMinNil(b bool)`

 SetLevelFourApplySingleUnitMinNil sets the value for LevelFourApplySingleUnitMin to be an explicit nil

### UnsetLevelFourApplySingleUnitMin
`func (o *TaxCode) UnsetLevelFourApplySingleUnitMin()`

UnsetLevelFourApplySingleUnitMin ensures that no value is present for LevelFourApplySingleUnitMin, not even an explicit nil
### GetLevelFourApplySingleUnitMax

`func (o *TaxCode) GetLevelFourApplySingleUnitMax() float64`

GetLevelFourApplySingleUnitMax returns the LevelFourApplySingleUnitMax field if non-nil, zero value otherwise.

### GetLevelFourApplySingleUnitMaxOk

`func (o *TaxCode) GetLevelFourApplySingleUnitMaxOk() (*float64, bool)`

GetLevelFourApplySingleUnitMaxOk returns a tuple with the LevelFourApplySingleUnitMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourApplySingleUnitMax

`func (o *TaxCode) SetLevelFourApplySingleUnitMax(v float64)`

SetLevelFourApplySingleUnitMax sets LevelFourApplySingleUnitMax field to given value.

### HasLevelFourApplySingleUnitMax

`func (o *TaxCode) HasLevelFourApplySingleUnitMax() bool`

HasLevelFourApplySingleUnitMax returns a boolean if a field has been set.

### SetLevelFourApplySingleUnitMaxNil

`func (o *TaxCode) SetLevelFourApplySingleUnitMaxNil(b bool)`

 SetLevelFourApplySingleUnitMaxNil sets the value for LevelFourApplySingleUnitMax to be an explicit nil

### UnsetLevelFourApplySingleUnitMax
`func (o *TaxCode) UnsetLevelFourApplySingleUnitMax()`

UnsetLevelFourApplySingleUnitMax ensures that no value is present for LevelFourApplySingleUnitMax, not even an explicit nil
### GetLevelFiveRate

`func (o *TaxCode) GetLevelFiveRate() float64`

GetLevelFiveRate returns the LevelFiveRate field if non-nil, zero value otherwise.

### GetLevelFiveRateOk

`func (o *TaxCode) GetLevelFiveRateOk() (*float64, bool)`

GetLevelFiveRateOk returns a tuple with the LevelFiveRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveRate

`func (o *TaxCode) SetLevelFiveRate(v float64)`

SetLevelFiveRate sets LevelFiveRate field to given value.

### HasLevelFiveRate

`func (o *TaxCode) HasLevelFiveRate() bool`

HasLevelFiveRate returns a boolean if a field has been set.

### SetLevelFiveRateNil

`func (o *TaxCode) SetLevelFiveRateNil(b bool)`

 SetLevelFiveRateNil sets the value for LevelFiveRate to be an explicit nil

### UnsetLevelFiveRate
`func (o *TaxCode) UnsetLevelFiveRate()`

UnsetLevelFiveRate ensures that no value is present for LevelFiveRate, not even an explicit nil
### GetLevelFiveRateType

`func (o *TaxCode) GetLevelFiveRateType() string`

GetLevelFiveRateType returns the LevelFiveRateType field if non-nil, zero value otherwise.

### GetLevelFiveRateTypeOk

`func (o *TaxCode) GetLevelFiveRateTypeOk() (*string, bool)`

GetLevelFiveRateTypeOk returns a tuple with the LevelFiveRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveRateType

`func (o *TaxCode) SetLevelFiveRateType(v string)`

SetLevelFiveRateType sets LevelFiveRateType field to given value.

### HasLevelFiveRateType

`func (o *TaxCode) HasLevelFiveRateType() bool`

HasLevelFiveRateType returns a boolean if a field has been set.

### SetLevelFiveRateTypeNil

`func (o *TaxCode) SetLevelFiveRateTypeNil(b bool)`

 SetLevelFiveRateTypeNil sets the value for LevelFiveRateType to be an explicit nil

### UnsetLevelFiveRateType
`func (o *TaxCode) UnsetLevelFiveRateType()`

UnsetLevelFiveRateType ensures that no value is present for LevelFiveRateType, not even an explicit nil
### GetLevelFiveTaxableMax

`func (o *TaxCode) GetLevelFiveTaxableMax() float64`

GetLevelFiveTaxableMax returns the LevelFiveTaxableMax field if non-nil, zero value otherwise.

### GetLevelFiveTaxableMaxOk

`func (o *TaxCode) GetLevelFiveTaxableMaxOk() (*float64, bool)`

GetLevelFiveTaxableMaxOk returns a tuple with the LevelFiveTaxableMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveTaxableMax

`func (o *TaxCode) SetLevelFiveTaxableMax(v float64)`

SetLevelFiveTaxableMax sets LevelFiveTaxableMax field to given value.

### HasLevelFiveTaxableMax

`func (o *TaxCode) HasLevelFiveTaxableMax() bool`

HasLevelFiveTaxableMax returns a boolean if a field has been set.

### SetLevelFiveTaxableMaxNil

`func (o *TaxCode) SetLevelFiveTaxableMaxNil(b bool)`

 SetLevelFiveTaxableMaxNil sets the value for LevelFiveTaxableMax to be an explicit nil

### UnsetLevelFiveTaxableMax
`func (o *TaxCode) UnsetLevelFiveTaxableMax()`

UnsetLevelFiveTaxableMax ensures that no value is present for LevelFiveTaxableMax, not even an explicit nil
### GetLevelFiveCaption

`func (o *TaxCode) GetLevelFiveCaption() string`

GetLevelFiveCaption returns the LevelFiveCaption field if non-nil, zero value otherwise.

### GetLevelFiveCaptionOk

`func (o *TaxCode) GetLevelFiveCaptionOk() (*string, bool)`

GetLevelFiveCaptionOk returns a tuple with the LevelFiveCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveCaption

`func (o *TaxCode) SetLevelFiveCaption(v string)`

SetLevelFiveCaption sets LevelFiveCaption field to given value.

### HasLevelFiveCaption

`func (o *TaxCode) HasLevelFiveCaption() bool`

HasLevelFiveCaption returns a boolean if a field has been set.

### GetLevelFiveTaxCodeXref

`func (o *TaxCode) GetLevelFiveTaxCodeXref() string`

GetLevelFiveTaxCodeXref returns the LevelFiveTaxCodeXref field if non-nil, zero value otherwise.

### GetLevelFiveTaxCodeXrefOk

`func (o *TaxCode) GetLevelFiveTaxCodeXrefOk() (*string, bool)`

GetLevelFiveTaxCodeXrefOk returns a tuple with the LevelFiveTaxCodeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveTaxCodeXref

`func (o *TaxCode) SetLevelFiveTaxCodeXref(v string)`

SetLevelFiveTaxCodeXref sets LevelFiveTaxCodeXref field to given value.

### HasLevelFiveTaxCodeXref

`func (o *TaxCode) HasLevelFiveTaxCodeXref() bool`

HasLevelFiveTaxCodeXref returns a boolean if a field has been set.

### GetLevelFiveAgencyXref

`func (o *TaxCode) GetLevelFiveAgencyXref() string`

GetLevelFiveAgencyXref returns the LevelFiveAgencyXref field if non-nil, zero value otherwise.

### GetLevelFiveAgencyXrefOk

`func (o *TaxCode) GetLevelFiveAgencyXrefOk() (*string, bool)`

GetLevelFiveAgencyXrefOk returns a tuple with the LevelFiveAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveAgencyXref

`func (o *TaxCode) SetLevelFiveAgencyXref(v string)`

SetLevelFiveAgencyXref sets LevelFiveAgencyXref field to given value.

### HasLevelFiveAgencyXref

`func (o *TaxCode) HasLevelFiveAgencyXref() bool`

HasLevelFiveAgencyXref returns a boolean if a field has been set.

### GetLevelFiveServicesFlag

`func (o *TaxCode) GetLevelFiveServicesFlag() bool`

GetLevelFiveServicesFlag returns the LevelFiveServicesFlag field if non-nil, zero value otherwise.

### GetLevelFiveServicesFlagOk

`func (o *TaxCode) GetLevelFiveServicesFlagOk() (*bool, bool)`

GetLevelFiveServicesFlagOk returns a tuple with the LevelFiveServicesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveServicesFlag

`func (o *TaxCode) SetLevelFiveServicesFlag(v bool)`

SetLevelFiveServicesFlag sets LevelFiveServicesFlag field to given value.

### HasLevelFiveServicesFlag

`func (o *TaxCode) HasLevelFiveServicesFlag() bool`

HasLevelFiveServicesFlag returns a boolean if a field has been set.

### SetLevelFiveServicesFlagNil

`func (o *TaxCode) SetLevelFiveServicesFlagNil(b bool)`

 SetLevelFiveServicesFlagNil sets the value for LevelFiveServicesFlag to be an explicit nil

### UnsetLevelFiveServicesFlag
`func (o *TaxCode) UnsetLevelFiveServicesFlag()`

UnsetLevelFiveServicesFlag ensures that no value is present for LevelFiveServicesFlag, not even an explicit nil
### GetLevelFiveExpensesFlag

`func (o *TaxCode) GetLevelFiveExpensesFlag() bool`

GetLevelFiveExpensesFlag returns the LevelFiveExpensesFlag field if non-nil, zero value otherwise.

### GetLevelFiveExpensesFlagOk

`func (o *TaxCode) GetLevelFiveExpensesFlagOk() (*bool, bool)`

GetLevelFiveExpensesFlagOk returns a tuple with the LevelFiveExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveExpensesFlag

`func (o *TaxCode) SetLevelFiveExpensesFlag(v bool)`

SetLevelFiveExpensesFlag sets LevelFiveExpensesFlag field to given value.

### HasLevelFiveExpensesFlag

`func (o *TaxCode) HasLevelFiveExpensesFlag() bool`

HasLevelFiveExpensesFlag returns a boolean if a field has been set.

### SetLevelFiveExpensesFlagNil

`func (o *TaxCode) SetLevelFiveExpensesFlagNil(b bool)`

 SetLevelFiveExpensesFlagNil sets the value for LevelFiveExpensesFlag to be an explicit nil

### UnsetLevelFiveExpensesFlag
`func (o *TaxCode) UnsetLevelFiveExpensesFlag()`

UnsetLevelFiveExpensesFlag ensures that no value is present for LevelFiveExpensesFlag, not even an explicit nil
### GetLevelFiveProductsFlag

`func (o *TaxCode) GetLevelFiveProductsFlag() bool`

GetLevelFiveProductsFlag returns the LevelFiveProductsFlag field if non-nil, zero value otherwise.

### GetLevelFiveProductsFlagOk

`func (o *TaxCode) GetLevelFiveProductsFlagOk() (*bool, bool)`

GetLevelFiveProductsFlagOk returns a tuple with the LevelFiveProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveProductsFlag

`func (o *TaxCode) SetLevelFiveProductsFlag(v bool)`

SetLevelFiveProductsFlag sets LevelFiveProductsFlag field to given value.

### HasLevelFiveProductsFlag

`func (o *TaxCode) HasLevelFiveProductsFlag() bool`

HasLevelFiveProductsFlag returns a boolean if a field has been set.

### SetLevelFiveProductsFlagNil

`func (o *TaxCode) SetLevelFiveProductsFlagNil(b bool)`

 SetLevelFiveProductsFlagNil sets the value for LevelFiveProductsFlag to be an explicit nil

### UnsetLevelFiveProductsFlag
`func (o *TaxCode) UnsetLevelFiveProductsFlag()`

UnsetLevelFiveProductsFlag ensures that no value is present for LevelFiveProductsFlag, not even an explicit nil
### GetLevelFiveApplySingleUnitFlag

`func (o *TaxCode) GetLevelFiveApplySingleUnitFlag() bool`

GetLevelFiveApplySingleUnitFlag returns the LevelFiveApplySingleUnitFlag field if non-nil, zero value otherwise.

### GetLevelFiveApplySingleUnitFlagOk

`func (o *TaxCode) GetLevelFiveApplySingleUnitFlagOk() (*bool, bool)`

GetLevelFiveApplySingleUnitFlagOk returns a tuple with the LevelFiveApplySingleUnitFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveApplySingleUnitFlag

`func (o *TaxCode) SetLevelFiveApplySingleUnitFlag(v bool)`

SetLevelFiveApplySingleUnitFlag sets LevelFiveApplySingleUnitFlag field to given value.

### HasLevelFiveApplySingleUnitFlag

`func (o *TaxCode) HasLevelFiveApplySingleUnitFlag() bool`

HasLevelFiveApplySingleUnitFlag returns a boolean if a field has been set.

### SetLevelFiveApplySingleUnitFlagNil

`func (o *TaxCode) SetLevelFiveApplySingleUnitFlagNil(b bool)`

 SetLevelFiveApplySingleUnitFlagNil sets the value for LevelFiveApplySingleUnitFlag to be an explicit nil

### UnsetLevelFiveApplySingleUnitFlag
`func (o *TaxCode) UnsetLevelFiveApplySingleUnitFlag()`

UnsetLevelFiveApplySingleUnitFlag ensures that no value is present for LevelFiveApplySingleUnitFlag, not even an explicit nil
### GetLevelFiveApplySingleUnitMin

`func (o *TaxCode) GetLevelFiveApplySingleUnitMin() float64`

GetLevelFiveApplySingleUnitMin returns the LevelFiveApplySingleUnitMin field if non-nil, zero value otherwise.

### GetLevelFiveApplySingleUnitMinOk

`func (o *TaxCode) GetLevelFiveApplySingleUnitMinOk() (*float64, bool)`

GetLevelFiveApplySingleUnitMinOk returns a tuple with the LevelFiveApplySingleUnitMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveApplySingleUnitMin

`func (o *TaxCode) SetLevelFiveApplySingleUnitMin(v float64)`

SetLevelFiveApplySingleUnitMin sets LevelFiveApplySingleUnitMin field to given value.

### HasLevelFiveApplySingleUnitMin

`func (o *TaxCode) HasLevelFiveApplySingleUnitMin() bool`

HasLevelFiveApplySingleUnitMin returns a boolean if a field has been set.

### SetLevelFiveApplySingleUnitMinNil

`func (o *TaxCode) SetLevelFiveApplySingleUnitMinNil(b bool)`

 SetLevelFiveApplySingleUnitMinNil sets the value for LevelFiveApplySingleUnitMin to be an explicit nil

### UnsetLevelFiveApplySingleUnitMin
`func (o *TaxCode) UnsetLevelFiveApplySingleUnitMin()`

UnsetLevelFiveApplySingleUnitMin ensures that no value is present for LevelFiveApplySingleUnitMin, not even an explicit nil
### GetLevelFiveApplySingleUnitMax

`func (o *TaxCode) GetLevelFiveApplySingleUnitMax() float64`

GetLevelFiveApplySingleUnitMax returns the LevelFiveApplySingleUnitMax field if non-nil, zero value otherwise.

### GetLevelFiveApplySingleUnitMaxOk

`func (o *TaxCode) GetLevelFiveApplySingleUnitMaxOk() (*float64, bool)`

GetLevelFiveApplySingleUnitMaxOk returns a tuple with the LevelFiveApplySingleUnitMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveApplySingleUnitMax

`func (o *TaxCode) SetLevelFiveApplySingleUnitMax(v float64)`

SetLevelFiveApplySingleUnitMax sets LevelFiveApplySingleUnitMax field to given value.

### HasLevelFiveApplySingleUnitMax

`func (o *TaxCode) HasLevelFiveApplySingleUnitMax() bool`

HasLevelFiveApplySingleUnitMax returns a boolean if a field has been set.

### SetLevelFiveApplySingleUnitMaxNil

`func (o *TaxCode) SetLevelFiveApplySingleUnitMaxNil(b bool)`

 SetLevelFiveApplySingleUnitMaxNil sets the value for LevelFiveApplySingleUnitMax to be an explicit nil

### UnsetLevelFiveApplySingleUnitMax
`func (o *TaxCode) UnsetLevelFiveApplySingleUnitMax()`

UnsetLevelFiveApplySingleUnitMax ensures that no value is present for LevelFiveApplySingleUnitMax, not even an explicit nil
### GetLevelSixRate

`func (o *TaxCode) GetLevelSixRate() float64`

GetLevelSixRate returns the LevelSixRate field if non-nil, zero value otherwise.

### GetLevelSixRateOk

`func (o *TaxCode) GetLevelSixRateOk() (*float64, bool)`

GetLevelSixRateOk returns a tuple with the LevelSixRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixRate

`func (o *TaxCode) SetLevelSixRate(v float64)`

SetLevelSixRate sets LevelSixRate field to given value.

### HasLevelSixRate

`func (o *TaxCode) HasLevelSixRate() bool`

HasLevelSixRate returns a boolean if a field has been set.

### SetLevelSixRateNil

`func (o *TaxCode) SetLevelSixRateNil(b bool)`

 SetLevelSixRateNil sets the value for LevelSixRate to be an explicit nil

### UnsetLevelSixRate
`func (o *TaxCode) UnsetLevelSixRate()`

UnsetLevelSixRate ensures that no value is present for LevelSixRate, not even an explicit nil
### GetLevelSixRateType

`func (o *TaxCode) GetLevelSixRateType() string`

GetLevelSixRateType returns the LevelSixRateType field if non-nil, zero value otherwise.

### GetLevelSixRateTypeOk

`func (o *TaxCode) GetLevelSixRateTypeOk() (*string, bool)`

GetLevelSixRateTypeOk returns a tuple with the LevelSixRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixRateType

`func (o *TaxCode) SetLevelSixRateType(v string)`

SetLevelSixRateType sets LevelSixRateType field to given value.

### HasLevelSixRateType

`func (o *TaxCode) HasLevelSixRateType() bool`

HasLevelSixRateType returns a boolean if a field has been set.

### SetLevelSixRateTypeNil

`func (o *TaxCode) SetLevelSixRateTypeNil(b bool)`

 SetLevelSixRateTypeNil sets the value for LevelSixRateType to be an explicit nil

### UnsetLevelSixRateType
`func (o *TaxCode) UnsetLevelSixRateType()`

UnsetLevelSixRateType ensures that no value is present for LevelSixRateType, not even an explicit nil
### GetLevelSixTaxableMax

`func (o *TaxCode) GetLevelSixTaxableMax() float64`

GetLevelSixTaxableMax returns the LevelSixTaxableMax field if non-nil, zero value otherwise.

### GetLevelSixTaxableMaxOk

`func (o *TaxCode) GetLevelSixTaxableMaxOk() (*float64, bool)`

GetLevelSixTaxableMaxOk returns a tuple with the LevelSixTaxableMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixTaxableMax

`func (o *TaxCode) SetLevelSixTaxableMax(v float64)`

SetLevelSixTaxableMax sets LevelSixTaxableMax field to given value.

### HasLevelSixTaxableMax

`func (o *TaxCode) HasLevelSixTaxableMax() bool`

HasLevelSixTaxableMax returns a boolean if a field has been set.

### SetLevelSixTaxableMaxNil

`func (o *TaxCode) SetLevelSixTaxableMaxNil(b bool)`

 SetLevelSixTaxableMaxNil sets the value for LevelSixTaxableMax to be an explicit nil

### UnsetLevelSixTaxableMax
`func (o *TaxCode) UnsetLevelSixTaxableMax()`

UnsetLevelSixTaxableMax ensures that no value is present for LevelSixTaxableMax, not even an explicit nil
### GetLevelSixCaption

`func (o *TaxCode) GetLevelSixCaption() string`

GetLevelSixCaption returns the LevelSixCaption field if non-nil, zero value otherwise.

### GetLevelSixCaptionOk

`func (o *TaxCode) GetLevelSixCaptionOk() (*string, bool)`

GetLevelSixCaptionOk returns a tuple with the LevelSixCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixCaption

`func (o *TaxCode) SetLevelSixCaption(v string)`

SetLevelSixCaption sets LevelSixCaption field to given value.

### HasLevelSixCaption

`func (o *TaxCode) HasLevelSixCaption() bool`

HasLevelSixCaption returns a boolean if a field has been set.

### GetLevelSixTaxCodeXref

`func (o *TaxCode) GetLevelSixTaxCodeXref() string`

GetLevelSixTaxCodeXref returns the LevelSixTaxCodeXref field if non-nil, zero value otherwise.

### GetLevelSixTaxCodeXrefOk

`func (o *TaxCode) GetLevelSixTaxCodeXrefOk() (*string, bool)`

GetLevelSixTaxCodeXrefOk returns a tuple with the LevelSixTaxCodeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixTaxCodeXref

`func (o *TaxCode) SetLevelSixTaxCodeXref(v string)`

SetLevelSixTaxCodeXref sets LevelSixTaxCodeXref field to given value.

### HasLevelSixTaxCodeXref

`func (o *TaxCode) HasLevelSixTaxCodeXref() bool`

HasLevelSixTaxCodeXref returns a boolean if a field has been set.

### GetLevelSixAgencyXref

`func (o *TaxCode) GetLevelSixAgencyXref() string`

GetLevelSixAgencyXref returns the LevelSixAgencyXref field if non-nil, zero value otherwise.

### GetLevelSixAgencyXrefOk

`func (o *TaxCode) GetLevelSixAgencyXrefOk() (*string, bool)`

GetLevelSixAgencyXrefOk returns a tuple with the LevelSixAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixAgencyXref

`func (o *TaxCode) SetLevelSixAgencyXref(v string)`

SetLevelSixAgencyXref sets LevelSixAgencyXref field to given value.

### HasLevelSixAgencyXref

`func (o *TaxCode) HasLevelSixAgencyXref() bool`

HasLevelSixAgencyXref returns a boolean if a field has been set.

### GetLevelSixServicesFlag

`func (o *TaxCode) GetLevelSixServicesFlag() bool`

GetLevelSixServicesFlag returns the LevelSixServicesFlag field if non-nil, zero value otherwise.

### GetLevelSixServicesFlagOk

`func (o *TaxCode) GetLevelSixServicesFlagOk() (*bool, bool)`

GetLevelSixServicesFlagOk returns a tuple with the LevelSixServicesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixServicesFlag

`func (o *TaxCode) SetLevelSixServicesFlag(v bool)`

SetLevelSixServicesFlag sets LevelSixServicesFlag field to given value.

### HasLevelSixServicesFlag

`func (o *TaxCode) HasLevelSixServicesFlag() bool`

HasLevelSixServicesFlag returns a boolean if a field has been set.

### SetLevelSixServicesFlagNil

`func (o *TaxCode) SetLevelSixServicesFlagNil(b bool)`

 SetLevelSixServicesFlagNil sets the value for LevelSixServicesFlag to be an explicit nil

### UnsetLevelSixServicesFlag
`func (o *TaxCode) UnsetLevelSixServicesFlag()`

UnsetLevelSixServicesFlag ensures that no value is present for LevelSixServicesFlag, not even an explicit nil
### GetLevelSixExpensesFlag

`func (o *TaxCode) GetLevelSixExpensesFlag() bool`

GetLevelSixExpensesFlag returns the LevelSixExpensesFlag field if non-nil, zero value otherwise.

### GetLevelSixExpensesFlagOk

`func (o *TaxCode) GetLevelSixExpensesFlagOk() (*bool, bool)`

GetLevelSixExpensesFlagOk returns a tuple with the LevelSixExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixExpensesFlag

`func (o *TaxCode) SetLevelSixExpensesFlag(v bool)`

SetLevelSixExpensesFlag sets LevelSixExpensesFlag field to given value.

### HasLevelSixExpensesFlag

`func (o *TaxCode) HasLevelSixExpensesFlag() bool`

HasLevelSixExpensesFlag returns a boolean if a field has been set.

### SetLevelSixExpensesFlagNil

`func (o *TaxCode) SetLevelSixExpensesFlagNil(b bool)`

 SetLevelSixExpensesFlagNil sets the value for LevelSixExpensesFlag to be an explicit nil

### UnsetLevelSixExpensesFlag
`func (o *TaxCode) UnsetLevelSixExpensesFlag()`

UnsetLevelSixExpensesFlag ensures that no value is present for LevelSixExpensesFlag, not even an explicit nil
### GetLevelSixProductsFlag

`func (o *TaxCode) GetLevelSixProductsFlag() bool`

GetLevelSixProductsFlag returns the LevelSixProductsFlag field if non-nil, zero value otherwise.

### GetLevelSixProductsFlagOk

`func (o *TaxCode) GetLevelSixProductsFlagOk() (*bool, bool)`

GetLevelSixProductsFlagOk returns a tuple with the LevelSixProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixProductsFlag

`func (o *TaxCode) SetLevelSixProductsFlag(v bool)`

SetLevelSixProductsFlag sets LevelSixProductsFlag field to given value.

### HasLevelSixProductsFlag

`func (o *TaxCode) HasLevelSixProductsFlag() bool`

HasLevelSixProductsFlag returns a boolean if a field has been set.

### SetLevelSixProductsFlagNil

`func (o *TaxCode) SetLevelSixProductsFlagNil(b bool)`

 SetLevelSixProductsFlagNil sets the value for LevelSixProductsFlag to be an explicit nil

### UnsetLevelSixProductsFlag
`func (o *TaxCode) UnsetLevelSixProductsFlag()`

UnsetLevelSixProductsFlag ensures that no value is present for LevelSixProductsFlag, not even an explicit nil
### GetLevelSixApplySingleUnitFlag

`func (o *TaxCode) GetLevelSixApplySingleUnitFlag() bool`

GetLevelSixApplySingleUnitFlag returns the LevelSixApplySingleUnitFlag field if non-nil, zero value otherwise.

### GetLevelSixApplySingleUnitFlagOk

`func (o *TaxCode) GetLevelSixApplySingleUnitFlagOk() (*bool, bool)`

GetLevelSixApplySingleUnitFlagOk returns a tuple with the LevelSixApplySingleUnitFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixApplySingleUnitFlag

`func (o *TaxCode) SetLevelSixApplySingleUnitFlag(v bool)`

SetLevelSixApplySingleUnitFlag sets LevelSixApplySingleUnitFlag field to given value.

### HasLevelSixApplySingleUnitFlag

`func (o *TaxCode) HasLevelSixApplySingleUnitFlag() bool`

HasLevelSixApplySingleUnitFlag returns a boolean if a field has been set.

### SetLevelSixApplySingleUnitFlagNil

`func (o *TaxCode) SetLevelSixApplySingleUnitFlagNil(b bool)`

 SetLevelSixApplySingleUnitFlagNil sets the value for LevelSixApplySingleUnitFlag to be an explicit nil

### UnsetLevelSixApplySingleUnitFlag
`func (o *TaxCode) UnsetLevelSixApplySingleUnitFlag()`

UnsetLevelSixApplySingleUnitFlag ensures that no value is present for LevelSixApplySingleUnitFlag, not even an explicit nil
### GetLevelSixApplySingleUnitMin

`func (o *TaxCode) GetLevelSixApplySingleUnitMin() float64`

GetLevelSixApplySingleUnitMin returns the LevelSixApplySingleUnitMin field if non-nil, zero value otherwise.

### GetLevelSixApplySingleUnitMinOk

`func (o *TaxCode) GetLevelSixApplySingleUnitMinOk() (*float64, bool)`

GetLevelSixApplySingleUnitMinOk returns a tuple with the LevelSixApplySingleUnitMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixApplySingleUnitMin

`func (o *TaxCode) SetLevelSixApplySingleUnitMin(v float64)`

SetLevelSixApplySingleUnitMin sets LevelSixApplySingleUnitMin field to given value.

### HasLevelSixApplySingleUnitMin

`func (o *TaxCode) HasLevelSixApplySingleUnitMin() bool`

HasLevelSixApplySingleUnitMin returns a boolean if a field has been set.

### SetLevelSixApplySingleUnitMinNil

`func (o *TaxCode) SetLevelSixApplySingleUnitMinNil(b bool)`

 SetLevelSixApplySingleUnitMinNil sets the value for LevelSixApplySingleUnitMin to be an explicit nil

### UnsetLevelSixApplySingleUnitMin
`func (o *TaxCode) UnsetLevelSixApplySingleUnitMin()`

UnsetLevelSixApplySingleUnitMin ensures that no value is present for LevelSixApplySingleUnitMin, not even an explicit nil
### GetLevelSixApplySingleUnitMax

`func (o *TaxCode) GetLevelSixApplySingleUnitMax() float64`

GetLevelSixApplySingleUnitMax returns the LevelSixApplySingleUnitMax field if non-nil, zero value otherwise.

### GetLevelSixApplySingleUnitMaxOk

`func (o *TaxCode) GetLevelSixApplySingleUnitMaxOk() (*float64, bool)`

GetLevelSixApplySingleUnitMaxOk returns a tuple with the LevelSixApplySingleUnitMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixApplySingleUnitMax

`func (o *TaxCode) SetLevelSixApplySingleUnitMax(v float64)`

SetLevelSixApplySingleUnitMax sets LevelSixApplySingleUnitMax field to given value.

### HasLevelSixApplySingleUnitMax

`func (o *TaxCode) HasLevelSixApplySingleUnitMax() bool`

HasLevelSixApplySingleUnitMax returns a boolean if a field has been set.

### SetLevelSixApplySingleUnitMaxNil

`func (o *TaxCode) SetLevelSixApplySingleUnitMaxNil(b bool)`

 SetLevelSixApplySingleUnitMaxNil sets the value for LevelSixApplySingleUnitMax to be an explicit nil

### UnsetLevelSixApplySingleUnitMax
`func (o *TaxCode) UnsetLevelSixApplySingleUnitMax()`

UnsetLevelSixApplySingleUnitMax ensures that no value is present for LevelSixApplySingleUnitMax, not even an explicit nil
### GetWorkRoleIds

`func (o *TaxCode) GetWorkRoleIds() []int32`

GetWorkRoleIds returns the WorkRoleIds field if non-nil, zero value otherwise.

### GetWorkRoleIdsOk

`func (o *TaxCode) GetWorkRoleIdsOk() (*[]int32, bool)`

GetWorkRoleIdsOk returns a tuple with the WorkRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRoleIds

`func (o *TaxCode) SetWorkRoleIds(v []int32)`

SetWorkRoleIds sets WorkRoleIds field to given value.

### HasWorkRoleIds

`func (o *TaxCode) HasWorkRoleIds() bool`

HasWorkRoleIds returns a boolean if a field has been set.

### GetAddAllWorkRoles

`func (o *TaxCode) GetAddAllWorkRoles() bool`

GetAddAllWorkRoles returns the AddAllWorkRoles field if non-nil, zero value otherwise.

### GetAddAllWorkRolesOk

`func (o *TaxCode) GetAddAllWorkRolesOk() (*bool, bool)`

GetAddAllWorkRolesOk returns a tuple with the AddAllWorkRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllWorkRoles

`func (o *TaxCode) SetAddAllWorkRoles(v bool)`

SetAddAllWorkRoles sets AddAllWorkRoles field to given value.

### HasAddAllWorkRoles

`func (o *TaxCode) HasAddAllWorkRoles() bool`

HasAddAllWorkRoles returns a boolean if a field has been set.

### SetAddAllWorkRolesNil

`func (o *TaxCode) SetAddAllWorkRolesNil(b bool)`

 SetAddAllWorkRolesNil sets the value for AddAllWorkRoles to be an explicit nil

### UnsetAddAllWorkRoles
`func (o *TaxCode) UnsetAddAllWorkRoles()`

UnsetAddAllWorkRoles ensures that no value is present for AddAllWorkRoles, not even an explicit nil
### GetRemoveAllWorkRoles

`func (o *TaxCode) GetRemoveAllWorkRoles() bool`

GetRemoveAllWorkRoles returns the RemoveAllWorkRoles field if non-nil, zero value otherwise.

### GetRemoveAllWorkRolesOk

`func (o *TaxCode) GetRemoveAllWorkRolesOk() (*bool, bool)`

GetRemoveAllWorkRolesOk returns a tuple with the RemoveAllWorkRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllWorkRoles

`func (o *TaxCode) SetRemoveAllWorkRoles(v bool)`

SetRemoveAllWorkRoles sets RemoveAllWorkRoles field to given value.

### HasRemoveAllWorkRoles

`func (o *TaxCode) HasRemoveAllWorkRoles() bool`

HasRemoveAllWorkRoles returns a boolean if a field has been set.

### SetRemoveAllWorkRolesNil

`func (o *TaxCode) SetRemoveAllWorkRolesNil(b bool)`

 SetRemoveAllWorkRolesNil sets the value for RemoveAllWorkRoles to be an explicit nil

### UnsetRemoveAllWorkRoles
`func (o *TaxCode) UnsetRemoveAllWorkRoles()`

UnsetRemoveAllWorkRoles ensures that no value is present for RemoveAllWorkRoles, not even an explicit nil
### GetExpenseTypeIds

`func (o *TaxCode) GetExpenseTypeIds() []int32`

GetExpenseTypeIds returns the ExpenseTypeIds field if non-nil, zero value otherwise.

### GetExpenseTypeIdsOk

`func (o *TaxCode) GetExpenseTypeIdsOk() (*[]int32, bool)`

GetExpenseTypeIdsOk returns a tuple with the ExpenseTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseTypeIds

`func (o *TaxCode) SetExpenseTypeIds(v []int32)`

SetExpenseTypeIds sets ExpenseTypeIds field to given value.

### HasExpenseTypeIds

`func (o *TaxCode) HasExpenseTypeIds() bool`

HasExpenseTypeIds returns a boolean if a field has been set.

### GetAddAllExpenseTypes

`func (o *TaxCode) GetAddAllExpenseTypes() bool`

GetAddAllExpenseTypes returns the AddAllExpenseTypes field if non-nil, zero value otherwise.

### GetAddAllExpenseTypesOk

`func (o *TaxCode) GetAddAllExpenseTypesOk() (*bool, bool)`

GetAddAllExpenseTypesOk returns a tuple with the AddAllExpenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllExpenseTypes

`func (o *TaxCode) SetAddAllExpenseTypes(v bool)`

SetAddAllExpenseTypes sets AddAllExpenseTypes field to given value.

### HasAddAllExpenseTypes

`func (o *TaxCode) HasAddAllExpenseTypes() bool`

HasAddAllExpenseTypes returns a boolean if a field has been set.

### SetAddAllExpenseTypesNil

`func (o *TaxCode) SetAddAllExpenseTypesNil(b bool)`

 SetAddAllExpenseTypesNil sets the value for AddAllExpenseTypes to be an explicit nil

### UnsetAddAllExpenseTypes
`func (o *TaxCode) UnsetAddAllExpenseTypes()`

UnsetAddAllExpenseTypes ensures that no value is present for AddAllExpenseTypes, not even an explicit nil
### GetRemoveAllExpenseTypes

`func (o *TaxCode) GetRemoveAllExpenseTypes() bool`

GetRemoveAllExpenseTypes returns the RemoveAllExpenseTypes field if non-nil, zero value otherwise.

### GetRemoveAllExpenseTypesOk

`func (o *TaxCode) GetRemoveAllExpenseTypesOk() (*bool, bool)`

GetRemoveAllExpenseTypesOk returns a tuple with the RemoveAllExpenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllExpenseTypes

`func (o *TaxCode) SetRemoveAllExpenseTypes(v bool)`

SetRemoveAllExpenseTypes sets RemoveAllExpenseTypes field to given value.

### HasRemoveAllExpenseTypes

`func (o *TaxCode) HasRemoveAllExpenseTypes() bool`

HasRemoveAllExpenseTypes returns a boolean if a field has been set.

### SetRemoveAllExpenseTypesNil

`func (o *TaxCode) SetRemoveAllExpenseTypesNil(b bool)`

 SetRemoveAllExpenseTypesNil sets the value for RemoveAllExpenseTypes to be an explicit nil

### UnsetRemoveAllExpenseTypes
`func (o *TaxCode) UnsetRemoveAllExpenseTypes()`

UnsetRemoveAllExpenseTypes ensures that no value is present for RemoveAllExpenseTypes, not even an explicit nil
### GetProductTypeIds

`func (o *TaxCode) GetProductTypeIds() []int32`

GetProductTypeIds returns the ProductTypeIds field if non-nil, zero value otherwise.

### GetProductTypeIdsOk

`func (o *TaxCode) GetProductTypeIdsOk() (*[]int32, bool)`

GetProductTypeIdsOk returns a tuple with the ProductTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeIds

`func (o *TaxCode) SetProductTypeIds(v []int32)`

SetProductTypeIds sets ProductTypeIds field to given value.

### HasProductTypeIds

`func (o *TaxCode) HasProductTypeIds() bool`

HasProductTypeIds returns a boolean if a field has been set.

### GetAddAllProductTypes

`func (o *TaxCode) GetAddAllProductTypes() bool`

GetAddAllProductTypes returns the AddAllProductTypes field if non-nil, zero value otherwise.

### GetAddAllProductTypesOk

`func (o *TaxCode) GetAddAllProductTypesOk() (*bool, bool)`

GetAddAllProductTypesOk returns a tuple with the AddAllProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllProductTypes

`func (o *TaxCode) SetAddAllProductTypes(v bool)`

SetAddAllProductTypes sets AddAllProductTypes field to given value.

### HasAddAllProductTypes

`func (o *TaxCode) HasAddAllProductTypes() bool`

HasAddAllProductTypes returns a boolean if a field has been set.

### SetAddAllProductTypesNil

`func (o *TaxCode) SetAddAllProductTypesNil(b bool)`

 SetAddAllProductTypesNil sets the value for AddAllProductTypes to be an explicit nil

### UnsetAddAllProductTypes
`func (o *TaxCode) UnsetAddAllProductTypes()`

UnsetAddAllProductTypes ensures that no value is present for AddAllProductTypes, not even an explicit nil
### GetRemoveAllProductTypes

`func (o *TaxCode) GetRemoveAllProductTypes() bool`

GetRemoveAllProductTypes returns the RemoveAllProductTypes field if non-nil, zero value otherwise.

### GetRemoveAllProductTypesOk

`func (o *TaxCode) GetRemoveAllProductTypesOk() (*bool, bool)`

GetRemoveAllProductTypesOk returns a tuple with the RemoveAllProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllProductTypes

`func (o *TaxCode) SetRemoveAllProductTypes(v bool)`

SetRemoveAllProductTypes sets RemoveAllProductTypes field to given value.

### HasRemoveAllProductTypes

`func (o *TaxCode) HasRemoveAllProductTypes() bool`

HasRemoveAllProductTypes returns a boolean if a field has been set.

### SetRemoveAllProductTypesNil

`func (o *TaxCode) SetRemoveAllProductTypesNil(b bool)`

 SetRemoveAllProductTypesNil sets the value for RemoveAllProductTypes to be an explicit nil

### UnsetRemoveAllProductTypes
`func (o *TaxCode) UnsetRemoveAllProductTypes()`

UnsetRemoveAllProductTypes ensures that no value is present for RemoveAllProductTypes, not even an explicit nil
### GetInfo

`func (o *TaxCode) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TaxCode) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TaxCode) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TaxCode) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


