# KPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Category** | Pointer to [**KPICategoryReference**](KPICategoryReference.md) |  | [optional] 
**DateFilter** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewKPI

`func NewKPI() *KPI`

NewKPI instantiates a new KPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKPIWithDefaults

`func NewKPIWithDefaults() *KPI`

NewKPIWithDefaults instantiates a new KPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KPI) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KPI) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KPI) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KPI) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KPI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KPI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KPI) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KPI) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *KPI) GetCategory() KPICategoryReference`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *KPI) GetCategoryOk() (*KPICategoryReference, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *KPI) SetCategory(v KPICategoryReference)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *KPI) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDateFilter

`func (o *KPI) GetDateFilter() string`

GetDateFilter returns the DateFilter field if non-nil, zero value otherwise.

### GetDateFilterOk

`func (o *KPI) GetDateFilterOk() (*string, bool)`

GetDateFilterOk returns a tuple with the DateFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFilter

`func (o *KPI) SetDateFilter(v string)`

SetDateFilter sets DateFilter field to given value.

### HasDateFilter

`func (o *KPI) HasDateFilter() bool`

HasDateFilter returns a boolean if a field has been set.

### GetSortOrder

`func (o *KPI) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *KPI) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *KPI) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *KPI) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *KPI) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *KPI) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetInactiveFlag

`func (o *KPI) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *KPI) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *KPI) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *KPI) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *KPI) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *KPI) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


