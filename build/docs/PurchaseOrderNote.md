# PurchaseOrderNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PurchaseHeaderRecID** | Pointer to **NullableInt32** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**NoteTypeReference**](NoteTypeReference.md) |  | [optional] 
**Flagged** | Pointer to **NullableBool** |  | [optional] 
**EnteredBy** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPurchaseOrderNote

`func NewPurchaseOrderNote() *PurchaseOrderNote`

NewPurchaseOrderNote instantiates a new PurchaseOrderNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderNoteWithDefaults

`func NewPurchaseOrderNoteWithDefaults() *PurchaseOrderNote`

NewPurchaseOrderNoteWithDefaults instantiates a new PurchaseOrderNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PurchaseOrderNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseOrderNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseOrderNote) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseOrderNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPurchaseHeaderRecID

`func (o *PurchaseOrderNote) GetPurchaseHeaderRecID() int32`

GetPurchaseHeaderRecID returns the PurchaseHeaderRecID field if non-nil, zero value otherwise.

### GetPurchaseHeaderRecIDOk

`func (o *PurchaseOrderNote) GetPurchaseHeaderRecIDOk() (*int32, bool)`

GetPurchaseHeaderRecIDOk returns a tuple with the PurchaseHeaderRecID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseHeaderRecID

`func (o *PurchaseOrderNote) SetPurchaseHeaderRecID(v int32)`

SetPurchaseHeaderRecID sets PurchaseHeaderRecID field to given value.

### HasPurchaseHeaderRecID

`func (o *PurchaseOrderNote) HasPurchaseHeaderRecID() bool`

HasPurchaseHeaderRecID returns a boolean if a field has been set.

### SetPurchaseHeaderRecIDNil

`func (o *PurchaseOrderNote) SetPurchaseHeaderRecIDNil(b bool)`

 SetPurchaseHeaderRecIDNil sets the value for PurchaseHeaderRecID to be an explicit nil

### UnsetPurchaseHeaderRecID
`func (o *PurchaseOrderNote) UnsetPurchaseHeaderRecID()`

UnsetPurchaseHeaderRecID ensures that no value is present for PurchaseHeaderRecID, not even an explicit nil
### GetText

`func (o *PurchaseOrderNote) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PurchaseOrderNote) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PurchaseOrderNote) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PurchaseOrderNote) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *PurchaseOrderNote) GetType() NoteTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PurchaseOrderNote) GetTypeOk() (*NoteTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PurchaseOrderNote) SetType(v NoteTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *PurchaseOrderNote) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlagged

`func (o *PurchaseOrderNote) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *PurchaseOrderNote) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *PurchaseOrderNote) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *PurchaseOrderNote) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### SetFlaggedNil

`func (o *PurchaseOrderNote) SetFlaggedNil(b bool)`

 SetFlaggedNil sets the value for Flagged to be an explicit nil

### UnsetFlagged
`func (o *PurchaseOrderNote) UnsetFlagged()`

UnsetFlagged ensures that no value is present for Flagged, not even an explicit nil
### GetEnteredBy

`func (o *PurchaseOrderNote) GetEnteredBy() string`

GetEnteredBy returns the EnteredBy field if non-nil, zero value otherwise.

### GetEnteredByOk

`func (o *PurchaseOrderNote) GetEnteredByOk() (*string, bool)`

GetEnteredByOk returns a tuple with the EnteredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredBy

`func (o *PurchaseOrderNote) SetEnteredBy(v string)`

SetEnteredBy sets EnteredBy field to given value.

### HasEnteredBy

`func (o *PurchaseOrderNote) HasEnteredBy() bool`

HasEnteredBy returns a boolean if a field has been set.

### GetInfo

`func (o *PurchaseOrderNote) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PurchaseOrderNote) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PurchaseOrderNote) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PurchaseOrderNote) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


