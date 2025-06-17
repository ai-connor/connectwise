# OpportunityType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Description** | **string** |  Max length: 50; | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOpportunityType

`func NewOpportunityType(description string, ) *OpportunityType`

NewOpportunityType instantiates a new OpportunityType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityTypeWithDefaults

`func NewOpportunityTypeWithDefaults() *OpportunityType`

NewOpportunityTypeWithDefaults instantiates a new OpportunityType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpportunityType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpportunityType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpportunityType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpportunityType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *OpportunityType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpportunityType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpportunityType) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInactiveFlag

`func (o *OpportunityType) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *OpportunityType) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *OpportunityType) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *OpportunityType) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *OpportunityType) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *OpportunityType) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *OpportunityType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OpportunityType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OpportunityType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OpportunityType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


