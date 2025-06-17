# UnpostedExpense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**DepartmentId** | Pointer to **NullableInt32** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**CreditAccount** | Pointer to **string** |  | [optional] 
**ExpenseDetailId** | Pointer to **NullableInt32** |  | [optional] 
**ExpenseType** | Pointer to [**ExpenseTypeReference**](ExpenseTypeReference.md) |  | [optional] 
**Classification** | Pointer to **NullableString** |  | [optional] 
**GlType** | Pointer to **NullableString** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**DateExpense** | Pointer to **string** |  | [optional] 
**ChargeCode** | Pointer to [**ChargeCodeReference**](ChargeCodeReference.md) |  | [optional] 
**ChargeDescription** | Pointer to **string** |  | [optional] 
**InPolicy** | Pointer to **NullableBool** |  | [optional] 
**PaymentMethod** | Pointer to [**PaymentMethodReference**](PaymentMethodReference.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**BillableAmount** | Pointer to **NullableFloat64** |  | [optional] 
**NonBillableAmount** | Pointer to **NullableFloat64** |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**AgreementAmountCovered** | Pointer to **NullableFloat64** |  | [optional] 
**Ticket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**ProjectPhase** | Pointer to [**ProjectPhaseReference**](ProjectPhaseReference.md) |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**AvalaraTaxFlag** | Pointer to **NullableBool** | Used to determine if Avalara tax is enabled. | [optional] 
**ItemTaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**SalesTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
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
**DateClosed** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUnpostedExpense

`func NewUnpostedExpense() *UnpostedExpense`

NewUnpostedExpense instantiates a new UnpostedExpense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnpostedExpenseWithDefaults

`func NewUnpostedExpenseWithDefaults() *UnpostedExpense`

NewUnpostedExpenseWithDefaults instantiates a new UnpostedExpense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnpostedExpense) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnpostedExpense) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnpostedExpense) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UnpostedExpense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocationId

`func (o *UnpostedExpense) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *UnpostedExpense) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *UnpostedExpense) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *UnpostedExpense) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *UnpostedExpense) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *UnpostedExpense) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetDepartmentId

`func (o *UnpostedExpense) GetDepartmentId() int32`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *UnpostedExpense) GetDepartmentIdOk() (*int32, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *UnpostedExpense) SetDepartmentId(v int32)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *UnpostedExpense) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### SetDepartmentIdNil

`func (o *UnpostedExpense) SetDepartmentIdNil(b bool)`

 SetDepartmentIdNil sets the value for DepartmentId to be an explicit nil

### UnsetDepartmentId
`func (o *UnpostedExpense) UnsetDepartmentId()`

UnsetDepartmentId ensures that no value is present for DepartmentId, not even an explicit nil
### GetCompany

`func (o *UnpostedExpense) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UnpostedExpense) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UnpostedExpense) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UnpostedExpense) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAccountNumber

`func (o *UnpostedExpense) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *UnpostedExpense) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *UnpostedExpense) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *UnpostedExpense) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetCreditAccount

`func (o *UnpostedExpense) GetCreditAccount() string`

GetCreditAccount returns the CreditAccount field if non-nil, zero value otherwise.

### GetCreditAccountOk

`func (o *UnpostedExpense) GetCreditAccountOk() (*string, bool)`

GetCreditAccountOk returns a tuple with the CreditAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccount

`func (o *UnpostedExpense) SetCreditAccount(v string)`

SetCreditAccount sets CreditAccount field to given value.

### HasCreditAccount

`func (o *UnpostedExpense) HasCreditAccount() bool`

HasCreditAccount returns a boolean if a field has been set.

### GetExpenseDetailId

`func (o *UnpostedExpense) GetExpenseDetailId() int32`

GetExpenseDetailId returns the ExpenseDetailId field if non-nil, zero value otherwise.

### GetExpenseDetailIdOk

`func (o *UnpostedExpense) GetExpenseDetailIdOk() (*int32, bool)`

GetExpenseDetailIdOk returns a tuple with the ExpenseDetailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseDetailId

`func (o *UnpostedExpense) SetExpenseDetailId(v int32)`

SetExpenseDetailId sets ExpenseDetailId field to given value.

### HasExpenseDetailId

`func (o *UnpostedExpense) HasExpenseDetailId() bool`

HasExpenseDetailId returns a boolean if a field has been set.

### SetExpenseDetailIdNil

`func (o *UnpostedExpense) SetExpenseDetailIdNil(b bool)`

 SetExpenseDetailIdNil sets the value for ExpenseDetailId to be an explicit nil

### UnsetExpenseDetailId
`func (o *UnpostedExpense) UnsetExpenseDetailId()`

UnsetExpenseDetailId ensures that no value is present for ExpenseDetailId, not even an explicit nil
### GetExpenseType

`func (o *UnpostedExpense) GetExpenseType() ExpenseTypeReference`

GetExpenseType returns the ExpenseType field if non-nil, zero value otherwise.

### GetExpenseTypeOk

`func (o *UnpostedExpense) GetExpenseTypeOk() (*ExpenseTypeReference, bool)`

GetExpenseTypeOk returns a tuple with the ExpenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseType

`func (o *UnpostedExpense) SetExpenseType(v ExpenseTypeReference)`

SetExpenseType sets ExpenseType field to given value.

### HasExpenseType

`func (o *UnpostedExpense) HasExpenseType() bool`

HasExpenseType returns a boolean if a field has been set.

### GetClassification

`func (o *UnpostedExpense) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *UnpostedExpense) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *UnpostedExpense) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *UnpostedExpense) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *UnpostedExpense) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *UnpostedExpense) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetGlType

`func (o *UnpostedExpense) GetGlType() string`

GetGlType returns the GlType field if non-nil, zero value otherwise.

### GetGlTypeOk

`func (o *UnpostedExpense) GetGlTypeOk() (*string, bool)`

GetGlTypeOk returns a tuple with the GlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlType

`func (o *UnpostedExpense) SetGlType(v string)`

SetGlType sets GlType field to given value.

### HasGlType

`func (o *UnpostedExpense) HasGlType() bool`

HasGlType returns a boolean if a field has been set.

### SetGlTypeNil

`func (o *UnpostedExpense) SetGlTypeNil(b bool)`

 SetGlTypeNil sets the value for GlType to be an explicit nil

### UnsetGlType
`func (o *UnpostedExpense) UnsetGlType()`

UnsetGlType ensures that no value is present for GlType, not even an explicit nil
### GetMember

`func (o *UnpostedExpense) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *UnpostedExpense) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *UnpostedExpense) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *UnpostedExpense) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetDateExpense

`func (o *UnpostedExpense) GetDateExpense() string`

GetDateExpense returns the DateExpense field if non-nil, zero value otherwise.

### GetDateExpenseOk

`func (o *UnpostedExpense) GetDateExpenseOk() (*string, bool)`

GetDateExpenseOk returns a tuple with the DateExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateExpense

`func (o *UnpostedExpense) SetDateExpense(v string)`

SetDateExpense sets DateExpense field to given value.

### HasDateExpense

`func (o *UnpostedExpense) HasDateExpense() bool`

HasDateExpense returns a boolean if a field has been set.

### GetChargeCode

`func (o *UnpostedExpense) GetChargeCode() ChargeCodeReference`

GetChargeCode returns the ChargeCode field if non-nil, zero value otherwise.

### GetChargeCodeOk

`func (o *UnpostedExpense) GetChargeCodeOk() (*ChargeCodeReference, bool)`

GetChargeCodeOk returns a tuple with the ChargeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeCode

`func (o *UnpostedExpense) SetChargeCode(v ChargeCodeReference)`

SetChargeCode sets ChargeCode field to given value.

### HasChargeCode

`func (o *UnpostedExpense) HasChargeCode() bool`

HasChargeCode returns a boolean if a field has been set.

### GetChargeDescription

`func (o *UnpostedExpense) GetChargeDescription() string`

GetChargeDescription returns the ChargeDescription field if non-nil, zero value otherwise.

### GetChargeDescriptionOk

`func (o *UnpostedExpense) GetChargeDescriptionOk() (*string, bool)`

GetChargeDescriptionOk returns a tuple with the ChargeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeDescription

`func (o *UnpostedExpense) SetChargeDescription(v string)`

SetChargeDescription sets ChargeDescription field to given value.

### HasChargeDescription

`func (o *UnpostedExpense) HasChargeDescription() bool`

HasChargeDescription returns a boolean if a field has been set.

### GetInPolicy

`func (o *UnpostedExpense) GetInPolicy() bool`

GetInPolicy returns the InPolicy field if non-nil, zero value otherwise.

### GetInPolicyOk

`func (o *UnpostedExpense) GetInPolicyOk() (*bool, bool)`

GetInPolicyOk returns a tuple with the InPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInPolicy

`func (o *UnpostedExpense) SetInPolicy(v bool)`

SetInPolicy sets InPolicy field to given value.

### HasInPolicy

`func (o *UnpostedExpense) HasInPolicy() bool`

HasInPolicy returns a boolean if a field has been set.

### SetInPolicyNil

`func (o *UnpostedExpense) SetInPolicyNil(b bool)`

 SetInPolicyNil sets the value for InPolicy to be an explicit nil

### UnsetInPolicy
`func (o *UnpostedExpense) UnsetInPolicy()`

UnsetInPolicy ensures that no value is present for InPolicy, not even an explicit nil
### GetPaymentMethod

`func (o *UnpostedExpense) GetPaymentMethod() PaymentMethodReference`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *UnpostedExpense) GetPaymentMethodOk() (*PaymentMethodReference, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *UnpostedExpense) SetPaymentMethod(v PaymentMethodReference)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *UnpostedExpense) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetCurrency

`func (o *UnpostedExpense) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnpostedExpense) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnpostedExpense) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnpostedExpense) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotal

`func (o *UnpostedExpense) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnpostedExpense) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnpostedExpense) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnpostedExpense) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *UnpostedExpense) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *UnpostedExpense) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetBillableAmount

`func (o *UnpostedExpense) GetBillableAmount() float64`

GetBillableAmount returns the BillableAmount field if non-nil, zero value otherwise.

### GetBillableAmountOk

`func (o *UnpostedExpense) GetBillableAmountOk() (*float64, bool)`

GetBillableAmountOk returns a tuple with the BillableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableAmount

`func (o *UnpostedExpense) SetBillableAmount(v float64)`

SetBillableAmount sets BillableAmount field to given value.

### HasBillableAmount

`func (o *UnpostedExpense) HasBillableAmount() bool`

HasBillableAmount returns a boolean if a field has been set.

### SetBillableAmountNil

`func (o *UnpostedExpense) SetBillableAmountNil(b bool)`

 SetBillableAmountNil sets the value for BillableAmount to be an explicit nil

### UnsetBillableAmount
`func (o *UnpostedExpense) UnsetBillableAmount()`

UnsetBillableAmount ensures that no value is present for BillableAmount, not even an explicit nil
### GetNonBillableAmount

`func (o *UnpostedExpense) GetNonBillableAmount() float64`

GetNonBillableAmount returns the NonBillableAmount field if non-nil, zero value otherwise.

### GetNonBillableAmountOk

`func (o *UnpostedExpense) GetNonBillableAmountOk() (*float64, bool)`

GetNonBillableAmountOk returns a tuple with the NonBillableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonBillableAmount

`func (o *UnpostedExpense) SetNonBillableAmount(v float64)`

SetNonBillableAmount sets NonBillableAmount field to given value.

### HasNonBillableAmount

`func (o *UnpostedExpense) HasNonBillableAmount() bool`

HasNonBillableAmount returns a boolean if a field has been set.

### SetNonBillableAmountNil

`func (o *UnpostedExpense) SetNonBillableAmountNil(b bool)`

 SetNonBillableAmountNil sets the value for NonBillableAmount to be an explicit nil

### UnsetNonBillableAmount
`func (o *UnpostedExpense) UnsetNonBillableAmount()`

UnsetNonBillableAmount ensures that no value is present for NonBillableAmount, not even an explicit nil
### GetAgreement

`func (o *UnpostedExpense) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *UnpostedExpense) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *UnpostedExpense) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *UnpostedExpense) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetAgreementAmountCovered

`func (o *UnpostedExpense) GetAgreementAmountCovered() float64`

GetAgreementAmountCovered returns the AgreementAmountCovered field if non-nil, zero value otherwise.

### GetAgreementAmountCoveredOk

`func (o *UnpostedExpense) GetAgreementAmountCoveredOk() (*float64, bool)`

GetAgreementAmountCoveredOk returns a tuple with the AgreementAmountCovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementAmountCovered

`func (o *UnpostedExpense) SetAgreementAmountCovered(v float64)`

SetAgreementAmountCovered sets AgreementAmountCovered field to given value.

### HasAgreementAmountCovered

`func (o *UnpostedExpense) HasAgreementAmountCovered() bool`

HasAgreementAmountCovered returns a boolean if a field has been set.

### SetAgreementAmountCoveredNil

`func (o *UnpostedExpense) SetAgreementAmountCoveredNil(b bool)`

 SetAgreementAmountCoveredNil sets the value for AgreementAmountCovered to be an explicit nil

### UnsetAgreementAmountCovered
`func (o *UnpostedExpense) UnsetAgreementAmountCovered()`

UnsetAgreementAmountCovered ensures that no value is present for AgreementAmountCovered, not even an explicit nil
### GetTicket

`func (o *UnpostedExpense) GetTicket() TicketReference`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *UnpostedExpense) GetTicketOk() (*TicketReference, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *UnpostedExpense) SetTicket(v TicketReference)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *UnpostedExpense) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetProject

`func (o *UnpostedExpense) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *UnpostedExpense) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *UnpostedExpense) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *UnpostedExpense) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectPhase

`func (o *UnpostedExpense) GetProjectPhase() ProjectPhaseReference`

GetProjectPhase returns the ProjectPhase field if non-nil, zero value otherwise.

### GetProjectPhaseOk

`func (o *UnpostedExpense) GetProjectPhaseOk() (*ProjectPhaseReference, bool)`

GetProjectPhaseOk returns a tuple with the ProjectPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPhase

`func (o *UnpostedExpense) SetProjectPhase(v ProjectPhaseReference)`

SetProjectPhase sets ProjectPhase field to given value.

### HasProjectPhase

`func (o *UnpostedExpense) HasProjectPhase() bool`

HasProjectPhase returns a boolean if a field has been set.

### GetTaxCode

`func (o *UnpostedExpense) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *UnpostedExpense) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *UnpostedExpense) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *UnpostedExpense) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetAvalaraTaxFlag

`func (o *UnpostedExpense) GetAvalaraTaxFlag() bool`

GetAvalaraTaxFlag returns the AvalaraTaxFlag field if non-nil, zero value otherwise.

### GetAvalaraTaxFlagOk

`func (o *UnpostedExpense) GetAvalaraTaxFlagOk() (*bool, bool)`

GetAvalaraTaxFlagOk returns a tuple with the AvalaraTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvalaraTaxFlag

`func (o *UnpostedExpense) SetAvalaraTaxFlag(v bool)`

SetAvalaraTaxFlag sets AvalaraTaxFlag field to given value.

### HasAvalaraTaxFlag

`func (o *UnpostedExpense) HasAvalaraTaxFlag() bool`

HasAvalaraTaxFlag returns a boolean if a field has been set.

### SetAvalaraTaxFlagNil

`func (o *UnpostedExpense) SetAvalaraTaxFlagNil(b bool)`

 SetAvalaraTaxFlagNil sets the value for AvalaraTaxFlag to be an explicit nil

### UnsetAvalaraTaxFlag
`func (o *UnpostedExpense) UnsetAvalaraTaxFlag()`

UnsetAvalaraTaxFlag ensures that no value is present for AvalaraTaxFlag, not even an explicit nil
### GetItemTaxableFlag

`func (o *UnpostedExpense) GetItemTaxableFlag() bool`

GetItemTaxableFlag returns the ItemTaxableFlag field if non-nil, zero value otherwise.

### GetItemTaxableFlagOk

`func (o *UnpostedExpense) GetItemTaxableFlagOk() (*bool, bool)`

GetItemTaxableFlagOk returns a tuple with the ItemTaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTaxableFlag

`func (o *UnpostedExpense) SetItemTaxableFlag(v bool)`

SetItemTaxableFlag sets ItemTaxableFlag field to given value.

### HasItemTaxableFlag

`func (o *UnpostedExpense) HasItemTaxableFlag() bool`

HasItemTaxableFlag returns a boolean if a field has been set.

### SetItemTaxableFlagNil

`func (o *UnpostedExpense) SetItemTaxableFlagNil(b bool)`

 SetItemTaxableFlagNil sets the value for ItemTaxableFlag to be an explicit nil

### UnsetItemTaxableFlag
`func (o *UnpostedExpense) UnsetItemTaxableFlag()`

UnsetItemTaxableFlag ensures that no value is present for ItemTaxableFlag, not even an explicit nil
### GetSalesTaxAmount

`func (o *UnpostedExpense) GetSalesTaxAmount() float64`

GetSalesTaxAmount returns the SalesTaxAmount field if non-nil, zero value otherwise.

### GetSalesTaxAmountOk

`func (o *UnpostedExpense) GetSalesTaxAmountOk() (*float64, bool)`

GetSalesTaxAmountOk returns a tuple with the SalesTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxAmount

`func (o *UnpostedExpense) SetSalesTaxAmount(v float64)`

SetSalesTaxAmount sets SalesTaxAmount field to given value.

### HasSalesTaxAmount

`func (o *UnpostedExpense) HasSalesTaxAmount() bool`

HasSalesTaxAmount returns a boolean if a field has been set.

### SetSalesTaxAmountNil

`func (o *UnpostedExpense) SetSalesTaxAmountNil(b bool)`

 SetSalesTaxAmountNil sets the value for SalesTaxAmount to be an explicit nil

### UnsetSalesTaxAmount
`func (o *UnpostedExpense) UnsetSalesTaxAmount()`

UnsetSalesTaxAmount ensures that no value is present for SalesTaxAmount, not even an explicit nil
### GetStateTaxFlag

`func (o *UnpostedExpense) GetStateTaxFlag() bool`

GetStateTaxFlag returns the StateTaxFlag field if non-nil, zero value otherwise.

### GetStateTaxFlagOk

`func (o *UnpostedExpense) GetStateTaxFlagOk() (*bool, bool)`

GetStateTaxFlagOk returns a tuple with the StateTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxFlag

`func (o *UnpostedExpense) SetStateTaxFlag(v bool)`

SetStateTaxFlag sets StateTaxFlag field to given value.

### HasStateTaxFlag

`func (o *UnpostedExpense) HasStateTaxFlag() bool`

HasStateTaxFlag returns a boolean if a field has been set.

### SetStateTaxFlagNil

`func (o *UnpostedExpense) SetStateTaxFlagNil(b bool)`

 SetStateTaxFlagNil sets the value for StateTaxFlag to be an explicit nil

### UnsetStateTaxFlag
`func (o *UnpostedExpense) UnsetStateTaxFlag()`

UnsetStateTaxFlag ensures that no value is present for StateTaxFlag, not even an explicit nil
### GetStateTaxXref

`func (o *UnpostedExpense) GetStateTaxXref() string`

GetStateTaxXref returns the StateTaxXref field if non-nil, zero value otherwise.

### GetStateTaxXrefOk

`func (o *UnpostedExpense) GetStateTaxXrefOk() (*string, bool)`

GetStateTaxXrefOk returns a tuple with the StateTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxXref

`func (o *UnpostedExpense) SetStateTaxXref(v string)`

SetStateTaxXref sets StateTaxXref field to given value.

### HasStateTaxXref

`func (o *UnpostedExpense) HasStateTaxXref() bool`

HasStateTaxXref returns a boolean if a field has been set.

### GetStateTaxAmount

`func (o *UnpostedExpense) GetStateTaxAmount() float64`

GetStateTaxAmount returns the StateTaxAmount field if non-nil, zero value otherwise.

### GetStateTaxAmountOk

`func (o *UnpostedExpense) GetStateTaxAmountOk() (*float64, bool)`

GetStateTaxAmountOk returns a tuple with the StateTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxAmount

`func (o *UnpostedExpense) SetStateTaxAmount(v float64)`

SetStateTaxAmount sets StateTaxAmount field to given value.

### HasStateTaxAmount

`func (o *UnpostedExpense) HasStateTaxAmount() bool`

HasStateTaxAmount returns a boolean if a field has been set.

### SetStateTaxAmountNil

`func (o *UnpostedExpense) SetStateTaxAmountNil(b bool)`

 SetStateTaxAmountNil sets the value for StateTaxAmount to be an explicit nil

### UnsetStateTaxAmount
`func (o *UnpostedExpense) UnsetStateTaxAmount()`

UnsetStateTaxAmount ensures that no value is present for StateTaxAmount, not even an explicit nil
### GetCountyTaxFlag

`func (o *UnpostedExpense) GetCountyTaxFlag() bool`

GetCountyTaxFlag returns the CountyTaxFlag field if non-nil, zero value otherwise.

### GetCountyTaxFlagOk

`func (o *UnpostedExpense) GetCountyTaxFlagOk() (*bool, bool)`

GetCountyTaxFlagOk returns a tuple with the CountyTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxFlag

`func (o *UnpostedExpense) SetCountyTaxFlag(v bool)`

SetCountyTaxFlag sets CountyTaxFlag field to given value.

### HasCountyTaxFlag

`func (o *UnpostedExpense) HasCountyTaxFlag() bool`

HasCountyTaxFlag returns a boolean if a field has been set.

### SetCountyTaxFlagNil

`func (o *UnpostedExpense) SetCountyTaxFlagNil(b bool)`

 SetCountyTaxFlagNil sets the value for CountyTaxFlag to be an explicit nil

### UnsetCountyTaxFlag
`func (o *UnpostedExpense) UnsetCountyTaxFlag()`

UnsetCountyTaxFlag ensures that no value is present for CountyTaxFlag, not even an explicit nil
### GetCountyTaxXref

`func (o *UnpostedExpense) GetCountyTaxXref() string`

GetCountyTaxXref returns the CountyTaxXref field if non-nil, zero value otherwise.

### GetCountyTaxXrefOk

`func (o *UnpostedExpense) GetCountyTaxXrefOk() (*string, bool)`

GetCountyTaxXrefOk returns a tuple with the CountyTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxXref

`func (o *UnpostedExpense) SetCountyTaxXref(v string)`

SetCountyTaxXref sets CountyTaxXref field to given value.

### HasCountyTaxXref

`func (o *UnpostedExpense) HasCountyTaxXref() bool`

HasCountyTaxXref returns a boolean if a field has been set.

### GetCountyTaxAmount

`func (o *UnpostedExpense) GetCountyTaxAmount() float64`

GetCountyTaxAmount returns the CountyTaxAmount field if non-nil, zero value otherwise.

### GetCountyTaxAmountOk

`func (o *UnpostedExpense) GetCountyTaxAmountOk() (*float64, bool)`

GetCountyTaxAmountOk returns a tuple with the CountyTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxAmount

`func (o *UnpostedExpense) SetCountyTaxAmount(v float64)`

SetCountyTaxAmount sets CountyTaxAmount field to given value.

### HasCountyTaxAmount

`func (o *UnpostedExpense) HasCountyTaxAmount() bool`

HasCountyTaxAmount returns a boolean if a field has been set.

### SetCountyTaxAmountNil

`func (o *UnpostedExpense) SetCountyTaxAmountNil(b bool)`

 SetCountyTaxAmountNil sets the value for CountyTaxAmount to be an explicit nil

### UnsetCountyTaxAmount
`func (o *UnpostedExpense) UnsetCountyTaxAmount()`

UnsetCountyTaxAmount ensures that no value is present for CountyTaxAmount, not even an explicit nil
### GetCityTaxFlag

`func (o *UnpostedExpense) GetCityTaxFlag() bool`

GetCityTaxFlag returns the CityTaxFlag field if non-nil, zero value otherwise.

### GetCityTaxFlagOk

`func (o *UnpostedExpense) GetCityTaxFlagOk() (*bool, bool)`

GetCityTaxFlagOk returns a tuple with the CityTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxFlag

`func (o *UnpostedExpense) SetCityTaxFlag(v bool)`

SetCityTaxFlag sets CityTaxFlag field to given value.

### HasCityTaxFlag

`func (o *UnpostedExpense) HasCityTaxFlag() bool`

HasCityTaxFlag returns a boolean if a field has been set.

### SetCityTaxFlagNil

`func (o *UnpostedExpense) SetCityTaxFlagNil(b bool)`

 SetCityTaxFlagNil sets the value for CityTaxFlag to be an explicit nil

### UnsetCityTaxFlag
`func (o *UnpostedExpense) UnsetCityTaxFlag()`

UnsetCityTaxFlag ensures that no value is present for CityTaxFlag, not even an explicit nil
### GetCityTaxXref

`func (o *UnpostedExpense) GetCityTaxXref() string`

GetCityTaxXref returns the CityTaxXref field if non-nil, zero value otherwise.

### GetCityTaxXrefOk

`func (o *UnpostedExpense) GetCityTaxXrefOk() (*string, bool)`

GetCityTaxXrefOk returns a tuple with the CityTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxXref

`func (o *UnpostedExpense) SetCityTaxXref(v string)`

SetCityTaxXref sets CityTaxXref field to given value.

### HasCityTaxXref

`func (o *UnpostedExpense) HasCityTaxXref() bool`

HasCityTaxXref returns a boolean if a field has been set.

### GetCityTaxAmount

`func (o *UnpostedExpense) GetCityTaxAmount() float64`

GetCityTaxAmount returns the CityTaxAmount field if non-nil, zero value otherwise.

### GetCityTaxAmountOk

`func (o *UnpostedExpense) GetCityTaxAmountOk() (*float64, bool)`

GetCityTaxAmountOk returns a tuple with the CityTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxAmount

`func (o *UnpostedExpense) SetCityTaxAmount(v float64)`

SetCityTaxAmount sets CityTaxAmount field to given value.

### HasCityTaxAmount

`func (o *UnpostedExpense) HasCityTaxAmount() bool`

HasCityTaxAmount returns a boolean if a field has been set.

### SetCityTaxAmountNil

`func (o *UnpostedExpense) SetCityTaxAmountNil(b bool)`

 SetCityTaxAmountNil sets the value for CityTaxAmount to be an explicit nil

### UnsetCityTaxAmount
`func (o *UnpostedExpense) UnsetCityTaxAmount()`

UnsetCityTaxAmount ensures that no value is present for CityTaxAmount, not even an explicit nil
### GetCountryTaxFlag

`func (o *UnpostedExpense) GetCountryTaxFlag() bool`

GetCountryTaxFlag returns the CountryTaxFlag field if non-nil, zero value otherwise.

### GetCountryTaxFlagOk

`func (o *UnpostedExpense) GetCountryTaxFlagOk() (*bool, bool)`

GetCountryTaxFlagOk returns a tuple with the CountryTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryTaxFlag

`func (o *UnpostedExpense) SetCountryTaxFlag(v bool)`

SetCountryTaxFlag sets CountryTaxFlag field to given value.

### HasCountryTaxFlag

`func (o *UnpostedExpense) HasCountryTaxFlag() bool`

HasCountryTaxFlag returns a boolean if a field has been set.

### SetCountryTaxFlagNil

`func (o *UnpostedExpense) SetCountryTaxFlagNil(b bool)`

 SetCountryTaxFlagNil sets the value for CountryTaxFlag to be an explicit nil

### UnsetCountryTaxFlag
`func (o *UnpostedExpense) UnsetCountryTaxFlag()`

UnsetCountryTaxFlag ensures that no value is present for CountryTaxFlag, not even an explicit nil
### GetCountryTaxXref

`func (o *UnpostedExpense) GetCountryTaxXref() string`

GetCountryTaxXref returns the CountryTaxXref field if non-nil, zero value otherwise.

### GetCountryTaxXrefOk

`func (o *UnpostedExpense) GetCountryTaxXrefOk() (*string, bool)`

GetCountryTaxXrefOk returns a tuple with the CountryTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryTaxXref

`func (o *UnpostedExpense) SetCountryTaxXref(v string)`

SetCountryTaxXref sets CountryTaxXref field to given value.

### HasCountryTaxXref

`func (o *UnpostedExpense) HasCountryTaxXref() bool`

HasCountryTaxXref returns a boolean if a field has been set.

### GetCountryTaxAmount

`func (o *UnpostedExpense) GetCountryTaxAmount() float64`

GetCountryTaxAmount returns the CountryTaxAmount field if non-nil, zero value otherwise.

### GetCountryTaxAmountOk

`func (o *UnpostedExpense) GetCountryTaxAmountOk() (*float64, bool)`

GetCountryTaxAmountOk returns a tuple with the CountryTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryTaxAmount

`func (o *UnpostedExpense) SetCountryTaxAmount(v float64)`

SetCountryTaxAmount sets CountryTaxAmount field to given value.

### HasCountryTaxAmount

`func (o *UnpostedExpense) HasCountryTaxAmount() bool`

HasCountryTaxAmount returns a boolean if a field has been set.

### SetCountryTaxAmountNil

`func (o *UnpostedExpense) SetCountryTaxAmountNil(b bool)`

 SetCountryTaxAmountNil sets the value for CountryTaxAmount to be an explicit nil

### UnsetCountryTaxAmount
`func (o *UnpostedExpense) UnsetCountryTaxAmount()`

UnsetCountryTaxAmount ensures that no value is present for CountryTaxAmount, not even an explicit nil
### GetCompositeTaxFlag

`func (o *UnpostedExpense) GetCompositeTaxFlag() bool`

GetCompositeTaxFlag returns the CompositeTaxFlag field if non-nil, zero value otherwise.

### GetCompositeTaxFlagOk

`func (o *UnpostedExpense) GetCompositeTaxFlagOk() (*bool, bool)`

GetCompositeTaxFlagOk returns a tuple with the CompositeTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeTaxFlag

`func (o *UnpostedExpense) SetCompositeTaxFlag(v bool)`

SetCompositeTaxFlag sets CompositeTaxFlag field to given value.

### HasCompositeTaxFlag

`func (o *UnpostedExpense) HasCompositeTaxFlag() bool`

HasCompositeTaxFlag returns a boolean if a field has been set.

### SetCompositeTaxFlagNil

`func (o *UnpostedExpense) SetCompositeTaxFlagNil(b bool)`

 SetCompositeTaxFlagNil sets the value for CompositeTaxFlag to be an explicit nil

### UnsetCompositeTaxFlag
`func (o *UnpostedExpense) UnsetCompositeTaxFlag()`

UnsetCompositeTaxFlag ensures that no value is present for CompositeTaxFlag, not even an explicit nil
### GetCompositeTaxXref

`func (o *UnpostedExpense) GetCompositeTaxXref() string`

GetCompositeTaxXref returns the CompositeTaxXref field if non-nil, zero value otherwise.

### GetCompositeTaxXrefOk

`func (o *UnpostedExpense) GetCompositeTaxXrefOk() (*string, bool)`

GetCompositeTaxXrefOk returns a tuple with the CompositeTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeTaxXref

`func (o *UnpostedExpense) SetCompositeTaxXref(v string)`

SetCompositeTaxXref sets CompositeTaxXref field to given value.

### HasCompositeTaxXref

`func (o *UnpostedExpense) HasCompositeTaxXref() bool`

HasCompositeTaxXref returns a boolean if a field has been set.

### GetCompositeTaxAmount

`func (o *UnpostedExpense) GetCompositeTaxAmount() float64`

GetCompositeTaxAmount returns the CompositeTaxAmount field if non-nil, zero value otherwise.

### GetCompositeTaxAmountOk

`func (o *UnpostedExpense) GetCompositeTaxAmountOk() (*float64, bool)`

GetCompositeTaxAmountOk returns a tuple with the CompositeTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeTaxAmount

`func (o *UnpostedExpense) SetCompositeTaxAmount(v float64)`

SetCompositeTaxAmount sets CompositeTaxAmount field to given value.

### HasCompositeTaxAmount

`func (o *UnpostedExpense) HasCompositeTaxAmount() bool`

HasCompositeTaxAmount returns a boolean if a field has been set.

### SetCompositeTaxAmountNil

`func (o *UnpostedExpense) SetCompositeTaxAmountNil(b bool)`

 SetCompositeTaxAmountNil sets the value for CompositeTaxAmount to be an explicit nil

### UnsetCompositeTaxAmount
`func (o *UnpostedExpense) UnsetCompositeTaxAmount()`

UnsetCompositeTaxAmount ensures that no value is present for CompositeTaxAmount, not even an explicit nil
### GetLevelSixTaxFlag

`func (o *UnpostedExpense) GetLevelSixTaxFlag() bool`

GetLevelSixTaxFlag returns the LevelSixTaxFlag field if non-nil, zero value otherwise.

### GetLevelSixTaxFlagOk

`func (o *UnpostedExpense) GetLevelSixTaxFlagOk() (*bool, bool)`

GetLevelSixTaxFlagOk returns a tuple with the LevelSixTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixTaxFlag

`func (o *UnpostedExpense) SetLevelSixTaxFlag(v bool)`

SetLevelSixTaxFlag sets LevelSixTaxFlag field to given value.

### HasLevelSixTaxFlag

`func (o *UnpostedExpense) HasLevelSixTaxFlag() bool`

HasLevelSixTaxFlag returns a boolean if a field has been set.

### SetLevelSixTaxFlagNil

`func (o *UnpostedExpense) SetLevelSixTaxFlagNil(b bool)`

 SetLevelSixTaxFlagNil sets the value for LevelSixTaxFlag to be an explicit nil

### UnsetLevelSixTaxFlag
`func (o *UnpostedExpense) UnsetLevelSixTaxFlag()`

UnsetLevelSixTaxFlag ensures that no value is present for LevelSixTaxFlag, not even an explicit nil
### GetLevelSixTaxXref

`func (o *UnpostedExpense) GetLevelSixTaxXref() string`

GetLevelSixTaxXref returns the LevelSixTaxXref field if non-nil, zero value otherwise.

### GetLevelSixTaxXrefOk

`func (o *UnpostedExpense) GetLevelSixTaxXrefOk() (*string, bool)`

GetLevelSixTaxXrefOk returns a tuple with the LevelSixTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixTaxXref

`func (o *UnpostedExpense) SetLevelSixTaxXref(v string)`

SetLevelSixTaxXref sets LevelSixTaxXref field to given value.

### HasLevelSixTaxXref

`func (o *UnpostedExpense) HasLevelSixTaxXref() bool`

HasLevelSixTaxXref returns a boolean if a field has been set.

### GetLevelSixTaxAmount

`func (o *UnpostedExpense) GetLevelSixTaxAmount() float64`

GetLevelSixTaxAmount returns the LevelSixTaxAmount field if non-nil, zero value otherwise.

### GetLevelSixTaxAmountOk

`func (o *UnpostedExpense) GetLevelSixTaxAmountOk() (*float64, bool)`

GetLevelSixTaxAmountOk returns a tuple with the LevelSixTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixTaxAmount

`func (o *UnpostedExpense) SetLevelSixTaxAmount(v float64)`

SetLevelSixTaxAmount sets LevelSixTaxAmount field to given value.

### HasLevelSixTaxAmount

`func (o *UnpostedExpense) HasLevelSixTaxAmount() bool`

HasLevelSixTaxAmount returns a boolean if a field has been set.

### SetLevelSixTaxAmountNil

`func (o *UnpostedExpense) SetLevelSixTaxAmountNil(b bool)`

 SetLevelSixTaxAmountNil sets the value for LevelSixTaxAmount to be an explicit nil

### UnsetLevelSixTaxAmount
`func (o *UnpostedExpense) UnsetLevelSixTaxAmount()`

UnsetLevelSixTaxAmount ensures that no value is present for LevelSixTaxAmount, not even an explicit nil
### GetDateClosed

`func (o *UnpostedExpense) GetDateClosed() string`

GetDateClosed returns the DateClosed field if non-nil, zero value otherwise.

### GetDateClosedOk

`func (o *UnpostedExpense) GetDateClosedOk() (*string, bool)`

GetDateClosedOk returns a tuple with the DateClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClosed

`func (o *UnpostedExpense) SetDateClosed(v string)`

SetDateClosed sets DateClosed field to given value.

### HasDateClosed

`func (o *UnpostedExpense) HasDateClosed() bool`

HasDateClosed returns a boolean if a field has been set.

### GetInfo

`func (o *UnpostedExpense) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UnpostedExpense) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UnpostedExpense) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *UnpostedExpense) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


