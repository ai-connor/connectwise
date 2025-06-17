# ReportCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewReportCard

`func NewReportCard(name string, ) *ReportCard`

NewReportCard instantiates a new ReportCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCardWithDefaults

`func NewReportCardWithDefaults() *ReportCard`

NewReportCardWithDefaults instantiates a new ReportCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReportCard) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportCard) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportCard) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReportCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ReportCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportCard) SetName(v string)`

SetName sets Name field to given value.


### GetInfo

`func (o *ReportCard) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ReportCard) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ReportCard) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ReportCard) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


