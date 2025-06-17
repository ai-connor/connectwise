# OpportunityToServiceTicketConversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TicketId** | Pointer to **int32** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**IncludeAllNotesFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeAllDocumentsFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeAllProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeNoteIds** | Pointer to **[]int32** |  | [optional] 
**IncludeDocumentIds** | Pointer to **[]int32** |  | [optional] 
**IncludeProductIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewOpportunityToServiceTicketConversion

`func NewOpportunityToServiceTicketConversion() *OpportunityToServiceTicketConversion`

NewOpportunityToServiceTicketConversion instantiates a new OpportunityToServiceTicketConversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityToServiceTicketConversionWithDefaults

`func NewOpportunityToServiceTicketConversionWithDefaults() *OpportunityToServiceTicketConversion`

NewOpportunityToServiceTicketConversionWithDefaults instantiates a new OpportunityToServiceTicketConversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicketId

`func (o *OpportunityToServiceTicketConversion) GetTicketId() int32`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *OpportunityToServiceTicketConversion) GetTicketIdOk() (*int32, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *OpportunityToServiceTicketConversion) SetTicketId(v int32)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *OpportunityToServiceTicketConversion) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### GetSummary

`func (o *OpportunityToServiceTicketConversion) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *OpportunityToServiceTicketConversion) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *OpportunityToServiceTicketConversion) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *OpportunityToServiceTicketConversion) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetIncludeAllNotesFlag

`func (o *OpportunityToServiceTicketConversion) GetIncludeAllNotesFlag() bool`

GetIncludeAllNotesFlag returns the IncludeAllNotesFlag field if non-nil, zero value otherwise.

### GetIncludeAllNotesFlagOk

`func (o *OpportunityToServiceTicketConversion) GetIncludeAllNotesFlagOk() (*bool, bool)`

GetIncludeAllNotesFlagOk returns a tuple with the IncludeAllNotesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllNotesFlag

`func (o *OpportunityToServiceTicketConversion) SetIncludeAllNotesFlag(v bool)`

SetIncludeAllNotesFlag sets IncludeAllNotesFlag field to given value.

### HasIncludeAllNotesFlag

`func (o *OpportunityToServiceTicketConversion) HasIncludeAllNotesFlag() bool`

HasIncludeAllNotesFlag returns a boolean if a field has been set.

### SetIncludeAllNotesFlagNil

`func (o *OpportunityToServiceTicketConversion) SetIncludeAllNotesFlagNil(b bool)`

 SetIncludeAllNotesFlagNil sets the value for IncludeAllNotesFlag to be an explicit nil

### UnsetIncludeAllNotesFlag
`func (o *OpportunityToServiceTicketConversion) UnsetIncludeAllNotesFlag()`

UnsetIncludeAllNotesFlag ensures that no value is present for IncludeAllNotesFlag, not even an explicit nil
### GetIncludeAllDocumentsFlag

`func (o *OpportunityToServiceTicketConversion) GetIncludeAllDocumentsFlag() bool`

GetIncludeAllDocumentsFlag returns the IncludeAllDocumentsFlag field if non-nil, zero value otherwise.

### GetIncludeAllDocumentsFlagOk

`func (o *OpportunityToServiceTicketConversion) GetIncludeAllDocumentsFlagOk() (*bool, bool)`

GetIncludeAllDocumentsFlagOk returns a tuple with the IncludeAllDocumentsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllDocumentsFlag

`func (o *OpportunityToServiceTicketConversion) SetIncludeAllDocumentsFlag(v bool)`

SetIncludeAllDocumentsFlag sets IncludeAllDocumentsFlag field to given value.

### HasIncludeAllDocumentsFlag

`func (o *OpportunityToServiceTicketConversion) HasIncludeAllDocumentsFlag() bool`

HasIncludeAllDocumentsFlag returns a boolean if a field has been set.

### SetIncludeAllDocumentsFlagNil

`func (o *OpportunityToServiceTicketConversion) SetIncludeAllDocumentsFlagNil(b bool)`

 SetIncludeAllDocumentsFlagNil sets the value for IncludeAllDocumentsFlag to be an explicit nil

### UnsetIncludeAllDocumentsFlag
`func (o *OpportunityToServiceTicketConversion) UnsetIncludeAllDocumentsFlag()`

UnsetIncludeAllDocumentsFlag ensures that no value is present for IncludeAllDocumentsFlag, not even an explicit nil
### GetIncludeAllProductsFlag

`func (o *OpportunityToServiceTicketConversion) GetIncludeAllProductsFlag() bool`

GetIncludeAllProductsFlag returns the IncludeAllProductsFlag field if non-nil, zero value otherwise.

### GetIncludeAllProductsFlagOk

`func (o *OpportunityToServiceTicketConversion) GetIncludeAllProductsFlagOk() (*bool, bool)`

GetIncludeAllProductsFlagOk returns a tuple with the IncludeAllProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllProductsFlag

`func (o *OpportunityToServiceTicketConversion) SetIncludeAllProductsFlag(v bool)`

SetIncludeAllProductsFlag sets IncludeAllProductsFlag field to given value.

### HasIncludeAllProductsFlag

`func (o *OpportunityToServiceTicketConversion) HasIncludeAllProductsFlag() bool`

HasIncludeAllProductsFlag returns a boolean if a field has been set.

### SetIncludeAllProductsFlagNil

`func (o *OpportunityToServiceTicketConversion) SetIncludeAllProductsFlagNil(b bool)`

 SetIncludeAllProductsFlagNil sets the value for IncludeAllProductsFlag to be an explicit nil

### UnsetIncludeAllProductsFlag
`func (o *OpportunityToServiceTicketConversion) UnsetIncludeAllProductsFlag()`

UnsetIncludeAllProductsFlag ensures that no value is present for IncludeAllProductsFlag, not even an explicit nil
### GetIncludeNoteIds

`func (o *OpportunityToServiceTicketConversion) GetIncludeNoteIds() []int32`

GetIncludeNoteIds returns the IncludeNoteIds field if non-nil, zero value otherwise.

### GetIncludeNoteIdsOk

`func (o *OpportunityToServiceTicketConversion) GetIncludeNoteIdsOk() (*[]int32, bool)`

GetIncludeNoteIdsOk returns a tuple with the IncludeNoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNoteIds

`func (o *OpportunityToServiceTicketConversion) SetIncludeNoteIds(v []int32)`

SetIncludeNoteIds sets IncludeNoteIds field to given value.

### HasIncludeNoteIds

`func (o *OpportunityToServiceTicketConversion) HasIncludeNoteIds() bool`

HasIncludeNoteIds returns a boolean if a field has been set.

### GetIncludeDocumentIds

`func (o *OpportunityToServiceTicketConversion) GetIncludeDocumentIds() []int32`

GetIncludeDocumentIds returns the IncludeDocumentIds field if non-nil, zero value otherwise.

### GetIncludeDocumentIdsOk

`func (o *OpportunityToServiceTicketConversion) GetIncludeDocumentIdsOk() (*[]int32, bool)`

GetIncludeDocumentIdsOk returns a tuple with the IncludeDocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDocumentIds

`func (o *OpportunityToServiceTicketConversion) SetIncludeDocumentIds(v []int32)`

SetIncludeDocumentIds sets IncludeDocumentIds field to given value.

### HasIncludeDocumentIds

`func (o *OpportunityToServiceTicketConversion) HasIncludeDocumentIds() bool`

HasIncludeDocumentIds returns a boolean if a field has been set.

### GetIncludeProductIds

`func (o *OpportunityToServiceTicketConversion) GetIncludeProductIds() []int32`

GetIncludeProductIds returns the IncludeProductIds field if non-nil, zero value otherwise.

### GetIncludeProductIdsOk

`func (o *OpportunityToServiceTicketConversion) GetIncludeProductIdsOk() (*[]int32, bool)`

GetIncludeProductIdsOk returns a tuple with the IncludeProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductIds

`func (o *OpportunityToServiceTicketConversion) SetIncludeProductIds(v []int32)`

SetIncludeProductIds sets IncludeProductIds field to given value.

### HasIncludeProductIds

`func (o *OpportunityToServiceTicketConversion) HasIncludeProductIds() bool`

HasIncludeProductIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


