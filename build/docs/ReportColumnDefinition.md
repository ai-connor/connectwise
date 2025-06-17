# ReportColumnDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**IsNullable** | Pointer to **bool** |  | [optional] 
**IdentityColumn** | Pointer to **bool** |  | [optional] 

## Methods

### NewReportColumnDefinition

`func NewReportColumnDefinition() *ReportColumnDefinition`

NewReportColumnDefinition instantiates a new ReportColumnDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportColumnDefinitionWithDefaults

`func NewReportColumnDefinitionWithDefaults() *ReportColumnDefinition`

NewReportColumnDefinitionWithDefaults instantiates a new ReportColumnDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReportColumnDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportColumnDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportColumnDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReportColumnDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsNullable

`func (o *ReportColumnDefinition) GetIsNullable() bool`

GetIsNullable returns the IsNullable field if non-nil, zero value otherwise.

### GetIsNullableOk

`func (o *ReportColumnDefinition) GetIsNullableOk() (*bool, bool)`

GetIsNullableOk returns a tuple with the IsNullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNullable

`func (o *ReportColumnDefinition) SetIsNullable(v bool)`

SetIsNullable sets IsNullable field to given value.

### HasIsNullable

`func (o *ReportColumnDefinition) HasIsNullable() bool`

HasIsNullable returns a boolean if a field has been set.

### GetIdentityColumn

`func (o *ReportColumnDefinition) GetIdentityColumn() bool`

GetIdentityColumn returns the IdentityColumn field if non-nil, zero value otherwise.

### GetIdentityColumnOk

`func (o *ReportColumnDefinition) GetIdentityColumnOk() (*bool, bool)`

GetIdentityColumnOk returns a tuple with the IdentityColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityColumn

`func (o *ReportColumnDefinition) SetIdentityColumn(v bool)`

SetIdentityColumn sets IdentityColumn field to given value.

### HasIdentityColumn

`func (o *ReportColumnDefinition) HasIdentityColumn() bool`

HasIdentityColumn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


