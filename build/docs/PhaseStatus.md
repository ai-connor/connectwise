# PhaseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**CollapsedFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**BoardAssociationIds** | Pointer to **[]int32** |  | [optional] 
**StatusIndicator** | Pointer to [**StatusIndicatorReference**](StatusIndicatorReference.md) |  | [optional] 
**CustomStatusIndicatorName** | Pointer to **string** | Required when statusIndicator is Custom. Max length: 30; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPhaseStatus

`func NewPhaseStatus(name string, ) *PhaseStatus`

NewPhaseStatus instantiates a new PhaseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhaseStatusWithDefaults

`func NewPhaseStatusWithDefaults() *PhaseStatus`

NewPhaseStatusWithDefaults instantiates a new PhaseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhaseStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhaseStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhaseStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PhaseStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PhaseStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhaseStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhaseStatus) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *PhaseStatus) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *PhaseStatus) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *PhaseStatus) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *PhaseStatus) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *PhaseStatus) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *PhaseStatus) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *PhaseStatus) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *PhaseStatus) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *PhaseStatus) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *PhaseStatus) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *PhaseStatus) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *PhaseStatus) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetCollapsedFlag

`func (o *PhaseStatus) GetCollapsedFlag() bool`

GetCollapsedFlag returns the CollapsedFlag field if non-nil, zero value otherwise.

### GetCollapsedFlagOk

`func (o *PhaseStatus) GetCollapsedFlagOk() (*bool, bool)`

GetCollapsedFlagOk returns a tuple with the CollapsedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapsedFlag

`func (o *PhaseStatus) SetCollapsedFlag(v bool)`

SetCollapsedFlag sets CollapsedFlag field to given value.

### HasCollapsedFlag

`func (o *PhaseStatus) HasCollapsedFlag() bool`

HasCollapsedFlag returns a boolean if a field has been set.

### SetCollapsedFlagNil

`func (o *PhaseStatus) SetCollapsedFlagNil(b bool)`

 SetCollapsedFlagNil sets the value for CollapsedFlag to be an explicit nil

### UnsetCollapsedFlag
`func (o *PhaseStatus) UnsetCollapsedFlag()`

UnsetCollapsedFlag ensures that no value is present for CollapsedFlag, not even an explicit nil
### GetClosedFlag

`func (o *PhaseStatus) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *PhaseStatus) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *PhaseStatus) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *PhaseStatus) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *PhaseStatus) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *PhaseStatus) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetBoardAssociationIds

`func (o *PhaseStatus) GetBoardAssociationIds() []int32`

GetBoardAssociationIds returns the BoardAssociationIds field if non-nil, zero value otherwise.

### GetBoardAssociationIdsOk

`func (o *PhaseStatus) GetBoardAssociationIdsOk() (*[]int32, bool)`

GetBoardAssociationIdsOk returns a tuple with the BoardAssociationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardAssociationIds

`func (o *PhaseStatus) SetBoardAssociationIds(v []int32)`

SetBoardAssociationIds sets BoardAssociationIds field to given value.

### HasBoardAssociationIds

`func (o *PhaseStatus) HasBoardAssociationIds() bool`

HasBoardAssociationIds returns a boolean if a field has been set.

### GetStatusIndicator

`func (o *PhaseStatus) GetStatusIndicator() StatusIndicatorReference`

GetStatusIndicator returns the StatusIndicator field if non-nil, zero value otherwise.

### GetStatusIndicatorOk

`func (o *PhaseStatus) GetStatusIndicatorOk() (*StatusIndicatorReference, bool)`

GetStatusIndicatorOk returns a tuple with the StatusIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusIndicator

`func (o *PhaseStatus) SetStatusIndicator(v StatusIndicatorReference)`

SetStatusIndicator sets StatusIndicator field to given value.

### HasStatusIndicator

`func (o *PhaseStatus) HasStatusIndicator() bool`

HasStatusIndicator returns a boolean if a field has been set.

### GetCustomStatusIndicatorName

`func (o *PhaseStatus) GetCustomStatusIndicatorName() string`

GetCustomStatusIndicatorName returns the CustomStatusIndicatorName field if non-nil, zero value otherwise.

### GetCustomStatusIndicatorNameOk

`func (o *PhaseStatus) GetCustomStatusIndicatorNameOk() (*string, bool)`

GetCustomStatusIndicatorNameOk returns a tuple with the CustomStatusIndicatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStatusIndicatorName

`func (o *PhaseStatus) SetCustomStatusIndicatorName(v string)`

SetCustomStatusIndicatorName sets CustomStatusIndicatorName field to given value.

### HasCustomStatusIndicatorName

`func (o *PhaseStatus) HasCustomStatusIndicatorName() bool`

HasCustomStatusIndicatorName returns a boolean if a field has been set.

### GetInfo

`func (o *PhaseStatus) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PhaseStatus) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PhaseStatus) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PhaseStatus) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


