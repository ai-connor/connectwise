# CompanyCompanyTypeAssociationCompanyTypeAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | [**CompanyTypeReference**](CompanyTypeReference.md) |  | 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyCompanyTypeAssociationCompanyTypeAssociation

`func NewCompanyCompanyTypeAssociationCompanyTypeAssociation(type_ CompanyTypeReference, company CompanyReference, ) *CompanyCompanyTypeAssociationCompanyTypeAssociation`

NewCompanyCompanyTypeAssociationCompanyTypeAssociation instantiates a new CompanyCompanyTypeAssociationCompanyTypeAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyCompanyTypeAssociationCompanyTypeAssociationWithDefaults

`func NewCompanyCompanyTypeAssociationCompanyTypeAssociationWithDefaults() *CompanyCompanyTypeAssociationCompanyTypeAssociation`

NewCompanyCompanyTypeAssociationCompanyTypeAssociationWithDefaults instantiates a new CompanyCompanyTypeAssociationCompanyTypeAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) GetType() CompanyTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) GetTypeOk() (*CompanyTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) SetType(v CompanyTypeReference)`

SetType sets Type field to given value.


### GetCompany

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetInfo

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyCompanyTypeAssociationCompanyTypeAssociation) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


