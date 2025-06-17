# ContactContactTypeAssociationContactTypeAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | [**ContactTypeReference**](ContactTypeReference.md) |  | 
**Contact** | [**ContactReference**](ContactReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewContactContactTypeAssociationContactTypeAssociation

`func NewContactContactTypeAssociationContactTypeAssociation(type_ ContactTypeReference, contact ContactReference, ) *ContactContactTypeAssociationContactTypeAssociation`

NewContactContactTypeAssociationContactTypeAssociation instantiates a new ContactContactTypeAssociationContactTypeAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactContactTypeAssociationContactTypeAssociationWithDefaults

`func NewContactContactTypeAssociationContactTypeAssociationWithDefaults() *ContactContactTypeAssociationContactTypeAssociation`

NewContactContactTypeAssociationContactTypeAssociationWithDefaults instantiates a new ContactContactTypeAssociationContactTypeAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactContactTypeAssociationContactTypeAssociation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactContactTypeAssociationContactTypeAssociation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactContactTypeAssociationContactTypeAssociation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContactContactTypeAssociationContactTypeAssociation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ContactContactTypeAssociationContactTypeAssociation) GetType() ContactTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactContactTypeAssociationContactTypeAssociation) GetTypeOk() (*ContactTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactContactTypeAssociationContactTypeAssociation) SetType(v ContactTypeReference)`

SetType sets Type field to given value.


### GetContact

`func (o *ContactContactTypeAssociationContactTypeAssociation) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ContactContactTypeAssociationContactTypeAssociation) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ContactContactTypeAssociationContactTypeAssociation) SetContact(v ContactReference)`

SetContact sets Contact field to given value.


### GetInfo

`func (o *ContactContactTypeAssociationContactTypeAssociation) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ContactContactTypeAssociationContactTypeAssociation) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ContactContactTypeAssociationContactTypeAssociation) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ContactContactTypeAssociationContactTypeAssociation) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


