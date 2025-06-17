# CompanyContactTypeAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | [**ContactTypeReference**](ContactTypeReference.md) |  | 
**Contact** | [**ContactReference**](ContactReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyContactTypeAssociation

`func NewCompanyContactTypeAssociation(type_ ContactTypeReference, contact ContactReference, ) *CompanyContactTypeAssociation`

NewCompanyContactTypeAssociation instantiates a new CompanyContactTypeAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyContactTypeAssociationWithDefaults

`func NewCompanyContactTypeAssociationWithDefaults() *CompanyContactTypeAssociation`

NewCompanyContactTypeAssociationWithDefaults instantiates a new CompanyContactTypeAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyContactTypeAssociation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyContactTypeAssociation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyContactTypeAssociation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyContactTypeAssociation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CompanyContactTypeAssociation) GetType() ContactTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyContactTypeAssociation) GetTypeOk() (*ContactTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyContactTypeAssociation) SetType(v ContactTypeReference)`

SetType sets Type field to given value.


### GetContact

`func (o *CompanyContactTypeAssociation) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *CompanyContactTypeAssociation) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *CompanyContactTypeAssociation) SetContact(v ContactReference)`

SetContact sets Contact field to given value.


### GetInfo

`func (o *CompanyContactTypeAssociation) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyContactTypeAssociation) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyContactTypeAssociation) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyContactTypeAssociation) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


