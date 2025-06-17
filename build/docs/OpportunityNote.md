# OpportunityNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OpportunityId** | Pointer to **NullableInt32** |  | [optional] 
**Type** | Pointer to [**NoteTypeReference**](NoteTypeReference.md) |  | [optional] 
**Text** | **string** |  | 
**Flagged** | Pointer to **NullableBool** |  | [optional] 
**EnteredBy** | Pointer to **string** |  | [optional] 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOpportunityNote

`func NewOpportunityNote(text string, ) *OpportunityNote`

NewOpportunityNote instantiates a new OpportunityNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityNoteWithDefaults

`func NewOpportunityNoteWithDefaults() *OpportunityNote`

NewOpportunityNoteWithDefaults instantiates a new OpportunityNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpportunityNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpportunityNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpportunityNote) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpportunityNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOpportunityId

`func (o *OpportunityNote) GetOpportunityId() int32`

GetOpportunityId returns the OpportunityId field if non-nil, zero value otherwise.

### GetOpportunityIdOk

`func (o *OpportunityNote) GetOpportunityIdOk() (*int32, bool)`

GetOpportunityIdOk returns a tuple with the OpportunityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityId

`func (o *OpportunityNote) SetOpportunityId(v int32)`

SetOpportunityId sets OpportunityId field to given value.

### HasOpportunityId

`func (o *OpportunityNote) HasOpportunityId() bool`

HasOpportunityId returns a boolean if a field has been set.

### SetOpportunityIdNil

`func (o *OpportunityNote) SetOpportunityIdNil(b bool)`

 SetOpportunityIdNil sets the value for OpportunityId to be an explicit nil

### UnsetOpportunityId
`func (o *OpportunityNote) UnsetOpportunityId()`

UnsetOpportunityId ensures that no value is present for OpportunityId, not even an explicit nil
### GetType

`func (o *OpportunityNote) GetType() NoteTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpportunityNote) GetTypeOk() (*NoteTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpportunityNote) SetType(v NoteTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *OpportunityNote) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *OpportunityNote) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OpportunityNote) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OpportunityNote) SetText(v string)`

SetText sets Text field to given value.


### GetFlagged

`func (o *OpportunityNote) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *OpportunityNote) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *OpportunityNote) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *OpportunityNote) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### SetFlaggedNil

`func (o *OpportunityNote) SetFlaggedNil(b bool)`

 SetFlaggedNil sets the value for Flagged to be an explicit nil

### UnsetFlagged
`func (o *OpportunityNote) UnsetFlagged()`

UnsetFlagged ensures that no value is present for Flagged, not even an explicit nil
### GetEnteredBy

`func (o *OpportunityNote) GetEnteredBy() string`

GetEnteredBy returns the EnteredBy field if non-nil, zero value otherwise.

### GetEnteredByOk

`func (o *OpportunityNote) GetEnteredByOk() (*string, bool)`

GetEnteredByOk returns a tuple with the EnteredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredBy

`func (o *OpportunityNote) SetEnteredBy(v string)`

SetEnteredBy sets EnteredBy field to given value.

### HasEnteredBy

`func (o *OpportunityNote) HasEnteredBy() bool`

HasEnteredBy returns a boolean if a field has been set.

### GetMobileGuid

`func (o *OpportunityNote) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *OpportunityNote) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *OpportunityNote) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *OpportunityNote) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *OpportunityNote) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *OpportunityNote) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetInfo

`func (o *OpportunityNote) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OpportunityNote) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OpportunityNote) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OpportunityNote) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


