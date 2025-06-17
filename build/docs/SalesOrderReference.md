# SalesOrderReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSalesOrderReference

`func NewSalesOrderReference() *SalesOrderReference`

NewSalesOrderReference instantiates a new SalesOrderReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesOrderReferenceWithDefaults

`func NewSalesOrderReferenceWithDefaults() *SalesOrderReference`

NewSalesOrderReferenceWithDefaults instantiates a new SalesOrderReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SalesOrderReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalesOrderReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalesOrderReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SalesOrderReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SalesOrderReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SalesOrderReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentifier

`func (o *SalesOrderReference) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *SalesOrderReference) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *SalesOrderReference) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *SalesOrderReference) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetInfo

`func (o *SalesOrderReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SalesOrderReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SalesOrderReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SalesOrderReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


