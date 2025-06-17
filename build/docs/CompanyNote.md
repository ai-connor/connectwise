# CompanyNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Text** | **string** |  | 
**Type** | Pointer to [**NoteTypeReference**](NoteTypeReference.md) |  | [optional] 
**Flagged** | Pointer to **NullableBool** |  | [optional] 
**EnteredBy** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyNote

`func NewCompanyNote(text string, ) *CompanyNote`

NewCompanyNote instantiates a new CompanyNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyNoteWithDefaults

`func NewCompanyNoteWithDefaults() *CompanyNote`

NewCompanyNoteWithDefaults instantiates a new CompanyNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyNote) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetText

`func (o *CompanyNote) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CompanyNote) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CompanyNote) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *CompanyNote) GetType() NoteTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyNote) GetTypeOk() (*NoteTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyNote) SetType(v NoteTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *CompanyNote) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlagged

`func (o *CompanyNote) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *CompanyNote) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *CompanyNote) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *CompanyNote) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### SetFlaggedNil

`func (o *CompanyNote) SetFlaggedNil(b bool)`

 SetFlaggedNil sets the value for Flagged to be an explicit nil

### UnsetFlagged
`func (o *CompanyNote) UnsetFlagged()`

UnsetFlagged ensures that no value is present for Flagged, not even an explicit nil
### GetEnteredBy

`func (o *CompanyNote) GetEnteredBy() string`

GetEnteredBy returns the EnteredBy field if non-nil, zero value otherwise.

### GetEnteredByOk

`func (o *CompanyNote) GetEnteredByOk() (*string, bool)`

GetEnteredByOk returns a tuple with the EnteredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredBy

`func (o *CompanyNote) SetEnteredBy(v string)`

SetEnteredBy sets EnteredBy field to given value.

### HasEnteredBy

`func (o *CompanyNote) HasEnteredBy() bool`

HasEnteredBy returns a boolean if a field has been set.

### GetCompany

`func (o *CompanyNote) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyNote) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyNote) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CompanyNote) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetInfo

`func (o *CompanyNote) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyNote) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyNote) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyNote) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


