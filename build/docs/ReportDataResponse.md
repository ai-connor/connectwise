# ReportDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnDefinitions** | Pointer to [**[]map[string]ReportColumnDefinition**](map[string]ReportColumnDefinition.md) |  | [optional] 
**RowValues** | Pointer to **[][]map[string]interface{}** |  | [optional] 

## Methods

### NewReportDataResponse

`func NewReportDataResponse() *ReportDataResponse`

NewReportDataResponse instantiates a new ReportDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportDataResponseWithDefaults

`func NewReportDataResponseWithDefaults() *ReportDataResponse`

NewReportDataResponseWithDefaults instantiates a new ReportDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnDefinitions

`func (o *ReportDataResponse) GetColumnDefinitions() []map[string]ReportColumnDefinition`

GetColumnDefinitions returns the ColumnDefinitions field if non-nil, zero value otherwise.

### GetColumnDefinitionsOk

`func (o *ReportDataResponse) GetColumnDefinitionsOk() (*[]map[string]ReportColumnDefinition, bool)`

GetColumnDefinitionsOk returns a tuple with the ColumnDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDefinitions

`func (o *ReportDataResponse) SetColumnDefinitions(v []map[string]ReportColumnDefinition)`

SetColumnDefinitions sets ColumnDefinitions field to given value.

### HasColumnDefinitions

`func (o *ReportDataResponse) HasColumnDefinitions() bool`

HasColumnDefinitions returns a boolean if a field has been set.

### GetRowValues

`func (o *ReportDataResponse) GetRowValues() [][]map[string]interface{}`

GetRowValues returns the RowValues field if non-nil, zero value otherwise.

### GetRowValuesOk

`func (o *ReportDataResponse) GetRowValuesOk() (*[][]map[string]interface{}, bool)`

GetRowValuesOk returns a tuple with the RowValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowValues

`func (o *ReportDataResponse) SetRowValues(v [][]map[string]interface{})`

SetRowValues sets RowValues field to given value.

### HasRowValues

`func (o *ReportDataResponse) HasRowValues() bool`

HasRowValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


