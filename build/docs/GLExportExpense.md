# GLExportExpense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DocumentDate** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 
**ApAccountNumber** | Pointer to **string** |  | [optional] 
**ApClass** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**GlClass** | Pointer to **string** |  | [optional] 
**GlTypeId** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PeriodStartDate** | Pointer to **string** |  | [optional] 
**PeriodEndDate** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**VendorNumber** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**CompanyAccountNumber** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**Offset** | Pointer to [**GLExportExpenseOffset**](GLExportExpenseOffset.md) |  | [optional] 

## Methods

### NewGLExportExpense

`func NewGLExportExpense() *GLExportExpense`

NewGLExportExpense instantiates a new GLExportExpense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportExpenseWithDefaults

`func NewGLExportExpenseWithDefaults() *GLExportExpense`

NewGLExportExpenseWithDefaults instantiates a new GLExportExpense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLExportExpense) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLExportExpense) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLExportExpense) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GLExportExpense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocumentDate

`func (o *GLExportExpense) GetDocumentDate() string`

GetDocumentDate returns the DocumentDate field if non-nil, zero value otherwise.

### GetDocumentDateOk

`func (o *GLExportExpense) GetDocumentDateOk() (*string, bool)`

GetDocumentDateOk returns a tuple with the DocumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDate

`func (o *GLExportExpense) SetDocumentDate(v string)`

SetDocumentDate sets DocumentDate field to given value.

### HasDocumentDate

`func (o *GLExportExpense) HasDocumentDate() bool`

HasDocumentDate returns a boolean if a field has been set.

### GetDocumentType

`func (o *GLExportExpense) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *GLExportExpense) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *GLExportExpense) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *GLExportExpense) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetApAccountNumber

`func (o *GLExportExpense) GetApAccountNumber() string`

GetApAccountNumber returns the ApAccountNumber field if non-nil, zero value otherwise.

### GetApAccountNumberOk

`func (o *GLExportExpense) GetApAccountNumberOk() (*string, bool)`

GetApAccountNumberOk returns a tuple with the ApAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApAccountNumber

`func (o *GLExportExpense) SetApAccountNumber(v string)`

SetApAccountNumber sets ApAccountNumber field to given value.

### HasApAccountNumber

`func (o *GLExportExpense) HasApAccountNumber() bool`

HasApAccountNumber returns a boolean if a field has been set.

### GetApClass

`func (o *GLExportExpense) GetApClass() string`

GetApClass returns the ApClass field if non-nil, zero value otherwise.

### GetApClassOk

`func (o *GLExportExpense) GetApClassOk() (*string, bool)`

GetApClassOk returns a tuple with the ApClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApClass

`func (o *GLExportExpense) SetApClass(v string)`

SetApClass sets ApClass field to given value.

### HasApClass

`func (o *GLExportExpense) HasApClass() bool`

HasApClass returns a boolean if a field has been set.

### GetAccountNumber

`func (o *GLExportExpense) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GLExportExpense) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GLExportExpense) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GLExportExpense) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetGlClass

`func (o *GLExportExpense) GetGlClass() string`

GetGlClass returns the GlClass field if non-nil, zero value otherwise.

### GetGlClassOk

`func (o *GLExportExpense) GetGlClassOk() (*string, bool)`

GetGlClassOk returns a tuple with the GlClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlClass

`func (o *GLExportExpense) SetGlClass(v string)`

SetGlClass sets GlClass field to given value.

### HasGlClass

`func (o *GLExportExpense) HasGlClass() bool`

HasGlClass returns a boolean if a field has been set.

### GetGlTypeId

`func (o *GLExportExpense) GetGlTypeId() string`

GetGlTypeId returns the GlTypeId field if non-nil, zero value otherwise.

### GetGlTypeIdOk

`func (o *GLExportExpense) GetGlTypeIdOk() (*string, bool)`

GetGlTypeIdOk returns a tuple with the GlTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlTypeId

`func (o *GLExportExpense) SetGlTypeId(v string)`

SetGlTypeId sets GlTypeId field to given value.

### HasGlTypeId

`func (o *GLExportExpense) HasGlTypeId() bool`

HasGlTypeId returns a boolean if a field has been set.

### GetMemo

`func (o *GLExportExpense) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GLExportExpense) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GLExportExpense) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GLExportExpense) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetDescription

`func (o *GLExportExpense) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GLExportExpense) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GLExportExpense) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GLExportExpense) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPeriodStartDate

`func (o *GLExportExpense) GetPeriodStartDate() string`

GetPeriodStartDate returns the PeriodStartDate field if non-nil, zero value otherwise.

### GetPeriodStartDateOk

`func (o *GLExportExpense) GetPeriodStartDateOk() (*string, bool)`

GetPeriodStartDateOk returns a tuple with the PeriodStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStartDate

`func (o *GLExportExpense) SetPeriodStartDate(v string)`

SetPeriodStartDate sets PeriodStartDate field to given value.

### HasPeriodStartDate

`func (o *GLExportExpense) HasPeriodStartDate() bool`

HasPeriodStartDate returns a boolean if a field has been set.

### GetPeriodEndDate

`func (o *GLExportExpense) GetPeriodEndDate() string`

GetPeriodEndDate returns the PeriodEndDate field if non-nil, zero value otherwise.

### GetPeriodEndDateOk

`func (o *GLExportExpense) GetPeriodEndDateOk() (*string, bool)`

GetPeriodEndDateOk returns a tuple with the PeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEndDate

`func (o *GLExportExpense) SetPeriodEndDate(v string)`

SetPeriodEndDate sets PeriodEndDate field to given value.

### HasPeriodEndDate

`func (o *GLExportExpense) HasPeriodEndDate() bool`

HasPeriodEndDate returns a boolean if a field has been set.

### GetMember

`func (o *GLExportExpense) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GLExportExpense) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GLExportExpense) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *GLExportExpense) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetVendorNumber

`func (o *GLExportExpense) GetVendorNumber() string`

GetVendorNumber returns the VendorNumber field if non-nil, zero value otherwise.

### GetVendorNumberOk

`func (o *GLExportExpense) GetVendorNumberOk() (*string, bool)`

GetVendorNumberOk returns a tuple with the VendorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumber

`func (o *GLExportExpense) SetVendorNumber(v string)`

SetVendorNumber sets VendorNumber field to given value.

### HasVendorNumber

`func (o *GLExportExpense) HasVendorNumber() bool`

HasVendorNumber returns a boolean if a field has been set.

### GetCompany

`func (o *GLExportExpense) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GLExportExpense) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GLExportExpense) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GLExportExpense) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyAccountNumber

`func (o *GLExportExpense) GetCompanyAccountNumber() string`

GetCompanyAccountNumber returns the CompanyAccountNumber field if non-nil, zero value otherwise.

### GetCompanyAccountNumberOk

`func (o *GLExportExpense) GetCompanyAccountNumberOk() (*string, bool)`

GetCompanyAccountNumberOk returns a tuple with the CompanyAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAccountNumber

`func (o *GLExportExpense) SetCompanyAccountNumber(v string)`

SetCompanyAccountNumber sets CompanyAccountNumber field to given value.

### HasCompanyAccountNumber

`func (o *GLExportExpense) HasCompanyAccountNumber() bool`

HasCompanyAccountNumber returns a boolean if a field has been set.

### GetProject

`func (o *GLExportExpense) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GLExportExpense) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GLExportExpense) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *GLExportExpense) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetCurrency

`func (o *GLExportExpense) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GLExportExpense) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GLExportExpense) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GLExportExpense) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotal

`func (o *GLExportExpense) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GLExportExpense) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GLExportExpense) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GLExportExpense) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GLExportExpense) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GLExportExpense) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetOffset

`func (o *GLExportExpense) GetOffset() GLExportExpenseOffset`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GLExportExpense) GetOffsetOk() (*GLExportExpenseOffset, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GLExportExpense) SetOffset(v GLExportExpenseOffset)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GLExportExpense) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


