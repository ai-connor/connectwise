# GLExportExpenseBillDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** |  | [optional] 
**DocumentDate** | Pointer to **string** |  | [optional] 
**GlTypeId** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**ExpenseClass** | Pointer to [**ClassificationReference**](ClassificationReference.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**Billable** | Pointer to **NullableBool** |  | [optional] 
**Reimbursable** | Pointer to **NullableBool** |  | [optional] 
**CompanyAdvance** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewGLExportExpenseBillDetail

`func NewGLExportExpenseBillDetail() *GLExportExpenseBillDetail`

NewGLExportExpenseBillDetail instantiates a new GLExportExpenseBillDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportExpenseBillDetailWithDefaults

`func NewGLExportExpenseBillDetailWithDefaults() *GLExportExpenseBillDetail`

NewGLExportExpenseBillDetailWithDefaults instantiates a new GLExportExpenseBillDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLExportExpenseBillDetail) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLExportExpenseBillDetail) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLExportExpenseBillDetail) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *GLExportExpenseBillDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocumentDate

`func (o *GLExportExpenseBillDetail) GetDocumentDate() string`

GetDocumentDate returns the DocumentDate field if non-nil, zero value otherwise.

### GetDocumentDateOk

`func (o *GLExportExpenseBillDetail) GetDocumentDateOk() (*string, bool)`

GetDocumentDateOk returns a tuple with the DocumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDate

`func (o *GLExportExpenseBillDetail) SetDocumentDate(v string)`

SetDocumentDate sets DocumentDate field to given value.

### HasDocumentDate

`func (o *GLExportExpenseBillDetail) HasDocumentDate() bool`

HasDocumentDate returns a boolean if a field has been set.

### GetGlTypeId

`func (o *GLExportExpenseBillDetail) GetGlTypeId() string`

GetGlTypeId returns the GlTypeId field if non-nil, zero value otherwise.

### GetGlTypeIdOk

`func (o *GLExportExpenseBillDetail) GetGlTypeIdOk() (*string, bool)`

GetGlTypeIdOk returns a tuple with the GlTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlTypeId

`func (o *GLExportExpenseBillDetail) SetGlTypeId(v string)`

SetGlTypeId sets GlTypeId field to given value.

### HasGlTypeId

`func (o *GLExportExpenseBillDetail) HasGlTypeId() bool`

HasGlTypeId returns a boolean if a field has been set.

### GetMemo

`func (o *GLExportExpenseBillDetail) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GLExportExpenseBillDetail) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GLExportExpenseBillDetail) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GLExportExpenseBillDetail) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetCompany

`func (o *GLExportExpenseBillDetail) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GLExportExpenseBillDetail) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GLExportExpenseBillDetail) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GLExportExpenseBillDetail) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAccountNumber

`func (o *GLExportExpenseBillDetail) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GLExportExpenseBillDetail) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GLExportExpenseBillDetail) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GLExportExpenseBillDetail) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetExpenseClass

`func (o *GLExportExpenseBillDetail) GetExpenseClass() ClassificationReference`

GetExpenseClass returns the ExpenseClass field if non-nil, zero value otherwise.

### GetExpenseClassOk

`func (o *GLExportExpenseBillDetail) GetExpenseClassOk() (*ClassificationReference, bool)`

GetExpenseClassOk returns a tuple with the ExpenseClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseClass

`func (o *GLExportExpenseBillDetail) SetExpenseClass(v ClassificationReference)`

SetExpenseClass sets ExpenseClass field to given value.

### HasExpenseClass

`func (o *GLExportExpenseBillDetail) HasExpenseClass() bool`

HasExpenseClass returns a boolean if a field has been set.

### GetCurrency

`func (o *GLExportExpenseBillDetail) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GLExportExpenseBillDetail) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GLExportExpenseBillDetail) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GLExportExpenseBillDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotal

`func (o *GLExportExpenseBillDetail) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GLExportExpenseBillDetail) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GLExportExpenseBillDetail) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GLExportExpenseBillDetail) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GLExportExpenseBillDetail) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GLExportExpenseBillDetail) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetBillable

`func (o *GLExportExpenseBillDetail) GetBillable() bool`

GetBillable returns the Billable field if non-nil, zero value otherwise.

### GetBillableOk

`func (o *GLExportExpenseBillDetail) GetBillableOk() (*bool, bool)`

GetBillableOk returns a tuple with the Billable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillable

`func (o *GLExportExpenseBillDetail) SetBillable(v bool)`

SetBillable sets Billable field to given value.

### HasBillable

`func (o *GLExportExpenseBillDetail) HasBillable() bool`

HasBillable returns a boolean if a field has been set.

### SetBillableNil

`func (o *GLExportExpenseBillDetail) SetBillableNil(b bool)`

 SetBillableNil sets the value for Billable to be an explicit nil

### UnsetBillable
`func (o *GLExportExpenseBillDetail) UnsetBillable()`

UnsetBillable ensures that no value is present for Billable, not even an explicit nil
### GetReimbursable

`func (o *GLExportExpenseBillDetail) GetReimbursable() bool`

GetReimbursable returns the Reimbursable field if non-nil, zero value otherwise.

### GetReimbursableOk

`func (o *GLExportExpenseBillDetail) GetReimbursableOk() (*bool, bool)`

GetReimbursableOk returns a tuple with the Reimbursable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReimbursable

`func (o *GLExportExpenseBillDetail) SetReimbursable(v bool)`

SetReimbursable sets Reimbursable field to given value.

### HasReimbursable

`func (o *GLExportExpenseBillDetail) HasReimbursable() bool`

HasReimbursable returns a boolean if a field has been set.

### SetReimbursableNil

`func (o *GLExportExpenseBillDetail) SetReimbursableNil(b bool)`

 SetReimbursableNil sets the value for Reimbursable to be an explicit nil

### UnsetReimbursable
`func (o *GLExportExpenseBillDetail) UnsetReimbursable()`

UnsetReimbursable ensures that no value is present for Reimbursable, not even an explicit nil
### GetCompanyAdvance

`func (o *GLExportExpenseBillDetail) GetCompanyAdvance() bool`

GetCompanyAdvance returns the CompanyAdvance field if non-nil, zero value otherwise.

### GetCompanyAdvanceOk

`func (o *GLExportExpenseBillDetail) GetCompanyAdvanceOk() (*bool, bool)`

GetCompanyAdvanceOk returns a tuple with the CompanyAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAdvance

`func (o *GLExportExpenseBillDetail) SetCompanyAdvance(v bool)`

SetCompanyAdvance sets CompanyAdvance field to given value.

### HasCompanyAdvance

`func (o *GLExportExpenseBillDetail) HasCompanyAdvance() bool`

HasCompanyAdvance returns a boolean if a field has been set.

### SetCompanyAdvanceNil

`func (o *GLExportExpenseBillDetail) SetCompanyAdvanceNil(b bool)`

 SetCompanyAdvanceNil sets the value for CompanyAdvance to be an explicit nil

### UnsetCompanyAdvance
`func (o *GLExportExpenseBillDetail) UnsetCompanyAdvance()`

UnsetCompanyAdvance ensures that no value is present for CompanyAdvance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


