# TaxCodeLevelReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTaxCodeLevelReference

`func NewTaxCodeLevelReference() *TaxCodeLevelReference`

NewTaxCodeLevelReference instantiates a new TaxCodeLevelReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCodeLevelReferenceWithDefaults

`func NewTaxCodeLevelReferenceWithDefaults() *TaxCodeLevelReference`

NewTaxCodeLevelReferenceWithDefaults instantiates a new TaxCodeLevelReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxCodeLevelReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxCodeLevelReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxCodeLevelReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaxCodeLevelReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaxCodeLevelReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaxCodeLevelReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TaxCodeLevelReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxCodeLevelReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxCodeLevelReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxCodeLevelReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInfo

`func (o *TaxCodeLevelReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TaxCodeLevelReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TaxCodeLevelReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TaxCodeLevelReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


