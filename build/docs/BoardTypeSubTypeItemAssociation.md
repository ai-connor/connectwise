# BoardTypeSubTypeItemAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**ServiceTypeReference**](ServiceTypeReference.md) |  | [optional] 
**SubType** | Pointer to [**ServiceSubTypeReference**](ServiceSubTypeReference.md) |  | [optional] 
**Item** | Pointer to [**ServiceItemReference**](ServiceItemReference.md) |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardTypeSubTypeItemAssociation

`func NewBoardTypeSubTypeItemAssociation() *BoardTypeSubTypeItemAssociation`

NewBoardTypeSubTypeItemAssociation instantiates a new BoardTypeSubTypeItemAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardTypeSubTypeItemAssociationWithDefaults

`func NewBoardTypeSubTypeItemAssociationWithDefaults() *BoardTypeSubTypeItemAssociation`

NewBoardTypeSubTypeItemAssociationWithDefaults instantiates a new BoardTypeSubTypeItemAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardTypeSubTypeItemAssociation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardTypeSubTypeItemAssociation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardTypeSubTypeItemAssociation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardTypeSubTypeItemAssociation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BoardTypeSubTypeItemAssociation) GetType() ServiceTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BoardTypeSubTypeItemAssociation) GetTypeOk() (*ServiceTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BoardTypeSubTypeItemAssociation) SetType(v ServiceTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *BoardTypeSubTypeItemAssociation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubType

`func (o *BoardTypeSubTypeItemAssociation) GetSubType() ServiceSubTypeReference`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *BoardTypeSubTypeItemAssociation) GetSubTypeOk() (*ServiceSubTypeReference, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *BoardTypeSubTypeItemAssociation) SetSubType(v ServiceSubTypeReference)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *BoardTypeSubTypeItemAssociation) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetItem

`func (o *BoardTypeSubTypeItemAssociation) GetItem() ServiceItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *BoardTypeSubTypeItemAssociation) GetItemOk() (*ServiceItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *BoardTypeSubTypeItemAssociation) SetItem(v ServiceItemReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *BoardTypeSubTypeItemAssociation) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetBoard

`func (o *BoardTypeSubTypeItemAssociation) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *BoardTypeSubTypeItemAssociation) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *BoardTypeSubTypeItemAssociation) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *BoardTypeSubTypeItemAssociation) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetInfo

`func (o *BoardTypeSubTypeItemAssociation) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardTypeSubTypeItemAssociation) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardTypeSubTypeItemAssociation) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardTypeSubTypeItemAssociation) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


