# GLExportAdjustmentTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 
**DocumentDate** | Pointer to **string** |  | [optional] 
**GlTypeID** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**GlClass** | Pointer to **string** |  | [optional] 
**AdjustmentDescription** | Pointer to **string** |  | [optional] 
**AdjustmentDetail** | Pointer to [**[]GLExportAdjustmentTransactionDetail**](GLExportAdjustmentTransactionDetail.md) |  | [optional] 

## Methods

### NewGLExportAdjustmentTransaction

`func NewGLExportAdjustmentTransaction() *GLExportAdjustmentTransaction`

NewGLExportAdjustmentTransaction instantiates a new GLExportAdjustmentTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportAdjustmentTransactionWithDefaults

`func NewGLExportAdjustmentTransactionWithDefaults() *GLExportAdjustmentTransaction`

NewGLExportAdjustmentTransactionWithDefaults instantiates a new GLExportAdjustmentTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLExportAdjustmentTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLExportAdjustmentTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLExportAdjustmentTransaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GLExportAdjustmentTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocumentType

`func (o *GLExportAdjustmentTransaction) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *GLExportAdjustmentTransaction) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *GLExportAdjustmentTransaction) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *GLExportAdjustmentTransaction) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetDocumentDate

`func (o *GLExportAdjustmentTransaction) GetDocumentDate() string`

GetDocumentDate returns the DocumentDate field if non-nil, zero value otherwise.

### GetDocumentDateOk

`func (o *GLExportAdjustmentTransaction) GetDocumentDateOk() (*string, bool)`

GetDocumentDateOk returns a tuple with the DocumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDate

`func (o *GLExportAdjustmentTransaction) SetDocumentDate(v string)`

SetDocumentDate sets DocumentDate field to given value.

### HasDocumentDate

`func (o *GLExportAdjustmentTransaction) HasDocumentDate() bool`

HasDocumentDate returns a boolean if a field has been set.

### GetGlTypeID

`func (o *GLExportAdjustmentTransaction) GetGlTypeID() string`

GetGlTypeID returns the GlTypeID field if non-nil, zero value otherwise.

### GetGlTypeIDOk

`func (o *GLExportAdjustmentTransaction) GetGlTypeIDOk() (*string, bool)`

GetGlTypeIDOk returns a tuple with the GlTypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlTypeID

`func (o *GLExportAdjustmentTransaction) SetGlTypeID(v string)`

SetGlTypeID sets GlTypeID field to given value.

### HasGlTypeID

`func (o *GLExportAdjustmentTransaction) HasGlTypeID() bool`

HasGlTypeID returns a boolean if a field has been set.

### GetAccountNumber

`func (o *GLExportAdjustmentTransaction) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GLExportAdjustmentTransaction) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GLExportAdjustmentTransaction) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GLExportAdjustmentTransaction) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetMemo

`func (o *GLExportAdjustmentTransaction) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GLExportAdjustmentTransaction) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GLExportAdjustmentTransaction) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GLExportAdjustmentTransaction) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetGlClass

`func (o *GLExportAdjustmentTransaction) GetGlClass() string`

GetGlClass returns the GlClass field if non-nil, zero value otherwise.

### GetGlClassOk

`func (o *GLExportAdjustmentTransaction) GetGlClassOk() (*string, bool)`

GetGlClassOk returns a tuple with the GlClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlClass

`func (o *GLExportAdjustmentTransaction) SetGlClass(v string)`

SetGlClass sets GlClass field to given value.

### HasGlClass

`func (o *GLExportAdjustmentTransaction) HasGlClass() bool`

HasGlClass returns a boolean if a field has been set.

### GetAdjustmentDescription

`func (o *GLExportAdjustmentTransaction) GetAdjustmentDescription() string`

GetAdjustmentDescription returns the AdjustmentDescription field if non-nil, zero value otherwise.

### GetAdjustmentDescriptionOk

`func (o *GLExportAdjustmentTransaction) GetAdjustmentDescriptionOk() (*string, bool)`

GetAdjustmentDescriptionOk returns a tuple with the AdjustmentDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentDescription

`func (o *GLExportAdjustmentTransaction) SetAdjustmentDescription(v string)`

SetAdjustmentDescription sets AdjustmentDescription field to given value.

### HasAdjustmentDescription

`func (o *GLExportAdjustmentTransaction) HasAdjustmentDescription() bool`

HasAdjustmentDescription returns a boolean if a field has been set.

### GetAdjustmentDetail

`func (o *GLExportAdjustmentTransaction) GetAdjustmentDetail() []GLExportAdjustmentTransactionDetail`

GetAdjustmentDetail returns the AdjustmentDetail field if non-nil, zero value otherwise.

### GetAdjustmentDetailOk

`func (o *GLExportAdjustmentTransaction) GetAdjustmentDetailOk() (*[]GLExportAdjustmentTransactionDetail, bool)`

GetAdjustmentDetailOk returns a tuple with the AdjustmentDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentDetail

`func (o *GLExportAdjustmentTransaction) SetAdjustmentDetail(v []GLExportAdjustmentTransactionDetail)`

SetAdjustmentDetail sets AdjustmentDetail field to given value.

### HasAdjustmentDetail

`func (o *GLExportAdjustmentTransaction) HasAdjustmentDetail() bool`

HasAdjustmentDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


