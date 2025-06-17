# ServiceTemplateReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceTemplateReference

`func NewServiceTemplateReference() *ServiceTemplateReference`

NewServiceTemplateReference instantiates a new ServiceTemplateReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTemplateReferenceWithDefaults

`func NewServiceTemplateReferenceWithDefaults() *ServiceTemplateReference`

NewServiceTemplateReferenceWithDefaults instantiates a new ServiceTemplateReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceTemplateReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceTemplateReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceTemplateReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceTemplateReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ServiceTemplateReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ServiceTemplateReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ServiceTemplateReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceTemplateReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceTemplateReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceTemplateReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSummary

`func (o *ServiceTemplateReference) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ServiceTemplateReference) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ServiceTemplateReference) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ServiceTemplateReference) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetInfo

`func (o *ServiceTemplateReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceTemplateReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceTemplateReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceTemplateReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


