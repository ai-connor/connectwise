# ProjectStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**NoTimeFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**StatusIndicator** | Pointer to [**StatusIndicatorReference**](StatusIndicatorReference.md) |  | [optional] 
**CustomStatusIndicatorName** | Pointer to **string** | Required when statusIndicator is Custom. Max length: 30; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectStatus

`func NewProjectStatus(name string, ) *ProjectStatus`

NewProjectStatus instantiates a new ProjectStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectStatusWithDefaults

`func NewProjectStatusWithDefaults() *ProjectStatus`

NewProjectStatusWithDefaults instantiates a new ProjectStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectStatus) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *ProjectStatus) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ProjectStatus) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ProjectStatus) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ProjectStatus) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ProjectStatus) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ProjectStatus) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *ProjectStatus) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ProjectStatus) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ProjectStatus) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ProjectStatus) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ProjectStatus) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ProjectStatus) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetNoTimeFlag

`func (o *ProjectStatus) GetNoTimeFlag() bool`

GetNoTimeFlag returns the NoTimeFlag field if non-nil, zero value otherwise.

### GetNoTimeFlagOk

`func (o *ProjectStatus) GetNoTimeFlagOk() (*bool, bool)`

GetNoTimeFlagOk returns a tuple with the NoTimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoTimeFlag

`func (o *ProjectStatus) SetNoTimeFlag(v bool)`

SetNoTimeFlag sets NoTimeFlag field to given value.

### HasNoTimeFlag

`func (o *ProjectStatus) HasNoTimeFlag() bool`

HasNoTimeFlag returns a boolean if a field has been set.

### SetNoTimeFlagNil

`func (o *ProjectStatus) SetNoTimeFlagNil(b bool)`

 SetNoTimeFlagNil sets the value for NoTimeFlag to be an explicit nil

### UnsetNoTimeFlag
`func (o *ProjectStatus) UnsetNoTimeFlag()`

UnsetNoTimeFlag ensures that no value is present for NoTimeFlag, not even an explicit nil
### GetClosedFlag

`func (o *ProjectStatus) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *ProjectStatus) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *ProjectStatus) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *ProjectStatus) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *ProjectStatus) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *ProjectStatus) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetStatusIndicator

`func (o *ProjectStatus) GetStatusIndicator() StatusIndicatorReference`

GetStatusIndicator returns the StatusIndicator field if non-nil, zero value otherwise.

### GetStatusIndicatorOk

`func (o *ProjectStatus) GetStatusIndicatorOk() (*StatusIndicatorReference, bool)`

GetStatusIndicatorOk returns a tuple with the StatusIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusIndicator

`func (o *ProjectStatus) SetStatusIndicator(v StatusIndicatorReference)`

SetStatusIndicator sets StatusIndicator field to given value.

### HasStatusIndicator

`func (o *ProjectStatus) HasStatusIndicator() bool`

HasStatusIndicator returns a boolean if a field has been set.

### GetCustomStatusIndicatorName

`func (o *ProjectStatus) GetCustomStatusIndicatorName() string`

GetCustomStatusIndicatorName returns the CustomStatusIndicatorName field if non-nil, zero value otherwise.

### GetCustomStatusIndicatorNameOk

`func (o *ProjectStatus) GetCustomStatusIndicatorNameOk() (*string, bool)`

GetCustomStatusIndicatorNameOk returns a tuple with the CustomStatusIndicatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStatusIndicatorName

`func (o *ProjectStatus) SetCustomStatusIndicatorName(v string)`

SetCustomStatusIndicatorName sets CustomStatusIndicatorName field to given value.

### HasCustomStatusIndicatorName

`func (o *ProjectStatus) HasCustomStatusIndicatorName() bool`

HasCustomStatusIndicatorName returns a boolean if a field has been set.

### GetInfo

`func (o *ProjectStatus) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectStatus) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectStatus) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectStatus) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


