# ProjectBoardKanbanSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Color** | Pointer to **string** |  Max length: 4; | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Statuses** | Pointer to [**[]ProjectBoardKanbanStatus**](ProjectBoardKanbanStatus.md) |  | [optional] 
**UpdatedBy** | Pointer to **string** |  Max length: 15; | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectBoardKanbanSetting

`func NewProjectBoardKanbanSetting(name string, ) *ProjectBoardKanbanSetting`

NewProjectBoardKanbanSetting instantiates a new ProjectBoardKanbanSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectBoardKanbanSettingWithDefaults

`func NewProjectBoardKanbanSettingWithDefaults() *ProjectBoardKanbanSetting`

NewProjectBoardKanbanSettingWithDefaults instantiates a new ProjectBoardKanbanSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectBoardKanbanSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectBoardKanbanSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectBoardKanbanSetting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectBoardKanbanSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectBoardKanbanSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectBoardKanbanSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectBoardKanbanSetting) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *ProjectBoardKanbanSetting) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ProjectBoardKanbanSetting) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ProjectBoardKanbanSetting) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ProjectBoardKanbanSetting) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetOrder

`func (o *ProjectBoardKanbanSetting) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ProjectBoardKanbanSetting) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ProjectBoardKanbanSetting) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ProjectBoardKanbanSetting) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetStatuses

`func (o *ProjectBoardKanbanSetting) GetStatuses() []ProjectBoardKanbanStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ProjectBoardKanbanSetting) GetStatusesOk() (*[]ProjectBoardKanbanStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ProjectBoardKanbanSetting) SetStatuses(v []ProjectBoardKanbanStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *ProjectBoardKanbanSetting) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ProjectBoardKanbanSetting) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ProjectBoardKanbanSetting) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ProjectBoardKanbanSetting) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ProjectBoardKanbanSetting) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ProjectBoardKanbanSetting) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ProjectBoardKanbanSetting) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ProjectBoardKanbanSetting) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ProjectBoardKanbanSetting) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


