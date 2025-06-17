# GLExportExpenseOffset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**DocumentDate** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**GlTypeId** | Pointer to **string** |  | [optional] 
**GlClass** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewGLExportExpenseOffset

`func NewGLExportExpenseOffset() *GLExportExpenseOffset`

NewGLExportExpenseOffset instantiates a new GLExportExpenseOffset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportExpenseOffsetWithDefaults

`func NewGLExportExpenseOffsetWithDefaults() *GLExportExpenseOffset`

NewGLExportExpenseOffsetWithDefaults instantiates a new GLExportExpenseOffset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLExportExpenseOffset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLExportExpenseOffset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLExportExpenseOffset) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GLExportExpenseOffset) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GLExportExpenseOffset) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GLExportExpenseOffset) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDocumentDate

`func (o *GLExportExpenseOffset) GetDocumentDate() string`

GetDocumentDate returns the DocumentDate field if non-nil, zero value otherwise.

### GetDocumentDateOk

`func (o *GLExportExpenseOffset) GetDocumentDateOk() (*string, bool)`

GetDocumentDateOk returns a tuple with the DocumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDate

`func (o *GLExportExpenseOffset) SetDocumentDate(v string)`

SetDocumentDate sets DocumentDate field to given value.

### HasDocumentDate

`func (o *GLExportExpenseOffset) HasDocumentDate() bool`

HasDocumentDate returns a boolean if a field has been set.

### GetDocumentType

`func (o *GLExportExpenseOffset) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *GLExportExpenseOffset) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *GLExportExpenseOffset) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *GLExportExpenseOffset) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetAccountNumber

`func (o *GLExportExpenseOffset) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GLExportExpenseOffset) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GLExportExpenseOffset) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GLExportExpenseOffset) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetGlTypeId

`func (o *GLExportExpenseOffset) GetGlTypeId() string`

GetGlTypeId returns the GlTypeId field if non-nil, zero value otherwise.

### GetGlTypeIdOk

`func (o *GLExportExpenseOffset) GetGlTypeIdOk() (*string, bool)`

GetGlTypeIdOk returns a tuple with the GlTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlTypeId

`func (o *GLExportExpenseOffset) SetGlTypeId(v string)`

SetGlTypeId sets GlTypeId field to given value.

### HasGlTypeId

`func (o *GLExportExpenseOffset) HasGlTypeId() bool`

HasGlTypeId returns a boolean if a field has been set.

### GetGlClass

`func (o *GLExportExpenseOffset) GetGlClass() string`

GetGlClass returns the GlClass field if non-nil, zero value otherwise.

### GetGlClassOk

`func (o *GLExportExpenseOffset) GetGlClassOk() (*string, bool)`

GetGlClassOk returns a tuple with the GlClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlClass

`func (o *GLExportExpenseOffset) SetGlClass(v string)`

SetGlClass sets GlClass field to given value.

### HasGlClass

`func (o *GLExportExpenseOffset) HasGlClass() bool`

HasGlClass returns a boolean if a field has been set.

### GetMember

`func (o *GLExportExpenseOffset) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GLExportExpenseOffset) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GLExportExpenseOffset) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *GLExportExpenseOffset) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetMemo

`func (o *GLExportExpenseOffset) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GLExportExpenseOffset) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GLExportExpenseOffset) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GLExportExpenseOffset) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetDescription

`func (o *GLExportExpenseOffset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GLExportExpenseOffset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GLExportExpenseOffset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GLExportExpenseOffset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTotal

`func (o *GLExportExpenseOffset) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GLExportExpenseOffset) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GLExportExpenseOffset) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GLExportExpenseOffset) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GLExportExpenseOffset) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GLExportExpenseOffset) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


