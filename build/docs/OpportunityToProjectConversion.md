# OpportunityToProjectConversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ProjectStatusReference**](ProjectStatusReference.md) |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  | [optional] 
**Board** | Pointer to [**ProjectBoardReference**](ProjectBoardReference.md) |  | [optional] 
**Manager** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**EstimatedStart** | Pointer to **string** |  | [optional] 
**EstimatedEnd** | Pointer to **string** |  | [optional] 
**IncludeAllNotesFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeAllDocumentsFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeAllProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeNoteIds** | Pointer to **[]int32** |  | [optional] 
**IncludeDocumentIds** | Pointer to **[]int32** |  | [optional] 
**IncludeProductIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewOpportunityToProjectConversion

`func NewOpportunityToProjectConversion() *OpportunityToProjectConversion`

NewOpportunityToProjectConversion instantiates a new OpportunityToProjectConversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityToProjectConversionWithDefaults

`func NewOpportunityToProjectConversionWithDefaults() *OpportunityToProjectConversion`

NewOpportunityToProjectConversionWithDefaults instantiates a new OpportunityToProjectConversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *OpportunityToProjectConversion) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OpportunityToProjectConversion) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OpportunityToProjectConversion) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *OpportunityToProjectConversion) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetName

`func (o *OpportunityToProjectConversion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpportunityToProjectConversion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpportunityToProjectConversion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpportunityToProjectConversion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *OpportunityToProjectConversion) GetStatus() ProjectStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpportunityToProjectConversion) GetStatusOk() (*ProjectStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpportunityToProjectConversion) SetStatus(v ProjectStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpportunityToProjectConversion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLocationId

`func (o *OpportunityToProjectConversion) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *OpportunityToProjectConversion) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *OpportunityToProjectConversion) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *OpportunityToProjectConversion) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *OpportunityToProjectConversion) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *OpportunityToProjectConversion) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetBusinessUnitId

`func (o *OpportunityToProjectConversion) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *OpportunityToProjectConversion) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *OpportunityToProjectConversion) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *OpportunityToProjectConversion) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *OpportunityToProjectConversion) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *OpportunityToProjectConversion) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetBoard

`func (o *OpportunityToProjectConversion) GetBoard() ProjectBoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *OpportunityToProjectConversion) GetBoardOk() (*ProjectBoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *OpportunityToProjectConversion) SetBoard(v ProjectBoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *OpportunityToProjectConversion) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetManager

`func (o *OpportunityToProjectConversion) GetManager() MemberReference`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *OpportunityToProjectConversion) GetManagerOk() (*MemberReference, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *OpportunityToProjectConversion) SetManager(v MemberReference)`

SetManager sets Manager field to given value.

### HasManager

`func (o *OpportunityToProjectConversion) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetEstimatedStart

`func (o *OpportunityToProjectConversion) GetEstimatedStart() string`

GetEstimatedStart returns the EstimatedStart field if non-nil, zero value otherwise.

### GetEstimatedStartOk

`func (o *OpportunityToProjectConversion) GetEstimatedStartOk() (*string, bool)`

GetEstimatedStartOk returns a tuple with the EstimatedStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStart

`func (o *OpportunityToProjectConversion) SetEstimatedStart(v string)`

SetEstimatedStart sets EstimatedStart field to given value.

### HasEstimatedStart

`func (o *OpportunityToProjectConversion) HasEstimatedStart() bool`

HasEstimatedStart returns a boolean if a field has been set.

### GetEstimatedEnd

`func (o *OpportunityToProjectConversion) GetEstimatedEnd() string`

GetEstimatedEnd returns the EstimatedEnd field if non-nil, zero value otherwise.

### GetEstimatedEndOk

`func (o *OpportunityToProjectConversion) GetEstimatedEndOk() (*string, bool)`

GetEstimatedEndOk returns a tuple with the EstimatedEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedEnd

`func (o *OpportunityToProjectConversion) SetEstimatedEnd(v string)`

SetEstimatedEnd sets EstimatedEnd field to given value.

### HasEstimatedEnd

`func (o *OpportunityToProjectConversion) HasEstimatedEnd() bool`

HasEstimatedEnd returns a boolean if a field has been set.

### GetIncludeAllNotesFlag

`func (o *OpportunityToProjectConversion) GetIncludeAllNotesFlag() bool`

GetIncludeAllNotesFlag returns the IncludeAllNotesFlag field if non-nil, zero value otherwise.

### GetIncludeAllNotesFlagOk

`func (o *OpportunityToProjectConversion) GetIncludeAllNotesFlagOk() (*bool, bool)`

GetIncludeAllNotesFlagOk returns a tuple with the IncludeAllNotesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllNotesFlag

`func (o *OpportunityToProjectConversion) SetIncludeAllNotesFlag(v bool)`

SetIncludeAllNotesFlag sets IncludeAllNotesFlag field to given value.

### HasIncludeAllNotesFlag

`func (o *OpportunityToProjectConversion) HasIncludeAllNotesFlag() bool`

HasIncludeAllNotesFlag returns a boolean if a field has been set.

### SetIncludeAllNotesFlagNil

`func (o *OpportunityToProjectConversion) SetIncludeAllNotesFlagNil(b bool)`

 SetIncludeAllNotesFlagNil sets the value for IncludeAllNotesFlag to be an explicit nil

### UnsetIncludeAllNotesFlag
`func (o *OpportunityToProjectConversion) UnsetIncludeAllNotesFlag()`

UnsetIncludeAllNotesFlag ensures that no value is present for IncludeAllNotesFlag, not even an explicit nil
### GetIncludeAllDocumentsFlag

`func (o *OpportunityToProjectConversion) GetIncludeAllDocumentsFlag() bool`

GetIncludeAllDocumentsFlag returns the IncludeAllDocumentsFlag field if non-nil, zero value otherwise.

### GetIncludeAllDocumentsFlagOk

`func (o *OpportunityToProjectConversion) GetIncludeAllDocumentsFlagOk() (*bool, bool)`

GetIncludeAllDocumentsFlagOk returns a tuple with the IncludeAllDocumentsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllDocumentsFlag

`func (o *OpportunityToProjectConversion) SetIncludeAllDocumentsFlag(v bool)`

SetIncludeAllDocumentsFlag sets IncludeAllDocumentsFlag field to given value.

### HasIncludeAllDocumentsFlag

`func (o *OpportunityToProjectConversion) HasIncludeAllDocumentsFlag() bool`

HasIncludeAllDocumentsFlag returns a boolean if a field has been set.

### SetIncludeAllDocumentsFlagNil

`func (o *OpportunityToProjectConversion) SetIncludeAllDocumentsFlagNil(b bool)`

 SetIncludeAllDocumentsFlagNil sets the value for IncludeAllDocumentsFlag to be an explicit nil

### UnsetIncludeAllDocumentsFlag
`func (o *OpportunityToProjectConversion) UnsetIncludeAllDocumentsFlag()`

UnsetIncludeAllDocumentsFlag ensures that no value is present for IncludeAllDocumentsFlag, not even an explicit nil
### GetIncludeAllProductsFlag

`func (o *OpportunityToProjectConversion) GetIncludeAllProductsFlag() bool`

GetIncludeAllProductsFlag returns the IncludeAllProductsFlag field if non-nil, zero value otherwise.

### GetIncludeAllProductsFlagOk

`func (o *OpportunityToProjectConversion) GetIncludeAllProductsFlagOk() (*bool, bool)`

GetIncludeAllProductsFlagOk returns a tuple with the IncludeAllProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllProductsFlag

`func (o *OpportunityToProjectConversion) SetIncludeAllProductsFlag(v bool)`

SetIncludeAllProductsFlag sets IncludeAllProductsFlag field to given value.

### HasIncludeAllProductsFlag

`func (o *OpportunityToProjectConversion) HasIncludeAllProductsFlag() bool`

HasIncludeAllProductsFlag returns a boolean if a field has been set.

### SetIncludeAllProductsFlagNil

`func (o *OpportunityToProjectConversion) SetIncludeAllProductsFlagNil(b bool)`

 SetIncludeAllProductsFlagNil sets the value for IncludeAllProductsFlag to be an explicit nil

### UnsetIncludeAllProductsFlag
`func (o *OpportunityToProjectConversion) UnsetIncludeAllProductsFlag()`

UnsetIncludeAllProductsFlag ensures that no value is present for IncludeAllProductsFlag, not even an explicit nil
### GetIncludeNoteIds

`func (o *OpportunityToProjectConversion) GetIncludeNoteIds() []int32`

GetIncludeNoteIds returns the IncludeNoteIds field if non-nil, zero value otherwise.

### GetIncludeNoteIdsOk

`func (o *OpportunityToProjectConversion) GetIncludeNoteIdsOk() (*[]int32, bool)`

GetIncludeNoteIdsOk returns a tuple with the IncludeNoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNoteIds

`func (o *OpportunityToProjectConversion) SetIncludeNoteIds(v []int32)`

SetIncludeNoteIds sets IncludeNoteIds field to given value.

### HasIncludeNoteIds

`func (o *OpportunityToProjectConversion) HasIncludeNoteIds() bool`

HasIncludeNoteIds returns a boolean if a field has been set.

### GetIncludeDocumentIds

`func (o *OpportunityToProjectConversion) GetIncludeDocumentIds() []int32`

GetIncludeDocumentIds returns the IncludeDocumentIds field if non-nil, zero value otherwise.

### GetIncludeDocumentIdsOk

`func (o *OpportunityToProjectConversion) GetIncludeDocumentIdsOk() (*[]int32, bool)`

GetIncludeDocumentIdsOk returns a tuple with the IncludeDocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDocumentIds

`func (o *OpportunityToProjectConversion) SetIncludeDocumentIds(v []int32)`

SetIncludeDocumentIds sets IncludeDocumentIds field to given value.

### HasIncludeDocumentIds

`func (o *OpportunityToProjectConversion) HasIncludeDocumentIds() bool`

HasIncludeDocumentIds returns a boolean if a field has been set.

### GetIncludeProductIds

`func (o *OpportunityToProjectConversion) GetIncludeProductIds() []int32`

GetIncludeProductIds returns the IncludeProductIds field if non-nil, zero value otherwise.

### GetIncludeProductIdsOk

`func (o *OpportunityToProjectConversion) GetIncludeProductIdsOk() (*[]int32, bool)`

GetIncludeProductIdsOk returns a tuple with the IncludeProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductIds

`func (o *OpportunityToProjectConversion) SetIncludeProductIds(v []int32)`

SetIncludeProductIds sets IncludeProductIds field to given value.

### HasIncludeProductIds

`func (o *OpportunityToProjectConversion) HasIncludeProductIds() bool`

HasIncludeProductIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


