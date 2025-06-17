# ProjectTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 200; | 
**Description** | Pointer to **string** |  | [optional] 
**ConnectWiseId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ProjectTypeReference**](ProjectTypeReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectTemplate

`func NewProjectTemplate(name string, ) *ProjectTemplate`

NewProjectTemplate instantiates a new ProjectTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTemplateWithDefaults

`func NewProjectTemplateWithDefaults() *ProjectTemplate`

NewProjectTemplateWithDefaults instantiates a new ProjectTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConnectWiseId

`func (o *ProjectTemplate) GetConnectWiseId() string`

GetConnectWiseId returns the ConnectWiseId field if non-nil, zero value otherwise.

### GetConnectWiseIdOk

`func (o *ProjectTemplate) GetConnectWiseIdOk() (*string, bool)`

GetConnectWiseIdOk returns a tuple with the ConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseId

`func (o *ProjectTemplate) SetConnectWiseId(v string)`

SetConnectWiseId sets ConnectWiseId field to given value.

### HasConnectWiseId

`func (o *ProjectTemplate) HasConnectWiseId() bool`

HasConnectWiseId returns a boolean if a field has been set.

### GetType

`func (o *ProjectTemplate) GetType() ProjectTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectTemplate) GetTypeOk() (*ProjectTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectTemplate) SetType(v ProjectTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *ProjectTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInfo

`func (o *ProjectTemplate) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectTemplate) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectTemplate) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectTemplate) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


