# PhaseStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**CollapsedFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**BoardAssociationIds** | Pointer to **[]int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPhaseStatusInfo

`func NewPhaseStatusInfo() *PhaseStatusInfo`

NewPhaseStatusInfo instantiates a new PhaseStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhaseStatusInfoWithDefaults

`func NewPhaseStatusInfoWithDefaults() *PhaseStatusInfo`

NewPhaseStatusInfoWithDefaults instantiates a new PhaseStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhaseStatusInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhaseStatusInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhaseStatusInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PhaseStatusInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PhaseStatusInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhaseStatusInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhaseStatusInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PhaseStatusInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *PhaseStatusInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *PhaseStatusInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *PhaseStatusInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *PhaseStatusInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *PhaseStatusInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *PhaseStatusInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *PhaseStatusInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *PhaseStatusInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *PhaseStatusInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *PhaseStatusInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *PhaseStatusInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *PhaseStatusInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetCollapsedFlag

`func (o *PhaseStatusInfo) GetCollapsedFlag() bool`

GetCollapsedFlag returns the CollapsedFlag field if non-nil, zero value otherwise.

### GetCollapsedFlagOk

`func (o *PhaseStatusInfo) GetCollapsedFlagOk() (*bool, bool)`

GetCollapsedFlagOk returns a tuple with the CollapsedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapsedFlag

`func (o *PhaseStatusInfo) SetCollapsedFlag(v bool)`

SetCollapsedFlag sets CollapsedFlag field to given value.

### HasCollapsedFlag

`func (o *PhaseStatusInfo) HasCollapsedFlag() bool`

HasCollapsedFlag returns a boolean if a field has been set.

### SetCollapsedFlagNil

`func (o *PhaseStatusInfo) SetCollapsedFlagNil(b bool)`

 SetCollapsedFlagNil sets the value for CollapsedFlag to be an explicit nil

### UnsetCollapsedFlag
`func (o *PhaseStatusInfo) UnsetCollapsedFlag()`

UnsetCollapsedFlag ensures that no value is present for CollapsedFlag, not even an explicit nil
### GetClosedFlag

`func (o *PhaseStatusInfo) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *PhaseStatusInfo) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *PhaseStatusInfo) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *PhaseStatusInfo) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *PhaseStatusInfo) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *PhaseStatusInfo) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetBoardAssociationIds

`func (o *PhaseStatusInfo) GetBoardAssociationIds() []int32`

GetBoardAssociationIds returns the BoardAssociationIds field if non-nil, zero value otherwise.

### GetBoardAssociationIdsOk

`func (o *PhaseStatusInfo) GetBoardAssociationIdsOk() (*[]int32, bool)`

GetBoardAssociationIdsOk returns a tuple with the BoardAssociationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardAssociationIds

`func (o *PhaseStatusInfo) SetBoardAssociationIds(v []int32)`

SetBoardAssociationIds sets BoardAssociationIds field to given value.

### HasBoardAssociationIds

`func (o *PhaseStatusInfo) HasBoardAssociationIds() bool`

HasBoardAssociationIds returns a boolean if a field has been set.

### GetInfo

`func (o *PhaseStatusInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PhaseStatusInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PhaseStatusInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PhaseStatusInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


