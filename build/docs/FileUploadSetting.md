# FileUploadSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RestrictFileTypesFlag** | **NullableBool** |  | 
**GlobalFileSizeLimit** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewFileUploadSetting

`func NewFileUploadSetting(restrictFileTypesFlag NullableBool, ) *FileUploadSetting`

NewFileUploadSetting instantiates a new FileUploadSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUploadSettingWithDefaults

`func NewFileUploadSettingWithDefaults() *FileUploadSetting`

NewFileUploadSettingWithDefaults instantiates a new FileUploadSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileUploadSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileUploadSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileUploadSetting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FileUploadSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRestrictFileTypesFlag

`func (o *FileUploadSetting) GetRestrictFileTypesFlag() bool`

GetRestrictFileTypesFlag returns the RestrictFileTypesFlag field if non-nil, zero value otherwise.

### GetRestrictFileTypesFlagOk

`func (o *FileUploadSetting) GetRestrictFileTypesFlagOk() (*bool, bool)`

GetRestrictFileTypesFlagOk returns a tuple with the RestrictFileTypesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictFileTypesFlag

`func (o *FileUploadSetting) SetRestrictFileTypesFlag(v bool)`

SetRestrictFileTypesFlag sets RestrictFileTypesFlag field to given value.


### SetRestrictFileTypesFlagNil

`func (o *FileUploadSetting) SetRestrictFileTypesFlagNil(b bool)`

 SetRestrictFileTypesFlagNil sets the value for RestrictFileTypesFlag to be an explicit nil

### UnsetRestrictFileTypesFlag
`func (o *FileUploadSetting) UnsetRestrictFileTypesFlag()`

UnsetRestrictFileTypesFlag ensures that no value is present for RestrictFileTypesFlag, not even an explicit nil
### GetGlobalFileSizeLimit

`func (o *FileUploadSetting) GetGlobalFileSizeLimit() int32`

GetGlobalFileSizeLimit returns the GlobalFileSizeLimit field if non-nil, zero value otherwise.

### GetGlobalFileSizeLimitOk

`func (o *FileUploadSetting) GetGlobalFileSizeLimitOk() (*int32, bool)`

GetGlobalFileSizeLimitOk returns a tuple with the GlobalFileSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalFileSizeLimit

`func (o *FileUploadSetting) SetGlobalFileSizeLimit(v int32)`

SetGlobalFileSizeLimit sets GlobalFileSizeLimit field to given value.

### HasGlobalFileSizeLimit

`func (o *FileUploadSetting) HasGlobalFileSizeLimit() bool`

HasGlobalFileSizeLimit returns a boolean if a field has been set.

### SetGlobalFileSizeLimitNil

`func (o *FileUploadSetting) SetGlobalFileSizeLimitNil(b bool)`

 SetGlobalFileSizeLimitNil sets the value for GlobalFileSizeLimit to be an explicit nil

### UnsetGlobalFileSizeLimit
`func (o *FileUploadSetting) UnsetGlobalFileSizeLimit()`

UnsetGlobalFileSizeLimit ensures that no value is present for GlobalFileSizeLimit, not even an explicit nil
### GetInfo

`func (o *FileUploadSetting) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *FileUploadSetting) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *FileUploadSetting) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *FileUploadSetting) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


