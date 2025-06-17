# BoardItemAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**SubTypeAssociationIds** | Pointer to **[]int32** | If addAllSubTypesFlag and removeAllSubTypesFlag are both false, this field is required. | [optional] 
**AddAllSubTypesFlag** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllSubTypesFlag** | Pointer to **NullableBool** |  | [optional] 
**Item** | Pointer to [**ServiceItemReference**](ServiceItemReference.md) |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardItemAssociation

`func NewBoardItemAssociation(id int32, ) *BoardItemAssociation`

NewBoardItemAssociation instantiates a new BoardItemAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardItemAssociationWithDefaults

`func NewBoardItemAssociationWithDefaults() *BoardItemAssociation`

NewBoardItemAssociationWithDefaults instantiates a new BoardItemAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardItemAssociation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardItemAssociation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardItemAssociation) SetId(v int32)`

SetId sets Id field to given value.


### GetSubTypeAssociationIds

`func (o *BoardItemAssociation) GetSubTypeAssociationIds() []int32`

GetSubTypeAssociationIds returns the SubTypeAssociationIds field if non-nil, zero value otherwise.

### GetSubTypeAssociationIdsOk

`func (o *BoardItemAssociation) GetSubTypeAssociationIdsOk() (*[]int32, bool)`

GetSubTypeAssociationIdsOk returns a tuple with the SubTypeAssociationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTypeAssociationIds

`func (o *BoardItemAssociation) SetSubTypeAssociationIds(v []int32)`

SetSubTypeAssociationIds sets SubTypeAssociationIds field to given value.

### HasSubTypeAssociationIds

`func (o *BoardItemAssociation) HasSubTypeAssociationIds() bool`

HasSubTypeAssociationIds returns a boolean if a field has been set.

### GetAddAllSubTypesFlag

`func (o *BoardItemAssociation) GetAddAllSubTypesFlag() bool`

GetAddAllSubTypesFlag returns the AddAllSubTypesFlag field if non-nil, zero value otherwise.

### GetAddAllSubTypesFlagOk

`func (o *BoardItemAssociation) GetAddAllSubTypesFlagOk() (*bool, bool)`

GetAddAllSubTypesFlagOk returns a tuple with the AddAllSubTypesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllSubTypesFlag

`func (o *BoardItemAssociation) SetAddAllSubTypesFlag(v bool)`

SetAddAllSubTypesFlag sets AddAllSubTypesFlag field to given value.

### HasAddAllSubTypesFlag

`func (o *BoardItemAssociation) HasAddAllSubTypesFlag() bool`

HasAddAllSubTypesFlag returns a boolean if a field has been set.

### SetAddAllSubTypesFlagNil

`func (o *BoardItemAssociation) SetAddAllSubTypesFlagNil(b bool)`

 SetAddAllSubTypesFlagNil sets the value for AddAllSubTypesFlag to be an explicit nil

### UnsetAddAllSubTypesFlag
`func (o *BoardItemAssociation) UnsetAddAllSubTypesFlag()`

UnsetAddAllSubTypesFlag ensures that no value is present for AddAllSubTypesFlag, not even an explicit nil
### GetRemoveAllSubTypesFlag

`func (o *BoardItemAssociation) GetRemoveAllSubTypesFlag() bool`

GetRemoveAllSubTypesFlag returns the RemoveAllSubTypesFlag field if non-nil, zero value otherwise.

### GetRemoveAllSubTypesFlagOk

`func (o *BoardItemAssociation) GetRemoveAllSubTypesFlagOk() (*bool, bool)`

GetRemoveAllSubTypesFlagOk returns a tuple with the RemoveAllSubTypesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllSubTypesFlag

`func (o *BoardItemAssociation) SetRemoveAllSubTypesFlag(v bool)`

SetRemoveAllSubTypesFlag sets RemoveAllSubTypesFlag field to given value.

### HasRemoveAllSubTypesFlag

`func (o *BoardItemAssociation) HasRemoveAllSubTypesFlag() bool`

HasRemoveAllSubTypesFlag returns a boolean if a field has been set.

### SetRemoveAllSubTypesFlagNil

`func (o *BoardItemAssociation) SetRemoveAllSubTypesFlagNil(b bool)`

 SetRemoveAllSubTypesFlagNil sets the value for RemoveAllSubTypesFlag to be an explicit nil

### UnsetRemoveAllSubTypesFlag
`func (o *BoardItemAssociation) UnsetRemoveAllSubTypesFlag()`

UnsetRemoveAllSubTypesFlag ensures that no value is present for RemoveAllSubTypesFlag, not even an explicit nil
### GetItem

`func (o *BoardItemAssociation) GetItem() ServiceItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *BoardItemAssociation) GetItemOk() (*ServiceItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *BoardItemAssociation) SetItem(v ServiceItemReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *BoardItemAssociation) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetBoard

`func (o *BoardItemAssociation) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *BoardItemAssociation) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *BoardItemAssociation) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *BoardItemAssociation) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetInfo

`func (o *BoardItemAssociation) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardItemAssociation) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardItemAssociation) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardItemAssociation) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


