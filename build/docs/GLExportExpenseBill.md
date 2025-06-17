# GLExportExpenseBill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**DocumentDate** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 
**DocumentNumber** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**GlClass** | Pointer to **string** |  | [optional] 
**ApAccountNumber** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**VendorNumber** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**Detail** | Pointer to [**[]GLExportExpenseBillDetail**](GLExportExpenseBillDetail.md) |  | [optional] 

## Methods

### NewGLExportExpenseBill

`func NewGLExportExpenseBill() *GLExportExpenseBill`

NewGLExportExpenseBill instantiates a new GLExportExpenseBill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportExpenseBillWithDefaults

`func NewGLExportExpenseBillWithDefaults() *GLExportExpenseBill`

NewGLExportExpenseBillWithDefaults instantiates a new GLExportExpenseBill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLExportExpenseBill) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLExportExpenseBill) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLExportExpenseBill) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GLExportExpenseBill) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GLExportExpenseBill) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GLExportExpenseBill) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDocumentDate

`func (o *GLExportExpenseBill) GetDocumentDate() string`

GetDocumentDate returns the DocumentDate field if non-nil, zero value otherwise.

### GetDocumentDateOk

`func (o *GLExportExpenseBill) GetDocumentDateOk() (*string, bool)`

GetDocumentDateOk returns a tuple with the DocumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDate

`func (o *GLExportExpenseBill) SetDocumentDate(v string)`

SetDocumentDate sets DocumentDate field to given value.

### HasDocumentDate

`func (o *GLExportExpenseBill) HasDocumentDate() bool`

HasDocumentDate returns a boolean if a field has been set.

### GetDocumentType

`func (o *GLExportExpenseBill) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *GLExportExpenseBill) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *GLExportExpenseBill) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *GLExportExpenseBill) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetDocumentNumber

`func (o *GLExportExpenseBill) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *GLExportExpenseBill) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *GLExportExpenseBill) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *GLExportExpenseBill) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### GetMemo

`func (o *GLExportExpenseBill) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GLExportExpenseBill) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GLExportExpenseBill) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GLExportExpenseBill) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetGlClass

`func (o *GLExportExpenseBill) GetGlClass() string`

GetGlClass returns the GlClass field if non-nil, zero value otherwise.

### GetGlClassOk

`func (o *GLExportExpenseBill) GetGlClassOk() (*string, bool)`

GetGlClassOk returns a tuple with the GlClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlClass

`func (o *GLExportExpenseBill) SetGlClass(v string)`

SetGlClass sets GlClass field to given value.

### HasGlClass

`func (o *GLExportExpenseBill) HasGlClass() bool`

HasGlClass returns a boolean if a field has been set.

### GetApAccountNumber

`func (o *GLExportExpenseBill) GetApAccountNumber() string`

GetApAccountNumber returns the ApAccountNumber field if non-nil, zero value otherwise.

### GetApAccountNumberOk

`func (o *GLExportExpenseBill) GetApAccountNumberOk() (*string, bool)`

GetApAccountNumberOk returns a tuple with the ApAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApAccountNumber

`func (o *GLExportExpenseBill) SetApAccountNumber(v string)`

SetApAccountNumber sets ApAccountNumber field to given value.

### HasApAccountNumber

`func (o *GLExportExpenseBill) HasApAccountNumber() bool`

HasApAccountNumber returns a boolean if a field has been set.

### GetMember

`func (o *GLExportExpenseBill) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GLExportExpenseBill) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GLExportExpenseBill) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *GLExportExpenseBill) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetVendorNumber

`func (o *GLExportExpenseBill) GetVendorNumber() string`

GetVendorNumber returns the VendorNumber field if non-nil, zero value otherwise.

### GetVendorNumberOk

`func (o *GLExportExpenseBill) GetVendorNumberOk() (*string, bool)`

GetVendorNumberOk returns a tuple with the VendorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumber

`func (o *GLExportExpenseBill) SetVendorNumber(v string)`

SetVendorNumber sets VendorNumber field to given value.

### HasVendorNumber

`func (o *GLExportExpenseBill) HasVendorNumber() bool`

HasVendorNumber returns a boolean if a field has been set.

### GetCurrency

`func (o *GLExportExpenseBill) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GLExportExpenseBill) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GLExportExpenseBill) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GLExportExpenseBill) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotal

`func (o *GLExportExpenseBill) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GLExportExpenseBill) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GLExportExpenseBill) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GLExportExpenseBill) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GLExportExpenseBill) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GLExportExpenseBill) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetDetail

`func (o *GLExportExpenseBill) GetDetail() []GLExportExpenseBillDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GLExportExpenseBill) GetDetailOk() (*[]GLExportExpenseBillDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GLExportExpenseBill) SetDetail(v []GLExportExpenseBillDetail)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GLExportExpenseBill) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


