# RmaActionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewRmaActionInfo

`func NewRmaActionInfo() *RmaActionInfo`

NewRmaActionInfo instantiates a new RmaActionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRmaActionInfoWithDefaults

`func NewRmaActionInfoWithDefaults() *RmaActionInfo`

NewRmaActionInfoWithDefaults instantiates a new RmaActionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RmaActionInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RmaActionInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RmaActionInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RmaActionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RmaActionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RmaActionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RmaActionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RmaActionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInfo

`func (o *RmaActionInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *RmaActionInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *RmaActionInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *RmaActionInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


