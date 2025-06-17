# UnitOfMeasure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**UomScheduleXref** | Pointer to **string** |  Max length: 31; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUnitOfMeasure

`func NewUnitOfMeasure(name string, ) *UnitOfMeasure`

NewUnitOfMeasure instantiates a new UnitOfMeasure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitOfMeasureWithDefaults

`func NewUnitOfMeasureWithDefaults() *UnitOfMeasure`

NewUnitOfMeasureWithDefaults instantiates a new UnitOfMeasure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnitOfMeasure) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnitOfMeasure) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnitOfMeasure) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UnitOfMeasure) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UnitOfMeasure) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnitOfMeasure) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnitOfMeasure) SetName(v string)`

SetName sets Name field to given value.


### GetInactiveFlag

`func (o *UnitOfMeasure) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *UnitOfMeasure) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *UnitOfMeasure) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *UnitOfMeasure) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *UnitOfMeasure) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *UnitOfMeasure) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetDefaultFlag

`func (o *UnitOfMeasure) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *UnitOfMeasure) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *UnitOfMeasure) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *UnitOfMeasure) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *UnitOfMeasure) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *UnitOfMeasure) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetUomScheduleXref

`func (o *UnitOfMeasure) GetUomScheduleXref() string`

GetUomScheduleXref returns the UomScheduleXref field if non-nil, zero value otherwise.

### GetUomScheduleXrefOk

`func (o *UnitOfMeasure) GetUomScheduleXrefOk() (*string, bool)`

GetUomScheduleXrefOk returns a tuple with the UomScheduleXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUomScheduleXref

`func (o *UnitOfMeasure) SetUomScheduleXref(v string)`

SetUomScheduleXref sets UomScheduleXref field to given value.

### HasUomScheduleXref

`func (o *UnitOfMeasure) HasUomScheduleXref() bool`

HasUomScheduleXref returns a boolean if a field has been set.

### GetInfo

`func (o *UnitOfMeasure) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UnitOfMeasure) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UnitOfMeasure) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *UnitOfMeasure) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


