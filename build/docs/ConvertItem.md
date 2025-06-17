# ConvertItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RecordType** | **NullableString** |  | 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**Phase** | Pointer to [**ProjectPhaseReference**](ProjectPhaseReference.md) |  | [optional] 
**WbsCode** | Pointer to **string** |  | [optional] 

## Methods

### NewConvertItem

`func NewConvertItem(recordType NullableString, ) *ConvertItem`

NewConvertItem instantiates a new ConvertItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertItemWithDefaults

`func NewConvertItemWithDefaults() *ConvertItem`

NewConvertItemWithDefaults instantiates a new ConvertItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConvertItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConvertItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConvertItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConvertItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *ConvertItem) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ConvertItem) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ConvertItem) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### SetRecordTypeNil

`func (o *ConvertItem) SetRecordTypeNil(b bool)`

 SetRecordTypeNil sets the value for RecordType to be an explicit nil

### UnsetRecordType
`func (o *ConvertItem) UnsetRecordType()`

UnsetRecordType ensures that no value is present for RecordType, not even an explicit nil
### GetProject

`func (o *ConvertItem) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ConvertItem) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ConvertItem) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *ConvertItem) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPhase

`func (o *ConvertItem) GetPhase() ProjectPhaseReference`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ConvertItem) GetPhaseOk() (*ProjectPhaseReference, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ConvertItem) SetPhase(v ProjectPhaseReference)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *ConvertItem) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetWbsCode

`func (o *ConvertItem) GetWbsCode() string`

GetWbsCode returns the WbsCode field if non-nil, zero value otherwise.

### GetWbsCodeOk

`func (o *ConvertItem) GetWbsCodeOk() (*string, bool)`

GetWbsCodeOk returns a tuple with the WbsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbsCode

`func (o *ConvertItem) SetWbsCode(v string)`

SetWbsCode sets WbsCode field to given value.

### HasWbsCode

`func (o *ConvertItem) HasWbsCode() bool`

HasWbsCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


