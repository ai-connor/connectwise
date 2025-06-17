# OpportunityToSalesOrderConversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SalesOrderId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IncludeAllNotesFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeAllDocumentsFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeAllProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeNoteIds** | Pointer to **[]int32** |  | [optional] 
**IncludeDocumentIds** | Pointer to **[]int32** |  | [optional] 
**IncludeProductIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewOpportunityToSalesOrderConversion

`func NewOpportunityToSalesOrderConversion() *OpportunityToSalesOrderConversion`

NewOpportunityToSalesOrderConversion instantiates a new OpportunityToSalesOrderConversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityToSalesOrderConversionWithDefaults

`func NewOpportunityToSalesOrderConversionWithDefaults() *OpportunityToSalesOrderConversion`

NewOpportunityToSalesOrderConversionWithDefaults instantiates a new OpportunityToSalesOrderConversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSalesOrderId

`func (o *OpportunityToSalesOrderConversion) GetSalesOrderId() int32`

GetSalesOrderId returns the SalesOrderId field if non-nil, zero value otherwise.

### GetSalesOrderIdOk

`func (o *OpportunityToSalesOrderConversion) GetSalesOrderIdOk() (*int32, bool)`

GetSalesOrderIdOk returns a tuple with the SalesOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrderId

`func (o *OpportunityToSalesOrderConversion) SetSalesOrderId(v int32)`

SetSalesOrderId sets SalesOrderId field to given value.

### HasSalesOrderId

`func (o *OpportunityToSalesOrderConversion) HasSalesOrderId() bool`

HasSalesOrderId returns a boolean if a field has been set.

### GetName

`func (o *OpportunityToSalesOrderConversion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpportunityToSalesOrderConversion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpportunityToSalesOrderConversion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpportunityToSalesOrderConversion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIncludeAllNotesFlag

`func (o *OpportunityToSalesOrderConversion) GetIncludeAllNotesFlag() bool`

GetIncludeAllNotesFlag returns the IncludeAllNotesFlag field if non-nil, zero value otherwise.

### GetIncludeAllNotesFlagOk

`func (o *OpportunityToSalesOrderConversion) GetIncludeAllNotesFlagOk() (*bool, bool)`

GetIncludeAllNotesFlagOk returns a tuple with the IncludeAllNotesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllNotesFlag

`func (o *OpportunityToSalesOrderConversion) SetIncludeAllNotesFlag(v bool)`

SetIncludeAllNotesFlag sets IncludeAllNotesFlag field to given value.

### HasIncludeAllNotesFlag

`func (o *OpportunityToSalesOrderConversion) HasIncludeAllNotesFlag() bool`

HasIncludeAllNotesFlag returns a boolean if a field has been set.

### SetIncludeAllNotesFlagNil

`func (o *OpportunityToSalesOrderConversion) SetIncludeAllNotesFlagNil(b bool)`

 SetIncludeAllNotesFlagNil sets the value for IncludeAllNotesFlag to be an explicit nil

### UnsetIncludeAllNotesFlag
`func (o *OpportunityToSalesOrderConversion) UnsetIncludeAllNotesFlag()`

UnsetIncludeAllNotesFlag ensures that no value is present for IncludeAllNotesFlag, not even an explicit nil
### GetIncludeAllDocumentsFlag

`func (o *OpportunityToSalesOrderConversion) GetIncludeAllDocumentsFlag() bool`

GetIncludeAllDocumentsFlag returns the IncludeAllDocumentsFlag field if non-nil, zero value otherwise.

### GetIncludeAllDocumentsFlagOk

`func (o *OpportunityToSalesOrderConversion) GetIncludeAllDocumentsFlagOk() (*bool, bool)`

GetIncludeAllDocumentsFlagOk returns a tuple with the IncludeAllDocumentsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllDocumentsFlag

`func (o *OpportunityToSalesOrderConversion) SetIncludeAllDocumentsFlag(v bool)`

SetIncludeAllDocumentsFlag sets IncludeAllDocumentsFlag field to given value.

### HasIncludeAllDocumentsFlag

`func (o *OpportunityToSalesOrderConversion) HasIncludeAllDocumentsFlag() bool`

HasIncludeAllDocumentsFlag returns a boolean if a field has been set.

### SetIncludeAllDocumentsFlagNil

`func (o *OpportunityToSalesOrderConversion) SetIncludeAllDocumentsFlagNil(b bool)`

 SetIncludeAllDocumentsFlagNil sets the value for IncludeAllDocumentsFlag to be an explicit nil

### UnsetIncludeAllDocumentsFlag
`func (o *OpportunityToSalesOrderConversion) UnsetIncludeAllDocumentsFlag()`

UnsetIncludeAllDocumentsFlag ensures that no value is present for IncludeAllDocumentsFlag, not even an explicit nil
### GetIncludeAllProductsFlag

`func (o *OpportunityToSalesOrderConversion) GetIncludeAllProductsFlag() bool`

GetIncludeAllProductsFlag returns the IncludeAllProductsFlag field if non-nil, zero value otherwise.

### GetIncludeAllProductsFlagOk

`func (o *OpportunityToSalesOrderConversion) GetIncludeAllProductsFlagOk() (*bool, bool)`

GetIncludeAllProductsFlagOk returns a tuple with the IncludeAllProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllProductsFlag

`func (o *OpportunityToSalesOrderConversion) SetIncludeAllProductsFlag(v bool)`

SetIncludeAllProductsFlag sets IncludeAllProductsFlag field to given value.

### HasIncludeAllProductsFlag

`func (o *OpportunityToSalesOrderConversion) HasIncludeAllProductsFlag() bool`

HasIncludeAllProductsFlag returns a boolean if a field has been set.

### SetIncludeAllProductsFlagNil

`func (o *OpportunityToSalesOrderConversion) SetIncludeAllProductsFlagNil(b bool)`

 SetIncludeAllProductsFlagNil sets the value for IncludeAllProductsFlag to be an explicit nil

### UnsetIncludeAllProductsFlag
`func (o *OpportunityToSalesOrderConversion) UnsetIncludeAllProductsFlag()`

UnsetIncludeAllProductsFlag ensures that no value is present for IncludeAllProductsFlag, not even an explicit nil
### GetIncludeNoteIds

`func (o *OpportunityToSalesOrderConversion) GetIncludeNoteIds() []int32`

GetIncludeNoteIds returns the IncludeNoteIds field if non-nil, zero value otherwise.

### GetIncludeNoteIdsOk

`func (o *OpportunityToSalesOrderConversion) GetIncludeNoteIdsOk() (*[]int32, bool)`

GetIncludeNoteIdsOk returns a tuple with the IncludeNoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNoteIds

`func (o *OpportunityToSalesOrderConversion) SetIncludeNoteIds(v []int32)`

SetIncludeNoteIds sets IncludeNoteIds field to given value.

### HasIncludeNoteIds

`func (o *OpportunityToSalesOrderConversion) HasIncludeNoteIds() bool`

HasIncludeNoteIds returns a boolean if a field has been set.

### GetIncludeDocumentIds

`func (o *OpportunityToSalesOrderConversion) GetIncludeDocumentIds() []int32`

GetIncludeDocumentIds returns the IncludeDocumentIds field if non-nil, zero value otherwise.

### GetIncludeDocumentIdsOk

`func (o *OpportunityToSalesOrderConversion) GetIncludeDocumentIdsOk() (*[]int32, bool)`

GetIncludeDocumentIdsOk returns a tuple with the IncludeDocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDocumentIds

`func (o *OpportunityToSalesOrderConversion) SetIncludeDocumentIds(v []int32)`

SetIncludeDocumentIds sets IncludeDocumentIds field to given value.

### HasIncludeDocumentIds

`func (o *OpportunityToSalesOrderConversion) HasIncludeDocumentIds() bool`

HasIncludeDocumentIds returns a boolean if a field has been set.

### GetIncludeProductIds

`func (o *OpportunityToSalesOrderConversion) GetIncludeProductIds() []int32`

GetIncludeProductIds returns the IncludeProductIds field if non-nil, zero value otherwise.

### GetIncludeProductIdsOk

`func (o *OpportunityToSalesOrderConversion) GetIncludeProductIdsOk() (*[]int32, bool)`

GetIncludeProductIdsOk returns a tuple with the IncludeProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductIds

`func (o *OpportunityToSalesOrderConversion) SetIncludeProductIds(v []int32)`

SetIncludeProductIds sets IncludeProductIds field to given value.

### HasIncludeProductIds

`func (o *OpportunityToSalesOrderConversion) HasIncludeProductIds() bool`

HasIncludeProductIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


