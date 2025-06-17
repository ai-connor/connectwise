# OpportunityToAgreementConversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreementId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**AgreementTypeReference**](AgreementTypeReference.md) |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**NoEndingDateFlag** | Pointer to **NullableBool** |  | [optional] 
**BillCycleId** | Pointer to **NullableInt32** |  | [optional] 
**BillOneTimeFlag** | Pointer to **NullableBool** |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  | [optional] 
**IncludeAllNotesFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeAllDocumentsFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeAllProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludeNoteIds** | Pointer to **[]int32** |  | [optional] 
**IncludeDocumentIds** | Pointer to **[]int32** |  | [optional] 
**IncludeProductIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewOpportunityToAgreementConversion

`func NewOpportunityToAgreementConversion() *OpportunityToAgreementConversion`

NewOpportunityToAgreementConversion instantiates a new OpportunityToAgreementConversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityToAgreementConversionWithDefaults

`func NewOpportunityToAgreementConversionWithDefaults() *OpportunityToAgreementConversion`

NewOpportunityToAgreementConversionWithDefaults instantiates a new OpportunityToAgreementConversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreementId

`func (o *OpportunityToAgreementConversion) GetAgreementId() int32`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *OpportunityToAgreementConversion) GetAgreementIdOk() (*int32, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *OpportunityToAgreementConversion) SetAgreementId(v int32)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *OpportunityToAgreementConversion) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### GetName

`func (o *OpportunityToAgreementConversion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpportunityToAgreementConversion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpportunityToAgreementConversion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpportunityToAgreementConversion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *OpportunityToAgreementConversion) GetType() AgreementTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpportunityToAgreementConversion) GetTypeOk() (*AgreementTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpportunityToAgreementConversion) SetType(v AgreementTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *OpportunityToAgreementConversion) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStartDate

`func (o *OpportunityToAgreementConversion) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *OpportunityToAgreementConversion) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *OpportunityToAgreementConversion) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *OpportunityToAgreementConversion) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *OpportunityToAgreementConversion) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *OpportunityToAgreementConversion) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *OpportunityToAgreementConversion) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *OpportunityToAgreementConversion) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNoEndingDateFlag

`func (o *OpportunityToAgreementConversion) GetNoEndingDateFlag() bool`

GetNoEndingDateFlag returns the NoEndingDateFlag field if non-nil, zero value otherwise.

### GetNoEndingDateFlagOk

`func (o *OpportunityToAgreementConversion) GetNoEndingDateFlagOk() (*bool, bool)`

GetNoEndingDateFlagOk returns a tuple with the NoEndingDateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoEndingDateFlag

`func (o *OpportunityToAgreementConversion) SetNoEndingDateFlag(v bool)`

SetNoEndingDateFlag sets NoEndingDateFlag field to given value.

### HasNoEndingDateFlag

`func (o *OpportunityToAgreementConversion) HasNoEndingDateFlag() bool`

HasNoEndingDateFlag returns a boolean if a field has been set.

### SetNoEndingDateFlagNil

`func (o *OpportunityToAgreementConversion) SetNoEndingDateFlagNil(b bool)`

 SetNoEndingDateFlagNil sets the value for NoEndingDateFlag to be an explicit nil

### UnsetNoEndingDateFlag
`func (o *OpportunityToAgreementConversion) UnsetNoEndingDateFlag()`

UnsetNoEndingDateFlag ensures that no value is present for NoEndingDateFlag, not even an explicit nil
### GetBillCycleId

`func (o *OpportunityToAgreementConversion) GetBillCycleId() int32`

GetBillCycleId returns the BillCycleId field if non-nil, zero value otherwise.

### GetBillCycleIdOk

`func (o *OpportunityToAgreementConversion) GetBillCycleIdOk() (*int32, bool)`

GetBillCycleIdOk returns a tuple with the BillCycleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCycleId

`func (o *OpportunityToAgreementConversion) SetBillCycleId(v int32)`

SetBillCycleId sets BillCycleId field to given value.

### HasBillCycleId

`func (o *OpportunityToAgreementConversion) HasBillCycleId() bool`

HasBillCycleId returns a boolean if a field has been set.

### SetBillCycleIdNil

`func (o *OpportunityToAgreementConversion) SetBillCycleIdNil(b bool)`

 SetBillCycleIdNil sets the value for BillCycleId to be an explicit nil

### UnsetBillCycleId
`func (o *OpportunityToAgreementConversion) UnsetBillCycleId()`

UnsetBillCycleId ensures that no value is present for BillCycleId, not even an explicit nil
### GetBillOneTimeFlag

`func (o *OpportunityToAgreementConversion) GetBillOneTimeFlag() bool`

GetBillOneTimeFlag returns the BillOneTimeFlag field if non-nil, zero value otherwise.

### GetBillOneTimeFlagOk

`func (o *OpportunityToAgreementConversion) GetBillOneTimeFlagOk() (*bool, bool)`

GetBillOneTimeFlagOk returns a tuple with the BillOneTimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOneTimeFlag

`func (o *OpportunityToAgreementConversion) SetBillOneTimeFlag(v bool)`

SetBillOneTimeFlag sets BillOneTimeFlag field to given value.

### HasBillOneTimeFlag

`func (o *OpportunityToAgreementConversion) HasBillOneTimeFlag() bool`

HasBillOneTimeFlag returns a boolean if a field has been set.

### SetBillOneTimeFlagNil

`func (o *OpportunityToAgreementConversion) SetBillOneTimeFlagNil(b bool)`

 SetBillOneTimeFlagNil sets the value for BillOneTimeFlag to be an explicit nil

### UnsetBillOneTimeFlag
`func (o *OpportunityToAgreementConversion) UnsetBillOneTimeFlag()`

UnsetBillOneTimeFlag ensures that no value is present for BillOneTimeFlag, not even an explicit nil
### GetLocationId

`func (o *OpportunityToAgreementConversion) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *OpportunityToAgreementConversion) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *OpportunityToAgreementConversion) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *OpportunityToAgreementConversion) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *OpportunityToAgreementConversion) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *OpportunityToAgreementConversion) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetBusinessUnitId

`func (o *OpportunityToAgreementConversion) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *OpportunityToAgreementConversion) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *OpportunityToAgreementConversion) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *OpportunityToAgreementConversion) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *OpportunityToAgreementConversion) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *OpportunityToAgreementConversion) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetIncludeAllNotesFlag

`func (o *OpportunityToAgreementConversion) GetIncludeAllNotesFlag() bool`

GetIncludeAllNotesFlag returns the IncludeAllNotesFlag field if non-nil, zero value otherwise.

### GetIncludeAllNotesFlagOk

`func (o *OpportunityToAgreementConversion) GetIncludeAllNotesFlagOk() (*bool, bool)`

GetIncludeAllNotesFlagOk returns a tuple with the IncludeAllNotesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllNotesFlag

`func (o *OpportunityToAgreementConversion) SetIncludeAllNotesFlag(v bool)`

SetIncludeAllNotesFlag sets IncludeAllNotesFlag field to given value.

### HasIncludeAllNotesFlag

`func (o *OpportunityToAgreementConversion) HasIncludeAllNotesFlag() bool`

HasIncludeAllNotesFlag returns a boolean if a field has been set.

### SetIncludeAllNotesFlagNil

`func (o *OpportunityToAgreementConversion) SetIncludeAllNotesFlagNil(b bool)`

 SetIncludeAllNotesFlagNil sets the value for IncludeAllNotesFlag to be an explicit nil

### UnsetIncludeAllNotesFlag
`func (o *OpportunityToAgreementConversion) UnsetIncludeAllNotesFlag()`

UnsetIncludeAllNotesFlag ensures that no value is present for IncludeAllNotesFlag, not even an explicit nil
### GetIncludeAllDocumentsFlag

`func (o *OpportunityToAgreementConversion) GetIncludeAllDocumentsFlag() bool`

GetIncludeAllDocumentsFlag returns the IncludeAllDocumentsFlag field if non-nil, zero value otherwise.

### GetIncludeAllDocumentsFlagOk

`func (o *OpportunityToAgreementConversion) GetIncludeAllDocumentsFlagOk() (*bool, bool)`

GetIncludeAllDocumentsFlagOk returns a tuple with the IncludeAllDocumentsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllDocumentsFlag

`func (o *OpportunityToAgreementConversion) SetIncludeAllDocumentsFlag(v bool)`

SetIncludeAllDocumentsFlag sets IncludeAllDocumentsFlag field to given value.

### HasIncludeAllDocumentsFlag

`func (o *OpportunityToAgreementConversion) HasIncludeAllDocumentsFlag() bool`

HasIncludeAllDocumentsFlag returns a boolean if a field has been set.

### SetIncludeAllDocumentsFlagNil

`func (o *OpportunityToAgreementConversion) SetIncludeAllDocumentsFlagNil(b bool)`

 SetIncludeAllDocumentsFlagNil sets the value for IncludeAllDocumentsFlag to be an explicit nil

### UnsetIncludeAllDocumentsFlag
`func (o *OpportunityToAgreementConversion) UnsetIncludeAllDocumentsFlag()`

UnsetIncludeAllDocumentsFlag ensures that no value is present for IncludeAllDocumentsFlag, not even an explicit nil
### GetIncludeAllProductsFlag

`func (o *OpportunityToAgreementConversion) GetIncludeAllProductsFlag() bool`

GetIncludeAllProductsFlag returns the IncludeAllProductsFlag field if non-nil, zero value otherwise.

### GetIncludeAllProductsFlagOk

`func (o *OpportunityToAgreementConversion) GetIncludeAllProductsFlagOk() (*bool, bool)`

GetIncludeAllProductsFlagOk returns a tuple with the IncludeAllProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllProductsFlag

`func (o *OpportunityToAgreementConversion) SetIncludeAllProductsFlag(v bool)`

SetIncludeAllProductsFlag sets IncludeAllProductsFlag field to given value.

### HasIncludeAllProductsFlag

`func (o *OpportunityToAgreementConversion) HasIncludeAllProductsFlag() bool`

HasIncludeAllProductsFlag returns a boolean if a field has been set.

### SetIncludeAllProductsFlagNil

`func (o *OpportunityToAgreementConversion) SetIncludeAllProductsFlagNil(b bool)`

 SetIncludeAllProductsFlagNil sets the value for IncludeAllProductsFlag to be an explicit nil

### UnsetIncludeAllProductsFlag
`func (o *OpportunityToAgreementConversion) UnsetIncludeAllProductsFlag()`

UnsetIncludeAllProductsFlag ensures that no value is present for IncludeAllProductsFlag, not even an explicit nil
### GetIncludeNoteIds

`func (o *OpportunityToAgreementConversion) GetIncludeNoteIds() []int32`

GetIncludeNoteIds returns the IncludeNoteIds field if non-nil, zero value otherwise.

### GetIncludeNoteIdsOk

`func (o *OpportunityToAgreementConversion) GetIncludeNoteIdsOk() (*[]int32, bool)`

GetIncludeNoteIdsOk returns a tuple with the IncludeNoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNoteIds

`func (o *OpportunityToAgreementConversion) SetIncludeNoteIds(v []int32)`

SetIncludeNoteIds sets IncludeNoteIds field to given value.

### HasIncludeNoteIds

`func (o *OpportunityToAgreementConversion) HasIncludeNoteIds() bool`

HasIncludeNoteIds returns a boolean if a field has been set.

### GetIncludeDocumentIds

`func (o *OpportunityToAgreementConversion) GetIncludeDocumentIds() []int32`

GetIncludeDocumentIds returns the IncludeDocumentIds field if non-nil, zero value otherwise.

### GetIncludeDocumentIdsOk

`func (o *OpportunityToAgreementConversion) GetIncludeDocumentIdsOk() (*[]int32, bool)`

GetIncludeDocumentIdsOk returns a tuple with the IncludeDocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDocumentIds

`func (o *OpportunityToAgreementConversion) SetIncludeDocumentIds(v []int32)`

SetIncludeDocumentIds sets IncludeDocumentIds field to given value.

### HasIncludeDocumentIds

`func (o *OpportunityToAgreementConversion) HasIncludeDocumentIds() bool`

HasIncludeDocumentIds returns a boolean if a field has been set.

### GetIncludeProductIds

`func (o *OpportunityToAgreementConversion) GetIncludeProductIds() []int32`

GetIncludeProductIds returns the IncludeProductIds field if non-nil, zero value otherwise.

### GetIncludeProductIdsOk

`func (o *OpportunityToAgreementConversion) GetIncludeProductIdsOk() (*[]int32, bool)`

GetIncludeProductIdsOk returns a tuple with the IncludeProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductIds

`func (o *OpportunityToAgreementConversion) SetIncludeProductIds(v []int32)`

SetIncludeProductIds sets IncludeProductIds field to given value.

### HasIncludeProductIds

`func (o *OpportunityToAgreementConversion) HasIncludeProductIds() bool`

HasIncludeProductIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


