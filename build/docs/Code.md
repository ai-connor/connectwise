# Code

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Description** | **string** |  | 
**BoardId** | Pointer to **NullableInt32** |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCode

`func NewCode(name string, description string, ) *Code`

NewCode instantiates a new Code object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeWithDefaults

`func NewCodeWithDefaults() *Code`

NewCodeWithDefaults instantiates a new Code object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Code) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Code) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Code) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Code) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Code) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Code) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Code) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Code) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Code) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Code) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetBoardId

`func (o *Code) GetBoardId() int32`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *Code) GetBoardIdOk() (*int32, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *Code) SetBoardId(v int32)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *Code) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### SetBoardIdNil

`func (o *Code) SetBoardIdNil(b bool)`

 SetBoardIdNil sets the value for BoardId to be an explicit nil

### UnsetBoardId
`func (o *Code) UnsetBoardId()`

UnsetBoardId ensures that no value is present for BoardId, not even an explicit nil
### GetLocationId

`func (o *Code) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Code) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Code) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *Code) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *Code) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *Code) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetBusinessUnitId

`func (o *Code) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *Code) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *Code) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *Code) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *Code) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *Code) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetInfo

`func (o *Code) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Code) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Code) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Code) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


