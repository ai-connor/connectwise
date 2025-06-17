# ServiceStatusReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Sort** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceStatusReference

`func NewServiceStatusReference() *ServiceStatusReference`

NewServiceStatusReference instantiates a new ServiceStatusReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStatusReferenceWithDefaults

`func NewServiceStatusReferenceWithDefaults() *ServiceStatusReference`

NewServiceStatusReferenceWithDefaults instantiates a new ServiceStatusReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceStatusReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceStatusReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceStatusReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceStatusReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ServiceStatusReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ServiceStatusReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ServiceStatusReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceStatusReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceStatusReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceStatusReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSort

`func (o *ServiceStatusReference) GetSort() int32`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ServiceStatusReference) GetSortOk() (*int32, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ServiceStatusReference) SetSort(v int32)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ServiceStatusReference) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *ServiceStatusReference) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *ServiceStatusReference) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetInfo

`func (o *ServiceStatusReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceStatusReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceStatusReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceStatusReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


