# CatalogItemReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCatalogItemReference

`func NewCatalogItemReference() *CatalogItemReference`

NewCatalogItemReference instantiates a new CatalogItemReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemReferenceWithDefaults

`func NewCatalogItemReferenceWithDefaults() *CatalogItemReference`

NewCatalogItemReferenceWithDefaults instantiates a new CatalogItemReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogItemReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogItemReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogItemReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogItemReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CatalogItemReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CatalogItemReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentifier

`func (o *CatalogItemReference) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CatalogItemReference) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CatalogItemReference) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CatalogItemReference) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *CatalogItemReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogItemReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogItemReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogItemReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInfo

`func (o *CatalogItemReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CatalogItemReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CatalogItemReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CatalogItemReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


