# ContactNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ContactId** | Pointer to **NullableInt32** |  | [optional] 
**Text** | **string** |  | 
**Type** | Pointer to [**NoteTypeReference**](NoteTypeReference.md) |  | [optional] 
**Flagged** | Pointer to **NullableBool** |  | [optional] 
**EnteredBy** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewContactNote

`func NewContactNote(text string, ) *ContactNote`

NewContactNote instantiates a new ContactNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactNoteWithDefaults

`func NewContactNoteWithDefaults() *ContactNote`

NewContactNoteWithDefaults instantiates a new ContactNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactNote) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContactNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContactId

`func (o *ContactNote) GetContactId() int32`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ContactNote) GetContactIdOk() (*int32, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ContactNote) SetContactId(v int32)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *ContactNote) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *ContactNote) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *ContactNote) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetText

`func (o *ContactNote) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ContactNote) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ContactNote) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *ContactNote) GetType() NoteTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactNote) GetTypeOk() (*NoteTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactNote) SetType(v NoteTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *ContactNote) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlagged

`func (o *ContactNote) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *ContactNote) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *ContactNote) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *ContactNote) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### SetFlaggedNil

`func (o *ContactNote) SetFlaggedNil(b bool)`

 SetFlaggedNil sets the value for Flagged to be an explicit nil

### UnsetFlagged
`func (o *ContactNote) UnsetFlagged()`

UnsetFlagged ensures that no value is present for Flagged, not even an explicit nil
### GetEnteredBy

`func (o *ContactNote) GetEnteredBy() string`

GetEnteredBy returns the EnteredBy field if non-nil, zero value otherwise.

### GetEnteredByOk

`func (o *ContactNote) GetEnteredByOk() (*string, bool)`

GetEnteredByOk returns a tuple with the EnteredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredBy

`func (o *ContactNote) SetEnteredBy(v string)`

SetEnteredBy sets EnteredBy field to given value.

### HasEnteredBy

`func (o *ContactNote) HasEnteredBy() bool`

HasEnteredBy returns a boolean if a field has been set.

### GetInfo

`func (o *ContactNote) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ContactNote) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ContactNote) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ContactNote) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


