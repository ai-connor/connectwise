# ReportCardDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Kpi** | [**KPIReference**](KPIReference.md) |  | 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**ReportCard** | Pointer to [**ReportCardReference**](ReportCardReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewReportCardDetail

`func NewReportCardDetail(kpi KPIReference, ) *ReportCardDetail`

NewReportCardDetail instantiates a new ReportCardDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCardDetailWithDefaults

`func NewReportCardDetailWithDefaults() *ReportCardDetail`

NewReportCardDetailWithDefaults instantiates a new ReportCardDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReportCardDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportCardDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportCardDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReportCardDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKpi

`func (o *ReportCardDetail) GetKpi() KPIReference`

GetKpi returns the Kpi field if non-nil, zero value otherwise.

### GetKpiOk

`func (o *ReportCardDetail) GetKpiOk() (*KPIReference, bool)`

GetKpiOk returns a tuple with the Kpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpi

`func (o *ReportCardDetail) SetKpi(v KPIReference)`

SetKpi sets Kpi field to given value.


### GetSortOrder

`func (o *ReportCardDetail) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ReportCardDetail) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ReportCardDetail) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ReportCardDetail) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *ReportCardDetail) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *ReportCardDetail) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetReportCard

`func (o *ReportCardDetail) GetReportCard() ReportCardReference`

GetReportCard returns the ReportCard field if non-nil, zero value otherwise.

### GetReportCardOk

`func (o *ReportCardDetail) GetReportCardOk() (*ReportCardReference, bool)`

GetReportCardOk returns a tuple with the ReportCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCard

`func (o *ReportCardDetail) SetReportCard(v ReportCardReference)`

SetReportCard sets ReportCard field to given value.

### HasReportCard

`func (o *ReportCardDetail) HasReportCard() bool`

HasReportCard returns a boolean if a field has been set.

### GetInfo

`func (o *ReportCardDetail) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ReportCardDetail) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ReportCardDetail) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ReportCardDetail) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


