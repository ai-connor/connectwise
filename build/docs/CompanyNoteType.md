# CompanyNoteType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | Pointer to **string** |  Max length: 15; | [optional] 
**Name** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ImportFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyNoteType

`func NewCompanyNoteType(name string, ) *CompanyNoteType`

NewCompanyNoteType instantiates a new CompanyNoteType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyNoteTypeWithDefaults

`func NewCompanyNoteTypeWithDefaults() *CompanyNoteType`

NewCompanyNoteTypeWithDefaults instantiates a new CompanyNoteType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyNoteType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyNoteType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyNoteType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyNoteType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *CompanyNoteType) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CompanyNoteType) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CompanyNoteType) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CompanyNoteType) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *CompanyNoteType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyNoteType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyNoteType) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *CompanyNoteType) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *CompanyNoteType) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *CompanyNoteType) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *CompanyNoteType) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *CompanyNoteType) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *CompanyNoteType) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetImportFlag

`func (o *CompanyNoteType) GetImportFlag() bool`

GetImportFlag returns the ImportFlag field if non-nil, zero value otherwise.

### GetImportFlagOk

`func (o *CompanyNoteType) GetImportFlagOk() (*bool, bool)`

GetImportFlagOk returns a tuple with the ImportFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportFlag

`func (o *CompanyNoteType) SetImportFlag(v bool)`

SetImportFlag sets ImportFlag field to given value.

### HasImportFlag

`func (o *CompanyNoteType) HasImportFlag() bool`

HasImportFlag returns a boolean if a field has been set.

### SetImportFlagNil

`func (o *CompanyNoteType) SetImportFlagNil(b bool)`

 SetImportFlagNil sets the value for ImportFlag to be an explicit nil

### UnsetImportFlag
`func (o *CompanyNoteType) UnsetImportFlag()`

UnsetImportFlag ensures that no value is present for ImportFlag, not even an explicit nil
### GetInfo

`func (o *CompanyNoteType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyNoteType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyNoteType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyNoteType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


