# ConvertToProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RecordType** | Pointer to **NullableString** |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**Phase** | [**ProjectPhaseReference**](ProjectPhaseReference.md) |  | 
**WbsCode** | **string** |  | 

## Methods

### NewConvertToProject

`func NewConvertToProject(phase ProjectPhaseReference, wbsCode string, ) *ConvertToProject`

NewConvertToProject instantiates a new ConvertToProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertToProjectWithDefaults

`func NewConvertToProjectWithDefaults() *ConvertToProject`

NewConvertToProjectWithDefaults instantiates a new ConvertToProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConvertToProject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConvertToProject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConvertToProject) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConvertToProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *ConvertToProject) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ConvertToProject) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ConvertToProject) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ConvertToProject) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### SetRecordTypeNil

`func (o *ConvertToProject) SetRecordTypeNil(b bool)`

 SetRecordTypeNil sets the value for RecordType to be an explicit nil

### UnsetRecordType
`func (o *ConvertToProject) UnsetRecordType()`

UnsetRecordType ensures that no value is present for RecordType, not even an explicit nil
### GetProject

`func (o *ConvertToProject) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ConvertToProject) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ConvertToProject) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *ConvertToProject) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPhase

`func (o *ConvertToProject) GetPhase() ProjectPhaseReference`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ConvertToProject) GetPhaseOk() (*ProjectPhaseReference, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ConvertToProject) SetPhase(v ProjectPhaseReference)`

SetPhase sets Phase field to given value.


### GetWbsCode

`func (o *ConvertToProject) GetWbsCode() string`

GetWbsCode returns the WbsCode field if non-nil, zero value otherwise.

### GetWbsCodeOk

`func (o *ConvertToProject) GetWbsCodeOk() (*string, bool)`

GetWbsCodeOk returns a tuple with the WbsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbsCode

`func (o *ConvertToProject) SetWbsCode(v string)`

SetWbsCode sets WbsCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


